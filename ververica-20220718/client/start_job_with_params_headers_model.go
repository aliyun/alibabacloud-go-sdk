// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartJobWithParamsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StartJobWithParamsHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *StartJobWithParamsHeaders
	GetWorkspace() *string
}

type StartJobWithParamsHeaders struct {
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

func (s StartJobWithParamsHeaders) String() string {
	return dara.Prettify(s)
}

func (s StartJobWithParamsHeaders) GoString() string {
	return s.String()
}

func (s *StartJobWithParamsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StartJobWithParamsHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *StartJobWithParamsHeaders) SetCommonHeaders(v map[string]*string) *StartJobWithParamsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartJobWithParamsHeaders) SetWorkspace(v string) *StartJobWithParamsHeaders {
	s.Workspace = &v
	return s
}

func (s *StartJobWithParamsHeaders) Validate() error {
	return dara.Validate(s)
}
