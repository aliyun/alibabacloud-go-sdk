// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentDraftLockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentDraftId(v string) *GetDeploymentDraftLockRequest
	GetDeploymentDraftId() *string
}

type GetDeploymentDraftLockRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c84d73be-40ad-4627-8bdd-fa1eba51b234
	DeploymentDraftId *string `json:"deploymentDraftId,omitempty" xml:"deploymentDraftId,omitempty"`
}

func (s GetDeploymentDraftLockRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentDraftLockRequest) GoString() string {
	return s.String()
}

func (s *GetDeploymentDraftLockRequest) GetDeploymentDraftId() *string {
	return s.DeploymentDraftId
}

func (s *GetDeploymentDraftLockRequest) SetDeploymentDraftId(v string) *GetDeploymentDraftLockRequest {
	s.DeploymentDraftId = &v
	return s
}

func (s *GetDeploymentDraftLockRequest) Validate() error {
	return dara.Validate(s)
}
