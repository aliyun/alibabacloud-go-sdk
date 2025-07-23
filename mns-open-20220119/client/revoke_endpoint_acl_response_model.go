// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeEndpointAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeEndpointAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeEndpointAclResponse
	GetStatusCode() *int32
	SetBody(v *RevokeEndpointAclResponseBody) *RevokeEndpointAclResponse
	GetBody() *RevokeEndpointAclResponseBody
}

type RevokeEndpointAclResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeEndpointAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeEndpointAclResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeEndpointAclResponse) GoString() string {
	return s.String()
}

func (s *RevokeEndpointAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeEndpointAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeEndpointAclResponse) GetBody() *RevokeEndpointAclResponseBody {
	return s.Body
}

func (s *RevokeEndpointAclResponse) SetHeaders(v map[string]*string) *RevokeEndpointAclResponse {
	s.Headers = v
	return s
}

func (s *RevokeEndpointAclResponse) SetStatusCode(v int32) *RevokeEndpointAclResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeEndpointAclResponse) SetBody(v *RevokeEndpointAclResponseBody) *RevokeEndpointAclResponse {
	s.Body = v
	return s
}

func (s *RevokeEndpointAclResponse) Validate() error {
	return dara.Validate(s)
}
