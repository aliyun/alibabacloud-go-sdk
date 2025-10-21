// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListDeploymentsHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ListDeploymentsHeaders
	GetWorkspace() *string
}

type ListDeploymentsHeaders struct {
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

func (s ListDeploymentsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentsHeaders) GoString() string {
	return s.String()
}

func (s *ListDeploymentsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListDeploymentsHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListDeploymentsHeaders) SetCommonHeaders(v map[string]*string) *ListDeploymentsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeploymentsHeaders) SetWorkspace(v string) *ListDeploymentsHeaders {
	s.Workspace = &v
	return s
}

func (s *ListDeploymentsHeaders) Validate() error {
	return dara.Validate(s)
}
