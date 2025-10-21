// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledPlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateScheduledPlanHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *CreateScheduledPlanHeaders
	GetWorkspace() *string
}

type CreateScheduledPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateScheduledPlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledPlanHeaders) GoString() string {
	return s.String()
}

func (s *CreateScheduledPlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateScheduledPlanHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateScheduledPlanHeaders) SetCommonHeaders(v map[string]*string) *CreateScheduledPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateScheduledPlanHeaders) SetWorkspace(v string) *CreateScheduledPlanHeaders {
	s.Workspace = &v
	return s
}

func (s *CreateScheduledPlanHeaders) Validate() error {
	return dara.Validate(s)
}
