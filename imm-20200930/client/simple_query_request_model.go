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
	// The list of aggregation field information.
	//
	// 	Notice: When you use an aggregation query, only the aggregation results are returned, and the list of matched metadata is not returned.</notice>
	Aggregations []*SimpleQueryRequestAggregations `json:"Aggregations,omitempty" xml:"Aggregations,omitempty" type:"Repeated"`
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
	Query *SimpleQuery `json:"Query,omitempty" xml:"Query,omitempty"`
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
	WithFields []*string `json:"WithFields,omitempty" xml:"WithFields,omitempty" type:"Repeated"`
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
	if s.Aggregations != nil {
		for _, item := range s.Aggregations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Query != nil {
		if err := s.Query.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SimpleQueryRequestAggregations struct {
	// The name of the field. For more information about supported fields, see [Supported fields and operators](https://help.aliyun.com/document_detail/2743991.html).
	//
	// example:
	//
	// Size
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The operator for the aggregation field.
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
