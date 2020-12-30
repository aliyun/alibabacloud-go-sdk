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

type AddFaceRequest struct {
	Group    *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Person   *string `json:"Person,omitempty" xml:"Person,omitempty"`
	Image    *string `json:"Image,omitempty" xml:"Image,omitempty"`
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s AddFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceRequest) GoString() string {
	return s.String()
}

func (s *AddFaceRequest) SetGroup(v string) *AddFaceRequest {
	s.Group = &v
	return s
}

func (s *AddFaceRequest) SetPerson(v string) *AddFaceRequest {
	s.Person = &v
	return s
}

func (s *AddFaceRequest) SetImage(v string) *AddFaceRequest {
	s.Image = &v
	return s
}

func (s *AddFaceRequest) SetImageUrl(v string) *AddFaceRequest {
	s.ImageUrl = &v
	return s
}

func (s *AddFaceRequest) SetContent(v string) *AddFaceRequest {
	s.Content = &v
	return s
}

type AddFaceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceResponseBody) SetMessage(v string) *AddFaceResponseBody {
	s.Message = &v
	return s
}

func (s *AddFaceResponseBody) SetRequestId(v string) *AddFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceResponseBody) SetData(v string) *AddFaceResponseBody {
	s.Data = &v
	return s
}

func (s *AddFaceResponseBody) SetCode(v string) *AddFaceResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceResponseBody) SetSuccess(v bool) *AddFaceResponseBody {
	s.Success = &v
	return s
}

type AddFaceResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceResponse) GoString() string {
	return s.String()
}

func (s *AddFaceResponse) SetHeaders(v map[string]*string) *AddFaceResponse {
	s.Headers = v
	return s
}

func (s *AddFaceResponse) SetBody(v *AddFaceResponseBody) *AddFaceResponse {
	s.Body = v
	return s
}

type DeleteFaceRequest struct {
	Group  *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Person *string `json:"Person,omitempty" xml:"Person,omitempty"`
	Image  *string `json:"Image,omitempty" xml:"Image,omitempty"`
}

func (s DeleteFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceRequest) SetGroup(v string) *DeleteFaceRequest {
	s.Group = &v
	return s
}

func (s *DeleteFaceRequest) SetPerson(v string) *DeleteFaceRequest {
	s.Person = &v
	return s
}

func (s *DeleteFaceRequest) SetImage(v string) *DeleteFaceRequest {
	s.Image = &v
	return s
}

type DeleteFaceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceResponseBody) SetMessage(v string) *DeleteFaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFaceResponseBody) SetRequestId(v string) *DeleteFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceResponseBody) SetData(v string) *DeleteFaceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFaceResponseBody) SetCode(v string) *DeleteFaceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceResponseBody) SetSuccess(v bool) *DeleteFaceResponseBody {
	s.Success = &v
	return s
}

type DeleteFaceResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceResponse) SetHeaders(v map[string]*string) *DeleteFaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceResponse) SetBody(v *DeleteFaceResponseBody) *DeleteFaceResponse {
	s.Body = v
	return s
}

type DetectFaceRequest struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s DetectFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceRequest) SetImageUrl(v string) *DetectFaceRequest {
	s.ImageUrl = &v
	return s
}

func (s *DetectFaceRequest) SetContent(v string) *DetectFaceRequest {
	s.Content = &v
	return s
}

type DetectFaceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetectFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetectFaceResponseBody) SetMessage(v string) *DetectFaceResponseBody {
	s.Message = &v
	return s
}

func (s *DetectFaceResponseBody) SetRequestId(v string) *DetectFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectFaceResponseBody) SetData(v string) *DetectFaceResponseBody {
	s.Data = &v
	return s
}

func (s *DetectFaceResponseBody) SetCode(v string) *DetectFaceResponseBody {
	s.Code = &v
	return s
}

func (s *DetectFaceResponseBody) SetSuccess(v bool) *DetectFaceResponseBody {
	s.Success = &v
	return s
}

type DetectFaceResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceResponse) GoString() string {
	return s.String()
}

func (s *DetectFaceResponse) SetHeaders(v map[string]*string) *DetectFaceResponse {
	s.Headers = v
	return s
}

func (s *DetectFaceResponse) SetBody(v *DetectFaceResponseBody) *DetectFaceResponse {
	s.Body = v
	return s
}

