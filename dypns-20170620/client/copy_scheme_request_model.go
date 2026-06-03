// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CopySchemeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CopySchemeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CopySchemeRequest
	GetResourceOwnerId() *int64
	SetSchemeId(v int64) *CopySchemeRequest
	GetSchemeId() *int64
}

type CopySchemeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
}

func (s CopySchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s CopySchemeRequest) GoString() string {
	return s.String()
}

func (s *CopySchemeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CopySchemeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CopySchemeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CopySchemeRequest) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *CopySchemeRequest) SetOwnerId(v int64) *CopySchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *CopySchemeRequest) SetResourceOwnerAccount(v string) *CopySchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CopySchemeRequest) SetResourceOwnerId(v int64) *CopySchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CopySchemeRequest) SetSchemeId(v int64) *CopySchemeRequest {
	s.SchemeId = &v
	return s
}

func (s *CopySchemeRequest) Validate() error {
	return dara.Validate(s)
}
