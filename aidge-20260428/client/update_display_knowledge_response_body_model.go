// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDisplayKnowledgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDisplayKnowledgeResponseBody
	GetCode() *string
	SetData(v *UpdateDisplayKnowledgeResponseBodyData) *UpdateDisplayKnowledgeResponseBody
	GetData() *UpdateDisplayKnowledgeResponseBodyData
	SetMessage(v string) *UpdateDisplayKnowledgeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDisplayKnowledgeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDisplayKnowledgeResponseBody
	GetSuccess() *bool
}

type UpdateDisplayKnowledgeResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The update result of the display knowledge base content.
	Data *UpdateDisplayKnowledgeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 39E8A74B-F99E-1195-A5FF-3ECC5F94F304
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDisplayKnowledgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisplayKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDisplayKnowledgeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDisplayKnowledgeResponseBody) GetData() *UpdateDisplayKnowledgeResponseBodyData {
	return s.Data
}

func (s *UpdateDisplayKnowledgeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDisplayKnowledgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDisplayKnowledgeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDisplayKnowledgeResponseBody) SetCode(v string) *UpdateDisplayKnowledgeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBody) SetData(v *UpdateDisplayKnowledgeResponseBodyData) *UpdateDisplayKnowledgeResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBody) SetMessage(v string) *UpdateDisplayKnowledgeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBody) SetRequestId(v string) *UpdateDisplayKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBody) SetSuccess(v bool) *UpdateDisplayKnowledgeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDisplayKnowledgeResponseBodyData struct {
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
	// success
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
	// UPDATING
	KnowledgeBaseStatus *string `json:"KnowledgeBaseStatus,omitempty" xml:"KnowledgeBaseStatus,omitempty"`
	// The knowledge base ID.
	//
	// example:
	//
	// rag_001
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
	// The processing results for each record.
	Results []*UpdateDisplayKnowledgeResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// ed9d8504-6141-9fbb-8345-4fa36433483f
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The usage information. This parameter is empty for free API operations.
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s UpdateDisplayKnowledgeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisplayKnowledgeResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDisplayKnowledgeResponseBodyData) GetAcceptedCount() *int32 {
	return s.AcceptedCount
}

func (s *UpdateDisplayKnowledgeResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDisplayKnowledgeResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDisplayKnowledgeResponseBodyData) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *UpdateDisplayKnowledgeResponseBodyData) GetKnowledgeBaseStatus() *string {
	return s.KnowledgeBaseStatus
}

func (s *UpdateDisplayKnowledgeResponseBodyData) GetRagId() *string {
	return s.RagId
}

func (s *UpdateDisplayKnowledgeResponseBodyData) GetResults() []*UpdateDisplayKnowledgeResponseBodyDataResults {
	return s.Results
}

func (s *UpdateDisplayKnowledgeResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDisplayKnowledgeResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateDisplayKnowledgeResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *UpdateDisplayKnowledgeResponseBodyData) SetAcceptedCount(v int32) *UpdateDisplayKnowledgeResponseBodyData {
	s.AcceptedCount = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyData) SetErrorCode(v string) *UpdateDisplayKnowledgeResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyData) SetErrorMessage(v string) *UpdateDisplayKnowledgeResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyData) SetFailedCount(v int32) *UpdateDisplayKnowledgeResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyData) SetKnowledgeBaseStatus(v string) *UpdateDisplayKnowledgeResponseBodyData {
	s.KnowledgeBaseStatus = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyData) SetRagId(v string) *UpdateDisplayKnowledgeResponseBodyData {
	s.RagId = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyData) SetResults(v []*UpdateDisplayKnowledgeResponseBodyDataResults) *UpdateDisplayKnowledgeResponseBodyData {
	s.Results = v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyData) SetSuccess(v bool) *UpdateDisplayKnowledgeResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyData) SetTaskId(v string) *UpdateDisplayKnowledgeResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyData) SetUsageMap(v map[string]*int64) *UpdateDisplayKnowledgeResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyData) Validate() error {
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

type UpdateDisplayKnowledgeResponseBodyDataResults struct {
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
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The record ID.
	//
	// example:
	//
	// 32238
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
	// UPDATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the record is accepted.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDisplayKnowledgeResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisplayKnowledgeResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) GetIndex() *int32 {
	return s.Index
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) GetRecordId() *int64 {
	return s.RecordId
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) GetSkuId() *string {
	return s.SkuId
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) GetStatus() *string {
	return s.Status
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) SetErrorCode(v string) *UpdateDisplayKnowledgeResponseBodyDataResults {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) SetErrorMessage(v string) *UpdateDisplayKnowledgeResponseBodyDataResults {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) SetIndex(v int32) *UpdateDisplayKnowledgeResponseBodyDataResults {
	s.Index = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) SetRecordId(v int64) *UpdateDisplayKnowledgeResponseBodyDataResults {
	s.RecordId = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) SetSkuId(v string) *UpdateDisplayKnowledgeResponseBodyDataResults {
	s.SkuId = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) SetStatus(v string) *UpdateDisplayKnowledgeResponseBodyDataResults {
	s.Status = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) SetSuccess(v bool) *UpdateDisplayKnowledgeResponseBodyDataResults {
	s.Success = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
