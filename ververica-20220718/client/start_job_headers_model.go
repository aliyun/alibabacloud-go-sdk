// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartJobHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StartJobHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *StartJobHeaders
	GetWorkspace() *string
}

type StartJobHeaders struct {
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

func (s StartJobHeaders) String() string {
	return dara.Prettify(s)
}

func (s StartJobHeaders) GoString() string {
	return s.String()
}

func (s *StartJobHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StartJobHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *StartJobHeaders) SetCommonHeaders(v map[string]*string) *StartJobHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartJobHeaders) SetWorkspace(v string) *StartJobHeaders {
	s.Workspace = &v
	return s
}

func (s *StartJobHeaders) Validate() error {
	return dara.Validate(s)
}
