// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceMQTTParameters interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SourceMQTTParameters
	GetInstanceId() *string
	SetRegionId(v string) *SourceMQTTParameters
	GetRegionId() *string
	SetTopic(v string) *SourceMQTTParameters
	GetTopic() *string
}

type SourceMQTTParameters struct {
	// The ID of the ApsaraMQ for MQTT instance.
	//
	// example:
	//
	// mqtt-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the ApsaraMQ for MQTT instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the topic in the ApsaraMQ for MQTT instance.
	//
	// example:
	//
	// testTopic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s SourceMQTTParameters) String() string {
	return dara.Prettify(s)
}

func (s SourceMQTTParameters) GoString() string {
	return s.String()
}

func (s *SourceMQTTParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SourceMQTTParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *SourceMQTTParameters) GetTopic() *string {
	return s.Topic
}

func (s *SourceMQTTParameters) SetInstanceId(v string) *SourceMQTTParameters {
	s.InstanceId = &v
	return s
}

func (s *SourceMQTTParameters) SetRegionId(v string) *SourceMQTTParameters {
	s.RegionId = &v
	return s
}

func (s *SourceMQTTParameters) SetTopic(v string) *SourceMQTTParameters {
	s.Topic = &v
	return s
}

func (s *SourceMQTTParameters) Validate() error {
	return dara.Validate(s)
}
