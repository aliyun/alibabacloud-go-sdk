// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ModifyUserPublicKeyRequest
	GetComment() *string
	SetInstanceId(v string) *ModifyUserPublicKeyRequest
	GetInstanceId() *string
	SetPublicKey(v string) *ModifyUserPublicKeyRequest
	GetPublicKey() *string
	SetPublicKeyId(v string) *ModifyUserPublicKeyRequest
	GetPublicKeyId() *string
	SetPublicKeyName(v string) *ModifyUserPublicKeyRequest
	GetPublicKeyName() *string
	SetRegionId(v string) *ModifyUserPublicKeyRequest
	GetRegionId() *string
}

type ModifyUserPublicKeyRequest struct {
	// The new description of the user group. The description can be up to 500 characters in length.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the bastion host on which you want to modify the public key of a user.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-nif236pmc1u
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new public key.
	//
	// >  Specify a Base64-encoded string.
	//
	// example:
	//
	// c3NoLWVkMjU1MTkgQUFBQUMzTnphQzFsWkRJMU5URTVBQUFBSUhVcjY4UENFYWFzZjFYRVpNYTVsMlNBQytHV3FpeXVsRVpndkV4dmlPM28gcm9vdEA5NjBkMmNhOTcwYjU=
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The ID of the public key that you want to modify.
	//
	// >  You can call the [ListUserPublicKeys](https://help.aliyun.com/document_detail/477555.html) operation to query the public key ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	PublicKeyId *string `json:"PublicKeyId,omitempty" xml:"PublicKeyId,omitempty"`
	// The name of the public key that you want to modify. This name can be up to 128 characters in length.
	//
	// example:
	//
	// name
	PublicKeyName *string `json:"PublicKeyName,omitempty" xml:"PublicKeyName,omitempty"`
	// The region ID of the bastion host that is used to modify the public key of the user.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyUserPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserPublicKeyRequest) GetComment() *string {
	return s.Comment
}

func (s *ModifyUserPublicKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyUserPublicKeyRequest) GetPublicKey() *string {
	return s.PublicKey
}

func (s *ModifyUserPublicKeyRequest) GetPublicKeyId() *string {
	return s.PublicKeyId
}

func (s *ModifyUserPublicKeyRequest) GetPublicKeyName() *string {
	return s.PublicKeyName
}

func (s *ModifyUserPublicKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyUserPublicKeyRequest) SetComment(v string) *ModifyUserPublicKeyRequest {
	s.Comment = &v
	return s
}

func (s *ModifyUserPublicKeyRequest) SetInstanceId(v string) *ModifyUserPublicKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserPublicKeyRequest) SetPublicKey(v string) *ModifyUserPublicKeyRequest {
	s.PublicKey = &v
	return s
}

func (s *ModifyUserPublicKeyRequest) SetPublicKeyId(v string) *ModifyUserPublicKeyRequest {
	s.PublicKeyId = &v
	return s
}

func (s *ModifyUserPublicKeyRequest) SetPublicKeyName(v string) *ModifyUserPublicKeyRequest {
	s.PublicKeyName = &v
	return s
}

func (s *ModifyUserPublicKeyRequest) SetRegionId(v string) *ModifyUserPublicKeyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
