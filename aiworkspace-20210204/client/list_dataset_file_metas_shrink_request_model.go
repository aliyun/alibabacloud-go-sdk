// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetFileMetasShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetVersion(v string) *ListDatasetFileMetasShrinkRequest
	GetDatasetVersion() *string
	SetEndFileUpdateTime(v string) *ListDatasetFileMetasShrinkRequest
	GetEndFileUpdateTime() *string
	SetEndTagUpdateTime(v string) *ListDatasetFileMetasShrinkRequest
	GetEndTagUpdateTime() *string
	SetMaxResults(v int32) *ListDatasetFileMetasShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDatasetFileMetasShrinkRequest
	GetNextToken() *string
	SetOrder(v string) *ListDatasetFileMetasShrinkRequest
	GetOrder() *string
	SetPageSize(v int32) *ListDatasetFileMetasShrinkRequest
	GetPageSize() *int32
	SetQueryContentTypeIncludeAnyShrink(v string) *ListDatasetFileMetasShrinkRequest
	GetQueryContentTypeIncludeAnyShrink() *string
	SetQueryExpression(v string) *ListDatasetFileMetasShrinkRequest
	GetQueryExpression() *string
	SetQueryFileDir(v string) *ListDatasetFileMetasShrinkRequest
	GetQueryFileDir() *string
	SetQueryFileName(v string) *ListDatasetFileMetasShrinkRequest
	GetQueryFileName() *string
	SetQueryFileTypeIncludeAnyShrink(v string) *ListDatasetFileMetasShrinkRequest
	GetQueryFileTypeIncludeAnyShrink() *string
	SetQueryImage(v string) *ListDatasetFileMetasShrinkRequest
	GetQueryImage() *string
	SetQueryTagsExcludeShrink(v string) *ListDatasetFileMetasShrinkRequest
	GetQueryTagsExcludeShrink() *string
	SetQueryTagsIncludeAllShrink(v string) *ListDatasetFileMetasShrinkRequest
	GetQueryTagsIncludeAllShrink() *string
	SetQueryTagsIncludeAnyShrink(v string) *ListDatasetFileMetasShrinkRequest
	GetQueryTagsIncludeAnyShrink() *string
	SetQueryText(v string) *ListDatasetFileMetasShrinkRequest
	GetQueryText() *string
	SetQueryType(v string) *ListDatasetFileMetasShrinkRequest
	GetQueryType() *string
	SetQueryVideo(v string) *ListDatasetFileMetasShrinkRequest
	GetQueryVideo() *string
	SetScoreThreshold(v float32) *ListDatasetFileMetasShrinkRequest
	GetScoreThreshold() *float32
	SetSortBy(v string) *ListDatasetFileMetasShrinkRequest
	GetSortBy() *string
	SetStartFileUpdateTime(v string) *ListDatasetFileMetasShrinkRequest
	GetStartFileUpdateTime() *string
	SetStartTagUpdateTime(v string) *ListDatasetFileMetasShrinkRequest
	GetStartTagUpdateTime() *string
	SetStatus(v string) *ListDatasetFileMetasShrinkRequest
	GetStatus() *string
	SetThumbnailMode(v string) *ListDatasetFileMetasShrinkRequest
	GetThumbnailMode() *string
	SetTopK(v int32) *ListDatasetFileMetasShrinkRequest
	GetTopK() *int32
	SetWorkspaceId(v string) *ListDatasetFileMetasShrinkRequest
	GetWorkspaceId() *string
}

type ListDatasetFileMetasShrinkRequest struct {
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
	PageSize                         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryContentTypeIncludeAnyShrink *string `json:"QueryContentTypeIncludeAny,omitempty" xml:"QueryContentTypeIncludeAny,omitempty"`
	QueryExpression                  *string `json:"QueryExpression,omitempty" xml:"QueryExpression,omitempty"`
	// example:
	//
	// cars/20250221/
	QueryFileDir *string `json:"QueryFileDir,omitempty" xml:"QueryFileDir,omitempty"`
	// example:
	//
	// shuima
	QueryFileName                 *string `json:"QueryFileName,omitempty" xml:"QueryFileName,omitempty"`
	QueryFileTypeIncludeAnyShrink *string `json:"QueryFileTypeIncludeAny,omitempty" xml:"QueryFileTypeIncludeAny,omitempty"`
	// example:
	//
	// oss://test-xxx-oss/car/0001.png
	QueryImage                *string `json:"QueryImage,omitempty" xml:"QueryImage,omitempty"`
	QueryTagsExcludeShrink    *string `json:"QueryTagsExclude,omitempty" xml:"QueryTagsExclude,omitempty"`
	QueryTagsIncludeAllShrink *string `json:"QueryTagsIncludeAll,omitempty" xml:"QueryTagsIncludeAll,omitempty"`
	QueryTagsIncludeAnyShrink *string `json:"QueryTagsIncludeAny,omitempty" xml:"QueryTagsIncludeAny,omitempty"`
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
	QueryType  *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	QueryVideo *string `json:"QueryVideo,omitempty" xml:"QueryVideo,omitempty"`
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
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s ListDatasetFileMetasShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetFileMetasShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetFileMetasShrinkRequest) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *ListDatasetFileMetasShrinkRequest) GetEndFileUpdateTime() *string {
	return s.EndFileUpdateTime
}

func (s *ListDatasetFileMetasShrinkRequest) GetEndTagUpdateTime() *string {
	return s.EndTagUpdateTime
}

