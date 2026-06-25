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
	// The channel type for running the command.
	//
	// example:
	//
	// EdsAgent
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The content of the command.
	//
	// example:
	//
	// ls
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The encoding method for the command content (`CommandContent`). This value is not case-sensitive.
	//
	// > An invalid value defaults to `PlainText`.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// A list of instance IDs. You can specify up to 50 instances per request.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The execution timeout in seconds. The command times out if it does not complete within this period. Defaults to 60 seconds.
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
