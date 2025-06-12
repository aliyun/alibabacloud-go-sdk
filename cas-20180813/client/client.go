// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ListPCAInstanceRequest struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// example:
	//
	// iTrusChina
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPCAInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPCAInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListPCAInstanceRequest) SetCurrentPage(v int64) *ListPCAInstanceRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListPCAInstanceRequest) SetShowSize(v int64) *ListPCAInstanceRequest {
	s.ShowSize = &v
	return s
}

func (s *ListPCAInstanceRequest) SetType(v string) *ListPCAInstanceRequest {
	s.Type = &v
	return s
}

type ListPCAInstanceResponseBody struct {
	// example:
	//
	// 1
	CurrentPage     *int64                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PCAInstanceList []*ListPCAInstanceResponseBodyPCAInstanceList `json:"PCAInstanceList,omitempty" xml:"PCAInstanceList,omitempty" type:"Repeated"`
	// example:
	//
	// EECA10D5-BD0F-4EF1-B3EA-B4578E5C6F8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPCAInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPCAInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListPCAInstanceResponseBody) SetCurrentPage(v int64) *ListPCAInstanceResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListPCAInstanceResponseBody) SetPCAInstanceList(v []*ListPCAInstanceResponseBodyPCAInstanceList) *ListPCAInstanceResponseBody {
	s.PCAInstanceList = v
	return s
}

func (s *ListPCAInstanceResponseBody) SetRequestId(v string) *ListPCAInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPCAInstanceResponseBody) SetShowSize(v int64) *ListPCAInstanceResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListPCAInstanceResponseBody) SetTotalCount(v int64) *ListPCAInstanceResponseBody {
	s.TotalCount = &v
	return s
}

type ListPCAInstanceResponseBodyPCAInstanceList struct {
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// qingqitest
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// example:
	//
	// 1efb82eb-19e5-620f-bdaa-11cc6cb2a720
	CaIdentifier *string `json:"CaIdentifier,omitempty" xml:"CaIdentifier,omitempty"`
	// example:
	//
	// 10
	CertCount *int64 `json:"CertCount,omitempty" xml:"CertCount,omitempty"`
	// example:
	//
	// 2024-08-09T10:05:18
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1ef3fefc-8065-60de-871e-d15cf842aab6
	InstanceUuid *string `json:"InstanceUuid,omitempty" xml:"InstanceUuid,omitempty"`
	// example:
	//
	// 0
	IssuedCertCount *int64 `json:"IssuedCertCount,omitempty" xml:"IssuedCertCount,omitempty"`
	// example:
	//
	// false
	RelateStatus *bool `json:"RelateStatus,omitempty" xml:"RelateStatus,omitempty"`
	// example:
	//
	// 1
	ShareFlag *int32 `json:"ShareFlag,omitempty" xml:"ShareFlag,omitempty"`
	// example:
	//
	// 0
	Trial *string `json:"Trial,omitempty" xml:"Trial,omitempty"`
}

func (s ListPCAInstanceResponseBodyPCAInstanceList) String() string {
	return tea.Prettify(s)
}

func (s ListPCAInstanceResponseBodyPCAInstanceList) GoString() string {
	return s.String()
}

func (s *ListPCAInstanceResponseBodyPCAInstanceList) SetAlgorithm(v string) *ListPCAInstanceResponseBodyPCAInstanceList {
	s.Algorithm = &v
	return s
}

func (s *ListPCAInstanceResponseBodyPCAInstanceList) SetAliasName(v string) *ListPCAInstanceResponseBodyPCAInstanceList {
	s.AliasName = &v
	return s
}

func (s *ListPCAInstanceResponseBodyPCAInstanceList) SetCaIdentifier(v string) *ListPCAInstanceResponseBodyPCAInstanceList {
	s.CaIdentifier = &v
	return s
}

