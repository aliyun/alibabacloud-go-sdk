// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteUserPublicKeyRequest
	GetInstanceId() *string
	SetPublicKeyId(v string) *DeleteUserPublicKeyRequest
	GetPublicKeyId() *string
	SetRegionId(v string) *DeleteUserPublicKeyRequest
	GetRegionId() *string
}

type DeleteUserPublicKeyRequest struct {
	// The ID of the Bastionhost instance to which the users to be queried belong.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the Bastionhost instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The public key ID.
	//
	// >  You can call the [ListUserPublicKeys](https://help.aliyun.com/document_detail/477555.html) operation to query the public key ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	PublicKeyId *string `json:"PublicKeyId,omitempty" xml:"PublicKeyId,omitempty"`
	// The region ID of the bastion host. For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteUserPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserPublicKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteUserPublicKeyRequest) GetPublicKeyId() *string {
	return s.PublicKeyId
}

func (s *DeleteUserPublicKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteUserPublicKeyRequest) SetInstanceId(v string) *DeleteUserPublicKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUserPublicKeyRequest) SetPublicKeyId(v string) *DeleteUserPublicKeyRequest {
	s.PublicKeyId = &v
	return s
}

func (s *DeleteUserPublicKeyRequest) SetRegionId(v string) *DeleteUserPublicKeyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUserPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
