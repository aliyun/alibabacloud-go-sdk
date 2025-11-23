// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserTenantsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListUserTenantsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListUserTenantsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListUserTenantsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUserTenantsResponseBody
	GetSuccess() *bool
	SetTenantList(v []*ListUserTenantsResponseBodyTenantList) *ListUserTenantsResponseBody
	GetTenantList() []*ListUserTenantsResponseBodyTenantList
}

type ListUserTenantsResponseBody struct {
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
	// C478D9DA-3615-50F6-A2BC-7855AD65****
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
	// The details of the tenants that were returned.
	TenantList []*ListUserTenantsResponseBodyTenantList `json:"TenantList,omitempty" xml:"TenantList,omitempty" type:"Repeated"`
}

func (s ListUserTenantsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserTenantsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserTenantsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListUserTenantsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListUserTenantsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserTenantsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserTenantsResponseBody) GetTenantList() []*ListUserTenantsResponseBodyTenantList {
	return s.TenantList
}

func (s *ListUserTenantsResponseBody) SetErrorCode(v string) *ListUserTenantsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUserTenantsResponseBody) SetErrorMessage(v string) *ListUserTenantsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListUserTenantsResponseBody) SetRequestId(v string) *ListUserTenantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserTenantsResponseBody) SetSuccess(v bool) *ListUserTenantsResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserTenantsResponseBody) SetTenantList(v []*ListUserTenantsResponseBodyTenantList) *ListUserTenantsResponseBody {
	s.TenantList = v
	return s
}

func (s *ListUserTenantsResponseBody) Validate() error {
	if s.TenantList != nil {
		for _, item := range s.TenantList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserTenantsResponseBodyTenantList struct {
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

func (s ListUserTenantsResponseBodyTenantList) String() string {
	return dara.Prettify(s)
}

func (s ListUserTenantsResponseBodyTenantList) GoString() string {
	return s.String()
}

func (s *ListUserTenantsResponseBodyTenantList) GetStatus() *string {
	return s.Status
}

func (s *ListUserTenantsResponseBodyTenantList) GetTenantName() *string {
	return s.TenantName
}

func (s *ListUserTenantsResponseBodyTenantList) GetTid() *int64 {
	return s.Tid
}

func (s *ListUserTenantsResponseBodyTenantList) SetStatus(v string) *ListUserTenantsResponseBodyTenantList {
	s.Status = &v
	return s
}

func (s *ListUserTenantsResponseBodyTenantList) SetTenantName(v string) *ListUserTenantsResponseBodyTenantList {
	s.TenantName = &v
	return s
}

func (s *ListUserTenantsResponseBodyTenantList) SetTid(v int64) *ListUserTenantsResponseBodyTenantList {
	s.Tid = &v
	return s
}

func (s *ListUserTenantsResponseBodyTenantList) Validate() error {
	return dara.Validate(s)
}
