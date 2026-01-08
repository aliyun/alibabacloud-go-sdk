// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGeneratePresignedUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilePath(v string) *GeneratePresignedUrlRequest
	GetFilePath() *string
	SetOwnerId(v int64) *GeneratePresignedUrlRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GeneratePresignedUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GeneratePresignedUrlRequest
	GetResourceOwnerId() *int64
}

type GeneratePresignedUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// /aaa/111/aa.png
	FilePath             *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GeneratePresignedUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GeneratePresignedUrlRequest) GoString() string {
	return s.String()
}

func (s *GeneratePresignedUrlRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *GeneratePresignedUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GeneratePresignedUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GeneratePresignedUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GeneratePresignedUrlRequest) SetFilePath(v string) *GeneratePresignedUrlRequest {
	s.FilePath = &v
	return s
}

func (s *GeneratePresignedUrlRequest) SetOwnerId(v int64) *GeneratePresignedUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GeneratePresignedUrlRequest) SetResourceOwnerAccount(v string) *GeneratePresignedUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GeneratePresignedUrlRequest) SetResourceOwnerId(v int64) *GeneratePresignedUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GeneratePresignedUrlRequest) Validate() error {
	return dara.Validate(s)
}
