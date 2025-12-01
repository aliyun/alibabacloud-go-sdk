// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocumentPageResponseBody
	GetCode() *string
	SetData(v []*GetDocumentPageResponseBodyData) *GetDocumentPageResponseBody
	GetData() []*GetDocumentPageResponseBodyData
	SetHttpStatusCode(v int32) *GetDocumentPageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDocumentPageResponseBody
	GetMessage() *string
	SetPageInfo(v *GetDocumentPageResponseBodyPageInfo) *GetDocumentPageResponseBody
	GetPageInfo() *GetDocumentPageResponseBodyPageInfo
	SetRequestId(v string) *GetDocumentPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDocumentPageResponseBody
	GetSuccess() *bool
}

type GetDocumentPageResponseBody struct {
	// API response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response data.
	Data []*GetDocumentPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Pagination information.
	PageInfo *GetDocumentPageResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 04DAD7B4-E1DA-5C2C-8E5C-A1EDC880CF60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDocumentPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocumentPageResponseBody) GetData() []*GetDocumentPageResponseBodyData {
	return s.Data
}

func (s *GetDocumentPageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDocumentPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocumentPageResponseBody) GetPageInfo() *GetDocumentPageResponseBodyPageInfo {
	return s.PageInfo
}

func (s *GetDocumentPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDocumentPageResponseBody) SetCode(v string) *GetDocumentPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentPageResponseBody) SetData(v []*GetDocumentPageResponseBodyData) *GetDocumentPageResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentPageResponseBody) SetHttpStatusCode(v int32) *GetDocumentPageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocumentPageResponseBody) SetMessage(v string) *GetDocumentPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentPageResponseBody) SetPageInfo(v *GetDocumentPageResponseBodyPageInfo) *GetDocumentPageResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetDocumentPageResponseBody) SetRequestId(v string) *GetDocumentPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentPageResponseBody) SetSuccess(v bool) *GetDocumentPageResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocumentPageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDocumentPageResponseBodyData struct {
	// Delivered by.
	//
	// example:
	//
	// luna
	DeliveredBy *string `json:"DeliveredBy,omitempty" xml:"DeliveredBy,omitempty"`
	// Report name.
	//
	// example:
	//
	// month report
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// Service report type.
	//
	// example:
	//
	// 3
	DocumentType *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// Primary key ID of the document.
	//
	// example:
	//
	// 346409
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Report status.
	//
	// example:
	//
	// uploaded
	ReportStatus *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// Report generation time.
	//
	// example:
	//
	// 2023-03-21 17:26:34
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s GetDocumentPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentPageResponseBodyData) GetDeliveredBy() *string {
	return s.DeliveredBy
}

func (s *GetDocumentPageResponseBodyData) GetDocumentName() *string {
	return s.DocumentName
}

func (s *GetDocumentPageResponseBodyData) GetDocumentType() *string {
	return s.DocumentType
}

func (s *GetDocumentPageResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetDocumentPageResponseBodyData) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *GetDocumentPageResponseBodyData) GetUploadTime() *string {
	return s.UploadTime
}

func (s *GetDocumentPageResponseBodyData) SetDeliveredBy(v string) *GetDocumentPageResponseBodyData {
	s.DeliveredBy = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) SetDocumentName(v string) *GetDocumentPageResponseBodyData {
	s.DocumentName = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) SetDocumentType(v string) *GetDocumentPageResponseBodyData {
	s.DocumentType = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) SetId(v int64) *GetDocumentPageResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) SetReportStatus(v string) *GetDocumentPageResponseBodyData {
	s.ReportStatus = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) SetUploadTime(v string) *GetDocumentPageResponseBodyData {
	s.UploadTime = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDocumentPageResponseBodyPageInfo struct {
	// The current page number in pagination queries.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number of items per page in the returned data.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of queried items.
	//
	// example:
	//
	// 3149
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetDocumentPageResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentPageResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetDocumentPageResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetDocumentPageResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDocumentPageResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetDocumentPageResponseBodyPageInfo) SetCurrentPage(v int32) *GetDocumentPageResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetDocumentPageResponseBodyPageInfo) SetPageSize(v int32) *GetDocumentPageResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetDocumentPageResponseBodyPageInfo) SetTotalCount(v int32) *GetDocumentPageResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *GetDocumentPageResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
