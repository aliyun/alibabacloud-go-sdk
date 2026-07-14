// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNotifyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *NotifyPolicyConfig) *UpdateNotifyPolicyRequest
	GetBody() *NotifyPolicyConfig
	SetWorkspace(v string) *UpdateNotifyPolicyRequest
	GetWorkspace() *string
}

type UpdateNotifyPolicyRequest struct {
	// The request body. This is the complete notification policy configuration object NotifyPolicyConfig.
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

func (s UpdateNotifyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNotifyPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateNotifyPolicyRequest) GetBody() *NotifyPolicyConfig {
	return s.Body
}

func (s *UpdateNotifyPolicyRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateNotifyPolicyRequest) SetBody(v *NotifyPolicyConfig) *UpdateNotifyPolicyRequest {
	s.Body = v
	return s
}

func (s *UpdateNotifyPolicyRequest) SetWorkspace(v string) *UpdateNotifyPolicyRequest {
	s.Workspace = &v
	return s
}

func (s *UpdateNotifyPolicyRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
