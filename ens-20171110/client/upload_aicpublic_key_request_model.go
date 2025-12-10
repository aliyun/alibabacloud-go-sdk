// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAICPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UploadAICPublicKeyRequest
	GetContent() *string
	SetDescription(v string) *UploadAICPublicKeyRequest
	GetDescription() *string
	SetKeyGroup(v string) *UploadAICPublicKeyRequest
	GetKeyGroup() *string
	SetKeyName(v string) *UploadAICPublicKeyRequest
	GetKeyName() *string
	SetKeyType(v string) *UploadAICPublicKeyRequest
	GetKeyType() *string
}

type UploadAICPublicKeyRequest struct {
	// Public Key
	//
	// This parameter is required.
	//
	// example:
	//
	// verify_dffeb6610035dcb77b413a59c3
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the document.
	//
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Public Key Grouping
	//
	// example:
	//
	// g-test
	KeyGroup *string `json:"KeyGroup,omitempty" xml:"KeyGroup,omitempty"`
	// Public Key Name
	//
	// This parameter is required.
	//
	// example:
	//
	// mykey
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// Public key type
	//
	// This parameter is required.
	//
	// example:
	//
	// adb
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
}

func (s UploadAICPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadAICPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *UploadAICPublicKeyRequest) GetContent() *string {
	return s.Content
}

func (s *UploadAICPublicKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *UploadAICPublicKeyRequest) GetKeyGroup() *string {
	return s.KeyGroup
}

func (s *UploadAICPublicKeyRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *UploadAICPublicKeyRequest) GetKeyType() *string {
	return s.KeyType
}

func (s *UploadAICPublicKeyRequest) SetContent(v string) *UploadAICPublicKeyRequest {
	s.Content = &v
	return s
}

func (s *UploadAICPublicKeyRequest) SetDescription(v string) *UploadAICPublicKeyRequest {
	s.Description = &v
	return s
}

func (s *UploadAICPublicKeyRequest) SetKeyGroup(v string) *UploadAICPublicKeyRequest {
	s.KeyGroup = &v
	return s
}

func (s *UploadAICPublicKeyRequest) SetKeyName(v string) *UploadAICPublicKeyRequest {
	s.KeyName = &v
	return s
}

func (s *UploadAICPublicKeyRequest) SetKeyType(v string) *UploadAICPublicKeyRequest {
	s.KeyType = &v
	return s
}

func (s *UploadAICPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
