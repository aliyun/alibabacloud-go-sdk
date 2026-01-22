// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountExistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAccountExistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAccountExistResponse
	GetStatusCode() *int32
	SetBody(v *CheckAccountExistResponseBody) *CheckAccountExistResponse
	GetBody() *CheckAccountExistResponseBody
}

type CheckAccountExistResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAccountExistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAccountExistResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountExistResponse) GoString() string {
	return s.String()
}

func (s *CheckAccountExistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAccountExistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAccountExistResponse) GetBody() *CheckAccountExistResponseBody {
	return s.Body
}

func (s *CheckAccountExistResponse) SetHeaders(v map[string]*string) *CheckAccountExistResponse {
	s.Headers = v
	return s
}

func (s *CheckAccountExistResponse) SetStatusCode(v int32) *CheckAccountExistResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccountExistResponse) SetBody(v *CheckAccountExistResponseBody) *CheckAccountExistResponse {
	s.Body = v
	return s
}

func (s *CheckAccountExistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
