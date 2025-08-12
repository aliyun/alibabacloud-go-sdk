// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableHostAvailabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v []*int64) *DisableHostAvailabilityRequest
	GetId() []*int64
	SetRegionId(v string) *DisableHostAvailabilityRequest
	GetRegionId() *string
}

type DisableHostAvailabilityRequest struct {
	// The ID of the availability monitoring task. Valid values of N: 1 to 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id       []*int64 `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
	RegionId *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableHostAvailabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableHostAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *DisableHostAvailabilityRequest) GetId() []*int64 {
	return s.Id
}

func (s *DisableHostAvailabilityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableHostAvailabilityRequest) SetId(v []*int64) *DisableHostAvailabilityRequest {
	s.Id = v
	return s
}

func (s *DisableHostAvailabilityRequest) SetRegionId(v string) *DisableHostAvailabilityRequest {
	s.RegionId = &v
	return s
}

func (s *DisableHostAvailabilityRequest) Validate() error {
	return dara.Validate(s)
}
