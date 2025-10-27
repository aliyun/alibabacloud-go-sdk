// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySQAConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifySQAConfigRequest
	GetDBClusterId() *string
	SetGroupName(v string) *ModifySQAConfigRequest
	GetGroupName() *string
	SetOwnerAccount(v string) *ModifySQAConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySQAConfigRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifySQAConfigRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifySQAConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySQAConfigRequest
	GetResourceOwnerId() *int64
	SetSQAStatus(v string) *ModifySQAConfigRequest
	GetSQAStatus() *string
}

type ModifySQAConfigRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-8vb369k7zxdt10tz0
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// >  You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/459446.html) operation to query the resource group name of a cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
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
	// Specifies whether to enable short query acceleration (SQA).
	//
	// This parameter is required.
	//
	// example:
	//
	// off
	SQAStatus *string `json:"SQAStatus,omitempty" xml:"SQAStatus,omitempty"`
}

func (s ModifySQAConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySQAConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifySQAConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifySQAConfigRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifySQAConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySQAConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySQAConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifySQAConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySQAConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySQAConfigRequest) GetSQAStatus() *string {
	return s.SQAStatus
}

func (s *ModifySQAConfigRequest) SetDBClusterId(v string) *ModifySQAConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifySQAConfigRequest) SetGroupName(v string) *ModifySQAConfigRequest {
	s.GroupName = &v
	return s
}

func (s *ModifySQAConfigRequest) SetOwnerAccount(v string) *ModifySQAConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySQAConfigRequest) SetOwnerId(v int64) *ModifySQAConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySQAConfigRequest) SetResourceGroupId(v string) *ModifySQAConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifySQAConfigRequest) SetResourceOwnerAccount(v string) *ModifySQAConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySQAConfigRequest) SetResourceOwnerId(v int64) *ModifySQAConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySQAConfigRequest) SetSQAStatus(v string) *ModifySQAConfigRequest {
	s.SQAStatus = &v
	return s
}

func (s *ModifySQAConfigRequest) Validate() error {
	return dara.Validate(s)
}
