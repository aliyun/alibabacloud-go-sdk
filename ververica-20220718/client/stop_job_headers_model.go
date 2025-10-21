// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopJobHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StopJobHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *StopJobHeaders
	GetWorkspace() *string
}

type StopJobHeaders struct {
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

func (s StopJobHeaders) String() string {
	return dara.Prettify(s)
}

func (s StopJobHeaders) GoString() string {
	return s.String()
}

func (s *StopJobHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StopJobHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *StopJobHeaders) SetCommonHeaders(v map[string]*string) *StopJobHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopJobHeaders) SetWorkspace(v string) *StopJobHeaders {
	s.Workspace = &v
	return s
}

func (s *StopJobHeaders) Validate() error {
	return dara.Validate(s)
}
