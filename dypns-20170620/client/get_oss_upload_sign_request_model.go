// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUploadSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientType(v string) *GetOssUploadSignRequest
	GetClientType() *string
	SetOwnerId(v int64) *GetOssUploadSignRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetOssUploadSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetOssUploadSignRequest
	GetResourceOwnerId() *int64
}

type GetOssUploadSignRequest struct {
	// This parameter is required.
	ClientType           *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetOssUploadSignRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadSignRequest) GoString() string {
	return s.String()
}

func (s *GetOssUploadSignRequest) GetClientType() *string {
	return s.ClientType
}

func (s *GetOssUploadSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetOssUploadSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetOssUploadSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetOssUploadSignRequest) SetClientType(v string) *GetOssUploadSignRequest {
	s.ClientType = &v
	return s
}

func (s *GetOssUploadSignRequest) SetOwnerId(v int64) *GetOssUploadSignRequest {
	s.OwnerId = &v
	return s
}

func (s *GetOssUploadSignRequest) SetResourceOwnerAccount(v string) *GetOssUploadSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetOssUploadSignRequest) SetResourceOwnerId(v int64) *GetOssUploadSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetOssUploadSignRequest) Validate() error {
	return dara.Validate(s)
}