type GetFaceAttributeRequest struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s GetFaceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFaceAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetFaceAttributeRequest) SetImageUrl(v string) *GetFaceAttributeRequest {
	s.ImageUrl = &v
	return s
}

func (s *GetFaceAttributeRequest) SetContent(v string) *GetFaceAttributeRequest {
	s.Content = &v
	return s
}

type GetFaceAttributeResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFaceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFaceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetFaceAttributeResponseBody) SetMessage(v string) *GetFaceAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *GetFaceAttributeResponseBody) SetRequestId(v string) *GetFaceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFaceAttributeResponseBody) SetData(v string) *GetFaceAttributeResponseBody {
	s.Data = &v
	return s
}

func (s *GetFaceAttributeResponseBody) SetCode(v string) *GetFaceAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *GetFaceAttributeResponseBody) SetSuccess(v bool) *GetFaceAttributeResponseBody {
	s.Success = &v
	return s
}

type GetFaceAttributeResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFaceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFaceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFaceAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetFaceAttributeResponse) SetHeaders(v map[string]*string) *GetFaceAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetFaceAttributeResponse) SetBody(v *GetFaceAttributeResponseBody) *GetFaceAttributeResponse {
	s.Body = v
	return s
}

type ListFaceRequest struct {
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Mark  *int64  `json:"Mark,omitempty" xml:"Mark,omitempty"`
}

func (s ListFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFaceRequest) GoString() string {
	return s.String()
}

func (s *ListFaceRequest) SetGroup(v string) *ListFaceRequest {
	s.Group = &v
	return s
}

func (s *ListFaceRequest) SetMark(v int64) *ListFaceRequest {
	s.Mark = &v
	return s
}

type ListFaceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFaceResponseBody) GoString() string {
	return s.String()
}

func (s *ListFaceResponseBody) SetMessage(v string) *ListFaceResponseBody {
	s.Message = &v
	return s
}

func (s *ListFaceResponseBody) SetRequestId(v string) *ListFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFaceResponseBody) SetData(v string) *ListFaceResponseBody {
	s.Data = &v
	return s
}

func (s *ListFaceResponseBody) SetCode(v string) *ListFaceResponseBody {
	s.Code = &v
	return s
}

func (s *ListFaceResponseBody) SetSuccess(v bool) *ListFaceResponseBody {
	s.Success = &v
	return s
}

type ListFaceResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFaceResponse) GoString() string {
	return s.String()
}

func (s *ListFaceResponse) SetHeaders(v map[string]*string) *ListFaceResponse {
	s.Headers = v
	return s
}

func (s *ListFaceResponse) SetBody(v *ListFaceResponseBody) *ListFaceResponse {
	s.Body = v
	return s
}

type ListGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupResponseBody) SetMessage(v string) *ListGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListGroupResponseBody) SetRequestId(v string) *ListGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupResponseBody) SetData(v string) *ListGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ListGroupResponseBody) SetCode(v string) *ListGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListGroupResponseBody) SetSuccess(v bool) *ListGroupResponseBody {
	s.Success = &v
	return s
}

type ListGroupResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupResponse) GoString() string {
	return s.String()
}

func (s *ListGroupResponse) SetHeaders(v map[string]*string) *ListGroupResponse {
	s.Headers = v
	return s
}

func (s *ListGroupResponse) SetBody(v *ListGroupResponseBody) *ListGroupResponse {
	s.Body = v
	return s
}

type RecognizeFaceRequest struct {
	Group    *string `json:"Group,omitempty" xml:"Group,omitempty"`
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s RecognizeFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFaceRequest) SetGroup(v string) *RecognizeFaceRequest {
	s.Group = &v
	return s
}

func (s *RecognizeFaceRequest) SetImageUrl(v string) *RecognizeFaceRequest {
	s.ImageUrl = &v
	return s
}

func (s *RecognizeFaceRequest) SetContent(v string) *RecognizeFaceRequest {
	s.Content = &v
	return s
}

type RecognizeFaceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RecognizeFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeFaceResponseBody) SetMessage(v string) *RecognizeFaceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeFaceResponseBody) SetRequestId(v string) *RecognizeFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeFaceResponseBody) SetData(v string) *RecognizeFaceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeFaceResponseBody) SetCode(v string) *RecognizeFaceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeFaceResponseBody) SetSuccess(v bool) *RecognizeFaceResponseBody {
	s.Success = &v
	return s
}

type RecognizeFaceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFaceResponse) SetHeaders(v map[string]*string) *RecognizeFaceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeFaceResponse) SetBody(v *RecognizeFaceResponseBody) *RecognizeFaceResponse {
	s.Body = v
	return s
}

type VerifyFaceRequest struct {
	ImageUrl1 *string `json:"ImageUrl1,omitempty" xml:"ImageUrl1,omitempty"`
	ImageUrl2 *string `json:"ImageUrl2,omitempty" xml:"ImageUrl2,omitempty"`
	Content1  *string `json:"Content1,omitempty" xml:"Content1,omitempty"`
	Content2  *string `json:"Content2,omitempty" xml:"Content2,omitempty"`
}

func (s VerifyFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyFaceRequest) GoString() string {
	return s.String()
}

func (s *VerifyFaceRequest) SetImageUrl1(v string) *VerifyFaceRequest {
	s.ImageUrl1 = &v
	return s
}

func (s *VerifyFaceRequest) SetImageUrl2(v string) *VerifyFaceRequest {
	s.ImageUrl2 = &v
	return s
}

func (s *VerifyFaceRequest) SetContent1(v string) *VerifyFaceRequest {
	s.Content1 = &v
	return s
}

func (s *VerifyFaceRequest) SetContent2(v string) *VerifyFaceRequest {
	s.Content2 = &v
	return s
}

type VerifyFaceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyFaceResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyFaceResponseBody) SetMessage(v string) *VerifyFaceResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyFaceResponseBody) SetRequestId(v string) *VerifyFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyFaceResponseBody) SetData(v string) *VerifyFaceResponseBody {
	s.Data = &v
	return s
}

func (s *VerifyFaceResponseBody) SetCode(v string) *VerifyFaceResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyFaceResponseBody) SetSuccess(v bool) *VerifyFaceResponseBody {
	s.Success = &v
	return s
}

type VerifyFaceResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyFaceResponse) GoString() string {
	return s.String()
}

func (s *VerifyFaceResponse) SetHeaders(v map[string]*string) *VerifyFaceResponse {
	s.Headers = v
	return s
}

func (s *VerifyFaceResponse) SetBody(v *VerifyFaceResponseBody) *VerifyFaceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("face"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddFaceWithOptions(request *AddFaceRequest, runtime *util.RuntimeOptions) (_result *AddFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddFace"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFace(request *AddFaceRequest) (_result *AddFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceResponse{}
	_body, _err := client.AddFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceWithOptions(request *DeleteFaceRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFace"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFace(request *DeleteFaceRequest) (_result *DeleteFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceResponse{}
	_body, _err := client.DeleteFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectFaceWithOptions(request *DetectFaceRequest, runtime *util.RuntimeOptions) (_result *DetectFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectFace"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectFace(request *DetectFaceRequest) (_result *DetectFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectFaceResponse{}
	_body, _err := client.DetectFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFaceAttributeWithOptions(request *GetFaceAttributeRequest, runtime *util.RuntimeOptions) (_result *GetFaceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetFaceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetFaceAttribute"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFaceAttribute(request *GetFaceAttributeRequest) (_result *GetFaceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFaceAttributeResponse{}
	_body, _err := client.GetFaceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFaceWithOptions(request *ListFaceRequest, runtime *util.RuntimeOptions) (_result *ListFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFace"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFace(request *ListFaceRequest) (_result *ListFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFaceResponse{}
	_body, _err := client.ListFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGroupWithOptions(runtime *util.RuntimeOptions) (_result *ListGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListGroup"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGroup() (_result *ListGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupResponse{}
	_body, _err := client.ListGroupWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeFaceWithOptions(request *RecognizeFaceRequest, runtime *util.RuntimeOptions) (_result *RecognizeFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeFace"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeFace(request *RecognizeFaceRequest) (_result *RecognizeFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeFaceResponse{}
	_body, _err := client.RecognizeFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyFaceWithOptions(request *VerifyFaceRequest, runtime *util.RuntimeOptions) (_result *VerifyFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyFace"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyFace(request *VerifyFaceRequest) (_result *VerifyFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyFaceResponse{}
	_body, _err := client.VerifyFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
