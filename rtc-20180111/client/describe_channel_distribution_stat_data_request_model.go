// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelDistributionStatDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeChannelDistributionStatDataRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeChannelDistributionStatDataRequest
	GetChannelId() *string
	SetCreatedTs(v int64) *DescribeChannelDistributionStatDataRequest
	GetCreatedTs() *int64
	SetDestroyedTs(v int64) *DescribeChannelDistributionStatDataRequest
	GetDestroyedTs() *int64
	SetStatDim(v string) *DescribeChannelDistributionStatDataRequest
	GetStatDim() *string
}

type DescribeChannelDistributionStatDataRequest struct {
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
	// The timestamp when the channel was created. You can query data from the last 30 days. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// The timestamp when the channel was released. The value is a UNIX timestamp in seconds. If this parameter is empty, the current time is used.
	//
	// > If the specified channel release time is later than the actual release time, the data from the creation time to the actual release time is returned.
	//
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// The statistical dimension. Valid values:
	//
	// - **OS**: Statistics by operating system.
	//
	// - **SDK_VERSION**: Statistics by SDK version.
	//
	// This parameter is required.
	//
	// example:
	//
	// OS
	StatDim *string `json:"StatDim,omitempty" xml:"StatDim,omitempty"`
}

func (s DescribeChannelDistributionStatDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelDistributionStatDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeChannelDistributionStatDataRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeChannelDistributionStatDataRequest) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeChannelDistributionStatDataRequest) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeChannelDistributionStatDataRequest) GetStatDim() *string {
	return s.StatDim
}

func (s *DescribeChannelDistributionStatDataRequest) SetAppId(v string) *DescribeChannelDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelDistributionStatDataRequest) SetChannelId(v string) *DescribeChannelDistributionStatDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelDistributionStatDataRequest) SetCreatedTs(v int64) *DescribeChannelDistributionStatDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelDistributionStatDataRequest) SetDestroyedTs(v int64) *DescribeChannelDistributionStatDataRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelDistributionStatDataRequest) SetStatDim(v string) *DescribeChannelDistributionStatDataRequest {
	s.StatDim = &v
	return s
}

func (s *DescribeChannelDistributionStatDataRequest) Validate() error {
	return dara.Validate(s)
}
