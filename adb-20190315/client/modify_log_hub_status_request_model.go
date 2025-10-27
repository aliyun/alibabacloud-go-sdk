// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogHubStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyLogHubStatusRequest
	GetDBClusterId() *string
	SetDeliverName(v string) *ModifyLogHubStatusRequest
	GetDeliverName() *string
	SetLogStoreName(v string) *ModifyLogHubStatusRequest
	GetLogStoreName() *string
	SetOwnerAccount(v string) *ModifyLogHubStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyLogHubStatusRequest
	GetOwnerId() *int64
	SetProjectName(v string) *ModifyLogHubStatusRequest
	GetProjectName() *string
	SetResourceOwnerAccount(v string) *ModifyLogHubStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyLogHubStatusRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ModifyLogHubStatusRequest
	GetStatus() *string
}

type ModifyLogHubStatusRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6j8370er80m6wf3
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the log shipping job.
	//
	// This parameter is required.
	//
	// example:
	//
	// loghub-mnp-report
	DeliverName *string `json:"DeliverName,omitempty" xml:"DeliverName,omitempty"`
	// The name of the Logstore.
	//
	// This parameter is required.
	//
	// example:
	//
	// pay-notify-wx
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the Simple Log Service project.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-adb
	ProjectName          *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the log shipping job.
	//
	// Valid values:
	//
	// 	- Delete
	//
	// 	- Pause
	//
	// 	- Restart
	//
	// This parameter is required.
	//
	// example:
	//
	// Pause
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyLogHubStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogHubStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogHubStatusRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyLogHubStatusRequest) GetDeliverName() *string {
	return s.DeliverName
}

func (s *ModifyLogHubStatusRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *ModifyLogHubStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyLogHubStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLogHubStatusRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ModifyLogHubStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyLogHubStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyLogHubStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyLogHubStatusRequest) SetDBClusterId(v string) *ModifyLogHubStatusRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyLogHubStatusRequest) SetDeliverName(v string) *ModifyLogHubStatusRequest {
	s.DeliverName = &v
	return s
}

func (s *ModifyLogHubStatusRequest) SetLogStoreName(v string) *ModifyLogHubStatusRequest {
	s.LogStoreName = &v
	return s
}

func (s *ModifyLogHubStatusRequest) SetOwnerAccount(v string) *ModifyLogHubStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLogHubStatusRequest) SetOwnerId(v int64) *ModifyLogHubStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLogHubStatusRequest) SetProjectName(v string) *ModifyLogHubStatusRequest {
	s.ProjectName = &v
	return s
}

func (s *ModifyLogHubStatusRequest) SetResourceOwnerAccount(v string) *ModifyLogHubStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLogHubStatusRequest) SetResourceOwnerId(v int64) *ModifyLogHubStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLogHubStatusRequest) SetStatus(v string) *ModifyLogHubStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyLogHubStatusRequest) Validate() error {
	return dara.Validate(s)
}
