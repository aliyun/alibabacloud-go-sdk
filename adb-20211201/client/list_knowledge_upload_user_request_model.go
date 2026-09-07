// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeUploadUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListKnowledgeUploadUserRequest
	GetDBClusterId() *string
	SetFileLocation(v string) *ListKnowledgeUploadUserRequest
	GetFileLocation() *string
}

type ListKnowledgeUploadUserRequest struct {
	// The ADB instance ID.
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

func (s ListKnowledgeUploadUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeUploadUserRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeUploadUserRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListKnowledgeUploadUserRequest) GetFileLocation() *string {
	return s.FileLocation
}

func (s *ListKnowledgeUploadUserRequest) SetDBClusterId(v string) *ListKnowledgeUploadUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListKnowledgeUploadUserRequest) SetFileLocation(v string) *ListKnowledgeUploadUserRequest {
	s.FileLocation = &v
	return s
}

func (s *ListKnowledgeUploadUserRequest) Validate() error {
	return dara.Validate(s)
}
