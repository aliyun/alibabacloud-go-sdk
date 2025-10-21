// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentTargetV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateDeploymentTargetV2Headers
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *UpdateDeploymentTargetV2Headers
	GetWorkspace() *string
}

type UpdateDeploymentTargetV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateDeploymentTargetV2Headers) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentTargetV2Headers) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentTargetV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateDeploymentTargetV2Headers) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateDeploymentTargetV2Headers) SetCommonHeaders(v map[string]*string) *UpdateDeploymentTargetV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *UpdateDeploymentTargetV2Headers) SetWorkspace(v string) *UpdateDeploymentTargetV2Headers {
	s.Workspace = &v
	return s
}

func (s *UpdateDeploymentTargetV2Headers) Validate() error {
	return dara.Validate(s)
}
