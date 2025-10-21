// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdfFunctionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteUdfFunctionHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteUdfFunctionHeaders
	GetWorkspace() *string
}

type DeleteUdfFunctionHeaders struct {
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

func (s DeleteUdfFunctionHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdfFunctionHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUdfFunctionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteUdfFunctionHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteUdfFunctionHeaders) SetCommonHeaders(v map[string]*string) *DeleteUdfFunctionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUdfFunctionHeaders) SetWorkspace(v string) *DeleteUdfFunctionHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteUdfFunctionHeaders) Validate() error {
	return dara.Validate(s)
}
