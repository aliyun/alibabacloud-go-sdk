// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSessionClusterHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StopSessionClusterHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *StopSessionClusterHeaders
	GetWorkspace() *string
}

type StopSessionClusterHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s StopSessionClusterHeaders) String() string {
	return dara.Prettify(s)
}

func (s StopSessionClusterHeaders) GoString() string {
	return s.String()
}

func (s *StopSessionClusterHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StopSessionClusterHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *StopSessionClusterHeaders) SetCommonHeaders(v map[string]*string) *StopSessionClusterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopSessionClusterHeaders) SetWorkspace(v string) *StopSessionClusterHeaders {
	s.Workspace = &v
	return s
}

func (s *StopSessionClusterHeaders) Validate() error {
	return dara.Validate(s)
}
