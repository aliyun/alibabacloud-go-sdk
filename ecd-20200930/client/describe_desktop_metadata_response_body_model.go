// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktops(v []*DescribeDesktopMetadataResponseBodyDesktops) *DescribeDesktopMetadataResponseBody
	GetDesktops() []*DescribeDesktopMetadataResponseBodyDesktops
	SetNextToken(v string) *DescribeDesktopMetadataResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDesktopMetadataResponseBody
	GetRequestId() *string
}

type DescribeDesktopMetadataResponseBody struct {
	Desktops []*DescribeDesktopMetadataResponseBodyDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" type:"Repeated"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDesktopMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopMetadataResponseBody) GetDesktops() []*DescribeDesktopMetadataResponseBodyDesktops {
	return s.Desktops
}

func (s *DescribeDesktopMetadataResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopMetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDesktopMetadataResponseBody) SetDesktops(v []*DescribeDesktopMetadataResponseBodyDesktops) *DescribeDesktopMetadataResponseBody {
	s.Desktops = v
	return s
}

func (s *DescribeDesktopMetadataResponseBody) SetNextToken(v string) *DescribeDesktopMetadataResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBody) SetRequestId(v string) *DescribeDesktopMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBody) Validate() error {
	if s.Desktops != nil {
		for _, item := range s.Desktops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDesktopMetadataResponseBodyDesktops struct {
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 2020-11-06T08:28Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// dg-3uiojcc0j4kh7****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// example:
	//
	// testDesktopName
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// example:
	//
	// Running
	DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DesktopType   *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// example:
	//
	// 2021-12-31T15:59Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// m-4zfb6zj728hhr****
	ImageId         *string   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	LocalName       *string   `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	ManagementFlags []*string `json:"ManagementFlags,omitempty" xml:"ManagementFlags,omitempty" type:"Repeated"`
	MemberEniIp     *string   `json:"MemberEniIp,omitempty" xml:"MemberEniIp,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-f3s3dgt8dtb0vlqc8
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// serverless_new
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// example:
	//
	// 2020-11-06T08:31Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDesktopMetadataResponseBodyDesktops) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopMetadataResponseBodyDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetDesktopStatus() *string {
	return s.DesktopStatus
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetDesktopType() *string {
	return s.DesktopType
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetManagementFlags() []*string {
	return s.ManagementFlags
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetMemberEniIp() *string {
	return s.MemberEniIp
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetChargeType(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.ChargeType = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetCreationTime(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetDesktopGroupId(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetDesktopId(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetDesktopName(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetDesktopStatus(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetDesktopType(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.DesktopType = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetExpiredTime(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetImageId(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetLocalName(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.LocalName = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetManagementFlags(v []*string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.ManagementFlags = v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetMemberEniIp(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.MemberEniIp = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetOfficeSiteId(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetPlatform(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.Platform = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetRegionId(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetResourceGroupId(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetResourceGroupName(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) SetStartTime(v string) *DescribeDesktopMetadataResponseBodyDesktops {
	s.StartTime = &v
	return s
}

func (s *DescribeDesktopMetadataResponseBodyDesktops) Validate() error {
	return dara.Validate(s)
}
