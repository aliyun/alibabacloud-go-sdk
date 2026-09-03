// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandContent(v string) *RunCommandRequest
	GetCommandContent() *string
	SetCommandRole(v string) *RunCommandRequest
	GetCommandRole() *string
	SetContentEncoding(v string) *RunCommandRequest
	GetContentEncoding() *string
	SetDesktopId(v []*string) *RunCommandRequest
	GetDesktopId() []*string
	SetEndUserId(v string) *RunCommandRequest
	GetEndUserId() *string
	SetRegionId(v string) *RunCommandRequest
	GetRegionId() *string
	SetTimeout(v int64) *RunCommandRequest
	GetTimeout() *int64
	SetType(v string) *RunCommandRequest
	GetType() *string
}

type RunCommandRequest struct {
	// The plaintext or Base64-encoded content of the script. The Base64-encoded script content cannot exceed 16 KB.
	//
	// > If the script content is Base64-encoded, set the ContentEncoding parameter to Base64.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipconfig
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The role used when the command is executed on the cloud computer.
	//
	// example:
	//
	// system
	CommandRole *string `json:"CommandRole,omitempty" xml:"CommandRole,omitempty"`
	// The encoding method of the script content.
	//
	// > If the specified value is not within the valid values, the value is treated as PlainText.
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The IDs of cloud computers. Valid values of N: 1 to 50. If you specify multiple cloud computers, the API call succeeds as long as the script is successfully executed on at least one cloud computer. If the script fails to be executed on all specified cloud computers, reset this parameter.
	//
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The ID of the end user. If this parameter is specified, the command is executed with the permissions of the end user.
	//
	// > The user must have a session record on the cloud computer (the user has logged on and connected to the cloud computer after it is started, and the connection was not preempted by another user). This parameter is not supported for Linux cloud computers.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timeout period for executing the script. Unit: seconds. Default value: 300. A timeout may occur when the script cannot run due to process issues, missing modules, or missing Cloud Assistant Agent. After a timeout, the script process is forcefully terminated.
	//
	// example:
	//
	// 3600
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The language type of the O&M script.
	//
	// This parameter is required.
	//
	// example:
	//
	// RunPowerShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RunCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *RunCommandRequest) GetCommandRole() *string {
	return s.CommandRole
}

func (s *RunCommandRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *RunCommandRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *RunCommandRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *RunCommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunCommandRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *RunCommandRequest) GetType() *string {
	return s.Type
}

func (s *RunCommandRequest) SetCommandContent(v string) *RunCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandRequest) SetCommandRole(v string) *RunCommandRequest {
	s.CommandRole = &v
	return s
}

func (s *RunCommandRequest) SetContentEncoding(v string) *RunCommandRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunCommandRequest) SetDesktopId(v []*string) *RunCommandRequest {
	s.DesktopId = v
	return s
}

func (s *RunCommandRequest) SetEndUserId(v string) *RunCommandRequest {
	s.EndUserId = &v
	return s
}

func (s *RunCommandRequest) SetRegionId(v string) *RunCommandRequest {
	s.RegionId = &v
	return s
}

func (s *RunCommandRequest) SetTimeout(v int64) *RunCommandRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandRequest) SetType(v string) *RunCommandRequest {
	s.Type = &v
	return s
}

func (s *RunCommandRequest) Validate() error {
	return dara.Validate(s)
}