func (s *ListDatasetFileMetasShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDatasetFileMetasShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDatasetFileMetasShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDatasetFileMetasShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetFileMetasShrinkRequest) GetQueryContentTypeIncludeAnyShrink() *string {
	return s.QueryContentTypeIncludeAnyShrink
}

func (s *ListDatasetFileMetasShrinkRequest) GetQueryExpression() *string {
	return s.QueryExpression
}

func (s *ListDatasetFileMetasShrinkRequest) GetQueryFileDir() *string {
	return s.QueryFileDir
}

func (s *ListDatasetFileMetasShrinkRequest) GetQueryFileName() *string {
	return s.QueryFileName
}

func (s *ListDatasetFileMetasShrinkRequest) GetQueryFileTypeIncludeAnyShrink() *string {
	return s.QueryFileTypeIncludeAnyShrink
}

func (s *ListDatasetFileMetasShrinkRequest) GetQueryImage() *string {
	return s.QueryImage
}

func (s *ListDatasetFileMetasShrinkRequest) GetQueryTagsExcludeShrink() *string {
	return s.QueryTagsExcludeShrink
}

func (s *ListDatasetFileMetasShrinkRequest) GetQueryTagsIncludeAllShrink() *string {
	return s.QueryTagsIncludeAllShrink
}

func (s *ListDatasetFileMetasShrinkRequest) GetQueryTagsIncludeAnyShrink() *string {
	return s.QueryTagsIncludeAnyShrink
}

func (s *ListDatasetFileMetasShrinkRequest) GetQueryText() *string {
	return s.QueryText
}

func (s *ListDatasetFileMetasShrinkRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *ListDatasetFileMetasShrinkRequest) GetQueryVideo() *string {
	return s.QueryVideo
}

func (s *ListDatasetFileMetasShrinkRequest) GetScoreThreshold() *float32 {
	return s.ScoreThreshold
}

func (s *ListDatasetFileMetasShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDatasetFileMetasShrinkRequest) GetStartFileUpdateTime() *string {
	return s.StartFileUpdateTime
}

func (s *ListDatasetFileMetasShrinkRequest) GetStartTagUpdateTime() *string {
	return s.StartTagUpdateTime
}

func (s *ListDatasetFileMetasShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDatasetFileMetasShrinkRequest) GetThumbnailMode() *string {
	return s.ThumbnailMode
}

func (s *ListDatasetFileMetasShrinkRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *ListDatasetFileMetasShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDatasetFileMetasShrinkRequest) SetDatasetVersion(v string) *ListDatasetFileMetasShrinkRequest {
	s.DatasetVersion = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetEndFileUpdateTime(v string) *ListDatasetFileMetasShrinkRequest {
	s.EndFileUpdateTime = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetEndTagUpdateTime(v string) *ListDatasetFileMetasShrinkRequest {
	s.EndTagUpdateTime = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetMaxResults(v int32) *ListDatasetFileMetasShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetNextToken(v string) *ListDatasetFileMetasShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetOrder(v string) *ListDatasetFileMetasShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetPageSize(v int32) *ListDatasetFileMetasShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetQueryContentTypeIncludeAnyShrink(v string) *ListDatasetFileMetasShrinkRequest {
	s.QueryContentTypeIncludeAnyShrink = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetQueryExpression(v string) *ListDatasetFileMetasShrinkRequest {
	s.QueryExpression = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetQueryFileDir(v string) *ListDatasetFileMetasShrinkRequest {
	s.QueryFileDir = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetQueryFileName(v string) *ListDatasetFileMetasShrinkRequest {
	s.QueryFileName = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetQueryFileTypeIncludeAnyShrink(v string) *ListDatasetFileMetasShrinkRequest {
	s.QueryFileTypeIncludeAnyShrink = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetQueryImage(v string) *ListDatasetFileMetasShrinkRequest {
	s.QueryImage = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetQueryTagsExcludeShrink(v string) *ListDatasetFileMetasShrinkRequest {
	s.QueryTagsExcludeShrink = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetQueryTagsIncludeAllShrink(v string) *ListDatasetFileMetasShrinkRequest {
	s.QueryTagsIncludeAllShrink = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetQueryTagsIncludeAnyShrink(v string) *ListDatasetFileMetasShrinkRequest {
	s.QueryTagsIncludeAnyShrink = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetQueryText(v string) *ListDatasetFileMetasShrinkRequest {
	s.QueryText = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetQueryType(v string) *ListDatasetFileMetasShrinkRequest {
	s.QueryType = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetQueryVideo(v string) *ListDatasetFileMetasShrinkRequest {
	s.QueryVideo = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetScoreThreshold(v float32) *ListDatasetFileMetasShrinkRequest {
	s.ScoreThreshold = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetSortBy(v string) *ListDatasetFileMetasShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetStartFileUpdateTime(v string) *ListDatasetFileMetasShrinkRequest {
	s.StartFileUpdateTime = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetStartTagUpdateTime(v string) *ListDatasetFileMetasShrinkRequest {
	s.StartTagUpdateTime = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetStatus(v string) *ListDatasetFileMetasShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetThumbnailMode(v string) *ListDatasetFileMetasShrinkRequest {
	s.ThumbnailMode = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetTopK(v int32) *ListDatasetFileMetasShrinkRequest {
	s.TopK = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) SetWorkspaceId(v string) *ListDatasetFileMetasShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDatasetFileMetasShrinkRequest) Validate() error {
	return dara.Validate(s)
}
