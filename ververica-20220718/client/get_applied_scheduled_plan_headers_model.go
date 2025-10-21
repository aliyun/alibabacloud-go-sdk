// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppliedScheduledPlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAppliedScheduledPlanHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetAppliedScheduledPlanHeaders
	GetWorkspace() *string
}

type GetAppliedScheduledPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetAppliedScheduledPlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAppliedScheduledPlanHeaders) GoString() string {
	return s.String()
}

func (s *GetAppliedScheduledPlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAppliedScheduledPlanHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetAppliedScheduledPlanHeaders) SetCommonHeaders(v map[string]*string) *GetAppliedScheduledPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAppliedScheduledPlanHeaders) SetWorkspace(v string) *GetAppliedScheduledPlanHeaders {
	s.Workspace = &v
	return s
}

func (s *GetAppliedScheduledPlanHeaders) Validate() error {
	return dara.Validate(s)
}
