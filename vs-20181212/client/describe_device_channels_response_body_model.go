// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceChannelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannels(v []*DescribeDeviceChannelsResponseBodyChannels) *DescribeDeviceChannelsResponseBody
	GetChannels() []*DescribeDeviceChannelsResponseBodyChannels
	SetPageCount(v int64) *DescribeDeviceChannelsResponseBody
	GetPageCount() *int64
	SetPageNum(v int64) *DescribeDeviceChannelsResponseBody
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeDeviceChannelsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDeviceChannelsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDeviceChannelsResponseBody
	GetTotalCount() *int64
}

type DescribeDeviceChannelsResponseBody struct {
	Channels []*DescribeDeviceChannelsResponseBodyChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 4641C72D-462E-4AEA-8485-FC267AF90B0A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeviceChannelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceChannelsResponseBody) GetChannels() []*DescribeDeviceChannelsResponseBodyChannels {
	return s.Channels
}

func (s *DescribeDeviceChannelsResponseBody) GetPageCount() *int64 {
	return s.PageCount
}

func (s *DescribeDeviceChannelsResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeDeviceChannelsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDeviceChannelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeviceChannelsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDeviceChannelsResponseBody) SetChannels(v []*DescribeDeviceChannelsResponseBodyChannels) *DescribeDeviceChannelsResponseBody {
	s.Channels = v
	return s
}

func (s *DescribeDeviceChannelsResponseBody) SetPageCount(v int64) *DescribeDeviceChannelsResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBody) SetPageNum(v int64) *DescribeDeviceChannelsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBody) SetPageSize(v int64) *DescribeDeviceChannelsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBody) SetRequestId(v string) *DescribeDeviceChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBody) SetTotalCount(v int64) *DescribeDeviceChannelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBody) Validate() error {
	if s.Channels != nil {
		for _, item := range s.Channels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeviceChannelsResponseBodyChannels struct {
	// example:
	//
	// 0
	ChannelId *int64 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 348*****380-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// on
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// example:
	//
	// 3100000****000000002
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// example:
	//
	// 323*****997-cn-qingdao
	StreamId *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
	// example:
	//
	// off
	StreamStatus *string `json:"StreamStatus,omitempty" xml:"StreamStatus,omitempty"`
}

func (s DescribeDeviceChannelsResponseBodyChannels) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceChannelsResponseBodyChannels) GoString() string {
	return s.String()
}

func (s *DescribeDeviceChannelsResponseBodyChannels) GetChannelId() *int64 {
	return s.ChannelId
}

func (s *DescribeDeviceChannelsResponseBodyChannels) GetDeviceId() *string {
	return s.DeviceId
}

func (s *DescribeDeviceChannelsResponseBodyChannels) GetDeviceStatus() *string {
	return s.DeviceStatus
}

func (s *DescribeDeviceChannelsResponseBodyChannels) GetGbId() *string {
	return s.GbId
}

func (s *DescribeDeviceChannelsResponseBodyChannels) GetName() *string {
	return s.Name
}

func (s *DescribeDeviceChannelsResponseBodyChannels) GetParams() *string {
	return s.Params
}

func (s *DescribeDeviceChannelsResponseBodyChannels) GetStreamId() *string {
	return s.StreamId
}

func (s *DescribeDeviceChannelsResponseBodyChannels) GetStreamStatus() *string {
	return s.StreamStatus
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetChannelId(v int64) *DescribeDeviceChannelsResponseBodyChannels {
	s.ChannelId = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetDeviceId(v string) *DescribeDeviceChannelsResponseBodyChannels {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetDeviceStatus(v string) *DescribeDeviceChannelsResponseBodyChannels {
	s.DeviceStatus = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetGbId(v string) *DescribeDeviceChannelsResponseBodyChannels {
	s.GbId = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetName(v string) *DescribeDeviceChannelsResponseBodyChannels {
	s.Name = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetParams(v string) *DescribeDeviceChannelsResponseBodyChannels {
	s.Params = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetStreamId(v string) *DescribeDeviceChannelsResponseBodyChannels {
	s.StreamId = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetStreamStatus(v string) *DescribeDeviceChannelsResponseBodyChannels {
	s.StreamStatus = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) Validate() error {
	return dara.Validate(s)
}
