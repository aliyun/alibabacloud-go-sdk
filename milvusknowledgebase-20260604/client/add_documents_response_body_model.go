// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddDocumentsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *AddDocumentsResponseBody
	GetCode() *int32
	SetData(v *AddDocumentsResponseBodyData) *AddDocumentsResponseBody
	GetData() *AddDocumentsResponseBodyData
	SetHttpStatusCode(v int32) *AddDocumentsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddDocumentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddDocumentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddDocumentsResponseBody
	GetSuccess() *bool
}

type AddDocumentsResponseBody struct {
	// The details of the permission verification failure.
	//
	// example:
	//
	// {"PolicyType":"AccountLevelIdentityBasedPolicy","AuthPrincipalOwnerId":"1234567890123456","AuthPrincipalType":"SubUser","AuthPrincipalDisplayName":"1234567890123456","NoPermissionType":"ImplicitDeny","AuthAction":"milvusknowledgebase:ListDatasets"}
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *AddDocumentsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DDDBE0E5-4314-156F-B7F1-C4BCFD25A509
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *AddDocumentsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddDocumentsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddDocumentsResponseBody) GetData() *AddDocumentsResponseBodyData {
	return s.Data
}

func (s *AddDocumentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddDocumentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDocumentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDocumentsResponseBody) SetAccessDeniedDetail(v string) *AddDocumentsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddDocumentsResponseBody) SetCode(v int32) *AddDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *AddDocumentsResponseBody) SetData(v *AddDocumentsResponseBodyData) *AddDocumentsResponseBody {
	s.Data = v
	return s
}

func (s *AddDocumentsResponseBody) SetHttpStatusCode(v int32) *AddDocumentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddDocumentsResponseBody) SetMessage(v string) *AddDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *AddDocumentsResponseBody) SetRequestId(v string) *AddDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDocumentsResponseBody) SetSuccess(v bool) *AddDocumentsResponseBody {
	s.Success = &v
	return s
}

