// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarkOssV2ResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *MarkOssV2ResultRequest
	GetEndDate() *string
	SetFreezeType(v string) *MarkOssV2ResultRequest
	GetFreezeType() *string
	SetOperation(v string) *MarkOssV2ResultRequest
	GetOperation() *string
	SetRequestIds(v string) *MarkOssV2ResultRequest
	GetRequestIds() *string
	SetStartDate(v string) *MarkOssV2ResultRequest
	GetStartDate() *string
	SetTaskName(v string) *MarkOssV2ResultRequest
	GetTaskName() *string
}

type MarkOssV2ResultRequest struct {
	// The end time.
	//
	// example:
	//
	// 2023-10-21 16:08:38 +0800
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The freeze type. This parameter is required when Operation is set to freeze. Valid values:
	//
	// - ACL: Modify file permissions.
	//
	// - COPY: Move the file to a directory. The destination directory is determined as follows: 1. The directory selected when the task was created takes priority. 2. If automatic freezing was not enabled during creation, or ACL freezing was configured, the directory selected during freezing in the console is used. 3. The default directory is alicip_riskfile_backup/.
	//
	// example:
	//
	// ACL
	FreezeType *string `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
	// The processing operation. Valid values:
	//
	// - freeze: freeze.
	//
	// - unfreeze: unfreeze.
	//
	// - misreport: false positive (not in violation).
	//
	// - missOut: missed violation.
	//
	// example:
	//
	// freeze
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestIds *string `json:"RequestIds,omitempty" xml:"RequestIds,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2023-08-21 16:08:38 +0800
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The task name.
	//
	// example:
	//
	// Image stock task 20240914100517757
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s MarkOssV2ResultRequest) String() string {
	return dara.Prettify(s)
}

func (s MarkOssV2ResultRequest) GoString() string {
	return s.String()
}

func (s *MarkOssV2ResultRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *MarkOssV2ResultRequest) GetFreezeType() *string {
	return s.FreezeType
}

func (s *MarkOssV2ResultRequest) GetOperation() *string {
	return s.Operation
}

func (s *MarkOssV2ResultRequest) GetRequestIds() *string {
	return s.RequestIds
}

func (s *MarkOssV2ResultRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *MarkOssV2ResultRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *MarkOssV2ResultRequest) SetEndDate(v string) *MarkOssV2ResultRequest {
	s.EndDate = &v
	return s
}

func (s *MarkOssV2ResultRequest) SetFreezeType(v string) *MarkOssV2ResultRequest {
	s.FreezeType = &v
	return s
}

func (s *MarkOssV2ResultRequest) SetOperation(v string) *MarkOssV2ResultRequest {
	s.Operation = &v
	return s
}

func (s *MarkOssV2ResultRequest) SetRequestIds(v string) *MarkOssV2ResultRequest {
	s.RequestIds = &v
	return s
}

func (s *MarkOssV2ResultRequest) SetStartDate(v string) *MarkOssV2ResultRequest {
	s.StartDate = &v
	return s
}

func (s *MarkOssV2ResultRequest) SetTaskName(v string) *MarkOssV2ResultRequest {
	s.TaskName = &v
	return s
}

func (s *MarkOssV2ResultRequest) Validate() error {
	return dara.Validate(s)
}
