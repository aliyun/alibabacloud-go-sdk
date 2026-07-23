// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShoppingAssistantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCitation(v *ShoppingAssistantResponseBodyCitation) *ShoppingAssistantResponseBody
	GetCitation() *ShoppingAssistantResponseBodyCitation
	SetContent(v string) *ShoppingAssistantResponseBody
	GetContent() *string
	SetConversationId(v string) *ShoppingAssistantResponseBody
	GetConversationId() *string
	SetErrorCode(v string) *ShoppingAssistantResponseBody
	GetErrorCode() *string
	SetEvent(v string) *ShoppingAssistantResponseBody
	GetEvent() *string
	SetRequestId(v string) *ShoppingAssistantResponseBody
	GetRequestId() *string
	SetResult(v *ShoppingAssistantResponseBodyResult) *ShoppingAssistantResponseBody
	GetResult() *ShoppingAssistantResponseBodyResult
	SetSessionId(v string) *ShoppingAssistantResponseBody
	GetSessionId() *string
	SetStopReason(v string) *ShoppingAssistantResponseBody
	GetStopReason() *string
}

type ShoppingAssistantResponseBody struct {
	// The citation information.
	Citation *ShoppingAssistantResponseBodyCitation `json:"Citation,omitempty" xml:"Citation,omitempty" type:"Struct"`
	// The returned content.
	//
	// example:
	//
	// Here are some light-colored long-sleeve shirts I picked for you:\\n.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The session ID.
	//
	// example:
	//
	// e47cfae9-c0cc-42e1-91e2-e67cdb0e7b96
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The event.
	//
	// example:
	//
	// analyze_requirement
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result details.
	Result *ShoppingAssistantResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The session ID.
	//
	// example:
	//
	// e47cfae9-c0cc-42e1-91e2-e67cdb0e7b96
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The stop reason.
	//
	// example:
	//
	// stop
	StopReason *string `json:"StopReason,omitempty" xml:"StopReason,omitempty"`
}

func (s ShoppingAssistantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ShoppingAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *ShoppingAssistantResponseBody) GetCitation() *ShoppingAssistantResponseBodyCitation {
	return s.Citation
}

func (s *ShoppingAssistantResponseBody) GetContent() *string {
	return s.Content
}

func (s *ShoppingAssistantResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *ShoppingAssistantResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ShoppingAssistantResponseBody) GetEvent() *string {
	return s.Event
}

func (s *ShoppingAssistantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ShoppingAssistantResponseBody) GetResult() *ShoppingAssistantResponseBodyResult {
	return s.Result
}

func (s *ShoppingAssistantResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *ShoppingAssistantResponseBody) GetStopReason() *string {
	return s.StopReason
}

func (s *ShoppingAssistantResponseBody) SetCitation(v *ShoppingAssistantResponseBodyCitation) *ShoppingAssistantResponseBody {
	s.Citation = v
	return s
}

func (s *ShoppingAssistantResponseBody) SetContent(v string) *ShoppingAssistantResponseBody {
	s.Content = &v
	return s
}

func (s *ShoppingAssistantResponseBody) SetConversationId(v string) *ShoppingAssistantResponseBody {
	s.ConversationId = &v
	return s
}

func (s *ShoppingAssistantResponseBody) SetErrorCode(v string) *ShoppingAssistantResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ShoppingAssistantResponseBody) SetEvent(v string) *ShoppingAssistantResponseBody {
	s.Event = &v
	return s
}

func (s *ShoppingAssistantResponseBody) SetRequestId(v string) *ShoppingAssistantResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShoppingAssistantResponseBody) SetResult(v *ShoppingAssistantResponseBodyResult) *ShoppingAssistantResponseBody {
	s.Result = v
	return s
}

func (s *ShoppingAssistantResponseBody) SetSessionId(v string) *ShoppingAssistantResponseBody {
	s.SessionId = &v
	return s
}

func (s *ShoppingAssistantResponseBody) SetStopReason(v string) *ShoppingAssistantResponseBody {
	s.StopReason = &v
	return s
}

