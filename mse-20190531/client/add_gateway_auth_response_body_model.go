// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddGatewayAuthResponseBody
	GetCode() *int32
	SetData(v int64) *AddGatewayAuthResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *AddGatewayAuthResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddGatewayAuthResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddGatewayAuthResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddGatewayAuthResponseBody
	GetSuccess() *bool
}

type AddGatewayAuthResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 3
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4279C00F-A5E1-53C6-A43B-751C1C524D0B
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

func (s AddGatewayAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayAuthResponseBody) GoString() string {
	return s.String()
}

func (s *AddGatewayAuthResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddGatewayAuthResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddGatewayAuthResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddGatewayAuthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddGatewayAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGatewayAuthResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddGatewayAuthResponseBody) SetCode(v int32) *AddGatewayAuthResponseBody {
	s.Code = &v
	return s
}

func (s *AddGatewayAuthResponseBody) SetData(v int64) *AddGatewayAuthResponseBody {
	s.Data = &v
	return s
}

func (s *AddGatewayAuthResponseBody) SetHttpStatusCode(v int32) *AddGatewayAuthResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddGatewayAuthResponseBody) SetMessage(v string) *AddGatewayAuthResponseBody {
	s.Message = &v
	return s
}

func (s *AddGatewayAuthResponseBody) SetRequestId(v string) *AddGatewayAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGatewayAuthResponseBody) SetSuccess(v bool) *AddGatewayAuthResponseBody {
	s.Success = &v
	return s
}

func (s *AddGatewayAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
