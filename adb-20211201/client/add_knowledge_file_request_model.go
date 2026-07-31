// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKnowledgeFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *AddKnowledgeFileRequest
	GetDBClusterId() *string
	SetFileLocation(v string) *AddKnowledgeFileRequest
	GetFileLocation() *string
	SetFileType(v string) *AddKnowledgeFileRequest
	GetFileType() *string
	SetIsDir(v bool) *AddKnowledgeFileRequest
	GetIsDir() *bool
	SetTags(v string) *AddKnowledgeFileRequest
	GetTags() *string
	SetUploadUser(v string) *AddKnowledgeFileRequest
	GetUploadUser() *string
}

type AddKnowledgeFileRequest struct {
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp19aaaaaa****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The file address. Currently, only OSS paths are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket_name/file/path
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
	// The file type.
	//
	// example:
	//
	// pdf
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// Specifies whether the file is a folder.
	//
	// example:
	//
	// false
	IsDir *bool `json:"IsDir,omitempty" xml:"IsDir,omitempty"`
	// The file tags in JSON format.
	//
	// example:
	//
	// {"type":"game"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The user who uploads the knowledge base file.
	//
	// example:
	//
	// user1
	UploadUser *string `json:"UploadUser,omitempty" xml:"UploadUser,omitempty"`
}

func (s AddKnowledgeFileRequest) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeFileRequest) GoString() string {
	return s.String()
}

func (s *AddKnowledgeFileRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AddKnowledgeFileRequest) GetFileLocation() *string {
	return s.FileLocation
}

func (s *AddKnowledgeFileRequest) GetFileType() *string {
	return s.FileType
}

func (s *AddKnowledgeFileRequest) GetIsDir() *bool {
	return s.IsDir
}

func (s *AddKnowledgeFileRequest) GetTags() *string {
	return s.Tags
}

func (s *AddKnowledgeFileRequest) GetUploadUser() *string {
	return s.UploadUser
}

func (s *AddKnowledgeFileRequest) SetDBClusterId(v string) *AddKnowledgeFileRequest {
	s.DBClusterId = &v
	return s
}

func (s *AddKnowledgeFileRequest) SetFileLocation(v string) *AddKnowledgeFileRequest {
	s.FileLocation = &v
	return s
}

func (s *AddKnowledgeFileRequest) SetFileType(v string) *AddKnowledgeFileRequest {
	s.FileType = &v
	return s
}

func (s *AddKnowledgeFileRequest) SetIsDir(v bool) *AddKnowledgeFileRequest {
	s.IsDir = &v
	return s
}

func (s *AddKnowledgeFileRequest) SetTags(v string) *AddKnowledgeFileRequest {
	s.Tags = &v
	return s
}

func (s *AddKnowledgeFileRequest) SetUploadUser(v string) *AddKnowledgeFileRequest {
	s.UploadUser = &v
	return s
}

func (s *AddKnowledgeFileRequest) Validate() error {
	return dara.Validate(s)
}
