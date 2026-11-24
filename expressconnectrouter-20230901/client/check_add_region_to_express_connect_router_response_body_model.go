// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAddRegionToExpressConnectRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CheckAddRegionToExpressConnectRouterResponseBody
	GetAccessDeniedDetail() *string
	SetAnyCrossBorderLink(v bool) *CheckAddRegionToExpressConnectRouterResponseBody
	GetAnyCrossBorderLink() *bool
	SetAnyInterRegionLink(v bool) *CheckAddRegionToExpressConnectRouterResponseBody
	GetAnyInterRegionLink() *bool
	SetCode(v string) *CheckAddRegionToExpressConnectRouterResponseBody
	GetCode() *string
	SetDynamicCode(v string) *CheckAddRegionToExpressConnectRouterResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CheckAddRegionToExpressConnectRouterResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *CheckAddRegionToExpressConnectRouterResponseBody
	GetHttpStatusCode() *int32
	SetIsCdtCrossBorderEnabled(v bool) *CheckAddRegionToExpressConnectRouterResponseBody
	GetIsCdtCrossBorderEnabled() *bool
	SetIsCdtInterRegionEnabled(v bool) *CheckAddRegionToExpressConnectRouterResponseBody
	GetIsCdtInterRegionEnabled() *bool
	SetIsUserAllowedToCreateCrossBorderLink(v bool) *CheckAddRegionToExpressConnectRouterResponseBody
	GetIsUserAllowedToCreateCrossBorderLink() *bool
	SetMessage(v string) *CheckAddRegionToExpressConnectRouterResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckAddRegionToExpressConnectRouterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckAddRegionToExpressConnectRouterResponseBody
	GetSuccess() *bool
}

type CheckAddRegionToExpressConnectRouterResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether the Express Connect Router (ECR) instance has cross-border connections between the Chinese mainland and regions outside the Chinese mainland. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	AnyCrossBorderLink *bool `json:"AnyCrossBorderLink,omitempty" xml:"AnyCrossBorderLink,omitempty"`
	// Indicates whether the Express Connect Router (ECR) instance has cross-region connections in the Chinese mainland. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	AnyInterRegionLink *bool `json:"AnyInterRegionLink,omitempty" xml:"AnyInterRegionLink,omitempty"`
	// Indicates whether the call was successful. A return value of 200 indicates that the request was successful. Any other value indicates that the request failed. For more information about the causes of failure, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace `%s` in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsInstanceId**, it indicates that the **DtsInstanceId*	- request parameter is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Indicates whether the cross-border Cloud Data Transfer (CDT) service is enabled for the account to which the Express Connect Router (ECR) instance belongs. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	IsCdtCrossBorderEnabled *bool `json:"IsCdtCrossBorderEnabled,omitempty" xml:"IsCdtCrossBorderEnabled,omitempty"`
	// Indicates whether the Cloud Data Transfer (CDT) service is enabled for the account to which the Express Connect Router (ECR) instance belongs. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	IsCdtInterRegionEnabled *bool `json:"IsCdtInterRegionEnabled,omitempty" xml:"IsCdtInterRegionEnabled,omitempty"`
	// Indicates whether the account to which the Express Connect Router (ECR) instance belongs is allowed to create cross-border connections. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	IsUserAllowedToCreateCrossBorderLink *bool `json:"IsUserAllowedToCreateCrossBorderLink,omitempty" xml:"IsUserAllowedToCreateCrossBorderLink,omitempty"`
	// The message returned for the operation.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **True**: The call was successful.
	//
	// - **False**: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckAddRegionToExpressConnectRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAddRegionToExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) GetAnyCrossBorderLink() *bool {
	return s.AnyCrossBorderLink
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) GetAnyInterRegionLink() *bool {
	return s.AnyInterRegionLink
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) GetIsCdtCrossBorderEnabled() *bool {
	return s.IsCdtCrossBorderEnabled
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) GetIsCdtInterRegionEnabled() *bool {
	return s.IsCdtInterRegionEnabled
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) GetIsUserAllowedToCreateCrossBorderLink() *bool {
	return s.IsUserAllowedToCreateCrossBorderLink
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetAnyCrossBorderLink(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.AnyCrossBorderLink = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetAnyInterRegionLink(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.AnyInterRegionLink = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetCode(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetDynamicCode(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetDynamicMessage(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetIsCdtCrossBorderEnabled(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.IsCdtCrossBorderEnabled = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetIsCdtInterRegionEnabled(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.IsCdtInterRegionEnabled = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetIsUserAllowedToCreateCrossBorderLink(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.IsUserAllowedToCreateCrossBorderLink = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetMessage(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetRequestId(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetSuccess(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
