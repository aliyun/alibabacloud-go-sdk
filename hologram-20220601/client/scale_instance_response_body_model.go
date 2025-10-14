// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ScaleInstanceResponseBodyData) *ScaleInstanceResponseBody
	GetData() *ScaleInstanceResponseBodyData
	SetErrorCode(v string) *ScaleInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ScaleInstanceResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *ScaleInstanceResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *ScaleInstanceResponseBody
	GetRequestId() *string
}

type ScaleInstanceResponseBody struct {
	// The returned data.
	Data *ScaleInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
}

func (s ScaleInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleInstanceResponseBody) GetData() *ScaleInstanceResponseBodyData {
	return s.Data
}

func (s *ScaleInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ScaleInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ScaleInstanceResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ScaleInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScaleInstanceResponseBody) SetData(v *ScaleInstanceResponseBodyData) *ScaleInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ScaleInstanceResponseBody) SetErrorCode(v string) *ScaleInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ScaleInstanceResponseBody) SetErrorMessage(v string) *ScaleInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ScaleInstanceResponseBody) SetHttpStatusCode(v string) *ScaleInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ScaleInstanceResponseBody) SetRequestId(v string) *ScaleInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ScaleInstanceResponseBodyData struct {
	// The error code returned.
	//
	// example:
	//
	// InvalidScaleType.Unsupported
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error details.
	//
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 219183853450000
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Indicates whether the change to specifications was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ScaleInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ScaleInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ScaleInstanceResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *ScaleInstanceResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ScaleInstanceResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *ScaleInstanceResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ScaleInstanceResponseBodyData) SetCode(v string) *ScaleInstanceResponseBodyData {
	s.Code = &v
	return s
}

func (s *ScaleInstanceResponseBodyData) SetMessage(v string) *ScaleInstanceResponseBodyData {
	s.Message = &v
	return s
}

func (s *ScaleInstanceResponseBodyData) SetOrderId(v string) *ScaleInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ScaleInstanceResponseBodyData) SetSuccess(v bool) *ScaleInstanceResponseBodyData {
	s.Success = &v
	return s
}

func (s *ScaleInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
