// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type EnterpriseDingtalkGroupMember struct {
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// true
	IsAdmin *bool `json:"IsAdmin,omitempty" xml:"IsAdmin,omitempty"`
	// 代表资源组的资源属性字段
	//
	// example:
	//
	// 130xxxxxxxx
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// 张三
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s EnterpriseDingtalkGroupMember) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseDingtalkGroupMember) GoString() string {
	return s.String()
}

func (s *EnterpriseDingtalkGroupMember) SetIsAdmin(v bool) *EnterpriseDingtalkGroupMember {
	s.IsAdmin = &v
	return s
}

func (s *EnterpriseDingtalkGroupMember) SetMobile(v string) *EnterpriseDingtalkGroupMember {
	s.Mobile = &v
	return s
}

func (s *EnterpriseDingtalkGroupMember) SetName(v string) *EnterpriseDingtalkGroupMember {
	s.Name = &v
	return s
}

type AddEnterpriseGroupMemberToTaskGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// E240815B72K5D3
	TaskOrderId *string `json:"TaskOrderId,omitempty" xml:"TaskOrderId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PtWoW82DJI1zcTwsT98kLIgAj7kfASzfC6StcpV7hKs=
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddEnterpriseGroupMemberToTaskGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEnterpriseGroupMemberToTaskGroupRequest) GoString() string {
	return s.String()
}

func (s *AddEnterpriseGroupMemberToTaskGroupRequest) SetTaskOrderId(v string) *AddEnterpriseGroupMemberToTaskGroupRequest {
	s.TaskOrderId = &v
	return s
}

func (s *AddEnterpriseGroupMemberToTaskGroupRequest) SetUserId(v string) *AddEnterpriseGroupMemberToTaskGroupRequest {
	s.UserId = &v
	return s
}

type AddEnterpriseGroupMemberToTaskGroupResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43135C31-E47A-5AD7-A693-6DC635201CE4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddEnterpriseGroupMemberToTaskGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEnterpriseGroupMemberToTaskGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddEnterpriseGroupMemberToTaskGroupResponseBody) SetCode(v string) *AddEnterpriseGroupMemberToTaskGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddEnterpriseGroupMemberToTaskGroupResponseBody) SetData(v string) *AddEnterpriseGroupMemberToTaskGroupResponseBody {
	s.Data = &v
	return s
}

func (s *AddEnterpriseGroupMemberToTaskGroupResponseBody) SetMessage(v string) *AddEnterpriseGroupMemberToTaskGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AddEnterpriseGroupMemberToTaskGroupResponseBody) SetRequestId(v string) *AddEnterpriseGroupMemberToTaskGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEnterpriseGroupMemberToTaskGroupResponseBody) SetSuccess(v bool) *AddEnterpriseGroupMemberToTaskGroupResponseBody {
	s.Success = &v
	return s
}

type AddEnterpriseGroupMemberToTaskGroupResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddEnterpriseGroupMemberToTaskGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddEnterpriseGroupMemberToTaskGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEnterpriseGroupMemberToTaskGroupResponse) GoString() string {
	return s.String()
}

func (s *AddEnterpriseGroupMemberToTaskGroupResponse) SetHeaders(v map[string]*string) *AddEnterpriseGroupMemberToTaskGroupResponse {
	s.Headers = v
	return s
}

func (s *AddEnterpriseGroupMemberToTaskGroupResponse) SetStatusCode(v int32) *AddEnterpriseGroupMemberToTaskGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddEnterpriseGroupMemberToTaskGroupResponse) SetBody(v *AddEnterpriseGroupMemberToTaskGroupResponseBody) *AddEnterpriseGroupMemberToTaskGroupResponse {
	s.Body = v
	return s
}

type CreateTaskOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	IsUrgent     *bool   `json:"IsUrgent,omitempty" xml:"IsUrgent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cid+lUpHxTIXt7DYqJDcpVxlA==
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 任务单标题：必填
	Overview *string `json:"Overview,omitempty" xml:"Overview,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 重要性描述
	UrgentDescription *string `json:"UrgentDescription,omitempty" xml:"UrgentDescription,omitempty"`
}

func (s CreateTaskOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskOrderRequest) SetCreateUserId(v string) *CreateTaskOrderRequest {
	s.CreateUserId = &v
	return s
}

func (s *CreateTaskOrderRequest) SetIsUrgent(v bool) *CreateTaskOrderRequest {
	s.IsUrgent = &v
	return s
}

func (s *CreateTaskOrderRequest) SetOpenGroupId(v string) *CreateTaskOrderRequest {
	s.OpenGroupId = &v
	return s
}

func (s *CreateTaskOrderRequest) SetOverview(v string) *CreateTaskOrderRequest {
	s.Overview = &v
	return s
}

func (s *CreateTaskOrderRequest) SetProductCode(v string) *CreateTaskOrderRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateTaskOrderRequest) SetUrgentDescription(v string) *CreateTaskOrderRequest {
	s.UrgentDescription = &v
	return s
}

type CreateTaskOrderResponseBody struct {
	// code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// E21111796147LE
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// msg
	//
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// 123
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTaskOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskOrderResponseBody) SetCode(v string) *CreateTaskOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskOrderResponseBody) SetData(v string) *CreateTaskOrderResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTaskOrderResponseBody) SetMessage(v string) *CreateTaskOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskOrderResponseBody) SetRequestId(v string) *CreateTaskOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskOrderResponseBody) SetSuccess(v bool) *CreateTaskOrderResponseBody {
	s.Success = &v
	return s
}

type CreateTaskOrderResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTaskOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskOrderResponse) SetHeaders(v map[string]*string) *CreateTaskOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskOrderResponse) SetStatusCode(v int32) *CreateTaskOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskOrderResponse) SetBody(v *CreateTaskOrderResponseBody) *CreateTaskOrderResponse {
	s.Body = v
	return s
}

type ListDdTaskOrderRequest struct {
	// createRealName
	//
	// example:
	//
	// Tom
	CreateRealName *string `json:"CreateRealName,omitempty" xml:"CreateRealName,omitempty"`
	// endTime
	//
	// example:
	//
	// 2024-08-20 14:09:16
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsUrgent *bool   `json:"IsUrgent,omitempty" xml:"IsUrgent,omitempty"`
	// openGroupId
	//
	// This parameter is required.
	//
	// example:
	//
	// DAWNN14N
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
	// pageNo
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// pageSize
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// startTime
	//
	// example:
	//
	// 2024-08-01 14:09:11
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// taskStatus
	//
	// example:
	//
	// dealingNode
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s ListDdTaskOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDdTaskOrderRequest) GoString() string {
	return s.String()
}

func (s *ListDdTaskOrderRequest) SetCreateRealName(v string) *ListDdTaskOrderRequest {
	s.CreateRealName = &v
	return s
}

func (s *ListDdTaskOrderRequest) SetEndTime(v string) *ListDdTaskOrderRequest {
	s.EndTime = &v
	return s
}

func (s *ListDdTaskOrderRequest) SetIsUrgent(v bool) *ListDdTaskOrderRequest {
	s.IsUrgent = &v
	return s
}

func (s *ListDdTaskOrderRequest) SetOpenGroupId(v string) *ListDdTaskOrderRequest {
	s.OpenGroupId = &v
	return s
}

func (s *ListDdTaskOrderRequest) SetPageNo(v int64) *ListDdTaskOrderRequest {
	s.PageNo = &v
	return s
}

func (s *ListDdTaskOrderRequest) SetPageSize(v int64) *ListDdTaskOrderRequest {
	s.PageSize = &v
	return s
}

func (s *ListDdTaskOrderRequest) SetStartTime(v string) *ListDdTaskOrderRequest {
	s.StartTime = &v
	return s
}

func (s *ListDdTaskOrderRequest) SetTaskStatus(v string) *ListDdTaskOrderRequest {
	s.TaskStatus = &v
	return s
}

type ListDdTaskOrderResponseBody struct {
	// code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	//
	// example:
	//
	// {   "msg": "请求成功",   "code": "200",   "data": [     {       "wfNodeId": "dealingNode",//任务单状态       "createEmpId": null,//创建人工号，可以忽略       "createTime": 1637571435000,//任务单创建时间       "orderId": "E2111221H1UKCZ",//任务单号       "isImportant": "normal",//是否紧急       "closeTime": null,//任务单关单时间       "taskTitle": "测试单16点44",//标题       "productType": "3270"//问题分类     },     {       "wfNodeId": "dealingNode",       "createEmpId": null,       "createTime": 1637820497000,       "orderId": "E211125CG111EM",       "isImportant": "normal",       "closeTime": null,       "taskTitle": "测试单14",       "class": "com.aliyun.dingtalklanding.dto.OpenTaskOrderDTO",       "productType": "3270"     }   ],   "success": true,   "requestId": "123",   "class": "com.aliyun.dingtalklanding.pop.dto.PopResultDTO" }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// msg
	//
	// example:
	//
	// 请求成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// 2F8557E4-742B-1CF7-8E83-28CD0C1F7B48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDdTaskOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDdTaskOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ListDdTaskOrderResponseBody) SetCode(v string) *ListDdTaskOrderResponseBody {
	s.Code = &v
	return s
}

func (s *ListDdTaskOrderResponseBody) SetData(v string) *ListDdTaskOrderResponseBody {
	s.Data = &v
	return s
}

func (s *ListDdTaskOrderResponseBody) SetMessage(v string) *ListDdTaskOrderResponseBody {
	s.Message = &v
	return s
}

func (s *ListDdTaskOrderResponseBody) SetRequestId(v string) *ListDdTaskOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDdTaskOrderResponseBody) SetSuccess(v bool) *ListDdTaskOrderResponseBody {
	s.Success = &v
	return s
}

type ListDdTaskOrderResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDdTaskOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDdTaskOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDdTaskOrderResponse) GoString() string {
	return s.String()
}

func (s *ListDdTaskOrderResponse) SetHeaders(v map[string]*string) *ListDdTaskOrderResponse {
	s.Headers = v
	return s
}

func (s *ListDdTaskOrderResponse) SetStatusCode(v int32) *ListDdTaskOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDdTaskOrderResponse) SetBody(v *ListDdTaskOrderResponseBody) *ListDdTaskOrderResponse {
	s.Body = v
	return s
}

type ListEnterpriseDingtalkGroupCustomerMembersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DAWNN14N
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
}

func (s ListEnterpriseDingtalkGroupCustomerMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupCustomerMembersRequest) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersRequest) SetOpenGroupId(v string) *ListEnterpriseDingtalkGroupCustomerMembersRequest {
	s.OpenGroupId = &v
	return s
}

type ListEnterpriseDingtalkGroupCustomerMembersResponseBody struct {
	// example:
	//
	// 200
	Code *string                                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListEnterpriseDingtalkGroupCustomerMembersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Invalid data
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 123
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEnterpriseDingtalkGroupCustomerMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupCustomerMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetCode(v string) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetData(v []*ListEnterpriseDingtalkGroupCustomerMembersResponseBodyData) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.Data = v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetMessage(v string) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetRequestId(v string) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetSuccess(v bool) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.Success = &v
	return s
}

type ListEnterpriseDingtalkGroupCustomerMembersResponseBodyData struct {
	IsAdmin *bool   `json:"IsAdmin,omitempty" xml:"IsAdmin,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UserId  *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListEnterpriseDingtalkGroupCustomerMembersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupCustomerMembersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBodyData) SetIsAdmin(v bool) *ListEnterpriseDingtalkGroupCustomerMembersResponseBodyData {
	s.IsAdmin = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBodyData) SetName(v string) *ListEnterpriseDingtalkGroupCustomerMembersResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBodyData) SetUserId(v string) *ListEnterpriseDingtalkGroupCustomerMembersResponseBodyData {
	s.UserId = &v
	return s
}

