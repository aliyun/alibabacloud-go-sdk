// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyName(v string) *ListTenantAppRequest
	GetKeyName() *string
	SetPageNumber(v int32) *ListTenantAppRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTenantAppRequest
	GetPageSize() *int32
	SetSourceType(v string) *ListTenantAppRequest
	GetSourceType() *string
}

type ListTenantAppRequest struct {
	// The application name keyword, matched by containment. If this parameter is not specified or is set to an empty string, no name-based filtering is applied. `%` can be used as a wildcard, and `_` is matched as a literal character.
	//
	// example:
	//
	// SampleEditor
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The page number, starting from 1. If this parameter is not specified or is set to a value less than or equal to 0, the value 1 is used.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 500. If this parameter is not specified, is set to a value less than or equal to 0, or is set to a value greater than 500, the value 20 is used.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The application source. Valid values:
	//
	// - MARKET: marketplace applications.
	//
	// - TENANT: applications uploaded by the current tenant.
	//
	// If this parameter is not specified, both types of visible applications are queried.
	//
	// example:
	//
	// TENANT
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListTenantAppRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTenantAppRequest) GoString() string {
	return s.String()
}

func (s *ListTenantAppRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *ListTenantAppRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTenantAppRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTenantAppRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListTenantAppRequest) SetKeyName(v string) *ListTenantAppRequest {
	s.KeyName = &v
	return s
}

func (s *ListTenantAppRequest) SetPageNumber(v int32) *ListTenantAppRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTenantAppRequest) SetPageSize(v int32) *ListTenantAppRequest {
	s.PageSize = &v
	return s
}

func (s *ListTenantAppRequest) SetSourceType(v string) *ListTenantAppRequest {
	s.SourceType = &v
	return s
}

func (s *ListTenantAppRequest) Validate() error {
	return dara.Validate(s)
}
