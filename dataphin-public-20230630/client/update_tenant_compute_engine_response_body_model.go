// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantComputeEngineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTenantComputeEngineResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateTenantComputeEngineResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateTenantComputeEngineResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTenantComputeEngineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTenantComputeEngineResponseBody
	GetSuccess() *bool
}

type UpdateTenantComputeEngineResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTenantComputeEngineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantComputeEngineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTenantComputeEngineResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTenantComputeEngineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateTenantComputeEngineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTenantComputeEngineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTenantComputeEngineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTenantComputeEngineResponseBody) SetCode(v string) *UpdateTenantComputeEngineResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTenantComputeEngineResponseBody) SetHttpStatusCode(v int32) *UpdateTenantComputeEngineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTenantComputeEngineResponseBody) SetMessage(v string) *UpdateTenantComputeEngineResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTenantComputeEngineResponseBody) SetRequestId(v string) *UpdateTenantComputeEngineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTenantComputeEngineResponseBody) SetSuccess(v bool) *UpdateTenantComputeEngineResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTenantComputeEngineResponseBody) Validate() error {
	return dara.Validate(s)
}
