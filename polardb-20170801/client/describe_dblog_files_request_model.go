// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBLogFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBLogFilesRequest
	GetDBClusterId() *string
	SetDBNodeId(v string) *DescribeDBLogFilesRequest
	GetDBNodeId() *string
	SetDescribeSimulateSwitchMode(v string) *DescribeDBLogFilesRequest
	GetDescribeSimulateSwitchMode() *string
	SetEndTime(v string) *DescribeDBLogFilesRequest
	GetEndTime() *string
	SetLogType(v string) *DescribeDBLogFilesRequest
	GetLogType() *string
	SetOwnerAccount(v string) *DescribeDBLogFilesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBLogFilesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBLogFilesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBLogFilesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeDBLogFilesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBLogFilesRequest
	GetResourceOwnerId() *int64
	SetSimulateListId(v string) *DescribeDBLogFilesRequest
	GetSimulateListId() *string
	SetSimulateModeList(v string) *DescribeDBLogFilesRequest
	GetSimulateModeList() *string
	SetSimulateStatusList(v string) *DescribeDBLogFilesRequest
	GetSimulateStatusList() *string
	SetStartTime(v string) *DescribeDBLogFilesRequest
	GetStartTime() *string
}

type DescribeDBLogFilesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the node in the PolarDB cluster.
	//
	// example:
	//
	// pi-*************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The mode to query simulation records. Valid values:
	//
	// - **0**: Queries the simulation records of a logical instance. You can specify the simulation ID.
	//
	// - **1**: Queries the simulation records of a physical instance based on a specified `SimulateListId`.
	//
	// - **2**: Queries the records of the most recent simulation in progress.
	//
	// example:
	//
	// 0
	DescribeSimulateSwitchMode *string `json:"DescribeSimulateSwitchMode,omitempty" xml:"DescribeSimulateSwitchMode,omitempty"`
	// The end of the time range to query. The time must be in UTC and formatted as `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// example:
	//
	// 2023-09-20T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The log type. Valid values:
	//
	// - **HaSwitchLogList**: A list of failover logs.
	//
	// - **HaSwitchLogInfo**: Details of a failover log.
	//
	// - **SimulateSwitchLogs**: Failover simulation logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// HaSwitchLogList
	LogType      *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be an integer that is greater than 0. The default value is **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. Valid values: 5 to 50. The default value is 10.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of a logical instance simulation record.
	//
	// example:
	//
	// "95"
	SimulateListId *string `json:"SimulateListId,omitempty" xml:"SimulateListId,omitempty"`
	// Filters the results by one or more simulation modes. Specify multiple modes as a comma-separated string. Valid values:
	//
	// - **0**: Fault is injected into the primary availability zone.
	//
	// - **1**: Fault is injected into the DB instance.
	//
	// - **2**: Fault is injected into the disaster recovery data center.
	//
	// example:
	//
	// 0,1,2
	SimulateModeList *string `json:"SimulateModeList,omitempty" xml:"SimulateModeList,omitempty"`
	// Filters the results by one or more simulation statuses. Specify multiple statuses as a comma-separated string. Valid values:
	//
	// - **0**: Waiting for scheduling
	//
	// - **1**: Succeeded
	//
	// - **2**: Running
	//
	// - **3**: Failed
	//
	// - **4**: Interrupted
	//
	// - **5**: Waiting for switchback
	//
	// - **6**: Canceled
	//
	// example:
	//
	// "0,1,2,3,4,5"
	SimulateStatusList *string `json:"SimulateStatusList,omitempty" xml:"SimulateStatusList,omitempty"`
	// The start of the time range to query. The time must be in UTC and formatted as `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// example:
	//
	// 2023-08-20T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBLogFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBLogFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBLogFilesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBLogFilesRequest) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDBLogFilesRequest) GetDescribeSimulateSwitchMode() *string {
	return s.DescribeSimulateSwitchMode
}

func (s *DescribeDBLogFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBLogFilesRequest) GetLogType() *string {
	return s.LogType
}

func (s *DescribeDBLogFilesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBLogFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBLogFilesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBLogFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBLogFilesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBLogFilesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBLogFilesRequest) GetSimulateListId() *string {
	return s.SimulateListId
}

func (s *DescribeDBLogFilesRequest) GetSimulateModeList() *string {
	return s.SimulateModeList
}

func (s *DescribeDBLogFilesRequest) GetSimulateStatusList() *string {
	return s.SimulateStatusList
}

func (s *DescribeDBLogFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBLogFilesRequest) SetDBClusterId(v string) *DescribeDBLogFilesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBLogFilesRequest) SetDBNodeId(v string) *DescribeDBLogFilesRequest {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBLogFilesRequest) SetDescribeSimulateSwitchMode(v string) *DescribeDBLogFilesRequest {
	s.DescribeSimulateSwitchMode = &v
	return s
}

func (s *DescribeDBLogFilesRequest) SetEndTime(v string) *DescribeDBLogFilesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBLogFilesRequest) SetLogType(v string) *DescribeDBLogFilesRequest {
	s.LogType = &v
	return s
}

func (s *DescribeDBLogFilesRequest) SetOwnerAccount(v string) *DescribeDBLogFilesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBLogFilesRequest) SetOwnerId(v int64) *DescribeDBLogFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBLogFilesRequest) SetPageNumber(v int32) *DescribeDBLogFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBLogFilesRequest) SetPageSize(v int32) *DescribeDBLogFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBLogFilesRequest) SetResourceOwnerAccount(v string) *DescribeDBLogFilesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBLogFilesRequest) SetResourceOwnerId(v int64) *DescribeDBLogFilesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBLogFilesRequest) SetSimulateListId(v string) *DescribeDBLogFilesRequest {
	s.SimulateListId = &v
	return s
}

func (s *DescribeDBLogFilesRequest) SetSimulateModeList(v string) *DescribeDBLogFilesRequest {
	s.SimulateModeList = &v
	return s
}

func (s *DescribeDBLogFilesRequest) SetSimulateStatusList(v string) *DescribeDBLogFilesRequest {
	s.SimulateStatusList = &v
	return s
}

func (s *DescribeDBLogFilesRequest) SetStartTime(v string) *DescribeDBLogFilesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBLogFilesRequest) Validate() error {
	return dara.Validate(s)
}
