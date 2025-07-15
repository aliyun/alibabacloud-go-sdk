// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAISubtitleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveAISubtitleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveAISubtitleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveAISubtitleResponseBody) *DeleteLiveAISubtitleResponse
	GetBody() *DeleteLiveAISubtitleResponseBody
}

type DeleteLiveAISubtitleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveAISubtitleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveAISubtitleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAISubtitleResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveAISubtitleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveAISubtitleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveAISubtitleResponse) GetBody() *DeleteLiveAISubtitleResponseBody {
	return s.Body
}

func (s *DeleteLiveAISubtitleResponse) SetHeaders(v map[string]*string) *DeleteLiveAISubtitleResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveAISubtitleResponse) SetStatusCode(v int32) *DeleteLiveAISubtitleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveAISubtitleResponse) SetBody(v *DeleteLiveAISubtitleResponseBody) *DeleteLiveAISubtitleResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveAISubtitleResponse) Validate() error {
	return dara.Validate(s)
}
