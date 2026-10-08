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
	// The list of aggregation field information.
	//
	// 	Notice: When you use an aggregation query, only the aggregation results are returned, and the list of matched metadata is not returned.</notice>
	AggregationsShrink *string `json:"Aggregations,omitempty" xml:"Aggregations,omitempty"`
	// The name of the dataset. For more information about how to obtain the dataset name, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// - When you perform a query for files without specifying the Aggregations parameter, this parameter specifies the maximum number of files to return. Valid values: 0 to 100.
	//
	// - When you specify the Aggregations parameter for aggregation statistics, this parameter specifies the maximum number of groups to return. Valid values: 0 to 2000.
	//
	// - If you do not specify this parameter or set it to 0, the default value is 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used for pagination when the total number of files exceeds the value of MaxResults.
	//
	// The list of files is returned in lexicographical order starting from NextToken.
	//
	// Set this parameter to empty when you call this operation for the first time.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpwZw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order of the sort fields. Valid values:
	//
	// - asc: ascending order
	//
	// - desc: descending order (default)
	//
	// >- You can separate multiple sort orders with commas (,), for example, asc,desc.
	//
	// > - The number of sort orders cannot exceed the number of sort fields. That is, the number of elements in the Order parameter must be less than or equal to the number of elements in the Sort parameter. For example, if Sort is set to Size,Filename, Order can be set to "asc,desc".
	//
	// > - If the number of sort orders is less than the number of sort fields, the default sort order for the unspecified fields is desc. For example, if Sort is set to Size,Filename and Order is set to asc, the default sort order for Filename is desc, which means descending order.
	//
	// example:
	//
	// asc,desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The name of the project. For more information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The simple query conditions. Click the link on the left to view details.
	QueryShrink *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The list of sort fields. For more information, see [Supported fields and operators](https://help.aliyun.com/document_detail/2743991.html).
	//
	// > - You can separate multiple sort fields with commas (,), for example, Size,Filename.
	//
	// > - You can specify a maximum of 5 sort fields.
	//
	// > - The order of the sort fields determines the sorting priority.
	//
	// example:
	//
	// Size,Filename
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Specifies the specific fields to return instead of all existing metadata fields. This can be used to reduce the size of the returned struct.
	//
	// If you do not specify this parameter or leave it empty, all fields are returned.
	WithFieldsShrink *string `json:"WithFields,omitempty" xml:"WithFields,omitempty"`
	// Specifies whether to return the total number of matched records. Valid values:
	//
	// - true: The TotalHits field is not returned.
	//
	// - false: The TotalHits field is returned.
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
