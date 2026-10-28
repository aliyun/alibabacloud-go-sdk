// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteKnowledgeFileRequest
	GetDBClusterId() *string
	SetFileId(v string) *DeleteKnowledgeFileRequest
	GetFileId() *string
	SetFileLocation(v string) *DeleteKnowledgeFileRequest
	GetFileLocation() *string
}

type DeleteKnowledgeFileRequest struct {
	// The ID of the ADB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp19aaaaaa****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the file.
	//
	// example:
	//
	// 2131213
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The path of the knowledge base document.
	//
	// example:
	//
	// oss://bucket/doc.pdf
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
}

func (s DeleteKnowledgeFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeFileRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteKnowledgeFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *DeleteKnowledgeFileRequest) GetFileLocation() *string {
	return s.FileLocation
}

func (s *DeleteKnowledgeFileRequest) SetDBClusterId(v string) *DeleteKnowledgeFileRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteKnowledgeFileRequest) SetFileId(v string) *DeleteKnowledgeFileRequest {
	s.FileId = &v
	return s
}

func (s *DeleteKnowledgeFileRequest) SetFileLocation(v string) *DeleteKnowledgeFileRequest {
	s.FileLocation = &v
	return s
}

func (s *DeleteKnowledgeFileRequest) Validate() error {
	return dara.Validate(s)
}
