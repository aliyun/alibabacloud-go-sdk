// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UploadPublicKeyRequest
	GetContent() *string
	SetDescription(v string) *UploadPublicKeyRequest
	GetDescription() *string
	SetKeyGroup(v string) *UploadPublicKeyRequest
	GetKeyGroup() *string
	SetKeyName(v string) *UploadPublicKeyRequest
	GetKeyName() *string
	SetKeyType(v string) *UploadPublicKeyRequest
	GetKeyType() *string
}

type UploadPublicKeyRequest struct {
	// Base64-encoded public key content.
	//
	// This parameter is required.
	//
	// example:
	//
	// verify_dffeb6610035dcb77b413a59c3*****
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Description of the public key.
	//
	// example:
	//
	// 测试使用
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Group for the public key. Used for public key management.
	//
	// 1. Length: 0 to 255 characters.
	//
	// 2. Valid characters: lowercase letters, digits, underscores (_), hyphens (-), and periods (.).
	//
	// 3. First character must be a letter or digit.
	//
	// example:
	//
	// g-test
	KeyGroup *string `json:"KeyGroup,omitempty" xml:"KeyGroup,omitempty"`
	// Name of the public key. Must be unique.
	//
	// 1. Length: 8 to 255 characters.
	//
	// 2. Valid characters: lowercase letters, digits, underscores (_), hyphens (-), and periods (.).
	//
	// 3. First character must be a letter or digit.
	//
	// 4. Prefix cannot be group-.
	//
	// This parameter is required.
	//
	// example:
	//
	// mykey-v1.0
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// Type of the public key. Valid values:
	//
	// - **adb**: ADB key.
	//
	// - **ssh**: SSH key.
	//
	// example:
	//
	// ssh
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
}

func (s UploadPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *UploadPublicKeyRequest) GetContent() *string {
	return s.Content
}

func (s *UploadPublicKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *UploadPublicKeyRequest) GetKeyGroup() *string {
	return s.KeyGroup
}

func (s *UploadPublicKeyRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *UploadPublicKeyRequest) GetKeyType() *string {
	return s.KeyType
}

func (s *UploadPublicKeyRequest) SetContent(v string) *UploadPublicKeyRequest {
	s.Content = &v
	return s
}

func (s *UploadPublicKeyRequest) SetDescription(v string) *UploadPublicKeyRequest {
	s.Description = &v
	return s
}

func (s *UploadPublicKeyRequest) SetKeyGroup(v string) *UploadPublicKeyRequest {
	s.KeyGroup = &v
	return s
}

func (s *UploadPublicKeyRequest) SetKeyName(v string) *UploadPublicKeyRequest {
	s.KeyName = &v
	return s
}

func (s *UploadPublicKeyRequest) SetKeyType(v string) *UploadPublicKeyRequest {
	s.KeyType = &v
	return s
}

func (s *UploadPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
