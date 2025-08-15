// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListInstanceAclResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListInstanceAclResponseBody
	GetCode() *string
	SetData(v *ListInstanceAclResponseBodyData) *ListInstanceAclResponseBody
	GetData() *ListInstanceAclResponseBodyData
	SetDynamicCode(v string) *ListInstanceAclResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListInstanceAclResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListInstanceAclResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInstanceAclResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstanceAclResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstanceAclResponseBody
	GetSuccess() *bool
}

type ListInstanceAclResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied due to the reason that the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListInstanceAclResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DA4D2F89-E2C8-5F04-936B-60D55B055FA7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListInstanceAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAclResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceAclResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListInstanceAclResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceAclResponseBody) GetData() *ListInstanceAclResponseBodyData {
	return s.Data
}

func (s *ListInstanceAclResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListInstanceAclResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListInstanceAclResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstanceAclResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceAclResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceAclResponseBody) SetAccessDeniedDetail(v string) *ListInstanceAclResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListInstanceAclResponseBody) SetCode(v string) *ListInstanceAclResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceAclResponseBody) SetData(v *ListInstanceAclResponseBodyData) *ListInstanceAclResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceAclResponseBody) SetDynamicCode(v string) *ListInstanceAclResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListInstanceAclResponseBody) SetDynamicMessage(v string) *ListInstanceAclResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListInstanceAclResponseBody) SetHttpStatusCode(v int32) *ListInstanceAclResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstanceAclResponseBody) SetMessage(v string) *ListInstanceAclResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceAclResponseBody) SetRequestId(v string) *ListInstanceAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceAclResponseBody) SetSuccess(v bool) *ListInstanceAclResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceAclResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceAclResponseBodyData struct {
	// The pagination information.
	List []*ListInstanceAclResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 24
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInstanceAclResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAclResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceAclResponseBodyData) GetList() []*ListInstanceAclResponseBodyDataList {
	return s.List
}

func (s *ListInstanceAclResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListInstanceAclResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListInstanceAclResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstanceAclResponseBodyData) SetList(v []*ListInstanceAclResponseBodyDataList) *ListInstanceAclResponseBodyData {
	s.List = v
	return s
}

func (s *ListInstanceAclResponseBodyData) SetPageNumber(v int64) *ListInstanceAclResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceAclResponseBodyData) SetPageSize(v int64) *ListInstanceAclResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstanceAclResponseBodyData) SetTotalCount(v int64) *ListInstanceAclResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceAclResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListInstanceAclResponseBodyDataList struct {
	// The ACL type.
	//
	// Valid value:
	//
	// 	- APACHE: open source ACL.
	//
	// example:
	//
	// APACHE
	AclType *string `json:"aclType,omitempty" xml:"aclType,omitempty"`
	// The types of the operations that are allowed by the ACL.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The decision result.
	//
	// Valid values:
	//
	// 	- Deny: Access is denied.
	//
	// 	- Allow: Access is allowed.
	//
	// example:
	//
	// Allow
	Decision *string `json:"decision,omitempty" xml:"decision,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The IP address whitelists.
	IpWhitelists []*string `json:"ipWhitelists,omitempty" xml:"ipWhitelists,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// test
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- Group
	//
	// 	- Topic
	//
	// example:
	//
	// Topic
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The username.
	//
	// example:
	//
	// test
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListInstanceAclResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAclResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListInstanceAclResponseBodyDataList) GetAclType() *string {
	return s.AclType
}

func (s *ListInstanceAclResponseBodyDataList) GetActions() []*string {
	return s.Actions
}

func (s *ListInstanceAclResponseBodyDataList) GetDecision() *string {
	return s.Decision
}

func (s *ListInstanceAclResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceAclResponseBodyDataList) GetIpWhitelists() []*string {
	return s.IpWhitelists
}

func (s *ListInstanceAclResponseBodyDataList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceAclResponseBodyDataList) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListInstanceAclResponseBodyDataList) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListInstanceAclResponseBodyDataList) GetUsername() *string {
	return s.Username
}

func (s *ListInstanceAclResponseBodyDataList) SetAclType(v string) *ListInstanceAclResponseBodyDataList {
	s.AclType = &v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetActions(v []*string) *ListInstanceAclResponseBodyDataList {
	s.Actions = v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetDecision(v string) *ListInstanceAclResponseBodyDataList {
	s.Decision = &v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetInstanceId(v string) *ListInstanceAclResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetIpWhitelists(v []*string) *ListInstanceAclResponseBodyDataList {
	s.IpWhitelists = v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetRegionId(v string) *ListInstanceAclResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetResourceName(v string) *ListInstanceAclResponseBodyDataList {
	s.ResourceName = &v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetResourceType(v string) *ListInstanceAclResponseBodyDataList {
	s.ResourceType = &v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetUsername(v string) *ListInstanceAclResponseBodyDataList {
	s.Username = &v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
