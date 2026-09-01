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
	// The page number of the current page when using paging.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EventStatus *int32 `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// The maximum number of entries per page when using paging.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The server name or IP address.
	//
	// example:
	//
	// test-ecs
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The running status of the machine. Valid values:
	//
	// - **monitoring**: Warning.
	//
	// - **blocking**: Blocking.
	//
	// - **studying**: Learning.
	//
	// - **study_finish**: Learning completed.
	//
	// example:
	//
	// studying
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The whitelist mode. Valid values:
	//
	// - **hash**: process hash
	//
	// - **path**: process path
	//
	// example:
	//
	// hash
	StudyMode *string `json:"StudyMode,omitempty" xml:"StudyMode,omitempty"`
	// The end of the model creation time range. The value is a timestamp in milliseconds.
	//
	// example:
	//
	// 1768891966346
	StudyTimeEnd *int64 `json:"StudyTimeEnd,omitempty" xml:"StudyTimeEnd,omitempty"`
	// The start of the model creation time range. The value is a timestamp in milliseconds.
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
