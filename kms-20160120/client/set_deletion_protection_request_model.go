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
	SetProtectedResourceArn(v string) *SetDeletionProtectionRequest
	GetProtectedResourceArn() *string
}

type SetDeletionProtectionRequest struct {
	// The description of deletion protection.
	//
	// >  This parameter takes effect only when you set the EnableDeletionProtection parameter to true.
	//
	// example:
	//
	// This key is being used by XXX service. You are protected from deletion.
	DeletionProtectionDescription *string `json:"DeletionProtectionDescription,omitempty" xml:"DeletionProtectionDescription,omitempty"`
	// Specifies whether to enable deletion protection. Valid values:
	//
	// 	- true: enables deletion protection.
	//
	// 	- false: disables deletion protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableDeletionProtection *bool   `json:"EnableDeletionProtection,omitempty" xml:"EnableDeletionProtection,omitempty"`
	KeyId                    *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ARN of the CMK for which you want to set deletion protection.
	//
	// You can call the [DescribeKey](https://help.aliyun.com/document_detail/28952.html) operation to query the CMK ARN.
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

func (s *SetDeletionProtectionRequest) SetProtectedResourceArn(v string) *SetDeletionProtectionRequest {
	s.ProtectedResourceArn = &v
	return s
}

func (s *SetDeletionProtectionRequest) Validate() error {
	return dara.Validate(s)
}
