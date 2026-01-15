// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds an image to an Image Search instance.
//
// Description:
//
// You can call this operation to add an image to an Image Search instance.
//
// > If you want to obtain more information about the service and technical support, click [Online Consulting](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or join the DingTalk group (ID 35035130).
//
// ## QPS limits
//
// By default, the concurrency limit for adding an image to instances whose image capacity specifications are 0.1 million images is 1. This means that the system can process up to one request of adding an image every second. By default, the concurrency limit for adding an image to instances of other image capacity specifications is 5. This means that the system can process up to five requests of adding an image every second.
//
// @param request - AddImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImageResponse
func (client *Client) AddImageWithContext(ctx context.Context, request *AddImageRequest, runtime *dara.RuntimeOptions) (_result *AddImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Crop) {
		body["Crop"] = request.Crop
	}

	if !dara.IsNil(request.CustomContent) {
		body["CustomContent"] = request.CustomContent
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IntAttr) {
		body["IntAttr"] = request.IntAttr
	}

	if !dara.IsNil(request.IntAttr2) {
		body["IntAttr2"] = request.IntAttr2
	}

	if !dara.IsNil(request.IntAttr3) {
		body["IntAttr3"] = request.IntAttr3
	}

	if !dara.IsNil(request.IntAttr4) {
		body["IntAttr4"] = request.IntAttr4
	}

	if !dara.IsNil(request.PicContent) {
		body["PicContent"] = request.PicContent
	}

	if !dara.IsNil(request.PicName) {
		body["PicName"] = request.PicName
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.StrAttr) {
		body["StrAttr"] = request.StrAttr
	}

	if !dara.IsNil(request.StrAttr2) {
		body["StrAttr2"] = request.StrAttr2
	}

	if !dara.IsNil(request.StrAttr3) {
		body["StrAttr3"] = request.StrAttr3
	}

	if !dara.IsNil(request.StrAttr4) {
		body["StrAttr4"] = request.StrAttr4
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddImage"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对比图片相似值
//
// @param request - CompareSimilarByImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareSimilarByImageResponse
func (client *Client) CompareSimilarByImageWithContext(ctx context.Context, request *CompareSimilarByImageRequest, runtime *dara.RuntimeOptions) (_result *CompareSimilarByImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PrimaryPicContent) {
		body["PrimaryPicContent"] = request.PrimaryPicContent
	}

	if !dara.IsNil(request.SecondaryPicContent) {
		body["SecondaryPicContent"] = request.SecondaryPicContent
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompareSimilarByImage"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompareSimilarByImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the DeleteImage operation and provides examples of this operation. You can call this operation to delete images from an Image Search instance.
//
// Description:
//
// This operation deletes images from an Image Search instance.
//
// >  A success response is returned even if the specified image does not exist on the instance. Therefore, you cannot determine whether the image exists on the instance based on the response.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 20. In this case, the system can process at most 20 requests every second.
//
// @param request - DeleteImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImageResponse
func (client *Client) DeleteImageWithContext(ctx context.Context, request *DeleteImageRequest, runtime *dara.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IsDeleteByFilter) {
		body["IsDeleteByFilter"] = request.IsDeleteByFilter
	}

	if !dara.IsNil(request.PicName) {
		body["PicName"] = request.PicName
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteImage"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the Detail operation and provides examples of this operation. You can call this operation to query instance details.
//
// Description:
//
// This operation queries instance details.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process only 1 request every second.
//
// @param request - DetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetailResponse
func (client *Client) DetailWithContext(ctx context.Context, request *DetailRequest, runtime *dara.RuntimeOptions) (_result *DetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Detail"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the DumpMeta operation and provides examples of this operation. You can call this operation to create a task for exporting metadata from an Image Search instance.
//
// Description:
//
// This operation creates a task for exporting metadata from an Image Search instance.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process at most 1 request every second.
//
// @param request - DumpMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DumpMetaResponse
func (client *Client) DumpMetaWithContext(ctx context.Context, request *DumpMetaRequest, runtime *dara.RuntimeOptions) (_result *DumpMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DumpMeta"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DumpMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the DumpMetaList operation and provides examples of this operation. You can call this operation to query tasks that are used for exporting metadata from an Image Search instance.
//
// Description:
//
// This operation queries tasks that are used for exporting metadata from an Image Search instance.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process at most 1 request every second.
//
// @param request - DumpMetaListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DumpMetaListResponse
func (client *Client) DumpMetaListWithContext(ctx context.Context, request *DumpMetaListRequest, runtime *dara.RuntimeOptions) (_result *DumpMetaListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DumpMetaList"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DumpMetaListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the IncreaseInstance operation and provides examples of this operation. You can call this operation to create a batch task on an Image Search instance.
//
// Description:
//
// This operation creates a batch task on an Image Search instance.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process at most 1 request every second.
//
// @param request - IncreaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IncreaseInstanceResponse
func (client *Client) IncreaseInstanceWithContext(ctx context.Context, request *IncreaseInstanceRequest, runtime *dara.RuntimeOptions) (_result *IncreaseInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.CallbackAddress) {
		query["CallbackAddress"] = request.CallbackAddress
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IncreaseInstance"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IncreaseInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the IncreaseList operation and provides examples of this operation. You can call this operation to query batch tasks on an Image Search instance.
//
// Description:
//
// This operation queries batch tasks on an Image Search instance.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process at most 1 request every second.
//
// @param request - IncreaseListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IncreaseListResponse
func (client *Client) IncreaseListWithContext(ctx context.Context, request *IncreaseListRequest, runtime *dara.RuntimeOptions) (_result *IncreaseListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IncreaseList"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IncreaseListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the SearchByName operation and provides examples of this operation. You can call this operation to search for images by image name on an Image Search instance.
//
// Description:
//
// This operation searches for images by image name on an Image Search instance.
//
// ## QPS limits
//
// The maximum number of queries per second is displayed in the Image Search console. The upper limit is specified when you purchase the instance. You can set the upper limit to 5 QPS or 10 QPS.
//
// @param request - SearchImageByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchImageByNameResponse
func (client *Client) SearchImageByNameWithContext(ctx context.Context, request *SearchImageByNameRequest, runtime *dara.RuntimeOptions) (_result *SearchImageByNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ScoreThreshold) {
		query["ScoreThreshold"] = request.ScoreThreshold
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.DistinctProductId) {
		body["DistinctProductId"] = request.DistinctProductId
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Num) {
		body["Num"] = request.Num
	}

	if !dara.IsNil(request.PicName) {
		body["PicName"] = request.PicName
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.Start) {
		body["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchImageByName"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchImageByNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the SearchByPic operation and provides examples of this operation. You can call this operation to search for images by image on an Image Search Instance.
//
// Description:
//
// This operation searches for images by image name on an Image Search instance.
//
// ## QPS limits
//
// The maximum number of queries per second is displayed in the Image Search console. The upper limit is specified when you purchase the instance. You can set the upper limit to 5 QPS or 10 QPS.
//
// ## SDK release notes
//
// The Image Search SDK has been upgraded to version 3.1.1, which supports multi-subject recognition and similarity scores. For more information, see [Image Search SDK for Java](/help/en/image-search/latest/version-v3-java-sdk).
//
// @param request - SearchImageByPicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchImageByPicResponse
func (client *Client) SearchImageByPicWithContext(ctx context.Context, request *SearchImageByPicRequest, runtime *dara.RuntimeOptions) (_result *SearchImageByPicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ScoreThreshold) {
		query["ScoreThreshold"] = request.ScoreThreshold
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Crop) {
		body["Crop"] = request.Crop
	}

	if !dara.IsNil(request.DistinctProductId) {
		body["DistinctProductId"] = request.DistinctProductId
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Num) {
		body["Num"] = request.Num
	}

	if !dara.IsNil(request.PicContent) {
		body["PicContent"] = request.PicContent
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.Start) {
		body["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchImageByPic"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchImageByPicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # SearchImageByText
//
// @param request - SearchImageByTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchImageByTextResponse
func (client *Client) SearchImageByTextWithContext(ctx context.Context, request *SearchImageByTextRequest, runtime *dara.RuntimeOptions) (_result *SearchImageByTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ScoreThreshold) {
		query["ScoreThreshold"] = request.ScoreThreshold
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DistinctProductId) {
		body["DistinctProductId"] = request.DistinctProductId
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Num) {
		body["Num"] = request.Num
	}

	if !dara.IsNil(request.Start) {
		body["Start"] = request.Start
	}

	if !dara.IsNil(request.Text) {
		body["Text"] = request.Text
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchImageByText"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchImageByTextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the UpdateImage operation and provides examples of this operation. You can call this operation to update image information on an Image Search instance.
//
// Description:
//
// This operation updates image information on an Image Search instance.
//
// > 	- Limits are imposed on the instance creation time.
//
// >	- This operation is supported by instances that are created in the Singapore (Singapore) region after December 2021. This operation is not supported in other regions.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 20. In this case, the system can process at most 20 requests every second.
//
// @param request - UpdateImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageResponse
func (client *Client) UpdateImageWithContext(ctx context.Context, request *UpdateImageRequest, runtime *dara.RuntimeOptions) (_result *UpdateImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IntAttr3) {
		query["IntAttr3"] = request.IntAttr3
	}

	if !dara.IsNil(request.IntAttr4) {
		query["IntAttr4"] = request.IntAttr4
	}

	if !dara.IsNil(request.StrAttr3) {
		query["StrAttr3"] = request.StrAttr3
	}

	if !dara.IsNil(request.StrAttr4) {
		query["StrAttr4"] = request.StrAttr4
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomContent) {
		body["CustomContent"] = request.CustomContent
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IntAttr) {
		body["IntAttr"] = request.IntAttr
	}

	if !dara.IsNil(request.IntAttr2) {
		body["IntAttr2"] = request.IntAttr2
	}

	if !dara.IsNil(request.PicName) {
		body["PicName"] = request.PicName
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.StrAttr) {
		body["StrAttr"] = request.StrAttr
	}

	if !dara.IsNil(request.StrAttr2) {
		body["StrAttr2"] = request.StrAttr2
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateImage"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
