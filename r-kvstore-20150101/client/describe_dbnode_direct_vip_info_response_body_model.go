// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBNodeDirectVipInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDirectVipInfo(v *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfo) *DescribeDBNodeDirectVipInfoResponseBody
	GetDirectVipInfo() *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfo
	SetInstanceId(v string) *DescribeDBNodeDirectVipInfoResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *DescribeDBNodeDirectVipInfoResponseBody
	GetRequestId() *string
}

type DescribeDBNodeDirectVipInfoResponseBody struct {
	// The VIP information of shards in the cluster instance.
	DirectVipInfo *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfo `json:"DirectVipInfo,omitempty" xml:"DirectVipInfo,omitempty" type:"Struct"`
	// The instance ID.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABAF95F6-35C1-4177-AF3A-70969EBD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBNodeDirectVipInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodeDirectVipInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBNodeDirectVipInfoResponseBody) GetDirectVipInfo() *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfo {
	return s.DirectVipInfo
}

func (s *DescribeDBNodeDirectVipInfoResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDBNodeDirectVipInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBNodeDirectVipInfoResponseBody) SetDirectVipInfo(v *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfo) *DescribeDBNodeDirectVipInfoResponseBody {
	s.DirectVipInfo = v
	return s
}

func (s *DescribeDBNodeDirectVipInfoResponseBody) SetInstanceId(v string) *DescribeDBNodeDirectVipInfoResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDBNodeDirectVipInfoResponseBody) SetRequestId(v string) *DescribeDBNodeDirectVipInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBNodeDirectVipInfoResponseBody) Validate() error {
	if s.DirectVipInfo != nil {
		if err := s.DirectVipInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfo struct {
	VipInfo []*DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo `json:"VipInfo,omitempty" xml:"VipInfo,omitempty" type:"Repeated"`
}

func (s DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfo) GetVipInfo() []*DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo {
	return s.VipInfo
}

func (s *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfo) SetVipInfo(v []*DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo) *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfo {
	s.VipInfo = v
	return s
}

func (s *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfo) Validate() error {
	if s.VipInfo != nil {
		for _, item := range s.VipInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo struct {
	// The network type of the security group. Valid values:
	//
	// 	- **vpc**: Virtual Private Cloud (VPC)
	//
	// example:
	//
	// vpc
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The shard ID.
	//
	// example:
	//
	// r-8vb3679b04551444-db-2
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The port number. Valid values: **1024*	- to **65535**. Default value: **6379**.
	//
	// example:
	//
	// 6379
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The VIP of the shard.
	//
	// example:
	//
	// 100.115.61.8
	Vip *string `json:"Vip,omitempty" xml:"Vip,omitempty"`
}

func (s DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo) GetNetType() *string {
	return s.NetType
}

func (s *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo) GetPort() *string {
	return s.Port
}

func (s *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo) GetVip() *string {
	return s.Vip
}

func (s *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo) SetNetType(v string) *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo {
	s.NetType = &v
	return s
}

func (s *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo) SetNodeId(v string) *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo) SetPort(v string) *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo {
	s.Port = &v
	return s
}

func (s *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo) SetVip(v string) *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo {
	s.Vip = &v
	return s
}

func (s *DescribeDBNodeDirectVipInfoResponseBodyDirectVipInfoVipInfo) Validate() error {
	return dara.Validate(s)
}
