// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceCustomizedDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DeleteInstanceCustomizedDomainRequest
	GetDomain() *string
	SetInstanceId(v string) *DeleteInstanceCustomizedDomainRequest
	GetInstanceId() *string
	SetModuleName(v string) *DeleteInstanceCustomizedDomainRequest
	GetModuleName() *string
}

type DeleteInstanceCustomizedDomainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// crtest.wgine-inc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cri-av3kxfkkiwofikl5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Registry
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DeleteInstanceCustomizedDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceCustomizedDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceCustomizedDomainRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteInstanceCustomizedDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceCustomizedDomainRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *DeleteInstanceCustomizedDomainRequest) SetDomain(v string) *DeleteInstanceCustomizedDomainRequest {
	s.Domain = &v
	return s
}

func (s *DeleteInstanceCustomizedDomainRequest) SetInstanceId(v string) *DeleteInstanceCustomizedDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceCustomizedDomainRequest) SetModuleName(v string) *DeleteInstanceCustomizedDomainRequest {
	s.ModuleName = &v
	return s
}

func (s *DeleteInstanceCustomizedDomainRequest) Validate() error {
	return dara.Validate(s)
}
