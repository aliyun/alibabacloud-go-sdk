// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateHaVipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHaVipId(v string) *UnassociateHaVipRequest
	GetHaVipId() *string
	SetInstanceId(v string) *UnassociateHaVipRequest
	GetInstanceId() *string
}

type UnassociateHaVipRequest struct {
	// The ID of the HAVIP that you want to disassociate.
	//
	// This parameter is required.
	//
	// example:
	//
	// havip-5p14t****
	HaVipId *string `json:"HaVipId,omitempty" xml:"HaVipId,omitempty"`
	// The ID of the ENS instance or ENI that you want to disassociate from the HAVIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-5ecpqvk****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UnassociateHaVipRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociateHaVipRequest) GoString() string {
	return s.String()
}

func (s *UnassociateHaVipRequest) GetHaVipId() *string {
	return s.HaVipId
}

func (s *UnassociateHaVipRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnassociateHaVipRequest) SetHaVipId(v string) *UnassociateHaVipRequest {
	s.HaVipId = &v
	return s
}

func (s *UnassociateHaVipRequest) SetInstanceId(v string) *UnassociateHaVipRequest {
	s.InstanceId = &v
	return s
}

func (s *UnassociateHaVipRequest) Validate() error {
	return dara.Validate(s)
}
