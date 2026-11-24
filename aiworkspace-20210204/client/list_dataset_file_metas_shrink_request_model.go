// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetFileMetasShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetFileMetaIdsShrink(v string) *ListDatasetFileMetasShrinkRequest
	GetDatasetFileMetaIdsShrink() *string
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
	// The list of metadata IDs to query.
	DatasetFileMetaIdsShrink *string `json:"DatasetFileMetaIds,omitempty" xml:"DatasetFileMetaIds,omitempty"`
	// The dataset version name.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The end time of the file update time range for the query. Specify the time in the ISO 8601 standard in UTC.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01.000Z
	EndFileUpdateTime *string `json:"EndFileUpdateTime,omitempty" xml:"EndFileUpdateTime,omitempty"`
	// The end time of the tag last update time range for the query. Specify the time in the ISO 8601 standard in UTC.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01.000Z
	EndTagUpdateTime *string `json:"EndTagUpdateTime,omitempty" xml:"EndTagUpdateTime,omitempty"`
	// The maximum number of results to return per query when NextToken is used. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination cursor.
	//
	// >
	//
	// > If this parameter is not specified, the first page of data is returned. If this parameter has a return value, it indicates that there is a next page. You can use the returned NextToken as a parameter to request the next page of data until Null is returned, which indicates that all data is retrieved.
	//
	// example:
	//
	// 90a6ee35-****-4cd4-927e-1f45e1cb8b62_1729644433000
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ascending or descending order for the specified sorting field in a paged query. This parameter is used together with SortBy during paging. Default value: DESC. Valid values:
	//
	// - ASC: Ascending order.
	//
	// - DESC: Descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// Deprecated
	//
	// The number of entries per page. If MaxResults is also specified, MaxResults prevails.
	//
	// > This parameter will soon be offline. We recommend that you use NextToken and MaxResults for paging in paged query operations.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search condition for "including any of the following file content types". You can select multiple file content types, and the search results only need to meet one of them. If this parameter is empty, this condition is not applied. Separate multiple file content types with commas.
	QueryContentTypeIncludeAnyShrink *string `json:"QueryContentTypeIncludeAny,omitempty" xml:"QueryContentTypeIncludeAny,omitempty"`
	// The query statement (DSL) is a domain-specific language used to express complex search conditions. It supports grouping, Boolean logic (AND/OR/NOT), range comparison (>, >=, <, <=), property existence (HAS/NOT HAS), token matching (:), and exact matching (=), and is suitable for advanced search scenarios.
	//
	// It is generally used for complex advanced search conditions.
	//
	// <notice>To avoid conflicts, do not mix this query statement with other query parameters.</notice>
	//
	// example:
	//
	// (FileUpdateTime > \\"2025-02-28T00:00:00Z\\" AND FileUpdateTime < \\"2025-05-30T09:27:29Z\\") AND FileDir:\\"blue_car\\" AND NOT FileName="toyota.jpg" AND (( Tags.all=\\"lane line\\" AND Tags.all=\\"barrier gate\\") OR NOT Tags.user=\\"rainy days\\" ) AND HAS SemanticIndexJobId AND Content:\\"a fallen water horse\\" AND TopK=100 AND SignMode=\\"PUBLIC\\"
	QueryExpression *string `json:"QueryExpression,omitempty" xml:"QueryExpression,omitempty"`
	// The file directory search condition. Fuzzy search is supported.
	//
	// example:
	//
	// cars/20250221/
	QueryFileDir *string `json:"QueryFileDir,omitempty" xml:"QueryFileDir,omitempty"`
	// The file name search condition. Fuzzy search is supported.
	//
	// example:
	//
	// car
	QueryFileName *string `json:"QueryFileName,omitempty" xml:"QueryFileName,omitempty"`
	// The search condition for "including any of the following file types". You can select multiple file types, and the search results only need to meet one of them. If this parameter is empty, this condition is not applied. Separate multiple file types with commas.
	QueryFileTypeIncludeAnyShrink *string `json:"QueryFileTypeIncludeAny,omitempty" xml:"QueryFileTypeIncludeAny,omitempty"`
	// The image information for image-to-image search.
	//
	// 	- Supports passing an OSS URL that allows public network access for the image. Format: oss://{bucket_name}/{object_path}. bucket_name is the bucket name, and object_path is the path of the file in the bucket.
	//
	// > This parameter takes effect only when QueryType is set to VECTOR or MIX.
	//
	// example:
	//
	// oss://test-xxx-oss/car/0001.png
	QueryImage *string `json:"QueryImage,omitempty" xml:"QueryImage,omitempty"`
	// The search condition for "excluding the following tags". You can select multiple tags, and the search results must not contain these tags. If this parameter is empty, this condition is not applied.
	//
	// > This parameter takes effect only when QueryType is set to TAG or MIX.
	QueryTagsExcludeShrink *string `json:"QueryTagsExclude,omitempty" xml:"QueryTagsExclude,omitempty"`
	// The search condition for "including all of the following tags". You can select multiple tags, and the search results must meet all these tags simultaneously. If this parameter is empty, this condition is not applied. Separate multiple tags with commas.
	//
	// > This parameter takes effect only when QueryType is set to TAG or MIX. When QueryType is set to TAG, QueryText is added to this condition.
	QueryTagsIncludeAllShrink *string `json:"QueryTagsIncludeAll,omitempty" xml:"QueryTagsIncludeAll,omitempty"`
	// The search condition for "including any of the following tags". You can select multiple tags, and the search results only need to meet one of them. If this parameter is empty, this condition is not applied. Separate multiple tags with commas.
	//
	// > This parameter takes effect only when QueryType is set to TAG or MIX.
	QueryTagsIncludeAnyShrink *string `json:"QueryTagsIncludeAny,omitempty" xml:"QueryTagsIncludeAny,omitempty"`
	// The text content to search.
	//
	// example:
	//
	// A fallen water
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// The retrieve type. Valid values:
	//
	// 	- MIX: Hybrid retrieve. This is the default value.
	//
	// 	- TAG: Label-only retrieve.
	//
	// 	- VECTOR: AISearch-only retrieve.
	//
	// example:
	//
	// MIX
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The video file information for video search.
	//
	// 	- Supports passing an OSS URL that allows public network access for the video. Format: oss://{bucket_name}/{object_path}. bucket_name is the bucket name, and object_path is the path of the file in the bucket.
	//
	// > This parameter takes effect only when QueryType is set to VECTOR or MIX.
	//
	// example:
	//
	// oss://test-xxx-oss/car/0001.mp4
	QueryVideo *string `json:"QueryVideo,omitempty" xml:"QueryVideo,omitempty"`
	// The similarity score threshold for filtering. Only results with a score greater than ScoreThreshold are returned.
	//
	// > This parameter takes effect only when QueryType is set to VECTOR or MIX.
	//
	// example:
	//
	// 0.6
	ScoreThreshold *float32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The sorting field for paged query. If this parameter is not specified, the results are sorted by retrieve relevance in descending order during paging. Valid values:
	//
	// 	- FileCreateTime: Sort by file creation time.
	//
	// 	- FileUpdateTime: Sort by file last modification time.
	//
	// example:
	//
	// FileCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The start time of the file update time range for the query. Specify the time in the ISO 8601 standard in UTC.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01.000Z
	StartFileUpdateTime *string `json:"StartFileUpdateTime,omitempty" xml:"StartFileUpdateTime,omitempty"`
	// The start time of the tag last update time range for the query. Specify the time in the ISO 8601 standard in UTC.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01.000Z
	StartTagUpdateTime *string `json:"StartTagUpdateTime,omitempty" xml:"StartTagUpdateTime,omitempty"`
	// The metadata status to query. Valid values:
	//
	// 	- ACTIVE: Query only undeleted data. This is the default value.
	//
	// 	- ALL: Query all data.
	//
	// 	- DELETED: Query only logically deleted data.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The thumbnail mode for images. Currently, only OSS files support thumbnails. Valid values:
	//
	// - Proportional scaling: p_{percentage}. percentage indicates the expected scaling ratio. Valid values: [1, 100]. Example: p_50 indicates that the thumbnail size is 50% of the original file size.
	//
	// - Fixed width and adaptive height: w_{width}. width indicates the expected image width. Valid values: [1, 16384]. Example: w_200 indicates that the image width is fixed at 200 pixels and the height is adaptively scaled.
	//
	// - Fixed height and adaptive width: h_{height}. height indicates the expected image height. Valid values: [1, 16384]. Example: h_100 indicates that the image height is fixed at 100 pixels and the width is adaptively scaled.
	//
	// - Fixed width and height with padding: m_pad,w_{width},h_{height},color_{RGB}. m_pad indicates that the image is scaled to the maximum size within the specified w and h rectangle. RGB specifies the color used to fill the blank space in the center. If this parameter is not specified, white is used by default. width indicates the expected image width, and height indicates the expected image height. Valid values of width and height: [1, 16384]. Example:
	//
	// - Fixed width and height with center cropping: m_fill,w_{width},h_{height}. m_fill indicates that the image is proportionally scaled to the minimum size that extends beyond the specified w and h rectangle, and the excess part is center-cropped. width indicates the expected image width, and height indicates the expected image height. Valid values of width and height: [1, 16384]. Example: m_fill,w_100,h_100 indicates that the image width and height are both fixed at 100 pixels for center cropping and scaling.
	//
	// - Forced width and height scaling: m_fixed,w_{width},h_{height}. width indicates the expected image width, and height indicates the expected image height. Valid values of width and height: [1, 16384]. Example: m_fixed,w_100,h_100 indicates that the image width and height are both fixed at 100 pixels for forced scaling.
	//
	// example:
	//
	// w_100
	ThumbnailMode *string `json:"ThumbnailMode,omitempty" xml:"ThumbnailMode,omitempty"`
	// The number of search results to return. A maximum of TopK search results are returned.
	//
	// > This parameter takes effect only when QueryType is set to VECTOR or MIX.
	//
	// example:
	//
	// 100
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// The ID of the workspace to which the dataset belongs. For more information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
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

func (s *ListDatasetFileMetasShrinkRequest) GetDatasetFileMetaIdsShrink() *string {
	return s.DatasetFileMetaIdsShrink
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

func (s *ListDatasetFileMetasShrinkRequest) SetDatasetFileMetaIdsShrink(v string) *ListDatasetFileMetasShrinkRequest {
	s.DatasetFileMetaIdsShrink = &v
	return s
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
