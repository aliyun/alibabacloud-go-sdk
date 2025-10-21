// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledPlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteScheduledPlanHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteScheduledPlanHeaders
	GetWorkspace() *string
}

type DeleteScheduledPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteScheduledPlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledPlanHeaders) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteScheduledPlanHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteScheduledPlanHeaders) SetCommonHeaders(v map[string]*string) *DeleteScheduledPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteScheduledPlanHeaders) SetWorkspace(v string) *DeleteScheduledPlanHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteScheduledPlanHeaders) Validate() error {
	return dara.Validate(s)
}
