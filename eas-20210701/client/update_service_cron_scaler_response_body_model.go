// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceCronScalerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateServiceCronScalerResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceCronScalerResponseBody
	GetRequestId() *string
}

type UpdateServiceCronScalerResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Update cron scaler for service [foo] successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceCronScalerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceCronScalerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceCronScalerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceCronScalerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceCronScalerResponseBody) SetMessage(v string) *UpdateServiceCronScalerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceCronScalerResponseBody) SetRequestId(v string) *UpdateServiceCronScalerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceCronScalerResponseBody) Validate() error {
	return dara.Validate(s)
}
