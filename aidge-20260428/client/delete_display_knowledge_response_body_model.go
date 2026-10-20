// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDisplayKnowledgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDisplayKnowledgeResponseBody
	GetCode() *string
	SetData(v *DeleteDisplayKnowledgeResponseBodyData) *DeleteDisplayKnowledgeResponseBody
	GetData() *DeleteDisplayKnowledgeResponseBodyData
	SetMessage(v string) *DeleteDisplayKnowledgeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDisplayKnowledgeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDisplayKnowledgeResponseBody
	GetSuccess() *bool
}

type DeleteDisplayKnowledgeResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The deletion results for the display knowledge base content.
	Data *DeleteDisplayKnowledgeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message or failure description.
	//
	// example:
	//
	// success
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

func (s DeleteDisplayKnowledgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDisplayKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDisplayKnowledgeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDisplayKnowledgeResponseBody) GetData() *DeleteDisplayKnowledgeResponseBodyData {
	return s.Data
}

func (s *DeleteDisplayKnowledgeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDisplayKnowledgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDisplayKnowledgeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDisplayKnowledgeResponseBody) SetCode(v string) *DeleteDisplayKnowledgeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBody) SetData(v *DeleteDisplayKnowledgeResponseBodyData) *DeleteDisplayKnowledgeResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBody) SetMessage(v string) *DeleteDisplayKnowledgeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBody) SetRequestId(v string) *DeleteDisplayKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBody) SetSuccess(v bool) *DeleteDisplayKnowledgeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteDisplayKnowledgeResponseBodyData struct {
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
	// The status of the knowledge base.
	//
	// example:
	//
	// UPDATING
	KnowledgeBaseStatus *string `json:"KnowledgeBaseStatus,omitempty" xml:"KnowledgeBaseStatus,omitempty"`
	// The knowledge base ID.
	//
	// example:
	//
	// rag_001
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
	// The number of visible records after the deletion.
	//
	// example:
	//
	// 50
	RemainingCount *int32 `json:"RemainingCount,omitempty" xml:"RemainingCount,omitempty"`
	// The processing results for each record.
	Results []*DeleteDisplayKnowledgeResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// Indicates whether the business processing is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The asynchronous task ID. This field is empty for synchronous API operations.
	//
	// example:
	//
	// task-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The usage information. This field is empty for free API operations.
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s DeleteDisplayKnowledgeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteDisplayKnowledgeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDisplayKnowledgeResponseBodyData) GetAcceptedCount() *int32 {
	return s.AcceptedCount
}

func (s *DeleteDisplayKnowledgeResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDisplayKnowledgeResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDisplayKnowledgeResponseBodyData) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *DeleteDisplayKnowledgeResponseBodyData) GetKnowledgeBaseStatus() *string {
	return s.KnowledgeBaseStatus
}

func (s *DeleteDisplayKnowledgeResponseBodyData) GetRagId() *string {
	return s.RagId
}

func (s *DeleteDisplayKnowledgeResponseBodyData) GetRemainingCount() *int32 {
	return s.RemainingCount
}

func (s *DeleteDisplayKnowledgeResponseBodyData) GetResults() []*DeleteDisplayKnowledgeResponseBodyDataResults {
	return s.Results
}

func (s *DeleteDisplayKnowledgeResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDisplayKnowledgeResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteDisplayKnowledgeResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *DeleteDisplayKnowledgeResponseBodyData) SetAcceptedCount(v int32) *DeleteDisplayKnowledgeResponseBodyData {
	s.AcceptedCount = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyData) SetErrorCode(v string) *DeleteDisplayKnowledgeResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyData) SetErrorMessage(v string) *DeleteDisplayKnowledgeResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyData) SetFailedCount(v int32) *DeleteDisplayKnowledgeResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyData) SetKnowledgeBaseStatus(v string) *DeleteDisplayKnowledgeResponseBodyData {
	s.KnowledgeBaseStatus = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyData) SetRagId(v string) *DeleteDisplayKnowledgeResponseBodyData {
	s.RagId = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyData) SetRemainingCount(v int32) *DeleteDisplayKnowledgeResponseBodyData {
	s.RemainingCount = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyData) SetResults(v []*DeleteDisplayKnowledgeResponseBodyDataResults) *DeleteDisplayKnowledgeResponseBodyData {
	s.Results = v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyData) SetSuccess(v bool) *DeleteDisplayKnowledgeResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyData) SetTaskId(v string) *DeleteDisplayKnowledgeResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyData) SetUsageMap(v map[string]*int64) *DeleteDisplayKnowledgeResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyData) Validate() error {
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

type DeleteDisplayKnowledgeResponseBodyDataResults struct {
	// The error code for the failure.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The cause of the failure.
	//
	// example:
	//
	// Success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The index in the original request array.
	//
	// example:
	//
	// 0
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The record ID.
	//
	// example:
	//
	// 2746611117
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The SKU ID. This field is empty for deletion results.
	//
	// example:
	//
	// sku-001
	SkuId *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	// The record status. This field is empty for deletion results.
	//
	// example:
	//
	// Empty for the deletion API
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the deletion is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDisplayKnowledgeResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s DeleteDisplayKnowledgeResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) GetIndex() *int32 {
	return s.Index
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) GetRecordId() *int64 {
	return s.RecordId
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) GetSkuId() *string {
	return s.SkuId
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) GetStatus() *string {
	return s.Status
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) SetErrorCode(v string) *DeleteDisplayKnowledgeResponseBodyDataResults {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) SetErrorMessage(v string) *DeleteDisplayKnowledgeResponseBodyDataResults {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) SetIndex(v int32) *DeleteDisplayKnowledgeResponseBodyDataResults {
	s.Index = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) SetRecordId(v int64) *DeleteDisplayKnowledgeResponseBodyDataResults {
	s.RecordId = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) SetSkuId(v string) *DeleteDisplayKnowledgeResponseBodyDataResults {
	s.SkuId = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) SetStatus(v string) *DeleteDisplayKnowledgeResponseBodyDataResults {
	s.Status = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) SetSuccess(v bool) *DeleteDisplayKnowledgeResponseBodyDataResults {
	s.Success = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
