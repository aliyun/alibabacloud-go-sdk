// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainProxyTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *CreateDomainProxyTokenRequest
	GetDomainId() *string
	SetInstanceId(v string) *CreateDomainProxyTokenRequest
	GetInstanceId() *string
}

type CreateDomainProxyTokenRequest struct {
	// The ID of the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateDomainProxyTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainProxyTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainProxyTokenRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *CreateDomainProxyTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDomainProxyTokenRequest) SetDomainId(v string) *CreateDomainProxyTokenRequest {
	s.DomainId = &v
	return s
}

func (s *CreateDomainProxyTokenRequest) SetInstanceId(v string) *CreateDomainProxyTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDomainProxyTokenRequest) Validate() error {
	return dara.Validate(s)
}
