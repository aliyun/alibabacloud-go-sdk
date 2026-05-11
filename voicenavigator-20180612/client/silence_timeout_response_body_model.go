// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSilenceTimeoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *SilenceTimeoutResponseBody
	GetAction() *string
	SetActionParams(v string) *SilenceTimeoutResponseBody
	GetActionParams() *string
	SetInterruptible(v bool) *SilenceTimeoutResponseBody
	GetInterruptible() *bool
	SetRequestId(v string) *SilenceTimeoutResponseBody
	GetRequestId() *string
	SetTextResponse(v string) *SilenceTimeoutResponseBody
	GetTextResponse() *string
}

type SilenceTimeoutResponseBody struct {
	// example:
	//
	// TransferToAgent
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// { "skillGroupId": "ABC"}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// example:
	//
	// false
	Interruptible *bool `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TextResponse *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
}

func (s SilenceTimeoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SilenceTimeoutResponseBody) GoString() string {
	return s.String()
}

func (s *SilenceTimeoutResponseBody) GetAction() *string {
	return s.Action
}

func (s *SilenceTimeoutResponseBody) GetActionParams() *string {
	return s.ActionParams
}

func (s *SilenceTimeoutResponseBody) GetInterruptible() *bool {
	return s.Interruptible
}

func (s *SilenceTimeoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SilenceTimeoutResponseBody) GetTextResponse() *string {
	return s.TextResponse
}

func (s *SilenceTimeoutResponseBody) SetAction(v string) *SilenceTimeoutResponseBody {
	s.Action = &v
	return s
}

func (s *SilenceTimeoutResponseBody) SetActionParams(v string) *SilenceTimeoutResponseBody {
	s.ActionParams = &v
	return s
}

func (s *SilenceTimeoutResponseBody) SetInterruptible(v bool) *SilenceTimeoutResponseBody {
	s.Interruptible = &v
	return s
}

func (s *SilenceTimeoutResponseBody) SetRequestId(v string) *SilenceTimeoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *SilenceTimeoutResponseBody) SetTextResponse(v string) *SilenceTimeoutResponseBody {
	s.TextResponse = &v
	return s
}

func (s *SilenceTimeoutResponseBody) Validate() error {
	return dara.Validate(s)
}
