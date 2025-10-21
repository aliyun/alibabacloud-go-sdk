// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetEventsHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetEventsHeaders
	GetWorkspace() *string
}

type GetEventsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetEventsHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetEventsHeaders) GoString() string {
	return s.String()
}

func (s *GetEventsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetEventsHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetEventsHeaders) SetCommonHeaders(v map[string]*string) *GetEventsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEventsHeaders) SetWorkspace(v string) *GetEventsHeaders {
	s.Workspace = &v
	return s
}

func (s *GetEventsHeaders) Validate() error {
	return dara.Validate(s)
}
