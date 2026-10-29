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
	// The encoded content of the public key. The default encoding format is Base64.
	//
	// This parameter is required.
	//
	// example:
	//
	// verify_dffeb6610035dcb77b413a59c3*****
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the public key.
	//
	// example:
	//
	// For testing purposes
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The group of the public key. Used for public key management.
	//
	// Naming conventions:
	//
	// 1. The length must be 0 to 255 characters.
	//
	// 2. The name can contain lowercase letters, digits, underscores (_), hyphens (-), and periods (.).
	//
	// 3. The first character must be a letter or a digit.
	//
	// example:
	//
	// g-test
	KeyGroup *string `json:"KeyGroup,omitempty" xml:"KeyGroup,omitempty"`
	// The name of the public key. The name must be unique.
	//
	// Naming conventions:
	//
	// 1. The length must be 8 to 255 characters.
	//
	// 2. The name can contain lowercase letters, digits, underscores (_), hyphens (-), and periods (.).
	//
	// 3. The first character must be a letter or a digit.
	//
	// 4. The prefix cannot be "group-".
	//
	// This parameter is required.
	//
	// example:
	//
	// mykey-v1.0
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The type of the public key. Valid values:
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
