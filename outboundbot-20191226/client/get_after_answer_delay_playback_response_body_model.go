// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAfterAnswerDelayPlaybackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAfterAnswerDelayPlayback(v int32) *GetAfterAnswerDelayPlaybackResponseBody
	GetAfterAnswerDelayPlayback() *int32
	SetCode(v string) *GetAfterAnswerDelayPlaybackResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetAfterAnswerDelayPlaybackResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAfterAnswerDelayPlaybackResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAfterAnswerDelayPlaybackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAfterAnswerDelayPlaybackResponseBody
	GetSuccess() *bool
}

type GetAfterAnswerDelayPlaybackResponseBody struct {
	// Latency of the response
	//
	// example:
	//
	// 1
	AfterAnswerDelayPlayback *int32 `json:"AfterAnswerDelayPlayback,omitempty" xml:"AfterAnswerDelayPlayback,omitempty"`
	// API status code
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
	// Response message
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
	// Indicates whether the request succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAfterAnswerDelayPlaybackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAfterAnswerDelayPlaybackResponseBody) GoString() string {
	return s.String()
}

func (s *GetAfterAnswerDelayPlaybackResponseBody) GetAfterAnswerDelayPlayback() *int32 {
	return s.AfterAnswerDelayPlayback
}

func (s *GetAfterAnswerDelayPlaybackResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAfterAnswerDelayPlaybackResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAfterAnswerDelayPlaybackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAfterAnswerDelayPlaybackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAfterAnswerDelayPlaybackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAfterAnswerDelayPlaybackResponseBody) SetAfterAnswerDelayPlayback(v int32) *GetAfterAnswerDelayPlaybackResponseBody {
	s.AfterAnswerDelayPlayback = &v
	return s
}

func (s *GetAfterAnswerDelayPlaybackResponseBody) SetCode(v string) *GetAfterAnswerDelayPlaybackResponseBody {
	s.Code = &v
	return s
}

func (s *GetAfterAnswerDelayPlaybackResponseBody) SetHttpStatusCode(v int32) *GetAfterAnswerDelayPlaybackResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAfterAnswerDelayPlaybackResponseBody) SetMessage(v string) *GetAfterAnswerDelayPlaybackResponseBody {
	s.Message = &v
	return s
}

func (s *GetAfterAnswerDelayPlaybackResponseBody) SetRequestId(v string) *GetAfterAnswerDelayPlaybackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAfterAnswerDelayPlaybackResponseBody) SetSuccess(v bool) *GetAfterAnswerDelayPlaybackResponseBody {
	s.Success = &v
	return s
}

func (s *GetAfterAnswerDelayPlaybackResponseBody) Validate() error {
	return dara.Validate(s)
}
