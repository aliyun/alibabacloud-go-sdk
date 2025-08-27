// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTripCCInfoQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessInstanceId(v string) *TripCCInfoQueryRequest
	GetBusinessInstanceId() *string
	SetNodeId(v string) *TripCCInfoQueryRequest
	GetNodeId() *string
	SetThirdBusinessId(v string) *TripCCInfoQueryRequest
	GetThirdBusinessId() *string
}

type TripCCInfoQueryRequest struct {
	// example:
	//
	// 2024060710160003300008684
	BusinessInstanceId *string `json:"business_instance_id,omitempty" xml:"business_instance_id,omitempty"`
	// example:
	//
	// 458003
	NodeId *string `json:"node_id,omitempty" xml:"node_id,omitempty"`
	// example:
	//
	// 2024060710160003300008684
	ThirdBusinessId *string `json:"third_business_id,omitempty" xml:"third_business_id,omitempty"`
}

func (s TripCCInfoQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s TripCCInfoQueryRequest) GoString() string {
	return s.String()
}

func (s *TripCCInfoQueryRequest) GetBusinessInstanceId() *string {
	return s.BusinessInstanceId
}

func (s *TripCCInfoQueryRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *TripCCInfoQueryRequest) GetThirdBusinessId() *string {
	return s.ThirdBusinessId
}

func (s *TripCCInfoQueryRequest) SetBusinessInstanceId(v string) *TripCCInfoQueryRequest {
	s.BusinessInstanceId = &v
	return s
}

func (s *TripCCInfoQueryRequest) SetNodeId(v string) *TripCCInfoQueryRequest {
	s.NodeId = &v
	return s
}

func (s *TripCCInfoQueryRequest) SetThirdBusinessId(v string) *TripCCInfoQueryRequest {
	s.ThirdBusinessId = &v
	return s
}

func (s *TripCCInfoQueryRequest) Validate() error {
	return dara.Validate(s)
}
