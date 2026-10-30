// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForkKnowledgeFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ForkKnowledgeFileRequest
	GetDBClusterId() *string
	SetFileId(v string) *ForkKnowledgeFileRequest
	GetFileId() *string
	SetFileLocation(v string) *ForkKnowledgeFileRequest
	GetFileLocation() *string
	SetTags(v string) *ForkKnowledgeFileRequest
	GetTags() *string
	SetUploadUser(v string) *ForkKnowledgeFileRequest
	GetUploadUser() *string
}

type ForkKnowledgeFileRequest struct {
	// The ADB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp19aaaaaa****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the knowledge base file. You must specify either this parameter or FileLocation, but not both. Use a decimal positive integer string without leading zeros. Valid values: 1 to 9223372036854775807.
	//
	// example:
	//
	// 123
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The address of a single file in the knowledge base. You must specify either this parameter or FileId, but not both. You cannot specify a directory. The address must uniquely locate an active file.
	//
	// example:
	//
	// oss://bucket/doc.pdf
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
	// The initial business tags for the new record. Use a JSON string of a JSON object or an array of tag_key/tag_value objects. If omitted, the source tags are not inherited.
	//
	// example:
	//
	// {"biz.game":"yys"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The initial authorized user for the new record. You can pass a JSON string of a single user identifier or a user array. If omitted, no initial user authorization is applied, and the source authorization is not inherited.
	//
	// example:
	//
	// ["alice","bob"]
	UploadUser *string `json:"UploadUser,omitempty" xml:"UploadUser,omitempty"`
}

func (s ForkKnowledgeFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ForkKnowledgeFileRequest) GoString() string {
	return s.String()
}

func (s *ForkKnowledgeFileRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ForkKnowledgeFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *ForkKnowledgeFileRequest) GetFileLocation() *string {
	return s.FileLocation
}

func (s *ForkKnowledgeFileRequest) GetTags() *string {
	return s.Tags
}

func (s *ForkKnowledgeFileRequest) GetUploadUser() *string {
	return s.UploadUser
}

func (s *ForkKnowledgeFileRequest) SetDBClusterId(v string) *ForkKnowledgeFileRequest {
	s.DBClusterId = &v
	return s
}

func (s *ForkKnowledgeFileRequest) SetFileId(v string) *ForkKnowledgeFileRequest {
	s.FileId = &v
	return s
}

func (s *ForkKnowledgeFileRequest) SetFileLocation(v string) *ForkKnowledgeFileRequest {
	s.FileLocation = &v
	return s
}

func (s *ForkKnowledgeFileRequest) SetTags(v string) *ForkKnowledgeFileRequest {
	s.Tags = &v
	return s
}

func (s *ForkKnowledgeFileRequest) SetUploadUser(v string) *ForkKnowledgeFileRequest {
	s.UploadUser = &v
	return s
}

func (s *ForkKnowledgeFileRequest) Validate() error {
	return dara.Validate(s)
}
