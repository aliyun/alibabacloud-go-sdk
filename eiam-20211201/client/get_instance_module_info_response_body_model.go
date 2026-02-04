// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceModuleInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v *GetInstanceModuleInfoResponseBodyModule) *GetInstanceModuleInfoResponseBody
	GetModule() *GetInstanceModuleInfoResponseBodyModule
	SetRequestId(v string) *GetInstanceModuleInfoResponseBody
	GetRequestId() *string
}

type GetInstanceModuleInfoResponseBody struct {
	Module *GetInstanceModuleInfoResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceModuleInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceModuleInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceModuleInfoResponseBody) GetModule() *GetInstanceModuleInfoResponseBodyModule {
	return s.Module
}

func (s *GetInstanceModuleInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceModuleInfoResponseBody) SetModule(v *GetInstanceModuleInfoResponseBodyModule) *GetInstanceModuleInfoResponseBody {
	s.Module = v
	return s
}

func (s *GetInstanceModuleInfoResponseBody) SetRequestId(v string) *GetInstanceModuleInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceModuleInfoResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceModuleInfoResponseBodyModule struct {
	// 二级模块信息
	Features []*GetInstanceModuleInfoResponseBodyModuleFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// 模块状态
	//
	// example:
	//
	// urn:alibaba:idaas:license:module:application
	ModuleKey *string `json:"ModuleKey,omitempty" xml:"ModuleKey,omitempty"`
	// 一级模块状态
	//
	// example:
	//
	// enabled
	ModuleStatus *string `json:"ModuleStatus,omitempty" xml:"ModuleStatus,omitempty"`
}

func (s GetInstanceModuleInfoResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceModuleInfoResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetInstanceModuleInfoResponseBodyModule) GetFeatures() []*GetInstanceModuleInfoResponseBodyModuleFeatures {
	return s.Features
}

func (s *GetInstanceModuleInfoResponseBodyModule) GetModuleKey() *string {
	return s.ModuleKey
}

func (s *GetInstanceModuleInfoResponseBodyModule) GetModuleStatus() *string {
	return s.ModuleStatus
}

func (s *GetInstanceModuleInfoResponseBodyModule) SetFeatures(v []*GetInstanceModuleInfoResponseBodyModuleFeatures) *GetInstanceModuleInfoResponseBodyModule {
	s.Features = v
	return s
}

func (s *GetInstanceModuleInfoResponseBodyModule) SetModuleKey(v string) *GetInstanceModuleInfoResponseBodyModule {
	s.ModuleKey = &v
	return s
}

func (s *GetInstanceModuleInfoResponseBodyModule) SetModuleStatus(v string) *GetInstanceModuleInfoResponseBodyModule {
	s.ModuleStatus = &v
	return s
}

func (s *GetInstanceModuleInfoResponseBodyModule) Validate() error {
	if s.Features != nil {
		for _, item := range s.Features {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceModuleInfoResponseBodyModuleFeatures struct {
	// 二级模块标识
	//
	// example:
	//
	// urn:alibaba:idaas:license:module:application:standard:oidc
	FeatureKey *string `json:"FeatureKey,omitempty" xml:"FeatureKey,omitempty"`
	// 二级模块状态
	//
	// example:
	//
	// enabled
	FeatureStatus *string `json:"FeatureStatus,omitempty" xml:"FeatureStatus,omitempty"`
}

func (s GetInstanceModuleInfoResponseBodyModuleFeatures) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceModuleInfoResponseBodyModuleFeatures) GoString() string {
	return s.String()
}

func (s *GetInstanceModuleInfoResponseBodyModuleFeatures) GetFeatureKey() *string {
	return s.FeatureKey
}

func (s *GetInstanceModuleInfoResponseBodyModuleFeatures) GetFeatureStatus() *string {
	return s.FeatureStatus
}

func (s *GetInstanceModuleInfoResponseBodyModuleFeatures) SetFeatureKey(v string) *GetInstanceModuleInfoResponseBodyModuleFeatures {
	s.FeatureKey = &v
	return s
}

func (s *GetInstanceModuleInfoResponseBodyModuleFeatures) SetFeatureStatus(v string) *GetInstanceModuleInfoResponseBodyModuleFeatures {
	s.FeatureStatus = &v
	return s
}

func (s *GetInstanceModuleInfoResponseBodyModuleFeatures) Validate() error {
	return dara.Validate(s)
}
