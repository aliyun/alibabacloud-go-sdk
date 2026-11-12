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
	// Specifies whether to clear binlogs. This parameter takes effect only when binlogs are disabled.
	//
	// example:
	//
	// false
	ClearBinlog *bool `json:"ClearBinlog,omitempty" xml:"ClearBinlog,omitempty"`
	// The cluster ID. The DBClusterStatus of the target cluster must be Running. You can call the DescribeDBClusters operation to check the current status. Usage notes: The DBClusterStatus of the cluster must be Running. You can call the DescribeDBClusters operation to verify the status.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**********
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The node IDs. You can set this parameter to modify the parameters of the cluster and the specified nodes. Separate multiple node IDs with commas (,).
	//
	// > If this parameter is left empty, only the cluster parameters are modified.
	//
	// example:
	//
	// pi-**********,pi-**********
	DBNodeIds *string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty"`
	// Specifies whether to execute the parameter modification and cluster restart immediately or at a scheduled time. Valid values:
	//
	// - **false**: The task is executed at a scheduled time. This is the default value.
	//
	// - **true**: The task is executed immediately.
	//
	// example:
	//
	// false
	FromTimeService *bool   `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The parameter template ID.
	//
	// example:
	//
	// pcpg-**********
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The JSON string of parameters and their values.
	//
	// example:
	//
	// {"wait_timeout":"86","innodb_old_blocks_time":"10"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The latest time to start the scheduled task. The time must be in the `YYYY-MM-DDThh:mm:ssZ` format and in UTC.
	//
	// > 	- The latest time must be at least 30 minutes later than the start time.
	//
	// > 	- If `PlannedStartTime` is specified but this parameter is not, the latest start time defaults to 30 minutes after the start time. For example, if `PlannedStartTime` is set to `2021-01-14T09:00:00Z` and this parameter is left empty, the task will start no later than `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest time to start the scheduled task. The time must be in the `YYYY-MM-DDThh:mm:ssZ` format and in UTC.
	//
	// > 	- The start time must be within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, the allowed start time ranges from `2021-01-14T09:00:00Z` to `2021-01-15T09:00:00Z`.
	//
	// > 	- If this parameter is left empty, the task is executed immediately by default.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of secondary clusters in the Global Database Network (GDN) that require parameter settings synchronization.
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
