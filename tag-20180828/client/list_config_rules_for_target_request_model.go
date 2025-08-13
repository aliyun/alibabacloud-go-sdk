// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigRulesForTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResult(v int32) *ListConfigRulesForTargetRequest
	GetMaxResult() *int32
	SetNextToken(v string) *ListConfigRulesForTargetRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListConfigRulesForTargetRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListConfigRulesForTargetRequest
	GetOwnerId() *int64
	SetPolicyType(v string) *ListConfigRulesForTargetRequest
	GetPolicyType() *string
	SetRegionId(v string) *ListConfigRulesForTargetRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListConfigRulesForTargetRequest
	GetResourceOwnerAccount() *string
	SetTagKey(v string) *ListConfigRulesForTargetRequest
	GetTagKey() *string
	SetTargetId(v string) *ListConfigRulesForTargetRequest
	GetTargetId() *string
	SetTargetType(v string) *ListConfigRulesForTargetRequest
	GetTargetType() *string
	SetUserType(v string) *ListConfigRulesForTargetRequest
	GetUserType() *string
}

type ListConfigRulesForTargetRequest struct {
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
	// The use scenario of the tag policy. This parameter specifies a filter condition for the query. Valid values:
	//
	// 	- tags: enables tags with specified tag values to be added to resources.
	//
	// 	- rg_inherit: enables resources in a resource group to automatically inherit tags from the resource group.
	//
	// example:
	//
	// tags
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The tag key. This parameter specifies a filter condition for the query.
	//
	// example:
	//
	// CostCenter
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The ID of the object. This parameter specifies a filter condition for the query.
	//
	// example:
	//
	// 134254031178****
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

func (s ListConfigRulesForTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRulesForTargetRequest) GoString() string {
	return s.String()
}

func (s *ListConfigRulesForTargetRequest) GetMaxResult() *int32 {
	return s.MaxResult
}

func (s *ListConfigRulesForTargetRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConfigRulesForTargetRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListConfigRulesForTargetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListConfigRulesForTargetRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListConfigRulesForTargetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListConfigRulesForTargetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListConfigRulesForTargetRequest) GetTagKey() *string {
	return s.TagKey
}

func (s *ListConfigRulesForTargetRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *ListConfigRulesForTargetRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListConfigRulesForTargetRequest) GetUserType() *string {
	return s.UserType
}

func (s *ListConfigRulesForTargetRequest) SetMaxResult(v int32) *ListConfigRulesForTargetRequest {
	s.MaxResult = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetNextToken(v string) *ListConfigRulesForTargetRequest {
	s.NextToken = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetOwnerAccount(v string) *ListConfigRulesForTargetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetOwnerId(v int64) *ListConfigRulesForTargetRequest {
	s.OwnerId = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetPolicyType(v string) *ListConfigRulesForTargetRequest {
	s.PolicyType = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetRegionId(v string) *ListConfigRulesForTargetRequest {
	s.RegionId = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetResourceOwnerAccount(v string) *ListConfigRulesForTargetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetTagKey(v string) *ListConfigRulesForTargetRequest {
	s.TagKey = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetTargetId(v string) *ListConfigRulesForTargetRequest {
	s.TargetId = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetTargetType(v string) *ListConfigRulesForTargetRequest {
	s.TargetType = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetUserType(v string) *ListConfigRulesForTargetRequest {
	s.UserType = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) Validate() error {
	return dara.Validate(s)
}
