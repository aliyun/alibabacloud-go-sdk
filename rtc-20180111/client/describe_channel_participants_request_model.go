// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelParticipantsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeChannelParticipantsRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeChannelParticipantsRequest
	GetChannelId() *string
	SetOrder(v string) *DescribeChannelParticipantsRequest
	GetOrder() *string
	SetOwnerId(v int64) *DescribeChannelParticipantsRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeChannelParticipantsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeChannelParticipantsRequest
	GetPageSize() *int32
}

type DescribeChannelParticipantsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aec****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// asc
	Order   *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeChannelParticipantsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelParticipantsRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelParticipantsRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeChannelParticipantsRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeChannelParticipantsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeChannelParticipantsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeChannelParticipantsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeChannelParticipantsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeChannelParticipantsRequest) SetAppId(v string) *DescribeChannelParticipantsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelParticipantsRequest) SetChannelId(v string) *DescribeChannelParticipantsRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelParticipantsRequest) SetOrder(v string) *DescribeChannelParticipantsRequest {
	s.Order = &v
	return s
}

func (s *DescribeChannelParticipantsRequest) SetOwnerId(v int64) *DescribeChannelParticipantsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeChannelParticipantsRequest) SetPageNum(v int32) *DescribeChannelParticipantsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeChannelParticipantsRequest) SetPageSize(v int32) *DescribeChannelParticipantsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeChannelParticipantsRequest) Validate() error {
	return dara.Validate(s)
}
