// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdfArtifactHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteUdfArtifactHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteUdfArtifactHeaders
	GetWorkspace() *string
}

type DeleteUdfArtifactHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteUdfArtifactHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdfArtifactHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUdfArtifactHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteUdfArtifactHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteUdfArtifactHeaders) SetCommonHeaders(v map[string]*string) *DeleteUdfArtifactHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUdfArtifactHeaders) SetWorkspace(v string) *DeleteUdfArtifactHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteUdfArtifactHeaders) Validate() error {
	return dara.Validate(s)
}
