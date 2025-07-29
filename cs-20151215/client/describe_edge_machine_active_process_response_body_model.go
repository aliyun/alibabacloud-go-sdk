// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeMachineActiveProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v string) *DescribeEdgeMachineActiveProcessResponseBody
	GetLogs() *string
	SetProgress(v int64) *DescribeEdgeMachineActiveProcessResponseBody
	GetProgress() *int64
	SetRequestId(v string) *DescribeEdgeMachineActiveProcessResponseBody
	GetRequestId() *string
	SetState(v string) *DescribeEdgeMachineActiveProcessResponseBody
	GetState() *string
	SetStep(v string) *DescribeEdgeMachineActiveProcessResponseBody
	GetStep() *string
}

type DescribeEdgeMachineActiveProcessResponseBody struct {
	// The activation progress list.
	//
	// example:
	//
	// [{\"content\":\"步骤 \"颁发激活凭证\" 执行开始\",\"id\":0,\"level\":3,\"timestamp\":1625994913000},{\"content\":\"步骤 \"颁发激活凭证\" 执行成功\",\"id\":1,\"level\":3,\"timestamp\":1625994914000},{\"content\":\"步骤 \"初始化主机配置\" 执行开始\",\"id\":2,\"level\":3,\"timestamp\":1625994975000},{\"content\":\"步骤 \"初始化主机配置\" 执行成功\",\"id\":3,\"level\":3,\"timestamp\":1625994975000}]
	Logs *string `json:"logs,omitempty" xml:"logs,omitempty"`
	// The activation progress.
	//
	// example:
	//
	// 100
	Progress *int64 `json:"progress,omitempty" xml:"progress,omitempty"`
	// The request ID.
	//
	// example:
	//
	// b62796a8-c5a6-4d3f-beb2-7650e4309cb1
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The activation status.
	//
	// example:
	//
	// ACTIVATED
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The activation step.
	//
	// example:
	//
	// 步骤 \"初始化主机配置\" 执行成功
	Step *string `json:"step,omitempty" xml:"step,omitempty"`
}

func (s DescribeEdgeMachineActiveProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMachineActiveProcessResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMachineActiveProcessResponseBody) GetLogs() *string {
	return s.Logs
}

func (s *DescribeEdgeMachineActiveProcessResponseBody) GetProgress() *int64 {
	return s.Progress
}

func (s *DescribeEdgeMachineActiveProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEdgeMachineActiveProcessResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeEdgeMachineActiveProcessResponseBody) GetStep() *string {
	return s.Step
}

func (s *DescribeEdgeMachineActiveProcessResponseBody) SetLogs(v string) *DescribeEdgeMachineActiveProcessResponseBody {
	s.Logs = &v
	return s
}

func (s *DescribeEdgeMachineActiveProcessResponseBody) SetProgress(v int64) *DescribeEdgeMachineActiveProcessResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribeEdgeMachineActiveProcessResponseBody) SetRequestId(v string) *DescribeEdgeMachineActiveProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEdgeMachineActiveProcessResponseBody) SetState(v string) *DescribeEdgeMachineActiveProcessResponseBody {
	s.State = &v
	return s
}

func (s *DescribeEdgeMachineActiveProcessResponseBody) SetStep(v string) *DescribeEdgeMachineActiveProcessResponseBody {
	s.Step = &v
	return s
}

func (s *DescribeEdgeMachineActiveProcessResponseBody) Validate() error {
	return dara.Validate(s)
}
