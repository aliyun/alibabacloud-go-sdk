// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFuzzyQueryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *FuzzyQueryShrinkRequest
	GetDatasetName() *string
	SetMaxResults(v int64) *FuzzyQueryShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *FuzzyQueryShrinkRequest
	GetNextToken() *string
	SetOrder(v string) *FuzzyQueryShrinkRequest
	GetOrder() *string
	SetProjectName(v string) *FuzzyQueryShrinkRequest
	GetProjectName() *string
	SetQuery(v string) *FuzzyQueryShrinkRequest
	GetQuery() *string
	SetSort(v string) *FuzzyQueryShrinkRequest
	GetSort() *string
	SetWithFieldsShrink(v string) *FuzzyQueryShrinkRequest
	GetWithFieldsShrink() *string
}

type FuzzyQueryShrinkRequest struct {
	// The name of the dataset. For more information about how to obtain the dataset name, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The maximum number of files to return. Valid values: 0 to 200.
	//
	// If you do not set this parameter or set it to 0, the default value is 100.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used for pagination when the total number of files exceeds the value of MaxResults.
	//
	// The list of file information is returned in lexicographical order starting from NextToken.
	//
	// Set this parameter to empty when you call this operation for the first time.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpwZw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order of the sort fields. Valid values:
	//
	// - asc: Ascending order.
	//
	// - desc: Descending order. This is the default value.
	//
	// > - You can separate multiple sort orders with commas (,), such as asc,desc.
	//
	// > - The number of sort orders cannot exceed the number of sort fields. That is, the number of elements in the Order parameter must be less than or equal to the number of elements in the Sort parameter. For example, if Sort is set to Size,Filename, Order can be set to desc or asc.
	//
	// > - If the number of sort orders is less than the number of sort fields, the default sort order for the unspecified fields is asc. For example, if Sort is set to Size,Filename and Order is set to asc, the default sort order for Filename is asc, which means ascending order.
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
	// The string used for the query. The string cannot exceed 1 MB in size.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba Cloud
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The list of fields by which to sort the results. For more information, see the [list of supported fields and operators](https://help.aliyun.com/document_detail/2743991.html).
	//
	// - You can separate multiple sort fields with commas (,), such as `Size,Filename`.
	//
	// - You can specify up to 5 sort fields.
	//
	// - The order of the sort fields determines the sorting priority.
	//
	// example:
	//
	// Size,Filename
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Specifies the fields to return. Only the values of the specified fields are returned instead of all existing metadata fields. You can use this parameter to reduce the size of the returned struct.
	//
	// If you do not specify this parameter or leave it empty, all fields are returned.
	WithFieldsShrink *string `json:"WithFields,omitempty" xml:"WithFields,omitempty"`
}

func (s FuzzyQueryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FuzzyQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *FuzzyQueryShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *FuzzyQueryShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *FuzzyQueryShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *FuzzyQueryShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *FuzzyQueryShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *FuzzyQueryShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *FuzzyQueryShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *FuzzyQueryShrinkRequest) GetWithFieldsShrink() *string {
	return s.WithFieldsShrink
}

func (s *FuzzyQueryShrinkRequest) SetDatasetName(v string) *FuzzyQueryShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *FuzzyQueryShrinkRequest) SetMaxResults(v int64) *FuzzyQueryShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *FuzzyQueryShrinkRequest) SetNextToken(v string) *FuzzyQueryShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *FuzzyQueryShrinkRequest) SetOrder(v string) *FuzzyQueryShrinkRequest {
	s.Order = &v
	return s
}

func (s *FuzzyQueryShrinkRequest) SetProjectName(v string) *FuzzyQueryShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *FuzzyQueryShrinkRequest) SetQuery(v string) *FuzzyQueryShrinkRequest {
	s.Query = &v
	return s
}

func (s *FuzzyQueryShrinkRequest) SetSort(v string) *FuzzyQueryShrinkRequest {
	s.Sort = &v
	return s
}

func (s *FuzzyQueryShrinkRequest) SetWithFieldsShrink(v string) *FuzzyQueryShrinkRequest {
	s.WithFieldsShrink = &v
	return s
}

func (s *FuzzyQueryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
