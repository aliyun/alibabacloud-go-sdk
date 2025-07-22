// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterAllowedPrefixHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody
	GetAccessDeniedDetail() *string
	SetAllowedPrefixHistoryList(v []*DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody
	GetAllowedPrefixHistoryList() []*DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList
	SetCode(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody
	GetSuccess() *bool
}

type DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The historical route prefixes.
	AllowedPrefixHistoryList []*DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList `json:"AllowedPrefixHistoryList,omitempty" xml:"AllowedPrefixHistoryList,omitempty" type:"Repeated"`
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

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) GetAllowedPrefixHistoryList() []*DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList {
	return s.AllowedPrefixHistoryList
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetAllowedPrefixHistoryList(v []*DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.AllowedPrefixHistoryList = v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetCode(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetMessage(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList struct {
	// The route prefix.
	AllowedPrefix []*string `json:"AllowedPrefix,omitempty" xml:"AllowedPrefix,omitempty" type:"Repeated"`
	// The time when the historical route prefix entry was created.
	//
	// example:
	//
	// 1673751163000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) GetAllowedPrefix() []*string {
	return s.AllowedPrefix
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) SetAllowedPrefix(v []*string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList {
	s.AllowedPrefix = v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) SetGmtCreate(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) Validate() error {
	return dara.Validate(s)
}
