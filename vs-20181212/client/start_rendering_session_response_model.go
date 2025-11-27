// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRenderingSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRenderingSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRenderingSessionResponse
	GetStatusCode() *int32
	SetBody(v *StartRenderingSessionResponseBody) *StartRenderingSessionResponse
	GetBody() *StartRenderingSessionResponseBody
}

type StartRenderingSessionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRenderingSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRenderingSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRenderingSessionResponse) GoString() string {
	return s.String()
}

func (s *StartRenderingSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRenderingSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRenderingSessionResponse) GetBody() *StartRenderingSessionResponseBody {
	return s.Body
}

func (s *StartRenderingSessionResponse) SetHeaders(v map[string]*string) *StartRenderingSessionResponse {
	s.Headers = v
	return s
}

func (s *StartRenderingSessionResponse) SetStatusCode(v int32) *StartRenderingSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRenderingSessionResponse) SetBody(v *StartRenderingSessionResponseBody) *StartRenderingSessionResponse {
	s.Body = v
	return s
}

func (s *StartRenderingSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
