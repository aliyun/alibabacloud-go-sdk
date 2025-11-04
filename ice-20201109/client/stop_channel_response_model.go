// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopChannelResponse
	GetStatusCode() *int32
	SetBody(v *StopChannelResponseBody) *StopChannelResponse
	GetBody() *StopChannelResponseBody
}

type StopChannelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s StopChannelResponse) GoString() string {
	return s.String()
}

func (s *StopChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopChannelResponse) GetBody() *StopChannelResponseBody {
	return s.Body
}

func (s *StopChannelResponse) SetHeaders(v map[string]*string) *StopChannelResponse {
	s.Headers = v
	return s
}

func (s *StopChannelResponse) SetStatusCode(v int32) *StopChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *StopChannelResponse) SetBody(v *StopChannelResponseBody) *StopChannelResponse {
	s.Body = v
	return s
}

func (s *StopChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
