// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpServiceRoleCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *PdpServiceRoleCreateCmd
	GetCompanyId() *int64
	SetName(v string) *PdpServiceRoleCreateCmd
	GetName() *string
	SetPbcId(v int64) *PdpServiceRoleCreateCmd
	GetPbcId() *int64
	SetServiceId(v int64) *PdpServiceRoleCreateCmd
	GetServiceId() *int64
}

type PdpServiceRoleCreateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PbcId *int64 `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s PdpServiceRoleCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s PdpServiceRoleCreateCmd) GoString() string {
	return s.String()
}

func (s *PdpServiceRoleCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PdpServiceRoleCreateCmd) GetName() *string {
	return s.Name
}

func (s *PdpServiceRoleCreateCmd) GetPbcId() *int64 {
	return s.PbcId
}

func (s *PdpServiceRoleCreateCmd) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *PdpServiceRoleCreateCmd) SetCompanyId(v int64) *PdpServiceRoleCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *PdpServiceRoleCreateCmd) SetName(v string) *PdpServiceRoleCreateCmd {
	s.Name = &v
	return s
}

func (s *PdpServiceRoleCreateCmd) SetPbcId(v int64) *PdpServiceRoleCreateCmd {
	s.PbcId = &v
	return s
}

func (s *PdpServiceRoleCreateCmd) SetServiceId(v int64) *PdpServiceRoleCreateCmd {
	s.ServiceId = &v
	return s
}

func (s *PdpServiceRoleCreateCmd) Validate() error {
	return dara.Validate(s)
}
