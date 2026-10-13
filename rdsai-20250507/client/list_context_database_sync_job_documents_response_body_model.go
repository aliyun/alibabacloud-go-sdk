// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncJobDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListContextDatabaseSyncJobDocumentsResponseBody
	GetCode() *string
	SetData(v *ListContextDatabaseSyncJobDocumentsResponseBodyData) *ListContextDatabaseSyncJobDocumentsResponseBody
	GetData() *ListContextDatabaseSyncJobDocumentsResponseBodyData
	SetHttpStatusCode(v int64) *ListContextDatabaseSyncJobDocumentsResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListContextDatabaseSyncJobDocumentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListContextDatabaseSyncJobDocumentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListContextDatabaseSyncJobDocumentsResponseBody
	GetSuccess() *bool
}

type ListContextDatabaseSyncJobDocumentsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Sample value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListContextDatabaseSyncJobDocumentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListContextDatabaseSyncJobDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBody) GetData() *ListContextDatabaseSyncJobDocumentsResponseBodyData {
	return s.Data
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBody) SetCode(v string) *ListContextDatabaseSyncJobDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBody) SetData(v *ListContextDatabaseSyncJobDocumentsResponseBodyData) *ListContextDatabaseSyncJobDocumentsResponseBody {
	s.Data = v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBody) SetHttpStatusCode(v int64) *ListContextDatabaseSyncJobDocumentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBody) SetMessage(v string) *ListContextDatabaseSyncJobDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBody) SetRequestId(v string) *ListContextDatabaseSyncJobDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBody) SetSuccess(v bool) *ListContextDatabaseSyncJobDocumentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListContextDatabaseSyncJobDocumentsResponseBodyData struct {
	// The list of documents.
	//
	// example:
	//
	// [{"ProviderDocumentId":"doc-1234567890", "Status":"completed", "ContentHash":"sha256:4b3b7b2e5f6a1c9d8e0f1234567890abcdef1234567890abcdef1234567890ab", "SourceUpdateTime":"2026-09-23T10:00:00+08:00"}]
	Documents []*ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// The page number, starting from 1.
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

func (s ListContextDatabaseSyncJobDocumentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobDocumentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyData) GetDocuments() []*ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments {
	return s.Documents
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyData) SetDocuments(v []*ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) *ListContextDatabaseSyncJobDocumentsResponseBodyData {
	s.Documents = v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyData) SetPage(v int32) *ListContextDatabaseSyncJobDocumentsResponseBodyData {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyData) SetPageSize(v int32) *ListContextDatabaseSyncJobDocumentsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyData) SetTotal(v int32) *ListContextDatabaseSyncJobDocumentsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyData) Validate() error {
	if s.Documents != nil {
		for _, item := range s.Documents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments struct {
	// The content hash.
	//
	// example:
	//
	// sha256:4b3b7b2e5f6a1c9d8e0f1234567890abcdef1234567890abcdef1234567890ab
	ContentHash *string `json:"ContentHash,omitempty" xml:"ContentHash,omitempty"`
	// The ContextDB document ID.
	//
	// example:
	//
	// cd-as1llqmkol****
	CtxdbDocumentId *string `json:"CtxdbDocumentId,omitempty" xml:"CtxdbDocumentId,omitempty"`
	// The source document ID.
	//
	// example:
	//
	// doc-1234567890
	ProviderDocumentId *string `json:"ProviderDocumentId,omitempty" xml:"ProviderDocumentId,omitempty"`
	// The retirement time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	RetiredAt *string `json:"RetiredAt,omitempty" xml:"RetiredAt,omitempty"`
	// The source update time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	SourceUpdateTime *string `json:"SourceUpdateTime,omitempty" xml:"SourceUpdateTime,omitempty"`
	// The status.
	//
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) GetContentHash() *string {
	return s.ContentHash
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) GetCtxdbDocumentId() *string {
	return s.CtxdbDocumentId
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) GetProviderDocumentId() *string {
	return s.ProviderDocumentId
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) GetRetiredAt() *string {
	return s.RetiredAt
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) GetSourceUpdateTime() *string {
	return s.SourceUpdateTime
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) GetStatus() *string {
	return s.Status
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) SetContentHash(v string) *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments {
	s.ContentHash = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) SetCtxdbDocumentId(v string) *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments {
	s.CtxdbDocumentId = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) SetProviderDocumentId(v string) *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments {
	s.ProviderDocumentId = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) SetRetiredAt(v string) *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments {
	s.RetiredAt = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) SetSourceUpdateTime(v string) *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments {
	s.SourceUpdateTime = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) SetStatus(v string) *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments {
	s.Status = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) SetUpdatedAt(v string) *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments {
	s.UpdatedAt = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponseBodyDataDocuments) Validate() error {
	return dara.Validate(s)
}
