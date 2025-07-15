// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAISubtitleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveAISubtitleResponseBody
	GetRequestId() *string
	SetSubtitleId(v string) *AddLiveAISubtitleResponseBody
	GetSubtitleId() *string
}

type AddLiveAISubtitleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0df228-4a64-af62-20e91b96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the subtitle template.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	SubtitleId *string `json:"SubtitleId,omitempty" xml:"SubtitleId,omitempty"`
}

func (s AddLiveAISubtitleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAISubtitleResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveAISubtitleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveAISubtitleResponseBody) GetSubtitleId() *string {
	return s.SubtitleId
}

func (s *AddLiveAISubtitleResponseBody) SetRequestId(v string) *AddLiveAISubtitleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveAISubtitleResponseBody) SetSubtitleId(v string) *AddLiveAISubtitleResponseBody {
	s.SubtitleId = &v
	return s
}

func (s *AddLiveAISubtitleResponseBody) Validate() error {
	return dara.Validate(s)
}
