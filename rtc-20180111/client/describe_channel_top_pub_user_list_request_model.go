// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelTopPubUserListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeChannelTopPubUserListRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeChannelTopPubUserListRequest
	GetChannelId() *string
	SetCreatedTs(v int64) *DescribeChannelTopPubUserListRequest
	GetCreatedTs() *int64
	SetDestroyedTs(v int64) *DescribeChannelTopPubUserListRequest
	GetDestroyedTs() *int64
}

type DescribeChannelTopPubUserListRequest struct {
	// The app ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The channel ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The timestamp when the channel was created. Data from the last 30 days can be queried. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1615893133
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// The timestamp when the channel was released. The value is a UNIX timestamp in seconds. If this parameter is left empty, the current time is used.
	//
	// > If the specified channel release time is later than the actual release time, the data from the creation time to the actual release time is returned.
	//
	// example:
	//
	// 1615893757
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
}

func (s DescribeChannelTopPubUserListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelTopPubUserListRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeChannelTopPubUserListRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeChannelTopPubUserListRequest) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeChannelTopPubUserListRequest) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeChannelTopPubUserListRequest) SetAppId(v string) *DescribeChannelTopPubUserListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelTopPubUserListRequest) SetChannelId(v string) *DescribeChannelTopPubUserListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelTopPubUserListRequest) SetCreatedTs(v int64) *DescribeChannelTopPubUserListRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelTopPubUserListRequest) SetDestroyedTs(v int64) *DescribeChannelTopPubUserListRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelTopPubUserListRequest) Validate() error {
	return dara.Validate(s)
}
