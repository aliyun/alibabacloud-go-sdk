// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarApplyAddResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CarApplyAddResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CarApplyAddResponse
	GetStatusCode() *int32
	SetBody(v *CarApplyAddResponseBody) *CarApplyAddResponse
	GetBody() *CarApplyAddResponseBody
}

type CarApplyAddResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CarApplyAddResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CarApplyAddResponse) String() string {
	return dara.Prettify(s)
}

func (s CarApplyAddResponse) GoString() string {
	return s.String()
}

func (s *CarApplyAddResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CarApplyAddResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CarApplyAddResponse) GetBody() *CarApplyAddResponseBody {
	return s.Body
}

func (s *CarApplyAddResponse) SetHeaders(v map[string]*string) *CarApplyAddResponse {
	s.Headers = v
	return s
}

func (s *CarApplyAddResponse) SetStatusCode(v int32) *CarApplyAddResponse {
	s.StatusCode = &v
	return s
}

func (s *CarApplyAddResponse) SetBody(v *CarApplyAddResponseBody) *CarApplyAddResponse {
	s.Body = v
	return s
}

func (s *CarApplyAddResponse) Validate() error {
	return dara.Validate(s)
}
