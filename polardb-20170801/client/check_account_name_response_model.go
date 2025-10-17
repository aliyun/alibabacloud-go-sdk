// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAccountNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAccountNameResponse
	GetStatusCode() *int32
	SetBody(v *CheckAccountNameResponseBody) *CheckAccountNameResponse
	GetBody() *CheckAccountNameResponseBody
}

type CheckAccountNameResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAccountNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAccountNameResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountNameResponse) GoString() string {
	return s.String()
}

func (s *CheckAccountNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAccountNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAccountNameResponse) GetBody() *CheckAccountNameResponseBody {
	return s.Body
}

func (s *CheckAccountNameResponse) SetHeaders(v map[string]*string) *CheckAccountNameResponse {
	s.Headers = v
	return s
}

func (s *CheckAccountNameResponse) SetStatusCode(v int32) *CheckAccountNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccountNameResponse) SetBody(v *CheckAccountNameResponseBody) *CheckAccountNameResponse {
	s.Body = v
	return s
}

func (s *CheckAccountNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
