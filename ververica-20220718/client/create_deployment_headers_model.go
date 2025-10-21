// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateDeploymentHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *CreateDeploymentHeaders
	GetWorkspace() *string
}

type CreateDeploymentHeaders struct {
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

func (s CreateDeploymentHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentHeaders) GoString() string {
	return s.String()
}

func (s *CreateDeploymentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateDeploymentHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateDeploymentHeaders) SetCommonHeaders(v map[string]*string) *CreateDeploymentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeploymentHeaders) SetWorkspace(v string) *CreateDeploymentHeaders {
	s.Workspace = &v
	return s
}

func (s *CreateDeploymentHeaders) Validate() error {
	return dara.Validate(s)
}
