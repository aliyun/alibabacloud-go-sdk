// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddGatewayDomainResponseBody
	GetCode() *int32
	SetData(v int64) *AddGatewayDomainResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *AddGatewayDomainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddGatewayDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddGatewayDomainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddGatewayDomainResponseBody
	GetSuccess() *bool
}

type AddGatewayDomainResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 100
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
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 16BA802F-B848-55DF-9F57-CD71B7272D1F
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

func (s AddGatewayDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddGatewayDomainResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddGatewayDomainResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddGatewayDomainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddGatewayDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddGatewayDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGatewayDomainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddGatewayDomainResponseBody) SetCode(v int32) *AddGatewayDomainResponseBody {
	s.Code = &v
	return s
}

func (s *AddGatewayDomainResponseBody) SetData(v int64) *AddGatewayDomainResponseBody {
	s.Data = &v
	return s
}

func (s *AddGatewayDomainResponseBody) SetHttpStatusCode(v int32) *AddGatewayDomainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddGatewayDomainResponseBody) SetMessage(v string) *AddGatewayDomainResponseBody {
	s.Message = &v
	return s
}

func (s *AddGatewayDomainResponseBody) SetRequestId(v string) *AddGatewayDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGatewayDomainResponseBody) SetSuccess(v bool) *AddGatewayDomainResponseBody {
	s.Success = &v
	return s
}

func (s *AddGatewayDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
