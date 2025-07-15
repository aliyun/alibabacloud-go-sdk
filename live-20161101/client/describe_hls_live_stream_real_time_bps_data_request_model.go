// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHlsLiveStreamRealTimeBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeHlsLiveStreamRealTimeBpsDataRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeHlsLiveStreamRealTimeBpsDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeHlsLiveStreamRealTimeBpsDataRequest
	GetRegionId() *string
	SetTime(v string) *DescribeHlsLiveStreamRealTimeBpsDataRequest
	GetTime() *string
}

type DescribeHlsLiveStreamRealTimeBpsDataRequest struct {
	// The domain names to query. Separate them with commas (,). A domain name cannot contain double-byte characters.
	//
	// example:
	//
	// live.aiyun.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-08-08T00:00:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeHlsLiveStreamRealTimeBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHlsLiveStreamRealTimeBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataRequest) GetTime() *string {
	return s.Time
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataRequest) SetDomainName(v string) *DescribeHlsLiveStreamRealTimeBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataRequest) SetOwnerId(v int64) *DescribeHlsLiveStreamRealTimeBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataRequest) SetRegionId(v string) *DescribeHlsLiveStreamRealTimeBpsDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataRequest) SetTime(v string) *DescribeHlsLiveStreamRealTimeBpsDataRequest {
	s.Time = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
