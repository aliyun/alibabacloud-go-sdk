// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitialSysomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitialSysomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitialSysomResponse
	GetStatusCode() *int32
	SetBody(v *InitialSysomResponseBody) *InitialSysomResponse
	GetBody() *InitialSysomResponseBody
}

type InitialSysomResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitialSysomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitialSysomResponse) String() string {
	return dara.Prettify(s)
}

func (s InitialSysomResponse) GoString() string {
	return s.String()
}

func (s *InitialSysomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitialSysomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitialSysomResponse) GetBody() *InitialSysomResponseBody {
	return s.Body
}

func (s *InitialSysomResponse) SetHeaders(v map[string]*string) *InitialSysomResponse {
	s.Headers = v
	return s
}

func (s *InitialSysomResponse) SetStatusCode(v int32) *InitialSysomResponse {
	s.StatusCode = &v
	return s
}

func (s *InitialSysomResponse) SetBody(v *InitialSysomResponseBody) *InitialSysomResponse {
	s.Body = v
	return s
}

func (s *InitialSysomResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
