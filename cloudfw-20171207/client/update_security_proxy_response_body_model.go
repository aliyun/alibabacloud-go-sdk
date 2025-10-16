// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v string) *UpdateSecurityProxyResponseBody
	GetModule() *string
	SetRequestId(v string) *UpdateSecurityProxyResponseBody
	GetRequestId() *string
}

type UpdateSecurityProxyResponseBody struct {
	// example:
	//
	// ips_server
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// 9E2CCAB4-E789-5BC9-88DC-5CE0358E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSecurityProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityProxyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecurityProxyResponseBody) GetModule() *string {
	return s.Module
}

func (s *UpdateSecurityProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSecurityProxyResponseBody) SetModule(v string) *UpdateSecurityProxyResponseBody {
	s.Module = &v
	return s
}

func (s *UpdateSecurityProxyResponseBody) SetRequestId(v string) *UpdateSecurityProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecurityProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
