// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *InvokeCommandRequest
	GetCommandId() *string
	SetInstanceIds(v string) *InvokeCommandRequest
	GetInstanceIds() *string
	SetParameters(v map[string]interface{}) *InvokeCommandRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *InvokeCommandRequest
	GetRegionId() *string
	SetUsername(v string) *InvokeCommandRequest
	GetUsername() *string
}

type InvokeCommandRequest struct {
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
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s InvokeCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeCommandRequest) GoString() string {
	return s.String()
}

func (s *InvokeCommandRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *InvokeCommandRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *InvokeCommandRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *InvokeCommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InvokeCommandRequest) GetUsername() *string {
	return s.Username
}

func (s *InvokeCommandRequest) SetCommandId(v string) *InvokeCommandRequest {
	s.CommandId = &v
	return s
}

func (s *InvokeCommandRequest) SetInstanceIds(v string) *InvokeCommandRequest {
	s.InstanceIds = &v
	return s
}

func (s *InvokeCommandRequest) SetParameters(v map[string]interface{}) *InvokeCommandRequest {
	s.Parameters = v
	return s
}

func (s *InvokeCommandRequest) SetRegionId(v string) *InvokeCommandRequest {
	s.RegionId = &v
	return s
}

func (s *InvokeCommandRequest) SetUsername(v string) *InvokeCommandRequest {
	s.Username = &v
	return s
}

func (s *InvokeCommandRequest) Validate() error {
	return dara.Validate(s)
}
