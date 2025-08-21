// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RealLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ConfigLayer4RealLimitRequest
	GetInstanceId() *string
	SetLimitValue(v int64) *ConfigLayer4RealLimitRequest
	GetLimitValue() *int64
}

type ConfigLayer4RealLimitRequest struct {
	// The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies the threshold of the clean bandwidth. Valid values: 0 to 15360. The value 0 indicates that rate limiting is never triggered. Unit: Mbit/s
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	LimitValue *int64 `json:"LimitValue,omitempty" xml:"LimitValue,omitempty"`
}

func (s ConfigLayer4RealLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RealLimitRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RealLimitRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigLayer4RealLimitRequest) GetLimitValue() *int64 {
	return s.LimitValue
}

func (s *ConfigLayer4RealLimitRequest) SetInstanceId(v string) *ConfigLayer4RealLimitRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigLayer4RealLimitRequest) SetLimitValue(v int64) *ConfigLayer4RealLimitRequest {
	s.LimitValue = &v
	return s
}

func (s *ConfigLayer4RealLimitRequest) Validate() error {
	return dara.Validate(s)
}
