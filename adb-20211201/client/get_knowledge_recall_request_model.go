// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeRecallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetKnowledgeRecallRequest
	GetDBClusterId() *string
	SetQuestion(v string) *GetKnowledgeRecallRequest
	GetQuestion() *string
	SetTopk(v int32) *GetKnowledgeRecallRequest
	GetTopk() *int32
	SetUser(v string) *GetKnowledgeRecallRequest
	GetUser() *string
}

type GetKnowledgeRecallRequest struct {
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp19aaaaaa****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The question for knowledge base recall.
	//
	// This parameter is required.
	//
	// example:
	//
	// What are the reports for the clothing category this month?
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// The top K number of related files to recall.
	//
	// example:
	//
	// 5
	Topk *int32 `json:"Topk,omitempty" xml:"Topk,omitempty"`
	// The username. Only files that this user has permission to access are recalled.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s GetKnowledgeRecallRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeRecallRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeRecallRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetKnowledgeRecallRequest) GetQuestion() *string {
	return s.Question
}

func (s *GetKnowledgeRecallRequest) GetTopk() *int32 {
	return s.Topk
}

func (s *GetKnowledgeRecallRequest) GetUser() *string {
	return s.User
}

func (s *GetKnowledgeRecallRequest) SetDBClusterId(v string) *GetKnowledgeRecallRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetKnowledgeRecallRequest) SetQuestion(v string) *GetKnowledgeRecallRequest {
	s.Question = &v
	return s
}

func (s *GetKnowledgeRecallRequest) SetTopk(v int32) *GetKnowledgeRecallRequest {
	s.Topk = &v
	return s
}

func (s *GetKnowledgeRecallRequest) SetUser(v string) *GetKnowledgeRecallRequest {
	s.User = &v
	return s
}

func (s *GetKnowledgeRecallRequest) Validate() error {
	return dara.Validate(s)
}
