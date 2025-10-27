// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadTasksRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeLoadTasksRecordsRequest
	GetDBClusterId() *string
	SetDBName(v string) *DescribeLoadTasksRecordsRequest
	GetDBName() *string
	SetEndTime(v string) *DescribeLoadTasksRecordsRequest
	GetEndTime() *string
	SetOrder(v string) *DescribeLoadTasksRecordsRequest
	GetOrder() *string
	SetOwnerAccount(v string) *DescribeLoadTasksRecordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLoadTasksRecordsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLoadTasksRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLoadTasksRecordsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeLoadTasksRecordsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLoadTasksRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLoadTasksRecordsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeLoadTasksRecordsRequest
	GetStartTime() *string
	SetState(v string) *DescribeLoadTasksRecordsRequest
	GetState() *string
}

type DescribeLoadTasksRecordsRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters in a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp2590j****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database that is involved in the import or export task.
	//
	// example:
	//
	// adb_demo
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-05-18T06:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The order in which to sort the tasks by field. Specify the field and the sort order in the JSON format. Example: `[{"Field":"CreateTime", "Type":"desc"}]`.
	//
	// >
	//
	// 	- `Field` specifies the field that is used to sort the tasks. Valid values of Field: `State`, `CreateTime`, `DBName`, `ProcessID`, `UpdateTime`, `JobName`, and `ProcessRows`.
	//
	// 	- `Type` specifies the sort order. Valid values of Type: `Desc` and `Asc`. The values are case-insensitive.
	//
	// example:
	//
	// [{"Field":"CreateTime", "Type":"desc"}]
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// > We recommend that you set the query start time to any point in time within 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-05-18T06:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the asynchronous import or export task to be queried. Valid values:
	//
	// 	- **INIT**: The task is being initialized.
	//
	// 	- **RUNNING**: The task is running.
	//
	// 	- **FINISH**: The task is successful.
	//
	// 	- **FAILED**: The task fails.
	//
	// example:
	//
	// FINISH
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeLoadTasksRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadTasksRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeLoadTasksRecordsRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeLoadTasksRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLoadTasksRecordsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeLoadTasksRecordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLoadTasksRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLoadTasksRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLoadTasksRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLoadTasksRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoadTasksRecordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLoadTasksRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLoadTasksRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLoadTasksRecordsRequest) GetState() *string {
	return s.State
}

func (s *DescribeLoadTasksRecordsRequest) SetDBClusterId(v string) *DescribeLoadTasksRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetDBName(v string) *DescribeLoadTasksRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetEndTime(v string) *DescribeLoadTasksRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetOrder(v string) *DescribeLoadTasksRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetOwnerAccount(v string) *DescribeLoadTasksRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetOwnerId(v int64) *DescribeLoadTasksRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetPageNumber(v int32) *DescribeLoadTasksRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetPageSize(v int32) *DescribeLoadTasksRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetRegionId(v string) *DescribeLoadTasksRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetResourceOwnerAccount(v string) *DescribeLoadTasksRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetResourceOwnerId(v int64) *DescribeLoadTasksRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetStartTime(v string) *DescribeLoadTasksRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetState(v string) *DescribeLoadTasksRecordsRequest {
	s.State = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) Validate() error {
	return dara.Validate(s)
}
