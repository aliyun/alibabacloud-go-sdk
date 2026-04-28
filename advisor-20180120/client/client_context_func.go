// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 根据多个维度获取用户最新的巡检结果，全量返回-openApi
//
// @param request - DescribeAdvicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvicesResponse
func (client *Client) DescribeAdvicesWithContext(ctx context.Context, request *DescribeAdvicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdvicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdviceId) {
		query["AdviceId"] = request.AdviceId
	}

	if !dara.IsNil(request.CheckId) {
		query["CheckId"] = request.CheckId
	}

	if !dara.IsNil(request.CheckPlanId) {
		query["CheckPlanId"] = request.CheckPlanId
	}

	if !dara.IsNil(request.ExcludeAdviceId) {
		query["ExcludeAdviceId"] = request.ExcludeAdviceId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdvices"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdvicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeAdvicesFlat分页
//
// @param request - DescribeAdvicesFlatPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvicesFlatPageResponse
func (client *Client) DescribeAdvicesFlatPageWithContext(ctx context.Context, request *DescribeAdvicesFlatPageRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdvicesFlatPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdviceId) {
		query["AdviceId"] = request.AdviceId
	}

	if !dara.IsNil(request.CheckId) {
		query["CheckId"] = request.CheckId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdvicesFlatPage"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdvicesFlatPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeAdvices分页
//
// @param request - DescribeAdvicesPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvicesPageResponse
func (client *Client) DescribeAdvicesPageWithContext(ctx context.Context, request *DescribeAdvicesPageRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdvicesPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdviceId) {
		query["AdviceId"] = request.AdviceId
	}

	if !dara.IsNil(request.CheckId) {
		query["CheckId"] = request.CheckId
	}

	if !dara.IsNil(request.CheckPlanId) {
		query["CheckPlanId"] = request.CheckPlanId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdvicesPage"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdvicesPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 巡检
//
// @param request - DescribeAdvisorChecksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvisorChecksResponse
func (client *Client) DescribeAdvisorChecksWithContext(ctx context.Context, request *DescribeAdvisorChecksRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdvisorChecksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdvisorChecks"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdvisorChecksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 巡检项设置-分页
//
// @param tmpReq - DescribeAdvisorChecksFoPagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvisorChecksFoPagesResponse
func (client *Client) DescribeAdvisorChecksFoPagesWithContext(ctx context.Context, tmpReq *DescribeAdvisorChecksFoPagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdvisorChecksFoPagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeAdvisorChecksFoPagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckTypes) {
		request.CheckTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckTypes, dara.String("CheckTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AssumeAliyunId) {
		query["AssumeAliyunId"] = request.AssumeAliyunId
	}

	if !dara.IsNil(request.BizCategory) {
		query["BizCategory"] = request.BizCategory
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.CheckTypesShrink) {
		query["CheckTypes"] = request.CheckTypesShrink
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdvisorChecksFoPages"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdvisorChecksFoPagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资源
//
// @param request - DescribeAdvisorResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvisorResourcesResponse
func (client *Client) DescribeAdvisorResourcesWithContext(ctx context.Context, request *DescribeAdvisorResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdvisorResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdvisorResources"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdvisorResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询成本优化结果详情
//
// @param tmpReq - DescribeCostCheckAdvicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCostCheckAdvicesResponse
func (client *Client) DescribeCostCheckAdvicesWithContext(ctx context.Context, tmpReq *DescribeCostCheckAdvicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCostCheckAdvicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeCostCheckAdvicesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssumeAliyunIdList) {
		request.AssumeAliyunIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssumeAliyunIdList, dara.String("AssumeAliyunIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RegionIds) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, dara.String("RegionIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceGroupIdList) {
		request.ResourceGroupIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceGroupIdList, dara.String("ResourceGroupIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceIds) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, dara.String("ResourceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagKeys) {
		request.TagKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKeys, dara.String("TagKeys"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagList) {
		request.TagListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagList, dara.String("TagList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagValues) {
		request.TagValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagValues, dara.String("TagValues"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AssumeAliyunIdListShrink) {
		query["AssumeAliyunIdList"] = request.AssumeAliyunIdListShrink
	}

	if !dara.IsNil(request.CheckId) {
		query["CheckId"] = request.CheckId
	}

	if !dara.IsNil(request.CheckPlanId) {
		query["CheckPlanId"] = request.CheckPlanId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionIdsShrink) {
		query["RegionIds"] = request.RegionIdsShrink
	}

	if !dara.IsNil(request.ResourceGroupIdListShrink) {
		query["ResourceGroupIdList"] = request.ResourceGroupIdListShrink
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceIdsShrink) {
		query["ResourceIds"] = request.ResourceIdsShrink
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.Severity) {
		query["Severity"] = request.Severity
	}

	if !dara.IsNil(request.TagKeysShrink) {
		query["TagKeys"] = request.TagKeysShrink
	}

	if !dara.IsNil(request.TagListShrink) {
		query["TagList"] = request.TagListShrink
	}

	if !dara.IsNil(request.TagValuesShrink) {
		query["TagValues"] = request.TagValuesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCostCheckAdvices"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCostCheckAdvicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询巡检项聚合成本优化结果概览
//
// @param tmpReq - DescribeCostCheckResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCostCheckResultsResponse
func (client *Client) DescribeCostCheckResultsWithContext(ctx context.Context, tmpReq *DescribeCostCheckResultsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCostCheckResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeCostCheckResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssumeAliyunIdList) {
		request.AssumeAliyunIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssumeAliyunIdList, dara.String("AssumeAliyunIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CheckIds) {
		request.CheckIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckIds, dara.String("CheckIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RegionIds) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, dara.String("RegionIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceGroupIdList) {
		request.ResourceGroupIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceGroupIdList, dara.String("ResourceGroupIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceIds) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, dara.String("ResourceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagKeys) {
		request.TagKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKeys, dara.String("TagKeys"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagList) {
		request.TagListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagList, dara.String("TagList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagValues) {
		request.TagValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagValues, dara.String("TagValues"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AssumeAliyunIdListShrink) {
		query["AssumeAliyunIdList"] = request.AssumeAliyunIdListShrink
	}

	if !dara.IsNil(request.CheckIdsShrink) {
		query["CheckIds"] = request.CheckIdsShrink
	}

	if !dara.IsNil(request.CheckPlanId) {
		query["CheckPlanId"] = request.CheckPlanId
	}

	if !dara.IsNil(request.GroupBy) {
		query["GroupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RegionIdsShrink) {
		query["RegionIds"] = request.RegionIdsShrink
	}

	if !dara.IsNil(request.ResourceGroupIdListShrink) {
		query["ResourceGroupIdList"] = request.ResourceGroupIdListShrink
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceIdsShrink) {
		query["ResourceIds"] = request.ResourceIdsShrink
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.Severity) {
		query["Severity"] = request.Severity
	}

	if !dara.IsNil(request.TagKeysShrink) {
		query["TagKeys"] = request.TagKeysShrink
	}

	if !dara.IsNil(request.TagListShrink) {
		query["TagList"] = request.TagListShrink
	}

	if !dara.IsNil(request.TagValuesShrink) {
		query["TagValues"] = request.TagValuesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCostCheckResults"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCostCheckResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 成本优化-概览
//
// @param tmpReq - DescribeCostOptimizationOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCostOptimizationOverviewResponse
func (client *Client) DescribeCostOptimizationOverviewWithContext(ctx context.Context, tmpReq *DescribeCostOptimizationOverviewRequest, runtime *dara.RuntimeOptions) (_result *DescribeCostOptimizationOverviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeCostOptimizationOverviewShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssumeAliyunIdList) {
		request.AssumeAliyunIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssumeAliyunIdList, dara.String("AssumeAliyunIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AssumeAliyunId) {
		query["AssumeAliyunId"] = request.AssumeAliyunId
	}

	if !dara.IsNil(request.AssumeAliyunIdListShrink) {
		query["AssumeAliyunIdList"] = request.AssumeAliyunIdListShrink
	}

	if !dara.IsNil(request.CheckPlanId) {
		query["CheckPlanId"] = request.CheckPlanId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCostOptimizationOverview"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCostOptimizationOverviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 历史
//
// @param request - GetHistoryAdvicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHistoryAdvicesResponse
func (client *Client) GetHistoryAdvicesWithContext(ctx context.Context, request *GetHistoryAdvicesRequest, runtime *dara.RuntimeOptions) (_result *GetHistoryAdvicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.Reverse) {
		query["Reverse"] = request.Reverse
	}

	if !dara.IsNil(request.Severity) {
		query["Severity"] = request.Severity
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHistoryAdvices"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHistoryAdvicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务执行进度(普通用户、RD单账号)
//
// @param request - GetInspectProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInspectProgressResponse
func (client *Client) GetInspectProgressWithContext(ctx context.Context, request *GetInspectProgressRequest, runtime *dara.RuntimeOptions) (_result *GetInspectProgressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssumeAliyunId) {
		query["AssumeAliyunId"] = request.AssumeAliyunId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInspectProgress"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInspectProgressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取云产品的列表
//
// @param request - GetProductListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProductListResponse
func (client *Client) GetProductListWithContext(ctx context.Context, request *GetProductListRequest, runtime *dara.RuntimeOptions) (_result *GetProductListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProductList"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProductListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据id获取任务状态
//
// @param request - GetTaskStatusByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskStatusByIdResponse
func (client *Client) GetTaskStatusByIdWithContext(ctx context.Context, request *GetTaskStatusByIdRequest, runtime *dara.RuntimeOptions) (_result *GetTaskStatusByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskStatusById"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskStatusByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 触发立即巡检
//
// @param tmpReq - RefreshAdvisorCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshAdvisorCheckResponse
func (client *Client) RefreshAdvisorCheckWithContext(ctx context.Context, tmpReq *RefreshAdvisorCheckRequest, runtime *dara.RuntimeOptions) (_result *RefreshAdvisorCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RefreshAdvisorCheckShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceDimensionList) {
		request.ResourceDimensionListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceDimensionList, dara.String("ResourceDimensionList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AssumeAliyunId) {
		query["AssumeAliyunId"] = request.AssumeAliyunId
	}

	if !dara.IsNil(request.CheckId) {
		query["CheckId"] = request.CheckId
	}

	if !dara.IsNil(request.CheckPlanId) {
		query["CheckPlanId"] = request.CheckPlanId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceDimensionListShrink) {
		body["ResourceDimensionList"] = request.ResourceDimensionListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshAdvisorCheck"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshAdvisorCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发起成本优化巡检
//
// @param tmpReq - RefreshAdvisorCostCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshAdvisorCostCheckResponse
func (client *Client) RefreshAdvisorCostCheckWithContext(ctx context.Context, tmpReq *RefreshAdvisorCostCheckRequest, runtime *dara.RuntimeOptions) (_result *RefreshAdvisorCostCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RefreshAdvisorCostCheckShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssumeAliyunIdList) {
		request.AssumeAliyunIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssumeAliyunIdList, dara.String("AssumeAliyunIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CheckIds) {
		request.CheckIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckIds, dara.String("CheckIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceIds) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, dara.String("ResourceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AssumeAliyunIdListShrink) {
		query["AssumeAliyunIdList"] = request.AssumeAliyunIdListShrink
	}

	if !dara.IsNil(request.CheckIdsShrink) {
		query["CheckIds"] = request.CheckIdsShrink
	}

	if !dara.IsNil(request.CheckPlanId) {
		query["CheckPlanId"] = request.CheckPlanId
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RefreshResource) {
		query["RefreshResource"] = request.RefreshResource
	}

	if !dara.IsNil(request.ResourceIdsShrink) {
		query["ResourceIds"] = request.ResourceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshAdvisorCostCheck"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshAdvisorCostCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # RefreshAdvisorResource
//
// @param request - RefreshAdvisorResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshAdvisorResourceResponse
func (client *Client) RefreshAdvisorResourceWithContext(ctx context.Context, request *RefreshAdvisorResourceRequest, runtime *dara.RuntimeOptions) (_result *RefreshAdvisorResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshAdvisorResource"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshAdvisorResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上报用户业务报警信息
//
// @param tmpReq - ReportBizAlertInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportBizAlertInfoResponse
func (client *Client) ReportBizAlertInfoWithContext(ctx context.Context, tmpReq *ReportBizAlertInfoRequest, runtime *dara.RuntimeOptions) (_result *ReportBizAlertInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReportBizAlertInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AlertUid) {
		request.AlertUidShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlertUid, dara.String("AlertUid"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertDescription) {
		query["AlertDescription"] = request.AlertDescription
	}

	if !dara.IsNil(request.AlertDetail) {
		query["AlertDetail"] = request.AlertDetail
	}

	if !dara.IsNil(request.AlertGrade) {
		query["AlertGrade"] = request.AlertGrade
	}

	if !dara.IsNil(request.AlertLabels) {
		query["AlertLabels"] = request.AlertLabels
	}

	if !dara.IsNil(request.AlertScene) {
		query["AlertScene"] = request.AlertScene
	}

	if !dara.IsNil(request.AlertToken) {
		query["AlertToken"] = request.AlertToken
	}

	if !dara.IsNil(request.AlertUidShrink) {
		query["AlertUid"] = request.AlertUidShrink
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportBizAlertInfo"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReportBizAlertInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
