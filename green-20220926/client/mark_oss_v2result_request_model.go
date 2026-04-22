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
	// example:
	//
	// 2023-10-21 16:08:38 +0800
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// ACL
	FreezeType *string `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
	// example:
	//
	// freeze
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestIds *string `json:"RequestIds,omitempty" xml:"RequestIds,omitempty"`
	// example:
	//
	// 2023-08-21 16:08:38 +0800
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TaskName  *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
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
