// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceCronScalerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateServiceCronScalerResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateServiceCronScalerResponseBody
	GetRequestId() *string
}

type CreateServiceCronScalerResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Create cron scaler for service [foo] successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceCronScalerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceCronScalerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceCronScalerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateServiceCronScalerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceCronScalerResponseBody) SetMessage(v string) *CreateServiceCronScalerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceCronScalerResponseBody) SetRequestId(v string) *CreateServiceCronScalerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceCronScalerResponseBody) Validate() error {
	return dara.Validate(s)
}
