// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetRealTimeRiskInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Ab63aqUKx9TmYRhPV7LnRLybLRQOEG6GQsr-oGlT8bo6-_z_oYxTqqutBz6R0dBognMN1DBvZImDadsKb8k_Hfcj8A8nrXtxch_nKVg5xARkRmMU2SU=
	Atoken *string `json:"atoken,omitempty" xml:"atoken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 62c792d188ccdf4b7ec1aa38
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// example:
	//
	// 360risk
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s GetRealTimeRiskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRealTimeRiskInfoRequest) GoString() string {
	return s.String()
}

func (s *GetRealTimeRiskInfoRequest) SetAtoken(v string) *GetRealTimeRiskInfoRequest {
	s.Atoken = &v
	return s
}

func (s *GetRealTimeRiskInfoRequest) SetDataSourceId(v string) *GetRealTimeRiskInfoRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetRealTimeRiskInfoRequest) SetExtra(v string) *GetRealTimeRiskInfoRequest {
	s.Extra = &v
	return s
}

type GetRealTimeRiskInfoResponseBody struct {
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 200
	Code *int64                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetRealTimeRiskInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 2343535353DF8984565
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetRealTimeRiskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRealTimeRiskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealTimeRiskInfoResponseBody) SetMsg(v string) *GetRealTimeRiskInfoResponseBody {
	s.Msg = &v
	return s
}

func (s *GetRealTimeRiskInfoResponseBody) SetCode(v int64) *GetRealTimeRiskInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetRealTimeRiskInfoResponseBody) SetData(v *GetRealTimeRiskInfoResponseBodyData) *GetRealTimeRiskInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetRealTimeRiskInfoResponseBody) SetRequestId(v string) *GetRealTimeRiskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRealTimeRiskInfoResponseBody) SetSuccess(v bool) *GetRealTimeRiskInfoResponseBody {
	s.Success = &v
	return s
}

type GetRealTimeRiskInfoResponseBodyData struct {
	// example:
	//
	// 360market
	AppChannel *string `json:"appChannel,omitempty" xml:"appChannel,omitempty"`
	// example:
	//
	// 1
	FakeDevice *string `json:"fakeDevice,omitempty" xml:"fakeDevice,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-000000000000
	Idfa *string `json:"idfa,omitempty" xml:"idfa,omitempty"`
	// example:
	//
	// sf58rw7erdfdgdgseeg2434
	Oaid *string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// example:
	//
	// 0
	ProxyDevice *string `json:"proxyDevice,omitempty" xml:"proxyDevice,omitempty"`
	// example:
	//
	// RISK
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	// example:
	//
	// 60
	RiskScore *string `json:"riskScore,omitempty" xml:"riskScore,omitempty"`
	// example:
	//
	// 815E045B-4A9E-AA76-BEFF-048C9B9F651A
	Zid *string `json:"zid,omitempty" xml:"zid,omitempty"`
}

func (s GetRealTimeRiskInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRealTimeRiskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRealTimeRiskInfoResponseBodyData) SetAppChannel(v string) *GetRealTimeRiskInfoResponseBodyData {
	s.AppChannel = &v
	return s
}

func (s *GetRealTimeRiskInfoResponseBodyData) SetFakeDevice(v string) *GetRealTimeRiskInfoResponseBodyData {
	s.FakeDevice = &v
	return s
}

func (s *GetRealTimeRiskInfoResponseBodyData) SetIdfa(v string) *GetRealTimeRiskInfoResponseBodyData {
	s.Idfa = &v
	return s
}

func (s *GetRealTimeRiskInfoResponseBodyData) SetOaid(v string) *GetRealTimeRiskInfoResponseBodyData {
	s.Oaid = &v
	return s
}

func (s *GetRealTimeRiskInfoResponseBodyData) SetProxyDevice(v string) *GetRealTimeRiskInfoResponseBodyData {
	s.ProxyDevice = &v
	return s
}

func (s *GetRealTimeRiskInfoResponseBodyData) SetRiskLevel(v string) *GetRealTimeRiskInfoResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *GetRealTimeRiskInfoResponseBodyData) SetRiskScore(v string) *GetRealTimeRiskInfoResponseBodyData {
	s.RiskScore = &v
	return s
}

