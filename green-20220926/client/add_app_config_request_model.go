// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAppConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassify(v string) *AddAppConfigRequest
	GetClassify() *string
	SetName(v string) *AddAppConfigRequest
	GetName() *string
	SetRegionId(v string) *AddAppConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *AddAppConfigRequest
	GetResourceType() *string
	SetSysAppId(v string) *AddAppConfigRequest
	GetSysAppId() *string
	SetType(v string) *AddAppConfigRequest
	GetType() *string
}

type AddAppConfigRequest struct {
	// The category.
	//
	// example:
	//
	// guard-scene
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The name.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// appId
	//
	// example:
	//
	// xxx
	SysAppId *string `json:"SysAppId,omitempty" xml:"SysAppId,omitempty"`
	// The type.
	//
	// example:
	//
	// plus
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddAppConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAppConfigRequest) GoString() string {
	return s.String()
}

func (s *AddAppConfigRequest) GetClassify() *string {
	return s.Classify
}

func (s *AddAppConfigRequest) GetName() *string {
	return s.Name
}

func (s *AddAppConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddAppConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *AddAppConfigRequest) GetSysAppId() *string {
	return s.SysAppId
}

func (s *AddAppConfigRequest) GetType() *string {
	return s.Type
}

func (s *AddAppConfigRequest) SetClassify(v string) *AddAppConfigRequest {
	s.Classify = &v
	return s
}

func (s *AddAppConfigRequest) SetName(v string) *AddAppConfigRequest {
	s.Name = &v
	return s
}

func (s *AddAppConfigRequest) SetRegionId(v string) *AddAppConfigRequest {
	s.RegionId = &v
	return s
}

func (s *AddAppConfigRequest) SetResourceType(v string) *AddAppConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *AddAppConfigRequest) SetSysAppId(v string) *AddAppConfigRequest {
	s.SysAppId = &v
	return s
}

func (s *AddAppConfigRequest) SetType(v string) *AddAppConfigRequest {
	s.Type = &v
	return s
}

func (s *AddAppConfigRequest) Validate() error {
	return dara.Validate(s)
}
