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
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether the ECR is used to establish connections between regions in the Chinese mainland and regions outside China. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AnyCrossBorderLink *bool `json:"AnyCrossBorderLink,omitempty" xml:"AnyCrossBorderLink,omitempty"`
	// Indicates whether the ECR is used to establish connections between regions in the Chinese mainland. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AnyInterRegionLink *bool `json:"AnyInterRegionLink,omitempty" xml:"AnyInterRegionLink,omitempty"`
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
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
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsInstanceId**, the request parameter **DtsInstanceId*	- is invalid.
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
	// Indicates whether the cross-border CDT service is activated for the account to which the ECR belongs. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsCdtCrossBorderEnabled *bool `json:"IsCdtCrossBorderEnabled,omitempty" xml:"IsCdtCrossBorderEnabled,omitempty"`
	// Indicates whether the CDT service is activated for the account to which the ECR belongs. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsCdtInterRegionEnabled *bool `json:"IsCdtInterRegionEnabled,omitempty" xml:"IsCdtInterRegionEnabled,omitempty"`
	// Indicates whether the account to which the ECR belongs can create cross-border connections. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsUserAllowedToCreateCrossBorderLink *bool `json:"IsUserAllowedToCreateCrossBorderLink,omitempty" xml:"IsUserAllowedToCreateCrossBorderLink,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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
