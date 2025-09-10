// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchHistoricalSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEdition(v string) *SearchHistoricalSnapshotsRequest
	GetEdition() *string
	SetLimit(v int32) *SearchHistoricalSnapshotsRequest
	GetLimit() *int32
	SetNextToken(v string) *SearchHistoricalSnapshotsRequest
	GetNextToken() *string
	SetOrder(v string) *SearchHistoricalSnapshotsRequest
	GetOrder() *string
	SetQuery(v []interface{}) *SearchHistoricalSnapshotsRequest
	GetQuery() []interface{}
	SetSortBy(v string) *SearchHistoricalSnapshotsRequest
	GetSortBy() *string
	SetSourceType(v string) *SearchHistoricalSnapshotsRequest
	GetSourceType() *string
}

type SearchHistoricalSnapshotsRequest struct {
	// example:
	//
	// BASIC
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The maximum number of rows that you want the current query to return. To query only the number of matched rows without the need to return specific data, you can set the Limit parameter to `0`. Then, the operation returns only the number of matched rows.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token that is required to obtain the next page of backup snapshots.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ordering mode. Valid values:
	//
	// 	- ASC (default): ascending order
	//
	// 	- DESC: descending order
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The query conditions. Example:
	//
	//     [
	//
	//       {
	//
	//         "field": "VaultId",
	//
	//         "value": "v-0003rf9m*****qx5",
	//
	//         "operation": "MATCH_TERM"
	//
	//       },
	//
	//       {
	//
	//         "field": "InstanceId",
	//
	//         "value": "i-bp1i20zq2*****e9368m",
	//
	//         "operation": "MATCH_TERM"
	//
	//       },
	//
	//       {
	//
	//         "field": "PlanId",
	//
	//         "value": "plan-0005vk*****gkd1iu4f",
	//
	//         "operation": "MATCH_TERM"
	//
	//       },
	//
	//       {
	//
	//         "field": "CompleteTime",
	//
	//         "value": "1626769913",
	//
	//         "operation": "GREATER_THAN_OR_EQUAL"
	//
	//       }
	//
	//     ]
	//
	// 	- The following fields are supported:
	//
	//     	- VaultId: specifies the ID of the backup vault. This field is required.
	//
	//     	- InstanceId: specifies the ID of the Elastic Compute Service (ECS) instance. If the SourceType parameter is set to ECS_FILE, this field is required.
	//
	//     	- Bucket: specifies the name of the Object Storage Service (OSS) bucket. If the SourceType parameter is set to OSS, this field is required.
	//
	//     	- FileSystemId: specifies the ID of the File Storage NAS (NAS) file system. If the SourceType parameter is set to NAS, this field is required.
	//
	//     	- CreateTime: specifies the time when the NAS file system was created. If the SourceType parameter is set to NAS, this field is required.
	//
	//     	- CompleteTime: specifies the time when the backup snapshot was completed.
	//
	//     	- PlanId: the ID of a backup plan.
	//
	// 	- The following operations are supported:
	//
	//     	- MATCH_TERM: exact match.
	//
	//     	- GREATER_THAN: greater than.
	//
	//     	- GREATER_THAN_OR_EQUAL: greater than or equal to.
	//
	//     	- LESS_THAN: less than.
	//
	//     	- LESS_THAN_OR_EQUAL: less than or equal to.
	//
	//     	- BETWEEN: specifies a JSON array as a range. The results must fall within the range in the `[Minimum value,Maximum value]` format.
	//
	//     	- IN: specifies an array as a collection. The results must fall within the collection.
	//
	//     	- NOT_IN: specifies an array as a collection. The results cannot fall within the collection.
	//
	// example:
	//
	// [   {     "field": "VaultId",     "value": "v-0003rf9m17pap3ltpqx5",     "operation": "MATCH_TERM"   },   {     "field": "InstanceId",     "value": "i-bp1i20zq2wuzdie9368m",     "operation": "MATCH_TERM"   },   {     "field": "PlanId",     "value": "plan-0005vkqhpesqgkd1iu4f",     "operation": "MATCH_TERM"   },   {     "field": "CompleteTime",     "value": 1626769913,     "operation": "GREATER_THAN_OR_EQUAL"   } ]
	Query []interface{} `json:"Query,omitempty" xml:"Query,omitempty" type:"Repeated"`
	// The field that is used to sort data.
	//
	// example:
	//
	// CreatedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: backup snapshots for Elastic Compute Service (ECS) files
	//
	// 	- **OSS**: backup snapshots for Object Storage Service (OSS) buckets
	//
	// 	- **NAS**: backup snapshots for Apsara File Storage NAS file systems
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s SearchHistoricalSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchHistoricalSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *SearchHistoricalSnapshotsRequest) GetEdition() *string {
	return s.Edition
}

func (s *SearchHistoricalSnapshotsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *SearchHistoricalSnapshotsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchHistoricalSnapshotsRequest) GetOrder() *string {
	return s.Order
}

func (s *SearchHistoricalSnapshotsRequest) GetQuery() []interface{} {
	return s.Query
}

func (s *SearchHistoricalSnapshotsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *SearchHistoricalSnapshotsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *SearchHistoricalSnapshotsRequest) SetEdition(v string) *SearchHistoricalSnapshotsRequest {
	s.Edition = &v
	return s
}

func (s *SearchHistoricalSnapshotsRequest) SetLimit(v int32) *SearchHistoricalSnapshotsRequest {
	s.Limit = &v
	return s
}

func (s *SearchHistoricalSnapshotsRequest) SetNextToken(v string) *SearchHistoricalSnapshotsRequest {
	s.NextToken = &v
	return s
}

func (s *SearchHistoricalSnapshotsRequest) SetOrder(v string) *SearchHistoricalSnapshotsRequest {
	s.Order = &v
	return s
}

func (s *SearchHistoricalSnapshotsRequest) SetQuery(v []interface{}) *SearchHistoricalSnapshotsRequest {
	s.Query = v
	return s
}

func (s *SearchHistoricalSnapshotsRequest) SetSortBy(v string) *SearchHistoricalSnapshotsRequest {
	s.SortBy = &v
	return s
}

func (s *SearchHistoricalSnapshotsRequest) SetSourceType(v string) *SearchHistoricalSnapshotsRequest {
	s.SourceType = &v
	return s
}

func (s *SearchHistoricalSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
