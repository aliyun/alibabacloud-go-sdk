// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDraftDeployResult interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactValidationDetail(v *ValidateStatementResult) *DraftDeployResult
	GetArtifactValidationDetail() *ValidateStatementResult
	SetDeploymentId(v string) *DraftDeployResult
	GetDeploymentId() *string
	SetMessage(v string) *DraftDeployResult
	GetMessage() *string
	SetSuccess(v bool) *DraftDeployResult
	GetSuccess() *bool
}

type DraftDeployResult struct {
	ArtifactValidationDetail *ValidateStatementResult `json:"artifactValidationDetail,omitempty" xml:"artifactValidationDetail,omitempty"`
	// example:
	//
	// 58718c99-3b29-4c5e-93bb-c9fc4ec6****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DraftDeployResult) String() string {
	return dara.Prettify(s)
}

func (s DraftDeployResult) GoString() string {
	return s.String()
}

func (s *DraftDeployResult) GetArtifactValidationDetail() *ValidateStatementResult {
	return s.ArtifactValidationDetail
}

func (s *DraftDeployResult) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *DraftDeployResult) GetMessage() *string {
	return s.Message
}

func (s *DraftDeployResult) GetSuccess() *bool {
	return s.Success
}

func (s *DraftDeployResult) SetArtifactValidationDetail(v *ValidateStatementResult) *DraftDeployResult {
	s.ArtifactValidationDetail = v
	return s
}

func (s *DraftDeployResult) SetDeploymentId(v string) *DraftDeployResult {
	s.DeploymentId = &v
	return s
}

func (s *DraftDeployResult) SetMessage(v string) *DraftDeployResult {
	s.Message = &v
	return s
}

func (s *DraftDeployResult) SetSuccess(v bool) *DraftDeployResult {
	s.Success = &v
	return s
}

func (s *DraftDeployResult) Validate() error {
	if s.ArtifactValidationDetail != nil {
		if err := s.ArtifactValidationDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
