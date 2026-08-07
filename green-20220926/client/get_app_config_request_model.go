// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *GetAppConfigRequest
	GetAgentId() *string
	SetAppId(v string) *GetAppConfigRequest
	GetAppId() *string
	SetAppVersion(v int64) *GetAppConfigRequest
	GetAppVersion() *int64
	SetRegionId(v string) *GetAppConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *GetAppConfigRequest
	GetResourceType() *string
}

type GetAppConfigRequest struct {
	// Agent ID。
	//
	// example:
	//
	// ag.abcxxx
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// App ID。
	//
	// example:
	//
	// txt_check_pro_agent_01
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application version number.
	//
	// example:
	//
	// 1785898163
	AppVersion *int64 `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
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

func (s GetAppConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAppConfigRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *GetAppConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAppConfigRequest) GetAppVersion() *int64 {
	return s.AppVersion
}

func (s *GetAppConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAppConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAppConfigRequest) SetAgentId(v string) *GetAppConfigRequest {
	s.AgentId = &v
	return s
}

func (s *GetAppConfigRequest) SetAppId(v string) *GetAppConfigRequest {
	s.AppId = &v
	return s
}

func (s *GetAppConfigRequest) SetAppVersion(v int64) *GetAppConfigRequest {
	s.AppVersion = &v
	return s
}

func (s *GetAppConfigRequest) SetRegionId(v string) *GetAppConfigRequest {
	s.RegionId = &v
	return s
}

func (s *GetAppConfigRequest) SetResourceType(v string) *GetAppConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *GetAppConfigRequest) Validate() error {
	return dara.Validate(s)
}
