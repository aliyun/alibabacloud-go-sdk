// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleQueryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregationsShrink(v string) *SimpleQueryShrinkRequest
	GetAggregationsShrink() *string
	SetDatasetName(v string) *SimpleQueryShrinkRequest
	GetDatasetName() *string
	SetMaxResults(v int32) *SimpleQueryShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *SimpleQueryShrinkRequest
	GetNextToken() *string
	SetOrder(v string) *SimpleQueryShrinkRequest
	GetOrder() *string
	SetProjectName(v string) *SimpleQueryShrinkRequest
	GetProjectName() *string
	SetQueryShrink(v string) *SimpleQueryShrinkRequest
	GetQueryShrink() *string
	SetSort(v string) *SimpleQueryShrinkRequest
	GetSort() *string
	SetWithFieldsShrink(v string) *SimpleQueryShrinkRequest
	GetWithFieldsShrink() *string
	SetWithoutTotalHits(v bool) *SimpleQueryShrinkRequest
	GetWithoutTotalHits() *bool
}

type SimpleQueryShrinkRequest struct {
	// The aggregations.
	//
	// >  If you perform an aggregate query, the aggregation returned in the response contains only statistical results, not the actual metadata.
	AggregationsShrink *string `json:"Aggregations,omitempty" xml:"Aggregations,omitempty"`
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 	- If the Aggregations parameter is not specified, this parameter specifies the maximum number of files that can be returned. Valid values: 1 to 100.
	//
	// 	- If the Aggregations parameter is specified, this parameter specifies the maximum number of aggregation groups that can be returned. Valid values: 0 to 2000.
	//
	// 	- If you do not specify this parameter or set the parameter to 0, the default value of 100 is used.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token is used in the next request to retrieve a new page of results if the total number of results exceeds the value of the MaxResults parameter.
	//
	// The next call to the operation returns results lexicographically after the NextToken parameter value.
	//
	// You do not need to specify this parameter in your initial request.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpwZw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order. Valid values:
	//
	// 	- asc: sorts the results in ascending order.
	//
	// 	- desc: sorts the results in descending order. This is the default value.
	//
	// 	- You can specify multiple sort orders that are separated by commas. Example: asc,desc.
	//
	// 	- The number of elements in the Order parameter must be less than or equal to the number of elements in the Sort parameter. For example, if the value of the Sort parameter is Size,Filename, you can set the Order parameter to desc,asc.
	//
	// 	- If the number of sort orders is less than the number of sort fields, the sort fields for which no sorting orders are explicitly specified use the asc order by default. For example, if you set Sort to Size,Filename and Order to asc, the Filename field defaults to the value of asc.
	//
	// example:
	//
	// asc,desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The query conditions.
	QueryShrink *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The sort fields. For more information, see [Supported fields and operators](https://help.aliyun.com/document_detail/2743991.html).
	//
	// >
	//
	// 	- If you specify multiple sort fields, separate them with commas (,), as in Size,Filename.
	//
	// 	- You can specify up to five sort fields.
	//
	// 	- The order of the sort fields determines their precedence in the sorting process.
	//
	// example:
	//
	// Size,Filename
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The fields that you want to include in the response. You can use this parameter to reduce the size of the response.
	//
	// If you do not specify this parameter or leave this parameter empty, the operation returns all metadata fields.
	WithFieldsShrink *string `json:"WithFields,omitempty" xml:"WithFields,omitempty"`
	// Specifies whether to return the total number of hits. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// if can be null:
	// true
	WithoutTotalHits *bool `json:"WithoutTotalHits,omitempty" xml:"WithoutTotalHits,omitempty"`
}

func (s SimpleQueryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SimpleQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *SimpleQueryShrinkRequest) GetAggregationsShrink() *string {
	return s.AggregationsShrink
}

func (s *SimpleQueryShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *SimpleQueryShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SimpleQueryShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SimpleQueryShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *SimpleQueryShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *SimpleQueryShrinkRequest) GetQueryShrink() *string {
	return s.QueryShrink
}

func (s *SimpleQueryShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *SimpleQueryShrinkRequest) GetWithFieldsShrink() *string {
	return s.WithFieldsShrink
}

func (s *SimpleQueryShrinkRequest) GetWithoutTotalHits() *bool {
	return s.WithoutTotalHits
}

func (s *SimpleQueryShrinkRequest) SetAggregationsShrink(v string) *SimpleQueryShrinkRequest {
	s.AggregationsShrink = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetDatasetName(v string) *SimpleQueryShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetMaxResults(v int32) *SimpleQueryShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetNextToken(v string) *SimpleQueryShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetOrder(v string) *SimpleQueryShrinkRequest {
	s.Order = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetProjectName(v string) *SimpleQueryShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetQueryShrink(v string) *SimpleQueryShrinkRequest {
	s.QueryShrink = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetSort(v string) *SimpleQueryShrinkRequest {
	s.Sort = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetWithFieldsShrink(v string) *SimpleQueryShrinkRequest {
	s.WithFieldsShrink = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetWithoutTotalHits(v bool) *SimpleQueryShrinkRequest {
	s.WithoutTotalHits = &v
	return s
}

func (s *SimpleQueryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
