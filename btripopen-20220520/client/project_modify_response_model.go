// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectModifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProjectModifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProjectModifyResponse
	GetStatusCode() *int32
	SetBody(v *ProjectModifyResponseBody) *ProjectModifyResponse
	GetBody() *ProjectModifyResponseBody
}

type ProjectModifyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProjectModifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProjectModifyResponse) String() string {
	return dara.Prettify(s)
}

func (s ProjectModifyResponse) GoString() string {
	return s.String()
}

func (s *ProjectModifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProjectModifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProjectModifyResponse) GetBody() *ProjectModifyResponseBody {
	return s.Body
}

func (s *ProjectModifyResponse) SetHeaders(v map[string]*string) *ProjectModifyResponse {
	s.Headers = v
	return s
}

func (s *ProjectModifyResponse) SetStatusCode(v int32) *ProjectModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *ProjectModifyResponse) SetBody(v *ProjectModifyResponseBody) *ProjectModifyResponse {
	s.Body = v
	return s
}

func (s *ProjectModifyResponse) Validate() error {
	return dara.Validate(s)
}
