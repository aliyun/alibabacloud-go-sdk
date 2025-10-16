// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcListLiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVpcListLiteResponseBody
	GetRequestId() *string
	SetVpcList(v []*DescribeVpcListLiteResponseBodyVpcList) *DescribeVpcListLiteResponseBody
	GetVpcList() []*DescribeVpcListLiteResponseBodyVpcList
}

type DescribeVpcListLiteResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 55E56A55-D93A-5614-AE00-BE2F8077F891
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the VPCs.
	VpcList []*DescribeVpcListLiteResponseBodyVpcList `json:"VpcList,omitempty" xml:"VpcList,omitempty" type:"Repeated"`
}

func (s DescribeVpcListLiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcListLiteResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcListLiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcListLiteResponseBody) GetVpcList() []*DescribeVpcListLiteResponseBodyVpcList {
	return s.VpcList
}

func (s *DescribeVpcListLiteResponseBody) SetRequestId(v string) *DescribeVpcListLiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcListLiteResponseBody) SetVpcList(v []*DescribeVpcListLiteResponseBodyVpcList) *DescribeVpcListLiteResponseBody {
	s.VpcList = v
	return s
}

func (s *DescribeVpcListLiteResponseBody) Validate() error {
	if s.VpcList != nil {
		for _, item := range s.VpcList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcListLiteResponseBodyVpcList struct {
	// The region ID of the VPC.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-wz9dqhljd10fk0b4eh885
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// Cloud_Firewall_VPC
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcListLiteResponseBodyVpcList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcListLiteResponseBodyVpcList) GoString() string {
	return s.String()
}

func (s *DescribeVpcListLiteResponseBodyVpcList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcListLiteResponseBodyVpcList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcListLiteResponseBodyVpcList) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcListLiteResponseBodyVpcList) SetRegionNo(v string) *DescribeVpcListLiteResponseBodyVpcList {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcListLiteResponseBodyVpcList) SetVpcId(v string) *DescribeVpcListLiteResponseBodyVpcList {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcListLiteResponseBodyVpcList) SetVpcName(v string) *DescribeVpcListLiteResponseBodyVpcList {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcListLiteResponseBodyVpcList) Validate() error {
	return dara.Validate(s)
}
