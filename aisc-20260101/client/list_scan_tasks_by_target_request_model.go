// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScanTasksByTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *ListScanTasksByTargetRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListScanTasksByTargetRequest
	GetPageSize() *string
	SetSampleLevel(v string) *ListScanTasksByTargetRequest
	GetSampleLevel() *string
	SetScanType(v string) *ListScanTasksByTargetRequest
	GetScanType() *string
	SetTargetId(v string) *ListScanTasksByTargetRequest
	GetTargetId() *string
	SetTaskStatus(v string) *ListScanTasksByTargetRequest
	GetTaskStatus() *string
}

type ListScanTasksByTargetRequest struct {
	// The page number, starting from 1. Values less than 1 are normalized to 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100. Values greater than 100 are clamped to 100. Values less than 1 return HTTP status code 400.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters by detection intensity. If this parameter is not specified, no filtering by intensity is applied.
	//
	// example:
	//
	// 2
	SampleLevel *string `json:"SampleLevel,omitempty" xml:"SampleLevel,omitempty"`
	// Filters by scan type. If this parameter is not specified, tasks of all scan types are returned.
	//
	// example:
	//
	// attack
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The unique identifier of the scan target. Only tasks under this target are queried. If the target does not exist or does not belong to the current tenant, HTTP status code 400 is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// target-abc123def4567
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// Filters by task status. If this parameter is not specified, tasks in all statuses are returned.
	//
	// example:
	//
	// completed
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s ListScanTasksByTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScanTasksByTargetRequest) GoString() string {
	return s.String()
}

func (s *ListScanTasksByTargetRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListScanTasksByTargetRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListScanTasksByTargetRequest) GetSampleLevel() *string {
	return s.SampleLevel
}

func (s *ListScanTasksByTargetRequest) GetScanType() *string {
	return s.ScanType
}

func (s *ListScanTasksByTargetRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *ListScanTasksByTargetRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *ListScanTasksByTargetRequest) SetPageNumber(v string) *ListScanTasksByTargetRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScanTasksByTargetRequest) SetPageSize(v string) *ListScanTasksByTargetRequest {
	s.PageSize = &v
	return s
}

func (s *ListScanTasksByTargetRequest) SetSampleLevel(v string) *ListScanTasksByTargetRequest {
	s.SampleLevel = &v
	return s
}

func (s *ListScanTasksByTargetRequest) SetScanType(v string) *ListScanTasksByTargetRequest {
	s.ScanType = &v
	return s
}

func (s *ListScanTasksByTargetRequest) SetTargetId(v string) *ListScanTasksByTargetRequest {
	s.TargetId = &v
	return s
}

func (s *ListScanTasksByTargetRequest) SetTaskStatus(v string) *ListScanTasksByTargetRequest {
	s.TaskStatus = &v
	return s
}

func (s *ListScanTasksByTargetRequest) Validate() error {
	return dara.Validate(s)
}
