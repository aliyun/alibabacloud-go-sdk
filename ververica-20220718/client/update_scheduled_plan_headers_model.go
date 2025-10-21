// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledPlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateScheduledPlanHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *UpdateScheduledPlanHeaders
	GetWorkspace() *string
}

type UpdateScheduledPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateScheduledPlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledPlanHeaders) GoString() string {
	return s.String()
}

func (s *UpdateScheduledPlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateScheduledPlanHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateScheduledPlanHeaders) SetCommonHeaders(v map[string]*string) *UpdateScheduledPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateScheduledPlanHeaders) SetWorkspace(v string) *UpdateScheduledPlanHeaders {
	s.Workspace = &v
	return s
}

func (s *UpdateScheduledPlanHeaders) Validate() error {
	return dara.Validate(s)
}
