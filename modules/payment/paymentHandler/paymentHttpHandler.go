package paymentHandler

import (
	"github.com/Rayato159/hello-sekai-shop-tutorial/config"
	"github.com/Rayato159/hello-sekai-shop-tutorial/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandlerService interface{}

	paymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentHttpHandlerService {
	return &paymentHttpHandler{
		cfg:            cfg,
		paymentUsecase: paymentUsecase,
	}
}
