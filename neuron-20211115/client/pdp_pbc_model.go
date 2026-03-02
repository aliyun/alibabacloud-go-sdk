// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpPbc interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *PdpPbc
	GetAlias() *string
	SetCompany(v string) *PdpPbc
	GetCompany() *string
	SetCompanyId(v int64) *PdpPbc
	GetCompanyId() *int64
	SetDescription(v string) *PdpPbc
	GetDescription() *string
	SetDeveloperId(v string) *PdpPbc
	GetDeveloperId() *string
	SetGitGroup(v string) *PdpPbc
	GetGitGroup() *string
	SetGitGroupInfo(v string) *PdpPbc
	GetGitGroupInfo() *string
	SetGmtCreate(v string) *PdpPbc
	GetGmtCreate() *string
	SetId(v int64) *PdpPbc
	GetId() *int64
	SetName(v string) *PdpPbc
	GetName() *string
	SetRequestId(v string) *PdpPbc
	GetRequestId() *string
	SetSnowbergDisplay(v bool) *PdpPbc
	GetSnowbergDisplay() *bool
	SetType(v string) *PdpPbc
	GetType() *string
}

type PdpPbc struct {
	// This parameter is required.
	//
	// example:
	//
	// 基础商品
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 多端商城
	Company *string `json:"company,omitempty" xml:"company,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// example:
	//
	// 基础商品描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	DeveloperId *string `json:"developerId,omitempty" xml:"developerId,omitempty"`
	// example:
	//
	// neuron-pdp-pbc
	GitGroup *string `json:"gitGroup,omitempty" xml:"gitGroup,omitempty"`
	// example:
	//
	// {"path":"neuron-pdp-pbc","webUrl":"https://codeup.aliyun.com/616d751bbc23ecd311af9711/neuron-pdp-pbc","name":"neuron-pdp-pbc","id":480037,"ownerId":663171,"pathWithNamespace":"616d751bbc23ecd311af9711/neuron-pdp-pbc","parentId":330007}
	GitGroupInfo *string `json:"gitGroupInfo,omitempty" xml:"gitGroupInfo,omitempty"`
	// example:
	//
	// 2022-05-01T00:00:00.000Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// product
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestId       *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SnowbergDisplay *bool   `json:"snowbergDisplay,omitempty" xml:"snowbergDisplay,omitempty"`
	// example:
	//
	// INNER
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PdpPbc) String() string {
	return dara.Prettify(s)
}

func (s PdpPbc) GoString() string {
	return s.String()
}

func (s *PdpPbc) GetAlias() *string {
	return s.Alias
}

func (s *PdpPbc) GetCompany() *string {
	return s.Company
}

func (s *PdpPbc) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PdpPbc) GetDescription() *string {
	return s.Description
}

func (s *PdpPbc) GetDeveloperId() *string {
	return s.DeveloperId
}

func (s *PdpPbc) GetGitGroup() *string {
	return s.GitGroup
}

func (s *PdpPbc) GetGitGroupInfo() *string {
	return s.GitGroupInfo
}

func (s *PdpPbc) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PdpPbc) GetId() *int64 {
	return s.Id
}

func (s *PdpPbc) GetName() *string {
	return s.Name
}

func (s *PdpPbc) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpPbc) GetSnowbergDisplay() *bool {
	return s.SnowbergDisplay
}

func (s *PdpPbc) GetType() *string {
	return s.Type
}

func (s *PdpPbc) SetAlias(v string) *PdpPbc {
	s.Alias = &v
	return s
}

func (s *PdpPbc) SetCompany(v string) *PdpPbc {
	s.Company = &v
	return s
}

func (s *PdpPbc) SetCompanyId(v int64) *PdpPbc {
	s.CompanyId = &v
	return s
}

func (s *PdpPbc) SetDescription(v string) *PdpPbc {
	s.Description = &v
	return s
}

func (s *PdpPbc) SetDeveloperId(v string) *PdpPbc {
	s.DeveloperId = &v
	return s
}

func (s *PdpPbc) SetGitGroup(v string) *PdpPbc {
	s.GitGroup = &v
	return s
}

func (s *PdpPbc) SetGitGroupInfo(v string) *PdpPbc {
	s.GitGroupInfo = &v
	return s
}

func (s *PdpPbc) SetGmtCreate(v string) *PdpPbc {
	s.GmtCreate = &v
	return s
}

func (s *PdpPbc) SetId(v int64) *PdpPbc {
	s.Id = &v
	return s
}

func (s *PdpPbc) SetName(v string) *PdpPbc {
	s.Name = &v
	return s
}

func (s *PdpPbc) SetRequestId(v string) *PdpPbc {
	s.RequestId = &v
	return s
}

func (s *PdpPbc) SetSnowbergDisplay(v bool) *PdpPbc {
	s.SnowbergDisplay = &v
	return s
}

func (s *PdpPbc) SetType(v string) *PdpPbc {
	s.Type = &v
	return s
}

func (s *PdpPbc) Validate() error {
	return dara.Validate(s)
}
