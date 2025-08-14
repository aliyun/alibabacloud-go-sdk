// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLiveTranscodeJobCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendLiveTranscodeJobCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendLiveTranscodeJobCommandResponse
	GetStatusCode() *int32
	SetBody(v *SendLiveTranscodeJobCommandResponseBody) *SendLiveTranscodeJobCommandResponse
	GetBody() *SendLiveTranscodeJobCommandResponseBody
}

type SendLiveTranscodeJobCommandResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendLiveTranscodeJobCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendLiveTranscodeJobCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s SendLiveTranscodeJobCommandResponse) GoString() string {
	return s.String()
}

func (s *SendLiveTranscodeJobCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendLiveTranscodeJobCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendLiveTranscodeJobCommandResponse) GetBody() *SendLiveTranscodeJobCommandResponseBody {
	return s.Body
}

func (s *SendLiveTranscodeJobCommandResponse) SetHeaders(v map[string]*string) *SendLiveTranscodeJobCommandResponse {
	s.Headers = v
	return s
}

func (s *SendLiveTranscodeJobCommandResponse) SetStatusCode(v int32) *SendLiveTranscodeJobCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *SendLiveTranscodeJobCommandResponse) SetBody(v *SendLiveTranscodeJobCommandResponseBody) *SendLiveTranscodeJobCommandResponse {
	s.Body = v
	return s
}

func (s *SendLiveTranscodeJobCommandResponse) Validate() error {
	return dara.Validate(s)
}
