// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DescribeExpressConnectRouterRouteEntriesResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *DescribeExpressConnectRouterRouteEntriesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody
	GetRequestId() *string
	SetRouteEntriesList(v []*DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) *DescribeExpressConnectRouterRouteEntriesResponseBody
	GetRouteEntriesList() []*DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList
	SetSuccess(v bool) *DescribeExpressConnectRouterRouteEntriesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeExpressConnectRouterRouteEntriesResponseBody
	GetTotalCount() *int32
}

type DescribeExpressConnectRouterRouteEntriesResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code that indicates whether the call is successful. A value of 200 indicates success. Other values indicate failure. For more information about the cause of failure, see Error codes.
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
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the **DtsJobId*	- request parameter is invalid.
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
	// The maximum number of entries returned. Valid values: 1 to 2147483647. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The message returned for the instance operation.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If **NextToken*	- is empty, no subsequent query is required.
	//
	// - If **NextToken*	- has a return value, the value indicates the token for the next query.
	//
	// example:
	//
	// gAAAAABkWwFTUMNCdWC0VMYOIylA56Hx6JUfCZlk5hQ5g_fnKmetN6303tqq5UJ2ouJzyT2fDOdzb-NqyEB5jcY8Z2euX7qHDA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of route entries.
	RouteEntriesList []*DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList `json:"RouteEntriesList,omitempty" xml:"RouteEntriesList,omitempty" type:"Repeated"`
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
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExpressConnectRouterRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) GetRouteEntriesList() []*DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	return s.RouteEntriesList
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetCode(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetMaxResults(v int32) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetMessage(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetNextToken(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetRouteEntriesList(v []*DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.RouteEntriesList = v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) Validate() error {
	if s.RouteEntriesList != nil {
		for _, item := range s.RouteEntriesList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList struct {
	// The AS path traversed by the route entry.
	//
	// example:
	//
	// [64993,64512]
	AsPath *string `json:"AsPath,omitempty" xml:"AsPath,omitempty"`
	// The community value carried by the BGP route entry.
	//
	// example:
	//
	// 9001:9263
	Community *string `json:"Community,omitempty" xml:"Community,omitempty"`
	// The destination CIDR block of the route entry.
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The MED value carried by the BGP route entry, which takes effect for routing between the Express Connect Router (ECR) and the TR.
	//
	// - The MED value can be set to 2000. After the configuration, the TR and the Express Connect Router (ECR) serve as the default path.
	//
	// - For a non-default path, the MED value is empty.
	//
	// - For TRs in the same Cloud Enterprise Network (CEN), only one associated object can have the MED value set to 2000.
	//
	// example:
	//
	// 2000
	Med *int64 `json:"Med,omitempty" xml:"Med,omitempty"`
	// The next hop instance ID.
	//
	// example:
	//
	// vbr-hp3u4u5f03tfuljis****
	NexthopInstanceId *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
	// The region where the next hop instance resides.
	//
	// example:
	//
	// cn-hangzhou
	NexthopInstanceRegionId *string `json:"NexthopInstanceRegionId,omitempty" xml:"NexthopInstanceRegionId,omitempty"`
	// The status of the Express Connect Router (ECR) instance.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) GetAsPath() *string {
	return s.AsPath
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) GetCommunity() *string {
	return s.Community
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) GetMed() *int64 {
	return s.Med
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) GetNexthopInstanceId() *string {
	return s.NexthopInstanceId
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) GetNexthopInstanceRegionId() *string {
	return s.NexthopInstanceRegionId
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) GetStatus() *string {
	return s.Status
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetAsPath(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.AsPath = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetCommunity(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.Community = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetDestinationCidrBlock(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetMed(v int64) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.Med = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetNexthopInstanceId(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.NexthopInstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetNexthopInstanceRegionId(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.NexthopInstanceRegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetStatus(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.Status = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) Validate() error {
	return dara.Validate(s)
}
