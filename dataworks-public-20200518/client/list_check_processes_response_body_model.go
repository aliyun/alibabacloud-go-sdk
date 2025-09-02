// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckProcessesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListCheckProcessesResponseBodyPagingInfo) *ListCheckProcessesResponseBody
	GetPagingInfo() *ListCheckProcessesResponseBodyPagingInfo
	SetRequestId(v string) *ListCheckProcessesResponseBody
	GetRequestId() *string
}

type ListCheckProcessesResponseBody struct {
	// The pagination information.
	PagingInfo *ListCheckProcessesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0000-ABCD-EF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCheckProcessesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCheckProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCheckProcessesResponseBody) GetPagingInfo() *ListCheckProcessesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListCheckProcessesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCheckProcessesResponseBody) SetPagingInfo(v *ListCheckProcessesResponseBodyPagingInfo) *ListCheckProcessesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListCheckProcessesResponseBody) SetRequestId(v string) *ListCheckProcessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCheckProcessesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCheckProcessesResponseBodyPagingInfo struct {
	// The check details of the extension.
	CheckProcesses []*ListCheckProcessesResponseBodyPagingInfoCheckProcesses `json:"CheckProcesses,omitempty" xml:"CheckProcesses,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries displayed on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCheckProcessesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCheckProcessesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListCheckProcessesResponseBodyPagingInfo) GetCheckProcesses() []*ListCheckProcessesResponseBodyPagingInfoCheckProcesses {
	return s.CheckProcesses
}

func (s *ListCheckProcessesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCheckProcessesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckProcessesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCheckProcessesResponseBodyPagingInfo) SetCheckProcesses(v []*ListCheckProcessesResponseBodyPagingInfoCheckProcesses) *ListCheckProcessesResponseBodyPagingInfo {
	s.CheckProcesses = v
	return s
}

func (s *ListCheckProcessesResponseBodyPagingInfo) SetPageNumber(v int32) *ListCheckProcessesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListCheckProcessesResponseBodyPagingInfo) SetPageSize(v int32) *ListCheckProcessesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListCheckProcessesResponseBodyPagingInfo) SetTotalCount(v int32) *ListCheckProcessesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCheckProcessesResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListCheckProcessesResponseBodyPagingInfoCheckProcesses struct {
	// Extension point event encoding.
	//
	// example:
	//
	// commit-file
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// The name of the extension point event.
	//
	// example:
	//
	// DnsEvent
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The English name of the event.
	//
	// example:
	//
	// Pre-event for Node Commit
	EventNameEn *string `json:"EventNameEn,omitempty" xml:"EventNameEn,omitempty"`
	// DataWorks the message ID of the open message. After an extended point event is triggered, you can obtain the message ID from the received event message.
	//
	// example:
	//
	// b824a5de-4223-4315-af3e-c4449d236db4
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The operator ID.
	//
	// example:
	//
	// 297635
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The ID of the process instance.
	//
	// example:
	//
	// rdk_generate_d395da25-b0d3-4114-b2a5-d0247444a661_none_3496903_365203
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The name of the check object, such as the file name or node name.
	//
	// example:
	//
	// odps_sql_test
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// example:
	//
	// 32563
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the extender check.
	//
	// - CHECKING CHECKING
	//
	// - PASSED the pass check
	//
	// - BLOCKED check failed
	//
	// example:
	//
	// CHECKING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCheckProcessesResponseBodyPagingInfoCheckProcesses) String() string {
	return dara.Prettify(s)
}

func (s ListCheckProcessesResponseBodyPagingInfoCheckProcesses) GoString() string {
	return s.String()
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) GetEventCode() *string {
	return s.EventCode
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) GetEventName() *string {
	return s.EventName
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) GetEventNameEn() *string {
	return s.EventNameEn
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) GetMessageId() *string {
	return s.MessageId
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) GetOperator() *string {
	return s.Operator
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) GetProcessName() *string {
	return s.ProcessName
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) GetStatus() *string {
	return s.Status
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) SetEventCode(v string) *ListCheckProcessesResponseBodyPagingInfoCheckProcesses {
	s.EventCode = &v
	return s
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) SetEventName(v string) *ListCheckProcessesResponseBodyPagingInfoCheckProcesses {
	s.EventName = &v
	return s
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) SetEventNameEn(v string) *ListCheckProcessesResponseBodyPagingInfoCheckProcesses {
	s.EventNameEn = &v
	return s
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) SetMessageId(v string) *ListCheckProcessesResponseBodyPagingInfoCheckProcesses {
	s.MessageId = &v
	return s
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) SetOperator(v string) *ListCheckProcessesResponseBodyPagingInfoCheckProcesses {
	s.Operator = &v
	return s
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) SetProcessId(v string) *ListCheckProcessesResponseBodyPagingInfoCheckProcesses {
	s.ProcessId = &v
	return s
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) SetProcessName(v string) *ListCheckProcessesResponseBodyPagingInfoCheckProcesses {
	s.ProcessName = &v
	return s
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) SetProjectId(v int64) *ListCheckProcessesResponseBodyPagingInfoCheckProcesses {
	s.ProjectId = &v
	return s
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) SetStatus(v string) *ListCheckProcessesResponseBodyPagingInfoCheckProcesses {
	s.Status = &v
	return s
}

func (s *ListCheckProcessesResponseBodyPagingInfoCheckProcesses) Validate() error {
	return dara.Validate(s)
}
