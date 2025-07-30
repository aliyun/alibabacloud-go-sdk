// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *SetDefaultDomainRequest
	GetDomainId() *string
	SetInstanceId(v string) *SetDefaultDomainRequest
	GetInstanceId() *string
}

type SetDefaultDomainRequest struct {
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

func (s SetDefaultDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultDomainRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultDomainRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *SetDefaultDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetDefaultDomainRequest) SetDomainId(v string) *SetDefaultDomainRequest {
	s.DomainId = &v
	return s
}

func (s *SetDefaultDomainRequest) SetInstanceId(v string) *SetDefaultDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *SetDefaultDomainRequest) Validate() error {
	return dara.Validate(s)
}
