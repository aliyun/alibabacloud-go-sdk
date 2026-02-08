// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAfterAnswerDelayPlaybackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveAfterAnswerDelayPlaybackResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SaveAfterAnswerDelayPlaybackResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveAfterAnswerDelayPlaybackResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveAfterAnswerDelayPlaybackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveAfterAnswerDelayPlaybackResponseBody
	GetSuccess() *bool
}

type SaveAfterAnswerDelayPlaybackResponseBody struct {
	// Response code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveAfterAnswerDelayPlaybackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveAfterAnswerDelayPlaybackResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAfterAnswerDelayPlaybackResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveAfterAnswerDelayPlaybackResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveAfterAnswerDelayPlaybackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveAfterAnswerDelayPlaybackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveAfterAnswerDelayPlaybackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveAfterAnswerDelayPlaybackResponseBody) SetCode(v string) *SaveAfterAnswerDelayPlaybackResponseBody {
	s.Code = &v
	return s
}

func (s *SaveAfterAnswerDelayPlaybackResponseBody) SetHttpStatusCode(v int32) *SaveAfterAnswerDelayPlaybackResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveAfterAnswerDelayPlaybackResponseBody) SetMessage(v string) *SaveAfterAnswerDelayPlaybackResponseBody {
	s.Message = &v
	return s
}

func (s *SaveAfterAnswerDelayPlaybackResponseBody) SetRequestId(v string) *SaveAfterAnswerDelayPlaybackResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveAfterAnswerDelayPlaybackResponseBody) SetSuccess(v bool) *SaveAfterAnswerDelayPlaybackResponseBody {
	s.Success = &v
	return s
}

func (s *SaveAfterAnswerDelayPlaybackResponseBody) Validate() error {
	return dara.Validate(s)
}
