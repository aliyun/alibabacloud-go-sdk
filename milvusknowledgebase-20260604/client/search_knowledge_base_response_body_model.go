// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SearchKnowledgeBaseResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *SearchKnowledgeBaseResponseBody
	GetCode() *int32
	SetHttpStatusCode(v int32) *SearchKnowledgeBaseResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SearchKnowledgeBaseResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *SearchKnowledgeBaseResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchKnowledgeBaseResponseBody
	GetPageSize() *int32
	SetQueryLabels(v []*string) *SearchKnowledgeBaseResponseBody
	GetQueryLabels() []*string
	SetRequestId(v string) *SearchKnowledgeBaseResponseBody
	GetRequestId() *string
	SetResults(v []*SearchKnowledgeBaseResponseBodyResults) *SearchKnowledgeBaseResponseBody
	GetResults() []*SearchKnowledgeBaseResponseBodyResults
	SetSuccess(v bool) *SearchKnowledgeBaseResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *SearchKnowledgeBaseResponseBody
	GetTotalCount() *int64
}

type SearchKnowledgeBaseResponseBody struct {
	// The details of the permission verification failure.
	//
	// example:
	//
	// {"PolicyType":"AccountLevelIdentityBasedPolicy","NoPermissionType":"ImplicitDeny","AuthAction":"milvusknowledgebase:SearchKnowledgeBase"}
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 0
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The query labels.
	QueryLabels []*string `json:"queryLabels,omitempty" xml:"queryLabels,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 019FCC42-90DE-56D3-A10D-3C06995DED17
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of retrieval results.
	Results []*SearchKnowledgeBaseResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of results.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SearchKnowledgeBaseResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SearchKnowledgeBaseResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SearchKnowledgeBaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchKnowledgeBaseResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchKnowledgeBaseResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchKnowledgeBaseResponseBody) GetQueryLabels() []*string {
	return s.QueryLabels
}

func (s *SearchKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchKnowledgeBaseResponseBody) GetResults() []*SearchKnowledgeBaseResponseBodyResults {
	return s.Results
}

func (s *SearchKnowledgeBaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchKnowledgeBaseResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchKnowledgeBaseResponseBody) SetAccessDeniedDetail(v string) *SearchKnowledgeBaseResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetCode(v int32) *SearchKnowledgeBaseResponseBody {
	s.Code = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetHttpStatusCode(v int32) *SearchKnowledgeBaseResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetMessage(v string) *SearchKnowledgeBaseResponseBody {
	s.Message = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetPageNumber(v int32) *SearchKnowledgeBaseResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetPageSize(v int32) *SearchKnowledgeBaseResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetQueryLabels(v []*string) *SearchKnowledgeBaseResponseBody {
	s.QueryLabels = v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetRequestId(v string) *SearchKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetResults(v []*SearchKnowledgeBaseResponseBodyResults) *SearchKnowledgeBaseResponseBody {
	s.Results = v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetSuccess(v bool) *SearchKnowledgeBaseResponseBody {
	s.Success = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetTotalCount(v int64) *SearchKnowledgeBaseResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchKnowledgeBaseResponseBodyResults struct {
	// The chunk ID.
	//
	// example:
	//
	// 7f0de3e041322a1d
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// The chunk content.
	//
	// example:
	//
	// Tax amount 1.59
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The content type.
	//
	// example:
	//
	// table
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The document ID.
	//
	// example:
	//
	// 539ddb688fe811f182f637422a0713b5
	DocumentId *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
	// The document name.
	//
	// example:
	//
	// invoice.pdf
	DocumentName *string `json:"documentName,omitempty" xml:"documentName,omitempty"`
	// The list of associated images.
	Images []*SearchKnowledgeBaseResponseBodyResultsImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// The knowledge base ID.
	//
	// example:
	//
	// kd-xxxxxxxxxx
	KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty" xml:"knowledgeBaseId,omitempty"`
	// The list of document locations.
	Locations []*SearchKnowledgeBaseResponseBodyResultsLocations `json:"locations,omitempty" xml:"locations,omitempty" type:"Repeated"`
	// The parent chunk ID.
	//
	// example:
	//
	// parent-chunk-id
	ParentChunkId *string `json:"parentChunkId,omitempty" xml:"parentChunkId,omitempty"`
	// The scalar columns of the structured knowledge base. The columns are returned by their original column names and are not used in retrieval.
	//
	// example:
	//
	// {"question":"How do I reset it?","category":"account"}
	ScalarFields interface{} `json:"scalarFields,omitempty" xml:"scalarFields,omitempty"`
	// The overall relevance score.
	//
	// example:
	//
	// 0.26136884
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// The relevance score details.
	ScoreDetails *SearchKnowledgeBaseResponseBodyResultsScoreDetails `json:"scoreDetails,omitempty" xml:"scoreDetails,omitempty" type:"Struct"`
	// The list of tags.
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s SearchKnowledgeBaseResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseResponseBodyResults) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetChunkId() *string {
	return s.ChunkId
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetContent() *string {
	return s.Content
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetContentType() *string {
	return s.ContentType
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetDocumentId() *string {
	return s.DocumentId
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetDocumentName() *string {
	return s.DocumentName
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetImages() []*SearchKnowledgeBaseResponseBodyResultsImages {
	return s.Images
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetLocations() []*SearchKnowledgeBaseResponseBodyResultsLocations {
	return s.Locations
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetParentChunkId() *string {
	return s.ParentChunkId
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetScalarFields() interface{} {
	return s.ScalarFields
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetScore() *float32 {
	return s.Score
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetScoreDetails() *SearchKnowledgeBaseResponseBodyResultsScoreDetails {
	return s.ScoreDetails
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetTags() []*string {
	return s.Tags
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetChunkId(v string) *SearchKnowledgeBaseResponseBodyResults {
	s.ChunkId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetContent(v string) *SearchKnowledgeBaseResponseBodyResults {
	s.Content = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetContentType(v string) *SearchKnowledgeBaseResponseBodyResults {
	s.ContentType = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetDocumentId(v string) *SearchKnowledgeBaseResponseBodyResults {
	s.DocumentId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetDocumentName(v string) *SearchKnowledgeBaseResponseBodyResults {
	s.DocumentName = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetImages(v []*SearchKnowledgeBaseResponseBodyResultsImages) *SearchKnowledgeBaseResponseBodyResults {
	s.Images = v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetKnowledgeBaseId(v string) *SearchKnowledgeBaseResponseBodyResults {
	s.KnowledgeBaseId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetLocations(v []*SearchKnowledgeBaseResponseBodyResultsLocations) *SearchKnowledgeBaseResponseBodyResults {
	s.Locations = v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetParentChunkId(v string) *SearchKnowledgeBaseResponseBodyResults {
	s.ParentChunkId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetScalarFields(v interface{}) *SearchKnowledgeBaseResponseBodyResults {
	s.ScalarFields = v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetScore(v float32) *SearchKnowledgeBaseResponseBodyResults {
	s.Score = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetScoreDetails(v *SearchKnowledgeBaseResponseBodyResultsScoreDetails) *SearchKnowledgeBaseResponseBodyResults {
	s.ScoreDetails = v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetTags(v []*string) *SearchKnowledgeBaseResponseBodyResults {
	s.Tags = v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Locations != nil {
		for _, item := range s.Locations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScoreDetails != nil {
		if err := s.ScoreDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchKnowledgeBaseResponseBodyResultsImages struct {
	// The image ID.
	//
	// example:
	//
	// kd-620ad908ec651-41253795bafd7a1c
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
	// The temporary access URL.
	//
	// example:
	//
	// https://example.com/signed-image
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SearchKnowledgeBaseResponseBodyResultsImages) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseResponseBodyResultsImages) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseResponseBodyResultsImages) GetImageId() *string {
	return s.ImageId
}

func (s *SearchKnowledgeBaseResponseBodyResultsImages) GetUrl() *string {
	return s.Url
}

func (s *SearchKnowledgeBaseResponseBodyResultsImages) SetImageId(v string) *SearchKnowledgeBaseResponseBodyResultsImages {
	s.ImageId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResultsImages) SetUrl(v string) *SearchKnowledgeBaseResponseBodyResultsImages {
	s.Url = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResultsImages) Validate() error {
	return dara.Validate(s)
}

type SearchKnowledgeBaseResponseBodyResultsLocations struct {
	// The bottom boundary.
	//
	// example:
	//
	// 364
	Bottom *int32 `json:"bottom,omitempty" xml:"bottom,omitempty"`
	// The left boundary.
	//
	// example:
	//
	// 13
	Left *int32 `json:"left,omitempty" xml:"left,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The right boundary.
	//
	// example:
	//
	// 566
	Right *int32 `json:"right,omitempty" xml:"right,omitempty"`
	// The top boundary.
	//
	// example:
	//
	// 12
	Top *int32 `json:"top,omitempty" xml:"top,omitempty"`
}

func (s SearchKnowledgeBaseResponseBodyResultsLocations) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseResponseBodyResultsLocations) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseResponseBodyResultsLocations) GetBottom() *int32 {
	return s.Bottom
}

func (s *SearchKnowledgeBaseResponseBodyResultsLocations) GetLeft() *int32 {
	return s.Left
}

func (s *SearchKnowledgeBaseResponseBodyResultsLocations) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchKnowledgeBaseResponseBodyResultsLocations) GetRight() *int32 {
	return s.Right
}

func (s *SearchKnowledgeBaseResponseBodyResultsLocations) GetTop() *int32 {
	return s.Top
}

func (s *SearchKnowledgeBaseResponseBodyResultsLocations) SetBottom(v int32) *SearchKnowledgeBaseResponseBodyResultsLocations {
	s.Bottom = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResultsLocations) SetLeft(v int32) *SearchKnowledgeBaseResponseBodyResultsLocations {
	s.Left = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResultsLocations) SetPageNumber(v int32) *SearchKnowledgeBaseResponseBodyResultsLocations {
	s.PageNumber = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResultsLocations) SetRight(v int32) *SearchKnowledgeBaseResponseBodyResultsLocations {
	s.Right = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResultsLocations) SetTop(v int32) *SearchKnowledgeBaseResponseBodyResultsLocations {
	s.Top = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResultsLocations) Validate() error {
	return dara.Validate(s)
}

type SearchKnowledgeBaseResponseBodyResultsScoreDetails struct {
	// The keyword relevance score.
	//
	// example:
	//
	// 0.0000000014285714
	KeywordScore *float32 `json:"keywordScore,omitempty" xml:"keywordScore,omitempty"`
	// The semantic relevance score.
	//
	// example:
	//
	// 0.5227377
	SemanticScore *float32 `json:"semanticScore,omitempty" xml:"semanticScore,omitempty"`
}

func (s SearchKnowledgeBaseResponseBodyResultsScoreDetails) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseResponseBodyResultsScoreDetails) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseResponseBodyResultsScoreDetails) GetKeywordScore() *float32 {
	return s.KeywordScore
}

func (s *SearchKnowledgeBaseResponseBodyResultsScoreDetails) GetSemanticScore() *float32 {
	return s.SemanticScore
}

func (s *SearchKnowledgeBaseResponseBodyResultsScoreDetails) SetKeywordScore(v float32) *SearchKnowledgeBaseResponseBodyResultsScoreDetails {
	s.KeywordScore = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResultsScoreDetails) SetSemanticScore(v float32) *SearchKnowledgeBaseResponseBodyResultsScoreDetails {
	s.SemanticScore = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResultsScoreDetails) Validate() error {
	return dara.Validate(s)
}
