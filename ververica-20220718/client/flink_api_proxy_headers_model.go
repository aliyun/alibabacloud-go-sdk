// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlinkApiProxyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlinkApiProxyHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *FlinkApiProxyHeaders
	GetWorkspace() *string
}

type FlinkApiProxyHeaders struct {
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

func (s FlinkApiProxyHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlinkApiProxyHeaders) GoString() string {
	return s.String()
}

func (s *FlinkApiProxyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlinkApiProxyHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *FlinkApiProxyHeaders) SetCommonHeaders(v map[string]*string) *FlinkApiProxyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlinkApiProxyHeaders) SetWorkspace(v string) *FlinkApiProxyHeaders {
	s.Workspace = &v
	return s
}

func (s *FlinkApiProxyHeaders) Validate() error {
	return dara.Validate(s)
}
