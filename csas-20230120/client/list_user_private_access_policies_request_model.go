// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPrivateAccessPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListUserPrivateAccessPoliciesRequest
	GetCurrentPage() *int32
	SetName(v string) *ListUserPrivateAccessPoliciesRequest
	GetName() *string
	SetPageSize(v int32) *ListUserPrivateAccessPoliciesRequest
	GetPageSize() *int32
	SetSaseUserId(v string) *ListUserPrivateAccessPoliciesRequest
	GetSaseUserId() *string
}

type ListUserPrivateAccessPoliciesRequest struct {
	// Current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Intranet access policy name. Length should be between 1 to 128 characters, supporting Chinese and case-sensitive English letters, and can include numbers, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Number of items per page for pagination. Range: 1~100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// User ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
}

func (s ListUserPrivateAccessPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserPrivateAccessPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListUserPrivateAccessPoliciesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUserPrivateAccessPoliciesRequest) GetName() *string {
	return s.Name
}

func (s *ListUserPrivateAccessPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserPrivateAccessPoliciesRequest) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListUserPrivateAccessPoliciesRequest) SetCurrentPage(v int32) *ListUserPrivateAccessPoliciesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesRequest) SetName(v string) *ListUserPrivateAccessPoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesRequest) SetPageSize(v int32) *ListUserPrivateAccessPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesRequest) SetSaseUserId(v string) *ListUserPrivateAccessPoliciesRequest {
	s.SaseUserId = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
