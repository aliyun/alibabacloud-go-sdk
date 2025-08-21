// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHaVipAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHaVipId(v string) *ModifyHaVipAttributeRequest
	GetHaVipId() *string
	SetName(v string) *ModifyHaVipAttributeRequest
	GetName() *string
}

type ModifyHaVipAttributeRequest struct {
	// The ID of the HAVIP that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// havip-52y28****
	HaVipId *string `json:"HaVipId,omitempty" xml:"HaVipId,omitempty"`
	// The name of the HAVIP. The name must be 1 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyHaVipAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHaVipAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyHaVipAttributeRequest) GetHaVipId() *string {
	return s.HaVipId
}

func (s *ModifyHaVipAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyHaVipAttributeRequest) SetHaVipId(v string) *ModifyHaVipAttributeRequest {
	s.HaVipId = &v
	return s
}

func (s *ModifyHaVipAttributeRequest) SetName(v string) *ModifyHaVipAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyHaVipAttributeRequest) Validate() error {
	return dara.Validate(s)
}
