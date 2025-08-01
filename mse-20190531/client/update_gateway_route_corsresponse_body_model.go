// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteCORSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayRouteCORSResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateGatewayRouteCORSResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateGatewayRouteCORSResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayRouteCORSResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayRouteCORSResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayRouteCORSResponseBody
	GetSuccess() *bool
}

type UpdateGatewayRouteCORSResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data structure.
	//
	// example:
	//
	// 618
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 41BE990E-44E4-5ED3-AEEB-A276941B1854
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayRouteCORSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteCORSResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteCORSResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayRouteCORSResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateGatewayRouteCORSResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayRouteCORSResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayRouteCORSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayRouteCORSResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayRouteCORSResponseBody) SetCode(v int32) *UpdateGatewayRouteCORSResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayRouteCORSResponseBody) SetData(v int64) *UpdateGatewayRouteCORSResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayRouteCORSResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayRouteCORSResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayRouteCORSResponseBody) SetMessage(v string) *UpdateGatewayRouteCORSResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayRouteCORSResponseBody) SetRequestId(v string) *UpdateGatewayRouteCORSResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayRouteCORSResponseBody) SetSuccess(v bool) *UpdateGatewayRouteCORSResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayRouteCORSResponseBody) Validate() error {
	return dara.Validate(s)
}
