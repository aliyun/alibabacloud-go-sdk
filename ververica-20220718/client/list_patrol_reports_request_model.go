// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPatrolReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v int64) *ListPatrolReportsRequest
	GetEndDate() *int64
	SetPage(v int32) *ListPatrolReportsRequest
	GetPage() *int32
	SetScopeType(v string) *ListPatrolReportsRequest
	GetScopeType() *string
	SetSize(v int32) *ListPatrolReportsRequest
	GetSize() *int32
	SetStartDate(v int64) *ListPatrolReportsRequest
	GetStartDate() *int64
	SetStatus(v string) *ListPatrolReportsRequest
	GetStatus() *string
	SetTriggerType(v string) *ListPatrolReportsRequest
	GetTriggerType() *string
}

type ListPatrolReportsRequest struct {
	// The end time of the query. Unit: milliseconds (UNIX timestamp).
	//
	// example:
	//
	// 1718086400000
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// Filters reports by inspection scope type. Valid values:
	//
	// - ALL
	//
	// - TAGS
	//
	// - DEPLOYMENTS
	//
	// example:
	//
	// ALL
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 200.
	//
	// example:
	//
	// 20
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The start time of the query. Unit: milliseconds (UNIX timestamp).
	//
	// example:
	//
	// 1718000000000
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// Filters reports by status. Valid values:
	//
	// - PENDING
	//
	// - IN_PROGRESS
	//
	// - COMPLETED
	//
	// - FAILED
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Filters reports by trigger type. Valid values:
	//
	// - CRON
	//
	// - MANUAL
	//
	// - INNER_API
	//
	// example:
	//
	// CRON
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s ListPatrolReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPatrolReportsRequest) GoString() string {
	return s.String()
}

func (s *ListPatrolReportsRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *ListPatrolReportsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListPatrolReportsRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *ListPatrolReportsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListPatrolReportsRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *ListPatrolReportsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListPatrolReportsRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListPatrolReportsRequest) SetEndDate(v int64) *ListPatrolReportsRequest {
	s.EndDate = &v
	return s
}

func (s *ListPatrolReportsRequest) SetPage(v int32) *ListPatrolReportsRequest {
	s.Page = &v
	return s
}

func (s *ListPatrolReportsRequest) SetScopeType(v string) *ListPatrolReportsRequest {
	s.ScopeType = &v
	return s
}

func (s *ListPatrolReportsRequest) SetSize(v int32) *ListPatrolReportsRequest {
	s.Size = &v
	return s
}

func (s *ListPatrolReportsRequest) SetStartDate(v int64) *ListPatrolReportsRequest {
	s.StartDate = &v
	return s
}

func (s *ListPatrolReportsRequest) SetStatus(v string) *ListPatrolReportsRequest {
	s.Status = &v
	return s
}

func (s *ListPatrolReportsRequest) SetTriggerType(v string) *ListPatrolReportsRequest {
	s.TriggerType = &v
	return s
}

func (s *ListPatrolReportsRequest) Validate() error {
	return dara.Validate(s)
}
