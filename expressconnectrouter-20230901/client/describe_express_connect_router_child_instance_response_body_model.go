// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterChildInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterChildInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetChildInstanceList(v []*DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) *DescribeExpressConnectRouterChildInstanceResponseBody
	GetChildInstanceList() []*DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList
	SetCode(v string) *DescribeExpressConnectRouterChildInstanceResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DescribeExpressConnectRouterChildInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeExpressConnectRouterChildInstanceResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DescribeExpressConnectRouterChildInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *DescribeExpressConnectRouterChildInstanceResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *DescribeExpressConnectRouterChildInstanceResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeExpressConnectRouterChildInstanceResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeExpressConnectRouterChildInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeExpressConnectRouterChildInstanceResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeExpressConnectRouterChildInstanceResponseBody
	GetTotalCount() *int32
}

type DescribeExpressConnectRouterChildInstanceResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The list of child instances.
	ChildInstanceList []*DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList `json:"ChildInstanceList,omitempty" xml:"ChildInstanceList,omitempty" type:"Repeated"`
	// The status code that indicates whether the call was successful. A value of 200 indicates that the attachment was successful. Other values indicate that the attachment failed. For more information about the causes of failures, see the error codes.
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
	// The maximum number of entries returned. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The message returned by the operation.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If NextToken is empty, no next query is required.
	//
	// - If NextToken is returned, the value indicates the token for the next query.
	//
	// example:
	//
	// AAAAAYws9fJ0Ur4MGm/5OkDoW/Zn0J0/sCjivzwX9oBcwFnWaaas/kSG+J/WzLOxJHS4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 05130E79-588D-5C40-A718-C4863A59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **True**: successful.
	//
	// - **False**: failed.
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

func (s DescribeExpressConnectRouterChildInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterChildInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) GetChildInstanceList() []*DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	return s.ChildInstanceList
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetChildInstanceList(v []*DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.ChildInstanceList = v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetCode(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetMaxResults(v int32) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetMessage(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetNextToken(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetTotalCount(v int32) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) Validate() error {
	if s.ChildInstanceList != nil {
		for _, item := range s.ChildInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList struct {
	// The ID used to associate the Express Connect Router (ECR) with the VBR.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The instance ID of the border router for routing.
	//
	// example:
	//
	// vbr-gw8vjq2zjux3ifsc9****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The name of the child instance.
	//
	// example:
	//
	// test
	ChildInstanceName *string `json:"ChildInstanceName,omitempty" xml:"ChildInstanceName,omitempty"`
	// The Alibaba Cloud account ID to which the child instance belongs.
	//
	// example:
	//
	// 112101171212****
	ChildInstanceOwnerId *int64 `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	// The region where the child instance resides.
	//
	// example:
	//
	// cn-hangzhou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of the child instance. Valid values: **VBR**: virtual border router.
	//
	// example:
	//
	// VBR
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	// The description of the child instance.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The time when the association was created. The time is displayed in UTC and follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2024-01-09T12:18:23Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the association was modified. The time is displayed in UTC and follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2024-01-09T12:18:23Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The Alibaba Cloud account ID to which the border router for routing belongs.
	//
	// example:
	//
	// 167509154715****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the VBR resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The deployment status of the associated resource. Valid values:
	//
	// - **CREATING**: creating.
	//
	// - **ACTIVE**: created.
	//
	// - **ASSOCIATING**: associating.
	//
	// - **DISSOCIATING**: dissociating.
	//
	// - **UPDATING**: updating.
	//
	// - **DELETING**: deleting.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GetAssociationId() *string {
	return s.AssociationId
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GetChildInstanceName() *string {
	return s.ChildInstanceName
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GetChildInstanceOwnerId() *int64 {
	return s.ChildInstanceOwnerId
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GetDescription() *string {
	return s.Description
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GetStatus() *string {
	return s.Status
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetAssociationId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetChildInstanceId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetChildInstanceName(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.ChildInstanceName = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetChildInstanceOwnerId(v int64) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetChildInstanceRegionId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetChildInstanceType(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetDescription(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.Description = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetEcrId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetGmtCreate(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetGmtModified(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.GmtModified = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetOwnerId(v int64) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.OwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetRegionId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.RegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetStatus(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.Status = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) Validate() error {
	return dara.Validate(s)
}
