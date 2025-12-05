// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateDeploymentDraftAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DraftValidateParams) *ValidateDeploymentDraftAsyncRequest
	GetBody() *DraftValidateParams
}

type ValidateDeploymentDraftAsyncRequest struct {
	Body *DraftValidateParams `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateDeploymentDraftAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateDeploymentDraftAsyncRequest) GoString() string {
	return s.String()
}

func (s *ValidateDeploymentDraftAsyncRequest) GetBody() *DraftValidateParams {
	return s.Body
}

func (s *ValidateDeploymentDraftAsyncRequest) SetBody(v *DraftValidateParams) *ValidateDeploymentDraftAsyncRequest {
	s.Body = v
	return s
}

func (s *ValidateDeploymentDraftAsyncRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
