// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterConfigsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v int64) *UpdateClusterConfigsShrinkRequest
	GetClusterId() *int64
	SetConfigType(v string) *UpdateClusterConfigsShrinkRequest
	GetConfigType() *string
	SetConfigValuesShrink(v string) *UpdateClusterConfigsShrinkRequest
	GetConfigValuesShrink() *string
	SetProjectId(v int64) *UpdateClusterConfigsShrinkRequest
	GetProjectId() *int64
}

type UpdateClusterConfigsShrinkRequest struct {
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
	ConfigValuesShrink *string `json:"ConfigValues,omitempty" xml:"ConfigValues,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5678
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateClusterConfigsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterConfigsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterConfigsShrinkRequest) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *UpdateClusterConfigsShrinkRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *UpdateClusterConfigsShrinkRequest) GetConfigValuesShrink() *string {
	return s.ConfigValuesShrink
}

func (s *UpdateClusterConfigsShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateClusterConfigsShrinkRequest) SetClusterId(v int64) *UpdateClusterConfigsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterConfigsShrinkRequest) SetConfigType(v string) *UpdateClusterConfigsShrinkRequest {
	s.ConfigType = &v
	return s
}

func (s *UpdateClusterConfigsShrinkRequest) SetConfigValuesShrink(v string) *UpdateClusterConfigsShrinkRequest {
	s.ConfigValuesShrink = &v
	return s
}

func (s *UpdateClusterConfigsShrinkRequest) SetProjectId(v int64) *UpdateClusterConfigsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateClusterConfigsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
