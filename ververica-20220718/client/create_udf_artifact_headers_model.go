// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUdfArtifactHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateUdfArtifactHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *CreateUdfArtifactHeaders
	GetWorkspace() *string
}

type CreateUdfArtifactHeaders struct {
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

func (s CreateUdfArtifactHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateUdfArtifactHeaders) GoString() string {
	return s.String()
}

func (s *CreateUdfArtifactHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateUdfArtifactHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateUdfArtifactHeaders) SetCommonHeaders(v map[string]*string) *CreateUdfArtifactHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateUdfArtifactHeaders) SetWorkspace(v string) *CreateUdfArtifactHeaders {
	s.Workspace = &v
	return s
}

func (s *CreateUdfArtifactHeaders) Validate() error {
	return dara.Validate(s)
}
