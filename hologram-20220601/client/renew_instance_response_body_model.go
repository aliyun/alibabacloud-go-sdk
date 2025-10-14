// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RenewInstanceResponseBodyData) *RenewInstanceResponseBody
	GetData() *RenewInstanceResponseBodyData
	SetErrorCode(v string) *RenewInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RenewInstanceResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *RenewInstanceResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *RenewInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *RenewInstanceResponseBody
	GetSuccess() *string
}

type RenewInstanceResponseBody struct {
	// The returned data.
	Data *RenewInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D3AE84AB-0873-5FC7-A4C4-8CF869D2FA70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result, which indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) GetData() *RenewInstanceResponseBodyData {
	return s.Data
}

func (s *RenewInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RenewInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RenewInstanceResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *RenewInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewInstanceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *RenewInstanceResponseBody) SetData(v *RenewInstanceResponseBodyData) *RenewInstanceResponseBody {
	s.Data = v
	return s
}

func (s *RenewInstanceResponseBody) SetErrorCode(v string) *RenewInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RenewInstanceResponseBody) SetErrorMessage(v string) *RenewInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RenewInstanceResponseBody) SetHttpStatusCode(v string) *RenewInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetSuccess(v string) *RenewInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RenewInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RenewInstanceResponseBodyData struct {
	// The error code returned.
	//
	// example:
	//
	// InvalidChargeType.UnRenewable
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error details.
	//
	// example:
	//
	// InvalidChargeType.UnRenewable
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 221625608580893
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Indicates whether the renewal was successful.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *RenewInstanceResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *RenewInstanceResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewInstanceResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *RenewInstanceResponseBodyData) SetCode(v string) *RenewInstanceResponseBodyData {
	s.Code = &v
	return s
}

func (s *RenewInstanceResponseBodyData) SetMessage(v string) *RenewInstanceResponseBodyData {
	s.Message = &v
	return s
}

func (s *RenewInstanceResponseBodyData) SetOrderId(v string) *RenewInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *RenewInstanceResponseBodyData) SetSuccess(v bool) *RenewInstanceResponseBodyData {
	s.Success = &v
	return s
}

func (s *RenewInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
