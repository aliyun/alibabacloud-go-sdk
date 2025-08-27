// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 换取accessToken接口
//
// @param request - AccessTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccessTokenResponse
func (client *Client) AccessTokenWithContext(ctx context.Context, request *AccessTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AccessTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppSecret) {
		query["app_secret"] = request.AppSecret
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AccessToken"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/btrip-open-auth/v1/access-token/action/take"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AccessTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建企业部门
//
// @param tmpReq - AddDepartmentRequest
//
// @param headers - AddDepartmentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDepartmentResponse
func (client *Client) AddDepartmentWithContext(ctx context.Context, tmpReq *AddDepartmentRequest, headers *AddDepartmentHeaders, runtime *dara.RuntimeOptions) (_result *AddDepartmentResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AddDepartmentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ManagerEmployeeIdList) {
		request.ManagerEmployeeIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ManagerEmployeeIdList, dara.String("manager_employee_id_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeptName) {
		body["dept_name"] = request.DeptName
	}

	if !dara.IsNil(request.ManagerEmployeeIdListShrink) {
		body["manager_employee_id_list"] = request.ManagerEmployeeIdListShrink
	}

	if !dara.IsNil(request.OutDeptId) {
		body["out_dept_id"] = request.OutDeptId
	}

	if !dara.IsNil(request.OutDeptPid) {
		body["out_dept_pid"] = request.OutDeptPid
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDepartment"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/department/v2/add"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDepartmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加员工
//
// @param tmpReq - AddEmployeeRequest
//
// @param headers - AddEmployeeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEmployeeResponse
func (client *Client) AddEmployeeWithContext(ctx context.Context, tmpReq *AddEmployeeRequest, headers *AddEmployeeHeaders, runtime *dara.RuntimeOptions) (_result *AddEmployeeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AddEmployeeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BaseCityCodeList) {
		request.BaseCityCodeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BaseCityCodeList, dara.String("base_city_code_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.BaseLocationList) {
		request.BaseLocationListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BaseLocationList, dara.String("base_location_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CertList) {
		request.CertListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CertList, dara.String("cert_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomRoleCodeList) {
		request.CustomRoleCodeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomRoleCodeList, dara.String("custom_role_code_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OutDeptIdList) {
		request.OutDeptIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutDeptIdList, dara.String("out_dept_id_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountEmail) {
		body["account_email"] = request.AccountEmail
	}

	if !dara.IsNil(request.AccountPhone) {
		body["account_phone"] = request.AccountPhone
	}

	if !dara.IsNil(request.Attribute) {
		body["attribute"] = request.Attribute
	}

	if !dara.IsNil(request.Avatar) {
		body["avatar"] = request.Avatar
	}

	if !dara.IsNil(request.BaseCityCodeListShrink) {
		body["base_city_code_list"] = request.BaseCityCodeListShrink
	}

	if !dara.IsNil(request.BaseLocationListShrink) {
		body["base_location_list"] = request.BaseLocationListShrink
	}

	if !dara.IsNil(request.Birthday) {
		body["birthday"] = request.Birthday
	}

	if !dara.IsNil(request.CertListShrink) {
		body["cert_list"] = request.CertListShrink
	}

	if !dara.IsNil(request.CustomRoleCodeListShrink) {
		body["custom_role_code_list"] = request.CustomRoleCodeListShrink
	}

	if !dara.IsNil(request.Email) {
		body["email"] = request.Email
	}

	if !dara.IsNil(request.Gender) {
		body["gender"] = request.Gender
	}

	if !dara.IsNil(request.IsAdmin) {
		body["is_admin"] = request.IsAdmin
	}

	if !dara.IsNil(request.IsBoss) {
		body["is_boss"] = request.IsBoss
	}

	if !dara.IsNil(request.IsDeptLeader) {
		body["is_dept_leader"] = request.IsDeptLeader
	}

	if !dara.IsNil(request.JobNo) {
		body["job_no"] = request.JobNo
	}

	if !dara.IsNil(request.ManagerUserId) {
		body["manager_user_id"] = request.ManagerUserId
	}

	if !dara.IsNil(request.OutDeptIdListShrink) {
		body["out_dept_id_list"] = request.OutDeptIdListShrink
	}

	if !dara.IsNil(request.Phone) {
		body["phone"] = request.Phone
	}

	if !dara.IsNil(request.PositionLevel) {
		body["position_level"] = request.PositionLevel
	}

	if !dara.IsNil(request.RealName) {
		body["real_name"] = request.RealName
	}

	if !dara.IsNil(request.RealNameEn) {
		body["real_name_en"] = request.RealNameEn
	}

	if !dara.IsNil(request.UnionId) {
		body["union_id"] = request.UnionId
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	if !dara.IsNil(request.UserNick) {
		body["user_nick"] = request.UserNick
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddEmployee"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/employee/v2/add"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddEmployeeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量新增企业自定义角色下人员
//
// @param tmpReq - AddEmployeesToCustomRoleRequest
//
// @param headers - AddEmployeesToCustomRoleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEmployeesToCustomRoleResponse
func (client *Client) AddEmployeesToCustomRoleWithContext(ctx context.Context, tmpReq *AddEmployeesToCustomRoleRequest, headers *AddEmployeesToCustomRoleHeaders, runtime *dara.RuntimeOptions) (_result *AddEmployeesToCustomRoleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AddEmployeesToCustomRoleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserIdList) {
		request.UserIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIdList, dara.String("user_id_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RoleId) {
		body["role_id"] = request.RoleId
	}

	if !dara.IsNil(request.UserIdListShrink) {
		body["user_id_list"] = request.UserIdListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddEmployeesToCustomRole"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/role/v1/customRoleEmployees/add"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddEmployeesToCustomRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增发票抬头适用人员
//
// @param tmpReq - AddInvoiceEntityRequest
//
// @param headers - AddInvoiceEntityHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddInvoiceEntityResponse
func (client *Client) AddInvoiceEntityWithContext(ctx context.Context, tmpReq *AddInvoiceEntityRequest, headers *AddInvoiceEntityHeaders, runtime *dara.RuntimeOptions) (_result *AddInvoiceEntityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AddInvoiceEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Entities) {
		request.EntitiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Entities, dara.String("entities"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EntitiesShrink) {
		body["entities"] = request.EntitiesShrink
	}

	if !dara.IsNil(request.ThirdPartId) {
		body["third_part_id"] = request.ThirdPartId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddInvoiceEntity"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/invoice/v1/entities"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddInvoiceEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 商旅功能页跳转
//
// @param request - AddressGetRequest
//
// @param headers - AddressGetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddressGetResponse
func (client *Client) AddressGetWithContext(ctx context.Context, request *AddressGetRequest, headers *AddressGetHeaders, runtime *dara.RuntimeOptions) (_result *AddressGetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		query["action_type"] = request.ActionType
	}

	if !dara.IsNil(request.ArrCityCode) {
		query["arr_city_code"] = request.ArrCityCode
	}

	if !dara.IsNil(request.ArrCityName) {
		query["arr_city_name"] = request.ArrCityName
	}

	if !dara.IsNil(request.CarScenesCode) {
		query["car_scenes_code"] = request.CarScenesCode
	}

	if !dara.IsNil(request.DepCityCode) {
		query["dep_city_code"] = request.DepCityCode
	}

	if !dara.IsNil(request.DepCityName) {
		query["dep_city_name"] = request.DepCityName
	}

	if !dara.IsNil(request.DepDate) {
		query["dep_date"] = request.DepDate
	}

	if !dara.IsNil(request.ItineraryId) {
		query["itinerary_id"] = request.ItineraryId
	}

	if !dara.IsNil(request.MiddlePage) {
		query["middle_page"] = request.MiddlePage
	}

	if !dara.IsNil(request.OrderId) {
		query["order_Id"] = request.OrderId
	}

	if !dara.IsNil(request.Phone) {
		query["phone"] = request.Phone
	}

	if !dara.IsNil(request.SessionParameters) {
		query["session_parameters"] = request.SessionParameters
	}

	if !dara.IsNil(request.SubCorpId) {
		query["sub_corp_id"] = request.SubCorpId
	}

	if !dara.IsNil(request.TaobaoCallbackUrl) {
		query["taobao_callback_url"] = request.TaobaoCallbackUrl
	}

	if !dara.IsNil(request.ThirdpartApplyId) {
		query["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !dara.IsNil(request.TravelerId) {
		query["traveler_id"] = request.TravelerId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	if !dara.IsNil(request.UseBookingProxy) {
		query["use_booking_proxy"] = request.UseBookingProxy
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddressGet"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/v1/address"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddressGetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询机场数据
//
// @param request - AirportSearchRequest
//
// @param headers - AirportSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AirportSearchResponse
func (client *Client) AirportSearchWithContext(ctx context.Context, request *AirportSearchRequest, headers *AirportSearchHeaders, runtime *dara.RuntimeOptions) (_result *AirportSearchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AirportSearch"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/city/v1/airport"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AirportSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 全量查询商旅城市行政区划编码信息
//
// @param headers - AllBaseCityInfoQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllBaseCityInfoQueryResponse
func (client *Client) AllBaseCityInfoQueryWithContext(ctx context.Context, headers *AllBaseCityInfoQueryHeaders, runtime *dara.RuntimeOptions) (_result *AllBaseCityInfoQueryResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripAccessToken) {
		realHeaders["x-acs-btrip-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripAccessToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllBaseCityInfoQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/city/v1/code"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AllBaseCityInfoQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建出差审批单
//
// @param tmpReq - ApplyAddRequest
//
// @param headers - ApplyAddHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyAddResponse
func (client *Client) ApplyAddWithContext(ctx context.Context, tmpReq *ApplyAddRequest, headers *ApplyAddHeaders, runtime *dara.RuntimeOptions) (_result *ApplyAddResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ApplyAddShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CarRule) {
		request.CarRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CarRule, dara.String("car_rule"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DefaultStandard) {
		request.DefaultStandardShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DefaultStandard, dara.String("default_standard"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExternalTravelerList) {
		request.ExternalTravelerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExternalTravelerList, dara.String("external_traveler_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExternalTravelerStandard) {
		request.ExternalTravelerStandardShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExternalTravelerStandard, dara.String("external_traveler_standard"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HotelShare) {
		request.HotelShareShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelShare, dara.String("hotel_share"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ItineraryList) {
		request.ItineraryListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItineraryList, dara.String("itinerary_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ItinerarySetList) {
		request.ItinerarySetListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItinerarySetList, dara.String("itinerary_set_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TravelerList) {
		request.TravelerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TravelerList, dara.String("traveler_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TravelerStandard) {
		request.TravelerStandardShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TravelerStandard, dara.String("traveler_standard"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Budget) {
		body["budget"] = request.Budget
	}

	if !dara.IsNil(request.BudgetMerge) {
		body["budget_merge"] = request.BudgetMerge
	}

	if !dara.IsNil(request.CarRuleShrink) {
		body["car_rule"] = request.CarRuleShrink
	}

	if !dara.IsNil(request.CorpName) {
		body["corp_name"] = request.CorpName
	}

	if !dara.IsNil(request.DefaultStandardShrink) {
		body["default_standard"] = request.DefaultStandardShrink
	}

	if !dara.IsNil(request.DepartId) {
		body["depart_id"] = request.DepartId
	}

	if !dara.IsNil(request.DepartName) {
		body["depart_name"] = request.DepartName
	}

	if !dara.IsNil(request.ExtendField) {
		body["extend_field"] = request.ExtendField
	}

	if !dara.IsNil(request.ExternalTravelerListShrink) {
		body["external_traveler_list"] = request.ExternalTravelerListShrink
	}

	if !dara.IsNil(request.ExternalTravelerStandardShrink) {
		body["external_traveler_standard"] = request.ExternalTravelerStandardShrink
	}

	if !dara.IsNil(request.FlightBudget) {
		body["flight_budget"] = request.FlightBudget
	}

	if !dara.IsNil(request.HotelBudget) {
		body["hotel_budget"] = request.HotelBudget
	}

	if !dara.IsNil(request.HotelShareShrink) {
		body["hotel_share"] = request.HotelShareShrink
	}

	if !dara.IsNil(request.InternationalFlightCabins) {
		body["international_flight_cabins"] = request.InternationalFlightCabins
	}

	if !dara.IsNil(request.IntlFlightBudget) {
		body["intl_flight_budget"] = request.IntlFlightBudget
	}

	if !dara.IsNil(request.IntlHotelBudget) {
		body["intl_hotel_budget"] = request.IntlHotelBudget
	}

	if !dara.IsNil(request.ItineraryListShrink) {
		body["itinerary_list"] = request.ItineraryListShrink
	}

	if !dara.IsNil(request.ItineraryRule) {
		body["itinerary_rule"] = request.ItineraryRule
	}

	if !dara.IsNil(request.ItinerarySetListShrink) {
		body["itinerary_set_list"] = request.ItinerarySetListShrink
	}

	if !dara.IsNil(request.LimitTraveler) {
		body["limit_traveler"] = request.LimitTraveler
	}

	if !dara.IsNil(request.MealBudget) {
		body["meal_budget"] = request.MealBudget
	}

	if !dara.IsNil(request.PaymentDepartmentId) {
		body["payment_department_id"] = request.PaymentDepartmentId
	}

	if !dara.IsNil(request.PaymentDepartmentName) {
		body["payment_department_name"] = request.PaymentDepartmentName
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.SubCorpId) {
		body["sub_corp_id"] = request.SubCorpId
	}

	if !dara.IsNil(request.ThirdpartApplyId) {
		body["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !dara.IsNil(request.ThirdpartBusinessId) {
		body["thirdpart_business_id"] = request.ThirdpartBusinessId
	}

	if !dara.IsNil(request.ThirdpartDepartId) {
		body["thirdpart_depart_id"] = request.ThirdpartDepartId
	}

	if !dara.IsNil(request.TogetherBookRule) {
		body["together_book_rule"] = request.TogetherBookRule
	}

	if !dara.IsNil(request.TrainBudget) {
		body["train_budget"] = request.TrainBudget
	}

	if !dara.IsNil(request.TravelerListShrink) {
		body["traveler_list"] = request.TravelerListShrink
	}

	if !dara.IsNil(request.TravelerStandardShrink) {
		body["traveler_standard"] = request.TravelerStandardShrink
	}

	if !dara.IsNil(request.TripCause) {
		body["trip_cause"] = request.TripCause
	}

	if !dara.IsNil(request.TripDay) {
		body["trip_day"] = request.TripDay
	}

	if !dara.IsNil(request.TripTitle) {
		body["trip_title"] = request.TripTitle
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	if !dara.IsNil(request.UnionNo) {
		body["union_no"] = request.UnionNo
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		body["user_name"] = request.UserName
	}

	if !dara.IsNil(request.VehicleBudget) {
		body["vehicle_budget"] = request.VehicleBudget
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyAdd"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/biz-trip"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyAddResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新出差审批单（状态）
//
// @param request - ApplyApproveRequest
//
// @param headers - ApplyApproveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyApproveResponse
func (client *Client) ApplyApproveWithContext(ctx context.Context, request *ApplyApproveRequest, headers *ApplyApproveHeaders, runtime *dara.RuntimeOptions) (_result *ApplyApproveResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplyId) {
		body["apply_id"] = request.ApplyId
	}

	if !dara.IsNil(request.Note) {
		body["note"] = request.Note
	}

	if !dara.IsNil(request.OperateTime) {
		body["operate_time"] = request.OperateTime
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.SubCorpId) {
		body["sub_corp_id"] = request.SubCorpId
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		body["user_name"] = request.UserName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyApprove"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/biz-trip/action/approve"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyApproveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 外部审批节点状态同步
//
// @param tmpReq - ApplyExternalNodeStatusUpdateRequest
//
// @param headers - ApplyExternalNodeStatusUpdateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyExternalNodeStatusUpdateResponse
func (client *Client) ApplyExternalNodeStatusUpdateWithContext(ctx context.Context, tmpReq *ApplyExternalNodeStatusUpdateRequest, headers *ApplyExternalNodeStatusUpdateHeaders, runtime *dara.RuntimeOptions) (_result *ApplyExternalNodeStatusUpdateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ApplyExternalNodeStatusUpdateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperationRecords) {
		request.OperationRecordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationRecords, dara.String("operation_records"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["node_id"] = request.NodeId
	}

	if !dara.IsNil(request.OperationRecordsShrink) {
		body["operation_records"] = request.OperationRecordsShrink
	}

	if !dara.IsNil(request.ProcessActionResult) {
		body["process_action_result"] = request.ProcessActionResult
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyExternalNodeStatusUpdate"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/external-nodes/action/status-update"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyExternalNodeStatusUpdateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请发票
//
// @param tmpReq - ApplyInvoiceTaskRequest
//
// @param headers - ApplyInvoiceTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyInvoiceTaskResponse
func (client *Client) ApplyInvoiceTaskWithContext(ctx context.Context, tmpReq *ApplyInvoiceTaskRequest, headers *ApplyInvoiceTaskHeaders, runtime *dara.RuntimeOptions) (_result *ApplyInvoiceTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ApplyInvoiceTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InvoiceTaskList) {
		request.InvoiceTaskListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InvoiceTaskList, dara.String("invoice_task_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BillDate) {
		body["bill_date"] = request.BillDate
	}

	if !dara.IsNil(request.InvoiceTaskListShrink) {
		body["invoice_task_list"] = request.InvoiceTaskListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyInvoiceTask"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/invoice/v1/apply-invoice-task"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyInvoiceTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询出差审批单列表
//
// @param request - ApplyListQueryRequest
//
// @param headers - ApplyListQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyListQueryResponse
func (client *Client) ApplyListQueryWithContext(ctx context.Context, request *ApplyListQueryRequest, headers *ApplyListQueryHeaders, runtime *dara.RuntimeOptions) (_result *ApplyListQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllApply) {
		query["all_apply"] = request.AllApply
	}

	if !dara.IsNil(request.DepartId) {
		query["depart_id"] = request.DepartId
	}

	if !dara.IsNil(request.EndTime) {
		query["end_time"] = request.EndTime
	}

	if !dara.IsNil(request.GmtModified) {
		query["gmt_modified"] = request.GmtModified
	}

	if !dara.IsNil(request.OnlyShangLvApply) {
		query["only_shang_lv_apply"] = request.OnlyShangLvApply
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["start_time"] = request.StartTime
	}

	if !dara.IsNil(request.SubCorpId) {
		query["sub_corp_id"] = request.SubCorpId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	if !dara.IsNil(request.UnionNo) {
		query["union_no"] = request.UnionNo
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyListQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/biz-trips"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyListQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新出差审批单
//
// @param tmpReq - ApplyModifyRequest
//
// @param headers - ApplyModifyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyModifyResponse
func (client *Client) ApplyModifyWithContext(ctx context.Context, tmpReq *ApplyModifyRequest, headers *ApplyModifyHeaders, runtime *dara.RuntimeOptions) (_result *ApplyModifyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ApplyModifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CarRule) {
		request.CarRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CarRule, dara.String("car_rule"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DefaultStandard) {
		request.DefaultStandardShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DefaultStandard, dara.String("default_standard"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExternalTravelerList) {
		request.ExternalTravelerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExternalTravelerList, dara.String("external_traveler_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExternalTravelerStandard) {
		request.ExternalTravelerStandardShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExternalTravelerStandard, dara.String("external_traveler_standard"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HotelShare) {
		request.HotelShareShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelShare, dara.String("hotel_share"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ItineraryList) {
		request.ItineraryListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItineraryList, dara.String("itinerary_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ItinerarySetList) {
		request.ItinerarySetListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItinerarySetList, dara.String("itinerary_set_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TravelerList) {
		request.TravelerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TravelerList, dara.String("traveler_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TravelerStandard) {
		request.TravelerStandardShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TravelerStandard, dara.String("traveler_standard"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Budget) {
		body["budget"] = request.Budget
	}

	if !dara.IsNil(request.BudgetMerge) {
		body["budget_merge"] = request.BudgetMerge
	}

	if !dara.IsNil(request.CarRuleShrink) {
		body["car_rule"] = request.CarRuleShrink
	}

	if !dara.IsNil(request.CorpName) {
		body["corp_name"] = request.CorpName
	}

	if !dara.IsNil(request.DefaultStandardShrink) {
		body["default_standard"] = request.DefaultStandardShrink
	}

	if !dara.IsNil(request.DepartId) {
		body["depart_id"] = request.DepartId
	}

	if !dara.IsNil(request.DepartName) {
		body["depart_name"] = request.DepartName
	}

	if !dara.IsNil(request.ExtendField) {
		body["extend_field"] = request.ExtendField
	}

	if !dara.IsNil(request.ExternalTravelerListShrink) {
		body["external_traveler_list"] = request.ExternalTravelerListShrink
	}

	if !dara.IsNil(request.ExternalTravelerStandardShrink) {
		body["external_traveler_standard"] = request.ExternalTravelerStandardShrink
	}

	if !dara.IsNil(request.FlightBudget) {
		body["flight_budget"] = request.FlightBudget
	}

	if !dara.IsNil(request.HotelBudget) {
		body["hotel_budget"] = request.HotelBudget
	}

	if !dara.IsNil(request.HotelShareShrink) {
		body["hotel_share"] = request.HotelShareShrink
	}

	if !dara.IsNil(request.IntlFlightBudget) {
		body["intl_flight_budget"] = request.IntlFlightBudget
	}

	if !dara.IsNil(request.IntlHotelBudget) {
		body["intl_hotel_budget"] = request.IntlHotelBudget
	}

	if !dara.IsNil(request.ItineraryListShrink) {
		body["itinerary_list"] = request.ItineraryListShrink
	}

	if !dara.IsNil(request.ItineraryRule) {
		body["itinerary_rule"] = request.ItineraryRule
	}

	if !dara.IsNil(request.ItinerarySetListShrink) {
		body["itinerary_set_list"] = request.ItinerarySetListShrink
	}

	if !dara.IsNil(request.LimitTraveler) {
		body["limit_traveler"] = request.LimitTraveler
	}

	if !dara.IsNil(request.MealBudget) {
		body["meal_budget"] = request.MealBudget
	}

	if !dara.IsNil(request.PaymentDepartmentId) {
		body["payment_department_id"] = request.PaymentDepartmentId
	}

	if !dara.IsNil(request.PaymentDepartmentName) {
		body["payment_department_name"] = request.PaymentDepartmentName
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.SubCorpId) {
		body["sub_corp_id"] = request.SubCorpId
	}

	if !dara.IsNil(request.ThirdpartApplyId) {
		body["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !dara.IsNil(request.ThirdpartBusinessId) {
		body["thirdpart_business_id"] = request.ThirdpartBusinessId
	}

	if !dara.IsNil(request.ThirdpartDepartId) {
		body["thirdpart_depart_id"] = request.ThirdpartDepartId
	}

	if !dara.IsNil(request.TogetherBookRule) {
		body["together_book_rule"] = request.TogetherBookRule
	}

	if !dara.IsNil(request.TrainBudget) {
		body["train_budget"] = request.TrainBudget
	}

	if !dara.IsNil(request.TravelerListShrink) {
		body["traveler_list"] = request.TravelerListShrink
	}

	if !dara.IsNil(request.TravelerStandardShrink) {
		body["traveler_standard"] = request.TravelerStandardShrink
	}

	if !dara.IsNil(request.TripCause) {
		body["trip_cause"] = request.TripCause
	}

	if !dara.IsNil(request.TripDay) {
		body["trip_day"] = request.TripDay
	}

	if !dara.IsNil(request.TripTitle) {
		body["trip_title"] = request.TripTitle
	}

	if !dara.IsNil(request.UnionNo) {
		body["union_no"] = request.UnionNo
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		body["user_name"] = request.UserName
	}

	if !dara.IsNil(request.VehicleBudget) {
		body["vehicle_budget"] = request.VehicleBudget
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyModify"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/biz-trip"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyModifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询出差审批单详情
//
// @param request - ApplyQueryRequest
//
// @param headers - ApplyQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyQueryResponse
func (client *Client) ApplyQueryWithContext(ctx context.Context, request *ApplyQueryRequest, headers *ApplyQueryHeaders, runtime *dara.RuntimeOptions) (_result *ApplyQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyId) {
		query["apply_id"] = request.ApplyId
	}

	if !dara.IsNil(request.ApplyShowId) {
		query["apply_show_id"] = request.ApplyShowId
	}

	if !dara.IsNil(request.SubCorpId) {
		query["sub_corp_id"] = request.SubCorpId
	}

	if !dara.IsNil(request.ThirdpartApplyId) {
		query["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/biz-trip"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行审批任务
//
// @param request - ApplyTripTaskExecuteRequest
//
// @param headers - ApplyTripTaskExecuteHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyTripTaskExecuteResponse
func (client *Client) ApplyTripTaskExecuteWithContext(ctx context.Context, request *ApplyTripTaskExecuteRequest, headers *ApplyTripTaskExecuteHeaders, runtime *dara.RuntimeOptions) (_result *ApplyTripTaskExecuteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ActionFrom) {
		body["action_from"] = request.ActionFrom
	}

	if !dara.IsNil(request.Comment) {
		body["comment"] = request.Comment
	}

	if !dara.IsNil(request.TaskAction) {
		body["task_action"] = request.TaskAction
	}

	if !dara.IsNil(request.TaskId) {
		body["task_id"] = request.TaskId
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		body["user_name"] = request.UserName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyTripTaskExecute"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/trip-task/action/execute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyTripTaskExecuteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索国内/国际（港澳台）城市基础行政区划数据
//
// @param request - BaseCityInfoSearchRequest
//
// @param headers - BaseCityInfoSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BaseCityInfoSearchResponse
func (client *Client) BaseCityInfoSearchWithContext(ctx context.Context, request *BaseCityInfoSearchRequest, headers *BaseCityInfoSearchHeaders, runtime *dara.RuntimeOptions) (_result *BaseCityInfoSearchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripAccessToken) {
		realHeaders["x-acs-btrip-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripAccessToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BaseCityInfoSearch"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/city/v1/cities/action/search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BaseCityInfoSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 商旅账单内容修改
//
// @param request - BtripBillInfoAdjustRequest
//
// @param headers - BtripBillInfoAdjustHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BtripBillInfoAdjustResponse
func (client *Client) BtripBillInfoAdjustWithContext(ctx context.Context, request *BtripBillInfoAdjustRequest, headers *BtripBillInfoAdjustHeaders, runtime *dara.RuntimeOptions) (_result *BtripBillInfoAdjustResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PrimaryId) {
		body["primary_id"] = request.PrimaryId
	}

	if !dara.IsNil(request.ThirdPartCostCenterId) {
		body["third_part_cost_center_id"] = request.ThirdPartCostCenterId
	}

	if !dara.IsNil(request.ThirdPartDepartmentId) {
		body["third_part_department_id"] = request.ThirdPartDepartmentId
	}

	if !dara.IsNil(request.ThirdPartInvoiceId) {
		body["third_part_invoice_id"] = request.ThirdPartInvoiceId
	}

	if !dara.IsNil(request.ThirdPartProjectId) {
		body["third_part_project_id"] = request.ThirdPartProjectId
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BtripBillInfoAdjust"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bill/v1/info/action/adjust"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BtripBillInfoAdjustResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步市内用车审批单
//
// @param tmpReq - CarApplyAddRequest
//
// @param headers - CarApplyAddHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CarApplyAddResponse
func (client *Client) CarApplyAddWithContext(ctx context.Context, tmpReq *CarApplyAddRequest, headers *CarApplyAddHeaders, runtime *dara.RuntimeOptions) (_result *CarApplyAddResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CarApplyAddShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TravelerStandard) {
		request.TravelerStandardShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TravelerStandard, dara.String("traveler_standard"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Cause) {
		body["cause"] = request.Cause
	}

	if !dara.IsNil(request.City) {
		body["city"] = request.City
	}

	if !dara.IsNil(request.CityCodeSet) {
		body["city_code_set"] = request.CityCodeSet
	}

	if !dara.IsNil(request.Date) {
		body["date"] = request.Date
	}

	if !dara.IsNil(request.FinishedDate) {
		body["finished_date"] = request.FinishedDate
	}

	if !dara.IsNil(request.ProjectCode) {
		body["project_code"] = request.ProjectCode
	}

	if !dara.IsNil(request.ProjectName) {
		body["project_name"] = request.ProjectName
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.ThirdPartApplyId) {
		body["third_part_apply_id"] = request.ThirdPartApplyId
	}

	if !dara.IsNil(request.ThirdPartCostCenterId) {
		body["third_part_cost_center_id"] = request.ThirdPartCostCenterId
	}

	if !dara.IsNil(request.ThirdPartInvoiceId) {
		body["third_part_invoice_id"] = request.ThirdPartInvoiceId
	}

	if !dara.IsNil(request.TimesTotal) {
		body["times_total"] = request.TimesTotal
	}

	if !dara.IsNil(request.TimesType) {
		body["times_type"] = request.TimesType
	}

	if !dara.IsNil(request.TimesUsed) {
		body["times_used"] = request.TimesUsed
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	if !dara.IsNil(request.TravelerStandardShrink) {
		body["traveler_standard"] = request.TravelerStandardShrink
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CarApplyAdd"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/car"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CarApplyAddResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新市内用车审批单
//
// @param request - CarApplyModifyRequest
//
// @param headers - CarApplyModifyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CarApplyModifyResponse
func (client *Client) CarApplyModifyWithContext(ctx context.Context, request *CarApplyModifyRequest, headers *CarApplyModifyHeaders, runtime *dara.RuntimeOptions) (_result *CarApplyModifyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OperateTime) {
		body["operate_time"] = request.OperateTime
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.ThirdPartApplyId) {
		body["third_part_apply_id"] = request.ThirdPartApplyId
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CarApplyModify"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/car"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CarApplyModifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询市内用车审批单
//
// @param request - CarApplyQueryRequest
//
// @param headers - CarApplyQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CarApplyQueryResponse
func (client *Client) CarApplyQueryWithContext(ctx context.Context, request *CarApplyQueryRequest, headers *CarApplyQueryHeaders, runtime *dara.RuntimeOptions) (_result *CarApplyQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreatedEndAt) {
		query["created_end_at"] = request.CreatedEndAt
	}

	if !dara.IsNil(request.CreatedStartAt) {
		query["created_start_at"] = request.CreatedStartAt
	}

	if !dara.IsNil(request.PageNumber) {
		query["page_number"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.ThirdPartApplyId) {
		query["third_part_apply_id"] = request.ThirdPartApplyId
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CarApplyQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/car"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CarApplyQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用车记账数据
//
// @param request - CarBillSettlementQueryRequest
//
// @param headers - CarBillSettlementQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CarBillSettlementQueryResponse
func (client *Client) CarBillSettlementQueryWithContext(ctx context.Context, request *CarBillSettlementQueryRequest, headers *CarBillSettlementQueryHeaders, runtime *dara.RuntimeOptions) (_result *CarBillSettlementQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillBatch) {
		query["bill_batch"] = request.BillBatch
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PeriodEnd) {
		query["period_end"] = request.PeriodEnd
	}

	if !dara.IsNil(request.PeriodStart) {
		query["period_start"] = request.PeriodStart
	}

	if !dara.IsNil(request.ScrollId) {
		query["scroll_id"] = request.ScrollId
	}

	if !dara.IsNil(request.ScrollMod) {
		query["scroll_mod"] = request.ScrollMod
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CarBillSettlementQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/car/v1/bill-settlement"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CarBillSettlementQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用车订单列表
//
// @param request - CarOrderListQueryRequest
//
// @param headers - CarOrderListQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CarOrderListQueryResponse
func (client *Client) CarOrderListQueryWithContext(ctx context.Context, request *CarOrderListQueryRequest, headers *CarOrderListQueryHeaders, runtime *dara.RuntimeOptions) (_result *CarOrderListQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllApply) {
		query["all_apply"] = request.AllApply
	}

	if !dara.IsNil(request.ApplyId) {
		query["apply_id"] = request.ApplyId
	}

	if !dara.IsNil(request.DepartId) {
		query["depart_id"] = request.DepartId
	}

	if !dara.IsNil(request.EndTime) {
		query["end_time"] = request.EndTime
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["start_time"] = request.StartTime
	}

	if !dara.IsNil(request.ThirdpartApplyId) {
		query["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !dara.IsNil(request.UpdateEndTime) {
		query["update_end_time"] = request.UpdateEndTime
	}

	if !dara.IsNil(request.UpdateStartTime) {
		query["update_start_time"] = request.UpdateStartTime
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CarOrderListQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/car/v1/order-list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CarOrderListQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用车订单查询
//
// @param request - CarOrderQueryRequest
//
// @param headers - CarOrderQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CarOrderQueryResponse
func (client *Client) CarOrderQueryWithContext(ctx context.Context, request *CarOrderQueryRequest, headers *CarOrderQueryHeaders, runtime *dara.RuntimeOptions) (_result *CarOrderQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.SubOrderId) {
		query["sub_order_id"] = request.SubOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CarOrderQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/car/v1/order"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CarOrderQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业用车场景
//
// @param headers - CarSceneQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CarSceneQueryResponse
func (client *Client) CarSceneQueryWithContext(ctx context.Context, headers *CarSceneQueryHeaders, runtime *dara.RuntimeOptions) (_result *CarSceneQueryResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CarSceneQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/car/v1/scenes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CarSceneQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 渠道商创建企业
//
// @param request - ChannelCorpCreateRequest
//
// @param headers - ChannelCorpCreateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChannelCorpCreateResponse
func (client *Client) ChannelCorpCreateWithContext(ctx context.Context, request *ChannelCorpCreateRequest, headers *ChannelCorpCreateHeaders, runtime *dara.RuntimeOptions) (_result *ChannelCorpCreateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdministratorName) {
		body["administrator_name"] = request.AdministratorName
	}

	if !dara.IsNil(request.AdministratorPhone) {
		body["administrator_phone"] = request.AdministratorPhone
	}

	if !dara.IsNil(request.City) {
		body["city"] = request.City
	}

	if !dara.IsNil(request.CorpName) {
		body["corp_name"] = request.CorpName
	}

	if !dara.IsNil(request.Province) {
		body["province"] = request.Province
	}

	if !dara.IsNil(request.Scope) {
		body["scope"] = request.Scope
	}

	if !dara.IsNil(request.ThirdCorpId) {
		body["third_corp_id"] = request.ThirdCorpId
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChannelCorpCreate"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/corp/v1/channelCorps"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChannelCorpCreateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询行政区划（市，区）基础数据
//
// @param request - CitySearchRequest
//
// @param headers - CitySearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CitySearchResponse
func (client *Client) CitySearchWithContext(ctx context.Context, request *CitySearchRequest, headers *CitySearchHeaders, runtime *dara.RuntimeOptions) (_result *CitySearchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CitySearch"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/city/v1/city"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CitySearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询退改审批信息
//
// @param request - CommonApplyQueryRequest
//
// @param headers - CommonApplyQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommonApplyQueryResponse
func (client *Client) CommonApplyQueryWithContext(ctx context.Context, request *CommonApplyQueryRequest, headers *CommonApplyQueryHeaders, runtime *dara.RuntimeOptions) (_result *CommonApplyQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyId) {
		query["apply_id"] = request.ApplyId
	}

	if !dara.IsNil(request.BizCategory) {
		query["biz_category"] = request.BizCategory
	}

	if !dara.IsNil(request.BusinessInstanceId) {
		query["business_instance_id"] = request.BusinessInstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CommonApplyQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/common"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CommonApplyQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 退改审批结果同步
//
// @param request - CommonApplySyncRequest
//
// @param headers - CommonApplySyncHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommonApplySyncResponse
func (client *Client) CommonApplySyncWithContext(ctx context.Context, request *CommonApplySyncRequest, headers *CommonApplySyncHeaders, runtime *dara.RuntimeOptions) (_result *CommonApplySyncResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyId) {
		query["apply_id"] = request.ApplyId
	}

	if !dara.IsNil(request.BizCategory) {
		query["biz_category"] = request.BizCategory
	}

	if !dara.IsNil(request.Remark) {
		query["remark"] = request.Remark
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.ThirdpartyFlowId) {
		query["thirdparty_flow_id"] = request.ThirdpartyFlowId
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CommonApplySync"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/syn-common"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CommonApplySyncResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务商机票记账数据
//
// @param request - CooperatorFlightBillSettlementQueryRequest
//
// @param headers - CooperatorFlightBillSettlementQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CooperatorFlightBillSettlementQueryResponse
func (client *Client) CooperatorFlightBillSettlementQueryWithContext(ctx context.Context, request *CooperatorFlightBillSettlementQueryRequest, headers *CooperatorFlightBillSettlementQueryHeaders, runtime *dara.RuntimeOptions) (_result *CooperatorFlightBillSettlementQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillBatch) {
		query["bill_batch"] = request.BillBatch
	}

	if !dara.IsNil(request.CooperatorId) {
		query["cooperator_id"] = request.CooperatorId
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PeriodEnd) {
		query["period_end"] = request.PeriodEnd
	}

	if !dara.IsNil(request.PeriodStart) {
		query["period_start"] = request.PeriodStart
	}

	if !dara.IsNil(request.ScrollId) {
		query["scroll_id"] = request.ScrollId
	}

	if !dara.IsNil(request.ScrollMod) {
		query["scroll_mod"] = request.ScrollMod
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CooperatorFlightBillSettlementQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cooperator-flight/v1/bill-settlement"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CooperatorFlightBillSettlementQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务商酒店记账数据
//
// @param request - CooperatorHotelBillSettlementQueryRequest
//
// @param headers - CooperatorHotelBillSettlementQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CooperatorHotelBillSettlementQueryResponse
func (client *Client) CooperatorHotelBillSettlementQueryWithContext(ctx context.Context, request *CooperatorHotelBillSettlementQueryRequest, headers *CooperatorHotelBillSettlementQueryHeaders, runtime *dara.RuntimeOptions) (_result *CooperatorHotelBillSettlementQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillBatch) {
		query["bill_batch"] = request.BillBatch
	}

	if !dara.IsNil(request.CooperatorId) {
		query["cooperator_id"] = request.CooperatorId
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PeriodEnd) {
		query["period_end"] = request.PeriodEnd
	}

	if !dara.IsNil(request.PeriodStart) {
		query["period_start"] = request.PeriodStart
	}

	if !dara.IsNil(request.ScrollId) {
		query["scroll_id"] = request.ScrollId
	}

	if !dara.IsNil(request.ScrollMod) {
		query["scroll_mod"] = request.ScrollMod
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CooperatorHotelBillSettlementQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cooperator-hotel/v1/bill-settlement"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CooperatorHotelBillSettlementQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店订单事件推送
//
// @param request - CooperatorHotelEventPushRequest
//
// @param headers - CooperatorHotelEventPushHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CooperatorHotelEventPushResponse
func (client *Client) CooperatorHotelEventPushWithContext(ctx context.Context, request *CooperatorHotelEventPushRequest, headers *CooperatorHotelEventPushHeaders, runtime *dara.RuntimeOptions) (_result *CooperatorHotelEventPushResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChangeOrderStatus) {
		body["change_order_status"] = request.ChangeOrderStatus
	}

	if !dara.IsNil(request.ChangeOrderStatusDesc) {
		body["change_order_status_desc"] = request.ChangeOrderStatusDesc
	}

	if !dara.IsNil(request.CooperatorOrderId) {
		body["cooperator_order_id"] = request.CooperatorOrderId
	}

	if !dara.IsNil(request.Event) {
		body["event"] = request.Event
	}

	if !dara.IsNil(request.EventDesc) {
		body["event_desc"] = request.EventDesc
	}

	if !dara.IsNil(request.EventTime) {
		body["event_time"] = request.EventTime
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CooperatorHotelEventPush"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/coop-hotel/v1/orders/events"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CooperatorHotelEventPushResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个人支付结果推送
//
// @param request - CooperatorSyncPayStatusRequest
//
// @param headers - CooperatorSyncPayStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CooperatorSyncPayStatusResponse
func (client *Client) CooperatorSyncPayStatusWithContext(ctx context.Context, request *CooperatorSyncPayStatusRequest, headers *CooperatorSyncPayStatusHeaders, runtime *dara.RuntimeOptions) (_result *CooperatorSyncPayStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CooperatorOrderId) {
		body["cooperator_order_id"] = request.CooperatorOrderId
	}

	if !dara.IsNil(request.CooperatorPayNo) {
		body["cooperator_pay_no"] = request.CooperatorPayNo
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.PayStatus) {
		body["pay_status"] = request.PayStatus
	}

	if !dara.IsNil(request.PayTime) {
		body["pay_time"] = request.PayTime
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CooperatorSyncPayStatus"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/coop-pay/v1/cashiers/status"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CooperatorSyncPayStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取关联可调用企业接口
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CorpAuthLinkInfoQueryResponse
func (client *Client) CorpAuthLinkInfoQueryWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CorpAuthLinkInfoQueryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CorpAuthLinkInfoQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/corp-authority-link/v1/info"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CorpAuthLinkInfoQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 换取CorpToken接口
//
// @param request - CorpTokenRequest
//
// @param headers - CorpTokenHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CorpTokenResponse
func (client *Client) CorpTokenWithContext(ctx context.Context, request *CorpTokenRequest, headers *CorpTokenHeaders, runtime *dara.RuntimeOptions) (_result *CorpTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppSecret) {
		query["app_secret"] = request.AppSecret
	}

	if !dara.IsNil(request.CorpId) {
		query["corp_id"] = request.CorpId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripAccessToken) {
		realHeaders["x-acs-btrip-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripAccessToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CorpToken"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/btrip-open-auth/v1/corp-token/action/take"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CorpTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除成本中心
//
// @param request - CostCenterDeleteRequest
//
// @param headers - CostCenterDeleteHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CostCenterDeleteResponse
func (client *Client) CostCenterDeleteWithContext(ctx context.Context, request *CostCenterDeleteRequest, headers *CostCenterDeleteHeaders, runtime *dara.RuntimeOptions) (_result *CostCenterDeleteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ThirdpartId) {
		query["thirdpart_id"] = request.ThirdpartId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CostCenterDelete"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/costcenter/v1/delete-costcenter"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CostCenterDeleteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改成本中心
//
// @param request - CostCenterModifyRequest
//
// @param headers - CostCenterModifyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CostCenterModifyResponse
func (client *Client) CostCenterModifyWithContext(ctx context.Context, request *CostCenterModifyRequest, headers *CostCenterModifyHeaders, runtime *dara.RuntimeOptions) (_result *CostCenterModifyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlipayNo) {
		body["alipay_no"] = request.AlipayNo
	}

	if !dara.IsNil(request.Disable) {
		body["disable"] = request.Disable
	}

	if !dara.IsNil(request.Number) {
		body["number"] = request.Number
	}

	if !dara.IsNil(request.Scope) {
		body["scope"] = request.Scope
	}

	if !dara.IsNil(request.ThirdpartId) {
		body["thirdpart_id"] = request.ThirdpartId
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CostCenterModify"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/costcenter/v1/modify-costcenter"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CostCenterModifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看成本中心
//
// @param request - CostCenterQueryRequest
//
// @param headers - CostCenterQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CostCenterQueryResponse
func (client *Client) CostCenterQueryWithContext(ctx context.Context, request *CostCenterQueryRequest, headers *CostCenterQueryHeaders, runtime *dara.RuntimeOptions) (_result *CostCenterQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Disable) {
		query["disable"] = request.Disable
	}

	if !dara.IsNil(request.NeedOrgEntity) {
		query["need_org_entity"] = request.NeedOrgEntity
	}

	if !dara.IsNil(request.ThirdpartId) {
		query["thirdpart_id"] = request.ThirdpartId
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CostCenterQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/costcenter/v1/costcenter"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CostCenterQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存成本中心
//
// @param request - CostCenterSaveRequest
//
// @param headers - CostCenterSaveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CostCenterSaveResponse
func (client *Client) CostCenterSaveWithContext(ctx context.Context, request *CostCenterSaveRequest, headers *CostCenterSaveHeaders, runtime *dara.RuntimeOptions) (_result *CostCenterSaveResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlipayNo) {
		body["alipay_no"] = request.AlipayNo
	}

	if !dara.IsNil(request.Disable) {
		body["disable"] = request.Disable
	}

	if !dara.IsNil(request.Number) {
		body["number"] = request.Number
	}

	if !dara.IsNil(request.Scope) {
		body["scope"] = request.Scope
	}

	if !dara.IsNil(request.ThirdpartId) {
		body["thirdpart_id"] = request.ThirdpartId
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CostCenterSave"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/costcenter/v1/save-costcenter"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CostCenterSaveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建企业自定义角色
//
// @param request - CreateCustomRoleRequest
//
// @param headers - CreateCustomRoleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomRoleResponse
func (client *Client) CreateCustomRoleWithContext(ctx context.Context, request *CreateCustomRoleRequest, headers *CreateCustomRoleHeaders, runtime *dara.RuntimeOptions) (_result *CreateCustomRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RoleId) {
		body["role_id"] = request.RoleId
	}

	if !dara.IsNil(request.RoleName) {
		body["role_name"] = request.RoleName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomRole"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/role/v1/customRoles/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建子企业
//
// @param request - CreateSubCorpRequest
//
// @param headers - CreateSubCorpHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubCorpResponse
func (client *Client) CreateSubCorpWithContext(ctx context.Context, request *CreateSubCorpRequest, headers *CreateSubCorpHeaders, runtime *dara.RuntimeOptions) (_result *CreateSubCorpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OuterCorpId) {
		body["outer_corp_id"] = request.OuterCorpId
	}

	if !dara.IsNil(request.OuterCorpName) {
		body["outer_corp_name"] = request.OuterCorpName
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSubCorp"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/sub_corps/v1/corps"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSubCorpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除企业自定义角色
//
// @param request - DeleteCustomRoleRequest
//
// @param headers - DeleteCustomRoleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomRoleResponse
func (client *Client) DeleteCustomRoleWithContext(ctx context.Context, request *DeleteCustomRoleRequest, headers *DeleteCustomRoleHeaders, runtime *dara.RuntimeOptions) (_result *DeleteCustomRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RoleId) {
		body["role_id"] = request.RoleId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomRole"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/role/v1/customRoles/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除企业部门
//
// @param request - DeleteDepartmentRequest
//
// @param headers - DeleteDepartmentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDepartmentResponse
func (client *Client) DeleteDepartmentWithContext(ctx context.Context, request *DeleteDepartmentRequest, headers *DeleteDepartmentHeaders, runtime *dara.RuntimeOptions) (_result *DeleteDepartmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OutDeptId) {
		body["out_dept_id"] = request.OutDeptId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDepartment"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/department/v2/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDepartmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除企业自定义角色下人员
//
// @param tmpReq - DeleteEmployeesFromCustomRoleRequest
//
// @param headers - DeleteEmployeesFromCustomRoleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEmployeesFromCustomRoleResponse
func (client *Client) DeleteEmployeesFromCustomRoleWithContext(ctx context.Context, tmpReq *DeleteEmployeesFromCustomRoleRequest, headers *DeleteEmployeesFromCustomRoleHeaders, runtime *dara.RuntimeOptions) (_result *DeleteEmployeesFromCustomRoleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteEmployeesFromCustomRoleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserIdList) {
		request.UserIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIdList, dara.String("user_id_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RoleId) {
		body["role_id"] = request.RoleId
	}

	if !dara.IsNil(request.UserIdListShrink) {
		body["user_id_list"] = request.UserIdListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEmployeesFromCustomRole"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/role/v1/customRoleEmployees/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEmployeesFromCustomRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除发票抬头适用人员
//
// @param tmpReq - DeleteInvoiceEntityRequest
//
// @param headers - DeleteInvoiceEntityHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInvoiceEntityResponse
func (client *Client) DeleteInvoiceEntityWithContext(ctx context.Context, tmpReq *DeleteInvoiceEntityRequest, headers *DeleteInvoiceEntityHeaders, runtime *dara.RuntimeOptions) (_result *DeleteInvoiceEntityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteInvoiceEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Entities) {
		request.EntitiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Entities, dara.String("entities"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DelAll) {
		query["del_all"] = request.DelAll
	}

	if !dara.IsNil(request.EntitiesShrink) {
		query["entities"] = request.EntitiesShrink
	}

	if !dara.IsNil(request.ThirdPartId) {
		query["third_part_id"] = request.ThirdPartId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInvoiceEntity"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/invoice/v1/entities"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInvoiceEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步外部平台部门信息至商旅内部
//
// @param tmpReq - DepartmentSaveRequest
//
// @param headers - DepartmentSaveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DepartmentSaveResponse
func (client *Client) DepartmentSaveWithContext(ctx context.Context, tmpReq *DepartmentSaveRequest, headers *DepartmentSaveHeaders, runtime *dara.RuntimeOptions) (_result *DepartmentSaveResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DepartmentSaveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DepartList) {
		request.DepartListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepartList, dara.String("depart_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DepartListShrink) {
		body["depart_list"] = request.DepartListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DepartmentSave"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/department/v1/department"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DepartmentSaveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量申请电子行程单
//
// @param tmpReq - ElectronicItineraryBatchApplyRequest
//
// @param headers - ElectronicItineraryBatchApplyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ElectronicItineraryBatchApplyResponse
func (client *Client) ElectronicItineraryBatchApplyWithContext(ctx context.Context, tmpReq *ElectronicItineraryBatchApplyRequest, headers *ElectronicItineraryBatchApplyHeaders, runtime *dara.RuntimeOptions) (_result *ElectronicItineraryBatchApplyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ElectronicItineraryBatchApplyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplyItineraryList) {
		request.ApplyItineraryListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplyItineraryList, dara.String("apply_itinerary_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplyItineraryListShrink) {
		body["apply_itinerary_list"] = request.ApplyItineraryListShrink
	}

	if !dara.IsNil(request.CanReprint) {
		body["can_reprint"] = request.CanReprint
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ElectronicItineraryBatchApply"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/invoice/v1/apply-itinerary-batch-task"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ElectronicItineraryBatchApplyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取电子行程单申请结果
//
// @param request - ElectronicItineraryGetApplyResultRequest
//
// @param headers - ElectronicItineraryGetApplyResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ElectronicItineraryGetApplyResultResponse
func (client *Client) ElectronicItineraryGetApplyResultWithContext(ctx context.Context, request *ElectronicItineraryGetApplyResultRequest, headers *ElectronicItineraryGetApplyResultHeaders, runtime *dara.RuntimeOptions) (_result *ElectronicItineraryGetApplyResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchApplyNo) {
		query["batch_apply_no"] = request.BatchApplyNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ElectronicItineraryGetApplyResult"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/invoice/v1/get-itinerary-batch-task"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ElectronicItineraryGetApplyResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加成本中心人员信息
//
// @param tmpReq - EntityAddRequest
//
// @param headers - EntityAddHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EntityAddResponse
func (client *Client) EntityAddWithContext(ctx context.Context, tmpReq *EntityAddRequest, headers *EntityAddHeaders, runtime *dara.RuntimeOptions) (_result *EntityAddResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &EntityAddShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EntityDOList) {
		request.EntityDOListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityDOList, dara.String("entity_d_o_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityDOListShrink) {
		body["entity_d_o_list"] = request.EntityDOListShrink
	}

	if !dara.IsNil(request.ThirdpartId) {
		body["thirdpart_id"] = request.ThirdpartId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EntityAdd"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/costcenter/v1/add-entity"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EntityAddResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除成本中心人员信息
//
// @param tmpReq - EntityDeleteRequest
//
// @param headers - EntityDeleteHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EntityDeleteResponse
func (client *Client) EntityDeleteWithContext(ctx context.Context, tmpReq *EntityDeleteRequest, headers *EntityDeleteHeaders, runtime *dara.RuntimeOptions) (_result *EntityDeleteResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &EntityDeleteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EntityDOList) {
		request.EntityDOListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityDOList, dara.String("entity_d_o_list"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DelAll) {
		query["del_all"] = request.DelAll
	}

	if !dara.IsNil(request.ThirdpartId) {
		query["thirdpart_id"] = request.ThirdpartId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityDOListShrink) {
		body["entity_d_o_list"] = request.EntityDOListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EntityDelete"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/costcenter/v1/entity/action/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EntityDeleteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置成本中心人员信息
//
// @param tmpReq - EntitySetRequest
//
// @param headers - EntitySetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EntitySetResponse
func (client *Client) EntitySetWithContext(ctx context.Context, tmpReq *EntitySetRequest, headers *EntitySetHeaders, runtime *dara.RuntimeOptions) (_result *EntitySetResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &EntitySetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EntityDOList) {
		request.EntityDOListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityDOList, dara.String("entity_d_o_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityDOListShrink) {
		body["entity_d_o_list"] = request.EntityDOListShrink
	}

	if !dara.IsNil(request.ThirdpartId) {
		body["thirdpart_id"] = request.ThirdpartId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EntitySet"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/costcenter/v1/set-entity"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EntitySetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预估价格查询
//
// @param request - EstimatedPriceQueryRequest
//
// @param headers - EstimatedPriceQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EstimatedPriceQueryResponse
func (client *Client) EstimatedPriceQueryWithContext(ctx context.Context, request *EstimatedPriceQueryRequest, headers *EstimatedPriceQueryHeaders, runtime *dara.RuntimeOptions) (_result *EstimatedPriceQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArrCity) {
		query["arr_city"] = request.ArrCity
	}

	if !dara.IsNil(request.Category) {
		query["category"] = request.Category
	}

	if !dara.IsNil(request.DepCity) {
		query["dep_city"] = request.DepCity
	}

	if !dara.IsNil(request.EndTime) {
		query["end_time"] = request.EndTime
	}

	if !dara.IsNil(request.ItineraryId) {
		query["itinerary_id"] = request.ItineraryId
	}

	if !dara.IsNil(request.StartTime) {
		query["start_time"] = request.StartTime
	}

	if !dara.IsNil(request.SubCorpId) {
		query["sub_corp_id"] = request.SubCorpId
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EstimatedPriceQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/costcenter/v1/estimated-price"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EstimatedPriceQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 超标审批结果同步
//
// @param request - ExceedApplySyncRequest
//
// @param headers - ExceedApplySyncHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExceedApplySyncResponse
func (client *Client) ExceedApplySyncWithContext(ctx context.Context, request *ExceedApplySyncRequest, headers *ExceedApplySyncHeaders, runtime *dara.RuntimeOptions) (_result *ExceedApplySyncResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyId) {
		query["apply_id"] = request.ApplyId
	}

	if !dara.IsNil(request.BizCategory) {
		query["biz_category"] = request.BizCategory
	}

	if !dara.IsNil(request.Remark) {
		query["remark"] = request.Remark
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.ThirdpartyFlowId) {
		query["thirdparty_flow_id"] = request.ThirdpartyFlowId
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExceedApplySync"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/syn-exceed"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExceedApplySyncResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加外部出行人与证件信息
//
// @param tmpReq - ExternalUserAddRequest
//
// @param headers - ExternalUserAddHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExternalUserAddResponse
func (client *Client) ExternalUserAddWithContext(ctx context.Context, tmpReq *ExternalUserAddRequest, headers *ExternalUserAddHeaders, runtime *dara.RuntimeOptions) (_result *ExternalUserAddResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ExternalUserAddShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CertRequestList) {
		request.CertRequestListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CertRequestList, dara.String("cert_request_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Birthday) {
		body["birthday"] = request.Birthday
	}

	if !dara.IsNil(request.CertRequestListShrink) {
		body["cert_request_list"] = request.CertRequestListShrink
	}

	if !dara.IsNil(request.Email) {
		body["email"] = request.Email
	}

	if !dara.IsNil(request.ExternalUserId) {
		body["external_user_id"] = request.ExternalUserId
	}

	if !dara.IsNil(request.Phone) {
		body["phone"] = request.Phone
	}

	if !dara.IsNil(request.RealName) {
		body["real_name"] = request.RealName
	}

	if !dara.IsNil(request.RealNameEn) {
		body["real_name_en"] = request.RealNameEn
	}

	if !dara.IsNil(request.UserType) {
		body["user_type"] = request.UserType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExternalUserAdd"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/user/v1/externalUsers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExternalUserAddResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除外部出行人
//
// @param headers - ExternalUserDeleteHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExternalUserDeleteResponse
func (client *Client) ExternalUserDeleteWithContext(ctx context.Context, externalUserId *string, headers *ExternalUserDeleteHeaders, runtime *dara.RuntimeOptions) (_result *ExternalUserDeleteResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExternalUserDelete"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/user/v1/externalUsers/" + dara.PercentEncode(dara.StringValue(externalUserId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExternalUserDeleteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询外部出行人
//
// @param headers - ExternalUserQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExternalUserQueryResponse
func (client *Client) ExternalUserQueryWithContext(ctx context.Context, externalUserId *string, headers *ExternalUserQueryHeaders, runtime *dara.RuntimeOptions) (_result *ExternalUserQueryResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExternalUserQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/user/v1/externalUsers/" + dara.PercentEncode(dara.StringValue(externalUserId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExternalUserQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改外部出行人与证件信息
//
// @param tmpReq - ExternalUserUpdateRequest
//
// @param headers - ExternalUserUpdateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExternalUserUpdateResponse
func (client *Client) ExternalUserUpdateWithContext(ctx context.Context, externalUserId *string, tmpReq *ExternalUserUpdateRequest, headers *ExternalUserUpdateHeaders, runtime *dara.RuntimeOptions) (_result *ExternalUserUpdateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ExternalUserUpdateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CertRequestList) {
		request.CertRequestListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CertRequestList, dara.String("cert_request_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Birthday) {
		body["birthday"] = request.Birthday
	}

	if !dara.IsNil(request.CertRequestListShrink) {
		body["cert_request_list"] = request.CertRequestListShrink
	}

	if !dara.IsNil(request.Email) {
		body["email"] = request.Email
	}

	if !dara.IsNil(request.Phone) {
		body["phone"] = request.Phone
	}

	if !dara.IsNil(request.RealName) {
		body["real_name"] = request.RealName
	}

	if !dara.IsNil(request.RealNameEn) {
		body["real_name_en"] = request.RealNameEn
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExternalUserUpdate"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/user/v1/externalUsers/" + dara.PercentEncode(dara.StringValue(externalUserId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExternalUserUpdateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询机票记账数据
//
// @param request - FlightBillSettlementQueryRequest
//
// @param headers - FlightBillSettlementQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightBillSettlementQueryResponse
func (client *Client) FlightBillSettlementQueryWithContext(ctx context.Context, request *FlightBillSettlementQueryRequest, headers *FlightBillSettlementQueryHeaders, runtime *dara.RuntimeOptions) (_result *FlightBillSettlementQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillBatch) {
		query["bill_batch"] = request.BillBatch
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PeriodEnd) {
		query["period_end"] = request.PeriodEnd
	}

	if !dara.IsNil(request.PeriodStart) {
		query["period_start"] = request.PeriodStart
	}

	if !dara.IsNil(request.ScrollId) {
		query["scroll_id"] = request.ScrollId
	}

	if !dara.IsNil(request.ScrollMod) {
		query["scroll_mod"] = request.ScrollMod
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightBillSettlementQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/flight/v1/bill-settlement"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightBillSettlementQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航班订单取消
//
// @param request - FlightCancelOrderRequest
//
// @param headers - FlightCancelOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightCancelOrderResponse
func (client *Client) FlightCancelOrderWithContext(ctx context.Context, request *FlightCancelOrderRequest, headers *FlightCancelOrderHeaders, runtime *dara.RuntimeOptions) (_result *FlightCancelOrderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisOrderId) {
		query["dis_order_id"] = request.DisOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightCancelOrder"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/order/action/cancel"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightCancelOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票订单取消
//
// @param request - FlightCancelOrderV2Request
//
// @param headers - FlightCancelOrderV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightCancelOrderV2Response
func (client *Client) FlightCancelOrderV2WithContext(ctx context.Context, request *FlightCancelOrderV2Request, headers *FlightCancelOrderV2Headers, runtime *dara.RuntimeOptions) (_result *FlightCancelOrderV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["out_order_id"] = request.OutOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightCancelOrderV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/order/action/cancel"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightCancelOrderV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航班订单创建
//
// @param tmpReq - FlightCreateOrderRequest
//
// @param headers - FlightCreateOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightCreateOrderResponse
func (client *Client) FlightCreateOrderWithContext(ctx context.Context, tmpReq *FlightCreateOrderRequest, headers *FlightCreateOrderHeaders, runtime *dara.RuntimeOptions) (_result *FlightCreateOrderResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FlightCreateOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ContactInfo) {
		request.ContactInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContactInfo, dara.String("contact_info"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OrderAttr) {
		request.OrderAttrShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderAttr, dara.String("order_attr"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TravelerInfoList) {
		request.TravelerInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TravelerInfoList, dara.String("traveler_info_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ArrAirportCode) {
		body["arr_airport_code"] = request.ArrAirportCode
	}

	if !dara.IsNil(request.ArrCityCode) {
		body["arr_city_code"] = request.ArrCityCode
	}

	if !dara.IsNil(request.AutoPay) {
		body["auto_pay"] = request.AutoPay
	}

	if !dara.IsNil(request.BuyerName) {
		body["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.BuyerUniqueKey) {
		body["buyer_unique_key"] = request.BuyerUniqueKey
	}

	if !dara.IsNil(request.ContactInfoShrink) {
		body["contact_info"] = request.ContactInfoShrink
	}

	if !dara.IsNil(request.DepAirportCode) {
		body["dep_airport_code"] = request.DepAirportCode
	}

	if !dara.IsNil(request.DepCityCode) {
		body["dep_city_code"] = request.DepCityCode
	}

	if !dara.IsNil(request.DepDate) {
		body["dep_date"] = request.DepDate
	}

	if !dara.IsNil(request.DisOrderId) {
		body["dis_order_id"] = request.DisOrderId
	}

	if !dara.IsNil(request.OrderAttrShrink) {
		body["order_attr"] = request.OrderAttrShrink
	}

	if !dara.IsNil(request.OrderParams) {
		body["order_params"] = request.OrderParams
	}

	if !dara.IsNil(request.OtaItemId) {
		body["ota_item_id"] = request.OtaItemId
	}

	if !dara.IsNil(request.Price) {
		body["price"] = request.Price
	}

	if !dara.IsNil(request.ReceiptAddress) {
		body["receipt_address"] = request.ReceiptAddress
	}

	if !dara.IsNil(request.ReceiptTarget) {
		body["receipt_target"] = request.ReceiptTarget
	}

	if !dara.IsNil(request.ReceiptTitle) {
		body["receipt_title"] = request.ReceiptTitle
	}

	if !dara.IsNil(request.TravelerInfoListShrink) {
		body["traveler_info_list"] = request.TravelerInfoListShrink
	}

	if !dara.IsNil(request.TripType) {
		body["trip_type"] = request.TripType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightCreateOrder"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/order/action/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightCreateOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票订单创建
//
// @param tmpReq - FlightCreateOrderV2Request
//
// @param headers - FlightCreateOrderV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightCreateOrderV2Response
func (client *Client) FlightCreateOrderV2WithContext(ctx context.Context, tmpReq *FlightCreateOrderV2Request, headers *FlightCreateOrderV2Headers, runtime *dara.RuntimeOptions) (_result *FlightCreateOrderV2Response, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FlightCreateOrderV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ContactInfo) {
		request.ContactInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContactInfo, dara.String("contact_info"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Travelers) {
		request.TravelersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Travelers, dara.String("travelers"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AsyncCreateOrderKey) {
		body["async_create_order_key"] = request.AsyncCreateOrderKey
	}

	if !dara.IsNil(request.AsyncCreateOrderMode) {
		body["async_create_order_mode"] = request.AsyncCreateOrderMode
	}

	if !dara.IsNil(request.BtripUserId) {
		body["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		body["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.ContactInfoShrink) {
		body["contact_info"] = request.ContactInfoShrink
	}

	if !dara.IsNil(request.IsvName) {
		body["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OtaItemId) {
		body["ota_item_id"] = request.OtaItemId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.TotalPriceCent) {
		body["total_price_cent"] = request.TotalPriceCent
	}

	if !dara.IsNil(request.TravelersShrink) {
		body["travelers"] = request.TravelersShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightCreateOrderV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/order/action/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightCreateOrderV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询飞机超标审批详情
//
// @param request - FlightExceedApplyQueryRequest
//
// @param headers - FlightExceedApplyQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightExceedApplyQueryResponse
func (client *Client) FlightExceedApplyQueryWithContext(ctx context.Context, request *FlightExceedApplyQueryRequest, headers *FlightExceedApplyQueryHeaders, runtime *dara.RuntimeOptions) (_result *FlightExceedApplyQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyId) {
		query["apply_id"] = request.ApplyId
	}

	if !dara.IsNil(request.BusinessInstanceId) {
		query["business_instance_id"] = request.BusinessInstanceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightExceedApplyQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/flight-exceed"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightExceedApplyQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询机票行程单扫描件
//
// @param request - FlightItineraryScanQueryRequest
//
// @param headers - FlightItineraryScanQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightItineraryScanQueryResponse
func (client *Client) FlightItineraryScanQueryWithContext(ctx context.Context, request *FlightItineraryScanQueryRequest, headers *FlightItineraryScanQueryHeaders, runtime *dara.RuntimeOptions) (_result *FlightItineraryScanQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillDate) {
		query["bill_date"] = request.BillDate
	}

	if !dara.IsNil(request.BillId) {
		query["bill_id"] = request.BillId
	}

	if !dara.IsNil(request.InvoiceSubTaskId) {
		query["invoice_sub_task_id"] = request.InvoiceSubTaskId
	}

	if !dara.IsNil(request.ItineraryNum) {
		query["itinerary_num"] = request.ItineraryNum
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.TicketNo) {
		query["ticket_no"] = request.TicketNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightItineraryScanQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/scan/v1/flight-itinerary"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightItineraryScanQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航班列表搜索
//
// @param request - FlightListingSearchRequest
//
// @param headers - FlightListingSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightListingSearchResponse
func (client *Client) FlightListingSearchWithContext(ctx context.Context, request *FlightListingSearchRequest, headers *FlightListingSearchHeaders, runtime *dara.RuntimeOptions) (_result *FlightListingSearchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AirlineCode) {
		query["airline_code"] = request.AirlineCode
	}

	if !dara.IsNil(request.ArrCityCode) {
		query["arr_city_code"] = request.ArrCityCode
	}

	if !dara.IsNil(request.CabinClass) {
		query["cabin_class"] = request.CabinClass
	}

	if !dara.IsNil(request.DepCityCode) {
		query["dep_city_code"] = request.DepCityCode
	}

	if !dara.IsNil(request.DepDate) {
		query["dep_date"] = request.DepDate
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightListingSearch"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/flight/action/listing-search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightListingSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航班列表搜索
//
// @param tmpReq - FlightListingSearchV2Request
//
// @param headers - FlightListingSearchV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightListingSearchV2Response
func (client *Client) FlightListingSearchV2WithContext(ctx context.Context, tmpReq *FlightListingSearchV2Request, headers *FlightListingSearchV2Headers, runtime *dara.RuntimeOptions) (_result *FlightListingSearchV2Response, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FlightListingSearchV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CabinTypeList) {
		request.CabinTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CabinTypeList, dara.String("cabin_type_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SearchJourneys) {
		request.SearchJourneysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchJourneys, dara.String("search_journeys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AirlineCode) {
		query["airline_code"] = request.AirlineCode
	}

	if !dara.IsNil(request.CabinTypeListShrink) {
		query["cabin_type_list"] = request.CabinTypeListShrink
	}

	if !dara.IsNil(request.DirectOnly) {
		query["direct_only"] = request.DirectOnly
	}

	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.NeedMultiClassPrice) {
		query["need_multi_class_price"] = request.NeedMultiClassPrice
	}

	if !dara.IsNil(request.NeedQueryServiceFee) {
		query["need_query_service_fee"] = request.NeedQueryServiceFee
	}

	if !dara.IsNil(request.NeedShareFlight) {
		query["need_share_flight"] = request.NeedShareFlight
	}

	if !dara.IsNil(request.NeedYCBestPrice) {
		query["need_y_c_best_price"] = request.NeedYCBestPrice
	}

	if !dara.IsNil(request.SearchJourneysShrink) {
		query["search_journeys"] = request.SearchJourneysShrink
	}

	if !dara.IsNil(request.SearchMode) {
		query["search_mode"] = request.SearchMode
	}

	if !dara.IsNil(request.TripType) {
		query["trip_type"] = request.TripType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightListingSearchV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/flight/action/listing-search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightListingSearchV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票改签申请
//
// @param tmpReq - FlightModifyApplyV2Request
//
// @param headers - FlightModifyApplyV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightModifyApplyV2Response
func (client *Client) FlightModifyApplyV2WithContext(ctx context.Context, tmpReq *FlightModifyApplyV2Request, headers *FlightModifyApplyV2Headers, runtime *dara.RuntimeOptions) (_result *FlightModifyApplyV2Response, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FlightModifyApplyV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PassengerSegmentRelations) {
		request.PassengerSegmentRelationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PassengerSegmentRelations, dara.String("passenger_segment_relations"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CacheKey) {
		body["cache_key"] = request.CacheKey
	}

	if !dara.IsNil(request.ContactPhone) {
		body["contact_phone"] = request.ContactPhone
	}

	if !dara.IsNil(request.IsvName) {
		body["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.ItemId) {
		body["item_id"] = request.ItemId
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutSubOrderId) {
		body["out_sub_order_id"] = request.OutSubOrderId
	}

	if !dara.IsNil(request.PassengerSegmentRelationsShrink) {
		body["passenger_segment_relations"] = request.PassengerSegmentRelationsShrink
	}

	if !dara.IsNil(request.Reason) {
		body["reason"] = request.Reason
	}

	if !dara.IsNil(request.SessionId) {
		body["session_id"] = request.SessionId
	}

	if !dara.IsNil(request.Voluntary) {
		body["voluntary"] = request.Voluntary
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightModifyApplyV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/modify/action/apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightModifyApplyV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票改签取消
//
// @param request - FlightModifyCancelV2Request
//
// @param headers - FlightModifyCancelV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightModifyCancelV2Response
func (client *Client) FlightModifyCancelV2WithContext(ctx context.Context, request *FlightModifyCancelV2Request, headers *FlightModifyCancelV2Headers, runtime *dara.RuntimeOptions) (_result *FlightModifyCancelV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutSubOrderId) {
		query["out_sub_order_id"] = request.OutSubOrderId
	}

	if !dara.IsNil(request.SubOrderId) {
		query["sub_order_id"] = request.SubOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightModifyCancelV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/modify/action/cancel"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightModifyCancelV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票改签列表搜索
//
// @param tmpReq - FlightModifyListingSearchV2Request
//
// @param headers - FlightModifyListingSearchV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightModifyListingSearchV2Response
func (client *Client) FlightModifyListingSearchV2WithContext(ctx context.Context, tmpReq *FlightModifyListingSearchV2Request, headers *FlightModifyListingSearchV2Headers, runtime *dara.RuntimeOptions) (_result *FlightModifyListingSearchV2Response, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FlightModifyListingSearchV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CabinClass) {
		request.CabinClassShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CabinClass, dara.String("cabin_class"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DepDate) {
		request.DepDateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepDate, dara.String("dep_date"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PassengerSegmentRelations) {
		request.PassengerSegmentRelationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PassengerSegmentRelations, dara.String("passenger_segment_relations"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SelectedSegments) {
		request.SelectedSegmentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SelectedSegments, dara.String("selected_segments"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CabinClassShrink) {
		query["cabin_class"] = request.CabinClassShrink
	}

	if !dara.IsNil(request.DepDateShrink) {
		query["dep_date"] = request.DepDateShrink
	}

	if !dara.IsNil(request.InterfaceCallerIsSupportRetry) {
		query["interface_caller_is_support_retry"] = request.InterfaceCallerIsSupportRetry
	}

	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.PassengerSegmentRelationsShrink) {
		query["passenger_segment_relations"] = request.PassengerSegmentRelationsShrink
	}

	if !dara.IsNil(request.SearchMode) {
		query["search_mode"] = request.SearchMode
	}

	if !dara.IsNil(request.SearchRetryToken) {
		query["search_retry_token"] = request.SearchRetryToken
	}

	if !dara.IsNil(request.SelectedSegmentsShrink) {
		query["selected_segments"] = request.SelectedSegmentsShrink
	}

	if !dara.IsNil(request.SessionId) {
		query["session_id"] = request.SessionId
	}

	if !dara.IsNil(request.Voluntary) {
		query["voluntary"] = request.Voluntary
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightModifyListingSearchV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/modify/action/listing-search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightModifyListingSearchV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票改签详情
//
// @param request - FlightModifyOrderDetailV2Request
//
// @param headers - FlightModifyOrderDetailV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightModifyOrderDetailV2Response
func (client *Client) FlightModifyOrderDetailV2WithContext(ctx context.Context, request *FlightModifyOrderDetailV2Request, headers *FlightModifyOrderDetailV2Headers, runtime *dara.RuntimeOptions) (_result *FlightModifyOrderDetailV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.ModifyApplyId) {
		query["modify_apply_id"] = request.ModifyApplyId
	}

	if !dara.IsNil(request.NeedQueryServiceFee) {
		query["need_query_service_fee"] = request.NeedQueryServiceFee
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutModifyApplyId) {
		query["out_modify_apply_id"] = request.OutModifyApplyId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["out_order_id"] = request.OutOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightModifyOrderDetailV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/modify/action/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightModifyOrderDetailV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票改签报价搜索
//
// @param tmpReq - FlightModifyOtaSearchV2Request
//
// @param headers - FlightModifyOtaSearchV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightModifyOtaSearchV2Response
func (client *Client) FlightModifyOtaSearchV2WithContext(ctx context.Context, tmpReq *FlightModifyOtaSearchV2Request, headers *FlightModifyOtaSearchV2Headers, runtime *dara.RuntimeOptions) (_result *FlightModifyOtaSearchV2Response, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FlightModifyOtaSearchV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CabinClass) {
		request.CabinClassShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CabinClass, dara.String("cabin_class"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DepDate) {
		request.DepDateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepDate, dara.String("dep_date"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PassengerSegmentRelations) {
		request.PassengerSegmentRelationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PassengerSegmentRelations, dara.String("passenger_segment_relations"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SelectedSegments) {
		request.SelectedSegmentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SelectedSegments, dara.String("selected_segments"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CabinClassShrink) {
		query["cabin_class"] = request.CabinClassShrink
	}

	if !dara.IsNil(request.DepDateShrink) {
		query["dep_date"] = request.DepDateShrink
	}

	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.PassengerSegmentRelationsShrink) {
		query["passenger_segment_relations"] = request.PassengerSegmentRelationsShrink
	}

	if !dara.IsNil(request.SelectedSegmentsShrink) {
		query["selected_segments"] = request.SelectedSegmentsShrink
	}

	if !dara.IsNil(request.SessionId) {
		query["session_id"] = request.SessionId
	}

	if !dara.IsNil(request.Voluntary) {
		query["voluntary"] = request.Voluntary
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightModifyOtaSearchV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/modify/action/ota-search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightModifyOtaSearchV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票改签支付
//
// @param tmpReq - FlightModifyPayV2Request
//
// @param headers - FlightModifyPayV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightModifyPayV2Response
func (client *Client) FlightModifyPayV2WithContext(ctx context.Context, tmpReq *FlightModifyPayV2Request, headers *FlightModifyPayV2Headers, runtime *dara.RuntimeOptions) (_result *FlightModifyPayV2Response, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FlightModifyPayV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExtParams) {
		request.ExtParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtParams, dara.String("ext_params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtParamsShrink) {
		body["ext_params"] = request.ExtParamsShrink
	}

	if !dara.IsNil(request.IsvName) {
		body["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.ModifyPayAmount) {
		body["modify_pay_amount"] = request.ModifyPayAmount
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutSubOrderId) {
		body["out_sub_order_id"] = request.OutSubOrderId
	}

	if !dara.IsNil(request.SubOrderId) {
		body["sub_order_id"] = request.SubOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightModifyPayV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/modify/action/pay"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightModifyPayV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航班订单明细信息
//
// @param request - FlightOrderDetailInfoRequest
//
// @param headers - FlightOrderDetailInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightOrderDetailInfoResponse
func (client *Client) FlightOrderDetailInfoWithContext(ctx context.Context, request *FlightOrderDetailInfoRequest, headers *FlightOrderDetailInfoHeaders, runtime *dara.RuntimeOptions) (_result *FlightOrderDetailInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisOrderId) {
		query["dis_order_id"] = request.DisOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightOrderDetailInfo"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/order/action/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightOrderDetailInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票订单详情
//
// @param request - FlightOrderDetailV2Request
//
// @param headers - FlightOrderDetailV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightOrderDetailV2Response
func (client *Client) FlightOrderDetailV2WithContext(ctx context.Context, request *FlightOrderDetailV2Request, headers *FlightOrderDetailV2Headers, runtime *dara.RuntimeOptions) (_result *FlightOrderDetailV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["out_order_id"] = request.OutOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightOrderDetailV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/order/action/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightOrderDetailV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询机票订单列表
//
// @param request - FlightOrderListQueryRequest
//
// @param headers - FlightOrderListQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightOrderListQueryResponse
func (client *Client) FlightOrderListQueryWithContext(ctx context.Context, request *FlightOrderListQueryRequest, headers *FlightOrderListQueryHeaders, runtime *dara.RuntimeOptions) (_result *FlightOrderListQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllApply) {
		query["all_apply"] = request.AllApply
	}

	if !dara.IsNil(request.ApplyId) {
		query["apply_id"] = request.ApplyId
	}

	if !dara.IsNil(request.DepartId) {
		query["depart_id"] = request.DepartId
	}

	if !dara.IsNil(request.EndTime) {
		query["end_time"] = request.EndTime
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["start_time"] = request.StartTime
	}

	if !dara.IsNil(request.ThirdpartApplyId) {
		query["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !dara.IsNil(request.UpdateEndTime) {
		query["update_end_time"] = request.UpdateEndTime
	}

	if !dara.IsNil(request.UpdateStartTime) {
		query["update_start_time"] = request.UpdateStartTime
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightOrderListQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/flight/v1/order-list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightOrderListQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票订单列表查询
//
// @param tmpReq - FlightOrderListQueryV2Request
//
// @param headers - FlightOrderListQueryV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightOrderListQueryV2Response
func (client *Client) FlightOrderListQueryV2WithContext(ctx context.Context, tmpReq *FlightOrderListQueryV2Request, headers *FlightOrderListQueryV2Headers, runtime *dara.RuntimeOptions) (_result *FlightOrderListQueryV2Response, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FlightOrderListQueryV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApproveId) {
		request.ApproveIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApproveId, dara.String("approve_id"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.BookerId) {
		request.BookerIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BookerId, dara.String("booker_id"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DepartId) {
		request.DepartIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepartId, dara.String("depart_id"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Supplier) {
		request.SupplierShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Supplier, dara.String("supplier"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ThirdpartApproveId) {
		request.ThirdpartApproveIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ThirdpartApproveId, dara.String("thirdpart_approve_id"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApproveIdShrink) {
		query["approve_id"] = request.ApproveIdShrink
	}

	if !dara.IsNil(request.BookerIdShrink) {
		query["booker_id"] = request.BookerIdShrink
	}

	if !dara.IsNil(request.DepartIdShrink) {
		query["depart_id"] = request.DepartIdShrink
	}

	if !dara.IsNil(request.EndDate) {
		query["end_date"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		query["page_Size"] = request.PageSize
	}

	if !dara.IsNil(request.ScrollId) {
		query["scroll_id"] = request.ScrollId
	}

	if !dara.IsNil(request.StartDate) {
		query["start_date"] = request.StartDate
	}

	if !dara.IsNil(request.SupplierShrink) {
		query["supplier"] = request.SupplierShrink
	}

	if !dara.IsNil(request.ThirdpartApproveIdShrink) {
		query["thirdpart_approve_id"] = request.ThirdpartApproveIdShrink
	}

	if !dara.IsNil(request.UpdateEndDate) {
		query["update_end_date"] = request.UpdateEndDate
	}

	if !dara.IsNil(request.UpdateStartDate) {
		query["update_start_date"] = request.UpdateStartDate
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightOrderListQueryV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/v2/Flight-order-list-query"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightOrderListQueryV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询机票订单详情（含票信息）
//
// @param request - FlightOrderQueryRequest
//
// @param headers - FlightOrderQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightOrderQueryResponse
func (client *Client) FlightOrderQueryWithContext(ctx context.Context, request *FlightOrderQueryRequest, headers *FlightOrderQueryHeaders, runtime *dara.RuntimeOptions) (_result *FlightOrderQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightOrderQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/flight/v1/order"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightOrderQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询退改规则行李额
//
// @param request - FlightOtaItemDetailRequest
//
// @param headers - FlightOtaItemDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightOtaItemDetailResponse
func (client *Client) FlightOtaItemDetailWithContext(ctx context.Context, request *FlightOtaItemDetailRequest, headers *FlightOtaItemDetailHeaders, runtime *dara.RuntimeOptions) (_result *FlightOtaItemDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OtaItemId) {
		query["ota_item_id"] = request.OtaItemId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightOtaItemDetail"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/flight/action/ota-item-detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightOtaItemDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航班最低价搜索
//
// @param request - FlightOtaSearchRequest
//
// @param headers - FlightOtaSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightOtaSearchResponse
func (client *Client) FlightOtaSearchWithContext(ctx context.Context, request *FlightOtaSearchRequest, headers *FlightOtaSearchHeaders, runtime *dara.RuntimeOptions) (_result *FlightOtaSearchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AirlineCode) {
		query["airline_code"] = request.AirlineCode
	}

	if !dara.IsNil(request.ArrCityCode) {
		query["arr_city_code"] = request.ArrCityCode
	}

	if !dara.IsNil(request.CabinClass) {
		query["cabin_class"] = request.CabinClass
	}

	if !dara.IsNil(request.CarrierFlightNo) {
		query["carrier_flight_no"] = request.CarrierFlightNo
	}

	if !dara.IsNil(request.DepCityCode) {
		query["dep_city_code"] = request.DepCityCode
	}

	if !dara.IsNil(request.DepDate) {
		query["dep_date"] = request.DepDate
	}

	if !dara.IsNil(request.FlightNo) {
		query["flight_no"] = request.FlightNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightOtaSearch"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/flight/action/ota-search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightOtaSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 单航班报价搜索
//
// @param tmpReq - FlightOtaSearchV2Request
//
// @param headers - FlightOtaSearchV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightOtaSearchV2Response
func (client *Client) FlightOtaSearchV2WithContext(ctx context.Context, tmpReq *FlightOtaSearchV2Request, headers *FlightOtaSearchV2Headers, runtime *dara.RuntimeOptions) (_result *FlightOtaSearchV2Response, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FlightOtaSearchV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CabinTypeList) {
		request.CabinTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CabinTypeList, dara.String("cabin_type_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SearchJourneys) {
		request.SearchJourneysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchJourneys, dara.String("search_journeys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CabinTypeListShrink) {
		query["cabin_type_list"] = request.CabinTypeListShrink
	}

	if !dara.IsNil(request.DirectOnly) {
		query["direct_only"] = request.DirectOnly
	}

	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.NeedShareFlight) {
		query["need_share_flight"] = request.NeedShareFlight
	}

	if !dara.IsNil(request.SearchJourneysShrink) {
		query["search_journeys"] = request.SearchJourneysShrink
	}

	if !dara.IsNil(request.SearchMode) {
		query["search_mode"] = request.SearchMode
	}

	if !dara.IsNil(request.TripType) {
		query["trip_type"] = request.TripType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightOtaSearchV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/flight/action/ota-search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightOtaSearchV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航班订单支付
//
// @param tmpReq - FlightPayOrderRequest
//
// @param headers - FlightPayOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightPayOrderResponse
func (client *Client) FlightPayOrderWithContext(ctx context.Context, tmpReq *FlightPayOrderRequest, headers *FlightPayOrderHeaders, runtime *dara.RuntimeOptions) (_result *FlightPayOrderResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FlightPayOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("extra"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CorpPayPrice) {
		body["corp_pay_price"] = request.CorpPayPrice
	}

	if !dara.IsNil(request.DisOrderId) {
		body["dis_order_id"] = request.DisOrderId
	}

	if !dara.IsNil(request.ExtraShrink) {
		body["extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.PersonalPayPrice) {
		body["personal_pay_price"] = request.PersonalPayPrice
	}

	if !dara.IsNil(request.TotalPayPrice) {
		body["total_pay_price"] = request.TotalPayPrice
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightPayOrder"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/order/action/pay"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightPayOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票订单支付
//
// @param request - FlightPayOrderV2Request
//
// @param headers - FlightPayOrderV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightPayOrderV2Response
func (client *Client) FlightPayOrderV2WithContext(ctx context.Context, request *FlightPayOrderV2Request, headers *FlightPayOrderV2Headers, runtime *dara.RuntimeOptions) (_result *FlightPayOrderV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IsvName) {
		body["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.TotalPrice) {
		body["total_price"] = request.TotalPrice
	}

	if !dara.IsNil(request.TotalServiceFeePrice) {
		body["total_service_fee_price"] = request.TotalServiceFeePrice
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightPayOrderV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/order/action/pay"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightPayOrderV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航班退票申请
//
// @param tmpReq - FlightRefundApplyRequest
//
// @param headers - FlightRefundApplyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightRefundApplyResponse
func (client *Client) FlightRefundApplyWithContext(ctx context.Context, tmpReq *FlightRefundApplyRequest, headers *FlightRefundApplyHeaders, runtime *dara.RuntimeOptions) (_result *FlightRefundApplyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FlightRefundApplyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("extra"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PassengerSegmentInfoList) {
		request.PassengerSegmentInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PassengerSegmentInfoList, dara.String("passenger_segment_info_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RefundVoucherInfo) {
		request.RefundVoucherInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RefundVoucherInfo, dara.String("refund_voucher_info"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CorpRefundPrice) {
		body["corp_refund_price"] = request.CorpRefundPrice
	}

	if !dara.IsNil(request.DisOrderId) {
		body["dis_order_id"] = request.DisOrderId
	}

	if !dara.IsNil(request.DisSubOrderId) {
		body["dis_sub_order_id"] = request.DisSubOrderId
	}

	if !dara.IsNil(request.DisplayRefundMoney) {
		body["display_refund_money"] = request.DisplayRefundMoney
	}

	if !dara.IsNil(request.ExtraShrink) {
		body["extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.IsVoluntary) {
		body["is_voluntary"] = request.IsVoluntary
	}

	if !dara.IsNil(request.ItemUnitIds) {
		body["item_unit_ids"] = request.ItemUnitIds
	}

	if !dara.IsNil(request.PassengerSegmentInfoListShrink) {
		body["passenger_segment_info_list"] = request.PassengerSegmentInfoListShrink
	}

	if !dara.IsNil(request.PersonalRefundPrice) {
		body["personal_refund_price"] = request.PersonalRefundPrice
	}

	if !dara.IsNil(request.ReasonDetail) {
		body["reason_detail"] = request.ReasonDetail
	}

	if !dara.IsNil(request.ReasonType) {
		body["reason_type"] = request.ReasonType
	}

	if !dara.IsNil(request.RefundVoucherInfoShrink) {
		body["refund_voucher_info"] = request.RefundVoucherInfoShrink
	}

	if !dara.IsNil(request.SessionId) {
		body["session_id"] = request.SessionId
	}

	if !dara.IsNil(request.TotalRefundPrice) {
		body["total_refund_price"] = request.TotalRefundPrice
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightRefundApply"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/refund/action/apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightRefundApplyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票退票申请
//
// @param tmpReq - FlightRefundApplyV2Request
//
// @param headers - FlightRefundApplyV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightRefundApplyV2Response
func (client *Client) FlightRefundApplyV2WithContext(ctx context.Context, tmpReq *FlightRefundApplyV2Request, headers *FlightRefundApplyV2Headers, runtime *dara.RuntimeOptions) (_result *FlightRefundApplyV2Response, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FlightRefundApplyV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PassengerSegmentRelations) {
		request.PassengerSegmentRelationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PassengerSegmentRelations, dara.String("passenger_segment_relations"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TicketNos) {
		request.TicketNosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TicketNos, dara.String("ticket_nos"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IsvName) {
		body["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutSubOrderId) {
		body["out_sub_order_id"] = request.OutSubOrderId
	}

	if !dara.IsNil(request.PassengerSegmentRelationsShrink) {
		body["passenger_segment_relations"] = request.PassengerSegmentRelationsShrink
	}

	if !dara.IsNil(request.PreCalType) {
		body["pre_cal_type"] = request.PreCalType
	}

	if !dara.IsNil(request.RefundReason) {
		body["refund_reason"] = request.RefundReason
	}

	if !dara.IsNil(request.RefundReasonType) {
		body["refund_reason_type"] = request.RefundReasonType
	}

	if !dara.IsNil(request.TicketNosShrink) {
		body["ticket_nos"] = request.TicketNosShrink
	}

	if !dara.IsNil(request.TotalRefundPrice) {
		body["total_refund_price"] = request.TotalRefundPrice
	}

	if !dara.IsNil(request.UploadPictUrls) {
		body["upload_pict_urls"] = request.UploadPictUrls
	}

	if !dara.IsNil(request.Voluntary) {
		body["voluntary"] = request.Voluntary
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightRefundApplyV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/refund/action/apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightRefundApplyV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航班退票详情
//
// @param request - FlightRefundDetailRequest
//
// @param headers - FlightRefundDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightRefundDetailResponse
func (client *Client) FlightRefundDetailWithContext(ctx context.Context, request *FlightRefundDetailRequest, headers *FlightRefundDetailHeaders, runtime *dara.RuntimeOptions) (_result *FlightRefundDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisOrderId) {
		query["dis_order_id"] = request.DisOrderId
	}

	if !dara.IsNil(request.DisSubOrderId) {
		query["dis_sub_order_id"] = request.DisSubOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightRefundDetail"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/refund/action/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightRefundDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票退票详情
//
// @param request - FlightRefundDetailV2Request
//
// @param headers - FlightRefundDetailV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightRefundDetailV2Response
func (client *Client) FlightRefundDetailV2WithContext(ctx context.Context, request *FlightRefundDetailV2Request, headers *FlightRefundDetailV2Headers, runtime *dara.RuntimeOptions) (_result *FlightRefundDetailV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutRefundApplyId) {
		query["out_refund_apply_id"] = request.OutRefundApplyId
	}

	if !dara.IsNil(request.RefundApplyId) {
		query["refund_apply_id"] = request.RefundApplyId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightRefundDetailV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/refund/action/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightRefundDetailV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票退票预计算
//
// @param tmpReq - FlightRefundPreCalRequest
//
// @param headers - FlightRefundPreCalHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightRefundPreCalResponse
func (client *Client) FlightRefundPreCalWithContext(ctx context.Context, tmpReq *FlightRefundPreCalRequest, headers *FlightRefundPreCalHeaders, runtime *dara.RuntimeOptions) (_result *FlightRefundPreCalResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FlightRefundPreCalShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PassengerSegmentInfoList) {
		request.PassengerSegmentInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PassengerSegmentInfoList, dara.String("passenger_segment_info_list"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DisOrderId) {
		query["dis_order_id"] = request.DisOrderId
	}

	if !dara.IsNil(request.IsVoluntary) {
		query["is_voluntary"] = request.IsVoluntary
	}

	if !dara.IsNil(request.PassengerSegmentInfoListShrink) {
		query["passenger_segment_info_list"] = request.PassengerSegmentInfoListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightRefundPreCal"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/refund/action/pre-cal"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightRefundPreCalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票退票费用预计算
//
// @param tmpReq - FlightRefundPreCalV2Request
//
// @param headers - FlightRefundPreCalV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightRefundPreCalV2Response
func (client *Client) FlightRefundPreCalV2WithContext(ctx context.Context, tmpReq *FlightRefundPreCalV2Request, headers *FlightRefundPreCalV2Headers, runtime *dara.RuntimeOptions) (_result *FlightRefundPreCalV2Response, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FlightRefundPreCalV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PassengerSegmentRelations) {
		request.PassengerSegmentRelationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PassengerSegmentRelations, dara.String("passenger_segment_relations"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TicketNos) {
		request.TicketNosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TicketNos, dara.String("ticket_nos"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.PassengerSegmentRelationsShrink) {
		query["passenger_segment_relations"] = request.PassengerSegmentRelationsShrink
	}

	if !dara.IsNil(request.PreCalType) {
		query["pre_cal_type"] = request.PreCalType
	}

	if !dara.IsNil(request.TicketNosShrink) {
		query["ticket_nos"] = request.TicketNosShrink
	}

	if !dara.IsNil(request.Voluntary) {
		query["voluntary"] = request.Voluntary
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightRefundPreCalV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v2/refund/action/pre-cal"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightRefundPreCalV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航班列表搜索
//
// @param request - FlightSearchListRequest
//
// @param headers - FlightSearchListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightSearchListResponse
func (client *Client) FlightSearchListWithContext(ctx context.Context, request *FlightSearchListRequest, headers *FlightSearchListHeaders, runtime *dara.RuntimeOptions) (_result *FlightSearchListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AirlineCode) {
		query["airline_code"] = request.AirlineCode
	}

	if !dara.IsNil(request.ArrCityCode) {
		query["arr_city_code"] = request.ArrCityCode
	}

	if !dara.IsNil(request.ArrCityName) {
		query["arr_city_name"] = request.ArrCityName
	}

	if !dara.IsNil(request.ArrDate) {
		query["arr_date"] = request.ArrDate
	}

	if !dara.IsNil(request.CabinClass) {
		query["cabin_class"] = request.CabinClass
	}

	if !dara.IsNil(request.DepCityCode) {
		query["dep_city_code"] = request.DepCityCode
	}

	if !dara.IsNil(request.DepCityName) {
		query["dep_city_name"] = request.DepCityName
	}

	if !dara.IsNil(request.DepDate) {
		query["dep_date"] = request.DepDate
	}

	if !dara.IsNil(request.FlightNo) {
		query["flight_no"] = request.FlightNo
	}

	if !dara.IsNil(request.NeedMultiClassPrice) {
		query["need_multi_class_price"] = request.NeedMultiClassPrice
	}

	if !dara.IsNil(request.TransferCityCode) {
		query["transfer_city_code"] = request.TransferCityCode
	}

	if !dara.IsNil(request.TransferFlightNo) {
		query["transfer_flight_no"] = request.TransferFlightNo
	}

	if !dara.IsNil(request.TransferLeaveDate) {
		query["transfer_leave_date"] = request.TransferLeaveDate
	}

	if !dara.IsNil(request.TripType) {
		query["trip_type"] = request.TripType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightSearchList"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/huge/dtb-flight/v1/flight/action/search-list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightSearchListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询福豆记账数据
//
// @param request - FuPointBillSettlementQueryRequest
//
// @param headers - FuPointBillSettlementQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FuPointBillSettlementQueryResponse
func (client *Client) FuPointBillSettlementQueryWithContext(ctx context.Context, request *FuPointBillSettlementQueryRequest, headers *FuPointBillSettlementQueryHeaders, runtime *dara.RuntimeOptions) (_result *FuPointBillSettlementQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillBatch) {
		query["bill_batch"] = request.BillBatch
	}

	if !dara.IsNil(request.CooperatorId) {
		query["cooperator_id"] = request.CooperatorId
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PeriodEnd) {
		query["period_end"] = request.PeriodEnd
	}

	if !dara.IsNil(request.PeriodStart) {
		query["period_start"] = request.PeriodStart
	}

	if !dara.IsNil(request.ScrollId) {
		query["scroll_id"] = request.ScrollId
	}

	if !dara.IsNil(request.ScrollMod) {
		query["scroll_mod"] = request.ScrollMod
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FuPointBillSettlementQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/fupoint/v1/bill-settlement"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FuPointBillSettlementQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 换取GroupCorpToken接口
//
// @param request - GroupCorpTokenRequest
//
// @param headers - GroupCorpTokenHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GroupCorpTokenResponse
func (client *Client) GroupCorpTokenWithContext(ctx context.Context, request *GroupCorpTokenRequest, headers *GroupCorpTokenHeaders, runtime *dara.RuntimeOptions) (_result *GroupCorpTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppSecret) {
		query["app_secret"] = request.AppSecret
	}

	if !dara.IsNil(request.CorpId) {
		query["corp_id"] = request.CorpId
	}

	if !dara.IsNil(request.SubCorpId) {
		query["sub_corp_id"] = request.SubCorpId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripAccessToken) {
		realHeaders["x-acs-btrip-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripAccessToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GroupCorpToken"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/btrip-open-auth/v1/group-corp-token/action/take"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GroupCorpTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 集团部门同步
//
// @param tmpReq - GroupDepartSaveRequest
//
// @param headers - GroupDepartSaveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GroupDepartSaveResponse
func (client *Client) GroupDepartSaveWithContext(ctx context.Context, tmpReq *GroupDepartSaveRequest, headers *GroupDepartSaveHeaders, runtime *dara.RuntimeOptions) (_result *GroupDepartSaveResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GroupDepartSaveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SubCorpIdList) {
		request.SubCorpIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubCorpIdList, dara.String("sub_corp_id_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeptName) {
		body["dept_name"] = request.DeptName
	}

	if !dara.IsNil(request.ManagerIds) {
		body["manager_ids"] = request.ManagerIds
	}

	if !dara.IsNil(request.OuterDeptId) {
		body["outer_dept_id"] = request.OuterDeptId
	}

	if !dara.IsNil(request.OuterDeptPid) {
		body["outer_dept_pid"] = request.OuterDeptPid
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.SubCorpIdListShrink) {
		body["sub_corp_id_list"] = request.SubCorpIdListShrink
	}

	if !dara.IsNil(request.SyncGroup) {
		body["sync_group"] = request.SyncGroup
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GroupDepartSave"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/sub_corps/v1/departs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GroupDepartSaveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 集团人员同步
//
// @param tmpReq - GroupUserSaveRequest
//
// @param headers - GroupUserSaveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GroupUserSaveResponse
func (client *Client) GroupUserSaveWithContext(ctx context.Context, tmpReq *GroupUserSaveRequest, headers *GroupUserSaveHeaders, runtime *dara.RuntimeOptions) (_result *GroupUserSaveResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GroupUserSaveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CertList) {
		request.CertListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CertList, dara.String("cert_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SubCorpIdList) {
		request.SubCorpIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubCorpIdList, dara.String("sub_corp_id_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseCityCode) {
		body["base_city_code"] = request.BaseCityCode
	}

	if !dara.IsNil(request.Birthday) {
		body["birthday"] = request.Birthday
	}

	if !dara.IsNil(request.CertListShrink) {
		body["cert_list"] = request.CertListShrink
	}

	if !dara.IsNil(request.Gender) {
		body["gender"] = request.Gender
	}

	if !dara.IsNil(request.JobNo) {
		body["job_no"] = request.JobNo
	}

	if !dara.IsNil(request.Phone) {
		body["phone"] = request.Phone
	}

	if !dara.IsNil(request.RealNameEn) {
		body["real_name_en"] = request.RealNameEn
	}

	if !dara.IsNil(request.SubCorpIdListShrink) {
		body["sub_corp_id_list"] = request.SubCorpIdListShrink
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		body["user_name"] = request.UserName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GroupUserSave"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/sub_corps/v1/users"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GroupUserSaveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店起价
//
// @param tmpReq - HotelAskingPriceRequest
//
// @param headers - HotelAskingPriceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelAskingPriceResponse
func (client *Client) HotelAskingPriceWithContext(ctx context.Context, tmpReq *HotelAskingPriceRequest, headers *HotelAskingPriceHeaders, runtime *dara.RuntimeOptions) (_result *HotelAskingPriceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &HotelAskingPriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Shids) {
		request.ShidsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Shids, dara.String("shids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdultNum) {
		query["adult_num"] = request.AdultNum
	}

	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.CheckInDate) {
		query["check_in_date"] = request.CheckInDate
	}

	if !dara.IsNil(request.CheckOutDate) {
		query["check_out_date"] = request.CheckOutDate
	}

	if !dara.IsNil(request.CityCode) {
		query["city_code"] = request.CityCode
	}

	if !dara.IsNil(request.CityName) {
		query["city_name"] = request.CityName
	}

	if !dara.IsNil(request.Dir) {
		query["dir"] = request.Dir
	}

	if !dara.IsNil(request.HotelStar) {
		query["hotel_star"] = request.HotelStar
	}

	if !dara.IsNil(request.IsProtocol) {
		query["is_protocol"] = request.IsProtocol
	}

	if !dara.IsNil(request.PaymentType) {
		query["payment_type"] = request.PaymentType
	}

	if !dara.IsNil(request.ShidsShrink) {
		query["shids"] = request.ShidsShrink
	}

	if !dara.IsNil(request.SortCode) {
		query["sort_code"] = request.SortCode
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelAskingPrice"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/hotels/action/asking-price"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelAskingPriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询酒店记账数据
//
// @param request - HotelBillSettlementQueryRequest
//
// @param headers - HotelBillSettlementQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelBillSettlementQueryResponse
func (client *Client) HotelBillSettlementQueryWithContext(ctx context.Context, request *HotelBillSettlementQueryRequest, headers *HotelBillSettlementQueryHeaders, runtime *dara.RuntimeOptions) (_result *HotelBillSettlementQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillBatch) {
		query["bill_batch"] = request.BillBatch
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PeriodEnd) {
		query["period_end"] = request.PeriodEnd
	}

	if !dara.IsNil(request.PeriodStart) {
		query["period_start"] = request.PeriodStart
	}

	if !dara.IsNil(request.ScrollId) {
		query["scroll_id"] = request.ScrollId
	}

	if !dara.IsNil(request.ScrollMod) {
		query["scroll_mod"] = request.ScrollMod
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelBillSettlementQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/hotel/v1/bill-settlement"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelBillSettlementQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店城市列表
//
// @param request - HotelCityCodeListRequest
//
// @param headers - HotelCityCodeListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelCityCodeListResponse
func (client *Client) HotelCityCodeListWithContext(ctx context.Context, request *HotelCityCodeListRequest, headers *HotelCityCodeListHeaders, runtime *dara.RuntimeOptions) (_result *HotelCityCodeListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CountryCode) {
		query["country_code"] = request.CountryCode
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelCityCodeList"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/city-codes/action/search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelCityCodeListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询酒店超标审批详情
//
// @param request - HotelExceedApplyQueryRequest
//
// @param headers - HotelExceedApplyQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelExceedApplyQueryResponse
func (client *Client) HotelExceedApplyQueryWithContext(ctx context.Context, request *HotelExceedApplyQueryRequest, headers *HotelExceedApplyQueryHeaders, runtime *dara.RuntimeOptions) (_result *HotelExceedApplyQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyId) {
		query["apply_id"] = request.ApplyId
	}

	if !dara.IsNil(request.BusinessInstanceId) {
		query["business_instance_id"] = request.BusinessInstanceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelExceedApplyQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/hotel-exceed"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelExceedApplyQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店详情页报价接口(直连)
//
// @param request - HotelGoodsQueryRequest
//
// @param headers - HotelGoodsQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelGoodsQueryResponse
func (client *Client) HotelGoodsQueryWithContext(ctx context.Context, request *HotelGoodsQueryRequest, headers *HotelGoodsQueryHeaders, runtime *dara.RuntimeOptions) (_result *HotelGoodsQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdultNum) {
		query["adult_num"] = request.AdultNum
	}

	if !dara.IsNil(request.AgreementPrice) {
		query["agreement_price"] = request.AgreementPrice
	}

	if !dara.IsNil(request.BeginDate) {
		query["begin_date"] = request.BeginDate
	}

	if !dara.IsNil(request.BreakfastIncluded) {
		query["breakfast_included"] = request.BreakfastIncluded
	}

	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.CityCode) {
		query["city_code"] = request.CityCode
	}

	if !dara.IsNil(request.EndDate) {
		query["end_date"] = request.EndDate
	}

	if !dara.IsNil(request.HotelId) {
		query["hotel_id"] = request.HotelId
	}

	if !dara.IsNil(request.PayOverType) {
		query["pay_over_type"] = request.PayOverType
	}

	if !dara.IsNil(request.PaymentType) {
		query["payment_type"] = request.PaymentType
	}

	if !dara.IsNil(request.SpecialInvoice) {
		query["special_invoice"] = request.SpecialInvoice
	}

	if !dara.IsNil(request.SuperMan) {
		query["super_man"] = request.SuperMan
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelGoodsQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/hotel-goods"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelGoodsQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店清单
//
// @param request - HotelIndexInfoRequest
//
// @param headers - HotelIndexInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelIndexInfoResponse
func (client *Client) HotelIndexInfoWithContext(ctx context.Context, request *HotelIndexInfoRequest, headers *HotelIndexInfoHeaders, runtime *dara.RuntimeOptions) (_result *HotelIndexInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CityCode) {
		query["city_code"] = request.CityCode
	}

	if !dara.IsNil(request.HotelStatus) {
		query["hotel_status"] = request.HotelStatus
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PageToken) {
		query["page_token"] = request.PageToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelIndexInfo"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/index-infos"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelIndexInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店订单取消
//
// @param request - HotelOrderCancelRequest
//
// @param headers - HotelOrderCancelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelOrderCancelResponse
func (client *Client) HotelOrderCancelWithContext(ctx context.Context, request *HotelOrderCancelRequest, headers *HotelOrderCancelHeaders, runtime *dara.RuntimeOptions) (_result *HotelOrderCancelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BtripOrderId) {
		query["btrip_order_id"] = request.BtripOrderId
	}

	if !dara.IsNil(request.DisOrderId) {
		query["dis_order_id"] = request.DisOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelOrderCancel"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/orders/action/cancel"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelOrderCancelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店订单修改申请
//
// @param tmpReq - HotelOrderChangeApplyRequest
//
// @param headers - HotelOrderChangeApplyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelOrderChangeApplyResponse
func (client *Client) HotelOrderChangeApplyWithContext(ctx context.Context, tmpReq *HotelOrderChangeApplyRequest, headers *HotelOrderChangeApplyHeaders, runtime *dara.RuntimeOptions) (_result *HotelOrderChangeApplyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &HotelOrderChangeApplyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoomInfoList) {
		request.RoomInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoomInfoList, dara.String("room_info_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		body["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.DisOrderId) {
		body["dis_order_id"] = request.DisOrderId
	}

	if !dara.IsNil(request.Reason) {
		body["reason"] = request.Reason
	}

	if !dara.IsNil(request.RoomInfoListShrink) {
		body["room_info_list"] = request.RoomInfoListShrink
	}

	if !dara.IsNil(request.SaleOrderId) {
		body["sale_order_id"] = request.SaleOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelOrderChangeApply"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/orders/action/change/apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelOrderChangeApplyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店订单修改详情
//
// @param request - HotelOrderChangeDetailRequest
//
// @param headers - HotelOrderChangeDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelOrderChangeDetailResponse
func (client *Client) HotelOrderChangeDetailWithContext(ctx context.Context, request *HotelOrderChangeDetailRequest, headers *HotelOrderChangeDetailHeaders, runtime *dara.RuntimeOptions) (_result *HotelOrderChangeDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		body["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.ChangeOrderId) {
		body["change_order_id"] = request.ChangeOrderId
	}

	if !dara.IsNil(request.DisOrderId) {
		body["dis_order_id"] = request.DisOrderId
	}

	if !dara.IsNil(request.SaleOrderId) {
		body["sale_order_id"] = request.SaleOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelOrderChangeDetail"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/orders/action/change/detail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelOrderChangeDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店订单创建
//
// @param tmpReq - HotelOrderCreateRequest
//
// @param headers - HotelOrderCreateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelOrderCreateResponse
func (client *Client) HotelOrderCreateWithContext(ctx context.Context, tmpReq *HotelOrderCreateRequest, headers *HotelOrderCreateHeaders, runtime *dara.RuntimeOptions) (_result *HotelOrderCreateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &HotelOrderCreateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InvoiceInfo) {
		request.InvoiceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InvoiceInfo, dara.String("invoice_info"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OccupantInfoList) {
		request.OccupantInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OccupantInfoList, dara.String("occupant_info_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PromotionInfo) {
		request.PromotionInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PromotionInfo, dara.String("promotion_info"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		body["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.CheckIn) {
		body["check_in"] = request.CheckIn
	}

	if !dara.IsNil(request.CheckOut) {
		body["check_out"] = request.CheckOut
	}

	if !dara.IsNil(request.ContractEmail) {
		body["contract_email"] = request.ContractEmail
	}

	if !dara.IsNil(request.ContractName) {
		body["contract_name"] = request.ContractName
	}

	if !dara.IsNil(request.ContractPhone) {
		body["contract_phone"] = request.ContractPhone
	}

	if !dara.IsNil(request.CorpPayPrice) {
		body["corp_pay_price"] = request.CorpPayPrice
	}

	if !dara.IsNil(request.DisOrderId) {
		body["dis_order_id"] = request.DisOrderId
	}

	if !dara.IsNil(request.Extra) {
		body["extra"] = request.Extra
	}

	if !dara.IsNil(request.InvoiceInfoShrink) {
		body["invoice_info"] = request.InvoiceInfoShrink
	}

	if !dara.IsNil(request.ItemId) {
		body["item_id"] = request.ItemId
	}

	if !dara.IsNil(request.ItineraryNo) {
		body["itinerary_no"] = request.ItineraryNo
	}

	if !dara.IsNil(request.OccupantInfoListShrink) {
		body["occupant_info_list"] = request.OccupantInfoListShrink
	}

	if !dara.IsNil(request.PersonPayPrice) {
		body["person_pay_price"] = request.PersonPayPrice
	}

	if !dara.IsNil(request.PromotionInfoShrink) {
		body["promotion_info"] = request.PromotionInfoShrink
	}

	if !dara.IsNil(request.RatePlanId) {
		body["rate_plan_id"] = request.RatePlanId
	}

	if !dara.IsNil(request.RoomId) {
		body["room_id"] = request.RoomId
	}

	if !dara.IsNil(request.RoomNum) {
		body["room_num"] = request.RoomNum
	}

	if !dara.IsNil(request.SellerId) {
		body["seller_id"] = request.SellerId
	}

	if !dara.IsNil(request.Shid) {
		body["shid"] = request.Shid
	}

	if !dara.IsNil(request.TotalOrderPrice) {
		body["total_order_price"] = request.TotalOrderPrice
	}

	if !dara.IsNil(request.ValidateResKey) {
		body["validate_res_key"] = request.ValidateResKey
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelOrderCreate"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/orders/action/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelOrderCreateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店订单明细信息
//
// @param request - HotelOrderDetailInfoRequest
//
// @param headers - HotelOrderDetailInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelOrderDetailInfoResponse
func (client *Client) HotelOrderDetailInfoWithContext(ctx context.Context, request *HotelOrderDetailInfoRequest, headers *HotelOrderDetailInfoHeaders, runtime *dara.RuntimeOptions) (_result *HotelOrderDetailInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BtripOrderId) {
		query["btrip_order_id"] = request.BtripOrderId
	}

	if !dara.IsNil(request.DisOrderId) {
		query["dis_order_id"] = request.DisOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelOrderDetailInfo"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/orders/action/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelOrderDetailInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自营酒店订单查询
//
// @param headers - HotelOrderInfoQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelOrderInfoQueryResponse
func (client *Client) HotelOrderInfoQueryWithContext(ctx context.Context, orderId *string, headers *HotelOrderInfoQueryHeaders, runtime *dara.RuntimeOptions) (_result *HotelOrderInfoQueryResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelOrderInfoQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/order/v1/hotelOrders/" + dara.PercentEncode(dara.StringValue(orderId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelOrderInfoQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询酒店订单列表
//
// @param request - HotelOrderListQueryRequest
//
// @param headers - HotelOrderListQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelOrderListQueryResponse
func (client *Client) HotelOrderListQueryWithContext(ctx context.Context, request *HotelOrderListQueryRequest, headers *HotelOrderListQueryHeaders, runtime *dara.RuntimeOptions) (_result *HotelOrderListQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllApply) {
		query["all_apply"] = request.AllApply
	}

	if !dara.IsNil(request.ApplyId) {
		query["apply_id"] = request.ApplyId
	}

	if !dara.IsNil(request.Category) {
		query["category"] = request.Category
	}

	if !dara.IsNil(request.DepartId) {
		query["depart_id"] = request.DepartId
	}

	if !dara.IsNil(request.EndTime) {
		query["end_time"] = request.EndTime
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["start_time"] = request.StartTime
	}

	if !dara.IsNil(request.ThirdpartApplyId) {
		query["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !dara.IsNil(request.UpdateEndTime) {
		query["update_end_time"] = request.UpdateEndTime
	}

	if !dara.IsNil(request.UpdateStartTime) {
		query["update_start_time"] = request.UpdateStartTime
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelOrderListQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/hotel/v1/order-list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelOrderListQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店订单支付
//
// @param request - HotelOrderPayRequest
//
// @param headers - HotelOrderPayHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelOrderPayResponse
func (client *Client) HotelOrderPayWithContext(ctx context.Context, request *HotelOrderPayRequest, headers *HotelOrderPayHeaders, runtime *dara.RuntimeOptions) (_result *HotelOrderPayResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BtripOrderId) {
		body["btrip_order_id"] = request.BtripOrderId
	}

	if !dara.IsNil(request.BtripUserId) {
		body["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.CompanyPayFee) {
		body["company_pay_fee"] = request.CompanyPayFee
	}

	if !dara.IsNil(request.PersonPayFee) {
		body["person_pay_fee"] = request.PersonPayFee
	}

	if !dara.IsNil(request.ThirdPayAccount) {
		body["third_pay_account"] = request.ThirdPayAccount
	}

	if !dara.IsNil(request.ThirdTradeNo) {
		body["third_trade_no"] = request.ThirdTradeNo
	}

	if !dara.IsNil(request.TotalPrice) {
		body["total_price"] = request.TotalPrice
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelOrderPay"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/orders/action/pay"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelOrderPayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店下单前校验
//
// @param tmpReq - HotelOrderPreValidateRequest
//
// @param headers - HotelOrderPreValidateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelOrderPreValidateResponse
func (client *Client) HotelOrderPreValidateWithContext(ctx context.Context, tmpReq *HotelOrderPreValidateRequest, headers *HotelOrderPreValidateHeaders, runtime *dara.RuntimeOptions) (_result *HotelOrderPreValidateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &HotelOrderPreValidateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DailyList) {
		request.DailyListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DailyList, dara.String("daily_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OccupantInfoList) {
		request.OccupantInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OccupantInfoList, dara.String("occupant_info_list"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.CheckIn) {
		query["check_in"] = request.CheckIn
	}

	if !dara.IsNil(request.CheckOut) {
		query["check_out"] = request.CheckOut
	}

	if !dara.IsNil(request.DailyListShrink) {
		query["daily_list"] = request.DailyListShrink
	}

	if !dara.IsNil(request.ItemId) {
		query["item_id"] = request.ItemId
	}

	if !dara.IsNil(request.NumberOfAdultsPerRoom) {
		query["number_of_adults_per_room"] = request.NumberOfAdultsPerRoom
	}

	if !dara.IsNil(request.OccupantInfoListShrink) {
		query["occupant_info_list"] = request.OccupantInfoListShrink
	}

	if !dara.IsNil(request.RatePlanId) {
		query["rate_plan_id"] = request.RatePlanId
	}

	if !dara.IsNil(request.RoomId) {
		query["room_id"] = request.RoomId
	}

	if !dara.IsNil(request.RoomNum) {
		query["room_num"] = request.RoomNum
	}

	if !dara.IsNil(request.SearchRoomPrice) {
		query["search_room_price"] = request.SearchRoomPrice
	}

	if !dara.IsNil(request.SellerId) {
		query["seller_id"] = request.SellerId
	}

	if !dara.IsNil(request.Shid) {
		query["shid"] = request.Shid
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelOrderPreValidate"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/orders/action/pre-validate"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelOrderPreValidateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店订单查询
//
// @param request - HotelOrderQueryRequest
//
// @param headers - HotelOrderQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelOrderQueryResponse
func (client *Client) HotelOrderQueryWithContext(ctx context.Context, request *HotelOrderQueryRequest, headers *HotelOrderQueryHeaders, runtime *dara.RuntimeOptions) (_result *HotelOrderQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelOrderQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/hotel/v1/order"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelOrderQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店拉动态拉取价格接口(落地)
//
// @param tmpReq - HotelPricePullRequest
//
// @param headers - HotelPricePullHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelPricePullResponse
func (client *Client) HotelPricePullWithContext(ctx context.Context, tmpReq *HotelPricePullRequest, headers *HotelPricePullHeaders, runtime *dara.RuntimeOptions) (_result *HotelPricePullResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &HotelPricePullShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HotelIds) {
		request.HotelIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelIds, dara.String("hotel_ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.CheckIn) {
		query["check_in"] = request.CheckIn
	}

	if !dara.IsNil(request.CheckOut) {
		query["check_out"] = request.CheckOut
	}

	if !dara.IsNil(request.CityCode) {
		query["city_code"] = request.CityCode
	}

	if !dara.IsNil(request.HotelIdsShrink) {
		query["hotel_ids"] = request.HotelIdsShrink
	}

	if !dara.IsNil(request.PaymentType) {
		query["payment_type"] = request.PaymentType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelPricePull"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/prices/action/pull"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelPricePullResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店静态房型详情
//
// @param tmpReq - HotelRoomInfoRequest
//
// @param headers - HotelRoomInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelRoomInfoResponse
func (client *Client) HotelRoomInfoWithContext(ctx context.Context, tmpReq *HotelRoomInfoRequest, headers *HotelRoomInfoHeaders, runtime *dara.RuntimeOptions) (_result *HotelRoomInfoResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &HotelRoomInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoomIds) {
		request.RoomIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoomIds, dara.String("room_ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RoomIdsShrink) {
		query["room_ids"] = request.RoomIdsShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelRoomInfo"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/room-infos"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelRoomInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店列表搜索接口(直连)
//
// @param tmpReq - HotelSearchRequest
//
// @param headers - HotelSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelSearchResponse
func (client *Client) HotelSearchWithContext(ctx context.Context, tmpReq *HotelSearchRequest, headers *HotelSearchHeaders, runtime *dara.RuntimeOptions) (_result *HotelSearchResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &HotelSearchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BrandCode) {
		request.BrandCodeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BrandCode, dara.String("brand_code"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Shids) {
		request.ShidsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Shids, dara.String("shids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdultNum) {
		query["adult_num"] = request.AdultNum
	}

	if !dara.IsNil(request.BrandCodeShrink) {
		query["brand_code"] = request.BrandCodeShrink
	}

	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.CheckInDate) {
		query["check_in_date"] = request.CheckInDate
	}

	if !dara.IsNil(request.CheckOutDate) {
		query["check_out_date"] = request.CheckOutDate
	}

	if !dara.IsNil(request.CityCode) {
		query["city_code"] = request.CityCode
	}

	if !dara.IsNil(request.Dir) {
		query["dir"] = request.Dir
	}

	if !dara.IsNil(request.Distance) {
		query["distance"] = request.Distance
	}

	if !dara.IsNil(request.DistrictCode) {
		query["district_code"] = request.DistrictCode
	}

	if !dara.IsNil(request.HotelStar) {
		query["hotel_star"] = request.HotelStar
	}

	if !dara.IsNil(request.IsProtocol) {
		query["is_protocol"] = request.IsProtocol
	}

	if !dara.IsNil(request.KeyWords) {
		query["key_words"] = request.KeyWords
	}

	if !dara.IsNil(request.Location) {
		query["location"] = request.Location
	}

	if !dara.IsNil(request.MaxPrice) {
		query["max_price"] = request.MaxPrice
	}

	if !dara.IsNil(request.MinPrice) {
		query["min_price"] = request.MinPrice
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PayOverType) {
		query["pay_over_type"] = request.PayOverType
	}

	if !dara.IsNil(request.PaymentType) {
		query["payment_type"] = request.PaymentType
	}

	if !dara.IsNil(request.ShidsShrink) {
		query["shids"] = request.ShidsShrink
	}

	if !dara.IsNil(request.SortCode) {
		query["sort_code"] = request.SortCode
	}

	if !dara.IsNil(request.SuperMan) {
		query["super_man"] = request.SuperMan
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelSearch"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/hotels/action/search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询酒店静态详情
//
// @param tmpReq - HotelStaticInfoRequest
//
// @param headers - HotelStaticInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelStaticInfoResponse
func (client *Client) HotelStaticInfoWithContext(ctx context.Context, tmpReq *HotelStaticInfoRequest, headers *HotelStaticInfoHeaders, runtime *dara.RuntimeOptions) (_result *HotelStaticInfoResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &HotelStaticInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HotelIds) {
		request.HotelIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelIds, dara.String("hotel_ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.HotelIdsShrink) {
		query["hotel_ids"] = request.HotelIdsShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelStaticInfo"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v1/static-infos"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelStaticInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店关键词搜索
//
// @param request - HotelSuggestV2Request
//
// @param headers - HotelSuggestV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelSuggestV2Response
func (client *Client) HotelSuggestV2WithContext(ctx context.Context, request *HotelSuggestV2Request, headers *HotelSuggestV2Headers, runtime *dara.RuntimeOptions) (_result *HotelSuggestV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.CheckIn) {
		query["check_in"] = request.CheckIn
	}

	if !dara.IsNil(request.CheckOut) {
		query["check_out"] = request.CheckOut
	}

	if !dara.IsNil(request.CityCode) {
		query["city_code"] = request.CityCode
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.SearchType) {
		query["search_type"] = request.SearchType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelSuggestV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-hotel/v2/suggest-infos"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelSuggestV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票订单详情
//
// @param request - IFlightOrderDetailQueryRequest
//
// @param headers - IFlightOrderDetailQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IFlightOrderDetailQueryResponse
func (client *Client) IFlightOrderDetailQueryWithContext(ctx context.Context, request *IFlightOrderDetailQueryRequest, headers *IFlightOrderDetailQueryHeaders, runtime *dara.RuntimeOptions) (_result *IFlightOrderDetailQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IFlightOrderDetailQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/v1/intlFlight-order-detail-query"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IFlightOrderDetailQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票订单列表
//
// @param tmpReq - IFlightOrderListQueryRequest
//
// @param headers - IFlightOrderListQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IFlightOrderListQueryResponse
func (client *Client) IFlightOrderListQueryWithContext(ctx context.Context, tmpReq *IFlightOrderListQueryRequest, headers *IFlightOrderListQueryHeaders, runtime *dara.RuntimeOptions) (_result *IFlightOrderListQueryResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &IFlightOrderListQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplyIdList) {
		request.ApplyIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplyIdList, dara.String("apply_id_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.BookTypeList) {
		request.BookTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BookTypeList, dara.String("book_type_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.BookerId) {
		request.BookerIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BookerId, dara.String("booker_id"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ThirdPartApplyIdList) {
		request.ThirdPartApplyIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ThirdPartApplyIdList, dara.String("third_part_apply_id_list"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyIdListShrink) {
		query["apply_id_list"] = request.ApplyIdListShrink
	}

	if !dara.IsNil(request.BookTypeListShrink) {
		query["book_type_list"] = request.BookTypeListShrink
	}

	if !dara.IsNil(request.BookerIdShrink) {
		query["booker_id"] = request.BookerIdShrink
	}

	if !dara.IsNil(request.EndDate) {
		query["end_date"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.ScrollId) {
		query["scroll_id"] = request.ScrollId
	}

	if !dara.IsNil(request.StartDate) {
		query["start_date"] = request.StartDate
	}

	if !dara.IsNil(request.ThirdPartApplyIdListShrink) {
		query["third_part_apply_id_list"] = request.ThirdPartApplyIdListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IFlightOrderListQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/v1/intlFlight-order-list-query"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IFlightOrderListQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询国际机票记账数据
//
// @param request - IeFlightBillSettlementQueryRequest
//
// @param headers - IeFlightBillSettlementQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IeFlightBillSettlementQueryResponse
func (client *Client) IeFlightBillSettlementQueryWithContext(ctx context.Context, request *IeFlightBillSettlementQueryRequest, headers *IeFlightBillSettlementQueryHeaders, runtime *dara.RuntimeOptions) (_result *IeFlightBillSettlementQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillBatch) {
		query["bill_batch"] = request.BillBatch
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PeriodEnd) {
		query["period_end"] = request.PeriodEnd
	}

	if !dara.IsNil(request.PeriodStart) {
		query["period_start"] = request.PeriodStart
	}

	if !dara.IsNil(request.ScrollId) {
		query["scroll_id"] = request.ScrollId
	}

	if !dara.IsNil(request.ScrollMod) {
		query["scroll_mod"] = request.ScrollMod
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IeFlightBillSettlementQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ie-flight/v1/bill-settlement"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IeFlightBillSettlementQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询国际/中国港澳台酒店记账数据
//
// @param request - IeHotelBillSettlementQueryRequest
//
// @param headers - IeHotelBillSettlementQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IeHotelBillSettlementQueryResponse
func (client *Client) IeHotelBillSettlementQueryWithContext(ctx context.Context, request *IeHotelBillSettlementQueryRequest, headers *IeHotelBillSettlementQueryHeaders, runtime *dara.RuntimeOptions) (_result *IeHotelBillSettlementQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillBatch) {
		query["bill_batch"] = request.BillBatch
	}

	if !dara.IsNil(request.Category) {
		query["category"] = request.Category
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PeriodEnd) {
		query["period_end"] = request.PeriodEnd
	}

	if !dara.IsNil(request.PeriodStart) {
		query["period_start"] = request.PeriodStart
	}

	if !dara.IsNil(request.ScrollId) {
		query["scroll_id"] = request.ScrollId
	}

	if !dara.IsNil(request.ScrollMod) {
		query["scroll_mod"] = request.ScrollMod
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IeHotelBillSettlementQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ie-hotel/v1/bill-settlement"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IeHotelBillSettlementQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询保险电子发票
//
// @param request - InsInvoiceScanQueryRequest
//
// @param headers - InsInvoiceScanQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsInvoiceScanQueryResponse
func (client *Client) InsInvoiceScanQueryWithContext(ctx context.Context, request *InsInvoiceScanQueryRequest, headers *InsInvoiceScanQueryHeaders, runtime *dara.RuntimeOptions) (_result *InsInvoiceScanQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillDate) {
		query["bill_date"] = request.BillDate
	}

	if !dara.IsNil(request.BillId) {
		query["bill_id"] = request.BillId
	}

	if !dara.IsNil(request.InvoiceSubTaskId) {
		query["invoice_sub_task_id"] = request.InvoiceSubTaskId
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsInvoiceScanQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/scan/v1/ins-invoice"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsInvoiceScanQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保险订单申请
//
// @param request - InsureOrderApplyRequest
//
// @param headers - InsureOrderApplyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsureOrderApplyResponse
func (client *Client) InsureOrderApplyWithContext(ctx context.Context, request *InsureOrderApplyRequest, headers *InsureOrderApplyHeaders, runtime *dara.RuntimeOptions) (_result *InsureOrderApplyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		body["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		body["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.InsOrderId) {
		body["ins_order_id"] = request.InsOrderId
	}

	if !dara.IsNil(request.IsvName) {
		body["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutSubOrderId) {
		body["out_sub_order_id"] = request.OutSubOrderId
	}

	if !dara.IsNil(request.SupplierCode) {
		body["supplier_code"] = request.SupplierCode
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsureOrderApply"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/insurances/action/apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsureOrderApplyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保险订单取消
//
// @param request - InsureOrderCancelRequest
//
// @param headers - InsureOrderCancelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsureOrderCancelResponse
func (client *Client) InsureOrderCancelWithContext(ctx context.Context, insOrderId *string, request *InsureOrderCancelRequest, headers *InsureOrderCancelHeaders, runtime *dara.RuntimeOptions) (_result *InsureOrderCancelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		query["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.SupplierCode) {
		query["supplier_code"] = request.SupplierCode
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsureOrderCancel"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/insurances/" + dara.PercentEncode(dara.StringValue(insOrderId)) + "/action/cancel"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsureOrderCancelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保险订单创建
//
// @param tmpReq - InsureOrderCreateRequest
//
// @param headers - InsureOrderCreateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsureOrderCreateResponse
func (client *Client) InsureOrderCreateWithContext(ctx context.Context, tmpReq *InsureOrderCreateRequest, headers *InsureOrderCreateHeaders, runtime *dara.RuntimeOptions) (_result *InsureOrderCreateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &InsureOrderCreateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Applicant) {
		request.ApplicantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Applicant, dara.String("applicant"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InsPersonAndSegmentList) {
		request.InsPersonAndSegmentListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InsPersonAndSegmentList, dara.String("ins_person_and_segment_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicantShrink) {
		body["applicant"] = request.ApplicantShrink
	}

	if !dara.IsNil(request.BtripUserId) {
		body["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		body["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.InsPersonAndSegmentListShrink) {
		body["ins_person_and_segment_list"] = request.InsPersonAndSegmentListShrink
	}

	if !dara.IsNil(request.IsvName) {
		body["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OutInsOrderId) {
		body["out_ins_order_id"] = request.OutInsOrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutSubOrderId) {
		body["out_sub_order_id"] = request.OutSubOrderId
	}

	if !dara.IsNil(request.SupplierCode) {
		body["supplier_code"] = request.SupplierCode
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsureOrderCreate"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/insurances/action/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsureOrderCreateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保险订单查询
//
// @param request - InsureOrderDetailRequest
//
// @param headers - InsureOrderDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsureOrderDetailResponse
func (client *Client) InsureOrderDetailWithContext(ctx context.Context, request *InsureOrderDetailRequest, headers *InsureOrderDetailHeaders, runtime *dara.RuntimeOptions) (_result *InsureOrderDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		query["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.InsOrderId) {
		query["ins_order_id"] = request.InsOrderId
	}

	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.SupplierCode) {
		query["supplier_code"] = request.SupplierCode
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsureOrderDetail"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/insurances/action/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsureOrderDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保险订单支付
//
// @param request - InsureOrderPayRequest
//
// @param headers - InsureOrderPayHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsureOrderPayResponse
func (client *Client) InsureOrderPayWithContext(ctx context.Context, insOrderId *string, request *InsureOrderPayRequest, headers *InsureOrderPayHeaders, runtime *dara.RuntimeOptions) (_result *InsureOrderPayResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		body["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		body["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.IsvName) {
		body["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutSubOrderId) {
		body["out_sub_order_id"] = request.OutSubOrderId
	}

	if !dara.IsNil(request.PaymentAmount) {
		body["payment_amount"] = request.PaymentAmount
	}

	if !dara.IsNil(request.SupplierCode) {
		body["supplier_code"] = request.SupplierCode
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsureOrderPay"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/insurances/" + dara.PercentEncode(dara.StringValue(insOrderId)) + "/action/pay"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsureOrderPayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保险订单退保
//
// @param tmpReq - InsureOrderRefundRequest
//
// @param headers - InsureOrderRefundHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsureOrderRefundResponse
func (client *Client) InsureOrderRefundWithContext(ctx context.Context, insOrderId *string, tmpReq *InsureOrderRefundRequest, headers *InsureOrderRefundHeaders, runtime *dara.RuntimeOptions) (_result *InsureOrderRefundResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &InsureOrderRefundShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PolicyNoList) {
		request.PolicyNoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PolicyNoList, dara.String("policy_no_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SubInsOrderIds) {
		request.SubInsOrderIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubInsOrderIds, dara.String("sub_ins_order_ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		body["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		body["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.IsvName) {
		body["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OutApplyId) {
		body["out_apply_id"] = request.OutApplyId
	}

	if !dara.IsNil(request.PolicyNoListShrink) {
		body["policy_no_list"] = request.PolicyNoListShrink
	}

	if !dara.IsNil(request.SubInsOrderIdsShrink) {
		body["sub_ins_order_ids"] = request.SubInsOrderIdsShrink
	}

	if !dara.IsNil(request.SupplierCode) {
		body["supplier_code"] = request.SupplierCode
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsureOrderRefund"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/insurances/" + dara.PercentEncode(dara.StringValue(insOrderId)) + "/action/refund"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsureOrderRefundResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询保单详情链接
//
// @param headers - InsureOrderUrlDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsureOrderUrlDetailResponse
func (client *Client) InsureOrderUrlDetailWithContext(ctx context.Context, insOrderId *string, headers *InsureOrderUrlDetailHeaders, runtime *dara.RuntimeOptions) (_result *InsureOrderUrlDetailResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsureOrderUrlDetail"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/insurances/" + dara.PercentEncode(dara.StringValue(insOrderId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsureOrderUrlDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 退保详情查询
//
// @param request - InsureRefundDetailRequest
//
// @param headers - InsureRefundDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsureRefundDetailResponse
func (client *Client) InsureRefundDetailWithContext(ctx context.Context, request *InsureRefundDetailRequest, headers *InsureRefundDetailHeaders, runtime *dara.RuntimeOptions) (_result *InsureRefundDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyId) {
		query["apply_id"] = request.ApplyId
	}

	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		query["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.InsOrderId) {
		query["ins_order_id"] = request.InsOrderId
	}

	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OutApplyId) {
		query["out_apply_id"] = request.OutApplyId
	}

	if !dara.IsNil(request.SupplierCode) {
		query["supplier_code"] = request.SupplierCode
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsureRefundDetail"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/insurances/action/refund-detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsureRefundDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票创建订单
//
// @param tmpReq - IntlFlightCreateOrderRequest
//
// @param headers - IntlFlightCreateOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightCreateOrderResponse
func (client *Client) IntlFlightCreateOrderWithContext(ctx context.Context, tmpReq *IntlFlightCreateOrderRequest, headers *IntlFlightCreateOrderHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightCreateOrderResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &IntlFlightCreateOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ContactInfo) {
		request.ContactInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContactInfo, dara.String("contact_info"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExtraInfo) {
		request.ExtraInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtraInfo, dara.String("extra_info"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PassengerList) {
		request.PassengerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PassengerList, dara.String("passenger_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AsyncCreateOrderKey) {
		body["async_create_order_key"] = request.AsyncCreateOrderKey
	}

	if !dara.IsNil(request.AsyncCreateOrderMode) {
		body["async_create_order_mode"] = request.AsyncCreateOrderMode
	}

	if !dara.IsNil(request.BtripUserId) {
		body["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		body["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.ContactInfoShrink) {
		body["contact_info"] = request.ContactInfoShrink
	}

	if !dara.IsNil(request.ExtraInfoShrink) {
		body["extra_info"] = request.ExtraInfoShrink
	}

	if !dara.IsNil(request.IsvName) {
		body["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderPrice) {
		body["order_price"] = request.OrderPrice
	}

	if !dara.IsNil(request.OtaItemId) {
		body["ota_item_id"] = request.OtaItemId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.PassengerListShrink) {
		body["passenger_list"] = request.PassengerListShrink
	}

	if !dara.IsNil(request.RenderKey) {
		body["render_key"] = request.RenderKey
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightCreateOrder"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/order/action/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightCreateOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票验舱验价
//
// @param tmpReq - IntlFlightInventoryPriceCheckRequest
//
// @param headers - IntlFlightInventoryPriceCheckHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightInventoryPriceCheckResponse
func (client *Client) IntlFlightInventoryPriceCheckWithContext(ctx context.Context, tmpReq *IntlFlightInventoryPriceCheckRequest, headers *IntlFlightInventoryPriceCheckHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightInventoryPriceCheckResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &IntlFlightInventoryPriceCheckShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PassengerList) {
		request.PassengerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PassengerList, dara.String("passenger_list"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		query["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderPrice) {
		query["order_price"] = request.OrderPrice
	}

	if !dara.IsNil(request.OtaItemId) {
		query["ota_item_id"] = request.OtaItemId
	}

	if !dara.IsNil(request.PassengerListShrink) {
		query["passenger_list"] = request.PassengerListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightInventoryPriceCheck"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/flights/action/inventory-price-check"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightInventoryPriceCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票航班搜索
//
// @param tmpReq - IntlFlightListingSearchRequest
//
// @param headers - IntlFlightListingSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightListingSearchResponse
func (client *Client) IntlFlightListingSearchWithContext(ctx context.Context, tmpReq *IntlFlightListingSearchRequest, headers *IntlFlightListingSearchHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightListingSearchResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &IntlFlightListingSearchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SearchJourneys) {
		request.SearchJourneysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchJourneys, dara.String("search_journeys"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SearchPassengerList) {
		request.SearchPassengerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchPassengerList, dara.String("search_passenger_list"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		query["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.CabinType) {
		query["cabin_type"] = request.CabinType
	}

	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OutWheelSearch) {
		query["out_wheel_search"] = request.OutWheelSearch
	}

	if !dara.IsNil(request.QueryRecordId) {
		query["query_record_id"] = request.QueryRecordId
	}

	if !dara.IsNil(request.SearchJourneysShrink) {
		query["search_journeys"] = request.SearchJourneysShrink
	}

	if !dara.IsNil(request.SearchMode) {
		query["search_mode"] = request.SearchMode
	}

	if !dara.IsNil(request.SearchPassengerListShrink) {
		query["search_passenger_list"] = request.SearchPassengerListShrink
	}

	if !dara.IsNil(request.Token) {
		query["token"] = request.Token
	}

	if !dara.IsNil(request.TripType) {
		query["trip_type"] = request.TripType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightListingSearch"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/flights/action/listing-search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightListingSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票订单取消
//
// @param request - IntlFlightOrderCancelRequest
//
// @param headers - IntlFlightOrderCancelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightOrderCancelResponse
func (client *Client) IntlFlightOrderCancelWithContext(ctx context.Context, request *IntlFlightOrderCancelRequest, headers *IntlFlightOrderCancelHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightOrderCancelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		body["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		body["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.IsvName) {
		body["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightOrderCancel"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/order/action/cancel"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightOrderCancelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票订单详情
//
// @param request - IntlFlightOrderDetailRequest
//
// @param headers - IntlFlightOrderDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightOrderDetailResponse
func (client *Client) IntlFlightOrderDetailWithContext(ctx context.Context, request *IntlFlightOrderDetailRequest, headers *IntlFlightOrderDetailHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightOrderDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		query["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["out_order_id"] = request.OutOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightOrderDetail"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/order/action/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightOrderDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票订单支付
//
// @param request - IntlFlightOrderPayRequest
//
// @param headers - IntlFlightOrderPayHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightOrderPayResponse
func (client *Client) IntlFlightOrderPayWithContext(ctx context.Context, request *IntlFlightOrderPayRequest, headers *IntlFlightOrderPayHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightOrderPayResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		body["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		body["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.IsvName) {
		body["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OrderPrice) {
		body["order_price"] = request.OrderPrice
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightOrderPay"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/order/action/pay"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightOrderPayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票订单支付前校验
//
// @param request - IntlFlightOrderPayCheckRequest
//
// @param headers - IntlFlightOrderPayCheckHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightOrderPayCheckResponse
func (client *Client) IntlFlightOrderPayCheckWithContext(ctx context.Context, request *IntlFlightOrderPayCheckRequest, headers *IntlFlightOrderPayCheckHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightOrderPayCheckResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		query["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["out_order_id"] = request.OutOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightOrderPayCheck"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/order/action/pay-check"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightOrderPayCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票报价商品详情
//
// @param request - IntlFlightOtaItemDetailRequest
//
// @param headers - IntlFlightOtaItemDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightOtaItemDetailResponse
func (client *Client) IntlFlightOtaItemDetailWithContext(ctx context.Context, otaItemId *string, request *IntlFlightOtaItemDetailRequest, headers *IntlFlightOtaItemDetailHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightOtaItemDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		query["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.Language) {
		query["language"] = request.Language
	}

	if !dara.IsNil(request.SupplierCode) {
		query["supplier_code"] = request.SupplierCode
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightOtaItemDetail"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/items/" + dara.PercentEncode(dara.StringValue(otaItemId)) + "/action/ota-get"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightOtaItemDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票航班报价查询
//
// @param tmpReq - IntlFlightOtaSearchRequest
//
// @param headers - IntlFlightOtaSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightOtaSearchResponse
func (client *Client) IntlFlightOtaSearchWithContext(ctx context.Context, tmpReq *IntlFlightOtaSearchRequest, headers *IntlFlightOtaSearchHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightOtaSearchResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &IntlFlightOtaSearchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SearchJourneys) {
		request.SearchJourneysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchJourneys, dara.String("search_journeys"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SearchPassengerList) {
		request.SearchPassengerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchPassengerList, dara.String("search_passenger_list"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BtripUserId) {
		query["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BuyerName) {
		query["buyer_name"] = request.BuyerName
	}

	if !dara.IsNil(request.CabinType) {
		query["cabin_type"] = request.CabinType
	}

	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.SearchJourneysShrink) {
		query["search_journeys"] = request.SearchJourneysShrink
	}

	if !dara.IsNil(request.SearchPassengerListShrink) {
		query["search_passenger_list"] = request.SearchPassengerListShrink
	}

	if !dara.IsNil(request.TripType) {
		query["trip_type"] = request.TripType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightOtaSearch"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/flights/action/ota-search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightOtaSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票改签申请
//
// @param tmpReq - IntlFlightReShopApplyRequest
//
// @param headers - IntlFlightReShopApplyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightReShopApplyResponse
func (client *Client) IntlFlightReShopApplyWithContext(ctx context.Context, tmpReq *IntlFlightReShopApplyRequest, headers *IntlFlightReShopApplyHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightReShopApplyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &IntlFlightReShopApplyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SelectedJourneys) {
		request.SelectedJourneysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SelectedJourneys, dara.String("selected_journeys"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SelectedPassengers) {
		request.SelectedPassengersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SelectedPassengers, dara.String("selected_passengers"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AsyncApplyKey) {
		body["async_apply_key"] = request.AsyncApplyKey
	}

	if !dara.IsNil(request.AsyncApplyMode) {
		body["async_apply_mode"] = request.AsyncApplyMode
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutReShopApplyId) {
		body["out_re_shop_apply_id"] = request.OutReShopApplyId
	}

	if !dara.IsNil(request.PassengerJourneyGroupKey) {
		body["passenger_journey_group_key"] = request.PassengerJourneyGroupKey
	}

	if !dara.IsNil(request.ReShopReasonCode) {
		body["re_shop_reason_code"] = request.ReShopReasonCode
	}

	if !dara.IsNil(request.SelectedJourneysShrink) {
		body["selected_journeys"] = request.SelectedJourneysShrink
	}

	if !dara.IsNil(request.SelectedPassengersShrink) {
		body["selected_passengers"] = request.SelectedPassengersShrink
	}

	if !dara.IsNil(request.UserIntentionMemo) {
		body["user_intention_memo"] = request.UserIntentionMemo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightReShopApply"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/flights/action/reshop/apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightReShopApplyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票改签取消
//
// @param request - IntlFlightReShopCancelRequest
//
// @param headers - IntlFlightReShopCancelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightReShopCancelResponse
func (client *Client) IntlFlightReShopCancelWithContext(ctx context.Context, request *IntlFlightReShopCancelRequest, headers *IntlFlightReShopCancelHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightReShopCancelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutReShopApplyId) {
		body["out_re_shop_apply_id"] = request.OutReShopApplyId
	}

	if !dara.IsNil(request.ReShopApplyId) {
		body["re_shop_apply_id"] = request.ReShopApplyId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightReShopCancel"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/flights/action/reshop/cancel"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightReShopCancelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票改签咨询
//
// @param request - IntlFlightReShopConsultRequest
//
// @param headers - IntlFlightReShopConsultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightReShopConsultResponse
func (client *Client) IntlFlightReShopConsultWithContext(ctx context.Context, request *IntlFlightReShopConsultRequest, headers *IntlFlightReShopConsultHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightReShopConsultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["out_order_id"] = request.OutOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightReShopConsult"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/flights/action/reshop/consult"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightReShopConsultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票改签详情
//
// @param request - IntlFlightReShopDetailRequest
//
// @param headers - IntlFlightReShopDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightReShopDetailResponse
func (client *Client) IntlFlightReShopDetailWithContext(ctx context.Context, request *IntlFlightReShopDetailRequest, headers *IntlFlightReShopDetailHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightReShopDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutReShopApplyId) {
		query["out_re_shop_apply_id"] = request.OutReShopApplyId
	}

	if !dara.IsNil(request.ReShopApplyId) {
		query["re_shop_apply_id"] = request.ReShopApplyId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightReShopDetail"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/flights/action/reshop/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightReShopDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票改签支付
//
// @param request - IntlFlightReShopPayRequest
//
// @param headers - IntlFlightReShopPayHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightReShopPayResponse
func (client *Client) IntlFlightReShopPayWithContext(ctx context.Context, request *IntlFlightReShopPayRequest, headers *IntlFlightReShopPayHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightReShopPayResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutReShopApplyId) {
		body["out_re_shop_apply_id"] = request.OutReShopApplyId
	}

	if !dara.IsNil(request.ReShopApplyId) {
		body["re_shop_apply_id"] = request.ReShopApplyId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightReShopPay"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/flights/action/reshop/pay"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightReShopPayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票退票申请
//
// @param tmpReq - IntlFlightRefundApplyRequest
//
// @param headers - IntlFlightRefundApplyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightRefundApplyResponse
func (client *Client) IntlFlightRefundApplyWithContext(ctx context.Context, tmpReq *IntlFlightRefundApplyRequest, headers *IntlFlightRefundApplyHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightRefundApplyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &IntlFlightRefundApplyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RefundSegmentList) {
		request.RefundSegmentListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RefundSegmentList, dara.String("refund_segment_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SelectedPassengers) {
		request.SelectedPassengersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SelectedPassengers, dara.String("selected_passengers"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutRefundApplyId) {
		body["out_refund_apply_id"] = request.OutRefundApplyId
	}

	if !dara.IsNil(request.PassengerJourneyGroupKey) {
		body["passenger_journey_group_key"] = request.PassengerJourneyGroupKey
	}

	if !dara.IsNil(request.RefundReasonCode) {
		body["refund_reason_code"] = request.RefundReasonCode
	}

	if !dara.IsNil(request.RefundSegmentListShrink) {
		body["refund_segment_list"] = request.RefundSegmentListShrink
	}

	if !dara.IsNil(request.SelectedPassengersShrink) {
		body["selected_passengers"] = request.SelectedPassengersShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightRefundApply"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/flights/action/refund/apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightRefundApplyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票退票咨询
//
// @param request - IntlFlightRefundConsultRequest
//
// @param headers - IntlFlightRefundConsultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightRefundConsultResponse
func (client *Client) IntlFlightRefundConsultWithContext(ctx context.Context, request *IntlFlightRefundConsultRequest, headers *IntlFlightRefundConsultHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightRefundConsultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["out_order_id"] = request.OutOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightRefundConsult"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/flights/action/refund/consult"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightRefundConsultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票退票详情
//
// @param request - IntlFlightRefundDetailRequest
//
// @param headers - IntlFlightRefundDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightRefundDetailResponse
func (client *Client) IntlFlightRefundDetailWithContext(ctx context.Context, request *IntlFlightRefundDetailRequest, headers *IntlFlightRefundDetailHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightRefundDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		query["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutRefundApplyId) {
		query["out_refund_apply_id"] = request.OutRefundApplyId
	}

	if !dara.IsNil(request.RefundApplyId) {
		query["refund_apply_id"] = request.RefundApplyId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightRefundDetail"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/flights/action/refund/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightRefundDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际机票航班可用证件查询
//
// @param request - IntlFlightSegmentAvailableCertRequest
//
// @param headers - IntlFlightSegmentAvailableCertHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntlFlightSegmentAvailableCertResponse
func (client *Client) IntlFlightSegmentAvailableCertWithContext(ctx context.Context, otaItemId *string, request *IntlFlightSegmentAvailableCertRequest, headers *IntlFlightSegmentAvailableCertHeaders, runtime *dara.RuntimeOptions) (_result *IntlFlightSegmentAvailableCertResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsvName) {
		query["isv_name"] = request.IsvName
	}

	if !dara.IsNil(request.Language) {
		query["language"] = request.Language
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		query["user_name"] = request.UserName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntlFlightSegmentAvailableCert"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/intl-flight/v1/items/" + dara.PercentEncode(dara.StringValue(otaItemId)) + "/action/segment-available-cert"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IntlFlightSegmentAvailableCertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增发票配置
//
// @param request - InvoiceAddRequest
//
// @param headers - InvoiceAddHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvoiceAddResponse
func (client *Client) InvoiceAddWithContext(ctx context.Context, request *InvoiceAddRequest, headers *InvoiceAddHeaders, runtime *dara.RuntimeOptions) (_result *InvoiceAddResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		body["address"] = request.Address
	}

	if !dara.IsNil(request.BankName) {
		body["bank_name"] = request.BankName
	}

	if !dara.IsNil(request.BankNo) {
		body["bank_no"] = request.BankNo
	}

	if !dara.IsNil(request.TaxNo) {
		body["tax_no"] = request.TaxNo
	}

	if !dara.IsNil(request.Tel) {
		body["tel"] = request.Tel
	}

	if !dara.IsNil(request.ThirdPartId) {
		body["third_part_id"] = request.ThirdPartId
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	if !dara.IsNil(request.UnitType) {
		body["unit_type"] = request.UnitType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvoiceAdd"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/invoice/v1/add-invoice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvoiceAddResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除发票抬头
//
// @param request - InvoiceDeleteRequest
//
// @param headers - InvoiceDeleteHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvoiceDeleteResponse
func (client *Client) InvoiceDeleteWithContext(ctx context.Context, request *InvoiceDeleteRequest, headers *InvoiceDeleteHeaders, runtime *dara.RuntimeOptions) (_result *InvoiceDeleteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ThirdPartId) {
		query["third_part_id"] = request.ThirdPartId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvoiceDelete"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/invoice/v1/invoice"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InvoiceDeleteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改发票配置
//
// @param request - InvoiceModifyRequest
//
// @param headers - InvoiceModifyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvoiceModifyResponse
func (client *Client) InvoiceModifyWithContext(ctx context.Context, request *InvoiceModifyRequest, headers *InvoiceModifyHeaders, runtime *dara.RuntimeOptions) (_result *InvoiceModifyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		body["address"] = request.Address
	}

	if !dara.IsNil(request.BankName) {
		body["bank_name"] = request.BankName
	}

	if !dara.IsNil(request.BankNo) {
		body["bank_no"] = request.BankNo
	}

	if !dara.IsNil(request.TaxNo) {
		body["tax_no"] = request.TaxNo
	}

	if !dara.IsNil(request.Tel) {
		body["tel"] = request.Tel
	}

	if !dara.IsNil(request.ThirdPartId) {
		body["third_part_id"] = request.ThirdPartId
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	if !dara.IsNil(request.UnitType) {
		body["unit_type"] = request.UnitType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvoiceModify"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/invoice/v1/invoice"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvoiceModifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增发票抬头可用员工
//
// @param tmpReq - InvoiceRuleAddRequest
//
// @param headers - InvoiceRuleAddHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvoiceRuleAddResponse
func (client *Client) InvoiceRuleAddWithContext(ctx context.Context, tmpReq *InvoiceRuleAddRequest, headers *InvoiceRuleAddHeaders, runtime *dara.RuntimeOptions) (_result *InvoiceRuleAddResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &InvoiceRuleAddShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Entities) {
		request.EntitiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Entities, dara.String("entities"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EntitiesShrink) {
		body["entities"] = request.EntitiesShrink
	}

	if !dara.IsNil(request.ThirdPartId) {
		body["third_part_id"] = request.ThirdPartId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvoiceRuleAdd"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/invoice/v1/invoice-rule"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvoiceRuleAddResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除发票抬头可用员工
//
// @param tmpReq - InvoiceRuleDeleteRequest
//
// @param headers - InvoiceRuleDeleteHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvoiceRuleDeleteResponse
func (client *Client) InvoiceRuleDeleteWithContext(ctx context.Context, tmpReq *InvoiceRuleDeleteRequest, headers *InvoiceRuleDeleteHeaders, runtime *dara.RuntimeOptions) (_result *InvoiceRuleDeleteResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &InvoiceRuleDeleteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Entities) {
		request.EntitiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Entities, dara.String("entities"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DelAll) {
		query["del_all"] = request.DelAll
	}

	if !dara.IsNil(request.EntitiesShrink) {
		query["entities"] = request.EntitiesShrink
	}

	if !dara.IsNil(request.ThirdPartId) {
		query["third_part_id"] = request.ThirdPartId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvoiceRuleDelete"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/invoice/v1/invoice-rule"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InvoiceRuleDeleteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存发票规则
//
// @param tmpReq - InvoiceRuleSaveRequest
//
// @param headers - InvoiceRuleSaveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvoiceRuleSaveResponse
func (client *Client) InvoiceRuleSaveWithContext(ctx context.Context, tmpReq *InvoiceRuleSaveRequest, headers *InvoiceRuleSaveHeaders, runtime *dara.RuntimeOptions) (_result *InvoiceRuleSaveResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &InvoiceRuleSaveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Entities) {
		request.EntitiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Entities, dara.String("entities"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllEmploye) {
		body["all_employe"] = request.AllEmploye
	}

	if !dara.IsNil(request.EntitiesShrink) {
		body["entities"] = request.EntitiesShrink
	}

	if !dara.IsNil(request.Scope) {
		body["scope"] = request.Scope
	}

	if !dara.IsNil(request.ThirdPartId) {
		body["third_part_id"] = request.ThirdPartId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvoiceRuleSave"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/invoice/v1/invoice-rule"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvoiceRuleSaveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索用户可用发票抬头
//
// @param request - InvoiceSearchRequest
//
// @param headers - InvoiceSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvoiceSearchResponse
func (client *Client) InvoiceSearchWithContext(ctx context.Context, request *InvoiceSearchRequest, headers *InvoiceSearchHeaders, runtime *dara.RuntimeOptions) (_result *InvoiceSearchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ThirdPartId) {
		query["third_part_id"] = request.ThirdPartId
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvoiceSearch"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/invoice/v1/invoice"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InvoiceSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 员工特殊角色修改
//
// @param tmpReq - IsvRuleSaveRequest
//
// @param headers - IsvRuleSaveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IsvRuleSaveResponse
func (client *Client) IsvRuleSaveWithContext(ctx context.Context, tmpReq *IsvRuleSaveRequest, headers *IsvRuleSaveHeaders, runtime *dara.RuntimeOptions) (_result *IsvRuleSaveResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &IsvRuleSaveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BookuserList) {
		request.BookuserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BookuserList, dara.String("bookuser_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplyNeed) {
		body["apply_need"] = request.ApplyNeed
	}

	if !dara.IsNil(request.BookType) {
		body["book_type"] = request.BookType
	}

	if !dara.IsNil(request.BookuserListShrink) {
		body["bookuser_list"] = request.BookuserListShrink
	}

	if !dara.IsNil(request.RuleNeed) {
		body["rule_need"] = request.RuleNeed
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IsvRuleSave"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/user/v1/rule"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IsvRuleSaveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户同步
//
// @param tmpReq - IsvUserSaveRequest
//
// @param headers - IsvUserSaveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IsvUserSaveResponse
func (client *Client) IsvUserSaveWithContext(ctx context.Context, tmpReq *IsvUserSaveRequest, headers *IsvUserSaveHeaders, runtime *dara.RuntimeOptions) (_result *IsvUserSaveResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &IsvUserSaveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserList) {
		request.UserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserList, dara.String("user_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UserListShrink) {
		body["user_list"] = request.UserListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IsvUserSave"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/isvuser/v1/isvuser"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IsvUserSaveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增用餐申请单
//
// @param tmpReq - MealApplyAddRequest
//
// @param headers - MealApplyAddHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MealApplyAddResponse
func (client *Client) MealApplyAddWithContext(ctx context.Context, tmpReq *MealApplyAddRequest, headers *MealApplyAddHeaders, runtime *dara.RuntimeOptions) (_result *MealApplyAddResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &MealApplyAddShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplyUser) {
		request.ApplyUserShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplyUser, dara.String("apply_user"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ItineraryList) {
		request.ItineraryListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItineraryList, dara.String("itinerary_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplyUserShrink) {
		body["apply_user"] = request.ApplyUserShrink
	}

	if !dara.IsNil(request.CostCenterId) {
		body["cost_center_id"] = request.CostCenterId
	}

	if !dara.IsNil(request.InvoiceId) {
		body["invoice_id"] = request.InvoiceId
	}

	if !dara.IsNil(request.ItineraryListShrink) {
		body["itinerary_list"] = request.ItineraryListShrink
	}

	if !dara.IsNil(request.MealAmount) {
		body["meal_amount"] = request.MealAmount
	}

	if !dara.IsNil(request.MealCause) {
		body["meal_cause"] = request.MealCause
	}

	if !dara.IsNil(request.ProjectCode) {
		body["project_code"] = request.ProjectCode
	}

	if !dara.IsNil(request.ProjectTitle) {
		body["project_title"] = request.ProjectTitle
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.ThirdPartApplyId) {
		body["third_part_apply_id"] = request.ThirdPartApplyId
	}

	if !dara.IsNil(request.ThirdPartCostCenterId) {
		body["third_part_cost_center_id"] = request.ThirdPartCostCenterId
	}

	if !dara.IsNil(request.ThirdPartInvoiceId) {
		body["third_part_invoice_id"] = request.ThirdPartInvoiceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MealApplyAdd"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/meal"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MealApplyAddResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新用餐申请单状态
//
// @param request - MealApplyApproveRequest
//
// @param headers - MealApplyApproveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MealApplyApproveResponse
func (client *Client) MealApplyApproveWithContext(ctx context.Context, request *MealApplyApproveRequest, headers *MealApplyApproveHeaders, runtime *dara.RuntimeOptions) (_result *MealApplyApproveResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OperateTime) {
		body["operate_time"] = request.OperateTime
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.ThirdPartApplyId) {
		body["third_part_apply_id"] = request.ThirdPartApplyId
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MealApplyApprove"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/meal"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MealApplyApproveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用餐申请单
//
// @param request - MealApplyQueryRequest
//
// @param headers - MealApplyQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MealApplyQueryResponse
func (client *Client) MealApplyQueryWithContext(ctx context.Context, request *MealApplyQueryRequest, headers *MealApplyQueryHeaders, runtime *dara.RuntimeOptions) (_result *MealApplyQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ThirdPartApplyId) {
		query["third_part_apply_id"] = request.ThirdPartApplyId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MealApplyQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/meal"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MealApplyQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询因公用餐记账数据
//
// @param request - MealBillSettlementQueryRequest
//
// @param headers - MealBillSettlementQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MealBillSettlementQueryResponse
func (client *Client) MealBillSettlementQueryWithContext(ctx context.Context, request *MealBillSettlementQueryRequest, headers *MealBillSettlementQueryHeaders, runtime *dara.RuntimeOptions) (_result *MealBillSettlementQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillBatch) {
		query["bill_batch"] = request.BillBatch
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PeriodEnd) {
		query["period_end"] = request.PeriodEnd
	}

	if !dara.IsNil(request.PeriodStart) {
		query["period_start"] = request.PeriodStart
	}

	if !dara.IsNil(request.ScrollId) {
		query["scroll_id"] = request.ScrollId
	}

	if !dara.IsNil(request.ScrollMod) {
		query["scroll_mod"] = request.ScrollMod
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MealBillSettlementQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/meal/v1/bill-settlement"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MealBillSettlementQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用餐订单详情
//
// @param request - MealOrderDetailQueryRequest
//
// @param headers - MealOrderDetailQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MealOrderDetailQueryResponse
func (client *Client) MealOrderDetailQueryWithContext(ctx context.Context, orderId *string, request *MealOrderDetailQueryRequest, headers *MealOrderDetailQueryHeaders, runtime *dara.RuntimeOptions) (_result *MealOrderDetailQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MealOrderDetailQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/meal/v1/orders/" + dara.PercentEncode(dara.StringValue(orderId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MealOrderDetailQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用餐订单列表
//
// @param request - MealOrderListQueryRequest
//
// @param headers - MealOrderListQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MealOrderListQueryResponse
func (client *Client) MealOrderListQueryWithContext(ctx context.Context, request *MealOrderListQueryRequest, headers *MealOrderListQueryHeaders, runtime *dara.RuntimeOptions) (_result *MealOrderListQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MealOrderListQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/meal/v1/orders"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MealOrderListQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 月账单确认
//
// @param request - MonthBillConfirmRequest
//
// @param headers - MonthBillConfirmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MonthBillConfirmResponse
func (client *Client) MonthBillConfirmWithContext(ctx context.Context, request *MonthBillConfirmRequest, headers *MonthBillConfirmHeaders, runtime *dara.RuntimeOptions) (_result *MonthBillConfirmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MailBillDate) {
		body["mail_bill_date"] = request.MailBillDate
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MonthBillConfirm"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bill/v1/status/action/confirm"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MonthBillConfirmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业月账单
//
// @param request - MonthBillGetRequest
//
// @param headers - MonthBillGetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MonthBillGetResponse
func (client *Client) MonthBillGetWithContext(ctx context.Context, request *MonthBillGetRequest, headers *MonthBillGetHeaders, runtime *dara.RuntimeOptions) (_result *MonthBillGetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillBatch) {
		query["bill_batch"] = request.BillBatch
	}

	if !dara.IsNil(request.BillMonth) {
		query["bill_month"] = request.BillMonth
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MonthBillGet"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/v1/month-bill"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MonthBillGetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询拆分版企业月账单
//
// @param tmpReq - MonthBillSplitGetRequest
//
// @param headers - MonthBillSplitGetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MonthBillSplitGetResponse
func (client *Client) MonthBillSplitGetWithContext(ctx context.Context, tmpReq *MonthBillSplitGetRequest, headers *MonthBillSplitGetHeaders, runtime *dara.RuntimeOptions) (_result *MonthBillSplitGetResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &MonthBillSplitGetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BillSplitKeyList) {
		request.BillSplitKeyListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BillSplitKeyList, dara.String("bill_split_key_list"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BillBatch) {
		query["bill_batch"] = request.BillBatch
	}

	if !dara.IsNil(request.BillMonth) {
		query["bill_month"] = request.BillMonth
	}

	if !dara.IsNil(request.BillSplitKeyListShrink) {
		query["bill_split_key_list"] = request.BillSplitKeyListShrink
	}

	if !dara.IsNil(request.BillSplitMode) {
		query["bill_split_mode"] = request.BillSplitMode
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MonthBillSplitGet"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/v1/month-bill-split"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MonthBillSplitGetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询订单退款明细
//
// @param request - OrderRefundDetailQueryRequest
//
// @param headers - OrderRefundDetailQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OrderRefundDetailQueryResponse
func (client *Client) OrderRefundDetailQueryWithContext(ctx context.Context, request *OrderRefundDetailQueryRequest, headers *OrderRefundDetailQueryHeaders, runtime *dara.RuntimeOptions) (_result *OrderRefundDetailQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CooperatorOrderId) {
		body["cooperator_order_id"] = request.CooperatorOrderId
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OrderRefundDetailQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/coop-hotel/v1/refund/action/detail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OrderRefundDetailQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加项目
//
// @param request - ProjectAddRequest
//
// @param headers - ProjectAddHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProjectAddResponse
func (client *Client) ProjectAddWithContext(ctx context.Context, request *ProjectAddRequest, headers *ProjectAddHeaders, runtime *dara.RuntimeOptions) (_result *ProjectAddResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ProjectName) {
		body["project_name"] = request.ProjectName
	}

	if !dara.IsNil(request.ThirdPartCostCenterId) {
		body["third_part_cost_center_id"] = request.ThirdPartCostCenterId
	}

	if !dara.IsNil(request.ThirdPartId) {
		body["third_part_id"] = request.ThirdPartId
	}

	if !dara.IsNil(request.ThirdPartInvoiceId) {
		body["third_part_invoice_id"] = request.ThirdPartInvoiceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ProjectAdd"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cost/v1/project"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ProjectAddResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除项目
//
// @param request - ProjectDeleteRequest
//
// @param headers - ProjectDeleteHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProjectDeleteResponse
func (client *Client) ProjectDeleteWithContext(ctx context.Context, request *ProjectDeleteRequest, headers *ProjectDeleteHeaders, runtime *dara.RuntimeOptions) (_result *ProjectDeleteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ThirdPartId) {
		query["third_part_id"] = request.ThirdPartId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ProjectDelete"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cost/v1/project"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ProjectDeleteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更项目
//
// @param request - ProjectModifyRequest
//
// @param headers - ProjectModifyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProjectModifyResponse
func (client *Client) ProjectModifyWithContext(ctx context.Context, request *ProjectModifyRequest, headers *ProjectModifyHeaders, runtime *dara.RuntimeOptions) (_result *ProjectModifyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ProjectName) {
		body["project_name"] = request.ProjectName
	}

	if !dara.IsNil(request.ThirdPartCostCenterId) {
		body["third_part_cost_center_id"] = request.ThirdPartCostCenterId
	}

	if !dara.IsNil(request.ThirdPartId) {
		body["third_part_id"] = request.ThirdPartId
	}

	if !dara.IsNil(request.ThirdPartInvoiceId) {
		body["third_part_invoice_id"] = request.ThirdPartInvoiceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ProjectModify"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cost/v1/project"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ProjectModifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业信息详情
//
// @param request - QueryCorpDetailInfoRequest
//
// @param headers - QueryCorpDetailInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCorpDetailInfoResponse
func (client *Client) QueryCorpDetailInfoWithContext(ctx context.Context, request *QueryCorpDetailInfoRequest, headers *QueryCorpDetailInfoHeaders, runtime *dara.RuntimeOptions) (_result *QueryCorpDetailInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["account_id"] = request.AccountId
	}

	if !dara.IsNil(request.TargetCorpId) {
		query["target_corp_id"] = request.TargetCorpId
	}

	if !dara.IsNil(request.TargetThirdCorpId) {
		query["target_third_corp_id"] = request.TargetThirdCorpId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCorpDetailInfo"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/corps/v1/corps/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCorpDetailInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个员工信息
//
// @param request - QueryEmployeeDetailRequest
//
// @param headers - QueryEmployeeDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEmployeeDetailResponse
func (client *Client) QueryEmployeeDetailWithContext(ctx context.Context, request *QueryEmployeeDetailRequest, headers *QueryEmployeeDetailHeaders, runtime *dara.RuntimeOptions) (_result *QueryEmployeeDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OutEmployeeId) {
		query["out_employee_id"] = request.OutEmployeeId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEmployeeDetail"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/user/v1/employeeDetail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEmployeeDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询子企业列表
//
// @param request - QueryGroupCorpListRequest
//
// @param headers - QueryGroupCorpListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGroupCorpListResponse
func (client *Client) QueryGroupCorpListWithContext(ctx context.Context, request *QueryGroupCorpListRequest, headers *QueryGroupCorpListHeaders, runtime *dara.RuntimeOptions) (_result *QueryGroupCorpListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryGroupCorpList"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/sub_corps/v1/corps/action/corpList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryGroupCorpListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 报销单查询
//
// @param request - QueryReimbursementOrderRequest
//
// @param headers - QueryReimbursementOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReimbursementOrderResponse
func (client *Client) QueryReimbursementOrderWithContext(ctx context.Context, request *QueryReimbursementOrderRequest, headers *QueryReimbursementOrderHeaders, runtime *dara.RuntimeOptions) (_result *QueryReimbursementOrderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReimbOrderNo) {
		query["reimb_order_no"] = request.ReimbOrderNo
	}

	if !dara.IsNil(request.SubCorpId) {
		query["sub_corp_id"] = request.SubCorpId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryReimbursementOrder"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/reimbursement/v1/order"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryReimbursementOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 单个人员同步
//
// @param tmpReq - SyncSingleUserRequest
//
// @param headers - SyncSingleUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncSingleUserResponse
func (client *Client) SyncSingleUserWithContext(ctx context.Context, tmpReq *SyncSingleUserRequest, headers *SyncSingleUserHeaders, runtime *dara.RuntimeOptions) (_result *SyncSingleUserResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SyncSingleUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ThirdDepartIdList) {
		request.ThirdDepartIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ThirdDepartIdList, dara.String("third_depart_id_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Email) {
		body["email"] = request.Email
	}

	if !dara.IsNil(request.JobNo) {
		body["job_no"] = request.JobNo
	}

	if !dara.IsNil(request.LeaveStatus) {
		body["leave_status"] = request.LeaveStatus
	}

	if !dara.IsNil(request.ManagerUserId) {
		body["manager_user_id"] = request.ManagerUserId
	}

	if !dara.IsNil(request.Phone) {
		body["phone"] = request.Phone
	}

	if !dara.IsNil(request.Position) {
		body["position"] = request.Position
	}

	if !dara.IsNil(request.PositionLevel) {
		body["position_level"] = request.PositionLevel
	}

	if !dara.IsNil(request.RealNameEn) {
		body["real_name_en"] = request.RealNameEn
	}

	if !dara.IsNil(request.ThirdDepartIdListShrink) {
		body["third_depart_id_list"] = request.ThirdDepartIdListShrink
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		body["user_name"] = request.UserName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncSingleUser"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/user/v1/single-user/action/sync"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncSingleUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步三方用户映射关系
//
// @param request - SyncThirdUserMappingRequest
//
// @param headers - SyncThirdUserMappingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncThirdUserMappingResponse
func (client *Client) SyncThirdUserMappingWithContext(ctx context.Context, request *SyncThirdUserMappingRequest, headers *SyncThirdUserMappingHeaders, runtime *dara.RuntimeOptions) (_result *SyncThirdUserMappingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.ThirdChannelType) {
		body["third_channel_type"] = request.ThirdChannelType
	}

	if !dara.IsNil(request.ThirdUserId) {
		body["third_user_id"] = request.ThirdUserId
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncThirdUserMapping"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/user/v1/third-users/action/mapping"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncThirdUserMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询淘宝账号信息
//
// @param headers - TBAccountInfoQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TBAccountInfoQueryResponse
func (client *Client) TBAccountInfoQueryWithContext(ctx context.Context, userId *string, headers *TBAccountInfoQueryHeaders, runtime *dara.RuntimeOptions) (_result *TBAccountInfoQueryResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("TBAccountInfoQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/account/v1/tb-accounts/" + dara.PercentEncode(dara.StringValue(userId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TBAccountInfoQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑淘宝账号
//
// @param headers - TBAccountUnbindHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TBAccountUnbindResponse
func (client *Client) TBAccountUnbindWithContext(ctx context.Context, userId *string, headers *TBAccountUnbindHeaders, runtime *dara.RuntimeOptions) (_result *TBAccountUnbindResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("TBAccountUnbind"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/account/v1/tb-accounts/" + dara.PercentEncode(dara.StringValue(userId)) + "/action/unbind"),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TBAccountUnbindResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票改签申请
//
// @param tmpReq - TicketChangingApplyRequest
//
// @param headers - TicketChangingApplyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TicketChangingApplyResponse
func (client *Client) TicketChangingApplyWithContext(ctx context.Context, tmpReq *TicketChangingApplyRequest, headers *TicketChangingApplyHeaders, runtime *dara.RuntimeOptions) (_result *TicketChangingApplyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TicketChangingApplyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ModifyFlightInfoList) {
		request.ModifyFlightInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModifyFlightInfoList, dara.String("modify_flight_info_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DisOrderId) {
		body["dis_order_id"] = request.DisOrderId
	}

	if !dara.IsNil(request.DisSubOrderId) {
		body["dis_sub_order_id"] = request.DisSubOrderId
	}

	if !dara.IsNil(request.IsVoluntary) {
		body["is_voluntary"] = request.IsVoluntary
	}

	if !dara.IsNil(request.ModifyFlightInfoListShrink) {
		body["modify_flight_info_list"] = request.ModifyFlightInfoListShrink
	}

	if !dara.IsNil(request.OtaItemId) {
		body["ota_item_id"] = request.OtaItemId
	}

	if !dara.IsNil(request.Reason) {
		body["reason"] = request.Reason
	}

	if !dara.IsNil(request.SessionId) {
		body["session_id"] = request.SessionId
	}

	if !dara.IsNil(request.WhetherRetry) {
		body["whether_retry"] = request.WhetherRetry
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TicketChangingApply"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/ticket-changing/action/apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TicketChangingApplyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票改签取消
//
// @param request - TicketChangingCancelRequest
//
// @param headers - TicketChangingCancelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TicketChangingCancelResponse
func (client *Client) TicketChangingCancelWithContext(ctx context.Context, request *TicketChangingCancelRequest, headers *TicketChangingCancelHeaders, runtime *dara.RuntimeOptions) (_result *TicketChangingCancelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisOrderId) {
		query["dis_order_id"] = request.DisOrderId
	}

	if !dara.IsNil(request.DisSubOrderId) {
		query["dis_sub_order_id"] = request.DisSubOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TicketChangingCancel"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/ticket-changing/action/cancel"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TicketChangingCancelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票改签详情
//
// @param request - TicketChangingDetailRequest
//
// @param headers - TicketChangingDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TicketChangingDetailResponse
func (client *Client) TicketChangingDetailWithContext(ctx context.Context, request *TicketChangingDetailRequest, headers *TicketChangingDetailHeaders, runtime *dara.RuntimeOptions) (_result *TicketChangingDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisOrderId) {
		query["dis_order_id"] = request.DisOrderId
	}

	if !dara.IsNil(request.DisSubOrderId) {
		query["dis_sub_order_id"] = request.DisSubOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TicketChangingDetail"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/ticket-changing/action/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TicketChangingDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票改签询价
//
// @param request - TicketChangingEnquiryRequest
//
// @param headers - TicketChangingEnquiryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TicketChangingEnquiryResponse
func (client *Client) TicketChangingEnquiryWithContext(ctx context.Context, request *TicketChangingEnquiryRequest, headers *TicketChangingEnquiryHeaders, runtime *dara.RuntimeOptions) (_result *TicketChangingEnquiryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArrCity) {
		query["arr_city"] = request.ArrCity
	}

	if !dara.IsNil(request.DepCity) {
		query["dep_city"] = request.DepCity
	}

	if !dara.IsNil(request.DisOrderId) {
		query["dis_order_id"] = request.DisOrderId
	}

	if !dara.IsNil(request.IsVoluntary) {
		query["is_voluntary"] = request.IsVoluntary
	}

	if !dara.IsNil(request.ModifyDepartDate) {
		query["modify_depart_date"] = request.ModifyDepartDate
	}

	if !dara.IsNil(request.ModifyFlightNo) {
		query["modify_flight_no"] = request.ModifyFlightNo
	}

	if !dara.IsNil(request.SessionId) {
		query["session_id"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TicketChangingEnquiry"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/ticket-changing/action/enquiry"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TicketChangingEnquiryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票改签可改签航班列表
//
// @param tmpReq - TicketChangingFlightListRequest
//
// @param headers - TicketChangingFlightListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TicketChangingFlightListResponse
func (client *Client) TicketChangingFlightListWithContext(ctx context.Context, tmpReq *TicketChangingFlightListRequest, headers *TicketChangingFlightListHeaders, runtime *dara.RuntimeOptions) (_result *TicketChangingFlightListResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TicketChangingFlightListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TravelerInfoList) {
		request.TravelerInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TravelerInfoList, dara.String("traveler_info_list"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ArrCity) {
		query["arr_city"] = request.ArrCity
	}

	if !dara.IsNil(request.DepCity) {
		query["dep_city"] = request.DepCity
	}

	if !dara.IsNil(request.DepDate) {
		query["dep_date"] = request.DepDate
	}

	if !dara.IsNil(request.DisOrderId) {
		query["dis_order_id"] = request.DisOrderId
	}

	if !dara.IsNil(request.IsVoluntary) {
		query["is_voluntary"] = request.IsVoluntary
	}

	if !dara.IsNil(request.TravelerInfoListShrink) {
		query["traveler_info_list"] = request.TravelerInfoListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TicketChangingFlightList"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/huge/dtb-flight/v1/ticket-changing-flight/action/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TicketChangingFlightListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机票改签航班支付
//
// @param tmpReq - TicketChangingPayRequest
//
// @param headers - TicketChangingPayHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TicketChangingPayResponse
func (client *Client) TicketChangingPayWithContext(ctx context.Context, tmpReq *TicketChangingPayRequest, headers *TicketChangingPayHeaders, runtime *dara.RuntimeOptions) (_result *TicketChangingPayResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TicketChangingPayShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("extra"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CorpPayPrice) {
		body["corp_pay_price"] = request.CorpPayPrice
	}

	if !dara.IsNil(request.DisOrderId) {
		body["dis_order_id"] = request.DisOrderId
	}

	if !dara.IsNil(request.DisSubOrderId) {
		body["dis_sub_order_id"] = request.DisSubOrderId
	}

	if !dara.IsNil(request.ExtraShrink) {
		body["extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.PersonalPayPrice) {
		body["personal_pay_price"] = request.PersonalPayPrice
	}

	if !dara.IsNil(request.TotalPayPrice) {
		body["total_pay_price"] = request.TotalPayPrice
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TicketChangingPay"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dtb-flight/v1/ticket-changing/action/pay"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TicketChangingPayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 火车票改签申请
//
// @param tmpReq - TrainApplyChangeRequest
//
// @param headers - TrainApplyChangeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainApplyChangeResponse
func (client *Client) TrainApplyChangeWithContext(ctx context.Context, tmpReq *TrainApplyChangeRequest, headers *TrainApplyChangeHeaders, runtime *dara.RuntimeOptions) (_result *TrainApplyChangeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TrainApplyChangeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChangeTrainInfoS) {
		request.ChangeTrainInfoSShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChangeTrainInfoS, dara.String("change_train_info_s"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeTrainInfoSShrink) {
		query["change_train_info_s"] = request.ChangeTrainInfoSShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptNoSeat) {
		body["accept_no_seat"] = request.AcceptNoSeat
	}

	if !dara.IsNil(request.ForceMatch) {
		body["force_match"] = request.ForceMatch
	}

	if !dara.IsNil(request.IsPayNow) {
		body["is_pay_now"] = request.IsPayNow
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutChangeApplyId) {
		body["out_change_apply_id"] = request.OutChangeApplyId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainApplyChange"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/change/apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainApplyChangeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 火车票退票申请
//
// @param tmpReq - TrainApplyRefundRequest
//
// @param headers - TrainApplyRefundHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainApplyRefundResponse
func (client *Client) TrainApplyRefundWithContext(ctx context.Context, tmpReq *TrainApplyRefundRequest, headers *TrainApplyRefundHeaders, runtime *dara.RuntimeOptions) (_result *TrainApplyRefundResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TrainApplyRefundShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RefundTrainInfos) {
		request.RefundTrainInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RefundTrainInfos, dara.String("refund_train_infos"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.OutRefundId) {
		body["out_refund_id"] = request.OutRefundId
	}

	if !dara.IsNil(request.RefundTrainInfosShrink) {
		body["refund_train_infos"] = request.RefundTrainInfosShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainApplyRefund"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/refund/apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainApplyRefundResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询火车票记账数据
//
// @param request - TrainBillSettlementQueryRequest
//
// @param headers - TrainBillSettlementQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainBillSettlementQueryResponse
func (client *Client) TrainBillSettlementQueryWithContext(ctx context.Context, request *TrainBillSettlementQueryRequest, headers *TrainBillSettlementQueryHeaders, runtime *dara.RuntimeOptions) (_result *TrainBillSettlementQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillBatch) {
		query["bill_batch"] = request.BillBatch
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PeriodEnd) {
		query["period_end"] = request.PeriodEnd
	}

	if !dara.IsNil(request.PeriodStart) {
		query["period_start"] = request.PeriodStart
	}

	if !dara.IsNil(request.ScrollId) {
		query["scroll_id"] = request.ScrollId
	}

	if !dara.IsNil(request.ScrollMod) {
		query["scroll_mod"] = request.ScrollMod
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainBillSettlementQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/bill-settlement"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainBillSettlementQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询火车超标审批详情
//
// @param request - TrainExceedApplyQueryRequest
//
// @param headers - TrainExceedApplyQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainExceedApplyQueryResponse
func (client *Client) TrainExceedApplyQueryWithContext(ctx context.Context, request *TrainExceedApplyQueryRequest, headers *TrainExceedApplyQueryHeaders, runtime *dara.RuntimeOptions) (_result *TrainExceedApplyQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyId) {
		query["apply_id"] = request.ApplyId
	}

	if !dara.IsNil(request.BusinessInstanceId) {
		query["business_instance_id"] = request.BusinessInstanceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainExceedApplyQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/train-exceed"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainExceedApplyQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 火车票改签费用预估
//
// @param tmpReq - TrainFeeCalculateChangeRequest
//
// @param headers - TrainFeeCalculateChangeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainFeeCalculateChangeResponse
func (client *Client) TrainFeeCalculateChangeWithContext(ctx context.Context, tmpReq *TrainFeeCalculateChangeRequest, headers *TrainFeeCalculateChangeHeaders, runtime *dara.RuntimeOptions) (_result *TrainFeeCalculateChangeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TrainFeeCalculateChangeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChangeTrainDetails) {
		request.ChangeTrainDetailsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChangeTrainDetails, dara.String("change_train_details"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChangeTrainDetailsShrink) {
		body["change_train_details"] = request.ChangeTrainDetailsShrink
	}

	if !dara.IsNil(request.DistributeOrderId) {
		body["distribute_order_id"] = request.DistributeOrderId
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainFeeCalculateChange"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/change/fee"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainFeeCalculateChangeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 火车票退票费用预估
//
// @param tmpReq - TrainFeeCalculateRefundRequest
//
// @param headers - TrainFeeCalculateRefundHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainFeeCalculateRefundResponse
func (client *Client) TrainFeeCalculateRefundWithContext(ctx context.Context, tmpReq *TrainFeeCalculateRefundRequest, headers *TrainFeeCalculateRefundHeaders, runtime *dara.RuntimeOptions) (_result *TrainFeeCalculateRefundResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TrainFeeCalculateRefundShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RefundTrainInfos) {
		request.RefundTrainInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RefundTrainInfos, dara.String("refund_train_infos"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DistributeOrderId) {
		body["distribute_order_id"] = request.DistributeOrderId
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.RefundTrainInfosShrink) {
		body["refund_train_infos"] = request.RefundTrainInfosShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainFeeCalculateRefund"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/refund/fee"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainFeeCalculateRefundResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 火车票车次详情查询
//
// @param request - TrainNoInfoSearchRequest
//
// @param headers - TrainNoInfoSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainNoInfoSearchResponse
func (client *Client) TrainNoInfoSearchWithContext(ctx context.Context, request *TrainNoInfoSearchRequest, headers *TrainNoInfoSearchHeaders, runtime *dara.RuntimeOptions) (_result *TrainNoInfoSearchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ArrLocation) {
		body["arr_location"] = request.ArrLocation
	}

	if !dara.IsNil(request.DepDate) {
		body["dep_date"] = request.DepDate
	}

	if !dara.IsNil(request.DepLocation) {
		body["dep_location"] = request.DepLocation
	}

	if !dara.IsNil(request.LineKey) {
		body["line_key"] = request.LineKey
	}

	if !dara.IsNil(request.MiddleDate) {
		body["middle_date"] = request.MiddleDate
	}

	if !dara.IsNil(request.MiddleStation) {
		body["middle_station"] = request.MiddleStation
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.TrainNo) {
		body["train_no"] = request.TrainNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainNoInfoSearch"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/search/info"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainNoInfoSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 火车票车次列表查询
//
// @param tmpReq - TrainNoListSearchRequest
//
// @param headers - TrainNoListSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainNoListSearchResponse
func (client *Client) TrainNoListSearchWithContext(ctx context.Context, tmpReq *TrainNoListSearchRequest, headers *TrainNoListSearchHeaders, runtime *dara.RuntimeOptions) (_result *TrainNoListSearchResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TrainNoListSearchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Option) {
		request.OptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Option, dara.String("option"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ArrLocation) {
		body["arr_location"] = request.ArrLocation
	}

	if !dara.IsNil(request.DepDate) {
		body["dep_date"] = request.DepDate
	}

	if !dara.IsNil(request.DepLocation) {
		body["dep_location"] = request.DepLocation
	}

	if !dara.IsNil(request.OptionShrink) {
		body["option"] = request.OptionShrink
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainNoListSearch"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/search/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainNoListSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 火车票订单取消
//
// @param request - TrainOrderCancelRequest
//
// @param headers - TrainOrderCancelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainOrderCancelResponse
func (client *Client) TrainOrderCancelWithContext(ctx context.Context, request *TrainOrderCancelRequest, headers *TrainOrderCancelHeaders, runtime *dara.RuntimeOptions) (_result *TrainOrderCancelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChangeOrderId) {
		body["change_order_id"] = request.ChangeOrderId
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutChangeOrderId) {
		body["out_change_order_id"] = request.OutChangeOrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainOrderCancel"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/order/cancel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainOrderCancelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 火车票改签确认
//
// @param request - TrainOrderChangeConfirmRequest
//
// @param headers - TrainOrderChangeConfirmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainOrderChangeConfirmResponse
func (client *Client) TrainOrderChangeConfirmWithContext(ctx context.Context, request *TrainOrderChangeConfirmRequest, headers *TrainOrderChangeConfirmHeaders, runtime *dara.RuntimeOptions) (_result *TrainOrderChangeConfirmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChangeApplyId) {
		body["change_apply_id"] = request.ChangeApplyId
	}

	if !dara.IsNil(request.ChangeSettleAmount) {
		body["change_settle_amount"] = request.ChangeSettleAmount
	}

	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutChangeApplyId) {
		body["out_change_apply_id"] = request.OutChangeApplyId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainOrderChangeConfirm"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/change/confirm"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainOrderChangeConfirmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 火车票正向预订
//
// @param tmpReq - TrainOrderCreateRequest
//
// @param headers - TrainOrderCreateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainOrderCreateResponse
func (client *Client) TrainOrderCreateWithContext(ctx context.Context, tmpReq *TrainOrderCreateRequest, headers *TrainOrderCreateHeaders, runtime *dara.RuntimeOptions) (_result *TrainOrderCreateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TrainOrderCreateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BookTrainInfos) {
		request.BookTrainInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BookTrainInfos, dara.String("book_train_infos"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.BusinessInfo) {
		request.BusinessInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BusinessInfo, dara.String("business_info"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ContactInfo) {
		request.ContactInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContactInfo, dara.String("contact_info"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PassengerOpenInfoS) {
		request.PassengerOpenInfoSShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PassengerOpenInfoS, dara.String("passenger_open_info_s"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptNoSeat) {
		body["accept_no_seat"] = request.AcceptNoSeat
	}

	if !dara.IsNil(request.BookTrainInfosShrink) {
		body["book_train_infos"] = request.BookTrainInfosShrink
	}

	if !dara.IsNil(request.BtripUserId) {
		body["btrip_user_id"] = request.BtripUserId
	}

	if !dara.IsNil(request.BtripUserName) {
		body["btrip_user_name"] = request.BtripUserName
	}

	if !dara.IsNil(request.BusinessInfoShrink) {
		body["business_info"] = request.BusinessInfoShrink
	}

	if !dara.IsNil(request.ContactInfoShrink) {
		body["contact_info"] = request.ContactInfoShrink
	}

	if !dara.IsNil(request.ForceMatch) {
		body["force_match"] = request.ForceMatch
	}

	if !dara.IsNil(request.IsPayNow) {
		body["is_pay_now"] = request.IsPayNow
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.PassengerOpenInfoSShrink) {
		body["passenger_open_info_s"] = request.PassengerOpenInfoSShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainOrderCreate"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/order/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainOrderCreateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 火车票订单详情
//
// @param request - TrainOrderDetailQueryRequest
//
// @param headers - TrainOrderDetailQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainOrderDetailQueryResponse
func (client *Client) TrainOrderDetailQueryWithContext(ctx context.Context, request *TrainOrderDetailQueryRequest, headers *TrainOrderDetailQueryHeaders, runtime *dara.RuntimeOptions) (_result *TrainOrderDetailQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainOrderDetailQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/order/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainOrderDetailQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询火车票订单列表
//
// @param request - TrainOrderListQueryRequest
//
// @param headers - TrainOrderListQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainOrderListQueryResponse
func (client *Client) TrainOrderListQueryWithContext(ctx context.Context, request *TrainOrderListQueryRequest, headers *TrainOrderListQueryHeaders, runtime *dara.RuntimeOptions) (_result *TrainOrderListQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllApply) {
		query["all_apply"] = request.AllApply
	}

	if !dara.IsNil(request.ApplyId) {
		query["apply_id"] = request.ApplyId
	}

	if !dara.IsNil(request.DepartId) {
		query["depart_id"] = request.DepartId
	}

	if !dara.IsNil(request.EndTime) {
		query["end_time"] = request.EndTime
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["start_time"] = request.StartTime
	}

	if !dara.IsNil(request.ThirdpartApplyId) {
		query["thirdpart_apply_id"] = request.ThirdpartApplyId
	}

	if !dara.IsNil(request.UpdateEndTime) {
		query["update_end_time"] = request.UpdateEndTime
	}

	if !dara.IsNil(request.UpdateStartTime) {
		query["update_start_time"] = request.UpdateStartTime
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainOrderListQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/order-list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainOrderListQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 火车票订单支付
//
// @param request - TrainOrderPayRequest
//
// @param headers - TrainOrderPayHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainOrderPayResponse
func (client *Client) TrainOrderPayWithContext(ctx context.Context, request *TrainOrderPayRequest, headers *TrainOrderPayHeaders, runtime *dara.RuntimeOptions) (_result *TrainOrderPayResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		body["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.OutOrderId) {
		body["out_order_id"] = request.OutOrderId
	}

	if !dara.IsNil(request.PayAmount) {
		body["pay_amount"] = request.PayAmount
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainOrderPay"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/order/pay"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainOrderPayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询火车票订单详情（含票信息）
//
// @param request - TrainOrderQueryRequest
//
// @param headers - TrainOrderQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainOrderQueryResponse
func (client *Client) TrainOrderQueryWithContext(ctx context.Context, request *TrainOrderQueryRequest, headers *TrainOrderQueryHeaders, runtime *dara.RuntimeOptions) (_result *TrainOrderQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainOrderQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/order"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainOrderQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 火车票订单查询V2
//
// @param request - TrainOrderQueryV2Request
//
// @param headers - TrainOrderQueryV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainOrderQueryV2Response
func (client *Client) TrainOrderQueryV2WithContext(ctx context.Context, request *TrainOrderQueryV2Request, headers *TrainOrderQueryV2Headers, runtime *dara.RuntimeOptions) (_result *TrainOrderQueryV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainOrderQueryV2"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v2/order"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainOrderQueryV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询火车站数据
//
// @param request - TrainStationSearchRequest
//
// @param headers - TrainStationSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainStationSearchResponse
func (client *Client) TrainStationSearchWithContext(ctx context.Context, request *TrainStationSearchRequest, headers *TrainStationSearchHeaders, runtime *dara.RuntimeOptions) (_result *TrainStationSearchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainStationSearch"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/city/v1/train"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainStationSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 火车票经停站查询
//
// @param request - TrainStopoverSearchRequest
//
// @param headers - TrainStopoverSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainStopoverSearchResponse
func (client *Client) TrainStopoverSearchWithContext(ctx context.Context, request *TrainStopoverSearchRequest, headers *TrainStopoverSearchHeaders, runtime *dara.RuntimeOptions) (_result *TrainStopoverSearchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ArrStation) {
		body["arr_station"] = request.ArrStation
	}

	if !dara.IsNil(request.DepStation) {
		body["dep_station"] = request.DepStation
	}

	if !dara.IsNil(request.TrainDate) {
		body["train_date"] = request.TrainDate
	}

	if !dara.IsNil(request.TrainNo) {
		body["train_no"] = request.TrainNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainStopoverSearch"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/train/v1/search/stopover"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainStopoverSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询火车票凭证扫描件
//
// @param request - TrainTicketScanQueryRequest
//
// @param headers - TrainTicketScanQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrainTicketScanQueryResponse
func (client *Client) TrainTicketScanQueryWithContext(ctx context.Context, request *TrainTicketScanQueryRequest, headers *TrainTicketScanQueryHeaders, runtime *dara.RuntimeOptions) (_result *TrainTicketScanQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillDate) {
		query["bill_date"] = request.BillDate
	}

	if !dara.IsNil(request.BillId) {
		query["bill_id"] = request.BillId
	}

	if !dara.IsNil(request.InvoiceSubTaskId) {
		query["invoice_sub_task_id"] = request.InvoiceSubTaskId
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.SerialNumber) {
		query["serial_number"] = request.SerialNumber
	}

	if !dara.IsNil(request.TicketNo) {
		query["ticket_no"] = request.TicketNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrainTicketScanQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/scan/v1/train-ticket"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TrainTicketScanQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询差标列表
//
// @param request - TravelStandardListQueryRequest
//
// @param headers - TravelStandardListQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TravelStandardListQueryResponse
func (client *Client) TravelStandardListQueryWithContext(ctx context.Context, request *TravelStandardListQueryRequest, headers *TravelStandardListQueryHeaders, runtime *dara.RuntimeOptions) (_result *TravelStandardListQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FromGroup) {
		query["from_group"] = request.FromGroup
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.RuleName) {
		query["rule_name"] = request.RuleName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TravelStandardListQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/travel-manage/v1/standards/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TravelStandardListQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询差标详情
//
// @param tmpReq - TravelStandardQueryRequest
//
// @param headers - TravelStandardQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TravelStandardQueryResponse
func (client *Client) TravelStandardQueryWithContext(ctx context.Context, tmpReq *TravelStandardQueryRequest, headers *TravelStandardQueryHeaders, runtime *dara.RuntimeOptions) (_result *TravelStandardQueryResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TravelStandardQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ServiceTypeList) {
		request.ServiceTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServiceTypeList, dara.String("service_type_list"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FromGroup) {
		query["from_group"] = request.FromGroup
	}

	if !dara.IsNil(request.RuleCode) {
		query["rule_code"] = request.RuleCode
	}

	if !dara.IsNil(request.ServiceTypeListShrink) {
		query["service_type_list"] = request.ServiceTypeListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TravelStandardQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/travel-manage/v1/standards/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TravelStandardQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增差旅标准关联人员实体
//
// @param tmpReq - TravelStandardRelateAddRequest
//
// @param headers - TravelStandardRelateAddHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TravelStandardRelateAddResponse
func (client *Client) TravelStandardRelateAddWithContext(ctx context.Context, tmpReq *TravelStandardRelateAddRequest, headers *TravelStandardRelateAddHeaders, runtime *dara.RuntimeOptions) (_result *TravelStandardRelateAddResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TravelStandardRelateAddShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddList) {
		request.AddListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddList, dara.String("add_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddListShrink) {
		body["add_list"] = request.AddListShrink
	}

	if !dara.IsNil(request.FromGroup) {
		body["from_group"] = request.FromGroup
	}

	if !dara.IsNil(request.RuleId) {
		body["rule_id"] = request.RuleId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TravelStandardRelateAdd"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/travel-manage/v1/standards/add-relate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TravelStandardRelateAddResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除差旅标准关联人员实体
//
// @param tmpReq - TravelStandardRelateDeleteRequest
//
// @param headers - TravelStandardRelateDeleteHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TravelStandardRelateDeleteResponse
func (client *Client) TravelStandardRelateDeleteWithContext(ctx context.Context, tmpReq *TravelStandardRelateDeleteRequest, headers *TravelStandardRelateDeleteHeaders, runtime *dara.RuntimeOptions) (_result *TravelStandardRelateDeleteResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TravelStandardRelateDeleteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RemoveList) {
		request.RemoveListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RemoveList, dara.String("remove_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FromGroup) {
		body["from_group"] = request.FromGroup
	}

	if !dara.IsNil(request.RemoveListShrink) {
		body["remove_list"] = request.RemoveListShrink
	}

	if !dara.IsNil(request.RuleId) {
		body["rule_id"] = request.RuleId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TravelStandardRelateDelete"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/travel-manage/v1/standards/delete-relate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TravelStandardRelateDeleteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询差旅标准关联人员实体
//
// @param request - TravelStandardRelateQueryRequest
//
// @param headers - TravelStandardRelateQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TravelStandardRelateQueryResponse
func (client *Client) TravelStandardRelateQueryWithContext(ctx context.Context, request *TravelStandardRelateQueryRequest, headers *TravelStandardRelateQueryHeaders, runtime *dara.RuntimeOptions) (_result *TravelStandardRelateQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FromGroup) {
		query["from_group"] = request.FromGroup
	}

	if !dara.IsNil(request.RuleId) {
		query["rule_id"] = request.RuleId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TravelStandardRelateQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/travel-manage/v1/standards/query-relate"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TravelStandardRelateQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新差旅标准绑定员工类型
//
// @param request - TravelStandardScopeSaveRequest
//
// @param headers - TravelStandardScopeSaveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TravelStandardScopeSaveResponse
func (client *Client) TravelStandardScopeSaveWithContext(ctx context.Context, request *TravelStandardScopeSaveRequest, headers *TravelStandardScopeSaveHeaders, runtime *dara.RuntimeOptions) (_result *TravelStandardScopeSaveResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FromGroup) {
		query["from_group"] = request.FromGroup
	}

	if !dara.IsNil(request.RuleId) {
		query["rule_id"] = request.RuleId
	}

	if !dara.IsNil(request.Scope) {
		query["scope"] = request.Scope
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TravelStandardScopeSave"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/travel-manage/v1/standards/save-scope"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TravelStandardScopeSaveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询业务流程
//
// @param request - TripBusinessInstanceQueryRequest
//
// @param headers - TripBusinessInstanceQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TripBusinessInstanceQueryResponse
func (client *Client) TripBusinessInstanceQueryWithContext(ctx context.Context, request *TripBusinessInstanceQueryRequest, headers *TripBusinessInstanceQueryHeaders, runtime *dara.RuntimeOptions) (_result *TripBusinessInstanceQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessInstanceId) {
		query["business_instance_id"] = request.BusinessInstanceId
	}

	if !dara.IsNil(request.ThirdBusinessId) {
		query["third_business_id"] = request.ThirdBusinessId
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		query["user_name"] = request.UserName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TripBusinessInstanceQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/business"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TripBusinessInstanceQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询抄送信息
//
// @param request - TripCCInfoQueryRequest
//
// @param headers - TripCCInfoQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TripCCInfoQueryResponse
func (client *Client) TripCCInfoQueryWithContext(ctx context.Context, request *TripCCInfoQueryRequest, headers *TripCCInfoQueryHeaders, runtime *dara.RuntimeOptions) (_result *TripCCInfoQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessInstanceId) {
		query["business_instance_id"] = request.BusinessInstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["node_id"] = request.NodeId
	}

	if !dara.IsNil(request.ThirdBusinessId) {
		query["third_business_id"] = request.ThirdBusinessId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TripCCInfoQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/cc"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TripCCInfoQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审批任务列表
//
// @param request - TripTaskQueryRequest
//
// @param headers - TripTaskQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TripTaskQueryResponse
func (client *Client) TripTaskQueryWithContext(ctx context.Context, request *TripTaskQueryRequest, headers *TripTaskQueryHeaders, runtime *dara.RuntimeOptions) (_result *TripTaskQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessInstanceId) {
		query["business_instance_id"] = request.BusinessInstanceId
	}

	if !dara.IsNil(request.ThirdBusinessId) {
		query["third_business_id"] = request.ThirdBusinessId
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		query["user_name"] = request.UserName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TripTaskQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apply/v1/tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TripTaskQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新企业自定义角色
//
// @param request - UpdateCustomRoleRequest
//
// @param headers - UpdateCustomRoleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomRoleResponse
func (client *Client) UpdateCustomRoleWithContext(ctx context.Context, request *UpdateCustomRoleRequest, headers *UpdateCustomRoleHeaders, runtime *dara.RuntimeOptions) (_result *UpdateCustomRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RoleId) {
		body["role_id"] = request.RoleId
	}

	if !dara.IsNil(request.RoleName) {
		body["role_name"] = request.RoleName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomRole"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/role/v1/customRoles/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改企业部门
//
// @param tmpReq - UpdateDepartmentRequest
//
// @param headers - UpdateDepartmentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDepartmentResponse
func (client *Client) UpdateDepartmentWithContext(ctx context.Context, tmpReq *UpdateDepartmentRequest, headers *UpdateDepartmentHeaders, runtime *dara.RuntimeOptions) (_result *UpdateDepartmentResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDepartmentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ManagerEmployeeIdList) {
		request.ManagerEmployeeIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ManagerEmployeeIdList, dara.String("manager_employee_id_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeptName) {
		body["dept_name"] = request.DeptName
	}

	if !dara.IsNil(request.ManagerEmployeeIdListShrink) {
		body["manager_employee_id_list"] = request.ManagerEmployeeIdListShrink
	}

	if !dara.IsNil(request.OutDeptId) {
		body["out_dept_id"] = request.OutDeptId
	}

	if !dara.IsNil(request.OutDeptPid) {
		body["out_dept_pid"] = request.OutDeptPid
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDepartment"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/department/v2/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDepartmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新员工信息
//
// @param tmpReq - UpdateEmployeeRequest
//
// @param headers - UpdateEmployeeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEmployeeResponse
func (client *Client) UpdateEmployeeWithContext(ctx context.Context, tmpReq *UpdateEmployeeRequest, headers *UpdateEmployeeHeaders, runtime *dara.RuntimeOptions) (_result *UpdateEmployeeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateEmployeeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BaseCityCodeList) {
		request.BaseCityCodeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BaseCityCodeList, dara.String("base_city_code_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.BaseLocationList) {
		request.BaseLocationListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BaseLocationList, dara.String("base_location_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CertList) {
		request.CertListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CertList, dara.String("cert_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomRoleCodeList) {
		request.CustomRoleCodeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomRoleCodeList, dara.String("custom_role_code_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OutDeptIdList) {
		request.OutDeptIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutDeptIdList, dara.String("out_dept_id_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountEmail) {
		body["account_email"] = request.AccountEmail
	}

	if !dara.IsNil(request.AccountPhone) {
		body["account_phone"] = request.AccountPhone
	}

	if !dara.IsNil(request.Attribute) {
		body["attribute"] = request.Attribute
	}

	if !dara.IsNil(request.Avatar) {
		body["avatar"] = request.Avatar
	}

	if !dara.IsNil(request.BaseCityCodeListShrink) {
		body["base_city_code_list"] = request.BaseCityCodeListShrink
	}

	if !dara.IsNil(request.BaseLocationListShrink) {
		body["base_location_list"] = request.BaseLocationListShrink
	}

	if !dara.IsNil(request.Birthday) {
		body["birthday"] = request.Birthday
	}

	if !dara.IsNil(request.CertListShrink) {
		body["cert_list"] = request.CertListShrink
	}

	if !dara.IsNil(request.CustomRoleCodeListShrink) {
		body["custom_role_code_list"] = request.CustomRoleCodeListShrink
	}

	if !dara.IsNil(request.Email) {
		body["email"] = request.Email
	}

	if !dara.IsNil(request.Gender) {
		body["gender"] = request.Gender
	}

	if !dara.IsNil(request.IsAdmin) {
		body["is_admin"] = request.IsAdmin
	}

	if !dara.IsNil(request.IsBoss) {
		body["is_boss"] = request.IsBoss
	}

	if !dara.IsNil(request.IsDeptLeader) {
		body["is_dept_leader"] = request.IsDeptLeader
	}

	if !dara.IsNil(request.JobNo) {
		body["job_no"] = request.JobNo
	}

	if !dara.IsNil(request.ManagerUserId) {
		body["manager_user_id"] = request.ManagerUserId
	}

	if !dara.IsNil(request.OutDeptIdListShrink) {
		body["out_dept_id_list"] = request.OutDeptIdListShrink
	}

	if !dara.IsNil(request.Phone) {
		body["phone"] = request.Phone
	}

	if !dara.IsNil(request.PositionLevel) {
		body["position_level"] = request.PositionLevel
	}

	if !dara.IsNil(request.RealName) {
		body["real_name"] = request.RealName
	}

	if !dara.IsNil(request.RealNameEn) {
		body["real_name_en"] = request.RealNameEn
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	if !dara.IsNil(request.UserNick) {
		body["user_nick"] = request.UserNick
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEmployee"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/employee/v2/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEmployeeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新员工在职状态
//
// @param request - UpdateEmployeeLeaveStatusRequest
//
// @param headers - UpdateEmployeeLeaveStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEmployeeLeaveStatusResponse
func (client *Client) UpdateEmployeeLeaveStatusWithContext(ctx context.Context, request *UpdateEmployeeLeaveStatusRequest, headers *UpdateEmployeeLeaveStatusHeaders, runtime *dara.RuntimeOptions) (_result *UpdateEmployeeLeaveStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IsLeave) {
		body["is_leave"] = request.IsLeave
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEmployeeLeaveStatus"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/employee/v2/updateLeaveStatus"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEmployeeLeaveStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人员查询
//
// @param request - UserQueryRequest
//
// @param headers - UserQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UserQueryResponse
func (client *Client) UserQueryWithContext(ctx context.Context, request *UserQueryRequest, headers *UserQueryHeaders, runtime *dara.RuntimeOptions) (_result *UserQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModifiedTimeGreaterOrEqualThan) {
		query["modified_time_greater_or_equal_than"] = request.ModifiedTimeGreaterOrEqualThan
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PageToken) {
		query["page_token"] = request.PageToken
	}

	if !dara.IsNil(request.ThirdPartJobNo) {
		query["third_part_job_no"] = request.ThirdPartJobNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UserQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/user/v1/user"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UserQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询增值服务记账数据
//
// @param request - VasBillSettlementQueryRequest
//
// @param headers - VasBillSettlementQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VasBillSettlementQueryResponse
func (client *Client) VasBillSettlementQueryWithContext(ctx context.Context, request *VasBillSettlementQueryRequest, headers *VasBillSettlementQueryHeaders, runtime *dara.RuntimeOptions) (_result *VasBillSettlementQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillBatch) {
		query["bill_batch"] = request.BillBatch
	}

	if !dara.IsNil(request.CooperatorId) {
		query["cooperator_id"] = request.CooperatorId
	}

	if !dara.IsNil(request.OrderId) {
		query["order_id"] = request.OrderId
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PeriodEnd) {
		query["period_end"] = request.PeriodEnd
	}

	if !dara.IsNil(request.PeriodStart) {
		query["period_start"] = request.PeriodStart
	}

	if !dara.IsNil(request.ScrollId) {
		query["scroll_id"] = request.ScrollId
	}

	if !dara.IsNil(request.ScrollMod) {
		query["scroll_mod"] = request.ScrollMod
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripCorpToken) {
		realHeaders["x-acs-btrip-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VasBillSettlementQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/vas/v1/bill-settlement"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &VasBillSettlementQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询增值税发票扫描件
//
// @param request - VatInvoiceScanQueryRequest
//
// @param headers - VatInvoiceScanQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VatInvoiceScanQueryResponse
func (client *Client) VatInvoiceScanQueryWithContext(ctx context.Context, request *VatInvoiceScanQueryRequest, headers *VatInvoiceScanQueryHeaders, runtime *dara.RuntimeOptions) (_result *VatInvoiceScanQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillDate) {
		query["bill_date"] = request.BillDate
	}

	if !dara.IsNil(request.BillId) {
		query["bill_id"] = request.BillId
	}

	if !dara.IsNil(request.InvoiceSubTaskId) {
		query["invoice_sub_task_id"] = request.InvoiceSubTaskId
	}

	if !dara.IsNil(request.PageNo) {
		query["page_no"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VatInvoiceScanQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/scan/v1/vat-invoice"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &VatInvoiceScanQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询账期待申请的发票数据
//
// @param request - WaitApplyInvoiceTaskDetailQueryRequest
//
// @param headers - WaitApplyInvoiceTaskDetailQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WaitApplyInvoiceTaskDetailQueryResponse
func (client *Client) WaitApplyInvoiceTaskDetailQueryWithContext(ctx context.Context, request *WaitApplyInvoiceTaskDetailQueryRequest, headers *WaitApplyInvoiceTaskDetailQueryHeaders, runtime *dara.RuntimeOptions) (_result *WaitApplyInvoiceTaskDetailQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillDate) {
		query["bill_date"] = request.BillDate
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsBtripSoCorpToken) {
		realHeaders["x-acs-btrip-so-corp-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsBtripSoCorpToken)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WaitApplyInvoiceTaskDetailQuery"),
		Version:     dara.String("2022-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/invoice/v1/wait-apply-task"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &WaitApplyInvoiceTaskDetailQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
