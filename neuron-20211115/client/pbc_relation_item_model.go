// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcRelationItem interface {
	dara.Model
	String() string
	GoString() string
	SetCompany(v string) *PbcRelationItem
	GetCompany() *string
	SetDeveloper(v string) *PbcRelationItem
	GetDeveloper() *string
	SetPbcName(v string) *PbcRelationItem
	GetPbcName() *string
	SetPbcVersion(v string) *PbcRelationItem
	GetPbcVersion() *string
	SetSummary(v string) *PbcRelationItem
	GetSummary() *string
}

type PbcRelationItem struct {
	// example:
	//
	// 企业服务
	Company *string `json:"company,omitempty" xml:"company,omitempty"`
	// example:
	//
	// 中驿
	Developer *string `json:"developer,omitempty" xml:"developer,omitempty"`
	// example:
	//
	// category
	PbcName *string `json:"pbcName,omitempty" xml:"pbcName,omitempty"`
	// example:
	//
	// 2022-04-01T00:00:00.000Z
	PbcVersion *string `json:"pbcVersion,omitempty" xml:"pbcVersion,omitempty"`
	// example:
	//
	// 资产市场
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
}

func (s PbcRelationItem) String() string {
	return dara.Prettify(s)
}

func (s PbcRelationItem) GoString() string {
	return s.String()
}

func (s *PbcRelationItem) GetCompany() *string {
	return s.Company
}

func (s *PbcRelationItem) GetDeveloper() *string {
	return s.Developer
}

func (s *PbcRelationItem) GetPbcName() *string {
	return s.PbcName
}

func (s *PbcRelationItem) GetPbcVersion() *string {
	return s.PbcVersion
}

func (s *PbcRelationItem) GetSummary() *string {
	return s.Summary
}

func (s *PbcRelationItem) SetCompany(v string) *PbcRelationItem {
	s.Company = &v
	return s
}

func (s *PbcRelationItem) SetDeveloper(v string) *PbcRelationItem {
	s.Developer = &v
	return s
}

func (s *PbcRelationItem) SetPbcName(v string) *PbcRelationItem {
	s.PbcName = &v
	return s
}

func (s *PbcRelationItem) SetPbcVersion(v string) *PbcRelationItem {
	s.PbcVersion = &v
	return s
}

func (s *PbcRelationItem) SetSummary(v string) *PbcRelationItem {
	s.Summary = &v
	return s
}

func (s *PbcRelationItem) Validate() error {
	return dara.Validate(s)
}
