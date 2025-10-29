// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAISubtitleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveAISubtitleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveAISubtitleResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveAISubtitleResponseBody) *AddLiveAISubtitleResponse
	GetBody() *AddLiveAISubtitleResponseBody
}

type AddLiveAISubtitleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveAISubtitleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveAISubtitleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAISubtitleResponse) GoString() string {
	return s.String()
}

func (s *AddLiveAISubtitleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveAISubtitleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveAISubtitleResponse) GetBody() *AddLiveAISubtitleResponseBody {
	return s.Body
}

func (s *AddLiveAISubtitleResponse) SetHeaders(v map[string]*string) *AddLiveAISubtitleResponse {
	s.Headers = v
	return s
}

func (s *AddLiveAISubtitleResponse) SetStatusCode(v int32) *AddLiveAISubtitleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveAISubtitleResponse) SetBody(v *AddLiveAISubtitleResponseBody) *AddLiveAISubtitleResponse {
	s.Body = v
	return s
}

func (s *AddLiveAISubtitleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
