// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyScheduledPlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ApplyScheduledPlanHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ApplyScheduledPlanHeaders
	GetWorkspace() *string
}

type ApplyScheduledPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ApplyScheduledPlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s ApplyScheduledPlanHeaders) GoString() string {
	return s.String()
}

func (s *ApplyScheduledPlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ApplyScheduledPlanHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ApplyScheduledPlanHeaders) SetCommonHeaders(v map[string]*string) *ApplyScheduledPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyScheduledPlanHeaders) SetWorkspace(v string) *ApplyScheduledPlanHeaders {
	s.Workspace = &v
	return s
}

func (s *ApplyScheduledPlanHeaders) Validate() error {
	return dara.Validate(s)
}
