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

type CreateDiagnosticReportRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s CreateDiagnosticReportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticReportRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticReportRequest) SetDBInstanceId(v string) *CreateDiagnosticReportRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDiagnosticReportRequest) SetStartTime(v string) *CreateDiagnosticReportRequest {
	s.StartTime = &v
	return s
}

func (s *CreateDiagnosticReportRequest) SetEndTime(v string) *CreateDiagnosticReportRequest {
	s.EndTime = &v
	return s
}

func (s *CreateDiagnosticReportRequest) SetCategory(v string) *CreateDiagnosticReportRequest {
	s.Category = &v
	return s
}

type CreateDiagnosticReportResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReportId  *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s CreateDiagnosticReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticReportResponseBody) SetRequestId(v string) *CreateDiagnosticReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiagnosticReportResponseBody) SetReportId(v string) *CreateDiagnosticReportResponseBody {
	s.ReportId = &v
	return s
}

type CreateDiagnosticReportResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDiagnosticReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDiagnosticReportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticReportResponse) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticReportResponse) SetHeaders(v map[string]*string) *CreateDiagnosticReportResponse {
	s.Headers = v
	return s
}

func (s *CreateDiagnosticReportResponse) SetBody(v *CreateDiagnosticReportResponseBody) *CreateDiagnosticReportResponse {
	s.Body = v
	return s
}

type DescribeDiagnosticReportListRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDiagnosticReportListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosticReportListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportListRequest) SetDBInstanceId(v string) *DescribeDiagnosticReportListRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDiagnosticReportListResponseBody struct {
	RequestId  *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReportList []*DescribeDiagnosticReportListResponseBodyReportList `json:"ReportList,omitempty" xml:"ReportList,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosticReportListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosticReportListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportListResponseBody) SetRequestId(v string) *DescribeDiagnosticReportListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBody) SetReportList(v []*DescribeDiagnosticReportListResponseBodyReportList) *DescribeDiagnosticReportListResponseBody {
	s.ReportList = v
	return s
}

type DescribeDiagnosticReportListResponseBodyReportList struct {
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DiagnosticTime *string `json:"DiagnosticTime,omitempty" xml:"DiagnosticTime,omitempty"`
	Score          *int32  `json:"Score,omitempty" xml:"Score,omitempty"`
	DownloadURL    *string `json:"DownloadURL,omitempty" xml:"DownloadURL,omitempty"`
}

func (s DescribeDiagnosticReportListResponseBodyReportList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosticReportListResponseBodyReportList) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportListResponseBodyReportList) SetEndTime(v string) *DescribeDiagnosticReportListResponseBodyReportList {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBodyReportList) SetStartTime(v string) *DescribeDiagnosticReportListResponseBodyReportList {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBodyReportList) SetDiagnosticTime(v string) *DescribeDiagnosticReportListResponseBodyReportList {
	s.DiagnosticTime = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBodyReportList) SetScore(v int32) *DescribeDiagnosticReportListResponseBodyReportList {
	s.Score = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBodyReportList) SetDownloadURL(v string) *DescribeDiagnosticReportListResponseBodyReportList {
	s.DownloadURL = &v
	return s
}

type DescribeDiagnosticReportListResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDiagnosticReportListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosticReportListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosticReportListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportListResponse) SetHeaders(v map[string]*string) *DescribeDiagnosticReportListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosticReportListResponse) SetBody(v *DescribeDiagnosticReportListResponseBody) *DescribeDiagnosticReportListResponse {
	s.Body = v
	return s
}

type GetDBInstanceTopologyRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s GetDBInstanceTopologyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDBInstanceTopologyRequest) GoString() string {
	return s.String()
}

func (s *GetDBInstanceTopologyRequest) SetDBInstanceId(v string) *GetDBInstanceTopologyRequest {
	s.DBInstanceId = &v
	return s
}

type GetDBInstanceTopologyResponseBody struct {
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetDBInstanceTopologyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetDBInstanceTopologyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDBInstanceTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *GetDBInstanceTopologyResponseBody) SetMessage(v string) *GetDBInstanceTopologyResponseBody {
	s.Message = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBody) SetRequestId(v string) *GetDBInstanceTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBody) SetData(v *GetDBInstanceTopologyResponseBodyData) *GetDBInstanceTopologyResponseBody {
	s.Data = v
	return s
}

func (s *GetDBInstanceTopologyResponseBody) SetCode(v string) *GetDBInstanceTopologyResponseBody {
	s.Code = &v
	return s
}

type GetDBInstanceTopologyResponseBodyData struct {
	Nodes          []*GetDBInstanceTopologyResponseBodyDataNodes       `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	DBInstanceName *string                                             `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	Connections    []*GetDBInstanceTopologyResponseBodyDataConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
}

func (s GetDBInstanceTopologyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDBInstanceTopologyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDBInstanceTopologyResponseBodyData) SetNodes(v []*GetDBInstanceTopologyResponseBodyDataNodes) *GetDBInstanceTopologyResponseBodyData {
	s.Nodes = v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyData) SetDBInstanceName(v string) *GetDBInstanceTopologyResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyData) SetConnections(v []*GetDBInstanceTopologyResponseBodyDataConnections) *GetDBInstanceTopologyResponseBodyData {
	s.Connections = v
	return s
}

type GetDBInstanceTopologyResponseBodyDataNodes struct {
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	Role                 *string `json:"Role,omitempty" xml:"Role,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
}

