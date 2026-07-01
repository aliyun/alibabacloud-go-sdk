// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRCSMobileCapableTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *QueryRCSMobileCapableTaskResultRequest
	GetTaskId() *string
}

type QueryRCSMobileCapableTaskResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryRCSMobileCapableTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRCSMobileCapableTaskResultRequest) GoString() string {
	return s.String()
}

func (s *QueryRCSMobileCapableTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryRCSMobileCapableTaskResultRequest) SetTaskId(v string) *QueryRCSMobileCapableTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *QueryRCSMobileCapableTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
