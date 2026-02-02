// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateChatSessionRequest
	GetInstanceId() *string
	SetLicense(v string) *CreateChatSessionRequest
	GetLicense() *string
	SetPlatform(v string) *CreateChatSessionRequest
	GetPlatform() *string
}

type CreateChatSessionRequest struct {
	// 需要在[数字人实时交互服务](https://common-buy.aliyun.com/?spm=a2c4g.11186623.0.0.457876812ETi6y&commodityCode=avatar_2dchat_public_cn)购买完成对应的服务购买，当前有可用的服务时，前往阿里云-[我的订单](https://billing-cost.console.aliyun.com/order/list)页面对应订单详情下进行查询
	//
	// This parameter is required.
	//
	// example:
	//
	// rmq-cn-9t946y43m1d
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 灵眸平台颁发的个人凭证（在使用端渲染数字人的场景下必填）。
	//
	// example:
	//
	// b9be4b25c2d38c409c376ffd2372be1
	License *string `json:"license,omitempty" xml:"license,omitempty"`
	// 运行SDK的平台（在使用端渲染数字人的场景下必填）。
	//
	// example:
	//
	// Web | Android | iOS
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
}

func (s CreateChatSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateChatSessionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateChatSessionRequest) GetLicense() *string {
	return s.License
}

func (s *CreateChatSessionRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateChatSessionRequest) SetInstanceId(v string) *CreateChatSessionRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateChatSessionRequest) SetLicense(v string) *CreateChatSessionRequest {
	s.License = &v
	return s
}

func (s *CreateChatSessionRequest) SetPlatform(v string) *CreateChatSessionRequest {
	s.Platform = &v
	return s
}

func (s *CreateChatSessionRequest) Validate() error {
	return dara.Validate(s)
}
