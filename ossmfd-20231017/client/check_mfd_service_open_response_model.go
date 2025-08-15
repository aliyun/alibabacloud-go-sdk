// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMfdServiceOpenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckMfdServiceOpenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckMfdServiceOpenResponse
	GetStatusCode() *int32
	SetBody(v *CheckMfdServiceOpenResponseBody) *CheckMfdServiceOpenResponse
	GetBody() *CheckMfdServiceOpenResponseBody
}

type CheckMfdServiceOpenResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckMfdServiceOpenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckMfdServiceOpenResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckMfdServiceOpenResponse) GoString() string {
	return s.String()
}

func (s *CheckMfdServiceOpenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckMfdServiceOpenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckMfdServiceOpenResponse) GetBody() *CheckMfdServiceOpenResponseBody {
	return s.Body
}

func (s *CheckMfdServiceOpenResponse) SetHeaders(v map[string]*string) *CheckMfdServiceOpenResponse {
	s.Headers = v
	return s
}

func (s *CheckMfdServiceOpenResponse) SetStatusCode(v int32) *CheckMfdServiceOpenResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckMfdServiceOpenResponse) SetBody(v *CheckMfdServiceOpenResponseBody) *CheckMfdServiceOpenResponse {
	s.Body = v
	return s
}

func (s *CheckMfdServiceOpenResponse) Validate() error {
	return dara.Validate(s)
}
