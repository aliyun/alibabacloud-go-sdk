// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSwitchNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeInstanceSwitchNetworkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceSwitchNetworkResponseBody
	GetSuccess() *bool
	SetVpcInfos(v *DescribeInstanceSwitchNetworkResponseBodyVpcInfos) *DescribeInstanceSwitchNetworkResponseBody
	GetVpcInfos() *DescribeInstanceSwitchNetworkResponseBodyVpcInfos
}

type DescribeInstanceSwitchNetworkResponseBody struct {
	// Indicates the ID of the request.
	//
	// example:
	//
	// 03E12FE3-1638-483E-A9B6-1A9120SER56T
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates the information about the virtual private cloud (VPC) in which the instance is deployed.
	VpcInfos *DescribeInstanceSwitchNetworkResponseBodyVpcInfos `json:"VpcInfos,omitempty" xml:"VpcInfos,omitempty" type:"Struct"`
}

func (s DescribeInstanceSwitchNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceSwitchNetworkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceSwitchNetworkResponseBody) GetVpcInfos() *DescribeInstanceSwitchNetworkResponseBodyVpcInfos {
	return s.VpcInfos
}

func (s *DescribeInstanceSwitchNetworkResponseBody) SetRequestId(v string) *DescribeInstanceSwitchNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBody) SetSuccess(v bool) *DescribeInstanceSwitchNetworkResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBody) SetVpcInfos(v *DescribeInstanceSwitchNetworkResponseBodyVpcInfos) *DescribeInstanceSwitchNetworkResponseBody {
	s.VpcInfos = v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceSwitchNetworkResponseBodyVpcInfos struct {
	VpcInfo []*DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo `json:"VpcInfo,omitempty" xml:"VpcInfo,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfos) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfos) GetVpcInfo() []*DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	return s.VpcInfo
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfos) SetVpcInfo(v []*DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) *DescribeInstanceSwitchNetworkResponseBodyVpcInfos {
	s.VpcInfo = v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo struct {
	// Indicates the ID of the region in which the instance is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates the ID of the VPC.
	//
	// example:
	//
	// vpc_id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Indicates the name of the VPC.
	//
	// example:
	//
	// vpc_name
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// Indicates information about the vSwitch to which the instance is connected.
	VswitchInfos *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos `json:"VswitchInfos,omitempty" xml:"VswitchInfos,omitempty" type:"Struct"`
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) GetVswitchInfos() *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos {
	return s.VswitchInfos
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetRegionId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetVpcId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetVpcName(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.VpcName = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetVswitchInfos(v *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.VswitchInfos = v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos struct {
	VswitchInfo []*DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo `json:"VswitchInfo,omitempty" xml:"VswitchInfo,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos) GetVswitchInfo() []*DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	return s.VswitchInfo
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos) SetVswitchInfo(v []*DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos {
	s.VswitchInfo = v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo struct {
	// Indicates the ID of the zone in which the instance is deployed.
	//
	// example:
	//
	// cn-hangzhou-a
	AzoneId *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	// Indicates whether you can change the network type of the instance.
	//
	// example:
	//
	// true
	DrdsSupported *bool `json:"DrdsSupported,omitempty" xml:"DrdsSupported,omitempty"`
	// Indicates the ID of the VPC.
	//
	// example:
	//
	// vpc_id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Indicates the ID of the vSwitch.
	//
	// example:
	//
	// vswitch_id
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// Indicates the name of the vSwitch.
	//
	// example:
	//
	// vswitch_name
	VswitchName *string `json:"VswitchName,omitempty" xml:"VswitchName,omitempty"`
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) GetAzoneId() *string {
	return s.AzoneId
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) GetDrdsSupported() *bool {
	return s.DrdsSupported
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) GetVswitchName() *string {
	return s.VswitchName
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetAzoneId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.AzoneId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetDrdsSupported(v bool) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.DrdsSupported = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetVpcId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetVswitchId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetVswitchName(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.VswitchName = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) Validate() error {
	return dara.Validate(s)
}
