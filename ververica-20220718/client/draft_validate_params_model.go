// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDraftValidateParams interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentDraftId(v string) *DraftValidateParams
	GetDeploymentDraftId() *string
	SetDeploymentTargetName(v string) *DraftValidateParams
	GetDeploymentTargetName() *string
}

type DraftValidateParams struct {
	DeploymentDraftId    *string `json:"deploymentDraftId,omitempty" xml:"deploymentDraftId,omitempty"`
	DeploymentTargetName *string `json:"deploymentTargetName,omitempty" xml:"deploymentTargetName,omitempty"`
}

func (s DraftValidateParams) String() string {
	return dara.Prettify(s)
}

func (s DraftValidateParams) GoString() string {
	return s.String()
}

func (s *DraftValidateParams) GetDeploymentDraftId() *string {
	return s.DeploymentDraftId
}

func (s *DraftValidateParams) GetDeploymentTargetName() *string {
	return s.DeploymentTargetName
}

func (s *DraftValidateParams) SetDeploymentDraftId(v string) *DraftValidateParams {
	s.DeploymentDraftId = &v
	return s
}

func (s *DraftValidateParams) SetDeploymentTargetName(v string) *DraftValidateParams {
	s.DeploymentTargetName = &v
	return s
}

func (s *DraftValidateParams) Validate() error {
	return dara.Validate(s)
}
