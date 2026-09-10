// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyntaxCheckAndTransformSqlConversionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v int64) *SyntaxCheckAndTransformSqlConversionTaskRequest
	GetTaskId() *int64
}

type SyntaxCheckAndTransformSqlConversionTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SyntaxCheckAndTransformSqlConversionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SyntaxCheckAndTransformSqlConversionTaskRequest) GoString() string {
	return s.String()
}

func (s *SyntaxCheckAndTransformSqlConversionTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *SyntaxCheckAndTransformSqlConversionTaskRequest) SetTaskId(v int64) *SyntaxCheckAndTransformSqlConversionTaskRequest {
	s.TaskId = &v
	return s
}

func (s *SyntaxCheckAndTransformSqlConversionTaskRequest) Validate() error {
	return dara.Validate(s)
}
