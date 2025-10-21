// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUdfArtifactsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUdfArtifactsHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetUdfArtifactsHeaders
	GetWorkspace() *string
}

type GetUdfArtifactsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 710d6a64d8c34d
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetUdfArtifactsHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUdfArtifactsHeaders) GoString() string {
	return s.String()
}

func (s *GetUdfArtifactsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUdfArtifactsHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetUdfArtifactsHeaders) SetCommonHeaders(v map[string]*string) *GetUdfArtifactsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUdfArtifactsHeaders) SetWorkspace(v string) *GetUdfArtifactsHeaders {
	s.Workspace = &v
	return s
}

func (s *GetUdfArtifactsHeaders) Validate() error {
	return dara.Validate(s)
}
