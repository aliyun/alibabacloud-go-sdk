// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatAiAgentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ChatAiAgentHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ChatAiAgentHeaders
	GetWorkspace() *string
}

type ChatAiAgentHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 710d6a64d8****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ChatAiAgentHeaders) String() string {
	return dara.Prettify(s)
}

func (s ChatAiAgentHeaders) GoString() string {
	return s.String()
}

func (s *ChatAiAgentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ChatAiAgentHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ChatAiAgentHeaders) SetCommonHeaders(v map[string]*string) *ChatAiAgentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChatAiAgentHeaders) SetWorkspace(v string) *ChatAiAgentHeaders {
	s.Workspace = &v
	return s
}

func (s *ChatAiAgentHeaders) Validate() error {
	return dara.Validate(s)
}
