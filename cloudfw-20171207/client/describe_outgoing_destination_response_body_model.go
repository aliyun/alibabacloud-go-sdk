// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDestinationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDstList(v []*DescribeOutgoingDestinationResponseBodyDstList) *DescribeOutgoingDestinationResponseBody
	GetDstList() []*DescribeOutgoingDestinationResponseBodyDstList
	SetRequestId(v string) *DescribeOutgoingDestinationResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeOutgoingDestinationResponseBody
	GetTotalCount() *int32
}

type DescribeOutgoingDestinationResponseBody struct {
	DstList []*DescribeOutgoingDestinationResponseBodyDstList `json:"DstList,omitempty" xml:"DstList,omitempty" type:"Repeated"`
	// example:
	//
	// A2845BA9-1642-5B27-9F04-8014DD94****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 42
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOutgoingDestinationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationResponseBody) GetDstList() []*DescribeOutgoingDestinationResponseBodyDstList {
	return s.DstList
}

func (s *DescribeOutgoingDestinationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOutgoingDestinationResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeOutgoingDestinationResponseBody) SetDstList(v []*DescribeOutgoingDestinationResponseBodyDstList) *DescribeOutgoingDestinationResponseBody {
	s.DstList = v
	return s
}

func (s *DescribeOutgoingDestinationResponseBody) SetRequestId(v string) *DescribeOutgoingDestinationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBody) SetTotalCount(v int32) *DescribeOutgoingDestinationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBody) Validate() error {
	if s.DstList != nil {
		for _, item := range s.DstList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOutgoingDestinationResponseBodyDstList struct {
	// example:
	//
	// test
	AclRecommendDetail *string `json:"AclRecommendDetail,omitempty" xml:"AclRecommendDetail,omitempty"`
	// example:
	//
	// Normal
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// example:
	//
	// example.com
	Business *string `json:"Business,omitempty" xml:"Business,omitempty"`
	// example:
	//
	// AliYun
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// categor_test
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// example.com
	DstDomain *string `json:"DstDomain,omitempty" xml:"DstDomain,omitempty"`
	// example:
	//
	// 101.6.15.XXX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// example:
	//
	// domain
	DstType *string `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// false
	HasAclRecommend *bool `json:"HasAclRecommend,omitempty" xml:"HasAclRecommend,omitempty"`
	// example:
	//
	// 0.0
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// example:
	//
	// false
	IsMarkNormal *bool `json:"IsMarkNormal,omitempty" xml:"IsMarkNormal,omitempty"`
	// example:
	//
	// 0.0
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// example:
	//
	// 10
	SessionCount *int64                                                   `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	TagList      []*DescribeOutgoingDestinationResponseBodyDstListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
}

func (s DescribeOutgoingDestinationResponseBodyDstList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationResponseBodyDstList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetAclRecommendDetail() *string {
	return s.AclRecommendDetail
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetAclStatus() *string {
	return s.AclStatus
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetBusiness() *string {
	return s.Business
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetCategoryId() *string {
	return s.CategoryId
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetDstDomain() *string {
	return s.DstDomain
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetDstType() *string {
	return s.DstType
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetHasAclRecommend() *bool {
	return s.HasAclRecommend
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetIsMarkNormal() *bool {
	return s.IsMarkNormal
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) GetTagList() []*DescribeOutgoingDestinationResponseBodyDstListTagList {
	return s.TagList
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetAclRecommendDetail(v string) *DescribeOutgoingDestinationResponseBodyDstList {
	s.AclRecommendDetail = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetAclStatus(v string) *DescribeOutgoingDestinationResponseBodyDstList {
	s.AclStatus = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetBusiness(v string) *DescribeOutgoingDestinationResponseBodyDstList {
	s.Business = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetCategoryId(v string) *DescribeOutgoingDestinationResponseBodyDstList {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetCategoryName(v string) *DescribeOutgoingDestinationResponseBodyDstList {
	s.CategoryName = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetDstDomain(v string) *DescribeOutgoingDestinationResponseBodyDstList {
	s.DstDomain = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetDstIP(v string) *DescribeOutgoingDestinationResponseBodyDstList {
	s.DstIP = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetDstType(v string) *DescribeOutgoingDestinationResponseBodyDstList {
	s.DstType = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetGroupName(v string) *DescribeOutgoingDestinationResponseBodyDstList {
	s.GroupName = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetHasAclRecommend(v bool) *DescribeOutgoingDestinationResponseBodyDstList {
	s.HasAclRecommend = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetInBytes(v int64) *DescribeOutgoingDestinationResponseBodyDstList {
	s.InBytes = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetIsMarkNormal(v bool) *DescribeOutgoingDestinationResponseBodyDstList {
	s.IsMarkNormal = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetOutBytes(v int64) *DescribeOutgoingDestinationResponseBodyDstList {
	s.OutBytes = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetSessionCount(v int64) *DescribeOutgoingDestinationResponseBodyDstList {
	s.SessionCount = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) SetTagList(v []*DescribeOutgoingDestinationResponseBodyDstListTagList) *DescribeOutgoingDestinationResponseBodyDstList {
	s.TagList = v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstList) Validate() error {
	if s.TagList != nil {
		for _, item := range s.TagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOutgoingDestinationResponseBodyDstListTagList struct {
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// test
	TagDescribe *string `json:"TagDescribe,omitempty" xml:"TagDescribe,omitempty"`
	// example:
	//
	// FirstFlow
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// verify
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeOutgoingDestinationResponseBodyDstListTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationResponseBodyDstListTagList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationResponseBodyDstListTagList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeOutgoingDestinationResponseBodyDstListTagList) GetTagDescribe() *string {
	return s.TagDescribe
}

func (s *DescribeOutgoingDestinationResponseBodyDstListTagList) GetTagId() *string {
	return s.TagId
}

func (s *DescribeOutgoingDestinationResponseBodyDstListTagList) GetTagName() *string {
	return s.TagName
}

func (s *DescribeOutgoingDestinationResponseBodyDstListTagList) SetRiskLevel(v int32) *DescribeOutgoingDestinationResponseBodyDstListTagList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstListTagList) SetTagDescribe(v string) *DescribeOutgoingDestinationResponseBodyDstListTagList {
	s.TagDescribe = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstListTagList) SetTagId(v string) *DescribeOutgoingDestinationResponseBodyDstListTagList {
	s.TagId = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstListTagList) SetTagName(v string) *DescribeOutgoingDestinationResponseBodyDstListTagList {
	s.TagName = &v
	return s
}

func (s *DescribeOutgoingDestinationResponseBodyDstListTagList) Validate() error {
	return dara.Validate(s)
}
