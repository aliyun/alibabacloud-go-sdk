// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledPlanExecutedHistoryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListScheduledPlanExecutedHistoryHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ListScheduledPlanExecutedHistoryHeaders
	GetWorkspace() *string
}

type ListScheduledPlanExecutedHistoryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListScheduledPlanExecutedHistoryHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPlanExecutedHistoryHeaders) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanExecutedHistoryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListScheduledPlanExecutedHistoryHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListScheduledPlanExecutedHistoryHeaders) SetCommonHeaders(v map[string]*string) *ListScheduledPlanExecutedHistoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListScheduledPlanExecutedHistoryHeaders) SetWorkspace(v string) *ListScheduledPlanExecutedHistoryHeaders {
	s.Workspace = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryHeaders) Validate() error {
	return dara.Validate(s)
}
