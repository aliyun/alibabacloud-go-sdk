// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportTaskValidationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstanceId(v string) *DescribeImportTaskValidationRequest
	GetDbInstanceId() *string
	SetOwnerId(v int64) *DescribeImportTaskValidationRequest
	GetOwnerId() *int64
	SetTaskId(v int64) *DescribeImportTaskValidationRequest
	GetTaskId() *int64
}

type DescribeImportTaskValidationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-xjkljj****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 41698****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeImportTaskValidationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportTaskValidationRequest) GoString() string {
	return s.String()
}

func (s *DescribeImportTaskValidationRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeImportTaskValidationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeImportTaskValidationRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeImportTaskValidationRequest) SetDbInstanceId(v string) *DescribeImportTaskValidationRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeImportTaskValidationRequest) SetOwnerId(v int64) *DescribeImportTaskValidationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeImportTaskValidationRequest) SetTaskId(v int64) *DescribeImportTaskValidationRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeImportTaskValidationRequest) Validate() error {
	return dara.Validate(s)
}
