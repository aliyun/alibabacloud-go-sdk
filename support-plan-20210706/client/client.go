// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type EnterpriseDingtalkGroupMember struct {
	// 是否企业钉群管理员
	IsAdmin *bool `json:"IsAdmin,omitempty" xml:"IsAdmin,omitempty"`
	// 成员手机号
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// 成员姓名
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

type ListEnterpriseDingtalkGroupCustomerMembersRequest struct {
	// 企业服务群ID
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
	// 接口请求的唯一ID, 每次调用requestID唯一
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用接口返回是否成功, true代表调用正常
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误信息, 当success=false的时候, 可以取到message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 接口请求结果返回码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 企业服务群成员列表
	Data []*EnterpriseDingtalkGroupMember `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListEnterpriseDingtalkGroupCustomerMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupCustomerMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetRequestId(v string) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetSuccess(v bool) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.Success = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetMessage(v string) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetCode(v string) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetData(v []*EnterpriseDingtalkGroupMember) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.Data = v
	return s
}

type ListEnterpriseDingtalkGroupCustomerMembersResponse struct {
	Headers map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEnterpriseDingtalkGroupCustomerMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponse) SetBody(v *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) *ListEnterpriseDingtalkGroupCustomerMembersResponse {
	s.Body = v
	return s
}

type ListEnterpriseDingtalkGroupsResponseBody struct {
	// 接口请求的唯一ID, 每次调用requestID唯一
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用接口返回是否成功, true代表调用正常
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误信息, 当success=false的时候, 可以取到message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 接口请求结果返回码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 服务钉群数组
	Data []*ListEnterpriseDingtalkGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListEnterpriseDingtalkGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetRequestId(v string) *ListEnterpriseDingtalkGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetSuccess(v bool) *ListEnterpriseDingtalkGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetMessage(v string) *ListEnterpriseDingtalkGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetCode(v string) *ListEnterpriseDingtalkGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetData(v []*ListEnterpriseDingtalkGroupsResponseBodyData) *ListEnterpriseDingtalkGroupsResponseBody {
	s.Data = v
	return s
}

type ListEnterpriseDingtalkGroupsResponseBodyData struct {
	// 钉群ID
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
	// 钉群名
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ListEnterpriseDingtalkGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupsResponseBodyData) SetOpenGroupId(v string) *ListEnterpriseDingtalkGroupsResponseBodyData {
	s.OpenGroupId = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBodyData) SetGroupName(v string) *ListEnterpriseDingtalkGroupsResponseBodyData {
	s.GroupName = &v
	return s
}

type ListEnterpriseDingtalkGroupsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEnterpriseDingtalkGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListEnterpriseDingtalkGroupsResponse) SetBody(v *ListEnterpriseDingtalkGroupsResponseBody) *ListEnterpriseDingtalkGroupsResponse {
	s.Body = v
	return s
}

type DeleteEnterpriseDingtalkGroupCustomerMemberRequest struct {
	OpenGroupId *string   `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
	Mobiles     []*string `json:"Mobiles,omitempty" xml:"Mobiles,omitempty" type:"Repeated"`
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberRequest) SetOpenGroupId(v string) *DeleteEnterpriseDingtalkGroupCustomerMemberRequest {
	s.OpenGroupId = &v
	return s
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberRequest) SetMobiles(v []*string) *DeleteEnterpriseDingtalkGroupCustomerMemberRequest {
	s.Mobiles = v
	return s
}

type DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest struct {
	OpenGroupId   *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
	MobilesShrink *string `json:"Mobiles,omitempty" xml:"Mobiles,omitempty"`
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest) SetOpenGroupId(v string) *DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest {
	s.OpenGroupId = &v
	return s
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest) SetMobilesShrink(v string) *DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest {
	s.MobilesShrink = &v
	return s
}

type DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody struct {
	// 接口请求的唯一ID, 每次调用requestID唯一
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用接口返回是否成功, true代表调用正常
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误信息, 当success=false的时候, 可以取到message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 接口请求结果返回码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody) SetRequestId(v string) *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody) SetSuccess(v bool) *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody) SetMessage(v string) *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody) SetCode(v string) *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.Code = &v
	return s
}

type DeleteEnterpriseDingtalkGroupCustomerMemberResponse struct {
	Headers map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberResponse) SetHeaders(v map[string]*string) *DeleteEnterpriseDingtalkGroupCustomerMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberResponse) SetBody(v *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody) *DeleteEnterpriseDingtalkGroupCustomerMemberResponse {
	s.Body = v
	return s
}

type GetEnterpriseDingtalkGroupCustomerMemberRequest struct {
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
}

func (s GetEnterpriseDingtalkGroupCustomerMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDingtalkGroupCustomerMemberRequest) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberRequest) SetOpenGroupId(v string) *GetEnterpriseDingtalkGroupCustomerMemberRequest {
	s.OpenGroupId = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberRequest) SetMobile(v string) *GetEnterpriseDingtalkGroupCustomerMemberRequest {
	s.Mobile = &v
	return s
}

