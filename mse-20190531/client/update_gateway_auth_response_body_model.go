// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayAuthResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateGatewayAuthResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateGatewayAuthResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayAuthResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayAuthResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayAuthResponseBody
	GetSuccess() *bool
}

type UpdateGatewayAuthResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 719
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 316F5F64-F73D-42DC-8632-01E308B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayAuthResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateGatewayAuthResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayAuthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayAuthResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayAuthResponseBody) SetCode(v int32) *UpdateGatewayAuthResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayAuthResponseBody) SetData(v int64) *UpdateGatewayAuthResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayAuthResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayAuthResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayAuthResponseBody) SetMessage(v string) *UpdateGatewayAuthResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayAuthResponseBody) SetRequestId(v string) *UpdateGatewayAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayAuthResponseBody) SetSuccess(v bool) *UpdateGatewayAuthResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
