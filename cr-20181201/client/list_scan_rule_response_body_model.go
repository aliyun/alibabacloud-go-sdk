// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScanRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScanRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListScanRuleResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListScanRuleResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListScanRuleResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListScanRuleResponseBody
	GetRequestId() *string
	SetScanRules(v []*ListScanRuleResponseBodyScanRules) *ListScanRuleResponseBody
	GetScanRules() []*ListScanRuleResponseBodyScanRules
	SetTotalCount(v int32) *ListScanRuleResponseBody
	GetTotalCount() *int32
}

type ListScanRuleResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request Id
	//
	// example:
	//
	// 2CB62B5E-605B-5A23-9110-728B8207A25C
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScanRules []*ListScanRuleResponseBodyScanRules `json:"ScanRules,omitempty" xml:"ScanRules,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScanRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScanRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListScanRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScanRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListScanRuleResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListScanRuleResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScanRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScanRuleResponseBody) GetScanRules() []*ListScanRuleResponseBodyScanRules {
	return s.ScanRules
}

func (s *ListScanRuleResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListScanRuleResponseBody) SetCode(v string) *ListScanRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListScanRuleResponseBody) SetIsSuccess(v bool) *ListScanRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListScanRuleResponseBody) SetPageNo(v int32) *ListScanRuleResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListScanRuleResponseBody) SetPageSize(v int32) *ListScanRuleResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScanRuleResponseBody) SetRequestId(v string) *ListScanRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScanRuleResponseBody) SetScanRules(v []*ListScanRuleResponseBodyScanRules) *ListScanRuleResponseBody {
	s.ScanRules = v
	return s
}

func (s *ListScanRuleResponseBody) SetTotalCount(v int32) *ListScanRuleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScanRuleResponseBody) Validate() error {
	if s.ScanRules != nil {
		for _, item := range s.ScanRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScanRuleResponseBodyScanRules struct {
	// example:
	//
	// 1702361810000
	CreateTime *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	RepoNames  []*string `json:"RepoNames,omitempty" xml:"RepoNames,omitempty" type:"Repeated"`
	// example:
	//
	// .*
	RepoTagFilterPattern *string `json:"RepoTagFilterPattern,omitempty" xml:"RepoTagFilterPattern,omitempty"`
	// example:
	//
	// cicd-prod
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// crscnr-2sdveqjhpzd****
	ScanRuleId *string `json:"ScanRuleId,omitempty" xml:"ScanRuleId,omitempty"`
	// example:
	//
	// REPO
	ScanScope *string `json:"ScanScope,omitempty" xml:"ScanScope,omitempty"`
	// example:
	//
	// SBOM
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// example:
	//
	// AUTO
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// example:
	//
	// 1764122725000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListScanRuleResponseBodyScanRules) String() string {
	return dara.Prettify(s)
}

func (s ListScanRuleResponseBodyScanRules) GoString() string {
	return s.String()
}

func (s *ListScanRuleResponseBodyScanRules) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListScanRuleResponseBodyScanRules) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *ListScanRuleResponseBodyScanRules) GetRepoNames() []*string {
	return s.RepoNames
}

func (s *ListScanRuleResponseBodyScanRules) GetRepoTagFilterPattern() *string {
	return s.RepoTagFilterPattern
}

func (s *ListScanRuleResponseBodyScanRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListScanRuleResponseBodyScanRules) GetScanRuleId() *string {
	return s.ScanRuleId
}

func (s *ListScanRuleResponseBodyScanRules) GetScanScope() *string {
	return s.ScanScope
}

func (s *ListScanRuleResponseBodyScanRules) GetScanType() *string {
	return s.ScanType
}

func (s *ListScanRuleResponseBodyScanRules) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListScanRuleResponseBodyScanRules) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListScanRuleResponseBodyScanRules) SetCreateTime(v int64) *ListScanRuleResponseBodyScanRules {
	s.CreateTime = &v
	return s
}

func (s *ListScanRuleResponseBodyScanRules) SetNamespaces(v []*string) *ListScanRuleResponseBodyScanRules {
	s.Namespaces = v
	return s
}

func (s *ListScanRuleResponseBodyScanRules) SetRepoNames(v []*string) *ListScanRuleResponseBodyScanRules {
	s.RepoNames = v
	return s
}

func (s *ListScanRuleResponseBodyScanRules) SetRepoTagFilterPattern(v string) *ListScanRuleResponseBodyScanRules {
	s.RepoTagFilterPattern = &v
	return s
}

func (s *ListScanRuleResponseBodyScanRules) SetRuleName(v string) *ListScanRuleResponseBodyScanRules {
	s.RuleName = &v
	return s
}

func (s *ListScanRuleResponseBodyScanRules) SetScanRuleId(v string) *ListScanRuleResponseBodyScanRules {
	s.ScanRuleId = &v
	return s
}

func (s *ListScanRuleResponseBodyScanRules) SetScanScope(v string) *ListScanRuleResponseBodyScanRules {
	s.ScanScope = &v
	return s
}

func (s *ListScanRuleResponseBodyScanRules) SetScanType(v string) *ListScanRuleResponseBodyScanRules {
	s.ScanType = &v
	return s
}

func (s *ListScanRuleResponseBodyScanRules) SetTriggerType(v string) *ListScanRuleResponseBodyScanRules {
	s.TriggerType = &v
	return s
}

func (s *ListScanRuleResponseBodyScanRules) SetUpdateTime(v int64) *ListScanRuleResponseBodyScanRules {
	s.UpdateTime = &v
	return s
}

func (s *ListScanRuleResponseBodyScanRules) Validate() error {
	return dara.Validate(s)
}
