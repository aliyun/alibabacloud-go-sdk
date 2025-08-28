// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayVideoFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *PlayVideoFileRequest
	GetCallId() *string
	SetCalledNumber(v string) *PlayVideoFileRequest
	GetCalledNumber() *string
	SetOnlyPhone(v bool) *PlayVideoFileRequest
	GetOnlyPhone() *bool
	SetOutId(v string) *PlayVideoFileRequest
	GetOutId() *string
	SetOwnerId(v int64) *PlayVideoFileRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *PlayVideoFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PlayVideoFileRequest
	GetResourceOwnerId() *int64
	SetVideoId(v string) *PlayVideoFileRequest
	GetVideoId() *string
}

type PlayVideoFileRequest struct {
	// example:
	//
	// 116012354148^1028137841****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// 示例值
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// example:
	//
	// false
	OnlyPhone *bool `json:"OnlyPhone,omitempty" xml:"OnlyPhone,omitempty"`
	// example:
	//
	// 342268*****
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s PlayVideoFileRequest) String() string {
	return dara.Prettify(s)
}

func (s PlayVideoFileRequest) GoString() string {
	return s.String()
}

func (s *PlayVideoFileRequest) GetCallId() *string {
	return s.CallId
}

func (s *PlayVideoFileRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *PlayVideoFileRequest) GetOnlyPhone() *bool {
	return s.OnlyPhone
}

func (s *PlayVideoFileRequest) GetOutId() *string {
	return s.OutId
}

func (s *PlayVideoFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PlayVideoFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PlayVideoFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PlayVideoFileRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *PlayVideoFileRequest) SetCallId(v string) *PlayVideoFileRequest {
	s.CallId = &v
	return s
}

func (s *PlayVideoFileRequest) SetCalledNumber(v string) *PlayVideoFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *PlayVideoFileRequest) SetOnlyPhone(v bool) *PlayVideoFileRequest {
	s.OnlyPhone = &v
	return s
}

func (s *PlayVideoFileRequest) SetOutId(v string) *PlayVideoFileRequest {
	s.OutId = &v
	return s
}

func (s *PlayVideoFileRequest) SetOwnerId(v int64) *PlayVideoFileRequest {
	s.OwnerId = &v
	return s
}

func (s *PlayVideoFileRequest) SetResourceOwnerAccount(v string) *PlayVideoFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PlayVideoFileRequest) SetResourceOwnerId(v int64) *PlayVideoFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PlayVideoFileRequest) SetVideoId(v string) *PlayVideoFileRequest {
	s.VideoId = &v
	return s
}

func (s *PlayVideoFileRequest) Validate() error {
	return dara.Validate(s)
}
