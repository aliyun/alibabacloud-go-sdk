// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranscribeChatVoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TranscribeChatVoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TranscribeChatVoiceResponse
	GetStatusCode() *int32
	SetBody(v *TranscribeChatVoiceResponseBody) *TranscribeChatVoiceResponse
	GetBody() *TranscribeChatVoiceResponseBody
}

type TranscribeChatVoiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranscribeChatVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TranscribeChatVoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s TranscribeChatVoiceResponse) GoString() string {
	return s.String()
}

func (s *TranscribeChatVoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TranscribeChatVoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TranscribeChatVoiceResponse) GetBody() *TranscribeChatVoiceResponseBody {
	return s.Body
}

func (s *TranscribeChatVoiceResponse) SetHeaders(v map[string]*string) *TranscribeChatVoiceResponse {
	s.Headers = v
	return s
}

func (s *TranscribeChatVoiceResponse) SetStatusCode(v int32) *TranscribeChatVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *TranscribeChatVoiceResponse) SetBody(v *TranscribeChatVoiceResponseBody) *TranscribeChatVoiceResponse {
	s.Body = v
	return s
}

func (s *TranscribeChatVoiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
