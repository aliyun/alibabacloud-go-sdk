// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWhiteIpListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestAliyunUid(v string) *WhiteIpListRequest
	GetDestAliyunUid() *string
	SetDestPrimaryVswId(v string) *WhiteIpListRequest
	GetDestPrimaryVswId() *string
	SetDestRoleName(v string) *WhiteIpListRequest
	GetDestRoleName() *string
	SetDestSecondaryVswId(v string) *WhiteIpListRequest
	GetDestSecondaryVswId() *string
	SetDestVpcId(v string) *WhiteIpListRequest
	GetDestVpcId() *string
	SetDestinationRegion(v string) *WhiteIpListRequest
	GetDestinationRegion() *string
	SetRegion(v string) *WhiteIpListRequest
	GetRegion() *string
	SetRegionId(v string) *WhiteIpListRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *WhiteIpListRequest
	GetResourceGroupId() *string
	SetSrcAliyunUid(v string) *WhiteIpListRequest
	GetSrcAliyunUid() *string
	SetSrcPrimaryVswId(v string) *WhiteIpListRequest
	GetSrcPrimaryVswId() *string
	SetSrcRoleName(v string) *WhiteIpListRequest
	GetSrcRoleName() *string
	SetSrcSecondaryVswId(v string) *WhiteIpListRequest
	GetSrcSecondaryVswId() *string
	SetSrcVpcId(v string) *WhiteIpListRequest
	GetSrcVpcId() *string
	SetType(v string) *WhiteIpListRequest
	GetType() *string
	SetZeroEtlJob(v bool) *WhiteIpListRequest
	GetZeroEtlJob() *bool
}

type WhiteIpListRequest struct {
	// The UID of the destination Alibaba Cloud account.
	//
	// example:
	//
	// ****
	DestAliyunUid *string `json:"DestAliyunUid,omitempty" xml:"DestAliyunUid,omitempty"`
	// The primary vSwitch of the destination for VPC NAT.
	//
	// example:
	//
	// ****
	DestPrimaryVswId *string `json:"DestPrimaryVswId,omitempty" xml:"DestPrimaryVswId,omitempty"`
	// The name of the destination role.
	//
	// example:
	//
	// ram-for-dts-sq
	DestRoleName *string `json:"DestRoleName,omitempty" xml:"DestRoleName,omitempty"`
	// The secondary vSwitch of the destination for VPC NAT.
	//
	// example:
	//
	// ****
	DestSecondaryVswId *string `json:"DestSecondaryVswId,omitempty" xml:"DestSecondaryVswId,omitempty"`
	// The ID of the destination VPC.
	//
	// example:
	//
	// ****
	DestVpcId *string `json:"DestVpcId,omitempty" xml:"DestVpcId,omitempty"`
	// The region ID of the destination instance. For details, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// > -  If the destination instance is a self-managed database with a public IP address or a third-party ApsaraDB database, you can set this parameter to **ap-southeast-1*	- or the area ID that is geographically closest to the database.
	//
	// -  This parameter is required when the DTS task is a data migration or data synchronization task.
	//
	// example:
	//
	// cn-hangzhou
	DestinationRegion *string `json:"DestinationRegion,omitempty" xml:"DestinationRegion,omitempty"`
	// The region ID of the source instance. For details, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// > If the source instance is a self-managed database with a public IP address or a third-party ApsaraDB database, you can set this parameter to **ap-southeast-1*	- or the area ID that is geographically closest to the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region to which the DTS instance belongs. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek26mat2ldb4oy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The UID of the source Alibaba Cloud account.
	//
	// example:
	//
	// 1971721963139419
	SrcAliyunUid *string `json:"SrcAliyunUid,omitempty" xml:"SrcAliyunUid,omitempty"`
	// The primary vSwitch of the source for VPC NAT.
	//
	// example:
	//
	// ****
	SrcPrimaryVswId *string `json:"SrcPrimaryVswId,omitempty" xml:"SrcPrimaryVswId,omitempty"`
	// The name of the source role.
	//
	// example:
	//
	// ram-for-dts
	SrcRoleName *string `json:"SrcRoleName,omitempty" xml:"SrcRoleName,omitempty"`
	// The secondary vSwitch of the source for VPC NAT.
	//
	// example:
	//
	// ****
	SrcSecondaryVswId *string `json:"SrcSecondaryVswId,omitempty" xml:"SrcSecondaryVswId,omitempty"`
	// The ID of the source VPC.
	//
	// example:
	//
	// ****
	SrcVpcId *string `json:"SrcVpcId,omitempty" xml:"SrcVpcId,omitempty"`
	// The connection method of the self-managed database or third-party ApsaraDB database. Valid values:
	//
	// - **internet**: connected over the Internet.
	//
	// - **vpc**: connected over Express Connect, VPN Gateway, or Smart Access Gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Specifies whether the node is a seamless integration (Zero-ETL) node. Valid values:
	//
	// - **true**: The node is a seamless integration (Zero-ETL) node.
	//
	// - **false**: The node is not a seamless integration (Zero-ETL) node.
	//
	// example:
	//
	// false
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s WhiteIpListRequest) String() string {
	return dara.Prettify(s)
}

