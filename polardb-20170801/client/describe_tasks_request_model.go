// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeTasksRequest
	GetDBClusterId() *string
	SetDBNodeId(v string) *DescribeTasksRequest
	GetDBNodeId() *string
	SetEndTime(v string) *DescribeTasksRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTasksRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTasksRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeTasksRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeTasksRequest
	GetStatus() *string
}

type DescribeTasksRequest struct {
	// The cluster ID.
	//
	// >  You must specify `DBNodeId` or `DBClusterId`. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of the clusters that belong to your Alibaba Cloud account, such as cluster IDs.
	//
	// example:
	//
	// pc-***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The node ID.
	//
	// >  You must specify `DBNodeId` or `DBClusterId`. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of the clusters that belong to your Alibaba Cloud account, such as node IDs.
	//
	// example:
	//
	// pi-***************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mmZ` format. The time must be in UTC. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-12-02T03:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30**, **50**, and **100**.
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-11-30T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the tasks that you want to query. Valid values:
	//
	// 	- **Waiting**: The task is pending.
	//
	// 	- **Running**: The task is running.
	//
	// 	- **Finished**: The task is completed.
	//
	// 	- **Closed**: The task is closed.
	//
	// 	- **Pause**: The task is paused.
	//
	// 	- **Stop**: The task is interrupted.
	//
	// >
	//
	// 	- If you do not specify this parameter, the operation returns the details of only the tasks that are in the **Waiting*	- or **Running*	- state for the cluster or node.
	//
	// 	- You can enter multiple task states. Separate multiple task states with commas (,).
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeTasksRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTasksRequest) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeTasksRequest) SetDBClusterId(v string) *DescribeTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTasksRequest) SetDBNodeId(v string) *DescribeTasksRequest {
	s.DBNodeId = &v
	return s
}

func (s *DescribeTasksRequest) SetEndTime(v string) *DescribeTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerAccount(v string) *DescribeTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerId(v int64) *DescribeTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetPageNumber(v int32) *DescribeTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksRequest) SetPageSize(v int32) *DescribeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceOwnerAccount(v string) *DescribeTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceOwnerId(v int64) *DescribeTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetStartTime(v string) *DescribeTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTasksRequest) SetStatus(v string) *DescribeTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeTasksRequest) Validate() error {
	return dara.Validate(s)
}
