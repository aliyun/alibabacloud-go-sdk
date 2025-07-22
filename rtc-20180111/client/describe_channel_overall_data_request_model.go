// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelOverallDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeChannelOverallDataRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeChannelOverallDataRequest
	GetChannelId() *string
	SetCreatedTs(v int64) *DescribeChannelOverallDataRequest
	GetCreatedTs() *int64
	SetDestroyedTs(v int64) *DescribeChannelOverallDataRequest
	GetDestroyedTs() *int64
}

type DescribeChannelOverallDataRequest struct {
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
	// 1615893133
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1615893757
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
}

func (s DescribeChannelOverallDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelOverallDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeChannelOverallDataRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeChannelOverallDataRequest) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeChannelOverallDataRequest) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeChannelOverallDataRequest) SetAppId(v string) *DescribeChannelOverallDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelOverallDataRequest) SetChannelId(v string) *DescribeChannelOverallDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelOverallDataRequest) SetCreatedTs(v int64) *DescribeChannelOverallDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelOverallDataRequest) SetDestroyedTs(v int64) *DescribeChannelOverallDataRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelOverallDataRequest) Validate() error {
	return dara.Validate(s)
}
