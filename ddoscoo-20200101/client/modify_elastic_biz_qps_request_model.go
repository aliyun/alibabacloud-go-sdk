// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticBizQpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyElasticBizQpsRequest
	GetInstanceId() *string
	SetMode(v string) *ModifyElasticBizQpsRequest
	GetMode() *string
	SetOpsElasticQps(v int64) *ModifyElasticBizQpsRequest
	GetOpsElasticQps() *int64
}

type ModifyElasticBizQpsRequest struct {
	// The ID of the Anti-DDoS instance.
	//
	// > You can call [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) to query the IDs of all Anti-DDoS instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metering method for the burstable QPS. Valid values:
	//
	// - **month**: Monthly 95th percentile
	//
	// - **day**: Daily 95th percentile
	//
	// This parameter is required.
	//
	// example:
	//
	// month
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The burstable QPS value. After burstable QPS is enabled, the default value is three times the basic QPS, but it cannot exceed the maximum burstable QPS:
	//
	// - **Anti-DDoS Proxy (Chinese Mainland) IPv4 instance**: The maximum value is 300,000 QPS.
	//
	// - **Anti-DDoS Proxy (Chinese Mainland) IPv6 instance**: The maximum value is 100,000 QPS.
	//
	// - **Anti-DDoS Proxy (Outside Chinese Mainland) instance**: The maximum value is 150,000 QPS.
	//
	// example:
	//
	// 300000
	OpsElasticQps *int64 `json:"OpsElasticQps,omitempty" xml:"OpsElasticQps,omitempty"`
}

func (s ModifyElasticBizQpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticBizQpsRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticBizQpsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyElasticBizQpsRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifyElasticBizQpsRequest) GetOpsElasticQps() *int64 {
	return s.OpsElasticQps
}

func (s *ModifyElasticBizQpsRequest) SetInstanceId(v string) *ModifyElasticBizQpsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyElasticBizQpsRequest) SetMode(v string) *ModifyElasticBizQpsRequest {
	s.Mode = &v
	return s
}

func (s *ModifyElasticBizQpsRequest) SetOpsElasticQps(v int64) *ModifyElasticBizQpsRequest {
	s.OpsElasticQps = &v
	return s
}

func (s *ModifyElasticBizQpsRequest) Validate() error {
	return dara.Validate(s)
}
