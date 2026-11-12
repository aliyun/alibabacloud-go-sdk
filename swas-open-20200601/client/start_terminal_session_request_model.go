// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTerminalSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandLine(v string) *StartTerminalSessionRequest
	GetCommandLine() *string
	SetInstanceId(v string) *StartTerminalSessionRequest
	GetInstanceId() *string
	SetRegionId(v string) *StartTerminalSessionRequest
	GetRegionId() *string
}

type StartTerminalSessionRequest struct {
	// The command to run after the session is started. The command can be up to 512 characters in length.
	//
	// example:
	//
	// ssh root@192.168.0.***
	CommandLine *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	// The ID of the Simple Application Server instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Simple Application Server instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartTerminalSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTerminalSessionRequest) GoString() string {
	return s.String()
}

func (s *StartTerminalSessionRequest) GetCommandLine() *string {
	return s.CommandLine
}

func (s *StartTerminalSessionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartTerminalSessionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartTerminalSessionRequest) SetCommandLine(v string) *StartTerminalSessionRequest {
	s.CommandLine = &v
	return s
}

func (s *StartTerminalSessionRequest) SetInstanceId(v string) *StartTerminalSessionRequest {
	s.InstanceId = &v
	return s
}

func (s *StartTerminalSessionRequest) SetRegionId(v string) *StartTerminalSessionRequest {
	s.RegionId = &v
	return s
}

func (s *StartTerminalSessionRequest) Validate() error {
	return dara.Validate(s)
}
