// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetFileMetasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetVersion(v string) *ListDatasetFileMetasRequest
	GetDatasetVersion() *string
	SetEndFileUpdateTime(v string) *ListDatasetFileMetasRequest
	GetEndFileUpdateTime() *string
	SetEndTagUpdateTime(v string) *ListDatasetFileMetasRequest
	GetEndTagUpdateTime() *string
	SetMaxResults(v int32) *ListDatasetFileMetasRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDatasetFileMetasRequest
	GetNextToken() *string
	SetOrder(v string) *ListDatasetFileMetasRequest
	GetOrder() *string
	SetPageSize(v int32) *ListDatasetFileMetasRequest
	GetPageSize() *int32
	SetQueryContentTypeIncludeAny(v []*string) *ListDatasetFileMetasRequest
	GetQueryContentTypeIncludeAny() []*string
	SetQueryExpression(v string) *ListDatasetFileMetasRequest
	GetQueryExpression() *string
	SetQueryFileDir(v string) *ListDatasetFileMetasRequest
	GetQueryFileDir() *string
	SetQueryFileName(v string) *ListDatasetFileMetasRequest
	GetQueryFileName() *string
	SetQueryFileTypeIncludeAny(v []*string) *ListDatasetFileMetasRequest
	GetQueryFileTypeIncludeAny() []*string
	SetQueryImage(v string) *ListDatasetFileMetasRequest
	GetQueryImage() *string
	SetQueryTagsExclude(v []*string) *ListDatasetFileMetasRequest
	GetQueryTagsExclude() []*string
	SetQueryTagsIncludeAll(v []*string) *ListDatasetFileMetasRequest
	GetQueryTagsIncludeAll() []*string
	SetQueryTagsIncludeAny(v []*string) *ListDatasetFileMetasRequest
	GetQueryTagsIncludeAny() []*string
	SetQueryText(v string) *ListDatasetFileMetasRequest
	GetQueryText() *string
	SetQueryType(v string) *ListDatasetFileMetasRequest
	GetQueryType() *string
	SetScoreThreshold(v float32) *ListDatasetFileMetasRequest
	GetScoreThreshold() *float32
	SetSortBy(v string) *ListDatasetFileMetasRequest
	GetSortBy() *string
	SetStartFileUpdateTime(v string) *ListDatasetFileMetasRequest
	GetStartFileUpdateTime() *string
	SetStartTagUpdateTime(v string) *ListDatasetFileMetasRequest
	GetStartTagUpdateTime() *string
	SetThumbnailMode(v string) *ListDatasetFileMetasRequest
	GetThumbnailMode() *string
	SetTopK(v int32) *ListDatasetFileMetasRequest
	GetTopK() *int32
	SetWorkspaceId(v string) *ListDatasetFileMetasRequest
	GetWorkspaceId() *string
}

