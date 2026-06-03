// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetSchemeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetSchemeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetSchemeRequest
	GetResourceOwnerId() *int64
	SetSchemeId(v int64) *GetSchemeRequest
	GetSchemeId() *int64
}

type GetSchemeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
}

func (s GetSchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSchemeRequest) GoString() string {
	return s.String()
}

func (s *GetSchemeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetSchemeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetSchemeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetSchemeRequest) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *GetSchemeRequest) SetOwnerId(v int64) *GetSchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSchemeRequest) SetResourceOwnerAccount(v string) *GetSchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSchemeRequest) SetResourceOwnerId(v int64) *GetSchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSchemeRequest) SetSchemeId(v int64) *GetSchemeRequest {
	s.SchemeId = &v
	return s
}

func (s *GetSchemeRequest) Validate() error {
	return dara.Validate(s)
}
