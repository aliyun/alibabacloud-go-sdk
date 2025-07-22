// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterAssociationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterAssociationResponseBody
	GetAccessDeniedDetail() *string
	SetAssociationList(v []*DescribeExpressConnectRouterAssociationResponseBodyAssociationList) *DescribeExpressConnectRouterAssociationResponseBody
	GetAssociationList() []*DescribeExpressConnectRouterAssociationResponseBodyAssociationList
	SetCode(v string) *DescribeExpressConnectRouterAssociationResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DescribeExpressConnectRouterAssociationResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeExpressConnectRouterAssociationResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DescribeExpressConnectRouterAssociationResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *DescribeExpressConnectRouterAssociationResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *DescribeExpressConnectRouterAssociationResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeExpressConnectRouterAssociationResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeExpressConnectRouterAssociationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeExpressConnectRouterAssociationResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeExpressConnectRouterAssociationResponseBody
	GetTotalCount() *int32
}

type DescribeExpressConnectRouterAssociationResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The associated resources.
	AssociationList []*DescribeExpressConnectRouterAssociationResponseBodyAssociationList `json:"AssociationList,omitempty" xml:"AssociationList,omitempty" type:"Repeated"`
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
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of DynamicMessage is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
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
	// The total number of entries returned. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 20
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
	// AAAAAYws9fJ0Ur4MGm/5OkDoW/Zn0J0/sCjivzwX9oBcwFnWaaas/kSG+J/WzLOxJHS4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 05130E79-588D-5C40-A718-C4863A59****
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
	// The total number of associated resources.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExpressConnectRouterAssociationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) GetAssociationList() []*DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	return s.AssociationList
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetAssociationList(v []*DescribeExpressConnectRouterAssociationResponseBodyAssociationList) *DescribeExpressConnectRouterAssociationResponseBody {
	s.AssociationList = v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetCode(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterAssociationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetMaxResults(v int32) *DescribeExpressConnectRouterAssociationResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetMessage(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetNextToken(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterAssociationResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetTotalCount(v int32) *DescribeExpressConnectRouterAssociationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExpressConnectRouterAssociationResponseBodyAssociationList struct {
	// The allowed route prefixes.
	AllowedPrefixes []*string `json:"AllowedPrefixes,omitempty" xml:"AllowedPrefixes,omitempty" type:"Repeated"`
	// The prefix route mode. Valid values:
	//
	// 	- MatchMode
	//
	// 	- IncrementalMode
	//
	// example:
	//
	// MatchMode
	AllowedPrefixesMode *string `json:"AllowedPrefixesMode,omitempty" xml:"AllowedPrefixesMode,omitempty"`
	// The ID of the association between the ECR and the VPC or TR.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The type of the associated resource. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **TR**
	//
	// example:
	//
	// VPC
	AssociationNodeType *string `json:"AssociationNodeType,omitempty" xml:"AssociationNodeType,omitempty"`
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-5510frtx8zi54q****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The description of the associated resource.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ECR ID.
	//
	// example:
	//
	// ecr-fu8rszhgv7623c****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The time when the association was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-01-09T12:18:23Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the association was modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-01-09T12:18:23Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the Alibaba Cloud account that owns the resource.
	//
	// example:
	//
	// 167509154715****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The deployment state of the associated resource. Valid values:
	//
	// 	- **CREATING**: The resource is being created.
	//
	// 	- **ACTIVE**: The resource is created.
	//
	// 	- **INACTIVE**: The TR is pending to be associated with the ECR.
	//
	// 	- **ASSOCIATING**: The resource is being associated.
	//
	// 	- **DISSOCIATING**: The resource is being disassociated.
	//
	// 	- **UPDATING**: The resource is being modified.
	//
	// 	- **DELETING**: The resource is being deleted.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The TR ID.
	//
	// example:
	//
	// tr-2ze4i71c6be454e2l****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the TR.
	//
	// example:
	//
	// 189159362009****
	TransitRouterOwnerId *int64 `json:"TransitRouterOwnerId,omitempty" xml:"TransitRouterOwnerId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-2zeeaxet4i2j1a7n7****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 146757288406****
	VpcOwnerId *int64 `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
}

func (s DescribeExpressConnectRouterAssociationResponseBodyAssociationList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetAllowedPrefixes() []*string {
	return s.AllowedPrefixes
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetAllowedPrefixesMode() *string {
	return s.AllowedPrefixesMode
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetAssociationId() *string {
	return s.AssociationId
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetAssociationNodeType() *string {
	return s.AssociationNodeType
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetCenId() *string {
	return s.CenId
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetDescription() *string {
	return s.Description
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetStatus() *string {
	return s.Status
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetTransitRouterOwnerId() *int64 {
	return s.TransitRouterOwnerId
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GetVpcOwnerId() *int64 {
	return s.VpcOwnerId
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetAllowedPrefixes(v []*string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.AllowedPrefixes = v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetAllowedPrefixesMode(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.AllowedPrefixesMode = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetAssociationId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetAssociationNodeType(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.AssociationNodeType = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetCenId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.CenId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetDescription(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.Description = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetEcrId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetGmtCreate(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetGmtModified(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.GmtModified = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetOwnerId(v int64) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.OwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetRegionId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.RegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetStatus(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.Status = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetTransitRouterId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetTransitRouterOwnerId(v int64) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.TransitRouterOwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetVpcId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.VpcId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetVpcOwnerId(v int64) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.VpcOwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) Validate() error {
	return dara.Validate(s)
}
