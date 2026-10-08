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
	// The ID of the application. You can call the [ListApplication](https://help.aliyun.com/document_detail/149390.html) operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bfa00cfb-9642-4292-bb78-1d7d4c86004c
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether the setting is active:
	//
	// - true: The setting is active.
	//
	// - false: The setting is not active.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the namespace. This ID cannot be changed after the namespace is created. The format is [unk]physical space identifier[unk].
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The ID of the region where the elastic compute unit (ECU) is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The total number of items that satisfy the threshold expression.
	//
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
