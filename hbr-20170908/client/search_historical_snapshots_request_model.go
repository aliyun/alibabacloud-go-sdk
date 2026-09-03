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
	// The edition type. Valid values: BASIC and STANDARD. Default value: STANDARD.
	//
	// example:
	//
	// BASIC
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The maximum number of results to return for this query.
	//
	// If you only want to obtain the row count without actual data, set `limit=0` to return no rows.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token required to retrieve the next page of snapshots.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order. Default value: ASC.
	//
	// - ASC: ascending order.
	//
	// - DESC: descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The query conditions. Example of Query:
	//
	// ```
	//
	// [
	//
	//   {
	//
	//     "field": "VaultId",
	//
	//     "value": "v-0003rf9m*****qx5",
	//
	//     "operation": "MATCH_TERM"
	//
	//   },
	//
	//   {
	//
	//     "field": "InstanceId",
	//
	//     "value": "i-bp1i20zq2*****e9368m",
	//
	//     "operation": "MATCH_TERM"
	//
	//   },
	//
	//   {
	//
	//     "field": "PlanId",
	//
	//     "value": "plan-0005vk*****gkd1iu4f",
	//
	//     "operation": "MATCH_TERM"
	//
	//   },
	//
	//   {
	//
	//     "field": "CompleteTime",
	//
	//     "value": "1626769913",
	//
	//     "operation": "GREATER_THAN_OR_EQUAL"
	//
	//   }
	//
	// ]
	//
	// ```
	//
	// - Supported fields:
	//
	//     - VaultId: required. The vault ID.
	//
	//     - InstanceId: required only when SourceType=ECS_FILE. The ECS instance ID.
	//
	//     - Bucket: required only when SourceType=OSS. The OSS bucket name.
	//
	//     - FileSystemId: required only when SourceType=NAS. The NAS file system ID.
	//
	//     - CreateTime: required only when SourceType=NAS. The creation time of the NAS file system.
	//
	//     - CompleteTime: the time when the snapshot is completed.
	//
	//     - PlanId: the backup plan ID.
	//
	// - Supported operations:
	//
	//     - MATCH_TERM: exact match.
	//
	//     - GREATER_THAN: greater than.
	//
	//     - GREATER_THAN_OR_EQUAL: greater than or equal to.
	//
	//     - LESS_THAN: less than.
	//
	//     - LESS_THAN_OR_EQUAL: less than or equal to.
	//
	//     - BETWEEN: range. The value is a JSON array `[lower bound, upper bound]`.
	//
	//     - IN: in the set. The value is an array.
	//
	//     - NOT_IN: not in the set. The value is an array.
	//
	// example:
	//
	// [   {     "field": "VaultId",     "value": "v-0003rf9m17pap3ltpqx5",     "operation": "MATCH_TERM"   },   {     "field": "InstanceId",     "value": "i-bp1i20zq2wuzdie9368m",     "operation": "MATCH_TERM"   },   {     "field": "PlanId",     "value": "plan-0005vkqhpesqgkd1iu4f",     "operation": "MATCH_TERM"   },   {     "field": "CompleteTime",     "value": 1626769913,     "operation": "GREATER_THAN_OR_EQUAL"   } ]
	Query []interface{} `json:"Query,omitempty" xml:"Query,omitempty" type:"Repeated"`
	// The field used for sorting.
	//
	// example:
	//
	// CreatedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: backup snapshots of ECS files.
	//
	// 	- **OSS**: backup snapshots of Alibaba Cloud OSS.
	//
	// 	- **NAS**: backup snapshots of Alibaba Cloud NAS.
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
