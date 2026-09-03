// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*GetSessionContentResponseBodyContent) *GetSessionContentResponseBody
	GetContent() []*GetSessionContentResponseBodyContent
	SetRequestId(v string) *GetSessionContentResponseBody
	GetRequestId() *string
	SetSessionId(v string) *GetSessionContentResponseBody
	GetSessionId() *string
}

type GetSessionContentResponseBody struct {
	// The conversation text content.
	Content []*GetSessionContentResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// UUID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// UUID
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s GetSessionContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSessionContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetSessionContentResponseBody) GetContent() []*GetSessionContentResponseBodyContent {
	return s.Content
}

func (s *GetSessionContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSessionContentResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *GetSessionContentResponseBody) SetContent(v []*GetSessionContentResponseBodyContent) *GetSessionContentResponseBody {
	s.Content = v
	return s
}

func (s *GetSessionContentResponseBody) SetRequestId(v string) *GetSessionContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSessionContentResponseBody) SetSessionId(v string) *GetSessionContentResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetSessionContentResponseBody) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSessionContentResponseBodyContent struct {
	// The detailed conversation content.
	AgentContents []*GetSessionContentResponseBodyContentAgentContents `json:"agentContents,omitempty" xml:"agentContents,omitempty" type:"Repeated"`
	// The time when the session occurred, in the yyyy-MM-dd HH:mm:ss,SSS format.
	//
	// example:
	//
	// 2026-09-03 04:08:30,637
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// Indicates whether tool confirmation is required.
	//
	// example:
	//
	// false
	ToolConfirm *bool `json:"toolConfirm,omitempty" xml:"toolConfirm,omitempty"`
	// The user content of the first message in the session.
	//
	// example:
	//
	// [{\\"role\\":\\"user\\",\\"content\\":\\"Hi OSS\\"}]
	UserContent *string `json:"userContent,omitempty" xml:"userContent,omitempty"`
}

func (s GetSessionContentResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetSessionContentResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetSessionContentResponseBodyContent) GetAgentContents() []*GetSessionContentResponseBodyContentAgentContents {
	return s.AgentContents
}

func (s *GetSessionContentResponseBodyContent) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetSessionContentResponseBodyContent) GetToolConfirm() *bool {
	return s.ToolConfirm
}

func (s *GetSessionContentResponseBodyContent) GetUserContent() *string {
	return s.UserContent
}

func (s *GetSessionContentResponseBodyContent) SetAgentContents(v []*GetSessionContentResponseBodyContentAgentContents) *GetSessionContentResponseBodyContent {
	s.AgentContents = v
	return s
}

func (s *GetSessionContentResponseBodyContent) SetTimestamp(v string) *GetSessionContentResponseBodyContent {
	s.Timestamp = &v
	return s
}

func (s *GetSessionContentResponseBodyContent) SetToolConfirm(v bool) *GetSessionContentResponseBodyContent {
	s.ToolConfirm = &v
	return s
}

func (s *GetSessionContentResponseBodyContent) SetUserContent(v string) *GetSessionContentResponseBodyContent {
	s.UserContent = &v
	return s
}

func (s *GetSessionContentResponseBodyContent) Validate() error {
	if s.AgentContents != nil {
		for _, item := range s.AgentContents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSessionContentResponseBodyContentAgentContents struct {
	// The detailed conversation chunk content. All chunks compose the complete response.
	//
	// example:
	//
	// {\\"id\\":\\"16aa7737-9e6a-4500-abd2-96c5f17d1188\\",\\"object\\":\\"chat.completion.chunk\\",\\"created\\":1788408524,\\"model\\":\\"oss-agent\\",\\"choices\\":[{\\"index\\":0,\\"delta\\":{\\"role\\":\\"assistant\\",\\"content\\":\\"，Hello， I am OSS Agent\\"}}]}
	AgentContent *string `json:"agentContent,omitempty" xml:"agentContent,omitempty"`
	// The time when the content was generated, in the yyyy-MM-dd HH:mm:ss,SSS format.
	//
	// example:
	//
	// 1774577589
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetSessionContentResponseBodyContentAgentContents) String() string {
	return dara.Prettify(s)
}

func (s GetSessionContentResponseBodyContentAgentContents) GoString() string {
	return s.String()
}

func (s *GetSessionContentResponseBodyContentAgentContents) GetAgentContent() *string {
	return s.AgentContent
}

func (s *GetSessionContentResponseBodyContentAgentContents) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetSessionContentResponseBodyContentAgentContents) SetAgentContent(v string) *GetSessionContentResponseBodyContentAgentContents {
	s.AgentContent = &v
	return s
}

func (s *GetSessionContentResponseBodyContentAgentContents) SetTimestamp(v string) *GetSessionContentResponseBodyContentAgentContents {
	s.Timestamp = &v
	return s
}

func (s *GetSessionContentResponseBodyContentAgentContents) Validate() error {
	return dara.Validate(s)
}
