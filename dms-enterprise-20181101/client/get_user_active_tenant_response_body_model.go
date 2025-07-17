// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserActiveTenantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetUserActiveTenantResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetUserActiveTenantResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetUserActiveTenantResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserActiveTenantResponseBody
	GetSuccess() *bool
	SetTenant(v *GetUserActiveTenantResponseBodyTenant) *GetUserActiveTenantResponseBody
	GetTenant() *GetUserActiveTenantResponseBodyTenant
}

type GetUserActiveTenantResponseBody struct {
	// The error code.
	//
	// example:
	//
	// TenantNotExist
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified tenant does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4B63CAC5-BD7F-5C7C-82C9-59DFFBC3C5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The details of the tenant.
	Tenant *GetUserActiveTenantResponseBodyTenant `json:"Tenant,omitempty" xml:"Tenant,omitempty" type:"Struct"`
}

func (s GetUserActiveTenantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserActiveTenantResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserActiveTenantResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetUserActiveTenantResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetUserActiveTenantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserActiveTenantResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserActiveTenantResponseBody) GetTenant() *GetUserActiveTenantResponseBodyTenant {
	return s.Tenant
}

func (s *GetUserActiveTenantResponseBody) SetErrorCode(v string) *GetUserActiveTenantResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserActiveTenantResponseBody) SetErrorMessage(v string) *GetUserActiveTenantResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserActiveTenantResponseBody) SetRequestId(v string) *GetUserActiveTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserActiveTenantResponseBody) SetSuccess(v bool) *GetUserActiveTenantResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserActiveTenantResponseBody) SetTenant(v *GetUserActiveTenantResponseBodyTenant) *GetUserActiveTenantResponseBody {
	s.Tenant = v
	return s
}

func (s *GetUserActiveTenantResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserActiveTenantResponseBodyTenant struct {
	// The status of the tenant. Valid values:
	//
	// 	- **ACTIVE**: The tenant is used to access DMS.
	//
	// 	- **IN_ACTIVE**: The tenant is not used.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the tenant.
	//
	// example:
	//
	// test_name
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetUserActiveTenantResponseBodyTenant) String() string {
	return dara.Prettify(s)
}

func (s GetUserActiveTenantResponseBodyTenant) GoString() string {
	return s.String()
}

func (s *GetUserActiveTenantResponseBodyTenant) GetStatus() *string {
	return s.Status
}

func (s *GetUserActiveTenantResponseBodyTenant) GetTenantName() *string {
	return s.TenantName
}

func (s *GetUserActiveTenantResponseBodyTenant) GetTid() *int64 {
	return s.Tid
}

func (s *GetUserActiveTenantResponseBodyTenant) SetStatus(v string) *GetUserActiveTenantResponseBodyTenant {
	s.Status = &v
	return s
}

func (s *GetUserActiveTenantResponseBodyTenant) SetTenantName(v string) *GetUserActiveTenantResponseBodyTenant {
	s.TenantName = &v
	return s
}

func (s *GetUserActiveTenantResponseBodyTenant) SetTid(v int64) *GetUserActiveTenantResponseBodyTenant {
	s.Tid = &v
	return s
}

func (s *GetUserActiveTenantResponseBodyTenant) Validate() error {
	return dara.Validate(s)
}
