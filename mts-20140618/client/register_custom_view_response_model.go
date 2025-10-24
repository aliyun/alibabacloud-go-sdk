// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterCustomViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterCustomViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterCustomViewResponse
	GetStatusCode() *int32
	SetBody(v *RegisterCustomViewResponseBody) *RegisterCustomViewResponse
	GetBody() *RegisterCustomViewResponseBody
}

type RegisterCustomViewResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterCustomViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterCustomViewResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterCustomViewResponse) GoString() string {
	return s.String()
}

func (s *RegisterCustomViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterCustomViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterCustomViewResponse) GetBody() *RegisterCustomViewResponseBody {
	return s.Body
}

func (s *RegisterCustomViewResponse) SetHeaders(v map[string]*string) *RegisterCustomViewResponse {
	s.Headers = v
	return s
}

func (s *RegisterCustomViewResponse) SetStatusCode(v int32) *RegisterCustomViewResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterCustomViewResponse) SetBody(v *RegisterCustomViewResponseBody) *RegisterCustomViewResponse {
	s.Body = v
	return s
}

func (s *RegisterCustomViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
