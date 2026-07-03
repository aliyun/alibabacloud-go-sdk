// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomizeRuleTestResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListCustomizeRuleTestResultRequest
	GetCurrentPage() *int32
	SetDetectionRuleId(v string) *ListCustomizeRuleTestResultRequest
	GetDetectionRuleId() *string
	SetEndTime(v int64) *ListCustomizeRuleTestResultRequest
	GetEndTime() *int64
	SetId(v int64) *ListCustomizeRuleTestResultRequest
	GetId() *int64
	SetPageSize(v int32) *ListCustomizeRuleTestResultRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListCustomizeRuleTestResultRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListCustomizeRuleTestResultRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListCustomizeRuleTestResultRequest
	GetRoleType() *int32
	SetStartTime(v int64) *ListCustomizeRuleTestResultRequest
	GetStartTime() *int64
	SetVerifyType(v string) *ListCustomizeRuleTestResultRequest
	GetVerifyType() *string
}

type ListCustomizeRuleTestResultRequest struct {
	// The page number. The value must be greater than or equal to 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the custom rule. You can obtain the rule ID from the rule list.
	//
	// example:
	//
	// dr-53np4nguf5jmh1vc****
	DetectionRuleId *string `json:"DetectionRuleId,omitempty" xml:"DetectionRuleId,omitempty"`
	// The end time.
	//
	// example:
	//
	// 1731797891000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the custom rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of entries per page. The maximum value is 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the Data Management center of Threat Analysis is located. Select a region based on the region where your assets are deployed. Valid values:
	//
	// - cn-hangzhou: assets in the Chinese mainland and China (Hong Kong)
	//
	// - ap-southeast-1: assets outside China
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. This parameter is used by an administrator to switch to the perspective of the member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view.
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts that belong to the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1723057091000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The verification result for the accuracy of alert fields based on the alert template.
	//
	// - true: The verification is passed. Alerts that are generated for enabled rules can be synchronized to the product.
	//
	// - false: The verification failed. Alerts cannot be synchronized to the product.
	//
	// example:
	//
	// true
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s ListCustomizeRuleTestResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizeRuleTestResultRequest) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCustomizeRuleTestResultRequest) GetDetectionRuleId() *string {
	return s.DetectionRuleId
}

func (s *ListCustomizeRuleTestResultRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListCustomizeRuleTestResultRequest) GetId() *int64 {
	return s.Id
}

func (s *ListCustomizeRuleTestResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomizeRuleTestResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCustomizeRuleTestResultRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListCustomizeRuleTestResultRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListCustomizeRuleTestResultRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListCustomizeRuleTestResultRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *ListCustomizeRuleTestResultRequest) SetCurrentPage(v int32) *ListCustomizeRuleTestResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetDetectionRuleId(v string) *ListCustomizeRuleTestResultRequest {
	s.DetectionRuleId = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetEndTime(v int64) *ListCustomizeRuleTestResultRequest {
	s.EndTime = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetId(v int64) *ListCustomizeRuleTestResultRequest {
	s.Id = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetPageSize(v int32) *ListCustomizeRuleTestResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetRegionId(v string) *ListCustomizeRuleTestResultRequest {
	s.RegionId = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetRoleFor(v int64) *ListCustomizeRuleTestResultRequest {
	s.RoleFor = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetRoleType(v int32) *ListCustomizeRuleTestResultRequest {
	s.RoleType = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetStartTime(v int64) *ListCustomizeRuleTestResultRequest {
	s.StartTime = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetVerifyType(v string) *ListCustomizeRuleTestResultRequest {
	s.VerifyType = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) Validate() error {
	return dara.Validate(s)
}
