// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDialogueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *DialogueResponseBody
	GetAction() *string
	SetActionParams(v string) *DialogueResponseBody
	GetActionParams() *string
	SetInterruptible(v bool) *DialogueResponseBody
	GetInterruptible() *bool
	SetRequestId(v string) *DialogueResponseBody
	GetRequestId() *string
	SetTextResponse(v string) *DialogueResponseBody
	GetTextResponse() *string
}

type DialogueResponseBody struct {
	// example:
	//
	// Broadcast
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// {\\"duration\\":2420,\\"endTime\\":1651717326805,\\"hangUpDirection\\":\\"client\\",\\"hasLastPlaybackCompleted\\":true,\\"startTime\\":1651717324385}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// example:
	//
	// true
	Interruptible *bool `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	// example:
	//
	// 9DB8BA95-4513-54B9-9C67-A28909CFB4AD
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TextResponse *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
}

func (s DialogueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DialogueResponseBody) GoString() string {
	return s.String()
}

func (s *DialogueResponseBody) GetAction() *string {
	return s.Action
}

func (s *DialogueResponseBody) GetActionParams() *string {
	return s.ActionParams
}

func (s *DialogueResponseBody) GetInterruptible() *bool {
	return s.Interruptible
}

func (s *DialogueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DialogueResponseBody) GetTextResponse() *string {
	return s.TextResponse
}

func (s *DialogueResponseBody) SetAction(v string) *DialogueResponseBody {
	s.Action = &v
	return s
}

func (s *DialogueResponseBody) SetActionParams(v string) *DialogueResponseBody {
	s.ActionParams = &v
	return s
}

func (s *DialogueResponseBody) SetInterruptible(v bool) *DialogueResponseBody {
	s.Interruptible = &v
	return s
}

func (s *DialogueResponseBody) SetRequestId(v string) *DialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DialogueResponseBody) SetTextResponse(v string) *DialogueResponseBody {
	s.TextResponse = &v
	return s
}

func (s *DialogueResponseBody) Validate() error {
	return dara.Validate(s)
}
