// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncRunItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListContextDatabaseSyncRunItemsResponseBody
	GetCode() *string
	SetData(v *ListContextDatabaseSyncRunItemsResponseBodyData) *ListContextDatabaseSyncRunItemsResponseBody
	GetData() *ListContextDatabaseSyncRunItemsResponseBodyData
	SetHttpStatusCode(v int64) *ListContextDatabaseSyncRunItemsResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListContextDatabaseSyncRunItemsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListContextDatabaseSyncRunItemsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListContextDatabaseSyncRunItemsResponseBody
	GetSuccess() *bool
}

type ListContextDatabaseSyncRunItemsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Sample value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListContextDatabaseSyncRunItemsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message.
	//
	// example:
	//
	// Sample value
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// id-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListContextDatabaseSyncRunItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunItemsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListContextDatabaseSyncRunItemsResponseBody) GetData() *ListContextDatabaseSyncRunItemsResponseBodyData {
	return s.Data
}

func (s *ListContextDatabaseSyncRunItemsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListContextDatabaseSyncRunItemsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContextDatabaseSyncRunItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseSyncRunItemsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListContextDatabaseSyncRunItemsResponseBody) SetCode(v string) *ListContextDatabaseSyncRunItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBody) SetData(v *ListContextDatabaseSyncRunItemsResponseBodyData) *ListContextDatabaseSyncRunItemsResponseBody {
	s.Data = v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBody) SetHttpStatusCode(v int64) *ListContextDatabaseSyncRunItemsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBody) SetMessage(v string) *ListContextDatabaseSyncRunItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBody) SetRequestId(v string) *ListContextDatabaseSyncRunItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBody) SetSuccess(v bool) *ListContextDatabaseSyncRunItemsResponseBody {
	s.Success = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListContextDatabaseSyncRunItemsResponseBodyData struct {
	// The data list.
	//
	// example:
	//
	// [{"SourceRecordId":"sr-as1llqmkol****", "ArtifactIndex":0, "ProviderDocumentId":"doc-1234567890", "Title":"Project Plan"}]
	Items []*ListContextDatabaseSyncRunItemsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListContextDatabaseSyncRunItemsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyData) GetItems() []*ListContextDatabaseSyncRunItemsResponseBodyDataItems {
	return s.Items
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyData) SetItems(v []*ListContextDatabaseSyncRunItemsResponseBodyDataItems) *ListContextDatabaseSyncRunItemsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyData) SetPage(v int32) *ListContextDatabaseSyncRunItemsResponseBodyData {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyData) SetPageSize(v int32) *ListContextDatabaseSyncRunItemsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyData) SetTotal(v int32) *ListContextDatabaseSyncRunItemsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextDatabaseSyncRunItemsResponseBodyDataItems struct {
	// The document shard index.
	//
	// example:
	//
	// 0
	ArtifactIndex *int32 `json:"ArtifactIndex,omitempty" xml:"ArtifactIndex,omitempty"`
	// The ContextDB document ID.
	//
	// example:
	//
	// cd-as1llqmkol****
	CtxdbDocumentId *string `json:"CtxdbDocumentId,omitempty" xml:"CtxdbDocumentId,omitempty"`
	// The fine-grained execution status of the current synchronization item, which indicates the processing result of a single document in the current run. Common valid values include uploaded, uploaded-pending-review, upload-accepted, upload-busy, metadata-updating, metadata-updated, migrated, failed, source-failed, replace, warning, deferred, resume, skip, skipped, unsupported, and cancelled.
	//
	// example:
	//
	// uploaded
	Decision *string `json:"Decision,omitempty" xml:"Decision,omitempty"`
	// The aggregated display status of the current synchronization item. Known mappings include success, failure, processing, and skipped. Unknown Decision values are passed through as-is.
	//
	// example:
	//
	// success
	DecisionCategory *string `json:"DecisionCategory,omitempty" xml:"DecisionCategory,omitempty"`
	// The source document ID.
	//
	// example:
	//
	// doc-1234567890
	ProviderDocumentId *string `json:"ProviderDocumentId,omitempty" xml:"ProviderDocumentId,omitempty"`
	// The processing reason or failure reason.
	//
	// example:
	//
	// The data source authorization has expired
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The size in bytes.
	//
	// example:
	//
	// 1048576
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	// The source path.
	//
	// example:
	//
	// /Development Documents/Project Plan
	SourcePath *string `json:"SourcePath,omitempty" xml:"SourcePath,omitempty"`
	// The source record ID.
	//
	// example:
	//
	// sr-as1llqmkol****
	SourceRecordId *string `json:"SourceRecordId,omitempty" xml:"SourceRecordId,omitempty"`
	// The title.
	//
	// example:
	//
	// Project Plan
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListContextDatabaseSyncRunItemsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunItemsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) GetArtifactIndex() *int32 {
	return s.ArtifactIndex
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) GetCtxdbDocumentId() *string {
	return s.CtxdbDocumentId
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) GetDecision() *string {
	return s.Decision
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) GetDecisionCategory() *string {
	return s.DecisionCategory
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) GetProviderDocumentId() *string {
	return s.ProviderDocumentId
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) GetReason() *string {
	return s.Reason
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) GetSourcePath() *string {
	return s.SourcePath
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) GetSourceRecordId() *string {
	return s.SourceRecordId
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) GetTitle() *string {
	return s.Title
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) SetArtifactIndex(v int32) *ListContextDatabaseSyncRunItemsResponseBodyDataItems {
	s.ArtifactIndex = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) SetCtxdbDocumentId(v string) *ListContextDatabaseSyncRunItemsResponseBodyDataItems {
	s.CtxdbDocumentId = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) SetDecision(v string) *ListContextDatabaseSyncRunItemsResponseBodyDataItems {
	s.Decision = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) SetDecisionCategory(v string) *ListContextDatabaseSyncRunItemsResponseBodyDataItems {
	s.DecisionCategory = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) SetProviderDocumentId(v string) *ListContextDatabaseSyncRunItemsResponseBodyDataItems {
	s.ProviderDocumentId = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) SetReason(v string) *ListContextDatabaseSyncRunItemsResponseBodyDataItems {
	s.Reason = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) SetSizeBytes(v int64) *ListContextDatabaseSyncRunItemsResponseBodyDataItems {
	s.SizeBytes = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) SetSourcePath(v string) *ListContextDatabaseSyncRunItemsResponseBodyDataItems {
	s.SourcePath = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) SetSourceRecordId(v string) *ListContextDatabaseSyncRunItemsResponseBodyDataItems {
	s.SourceRecordId = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) SetTitle(v string) *ListContextDatabaseSyncRunItemsResponseBodyDataItems {
	s.Title = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) SetUpdatedAt(v string) *ListContextDatabaseSyncRunItemsResponseBodyDataItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
