// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppInfoRequest
	GetAppId() *string
	SetName(v string) *ModifyAppInfoRequest
	GetName() *string
	SetRegionId(v string) *ModifyAppInfoRequest
	GetRegionId() *string
	SetResourceType(v string) *ModifyAppInfoRequest
	GetResourceType() *string
}

type ModifyAppInfoRequest struct {
	// App ID。
	//
	// example:
	//
	// txt_check_agent_01
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The app name.
	//
	// example:
	//
	// Custom text moderation
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
}

func (s ModifyAppInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppInfoRequest) GetName() *string {
	return s.Name
}

func (s *ModifyAppInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAppInfoRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ModifyAppInfoRequest) SetAppId(v string) *ModifyAppInfoRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppInfoRequest) SetName(v string) *ModifyAppInfoRequest {
	s.Name = &v
	return s
}

func (s *ModifyAppInfoRequest) SetRegionId(v string) *ModifyAppInfoRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAppInfoRequest) SetResourceType(v string) *ModifyAppInfoRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyAppInfoRequest) Validate() error {
	return dara.Validate(s)
}
