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
	// A list of metadata IDs to query.
	DatasetFileMetaIdsShrink *string `json:"DatasetFileMetaIds,omitempty" xml:"DatasetFileMetaIds,omitempty"`
	// The version name of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The start time for the query that filters files by update time. The time must be a UTC timestamp in ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01.000Z
	EndFileUpdateTime *string `json:"EndFileUpdateTime,omitempty" xml:"EndFileUpdateTime,omitempty"`
	// The start time for querying tags by their last update time. The time must be in UTC and in the ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01.000Z
	EndTagUpdateTime *string `json:"EndTagUpdateTime,omitempty" xml:"EndTagUpdateTime,omitempty"`
	// The end of the time range for a query that filters tags by their last update time. The time is a UTC timestamp in ISO 8601 format.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// > If you do not specify this parameter, the first page of results is returned. If a value is returned for this parameter, more results are available. To get the next page, use the returned token in your next request. Repeat this process until no token is returned, which indicates that all results have been retrieved.
	//
	// example:
	//
	// 90a6ee35-****-4cd4-927e-1f45e1cb8b62_1729644433000
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order for the specified field in a paginated query. Use this parameter with \\`SortBy\\`. The default value is \\`DESC\\`. Valid values:
	//
	// - ASC: Ascending.
	//
	// - DESC: Descending.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// Deprecated
	//
	// The number of entries per page. If you also specify \\`MaxResults\\`, the value of \\`MaxResults\\` takes precedence.
	//
	// > This parameter is deprecated. Use \\`NextToken\\` and \\`MaxResults\\` for paginated queries.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A search condition to include any of the specified content types. The search results must match at least one of these types. You can specify multiple content types. If this parameter is empty, this condition is not applied. Use commas to separate multiple types in the array.
	QueryContentTypeIncludeAnyShrink *string `json:"QueryContentTypeIncludeAny,omitempty" xml:"QueryContentTypeIncludeAny,omitempty"`
	// The maximum number of results to return per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// (FileUpdateTime > \\"2025-02-28T00:00:00Z\\" AND FileUpdateTime < \\"2025-05-30T09:27:29Z\\") AND FileDir:\\"blue_car\\" AND NOT FileName="toyota.jpg" AND (( Tags.all=\\"lane line\\" AND Tags.all=\\"barrier gate\\") OR NOT Tags.user=\\"rainy days\\" ) AND HAS SemanticIndexJobId AND Content:\\"a fallen water horse\\" AND TopK=100 AND SignMode=\\"PUBLIC\\"
	QueryExpression *string `json:"QueryExpression,omitempty" xml:"QueryExpression,omitempty"`
	// The name of the file to retrieve. This parameter supports fuzzy search.
	//
	// example:
	//
	// cars/20250221/
	QueryFileDir *string `json:"QueryFileDir,omitempty" xml:"QueryFileDir,omitempty"`
	// The tags to exclude from the query results. If you do not specify any tags, this filter is not applied.
	//
	// > This parameter is valid only when QueryType is set to TAG or MIX.
	//
	// example:
	//
	// car
	QueryFileName *string `json:"QueryFileName,omitempty" xml:"QueryFileName,omitempty"`
	// The search keyword for the file directory. Fuzzy search is supported.
	QueryFileTypeIncludeAnyShrink *string `json:"QueryFileTypeIncludeAny,omitempty" xml:"QueryFileTypeIncludeAny,omitempty"`
	// The image information to use for an image-based search.
	//
	// - Specify the public URL of an image in an OSS bucket. The format is \\`oss\\://{bucket_name}/{object_path}\\`. \\`bucket_name\\` is the name of the bucket, and \\`object_path\\` is the path of the file in the bucket.
	//
	// > This parameter is valid only when \\`QueryType\\` is set to \\`VECTOR\\` or \\`MIX\\`.
	//
	// example:
	//
	// oss://test-xxx-oss/car/0001.png
	QueryImage *string `json:"QueryImage,omitempty" xml:"QueryImage,omitempty"`
	// A comma-separated list of tags. The query returns files that match at least one of the specified tags. If you do not specify this parameter, this filter is ignored.
	//
	// > This parameter is valid only when QueryType is set to TAG or MIX.
	QueryTagsExcludeShrink *string `json:"QueryTagsExclude,omitempty" xml:"QueryTagsExclude,omitempty"`
	// The metadata IDs to query.
	QueryTagsIncludeAllShrink *string `json:"QueryTagsIncludeAll,omitempty" xml:"QueryTagsIncludeAll,omitempty"`
	// A condition that retrieves items that have all of the specified tags. The tags are specified as a comma-separated array. This condition is not applied if the parameter is empty.
	//
	// > This parameter takes effect only when QueryType is set to TAG or MIX. If QueryType is set to TAG, the value of QueryText is also added to this condition.
	QueryTagsIncludeAnyShrink *string `json:"QueryTagsIncludeAny,omitempty" xml:"QueryTagsIncludeAny,omitempty"`
	// The text to search for.
	//
	// example:
	//
	// A fallen water
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// The search type. Valid values:
	//
	// - MIX: Mixed search. This is the default value.
	//
	// - TAG: Searches by tag only.
	//
	// - VECTOR: Searches by vector only.
	//
	// example:
	//
	// MIX
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The status of the metadata to query.
	//
	// - ACTIVE: Returns metadata for active files. This is the default value.
	//
	// - ALL: Returns metadata for all files.
	//
	// - DELETED: Returns metadata for logically deleted files.
	//
	// example:
	//
	// oss://test-xxx-oss/car/0001.mp4
	QueryVideo *string `json:"QueryVideo,omitempty" xml:"QueryVideo,omitempty"`
	// The similarity score threshold. Only results with a score greater than this threshold are returned.
	//
	// > This parameter is valid only when \\`QueryType\\` is set to \\`VECTOR\\` or \\`MIX\\`.
	//
	// example:
	//
	// 0.6
	ScoreThreshold *float32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The field to sort by for paginated queries. If you do not specify this parameter, results are sorted by relevance from high to low. Other valid values are as follows:
	//
	// - FileCreateTime: Sort by file creation time.
	//
	// - FileUpdateTime: Sort by file last modified time.
	//
	// example:
	//
	// FileCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The end of the time range for a query based on file update time. The value is a UTC timestamp in ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01.000Z
	StartFileUpdateTime *string `json:"StartFileUpdateTime,omitempty" xml:"StartFileUpdateTime,omitempty"`
	// The file content types. The query returns files that match any of the specified types. You can specify multiple types and separate them with commas. If this parameter is empty, this filter is ignored.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-01-12T14:36:01.000Z
	StartTagUpdateTime *string `json:"StartTagUpdateTime,omitempty" xml:"StartTagUpdateTime,omitempty"`
	// A query statement, also known as a Domain-Specific Language (DSL) query, lets you express complex retrieval conditions. It supports grouping, Boolean logic (AND/OR/NOT), range comparisons (>, >=, <, <=), property existence (HAS/NOT HAS), tokenized matches (:), and exact matches (=). Use DSL for advanced retrieval scenarios.
	//
	// 	Notice: To avoid conflicts, do not use this query statement with other query parameters.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The mode for generating image thumbnails. Thumbnails are supported only for files in OSS.
	//
	// - Proportional scaling: \\`p_{percentage}\\`. The \\`percentage\\` parameter specifies the scaling ratio. Valid values: 1 to 100. For example, \\`p_50\\` scales the image to 50% of its original size.
	//
	// - Fixed width, adaptive height: \\`w_{width}\\`. The \\`width\\` parameter specifies the image width. Valid values: 1 to 16,384. For example, \\`w_200\\` sets the image width to 200 pixels and scales the height adaptively.
	//
	// - Fixed height, adaptive width: \\`h_{height}\\`. The \\`height\\` parameter specifies the image height. Valid values: 1 to 16,384. For example, \\`h_100\\` sets the image height to 100 pixels and scales the width adaptively.
	//
	// - Fixed width and height with padding: \\`m_pad,w_{width},h_{height},color_{RGB}\\`. The \\`m_pad\\` parameter scales the image to the maximum size that fits within a rectangle of the specified width and height. The \\`RGB\\` parameter specifies the color for the centered padding in the empty areas. If you do not specify this parameter, the empty areas are filled with white by default. The \\`width\\` and \\`height\\` parameters specify the image width and height. The values for both \\`width\\` and \\`height\\` must be between 1 and 16,384.
	//
	// - Fixed width and height with center crop: \\`m_fill,w_{width},h_{height}\\`. The \\`m_fill\\` parameter proportionally scales the image to the minimum size that covers the specified width and height, and then crops the excess from the center. The \\`width\\` and \\`height\\` parameters specify the image width and height. The values for both \\`width\\` and \\`height\\` must be between 1 and 16,384. For example, \\`m_fill,w_100,h_100\\` scales and crops the image to 100 × 100 pixels from the center.
	//
	// - Forced width and height scaling: \\`m_fixed,w_{width},h_{height}\\`. The \\`width\\` and \\`height\\` parameters specify the image width and height. The values for both \\`width\\` and \\`height\\` must be between 1 and 16,384. For example, \\`m_fixed,w_100,h_100\\` forces the image to be scaled to 100 × 100 pixels.
	//
	// example:
	//
	// w_100
	ThumbnailMode *string `json:"ThumbnailMode,omitempty" xml:"ThumbnailMode,omitempty"`
	// The maximum number of search results to return.
	//
	// > This parameter is valid only when \\`QueryType\\` is set to \\`VECTOR\\` or \\`MIX\\`.
	//
	// example:
	//
	// 100
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// The ID of the workspace where the dataset is located. For more information, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
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
