// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainDomainProxyTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *ObtainDomainProxyTokenRequest
	GetDomainId() *string
	SetDomainProxyTokenId(v string) *ObtainDomainProxyTokenRequest
	GetDomainProxyTokenId() *string
	SetInstanceId(v string) *ObtainDomainProxyTokenRequest
	GetInstanceId() *string
}

type ObtainDomainProxyTokenRequest struct {
	// The ID of the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The ID of the proxy token of the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pt_examplexxxx
	DomainProxyTokenId *string `json:"DomainProxyTokenId,omitempty" xml:"DomainProxyTokenId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ObtainDomainProxyTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ObtainDomainProxyTokenRequest) GoString() string {
	return s.String()
}

func (s *ObtainDomainProxyTokenRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *ObtainDomainProxyTokenRequest) GetDomainProxyTokenId() *string {
	return s.DomainProxyTokenId
}

func (s *ObtainDomainProxyTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ObtainDomainProxyTokenRequest) SetDomainId(v string) *ObtainDomainProxyTokenRequest {
	s.DomainId = &v
	return s
}

func (s *ObtainDomainProxyTokenRequest) SetDomainProxyTokenId(v string) *ObtainDomainProxyTokenRequest {
	s.DomainProxyTokenId = &v
	return s
}

func (s *ObtainDomainProxyTokenRequest) SetInstanceId(v string) *ObtainDomainProxyTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *ObtainDomainProxyTokenRequest) Validate() error {
	return dara.Validate(s)
}
