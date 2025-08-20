// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetStackRequest
	GetClientToken() *string
	SetLogOption(v string) *GetStackRequest
	GetLogOption() *string
	SetOutputOption(v string) *GetStackRequest
	GetOutputOption() *string
	SetRegionId(v string) *GetStackRequest
	GetRegionId() *string
	SetShowResourceProgress(v string) *GetStackRequest
	GetShowResourceProgress() *string
	SetStackId(v string) *GetStackRequest
	GetStackId() *string
}

type GetStackRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\\
	//
	// The token can be up to 64 characters in length.\\
	//
	// For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The option for returning logs. Valid values:
	//
	// 	- None: does not return logs.
	//
	// 	- Stack (default): returns the logs of the stack.
	//
	// 	- Resource: returns the logs of resources in the stack.
	//
	// 	- All: returns all logs.
	//
	// example:
	//
	// Stack
	LogOption *string `json:"LogOption,omitempty" xml:"LogOption,omitempty"`
	// Specifies whether to return Outputs. Valid values:
	//
	// 	- Enabled (default)
	//
	// 	- Disabled
	//
	// >  The Outputs parameter requires a long period of time to calculate. If you do not require Outputs of the stack, we recommend that you set OutputOption to Disabled to improve the response speed of the GetStack operation.
	//
	// example:
	//
	// Disabled
	OutputOption *string `json:"OutputOption,omitempty" xml:"OutputOption,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to return information about ResourceProgress. Valid values:
	//
	// 	- Disabled (default): does not return information about ResourceProgress.
	//
	// 	- PercentageOnly: returns StackOperationProgress and StackActionProgress of ResourceProgress.
	//
	// >  ROS and Terraform stacks are supported. Creation, resumed creation, update, deletion, import, and rollback operations on stacks are supported.
	//
	// 	- EnabledIfCreateStack (not recommend): returns \\*Count and InProgressResourceDetails of ResourceProgress only during a stack creation operation.
	//
	// >  During a creation operation, a stack is in one of the following states: CREATE_IN_PROGRESS, CREATE_COMPLETE, CREATE_FAILED, CREATE_ROLLBACK_IN_PROGRESS, CREATE_ROLLBACK_COMPLETE, and CREATE_ROLLBACK_FAILED.
	//
	// example:
	//
	// Disabled
	ShowResourceProgress *string `json:"ShowResourceProgress,omitempty" xml:"ShowResourceProgress,omitempty"`
	// The stack ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c754d2a4-28f1-46df-b557-9586173a****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s GetStackRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStackRequest) GoString() string {
	return s.String()
}

func (s *GetStackRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetStackRequest) GetLogOption() *string {
	return s.LogOption
}

func (s *GetStackRequest) GetOutputOption() *string {
	return s.OutputOption
}

func (s *GetStackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetStackRequest) GetShowResourceProgress() *string {
	return s.ShowResourceProgress
}

func (s *GetStackRequest) GetStackId() *string {
	return s.StackId
}

func (s *GetStackRequest) SetClientToken(v string) *GetStackRequest {
	s.ClientToken = &v
	return s
}

func (s *GetStackRequest) SetLogOption(v string) *GetStackRequest {
	s.LogOption = &v
	return s
}

func (s *GetStackRequest) SetOutputOption(v string) *GetStackRequest {
	s.OutputOption = &v
	return s
}

func (s *GetStackRequest) SetRegionId(v string) *GetStackRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackRequest) SetShowResourceProgress(v string) *GetStackRequest {
	s.ShowResourceProgress = &v
	return s
}

func (s *GetStackRequest) SetStackId(v string) *GetStackRequest {
	s.StackId = &v
	return s
}

func (s *GetStackRequest) Validate() error {
	return dara.Validate(s)
}
