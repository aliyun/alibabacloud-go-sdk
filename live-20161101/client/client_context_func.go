// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds a component to a production studio.
//
// Description:
//
// Before calling this operation, create a production studio and review its layout list. This operation adds components such as images, text, and captions. For more information about creating a production studio using an API call, see [Create a production studio](https://help.aliyun.com/document_detail/2848009.html).
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - AddCasterComponentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCasterComponentResponse
func (client *Client) AddCasterComponentWithContext(ctx context.Context, request *AddCasterComponentRequest, runtime *dara.RuntimeOptions) (_result *AddCasterComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaptionLayerContent) {
		query["CaptionLayerContent"] = request.CaptionLayerContent
	}

	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.ComponentLayer) {
		query["ComponentLayer"] = request.ComponentLayer
	}

	if !dara.IsNil(request.ComponentName) {
		query["ComponentName"] = request.ComponentName
	}

	if !dara.IsNil(request.ComponentType) {
		query["ComponentType"] = request.ComponentType
	}

	if !dara.IsNil(request.Effect) {
		query["Effect"] = request.Effect
	}

	if !dara.IsNil(request.HtmlLayerContent) {
		query["HtmlLayerContent"] = request.HtmlLayerContent
	}

	if !dara.IsNil(request.ImageLayerContent) {
		query["ImageLayerContent"] = request.ImageLayerContent
	}

	if !dara.IsNil(request.LayerOrder) {
		query["LayerOrder"] = request.LayerOrder
	}

	if !dara.IsNil(request.LocationId) {
		query["LocationId"] = request.LocationId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TextLayerContent) {
		query["TextLayerContent"] = request.TextLayerContent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCasterComponent"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCasterComponentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an episode to a production studio.
//
// Description:
//
// Before you call this operation, obtain the ID of the production studio. The ID is generated when the production studio is created.
//
// - If you create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the value of the returned CasterId parameter.
//
// - If you create a production studio in the LIVE console, go to **LIVE Console*	- > **Production Studio*	- > **Production Studio*	- to view the name of the production studio.
//
// > The name of the production studio in the production studio list serves as the production studio ID.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 4 for each account. API calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - AddCasterEpisodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCasterEpisodeResponse
func (client *Client) AddCasterEpisodeWithContext(ctx context.Context, request *AddCasterEpisodeRequest, runtime *dara.RuntimeOptions) (_result *AddCasterEpisodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.ComponentId) {
		query["ComponentId"] = request.ComponentId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EpisodeName) {
		query["EpisodeName"] = request.EpisodeName
	}

	if !dara.IsNil(request.EpisodeType) {
		query["EpisodeType"] = request.EpisodeType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.SwitchType) {
		query["SwitchType"] = request.SwitchType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCasterEpisode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCasterEpisodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a program list to a production studio.
//
// Description:
//
// Create a production studio and obtain its configuration information before calling this operation to add a program list to the production studio. To create a production studio by using an API operation, see [Create a production studio](https://help.aliyun.com/document_detail/2848009.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 4 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation as needed.
//
// @param request - AddCasterEpisodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCasterEpisodeGroupResponse
func (client *Client) AddCasterEpisodeGroupWithContext(ctx context.Context, request *AddCasterEpisodeGroupRequest, runtime *dara.RuntimeOptions) (_result *AddCasterEpisodeGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Item) {
		query["Item"] = request.Item
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepeatNum) {
		query["RepeatNum"] = request.RepeatNum
	}

	if !dara.IsNil(request.SideOutputUrl) {
		query["SideOutputUrl"] = request.SideOutputUrl
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCasterEpisodeGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCasterEpisodeGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds content to an episode group of a production studio.
//
// Description:
//
// Before you call this operation, you must create a production studio and an episode group. For more information, see [Create a production studio](https://help.aliyun.com/document_detail/2848009.html).
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 4 for each user. If you exceed this limit, your API calls are throttled. This may impact your business. Plan your calls accordingly.
//
// @param request - AddCasterEpisodeGroupContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCasterEpisodeGroupContentResponse
func (client *Client) AddCasterEpisodeGroupContentWithContext(ctx context.Context, request *AddCasterEpisodeGroupContentRequest, runtime *dara.RuntimeOptions) (_result *AddCasterEpisodeGroupContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCasterEpisodeGroupContent"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCasterEpisodeGroupContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a layout to a production studio.
//
// Description:
//
// Before you call this operation, you must create a production studio and add video sources. For more information about how to create a production studio, see [CreateCaster](https://help.aliyun.com/document_detail/2848009.html).
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second per user. API calls that exceed this limit are throttled, which may affect your business. Do not exceed this limit.
//
// @param request - AddCasterLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCasterLayoutResponse
func (client *Client) AddCasterLayoutWithContext(ctx context.Context, request *AddCasterLayoutRequest, runtime *dara.RuntimeOptions) (_result *AddCasterLayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AudioLayer) {
		query["AudioLayer"] = request.AudioLayer
	}

	if !dara.IsNil(request.BlendList) {
		query["BlendList"] = request.BlendList
	}

	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.MixList) {
		query["MixList"] = request.MixList
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VideoLayer) {
		query["VideoLayer"] = request.VideoLayer
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCasterLayout"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCasterLayoutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a program list to a production studio.
//
// Description:
//
// Create a production studio and add video resources to it before calling this operation to add a program list. This operation currently supports only two node types: video source and component. To create a production studio by using an API operation, see [CreateCaster](https://help.aliyun.com/document_detail/2848009.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - AddCasterProgramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCasterProgramResponse
func (client *Client) AddCasterProgramWithContext(ctx context.Context, request *AddCasterProgramRequest, runtime *dara.RuntimeOptions) (_result *AddCasterProgramResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.Episode) {
		query["Episode"] = request.Episode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCasterProgram"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCasterProgramResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a video source to a production studio. The number of video sources is limited by the number of input channels of the production studio.
//
// Description:
//
// Create a production studio before calling this operation to add a video source. The number of video sources is limited by the number of input channels of the production studio. To create a production studio by using an API operation, see [CreateCaster](https://help.aliyun.com/document_detail/2848009.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - AddCasterVideoResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCasterVideoResourceResponse
func (client *Client) AddCasterVideoResourceWithContext(ctx context.Context, request *AddCasterVideoResourceRequest, runtime *dara.RuntimeOptions) (_result *AddCasterVideoResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginOffset) {
		query["BeginOffset"] = request.BeginOffset
	}

	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.EndOffset) {
		query["EndOffset"] = request.EndOffset
	}

	if !dara.IsNil(request.FixedDelayDuration) {
		query["FixedDelayDuration"] = request.FixedDelayDuration
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.LiveStreamUrl) {
		query["LiveStreamUrl"] = request.LiveStreamUrl
	}

	if !dara.IsNil(request.LocationId) {
		query["LocationId"] = request.LocationId
	}

	if !dara.IsNil(request.MaterialId) {
		query["MaterialId"] = request.MaterialId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PtsCallbackInterval) {
		query["PtsCallbackInterval"] = request.PtsCallbackInterval
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepeatNum) {
		query["RepeatNum"] = request.RepeatNum
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.VodUrl) {
		query["VodUrl"] = request.VodUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCasterVideoResource"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCasterVideoResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a custom transcoding configuration.
//
// Description:
//
// This operation supports the following types of custom transcoding templates:
//
// - h264: H.264 standard transcoding.
//
// - h264-nbhd: H.264 Narrowband HD™ transcoding.
//
// - h265: H.265 standard transcoding.
//
// - h265-nbhd: H.265 Narrowband HD™ transcoding.
//
// - audio: audio-only transcoding.
//
// ## QPS limit
//
// You can call this operation up to 6,000 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions. For more information, see [QPS limit](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - AddCustomLiveStreamTranscodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomLiveStreamTranscodeResponse
func (client *Client) AddCustomLiveStreamTranscodeWithContext(ctx context.Context, request *AddCustomLiveStreamTranscodeRequest, runtime *dara.RuntimeOptions) (_result *AddCustomLiveStreamTranscodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.AudioBitrate) {
		query["AudioBitrate"] = request.AudioBitrate
	}

	if !dara.IsNil(request.AudioChannelNum) {
		query["AudioChannelNum"] = request.AudioChannelNum
	}

	if !dara.IsNil(request.AudioCodec) {
		query["AudioCodec"] = request.AudioCodec
	}

	if !dara.IsNil(request.AudioProfile) {
		query["AudioProfile"] = request.AudioProfile
	}

	if !dara.IsNil(request.AudioRate) {
		query["AudioRate"] = request.AudioRate
	}

	if !dara.IsNil(request.BitrateWithSource) {
		query["BitrateWithSource"] = request.BitrateWithSource
	}

	if !dara.IsNil(request.DeInterlaced) {
		query["DeInterlaced"] = request.DeInterlaced
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EncryptParameters) {
		query["EncryptParameters"] = request.EncryptParameters
	}

	if !dara.IsNil(request.ExtWithSource) {
		query["ExtWithSource"] = request.ExtWithSource
	}

	if !dara.IsNil(request.FPS) {
		query["FPS"] = request.FPS
	}

	if !dara.IsNil(request.FpsWithSource) {
		query["FpsWithSource"] = request.FpsWithSource
	}

	if !dara.IsNil(request.Gop) {
		query["Gop"] = request.Gop
	}

	if !dara.IsNil(request.Height) {
		query["Height"] = request.Height
	}

	if !dara.IsNil(request.KmsKeyExpireInterval) {
		query["KmsKeyExpireInterval"] = request.KmsKeyExpireInterval
	}

	if !dara.IsNil(request.KmsKeyID) {
		query["KmsKeyID"] = request.KmsKeyID
	}

	if !dara.IsNil(request.KmsUID) {
		query["KmsUID"] = request.KmsUID
	}

	if !dara.IsNil(request.Lazy) {
		query["Lazy"] = request.Lazy
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Profile) {
		query["Profile"] = request.Profile
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResWithSource) {
		query["ResWithSource"] = request.ResWithSource
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.VideoBitrate) {
		query["VideoBitrate"] = request.VideoBitrate
	}

	if !dara.IsNil(request.Width) {
		query["Width"] = request.Width
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCustomLiveStreamTranscode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCustomLiveStreamTranscodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a subtitle rule.
//
// Description:
//
// - After you add a subtitle template by calling the [AddLiveAISubtitle](https://help.aliyun.com/document_detail/2848222.html) operation, you can call this operation to add a subtitle rule.
//
// - To play a subtitle stream, append _{Subtitle template name} to the StreamName in the playback URL.
//
//   - RTMP: rtmp\\://example.aliyundoc.com/app/stream_{Subtitle template name}?auth_key={access token}
//
//   - FLV: http\\://example.aliyundoc.com/app/stream_{Subtitle template name}.flv?auth_key={access token}
//
//   - M3U8: http\\://example.aliyundoc.com/app/stream_{Subtitle template name}.m3u8?auth_key={access token}
//
//     Notice:
//
// The real-time subtitle feature is in invitational preview. You can add up to 300 subtitle templates.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 60 per user. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - AddLiveAIProduceRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveAIProduceRulesResponse
func (client *Client) AddLiveAIProduceRulesWithContext(ctx context.Context, request *AddLiveAIProduceRulesRequest, runtime *dara.RuntimeOptions) (_result *AddLiveAIProduceRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.IsLazy) {
		query["IsLazy"] = request.IsLazy
	}

	if !dara.IsNil(request.LiveTemplate) {
		query["LiveTemplate"] = request.LiveTemplate
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StudioName) {
		query["StudioName"] = request.StudioName
	}

	if !dara.IsNil(request.SubtitleName) {
		query["SubtitleName"] = request.SubtitleName
	}

	if !dara.IsNil(request.Suffix) {
		query["Suffix"] = request.Suffix
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveAIProduceRules"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveAIProduceRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a subtitle template by calling AddLiveAISubtitle.
//
// Description:
//
// ## Operation description
//
// - This operation allows you to add a live subtitle template configuration. The template configuration describes subtitle content, layout, and other information.
//
// - After you add a subtitle template, call the [AddLiveAIProduceRules](https://help.aliyun.com/document_detail/2799676.html) operation to add subtitle rules. The subtitles take effect in the stream only after you re-ingest the stream.
//
// - Real-time subtitles are supported only in the Beijing, Shanghai, Singapore, Indonesia, and Saudi Arabia regions.
//
//	Notice: The real-time subtitle feature is in public preview. Each user can add up to 300 subtitle templates. The feature is free of charge during the public preview. After the public preview ends, standard billing applies. The specific date will be announced separately.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 60 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as appropriate. For more information, see [QPS limit](https://help.aliyun.com/document_detail/343507.html).
//
// @param tmpReq - AddLiveAISubtitleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveAISubtitleResponse
func (client *Client) AddLiveAISubtitleWithContext(ctx context.Context, tmpReq *AddLiveAISubtitleRequest, runtime *dara.RuntimeOptions) (_result *AddLiveAISubtitleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddLiveAISubtitleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PositionNormalized) {
		request.PositionNormalizedShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PositionNormalized, dara.String("PositionNormalized"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BgColor) {
		query["BgColor"] = request.BgColor
	}

	if !dara.IsNil(request.BgWidthNormalized) {
		query["BgWidthNormalized"] = request.BgWidthNormalized
	}

	if !dara.IsNil(request.BorderWidthNormalized) {
		query["BorderWidthNormalized"] = request.BorderWidthNormalized
	}

	if !dara.IsNil(request.CopyFrom) {
		query["CopyFrom"] = request.CopyFrom
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DstLanguage) {
		query["DstLanguage"] = request.DstLanguage
	}

	if !dara.IsNil(request.FontColor) {
		query["FontColor"] = request.FontColor
	}

	if !dara.IsNil(request.FontName) {
		query["FontName"] = request.FontName
	}

	if !dara.IsNil(request.FontSizeNormalized) {
		query["FontSizeNormalized"] = request.FontSizeNormalized
	}

	if !dara.IsNil(request.Height) {
		query["Height"] = request.Height
	}

	if !dara.IsNil(request.MaxLines) {
		query["MaxLines"] = request.MaxLines
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PositionNormalizedShrink) {
		query["PositionNormalized"] = request.PositionNormalizedShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ShowSourceLan) {
		query["ShowSourceLan"] = request.ShowSourceLan
	}

	if !dara.IsNil(request.SrcLanguage) {
		query["SrcLanguage"] = request.SrcLanguage
	}

	if !dara.IsNil(request.SubtitleName) {
		query["SubtitleName"] = request.SubtitleName
	}

	if !dara.IsNil(request.Width) {
		query["Width"] = request.Width
	}

	if !dara.IsNil(request.WordPerLine) {
		query["WordPerLine"] = request.WordPerLine
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveAISubtitle"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveAISubtitleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures recording for an application and saves the output to Object Storage Service (OSS).
//
// Description:
//
// - Before you use this operation, make sure that you fully understand the billing methods and pricing of live stream recording. For billing details, see [Live stream recording fees](https://help.aliyun.com/document_detail/195287.html).
//
// - If you use the method of storing recordings in OSS to configure live stream recording, activate OSS and create a bucket. For more information, see [Configure OSS](https://help.aliyun.com/document_detail/84932.html).
//
// - Recordings stored in OSS incur storage fees. For billing details in OSS, see [Storage fees](https://help.aliyun.com/document_detail/173534.html).
//
// - The OSS bucket must be in the same region as the live center of the streaming domain. Cross-region recording is not supported.
//
// - The live stream recording feature records live content and saves it to a specified location for on-demand playback. Recordings stored in OSS support multiple container formats (TS, MP4, FLV, and CMAF) and custom recording policies (automatic recording, on-demand recording, and manual recording). Call this operation to configure recording templates. For more information about live stream recording, see [Live stream recording](https://help.aliyun.com/document_detail/199357.html).
//
// - The triplet (DomainName, AppName, StreamName) can correspond to only one configuration. If a configuration already exists for the triplet, calling this operation to add another configuration returns a configuration-already-exists error.
//
// - Configurations set through this operation take effect only after the live stream is re-ingested and remain effective permanently.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 30 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - AddLiveAppRecordConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveAppRecordConfigResponse
func (client *Client) AddLiveAppRecordConfigWithContext(ctx context.Context, request *AddLiveAppRecordConfigRequest, runtime *dara.RuntimeOptions) (_result *AddLiveAppRecordConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DelayTime) {
		query["DelayTime"] = request.DelayTime
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OnDemand) {
		query["OnDemand"] = request.OnDemand
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RecordFormat) {
		query["RecordFormat"] = request.RecordFormat
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.TranscodeRecordFormat) {
		query["TranscodeRecordFormat"] = request.TranscodeRecordFormat
	}

	if !dara.IsNil(request.TranscodeTemplates) {
		query["TranscodeTemplates"] = request.TranscodeTemplates
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveAppRecordConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveAppRecordConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the snapshot feature for a streaming domain. The captured snapshots are stored in Object Storage Service (OSS). The configuration takes effect after you restart stream ingest.
//
// Description:
//
// - Before you call this operation, make sure that you fully understand the billing method and pricing of live stream snapshots in ApsaraVideo Live. For more information, see [Billing of live stream snapshots](https://help.aliyun.com/document_detail/195286.html).
//
// - Make sure that Object Storage Service (OSS) is activated and a specific bucket is created. This way, ApsaraVideo Live can store live stream snapshots in the bucket. For more information, see [Configure OSS](https://help.aliyun.com/document_detail/84932.html).
//
// - If you store snapshots in OSS, storage fees are generated. For more information, see [Storage fees](https://help.aliyun.com/document_detail/173534.html).
//
// - The OSS bucket must reside in the same region as the live center of the streaming domain. Cross-region snapshot capture is not supported.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - AddLiveAppSnapshotConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveAppSnapshotConfigResponse
func (client *Client) AddLiveAppSnapshotConfigWithContext(ctx context.Context, request *AddLiveAppSnapshotConfigRequest, runtime *dara.RuntimeOptions) (_result *AddLiveAppSnapshotConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Callback) {
		query["Callback"] = request.Callback
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OverwriteOssObject) {
		query["OverwriteOssObject"] = request.OverwriteOssObject
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SequenceOssObject) {
		query["SequenceOssObject"] = request.SequenceOssObject
	}

	if !dara.IsNil(request.TimeInterval) {
		query["TimeInterval"] = request.TimeInterval
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveAppSnapshotConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveAppSnapshotConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an audio moderation configuration.
//
// Description:
//
// - The audio moderation feature detects sensitive content in audio from live streams and uses callbacks to notify you of violations in real time. You can then manually review the content and take appropriate actions.
//
// - Currently, only some live centers support automated review. For supported regions, see [Service regions](https://help.aliyun.com/document_detail/193730.html).
//
// <props="china">
//
// Before you call this API, make sure that you understand the billing methods and pricing of the live audio moderation service. For more information, see [Automated review fees](https://help.aliyun.com/document_detail/195288.html).
//
// ## QPS limits
//
// You can call this operation up to 10 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - AddLiveAudioAuditConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveAudioAuditConfigResponse
func (client *Client) AddLiveAudioAuditConfigWithContext(ctx context.Context, request *AddLiveAudioAuditConfigRequest, runtime *dara.RuntimeOptions) (_result *AddLiveAudioAuditConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssObject) {
		query["OssObject"] = request.OssObject
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveAudioAuditConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveAudioAuditConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a callback configuration for audio moderation.
//
// Description:
//
// - The automated review service returns review results based on the configured callback template.
//
// - The automated review feature is available only in some live centers. For supported regions, see [Service regions](https://help.aliyun.com/document_detail/193730.html).
//
// ## QPS limit
//
// You can call this operation up to 10 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - AddLiveAudioAuditNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveAudioAuditNotifyConfigResponse
func (client *Client) AddLiveAudioAuditNotifyConfigWithContext(ctx context.Context, request *AddLiveAudioAuditNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *AddLiveAudioAuditNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Callback) {
		query["Callback"] = request.Callback
	}

	if !dara.IsNil(request.CallbackTemplate) {
		query["CallbackTemplate"] = request.CallbackTemplate
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveAudioAuditNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveAudioAuditNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a live center stream relay configuration.
//
// Description:
//
// The single-user QPS limit of this API is 100 calls per second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please make calls appropriately.
//
// @param request - AddLiveCenterTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveCenterTransferResponse
func (client *Client) AddLiveCenterTransferWithContext(ctx context.Context, request *AddLiveCenterTransferRequest, runtime *dara.RuntimeOptions) (_result *AddLiveCenterTransferResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DstUrl) {
		query["DstUrl"] = request.DstUrl
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.TransferArgs) {
		query["TransferArgs"] = request.TransferArgs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveCenterTransfer"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveCenterTransferResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures callbacks for video moderation results. As a result, a callback URL that is used to receive the callback notifications is added.
//
// Description:
//
// - The automated review feature sends notifications about violations to the callback URL in real time. Then, you can manually review the content and take actions accordingly.
//
// - Only some live centers support the automated review feature. For more information, see [Supported regions](https://help.aliyun.com/document_detail/193730.html).
//
// ## QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit on an API operation in ApsaraVideo Live](https://www.alibabacloud.com/help/en/apsaravideo-live/latest/qps-limit-on-an-api-operation-in-apsaravideo-live).
//
// @param request - AddLiveDetectNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveDetectNotifyConfigResponse
func (client *Client) AddLiveDetectNotifyConfigWithContext(ctx context.Context, request *AddLiveDetectNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *AddLiveDetectNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.NotifyUrl) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveDetectNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveDetectNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a live streaming domain name. You can add only one domain name at a time.
//
// Description:
//
// - Before you add a live streaming domain name, you must activate ApsaraVideo Live. For more information, see [Activate the service](https://help.aliyun.com/document_detail/195292.html).
//
// - Before you add a new domain name, you must first verify the domain name ownership and then call this operation to add the domain name. You can use DNS resolution verification or file verification. For more information, see [Verify domain name ownership](https://help.aliyun.com/document_detail/184466.html).
//
// - ApsaraVideo Live requires both stream ingest and streaming. You must add an ingest domain and a streaming domain separately. You can commit only one domain name at a time.
//
// - After you add a domain name, you must configure CNAME resolution for the domain name. For more information, see [Configure CNAME resolution](https://help.aliyun.com/document_detail/84929.html).
//
// - After you add an ingest domain and a streaming domain, you must associate the associated domains before you can use ApsaraVideo Live. For more information, see [Associated domain](https://help.aliyun.com/document_detail/199338.html).
//
//	Notice: Starting from February 19, 2019, domain names added by using AddLiveDomain do not support live center ingest. New domain names added by using CDN API operations also do not support live center ingest. Use edge ingest to add an ingest domain (call AddLiveDomain with the business type set to liveEdge) and a streaming domain (call AddLiveDomain with the business type set to liveVideo), and then associate them (call [AddLiveDomainMapping](https://help.aliyun.com/document_detail/2847792.html)). Domain names added before February 19, 2019 are not affected.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 30 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - AddLiveDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveDomainResponse
func (client *Client) AddLiveDomainWithContext(ctx context.Context, request *AddLiveDomainRequest, runtime *dara.RuntimeOptions) (_result *AddLiveDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckUrl) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.LiveDomainType) {
		query["LiveDomainType"] = request.LiveDomainType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveDomain"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates the mapping between a streaming domain and an ingest domain.
//
// Description:
//
// Call the [AddLiveDomain](https://help.aliyun.com/document_detail/88327.html) operation to add a streaming domain and an ingest domain, and then call this operation to create the mapping between the two domain names.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - AddLiveDomainMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveDomainMappingResponse
func (client *Client) AddLiveDomainMappingWithContext(ctx context.Context, request *AddLiveDomainMappingRequest, runtime *dara.RuntimeOptions) (_result *AddLiveDomainMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PullDomain) {
		query["PullDomain"] = request.PullDomain
	}

	if !dara.IsNil(request.PushDomain) {
		query["PushDomain"] = request.PushDomain
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveDomainMapping"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveDomainMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Maps a sub-streaming domain to a main streaming domain.
//
// Description:
//
// Before you call this operation, you must add the main streaming domain and sub-streaming domain by calling [AddLiveDomain](https://help.aliyun.com/document_detail/88327.html).
//
// ## QPS limit
//
// You can call this operation up to 1,000 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - AddLiveDomainPlayMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveDomainPlayMappingResponse
func (client *Client) AddLiveDomainPlayMappingWithContext(ctx context.Context, request *AddLiveDomainPlayMappingRequest, runtime *dara.RuntimeOptions) (_result *AddLiveDomainPlayMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayDomain) {
		query["PlayDomain"] = request.PlayDomain
	}

	if !dara.IsNil(request.PullDomain) {
		query["PullDomain"] = request.PullDomain
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveDomainPlayMapping"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveDomainPlayMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The AddLiveMessageGroupBand operation mutes users in a group.
//
// Description:
//
// Before you call this operation, you must call the [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html) operation to create an interactive message group.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. API calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param tmpReq - AddLiveMessageGroupBandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveMessageGroupBandResponse
func (client *Client) AddLiveMessageGroupBandWithContext(ctx context.Context, tmpReq *AddLiveMessageGroupBandRequest, runtime *dara.RuntimeOptions) (_result *AddLiveMessageGroupBandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddLiveMessageGroupBandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BannedUsers) {
		request.BannedUsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BannedUsers, dara.String("BannedUsers"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BannedUsersShrink) {
		query["BannedUsers"] = request.BannedUsersShrink
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveMessageGroupBand"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveMessageGroupBandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the AddLivePackageConfig operation to add a live stream packaging configuration.
//
// Description:
//
// - When you add a packaging configuration to a domain name for the first time, the related acceleration configurations for the playback domain name are also applied. The configurations take effect in 3 to 5 minutes.
//
// - If the playback domain name is in a region outside China, such as Singapore, Germany, Japan, or Indonesia, high latency may occur. After you add the configuration, test it to ensure it works as expected.
//
// - After you add a live stream packaging configuration, you must restart the stream ingest for the configuration to take effect.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 300 calls per minute. If you exceed the limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - AddLivePackageConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLivePackageConfigResponse
func (client *Client) AddLivePackageConfigWithContext(ctx context.Context, request *AddLivePackageConfigRequest, runtime *dara.RuntimeOptions) (_result *AddLivePackageConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.IgnoreTranscode) {
		query["IgnoreTranscode"] = request.IgnoreTranscode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PartDuration) {
		query["PartDuration"] = request.PartDuration
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SegmentDuration) {
		query["SegmentDuration"] = request.SegmentDuration
	}

	if !dara.IsNil(request.SegmentNum) {
		query["SegmentNum"] = request.SegmentNum
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLivePackageConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLivePackageConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scheduled stream pulling configuration, which includes parameters such as the origin URL, start time, and end time.
//
// Description:
//
// - Before you call this operation, make sure that you understand the billing method and pricing of ApsaraVideo Live. For more information, see [](t1961174.xdita#).
//
// - Stream pulling refers to the process of pulling live streams from third-party URLs to a live center of ApsaraVideo Live for CDN acceleration.
//
// - This operation supports only scheduled stream pulling. You can specify a start time and an end time to pull a live stream during a specific time period.
//
// - The console supports both scheduled and triggered stream pulling. For more information, see [Configure stream pulling](https://help.aliyun.com/document_detail/199452.html).
//
// - You can specify custom values for the AppName and StreamName parameters. Streaming URLs are generated based on AppName and StreamName. You can use the [](t2020590.xdita#) to generate a streaming URL.
//
// - Each stream pulling configuration must be unique. The combination of DomainName, AppName, and StreamName can only be associated with one active configuration. Attempting to add a duplicate configuration will result in an error.
//
// ## QPS limit
//
// You can call this operation up to 15 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - AddLivePullStreamInfoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLivePullStreamInfoConfigResponse
func (client *Client) AddLivePullStreamInfoConfigWithContext(ctx context.Context, request *AddLivePullStreamInfoConfigRequest, runtime *dara.RuntimeOptions) (_result *AddLivePullStreamInfoConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceUrl) {
		query["SourceUrl"] = request.SourceUrl
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLivePullStreamInfoConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLivePullStreamInfoConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a domain-level recording callback configuration.
//
// Description:
//
// Before calling this operation to add a domain-level recording callback configuration, query the live recording callback configuration first. To query the live recording callback configuration by using an API operation, see [Query live recording callback configuration](https://help.aliyun.com/document_detail/2847893.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 30 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - AddLiveRecordNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveRecordNotifyConfigResponse
func (client *Client) AddLiveRecordNotifyConfigWithContext(ctx context.Context, request *AddLiveRecordNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *AddLiveRecordNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.NeedStatusNotify) {
		query["NeedStatusNotify"] = request.NeedStatusNotify
	}

	if !dara.IsNil(request.NotifyAuthKey) {
		query["NotifyAuthKey"] = request.NotifyAuthKey
	}

	if !dara.IsNil(request.NotifyReqAuth) {
		query["NotifyReqAuth"] = request.NotifyReqAuth
	}

	if !dara.IsNil(request.NotifyUrl) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !dara.IsNil(request.OnDemandUrl) {
		query["OnDemandUrl"] = request.OnDemandUrl
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveRecordNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveRecordNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a live-to-VOD configuration to store recordings in ApsaraVideo VOD.
//
// Description:
//
// - Storing recordings in ApsaraVideo for VOD triggers automatic operations like video merging and transcoding, which incur video processing fees within the service. For details, see [video editing billing](https://help.aliyun.com/document_detail/188310.html) and [video transcoding billing](https://help.aliyun.com/document_detail/188308.html). For FAQs on automatic merging and transcoding, see the [live-to-VOD FAQ](https://help.aliyun.com/document_detail/99726.html).
//
// - Alibaba Finance Cloud accounts do not support the live-to-VOD feature.
//
// ## QPS limit
//
// The QPS limit for this API is 1,000 calls per minute per user. Exceeding this limit will cause API calls to be throttled, which can impact your business. To avoid throttling, call the API at a reasonable rate.
//
// @param request - AddLiveRecordVodConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveRecordVodConfigResponse
func (client *Client) AddLiveRecordVodConfigWithContext(ctx context.Context, request *AddLiveRecordVodConfigRequest, runtime *dara.RuntimeOptions) (_result *AddLiveRecordVodConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AutoCompose) {
		query["AutoCompose"] = request.AutoCompose
	}

	if !dara.IsNil(request.ComposeVodTranscodeGroupId) {
		query["ComposeVodTranscodeGroupId"] = request.ComposeVodTranscodeGroupId
	}

	if !dara.IsNil(request.CycleDuration) {
		query["CycleDuration"] = request.CycleDuration
	}

	if !dara.IsNil(request.DelayTime) {
		query["DelayTime"] = request.DelayTime
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OnDemand) {
		query["OnDemand"] = request.OnDemand
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RecordContent) {
		query["RecordContent"] = request.RecordContent
	}

	if !dara.IsNil(request.RecordFormat) {
		query["RecordFormat"] = request.RecordFormat
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SpaceId) {
		query["SpaceId"] = request.SpaceId
	}

	if !dara.IsNil(request.StorageLocation) {
		query["StorageLocation"] = request.StorageLocation
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.TranscodeTemplates) {
		query["TranscodeTemplates"] = request.TranscodeTemplates
	}

	if !dara.IsNil(request.VodTranscodeGroupId) {
		query["VodTranscodeGroupId"] = request.VodTranscodeGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveRecordVodConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveRecordVodConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a video moderation configuration for live streams in an application under a domain name.
//
// Description:
//
// - The live streaming audit function identifies and reviews non-compliant sensitive content at the domain and App level, and promptly notifies users of such violations via callbacks. Users can then review the content and take appropriate actions.
//
// - Currently, only some live streaming centers support intelligent auditing. For a list of live streaming centers that support this feature, please refer to [Service Regions](https://help.aliyun.com/document_detail/193730.html).
//
// ## QPS Limitation
//
// The QPS limit for this API per user is 30 requests/second. Exceeding this limit will result in API throttling, which may impact your services. Please use the API judiciously. For more information, see [QPS Limitations](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - AddLiveSnapshotDetectPornConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveSnapshotDetectPornConfigResponse
func (client *Client) AddLiveSnapshotDetectPornConfigWithContext(ctx context.Context, request *AddLiveSnapshotDetectPornConfigRequest, runtime *dara.RuntimeOptions) (_result *AddLiveSnapshotDetectPornConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssObject) {
		query["OssObject"] = request.OssObject
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveSnapshotDetectPornConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveSnapshotDetectPornConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures snapshot callbacks.
//
// Description:
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - AddLiveSnapshotNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveSnapshotNotifyConfigResponse
func (client *Client) AddLiveSnapshotNotifyConfigWithContext(ctx context.Context, request *AddLiveSnapshotNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *AddLiveSnapshotNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.NotifyAuthKey) {
		query["NotifyAuthKey"] = request.NotifyAuthKey
	}

	if !dara.IsNil(request.NotifyReqAuth) {
		query["NotifyReqAuth"] = request.NotifyReqAuth
	}

	if !dara.IsNil(request.NotifyUrl) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveSnapshotNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveSnapshotNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call AddLiveStreamMerge to add a primary/backup stream merge configuration.
//
// Description:
//
// Exceeding the API limit of 100 QPS per account triggers throttling, which may disrupt your service. Please make calls reasonably.
//
// @param request - AddLiveStreamMergeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveStreamMergeResponse
func (client *Client) AddLiveStreamMergeWithContext(ctx context.Context, request *AddLiveStreamMergeRequest, runtime *dara.RuntimeOptions) (_result *AddLiveStreamMergeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InAppName1) {
		query["InAppName1"] = request.InAppName1
	}

	if !dara.IsNil(request.InAppName2) {
		query["InAppName2"] = request.InAppName2
	}

	if !dara.IsNil(request.InStreamName1) {
		query["InStreamName1"] = request.InStreamName1
	}

	if !dara.IsNil(request.InStreamName2) {
		query["InStreamName2"] = request.InStreamName2
	}

	if !dara.IsNil(request.LiveMerger) {
		query["LiveMerger"] = request.LiveMerger
	}

	if !dara.IsNil(request.MergeParameters) {
		query["MergeParameters"] = request.MergeParameters
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SelectAppName) {
		query["SelectAppName"] = request.SelectAppName
	}

	if !dara.IsNil(request.SelectStreamName) {
		query["SelectStreamName"] = request.SelectStreamName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveStreamMerge"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveStreamMergeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls AddLiveStreamTranscode to add transcoding configuration information.
//
// Description:
//
// - Using the live stream transcoding feature incurs transcoding fees, which are calculated based on the transcoding standard, resolution specification, and total transcoding length. For billing rules, see [Live stream transcoding fees](https://help.aliyun.com/document_detail/90424.html).
//
// - First obtain the user KMS master key ID through Key Management Service (KMS), and then invoke this operation to add default transcoding configuration information. Currently, this operation supports only two types of transcoding templates: standard quality templates and Narrowband HD™ transcoding templates.
//
// - Alibaba Cloud KMS provides default keys for server-side encryption of cloud services at no cost, and no instance purchase is required. However, if you need to build a custom application cryptographic solution, use the credential feature, or manage the key lifecycle, you must purchase a software or hardware key management instance. For billing details, see [Billing](https://help.aliyun.com/document_detail/600418.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 30 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Invoke this operation appropriately.
//
// @param request - AddLiveStreamTranscodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveStreamTranscodeResponse
func (client *Client) AddLiveStreamTranscodeWithContext(ctx context.Context, request *AddLiveStreamTranscodeRequest, runtime *dara.RuntimeOptions) (_result *AddLiveStreamTranscodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EncryptParameters) {
		query["EncryptParameters"] = request.EncryptParameters
	}

	if !dara.IsNil(request.Lazy) {
		query["Lazy"] = request.Lazy
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveStreamTranscode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveStreamTranscodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a watermark template.
//
// Description:
//
// - This API creates a live stream watermark template, which defines the watermark\\"s content, layout, and other properties.
//
// - After you add a watermark template, you must call the [AddLiveStreamWatermarkRule](https://help.aliyun.com/document_detail/2848100.html) operation to create a rule that applies the template. The watermark appears on the stream after you restart the stream ingest.
//
// ## QPS limit
//
// You can call this operation up to 60 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - AddLiveStreamWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveStreamWatermarkResponse
func (client *Client) AddLiveStreamWatermarkWithContext(ctx context.Context, request *AddLiveStreamWatermarkRequest, runtime *dara.RuntimeOptions) (_result *AddLiveStreamWatermarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Height) {
		query["Height"] = request.Height
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OffsetCorner) {
		query["OffsetCorner"] = request.OffsetCorner
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PictureUrl) {
		query["PictureUrl"] = request.PictureUrl
	}

	if !dara.IsNil(request.RefHeight) {
		query["RefHeight"] = request.RefHeight
	}

	if !dara.IsNil(request.RefWidth) {
		query["RefWidth"] = request.RefWidth
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Transparency) {
		query["Transparency"] = request.Transparency
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.XOffset) {
		query["XOffset"] = request.XOffset
	}

	if !dara.IsNil(request.YOffset) {
		query["YOffset"] = request.YOffset
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveStreamWatermark"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveStreamWatermarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a watermark rule.
//
// Description:
//
// After you add a watermark template by calling the [AddLiveStreamWatermark](https://help.aliyun.com/document_detail/2848096.html) operation, call this operation to create a rule that applies the template.
//
// ## QPS limits
//
// You can call this operation up to 60 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - AddLiveStreamWatermarkRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveStreamWatermarkRuleResponse
func (client *Client) AddLiveStreamWatermarkRuleWithContext(ctx context.Context, request *AddLiveStreamWatermarkRuleRequest, runtime *dara.RuntimeOptions) (_result *AddLiveStreamWatermarkRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLiveStreamWatermarkRule"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLiveStreamWatermarkRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a playlist item.
//
// Description:
//
// Create a production studio, add a production studio layout and production studio components, and then call this operation to add a playlist item. If no playlist has been created, the system automatically creates one. To create a production studio by using an API operation, see [CreateCaster](https://help.aliyun.com/document_detail/2848009.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - AddPlaylistItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPlaylistItemsResponse
func (client *Client) AddPlaylistItemsWithContext(ctx context.Context, request *AddPlaylistItemsRequest, runtime *dara.RuntimeOptions) (_result *AddPlaylistItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProgramConfig) {
		query["ProgramConfig"] = request.ProgramConfig
	}

	if !dara.IsNil(request.ProgramId) {
		query["ProgramId"] = request.ProgramId
	}

	if !dara.IsNil(request.ProgramItems) {
		query["ProgramItems"] = request.ProgramItems
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPlaylistItems"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPlaylistItemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a custom transcoding configuration for a streaming domain with Real-Time Streaming (RTS) enabled.
//
// Description:
//
// ## Usage notes
//
// You can call this operation to add a custom RTS transcoding configuration. This operation supports only the following types of custom transcoding templates: h264, h264-nbhd, h264-origin, and audio.
//
// ## QPS limits
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - AddRtsLiveStreamTranscodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRtsLiveStreamTranscodeResponse
func (client *Client) AddRtsLiveStreamTranscodeWithContext(ctx context.Context, request *AddRtsLiveStreamTranscodeRequest, runtime *dara.RuntimeOptions) (_result *AddRtsLiveStreamTranscodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.AudioBitrate) {
		query["AudioBitrate"] = request.AudioBitrate
	}

	if !dara.IsNil(request.AudioChannelNum) {
		query["AudioChannelNum"] = request.AudioChannelNum
	}

	if !dara.IsNil(request.AudioCodec) {
		query["AudioCodec"] = request.AudioCodec
	}

	if !dara.IsNil(request.AudioProfile) {
		query["AudioProfile"] = request.AudioProfile
	}

	if !dara.IsNil(request.AudioRate) {
		query["AudioRate"] = request.AudioRate
	}

	if !dara.IsNil(request.DeleteBframes) {
		query["DeleteBframes"] = request.DeleteBframes
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.FPS) {
		query["FPS"] = request.FPS
	}

	if !dara.IsNil(request.Gop) {
		query["Gop"] = request.Gop
	}

	if !dara.IsNil(request.Height) {
		query["Height"] = request.Height
	}

	if !dara.IsNil(request.Lazy) {
		query["Lazy"] = request.Lazy
	}

	if !dara.IsNil(request.Opus) {
		query["Opus"] = request.Opus
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Profile) {
		query["Profile"] = request.Profile
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.VideoBitrate) {
		query["VideoBitrate"] = request.VideoBitrate
	}

	if !dara.IsNil(request.Width) {
		query["Width"] = request.Width
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRtsLiveStreamTranscode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRtsLiveStreamTranscodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a show to a playlist.
//
// Description:
//
// Before calling this operation, create a playlist mode production studio (NormType=6) and add video resources to it. To create a production studio by using an API operation, refer to [CreateCaster](https://help.aliyun.com/document_detail/2848009.html).
//
// Each playlist can contain up to 1000 shows.
//
//	Notice:
//
// - When using video-on-demand (VOD) resources, use managed Bucket resources first. Resources in your own Bucket may expire. If you use resources in your own Bucket, check the resource validity period.
//
// - Use ApsaraVideo Live and ApsaraVideo VOD resources as input for the production studio. Resources from third-party URLs may fail to play. Verify the quality and validity of such resources.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - AddShowIntoShowListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddShowIntoShowListResponse
func (client *Client) AddShowIntoShowListWithContext(ctx context.Context, request *AddShowIntoShowListRequest, runtime *dara.RuntimeOptions) (_result *AddShowIntoShowListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.LiveInputType) {
		query["LiveInputType"] = request.LiveInputType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepeatTimes) {
		query["RepeatTimes"] = request.RepeatTimes
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ResourceUrl) {
		query["ResourceUrl"] = request.ResourceUrl
	}

	if !dara.IsNil(request.ShowName) {
		query["ShowName"] = request.ShowName
	}

	if !dara.IsNil(request.Spot) {
		query["Spot"] = request.Spot
	}

	if !dara.IsNil(request.IsBatchMode) {
		query["isBatchMode"] = request.IsBatchMode
	}

	if !dara.IsNil(request.ShowList) {
		query["showList"] = request.ShowList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddShowIntoShowList"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddShowIntoShowListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds layout settings for a virtual production studio.
//
// Description:
//
// You can call this operation to add layout settings for a virtual production studio. This operation supports both common and studio layouts.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 per user. If you exceed this limit, your API calls are throttled. This may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - AddStudioLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddStudioLayoutResponse
func (client *Client) AddStudioLayoutWithContext(ctx context.Context, request *AddStudioLayoutRequest, runtime *dara.RuntimeOptions) (_result *AddStudioLayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BgImageConfig) {
		query["BgImageConfig"] = request.BgImageConfig
	}

	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.CommonConfig) {
		query["CommonConfig"] = request.CommonConfig
	}

	if !dara.IsNil(request.LayerOrderConfigList) {
		query["LayerOrderConfigList"] = request.LayerOrderConfigList
	}

	if !dara.IsNil(request.LayoutName) {
		query["LayoutName"] = request.LayoutName
	}

	if !dara.IsNil(request.LayoutType) {
		query["LayoutType"] = request.LayoutType
	}

	if !dara.IsNil(request.MediaInputConfigList) {
		query["MediaInputConfigList"] = request.MediaInputConfigList
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScreenInputConfigList) {
		query["ScreenInputConfigList"] = request.ScreenInputConfigList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddStudioLayout"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddStudioLayoutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Inserts Supplemental Enhancement Information (SEI) data into transcoded streams.
//
// Description:
//
// First, obtain the streaming domain. Then, call this operation to insert SEI into the transcoded streams. The stream name must be the same as the source stream to ensure that SEI is inserted into all transcoded streams.
//
// ## QPS limit
//
// You can call this operation up to 6,000 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - AddTrancodeSEIRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTrancodeSEIResponse
func (client *Client) AddTrancodeSEIWithContext(ctx context.Context, request *AddTrancodeSEIRequest, runtime *dara.RuntimeOptions) (_result *AddTrancodeSEIResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Delay) {
		query["Delay"] = request.Delay
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Pattern) {
		query["Pattern"] = request.Pattern
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Repeat) {
		query["Repeat"] = request.Repeat
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.Text) {
		query["Text"] = request.Text
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTrancodeSEI"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTrancodeSEIResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the BanLiveMessageGroup operation to mute a user.
//
// Description:
//
// Before you call this operation, you must call [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html) to create an interactive message group.
//
// ## QPS limits
//
// Each user can call this operation up to 10 times per second. API calls that exceed this limit are throttled, which may impact your business. We recommend that you call this operation within this limit.
//
// @param tmpReq - BanLiveMessageGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BanLiveMessageGroupResponse
func (client *Client) BanLiveMessageGroupWithContext(ctx context.Context, tmpReq *BanLiveMessageGroupRequest, runtime *dara.RuntimeOptions) (_result *BanLiveMessageGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BanLiveMessageGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExceptUsers) {
		request.ExceptUsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExceptUsers, dara.String("ExceptUsers"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.ExceptUsersShrink) {
		query["ExceptUsers"] = request.ExceptUsersShrink
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BanLiveMessageGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BanLiveMessageGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the configurations of multiple domain names at a time.
//
// Description:
//
// Obtain the domain names for which you want to delete the configurations, and then call this operation to delete the configurations of these domain domains at a time.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - BatchDeleteLiveDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteLiveDomainConfigsResponse
func (client *Client) BatchDeleteLiveDomainConfigsWithContext(ctx context.Context, request *BatchDeleteLiveDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteLiveDomainConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteLiveDomainConfigs"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteLiveDomainConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Queries whether one or more users are online based on their UIDs
//
// Description:
//
// ## Usage notes
//
// This operation supports batch queries. You can query the online status of up to 20 users at a time.
//
// ## QPS limits
//
// The single-user queries per second (QPS) limit for this operation is 100 times/second. If you exceed this limit, API calls will be throttled, which may affect your business. You should make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - BatchGetOnlineUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetOnlineUsersResponse
func (client *Client) BatchGetOnlineUsersWithContext(ctx context.Context, request *BatchGetOnlineUsersRequest, runtime *dara.RuntimeOptions) (_result *BatchGetOnlineUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.UserIds) {
		body["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetOnlineUsers"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetOnlineUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures domain names in batches.
//
// Description:
//
// Obtain the ApsaraVideo Live domain names to configure, and then call this operation to configure domain names in batches.
//
// ## Rate limit
//
// The single-user QPS limit for this operation is 30 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - BatchSetLiveDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetLiveDomainConfigsResponse
func (client *Client) BatchSetLiveDomainConfigsWithContext(ctx context.Context, request *BatchSetLiveDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchSetLiveDomainConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Functions) {
		query["Functions"] = request.Functions
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSetLiveDomainConfigs"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSetLiveDomainConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the muting of a message group user.
//
// Description:
//
// ## QPS limits
//
// The single-user QPS limit for this operation is 100 times/second. If this limit is exceeded, API calls will be throttled, which may affect your business. You should make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - CancelMuteAllGroupUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelMuteAllGroupUserResponse
func (client *Client) CancelMuteAllGroupUserWithContext(ctx context.Context, request *CancelMuteAllGroupUserRequest, runtime *dara.RuntimeOptions) (_result *CancelMuteAllGroupUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BroadCastType) {
		body["BroadCastType"] = request.BroadCastType
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OperatorUserId) {
		body["OperatorUserId"] = request.OperatorUserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelMuteAllGroupUser"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelMuteAllGroupUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls CancelMuteGroupUser to batch unmute members in a message group.
//
// Description:
//
// ## QPS limits
//
// The QPS limit for this API is 100 queries per second (QPS) per user. If the limit is exceeded, API calls will be throttled, which may affect your business. You can call the API properly to avoid this issue. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param tmpReq - CancelMuteGroupUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelMuteGroupUserResponse
func (client *Client) CancelMuteGroupUserWithContext(ctx context.Context, tmpReq *CancelMuteGroupUserRequest, runtime *dara.RuntimeOptions) (_result *CancelMuteGroupUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CancelMuteGroupUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CancelMuteUserList) {
		request.CancelMuteUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CancelMuteUserList, dara.String("CancelMuteUserList"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BroadCastType) {
		body["BroadCastType"] = request.BroadCastType
	}

	if !dara.IsNil(request.CancelMuteUserListShrink) {
		body["CancelMuteUserList"] = request.CancelMuteUserListShrink
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OperatorUserId) {
		body["OperatorUserId"] = request.OperatorUserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelMuteGroupUser"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelMuteGroupUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a domain name in ApsaraVideo Live to a specified resource group.
//
// Description:
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - ChangeLiveDomainResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeLiveDomainResourceGroupResponse
func (client *Client) ChangeLiveDomainResourceGroupWithContext(ctx context.Context, request *ChangeLiveDomainResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeLiveDomainResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeLiveDomainResourceGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeLiveDomainResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether a user is in an interactive messaging group.
//
// Description:
//
// Before you call this operation, make sure that you have called the [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html) operation to create an interactive messaging group.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param tmpReq - CheckLiveMessageUsersInGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckLiveMessageUsersInGroupResponse
func (client *Client) CheckLiveMessageUsersInGroupWithContext(ctx context.Context, tmpReq *CheckLiveMessageUsersInGroupRequest, runtime *dara.RuntimeOptions) (_result *CheckLiveMessageUsersInGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CheckLiveMessageUsersInGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserIds) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, dara.String("UserIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.UserIdsShrink) {
		query["UserIds"] = request.UserIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckLiveMessageUsersInGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckLiveMessageUsersInGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether one or more specified users are online.
//
// Description:
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param tmpReq - CheckLiveMessageUsersOnlineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckLiveMessageUsersOnlineResponse
func (client *Client) CheckLiveMessageUsersOnlineWithContext(ctx context.Context, tmpReq *CheckLiveMessageUsersOnlineRequest, runtime *dara.RuntimeOptions) (_result *CheckLiveMessageUsersOnlineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CheckLiveMessageUsersOnlineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserIds) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, dara.String("UserIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.UserIdsShrink) {
		query["UserIds"] = request.UserIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckLiveMessageUsersOnline"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckLiveMessageUsersOnlineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the time shifting service for a specified domain name (application or live stream).
//
// Description:
//
// Obtain the streaming domain first, and then call this operation to disable the time shifting service for a specified domain name (application or live stream).
//
// Before calling this operation, call OpenLiveShift to enable the time shifting service. The AppName and StreamName specified when disabling the service must be the same as those specified when enabling the service. Wildcards (*) are supported.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - CloseLiveShiftRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseLiveShiftResponse
func (client *Client) CloseLiveShiftWithContext(ctx context.Context, request *CloseLiveShiftRequest, runtime *dara.RuntimeOptions) (_result *CloseLiveShiftResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseLiveShift"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseLiveShiftResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Duplicates a production studio and creates a new production studio instance.
//
// Description:
//
// You can call this operation to copy a specified production studio, which creates a new production studio instance.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - CopyCasterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyCasterResponse
func (client *Client) CopyCasterWithContext(ctx context.Context, request *CopyCasterRequest, runtime *dara.RuntimeOptions) (_result *CopyCasterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterName) {
		query["CasterName"] = request.CasterName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SrcCasterId) {
		query["SrcCasterId"] = request.SrcCasterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyCaster"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyCasterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies the configuration of a PVW scene to a PGM scene.
//
// Description:
//
// - Production Studio is billed based on output specifications, transcoding specifications, and the duration of usage. For more information, see [Production studio pricing](https://help.aliyun.com/document_detail/64531.html).
//
// - You can call this operation to copy the configuration from a source scene to a destination scene. You can only copy the configuration from a PVW scene to a PGM scene. A PVW scene is a preview scene, and a PGM scene is a program scene.
//
// - The PVW scene and the PGM scene must be in the same production studio.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second per user. API calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - CopyCasterSceneConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyCasterSceneConfigResponse
func (client *Client) CopyCasterSceneConfigWithContext(ctx context.Context, request *CopyCasterSceneConfigRequest, runtime *dara.RuntimeOptions) (_result *CopyCasterSceneConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.FromSceneId) {
		query["FromSceneId"] = request.FromSceneId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ToSceneId) {
		query["ToSceneId"] = request.ToSceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyCasterSceneConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyCasterSceneConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the CreateCaster operation to create a production studio.
//
// Description:
//
// You must activate ApsaraVideo Live before you call this operation. For more information, see [Activate ApsaraVideo Live](https://help.aliyun.com/document_detail/60361.html). <props="china">This operation supports the following types of production studios: Standard, Lightweight Carousel, Virtual Studio, and New Playlist (Carousel). <props="intl">This operation supports the following types of production studios: Standard and New Playlist (Carousel).
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If you exceed this limit, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - CreateCasterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCasterResponse
func (client *Client) CreateCasterWithContext(ctx context.Context, request *CreateCasterRequest, runtime *dara.RuntimeOptions) (_result *CreateCasterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterName) {
		query["CasterName"] = request.CasterName
	}

	if !dara.IsNil(request.CasterTemplate) {
		query["CasterTemplate"] = request.CasterTemplate
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.NormType) {
		query["NormType"] = request.NormType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PurchaseTime) {
		query["PurchaseTime"] = request.PurchaseTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCaster"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCasterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom stream mixing template.
//
// Description:
//
// After you call this operation to create a custom template, record the template name. To use the custom template, set the MixStreamTemplate parameter to the template name when you call the [CreateMixStream](https://help.aliyun.com/document_detail/2848087.html) operation to create a stream mixing task.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 for each user. API calls that exceed this limit are throttled, which can affect your business. Plan your calls accordingly.
//
// @param request - CreateCustomTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomTemplateResponse
func (client *Client) CreateCustomTemplateWithContext(ctx context.Context, request *CreateCustomTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomTemplate) {
		query["CustomTemplate"] = request.CustomTemplate
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomTemplate"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an edge transcoding job.
//
// Description:
//
// - You can call this operation to create an edge transcoding job.
//
// - Before you call this operation, you must have permission to access the edge transcoding service.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 6,000 calls per minute for each user. If you exceed this limit, API calls are throttled, which can affect your business. Plan your calls accordingly.
//
// @param request - CreateEdgeTranscodeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEdgeTranscodeJobResponse
func (client *Client) CreateEdgeTranscodeJobWithContext(ctx context.Context, request *CreateEdgeTranscodeJobRequest, runtime *dara.RuntimeOptions) (_result *CreateEdgeTranscodeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamInput) {
		query["StreamInput"] = request.StreamInput
	}

	if !dara.IsNil(request.StreamOutput) {
		query["StreamOutput"] = request.StreamOutput
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEdgeTranscodeJob"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEdgeTranscodeJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a callback for subscribing to channel messages.
//
// Description:
//
// Creates a callback for subscribing to channel messages. For example, when creating a callback, you can configure parameters such as the callback URL and event types.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - CreateEventSubRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEventSubResponse
func (client *Client) CreateEventSubWithContext(ctx context.Context, request *CreateEventSubRequest, runtime *dara.RuntimeOptions) (_result *CreateEventSubResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Events) {
		query["Events"] = request.Events
	}

	if !dara.IsNil(request.Users) {
		query["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEventSub"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEventSubResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a lightweight virtual studio template.
//
// Description:
//
//	Notice: The lightweight virtual studio feature is in public preview. Each user can create up to 300 templates. The feature is free of charge during the public preview. After the public preview ends, standard fees will apply. The specific date will be announced separately..
//
// @param tmpReq - CreateLiveAIStudioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLiveAIStudioResponse
func (client *Client) CreateLiveAIStudioWithContext(ctx context.Context, tmpReq *CreateLiveAIStudioRequest, runtime *dara.RuntimeOptions) (_result *CreateLiveAIStudioResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateLiveAIStudioShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MattingLayout) {
		request.MattingLayoutShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MattingLayout, dara.String("MattingLayout"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MediaLayout) {
		request.MediaLayoutShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MediaLayout, dara.String("MediaLayout"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BackgroundResourceId) {
		query["BackgroundResourceId"] = request.BackgroundResourceId
	}

	if !dara.IsNil(request.BackgroundResourceUrl) {
		query["BackgroundResourceUrl"] = request.BackgroundResourceUrl
	}

	if !dara.IsNil(request.BackgroundType) {
		query["BackgroundType"] = request.BackgroundType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Height) {
		query["Height"] = request.Height
	}

	if !dara.IsNil(request.MattingLayoutShrink) {
		query["MattingLayout"] = request.MattingLayoutShrink
	}

	if !dara.IsNil(request.MattingType) {
		query["MattingType"] = request.MattingType
	}

	if !dara.IsNil(request.MediaLayoutShrink) {
		query["MediaLayout"] = request.MediaLayoutShrink
	}

	if !dara.IsNil(request.MediaResourceId) {
		query["MediaResourceId"] = request.MediaResourceId
	}

	if !dara.IsNil(request.MediaResourceUrl) {
		query["MediaResourceUrl"] = request.MediaResourceUrl
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StudioName) {
		query["StudioName"] = request.StudioName
	}

	if !dara.IsNil(request.Width) {
		query["Width"] = request.Width
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLiveAIStudio"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLiveAIStudioResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a stream delay configuration.
//
// Description:
//
// Stream delay differs from latency caused by streaming protocols. Stream delay is a feature that lets you delay the playback of a live stream processed in the cloud.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 60 for each user. If you exceed the limit, API calls are throttled, which may affect your business. Call this operation within the specified limit.
//
// @param request - CreateLiveDelayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLiveDelayConfigResponse
func (client *Client) CreateLiveDelayConfigWithContext(ctx context.Context, request *CreateLiveDelayConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateLiveDelayConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.DelayTime) {
		query["DelayTime"] = request.DelayTime
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	if !dara.IsNil(request.TaskTriggerMode) {
		query["TaskTriggerMode"] = request.TaskTriggerMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLiveDelayConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLiveDelayConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an interactive messaging application by calling CreateLiveMessageApp.
//
// Description:
//
// - When calling other interactive messaging API operations, the data center must be the same as the one specified when creating the interactive messaging application.
//
// - A maximum of 300 interactive messaging applications can be created under a single Alibaba Cloud account.
//
// ## QPS limit
//
// The single-user QPS limit for this API operation is 50 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - CreateLiveMessageAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLiveMessageAppResponse
func (client *Client) CreateLiveMessageAppWithContext(ctx context.Context, request *CreateLiveMessageAppRequest, runtime *dara.RuntimeOptions) (_result *CreateLiveMessageAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AuditType) {
		query["AuditType"] = request.AuditType
	}

	if !dara.IsNil(request.AuditUrl) {
		query["AuditUrl"] = request.AuditUrl
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.EventCallbackUrl) {
		query["EventCallbackUrl"] = request.EventCallbackUrl
	}

	if !dara.IsNil(request.MsgLifeCycle) {
		query["MsgLifeCycle"] = request.MsgLifeCycle
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLiveMessageApp"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLiveMessageAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an interactive messaging group.
//
// Description:
//
//	  Before you call this operation, make sure that you have called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create an interactive messaging application.
//
//		- You can create up to 5,000 interactive messaging groups in an interactive messaging application.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param tmpReq - CreateLiveMessageGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLiveMessageGroupResponse
func (client *Client) CreateLiveMessageGroupWithContext(ctx context.Context, tmpReq *CreateLiveMessageGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateLiveMessageGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateLiveMessageGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Administrators) {
		request.AdministratorsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Administrators, dara.String("Administrators"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdministratorsShrink) {
		query["Administrators"] = request.AdministratorsShrink
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CreatorId) {
		query["CreatorId"] = request.CreatorId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupInfo) {
		query["GroupInfo"] = request.GroupInfo
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLiveMessageGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLiveMessageGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Alibaba Cloud Global Accelerator (GA) instance and attaches it to a live streaming link.
//
// Description:
//
// - This operation creates an Alibaba Cloud Global Accelerator (GA) instance and attaches it to a live streaming link. You must specify the stream-level granularity and indicate the acceleration start point and end point.
//
// - The template takes effect only when the AppName and StreamName values match the AppName and StreamName in the streaming URL.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If this limit is exceeded, the API invoke is throttled, which may affect your business. Invoke this operation appropriately.
//
// @param request - CreateLivePrivateLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLivePrivateLineResponse
func (client *Client) CreateLivePrivateLineWithContext(ctx context.Context, request *CreateLivePrivateLineRequest, runtime *dara.RuntimeOptions) (_result *CreateLivePrivateLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccelerationArea) {
		query["AccelerationArea"] = request.AccelerationArea
	}

	if !dara.IsNil(request.AccelerationType) {
		query["AccelerationType"] = request.AccelerationType
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxBandwidth) {
		query["MaxBandwidth"] = request.MaxBandwidth
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Reuse) {
		query["Reuse"] = request.Reuse
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.VideoCenter) {
		query["VideoCenter"] = request.VideoCenter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLivePrivateLine"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLivePrivateLineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call CreateLivePullToPush to create a pull-to-push task.
//
// Description:
//
//	Notice: Pull-to-push is a paid feature. Billing officially starts from 00:00 on December 5, 2025.
//
// - For pricing details, see [Pull-to-push pricing](https://help.aliyun.com/document_detail/2997901.html).
//
// - Call this operation to create a pull-to-push task.
//
// - Supports creating live stream pull tasks and VOD pull tasks.
//
// - After a task is created, it starts running at the specified start time and automatically stops and is deleted at the specified end time.
//
// - The push destination URL specified in the task must not be used by other tasks. Otherwise, multiple tasks pushing to the same URL simultaneously will cause push failures.
//
// - Pull-to-push callback events include task running status change callbacks and task exit callbacks. For more information, see [Pull-to-push event callbacks](https://help.aliyun.com/document_detail/2846768.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please call this operation appropriately.
//
// @param tmpReq - CreateLivePullToPushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLivePullToPushResponse
func (client *Client) CreateLivePullToPushWithContext(ctx context.Context, tmpReq *CreateLivePullToPushRequest, runtime *dara.RuntimeOptions) (_result *CreateLivePullToPushResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateLivePullToPushShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SourceUrls) {
		request.SourceUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceUrls, dara.String("SourceUrls"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthKey) {
		query["AuthKey"] = request.AuthKey
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.DstUrl) {
		query["DstUrl"] = request.DstUrl
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FileIndex) {
		query["FileIndex"] = request.FileIndex
	}

	if !dara.IsNil(request.NotifyItemSwitch) {
		query["NotifyItemSwitch"] = request.NotifyItemSwitch
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepeatNumber) {
		query["RepeatNumber"] = request.RepeatNumber
	}

	if !dara.IsNil(request.ReqAuth) {
		query["ReqAuth"] = request.ReqAuth
	}

	if !dara.IsNil(request.RetryCount) {
		query["RetryCount"] = request.RetryCount
	}

	if !dara.IsNil(request.RetryInterval) {
		query["RetryInterval"] = request.RetryInterval
	}

	if !dara.IsNil(request.SourceProtocol) {
		query["SourceProtocol"] = request.SourceProtocol
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.SourceUrlsShrink) {
		query["SourceUrls"] = request.SourceUrlsShrink
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLivePullToPush"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLivePullToPushResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls CreateLiveRealTimeLogDelivery to create a real-time log delivery configuration for a domain name.
//
// Description:
//
// Obtain the streaming domain first, and then call this operation to create a real-time log delivery configuration for the domain name.
//
// The resources corresponding to the Project, Logstore, and Region parameters must be created in Simple Log Service (SLS) in advance.
//
// Currently, only streaming domains can be configured. To push upstream real-time logs (that is, to configure an ingest domain), [submit a ticket](https://workorder.console.aliyun.com/console.htm#/ticket/add?productCode=live&commonQuestionId=4545&isSmart=true&iatraceid=1608439120675-2a5c48de0b84805313c708&channel=selfservice).
//
// Currently, only streaming domains can be configured. To push upstream real-time logs (that is, to configure an ingest domain), [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.12818093.nav-right.dticket.6cb216d07otFWR#/ticket/createIndex).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 6,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - CreateLiveRealTimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLiveRealTimeLogDeliveryResponse
func (client *Client) CreateLiveRealTimeLogDeliveryWithContext(ctx context.Context, request *CreateLiveRealTimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *CreateLiveRealTimeLogDeliveryResponse, _err error) {
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
		Action:      dara.String("CreateLiveRealTimeLogDelivery"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLiveRealTimeLogDeliveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a monitoring session.
//
// Description:
//
// Call this operation to create a monitoring session. Ensure that the required parameters are configured.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed the limit, your API calls are throttled. This may affect your business. Plan your API calls accordingly.
//
// @param request - CreateLiveStreamMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLiveStreamMonitorResponse
func (client *Client) CreateLiveStreamMonitorWithContext(ctx context.Context, request *CreateLiveStreamMonitorRequest, runtime *dara.RuntimeOptions) (_result *CreateLiveStreamMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.DingTalkWebHookUrl) {
		query["DingTalkWebHookUrl"] = request.DingTalkWebHookUrl
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.InputList) {
		query["InputList"] = request.InputList
	}

	if !dara.IsNil(request.MonitorConfig) {
		query["MonitorConfig"] = request.MonitorConfig
	}

	if !dara.IsNil(request.MonitorName) {
		query["MonitorName"] = request.MonitorName
	}

	if !dara.IsNil(request.OutputTemplate) {
		query["OutputTemplate"] = request.OutputTemplate
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLiveStreamMonitor"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLiveStreamMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an M3U8 index file for a specified time range.
//
// Description:
//
// You have configured OSS. For more information, see [Configure OSS](https://help.aliyun.com/document_detail/84932.html).
//
// Live recording indexing records a live video stream in M3U8 format, stores it in OSS, and performs real-time clipping on the stored TS segment index files.
//
// > - To create a recording index, the live stream must have had stream ingest activity. If no live streaming occurred within the specified time range or the stream name is incorrect, the recording index creation is failed.
//
// > - Make sure that DomainName, AppName, and StreamName are correct. Otherwise, the InvalidStream.NotFound error is returned.
//
// > - The interval between StartTime and EndTime must be at least the duration of one TS segment (30 seconds by default).
//
// > - EndTime must be later than StartTime, and the interval cannot exceed 4 days.
//
// > - TS segment information is retained in the ApsaraVideo Live system for only 3 months. You can create an M3U8 file only from recordings within the last 3 months.
//
// > - TS segment files are stored in OSS. The retention period is determined by the OSS storage configuration. For more information, see [Settings lifecycle rules](https://help.aliyun.com/document_detail/31904.html).
//
// > - Information about created M3U8 index files is retained in the ApsaraVideo Live system for only 6 months. You can query only the information of index files created within the last 6 months.
//
// > - M3U8 index files are stored in OSS. The retention period is determined by the OSS storage configuration.
//
// > - If the M3U8 and TS files are stored in different buckets, the TS paths in the M3U8 file are in HTTP format.
//
// ## QPS limit
//
// The single-user QPS limit for this API is 45 calls per second. If this limit is exceeded, throttling is triggered, which may affect your business. Call this operation as appropriate.
//
// @param request - CreateLiveStreamRecordIndexFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLiveStreamRecordIndexFilesResponse
func (client *Client) CreateLiveStreamRecordIndexFilesWithContext(ctx context.Context, request *CreateLiveStreamRecordIndexFilesRequest, runtime *dara.RuntimeOptions) (_result *CreateLiveStreamRecordIndexFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndTimeIncluded) {
		query["EndTimeIncluded"] = request.EndTimeIncluded
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssObject) {
		query["OssObject"] = request.OssObject
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLiveStreamRecordIndexFiles"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLiveStreamRecordIndexFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call CreateMessageApp to create an interactive message application.
//
// Description:
//
// ## QPS limits
//
// The QPS limit for this API is 100 queries per second (QPS) per user. If the limit is exceeded, API calls will be throttled, which may affect your business. You can call this API at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param tmpReq - CreateMessageAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMessageAppResponse
func (client *Client) CreateMessageAppWithContext(ctx context.Context, tmpReq *CreateMessageAppRequest, runtime *dara.RuntimeOptions) (_result *CreateMessageAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMessageAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AppConfig) {
		request.AppConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppConfig, dara.String("AppConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Extension) {
		request.ExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extension, dara.String("Extension"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppConfigShrink) {
		body["AppConfig"] = request.AppConfigShrink
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ExtensionShrink) {
		body["Extension"] = request.ExtensionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMessageApp"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMessageAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a message group.
//
// Description:
//
// ## QPS limits
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions. Consider this limit when calling this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param tmpReq - CreateMessageGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMessageGroupResponse
func (client *Client) CreateMessageGroupWithContext(ctx context.Context, tmpReq *CreateMessageGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateMessageGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMessageGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Extension) {
		request.ExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extension, dara.String("Extension"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CreatorId) {
		body["CreatorId"] = request.CreatorId
	}

	if !dara.IsNil(request.ExtensionShrink) {
		body["Extension"] = request.ExtensionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMessageGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMessageGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a stream mixing task.
//
// Description:
//
//	Notice:
//
// Cloud stream mixing is a paid feature. This feature is in public preview and is currently free of charge. Standard fees will apply after the public preview period. The end date of the public preview will be announced at a later date.
//
// </notice>
//
// You can call this operation to create a stream mixing task. This operation supports both preset and custom layouts.
//
// ## QPS limit
//
// A single user can make up to 10 queries per second (QPS). Calls that exceed this limit are throttled. Throttling may affect your business operations. We recommend that you plan your calls accordingly.
//
// @param request - CreateMixStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMixStreamResponse
func (client *Client) CreateMixStreamWithContext(ctx context.Context, request *CreateMixStreamRequest, runtime *dara.RuntimeOptions) (_result *CreateMixStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallbackConfig) {
		query["CallbackConfig"] = request.CallbackConfig
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.InputStreamList) {
		query["InputStreamList"] = request.InputStreamList
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.OutputConfig) {
		query["OutputConfig"] = request.OutputConfig
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMixStream"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMixStreamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a WHIP ingest URL for an ApsaraVideo Real-time Communication channel by calling CreateRTCWhipStreamAddress.
//
// Description:
//
// ## Usage notes
//
// 1. Call the CreateRTCWhipStreamAddress operation to create a WHIP ingest URL (WhipAddress) for the specified RTC channel.
//
// 2. Use OBS to ingest a stream by using the WHIP protocol.
//
//   - Run the OBS streaming tool.
//
//   - In the menu bar, choose **File > Settings**.
//
//   - On the Settings page, select **Stream**, configure the following information, and then click **OK**.
//
//     | Parameter | Description |
//
//     | ------ | ------|
//
//     | Service| Set Service to **WHIP**.|
//
//     | Server | Use the WhipAddress generated by the operation in Step 1.|
//
//     | Bearer Token| Leave the stream key empty. |
//
//     ![](https://img.alicdn.com/imgextra/i3/O1CN01xaAEK61umh8yP8NFe_!!6000000006080-2-tps-1746-685.png)
//
// 3. Other users can join the corresponding channel to watch the stream.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 40 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - CreateRTCWhipStreamAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRTCWhipStreamAddressResponse
func (client *Client) CreateRTCWhipStreamAddressWithContext(ctx context.Context, request *CreateRTCWhipStreamAddressRequest, runtime *dara.RuntimeOptions) (_result *CreateRTCWhipStreamAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRTCWhipStreamAddress"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRTCWhipStreamAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ingest URL for an RTMP stream in a channel.
//
// Description:
//
// 1. Call this operation to generate an RTMP ingest URL.
//
// 2. Ingest a stream to the RTMP URL. Other users in the channel can view this stream.
//
// 3. Stop stream ingest. Other users in the channel see the RTMP stream user leave the channel.
//
// > - You can repeat steps 2 and 3 while the RTMP URL is valid.
//
// >- Call [DescribeChannelParticipants](https://help.aliyun.com/document_detail/2848193.html) to query the channel user list and periodically check whether the RTMP stream user is still in the channel. If the user is no longer in the channel, stream ingest may have been interrupted. Stop stream ingest and go back to step 2 to re-ingest.
//
// ## Before you begin
//
// When using ApsaraVideo Real-time Communication, you typically join a channel and ingest RTC streams through the Alibaba Cloud ARTC SDK. However, in certain special scenarios, Alibaba Cloud allows you to use the RTMP protocol for stream ingest (such as OBS stream ingest). Alibaba Cloud then converts the RTMP stream to an RTC stream for distribution and user joining. You can use this operation to implement this capability. This operation generates an RTMP ingest URL. After you complete stream ingest, Alibaba Cloud automatically converts the stream to an RTC stream.
//
// If your business is a pure live streaming scenario, do not use this operation. Refer to [Generate ingest URLs and streaming URLs](https://help.aliyun.com/document_detail/198676.html) to quickly implement RTMP stream ingest and live playback.
//
// ## Rate limit
//
// The single-user queries per second (QPS) limit for this operation is 100. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - CreateRoomRealTimeStreamAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoomRealTimeStreamAddressResponse
func (client *Client) CreateRoomRealTimeStreamAddressWithContext(ctx context.Context, request *CreateRoomRealTimeStreamAddressRequest, runtime *dara.RuntimeOptions) (_result *CreateRoomRealTimeStreamAddressResponse, _err error) {
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
		Action:      dara.String("CreateRoomRealTimeStreamAddress"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRoomRealTimeStreamAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a real-time subtitle task to transcribe an audio stream into text in real time.
//
// Description:
//
// This operation is currently in maintenance mode. Use intelligent workflows to implement this capability. For more information, see [Intelligent workflow configuration](https://help.aliyun.com/document_detail/2985843.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 20 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - CreateRtcAsrTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRtcAsrTaskResponse
func (client *Client) CreateRtcAsrTaskWithContext(ctx context.Context, request *CreateRtcAsrTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateRtcAsrTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthKey) {
		query["AuthKey"] = request.AuthKey
	}

	if !dara.IsNil(request.AutoTerminateDelay) {
		query["AutoTerminateDelay"] = request.AutoTerminateDelay
	}

	if !dara.IsNil(request.AutoTerminateEnabled) {
		query["AutoTerminateEnabled"] = request.AutoTerminateEnabled
	}

	if !dara.IsNil(request.CallbackURL) {
		query["CallbackURL"] = request.CallbackURL
	}

	if !dara.IsNil(request.ChannelID) {
		query["ChannelID"] = request.ChannelID
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReportInterval) {
		query["ReportInterval"] = request.ReportInterval
	}

	if !dara.IsNil(request.RtcUserId) {
		query["RtcUserId"] = request.RtcUserId
	}

	if !dara.IsNil(request.SDKAppID) {
		query["SDKAppID"] = request.SDKAppID
	}

	if !dara.IsNil(request.StreamURL) {
		query["StreamURL"] = request.StreamURL
	}

	if !dara.IsNil(request.TargetLanguages) {
		query["TargetLanguages"] = request.TargetLanguages
	}

	if !dara.IsNil(request.TranslateEnabled) {
		query["TranslateEnabled"] = request.TranslateEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRtcAsrTask"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRtcAsrTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an event subscription for stream mixing and relaying.
//
// Description:
//
// Creates an event subscription for stream mixing and relaying. When you create a subscription, you can configure parameters such as the callback URL, the application to subscribe to, and channel information.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - CreateRtcMPUEventSubRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRtcMPUEventSubResponse
func (client *Client) CreateRtcMPUEventSubWithContext(ctx context.Context, request *CreateRtcMPUEventSubRequest, runtime *dara.RuntimeOptions) (_result *CreateRtcMPUEventSubResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.ChannelIds) {
		query["ChannelIds"] = request.ChannelIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRtcMPUEventSub"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRtcMPUEventSubResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a production studio.
//
// Description:
//
// - A production studio must be shut down before you can delete it. Otherwise, the operation fails.
//
// - When you delete a production studio, its associated scenes, components, and layouts are also deleted.
//
// - You cannot recover a deleted production studio.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 for each user. Calls that exceed this limit are throttled, which can affect your business. Plan your calls accordingly.
//
// @param request - DeleteCasterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCasterResponse
func (client *Client) DeleteCasterWithContext(ctx context.Context, request *DeleteCasterRequest, runtime *dara.RuntimeOptions) (_result *DeleteCasterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCaster"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCasterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a component from a production studio.
//
// Description:
//
// Call the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation to create a production studio. You can then call this operation to delete a component from the production studio.
//
// ## QPS limit
//
// This operation is limited to 10 queries per second (QPS) per user. Calls that exceed this limit are throttled, which may affect your business. We recommend that you plan your calls accordingly.
//
// @param request - DeleteCasterComponentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCasterComponentResponse
func (client *Client) DeleteCasterComponentWithContext(ctx context.Context, request *DeleteCasterComponentRequest, runtime *dara.RuntimeOptions) (_result *DeleteCasterComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.ComponentId) {
		query["ComponentId"] = request.ComponentId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCasterComponent"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCasterComponentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an episode from a production studio.
//
// Description:
//
// Before you call this operation, you must obtain the production studio ID and the episode ID.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 4. API calls that exceed this limit are throttled, which may impact your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DeleteCasterEpisodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCasterEpisodeResponse
func (client *Client) DeleteCasterEpisodeWithContext(ctx context.Context, request *DeleteCasterEpisodeRequest, runtime *dara.RuntimeOptions) (_result *DeleteCasterEpisodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.EpisodeId) {
		query["EpisodeId"] = request.EpisodeId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCasterEpisode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCasterEpisodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an episode list in a production studio.
//
// Description:
//
// First, call the [AddCasterEpisodeGroup](https://help.aliyun.com/document_detail/2848071.html) operation to add an episode list to a production studio. You can then call this operation to delete the episode list.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 4 per user. API calls that exceed this limit are throttled, which may affect your business. Call this operation within the specified limit.
//
// @param request - DeleteCasterEpisodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCasterEpisodeGroupResponse
func (client *Client) DeleteCasterEpisodeGroupWithContext(ctx context.Context, request *DeleteCasterEpisodeGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteCasterEpisodeGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProgramId) {
		query["ProgramId"] = request.ProgramId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCasterEpisodeGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCasterEpisodeGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a layout in a production studio.
//
// Description:
//
// Call the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation to create a production studio. You can then call this operation to delete a layout in the production studio.
//
// ## QPS limit
//
// This operation is limited to 10 queries per second (QPS) for each user. API calls that exceed this limit are throttled, which may affect your business. Call this operation at a reasonable rate.
//
// @param request - DeleteCasterLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCasterLayoutResponse
func (client *Client) DeleteCasterLayoutWithContext(ctx context.Context, request *DeleteCasterLayoutRequest, runtime *dara.RuntimeOptions) (_result *DeleteCasterLayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCasterLayout"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCasterLayoutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the episode list for carousel playback in a production studio.
//
// Description:
//
// First, call the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation to create a production studio. Then, call this operation to delete the program of the production studio.
//
// ## Usage limits
//
// You can call this operation up to 4 queries per second (QPS) per account. If you exceed this limit, API calls are throttled. This can affect your business. Plan your calls accordingly.
//
// @param request - DeleteCasterProgramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCasterProgramResponse
func (client *Client) DeleteCasterProgramWithContext(ctx context.Context, request *DeleteCasterProgramRequest, runtime *dara.RuntimeOptions) (_result *DeleteCasterProgramResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCasterProgram"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCasterProgramResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the scene configuration of a production studio.
//
// Description:
//
// Calls this operation to delete the scene configuration of a production studio. This operation currently supports the following scene configuration types: component configuration, layout configuration, and component and layout configuration.
//
// The scene specified by SceneId must be in the started state (Status=1). You can check the current state of the scene by using the Scene.Status field of the DescribeCasterScenes operation.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DeleteCasterSceneConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCasterSceneConfigResponse
func (client *Client) DeleteCasterSceneConfigWithContext(ctx context.Context, request *DeleteCasterSceneConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteCasterSceneConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCasterSceneConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCasterSceneConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an input source from a production studio.
//
// Description:
//
// You must first call the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation to create a production studio. You can then call this operation to remove a video resource from the production studio.
//
// ## QPS limit
//
// The limit for this operation is 10 queries per second (QPS) per user. Calls that exceed this limit are throttled, which may impact your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DeleteCasterVideoResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCasterVideoResourceResponse
func (client *Client) DeleteCasterVideoResourceWithContext(ctx context.Context, request *DeleteCasterVideoResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteCasterVideoResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCasterVideoResource"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCasterVideoResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Closes a channel, which causes all current members in the channel to leave, but does not affect subsequent reuse of the channel.
//
// Description:
//
// A channel is implicitly created when an RTC client SDK joins a session. You can call ListRTCLiveRooms to query existing channels.
//
// The queries per second (QPS) limit for a single user is 100. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DeleteChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChannelResponse
func (client *Client) DeleteChannelWithContext(ctx context.Context, request *DeleteChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChannel"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom stream mixing template.
//
// Description:
//
// Obtain the name of the custom stream mixing template to delete, and then call this operation.
//
// ## QPS limit
//
// You can call this operation up to 10 times per second per user. If you exceed this limit, your API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DeleteCustomTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomTemplateResponse
func (client *Client) DeleteCustomTemplateWithContext(ctx context.Context, request *DeleteCustomTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomTemplate"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an edge transcoding job.
//
// Description:
//
// - This operation deletes an edge transcoding job.
//
// - You must have permission to access the edge transcoding service to call this operation.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 6,000 calls per minute. If you exceed this limit, your API calls are throttled, which may affect your business. Call this operation at a reasonable rate.
//
// @param request - DeleteEdgeTranscodeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEdgeTranscodeJobResponse
func (client *Client) DeleteEdgeTranscodeJobWithContext(ctx context.Context, request *DeleteEdgeTranscodeJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteEdgeTranscodeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEdgeTranscodeJob"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEdgeTranscodeJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a callback that is used to subscribe to channel or user events.
//
// Description:
//
// Before you call this operation, make sure that you have called the [CreateEventSubscribe](https://help.aliyun.com/document_detail/2848209.html) operation to create a callback that is used to subscribe to channel or user events.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DeleteEventSubRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventSubResponse
func (client *Client) DeleteEventSubWithContext(ctx context.Context, request *DeleteEventSubRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventSubResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SubscribeId) {
		query["SubscribeId"] = request.SubscribeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventSub"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventSubResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a subtitle rule.
//
// Description:
//
// This operation deletes a specified subtitle rule.
//
//	Notice: The intelligent subtitling feature is currently in invitational preview. A single user can add up to 300 subtitle templates.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 60 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation as needed.
//
// @param request - DeleteLiveAIProduceRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveAIProduceRulesResponse
func (client *Client) DeleteLiveAIProduceRulesWithContext(ctx context.Context, request *DeleteLiveAIProduceRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveAIProduceRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RulesId) {
		query["RulesId"] = request.RulesId
	}

	if !dara.IsNil(request.SuffixName) {
		query["SuffixName"] = request.SuffixName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveAIProduceRules"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveAIProduceRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virtual studio template.
//
// Description:
//
// Before you delete a virtual studio template, you must detach all associated rules. Otherwise, an error is reported.
//
//	Notice:
//
// The lightweight virtual studio feature is in invitational preview. Each user can add a maximum of 300 templates.
//
// @param request - DeleteLiveAIStudioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveAIStudioResponse
func (client *Client) DeleteLiveAIStudioWithContext(ctx context.Context, request *DeleteLiveAIStudioRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveAIStudioResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StudioId) {
		query["StudioId"] = request.StudioId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveAIStudio"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveAIStudioResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified subtitle template.
//
// Description:
//
// You can call this operation to delete a specified caption template for a live stream.
//
//	Notice:
//
// The real-time caption feature is in invitational preview. Each user can add up to 300 caption templates.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 60 calls per second for each user. API calls that exceed this limit are throttled, which may affect your business. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DeleteLiveAISubtitleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveAISubtitleResponse
func (client *Client) DeleteLiveAISubtitleWithContext(ctx context.Context, request *DeleteLiveAISubtitleRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveAISubtitleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SubtitleId) {
		query["SubtitleId"] = request.SubtitleId
	}

	if !dara.IsNil(request.SubtitleName) {
		query["SubtitleName"] = request.SubtitleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveAISubtitle"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveAISubtitleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a recording configuration at the AppName level.
//
// Description:
//
// Obtain the main streaming domain, then call this operation to delete a recording configuration at the AppName level.
//
// ## QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DeleteLiveAppRecordConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveAppRecordConfigResponse
func (client *Client) DeleteLiveAppRecordConfigWithContext(ctx context.Context, request *DeleteLiveAppRecordConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveAppRecordConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveAppRecordConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveAppRecordConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the snapshot configuration for live streams in an application. The deletion takes effect after you restart stream ingest.
//
// Description:
//
// You can call this operation to delete the snapshot configuration for live streams in an application. The deletion takes effect after you restart stream ingest.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DeleteLiveAppSnapshotConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveAppSnapshotConfigResponse
func (client *Client) DeleteLiveAppSnapshotConfigWithContext(ctx context.Context, request *DeleteLiveAppSnapshotConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveAppSnapshotConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveAppSnapshotConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveAppSnapshotConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an audio moderation configuration.
//
// Description:
//
// - You can call this operation to delete the automated audio review configuration for a specified streaming domain.
//
// - The automated review feature is available only in specific regions. For supported regions, see [Service regions](https://help.aliyun.com/document_detail/193730.html).
//
// ## QPS limits
//
// You can call this operation up to 10 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLiveAudioAuditConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveAudioAuditConfigResponse
func (client *Client) DeleteLiveAudioAuditConfigWithContext(ctx context.Context, request *DeleteLiveAudioAuditConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveAudioAuditConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveAudioAuditConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveAudioAuditConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the callback configuration for automated audio review from a streaming domain.
//
// Description:
//
// Only some live centers support automated review. For supported regions, see [Service regions](https://help.aliyun.com/document_detail/193730.html).
//
// ## QPS limits
//
// You can call this operation up to 10 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLiveAudioAuditNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveAudioAuditNotifyConfigResponse
func (client *Client) DeleteLiveAudioAuditNotifyConfigWithContext(ctx context.Context, request *DeleteLiveAudioAuditNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveAudioAuditNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveAudioAuditNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveAudioAuditNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a live center stream relay configuration.
//
// Description:
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLiveCenterTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveCenterTransferResponse
func (client *Client) DeleteLiveCenterTransferWithContext(ctx context.Context, request *DeleteLiveCenterTransferRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveCenterTransferResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DstUrl) {
		query["DstUrl"] = request.DstUrl
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveCenterTransfer"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveCenterTransferResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a stream delay configuration.
//
// Description:
//
// The queries per second (QPS) limit for a single user is 60. If you exceed the limit, API calls are throttled. This may impact your business. Plan your API calls accordingly.
//
// @param request - DeleteLiveDelayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveDelayConfigResponse
func (client *Client) DeleteLiveDelayConfigWithContext(ctx context.Context, request *DeleteLiveDelayConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveDelayConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveDelayConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveDelayConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the configuration of callbacks for video moderation results.
//
// Description:
//
// ## [](#)Usage notes
//
// - Obtain the main streaming domain, and then call this operation to delete the configuration of callbacks for video moderation results.
//
// - Only some live centers support the content moderation feature. For more information, see [Supported regions](https://help.aliyun.com/document_detail/193730.html).
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DeleteLiveDetectNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveDetectNotifyConfigResponse
func (client *Client) DeleteLiveDetectNotifyConfigWithContext(ctx context.Context, request *DeleteLiveDetectNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveDetectNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveDetectNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveDetectNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes all records of the specified live streaming domain.
//
// Description:
//
// - A successful call to the DeleteLiveDomain operation deletes all records for the specified live streaming domain. Use this operation with caution.
//
// - To prevent the domain name from becoming inaccessible, you must restore its A record with your DNS service provider before you delete the domain name. If your domain name registrar is Alibaba Cloud, log on to the [Alibaba Cloud DNS console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdns.console.aliyun.com%2F%3Fspm%3Da2c4g.11186623.0.0.3cda841fcvk7Qs\\&lang=zh). Navigate to the **Public Zone*	- page. Find the domain name, click **Settings**, and change the CNAME record to an A record. If your domain name is registered with another registrar, perform a similar configuration with that registrar.
//
// - If you only need to disable a live streaming domain, call the [StopLiveDomain](https://help.aliyun.com/document_detail/2847799.html) operation.
//
// ## QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLiveDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveDomainResponse
func (client *Client) DeleteLiveDomainWithContext(ctx context.Context, request *DeleteLiveDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveDomain"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the mapping between a streaming domain and an ingest domain.
//
// Description:
//
// Make sure that the streaming domain has been mapped to the ingest domain before you call this operation to delete the mapping. For more information about how to map a streaming domain to an ingest domain, see [AddLiveDomainMapping](https://help.aliyun.com/document_detail/88782.html).
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DeleteLiveDomainMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveDomainMappingResponse
func (client *Client) DeleteLiveDomainMappingWithContext(ctx context.Context, request *DeleteLiveDomainMappingRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveDomainMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PullDomain) {
		query["PullDomain"] = request.PullDomain
	}

	if !dara.IsNil(request.PushDomain) {
		query["PushDomain"] = request.PushDomain
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveDomainMapping"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveDomainMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the mapping between a main streaming domain and a sub-streaming domain.
//
// Description:
//
// ## QPS limits
//
// You can call this operation up to 1,000 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLiveDomainPlayMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveDomainPlayMappingResponse
func (client *Client) DeleteLiveDomainPlayMappingWithContext(ctx context.Context, request *DeleteLiveDomainPlayMappingRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveDomainPlayMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayDomain) {
		query["PlayDomain"] = request.PlayDomain
	}

	if !dara.IsNil(request.PullDomain) {
		query["PullDomain"] = request.PullDomain
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveDomainPlayMapping"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveDomainPlayMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a live stream relay configuration.
//
// Description:
//
// Calling DeleteLiveEdgeTransfer only deletes the live edge relay configuration for a specific domain name. It does not affect live stream relays that are already in progress based on that configuration. The following table describes typical scenarios:
//
// | Scenario        | Analysis           | Result   |
//
// | -------------- | -------------- | ------ |
//
// | 1. You call SetLiveEdgeTransfer to configure live edge relay, start stream ingest, and then call DeleteLiveEdgeTransfer to delete the configuration during the stream.      | The live edge relay configuration exists when stream ingest starts.        | Stream ingest is not affected, and the live stream relay is not interrupted.  |
//
// | 2. You call DeleteLiveEdgeTransfer to delete the live edge relay configuration during stream ingest, stop the stream, and then restart stream ingest.       | The live edge relay configuration no longer exists when stream ingest restarts.       |   Live edge relay is not started.  |
//
// | 3. You start stream ingest after calling DeleteLiveEdgeTransfer. | The live edge relay configuration does not exist at this point.      |    Live edge relay is not started.  |
//
// ## QPS limit
//
// The single-user QPS limit for this API is 100 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DeleteLiveEdgeTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveEdgeTransferResponse
func (client *Client) DeleteLiveEdgeTransferWithContext(ctx context.Context, request *DeleteLiveEdgeTransferRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveEdgeTransferResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveEdgeTransfer"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveEdgeTransferResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a configuration of triggered stream pulling.
//
// Description:
//
// This operation supports deleting only the configurations of triggered stream pulling. If AppName is set to ali_all_app, such configurations for all applications under the domain name are deleted.
//
// ## QPS limits
//
// You can call this operation up to 1,000 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLiveLazyPullStreamInfoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveLazyPullStreamInfoConfigResponse
func (client *Client) DeleteLiveLazyPullStreamInfoConfigWithContext(ctx context.Context, request *DeleteLiveLazyPullStreamInfoConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveLazyPullStreamInfoConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveLazyPullStreamInfoConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveLazyPullStreamInfoConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an interactive messaging group.
//
// Description:
//
//	  Before you call this operation, make sure that you have called the [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html) operation to create an interactive messaging group.
//
//		- After you delete an interactive messaging group, it is no longer available. Every user in the group is notified that the group is closed.
//
//		- After you delete an interactive messaging group, messages in the group are retained for 30 days.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DeleteLiveMessageGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveMessageGroupResponse
func (client *Client) DeleteLiveMessageGroupWithContext(ctx context.Context, request *DeleteLiveMessageGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveMessageGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OperatorId) {
		query["OperatorId"] = request.OperatorId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveMessageGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveMessageGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a message that was sent to an interactive messaging group.
//
// Description:
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DeleteLiveMessageGroupMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveMessageGroupMessageResponse
func (client *Client) DeleteLiveMessageGroupMessageWithContext(ctx context.Context, request *DeleteLiveMessageGroupMessageRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveMessageGroupMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.DeleterId) {
		query["DeleterId"] = request.DeleterId
	}

	if !dara.IsNil(request.DeleterInfo) {
		query["DeleterInfo"] = request.DeleterInfo
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveMessageGroupMessage"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveMessageGroupMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes (recalls) a specific one-to-one chat message.
//
// Description:
//
// The single-user QPS limit for this operation is 50 calls per second. Exceeding this limit will trigger throttling, which may affect your business. Call this operation appropriately.
//
// @param request - DeleteLiveMessageUserMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveMessageUserMessageResponse
func (client *Client) DeleteLiveMessageUserMessageWithContext(ctx context.Context, request *DeleteLiveMessageUserMessageRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveMessageUserMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.DeleterId) {
		query["DeleterId"] = request.DeleterId
	}

	if !dara.IsNil(request.DeleterInfo) {
		query["DeleterInfo"] = request.DeleterInfo
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.ReceiverId) {
		query["ReceiverId"] = request.ReceiverId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveMessageUserMessage"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveMessageUserMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a live stream encapsulation configuration.
//
// Description:
//
// You can call this operation to delete a live stream encapsulation configuration. The change takes effect the next time you ingest the stream.
//
// ## QPS limits
//
// You can call this operation up to 300 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLivePackageConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLivePackageConfigResponse
func (client *Client) DeleteLivePackageConfigWithContext(ctx context.Context, request *DeleteLivePackageConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLivePackageConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLivePackageConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLivePackageConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds an Alibaba Cloud Global Accelerator (GA) instance from a live streaming link.
//
// Description:
//
// After unbinding, your live stream ingest and streaming links no longer use GA for back-to-origin. The accelerator instance still exists after unbinding. To release the accelerator instance, delete it in the Global Accelerator console.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - DeleteLivePrivateLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLivePrivateLineResponse
func (client *Client) DeleteLivePrivateLineWithContext(ctx context.Context, request *DeleteLivePrivateLineRequest, runtime *dara.RuntimeOptions) (_result *DeleteLivePrivateLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccelerationType) {
		query["AccelerationType"] = request.AccelerationType
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLivePrivateLine"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLivePrivateLineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scheduled stream pulling configuration.
//
// Description:
//
// This operation supports deleting only the configurations of scheduled stream pulling.
//
// ## QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLivePullStreamInfoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLivePullStreamInfoConfigResponse
func (client *Client) DeleteLivePullStreamInfoConfigWithContext(ctx context.Context, request *DeleteLivePullStreamInfoConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLivePullStreamInfoConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLivePullStreamInfoConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLivePullStreamInfoConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a stream pulling-to-pushing task by calling DeleteLivePullToPush.
//
// Description:
//
// - Deletes a stream pulling-to-pushing task.
//
// - Deleting a task cleans up the task. A running task is stopped immediately and cannot be restarted.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DeleteLivePullToPushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLivePullToPushResponse
func (client *Client) DeleteLivePullToPushWithContext(ctx context.Context, request *DeleteLivePullToPushRequest, runtime *dara.RuntimeOptions) (_result *DeleteLivePullToPushResponse, _err error) {
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
		Action:      dara.String("DeleteLivePullToPush"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLivePullToPushResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Logstore for real-time log delivery.
//
// Description:
//
// - You can call this operation to delete a Logstore used for real-time log delivery. Make sure that you set the parameters correctly.
//
// - You can call the [DescribeLiveDomainRealtimeLogDelivery](https://help.aliyun.com/document_detail/2848121.html) operation to query the Project, Logstore, and Region parameters.
//
// ## QPS limits
//
// The rate limit for this operation is 6,000 calls per minute for each user. If you exceed the limit, API calls are throttled. This may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DeleteLiveRealTimeLogLogstoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveRealTimeLogLogstoreResponse
func (client *Client) DeleteLiveRealTimeLogLogstoreWithContext(ctx context.Context, request *DeleteLiveRealTimeLogLogstoreRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveRealTimeLogLogstoreResponse, _err error) {
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
		Action:      dara.String("DeleteLiveRealTimeLogLogstore"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveRealTimeLogLogstoreResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a real-time log delivery configuration.
//
// Description:
//
// <props="china">
//
// - Real-time log delivery can be configured only for streaming domains. To configure real-time log delivery for ingest domains, [submit a ticket](https://workorder.console.aliyun.com/console.htm#/ticket/add?productCode=live\\&commonQuestionId=4545\\&isSmart=true\\&iatraceid=1608439120675-2a5c48de0b84805313c708\\&channel=selfservice).
//
// <props="intl">Real-time log delivery can be configured only for streaming domains. To configure real-time log delivery for ingest domains, [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.12818093.nav-right.dticket.6cb216d07otFWR#/ticket/createIndex).
//
// - You can call [DescribeLiveDomainRealtimeLogDelivery](https://help.aliyun.com/document_detail/2848121.html) to query the Project, Logstore, and Region parameters.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 6,000 calls per minute for each user. API calls that exceed this limit are throttled, which may affect your business. Call this operation at a reasonable rate.
//
// @param request - DeleteLiveRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveRealtimeLogDeliveryResponse
func (client *Client) DeleteLiveRealtimeLogDeliveryWithContext(ctx context.Context, request *DeleteLiveRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveRealtimeLogDeliveryResponse, _err error) {
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
		Action:      dara.String("DeleteLiveRealtimeLogDelivery"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the configuration of callbacks for live stream recording under a domain name.
//
// Description:
//
// Obtain the main streaming domain, and then call this operation to delete the configuration of callbacks for live stream recording under the main streaming domain.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DeleteLiveRecordNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveRecordNotifyConfigResponse
func (client *Client) DeleteLiveRecordNotifyConfigWithContext(ctx context.Context, request *DeleteLiveRecordNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveRecordNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveRecordNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveRecordNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Live-to-VOD configuration.
//
// Description:
//
// You can call this operation to delete the live-to-VOD configuration for a specified streaming domain.
//
// ## QPS limit
//
// You can call this operation up to 1,000 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLiveRecordVodConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveRecordVodConfigResponse
func (client *Client) DeleteLiveRecordVodConfigWithContext(ctx context.Context, request *DeleteLiveRecordVodConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveRecordVodConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveRecordVodConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveRecordVodConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a video moderation configuration.
//
// Description:
//
// - Obtain the main streaming domain, and then call this operation to delete a video moderation configuration.
//
// - Only some live centers support the content moderation feature. For more information, see [Supported regions](https://help.aliyun.com/document_detail/193730.html).
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DeleteLiveSnapshotDetectPornConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveSnapshotDetectPornConfigResponse
func (client *Client) DeleteLiveSnapshotDetectPornConfigWithContext(ctx context.Context, request *DeleteLiveSnapshotDetectPornConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveSnapshotDetectPornConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveSnapshotDetectPornConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveSnapshotDetectPornConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the snapshot callback configuration.
//
// Description:
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLiveSnapshotNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveSnapshotNotifyConfigResponse
func (client *Client) DeleteLiveSnapshotNotifyConfigWithContext(ctx context.Context, request *DeleteLiveSnapshotNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveSnapshotNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveSnapshotNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveSnapshotNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DeleteLiveSpecificStagingConfig operation to delete domain name configurations in a grayscale environment.
//
// Description:
//
// ## QPS limits
//
// You can call this operation up to 20 times per second per account. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// ## Queries Per Second (QPS) limits
//
// You can call this operation up to 20 times per second per account. If you exceed this limit, your API calls are throttled, which may impact your business. We recommend that you plan your calls accordingly. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DeleteLiveSpecificStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveSpecificStagingConfigResponse
func (client *Client) DeleteLiveSpecificStagingConfigWithContext(ctx context.Context, request *DeleteLiveSpecificStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveSpecificStagingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveSpecificStagingConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveSpecificStagingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes blocking configurations for a specific live stream.
//
// Description:
//
// You can call this operation to delete a stream-level block configuration.
//
// ## QPS limits
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLiveStreamBlockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveStreamBlockResponse
func (client *Client) DeleteLiveStreamBlockWithContext(ctx context.Context, request *DeleteLiveStreamBlockRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveStreamBlockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveStreamBlock"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveStreamBlockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an active merged stream.
//
// Description:
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLiveStreamMergeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveStreamMergeResponse
func (client *Client) DeleteLiveStreamMergeWithContext(ctx context.Context, request *DeleteLiveStreamMergeRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveStreamMergeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveStreamMerge"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveStreamMergeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a monitoring session.
//
// Description:
//
// Obtain the **MonitorId*	- from the response of the [CreateLiveStreamMonitor](https://help.aliyun.com/document_detail/2848129.html) operation. Then, call this operation to delete the monitoring session.
//
//	Notice:
//
// You cannot delete a monitoring session that is running. Attempting to do so returns a 400 error.
//
// ## QPS limits
//
// This operation is limited to 10 queries per second (QPS) per account. Calls that exceed this limit are throttled, which can affect your business. To avoid interruptions to your business, make sure that you do not exceed this limit.
//
// @param request - DeleteLiveStreamMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveStreamMonitorResponse
func (client *Client) DeleteLiveStreamMonitorWithContext(ctx context.Context, request *DeleteLiveStreamMonitorRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveStreamMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MonitorId) {
		query["MonitorId"] = request.MonitorId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveStreamMonitor"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveStreamMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the index files of live stream recordings.
//
// Description:
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLiveStreamRecordIndexFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveStreamRecordIndexFilesResponse
func (client *Client) DeleteLiveStreamRecordIndexFilesWithContext(ctx context.Context, request *DeleteLiveStreamRecordIndexFilesRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveStreamRecordIndexFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemoveFile) {
		query["RemoveFile"] = request.RemoveFile
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveStreamRecordIndexFiles"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveStreamRecordIndexFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a transcoding configuration.
//
// Description:
//
// Standard, Narrowband HD™, and custom transcoding templates are supported.
//
// ## QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLiveStreamTranscodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveStreamTranscodeResponse
func (client *Client) DeleteLiveStreamTranscodeWithContext(ctx context.Context, request *DeleteLiveStreamTranscodeRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveStreamTranscodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveStreamTranscode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveStreamTranscodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified watermark template.
//
// Description:
//
// This operation deletes a live streaming watermark template based on a specified template ID.
//
// ## QPS Limits
//
// You can call this operation up to 60 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLiveStreamWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveStreamWatermarkResponse
func (client *Client) DeleteLiveStreamWatermarkWithContext(ctx context.Context, request *DeleteLiveStreamWatermarkRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveStreamWatermarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveStreamWatermark"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveStreamWatermarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a watermark rule.
//
// Description:
//
// This operation deletes the watermark rule based on the specified rule ID.
//
// ## QPS limit
//
// You can call this operation up to 60 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteLiveStreamWatermarkRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveStreamWatermarkRuleResponse
func (client *Client) DeleteLiveStreamWatermarkRuleWithContext(ctx context.Context, request *DeleteLiveStreamWatermarkRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveStreamWatermarkRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveStreamWatermarkRule"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveStreamWatermarkRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the callback configuration of stream ingest for an ingest domain.
//
// Description:
//
// After you obtain an ingest domain, you can call this operation to delete the stream ingest callback configuration.
//
// ## QPS limits
//
// This operation is limited to 15 queries per second (QPS) for each user. If you exceed the limit, API calls are throttled, which can affect your business. Plan your calls accordingly.
//
// @param request - DeleteLiveStreamsNotifyUrlConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveStreamsNotifyUrlConfigResponse
func (client *Client) DeleteLiveStreamsNotifyUrlConfigWithContext(ctx context.Context, request *DeleteLiveStreamsNotifyUrlConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveStreamsNotifyUrlConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveStreamsNotifyUrlConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveStreamsNotifyUrlConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call DeleteMessageApp to delete an interactive message application.
//
// Description:
//
// ## QPS limits
//
// The QPS limit for this API is 100 queries per second (QPS) per user. API calls that exceed this limit are throttled, which may affect your business. You can call this API at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DeleteMessageAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMessageAppResponse
func (client *Client) DeleteMessageAppWithContext(ctx context.Context, request *DeleteMessageAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteMessageAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMessageApp"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMessageAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a stream mixing task.
//
// Description:
//
// Before you call this operation, you must create a stream mixing task by calling the [CreateMixStream](https://help.aliyun.com/document_detail/2848087.html) operation. If you no longer need the mixed stream, you must delete the task. Otherwise, the stream will be continuously ingested.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 for each user. API calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - DeleteMixStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMixStreamResponse
func (client *Client) DeleteMixStreamWithContext(ctx context.Context, request *DeleteMixStreamRequest, runtime *dara.RuntimeOptions) (_result *DeleteMixStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.MixStreamId) {
		query["MixStreamId"] = request.MixStreamId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMixStream"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMixStreamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an entire playlist.
//
// Description:
//
// After you add items to a playlist by calling the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation, you can call this operation to delete the entire playlist.
//
// ## QPS limit
//
// Each user can make up to 10 queries per second (QPS). If you exceed this limit, API calls are throttled, which can affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DeletePlaylistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePlaylistResponse
func (client *Client) DeletePlaylistWithContext(ctx context.Context, request *DeletePlaylistRequest, runtime *dara.RuntimeOptions) (_result *DeletePlaylistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProgramId) {
		query["ProgramId"] = request.ProgramId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePlaylist"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePlaylistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes items from a playlist.
//
// Description:
//
// You can call the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation to add items to a playlist. Then, you can call this operation to delete the items.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second per user. If you exceed this limit, your API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DeletePlaylistItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePlaylistItemsResponse
func (client *Client) DeletePlaylistItemsWithContext(ctx context.Context, request *DeletePlaylistItemsRequest, runtime *dara.RuntimeOptions) (_result *DeletePlaylistItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProgramId) {
		query["ProgramId"] = request.ProgramId
	}

	if !dara.IsNil(request.ProgramItemIds) {
		query["ProgramItemIds"] = request.ProgramItemIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePlaylistItems"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePlaylistItemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call `DeleteRtcAsrTask` to delete a real-time speech-to-text or translation task.
//
// Description:
//
// The call frequency for this API is limited to 20 queries per second (QPS) per user. If you exceed this limit, your API calls are throttled. This may impact your business. Plan your calls accordingly.
//
// @param request - DeleteRtcAsrTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRtcAsrTaskResponse
func (client *Client) DeleteRtcAsrTaskWithContext(ctx context.Context, request *DeleteRtcAsrTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteRtcAsrTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRtcAsrTask"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRtcAsrTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a subscription to mixed-stream relay events.
//
// Description:
//
// - You can call this operation to delete a subscription to mixed-stream relay events.
//
// - Before you call this operation, make sure that you have called the CreateRtcMPUEventSub operation to create the subscription.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DeleteRtcMPUEventSubRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRtcMPUEventSubResponse
func (client *Client) DeleteRtcMPUEventSubWithContext(ctx context.Context, request *DeleteRtcMPUEventSubRequest, runtime *dara.RuntimeOptions) (_result *DeleteRtcMPUEventSubResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRtcMPUEventSub"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRtcMPUEventSubResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the authentication configuration for snapshot callbacks.
//
// Description:
//
// After the deletion, callbacks for new streams will no longer be authenticated.
//
// ## QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteSnapshotCallbackAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotCallbackAuthResponse
func (client *Client) DeleteSnapshotCallbackAuthWithContext(ctx context.Context, request *DeleteSnapshotCallbackAuthRequest, runtime *dara.RuntimeOptions) (_result *DeleteSnapshotCallbackAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSnapshotCallbackAuth"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSnapshotCallbackAuthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes specified snapshot files.
//
// Description:
//
// This operation supports deleting snapshot files generated within a year.
//
// ## QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DeleteSnapshotFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotFilesResponse
func (client *Client) DeleteSnapshotFilesWithContext(ctx context.Context, request *DeleteSnapshotFilesRequest, runtime *dara.RuntimeOptions) (_result *DeleteSnapshotFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.CreateTimestampList) {
		query["CreateTimestampList"] = request.CreateTimestampList
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemoveFile) {
		query["RemoveFile"] = request.RemoveFile
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSnapshotFiles"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSnapshotFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a layout in a virtual studio.
//
// Description:
//
// You can call this operation to delete a layout from a production studio. You can delete only one layout in each call.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If the number of calls per second exceeds the limit, throttling is triggered. This may affect your business. Plan your calls accordingly.
//
// @param request - DeleteStudioLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStudioLayoutResponse
func (client *Client) DeleteStudioLayoutWithContext(ctx context.Context, request *DeleteStudioLayoutRequest, runtime *dara.RuntimeOptions) (_result *DeleteStudioLayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStudioLayout"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStudioLayoutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries scheduled tasks for automatic start and stop.
//
// Description:
//
// This operation queries information about scheduled tasks for automatic start and stop. When you call this operation, ensure that the parameters meet the requirements.
//
// ## QPS limit
//
// A single user can make a maximum of 10 queries per second (QPS). If you exceed this limit, API calls are throttled. This may affect your business. Therefore, call this operation at a reasonable rate.
//
// @param request - DescribeAutoShowListTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoShowListTasksResponse
func (client *Client) DescribeAutoShowListTasksWithContext(ctx context.Context, request *DescribeAutoShowListTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoShowListTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAutoShowListTasks"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoShowListTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation queries the channels of a production studio.
//
// Description:
//
// After you create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, you can call this operation to query its channels.
//
// ## QPS limit
//
// This operation is limited to 15 queries per second (QPS) for each user. If you exceed this limit, API calls are throttled, which can affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DescribeCasterChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCasterChannelsResponse
func (client *Client) DescribeCasterChannelsWithContext(ctx context.Context, request *DescribeCasterChannelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCasterChannelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCasterChannels"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCasterChannelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of components for a production studio.
//
// Description:
//
// Call the [AddCasterComponent](https://help.aliyun.com/document_detail/2848030.html) operation to add components to a production studio. You can then call this operation to query the list of components.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 15. If you exceed the limit, API calls are throttled. This may affect your business, so ensure that you call this operation within the limit.
//
// @param request - DescribeCasterComponentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCasterComponentsResponse
func (client *Client) DescribeCasterComponentsWithContext(ctx context.Context, request *DescribeCasterComponentsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCasterComponentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.ComponentId) {
		query["ComponentId"] = request.ComponentId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCasterComponents"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCasterComponentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration information of a production studio.
//
// Description:
//
// Create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, and then call this operation to query the configuration information of the production studio.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 15 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeCasterConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCasterConfigResponse
func (client *Client) DescribeCasterConfigWithContext(ctx context.Context, request *DescribeCasterConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeCasterConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCasterConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCasterConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the layouts of a production studio.
//
// Description:
//
// You can call this API operation to query the layouts of a production studio. If you do not specify a layout ID, all layouts of the studio are returned.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this API operation is 15 calls per second for each user. If you exceed this limit, API calls are throttled. This can affect your business, so you should call this API operation at a reasonable rate.
//
// @param request - DescribeCasterLayoutsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCasterLayoutsResponse
func (client *Client) DescribeCasterLayoutsWithContext(ctx context.Context, request *DescribeCasterLayoutsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCasterLayoutsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCasterLayouts"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCasterLayoutsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the program list of a production studio.
//
// Description:
//
// You can call this operation to query the program list of a production studio. This operation supports video source and component nodes.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 4 calls per second. If the limit is exceeded, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeCasterProgramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCasterProgramResponse
func (client *Client) DescribeCasterProgramWithContext(ctx context.Context, request *DescribeCasterProgramRequest, runtime *dara.RuntimeOptions) (_result *DescribeCasterProgramResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EpisodeId) {
		query["EpisodeId"] = request.EpisodeId
	}

	if !dara.IsNil(request.EpisodeType) {
		query["EpisodeType"] = request.EpisodeType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCasterProgram"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCasterProgramResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the audio configuration of a scenario.
//
// Description:
//
// Before you call this operation, you must create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 15 calls per second. API calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - DescribeCasterSceneAudioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCasterSceneAudioResponse
func (client *Client) DescribeCasterSceneAudioWithContext(ctx context.Context, request *DescribeCasterSceneAudioRequest, runtime *dara.RuntimeOptions) (_result *DescribeCasterSceneAudioResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCasterSceneAudio"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCasterSceneAudioResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of scenes in a production studio.
//
// Description:
//
// Create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, and then call this operation to query the list of scenes in the production studio.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 15 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation as needed.
//
// @param request - DescribeCasterScenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCasterScenesResponse
func (client *Client) DescribeCasterScenesWithContext(ctx context.Context, request *DescribeCasterScenesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCasterScenesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCasterScenes"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCasterScenesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the stream URL of a production studio.
//
// Description:
//
// Before you call this operation, call the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation to create a production studio. If a production studio has not been created, the InvalidScene.NotFound error is returned.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 5 for each user. Calls that exceed this limit are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeCasterStreamUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCasterStreamUrlResponse
func (client *Client) DescribeCasterStreamUrlWithContext(ctx context.Context, request *DescribeCasterStreamUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeCasterStreamUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCasterStreamUrl"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCasterStreamUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the video sources of a production studio.
//
// Description:
//
// Call the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation to create a production studio. You can then call this operation to query the video sources of the production studio.
//
// ## QPS limit
//
// This operation is limited to 15 queries per second (QPS) for each user. If you exceed this limit, your API calls are throttled. Throttling may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DescribeCasterVideoResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCasterVideoResourcesResponse
func (client *Client) DescribeCasterVideoResourcesWithContext(ctx context.Context, request *DescribeCasterVideoResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCasterVideoResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCasterVideoResources"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCasterVideoResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of production studios.
//
// Description:
//
// Create a production studio by calling the CreateCaster operation, and then call this operation to query the list of production studios. A production studio can be in the idle or streaming state.
//
// >
//
// > - The account that calls this operation must have [ApsaraVideo Live activated](https://live.console.aliyun.com/#/overview).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 15 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - DescribeCastersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCastersResponse
func (client *Client) DescribeCastersWithContext(ctx context.Context, request *DescribeCastersRequest, runtime *dara.RuntimeOptions) (_result *DescribeCastersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.CasterName) {
		query["CasterName"] = request.CasterName
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.NormType) {
		query["NormType"] = request.NormType
	}

	if !dara.IsNil(request.OrderByModifyAsc) {
		query["OrderByModifyAsc"] = request.OrderByModifyAsc
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCasters"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCastersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries online users in a channel.
//
// Description:
//
// You can call this operation to query online users in a channel. The returned result does not include details about the users.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DescribeChannelParticipantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelParticipantsResponse
func (client *Client) DescribeChannelParticipantsWithContext(ctx context.Context, request *DescribeChannelParticipantsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelParticipantsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelParticipants"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelParticipantsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of online users in a channel.
//
// Description:
//
// You can call this operation to query information about online users in a channel, such as the total number of users during live streaming.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DescribeChannelUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelUsersResponse
func (client *Client) DescribeChannelUsersWithContext(ctx context.Context, request *DescribeChannelUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelUsers"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage data of a domain name in a specific billing region.
//
// Description:
//
// - This operation supports batch domain name queries. Separate multiple domain names with commas (,). You can query up to 100 domain names at a time. If the DomainName parameter is empty, data for all domain names under the account is returned.
//
// - Usage data includes three types: traffic, bandwidth, and requests, measured in bytes, bit/s, and count respectively.
//
// - If you do not specify the Interval parameter, you can query data for up to the last year, and the maximum time span per query is 31 days. For a query period of 1 to 3 days, data is returned at hourly granularity. For a query period longer than 3 days, data is returned at daily granularity.
//
// - When you specify the Interval parameter, the supported maximum time span per query, the historical data range, and the data delay are as follows:
//
// | Time granularity | Maximum time span per query | Historical data range | Data delay |
//
// | -------------- | -------------- | ------ | ------ |
//
// | 5 minutes | 3 days | 93 days | 15 minutes |
//
// | 1 hour | 31 days | 186 days | 4 hours |
//
// | 1 day | 90 days | 366 days | 4:00 AM the next day |
//
// ## QPS limit
//
// The QPS limit for a single user on this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeDomainUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainUsageDataResponse
func (client *Client) DescribeDomainUsageDataWithContext(ctx context.Context, request *DescribeDomainUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainUsageDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.DataProtocol) {
		query["DataProtocol"] = request.DataProtocol
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Field) {
		query["Field"] = request.Field
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainUsageData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the integrity of a domain name.
//
// Description:
//
// You can call this operation to obtain the integrity of a domain name.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If you exceed this limit, API calls are throttled, which can affect your business. We recommend that you call this operation at a reasonable rate. For more information, see [QPS limit](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DescribeDomainWithIntegrityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainWithIntegrityResponse
func (client *Client) DescribeDomainWithIntegrityWithContext(ctx context.Context, request *DescribeDomainWithIntegrityRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainWithIntegrityResponse, _err error) {
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
		Action:      dara.String("DescribeDomainWithIntegrity"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainWithIntegrityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询HLS直播流的实时在线人数和带宽信息
//
// @param request - DescribeHlsLiveStreamRealTimeBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHlsLiveStreamRealTimeBpsDataResponse
func (client *Client) DescribeHlsLiveStreamRealTimeBpsDataWithContext(ctx context.Context, request *DescribeHlsLiveStreamRealTimeBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeHlsLiveStreamRealTimeBpsDataResponse, _err error) {
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
		Action:      dara.String("DescribeHlsLiveStreamRealTimeBpsData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHlsLiveStreamRealTimeBpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of caption rules.
//
// Description:
//
// This operation queries a list of caption rules. Before you call this operation, ensure that the parameters are correctly configured.
//
//	Notice:
//
// The real-time caption feature is in invitational preview. Each user can add a maximum of 300 caption templates.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 60. Exceeding this limit results in API call throttling, which can affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DescribeLiveAIProduceRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveAIProduceRulesResponse
func (client *Client) DescribeLiveAIProduceRulesWithContext(ctx context.Context, request *DescribeLiveAIProduceRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveAIProduceRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RulesId) {
		query["RulesId"] = request.RulesId
	}

	if !dara.IsNil(request.SuffixName) {
		query["SuffixName"] = request.SuffixName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveAIProduceRules"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveAIProduceRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the virtual studio templates in your account.
//
// @param request - DescribeLiveAIStudioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveAIStudioResponse
func (client *Client) DescribeLiveAIStudioWithContext(ctx context.Context, request *DescribeLiveAIStudioRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveAIStudioResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StudioId) {
		query["StudioId"] = request.StudioId
	}

	if !dara.IsNil(request.StudioName) {
		query["StudioName"] = request.StudioName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveAIStudio"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveAIStudioResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeLiveAISubtitle operation to query a list of caption templates.
//
// Description:
//
// This operation supports paging. You can specify paging parameters to query a list of caption templates. When you call this operation, ensure that the parameter settings are valid.
//
//	Notice:
//
// The real-time captioning feature is in invitational preview. Each user can add a maximum of 300 caption templates.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 60. If the number of calls per second exceeds the limit, throttling is triggered, which may affect your business. We recommend that you take note of the limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DescribeLiveAISubtitleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveAISubtitleResponse
func (client *Client) DescribeLiveAISubtitleWithContext(ctx context.Context, request *DescribeLiveAISubtitleRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveAISubtitleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SubtitleId) {
		query["SubtitleId"] = request.SubtitleId
	}

	if !dara.IsNil(request.SubtitleName) {
		query["SubtitleName"] = request.SubtitleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveAISubtitle"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveAISubtitleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the audio moderation configurations.
//
// Description:
//
// - You can call this API operation to query the audio moderation configurations for a specific streaming domain.
//
// - Automated review is supported only in some live centers. For supported regions, see [Service regions](https://help.aliyun.com/document_detail/193730.html).
//
// ## QPS limit
//
// You can call this operation up to 10 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveAudioAuditConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveAudioAuditConfigResponse
func (client *Client) DescribeLiveAudioAuditConfigWithContext(ctx context.Context, request *DescribeLiveAudioAuditConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveAudioAuditConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveAudioAuditConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveAudioAuditConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the callback configuration for audio moderation.
//
// Description:
//
// - You can call this operation to query the callback configuration for audio moderation for a specified streaming domain.
//
// - Automated review is available only in some live centers. For supported regions, see [Service regions](https://help.aliyun.com/document_detail/193730.html).
//
// ## QPS limit
//
// You can call this operation up to 10 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveAudioAuditNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveAudioAuditNotifyConfigResponse
func (client *Client) DescribeLiveAudioAuditNotifyConfigWithContext(ctx context.Context, request *DescribeLiveAudioAuditNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveAudioAuditNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveAudioAuditNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveAudioAuditNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询直播CDN诊断信息
//
// @param request - DescribeLiveCdnDiagnoseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveCdnDiagnoseInfoResponse
func (client *Client) DescribeLiveCdnDiagnoseInfoWithContext(ctx context.Context, request *DescribeLiveCdnDiagnoseInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveCdnDiagnoseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.App) {
		query["app"] = request.App
	}

	if !dara.IsNil(request.Domain) {
		query["domain"] = request.Domain
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.IntervalType) {
		query["intervalType"] = request.IntervalType
	}

	if !dara.IsNil(request.Phase) {
		query["phase"] = request.Phase
	}

	if !dara.IsNil(request.RequestType) {
		query["requestType"] = request.RequestType
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["streamName"] = request.StreamName
	}

	if !dara.IsNil(request.StreamSuffix) {
		query["streamSuffix"] = request.StreamSuffix
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveCdnDiagnoseInfo"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveCdnDiagnoseInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the audio and video frame rates and bitrates of a stream in a live center.
//
// Description:
//
// The time granularity for the returned data is 5 seconds. The maximum time range to query is 3 hours. You can query data in the last 30 days.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DescribeLiveCenterStreamRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveCenterStreamRateDataResponse
func (client *Client) DescribeLiveCenterStreamRateDataWithContext(ctx context.Context, request *DescribeLiveCenterStreamRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveCenterStreamRateDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveCenterStreamRateData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveCenterStreamRateDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the live center stream relay configuration by calling DescribeLiveCenterTransfer.
//
// Description:
//
// The queries per second (QPS) limit for a single user on this operation is 100. Requests that exceed this limit are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLiveCenterTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveCenterTransferResponse
func (client *Client) DescribeLiveCenterTransferWithContext(ctx context.Context, request *DescribeLiveCenterTransferRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveCenterTransferResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DstUrl) {
		query["DstUrl"] = request.DstUrl
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveCenterTransfer"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveCenterTransferResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a certificate.
//
// Description:
//
// Before you call this operation, get the certificate name from the [Certificates](https://help.aliyun.com/document_detail/2584962.html) page in the ApsaraVideo Live console.
//
// ## QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions. For more information, see [](t2136805.xdita#).
//
// @param request - DescribeLiveCertificateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveCertificateDetailResponse
func (client *Client) DescribeLiveCertificateDetailWithContext(ctx context.Context, request *DescribeLiveCertificateDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveCertificateDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveCertificateDetail"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveCertificateDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificates of one or more specified domain names or all certificates within your Alibaba Cloud account.
//
// Description:
//
// If you specify one or more domain names in the request, the certificates of the domain names are returned. If you do not specify a domain name in the request, all certificates within your Alibaba Cloud account are returned.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DescribeLiveCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveCertificateListResponse
func (client *Client) DescribeLiveCertificateListWithContext(ctx context.Context, request *DescribeLiveCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveCertificateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveCertificateList"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveCertificateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls DescribeLiveDelayConfig to query the stream delay configuration.
//
// Description:
//
// The number of queries per second (QPS) for this operation is limited to 60 per user. Calls that exceed this limit are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeLiveDelayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDelayConfigResponse
func (client *Client) DescribeLiveDelayConfigWithContext(ctx context.Context, request *DescribeLiveDelayConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDelayConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDelayConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDelayConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls DescribeLiveDelayedStreamingUsage to query delayed streaming duration usage data.
//
// Description:
//
// - Queries user delayed streaming service usage data. The default granularity is 1 hour.
//
// - Maximum query time span: 31 days.
//
// - Minimum query time granularity: 1 hour.
//
// - Maximum query time range: 31 days.
//
// ## QPS limit
//
// The single-user QPS limit for this API is 5 queries per second. If the limit is exceeded, API calls are throttled, which may affect your business. We recommend that you call this API at a reasonable frequency.
//
// @param request - DescribeLiveDelayedStreamingUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDelayedStreamingUsageResponse
func (client *Client) DescribeLiveDelayedStreamingUsageWithContext(ctx context.Context, request *DescribeLiveDelayedStreamingUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDelayedStreamingUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDelayedStreamingUsage"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDelayedStreamingUsageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of callbacks for video moderation results.
//
// Description:
//
// - Obtain the main streaming domain, and then call this operation to query the configuration of callbacks for video moderation results.
//
// - Only some live centers support the automated review feature. For more information, see [Supported regions](https://help.aliyun.com/document_detail/193730.html).
//
// ## QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit on an API operation in ApsaraVideo Live](https://www.alibabacloud.com/help/en/apsaravideo-live/latest/qps-limit-on-an-api-operation-in-apsaravideo-live).
//
// @param request - DescribeLiveDetectNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDetectNotifyConfigResponse
func (client *Client) DescribeLiveDetectNotifyConfigWithContext(ctx context.Context, request *DescribeLiveDetectNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDetectNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDetectNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDetectNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeLiveDetectPornData operation to query usage data for live stream moderation.
//
// Description:
//
// - The minimum data granularity is 5 minutes. If the `StartTime` parameter is empty, the service queries data from the last 24 hours by default.
//
// - You can query data from the last 90 days.
//
// - You can query network bandwidth data for each time interval.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 10. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeLiveDetectPornDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDetectPornDataResponse
func (client *Client) DescribeLiveDetectPornDataWithContext(ctx context.Context, request *DescribeLiveDetectPornDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDetectPornDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Fee) {
		query["Fee"] = request.Fee
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDetectPornData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDetectPornDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network bandwidth monitoring data of a live streaming domain.
//
// Description:
//
// - If you do not specify StartTime or EndTime, data of the last 24 hours is returned by default. You can also query data for a specific time range by specifying both StartTime and EndTime.
//
// - If you specify only StartTime without EndTime, data within 1 hour from StartTime is returned.
//
// - If you specify only EndTime without StartTime, data within 1 hour before EndTime is returned.
//
// - You can query data of up to the last 90 days.
//
// - This is a monitoring data API. The data collection and processing method differs from that used for billing. Do not use this API to calculate usage for billing reconciliation.
//
// ## QPS limit
//
// The single-user QPS limit for this API is 100 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this API at a reasonable frequency.
//
// @param request - DescribeLiveDomainBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainBpsDataResponse
func (client *Client) DescribeLiveDomainBpsDataWithContext(ctx context.Context, request *DescribeLiveDomainBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainBpsDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainBpsData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainBpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bandwidth and traffic data of a domain name by protocol by calling DescribeLiveDomainBpsDataByLayer.
//
// Description:
//
// The storage duration of data is up to 90 days.
//
// You can query data for multiple domain names at a time. Separate multiple domain names with commas (,). A maximum of 500 domain names are supported. Data for multiple domain names is returned as aggregation results.
//
// The time granularity of returned data varies based on the time range specified by **StartTime*	- and **EndTime**:
//
// - Time range ≤ 3 days: The time granularity is 5 minutes.
//
// - 3 days < time range ≤ 31 days: The time granularity is 1 hour.
//
// - Time range > 31 days: The time granularity is 1 day.
//
// >If neither **StartTime*	- nor **EndTime*	- is specified, data for the last 24 hours is returned by default.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 20 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Invoke this operation appropriately.
//
// @param request - DescribeLiveDomainBpsDataByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainBpsDataByLayerResponse
func (client *Client) DescribeLiveDomainBpsDataByLayerWithContext(ctx context.Context, request *DescribeLiveDomainBpsDataByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainBpsDataByLayerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.Layer) {
		query["Layer"] = request.Layer
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainBpsDataByLayer"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainBpsDataByLayerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves accelerated domain names based on certificate information.
//
// Description:
//
// Calls this operation to query accelerated domain names that match the specified certificate information.
//
// ## QPS limit
//
// The QPS limit for a single user is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DescribeLiveDomainByCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainByCertificateResponse
func (client *Client) DescribeLiveDomainByCertificateWithContext(ctx context.Context, request *DescribeLiveDomainByCertificateRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainByCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SSLPub) {
		query["SSLPub"] = request.SSLPub
	}

	if !dara.IsNil(request.SSLStatus) {
		query["SSLStatus"] = request.SSLStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainByCertificate"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainByCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the certificate information for a specified domain name.
//
// Description:
//
// You must specify a domain name whose certificate information you want to query.
//
// ## QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveDomainCertificateInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainCertificateInfoResponse
func (client *Client) DescribeLiveDomainCertificateInfoWithContext(ctx context.Context, request *DescribeLiveDomainCertificateInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainCertificateInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainCertificateInfo"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainCertificateInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more configurations of a domain name.
//
// Description:
//
// You can call this operation to query multiple configurations at a time.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DescribeLiveDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainConfigsResponse
func (client *Client) DescribeLiveDomainConfigsWithContext(ctx context.Context, request *DescribeLiveDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainConfigs"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic configuration information of a specified ingest domain or streaming domain.
//
// Description:
//
// This operation queries the basic configuration information of a specified ingest domain or streaming domain. The configuration of a newly added domain name takes several minutes to complete. You can query it after the configuration is complete.
//
// ## QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainDetailResponse
func (client *Client) DescribeLiveDomainDetailWithContext(ctx context.Context, request *DescribeLiveDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainDetail"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 直播离线日志查询地址
//
// @param request - DescribeLiveDomainEdgeLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainEdgeLogResponse
func (client *Client) DescribeLiveDomainEdgeLogWithContext(ctx context.Context, request *DescribeLiveDomainEdgeLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainEdgeLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainEdgeLog"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainEdgeLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the frame rate and bitrate data of streams for an ingest domain.
//
// Description:
//
// You can call this API to query the real-time bitrate and frame rate of a stream. This helps you monitor the quality of stream ingest. Data collection and statistics are subject to a delay. Query data that is at least 5 minutes old.
//
// > Use this API to replace the deprecated DescribeLiveStreamsFrameRateAndBitRateData API.
//
// ## QPS limits
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveDomainFrameRateAndBitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainFrameRateAndBitRateDataResponse
func (client *Client) DescribeLiveDomainFrameRateAndBitRateDataWithContext(ctx context.Context, request *DescribeLiveDomainFrameRateAndBitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainFrameRateAndBitRateDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QueryTime) {
		query["QueryTime"] = request.QueryTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainFrameRateAndBitRateData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainFrameRateAndBitRateDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the maximum numbers of ingested and transcoded streams for a streaming domain.
//
// Description:
//
// This operation supports only main streaming domains.
//
// ## QPS limit
//
// You can call this operation up to 5 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions. For more information, see [QPS limit](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DescribeLiveDomainLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainLimitResponse
func (client *Client) DescribeLiveDomainLimitWithContext(ctx context.Context, request *DescribeLiveDomainLimitRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainLimitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainLimit"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainLimitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the download URLs of raw access logs for a specified domain name in ApsaraVideo Live.
//
// Description:
//
// - Each API call supports querying offline logs for only a single domain name.
//
// - The optional parameters StartTime and EndTime must be specified together. After you specify the start time and end time, logs within the specified time range are queried.
//
// - If you do not specify StartTime and EndTime, log data from the past 24 hours is returned by default.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLiveDomainLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainLogResponse
func (client *Client) DescribeLiveDomainLogWithContext(ctx context.Context, request *DescribeLiveDomainLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainLog"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 直播离线日志扩展接口(大客定制)
//
// @param request - DescribeLiveDomainLogExTtlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainLogExTtlResponse
func (client *Client) DescribeLiveDomainLogExTtlWithContext(ctx context.Context, request *DescribeLiveDomainLogExTtlRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainLogExTtlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainLogExTtl"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainLogExTtlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the mapping between ingest domains, main streaming domains, and sub-streaming domains.
//
// Description:
//
// Before you call this operation, identify the ingest or streaming domain that you want to query.
//
// ## QPS limits
//
// You can call this operation up to 500 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveDomainMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainMappingResponse
func (client *Client) DescribeLiveDomainMappingWithContext(ctx context.Context, request *DescribeLiveDomainMappingRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainMappingResponse, _err error) {
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
		Action:      dara.String("DescribeLiveDomainMapping"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring usage data of a domain name by calling DescribeLiveDomainMonitoringUsageData.
//
// Description:
//
// - You can query multiple domain names at a time. Separate multiple domain names with commas (,).
//
// - You can query data of up to the last 90 days.
//
// - The time granularity for querying data is hour or day.
//
// - The maximum query time span is 31 days.
//
// ## QPS limit
//
// The QPS limit for a single user on this operation is 20 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLiveDomainMonitoringUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainMonitoringUsageDataResponse
func (client *Client) DescribeLiveDomainMonitoringUsageDataWithContext(ctx context.Context, request *DescribeLiveDomainMonitoringUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainMonitoringUsageDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainMonitoringUsageData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainMonitoringUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration status of dual-stream disaster recovery.
//
// @param request - DescribeLiveDomainMultiStreamConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainMultiStreamConfigResponse
func (client *Client) DescribeLiveDomainMultiStreamConfigWithContext(ctx context.Context, request *DescribeLiveDomainMultiStreamConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainMultiStreamConfigResponse, _err error) {
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
		Action:      dara.String("DescribeLiveDomainMultiStreamConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainMultiStreamConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the number of online viewers for all live streams on a specified domain.
//
// Description:
//
// This API only supports viewer counts for FLV, RTS, and RTMP streams. HLS streams are not supported. The data collection has a delay. For accuracy, query for data that is at least 5 minutes old.
//
// > This API replaces the deprecated `DescribeLiveStreamOnlineUserNum` endpoint.
//
// ## QPS limits
//
// You can call this operation up to 200 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveDomainOnlineUserNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainOnlineUserNumResponse
func (client *Client) DescribeLiveDomainOnlineUserNumWithContext(ctx context.Context, request *DescribeLiveDomainOnlineUserNumRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainOnlineUserNumResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QueryTime) {
		query["QueryTime"] = request.QueryTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainOnlineUserNum"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainOnlineUserNumResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the stream ingest status codes for an ingest domain within a specified time range.
//
// Description:
//
// - This operation is a monitoring data API. The data collection and processing method differs from billing and cannot be used to calculate usage for reconciliation purposes.
//
// - You can query data from the last 90 days.
//
// - Data is delayed by 3 to 5 minutes.
//
// ## QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLiveDomainPublishErrorCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainPublishErrorCodeResponse
func (client *Client) DescribeLiveDomainPublishErrorCodeWithContext(ctx context.Context, request *DescribeLiveDomainPublishErrorCodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainPublishErrorCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainPublishErrorCode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainPublishErrorCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network bandwidth monitoring data for stream ingest on an ingest domain by calling DescribeLiveDomainPushBpsData.
//
// Description:
//
// - The bandwidth data is measured in bit/s.
//
// - Batch domain name queries are supported. Separate multiple domain names with commas (,).
//
// - If you do not specify StartTime or EndTime, data from the last 24 hours is returned by default. You can also query data for a specific time range by specifying both StartTime and EndTime.
//
// - You can query data from the last 90 days.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLiveDomainPushBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainPushBpsDataResponse
func (client *Client) DescribeLiveDomainPushBpsDataWithContext(ctx context.Context, request *DescribeLiveDomainPushBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainPushBpsDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainPushBpsData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainPushBpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the push traffic monitoring data of an ingest domain by calling DescribeLiveDomainPushTrafficData.
//
// Description:
//
// - The unit of the traffic monitoring data is bytes.
//
// - Batch domain name queries are supported. Separate multiple domain names with commas (,).
//
// - If you do not specify StartTime or EndTime, data of the last 24 hours is returned by default. You can also query data for a specific time range by specifying both StartTime and EndTime.
//
// - You can query data of the last 90 days.
//
// ## QPS limit
//
// The QPS limit for a single user on this operation is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as appropriate.
//
// @param request - DescribeLiveDomainPushTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainPushTrafficDataResponse
func (client *Client) DescribeLiveDomainPushTrafficDataWithContext(ctx context.Context, request *DescribeLiveDomainPushTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainPushTrafficDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainPushTrafficData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainPushTrafficDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the page view (PV) and unique visitor (UV) data for a specified streaming domain.
//
// Description:
//
// - You can call this operation to view the geographic distribution of your video viewers, check access rankings for your accelerated domain names, and count the number of unique IP addresses that send requests to the domain name within a specific period.
//
// - If you do not specify the StartTime and EndTime parameters, data from the last 24 hours is queried by default. You can also query data for a specified time range.
//
// - You can query only one domain name at a time.
//
// - You can query data from the last 90 days.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 100. If you exceed this limit, API calls are throttled. This may affect your business. We recommend that you call this operation at a reasonable frequency.
//
// @param request - DescribeLiveDomainPvUvDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainPvUvDataResponse
func (client *Client) DescribeLiveDomainPvUvDataWithContext(ctx context.Context, request *DescribeLiveDomainPvUvDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainPvUvDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainPvUvData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainPvUvDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DescribeLiveDomainRealTimeBpsData operation to retrieve bandwidth data for a domain name at a 1-minute granularity.
//
// Description:
//
// - You can call this operation to query the traffic and bandwidth usage of a specified domain name.
//
// - You can query data from the last 7 days. The time range of a single query cannot exceed 24 hours.
//
// - If you do not specify StartTime and EndTime, data for the last hour is returned by default.
//
// - This operation provides monitoring data. The data collection and processing methods are different from those used for billing. You cannot use the data from this operation for billing reconciliation.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If you exceed this limit, your API calls are throttled. This may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DescribeLiveDomainRealTimeBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainRealTimeBpsDataResponse
func (client *Client) DescribeLiveDomainRealTimeBpsDataWithContext(ctx context.Context, request *DescribeLiveDomainRealTimeBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainRealTimeBpsDataResponse, _err error) {
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
		Action:      dara.String("DescribeLiveDomainRealTimeBpsData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainRealTimeBpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the proportions of HTTP status codes for an accelerated domain name at a 1-minute granularity.
//
// Description:
//
// The following table describes the data timestamp granularity supported by this operation:
//
// |Time granularity|Maximum time range per query|Historical data available|Data latency|
//
// |-----|------|-------|-------|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|Typically 4 hours|
//
// |1 day|Unlimited|366 days|After 04:00 on the next day|
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Invoke this operation as needed.
//
// @param request - DescribeLiveDomainRealTimeHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainRealTimeHttpCodeDataResponse
func (client *Client) DescribeLiveDomainRealTimeHttpCodeDataWithContext(ctx context.Context, request *DescribeLiveDomainRealTimeHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainRealTimeHttpCodeDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainRealTimeHttpCodeData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries real-time traffic monitoring data for live streaming domain names.
//
// Description:
//
// - Call this operation to query the real-time traffic and bandwidth of a streaming domain name for a specific region, carrier, and time period.
//
// - If you do not specify the StartTime and EndTime parameters, data from the last hour is queried by default. To query data for a specific time range, you must specify both the StartTime and EndTime parameters.
//
// - This operation returns monitoring data. This data is collected and processed differently from the data used for billing. Therefore, you cannot use the returned data for billing reconciliation.
//
// - You can query data from the last 90 days.
//
// ## QPS limits
//
// This operation has no queries per second (QPS) limit for a single user. You can call this operation as needed.
//
// @param request - DescribeLiveDomainRealTimeTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainRealTimeTrafficDataResponse
func (client *Client) DescribeLiveDomainRealTimeTrafficDataWithContext(ctx context.Context, request *DescribeLiveDomainRealTimeTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainRealTimeTrafficDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainRealTimeTrafficData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainRealTimeTrafficDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about real-time log delivery for a domain name.
//
// Description:
//
// You can call this operation to query the status of real-time log delivery for a domain name. Ensure that the parameter settings are valid.
//
// ## QPS limits
//
// This operation is limited to 6,000 queries per second (QPS) per user. Calls that exceed this limit are throttled, which can affect your business. Plan your calls accordingly.
//
// @param request - DescribeLiveDomainRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainRealtimeLogDeliveryResponse
func (client *Client) DescribeLiveDomainRealtimeLogDeliveryWithContext(ctx context.Context, request *DescribeLiveDomainRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainRealtimeLogDeliveryResponse, _err error) {
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
		Action:      dara.String("DescribeLiveDomainRealtimeLogDelivery"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of concurrent recording streams and container format conversion usage for ApsaraVideo Live.
//
// Description:
//
// - Collects statistics on daily concurrent recording streams and sampling data at different time intervals. You can use this operation to query the peak number of daily or monthly concurrent recording streams.
//
// - Time shifting streams are not counted as recording streams.
//
// - Supports domain-level queries and batch domain queries. Separate multiple domain names with commas (,).
//
// - Data granularity: 1 minute. Maximum query span: 24 hours. Maximum data retention: 60 days.
//
// - Data granularity: 1 hour. Maximum query span: 31 days. Maximum data retention: 180 days.
//
// - Data granularity: 1 day. Maximum query span: 90 days. Maximum data retention: 366 days.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLiveDomainRecordUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainRecordUsageDataResponse
func (client *Client) DescribeLiveDomainRecordUsageDataWithContext(ctx context.Context, request *DescribeLiveDomainRecordUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainRecordUsageDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainRecordUsageData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainRecordUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of snapshots captured for a live streaming domain.
//
// Description:
//
// - You can use this operation to obtain the total number of snapshots captured per day.
//
// - You can query data from the last 90 days.
//
// ## QPS limit
//
// This operation does not have a per-user QPS limit. Call this operation as needed.
//
// @param request - DescribeLiveDomainSnapshotDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainSnapshotDataResponse
func (client *Client) DescribeLiveDomainSnapshotDataWithContext(ctx context.Context, request *DescribeLiveDomainSnapshotDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainSnapshotDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainSnapshotData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainSnapshotDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the feature configurations of an accelerated domain name in the canary release environment.
//
// Description:
//
// ## Usage notes
//
// You can call this operation to query the staging environment configuration of a specified accelerated domain name.
//
// ## QPS limits
//
// The single-user limit for this API operation is 30 queries per second (QPS). If you exceed this limit, API calls are throttled. This may affect your business. Pace your calls to stay within the limit. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DescribeLiveDomainStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainStagingConfigResponse
func (client *Client) DescribeLiveDomainStagingConfigWithContext(ctx context.Context, request *DescribeLiveDomainStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainStagingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainStagingConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainStagingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the transcoding usage data of a domain name by calling DescribeLiveDomainStreamTranscodeData.
//
// Description:
//
// - You can use this operation to query network bandwidth data for each time interval.
//
// - Batch domain name queries are supported. Separate multiple domain names with commas (,).
//
// - You can query data from the last 90 days.
//
// - The data time granularity is hour or day.
//
// - For the billing tiers that correspond to different transcoding types and transcoding resolutions, see the billing tier description for different instance specifications in [Live stream transcoding billing](https://help.aliyun.com/document_detail/90424.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 20 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Invoke this operation appropriately.
//
// @param request - DescribeLiveDomainStreamTranscodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainStreamTranscodeDataResponse
func (client *Client) DescribeLiveDomainStreamTranscodeDataWithContext(ctx context.Context, request *DescribeLiveDomainStreamTranscodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainStreamTranscodeDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Precision) {
		query["Precision"] = request.Precision
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Split) {
		query["Split"] = request.Split
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainStreamTranscodeData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainStreamTranscodeDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries time shifting usage data for live streaming.
//
// Description:
//
// - Queries time shifting usage data for each time interval.
//
// - Retrieves data for up to the last 90 days.
//
// - The data interval is fixed at 1 hour.
//
// - The maximum data timestamp span for a single query is 31 days.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API invoke is throttled, which may affect your business. Invoke this operation as needed.
//
// @param request - DescribeLiveDomainTimeShiftDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainTimeShiftDataResponse
func (client *Client) DescribeLiveDomainTimeShiftDataWithContext(ctx context.Context, request *DescribeLiveDomainTimeShiftDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainTimeShiftDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainTimeShiftData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainTimeShiftDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries network traffic monitoring data for live streaming domains.
//
// Description:
//
// - If you do not specify StartTime or EndTime, data from the past 24 hours is returned by default.
//
// - This is a monitoring data API. The data collection and processing method differs from that used for billing. Do not use this API to calculate usage for billing reconciliation.
//
// ## QPS limit
//
// The single-user QPS limit for this API is 100 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this API at a reasonable frequency.
//
// @param request - DescribeLiveDomainTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainTrafficDataResponse
func (client *Client) DescribeLiveDomainTrafficDataWithContext(ctx context.Context, request *DescribeLiveDomainTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainTrafficDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainTrafficData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainTrafficDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询直播域名转码参数
//
// @param request - DescribeLiveDomainTranscodeParamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDomainTranscodeParamsResponse
func (client *Client) DescribeLiveDomainTranscodeParamsWithContext(ctx context.Context, request *DescribeLiveDomainTranscodeParamsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDomainTranscodeParamsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.App) {
		query["app"] = request.App
	}

	if !dara.IsNil(request.Pushdomain) {
		query["pushdomain"] = request.Pushdomain
	}

	if !dara.IsNil(request.TemplateName) {
		query["template_name"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDomainTranscodeParams"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDomainTranscodeParamsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries DRM usage data within a specified time range.
//
// Description:
//
// ### Operation description
//
// - You can query data from the last 90 days.
//
// - The maximum time span is 31 days.
//
// ### QPS limit
//
// Each user can make up to 20 queries per second (QPS). If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation properly. For more information, see [QPS limit](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DescribeLiveDrmUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveDrmUsageDataResponse
func (client *Client) DescribeLiveDrmUsageDataWithContext(ctx context.Context, request *DescribeLiveDrmUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveDrmUsageDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveDrmUsageData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveDrmUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the configuration of edge stream relay.
//
// Description:
//
// Get the ingest domain, then call this operation to query the configuration of edge stream relay.
//
// ## QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveEdgeTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveEdgeTransferResponse
func (client *Client) DescribeLiveEdgeTransferWithContext(ctx context.Context, request *DescribeLiveEdgeTransferRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveEdgeTransferResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveEdgeTransfer"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveEdgeTransferResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call DescribeLiveGrtnDuration to query co-hosting call duration usage data.
//
// Description:
//
// - Maximum query span: 31 days.
//
// - Minimum query granularity: 5 minutes.
//
// - Maximum query range: Data from the last 90 days.
//
// ## QPS limit
//
// The single-user QPS limit for this API is 10 requests per second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please call this API appropriately.
//
// @param request - DescribeLiveGrtnDurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveGrtnDurationResponse
func (client *Client) DescribeLiveGrtnDurationWithContext(ctx context.Context, request *DescribeLiveGrtnDurationRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveGrtnDurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveGrtnDuration"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveGrtnDurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all configured certificate information for the certificate service of a user.
//
// Description:
//
// Calls this operation to query all configured certificates and domain name information for the user in the certificate service.
//
// Before calling this operation, make sure that at least one live streaming domain under the account has HTTPS certificates enabled through SetLiveDomainCertificate. Otherwise, the error NoHttpsDomain(400) is returned.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLiveHttpsDomainListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveHttpsDomainListResponse
func (client *Client) DescribeLiveHttpsDomainListWithContext(ctx context.Context, request *DescribeLiveHttpsDomainListRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveHttpsDomainListResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveHttpsDomainList"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveHttpsDomainListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries metric data for ApsaraVideo Real-time Communication (ARTC).
//
// Description:
//
// - You can query data from the past 30 days. The query time range for a single request is limited to 24 hours.
//
// - The data granularity is 5 minutes.
//
// ## QPS limit
//
// The QPS limit for this operation is 50 requests per second per account. Exceeding this limit triggers throttling, which can disrupt your services.
//
// @param request - DescribeLiveInteractionMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveInteractionMetricDataResponse
func (client *Client) DescribeLiveInteractionMetricDataWithContext(ctx context.Context, request *DescribeLiveInteractionMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveInteractionMetricDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BeginTs) {
		query["BeginTs"] = request.BeginTs
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.Os) {
		query["Os"] = request.Os
	}

	if !dara.IsNil(request.TerminalType) {
		query["TerminalType"] = request.TerminalType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveInteractionMetricData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveInteractionMetricDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a specified IP address belongs to an Alibaba Cloud point of presence (POP).
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveIpInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveIpInfoResponse
func (client *Client) DescribeLiveIpInfoWithContext(ctx context.Context, request *DescribeLiveIpInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveIpInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IP) {
		query["IP"] = request.IP
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveIpInfo"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveIpInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the configurations of triggered stream pulling for a streaming domain.
//
// Description:
//
// This operation supports retrieving only the configurations of triggered stream pulling.
//
// ## QPS limit
//
// You can call this operation up to 1,000 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveLazyPullStreamConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveLazyPullStreamConfigResponse
func (client *Client) DescribeLiveLazyPullStreamConfigWithContext(ctx context.Context, request *DescribeLiveLazyPullStreamConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveLazyPullStreamConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveLazyPullStreamConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveLazyPullStreamConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Interactive Message App
//
// Description:
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DescribeLiveMessageAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveMessageAppResponse
func (client *Client) DescribeLiveMessageAppWithContext(ctx context.Context, request *DescribeLiveMessageAppRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveMessageAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveMessageApp"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveMessageAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a live interactive group.
//
// Description:
//
// Before calling this operation, you must have already called [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848162.html) to create an interactive messaging group.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLiveMessageGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveMessageGroupResponse
func (client *Client) DescribeLiveMessageGroupWithContext(ctx context.Context, request *DescribeLiveMessageGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveMessageGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveMessageGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveMessageGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the mute status of users.
//
// Description:
//
// Before you call this operation, make sure that you have called the [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html) operation to create an interactive messaging group.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DescribeLiveMessageGroupBandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveMessageGroupBandResponse
func (client *Client) DescribeLiveMessageGroupBandWithContext(ctx context.Context, request *DescribeLiveMessageGroupBandRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveMessageGroupBandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveMessageGroupBand"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveMessageGroupBandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the live packaging configuration under a streaming domain.
//
// Description:
//
// Obtain the primary streaming domain first, and then call this operation to query the live packaging configuration under the streaming domain.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 300 calls per minute. If the limit is exceeded, API calls will be throttled, which may affect your business. Please call this operation as appropriate.
//
// @param request - DescribeLivePackageConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLivePackageConfigResponse
func (client *Client) DescribeLivePackageConfigWithContext(ctx context.Context, request *DescribeLivePackageConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLivePackageConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLivePackageConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLivePackageConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available origin points for Global Accelerator links.
//
// Description:
//
// You can call this operation to query available Global Accelerator regions before you call the CreateLivePrivateLine operation. These regions can be used as origin points for acceleration.
//
// ## Limits
//
// You can make up to 50 queries per second (QPS) per user. If you exceed this limit, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DescribeLivePrivateLineAreasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLivePrivateLineAreasResponse
func (client *Client) DescribeLivePrivateLineAreasWithContext(ctx context.Context, request *DescribeLivePrivateLineAreasRequest, runtime *dara.RuntimeOptions) (_result *DescribeLivePrivateLineAreasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLivePrivateLineAreas"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLivePrivateLineAreasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attachment information between an Alibaba Cloud Global Accelerator (GA) instance and a live streaming link.
//
// Description:
//
// When the request parameter IsGaInstance is set to yes, the Alibaba Cloud Global Accelerator (GA) instance status is queried. When it is set to no, the attachment details between the GA instance and the live streaming link are queried.
//
// ## QPS limit
//
// The single-user QPS limit for this API is 50 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Invoke this operation appropriately.
//
// @param request - DescribeLivePrivateLineAvailGARequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLivePrivateLineAvailGAResponse
func (client *Client) DescribeLivePrivateLineAvailGAWithContext(ctx context.Context, request *DescribeLivePrivateLineAvailGARequest, runtime *dara.RuntimeOptions) (_result *DescribeLivePrivateLineAvailGAResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccelerationArea) {
		query["AccelerationArea"] = request.AccelerationArea
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.IsGaInstance) {
		query["IsGaInstance"] = request.IsGaInstance
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.VideoCenter) {
		query["VideoCenter"] = request.VideoCenter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLivePrivateLineAvailGA"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLivePrivateLineAvailGAResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls DescribeLiveProducerUsageData to query producer studio usage data.
//
// Description:
//
// The minimum data query granularity is 1 hour. The maximum query time span is 31 days. The maximum historical query range is the last 90 days.
//
// ## QPS limit
//
// The single-user QPS limit for this API operation is 5 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation properly.
//
// @param request - DescribeLiveProducerUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveProducerUsageDataResponse
func (client *Client) DescribeLiveProducerUsageDataWithContext(ctx context.Context, request *DescribeLiveProducerUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveProducerUsageDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Instance) {
		query["Instance"] = request.Instance
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.App) {
		query["app"] = request.App
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveProducerUsageData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveProducerUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the scheduled stream pulling configurations for a domain name.
//
// Description:
//
// This operation supports retrieving only the configurations of scheduled stream pulling.
//
// ## QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLivePullStreamConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLivePullStreamConfigResponse
func (client *Client) DescribeLivePullStreamConfigWithContext(ctx context.Context, request *DescribeLivePullStreamConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLivePullStreamConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLivePullStreamConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLivePullStreamConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a stream pulling-to-pushing task.
//
// Description:
//
// - Queries a stream pulling-to-pushing task.
//
// - You can query the configuration and status information of a task with a specified ID.
//
// ## QPS limit
//
// The single-user QPS limit for this API is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this API appropriately.
//
// @param request - DescribeLivePullToPushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLivePullToPushResponse
func (client *Client) DescribeLivePullToPushWithContext(ctx context.Context, request *DescribeLivePullToPushRequest, runtime *dara.RuntimeOptions) (_result *DescribeLivePullToPushResponse, _err error) {
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
		Action:      dara.String("DescribeLivePullToPush"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLivePullToPushResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries stream pulling and pushing tasks by using fuzzy match.
//
// Description:
//
// - Invoke this operation to query the list of stream pulling and stream ingest nodes.
//
// - Supports paging query of the node list, and fuzzy search by node ID, node name, and destination stream ingest URL.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Invoke this operation appropriately.
//
// @param request - DescribeLivePullToPushListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLivePullToPushListResponse
func (client *Client) DescribeLivePullToPushListWithContext(ctx context.Context, request *DescribeLivePullToPushListRequest, runtime *dara.RuntimeOptions) (_result *DescribeLivePullToPushListResponse, _err error) {
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
		Action:      dara.String("DescribeLivePullToPushList"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLivePullToPushListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the download URLs of offline logs for live stream relay by calling DescribeLivePushProxyLog.
//
// Description:
//
// - Data time granularity: 1 hour.
//
// - Maximum query range: data from the last 31 days.
//
// - If you do not specify StartTime or EndTime, this operation reads data from the last 24 hours by default. If you specify StartTime and EndTime, data is queried based on the specified time range.
//
// ## QPS limit
//
// The QPS limit for a single user on this operation is 100 calls per second. If the limit is exceeded, API invocations are throttled, which may affect your business. Invoke this operation as needed.
//
// @param request - DescribeLivePushProxyLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLivePushProxyLogResponse
func (client *Client) DescribeLivePushProxyLogWithContext(ctx context.Context, request *DescribeLivePushProxyLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeLivePushProxyLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLivePushProxyLog"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLivePushProxyLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage data of live center relay by calling DescribeLivePushProxyUsageData.
//
// Description:
//
// - Queries the usage data of live center relay.
//
// - Maximum query span: 31 days.
//
// - Minimum query granularity: 1 day.
//
// - Maximum query range: data from the last 90 days.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 5 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLivePushProxyUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLivePushProxyUsageDataResponse
func (client *Client) DescribeLivePushProxyUsageDataWithContext(ctx context.Context, request *DescribeLivePushProxyUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLivePushProxyUsageDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLivePushProxyUsageData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLivePushProxyUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes DescribeLiveRealtimeDeliveryAcc to query the number of real-time log delivery attempts.
//
// Description:
//
// - Queries statistics on real-time log delivery attempts, including the number of successful and failed log delivery attempts.
//
// - Supports queries by UID dimension.
//
// - The billable count includes both successful and failed log delivery attempts.
//
// ## QPS limit
//
// The single-user QPS limit for this API is 100 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Invoke this API appropriately.
//
// @param request - DescribeLiveRealtimeDeliveryAccRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveRealtimeDeliveryAccResponse
func (client *Client) DescribeLiveRealtimeDeliveryAccWithContext(ctx context.Context, request *DescribeLiveRealtimeDeliveryAccRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveRealtimeDeliveryAccResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.LogStore) {
		query["LogStore"] = request.LogStore
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Project) {
		query["Project"] = request.Project
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveRealtimeDeliveryAcc"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveRealtimeDeliveryAccResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorization status for real-time log delivery.
//
// Description:
//
// You can call this operation to query the authorization status for real-time log delivery.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 100 calls per second per user. If you exceed this limit, throttling is triggered, which may affect your business. Plan your API calls accordingly. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DescribeLiveRealtimeLogAuthorizedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveRealtimeLogAuthorizedResponse
func (client *Client) DescribeLiveRealtimeLogAuthorizedWithContext(ctx context.Context, request *DescribeLiveRealtimeLogAuthorizedRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveRealtimeLogAuthorizedResponse, _err error) {
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
		Action:      dara.String("DescribeLiveRealtimeLogAuthorized"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveRealtimeLogAuthorizedResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all recording configurations of apps under a live streaming domain.
//
// Description:
//
// Obtain the ingest domain first, and then call this operation to query all app recording configurations under a live streaming domain.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLiveRecordConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveRecordConfigResponse
func (client *Client) DescribeLiveRecordConfigWithContext(ctx context.Context, request *DescribeLiveRecordConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveRecordConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveRecordConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveRecordConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation queries the recording callback configuration for a domain name.
//
// Description:
//
// You can call this operation to query the recording callback configuration for a streaming domain name.
//
// ## QPS limit
//
// This operation supports up to 50 queries per second (QPS) per user. If you exceed this limit, your API calls are throttled. Throttling may affect your business operations, so make sure to stay within the specified limit.
//
// @param request - DescribeLiveRecordNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveRecordNotifyConfigResponse
func (client *Client) DescribeLiveRecordNotifyConfigWithContext(ctx context.Context, request *DescribeLiveRecordNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveRecordNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveRecordNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveRecordNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the callback records for live stream recordings stored in OSS.
//
// Description:
//
// The China site (Chinese) QPS limit for a single user on this operation is 100 calls per second. Exceeding this limit results in throttling, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLiveRecordNotifyRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveRecordNotifyRecordsResponse
func (client *Client) DescribeLiveRecordNotifyRecordsWithContext(ctx context.Context, request *DescribeLiveRecordNotifyRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveRecordNotifyRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveRecordNotifyRecords"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveRecordNotifyRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries live-to-VOD configurations of a streaming domain.
//
// Description:
//
// ### **QPS limit**
//
// This API is limited to 1,000 queries per minute for each account. If you exceed this limit, API calls are throttled, which can affect your business.
//
// @param request - DescribeLiveRecordVodConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveRecordVodConfigsResponse
func (client *Client) DescribeLiveRecordVodConfigsWithContext(ctx context.Context, request *DescribeLiveRecordVodConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveRecordVodConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveRecordVodConfigs"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveRecordVodConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cloud recording usage of ApsaraVideo Real-time Communication.
//
// Description:
//
// - Queries the recording length for each specification of cloud recording on a daily basis.
//
// - Supports queries at the ApsaraVideo Real-time Communication application granularity.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - DescribeLiveRtcRecordUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveRtcRecordUsageDataResponse
func (client *Client) DescribeLiveRtcRecordUsageDataWithContext(ctx context.Context, request *DescribeLiveRtcRecordUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveRtcRecordUsageDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.RecordMode) {
		query["RecordMode"] = request.RecordMode
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveRtcRecordUsageData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveRtcRecordUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the time-shifting configurations for a specified domain name.
//
// Description:
//
// After you obtain a streaming domain name, you can call this operation to query the time-shifting configurations for the specified domain name.
//
// ## QPS limit
//
// You can make up to 10 queries per second (QPS) for each user. If you exceed this limit, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DescribeLiveShiftConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveShiftConfigsResponse
func (client *Client) DescribeLiveShiftConfigsWithContext(ctx context.Context, request *DescribeLiveShiftConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveShiftConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveShiftConfigs"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveShiftConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the snapshot configurations under a streaming domain.
//
// Description:
//
// Obtain the streaming domain first, and then call this operation to query the snapshot configurations under the streaming domain.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - DescribeLiveSnapshotConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveSnapshotConfigResponse
func (client *Client) DescribeLiveSnapshotConfigWithContext(ctx context.Context, request *DescribeLiveSnapshotConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveSnapshotConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveSnapshotConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveSnapshotConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the content moderation configurations for live streaming.
//
// Description:
//
// - Obtain the streamer streaming domain first, and then invoke this operation to query the content moderation configurations for live streaming. This operation supports sorting in ascending and descending order.
//
// - Currently, only some live centers support intelligent content moderation for live streaming. For information about the live centers that support this feature, see [Service regions](https://help.aliyun.com/document_detail/193730.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation as needed.
//
// @param request - DescribeLiveSnapshotDetectPornConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveSnapshotDetectPornConfigResponse
func (client *Client) DescribeLiveSnapshotDetectPornConfigWithContext(ctx context.Context, request *DescribeLiveSnapshotDetectPornConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveSnapshotDetectPornConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveSnapshotDetectPornConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveSnapshotDetectPornConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration for snapshot callbacks.
//
// Description:
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveSnapshotNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveSnapshotNotifyConfigResponse
func (client *Client) DescribeLiveSnapshotNotifyConfigWithContext(ctx context.Context, request *DescribeLiveSnapshotNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveSnapshotNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveSnapshotNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveSnapshotNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authentication status of an active stream.
//
// Description:
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveStreamAuthCheckingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamAuthCheckingResponse
func (client *Client) DescribeLiveStreamAuthCheckingWithContext(ctx context.Context, request *DescribeLiveStreamAuthCheckingRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamAuthCheckingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamAuthChecking"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamAuthCheckingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the frame rates and bitrates of an RTMP live stream within a specified time range. This operation is applicable to retrieving historical data.
//
// Description:
//
// Queries the frame rates and bitrates of an RTMP live stream within a specified time range. This operation is applicable to retrieving historical data.
//
// ## QPS limit
//
// The maximum number of queries per second (QPS) per user for this operation is 50. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as appropriate.
//
// @param request - DescribeLiveStreamBitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamBitRateDataResponse
func (client *Client) DescribeLiveStreamBitRateDataWithContext(ctx context.Context, request *DescribeLiveStreamBitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamBitRateDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamBitRateData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamBitRateDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the real-time count of active original and transcoded streams for a specified streaming domain.
//
// Description:
//
// Before you call this operation, obtain the streaming domain name in the console. The returned stream count includes streams encoded in H.264 and H.265 formats.
//
// ## QPS limit
//
// You can call this operation only one time per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveStreamCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamCountResponse
func (client *Client) DescribeLiveStreamCountWithContext(ctx context.Context, request *DescribeLiveStreamCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamCountResponse, _err error) {
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
		Action:      dara.String("DescribeLiveStreamCount"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the live streaming latency.
//
// Description:
//
// You must first obtain a streaming domain name. You can then call this operation to query the live streaming latency.
//
// ## QPS limits
//
// This operation supports up to 1,000 queries per second (QPS) per user. If you exceed the limit, your calls are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - DescribeLiveStreamDelayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamDelayConfigResponse
func (client *Client) DescribeLiveStreamDelayConfigWithContext(ctx context.Context, request *DescribeLiveStreamDelayConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamDelayConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamDelayConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamDelayConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed audio and video frame rates and bitrates of an RTMP live stream.
//
// Description:
//
// - Call this operation to query a set of audio and video frame rates and bitrates of an RTMP live stream within a specified time range.
//
// - This operation is a monitoring data operation. The data collection and processing method differs from that used for billing. Do not use this operation to calculate usage for billing reconciliation.
//
// - You can query historical data within the last 90 days.
//
// - Data latency is 3 to 5 minutes.
//
// - The maximum time span for a single request is 1 hour.
//
// @param request - DescribeLiveStreamDetailFrameRateAndBitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamDetailFrameRateAndBitRateDataResponse
func (client *Client) DescribeLiveStreamDetailFrameRateAndBitRateDataWithContext(ctx context.Context, request *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamDetailFrameRateAndBitRateDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamDetailFrameRateAndBitRateData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamDetailFrameRateAndBitRateDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of historical online users for a live stream.
//
// Description:
//
// - The data returned by this operation is delayed for an average of 2 to 5 minutes.
//
// - This operation queries the number of historical online users for only Flash Video (FLV) and Real-Time Messaging Protocol (RTMP) streams.
//
// - This operation does not query the number of viewers that are watching transcoded streams.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DescribeLiveStreamHistoryUserNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamHistoryUserNumResponse
func (client *Client) DescribeLiveStreamHistoryUserNumWithContext(ctx context.Context, request *DescribeLiveStreamHistoryUserNumRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamHistoryUserNumResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamHistoryUserNum"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamHistoryUserNumResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query primary/backup stream merge configuration.
//
// Description:
//
// The QPS limit for a single user of this API is 100 calls per second. If this limit is exceeded, API calls will be throttled, which may affect your business. Please call this API appropriately.
//
// @param request - DescribeLiveStreamMergeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamMergeResponse
func (client *Client) DescribeLiveStreamMergeWithContext(ctx context.Context, request *DescribeLiveStreamMergeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamMergeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamMerge"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamMergeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries stream-level batch data for a specified domain name. A maximum of 5,000 rows of data can be returned per query.
//
// Description:
//
// If this is the first time you use this operation to query stream-level data for a specified domain name, <props="china">[submit a ticket](https://workorder.console.aliyun.com/console.htm#/ticket/add?productCode=live&commonQuestionId=4545&isSmart=true&iatraceid=1608439120675-2a5c48de0b84805313c708&channel=selfservice)<props="intl">[submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.12818093.nav-right.dticket.6cb216d07otFWR#/ticket/createIndex) to request backend configuration before using this operation.
//
// Provide the following information in the ticket:
//
// - The domain name to query.
//
// - The maximum number of concurrent live streams under the domain name.
//
// - The maximum number of concurrent viewers per live stream.
//
// - The protocol types included in client requests.
//
//	Notice: This operation will no longer be maintained after September 31, 2025. Switch to the new stream-level operation [DescribeLiveUserStreamMetricData](https://help.aliyun.com/document_detail/2948552.html) promptly. The new stream-level operation does not require backend configuration..
//
// ## Before you begin
//
// - Online viewer counting for HLS is not supported by default.
//
// - Only a single domain name can be queried at a time.
//
// - Maximum large query time span: 24 hours.
//
// - Minimum query granularity: 1 minute.
//
// - Maximum query range: 31 days.
//
// ## Rate limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLiveStreamMetricDetailDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamMetricDetailDataResponse
func (client *Client) DescribeLiveStreamMetricDetailDataWithContext(ctx context.Context, request *DescribeLiveStreamMetricDetailDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamMetricDetailDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamMetricDetailData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamMetricDetailDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of created monitoring sessions.
//
// Description:
//
// First, [create a monitoring session](https://help.aliyun.com/document_detail/2848129.html). You can then call this operation to query the list of monitoring sessions. When you call this operation, ensure that the required parameters are configured.
//
// ## QPS limit
//
// This operation is limited to 15 queries per second (QPS) for each user. API calls that exceed this limit are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeLiveStreamMonitorListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamMonitorListResponse
func (client *Client) DescribeLiveStreamMonitorListWithContext(ctx context.Context, request *DescribeLiveStreamMonitorListRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamMonitorListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MonitorId) {
		query["MonitorId"] = request.MonitorId
	}

	if !dara.IsNil(request.OrderRule) {
		query["OrderRule"] = request.OrderRule
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamMonitorList"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamMonitorListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Only data from the last 3 days can be queried.
//
// Description:
//
// This operation only supports querying data from the last 3 days.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLiveStreamPreloadTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamPreloadTasksResponse
func (client *Client) DescribeLiveStreamPreloadTasksWithContext(ctx context.Context, request *DescribeLiveStreamPreloadTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamPreloadTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlayUrl) {
		query["PlayUrl"] = request.PlayUrl
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamPreloadTasks"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamPreloadTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries stream ingest data at the app and stream granularity for a specified domain name.
//
// Description:
//
// - Only a single domain name can be queried. An error is returned if you specify multiple domain names.
//
// - Maximum query time range per request: 24 hours.
//
// - Minimum query granularity: 1 minute.
//
// - Maximum query period: 31 days.
//
// - This is a monitoring data API. The data collection and processing method differs from that used for billing. Do not use this API to calculate usage for billing reconciliation.
//
// ## QPS limit
//
// The maximum number of queries per second (QPS) per user for this operation is 10. If the number of calls per second exceeds the limit, throttling is triggered. This may affect your business. Call this operation as appropriate.
//
// @param request - DescribeLiveStreamPushMetricDetailDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamPushMetricDetailDataResponse
func (client *Client) DescribeLiveStreamPushMetricDetailDataWithContext(ctx context.Context, request *DescribeLiveStreamPushMetricDetailDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamPushMetricDetailDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamPushMetricDetailData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamPushMetricDetailDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the recordings of a live stream.
//
// Description:
//
// Obtain the main streaming domain, and then call this operation to query the recordings of a live stream.
//
// The information returned by this API is organized by recording task. To get information about specific recording files within a task, use the [DescribeLiveStreamRecordIndexFiles](https://help.aliyun.com/document_detail/2847890.html) and [DescribeLiveStreamRecordIndexFile](https://help.aliyun.com/document_detail/2847889.html) APIs.
//
// ## QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions.
//
// @param request - DescribeLiveStreamRecordContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamRecordContentResponse
func (client *Client) DescribeLiveStreamRecordContentWithContext(ctx context.Context, request *DescribeLiveStreamRecordContentRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamRecordContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamRecordContent"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamRecordContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a single recording manifest.
//
// Description:
//
// Metadata for created M3U8 manifests is retained in ApsaraVideo Live for 6 months. You can only query information about manifests created within this period. The M3U8 manifest files are stored in OSS, and their retention period is determined by your OSS storage configuration.
//
// ## QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions.
//
// @param request - DescribeLiveStreamRecordIndexFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamRecordIndexFileResponse
func (client *Client) DescribeLiveStreamRecordIndexFileWithContext(ctx context.Context, request *DescribeLiveStreamRecordIndexFileRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamRecordIndexFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamRecordIndexFile"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamRecordIndexFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all recording index files within a specified time period.
//
// Description:
//
// - Information about created M3U8 index files is retained in the ApsaraVideo Live system for only 6 months. You can query only the information about index files that were created within the last 6 months.
//
//   - M3U8 index files are stored in OSS. The retention period is determined by the storage configuration of OSS.
//
// - By default, recording returns HTTP URLs. To use HTTPS, configure the certificate and change HTTP to HTTPS.
//
// ## QPS limit
//
// The QPS limit on this API is 15 calls per second per user. If this limit is exceeded, API calls are throttled, which may affect your business. Call this API at an appropriate frequency.
//
// @param request - DescribeLiveStreamRecordIndexFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamRecordIndexFilesResponse
func (client *Client) DescribeLiveStreamRecordIndexFilesWithContext(ctx context.Context, request *DescribeLiveStreamRecordIndexFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamRecordIndexFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamRecordIndexFiles"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamRecordIndexFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries snapshot content within a specified time range.
//
// Description:
//
// You can call this operation to query snapshot data only within the last year.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLiveStreamSnapshotInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamSnapshotInfoResponse
func (client *Client) DescribeLiveStreamSnapshotInfoWithContext(ctx context.Context, request *DescribeLiveStreamSnapshotInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamSnapshotInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamSnapshotInfo"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamSnapshotInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time status of a single stream.
//
// Description:
//
// After you obtain a live streaming domain name, you can call this operation to query the real-time status of a single stream. For details about the offline status, use the data returned by the stream ingest callback. This operation does not provide a breakdown of the offline status.
//
// ## QPS limits
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveStreamStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamStateResponse
func (client *Client) DescribeLiveStreamStateWithContext(ctx context.Context, request *DescribeLiveStreamStateRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamState"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamStateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries transcoding configurations.
//
// Description:
//
// You must obtain the streaming domain before you call this operation to query the transcoding configurations.
//
// ## QPS limit
//
// You can call this operation up to 60 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveStreamTranscodeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamTranscodeInfoResponse
func (client *Client) DescribeLiveStreamTranscodeInfoWithContext(ctx context.Context, request *DescribeLiveStreamTranscodeInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamTranscodeInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainTranscodeName) {
		query["DomainTranscodeName"] = request.DomainTranscodeName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamTranscodeInfo"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamTranscodeInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries transcoding data for a specified domain name at the app and stream granularity.
//
// Description:
//
// - Maximum query span: 24 hours.
//
// - Minimum query granularity: 5 minutes.
//
// - Maximum query range: data from the last 31 days.
//
// ## QPS limit
//
// The single-user QPS limit for this API is 10 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DescribeLiveStreamTranscodeMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamTranscodeMetricDataResponse
func (client *Client) DescribeLiveStreamTranscodeMetricDataWithContext(ctx context.Context, request *DescribeLiveStreamTranscodeMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamTranscodeMetricDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamTranscodeMetricData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamTranscodeMetricDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of transcoding streams in real time by calling DescribeLiveStreamTranscodeStreamNum.
//
// Description:
//
// The QPS limit for a single user on this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DescribeLiveStreamTranscodeStreamNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamTranscodeStreamNumResponse
func (client *Client) DescribeLiveStreamTranscodeStreamNumWithContext(ctx context.Context, request *DescribeLiveStreamTranscodeStreamNumRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamTranscodeStreamNumResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SplitType) {
		query["SplitType"] = request.SplitType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamTranscodeStreamNum"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamTranscodeStreamNumResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of watermark rules.
//
// Description:
//
// This operation supports pagination.
//
// ## QPS limit
//
// You can call this operation up to 60 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveStreamWatermarkRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamWatermarkRulesResponse
func (client *Client) DescribeLiveStreamWatermarkRulesWithContext(ctx context.Context, request *DescribeLiveStreamWatermarkRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamWatermarkRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamWatermarkRules"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamWatermarkRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries watermark templates.
//
// Description:
//
// This operation supports paging.
//
// ## QPS limit
//
// You can call this operation up to 60 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveStreamWatermarksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamWatermarksResponse
func (client *Client) DescribeLiveStreamWatermarksWithContext(ctx context.Context, request *DescribeLiveStreamWatermarksRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamWatermarksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamWatermarks"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamWatermarksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the live streams that are blacklisted under a streaming domain.
//
// Description:
//
// The streaming URLs refer to the URLs for playback.
//
// ## QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveStreamsBlockListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamsBlockListResponse
func (client *Client) DescribeLiveStreamsBlockListWithContext(ctx context.Context, request *DescribeLiveStreamsBlockListRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamsBlockListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamsBlockList"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamsBlockListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operations performed on live streams for a specified domain name or application.
//
// Description:
//
// The operations include all API operations that were called on live streams.
//
// ## QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveStreamsControlHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamsControlHistoryResponse
func (client *Client) DescribeLiveStreamsControlHistoryWithContext(ctx context.Context, request *DescribeLiveStreamsControlHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamsControlHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamsControlHistory"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamsControlHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries stream callback records.
//
// Description:
//
// The QPS limit for a single user of this API operation is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - DescribeLiveStreamsNotifyRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamsNotifyRecordsResponse
func (client *Client) DescribeLiveStreamsNotifyRecordsWithContext(ctx context.Context, request *DescribeLiveStreamsNotifyRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamsNotifyRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamsNotifyRecords"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamsNotifyRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the callback configuration for stream ingest.
//
// Description:
//
// You can call this API to query the webhook address and authentication information for an ingest domain.
//
// ## QPS limits
//
// You can call this operation up to 1,000 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveStreamsNotifyUrlConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamsNotifyUrlConfigResponse
func (client *Client) DescribeLiveStreamsNotifyUrlConfigWithContext(ctx context.Context, request *DescribeLiveStreamsNotifyUrlConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamsNotifyUrlConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamsNotifyUrlConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamsNotifyUrlConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all active streams under a specified domain name or a specified application under a domain name.
//
// Description:
//
// This operation supports the following stream types:
//
// - all: Queries all streams.
//
// - raw: Queries raw streams.
//
// - trans: Queries transcoded streams.
//
// ## QPS limit
//
// The QPS limit for a single user is 10,000 calls per minute. If the limit is exceeded, the API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DescribeLiveStreamsOnlineListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamsOnlineListResponse
func (client *Client) DescribeLiveStreamsOnlineListWithContext(ctx context.Context, request *DescribeLiveStreamsOnlineListRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamsOnlineListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OnlyStream) {
		query["OnlyStream"] = request.OnlyStream
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamsOnlineList"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamsOnlineListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the historical stream ingest records of a domain name, an application under a domain name, or a specific stream within a specified time range.
//
// Description:
//
// You can call this operation to query historical streams within the last 30 days. The returned data contains the active stream information within the specified time range. This operation supports the following sorting methods:
//
// - stream_name_desc: sorts by live stream name in descending order.
//
// - stream_name_asc: sorts by live stream name in ascending order.
//
// - publish_time_desc: sorts by stream ingest time in descending order.
//
// - publish_time_asc: sorts by stream ingest time in ascending order.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 3 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeLiveStreamsPublishListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamsPublishListResponse
func (client *Client) DescribeLiveStreamsPublishListWithContext(ctx context.Context, request *DescribeLiveStreamsPublishListRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamsPublishListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamsPublishList"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamsPublishListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cumulative number of live streams at the day granularity by calling the DescribeLiveStreamsTotalCount operation.
//
// Description:
//
// - Maximum query time span: 15 days.
//
// - Maximum query time range: up to 1.5 years of historical data.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 100 calls per minute. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - DescribeLiveStreamsTotalCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveStreamsTotalCountResponse
func (client *Client) DescribeLiveStreamsTotalCountWithContext(ctx context.Context, request *DescribeLiveStreamsTotalCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveStreamsTotalCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Typ) {
		query["Typ"] = request.Typ
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveStreamsTotalCount"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveStreamsTotalCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags of ApsaraVideo Live resources.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
// @param request - DescribeLiveTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveTagResourcesResponse
func (client *Client) DescribeLiveTagResourcesWithContext(ctx context.Context, request *DescribeLiveTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveTagResources"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveTagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domains of a user ranked by traffic volume by calling DescribeLiveTopDomainsByFlow.
//
// Description:
//
// - If you do not specify StartTime or EndTime, data for the current month is returned by default. You can also query data for a specified time range. Both StartTime and EndTime must be specified together.
//
// - You can query data for up to 90 days.
//
// ## QPS limit
//
// The QPS limit for a single user on this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - DescribeLiveTopDomainsByFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveTopDomainsByFlowResponse
func (client *Client) DescribeLiveTopDomainsByFlowWithContext(ctx context.Context, request *DescribeLiveTopDomainsByFlowRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveTopDomainsByFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveTopDomainsByFlow"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveTopDomainsByFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取直播指定域名的原始访问日志的下载地址
//
// @param request - DescribeLiveTrafficDomainLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveTrafficDomainLogResponse
func (client *Client) DescribeLiveTrafficDomainLogWithContext(ctx context.Context, request *DescribeLiveTrafficDomainLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveTrafficDomainLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveTrafficDomainLog"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveTrafficDomainLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the audio and video parameters of an upstream ingest stream within a specified time range.
//
// Description:
//
// - The maximum time range for a query is 24 hours.
//
// - The minimum time range for a query is 1 minute.
//
// - You can query data from the last 31 days.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per minute. If you exceed this limit, API calls are throttled, which may affect your business. We recommend that you call this API at a reasonable rate.
//
// @param request - DescribeLiveUpVideoAudioInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveUpVideoAudioInfoResponse
func (client *Client) DescribeLiveUpVideoAudioInfoWithContext(ctx context.Context, request *DescribeLiveUpVideoAudioInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveUpVideoAudioInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveUpVideoAudioInfo"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveUpVideoAudioInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the bill prediction for a live streaming user.
//
// Description:
//
// This operation predicts your usage data for the current month. The prediction is based on the billing method that is active on the first day of the month. This operation supports predictions for monthly billing methods only and provides data at the user level. The query time range starts at 00:00 on the first day of the month and ends two hours before the current time.
//
// - Monthly 95th percentile: The highest data point after the top 5% of data points are removed from the specified time range.
//
// - Monthly average of daily peak bandwidth: The sum of daily peak bandwidth values divided by the number of days in the time range. The current day\\"s data is not included.
//
// - Monthly 4th peak: The fourth-highest peak bandwidth in the specified time range. If the time range is less than four days, the predicted value is 0.
//
// - Monthly average of daily 95th percentile peak: The sum of daily 95th percentile peak values divided by the number of days in the time range. The current day\\"s data is not included.
//
// - Nightly 95th percentile: The highest data point after the top 5% of data points are removed from the specified time range.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 1. If you exceed this limit, your API calls are throttled, which may affect your business. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DescribeLiveUserBillPredictionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveUserBillPredictionResponse
func (client *Client) DescribeLiveUserBillPredictionWithContext(ctx context.Context, request *DescribeLiveUserBillPredictionRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveUserBillPredictionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveUserBillPrediction"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveUserBillPredictionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries domain names of ApsaraVideo Live in your Alibaba Cloud account.
//
// Description:
//
// ## [](#)Usage notes
//
// You can call this operation to query all domain names of ApsaraVideo Live within your Alibaba Cloud account. The supported types of domain names are streaming domains and edge ingest domains.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DescribeLiveUserDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveUserDomainsResponse
func (client *Client) DescribeLiveUserDomainsWithContext(ctx context.Context, request *DescribeLiveUserDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveUserDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainSearchType) {
		query["DomainSearchType"] = request.DomainSearchType
	}

	if !dara.IsNil(request.DomainStatus) {
		query["DomainStatus"] = request.DomainStatus
	}

	if !dara.IsNil(request.LiveDomainType) {
		query["LiveDomainType"] = request.LiveDomainType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionName) {
		query["RegionName"] = request.RegionName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveUserDomains"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveUserDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries stream-level batch data for a specified streaming domain. Multiple protocols are supported.
//
// Description:
//
// > Querying new connections for the HLS protocol is not supported. Concurrent connections for HLS are counted based on requests in the default format. Requests in special formats require configuration by <props="china">[submit a ticket](https://workorder.console.aliyun.com/console.htm#/ticket/add?productCode=live&commonQuestionId=4545&isSmart=true&iatraceid=1608439120675-2a5c48de0b84805313c708&channel=selfservice)<props="intl">[submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.12818093.nav-right.dticket.6cb216d07otFWR#/ticket/createIndex). The default formats are as follows:
//
// > - m3u8 request example: http(s)://example.aliyundoc.com/Appname/ StreamName.m3u8
//
// > - ts request example: http(s)://example.aliyundoc.com/Appname/ StreamName/153xxxxxxxx_137xxxxx.ts.
//
// ## Settings
//
// - **Single query limit**: A maximum of 5000 rows of data can be returned per query.
//
// - **Domain name query limit**: Only a single domain name is supported. An error is returned if multiple domain names are specified.
//
// - **Time span limit**: The maximum query time span is 24 hours.
//
// - **Time granularity limit**: The minimum query granularity is 1 minute.
//
// - **Query range limit**: The maximum query range is 31 days.
//
// - **Call frequency limit**: The maximum call frequency per user is 10 calls per second.
//
// - **Special parameter combination**: When `DomainName` is not empty and both `AppName` and `StreamName` are set to `all`, the aggregate data for the streaming domain is returned.
//
// @param request - DescribeLiveUserStreamMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveUserStreamMetricDataResponse
func (client *Client) DescribeLiveUserStreamMetricDataWithContext(ctx context.Context, request *DescribeLiveUserStreamMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveUserStreamMetricDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveUserStreamMetricData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveUserStreamMetricDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves user tags.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DescribeLiveUserTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveUserTagsResponse
func (client *Client) DescribeLiveUserTagsWithContext(ctx context.Context, request *DescribeLiveUserTagsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveUserTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveUserTags"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveUserTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取直播指定域名的原始访问日志的下载地址
//
// @param request - DescribeLiveUserTrafficLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveUserTrafficLogResponse
func (client *Client) DescribeLiveUserTrafficLogWithContext(ctx context.Context, request *DescribeLiveUserTrafficLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveUserTrafficLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveUserTrafficLog"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveUserTrafficLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ownership verification content for a specified domain name.
//
// Description:
//
// - You can call this operation to query the verification content for a single domain name.
//
// - Each user can call this operation up to 30 times per second.
//
// - You must specify the domain name that you want to authenticate.
//
// - A successful call returns the verification content and a request ID for subsequent operations or for your records.
//
// @param request - DescribeLiveVerifyContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLiveVerifyContentResponse
func (client *Client) DescribeLiveVerifyContentWithContext(ctx context.Context, request *DescribeLiveVerifyContentRequest, runtime *dara.RuntimeOptions) (_result *DescribeLiveVerifyContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLiveVerifyContent"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLiveVerifyContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// DescribeMeterLiveBypassDuration.
//
// Description:
//
// ## Operation description
//
// - Maximum query span: 31 days.
//
// - Minimum query granularity: 5 minutes.
//
// - Maximum query range: data from the last 90 days.
//
// ## QPS limit
//
// A single user can call this operation up to 10 times per second. If the number of calls exceeds the limit, throttling is triggered, which may affect your business. Refer to [QPS limit](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DescribeMeterLiveBypassDurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMeterLiveBypassDurationResponse
func (client *Client) DescribeMeterLiveBypassDurationWithContext(ctx context.Context, request *DescribeMeterLiveBypassDurationRequest, runtime *dara.RuntimeOptions) (_result *DescribeMeterLiveBypassDurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMeterLiveBypassDuration"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMeterLiveBypassDurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries stream mixing tasks.
//
// Description:
//
// Call the [CreateMixStream](https://help.aliyun.com/document_detail/2848087.html) operation to create a stream mixing task. You can then call this operation to query the list of stream mixing tasks.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 5 for each user. API calls that exceed this limit are throttled. This can affect your business. Plan your API calls accordingly.
//
// @param request - DescribeMixStreamListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMixStreamListResponse
func (client *Client) DescribeMixStreamListWithContext(ctx context.Context, request *DescribeMixStreamListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMixStreamListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MixStreamId) {
		query["MixStreamId"] = request.MixStreamId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMixStreamList"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMixStreamListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the first frame latency composition within a specified time period.
//
// Description:
//
// Queries the first frame latency composition within a specified time range. QPS limit: The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency. For more information, see QPS limit.
//
// @param tmpReq - DescribeRTSNativeSDKFirstFrameCostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRTSNativeSDKFirstFrameCostResponse
func (client *Client) DescribeRTSNativeSDKFirstFrameCostWithContext(ctx context.Context, tmpReq *DescribeRTSNativeSDKFirstFrameCostRequest, runtime *dara.RuntimeOptions) (_result *DescribeRTSNativeSDKFirstFrameCostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeRTSNativeSDKFirstFrameCostShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DomainNameList) {
		request.DomainNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DomainNameList, dara.String("DomainNameList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataInterval) {
		query["DataInterval"] = request.DataInterval
	}

	if !dara.IsNil(request.DomainNameListShrink) {
		query["DomainNameList"] = request.DomainNameListShrink
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRTSNativeSDKFirstFrameCost"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRTSNativeSDKFirstFrameCostResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the average first frame delay within a specified time range.
//
// Description:
//
// Queries the first frame delay within a specified time range.
//
// # QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately. For more information, see QPS limit.
//
// @param tmpReq - DescribeRTSNativeSDKFirstFrameDelayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRTSNativeSDKFirstFrameDelayResponse
func (client *Client) DescribeRTSNativeSDKFirstFrameDelayWithContext(ctx context.Context, tmpReq *DescribeRTSNativeSDKFirstFrameDelayRequest, runtime *dara.RuntimeOptions) (_result *DescribeRTSNativeSDKFirstFrameDelayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeRTSNativeSDKFirstFrameDelayShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DomainNameList) {
		request.DomainNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DomainNameList, dara.String("DomainNameList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataInterval) {
		query["DataInterval"] = request.DataInterval
	}

	if !dara.IsNil(request.DomainNameListShrink) {
		query["DomainNameList"] = request.DomainNameListShrink
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRTSNativeSDKFirstFrameDelay"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRTSNativeSDKFirstFrameDelayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the reasons for playback failures (in the form of status codes) within a specified time period.
//
// Description:
//
// Queries the reasons for playback failures (in the form of status codes) within a specified time range. QPS limit: The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at an appropriate frequency. For more information, refer to QPS limits.
//
// @param tmpReq - DescribeRTSNativeSDKPlayFailStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRTSNativeSDKPlayFailStatusResponse
func (client *Client) DescribeRTSNativeSDKPlayFailStatusWithContext(ctx context.Context, tmpReq *DescribeRTSNativeSDKPlayFailStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeRTSNativeSDKPlayFailStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeRTSNativeSDKPlayFailStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DomainNameList) {
		request.DomainNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DomainNameList, dara.String("DomainNameList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataInterval) {
		query["DataInterval"] = request.DataInterval
	}

	if !dara.IsNil(request.DomainNameListShrink) {
		query["DomainNameList"] = request.DomainNameListShrink
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRTSNativeSDKPlayFailStatus"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRTSNativeSDKPlayFailStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the playback duration within a specified time range.
//
// Description:
//
// Queries the playback duration within a specified time range. QPS limit: The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Use this operation appropriately. For more information, refer to QPS limits.
//
// @param tmpReq - DescribeRTSNativeSDKPlayTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRTSNativeSDKPlayTimeResponse
func (client *Client) DescribeRTSNativeSDKPlayTimeWithContext(ctx context.Context, tmpReq *DescribeRTSNativeSDKPlayTimeRequest, runtime *dara.RuntimeOptions) (_result *DescribeRTSNativeSDKPlayTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeRTSNativeSDKPlayTimeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DomainNameList) {
		request.DomainNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DomainNameList, dara.String("DomainNameList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataInterval) {
		query["DataInterval"] = request.DataInterval
	}

	if !dara.IsNil(request.DomainNameListShrink) {
		query["DomainNameList"] = request.DomainNameListShrink
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRTSNativeSDKPlayTime"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRTSNativeSDKPlayTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the playback count within a specified time period, including the total count and the successful count.
//
// Description:
//
// # QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately. For more information, refer to QPS limit.
//
// @param tmpReq - DescribeRTSNativeSDKVvDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRTSNativeSDKVvDataResponse
func (client *Client) DescribeRTSNativeSDKVvDataWithContext(ctx context.Context, tmpReq *DescribeRTSNativeSDKVvDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeRTSNativeSDKVvDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeRTSNativeSDKVvDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DomainNameList) {
		request.DomainNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DomainNameList, dara.String("DomainNameList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataInterval) {
		query["DataInterval"] = request.DataInterval
	}

	if !dara.IsNil(request.DomainNameListShrink) {
		query["DomainNameList"] = request.DomainNameListShrink
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRTSNativeSDKVvData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRTSNativeSDKVvDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cloud recording files and task information for RTC.
//
// Description:
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeRtcCloudRecordingFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcCloudRecordingFilesResponse
func (client *Client) DescribeRtcCloudRecordingFilesWithContext(ctx context.Context, request *DescribeRtcCloudRecordingFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcCloudRecordingFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRtcCloudRecordingFiles"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRtcCloudRecordingFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a cloud transcoding task.
//
// @param request - DescribeRtcCloudTranscodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcCloudTranscodeResponse
func (client *Client) DescribeRtcCloudTranscodeWithContext(ctx context.Context, request *DescribeRtcCloudTranscodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcCloudTranscodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRtcCloudTranscode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRtcCloudTranscodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the event subscription for stream mixing and forwarding.
//
// Description:
//
// - This operation queries the event subscription for stream mixing and forwarding.
//
// - Before calling this operation, you must have already called CreateRtcMPUEventSub to create a stream mixing and forwarding event subscription.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeRtcMPUEventSubRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcMPUEventSubResponse
func (client *Client) DescribeRtcMPUEventSubWithContext(ctx context.Context, request *DescribeRtcMPUEventSubRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcMPUEventSubResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRtcMPUEventSub"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRtcMPUEventSubResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a playlist.
//
// Description:
//
// Before calling this operation, you must add a show to the playlist. To add a show by calling an API operation, see [Add a show to a playlist](https://help.aliyun.com/document_detail/2848051.html).
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled. This may affect your business operations. Plan your calls accordingly.
//
// @param request - DescribeShowListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeShowListResponse
func (client *Client) DescribeShowListWithContext(ctx context.Context, request *DescribeShowListRequest, runtime *dara.RuntimeOptions) (_result *DescribeShowListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeShowList"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeShowListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries region blocking configurations of a specific live stream.
//
// Description:
//
// ## QPS limit
//
// You can call this operation up to 50 times per second per user. If the number of calls per second exceeds the limit, throttling is triggered. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeStreamLocationBlockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamLocationBlockResponse
func (client *Client) DescribeStreamLocationBlockWithContext(ctx context.Context, request *DescribeStreamLocationBlockRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamLocationBlockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BlockType) {
		query["BlockType"] = request.BlockType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStreamLocationBlock"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStreamLocationBlockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the layout settings of a virtual studio.
//
// Description:
//
// Before calling this operation, add layout settings for a virtual studio by calling the [AddStudioLayout](https://help.aliyun.com/document_detail/2848062.html) operation. Then call this operation to retrieve the virtual studio layout settings.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 15 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeStudioLayoutsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStudioLayoutsResponse
func (client *Client) DescribeStudioLayoutsWithContext(ctx context.Context, request *DescribeStudioLayoutsRequest, runtime *dara.RuntimeOptions) (_result *DescribeStudioLayoutsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStudioLayouts"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStudioLayoutsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation queries stream pulling information for a specified Toutiao live stream.
//
// Description:
//
// You can call this API to query stream pulling information for a specified Toutiao live stream.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 100. If you exceed this limit, API calls are throttled. This may affect your business. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DescribeToutiaoLivePlayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeToutiaoLivePlayResponse
func (client *Client) DescribeToutiaoLivePlayWithContext(ctx context.Context, request *DescribeToutiaoLivePlayRequest, runtime *dara.RuntimeOptions) (_result *DescribeToutiaoLivePlayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeToutiaoLivePlay"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeToutiaoLivePlayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the stream ingest information for a specified Toutiao live stream.
//
// Description:
//
// You can call this API to query the stream ingest information for a specified Toutiao live stream.
//
// ## QPS limits
//
// The QPS limit for this API is 100 calls per second per user. Calls that exceed this limit are throttled, which may impact your business. We recommend that you call the API within this limit. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DescribeToutiaoLivePublishRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeToutiaoLivePublishResponse
func (client *Client) DescribeToutiaoLivePublishWithContext(ctx context.Context, request *DescribeToutiaoLivePublishRequest, runtime *dara.RuntimeOptions) (_result *DescribeToutiaoLivePublishResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeToutiaoLivePublish"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeToutiaoLivePublishResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增查询 uid 级别或域名app级别在线流
//
// @param request - DescribeUidOnlineStreamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUidOnlineStreamsResponse
func (client *Client) DescribeUidOnlineStreamsWithContext(ctx context.Context, request *DescribeUidOnlineStreamsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUidOnlineStreamsResponse, _err error) {
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
		Action:      dara.String("DescribeUidOnlineStreams"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUidOnlineStreamsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the daily peak upstream bandwidth of a user.
//
// Description:
//
// Calls the operation to query the daily peak upstream bandwidth of a user.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user for this operation is 5. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed. For more information, see [QPS limit](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DescribeUpBpsPeakDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUpBpsPeakDataResponse
func (client *Client) DescribeUpBpsPeakDataWithContext(ctx context.Context, request *DescribeUpBpsPeakDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUpBpsPeakDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainSwitch) {
		query["DomainSwitch"] = request.DomainSwitch
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUpBpsPeakData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUpBpsPeakDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the daily peak upstream bandwidth of a specific dedicated line for a user.
//
// Description:
//
// Calls the operation to query the daily peak upstream bandwidth of a specific dedicated line for a user.
//
// ## QPS limit
//
// The QPS limit for a single user on this operation is 5 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DescribeUpBpsPeakOfLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUpBpsPeakOfLineResponse
func (client *Client) DescribeUpBpsPeakOfLineWithContext(ctx context.Context, request *DescribeUpBpsPeakOfLineRequest, runtime *dara.RuntimeOptions) (_result *DescribeUpBpsPeakOfLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainSwitch) {
		query["DomainSwitch"] = request.DomainSwitch
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Line) {
		query["Line"] = request.Line
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUpBpsPeakOfLine"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUpBpsPeakOfLineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the daily peak number of concurrent upstream ingest streams for a user.
//
// Description:
//
// Calls the operation to query the daily peak number of concurrent upstream ingest streams for a user.
//
// ## QPS limit
//
// The QPS limit for a single user on this operation is 10 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - DescribeUpPeakPublishStreamDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUpPeakPublishStreamDataResponse
func (client *Client) DescribeUpPeakPublishStreamDataWithContext(ctx context.Context, request *DescribeUpPeakPublishStreamDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUpPeakPublishStreamDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainSwitch) {
		query["DomainSwitch"] = request.DomainSwitch
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUpPeakPublishStreamData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUpPeakPublishStreamDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call DisableLiveRealtimeLogDelivery to pause real-time log delivery for a domain name.
//
// Description:
//
// This operation pauses real-time log delivery for a specified streaming domain. Before you call this operation, you must specify a streaming domain that has real-time log delivery enabled. <props="china">Currently, this feature is available only for streaming domains. To push upstream real-time logs from an ingest domain, you must [submit a ticket](https://workorder.console.aliyun.com/console.htm#/ticket/add?productCode=live\\&commonQuestionId=4545\\&isSmart=true\\&iatraceid=1608439120675-2a5c48de0b84805313c708\\&channel=selfservice). <props="intl">Currently, this feature is available only for streaming domains. To push upstream real-time logs from an ingest domain, you must [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.12818093.nav-right.dticket.6cb216d07otFWR#/ticket/createIndex).
//
// ## QPS limit
//
// A single user can make up to 6,000 queries per second (QPS) to this API. If you exceed this limit, API calls are throttled, which can affect your business. We recommend that you call this API at a reasonable rate.
//
// @param request - DisableLiveRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableLiveRealtimeLogDeliveryResponse
func (client *Client) DisableLiveRealtimeLogDeliveryWithContext(ctx context.Context, request *DisableLiveRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *DisableLiveRealtimeLogDeliveryResponse, _err error) {
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
		Action:      dara.String("DisableLiveRealtimeLogDelivery"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dynamically updates a watermark.
//
// Description:
//
// Dynamically updating a watermark involves replacing the watermark template ID during a live stream. Before you call this operation, you must obtain the ID of the watermark template that you want to use. To obtain the IDs of available watermark templates, call the [DescribeLiveStreamWatermarks](https://help.aliyun.com/document_detail/2848102.html) operation.
//
// ## QPS limit
//
// You can call this operation up to 60 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - DynamicUpdateWaterMarkStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DynamicUpdateWaterMarkStreamRuleResponse
func (client *Client) DynamicUpdateWaterMarkStreamRuleWithContext(ctx context.Context, request *DynamicUpdateWaterMarkStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *DynamicUpdateWaterMarkStreamRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DynamicUpdateWaterMarkStreamRule"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DynamicUpdateWaterMarkStreamRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits a playlist.
//
// Description:
//
// This operation performs a full edit. You can use this operation to edit configuration information or replace multiple playlist items.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, your API calls are throttled. This throttling can affect your business. Call this API within the specified limit.
//
// @param request - EditPlaylistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditPlaylistResponse
func (client *Client) EditPlaylistWithContext(ctx context.Context, request *EditPlaylistRequest, runtime *dara.RuntimeOptions) (_result *EditPlaylistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProgramConfig) {
		query["ProgramConfig"] = request.ProgramConfig
	}

	if !dara.IsNil(request.ProgramId) {
		query["ProgramId"] = request.ProgramId
	}

	if !dara.IsNil(request.ProgramItems) {
		query["ProgramItems"] = request.ProgramItems
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditPlaylist"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditPlaylistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a video clip task.
//
// Description:
//
// This operation allows you to add a video clip task by specifying CasterId and ShowId.
//
// After a video clip task is added, no automatic notification is sent. You can call the [GetEditingJobInfo](https://help.aliyun.com/document_detail/2848059.html) operation to query the task status.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - EditShowAndReplaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditShowAndReplaceResponse
func (client *Client) EditShowAndReplaceWithContext(ctx context.Context, request *EditShowAndReplaceRequest, runtime *dara.RuntimeOptions) (_result *EditShowAndReplaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ShowId) {
		query["ShowId"] = request.ShowId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StorageInfo) {
		query["StorageInfo"] = request.StorageInfo
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditShowAndReplace"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditShowAndReplaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches a production studio to the standby resource.
//
// Description:
//
// You can call this operation to urgently switch a specified scene to the standby video. This operation applies only to Program (PGM) scenes.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 per user. If you exceed this limit, API calls are throttled, which may impact your business.
//
// @param request - EffectCasterUrgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EffectCasterUrgentResponse
func (client *Client) EffectCasterUrgentWithContext(ctx context.Context, request *EffectCasterUrgentRequest, runtime *dara.RuntimeOptions) (_result *EffectCasterUrgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EffectCasterUrgent"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EffectCasterUrgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the standby video of a production studio.
//
// Description:
//
// This operation adds a video resource to a specified scenario. The resource becomes active when it is referenced by the scenario.
//
// ## QPS limits
//
// The queries per second (QPS) limit for each user is 10. If you exceed this limit, API calls are throttled. Throttling may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - EffectCasterVideoResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EffectCasterVideoResourceResponse
func (client *Client) EffectCasterVideoResourceWithContext(ctx context.Context, request *EffectCasterVideoResourceRequest, runtime *dara.RuntimeOptions) (_result *EffectCasterVideoResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EffectCasterVideoResource"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EffectCasterVideoResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the EnableLiveRealtimeLogDelivery operation to enable real-time log delivery for a domain name.
//
// Description:
//
// <props="china">Currently, you can configure only streaming domains. To deliver real-time logs for ingest domains, [submit a ticket](https://workorder.console.aliyun.com/console.htm#/ticket/add?productCode=live\\&commonQuestionId=4545\\&isSmart=true\\&iatraceid=1608439120675-2a5c48de0b84805313c708\\&channel=selfservice). <props="intl">Currently, you can configure only streaming domains. To deliver real-time logs for ingest domains, [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.12818093.nav-right.dticket.6cb216d07otFWR#/ticket/createIndex).
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 6,000 calls per minute. If you exceed the limit, API calls are throttled, which can affect your business. Call this operation at a reasonable rate.
//
// @param request - EnableLiveRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableLiveRealtimeLogDeliveryResponse
func (client *Client) EnableLiveRealtimeLogDeliveryWithContext(ctx context.Context, request *EnableLiveRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *EnableLiveRealtimeLogDeliveryResponse, _err error) {
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
		Action:      dara.String("EnableLiveRealtimeLogDelivery"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables stream ingest for a specified stream. You can preset a time to resume the stream.
//
// Description:
//
// Calls this operation to disable stream ingest for a specified stream. You can preset a time to resume the stream. If no preset time is specified, call the [ResumeLiveStream](https://help.aliyun.com/document_detail/2847831.html) operation to resume the live stream. This operation currently supports only publisher (streamer ingest).
//
// > - This operation disables live streams by adding them to a blacklist. The upper limit is 10,000 streams. If this limit is exceeded, the disable operation fails. Monitor the number of currently disabled streams. Call the [DescribeLiveStreamsBlockList](https://help.aliyun.com/document_detail/2847825.html) operation to query the number of disabled streams.
//
// > - If you only interrupt a live stream without adding it to the blacklist, the stream does not count toward the 10,000-stream blacklist quota.
//
// ## QPS limit
//
// The maximum queries per second (QPS) per user for this operation is 20. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - ForbidLiveStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ForbidLiveStreamResponse
func (client *Client) ForbidLiveStreamWithContext(ctx context.Context, request *ForbidLiveStreamRequest, runtime *dara.RuntimeOptions) (_result *ForbidLiveStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.LiveStreamType) {
		query["LiveStreamType"] = request.LiveStreamType
	}

	if !dara.IsNil(request.Oneshot) {
		query["Oneshot"] = request.Oneshot
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResumeTime) {
		query["ResumeTime"] = request.ResumeTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ForbidLiveStream"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ForbidLiveStreamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all custom stream mix templates.
//
// Description:
//
// You can call this operation to retrieve all custom stream mix templates. The operation returns a list of template names and template configurations.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If this limit is exceeded, throttling is triggered, which may affect your business. We recommend that you are aware of this limit when you call this operation.
//
// @param request - GetAllCustomTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAllCustomTemplatesResponse
func (client *Client) GetAllCustomTemplatesWithContext(ctx context.Context, request *GetAllCustomTemplatesRequest, runtime *dara.RuntimeOptions) (_result *GetAllCustomTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAllCustomTemplates"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAllCustomTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a custom stream mixing template.
//
// Description:
//
// First, obtain the name of the custom template that you want to query. Then, call this operation to retrieve the custom stream mixing template.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 per user. If you exceed this limit, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - GetCustomTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomTemplateResponse
func (client *Client) GetCustomTemplateWithContext(ctx context.Context, request *GetCustomTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetCustomTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomTemplate"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an edge transcoding job.
//
// Description:
//
// To call this API operation, you must have permissions for the edge transcoding service.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user on this API is 6,000 calls per minute. API calls that exceed this limit are throttled. This may affect your business. We recommend that you call the API at a reasonable rate.
//
// @param request - GetEdgeTranscodeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEdgeTranscodeJobResponse
func (client *Client) GetEdgeTranscodeJobWithContext(ctx context.Context, request *GetEdgeTranscodeJobRequest, runtime *dara.RuntimeOptions) (_result *GetEdgeTranscodeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEdgeTranscodeJob"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEdgeTranscodeJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an edge transcoding template.
//
// Description:
//
// - This operation queries the details of an edge transcoding template.
//
// - You must have permissions to access the edge transcoding service before you can call this operation.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 6,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - GetEdgeTranscodeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEdgeTranscodeTemplateResponse
func (client *Client) GetEdgeTranscodeTemplateWithContext(ctx context.Context, request *GetEdgeTranscodeTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetEdgeTranscodeTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEdgeTranscodeTemplate"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEdgeTranscodeTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a video clip task.
//
// Description:
//
// - This operation allows you to query information about a video clip task for a specified CasterId and ShowId. Make sure that the parameters are set correctly when you call this operation.
//
// - If you specify the ShowId request parameter, the response returns the video clip task information for the specified show in the playlist.
//
// - If you do not specify the ShowId request parameter, the response returns the video clip task information for the entire playlist.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - GetEditingJobInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEditingJobInfoResponse
func (client *Client) GetEditingJobInfoWithContext(ctx context.Context, request *GetEditingJobInfoRequest, runtime *dara.RuntimeOptions) (_result *GetEditingJobInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ShowId) {
		query["ShowId"] = request.ShowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEditingJobInfo"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEditingJobInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the GetMessageApp operation to retrieve the details of a specified interactive message application.
//
// Description:
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 100. If you exceed this limit, API calls are throttled, which may affect your business. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - GetMessageAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageAppResponse
func (client *Client) GetMessageAppWithContext(ctx context.Context, request *GetMessageAppRequest, runtime *dara.RuntimeOptions) (_result *GetMessageAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMessageApp"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMessageAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call GetMessageGroup to retrieve message group details.
//
// Description:
//
// ## Usage notes
//
// The QPS limit for this API is 100 queries per second per user. If the limit is exceeded, API calls will be throttled, which may affect your business. You can make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - GetMessageGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageGroupResponse
func (client *Client) GetMessageGroupWithContext(ctx context.Context, request *GetMessageGroupRequest, runtime *dara.RuntimeOptions) (_result *GetMessageGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMessageGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMessageGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call GetMessageToken to obtain a persistent connection token. The client can use this token to communicate with various atomic capabilities through the LWP protocol via persistent connection.
//
// Description:
//
// ## Usage notes
//
// First obtain the client UserId, DeviceId, and DeviceType information, then pass them to the server. The server uses this interface to obtain the authentication token and returns it to the client. Different users need to use different UserIds, and different terminal devices need to use different DeviceIds.
//
// ## QPS limits
//
// The QPS limit for this API is 100 queries per second (QPS) per user. If you exceed this limit, API calls will be throttled, which may affect your business. You can call this API at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - GetMessageTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageTokenResponse
func (client *Client) GetMessageTokenWithContext(ctx context.Context, request *GetMessageTokenRequest, runtime *dara.RuntimeOptions) (_result *GetMessageTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DeviceId) {
		body["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.DeviceType) {
		body["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMessageToken"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMessageTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取转码任务状态
//
// @param request - GetTranscodeTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranscodeTaskStatusResponse
func (client *Client) GetTranscodeTaskStatusWithContext(ctx context.Context, request *GetTranscodeTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *GetTranscodeTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.PushDomain) {
		query["PushDomain"] = request.PushDomain
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.TranscodingTemplate) {
		query["TranscodingTemplate"] = request.TranscodingTemplate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTranscodeTaskStatus"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTranscodeTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Prefetches Real-Time Communication (RTC) streams.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. API calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - HotLiveRtcStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotLiveRtcStreamResponse
func (client *Client) HotLiveRtcStreamWithContext(ctx context.Context, request *HotLiveRtcStreamRequest, runtime *dara.RuntimeOptions) (_result *HotLiveRtcStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AudioMsid) {
		query["AudioMsid"] = request.AudioMsid
	}

	if !dara.IsNil(request.ConnectionTimeout) {
		query["ConnectionTimeout"] = request.ConnectionTimeout
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.MediaTimeout) {
		query["MediaTimeout"] = request.MediaTimeout
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.VideoMsid) {
		query["VideoMsid"] = request.VideoMsid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotLiveRtcStream"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HotLiveRtcStreamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scheduled task.
//
// Description:
//
// This operation allows you to configure a scheduled task to start and stop a playlist at specified times. Make sure that the parameters are set correctly when you call this operation.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - InitializeAutoShowListTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeAutoShowListTaskResponse
func (client *Client) InitializeAutoShowListTaskWithContext(ctx context.Context, request *InitializeAutoShowListTaskRequest, runtime *dara.RuntimeOptions) (_result *InitializeAutoShowListTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallBackUrl) {
		query["CallBackUrl"] = request.CallBackUrl
	}

	if !dara.IsNil(request.CasterConfig) {
		query["CasterConfig"] = request.CasterConfig
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitializeAutoShowListTask"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitializeAutoShowListTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Join a message group.
//
// Description:
//
// The QPS limit for this API is 200 calls per second per user. If this limit is exceeded, API calls will be throttled, which may affect your business. Please use this API responsibly.
//
// @param request - JoinMessageGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JoinMessageGroupResponse
func (client *Client) JoinMessageGroupWithContext(ctx context.Context, request *JoinMessageGroupRequest, runtime *dara.RuntimeOptions) (_result *JoinMessageGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BroadCastStatistics) {
		body["BroadCastStatistics"] = request.BroadCastStatistics
	}

	if !dara.IsNil(request.BroadCastType) {
		body["BroadCastType"] = request.BroadCastType
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("JoinMessageGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &JoinMessageGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a user from an interactive messaging group.
//
// Description:
//
// Before you call this operation, make sure that you have called the [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html) operation to create an interactive messaging group.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - KickLiveMessageGroupUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KickLiveMessageGroupUserResponse
func (client *Client) KickLiveMessageGroupUserWithContext(ctx context.Context, request *KickLiveMessageGroupUserRequest, runtime *dara.RuntimeOptions) (_result *KickLiveMessageGroupUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.KickoffUser) {
		query["KickoffUser"] = request.KickoffUser
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KickLiveMessageGroupUser"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KickLiveMessageGroupUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call LeaveMessageGroup to allow users to exit a message group.
//
// Description:
//
// ## Usage notes
//
// The QPS limit for this API is 100 calls per second per user. If this limit is exceeded, API calls will be throttled, which may affect your business. You can call this API properly to avoid issues. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - LeaveMessageGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LeaveMessageGroupResponse
func (client *Client) LeaveMessageGroupWithContext(ctx context.Context, request *LeaveMessageGroupRequest, runtime *dara.RuntimeOptions) (_result *LeaveMessageGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BroadCastStatistics) {
		body["BroadCastStatistics"] = request.BroadCastStatistics
	}

	if !dara.IsNil(request.BroadCastType) {
		body["BroadCastType"] = request.BroadCastType
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LeaveMessageGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LeaveMessageGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of edge transcoding tasks.
//
// Description:
//
// - This operation queries the list of edge transcoding tasks.
//
// - You must have the permissions to access the edge transcoding service before you can call this operation.
//
// - This operation returns only tasks from the last 180 days, and the last operation time must be within this 180-day period.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 6,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - ListEdgeTranscodeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEdgeTranscodeJobResponse
func (client *Client) ListEdgeTranscodeJobWithContext(ctx context.Context, request *ListEdgeTranscodeJobRequest, runtime *dara.RuntimeOptions) (_result *ListEdgeTranscodeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEdgeTranscodeJob"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEdgeTranscodeJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of edge transcoding templates.
//
// Description:
//
// - This operation queries the list of edge transcoding templates.
//
// - You must have permissions to access the edge transcoding service before you can call this operation.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 6,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - ListEdgeTranscodeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEdgeTranscodeTemplateResponse
func (client *Client) ListEdgeTranscodeTemplateWithContext(ctx context.Context, request *ListEdgeTranscodeTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListEdgeTranscodeTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.VideoCodec) {
		query["VideoCodec"] = request.VideoCodec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEdgeTranscodeTemplate"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEdgeTranscodeTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the events generated in channels to which you subscribe.
//
// Description:
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - ListEventSubRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventSubResponse
func (client *Client) ListEventSubWithContext(ctx context.Context, request *ListEventSubRequest, runtime *dara.RuntimeOptions) (_result *ListEventSubResponse, _err error) {
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
		Action:      dara.String("ListEventSub"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEventSubResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries callback records.
//
// Description:
//
// - Maximum query span: 7 days.
//
// - Minimum query granularity: 1 minute.
//
// - Maximum query range: data from the last 7 days.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - ListEventSubEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventSubEventResponse
func (client *Client) ListEventSubEventWithContext(ctx context.Context, request *ListEventSubEventRequest, runtime *dara.RuntimeOptions) (_result *ListEventSubEventResponse, _err error) {
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
		Action:      dara.String("ListEventSubEvent"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEventSubEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves stream delay configurations.
//
// Description:
//
// The queries per second (QPS) limit for a single user is 60. If you exceed this limit, API calls are throttled, which may impact your business. Ensure that you call this operation within this limit.
//
// @param request - ListLiveDelayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveDelayConfigResponse
func (client *Client) ListLiveDelayConfigWithContext(ctx context.Context, request *ListLiveDelayConfigRequest, runtime *dara.RuntimeOptions) (_result *ListLiveDelayConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveDelayConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveDelayConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of interactive messaging applications by calling ListLiveMessageApps.
//
// Description:
//
// The China (Shanghai) region is replaced by the Singapore region in the example. The China (Shanghai) region is replaced by the Singapore region in the example. The per-user QPS limit for this operation is 50. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - ListLiveMessageAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveMessageAppsResponse
func (client *Client) ListLiveMessageAppsWithContext(ctx context.Context, request *ListLiveMessageAppsRequest, runtime *dara.RuntimeOptions) (_result *ListLiveMessageAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveMessageApps"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveMessageAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of interactive message groups with pagination support.
//
// Description:
//
// Before calling this operation, you must have already called [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html) to create an interactive message group.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - ListLiveMessageGroupByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveMessageGroupByPageResponse
func (client *Client) ListLiveMessageGroupByPageWithContext(ctx context.Context, request *ListLiveMessageGroupByPageRequest, runtime *dara.RuntimeOptions) (_result *ListLiveMessageGroupByPageResponse, _err error) {
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
		Action:      dara.String("ListLiveMessageGroupByPage"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveMessageGroupByPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the message list of a group by calling ListLiveMessageGroupMessages.
//
// Description:
//
// Before calling this operation, you must have already created an interactive message group by calling [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - ListLiveMessageGroupMessagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveMessageGroupMessagesResponse
func (client *Client) ListLiveMessageGroupMessagesWithContext(ctx context.Context, request *ListLiveMessageGroupMessagesRequest, runtime *dara.RuntimeOptions) (_result *ListLiveMessageGroupMessagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MsgType) {
		query["MsgType"] = request.MsgType
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveMessageGroupMessages"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveMessageGroupMessagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of users in a group by calling ListLiveMessageGroupUsers.
//
// Description:
//
// - Before calling this operation, you must have already created an interactive messaging group by calling [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html).
//
// - For super-large groups (groups with more than 2,000 members), member list queries are no longer supported. Additionally, notifications for members joining or leaving the group are sent at intervals of at least 5 seconds. The notifications do not display the full list of users who joined or left, but they display the exact group member count. Once a group is upgraded to a super-large group, the member list is immediately cleared. The group cannot be reverted to a regular group until all members leave the group (the group is closed). When the group is reopened, it is restored to a regular group.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - ListLiveMessageGroupUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveMessageGroupUsersResponse
func (client *Client) ListLiveMessageGroupUsersWithContext(ctx context.Context, request *ListLiveMessageGroupUsersRequest, runtime *dara.RuntimeOptions) (_result *ListLiveMessageGroupUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveMessageGroupUsers"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveMessageGroupUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the interactive messaging groups in an interactive messaging application.
//
// Description:
//
// Before you call this operation, make sure that you have called the [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html) operation to create an interactive messaging group.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - ListLiveMessageGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveMessageGroupsResponse
func (client *Client) ListLiveMessageGroupsWithContext(ctx context.Context, request *ListLiveMessageGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListLiveMessageGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.GroupStatus) {
		query["GroupStatus"] = request.GroupStatus
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveMessageGroups"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveMessageGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ListLiveRealtimeLogDelivery operation to query all real-time log delivery configurations.
//
// Description:
//
// You can call this operation to query all real-time log delivery configurations. Make sure that the parameters are set as required.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 6,000 calls per minute for each user. API calls that exceed this limit are throttled, which may impact your business. Call this operation at a reasonable rate.
//
// @param request - ListLiveRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveRealtimeLogDeliveryResponse
func (client *Client) ListLiveRealtimeLogDeliveryWithContext(ctx context.Context, request *ListLiveRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *ListLiveRealtimeLogDeliveryResponse, _err error) {
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
		Action:      dara.String("ListLiveRealtimeLogDelivery"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call ListLiveRealtimeLogDeliveryDomains to query all streaming domains that are configured for the real-time log delivery service.
//
// Description:
//
// - This operation queries all streaming domains that are configured for the real-time log delivery service. The response indicates whether the service is online or offline for each domain.
//
// - You can call [DescribeLiveDomainRealtimeLogDelivery](https://help.aliyun.com/document_detail/2848121.html) to query parameters such as Project, Logstore, and Region.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 6,000 calls per minute. If you exceed this limit, API calls are throttled. Throttling may affect your business. We recommend that you call this operation within the specified limit.
//
// @param request - ListLiveRealtimeLogDeliveryDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveRealtimeLogDeliveryDomainsResponse
func (client *Client) ListLiveRealtimeLogDeliveryDomainsWithContext(ctx context.Context, request *ListLiveRealtimeLogDeliveryDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListLiveRealtimeLogDeliveryDomainsResponse, _err error) {
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
		Action:      dara.String("ListLiveRealtimeLogDeliveryDomains"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveRealtimeLogDeliveryDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all real-time log delivery services.
//
// Description:
//
// Make sure that you configure the parameters as required when you call this operation.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 6,000 calls per minute. If you exceed this limit, API calls are throttled. This can affect your business, so you should plan your calls accordingly.
//
// @param request - ListLiveRealtimeLogDeliveryInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveRealtimeLogDeliveryInfosResponse
func (client *Client) ListLiveRealtimeLogDeliveryInfosWithContext(ctx context.Context, request *ListLiveRealtimeLogDeliveryInfosRequest, runtime *dara.RuntimeOptions) (_result *ListLiveRealtimeLogDeliveryInfosResponse, _err error) {
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
		Action:      dara.String("ListLiveRealtimeLogDeliveryInfos"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveRealtimeLogDeliveryInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resources and tags that meet the specified conditions.
//
// @param request - ListLiveTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveTagResourcesResponse
func (client *Client) ListLiveTagResourcesWithContext(ctx context.Context, request *ListLiveTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListLiveTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TagOwnerBid) {
		query["TagOwnerBid"] = request.TagOwnerBid
	}

	if !dara.IsNil(request.TagOwnerUid) {
		query["TagOwnerUid"] = request.TagOwnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveTagResources"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveTagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the message list.
//
// Description:
//
// ## QPS limits
//
// The single-user QPS limit for this API operation is 100 queries per second (QPS). If the limit is exceeded, API calls will be throttled, which may affect your business. You can call the API operation properly to avoid this issue. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - ListMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessageResponse
func (client *Client) ListMessageWithContext(ctx context.Context, request *ListMessageRequest, runtime *dara.RuntimeOptions) (_result *ListMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortType) {
		body["SortType"] = request.SortType
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMessage"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call ListMessageApp to query the list of interactive message applications.
//
// Description:
//
// ## QPS limits
//
// The single-user QPS limit for this API is 100 queries per second (QPS). API calls exceeding this limit will be throttled, which may affect your business. You can call this API at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - ListMessageAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessageAppResponse
func (client *Client) ListMessageAppWithContext(ctx context.Context, request *ListMessageAppRequest, runtime *dara.RuntimeOptions) (_result *ListMessageAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortType) {
		body["SortType"] = request.SortType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMessageApp"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMessageAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call ListMessageGroup to query the message group list for a specified user.
//
// Description:
//
// ## QPS limits
//
// The single-user QPS limit for this API is 100 queries per second. If this limit is exceeded, API calls will be throttled, which may affect your business. You can call this API at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - ListMessageGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessageGroupResponse
func (client *Client) ListMessageGroupWithContext(ctx context.Context, request *ListMessageGroupRequest, runtime *dara.RuntimeOptions) (_result *ListMessageGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortType) {
		body["SortType"] = request.SortType
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMessageGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMessageGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query message group members.
//
// Description:
//
// ## QPS limits
//
// The single-user QPS limit for this API is 100 queries per second (QPS). If you exceed this limit, API calls will be throttled, which may affect your business. You can make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - ListMessageGroupUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessageGroupUserResponse
func (client *Client) ListMessageGroupUserWithContext(ctx context.Context, request *ListMessageGroupUserRequest, runtime *dara.RuntimeOptions) (_result *ListMessageGroupUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortType) {
		body["SortType"] = request.SortType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMessageGroupUser"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMessageGroupUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call ListMessageGroupUserById to query user information by user ID list.
//
// Description:
//
// ## QPS limits
//
// The QPS limit for this API is 100 queries per second per user. If the limit is exceeded, API calls will be throttled, which may affect your business. You can call the API properly to avoid this issue. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param tmpReq - ListMessageGroupUserByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessageGroupUserByIdResponse
func (client *Client) ListMessageGroupUserByIdWithContext(ctx context.Context, tmpReq *ListMessageGroupUserByIdRequest, runtime *dara.RuntimeOptions) (_result *ListMessageGroupUserByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMessageGroupUserByIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserIdList) {
		request.UserIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIdList, dara.String("UserIdList"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.UserIdListShrink) {
		body["UserIdList"] = request.UserIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMessageGroupUserById"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMessageGroupUserByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call ListMuteGroupUser to query the list of muted members in a message group.
//
// Description:
//
// ## QPS limits
//
// The QPS limit for this API is 100 queries per second (QPS) per user. If the limit is exceeded, API calls will be throttled, which may affect your business. You can call this API at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - ListMuteGroupUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMuteGroupUserResponse
func (client *Client) ListMuteGroupUserWithContext(ctx context.Context, request *ListMuteGroupUserRequest, runtime *dara.RuntimeOptions) (_result *ListMuteGroupUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OperatorUserId) {
		body["OperatorUserId"] = request.OperatorUserId
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMuteGroupUser"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMuteGroupUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries episode lists.
//
// Description:
//
// Call the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation to add items to a playlist before you query it.
//
// ## QPS limits
//
// This operation is limited to 10 queries per second (QPS) per user. API calls that exceed this limit are throttled. Throttling may affect your business. Plan your calls accordingly.
//
// @param request - ListPlaylistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPlaylistResponse
func (client *Client) ListPlaylistWithContext(ctx context.Context, request *ListPlaylistRequest, runtime *dara.RuntimeOptions) (_result *ListPlaylistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProgramId) {
		query["ProgramId"] = request.ProgramId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPlaylist"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPlaylistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the items in a specified playlist.
//
// Description:
//
// Before you call this operation, you must call the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation to add items to a playlist.
//
// ## QPS limit
//
// This operation has a queries per second (QPS) limit of 10 calls per user. If you exceed this limit, your API calls are throttled. Throttling can affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - ListPlaylistItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPlaylistItemsResponse
func (client *Client) ListPlaylistItemsWithContext(ctx context.Context, request *ListPlaylistItemsRequest, runtime *dara.RuntimeOptions) (_result *ListPlaylistItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProgramId) {
		query["ProgramId"] = request.ProgramId
	}

	if !dara.IsNil(request.ProgramItemIds) {
		query["ProgramItemIds"] = request.ProgramItemIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPlaylistItems"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPlaylistItemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of online channels.
//
// @param request - ListRTCLiveRoomsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRTCLiveRoomsResponse
func (client *Client) ListRTCLiveRoomsWithContext(ctx context.Context, request *ListRTCLiveRoomsRequest, runtime *dara.RuntimeOptions) (_result *ListRTCLiveRoomsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRTCLiveRooms"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRTCLiveRoomsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the event subscription records for stream mixing and forwarding.
//
// Description:
//
// Queries the event subscription records for stream mixing and forwarding. You can query data from the last seven days.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - ListRtcMPUEventSubRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRtcMPUEventSubRecordResponse
func (client *Client) ListRtcMPUEventSubRecordWithContext(ctx context.Context, request *ListRtcMPUEventSubRecordRequest, runtime *dara.RuntimeOptions) (_result *ListRtcMPUEventSubRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.SubId) {
		query["SubId"] = request.SubId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRtcMPUEventSubRecord"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRtcMPUEventSubRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the parameter details of a stream mixing and relaying task.
//
// Description:
//
// - This operation queries the parameter details of a stream mixing and relaying task. Only stream mixing and relaying tasks created by using API operations can be queried.
//
// - For a paged query, the query results are sorted by task update time in descending order by default.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If the limit is exceeded, the API invocation is throttled, which may affect your business. Invoke this operation as needed.
//
// @param request - ListRtcMPUTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRtcMPUTaskDetailResponse
func (client *Client) ListRtcMPUTaskDetailWithContext(ctx context.Context, request *ListRtcMPUTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *ListRtcMPUTaskDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRtcMPUTaskDetail"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRtcMPUTaskDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 直播回源监控数据
//
// @param tmpReq - LiveUpstreamQosDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LiveUpstreamQosDataResponse
func (client *Client) LiveUpstreamQosDataWithContext(ctx context.Context, tmpReq *LiveUpstreamQosDataRequest, runtime *dara.RuntimeOptions) (_result *LiveUpstreamQosDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &LiveUpstreamQosDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CdnDomains) {
		request.CdnDomainsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CdnDomains, dara.String("CdnDomains"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CdnIsps) {
		request.CdnIspsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CdnIsps, dara.String("CdnIsps"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CdnProvinces) {
		request.CdnProvincesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CdnProvinces, dara.String("CdnProvinces"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.KwaiSidcs) {
		request.KwaiSidcsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KwaiSidcs, dara.String("KwaiSidcs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.KwaiTsc) {
		request.KwaiTscShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KwaiTsc, dara.String("KwaiTsc"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UpstreamDomains) {
		request.UpstreamDomainsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpstreamDomains, dara.String("UpstreamDomains"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CdnDomainsShrink) {
		query["CdnDomains"] = request.CdnDomainsShrink
	}

	if !dara.IsNil(request.CdnIspsShrink) {
		query["CdnIsps"] = request.CdnIspsShrink
	}

	if !dara.IsNil(request.CdnProvincesShrink) {
		query["CdnProvinces"] = request.CdnProvincesShrink
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.KwaiSidcsShrink) {
		query["KwaiSidcs"] = request.KwaiSidcsShrink
	}

	if !dara.IsNil(request.KwaiTscShrink) {
		query["KwaiTsc"] = request.KwaiTscShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UpstreamDomainsShrink) {
		query["UpstreamDomains"] = request.UpstreamDomainsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LiveUpstreamQosData"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LiveUpstreamQosDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 咪咕定制直播拉转推启动接口
//
// @param request - MiguLivePullToPushStartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MiguLivePullToPushStartResponse
func (client *Client) MiguLivePullToPushStartWithContext(ctx context.Context, request *MiguLivePullToPushStartRequest, runtime *dara.RuntimeOptions) (_result *MiguLivePullToPushStartResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.MiguData) {
		query["MiguData"] = request.MiguData
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MiguLivePullToPushStart"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MiguLivePullToPushStartResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 咪咕定制直播拉转推启动接口
//
// @param request - MiguLivePullToPushStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MiguLivePullToPushStatusResponse
func (client *Client) MiguLivePullToPushStatusWithContext(ctx context.Context, request *MiguLivePullToPushStatusRequest, runtime *dara.RuntimeOptions) (_result *MiguLivePullToPushStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.MiguData) {
		query["MiguData"] = request.MiguData
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MiguLivePullToPushStatus"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MiguLivePullToPushStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a component of a production studio.
//
// Description:
//
// You can call this operation to modify a production studio component, such as a text, image, or translation caption component.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - ModifyCasterComponentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCasterComponentResponse
func (client *Client) ModifyCasterComponentWithContext(ctx context.Context, request *ModifyCasterComponentRequest, runtime *dara.RuntimeOptions) (_result *ModifyCasterComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaptionLayerContent) {
		query["CaptionLayerContent"] = request.CaptionLayerContent
	}

	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.ComponentId) {
		query["ComponentId"] = request.ComponentId
	}

	if !dara.IsNil(request.ComponentLayer) {
		query["ComponentLayer"] = request.ComponentLayer
	}

	if !dara.IsNil(request.ComponentName) {
		query["ComponentName"] = request.ComponentName
	}

	if !dara.IsNil(request.ComponentType) {
		query["ComponentType"] = request.ComponentType
	}

	if !dara.IsNil(request.Effect) {
		query["Effect"] = request.Effect
	}

	if !dara.IsNil(request.ImageLayerContent) {
		query["ImageLayerContent"] = request.ImageLayerContent
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TextLayerContent) {
		query["TextLayerContent"] = request.TextLayerContent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCasterComponent"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCasterComponentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a production studio episode. The episode type cannot be modified.
//
// Description:
//
// You can call this operation to modify the configuration of a production studio episode. The episode type cannot be modified.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 4 calls per second. If this limit is exceeded, API calls are throttled. This may affect your business. We recommend that you call this API operation at a reasonable rate.
//
// @param request - ModifyCasterEpisodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCasterEpisodeResponse
func (client *Client) ModifyCasterEpisodeWithContext(ctx context.Context, request *ModifyCasterEpisodeRequest, runtime *dara.RuntimeOptions) (_result *ModifyCasterEpisodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.ComponentId) {
		query["ComponentId"] = request.ComponentId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EpisodeId) {
		query["EpisodeId"] = request.EpisodeId
	}

	if !dara.IsNil(request.EpisodeName) {
		query["EpisodeName"] = request.EpisodeName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.SwitchType) {
		query["SwitchType"] = request.SwitchType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCasterEpisode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCasterEpisodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a layout configuration. Only the items to be modified need to be passed. Items that do not require modification do not need to be included.
//
// Description:
//
// Create a production studio by calling the [CreateCaster operation](https://help.aliyun.com/document_detail/2848009.html) first, and then call this operation to modify the layout configuration. Only the items to be modified need to be passed. Items that do not require modification do not need to be included. This operation currently supports the following element fill modes: default and adaptive.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - ModifyCasterLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCasterLayoutResponse
func (client *Client) ModifyCasterLayoutWithContext(ctx context.Context, request *ModifyCasterLayoutRequest, runtime *dara.RuntimeOptions) (_result *ModifyCasterLayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AudioLayer) {
		query["AudioLayer"] = request.AudioLayer
	}

	if !dara.IsNil(request.BlendList) {
		query["BlendList"] = request.BlendList
	}

	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.MixList) {
		query["MixList"] = request.MixList
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VideoLayer) {
		query["VideoLayer"] = request.VideoLayer
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCasterLayout"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCasterLayoutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation modifies the program list of a production studio.
//
// Description:
//
// You can call the [AddCasterProgram](https://help.aliyun.com/document_detail/2848074.html) operation to add a program list for a production studio. Then, you can call this operation to modify the program list. This operation supports programs of the video source and component types.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 4 calls per second. If you exceed this limit, API calls are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - ModifyCasterProgramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCasterProgramResponse
func (client *Client) ModifyCasterProgramWithContext(ctx context.Context, request *ModifyCasterProgramRequest, runtime *dara.RuntimeOptions) (_result *ModifyCasterProgramResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.Episode) {
		query["Episode"] = request.Episode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCasterProgram"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCasterProgramResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the video resources of a production studio.
//
// Description:
//
// Before calling this operation, you must create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation.
//
// ## QPS limit
//
// This operation supports up to 10 queries per second (QPS) per user. If you exceed this limit, your API calls are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - ModifyCasterVideoResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCasterVideoResourceResponse
func (client *Client) ModifyCasterVideoResourceWithContext(ctx context.Context, request *ModifyCasterVideoResourceRequest, runtime *dara.RuntimeOptions) (_result *ModifyCasterVideoResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginOffset) {
		query["BeginOffset"] = request.BeginOffset
	}

	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.EndOffset) {
		query["EndOffset"] = request.EndOffset
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.LiveStreamUrl) {
		query["LiveStreamUrl"] = request.LiveStreamUrl
	}

	if !dara.IsNil(request.MaterialId) {
		query["MaterialId"] = request.MaterialId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PtsCallbackInterval) {
		query["PtsCallbackInterval"] = request.PtsCallbackInterval
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepeatNum) {
		query["RepeatNum"] = request.RepeatNum
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.VodUrl) {
		query["VodUrl"] = request.VodUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCasterVideoResource"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCasterVideoResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API is used to modify the specified virtual studio template.
//
// @param tmpReq - ModifyLiveAIStudioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLiveAIStudioResponse
func (client *Client) ModifyLiveAIStudioWithContext(ctx context.Context, tmpReq *ModifyLiveAIStudioRequest, runtime *dara.RuntimeOptions) (_result *ModifyLiveAIStudioResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyLiveAIStudioShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MattingLayout) {
		request.MattingLayoutShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MattingLayout, dara.String("MattingLayout"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MediaLayout) {
		request.MediaLayoutShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MediaLayout, dara.String("MediaLayout"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BackgroundResourceId) {
		query["BackgroundResourceId"] = request.BackgroundResourceId
	}

	if !dara.IsNil(request.BackgroundResourceUrl) {
		query["BackgroundResourceUrl"] = request.BackgroundResourceUrl
	}

	if !dara.IsNil(request.BackgroundType) {
		query["BackgroundType"] = request.BackgroundType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Height) {
		query["Height"] = request.Height
	}

	if !dara.IsNil(request.MattingLayoutShrink) {
		query["MattingLayout"] = request.MattingLayoutShrink
	}

	if !dara.IsNil(request.MattingType) {
		query["MattingType"] = request.MattingType
	}

	if !dara.IsNil(request.MediaLayoutShrink) {
		query["MediaLayout"] = request.MediaLayoutShrink
	}

	if !dara.IsNil(request.MediaResourceId) {
		query["MediaResourceId"] = request.MediaResourceId
	}

	if !dara.IsNil(request.MediaResourceUrl) {
		query["MediaResourceUrl"] = request.MediaResourceUrl
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StudioName) {
		query["StudioName"] = request.StudioName
	}

	if !dara.IsNil(request.Width) {
		query["Width"] = request.Width
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLiveAIStudio"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLiveAIStudioResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the acceleration region of a domain name.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 100 queries per second (QPS) per user. API calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - ModifyLiveDomainSchdmByPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLiveDomainSchdmByPropertyResponse
func (client *Client) ModifyLiveDomainSchdmByPropertyWithContext(ctx context.Context, request *ModifyLiveDomainSchdmByPropertyRequest, runtime *dara.RuntimeOptions) (_result *ModifyLiveDomainSchdmByPropertyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Property) {
		query["Property"] = request.Property
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLiveDomainSchdmByProperty"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLiveDomainSchdmByPropertyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the security audit settings of an interactive messaging application.
//
// Description:
//
// The China site Chinese edition of this API has a single-user QPS limit of 50 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this API at a reasonable frequency.
//
// @param request - ModifyLiveMessageAppAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLiveMessageAppAuditResponse
func (client *Client) ModifyLiveMessageAppAuditWithContext(ctx context.Context, request *ModifyLiveMessageAppAuditRequest, runtime *dara.RuntimeOptions) (_result *ModifyLiveMessageAppAuditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AuditType) {
		query["AuditType"] = request.AuditType
	}

	if !dara.IsNil(request.AuditUrl) {
		query["AuditUrl"] = request.AuditUrl
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLiveMessageAppAudit"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLiveMessageAppAuditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the callback settings of an interactive messaging application.
//
// Description:
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - ModifyLiveMessageAppCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLiveMessageAppCallbackResponse
func (client *Client) ModifyLiveMessageAppCallbackWithContext(ctx context.Context, request *ModifyLiveMessageAppCallbackRequest, runtime *dara.RuntimeOptions) (_result *ModifyLiveMessageAppCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.EventCallbackUrl) {
		query["EventCallbackUrl"] = request.EventCallbackUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLiveMessageAppCallback"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLiveMessageAppCallbackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables or enables an interactive messaging application.
//
// Description:
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - ModifyLiveMessageAppDisableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLiveMessageAppDisableResponse
func (client *Client) ModifyLiveMessageAppDisableWithContext(ctx context.Context, request *ModifyLiveMessageAppDisableRequest, runtime *dara.RuntimeOptions) (_result *ModifyLiveMessageAppDisableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.Disable) {
		query["Disable"] = request.Disable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLiveMessageAppDisable"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLiveMessageAppDisableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ModifyLiveMessageGroup operation to modify information about an interactive message group.
//
// Description:
//
// Before you call this operation, call [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html) to create an interactive message group.
//
// ## QPS limit
//
// This operation has a queries per second (QPS) limit of 50 for each user. If you exceed the limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param tmpReq - ModifyLiveMessageGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLiveMessageGroupResponse
func (client *Client) ModifyLiveMessageGroupWithContext(ctx context.Context, tmpReq *ModifyLiveMessageGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyLiveMessageGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyLiveMessageGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AdminList) {
		request.AdminListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdminList, dara.String("AdminList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdminListShrink) {
		query["AdminList"] = request.AdminListShrink
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupInfo) {
		query["GroupInfo"] = request.GroupInfo
	}

	if !dara.IsNil(request.ModifyAdmin) {
		query["ModifyAdmin"] = request.ModifyAdmin
	}

	if !dara.IsNil(request.ModifyInfo) {
		query["ModifyInfo"] = request.ModifyInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLiveMessageGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLiveMessageGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the mute status of users in a group.
//
// Description:
//
// Before you call this operation, you must call the [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html) operation to create an interactive messaging group.
//
// ## QPS limit
//
// A single user can make up to 10 queries per second (QPS). If you exceed this limit, API calls are throttled, which may impact your business. Call this operation at a reasonable rate.
//
// @param tmpReq - ModifyLiveMessageGroupBandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLiveMessageGroupBandResponse
func (client *Client) ModifyLiveMessageGroupBandWithContext(ctx context.Context, tmpReq *ModifyLiveMessageGroupBandRequest, runtime *dara.RuntimeOptions) (_result *ModifyLiveMessageGroupBandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyLiveMessageGroupBandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BannnedUsers) {
		request.BannnedUsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BannnedUsers, dara.String("BannnedUsers"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.ExceptUsers) {
		request.ExceptUsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExceptUsers, dara.String("ExceptUsers"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BannedAll) {
		query["BannedAll"] = request.BannedAll
	}

	if !dara.IsNil(request.BannnedUsersShrink) {
		query["BannnedUsers"] = request.BannnedUsersShrink
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.ExceptUsersShrink) {
		query["ExceptUsers"] = request.ExceptUsersShrink
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLiveMessageGroupBand"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLiveMessageGroupBandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a user.
//
// Description:
//
// 本接口的单用户QPS限制为50次/秒。超过限制，API调用会被限流，这可能会影响您的业务，请合理调用。
//
// @param request - ModifyLiveMessageUserInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLiveMessageUserInfoResponse
func (client *Client) ModifyLiveMessageUserInfoWithContext(ctx context.Context, request *ModifyLiveMessageUserInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyLiveMessageUserInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserMetaInfo) {
		query["UserMetaInfo"] = request.UserMetaInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLiveMessageUserInfo"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLiveMessageUserInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the real-time log delivery configuration for a domain name.
//
// Description:
//
// - Call this operation to modify the real-time log delivery configuration for a domain name. A domain name can deliver logs to only one Logstore at a time. <props="china">Currently, only streaming domains can be configured. To push real-time upstream logs (by configuring an ingest domain), [submit a ticket](https://workorder.console.aliyun.com/console.htm#/ticket/add?productCode=live\\&commonQuestionId=4545\\&isSmart=true\\&iatraceid=1608439120675-2a5c48de0b84805313c708\\&channel=selfservice). <props="intl">Currently, only streaming domains can be configured. To push real-time upstream logs (by configuring an ingest domain), [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.12818093.nav-right.dticket.6cb216d07otFWR#/ticket/createIndex).
//
// - Call [DescribeLiveDomainRealtimeLogDelivery](https://help.aliyun.com/document_detail/2848121.html) to query information about the Project, Logstore, and Region parameters.
//
// ## QPS limits
//
// You can call this operation up to 6,000 times per minute per user. If you exceed the queries per second (QPS) limit, API calls are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - ModifyLiveRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLiveRealtimeLogDeliveryResponse
func (client *Client) ModifyLiveRealtimeLogDeliveryWithContext(ctx context.Context, request *ModifyLiveRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *ModifyLiveRealtimeLogDeliveryResponse, _err error) {
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
		Action:      dara.String("ModifyLiveRealtimeLogDelivery"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an episode list.
//
// Description:
//
// This API operation modifies the position of a show in a playlist, the total number of playback loops for the playlist, and the specific playback time of the highest-priority show.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this API operation is 10 calls per second per user. If this limit is exceeded, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - ModifyShowListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyShowListResponse
func (client *Client) ModifyShowListWithContext(ctx context.Context, request *ModifyShowListRequest, runtime *dara.RuntimeOptions) (_result *ModifyShowListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.HighPriorityShowId) {
		query["HighPriorityShowId"] = request.HighPriorityShowId
	}

	if !dara.IsNil(request.HighPriorityShowStartTime) {
		query["HighPriorityShowStartTime"] = request.HighPriorityShowStartTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepeatTimes) {
		query["RepeatTimes"] = request.RepeatTimes
	}

	if !dara.IsNil(request.ShowId) {
		query["ShowId"] = request.ShowId
	}

	if !dara.IsNil(request.Spot) {
		query["Spot"] = request.Spot
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyShowList"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyShowListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the layout of a production studio.
//
// Description:
//
// You can call this operation to modify the layout of a production studio. When you modify the layout settings, pass only the parameters that you want to change.
//
// ## QPS limit
//
// The QPS limit for this operation is 10 calls per second for each user. If you exceed the limit, API calls are throttled. This may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - ModifyStudioLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyStudioLayoutResponse
func (client *Client) ModifyStudioLayoutWithContext(ctx context.Context, request *ModifyStudioLayoutRequest, runtime *dara.RuntimeOptions) (_result *ModifyStudioLayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BgImageConfig) {
		query["BgImageConfig"] = request.BgImageConfig
	}

	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.CommonConfig) {
		query["CommonConfig"] = request.CommonConfig
	}

	if !dara.IsNil(request.LayerOrderConfigList) {
		query["LayerOrderConfigList"] = request.LayerOrderConfigList
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.LayoutName) {
		query["LayoutName"] = request.LayoutName
	}

	if !dara.IsNil(request.MediaInputConfigList) {
		query["MediaInputConfigList"] = request.MediaInputConfigList
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScreenInputConfigList) {
		query["ScreenInputConfigList"] = request.ScreenInputConfigList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyStudioLayout"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyStudioLayoutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call MuteAllGroupUser to mute a message group, enabling group-wide muting.
//
// Description:
//
// ## QPS limits
//
// The QPS limit for this API is 100 queries per second (QPS) per user. If you exceed this limit, API calls will be throttled, which may affect your business. You can make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - MuteAllGroupUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MuteAllGroupUserResponse
func (client *Client) MuteAllGroupUserWithContext(ctx context.Context, request *MuteAllGroupUserRequest, runtime *dara.RuntimeOptions) (_result *MuteAllGroupUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BroadCastType) {
		body["BroadCastType"] = request.BroadCastType
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OperatorUserId) {
		body["OperatorUserId"] = request.OperatorUserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MuteAllGroupUser"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MuteAllGroupUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Mutes members in a message group in batches.
//
// Description:
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/343507.html).
//
// @param tmpReq - MuteGroupUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MuteGroupUserResponse
func (client *Client) MuteGroupUserWithContext(ctx context.Context, tmpReq *MuteGroupUserRequest, runtime *dara.RuntimeOptions) (_result *MuteGroupUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &MuteGroupUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MuteUserList) {
		request.MuteUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MuteUserList, dara.String("MuteUserList"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BroadCastType) {
		body["BroadCastType"] = request.BroadCastType
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MuteTime) {
		body["MuteTime"] = request.MuteTime
	}

	if !dara.IsNil(request.MuteUserListShrink) {
		body["MuteUserList"] = request.MuteUserListShrink
	}

	if !dara.IsNil(request.OperatorUserId) {
		body["OperatorUserId"] = request.OperatorUserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MuteGroupUser"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MuteGroupUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the OpenLiveShift operation to enable time shifting for a specific domain name, application, or stream.
//
// Description:
//
// You cannot configure time shifting and delayed transcoding at the same time.
//
// ## QPS limit
//
// This operation supports up to 10 queries per second (QPS) per user. If you exceed this limit, the system throttles your API calls, which can impact your business. To prevent throttling, call this operation at a reasonable rate.
//
// @param request - OpenLiveShiftRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenLiveShiftResponse
func (client *Client) OpenLiveShiftWithContext(ctx context.Context, request *OpenLiveShiftRequest, runtime *dara.RuntimeOptions) (_result *OpenLiveShiftResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.IgnoreTranscode) {
		query["IgnoreTranscode"] = request.IgnoreTranscode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.Vision) {
		query["Vision"] = request.Vision
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenLiveShift"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenLiveShiftResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually switches a show.
//
// Description:
//
// After you add a show and start live streaming, you can call this operation to switch the show. To add a show by using an API operation, see [Add a show to the show list](https://help.aliyun.com/document_detail/2848051.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - PlayChoosenShowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PlayChoosenShowResponse
func (client *Client) PlayChoosenShowWithContext(ctx context.Context, request *PlayChoosenShowRequest, runtime *dara.RuntimeOptions) (_result *PlayChoosenShowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ShowId) {
		query["ShowId"] = request.ShowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PlayChoosenShow"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PlayChoosenShowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes the configurations of an accelerated domain name from the canary release environment to the production environment.
//
// Description:
//
// ## Usage notes
//
// You can call the [DescribeLiveDomainStagingConfig](https://help.aliyun.com/document_detail/297374.html) operation to obtain the feature name. Then, you can call this operation to publish the configurations from the canary release environment to the production environment.
//
// ## QPS limits
//
// You can call this operation up to 30 times per second per user. Throttling is triggered if the number of calls exceeds the limit. This may affect your business operations. We recommend that you plan your calls accordingly. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - PublishLiveStagingConfigToProductionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishLiveStagingConfigToProductionResponse
func (client *Client) PublishLiveStagingConfigToProductionWithContext(ctx context.Context, request *PublishLiveStagingConfigToProductionRequest, runtime *dara.RuntimeOptions) (_result *PublishLiveStagingConfigToProductionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishLiveStagingConfigToProduction"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishLiveStagingConfigToProductionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于修改指定直播流的录制文件存储时长。
//
// Description:
//
// ## 请求说明
//
// - 该接口允许用户为一个或多个指定的直播流设置新的录制文件存储期限。
//
// - `Tag` 字段必须符合格式 `[0-9]+days`，表示直播结束后录制内容将被保存的天数。
//
// - 如果对某个流的存储时间修改失败，错误信息会被记录在返回结果中。对于失败的情况，调用方应重试最多3次；如果超过重试次数仍失败，则视为最终失败。
//
// - 为了支持未来可能的需求变化（如更长的存储周期），请确保您的系统能够处理不同的时间段值。
//
// - 成功执行后，供应商会通过异步回调的方式通知调用方所有操作的结果。若回调失败，将按照1小时、2小时、4小时的时间间隔尝试重新发送，直至成功或达到最大重试次数。
//
// @param tmpReq - PutRecordStorageLifeCycleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutRecordStorageLifeCycleResponse
func (client *Client) PutRecordStorageLifeCycleWithContext(ctx context.Context, tmpReq *PutRecordStorageLifeCycleRequest, runtime *dara.RuntimeOptions) (_result *PutRecordStorageLifeCycleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PutRecordStorageLifeCycleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StreamIds) {
		request.StreamIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamIds, dara.String("StreamIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.StreamIdsShrink) {
		body["StreamIds"] = request.StreamIdsShrink
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UnixTimestamp) {
		body["UnixTimestamp"] = request.UnixTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutRecordStorageLifeCycle"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutRecordStorageLifeCycleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries dual-stream disaster recovery online records.
//
// @param request - QueryLiveDomainMultiStreamListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLiveDomainMultiStreamListResponse
func (client *Client) QueryLiveDomainMultiStreamListWithContext(ctx context.Context, request *QueryLiveDomainMultiStreamListRequest, runtime *dara.RuntimeOptions) (_result *QueryLiveDomainMultiStreamListResponse, _err error) {
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
		Action:      dara.String("QueryLiveDomainMultiStreamList"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryLiveDomainMultiStreamListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call QueryMessageApp to query interactive message applications based on specified conditions.
//
// Description:
//
// ## QPS limits
//
// The QPS limit for this API is 100 queries per second per user. If the limit is exceeded, API calls will be throttled, which may affect your business. You can call the API properly to avoid this issue. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - QueryMessageAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMessageAppResponse
func (client *Client) QueryMessageAppWithContext(ctx context.Context, request *QueryMessageAppRequest, runtime *dara.RuntimeOptions) (_result *QueryMessageAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortType) {
		body["SortType"] = request.SortType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMessageApp"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMessageAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of created real-time speech-to-text or translation tasks.
//
// Description:
//
// The queries per second (QPS) limit for a single user is 20 calls per second. If you exceed this limit, your API calls are throttled, which can affect your business. Call this API only as needed.
//
// @param request - QueryRtcAsrTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRtcAsrTasksResponse
func (client *Client) QueryRtcAsrTasksWithContext(ctx context.Context, request *QueryRtcAsrTasksRequest, runtime *dara.RuntimeOptions) (_result *QueryRtcAsrTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRtcAsrTasks"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRtcAsrTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authentication configuration for snapshot callbacks.
//
// Description:
//
// You can call this operation to query the authentication configuration for snapshot callbacks for a streaming domain. Before you call this operation, you must configure authentication. For more information, see [Set snapshot callback authentication](https://help.aliyun.com/document_detail/2847907.html).
//
// ## QPS limit
//
// This operation is limited to 30 queries per second (QPS) per user. If you exceed this limit, throttling is triggered, which can affect your services.
//
// @param request - QuerySnapshotCallbackAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySnapshotCallbackAuthResponse
func (client *Client) QuerySnapshotCallbackAuthWithContext(ctx context.Context, request *QuerySnapshotCallbackAuthRequest, runtime *dara.RuntimeOptions) (_result *QuerySnapshotCallbackAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySnapshotCallbackAuth"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySnapshotCallbackAuthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs on-demand manual recording. For example, dynamically starts or stops recording.
//
// Description:
//
// - Before using this operation, make sure that you fully understand the billing method and pricing of live stream recording. For billing details, see [Live stream recording fees](https://help.aliyun.com/document_detail/195287.html).
//
// - This operation can only control live streams for which a recording configuration has taken effect. Complete [adding a recording configuration](https://help.aliyun.com/document_detail/2847881.html) first.
//
// - Before calling this operation, make sure that the target stream (DomainName/AppName/StreamName) is in an active stream ingest state.
//
// - If a live stream is being recorded (through automatic recording or manual recording), you can call this operation to stop recording the stream. However, if you call this operation to start recording when recording is already started, the TaskAlreadyStarted error is returned, indicating that the task has already been started.
//
// - If a live stream that is manually started for recording is interrupted, recording stops. If automatic recording is not configured, recording does not automatically start after stream ingest resumes.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - RealTimeRecordCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RealTimeRecordCommandResponse
func (client *Client) RealTimeRecordCommandWithContext(ctx context.Context, request *RealTimeRecordCommandRequest, runtime *dara.RuntimeOptions) (_result *RealTimeRecordCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Command) {
		query["Command"] = request.Command
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RealTimeRecordCommand"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RealTimeRecordCommandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a deleted interactive messaging group.
//
// Description:
//
//	  You can call this operation to restore a deleted interactive messaging group within 30 days after you call the [DeleteLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html) to delete the group.
//
//		- After you restore a group, the messages that were stored in the group before it was deleted can still be queried.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - RecoverLiveMessageDeletedGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecoverLiveMessageDeletedGroupResponse
func (client *Client) RecoverLiveMessageDeletedGroupWithContext(ctx context.Context, request *RecoverLiveMessageDeletedGroupRequest, runtime *dara.RuntimeOptions) (_result *RecoverLiveMessageDeletedGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecoverLiveMessageDeletedGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecoverLiveMessageDeletedGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unmutes users in a live interactive message group.
//
// Description:
//
// You must call [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html) to create an interactive message group before you call this operation.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. Exceeding this limit triggers API throttling, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param tmpReq - RemoveLiveMessageGroupBandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveLiveMessageGroupBandResponse
func (client *Client) RemoveLiveMessageGroupBandWithContext(ctx context.Context, tmpReq *RemoveLiveMessageGroupBandRequest, runtime *dara.RuntimeOptions) (_result *RemoveLiveMessageGroupBandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveLiveMessageGroupBandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UnbannedUsers) {
		request.UnbannedUsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UnbannedUsers, dara.String("UnbannedUsers"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.UnbannedUsersShrink) {
		query["UnbannedUsers"] = request.UnbannedUsersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveLiveMessageGroupBand"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveLiveMessageGroupBandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an episode from an episode list.
//
// Description:
//
// This operation deletes a show from a show list. You must first call the [AddShowIntoShowList](https://help.aliyun.com/document_detail/2848051.html) operation to add the show.
//
// ## QPS limit
//
// The limit for this operation is 10 queries per second (QPS) per user. API calls that exceed this limit are throttled, which may affect your business. Plan your API calls to avoid exceeding this limit.
//
// @param request - RemoveShowFromShowListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveShowFromShowListResponse
func (client *Client) RemoveShowFromShowListWithContext(ctx context.Context, request *RemoveShowFromShowListRequest, runtime *dara.RuntimeOptions) (_result *RemoveShowFromShowListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ShowId) {
		query["ShowId"] = request.ShowId
	}

	if !dara.IsNil(request.IsBatchMode) {
		query["isBatchMode"] = request.IsBatchMode
	}

	if !dara.IsNil(request.ShowIdList) {
		query["showIdList"] = request.ShowIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveShowFromShowList"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveShowFromShowListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes users from a channel.
//
// Description:
//
// You can call this operation to remove one or more users from a channel.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - RemoveTerminalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTerminalsResponse
func (client *Client) RemoveTerminalsWithContext(ctx context.Context, request *RemoveTerminalsRequest, runtime *dara.RuntimeOptions) (_result *RemoveTerminalsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TerminalIds) {
		query["TerminalIds"] = request.TerminalIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTerminals"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTerminalsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a production studio.
//
// Description:
//
// - This operation supports only production studios in playlist mode (carousel) or general mode. Virtual studios are not supported.
//
// - When you restart a production studio, its current settings, such as resolution and screen orientation, are reloaded to restore the previous playback status.
//
// ## QPS limits
//
// This operation is limited to 10 queries per second (QPS) per user. Calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - RestartCasterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartCasterResponse
func (client *Client) RestartCasterWithContext(ctx context.Context, request *RestartCasterRequest, runtime *dara.RuntimeOptions) (_result *RestartCasterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartCaster"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartCasterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a stream pulling task.
//
// Description:
//
// - Calls this operation to restart a stream pulling task.
//
// - You can restart a task that is running (including tasks in abnormal retry status) or stopped. A running task is stopped and then started again. A non-running task is started directly.
//
// - If the task has not reached the configured start time, the restart does not take effect.
//
// - The restarted task runs based on the latest task configuration, which causes stream ingest interruption.
//
// - After a video-on-demand task is restarted, playback starts from the beginning based on the latest playlist. You can call the update operation to set the video index and video playback progress to resume playback.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - RestartLivePullToPushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartLivePullToPushResponse
func (client *Client) RestartLivePullToPushWithContext(ctx context.Context, request *RestartLivePullToPushRequest, runtime *dara.RuntimeOptions) (_result *RestartLivePullToPushResponse, _err error) {
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
		Action:      dara.String("RestartLivePullToPush"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartLivePullToPushResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重启转码任务
//
// @param request - RestartTranscodeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartTranscodeTaskResponse
func (client *Client) RestartTranscodeTaskWithContext(ctx context.Context, request *RestartTranscodeTaskRequest, runtime *dara.RuntimeOptions) (_result *RestartTranscodeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.PushDomain) {
		query["PushDomain"] = request.PushDomain
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.TranscodingTemplate) {
		query["TranscodingTemplate"] = request.TranscodingTemplate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartTranscodeTask"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartTranscodeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes stream ingest for a specified stream.
//
// Description:
//
// Calls this operation to resume stream ingest for a specified stream. This operation currently supports only publisher (streamer ingest).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 30 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - ResumeLiveStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeLiveStreamResponse
func (client *Client) ResumeLiveStreamWithContext(ctx context.Context, request *ResumeLiveStreamRequest, runtime *dara.RuntimeOptions) (_result *ResumeLiveStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.LiveStreamType) {
		query["LiveStreamType"] = request.LiveStreamType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeLiveStream"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeLiveStreamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back the configurations of an accelerated domain name in the canary release environment.
//
// Description:
//
// ## Usage notes
//
// You must first call the [DescribeLiveDomainStagingConfig](https://help.aliyun.com/document_detail/297374.html) operation to obtain the feature name. Then, you can call this operation to roll back the configurations in the canary release environment.
//
// ## QPS limits
//
// You can call this operation up to 30 times per second per user. API calls that exceed this limit are throttled. This can affect your business. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - RollbackLiveStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackLiveStagingConfigResponse
func (client *Client) RollbackLiveStagingConfigWithContext(ctx context.Context, request *RollbackLiveStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *RollbackLiveStagingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackLiveStagingConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackLiveStagingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call SendLike to configure system like messages.
//
// Description:
//
// ## QPS limits
//
// The QPS limit for this API is 100 queries per second (QPS) per user. If the limit is exceeded, API calls will be throttled, which may affect your business. You can call the API properly to avoid this issue. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - SendLikeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendLikeResponse
func (client *Client) SendLikeWithContext(ctx context.Context, request *SendLikeRequest, runtime *dara.RuntimeOptions) (_result *SendLikeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BroadCastType) {
		body["BroadCastType"] = request.BroadCastType
	}

	if !dara.IsNil(request.Count) {
		body["Count"] = request.Count
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OperatorUserId) {
		body["OperatorUserId"] = request.OperatorUserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendLike"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendLikeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a message to a group.
//
// Description:
//
// Before you call this operation, you must call [](t2586945.xdita#)to create an interactive messaging group. You can send messages to a group only if the group is active, which means that one or more users have joined the group. Offline messages are not supported. If you fail to send a message, check whether there are users in the group. If you want to send an offline message, we recommend that you store the message locally and send it after users come online.
//
// ## QPS limits
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you consider this limit when calling this operation.
//
// @param request - SendLiveMessageGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendLiveMessageGroupResponse
func (client *Client) SendLiveMessageGroupWithContext(ctx context.Context, request *SendLiveMessageGroupRequest, runtime *dara.RuntimeOptions) (_result *SendLiveMessageGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MsgTid) {
		query["MsgTid"] = request.MsgTid
	}

	if !dara.IsNil(request.MsgType) {
		query["MsgType"] = request.MsgType
	}

	if !dara.IsNil(request.NoCache) {
		query["NoCache"] = request.NoCache
	}

	if !dara.IsNil(request.NoStorage) {
		query["NoStorage"] = request.NoStorage
	}

	if !dara.IsNil(request.SenderId) {
		query["SenderId"] = request.SenderId
	}

	if !dara.IsNil(request.SenderMetaInfo) {
		query["SenderMetaInfo"] = request.SenderMetaInfo
	}

	if !dara.IsNil(request.StaticsIncrease) {
		query["StaticsIncrease"] = request.StaticsIncrease
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendLiveMessageGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendLiveMessageGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a message to a specified user. The user is identified by ReceiverId.
//
// Description:
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - SendLiveMessageUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendLiveMessageUserResponse
func (client *Client) SendLiveMessageUserWithContext(ctx context.Context, request *SendLiveMessageUserRequest, runtime *dara.RuntimeOptions) (_result *SendLiveMessageUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.HighReliability) {
		query["HighReliability"] = request.HighReliability
	}

	if !dara.IsNil(request.MsgTid) {
		query["MsgTid"] = request.MsgTid
	}

	if !dara.IsNil(request.MsgType) {
		query["MsgType"] = request.MsgType
	}

	if !dara.IsNil(request.ReceiverId) {
		query["ReceiverId"] = request.ReceiverId
	}

	if !dara.IsNil(request.SenderId) {
		query["SenderId"] = request.SenderId
	}

	if !dara.IsNil(request.SenderInfo) {
		query["SenderInfo"] = request.SenderInfo
	}

	if !dara.IsNil(request.Storage) {
		query["Storage"] = request.Storage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendLiveMessageUser"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendLiveMessageUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call SendMessageToGroup to send messages to all members of a message group.
//
// Description:
//
// ## QPS limits
//
// The single-user QPS limit for this API is 100 queries per second (QPS). If the limit is exceeded, API calls will be throttled, which may affect your business. You can call the API properly to avoid this issue. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - SendMessageToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageToGroupResponse
func (client *Client) SendMessageToGroupWithContext(ctx context.Context, request *SendMessageToGroupRequest, runtime *dara.RuntimeOptions) (_result *SendMessageToGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SkipAudit) {
		query["SkipAudit"] = request.SkipAudit
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OperatorUserId) {
		body["OperatorUserId"] = request.OperatorUserId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendMessageToGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendMessageToGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call SendMessageToGroupUsers to send messages to specified users in a message group.
//
// Description:
//
// ## QPS limits
//
// The single-user QPS limit for this API is 100 queries per second (QPS). If this limit is exceeded, API calls will be throttled, which may affect your business. You can make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param tmpReq - SendMessageToGroupUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageToGroupUsersResponse
func (client *Client) SendMessageToGroupUsersWithContext(ctx context.Context, tmpReq *SendMessageToGroupUsersRequest, runtime *dara.RuntimeOptions) (_result *SendMessageToGroupUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendMessageToGroupUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReceiverIdList) {
		request.ReceiverIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReceiverIdList, dara.String("ReceiverIdList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SkipAudit) {
		query["SkipAudit"] = request.SkipAudit
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OperatorUserId) {
		body["OperatorUserId"] = request.OperatorUserId
	}

	if !dara.IsNil(request.ReceiverIdListShrink) {
		body["ReceiverIdList"] = request.ReceiverIdListShrink
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendMessageToGroupUsers"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendMessageToGroupUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets a channel for a production studio.
//
// Description:
//
// Call the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation to create a production studio. You can then call this operation to set a channel for the production studio.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 for each user. Exceeding this limit triggers throttling, which may affect your business. Call this operation at a reasonable rate to prevent interruptions.
//
// @param request - SetCasterChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCasterChannelResponse
func (client *Client) SetCasterChannelWithContext(ctx context.Context, request *SetCasterChannelRequest, runtime *dara.RuntimeOptions) (_result *SetCasterChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.FaceBeauty) {
		query["FaceBeauty"] = request.FaceBeauty
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayStatus) {
		query["PlayStatus"] = request.PlayStatus
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.SeekOffset) {
		query["SeekOffset"] = request.SeekOffset
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCasterChannel"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCasterChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures detailed settings for a production studio, including the name, transcoding configuration, recording configuration, and other parameters.
//
// Description:
//
// Create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation first, and then call this operation to configure detailed settings for the production studio.
//
//	Warning: This operation fully replaces the existing configuration. If you set a parameter to empty, the existing configuration of that parameter in the production studio is cleared.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - SetCasterConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCasterConfigResponse
func (client *Client) SetCasterConfigWithContext(ctx context.Context, request *SetCasterConfigRequest, runtime *dara.RuntimeOptions) (_result *SetCasterConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoSwitchUrgentConfig) {
		query["AutoSwitchUrgentConfig"] = request.AutoSwitchUrgentConfig
	}

	if !dara.IsNil(request.AutoSwitchUrgentOn) {
		query["AutoSwitchUrgentOn"] = request.AutoSwitchUrgentOn
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.CasterName) {
		query["CasterName"] = request.CasterName
	}

	if !dara.IsNil(request.ChannelEnable) {
		query["ChannelEnable"] = request.ChannelEnable
	}

	if !dara.IsNil(request.Delay) {
		query["Delay"] = request.Delay
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProgramEffect) {
		query["ProgramEffect"] = request.ProgramEffect
	}

	if !dara.IsNil(request.ProgramName) {
		query["ProgramName"] = request.ProgramName
	}

	if !dara.IsNil(request.RecordConfig) {
		query["RecordConfig"] = request.RecordConfig
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SideOutputUrl) {
		query["SideOutputUrl"] = request.SideOutputUrl
	}

	if !dara.IsNil(request.SideOutputUrlList) {
		query["SideOutputUrlList"] = request.SideOutputUrlList
	}

	if !dara.IsNil(request.SyncGroupsConfig) {
		query["SyncGroupsConfig"] = request.SyncGroupsConfig
	}

	if !dara.IsNil(request.TranscodeConfig) {
		query["TranscodeConfig"] = request.TranscodeConfig
	}

	if !dara.IsNil(request.UrgentImageId) {
		query["UrgentImageId"] = request.UrgentImageId
	}

	if !dara.IsNil(request.UrgentImageUrl) {
		query["UrgentImageUrl"] = request.UrgentImageUrl
	}

	if !dara.IsNil(request.UrgentLiveStreamUrl) {
		query["UrgentLiveStreamUrl"] = request.UrgentLiveStreamUrl
	}

	if !dara.IsNil(request.UrgentMaterialId) {
		query["UrgentMaterialId"] = request.UrgentMaterialId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCasterConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCasterConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the full scene configuration for a production studio by clearing the existing scene configuration and applying layout information to the specified scene.
//
// Description:
//
// Calls this operation to set the full scene configuration for a production studio by clearing the existing scene configuration and applying layout information to the specified scene.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - SetCasterSceneConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCasterSceneConfigResponse
func (client *Client) SetCasterSceneConfigWithContext(ctx context.Context, request *SetCasterSceneConfigRequest, runtime *dara.RuntimeOptions) (_result *SetCasterSceneConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.ComponentId) {
		query["ComponentId"] = request.ComponentId
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCasterSceneConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCasterSceneConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures whether the certificate feature is enabled for a domain name and modifies certificate information.
//
// Description:
//
// Obtain the live streaming domain name first, and then call this operation to configure whether the certificate is enabled for the domain name and modify certificate information.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 30 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - SetLiveDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLiveDomainCertificateResponse
func (client *Client) SetLiveDomainCertificateWithContext(ctx context.Context, request *SetLiveDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetLiveDomainCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.CertType) {
		query["CertType"] = request.CertType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ForceSet) {
		query["ForceSet"] = request.ForceSet
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SSLPri) {
		query["SSLPri"] = request.SSLPri
	}

	if !dara.IsNil(request.SSLProtocol) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !dara.IsNil(request.SSLPub) {
		query["SSLPub"] = request.SSLPub
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLiveDomainCertificate"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLiveDomainCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls SetLiveDomainMultiStreamConfig to enable or disable the dual-stream disaster recovery switch.
//
// Description:
//
// Calls this operation to enable the dual-stream disaster recovery feature, which allows stream ingest to the same live stream name.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - SetLiveDomainMultiStreamConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLiveDomainMultiStreamConfigResponse
func (client *Client) SetLiveDomainMultiStreamConfigWithContext(ctx context.Context, request *SetLiveDomainMultiStreamConfigRequest, runtime *dara.RuntimeOptions) (_result *SetLiveDomainMultiStreamConfigResponse, _err error) {
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
		Action:      dara.String("SetLiveDomainMultiStreamConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLiveDomainMultiStreamConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually switches between the active stream and standby stream.
//
// @param request - SetLiveDomainMultiStreamMasterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLiveDomainMultiStreamMasterResponse
func (client *Client) SetLiveDomainMultiStreamMasterWithContext(ctx context.Context, request *SetLiveDomainMultiStreamMasterRequest, runtime *dara.RuntimeOptions) (_result *SetLiveDomainMultiStreamMasterResponse, _err error) {
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
		Action:      dara.String("SetLiveDomainMultiStreamMaster"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLiveDomainMultiStreamMasterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the automatic mode switch for dual-stream disaster recovery.
//
// @param request - SetLiveDomainMultiStreamOptimalModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLiveDomainMultiStreamOptimalModeResponse
func (client *Client) SetLiveDomainMultiStreamOptimalModeWithContext(ctx context.Context, request *SetLiveDomainMultiStreamOptimalModeRequest, runtime *dara.RuntimeOptions) (_result *SetLiveDomainMultiStreamOptimalModeResponse, _err error) {
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
		Action:      dara.String("SetLiveDomainMultiStreamOptimalMode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLiveDomainMultiStreamOptimalModeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets or modifies domain configurations in the canary release environment.
//
// Description:
//
// ## Usage notes
//
// Obtain an accelerated domain name. You can then call this operation to set or modify domain name configurations in the staging environment. For more information, see **Functions format description**.
//
// ## QPS limit
//
// This operation is limited to 30 queries per second (QPS) per user. API calls that exceed this limit are throttled, which may affect your business. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - SetLiveDomainStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLiveDomainStagingConfigResponse
func (client *Client) SetLiveDomainStagingConfigWithContext(ctx context.Context, request *SetLiveDomainStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *SetLiveDomainStagingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Functions) {
		query["Functions"] = request.Functions
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLiveDomainStagingConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLiveDomainStagingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls SetLiveEdgeTransfer to configure live stream relay settings.
//
// Description:
//
// The live edge stream relay configuration made by calling SetLiveEdgeTransfer only takes effect for streams that start after the configuration is completed. The following are examples of some typical scenarios:
//
// | Scenario | Analysis | Result |
//
// | -------------- | -------------- | ------ |
//
// | 1. The user has already started streaming before calling SetLiveEdgeTransfer. | The live edge stream relay configuration does not exist at this point. | The stream is not affected by the SetLiveEdgeTransfer configuration, meaning live edge stream relay will not be initiated. |
//
// | 2. The user interrupts streaming that was started before calling SetLiveEdgeTransfer, and then resumes streaming. | The live edge stream relay configuration already exists at this point. | The resumed stream will initiate live stream relay according to the SetLiveEdgeTransfer configuration. |
//
// | 3. The user starts streaming after calling SetLiveEdgeTransfer. | The live edge stream relay configuration already exists at this point. | This stream will initiate live stream relay according to the SetLiveEdgeTransfer configuration. |
//
// ## QPS Limit
//
// The single-user QPS limit for this API is 100 calls/second. Exceeding the limit will result in API throttling, which may affect your business. Please call this API appropriately.
//
// @param request - SetLiveEdgeTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLiveEdgeTransferResponse
func (client *Client) SetLiveEdgeTransferWithContext(ctx context.Context, request *SetLiveEdgeTransferRequest, runtime *dara.RuntimeOptions) (_result *SetLiveEdgeTransferResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.HttpDns) {
		query["HttpDns"] = request.HttpDns
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.TargetDomainList) {
		query["TargetDomainList"] = request.TargetDomainList
	}

	if !dara.IsNil(request.TransferArgs) {
		query["TransferArgs"] = request.TransferArgs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLiveEdgeTransfer"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLiveEdgeTransferResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures triggered stream pulling.
//
// Description:
//
// This API configures triggered stream pulling. When a live stream starts on the origin server, ApsaraVideo Live automatically pulls the stream for live playback.
//
// > This API does not support the IPv6 protocol.
//
// ## QPS limits
//
// You can call this operation up to 1,000 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - SetLiveLazyPullStreamInfoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLiveLazyPullStreamInfoConfigResponse
func (client *Client) SetLiveLazyPullStreamInfoConfigWithContext(ctx context.Context, request *SetLiveLazyPullStreamInfoConfigRequest, runtime *dara.RuntimeOptions) (_result *SetLiveLazyPullStreamInfoConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PullAppName) {
		query["PullAppName"] = request.PullAppName
	}

	if !dara.IsNil(request.PullDomainName) {
		query["PullDomainName"] = request.PullDomainName
	}

	if !dara.IsNil(request.PullProtocol) {
		query["PullProtocol"] = request.PullProtocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TranscodeLazy) {
		query["TranscodeLazy"] = request.TranscodeLazy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLiveLazyPullStreamInfoConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLiveLazyPullStreamInfoConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures custom supplemental enhancement information (SEI) for co-streaming.
//
// Description:
//
// ## [](#)Usage notes
//
// You can call this operation to configure custom SEI for a mixed-stream relay task.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - SetLiveMpuTaskSeiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLiveMpuTaskSeiResponse
func (client *Client) SetLiveMpuTaskSeiWithContext(ctx context.Context, request *SetLiveMpuTaskSeiRequest, runtime *dara.RuntimeOptions) (_result *SetLiveMpuTaskSeiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CustomSei) {
		query["CustomSei"] = request.CustomSei
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLiveMpuTaskSei"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLiveMpuTaskSeiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets stream-level blocking for live streams by calling SetLiveStreamBlock.
//
// Description:
//
// You can configure up to 200 live stream area blocking rules for a domain name. If duplicate rules with the same AppName and StreamName exist in the live stream area blocking rules, the most recently updated rule takes effect.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If this limit is exceeded, the API invocations are throttled, which may affect your business. Invoke this operation appropriately.
//
// @param request - SetLiveStreamBlockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLiveStreamBlockResponse
func (client *Client) SetLiveStreamBlockWithContext(ctx context.Context, request *SetLiveStreamBlockRequest, runtime *dara.RuntimeOptions) (_result *SetLiveStreamBlockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BlockType) {
		query["BlockType"] = request.BlockType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.LocationList) {
		query["LocationList"] = request.LocationList
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReleaseTime) {
		query["ReleaseTime"] = request.ReleaseTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLiveStreamBlock"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLiveStreamBlockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call SetLiveStreamDelayConfig to configure live streaming latency.
//
// Description:
//
// - This operation configures the latency for a streaming domain.
//
// - The server-side cache stores an integer number of Groups of Pictures (GOPs). Therefore, the latency cannot be less than the GOP size. The actual latency is estimated based on the GOP size. For example, if you set RtmpDelay to 4 seconds for an RTMP stream and the GOP size is 2 seconds, the actual latency ranges from 2 seconds (4 - 2) to 6 seconds (4 + 2). If the GOP size is larger than the configured RtmpDelay, for example, GOP = 5 seconds and RtmpDelay = 4 seconds, the latency fluctuates between 0 and 9 seconds.
//
// - Latency configurations do not take effect for audio-only streams. The default latency is close to 0 seconds.
//
// - For an HLS configuration, the segment size is calculated as \\`Delay / 3\\` and rounded down to the nearest integer. The minimum segment size is 1 second. The maximum number of segments is 4 if the segment size is 3 seconds or more. Otherwise, the maximum is 6 segments.
//
// - The actual HLS segment size cannot be smaller than the GOP size.
//
// - The HLS latency is calculated as: Configured Segment Size × 3.
//
// - If you do not call this operation, the system uses the default values. The default latency is 2 seconds for RTMP and 4 seconds for FLV. For HLS, the default segment size is 5 seconds, which results in a latency of 15 seconds with a maximum of 6 segments.
//
// ## Queries per second (QPS) limit
//
// The QPS limit for this operation is 1,000 calls per second per user. If you exceed this limit, API calls are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - SetLiveStreamDelayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLiveStreamDelayConfigResponse
func (client *Client) SetLiveStreamDelayConfigWithContext(ctx context.Context, request *SetLiveStreamDelayConfigRequest, runtime *dara.RuntimeOptions) (_result *SetLiveStreamDelayConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FlvDelay) {
		query["FlvDelay"] = request.FlvDelay
	}

	if !dara.IsNil(request.FlvLevel) {
		query["FlvLevel"] = request.FlvLevel
	}

	if !dara.IsNil(request.HlsDelay) {
		query["HlsDelay"] = request.HlsDelay
	}

	if !dara.IsNil(request.HlsLevel) {
		query["HlsLevel"] = request.HlsLevel
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RtmpDelay) {
		query["RtmpDelay"] = request.RtmpDelay
	}

	if !dara.IsNil(request.RtmpLevel) {
		query["RtmpLevel"] = request.RtmpLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLiveStreamDelayConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLiveStreamDelayConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets live stream prefetch tasks. You can prefetch up to 100 live streams in a batch.
//
// Description:
//
// - You can call this operation to set live stream prefetch tasks in batches. You can set up to 100 live stream URLs at a time.
//
// - Live stream prefetch does not support HLS URLs.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - SetLiveStreamPreloadTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLiveStreamPreloadTasksResponse
func (client *Client) SetLiveStreamPreloadTasksWithContext(ctx context.Context, request *SetLiveStreamPreloadTasksRequest, runtime *dara.RuntimeOptions) (_result *SetLiveStreamPreloadTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayUrl) {
		query["PlayUrl"] = request.PlayUrl
	}

	if !dara.IsNil(request.PreloadedEndTime) {
		query["PreloadedEndTime"] = request.PreloadedEndTime
	}

	if !dara.IsNil(request.PreloadedStartTime) {
		query["PreloadedStartTime"] = request.PreloadedStartTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLiveStreamPreloadTasks"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLiveStreamPreloadTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the stream ingest callback configuration.
//
// Description:
//
// - This operation allows you to configure the callback URL and authentication information for an ingest domain.
//
// - The real-time stream status callback promptly notifies you of the results of stream ingest or stream disconnection operations. For more information, see [Stream ingest callback format description](https://help.aliyun.com/document_detail/54787.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 15 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - SetLiveStreamsNotifyUrlConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLiveStreamsNotifyUrlConfigResponse
func (client *Client) SetLiveStreamsNotifyUrlConfigWithContext(ctx context.Context, request *SetLiveStreamsNotifyUrlConfigRequest, runtime *dara.RuntimeOptions) (_result *SetLiveStreamsNotifyUrlConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ExceptionNotifyUrl) {
		query["ExceptionNotifyUrl"] = request.ExceptionNotifyUrl
	}

	if !dara.IsNil(request.NotifyAuthKey) {
		query["NotifyAuthKey"] = request.NotifyAuthKey
	}

	if !dara.IsNil(request.NotifyReqAuth) {
		query["NotifyReqAuth"] = request.NotifyReqAuth
	}

	if !dara.IsNil(request.NotifyUrl) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SwitchNotifyUrl) {
		query["SwitchNotifyUrl"] = request.SwitchNotifyUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLiveStreamsNotifyUrlConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLiveStreamsNotifyUrlConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the background for a playlist mode production studio.
//
// Description:
//
// - Create a playlist mode production studio before calling this operation to add background materials. To create a production studio by using an API operation, see [CreateCaster](https://help.aliyun.com/document_detail/2848009.html).
//
// - You can use this operation to create, update, or delete a playlist background. To delete the background, leave the ResourceType, ResourceUrl, and MaterialId parameters empty.
//
//	Notice:
//
// - When using ApsaraVideo VOD resources, use managed bucket resources first. Resources in your own bucket may expire. If you use resources in your own bucket, check the resource validity period.
//
// - Use ApsaraVideo Live and ApsaraVideo VOD resources as material input first. Resources from third-party URLs may fail to play. Verify the quality and validity of such resources.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - SetShowListBackgroundRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetShowListBackgroundResponse
func (client *Client) SetShowListBackgroundWithContext(ctx context.Context, request *SetShowListBackgroundRequest, runtime *dara.RuntimeOptions) (_result *SetShowListBackgroundResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.MaterialId) {
		query["MaterialId"] = request.MaterialId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ResourceUrl) {
		query["ResourceUrl"] = request.ResourceUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetShowListBackground"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetShowListBackgroundResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures authentication for snapshot callbacks.
//
// Description:
//
// - Before calling this API, configure a callback URL first. For more information, see [AddLiveAppSnapshotConfig](https://help.aliyun.com/document_detail/2847897.html).
//
// - The snapshot service supports adding a signature header to HTTP or HTTPS callback requests. This allows the receiving server to authenticate the signature and prevent unauthorized or invalid requests. For more information, see [Usage notes for callback authentication](https://help.aliyun.com/document_detail/417349.html).
//
// ## QPS limits
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - SetSnapshotCallbackAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSnapshotCallbackAuthResponse
func (client *Client) SetSnapshotCallbackAuthWithContext(ctx context.Context, request *SetSnapshotCallbackAuthRequest, runtime *dara.RuntimeOptions) (_result *SetSnapshotCallbackAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallbackAuthKey) {
		query["CallbackAuthKey"] = request.CallbackAuthKey
	}

	if !dara.IsNil(request.CallbackReqAuth) {
		query["CallbackReqAuth"] = request.CallbackReqAuth
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSnapshotCallbackAuth"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSnapshotCallbackAuthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a production studio. If the PVW and PGM scenes do not exist, they are created. The PVW and PGM scenes are started, and the underlying audio and video processing tasks are initiated.
//
// Description:
//
// Create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, and then call this operation to start the production studio. If the PVW and PGM scenes do not exist, they are created. The PVW and PGM scenes are started, and the underlying audio and video processing tasks are initiated.
//
// Before calling this operation, call SetCasterConfig to configure DomainName. If DomainName is not configured, the error InvalidDomainName.NotFound is returned.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as appropriate.
//
// @param request - StartCasterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCasterResponse
func (client *Client) StartCasterWithContext(ctx context.Context, request *StartCasterRequest, runtime *dara.RuntimeOptions) (_result *StartCasterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartCaster"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartCasterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a specified Program View (PVW) scene. This operation is used to start preview scenes.
//
// Description:
//
// You can call this operation to start a specified Program View (PVW) scene. A PVW scene is a preview scene.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 10. If you exceed this limit, API calls are throttled. Throttling can affect your business. We recommend that you call this operation within the specified limit.
//
// @param request - StartCasterSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCasterSceneResponse
func (client *Client) StartCasterSceneWithContext(ctx context.Context, request *StartCasterSceneRequest, runtime *dara.RuntimeOptions) (_result *StartCasterSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartCasterScene"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartCasterSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an edge transcoding task.
//
// Description:
//
// - This operation is used to start an edge transcoding job.
//
// - Before you call this operation, make sure that you have the required permissions to access the edge transcoding service and that the specified transcoding job is not running.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 6,000 calls per minute for each user. API calls that exceed this limit are throttled, which may affect your business. We recommend that you call this operation as needed.
//
// @param request - StartEdgeTranscodeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartEdgeTranscodeJobResponse
func (client *Client) StartEdgeTranscodeJobWithContext(ctx context.Context, request *StartEdgeTranscodeJobRequest, runtime *dara.RuntimeOptions) (_result *StartEdgeTranscodeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartEdgeTranscodeJob"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartEdgeTranscodeJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a disabled live streaming domain. This operation changes the DomainStatus to online.
//
// Description:
//
// This API operation fails if your account has an overdue payment or if the domain name is in an invalid state.
//
// ## QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - StartLiveDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartLiveDomainResponse
func (client *Client) StartLiveDomainWithContext(ctx context.Context, request *StartLiveDomainRequest, runtime *dara.RuntimeOptions) (_result *StartLiveDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartLiveDomain"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartLiveDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a stream mixing and transcoding task.
//
// Description:
//
// By default, each application ID supports a maximum of 200 single-stream ingest tasks and 40 stream mixing and transcoding tasks. To increase the quota, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket).
//
// ### Stream mixing task lifecycle
//
// **Start**
//
// - When a streamer starts streaming for the first time, you can call StartLiveMPUTask to start a bypass task.
//
//   - If no users are in the channel, a "channel does not exist" error is returned.
//
//   - The bypass stream is output only when a user starts stream ingest. If the user in a single-stream task does not ingest a stream, the bypass stream cannot be played.
//
//   - For a stream mixing task, at least one user must be ingesting a stream for the bypass stream to be playable. The layout area for users who are not ingesting streams shows a black screen.
//
// - You can record the bypass task status, task type, and task parameters on your business server.
//
//   - Task status: Started, Stopped.
//
//   - Task type: Single-stream, Stream mixing.
//
//   - Task parameters: The latest input parameters. For example, after a successful call to UpdateLiveMPUTask, record the latest task parameters.
//
// - In co-streaming or PK scenarios, if a task has been updated to a stream mixing task and the streamer unexpectedly leaves and then rejoins the channel, your business server can call StartLiveMPUTask to restart the stream mixing task based on the saved task type and parameters.
//
//   - If the system has not automatically cleared the task before you start it, the task starts successfully.
//
//   - If the system has not yet cleared the task, a **Task already exists*	- error code is returned.
//
// **End**
//
// - When a streamer leaves the channel, call [StopLiveMPUTask](https://help.aliyun.com/document_detail/2362742.html) to stop the bypass task.
//
// - If all users in the task leave the channel and StopLiveMPUTask is not called, the system automatically stops the bypass task after 2 minutes.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user for this API is 500 calls/second. If you exceed this limit, API calls are throttled. This may affect your business. We recommend that you call this API reasonably.
//
// @param tmpReq - StartLiveMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartLiveMPUTaskResponse
func (client *Client) StartLiveMPUTaskWithContext(ctx context.Context, tmpReq *StartLiveMPUTaskRequest, runtime *dara.RuntimeOptions) (_result *StartLiveMPUTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartLiveMPUTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MultiStreamURL) {
		request.MultiStreamURLShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MultiStreamURL, dara.String("MultiStreamURL"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SeiParams) {
		request.SeiParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SeiParams, dara.String("SeiParams"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SingleSubParams) {
		request.SingleSubParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SingleSubParams, dara.String("SingleSubParams"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TranscodeParams) {
		request.TranscodeParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TranscodeParams, dara.String("TranscodeParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.MaxIdleTime) {
		query["MaxIdleTime"] = request.MaxIdleTime
	}

	if !dara.IsNil(request.MixMode) {
		query["MixMode"] = request.MixMode
	}

	if !dara.IsNil(request.MultiStreamURLShrink) {
		query["MultiStreamURL"] = request.MultiStreamURLShrink
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.SeiParamsShrink) {
		query["SeiParams"] = request.SeiParamsShrink
	}

	if !dara.IsNil(request.SingleSubParamsShrink) {
		query["SingleSubParams"] = request.SingleSubParamsShrink
	}

	if !dara.IsNil(request.StreamURL) {
		query["StreamURL"] = request.StreamURL
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TranscodeParamsShrink) {
		query["TranscodeParams"] = request.TranscodeParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartLiveMPUTask"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartLiveMPUTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts live monitoring.
//
// Description:
//
// You can call the [CreateLiveStreamMonitor](https://help.aliyun.com/document_detail/2848129.html) operation to create a monitoring session. Then, you can use the **MonitorId*	- value from the response to start the session.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. Calls that exceed this limit are throttled. Throttling may affect your business. Plan your calls accordingly.
//
// @param request - StartLiveStreamMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartLiveStreamMonitorResponse
func (client *Client) StartLiveStreamMonitorWithContext(ctx context.Context, request *StartLiveStreamMonitorRequest, runtime *dara.RuntimeOptions) (_result *StartLiveStreamMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MonitorId) {
		query["MonitorId"] = request.MonitorId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartLiveStreamMonitor"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartLiveStreamMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a playlist.
//
// Description:
//
// Call the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation to add items to a playlist. Then, call this operation to start the playlist.
//
// ## QPS limit
//
// This operation has a queries per second (QPS) limit of 10 for each user. If you exceed this limit, your API calls are throttled, which may affect your business. Call this operation within the specified limit.
//
// @param request - StartPlaylistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartPlaylistResponse
func (client *Client) StartPlaylistWithContext(ctx context.Context, request *StartPlaylistRequest, runtime *dara.RuntimeOptions) (_result *StartPlaylistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProgramId) {
		query["ProgramId"] = request.ProgramId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResumeMode) {
		query["ResumeMode"] = request.ResumeMode
	}

	if !dara.IsNil(request.StartItemId) {
		query["StartItemId"] = request.StartItemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartPlaylist"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartPlaylistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an RTC cloud recording task.
//
// Description:
//
// Cloud recording is a paid feature. For billing details, see [Cloud recording fees](https://help.aliyun.com/document_detail/2976391.html).
//
// ## Service registration
//
// ## QPS limit.
//
// @param tmpReq - StartRtcCloudRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRtcCloudRecordingResponse
func (client *Client) StartRtcCloudRecordingWithContext(ctx context.Context, tmpReq *StartRtcCloudRecordingRequest, runtime *dara.RuntimeOptions) (_result *StartRtcCloudRecordingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartRtcCloudRecordingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MixLayoutParams) {
		request.MixLayoutParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MixLayoutParams, dara.String("MixLayoutParams"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MixTranscodeParams) {
		request.MixTranscodeParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MixTranscodeParams, dara.String("MixTranscodeParams"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RecordParams) {
		request.RecordParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordParams, dara.String("RecordParams"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StorageParams) {
		request.StorageParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StorageParams, dara.String("StorageParams"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SubscribeParams) {
		request.SubscribeParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubscribeParams, dara.String("SubscribeParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.MaxIdleTime) {
		query["MaxIdleTime"] = request.MaxIdleTime
	}

	if !dara.IsNil(request.MixLayoutParamsShrink) {
		query["MixLayoutParams"] = request.MixLayoutParamsShrink
	}

	if !dara.IsNil(request.MixTranscodeParamsShrink) {
		query["MixTranscodeParams"] = request.MixTranscodeParamsShrink
	}

	if !dara.IsNil(request.NotifyAuthKey) {
		query["NotifyAuthKey"] = request.NotifyAuthKey
	}

	if !dara.IsNil(request.NotifyFileUploadedFormat) {
		query["NotifyFileUploadedFormat"] = request.NotifyFileUploadedFormat
	}

	if !dara.IsNil(request.NotifyUrl) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !dara.IsNil(request.RecordParamsShrink) {
		query["RecordParams"] = request.RecordParamsShrink
	}

	if !dara.IsNil(request.StorageParamsShrink) {
		query["StorageParams"] = request.StorageParamsShrink
	}

	if !dara.IsNil(request.SubscribeParamsShrink) {
		query["SubscribeParams"] = request.SubscribeParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRtcCloudRecording"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRtcCloudRecordingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cloud transcoding job.
//
// Description:
//
// ## Endpoints
//
// The following endpoints are available for this operation.
//
// | Region    | Region ID      | Public endpoint                  |
//
// | --------- | -------------- | -------------------------------- |
//
// | Singapore | ap-southeast-1 | live.ap-southeast-1.aliyuncs.com |
//
// @param tmpReq - StartRtcCloudTranscodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRtcCloudTranscodeResponse
func (client *Client) StartRtcCloudTranscodeWithContext(ctx context.Context, tmpReq *StartRtcCloudTranscodeRequest, runtime *dara.RuntimeOptions) (_result *StartRtcCloudTranscodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartRtcCloudTranscodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InputParam) {
		request.InputParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputParam, dara.String("InputParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OutputParams) {
		request.OutputParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutputParams, dara.String("OutputParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.InputParamShrink) {
		query["InputParam"] = request.InputParamShrink
	}

	if !dara.IsNil(request.MaxIdleTime) {
		query["MaxIdleTime"] = request.MaxIdleTime
	}

	if !dara.IsNil(request.OutputParamsShrink) {
		query["OutputParams"] = request.OutputParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRtcCloudTranscode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRtcCloudTranscodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a production studio. This stops the PVW (preview scene) and PGM (program output scene) scenes.
//
// Description:
//
// You must first call the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation to create a production studio. You can then call this operation to stop the production studio, which stops the Preview (PVW) and Program (PGM) scenes.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second per user. API calls that exceed this limit are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - StopCasterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCasterResponse
func (client *Client) StopCasterWithContext(ctx context.Context, request *StopCasterRequest, runtime *dara.RuntimeOptions) (_result *StopCasterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopCaster"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopCasterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a specified preview (PVW) scene.
//
// Description:
//
// Call this operation to stop a specified preview (PVW) scene.
//
// ## QPS limit
//
// This operation is limited to 10 queries per second (QPS) per user. Exceeding this limit results in API call throttling, which may affect your business. We recommend that you call this operation at a reasonable frequency.
//
// @param request - StopCasterSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCasterSceneResponse
func (client *Client) StopCasterSceneWithContext(ctx context.Context, request *StopCasterSceneRequest, runtime *dara.RuntimeOptions) (_result *StopCasterSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopCasterScene"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopCasterSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an edge transcoding task.
//
// Description:
//
// - This operation stops an edge transcoding job.
//
// - To call this operation, you must have permissions to access the edge transcoding service, and the transcoding job must be in the running state.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 6,000 calls per minute for each account. Calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - StopEdgeTranscodeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopEdgeTranscodeJobResponse
func (client *Client) StopEdgeTranscodeJobWithContext(ctx context.Context, request *StopEdgeTranscodeJobRequest, runtime *dara.RuntimeOptions) (_result *StopEdgeTranscodeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopEdgeTranscodeJob"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopEdgeTranscodeJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a live streaming domain. This operation changes the DomainStatus to offline.
//
// Description:
//
// After a live streaming domain is disabled, its information is retained. The system automatically performs an origin fetch for requests to the domain.
//
// ## QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - StopLiveDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopLiveDomainResponse
func (client *Client) StopLiveDomainWithContext(ctx context.Context, request *StopLiveDomainRequest, runtime *dara.RuntimeOptions) (_result *StopLiveDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopLiveDomain"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopLiveDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a stream mixing and forwarding task.
//
// Description:
//
// - Before you call this operation, you must have already called [StartLiveMPUTask](https://help.aliyun.com/document_detail/2848199.html) to start a stream mixing and forwarding task.
//
// - If you need to stop a stream mixing and forwarding task but the task is abnormal (StopLiveMPUTask was not called to stop the task), the task automatically stops 2 minutes after the last user leaves the channel. To resume stream mixing and forwarding after the task stops, call the [StartLiveMPUTask](https://help.aliyun.com/document_detail/2848199.html) operation again.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 500 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - StopLiveMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopLiveMPUTaskResponse
func (client *Client) StopLiveMPUTaskWithContext(ctx context.Context, request *StopLiveMPUTaskRequest, runtime *dara.RuntimeOptions) (_result *StopLiveMPUTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopLiveMPUTask"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopLiveMPUTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a stream pulling task.
//
// Description:
//
// - Call this operation to stop a stream pulling task.
//
// - You can forcibly stop a running task (including tasks in the abnormal retry state). This operation does not take effect on tasks that are not running.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - StopLivePullToPushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopLivePullToPushResponse
func (client *Client) StopLivePullToPushWithContext(ctx context.Context, request *StopLivePullToPushRequest, runtime *dara.RuntimeOptions) (_result *StopLivePullToPushResponse, _err error) {
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
		Action:      dara.String("StopLivePullToPush"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopLivePullToPushResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a monitoring session.
//
// Description:
//
// You can obtain a **MonitorId*	- by calling the [CreateLiveStreamMonitor](https://help.aliyun.com/document_detail/2848129.html) operation. After a monitoring session starts, call this operation to stop the session.
//
// ## QPS limit
//
// This operation is limited to 10 queries per second (QPS) per user. Calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - StopLiveStreamMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopLiveStreamMonitorResponse
func (client *Client) StopLiveStreamMonitorWithContext(ctx context.Context, request *StopLiveStreamMonitorRequest, runtime *dara.RuntimeOptions) (_result *StopLiveStreamMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MonitorId) {
		query["MonitorId"] = request.MonitorId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopLiveStreamMonitor"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopLiveStreamMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops playing an episode list.
//
// Description:
//
// First, call the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation to add items to a playlist. You can then call this operation to stop the playlist.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation based on your business needs.
//
// @param request - StopPlaylistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopPlaylistResponse
func (client *Client) StopPlaylistWithContext(ctx context.Context, request *StopPlaylistRequest, runtime *dara.RuntimeOptions) (_result *StopPlaylistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProgramId) {
		query["ProgramId"] = request.ProgramId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopPlaylist"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopPlaylistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a live subtitle task.
//
// Description:
//
// This operation is limited to 20 queries per second (QPS) for each account. If you exceed the limit, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - StopRtcAsrTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRtcAsrTaskResponse
func (client *Client) StopRtcAsrTaskWithContext(ctx context.Context, request *StopRtcAsrTaskRequest, runtime *dara.RuntimeOptions) (_result *StopRtcAsrTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopRtcAsrTask"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopRtcAsrTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an RTC cloud recording task.
//
// Description:
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 50 calls per second. If the number of calls exceeds the limit, throttling is triggered, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - StopRtcCloudRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRtcCloudRecordingResponse
func (client *Client) StopRtcCloudRecordingWithContext(ctx context.Context, request *StopRtcCloudRecordingRequest, runtime *dara.RuntimeOptions) (_result *StopRtcCloudRecordingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopRtcCloudRecording"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopRtcCloudRecordingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a cloud transcoding job.
//
// @param request - StopRtcCloudTranscodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRtcCloudTranscodeResponse
func (client *Client) StopRtcCloudTranscodeWithContext(ctx context.Context, request *StopRtcCloudTranscodeRequest, runtime *dara.RuntimeOptions) (_result *StopRtcCloudTranscodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopRtcCloudTranscode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopRtcCloudTranscodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds labels to ApsaraVideo Live resources by calling TagLiveResources.
//
// Description:
//
// The maximum number of times that each user can call this operation per second is 100.
//
// @param request - TagLiveResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagLiveResourcesResponse
func (client *Client) TagLiveResourcesWithContext(ctx context.Context, request *TagLiveResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagLiveResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagLiveResources"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagLiveResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls UnTagLiveResources to delete tags from ApsaraVideo Live resources.
//
// Description:
//
// The maximum call frequency for a single user is 100 calls per second.
//
// @param request - UnTagLiveResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagLiveResourcesResponse
func (client *Client) UnTagLiveResourcesWithContext(ctx context.Context, request *UnTagLiveResourcesRequest, runtime *dara.RuntimeOptions) (_result *UnTagLiveResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnTagLiveResources"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnTagLiveResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unmutes a group.
//
// Description:
//
//	  Before you call this operation, make sure that you have called the [CreateLiveMessageGroup](https://help.aliyun.com/document_detail/2848163.html) operation to create an interactive messaging group.
//
//		- If a user was muted by calling the AddLiveMessageGroupBand operation, the user remains muted even after you call the UnbanLiveMessageGroup operation.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 10 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - UnbanLiveMessageGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbanLiveMessageGroupResponse
func (client *Client) UnbanLiveMessageGroupWithContext(ctx context.Context, request *UnbanLiveMessageGroupRequest, runtime *dara.RuntimeOptions) (_result *UnbanLiveMessageGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataCenter) {
		query["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbanLiveMessageGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbanLiveMessageGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the resource group of a production studio.
//
// Description:
//
// ## Usage notes
//
// When you call this operation, your account must have permissions on both the source and destination resource groups.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second per user. If you exceed the limit, your API calls are throttled. This may affect your business. Make sure that you call this operation within the limit. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param request - UpdateCasterResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCasterResourceGroupResponse
func (client *Client) UpdateCasterResourceGroupWithContext(ctx context.Context, request *UpdateCasterResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateCasterResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCasterResourceGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCasterResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the audio configurations of a scene. This operation also lets you add an audio configuration or apply an existing audio configuration to a new scene.
//
// Description:
//
// First, call the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation to create a production studio. Then, you can call this operation to update the audio configuration of a scene. This operation supports the audio mixing mode and the audio-follows-video (AFV) mode.
//
// ## QPS limits
//
// You can make up to 10 queries per second (QPS) per Alibaba Cloud account. API calls that exceed this limit are throttled, which may affect your business. We recommend that you adhere to this limit.
//
// @param request - UpdateCasterSceneAudioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCasterSceneAudioResponse
func (client *Client) UpdateCasterSceneAudioWithContext(ctx context.Context, request *UpdateCasterSceneAudioRequest, runtime *dara.RuntimeOptions) (_result *UpdateCasterSceneAudioResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AudioLayer) {
		query["AudioLayer"] = request.AudioLayer
	}

	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.FollowEnable) {
		query["FollowEnable"] = request.FollowEnable
	}

	if !dara.IsNil(request.MixList) {
		query["MixList"] = request.MixList
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCasterSceneAudio"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCasterSceneAudioResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a scene, including the layout, without invalidating the previous configurations. This operation is more efficient than the SetCasterSceneConfig operation.
//
// Description:
//
// You can call this operation to incrementally modify the configuration of a scene, including its layout. This operation preserves existing settings and is more efficient than a full update.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second per user. Calls that exceed this limit are throttled, which may affect your business. We recommend that you plan your calls accordingly.
//
// @param request - UpdateCasterSceneConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCasterSceneConfigResponse
func (client *Client) UpdateCasterSceneConfigWithContext(ctx context.Context, request *UpdateCasterSceneConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateCasterSceneConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CasterId) {
		query["CasterId"] = request.CasterId
	}

	if !dara.IsNil(request.ComponentId) {
		query["ComponentId"] = request.ComponentId
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCasterSceneConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCasterSceneConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom transcoding configuration.
//
// Description:
//
// This operation supports the following types of custom transcoding templates:
//
// - h264: H.264 standard transcoding.
//
// - h264-nbhd: H.264 Narrowband HD™ transcoding.
//
// - h265: H.265 standard transcoding.
//
// - h265-nbhd: H.265 Narrowband HD™ transcoding.
//
// - audio: an audio-only transcoding.
//
// ## QPS limit
//
// You can call this operation up to 6,000 times per second per account.
//
// @param request - UpdateCustomLiveStreamTranscodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomLiveStreamTranscodeResponse
func (client *Client) UpdateCustomLiveStreamTranscodeWithContext(ctx context.Context, request *UpdateCustomLiveStreamTranscodeRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomLiveStreamTranscodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.AudioBitrate) {
		query["AudioBitrate"] = request.AudioBitrate
	}

	if !dara.IsNil(request.AudioChannelNum) {
		query["AudioChannelNum"] = request.AudioChannelNum
	}

	if !dara.IsNil(request.AudioCodec) {
		query["AudioCodec"] = request.AudioCodec
	}

	if !dara.IsNil(request.AudioProfile) {
		query["AudioProfile"] = request.AudioProfile
	}

	if !dara.IsNil(request.AudioRate) {
		query["AudioRate"] = request.AudioRate
	}

	if !dara.IsNil(request.BitrateWithSource) {
		query["BitrateWithSource"] = request.BitrateWithSource
	}

	if !dara.IsNil(request.DeInterlaced) {
		query["DeInterlaced"] = request.DeInterlaced
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EncryptParameters) {
		query["EncryptParameters"] = request.EncryptParameters
	}

	if !dara.IsNil(request.ExtWithSource) {
		query["ExtWithSource"] = request.ExtWithSource
	}

	if !dara.IsNil(request.FPS) {
		query["FPS"] = request.FPS
	}

	if !dara.IsNil(request.FpsWithSource) {
		query["FpsWithSource"] = request.FpsWithSource
	}

	if !dara.IsNil(request.Gop) {
		query["Gop"] = request.Gop
	}

	if !dara.IsNil(request.Height) {
		query["Height"] = request.Height
	}

	if !dara.IsNil(request.Lazy) {
		query["Lazy"] = request.Lazy
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Profile) {
		query["Profile"] = request.Profile
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResWithSource) {
		query["ResWithSource"] = request.ResWithSource
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.VideoBitrate) {
		query["VideoBitrate"] = request.VideoBitrate
	}

	if !dara.IsNil(request.Width) {
		query["Width"] = request.Width
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomLiveStreamTranscode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomLiveStreamTranscodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an edge transcoding job.
//
// Description:
//
// - This operation updates an edge transcoding job.
//
// - To call this operation, you must have permissions to access the edge transcoding service. The transcoding job must not have been started.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 6,000 calls per minute per user. If you exceed this limit, API calls are throttled. This may impact your business. Plan your calls accordingly.
//
// @param request - UpdateEdgeTranscodeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEdgeTranscodeJobResponse
func (client *Client) UpdateEdgeTranscodeJobWithContext(ctx context.Context, request *UpdateEdgeTranscodeJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateEdgeTranscodeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamInput) {
		query["StreamInput"] = request.StreamInput
	}

	if !dara.IsNil(request.StreamOutput) {
		query["StreamOutput"] = request.StreamOutput
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEdgeTranscodeJob"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEdgeTranscodeJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a callback that is used to subscribe to channel or user events.
//
// Description:
//
// - Before you call this operation, make sure that you have called the [CreateEventSub](https://help.aliyun.com/document_detail/2848209.html) operation to create a callback that is used to subscribe to channel or user events.
//
// - An existing channel that you specify in this operation still uses its original callback configuration. The updated configuration can apply to the channel only if you restart the channel after it is closed for longer than 20 minutes.
//
// - If you only want to update specific parameters, you must also specify the other required parameters with their original values.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - UpdateEventSubRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEventSubResponse
func (client *Client) UpdateEventSubWithContext(ctx context.Context, request *UpdateEventSubRequest, runtime *dara.RuntimeOptions) (_result *UpdateEventSubResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Events) {
		query["Events"] = request.Events
	}

	if !dara.IsNil(request.SubscribeId) {
		query["SubscribeId"] = request.SubscribeId
	}

	if !dara.IsNil(request.Users) {
		query["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEventSub"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEventSubResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a subtitle rule.
//
// Description:
//
// Updates the parameters of a specified subtitle rule.
//
//	Notice: The real-time subtitle feature is currently in invitational preview. Each user can add up to 300 subtitle templates.
//
// ## QPS limit
//
// The single-user QPS limit for this API is 60 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - UpdateLiveAIProduceRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveAIProduceRulesResponse
func (client *Client) UpdateLiveAIProduceRulesWithContext(ctx context.Context, request *UpdateLiveAIProduceRulesRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveAIProduceRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.IsLazy) {
		query["IsLazy"] = request.IsLazy
	}

	if !dara.IsNil(request.LiveTemplate) {
		query["LiveTemplate"] = request.LiveTemplate
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RulesId) {
		query["RulesId"] = request.RulesId
	}

	if !dara.IsNil(request.StudioName) {
		query["StudioName"] = request.StudioName
	}

	if !dara.IsNil(request.SubtitleId) {
		query["SubtitleId"] = request.SubtitleId
	}

	if !dara.IsNil(request.SubtitleName) {
		query["SubtitleName"] = request.SubtitleName
	}

	if !dara.IsNil(request.Suffix) {
		query["Suffix"] = request.Suffix
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveAIProduceRules"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveAIProduceRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a subtitle template by calling UpdateLiveAISubtitle.
//
// Description:
//
// ## Operation description
//
// This operation updates the parameters of a specified subtitle template.
//
//	Notice: The real-time subtitle feature is currently in invitational preview. Each user can add up to 300 subtitle templates.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 60 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/343507.html).
//
// @param tmpReq - UpdateLiveAISubtitleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveAISubtitleResponse
func (client *Client) UpdateLiveAISubtitleWithContext(ctx context.Context, tmpReq *UpdateLiveAISubtitleRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveAISubtitleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateLiveAISubtitleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PositionNormalized) {
		request.PositionNormalizedShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PositionNormalized, dara.String("PositionNormalized"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BgColor) {
		query["BgColor"] = request.BgColor
	}

	if !dara.IsNil(request.BgWidthNormalized) {
		query["BgWidthNormalized"] = request.BgWidthNormalized
	}

	if !dara.IsNil(request.BorderWidthNormalized) {
		query["BorderWidthNormalized"] = request.BorderWidthNormalized
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DstLanguage) {
		query["DstLanguage"] = request.DstLanguage
	}

	if !dara.IsNil(request.FontColor) {
		query["FontColor"] = request.FontColor
	}

	if !dara.IsNil(request.FontName) {
		query["FontName"] = request.FontName
	}

	if !dara.IsNil(request.FontSizeNormalized) {
		query["FontSizeNormalized"] = request.FontSizeNormalized
	}

	if !dara.IsNil(request.Height) {
		query["Height"] = request.Height
	}

	if !dara.IsNil(request.MaxLines) {
		query["MaxLines"] = request.MaxLines
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PositionNormalizedShrink) {
		query["PositionNormalized"] = request.PositionNormalizedShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ShowSourceLan) {
		query["ShowSourceLan"] = request.ShowSourceLan
	}

	if !dara.IsNil(request.SrcLanguage) {
		query["SrcLanguage"] = request.SrcLanguage
	}

	if !dara.IsNil(request.SubtitleId) {
		query["SubtitleId"] = request.SubtitleId
	}

	if !dara.IsNil(request.SubtitleName) {
		query["SubtitleName"] = request.SubtitleName
	}

	if !dara.IsNil(request.Width) {
		query["Width"] = request.Width
	}

	if !dara.IsNil(request.WordPerLine) {
		query["WordPerLine"] = request.WordPerLine
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveAISubtitle"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveAISubtitleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a live stream recording configuration that saves the output to Object Storage Service (OSS).
//
// Description:
//
// ## QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - UpdateLiveAppRecordConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveAppRecordConfigResponse
func (client *Client) UpdateLiveAppRecordConfigWithContext(ctx context.Context, request *UpdateLiveAppRecordConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveAppRecordConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DelayTime) {
		query["DelayTime"] = request.DelayTime
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OnDemand) {
		query["OnDemand"] = request.OnDemand
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RecordFormat) {
		query["RecordFormat"] = request.RecordFormat
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.TranscodeRecordFormat) {
		query["TranscodeRecordFormat"] = request.TranscodeRecordFormat
	}

	if !dara.IsNil(request.TranscodeTemplates) {
		query["TranscodeTemplates"] = request.TranscodeTemplates
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveAppRecordConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveAppRecordConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a snapshot configuration of a streaming domain. The captured snapshots are stored in Object Storage Service (OSS). The modification takes effect after you restart stream ingest.
//
// Description:
//
// You can call this operation to modify a snapshot configuration of a streaming domain. The captured snapshots are stored in OSS. The modification takes effect after you restart stream ingest.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - UpdateLiveAppSnapshotConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveAppSnapshotConfigResponse
func (client *Client) UpdateLiveAppSnapshotConfigWithContext(ctx context.Context, request *UpdateLiveAppSnapshotConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveAppSnapshotConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Callback) {
		query["Callback"] = request.Callback
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OverwriteOssObject) {
		query["OverwriteOssObject"] = request.OverwriteOssObject
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SequenceOssObject) {
		query["SequenceOssObject"] = request.SequenceOssObject
	}

	if !dara.IsNil(request.TimeInterval) {
		query["TimeInterval"] = request.TimeInterval
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveAppSnapshotConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveAppSnapshotConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the voice content moderation configuration for live streaming.
//
// Description:
//
// - Obtain the streamer streaming domain before invoking this operation to update the voice content moderation configuration for live streaming.
//
// - Before invoking this operation, create a voice content moderation configuration by calling AddLiveAudioAuditConfig.
//
// - Currently, only some live centers support intelligent content moderation for live streaming. For information about the live centers that support this feature, see [Service regions](https://help.aliyun.com/document_detail/193730.html).
//
// ## QPS limit
//
// The maximum number of queries per second (QPS) per user for this operation is 10. If the number of calls per second exceeds the limit, throttling is triggered. This may affect your business. Invoke this operation as needed.
//
// @param request - UpdateLiveAudioAuditConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveAudioAuditConfigResponse
func (client *Client) UpdateLiveAudioAuditConfigWithContext(ctx context.Context, request *UpdateLiveAudioAuditConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveAudioAuditConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssObject) {
		query["OssObject"] = request.OssObject
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveAudioAuditConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveAudioAuditConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the callback configuration for automated audio review.
//
// Description:
//
// - Obtain the streaming domain and then call this operation to update the callback configuration for automated audio review.
//
// - Live audio moderation is available only in specific regions. For supported regions, see [Service regions](https://help.aliyun.com/document_detail/193730.html).
//
// ## QPS limits
//
// You can call this operation up to 10 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - UpdateLiveAudioAuditNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveAudioAuditNotifyConfigResponse
func (client *Client) UpdateLiveAudioAuditNotifyConfigWithContext(ctx context.Context, request *UpdateLiveAudioAuditNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveAudioAuditNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Callback) {
		query["Callback"] = request.Callback
	}

	if !dara.IsNil(request.CallbackTemplate) {
		query["CallbackTemplate"] = request.CallbackTemplate
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveAudioAuditNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveAudioAuditNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the live center transfer configuration by calling UpdateLiveCenterTransfer.
//
// Description:
//
// This operation only supports updating configurations related to the **transfer validity period*	- (TransferArgs, StartTime, EndTime).
//
// ## QPS limit
//
// The QPS limit for this operation is 100 calls per second per user. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable rate.
//
// @param request - UpdateLiveCenterTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveCenterTransferResponse
func (client *Client) UpdateLiveCenterTransferWithContext(ctx context.Context, request *UpdateLiveCenterTransferRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveCenterTransferResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DstUrl) {
		query["DstUrl"] = request.DstUrl
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.TransferArgs) {
		query["TransferArgs"] = request.TransferArgs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveCenterTransfer"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveCenterTransferResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the live stream delay configuration.
//
// Description:
//
// You can call this operation up to 60 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - UpdateLiveDelayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveDelayConfigResponse
func (client *Client) UpdateLiveDelayConfigWithContext(ctx context.Context, request *UpdateLiveDelayConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveDelayConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.DelayTime) {
		query["DelayTime"] = request.DelayTime
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	if !dara.IsNil(request.TaskTriggerMode) {
		query["TaskTriggerMode"] = request.TaskTriggerMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveDelayConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveDelayConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of callbacks for video moderation results. As a result, the callback URL that is used to receive the callback notifications is changed.
//
// Description:
//
// - Obtain the main streaming domain, and then call this operation to modify the configuration of callbacks for video moderation results.
//
// - Only some live centers support the content moderation feature. For more information, see [Supported regions](https://help.aliyun.com/document_detail/193730.html).
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - UpdateLiveDetectNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveDetectNotifyConfigResponse
func (client *Client) UpdateLiveDetectNotifyConfigWithContext(ctx context.Context, request *UpdateLiveDetectNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveDetectNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.NotifyUrl) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveDetectNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveDetectNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a mixed-stream relay task.
//
// Description:
//
// Make sure that a mixed-stream relay task is created before you call this operation. You can call the [StartLiveMPUTask](https://help.aliyun.com/document_detail/2848199.html) operation to create a mixed-stream relay task.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 500 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param tmpReq - UpdateLiveMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveMPUTaskResponse
func (client *Client) UpdateLiveMPUTaskWithContext(ctx context.Context, tmpReq *UpdateLiveMPUTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveMPUTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateLiveMPUTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MultiStreamURL) {
		request.MultiStreamURLShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MultiStreamURL, dara.String("MultiStreamURL"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SeiParams) {
		request.SeiParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SeiParams, dara.String("SeiParams"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SingleSubParams) {
		request.SingleSubParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SingleSubParams, dara.String("SingleSubParams"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TranscodeParams) {
		request.TranscodeParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TranscodeParams, dara.String("TranscodeParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.MixMode) {
		query["MixMode"] = request.MixMode
	}

	if !dara.IsNil(request.MultiStreamURLShrink) {
		query["MultiStreamURL"] = request.MultiStreamURLShrink
	}

	if !dara.IsNil(request.SeiParamsShrink) {
		query["SeiParams"] = request.SeiParamsShrink
	}

	if !dara.IsNil(request.SingleSubParamsShrink) {
		query["SingleSubParams"] = request.SingleSubParamsShrink
	}

	if !dara.IsNil(request.StreamURL) {
		query["StreamURL"] = request.StreamURL
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TranscodeParamsShrink) {
		query["TranscodeParams"] = request.TranscodeParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveMPUTask"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveMPUTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates live encapsulation configurations.
//
// Description:
//
// You can call this operation to update live encapsulation configurations. The new configurations take effect after you restart the stream ingest.
//
// ## QPS limit
//
// This operation is limited to 300 queries per second (QPS) per user. If you exceed this limit, API calls are throttled. This can affect your business operations. Plan your calls accordingly.
//
// @param request - UpdateLivePackageConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLivePackageConfigResponse
func (client *Client) UpdateLivePackageConfigWithContext(ctx context.Context, request *UpdateLivePackageConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLivePackageConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.IgnoreTranscode) {
		query["IgnoreTranscode"] = request.IgnoreTranscode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PartDuration) {
		query["PartDuration"] = request.PartDuration
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SegmentDuration) {
		query["SegmentDuration"] = request.SegmentDuration
	}

	if !dara.IsNil(request.SegmentNum) {
		query["SegmentNum"] = request.SegmentNum
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLivePackageConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLivePackageConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of live stream pulling, including the addresses of the origin server, start time, and end time.
//
// Description:
//
// This operation is applicable to fixed stream pulling. Invoke this operation to update the configuration of live stream pulling, including the addresses of the origin server, start time, and end time.
//
// Before invoking this operation, use AddLivePullStreamInfoConfig to create a stream pulling configuration for the specified (DomainName, AppName, StreamName).
//
// > After you invoke this operation to modify the configuration, the live streaming service re-executes fixed stream pulling based on the modified configuration. Make sure that the modification does not affect your online services.
//
// ## QPS limit
//
// The QPS limit for a single user on this operation is 1000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Invoke this operation appropriately.
//
// @param request - UpdateLivePullStreamInfoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLivePullStreamInfoConfigResponse
func (client *Client) UpdateLivePullStreamInfoConfigWithContext(ctx context.Context, request *UpdateLivePullStreamInfoConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLivePullStreamInfoConfigResponse, _err error) {
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
		Action:      dara.String("UpdateLivePullStreamInfoConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLivePullStreamInfoConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a stream pulling task by calling UpdateLivePullToPush.
//
// Description:
//
// - Call this operation to update a stream pulling task.
//
// - If the task has not reached the configured start time, you can update the Region parameter.
//
// - If the task is running (including in an abnormal retry state), only CallbackUrl and RepeatTime can be updated, and the updates take effect immediately.
//
// - If the task is stopped, all parameters except Region can be updated.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param tmpReq - UpdateLivePullToPushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLivePullToPushResponse
func (client *Client) UpdateLivePullToPushWithContext(ctx context.Context, tmpReq *UpdateLivePullToPushRequest, runtime *dara.RuntimeOptions) (_result *UpdateLivePullToPushResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateLivePullToPushShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SourceUrls) {
		request.SourceUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceUrls, dara.String("SourceUrls"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLivePullToPush"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLivePullToPushResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the domain-level recording callback configuration.
//
// Description:
//
// When you update the domain-level recording callback configuration, you can modify the following settings:
//
// - The callback URL for recording events, including recording file generation event callbacks and recording task status callbacks. For more information, see [Recording event callbacks](https://help.aliyun.com/document_detail/55016.html).
//
// - The on-demand recording callback URL. For more information, see [On-demand recording callback](https://help.aliyun.com/document_detail/85910.html).
//
// - Whether recording task status callbacks are required.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 30 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately.
//
// @param request - UpdateLiveRecordNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveRecordNotifyConfigResponse
func (client *Client) UpdateLiveRecordNotifyConfigWithContext(ctx context.Context, request *UpdateLiveRecordNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveRecordNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.NeedStatusNotify) {
		query["NeedStatusNotify"] = request.NeedStatusNotify
	}

	if !dara.IsNil(request.NotifyAuthKey) {
		query["NotifyAuthKey"] = request.NotifyAuthKey
	}

	if !dara.IsNil(request.NotifyReqAuth) {
		query["NotifyReqAuth"] = request.NotifyReqAuth
	}

	if !dara.IsNil(request.NotifyUrl) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !dara.IsNil(request.OnDemandUrl) {
		query["OnDemandUrl"] = request.OnDemandUrl
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveRecordNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveRecordNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update a live-to-VOD configuration.
//
// Description:
//
// This operation has a rate limit of 1,000 calls per minute per account. If you exceed this limit, your API calls will be rate-limited, which may interrupt your service.
//
// @param request - UpdateLiveRecordVodConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveRecordVodConfigResponse
func (client *Client) UpdateLiveRecordVodConfigWithContext(ctx context.Context, request *UpdateLiveRecordVodConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveRecordVodConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AutoCompose) {
		query["AutoCompose"] = request.AutoCompose
	}

	if !dara.IsNil(request.ComposeVodTranscodeGroupId) {
		query["ComposeVodTranscodeGroupId"] = request.ComposeVodTranscodeGroupId
	}

	if !dara.IsNil(request.CycleDuration) {
		query["CycleDuration"] = request.CycleDuration
	}

	if !dara.IsNil(request.DelayTime) {
		query["DelayTime"] = request.DelayTime
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OnDemand) {
		query["OnDemand"] = request.OnDemand
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RecordFormat) {
		query["RecordFormat"] = request.RecordFormat
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.TranscodeTemplates) {
		query["TranscodeTemplates"] = request.TranscodeTemplates
	}

	if !dara.IsNil(request.VodTranscodeGroupId) {
		query["VodTranscodeGroupId"] = request.VodTranscodeGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveRecordVodConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveRecordVodConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the content moderation configuration for live streaming.
//
// Description:
//
// - Obtain the streaming domain of the streamer first, and then invoke this operation to update the content moderation configuration for live streaming.
//
// - Currently, only some live centers support intelligent content moderation for live streaming. For more information about the live centers that support this feature, see [Service regions](https://help.aliyun.com/document_detail/193730.html).
//
// - Before you invoke this operation, you must have already created a content moderation configuration for the specified DomainName and AppName by invoking AddLiveSnapshotDetectPornConfig.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 30 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Invoke this operation appropriately.
//
// @param request - UpdateLiveSnapshotDetectPornConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveSnapshotDetectPornConfigResponse
func (client *Client) UpdateLiveSnapshotDetectPornConfigWithContext(ctx context.Context, request *UpdateLiveSnapshotDetectPornConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveSnapshotDetectPornConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssObject) {
		query["OssObject"] = request.OssObject
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveSnapshotDetectPornConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveSnapshotDetectPornConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the snapshot callback configuration.
//
// Description:
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - UpdateLiveSnapshotNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveSnapshotNotifyConfigResponse
func (client *Client) UpdateLiveSnapshotNotifyConfigWithContext(ctx context.Context, request *UpdateLiveSnapshotNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveSnapshotNotifyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.NotifyAuthKey) {
		query["NotifyAuthKey"] = request.NotifyAuthKey
	}

	if !dara.IsNil(request.NotifyReqAuth) {
		query["NotifyReqAuth"] = request.NotifyReqAuth
	}

	if !dara.IsNil(request.NotifyUrl) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveSnapshotNotifyConfig"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveSnapshotNotifyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of a monitoring session.
//
// Description:
//
// This operation updates the configuration of a monitoring session. If the monitoring session is running, updates to the input source configuration take effect in real time.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10. API calls that exceed this limit are throttled. Throttling can affect your business, so we recommend that you plan your calls accordingly.
//
// @param request - UpdateLiveStreamMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveStreamMonitorResponse
func (client *Client) UpdateLiveStreamMonitorWithContext(ctx context.Context, request *UpdateLiveStreamMonitorRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveStreamMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.DingTalkWebHookUrl) {
		query["DingTalkWebHookUrl"] = request.DingTalkWebHookUrl
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.InputList) {
		query["InputList"] = request.InputList
	}

	if !dara.IsNil(request.MonitorConfig) {
		query["MonitorConfig"] = request.MonitorConfig
	}

	if !dara.IsNil(request.MonitorId) {
		query["MonitorId"] = request.MonitorId
	}

	if !dara.IsNil(request.MonitorName) {
		query["MonitorName"] = request.MonitorName
	}

	if !dara.IsNil(request.OutputTemplate) {
		query["OutputTemplate"] = request.OutputTemplate
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveStreamMonitor"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveStreamMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a default transcoding configuration.
//
// Description:
//
// To modify the encryption settings, first obtain a Customer Master Key (CMK) ID from Key Management Service (KMS). This operation supports only standard and Narrowband HD™ transcoding templates.
//
// ## QPS limits
//
// ou can call this operation up to 30 times per second per account.
//
// @param request - UpdateLiveStreamTranscodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveStreamTranscodeResponse
func (client *Client) UpdateLiveStreamTranscodeWithContext(ctx context.Context, request *UpdateLiveStreamTranscodeRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveStreamTranscodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EncryptParameters) {
		query["EncryptParameters"] = request.EncryptParameters
	}

	if !dara.IsNil(request.Lazy) {
		query["Lazy"] = request.Lazy
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveStreamTranscode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveStreamTranscodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a watermark template.
//
// Description:
//
// This operation updates the parameters of a specified watermark template.
//
// ## QPS limits
//
// You can call this operation up to 60 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - UpdateLiveStreamWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveStreamWatermarkResponse
func (client *Client) UpdateLiveStreamWatermarkWithContext(ctx context.Context, request *UpdateLiveStreamWatermarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveStreamWatermarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Height) {
		query["Height"] = request.Height
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OffsetCorner) {
		query["OffsetCorner"] = request.OffsetCorner
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PictureUrl) {
		query["PictureUrl"] = request.PictureUrl
	}

	if !dara.IsNil(request.RefHeight) {
		query["RefHeight"] = request.RefHeight
	}

	if !dara.IsNil(request.RefWidth) {
		query["RefWidth"] = request.RefWidth
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Transparency) {
		query["Transparency"] = request.Transparency
	}

	if !dara.IsNil(request.XOffset) {
		query["XOffset"] = request.XOffset
	}

	if !dara.IsNil(request.YOffset) {
		query["YOffset"] = request.YOffset
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveStreamWatermark"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveStreamWatermarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a watermark rule.
//
// Description:
//
// This operation updates the parameters of a watermark rule based on the specified rule ID.
//
// ## QPS limit
//
// You can call this operation up to 60 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - UpdateLiveStreamWatermarkRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveStreamWatermarkRuleResponse
func (client *Client) UpdateLiveStreamWatermarkRuleWithContext(ctx context.Context, request *UpdateLiveStreamWatermarkRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveStreamWatermarkRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveStreamWatermarkRule"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveStreamWatermarkRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call UpdateMessageApp to update interactive message application information.
//
// Description:
//
// ## QPS limits
//
// The single-user QPS limit for this API is 100 queries per second (QPS). API calls that exceed this limit are throttled, which may affect your business. You can call this API at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param tmpReq - UpdateMessageAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMessageAppResponse
func (client *Client) UpdateMessageAppWithContext(ctx context.Context, tmpReq *UpdateMessageAppRequest, runtime *dara.RuntimeOptions) (_result *UpdateMessageAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMessageAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AppConfig) {
		request.AppConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppConfig, dara.String("AppConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Extension) {
		request.ExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extension, dara.String("Extension"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppConfigShrink) {
		body["AppConfig"] = request.AppConfigShrink
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ExtensionShrink) {
		body["Extension"] = request.ExtensionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMessageApp"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMessageAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the UpdateMessageGroup operation to update message group information.
//
// Description:
//
// ## QPS limits
//
// The QPS limit for this API is 100 queries per second (QPS) per user. If the limit is exceeded, API calls will be throttled, which may affect your business. You can call the API properly to avoid this issue. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
//
// @param tmpReq - UpdateMessageGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMessageGroupResponse
func (client *Client) UpdateMessageGroupWithContext(ctx context.Context, tmpReq *UpdateMessageGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateMessageGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMessageGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Extension) {
		request.ExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extension, dara.String("Extension"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ExtensionShrink) {
		body["Extension"] = request.ExtensionShrink
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMessageGroup"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMessageGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a stream mix task.
//
// Description:
//
// You can call this API operation to update a stream mix task. You can update the input sources and layout, but you cannot update parameters such as the output resolution.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, your API calls are throttled. This may impact your business. Plan your calls accordingly.
//
// @param request - UpdateMixStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMixStreamResponse
func (client *Client) UpdateMixStreamWithContext(ctx context.Context, request *UpdateMixStreamRequest, runtime *dara.RuntimeOptions) (_result *UpdateMixStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.InputStreamList) {
		query["InputStreamList"] = request.InputStreamList
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.MixStreamId) {
		query["MixStreamId"] = request.MixStreamId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMixStream"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMixStreamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an RTC cloud recording task.
//
// Description:
//
// Single-stream recording supports updating subscription parameters. Stream mixing recording supports updating only the subscribed user streams.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation as appropriate.
//
// @param tmpReq - UpdateRtcCloudRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRtcCloudRecordingResponse
func (client *Client) UpdateRtcCloudRecordingWithContext(ctx context.Context, tmpReq *UpdateRtcCloudRecordingRequest, runtime *dara.RuntimeOptions) (_result *UpdateRtcCloudRecordingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateRtcCloudRecordingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MixLayoutParams) {
		request.MixLayoutParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MixLayoutParams, dara.String("MixLayoutParams"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SubscribeParams) {
		request.SubscribeParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubscribeParams, dara.String("SubscribeParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MixLayoutParamsShrink) {
		query["MixLayoutParams"] = request.MixLayoutParamsShrink
	}

	if !dara.IsNil(request.SubscribeParamsShrink) {
		query["SubscribeParams"] = request.SubscribeParamsShrink
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRtcCloudRecording"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRtcCloudRecordingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a subscription to mixed-stream relay events.
//
// Description:
//
// - You can call this operation to update a subscription to mixed-stream relay events. You can modify parameters such as the callback URL and channel IDs.
//
// - Before you call this operation, make sure that you have called the CreateRtcMPUEventSub operation to create the subscription.
//
// ## [](#qps-)QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - UpdateRtcMPUEventSubRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRtcMPUEventSubResponse
func (client *Client) UpdateRtcMPUEventSubWithContext(ctx context.Context, request *UpdateRtcMPUEventSubRequest, runtime *dara.RuntimeOptions) (_result *UpdateRtcMPUEventSubResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.ChannelIds) {
		query["ChannelIds"] = request.ChannelIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRtcMPUEventSub"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRtcMPUEventSubResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a custom Real-time Streaming (RTS) transcoding template.
//
// Description:
//
// This API operation supports only the following types of custom transcoding templates: h264, h264-nbhd, h264-origin, and audio.
//
// ## QPS limit
//
// You can call this operation up to 10 times per second per account.
//
// @param request - UpdateRtsLiveStreamTranscodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRtsLiveStreamTranscodeResponse
func (client *Client) UpdateRtsLiveStreamTranscodeWithContext(ctx context.Context, request *UpdateRtsLiveStreamTranscodeRequest, runtime *dara.RuntimeOptions) (_result *UpdateRtsLiveStreamTranscodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.AudioBitrate) {
		query["AudioBitrate"] = request.AudioBitrate
	}

	if !dara.IsNil(request.AudioChannelNum) {
		query["AudioChannelNum"] = request.AudioChannelNum
	}

	if !dara.IsNil(request.AudioCodec) {
		query["AudioCodec"] = request.AudioCodec
	}

	if !dara.IsNil(request.AudioProfile) {
		query["AudioProfile"] = request.AudioProfile
	}

	if !dara.IsNil(request.AudioRate) {
		query["AudioRate"] = request.AudioRate
	}

	if !dara.IsNil(request.DeleteBframes) {
		query["DeleteBframes"] = request.DeleteBframes
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.FPS) {
		query["FPS"] = request.FPS
	}

	if !dara.IsNil(request.Gop) {
		query["Gop"] = request.Gop
	}

	if !dara.IsNil(request.Height) {
		query["Height"] = request.Height
	}

	if !dara.IsNil(request.Lazy) {
		query["Lazy"] = request.Lazy
	}

	if !dara.IsNil(request.Opus) {
		query["Opus"] = request.Opus
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Profile) {
		query["Profile"] = request.Profile
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.VideoBitrate) {
		query["VideoBitrate"] = request.VideoBitrate
	}

	if !dara.IsNil(request.Width) {
		query["Width"] = request.Width
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRtsLiveStreamTranscode"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRtsLiveStreamTranscodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the ownership of a domain name.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions.
//
// @param request - VerifyLiveDomainOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyLiveDomainOwnerResponse
func (client *Client) VerifyLiveDomainOwnerWithContext(ctx context.Context, request *VerifyLiveDomainOwnerRequest, runtime *dara.RuntimeOptions) (_result *VerifyLiveDomainOwnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VerifyType) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyLiveDomainOwner"),
		Version:     dara.String("2016-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyLiveDomainOwnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
