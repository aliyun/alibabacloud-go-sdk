// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDisplayKnowledgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryDisplayKnowledgeResponseBody
	GetCode() *string
	SetData(v *QueryDisplayKnowledgeResponseBodyData) *QueryDisplayKnowledgeResponseBody
	GetData() *QueryDisplayKnowledgeResponseBodyData
	SetMessage(v string) *QueryDisplayKnowledgeResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryDisplayKnowledgeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryDisplayKnowledgeResponseBody
	GetSuccess() *bool
}

type QueryDisplayKnowledgeResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result of the display knowledge base content.
	Data *QueryDisplayKnowledgeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message or failure description.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 70CBEFDF-BB17-1EB3-8A21-569F3124738F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDisplayKnowledgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDisplayKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDisplayKnowledgeResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryDisplayKnowledgeResponseBody) GetData() *QueryDisplayKnowledgeResponseBodyData {
	return s.Data
}

func (s *QueryDisplayKnowledgeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryDisplayKnowledgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDisplayKnowledgeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDisplayKnowledgeResponseBody) SetCode(v string) *QueryDisplayKnowledgeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBody) SetData(v *QueryDisplayKnowledgeResponseBodyData) *QueryDisplayKnowledgeResponseBody {
	s.Data = v
	return s
}

func (s *QueryDisplayKnowledgeResponseBody) SetMessage(v string) *QueryDisplayKnowledgeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBody) SetRequestId(v string) *QueryDisplayKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBody) SetSuccess(v bool) *QueryDisplayKnowledgeResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDisplayKnowledgeResponseBodyData struct {
	// The business error code.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The business error message.
	//
	// example:
	//
	// Success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The knowledge base status.
	//
	// example:
	//
	// AVAILABLE - Available
	KnowledgeBaseStatus *string `json:"KnowledgeBaseStatus,omitempty" xml:"KnowledgeBaseStatus,omitempty"`
	// The page number.
	//
	// example:
	//
	// 10
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The knowledge base ID.
	//
	// example:
	//
	// rag_xxx
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
	// The records on the current page.
	Records []*QueryDisplayKnowledgeResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// Indicates whether the business processing is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The asynchronous task ID. This parameter is empty for synchronous API operations.
	//
	// example:
	//
	// a8323ada-a196-9061-976f-90e38b27323a
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 17
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 20
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// The usage information. This parameter is empty for free API operations.
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s QueryDisplayKnowledgeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryDisplayKnowledgeResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDisplayKnowledgeResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryDisplayKnowledgeResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryDisplayKnowledgeResponseBodyData) GetKnowledgeBaseStatus() *string {
	return s.KnowledgeBaseStatus
}

func (s *QueryDisplayKnowledgeResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryDisplayKnowledgeResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDisplayKnowledgeResponseBodyData) GetRagId() *string {
	return s.RagId
}

func (s *QueryDisplayKnowledgeResponseBodyData) GetRecords() []*QueryDisplayKnowledgeResponseBodyDataRecords {
	return s.Records
}

func (s *QueryDisplayKnowledgeResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDisplayKnowledgeResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryDisplayKnowledgeResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryDisplayKnowledgeResponseBodyData) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *QueryDisplayKnowledgeResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *QueryDisplayKnowledgeResponseBodyData) SetErrorCode(v string) *QueryDisplayKnowledgeResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyData) SetErrorMessage(v string) *QueryDisplayKnowledgeResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyData) SetKnowledgeBaseStatus(v string) *QueryDisplayKnowledgeResponseBodyData {
	s.KnowledgeBaseStatus = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyData) SetPageNum(v int32) *QueryDisplayKnowledgeResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyData) SetPageSize(v int32) *QueryDisplayKnowledgeResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyData) SetRagId(v string) *QueryDisplayKnowledgeResponseBodyData {
	s.RagId = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyData) SetRecords(v []*QueryDisplayKnowledgeResponseBodyDataRecords) *QueryDisplayKnowledgeResponseBodyData {
	s.Records = v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyData) SetSuccess(v bool) *QueryDisplayKnowledgeResponseBodyData {
	s.Success = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyData) SetTaskId(v string) *QueryDisplayKnowledgeResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyData) SetTotal(v int32) *QueryDisplayKnowledgeResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyData) SetTotalPage(v int32) *QueryDisplayKnowledgeResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyData) SetUsageMap(v map[string]*int64) *QueryDisplayKnowledgeResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDisplayKnowledgeResponseBodyDataRecords struct {
	// The description of the failure cause. This parameter is returned only when the status is FAILED.
	//
	// example:
	//
	// Image is inaccessible
	FailureMessage *string `json:"FailureMessage,omitempty" xml:"FailureMessage,omitempty"`
	// The code of the failure cause. This parameter is returned only when the status is FAILED.
	//
	// example:
	//
	// Image is inaccessible
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	// The last modified time.
	//
	// example:
	//
	// 2026-08-01T15:53:46Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The 24-hour signed URL of the transferred image.
	//
	// example:
	//
	// https://xxx/2190149752-0-cib.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The expiration time of the signed URL.
	//
	// example:
	//
	// Date
	ImageUrlExpireTime *string `json:"ImageUrlExpireTime,omitempty" xml:"ImageUrlExpireTime,omitempty"`
	// The record ID.
	//
	// example:
	//
	// 4410879640480000
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// SKU ID
	//
	// example:
	//
	// sku-001
	SkuId *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	// The SKU name.
	//
	// example:
	//
	// Product side image
	SkuName *string `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
	// The record status.
	//
	// example:
	//
	// UPDATING - Updating
	//
	// EMBEDDED - Vectorization completed
	//
	// FAILED - Processing failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryDisplayKnowledgeResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s QueryDisplayKnowledgeResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) GetFailureMessage() *string {
	return s.FailureMessage
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) GetFailureReason() *string {
	return s.FailureReason
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) GetImageUrlExpireTime() *string {
	return s.ImageUrlExpireTime
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) GetRecordId() *int64 {
	return s.RecordId
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) GetSkuId() *string {
	return s.SkuId
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) GetSkuName() *string {
	return s.SkuName
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) GetStatus() *string {
	return s.Status
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) SetFailureMessage(v string) *QueryDisplayKnowledgeResponseBodyDataRecords {
	s.FailureMessage = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) SetFailureReason(v string) *QueryDisplayKnowledgeResponseBodyDataRecords {
	s.FailureReason = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) SetGmtModified(v string) *QueryDisplayKnowledgeResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) SetImageUrl(v string) *QueryDisplayKnowledgeResponseBodyDataRecords {
	s.ImageUrl = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) SetImageUrlExpireTime(v string) *QueryDisplayKnowledgeResponseBodyDataRecords {
	s.ImageUrlExpireTime = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) SetRecordId(v int64) *QueryDisplayKnowledgeResponseBodyDataRecords {
	s.RecordId = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) SetSkuId(v string) *QueryDisplayKnowledgeResponseBodyDataRecords {
	s.SkuId = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) SetSkuName(v string) *QueryDisplayKnowledgeResponseBodyDataRecords {
	s.SkuName = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) SetStatus(v string) *QueryDisplayKnowledgeResponseBodyDataRecords {
	s.Status = &v
	return s
}

func (s *QueryDisplayKnowledgeResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
