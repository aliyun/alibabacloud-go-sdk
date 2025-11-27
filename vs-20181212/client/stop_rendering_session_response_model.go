// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRenderingSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopRenderingSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopRenderingSessionResponse
	GetStatusCode() *int32
	SetBody(v *StopRenderingSessionResponseBody) *StopRenderingSessionResponse
	GetBody() *StopRenderingSessionResponseBody
}

type StopRenderingSessionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRenderingSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRenderingSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s StopRenderingSessionResponse) GoString() string {
	return s.String()
}

func (s *StopRenderingSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopRenderingSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopRenderingSessionResponse) GetBody() *StopRenderingSessionResponseBody {
	return s.Body
}

func (s *StopRenderingSessionResponse) SetHeaders(v map[string]*string) *StopRenderingSessionResponse {
	s.Headers = v
	return s
}

func (s *StopRenderingSessionResponse) SetStatusCode(v int32) *StopRenderingSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRenderingSessionResponse) SetBody(v *StopRenderingSessionResponseBody) *StopRenderingSessionResponse {
	s.Body = v
	return s
}

func (s *StopRenderingSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
