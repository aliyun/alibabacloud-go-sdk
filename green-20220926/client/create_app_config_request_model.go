// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateAppConfigRequest
	GetName() *string
	SetRegionId(v string) *CreateAppConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *CreateAppConfigRequest
	GetResourceType() *string
	SetSysAppId(v string) *CreateAppConfigRequest
	GetSysAppId() *string
	SetType(v string) *CreateAppConfigRequest
	GetType() *string
}

type CreateAppConfigRequest struct {
	// The name.
	//
	// example:
	//
	// CustomReview
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
	// agent_text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The system app ID.
	//
	// example:
	//
	// txt_check_agent
	SysAppId *string `json:"SysAppId,omitempty" xml:"SysAppId,omitempty"`
	// The type.
	//
	// example:
	//
	// plus
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateAppConfigRequest) GetName() *string {
	return s.Name
}

func (s *CreateAppConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAppConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateAppConfigRequest) GetSysAppId() *string {
	return s.SysAppId
}

func (s *CreateAppConfigRequest) GetType() *string {
	return s.Type
}

func (s *CreateAppConfigRequest) SetName(v string) *CreateAppConfigRequest {
	s.Name = &v
	return s
}

func (s *CreateAppConfigRequest) SetRegionId(v string) *CreateAppConfigRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAppConfigRequest) SetResourceType(v string) *CreateAppConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateAppConfigRequest) SetSysAppId(v string) *CreateAppConfigRequest {
	s.SysAppId = &v
	return s
}

func (s *CreateAppConfigRequest) SetType(v string) *CreateAppConfigRequest {
	s.Type = &v
	return s
}

func (s *CreateAppConfigRequest) Validate() error {
	return dara.Validate(s)
}
