// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionClusterHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSessionClusterHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetSessionClusterHeaders
	GetWorkspace() *string
}

type GetSessionClusterHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetSessionClusterHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSessionClusterHeaders) GoString() string {
	return s.String()
}

func (s *GetSessionClusterHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSessionClusterHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetSessionClusterHeaders) SetCommonHeaders(v map[string]*string) *GetSessionClusterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSessionClusterHeaders) SetWorkspace(v string) *GetSessionClusterHeaders {
	s.Workspace = &v
	return s
}

func (s *GetSessionClusterHeaders) Validate() error {
	return dara.Validate(s)
}
