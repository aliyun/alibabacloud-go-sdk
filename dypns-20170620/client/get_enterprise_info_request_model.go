// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnterpriseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *GetEnterpriseInfoRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *GetEnterpriseInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetEnterpriseInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetEnterpriseInfoRequest
	GetResourceOwnerId() *int64
}

type GetEnterpriseInfoRequest struct {
	// This parameter is required.
	EnterpriseId         *int64  `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetEnterpriseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseInfoRequest) GoString() string {
	return s.String()
}

func (s *GetEnterpriseInfoRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *GetEnterpriseInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetEnterpriseInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetEnterpriseInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetEnterpriseInfoRequest) SetEnterpriseId(v int64) *GetEnterpriseInfoRequest {
	s.EnterpriseId = &v
	return s
}

func (s *GetEnterpriseInfoRequest) SetOwnerId(v int64) *GetEnterpriseInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *GetEnterpriseInfoRequest) SetResourceOwnerAccount(v string) *GetEnterpriseInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetEnterpriseInfoRequest) SetResourceOwnerId(v int64) *GetEnterpriseInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetEnterpriseInfoRequest) Validate() error {
	return dara.Validate(s)
}
