// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAntiBruteForceRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) *DescribeInstanceAntiBruteForceRulesResponseBody
	GetPageInfo() *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo
	SetRequestId(v string) *DescribeInstanceAntiBruteForceRulesResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeInstanceAntiBruteForceRulesResponseBodyRules) *DescribeInstanceAntiBruteForceRulesResponseBody
	GetRules() []*DescribeInstanceAntiBruteForceRulesResponseBodyRules
}

type DescribeInstanceAntiBruteForceRulesResponseBody struct {
	// The pagination information.
	PageInfo *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 97286A-4A6B-4A4-95FA-EC7E3E2451
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array consisting of the servers to which a defense rule is applied.
	Rules []*DescribeInstanceAntiBruteForceRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAntiBruteForceRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAntiBruteForceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBody) GetPageInfo() *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBody) GetRules() []*DescribeInstanceAntiBruteForceRulesResponseBodyRules {
	return s.Rules
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBody) SetPageInfo(v *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) *DescribeInstanceAntiBruteForceRulesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBody) SetRequestId(v string) *DescribeInstanceAntiBruteForceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBody) SetRules(v []*DescribeInstanceAntiBruteForceRulesResponseBodyRules) *DescribeInstanceAntiBruteForceRulesResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo struct {
	// The number of servers returned on the current page.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of servers returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) SetCount(v int32) *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) SetPageSize(v int32) *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAntiBruteForceRulesResponseBodyRules struct {
	// The ID of the defense rule.
	//
	// example:
	//
	// 215779601
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the defense rule.
	//
	// example:
	//
	// TestRule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The UUID of the server to which the defense rule is applied.
	//
	// example:
	//
	// 4fe8e1cd-3c37-4851-b9de-124da32c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeInstanceAntiBruteForceRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAntiBruteForceRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyRules) GetId() *int64 {
	return s.Id
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyRules) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyRules) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyRules) SetId(v int64) *DescribeInstanceAntiBruteForceRulesResponseBodyRules {
	s.Id = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyRules) SetName(v string) *DescribeInstanceAntiBruteForceRulesResponseBodyRules {
	s.Name = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyRules) SetUuid(v string) *DescribeInstanceAntiBruteForceRulesResponseBodyRules {
	s.Uuid = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
