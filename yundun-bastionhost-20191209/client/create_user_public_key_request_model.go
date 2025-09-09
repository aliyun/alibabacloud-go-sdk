// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateUserPublicKeyRequest
	GetComment() *string
	SetInstanceId(v string) *CreateUserPublicKeyRequest
	GetInstanceId() *string
	SetPublicKey(v string) *CreateUserPublicKeyRequest
	GetPublicKey() *string
	SetPublicKeyName(v string) *CreateUserPublicKeyRequest
	GetPublicKeyName() *string
	SetRegionId(v string) *CreateUserPublicKeyRequest
	GetRegionId() *string
	SetUserId(v string) *CreateUserPublicKeyRequest
	GetUserId() *string
}

type CreateUserPublicKeyRequest struct {
	// The description of the public key. The description can be up to 500 characters in length.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the bastion host on which you want to create a public key for the user.
	//
	// > You can call the [listinstances](https://help.aliyun.com/document_detail/204522.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The public key. Encode the value by using the Base64 algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// c3NoLWVkMjU1MTkgQUFBQUMzTnphQzFsWkRJMU5URTVBQUFBSUxGQnQxUUpyT3IxK2hTTGRkbERMZUx4WGRIZ3hBalBxWHJIbWNFNWxqSk8gbm93Y29kZXJAbm93Y29kZXJkZU1hY0Jvb2stUHJvLmxvY2Fs
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The name of the public key.
	//
	// This parameter is required.
	//
	// example:
	//
	// Public key of a user
	PublicKeyName *string `json:"PublicKeyName,omitempty" xml:"PublicKeyName,omitempty"`
	// Specifies the region ID of the bastion host on which you want to create a public key for the user.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user for whom you want to create a public key.
	//
	// >  You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateUserPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateUserPublicKeyRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateUserPublicKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUserPublicKeyRequest) GetPublicKey() *string {
	return s.PublicKey
}

func (s *CreateUserPublicKeyRequest) GetPublicKeyName() *string {
	return s.PublicKeyName
}

func (s *CreateUserPublicKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateUserPublicKeyRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateUserPublicKeyRequest) SetComment(v string) *CreateUserPublicKeyRequest {
	s.Comment = &v
	return s
}

func (s *CreateUserPublicKeyRequest) SetInstanceId(v string) *CreateUserPublicKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserPublicKeyRequest) SetPublicKey(v string) *CreateUserPublicKeyRequest {
	s.PublicKey = &v
	return s
}

func (s *CreateUserPublicKeyRequest) SetPublicKeyName(v string) *CreateUserPublicKeyRequest {
	s.PublicKeyName = &v
	return s
}

func (s *CreateUserPublicKeyRequest) SetRegionId(v string) *CreateUserPublicKeyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateUserPublicKeyRequest) SetUserId(v string) *CreateUserPublicKeyRequest {
	s.UserId = &v
	return s
}

func (s *CreateUserPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
