// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAccountDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAccountDeleteResponse
	GetStatusCode() *int32
	SetBody(v *CheckAccountDeleteResponseBody) *CheckAccountDeleteResponse
	GetBody() *CheckAccountDeleteResponseBody
}

type CheckAccountDeleteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAccountDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAccountDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountDeleteResponse) GoString() string {
	return s.String()
}

func (s *CheckAccountDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAccountDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAccountDeleteResponse) GetBody() *CheckAccountDeleteResponseBody {
	return s.Body
}

func (s *CheckAccountDeleteResponse) SetHeaders(v map[string]*string) *CheckAccountDeleteResponse {
	s.Headers = v
	return s
}

func (s *CheckAccountDeleteResponse) SetStatusCode(v int32) *CheckAccountDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccountDeleteResponse) SetBody(v *CheckAccountDeleteResponseBody) *CheckAccountDeleteResponse {
	s.Body = v
	return s
}

func (s *CheckAccountDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
