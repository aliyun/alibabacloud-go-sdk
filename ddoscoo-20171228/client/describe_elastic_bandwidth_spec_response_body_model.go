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
	ElasticBandwidthSpec []*string `json:"ElasticBandwidthSpec,omitempty" xml:"ElasticBandwidthSpec,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
