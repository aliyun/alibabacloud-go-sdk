// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantSagRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGrantRules(v *DescribeGrantSagRulesResponseBodyGrantRules) *DescribeGrantSagRulesResponseBody
	GetGrantRules() *DescribeGrantSagRulesResponseBodyGrantRules
	SetPageNumber(v int32) *DescribeGrantSagRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGrantSagRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeGrantSagRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeGrantSagRulesResponseBody
	GetTotalCount() *int32
}

type DescribeGrantSagRulesResponseBody struct {
	GrantRules *DescribeGrantSagRulesResponseBodyGrantRules `json:"GrantRules,omitempty" xml:"GrantRules,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6E1674AC-083C-4031-B047-7A66E418E0C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGrantSagRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantSagRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGrantSagRulesResponseBody) GetGrantRules() *DescribeGrantSagRulesResponseBodyGrantRules {
	return s.GrantRules
}

func (s *DescribeGrantSagRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGrantSagRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGrantSagRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGrantSagRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeGrantSagRulesResponseBody) SetGrantRules(v *DescribeGrantSagRulesResponseBodyGrantRules) *DescribeGrantSagRulesResponseBody {
	s.GrantRules = v
	return s
}

func (s *DescribeGrantSagRulesResponseBody) SetPageNumber(v int32) *DescribeGrantSagRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGrantSagRulesResponseBody) SetPageSize(v int32) *DescribeGrantSagRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGrantSagRulesResponseBody) SetRequestId(v string) *DescribeGrantSagRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGrantSagRulesResponseBody) SetTotalCount(v int32) *DescribeGrantSagRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGrantSagRulesResponseBody) Validate() error {
	if s.GrantRules != nil {
		if err := s.GrantRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGrantSagRulesResponseBodyGrantRules struct {
	GrantRule []*DescribeGrantSagRulesResponseBodyGrantRulesGrantRule `json:"GrantRule,omitempty" xml:"GrantRule,omitempty" type:"Repeated"`
}

func (s DescribeGrantSagRulesResponseBodyGrantRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantSagRulesResponseBodyGrantRules) GoString() string {
	return s.String()
}

func (s *DescribeGrantSagRulesResponseBodyGrantRules) GetGrantRule() []*DescribeGrantSagRulesResponseBodyGrantRulesGrantRule {
	return s.GrantRule
}

func (s *DescribeGrantSagRulesResponseBodyGrantRules) SetGrantRule(v []*DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) *DescribeGrantSagRulesResponseBodyGrantRules {
	s.GrantRule = v
	return s
}

func (s *DescribeGrantSagRulesResponseBodyGrantRules) Validate() error {
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

type DescribeGrantSagRulesResponseBodyGrantRulesGrantRule struct {
	CcnInstanceId       *string `json:"CcnInstanceId,omitempty" xml:"CcnInstanceId,omitempty"`
	CcnUid              *int64  `json:"CcnUid,omitempty" xml:"CcnUid,omitempty"`
	CreateTime          *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GrantTrafficService *bool   `json:"GrantTrafficService,omitempty" xml:"GrantTrafficService,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SmartAGId           *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) GoString() string {
	return s.String()
}

func (s *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) GetCcnInstanceId() *string {
	return s.CcnInstanceId
}

func (s *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) GetCcnUid() *int64 {
	return s.CcnUid
}

func (s *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) GetGrantTrafficService() *bool {
	return s.GrantTrafficService
}

func (s *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) SetCcnInstanceId(v string) *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule {
	s.CcnInstanceId = &v
	return s
}

func (s *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) SetCcnUid(v int64) *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule {
	s.CcnUid = &v
	return s
}

func (s *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) SetCreateTime(v int64) *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule {
	s.CreateTime = &v
	return s
}

func (s *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) SetGrantTrafficService(v bool) *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule {
	s.GrantTrafficService = &v
	return s
}

func (s *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) SetInstanceId(v string) *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule {
	s.InstanceId = &v
	return s
}

func (s *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) SetSmartAGId(v string) *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule {
	s.SmartAGId = &v
	return s
}

func (s *DescribeGrantSagRulesResponseBodyGrantRulesGrantRule) Validate() error {
	return dara.Validate(s)
}
