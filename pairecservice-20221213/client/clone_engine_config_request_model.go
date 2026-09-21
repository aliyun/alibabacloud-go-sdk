// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneEngineConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigValue(v string) *CloneEngineConfigRequest
	GetConfigValue() *string
	SetDescription(v string) *CloneEngineConfigRequest
	GetDescription() *string
	SetEnvironment(v string) *CloneEngineConfigRequest
	GetEnvironment() *string
	SetInstanceId(v string) *CloneEngineConfigRequest
	GetInstanceId() *string
	SetSceneId(v string) *CloneEngineConfigRequest
	GetSceneId() *string
}

type CloneEngineConfigRequest struct {
	// The content of the DPI engine configuration.
	//
	// example:
	//
	// {}
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The description.
	//
	// example:
	//
	// this is a clone config.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The runtime environment. Valid values:
	//
	// - Daily: daily environment.
	//
	// - Pre: staging environment.
	//
	// - Prod: production environment.
	//
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The instance ID. For information about how to obtain the instance ID, see [ListInstances](https://help.aliyun.com/document_detail/2411819.html).
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The scene.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s CloneEngineConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneEngineConfigRequest) GoString() string {
	return s.String()
}

func (s *CloneEngineConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *CloneEngineConfigRequest) GetDescription() *string {
	return s.Description
}

func (s *CloneEngineConfigRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *CloneEngineConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CloneEngineConfigRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *CloneEngineConfigRequest) SetConfigValue(v string) *CloneEngineConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *CloneEngineConfigRequest) SetDescription(v string) *CloneEngineConfigRequest {
	s.Description = &v
	return s
}

func (s *CloneEngineConfigRequest) SetEnvironment(v string) *CloneEngineConfigRequest {
	s.Environment = &v
	return s
}

func (s *CloneEngineConfigRequest) SetInstanceId(v string) *CloneEngineConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CloneEngineConfigRequest) SetSceneId(v string) *CloneEngineConfigRequest {
	s.SceneId = &v
	return s
}

func (s *CloneEngineConfigRequest) Validate() error {
	return dara.Validate(s)
}
