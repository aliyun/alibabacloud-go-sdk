// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQoeMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeQoeMetricDataRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeQoeMetricDataRequest
	GetChannelId() *string
	SetCreatedTs(v int64) *DescribeQoeMetricDataRequest
	GetCreatedTs() *int64
	SetDestroyedTs(v int64) *DescribeQoeMetricDataRequest
	GetDestroyedTs() *int64
	SetUserId(v string) *DescribeQoeMetricDataRequest
	GetUserId() *string
}

type DescribeQoeMetricDataRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeQoeMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQoeMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeQoeMetricDataRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeQoeMetricDataRequest) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeQoeMetricDataRequest) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeQoeMetricDataRequest) GetUserId() *string {
	return s.UserId
}

func (s *DescribeQoeMetricDataRequest) SetAppId(v string) *DescribeQoeMetricDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQoeMetricDataRequest) SetChannelId(v string) *DescribeQoeMetricDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeQoeMetricDataRequest) SetCreatedTs(v int64) *DescribeQoeMetricDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeQoeMetricDataRequest) SetDestroyedTs(v int64) *DescribeQoeMetricDataRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeQoeMetricDataRequest) SetUserId(v string) *DescribeQoeMetricDataRequest {
	s.UserId = &v
	return s
}

func (s *DescribeQoeMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
