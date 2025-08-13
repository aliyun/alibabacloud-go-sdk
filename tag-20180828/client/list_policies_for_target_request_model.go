// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesForTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResult(v int32) *ListPoliciesForTargetRequest
	GetMaxResult() *int32
	SetNextToken(v string) *ListPoliciesForTargetRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListPoliciesForTargetRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListPoliciesForTargetRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListPoliciesForTargetRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListPoliciesForTargetRequest
	GetResourceOwnerAccount() *string
	SetTargetId(v string) *ListPoliciesForTargetRequest
	GetTargetId() *string
	SetTargetType(v string) *ListPoliciesForTargetRequest
	GetTargetType() *string
}

type ListPoliciesForTargetRequest struct {
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
	// The region ID. Set the value to cn-shanghai.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the object. This parameter specifies a filter condition for the query.
	//
	// example:
	//
	// 154950938137****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. This parameter specifies a filter condition for the query. Valid values:
	//
	// 	- USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	//
	// 	- ROOT: the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// 	- FOLDER: a folder other than the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// 	- ACCOUNT: a member in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// >  The value of this parameter is not case-sensitive.
	//
	// example:
	//
	// ACCOUNT
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListPoliciesForTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForTargetRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesForTargetRequest) GetMaxResult() *int32 {
	return s.MaxResult
}

func (s *ListPoliciesForTargetRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPoliciesForTargetRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListPoliciesForTargetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListPoliciesForTargetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPoliciesForTargetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListPoliciesForTargetRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *ListPoliciesForTargetRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListPoliciesForTargetRequest) SetMaxResult(v int32) *ListPoliciesForTargetRequest {
	s.MaxResult = &v
	return s
}

func (s *ListPoliciesForTargetRequest) SetNextToken(v string) *ListPoliciesForTargetRequest {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesForTargetRequest) SetOwnerAccount(v string) *ListPoliciesForTargetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListPoliciesForTargetRequest) SetOwnerId(v int64) *ListPoliciesForTargetRequest {
	s.OwnerId = &v
	return s
}

func (s *ListPoliciesForTargetRequest) SetRegionId(v string) *ListPoliciesForTargetRequest {
	s.RegionId = &v
	return s
}

func (s *ListPoliciesForTargetRequest) SetResourceOwnerAccount(v string) *ListPoliciesForTargetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListPoliciesForTargetRequest) SetTargetId(v string) *ListPoliciesForTargetRequest {
	s.TargetId = &v
	return s
}

func (s *ListPoliciesForTargetRequest) SetTargetType(v string) *ListPoliciesForTargetRequest {
	s.TargetType = &v
	return s
}

func (s *ListPoliciesForTargetRequest) Validate() error {
	return dara.Validate(s)
}
