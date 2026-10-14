// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *GetAgentConfigRequest
	GetAgentId() *string
	SetConfigId(v string) *GetAgentConfigRequest
	GetConfigId() *string
	SetConfigName(v string) *GetAgentConfigRequest
	GetConfigName() *string
	SetCurrent(v string) *GetAgentConfigRequest
	GetCurrent() *string
	SetPageSize(v string) *GetAgentConfigRequest
	GetPageSize() *string
	SetRegion(v string) *GetAgentConfigRequest
	GetRegion() *string
}

type GetAgentConfigRequest struct {
	// The ID of the component for which you want to retrieve configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The ID of the component configuration to retrieve.
	//
	// example:
	//
	// 4d5ecf02-3d01-11f0-8c90-00163e1029af
	ConfigId *string `json:"configId,omitempty" xml:"configId,omitempty"`
	// The name of the configuration.
	//
	// example:
	//
	// Configuration-x
	ConfigName *string `json:"configName,omitempty" xml:"configName,omitempty"`
	// The current page number. This field is available when page-based pagination is used.
	//
	// example:
	//
	// 1
	Current *string `json:"current,omitempty" xml:"current,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s GetAgentConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAgentConfigRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *GetAgentConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *GetAgentConfigRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *GetAgentConfigRequest) GetCurrent() *string {
	return s.Current
}

func (s *GetAgentConfigRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetAgentConfigRequest) GetRegion() *string {
	return s.Region
}

func (s *GetAgentConfigRequest) SetAgentId(v string) *GetAgentConfigRequest {
	s.AgentId = &v
	return s
}

func (s *GetAgentConfigRequest) SetConfigId(v string) *GetAgentConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *GetAgentConfigRequest) SetConfigName(v string) *GetAgentConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *GetAgentConfigRequest) SetCurrent(v string) *GetAgentConfigRequest {
	s.Current = &v
	return s
}

func (s *GetAgentConfigRequest) SetPageSize(v string) *GetAgentConfigRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentConfigRequest) SetRegion(v string) *GetAgentConfigRequest {
	s.Region = &v
	return s
}

func (s *GetAgentConfigRequest) Validate() error {
	return dara.Validate(s)
}
