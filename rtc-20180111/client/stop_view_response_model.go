// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopViewResponse
	GetStatusCode() *int32
	SetBody(v *StopViewResponseBody) *StopViewResponse
	GetBody() *StopViewResponseBody
}

type StopViewResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopViewResponse) String() string {
	return dara.Prettify(s)
}

func (s StopViewResponse) GoString() string {
	return s.String()
}

func (s *StopViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopViewResponse) GetBody() *StopViewResponseBody {
	return s.Body
}

func (s *StopViewResponse) SetHeaders(v map[string]*string) *StopViewResponse {
	s.Headers = v
	return s
}

func (s *StopViewResponse) SetStatusCode(v int32) *StopViewResponse {
	s.StatusCode = &v
	return s
}

func (s *StopViewResponse) SetBody(v *StopViewResponseBody) *StopViewResponse {
	s.Body = v
	return s
}

func (s *StopViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
