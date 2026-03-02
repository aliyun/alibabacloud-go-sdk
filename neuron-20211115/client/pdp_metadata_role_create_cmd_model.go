// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpMetadataRoleCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *PdpMetadataRoleCreateCmd
	GetCompanyId() *int64
	SetName(v string) *PdpMetadataRoleCreateCmd
	GetName() *string
	SetPbcId(v int64) *PdpMetadataRoleCreateCmd
	GetPbcId() *int64
}

type PdpMetadataRoleCreateCmd struct {
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
}

func (s PdpMetadataRoleCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s PdpMetadataRoleCreateCmd) GoString() string {
	return s.String()
}

func (s *PdpMetadataRoleCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PdpMetadataRoleCreateCmd) GetName() *string {
	return s.Name
}

func (s *PdpMetadataRoleCreateCmd) GetPbcId() *int64 {
	return s.PbcId
}

func (s *PdpMetadataRoleCreateCmd) SetCompanyId(v int64) *PdpMetadataRoleCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *PdpMetadataRoleCreateCmd) SetName(v string) *PdpMetadataRoleCreateCmd {
	s.Name = &v
	return s
}

func (s *PdpMetadataRoleCreateCmd) SetPbcId(v int64) *PdpMetadataRoleCreateCmd {
	s.PbcId = &v
	return s
}

func (s *PdpMetadataRoleCreateCmd) Validate() error {
	return dara.Validate(s)
}
