// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLatestJobStartLogHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetLatestJobStartLogHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetLatestJobStartLogHeaders
	GetWorkspace() *string
}

type GetLatestJobStartLogHeaders struct {
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

func (s GetLatestJobStartLogHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetLatestJobStartLogHeaders) GoString() string {
	return s.String()
}

func (s *GetLatestJobStartLogHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetLatestJobStartLogHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetLatestJobStartLogHeaders) SetCommonHeaders(v map[string]*string) *GetLatestJobStartLogHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLatestJobStartLogHeaders) SetWorkspace(v string) *GetLatestJobStartLogHeaders {
	s.Workspace = &v
	return s
}

func (s *GetLatestJobStartLogHeaders) Validate() error {
	return dara.Validate(s)
}
