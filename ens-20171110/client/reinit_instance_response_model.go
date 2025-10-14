// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinitInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReinitInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReinitInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReinitInstanceResponseBody) *ReinitInstanceResponse
	GetBody() *ReinitInstanceResponseBody
}

type ReinitInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReinitInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReinitInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReinitInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReinitInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReinitInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReinitInstanceResponse) GetBody() *ReinitInstanceResponseBody {
	return s.Body
}

func (s *ReinitInstanceResponse) SetHeaders(v map[string]*string) *ReinitInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReinitInstanceResponse) SetStatusCode(v int32) *ReinitInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReinitInstanceResponse) SetBody(v *ReinitInstanceResponseBody) *ReinitInstanceResponse {
	s.Body = v
	return s
}

func (s *ReinitInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
