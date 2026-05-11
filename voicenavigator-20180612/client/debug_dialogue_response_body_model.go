// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugDialogueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *DebugDialogueResponseBody
	GetAction() *string
	SetActionParams(v string) *DebugDialogueResponseBody
	GetActionParams() *string
	SetInterruptible(v bool) *DebugDialogueResponseBody
	GetInterruptible() *bool
	SetRequestId(v string) *DebugDialogueResponseBody
	GetRequestId() *string
	SetTextResponse(v string) *DebugDialogueResponseBody
	GetTextResponse() *string
}

type DebugDialogueResponseBody struct {
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
	// d74d6290-7cbe-4436-b5d7-014ebb0f4060
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 80d11be3-faad-4101-b4b0-59dbea28aaf0
	TextResponse *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
}

func (s DebugDialogueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DebugDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *DebugDialogueResponseBody) GetAction() *string {
	return s.Action
}

func (s *DebugDialogueResponseBody) GetActionParams() *string {
	return s.ActionParams
}

func (s *DebugDialogueResponseBody) GetInterruptible() *bool {
	return s.Interruptible
}

func (s *DebugDialogueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DebugDialogueResponseBody) GetTextResponse() *string {
	return s.TextResponse
}

func (s *DebugDialogueResponseBody) SetAction(v string) *DebugDialogueResponseBody {
	s.Action = &v
	return s
}

func (s *DebugDialogueResponseBody) SetActionParams(v string) *DebugDialogueResponseBody {
	s.ActionParams = &v
	return s
}

func (s *DebugDialogueResponseBody) SetInterruptible(v bool) *DebugDialogueResponseBody {
	s.Interruptible = &v
	return s
}

func (s *DebugDialogueResponseBody) SetRequestId(v string) *DebugDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugDialogueResponseBody) SetTextResponse(v string) *DebugDialogueResponseBody {
	s.TextResponse = &v
	return s
}

func (s *DebugDialogueResponseBody) Validate() error {
	return dara.Validate(s)
}
