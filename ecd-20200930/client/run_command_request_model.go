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
	// The content of the command. The command content can be plaintext or Base64-encoded.\\
	//
	// The Base64-encoded command content cannot exceed 16 KB in size.
	//
	// > If the command content is Base64-encoded, you must set the ContentEncoding parameter to Base64.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipconfig
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	CommandRole    *string `json:"CommandRole,omitempty" xml:"CommandRole,omitempty"`
	// The encoding mode of the command content. Valid values:
	//
	// 	- PlainText: The command content is not encoded.
	//
	// 	- Base64: The command content is Base64-encoded.
	//
	// Default value: PlainText. If the specified value of this parameter is invalid, PlainText is used by default.
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The ID of cloud desktop N. Valid values of N: 1 to 50.\\
	//
	// If multiple cloud desktops are specified and the command execution succeeds on at least one of the cloud desktops, the operation is considered successful. If multiple cloud desktops are specified and the command execution fails on all the cloud desktops, verify the value of the parameter and try again.
	//
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The ID of the end user. If you specify a value, you run the command as the end user that is granted specific permissions. Note: The end user has sessions on a cloud computer. That is, when the cloud computer is started, the end user logs on to an Alibaba Cloud Workspace client and connects to the cloud computer, and the cloud computer is not preempted by another end user during the connection. This parameter is not available for Linux cloud computers.
	//
	// example:
	//
	// User1
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timeout period for the command to run. Unit: seconds. Default value: 60.\\
	//
	// A timeout error occurs if the command cannot be run because the process slows down or because a specific module or the Cloud Assistant client does not exist. When a timeout error occurs, the command process is forcibly terminated.
	//
	// example:
	//
	// 3600
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The language of the O\\&M command. Valid values:
	//
	// 	- RunBatScript
	//
	// 	- RunPowerShellScript
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
