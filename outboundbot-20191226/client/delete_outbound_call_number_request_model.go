// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOutboundCallNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteOutboundCallNumberRequest
	GetInstanceId() *string
	SetOutboundCallNumberId(v string) *DeleteOutboundCallNumberRequest
	GetOutboundCallNumberId() *string
}

type DeleteOutboundCallNumberRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 33014787-cc13-49d3-ab2f-a98aa8f15fbb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ffa367e0-58f3-43b6-9615-c63db99c5add
	OutboundCallNumberId *string `json:"OutboundCallNumberId,omitempty" xml:"OutboundCallNumberId,omitempty"`
}

func (s DeleteOutboundCallNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOutboundCallNumberRequest) GoString() string {
	return s.String()
}

func (s *DeleteOutboundCallNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteOutboundCallNumberRequest) GetOutboundCallNumberId() *string {
	return s.OutboundCallNumberId
}

func (s *DeleteOutboundCallNumberRequest) SetInstanceId(v string) *DeleteOutboundCallNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteOutboundCallNumberRequest) SetOutboundCallNumberId(v string) *DeleteOutboundCallNumberRequest {
	s.OutboundCallNumberId = &v
	return s
}

func (s *DeleteOutboundCallNumberRequest) Validate() error {
	return dara.Validate(s)
}
