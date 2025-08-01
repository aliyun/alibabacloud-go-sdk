// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayServiceResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateGatewayServiceResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateGatewayServiceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayServiceResponseBody
	GetSuccess() *bool
}

type UpdateGatewayServiceResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 322
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
	// AF21683A-29C7-4853-AC0F-B5ADEE4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayServiceResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateGatewayServiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayServiceResponseBody) SetCode(v int32) *UpdateGatewayServiceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayServiceResponseBody) SetData(v int64) *UpdateGatewayServiceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayServiceResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayServiceResponseBody) SetMessage(v string) *UpdateGatewayServiceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayServiceResponseBody) SetRequestId(v string) *UpdateGatewayServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayServiceResponseBody) SetSuccess(v bool) *UpdateGatewayServiceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
