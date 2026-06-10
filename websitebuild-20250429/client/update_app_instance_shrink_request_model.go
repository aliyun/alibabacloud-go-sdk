// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *UpdateAppInstanceShrinkRequest
	GetApplicationType() *string
	SetBizId(v string) *UpdateAppInstanceShrinkRequest
	GetBizId() *string
	SetClientToken(v string) *UpdateAppInstanceShrinkRequest
	GetClientToken() *string
	SetDeployArea(v string) *UpdateAppInstanceShrinkRequest
	GetDeployArea() *string
	SetDescription(v string) *UpdateAppInstanceShrinkRequest
	GetDescription() *string
	SetExtend(v string) *UpdateAppInstanceShrinkRequest
	GetExtend() *string
	SetIconUrl(v string) *UpdateAppInstanceShrinkRequest
	GetIconUrl() *string
	SetName(v string) *UpdateAppInstanceShrinkRequest
	GetName() *string
	SetPaymentType(v string) *UpdateAppInstanceShrinkRequest
	GetPaymentType() *string
	SetResourceGroupId(v string) *UpdateAppInstanceShrinkRequest
	GetResourceGroupId() *string
	SetSiteVersion(v string) *UpdateAppInstanceShrinkRequest
	GetSiteVersion() *string
	SetTagsShrink(v string) *UpdateAppInstanceShrinkRequest
	GetTagsShrink() *string
	SetThumbnailUrl(v string) *UpdateAppInstanceShrinkRequest
	GetThumbnailUrl() *string
}

type UpdateAppInstanceShrinkRequest struct {
	// Application type
	//
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// Business ID
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Ensures the idempotence of the request. Generate a unique value from your client for this parameter to guarantee uniqueness across different requests. ClientToken supports only ASCII characters and must not exceed 64 characters.
	//
	// example:
	//
	// 123
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Deployment area
	//
	// example:
	//
	// ChineseMainland
	DeployArea *string `json:"DeployArea,omitempty" xml:"DeployArea,omitempty"`
	// Application description
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Extension information
	//
	// example:
	//
	// {\\"user_device_id\\":\\"6bef45cb0c76de284d24de074c088b73\\"}\\n
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// Application icon
	//
	// example:
	//
	// https://app-center-icon-prod-shanghai.oss-cn-shanghai.aliyuncs.com/tenant/1864953777494693/1753841032702_WX20250729-171155%402x.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// Application name
	//
	// example:
	//
	// 大角鹿
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Payment type
	//
	// example:
	//
	// AUTO_PAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aek3cqkrqibqkby
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Website version
	//
	// example:
	//
	// Basic_Edition
	SiteVersion *string `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Application thumbnail
	//
	// example:
	//
	// http://www.aliyun.com
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
}

func (s UpdateAppInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceShrinkRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *UpdateAppInstanceShrinkRequest) GetBizId() *string {
	return s.BizId
}

func (s *UpdateAppInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAppInstanceShrinkRequest) GetDeployArea() *string {
	return s.DeployArea
}

func (s *UpdateAppInstanceShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAppInstanceShrinkRequest) GetExtend() *string {
	return s.Extend
}

func (s *UpdateAppInstanceShrinkRequest) GetIconUrl() *string {
	return s.IconUrl
}

func (s *UpdateAppInstanceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAppInstanceShrinkRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *UpdateAppInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateAppInstanceShrinkRequest) GetSiteVersion() *string {
	return s.SiteVersion
}

func (s *UpdateAppInstanceShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdateAppInstanceShrinkRequest) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *UpdateAppInstanceShrinkRequest) SetApplicationType(v string) *UpdateAppInstanceShrinkRequest {
	s.ApplicationType = &v
	return s
}

func (s *UpdateAppInstanceShrinkRequest) SetBizId(v string) *UpdateAppInstanceShrinkRequest {
	s.BizId = &v
	return s
}

func (s *UpdateAppInstanceShrinkRequest) SetClientToken(v string) *UpdateAppInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAppInstanceShrinkRequest) SetDeployArea(v string) *UpdateAppInstanceShrinkRequest {
	s.DeployArea = &v
	return s
}

func (s *UpdateAppInstanceShrinkRequest) SetDescription(v string) *UpdateAppInstanceShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateAppInstanceShrinkRequest) SetExtend(v string) *UpdateAppInstanceShrinkRequest {
	s.Extend = &v
	return s
}

func (s *UpdateAppInstanceShrinkRequest) SetIconUrl(v string) *UpdateAppInstanceShrinkRequest {
	s.IconUrl = &v
	return s
}

func (s *UpdateAppInstanceShrinkRequest) SetName(v string) *UpdateAppInstanceShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateAppInstanceShrinkRequest) SetPaymentType(v string) *UpdateAppInstanceShrinkRequest {
	s.PaymentType = &v
	return s
}

func (s *UpdateAppInstanceShrinkRequest) SetResourceGroupId(v string) *UpdateAppInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateAppInstanceShrinkRequest) SetSiteVersion(v string) *UpdateAppInstanceShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *UpdateAppInstanceShrinkRequest) SetTagsShrink(v string) *UpdateAppInstanceShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateAppInstanceShrinkRequest) SetThumbnailUrl(v string) *UpdateAppInstanceShrinkRequest {
	s.ThumbnailUrl = &v
	return s
}

func (s *UpdateAppInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
