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
	SetPath(v string) *GetKnowledgeRecallRequest
	GetPath() *string
	SetQuestion(v string) *GetKnowledgeRecallRequest
	GetQuestion() *string
	SetTags(v string) *GetKnowledgeRecallRequest
	GetTags() *string
	SetTopk(v int32) *GetKnowledgeRecallRequest
	GetTopk() *int32
	SetUser(v string) *GetKnowledgeRecallRequest
	GetUser() *string
}

type GetKnowledgeRecallRequest struct {
	// The ID of the ADB MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp19aaaaaa****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The file path prefix. Only files that match the specified path prefix are recalled.
	//
	// example:
	//
	// oss://bucketName/path/prefix/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The question for knowledge base recall.
	//
	// This parameter is required.
	//
	// example:
	//
	// What are the reports for the clothing category this month?
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// The list of tags in JSON format.
	//
	// example:
	//
	// {   "tag_key1": ["tag_key1_value1", "tag_key1_value2"],   "tag_key2": ["tag_key2_value"] }
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The top K associated files to recall.
	//
	// example:
	//
	// 5
	Topk *int32 `json:"Topk,omitempty" xml:"Topk,omitempty"`
	// The username. Only files that the specified user has permission to access are recalled.
	//
	// example:
	//
	// user_name1
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

func (s *GetKnowledgeRecallRequest) GetPath() *string {
	return s.Path
}

func (s *GetKnowledgeRecallRequest) GetQuestion() *string {
	return s.Question
}

func (s *GetKnowledgeRecallRequest) GetTags() *string {
	return s.Tags
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

func (s *GetKnowledgeRecallRequest) SetPath(v string) *GetKnowledgeRecallRequest {
	s.Path = &v
	return s
}

func (s *GetKnowledgeRecallRequest) SetQuestion(v string) *GetKnowledgeRecallRequest {
	s.Question = &v
	return s
}

func (s *GetKnowledgeRecallRequest) SetTags(v string) *GetKnowledgeRecallRequest {
	s.Tags = &v
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
