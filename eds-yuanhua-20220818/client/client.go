// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetDumpsResponseBody struct {
	Content    interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	ErrorCode  *int32      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorLevel *string     `json:"ErrorLevel,omitempty" xml:"ErrorLevel,omitempty"`
	ErrorMsg   *string     `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Success    *string     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDumpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDumpsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDumpsResponseBody) SetContent(v interface{}) *GetDumpsResponseBody {
	s.Content = v
	return s
}

func (s *GetDumpsResponseBody) SetErrorCode(v int32) *GetDumpsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDumpsResponseBody) SetErrorLevel(v string) *GetDumpsResponseBody {
	s.ErrorLevel = &v
	return s
}

func (s *GetDumpsResponseBody) SetErrorMsg(v string) *GetDumpsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetDumpsResponseBody) SetSuccess(v string) *GetDumpsResponseBody {
	s.Success = &v
	return s
}

type GetDumpsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDumpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDumpsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDumpsResponse) GoString() string {
	return s.String()
}

func (s *GetDumpsResponse) SetHeaders(v map[string]*string) *GetDumpsResponse {
	s.Headers = v
	return s
}

func (s *GetDumpsResponse) SetStatusCode(v int32) *GetDumpsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDumpsResponse) SetBody(v *GetDumpsResponseBody) *GetDumpsResponse {
	s.Body = v
	return s
}

type ShowTrackpointDataServiceRequest struct {
	DesktopID  *string `json:"DesktopID,omitempty" xml:"DesktopID,omitempty"`
	HostName   *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	TimeBegin  *int64  `json:"TimeBegin,omitempty" xml:"TimeBegin,omitempty"`
	TimeEnd    *int64  `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
	DataType   *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	Module     *string `json:"module,omitempty" xml:"module,omitempty"`
}

func (s ShowTrackpointDataServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ShowTrackpointDataServiceRequest) GoString() string {
	return s.String()
}

func (s *ShowTrackpointDataServiceRequest) SetDesktopID(v string) *ShowTrackpointDataServiceRequest {
	s.DesktopID = &v
	return s
}

func (s *ShowTrackpointDataServiceRequest) SetHostName(v string) *ShowTrackpointDataServiceRequest {
	s.HostName = &v
	return s
}

func (s *ShowTrackpointDataServiceRequest) SetInstanceID(v string) *ShowTrackpointDataServiceRequest {
	s.InstanceID = &v
	return s
}

func (s *ShowTrackpointDataServiceRequest) SetTimeBegin(v int64) *ShowTrackpointDataServiceRequest {
	s.TimeBegin = &v
	return s
}

func (s *ShowTrackpointDataServiceRequest) SetTimeEnd(v int64) *ShowTrackpointDataServiceRequest {
	s.TimeEnd = &v
	return s
}

func (s *ShowTrackpointDataServiceRequest) SetDataType(v string) *ShowTrackpointDataServiceRequest {
	s.DataType = &v
	return s
}

func (s *ShowTrackpointDataServiceRequest) SetModule(v string) *ShowTrackpointDataServiceRequest {
	s.Module = &v
	return s
}

type ShowTrackpointDataServiceResponseBody struct {
	ErrorCode  *int64      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorLevel *string     `json:"ErrorLevel,omitempty" xml:"ErrorLevel,omitempty"`
	ErrorMsg   *string     `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId  *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Content    interface{} `json:"content,omitempty" xml:"content,omitempty"`
}

func (s ShowTrackpointDataServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ShowTrackpointDataServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ShowTrackpointDataServiceResponseBody) SetErrorCode(v int64) *ShowTrackpointDataServiceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ShowTrackpointDataServiceResponseBody) SetErrorLevel(v string) *ShowTrackpointDataServiceResponseBody {
	s.ErrorLevel = &v
	return s
}

func (s *ShowTrackpointDataServiceResponseBody) SetErrorMsg(v string) *ShowTrackpointDataServiceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ShowTrackpointDataServiceResponseBody) SetRequestId(v string) *ShowTrackpointDataServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShowTrackpointDataServiceResponseBody) SetSuccess(v bool) *ShowTrackpointDataServiceResponseBody {
	s.Success = &v
	return s
}

func (s *ShowTrackpointDataServiceResponseBody) SetContent(v interface{}) *ShowTrackpointDataServiceResponseBody {
	s.Content = v
	return s
}

type ShowTrackpointDataServiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ShowTrackpointDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ShowTrackpointDataServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ShowTrackpointDataServiceResponse) GoString() string {
	return s.String()
}

func (s *ShowTrackpointDataServiceResponse) SetHeaders(v map[string]*string) *ShowTrackpointDataServiceResponse {
	s.Headers = v
	return s
}

func (s *ShowTrackpointDataServiceResponse) SetStatusCode(v int32) *ShowTrackpointDataServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ShowTrackpointDataServiceResponse) SetBody(v *ShowTrackpointDataServiceResponseBody) *ShowTrackpointDataServiceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("eds-yuanhua"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetDumpsWithOptions(runtime *util.RuntimeOptions) (_result *GetDumpsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetDumps"),
		Version:     tea.String("2022-08-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDumpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDumps() (_result *GetDumpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDumpsResponse{}
	_body, _err := client.GetDumpsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ShowTrackpointDataServiceWithOptions(request *ShowTrackpointDataServiceRequest, runtime *util.RuntimeOptions) (_result *ShowTrackpointDataServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ShowTrackpointDataService"),
		Version:     tea.String("2022-08-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ShowTrackpointDataServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ShowTrackpointDataService(request *ShowTrackpointDataServiceRequest) (_result *ShowTrackpointDataServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ShowTrackpointDataServiceResponse{}
	_body, _err := client.ShowTrackpointDataServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
