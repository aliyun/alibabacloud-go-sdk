// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceList(v *DescribeInstanceResponseBodyInstanceList) *DescribeInstanceResponseBody
	GetInstanceList() *DescribeInstanceResponseBodyInstanceList
	SetRequestId(v string) *DescribeInstanceResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeInstanceResponseBody
	GetTotal() *int32
}

type DescribeInstanceResponseBody struct {
	InstanceList *DescribeInstanceResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C728D7E9-9A39-52E0-966B-5C33118BDBB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the assets.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) GetInstanceList() *DescribeInstanceResponseBodyInstanceList {
	return s.InstanceList
}

func (s *DescribeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeInstanceResponseBody) SetInstanceList(v *DescribeInstanceResponseBodyInstanceList) *DescribeInstanceResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetTotal(v int32) *DescribeInstanceResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeInstanceResponseBody) Validate() error {
	if s.InstanceList != nil {
		if err := s.InstanceList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceResponseBodyInstanceList struct {
	Instance []*DescribeInstanceResponseBodyInstanceListInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceList) GetInstance() []*DescribeInstanceResponseBodyInstanceListInstance {
	return s.Instance
}

func (s *DescribeInstanceResponseBodyInstanceList) SetInstance(v []*DescribeInstanceResponseBodyInstanceListInstance) *DescribeInstanceResponseBodyInstanceList {
	s.Instance = v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceList) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceResponseBodyInstanceListInstance struct {
	BlackholeThreshold  *int32  `json:"BlackholeThreshold,omitempty" xml:"BlackholeThreshold,omitempty"`
	DefenseBpsThreshold *int32  `json:"DefenseBpsThreshold,omitempty" xml:"DefenseBpsThreshold,omitempty"`
	DefensePpsThreshold *int32  `json:"DefensePpsThreshold,omitempty" xml:"DefensePpsThreshold,omitempty"`
	ElasticThreshold    *int32  `json:"ElasticThreshold,omitempty" xml:"ElasticThreshold,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIp          *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	InstanceName        *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus      *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType        *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IpVersion           *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	IsBgppack           *bool   `json:"IsBgppack,omitempty" xml:"IsBgppack,omitempty"`
}

func (s DescribeInstanceResponseBodyInstanceListInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceListInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetBlackholeThreshold() *int32 {
	return s.BlackholeThreshold
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetDefenseBpsThreshold() *int32 {
	return s.DefenseBpsThreshold
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetDefensePpsThreshold() *int32 {
	return s.DefensePpsThreshold
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetElasticThreshold() *int32 {
	return s.ElasticThreshold
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetInstanceIp() *string {
	return s.InstanceIp
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) GetIsBgppack() *bool {
	return s.IsBgppack
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetBlackholeThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.BlackholeThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetDefenseBpsThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.DefenseBpsThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetDefensePpsThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.DefensePpsThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetElasticThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.ElasticThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceId(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceIp(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceIp = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceName(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceStatus(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceType(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetIpVersion(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.IpVersion = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetIsBgppack(v bool) *DescribeInstanceResponseBodyInstanceListInstance {
	s.IsBgppack = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) Validate() error {
	return dara.Validate(s)
}
