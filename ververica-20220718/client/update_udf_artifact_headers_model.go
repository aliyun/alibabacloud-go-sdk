// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUdfArtifactHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateUdfArtifactHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *UpdateUdfArtifactHeaders
	GetWorkspace() *string
}

type UpdateUdfArtifactHeaders struct {
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

func (s UpdateUdfArtifactHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateUdfArtifactHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUdfArtifactHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateUdfArtifactHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateUdfArtifactHeaders) SetCommonHeaders(v map[string]*string) *UpdateUdfArtifactHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUdfArtifactHeaders) SetWorkspace(v string) *UpdateUdfArtifactHeaders {
	s.Workspace = &v
	return s
}

func (s *UpdateUdfArtifactHeaders) Validate() error {
	return dara.Validate(s)
}
