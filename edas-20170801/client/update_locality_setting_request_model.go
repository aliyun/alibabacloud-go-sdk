// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocalitySettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateLocalitySettingRequest
	GetAppId() *string
	SetEnabled(v bool) *UpdateLocalitySettingRequest
	GetEnabled() *bool
	SetNamespaceId(v string) *UpdateLocalitySettingRequest
	GetNamespaceId() *string
	SetRegion(v string) *UpdateLocalitySettingRequest
	GetRegion() *string
	SetThreshold(v float32) *UpdateLocalitySettingRequest
	GetThreshold() *float32
}

type UpdateLocalitySettingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bfa00cfb-9642-4292-bb78-1d7d4c86004c
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 15
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateLocalitySettingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocalitySettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateLocalitySettingRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateLocalitySettingRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateLocalitySettingRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateLocalitySettingRequest) GetRegion() *string {
	return s.Region
}

func (s *UpdateLocalitySettingRequest) GetThreshold() *float32 {
	return s.Threshold
}

func (s *UpdateLocalitySettingRequest) SetAppId(v string) *UpdateLocalitySettingRequest {
	s.AppId = &v
	return s
}

func (s *UpdateLocalitySettingRequest) SetEnabled(v bool) *UpdateLocalitySettingRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateLocalitySettingRequest) SetNamespaceId(v string) *UpdateLocalitySettingRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateLocalitySettingRequest) SetRegion(v string) *UpdateLocalitySettingRequest {
	s.Region = &v
	return s
}

func (s *UpdateLocalitySettingRequest) SetThreshold(v float32) *UpdateLocalitySettingRequest {
	s.Threshold = &v
	return s
}

func (s *UpdateLocalitySettingRequest) Validate() error {
	return dara.Validate(s)
}
