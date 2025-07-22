// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisabledExpressConnectRouterRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody
	GetCode() *string
	SetDisabledRouteEntryList(v []*DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody
	GetDisabledRouteEntryList() []*DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList
	SetDynamicCode(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody
	GetTotalCount() *int32
}

type DescribeDisabledExpressConnectRouterRouteEntriesResponseBody struct {
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
	// The routes that are queried.
	DisabledRouteEntryList []*DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList `json:"DisabledRouteEntryList,omitempty" xml:"DisabledRouteEntryList,omitempty" type:"Repeated"`
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
	// The total number of entries returned. Valid values: 1 to 2147483647. Default value: 10.
	//
	// example:
	//
	// 6
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If NextToken is empty, no next page exists.
	//
	// 	- If a value of NextToken is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// gAAAAABkDTaRFnmxUoMLVOn8YTIgYFeL2ch8j0sJs8VCIU8SS5438m3D9X1VqspCcaUEHRN9I_AfVwMhZHAhcNivifK_OtQxJQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether routes are disabled by the ECR. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of routes.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GetDisabledRouteEntryList() []*DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	return s.DisabledRouteEntryList
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetCode(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetDisabledRouteEntryList(v []*DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.DisabledRouteEntryList = v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetDynamicCode(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetDynamicMessage(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetHttpStatusCode(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetMaxResults(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetMessage(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetNextToken(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetRequestId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetSuccess(v bool) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList struct {
	// The destination CIDR block of the route.
	//
	// example:
	//
	// 192.168.100.110/32
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ECR ID.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The time when the route entry was created.
	//
	// example:
	//
	// 1682317345
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the next-hop instance.
	//
	// example:
	//
	// br-hp3u4u5f03tfuljis****
	NexthopInstanceId *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
	// The region ID of the next-hop instance.
	//
	// example:
	//
	// cn-hangzhou
	NexthopInstanceRegionId *string `json:"NexthopInstanceRegionId,omitempty" xml:"NexthopInstanceRegionId,omitempty"`
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) GetNexthopInstanceId() *string {
	return s.NexthopInstanceId
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) GetNexthopInstanceRegionId() *string {
	return s.NexthopInstanceRegionId
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetDestinationCidrBlock(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetEcrId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.EcrId = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetGmtCreate(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetNexthopInstanceId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.NexthopInstanceId = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetNexthopInstanceRegionId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.NexthopInstanceRegionId = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) Validate() error {
	return dara.Validate(s)
}
