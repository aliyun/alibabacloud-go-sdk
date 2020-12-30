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

type ModifyDBClusterMonitorRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s ModifyDBClusterMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMonitorRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMonitorRequest) SetOwnerId(v int64) *ModifyDBClusterMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetResourceOwnerId(v int64) *ModifyDBClusterMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetDBClusterId(v string) *ModifyDBClusterMonitorRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetPeriod(v string) *ModifyDBClusterMonitorRequest {
	s.Period = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetOwnerAccount(v string) *ModifyDBClusterMonitorRequest {
	s.OwnerAccount = &v
	return s
}

type ModifyDBClusterMonitorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMonitorResponseBody) SetRequestId(v string) *ModifyDBClusterMonitorResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterMonitorResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMonitorResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMonitorResponse) SetHeaders(v map[string]*string) *ModifyDBClusterMonitorResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterMonitorResponse) SetBody(v *ModifyDBClusterMonitorResponseBody) *ModifyDBClusterMonitorResponse {
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
		"cn-qingdao":                  tea.String("polardb.aliyuncs.com"),
		"cn-beijing":                  tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("polardb.aliyuncs.com"),
		"cn-shanghai":                 tea.String("polardb.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("polardb.aliyuncs.com"),
		"cn-hongkong":                 tea.String("polardb.aliyuncs.com"),
		"ap-southeast-1":              tea.String("polardb.aliyuncs.com"),
		"us-west-1":                   tea.String("polardb.aliyuncs.com"),
		"us-east-1":                   tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("polardb.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("polardb.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("polardb.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("polardb.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("polardb.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("polardb.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("polardb.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("polardb.aliyuncs.com"),
		"cn-edge-1":                   tea.String("polardb.aliyuncs.com"),
		"cn-fujian":                   tea.String("polardb.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("polardb.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("polardb.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("polardb.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("polardb.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("polardb.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("polardb.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("polardb.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("polardb.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("polardb.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("polardb.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("polardb.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("polardb.aliyuncs.com"),
		"cn-wuhan":                    tea.String("polardb.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("polardb.aliyuncs.com"),
		"cn-yushanfang":               tea.String("polardb.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("polardb.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("polardb.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("polardb.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("polardb.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("polardb.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("polardb.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("polardb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ModifyDBClusterMonitorWithOptions(request *ModifyDBClusterMonitorRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterMonitor"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterMonitor(request *ModifyDBClusterMonitorRequest) (_result *ModifyDBClusterMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterMonitorResponse{}
	_body, _err := client.ModifyDBClusterMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
