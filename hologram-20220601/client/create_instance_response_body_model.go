// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateInstanceResponseBodyData) *CreateInstanceResponseBody
	GetData() *CreateInstanceResponseBodyData
	SetErrorCode(v string) *CreateInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateInstanceResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *CreateInstanceResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *CreateInstanceResponseBody
	GetRequestId() *string
}

type CreateInstanceResponseBody struct {
	// The returned data.
	Data *CreateInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 9CC37B9F-F4B4-5FF1-939B-AEE78DC70130
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) GetData() *CreateInstanceResponseBodyData {
	return s.Data
}

func (s *CreateInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateInstanceResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceResponseBody) SetData(v *CreateInstanceResponseBodyData) *CreateInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateInstanceResponseBody) SetErrorCode(v string) *CreateInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateInstanceResponseBody) SetErrorMessage(v string) *CreateInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateInstanceResponseBody) SetHttpStatusCode(v string) *CreateInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInstanceResponseBodyData struct {
	// The error code returned.
	//
	// example:
	//
	// InvalidVpcOrVSwitch.NotAvailable
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// hgpostcn-cn-xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The error details.
	//
	// example:
	//
	// Vpc is not available
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 217523224780172
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Indicates whether the instance was created.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CreateInstanceResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateInstanceResponseBodyData) GetSuccess() *string {
	return s.Success
}

func (s *CreateInstanceResponseBodyData) SetCode(v string) *CreateInstanceResponseBodyData {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetInstanceId(v string) *CreateInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetMessage(v string) *CreateInstanceResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetOrderId(v string) *CreateInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetSuccess(v string) *CreateInstanceResponseBodyData {
	s.Success = &v
	return s
}

func (s *CreateInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
