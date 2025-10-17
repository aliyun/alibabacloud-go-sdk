// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopStreamingOutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopStreamingOutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopStreamingOutResponse
	GetStatusCode() *int32
	SetBody(v *StopStreamingOutResponseBody) *StopStreamingOutResponse
	GetBody() *StopStreamingOutResponseBody
}

type StopStreamingOutResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopStreamingOutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopStreamingOutResponse) String() string {
	return dara.Prettify(s)
}

func (s StopStreamingOutResponse) GoString() string {
	return s.String()
}

func (s *StopStreamingOutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopStreamingOutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopStreamingOutResponse) GetBody() *StopStreamingOutResponseBody {
	return s.Body
}

func (s *StopStreamingOutResponse) SetHeaders(v map[string]*string) *StopStreamingOutResponse {
	s.Headers = v
	return s
}

func (s *StopStreamingOutResponse) SetStatusCode(v int32) *StopStreamingOutResponse {
	s.StatusCode = &v
	return s
}

func (s *StopStreamingOutResponse) SetBody(v *StopStreamingOutResponseBody) *StopStreamingOutResponse {
	s.Body = v
	return s
}

func (s *StopStreamingOutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
