// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostAvailabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v []*int64) *DeleteHostAvailabilityRequest
	GetId() []*int64
	SetRegionId(v string) *DeleteHostAvailabilityRequest
	GetRegionId() *string
}

type DeleteHostAvailabilityRequest struct {
	// The task ID. Valid values of N: 1 to 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12456
	Id       []*int64 `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
	RegionId *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteHostAvailabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *DeleteHostAvailabilityRequest) GetId() []*int64 {
	return s.Id
}

func (s *DeleteHostAvailabilityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHostAvailabilityRequest) SetId(v []*int64) *DeleteHostAvailabilityRequest {
	s.Id = v
	return s
}

func (s *DeleteHostAvailabilityRequest) SetRegionId(v string) *DeleteHostAvailabilityRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHostAvailabilityRequest) Validate() error {
	return dara.Validate(s)
}