func (s *ListPCAInstanceResponseBodyPCAInstanceList) SetCertCount(v int64) *ListPCAInstanceResponseBodyPCAInstanceList {
	s.CertCount = &v
	return s
}

func (s *ListPCAInstanceResponseBodyPCAInstanceList) SetEndTime(v int64) *ListPCAInstanceResponseBodyPCAInstanceList {
	s.EndTime = &v
	return s
}

func (s *ListPCAInstanceResponseBodyPCAInstanceList) SetInstanceUuid(v string) *ListPCAInstanceResponseBodyPCAInstanceList {
	s.InstanceUuid = &v
	return s
}

func (s *ListPCAInstanceResponseBodyPCAInstanceList) SetIssuedCertCount(v int64) *ListPCAInstanceResponseBodyPCAInstanceList {
	s.IssuedCertCount = &v
	return s
}

func (s *ListPCAInstanceResponseBodyPCAInstanceList) SetRelateStatus(v bool) *ListPCAInstanceResponseBodyPCAInstanceList {
	s.RelateStatus = &v
	return s
}

func (s *ListPCAInstanceResponseBodyPCAInstanceList) SetShareFlag(v int32) *ListPCAInstanceResponseBodyPCAInstanceList {
	s.ShareFlag = &v
	return s
}

func (s *ListPCAInstanceResponseBodyPCAInstanceList) SetTrial(v string) *ListPCAInstanceResponseBodyPCAInstanceList {
	s.Trial = &v
	return s
}

type ListPCAInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPCAInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPCAInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPCAInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListPCAInstanceResponse) SetHeaders(v map[string]*string) *ListPCAInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListPCAInstanceResponse) SetStatusCode(v int32) *ListPCAInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPCAInstanceResponse) SetBody(v *ListPCAInstanceResponseBody) *ListPCAInstanceResponse {
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
	client.EndpointMap = map[string]*string{
		"cn-hangzhou":                 tea.String("cas.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("cas.aliyuncs.com"),
		"ap-southeast-3":              tea.String("cas.aliyuncs.com"),
		"ap-southeast-5":              tea.String("cas.aliyuncs.com"),
		"cn-beijing":                  tea.String("cas.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("cas.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("cas.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("cas.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("cas.aliyuncs.com"),
		"cn-chengdu":                  tea.String("cas.aliyuncs.com"),
		"cn-edge-1":                   tea.String("cas.aliyuncs.com"),
		"cn-fujian":                   tea.String("cas.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("cas.aliyuncs.com"),
		"cn-hongkong":                 tea.String("cas.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("cas.aliyuncs.com"),
		"cn-huhehaote":                tea.String("cas.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("cas.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("cas.aliyuncs.com"),
		"cn-qingdao":                  tea.String("cas.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("cas.aliyuncs.com"),
		"cn-shanghai":                 tea.String("cas.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("cas.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("cas.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("cas.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("cas.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("cas.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("cas.aliyuncs.com"),
		"cn-wuhan":                    tea.String("cas.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("cas.aliyuncs.com"),
		"cn-yushanfang":               tea.String("cas.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("cas.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("cas.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("cas.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("cas.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("cas.aliyuncs.com"),
		"eu-west-1":                   tea.String("cas.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("cas.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("cas.aliyuncs.com"),
		"us-east-1":                   tea.String("cas.aliyuncs.com"),
		"us-west-1":                   tea.String("cas.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 获取私有证书实例
//
// @param request - ListPCAInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPCAInstanceResponse
func (client *Client) ListPCAInstanceWithOptions(request *ListPCAInstanceRequest, runtime *util.RuntimeOptions) (_result *ListPCAInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPCAInstance"),
		Version:     tea.String("2018-08-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPCAInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取私有证书实例
//
// @param request - ListPCAInstanceRequest
//
// @return ListPCAInstanceResponse
func (client *Client) ListPCAInstance(request *ListPCAInstanceRequest) (_result *ListPCAInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPCAInstanceResponse{}
	_body, _err := client.ListPCAInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
