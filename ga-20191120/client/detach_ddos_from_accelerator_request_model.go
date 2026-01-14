// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDdosFromAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DetachDdosFromAcceleratorRequest
	GetAcceleratorId() *string
	SetDdosConfigList(v []*DetachDdosFromAcceleratorRequestDdosConfigList) *DetachDdosFromAcceleratorRequest
	GetDdosConfigList() []*DetachDdosFromAcceleratorRequestDdosConfigList
	SetDryRun(v bool) *DetachDdosFromAcceleratorRequest
	GetDryRun() *bool
	SetRegionId(v string) *DetachDdosFromAcceleratorRequest
	GetRegionId() *string
}

type DetachDdosFromAcceleratorRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId  *string                                           `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	DdosConfigList []*DetachDdosFromAcceleratorRequestDdosConfigList `json:"DdosConfigList,omitempty" xml:"DdosConfigList,omitempty" type:"Repeated"`
	DryRun         *bool                                             `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachDdosFromAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachDdosFromAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *DetachDdosFromAcceleratorRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DetachDdosFromAcceleratorRequest) GetDdosConfigList() []*DetachDdosFromAcceleratorRequestDdosConfigList {
	return s.DdosConfigList
}

func (s *DetachDdosFromAcceleratorRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DetachDdosFromAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachDdosFromAcceleratorRequest) SetAcceleratorId(v string) *DetachDdosFromAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DetachDdosFromAcceleratorRequest) SetDdosConfigList(v []*DetachDdosFromAcceleratorRequestDdosConfigList) *DetachDdosFromAcceleratorRequest {
	s.DdosConfigList = v
	return s
}

func (s *DetachDdosFromAcceleratorRequest) SetDryRun(v bool) *DetachDdosFromAcceleratorRequest {
	s.DryRun = &v
	return s
}

func (s *DetachDdosFromAcceleratorRequest) SetRegionId(v string) *DetachDdosFromAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *DetachDdosFromAcceleratorRequest) Validate() error {
	if s.DdosConfigList != nil {
		for _, item := range s.DdosConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetachDdosFromAcceleratorRequestDdosConfigList struct {
	DdosId       *string `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
}

func (s DetachDdosFromAcceleratorRequestDdosConfigList) String() string {
	return dara.Prettify(s)
}

func (s DetachDdosFromAcceleratorRequestDdosConfigList) GoString() string {
	return s.String()
}

func (s *DetachDdosFromAcceleratorRequestDdosConfigList) GetDdosId() *string {
	return s.DdosId
}

func (s *DetachDdosFromAcceleratorRequestDdosConfigList) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *DetachDdosFromAcceleratorRequestDdosConfigList) SetDdosId(v string) *DetachDdosFromAcceleratorRequestDdosConfigList {
	s.DdosId = &v
	return s
}

func (s *DetachDdosFromAcceleratorRequestDdosConfigList) SetDdosRegionId(v string) *DetachDdosFromAcceleratorRequestDdosConfigList {
	s.DdosRegionId = &v
	return s
}

func (s *DetachDdosFromAcceleratorRequestDdosConfigList) Validate() error {
	return dara.Validate(s)
}
