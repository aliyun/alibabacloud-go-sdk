// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckReadableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckReadableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckReadableResponse
	GetStatusCode() *int32
	SetBody(v *CheckReadableResponseBody) *CheckReadableResponse
	GetBody() *CheckReadableResponseBody
}

type CheckReadableResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckReadableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckReadableResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckReadableResponse) GoString() string {
	return s.String()
}

func (s *CheckReadableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckReadableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckReadableResponse) GetBody() *CheckReadableResponseBody {
	return s.Body
}

func (s *CheckReadableResponse) SetHeaders(v map[string]*string) *CheckReadableResponse {
	s.Headers = v
	return s
}

func (s *CheckReadableResponse) SetStatusCode(v int32) *CheckReadableResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckReadableResponse) SetBody(v *CheckReadableResponseBody) *CheckReadableResponse {
	s.Body = v
	return s
}

func (s *CheckReadableResponse) Validate() error {
	return dara.Validate(s)
}
