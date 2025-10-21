// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteDeploymentHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteDeploymentHeaders
	GetWorkspace() *string
}

type DeleteDeploymentHeaders struct {
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

func (s DeleteDeploymentHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteDeploymentHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteDeploymentHeaders) SetCommonHeaders(v map[string]*string) *DeleteDeploymentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDeploymentHeaders) SetWorkspace(v string) *DeleteDeploymentHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteDeploymentHeaders) Validate() error {
	return dara.Validate(s)
}
