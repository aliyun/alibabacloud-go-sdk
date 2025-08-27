// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProjectDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProjectDeleteResponse
	GetStatusCode() *int32
	SetBody(v *ProjectDeleteResponseBody) *ProjectDeleteResponse
	GetBody() *ProjectDeleteResponseBody
}

type ProjectDeleteResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProjectDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProjectDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s ProjectDeleteResponse) GoString() string {
	return s.String()
}

func (s *ProjectDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProjectDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProjectDeleteResponse) GetBody() *ProjectDeleteResponseBody {
	return s.Body
}

func (s *ProjectDeleteResponse) SetHeaders(v map[string]*string) *ProjectDeleteResponse {
	s.Headers = v
	return s
}

func (s *ProjectDeleteResponse) SetStatusCode(v int32) *ProjectDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *ProjectDeleteResponse) SetBody(v *ProjectDeleteResponseBody) *ProjectDeleteResponse {
	s.Body = v
	return s
}

func (s *ProjectDeleteResponse) Validate() error {
	return dara.Validate(s)
}
