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
}

type DescribeElasticBandwidthSpecRequest struct {
	// The ID of the instance.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *DescribeElasticBandwidthSpecRequest) SetInstanceId(v string) *DescribeElasticBandwidthSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeElasticBandwidthSpecRequest) Validate() error {
	return dara.Validate(s)
}
