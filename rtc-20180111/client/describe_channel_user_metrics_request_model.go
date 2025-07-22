// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelUserMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeChannelUserMetricsRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeChannelUserMetricsRequest
	GetChannelId() *string
	SetCreatedTs(v int64) *DescribeChannelUserMetricsRequest
	GetCreatedTs() *int64
	SetDestroyedTs(v int64) *DescribeChannelUserMetricsRequest
	GetDestroyedTs() *int64
}

type DescribeChannelUserMetricsRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
}

func (s DescribeChannelUserMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelUserMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeChannelUserMetricsRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeChannelUserMetricsRequest) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeChannelUserMetricsRequest) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeChannelUserMetricsRequest) SetAppId(v string) *DescribeChannelUserMetricsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelUserMetricsRequest) SetChannelId(v string) *DescribeChannelUserMetricsRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelUserMetricsRequest) SetCreatedTs(v int64) *DescribeChannelUserMetricsRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelUserMetricsRequest) SetDestroyedTs(v int64) *DescribeChannelUserMetricsRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelUserMetricsRequest) Validate() error {
	return dara.Validate(s)
}
