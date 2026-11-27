// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePubUserListBySubUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribePubUserListBySubUserRequest
	GetAppId() *string
	SetChannelId(v string) *DescribePubUserListBySubUserRequest
	GetChannelId() *string
	SetCreatedTs(v int64) *DescribePubUserListBySubUserRequest
	GetCreatedTs() *int64
	SetDestroyedTs(v int64) *DescribePubUserListBySubUserRequest
	GetDestroyedTs() *int64
	SetSubUserId(v string) *DescribePubUserListBySubUserRequest
	GetSubUserId() *string
}

type DescribePubUserListBySubUserRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Channel ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The channel creation time. You can query data from the last 30 days. The value is represented as a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// The channel release time, represented as a UNIX timestamp in seconds. If this parameter is empty, the current time is used.
	//
	// > If the provided channel release time exceeds the actual release time, data from the creation time to the actual release time will be returned.
	//
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// Subscriber User ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// testuserid
	SubUserId *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribePubUserListBySubUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePubUserListBySubUserRequest) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribePubUserListBySubUserRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribePubUserListBySubUserRequest) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribePubUserListBySubUserRequest) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribePubUserListBySubUserRequest) GetSubUserId() *string {
	return s.SubUserId
}

func (s *DescribePubUserListBySubUserRequest) SetAppId(v string) *DescribePubUserListBySubUserRequest {
	s.AppId = &v
	return s
}

func (s *DescribePubUserListBySubUserRequest) SetChannelId(v string) *DescribePubUserListBySubUserRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribePubUserListBySubUserRequest) SetCreatedTs(v int64) *DescribePubUserListBySubUserRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserRequest) SetDestroyedTs(v int64) *DescribePubUserListBySubUserRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserRequest) SetSubUserId(v string) *DescribePubUserListBySubUserRequest {
	s.SubUserId = &v
	return s
}

func (s *DescribePubUserListBySubUserRequest) Validate() error {
	return dara.Validate(s)
}
