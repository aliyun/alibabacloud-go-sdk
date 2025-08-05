// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAssetStatisticResponseBody
	GetRequestId() *string
	SetResourceSpecStatistic(v *DescribeAssetStatisticResponseBodyResourceSpecStatistic) *DescribeAssetStatisticResponseBody
	GetResourceSpecStatistic() *DescribeAssetStatisticResponseBodyResourceSpecStatistic
}

type DescribeAssetStatisticResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 850A84******25g4d2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics on specifications.
	ResourceSpecStatistic *DescribeAssetStatisticResponseBodyResourceSpecStatistic `json:"ResourceSpecStatistic,omitempty" xml:"ResourceSpecStatistic,omitempty" type:"Struct"`
}

func (s DescribeAssetStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAssetStatisticResponseBody) GetResourceSpecStatistic() *DescribeAssetStatisticResponseBodyResourceSpecStatistic {
	return s.ResourceSpecStatistic
}

func (s *DescribeAssetStatisticResponseBody) SetRequestId(v string) *DescribeAssetStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetStatisticResponseBody) SetResourceSpecStatistic(v *DescribeAssetStatisticResponseBodyResourceSpecStatistic) *DescribeAssetStatisticResponseBody {
	s.ResourceSpecStatistic = v
	return s
}

func (s *DescribeAssetStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAssetStatisticResponseBodyResourceSpecStatistic struct {
	// The number of public IP addresses that can be protected.
	//
	// example:
	//
	// 20
	IpNumSpec *int32 `json:"IpNumSpec,omitempty" xml:"IpNumSpec,omitempty"`
	// The number of public IP addresses that are protected.
	//
	// example:
	//
	// 10
	IpNumUsed *int32 `json:"IpNumUsed,omitempty" xml:"IpNumUsed,omitempty"`
	// The number of public IP addresses that can enable data leakage detection.
	//
	// example:
	//
	// 0
	SensitiveDataIpNumSpec *int64 `json:"SensitiveDataIpNumSpec,omitempty" xml:"SensitiveDataIpNumSpec,omitempty"`
	// The number of public IP addresses that enabled data leakage detection.
	//
	// example:
	//
	// 0
	SensitiveDataIpNumUsed *int64 `json:"SensitiveDataIpNumUsed,omitempty" xml:"SensitiveDataIpNumUsed,omitempty"`
}

func (s DescribeAssetStatisticResponseBodyResourceSpecStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetStatisticResponseBodyResourceSpecStatistic) GoString() string {
	return s.String()
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) GetIpNumSpec() *int32 {
	return s.IpNumSpec
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) GetIpNumUsed() *int32 {
	return s.IpNumUsed
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) GetSensitiveDataIpNumSpec() *int64 {
	return s.SensitiveDataIpNumSpec
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) GetSensitiveDataIpNumUsed() *int64 {
	return s.SensitiveDataIpNumUsed
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) SetIpNumSpec(v int32) *DescribeAssetStatisticResponseBodyResourceSpecStatistic {
	s.IpNumSpec = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) SetIpNumUsed(v int32) *DescribeAssetStatisticResponseBodyResourceSpecStatistic {
	s.IpNumUsed = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) SetSensitiveDataIpNumSpec(v int64) *DescribeAssetStatisticResponseBodyResourceSpecStatistic {
	s.SensitiveDataIpNumSpec = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) SetSensitiveDataIpNumUsed(v int64) *DescribeAssetStatisticResponseBodyResourceSpecStatistic {
	s.SensitiveDataIpNumUsed = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) Validate() error {
	return dara.Validate(s)
}
