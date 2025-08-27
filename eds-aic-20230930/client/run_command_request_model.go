// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *RunCommandRequest
	GetAgentType() *string
	SetCommandContent(v string) *RunCommandRequest
	GetCommandContent() *string
	SetContentEncoding(v string) *RunCommandRequest
	GetContentEncoding() *string
	SetInstanceIds(v []*string) *RunCommandRequest
	GetInstanceIds() []*string
	SetTimeout(v int64) *RunCommandRequest
	GetTimeout() *int64
}

type RunCommandRequest struct {
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The content of the command.
	//
	// example:
	//
	// ls
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The encoding method of the command content (`CommandContent`). The value is not case-sensitive.
	//
	// >  If you set the value to an invalid encoding method, the system will process the command content as `PlainText`.
	//
	// Valid values:
	//
	// 	- Base64: encodes the command content in Base64.
	//
	// 	- PlainText (default): does not encode the command content. The command content is input as plain text.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The IDs of the cloud phone instances. You can specify a maximum of 50 cloud phone instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The timeout period of the command execution. If the command execution exceeds the timeout period, it will be considered timed out. If you leave this parameter empty, it defaults to 60.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s RunCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *RunCommandRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *RunCommandRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *RunCommandRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RunCommandRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *RunCommandRequest) SetAgentType(v string) *RunCommandRequest {
	s.AgentType = &v
	return s
}

func (s *RunCommandRequest) SetCommandContent(v string) *RunCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandRequest) SetContentEncoding(v string) *RunCommandRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunCommandRequest) SetInstanceIds(v []*string) *RunCommandRequest {
	s.InstanceIds = v
	return s
}

func (s *RunCommandRequest) SetTimeout(v int64) *RunCommandRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandRequest) Validate() error {
	return dara.Validate(s)
}
