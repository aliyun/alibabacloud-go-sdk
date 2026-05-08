// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckSessionResponse
	GetStatusCode() *int32
	SetBody(v *CheckSessionResponseBody) *CheckSessionResponse
	GetBody() *CheckSessionResponseBody
}

type CheckSessionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckSessionResponse) GoString() string {
	return s.String()
}

func (s *CheckSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckSessionResponse) GetBody() *CheckSessionResponseBody {
	return s.Body
}

func (s *CheckSessionResponse) SetHeaders(v map[string]*string) *CheckSessionResponse {
	s.Headers = v
	return s
}

func (s *CheckSessionResponse) SetStatusCode(v int32) *CheckSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSessionResponse) SetBody(v *CheckSessionResponseBody) *CheckSessionResponse {
	s.Body = v
	return s
}

func (s *CheckSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
