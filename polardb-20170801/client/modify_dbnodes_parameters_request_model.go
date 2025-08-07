// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodesParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBNodesParametersRequest
	GetDBClusterId() *string
	SetDBNodeIds(v string) *ModifyDBNodesParametersRequest
	GetDBNodeIds() *string
	SetFromTimeService(v bool) *ModifyDBNodesParametersRequest
	GetFromTimeService() *bool
	SetOwnerAccount(v string) *ModifyDBNodesParametersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBNodesParametersRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *ModifyDBNodesParametersRequest
	GetParameterGroupId() *string
	SetParameters(v string) *ModifyDBNodesParametersRequest
	GetParameters() *string
	SetPlannedEndTime(v string) *ModifyDBNodesParametersRequest
	GetPlannedEndTime() *string
	SetPlannedStartTime(v string) *ModifyDBNodesParametersRequest
	GetPlannedStartTime() *string
	SetResourceOwnerAccount(v string) *ModifyDBNodesParametersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBNodesParametersRequest
	GetResourceOwnerId() *int64
}

type ModifyDBNodesParametersRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the node. You can specify multiple node IDs. Separate multiple node IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// pi-****************， pi-****************
	DBNodeIds *string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty"`
	// Specifies whether to immediately run the task to modify parameters and restart the cluster. Valid values: false: runs the task on schedule. true: runs the task immediately. Default value: false.
	//
	// example:
	//
	// false
	FromTimeService *bool   `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the parameter template that is used for the cluster.
	//
	// example:
	//
	// pcpg-**************
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The JSON string that specifies the parameter and its value.
	//
	// example:
	//
	// {"wait_timeout":"86","innodb_old_blocks_time":"10"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The latest start time to run the task. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-04-28T14:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest start time to run the task to upgrade the kernel version of the cluster. The task runs within a specified period of time. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-04-28T14:00:00Z
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBNodesParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodesParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBNodesParametersRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBNodesParametersRequest) GetDBNodeIds() *string {
	return s.DBNodeIds
}

func (s *ModifyDBNodesParametersRequest) GetFromTimeService() *bool {
	return s.FromTimeService
}

func (s *ModifyDBNodesParametersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBNodesParametersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBNodesParametersRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ModifyDBNodesParametersRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyDBNodesParametersRequest) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *ModifyDBNodesParametersRequest) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *ModifyDBNodesParametersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBNodesParametersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBNodesParametersRequest) SetDBClusterId(v string) *ModifyDBNodesParametersRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodesParametersRequest) SetDBNodeIds(v string) *ModifyDBNodesParametersRequest {
	s.DBNodeIds = &v
	return s
}

func (s *ModifyDBNodesParametersRequest) SetFromTimeService(v bool) *ModifyDBNodesParametersRequest {
	s.FromTimeService = &v
	return s
}

func (s *ModifyDBNodesParametersRequest) SetOwnerAccount(v string) *ModifyDBNodesParametersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBNodesParametersRequest) SetOwnerId(v int64) *ModifyDBNodesParametersRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBNodesParametersRequest) SetParameterGroupId(v string) *ModifyDBNodesParametersRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *ModifyDBNodesParametersRequest) SetParameters(v string) *ModifyDBNodesParametersRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyDBNodesParametersRequest) SetPlannedEndTime(v string) *ModifyDBNodesParametersRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *ModifyDBNodesParametersRequest) SetPlannedStartTime(v string) *ModifyDBNodesParametersRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *ModifyDBNodesParametersRequest) SetResourceOwnerAccount(v string) *ModifyDBNodesParametersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBNodesParametersRequest) SetResourceOwnerId(v int64) *ModifyDBNodesParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBNodesParametersRequest) Validate() error {
	return dara.Validate(s)
}
