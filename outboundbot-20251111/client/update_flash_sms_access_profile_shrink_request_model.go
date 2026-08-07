// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlashSmsAccessProfileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessProfileShrink(v string) *UpdateFlashSmsAccessProfileShrinkRequest
	GetAccessProfileShrink() *string
	SetAccessProfileId(v string) *UpdateFlashSmsAccessProfileShrinkRequest
	GetAccessProfileId() *string
	SetInstanceId(v string) *UpdateFlashSmsAccessProfileShrinkRequest
	GetInstanceId() *string
	SetProviderId(v string) *UpdateFlashSmsAccessProfileShrinkRequest
	GetProviderId() *string
}

type UpdateFlashSmsAccessProfileShrinkRequest struct {
	// 接入配置
	AccessProfileShrink *string `json:"AccessProfile,omitempty" xml:"AccessProfile,omitempty"`
	// 接入配置ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// 实例ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 供应商ID
	//
	// example:
	//
	// Uincall
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
}

func (s UpdateFlashSmsAccessProfileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlashSmsAccessProfileShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlashSmsAccessProfileShrinkRequest) GetAccessProfileShrink() *string {
	return s.AccessProfileShrink
}

func (s *UpdateFlashSmsAccessProfileShrinkRequest) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *UpdateFlashSmsAccessProfileShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateFlashSmsAccessProfileShrinkRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *UpdateFlashSmsAccessProfileShrinkRequest) SetAccessProfileShrink(v string) *UpdateFlashSmsAccessProfileShrinkRequest {
	s.AccessProfileShrink = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileShrinkRequest) SetAccessProfileId(v string) *UpdateFlashSmsAccessProfileShrinkRequest {
	s.AccessProfileId = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileShrinkRequest) SetInstanceId(v string) *UpdateFlashSmsAccessProfileShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileShrinkRequest) SetProviderId(v string) *UpdateFlashSmsAccessProfileShrinkRequest {
	s.ProviderId = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
