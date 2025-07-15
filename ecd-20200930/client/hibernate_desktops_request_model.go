// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHibernateDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *HibernateDesktopsRequest
	GetDesktopId() []*string
	SetRegionId(v string) *HibernateDesktopsRequest
	GetRegionId() *string
}

type HibernateDesktopsRequest struct {
	// The IDs of the cloud desktops. You can specify 1 to 20 cloud desktop IDs.
	//
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions supported by Elastic Desktop Service (EDS).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s HibernateDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s HibernateDesktopsRequest) GoString() string {
	return s.String()
}

func (s *HibernateDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *HibernateDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *HibernateDesktopsRequest) SetDesktopId(v []*string) *HibernateDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *HibernateDesktopsRequest) SetRegionId(v string) *HibernateDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *HibernateDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
