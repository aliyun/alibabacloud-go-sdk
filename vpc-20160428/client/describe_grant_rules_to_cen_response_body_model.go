// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantRulesToCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCenGrantRules(v *DescribeGrantRulesToCenResponseBodyCenGrantRules) *DescribeGrantRulesToCenResponseBody
	GetCenGrantRules() *DescribeGrantRulesToCenResponseBodyCenGrantRules
	SetPageNumber(v int32) *DescribeGrantRulesToCenResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGrantRulesToCenResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeGrantRulesToCenResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeGrantRulesToCenResponseBody
	GetTotalCount() *int32
}

type DescribeGrantRulesToCenResponseBody struct {
	// The information about the authorization.
	CenGrantRules *DescribeGrantRulesToCenResponseBodyCenGrantRules `json:"CenGrantRules,omitempty" xml:"CenGrantRules,omitempty" type:"Struct"`
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
	// F5BB78C8-5F41-464F-B9FF-5E0A7198BA26
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGrantRulesToCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToCenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenResponseBody) GetCenGrantRules() *DescribeGrantRulesToCenResponseBodyCenGrantRules {
	return s.CenGrantRules
}

func (s *DescribeGrantRulesToCenResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGrantRulesToCenResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGrantRulesToCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGrantRulesToCenResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeGrantRulesToCenResponseBody) SetCenGrantRules(v *DescribeGrantRulesToCenResponseBodyCenGrantRules) *DescribeGrantRulesToCenResponseBody {
	s.CenGrantRules = v
	return s
}

func (s *DescribeGrantRulesToCenResponseBody) SetPageNumber(v int32) *DescribeGrantRulesToCenResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBody) SetPageSize(v int32) *DescribeGrantRulesToCenResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBody) SetRequestId(v string) *DescribeGrantRulesToCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBody) SetTotalCount(v int32) *DescribeGrantRulesToCenResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBody) Validate() error {
	if s.CenGrantRules != nil {
		if err := s.CenGrantRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGrantRulesToCenResponseBodyCenGrantRules struct {
	CbnGrantRule []*DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule `json:"CbnGrantRule,omitempty" xml:"CbnGrantRule,omitempty" type:"Repeated"`
}

func (s DescribeGrantRulesToCenResponseBodyCenGrantRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToCenResponseBodyCenGrantRules) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenResponseBodyCenGrantRules) GetCbnGrantRule() []*DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule {
	return s.CbnGrantRule
}

func (s *DescribeGrantRulesToCenResponseBodyCenGrantRules) SetCbnGrantRule(v []*DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule) *DescribeGrantRulesToCenResponseBodyCenGrantRules {
	s.CbnGrantRule = v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyCenGrantRules) Validate() error {
	if s.CbnGrantRule != nil {
		for _, item := range s.CbnGrantRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule struct {
	// The ID of the authorized CEN instance.
	//
	// example:
	//
	// cen-9gsm1q2yh1prpt****
	CenInstanceId *string `json:"CenInstanceId,omitempty" xml:"CenInstanceId,omitempty"`
	// The UID of the Alibaba Cloud account to which the authorized CEN instance belongs.
	//
	// example:
	//
	// 132193271328****
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2019-11-15T09:26:36Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
}

func (s DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule) GetCenInstanceId() *string {
	return s.CenInstanceId
}

func (s *DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule) GetCenOwnerId() *int64 {
	return s.CenOwnerId
}

func (s *DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule) SetCenInstanceId(v string) *DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule {
	s.CenInstanceId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule) SetCenOwnerId(v int64) *DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule {
	s.CenOwnerId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule) SetCreationTime(v string) *DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule {
	s.CreationTime = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyCenGrantRulesCbnGrantRule) Validate() error {
	return dara.Validate(s)
}
