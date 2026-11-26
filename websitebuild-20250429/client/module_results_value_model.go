// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleResultsValue interface {
	dara.Model
	String() string
	GoString() string
	SetDesiredTotal(v int64) *ModuleResultsValue
	GetDesiredTotal() *int64
	SetErrorCode(v string) *ModuleResultsValue
	GetErrorCode() *string
	SetErrorMessage(v string) *ModuleResultsValue
	GetErrorMessage() *string
	SetLimit(v int64) *ModuleResultsValue
	GetLimit() *int64
	SetPassed(v bool) *ModuleResultsValue
	GetPassed() *bool
	SetResourceCode(v string) *ModuleResultsValue
	GetResourceCode() *string
	SetUpgradeAvailable(v bool) *ModuleResultsValue
	GetUpgradeAvailable() *bool
}

type ModuleResultsValue struct {
	// Returned only for APP_WORKSPACE. This parameter indicates the target total quantity for this check, which is the used quantity in the saved configuration plus the ResourceValue. If a compatible DesiredTotal is provided, this total quantity is used directly. The used quantity for issues is the sum of normal and word-of-mouth issues in the saved configuration, and the used quantity for competitors is the number of competitors. This parameter is empty for APP and USER checks.
	//
	// example:
	//
	// 21
	DesiredTotal *int64 `json:"DesiredTotal,omitempty" xml:"DesiredTotal,omitempty"`
	// The error code when the check fails. This parameter is null when the check passes.
	//
	// example:
	//
	// Resource.Control.No.Usage
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message when the check fails. This parameter is null when the check passes.
	//
	// example:
	//
	// Quota exhausted
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Returned only for APP_WORKSPACE quantity pre-checks. This parameter indicates the maximum number of issues or competitors allowed by the current commodity version, not the remaining quota. This parameter is empty for APP and USER checks or when the version benefits cannot be determined.
	//
	// example:
	//
	// 20
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// Indicates whether the check passed.
	//
	// example:
	//
	// true
	Passed *bool `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// The resource identifier.
	//
	// example:
	//
	// InspirationTokens
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// Indicates whether a purchasable higher version exists that can accommodate the DesiredTotal. This parameter is true only when the capacity of the current version is insufficient and a higher version meets the requirements. Otherwise, this parameter is false. This parameter may be empty for APP and USER checks. This field does not indicate that an upgrade has been performed or an upgrade order has been created.
	//
	// example:
	//
	// true
	UpgradeAvailable *bool `json:"UpgradeAvailable,omitempty" xml:"UpgradeAvailable,omitempty"`
}

func (s ModuleResultsValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleResultsValue) GoString() string {
	return s.String()
}

func (s *ModuleResultsValue) GetDesiredTotal() *int64 {
	return s.DesiredTotal
}

func (s *ModuleResultsValue) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ModuleResultsValue) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ModuleResultsValue) GetLimit() *int64 {
	return s.Limit
}

func (s *ModuleResultsValue) GetPassed() *bool {
	return s.Passed
}

func (s *ModuleResultsValue) GetResourceCode() *string {
	return s.ResourceCode
}

func (s *ModuleResultsValue) GetUpgradeAvailable() *bool {
	return s.UpgradeAvailable
}

func (s *ModuleResultsValue) SetDesiredTotal(v int64) *ModuleResultsValue {
	s.DesiredTotal = &v
	return s
}

func (s *ModuleResultsValue) SetErrorCode(v string) *ModuleResultsValue {
	s.ErrorCode = &v
	return s
}

func (s *ModuleResultsValue) SetErrorMessage(v string) *ModuleResultsValue {
	s.ErrorMessage = &v
	return s
}

func (s *ModuleResultsValue) SetLimit(v int64) *ModuleResultsValue {
	s.Limit = &v
	return s
}

func (s *ModuleResultsValue) SetPassed(v bool) *ModuleResultsValue {
	s.Passed = &v
	return s
}

func (s *ModuleResultsValue) SetResourceCode(v string) *ModuleResultsValue {
	s.ResourceCode = &v
	return s
}

func (s *ModuleResultsValue) SetUpgradeAvailable(v bool) *ModuleResultsValue {
	s.UpgradeAvailable = &v
	return s
}

func (s *ModuleResultsValue) Validate() error {
	return dara.Validate(s)
}
