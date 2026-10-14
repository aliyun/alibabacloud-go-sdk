// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticBandWidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticBandwidth(v int32) *ModifyElasticBandWidthRequest
	GetElasticBandwidth() *int32
	SetInstanceId(v string) *ModifyElasticBandWidthRequest
	GetInstanceId() *string
}

type ModifyElasticBandWidthRequest struct {
	// The burstable protection bandwidth that you want to set. Unit: Gbps.
	//
	// > You can call [DescribeElasticBandwidthSpec](https://help.aliyun.com/document_detail/91502.html) to query the available burstable protection bandwidth specifications.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	ElasticBandwidth *int32 `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	// The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// > The instance must be in a normal state. You can call [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) to query the IDs of all Anti-DDoS Pro or Anti-DDoS Premium instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyElasticBandWidthRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticBandWidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandWidthRequest) GetElasticBandwidth() *int32 {
	return s.ElasticBandwidth
}

func (s *ModifyElasticBandWidthRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyElasticBandWidthRequest) SetElasticBandwidth(v int32) *ModifyElasticBandWidthRequest {
	s.ElasticBandwidth = &v
	return s
}

func (s *ModifyElasticBandWidthRequest) SetInstanceId(v string) *ModifyElasticBandWidthRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyElasticBandWidthRequest) Validate() error {
	return dara.Validate(s)
}
