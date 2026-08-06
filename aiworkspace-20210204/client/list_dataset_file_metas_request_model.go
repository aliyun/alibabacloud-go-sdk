// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetFileMetasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetFileMetaIds(v []*string) *ListDatasetFileMetasRequest
	GetDatasetFileMetaIds() []*string
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
	SetQueryVideo(v string) *ListDatasetFileMetasRequest
	GetQueryVideo() *string
	SetScoreThreshold(v float32) *ListDatasetFileMetasRequest
	GetScoreThreshold() *float32
	SetSortBy(v string) *ListDatasetFileMetasRequest
	GetSortBy() *string
	SetStartFileUpdateTime(v string) *ListDatasetFileMetasRequest
	GetStartFileUpdateTime() *string
	SetStartTagUpdateTime(v string) *ListDatasetFileMetasRequest
	GetStartTagUpdateTime() *string
	SetStatus(v string) *ListDatasetFileMetasRequest
	GetStatus() *string
	SetThumbnailMode(v string) *ListDatasetFileMetasRequest
	GetThumbnailMode() *string
	SetTopK(v int32) *ListDatasetFileMetasRequest
	GetTopK() *int32
	SetWorkspaceId(v string) *ListDatasetFileMetasRequest
	GetWorkspaceId() *string
}

