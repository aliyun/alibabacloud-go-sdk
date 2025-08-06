// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListAppPolicyRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *ListAppPolicyRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListAppPolicyRequest
	GetPageSize() *int64
	SetPolicyType(v string) *ListAppPolicyRequest
	GetPolicyType() *string
	SetResourceOwnerAccount(v string) *ListAppPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListAppPolicyRequest
	GetResourceOwnerId() *int64
}

type ListAppPolicyRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PolicyType           *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAppPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListAppPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAppPolicyRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListAppPolicyRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAppPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListAppPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAppPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAppPolicyRequest) SetOwnerId(v int64) *ListAppPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAppPolicyRequest) SetPageNo(v int64) *ListAppPolicyRequest {
	s.PageNo = &v
	return s
}

func (s *ListAppPolicyRequest) SetPageSize(v int64) *ListAppPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppPolicyRequest) SetPolicyType(v string) *ListAppPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *ListAppPolicyRequest) SetResourceOwnerAccount(v string) *ListAppPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAppPolicyRequest) SetResourceOwnerId(v int64) *ListAppPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAppPolicyRequest) Validate() error {
	return dara.Validate(s)
}
