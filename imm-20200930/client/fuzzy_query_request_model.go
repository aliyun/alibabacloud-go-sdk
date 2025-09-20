// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFuzzyQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *FuzzyQueryRequest
	GetDatasetName() *string
	SetMaxResults(v int64) *FuzzyQueryRequest
	GetMaxResults() *int64
	SetNextToken(v string) *FuzzyQueryRequest
	GetNextToken() *string
	SetOrder(v string) *FuzzyQueryRequest
	GetOrder() *string
	SetProjectName(v string) *FuzzyQueryRequest
	GetProjectName() *string
	SetQuery(v string) *FuzzyQueryRequest
	GetQuery() *string
	SetSort(v string) *FuzzyQueryRequest
	GetSort() *string
	SetWithFields(v []*string) *FuzzyQueryRequest
	GetWithFields() []*string
}

type FuzzyQueryRequest struct {
	// The name of the dataset. You can obtain the name of the dataset from the response of the [CreateDataset](https://help.aliyun.com/document_detail/478160.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The maximum number of entries to return. Valid values: 0 to 200.
	//
	// Default value: 100.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the total number of files is greater than the value of MaxResults, you must specify NextToken.
	//
	// The file information is returned in alphabetical order starting from the value of NextToken.
	//
	// You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpwZw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sorting method. Valid values:
	//
	// 	- asc: ascending order.
	//
	// 	- desc (default): descending order.
	//
	// >
	//
	// 	- Separate multiple sorting methods with commas (,). Example: asc,desc.
	//
	// 	- The number of values for Order must be less than or equal to the number of values for Sort. For example, if you set Sort to Size,Filename, you can set Order only to desc or asc.
	//
	// 	- If the number of values for Order is less than the number of values for Sort, the unsorted fields are default to the value of asc. For example, if you set Sort to Size,Filename and Order to asc, the Filename field is default to the value of asc.
	//
	// example:
	//
	// asc,desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The query content. The value can be up to 1 MB in size.
	//
	// This parameter is required.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The sort fields. For more information, see [Supported fields and operators](https://help.aliyun.com/document_detail/2743991.html).
	//
	// 	- Separate multiple sort fields with commas (,). Example: `Size,Filename`.
	//
	// 	- You can specify up to five sort fields.
	//
	// 	- The priority order of sorting is determined based on the order of the sort fields.
	//
	// example:
	//
	// Size,Filename
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The fields that you want to include in the response. To help reduce the size of the response, include only necessary metadata fields.
	//
	// If you do not specify this parameter or set the value to null, all existing metadata fields are returned.
	WithFields []*string `json:"WithFields,omitempty" xml:"WithFields,omitempty" type:"Repeated"`
}

func (s FuzzyQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s FuzzyQueryRequest) GoString() string {
	return s.String()
}

func (s *FuzzyQueryRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *FuzzyQueryRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *FuzzyQueryRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *FuzzyQueryRequest) GetOrder() *string {
	return s.Order
}

func (s *FuzzyQueryRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *FuzzyQueryRequest) GetQuery() *string {
	return s.Query
}

func (s *FuzzyQueryRequest) GetSort() *string {
	return s.Sort
}

func (s *FuzzyQueryRequest) GetWithFields() []*string {
	return s.WithFields
}

func (s *FuzzyQueryRequest) SetDatasetName(v string) *FuzzyQueryRequest {
	s.DatasetName = &v
	return s
}

func (s *FuzzyQueryRequest) SetMaxResults(v int64) *FuzzyQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *FuzzyQueryRequest) SetNextToken(v string) *FuzzyQueryRequest {
	s.NextToken = &v
	return s
}

func (s *FuzzyQueryRequest) SetOrder(v string) *FuzzyQueryRequest {
	s.Order = &v
	return s
}

func (s *FuzzyQueryRequest) SetProjectName(v string) *FuzzyQueryRequest {
	s.ProjectName = &v
	return s
}

func (s *FuzzyQueryRequest) SetQuery(v string) *FuzzyQueryRequest {
	s.Query = &v
	return s
}

func (s *FuzzyQueryRequest) SetSort(v string) *FuzzyQueryRequest {
	s.Sort = &v
	return s
}

func (s *FuzzyQueryRequest) SetWithFields(v []*string) *FuzzyQueryRequest {
	s.WithFields = v
	return s
}

func (s *FuzzyQueryRequest) Validate() error {
	return dara.Validate(s)
}
