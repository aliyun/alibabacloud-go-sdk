// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *UpdateAppInstanceRequest
	GetApplicationType() *string
	SetBizId(v string) *UpdateAppInstanceRequest
	GetBizId() *string
	SetClientToken(v string) *UpdateAppInstanceRequest
	GetClientToken() *string
	SetDeployArea(v string) *UpdateAppInstanceRequest
	GetDeployArea() *string
	SetDescription(v string) *UpdateAppInstanceRequest
	GetDescription() *string
	SetExtend(v string) *UpdateAppInstanceRequest
	GetExtend() *string
	SetIconUrl(v string) *UpdateAppInstanceRequest
	GetIconUrl() *string
	SetName(v string) *UpdateAppInstanceRequest
	GetName() *string
	SetPaymentType(v string) *UpdateAppInstanceRequest
	GetPaymentType() *string
	SetResourceGroupId(v string) *UpdateAppInstanceRequest
	GetResourceGroupId() *string
	SetSiteVersion(v string) *UpdateAppInstanceRequest
	GetSiteVersion() *string
	SetTags(v []*UpdateAppInstanceRequestTags) *UpdateAppInstanceRequest
	GetTags() []*UpdateAppInstanceRequestTags
	SetThumbnailUrl(v string) *UpdateAppInstanceRequest
	GetThumbnailUrl() *string
}

type UpdateAppInstanceRequest struct {
	// The application type.
	//
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The business ID.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Generate a unique value from your client. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The deployment region.
	//
	// example:
	//
	// ChineseMainland
	DeployArea *string `json:"DeployArea,omitempty" xml:"DeployArea,omitempty"`
	// The application description.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extended information.
	//
	// example:
	//
	// {\\"user_device_id\\":\\"6bef45cb0c76de284d24de074c088b73\\"}\\n
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The application icon URL.
	//
	// example:
	//
	// https://app-center-icon-prod-shanghai.oss-cn-shanghai.aliyuncs.com/tenant/1864953777494693/1753841032702_WX20250729-171155%402x.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// The application name.
	//
	// example:
	//
	// 大角鹿
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The payment type.
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
	// The website version.
	//
	// example:
	//
	// Basic_Edition
	SiteVersion *string `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// The tags.
	Tags []*UpdateAppInstanceRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The application thumbnail URL.
	//
	// example:
	//
	// http://www.aliyun.com
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
}

func (s UpdateAppInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *UpdateAppInstanceRequest) GetBizId() *string {
	return s.BizId
}

func (s *UpdateAppInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAppInstanceRequest) GetDeployArea() *string {
	return s.DeployArea
}

func (s *UpdateAppInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAppInstanceRequest) GetExtend() *string {
	return s.Extend
}

func (s *UpdateAppInstanceRequest) GetIconUrl() *string {
	return s.IconUrl
}

func (s *UpdateAppInstanceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAppInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *UpdateAppInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateAppInstanceRequest) GetSiteVersion() *string {
	return s.SiteVersion
}

func (s *UpdateAppInstanceRequest) GetTags() []*UpdateAppInstanceRequestTags {
	return s.Tags
}

func (s *UpdateAppInstanceRequest) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *UpdateAppInstanceRequest) SetApplicationType(v string) *UpdateAppInstanceRequest {
	s.ApplicationType = &v
	return s
}

func (s *UpdateAppInstanceRequest) SetBizId(v string) *UpdateAppInstanceRequest {
	s.BizId = &v
	return s
}

func (s *UpdateAppInstanceRequest) SetClientToken(v string) *UpdateAppInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAppInstanceRequest) SetDeployArea(v string) *UpdateAppInstanceRequest {
	s.DeployArea = &v
	return s
}

func (s *UpdateAppInstanceRequest) SetDescription(v string) *UpdateAppInstanceRequest {
	s.Description = &v
	return s
}

func (s *UpdateAppInstanceRequest) SetExtend(v string) *UpdateAppInstanceRequest {
	s.Extend = &v
	return s
}

func (s *UpdateAppInstanceRequest) SetIconUrl(v string) *UpdateAppInstanceRequest {
	s.IconUrl = &v
	return s
}

func (s *UpdateAppInstanceRequest) SetName(v string) *UpdateAppInstanceRequest {
	s.Name = &v
	return s
}

func (s *UpdateAppInstanceRequest) SetPaymentType(v string) *UpdateAppInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *UpdateAppInstanceRequest) SetResourceGroupId(v string) *UpdateAppInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateAppInstanceRequest) SetSiteVersion(v string) *UpdateAppInstanceRequest {
	s.SiteVersion = &v
	return s
}

func (s *UpdateAppInstanceRequest) SetTags(v []*UpdateAppInstanceRequestTags) *UpdateAppInstanceRequest {
	s.Tags = v
	return s
}

func (s *UpdateAppInstanceRequest) SetThumbnailUrl(v string) *UpdateAppInstanceRequest {
	s.ThumbnailUrl = &v
	return s
}

func (s *UpdateAppInstanceRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAppInstanceRequestTags struct {
	// The tag key of the instance. Valid values of N: **1*	- to **20**. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// Group
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the instance. Valid values of N: **1*	- to **20**. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// 15
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s UpdateAppInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceRequestTags) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateAppInstanceRequestTags) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateAppInstanceRequestTags) SetTagKey(v string) *UpdateAppInstanceRequestTags {
	s.TagKey = &v
	return s
}

func (s *UpdateAppInstanceRequestTags) SetTagValue(v string) *UpdateAppInstanceRequestTags {
	s.TagValue = &v
	return s
}

func (s *UpdateAppInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}
