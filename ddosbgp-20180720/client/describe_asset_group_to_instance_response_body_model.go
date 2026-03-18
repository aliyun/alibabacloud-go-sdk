// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetGroupToInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeAssetGroupToInstanceResponseBodyDataList) *DescribeAssetGroupToInstanceResponseBody
	GetDataList() []*DescribeAssetGroupToInstanceResponseBodyDataList
	SetRequestId(v string) *DescribeAssetGroupToInstanceResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeAssetGroupToInstanceResponseBody
	GetTotal() *int64
}

type DescribeAssetGroupToInstanceResponseBody struct {
	// The response parameters.
	DataList []*DescribeAssetGroupToInstanceResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C73C59B9-9F5C-57FF-A394-13EC8FC3B2FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAssetGroupToInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetGroupToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupToInstanceResponseBody) GetDataList() []*DescribeAssetGroupToInstanceResponseBodyDataList {
	return s.DataList
}

func (s *DescribeAssetGroupToInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAssetGroupToInstanceResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeAssetGroupToInstanceResponseBody) SetDataList(v []*DescribeAssetGroupToInstanceResponseBodyDataList) *DescribeAssetGroupToInstanceResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBody) SetRequestId(v string) *DescribeAssetGroupToInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBody) SetTotal(v int64) *DescribeAssetGroupToInstanceResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAssetGroupToInstanceResponseBodyDataList struct {
	// The ID of the Anti-DDoS Origin instance of a paid edition.
	//
	// example:
	//
	// ddosbgp-cn-7212zaa5v***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The UID of the member to which the asset belongs.
	//
	// example:
	//
	// 170858869679****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the asset.
	//
	// example:
	//
	// waf_v2_public_cn-lbj382l****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the asset.
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

func (s DescribeAssetGroupToInstanceResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetGroupToInstanceResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) GetRegion() *string {
	return s.Region
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) GetType() *string {
	return s.Type
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetInstanceId(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetMemberUid(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.MemberUid = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetName(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetRegion(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.Region = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetType(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
