// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppSecurityGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppSecurityGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppSecurityGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppSecurityGroupResponseBody) *UpdateAppSecurityGroupResponse
	GetBody() *UpdateAppSecurityGroupResponseBody
}

type UpdateAppSecurityGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppSecurityGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppSecurityGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppSecurityGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppSecurityGroupResponse) GetBody() *UpdateAppSecurityGroupResponseBody {
	return s.Body
}

func (s *UpdateAppSecurityGroupResponse) SetHeaders(v map[string]*string) *UpdateAppSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppSecurityGroupResponse) SetStatusCode(v int32) *UpdateAppSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppSecurityGroupResponse) SetBody(v *UpdateAppSecurityGroupResponseBody) *UpdateAppSecurityGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateAppSecurityGroupResponse) Validate() error {
	return dara.Validate(s)
}
