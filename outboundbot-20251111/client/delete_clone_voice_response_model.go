// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloneVoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloneVoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloneVoiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloneVoiceResponseBody) *DeleteCloneVoiceResponse
	GetBody() *DeleteCloneVoiceResponseBody
}

type DeleteCloneVoiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloneVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloneVoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloneVoiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloneVoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloneVoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloneVoiceResponse) GetBody() *DeleteCloneVoiceResponseBody {
	return s.Body
}

func (s *DeleteCloneVoiceResponse) SetHeaders(v map[string]*string) *DeleteCloneVoiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloneVoiceResponse) SetStatusCode(v int32) *DeleteCloneVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloneVoiceResponse) SetBody(v *DeleteCloneVoiceResponseBody) *DeleteCloneVoiceResponse {
	s.Body = v
	return s
}

func (s *DeleteCloneVoiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
