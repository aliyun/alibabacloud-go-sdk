// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDiagnosisJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListDataDiagnosisJobsRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListDataDiagnosisJobsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDataDiagnosisJobsRequest
	GetPageSize() *string
	SetStatus(v string) *ListDataDiagnosisJobsRequest
	GetStatus() *string
	SetTypes(v []*string) *ListDataDiagnosisJobsRequest
	GetTypes() []*string
}

type ListDataDiagnosisJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Initializing
	Status *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Types  []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s ListDataDiagnosisJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosisJobsRequest) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosisJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataDiagnosisJobsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataDiagnosisJobsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataDiagnosisJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDataDiagnosisJobsRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListDataDiagnosisJobsRequest) SetInstanceId(v string) *ListDataDiagnosisJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDataDiagnosisJobsRequest) SetPageNumber(v string) *ListDataDiagnosisJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataDiagnosisJobsRequest) SetPageSize(v string) *ListDataDiagnosisJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataDiagnosisJobsRequest) SetStatus(v string) *ListDataDiagnosisJobsRequest {
	s.Status = &v
	return s
}

func (s *ListDataDiagnosisJobsRequest) SetTypes(v []*string) *ListDataDiagnosisJobsRequest {
	s.Types = v
	return s
}

func (s *ListDataDiagnosisJobsRequest) Validate() error {
	return dara.Validate(s)
}
