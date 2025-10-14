// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsVswitchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeRdsVswitchesResponseBodyData) *DescribeRdsVswitchesResponseBody
	GetData() *DescribeRdsVswitchesResponseBodyData
	SetRequestId(v string) *DescribeRdsVswitchesResponseBody
	GetRequestId() *string
}

type DescribeRdsVswitchesResponseBody struct {
	Data *DescribeRdsVswitchesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRdsVswitchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVswitchesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsVswitchesResponseBody) GetData() *DescribeRdsVswitchesResponseBodyData {
	return s.Data
}

func (s *DescribeRdsVswitchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRdsVswitchesResponseBody) SetData(v *DescribeRdsVswitchesResponseBodyData) *DescribeRdsVswitchesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRdsVswitchesResponseBody) SetRequestId(v string) *DescribeRdsVswitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsVswitchesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRdsVswitchesResponseBodyData struct {
	VswitchList []*DescribeRdsVswitchesResponseBodyDataVswitchList `json:"VswitchList,omitempty" xml:"VswitchList,omitempty" type:"Repeated"`
}

func (s DescribeRdsVswitchesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVswitchesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRdsVswitchesResponseBodyData) GetVswitchList() []*DescribeRdsVswitchesResponseBodyDataVswitchList {
	return s.VswitchList
}

func (s *DescribeRdsVswitchesResponseBodyData) SetVswitchList(v []*DescribeRdsVswitchesResponseBodyDataVswitchList) *DescribeRdsVswitchesResponseBodyData {
	s.VswitchList = v
	return s
}

func (s *DescribeRdsVswitchesResponseBodyData) Validate() error {
	if s.VswitchList != nil {
		for _, item := range s.VswitchList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRdsVswitchesResponseBodyDataVswitchList struct {
	// example:
	//
	// 16378
	AvailabeIpCount *string `json:"AvailabeIpCount,omitempty" xml:"AvailabeIpCount,omitempty"`
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// example:
	//
	// describe
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// babac91eff704edf9bdccfaa6ba4efce
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// drdshbgae0han226
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// cn-hangzhou-j
	IzNo *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	// example:
	//
	// szt-backup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// vpc id。
	//
	// example:
	//
	// vpc-uf6lis7xmm****t9wrxb5
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeRdsVswitchesResponseBodyDataVswitchList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVswitchesResponseBodyDataVswitchList) GoString() string {
	return s.String()
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) GetAvailabeIpCount() *string {
	return s.AvailabeIpCount
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) GetDescription() *string {
	return s.Description
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) GetId() *int64 {
	return s.Id
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) GetIsDefault() *string {
	return s.IsDefault
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) GetIzNo() *string {
	return s.IzNo
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) GetName() *string {
	return s.Name
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) SetAvailabeIpCount(v string) *DescribeRdsVswitchesResponseBodyDataVswitchList {
	s.AvailabeIpCount = &v
	return s
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) SetCidrBlock(v string) *DescribeRdsVswitchesResponseBodyDataVswitchList {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) SetDescription(v string) *DescribeRdsVswitchesResponseBodyDataVswitchList {
	s.Description = &v
	return s
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) SetId(v int64) *DescribeRdsVswitchesResponseBodyDataVswitchList {
	s.Id = &v
	return s
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) SetInstanceId(v string) *DescribeRdsVswitchesResponseBodyDataVswitchList {
	s.InstanceId = &v
	return s
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) SetIsDefault(v string) *DescribeRdsVswitchesResponseBodyDataVswitchList {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) SetIzNo(v string) *DescribeRdsVswitchesResponseBodyDataVswitchList {
	s.IzNo = &v
	return s
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) SetName(v string) *DescribeRdsVswitchesResponseBodyDataVswitchList {
	s.Name = &v
	return s
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) SetVpcInstanceId(v string) *DescribeRdsVswitchesResponseBodyDataVswitchList {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeRdsVswitchesResponseBodyDataVswitchList) Validate() error {
	return dara.Validate(s)
}
