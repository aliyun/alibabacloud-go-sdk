// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSessionClusterHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateSessionClusterHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *UpdateSessionClusterHeaders
	GetWorkspace() *string
}

type UpdateSessionClusterHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 710d6a64d8c34d
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateSessionClusterHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateSessionClusterHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSessionClusterHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateSessionClusterHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateSessionClusterHeaders) SetCommonHeaders(v map[string]*string) *UpdateSessionClusterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSessionClusterHeaders) SetWorkspace(v string) *UpdateSessionClusterHeaders {
	s.Workspace = &v
	return s
}

func (s *UpdateSessionClusterHeaders) Validate() error {
	return dara.Validate(s)
}
