// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfigs(v []*UpdateApplicationConfig) *UpdateApplicationConfigsRequest
	GetApplicationConfigs() []*UpdateApplicationConfig
	SetApplicationName(v string) *UpdateApplicationConfigsRequest
	GetApplicationName() *string
	SetClusterId(v string) *UpdateApplicationConfigsRequest
	GetClusterId() *string
	SetConfigAction(v string) *UpdateApplicationConfigsRequest
	GetConfigAction() *string
	SetConfigScope(v string) *UpdateApplicationConfigsRequest
	GetConfigScope() *string
	SetDescription(v string) *UpdateApplicationConfigsRequest
	GetDescription() *string
	SetNodeGroupId(v string) *UpdateApplicationConfigsRequest
	GetNodeGroupId() *string
	SetNodeId(v string) *UpdateApplicationConfigsRequest
	GetNodeId() *string
	SetRefreshConfig(v bool) *UpdateApplicationConfigsRequest
	GetRefreshConfig() *bool
	SetRegionId(v string) *UpdateApplicationConfigsRequest
	GetRegionId() *string
}

type UpdateApplicationConfigsRequest struct {
	// 应用配置列表。
	//
	// This parameter is required.
	ApplicationConfigs []*UpdateApplicationConfig `json:"ApplicationConfigs,omitempty" xml:"ApplicationConfigs,omitempty" type:"Repeated"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// HDFS
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-e6a9d46e9267****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The operation performed on configuration items. Valid values:
	//
	// 	- ADD
	//
	// 	- UPDATE
	//
	// 	- DELETE
	//
	// example:
	//
	// ADD
	ConfigAction *string `json:"ConfigAction,omitempty" xml:"ConfigAction,omitempty"`
	// The operation scope. Valid values:
	//
	// 	- CLUSTER
	//
	// 	- NODE_GROUP
	//
	// example:
	//
	// CLUSTER
	ConfigScope *string `json:"ConfigScope,omitempty" xml:"ConfigScope,omitempty"`
	// The description.
	//
	// example:
	//
	// 更新YARN内存配置。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The node ID.
	//
	// example:
	//
	// i-bp1cudc25w2bfwl5****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Specifies whether to refresh the configurations. Default value: True.
	//
	// example:
	//
	// true
	RefreshConfig *bool `json:"RefreshConfig,omitempty" xml:"RefreshConfig,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateApplicationConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationConfigsRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationConfigsRequest) GetApplicationConfigs() []*UpdateApplicationConfig {
	return s.ApplicationConfigs
}

func (s *UpdateApplicationConfigsRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *UpdateApplicationConfigsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateApplicationConfigsRequest) GetConfigAction() *string {
	return s.ConfigAction
}

func (s *UpdateApplicationConfigsRequest) GetConfigScope() *string {
	return s.ConfigScope
}

func (s *UpdateApplicationConfigsRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationConfigsRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *UpdateApplicationConfigsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateApplicationConfigsRequest) GetRefreshConfig() *bool {
	return s.RefreshConfig
}

func (s *UpdateApplicationConfigsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateApplicationConfigsRequest) SetApplicationConfigs(v []*UpdateApplicationConfig) *UpdateApplicationConfigsRequest {
	s.ApplicationConfigs = v
	return s
}

func (s *UpdateApplicationConfigsRequest) SetApplicationName(v string) *UpdateApplicationConfigsRequest {
	s.ApplicationName = &v
	return s
}

func (s *UpdateApplicationConfigsRequest) SetClusterId(v string) *UpdateApplicationConfigsRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateApplicationConfigsRequest) SetConfigAction(v string) *UpdateApplicationConfigsRequest {
	s.ConfigAction = &v
	return s
}

func (s *UpdateApplicationConfigsRequest) SetConfigScope(v string) *UpdateApplicationConfigsRequest {
	s.ConfigScope = &v
	return s
}

func (s *UpdateApplicationConfigsRequest) SetDescription(v string) *UpdateApplicationConfigsRequest {
	s.Description = &v
	return s
}

func (s *UpdateApplicationConfigsRequest) SetNodeGroupId(v string) *UpdateApplicationConfigsRequest {
	s.NodeGroupId = &v
	return s
}

func (s *UpdateApplicationConfigsRequest) SetNodeId(v string) *UpdateApplicationConfigsRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateApplicationConfigsRequest) SetRefreshConfig(v bool) *UpdateApplicationConfigsRequest {
	s.RefreshConfig = &v
	return s
}

func (s *UpdateApplicationConfigsRequest) SetRegionId(v string) *UpdateApplicationConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateApplicationConfigsRequest) Validate() error {
	return dara.Validate(s)
}
