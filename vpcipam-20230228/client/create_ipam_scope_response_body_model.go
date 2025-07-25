// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpamScopeId(v string) *CreateIpamScopeResponseBody
	GetIpamScopeId() *string
	SetRequestId(v string) *CreateIpamScopeResponseBody
	GetRequestId() *string
}

type CreateIpamScopeResponseBody struct {
	// The ID of the IPAM scope.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E897D16A-50EB-543F-B002-C5A26AB818FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpamScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamScopeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpamScopeResponseBody) GetIpamScopeId() *string {
	return s.IpamScopeId
}

func (s *CreateIpamScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIpamScopeResponseBody) SetIpamScopeId(v string) *CreateIpamScopeResponseBody {
	s.IpamScopeId = &v
	return s
}

func (s *CreateIpamScopeResponseBody) SetRequestId(v string) *CreateIpamScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpamScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
