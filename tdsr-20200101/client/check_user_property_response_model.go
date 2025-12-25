// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserPropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckUserPropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckUserPropertyResponse
	GetStatusCode() *int32
	SetBody(v *CheckUserPropertyResponseBody) *CheckUserPropertyResponse
	GetBody() *CheckUserPropertyResponseBody
}

type CheckUserPropertyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUserPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUserPropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckUserPropertyResponse) GoString() string {
	return s.String()
}

func (s *CheckUserPropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckUserPropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckUserPropertyResponse) GetBody() *CheckUserPropertyResponseBody {
	return s.Body
}

func (s *CheckUserPropertyResponse) SetHeaders(v map[string]*string) *CheckUserPropertyResponse {
	s.Headers = v
	return s
}

func (s *CheckUserPropertyResponse) SetStatusCode(v int32) *CheckUserPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUserPropertyResponse) SetBody(v *CheckUserPropertyResponseBody) *CheckUserPropertyResponse {
	s.Body = v
	return s
}

func (s *CheckUserPropertyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
