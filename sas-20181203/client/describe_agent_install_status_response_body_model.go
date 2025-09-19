// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgentInstallStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAegisClientInvokeStatusResponseList(v []*DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList) *DescribeAgentInstallStatusResponseBody
	GetAegisClientInvokeStatusResponseList() []*DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList
	SetRequestId(v string) *DescribeAgentInstallStatusResponseBody
	GetRequestId() *string
}

type DescribeAgentInstallStatusResponseBody struct {
	// The status of servers.
	AegisClientInvokeStatusResponseList []*DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList `json:"AegisClientInvokeStatusResponseList,omitempty" xml:"AegisClientInvokeStatusResponseList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB3936FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAgentInstallStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgentInstallStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgentInstallStatusResponseBody) GetAegisClientInvokeStatusResponseList() []*DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList {
	return s.AegisClientInvokeStatusResponseList
}

func (s *DescribeAgentInstallStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgentInstallStatusResponseBody) SetAegisClientInvokeStatusResponseList(v []*DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList) *DescribeAgentInstallStatusResponseBody {
	s.AegisClientInvokeStatusResponseList = v
	return s
}

func (s *DescribeAgentInstallStatusResponseBody) SetRequestId(v string) *DescribeAgentInstallStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgentInstallStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList struct {
	// The returned message.
	//
	// example:
	//
	// Installed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The installation status. Valid value:
	//
	// 	- **-1**: The agent is not installed.
	//
	// 	- **0**: The agent is installed.
	//
	// 	- **1**: Failed to create a directory in the client.
	//
	// 	- **2**: Failed to download the installation package.
	//
	// 	- **3**: The installation file does not exist.
	//
	// 	- **4**: The verification information of the installation file does not exist.
	//
	// 	- **5**: Failed to verify the installation file.
	//
	// 	- **6**: Failed to execute the installation file.
	//
	// 	- **7**: You do not have the required permissions. The installation failed.
	//
	// 	- **8**: No client process is detected.
	//
	// 	- **100**: The installation failed due to an unknown error.
	//
	// 	- **1001**: The installation failed. One-click installation is not supported in this region.
	//
	// 	- **1002**: The installation failed. Servers that are not provided by Alibaba Cloud are not supported. Install the agent by executing a script on the server.
	//
	// 	- **1003**: The installation failed. The operating system is not supported.
	//
	// 	- **1004**: An internal error occurred. Try again later.
	//
	// 	- **1005**: The Elastic Compute Service (ECS) instance is not started. Start the ECS instance and try again.
	//
	// 	- **1006**: One-click installation is not supported for ECS instances of the classic network type.
	//
	// 	- **1007**: The running command is manually stopped.
	//
	// 	- **1008**: Cloud Assistant is not installed. You cannot install the client.
	//
	// 	- **1009**: The command execution timed out. Try again later.
	//
	// 	- **1010**: The machine is already online. You do not need to install a client.
	//
	// example:
	//
	// 1
	ResuleCode *string `json:"ResuleCode,omitempty" xml:"ResuleCode,omitempty"`
	// The installation result. Valid value:
	//
	// 	- **-1**: The agent is not installed.
	//
	// 	- **0**: The agent is being installed.
	//
	// 	- **1**: The agent is installed.
	//
	// 	- **2**: The installation failed.
	//
	// example:
	//
	// 0
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// d123f6ae-9749-4338-8c7f-3c2c1ead****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList) GoString() string {
	return s.String()
}

func (s *DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList) GetMessage() *string {
	return s.Message
}

func (s *DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList) GetResuleCode() *string {
	return s.ResuleCode
}

func (s *DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList) GetResult() *int32 {
	return s.Result
}

func (s *DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList) SetMessage(v string) *DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList {
	s.Message = &v
	return s
}

func (s *DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList) SetResuleCode(v string) *DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList {
	s.ResuleCode = &v
	return s
}

func (s *DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList) SetResult(v int32) *DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList {
	s.Result = &v
	return s
}

func (s *DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList) SetUuid(v string) *DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList {
	s.Uuid = &v
	return s
}

func (s *DescribeAgentInstallStatusResponseBodyAegisClientInvokeStatusResponseList) Validate() error {
	return dara.Validate(s)
}
