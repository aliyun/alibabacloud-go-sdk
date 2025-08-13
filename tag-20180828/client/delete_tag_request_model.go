// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *DeleteTagRequest
	GetKey() *string
	SetOwnerAccount(v string) *DeleteTagRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTagRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteTagRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteTagRequest
	GetResourceOwnerAccount() *string
	SetValue(v string) *DeleteTagRequest
	GetValue() *string
}

type DeleteTagRequest struct {
	// The tag key.
	//
	// If no tag value is associated with a tag key, you can specify the `Key` parameter without specifying the Value parameter to delete the tag key. Otherwise, you must specify both the `Key` and `Value` parameters to delete a preset tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// Environment
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  Only `cn-hangzhou` is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DeleteTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagRequest) GetKey() *string {
	return s.Key
}

func (s *DeleteTagRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTagRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTagRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTagRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTagRequest) GetValue() *string {
	return s.Value
}

func (s *DeleteTagRequest) SetKey(v string) *DeleteTagRequest {
	s.Key = &v
	return s
}

func (s *DeleteTagRequest) SetOwnerAccount(v string) *DeleteTagRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTagRequest) SetOwnerId(v int64) *DeleteTagRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTagRequest) SetRegionId(v string) *DeleteTagRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTagRequest) SetResourceOwnerAccount(v string) *DeleteTagRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTagRequest) SetValue(v string) *DeleteTagRequest {
	s.Value = &v
	return s
}

func (s *DeleteTagRequest) Validate() error {
	return dara.Validate(s)
}
