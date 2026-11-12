// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCommandAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *UpdateCommandAttributeRequest
	GetCommandId() *string
	SetDescription(v string) *UpdateCommandAttributeRequest
	GetDescription() *string
	SetName(v string) *UpdateCommandAttributeRequest
	GetName() *string
	SetRegionId(v string) *UpdateCommandAttributeRequest
	GetRegionId() *string
	SetTimeout(v int64) *UpdateCommandAttributeRequest
	GetTimeout() *int64
	SetWorkingDir(v string) *UpdateCommandAttributeRequest
	GetWorkingDir() *string
}

type UpdateCommandAttributeRequest struct {
	// The command ID. You can call the [DescribeCommands](https://help.aliyun.com/document_detail/64843.html) operation to query all available command IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-sh02yh0932w****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The description of the command. The description supports all character sets and can be up to 512 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the command. The name supports all character sets and can be up to 128 characters in length.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maximum timeout period for the command execution on the ECS instance. Unit: seconds. When a command that you created cannot be run, the command execution times out. When the execution times out, the command process is forcefully terminated and the PID of the command is canceled. Default value: 60.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The working directory of the command. The directory can be up to 200 characters in length.
	//
	// example:
	//
	// /home/
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s UpdateCommandAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommandAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateCommandAttributeRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *UpdateCommandAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCommandAttributeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCommandAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCommandAttributeRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *UpdateCommandAttributeRequest) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *UpdateCommandAttributeRequest) SetCommandId(v string) *UpdateCommandAttributeRequest {
	s.CommandId = &v
	return s
}

func (s *UpdateCommandAttributeRequest) SetDescription(v string) *UpdateCommandAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateCommandAttributeRequest) SetName(v string) *UpdateCommandAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateCommandAttributeRequest) SetRegionId(v string) *UpdateCommandAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateCommandAttributeRequest) SetTimeout(v int64) *UpdateCommandAttributeRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateCommandAttributeRequest) SetWorkingDir(v string) *UpdateCommandAttributeRequest {
	s.WorkingDir = &v
	return s
}

func (s *UpdateCommandAttributeRequest) Validate() error {
	return dara.Validate(s)
}
