// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndPointMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeEndPointMetricDataRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeEndPointMetricDataRequest
	GetChannelId() *string
	SetCreatedTs(v int64) *DescribeEndPointMetricDataRequest
	GetCreatedTs() *int64
	SetDestroyedTs(v int64) *DescribeEndPointMetricDataRequest
	GetDestroyedTs() *int64
	SetMetrics(v string) *DescribeEndPointMetricDataRequest
	GetMetrics() *string
	SetPubCallIdList(v string) *DescribeEndPointMetricDataRequest
	GetPubCallIdList() *string
	SetPubUserId(v string) *DescribeEndPointMetricDataRequest
	GetPubUserId() *string
	SetSubUserId(v string) *DescribeEndPointMetricDataRequest
	GetSubUserId() *string
}

type DescribeEndPointMetricDataRequest struct {
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
	// VIDEO_STUCK_CAMERA
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// example:
	//
	// testcall1,testcall2
	PubCallIdList *string `json:"PubCallIdList,omitempty" xml:"PubCallIdList,omitempty"`
	// example:
	//
	// testuserid
	PubUserId *string `json:"PubUserId,omitempty" xml:"PubUserId,omitempty"`
	// example:
	//
	// testuserid
	SubUserId *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribeEndPointMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndPointMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeEndPointMetricDataRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeEndPointMetricDataRequest) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeEndPointMetricDataRequest) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeEndPointMetricDataRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *DescribeEndPointMetricDataRequest) GetPubCallIdList() *string {
	return s.PubCallIdList
}

func (s *DescribeEndPointMetricDataRequest) GetPubUserId() *string {
	return s.PubUserId
}

func (s *DescribeEndPointMetricDataRequest) GetSubUserId() *string {
	return s.SubUserId
}

func (s *DescribeEndPointMetricDataRequest) SetAppId(v string) *DescribeEndPointMetricDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetChannelId(v string) *DescribeEndPointMetricDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetCreatedTs(v int64) *DescribeEndPointMetricDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetDestroyedTs(v int64) *DescribeEndPointMetricDataRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetMetrics(v string) *DescribeEndPointMetricDataRequest {
	s.Metrics = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetPubCallIdList(v string) *DescribeEndPointMetricDataRequest {
	s.PubCallIdList = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetPubUserId(v string) *DescribeEndPointMetricDataRequest {
	s.PubUserId = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetSubUserId(v string) *DescribeEndPointMetricDataRequest {
	s.SubUserId = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
