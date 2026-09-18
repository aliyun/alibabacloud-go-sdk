// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAlertRulesInput interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *QueryAlertRulesFilter) *QueryAlertRulesInput
	GetFilter() *QueryAlertRulesFilter
	SetPagination(v *Pagination) *QueryAlertRulesInput
	GetPagination() *Pagination
	SetWorkspace(v string) *QueryAlertRulesInput
	GetWorkspace() *string
}

type QueryAlertRulesInput struct {
	// The comprehensive filter conditions for querying alert rules. Supports combined filtering by multiple dimensions such as name, status, severity level, tags, and notification channels.
	//
	// example:
	//
	// {"status":{"eq":"Alarm"},"severityLevels":{"contains":["CRITICAL","ERROR"]}}
	Filter *QueryAlertRulesFilter `json:"filter,omitempty" xml:"filter,omitempty"`
	// The pagination mode parameter. keyset indicates keyset-based pagination.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize":20}
	Pagination *Pagination `json:"pagination,omitempty" xml:"pagination,omitempty"`
	// The observable workspace name. If this parameter is not specified, all alert rules are returned based on the caller identity (callerUserId + callerRegionId). If this parameter is specified, only the rules under the specified workspace are returned.
	//
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s QueryAlertRulesInput) String() string {
	return dara.Prettify(s)
}

func (s QueryAlertRulesInput) GoString() string {
	return s.String()
}

func (s *QueryAlertRulesInput) GetFilter() *QueryAlertRulesFilter {
	return s.Filter
}

func (s *QueryAlertRulesInput) GetPagination() *Pagination {
	return s.Pagination
}

func (s *QueryAlertRulesInput) GetWorkspace() *string {
	return s.Workspace
}

func (s *QueryAlertRulesInput) SetFilter(v *QueryAlertRulesFilter) *QueryAlertRulesInput {
	s.Filter = v
	return s
}

func (s *QueryAlertRulesInput) SetPagination(v *Pagination) *QueryAlertRulesInput {
	s.Pagination = v
	return s
}

func (s *QueryAlertRulesInput) SetWorkspace(v string) *QueryAlertRulesInput {
	s.Workspace = &v
	return s
}

func (s *QueryAlertRulesInput) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	if s.Pagination != nil {
		if err := s.Pagination.Validate(); err != nil {
			return err
		}
	}
	return nil
}
