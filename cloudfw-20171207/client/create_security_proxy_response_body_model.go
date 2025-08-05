// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProxyId(v string) *CreateSecurityProxyResponseBody
	GetProxyId() *string
	SetRequestId(v string) *CreateSecurityProxyResponseBody
	GetRequestId() *string
}

type CreateSecurityProxyResponseBody struct {
	// The ID of the NAT firewall.
	//
	// example:
	//
	// proxy-nat97ac4d7cc3834a5daf40
	ProxyId *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15FCCC52-1E23-57AE-B5EF-3E00A3DC3CAB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSecurityProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityProxyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityProxyResponseBody) GetProxyId() *string {
	return s.ProxyId
}

func (s *CreateSecurityProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecurityProxyResponseBody) SetProxyId(v string) *CreateSecurityProxyResponseBody {
	s.ProxyId = &v
	return s
}

func (s *CreateSecurityProxyResponseBody) SetRequestId(v string) *CreateSecurityProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
