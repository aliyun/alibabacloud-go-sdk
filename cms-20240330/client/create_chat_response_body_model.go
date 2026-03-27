// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*CreateChatResponseBodyMessages) *CreateChatResponseBody
	GetMessages() []*CreateChatResponseBodyMessages
	SetRequestId(v string) *CreateChatResponseBody
	GetRequestId() *string
	SetTraceId(v string) *CreateChatResponseBody
	GetTraceId() *string
}

type CreateChatResponseBody struct {
	Messages []*CreateChatResponseBodyMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// 0CEC5375-XXXX-XXXX-XXXX-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 21504600000008405622576e3b48
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CreateChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChatResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatResponseBody) GetMessages() []*CreateChatResponseBodyMessages {
	return s.Messages
}

func (s *CreateChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChatResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateChatResponseBody) SetMessages(v []*CreateChatResponseBodyMessages) *CreateChatResponseBody {
	s.Messages = v
	return s
}

func (s *CreateChatResponseBody) SetRequestId(v string) *CreateChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChatResponseBody) SetTraceId(v string) *CreateChatResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateChatResponseBody) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateChatResponseBodyMessages struct {
	Agents    []map[string]interface{} `json:"agents,omitempty" xml:"agents,omitempty" type:"Repeated"`
	Artifacts []map[string]interface{} `json:"artifacts,omitempty" xml:"artifacts,omitempty" type:"Repeated"`
	// example:
	//
	// call_search_001
	CallId   *string                  `json:"callId,omitempty" xml:"callId,omitempty"`
	Contents []map[string]interface{} `json:"contents,omitempty" xml:"contents,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Detail *string                  `json:"detail,omitempty" xml:"detail,omitempty"`
	Events []map[string]interface{} `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// example:
	//
	// sess_abc123
	ParentCallId *string `json:"parentCallId,omitempty" xml:"parentCallId,omitempty"`
	// example:
	//
	// tool
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// 1
	Seq *int32 `json:"seq,omitempty" xml:"seq,omitempty"`
	// example:
	//
	// 1765000005
	Timestamp *string                  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Tools     []map[string]interface{} `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	// example:
	//
	// done
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateChatResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s CreateChatResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *CreateChatResponseBodyMessages) GetAgents() []map[string]interface{} {
	return s.Agents
}

func (s *CreateChatResponseBodyMessages) GetArtifacts() []map[string]interface{} {
	return s.Artifacts
}

func (s *CreateChatResponseBodyMessages) GetCallId() *string {
	return s.CallId
}

func (s *CreateChatResponseBodyMessages) GetContents() []map[string]interface{} {
	return s.Contents
}

func (s *CreateChatResponseBodyMessages) GetDetail() *string {
	return s.Detail
}

func (s *CreateChatResponseBodyMessages) GetEvents() []map[string]interface{} {
	return s.Events
}

func (s *CreateChatResponseBodyMessages) GetParentCallId() *string {
	return s.ParentCallId
}

func (s *CreateChatResponseBodyMessages) GetRole() *string {
	return s.Role
}

func (s *CreateChatResponseBodyMessages) GetSeq() *int32 {
	return s.Seq
}

func (s *CreateChatResponseBodyMessages) GetTimestamp() *string {
	return s.Timestamp
}

func (s *CreateChatResponseBodyMessages) GetTools() []map[string]interface{} {
	return s.Tools
}

func (s *CreateChatResponseBodyMessages) GetType() *string {
	return s.Type
}

func (s *CreateChatResponseBodyMessages) GetVersion() *string {
	return s.Version
}

func (s *CreateChatResponseBodyMessages) SetAgents(v []map[string]interface{}) *CreateChatResponseBodyMessages {
	s.Agents = v
	return s
}

func (s *CreateChatResponseBodyMessages) SetArtifacts(v []map[string]interface{}) *CreateChatResponseBodyMessages {
	s.Artifacts = v
	return s
}

func (s *CreateChatResponseBodyMessages) SetCallId(v string) *CreateChatResponseBodyMessages {
	s.CallId = &v
	return s
}

func (s *CreateChatResponseBodyMessages) SetContents(v []map[string]interface{}) *CreateChatResponseBodyMessages {
	s.Contents = v
	return s
}

func (s *CreateChatResponseBodyMessages) SetDetail(v string) *CreateChatResponseBodyMessages {
	s.Detail = &v
	return s
}

func (s *CreateChatResponseBodyMessages) SetEvents(v []map[string]interface{}) *CreateChatResponseBodyMessages {
	s.Events = v
	return s
}

func (s *CreateChatResponseBodyMessages) SetParentCallId(v string) *CreateChatResponseBodyMessages {
	s.ParentCallId = &v
	return s
}

func (s *CreateChatResponseBodyMessages) SetRole(v string) *CreateChatResponseBodyMessages {
	s.Role = &v
	return s
}

func (s *CreateChatResponseBodyMessages) SetSeq(v int32) *CreateChatResponseBodyMessages {
	s.Seq = &v
	return s
}

func (s *CreateChatResponseBodyMessages) SetTimestamp(v string) *CreateChatResponseBodyMessages {
	s.Timestamp = &v
	return s
}

func (s *CreateChatResponseBodyMessages) SetTools(v []map[string]interface{}) *CreateChatResponseBodyMessages {
	s.Tools = v
	return s
}

func (s *CreateChatResponseBodyMessages) SetType(v string) *CreateChatResponseBodyMessages {
	s.Type = &v
	return s
}

func (s *CreateChatResponseBodyMessages) SetVersion(v string) *CreateChatResponseBodyMessages {
	s.Version = &v
	return s
}

func (s *CreateChatResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
