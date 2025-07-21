// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteTagRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTagRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTagRequest
	GetResourceOwnerId() *int64
	SetTagId(v int32) *DeleteTagRequest
	GetTagId() *int32
}

type DeleteTagRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the tag
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	TagId *int32 `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s DeleteTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTagRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTagRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTagRequest) GetTagId() *int32 {
	return s.TagId
}

func (s *DeleteTagRequest) SetOwnerId(v int64) *DeleteTagRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTagRequest) SetResourceOwnerAccount(v string) *DeleteTagRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTagRequest) SetResourceOwnerId(v int64) *DeleteTagRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTagRequest) SetTagId(v int32) *DeleteTagRequest {
	s.TagId = &v
	return s
}

func (s *DeleteTagRequest) Validate() error {
	return dara.Validate(s)
}
