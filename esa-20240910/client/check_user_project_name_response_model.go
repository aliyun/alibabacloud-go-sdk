// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserProjectNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckUserProjectNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckUserProjectNameResponse
	GetStatusCode() *int32
	SetBody(v *CheckUserProjectNameResponseBody) *CheckUserProjectNameResponse
	GetBody() *CheckUserProjectNameResponseBody
}

type CheckUserProjectNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUserProjectNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUserProjectNameResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckUserProjectNameResponse) GoString() string {
	return s.String()
}

func (s *CheckUserProjectNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckUserProjectNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckUserProjectNameResponse) GetBody() *CheckUserProjectNameResponseBody {
	return s.Body
}

func (s *CheckUserProjectNameResponse) SetHeaders(v map[string]*string) *CheckUserProjectNameResponse {
	s.Headers = v
	return s
}

func (s *CheckUserProjectNameResponse) SetStatusCode(v int32) *CheckUserProjectNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUserProjectNameResponse) SetBody(v *CheckUserProjectNameResponseBody) *CheckUserProjectNameResponse {
	s.Body = v
	return s
}

func (s *CheckUserProjectNameResponse) Validate() error {
	return dara.Validate(s)
}
