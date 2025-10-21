// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentDraftHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateDeploymentDraftHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *CreateDeploymentDraftHeaders
	GetWorkspace() *string
}

type CreateDeploymentDraftHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateDeploymentDraftHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentDraftHeaders) GoString() string {
	return s.String()
}

func (s *CreateDeploymentDraftHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateDeploymentDraftHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateDeploymentDraftHeaders) SetCommonHeaders(v map[string]*string) *CreateDeploymentDraftHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeploymentDraftHeaders) SetWorkspace(v string) *CreateDeploymentDraftHeaders {
	s.Workspace = &v
	return s
}

func (s *CreateDeploymentDraftHeaders) Validate() error {
	return dara.Validate(s)
}
