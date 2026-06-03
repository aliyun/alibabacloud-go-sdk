// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteSchemeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteSchemeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSchemeRequest
	GetResourceOwnerId() *int64
	SetSchemeId(v int64) *DeleteSchemeRequest
	GetSchemeId() *int64
}

type DeleteSchemeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
}

func (s DeleteSchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchemeRequest) GoString() string {
	return s.String()
}

func (s *DeleteSchemeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSchemeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSchemeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSchemeRequest) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *DeleteSchemeRequest) SetOwnerId(v int64) *DeleteSchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSchemeRequest) SetResourceOwnerAccount(v string) *DeleteSchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSchemeRequest) SetResourceOwnerId(v int64) *DeleteSchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSchemeRequest) SetSchemeId(v int64) *DeleteSchemeRequest {
	s.SchemeId = &v
	return s
}

func (s *DeleteSchemeRequest) Validate() error {
	return dara.Validate(s)
}
