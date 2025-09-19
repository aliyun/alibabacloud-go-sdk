// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceCount(v int32) *DescribeImageStatisticsResponseBody
	GetInstanceCount() *int32
	SetRequestId(v string) *DescribeImageStatisticsResponseBody
	GetRequestId() *string
	SetRiskInstanceCount(v int32) *DescribeImageStatisticsResponseBody
	GetRiskInstanceCount() *int32
}

type DescribeImageStatisticsResponseBody struct {
	// The number of container images in your assets. Only Container Registry Enterprise Edition instances are counted.
	//
	// example:
	//
	// 5
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4E5BFDCF-B9DD-430D-9DA4-151BCB581C9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of container images that have security risks. Only Container Registry Enterprise Edition instances are counted.
	//
	// example:
	//
	// 2
	RiskInstanceCount *int32 `json:"RiskInstanceCount,omitempty" xml:"RiskInstanceCount,omitempty"`
}

func (s DescribeImageStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageStatisticsResponseBody) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeImageStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageStatisticsResponseBody) GetRiskInstanceCount() *int32 {
	return s.RiskInstanceCount
}

func (s *DescribeImageStatisticsResponseBody) SetInstanceCount(v int32) *DescribeImageStatisticsResponseBody {
	s.InstanceCount = &v
	return s
}

func (s *DescribeImageStatisticsResponseBody) SetRequestId(v string) *DescribeImageStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageStatisticsResponseBody) SetRiskInstanceCount(v int32) *DescribeImageStatisticsResponseBody {
	s.RiskInstanceCount = &v
	return s
}

func (s *DescribeImageStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
