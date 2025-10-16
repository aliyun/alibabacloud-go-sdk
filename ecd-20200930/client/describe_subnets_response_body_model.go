// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubnetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeSubnetsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeSubnetsResponseBody
	GetRequestId() *string
	SetSubnets(v []*DescribeSubnetsResponseBodySubnets) *DescribeSubnetsResponseBody
	GetSubnets() []*DescribeSubnetsResponseBodySubnets
}

type DescribeSubnetsResponseBody struct {
	NextToken *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Subnets   []*DescribeSubnetsResponseBodySubnets `json:"Subnets,omitempty" xml:"Subnets,omitempty" type:"Repeated"`
}

func (s DescribeSubnetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubnetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubnetsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSubnetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSubnetsResponseBody) GetSubnets() []*DescribeSubnetsResponseBodySubnets {
	return s.Subnets
}

func (s *DescribeSubnetsResponseBody) SetNextToken(v string) *DescribeSubnetsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSubnetsResponseBody) SetRequestId(v string) *DescribeSubnetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubnetsResponseBody) SetSubnets(v []*DescribeSubnetsResponseBodySubnets) *DescribeSubnetsResponseBody {
	s.Subnets = v
	return s
}

func (s *DescribeSubnetsResponseBody) Validate() error {
	if s.Subnets != nil {
		for _, item := range s.Subnets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSubnetsResponseBodySubnets struct {
	CidrBlock             *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OfficeSiteId          *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubnetId              *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	TotalEdsCount         *int32  `json:"TotalEdsCount,omitempty" xml:"TotalEdsCount,omitempty"`
	TotalEdsCountForGroup *int32  `json:"TotalEdsCountForGroup,omitempty" xml:"TotalEdsCountForGroup,omitempty"`
	ZoneId                *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSubnetsResponseBodySubnets) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubnetsResponseBodySubnets) GoString() string {
	return s.String()
}

func (s *DescribeSubnetsResponseBodySubnets) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeSubnetsResponseBodySubnets) GetName() *string {
	return s.Name
}

func (s *DescribeSubnetsResponseBodySubnets) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeSubnetsResponseBodySubnets) GetStatus() *string {
	return s.Status
}

func (s *DescribeSubnetsResponseBodySubnets) GetSubnetId() *string {
	return s.SubnetId
}

func (s *DescribeSubnetsResponseBodySubnets) GetTotalEdsCount() *int32 {
	return s.TotalEdsCount
}

func (s *DescribeSubnetsResponseBodySubnets) GetTotalEdsCountForGroup() *int32 {
	return s.TotalEdsCountForGroup
}

func (s *DescribeSubnetsResponseBodySubnets) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeSubnetsResponseBodySubnets) SetCidrBlock(v string) *DescribeSubnetsResponseBodySubnets {
	s.CidrBlock = &v
	return s
}

func (s *DescribeSubnetsResponseBodySubnets) SetName(v string) *DescribeSubnetsResponseBodySubnets {
	s.Name = &v
	return s
}

func (s *DescribeSubnetsResponseBodySubnets) SetOfficeSiteId(v string) *DescribeSubnetsResponseBodySubnets {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeSubnetsResponseBodySubnets) SetStatus(v string) *DescribeSubnetsResponseBodySubnets {
	s.Status = &v
	return s
}

func (s *DescribeSubnetsResponseBodySubnets) SetSubnetId(v string) *DescribeSubnetsResponseBodySubnets {
	s.SubnetId = &v
	return s
}

func (s *DescribeSubnetsResponseBodySubnets) SetTotalEdsCount(v int32) *DescribeSubnetsResponseBodySubnets {
	s.TotalEdsCount = &v
	return s
}

func (s *DescribeSubnetsResponseBodySubnets) SetTotalEdsCountForGroup(v int32) *DescribeSubnetsResponseBodySubnets {
	s.TotalEdsCountForGroup = &v
	return s
}

func (s *DescribeSubnetsResponseBodySubnets) SetZoneId(v string) *DescribeSubnetsResponseBodySubnets {
	s.ZoneId = &v
	return s
}

func (s *DescribeSubnetsResponseBodySubnets) Validate() error {
	return dara.Validate(s)
}
