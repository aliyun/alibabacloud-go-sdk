// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Describes the syntax and provides examples of the AddImage operation, which adds image information to an Image Search instance.
//
// Description:
//
// ## Description
//
// This operation adds image information to an Image Search instance.
//
// ## QPS limit
//
// An instance with a maximum image capacity of 100,000 has a default concurrency of 1, which means that a maximum of 1 image addition request can be processed per second.
//
// Instances with other image capacities have a default concurrency of 5, which means that a maximum of 5 image addition requests can be processed per second.
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
// # CheckImageExists
//
// Description:
//
// ## How-To
//
// This API is used to query image information in an Image Search instance based on an image.
//
// ## QPS Limit
//
// The default maximum queries per second (QPS) for query operations can be viewed in the console. It corresponds to the Visit Frequency (QPS) you selected when purchasing the instance. Supported QPS values are 1, 5, and 10.
//
// ### SDK Version Guide
//
// Upgrade the Image SDK to version V3.1.1 to use the "subject identification" and "similarity score" features. For more information, see [Java SDK](https://help.aliyun.com/document_detail/179188.html).
//
// @param request - CheckImageExistsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckImageExistsResponse
func (client *Client) CheckImageExistsWithContext(ctx context.Context, request *CheckImageExistsRequest, runtime *dara.RuntimeOptions) (_result *CheckImageExistsResponse, _err error) {
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
		Action:      dara.String("CheckImageExists"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckImageExistsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Compares two images and returns a similarity score.
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
// This topic describes the syntax and examples of the DeleteImage operation, which is used to delete image information from an Image Search instance.
//
// Description:
//
// ## Operation description
//
// This operation is used to delete image information from an Image Search instance.
//
// >- If the specified image does not exist in the Image Search instance, this operation still returns a success response. Do not use the response to determine whether the image exists.
//
// ## QPS limit
//
// The default concurrency for delete operations is 20, which means a maximum of 20 delete requests can be processed per second.
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
// This topic describes the syntax and examples of the Detail operation, which queries information about an Image Search instance by name.
//
// Description:
//
// ## Operation description
//
// This operation queries instance information from an Image Search instance.
//
// > For more product details or technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for query operations is 1, which means a maximum of 1 request is processed per second.
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
// This topic describes the syntax and examples of the DumpMeta operation, which creates a metadata export task for Image Search by name.
//
// Description:
//
// ## Operation description
//
// This operation submits a metadata export task to an Image Search instance.
//
// > For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for submit operations is 1, which means a maximum of 1 request is processed per second.
//
// > You cannot submit a new metadata export task while the previous metadata export task is still in progress.
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
// Describes the syntax and provides examples of the DumpMetaList operation, which queries the list of metadata export tasks in an Image Search instance.
//
// Description:
//
// ## Operation description
//
// This operation queries metadata export tasks in an Image Search instance.
//
// > For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for query operations is 1, which means a maximum of 1 request is processed per second.
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
// Describes the syntax and provides examples of the IncreaseInstance operation, which is used to create a batch task for an Image Search instance by name.
//
// Description:
//
// ## Operation description
//
// This operation is used to submit a batch task to an Image Search instance.
//
// > <props="china">For more information about the product or technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us by using DingTalk group 35035130.
//
// ## QPS limit
//
// Only one batch task can run at a time.
//
// > You cannot submit a new batch task until the previous batch task is complete.
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
// Queries the list of batch tasks in an Image Search instance by calling the IncreaseList operation. This topic describes the syntax and provides examples.
//
// Description:
//
// ## Operation description
//
// This operation is used to query batch tasks in an Image Search instance.
//
// > For more product details or technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for query operations is 1, which means a maximum of 1 request is processed per second.
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
// This topic describes the syntax and examples of SearchImageByFilter, which is used to query image information in an Image Search instance based on filter conditions.
//
// @param request - SearchImageByFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchImageByFilterResponse
func (client *Client) SearchImageByFilterWithContext(ctx context.Context, request *SearchImageByFilterRequest, runtime *dara.RuntimeOptions) (_result *SearchImageByFilterResponse, _err error) {
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

	if !dara.IsNil(request.Num) {
		body["Num"] = request.Num
	}

	if !dara.IsNil(request.Start) {
		body["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchImageByFilter"),
		Version:     dara.String("2020-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchImageByFilterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the syntax and examples of the SearchByName operation, which is used to query image information in an Image Search instance by name.
//
// Description:
//
// ### Operation description
//
// This operation queries image information in an Image Search instance by name (ProductId and PicName).
//
// > For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ### QPS limit
//
// The default maximum query rate can be viewed in the console. It is the QPS value you selected at the time of purchase. Currently supported values are 1 QPS, 5 QPS, and 10 QPS.
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
// This topic describes the syntax and examples of SearchByPic, which is used to search for image information in an Image Search instance by image.
//
// Description:
//
// ## Operation description
//
// This operation is used to search for image information in an Image Search instance by image.
//
// > <props="china">For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through DingTalk group (35035130).
//
// ## QPS limit
//
// You can view the default maximum access frequency for query operations in the console. The frequency is the QPS value that you selected when you made the purchase. The supported values are 1 QPS, 5 QPS, and 10 QPS.
//
// ### SDK version description
//
// Upgrade the Image Search SDK to V3.1.1 to use the multi-subject identification and similarity score features. For more information, see [Java SDK](https://help.aliyun.com/document_detail/179188.html).
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
// This topic describes the syntax and examples of SearchImageByText, which is used to search for image information in an Image Search instance based on text.
//
// Description:
//
// ## Operation description
//
// This operation is used to search for image information in an Image Search instance based on text. This operation is available only for instances whose service type is product multimodal search.
//
// > <props="china">For more product details and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through the DingTalk group (35035130).
//
// ## QPS limit
//
// You can view the default maximum access frequency for query operations in the console. The frequency is the QPS value you selected at the time of purchase. Currently supported values are 1 QPS, 5 QPS, and 10 QPS.
//
// ### SDK version description
//
// Upgrade the Image Search SDK to V3.1.1 to use the multi-subject identification and similarity score features. For more information, see [Java SDK](https://help.aliyun.com/document_detail/179188.html).
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
// Updates the image information in an Image Search instance.
//
// Description:
//
// ## Usage notes
//
// This operation updates the image information in an Image Search instance based on the product ID and image name.
//
// > - The instance must meet the creation date requirements.
//
// <props="china">
//
// - Instances created after June 2021 in the Shanghai and Hangzhou regions are supported. Instances in other regions can be used normally.
//
// <props="intl">
//
// - Instances created after December 2021 in the Singapore region are supported. Instances in other regions are currently unavailable.
//
// - For more information about the product and technical support, click [Online Consultation](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or contact us through the DingTalk group (35035130).
//
// ## QPS limit
//
// The default concurrency for update operations is 20, which means that a maximum of 20 requests can be processed per second.
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
