// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpPbcManagerRoleCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *PdpPbcManagerRoleCreateCmd
	GetCompanyId() *int64
	SetName(v string) *PdpPbcManagerRoleCreateCmd
	GetName() *string
	SetPbcId(v int64) *PdpPbcManagerRoleCreateCmd
	GetPbcId() *int64
}

type PdpPbcManagerRoleCreateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// product
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PbcId *int64 `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
}

func (s PdpPbcManagerRoleCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s PdpPbcManagerRoleCreateCmd) GoString() string {
	return s.String()
}

func (s *PdpPbcManagerRoleCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PdpPbcManagerRoleCreateCmd) GetName() *string {
	return s.Name
}

func (s *PdpPbcManagerRoleCreateCmd) GetPbcId() *int64 {
	return s.PbcId
}

func (s *PdpPbcManagerRoleCreateCmd) SetCompanyId(v int64) *PdpPbcManagerRoleCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *PdpPbcManagerRoleCreateCmd) SetName(v string) *PdpPbcManagerRoleCreateCmd {
	s.Name = &v
	return s
}

func (s *PdpPbcManagerRoleCreateCmd) SetPbcId(v int64) *PdpPbcManagerRoleCreateCmd {
	s.PbcId = &v
	return s
}

func (s *PdpPbcManagerRoleCreateCmd) Validate() error {
	return dara.Validate(s)
}
