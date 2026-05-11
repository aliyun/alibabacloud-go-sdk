// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugCollectedNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *DebugCollectedNumberResponseBody
	GetAction() *string
	SetActionParams(v string) *DebugCollectedNumberResponseBody
	GetActionParams() *string
	SetInterruptible(v bool) *DebugCollectedNumberResponseBody
	GetInterruptible() *bool
	SetRequestId(v string) *DebugCollectedNumberResponseBody
	GetRequestId() *string
	SetTextResponse(v string) *DebugCollectedNumberResponseBody
	GetTextResponse() *string
}

type DebugCollectedNumberResponseBody struct {
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
	// abb4aa26-3a8e-43dd-82f8-0c3898c9c67f
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TextResponse *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
}

func (s DebugCollectedNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DebugCollectedNumberResponseBody) GoString() string {
	return s.String()
}

func (s *DebugCollectedNumberResponseBody) GetAction() *string {
	return s.Action
}

func (s *DebugCollectedNumberResponseBody) GetActionParams() *string {
	return s.ActionParams
}

func (s *DebugCollectedNumberResponseBody) GetInterruptible() *bool {
	return s.Interruptible
}

func (s *DebugCollectedNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DebugCollectedNumberResponseBody) GetTextResponse() *string {
	return s.TextResponse
}

func (s *DebugCollectedNumberResponseBody) SetAction(v string) *DebugCollectedNumberResponseBody {
	s.Action = &v
	return s
}

func (s *DebugCollectedNumberResponseBody) SetActionParams(v string) *DebugCollectedNumberResponseBody {
	s.ActionParams = &v
	return s
}

func (s *DebugCollectedNumberResponseBody) SetInterruptible(v bool) *DebugCollectedNumberResponseBody {
	s.Interruptible = &v
	return s
}

func (s *DebugCollectedNumberResponseBody) SetRequestId(v string) *DebugCollectedNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugCollectedNumberResponseBody) SetTextResponse(v string) *DebugCollectedNumberResponseBody {
	s.TextResponse = &v
	return s
}

func (s *DebugCollectedNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
