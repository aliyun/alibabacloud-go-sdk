// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppInstanceProfile interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *AppInstanceProfile
	GetApplicationType() *string
	SetApplicationTypeText(v string) *AppInstanceProfile
	GetApplicationTypeText() *string
	SetBizId(v string) *AppInstanceProfile
	GetBizId() *string
	SetCommodityCode(v string) *AppInstanceProfile
	GetCommodityCode() *string
	SetCustomerService(v string) *AppInstanceProfile
	GetCustomerService() *string
	SetDeployArea(v string) *AppInstanceProfile
	GetDeployArea() *string
	SetInstanceId(v string) *AppInstanceProfile
	GetInstanceId() *string
	SetOrdTime(v string) *AppInstanceProfile
	GetOrdTime() *string
	SetOrderId(v string) *AppInstanceProfile
	GetOrderId() *string
	SetPayTime(v string) *AppInstanceProfile
	GetPayTime() *string
	SetSeoSite(v string) *AppInstanceProfile
	GetSeoSite() *string
	SetSiteVersion(v string) *AppInstanceProfile
	GetSiteVersion() *string
	SetSiteVersionText(v string) *AppInstanceProfile
	GetSiteVersionText() *string
	SetSource(v string) *AppInstanceProfile
	GetSource() *string
	SetTemplateEtag(v string) *AppInstanceProfile
	GetTemplateEtag() *string
	SetTemplateId(v string) *AppInstanceProfile
	GetTemplateId() *string
}

type AppInstanceProfile struct {
	ApplicationType     *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	ApplicationTypeText *string `json:"ApplicationTypeText,omitempty" xml:"ApplicationTypeText,omitempty"`
	BizId               *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CommodityCode       *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CustomerService     *string `json:"CustomerService,omitempty" xml:"CustomerService,omitempty"`
	DeployArea          *string `json:"DeployArea,omitempty" xml:"DeployArea,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrdTime             *string `json:"OrdTime,omitempty" xml:"OrdTime,omitempty"`
	OrderId             *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PayTime             *string `json:"PayTime,omitempty" xml:"PayTime,omitempty"`
	SeoSite             *string `json:"SeoSite,omitempty" xml:"SeoSite,omitempty"`
	SiteVersion         *string `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	SiteVersionText     *string `json:"SiteVersionText,omitempty" xml:"SiteVersionText,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TemplateEtag        *string `json:"TemplateEtag,omitempty" xml:"TemplateEtag,omitempty"`
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s AppInstanceProfile) String() string {
	return dara.Prettify(s)
}

func (s AppInstanceProfile) GoString() string {
	return s.String()
}

func (s *AppInstanceProfile) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *AppInstanceProfile) GetApplicationTypeText() *string {
	return s.ApplicationTypeText
}

func (s *AppInstanceProfile) GetBizId() *string {
	return s.BizId
}

func (s *AppInstanceProfile) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *AppInstanceProfile) GetCustomerService() *string {
	return s.CustomerService
}

func (s *AppInstanceProfile) GetDeployArea() *string {
	return s.DeployArea
}

func (s *AppInstanceProfile) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AppInstanceProfile) GetOrdTime() *string {
	return s.OrdTime
}

func (s *AppInstanceProfile) GetOrderId() *string {
	return s.OrderId
}

func (s *AppInstanceProfile) GetPayTime() *string {
	return s.PayTime
}

func (s *AppInstanceProfile) GetSeoSite() *string {
	return s.SeoSite
}

func (s *AppInstanceProfile) GetSiteVersion() *string {
	return s.SiteVersion
}

func (s *AppInstanceProfile) GetSiteVersionText() *string {
	return s.SiteVersionText
}

func (s *AppInstanceProfile) GetSource() *string {
	return s.Source
}

func (s *AppInstanceProfile) GetTemplateEtag() *string {
	return s.TemplateEtag
}

func (s *AppInstanceProfile) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AppInstanceProfile) SetApplicationType(v string) *AppInstanceProfile {
	s.ApplicationType = &v
	return s
}

func (s *AppInstanceProfile) SetApplicationTypeText(v string) *AppInstanceProfile {
	s.ApplicationTypeText = &v
	return s
}

func (s *AppInstanceProfile) SetBizId(v string) *AppInstanceProfile {
	s.BizId = &v
	return s
}

func (s *AppInstanceProfile) SetCommodityCode(v string) *AppInstanceProfile {
	s.CommodityCode = &v
	return s
}

func (s *AppInstanceProfile) SetCustomerService(v string) *AppInstanceProfile {
	s.CustomerService = &v
	return s
}

func (s *AppInstanceProfile) SetDeployArea(v string) *AppInstanceProfile {
	s.DeployArea = &v
	return s
}

func (s *AppInstanceProfile) SetInstanceId(v string) *AppInstanceProfile {
	s.InstanceId = &v
	return s
}

func (s *AppInstanceProfile) SetOrdTime(v string) *AppInstanceProfile {
	s.OrdTime = &v
	return s
}

func (s *AppInstanceProfile) SetOrderId(v string) *AppInstanceProfile {
	s.OrderId = &v
	return s
}

func (s *AppInstanceProfile) SetPayTime(v string) *AppInstanceProfile {
	s.PayTime = &v
	return s
}

func (s *AppInstanceProfile) SetSeoSite(v string) *AppInstanceProfile {
	s.SeoSite = &v
	return s
}

func (s *AppInstanceProfile) SetSiteVersion(v string) *AppInstanceProfile {
	s.SiteVersion = &v
	return s
}

func (s *AppInstanceProfile) SetSiteVersionText(v string) *AppInstanceProfile {
	s.SiteVersionText = &v
	return s
}

func (s *AppInstanceProfile) SetSource(v string) *AppInstanceProfile {
	s.Source = &v
	return s
}

func (s *AppInstanceProfile) SetTemplateEtag(v string) *AppInstanceProfile {
	s.TemplateEtag = &v
	return s
}

func (s *AppInstanceProfile) SetTemplateId(v string) *AppInstanceProfile {
	s.TemplateId = &v
	return s
}

func (s *AppInstanceProfile) Validate() error {
	return dara.Validate(s)
}
