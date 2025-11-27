// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountNameAvailableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAccountNameAvailableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAccountNameAvailableResponse
	GetStatusCode() *int32
	SetBody(v *CheckAccountNameAvailableResponseBody) *CheckAccountNameAvailableResponse
	GetBody() *CheckAccountNameAvailableResponseBody
}

type CheckAccountNameAvailableResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAccountNameAvailableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAccountNameAvailableResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountNameAvailableResponse) GoString() string {
	return s.String()
}

func (s *CheckAccountNameAvailableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAccountNameAvailableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAccountNameAvailableResponse) GetBody() *CheckAccountNameAvailableResponseBody {
	return s.Body
}

func (s *CheckAccountNameAvailableResponse) SetHeaders(v map[string]*string) *CheckAccountNameAvailableResponse {
	s.Headers = v
	return s
}

func (s *CheckAccountNameAvailableResponse) SetStatusCode(v int32) *CheckAccountNameAvailableResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccountNameAvailableResponse) SetBody(v *CheckAccountNameAvailableResponseBody) *CheckAccountNameAvailableResponse {
	s.Body = v
	return s
}

func (s *CheckAccountNameAvailableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
