// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartJobRequestBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v string) *StartJobRequestBody
	GetDeploymentId() *string
	SetResourceSettingSpec(v *BriefResourceSetting) *StartJobRequestBody
	GetResourceSettingSpec() *BriefResourceSetting
	SetRestoreStrategy(v *DeploymentRestoreStrategy) *StartJobRequestBody
	GetRestoreStrategy() *DeploymentRestoreStrategy
}

type StartJobRequestBody struct {
	// example:
	//
	// 5a19a71b-1c42-4f34-94fd-86cf60782c81
	DeploymentId        *string                    `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	ResourceSettingSpec *BriefResourceSetting      `json:"resourceSettingSpec,omitempty" xml:"resourceSettingSpec,omitempty"`
	RestoreStrategy     *DeploymentRestoreStrategy `json:"restoreStrategy,omitempty" xml:"restoreStrategy,omitempty"`
}

func (s StartJobRequestBody) String() string {
	return dara.Prettify(s)
}

func (s StartJobRequestBody) GoString() string {
	return s.String()
}

func (s *StartJobRequestBody) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *StartJobRequestBody) GetResourceSettingSpec() *BriefResourceSetting {
	return s.ResourceSettingSpec
}

func (s *StartJobRequestBody) GetRestoreStrategy() *DeploymentRestoreStrategy {
	return s.RestoreStrategy
}

func (s *StartJobRequestBody) SetDeploymentId(v string) *StartJobRequestBody {
	s.DeploymentId = &v
	return s
}

func (s *StartJobRequestBody) SetResourceSettingSpec(v *BriefResourceSetting) *StartJobRequestBody {
	s.ResourceSettingSpec = v
	return s
}

func (s *StartJobRequestBody) SetRestoreStrategy(v *DeploymentRestoreStrategy) *StartJobRequestBody {
	s.RestoreStrategy = v
	return s
}

func (s *StartJobRequestBody) Validate() error {
	if s.ResourceSettingSpec != nil {
		if err := s.ResourceSettingSpec.Validate(); err != nil {
			return err
		}
	}
	if s.RestoreStrategy != nil {
		if err := s.RestoreStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
