// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckVodDefaultRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckVodDefaultRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckVodDefaultRoleResponse
	GetStatusCode() *int32
	SetBody(v *CheckVodDefaultRoleResponseBody) *CheckVodDefaultRoleResponse
	GetBody() *CheckVodDefaultRoleResponseBody
}

type CheckVodDefaultRoleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckVodDefaultRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckVodDefaultRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckVodDefaultRoleResponse) GoString() string {
	return s.String()
}

func (s *CheckVodDefaultRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckVodDefaultRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckVodDefaultRoleResponse) GetBody() *CheckVodDefaultRoleResponseBody {
	return s.Body
}

func (s *CheckVodDefaultRoleResponse) SetHeaders(v map[string]*string) *CheckVodDefaultRoleResponse {
	s.Headers = v
	return s
}

func (s *CheckVodDefaultRoleResponse) SetStatusCode(v int32) *CheckVodDefaultRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckVodDefaultRoleResponse) SetBody(v *CheckVodDefaultRoleResponseBody) *CheckVodDefaultRoleResponse {
	s.Body = v
	return s
}

func (s *CheckVodDefaultRoleResponse) Validate() error {
	return dara.Validate(s)
}
