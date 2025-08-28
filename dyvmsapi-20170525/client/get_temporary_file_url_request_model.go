// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemporaryFileUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetTemporaryFileUrlRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetTemporaryFileUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetTemporaryFileUrlRequest
	GetResourceOwnerId() *int64
	SetVideoId(v string) *GetTemporaryFileUrlRequest
	GetVideoId() *string
}

type GetTemporaryFileUrlRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 45a04670582571eebf9e4531948c****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetTemporaryFileUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemporaryFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetTemporaryFileUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetTemporaryFileUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetTemporaryFileUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetTemporaryFileUrlRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *GetTemporaryFileUrlRequest) SetOwnerId(v int64) *GetTemporaryFileUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetTemporaryFileUrlRequest) SetResourceOwnerAccount(v string) *GetTemporaryFileUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetTemporaryFileUrlRequest) SetResourceOwnerId(v int64) *GetTemporaryFileUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetTemporaryFileUrlRequest) SetVideoId(v string) *GetTemporaryFileUrlRequest {
	s.VideoId = &v
	return s
}

func (s *GetTemporaryFileUrlRequest) Validate() error {
	return dara.Validate(s)
}
