// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySimilarImageClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomLabels(v string) *QuerySimilarImageClustersRequest
	GetCustomLabels() *string
	SetDatasetName(v string) *QuerySimilarImageClustersRequest
	GetDatasetName() *string
	SetMaxResults(v int32) *QuerySimilarImageClustersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QuerySimilarImageClustersRequest
	GetNextToken() *string
	SetOrder(v string) *QuerySimilarImageClustersRequest
	GetOrder() *string
	SetProjectName(v string) *QuerySimilarImageClustersRequest
	GetProjectName() *string
	SetSort(v string) *QuerySimilarImageClustersRequest
	GetSort() *string
}

type QuerySimilarImageClustersRequest struct {
	// The custom tags, which are used to filter tasks.
	//
	// example:
	//
	// {"key": "val"}
	CustomLabels *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset. For more information, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The number of entries per page. Value range: 0 to 100. Default value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// If the total number of clusters is greater than the value of MaxResults, you must specify this parameter. The next call to the operation returns results lexicographically after the NextToken parameter value.
	//
	// >  The first time you call this operation in a query, set this parameter to null.
	//
	// example:
	//
	// CAESEgoQCg4KClVwZGF0ZVRpbWUQARgBIs8ECgkAAJLUwUCAQ****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sorting order. Valid values:
	//
	// 	- asc: ascending order.
	//
	// 	- desc: descending order. This is the default value.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The sorting field.
	//
	// 	- CreateTime: the time when the clusters were created.
	//
	// 	- UpdateTime: the time when the clusters were updated. This is the default value.
	//
	// example:
	//
	// UpdateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s QuerySimilarImageClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySimilarImageClustersRequest) GoString() string {
	return s.String()
}

func (s *QuerySimilarImageClustersRequest) GetCustomLabels() *string {
	return s.CustomLabels
}

func (s *QuerySimilarImageClustersRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *QuerySimilarImageClustersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QuerySimilarImageClustersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QuerySimilarImageClustersRequest) GetOrder() *string {
	return s.Order
}

func (s *QuerySimilarImageClustersRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *QuerySimilarImageClustersRequest) GetSort() *string {
	return s.Sort
}

func (s *QuerySimilarImageClustersRequest) SetCustomLabels(v string) *QuerySimilarImageClustersRequest {
	s.CustomLabels = &v
	return s
}

func (s *QuerySimilarImageClustersRequest) SetDatasetName(v string) *QuerySimilarImageClustersRequest {
	s.DatasetName = &v
	return s
}

func (s *QuerySimilarImageClustersRequest) SetMaxResults(v int32) *QuerySimilarImageClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *QuerySimilarImageClustersRequest) SetNextToken(v string) *QuerySimilarImageClustersRequest {
	s.NextToken = &v
	return s
}

func (s *QuerySimilarImageClustersRequest) SetOrder(v string) *QuerySimilarImageClustersRequest {
	s.Order = &v
	return s
}

func (s *QuerySimilarImageClustersRequest) SetProjectName(v string) *QuerySimilarImageClustersRequest {
	s.ProjectName = &v
	return s
}

func (s *QuerySimilarImageClustersRequest) SetSort(v string) *QuerySimilarImageClustersRequest {
	s.Sort = &v
	return s
}

func (s *QuerySimilarImageClustersRequest) Validate() error {
	return dara.Validate(s)
}
