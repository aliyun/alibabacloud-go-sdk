// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetThreadDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDigitalEmployeeName(v string) *GetThreadDataResponseBody
	GetDigitalEmployeeName() *string
	SetMaxResults(v int64) *GetThreadDataResponseBody
	GetMaxResults() *int64
	SetMessages(v []*GetThreadDataResponseBodyMessages) *GetThreadDataResponseBody
	GetMessages() []*GetThreadDataResponseBodyMessages
	SetNextToken(v string) *GetThreadDataResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetThreadDataResponseBody
	GetRequestId() *string
	SetThreadId(v string) *GetThreadDataResponseBody
	GetThreadId() *string
}

type GetThreadDataResponseBody struct {
	// example:
	//
	// test
	DigitalEmployeeName *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// example:
	//
	// 2
	MaxResults *int64                               `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	Messages   []*GetThreadDataResponseBodyMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
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

func (s *GetThreadDataResponseBody) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *GetThreadDataResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *GetThreadDataResponseBody) GetMessages() []*GetThreadDataResponseBodyMessages {
	return s.Messages
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

func (s *GetThreadDataResponseBody) SetDigitalEmployeeName(v string) *GetThreadDataResponseBody {
	s.DigitalEmployeeName = &v
	return s
}

func (s *GetThreadDataResponseBody) SetMaxResults(v int64) *GetThreadDataResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetThreadDataResponseBody) SetMessages(v []*GetThreadDataResponseBodyMessages) *GetThreadDataResponseBody {
	s.Messages = v
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

type GetThreadDataResponseBodyMessages struct {
	// example:
	//
	// 205190712643664705
	CallerUid *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	// example:
	//
	// test
	DigitalEmployeeName *string                  `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	Items               []map[string]interface{} `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 7F0000012B1B668BC3D59A7EF8A00063
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// example:
	//
	// 1560138499250147
	OwnerUid *string `json:"ownerUid,omitempty" xml:"ownerUid,omitempty"`
	// example:
	//
	// xxxx
	ParentMessageId *string `json:"parentMessageId,omitempty" xml:"parentMessageId,omitempty"`
	// example:
	//
	// cn-qingdao
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// jr-c2b000da0e41b543
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// 98958d65-6cdb-4f40-8f46-f5e49f13c860
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
	// example:
	//
	// 1765359068
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// example:
	//
	// 3b5287ba17572104610774286d0096
	TraceId   *string            `json:"traceId,omitempty" xml:"traceId,omitempty"`
	Variables map[string]*string `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s GetThreadDataResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s GetThreadDataResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *GetThreadDataResponseBodyMessages) GetCallerUid() *string {
	return s.CallerUid
}

func (s *GetThreadDataResponseBodyMessages) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *GetThreadDataResponseBodyMessages) GetItems() []map[string]interface{} {
	return s.Items
}

func (s *GetThreadDataResponseBodyMessages) GetMessageId() *string {
	return s.MessageId
}

func (s *GetThreadDataResponseBodyMessages) GetOwnerUid() *string {
	return s.OwnerUid
}

func (s *GetThreadDataResponseBodyMessages) GetParentMessageId() *string {
	return s.ParentMessageId
}

func (s *GetThreadDataResponseBodyMessages) GetRegion() *string {
	return s.Region
}

func (s *GetThreadDataResponseBodyMessages) GetRole() *string {
	return s.Role
}

func (s *GetThreadDataResponseBodyMessages) GetRunId() *string {
	return s.RunId
}

func (s *GetThreadDataResponseBodyMessages) GetThreadId() *string {
	return s.ThreadId
}

func (s *GetThreadDataResponseBodyMessages) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetThreadDataResponseBodyMessages) GetTraceId() *string {
	return s.TraceId
}

func (s *GetThreadDataResponseBodyMessages) GetVariables() map[string]*string {
	return s.Variables
}

func (s *GetThreadDataResponseBodyMessages) SetCallerUid(v string) *GetThreadDataResponseBodyMessages {
	s.CallerUid = &v
	return s
}

func (s *GetThreadDataResponseBodyMessages) SetDigitalEmployeeName(v string) *GetThreadDataResponseBodyMessages {
	s.DigitalEmployeeName = &v
	return s
}

func (s *GetThreadDataResponseBodyMessages) SetItems(v []map[string]interface{}) *GetThreadDataResponseBodyMessages {
	s.Items = v
	return s
}

func (s *GetThreadDataResponseBodyMessages) SetMessageId(v string) *GetThreadDataResponseBodyMessages {
	s.MessageId = &v
	return s
}

func (s *GetThreadDataResponseBodyMessages) SetOwnerUid(v string) *GetThreadDataResponseBodyMessages {
	s.OwnerUid = &v
	return s
}

func (s *GetThreadDataResponseBodyMessages) SetParentMessageId(v string) *GetThreadDataResponseBodyMessages {
	s.ParentMessageId = &v
	return s
}

func (s *GetThreadDataResponseBodyMessages) SetRegion(v string) *GetThreadDataResponseBodyMessages {
	s.Region = &v
	return s
}

func (s *GetThreadDataResponseBodyMessages) SetRole(v string) *GetThreadDataResponseBodyMessages {
	s.Role = &v
	return s
}

func (s *GetThreadDataResponseBodyMessages) SetRunId(v string) *GetThreadDataResponseBodyMessages {
	s.RunId = &v
	return s
}

func (s *GetThreadDataResponseBodyMessages) SetThreadId(v string) *GetThreadDataResponseBodyMessages {
	s.ThreadId = &v
	return s
}

func (s *GetThreadDataResponseBodyMessages) SetTimestamp(v string) *GetThreadDataResponseBodyMessages {
	s.Timestamp = &v
	return s
}

func (s *GetThreadDataResponseBodyMessages) SetTraceId(v string) *GetThreadDataResponseBodyMessages {
	s.TraceId = &v
	return s
}

func (s *GetThreadDataResponseBodyMessages) SetVariables(v map[string]*string) *GetThreadDataResponseBodyMessages {
	s.Variables = v
	return s
}

func (s *GetThreadDataResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
