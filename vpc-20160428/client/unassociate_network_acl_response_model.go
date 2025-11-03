// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateNetworkAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassociateNetworkAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassociateNetworkAclResponse
	GetStatusCode() *int32
	SetBody(v *UnassociateNetworkAclResponseBody) *UnassociateNetworkAclResponse
	GetBody() *UnassociateNetworkAclResponseBody
}

type UnassociateNetworkAclResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassociateNetworkAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassociateNetworkAclResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassociateNetworkAclResponse) GoString() string {
	return s.String()
}

func (s *UnassociateNetworkAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassociateNetworkAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassociateNetworkAclResponse) GetBody() *UnassociateNetworkAclResponseBody {
	return s.Body
}

func (s *UnassociateNetworkAclResponse) SetHeaders(v map[string]*string) *UnassociateNetworkAclResponse {
	s.Headers = v
	return s
}

func (s *UnassociateNetworkAclResponse) SetStatusCode(v int32) *UnassociateNetworkAclResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassociateNetworkAclResponse) SetBody(v *UnassociateNetworkAclResponseBody) *UnassociateNetworkAclResponse {
	s.Body = v
	return s
}

func (s *UnassociateNetworkAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
