// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNotifyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *NotifyPolicyConfig) *CreateNotifyPolicyRequest
	GetBody() *NotifyPolicyConfig
	SetWorkspace(v string) *CreateNotifyPolicyRequest
	GetWorkspace() *string
}

type CreateNotifyPolicyRequest struct {
	// The request body, which is the complete notification policy configuration object NotifyPolicyConfig.
	Body *NotifyPolicyConfig `json:"body,omitempty" xml:"body,omitempty"`
	// The workspace ID. This parameter is used to isolate notification policy resources across different business spaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// default-cms-xxxx-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateNotifyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNotifyPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateNotifyPolicyRequest) GetBody() *NotifyPolicyConfig {
	return s.Body
}

func (s *CreateNotifyPolicyRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateNotifyPolicyRequest) SetBody(v *NotifyPolicyConfig) *CreateNotifyPolicyRequest {
	s.Body = v
	return s
}

func (s *CreateNotifyPolicyRequest) SetWorkspace(v string) *CreateNotifyPolicyRequest {
	s.Workspace = &v
	return s
}

func (s *CreateNotifyPolicyRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
