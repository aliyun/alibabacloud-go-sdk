// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ListDeliveryHistoryJobsRequest struct {
}

func (s ListDeliveryHistoryJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryHistoryJobsRequest) GoString() string {
	return s.String()
}

type ListDeliveryHistoryJobsResponse struct {
	RequestId           *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount          *int                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	DeliveryHistoryJobs []*ListDeliveryHistoryJobsResponseDeliveryHistoryJobs `json:"DeliveryHistoryJobs,omitempty" xml:"DeliveryHistoryJobs,omitempty" require:"true" type:"Repeated"`
}

func (s ListDeliveryHistoryJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryHistoryJobsResponse) GoString() string {
	return s.String()
}

func (s *ListDeliveryHistoryJobsResponse) SetRequestId(v string) *ListDeliveryHistoryJobsResponse {
	s.RequestId = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponse) SetTotalCount(v int) *ListDeliveryHistoryJobsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponse) SetDeliveryHistoryJobs(v []*ListDeliveryHistoryJobsResponseDeliveryHistoryJobs) *ListDeliveryHistoryJobsResponse {
	s.DeliveryHistoryJobs = v
	return s
}

type ListDeliveryHistoryJobsResponseDeliveryHistoryJobs struct {
	TrailName   *string `json:"TrailName,omitempty" xml:"TrailName,omitempty" require:"true"`
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty" require:"true"`
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty" require:"true"`
	HomeRegion  *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty" require:"true"`
}

func (s ListDeliveryHistoryJobsResponseDeliveryHistoryJobs) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryHistoryJobsResponseDeliveryHistoryJobs) GoString() string {
	return s.String()
}

func (s *ListDeliveryHistoryJobsResponseDeliveryHistoryJobs) SetTrailName(v string) *ListDeliveryHistoryJobsResponseDeliveryHistoryJobs {
	s.TrailName = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseDeliveryHistoryJobs) SetCreatedTime(v string) *ListDeliveryHistoryJobsResponseDeliveryHistoryJobs {
	s.CreatedTime = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseDeliveryHistoryJobs) SetUpdatedTime(v string) *ListDeliveryHistoryJobsResponseDeliveryHistoryJobs {
	s.UpdatedTime = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseDeliveryHistoryJobs) SetHomeRegion(v string) *ListDeliveryHistoryJobsResponseDeliveryHistoryJobs {
	s.HomeRegion = &v
	return s
}

type CreateDeliveryHistoryJobRequest struct {
	TrailName   *string `json:"TrailName,omitempty" xml:"TrailName,omitempty" require:"true"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateDeliveryHistoryJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryHistoryJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDeliveryHistoryJobRequest) SetTrailName(v string) *CreateDeliveryHistoryJobRequest {
	s.TrailName = &v
	return s
}

func (s *CreateDeliveryHistoryJobRequest) SetClientToken(v string) *CreateDeliveryHistoryJobRequest {
	s.ClientToken = &v
	return s
}

type CreateDeliveryHistoryJobResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateDeliveryHistoryJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryHistoryJobResponse) GoString() string {
	return s.String()
}

func (s *CreateDeliveryHistoryJobResponse) SetRequestId(v string) *CreateDeliveryHistoryJobResponse {
	s.RequestId = &v
	return s
}

type LookupEventsRequest struct {
	NextToken       *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults      *string                               `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	StartTime       *string                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LookupAttribute []*LookupEventsRequestLookupAttribute `json:"LookupAttribute,omitempty" xml:"LookupAttribute,omitempty" type:"Repeated"`
	Direction       *string                               `json:"Direction,omitempty" xml:"Direction,omitempty"`
}

func (s LookupEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s LookupEventsRequest) GoString() string {
	return s.String()
}

func (s *LookupEventsRequest) SetNextToken(v string) *LookupEventsRequest {
	s.NextToken = &v
	return s
}

func (s *LookupEventsRequest) SetMaxResults(v string) *LookupEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *LookupEventsRequest) SetStartTime(v string) *LookupEventsRequest {
	s.StartTime = &v
	return s
}

func (s *LookupEventsRequest) SetEndTime(v string) *LookupEventsRequest {
	s.EndTime = &v
	return s
}

