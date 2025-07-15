// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateNetworkAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateNetworkAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateNetworkAclResponse
	GetStatusCode() *int32
	SetBody(v *AssociateNetworkAclResponseBody) *AssociateNetworkAclResponse
	GetBody() *AssociateNetworkAclResponseBody
}

type AssociateNetworkAclResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateNetworkAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateNetworkAclResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateNetworkAclResponse) GoString() string {
	return s.String()
}

func (s *AssociateNetworkAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateNetworkAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateNetworkAclResponse) GetBody() *AssociateNetworkAclResponseBody {
	return s.Body
}

func (s *AssociateNetworkAclResponse) SetHeaders(v map[string]*string) *AssociateNetworkAclResponse {
	s.Headers = v
	return s
}

func (s *AssociateNetworkAclResponse) SetStatusCode(v int32) *AssociateNetworkAclResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateNetworkAclResponse) SetBody(v *AssociateNetworkAclResponseBody) *AssociateNetworkAclResponse {
	s.Body = v
	return s
}

func (s *AssociateNetworkAclResponse) Validate() error {
	return dara.Validate(s)
}
