// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceGrantedToExpressConnectRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody
	GetDynamicMessage() *string
	SetEcrGrantedInstanceList(v []*DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) *DescribeInstanceGrantedToExpressConnectRouterResponseBody
	GetEcrGrantedInstanceList() []*DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList
	SetHttpStatusCode(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceGrantedToExpressConnectRouterResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponseBody
	GetTotalCount() *int32
}

type DescribeInstanceGrantedToExpressConnectRouterResponseBody struct {
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
	// The network instances whose permissions are granted to the ECR.
	EcrGrantedInstanceList []*DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList `json:"EcrGrantedInstanceList,omitempty" xml:"EcrGrantedInstanceList,omitempty" type:"Repeated"`
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
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFlMqGuJ10uN3l+FX/cBrGDNXUOUifNeOuAJlT4dc3vsWD6DsNSFAC2FtpeH5QOSG2WFdyRoun7gSLCm5o69YnAQ==
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
	// The total number of network instances whose permissions are granted to the ECR.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) GetEcrGrantedInstanceList() []*DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	return s.EcrGrantedInstanceList
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetCode(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetDynamicCode(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetDynamicMessage(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetEcrGrantedInstanceList(v []*DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.EcrGrantedInstanceList = v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetMaxResults(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetMessage(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetNextToken(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetRequestId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetSuccess(v bool) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetTotalCount(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList struct {
	// The ECR ID.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the ECR to which you want to grant permissions.
	//
	// example:
	//
	// 1456408190545060
	EcrOwnerAliUid *string `json:"EcrOwnerAliUid,omitempty" xml:"EcrOwnerAliUid,omitempty"`
	// The time when the network instance was created.
	//
	// example:
	//
	// 1669023139000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the network instance was modified.
	//
	// example:
	//
	// 1669023139000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The authorization ID.
	//
	// example:
	//
	// gr-8gdelo13mi99g1****
	GrantId *string `json:"GrantId,omitempty" xml:"GrantId,omitempty"`
	// The ID of the network instance.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the Alibaba Cloud enterprise account that owns the network instance.
	//
	// example:
	//
	// 26842
	NodeOwnerBid *string `json:"NodeOwnerBid,omitempty" xml:"NodeOwnerBid,omitempty"`
	// The ID of the Alibaba Cloud account that owns the network instance.
	//
	// example:
	//
	// 129845258050****
	NodeOwnerUid *int64 `json:"NodeOwnerUid,omitempty" xml:"NodeOwnerUid,omitempty"`
	// The region ID of the network instance.
	//
	// example:
	//
	// cn-hangzhou
	NodeRegionId *string `json:"NodeRegionId,omitempty" xml:"NodeRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VBR**
	//
	// 	- **VPC**
	//
	// example:
	//
	// VBR
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The state of the network instance.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) GetEcrOwnerAliUid() *string {
	return s.EcrOwnerAliUid
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) GetGrantId() *string {
	return s.GrantId
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) GetNodeOwnerBid() *string {
	return s.NodeOwnerBid
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) GetNodeOwnerUid() *int64 {
	return s.NodeOwnerUid
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) GetNodeRegionId() *string {
	return s.NodeRegionId
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetEcrId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.EcrId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetEcrOwnerAliUid(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.EcrOwnerAliUid = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetGmtCreate(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetGmtModified(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.GmtModified = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetGrantId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.GrantId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeOwnerBid(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeOwnerBid = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeOwnerUid(v int64) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeOwnerUid = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeRegionId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeRegionId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeType(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeType = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetStatus(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.Status = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) Validate() error {
	return dara.Validate(s)
}
