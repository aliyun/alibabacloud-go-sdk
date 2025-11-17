// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceAdbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopInstanceAdbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopInstanceAdbResponse
	GetStatusCode() *int32
	SetBody(v *StopInstanceAdbResponseBody) *StopInstanceAdbResponse
	GetBody() *StopInstanceAdbResponseBody
}

type StopInstanceAdbResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopInstanceAdbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopInstanceAdbResponse) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceAdbResponse) GoString() string {
	return s.String()
}

func (s *StopInstanceAdbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopInstanceAdbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopInstanceAdbResponse) GetBody() *StopInstanceAdbResponseBody {
	return s.Body
}

func (s *StopInstanceAdbResponse) SetHeaders(v map[string]*string) *StopInstanceAdbResponse {
	s.Headers = v
	return s
}

func (s *StopInstanceAdbResponse) SetStatusCode(v int32) *StopInstanceAdbResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInstanceAdbResponse) SetBody(v *StopInstanceAdbResponseBody) *StopInstanceAdbResponse {
	s.Body = v
	return s
}

func (s *StopInstanceAdbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
