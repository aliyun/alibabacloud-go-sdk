// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayDomainResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateGatewayDomainResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateGatewayDomainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayDomainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayDomainResponseBody
	GetSuccess() *bool
}

type UpdateGatewayDomainResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 403
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 94
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
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6F6C6DE4-DB33-5791-B210-ED2E6FEFFE6F
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

func (s UpdateGatewayDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayDomainResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayDomainResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateGatewayDomainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayDomainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayDomainResponseBody) SetCode(v int32) *UpdateGatewayDomainResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayDomainResponseBody) SetData(v int64) *UpdateGatewayDomainResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayDomainResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayDomainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayDomainResponseBody) SetMessage(v string) *UpdateGatewayDomainResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayDomainResponseBody) SetRequestId(v string) *UpdateGatewayDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayDomainResponseBody) SetSuccess(v bool) *UpdateGatewayDomainResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
