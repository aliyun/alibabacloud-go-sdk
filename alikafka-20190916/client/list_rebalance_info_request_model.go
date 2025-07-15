// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRebalanceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *ListRebalanceInfoRequest
	GetConsumerId() *string
	SetInstanceId(v string) *ListRebalanceInfoRequest
	GetInstanceId() *string
	SetRegionId(v string) *ListRebalanceInfoRequest
	GetRegionId() *string
}

type ListRebalanceInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// kafka-test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRebalanceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRebalanceInfoRequest) GoString() string {
	return s.String()
}

func (s *ListRebalanceInfoRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ListRebalanceInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRebalanceInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRebalanceInfoRequest) SetConsumerId(v string) *ListRebalanceInfoRequest {
	s.ConsumerId = &v
	return s
}

func (s *ListRebalanceInfoRequest) SetInstanceId(v string) *ListRebalanceInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRebalanceInfoRequest) SetRegionId(v string) *ListRebalanceInfoRequest {
	s.RegionId = &v
	return s
}

func (s *ListRebalanceInfoRequest) Validate() error {
	return dara.Validate(s)
}
