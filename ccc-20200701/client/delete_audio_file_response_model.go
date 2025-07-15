// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAudioFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAudioFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAudioFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAudioFileResponseBody) *DeleteAudioFileResponse
	GetBody() *DeleteAudioFileResponseBody
}

type DeleteAudioFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAudioFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAudioFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAudioFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteAudioFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAudioFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAudioFileResponse) GetBody() *DeleteAudioFileResponseBody {
	return s.Body
}

func (s *DeleteAudioFileResponse) SetHeaders(v map[string]*string) *DeleteAudioFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteAudioFileResponse) SetStatusCode(v int32) *DeleteAudioFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAudioFileResponse) SetBody(v *DeleteAudioFileResponseBody) *DeleteAudioFileResponse {
	s.Body = v
	return s
}

func (s *DeleteAudioFileResponse) Validate() error {
	return dara.Validate(s)
}
