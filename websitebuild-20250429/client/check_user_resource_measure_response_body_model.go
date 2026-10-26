// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserResourceMeasureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CheckUserResourceMeasureResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CheckUserResourceMeasureResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CheckUserResourceMeasureResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CheckUserResourceMeasureResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CheckUserResourceMeasureResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CheckUserResourceMeasureResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CheckUserResourceMeasureResponseBodyModule) *CheckUserResourceMeasureResponseBody
	GetModule() *CheckUserResourceMeasureResponseBodyModule
	SetRequestId(v string) *CheckUserResourceMeasureResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CheckUserResourceMeasureResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CheckUserResourceMeasureResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CheckUserResourceMeasureResponseBody
	GetSynchro() *bool
}

type CheckUserResourceMeasureResponseBody struct {
	// The detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Specifies whether retries are allowed. Valid values:
	//
	// - false: Retries are not allowed.
	//
	// - true: Retries are allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The name of the application.
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
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The returned error arguments.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The returned object.
	Module *CheckUserResourceMeasureResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// The root error message.
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

func (s CheckUserResourceMeasureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckUserResourceMeasureResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserResourceMeasureResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CheckUserResourceMeasureResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CheckUserResourceMeasureResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CheckUserResourceMeasureResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CheckUserResourceMeasureResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CheckUserResourceMeasureResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CheckUserResourceMeasureResponseBody) GetModule() *CheckUserResourceMeasureResponseBodyModule {
	return s.Module
}

func (s *CheckUserResourceMeasureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckUserResourceMeasureResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CheckUserResourceMeasureResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CheckUserResourceMeasureResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CheckUserResourceMeasureResponseBody) SetAccessDeniedDetail(v string) *CheckUserResourceMeasureResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetAllowRetry(v bool) *CheckUserResourceMeasureResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetAppName(v string) *CheckUserResourceMeasureResponseBody {
	s.AppName = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetDynamicCode(v string) *CheckUserResourceMeasureResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetDynamicMessage(v string) *CheckUserResourceMeasureResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetErrorArgs(v []interface{}) *CheckUserResourceMeasureResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetModule(v *CheckUserResourceMeasureResponseBodyModule) *CheckUserResourceMeasureResponseBody {
	s.Module = v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetRequestId(v string) *CheckUserResourceMeasureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetRootErrorCode(v string) *CheckUserResourceMeasureResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetRootErrorMsg(v string) *CheckUserResourceMeasureResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetSynchro(v bool) *CheckUserResourceMeasureResponseBody {
	s.Synchro = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckUserResourceMeasureResponseBodyModule struct {
	// Returned only for APP_WORKSPACE: the target total for this check, which is the used quota in the saved configuration plus ResourceValue. If the compatible DesiredTotal is provided, this total is used directly. The used quota for questions is the sum of common and word-of-mouth questions in the saved configuration. The used quota for competitors is the number of competitors. This field is empty for APP and USER checks.
	//
	// example:
	//
	// 21
	DesiredTotal *int64 `json:"DesiredTotal,omitempty" xml:"DesiredTotal,omitempty"`
	// The error code when the check fails. This field is null when the check passes.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message when the check fails. This field is null when the check passes.
	//
	// example:
	//
	// stream push failed
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Returned only for APP_WORKSPACE quantity dry runs: the maximum number of questions or competitors allowed by the current commodity edition, not the remaining quota. This field is empty for APP and USER checks or when the edition benefits cannot be determined.
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
	// The resource code.
	//
	// example:
	//
	// TransitRouterVpcAttachment
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// Indicates whether a purchasable higher edition exists that can accommodate the DesiredTotal. This field is true only when the current edition capacity is insufficient and a higher edition meets the requirements. Otherwise, it is false. This field may be empty for APP and USER checks. This field does not indicate that an upgrade has been performed or an upgrade order has been created.
	//
	// example:
	//
	// true
	UpgradeAvailable *bool `json:"UpgradeAvailable,omitempty" xml:"UpgradeAvailable,omitempty"`
}

func (s CheckUserResourceMeasureResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CheckUserResourceMeasureResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CheckUserResourceMeasureResponseBodyModule) GetDesiredTotal() *int64 {
	return s.DesiredTotal
}

func (s *CheckUserResourceMeasureResponseBodyModule) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CheckUserResourceMeasureResponseBodyModule) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CheckUserResourceMeasureResponseBodyModule) GetLimit() *int64 {
	return s.Limit
}

func (s *CheckUserResourceMeasureResponseBodyModule) GetPassed() *bool {
	return s.Passed
}

func (s *CheckUserResourceMeasureResponseBodyModule) GetResourceCode() *string {
	return s.ResourceCode
}

func (s *CheckUserResourceMeasureResponseBodyModule) GetUpgradeAvailable() *bool {
	return s.UpgradeAvailable
}

func (s *CheckUserResourceMeasureResponseBodyModule) SetDesiredTotal(v int64) *CheckUserResourceMeasureResponseBodyModule {
	s.DesiredTotal = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBodyModule) SetErrorCode(v string) *CheckUserResourceMeasureResponseBodyModule {
	s.ErrorCode = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBodyModule) SetErrorMessage(v string) *CheckUserResourceMeasureResponseBodyModule {
	s.ErrorMessage = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBodyModule) SetLimit(v int64) *CheckUserResourceMeasureResponseBodyModule {
	s.Limit = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBodyModule) SetPassed(v bool) *CheckUserResourceMeasureResponseBodyModule {
	s.Passed = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBodyModule) SetResourceCode(v string) *CheckUserResourceMeasureResponseBodyModule {
	s.ResourceCode = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBodyModule) SetUpgradeAvailable(v bool) *CheckUserResourceMeasureResponseBodyModule {
	s.UpgradeAvailable = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