type ListDatasetFileMetasRequest struct {
	// The list of metadata IDs to query.
	DatasetFileMetaIds []*string `json:"DatasetFileMetaIds,omitempty" xml:"DatasetFileMetaIds,omitempty" type:"Repeated"`
	// The dataset version name.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The end time for the file update time query range. The value is a UTC timestamp in ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01.000Z
	EndFileUpdateTime *string `json:"EndFileUpdateTime,omitempty" xml:"EndFileUpdateTime,omitempty"`
	// The end time for the tag last update time query range. The value is a UTC timestamp in ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01.000Z
	EndTagUpdateTime *string `json:"EndTagUpdateTime,omitempty" xml:"EndTagUpdateTime,omitempty"`
	// The maximum number of results to return per request when using NextToken-based pagination. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// >
	//
	// > If this parameter is not specified, the first page of data is returned. If a value is returned for this parameter, more pages are available. Pass the returned NextToken value as a request parameter to retrieve the next page, until no NextToken value is returned, which indicates that all data has been retrieved.
	//
	// example:
	//
	// 90a6ee35-****-4cd4-927e-1f45e1cb8b62_1729644433000
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sorting order for the specified sort field in paging queries. Used together with SortBy. Default value: DESC. Valid values:
	//
	// - ASC: ascending order.
	//
	// - DESC: descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// Deprecated
	//
	// The number of entries per page. If MaxResults is also specified, MaxResults takes precedence.
	//
	// > This parameter will be offline soon. Use NextToken and MaxResults to perform paging operations.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search condition for "include any of the following content types". You can select multiple content types, and the query results need to match only one of them. If empty, this condition is not applied. Array values are separated by commas.
	QueryContentTypeIncludeAny []*string `json:"QueryContentTypeIncludeAny,omitempty" xml:"QueryContentTypeIncludeAny,omitempty" type:"Repeated"`
	// The query statement (DSL) is a domain-specific language for expressing complex retrieve conditions. It supports grouping, Boolean logic (AND/OR/NOT), range comparisons (>, >=, <, <=), property existence (HAS/NOT HAS), tokenized matching (:), and exact match (=), suitable for advanced retrieve scenarios.
	//
	// Generally used for complex advanced conditional retrieve operations.
	//
	// <notice>To avoid conflicts, after setting this query statement, do not use it together with other query parameters.</notice>
	//
	// example:
	//
	// (FileUpdateTime > \\"2025-02-28T00:00:00Z\\" AND FileUpdateTime < \\"2025-05-30T09:27:29Z\\") AND FileDir:\\"blue_car\\" AND NOT FileName="toyota.jpg" AND (( Tags.all=\\"lane line\\" AND Tags.all=\\"barrier gate\\") OR NOT Tags.user=\\"rainy days\\" ) AND HAS SemanticIndexJobId AND Content:\\"a fallen water horse\\" AND TopK=100 AND SignMode=\\"PUBLIC\\"
	QueryExpression *string `json:"QueryExpression,omitempty" xml:"QueryExpression,omitempty"`
	// The file directory search condition. Fuzzy match is supported.
	//
	// example:
	//
	// cars/20250221/
	QueryFileDir *string `json:"QueryFileDir,omitempty" xml:"QueryFileDir,omitempty"`
	// The file name search condition. Fuzzy match is supported.
	//
	// example:
	//
	// car
	QueryFileName *string `json:"QueryFileName,omitempty" xml:"QueryFileName,omitempty"`
	// The search condition for "include any of the following file types". You can select multiple file types, and the query results need to match only one of them. If empty, this condition is not applied. Array values are separated by commas.
	QueryFileTypeIncludeAny []*string `json:"QueryFileTypeIncludeAny,omitempty" xml:"QueryFileTypeIncludeAny,omitempty" type:"Repeated"`
	// The image information for image-to-image search.
	//
	// 	- Supports a public network access OSS URL in the format: oss://{bucket_name}/{object_path}, where bucket_name is the bucket name and object_path is the file path in the bucket.
	//
	// > This parameter takes effect only when QueryType is set to VECTOR or MIX.
	//
	// example:
	//
	// oss://test-xxx-oss/car/0001.png
	QueryImage *string `json:"QueryImage,omitempty" xml:"QueryImage,omitempty"`
	// The search condition for "exclude the following tags". You can select multiple tags, and the query results must not contain any of them. If empty, this condition is not applied.
	//
	// > This parameter takes effect only when QueryType is set to TAG or MIX.
	QueryTagsExclude []*string `json:"QueryTagsExclude,omitempty" xml:"QueryTagsExclude,omitempty" type:"Repeated"`
	// The search condition for "include all of the following tags". You can select multiple tags, and the query results must match all of them. If empty, this condition is not applied. Array values are separated by commas.
	//
	// > This parameter takes effect only when QueryType is set to TAG or MIX. When QueryType is set to TAG, QueryText is added to this condition.
	QueryTagsIncludeAll []*string `json:"QueryTagsIncludeAll,omitempty" xml:"QueryTagsIncludeAll,omitempty" type:"Repeated"`
	// The search condition for "include any of the following tags". You can select multiple tags, and the query results need to match only one of them. If empty, this condition is not applied. Array values are separated by commas.
	//
	// > This parameter takes effect only when QueryType is set to TAG or MIX.
	QueryTagsIncludeAny []*string `json:"QueryTagsIncludeAny,omitempty" xml:"QueryTagsIncludeAny,omitempty" type:"Repeated"`
	// The text content to search for.
	//
	// example:
	//
	// A fallen water
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// The retrieve type. Valid values:
	//
	// 	- MIX: hybrid retrieve (default).
	//
	// 	- TAG: label-only retrieve.
	//
	// 	- VECTOR: vector retrieve only.
	//
	// example:
	//
	// MIX
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The video file information for video-based search.
	//
	// 	- Supports a public network access OSS URL in the format: oss://{bucket_name}/{object_path}, where bucket_name is the bucket name and object_path is the file path in the bucket.
	//
	// > This parameter takes effect only when QueryType is set to VECTOR or MIX.
	//
	// example:
	//
	// oss://test-xxx-oss/car/0001.mp4
	QueryVideo *string `json:"QueryVideo,omitempty" xml:"QueryVideo,omitempty"`
	// The similarity score threshold. Only results with a score greater than ScoreThreshold are returned.
	//
	// > This parameter takes effect only when QueryType is set to VECTOR or MIX.
	//
	// example:
	//
	// 0.6
	ScoreThreshold *float32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The sorting field for paging queries. By default, results are sorted by retrieve relevance in descending order. Valid values:
	//
	// 	- FileCreateTime: sorting by file creation time.
	//
	// 	- FileUpdateTime: sorting by file last modification time.
	//
	// example:
	//
	// FileCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The start time for the file update time query range. The value is a UTC timestamp in ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01.000Z
	StartFileUpdateTime *string `json:"StartFileUpdateTime,omitempty" xml:"StartFileUpdateTime,omitempty"`
	// The start time for the tag last update time query range. The value is a UTC timestamp in ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01.000Z
	StartTagUpdateTime *string `json:"StartTagUpdateTime,omitempty" xml:"StartTagUpdateTime,omitempty"`
	// The metadata status to query. Valid values:
	//
	// 	- ACTIVE: queries only non-deleted data (default).
	//
	// 	- ALL: queries all data.
	//
	// 	- DELETED: queries only logically deleted data.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The thumbnail mode for images. Currently, only OSS files support thumbnails:
	//
	// - Proportional scaling: p_{percentage}, where percentage specifies the desired scaling ratio. Valid values: [1, 100]. Example: p_50 uses 50% of the original file size as the thumbnail.
	//
	// - Fixed width with adaptive height: w_{width}, where width specifies the desired image width. Valid values: [1, 16384]. Example: w_200 fixes the image width to 200 pixels and adaptively scales the height.
	//
	// - Fixed height with adaptive width: h_{height}, where height specifies the desired image height. Valid values: [1, 16384]. Example: h_100 fixes the image height to 100 pixels and adaptively scales the width.
	//
	// - Fixed dimensions with padding: m_pad,w_{width},h_{height},color_{RGB}. m_pad scales the image to the largest size that fits within the specified width and height rectangle. RGB specifies the fill color for blank areas. If not specified, white is used by default. width specifies the desired image width and height specifies the desired image height. Valid values for both width and height: [1, 16384].
	//
	// - Fixed dimensions with center cropping: m_fill,w_{width},h_{height}. m_fill proportionally scales the image to the smallest size that extends beyond the specified width and height rectangle, and center-crops the excess. width specifies the desired image width and height specifies the desired image height. Valid values for both width and height: [1, 16384]. Example: m_fill,w_100,h_100 fixes both width and height to 100 pixels with center cropping.
	//
	// - Forced dimensions: m_fixed,w_{width},h_{height}. width specifies the desired image width and height specifies the desired image height. Valid values for both width and height: [1, 16384]. Example: m_fixed,w_100,h_100 forces both width and height to 100 pixels.
	//
	// example:
	//
	// w_100
	ThumbnailMode *string `json:"ThumbnailMode,omitempty" xml:"ThumbnailMode,omitempty"`
	// The maximum number of results to return. Only the top K results are returned.
	//
	// > This parameter takes effect only when QueryType is set to VECTOR or MIX.
	//
	// example:
	//
	// 100
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// The workspace ID where the dataset resides. For information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
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

func (s *ListDatasetFileMetasRequest) GetDatasetFileMetaIds() []*string {
	return s.DatasetFileMetaIds
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

func (s *ListDatasetFileMetasRequest) GetQueryVideo() *string {
	return s.QueryVideo
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

func (s *ListDatasetFileMetasRequest) GetStatus() *string {
	return s.Status
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

func (s *ListDatasetFileMetasRequest) SetDatasetFileMetaIds(v []*string) *ListDatasetFileMetasRequest {
	s.DatasetFileMetaIds = v
	return s
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

func (s *ListDatasetFileMetasRequest) SetQueryVideo(v string) *ListDatasetFileMetasRequest {
	s.QueryVideo = &v
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

func (s *ListDatasetFileMetasRequest) SetStatus(v string) *ListDatasetFileMetasRequest {
	s.Status = &v
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
