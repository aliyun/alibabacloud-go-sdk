// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingAssetListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetList(v []*DescribeOutgoingAssetListResponseBodyAssetList) *DescribeOutgoingAssetListResponseBody
	GetAssetList() []*DescribeOutgoingAssetListResponseBodyAssetList
	SetRequestId(v string) *DescribeOutgoingAssetListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeOutgoingAssetListResponseBody
	GetTotalCount() *int32
}

type DescribeOutgoingAssetListResponseBody struct {
	AssetList []*DescribeOutgoingAssetListResponseBodyAssetList `json:"AssetList,omitempty" xml:"AssetList,omitempty" type:"Repeated"`
	// example:
	//
	// 7A515672-FAAE-584F-B51C-B2586E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOutgoingAssetListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingAssetListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingAssetListResponseBody) GetAssetList() []*DescribeOutgoingAssetListResponseBodyAssetList {
	return s.AssetList
}

func (s *DescribeOutgoingAssetListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOutgoingAssetListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeOutgoingAssetListResponseBody) SetAssetList(v []*DescribeOutgoingAssetListResponseBodyAssetList) *DescribeOutgoingAssetListResponseBody {
	s.AssetList = v
	return s
}

func (s *DescribeOutgoingAssetListResponseBody) SetRequestId(v string) *DescribeOutgoingAssetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBody) SetTotalCount(v int32) *DescribeOutgoingAssetListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBody) Validate() error {
	if s.AssetList != nil {
		for _, item := range s.AssetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOutgoingAssetListResponseBodyAssetList struct {
	// example:
	//
	// i-8vbcmllue4d94nto****
	AssetInstanceId *string `json:"AssetInstanceId,omitempty" xml:"AssetInstanceId,omitempty"`
	// example:
	//
	// test
	AssetInstanceName *string `json:"AssetInstanceName,omitempty" xml:"AssetInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	AssetsRegion *string `json:"AssetsRegion,omitempty" xml:"AssetsRegion,omitempty"`
	// example:
	//
	// subscribe
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 244438.0
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// example:
	//
	// 12
	IpsHitCnt *int64 `json:"IpsHitCnt,omitempty" xml:"IpsHitCnt,omitempty"`
	// example:
	//
	// ngw-bp1utx6wj4x9qu9tl****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// example:
	//
	// ngw-test
	NatGatewayName *string `json:"NatGatewayName,omitempty" xml:"NatGatewayName,omitempty"`
	// example:
	//
	// 100
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// example:
	//
	// 10
	OutgoingDomainCnt *int64 `json:"OutgoingDomainCnt,omitempty" xml:"OutgoingDomainCnt,omitempty"`
	// example:
	//
	// 12
	OutgoingDstIPCnt *int64 `json:"OutgoingDstIPCnt,omitempty" xml:"OutgoingDstIPCnt,omitempty"`
	// example:
	//
	// 10.21.242XXX
	PrivateIP     *string   `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	PrivateIPList []*string `json:"PrivateIPList,omitempty" xml:"PrivateIPList,omitempty" type:"Repeated"`
	// example:
	//
	// 47.96.181.XXX
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// example:
	//
	// EcsPublicIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// risk
	SecurityRisk *string `json:"SecurityRisk,omitempty" xml:"SecurityRisk,omitempty"`
	// example:
	//
	// 2
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// example:
	//
	// 12498767
	TotalBytes *int64 `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
}

func (s DescribeOutgoingAssetListResponseBodyAssetList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingAssetListResponseBodyAssetList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetAssetInstanceId() *string {
	return s.AssetInstanceId
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetAssetInstanceName() *string {
	return s.AssetInstanceName
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetAssetsRegion() *string {
	return s.AssetsRegion
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetIpsHitCnt() *int64 {
	return s.IpsHitCnt
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetNatGatewayName() *string {
	return s.NatGatewayName
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetOutgoingDomainCnt() *int64 {
	return s.OutgoingDomainCnt
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetOutgoingDstIPCnt() *int64 {
	return s.OutgoingDstIPCnt
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetPrivateIP() *string {
	return s.PrivateIP
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetPrivateIPList() []*string {
	return s.PrivateIPList
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetPublicIP() *string {
	return s.PublicIP
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetSecurityRisk() *string {
	return s.SecurityRisk
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetAssetInstanceId(v string) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.AssetInstanceId = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetAssetInstanceName(v string) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.AssetInstanceName = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetAssetsRegion(v string) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.AssetsRegion = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetGroupName(v string) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.GroupName = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetInBytes(v int64) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.InBytes = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetIpsHitCnt(v int64) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.IpsHitCnt = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetNatGatewayId(v string) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetNatGatewayName(v string) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.NatGatewayName = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetOutBytes(v int64) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.OutBytes = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetOutgoingDomainCnt(v int64) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.OutgoingDomainCnt = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetOutgoingDstIPCnt(v int64) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.OutgoingDstIPCnt = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetPrivateIP(v string) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.PrivateIP = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetPrivateIPList(v []*string) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.PrivateIPList = v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetPublicIP(v string) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.PublicIP = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetResourceType(v string) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.ResourceType = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetSecurityRisk(v string) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.SecurityRisk = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetSessionCount(v int64) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.SessionCount = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) SetTotalBytes(v int64) *DescribeOutgoingAssetListResponseBodyAssetList {
	s.TotalBytes = &v
	return s
}

func (s *DescribeOutgoingAssetListResponseBodyAssetList) Validate() error {
	return dara.Validate(s)
}
