// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeCommandShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *InvokeCommandShrinkRequest
	GetCommandId() *string
	SetInstanceIds(v string) *InvokeCommandShrinkRequest
	GetInstanceIds() *string
	SetParametersShrink(v string) *InvokeCommandShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *InvokeCommandShrinkRequest
	GetRegionId() *string
	SetUsername(v string) *InvokeCommandShrinkRequest
	GetUsername() *string
}

type InvokeCommandShrinkRequest struct {
	// The command ID. You can call the DescribeCommands operation to query all available command IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-sh02yh0932w****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 50 IDs of simple application servers. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc****", "2ad1ae67295445f598017123dc****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The key-value pairs of custom parameters to specify when the custom parameter feature is enabled.
	//
	// 	- You can specify up to 10 custom parameters. Each key in a Map collection cannot be an empty string and can be up to 64 characters in length.
	//
	// 	- Values in a Map collection can be empty strings. The total length of the custom parameters and the original command cannot exceed 18 KB after they are encoded in Base64.
	//
	// 	- The custom parameter names that you specify for the Parameters parameter must be included in the custom parameter names that you specified when you created the command.
	//
	// 	- You can use empty strings to represent the custom parameters that are not specified. If you want to disable the custom parameter feature, you can leave this parameter empty.
	//
	// example:
	//
	// {"delayed_insert_timeout":"600","max_length_for_sort_data":"2048"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the user who runs the command in a simple application server. The username cannot exceed 255 characters in length.
	//
	// 	- For Linux servers, the default value is the root username.
	//
	// 	- For Windows servers, the default value is the system username.
	//
	// You can change the user to run the command only for Linux simple application servers.
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s InvokeCommandShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *InvokeCommandShrinkRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *InvokeCommandShrinkRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *InvokeCommandShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *InvokeCommandShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InvokeCommandShrinkRequest) GetUsername() *string {
	return s.Username
}

func (s *InvokeCommandShrinkRequest) SetCommandId(v string) *InvokeCommandShrinkRequest {
	s.CommandId = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetInstanceIds(v string) *InvokeCommandShrinkRequest {
	s.InstanceIds = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetParametersShrink(v string) *InvokeCommandShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetRegionId(v string) *InvokeCommandShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetUsername(v string) *InvokeCommandShrinkRequest {
	s.Username = &v
	return s
}

func (s *InvokeCommandShrinkRequest) Validate() error {
	return dara.Validate(s)
}
