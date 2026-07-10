// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSqlExecutionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StartSqlExecutionHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *StartSqlExecutionHeaders
	GetWorkspace() *string
}

type StartSqlExecutionHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s StartSqlExecutionHeaders) String() string {
	return dara.Prettify(s)
}

func (s StartSqlExecutionHeaders) GoString() string {
	return s.String()
}

func (s *StartSqlExecutionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StartSqlExecutionHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *StartSqlExecutionHeaders) SetCommonHeaders(v map[string]*string) *StartSqlExecutionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartSqlExecutionHeaders) SetWorkspace(v string) *StartSqlExecutionHeaders {
	s.Workspace = &v
	return s
}

func (s *StartSqlExecutionHeaders) Validate() error {
	return dara.Validate(s)
}
