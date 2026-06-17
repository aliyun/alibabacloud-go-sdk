// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRateLimitPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeRateLimitPolicyResponseBodyItems) *DescribeRateLimitPolicyResponseBody
	GetItems() []*DescribeRateLimitPolicyResponseBodyItems
	SetPageNumber(v string) *DescribeRateLimitPolicyResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *DescribeRateLimitPolicyResponseBody
	GetPageRecordCount() *string
	SetPageSize(v string) *DescribeRateLimitPolicyResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeRateLimitPolicyResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribeRateLimitPolicyResponseBody
	GetTotalRecordCount() *string
}

type DescribeRateLimitPolicyResponseBody struct {
	// An array of rate limit policy objects.
	Items []*DescribeRateLimitPolicyResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The returned page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CED079B7-A408-41A1-BFF1-EC608E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of matching entries.
	//
	// example:
	//
	// 1
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeRateLimitPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRateLimitPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRateLimitPolicyResponseBody) GetItems() []*DescribeRateLimitPolicyResponseBodyItems {
	return s.Items
}

func (s *DescribeRateLimitPolicyResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeRateLimitPolicyResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribeRateLimitPolicyResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeRateLimitPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRateLimitPolicyResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribeRateLimitPolicyResponseBody) SetItems(v []*DescribeRateLimitPolicyResponseBodyItems) *DescribeRateLimitPolicyResponseBody {
	s.Items = v
	return s
}

func (s *DescribeRateLimitPolicyResponseBody) SetPageNumber(v string) *DescribeRateLimitPolicyResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBody) SetPageRecordCount(v string) *DescribeRateLimitPolicyResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBody) SetPageSize(v string) *DescribeRateLimitPolicyResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBody) SetRequestId(v string) *DescribeRateLimitPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBody) SetTotalRecordCount(v string) *DescribeRateLimitPolicyResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRateLimitPolicyResponseBodyItems struct {
	// The creation time.
	//
	// example:
	//
	// 2025-12-01T17:52:05+08:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-01-19T16:47:25+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the gateway instance.
	//
	// example:
	//
	// pg-xxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The ID of the rate limit policy.
	//
	// example:
	//
	// 02eccf7c61cf4d05a543075ee907f3**
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The policy type.
	//
	// example:
	//
	// RateLimit
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The maximum requests per minute (RPM).
	//
	// example:
	//
	// 10
	RateLimitRpm *string `json:"RateLimitRpm,omitempty" xml:"RateLimitRpm,omitempty"`
	// The maximum tokens per minute (TPM).
	//
	// example:
	//
	// 10
	RateLimitTpm *string `json:"RateLimitTpm,omitempty" xml:"RateLimitTpm,omitempty"`
	// The ID of the object within the policy\\"s scope, such as a consumer group or a consumer.
	//
	// example:
	//
	// cg-xxxxxxx
	ScopeRefId *string `json:"ScopeRefId,omitempty" xml:"ScopeRefId,omitempty"`
	// The scope of the policy. Valid values:
	//
	// - **ConsumerGroup**: The policy applies to a consumer group.
	//
	// - **Consumer**: The policy applies to a specific consumer.
	//
	// example:
	//
	// ConsumerGroup
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The status of the policy. Valid values:
	//
	// - **Enabled**: The policy is enabled.
	//
	// - **Disabled**: The policy is disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRateLimitPolicyResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeRateLimitPolicyResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeRateLimitPolicyResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeRateLimitPolicyResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeRateLimitPolicyResponseBodyItems) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeRateLimitPolicyResponseBodyItems) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeRateLimitPolicyResponseBodyItems) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeRateLimitPolicyResponseBodyItems) GetRateLimitRpm() *string {
	return s.RateLimitRpm
}

func (s *DescribeRateLimitPolicyResponseBodyItems) GetRateLimitTpm() *string {
	return s.RateLimitTpm
}

func (s *DescribeRateLimitPolicyResponseBodyItems) GetScopeRefId() *string {
	return s.ScopeRefId
}

func (s *DescribeRateLimitPolicyResponseBodyItems) GetScopeType() *string {
	return s.ScopeType
}

func (s *DescribeRateLimitPolicyResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeRateLimitPolicyResponseBodyItems) SetGmtCreated(v string) *DescribeRateLimitPolicyResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBodyItems) SetGmtModified(v string) *DescribeRateLimitPolicyResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBodyItems) SetGwClusterId(v string) *DescribeRateLimitPolicyResponseBodyItems {
	s.GwClusterId = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBodyItems) SetPolicyId(v string) *DescribeRateLimitPolicyResponseBodyItems {
	s.PolicyId = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBodyItems) SetPolicyType(v string) *DescribeRateLimitPolicyResponseBodyItems {
	s.PolicyType = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBodyItems) SetRateLimitRpm(v string) *DescribeRateLimitPolicyResponseBodyItems {
	s.RateLimitRpm = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBodyItems) SetRateLimitTpm(v string) *DescribeRateLimitPolicyResponseBodyItems {
	s.RateLimitTpm = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBodyItems) SetScopeRefId(v string) *DescribeRateLimitPolicyResponseBodyItems {
	s.ScopeRefId = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBodyItems) SetScopeType(v string) *DescribeRateLimitPolicyResponseBodyItems {
	s.ScopeType = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBodyItems) SetStatus(v string) *DescribeRateLimitPolicyResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeRateLimitPolicyResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
