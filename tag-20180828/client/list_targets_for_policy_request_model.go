// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTargetsForPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResult(v int32) *ListTargetsForPolicyRequest
	GetMaxResult() *int32
	SetNextToken(v string) *ListTargetsForPolicyRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTargetsForPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTargetsForPolicyRequest
	GetOwnerId() *int64
	SetPolicyId(v string) *ListTargetsForPolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *ListTargetsForPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTargetsForPolicyRequest
	GetResourceOwnerAccount() *string
}

type ListTargetsForPolicyRequest struct {
	// The number of entries to return on each page.
	//
	// Default value: 50. Maximum value: 1000.
	//
	// example:
	//
	// 50
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the tag policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// p-de62a0bf400e4b69****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s ListTargetsForPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTargetsForPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListTargetsForPolicyRequest) GetMaxResult() *int32 {
	return s.MaxResult
}

func (s *ListTargetsForPolicyRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTargetsForPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTargetsForPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTargetsForPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListTargetsForPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTargetsForPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTargetsForPolicyRequest) SetMaxResult(v int32) *ListTargetsForPolicyRequest {
	s.MaxResult = &v
	return s
}

func (s *ListTargetsForPolicyRequest) SetNextToken(v string) *ListTargetsForPolicyRequest {
	s.NextToken = &v
	return s
}

func (s *ListTargetsForPolicyRequest) SetOwnerAccount(v string) *ListTargetsForPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTargetsForPolicyRequest) SetOwnerId(v int64) *ListTargetsForPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTargetsForPolicyRequest) SetPolicyId(v string) *ListTargetsForPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ListTargetsForPolicyRequest) SetRegionId(v string) *ListTargetsForPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ListTargetsForPolicyRequest) SetResourceOwnerAccount(v string) *ListTargetsForPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTargetsForPolicyRequest) Validate() error {
	return dara.Validate(s)
}
