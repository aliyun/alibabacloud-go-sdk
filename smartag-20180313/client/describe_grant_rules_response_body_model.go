// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGrantRules(v *DescribeGrantRulesResponseBodyGrantRules) *DescribeGrantRulesResponseBody
	GetGrantRules() *DescribeGrantRulesResponseBodyGrantRules
	SetPageNumber(v int32) *DescribeGrantRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGrantRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeGrantRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeGrantRulesResponseBody
	GetTotalCount() *int32
}

type DescribeGrantRulesResponseBody struct {
	GrantRules *DescribeGrantRulesResponseBodyGrantRules `json:"GrantRules,omitempty" xml:"GrantRules,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FA579C2D-84A0-4BA1-B9C3-1E5A3B15F1B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGrantRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesResponseBody) GetGrantRules() *DescribeGrantRulesResponseBodyGrantRules {
	return s.GrantRules
}

func (s *DescribeGrantRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGrantRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGrantRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGrantRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeGrantRulesResponseBody) SetGrantRules(v *DescribeGrantRulesResponseBodyGrantRules) *DescribeGrantRulesResponseBody {
	s.GrantRules = v
	return s
}

func (s *DescribeGrantRulesResponseBody) SetPageNumber(v int32) *DescribeGrantRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGrantRulesResponseBody) SetPageSize(v int32) *DescribeGrantRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGrantRulesResponseBody) SetRequestId(v string) *DescribeGrantRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGrantRulesResponseBody) SetTotalCount(v int32) *DescribeGrantRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGrantRulesResponseBody) Validate() error {
	if s.GrantRules != nil {
		if err := s.GrantRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGrantRulesResponseBodyGrantRules struct {
	GrantRule []*DescribeGrantRulesResponseBodyGrantRulesGrantRule `json:"GrantRule,omitempty" xml:"GrantRule,omitempty" type:"Repeated"`
}

func (s DescribeGrantRulesResponseBodyGrantRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesResponseBodyGrantRules) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesResponseBodyGrantRules) GetGrantRule() []*DescribeGrantRulesResponseBodyGrantRulesGrantRule {
	return s.GrantRule
}

func (s *DescribeGrantRulesResponseBodyGrantRules) SetGrantRule(v []*DescribeGrantRulesResponseBodyGrantRulesGrantRule) *DescribeGrantRulesResponseBodyGrantRules {
	s.GrantRule = v
	return s
}

func (s *DescribeGrantRulesResponseBodyGrantRules) Validate() error {
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

type DescribeGrantRulesResponseBodyGrantRulesGrantRule struct {
	CcnInstanceId       *string `json:"CcnInstanceId,omitempty" xml:"CcnInstanceId,omitempty"`
	CcnUid              *int64  `json:"CcnUid,omitempty" xml:"CcnUid,omitempty"`
	CenInstanceId       *string `json:"CenInstanceId,omitempty" xml:"CenInstanceId,omitempty"`
	CenUid              *int64  `json:"CenUid,omitempty" xml:"CenUid,omitempty"`
	GmtCreate           *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified         *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GrantRuleId         *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	GrantTrafficService *bool   `json:"GrantTrafficService,omitempty" xml:"GrantTrafficService,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGrantRulesResponseBodyGrantRulesGrantRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesResponseBodyGrantRulesGrantRule) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) GetCcnInstanceId() *string {
	return s.CcnInstanceId
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) GetCcnUid() *int64 {
	return s.CcnUid
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) GetCenInstanceId() *string {
	return s.CenInstanceId
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) GetCenUid() *int64 {
	return s.CenUid
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) GetGrantRuleId() *string {
	return s.GrantRuleId
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) GetGrantTrafficService() *bool {
	return s.GrantTrafficService
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) SetCcnInstanceId(v string) *DescribeGrantRulesResponseBodyGrantRulesGrantRule {
	s.CcnInstanceId = &v
	return s
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) SetCcnUid(v int64) *DescribeGrantRulesResponseBodyGrantRulesGrantRule {
	s.CcnUid = &v
	return s
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) SetCenInstanceId(v string) *DescribeGrantRulesResponseBodyGrantRulesGrantRule {
	s.CenInstanceId = &v
	return s
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) SetCenUid(v int64) *DescribeGrantRulesResponseBodyGrantRulesGrantRule {
	s.CenUid = &v
	return s
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) SetGmtCreate(v int64) *DescribeGrantRulesResponseBodyGrantRulesGrantRule {
	s.GmtCreate = &v
	return s
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) SetGmtModified(v int64) *DescribeGrantRulesResponseBodyGrantRulesGrantRule {
	s.GmtModified = &v
	return s
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) SetGrantRuleId(v string) *DescribeGrantRulesResponseBodyGrantRulesGrantRule {
	s.GrantRuleId = &v
	return s
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) SetGrantTrafficService(v bool) *DescribeGrantRulesResponseBodyGrantRulesGrantRule {
	s.GrantTrafficService = &v
	return s
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) SetRegionId(v string) *DescribeGrantRulesResponseBodyGrantRulesGrantRule {
	s.RegionId = &v
	return s
}

func (s *DescribeGrantRulesResponseBodyGrantRulesGrantRule) Validate() error {
	return dara.Validate(s)
}
