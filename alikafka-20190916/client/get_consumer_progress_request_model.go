// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *GetConsumerProgressRequest
	GetConsumerId() *string
	SetHideLastTimestamp(v bool) *GetConsumerProgressRequest
	GetHideLastTimestamp() *bool
	SetInstanceId(v string) *GetConsumerProgressRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetConsumerProgressRequest
	GetRegionId() *string
}

type GetConsumerProgressRequest struct {
	// The name of the consumer group.
	//
	// This parameter is required.
	//
	// example:
	//
	// kafka-test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// Specifies whether to hide the \\`LastTimestamp\\` parameter. The default value is false. Set this parameter to true for better performance.
	//
	// > - If you set this parameter to true, -1 is returned for the \\`LastTimestamp\\` parameter. Otherwise, a specific value is returned. This parameter is supported only for topics of cloud storage on provisioned instances.
	//
	// >
	//
	// > - This operation processes a large amount of data and consumes a high degree of performance. Set this parameter to true to reduce the processing time.
	//
	// example:
	//
	// true
	HideLastTimestamp *bool `json:"HideLastTimestamp,omitempty" xml:"HideLastTimestamp,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetConsumerProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerProgressRequest) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *GetConsumerProgressRequest) GetHideLastTimestamp() *bool {
	return s.HideLastTimestamp
}

func (s *GetConsumerProgressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetConsumerProgressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetConsumerProgressRequest) SetConsumerId(v string) *GetConsumerProgressRequest {
	s.ConsumerId = &v
	return s
}

func (s *GetConsumerProgressRequest) SetHideLastTimestamp(v bool) *GetConsumerProgressRequest {
	s.HideLastTimestamp = &v
	return s
}

func (s *GetConsumerProgressRequest) SetInstanceId(v string) *GetConsumerProgressRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerProgressRequest) SetRegionId(v string) *GetConsumerProgressRequest {
	s.RegionId = &v
	return s
}

func (s *GetConsumerProgressRequest) Validate() error {
	return dara.Validate(s)
}
