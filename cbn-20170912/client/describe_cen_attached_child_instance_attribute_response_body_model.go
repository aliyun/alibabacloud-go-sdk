// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenAttachedChildInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody
	GetCenId() *string
	SetChildInstanceAttachTime(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody
	GetChildInstanceAttachTime() *string
	SetChildInstanceAttributes(v *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes) *DescribeCenAttachedChildInstanceAttributeResponseBody
	GetChildInstanceAttributes() *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes
	SetChildInstanceId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody
	GetChildInstanceId() *string
	SetChildInstanceName(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody
	GetChildInstanceName() *string
	SetChildInstanceOwnerId(v int64) *DescribeCenAttachedChildInstanceAttributeResponseBody
	GetChildInstanceOwnerId() *int64
	SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody
	GetChildInstanceRegionId() *string
	SetChildInstanceType(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody
	GetChildInstanceType() *string
	SetManagedService(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody
	GetManagedService() *string
	SetRequestId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody
	GetStatus() *string
}

type DescribeCenAttachedChildInstanceAttributeResponseBody struct {
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-5mv960yjhja0dh****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The time when the network instance was attached to the CEN instance.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-07-30T07:53Z
	ChildInstanceAttachTime *string `json:"ChildInstanceAttachTime,omitempty" xml:"ChildInstanceAttachTime,omitempty"`
	// The details about the network instance.
	ChildInstanceAttributes *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes `json:"ChildInstanceAttributes,omitempty" xml:"ChildInstanceAttributes,omitempty" type:"Struct"`
	// The ID of the network instance.
	//
	// example:
	//
	// vpc-2zebdboka7d7t37vo****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The name of the network instance.
	//
	// example:
	//
	// defaultvpc
	ChildInstanceName *string `json:"ChildInstanceName,omitempty" xml:"ChildInstanceName,omitempty"`
	// The ID of the Alibaba Cloud account to which the network instance belongs.
	//
	// example:
	//
	// 1688000000000000
	ChildInstanceOwnerId *int64 `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	// The region ID of the network instance.
	//
	// example:
	//
	// cn-beijing
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**: VPC
	//
	// 	- **VBR**: VBR
	//
	// 	- **CCN**: CCN instance
	//
	// example:
	//
	// VPC
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ManagedService    *string `json:"ManagedService,omitempty" xml:"ManagedService,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ADD98358-D265-4060-87CB-A2427F5A8944
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the network instance is attached to the CEN instance.
	//
	// 	- **Attaching**: The network instance is being attached to the CEN instance.
	//
	// 	- **Attached**: The network instance is attached to the CEN instance.
	//
	// 	- **Detaching**: The network instance is being detached from the CEN instance.
	//
	// example:
	//
	// Attached
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenAttachedChildInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenAttachedChildInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) GetChildInstanceAttachTime() *string {
	return s.ChildInstanceAttachTime
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) GetChildInstanceAttributes() *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes {
	return s.ChildInstanceAttributes
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) GetChildInstanceName() *string {
	return s.ChildInstanceName
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) GetChildInstanceOwnerId() *int64 {
	return s.ChildInstanceOwnerId
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) GetManagedService() *string {
	return s.ManagedService
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetCenId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.CenId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceAttachTime(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceAttachTime = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceAttributes(v *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceAttributes = v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceName(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceName = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceOwnerId(v int64) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceType(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetManagedService(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ManagedService = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetRequestId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetStatus(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes struct {
	// The IPv4 CIDR block of the VPC.
	//
	// example:
	//
	// 192.168.0.0/16
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The IPv6 CIDR block of the VPC.
	//
	// example:
	//
	// 2408:XXXX:0:a600::/56
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	// The IPv6 CIDR blocks of the VPC.
	Ipv6CidrBlocks *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocks `json:"Ipv6CidrBlocks,omitempty" xml:"Ipv6CidrBlocks,omitempty" type:"Struct"`
	// The information about the VPC secondary CIDR block.
	SecondaryCidrBlocks *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesSecondaryCidrBlocks `json:"SecondaryCidrBlocks,omitempty" xml:"SecondaryCidrBlocks,omitempty" type:"Struct"`
}

func (s DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes) GetIpv6CidrBlocks() *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocks {
	return s.Ipv6CidrBlocks
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes) GetSecondaryCidrBlocks() *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesSecondaryCidrBlocks {
	return s.SecondaryCidrBlocks
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes) SetCidrBlock(v string) *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes {
	s.CidrBlock = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes) SetIpv6CidrBlock(v string) *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes) SetIpv6CidrBlocks(v *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocks) *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes {
	s.Ipv6CidrBlocks = v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes) SetSecondaryCidrBlocks(v *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesSecondaryCidrBlocks) *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes {
	s.SecondaryCidrBlocks = v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributes) Validate() error {
	return dara.Validate(s)
}

type DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocks struct {
	Ipv6CidrBlock []*DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocksIpv6CidrBlock `json:"ipv6CidrBlock,omitempty" xml:"ipv6CidrBlock,omitempty" type:"Repeated"`
}

func (s DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocks) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocks) GetIpv6CidrBlock() []*DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocksIpv6CidrBlock {
	return s.Ipv6CidrBlock
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocks) SetIpv6CidrBlock(v []*DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocksIpv6CidrBlock) *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocks {
	s.Ipv6CidrBlock = v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocks) Validate() error {
	return dara.Validate(s)
}

type DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocksIpv6CidrBlock struct {
	// The IPv6 CIDR block of the VPC.
	//
	// example:
	//
	// 2408:XXXX:0:6a::/56
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	// The type of the IPv6 CIDR block of the VPC. Valid values:
	//
	// 	- BGP (default): Alibaba Cloud Border Gateway Protocol (BGP) IPv6
	//
	// 	- ChinaMobile: China Mobile (single line)
	//
	// 	- ChinaUnicom: China Unicom (single line)
	//
	// 	- ChinaTelecom: China Telecom (single line)
	//
	// >  If you are on the whitelist of single-line bandwidth, you can set this parameter to ChinaTelecom, ChinaUnicom, or ChinaMobile.
	//
	// example:
	//
	// BGP
	Ipv6Isp *string `json:"Ipv6Isp,omitempty" xml:"Ipv6Isp,omitempty"`
}

func (s DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocksIpv6CidrBlock) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocksIpv6CidrBlock) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocksIpv6CidrBlock) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocksIpv6CidrBlock) GetIpv6Isp() *string {
	return s.Ipv6Isp
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocksIpv6CidrBlock) SetIpv6CidrBlock(v string) *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocksIpv6CidrBlock {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocksIpv6CidrBlock) SetIpv6Isp(v string) *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocksIpv6CidrBlock {
	s.Ipv6Isp = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesIpv6CidrBlocksIpv6CidrBlock) Validate() error {
	return dara.Validate(s)
}

type DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesSecondaryCidrBlocks struct {
	SecondaryCidrBlock []*string `json:"secondaryCidrBlock,omitempty" xml:"secondaryCidrBlock,omitempty" type:"Repeated"`
}

func (s DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesSecondaryCidrBlocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesSecondaryCidrBlocks) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesSecondaryCidrBlocks) GetSecondaryCidrBlock() []*string {
	return s.SecondaryCidrBlock
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesSecondaryCidrBlocks) SetSecondaryCidrBlock(v []*string) *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesSecondaryCidrBlocks {
	s.SecondaryCidrBlock = v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBodyChildInstanceAttributesSecondaryCidrBlocks) Validate() error {
	return dara.Validate(s)
}
