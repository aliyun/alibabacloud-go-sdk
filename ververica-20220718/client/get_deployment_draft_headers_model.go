// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentDraftHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeploymentDraftHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetDeploymentDraftHeaders
	GetWorkspace() *string
}

type GetDeploymentDraftHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetDeploymentDraftHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentDraftHeaders) GoString() string {
	return s.String()
}

func (s *GetDeploymentDraftHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeploymentDraftHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetDeploymentDraftHeaders) SetCommonHeaders(v map[string]*string) *GetDeploymentDraftHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeploymentDraftHeaders) SetWorkspace(v string) *GetDeploymentDraftHeaders {
	s.Workspace = &v
	return s
}

func (s *GetDeploymentDraftHeaders) Validate() error {
	return dara.Validate(s)
}
