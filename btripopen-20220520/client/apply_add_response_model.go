// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAddResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyAddResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyAddResponse
	GetStatusCode() *int32
	SetBody(v *ApplyAddResponseBody) *ApplyAddResponse
	GetBody() *ApplyAddResponseBody
}

type ApplyAddResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyAddResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyAddResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddResponse) GoString() string {
	return s.String()
}

func (s *ApplyAddResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyAddResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyAddResponse) GetBody() *ApplyAddResponseBody {
	return s.Body
}

func (s *ApplyAddResponse) SetHeaders(v map[string]*string) *ApplyAddResponse {
	s.Headers = v
	return s
}

func (s *ApplyAddResponse) SetStatusCode(v int32) *ApplyAddResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyAddResponse) SetBody(v *ApplyAddResponseBody) *ApplyAddResponse {
	s.Body = v
	return s
}

func (s *ApplyAddResponse) Validate() error {
	return dara.Validate(s)
}