func (s WhiteIpListRequest) GoString() string {
	return s.String()
}

func (s *WhiteIpListRequest) GetDestAliyunUid() *string {
	return s.DestAliyunUid
}

func (s *WhiteIpListRequest) GetDestPrimaryVswId() *string {
	return s.DestPrimaryVswId
}

func (s *WhiteIpListRequest) GetDestRoleName() *string {
	return s.DestRoleName
}

func (s *WhiteIpListRequest) GetDestSecondaryVswId() *string {
	return s.DestSecondaryVswId
}

func (s *WhiteIpListRequest) GetDestVpcId() *string {
	return s.DestVpcId
}

func (s *WhiteIpListRequest) GetDestinationRegion() *string {
	return s.DestinationRegion
}

func (s *WhiteIpListRequest) GetRegion() *string {
	return s.Region
}

func (s *WhiteIpListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *WhiteIpListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *WhiteIpListRequest) GetSrcAliyunUid() *string {
	return s.SrcAliyunUid
}

func (s *WhiteIpListRequest) GetSrcPrimaryVswId() *string {
	return s.SrcPrimaryVswId
}

func (s *WhiteIpListRequest) GetSrcRoleName() *string {
	return s.SrcRoleName
}

func (s *WhiteIpListRequest) GetSrcSecondaryVswId() *string {
	return s.SrcSecondaryVswId
}

func (s *WhiteIpListRequest) GetSrcVpcId() *string {
	return s.SrcVpcId
}

func (s *WhiteIpListRequest) GetType() *string {
	return s.Type
}

func (s *WhiteIpListRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *WhiteIpListRequest) SetDestAliyunUid(v string) *WhiteIpListRequest {
	s.DestAliyunUid = &v
	return s
}

func (s *WhiteIpListRequest) SetDestPrimaryVswId(v string) *WhiteIpListRequest {
	s.DestPrimaryVswId = &v
	return s
}

func (s *WhiteIpListRequest) SetDestRoleName(v string) *WhiteIpListRequest {
	s.DestRoleName = &v
	return s
}

func (s *WhiteIpListRequest) SetDestSecondaryVswId(v string) *WhiteIpListRequest {
	s.DestSecondaryVswId = &v
	return s
}

func (s *WhiteIpListRequest) SetDestVpcId(v string) *WhiteIpListRequest {
	s.DestVpcId = &v
	return s
}

func (s *WhiteIpListRequest) SetDestinationRegion(v string) *WhiteIpListRequest {
	s.DestinationRegion = &v
	return s
}

func (s *WhiteIpListRequest) SetRegion(v string) *WhiteIpListRequest {
	s.Region = &v
	return s
}

func (s *WhiteIpListRequest) SetRegionId(v string) *WhiteIpListRequest {
	s.RegionId = &v
	return s
}

func (s *WhiteIpListRequest) SetResourceGroupId(v string) *WhiteIpListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *WhiteIpListRequest) SetSrcAliyunUid(v string) *WhiteIpListRequest {
	s.SrcAliyunUid = &v
	return s
}

func (s *WhiteIpListRequest) SetSrcPrimaryVswId(v string) *WhiteIpListRequest {
	s.SrcPrimaryVswId = &v
	return s
}

func (s *WhiteIpListRequest) SetSrcRoleName(v string) *WhiteIpListRequest {
	s.SrcRoleName = &v
	return s
}

func (s *WhiteIpListRequest) SetSrcSecondaryVswId(v string) *WhiteIpListRequest {
	s.SrcSecondaryVswId = &v
	return s
}

func (s *WhiteIpListRequest) SetSrcVpcId(v string) *WhiteIpListRequest {
	s.SrcVpcId = &v
	return s
}

func (s *WhiteIpListRequest) SetType(v string) *WhiteIpListRequest {
	s.Type = &v
	return s
}

func (s *WhiteIpListRequest) SetZeroEtlJob(v bool) *WhiteIpListRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *WhiteIpListRequest) Validate() error {
	return dara.Validate(s)
}
