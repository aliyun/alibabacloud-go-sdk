// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewaySpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewaySpecResponseBody
	GetCode() *int32
	SetData(v string) *UpdateGatewaySpecResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateGatewaySpecResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewaySpecResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewaySpecResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewaySpecResponseBody
	GetSuccess() *bool
}

type UpdateGatewaySpecResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// gw-892ehbv7gg56******
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 316F5F64-F73D-42DC-8632-01E308B6****
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

func (s UpdateGatewaySpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewaySpecResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewaySpecResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewaySpecResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateGatewaySpecResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewaySpecResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewaySpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewaySpecResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewaySpecResponseBody) SetCode(v int32) *UpdateGatewaySpecResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewaySpecResponseBody) SetData(v string) *UpdateGatewaySpecResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewaySpecResponseBody) SetHttpStatusCode(v int32) *UpdateGatewaySpecResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewaySpecResponseBody) SetMessage(v string) *UpdateGatewaySpecResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewaySpecResponseBody) SetRequestId(v string) *UpdateGatewaySpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewaySpecResponseBody) SetSuccess(v bool) *UpdateGatewaySpecResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewaySpecResponseBody) Validate() error {
	return dara.Validate(s)
}
