// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugBeginDialogueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *DebugBeginDialogueResponseBody
	GetAction() *string
	SetActionParams(v string) *DebugBeginDialogueResponseBody
	GetActionParams() *string
	SetInterruptible(v bool) *DebugBeginDialogueResponseBody
	GetInterruptible() *bool
	SetRequestId(v string) *DebugBeginDialogueResponseBody
	GetRequestId() *string
	SetTextResponse(v string) *DebugBeginDialogueResponseBody
	GetTextResponse() *string
}

type DebugBeginDialogueResponseBody struct {
	// example:
	//
	// Broadcast
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// {}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// example:
	//
	// true
	Interruptible *bool `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TextResponse *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
}

func (s DebugBeginDialogueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DebugBeginDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *DebugBeginDialogueResponseBody) GetAction() *string {
	return s.Action
}

func (s *DebugBeginDialogueResponseBody) GetActionParams() *string {
	return s.ActionParams
}

func (s *DebugBeginDialogueResponseBody) GetInterruptible() *bool {
	return s.Interruptible
}

func (s *DebugBeginDialogueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DebugBeginDialogueResponseBody) GetTextResponse() *string {
	return s.TextResponse
}

func (s *DebugBeginDialogueResponseBody) SetAction(v string) *DebugBeginDialogueResponseBody {
	s.Action = &v
	return s
}

func (s *DebugBeginDialogueResponseBody) SetActionParams(v string) *DebugBeginDialogueResponseBody {
	s.ActionParams = &v
	return s
}

func (s *DebugBeginDialogueResponseBody) SetInterruptible(v bool) *DebugBeginDialogueResponseBody {
	s.Interruptible = &v
	return s
}

func (s *DebugBeginDialogueResponseBody) SetRequestId(v string) *DebugBeginDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugBeginDialogueResponseBody) SetTextResponse(v string) *DebugBeginDialogueResponseBody {
	s.TextResponse = &v
	return s
}

func (s *DebugBeginDialogueResponseBody) Validate() error {
	return dara.Validate(s)
}
