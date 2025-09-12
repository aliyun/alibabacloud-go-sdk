// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceSecurityGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceSecurityGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceSecurityGroupsResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceSecurityGroupsResponseBody) *GetInstanceSecurityGroupsResponse
	GetBody() *GetInstanceSecurityGroupsResponseBody
}

type GetInstanceSecurityGroupsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceSecurityGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceSecurityGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceSecurityGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceSecurityGroupsResponse) GetBody() *GetInstanceSecurityGroupsResponseBody {
	return s.Body
}

func (s *GetInstanceSecurityGroupsResponse) SetHeaders(v map[string]*string) *GetInstanceSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceSecurityGroupsResponse) SetStatusCode(v int32) *GetInstanceSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceSecurityGroupsResponse) SetBody(v *GetInstanceSecurityGroupsResponseBody) *GetInstanceSecurityGroupsResponse {
	s.Body = v
	return s
}

func (s *GetInstanceSecurityGroupsResponse) Validate() error {
	return dara.Validate(s)
}
