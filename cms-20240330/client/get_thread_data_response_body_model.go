// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetThreadDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetThreadDataResponseBodyData) *GetThreadDataResponseBody
	GetData() []*GetThreadDataResponseBodyData
	SetDigitalEmployeeName(v string) *GetThreadDataResponseBody
	GetDigitalEmployeeName() *string
	SetMaxResults(v int64) *GetThreadDataResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *GetThreadDataResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetThreadDataResponseBody
	GetRequestId() *string
	SetThreadId(v string) *GetThreadDataResponseBody
	GetThreadId() *string
}

type GetThreadDataResponseBody struct {
	Data []*GetThreadDataResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// test
	DigitalEmployeeName *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// example:
	//
	// 2
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// aliding_thread_448d05c048a3481f8c19bc1a6038f8f6
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s GetThreadDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetThreadDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetThreadDataResponseBody) GetData() []*GetThreadDataResponseBodyData {
	return s.Data
}

func (s *GetThreadDataResponseBody) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *GetThreadDataResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *GetThreadDataResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetThreadDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetThreadDataResponseBody) GetThreadId() *string {
	return s.ThreadId
}

func (s *GetThreadDataResponseBody) SetData(v []*GetThreadDataResponseBodyData) *GetThreadDataResponseBody {
	s.Data = v
	return s
}

func (s *GetThreadDataResponseBody) SetDigitalEmployeeName(v string) *GetThreadDataResponseBody {
	s.DigitalEmployeeName = &v
	return s
}

func (s *GetThreadDataResponseBody) SetMaxResults(v int64) *GetThreadDataResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetThreadDataResponseBody) SetNextToken(v string) *GetThreadDataResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetThreadDataResponseBody) SetRequestId(v string) *GetThreadDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetThreadDataResponseBody) SetThreadId(v string) *GetThreadDataResponseBody {
	s.ThreadId = &v
	return s
}

func (s *GetThreadDataResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetThreadDataResponseBodyData struct {
	Messages  []*GetThreadDataResponseBodyDataMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TraceId   *string                                  `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GetThreadDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetThreadDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetThreadDataResponseBodyData) GetMessages() []*GetThreadDataResponseBodyDataMessages {
	return s.Messages
}

func (s *GetThreadDataResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *GetThreadDataResponseBodyData) GetTraceId() *string {
	return s.TraceId
}

func (s *GetThreadDataResponseBodyData) SetMessages(v []*GetThreadDataResponseBodyDataMessages) *GetThreadDataResponseBodyData {
	s.Messages = v
	return s
}

func (s *GetThreadDataResponseBodyData) SetRequestId(v string) *GetThreadDataResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetThreadDataResponseBodyData) SetTraceId(v string) *GetThreadDataResponseBodyData {
	s.TraceId = &v
	return s
}

func (s *GetThreadDataResponseBodyData) Validate() error {
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

type GetThreadDataResponseBodyDataMessages struct {
	Agents       []map[string]interface{} `json:"agents,omitempty" xml:"agents,omitempty" type:"Repeated"`
	Artifacts    []map[string]interface{} `json:"artifacts,omitempty" xml:"artifacts,omitempty" type:"Repeated"`
	CallId       *string                  `json:"callId,omitempty" xml:"callId,omitempty"`
	Contents     []map[string]interface{} `json:"contents,omitempty" xml:"contents,omitempty" type:"Repeated"`
	Detail       *string                  `json:"detail,omitempty" xml:"detail,omitempty"`
	Events       []map[string]interface{} `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	ParentCallId *string                  `json:"parentCallId,omitempty" xml:"parentCallId,omitempty"`
	Role         *string                  `json:"role,omitempty" xml:"role,omitempty"`
	Seq          *int32                   `json:"seq,omitempty" xml:"seq,omitempty"`
	Timestamp    *string                  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Tools        []map[string]interface{} `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	Type         *string                  `json:"type,omitempty" xml:"type,omitempty"`
	Version      *string                  `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetThreadDataResponseBodyDataMessages) String() string {
	return dara.Prettify(s)
}

func (s GetThreadDataResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *GetThreadDataResponseBodyDataMessages) GetAgents() []map[string]interface{} {
	return s.Agents
}

func (s *GetThreadDataResponseBodyDataMessages) GetArtifacts() []map[string]interface{} {
	return s.Artifacts
}

func (s *GetThreadDataResponseBodyDataMessages) GetCallId() *string {
	return s.CallId
}

func (s *GetThreadDataResponseBodyDataMessages) GetContents() []map[string]interface{} {
	return s.Contents
}

func (s *GetThreadDataResponseBodyDataMessages) GetDetail() *string {
	return s.Detail
}

func (s *GetThreadDataResponseBodyDataMessages) GetEvents() []map[string]interface{} {
	return s.Events
}

func (s *GetThreadDataResponseBodyDataMessages) GetParentCallId() *string {
	return s.ParentCallId
}

func (s *GetThreadDataResponseBodyDataMessages) GetRole() *string {
	return s.Role
}

func (s *GetThreadDataResponseBodyDataMessages) GetSeq() *int32 {
	return s.Seq
}

func (s *GetThreadDataResponseBodyDataMessages) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetThreadDataResponseBodyDataMessages) GetTools() []map[string]interface{} {
	return s.Tools
}

