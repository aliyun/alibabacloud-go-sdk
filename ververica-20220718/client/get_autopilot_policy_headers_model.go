// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutopilotPolicyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAutopilotPolicyHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetAutopilotPolicyHeaders
	GetWorkspace() *string
}

type GetAutopilotPolicyHeaders struct {
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

func (s GetAutopilotPolicyHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAutopilotPolicyHeaders) GoString() string {
	return s.String()
}

func (s *GetAutopilotPolicyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAutopilotPolicyHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetAutopilotPolicyHeaders) SetCommonHeaders(v map[string]*string) *GetAutopilotPolicyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAutopilotPolicyHeaders) SetWorkspace(v string) *GetAutopilotPolicyHeaders {
	s.Workspace = &v
	return s
}

func (s *GetAutopilotPolicyHeaders) Validate() error {
	return dara.Validate(s)
}
