// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterParametersRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The ID of the cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query information about all clusters that are deployed in a specified region, such as the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies an immediate or scheduled task to modify parameters and restart the cluster. Valid values:
	//
	// 	- false: scheduled task
	//
	// 	- true: immediate task
	//
	// example:
	//
	// false
	FromTimeService *bool   `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the parameter template.
	//
	// >
	//
	// 	- You can call the [DescribeParameterGroups](https://help.aliyun.com/document_detail/207178.html) operation to query the parameter template ID.
	//
	// 	- You must specify this parameter or the `Parameters` parameter.
	//
	// 	- This parameter is valid only for a PolarDB for MySQL cluster.
	//
	// example:
	//
	// pcpg-**************
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The JSON string that consists of parameters and values. The parameter values are strings, for example, `{"wait_timeout":"86","innodb_old_blocks_time":"10"}`.
	//
	// >
	//
	// 	- You can call the [DescribeDBClusterParameters](https://help.aliyun.com/document_detail/98122.html) operation to query the parameters of the PolarDB cluster.
	//
	// 	- This parameter is required for a PolarDB for Oracle or PolarDB for PostgreSQL cluster.
	//
	// 	- For PolarDB for MySQL clusters, you must specify this parameter or the `ParameterGroupId` parameter.
	//
	// example:
	//
	// {"wait_timeout":"86","innodb_old_blocks_time":"10"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The latest start time to run the task. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// >
	//
	// 	- The value of this parameter must be at least 30 minutes later than the value of the PlannedStartTime parameter.
	//
	// 	- By default, if you specify the `PlannedStartTime` parameter but do not specify the PlannedEndTime parameter, the latest start time of the task is set to a value that is calculated by using the following formula: `Value of the PlannedEndTime parameter + 30 minutes`. For example, if you set the `PlannedStartTime` parameter to `2021-01-14T09:00:00Z` and you do not specify the PlannedEndTime parameter, the latest start time of the task is set to `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2022-04-28T14:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest time to upgrade the specifications within the scheduled time period. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// >
	//
	// 	- The earliest start time of the task can be a point in time within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, you can specify a point in the time range from `2021-01-14T09:00:00Z` to `2021-01-15T09:00:00Z`.
	//
	// 	- If this parameter is empty, the upgrade task is immediately performed.
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
