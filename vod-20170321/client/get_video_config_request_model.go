// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthInfo(v string) *GetVideoConfigRequest
	GetAuthInfo() *string
	SetOwnerId(v int64) *GetVideoConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetVideoConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetVideoConfigRequest
	GetResourceOwnerId() *int64
	SetVideoId(v string) *GetVideoConfigRequest
	GetVideoId() *string
}

type GetVideoConfigRequest struct {
	AuthInfo             *string `json:"AuthInfo,omitempty" xml:"AuthInfo,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoConfigRequest) GoString() string {
	return s.String()
}

func (s *GetVideoConfigRequest) GetAuthInfo() *string {
	return s.AuthInfo
}

func (s *GetVideoConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetVideoConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetVideoConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetVideoConfigRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *GetVideoConfigRequest) SetAuthInfo(v string) *GetVideoConfigRequest {
	s.AuthInfo = &v
	return s
}

func (s *GetVideoConfigRequest) SetOwnerId(v int64) *GetVideoConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *GetVideoConfigRequest) SetResourceOwnerAccount(v string) *GetVideoConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetVideoConfigRequest) SetResourceOwnerId(v int64) *GetVideoConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetVideoConfigRequest) SetVideoId(v string) *GetVideoConfigRequest {
	s.VideoId = &v
	return s
}

func (s *GetVideoConfigRequest) Validate() error {
	return dara.Validate(s)
}
