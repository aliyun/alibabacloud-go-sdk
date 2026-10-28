// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveKnowledgeTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *RemoveKnowledgeTagsRequest
	GetDBClusterId() *string
	SetFileId(v string) *RemoveKnowledgeTagsRequest
	GetFileId() *string
	SetFileLocation(v string) *RemoveKnowledgeTagsRequest
	GetFileLocation() *string
	SetTags(v string) *RemoveKnowledgeTagsRequest
	GetTags() *string
}

type RemoveKnowledgeTagsRequest struct {
	// The database cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp19aaaaaa****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 2131234
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The knowledge base document location.
	//
	// example:
	//
	// oss://bucket/doc.pdf
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
	// The JSON string of the tag array to be deleted.
	//
	// example:
	//
	// [{"tag_key":"biz.scene","tag_value":"test"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s RemoveKnowledgeTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveKnowledgeTagsRequest) GoString() string {
	return s.String()
}

func (s *RemoveKnowledgeTagsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RemoveKnowledgeTagsRequest) GetFileId() *string {
	return s.FileId
}

func (s *RemoveKnowledgeTagsRequest) GetFileLocation() *string {
	return s.FileLocation
}

func (s *RemoveKnowledgeTagsRequest) GetTags() *string {
	return s.Tags
}

func (s *RemoveKnowledgeTagsRequest) SetDBClusterId(v string) *RemoveKnowledgeTagsRequest {
	s.DBClusterId = &v
	return s
}

func (s *RemoveKnowledgeTagsRequest) SetFileId(v string) *RemoveKnowledgeTagsRequest {
	s.FileId = &v
	return s
}

func (s *RemoveKnowledgeTagsRequest) SetFileLocation(v string) *RemoveKnowledgeTagsRequest {
	s.FileLocation = &v
	return s
}

func (s *RemoveKnowledgeTagsRequest) SetTags(v string) *RemoveKnowledgeTagsRequest {
	s.Tags = &v
	return s
}

func (s *RemoveKnowledgeTagsRequest) Validate() error {
	return dara.Validate(s)
}
