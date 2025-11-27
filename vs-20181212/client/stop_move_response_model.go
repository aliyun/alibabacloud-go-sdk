// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMoveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopMoveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopMoveResponse
	GetStatusCode() *int32
	SetBody(v *StopMoveResponseBody) *StopMoveResponse
	GetBody() *StopMoveResponseBody
}

type StopMoveResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopMoveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopMoveResponse) String() string {
	return dara.Prettify(s)
}

func (s StopMoveResponse) GoString() string {
	return s.String()
}

func (s *StopMoveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopMoveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopMoveResponse) GetBody() *StopMoveResponseBody {
	return s.Body
}

func (s *StopMoveResponse) SetHeaders(v map[string]*string) *StopMoveResponse {
	s.Headers = v
	return s
}

func (s *StopMoveResponse) SetStatusCode(v int32) *StopMoveResponse {
	s.StatusCode = &v
	return s
}

func (s *StopMoveResponse) SetBody(v *StopMoveResponseBody) *StopMoveResponse {
	s.Body = v
	return s
}

func (s *StopMoveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
