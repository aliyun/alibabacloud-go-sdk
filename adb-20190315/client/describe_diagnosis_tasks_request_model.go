// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDiagnosisTasksRequest
	GetDBClusterId() *string
	SetHost(v string) *DescribeDiagnosisTasksRequest
	GetHost() *string
	SetOrder(v string) *DescribeDiagnosisTasksRequest
	GetOrder() *string
	SetPageNumber(v int32) *DescribeDiagnosisTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDiagnosisTasksRequest
	GetPageSize() *int32
	SetProcessId(v string) *DescribeDiagnosisTasksRequest
	GetProcessId() *string
	SetRegionId(v string) *DescribeDiagnosisTasksRequest
	GetRegionId() *string
	SetStageId(v string) *DescribeDiagnosisTasksRequest
	GetStageId() *string
	SetState(v string) *DescribeDiagnosisTasksRequest
	GetState() *string
}

type DescribeDiagnosisTasksRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The IP address from which the query was initiated.
	//
	// example:
	//
	// 192.168.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The order in which to sort the tasks by field. Specify the value in the JSON format. Example: `[{"Field":"CreateTime", "Type":"desc"}]`.
	//
	// >
	//
	// 	- `Field` specifies the field that is used to sort the tasks. Valid values of Field: `State`, `CreateTime`, `DBName`, `ProcessID`, `UpdateTime`, `JobName`, and `ProcessRows`.
	//
	// 	- `Type` specifies the sort order. Valid values of Type: `Desc` and `Asc`. The values are case-insensitive.
	//
	// example:
	//
	// [{"Field":"StartTime", "Type": "desc" }]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- 30 (default)
	//
	// 	- 50
	//
	// 	- 100
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query ID.
	//
	// > You can call the [DescribeProcessList](https://help.aliyun.com/document_detail/190092.html) operation to query the IDs of queries that are being executed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 202011191048151921681492420315100****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of a stage in the query that is specified by the `ProcessId` parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// Stage[26]
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The state of the asynchronous import or export task to be queried. Valid values:
	//
	// 	- **RUNNING**
	//
	// 	- **FINISHED**
	//
	// 	- **FAILED**
	//
	// example:
	//
	// RUNNING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeDiagnosisTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisTasksRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDiagnosisTasksRequest) GetHost() *string {
	return s.Host
}

func (s *DescribeDiagnosisTasksRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeDiagnosisTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDiagnosisTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDiagnosisTasksRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *DescribeDiagnosisTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiagnosisTasksRequest) GetStageId() *string {
	return s.StageId
}

func (s *DescribeDiagnosisTasksRequest) GetState() *string {
	return s.State
}

func (s *DescribeDiagnosisTasksRequest) SetDBClusterId(v string) *DescribeDiagnosisTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetHost(v string) *DescribeDiagnosisTasksRequest {
	s.Host = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetOrder(v string) *DescribeDiagnosisTasksRequest {
	s.Order = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetPageNumber(v int32) *DescribeDiagnosisTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetPageSize(v int32) *DescribeDiagnosisTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetProcessId(v string) *DescribeDiagnosisTasksRequest {
	s.ProcessId = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetRegionId(v string) *DescribeDiagnosisTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetStageId(v string) *DescribeDiagnosisTasksRequest {
	s.StageId = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetState(v string) *DescribeDiagnosisTasksRequest {
	s.State = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) Validate() error {
	return dara.Validate(s)
}
