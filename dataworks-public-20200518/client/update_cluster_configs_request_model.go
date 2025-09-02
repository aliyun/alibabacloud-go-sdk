// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v int64) *UpdateClusterConfigsRequest
	GetClusterId() *int64
	SetConfigType(v string) *UpdateClusterConfigsRequest
	GetConfigType() *string
	SetConfigValues(v []*ClusterConfig) *UpdateClusterConfigsRequest
	GetConfigValues() []*ClusterConfig
	SetProjectId(v int64) *UpdateClusterConfigsRequest
	GetProjectId() *int64
}

type UpdateClusterConfigsRequest struct {
	// The ID of the cluster associated with DataWorks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ClusterId *int64 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The configuration type of the cluster. Valid values:
	//
	// 	- SPARK_CONF: SPARK parameters
	//
	// This parameter is required.
	//
	// example:
	//
	// SPARK_CONF
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The configuration information of the cluster submodule.
	//
	// This parameter is required.
	ConfigValues []*ClusterConfig `json:"ConfigValues,omitempty" xml:"ConfigValues,omitempty" type:"Repeated"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5678
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateClusterConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterConfigsRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterConfigsRequest) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *UpdateClusterConfigsRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *UpdateClusterConfigsRequest) GetConfigValues() []*ClusterConfig {
	return s.ConfigValues
}

func (s *UpdateClusterConfigsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateClusterConfigsRequest) SetClusterId(v int64) *UpdateClusterConfigsRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterConfigsRequest) SetConfigType(v string) *UpdateClusterConfigsRequest {
	s.ConfigType = &v
	return s
}

func (s *UpdateClusterConfigsRequest) SetConfigValues(v []*ClusterConfig) *UpdateClusterConfigsRequest {
	s.ConfigValues = v
	return s
}

func (s *UpdateClusterConfigsRequest) SetProjectId(v int64) *UpdateClusterConfigsRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateClusterConfigsRequest) Validate() error {
	return dara.Validate(s)
}
