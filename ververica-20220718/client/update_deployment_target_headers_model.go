// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentTargetHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateDeploymentTargetHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *UpdateDeploymentTargetHeaders
	GetWorkspace() *string
}

type UpdateDeploymentTargetHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateDeploymentTargetHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentTargetHeaders) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentTargetHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateDeploymentTargetHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateDeploymentTargetHeaders) SetCommonHeaders(v map[string]*string) *UpdateDeploymentTargetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateDeploymentTargetHeaders) SetWorkspace(v string) *UpdateDeploymentTargetHeaders {
	s.Workspace = &v
	return s
}

func (s *UpdateDeploymentTargetHeaders) Validate() error {
	return dara.Validate(s)
}
