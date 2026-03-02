// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronApp interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *NeuronApp
	GetAccessType() *string
	SetAccountId(v string) *NeuronApp
	GetAccountId() *string
	SetAlias(v string) *NeuronApp
	GetAlias() *string
	SetBizType(v string) *NeuronApp
	GetBizType() *string
	SetDescription(v string) *NeuronApp
	GetDescription() *string
	SetEnterpriseId(v int64) *NeuronApp
	GetEnterpriseId() *int64
	SetGmtCreate(v string) *NeuronApp
	GetGmtCreate() *string
	SetGmtModified(v string) *NeuronApp
	GetGmtModified() *string
	SetIconUrl(v string) *NeuronApp
	GetIconUrl() *string
	SetId(v int64) *NeuronApp
	GetId() *int64
	SetName(v string) *NeuronApp
	GetName() *string
	SetOwner(v string) *NeuronApp
	GetOwner() *string
	SetProductId(v int64) *NeuronApp
	GetProductId() *int64
	SetStatus(v string) *NeuronApp
	GetStatus() *string
}

type NeuronApp struct {
	// This parameter is required.
	//
	// example:
	//
	// PUBLIC
	AccessType *string `json:"accessType,omitempty" xml:"accessType,omitempty"`
	// example:
	//
	// 1343424
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 分销APP
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ORDER
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 升级订单功能
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	EnterpriseId *int64  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	GmtCreate    *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified  *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://wwww.ali.com
	IconUrl *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// order
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 多端商城
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// NEW
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s NeuronApp) String() string {
	return dara.Prettify(s)
}

func (s NeuronApp) GoString() string {
	return s.String()
}

func (s *NeuronApp) GetAccessType() *string {
	return s.AccessType
}

func (s *NeuronApp) GetAccountId() *string {
	return s.AccountId
}

func (s *NeuronApp) GetAlias() *string {
	return s.Alias
}

func (s *NeuronApp) GetBizType() *string {
	return s.BizType
}

func (s *NeuronApp) GetDescription() *string {
	return s.Description
}

func (s *NeuronApp) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *NeuronApp) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *NeuronApp) GetGmtModified() *string {
	return s.GmtModified
}

func (s *NeuronApp) GetIconUrl() *string {
	return s.IconUrl
}

func (s *NeuronApp) GetId() *int64 {
	return s.Id
}

func (s *NeuronApp) GetName() *string {
	return s.Name
}

func (s *NeuronApp) GetOwner() *string {
	return s.Owner
}

func (s *NeuronApp) GetProductId() *int64 {
	return s.ProductId
}

func (s *NeuronApp) GetStatus() *string {
	return s.Status
}

func (s *NeuronApp) SetAccessType(v string) *NeuronApp {
	s.AccessType = &v
	return s
}

func (s *NeuronApp) SetAccountId(v string) *NeuronApp {
	s.AccountId = &v
	return s
}

func (s *NeuronApp) SetAlias(v string) *NeuronApp {
	s.Alias = &v
	return s
}

func (s *NeuronApp) SetBizType(v string) *NeuronApp {
	s.BizType = &v
	return s
}

func (s *NeuronApp) SetDescription(v string) *NeuronApp {
	s.Description = &v
	return s
}

func (s *NeuronApp) SetEnterpriseId(v int64) *NeuronApp {
	s.EnterpriseId = &v
	return s
}

func (s *NeuronApp) SetGmtCreate(v string) *NeuronApp {
	s.GmtCreate = &v
	return s
}

func (s *NeuronApp) SetGmtModified(v string) *NeuronApp {
	s.GmtModified = &v
	return s
}

func (s *NeuronApp) SetIconUrl(v string) *NeuronApp {
	s.IconUrl = &v
	return s
}

func (s *NeuronApp) SetId(v int64) *NeuronApp {
	s.Id = &v
	return s
}

func (s *NeuronApp) SetName(v string) *NeuronApp {
	s.Name = &v
	return s
}

func (s *NeuronApp) SetOwner(v string) *NeuronApp {
	s.Owner = &v
	return s
}

func (s *NeuronApp) SetProductId(v int64) *NeuronApp {
	s.ProductId = &v
	return s
}

func (s *NeuronApp) SetStatus(v string) *NeuronApp {
	s.Status = &v
	return s
}

func (s *NeuronApp) Validate() error {
	return dara.Validate(s)
}
