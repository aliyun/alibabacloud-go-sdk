// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventTopAttackAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssets(v []*DescribeRiskEventTopAttackAssetResponseBodyAssets) *DescribeRiskEventTopAttackAssetResponseBody
	GetAssets() []*DescribeRiskEventTopAttackAssetResponseBodyAssets
	SetRequestId(v string) *DescribeRiskEventTopAttackAssetResponseBody
	GetRequestId() *string
}

type DescribeRiskEventTopAttackAssetResponseBody struct {
	// The statistics of the assets.
	Assets []*DescribeRiskEventTopAttackAssetResponseBodyAssets `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B9BF7C33-9A23-5096-8568-A3DACAF0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRiskEventTopAttackAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventTopAttackAssetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventTopAttackAssetResponseBody) GetAssets() []*DescribeRiskEventTopAttackAssetResponseBodyAssets {
	return s.Assets
}

func (s *DescribeRiskEventTopAttackAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskEventTopAttackAssetResponseBody) SetAssets(v []*DescribeRiskEventTopAttackAssetResponseBodyAssets) *DescribeRiskEventTopAttackAssetResponseBody {
	s.Assets = v
	return s
}

func (s *DescribeRiskEventTopAttackAssetResponseBody) SetRequestId(v string) *DescribeRiskEventTopAttackAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetResponseBody) Validate() error {
	if s.Assets != nil {
		for _, item := range s.Assets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskEventTopAttackAssetResponseBodyAssets struct {
	// The number of attacks.
	//
	// example:
	//
	// 12
	AttackCnt *int32 `json:"AttackCnt,omitempty" xml:"AttackCnt,omitempty"`
	// The number of dropped attempts.
	//
	// example:
	//
	// 8
	DropCnt *int32 `json:"DropCnt,omitempty" xml:"DropCnt,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 10.3.54.XXX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-chengdu
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The ID of the asset instance.
	//
	// example:
	//
	// i-8vbdrjrxzt78****
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The name of the resource instance.
	//
	// example:
	//
	// test_resource
	ResourceInstanceName *string `json:"ResourceInstanceName,omitempty" xml:"ResourceInstanceName,omitempty"`
	// The type of the resource instance.
	//
	// example:
	//
	// EcsPublicIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeRiskEventTopAttackAssetResponseBodyAssets) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventTopAttackAssetResponseBodyAssets) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) GetAttackCnt() *int32 {
	return s.AttackCnt
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) GetDropCnt() *int32 {
	return s.DropCnt
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) GetIp() *string {
	return s.Ip
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) GetResourceInstanceName() *string {
	return s.ResourceInstanceName
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) SetAttackCnt(v int32) *DescribeRiskEventTopAttackAssetResponseBodyAssets {
	s.AttackCnt = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) SetDropCnt(v int32) *DescribeRiskEventTopAttackAssetResponseBodyAssets {
	s.DropCnt = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) SetIp(v string) *DescribeRiskEventTopAttackAssetResponseBodyAssets {
	s.Ip = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) SetRegionNo(v string) *DescribeRiskEventTopAttackAssetResponseBodyAssets {
	s.RegionNo = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) SetResourceInstanceId(v string) *DescribeRiskEventTopAttackAssetResponseBodyAssets {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) SetResourceInstanceName(v string) *DescribeRiskEventTopAttackAssetResponseBodyAssets {
	s.ResourceInstanceName = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) SetResourceType(v string) *DescribeRiskEventTopAttackAssetResponseBodyAssets {
	s.ResourceType = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetResponseBodyAssets) Validate() error {
	return dara.Validate(s)
}