func (s *GetRealTimeRiskInfoResponseBodyData) SetZid(v string) *GetRealTimeRiskInfoResponseBodyData {
	s.Zid = &v
	return s
}

type GetRealTimeRiskInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRealTimeRiskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRealTimeRiskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRealTimeRiskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetRealTimeRiskInfoResponse) SetHeaders(v map[string]*string) *GetRealTimeRiskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetRealTimeRiskInfoResponse) SetStatusCode(v int32) *GetRealTimeRiskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealTimeRiskInfoResponse) SetBody(v *GetRealTimeRiskInfoResponseBody) *GetRealTimeRiskInfoResponse {
	s.Body = v
	return s
}

type GetZidTagByAtokenRequest struct {
	// atoken
	//
	// This parameter is required.
	//
	// example:
	//
	// AcNh1pZFEJ9wEO1RkqK2S9Qt0DHxZR6riGmx4CpGrxCpf3_LxvQNQQGaLjGWRg_-Un1yC47qfEiwRysoCJgQrBSa7b274rzhuIW4QDer7J41X18PiHI=
	Atoken *string `json:"atoken,omitempty" xml:"atoken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 59892ebcaed179694b000104
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
}

func (s GetZidTagByAtokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetZidTagByAtokenRequest) GoString() string {
	return s.String()
}

func (s *GetZidTagByAtokenRequest) SetAtoken(v string) *GetZidTagByAtokenRequest {
	s.Atoken = &v
	return s
}

func (s *GetZidTagByAtokenRequest) SetDataSourceId(v string) *GetZidTagByAtokenRequest {
	s.DataSourceId = &v
	return s
}

type GetZidTagByAtokenResponseBody struct {
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// code
	//
	// example:
	//
	// 200
	Code *int64                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetZidTagByAtokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 2343535353DF8984565
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// success
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetZidTagByAtokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetZidTagByAtokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetZidTagByAtokenResponseBody) SetMsg(v string) *GetZidTagByAtokenResponseBody {
	s.Msg = &v
	return s
}

func (s *GetZidTagByAtokenResponseBody) SetCode(v int64) *GetZidTagByAtokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetZidTagByAtokenResponseBody) SetData(v *GetZidTagByAtokenResponseBodyData) *GetZidTagByAtokenResponseBody {
	s.Data = v
	return s
}

func (s *GetZidTagByAtokenResponseBody) SetRequestId(v string) *GetZidTagByAtokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetZidTagByAtokenResponseBody) SetSuccess(v bool) *GetZidTagByAtokenResponseBody {
	s.Success = &v
	return s
}

type GetZidTagByAtokenResponseBodyData struct {
	// aHook
	//
	// example:
	//
	// 1
	AHook *string `json:"aHook,omitempty" xml:"aHook,omitempty"`
	// example:
	//
	// 1
	Debug *string `json:"debug,omitempty" xml:"debug,omitempty"`
	// example:
	//
	// 0
	DoubleOpen *string `json:"doubleOpen,omitempty" xml:"doubleOpen,omitempty"`
	// javaHook
	//
	// example:
	//
	// 1
	JavaHook *string `json:"javaHook,omitempty" xml:"javaHook,omitempty"`
	// example:
	//
	// 1
	NativeHook *string `json:"nativeHook,omitempty" xml:"nativeHook,omitempty"`
	// example:
	//
	// 1
	Root *string `json:"root,omitempty" xml:"root,omitempty"`
	// example:
	//
	// 1
	Simulator *string `json:"simulator,omitempty" xml:"simulator,omitempty"`
	// example:
	//
	// 0
	VpnProxy *string `json:"vpnProxy,omitempty" xml:"vpnProxy,omitempty"`
	// example:
	//
	// 1
	WifiProxy *string `json:"wifiProxy,omitempty" xml:"wifiProxy,omitempty"`
	// example:
	//
	// F6DD4AD2-B7B8-AA3A-34B8-EDE55FC82FC6
	Zid *string `json:"zid,omitempty" xml:"zid,omitempty"`
}

func (s GetZidTagByAtokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetZidTagByAtokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetZidTagByAtokenResponseBodyData) SetAHook(v string) *GetZidTagByAtokenResponseBodyData {
	s.AHook = &v
	return s
}

func (s *GetZidTagByAtokenResponseBodyData) SetDebug(v string) *GetZidTagByAtokenResponseBodyData {
	s.Debug = &v
	return s
}

func (s *GetZidTagByAtokenResponseBodyData) SetDoubleOpen(v string) *GetZidTagByAtokenResponseBodyData {
	s.DoubleOpen = &v
	return s
}

func (s *GetZidTagByAtokenResponseBodyData) SetJavaHook(v string) *GetZidTagByAtokenResponseBodyData {
	s.JavaHook = &v
	return s
}

func (s *GetZidTagByAtokenResponseBodyData) SetNativeHook(v string) *GetZidTagByAtokenResponseBodyData {
	s.NativeHook = &v
	return s
}

func (s *GetZidTagByAtokenResponseBodyData) SetRoot(v string) *GetZidTagByAtokenResponseBodyData {
	s.Root = &v
	return s
}

func (s *GetZidTagByAtokenResponseBodyData) SetSimulator(v string) *GetZidTagByAtokenResponseBodyData {
	s.Simulator = &v
	return s
}

func (s *GetZidTagByAtokenResponseBodyData) SetVpnProxy(v string) *GetZidTagByAtokenResponseBodyData {
	s.VpnProxy = &v
	return s
}

func (s *GetZidTagByAtokenResponseBodyData) SetWifiProxy(v string) *GetZidTagByAtokenResponseBodyData {
	s.WifiProxy = &v
	return s
}

func (s *GetZidTagByAtokenResponseBodyData) SetZid(v string) *GetZidTagByAtokenResponseBodyData {
	s.Zid = &v
	return s
}

type GetZidTagByAtokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetZidTagByAtokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetZidTagByAtokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetZidTagByAtokenResponse) GoString() string {
	return s.String()
}

func (s *GetZidTagByAtokenResponse) SetHeaders(v map[string]*string) *GetZidTagByAtokenResponse {
	s.Headers = v
	return s
}

func (s *GetZidTagByAtokenResponse) SetStatusCode(v int32) *GetZidTagByAtokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetZidTagByAtokenResponse) SetBody(v *GetZidTagByAtokenResponseBody) *GetZidTagByAtokenResponse {
	s.Body = v
	return s
}

type GetZidTagScoreByAtokenRequest struct {
	// atoken
	//
	// This parameter is required.
	//
	// example:
	//
	// AcNh1pZFEJ9wEO1RkqK2S9Qt0DHxZR6riGmx4CpGrxCpf3_LxvQNQQGaLjGWRg_-Un1yC47qfEiwRysoCJgQrBSa7b274rzhuIW4QDer7J41X18PiHI=
	Atoken *string `json:"atoken,omitempty" xml:"atoken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 59892ebcaed179694b000104
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
}

func (s GetZidTagScoreByAtokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetZidTagScoreByAtokenRequest) GoString() string {
	return s.String()
}

func (s *GetZidTagScoreByAtokenRequest) SetAtoken(v string) *GetZidTagScoreByAtokenRequest {
	s.Atoken = &v
	return s
}

func (s *GetZidTagScoreByAtokenRequest) SetDataSourceId(v string) *GetZidTagScoreByAtokenRequest {
	s.DataSourceId = &v
	return s
}

type GetZidTagScoreByAtokenResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code *int64                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Msg  *string                                 `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Data *GetZidTagScoreByAtokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 2343535353DF8984565
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// success
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetZidTagScoreByAtokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetZidTagScoreByAtokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetZidTagScoreByAtokenResponseBody) SetCode(v int64) *GetZidTagScoreByAtokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBody) SetMsg(v string) *GetZidTagScoreByAtokenResponseBody {
	s.Msg = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBody) SetData(v *GetZidTagScoreByAtokenResponseBodyData) *GetZidTagScoreByAtokenResponseBody {
	s.Data = v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBody) SetRequestId(v string) *GetZidTagScoreByAtokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBody) SetSuccess(v bool) *GetZidTagScoreByAtokenResponseBody {
	s.Success = &v
	return s
}

type GetZidTagScoreByAtokenResponseBodyData struct {
	// aHook
	//
	// example:
	//
	// 0
	AHook *string `json:"aHook,omitempty" xml:"aHook,omitempty"`
	// example:
	//
	// 1
	Debug *string `json:"debug,omitempty" xml:"debug,omitempty"`
	// example:
	//
	// 1
	DoubleOpen *string `json:"doubleOpen,omitempty" xml:"doubleOpen,omitempty"`
	// javaHook
	//
	// example:
	//
	// 0
	JavaHook *string `json:"javaHook,omitempty" xml:"javaHook,omitempty"`
	// example:
	//
	// 0
	NativeHook *string `json:"nativeHook,omitempty" xml:"nativeHook,omitempty"`
	// example:
	//
	// RISK
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	// example:
	//
	// 60
	RiskScore *string `json:"riskScore,omitempty" xml:"riskScore,omitempty"`
	// example:
	//
	// 0
	Root *string `json:"root,omitempty" xml:"root,omitempty"`
	// example:
	//
	// 1
	Simulator *string `json:"simulator,omitempty" xml:"simulator,omitempty"`
	// example:
	//
	// 0
	VpnProxy *string `json:"vpnProxy,omitempty" xml:"vpnProxy,omitempty"`
	// example:
	//
	// 0
	WifiProxy *string `json:"wifiProxy,omitempty" xml:"wifiProxy,omitempty"`
	// example:
	//
	// F6DD4AD2-B7B8-AA3A-34B8-EDE55FC82FC6
	Zid *string `json:"zid,omitempty" xml:"zid,omitempty"`
}

func (s GetZidTagScoreByAtokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetZidTagScoreByAtokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetZidTagScoreByAtokenResponseBodyData) SetAHook(v string) *GetZidTagScoreByAtokenResponseBodyData {
	s.AHook = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBodyData) SetDebug(v string) *GetZidTagScoreByAtokenResponseBodyData {
	s.Debug = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBodyData) SetDoubleOpen(v string) *GetZidTagScoreByAtokenResponseBodyData {
	s.DoubleOpen = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBodyData) SetJavaHook(v string) *GetZidTagScoreByAtokenResponseBodyData {
	s.JavaHook = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBodyData) SetNativeHook(v string) *GetZidTagScoreByAtokenResponseBodyData {
	s.NativeHook = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBodyData) SetRiskLevel(v string) *GetZidTagScoreByAtokenResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBodyData) SetRiskScore(v string) *GetZidTagScoreByAtokenResponseBodyData {
	s.RiskScore = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBodyData) SetRoot(v string) *GetZidTagScoreByAtokenResponseBodyData {
	s.Root = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBodyData) SetSimulator(v string) *GetZidTagScoreByAtokenResponseBodyData {
	s.Simulator = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBodyData) SetVpnProxy(v string) *GetZidTagScoreByAtokenResponseBodyData {
	s.VpnProxy = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBodyData) SetWifiProxy(v string) *GetZidTagScoreByAtokenResponseBodyData {
	s.WifiProxy = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponseBodyData) SetZid(v string) *GetZidTagScoreByAtokenResponseBodyData {
	s.Zid = &v
	return s
}

type GetZidTagScoreByAtokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetZidTagScoreByAtokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetZidTagScoreByAtokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetZidTagScoreByAtokenResponse) GoString() string {
	return s.String()
}

func (s *GetZidTagScoreByAtokenResponse) SetHeaders(v map[string]*string) *GetZidTagScoreByAtokenResponse {
	s.Headers = v
	return s
}

func (s *GetZidTagScoreByAtokenResponse) SetStatusCode(v int32) *GetZidTagScoreByAtokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetZidTagScoreByAtokenResponse) SetBody(v *GetZidTagScoreByAtokenResponseBody) *GetZidTagScoreByAtokenResponse {
	s.Body = v
	return s
}

type ListChannelRiskDetailsRequest struct {
	// example:
	//
	// 360market
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 59892ebcaed179694b000104
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20230418
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 1
	IsAllChannel *string `json:"isAllChannel,omitempty" xml:"isAllChannel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20230418
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ListChannelRiskDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChannelRiskDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListChannelRiskDetailsRequest) SetChannel(v string) *ListChannelRiskDetailsRequest {
	s.Channel = &v
	return s
}

