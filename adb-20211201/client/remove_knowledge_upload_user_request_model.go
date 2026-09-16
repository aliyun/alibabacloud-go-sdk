// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveKnowledgeUploadUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *RemoveKnowledgeUploadUserRequest
	GetDBClusterId() *string
	SetFileLocation(v string) *RemoveKnowledgeUploadUserRequest
	GetFileLocation() *string
	SetUsers(v string) *RemoveKnowledgeUploadUserRequest
	GetUsers() *string
}

type RemoveKnowledgeUploadUserRequest struct {
	// The ID of the ADB instance.
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
	// The JSON string of the array of authorized users to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["alice","bob"]
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s RemoveKnowledgeUploadUserRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveKnowledgeUploadUserRequest) GoString() string {
	return s.String()
}

func (s *RemoveKnowledgeUploadUserRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RemoveKnowledgeUploadUserRequest) GetFileLocation() *string {
	return s.FileLocation
}

func (s *RemoveKnowledgeUploadUserRequest) GetUsers() *string {
	return s.Users
}

func (s *RemoveKnowledgeUploadUserRequest) SetDBClusterId(v string) *RemoveKnowledgeUploadUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *RemoveKnowledgeUploadUserRequest) SetFileLocation(v string) *RemoveKnowledgeUploadUserRequest {
	s.FileLocation = &v
	return s
}

func (s *RemoveKnowledgeUploadUserRequest) SetUsers(v string) *RemoveKnowledgeUploadUserRequest {
	s.Users = &v
	return s
}

func (s *RemoveKnowledgeUploadUserRequest) Validate() error {
	return dara.Validate(s)
}
