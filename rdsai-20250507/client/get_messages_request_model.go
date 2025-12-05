// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *GetMessagesRequest
	GetConversationId() *string
	SetFirstId(v string) *GetMessagesRequest
	GetFirstId() *string
	SetLimit(v int64) *GetMessagesRequest
	GetLimit() *int64
}

type GetMessagesRequest struct {
	// example:
	//
	// 941c6f59-acf5-4e11-9adc-31e52e1f****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 038866af-a050-4bc5-bfad-b7bfc838****
	FirstId *string `json:"FirstId,omitempty" xml:"FirstId,omitempty"`
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s GetMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMessagesRequest) GoString() string {
	return s.String()
}

func (s *GetMessagesRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetMessagesRequest) GetFirstId() *string {
	return s.FirstId
}

func (s *GetMessagesRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *GetMessagesRequest) SetConversationId(v string) *GetMessagesRequest {
	s.ConversationId = &v
	return s
}

func (s *GetMessagesRequest) SetFirstId(v string) *GetMessagesRequest {
	s.FirstId = &v
	return s
}

func (s *GetMessagesRequest) SetLimit(v int64) *GetMessagesRequest {
	s.Limit = &v
	return s
}

func (s *GetMessagesRequest) Validate() error {
	return dara.Validate(s)
}
