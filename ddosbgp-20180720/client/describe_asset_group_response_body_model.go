// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetGroupList(v []*DescribeAssetGroupResponseBodyAssetGroupList) *DescribeAssetGroupResponseBody
	GetAssetGroupList() []*DescribeAssetGroupResponseBodyAssetGroupList
	SetRequestId(v string) *DescribeAssetGroupResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeAssetGroupResponseBody
	GetTotal() *int64
}

type DescribeAssetGroupResponseBody struct {
	// The information about the asset.
	AssetGroupList []*DescribeAssetGroupResponseBodyAssetGroupList `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 487EC0F7-8D14-504E-914E-3A1BC314B581
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAssetGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupResponseBody) GetAssetGroupList() []*DescribeAssetGroupResponseBodyAssetGroupList {
	return s.AssetGroupList
}

func (s *DescribeAssetGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAssetGroupResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeAssetGroupResponseBody) SetAssetGroupList(v []*DescribeAssetGroupResponseBodyAssetGroupList) *DescribeAssetGroupResponseBody {
	s.AssetGroupList = v
	return s
}

func (s *DescribeAssetGroupResponseBody) SetRequestId(v string) *DescribeAssetGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetGroupResponseBody) SetTotal(v int64) *DescribeAssetGroupResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeAssetGroupResponseBody) Validate() error {
	if s.AssetGroupList != nil {
		for _, item := range s.AssetGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAssetGroupResponseBodyAssetGroupList struct {
	// The ID of the asset.
	//
	// example:
	//
	// waf_v2_public_cn-lbj382l****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region to which the asset belongs.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The type of the asset.
	//
	// example:
	//
	// waf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAssetGroupResponseBodyAssetGroupList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetGroupResponseBodyAssetGroupList) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupResponseBodyAssetGroupList) GetName() *string {
	return s.Name
}

func (s *DescribeAssetGroupResponseBodyAssetGroupList) GetRegion() *string {
	return s.Region
}

func (s *DescribeAssetGroupResponseBodyAssetGroupList) GetType() *string {
	return s.Type
}

func (s *DescribeAssetGroupResponseBodyAssetGroupList) SetName(v string) *DescribeAssetGroupResponseBodyAssetGroupList {
	s.Name = &v
	return s
}

func (s *DescribeAssetGroupResponseBodyAssetGroupList) SetRegion(v string) *DescribeAssetGroupResponseBodyAssetGroupList {
	s.Region = &v
	return s
}

func (s *DescribeAssetGroupResponseBodyAssetGroupList) SetType(v string) *DescribeAssetGroupResponseBodyAssetGroupList {
	s.Type = &v
	return s
}

func (s *DescribeAssetGroupResponseBodyAssetGroupList) Validate() error {
	return dara.Validate(s)
}
