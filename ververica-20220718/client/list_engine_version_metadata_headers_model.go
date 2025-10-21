// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEngineVersionMetadataHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListEngineVersionMetadataHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ListEngineVersionMetadataHeaders
	GetWorkspace() *string
}

type ListEngineVersionMetadataHeaders struct {
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

func (s ListEngineVersionMetadataHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListEngineVersionMetadataHeaders) GoString() string {
	return s.String()
}

func (s *ListEngineVersionMetadataHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListEngineVersionMetadataHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListEngineVersionMetadataHeaders) SetCommonHeaders(v map[string]*string) *ListEngineVersionMetadataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListEngineVersionMetadataHeaders) SetWorkspace(v string) *ListEngineVersionMetadataHeaders {
	s.Workspace = &v
	return s
}

func (s *ListEngineVersionMetadataHeaders) Validate() error {
	return dara.Validate(s)
}
