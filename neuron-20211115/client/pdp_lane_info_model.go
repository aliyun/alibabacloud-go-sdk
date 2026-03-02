// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpLaneInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *PdpLaneInfo
	GetAlias() *string
	SetCompanyId(v int64) *PdpLaneInfo
	GetCompanyId() *int64
	SetCreatorName(v string) *PdpLaneInfo
	GetCreatorName() *string
	SetCreatorUid(v string) *PdpLaneInfo
	GetCreatorUid() *string
	SetDescription(v string) *PdpLaneInfo
	GetDescription() *string
	SetEnv(v string) *PdpLaneInfo
	GetEnv() *string
	SetGmtCreate(v string) *PdpLaneInfo
	GetGmtCreate() *string
	SetId(v int64) *PdpLaneInfo
	GetId() *int64
	SetInitGroupFlag(v bool) *PdpLaneInfo
	GetInitGroupFlag() *bool
	SetInletServices(v string) *PdpLaneInfo
	GetInletServices() *string
	SetName(v string) *PdpLaneInfo
	GetName() *string
	SetProductId(v int64) *PdpLaneInfo
	GetProductId() *int64
	SetProductName(v string) *PdpLaneInfo
	GetProductName() *string
	SetServiceGroupIds(v string) *PdpLaneInfo
	GetServiceGroupIds() *string
	SetType(v string) *PdpLaneInfo
	GetType() *string
}

type PdpLaneInfo struct {
	// example:
	//
	// 灰度
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// 1
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// example:
	//
	// 张三
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// 252333049301505383
	CreatorUid  *string `json:"creatorUid,omitempty" xml:"creatorUid,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// 2023-07-04T00:00:00.000Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	InitGroupFlag *bool `json:"initGroupFlag,omitempty" xml:"initGroupFlag,omitempty"`
	// example:
	//
	// 1，2
	InletServices *string `json:"inletServices,omitempty" xml:"inletServices,omitempty"`
	// example:
	//
	// gray
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// linkedmall
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// 1，2
	ServiceGroupIds *string `json:"serviceGroupIds,omitempty" xml:"serviceGroupIds,omitempty"`
	// example:
	//
	// SYSTEM
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PdpLaneInfo) String() string {
	return dara.Prettify(s)
}

func (s PdpLaneInfo) GoString() string {
	return s.String()
}

func (s *PdpLaneInfo) GetAlias() *string {
	return s.Alias
}

func (s *PdpLaneInfo) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PdpLaneInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *PdpLaneInfo) GetCreatorUid() *string {
	return s.CreatorUid
}

func (s *PdpLaneInfo) GetDescription() *string {
	return s.Description
}

func (s *PdpLaneInfo) GetEnv() *string {
	return s.Env
}

func (s *PdpLaneInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PdpLaneInfo) GetId() *int64 {
	return s.Id
}

func (s *PdpLaneInfo) GetInitGroupFlag() *bool {
	return s.InitGroupFlag
}

func (s *PdpLaneInfo) GetInletServices() *string {
	return s.InletServices
}

func (s *PdpLaneInfo) GetName() *string {
	return s.Name
}

func (s *PdpLaneInfo) GetProductId() *int64 {
	return s.ProductId
}

func (s *PdpLaneInfo) GetProductName() *string {
	return s.ProductName
}

func (s *PdpLaneInfo) GetServiceGroupIds() *string {
	return s.ServiceGroupIds
}

func (s *PdpLaneInfo) GetType() *string {
	return s.Type
}

func (s *PdpLaneInfo) SetAlias(v string) *PdpLaneInfo {
	s.Alias = &v
	return s
}

func (s *PdpLaneInfo) SetCompanyId(v int64) *PdpLaneInfo {
	s.CompanyId = &v
	return s
}

func (s *PdpLaneInfo) SetCreatorName(v string) *PdpLaneInfo {
	s.CreatorName = &v
	return s
}

func (s *PdpLaneInfo) SetCreatorUid(v string) *PdpLaneInfo {
	s.CreatorUid = &v
	return s
}

func (s *PdpLaneInfo) SetDescription(v string) *PdpLaneInfo {
	s.Description = &v
	return s
}

func (s *PdpLaneInfo) SetEnv(v string) *PdpLaneInfo {
	s.Env = &v
	return s
}

func (s *PdpLaneInfo) SetGmtCreate(v string) *PdpLaneInfo {
	s.GmtCreate = &v
	return s
}

func (s *PdpLaneInfo) SetId(v int64) *PdpLaneInfo {
	s.Id = &v
	return s
}

func (s *PdpLaneInfo) SetInitGroupFlag(v bool) *PdpLaneInfo {
	s.InitGroupFlag = &v
	return s
}

func (s *PdpLaneInfo) SetInletServices(v string) *PdpLaneInfo {
	s.InletServices = &v
	return s
}

func (s *PdpLaneInfo) SetName(v string) *PdpLaneInfo {
	s.Name = &v
	return s
}

func (s *PdpLaneInfo) SetProductId(v int64) *PdpLaneInfo {
	s.ProductId = &v
	return s
}

func (s *PdpLaneInfo) SetProductName(v string) *PdpLaneInfo {
	s.ProductName = &v
	return s
}

func (s *PdpLaneInfo) SetServiceGroupIds(v string) *PdpLaneInfo {
	s.ServiceGroupIds = &v
	return s
}

func (s *PdpLaneInfo) SetType(v string) *PdpLaneInfo {
	s.Type = &v
	return s
}

func (s *PdpLaneInfo) Validate() error {
	return dara.Validate(s)
}
