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
	SetImages(v string) *GetKnowledgeRecallRequest
	GetImages() *string
	SetPath(v string) *GetKnowledgeRecallRequest
	GetPath() *string
	SetQuestion(v string) *GetKnowledgeRecallRequest
	GetQuestion() *string
	SetSemanticRules(v string) *GetKnowledgeRecallRequest
	GetSemanticRules() *string
	SetTags(v string) *GetKnowledgeRecallRequest
	GetTags() *string
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
	// The array of images to search for image retrieval. Currently, only one image is supported. The image must be Base64-encoded and passed in JSON format. Set the base64 key to the encoded image and the mime_type key to the image format. For more information, see the examples.
	//
	// example:
	//
	// [{"base64": "/9j/4AAQSkZJRgABAQ...", "mime_type": "image/jpeg"}]
	Images *string `json:"Images,omitempty" xml:"Images,omitempty"`
	// The file path prefix. Only files that match the path prefix are retrieved.
	//
	// example:
	//
	// oss://bucketName/path/prefix/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The question for knowledge base retrieval.
	//
	// This parameter is required.
	//
	// example:
	//
	// What are the reports for the clothing category this month?
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// The semantic rules. Currently, only "all" and "" are supported.
	//
	// example:
	//
	// all
	SemanticRules *string `json:"SemanticRules,omitempty" xml:"SemanticRules,omitempty"`
	// The list of tags in JSON format.
	//
	// example:
	//
	// {   "tag_key1": ["tag_key1_value1", "tag_key1_value2"],   "tag_key2": ["tag_key2_value"] }
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The top K value for retrieving associated files.
	//
	// example:
	//
	// 5
	Topk *int32 `json:"Topk,omitempty" xml:"Topk,omitempty"`
	// The username. Only files that the user has permissions to access are retrieved.
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

func (s *GetKnowledgeRecallRequest) GetImages() *string {
	return s.Images
}

func (s *GetKnowledgeRecallRequest) GetPath() *string {
	return s.Path
}

func (s *GetKnowledgeRecallRequest) GetQuestion() *string {
	return s.Question
}

func (s *GetKnowledgeRecallRequest) GetSemanticRules() *string {
	return s.SemanticRules
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

func (s *GetKnowledgeRecallRequest) SetImages(v string) *GetKnowledgeRecallRequest {
	s.Images = &v
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

func (s *GetKnowledgeRecallRequest) SetSemanticRules(v string) *GetKnowledgeRecallRequest {
	s.SemanticRules = &v
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
