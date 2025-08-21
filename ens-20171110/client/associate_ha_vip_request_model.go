// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateHaVipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHaVipId(v string) *AssociateHaVipRequest
	GetHaVipId() *string
	SetInstanceId(v string) *AssociateHaVipRequest
	GetInstanceId() *string
	SetInstanceType(v string) *AssociateHaVipRequest
	GetInstanceType() *string
}

type AssociateHaVipRequest struct {
	// The ID of the HAVIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// havip-5p14t****
	HaVipId *string `json:"HaVipId,omitempty" xml:"HaVipId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-50c4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the instance to be associated with the HAVIP. Valid values:
	//
	// 	- EnsInstance (default): ENS instance
	//
	// 	- NetworkInterface: elastic network interface (ENI)
	//
	// example:
	//
	// EnsInstance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s AssociateHaVipRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateHaVipRequest) GoString() string {
	return s.String()
}

func (s *AssociateHaVipRequest) GetHaVipId() *string {
	return s.HaVipId
}

func (s *AssociateHaVipRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssociateHaVipRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *AssociateHaVipRequest) SetHaVipId(v string) *AssociateHaVipRequest {
	s.HaVipId = &v
	return s
}

func (s *AssociateHaVipRequest) SetInstanceId(v string) *AssociateHaVipRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateHaVipRequest) SetInstanceType(v string) *AssociateHaVipRequest {
	s.InstanceType = &v
	return s
}

func (s *AssociateHaVipRequest) Validate() error {
	return dara.Validate(s)
}
