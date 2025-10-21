// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopApplyScheduledPlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StopApplyScheduledPlanHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *StopApplyScheduledPlanHeaders
	GetWorkspace() *string
}

type StopApplyScheduledPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s StopApplyScheduledPlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s StopApplyScheduledPlanHeaders) GoString() string {
	return s.String()
}

func (s *StopApplyScheduledPlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StopApplyScheduledPlanHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *StopApplyScheduledPlanHeaders) SetCommonHeaders(v map[string]*string) *StopApplyScheduledPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopApplyScheduledPlanHeaders) SetWorkspace(v string) *StopApplyScheduledPlanHeaders {
	s.Workspace = &v
	return s
}

func (s *StopApplyScheduledPlanHeaders) Validate() error {
	return dara.Validate(s)
}
