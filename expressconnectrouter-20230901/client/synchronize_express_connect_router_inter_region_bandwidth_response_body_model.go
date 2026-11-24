// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSynchronizeExpressConnectRouterInterRegionBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody
	GetCode() *string
	SetDynamicCode(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody
	GetMessage() *string
	SetRequestId(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody
	GetSuccess() *bool
}

type SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. A value of 200 indicates that the call was successful. For other return values, the call failed. For more information, see the error codes.
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
	// The dynamic error message that is used to replace the `%s` placeholder in the **ErrMessage*	- parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsInstanceId**, the specified **DtsInstanceId*	- parameter is invalid.
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
	// The response message.
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

func (s SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) GetCode() *string {
	return s.Code
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetAccessDeniedDetail(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetCode(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.Code = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetDynamicCode(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetDynamicMessage(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetHttpStatusCode(v int32) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetMessage(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.Message = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetRequestId(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetSuccess(v bool) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.Success = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
