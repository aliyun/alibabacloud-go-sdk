// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUploadForEnterpriseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetOssUploadForEnterpriseRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetOssUploadForEnterpriseRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetOssUploadForEnterpriseRequest
	GetResourceOwnerId() *int64
}

type GetOssUploadForEnterpriseRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetOssUploadForEnterpriseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadForEnterpriseRequest) GoString() string {
	return s.String()
}

func (s *GetOssUploadForEnterpriseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetOssUploadForEnterpriseRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetOssUploadForEnterpriseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetOssUploadForEnterpriseRequest) SetOwnerId(v int64) *GetOssUploadForEnterpriseRequest {
	s.OwnerId = &v
	return s
}

func (s *GetOssUploadForEnterpriseRequest) SetResourceOwnerAccount(v string) *GetOssUploadForEnterpriseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetOssUploadForEnterpriseRequest) SetResourceOwnerId(v int64) *GetOssUploadForEnterpriseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetOssUploadForEnterpriseRequest) Validate() error {
	return dara.Validate(s)
}
