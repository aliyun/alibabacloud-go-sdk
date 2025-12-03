// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopEaisEiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopEaisEiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopEaisEiResponse
	GetStatusCode() *int32
	SetBody(v *StopEaisEiResponseBody) *StopEaisEiResponse
	GetBody() *StopEaisEiResponseBody
}

type StopEaisEiResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopEaisEiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopEaisEiResponse) String() string {
	return dara.Prettify(s)
}

func (s StopEaisEiResponse) GoString() string {
	return s.String()
}

func (s *StopEaisEiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopEaisEiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopEaisEiResponse) GetBody() *StopEaisEiResponseBody {
	return s.Body
}

func (s *StopEaisEiResponse) SetHeaders(v map[string]*string) *StopEaisEiResponse {
	s.Headers = v
	return s
}

func (s *StopEaisEiResponse) SetStatusCode(v int32) *StopEaisEiResponse {
	s.StatusCode = &v
	return s
}

func (s *StopEaisEiResponse) SetBody(v *StopEaisEiResponseBody) *StopEaisEiResponse {
	s.Body = v
	return s
}

func (s *StopEaisEiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
