// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTablePartitionDiagnoseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeTablePartitionDiagnoseResponseBody
	GetAccessDeniedDetail() *string
	SetDBClusterId(v string) *DescribeTablePartitionDiagnoseResponseBody
	GetDBClusterId() *string
	SetDetectionItems(v []*DescribeTablePartitionDiagnoseResponseBodyDetectionItems) *DescribeTablePartitionDiagnoseResponseBody
	GetDetectionItems() []*DescribeTablePartitionDiagnoseResponseBodyDetectionItems
	SetItems(v []*DescribeTablePartitionDiagnoseResponseBodyItems) *DescribeTablePartitionDiagnoseResponseBody
	GetItems() []*DescribeTablePartitionDiagnoseResponseBodyItems
	SetPageNumber(v int32) *DescribeTablePartitionDiagnoseResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTablePartitionDiagnoseResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTablePartitionDiagnoseResponseBody
	GetRequestId() *string
	SetSuggestMaxRecordsPerPartition(v int64) *DescribeTablePartitionDiagnoseResponseBody
	GetSuggestMaxRecordsPerPartition() *int64
	SetSuggestMinRecordsPerPartition(v int64) *DescribeTablePartitionDiagnoseResponseBody
	GetSuggestMinRecordsPerPartition() *int64
	SetTotalCount(v int32) *DescribeTablePartitionDiagnoseResponseBody
	GetTotalCount() *int32
}

