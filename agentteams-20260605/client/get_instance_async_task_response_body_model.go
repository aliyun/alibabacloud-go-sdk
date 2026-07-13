// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAsyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceAsyncTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetInstanceAsyncTaskResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*GetInstanceAsyncTaskResponseBodyItems) *GetInstanceAsyncTaskResponseBody
	GetItems() []*GetInstanceAsyncTaskResponseBodyItems
	SetMaxResults(v int32) *GetInstanceAsyncTaskResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *GetInstanceAsyncTaskResponseBody
	GetMessage() *string
	SetNextToken(v string) *GetInstanceAsyncTaskResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetInstanceAsyncTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceAsyncTaskResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetInstanceAsyncTaskResponseBody
	GetTotalCount() *int32
}

type GetInstanceAsyncTaskResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          []*GetInstanceAsyncTaskResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetInstanceAsyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceAsyncTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceAsyncTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceAsyncTaskResponseBody) GetItems() []*GetInstanceAsyncTaskResponseBodyItems {
	return s.Items
}

func (s *GetInstanceAsyncTaskResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetInstanceAsyncTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceAsyncTaskResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetInstanceAsyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceAsyncTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceAsyncTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetInstanceAsyncTaskResponseBody) SetCode(v string) *GetInstanceAsyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBody) SetHttpStatusCode(v int32) *GetInstanceAsyncTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBody) SetItems(v []*GetInstanceAsyncTaskResponseBodyItems) *GetInstanceAsyncTaskResponseBody {
	s.Items = v
	return s
}

func (s *GetInstanceAsyncTaskResponseBody) SetMaxResults(v int32) *GetInstanceAsyncTaskResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBody) SetMessage(v string) *GetInstanceAsyncTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBody) SetNextToken(v string) *GetInstanceAsyncTaskResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBody) SetRequestId(v string) *GetInstanceAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBody) SetSuccess(v bool) *GetInstanceAsyncTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBody) SetTotalCount(v int32) *GetInstanceAsyncTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBody) Validate() error {
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

type GetInstanceAsyncTaskResponseBodyItems struct {
	CurrentStep          *string                                               `json:"CurrentStep,omitempty" xml:"CurrentStep,omitempty"`
	RecoveryMessage      *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage `json:"RecoveryMessage,omitempty" xml:"RecoveryMessage,omitempty" type:"Struct"`
	TaskCode             *string                                               `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	TaskId               *string                                               `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus           *string                                               `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	WaitingForUserAction *bool                                                 `json:"WaitingForUserAction,omitempty" xml:"WaitingForUserAction,omitempty"`
}

func (s GetInstanceAsyncTaskResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAsyncTaskResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetInstanceAsyncTaskResponseBodyItems) GetCurrentStep() *string {
	return s.CurrentStep
}

func (s *GetInstanceAsyncTaskResponseBodyItems) GetRecoveryMessage() *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage {
	return s.RecoveryMessage
}

func (s *GetInstanceAsyncTaskResponseBodyItems) GetTaskCode() *string {
	return s.TaskCode
}

func (s *GetInstanceAsyncTaskResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *GetInstanceAsyncTaskResponseBodyItems) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetInstanceAsyncTaskResponseBodyItems) GetWaitingForUserAction() *bool {
	return s.WaitingForUserAction
}

func (s *GetInstanceAsyncTaskResponseBodyItems) SetCurrentStep(v string) *GetInstanceAsyncTaskResponseBodyItems {
	s.CurrentStep = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBodyItems) SetRecoveryMessage(v *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) *GetInstanceAsyncTaskResponseBodyItems {
	s.RecoveryMessage = v
	return s
}

func (s *GetInstanceAsyncTaskResponseBodyItems) SetTaskCode(v string) *GetInstanceAsyncTaskResponseBodyItems {
	s.TaskCode = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBodyItems) SetTaskId(v string) *GetInstanceAsyncTaskResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBodyItems) SetTaskStatus(v string) *GetInstanceAsyncTaskResponseBodyItems {
	s.TaskStatus = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBodyItems) SetWaitingForUserAction(v bool) *GetInstanceAsyncTaskResponseBodyItems {
	s.WaitingForUserAction = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBodyItems) Validate() error {
	if s.RecoveryMessage != nil {
		if err := s.RecoveryMessage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage struct {
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OccurredAt         *string `json:"OccurredAt,omitempty" xml:"OccurredAt,omitempty"`
	RecoverySuggestion *string `json:"RecoverySuggestion,omitempty" xml:"RecoverySuggestion,omitempty"`
	Retryable          *bool   `json:"Retryable,omitempty" xml:"Retryable,omitempty"`
	Source             *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Type               *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) GoString() string {
	return s.String()
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) GetCode() *string {
	return s.Code
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) GetOccurredAt() *string {
	return s.OccurredAt
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) GetRecoverySuggestion() *string {
	return s.RecoverySuggestion
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) GetRetryable() *bool {
	return s.Retryable
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) GetSource() *string {
	return s.Source
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) GetType() *string {
	return s.Type
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) SetCode(v string) *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage {
	s.Code = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) SetMessage(v string) *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage {
	s.Message = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) SetOccurredAt(v string) *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage {
	s.OccurredAt = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) SetRecoverySuggestion(v string) *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage {
	s.RecoverySuggestion = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) SetRetryable(v bool) *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage {
	s.Retryable = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) SetSource(v string) *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage {
	s.Source = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) SetType(v string) *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage {
	s.Type = &v
	return s
}

func (s *GetInstanceAsyncTaskResponseBodyItemsRecoveryMessage) Validate() error {
	return dara.Validate(s)
}
