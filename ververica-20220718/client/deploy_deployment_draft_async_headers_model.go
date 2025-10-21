// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployDeploymentDraftAsyncHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeployDeploymentDraftAsyncHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeployDeploymentDraftAsyncHeaders
	GetWorkspace() *string
}

type DeployDeploymentDraftAsyncHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeployDeploymentDraftAsyncHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeployDeploymentDraftAsyncHeaders) GoString() string {
	return s.String()
}

func (s *DeployDeploymentDraftAsyncHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeployDeploymentDraftAsyncHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeployDeploymentDraftAsyncHeaders) SetCommonHeaders(v map[string]*string) *DeployDeploymentDraftAsyncHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeployDeploymentDraftAsyncHeaders) SetWorkspace(v string) *DeployDeploymentDraftAsyncHeaders {
	s.Workspace = &v
	return s
}

func (s *DeployDeploymentDraftAsyncHeaders) Validate() error {
	return dara.Validate(s)
}
