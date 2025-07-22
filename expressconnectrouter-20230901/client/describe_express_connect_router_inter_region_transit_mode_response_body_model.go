// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterInterRegionTransitModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody
	GetHttpStatusCode() *int32
	SetInterRegionTransitModeList(v []*DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody
	GetInterRegionTransitModeList() []*DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList
	SetMessage(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody
	GetSuccess() *bool
}

type DescribeExpressConnectRouterInterRegionTransitModeResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
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
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
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
	// The cross-region forwarding modes.
	InterRegionTransitModeList []*DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList `json:"InterRegionTransitModeList,omitempty" xml:"InterRegionTransitModeList,omitempty" type:"Repeated"`
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

func (s DescribeExpressConnectRouterInterRegionTransitModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) GetInterRegionTransitModeList() []*DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList {
	return s.InterRegionTransitModeList
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetCode(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetInterRegionTransitModeList(v []*DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.InterRegionTransitModeList = v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetMessage(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList struct {
	// The cross-region forwarding mode of the ECR. Valid values:
	//
	// 	- **ECMP**: the load balancing mode.
	//
	// 	- **NearBy**: the nearby forwarding mode.
	//
	// example:
	//
	// ECMP
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The region ID of the ECR.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) GetMode() *string {
	return s.Mode
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) SetMode(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList {
	s.Mode = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) SetRegionId(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList {
	s.RegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) Validate() error {
	return dara.Validate(s)
}
