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
	SetEvidenceNextToken(v string) *SearchKnowledgeBaseResponseBody
	GetEvidenceNextToken() *string
	SetEvidenceResults(v []*SearchKnowledgeBaseResponseBodyEvidenceResults) *SearchKnowledgeBaseResponseBody
	GetEvidenceResults() []*SearchKnowledgeBaseResponseBodyEvidenceResults
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
	// The access denial details.
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
	// Optional. Conditionally returned. Required when includeWiki is set to true. The cursor for the next page of evidence. Null for the last page. Not used for legacy results pagination. Null is allowed. The not-equal-to condition can be omitted. Constraints: {"minLength":1,"maxLength":4096}.
	//
	// example:
	//
	// example-opaque-next-token
	EvidenceNextToken *string `json:"evidenceNextToken,omitempty" xml:"evidenceNextToken,omitempty"`
	// Optional. Conditionally returned. Returned only when includeWiki is set to true. The current page of unified and merged evidence. Constraints: {"minItems":0,"maxItems":100}.
	//
	// example:
	//
	// [{"evidenceId":"wiki_section_page_1_sec_1_rev_7","kind":"wiki_section","knowledgeBaseId":"example_kb","documentId":null,"chunkId":null,"pageId":"page_1","sectionId":"sec_1","entityId":null,"relationId":null,"revisionId":"rev_7","batchId":"batch_7","title":"Warranty instructions","content":"The warranty period for Product A is 18 months.","score":0.62,"sourceRefs":[{"documentId":"doc_manual","chunkId":"chunk_1","sourceRevisionId":"source_rev_7","version":"DRAFT"}],"sourceRefCount":1,"sourceRefsTruncated":false,"sourceTaskId":"11111111111111111111111111111111"}]
	EvidenceResults []*SearchKnowledgeBaseResponseBodyEvidenceResults `json:"evidenceResults,omitempty" xml:"evidenceResults,omitempty" type:"Repeated"`
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
	// The number of items per page.
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
	// Indicates whether the request is successful.
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

func (s *SearchKnowledgeBaseResponseBody) GetEvidenceNextToken() *string {
	return s.EvidenceNextToken
}

