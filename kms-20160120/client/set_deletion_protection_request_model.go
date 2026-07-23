// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeletionProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletionProtectionDescription(v string) *SetDeletionProtectionRequest
	GetDeletionProtectionDescription() *string
	SetEnableDeletionProtection(v bool) *SetDeletionProtectionRequest
	GetEnableDeletionProtection() *bool
	SetKeyId(v string) *SetDeletionProtectionRequest
	GetKeyId() *string
	SetKmsInstanceId(v string) *SetDeletionProtectionRequest
	GetKmsInstanceId() *string
	SetProtectedResourceArn(v string) *SetDeletionProtectionRequest
	GetProtectedResourceArn() *string
}

type SetDeletionProtectionRequest struct {
	// The description of deletion protection.
	//
	// > This parameter is valid only when EnableDeletionProtection is set to true.
	//
	// example:
	//
	// The CMK is being used by XXX. Deletion protection is set.
	DeletionProtectionDescription *string `json:"DeletionProtectionDescription,omitempty" xml:"DeletionProtectionDescription,omitempty"`
	// Specifies whether to enable deletion protection. Valid values:
	//
	// - true: enables deletion protection.
	//
	// - false (default): disables deletion protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitempty" xml:"EnableDeletionProtection,omitempty"`
	// The ID of the key.
	//
	// example:
	//
	// key-hzz65f3a68554s6ms****
	KeyId         *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KmsInstanceId *string `json:"KmsInstanceId,omitempty" xml:"KmsInstanceId,omitempty"`
	// The ARN of the CMK for which you want to set deletion protection.
	//
	// You can call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the CMK ARN (Arn).
	//
	// example:
	//
	// acs:kms:cn-hangzhou:123213123****:key/0225f411-b21d-46d1-be5b-93931c82****
	ProtectedResourceArn *string `json:"ProtectedResourceArn,omitempty" xml:"ProtectedResourceArn,omitempty"`
}

func (s SetDeletionProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetDeletionProtectionRequest) GetDeletionProtectionDescription() *string {
	return s.DeletionProtectionDescription
}

func (s *SetDeletionProtectionRequest) GetEnableDeletionProtection() *bool {
	return s.EnableDeletionProtection
}

func (s *SetDeletionProtectionRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *SetDeletionProtectionRequest) GetKmsInstanceId() *string {
	return s.KmsInstanceId
}

func (s *SetDeletionProtectionRequest) GetProtectedResourceArn() *string {
	return s.ProtectedResourceArn
}

func (s *SetDeletionProtectionRequest) SetDeletionProtectionDescription(v string) *SetDeletionProtectionRequest {
	s.DeletionProtectionDescription = &v
	return s
}

func (s *SetDeletionProtectionRequest) SetEnableDeletionProtection(v bool) *SetDeletionProtectionRequest {
	s.EnableDeletionProtection = &v
	return s
}

func (s *SetDeletionProtectionRequest) SetKeyId(v string) *SetDeletionProtectionRequest {
	s.KeyId = &v
	return s
}

func (s *SetDeletionProtectionRequest) SetKmsInstanceId(v string) *SetDeletionProtectionRequest {
	s.KmsInstanceId = &v
	return s
}

func (s *SetDeletionProtectionRequest) SetProtectedResourceArn(v string) *SetDeletionProtectionRequest {
	s.ProtectedResourceArn = &v
	return s
}

func (s *SetDeletionProtectionRequest) Validate() error {
	return dara.Validate(s)
}
