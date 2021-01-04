// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RetrieveFaceRequest struct {
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Face64String *string `json:"Face64String,omitempty" xml:"Face64String,omitempty"`
	FaceUrl      *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
}

func (s RetrieveFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s RetrieveFaceRequest) GoString() string {
	return s.String()
}

func (s *RetrieveFaceRequest) SetProjectId(v string) *RetrieveFaceRequest {
	s.ProjectId = &v
	return s
}

func (s *RetrieveFaceRequest) SetFace64String(v string) *RetrieveFaceRequest {
	s.Face64String = &v
	return s
}

func (s *RetrieveFaceRequest) SetFaceUrl(v string) *RetrieveFaceRequest {
	s.FaceUrl = &v
	return s
}

type RetrieveFaceResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RetrieveFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RetrieveFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetrieveFaceResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveFaceResponseBody) SetRequestId(v string) *RetrieveFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetrieveFaceResponseBody) SetData(v *RetrieveFaceResponseBodyData) *RetrieveFaceResponseBody {
	s.Data = v
	return s
}

func (s *RetrieveFaceResponseBody) SetSuccess(v bool) *RetrieveFaceResponseBody {
	s.Success = &v
	return s
}

type RetrieveFaceResponseBodyData struct {
	Data []*RetrieveFaceResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s RetrieveFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RetrieveFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetrieveFaceResponseBodyData) SetData(v []*RetrieveFaceResponseBodyDataData) *RetrieveFaceResponseBodyData {
	s.Data = v
	return s
}

type RetrieveFaceResponseBodyDataData struct {
	UserId   *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Rate     *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s RetrieveFaceResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s RetrieveFaceResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *RetrieveFaceResponseBodyDataData) SetUserId(v int64) *RetrieveFaceResponseBodyDataData {
	s.UserId = &v
	return s
}

func (s *RetrieveFaceResponseBodyDataData) SetUserName(v string) *RetrieveFaceResponseBodyDataData {
	s.UserName = &v
	return s
}

func (s *RetrieveFaceResponseBodyDataData) SetRate(v string) *RetrieveFaceResponseBodyDataData {
	s.Rate = &v
	return s
}

type RetrieveFaceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RetrieveFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RetrieveFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s RetrieveFaceResponse) GoString() string {
	return s.String()
}

func (s *RetrieveFaceResponse) SetHeaders(v map[string]*string) *RetrieveFaceResponse {
	s.Headers = v
	return s
}

func (s *RetrieveFaceResponse) SetBody(v *RetrieveFaceResponseBody) *RetrieveFaceResponse {
	s.Body = v
	return s
}

type UploadIdentifyRecordRequest struct {
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName               *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	ProjectId              *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	IotId                  *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IdentifyingImageBase64 *string `json:"IdentifyingImageBase64,omitempty" xml:"IdentifyingImageBase64,omitempty"`
	IdentifyingTime        *int64  `json:"IdentifyingTime,omitempty" xml:"IdentifyingTime,omitempty"`
	IdentifyingImageUrl    *string `json:"IdentifyingImageUrl,omitempty" xml:"IdentifyingImageUrl,omitempty"`
	DeviceName             *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	ProductKey             *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	DeviceSecret           *string `json:"DeviceSecret,omitempty" xml:"DeviceSecret,omitempty"`
	Ext                    *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
}

func (s UploadIdentifyRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadIdentifyRecordRequest) GoString() string {
	return s.String()
}

func (s *UploadIdentifyRecordRequest) SetUserId(v string) *UploadIdentifyRecordRequest {
	s.UserId = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetUserName(v string) *UploadIdentifyRecordRequest {
	s.UserName = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetProjectId(v string) *UploadIdentifyRecordRequest {
	s.ProjectId = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetIotId(v string) *UploadIdentifyRecordRequest {
	s.IotId = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetIdentifyingImageBase64(v string) *UploadIdentifyRecordRequest {
	s.IdentifyingImageBase64 = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetIdentifyingTime(v int64) *UploadIdentifyRecordRequest {
	s.IdentifyingTime = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetIdentifyingImageUrl(v string) *UploadIdentifyRecordRequest {
	s.IdentifyingImageUrl = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetDeviceName(v string) *UploadIdentifyRecordRequest {
	s.DeviceName = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetProductKey(v string) *UploadIdentifyRecordRequest {
	s.ProductKey = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetDeviceSecret(v string) *UploadIdentifyRecordRequest {
	s.DeviceSecret = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetExt(v string) *UploadIdentifyRecordRequest {
	s.Ext = &v
	return s
}

type UploadIdentifyRecordResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadIdentifyRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadIdentifyRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UploadIdentifyRecordResponseBody) SetRequestId(v string) *UploadIdentifyRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadIdentifyRecordResponseBody) SetHttpStatusCode(v int32) *UploadIdentifyRecordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UploadIdentifyRecordResponseBody) SetSuccess(v bool) *UploadIdentifyRecordResponseBody {
	s.Success = &v
	return s
}

type UploadIdentifyRecordResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadIdentifyRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadIdentifyRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadIdentifyRecordResponse) GoString() string {
	return s.String()
}

func (s *UploadIdentifyRecordResponse) SetHeaders(v map[string]*string) *UploadIdentifyRecordResponse {
	s.Headers = v
	return s
}

func (s *UploadIdentifyRecordResponse) SetBody(v *UploadIdentifyRecordResponseBody) *UploadIdentifyRecordResponse {
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
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudauth-console"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) RetrieveFaceWithOptions(request *RetrieveFaceRequest, runtime *util.RuntimeOptions) (_result *RetrieveFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RetrieveFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RetrieveFace"), tea.String("2019-04-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetrieveFace(request *RetrieveFaceRequest) (_result *RetrieveFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetrieveFaceResponse{}
	_body, _err := client.RetrieveFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadIdentifyRecordWithOptions(request *UploadIdentifyRecordRequest, runtime *util.RuntimeOptions) (_result *UploadIdentifyRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadIdentifyRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadIdentifyRecord"), tea.String("2019-04-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadIdentifyRecord(request *UploadIdentifyRecordRequest) (_result *UploadIdentifyRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadIdentifyRecordResponse{}
	_body, _err := client.UploadIdentifyRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
