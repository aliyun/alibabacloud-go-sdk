// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantSagVbrRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGrantRules(v *DescribeGrantSagVbrRulesResponseBodyGrantRules) *DescribeGrantSagVbrRulesResponseBody
	GetGrantRules() *DescribeGrantSagVbrRulesResponseBodyGrantRules
	SetPageNumber(v int32) *DescribeGrantSagVbrRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGrantSagVbrRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeGrantSagVbrRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeGrantSagVbrRulesResponseBody
	GetTotalCount() *int32
}

type DescribeGrantSagVbrRulesResponseBody struct {
	GrantRules *DescribeGrantSagVbrRulesResponseBodyGrantRules `json:"GrantRules,omitempty" xml:"GrantRules,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 46E98E69-FBA2-423E-9E5A-A3C6D843FED1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of authorization rules.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGrantSagVbrRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantSagVbrRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGrantSagVbrRulesResponseBody) GetGrantRules() *DescribeGrantSagVbrRulesResponseBodyGrantRules {
	return s.GrantRules
}

func (s *DescribeGrantSagVbrRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGrantSagVbrRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGrantSagVbrRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGrantSagVbrRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeGrantSagVbrRulesResponseBody) SetGrantRules(v *DescribeGrantSagVbrRulesResponseBodyGrantRules) *DescribeGrantSagVbrRulesResponseBody {
	s.GrantRules = v
	return s
}

func (s *DescribeGrantSagVbrRulesResponseBody) SetPageNumber(v int32) *DescribeGrantSagVbrRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGrantSagVbrRulesResponseBody) SetPageSize(v int32) *DescribeGrantSagVbrRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGrantSagVbrRulesResponseBody) SetRequestId(v string) *DescribeGrantSagVbrRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGrantSagVbrRulesResponseBody) SetTotalCount(v int32) *DescribeGrantSagVbrRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGrantSagVbrRulesResponseBody) Validate() error {
	if s.GrantRules != nil {
		if err := s.GrantRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGrantSagVbrRulesResponseBodyGrantRules struct {
	GrantRule []*DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule `json:"GrantRule,omitempty" xml:"GrantRule,omitempty" type:"Repeated"`
}

func (s DescribeGrantSagVbrRulesResponseBodyGrantRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantSagVbrRulesResponseBodyGrantRules) GoString() string {
	return s.String()
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRules) GetGrantRule() []*DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule {
	return s.GrantRule
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRules) SetGrantRule(v []*DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) *DescribeGrantSagVbrRulesResponseBodyGrantRules {
	s.GrantRule = v
	return s
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRules) Validate() error {
	if s.GrantRule != nil {
		for _, item := range s.GrantRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule struct {
	Bound         *bool   `json:"Bound,omitempty" xml:"Bound,omitempty"`
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SmartAGId     *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	SmartAGUid    *int64  `json:"SmartAGUid,omitempty" xml:"SmartAGUid,omitempty"`
	VbrInstanceId *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	VbrRegionId   *string `json:"VbrRegionId,omitempty" xml:"VbrRegionId,omitempty"`
	VbrUid        *int64  `json:"VbrUid,omitempty" xml:"VbrUid,omitempty"`
}

func (s DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) GoString() string {
	return s.String()
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) GetBound() *bool {
	return s.Bound
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) GetSmartAGUid() *int64 {
	return s.SmartAGUid
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) GetVbrInstanceId() *string {
	return s.VbrInstanceId
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) GetVbrRegionId() *string {
	return s.VbrRegionId
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) GetVbrUid() *int64 {
	return s.VbrUid
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) SetBound(v bool) *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule {
	s.Bound = &v
	return s
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) SetCreateTime(v int64) *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule {
	s.CreateTime = &v
	return s
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) SetInstanceId(v string) *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule {
	s.InstanceId = &v
	return s
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) SetSmartAGId(v string) *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule {
	s.SmartAGId = &v
	return s
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) SetSmartAGUid(v int64) *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule {
	s.SmartAGUid = &v
	return s
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) SetVbrInstanceId(v string) *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule {
	s.VbrInstanceId = &v
	return s
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) SetVbrRegionId(v string) *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule {
	s.VbrRegionId = &v
	return s
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) SetVbrUid(v int64) *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule {
	s.VbrUid = &v
	return s
}

func (s *DescribeGrantSagVbrRulesResponseBodyGrantRulesGrantRule) Validate() error {
	return dara.Validate(s)
}
