// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainFrameRateAndBitRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainFrameRateAndBitRateDataRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveDomainFrameRateAndBitRateDataRequest
	GetOwnerId() *int64
	SetQueryTime(v string) *DescribeLiveDomainFrameRateAndBitRateDataRequest
	GetQueryTime() *string
	SetRegionId(v string) *DescribeLiveDomainFrameRateAndBitRateDataRequest
	GetRegionId() *string
}

type DescribeLiveDomainFrameRateAndBitRateDataRequest struct {
	// The ingest domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The point of time to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-02-21T08:00:00Z
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveDomainFrameRateAndBitRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainFrameRateAndBitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataRequest) GetQueryTime() *string {
	return s.QueryTime
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataRequest) SetDomainName(v string) *DescribeLiveDomainFrameRateAndBitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataRequest) SetOwnerId(v int64) *DescribeLiveDomainFrameRateAndBitRateDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataRequest) SetQueryTime(v string) *DescribeLiveDomainFrameRateAndBitRateDataRequest {
	s.QueryTime = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataRequest) SetRegionId(v string) *DescribeLiveDomainFrameRateAndBitRateDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataRequest) Validate() error {
	return dara.Validate(s)
}
