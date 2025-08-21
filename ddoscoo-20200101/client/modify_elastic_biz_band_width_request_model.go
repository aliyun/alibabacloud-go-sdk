// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticBizBandWidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticBizBandwidth(v int32) *ModifyElasticBizBandWidthRequest
	GetElasticBizBandwidth() *int32
	SetInstanceId(v string) *ModifyElasticBizBandWidthRequest
	GetInstanceId() *string
	SetMode(v string) *ModifyElasticBizBandWidthRequest
	GetMode() *string
}

type ModifyElasticBizBandWidthRequest struct {
	// The burstable clean bandwidth. Unit: Mbit/s. The burstable clean bandwidth cannot exceed nine times the clean bandwidth of your Anti-DDoS Pro or Anti-DDoS Premium instance, and the sum of the clean bandwidth and the burstable clean bandwidth cannot exceed the maximum clean bandwidth that is supported by your instance. The value 0 indicates that the burstable clean bandwidth feature is disabled. You can disable the burstable clean bandwidth feature once a month.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ElasticBizBandwidth *int32 `json:"ElasticBizBandwidth,omitempty" xml:"ElasticBizBandwidth,omitempty"`
	// The ID of the instance.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-tl32morr****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metering method of the burstable clean bandwidth feature. Valid values:
	//
	// 	- **month**: the metering method of monthly 95th percentile
	//
	// 	- **day**: the metering method of daily 95th percentile
	//
	// example:
	//
	// month
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ModifyElasticBizBandWidthRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticBizBandWidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticBizBandWidthRequest) GetElasticBizBandwidth() *int32 {
	return s.ElasticBizBandwidth
}

func (s *ModifyElasticBizBandWidthRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyElasticBizBandWidthRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifyElasticBizBandWidthRequest) SetElasticBizBandwidth(v int32) *ModifyElasticBizBandWidthRequest {
	s.ElasticBizBandwidth = &v
	return s
}

func (s *ModifyElasticBizBandWidthRequest) SetInstanceId(v string) *ModifyElasticBizBandWidthRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyElasticBizBandWidthRequest) SetMode(v string) *ModifyElasticBizBandWidthRequest {
	s.Mode = &v
	return s
}

func (s *ModifyElasticBizBandWidthRequest) Validate() error {
	return dara.Validate(s)
}