func (s *ListChannelRiskDetailsRequest) SetDataSourceId(v string) *ListChannelRiskDetailsRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListChannelRiskDetailsRequest) SetEnd(v string) *ListChannelRiskDetailsRequest {
	s.End = &v
	return s
}

func (s *ListChannelRiskDetailsRequest) SetIsAllChannel(v string) *ListChannelRiskDetailsRequest {
	s.IsAllChannel = &v
	return s
}

func (s *ListChannelRiskDetailsRequest) SetStart(v string) *ListChannelRiskDetailsRequest {
	s.Start = &v
	return s
}

type ListChannelRiskDetailsResponseBody struct {
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 200
	Code *int64                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListChannelRiskDetailsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 2343535353DF8984565
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListChannelRiskDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChannelRiskDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListChannelRiskDetailsResponseBody) SetMsg(v string) *ListChannelRiskDetailsResponseBody {
	s.Msg = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBody) SetCode(v int64) *ListChannelRiskDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBody) SetData(v *ListChannelRiskDetailsResponseBodyData) *ListChannelRiskDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ListChannelRiskDetailsResponseBody) SetRequestId(v string) *ListChannelRiskDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBody) SetSuccess(v bool) *ListChannelRiskDetailsResponseBody {
	s.Success = &v
	return s
}

