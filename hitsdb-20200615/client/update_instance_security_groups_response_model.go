// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceSecurityGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceSecurityGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceSecurityGroupsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceSecurityGroupsResponseBody) *UpdateInstanceSecurityGroupsResponse
	GetBody() *UpdateInstanceSecurityGroupsResponseBody
}

type UpdateInstanceSecurityGroupsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceSecurityGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSecurityGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceSecurityGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceSecurityGroupsResponse) GetBody() *UpdateInstanceSecurityGroupsResponseBody {
	return s.Body
}

func (s *UpdateInstanceSecurityGroupsResponse) SetHeaders(v map[string]*string) *UpdateInstanceSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceSecurityGroupsResponse) SetStatusCode(v int32) *UpdateInstanceSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsResponse) SetBody(v *UpdateInstanceSecurityGroupsResponseBody) *UpdateInstanceSecurityGroupsResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceSecurityGroupsResponse) Validate() error {
	return dara.Validate(s)
}
