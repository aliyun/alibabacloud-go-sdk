// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoFieldUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetVideoFieldUrlRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetVideoFieldUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetVideoFieldUrlRequest
	GetResourceOwnerId() *int64
	SetVideoFile(v string) *GetVideoFieldUrlRequest
	GetVideoFile() *string
}

type GetVideoFieldUrlRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值示例值
	VideoFile *string `json:"VideoFile,omitempty" xml:"VideoFile,omitempty"`
}

func (s GetVideoFieldUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoFieldUrlRequest) GoString() string {
	return s.String()
}

func (s *GetVideoFieldUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetVideoFieldUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetVideoFieldUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetVideoFieldUrlRequest) GetVideoFile() *string {
	return s.VideoFile
}

func (s *GetVideoFieldUrlRequest) SetOwnerId(v int64) *GetVideoFieldUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetVideoFieldUrlRequest) SetResourceOwnerAccount(v string) *GetVideoFieldUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetVideoFieldUrlRequest) SetResourceOwnerId(v int64) *GetVideoFieldUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetVideoFieldUrlRequest) SetVideoFile(v string) *GetVideoFieldUrlRequest {
	s.VideoFile = &v
	return s
}

func (s *GetVideoFieldUrlRequest) Validate() error {
	return dara.Validate(s)
}