type GetEnterpriseDingtalkGroupCustomerMemberResponseBody struct {
	// 接口请求的唯一ID, 每次调用requestID唯一
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用接口返回是否成功, true代表调用正常
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误信息, 当success=false的时候, 可以取到message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 接口请求结果返回码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 成员信息列表
	Data *EnterpriseDingtalkGroupMember `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetEnterpriseDingtalkGroupCustomerMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDingtalkGroupCustomerMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponseBody) SetRequestId(v string) *GetEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponseBody) SetSuccess(v bool) *GetEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.Success = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponseBody) SetMessage(v string) *GetEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.Message = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponseBody) SetCode(v string) *GetEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.Code = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponseBody) SetData(v *EnterpriseDingtalkGroupMember) *GetEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.Data = v
	return s
}

type GetEnterpriseDingtalkGroupCustomerMemberResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEnterpriseDingtalkGroupCustomerMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnterpriseDingtalkGroupCustomerMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDingtalkGroupCustomerMemberResponse) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponse) SetHeaders(v map[string]*string) *GetEnterpriseDingtalkGroupCustomerMemberResponse {
	s.Headers = v
	return s
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponse) SetBody(v *GetEnterpriseDingtalkGroupCustomerMemberResponseBody) *GetEnterpriseDingtalkGroupCustomerMemberResponse {
	s.Body = v
	return s
}

type GetEnterpriseDingtalkGroupRequest struct {
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
}

func (s GetEnterpriseDingtalkGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDingtalkGroupRequest) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDingtalkGroupRequest) SetOpenGroupId(v string) *GetEnterpriseDingtalkGroupRequest {
	s.OpenGroupId = &v
	return s
}

type GetEnterpriseDingtalkGroupResponseBody struct {
	// 接口请求的唯一ID, 每次调用requestID唯一
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用接口返回是否成功, true代表调用正常
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误信息, 当success=false的时候, 可以取到message
	Message *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Data    *GetEnterpriseDingtalkGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 接口请求结果返回码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetEnterpriseDingtalkGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDingtalkGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDingtalkGroupResponseBody) SetRequestId(v string) *GetEnterpriseDingtalkGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupResponseBody) SetSuccess(v bool) *GetEnterpriseDingtalkGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupResponseBody) SetMessage(v string) *GetEnterpriseDingtalkGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupResponseBody) SetData(v *GetEnterpriseDingtalkGroupResponseBodyData) *GetEnterpriseDingtalkGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetEnterpriseDingtalkGroupResponseBody) SetCode(v string) *GetEnterpriseDingtalkGroupResponseBody {
	s.Code = &v
	return s
}

type GetEnterpriseDingtalkGroupResponseBodyData struct {
	// 企业服务群的ID
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
	// 企业服务群的群名
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetEnterpriseDingtalkGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDingtalkGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDingtalkGroupResponseBodyData) SetOpenGroupId(v string) *GetEnterpriseDingtalkGroupResponseBodyData {
	s.OpenGroupId = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupResponseBodyData) SetGroupName(v string) *GetEnterpriseDingtalkGroupResponseBodyData {
	s.GroupName = &v
	return s
}

type GetEnterpriseDingtalkGroupResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEnterpriseDingtalkGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnterpriseDingtalkGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDingtalkGroupResponse) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDingtalkGroupResponse) SetHeaders(v map[string]*string) *GetEnterpriseDingtalkGroupResponse {
	s.Headers = v
	return s
}

func (s *GetEnterpriseDingtalkGroupResponse) SetBody(v *GetEnterpriseDingtalkGroupResponseBody) *GetEnterpriseDingtalkGroupResponse {
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

func (client *Client) ListEnterpriseDingtalkGroupCustomerMembersWithOptions(request *ListEnterpriseDingtalkGroupCustomerMembersRequest, runtime *util.RuntimeOptions) (_result *ListEnterpriseDingtalkGroupCustomerMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListEnterpriseDingtalkGroupCustomerMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListEnterpriseDingtalkGroupCustomerMembers"), tea.String("2021-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListEnterpriseDingtalkGroupsWithOptions(runtime *util.RuntimeOptions) (_result *ListEnterpriseDingtalkGroupsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListEnterpriseDingtalkGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListEnterpriseDingtalkGroups"), tea.String("2021-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteEnterpriseDingtalkGroupCustomerMemberWithOptions(tmpReq *DeleteEnterpriseDingtalkGroupCustomerMemberRequest, runtime *util.RuntimeOptions) (_result *DeleteEnterpriseDingtalkGroupCustomerMemberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Mobiles)) {
		request.MobilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Mobiles, tea.String("Mobiles"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEnterpriseDingtalkGroupCustomerMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEnterpriseDingtalkGroupCustomerMember"), tea.String("2021-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEnterpriseDingtalkGroupCustomerMember(request *DeleteEnterpriseDingtalkGroupCustomerMemberRequest) (_result *DeleteEnterpriseDingtalkGroupCustomerMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEnterpriseDingtalkGroupCustomerMemberResponse{}
	_body, _err := client.DeleteEnterpriseDingtalkGroupCustomerMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnterpriseDingtalkGroupCustomerMemberWithOptions(request *GetEnterpriseDingtalkGroupCustomerMemberRequest, runtime *util.RuntimeOptions) (_result *GetEnterpriseDingtalkGroupCustomerMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetEnterpriseDingtalkGroupCustomerMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetEnterpriseDingtalkGroupCustomerMember"), tea.String("2021-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnterpriseDingtalkGroupCustomerMember(request *GetEnterpriseDingtalkGroupCustomerMemberRequest) (_result *GetEnterpriseDingtalkGroupCustomerMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEnterpriseDingtalkGroupCustomerMemberResponse{}
	_body, _err := client.GetEnterpriseDingtalkGroupCustomerMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnterpriseDingtalkGroupWithOptions(request *GetEnterpriseDingtalkGroupRequest, runtime *util.RuntimeOptions) (_result *GetEnterpriseDingtalkGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetEnterpriseDingtalkGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetEnterpriseDingtalkGroup"), tea.String("2021-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnterpriseDingtalkGroup(request *GetEnterpriseDingtalkGroupRequest) (_result *GetEnterpriseDingtalkGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEnterpriseDingtalkGroupResponse{}
	_body, _err := client.GetEnterpriseDingtalkGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
