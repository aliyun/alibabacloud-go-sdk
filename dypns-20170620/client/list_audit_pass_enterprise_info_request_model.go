// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuditPassEnterpriseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListAuditPassEnterpriseInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListAuditPassEnterpriseInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListAuditPassEnterpriseInfoRequest
	GetResourceOwnerId() *int64
}

type ListAuditPassEnterpriseInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAuditPassEnterpriseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuditPassEnterpriseInfoRequest) GoString() string {
	return s.String()
}

func (s *ListAuditPassEnterpriseInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAuditPassEnterpriseInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAuditPassEnterpriseInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAuditPassEnterpriseInfoRequest) SetOwnerId(v int64) *ListAuditPassEnterpriseInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAuditPassEnterpriseInfoRequest) SetResourceOwnerAccount(v string) *ListAuditPassEnterpriseInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAuditPassEnterpriseInfoRequest) SetResourceOwnerId(v int64) *ListAuditPassEnterpriseInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAuditPassEnterpriseInfoRequest) Validate() error {
	return dara.Validate(s)
}
