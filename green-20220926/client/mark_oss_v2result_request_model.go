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
	// The end time. Specify the time in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-10-21 16:08:38 +0800
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The freeze type. This parameter is required when Operation is set to freeze. Valid values:
	//
	// - ACL: Modifies file permissions.
	//
	// - COPY: Moves the file directory. Destination directory rules: 1. The file directory selected during task creation takes precedence. 2. If automatic freezing is not enabled during creation, or if ACL freezing is configured, the directory selected in the console during freezing is used. 3. The default directory is alicip_riskfile_backup/.
	//
	// example:
	//
	// ACL
	FreezeType *string `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
	// The handling operation. Valid values:
	//
	// - freeze: Freezes the object.
	//
	// - unfreeze: Unfreezes the object.
	//
	// - misreport: Reports a false positive.
	//
	// - missOut: Reports a missed violation.
	//
	// example:
	//
	// freeze
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The request IDs.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestIds *string `json:"RequestIds,omitempty" xml:"RequestIds,omitempty"`
	// The start time. Specify the time in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-08-21 16:08:38 +0800
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The task name.
	//
	// example:
	//
	// Image inventory task 20240914100517757
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