type DescribeTablePartitionDiagnoseResponseBody struct {
	// The information about the request denial.
	//
	// example:
	//
	// {
	//
	//     "PolicyType": "AccountLevelIdentityBasedPolicy",
	//
	//     "AuthPrincipalOwnerId": "1*****************7",
	//
	//     "EncodedDiagnosticMessage": "AQIBIAAAAOPdwKY2QLOvgMEc7SkkoJfj1kvZwsaRqNYMh10Tv0wTe0fCzaCdrvgazfNb0EnJKETgXyhR+3BIQjx9WAqZryejBsp1Bl4qI5En/D9dEhcXAtKCxCmE2kZCiEzpy8BoEUt+bs0DmlaGWO5xkEpttypLIB4rUhDvZd+zwPg4EXk4KSSWSWsurxtqDkKEMshKlQFBTKvJcKwyhk62IeYly4hQ+5IpXjkh1GQXuDRCQ==",
	//
	//     "AuthPrincipalType": "SubUser",
	//
	//     "AuthPrincipalDisplayName": "2***************9",
	//
	//     "NoPermissionType": "ImplicitDeny",
	//
	//     "AuthAction": "adb:DescribeExcessivePrimaryKeys"
	//
	// }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-bp171g24yvbxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried detection items and detection results.
	DetectionItems []*DescribeTablePartitionDiagnoseResponseBodyDetectionItems `json:"DetectionItems,omitempty" xml:"DetectionItems,omitempty" type:"Repeated"`
	// The queried partition diagnostic information.
	Items []*DescribeTablePartitionDiagnoseResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The recommended maximum number of rows in each partition.
	//
	// example:
	//
	// 640000000
	SuggestMaxRecordsPerPartition *int64 `json:"SuggestMaxRecordsPerPartition,omitempty" xml:"SuggestMaxRecordsPerPartition,omitempty"`
	// The recommended minimum number of rows in each partition.
	//
	// example:
	//
	// 64000000
	SuggestMinRecordsPerPartition *int64 `json:"SuggestMinRecordsPerPartition,omitempty" xml:"SuggestMinRecordsPerPartition,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 185
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTablePartitionDiagnoseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeTablePartitionDiagnoseResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTablePartitionDiagnoseResponseBody) GetDetectionItems() []*DescribeTablePartitionDiagnoseResponseBodyDetectionItems {
	return s.DetectionItems
}

func (s *DescribeTablePartitionDiagnoseResponseBody) GetItems() []*DescribeTablePartitionDiagnoseResponseBodyItems {
	return s.Items
}

func (s *DescribeTablePartitionDiagnoseResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTablePartitionDiagnoseResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTablePartitionDiagnoseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTablePartitionDiagnoseResponseBody) GetSuggestMaxRecordsPerPartition() *int64 {
	return s.SuggestMaxRecordsPerPartition
}

func (s *DescribeTablePartitionDiagnoseResponseBody) GetSuggestMinRecordsPerPartition() *int64 {
	return s.SuggestMinRecordsPerPartition
}

func (s *DescribeTablePartitionDiagnoseResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetAccessDeniedDetail(v string) *DescribeTablePartitionDiagnoseResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetDBClusterId(v string) *DescribeTablePartitionDiagnoseResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetDetectionItems(v []*DescribeTablePartitionDiagnoseResponseBodyDetectionItems) *DescribeTablePartitionDiagnoseResponseBody {
	s.DetectionItems = v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetItems(v []*DescribeTablePartitionDiagnoseResponseBodyItems) *DescribeTablePartitionDiagnoseResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetPageNumber(v int32) *DescribeTablePartitionDiagnoseResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetPageSize(v int32) *DescribeTablePartitionDiagnoseResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetRequestId(v string) *DescribeTablePartitionDiagnoseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetSuggestMaxRecordsPerPartition(v int64) *DescribeTablePartitionDiagnoseResponseBody {
	s.SuggestMaxRecordsPerPartition = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetSuggestMinRecordsPerPartition(v int64) *DescribeTablePartitionDiagnoseResponseBody {
	s.SuggestMinRecordsPerPartition = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetTotalCount(v int32) *DescribeTablePartitionDiagnoseResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) Validate() error {
	if s.DetectionItems != nil {
		for _, item := range s.DetectionItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTablePartitionDiagnoseResponseBodyDetectionItems struct {
	// The detection result.
	//
	// example:
	//
	// Multiple oversized non-partitioned tables are detected.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the detection item.
	//
	// example:
	//
	// Improper partition field diagnostics
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The severity level of the detection result. Valid values:
	//
	// 	- NORMAL
	//
	// 	- WARNING
	//
	// 	- CRITICAL
	//
	// example:
	//
	// WARNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeTablePartitionDiagnoseResponseBodyDetectionItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseResponseBodyDetectionItems) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseResponseBodyDetectionItems) GetMessage() *string {
	return s.Message
}

func (s *DescribeTablePartitionDiagnoseResponseBodyDetectionItems) GetName() *string {
	return s.Name
}

func (s *DescribeTablePartitionDiagnoseResponseBodyDetectionItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeTablePartitionDiagnoseResponseBodyDetectionItems) SetMessage(v string) *DescribeTablePartitionDiagnoseResponseBodyDetectionItems {
	s.Message = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyDetectionItems) SetName(v string) *DescribeTablePartitionDiagnoseResponseBodyDetectionItems {
	s.Name = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyDetectionItems) SetStatus(v string) *DescribeTablePartitionDiagnoseResponseBodyDetectionItems {
	s.Status = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyDetectionItems) Validate() error {
	return dara.Validate(s)
}

type DescribeTablePartitionDiagnoseResponseBodyItems struct {
	// The improper partitions.
	//
	// example:
	//
	// [20210110, 20210113,20210123]
	PartitionDetail *string `json:"PartitionDetail,omitempty" xml:"PartitionDetail,omitempty"`
	// The number of partitions.
	//
	// example:
	//
	// 2
	PartitionNumber *int32 `json:"PartitionNumber,omitempty" xml:"PartitionNumber,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test_db
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The storage percentage of the table. Unit: %.
	//
	// >  Formula: Table storage percentage = Total data size of a table/Total data size of the cluster × 100%.
	//
	// example:
	//
	// 66.23
	SpaceRatio *float64 `json:"SpaceRatio,omitempty" xml:"SpaceRatio,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// user
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The total data size of the table. Unit: bytes.
	//
	// example:
	//
	// 42949672960
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeTablePartitionDiagnoseResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) GetPartitionDetail() *string {
	return s.PartitionDetail
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) GetPartitionNumber() *int32 {
	return s.PartitionNumber
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) GetSpaceRatio() *float64 {
	return s.SpaceRatio
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) GetTableName() *string {
	return s.TableName
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetPartitionDetail(v string) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.PartitionDetail = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetPartitionNumber(v int32) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.PartitionNumber = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetSchemaName(v string) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetSpaceRatio(v float64) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.SpaceRatio = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetTableName(v string) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetTotalSize(v int64) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.TotalSize = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
