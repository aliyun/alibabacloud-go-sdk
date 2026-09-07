// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListKnowledgeTagsRequest
	GetDBClusterId() *string
	SetFileLocation(v string) *ListKnowledgeTagsRequest
	GetFileLocation() *string
}

type ListKnowledgeTagsRequest struct {
	// The ID of the AnalyticDB for MySQL instance.
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
	// oss://bucket/doc.pdf
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
}

func (s ListKnowledgeTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeTagsRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeTagsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListKnowledgeTagsRequest) GetFileLocation() *string {
	return s.FileLocation
}

func (s *ListKnowledgeTagsRequest) SetDBClusterId(v string) *ListKnowledgeTagsRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListKnowledgeTagsRequest) SetFileLocation(v string) *ListKnowledgeTagsRequest {
	s.FileLocation = &v
	return s
}

func (s *ListKnowledgeTagsRequest) Validate() error {
	return dara.Validate(s)
}
