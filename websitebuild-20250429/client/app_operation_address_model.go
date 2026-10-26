// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppOperationAddress interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*AppOperateAction) *AppOperationAddress
	GetActions() []*AppOperateAction
	SetAiCustomerConfigUrl(v string) *AppOperationAddress
	GetAiCustomerConfigUrl() *string
	SetAiDesignUrl(v string) *AppOperationAddress
	GetAiDesignUrl() *string
	SetAppPublishUrl(v string) *AppOperationAddress
	GetAppPublishUrl() *string
	SetDashboardActions(v []*AppOperateAction) *AppOperationAddress
	GetDashboardActions() []*AppOperateAction
	SetDesignUrl(v string) *AppOperationAddress
	GetDesignUrl() *string
	SetInstanceLoginUrl(v string) *AppOperationAddress
	GetInstanceLoginUrl() *string
	SetRenewBuyUrl(v string) *AppOperationAddress
	GetRenewBuyUrl() *string
	SetServerDeliveryUrl(v string) *AppOperationAddress
	GetServerDeliveryUrl() *string
	SetUpgradeBuyUrl(v string) *AppOperationAddress
	GetUpgradeBuyUrl() *string
}

type AppOperationAddress struct {
	// The action button list.
	Actions []*AppOperateAction `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The customer service design URL.
	AiCustomerConfigUrl *string `json:"AiCustomerConfigUrl,omitempty" xml:"AiCustomerConfigUrl,omitempty"`
	// The AI design URL.
	AiDesignUrl *string `json:"AiDesignUrl,omitempty" xml:"AiDesignUrl,omitempty"`
	// The application publishing URL.
	//
	// example:
	//
	// https://abcd.scd.wanwang.xin
	AppPublishUrl *string `json:"AppPublishUrl,omitempty" xml:"AppPublishUrl,omitempty"`
	// The overview page action button list.
	DashboardActions []*AppOperateAction `json:"DashboardActions,omitempty" xml:"DashboardActions,omitempty" type:"Repeated"`
	// The professional design URL.
	//
	// example:
	//
	// https://zero.msea.aliyun.com/forward/v1/website/websiteLogin?bizId=WS20250915163734000001&callBackUrl=/Designer/index
	DesignUrl *string `json:"DesignUrl,omitempty" xml:"DesignUrl,omitempty"`
	// The application instance logon URL.
	//
	// example:
	//
	// https://zero.msea.aliyun.com/forward/v1/website/websiteLogin?bizId=WS20250915163734000001&callBackUrl=/Admin/SiteAdmin/Index
	InstanceLoginUrl *string `json:"InstanceLoginUrl,omitempty" xml:"InstanceLoginUrl,omitempty"`
	// The renewal purchase URL.
	//
	// example:
	//
	// https://common-buy.aliyun.com/?orderType=RENEW&commodityCode=domain_website_public_cn&instanceId=12313213
	RenewBuyUrl *string `json:"RenewBuyUrl,omitempty" xml:"RenewBuyUrl,omitempty"`
	// The service delivery URL.
	//
	// example:
	//
	// https://zero.msea.aliyun.com/forward/v1/website/websiteLogin?bizId=WS20250915163734000001&callBackUrl=/Admin/SiteAdmin/Index
	ServerDeliveryUrl *string `json:"ServerDeliveryUrl,omitempty" xml:"ServerDeliveryUrl,omitempty"`
	// The upgrade purchase URL.
	//
	// example:
	//
	// https://common-buy.aliyun.com/?orderType=UPGRADE&commodityCode=domain_website_public_cn&instanceId=12313213
	UpgradeBuyUrl *string `json:"UpgradeBuyUrl,omitempty" xml:"UpgradeBuyUrl,omitempty"`
}

func (s AppOperationAddress) String() string {
	return dara.Prettify(s)
}

func (s AppOperationAddress) GoString() string {
	return s.String()
}

func (s *AppOperationAddress) GetActions() []*AppOperateAction {
	return s.Actions
}

func (s *AppOperationAddress) GetAiCustomerConfigUrl() *string {
	return s.AiCustomerConfigUrl
}

func (s *AppOperationAddress) GetAiDesignUrl() *string {
	return s.AiDesignUrl
}

func (s *AppOperationAddress) GetAppPublishUrl() *string {
	return s.AppPublishUrl
}

func (s *AppOperationAddress) GetDashboardActions() []*AppOperateAction {
	return s.DashboardActions
}

func (s *AppOperationAddress) GetDesignUrl() *string {
	return s.DesignUrl
}

func (s *AppOperationAddress) GetInstanceLoginUrl() *string {
	return s.InstanceLoginUrl
}

func (s *AppOperationAddress) GetRenewBuyUrl() *string {
	return s.RenewBuyUrl
}

func (s *AppOperationAddress) GetServerDeliveryUrl() *string {
	return s.ServerDeliveryUrl
}

func (s *AppOperationAddress) GetUpgradeBuyUrl() *string {
	return s.UpgradeBuyUrl
}

func (s *AppOperationAddress) SetActions(v []*AppOperateAction) *AppOperationAddress {
	s.Actions = v
	return s
}

func (s *AppOperationAddress) SetAiCustomerConfigUrl(v string) *AppOperationAddress {
	s.AiCustomerConfigUrl = &v
	return s
}

func (s *AppOperationAddress) SetAiDesignUrl(v string) *AppOperationAddress {
	s.AiDesignUrl = &v
	return s
}

func (s *AppOperationAddress) SetAppPublishUrl(v string) *AppOperationAddress {
	s.AppPublishUrl = &v
	return s
}

func (s *AppOperationAddress) SetDashboardActions(v []*AppOperateAction) *AppOperationAddress {
	s.DashboardActions = v
	return s
}

func (s *AppOperationAddress) SetDesignUrl(v string) *AppOperationAddress {
	s.DesignUrl = &v
	return s
}

func (s *AppOperationAddress) SetInstanceLoginUrl(v string) *AppOperationAddress {
	s.InstanceLoginUrl = &v
	return s
}

func (s *AppOperationAddress) SetRenewBuyUrl(v string) *AppOperationAddress {
	s.RenewBuyUrl = &v
	return s
}

func (s *AppOperationAddress) SetServerDeliveryUrl(v string) *AppOperationAddress {
	s.ServerDeliveryUrl = &v
	return s
}

func (s *AppOperationAddress) SetUpgradeBuyUrl(v string) *AppOperationAddress {
	s.UpgradeBuyUrl = &v
	return s
}

func (s *AppOperationAddress) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DashboardActions != nil {
		for _, item := range s.DashboardActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