func (s GetDBInstanceTopologyResponseBodyDataNodes) String() string {
	return tea.Prettify(s)
}

func (s GetDBInstanceTopologyResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) SetZoneId(v string) *GetDBInstanceTopologyResponseBodyDataNodes {
	s.ZoneId = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) SetDedicatedHostId(v string) *GetDBInstanceTopologyResponseBodyDataNodes {
	s.DedicatedHostId = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) SetRole(v string) *GetDBInstanceTopologyResponseBodyDataNodes {
	s.Role = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) SetDedicatedHostGroupId(v string) *GetDBInstanceTopologyResponseBodyDataNodes {
	s.DedicatedHostGroupId = &v
	return s
}

type GetDBInstanceTopologyResponseBodyDataConnections struct {
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	NetType          *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	ZoneId           *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetDBInstanceTopologyResponseBodyDataConnections) String() string {
	return tea.Prettify(s)
}

func (s GetDBInstanceTopologyResponseBodyDataConnections) GoString() string {
	return s.String()
}

func (s *GetDBInstanceTopologyResponseBodyDataConnections) SetConnectionString(v string) *GetDBInstanceTopologyResponseBodyDataConnections {
	s.ConnectionString = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataConnections) SetNetType(v string) *GetDBInstanceTopologyResponseBodyDataConnections {
	s.NetType = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataConnections) SetZoneId(v string) *GetDBInstanceTopologyResponseBodyDataConnections {
	s.ZoneId = &v
	return s
}

type GetDBInstanceTopologyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDBInstanceTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDBInstanceTopologyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDBInstanceTopologyResponse) GoString() string {
	return s.String()
}

func (s *GetDBInstanceTopologyResponse) SetHeaders(v map[string]*string) *GetDBInstanceTopologyResponse {
	s.Headers = v
	return s
}

func (s *GetDBInstanceTopologyResponse) SetBody(v *GetDBInstanceTopologyResponseBody) *GetDBInstanceTopologyResponse {
	s.Body = v
	return s
}

type MigrateConnectionToOtherZoneRequest struct {
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	ZoneId           *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s MigrateConnectionToOtherZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s MigrateConnectionToOtherZoneRequest) GoString() string {
	return s.String()
}

func (s *MigrateConnectionToOtherZoneRequest) SetDBInstanceId(v string) *MigrateConnectionToOtherZoneRequest {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateConnectionToOtherZoneRequest) SetConnectionString(v string) *MigrateConnectionToOtherZoneRequest {
	s.ConnectionString = &v
	return s
}

func (s *MigrateConnectionToOtherZoneRequest) SetZoneId(v string) *MigrateConnectionToOtherZoneRequest {
	s.ZoneId = &v
	return s
}

type MigrateConnectionToOtherZoneResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s MigrateConnectionToOtherZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MigrateConnectionToOtherZoneResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateConnectionToOtherZoneResponseBody) SetMessage(v string) *MigrateConnectionToOtherZoneResponseBody {
	s.Message = &v
	return s
}

func (s *MigrateConnectionToOtherZoneResponseBody) SetRequestId(v string) *MigrateConnectionToOtherZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateConnectionToOtherZoneResponseBody) SetCode(v string) *MigrateConnectionToOtherZoneResponseBody {
	s.Code = &v
	return s
}

type MigrateConnectionToOtherZoneResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MigrateConnectionToOtherZoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MigrateConnectionToOtherZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s MigrateConnectionToOtherZoneResponse) GoString() string {
	return s.String()
}

