// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRCSSignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackgroundImage(v string) *UpdateRCSSignatureRequest
	GetBackgroundImage() *string
	SetBubbleColor(v string) *UpdateRCSSignatureRequest
	GetBubbleColor() *string
	SetCategory(v int64) *UpdateRCSSignatureRequest
	GetCategory() *int64
	SetDescription(v string) *UpdateRCSSignatureRequest
	GetDescription() *string
	SetLatitude(v string) *UpdateRCSSignatureRequest
	GetLatitude() *string
	SetLogo(v string) *UpdateRCSSignatureRequest
	GetLogo() *string
	SetLongitude(v string) *UpdateRCSSignatureRequest
	GetLongitude() *string
	SetOfficeAddress(v string) *UpdateRCSSignatureRequest
	GetOfficeAddress() *string
	SetServiceEmail(v string) *UpdateRCSSignatureRequest
	GetServiceEmail() *string
	SetServicePhone(v string) *UpdateRCSSignatureRequest
	GetServicePhone() *string
	SetServiceTerms(v string) *UpdateRCSSignatureRequest
	GetServiceTerms() *string
	SetServiceWebsite(v string) *UpdateRCSSignatureRequest
	GetServiceWebsite() *string
	SetSignName(v string) *UpdateRCSSignatureRequest
	GetSignName() *string
}

type UpdateRCSSignatureRequest struct {
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
	// 24
	Category *int64 `json:"Category,omitempty" xml:"Category,omitempty"`
	// 描述信息
	//
	// example:
	//
	// 示例值示例值示例值
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
	// 示例值示例值示例值
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
	// 示例值
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
	// 示例值示例值
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
	// 签名名称（用于定位5G签名）
	//
	// example:
	//
	// 示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s UpdateRCSSignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRCSSignatureRequest) GoString() string {
	return s.String()
}

func (s *UpdateRCSSignatureRequest) GetBackgroundImage() *string {
	return s.BackgroundImage
}

func (s *UpdateRCSSignatureRequest) GetBubbleColor() *string {
	return s.BubbleColor
}

func (s *UpdateRCSSignatureRequest) GetCategory() *int64 {
	return s.Category
}

func (s *UpdateRCSSignatureRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRCSSignatureRequest) GetLatitude() *string {
	return s.Latitude
}

func (s *UpdateRCSSignatureRequest) GetLogo() *string {
	return s.Logo
}

func (s *UpdateRCSSignatureRequest) GetLongitude() *string {
	return s.Longitude
}

func (s *UpdateRCSSignatureRequest) GetOfficeAddress() *string {
	return s.OfficeAddress
}

func (s *UpdateRCSSignatureRequest) GetServiceEmail() *string {
	return s.ServiceEmail
}

func (s *UpdateRCSSignatureRequest) GetServicePhone() *string {
	return s.ServicePhone
}

func (s *UpdateRCSSignatureRequest) GetServiceTerms() *string {
	return s.ServiceTerms
}

func (s *UpdateRCSSignatureRequest) GetServiceWebsite() *string {
	return s.ServiceWebsite
}

func (s *UpdateRCSSignatureRequest) GetSignName() *string {
	return s.SignName
}

func (s *UpdateRCSSignatureRequest) SetBackgroundImage(v string) *UpdateRCSSignatureRequest {
	s.BackgroundImage = &v
	return s
}

func (s *UpdateRCSSignatureRequest) SetBubbleColor(v string) *UpdateRCSSignatureRequest {
	s.BubbleColor = &v
	return s
}

func (s *UpdateRCSSignatureRequest) SetCategory(v int64) *UpdateRCSSignatureRequest {
	s.Category = &v
	return s
}

func (s *UpdateRCSSignatureRequest) SetDescription(v string) *UpdateRCSSignatureRequest {
	s.Description = &v
	return s
}

func (s *UpdateRCSSignatureRequest) SetLatitude(v string) *UpdateRCSSignatureRequest {
	s.Latitude = &v
	return s
}

func (s *UpdateRCSSignatureRequest) SetLogo(v string) *UpdateRCSSignatureRequest {
	s.Logo = &v
	return s
}

func (s *UpdateRCSSignatureRequest) SetLongitude(v string) *UpdateRCSSignatureRequest {
	s.Longitude = &v
	return s
}

func (s *UpdateRCSSignatureRequest) SetOfficeAddress(v string) *UpdateRCSSignatureRequest {
	s.OfficeAddress = &v
	return s
}

func (s *UpdateRCSSignatureRequest) SetServiceEmail(v string) *UpdateRCSSignatureRequest {
	s.ServiceEmail = &v
	return s
}

func (s *UpdateRCSSignatureRequest) SetServicePhone(v string) *UpdateRCSSignatureRequest {
	s.ServicePhone = &v
	return s
}

func (s *UpdateRCSSignatureRequest) SetServiceTerms(v string) *UpdateRCSSignatureRequest {
	s.ServiceTerms = &v
	return s
}

func (s *UpdateRCSSignatureRequest) SetServiceWebsite(v string) *UpdateRCSSignatureRequest {
	s.ServiceWebsite = &v
	return s
}

func (s *UpdateRCSSignatureRequest) SetSignName(v string) *UpdateRCSSignatureRequest {
	s.SignName = &v
	return s
}

func (s *UpdateRCSSignatureRequest) Validate() error {
	return dara.Validate(s)
}
