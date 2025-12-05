// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateDeploymentDraftAsyncHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ValidateDeploymentDraftAsyncHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ValidateDeploymentDraftAsyncHeaders
	GetWorkspace() *string
}

type ValidateDeploymentDraftAsyncHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ValidateDeploymentDraftAsyncHeaders) String() string {
	return dara.Prettify(s)
}

func (s ValidateDeploymentDraftAsyncHeaders) GoString() string {
	return s.String()
}

func (s *ValidateDeploymentDraftAsyncHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ValidateDeploymentDraftAsyncHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ValidateDeploymentDraftAsyncHeaders) SetCommonHeaders(v map[string]*string) *ValidateDeploymentDraftAsyncHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ValidateDeploymentDraftAsyncHeaders) SetWorkspace(v string) *ValidateDeploymentDraftAsyncHeaders {
	s.Workspace = &v
	return s
}

func (s *ValidateDeploymentDraftAsyncHeaders) Validate() error {
	return dara.Validate(s)
}
