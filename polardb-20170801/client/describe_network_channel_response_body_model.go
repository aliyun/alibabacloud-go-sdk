// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannelInfos(v []*DescribeNetworkChannelResponseBodyChannelInfos) *DescribeNetworkChannelResponseBody
	GetChannelInfos() []*DescribeNetworkChannelResponseBodyChannelInfos
	SetRequestId(v string) *DescribeNetworkChannelResponseBody
	GetRequestId() *string
}

type DescribeNetworkChannelResponseBody struct {
	ChannelInfos []*DescribeNetworkChannelResponseBodyChannelInfos `json:"ChannelInfos,omitempty" xml:"ChannelInfos,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNetworkChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkChannelResponseBody) GetChannelInfos() []*DescribeNetworkChannelResponseBodyChannelInfos {
	return s.ChannelInfos
}

func (s *DescribeNetworkChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkChannelResponseBody) SetChannelInfos(v []*DescribeNetworkChannelResponseBodyChannelInfos) *DescribeNetworkChannelResponseBody {
	s.ChannelInfos = v
	return s
}

func (s *DescribeNetworkChannelResponseBody) SetRequestId(v string) *DescribeNetworkChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkChannelResponseBody) Validate() error {
	if s.ChannelInfos != nil {
		for _, item := range s.ChannelInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkChannelResponseBodyChannelInfos struct {
	// example:
	//
	// ch4
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// test update
	Notes *string `json:"Notes,omitempty" xml:"Notes,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// pc-*************
	TargetDBClusterId *string `json:"TargetDBClusterId,omitempty" xml:"TargetDBClusterId,omitempty"`
	// example:
	//
	// 33.*.*.240
	TargetIp *string `json:"TargetIp,omitempty" xml:"TargetIp,omitempty"`
	// example:
	//
	// 3389
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// example:
	//
	// polardb_for_postgresql
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// example:
	//
	// vpc-2ze13g2c6j7j2jl*******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNetworkChannelResponseBodyChannelInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkChannelResponseBodyChannelInfos) GoString() string {
	return s.String()
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) GetChannelName() *string {
	return s.ChannelName
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) GetNotes() *string {
	return s.Notes
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) GetTargetDBClusterId() *string {
	return s.TargetDBClusterId
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) GetTargetIp() *string {
	return s.TargetIp
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) GetTargetPort() *string {
	return s.TargetPort
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) SetChannelName(v string) *DescribeNetworkChannelResponseBodyChannelInfos {
	s.ChannelName = &v
	return s
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) SetDBClusterId(v string) *DescribeNetworkChannelResponseBodyChannelInfos {
	s.DBClusterId = &v
	return s
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) SetNotes(v string) *DescribeNetworkChannelResponseBodyChannelInfos {
	s.Notes = &v
	return s
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) SetRegionId(v string) *DescribeNetworkChannelResponseBodyChannelInfos {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) SetTargetDBClusterId(v string) *DescribeNetworkChannelResponseBodyChannelInfos {
	s.TargetDBClusterId = &v
	return s
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) SetTargetIp(v string) *DescribeNetworkChannelResponseBodyChannelInfos {
	s.TargetIp = &v
	return s
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) SetTargetPort(v string) *DescribeNetworkChannelResponseBodyChannelInfos {
	s.TargetPort = &v
	return s
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) SetTargetType(v string) *DescribeNetworkChannelResponseBodyChannelInfos {
	s.TargetType = &v
	return s
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) SetVpcId(v string) *DescribeNetworkChannelResponseBodyChannelInfos {
	s.VpcId = &v
	return s
}

func (s *DescribeNetworkChannelResponseBodyChannelInfos) Validate() error {
	return dara.Validate(s)
}
