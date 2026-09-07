// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKnowledgeTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *AddKnowledgeTagsRequest
	GetDBClusterId() *string
	SetFileLocation(v string) *AddKnowledgeTagsRequest
	GetFileLocation() *string
	SetTags(v string) *AddKnowledgeTagsRequest
	GetTags() *string
}

type AddKnowledgeTagsRequest struct {
	// The database cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp19aaaaaa****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The location of the knowledge base document.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucketName/path/to/file.pdf
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
	// The JSON string of the tag array. Each tag must be in the following format: {"tag_key":"key_name","tag_value":"value"}.
	//
	// example:
	//
	// [{"tag_key":"biz.scene","tag_value":"test"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AddKnowledgeTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeTagsRequest) GoString() string {
	return s.String()
}

func (s *AddKnowledgeTagsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AddKnowledgeTagsRequest) GetFileLocation() *string {
	return s.FileLocation
}

func (s *AddKnowledgeTagsRequest) GetTags() *string {
	return s.Tags
}

func (s *AddKnowledgeTagsRequest) SetDBClusterId(v string) *AddKnowledgeTagsRequest {
	s.DBClusterId = &v
	return s
}

func (s *AddKnowledgeTagsRequest) SetFileLocation(v string) *AddKnowledgeTagsRequest {
	s.FileLocation = &v
	return s
}

func (s *AddKnowledgeTagsRequest) SetTags(v string) *AddKnowledgeTagsRequest {
	s.Tags = &v
	return s
}

func (s *AddKnowledgeTagsRequest) Validate() error {
	return dara.Validate(s)
}
