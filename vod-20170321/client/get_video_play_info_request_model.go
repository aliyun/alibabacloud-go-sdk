// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoPlayInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *GetVideoPlayInfoRequest
	GetChannel() *string
	SetClientTS(v int64) *GetVideoPlayInfoRequest
	GetClientTS() *int64
	SetClientVersion(v string) *GetVideoPlayInfoRequest
	GetClientVersion() *string
	SetOwnerId(v int64) *GetVideoPlayInfoRequest
	GetOwnerId() *int64
	SetPlaySign(v string) *GetVideoPlayInfoRequest
	GetPlaySign() *string
	SetResourceOwnerAccount(v string) *GetVideoPlayInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetVideoPlayInfoRequest
	GetResourceOwnerId() *int64
	SetSignVersion(v string) *GetVideoPlayInfoRequest
	GetSignVersion() *string
	SetVideoId(v string) *GetVideoPlayInfoRequest
	GetVideoId() *string
}

type GetVideoPlayInfoRequest struct {
	// This parameter is required.
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// This parameter is required.
	ClientTS *int64 `json:"ClientTS,omitempty" xml:"ClientTS,omitempty"`
	// This parameter is required.
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PlaySign             *string `json:"PlaySign,omitempty" xml:"PlaySign,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SignVersion *string `json:"SignVersion,omitempty" xml:"SignVersion,omitempty"`
	// This parameter is required.
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoPlayInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPlayInfoRequest) GoString() string {
	return s.String()
}

func (s *GetVideoPlayInfoRequest) GetChannel() *string {
	return s.Channel
}

func (s *GetVideoPlayInfoRequest) GetClientTS() *int64 {
	return s.ClientTS
}

func (s *GetVideoPlayInfoRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *GetVideoPlayInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetVideoPlayInfoRequest) GetPlaySign() *string {
	return s.PlaySign
}

func (s *GetVideoPlayInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetVideoPlayInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetVideoPlayInfoRequest) GetSignVersion() *string {
	return s.SignVersion
}

func (s *GetVideoPlayInfoRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *GetVideoPlayInfoRequest) SetChannel(v string) *GetVideoPlayInfoRequest {
	s.Channel = &v
	return s
}

func (s *GetVideoPlayInfoRequest) SetClientTS(v int64) *GetVideoPlayInfoRequest {
	s.ClientTS = &v
	return s
}

func (s *GetVideoPlayInfoRequest) SetClientVersion(v string) *GetVideoPlayInfoRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetVideoPlayInfoRequest) SetOwnerId(v int64) *GetVideoPlayInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *GetVideoPlayInfoRequest) SetPlaySign(v string) *GetVideoPlayInfoRequest {
	s.PlaySign = &v
	return s
}

func (s *GetVideoPlayInfoRequest) SetResourceOwnerAccount(v string) *GetVideoPlayInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetVideoPlayInfoRequest) SetResourceOwnerId(v int64) *GetVideoPlayInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetVideoPlayInfoRequest) SetSignVersion(v string) *GetVideoPlayInfoRequest {
	s.SignVersion = &v
	return s
}

func (s *GetVideoPlayInfoRequest) SetVideoId(v string) *GetVideoPlayInfoRequest {
	s.VideoId = &v
	return s
}

func (s *GetVideoPlayInfoRequest) Validate() error {
	return dara.Validate(s)
}
