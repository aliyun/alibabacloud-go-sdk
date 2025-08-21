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
	// This parameter is required.
	//
	// example:
	//
	// verify_dffeb6610035dcb77b413a59c3*****
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// g-test
	KeyGroup *string `json:"KeyGroup,omitempty" xml:"KeyGroup,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mykey-v1.0
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
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
