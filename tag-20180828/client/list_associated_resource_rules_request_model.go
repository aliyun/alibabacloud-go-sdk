// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssociatedResourceRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResult(v int32) *ListAssociatedResourceRulesRequest
	GetMaxResult() *int32
	SetNextToken(v string) *ListAssociatedResourceRulesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListAssociatedResourceRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListAssociatedResourceRulesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListAssociatedResourceRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListAssociatedResourceRulesRequest
	GetResourceOwnerAccount() *string
	SetSettingName(v []*string) *ListAssociatedResourceRulesRequest
	GetSettingName() []*string
	SetStatus(v string) *ListAssociatedResourceRulesRequest
	GetStatus() *string
}

type ListAssociatedResourceRulesRequest struct {
	// The number of entries to return on each page.
	//
	// Default Value: 50. Maximum Value: 100.
	//
	// example:
	//
	// 50
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The token returned from a previous call to retrieve the next page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The names of the associated resource rules.
	SettingName []*string `json:"SettingName,omitempty" xml:"SettingName,omitempty" type:"Repeated"`
	// The status of the associated resource rules to query. Valid values:
	//
	// - Enable: The rule is enabled.
	//
	// - Disable: The rule is disabled.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAssociatedResourceRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAssociatedResourceRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAssociatedResourceRulesRequest) GetMaxResult() *int32 {
	return s.MaxResult
}

func (s *ListAssociatedResourceRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAssociatedResourceRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAssociatedResourceRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAssociatedResourceRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAssociatedResourceRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAssociatedResourceRulesRequest) GetSettingName() []*string {
	return s.SettingName
}

func (s *ListAssociatedResourceRulesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAssociatedResourceRulesRequest) SetMaxResult(v int32) *ListAssociatedResourceRulesRequest {
	s.MaxResult = &v
	return s
}

func (s *ListAssociatedResourceRulesRequest) SetNextToken(v string) *ListAssociatedResourceRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAssociatedResourceRulesRequest) SetOwnerAccount(v string) *ListAssociatedResourceRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAssociatedResourceRulesRequest) SetOwnerId(v int64) *ListAssociatedResourceRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAssociatedResourceRulesRequest) SetRegionId(v string) *ListAssociatedResourceRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListAssociatedResourceRulesRequest) SetResourceOwnerAccount(v string) *ListAssociatedResourceRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAssociatedResourceRulesRequest) SetSettingName(v []*string) *ListAssociatedResourceRulesRequest {
	s.SettingName = v
	return s
}

func (s *ListAssociatedResourceRulesRequest) SetStatus(v string) *ListAssociatedResourceRulesRequest {
	s.Status = &v
	return s
}

func (s *ListAssociatedResourceRulesRequest) Validate() error {
	return dara.Validate(s)
}
