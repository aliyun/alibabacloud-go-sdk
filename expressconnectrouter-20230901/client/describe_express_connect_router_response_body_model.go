// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribeExpressConnectRouterResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DescribeExpressConnectRouterResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeExpressConnectRouterResponseBody
	GetDynamicMessage() *string
	SetEcrList(v []*DescribeExpressConnectRouterResponseBodyEcrList) *DescribeExpressConnectRouterResponseBody
	GetEcrList() []*DescribeExpressConnectRouterResponseBodyEcrList
	SetHttpStatusCode(v int32) *DescribeExpressConnectRouterResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *DescribeExpressConnectRouterResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *DescribeExpressConnectRouterResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeExpressConnectRouterResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeExpressConnectRouterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeExpressConnectRouterResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeExpressConnectRouterResponseBody
	GetTotalCount() *int32
}

type DescribeExpressConnectRouterResponseBody struct {
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
	// IllegalParamFormat.Name
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of Name ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The ECRs.
	EcrList []*DescribeExpressConnectRouterResponseBodyEcrList `json:"EcrList,omitempty" xml:"EcrList,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The total number of entries returned. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 10
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
	// AAAAAdDWBF2w6Olxc+cMPjUtUMpttDGZkffvHCfhBKKNEyCVaq+WUEzuUWpp9+QOApNf6g==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// The total number of ECRs.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExpressConnectRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeExpressConnectRouterResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeExpressConnectRouterResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeExpressConnectRouterResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeExpressConnectRouterResponseBody) GetEcrList() []*DescribeExpressConnectRouterResponseBodyEcrList {
	return s.EcrList
}

func (s *DescribeExpressConnectRouterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeExpressConnectRouterResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeExpressConnectRouterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeExpressConnectRouterResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeExpressConnectRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpressConnectRouterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeExpressConnectRouterResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetCode(v string) *DescribeExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetEcrList(v []*DescribeExpressConnectRouterResponseBodyEcrList) *DescribeExpressConnectRouterResponseBody {
	s.EcrList = v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetMaxResults(v int32) *DescribeExpressConnectRouterResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetMessage(v string) *DescribeExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetNextToken(v string) *DescribeExpressConnectRouterResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetTotalCount(v int32) *DescribeExpressConnectRouterResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExpressConnectRouterResponseBodyEcrList struct {
	// The autonomous system number (ASN) of the ECR.
	//
	// example:
	//
	// 45104
	AlibabaSideAsn *int64 `json:"AlibabaSideAsn,omitempty" xml:"AlibabaSideAsn,omitempty"`
	// The business state of the ECR. Valid values:
	//
	// 	- **Normal:*	- The ECR is running as expected.
	//
	// 	- **FinancialLocked**: The ECR is locked due to overdue payments.
	//
	// example:
	//
	// Normal
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// The description of the ECR.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ECR ID.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The time when the ECR was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-16T01:44:50Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the ECR was modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-16T01:44:50Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name of the ECR.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the Alibaba Cloud account that owns the ECR.
	//
	// example:
	//
	// 170646818729****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group to which the ECR belongs.
	//
	// example:
	//
	// rg-aekzuscospt****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The deployment state of the ECR. Valid values:
	//
	// 	- **ACTIVE**: The ECR is created.
	//
	// 	- **UPDATING**: The ECR is being modified.
	//
	// 	- **ASSOCIATING**: The ECR is being associated with resources.
	//
	// 	- **DISSOCIATING**: The resource is being disassociated from resources.
	//
	// 	- **LOCKED_ATTACHING**: The ECR is locked because it is being associated with an external system.
	//
	// 	- **LOCKED_DETACHING**: The ECR is locked because it is being disassociated from an external system.
	//
	// 	- **RECLAIMING**: The ECR is waiting to release resources.
	//
	// 	- **DELETING**: The ECR is being deleted.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*DescribeExpressConnectRouterResponseBodyEcrListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeExpressConnectRouterResponseBodyEcrList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterResponseBodyEcrList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) GetAlibabaSideAsn() *int64 {
	return s.AlibabaSideAsn
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) GetBizStatus() *string {
	return s.BizStatus
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) GetDescription() *string {
	return s.Description
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) GetName() *string {
	return s.Name
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) GetStatus() *string {
	return s.Status
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) GetTags() []*DescribeExpressConnectRouterResponseBodyEcrListTags {
	return s.Tags
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetAlibabaSideAsn(v int64) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.AlibabaSideAsn = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetBizStatus(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.BizStatus = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetDescription(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.Description = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetEcrId(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetGmtCreate(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetGmtModified(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.GmtModified = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetName(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.Name = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetOwnerId(v int64) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.OwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetResourceGroupId(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetStatus(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.Status = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetTags(v []*DescribeExpressConnectRouterResponseBodyEcrListTags) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.Tags = v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) Validate() error {
	return dara.Validate(s)
}

type DescribeExpressConnectRouterResponseBodyEcrListTags struct {
	// The tag key.
	//
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeExpressConnectRouterResponseBodyEcrListTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterResponseBodyEcrListTags) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) SetTagKey(v string) *DescribeExpressConnectRouterResponseBodyEcrListTags {
	s.TagKey = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) SetTagValue(v string) *DescribeExpressConnectRouterResponseBodyEcrListTags {
	s.TagValue = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) Validate() error {
	return dara.Validate(s)
}
