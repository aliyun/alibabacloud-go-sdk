// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceModuleStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureKey(v string) *CheckInstanceModuleStatusRequest
	GetFeatureKey() *string
	SetInstanceId(v string) *CheckInstanceModuleStatusRequest
	GetInstanceId() *string
	SetModuleKey(v string) *CheckInstanceModuleStatusRequest
	GetModuleKey() *string
	SetSubFeatureKey(v string) *CheckInstanceModuleStatusRequest
	GetSubFeatureKey() *string
}

type CheckInstanceModuleStatusRequest struct {
	// 二级模块标识
	//
	// example:
	//
	// urn:alibaba:idaas:license:module:ud:customField
	FeatureKey *string `json:"FeatureKey,omitempty" xml:"FeatureKey,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 一级模块标识，必填
	//
	// This parameter is required.
	//
	// example:
	//
	// urn:alibaba:idaas:license:module:ud
	ModuleKey *string `json:"ModuleKey,omitempty" xml:"ModuleKey,omitempty"`
	// example:
	//
	// urn:alibaba:idaas:license:module:idp:alibaba:dingtalk:pull:advanced_configuration
	SubFeatureKey *string `json:"SubFeatureKey,omitempty" xml:"SubFeatureKey,omitempty"`
}

func (s CheckInstanceModuleStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceModuleStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckInstanceModuleStatusRequest) GetFeatureKey() *string {
	return s.FeatureKey
}

func (s *CheckInstanceModuleStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CheckInstanceModuleStatusRequest) GetModuleKey() *string {
	return s.ModuleKey
}

func (s *CheckInstanceModuleStatusRequest) GetSubFeatureKey() *string {
	return s.SubFeatureKey
}

func (s *CheckInstanceModuleStatusRequest) SetFeatureKey(v string) *CheckInstanceModuleStatusRequest {
	s.FeatureKey = &v
	return s
}

func (s *CheckInstanceModuleStatusRequest) SetInstanceId(v string) *CheckInstanceModuleStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *CheckInstanceModuleStatusRequest) SetModuleKey(v string) *CheckInstanceModuleStatusRequest {
	s.ModuleKey = &v
	return s
}

func (s *CheckInstanceModuleStatusRequest) SetSubFeatureKey(v string) *CheckInstanceModuleStatusRequest {
	s.SubFeatureKey = &v
	return s
}

func (s *CheckInstanceModuleStatusRequest) Validate() error {
	return dara.Validate(s)
}
