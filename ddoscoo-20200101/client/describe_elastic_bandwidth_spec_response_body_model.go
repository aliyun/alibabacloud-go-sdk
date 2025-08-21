// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticBandwidthSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElasticBandwidthSpec(v []*string) *DescribeElasticBandwidthSpecResponseBody
	GetElasticBandwidthSpec() []*string
	SetRequestId(v string) *DescribeElasticBandwidthSpecResponseBody
	GetRequestId() *string
}

type DescribeElasticBandwidthSpecResponseBody struct {
	// An array that consists of the available burstable protection bandwidths. Unit: Gbit/s.
	ElasticBandwidthSpec []*string `json:"ElasticBandwidthSpec,omitempty" xml:"ElasticBandwidthSpec,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElasticBandwidthSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticBandwidthSpecResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticBandwidthSpecResponseBody) GetElasticBandwidthSpec() []*string {
	return s.ElasticBandwidthSpec
}

func (s *DescribeElasticBandwidthSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticBandwidthSpecResponseBody) SetElasticBandwidthSpec(v []*string) *DescribeElasticBandwidthSpecResponseBody {
	s.ElasticBandwidthSpec = v
	return s
}

func (s *DescribeElasticBandwidthSpecResponseBody) SetRequestId(v string) *DescribeElasticBandwidthSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticBandwidthSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
