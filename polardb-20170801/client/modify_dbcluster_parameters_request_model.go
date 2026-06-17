// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClearBinlog(v bool) *ModifyDBClusterParametersRequest
	GetClearBinlog() *bool
	SetDBClusterId(v string) *ModifyDBClusterParametersRequest
	GetDBClusterId() *string
	SetFromTimeService(v bool) *ModifyDBClusterParametersRequest
	GetFromTimeService() *bool
	SetOwnerAccount(v string) *ModifyDBClusterParametersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterParametersRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *ModifyDBClusterParametersRequest
	GetParameterGroupId() *string
	SetParameters(v string) *ModifyDBClusterParametersRequest
	GetParameters() *string
	SetPlannedEndTime(v string) *ModifyDBClusterParametersRequest
	GetPlannedEndTime() *string
	SetPlannedStartTime(v string) *ModifyDBClusterParametersRequest
	GetPlannedStartTime() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterParametersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterParametersRequest
	GetResourceOwnerId() *int64
}

type ModifyDBClusterParametersRequest struct {
	ClearBinlog *bool `json:"ClearBinlog,omitempty" xml:"ClearBinlog,omitempty"`
	// The cluster ID.
	//
	// > Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to view information about all clusters in the destination region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to immediately modify the parameters and restart the cluster. Valid values:
	//
	// - false (default): The modification is scheduled.
	//
	// - true: The modification is performed immediately.
	//
	// example:
	//
	// false
	FromTimeService *bool   `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the parameter template.
	//
	// > - Call the [DescribeParameterGroups](https://help.aliyun.com/document_detail/207178.html) operation to view the ID of the parameter template.
	//
	// >
	//
	// > - You must specify this parameter or the `Parameters` parameter.
	//
	// >
	//
	// > - This parameter is supported only by PolarDB for MySQL.
	//
	// example:
	//
	// pcpg-**************
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// A JSON string that consists of parameters and their values. The parameter values must be strings. For example: `{"wait_timeout":"86","innodb_old_blocks_time":"10"}`.
	//
	// > - Call the [DescribeDBClusterParameters](https://help.aliyun.com/document_detail/98122.html) operation to view the parameters of a PolarDB cluster.
	//
	// >
	//
	// > - This parameter is required if the destination cluster is a PolarDB for PostgreSQL or PolarDB for PostgreSQL (Oracle compatible) cluster.
	//
	// >
	//
	// > - If the destination cluster is a PolarDB for MySQL cluster, you must specify this parameter or the `ParameterGroupId` parameter.
	//
	// example:
	//
	// {"wait_timeout":"86","innodb_old_blocks_time":"10"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The latest time to start the scheduled task. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// > - The latest start time must be at least 30 minutes later than the earliest start time.
	//
	// >
	//
	// > - If you specify `PlannedStartTime` but not this parameter, the latest time to start the task is `PlannedStartTime + 30 minutes` by default. For example, if you set `PlannedStartTime` to `2021-01-14T09:00:00Z` and leave this parameter empty, the task starts no later than `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2022-04-28T14:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest time to start the scheduled task. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// > - The start time can be any point in time within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, you can set the start time to a value in the range of `2021-01-14T09:00:00Z` to `2021-01-15T09:00:00Z`.
	//
	// >
	//
	// > - If you leave this parameter empty, the task is executed immediately.
	//
	// example:
	//
	// 2022-04-28T14:00:00Z
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterParametersRequest) GetClearBinlog() *bool {
	return s.ClearBinlog
}

func (s *ModifyDBClusterParametersRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterParametersRequest) GetFromTimeService() *bool {
	return s.FromTimeService
}

func (s *ModifyDBClusterParametersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterParametersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterParametersRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ModifyDBClusterParametersRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyDBClusterParametersRequest) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *ModifyDBClusterParametersRequest) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *ModifyDBClusterParametersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterParametersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterParametersRequest) SetClearBinlog(v bool) *ModifyDBClusterParametersRequest {
	s.ClearBinlog = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetDBClusterId(v string) *ModifyDBClusterParametersRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetFromTimeService(v bool) *ModifyDBClusterParametersRequest {
	s.FromTimeService = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetOwnerAccount(v string) *ModifyDBClusterParametersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetOwnerId(v int64) *ModifyDBClusterParametersRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetParameterGroupId(v string) *ModifyDBClusterParametersRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetParameters(v string) *ModifyDBClusterParametersRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetPlannedEndTime(v string) *ModifyDBClusterParametersRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetPlannedStartTime(v string) *ModifyDBClusterParametersRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterParametersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetResourceOwnerId(v int64) *ModifyDBClusterParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) Validate() error {
	return dara.Validate(s)
}