func (s *LookupEventsRequest) SetLookupAttribute(v []*LookupEventsRequestLookupAttribute) *LookupEventsRequest {
	s.LookupAttribute = v
	return s
}

func (s *LookupEventsRequest) SetDirection(v string) *LookupEventsRequest {
	s.Direction = &v
	return s
}

type LookupEventsRequestLookupAttribute struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s LookupEventsRequestLookupAttribute) String() string {
	return tea.Prettify(s)
}

func (s LookupEventsRequestLookupAttribute) GoString() string {
	return s.String()
}

func (s *LookupEventsRequestLookupAttribute) SetKey(v string) *LookupEventsRequestLookupAttribute {
	s.Key = &v
	return s
}

func (s *LookupEventsRequestLookupAttribute) SetValue(v string) *LookupEventsRequestLookupAttribute {
	s.Value = &v
	return s
}

type LookupEventsResponse struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken *string                  `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	StartTime *string                  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *string                  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Events    []map[string]interface{} `json:"Events,omitempty" xml:"Events,omitempty" require:"true" type:"Repeated"`
}

func (s LookupEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s LookupEventsResponse) GoString() string {
	return s.String()
}

func (s *LookupEventsResponse) SetRequestId(v string) *LookupEventsResponse {
	s.RequestId = &v
	return s
}

func (s *LookupEventsResponse) SetNextToken(v string) *LookupEventsResponse {
	s.NextToken = &v
	return s
}

func (s *LookupEventsResponse) SetStartTime(v string) *LookupEventsResponse {
	s.StartTime = &v
	return s
}

func (s *LookupEventsResponse) SetEndTime(v string) *LookupEventsResponse {
	s.EndTime = &v
	return s
}

func (s *LookupEventsResponse) SetEvents(v []map[string]interface{}) *LookupEventsResponse {
	s.Events = v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("actiontrail.ap-northeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("actiontrail.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("actiontrail.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("actiontrail.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("actiontrail.aliyuncs.com"),
		"cn-edge-1":                   tea.String("actiontrail.aliyuncs.com"),
		"cn-fujian":                   tea.String("actiontrail.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("actiontrail.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("actiontrail.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("actiontrail.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("actiontrail.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("actiontrail.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("actiontrail.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("actiontrail.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("actiontrail.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("actiontrail.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("actiontrail.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("actiontrail.aliyuncs.com"),
		"cn-wuhan":                    tea.String("actiontrail.aliyuncs.com"),
		"cn-yushanfang":               tea.String("actiontrail.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("actiontrail.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("actiontrail.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("actiontrail.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("actiontrail.ap-northeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("actiontrail.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("actiontrail"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) ListDeliveryHistoryJobsWithOptions(request *ListDeliveryHistoryJobsRequest, runtime *util.RuntimeOptions) (_result *ListDeliveryHistoryJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListDeliveryHistoryJobsResponse{}
	_body, _err := client.DoRequest(tea.String("ListDeliveryHistoryJobs"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-06"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeliveryHistoryJobs(request *ListDeliveryHistoryJobsRequest) (_result *ListDeliveryHistoryJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeliveryHistoryJobsResponse{}
	_body, _err := client.ListDeliveryHistoryJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeliveryHistoryJobWithOptions(request *CreateDeliveryHistoryJobRequest, runtime *util.RuntimeOptions) (_result *CreateDeliveryHistoryJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDeliveryHistoryJobResponse{}
	_body, _err := client.DoRequest(tea.String("CreateDeliveryHistoryJob"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-06"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeliveryHistoryJob(request *CreateDeliveryHistoryJobRequest) (_result *CreateDeliveryHistoryJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeliveryHistoryJobResponse{}
	_body, _err := client.CreateDeliveryHistoryJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LookupEventsWithOptions(request *LookupEventsRequest, runtime *util.RuntimeOptions) (_result *LookupEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &LookupEventsResponse{}
	_body, _err := client.DoRequest(tea.String("LookupEvents"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-06"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LookupEvents(request *LookupEventsRequest) (_result *LookupEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LookupEventsResponse{}
	_body, _err := client.LookupEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