func (s *ShoppingAssistantResponseBody) Validate() error {
	if s.Citation != nil {
		if err := s.Citation.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ShoppingAssistantResponseBodyCitation struct {
	// The ID of the `item`.
	//
	// example:
	//
	// 0005
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The reference data type. Fixed value: `item`.
	//
	// example:
	//
	// item
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ShoppingAssistantResponseBodyCitation) String() string {
	return dara.Prettify(s)
}

func (s ShoppingAssistantResponseBodyCitation) GoString() string {
	return s.String()
}

func (s *ShoppingAssistantResponseBodyCitation) GetItemId() *string {
	return s.ItemId
}

func (s *ShoppingAssistantResponseBodyCitation) GetType() *string {
	return s.Type
}

func (s *ShoppingAssistantResponseBodyCitation) SetItemId(v string) *ShoppingAssistantResponseBodyCitation {
	s.ItemId = &v
	return s
}

func (s *ShoppingAssistantResponseBodyCitation) SetType(v string) *ShoppingAssistantResponseBodyCitation {
	s.Type = &v
	return s
}

func (s *ShoppingAssistantResponseBodyCitation) Validate() error {
	return dara.Validate(s)
}

type ShoppingAssistantResponseBodyResult struct {
	// The citation information.
	Citation *ShoppingAssistantResponseBodyResultCitation `json:"Citation,omitempty" xml:"Citation,omitempty" type:"Struct"`
	// The returned content.
	//
	// example:
	//
	// Here are some light-colored long-sleeve shirts I picked for you:\\\\n.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The step information.
	StepInfo *ShoppingAssistantResponseBodyResultStepInfo `json:"StepInfo,omitempty" xml:"StepInfo,omitempty" type:"Struct"`
	// The stop reason.
	//
	// example:
	//
	// stop
	StopReason *string `json:"StopReason,omitempty" xml:"StopReason,omitempty"`
}

func (s ShoppingAssistantResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ShoppingAssistantResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ShoppingAssistantResponseBodyResult) GetCitation() *ShoppingAssistantResponseBodyResultCitation {
	return s.Citation
}

func (s *ShoppingAssistantResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *ShoppingAssistantResponseBodyResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ShoppingAssistantResponseBodyResult) GetStepInfo() *ShoppingAssistantResponseBodyResultStepInfo {
	return s.StepInfo
}

func (s *ShoppingAssistantResponseBodyResult) GetStopReason() *string {
	return s.StopReason
}

func (s *ShoppingAssistantResponseBodyResult) SetCitation(v *ShoppingAssistantResponseBodyResultCitation) *ShoppingAssistantResponseBodyResult {
	s.Citation = v
	return s
}

func (s *ShoppingAssistantResponseBodyResult) SetContent(v string) *ShoppingAssistantResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ShoppingAssistantResponseBodyResult) SetErrorCode(v string) *ShoppingAssistantResponseBodyResult {
	s.ErrorCode = &v
	return s
}

func (s *ShoppingAssistantResponseBodyResult) SetStepInfo(v *ShoppingAssistantResponseBodyResultStepInfo) *ShoppingAssistantResponseBodyResult {
	s.StepInfo = v
	return s
}

func (s *ShoppingAssistantResponseBodyResult) SetStopReason(v string) *ShoppingAssistantResponseBodyResult {
	s.StopReason = &v
	return s
}

func (s *ShoppingAssistantResponseBodyResult) Validate() error {
	if s.Citation != nil {
		if err := s.Citation.Validate(); err != nil {
			return err
		}
	}
	if s.StepInfo != nil {
		if err := s.StepInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ShoppingAssistantResponseBodyResultCitation struct {
	// The ID of the item.
	//
	// example:
	//
	// 48
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The reference data type. Fixed value: item.
	//
	// example:
	//
	// item
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ShoppingAssistantResponseBodyResultCitation) String() string {
	return dara.Prettify(s)
}

func (s ShoppingAssistantResponseBodyResultCitation) GoString() string {
	return s.String()
}

func (s *ShoppingAssistantResponseBodyResultCitation) GetItemId() *string {
	return s.ItemId
}

func (s *ShoppingAssistantResponseBodyResultCitation) GetType() *string {
	return s.Type
}

func (s *ShoppingAssistantResponseBodyResultCitation) SetItemId(v string) *ShoppingAssistantResponseBodyResultCitation {
	s.ItemId = &v
	return s
}

func (s *ShoppingAssistantResponseBodyResultCitation) SetType(v string) *ShoppingAssistantResponseBodyResultCitation {
	s.Type = &v
	return s
}

func (s *ShoppingAssistantResponseBodyResultCitation) Validate() error {
	return dara.Validate(s)
}

type ShoppingAssistantResponseBodyResultStepInfo struct {
	// The step.
	//
	// example:
	//
	// analyze_requirement
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s ShoppingAssistantResponseBodyResultStepInfo) String() string {
	return dara.Prettify(s)
}

func (s ShoppingAssistantResponseBodyResultStepInfo) GoString() string {
	return s.String()
}

func (s *ShoppingAssistantResponseBodyResultStepInfo) GetStep() *string {
	return s.Step
}

func (s *ShoppingAssistantResponseBodyResultStepInfo) SetStep(v string) *ShoppingAssistantResponseBodyResultStepInfo {
	s.Step = &v
	return s
}

func (s *ShoppingAssistantResponseBodyResultStepInfo) Validate() error {
	return dara.Validate(s)
}
