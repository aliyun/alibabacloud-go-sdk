// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDataSourcesResponseBody
	GetCode() *string
	SetDataSources(v []*DescribeDataSourcesResponseBodyDataSources) *DescribeDataSourcesResponseBody
	GetDataSources() []*DescribeDataSourcesResponseBodyDataSources
	SetMessage(v string) *DescribeDataSourcesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeDataSourcesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDataSourcesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDataSourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDataSourcesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeDataSourcesResponseBody
	GetTotalCount() *int64
}

type DescribeDataSourcesResponseBody struct {
	// example:
	//
	// 200
	Code        *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	DataSources []*DescribeDataSourcesResponseBodyDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 36A5CD24-****-****-****-5F30C3F1753F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataSourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDataSourcesResponseBody) GetDataSources() []*DescribeDataSourcesResponseBodyDataSources {
	return s.DataSources
}

func (s *DescribeDataSourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDataSourcesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataSourcesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataSourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDataSourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDataSourcesResponseBody) SetCode(v string) *DescribeDataSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDataSourcesResponseBody) SetDataSources(v []*DescribeDataSourcesResponseBodyDataSources) *DescribeDataSourcesResponseBody {
	s.DataSources = v
	return s
}

func (s *DescribeDataSourcesResponseBody) SetMessage(v string) *DescribeDataSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDataSourcesResponseBody) SetPageNumber(v int32) *DescribeDataSourcesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataSourcesResponseBody) SetPageSize(v int32) *DescribeDataSourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataSourcesResponseBody) SetRequestId(v string) *DescribeDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataSourcesResponseBody) SetSuccess(v bool) *DescribeDataSourcesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDataSourcesResponseBody) SetTotalCount(v int64) *DescribeDataSourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataSourcesResponseBody) Validate() error {
	if s.DataSources != nil {
		for _, item := range s.DataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataSourcesResponseBodyDataSources struct {
	// example:
	//
	// [
	//
	//       "/home/alice/log"
	//
	// ]
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// {
	//
	//       "dataServerAddresses": [
	//
	//             {
	//
	//                   "host": "123.123.123.123",
	//
	//                   "port": "8080"
	//
	//             }
	//
	//       ],
	//
	//       "sharePath": "/share",
	//
	//       "mountOptions": "",
	//
	//       "fileSystemType": "nfs"
	//
	// }
	ConnectionInfo *string `json:"ConnectionInfo,omitempty" xml:"ConnectionInfo,omitempty"`
	// example:
	//
	// 1770257653
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// ds-000******2nqeo
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// COMMON_NAS
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// example:
	//
	// COMMON_NAS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// [
	//
	//       "/home/alice/log"
	//
	// ]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// example:
	//
	// [
	//
	//       "/home/alice",
	//
	//       "/home/bob"
	//
	// ]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// example:
	//
	// true
	IndexAvailable *bool `json:"IndexAvailable,omitempty" xml:"IndexAvailable,omitempty"`
	// example:
	//
	// META
	IndexLevel *string `json:"IndexLevel,omitempty" xml:"IndexLevel,omitempty"`
	// example:
	//
	// 1745454604
	IndexUpdateTime *bool `json:"IndexUpdateTime,omitempty" xml:"IndexUpdateTime,omitempty"`
	// example:
	//
	// false
	Indexing *bool `json:"Indexing,omitempty" xml:"Indexing,omitempty"`
	// example:
	//
	// {}
	Options *string   `json:"Options,omitempty" xml:"Options,omitempty"`
	Paths   []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// example:
	//
	// plan-123***7890
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// example:
	//
	// I|1729493847|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// example:
	//
	// 6:21:10240
	SpeedLimit *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	// example:
	//
	// 1745454604
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s DescribeDataSourcesResponseBodyDataSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourcesResponseBodyDataSources) GoString() string {
	return s.String()
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetConnectionInfo() *string {
	return s.ConnectionInfo
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetExclude() *string {
	return s.Exclude
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetInclude() *string {
	return s.Include
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetIndexAvailable() *bool {
	return s.IndexAvailable
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetIndexLevel() *string {
	return s.IndexLevel
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetIndexUpdateTime() *bool {
	return s.IndexUpdateTime
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetIndexing() *bool {
	return s.Indexing
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetOptions() *string {
	return s.Options
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetPaths() []*string {
	return s.Paths
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetPlanId() *string {
	return s.PlanId
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetSchedule() *string {
	return s.Schedule
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetSpeedLimit() *string {
	return s.SpeedLimit
}

func (s *DescribeDataSourcesResponseBodyDataSources) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetClusterId(v string) *DescribeDataSourcesResponseBodyDataSources {
	s.ClusterId = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetConnectionInfo(v string) *DescribeDataSourcesResponseBodyDataSources {
	s.ConnectionInfo = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetCreatedTime(v int64) *DescribeDataSourcesResponseBodyDataSources {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetDataSourceId(v string) *DescribeDataSourcesResponseBodyDataSources {
	s.DataSourceId = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetDataSourceName(v string) *DescribeDataSourcesResponseBodyDataSources {
	s.DataSourceName = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetDataSourceType(v string) *DescribeDataSourcesResponseBodyDataSources {
	s.DataSourceType = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetExclude(v string) *DescribeDataSourcesResponseBodyDataSources {
	s.Exclude = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetInclude(v string) *DescribeDataSourcesResponseBodyDataSources {
	s.Include = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetIndexAvailable(v bool) *DescribeDataSourcesResponseBodyDataSources {
	s.IndexAvailable = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetIndexLevel(v string) *DescribeDataSourcesResponseBodyDataSources {
	s.IndexLevel = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetIndexUpdateTime(v bool) *DescribeDataSourcesResponseBodyDataSources {
	s.IndexUpdateTime = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetIndexing(v bool) *DescribeDataSourcesResponseBodyDataSources {
	s.Indexing = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetOptions(v string) *DescribeDataSourcesResponseBodyDataSources {
	s.Options = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetPaths(v []*string) *DescribeDataSourcesResponseBodyDataSources {
	s.Paths = v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetPlanId(v string) *DescribeDataSourcesResponseBodyDataSources {
	s.PlanId = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetSchedule(v string) *DescribeDataSourcesResponseBodyDataSources {
	s.Schedule = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetSpeedLimit(v string) *DescribeDataSourcesResponseBodyDataSources {
	s.SpeedLimit = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) SetUpdatedTime(v int64) *DescribeDataSourcesResponseBodyDataSources {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeDataSourcesResponseBodyDataSources) Validate() error {
	return dara.Validate(s)
}
