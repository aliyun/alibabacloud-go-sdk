// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomerLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCustomerLabelResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteCustomerLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomerLabelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomerLabelResponseBody
	GetSuccess() *bool
}

type DeleteCustomerLabelResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomerLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomerLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomerLabelResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCustomerLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomerLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomerLabelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomerLabelResponseBody) SetCode(v string) *DeleteCustomerLabelResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomerLabelResponseBody) SetMessage(v string) *DeleteCustomerLabelResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomerLabelResponseBody) SetRequestId(v string) *DeleteCustomerLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomerLabelResponseBody) SetSuccess(v bool) *DeleteCustomerLabelResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomerLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
