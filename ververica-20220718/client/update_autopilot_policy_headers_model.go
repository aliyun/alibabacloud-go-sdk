// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutopilotPolicyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateAutopilotPolicyHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *UpdateAutopilotPolicyHeaders
	GetWorkspace() *string
}

type UpdateAutopilotPolicyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateAutopilotPolicyHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutopilotPolicyHeaders) GoString() string {
	return s.String()
}

func (s *UpdateAutopilotPolicyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateAutopilotPolicyHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateAutopilotPolicyHeaders) SetCommonHeaders(v map[string]*string) *UpdateAutopilotPolicyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateAutopilotPolicyHeaders) SetWorkspace(v string) *UpdateAutopilotPolicyHeaders {
	s.Workspace = &v
	return s
}

func (s *UpdateAutopilotPolicyHeaders) Validate() error {
	return dara.Validate(s)
}
