// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentDraftHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteDeploymentDraftHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteDeploymentDraftHeaders
	GetWorkspace() *string
}

type DeleteDeploymentDraftHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteDeploymentDraftHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentDraftHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentDraftHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteDeploymentDraftHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteDeploymentDraftHeaders) SetCommonHeaders(v map[string]*string) *DeleteDeploymentDraftHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDeploymentDraftHeaders) SetWorkspace(v string) *DeleteDeploymentDraftHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteDeploymentDraftHeaders) Validate() error {
	return dara.Validate(s)
}
