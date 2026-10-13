// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVoiceCloneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVoiceCloneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVoiceCloneResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVoiceCloneResponseBody) *DeleteVoiceCloneResponse
	GetBody() *DeleteVoiceCloneResponseBody
}

type DeleteVoiceCloneResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVoiceCloneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVoiceCloneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVoiceCloneResponse) GoString() string {
	return s.String()
}

func (s *DeleteVoiceCloneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVoiceCloneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVoiceCloneResponse) GetBody() *DeleteVoiceCloneResponseBody {
	return s.Body
}

func (s *DeleteVoiceCloneResponse) SetHeaders(v map[string]*string) *DeleteVoiceCloneResponse {
	s.Headers = v
	return s
}

func (s *DeleteVoiceCloneResponse) SetStatusCode(v int32) *DeleteVoiceCloneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVoiceCloneResponse) SetBody(v *DeleteVoiceCloneResponseBody) *DeleteVoiceCloneResponse {
	s.Body = v
	return s
}

func (s *DeleteVoiceCloneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
