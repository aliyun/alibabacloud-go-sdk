// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseInstanceResponse
	GetStatusCode() *int32
	SetBody(v *PauseInstanceResponseBody) *PauseInstanceResponse
	GetBody() *PauseInstanceResponseBody
}

type PauseInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseInstanceResponse) GoString() string {
	return s.String()
}

func (s *PauseInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseInstanceResponse) GetBody() *PauseInstanceResponseBody {
	return s.Body
}

func (s *PauseInstanceResponse) SetHeaders(v map[string]*string) *PauseInstanceResponse {
	s.Headers = v
	return s
}

func (s *PauseInstanceResponse) SetStatusCode(v int32) *PauseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseInstanceResponse) SetBody(v *PauseInstanceResponseBody) *PauseInstanceResponse {
	s.Body = v
	return s
}

func (s *PauseInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
