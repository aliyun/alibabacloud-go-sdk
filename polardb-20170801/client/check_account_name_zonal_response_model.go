// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountNameZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAccountNameZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAccountNameZonalResponse
	GetStatusCode() *int32
	SetBody(v *CheckAccountNameZonalResponseBody) *CheckAccountNameZonalResponse
	GetBody() *CheckAccountNameZonalResponseBody
}

type CheckAccountNameZonalResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAccountNameZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAccountNameZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountNameZonalResponse) GoString() string {
	return s.String()
}

func (s *CheckAccountNameZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAccountNameZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAccountNameZonalResponse) GetBody() *CheckAccountNameZonalResponseBody {
	return s.Body
}

func (s *CheckAccountNameZonalResponse) SetHeaders(v map[string]*string) *CheckAccountNameZonalResponse {
	s.Headers = v
	return s
}

func (s *CheckAccountNameZonalResponse) SetStatusCode(v int32) *CheckAccountNameZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccountNameZonalResponse) SetBody(v *CheckAccountNameZonalResponseBody) *CheckAccountNameZonalResponse {
	s.Body = v
	return s
}

func (s *CheckAccountNameZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
