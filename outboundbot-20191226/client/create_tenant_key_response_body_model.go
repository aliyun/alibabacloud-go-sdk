// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTenantKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTenantKeyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateTenantKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTenantKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTenantKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTenantKeyResponseBody
	GetSuccess() *bool
}

type CreateTenantKeyResponseBody struct {
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
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTenantKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTenantKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTenantKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTenantKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTenantKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTenantKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTenantKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTenantKeyResponseBody) SetCode(v string) *CreateTenantKeyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTenantKeyResponseBody) SetHttpStatusCode(v int32) *CreateTenantKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTenantKeyResponseBody) SetMessage(v string) *CreateTenantKeyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTenantKeyResponseBody) SetRequestId(v string) *CreateTenantKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTenantKeyResponseBody) SetSuccess(v bool) *CreateTenantKeyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTenantKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
