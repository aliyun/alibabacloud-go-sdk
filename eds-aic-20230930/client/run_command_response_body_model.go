// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvokeId(v string) *RunCommandResponseBody
	GetInvokeId() *string
	SetRequestId(v string) *RunCommandResponseBody
	GetRequestId() *string
	SetRunCommandInfos(v []*RunCommandResponseBodyRunCommandInfos) *RunCommandResponseBody
	GetRunCommandInfos() []*RunCommandResponseBodyRunCommandInfos
	SetTaskId(v string) *RunCommandResponseBody
	GetTaskId() *string
}

type RunCommandResponseBody struct {
	// The command execution ID. You can use this ID to query the command execution result.
	//
	// example:
	//
	// t-gov2ujrk32v4****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 440D7342-5E7C-B2DB-D0B4EAC2BDF1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of command execution information for each cloud phone instance.
	RunCommandInfos []*RunCommandResponseBodyRunCommandInfos `json:"RunCommandInfos,omitempty" xml:"RunCommandInfos,omitempty" type:"Repeated"`
	// The batch task ID.
	//
	// example:
	//
	// t-bp67acfmxazb4p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RunCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBody) GetInvokeId() *string {
	return s.InvokeId
}

func (s *RunCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunCommandResponseBody) GetRunCommandInfos() []*RunCommandResponseBodyRunCommandInfos {
	return s.RunCommandInfos
}

func (s *RunCommandResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RunCommandResponseBody) SetInvokeId(v string) *RunCommandResponseBody {
	s.InvokeId = &v
	return s
}

func (s *RunCommandResponseBody) SetRequestId(v string) *RunCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCommandResponseBody) SetRunCommandInfos(v []*RunCommandResponseBodyRunCommandInfos) *RunCommandResponseBody {
	s.RunCommandInfos = v
	return s
}

func (s *RunCommandResponseBody) SetTaskId(v string) *RunCommandResponseBody {
	s.TaskId = &v
	return s
}

func (s *RunCommandResponseBody) Validate() error {
	if s.RunCommandInfos != nil {
		for _, item := range s.RunCommandInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunCommandResponseBodyRunCommandInfos struct {
	// The instance-level child task ID.
	//
	// example:
	//
	// t-bp67acfmxazb4p****
	ChildTaskId *string `json:"ChildTaskId,omitempty" xml:"ChildTaskId,omitempty"`
	// The cloud phone instance ID.
	//
	// example:
	//
	// acp-uto81vfd8t8z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The command execution ID.
	//
	// example:
	//
	// B8ED2BA9-0C6A-5643-818F-B5D60A64****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
}

func (s RunCommandResponseBodyRunCommandInfos) String() string {
	return dara.Prettify(s)
}

func (s RunCommandResponseBodyRunCommandInfos) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBodyRunCommandInfos) GetChildTaskId() *string {
	return s.ChildTaskId
}

func (s *RunCommandResponseBodyRunCommandInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RunCommandResponseBodyRunCommandInfos) GetInvokeId() *string {
	return s.InvokeId
}

func (s *RunCommandResponseBodyRunCommandInfos) SetChildTaskId(v string) *RunCommandResponseBodyRunCommandInfos {
	s.ChildTaskId = &v
	return s
}

func (s *RunCommandResponseBodyRunCommandInfos) SetInstanceId(v string) *RunCommandResponseBodyRunCommandInfos {
	s.InstanceId = &v
	return s
}

func (s *RunCommandResponseBodyRunCommandInfos) SetInvokeId(v string) *RunCommandResponseBodyRunCommandInfos {
	s.InvokeId = &v
	return s
}

func (s *RunCommandResponseBodyRunCommandInfos) Validate() error {
	return dara.Validate(s)
}
