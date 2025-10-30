// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckProjectHasDependencyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckProjectHasDependencyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckProjectHasDependencyResponse
	GetStatusCode() *int32
	SetBody(v *CheckProjectHasDependencyResponseBody) *CheckProjectHasDependencyResponse
	GetBody() *CheckProjectHasDependencyResponseBody
}

type CheckProjectHasDependencyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckProjectHasDependencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckProjectHasDependencyResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckProjectHasDependencyResponse) GoString() string {
	return s.String()
}

func (s *CheckProjectHasDependencyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckProjectHasDependencyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckProjectHasDependencyResponse) GetBody() *CheckProjectHasDependencyResponseBody {
	return s.Body
}

func (s *CheckProjectHasDependencyResponse) SetHeaders(v map[string]*string) *CheckProjectHasDependencyResponse {
	s.Headers = v
	return s
}

func (s *CheckProjectHasDependencyResponse) SetStatusCode(v int32) *CheckProjectHasDependencyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckProjectHasDependencyResponse) SetBody(v *CheckProjectHasDependencyResponseBody) *CheckProjectHasDependencyResponse {
	s.Body = v
	return s
}

func (s *CheckProjectHasDependencyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
