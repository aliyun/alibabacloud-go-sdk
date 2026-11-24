// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterRegionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribeExpressConnectRouterRegionResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DescribeExpressConnectRouterRegionResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeExpressConnectRouterRegionResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DescribeExpressConnectRouterRegionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeExpressConnectRouterRegionResponseBody
	GetMessage() *string
	SetRegionIdList(v []*string) *DescribeExpressConnectRouterRegionResponseBody
	GetRegionIdList() []*string
	SetRequestId(v string) *DescribeExpressConnectRouterRegionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeExpressConnectRouterRegionResponseBody
	GetSuccess() *bool
}

type DescribeExpressConnectRouterRegionResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether the call is successful. A value of 200 indicates success, and other values indicate failure. For the cause of failure, see error codes.
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
	// The dynamic error message, which is used to replace `%s` in the error message of the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the **DtsJobId*	- request parameter is invalid.
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
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of regions where sub-resources are located.
	RegionIdList []*string `json:"RegionIdList,omitempty" xml:"RegionIdList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **True**: The call is successful.
	//
	// - **False**: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExpressConnectRouterRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRegionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeExpressConnectRouterRegionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeExpressConnectRouterRegionResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeExpressConnectRouterRegionResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeExpressConnectRouterRegionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeExpressConnectRouterRegionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeExpressConnectRouterRegionResponseBody) GetRegionIdList() []*string {
	return s.RegionIdList
}

func (s *DescribeExpressConnectRouterRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpressConnectRouterRegionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetCode(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterRegionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetMessage(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetRegionIdList(v []*string) *DescribeExpressConnectRouterRegionResponseBody {
	s.RegionIdList = v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterRegionResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) Validate() error {
	return dara.Validate(s)
}
