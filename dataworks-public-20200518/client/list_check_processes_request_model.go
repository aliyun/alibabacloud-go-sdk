// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckProcessesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventCode(v string) *ListCheckProcessesRequest
	GetEventCode() *string
	SetMessageId(v string) *ListCheckProcessesRequest
	GetMessageId() *string
	SetOperator(v string) *ListCheckProcessesRequest
	GetOperator() *string
	SetPageNumber(v int32) *ListCheckProcessesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCheckProcessesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListCheckProcessesRequest
	GetProjectId() *int64
	SetStatus(v string) *ListCheckProcessesRequest
	GetStatus() *string
}

type ListCheckProcessesRequest struct {
	// Extension point event encoding.
	//
	// This parameter is required.
	//
	// example:
	//
	// commit-file
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// The message ID in DataWorks OpenEvent. You can obtain the ID from a received message when an extension point event is triggered.
	//
	// example:
	//
	// 03400b03-b721-4c34-8727-2****1
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The operator ID.
	//
	// example:
	//
	// 123333232
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 123465
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The check status of the extension. Valid values:
	//
	// 	- CHECKING
	//
	// 	- PASSED
	//
	// 	- BLOCKED
	//
	// example:
	//
	// True
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCheckProcessesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCheckProcessesRequest) GoString() string {
	return s.String()
}

func (s *ListCheckProcessesRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *ListCheckProcessesRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *ListCheckProcessesRequest) GetOperator() *string {
	return s.Operator
}

func (s *ListCheckProcessesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCheckProcessesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckProcessesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCheckProcessesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCheckProcessesRequest) SetEventCode(v string) *ListCheckProcessesRequest {
	s.EventCode = &v
	return s
}

func (s *ListCheckProcessesRequest) SetMessageId(v string) *ListCheckProcessesRequest {
	s.MessageId = &v
	return s
}

func (s *ListCheckProcessesRequest) SetOperator(v string) *ListCheckProcessesRequest {
	s.Operator = &v
	return s
}

func (s *ListCheckProcessesRequest) SetPageNumber(v int32) *ListCheckProcessesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCheckProcessesRequest) SetPageSize(v int32) *ListCheckProcessesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCheckProcessesRequest) SetProjectId(v int64) *ListCheckProcessesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListCheckProcessesRequest) SetStatus(v string) *ListCheckProcessesRequest {
	s.Status = &v
	return s
}

func (s *ListCheckProcessesRequest) Validate() error {
	return dara.Validate(s)
}
