// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConversationDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConversationDetails(v []*ListConversationDetailsResponseBodyConversationDetails) *ListConversationDetailsResponseBody
	GetConversationDetails() []*ListConversationDetailsResponseBodyConversationDetails
	SetRequestId(v string) *ListConversationDetailsResponseBody
	GetRequestId() *string
}

type ListConversationDetailsResponseBody struct {
	ConversationDetails []*ListConversationDetailsResponseBodyConversationDetails `json:"ConversationDetails,omitempty" xml:"ConversationDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConversationDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConversationDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConversationDetailsResponseBody) GetConversationDetails() []*ListConversationDetailsResponseBodyConversationDetails {
	return s.ConversationDetails
}

func (s *ListConversationDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConversationDetailsResponseBody) SetConversationDetails(v []*ListConversationDetailsResponseBodyConversationDetails) *ListConversationDetailsResponseBody {
	s.ConversationDetails = v
	return s
}

func (s *ListConversationDetailsResponseBody) SetRequestId(v string) *ListConversationDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConversationDetailsResponseBody) Validate() error {
	if s.ConversationDetails != nil {
		for _, item := range s.ConversationDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConversationDetailsResponseBodyConversationDetails struct {
	// example:
	//
	// Dialogue
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// {}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 1582266750353
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	SequenceId *string `json:"SequenceId,omitempty" xml:"SequenceId,omitempty"`
	// example:
	//
	// Chatbot
	Speaker   *string `json:"Speaker,omitempty" xml:"Speaker,omitempty"`
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s ListConversationDetailsResponseBodyConversationDetails) String() string {
	return dara.Prettify(s)
}

func (s ListConversationDetailsResponseBodyConversationDetails) GoString() string {
	return s.String()
}

func (s *ListConversationDetailsResponseBodyConversationDetails) GetAction() *string {
	return s.Action
}

func (s *ListConversationDetailsResponseBodyConversationDetails) GetActionParams() *string {
	return s.ActionParams
}

func (s *ListConversationDetailsResponseBodyConversationDetails) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListConversationDetailsResponseBodyConversationDetails) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListConversationDetailsResponseBodyConversationDetails) GetSequenceId() *string {
	return s.SequenceId
}

func (s *ListConversationDetailsResponseBodyConversationDetails) GetSpeaker() *string {
	return s.Speaker
}

func (s *ListConversationDetailsResponseBodyConversationDetails) GetUtterance() *string {
	return s.Utterance
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetAction(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.Action = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetActionParams(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.ActionParams = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetConversationId(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.ConversationId = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetCreateTime(v int64) *ListConversationDetailsResponseBodyConversationDetails {
	s.CreateTime = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetSequenceId(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.SequenceId = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetSpeaker(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.Speaker = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetUtterance(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.Utterance = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) Validate() error {
	return dara.Validate(s)
}
