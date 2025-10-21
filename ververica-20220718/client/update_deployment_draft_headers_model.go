// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentDraftHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateDeploymentDraftHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *UpdateDeploymentDraftHeaders
	GetWorkspace() *string
}

type UpdateDeploymentDraftHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateDeploymentDraftHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentDraftHeaders) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentDraftHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateDeploymentDraftHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateDeploymentDraftHeaders) SetCommonHeaders(v map[string]*string) *UpdateDeploymentDraftHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateDeploymentDraftHeaders) SetWorkspace(v string) *UpdateDeploymentDraftHeaders {
	s.Workspace = &v
	return s
}

func (s *UpdateDeploymentDraftHeaders) Validate() error {
	return dara.Validate(s)
}
