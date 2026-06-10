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
	// The script content in plaintext or Base64-encoded format.<br>
	//
	// The Base64-encoded script content cannot exceed 16 KB.<br>
	//
	// > If the script content is Base64-encoded, you must set the `ContentEncoding` parameter to `Base64`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipconfig
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	CommandRole    *string `json:"CommandRole,omitempty" xml:"CommandRole,omitempty"`
	// The encoding mode of the script content.
	//
	// > If you specify a value that is not a valid enumeration member, the system defaults to `PlainText`.
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The IDs of the cloud computers on which to run the script. You can specify up to 50 IDs.<br>
	//
	// The API call is considered successful if the script runs on at least one of the specified cloud computers. The call fails only if the script fails on all of them.<br>
	//
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// If you specify this parameter, the command runs with the permissions of the specified end user.
	//
	// > This user must have a session history on the cloud computer. This means the user must have logged in after the cloud computer started and their session was not taken over by another user. This parameter is not supported for Linux cloud computers.
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
	// The script execution timeout, in seconds. Default value: 300.<br>
	//
	// A command times out if the script cannot be run due to issues such as process conflicts, missing modules, or an unavailable Cloud Assistant client. When a command times out, the system forcibly terminates the script process.<br>
	//
	// example:
	//
	// 3600
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The type of the script.
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
