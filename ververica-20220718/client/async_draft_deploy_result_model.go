// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncDraftDeployResult interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactValidationDetail(v *ValidateStatementResult) *AsyncDraftDeployResult
	GetArtifactValidationDetail() *ValidateStatementResult
	SetDeploymentId(v string) *AsyncDraftDeployResult
	GetDeploymentId() *string
	SetMessage(v string) *AsyncDraftDeployResult
	GetMessage() *string
	SetSuccess(v bool) *AsyncDraftDeployResult
	GetSuccess() *bool
	SetTicketStatus(v string) *AsyncDraftDeployResult
	GetTicketStatus() *string
}

type AsyncDraftDeployResult struct {
	ArtifactValidationDetail *ValidateStatementResult `json:"artifactValidationDetail,omitempty" xml:"artifactValidationDetail,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// "Validation error: SQL validate failed"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// FINISHED
	TicketStatus *string `json:"ticketStatus,omitempty" xml:"ticketStatus,omitempty"`
}

func (s AsyncDraftDeployResult) String() string {
	return dara.Prettify(s)
}

func (s AsyncDraftDeployResult) GoString() string {
	return s.String()
}

func (s *AsyncDraftDeployResult) GetArtifactValidationDetail() *ValidateStatementResult {
	return s.ArtifactValidationDetail
}

func (s *AsyncDraftDeployResult) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *AsyncDraftDeployResult) GetMessage() *string {
	return s.Message
}

func (s *AsyncDraftDeployResult) GetSuccess() *bool {
	return s.Success
}

func (s *AsyncDraftDeployResult) GetTicketStatus() *string {
	return s.TicketStatus
}

func (s *AsyncDraftDeployResult) SetArtifactValidationDetail(v *ValidateStatementResult) *AsyncDraftDeployResult {
	s.ArtifactValidationDetail = v
	return s
}

func (s *AsyncDraftDeployResult) SetDeploymentId(v string) *AsyncDraftDeployResult {
	s.DeploymentId = &v
	return s
}

func (s *AsyncDraftDeployResult) SetMessage(v string) *AsyncDraftDeployResult {
	s.Message = &v
	return s
}

func (s *AsyncDraftDeployResult) SetSuccess(v bool) *AsyncDraftDeployResult {
	s.Success = &v
	return s
}

func (s *AsyncDraftDeployResult) SetTicketStatus(v string) *AsyncDraftDeployResult {
	s.TicketStatus = &v
	return s
}

func (s *AsyncDraftDeployResult) Validate() error {
	if s.ArtifactValidationDetail != nil {
		if err := s.ArtifactValidationDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
