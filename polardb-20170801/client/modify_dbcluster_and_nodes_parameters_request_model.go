// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAndNodesParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClearBinlog(v bool) *ModifyDBClusterAndNodesParametersRequest
	GetClearBinlog() *bool
	SetDBClusterId(v string) *ModifyDBClusterAndNodesParametersRequest
	GetDBClusterId() *string
	SetDBNodeIds(v string) *ModifyDBClusterAndNodesParametersRequest
	GetDBNodeIds() *string
	SetFromTimeService(v bool) *ModifyDBClusterAndNodesParametersRequest
	GetFromTimeService() *bool
	SetOwnerAccount(v string) *ModifyDBClusterAndNodesParametersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterAndNodesParametersRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *ModifyDBClusterAndNodesParametersRequest
	GetParameterGroupId() *string
	SetParameters(v string) *ModifyDBClusterAndNodesParametersRequest
	GetParameters() *string
	SetPlannedEndTime(v string) *ModifyDBClusterAndNodesParametersRequest
	GetPlannedEndTime() *string
	SetPlannedStartTime(v string) *ModifyDBClusterAndNodesParametersRequest
	GetPlannedStartTime() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterAndNodesParametersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterAndNodesParametersRequest
	GetResourceOwnerId() *int64
	SetStandbyClusterIdListNeedToSync(v string) *ModifyDBClusterAndNodesParametersRequest
	GetStandbyClusterIdListNeedToSync() *string
}

type ModifyDBClusterAndNodesParametersRequest struct {
	// Specifies whether to clear binlog. This parameter is valid only when binlog is disabled.
	ClearBinlog *bool `json:"ClearBinlog,omitempty" xml:"ClearBinlog,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**********
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The IDs of the nodes. By setting this parameter, you can modify the parameters of the cluster and specified nodes. Separate multiple node IDs with commas (,).
	//
	// > If this parameter is not specified, only the cluster parameters are modified.
	//
	// example:
	//
	// pi-**********,pi-**********
	DBNodeIds *string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty"`
	// Specifies whether to immediately or scheduledly modify parameters and restart the cluster. Valid values:
	//
	// - **false*	- (default): scheduled execution
	//
	// - **true**: immediate execution
	//
	// example:
	//
	// false
	FromTimeService *bool   `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the parameter template.
	//
	// example:
	//
	// pcpg-**********
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The JSON string that consists of parameters and their values.
	//
	// example:
	//
	// {"wait_timeout":"86","innodb_old_blocks_time":"10"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The latest time to start the scheduled task. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// > - The latest start time must be 30 minutes or more later than the earliest start time.
	//
	// >
	//
	// > - If you specify `PlannedStartTime` but not this parameter, the latest start time of the task is `the earliest start time + 30 minutes` by default. For example, if you set `PlannedStartTime` to `2021-01-14T09:00:00Z` and leave this parameter empty, the task will start no later than `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest time to start the scheduled task. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// > - The start time can be any time within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, you can specify a time that ranges from `2021-01-14T09:00:00Z` to `2021-01-15T09:00:00Z`.
	//
	// >
	//
	// > - If you leave this parameter empty, the task is immediately executed.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of GDN standby clusters to which you want to synchronize the parameter settings.
	//
	// example:
	//
	// gdn-**********,gdn-**********
	StandbyClusterIdListNeedToSync *string `json:"StandbyClusterIdListNeedToSync,omitempty" xml:"StandbyClusterIdListNeedToSync,omitempty"`
}

func (s ModifyDBClusterAndNodesParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAndNodesParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAndNodesParametersRequest) GetClearBinlog() *bool {
	return s.ClearBinlog
}

func (s *ModifyDBClusterAndNodesParametersRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterAndNodesParametersRequest) GetDBNodeIds() *string {
	return s.DBNodeIds
}

func (s *ModifyDBClusterAndNodesParametersRequest) GetFromTimeService() *bool {
	return s.FromTimeService
}

func (s *ModifyDBClusterAndNodesParametersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterAndNodesParametersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterAndNodesParametersRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ModifyDBClusterAndNodesParametersRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyDBClusterAndNodesParametersRequest) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *ModifyDBClusterAndNodesParametersRequest) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *ModifyDBClusterAndNodesParametersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterAndNodesParametersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterAndNodesParametersRequest) GetStandbyClusterIdListNeedToSync() *string {
	return s.StandbyClusterIdListNeedToSync
}

func (s *ModifyDBClusterAndNodesParametersRequest) SetClearBinlog(v bool) *ModifyDBClusterAndNodesParametersRequest {
	s.ClearBinlog = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersRequest) SetDBClusterId(v string) *ModifyDBClusterAndNodesParametersRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersRequest) SetDBNodeIds(v string) *ModifyDBClusterAndNodesParametersRequest {
	s.DBNodeIds = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersRequest) SetFromTimeService(v bool) *ModifyDBClusterAndNodesParametersRequest {
	s.FromTimeService = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersRequest) SetOwnerAccount(v string) *ModifyDBClusterAndNodesParametersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersRequest) SetOwnerId(v int64) *ModifyDBClusterAndNodesParametersRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersRequest) SetParameterGroupId(v string) *ModifyDBClusterAndNodesParametersRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersRequest) SetParameters(v string) *ModifyDBClusterAndNodesParametersRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersRequest) SetPlannedEndTime(v string) *ModifyDBClusterAndNodesParametersRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersRequest) SetPlannedStartTime(v string) *ModifyDBClusterAndNodesParametersRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterAndNodesParametersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersRequest) SetResourceOwnerId(v int64) *ModifyDBClusterAndNodesParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersRequest) SetStandbyClusterIdListNeedToSync(v string) *ModifyDBClusterAndNodesParametersRequest {
	s.StandbyClusterIdListNeedToSync = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersRequest) Validate() error {
	return dara.Validate(s)
}