type ListChannelRiskDetailsResponseBodyData struct {
	RiskDetails []*ListChannelRiskDetailsResponseBodyDataRiskDetails `json:"riskDetails,omitempty" xml:"riskDetails,omitempty" type:"Repeated"`
	RiskSumary  []*ListChannelRiskDetailsResponseBodyDataRiskSumary  `json:"riskSumary,omitempty" xml:"riskSumary,omitempty" type:"Repeated"`
}

func (s ListChannelRiskDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListChannelRiskDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChannelRiskDetailsResponseBodyData) SetRiskDetails(v []*ListChannelRiskDetailsResponseBodyDataRiskDetails) *ListChannelRiskDetailsResponseBodyData {
	s.RiskDetails = v
	return s
}

func (s *ListChannelRiskDetailsResponseBodyData) SetRiskSumary(v []*ListChannelRiskDetailsResponseBodyDataRiskSumary) *ListChannelRiskDetailsResponseBodyData {
	s.RiskSumary = v
	return s
}

type ListChannelRiskDetailsResponseBodyDataRiskDetails struct {
	// example:
	//
	// 0
	An *string `json:"an,omitempty" xml:"an,omitempty"`
	// example:
	//
	// v1.1.0
	Av *string `json:"av,omitempty" xml:"av,omitempty"`
	// example:
	//
	// 1
	Bn *string `json:"bn,omitempty" xml:"bn,omitempty"`
	// example:
	//
	// 360market
	C *string `json:"c,omitempty" xml:"c,omitempty"`
	// example:
	//
	// 20230410
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// 1
	Fd *string `json:"fd,omitempty" xml:"fd,omitempty"`
	// example:
	//
	// 0000-0000-0000000
	Idfa *string `json:"idfa,omitempty" xml:"idfa,omitempty"`
	// example:
	//
	// 0
	Jb *string `json:"jb,omitempty" xml:"jb,omitempty"`
	// example:
	//
	// 343rUFD834343KJDDFS
	Oaid *string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// example:
	//
	// 1
	Py *string `json:"py,omitempty" xml:"py,omitempty"`
	// example:
	//
	// RISK
	Rl *string `json:"rl,omitempty" xml:"rl,omitempty"`
	// example:
	//
	// 60
	Rs *string `json:"rs,omitempty" xml:"rs,omitempty"`
	// example:
	//
	// 815E045B-4A9E-AA76-BEFF-048C9B9F651A
	Zid *string `json:"zid,omitempty" xml:"zid,omitempty"`
}

func (s ListChannelRiskDetailsResponseBodyDataRiskDetails) String() string {
	return tea.Prettify(s)
}

