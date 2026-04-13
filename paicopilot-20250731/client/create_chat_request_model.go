// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtraData(v string) *CreateChatRequest
	GetExtraData() *string
	SetPayload(v map[string]interface{}) *CreateChatRequest
	GetPayload() map[string]interface{}
	SetQuestion(v *ChatDetail) *CreateChatRequest
	GetQuestion() *ChatDetail
	SetTitle(v string) *CreateChatRequest
	GetTitle() *string
}

type CreateChatRequest struct {
	// example:
	//
	// {"solutionDetail":{"formValues":{"params":{"InstanceId":"dsw-g54******qg9"},"content":{"EcsSpec":"ecs.gn6i-c8g1.2xlarge"}},"success":true}}
	ExtraData *string                `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	Payload   map[string]interface{} `json:"Payload,omitempty" xml:"Payload,omitempty"`
	Question  *ChatDetail            `json:"Question,omitempty" xml:"Question,omitempty"`
	Title     *string                `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateChatRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatRequest) GoString() string {
	return s.String()
}

func (s *CreateChatRequest) GetExtraData() *string {
	return s.ExtraData
}

func (s *CreateChatRequest) GetPayload() map[string]interface{} {
	return s.Payload
}

func (s *CreateChatRequest) GetQuestion() *ChatDetail {
	return s.Question
}

func (s *CreateChatRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateChatRequest) SetExtraData(v string) *CreateChatRequest {
	s.ExtraData = &v
	return s
}

func (s *CreateChatRequest) SetPayload(v map[string]interface{}) *CreateChatRequest {
	s.Payload = v
	return s
}

func (s *CreateChatRequest) SetQuestion(v *ChatDetail) *CreateChatRequest {
	s.Question = v
	return s
}

func (s *CreateChatRequest) SetTitle(v string) *CreateChatRequest {
	s.Title = &v
	return s
}

func (s *CreateChatRequest) Validate() error {
	if s.Question != nil {
		if err := s.Question.Validate(); err != nil {
			return err
		}
	}
	return nil
}
