// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentDraftLockHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeploymentDraftLockHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetDeploymentDraftLockHeaders
	GetWorkspace() *string
}

type GetDeploymentDraftLockHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetDeploymentDraftLockHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentDraftLockHeaders) GoString() string {
	return s.String()
}

func (s *GetDeploymentDraftLockHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeploymentDraftLockHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetDeploymentDraftLockHeaders) SetCommonHeaders(v map[string]*string) *GetDeploymentDraftLockHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeploymentDraftLockHeaders) SetWorkspace(v string) *GetDeploymentDraftLockHeaders {
	s.Workspace = &v
	return s
}

func (s *GetDeploymentDraftLockHeaders) Validate() error {
	return dara.Validate(s)
}