func (s *MigrateConnectionToOtherZoneResponse) SetHeaders(v map[string]*string) *MigrateConnectionToOtherZoneResponse {
	s.Headers = v
	return s
}

func (s *MigrateConnectionToOtherZoneResponse) SetBody(v *MigrateConnectionToOtherZoneResponseBody) *MigrateConnectionToOtherZoneResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceMonitorRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s ModifyDBInstanceMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMonitorRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMonitorRequest) SetOwnerId(v int64) *ModifyDBInstanceMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetClientToken(v string) *ModifyDBInstanceMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetDBInstanceId(v string) *ModifyDBInstanceMonitorRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetPeriod(v string) *ModifyDBInstanceMonitorRequest {
	s.Period = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetOwnerAccount(v string) *ModifyDBInstanceMonitorRequest {
	s.OwnerAccount = &v
	return s
}

type ModifyDBInstanceMonitorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMonitorResponseBody) SetRequestId(v string) *ModifyDBInstanceMonitorResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceMonitorResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMonitorResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMonitorResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceMonitorResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceMonitorResponse) SetBody(v *ModifyDBInstanceMonitorResponseBody) *ModifyDBInstanceMonitorResponse {
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
		"cn-qingdao":                  tea.String("rds.aliyuncs.com"),
		"cn-beijing":                  tea.String("rds.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("rds.aliyuncs.com"),
		"cn-shanghai":                 tea.String("rds.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("rds.aliyuncs.com"),
		"cn-hongkong":                 tea.String("rds.aliyuncs.com"),
		"ap-southeast-1":              tea.String("rds.aliyuncs.com"),
		"us-west-1":                   tea.String("rds.aliyuncs.com"),
		"us-east-1":                   tea.String("rds.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("rds.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("rds.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("rds.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("rds.ap-northeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("rds.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("rds.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("rds.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("rds.aliyuncs.com"),
		"cn-edge-1":                   tea.String("rds.aliyuncs.com"),
		"cn-fujian":                   tea.String("rds.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("rds.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("rds.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("rds.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("rds.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("rds.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("rds.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("rds.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("rds.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("rds.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("rds.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("rds.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("rds.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("rds.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("rds.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("rds.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("rds.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("rds.aliyuncs.com"),
		"cn-wuhan":                    tea.String("rds.aliyuncs.com"),
		"cn-yushanfang":               tea.String("rds.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("rds.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("rds.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("rds.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("rds.ap-northeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("rds.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("rds"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateDiagnosticReportWithOptions(request *CreateDiagnosticReportRequest, runtime *util.RuntimeOptions) (_result *CreateDiagnosticReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDiagnosticReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDiagnosticReport"), tea.String("2014-08-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDiagnosticReport(request *CreateDiagnosticReportRequest) (_result *CreateDiagnosticReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDiagnosticReportResponse{}
	_body, _err := client.CreateDiagnosticReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosticReportListWithOptions(request *DescribeDiagnosticReportListRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosticReportListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDiagnosticReportListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDiagnosticReportList"), tea.String("2014-08-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosticReportList(request *DescribeDiagnosticReportListRequest) (_result *DescribeDiagnosticReportListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosticReportListResponse{}
	_body, _err := client.DescribeDiagnosticReportListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDBInstanceTopologyWithOptions(request *GetDBInstanceTopologyRequest, runtime *util.RuntimeOptions) (_result *GetDBInstanceTopologyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDBInstanceTopologyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDBInstanceTopology"), tea.String("2014-08-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDBInstanceTopology(request *GetDBInstanceTopologyRequest) (_result *GetDBInstanceTopologyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDBInstanceTopologyResponse{}
	_body, _err := client.GetDBInstanceTopologyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MigrateConnectionToOtherZoneWithOptions(request *MigrateConnectionToOtherZoneRequest, runtime *util.RuntimeOptions) (_result *MigrateConnectionToOtherZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MigrateConnectionToOtherZoneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MigrateConnectionToOtherZone"), tea.String("2014-08-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MigrateConnectionToOtherZone(request *MigrateConnectionToOtherZoneRequest) (_result *MigrateConnectionToOtherZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MigrateConnectionToOtherZoneResponse{}
	_body, _err := client.MigrateConnectionToOtherZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceMonitorWithOptions(request *ModifyDBInstanceMonitorRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceMonitor"), tea.String("2014-08-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceMonitor(request *ModifyDBInstanceMonitorRequest) (_result *ModifyDBInstanceMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceMonitorResponse{}
	_body, _err := client.ModifyDBInstanceMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
