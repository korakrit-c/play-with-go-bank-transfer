package handler

import (
	"bank-transfer/errs"
	"bank-transfer/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accService service.AccountService
}

func NewAccountHandler(accService service.AccountService) accountHandler {
	return accountHandler{accService: accService}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerID, err := strconv.Atoi(mux.Vars(r)["customerID"])
	if err != nil {
		handleError(w, err)
		return
	}

	if r.Header.Get("content-type") != "application/json" {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	request := service.NewAccountRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, errs.NewValidationError("request body incorrect formant"))
		return
	}

	resposne, err := h.accService.NewAccount(customerID, request)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(resposne)

}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	customerID, err := strconv.Atoi(mux.Vars(r)["customerID"])
	if err != nil {
		handleError(w, err)
		return
	}

	responses, err := h.accService.GetAccounts(customerID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(responses)
}
