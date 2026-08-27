// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateConversationResponseBody
	GetCode() *string
	SetConversationId(v string) *CreateConversationResponseBody
	GetConversationId() *string
	SetCreatedAt(v string) *CreateConversationResponseBody
	GetCreatedAt() *string
	SetMessage(v string) *CreateConversationResponseBody
	GetMessage() *string
	SetMetadata(v map[string]interface{}) *CreateConversationResponseBody
	GetMetadata() map[string]interface{}
	SetRequestId(v string) *CreateConversationResponseBody
	GetRequestId() *string
	SetTitle(v string) *CreateConversationResponseBody
	GetTitle() *string
}

type CreateConversationResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3a59769c-e631-4b48-84f3-c9bf3a8ae940
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// The time when the share was created.
	//
	// example:
	//
	// 1784513941206
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The status code description.
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// A reserved field for extension use.
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5241B90-8FF4-565C-977A-0CE1842AED72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The appointment title.
	//
	// example:
	//
	// Incident RCA: alert-0885feb7-3d4b-4da5-90f0-0119dfbbf555:up0shc25tp0kueo0afeobvhk81
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConversationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateConversationResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *CreateConversationResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateConversationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateConversationResponseBody) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *CreateConversationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConversationResponseBody) GetTitle() *string {
	return s.Title
}

func (s *CreateConversationResponseBody) SetCode(v string) *CreateConversationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConversationResponseBody) SetConversationId(v string) *CreateConversationResponseBody {
	s.ConversationId = &v
	return s
}

func (s *CreateConversationResponseBody) SetCreatedAt(v string) *CreateConversationResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateConversationResponseBody) SetMessage(v string) *CreateConversationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConversationResponseBody) SetMetadata(v map[string]interface{}) *CreateConversationResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateConversationResponseBody) SetRequestId(v string) *CreateConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConversationResponseBody) SetTitle(v string) *CreateConversationResponseBody {
	s.Title = &v
	return s
}

func (s *CreateConversationResponseBody) Validate() error {
	return dara.Validate(s)
}
