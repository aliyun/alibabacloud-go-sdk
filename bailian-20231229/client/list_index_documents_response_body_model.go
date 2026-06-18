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
	// The error status code.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data field returned by the operation.
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
	// The status code returned by the operation.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIndexDocumentsResponseBodyData struct {
	// The list of files in the knowledge base, sorted by file import time in descending order (consistent with the console).
	Documents []*ListIndexDocumentsResponseBodyDataDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// The knowledge base ID.
	//
	// example:
	//
	// pno97txxxx
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The returned page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The returned number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned results.
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

type ListIndexDocumentsResponseBodyDataDocuments struct {
	// The error status code for the file import.
	//
	// example:
	//
	// 110002
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The file format type. Valid values: pdf, docx, doc, txt, md, pptx, ppt, png, jpg, jpeg, bmp, gif, and EXCEL.
	//
	// example:
	//
	// pdf
	DocumentType *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// The time when the file was imported to the knowledge base, in UNIX timestamp format.
	//
	// example:
	//
	// 1744856423000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The file ID.
	//
	// example:
	//
	// doc_c134aa2073204a5d936d870bf960f56axxxxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message for the file import.
	//
	// example:
	//
	// check fileUrlKey[file_path] / fileNameKey[null] / fileExtensionKey[file_extension] is invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The file name.
	//
	// example:
	//
	// product-overview
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The file size, in bytes.
	//
	// example:
	//
	// 996765
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// <props="china">
	//
	// For document search or audio/video search knowledge bases, this parameter specifies the category ID, which is the `CategoryId` returned by the **AddCategory*	- operation. You can also obtain the category ID by clicking the ID icon next to the category name on the Files tab of the [Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center) page.
	//
	//
	// For data query or image Q&A knowledge bases, this parameter specifies the data table ID. You can obtain the data table ID by clicking the ID icon next to the data table name on the Tables tab of the [Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center) page.
	//
	//
	//
	//
	// <props="intl">
	//
	// For document search knowledge bases, this parameter specifies the category ID, which is the `CategoryId` returned by the **AddCategory*	- operation. You can also obtain the category ID by clicking the ID icon next to the category name on the Files tab of the [Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center) page.
	//
	//
	// For data query or image Q&A knowledge bases, this parameter specifies the data table ID. You can obtain the data table ID by clicking the ID icon next to the data table name on the Tables tab of the [Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) page.
	//
	// .
	//
	// example:
	//
	// cate_21a407a3372c4ba7aedc649709143f0cxxxxxxxx
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The file import status. Valid values:
	//
	// - INSERT_ERROR: The file failed to be imported.
	//
	// - RUNNING: The file is being imported.
	//
	// - DELETED: The file has been deleted.
	//
	// - FINISH: The file was imported.
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