func (s *SearchKnowledgeBaseResponseBody) GetEvidenceResults() []*SearchKnowledgeBaseResponseBodyEvidenceResults {
	return s.EvidenceResults
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

func (s *SearchKnowledgeBaseResponseBody) SetEvidenceNextToken(v string) *SearchKnowledgeBaseResponseBody {
	s.EvidenceNextToken = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBody) SetEvidenceResults(v []*SearchKnowledgeBaseResponseBodyEvidenceResults) *SearchKnowledgeBaseResponseBody {
	s.EvidenceResults = v
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
	if s.EvidenceResults != nil {
		for _, item := range s.EvidenceResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type SearchKnowledgeBaseResponseBodyEvidenceResults struct {
	// Required. Always returned. Null is allowed. The semantics of null values follow the field description. Null is allowed. The not-equal-to condition can be omitted. Constraints: {"minLength":1,"maxLength":128,"pattern":"^[A-Za-z0-9][A-Za-z0-9_-]{0,127}$"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// null
	BatchId *string `json:"batchId,omitempty" xml:"batchId,omitempty"`
	// Required. Always returned. Null is allowed. The semantics of null values follow the field description. Null is allowed. The not-equal-to condition can be omitted. Constraints: {"minLength":1,"maxLength":128,"pattern":"^[A-Za-z0-9][A-Za-z0-9_-]{0,127}$"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// null
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// Required. Always returned. The readable evidence text. Constraints: {"maxLength":65535}.
	//
	// This parameter is required.
	//
	// example:
	//
	// The warranty period for Product A is 18 months
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Required. Always returned. Null is allowed. The semantics of null values follow the field description. Null is allowed. The not-equal-to condition can be omitted. Constraints: {"minLength":1,"maxLength":128,"pattern":"^[A-Za-z0-9][A-Za-z0-9_-]{0,127}$"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// null
	DocumentId *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
	// Required. Always returned. Null is allowed. The semantics of null values follow the field description. Null is allowed. The not-equal-to condition can be omitted. Constraints: {"minLength":1,"maxLength":128,"pattern":"^[A-Za-z0-9][A-Za-z0-9_-]{0,127}$"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// null
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// Required. Always returned. The opaque business ID generated by the server. Validates ownership and does not accept arbitrary paths. Constraints: {"minLength":1,"maxLength":128,"pattern":"^[A-Za-z0-9][A-Za-z0-9_-]{0,127}$"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// wiki_section_page_1_sec_1_rev_7
	EvidenceId *string `json:"evidenceId,omitempty" xml:"evidenceId,omitempty"`
	// Required. Always returned. The actual evidence object type. Constraints: {"enum":["raw","wiki_page","wiki_section","graph_fact"],"maxLength":64}.
	//
	// This parameter is required.
	//
	// example:
	//
	// wiki_section
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// Required. Always returned. The opaque business ID generated by the server. Validates ownership and does not accept arbitrary paths. Constraints: {"minLength":1,"maxLength":128,"pattern":"^[A-Za-z0-9][A-Za-z0-9_-]{0,127}$"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// example_kb
	KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty" xml:"knowledgeBaseId,omitempty"`
	// Required. Always returned. Null is allowed. The semantics of null values follow the field description. Null is allowed. The not-equal-to condition can be omitted. Constraints: {"minLength":1,"maxLength":128,"pattern":"^[A-Za-z0-9][A-Za-z0-9_-]{0,127}$"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// page_1
	PageId *string `json:"pageId,omitempty" xml:"pageId,omitempty"`
	// Required. Always returned. Null is allowed. The semantics of null values follow the field description. Null is allowed. The not-equal-to condition can be omitted. Constraints: {"minLength":1,"maxLength":128,"pattern":"^[A-Za-z0-9][A-Za-z0-9_-]{0,127}$"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// null
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// Required. Always returned. The opaque business ID generated by the server. Validates ownership and does not accept arbitrary paths. Constraints: {"minLength":1,"maxLength":128,"pattern":"^[A-Za-z0-9][A-Za-z0-9_-]{0,127}$"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// example_value
	RevisionId *string `json:"revisionId,omitempty" xml:"revisionId,omitempty"`
	// Required. Always returned. The finite numerical value used for unified ranking. For hybrid, it is the relevance of the existing Dealer.rerank. For rerank, it is the text and model hybrid relevance of the existing Dealer.rerank_by_model, along the semanticWeight and rank feature, which is not a relevance probability. The minScore is applied according to the existing unified scoring rules. Can be null for non-ranking fact evidence. Null is allowed. The not-equal-to condition can be omitted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.62
	Score *float64 `json:"score,omitempty" xml:"score,omitempty"`
	// Required. Always returned. Null is allowed. The semantics of null values follow the field description. Null is allowed. The not-equal-to condition can be omitted. Constraints: {"minLength":1,"maxLength":128,"pattern":"^[A-Za-z0-9][A-Za-z0-9_-]{0,127}$"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// sec_1
	SectionId *string `json:"sectionId,omitempty" xml:"sectionId,omitempty"`
	// Required. Always returned. The total number of source chunk references. Authorization is performed on the full set. Constraints: {"minimum":1,"maximum":9223372036854775807}.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SourceRefCount *int64 `json:"sourceRefCount,omitempty" xml:"sourceRefCount,omitempty"`
	// Required. Always returned. The actually supported sources. Constraints: {"minItems":1,"maxItems":1000}.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"documentId":"doc_manual","chunkId":"chunk_1","sourceRevisionId":"source_rev_7","version":"DRAFT","contentHash":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}]
	SourceRefs []*SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs `json:"sourceRefs,omitempty" xml:"sourceRefs,omitempty" type:"Repeated"`
	// Required. Always returned. True when a single sourceRefs contains only the first 1000 items. You can paginate and verify by using sourceTaskId.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	SourceRefsTruncated *bool `json:"sourceRefsTruncated,omitempty" xml:"sourceRefsTruncated,omitempty"`
	// Required. Always returned. The entry point for complete sources. Must be a non-null locator when sourceRefsTruncated is true. The source operation references held by the artifact or revision remain readable, and archiving task details does not delete this source manifest. Null is allowed. The not-equal-to condition can be omitted. Constraints: {"minLength":32,"maxLength":32,"pattern":"^[0-9a-f]{32}$"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// null
	SourceTaskId *string `json:"sourceTaskId,omitempty" xml:"sourceTaskId,omitempty"`
	// Required. Always returned. The display title after authorization. Constraints: {"maxLength":1024}.
	//
	// This parameter is required.
	//
	// example:
	//
	// Warranty instructions
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SearchKnowledgeBaseResponseBodyEvidenceResults) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseResponseBodyEvidenceResults) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetBatchId() *string {
	return s.BatchId
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetChunkId() *string {
	return s.ChunkId
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetContent() *string {
	return s.Content
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetDocumentId() *string {
	return s.DocumentId
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetEntityId() *string {
	return s.EntityId
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetEvidenceId() *string {
	return s.EvidenceId
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetKind() *string {
	return s.Kind
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetPageId() *string {
	return s.PageId
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetRelationId() *string {
	return s.RelationId
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetRevisionId() *string {
	return s.RevisionId
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetScore() *float64 {
	return s.Score
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetSectionId() *string {
	return s.SectionId
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetSourceRefCount() *int64 {
	return s.SourceRefCount
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetSourceRefs() []*SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs {
	return s.SourceRefs
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetSourceRefsTruncated() *bool {
	return s.SourceRefsTruncated
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) GetTitle() *string {
	return s.Title
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetBatchId(v string) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.BatchId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetChunkId(v string) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.ChunkId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetContent(v string) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.Content = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetDocumentId(v string) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.DocumentId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetEntityId(v string) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.EntityId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetEvidenceId(v string) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.EvidenceId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetKind(v string) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.Kind = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetKnowledgeBaseId(v string) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.KnowledgeBaseId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetPageId(v string) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.PageId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetRelationId(v string) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.RelationId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetRevisionId(v string) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.RevisionId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetScore(v float64) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.Score = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetSectionId(v string) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.SectionId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetSourceRefCount(v int64) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.SourceRefCount = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetSourceRefs(v []*SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.SourceRefs = v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetSourceRefsTruncated(v bool) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.SourceRefsTruncated = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetSourceTaskId(v string) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.SourceTaskId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) SetTitle(v string) *SearchKnowledgeBaseResponseBodyEvidenceResults {
	s.Title = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResults) Validate() error {
	if s.SourceRefs != nil {
		for _, item := range s.SourceRefs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs struct {
	// Required. Always returned. The opaque business ID generated by the server. Validates ownership and does not accept arbitrary paths. Constraints: {"minLength":1,"maxLength":128,"pattern":"^[A-Za-z0-9][A-Za-z0-9_-]{0,127}$"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// example_value
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// Required. Always returned. The opaque business ID generated by the server. Validates ownership and does not accept arbitrary paths. Constraints: {"minLength":1,"maxLength":128,"pattern":"^[A-Za-z0-9][A-Za-z0-9_-]{0,127}$"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// example_value
	DocumentId *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
	// Optional. Conditionally returned. The document name returned only after authorization. Constraints: {"maxLength":1024}.
	//
	// example:
	//
	// example_value
	DocumentName *string `json:"documentName,omitempty" xml:"documentName,omitempty"`
	// Optional. Conditionally returned. Null is allowed. The semantics of null values follow the field description. Null is allowed. The not-equal-to condition can be omitted. Constraints: {"minimum":0,"maximum":9223372036854775807}.
	//
	// example:
	//
	// 0
	MediaEndMs *int64 `json:"mediaEndMs,omitempty" xml:"mediaEndMs,omitempty"`
	// Optional. Conditionally returned. Null is allowed. The semantics of null values follow the field description. Null is allowed. The not-equal-to condition can be omitted. Constraints: {"minimum":0,"maximum":9223372036854775807}.
	//
	// example:
	//
	// 0
	MediaStartMs *int64 `json:"mediaStartMs,omitempty" xml:"mediaStartMs,omitempty"`
	// Optional. Conditionally returned. Null is allowed. The semantics of null values follow the field description. Null is allowed. The not-equal-to condition can be omitted. Constraints: {"minimum":1,"maximum":9223372036854775807}.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Required. Always returned. The opaque business ID generated by the server. Validates ownership and does not accept arbitrary paths. Constraints: {"minLength":1,"maxLength":128,"pattern":"^[A-Za-z0-9][A-Za-z0-9_-]{0,127}$"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// example_value
	SourceRevisionId *string `json:"sourceRevisionId,omitempty" xml:"sourceRevisionId,omitempty"`
	// Required. Always returned. The display name of the version. Does not fall back to draft if invalid or unavailable. Constraints: {"maxLength":64,"pattern":"^(DRAFT|LATEST_PUBLISHED|v[1-9][0-9]*)$","default":"DRAFT"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// DRAFT
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) GetChunkId() *string {
	return s.ChunkId
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) GetDocumentId() *string {
	return s.DocumentId
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) GetDocumentName() *string {
	return s.DocumentName
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) GetMediaEndMs() *int64 {
	return s.MediaEndMs
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) GetMediaStartMs() *int64 {
	return s.MediaStartMs
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) GetSourceRevisionId() *string {
	return s.SourceRevisionId
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) GetVersion() *string {
	return s.Version
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) SetChunkId(v string) *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs {
	s.ChunkId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) SetDocumentId(v string) *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs {
	s.DocumentId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) SetDocumentName(v string) *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs {
	s.DocumentName = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) SetMediaEndMs(v int64) *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs {
	s.MediaEndMs = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) SetMediaStartMs(v int64) *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs {
	s.MediaStartMs = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) SetPageNumber(v int64) *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs {
	s.PageNumber = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) SetSourceRevisionId(v string) *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs {
	s.SourceRevisionId = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) SetVersion(v string) *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs {
	s.Version = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyEvidenceResultsSourceRefs) Validate() error {
	return dara.Validate(s)
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
	// The ID of the knowledge base.
	//
	// example:
	//
	// kd-xxxxxxxxxx
	KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty" xml:"knowledgeBaseId,omitempty"`
	// The list of document locations.
	Locations []*SearchKnowledgeBaseResponseBodyResultsLocations `json:"locations,omitempty" xml:"locations,omitempty" type:"Repeated"`
	// The end time of the audio or video chunk in the original media, in milliseconds. This parameter is not returned for non-media chunks.
	//
	// example:
	//
	// 8400
	MediaEndMs *int64 `json:"mediaEndMs,omitempty" xml:"mediaEndMs,omitempty"`
	// The start time of the audio or video chunk in the original media, in milliseconds. This parameter is not returned for non-media chunks.
	//
	// example:
	//
	// 1250
	MediaStartMs *int64 `json:"mediaStartMs,omitempty" xml:"mediaStartMs,omitempty"`
	// The time range of each sentence or paragraph within the chunk in the original media. This parameter is not returned for non-audio chunks.
	MediaTimeline []*SearchKnowledgeBaseResponseBodyResultsMediaTimeline `json:"mediaTimeline,omitempty" xml:"mediaTimeline,omitempty" type:"Repeated"`
	// The parent chunk ID.
	//
	// example:
	//
	// parent-chunk-id
	ParentChunkId *string `json:"parentChunkId,omitempty" xml:"parentChunkId,omitempty"`
	// The scalar columns of the structured knowledge base. The columns are returned by their original names and are not used for retrieval.
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

func (s *SearchKnowledgeBaseResponseBodyResults) GetMediaEndMs() *int64 {
	return s.MediaEndMs
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetMediaStartMs() *int64 {
	return s.MediaStartMs
}

func (s *SearchKnowledgeBaseResponseBodyResults) GetMediaTimeline() []*SearchKnowledgeBaseResponseBodyResultsMediaTimeline {
	return s.MediaTimeline
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

func (s *SearchKnowledgeBaseResponseBodyResults) SetMediaEndMs(v int64) *SearchKnowledgeBaseResponseBodyResults {
	s.MediaEndMs = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetMediaStartMs(v int64) *SearchKnowledgeBaseResponseBodyResults {
	s.MediaStartMs = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResults) SetMediaTimeline(v []*SearchKnowledgeBaseResponseBodyResultsMediaTimeline) *SearchKnowledgeBaseResponseBodyResults {
	s.MediaTimeline = v
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
	if s.MediaTimeline != nil {
		for _, item := range s.MediaTimeline {
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
	// The bottom edge.
	//
	// example:
	//
	// 364
	Bottom *int32 `json:"bottom,omitempty" xml:"bottom,omitempty"`
	// The left edge.
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
	// The right edge.
	//
	// example:
	//
	// 566
	Right *int32 `json:"right,omitempty" xml:"right,omitempty"`
	// The top edge.
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

type SearchKnowledgeBaseResponseBodyResultsMediaTimeline struct {
	// The end time of the segment.
	//
	// example:
	//
	// 4200
	EndMs *int64 `json:"endMs,omitempty" xml:"endMs,omitempty"`
	// The start time of the segment.
	//
	// example:
	//
	// 1250
	StartMs *int64 `json:"startMs,omitempty" xml:"startMs,omitempty"`
	// The text of the segment.
	//
	// example:
	//
	// The first sentence
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s SearchKnowledgeBaseResponseBodyResultsMediaTimeline) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseResponseBodyResultsMediaTimeline) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseResponseBodyResultsMediaTimeline) GetEndMs() *int64 {
	return s.EndMs
}

func (s *SearchKnowledgeBaseResponseBodyResultsMediaTimeline) GetStartMs() *int64 {
	return s.StartMs
}

func (s *SearchKnowledgeBaseResponseBodyResultsMediaTimeline) GetText() *string {
	return s.Text
}

func (s *SearchKnowledgeBaseResponseBodyResultsMediaTimeline) SetEndMs(v int64) *SearchKnowledgeBaseResponseBodyResultsMediaTimeline {
	s.EndMs = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResultsMediaTimeline) SetStartMs(v int64) *SearchKnowledgeBaseResponseBodyResultsMediaTimeline {
	s.StartMs = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResultsMediaTimeline) SetText(v string) *SearchKnowledgeBaseResponseBodyResultsMediaTimeline {
	s.Text = &v
	return s
}

func (s *SearchKnowledgeBaseResponseBodyResultsMediaTimeline) Validate() error {
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
