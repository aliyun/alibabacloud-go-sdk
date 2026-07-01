// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeToRCSSignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackgroundImage(v string) *UpgradeToRCSSignatureRequest
	GetBackgroundImage() *string
	SetBubbleColor(v string) *UpgradeToRCSSignatureRequest
	GetBubbleColor() *string
	SetCategory(v int64) *UpgradeToRCSSignatureRequest
	GetCategory() *int64
	SetDescription(v string) *UpgradeToRCSSignatureRequest
	GetDescription() *string
	SetLatitude(v string) *UpgradeToRCSSignatureRequest
	GetLatitude() *string
	SetLogo(v string) *UpgradeToRCSSignatureRequest
	GetLogo() *string
	SetLongitude(v string) *UpgradeToRCSSignatureRequest
	GetLongitude() *string
	SetOfficeAddress(v string) *UpgradeToRCSSignatureRequest
	GetOfficeAddress() *string
	SetServiceEmail(v string) *UpgradeToRCSSignatureRequest
	GetServiceEmail() *string
	SetServicePhone(v string) *UpgradeToRCSSignatureRequest
	GetServicePhone() *string
	SetServiceTerms(v string) *UpgradeToRCSSignatureRequest
	GetServiceTerms() *string
	SetServiceWebsite(v string) *UpgradeToRCSSignatureRequest
	GetServiceWebsite() *string
	SetSignName(v string) *UpgradeToRCSSignatureRequest
	GetSignName() *string
}

type UpgradeToRCSSignatureRequest struct {
	// 背景图片ossKey
	//
	// example:
	//
	// 示例值示例值
	BackgroundImage *string `json:"BackgroundImage,omitempty" xml:"BackgroundImage,omitempty"`
	// 气泡颜色
	//
	// example:
	//
	// 示例值示例值示例值
	BubbleColor *string `json:"BubbleColor,omitempty" xml:"BubbleColor,omitempty"`
	// 行业类型
	//
	// example:
	//
	// 97
	Category *int64 `json:"Category,omitempty" xml:"Category,omitempty"`
	// 描述信息
	//
	// example:
	//
	// 示例值示例值
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 纬度
	//
	// example:
	//
	// 示例值示例值
	Latitude *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// logo图片ossKey
	//
	// example:
	//
	// 示例值示例值
	Logo *string `json:"Logo,omitempty" xml:"Logo,omitempty"`
	// 经度
	//
	// example:
	//
	// 示例值示例值
	Longitude *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	// 办公地址
	//
	// example:
	//
	// 示例值示例值
	OfficeAddress *string `json:"OfficeAddress,omitempty" xml:"OfficeAddress,omitempty"`
	// 服务邮箱
	//
	// example:
	//
	// 示例值示例值
	ServiceEmail *string `json:"ServiceEmail,omitempty" xml:"ServiceEmail,omitempty"`
	// 服务电话
	//
	// example:
	//
	// 示例值示例值示例值
	ServicePhone *string `json:"ServicePhone,omitempty" xml:"ServicePhone,omitempty"`
	// 服务条款URL
	//
	// example:
	//
	// 示例值示例值
	ServiceTerms *string `json:"ServiceTerms,omitempty" xml:"ServiceTerms,omitempty"`
	// 服务官网URL
	//
	// example:
	//
	// 示例值示例值
	ServiceWebsite *string `json:"ServiceWebsite,omitempty" xml:"ServiceWebsite,omitempty"`
	// 签名名称
	//
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s UpgradeToRCSSignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeToRCSSignatureRequest) GoString() string {
	return s.String()
}

func (s *UpgradeToRCSSignatureRequest) GetBackgroundImage() *string {
	return s.BackgroundImage
}

func (s *UpgradeToRCSSignatureRequest) GetBubbleColor() *string {
	return s.BubbleColor
}

func (s *UpgradeToRCSSignatureRequest) GetCategory() *int64 {
	return s.Category
}

func (s *UpgradeToRCSSignatureRequest) GetDescription() *string {
	return s.Description
}

func (s *UpgradeToRCSSignatureRequest) GetLatitude() *string {
	return s.Latitude
}

func (s *UpgradeToRCSSignatureRequest) GetLogo() *string {
	return s.Logo
}

func (s *UpgradeToRCSSignatureRequest) GetLongitude() *string {
	return s.Longitude
}

func (s *UpgradeToRCSSignatureRequest) GetOfficeAddress() *string {
	return s.OfficeAddress
}

func (s *UpgradeToRCSSignatureRequest) GetServiceEmail() *string {
	return s.ServiceEmail
}

func (s *UpgradeToRCSSignatureRequest) GetServicePhone() *string {
	return s.ServicePhone
}

func (s *UpgradeToRCSSignatureRequest) GetServiceTerms() *string {
	return s.ServiceTerms
}

func (s *UpgradeToRCSSignatureRequest) GetServiceWebsite() *string {
	return s.ServiceWebsite
}

func (s *UpgradeToRCSSignatureRequest) GetSignName() *string {
	return s.SignName
}

func (s *UpgradeToRCSSignatureRequest) SetBackgroundImage(v string) *UpgradeToRCSSignatureRequest {
	s.BackgroundImage = &v
	return s
}

func (s *UpgradeToRCSSignatureRequest) SetBubbleColor(v string) *UpgradeToRCSSignatureRequest {
	s.BubbleColor = &v
	return s
}

func (s *UpgradeToRCSSignatureRequest) SetCategory(v int64) *UpgradeToRCSSignatureRequest {
	s.Category = &v
	return s
}

func (s *UpgradeToRCSSignatureRequest) SetDescription(v string) *UpgradeToRCSSignatureRequest {
	s.Description = &v
	return s
}

func (s *UpgradeToRCSSignatureRequest) SetLatitude(v string) *UpgradeToRCSSignatureRequest {
	s.Latitude = &v
	return s
}

func (s *UpgradeToRCSSignatureRequest) SetLogo(v string) *UpgradeToRCSSignatureRequest {
	s.Logo = &v
	return s
}

func (s *UpgradeToRCSSignatureRequest) SetLongitude(v string) *UpgradeToRCSSignatureRequest {
	s.Longitude = &v
	return s
}

func (s *UpgradeToRCSSignatureRequest) SetOfficeAddress(v string) *UpgradeToRCSSignatureRequest {
	s.OfficeAddress = &v
	return s
}

func (s *UpgradeToRCSSignatureRequest) SetServiceEmail(v string) *UpgradeToRCSSignatureRequest {
	s.ServiceEmail = &v
	return s
}

func (s *UpgradeToRCSSignatureRequest) SetServicePhone(v string) *UpgradeToRCSSignatureRequest {
	s.ServicePhone = &v
	return s
}

func (s *UpgradeToRCSSignatureRequest) SetServiceTerms(v string) *UpgradeToRCSSignatureRequest {
	s.ServiceTerms = &v
	return s
}

func (s *UpgradeToRCSSignatureRequest) SetServiceWebsite(v string) *UpgradeToRCSSignatureRequest {
	s.ServiceWebsite = &v
	return s
}

func (s *UpgradeToRCSSignatureRequest) SetSignName(v string) *UpgradeToRCSSignatureRequest {
	s.SignName = &v
	return s
}

func (s *UpgradeToRCSSignatureRequest) Validate() error {
	return dara.Validate(s)
}
