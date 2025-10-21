// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionClustersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListSessionClustersHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ListSessionClustersHeaders
	GetWorkspace() *string
}

type ListSessionClustersHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListSessionClustersHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListSessionClustersHeaders) GoString() string {
	return s.String()
}

func (s *ListSessionClustersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListSessionClustersHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListSessionClustersHeaders) SetCommonHeaders(v map[string]*string) *ListSessionClustersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSessionClustersHeaders) SetWorkspace(v string) *ListSessionClustersHeaders {
	s.Workspace = &v
	return s
}

func (s *ListSessionClustersHeaders) Validate() error {
	return dara.Validate(s)
}
