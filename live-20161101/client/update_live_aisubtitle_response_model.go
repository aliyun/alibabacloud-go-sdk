// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAISubtitleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveAISubtitleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveAISubtitleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveAISubtitleResponseBody) *UpdateLiveAISubtitleResponse
	GetBody() *UpdateLiveAISubtitleResponseBody
}

type UpdateLiveAISubtitleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveAISubtitleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveAISubtitleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAISubtitleResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveAISubtitleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveAISubtitleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveAISubtitleResponse) GetBody() *UpdateLiveAISubtitleResponseBody {
	return s.Body
}

func (s *UpdateLiveAISubtitleResponse) SetHeaders(v map[string]*string) *UpdateLiveAISubtitleResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveAISubtitleResponse) SetStatusCode(v int32) *UpdateLiveAISubtitleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveAISubtitleResponse) SetBody(v *UpdateLiveAISubtitleResponseBody) *UpdateLiveAISubtitleResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveAISubtitleResponse) Validate() error {
	return dara.Validate(s)
}
