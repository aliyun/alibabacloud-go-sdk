// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBeginDialogueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *BeginDialogueResponseBody
	GetAction() *string
	SetActionParams(v string) *BeginDialogueResponseBody
	GetActionParams() *string
	SetInterruptible(v bool) *BeginDialogueResponseBody
	GetInterruptible() *bool
	SetRequestId(v string) *BeginDialogueResponseBody
	GetRequestId() *string
	SetTextResponse(v string) *BeginDialogueResponseBody
	GetTextResponse() *string
}

type BeginDialogueResponseBody struct {
	// example:
	//
	// Broadcast
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// {\\"duration\\":31340,\\"endTime\\":1638243934786,\\"hangUpDirection\\":\\"ivr\\",\\"startTime\\":1638243903446}
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

func (s BeginDialogueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BeginDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *BeginDialogueResponseBody) GetAction() *string {
	return s.Action
}

func (s *BeginDialogueResponseBody) GetActionParams() *string {
	return s.ActionParams
}

func (s *BeginDialogueResponseBody) GetInterruptible() *bool {
	return s.Interruptible
}

func (s *BeginDialogueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BeginDialogueResponseBody) GetTextResponse() *string {
	return s.TextResponse
}

func (s *BeginDialogueResponseBody) SetAction(v string) *BeginDialogueResponseBody {
	s.Action = &v
	return s
}

func (s *BeginDialogueResponseBody) SetActionParams(v string) *BeginDialogueResponseBody {
	s.ActionParams = &v
	return s
}

func (s *BeginDialogueResponseBody) SetInterruptible(v bool) *BeginDialogueResponseBody {
	s.Interruptible = &v
	return s
}

func (s *BeginDialogueResponseBody) SetRequestId(v string) *BeginDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *BeginDialogueResponseBody) SetTextResponse(v string) *BeginDialogueResponseBody {
	s.TextResponse = &v
	return s
}

func (s *BeginDialogueResponseBody) Validate() error {
	return dara.Validate(s)
}
