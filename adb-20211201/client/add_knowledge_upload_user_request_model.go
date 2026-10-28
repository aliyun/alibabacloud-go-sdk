// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKnowledgeUploadUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *AddKnowledgeUploadUserRequest
	GetDBClusterId() *string
	SetFileId(v string) *AddKnowledgeUploadUserRequest
	GetFileId() *string
	SetFileLocation(v string) *AddKnowledgeUploadUserRequest
	GetFileLocation() *string
	SetUsers(v string) *AddKnowledgeUploadUserRequest
	GetUsers() *string
}

type AddKnowledgeUploadUserRequest struct {
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
	// 3214123
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The path of the knowledge base document.
	//
	// example:
	//
	// oss://bucketName/path/to/file.pdf
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
	// The JSON string of the authorized users array.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["alice","bob"]
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s AddKnowledgeUploadUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeUploadUserRequest) GoString() string {
	return s.String()
}

func (s *AddKnowledgeUploadUserRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AddKnowledgeUploadUserRequest) GetFileId() *string {
	return s.FileId
}

func (s *AddKnowledgeUploadUserRequest) GetFileLocation() *string {
	return s.FileLocation
}

func (s *AddKnowledgeUploadUserRequest) GetUsers() *string {
	return s.Users
}

func (s *AddKnowledgeUploadUserRequest) SetDBClusterId(v string) *AddKnowledgeUploadUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *AddKnowledgeUploadUserRequest) SetFileId(v string) *AddKnowledgeUploadUserRequest {
	s.FileId = &v
	return s
}

func (s *AddKnowledgeUploadUserRequest) SetFileLocation(v string) *AddKnowledgeUploadUserRequest {
	s.FileLocation = &v
	return s
}

func (s *AddKnowledgeUploadUserRequest) SetUsers(v string) *AddKnowledgeUploadUserRequest {
	s.Users = &v
	return s
}

func (s *AddKnowledgeUploadUserRequest) Validate() error {
	return dara.Validate(s)
}