func (s ListChannelRiskDetailsResponseBodyDataRiskDetails) GoString() string {
	return s.String()
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskDetails) SetAn(v string) *ListChannelRiskDetailsResponseBodyDataRiskDetails {
	s.An = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskDetails) SetAv(v string) *ListChannelRiskDetailsResponseBodyDataRiskDetails {
	s.Av = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskDetails) SetBn(v string) *ListChannelRiskDetailsResponseBodyDataRiskDetails {
	s.Bn = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskDetails) SetC(v string) *ListChannelRiskDetailsResponseBodyDataRiskDetails {
	s.C = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskDetails) SetDate(v string) *ListChannelRiskDetailsResponseBodyDataRiskDetails {
	s.Date = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskDetails) SetFd(v string) *ListChannelRiskDetailsResponseBodyDataRiskDetails {
	s.Fd = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskDetails) SetIdfa(v string) *ListChannelRiskDetailsResponseBodyDataRiskDetails {
	s.Idfa = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskDetails) SetJb(v string) *ListChannelRiskDetailsResponseBodyDataRiskDetails {
	s.Jb = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskDetails) SetOaid(v string) *ListChannelRiskDetailsResponseBodyDataRiskDetails {
	s.Oaid = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskDetails) SetPy(v string) *ListChannelRiskDetailsResponseBodyDataRiskDetails {
	s.Py = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskDetails) SetRl(v string) *ListChannelRiskDetailsResponseBodyDataRiskDetails {
	s.Rl = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskDetails) SetRs(v string) *ListChannelRiskDetailsResponseBodyDataRiskDetails {
	s.Rs = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskDetails) SetZid(v string) *ListChannelRiskDetailsResponseBodyDataRiskDetails {
	s.Zid = &v
	return s
}

type ListChannelRiskDetailsResponseBodyDataRiskSumary struct {
	// example:
	//
	// 20230418
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// 2
	RiskZidEmuDistinctNew *string `json:"riskZidEmuDistinctNew,omitempty" xml:"riskZidEmuDistinctNew,omitempty"`
}

func (s ListChannelRiskDetailsResponseBodyDataRiskSumary) String() string {
	return tea.Prettify(s)
}

func (s ListChannelRiskDetailsResponseBodyDataRiskSumary) GoString() string {
	return s.String()
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskSumary) SetDate(v string) *ListChannelRiskDetailsResponseBodyDataRiskSumary {
	s.Date = &v
	return s
}

func (s *ListChannelRiskDetailsResponseBodyDataRiskSumary) SetRiskZidEmuDistinctNew(v string) *ListChannelRiskDetailsResponseBodyDataRiskSumary {
	s.RiskZidEmuDistinctNew = &v
	return s
}

type ListChannelRiskDetailsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChannelRiskDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChannelRiskDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChannelRiskDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListChannelRiskDetailsResponse) SetHeaders(v map[string]*string) *ListChannelRiskDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListChannelRiskDetailsResponse) SetStatusCode(v int32) *ListChannelRiskDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChannelRiskDetailsResponse) SetBody(v *ListChannelRiskDetailsResponseBody) *ListChannelRiskDetailsResponse {
	s.Body = v
	return s
}

type ListUninstallDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 62c792d188ccdf4b7ec1aa38
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20231120
	EndDs *string `json:"endDs,omitempty" xml:"endDs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20231120
	StartDs *string `json:"startDs,omitempty" xml:"startDs,omitempty"`
}

func (s ListUninstallDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUninstallDetailRequest) GoString() string {
	return s.String()
}

func (s *ListUninstallDetailRequest) SetDataSourceId(v string) *ListUninstallDetailRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListUninstallDetailRequest) SetEndDs(v string) *ListUninstallDetailRequest {
	s.EndDs = &v
	return s
}

func (s *ListUninstallDetailRequest) SetStartDs(v string) *ListUninstallDetailRequest {
	s.StartDs = &v
	return s
}

type ListUninstallDetailResponseBody struct {
	// msg
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// code
	//
	// example:
	//
	// 200
	Code *int64                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListUninstallDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 8797ddc1857a66deb6b5d14
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListUninstallDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUninstallDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListUninstallDetailResponseBody) SetMsg(v string) *ListUninstallDetailResponseBody {
	s.Msg = &v
	return s
}

func (s *ListUninstallDetailResponseBody) SetSuccess(v bool) *ListUninstallDetailResponseBody {
	s.Success = &v
	return s
}

func (s *ListUninstallDetailResponseBody) SetCode(v int64) *ListUninstallDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListUninstallDetailResponseBody) SetData(v *ListUninstallDetailResponseBodyData) *ListUninstallDetailResponseBody {
	s.Data = v
	return s
}

func (s *ListUninstallDetailResponseBody) SetRequestId(v string) *ListUninstallDetailResponseBody {
	s.RequestId = &v
	return s
}

type ListUninstallDetailResponseBodyData struct {
	Details []*ListUninstallDetailResponseBodyDataDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
}

func (s ListUninstallDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUninstallDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUninstallDetailResponseBodyData) SetDetails(v []*ListUninstallDetailResponseBodyDataDetails) *ListUninstallDetailResponseBodyData {
	s.Details = v
	return s
}

type ListUninstallDetailResponseBodyDataDetails struct {
	// example:
	//
	// 1700278078773
	ActiveDatetime *string `json:"activeDatetime,omitempty" xml:"activeDatetime,omitempty"`
	// example:
	//
	// 阳江市
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// deviceBrand
	//
	// example:
	//
	// HONOR
	DeviceBrand *string `json:"deviceBrand,omitempty" xml:"deviceBrand,omitempty"`
	// deviceModel
	//
	// example:
	//
	// KOZ-AL00
	DeviceModel *string `json:"deviceModel,omitempty" xml:"deviceModel,omitempty"`
	// firstActiveDatetime
	//
	// example:
	//
	// 2023-11-12 22:41:46
	FirstActiveDatetime *string `json:"firstActiveDatetime,omitempty" xml:"firstActiveDatetime,omitempty"`
	// idfa
	//
	// example:
	//
	// 5ce253a8-2917-44da-a1fe-9d21051f9015
	Idfa *string `json:"idfa,omitempty" xml:"idfa,omitempty"`
	// imei
	//
	// example:
	//
	// None
	Imei *string `json:"imei,omitempty" xml:"imei,omitempty"`
	// installAppVersion
	//
	// example:
	//
	// 23.05.06
	InstallAppVersion *string `json:"installAppVersion,omitempty" xml:"installAppVersion,omitempty"`
	// installChannel
	//
	// example:
	//
	// umeng
	InstallChannel *string `json:"installChannel,omitempty" xml:"installChannel,omitempty"`
	// oaid
	//
	// example:
	//
	// ai357e2b5358797ddc
	Oaid *string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// osVersion
	//
	// example:
	//
	// 10
	OsVersion *string `json:"osVersion,omitempty" xml:"osVersion,omitempty"`
	// puid
	//
	// example:
	//
	// zhangfei
	Puid *string `json:"puid,omitempty" xml:"puid,omitempty"`
	// example:
	//
	// ai357e2b5358797ddc1857a66deb6b5d14
	Umid *string `json:"umid,omitempty" xml:"umid,omitempty"`
	// example:
	//
	// 3
	UninstallCount *int32 `json:"uninstallCount,omitempty" xml:"uninstallCount,omitempty"`
	// example:
	//
	// 20231118
	UninstallDatetime *string `json:"uninstallDatetime,omitempty" xml:"uninstallDatetime,omitempty"`
	// zid
	//
	// example:
	//
	// ddc1857a66deb6b5
	Zid *string `json:"zid,omitempty" xml:"zid,omitempty"`
}

func (s ListUninstallDetailResponseBodyDataDetails) String() string {
	return tea.Prettify(s)
}

func (s ListUninstallDetailResponseBodyDataDetails) GoString() string {
	return s.String()
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetActiveDatetime(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.ActiveDatetime = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetCity(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.City = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetDeviceBrand(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.DeviceBrand = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetDeviceModel(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.DeviceModel = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetFirstActiveDatetime(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.FirstActiveDatetime = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetIdfa(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.Idfa = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetImei(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.Imei = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetInstallAppVersion(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.InstallAppVersion = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetInstallChannel(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.InstallChannel = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetOaid(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.Oaid = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetOsVersion(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.OsVersion = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetPuid(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.Puid = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetUmid(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.Umid = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetUninstallCount(v int32) *ListUninstallDetailResponseBodyDataDetails {
	s.UninstallCount = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetUninstallDatetime(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.UninstallDatetime = &v
	return s
}

func (s *ListUninstallDetailResponseBodyDataDetails) SetZid(v string) *ListUninstallDetailResponseBodyDataDetails {
	s.Zid = &v
	return s
}

type ListUninstallDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUninstallDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUninstallDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUninstallDetailResponse) GoString() string {
	return s.String()
}

func (s *ListUninstallDetailResponse) SetHeaders(v map[string]*string) *ListUninstallDetailResponse {
	s.Headers = v
	return s
}

func (s *ListUninstallDetailResponse) SetStatusCode(v int32) *ListUninstallDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUninstallDetailResponse) SetBody(v *ListUninstallDetailResponseBody) *ListUninstallDetailResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("antirisk"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 获取实时反作弊信息
//
// @param request - GetRealTimeRiskInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRealTimeRiskInfoResponse
func (client *Client) GetRealTimeRiskInfoWithOptions(request *GetRealTimeRiskInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRealTimeRiskInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Atoken)) {
		query["atoken"] = request.Atoken
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		query["extra"] = request.Extra
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRealTimeRiskInfo"),
		Version:     tea.String("2022-11-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/anti/getRealTimeRiskInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRealTimeRiskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实时反作弊信息
//
// @param request - GetRealTimeRiskInfoRequest
//
// @return GetRealTimeRiskInfoResponse
func (client *Client) GetRealTimeRiskInfo(request *GetRealTimeRiskInfoRequest) (_result *GetRealTimeRiskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRealTimeRiskInfoResponse{}
	_body, _err := client.GetRealTimeRiskInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// atoken换zid+tags
//
// @param request - GetZidTagByAtokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetZidTagByAtokenResponse
func (client *Client) GetZidTagByAtokenWithOptions(request *GetZidTagByAtokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetZidTagByAtokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Atoken)) {
		query["atoken"] = request.Atoken
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetZidTagByAtoken"),
		Version:     tea.String("2022-11-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/anti/getZidTagByAtoken"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetZidTagByAtokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// atoken换zid+tags
//
// @param request - GetZidTagByAtokenRequest
//
// @return GetZidTagByAtokenResponse
func (client *Client) GetZidTagByAtoken(request *GetZidTagByAtokenRequest) (_result *GetZidTagByAtokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetZidTagByAtokenResponse{}
	_body, _err := client.GetZidTagByAtokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// atoken换zid+tags+风险分
//
// @param request - GetZidTagScoreByAtokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetZidTagScoreByAtokenResponse
func (client *Client) GetZidTagScoreByAtokenWithOptions(request *GetZidTagScoreByAtokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetZidTagScoreByAtokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Atoken)) {
		query["atoken"] = request.Atoken
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetZidTagScoreByAtoken"),
		Version:     tea.String("2022-11-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/anti/getZidTagScoreByAtoken"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetZidTagScoreByAtokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// atoken换zid+tags+风险分
//
// @param request - GetZidTagScoreByAtokenRequest
//
// @return GetZidTagScoreByAtokenResponse
func (client *Client) GetZidTagScoreByAtoken(request *GetZidTagScoreByAtokenRequest) (_result *GetZidTagScoreByAtokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetZidTagScoreByAtokenResponse{}
	_body, _err := client.GetZidTagScoreByAtokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 渠道风险明细
//
// @param request - ListChannelRiskDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChannelRiskDetailsResponse
func (client *Client) ListChannelRiskDetailsWithOptions(request *ListChannelRiskDetailsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListChannelRiskDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		query["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.IsAllChannel)) {
		query["isAllChannel"] = request.IsAllChannel
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListChannelRiskDetails"),
		Version:     tea.String("2022-11-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/anti/listChannelRiskDetails"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListChannelRiskDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 渠道风险明细
//
// @param request - ListChannelRiskDetailsRequest
//
// @return ListChannelRiskDetailsResponse
func (client *Client) ListChannelRiskDetails(request *ListChannelRiskDetailsRequest) (_result *ListChannelRiskDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListChannelRiskDetailsResponse{}
	_body, _err := client.ListChannelRiskDetailsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 卸载明细列表
//
// @param request - ListUninstallDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUninstallDetailResponse
func (client *Client) ListUninstallDetailWithOptions(request *ListUninstallDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUninstallDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDs)) {
		query["endDs"] = request.EndDs
	}

	if !tea.BoolValue(util.IsUnset(request.StartDs)) {
		query["startDs"] = request.StartDs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUninstallDetail"),
		Version:     tea.String("2022-11-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/uninstall/listUninstallDetail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUninstallDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 卸载明细列表
//
// @param request - ListUninstallDetailRequest
//
// @return ListUninstallDetailResponse
func (client *Client) ListUninstallDetail(request *ListUninstallDetailRequest) (_result *ListUninstallDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUninstallDetailResponse{}
	_body, _err := client.ListUninstallDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
