// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionClusterHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateSessionClusterHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *CreateSessionClusterHeaders
	GetWorkspace() *string
}

type CreateSessionClusterHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateSessionClusterHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionClusterHeaders) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateSessionClusterHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateSessionClusterHeaders) SetCommonHeaders(v map[string]*string) *CreateSessionClusterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSessionClusterHeaders) SetWorkspace(v string) *CreateSessionClusterHeaders {
	s.Workspace = &v
	return s
}

func (s *CreateSessionClusterHeaders) Validate() error {
	return dara.Validate(s)
}
