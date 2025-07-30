// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainProxyTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *DeleteDomainProxyTokenRequest
	GetDomainId() *string
	SetDomainProxyTokenId(v string) *DeleteDomainProxyTokenRequest
	GetDomainProxyTokenId() *string
	SetInstanceId(v string) *DeleteDomainProxyTokenRequest
	GetInstanceId() *string
}

type DeleteDomainProxyTokenRequest struct {
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

func (s DeleteDomainProxyTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainProxyTokenRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainProxyTokenRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *DeleteDomainProxyTokenRequest) GetDomainProxyTokenId() *string {
	return s.DomainProxyTokenId
}

func (s *DeleteDomainProxyTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDomainProxyTokenRequest) SetDomainId(v string) *DeleteDomainProxyTokenRequest {
	s.DomainId = &v
	return s
}

func (s *DeleteDomainProxyTokenRequest) SetDomainProxyTokenId(v string) *DeleteDomainProxyTokenRequest {
	s.DomainProxyTokenId = &v
	return s
}

func (s *DeleteDomainProxyTokenRequest) SetInstanceId(v string) *DeleteDomainProxyTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDomainProxyTokenRequest) Validate() error {
	return dara.Validate(s)
}
