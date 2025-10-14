// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunServiceScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommandResults(v *RunServiceScheduleResponseBodyCommandResults) *RunServiceScheduleResponseBody
	GetCommandResults() *RunServiceScheduleResponseBodyCommandResults
	SetIndex(v int32) *RunServiceScheduleResponseBody
	GetIndex() *int32
	SetInstanceId(v string) *RunServiceScheduleResponseBody
	GetInstanceId() *string
	SetInstanceIp(v string) *RunServiceScheduleResponseBody
	GetInstanceIp() *string
	SetInstancePort(v int32) *RunServiceScheduleResponseBody
	GetInstancePort() *int32
	SetRequestId(v string) *RunServiceScheduleResponseBody
	GetRequestId() *string
	SetRequestRepeated(v string) *RunServiceScheduleResponseBody
	GetRequestRepeated() *string
	SetTcpPorts(v bool) *RunServiceScheduleResponseBody
	GetTcpPorts() *bool
}

type RunServiceScheduleResponseBody struct {
	// The execution results of the commands.
	CommandResults *RunServiceScheduleResponseBodyCommandResults `json:"CommandResults,omitempty" xml:"CommandResults,omitempty" type:"Struct"`
	// The index number of the scheduled virtual device (pod).
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The ID of the scheduled instance.
	//
	// example:
	//
	// i-5qvji3mom4ec013dyygmtxlkj
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the scheduled instance.
	//
	// example:
	//
	// 172.16.246.11
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The start port of the scheduled instance.
	//
	// example:
	//
	// 1024
	InstancePort *int32 `json:"InstancePort,omitempty" xml:"InstancePort,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is repeated. This parameter is not returned if ServcieAction is set to Console.
	//
	// example:
	//
	// false
	RequestRepeated *string `json:"RequestRepeated,omitempty" xml:"RequestRepeated,omitempty"`
	// The TCP port range of the scheduled instance or container. The value is in the ${from}-$-{to} format. Example: 80-88.
	//
	// example:
	//
	// "80-88"
	TcpPorts *bool `json:"TcpPorts,omitempty" xml:"TcpPorts,omitempty"`
}

func (s RunServiceScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunServiceScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *RunServiceScheduleResponseBody) GetCommandResults() *RunServiceScheduleResponseBodyCommandResults {
	return s.CommandResults
}

func (s *RunServiceScheduleResponseBody) GetIndex() *int32 {
	return s.Index
}

func (s *RunServiceScheduleResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RunServiceScheduleResponseBody) GetInstanceIp() *string {
	return s.InstanceIp
}

func (s *RunServiceScheduleResponseBody) GetInstancePort() *int32 {
	return s.InstancePort
}

func (s *RunServiceScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunServiceScheduleResponseBody) GetRequestRepeated() *string {
	return s.RequestRepeated
}

func (s *RunServiceScheduleResponseBody) GetTcpPorts() *bool {
	return s.TcpPorts
}

func (s *RunServiceScheduleResponseBody) SetCommandResults(v *RunServiceScheduleResponseBodyCommandResults) *RunServiceScheduleResponseBody {
	s.CommandResults = v
	return s
}

func (s *RunServiceScheduleResponseBody) SetIndex(v int32) *RunServiceScheduleResponseBody {
	s.Index = &v
	return s
}

func (s *RunServiceScheduleResponseBody) SetInstanceId(v string) *RunServiceScheduleResponseBody {
	s.InstanceId = &v
	return s
}

func (s *RunServiceScheduleResponseBody) SetInstanceIp(v string) *RunServiceScheduleResponseBody {
	s.InstanceIp = &v
	return s
}

func (s *RunServiceScheduleResponseBody) SetInstancePort(v int32) *RunServiceScheduleResponseBody {
	s.InstancePort = &v
	return s
}

func (s *RunServiceScheduleResponseBody) SetRequestId(v string) *RunServiceScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunServiceScheduleResponseBody) SetRequestRepeated(v string) *RunServiceScheduleResponseBody {
	s.RequestRepeated = &v
	return s
}

func (s *RunServiceScheduleResponseBody) SetTcpPorts(v bool) *RunServiceScheduleResponseBody {
	s.TcpPorts = &v
	return s
}

func (s *RunServiceScheduleResponseBody) Validate() error {
	if s.CommandResults != nil {
		if err := s.CommandResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunServiceScheduleResponseBodyCommandResults struct {
	CommandResult []*RunServiceScheduleResponseBodyCommandResultsCommandResult `json:"CommandResult,omitempty" xml:"CommandResult,omitempty" type:"Repeated"`
}

func (s RunServiceScheduleResponseBodyCommandResults) String() string {
	return dara.Prettify(s)
}

func (s RunServiceScheduleResponseBodyCommandResults) GoString() string {
	return s.String()
}

func (s *RunServiceScheduleResponseBodyCommandResults) GetCommandResult() []*RunServiceScheduleResponseBodyCommandResultsCommandResult {
	return s.CommandResult
}

func (s *RunServiceScheduleResponseBodyCommandResults) SetCommandResult(v []*RunServiceScheduleResponseBodyCommandResultsCommandResult) *RunServiceScheduleResponseBodyCommandResults {
	s.CommandResult = v
	return s
}

func (s *RunServiceScheduleResponseBodyCommandResults) Validate() error {
	if s.CommandResult != nil {
		for _, item := range s.CommandResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunServiceScheduleResponseBodyCommandResultsCommandResult struct {
	// The command.
	//
	// example:
	//
	// ls -l /data
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The name of the container.
	//
	// example:
	//
	// android
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The execution result of the command.
	//
	// example:
	//
	// success
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
}

func (s RunServiceScheduleResponseBodyCommandResultsCommandResult) String() string {
	return dara.Prettify(s)
}

func (s RunServiceScheduleResponseBodyCommandResultsCommandResult) GoString() string {
	return s.String()
}

func (s *RunServiceScheduleResponseBodyCommandResultsCommandResult) GetCommand() *string {
	return s.Command
}

func (s *RunServiceScheduleResponseBodyCommandResultsCommandResult) GetContainerName() *string {
	return s.ContainerName
}

func (s *RunServiceScheduleResponseBodyCommandResultsCommandResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *RunServiceScheduleResponseBodyCommandResultsCommandResult) SetCommand(v string) *RunServiceScheduleResponseBodyCommandResultsCommandResult {
	s.Command = &v
	return s
}

func (s *RunServiceScheduleResponseBodyCommandResultsCommandResult) SetContainerName(v string) *RunServiceScheduleResponseBodyCommandResultsCommandResult {
	s.ContainerName = &v
	return s
}

func (s *RunServiceScheduleResponseBodyCommandResultsCommandResult) SetResultMsg(v string) *RunServiceScheduleResponseBodyCommandResultsCommandResult {
	s.ResultMsg = &v
	return s
}

func (s *RunServiceScheduleResponseBodyCommandResultsCommandResult) Validate() error {
	return dara.Validate(s)
}
