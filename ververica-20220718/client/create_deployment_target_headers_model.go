// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentTargetHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateDeploymentTargetHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *CreateDeploymentTargetHeaders
	GetWorkspace() *string
}

type CreateDeploymentTargetHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateDeploymentTargetHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentTargetHeaders) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTargetHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateDeploymentTargetHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateDeploymentTargetHeaders) SetCommonHeaders(v map[string]*string) *CreateDeploymentTargetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeploymentTargetHeaders) SetWorkspace(v string) *CreateDeploymentTargetHeaders {
	s.Workspace = &v
	return s
}

func (s *CreateDeploymentTargetHeaders) Validate() error {
	return dara.Validate(s)
}
