// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndPointEventListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeEndPointEventListRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeEndPointEventListRequest
	GetChannelId() *string
	SetCreatedTs(v int64) *DescribeEndPointEventListRequest
	GetCreatedTs() *int64
	SetDestroyedTs(v int64) *DescribeEndPointEventListRequest
	GetDestroyedTs() *int64
	SetUserIdList(v string) *DescribeEndPointEventListRequest
	GetUserIdList() *string
}

type DescribeEndPointEventListRequest struct {
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
	// The start time of the event query time range. This time cannot be earlier than the actual creation time of the channel or 30 days ago. The time is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// The end time of the event query time range. This time cannot be later than the release time of the channel. The time is a UNIX timestamp in seconds. If this parameter is left empty, the current time is used.
	//
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// The list of user IDs. Separate multiple user IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// testuserid1,testuserid2
	UserIdList *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s DescribeEndPointEventListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndPointEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeEndPointEventListRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeEndPointEventListRequest) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeEndPointEventListRequest) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeEndPointEventListRequest) GetUserIdList() *string {
	return s.UserIdList
}

func (s *DescribeEndPointEventListRequest) SetAppId(v string) *DescribeEndPointEventListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeEndPointEventListRequest) SetChannelId(v string) *DescribeEndPointEventListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeEndPointEventListRequest) SetCreatedTs(v int64) *DescribeEndPointEventListRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeEndPointEventListRequest) SetDestroyedTs(v int64) *DescribeEndPointEventListRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeEndPointEventListRequest) SetUserIdList(v string) *DescribeEndPointEventListRequest {
	s.UserIdList = &v
	return s
}

func (s *DescribeEndPointEventListRequest) Validate() error {
	return dara.Validate(s)
}
