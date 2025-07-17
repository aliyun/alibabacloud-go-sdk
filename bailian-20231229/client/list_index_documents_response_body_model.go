// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIndexDocumentsResponseBody
	GetCode() *string
	SetData(v *ListIndexDocumentsResponseBodyData) *ListIndexDocumentsResponseBody
	GetData() *ListIndexDocumentsResponseBodyData
	SetMessage(v string) *ListIndexDocumentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIndexDocumentsResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListIndexDocumentsResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ListIndexDocumentsResponseBody
	GetSuccess() *bool
}

type ListIndexDocumentsResponseBody struct {
	// HTTP status code
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListIndexDocumentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 35A267BF-xxxx-54DB-8394-AA3B0742D833
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indications whether the API call is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIndexDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIndexDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndexDocumentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIndexDocumentsResponseBody) GetData() *ListIndexDocumentsResponseBodyData {
	return s.Data
}

func (s *ListIndexDocumentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIndexDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIndexDocumentsResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListIndexDocumentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListIndexDocumentsResponseBody) SetCode(v string) *ListIndexDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListIndexDocumentsResponseBody) SetData(v *ListIndexDocumentsResponseBodyData) *ListIndexDocumentsResponseBody {
	s.Data = v
	return s
}

func (s *ListIndexDocumentsResponseBody) SetMessage(v string) *ListIndexDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListIndexDocumentsResponseBody) SetRequestId(v string) *ListIndexDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndexDocumentsResponseBody) SetStatus(v string) *ListIndexDocumentsResponseBody {
	s.Status = &v
	return s
}

func (s *ListIndexDocumentsResponseBody) SetSuccess(v bool) *ListIndexDocumentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListIndexDocumentsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIndexDocumentsResponseBodyData struct {
	// The list of documents in the knowledge base.
	Documents []*ListIndexDocumentsResponseBodyDataDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// The primary key ID of the knowledge base.
	//
	// example:
	//
	// pno97tn8iu
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The specified page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The specified number of documents on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of documents returned.
	//
	// example:
	//
	// 2437
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIndexDocumentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIndexDocumentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIndexDocumentsResponseBodyData) GetDocuments() []*ListIndexDocumentsResponseBodyDataDocuments {
	return s.Documents
}

func (s *ListIndexDocumentsResponseBodyData) GetIndexId() *string {
	return s.IndexId
}

func (s *ListIndexDocumentsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIndexDocumentsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIndexDocumentsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIndexDocumentsResponseBodyData) SetDocuments(v []*ListIndexDocumentsResponseBodyDataDocuments) *ListIndexDocumentsResponseBodyData {
	s.Documents = v
	return s
}

func (s *ListIndexDocumentsResponseBodyData) SetIndexId(v string) *ListIndexDocumentsResponseBodyData {
	s.IndexId = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyData) SetPageNumber(v int32) *ListIndexDocumentsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyData) SetPageSize(v int32) *ListIndexDocumentsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyData) SetTotalCount(v int64) *ListIndexDocumentsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListIndexDocumentsResponseBodyDataDocuments struct {
	// The error status code of document import.
	//
	// example:
	//
	// 110002
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The format of the document. Valid values: pdf, docx, doc, txt, md, pptx, ppt, and EXCEL.
	//
	// example:
	//
	// pdf
	DocumentType *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	GmtModified  *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The primary key ID of the document.
	//
	// example:
	//
	// doc_c134aa2073204a5d936d870bf960f56a10024701
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message of document import.
	//
	// example:
	//
	// check fileUrlKey[file_path] / fileNameKey[null] / fileExtensionKey[file_extension] is invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the document.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The size of the document. Unit: bytes.
	//
	// example:
	//
	// 996764
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// For unstructured knowledge base, this parameter is the category ID. To view the category ID, you can click the ID icon next to the category name on the Unstructured Data tab of the [Data Management](https://bailian.console.aliyun.com/#/data-center) page.
	//
	// For structured knowledge base, this parameter is the data table ID. To view the table ID, you can click the ID icon next to the table name on the Structured Data tab of the [Data Management](https://bailian.console.aliyun.com/#/data-center) page.
	//
	// example:
	//
	// cate_21a407a3372c4ba7aedc649709143f0c10021401
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The import status of the document. Valid values:
	//
	// 	- INSERT_ERROR
	//
	// 	- RUNNING
	//
	// 	- DELETED
	//
	// 	- FINISH
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListIndexDocumentsResponseBodyDataDocuments) String() string {
	return dara.Prettify(s)
}

func (s ListIndexDocumentsResponseBodyDataDocuments) GoString() string {
	return s.String()
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) GetCode() *string {
	return s.Code
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) GetDocumentType() *string {
	return s.DocumentType
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) GetId() *string {
	return s.Id
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) GetMessage() *string {
	return s.Message
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) GetName() *string {
	return s.Name
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) GetSize() *int32 {
	return s.Size
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) GetSourceId() *string {
	return s.SourceId
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) GetStatus() *string {
	return s.Status
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetCode(v string) *ListIndexDocumentsResponseBodyDataDocuments {
	s.Code = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetDocumentType(v string) *ListIndexDocumentsResponseBodyDataDocuments {
	s.DocumentType = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetGmtModified(v int64) *ListIndexDocumentsResponseBodyDataDocuments {
	s.GmtModified = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetId(v string) *ListIndexDocumentsResponseBodyDataDocuments {
	s.Id = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetMessage(v string) *ListIndexDocumentsResponseBodyDataDocuments {
	s.Message = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetName(v string) *ListIndexDocumentsResponseBodyDataDocuments {
	s.Name = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetSize(v int32) *ListIndexDocumentsResponseBodyDataDocuments {
	s.Size = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetSourceId(v string) *ListIndexDocumentsResponseBodyDataDocuments {
	s.SourceId = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetStatus(v string) *ListIndexDocumentsResponseBodyDataDocuments {
	s.Status = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) Validate() error {
	return dara.Validate(s)
}
