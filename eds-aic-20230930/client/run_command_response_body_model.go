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
}

type RunCommandResponseBody struct {
	// The ID of the command execution. You can use the command execution ID to query the output of a command.
	//
	// example:
	//
	// t-gov2ujrk32v4****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 440D7342-5E7C-B2DB-D0B4EAC2BDF1****
	RequestId       *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RunCommandInfos []*RunCommandResponseBodyRunCommandInfos `json:"RunCommandInfos,omitempty" xml:"RunCommandInfos,omitempty" type:"Repeated"`
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

func (s *RunCommandResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunCommandResponseBodyRunCommandInfos struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InvokeId   *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
}

func (s RunCommandResponseBodyRunCommandInfos) String() string {
	return dara.Prettify(s)
}

func (s RunCommandResponseBodyRunCommandInfos) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBodyRunCommandInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RunCommandResponseBodyRunCommandInfos) GetInvokeId() *string {
	return s.InvokeId
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
