// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLineageInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetLineageInfoHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetLineageInfoHeaders
	GetWorkspace() *string
}

type GetLineageInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 710d6a64d8****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetLineageInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetLineageInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetLineageInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetLineageInfoHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetLineageInfoHeaders) SetCommonHeaders(v map[string]*string) *GetLineageInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLineageInfoHeaders) SetWorkspace(v string) *GetLineageInfoHeaders {
	s.Workspace = &v
	return s
}

func (s *GetLineageInfoHeaders) Validate() error {
	return dara.Validate(s)
}
