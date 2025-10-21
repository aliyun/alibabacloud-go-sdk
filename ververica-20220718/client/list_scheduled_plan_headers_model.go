// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledPlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListScheduledPlanHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ListScheduledPlanHeaders
	GetWorkspace() *string
}

type ListScheduledPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListScheduledPlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPlanHeaders) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListScheduledPlanHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListScheduledPlanHeaders) SetCommonHeaders(v map[string]*string) *ListScheduledPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListScheduledPlanHeaders) SetWorkspace(v string) *ListScheduledPlanHeaders {
	s.Workspace = &v
	return s
}

func (s *ListScheduledPlanHeaders) Validate() error {
	return dara.Validate(s)
}
