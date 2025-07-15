// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrafficControlsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeTrafficControlsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTrafficControlsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTrafficControlsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeTrafficControlsResponseBody
	GetTotalCount() *int32
	SetTrafficControls(v *DescribeTrafficControlsResponseBodyTrafficControls) *DescribeTrafficControlsResponseBody
	GetTrafficControls() *DescribeTrafficControlsResponseBodyTrafficControls
}

type DescribeTrafficControlsResponseBody struct {
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
	// 93D91A99-F093-4596-87BA-3C4FBFD3FD8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The returned throttling policy information. It is an array consisting of TrafficControl data.
	TrafficControls *DescribeTrafficControlsResponseBodyTrafficControls `json:"TrafficControls,omitempty" xml:"TrafficControls,omitempty" type:"Struct"`
}

func (s DescribeTrafficControlsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficControlsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTrafficControlsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTrafficControlsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTrafficControlsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTrafficControlsResponseBody) GetTrafficControls() *DescribeTrafficControlsResponseBodyTrafficControls {
	return s.TrafficControls
}

func (s *DescribeTrafficControlsResponseBody) SetPageNumber(v int32) *DescribeTrafficControlsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTrafficControlsResponseBody) SetPageSize(v int32) *DescribeTrafficControlsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTrafficControlsResponseBody) SetRequestId(v string) *DescribeTrafficControlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrafficControlsResponseBody) SetTotalCount(v int32) *DescribeTrafficControlsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTrafficControlsResponseBody) SetTrafficControls(v *DescribeTrafficControlsResponseBodyTrafficControls) *DescribeTrafficControlsResponseBody {
	s.TrafficControls = v
	return s
}

func (s *DescribeTrafficControlsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTrafficControlsResponseBodyTrafficControls struct {
	TrafficControl []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl `json:"TrafficControl,omitempty" xml:"TrafficControl,omitempty" type:"Repeated"`
}

func (s DescribeTrafficControlsResponseBodyTrafficControls) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficControlsResponseBodyTrafficControls) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponseBodyTrafficControls) GetTrafficControl() []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	return s.TrafficControl
}

func (s *DescribeTrafficControlsResponseBodyTrafficControls) SetTrafficControl(v []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) *DescribeTrafficControlsResponseBodyTrafficControls {
	s.TrafficControl = v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControls) Validate() error {
	return dara.Validate(s)
}

type DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl struct {
	// The default throttling value for each API.
	//
	// example:
	//
	// 20000
	ApiDefault *int32 `json:"ApiDefault,omitempty" xml:"ApiDefault,omitempty"`
	// The default throttling value for each app.
	//
	// example:
	//
	// 8000
	AppDefault *int32 `json:"AppDefault,omitempty" xml:"AppDefault,omitempty"`
	// The creation time (UTC) of the throttling policy.
	//
	// example:
	//
	// 2016-01-27T10:19:39Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description of the throttling policy.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The last modification time (UTC) of the throttling policy.
	//
	// example:
	//
	// 2016-01-27T10:34:38Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The returned information about a special throttling policy. It is an array consisting of SpecialPolicy data.
	SpecialPolicies *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies `json:"SpecialPolicies,omitempty" xml:"SpecialPolicies,omitempty" type:"Struct"`
	// The ID of the throttling policy.
	//
	// example:
	//
	// cfed6c970d45481dbe136d6b5ac68c41
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
	// The name of the throttling policy.
	//
	// example:
	//
	// wulingtestq1
	TrafficControlName *string `json:"TrafficControlName,omitempty" xml:"TrafficControlName,omitempty"`
	// The unit to be used in the throttling policy. Valid values:
	//
	// 	- MINUTE
	//
	// 	- HOUR
	//
	// 	- DAY
	//
	// example:
	//
	// Minute
	TrafficControlUnit *string `json:"TrafficControlUnit,omitempty" xml:"TrafficControlUnit,omitempty"`
	// The default throttling value for each user.
	//
	// example:
	//
	// 15000
	UserDefault *int32 `json:"UserDefault,omitempty" xml:"UserDefault,omitempty"`
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) GetApiDefault() *int32 {
	return s.ApiDefault
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) GetAppDefault() *int32 {
	return s.AppDefault
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) GetDescription() *string {
	return s.Description
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) GetSpecialPolicies() *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies {
	return s.SpecialPolicies
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) GetTrafficControlId() *string {
	return s.TrafficControlId
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) GetTrafficControlName() *string {
	return s.TrafficControlName
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) GetTrafficControlUnit() *string {
	return s.TrafficControlUnit
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) GetUserDefault() *int32 {
	return s.UserDefault
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetApiDefault(v int32) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.ApiDefault = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetAppDefault(v int32) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.AppDefault = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetCreatedTime(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.CreatedTime = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetDescription(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.Description = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetModifiedTime(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetSpecialPolicies(v *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.SpecialPolicies = v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetTrafficControlId(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.TrafficControlId = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetTrafficControlName(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.TrafficControlName = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetTrafficControlUnit(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.TrafficControlUnit = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetUserDefault(v int32) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.UserDefault = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) Validate() error {
	return dara.Validate(s)
}

type DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies struct {
	SpecialPolicy []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy `json:"SpecialPolicy,omitempty" xml:"SpecialPolicy,omitempty" type:"Repeated"`
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies) GetSpecialPolicy() []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy {
	return s.SpecialPolicy
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies) SetSpecialPolicy(v []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies {
	s.SpecialPolicy = v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies) Validate() error {
	return dara.Validate(s)
}

type DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy struct {
	// The type of the special throttling policy. Valid values:
	//
	// 	- **APP**
	//
	// 	- **USER**
	//
	// example:
	//
	// USER
	SpecialType *string `json:"SpecialType,omitempty" xml:"SpecialType,omitempty"`
	// The returned information about a special throttling policy. It is an array consisting of Special data.
	Specials *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials `json:"Specials,omitempty" xml:"Specials,omitempty" type:"Struct"`
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy) GetSpecialType() *string {
	return s.SpecialType
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy) GetSpecials() *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials {
	return s.Specials
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy) SetSpecialType(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy {
	s.SpecialType = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy) SetSpecials(v *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy {
	s.Specials = v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy) Validate() error {
	return dara.Validate(s)
}

type DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials struct {
	Special []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial `json:"Special,omitempty" xml:"Special,omitempty" type:"Repeated"`
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials) GetSpecial() []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial {
	return s.Special
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials) SetSpecial(v []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials {
	s.Special = v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials) Validate() error {
	return dara.Validate(s)
}

type DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial struct {
	// The AppId or user account corresponding to SpecialType.
	//
	// example:
	//
	// test_wg@aliyun.com
	SpecialKey *string `json:"SpecialKey,omitempty" xml:"SpecialKey,omitempty"`
	// The throttling value.
	//
	// example:
	//
	// 100
	TrafficValue *int32 `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial) GetSpecialKey() *string {
	return s.SpecialKey
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial) GetTrafficValue() *int32 {
	return s.TrafficValue
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial) SetSpecialKey(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial {
	s.SpecialKey = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial) SetTrafficValue(v int32) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial {
	s.TrafficValue = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial) Validate() error {
	return dara.Validate(s)
}
