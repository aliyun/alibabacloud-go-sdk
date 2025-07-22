// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelAreaDistributionStatDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeChannelAreaDistributionStatDataRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeChannelAreaDistributionStatDataRequest
	GetChannelId() *string
	SetCreatedTs(v int64) *DescribeChannelAreaDistributionStatDataRequest
	GetCreatedTs() *int64
	SetDestroyedTs(v int64) *DescribeChannelAreaDistributionStatDataRequest
	GetDestroyedTs() *int64
	SetParentArea(v string) *DescribeChannelAreaDistributionStatDataRequest
	GetParentArea() *string
}

type DescribeChannelAreaDistributionStatDataRequest struct {
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
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	ParentArea  *string `json:"ParentArea,omitempty" xml:"ParentArea,omitempty"`
}

func (s DescribeChannelAreaDistributionStatDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelAreaDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelAreaDistributionStatDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeChannelAreaDistributionStatDataRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeChannelAreaDistributionStatDataRequest) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeChannelAreaDistributionStatDataRequest) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeChannelAreaDistributionStatDataRequest) GetParentArea() *string {
	return s.ParentArea
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetAppId(v string) *DescribeChannelAreaDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetChannelId(v string) *DescribeChannelAreaDistributionStatDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetCreatedTs(v int64) *DescribeChannelAreaDistributionStatDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetDestroyedTs(v int64) *DescribeChannelAreaDistributionStatDataRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetParentArea(v string) *DescribeChannelAreaDistributionStatDataRequest {
	s.ParentArea = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataRequest) Validate() error {
	return dara.Validate(s)
}
