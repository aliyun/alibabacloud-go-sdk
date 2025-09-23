// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticBandwidthSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeElasticBandwidthSpecRequest
	GetInstanceId() *string
	SetSourceIp(v string) *DescribeElasticBandwidthSpecRequest
	GetSourceIp() *string
}

type DescribeElasticBandwidthSpecRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeElasticBandwidthSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticBandwidthSpecRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticBandwidthSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeElasticBandwidthSpecRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeElasticBandwidthSpecRequest) SetInstanceId(v string) *DescribeElasticBandwidthSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeElasticBandwidthSpecRequest) SetSourceIp(v string) *DescribeElasticBandwidthSpecRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeElasticBandwidthSpecRequest) Validate() error {
	return dara.Validate(s)
}
