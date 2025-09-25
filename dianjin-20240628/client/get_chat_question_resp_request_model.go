// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatQuestionRespRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *GetChatQuestionRespRequest
	GetBatchId() *string
	SetSessionId(v string) *GetChatQuestionRespRequest
	GetSessionId() *string
}

type GetChatQuestionRespRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1869307330227937280
	BatchId *string `json:"batchId,omitempty" xml:"batchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 237645726354
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s GetChatQuestionRespRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatQuestionRespRequest) GoString() string {
	return s.String()
}

func (s *GetChatQuestionRespRequest) GetBatchId() *string {
	return s.BatchId
}

func (s *GetChatQuestionRespRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetChatQuestionRespRequest) SetBatchId(v string) *GetChatQuestionRespRequest {
	s.BatchId = &v
	return s
}

func (s *GetChatQuestionRespRequest) SetSessionId(v string) *GetChatQuestionRespRequest {
	s.SessionId = &v
	return s
}

func (s *GetChatQuestionRespRequest) Validate() error {
	return dara.Validate(s)
}
