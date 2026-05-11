// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCollectedNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *CollectedNumberResponseBody
	GetAction() *string
	SetActionParams(v string) *CollectedNumberResponseBody
	GetActionParams() *string
	SetInterruptible(v bool) *CollectedNumberResponseBody
	GetInterruptible() *bool
	SetRequestId(v string) *CollectedNumberResponseBody
	GetRequestId() *string
	SetTextResponse(v string) *CollectedNumberResponseBody
	GetTextResponse() *string
}

type CollectedNumberResponseBody struct {
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
	// false
	Interruptible *bool `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TextResponse *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
}

func (s CollectedNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CollectedNumberResponseBody) GoString() string {
	return s.String()
}

func (s *CollectedNumberResponseBody) GetAction() *string {
	return s.Action
}

func (s *CollectedNumberResponseBody) GetActionParams() *string {
	return s.ActionParams
}

func (s *CollectedNumberResponseBody) GetInterruptible() *bool {
	return s.Interruptible
}

func (s *CollectedNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CollectedNumberResponseBody) GetTextResponse() *string {
	return s.TextResponse
}

func (s *CollectedNumberResponseBody) SetAction(v string) *CollectedNumberResponseBody {
	s.Action = &v
	return s
}

func (s *CollectedNumberResponseBody) SetActionParams(v string) *CollectedNumberResponseBody {
	s.ActionParams = &v
	return s
}

func (s *CollectedNumberResponseBody) SetInterruptible(v bool) *CollectedNumberResponseBody {
	s.Interruptible = &v
	return s
}

func (s *CollectedNumberResponseBody) SetRequestId(v string) *CollectedNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CollectedNumberResponseBody) SetTextResponse(v string) *CollectedNumberResponseBody {
	s.TextResponse = &v
	return s
}

func (s *CollectedNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
