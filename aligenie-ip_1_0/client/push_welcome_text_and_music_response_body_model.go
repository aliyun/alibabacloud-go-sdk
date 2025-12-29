// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushWelcomeTextAndMusicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *PushWelcomeTextAndMusicResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *PushWelcomeTextAndMusicResponseBody
	GetMessage() *string
	SetRequestId(v string) *PushWelcomeTextAndMusicResponseBody
	GetRequestId() *string
	SetResult(v bool) *PushWelcomeTextAndMusicResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *PushWelcomeTextAndMusicResponseBody
	GetStatusCode() *int32
}

type PushWelcomeTextAndMusicResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F7E2****B7C94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s PushWelcomeTextAndMusicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushWelcomeTextAndMusicResponseBody) GoString() string {
	return s.String()
}

func (s *PushWelcomeTextAndMusicResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *PushWelcomeTextAndMusicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PushWelcomeTextAndMusicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushWelcomeTextAndMusicResponseBody) GetResult() *bool {
	return s.Result
}

func (s *PushWelcomeTextAndMusicResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushWelcomeTextAndMusicResponseBody) SetExtentions(v map[string]interface{}) *PushWelcomeTextAndMusicResponseBody {
	s.Extentions = v
	return s
}

func (s *PushWelcomeTextAndMusicResponseBody) SetMessage(v string) *PushWelcomeTextAndMusicResponseBody {
	s.Message = &v
	return s
}

func (s *PushWelcomeTextAndMusicResponseBody) SetRequestId(v string) *PushWelcomeTextAndMusicResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushWelcomeTextAndMusicResponseBody) SetResult(v bool) *PushWelcomeTextAndMusicResponseBody {
	s.Result = &v
	return s
}

func (s *PushWelcomeTextAndMusicResponseBody) SetStatusCode(v int32) *PushWelcomeTextAndMusicResponseBody {
	s.StatusCode = &v
	return s
}

func (s *PushWelcomeTextAndMusicResponseBody) Validate() error {
	return dara.Validate(s)
}
