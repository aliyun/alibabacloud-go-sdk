// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// ai播客生成任务结果查询
//
// @param request - PodcastTaskResultQueryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PodcastTaskResultQueryResponse
func (client *Client) PodcastTaskResultQueryWithContext(ctx context.Context, request *PodcastTaskResultQueryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PodcastTaskResultQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PodcastTaskResultQuery"),
		Version:     dara.String("2025-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/podcast/task"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PodcastTaskResultQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ai播客生成任务提交
//
// @param tmpReq - PodcastTaskSubmitRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PodcastTaskSubmitResponse
func (client *Client) PodcastTaskSubmitWithContext(ctx context.Context, tmpReq *PodcastTaskSubmitRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PodcastTaskSubmitResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &PodcastTaskSubmitShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileUrls) {
		request.FileUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileUrls, dara.String("fileUrls"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Voices) {
		request.VoicesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Voices, dara.String("voices"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Counts) {
		body["counts"] = request.Counts
	}

	if !dara.IsNil(request.FileUrlsShrink) {
		body["fileUrls"] = request.FileUrlsShrink
	}

	if !dara.IsNil(request.SourceLang) {
		body["sourceLang"] = request.SourceLang
	}

	if !dara.IsNil(request.Text) {
		body["text"] = request.Text
	}

	if !dara.IsNil(request.Topic) {
		body["topic"] = request.Topic
	}

	if !dara.IsNil(request.VoicesShrink) {
		body["voices"] = request.VoicesShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PodcastTaskSubmit"),
		Version:     dara.String("2025-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/podcast/task/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PodcastTaskSubmitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
