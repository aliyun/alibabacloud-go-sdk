// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGenerateResourcePlanResultHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetGenerateResourcePlanResultHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetGenerateResourcePlanResultHeaders
	GetWorkspace() *string
}

type GetGenerateResourcePlanResultHeaders struct {
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

func (s GetGenerateResourcePlanResultHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetGenerateResourcePlanResultHeaders) GoString() string {
	return s.String()
}

func (s *GetGenerateResourcePlanResultHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetGenerateResourcePlanResultHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetGenerateResourcePlanResultHeaders) SetCommonHeaders(v map[string]*string) *GetGenerateResourcePlanResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGenerateResourcePlanResultHeaders) SetWorkspace(v string) *GetGenerateResourcePlanResultHeaders {
	s.Workspace = &v
	return s
}

func (s *GetGenerateResourcePlanResultHeaders) Validate() error {
	return dara.Validate(s)
}
