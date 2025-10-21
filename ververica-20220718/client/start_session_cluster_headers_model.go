// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSessionClusterHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StartSessionClusterHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *StartSessionClusterHeaders
	GetWorkspace() *string
}

type StartSessionClusterHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s StartSessionClusterHeaders) String() string {
	return dara.Prettify(s)
}

func (s StartSessionClusterHeaders) GoString() string {
	return s.String()
}

func (s *StartSessionClusterHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StartSessionClusterHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *StartSessionClusterHeaders) SetCommonHeaders(v map[string]*string) *StartSessionClusterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartSessionClusterHeaders) SetWorkspace(v string) *StartSessionClusterHeaders {
	s.Workspace = &v
	return s
}

func (s *StartSessionClusterHeaders) Validate() error {
	return dara.Validate(s)
}