func (s *GetThreadDataResponseBodyDataMessages) GetType() *string {
	return s.Type
}

func (s *GetThreadDataResponseBodyDataMessages) GetVersion() *string {
	return s.Version
}

func (s *GetThreadDataResponseBodyDataMessages) SetAgents(v []map[string]interface{}) *GetThreadDataResponseBodyDataMessages {
	s.Agents = v
	return s
}

func (s *GetThreadDataResponseBodyDataMessages) SetArtifacts(v []map[string]interface{}) *GetThreadDataResponseBodyDataMessages {
	s.Artifacts = v
	return s
}

func (s *GetThreadDataResponseBodyDataMessages) SetCallId(v string) *GetThreadDataResponseBodyDataMessages {
	s.CallId = &v
	return s
}

func (s *GetThreadDataResponseBodyDataMessages) SetContents(v []map[string]interface{}) *GetThreadDataResponseBodyDataMessages {
	s.Contents = v
	return s
}

func (s *GetThreadDataResponseBodyDataMessages) SetDetail(v string) *GetThreadDataResponseBodyDataMessages {
	s.Detail = &v
	return s
}

func (s *GetThreadDataResponseBodyDataMessages) SetEvents(v []map[string]interface{}) *GetThreadDataResponseBodyDataMessages {
	s.Events = v
	return s
}

func (s *GetThreadDataResponseBodyDataMessages) SetParentCallId(v string) *GetThreadDataResponseBodyDataMessages {
	s.ParentCallId = &v
	return s
}

func (s *GetThreadDataResponseBodyDataMessages) SetRole(v string) *GetThreadDataResponseBodyDataMessages {
	s.Role = &v
	return s
}

func (s *GetThreadDataResponseBodyDataMessages) SetSeq(v int32) *GetThreadDataResponseBodyDataMessages {
	s.Seq = &v
	return s
}

func (s *GetThreadDataResponseBodyDataMessages) SetTimestamp(v string) *GetThreadDataResponseBodyDataMessages {
	s.Timestamp = &v
	return s
}

func (s *GetThreadDataResponseBodyDataMessages) SetTools(v []map[string]interface{}) *GetThreadDataResponseBodyDataMessages {
	s.Tools = v
	return s
}

func (s *GetThreadDataResponseBodyDataMessages) SetType(v string) *GetThreadDataResponseBodyDataMessages {
	s.Type = &v
	return s
}

func (s *GetThreadDataResponseBodyDataMessages) SetVersion(v string) *GetThreadDataResponseBodyDataMessages {
	s.Version = &v
	return s
}

func (s *GetThreadDataResponseBodyDataMessages) Validate() error {
	return dara.Validate(s)
}
