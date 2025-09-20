// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregations(v []*SimpleQueryRequestAggregations) *SimpleQueryRequest
	GetAggregations() []*SimpleQueryRequestAggregations
	SetDatasetName(v string) *SimpleQueryRequest
	GetDatasetName() *string
	SetMaxResults(v int32) *SimpleQueryRequest
	GetMaxResults() *int32
	SetNextToken(v string) *SimpleQueryRequest
	GetNextToken() *string
	SetOrder(v string) *SimpleQueryRequest
	GetOrder() *string
	SetProjectName(v string) *SimpleQueryRequest
	GetProjectName() *string
	SetQuery(v *SimpleQuery) *SimpleQueryRequest
	GetQuery() *SimpleQuery
	SetSort(v string) *SimpleQueryRequest
	GetSort() *string
	SetWithFields(v []*string) *SimpleQueryRequest
	GetWithFields() []*string
	SetWithoutTotalHits(v bool) *SimpleQueryRequest
	GetWithoutTotalHits() *bool
}

type SimpleQueryRequest struct {
	// The aggregations.
	//
	// >  If you perform an aggregate query, the aggregation returned in the response contains only statistical results, not the actual metadata.
	Aggregations []*SimpleQueryRequestAggregations `json:"Aggregations,omitempty" xml:"Aggregations,omitempty" type:"Repeated"`
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
	Query *SimpleQuery `json:"Query,omitempty" xml:"Query,omitempty"`
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
	WithFields []*string `json:"WithFields,omitempty" xml:"WithFields,omitempty" type:"Repeated"`
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

func (s SimpleQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s SimpleQueryRequest) GoString() string {
	return s.String()
}

func (s *SimpleQueryRequest) GetAggregations() []*SimpleQueryRequestAggregations {
	return s.Aggregations
}

func (s *SimpleQueryRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *SimpleQueryRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SimpleQueryRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SimpleQueryRequest) GetOrder() *string {
	return s.Order
}

func (s *SimpleQueryRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *SimpleQueryRequest) GetQuery() *SimpleQuery {
	return s.Query
}

func (s *SimpleQueryRequest) GetSort() *string {
	return s.Sort
}

func (s *SimpleQueryRequest) GetWithFields() []*string {
	return s.WithFields
}

func (s *SimpleQueryRequest) GetWithoutTotalHits() *bool {
	return s.WithoutTotalHits
}

func (s *SimpleQueryRequest) SetAggregations(v []*SimpleQueryRequestAggregations) *SimpleQueryRequest {
	s.Aggregations = v
	return s
}

func (s *SimpleQueryRequest) SetDatasetName(v string) *SimpleQueryRequest {
	s.DatasetName = &v
	return s
}

func (s *SimpleQueryRequest) SetMaxResults(v int32) *SimpleQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *SimpleQueryRequest) SetNextToken(v string) *SimpleQueryRequest {
	s.NextToken = &v
	return s
}

func (s *SimpleQueryRequest) SetOrder(v string) *SimpleQueryRequest {
	s.Order = &v
	return s
}

func (s *SimpleQueryRequest) SetProjectName(v string) *SimpleQueryRequest {
	s.ProjectName = &v
	return s
}

func (s *SimpleQueryRequest) SetQuery(v *SimpleQuery) *SimpleQueryRequest {
	s.Query = v
	return s
}

func (s *SimpleQueryRequest) SetSort(v string) *SimpleQueryRequest {
	s.Sort = &v
	return s
}

func (s *SimpleQueryRequest) SetWithFields(v []*string) *SimpleQueryRequest {
	s.WithFields = v
	return s
}

func (s *SimpleQueryRequest) SetWithoutTotalHits(v bool) *SimpleQueryRequest {
	s.WithoutTotalHits = &v
	return s
}

func (s *SimpleQueryRequest) Validate() error {
	return dara.Validate(s)
}

type SimpleQueryRequestAggregations struct {
	// The name of the field. For more information about supported fields, see [Supported fields and operators](https://help.aliyun.com/document_detail/2743991.html).
	//
	// example:
	//
	// Size
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The operator.
	//
	// Enumerated values:
	//
	// 	- average: calculates the average number.
	//
	// 	- min: finds the minimum value.
	//
	// 	- max: finds the maximum value.
	//
	// 	- count: counts the number of results.
	//
	// 	- distinct: counts the number of distinct results.
	//
	// 	- sum: calculates the sum of all matching results..
	//
	// 	- group: counts the number of results by group. The results are sorted by the count number in descending order.
	//
	// example:
	//
	// sum
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s SimpleQueryRequestAggregations) String() string {
	return dara.Prettify(s)
}

func (s SimpleQueryRequestAggregations) GoString() string {
	return s.String()
}

func (s *SimpleQueryRequestAggregations) GetField() *string {
	return s.Field
}

func (s *SimpleQueryRequestAggregations) GetOperation() *string {
	return s.Operation
}

func (s *SimpleQueryRequestAggregations) SetField(v string) *SimpleQueryRequestAggregations {
	s.Field = &v
	return s
}

func (s *SimpleQueryRequestAggregations) SetOperation(v string) *SimpleQueryRequestAggregations {
	s.Operation = &v
	return s
}

func (s *SimpleQueryRequestAggregations) Validate() error {
	return dara.Validate(s)
}
