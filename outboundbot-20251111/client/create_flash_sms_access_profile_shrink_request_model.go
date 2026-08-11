// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlashSmsAccessProfileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessProfileShrink(v string) *CreateFlashSmsAccessProfileShrinkRequest
	GetAccessProfileShrink() *string
	SetInstanceId(v string) *CreateFlashSmsAccessProfileShrinkRequest
	GetInstanceId() *string
	SetProviderId(v string) *CreateFlashSmsAccessProfileShrinkRequest
	GetProviderId() *string
}

type CreateFlashSmsAccessProfileShrinkRequest struct {
	// The access configuration.
	AccessProfileShrink *string `json:"AccessProfile,omitempty" xml:"AccessProfile,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The provider ID. Valid values:
	//
	// - Uincall: Beijing Youyin Communication Co., Ltd.
	//
	// - ChuangLan: Beijing Chuanglan Yunzhi Information Co., Ltd.
	//
	// - ChinaMobile: China Mobile.
	//
	// - ShangHaiTianNan: Shanghai Tiannan.
	//
	// - HeDao: Galaxis.
	//
	// - DySms: Alibaba Communication.
	//
	// example:
	//
	// Uincall
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
}

func (s CreateFlashSmsAccessProfileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlashSmsAccessProfileShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFlashSmsAccessProfileShrinkRequest) GetAccessProfileShrink() *string {
	return s.AccessProfileShrink
}

func (s *CreateFlashSmsAccessProfileShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateFlashSmsAccessProfileShrinkRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *CreateFlashSmsAccessProfileShrinkRequest) SetAccessProfileShrink(v string) *CreateFlashSmsAccessProfileShrinkRequest {
	s.AccessProfileShrink = &v
	return s
}

func (s *CreateFlashSmsAccessProfileShrinkRequest) SetInstanceId(v string) *CreateFlashSmsAccessProfileShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFlashSmsAccessProfileShrinkRequest) SetProviderId(v string) *CreateFlashSmsAccessProfileShrinkRequest {
	s.ProviderId = &v
	return s
}

func (s *CreateFlashSmsAccessProfileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
