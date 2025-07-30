// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDomainProxyTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *DisableDomainProxyTokenRequest
	GetDomainId() *string
	SetDomainProxyTokenId(v string) *DisableDomainProxyTokenRequest
	GetDomainProxyTokenId() *string
	SetInstanceId(v string) *DisableDomainProxyTokenRequest
	GetInstanceId() *string
}

type DisableDomainProxyTokenRequest struct {
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

func (s DisableDomainProxyTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDomainProxyTokenRequest) GoString() string {
	return s.String()
}

func (s *DisableDomainProxyTokenRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *DisableDomainProxyTokenRequest) GetDomainProxyTokenId() *string {
	return s.DomainProxyTokenId
}

func (s *DisableDomainProxyTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableDomainProxyTokenRequest) SetDomainId(v string) *DisableDomainProxyTokenRequest {
	s.DomainId = &v
	return s
}

func (s *DisableDomainProxyTokenRequest) SetDomainProxyTokenId(v string) *DisableDomainProxyTokenRequest {
	s.DomainProxyTokenId = &v
	return s
}

func (s *DisableDomainProxyTokenRequest) SetInstanceId(v string) *DisableDomainProxyTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableDomainProxyTokenRequest) Validate() error {
	return dara.Validate(s)
}
