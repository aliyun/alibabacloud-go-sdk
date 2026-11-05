// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnknownThreatDetectMachineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListUnknownThreatDetectMachineRequest
	GetCurrentPage() *int32
	SetEventStatus(v int32) *ListUnknownThreatDetectMachineRequest
	GetEventStatus() *int32
	SetPageSize(v int32) *ListUnknownThreatDetectMachineRequest
	GetPageSize() *int32
	SetRemark(v string) *ListUnknownThreatDetectMachineRequest
	GetRemark() *string
	SetStatus(v string) *ListUnknownThreatDetectMachineRequest
	GetStatus() *string
	SetStudyMode(v string) *ListUnknownThreatDetectMachineRequest
	GetStudyMode() *string
	SetStudyTimeEnd(v int64) *ListUnknownThreatDetectMachineRequest
	GetStudyTimeEnd() *int64
	SetStudyTimeStart(v int64) *ListUnknownThreatDetectMachineRequest
	GetStudyTimeStart() *int64
	SetUuid(v string) *ListUnknownThreatDetectMachineRequest
	GetUuid() *string
}

type ListUnknownThreatDetectMachineRequest struct {
	// The page number of the current page in a paging query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The status of the alert event. Valid values:
	//
	// - **1**: pending
	//
	// - **2**: ignored
	//
	// - **4**: confirmed
	//
	// - **8**: marked as false positive
	//
	// - **16**: processing
	//
	// - **32**: handled
	//
	// - **64**: expired
	//
	// - **604**: marked as false positive by the system
	//
	// example:
	//
	// 1
	EventStatus *int32 `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// The maximum number of entries to return on each page in a paging query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name or IP address of the server.
	//
	// example:
	//
	// test-ecs
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The running status of the server. Valid values:
	//
	// - **monitoring**: alerting
	//
	// - **blocking**: blocking
	//
	// - **studying**: learning
	//
	// - **study_finish**: learning completed
	//
	// example:
	//
	// studying
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The whitelist mode. Valid values:
	//
	// - **hash**: the hash of the process
	//
	// - **path**: the path of the process
	//
	// example:
	//
	// hash
	StudyMode *string `json:"StudyMode,omitempty" xml:"StudyMode,omitempty"`
	// The end time of the time range when the model was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1768891966346
	StudyTimeEnd *int64 `json:"StudyTimeEnd,omitempty" xml:"StudyTimeEnd,omitempty"`
	// The start time of the time range when the model was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1768891966344
	StudyTimeStart *int64 `json:"StudyTimeStart,omitempty" xml:"StudyTimeStart,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// d2d94e8b-bb25-4744-8004-1e08a53c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListUnknownThreatDetectMachineRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectMachineRequest) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectMachineRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUnknownThreatDetectMachineRequest) GetEventStatus() *int32 {
	return s.EventStatus
}

func (s *ListUnknownThreatDetectMachineRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUnknownThreatDetectMachineRequest) GetRemark() *string {
	return s.Remark
}

func (s *ListUnknownThreatDetectMachineRequest) GetStatus() *string {
	return s.Status
}

func (s *ListUnknownThreatDetectMachineRequest) GetStudyMode() *string {
	return s.StudyMode
}

func (s *ListUnknownThreatDetectMachineRequest) GetStudyTimeEnd() *int64 {
	return s.StudyTimeEnd
}

func (s *ListUnknownThreatDetectMachineRequest) GetStudyTimeStart() *int64 {
	return s.StudyTimeStart
}

func (s *ListUnknownThreatDetectMachineRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListUnknownThreatDetectMachineRequest) SetCurrentPage(v int32) *ListUnknownThreatDetectMachineRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUnknownThreatDetectMachineRequest) SetEventStatus(v int32) *ListUnknownThreatDetectMachineRequest {
	s.EventStatus = &v
	return s
}

func (s *ListUnknownThreatDetectMachineRequest) SetPageSize(v int32) *ListUnknownThreatDetectMachineRequest {
	s.PageSize = &v
	return s
}

func (s *ListUnknownThreatDetectMachineRequest) SetRemark(v string) *ListUnknownThreatDetectMachineRequest {
	s.Remark = &v
	return s
}

func (s *ListUnknownThreatDetectMachineRequest) SetStatus(v string) *ListUnknownThreatDetectMachineRequest {
	s.Status = &v
	return s
}

func (s *ListUnknownThreatDetectMachineRequest) SetStudyMode(v string) *ListUnknownThreatDetectMachineRequest {
	s.StudyMode = &v
	return s
}

func (s *ListUnknownThreatDetectMachineRequest) SetStudyTimeEnd(v int64) *ListUnknownThreatDetectMachineRequest {
	s.StudyTimeEnd = &v
	return s
}

func (s *ListUnknownThreatDetectMachineRequest) SetStudyTimeStart(v int64) *ListUnknownThreatDetectMachineRequest {
	s.StudyTimeStart = &v
	return s
}

func (s *ListUnknownThreatDetectMachineRequest) SetUuid(v string) *ListUnknownThreatDetectMachineRequest {
	s.Uuid = &v
	return s
}

func (s *ListUnknownThreatDetectMachineRequest) Validate() error {
	return dara.Validate(s)
}
