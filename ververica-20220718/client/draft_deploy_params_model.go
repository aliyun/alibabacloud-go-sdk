// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDraftDeployParams interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentDraftId(v string) *DraftDeployParams
	GetDeploymentDraftId() *string
	SetDeploymentTarget(v *BriefDeploymentTarget) *DraftDeployParams
	GetDeploymentTarget() *BriefDeploymentTarget
	SetSkipValidate(v bool) *DraftDeployParams
	GetSkipValidate() *bool
}

type DraftDeployParams struct {
	DeploymentDraftId *string                `json:"deploymentDraftId,omitempty" xml:"deploymentDraftId,omitempty"`
	DeploymentTarget  *BriefDeploymentTarget `json:"deploymentTarget,omitempty" xml:"deploymentTarget,omitempty"`
	// example:
	//
	// false
	SkipValidate *bool `json:"skipValidate,omitempty" xml:"skipValidate,omitempty"`
}

func (s DraftDeployParams) String() string {
	return dara.Prettify(s)
}

func (s DraftDeployParams) GoString() string {
	return s.String()
}

func (s *DraftDeployParams) GetDeploymentDraftId() *string {
	return s.DeploymentDraftId
}

func (s *DraftDeployParams) GetDeploymentTarget() *BriefDeploymentTarget {
	return s.DeploymentTarget
}

func (s *DraftDeployParams) GetSkipValidate() *bool {
	return s.SkipValidate
}

func (s *DraftDeployParams) SetDeploymentDraftId(v string) *DraftDeployParams {
	s.DeploymentDraftId = &v
	return s
}

func (s *DraftDeployParams) SetDeploymentTarget(v *BriefDeploymentTarget) *DraftDeployParams {
	s.DeploymentTarget = v
	return s
}

func (s *DraftDeployParams) SetSkipValidate(v bool) *DraftDeployParams {
	s.SkipValidate = &v
	return s
}

func (s *DraftDeployParams) Validate() error {
	if s.DeploymentTarget != nil {
		if err := s.DeploymentTarget.Validate(); err != nil {
			return err
		}
	}
	return nil
}
