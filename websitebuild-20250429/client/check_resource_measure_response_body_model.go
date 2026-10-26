// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckResourceMeasureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CheckResourceMeasureResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CheckResourceMeasureResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CheckResourceMeasureResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CheckResourceMeasureResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CheckResourceMeasureResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CheckResourceMeasureResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CheckResourceMeasureResponseBodyModule) *CheckResourceMeasureResponseBody
	GetModule() *CheckResourceMeasureResponseBodyModule
	SetRequestId(v string) *CheckResourceMeasureResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CheckResourceMeasureResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CheckResourceMeasureResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CheckResourceMeasureResponseBody
	GetSynchro() *bool
}

type CheckResourceMeasureResponseBody struct {
	// The detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Specifies whether retries are allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name.
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace `%s` in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters returned.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The application module.
	Module *CheckResourceMeasureResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The root error code.
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// The exception message.
	//
	// example:
	//
	// System exception
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// The backup parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s CheckResourceMeasureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckResourceMeasureResponseBody) GoString() string {
	return s.String()
}

func (s *CheckResourceMeasureResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CheckResourceMeasureResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CheckResourceMeasureResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CheckResourceMeasureResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CheckResourceMeasureResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CheckResourceMeasureResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CheckResourceMeasureResponseBody) GetModule() *CheckResourceMeasureResponseBodyModule {
	return s.Module
}

func (s *CheckResourceMeasureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckResourceMeasureResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CheckResourceMeasureResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CheckResourceMeasureResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CheckResourceMeasureResponseBody) SetAccessDeniedDetail(v string) *CheckResourceMeasureResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetAllowRetry(v bool) *CheckResourceMeasureResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetAppName(v string) *CheckResourceMeasureResponseBody {
	s.AppName = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetDynamicCode(v string) *CheckResourceMeasureResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetDynamicMessage(v string) *CheckResourceMeasureResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetErrorArgs(v []interface{}) *CheckResourceMeasureResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetModule(v *CheckResourceMeasureResponseBodyModule) *CheckResourceMeasureResponseBody {
	s.Module = v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetRequestId(v string) *CheckResourceMeasureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetRootErrorCode(v string) *CheckResourceMeasureResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetRootErrorMsg(v string) *CheckResourceMeasureResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetSynchro(v bool) *CheckResourceMeasureResponseBody {
	s.Synchro = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckResourceMeasureResponseBodyModule struct {
	// The desired total returned only for APP_WORKSPACE. This is the target total for the current check, which is the sum of the existing usage in the saved configuration and ResourceValue. If the compatible DesiredTotal is provided, this total is used directly. The existing usage of issues is the sum of common and word-of-mouth issues in the saved configuration, and the existing usage of competitors is the number of competitors. This parameter is empty for APP and USER checks.
	//
	// example:
	//
	// 21
	DesiredTotal *int64 `json:"DesiredTotal,omitempty" xml:"DesiredTotal,omitempty"`
	// The error code returned when the check fails. This parameter is null when the check passes.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the check fails. This parameter is null when the check passes.
	//
	// example:
	//
	// stream push failed
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The limit returned only for APP_WORKSPACE quantity pre-checks. This is the maximum number of issues or competitors allowed by the current commodity edition, not the remaining quota. This parameter is empty for APP and USER checks or when the edition benefits cannot be determined.
	//
	// example:
	//
	// 20
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// Indicates whether the check is passed.
	//
	// example:
	//
	// False
	Passed *bool `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// The resource identifier.
	//
	// example:
	//
	// TransitRouterVpcAttachment
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// Indicates whether a purchasable higher edition that can accommodate the DesiredTotal exists. This parameter is true only when the capacity of the current edition is insufficient and a higher edition meets the requirements. Otherwise, this parameter is false. This parameter may be empty for APP and USER checks. This field does not indicate that an upgrade has been performed or an upgrade order has been created.
	//
	// example:
	//
	// true
	UpgradeAvailable *bool `json:"UpgradeAvailable,omitempty" xml:"UpgradeAvailable,omitempty"`
}

func (s CheckResourceMeasureResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CheckResourceMeasureResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CheckResourceMeasureResponseBodyModule) GetDesiredTotal() *int64 {
	return s.DesiredTotal
}

func (s *CheckResourceMeasureResponseBodyModule) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CheckResourceMeasureResponseBodyModule) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CheckResourceMeasureResponseBodyModule) GetLimit() *int64 {
	return s.Limit
}

func (s *CheckResourceMeasureResponseBodyModule) GetPassed() *bool {
	return s.Passed
}

func (s *CheckResourceMeasureResponseBodyModule) GetResourceCode() *string {
	return s.ResourceCode
}

func (s *CheckResourceMeasureResponseBodyModule) GetUpgradeAvailable() *bool {
	return s.UpgradeAvailable
}

func (s *CheckResourceMeasureResponseBodyModule) SetDesiredTotal(v int64) *CheckResourceMeasureResponseBodyModule {
	s.DesiredTotal = &v
	return s
}

func (s *CheckResourceMeasureResponseBodyModule) SetErrorCode(v string) *CheckResourceMeasureResponseBodyModule {
	s.ErrorCode = &v
	return s
}

func (s *CheckResourceMeasureResponseBodyModule) SetErrorMessage(v string) *CheckResourceMeasureResponseBodyModule {
	s.ErrorMessage = &v
	return s
}

func (s *CheckResourceMeasureResponseBodyModule) SetLimit(v int64) *CheckResourceMeasureResponseBodyModule {
	s.Limit = &v
	return s
}

func (s *CheckResourceMeasureResponseBodyModule) SetPassed(v bool) *CheckResourceMeasureResponseBodyModule {
	s.Passed = &v
	return s
}

func (s *CheckResourceMeasureResponseBodyModule) SetResourceCode(v string) *CheckResourceMeasureResponseBodyModule {
	s.ResourceCode = &v
	return s
}

func (s *CheckResourceMeasureResponseBodyModule) SetUpgradeAvailable(v bool) *CheckResourceMeasureResponseBodyModule {
	s.UpgradeAvailable = &v
	return s
}

func (s *CheckResourceMeasureResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