type ListDatasetFileMetasRequest struct {
	// The dataset version.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The update time range to query. The end time. The time follows the ISO 8601 standard. This parameter is valid only when QueryType is set to TAG.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01Z
	EndFileUpdateTime *string `json:"EndFileUpdateTime,omitempty" xml:"EndFileUpdateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01.000Z
	EndTagUpdateTime *string `json:"EndTagUpdateTime,omitempty" xml:"EndTagUpdateTime,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// >  If you do not configure this parameter, the data on the first page is returned. A return value other than Null of this parameter indicates that not all entries have been returned. You can use this value as an input parameter to obtain entries on the next page. The value Null indicates that all query results have been returned.
	//
	// example:
	//
	// 90a6ee35-****-4cd4-927e-1f45e1cb8b62_1729644433000
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The order in which the entries are sorted by the specific field on the returned page. This parameter must be used together with SortBy. Default value: ASC.
	//
	// 	- ASC
	//
	// 	- DESC
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// Deprecated
	//
	// The number of entries per page. Default value: 10. Maximum value: 1000.
	//
	// example:
	//
	// 10
	PageSize                   *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryContentTypeIncludeAny []*string `json:"QueryContentTypeIncludeAny,omitempty" xml:"QueryContentTypeIncludeAny,omitempty" type:"Repeated"`
	QueryExpression            *string   `json:"QueryExpression,omitempty" xml:"QueryExpression,omitempty"`
	// example:
	//
	// cars/20250221/
	QueryFileDir *string `json:"QueryFileDir,omitempty" xml:"QueryFileDir,omitempty"`
	// example:
	//
	// shuima
	QueryFileName           *string   `json:"QueryFileName,omitempty" xml:"QueryFileName,omitempty"`
	QueryFileTypeIncludeAny []*string `json:"QueryFileTypeIncludeAny,omitempty" xml:"QueryFileTypeIncludeAny,omitempty" type:"Repeated"`
	// example:
	//
	// oss://test-xxx-oss/car/0001.png
	QueryImage          *string   `json:"QueryImage,omitempty" xml:"QueryImage,omitempty"`
	QueryTagsExclude    []*string `json:"QueryTagsExclude,omitempty" xml:"QueryTagsExclude,omitempty" type:"Repeated"`
	QueryTagsIncludeAll []*string `json:"QueryTagsIncludeAll,omitempty" xml:"QueryTagsIncludeAll,omitempty" type:"Repeated"`
	QueryTagsIncludeAny []*string `json:"QueryTagsIncludeAny,omitempty" xml:"QueryTagsIncludeAny,omitempty" type:"Repeated"`
	// The text content to be queried.
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// The retrieval type.
	//
	// 	- TAG (default)
	//
	// 	- VECTOR
	//
	// example:
	//
	// TAG
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The similarity score. Only dataset files whose similarity score is greater than the value of ScoreThreshold are returned. This parameter is valid only when QueryType is set to VECTOR.
	//
	// example:
	//
	// 0.6
	ScoreThreshold *float32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The field used to sort the results. Default value: GmtCreateTime. Valid values:
	//
	// 	- FileCreateTime (default): The results are sorted by the time when the file is created.
	//
	// 	- FileUpdateTime: The results are sorted by the time when the file is last modified.
	//
	// example:
	//
	// FileCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The update time range to query. The start time. The time follows the ISO 8601 standard. This parameter is valid only when QueryType is set to TAG.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01Z
	StartFileUpdateTime *string `json:"StartFileUpdateTime,omitempty" xml:"StartFileUpdateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01.000Z
	StartTagUpdateTime *string `json:"StartTagUpdateTime,omitempty" xml:"StartTagUpdateTime,omitempty"`
	// example:
	//
	// w_100
	ThumbnailMode *string `json:"ThumbnailMode,omitempty" xml:"ThumbnailMode,omitempty"`
	// The number of search results to return. A maximum of Top K search results can be returned. This parameter is valid only when QueryType is set to VECTOR.
	//
	// example:
	//
	// 100
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// The ID of the workspace to which the dataset belongs. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 105173
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasetFileMetasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetFileMetasRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetFileMetasRequest) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *ListDatasetFileMetasRequest) GetEndFileUpdateTime() *string {
	return s.EndFileUpdateTime
}

func (s *ListDatasetFileMetasRequest) GetEndTagUpdateTime() *string {
	return s.EndTagUpdateTime
}

func (s *ListDatasetFileMetasRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDatasetFileMetasRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDatasetFileMetasRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDatasetFileMetasRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetFileMetasRequest) GetQueryContentTypeIncludeAny() []*string {
	return s.QueryContentTypeIncludeAny
}

func (s *ListDatasetFileMetasRequest) GetQueryExpression() *string {
	return s.QueryExpression
}

func (s *ListDatasetFileMetasRequest) GetQueryFileDir() *string {
	return s.QueryFileDir
}

func (s *ListDatasetFileMetasRequest) GetQueryFileName() *string {
	return s.QueryFileName
}

func (s *ListDatasetFileMetasRequest) GetQueryFileTypeIncludeAny() []*string {
	return s.QueryFileTypeIncludeAny
}

func (s *ListDatasetFileMetasRequest) GetQueryImage() *string {
	return s.QueryImage
}

func (s *ListDatasetFileMetasRequest) GetQueryTagsExclude() []*string {
	return s.QueryTagsExclude
}

func (s *ListDatasetFileMetasRequest) GetQueryTagsIncludeAll() []*string {
	return s.QueryTagsIncludeAll
}

func (s *ListDatasetFileMetasRequest) GetQueryTagsIncludeAny() []*string {
	return s.QueryTagsIncludeAny
}

func (s *ListDatasetFileMetasRequest) GetQueryText() *string {
	return s.QueryText
}

func (s *ListDatasetFileMetasRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *ListDatasetFileMetasRequest) GetScoreThreshold() *float32 {
	return s.ScoreThreshold
}

func (s *ListDatasetFileMetasRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDatasetFileMetasRequest) GetStartFileUpdateTime() *string {
	return s.StartFileUpdateTime
}

func (s *ListDatasetFileMetasRequest) GetStartTagUpdateTime() *string {
	return s.StartTagUpdateTime
}

func (s *ListDatasetFileMetasRequest) GetThumbnailMode() *string {
	return s.ThumbnailMode
}

func (s *ListDatasetFileMetasRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *ListDatasetFileMetasRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDatasetFileMetasRequest) SetDatasetVersion(v string) *ListDatasetFileMetasRequest {
	s.DatasetVersion = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetEndFileUpdateTime(v string) *ListDatasetFileMetasRequest {
	s.EndFileUpdateTime = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetEndTagUpdateTime(v string) *ListDatasetFileMetasRequest {
	s.EndTagUpdateTime = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetMaxResults(v int32) *ListDatasetFileMetasRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetNextToken(v string) *ListDatasetFileMetasRequest {
	s.NextToken = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetOrder(v string) *ListDatasetFileMetasRequest {
	s.Order = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetPageSize(v int32) *ListDatasetFileMetasRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetQueryContentTypeIncludeAny(v []*string) *ListDatasetFileMetasRequest {
	s.QueryContentTypeIncludeAny = v
	return s
}

func (s *ListDatasetFileMetasRequest) SetQueryExpression(v string) *ListDatasetFileMetasRequest {
	s.QueryExpression = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetQueryFileDir(v string) *ListDatasetFileMetasRequest {
	s.QueryFileDir = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetQueryFileName(v string) *ListDatasetFileMetasRequest {
	s.QueryFileName = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetQueryFileTypeIncludeAny(v []*string) *ListDatasetFileMetasRequest {
	s.QueryFileTypeIncludeAny = v
	return s
}

func (s *ListDatasetFileMetasRequest) SetQueryImage(v string) *ListDatasetFileMetasRequest {
	s.QueryImage = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetQueryTagsExclude(v []*string) *ListDatasetFileMetasRequest {
	s.QueryTagsExclude = v
	return s
}

func (s *ListDatasetFileMetasRequest) SetQueryTagsIncludeAll(v []*string) *ListDatasetFileMetasRequest {
	s.QueryTagsIncludeAll = v
	return s
}

func (s *ListDatasetFileMetasRequest) SetQueryTagsIncludeAny(v []*string) *ListDatasetFileMetasRequest {
	s.QueryTagsIncludeAny = v
	return s
}

func (s *ListDatasetFileMetasRequest) SetQueryText(v string) *ListDatasetFileMetasRequest {
	s.QueryText = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetQueryType(v string) *ListDatasetFileMetasRequest {
	s.QueryType = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetScoreThreshold(v float32) *ListDatasetFileMetasRequest {
	s.ScoreThreshold = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetSortBy(v string) *ListDatasetFileMetasRequest {
	s.SortBy = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetStartFileUpdateTime(v string) *ListDatasetFileMetasRequest {
	s.StartFileUpdateTime = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetStartTagUpdateTime(v string) *ListDatasetFileMetasRequest {
	s.StartTagUpdateTime = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetThumbnailMode(v string) *ListDatasetFileMetasRequest {
	s.ThumbnailMode = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetTopK(v int32) *ListDatasetFileMetasRequest {
	s.TopK = &v
	return s
}

func (s *ListDatasetFileMetasRequest) SetWorkspaceId(v string) *ListDatasetFileMetasRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDatasetFileMetasRequest) Validate() error {
	return dara.Validate(s)
}
