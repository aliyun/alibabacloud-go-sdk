// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDisplayKnowledgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddDisplayKnowledgeResponseBody
	GetCode() *string
	SetData(v *AddDisplayKnowledgeResponseBodyData) *AddDisplayKnowledgeResponseBody
	GetData() *AddDisplayKnowledgeResponseBodyData
	SetMessage(v string) *AddDisplayKnowledgeResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddDisplayKnowledgeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddDisplayKnowledgeResponseBody
	GetSuccess() *bool
}

type AddDisplayKnowledgeResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of adding content to the display knowledge base.
	Data *AddDisplayKnowledgeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDisplayKnowledgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDisplayKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *AddDisplayKnowledgeResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddDisplayKnowledgeResponseBody) GetData() *AddDisplayKnowledgeResponseBodyData {
	return s.Data
}

func (s *AddDisplayKnowledgeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddDisplayKnowledgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDisplayKnowledgeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDisplayKnowledgeResponseBody) SetCode(v string) *AddDisplayKnowledgeResponseBody {
	s.Code = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBody) SetData(v *AddDisplayKnowledgeResponseBodyData) *AddDisplayKnowledgeResponseBody {
	s.Data = v
	return s
}

func (s *AddDisplayKnowledgeResponseBody) SetMessage(v string) *AddDisplayKnowledgeResponseBody {
	s.Message = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBody) SetRequestId(v string) *AddDisplayKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBody) SetSuccess(v bool) *AddDisplayKnowledgeResponseBody {
	s.Success = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddDisplayKnowledgeResponseBodyData struct {
	// The number of accepted records.
	//
	// example:
	//
	// 10
	AcceptedCount *int32 `json:"AcceptedCount,omitempty" xml:"AcceptedCount,omitempty"`
	// The business error code.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The business error message.
	//
	// example:
	//
	// Success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The number of failed records.
	//
	// example:
	//
	// 0
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The knowledge base status.
	//
	// example:
	//
	// UPDATING-Updating
	//
	// AVAILABLE-Available
	//
	// UNAVAILABLE-Unavailable
	KnowledgeBaseStatus *string `json:"KnowledgeBaseStatus,omitempty" xml:"KnowledgeBaseStatus,omitempty"`
	// The knowledge base ID.
	//
	// example:
	//
	// rag_xxxx
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
	// The processing results for each record.
	Results []*AddDisplayKnowledgeResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// 0ea3b66e88a543658520c994f08896a0
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The usage information. This parameter is empty for free API operations.
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s AddDisplayKnowledgeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddDisplayKnowledgeResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddDisplayKnowledgeResponseBodyData) GetAcceptedCount() *int32 {
	return s.AcceptedCount
}

func (s *AddDisplayKnowledgeResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddDisplayKnowledgeResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddDisplayKnowledgeResponseBodyData) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *AddDisplayKnowledgeResponseBodyData) GetKnowledgeBaseStatus() *string {
	return s.KnowledgeBaseStatus
}

func (s *AddDisplayKnowledgeResponseBodyData) GetRagId() *string {
	return s.RagId
}

func (s *AddDisplayKnowledgeResponseBodyData) GetResults() []*AddDisplayKnowledgeResponseBodyDataResults {
	return s.Results
}

func (s *AddDisplayKnowledgeResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *AddDisplayKnowledgeResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *AddDisplayKnowledgeResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *AddDisplayKnowledgeResponseBodyData) SetAcceptedCount(v int32) *AddDisplayKnowledgeResponseBodyData {
	s.AcceptedCount = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyData) SetErrorCode(v string) *AddDisplayKnowledgeResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyData) SetErrorMessage(v string) *AddDisplayKnowledgeResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyData) SetFailedCount(v int32) *AddDisplayKnowledgeResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyData) SetKnowledgeBaseStatus(v string) *AddDisplayKnowledgeResponseBodyData {
	s.KnowledgeBaseStatus = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyData) SetRagId(v string) *AddDisplayKnowledgeResponseBodyData {
	s.RagId = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyData) SetResults(v []*AddDisplayKnowledgeResponseBodyDataResults) *AddDisplayKnowledgeResponseBodyData {
	s.Results = v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyData) SetSuccess(v bool) *AddDisplayKnowledgeResponseBodyData {
	s.Success = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyData) SetTaskId(v string) *AddDisplayKnowledgeResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyData) SetUsageMap(v map[string]*int64) *AddDisplayKnowledgeResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyData) Validate() error {
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

type AddDisplayKnowledgeResponseBodyDataResults struct {
	// The failure error code.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The failure reason.
	//
	// example:
	//
	// The deviceName format is incorrect.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The index in the original request array.
	//
	// example:
	//
	// 3
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The record ID.
	//
	// example:
	//
	// 204737029000011
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The effective SKU ID.
	//
	// example:
	//
	// sku-001
	SkuId *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	// The record status.
	//
	// example:
	//
	// UPDATING-Updating
	//
	// EMBEDDED-Vectorization completed
	//
	// FAILED-Failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the record is accepted.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDisplayKnowledgeResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s AddDisplayKnowledgeResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) GetIndex() *int32 {
	return s.Index
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) GetRecordId() *int64 {
	return s.RecordId
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) GetSkuId() *string {
	return s.SkuId
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) GetStatus() *string {
	return s.Status
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) GetSuccess() *bool {
	return s.Success
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) SetErrorCode(v string) *AddDisplayKnowledgeResponseBodyDataResults {
	s.ErrorCode = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) SetErrorMessage(v string) *AddDisplayKnowledgeResponseBodyDataResults {
	s.ErrorMessage = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) SetIndex(v int32) *AddDisplayKnowledgeResponseBodyDataResults {
	s.Index = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) SetRecordId(v int64) *AddDisplayKnowledgeResponseBodyDataResults {
	s.RecordId = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) SetSkuId(v string) *AddDisplayKnowledgeResponseBodyDataResults {
	s.SkuId = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) SetStatus(v string) *AddDisplayKnowledgeResponseBodyDataResults {
	s.Status = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) SetSuccess(v bool) *AddDisplayKnowledgeResponseBodyDataResults {
	s.Success = &v
	return s
}

func (s *AddDisplayKnowledgeResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
