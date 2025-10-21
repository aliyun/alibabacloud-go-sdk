// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSessionClusterHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteSessionClusterHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteSessionClusterHeaders
	GetWorkspace() *string
}

type DeleteSessionClusterHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteSessionClusterHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteSessionClusterHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSessionClusterHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteSessionClusterHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteSessionClusterHeaders) SetCommonHeaders(v map[string]*string) *DeleteSessionClusterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSessionClusterHeaders) SetWorkspace(v string) *DeleteSessionClusterHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteSessionClusterHeaders) Validate() error {
	return dara.Validate(s)
}
