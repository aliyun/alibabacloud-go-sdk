// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOutboundCallRestrictionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteOutboundCallRestrictionRequest
	GetInstanceId() *string
	SetRestrictionIdList(v []*string) *DeleteOutboundCallRestrictionRequest
	GetRestrictionIdList() []*string
}

type DeleteOutboundCallRestrictionRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of outbound call restriction IDs.
	RestrictionIdList []*string `json:"RestrictionIdList,omitempty" xml:"RestrictionIdList,omitempty" type:"Repeated"`
}

func (s DeleteOutboundCallRestrictionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOutboundCallRestrictionRequest) GoString() string {
	return s.String()
}

func (s *DeleteOutboundCallRestrictionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteOutboundCallRestrictionRequest) GetRestrictionIdList() []*string {
	return s.RestrictionIdList
}

func (s *DeleteOutboundCallRestrictionRequest) SetInstanceId(v string) *DeleteOutboundCallRestrictionRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteOutboundCallRestrictionRequest) SetRestrictionIdList(v []*string) *DeleteOutboundCallRestrictionRequest {
	s.RestrictionIdList = v
	return s
}

func (s *DeleteOutboundCallRestrictionRequest) Validate() error {
	return dara.Validate(s)
}
