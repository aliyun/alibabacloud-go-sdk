// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResubmitConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyResubmitConfigShrinkRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyResubmitConfigShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyResubmitConfigShrinkRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifyResubmitConfigShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyResubmitConfigShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyResubmitConfigShrinkRequest
	GetResourceOwnerId() *int64
	SetRulesShrink(v string) *ModifyResubmitConfigShrinkRequest
	GetRulesShrink() *string
}

type ModifyResubmitConfigShrinkRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6wjk5xxxxxxxxxx
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The job resubmission rules.
	//
	// This parameter is required.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s ModifyResubmitConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyResubmitConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyResubmitConfigShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyResubmitConfigShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyResubmitConfigShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyResubmitConfigShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyResubmitConfigShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyResubmitConfigShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyResubmitConfigShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *ModifyResubmitConfigShrinkRequest) SetDBClusterId(v string) *ModifyResubmitConfigShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyResubmitConfigShrinkRequest) SetOwnerAccount(v string) *ModifyResubmitConfigShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyResubmitConfigShrinkRequest) SetOwnerId(v int64) *ModifyResubmitConfigShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyResubmitConfigShrinkRequest) SetResourceGroupId(v string) *ModifyResubmitConfigShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyResubmitConfigShrinkRequest) SetResourceOwnerAccount(v string) *ModifyResubmitConfigShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyResubmitConfigShrinkRequest) SetResourceOwnerId(v int64) *ModifyResubmitConfigShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyResubmitConfigShrinkRequest) SetRulesShrink(v string) *ModifyResubmitConfigShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *ModifyResubmitConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
