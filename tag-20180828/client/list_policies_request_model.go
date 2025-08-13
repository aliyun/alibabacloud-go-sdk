// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResult(v int32) *ListPoliciesRequest
	GetMaxResult() *int32
	SetNextToken(v string) *ListPoliciesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListPoliciesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListPoliciesRequest
	GetOwnerId() *int64
	SetPolicyIds(v []*string) *ListPoliciesRequest
	GetPolicyIds() []*string
	SetPolicyNames(v []*string) *ListPoliciesRequest
	GetPolicyNames() []*string
	SetRegionId(v string) *ListPoliciesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListPoliciesRequest
	GetResourceOwnerAccount() *string
	SetUserType(v string) *ListPoliciesRequest
	GetUserType() *string
}

type ListPoliciesRequest struct {
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
	// The ID of a tag policy. This parameter specifies a filter condition for the query.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The name of a tag policy. This parameter specifies a filter condition for the query.
	PolicyNames []*string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty" type:"Repeated"`
	// The region ID. Set the value to cn-shanghai.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The mode of the Tag Policy feature. This parameter specifies a filter condition for the query. Valid values:
	//
	// 	- USER: single-account mode
	//
	// 	- RD: multi-account mode
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](https://help.aliyun.com/document_detail/417434.html).
	//
	// >  The value of this parameter is not case-sensitive.
	//
	// example:
	//
	// USER
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s ListPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) GetMaxResult() *int32 {
	return s.MaxResult
}

func (s *ListPoliciesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPoliciesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListPoliciesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListPoliciesRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListPoliciesRequest) GetPolicyNames() []*string {
	return s.PolicyNames
}

func (s *ListPoliciesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPoliciesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListPoliciesRequest) GetUserType() *string {
	return s.UserType
}

func (s *ListPoliciesRequest) SetMaxResult(v int32) *ListPoliciesRequest {
	s.MaxResult = &v
	return s
}

func (s *ListPoliciesRequest) SetNextToken(v string) *ListPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesRequest) SetOwnerAccount(v string) *ListPoliciesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListPoliciesRequest) SetOwnerId(v int64) *ListPoliciesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListPoliciesRequest) SetPolicyIds(v []*string) *ListPoliciesRequest {
	s.PolicyIds = v
	return s
}

func (s *ListPoliciesRequest) SetPolicyNames(v []*string) *ListPoliciesRequest {
	s.PolicyNames = v
	return s
}

func (s *ListPoliciesRequest) SetRegionId(v string) *ListPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPoliciesRequest) SetResourceOwnerAccount(v string) *ListPoliciesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListPoliciesRequest) SetUserType(v string) *ListPoliciesRequest {
	s.UserType = &v
	return s
}

func (s *ListPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
