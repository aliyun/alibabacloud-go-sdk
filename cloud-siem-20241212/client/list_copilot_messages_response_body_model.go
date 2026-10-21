// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCopilotMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiVersion(v string) *ListCopilotMessagesResponseBody
	GetApiVersion() *string
	SetItems(v []*ListCopilotMessagesResponseBodyItems) *ListCopilotMessagesResponseBody
	GetItems() []*ListCopilotMessagesResponseBodyItems
	SetMaxResults(v int32) *ListCopilotMessagesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListCopilotMessagesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCopilotMessagesResponseBody
	GetRequestId() *string
}

type ListCopilotMessagesResponseBody struct {
	// The API version.
	//
	// >You can call the [DescribeGroupProductions](~~DescribeGroupProductions~~) operation to obtain this parameter.
	//
	// example:
	//
	// v1.2
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// The list of conversation messages.
	Items []*ListCopilotMessagesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum amount of data read in this query.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCopilotMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCopilotMessagesResponseBody) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *ListCopilotMessagesResponseBody) GetItems() []*ListCopilotMessagesResponseBodyItems {
	return s.Items
}

func (s *ListCopilotMessagesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCopilotMessagesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCopilotMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCopilotMessagesResponseBody) SetApiVersion(v string) *ListCopilotMessagesResponseBody {
	s.ApiVersion = &v
	return s
}

func (s *ListCopilotMessagesResponseBody) SetItems(v []*ListCopilotMessagesResponseBodyItems) *ListCopilotMessagesResponseBody {
	s.Items = v
	return s
}

func (s *ListCopilotMessagesResponseBody) SetMaxResults(v int32) *ListCopilotMessagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCopilotMessagesResponseBody) SetNextToken(v string) *ListCopilotMessagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCopilotMessagesResponseBody) SetRequestId(v string) *ListCopilotMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCopilotMessagesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCopilotMessagesResponseBodyItems struct {
	// The client message ID.
	//
	// example:
	//
	// 366932fb-9a83-4a2d-94dd-c925b35f****
	ClientMessageId *string `json:"ClientMessageId,omitempty" xml:"ClientMessageId,omitempty"`
	// The sent content.
	//
	// example:
	//
	// The correlation analysis of alerts and network logs has been completed
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-01-11T02:01:36Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The message ID.
	//
	// example:
	//
	// AB917534-C484-5BA0-9077-0D5D1D2F****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The message sequence number.
	//
	// example:
	//
	// 1
	MessageSequence *int64 `json:"MessageSequence,omitempty" xml:"MessageSequence,omitempty"`
	// The message role.
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The status.
	//
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The primary key ID.
	//
	// example:
	//
	// dc18b894-19a3-45ee-8768-a59478e0****
	TurnId *string `json:"TurnId,omitempty" xml:"TurnId,omitempty"`
}

func (s ListCopilotMessagesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotMessagesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListCopilotMessagesResponseBodyItems) GetClientMessageId() *string {
	return s.ClientMessageId
}

func (s *ListCopilotMessagesResponseBodyItems) GetContent() *string {
	return s.Content
}

func (s *ListCopilotMessagesResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListCopilotMessagesResponseBodyItems) GetMessageId() *string {
	return s.MessageId
}

func (s *ListCopilotMessagesResponseBodyItems) GetMessageSequence() *int64 {
	return s.MessageSequence
}

func (s *ListCopilotMessagesResponseBodyItems) GetRole() *string {
	return s.Role
}

func (s *ListCopilotMessagesResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListCopilotMessagesResponseBodyItems) GetTurnId() *string {
	return s.TurnId
}

func (s *ListCopilotMessagesResponseBodyItems) SetClientMessageId(v string) *ListCopilotMessagesResponseBodyItems {
	s.ClientMessageId = &v
	return s
}

func (s *ListCopilotMessagesResponseBodyItems) SetContent(v string) *ListCopilotMessagesResponseBodyItems {
	s.Content = &v
	return s
}

func (s *ListCopilotMessagesResponseBodyItems) SetCreatedAt(v string) *ListCopilotMessagesResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListCopilotMessagesResponseBodyItems) SetMessageId(v string) *ListCopilotMessagesResponseBodyItems {
	s.MessageId = &v
	return s
}

func (s *ListCopilotMessagesResponseBodyItems) SetMessageSequence(v int64) *ListCopilotMessagesResponseBodyItems {
	s.MessageSequence = &v
	return s
}

func (s *ListCopilotMessagesResponseBodyItems) SetRole(v string) *ListCopilotMessagesResponseBodyItems {
	s.Role = &v
	return s
}

func (s *ListCopilotMessagesResponseBodyItems) SetStatus(v string) *ListCopilotMessagesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListCopilotMessagesResponseBodyItems) SetTurnId(v string) *ListCopilotMessagesResponseBodyItems {
	s.TurnId = &v
	return s
}

func (s *ListCopilotMessagesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