func (s *AddDocumentsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddDocumentsResponseBodyData struct {
	// The list of documents.
	Documents []*AddDocumentsResponseBodyDataDocuments `json:"documents,omitempty" xml:"documents,omitempty" type:"Repeated"`
	// The list of errors.
	Errors []*string `json:"errors,omitempty" xml:"errors,omitempty" type:"Repeated"`
}

func (s AddDocumentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddDocumentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddDocumentsResponseBodyData) GetDocuments() []*AddDocumentsResponseBodyDataDocuments {
	return s.Documents
}

func (s *AddDocumentsResponseBodyData) GetErrors() []*string {
	return s.Errors
}

func (s *AddDocumentsResponseBodyData) SetDocuments(v []*AddDocumentsResponseBodyDataDocuments) *AddDocumentsResponseBodyData {
	s.Documents = v
	return s
}

func (s *AddDocumentsResponseBodyData) SetErrors(v []*string) *AddDocumentsResponseBodyData {
	s.Errors = v
	return s
}

func (s *AddDocumentsResponseBodyData) Validate() error {
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

type AddDocumentsResponseBodyDataDocuments struct {
	// The chunk count.
	//
	// example:
	//
	// 0
	ChunkCount *int32 `json:"chunkCount,omitempty" xml:"chunkCount,omitempty"`
	// The chunk method.
	//
	// example:
	//
	// naive
	ChunkMethod *string `json:"chunkMethod,omitempty" xml:"chunkMethod,omitempty"`
	// The ID of the knowledge base.
	//
	// example:
	//
	// kb-123
	DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	// The document ID.
	//
	// example:
	//
	// doc-123
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The object path.
	//
	// example:
	//
	// uploaded/doc-id/example.pdf
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// The name of the document.
	//
	// example:
	//
	// example.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The processing progress.
	//
	// example:
	//
	// 0
	Progress *float32 `json:"progress,omitempty" xml:"progress,omitempty"`
	// The processing status.
	//
	// example:
	//
	// RUNNING
	Run *string `json:"run,omitempty" xml:"run,omitempty"`
	// The size of the file.
	//
	// example:
	//
	// 1024
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The file extension.
	//
	// example:
	//
	// pdf
	Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
	// The thumbnail.
	//
	// example:
	//
	// ""
	Thumbnail *string `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	// The token count.
	//
	// example:
	//
	// 0
	TokenCount *int32 `json:"tokenCount,omitempty" xml:"tokenCount,omitempty"`
}

func (s AddDocumentsResponseBodyDataDocuments) String() string {
	return dara.Prettify(s)
}

func (s AddDocumentsResponseBodyDataDocuments) GoString() string {
	return s.String()
}

func (s *AddDocumentsResponseBodyDataDocuments) GetChunkCount() *int32 {
	return s.ChunkCount
}

func (s *AddDocumentsResponseBodyDataDocuments) GetChunkMethod() *string {
	return s.ChunkMethod
}

func (s *AddDocumentsResponseBodyDataDocuments) GetDatasetId() *string {
	return s.DatasetId
}

func (s *AddDocumentsResponseBodyDataDocuments) GetId() *string {
	return s.Id
}

func (s *AddDocumentsResponseBodyDataDocuments) GetLocation() *string {
	return s.Location
}

func (s *AddDocumentsResponseBodyDataDocuments) GetName() *string {
	return s.Name
}

func (s *AddDocumentsResponseBodyDataDocuments) GetProgress() *float32 {
	return s.Progress
}

func (s *AddDocumentsResponseBodyDataDocuments) GetRun() *string {
	return s.Run
}

func (s *AddDocumentsResponseBodyDataDocuments) GetSize() *int64 {
	return s.Size
}

func (s *AddDocumentsResponseBodyDataDocuments) GetSuffix() *string {
	return s.Suffix
}

func (s *AddDocumentsResponseBodyDataDocuments) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *AddDocumentsResponseBodyDataDocuments) GetTokenCount() *int32 {
	return s.TokenCount
}

func (s *AddDocumentsResponseBodyDataDocuments) SetChunkCount(v int32) *AddDocumentsResponseBodyDataDocuments {
	s.ChunkCount = &v
	return s
}

func (s *AddDocumentsResponseBodyDataDocuments) SetChunkMethod(v string) *AddDocumentsResponseBodyDataDocuments {
	s.ChunkMethod = &v
	return s
}

func (s *AddDocumentsResponseBodyDataDocuments) SetDatasetId(v string) *AddDocumentsResponseBodyDataDocuments {
	s.DatasetId = &v
	return s
}

func (s *AddDocumentsResponseBodyDataDocuments) SetId(v string) *AddDocumentsResponseBodyDataDocuments {
	s.Id = &v
	return s
}

func (s *AddDocumentsResponseBodyDataDocuments) SetLocation(v string) *AddDocumentsResponseBodyDataDocuments {
	s.Location = &v
	return s
}

func (s *AddDocumentsResponseBodyDataDocuments) SetName(v string) *AddDocumentsResponseBodyDataDocuments {
	s.Name = &v
	return s
}

func (s *AddDocumentsResponseBodyDataDocuments) SetProgress(v float32) *AddDocumentsResponseBodyDataDocuments {
	s.Progress = &v
	return s
}

func (s *AddDocumentsResponseBodyDataDocuments) SetRun(v string) *AddDocumentsResponseBodyDataDocuments {
	s.Run = &v
	return s
}

func (s *AddDocumentsResponseBodyDataDocuments) SetSize(v int64) *AddDocumentsResponseBodyDataDocuments {
	s.Size = &v
	return s
}

func (s *AddDocumentsResponseBodyDataDocuments) SetSuffix(v string) *AddDocumentsResponseBodyDataDocuments {
	s.Suffix = &v
	return s
}

func (s *AddDocumentsResponseBodyDataDocuments) SetThumbnail(v string) *AddDocumentsResponseBodyDataDocuments {
	s.Thumbnail = &v
	return s
}

func (s *AddDocumentsResponseBodyDataDocuments) SetTokenCount(v int32) *AddDocumentsResponseBodyDataDocuments {
	s.TokenCount = &v
	return s
}

func (s *AddDocumentsResponseBodyDataDocuments) Validate() error {
	return dara.Validate(s)
}
