// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectAddResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProjectAddResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProjectAddResponse
	GetStatusCode() *int32
	SetBody(v *ProjectAddResponseBody) *ProjectAddResponse
	GetBody() *ProjectAddResponseBody
}

type ProjectAddResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProjectAddResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProjectAddResponse) String() string {
	return dara.Prettify(s)
}

func (s ProjectAddResponse) GoString() string {
	return s.String()
}

func (s *ProjectAddResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProjectAddResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProjectAddResponse) GetBody() *ProjectAddResponseBody {
	return s.Body
}

func (s *ProjectAddResponse) SetHeaders(v map[string]*string) *ProjectAddResponse {
	s.Headers = v
	return s
}

func (s *ProjectAddResponse) SetStatusCode(v int32) *ProjectAddResponse {
	s.StatusCode = &v
	return s
}

func (s *ProjectAddResponse) SetBody(v *ProjectAddResponseBody) *ProjectAddResponse {
	s.Body = v
	return s
}

func (s *ProjectAddResponse) Validate() error {
	return dara.Validate(s)
}