type ListEnterpriseDingtalkGroupCustomerMembersResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnterpriseDingtalkGroupCustomerMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnterpriseDingtalkGroupCustomerMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupCustomerMembersResponse) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponse) SetHeaders(v map[string]*string) *ListEnterpriseDingtalkGroupCustomerMembersResponse {
	s.Headers = v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponse) SetStatusCode(v int32) *ListEnterpriseDingtalkGroupCustomerMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponse) SetBody(v *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) *ListEnterpriseDingtalkGroupCustomerMembersResponse {
	s.Body = v
	return s
}

type ListEnterpriseDingtalkGroupsResponseBody struct {
	// example:
	//
	// true
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListEnterpriseDingtalkGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Data Invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 12xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEnterpriseDingtalkGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetCode(v string) *ListEnterpriseDingtalkGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetData(v []*ListEnterpriseDingtalkGroupsResponseBodyData) *ListEnterpriseDingtalkGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetMessage(v string) *ListEnterpriseDingtalkGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetRequestId(v string) *ListEnterpriseDingtalkGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetSuccess(v bool) *ListEnterpriseDingtalkGroupsResponseBody {
	s.Success = &v
	return s
}

type ListEnterpriseDingtalkGroupsResponseBodyData struct {
	// example:
	//
	// A公司服务群
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 123
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
}

func (s ListEnterpriseDingtalkGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupsResponseBodyData) SetGroupName(v string) *ListEnterpriseDingtalkGroupsResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBodyData) SetOpenGroupId(v string) *ListEnterpriseDingtalkGroupsResponseBodyData {
	s.OpenGroupId = &v
	return s
}

type ListEnterpriseDingtalkGroupsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnterpriseDingtalkGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnterpriseDingtalkGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupsResponse) SetHeaders(v map[string]*string) *ListEnterpriseDingtalkGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponse) SetStatusCode(v int32) *ListEnterpriseDingtalkGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponse) SetBody(v *ListEnterpriseDingtalkGroupsResponseBody) *ListEnterpriseDingtalkGroupsResponse {
	s.Body = v
	return s
}

type ListProductByGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidXcezGVQJjiWy2PzXylGwvg==
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
}

func (s ListProductByGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductByGroupRequest) GoString() string {
	return s.String()
}

func (s *ListProductByGroupRequest) SetOpenGroupId(v string) *ListProductByGroupRequest {
	s.OpenGroupId = &v
	return s
}

type ListProductByGroupResponseBody struct {
	// code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	//
	// example:
	//
	// [     {       "wfNodeId": "dealingNode",//任务单状态       "createEmpId": null,//创建人工号，可以忽略       "createTime": 1637571435000,//任务单创建时间       "orderId": "E2111221H1UKCZ",//任务单号       "isImportant": "normal",//是否紧急       "closeTime": null,//任务单关单时间       "taskTitle": "测试单16点44",//标题       "productType": "3270"//问题分类     },     {       "wfNodeId": "dealingNode",       "createEmpId": null,       "createTime": 1637820497000,       "orderId": "E211125CG111EM",       "isImportant": "normal",       "closeTime": null,       "taskTitle": "测试单14",       "class": "com.aliyun.dingtalklanding.dto.OpenTaskOrderDTO",       "productType": "3270"     }   ]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// msg
	//
	// example:
	//
	// 请求成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// 123
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProductByGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductByGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductByGroupResponseBody) SetCode(v string) *ListProductByGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductByGroupResponseBody) SetData(v string) *ListProductByGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ListProductByGroupResponseBody) SetMessage(v string) *ListProductByGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListProductByGroupResponseBody) SetRequestId(v string) *ListProductByGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductByGroupResponseBody) SetSuccess(v bool) *ListProductByGroupResponseBody {
	s.Success = &v
	return s
}

type ListProductByGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductByGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductByGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductByGroupResponse) GoString() string {
	return s.String()
}

func (s *ListProductByGroupResponse) SetHeaders(v map[string]*string) *ListProductByGroupResponse {
	s.Headers = v
	return s
}

func (s *ListProductByGroupResponse) SetStatusCode(v int32) *ListProductByGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductByGroupResponse) SetBody(v *ListProductByGroupResponseBody) *ListProductByGroupResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("support-plan"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加客户服务主群人员进子群
//
// @param request - AddEnterpriseGroupMemberToTaskGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEnterpriseGroupMemberToTaskGroupResponse
func (client *Client) AddEnterpriseGroupMemberToTaskGroupWithOptions(request *AddEnterpriseGroupMemberToTaskGroupRequest, runtime *util.RuntimeOptions) (_result *AddEnterpriseGroupMemberToTaskGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskOrderId)) {
		query["TaskOrderId"] = request.TaskOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddEnterpriseGroupMemberToTaskGroup"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AddEnterpriseGroupMemberToTaskGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AddEnterpriseGroupMemberToTaskGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 添加客户服务主群人员进子群
//
// @param request - AddEnterpriseGroupMemberToTaskGroupRequest
//
// @return AddEnterpriseGroupMemberToTaskGroupResponse
func (client *Client) AddEnterpriseGroupMemberToTaskGroup(request *AddEnterpriseGroupMemberToTaskGroupRequest) (_result *AddEnterpriseGroupMemberToTaskGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddEnterpriseGroupMemberToTaskGroupResponse{}
	_body, _err := client.AddEnterpriseGroupMemberToTaskGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建工单
//
// @param request - CreateTaskOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskOrderResponse
func (client *Client) CreateTaskOrderWithOptions(request *CreateTaskOrderRequest, runtime *util.RuntimeOptions) (_result *CreateTaskOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateUserId)) {
		query["CreateUserId"] = request.CreateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.IsUrgent)) {
		query["IsUrgent"] = request.IsUrgent
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupId)) {
		query["OpenGroupId"] = request.OpenGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Overview)) {
		query["Overview"] = request.Overview
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.UrgentDescription)) {
		query["UrgentDescription"] = request.UrgentDescription
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTaskOrder"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateTaskOrderResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateTaskOrderResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建工单
//
// @param request - CreateTaskOrderRequest
//
// @return CreateTaskOrderResponse
func (client *Client) CreateTaskOrder(request *CreateTaskOrderRequest) (_result *CreateTaskOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTaskOrderResponse{}
	_body, _err := client.CreateTaskOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ListDdTaskOrder
//
// @param request - ListDdTaskOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDdTaskOrderResponse
func (client *Client) ListDdTaskOrderWithOptions(request *ListDdTaskOrderRequest, runtime *util.RuntimeOptions) (_result *ListDdTaskOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateRealName)) {
		query["CreateRealName"] = request.CreateRealName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsUrgent)) {
		query["IsUrgent"] = request.IsUrgent
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupId)) {
		query["OpenGroupId"] = request.OpenGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		query["TaskStatus"] = request.TaskStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDdTaskOrder"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListDdTaskOrderResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListDdTaskOrderResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// ListDdTaskOrder
//
// @param request - ListDdTaskOrderRequest
//
// @return ListDdTaskOrderResponse
func (client *Client) ListDdTaskOrder(request *ListDdTaskOrderRequest) (_result *ListDdTaskOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDdTaskOrderResponse{}
	_body, _err := client.ListDdTaskOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取钉群中所有客户侧人员信息
//
// @param request - ListEnterpriseDingtalkGroupCustomerMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnterpriseDingtalkGroupCustomerMembersResponse
func (client *Client) ListEnterpriseDingtalkGroupCustomerMembersWithOptions(request *ListEnterpriseDingtalkGroupCustomerMembersRequest, runtime *util.RuntimeOptions) (_result *ListEnterpriseDingtalkGroupCustomerMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenGroupId)) {
		query["OpenGroupId"] = request.OpenGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnterpriseDingtalkGroupCustomerMembers"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListEnterpriseDingtalkGroupCustomerMembersResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListEnterpriseDingtalkGroupCustomerMembersResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取钉群中所有客户侧人员信息
//
// @param request - ListEnterpriseDingtalkGroupCustomerMembersRequest
//
// @return ListEnterpriseDingtalkGroupCustomerMembersResponse
func (client *Client) ListEnterpriseDingtalkGroupCustomerMembers(request *ListEnterpriseDingtalkGroupCustomerMembersRequest) (_result *ListEnterpriseDingtalkGroupCustomerMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEnterpriseDingtalkGroupCustomerMembersResponse{}
	_body, _err := client.ListEnterpriseDingtalkGroupCustomerMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询所有企业钉群成员
//
// @param request - ListEnterpriseDingtalkGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnterpriseDingtalkGroupsResponse
func (client *Client) ListEnterpriseDingtalkGroupsWithOptions(runtime *util.RuntimeOptions) (_result *ListEnterpriseDingtalkGroupsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListEnterpriseDingtalkGroups"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListEnterpriseDingtalkGroupsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListEnterpriseDingtalkGroupsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询所有企业钉群成员
//
// @return ListEnterpriseDingtalkGroupsResponse
func (client *Client) ListEnterpriseDingtalkGroups() (_result *ListEnterpriseDingtalkGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEnterpriseDingtalkGroupsResponse{}
	_body, _err := client.ListEnterpriseDingtalkGroupsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取问题分类
//
// @param request - ListProductByGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductByGroupResponse
func (client *Client) ListProductByGroupWithOptions(request *ListProductByGroupRequest, runtime *util.RuntimeOptions) (_result *ListProductByGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenGroupId)) {
		query["OpenGroupId"] = request.OpenGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductByGroup"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListProductByGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListProductByGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取问题分类
//
// @param request - ListProductByGroupRequest
//
// @return ListProductByGroupResponse
func (client *Client) ListProductByGroup(request *ListProductByGroupRequest) (_result *ListProductByGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProductByGroupResponse{}
	_body, _err := client.ListProductByGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
