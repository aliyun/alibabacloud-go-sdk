// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentTargetHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteDeploymentTargetHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteDeploymentTargetHeaders
	GetWorkspace() *string
}

type DeleteDeploymentTargetHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteDeploymentTargetHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentTargetHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentTargetHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteDeploymentTargetHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteDeploymentTargetHeaders) SetCommonHeaders(v map[string]*string) *DeleteDeploymentTargetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDeploymentTargetHeaders) SetWorkspace(v string) *DeleteDeploymentTargetHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteDeploymentTargetHeaders) Validate() error {
	return dara.Validate(s)
}
