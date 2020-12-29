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

type AddCustomLineRequest struct {
	Lang         *string                          `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string                          `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string                          `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	LineName     *string                          `json:"LineName,omitempty" xml:"LineName,omitempty"`
	IpSegment    []*AddCustomLineRequestIpSegment `json:"IpSegment,omitempty" xml:"IpSegment,omitempty" type:"Repeated"`
}

func (s AddCustomLineRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCustomLineRequest) GoString() string {
	return s.String()
}

func (s *AddCustomLineRequest) SetLang(v string) *AddCustomLineRequest {
	s.Lang = &v
	return s
}

func (s *AddCustomLineRequest) SetUserClientIp(v string) *AddCustomLineRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddCustomLineRequest) SetDomainName(v string) *AddCustomLineRequest {
	s.DomainName = &v
	return s
}

func (s *AddCustomLineRequest) SetLineName(v string) *AddCustomLineRequest {
	s.LineName = &v
	return s
}

func (s *AddCustomLineRequest) SetIpSegment(v []*AddCustomLineRequestIpSegment) *AddCustomLineRequest {
	s.IpSegment = v
	return s
}

type AddCustomLineRequestIpSegment struct {
	EndIp   *string `json:"EndIp,omitempty" xml:"EndIp,omitempty"`
	StartIp *string `json:"StartIp,omitempty" xml:"StartIp,omitempty"`
}

func (s AddCustomLineRequestIpSegment) String() string {
	return tea.Prettify(s)
}

func (s AddCustomLineRequestIpSegment) GoString() string {
	return s.String()
}

func (s *AddCustomLineRequestIpSegment) SetEndIp(v string) *AddCustomLineRequestIpSegment {
	s.EndIp = &v
	return s
}

func (s *AddCustomLineRequestIpSegment) SetStartIp(v string) *AddCustomLineRequestIpSegment {
	s.StartIp = &v
	return s
}

type AddCustomLineResponseBody struct {
	LineId    *int64  `json:"LineId,omitempty" xml:"LineId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LineCode  *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
}

func (s AddCustomLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCustomLineResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomLineResponseBody) SetLineId(v int64) *AddCustomLineResponseBody {
	s.LineId = &v
	return s
}

func (s *AddCustomLineResponseBody) SetRequestId(v string) *AddCustomLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomLineResponseBody) SetLineCode(v string) *AddCustomLineResponseBody {
	s.LineCode = &v
	return s
}

type AddCustomLineResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddCustomLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddCustomLineResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCustomLineResponse) GoString() string {
	return s.String()
}

func (s *AddCustomLineResponse) SetHeaders(v map[string]*string) *AddCustomLineResponse {
	s.Headers = v
	return s
}

func (s *AddCustomLineResponse) SetBody(v *AddCustomLineResponseBody) *AddCustomLineResponse {
	s.Body = v
	return s
}

type AddDnsGtmAccessStrategyRequest struct {
	Lang                        *string                                           `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp                *string                                           `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId                  *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StrategyName                *string                                           `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	Lines                       *string                                           `json:"Lines,omitempty" xml:"Lines,omitempty"`
	DefaultAddrPoolType         *string                                           `json:"DefaultAddrPoolType,omitempty" xml:"DefaultAddrPoolType,omitempty"`
	DefaultLbaStrategy          *string                                           `json:"DefaultLbaStrategy,omitempty" xml:"DefaultLbaStrategy,omitempty"`
	DefaultMinAvailableAddrNum  *int32                                            `json:"DefaultMinAvailableAddrNum,omitempty" xml:"DefaultMinAvailableAddrNum,omitempty"`
	DefaultMaxReturnAddrNum     *int32                                            `json:"DefaultMaxReturnAddrNum,omitempty" xml:"DefaultMaxReturnAddrNum,omitempty"`
	DefaultLatencyOptimization  *string                                           `json:"DefaultLatencyOptimization,omitempty" xml:"DefaultLatencyOptimization,omitempty"`
	FailoverAddrPoolType        *string                                           `json:"FailoverAddrPoolType,omitempty" xml:"FailoverAddrPoolType,omitempty"`
	FailoverLbaStrategy         *string                                           `json:"FailoverLbaStrategy,omitempty" xml:"FailoverLbaStrategy,omitempty"`
	FailoverMinAvailableAddrNum *int32                                            `json:"FailoverMinAvailableAddrNum,omitempty" xml:"FailoverMinAvailableAddrNum,omitempty"`
	FailoverMaxReturnAddrNum    *int32                                            `json:"FailoverMaxReturnAddrNum,omitempty" xml:"FailoverMaxReturnAddrNum,omitempty"`
	FailoverLatencyOptimization *string                                           `json:"FailoverLatencyOptimization,omitempty" xml:"FailoverLatencyOptimization,omitempty"`
	StrategyMode                *string                                           `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
	DefaultAddrPool             []*AddDnsGtmAccessStrategyRequestDefaultAddrPool  `json:"DefaultAddrPool,omitempty" xml:"DefaultAddrPool,omitempty" type:"Repeated"`
	FailoverAddrPool            []*AddDnsGtmAccessStrategyRequestFailoverAddrPool `json:"FailoverAddrPool,omitempty" xml:"FailoverAddrPool,omitempty" type:"Repeated"`
}

func (s AddDnsGtmAccessStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDnsGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAccessStrategyRequest) SetLang(v string) *AddDnsGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetUserClientIp(v string) *AddDnsGtmAccessStrategyRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetInstanceId(v string) *AddDnsGtmAccessStrategyRequest {
	s.InstanceId = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetStrategyName(v string) *AddDnsGtmAccessStrategyRequest {
	s.StrategyName = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetLines(v string) *AddDnsGtmAccessStrategyRequest {
	s.Lines = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetDefaultAddrPoolType(v string) *AddDnsGtmAccessStrategyRequest {
	s.DefaultAddrPoolType = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetDefaultLbaStrategy(v string) *AddDnsGtmAccessStrategyRequest {
	s.DefaultLbaStrategy = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetDefaultMinAvailableAddrNum(v int32) *AddDnsGtmAccessStrategyRequest {
	s.DefaultMinAvailableAddrNum = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetDefaultMaxReturnAddrNum(v int32) *AddDnsGtmAccessStrategyRequest {
	s.DefaultMaxReturnAddrNum = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetDefaultLatencyOptimization(v string) *AddDnsGtmAccessStrategyRequest {
	s.DefaultLatencyOptimization = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetFailoverAddrPoolType(v string) *AddDnsGtmAccessStrategyRequest {
	s.FailoverAddrPoolType = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetFailoverLbaStrategy(v string) *AddDnsGtmAccessStrategyRequest {
	s.FailoverLbaStrategy = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetFailoverMinAvailableAddrNum(v int32) *AddDnsGtmAccessStrategyRequest {
	s.FailoverMinAvailableAddrNum = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetFailoverMaxReturnAddrNum(v int32) *AddDnsGtmAccessStrategyRequest {
	s.FailoverMaxReturnAddrNum = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetFailoverLatencyOptimization(v string) *AddDnsGtmAccessStrategyRequest {
	s.FailoverLatencyOptimization = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetStrategyMode(v string) *AddDnsGtmAccessStrategyRequest {
	s.StrategyMode = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetDefaultAddrPool(v []*AddDnsGtmAccessStrategyRequestDefaultAddrPool) *AddDnsGtmAccessStrategyRequest {
	s.DefaultAddrPool = v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetFailoverAddrPool(v []*AddDnsGtmAccessStrategyRequestFailoverAddrPool) *AddDnsGtmAccessStrategyRequest {
	s.FailoverAddrPool = v
	return s
}

type AddDnsGtmAccessStrategyRequestDefaultAddrPool struct {
	LbaWeight *int32  `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AddDnsGtmAccessStrategyRequestDefaultAddrPool) String() string {
	return tea.Prettify(s)
}

func (s AddDnsGtmAccessStrategyRequestDefaultAddrPool) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAccessStrategyRequestDefaultAddrPool) SetLbaWeight(v int32) *AddDnsGtmAccessStrategyRequestDefaultAddrPool {
	s.LbaWeight = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequestDefaultAddrPool) SetId(v string) *AddDnsGtmAccessStrategyRequestDefaultAddrPool {
	s.Id = &v
	return s
}

type AddDnsGtmAccessStrategyRequestFailoverAddrPool struct {
	LbaWeight *int32  `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AddDnsGtmAccessStrategyRequestFailoverAddrPool) String() string {
	return tea.Prettify(s)
}

func (s AddDnsGtmAccessStrategyRequestFailoverAddrPool) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAccessStrategyRequestFailoverAddrPool) SetLbaWeight(v int32) *AddDnsGtmAccessStrategyRequestFailoverAddrPool {
	s.LbaWeight = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequestFailoverAddrPool) SetId(v string) *AddDnsGtmAccessStrategyRequestFailoverAddrPool {
	s.Id = &v
	return s
}

type AddDnsGtmAccessStrategyResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s AddDnsGtmAccessStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDnsGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAccessStrategyResponseBody) SetRequestId(v string) *AddDnsGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDnsGtmAccessStrategyResponseBody) SetStrategyId(v string) *AddDnsGtmAccessStrategyResponseBody {
	s.StrategyId = &v
	return s
}

type AddDnsGtmAccessStrategyResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDnsGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDnsGtmAccessStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDnsGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *AddDnsGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *AddDnsGtmAccessStrategyResponse) SetBody(v *AddDnsGtmAccessStrategyResponseBody) *AddDnsGtmAccessStrategyResponse {
	s.Body = v
	return s
}

type AddDnsGtmAddressPoolRequest struct {
	UserClientIp      *string                                   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang              *string                                   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId        *string                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name              *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Type              *string                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	LbaStrategy       *string                                   `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	MonitorStatus     *string                                   `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	ProtocolType      *string                                   `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Interval          *int32                                    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	EvaluationCount   *int32                                    `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Timeout           *int32                                    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	MonitorExtendInfo *string                                   `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
	Addr              []*AddDnsGtmAddressPoolRequestAddr        `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Repeated"`
	IspCityNode       []*AddDnsGtmAddressPoolRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
}

func (s AddDnsGtmAddressPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDnsGtmAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAddressPoolRequest) SetUserClientIp(v string) *AddDnsGtmAddressPoolRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetLang(v string) *AddDnsGtmAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetInstanceId(v string) *AddDnsGtmAddressPoolRequest {
	s.InstanceId = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetName(v string) *AddDnsGtmAddressPoolRequest {
	s.Name = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetType(v string) *AddDnsGtmAddressPoolRequest {
	s.Type = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetLbaStrategy(v string) *AddDnsGtmAddressPoolRequest {
	s.LbaStrategy = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetMonitorStatus(v string) *AddDnsGtmAddressPoolRequest {
	s.MonitorStatus = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetProtocolType(v string) *AddDnsGtmAddressPoolRequest {
	s.ProtocolType = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetInterval(v int32) *AddDnsGtmAddressPoolRequest {
	s.Interval = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetEvaluationCount(v int32) *AddDnsGtmAddressPoolRequest {
	s.EvaluationCount = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetTimeout(v int32) *AddDnsGtmAddressPoolRequest {
	s.Timeout = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetMonitorExtendInfo(v string) *AddDnsGtmAddressPoolRequest {
	s.MonitorExtendInfo = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetAddr(v []*AddDnsGtmAddressPoolRequestAddr) *AddDnsGtmAddressPoolRequest {
	s.Addr = v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetIspCityNode(v []*AddDnsGtmAddressPoolRequestIspCityNode) *AddDnsGtmAddressPoolRequest {
	s.IspCityNode = v
	return s
}

type AddDnsGtmAddressPoolRequestAddr struct {
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	LbaWeight     *int32  `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	Addr          *string `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Mode          *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s AddDnsGtmAddressPoolRequestAddr) String() string {
	return tea.Prettify(s)
}

func (s AddDnsGtmAddressPoolRequestAddr) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAddressPoolRequestAddr) SetAttributeInfo(v string) *AddDnsGtmAddressPoolRequestAddr {
	s.AttributeInfo = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequestAddr) SetRemark(v string) *AddDnsGtmAddressPoolRequestAddr {
	s.Remark = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequestAddr) SetLbaWeight(v int32) *AddDnsGtmAddressPoolRequestAddr {
	s.LbaWeight = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequestAddr) SetAddr(v string) *AddDnsGtmAddressPoolRequestAddr {
	s.Addr = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequestAddr) SetMode(v string) *AddDnsGtmAddressPoolRequestAddr {
	s.Mode = &v
	return s
}

type AddDnsGtmAddressPoolRequestIspCityNode struct {
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	IspCode  *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
}

func (s AddDnsGtmAddressPoolRequestIspCityNode) String() string {
	return tea.Prettify(s)
}

func (s AddDnsGtmAddressPoolRequestIspCityNode) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAddressPoolRequestIspCityNode) SetCityCode(v string) *AddDnsGtmAddressPoolRequestIspCityNode {
	s.CityCode = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequestIspCityNode) SetIspCode(v string) *AddDnsGtmAddressPoolRequestIspCityNode {
	s.IspCode = &v
	return s
}

type AddDnsGtmAddressPoolResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AddrPoolId      *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
}

func (s AddDnsGtmAddressPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDnsGtmAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAddressPoolResponseBody) SetRequestId(v string) *AddDnsGtmAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDnsGtmAddressPoolResponseBody) SetAddrPoolId(v string) *AddDnsGtmAddressPoolResponseBody {
	s.AddrPoolId = &v
	return s
}

func (s *AddDnsGtmAddressPoolResponseBody) SetMonitorConfigId(v string) *AddDnsGtmAddressPoolResponseBody {
	s.MonitorConfigId = &v
	return s
}

type AddDnsGtmAddressPoolResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDnsGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDnsGtmAddressPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDnsGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAddressPoolResponse) SetHeaders(v map[string]*string) *AddDnsGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *AddDnsGtmAddressPoolResponse) SetBody(v *AddDnsGtmAddressPoolResponseBody) *AddDnsGtmAddressPoolResponse {
	s.Body = v
	return s
}

type AddDnsGtmMonitorRequest struct {
	UserClientIp      *string                               `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang              *string                               `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AddrPoolId        *string                               `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	ProtocolType      *string                               `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Interval          *int32                                `json:"Interval,omitempty" xml:"Interval,omitempty"`
	EvaluationCount   *int32                                `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Timeout           *int32                                `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	MonitorExtendInfo *string                               `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
	IspCityNode       []*AddDnsGtmMonitorRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
}

func (s AddDnsGtmMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDnsGtmMonitorRequest) GoString() string {
	return s.String()
}

func (s *AddDnsGtmMonitorRequest) SetUserClientIp(v string) *AddDnsGtmMonitorRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetLang(v string) *AddDnsGtmMonitorRequest {
	s.Lang = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetAddrPoolId(v string) *AddDnsGtmMonitorRequest {
	s.AddrPoolId = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetProtocolType(v string) *AddDnsGtmMonitorRequest {
	s.ProtocolType = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetInterval(v int32) *AddDnsGtmMonitorRequest {
	s.Interval = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetEvaluationCount(v int32) *AddDnsGtmMonitorRequest {
	s.EvaluationCount = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetTimeout(v int32) *AddDnsGtmMonitorRequest {
	s.Timeout = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetMonitorExtendInfo(v string) *AddDnsGtmMonitorRequest {
	s.MonitorExtendInfo = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetIspCityNode(v []*AddDnsGtmMonitorRequestIspCityNode) *AddDnsGtmMonitorRequest {
	s.IspCityNode = v
	return s
}

type AddDnsGtmMonitorRequestIspCityNode struct {
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	IspCode  *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
}

func (s AddDnsGtmMonitorRequestIspCityNode) String() string {
	return tea.Prettify(s)
}

func (s AddDnsGtmMonitorRequestIspCityNode) GoString() string {
	return s.String()
}

func (s *AddDnsGtmMonitorRequestIspCityNode) SetCityCode(v string) *AddDnsGtmMonitorRequestIspCityNode {
	s.CityCode = &v
	return s
}

func (s *AddDnsGtmMonitorRequestIspCityNode) SetIspCode(v string) *AddDnsGtmMonitorRequestIspCityNode {
	s.IspCode = &v
	return s
}

type AddDnsGtmMonitorResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
}

func (s AddDnsGtmMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDnsGtmMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *AddDnsGtmMonitorResponseBody) SetRequestId(v string) *AddDnsGtmMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDnsGtmMonitorResponseBody) SetMonitorConfigId(v string) *AddDnsGtmMonitorResponseBody {
	s.MonitorConfigId = &v
	return s
}

type AddDnsGtmMonitorResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDnsGtmMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDnsGtmMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDnsGtmMonitorResponse) GoString() string {
	return s.String()
}

func (s *AddDnsGtmMonitorResponse) SetHeaders(v map[string]*string) *AddDnsGtmMonitorResponse {
	s.Headers = v
	return s
}

func (s *AddDnsGtmMonitorResponse) SetBody(v *AddDnsGtmMonitorResponseBody) *AddDnsGtmMonitorResponse {
	s.Body = v
	return s
}

type AddDomainRequest struct {
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	UserClientIp    *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s AddDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDomainRequest) GoString() string {
	return s.String()
}

func (s *AddDomainRequest) SetLang(v string) *AddDomainRequest {
	s.Lang = &v
	return s
}

func (s *AddDomainRequest) SetDomainName(v string) *AddDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddDomainRequest) SetGroupId(v string) *AddDomainRequest {
	s.GroupId = &v
	return s
}

func (s *AddDomainRequest) SetResourceGroupId(v string) *AddDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddDomainRequest) SetUserClientIp(v string) *AddDomainRequest {
	s.UserClientIp = &v
	return s
}

type AddDomainResponseBody struct {
	GroupName  *string   `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	DomainId   *string   `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	PunyCode   *string   `json:"PunyCode,omitempty" xml:"PunyCode,omitempty"`
	DnsServers []*string `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Repeated"`
	GroupId    *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s AddDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddDomainResponseBody) SetGroupName(v string) *AddDomainResponseBody {
	s.GroupName = &v
	return s
}

func (s *AddDomainResponseBody) SetDomainId(v string) *AddDomainResponseBody {
	s.DomainId = &v
	return s
}

func (s *AddDomainResponseBody) SetRequestId(v string) *AddDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDomainResponseBody) SetDomainName(v string) *AddDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *AddDomainResponseBody) SetPunyCode(v string) *AddDomainResponseBody {
	s.PunyCode = &v
	return s
}

func (s *AddDomainResponseBody) SetDnsServers(v []*string) *AddDomainResponseBody {
	s.DnsServers = v
	return s
}

func (s *AddDomainResponseBody) SetGroupId(v string) *AddDomainResponseBody {
	s.GroupId = &v
	return s
}

type AddDomainResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDomainResponse) GoString() string {
	return s.String()
}

func (s *AddDomainResponse) SetHeaders(v map[string]*string) *AddDomainResponse {
	s.Headers = v
	return s
}

func (s *AddDomainResponse) SetBody(v *AddDomainResponseBody) *AddDomainResponse {
	s.Body = v
	return s
}

type AddDomainBackupRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	PeriodType   *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s AddDomainBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDomainBackupRequest) GoString() string {
	return s.String()
}

func (s *AddDomainBackupRequest) SetLang(v string) *AddDomainBackupRequest {
	s.Lang = &v
	return s
}

func (s *AddDomainBackupRequest) SetDomainName(v string) *AddDomainBackupRequest {
	s.DomainName = &v
	return s
}

func (s *AddDomainBackupRequest) SetPeriodType(v string) *AddDomainBackupRequest {
	s.PeriodType = &v
	return s
}

func (s *AddDomainBackupRequest) SetUserClientIp(v string) *AddDomainBackupRequest {
	s.UserClientIp = &v
	return s
}

type AddDomainBackupResponseBody struct {
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s AddDomainBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDomainBackupResponseBody) GoString() string {
	return s.String()
}

func (s *AddDomainBackupResponseBody) SetPeriodType(v string) *AddDomainBackupResponseBody {
	s.PeriodType = &v
	return s
}

func (s *AddDomainBackupResponseBody) SetRequestId(v string) *AddDomainBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDomainBackupResponseBody) SetDomainName(v string) *AddDomainBackupResponseBody {
	s.DomainName = &v
	return s
}

type AddDomainBackupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDomainBackupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDomainBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDomainBackupResponse) GoString() string {
	return s.String()
}

func (s *AddDomainBackupResponse) SetHeaders(v map[string]*string) *AddDomainBackupResponse {
	s.Headers = v
	return s
}

func (s *AddDomainBackupResponse) SetBody(v *AddDomainBackupResponseBody) *AddDomainBackupResponse {
	s.Body = v
	return s
}

type AddDomainGroupRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s AddDomainGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDomainGroupRequest) GoString() string {
	return s.String()
}

func (s *AddDomainGroupRequest) SetLang(v string) *AddDomainGroupRequest {
	s.Lang = &v
	return s
}

func (s *AddDomainGroupRequest) SetUserClientIp(v string) *AddDomainGroupRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddDomainGroupRequest) SetGroupName(v string) *AddDomainGroupRequest {
	s.GroupName = &v
	return s
}

type AddDomainGroupResponseBody struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s AddDomainGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDomainGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddDomainGroupResponseBody) SetGroupName(v string) *AddDomainGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *AddDomainGroupResponseBody) SetRequestId(v string) *AddDomainGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDomainGroupResponseBody) SetGroupId(v string) *AddDomainGroupResponseBody {
	s.GroupId = &v
	return s
}

type AddDomainGroupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDomainGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDomainGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDomainGroupResponse) GoString() string {
	return s.String()
}

func (s *AddDomainGroupResponse) SetHeaders(v map[string]*string) *AddDomainGroupResponse {
	s.Headers = v
	return s
}

func (s *AddDomainGroupResponse) SetBody(v *AddDomainGroupResponseBody) *AddDomainGroupResponse {
	s.Body = v
	return s
}

type AddDomainRecordRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RR           *string `json:"RR,omitempty" xml:"RR,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TTL          *int64  `json:"TTL,omitempty" xml:"TTL,omitempty"`
	Priority     *int64  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Line         *string `json:"Line,omitempty" xml:"Line,omitempty"`
}

func (s AddDomainRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDomainRecordRequest) GoString() string {
	return s.String()
}

func (s *AddDomainRecordRequest) SetLang(v string) *AddDomainRecordRequest {
	s.Lang = &v
	return s
}

func (s *AddDomainRecordRequest) SetUserClientIp(v string) *AddDomainRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddDomainRecordRequest) SetDomainName(v string) *AddDomainRecordRequest {
	s.DomainName = &v
	return s
}

func (s *AddDomainRecordRequest) SetRR(v string) *AddDomainRecordRequest {
	s.RR = &v
	return s
}

func (s *AddDomainRecordRequest) SetType(v string) *AddDomainRecordRequest {
	s.Type = &v
	return s
}

func (s *AddDomainRecordRequest) SetValue(v string) *AddDomainRecordRequest {
	s.Value = &v
	return s
}

func (s *AddDomainRecordRequest) SetTTL(v int64) *AddDomainRecordRequest {
	s.TTL = &v
	return s
}

func (s *AddDomainRecordRequest) SetPriority(v int64) *AddDomainRecordRequest {
	s.Priority = &v
	return s
}

func (s *AddDomainRecordRequest) SetLine(v string) *AddDomainRecordRequest {
	s.Line = &v
	return s
}

type AddDomainRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s AddDomainRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDomainRecordResponseBody) GoString() string {
	return s.String()
}

func (s *AddDomainRecordResponseBody) SetRequestId(v string) *AddDomainRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDomainRecordResponseBody) SetRecordId(v string) *AddDomainRecordResponseBody {
	s.RecordId = &v
	return s
}

type AddDomainRecordResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDomainRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDomainRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDomainRecordResponse) GoString() string {
	return s.String()
}

func (s *AddDomainRecordResponse) SetHeaders(v map[string]*string) *AddDomainRecordResponse {
	s.Headers = v
	return s
}

func (s *AddDomainRecordResponse) SetBody(v *AddDomainRecordResponseBody) *AddDomainRecordResponse {
	s.Body = v
	return s
}

type AddGtmAccessStrategyRequest struct {
	Lang               *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp       *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StrategyName       *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	DefaultAddrPoolId  *string `json:"DefaultAddrPoolId,omitempty" xml:"DefaultAddrPoolId,omitempty"`
	FailoverAddrPoolId *string `json:"FailoverAddrPoolId,omitempty" xml:"FailoverAddrPoolId,omitempty"`
	AccessLines        *string `json:"AccessLines,omitempty" xml:"AccessLines,omitempty"`
}

func (s AddGtmAccessStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *AddGtmAccessStrategyRequest) SetLang(v string) *AddGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *AddGtmAccessStrategyRequest) SetUserClientIp(v string) *AddGtmAccessStrategyRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddGtmAccessStrategyRequest) SetInstanceId(v string) *AddGtmAccessStrategyRequest {
	s.InstanceId = &v
	return s
}

func (s *AddGtmAccessStrategyRequest) SetStrategyName(v string) *AddGtmAccessStrategyRequest {
	s.StrategyName = &v
	return s
}

func (s *AddGtmAccessStrategyRequest) SetDefaultAddrPoolId(v string) *AddGtmAccessStrategyRequest {
	s.DefaultAddrPoolId = &v
	return s
}

func (s *AddGtmAccessStrategyRequest) SetFailoverAddrPoolId(v string) *AddGtmAccessStrategyRequest {
	s.FailoverAddrPoolId = &v
	return s
}

func (s *AddGtmAccessStrategyRequest) SetAccessLines(v string) *AddGtmAccessStrategyRequest {
	s.AccessLines = &v
	return s
}

type AddGtmAccessStrategyResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s AddGtmAccessStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *AddGtmAccessStrategyResponseBody) SetRequestId(v string) *AddGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGtmAccessStrategyResponseBody) SetStrategyId(v string) *AddGtmAccessStrategyResponseBody {
	s.StrategyId = &v
	return s
}

type AddGtmAccessStrategyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddGtmAccessStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *AddGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *AddGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *AddGtmAccessStrategyResponse) SetBody(v *AddGtmAccessStrategyResponseBody) *AddGtmAccessStrategyResponse {
	s.Body = v
	return s
}

type AddGtmAddressPoolRequest struct {
	UserClientIp        *string                                `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang                *string                                `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId          *string                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name                *string                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Type                *string                                `json:"Type,omitempty" xml:"Type,omitempty"`
	MinAvailableAddrNum *int32                                 `json:"MinAvailableAddrNum,omitempty" xml:"MinAvailableAddrNum,omitempty"`
	MonitorStatus       *string                                `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	ProtocolType        *string                                `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Interval            *int32                                 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	EvaluationCount     *int32                                 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Timeout             *int32                                 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	MonitorExtendInfo   *string                                `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
	Addr                []*AddGtmAddressPoolRequestAddr        `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Repeated"`
	IspCityNode         []*AddGtmAddressPoolRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
}

func (s AddGtmAddressPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGtmAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *AddGtmAddressPoolRequest) SetUserClientIp(v string) *AddGtmAddressPoolRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetLang(v string) *AddGtmAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetInstanceId(v string) *AddGtmAddressPoolRequest {
	s.InstanceId = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetName(v string) *AddGtmAddressPoolRequest {
	s.Name = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetType(v string) *AddGtmAddressPoolRequest {
	s.Type = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetMinAvailableAddrNum(v int32) *AddGtmAddressPoolRequest {
	s.MinAvailableAddrNum = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetMonitorStatus(v string) *AddGtmAddressPoolRequest {
	s.MonitorStatus = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetProtocolType(v string) *AddGtmAddressPoolRequest {
	s.ProtocolType = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetInterval(v int32) *AddGtmAddressPoolRequest {
	s.Interval = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetEvaluationCount(v int32) *AddGtmAddressPoolRequest {
	s.EvaluationCount = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetTimeout(v int32) *AddGtmAddressPoolRequest {
	s.Timeout = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetMonitorExtendInfo(v string) *AddGtmAddressPoolRequest {
	s.MonitorExtendInfo = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetAddr(v []*AddGtmAddressPoolRequestAddr) *AddGtmAddressPoolRequest {
	s.Addr = v
	return s
}

func (s *AddGtmAddressPoolRequest) SetIspCityNode(v []*AddGtmAddressPoolRequestIspCityNode) *AddGtmAddressPoolRequest {
	s.IspCityNode = v
	return s
}

type AddGtmAddressPoolRequestAddr struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	LbaWeight *int32  `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	Mode      *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s AddGtmAddressPoolRequestAddr) String() string {
	return tea.Prettify(s)
}

func (s AddGtmAddressPoolRequestAddr) GoString() string {
	return s.String()
}

func (s *AddGtmAddressPoolRequestAddr) SetValue(v string) *AddGtmAddressPoolRequestAddr {
	s.Value = &v
	return s
}

func (s *AddGtmAddressPoolRequestAddr) SetLbaWeight(v int32) *AddGtmAddressPoolRequestAddr {
	s.LbaWeight = &v
	return s
}

func (s *AddGtmAddressPoolRequestAddr) SetMode(v string) *AddGtmAddressPoolRequestAddr {
	s.Mode = &v
	return s
}

type AddGtmAddressPoolRequestIspCityNode struct {
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	IspCode  *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
}

func (s AddGtmAddressPoolRequestIspCityNode) String() string {
	return tea.Prettify(s)
}

func (s AddGtmAddressPoolRequestIspCityNode) GoString() string {
	return s.String()
}

func (s *AddGtmAddressPoolRequestIspCityNode) SetCityCode(v string) *AddGtmAddressPoolRequestIspCityNode {
	s.CityCode = &v
	return s
}

func (s *AddGtmAddressPoolRequestIspCityNode) SetIspCode(v string) *AddGtmAddressPoolRequestIspCityNode {
	s.IspCode = &v
	return s
}

type AddGtmAddressPoolResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AddrPoolId      *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
}

func (s AddGtmAddressPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGtmAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *AddGtmAddressPoolResponseBody) SetRequestId(v string) *AddGtmAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGtmAddressPoolResponseBody) SetAddrPoolId(v string) *AddGtmAddressPoolResponseBody {
	s.AddrPoolId = &v
	return s
}

func (s *AddGtmAddressPoolResponseBody) SetMonitorConfigId(v string) *AddGtmAddressPoolResponseBody {
	s.MonitorConfigId = &v
	return s
}

type AddGtmAddressPoolResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddGtmAddressPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *AddGtmAddressPoolResponse) SetHeaders(v map[string]*string) *AddGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *AddGtmAddressPoolResponse) SetBody(v *AddGtmAddressPoolResponseBody) *AddGtmAddressPoolResponse {
	s.Body = v
	return s
}

type AddGtmMonitorRequest struct {
	UserClientIp      *string                            `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang              *string                            `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AddrPoolId        *string                            `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	ProtocolType      *string                            `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Interval          *int32                             `json:"Interval,omitempty" xml:"Interval,omitempty"`
	EvaluationCount   *int32                             `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Timeout           *int32                             `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	MonitorExtendInfo *string                            `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
	IspCityNode       []*AddGtmMonitorRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
}

func (s AddGtmMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGtmMonitorRequest) GoString() string {
	return s.String()
}

func (s *AddGtmMonitorRequest) SetUserClientIp(v string) *AddGtmMonitorRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddGtmMonitorRequest) SetLang(v string) *AddGtmMonitorRequest {
	s.Lang = &v
	return s
}

func (s *AddGtmMonitorRequest) SetAddrPoolId(v string) *AddGtmMonitorRequest {
	s.AddrPoolId = &v
	return s
}

func (s *AddGtmMonitorRequest) SetProtocolType(v string) *AddGtmMonitorRequest {
	s.ProtocolType = &v
	return s
}

func (s *AddGtmMonitorRequest) SetInterval(v int32) *AddGtmMonitorRequest {
	s.Interval = &v
	return s
}

func (s *AddGtmMonitorRequest) SetEvaluationCount(v int32) *AddGtmMonitorRequest {
	s.EvaluationCount = &v
	return s
}

func (s *AddGtmMonitorRequest) SetTimeout(v int32) *AddGtmMonitorRequest {
	s.Timeout = &v
	return s
}

func (s *AddGtmMonitorRequest) SetMonitorExtendInfo(v string) *AddGtmMonitorRequest {
	s.MonitorExtendInfo = &v
	return s
}

func (s *AddGtmMonitorRequest) SetIspCityNode(v []*AddGtmMonitorRequestIspCityNode) *AddGtmMonitorRequest {
	s.IspCityNode = v
	return s
}

type AddGtmMonitorRequestIspCityNode struct {
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	IspCode  *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
}

func (s AddGtmMonitorRequestIspCityNode) String() string {
	return tea.Prettify(s)
}

func (s AddGtmMonitorRequestIspCityNode) GoString() string {
	return s.String()
}

func (s *AddGtmMonitorRequestIspCityNode) SetCityCode(v string) *AddGtmMonitorRequestIspCityNode {
	s.CityCode = &v
	return s
}

func (s *AddGtmMonitorRequestIspCityNode) SetIspCode(v string) *AddGtmMonitorRequestIspCityNode {
	s.IspCode = &v
	return s
}

type AddGtmMonitorResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
}

func (s AddGtmMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGtmMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *AddGtmMonitorResponseBody) SetRequestId(v string) *AddGtmMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGtmMonitorResponseBody) SetMonitorConfigId(v string) *AddGtmMonitorResponseBody {
	s.MonitorConfigId = &v
	return s
}

type AddGtmMonitorResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddGtmMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddGtmMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGtmMonitorResponse) GoString() string {
	return s.String()
}

func (s *AddGtmMonitorResponse) SetHeaders(v map[string]*string) *AddGtmMonitorResponse {
	s.Headers = v
	return s
}

func (s *AddGtmMonitorResponse) SetBody(v *AddGtmMonitorResponseBody) *AddGtmMonitorResponse {
	s.Body = v
	return s
}

type AddGtmRecoveryPlanRequest struct {
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp  *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	FaultAddrPool *string `json:"FaultAddrPool,omitempty" xml:"FaultAddrPool,omitempty"`
}

func (s AddGtmRecoveryPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGtmRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *AddGtmRecoveryPlanRequest) SetLang(v string) *AddGtmRecoveryPlanRequest {
	s.Lang = &v
	return s
}

func (s *AddGtmRecoveryPlanRequest) SetUserClientIp(v string) *AddGtmRecoveryPlanRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddGtmRecoveryPlanRequest) SetName(v string) *AddGtmRecoveryPlanRequest {
	s.Name = &v
	return s
}

func (s *AddGtmRecoveryPlanRequest) SetRemark(v string) *AddGtmRecoveryPlanRequest {
	s.Remark = &v
	return s
}

func (s *AddGtmRecoveryPlanRequest) SetFaultAddrPool(v string) *AddGtmRecoveryPlanRequest {
	s.FaultAddrPool = &v
	return s
}

type AddGtmRecoveryPlanResponseBody struct {
	RecoveryPlanId *string `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddGtmRecoveryPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGtmRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *AddGtmRecoveryPlanResponseBody) SetRecoveryPlanId(v string) *AddGtmRecoveryPlanResponseBody {
	s.RecoveryPlanId = &v
	return s
}

func (s *AddGtmRecoveryPlanResponseBody) SetRequestId(v string) *AddGtmRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

type AddGtmRecoveryPlanResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddGtmRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddGtmRecoveryPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGtmRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *AddGtmRecoveryPlanResponse) SetHeaders(v map[string]*string) *AddGtmRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *AddGtmRecoveryPlanResponse) SetBody(v *AddGtmRecoveryPlanResponseBody) *AddGtmRecoveryPlanResponse {
	s.Body = v
	return s
}

type BindInstanceDomainsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DomainNames  *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
}

func (s BindInstanceDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s BindInstanceDomainsRequest) GoString() string {
	return s.String()
}

func (s *BindInstanceDomainsRequest) SetLang(v string) *BindInstanceDomainsRequest {
	s.Lang = &v
	return s
}

func (s *BindInstanceDomainsRequest) SetUserClientIp(v string) *BindInstanceDomainsRequest {
	s.UserClientIp = &v
	return s
}

func (s *BindInstanceDomainsRequest) SetInstanceId(v string) *BindInstanceDomainsRequest {
	s.InstanceId = &v
	return s
}

func (s *BindInstanceDomainsRequest) SetDomainNames(v string) *BindInstanceDomainsRequest {
	s.DomainNames = &v
	return s
}

type BindInstanceDomainsResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FailedCount  *int32  `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	SuccessCount *int32  `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s BindInstanceDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindInstanceDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *BindInstanceDomainsResponseBody) SetRequestId(v string) *BindInstanceDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindInstanceDomainsResponseBody) SetFailedCount(v int32) *BindInstanceDomainsResponseBody {
	s.FailedCount = &v
	return s
}

func (s *BindInstanceDomainsResponseBody) SetSuccessCount(v int32) *BindInstanceDomainsResponseBody {
	s.SuccessCount = &v
	return s
}

type BindInstanceDomainsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindInstanceDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindInstanceDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s BindInstanceDomainsResponse) GoString() string {
	return s.String()
}

func (s *BindInstanceDomainsResponse) SetHeaders(v map[string]*string) *BindInstanceDomainsResponse {
	s.Headers = v
	return s
}

func (s *BindInstanceDomainsResponse) SetBody(v *BindInstanceDomainsResponseBody) *BindInstanceDomainsResponse {
	s.Body = v
	return s
}

type ChangeDomainGroupRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ChangeDomainGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeDomainGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeDomainGroupRequest) SetLang(v string) *ChangeDomainGroupRequest {
	s.Lang = &v
	return s
}

func (s *ChangeDomainGroupRequest) SetUserClientIp(v string) *ChangeDomainGroupRequest {
	s.UserClientIp = &v
	return s
}

func (s *ChangeDomainGroupRequest) SetDomainName(v string) *ChangeDomainGroupRequest {
	s.DomainName = &v
	return s
}

func (s *ChangeDomainGroupRequest) SetGroupId(v string) *ChangeDomainGroupRequest {
	s.GroupId = &v
	return s
}

type ChangeDomainGroupResponseBody struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ChangeDomainGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeDomainGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeDomainGroupResponseBody) SetGroupName(v string) *ChangeDomainGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *ChangeDomainGroupResponseBody) SetRequestId(v string) *ChangeDomainGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeDomainGroupResponseBody) SetGroupId(v string) *ChangeDomainGroupResponseBody {
	s.GroupId = &v
	return s
}

type ChangeDomainGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeDomainGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeDomainGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeDomainGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeDomainGroupResponse) SetHeaders(v map[string]*string) *ChangeDomainGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeDomainGroupResponse) SetBody(v *ChangeDomainGroupResponseBody) *ChangeDomainGroupResponse {
	s.Body = v
	return s
}

type ChangeDomainOfDnsProductRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NewDomain    *string `json:"NewDomain,omitempty" xml:"NewDomain,omitempty"`
	Force        *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s ChangeDomainOfDnsProductRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeDomainOfDnsProductRequest) GoString() string {
	return s.String()
}

func (s *ChangeDomainOfDnsProductRequest) SetLang(v string) *ChangeDomainOfDnsProductRequest {
	s.Lang = &v
	return s
}

func (s *ChangeDomainOfDnsProductRequest) SetUserClientIp(v string) *ChangeDomainOfDnsProductRequest {
	s.UserClientIp = &v
	return s
}

func (s *ChangeDomainOfDnsProductRequest) SetInstanceId(v string) *ChangeDomainOfDnsProductRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeDomainOfDnsProductRequest) SetNewDomain(v string) *ChangeDomainOfDnsProductRequest {
	s.NewDomain = &v
	return s
}

func (s *ChangeDomainOfDnsProductRequest) SetForce(v bool) *ChangeDomainOfDnsProductRequest {
	s.Force = &v
	return s
}

type ChangeDomainOfDnsProductResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OriginalDomain *string `json:"OriginalDomain,omitempty" xml:"OriginalDomain,omitempty"`
}

func (s ChangeDomainOfDnsProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeDomainOfDnsProductResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeDomainOfDnsProductResponseBody) SetRequestId(v string) *ChangeDomainOfDnsProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeDomainOfDnsProductResponseBody) SetOriginalDomain(v string) *ChangeDomainOfDnsProductResponseBody {
	s.OriginalDomain = &v
	return s
}

type ChangeDomainOfDnsProductResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeDomainOfDnsProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeDomainOfDnsProductResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeDomainOfDnsProductResponse) GoString() string {
	return s.String()
}

func (s *ChangeDomainOfDnsProductResponse) SetHeaders(v map[string]*string) *ChangeDomainOfDnsProductResponse {
	s.Headers = v
	return s
}

func (s *ChangeDomainOfDnsProductResponse) SetBody(v *ChangeDomainOfDnsProductResponseBody) *ChangeDomainOfDnsProductResponse {
	s.Body = v
	return s
}

type CopyGtmConfigRequest struct {
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceId     *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	TargetId     *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	CopyType     *string `json:"CopyType,omitempty" xml:"CopyType,omitempty"`
}

func (s CopyGtmConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyGtmConfigRequest) GoString() string {
	return s.String()
}

func (s *CopyGtmConfigRequest) SetUserClientIp(v string) *CopyGtmConfigRequest {
	s.UserClientIp = &v
	return s
}

func (s *CopyGtmConfigRequest) SetLang(v string) *CopyGtmConfigRequest {
	s.Lang = &v
	return s
}

func (s *CopyGtmConfigRequest) SetSourceId(v string) *CopyGtmConfigRequest {
	s.SourceId = &v
	return s
}

func (s *CopyGtmConfigRequest) SetTargetId(v string) *CopyGtmConfigRequest {
	s.TargetId = &v
	return s
}

func (s *CopyGtmConfigRequest) SetCopyType(v string) *CopyGtmConfigRequest {
	s.CopyType = &v
	return s
}

type CopyGtmConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyGtmConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyGtmConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CopyGtmConfigResponseBody) SetRequestId(v string) *CopyGtmConfigResponseBody {
	s.RequestId = &v
	return s
}

type CopyGtmConfigResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CopyGtmConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CopyGtmConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyGtmConfigResponse) GoString() string {
	return s.String()
}

func (s *CopyGtmConfigResponse) SetHeaders(v map[string]*string) *CopyGtmConfigResponse {
	s.Headers = v
	return s
}

func (s *CopyGtmConfigResponse) SetBody(v *CopyGtmConfigResponseBody) *CopyGtmConfigResponse {
	s.Body = v
	return s
}

type DeleteCustomLinesRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	LineIds      *string `json:"LineIds,omitempty" xml:"LineIds,omitempty"`
}

func (s DeleteCustomLinesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomLinesRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomLinesRequest) SetLang(v string) *DeleteCustomLinesRequest {
	s.Lang = &v
	return s
}

func (s *DeleteCustomLinesRequest) SetUserClientIp(v string) *DeleteCustomLinesRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteCustomLinesRequest) SetLineIds(v string) *DeleteCustomLinesRequest {
	s.LineIds = &v
	return s
}

type DeleteCustomLinesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomLinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomLinesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomLinesResponseBody) SetRequestId(v string) *DeleteCustomLinesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCustomLinesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCustomLinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCustomLinesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomLinesResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomLinesResponse) SetHeaders(v map[string]*string) *DeleteCustomLinesResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomLinesResponse) SetBody(v *DeleteCustomLinesResponseBody) *DeleteCustomLinesResponse {
	s.Body = v
	return s
}

type DeleteDnsGtmAccessStrategyRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	StrategyId   *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DeleteDnsGtmAccessStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDnsGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteDnsGtmAccessStrategyRequest) SetLang(v string) *DeleteDnsGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDnsGtmAccessStrategyRequest) SetUserClientIp(v string) *DeleteDnsGtmAccessStrategyRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteDnsGtmAccessStrategyRequest) SetStrategyId(v string) *DeleteDnsGtmAccessStrategyRequest {
	s.StrategyId = &v
	return s
}

type DeleteDnsGtmAccessStrategyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDnsGtmAccessStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDnsGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDnsGtmAccessStrategyResponseBody) SetRequestId(v string) *DeleteDnsGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDnsGtmAccessStrategyResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDnsGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDnsGtmAccessStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDnsGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteDnsGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *DeleteDnsGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteDnsGtmAccessStrategyResponse) SetBody(v *DeleteDnsGtmAccessStrategyResponseBody) *DeleteDnsGtmAccessStrategyResponse {
	s.Body = v
	return s
}

type DeleteDnsGtmAddressPoolRequest struct {
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AddrPoolId   *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
}

func (s DeleteDnsGtmAddressPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDnsGtmAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteDnsGtmAddressPoolRequest) SetUserClientIp(v string) *DeleteDnsGtmAddressPoolRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteDnsGtmAddressPoolRequest) SetLang(v string) *DeleteDnsGtmAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDnsGtmAddressPoolRequest) SetAddrPoolId(v string) *DeleteDnsGtmAddressPoolRequest {
	s.AddrPoolId = &v
	return s
}

type DeleteDnsGtmAddressPoolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDnsGtmAddressPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDnsGtmAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDnsGtmAddressPoolResponseBody) SetRequestId(v string) *DeleteDnsGtmAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDnsGtmAddressPoolResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDnsGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDnsGtmAddressPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDnsGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteDnsGtmAddressPoolResponse) SetHeaders(v map[string]*string) *DeleteDnsGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteDnsGtmAddressPoolResponse) SetBody(v *DeleteDnsGtmAddressPoolResponseBody) *DeleteDnsGtmAddressPoolResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) SetLang(v string) *DeleteDomainRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDomainRequest) SetUserClientIp(v string) *DeleteDomainRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteDomainRequest) SetDomainName(v string) *DeleteDomainRequest {
	s.DomainName = &v
	return s
}

type DeleteDomainResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DeleteDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBody) SetRequestId(v string) *DeleteDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainResponseBody) SetDomainName(v string) *DeleteDomainResponseBody {
	s.DomainName = &v
	return s
}

type DeleteDomainResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetBody(v *DeleteDomainResponseBody) *DeleteDomainResponse {
	s.Body = v
	return s
}

type DeleteDomainGroupRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DeleteDomainGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainGroupRequest) SetLang(v string) *DeleteDomainGroupRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDomainGroupRequest) SetUserClientIp(v string) *DeleteDomainGroupRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteDomainGroupRequest) SetGroupId(v string) *DeleteDomainGroupRequest {
	s.GroupId = &v
	return s
}

type DeleteDomainGroupResponseBody struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainGroupResponseBody) SetGroupName(v string) *DeleteDomainGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *DeleteDomainGroupResponseBody) SetRequestId(v string) *DeleteDomainGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDomainGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDomainGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainGroupResponse) SetHeaders(v map[string]*string) *DeleteDomainGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainGroupResponse) SetBody(v *DeleteDomainGroupResponseBody) *DeleteDomainGroupResponse {
	s.Body = v
	return s
}

type DeleteDomainRecordRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	RecordId     *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s DeleteDomainRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRecordRequest) SetLang(v string) *DeleteDomainRecordRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDomainRecordRequest) SetUserClientIp(v string) *DeleteDomainRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteDomainRecordRequest) SetRecordId(v string) *DeleteDomainRecordRequest {
	s.RecordId = &v
	return s
}

type DeleteDomainRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s DeleteDomainRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainRecordResponseBody) SetRequestId(v string) *DeleteDomainRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainRecordResponseBody) SetRecordId(v string) *DeleteDomainRecordResponseBody {
	s.RecordId = &v
	return s
}

type DeleteDomainRecordResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDomainRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDomainRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainRecordResponse) SetHeaders(v map[string]*string) *DeleteDomainRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainRecordResponse) SetBody(v *DeleteDomainRecordResponseBody) *DeleteDomainRecordResponse {
	s.Body = v
	return s
}

type DeleteGtmAccessStrategyRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	StrategyId   *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DeleteGtmAccessStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteGtmAccessStrategyRequest) SetLang(v string) *DeleteGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteGtmAccessStrategyRequest) SetUserClientIp(v string) *DeleteGtmAccessStrategyRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteGtmAccessStrategyRequest) SetStrategyId(v string) *DeleteGtmAccessStrategyRequest {
	s.StrategyId = &v
	return s
}

type DeleteGtmAccessStrategyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGtmAccessStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGtmAccessStrategyResponseBody) SetRequestId(v string) *DeleteGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGtmAccessStrategyResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGtmAccessStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *DeleteGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteGtmAccessStrategyResponse) SetBody(v *DeleteGtmAccessStrategyResponseBody) *DeleteGtmAccessStrategyResponse {
	s.Body = v
	return s
}

type DeleteGtmAddressPoolRequest struct {
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AddrPoolId   *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
}

func (s DeleteGtmAddressPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGtmAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteGtmAddressPoolRequest) SetUserClientIp(v string) *DeleteGtmAddressPoolRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteGtmAddressPoolRequest) SetLang(v string) *DeleteGtmAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *DeleteGtmAddressPoolRequest) SetAddrPoolId(v string) *DeleteGtmAddressPoolRequest {
	s.AddrPoolId = &v
	return s
}

type DeleteGtmAddressPoolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGtmAddressPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGtmAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGtmAddressPoolResponseBody) SetRequestId(v string) *DeleteGtmAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGtmAddressPoolResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGtmAddressPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteGtmAddressPoolResponse) SetHeaders(v map[string]*string) *DeleteGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteGtmAddressPoolResponse) SetBody(v *DeleteGtmAddressPoolResponseBody) *DeleteGtmAddressPoolResponse {
	s.Body = v
	return s
}

type DeleteGtmRecoveryPlanRequest struct {
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp   *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	RecoveryPlanId *int64  `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
}

func (s DeleteGtmRecoveryPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGtmRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteGtmRecoveryPlanRequest) SetLang(v string) *DeleteGtmRecoveryPlanRequest {
	s.Lang = &v
	return s
}

func (s *DeleteGtmRecoveryPlanRequest) SetUserClientIp(v string) *DeleteGtmRecoveryPlanRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteGtmRecoveryPlanRequest) SetRecoveryPlanId(v int64) *DeleteGtmRecoveryPlanRequest {
	s.RecoveryPlanId = &v
	return s
}

type DeleteGtmRecoveryPlanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGtmRecoveryPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGtmRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGtmRecoveryPlanResponseBody) SetRequestId(v string) *DeleteGtmRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGtmRecoveryPlanResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGtmRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGtmRecoveryPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGtmRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteGtmRecoveryPlanResponse) SetHeaders(v map[string]*string) *DeleteGtmRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteGtmRecoveryPlanResponse) SetBody(v *DeleteGtmRecoveryPlanResponseBody) *DeleteGtmRecoveryPlanResponse {
	s.Body = v
	return s
}

type DeleteSubDomainRecordsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RR           *string `json:"RR,omitempty" xml:"RR,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteSubDomainRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubDomainRecordsRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubDomainRecordsRequest) SetLang(v string) *DeleteSubDomainRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DeleteSubDomainRecordsRequest) SetUserClientIp(v string) *DeleteSubDomainRecordsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteSubDomainRecordsRequest) SetDomainName(v string) *DeleteSubDomainRecordsRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteSubDomainRecordsRequest) SetRR(v string) *DeleteSubDomainRecordsRequest {
	s.RR = &v
	return s
}

func (s *DeleteSubDomainRecordsRequest) SetType(v string) *DeleteSubDomainRecordsRequest {
	s.Type = &v
	return s
}

type DeleteSubDomainRecordsResponseBody struct {
	RR         *string `json:"RR,omitempty" xml:"RR,omitempty"`
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSubDomainRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubDomainRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubDomainRecordsResponseBody) SetRR(v string) *DeleteSubDomainRecordsResponseBody {
	s.RR = &v
	return s
}

func (s *DeleteSubDomainRecordsResponseBody) SetTotalCount(v string) *DeleteSubDomainRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DeleteSubDomainRecordsResponseBody) SetRequestId(v string) *DeleteSubDomainRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSubDomainRecordsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSubDomainRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSubDomainRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubDomainRecordsResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubDomainRecordsResponse) SetHeaders(v map[string]*string) *DeleteSubDomainRecordsResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubDomainRecordsResponse) SetBody(v *DeleteSubDomainRecordsResponseBody) *DeleteSubDomainRecordsResponse {
	s.Body = v
	return s
}

type DescribeBatchResultCountRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	BatchType    *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
}

func (s DescribeBatchResultCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchResultCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultCountRequest) SetLang(v string) *DescribeBatchResultCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBatchResultCountRequest) SetUserClientIp(v string) *DescribeBatchResultCountRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeBatchResultCountRequest) SetTaskId(v int64) *DescribeBatchResultCountRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeBatchResultCountRequest) SetBatchType(v string) *DescribeBatchResultCountRequest {
	s.BatchType = &v
	return s
}

type DescribeBatchResultCountResponseBody struct {
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalCount   *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FailedCount  *int32  `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	SuccessCount *int32  `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	BatchType    *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeBatchResultCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchResultCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultCountResponseBody) SetStatus(v int32) *DescribeBatchResultCountResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) SetTotalCount(v int32) *DescribeBatchResultCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) SetTaskId(v int64) *DescribeBatchResultCountResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) SetRequestId(v string) *DescribeBatchResultCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) SetFailedCount(v int32) *DescribeBatchResultCountResponseBody {
	s.FailedCount = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) SetSuccessCount(v int32) *DescribeBatchResultCountResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) SetBatchType(v string) *DescribeBatchResultCountResponseBody {
	s.BatchType = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) SetReason(v string) *DescribeBatchResultCountResponseBody {
	s.Reason = &v
	return s
}

type DescribeBatchResultCountResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBatchResultCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBatchResultCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchResultCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultCountResponse) SetHeaders(v map[string]*string) *DescribeBatchResultCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeBatchResultCountResponse) SetBody(v *DescribeBatchResultCountResponseBody) *DescribeBatchResultCountResponse {
	s.Body = v
	return s
}

type DescribeBatchResultDetailRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	BatchType    *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBatchResultDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchResultDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultDetailRequest) SetLang(v string) *DescribeBatchResultDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBatchResultDetailRequest) SetUserClientIp(v string) *DescribeBatchResultDetailRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeBatchResultDetailRequest) SetPageNumber(v int32) *DescribeBatchResultDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBatchResultDetailRequest) SetPageSize(v int32) *DescribeBatchResultDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBatchResultDetailRequest) SetTaskId(v int64) *DescribeBatchResultDetailRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeBatchResultDetailRequest) SetBatchType(v string) *DescribeBatchResultDetailRequest {
	s.BatchType = &v
	return s
}

func (s *DescribeBatchResultDetailRequest) SetStatus(v string) *DescribeBatchResultDetailRequest {
	s.Status = &v
	return s
}

type DescribeBatchResultDetailResponseBody struct {
	TotalCount         *int64                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	BatchResultDetails []*DescribeBatchResultDetailResponseBodyBatchResultDetails `json:"BatchResultDetails,omitempty" xml:"BatchResultDetails,omitempty" type:"Repeated"`
	PageSize           *int64                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber         *int64                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeBatchResultDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchResultDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultDetailResponseBody) SetTotalCount(v int64) *DescribeBatchResultDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBody) SetBatchResultDetails(v []*DescribeBatchResultDetailResponseBodyBatchResultDetails) *DescribeBatchResultDetailResponseBody {
	s.BatchResultDetails = v
	return s
}

func (s *DescribeBatchResultDetailResponseBody) SetPageSize(v int64) *DescribeBatchResultDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBody) SetRequestId(v string) *DescribeBatchResultDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBody) SetPageNumber(v int64) *DescribeBatchResultDetailResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeBatchResultDetailResponseBodyBatchResultDetails struct {
	Status         *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Remark         *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RecordId       *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Rr             *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	Priority       *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RrStatus       *string `json:"RrStatus,omitempty" xml:"RrStatus,omitempty"`
	OperateDateStr *string `json:"OperateDateStr,omitempty" xml:"OperateDateStr,omitempty"`
	NewValue       *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	Value          *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Ttl            *string `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	BatchType      *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	Line           *string `json:"Line,omitempty" xml:"Line,omitempty"`
	NewRr          *string `json:"NewRr,omitempty" xml:"NewRr,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeBatchResultDetailResponseBodyBatchResultDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchResultDetailResponseBodyBatchResultDetails) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetStatus(v bool) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.Status = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetType(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.Type = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetDomain(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.Domain = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetRemark(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.Remark = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetRecordId(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.RecordId = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetRr(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.Rr = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetPriority(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.Priority = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetRrStatus(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.RrStatus = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetOperateDateStr(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.OperateDateStr = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetNewValue(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.NewValue = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetValue(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.Value = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetTtl(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.Ttl = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetBatchType(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.BatchType = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetLine(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.Line = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetNewRr(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.NewRr = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetReason(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.Reason = &v
	return s
}

type DescribeBatchResultDetailResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBatchResultDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBatchResultDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchResultDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultDetailResponse) SetHeaders(v map[string]*string) *DescribeBatchResultDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeBatchResultDetailResponse) SetBody(v *DescribeBatchResultDetailResponseBody) *DescribeBatchResultDetailResponse {
	s.Body = v
	return s
}

type DescribeCustomLineRequest struct {
	LineId       *int64  `json:"LineId,omitempty" xml:"LineId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeCustomLineRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomLineRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomLineRequest) SetLineId(v int64) *DescribeCustomLineRequest {
	s.LineId = &v
	return s
}

func (s *DescribeCustomLineRequest) SetLang(v string) *DescribeCustomLineRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCustomLineRequest) SetUserClientIp(v string) *DescribeCustomLineRequest {
	s.UserClientIp = &v
	return s
}

type DescribeCustomLineResponseBody struct {
	RequestId     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IpSegmentList []*DescribeCustomLineResponseBodyIpSegmentList `json:"IpSegmentList,omitempty" xml:"IpSegmentList,omitempty" type:"Repeated"`
	DomainName    *string                                        `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Id            *int64                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	Code          *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Name          *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeCustomLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomLineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomLineResponseBody) SetRequestId(v string) *DescribeCustomLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomLineResponseBody) SetIpSegmentList(v []*DescribeCustomLineResponseBodyIpSegmentList) *DescribeCustomLineResponseBody {
	s.IpSegmentList = v
	return s
}

func (s *DescribeCustomLineResponseBody) SetDomainName(v string) *DescribeCustomLineResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeCustomLineResponseBody) SetId(v int64) *DescribeCustomLineResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeCustomLineResponseBody) SetCode(v string) *DescribeCustomLineResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomLineResponseBody) SetName(v string) *DescribeCustomLineResponseBody {
	s.Name = &v
	return s
}

type DescribeCustomLineResponseBodyIpSegmentList struct {
	EndIp   *string `json:"EndIp,omitempty" xml:"EndIp,omitempty"`
	StartIp *string `json:"StartIp,omitempty" xml:"StartIp,omitempty"`
}

func (s DescribeCustomLineResponseBodyIpSegmentList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomLineResponseBodyIpSegmentList) GoString() string {
	return s.String()
}

func (s *DescribeCustomLineResponseBodyIpSegmentList) SetEndIp(v string) *DescribeCustomLineResponseBodyIpSegmentList {
	s.EndIp = &v
	return s
}

func (s *DescribeCustomLineResponseBodyIpSegmentList) SetStartIp(v string) *DescribeCustomLineResponseBodyIpSegmentList {
	s.StartIp = &v
	return s
}

type DescribeCustomLineResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCustomLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomLineResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomLineResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomLineResponse) SetHeaders(v map[string]*string) *DescribeCustomLineResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomLineResponse) SetBody(v *DescribeCustomLineResponseBody) *DescribeCustomLineResponse {
	s.Body = v
	return s
}

type DescribeCustomLinesRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeCustomLinesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomLinesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomLinesRequest) SetLang(v string) *DescribeCustomLinesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCustomLinesRequest) SetUserClientIp(v string) *DescribeCustomLinesRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeCustomLinesRequest) SetPageNumber(v int64) *DescribeCustomLinesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomLinesRequest) SetPageSize(v int64) *DescribeCustomLinesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomLinesRequest) SetDomainName(v string) *DescribeCustomLinesRequest {
	s.DomainName = &v
	return s
}

type DescribeCustomLinesResponseBody struct {
	PageSize    *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber  *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalPages  *int32                                        `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	CustomLines []*DescribeCustomLinesResponseBodyCustomLines `json:"CustomLines,omitempty" xml:"CustomLines,omitempty" type:"Repeated"`
	TotalItems  *int32                                        `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
}

func (s DescribeCustomLinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomLinesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomLinesResponseBody) SetPageSize(v int32) *DescribeCustomLinesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomLinesResponseBody) SetRequestId(v string) *DescribeCustomLinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomLinesResponseBody) SetPageNumber(v int32) *DescribeCustomLinesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomLinesResponseBody) SetTotalPages(v int32) *DescribeCustomLinesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeCustomLinesResponseBody) SetCustomLines(v []*DescribeCustomLinesResponseBodyCustomLines) *DescribeCustomLinesResponseBody {
	s.CustomLines = v
	return s
}

func (s *DescribeCustomLinesResponseBody) SetTotalItems(v int32) *DescribeCustomLinesResponseBody {
	s.TotalItems = &v
	return s
}

type DescribeCustomLinesResponseBodyCustomLines struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeCustomLinesResponseBodyCustomLines) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomLinesResponseBodyCustomLines) GoString() string {
	return s.String()
}

func (s *DescribeCustomLinesResponseBodyCustomLines) SetCode(v string) *DescribeCustomLinesResponseBodyCustomLines {
	s.Code = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLines) SetName(v string) *DescribeCustomLinesResponseBodyCustomLines {
	s.Name = &v
	return s
}

func (s *DescribeCustomLinesResponseBodyCustomLines) SetId(v int64) *DescribeCustomLinesResponseBodyCustomLines {
	s.Id = &v
	return s
}

type DescribeCustomLinesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCustomLinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomLinesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomLinesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomLinesResponse) SetHeaders(v map[string]*string) *DescribeCustomLinesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomLinesResponse) SetBody(v *DescribeCustomLinesResponseBody) *DescribeCustomLinesResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmAccessStrategiesRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StrategyMode *string `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
}

func (s DescribeDnsGtmAccessStrategiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesRequest) SetLang(v string) *DescribeDnsGtmAccessStrategiesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesRequest) SetUserClientIp(v string) *DescribeDnsGtmAccessStrategiesRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesRequest) SetInstanceId(v string) *DescribeDnsGtmAccessStrategiesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesRequest) SetPageNumber(v int32) *DescribeDnsGtmAccessStrategiesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesRequest) SetPageSize(v int32) *DescribeDnsGtmAccessStrategiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesRequest) SetStrategyMode(v string) *DescribeDnsGtmAccessStrategiesRequest {
	s.StrategyMode = &v
	return s
}

type DescribeDnsGtmAccessStrategiesResponseBody struct {
	PageSize   *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Strategies []*DescribeDnsGtmAccessStrategiesResponseBodyStrategies `json:"Strategies,omitempty" xml:"Strategies,omitempty" type:"Repeated"`
	TotalPages *int32                                                  `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems *int32                                                  `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
}

func (s DescribeDnsGtmAccessStrategiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) SetPageSize(v int32) *DescribeDnsGtmAccessStrategiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) SetRequestId(v string) *DescribeDnsGtmAccessStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) SetPageNumber(v int32) *DescribeDnsGtmAccessStrategiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) SetStrategies(v []*DescribeDnsGtmAccessStrategiesResponseBodyStrategies) *DescribeDnsGtmAccessStrategiesResponseBody {
	s.Strategies = v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) SetTotalPages(v int32) *DescribeDnsGtmAccessStrategiesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) SetTotalItems(v int32) *DescribeDnsGtmAccessStrategiesResponseBody {
	s.TotalItems = &v
	return s
}

type DescribeDnsGtmAccessStrategiesResponseBodyStrategies struct {
	EffectiveLbaStrategy       *string                                                                   `json:"EffectiveLbaStrategy,omitempty" xml:"EffectiveLbaStrategy,omitempty"`
	StrategyId                 *string                                                                   `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	StrategyName               *string                                                                   `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	EffectiveAddrPoolGroupType *string                                                                   `json:"EffectiveAddrPoolGroupType,omitempty" xml:"EffectiveAddrPoolGroupType,omitempty"`
	CreateTime                 *string                                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EffectiveAddrPools         []*DescribeDnsGtmAccessStrategiesResponseBodyStrategiesEffectiveAddrPools `json:"EffectiveAddrPools,omitempty" xml:"EffectiveAddrPools,omitempty" type:"Repeated"`
	CreateTimestamp            *int64                                                                    `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EffectiveAddrPoolType      *string                                                                   `json:"EffectiveAddrPoolType,omitempty" xml:"EffectiveAddrPoolType,omitempty"`
	Lines                      []*DescribeDnsGtmAccessStrategiesResponseBodyStrategiesLines              `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategies) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategies) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategies) SetEffectiveLbaStrategy(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategies {
	s.EffectiveLbaStrategy = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategies) SetStrategyId(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategies {
	s.StrategyId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategies) SetStrategyName(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategies {
	s.StrategyName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategies) SetEffectiveAddrPoolGroupType(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategies {
	s.EffectiveAddrPoolGroupType = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategies) SetCreateTime(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategies {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategies) SetEffectiveAddrPools(v []*DescribeDnsGtmAccessStrategiesResponseBodyStrategiesEffectiveAddrPools) *DescribeDnsGtmAccessStrategiesResponseBodyStrategies {
	s.EffectiveAddrPools = v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategies) SetCreateTimestamp(v int64) *DescribeDnsGtmAccessStrategiesResponseBodyStrategies {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategies) SetEffectiveAddrPoolType(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategies {
	s.EffectiveAddrPoolType = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategies) SetLines(v []*DescribeDnsGtmAccessStrategiesResponseBodyStrategiesLines) *DescribeDnsGtmAccessStrategiesResponseBodyStrategies {
	s.Lines = v
	return s
}

type DescribeDnsGtmAccessStrategiesResponseBodyStrategiesEffectiveAddrPools struct {
	LbaWeight *int32  `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AddrCount *int32  `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategiesEffectiveAddrPools) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategiesEffectiveAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesEffectiveAddrPools) SetLbaWeight(v int32) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesEffectiveAddrPools {
	s.LbaWeight = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesEffectiveAddrPools) SetName(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesEffectiveAddrPools {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesEffectiveAddrPools) SetAddrCount(v int32) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesEffectiveAddrPools {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesEffectiveAddrPools) SetId(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesEffectiveAddrPools {
	s.Id = &v
	return s
}

type DescribeDnsGtmAccessStrategiesResponseBodyStrategiesLines struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	LineCode  *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	LineName  *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategiesLines) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategiesLines) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesLines) SetGroupName(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesLines {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesLines) SetLineCode(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesLines {
	s.LineCode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesLines) SetLineName(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesLines {
	s.LineName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesLines) SetGroupCode(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesLines {
	s.GroupCode = &v
	return s
}

type DescribeDnsGtmAccessStrategiesResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmAccessStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmAccessStrategiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmAccessStrategiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponse) SetBody(v *DescribeDnsGtmAccessStrategiesResponseBody) *DescribeDnsGtmAccessStrategiesResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmAccessStrategyRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	StrategyId   *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyRequest) SetLang(v string) *DescribeDnsGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyRequest) SetUserClientIp(v string) *DescribeDnsGtmAccessStrategyRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyRequest) SetStrategyId(v string) *DescribeDnsGtmAccessStrategyRequest {
	s.StrategyId = &v
	return s
}

type DescribeDnsGtmAccessStrategyResponseBody struct {
	FailoverMinAvailableAddrNum *int32                                                       `json:"FailoverMinAvailableAddrNum,omitempty" xml:"FailoverMinAvailableAddrNum,omitempty"`
	DefaultAddrPoolType         *string                                                      `json:"DefaultAddrPoolType,omitempty" xml:"DefaultAddrPoolType,omitempty"`
	DefaultAvailableAddrNum     *int32                                                       `json:"DefaultAvailableAddrNum,omitempty" xml:"DefaultAvailableAddrNum,omitempty"`
	StrategyId                  *string                                                      `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	FailoverAddrPoolGroupStatus *string                                                      `json:"FailoverAddrPoolGroupStatus,omitempty" xml:"FailoverAddrPoolGroupStatus,omitempty"`
	FailoverAvailableAddrNum    *int32                                                       `json:"FailoverAvailableAddrNum,omitempty" xml:"FailoverAvailableAddrNum,omitempty"`
	FailoverLbaStrategy         *string                                                      `json:"FailoverLbaStrategy,omitempty" xml:"FailoverLbaStrategy,omitempty"`
	DefaultMaxReturnAddrNum     *int32                                                       `json:"DefaultMaxReturnAddrNum,omitempty" xml:"DefaultMaxReturnAddrNum,omitempty"`
	StrategyMode                *string                                                      `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
	CreateTimestamp             *int64                                                       `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	DefaultLbaStrategy          *string                                                      `json:"DefaultLbaStrategy,omitempty" xml:"DefaultLbaStrategy,omitempty"`
	DefaultAddrPoolGroupStatus  *string                                                      `json:"DefaultAddrPoolGroupStatus,omitempty" xml:"DefaultAddrPoolGroupStatus,omitempty"`
	FailoverAddrPoolType        *string                                                      `json:"FailoverAddrPoolType,omitempty" xml:"FailoverAddrPoolType,omitempty"`
	RequestId                   *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId                  *string                                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	FailoverAddrPools           []*DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools `json:"FailoverAddrPools,omitempty" xml:"FailoverAddrPools,omitempty" type:"Repeated"`
	DefaultLatencyOptimization  *string                                                      `json:"DefaultLatencyOptimization,omitempty" xml:"DefaultLatencyOptimization,omitempty"`
	EffectiveAddrPoolGroupType  *string                                                      `json:"EffectiveAddrPoolGroupType,omitempty" xml:"EffectiveAddrPoolGroupType,omitempty"`
	CreateTime                  *string                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DefaultAddrPools            []*DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools  `json:"DefaultAddrPools,omitempty" xml:"DefaultAddrPools,omitempty" type:"Repeated"`
	DefaultMinAvailableAddrNum  *int32                                                       `json:"DefaultMinAvailableAddrNum,omitempty" xml:"DefaultMinAvailableAddrNum,omitempty"`
	FailoverLatencyOptimization *string                                                      `json:"FailoverLatencyOptimization,omitempty" xml:"FailoverLatencyOptimization,omitempty"`
	StrategyName                *string                                                      `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	FailoverMaxReturnAddrNum    *int32                                                       `json:"FailoverMaxReturnAddrNum,omitempty" xml:"FailoverMaxReturnAddrNum,omitempty"`
	AccessMode                  *string                                                      `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	Lines                       []*DescribeDnsGtmAccessStrategyResponseBodyLines             `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverMinAvailableAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverMinAvailableAddrNum = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultAddrPoolType(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultAddrPoolType = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultAvailableAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultAvailableAddrNum = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetStrategyId(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverAddrPoolGroupStatus(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverAddrPoolGroupStatus = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverAvailableAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverAvailableAddrNum = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverLbaStrategy(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverLbaStrategy = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultMaxReturnAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultMaxReturnAddrNum = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetStrategyMode(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.StrategyMode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetCreateTimestamp(v int64) *DescribeDnsGtmAccessStrategyResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultLbaStrategy(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultLbaStrategy = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultAddrPoolGroupStatus(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultAddrPoolGroupStatus = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverAddrPoolType(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverAddrPoolType = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetRequestId(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetInstanceId(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverAddrPools(v []*DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverAddrPools = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultLatencyOptimization(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultLatencyOptimization = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetEffectiveAddrPoolGroupType(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.EffectiveAddrPoolGroupType = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetCreateTime(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultAddrPools(v []*DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultAddrPools = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultMinAvailableAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultMinAvailableAddrNum = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverLatencyOptimization(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverLatencyOptimization = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetStrategyName(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.StrategyName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverMaxReturnAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverMaxReturnAddrNum = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetAccessMode(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.AccessMode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetLines(v []*DescribeDnsGtmAccessStrategyResponseBodyLines) *DescribeDnsGtmAccessStrategyResponseBody {
	s.Lines = v
	return s
}

type DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools struct {
	LbaWeight *int32  `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AddrCount *int32  `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools) SetLbaWeight(v int32) *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools {
	s.LbaWeight = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools) SetName(v string) *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools) SetAddrCount(v int32) *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools) SetId(v string) *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools {
	s.Id = &v
	return s
}

type DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools struct {
	LbaWeight *int32  `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AddrCount *int32  `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools) SetLbaWeight(v int32) *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools {
	s.LbaWeight = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools) SetName(v string) *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools) SetAddrCount(v int32) *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools) SetId(v string) *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools {
	s.Id = &v
	return s
}

type DescribeDnsGtmAccessStrategyResponseBodyLines struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	LineCode  *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	LineName  *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyResponseBodyLines) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyResponseBodyLines) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLines) SetGroupName(v string) *DescribeDnsGtmAccessStrategyResponseBodyLines {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLines) SetLineCode(v string) *DescribeDnsGtmAccessStrategyResponseBodyLines {
	s.LineCode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLines) SetLineName(v string) *DescribeDnsGtmAccessStrategyResponseBodyLines {
	s.LineName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLines) SetGroupCode(v string) *DescribeDnsGtmAccessStrategyResponseBodyLines {
	s.GroupCode = &v
	return s
}

type DescribeDnsGtmAccessStrategyResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmAccessStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponse) SetBody(v *DescribeDnsGtmAccessStrategyResponseBody) *DescribeDnsGtmAccessStrategyResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmAccessStrategyAvailableConfigRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StrategyMode *string `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigRequest) SetLang(v string) *DescribeDnsGtmAccessStrategyAvailableConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigRequest) SetUserClientIp(v string) *DescribeDnsGtmAccessStrategyAvailableConfigRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigRequest) SetInstanceId(v string) *DescribeDnsGtmAccessStrategyAvailableConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigRequest) SetStrategyMode(v string) *DescribeDnsGtmAccessStrategyAvailableConfigRequest {
	s.StrategyMode = &v
	return s
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBody struct {
	DomainAddrPools       []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools `json:"DomainAddrPools,omitempty" xml:"DomainAddrPools,omitempty" type:"Repeated"`
	Ipv4AddrPools         []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools   `json:"Ipv4AddrPools,omitempty" xml:"Ipv4AddrPools,omitempty" type:"Repeated"`
	RequestId             *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Ipv6AddrPools         []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools   `json:"Ipv6AddrPools,omitempty" xml:"Ipv6AddrPools,omitempty" type:"Repeated"`
	SuggestSetDefaultLine *bool                                                                     `json:"SuggestSetDefaultLine,omitempty" xml:"SuggestSetDefaultLine,omitempty"`
	Lines                 []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines           `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetDomainAddrPools(v []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.DomainAddrPools = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetIpv4AddrPools(v []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.Ipv4AddrPools = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetRequestId(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetIpv6AddrPools(v []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.Ipv6AddrPools = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetSuggestSetDefaultLine(v bool) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.SuggestSetDefaultLine = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetLines(v []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.Lines = v
	return s
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AddrCount *int32  `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools) SetName(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools) SetAddrCount(v int32) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools) SetId(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools {
	s.Id = &v
	return s
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AddrCount *int32  `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools) SetName(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools) SetAddrCount(v int32) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools) SetId(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools {
	s.Id = &v
	return s
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AddrCount *int32  `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools) SetName(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools) SetAddrCount(v int32) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools) SetId(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools {
	s.Id = &v
	return s
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines struct {
	FatherCode *string `json:"FatherCode,omitempty" xml:"FatherCode,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	LineCode   *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	LineName   *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
	GroupCode  *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) SetFatherCode(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines {
	s.FatherCode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) SetGroupName(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) SetLineCode(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines {
	s.LineCode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) SetLineName(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines {
	s.LineName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) SetGroupCode(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines {
	s.GroupCode = &v
	return s
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponse struct {
	Headers map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmAccessStrategyAvailableConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponse) SetBody(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) *DescribeDnsGtmAccessStrategyAvailableConfigResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmAddrAttributeInfoRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Addrs        *string `json:"Addrs,omitempty" xml:"Addrs,omitempty"`
}

func (s DescribeDnsGtmAddrAttributeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAddrAttributeInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddrAttributeInfoRequest) SetLang(v string) *DescribeDnsGtmAddrAttributeInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoRequest) SetUserClientIp(v string) *DescribeDnsGtmAddrAttributeInfoRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoRequest) SetType(v string) *DescribeDnsGtmAddrAttributeInfoRequest {
	s.Type = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoRequest) SetAddrs(v string) *DescribeDnsGtmAddrAttributeInfoRequest {
	s.Addrs = &v
	return s
}

type DescribeDnsGtmAddrAttributeInfoResponseBody struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Addr      []*DescribeDnsGtmAddrAttributeInfoResponseBodyAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAddrAttributeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAddrAttributeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBody) SetRequestId(v string) *DescribeDnsGtmAddrAttributeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBody) SetAddr(v []*DescribeDnsGtmAddrAttributeInfoResponseBodyAddr) *DescribeDnsGtmAddrAttributeInfoResponseBody {
	s.Addr = v
	return s
}

type DescribeDnsGtmAddrAttributeInfoResponseBodyAddr struct {
	AttributeInfo *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty" type:"Struct"`
	Addr          *string                                                       `json:"Addr,omitempty" xml:"Addr,omitempty"`
}

func (s DescribeDnsGtmAddrAttributeInfoResponseBodyAddr) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAddrAttributeInfoResponseBodyAddr) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddr) SetAttributeInfo(v *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddr {
	s.AttributeInfo = v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddr) SetAddr(v string) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddr {
	s.Addr = &v
	return s
}

type DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo struct {
	FatherCode *string `json:"FatherCode,omitempty" xml:"FatherCode,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	LineCode   *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	LineName   *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
	GroupCode  *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
}

func (s DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo) SetFatherCode(v string) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo {
	s.FatherCode = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo) SetGroupName(v string) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo) SetLineCode(v string) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo {
	s.LineCode = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo) SetLineName(v string) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo {
	s.LineName = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo) SetGroupCode(v string) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAttributeInfo {
	s.GroupCode = &v
	return s
}

type DescribeDnsGtmAddrAttributeInfoResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmAddrAttributeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmAddrAttributeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAddrAttributeInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddrAttributeInfoResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmAddrAttributeInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponse) SetBody(v *DescribeDnsGtmAddrAttributeInfoResponseBody) *DescribeDnsGtmAddrAttributeInfoResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmAddressPoolAvailableConfigRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDnsGtmAddressPoolAvailableConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAddressPoolAvailableConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigRequest) SetLang(v string) *DescribeDnsGtmAddressPoolAvailableConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigRequest) SetUserClientIp(v string) *DescribeDnsGtmAddressPoolAvailableConfigRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigRequest) SetInstanceId(v string) *DescribeDnsGtmAddressPoolAvailableConfigRequest {
	s.InstanceId = &v
	return s
}

type DescribeDnsGtmAddressPoolAvailableConfigResponseBody struct {
	RequestId      *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AttributeInfos []*DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos `json:"AttributeInfos,omitempty" xml:"AttributeInfos,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAddressPoolAvailableConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAddressPoolAvailableConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBody) SetRequestId(v string) *DescribeDnsGtmAddressPoolAvailableConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBody) SetAttributeInfos(v []*DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) *DescribeDnsGtmAddressPoolAvailableConfigResponseBody {
	s.AttributeInfos = v
	return s
}

type DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos struct {
	FatherCode *string `json:"FatherCode,omitempty" xml:"FatherCode,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	LineCode   *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	LineName   *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
	GroupCode  *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
}

func (s DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) SetFatherCode(v string) *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos {
	s.FatherCode = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) SetGroupName(v string) *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) SetLineCode(v string) *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos {
	s.LineCode = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) SetLineName(v string) *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos {
	s.LineName = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) SetGroupCode(v string) *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos {
	s.GroupCode = &v
	return s
}

type DescribeDnsGtmAddressPoolAvailableConfigResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmAddressPoolAvailableConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmAddressPoolAvailableConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAddressPoolAvailableConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmAddressPoolAvailableConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponse) SetBody(v *DescribeDnsGtmAddressPoolAvailableConfigResponseBody) *DescribeDnsGtmAddressPoolAvailableConfigResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmAvailableAlertGroupRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeDnsGtmAvailableAlertGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAvailableAlertGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAvailableAlertGroupRequest) SetLang(v string) *DescribeDnsGtmAvailableAlertGroupRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmAvailableAlertGroupRequest) SetUserClientIp(v string) *DescribeDnsGtmAvailableAlertGroupRequest {
	s.UserClientIp = &v
	return s
}

type DescribeDnsGtmAvailableAlertGroupResponseBody struct {
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AvailableAlertGroup *string `json:"AvailableAlertGroup,omitempty" xml:"AvailableAlertGroup,omitempty"`
}

func (s DescribeDnsGtmAvailableAlertGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAvailableAlertGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAvailableAlertGroupResponseBody) SetRequestId(v string) *DescribeDnsGtmAvailableAlertGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmAvailableAlertGroupResponseBody) SetAvailableAlertGroup(v string) *DescribeDnsGtmAvailableAlertGroupResponseBody {
	s.AvailableAlertGroup = &v
	return s
}

type DescribeDnsGtmAvailableAlertGroupResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmAvailableAlertGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmAvailableAlertGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmAvailableAlertGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAvailableAlertGroupResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmAvailableAlertGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmAvailableAlertGroupResponse) SetBody(v *DescribeDnsGtmAvailableAlertGroupResponseBody) *DescribeDnsGtmAvailableAlertGroupResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmInstanceRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDnsGtmInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceRequest) SetLang(v string) *DescribeDnsGtmInstanceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmInstanceRequest) SetUserClientIp(v string) *DescribeDnsGtmInstanceRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsGtmInstanceRequest) SetInstanceId(v string) *DescribeDnsGtmInstanceRequest {
	s.InstanceId = &v
	return s
}

type DescribeDnsGtmInstanceResponseBody struct {
	ExpireTimestamp *int64                                       `json:"ExpireTimestamp,omitempty" xml:"ExpireTimestamp,omitempty"`
	RequestId       *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId *string                                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceId      *string                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskQuota       *int32                                       `json:"TaskQuota,omitempty" xml:"TaskQuota,omitempty"`
	Config          *DescribeDnsGtmInstanceResponseBodyConfig    `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	CreateTime      *string                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SmsQuota        *int32                                       `json:"SmsQuota,omitempty" xml:"SmsQuota,omitempty"`
	VersionCode     *string                                      `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	PaymentType     *string                                      `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	ExpireTime      *string                                      `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CreateTimestamp *int64                                       `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	UsedQuota       *DescribeDnsGtmInstanceResponseBodyUsedQuota `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty" type:"Struct"`
}

func (s DescribeDnsGtmInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceResponseBody) SetExpireTimestamp(v int64) *DescribeDnsGtmInstanceResponseBody {
	s.ExpireTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetRequestId(v string) *DescribeDnsGtmInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetResourceGroupId(v string) *DescribeDnsGtmInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetInstanceId(v string) *DescribeDnsGtmInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetTaskQuota(v int32) *DescribeDnsGtmInstanceResponseBody {
	s.TaskQuota = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetConfig(v *DescribeDnsGtmInstanceResponseBodyConfig) *DescribeDnsGtmInstanceResponseBody {
	s.Config = v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetCreateTime(v string) *DescribeDnsGtmInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetSmsQuota(v int32) *DescribeDnsGtmInstanceResponseBody {
	s.SmsQuota = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetVersionCode(v string) *DescribeDnsGtmInstanceResponseBody {
	s.VersionCode = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetPaymentType(v string) *DescribeDnsGtmInstanceResponseBody {
	s.PaymentType = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetExpireTime(v string) *DescribeDnsGtmInstanceResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetCreateTimestamp(v int64) *DescribeDnsGtmInstanceResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetUsedQuota(v *DescribeDnsGtmInstanceResponseBodyUsedQuota) *DescribeDnsGtmInstanceResponseBody {
	s.UsedQuota = v
	return s
}

type DescribeDnsGtmInstanceResponseBodyConfig struct {
	Ttl                  *int32                                                 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	AlertGroup           *string                                                `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	CnameType            *string                                                `json:"CnameType,omitempty" xml:"CnameType,omitempty"`
	StrategyMode         *string                                                `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
	InstanceName         *string                                                `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PublicCnameMode      *string                                                `json:"PublicCnameMode,omitempty" xml:"PublicCnameMode,omitempty"`
	AlertConfig          []*DescribeDnsGtmInstanceResponseBodyConfigAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
	PublicUserDomainName *string                                                `json:"PublicUserDomainName,omitempty" xml:"PublicUserDomainName,omitempty"`
	PubicZoneName        *string                                                `json:"PubicZoneName,omitempty" xml:"PubicZoneName,omitempty"`
}

func (s DescribeDnsGtmInstanceResponseBodyConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetTtl(v int32) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.Ttl = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetAlertGroup(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.AlertGroup = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetCnameType(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.CnameType = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetStrategyMode(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.StrategyMode = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetInstanceName(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.InstanceName = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetPublicCnameMode(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.PublicCnameMode = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetAlertConfig(v []*DescribeDnsGtmInstanceResponseBodyConfigAlertConfig) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.AlertConfig = v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetPublicUserDomainName(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.PublicUserDomainName = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetPubicZoneName(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.PubicZoneName = &v
	return s
}

type DescribeDnsGtmInstanceResponseBodyConfigAlertConfig struct {
	SmsNotice   *bool   `json:"SmsNotice,omitempty" xml:"SmsNotice,omitempty"`
	NoticeType  *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	EmailNotice *bool   `json:"EmailNotice,omitempty" xml:"EmailNotice,omitempty"`
}

func (s DescribeDnsGtmInstanceResponseBodyConfigAlertConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceResponseBodyConfigAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfig) SetSmsNotice(v bool) *DescribeDnsGtmInstanceResponseBodyConfigAlertConfig {
	s.SmsNotice = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfig) SetNoticeType(v string) *DescribeDnsGtmInstanceResponseBodyConfigAlertConfig {
	s.NoticeType = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfig) SetEmailNotice(v bool) *DescribeDnsGtmInstanceResponseBodyConfigAlertConfig {
	s.EmailNotice = &v
	return s
}

type DescribeDnsGtmInstanceResponseBodyUsedQuota struct {
	EmailUsedCount *int32 `json:"EmailUsedCount,omitempty" xml:"EmailUsedCount,omitempty"`
	TaskUsedCount  *int32 `json:"TaskUsedCount,omitempty" xml:"TaskUsedCount,omitempty"`
	SmsUsedCount   *int32 `json:"SmsUsedCount,omitempty" xml:"SmsUsedCount,omitempty"`
}

func (s DescribeDnsGtmInstanceResponseBodyUsedQuota) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceResponseBodyUsedQuota) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceResponseBodyUsedQuota) SetEmailUsedCount(v int32) *DescribeDnsGtmInstanceResponseBodyUsedQuota {
	s.EmailUsedCount = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyUsedQuota) SetTaskUsedCount(v int32) *DescribeDnsGtmInstanceResponseBodyUsedQuota {
	s.TaskUsedCount = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyUsedQuota) SetSmsUsedCount(v int32) *DescribeDnsGtmInstanceResponseBodyUsedQuota {
	s.SmsUsedCount = &v
	return s
}

type DescribeDnsGtmInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmInstanceResponse) SetBody(v *DescribeDnsGtmInstanceResponseBody) *DescribeDnsGtmInstanceResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmInstanceAddressPoolRequest struct {
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AddrPoolId   *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
}

func (s DescribeDnsGtmInstanceAddressPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolRequest) SetUserClientIp(v string) *DescribeDnsGtmInstanceAddressPoolRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolRequest) SetLang(v string) *DescribeDnsGtmInstanceAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolRequest) SetAddrPoolId(v string) *DescribeDnsGtmInstanceAddressPoolRequest {
	s.AddrPoolId = &v
	return s
}

type DescribeDnsGtmInstanceAddressPoolResponseBody struct {
	Addrs           []*DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs `json:"Addrs,omitempty" xml:"Addrs,omitempty" type:"Repeated"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LbaStrategy     *string                                               `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	CreateTime      *string                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AddrCount       *int32                                                `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	Name            *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Type            *string                                               `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime      *string                                               `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	AddrPoolId      *string                                               `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	UpdateTimestamp *int64                                                `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	MonitorConfigId *string                                               `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	MonitorStatus   *string                                               `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	CreateTimestamp *int64                                                `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeDnsGtmInstanceAddressPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetAddrs(v []*DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.Addrs = v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetRequestId(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetLbaStrategy(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.LbaStrategy = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetCreateTime(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetAddrCount(v int32) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetName(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetType(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetUpdateTime(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetAddrPoolId(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetUpdateTimestamp(v int64) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetMonitorConfigId(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetMonitorStatus(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.MonitorStatus = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetCreateTimestamp(v int64) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.CreateTimestamp = &v
	return s
}

type DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs struct {
	UpdateTimestamp *int64  `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	AttributeInfo   *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	AlertStatus     *string `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	LbaWeight       *int32  `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	Addr            *string `json:"Addr,omitempty" xml:"Addr,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) SetUpdateTimestamp(v int64) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) SetAttributeInfo(v string) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs {
	s.AttributeInfo = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) SetUpdateTime(v string) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) SetAlertStatus(v string) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs {
	s.AlertStatus = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) SetRemark(v string) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs {
	s.Remark = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) SetLbaWeight(v int32) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs {
	s.LbaWeight = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) SetAddr(v string) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs {
	s.Addr = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) SetCreateTime(v string) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) SetMode(v string) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs {
	s.Mode = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) SetCreateTimestamp(v int64) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs {
	s.CreateTimestamp = &v
	return s
}

type DescribeDnsGtmInstanceAddressPoolResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmInstanceAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmInstanceAddressPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponse) SetBody(v *DescribeDnsGtmInstanceAddressPoolResponseBody) *DescribeDnsGtmInstanceAddressPoolResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmInstanceAddressPoolsRequest struct {
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDnsGtmInstanceAddressPoolsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolsRequest) SetUserClientIp(v string) *DescribeDnsGtmInstanceAddressPoolsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsRequest) SetLang(v string) *DescribeDnsGtmInstanceAddressPoolsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsRequest) SetInstanceId(v string) *DescribeDnsGtmInstanceAddressPoolsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsRequest) SetPageNumber(v int32) *DescribeDnsGtmInstanceAddressPoolsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsRequest) SetPageSize(v int32) *DescribeDnsGtmInstanceAddressPoolsRequest {
	s.PageSize = &v
	return s
}

type DescribeDnsGtmInstanceAddressPoolsResponseBody struct {
	PageSize   *int32                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalPages *int32                                                     `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems *int32                                                     `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	AddrPools  []*DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools `json:"AddrPools,omitempty" xml:"AddrPools,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmInstanceAddressPoolsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) SetPageSize(v int32) *DescribeDnsGtmInstanceAddressPoolsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) SetRequestId(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) SetPageNumber(v int32) *DescribeDnsGtmInstanceAddressPoolsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) SetTotalPages(v int32) *DescribeDnsGtmInstanceAddressPoolsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) SetTotalItems(v int32) *DescribeDnsGtmInstanceAddressPoolsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) SetAddrPools(v []*DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) *DescribeDnsGtmInstanceAddressPoolsResponseBody {
	s.AddrPools = v
	return s
}

type DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools struct {
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTimestamp *int64  `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	MonitorStatus   *string `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AddrPoolId      *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	LbaStrategy     *string `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AddrCount       *int32  `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) SetType(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.Type = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) SetUpdateTimestamp(v int64) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) SetMonitorStatus(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.MonitorStatus = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) SetUpdateTime(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) SetCreateTime(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) SetAddrPoolId(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) SetLbaStrategy(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.LbaStrategy = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) SetName(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) SetAddrCount(v int32) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) SetMonitorConfigId(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) SetCreateTimestamp(v int64) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.CreateTimestamp = &v
	return s
}

type DescribeDnsGtmInstanceAddressPoolsResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmInstanceAddressPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmInstanceAddressPoolsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceAddressPoolsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponse) SetBody(v *DescribeDnsGtmInstanceAddressPoolsResponseBody) *DescribeDnsGtmInstanceAddressPoolsResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmInstancesRequest struct {
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp    *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Keyword         *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDnsGtmInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstancesRequest) SetLang(v string) *DescribeDnsGtmInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmInstancesRequest) SetUserClientIp(v string) *DescribeDnsGtmInstancesRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsGtmInstancesRequest) SetPageNumber(v int32) *DescribeDnsGtmInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmInstancesRequest) SetPageSize(v int32) *DescribeDnsGtmInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmInstancesRequest) SetKeyword(v string) *DescribeDnsGtmInstancesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDnsGtmInstancesRequest) SetResourceGroupId(v string) *DescribeDnsGtmInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDnsGtmInstancesResponseBody struct {
	PageSize     *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber   *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	GtmInstances []*DescribeDnsGtmInstancesResponseBodyGtmInstances `json:"GtmInstances,omitempty" xml:"GtmInstances,omitempty" type:"Repeated"`
	TotalPages   *int32                                             `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems   *int32                                             `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
}

func (s DescribeDnsGtmInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstancesResponseBody) SetPageSize(v int32) *DescribeDnsGtmInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBody) SetRequestId(v string) *DescribeDnsGtmInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBody) SetPageNumber(v int32) *DescribeDnsGtmInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBody) SetGtmInstances(v []*DescribeDnsGtmInstancesResponseBodyGtmInstances) *DescribeDnsGtmInstancesResponseBody {
	s.GtmInstances = v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBody) SetTotalPages(v int32) *DescribeDnsGtmInstancesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBody) SetTotalItems(v int32) *DescribeDnsGtmInstancesResponseBody {
	s.TotalItems = &v
	return s
}

type DescribeDnsGtmInstancesResponseBodyGtmInstances struct {
	PaymentType     *string                                                   `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	ExpireTime      *string                                                   `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CreateTime      *string                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SmsQuota        *int32                                                    `json:"SmsQuota,omitempty" xml:"SmsQuota,omitempty"`
	InstanceId      *string                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Config          *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig    `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	ExpireTimestamp *int64                                                    `json:"ExpireTimestamp,omitempty" xml:"ExpireTimestamp,omitempty"`
	ResourceGroupId *string                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	VersionCode     *string                                                   `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	UsedQuota       *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty" type:"Struct"`
	TaskQuota       *int32                                                    `json:"TaskQuota,omitempty" xml:"TaskQuota,omitempty"`
	CreateTimestamp *int64                                                    `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstances) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetPaymentType(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.PaymentType = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetExpireTime(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetCreateTime(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetSmsQuota(v int32) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.SmsQuota = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetInstanceId(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetConfig(v *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.Config = v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetExpireTimestamp(v int64) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.ExpireTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetResourceGroupId(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetVersionCode(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.VersionCode = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetUsedQuota(v *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.UsedQuota = v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetTaskQuota(v int32) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.TaskQuota = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetCreateTimestamp(v int64) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.CreateTimestamp = &v
	return s
}

type DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig struct {
	Ttl                  *int32                                                              `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	AlertGroup           *string                                                             `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	PublicZoneName       *string                                                             `json:"PublicZoneName,omitempty" xml:"PublicZoneName,omitempty"`
	CnameType            *string                                                             `json:"CnameType,omitempty" xml:"CnameType,omitempty"`
	StrategyMode         *string                                                             `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
	InstanceName         *string                                                             `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PublicCnameMode      *string                                                             `json:"PublicCnameMode,omitempty" xml:"PublicCnameMode,omitempty"`
	AlertConfig          []*DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
	PublicUserDomainName *string                                                             `json:"PublicUserDomainName,omitempty" xml:"PublicUserDomainName,omitempty"`
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetTtl(v int32) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.Ttl = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetAlertGroup(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.AlertGroup = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetPublicZoneName(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.PublicZoneName = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetCnameType(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.CnameType = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetStrategyMode(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.StrategyMode = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetInstanceName(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.InstanceName = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetPublicCnameMode(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.PublicCnameMode = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetAlertConfig(v []*DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.AlertConfig = v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetPublicUserDomainName(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.PublicUserDomainName = &v
	return s
}

type DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig struct {
	SmsNotice   *string `json:"SmsNotice,omitempty" xml:"SmsNotice,omitempty"`
	NoticeType  *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	EmailNotice *string `json:"EmailNotice,omitempty" xml:"EmailNotice,omitempty"`
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) SetSmsNotice(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig {
	s.SmsNotice = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) SetNoticeType(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig {
	s.NoticeType = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) SetEmailNotice(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig {
	s.EmailNotice = &v
	return s
}

type DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota struct {
	EmailUsedCount *int32 `json:"EmailUsedCount,omitempty" xml:"EmailUsedCount,omitempty"`
	TaskUsedCount  *int32 `json:"TaskUsedCount,omitempty" xml:"TaskUsedCount,omitempty"`
	SmsUsedCount   *int32 `json:"SmsUsedCount,omitempty" xml:"SmsUsedCount,omitempty"`
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) SetEmailUsedCount(v int32) *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota {
	s.EmailUsedCount = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) SetTaskUsedCount(v int32) *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota {
	s.TaskUsedCount = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) SetSmsUsedCount(v int32) *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota {
	s.SmsUsedCount = &v
	return s
}

type DescribeDnsGtmInstancesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstancesResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmInstancesResponse) SetBody(v *DescribeDnsGtmInstancesResponseBody) *DescribeDnsGtmInstancesResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmInstanceStatusRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDnsGtmInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceStatusRequest) SetLang(v string) *DescribeDnsGtmInstanceStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusRequest) SetUserClientIp(v string) *DescribeDnsGtmInstanceStatusRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusRequest) SetInstanceId(v string) *DescribeDnsGtmInstanceStatusRequest {
	s.InstanceId = &v
	return s
}

type DescribeDnsGtmInstanceStatusResponseBody struct {
	StrategyNotAvailableNum      *int32  `json:"StrategyNotAvailableNum,omitempty" xml:"StrategyNotAvailableNum,omitempty"`
	AddrAvailableNum             *int32  `json:"AddrAvailableNum,omitempty" xml:"AddrAvailableNum,omitempty"`
	RequestId                    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SwitchToFailoverStrategyNum  *int32  `json:"SwitchToFailoverStrategyNum,omitempty" xml:"SwitchToFailoverStrategyNum,omitempty"`
	AddrNotAvailableNum          *int32  `json:"AddrNotAvailableNum,omitempty" xml:"AddrNotAvailableNum,omitempty"`
	AddrPoolGroupNotAvailableNum *int32  `json:"AddrPoolGroupNotAvailableNum,omitempty" xml:"AddrPoolGroupNotAvailableNum,omitempty"`
}

func (s DescribeDnsGtmInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) SetStrategyNotAvailableNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody {
	s.StrategyNotAvailableNum = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) SetAddrAvailableNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody {
	s.AddrAvailableNum = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) SetRequestId(v string) *DescribeDnsGtmInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) SetSwitchToFailoverStrategyNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody {
	s.SwitchToFailoverStrategyNum = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) SetAddrNotAvailableNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody {
	s.AddrNotAvailableNum = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) SetAddrPoolGroupNotAvailableNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody {
	s.AddrPoolGroupNotAvailableNum = &v
	return s
}

type DescribeDnsGtmInstanceStatusResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceStatusResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponse) SetBody(v *DescribeDnsGtmInstanceStatusResponseBody) *DescribeDnsGtmInstanceStatusResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmInstanceSystemCnameRequest struct {
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDnsGtmInstanceSystemCnameRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceSystemCnameRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceSystemCnameRequest) SetUserClientIp(v string) *DescribeDnsGtmInstanceSystemCnameRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsGtmInstanceSystemCnameRequest) SetLang(v string) *DescribeDnsGtmInstanceSystemCnameRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmInstanceSystemCnameRequest) SetInstanceId(v string) *DescribeDnsGtmInstanceSystemCnameRequest {
	s.InstanceId = &v
	return s
}

type DescribeDnsGtmInstanceSystemCnameResponseBody struct {
	SystemCname *string `json:"SystemCname,omitempty" xml:"SystemCname,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDnsGtmInstanceSystemCnameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceSystemCnameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceSystemCnameResponseBody) SetSystemCname(v string) *DescribeDnsGtmInstanceSystemCnameResponseBody {
	s.SystemCname = &v
	return s
}

func (s *DescribeDnsGtmInstanceSystemCnameResponseBody) SetRequestId(v string) *DescribeDnsGtmInstanceSystemCnameResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDnsGtmInstanceSystemCnameResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmInstanceSystemCnameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmInstanceSystemCnameResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmInstanceSystemCnameResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceSystemCnameResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceSystemCnameResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmInstanceSystemCnameResponse) SetBody(v *DescribeDnsGtmInstanceSystemCnameResponseBody) *DescribeDnsGtmInstanceSystemCnameResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmLogsRequest struct {
	UserClientIp   *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Keyword        *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTimestamp *int64  `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	EndTimestamp   *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
}

func (s DescribeDnsGtmLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmLogsRequest) SetUserClientIp(v string) *DescribeDnsGtmLogsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsGtmLogsRequest) SetLang(v string) *DescribeDnsGtmLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmLogsRequest) SetInstanceId(v string) *DescribeDnsGtmLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmLogsRequest) SetKeyword(v string) *DescribeDnsGtmLogsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDnsGtmLogsRequest) SetPageNumber(v int32) *DescribeDnsGtmLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmLogsRequest) SetPageSize(v int32) *DescribeDnsGtmLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmLogsRequest) SetStartTimestamp(v int64) *DescribeDnsGtmLogsRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeDnsGtmLogsRequest) SetEndTimestamp(v int64) *DescribeDnsGtmLogsRequest {
	s.EndTimestamp = &v
	return s
}

type DescribeDnsGtmLogsResponseBody struct {
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalPages *int32                                `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Logs       []*DescribeDnsGtmLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	TotalItems *int32                                `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
}

func (s DescribeDnsGtmLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmLogsResponseBody) SetPageSize(v int32) *DescribeDnsGtmLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBody) SetRequestId(v string) *DescribeDnsGtmLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBody) SetPageNumber(v int32) *DescribeDnsGtmLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBody) SetTotalPages(v int32) *DescribeDnsGtmLogsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBody) SetLogs(v []*DescribeDnsGtmLogsResponseBodyLogs) *DescribeDnsGtmLogsResponseBody {
	s.Logs = v
	return s
}

func (s *DescribeDnsGtmLogsResponseBody) SetTotalItems(v int32) *DescribeDnsGtmLogsResponseBody {
	s.TotalItems = &v
	return s
}

type DescribeDnsGtmLogsResponseBodyLogs struct {
	OperTimestamp *int64  `json:"OperTimestamp,omitempty" xml:"OperTimestamp,omitempty"`
	EntityId      *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType    *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	OperTime      *string `json:"OperTime,omitempty" xml:"OperTime,omitempty"`
	OperAction    *string `json:"OperAction,omitempty" xml:"OperAction,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	EntityName    *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeDnsGtmLogsResponseBodyLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmLogsResponseBodyLogs) SetOperTimestamp(v int64) *DescribeDnsGtmLogsResponseBodyLogs {
	s.OperTimestamp = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogs) SetEntityId(v string) *DescribeDnsGtmLogsResponseBodyLogs {
	s.EntityId = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogs) SetEntityType(v string) *DescribeDnsGtmLogsResponseBodyLogs {
	s.EntityType = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogs) SetOperTime(v string) *DescribeDnsGtmLogsResponseBodyLogs {
	s.OperTime = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogs) SetOperAction(v string) *DescribeDnsGtmLogsResponseBodyLogs {
	s.OperAction = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogs) SetContent(v string) *DescribeDnsGtmLogsResponseBodyLogs {
	s.Content = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogs) SetEntityName(v string) *DescribeDnsGtmLogsResponseBodyLogs {
	s.EntityName = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogs) SetId(v int64) *DescribeDnsGtmLogsResponseBodyLogs {
	s.Id = &v
	return s
}

type DescribeDnsGtmLogsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmLogsResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmLogsResponse) SetBody(v *DescribeDnsGtmLogsResponseBody) *DescribeDnsGtmLogsResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmMonitorAvailableConfigRequest struct {
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeDnsGtmMonitorAvailableConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigRequest) SetUserClientIp(v string) *DescribeDnsGtmMonitorAvailableConfigRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigRequest) SetLang(v string) *DescribeDnsGtmMonitorAvailableConfigRequest {
	s.Lang = &v
	return s
}

type DescribeDnsGtmMonitorAvailableConfigResponseBody struct {
	Ipv4IspCityNodes       []*DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes       `json:"Ipv4IspCityNodes,omitempty" xml:"Ipv4IspCityNodes,omitempty" type:"Repeated"`
	DomainIpv4IspCityNodes []*DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes `json:"DomainIpv4IspCityNodes,omitempty" xml:"DomainIpv4IspCityNodes,omitempty" type:"Repeated"`
	RequestId              *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainIpv6IspCityNodes []*DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes `json:"DomainIpv6IspCityNodes,omitempty" xml:"DomainIpv6IspCityNodes,omitempty" type:"Repeated"`
	Ipv6IspCityNodes       []*DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes       `json:"Ipv6IspCityNodes,omitempty" xml:"Ipv6IspCityNodes,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) SetIpv4IspCityNodes(v []*DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) *DescribeDnsGtmMonitorAvailableConfigResponseBody {
	s.Ipv4IspCityNodes = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) SetDomainIpv4IspCityNodes(v []*DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) *DescribeDnsGtmMonitorAvailableConfigResponseBody {
	s.DomainIpv4IspCityNodes = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) SetRequestId(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) SetDomainIpv6IspCityNodes(v []*DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) *DescribeDnsGtmMonitorAvailableConfigResponseBody {
	s.DomainIpv6IspCityNodes = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) SetIpv6IspCityNodes(v []*DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) *DescribeDnsGtmMonitorAvailableConfigResponseBody {
	s.Ipv6IspCityNodes = v
	return s
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes struct {
	CityCode        *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IspCode         *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	CityName        *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	IspName         *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	GroupType       *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	DefaultSelected *bool   `json:"DefaultSelected,omitempty" xml:"DefaultSelected,omitempty"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) SetCityCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes {
	s.CityCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) SetGroupName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) SetIspCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes {
	s.IspCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) SetCityName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes {
	s.CityName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) SetIspName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes {
	s.IspName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) SetGroupType(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes {
	s.GroupType = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) SetDefaultSelected(v bool) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes {
	s.DefaultSelected = &v
	return s
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes struct {
	CityCode        *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IspCode         *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	CityName        *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	IspName         *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	GroupType       *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	DefaultSelected *bool   `json:"DefaultSelected,omitempty" xml:"DefaultSelected,omitempty"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) SetCityCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes {
	s.CityCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) SetGroupName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) SetIspCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes {
	s.IspCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) SetCityName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes {
	s.CityName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) SetIspName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes {
	s.IspName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) SetGroupType(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes {
	s.GroupType = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) SetDefaultSelected(v bool) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes {
	s.DefaultSelected = &v
	return s
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes struct {
	CityCode        *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IspCode         *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	CityName        *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	IspName         *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	GroupType       *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	DefaultSelected *bool   `json:"DefaultSelected,omitempty" xml:"DefaultSelected,omitempty"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) SetCityCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes {
	s.CityCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) SetGroupName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) SetIspCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes {
	s.IspCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) SetCityName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes {
	s.CityName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) SetIspName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes {
	s.IspName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) SetGroupType(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes {
	s.GroupType = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) SetDefaultSelected(v bool) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes {
	s.DefaultSelected = &v
	return s
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes struct {
	CityCode        *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IspCode         *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	CityName        *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	IspName         *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	GroupType       *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	DefaultSelected *bool   `json:"DefaultSelected,omitempty" xml:"DefaultSelected,omitempty"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) SetCityCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes {
	s.CityCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) SetGroupName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) SetIspCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes {
	s.IspCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) SetCityName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes {
	s.CityName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) SetIspName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes {
	s.IspName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) SetGroupType(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes {
	s.GroupType = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) SetDefaultSelected(v bool) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes {
	s.DefaultSelected = &v
	return s
}

type DescribeDnsGtmMonitorAvailableConfigResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmMonitorAvailableConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmMonitorAvailableConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponse) SetBody(v *DescribeDnsGtmMonitorAvailableConfigResponseBody) *DescribeDnsGtmMonitorAvailableConfigResponse {
	s.Body = v
	return s
}

type DescribeDnsGtmMonitorConfigRequest struct {
	UserClientIp    *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
}

func (s DescribeDnsGtmMonitorConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmMonitorConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorConfigRequest) SetUserClientIp(v string) *DescribeDnsGtmMonitorConfigRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigRequest) SetLang(v string) *DescribeDnsGtmMonitorConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigRequest) SetMonitorConfigId(v string) *DescribeDnsGtmMonitorConfigRequest {
	s.MonitorConfigId = &v
	return s
}

type DescribeDnsGtmMonitorConfigResponseBody struct {
	RequestId         *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Timeout           *int32                                                 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	ProtocolType      *string                                                `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	IspCityNodes      []*DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Repeated"`
	CreateTime        *string                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UpdateTime        *string                                                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	EvaluationCount   *int32                                                 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	UpdateTimestamp   *int64                                                 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	MonitorExtendInfo *string                                                `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
	MonitorConfigId   *string                                                `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	CreateTimestamp   *int64                                                 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Interval          *int32                                                 `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeDnsGtmMonitorConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmMonitorConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetRequestId(v string) *DescribeDnsGtmMonitorConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetTimeout(v int32) *DescribeDnsGtmMonitorConfigResponseBody {
	s.Timeout = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetProtocolType(v string) *DescribeDnsGtmMonitorConfigResponseBody {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetIspCityNodes(v []*DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) *DescribeDnsGtmMonitorConfigResponseBody {
	s.IspCityNodes = v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetCreateTime(v string) *DescribeDnsGtmMonitorConfigResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetUpdateTime(v string) *DescribeDnsGtmMonitorConfigResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetEvaluationCount(v int32) *DescribeDnsGtmMonitorConfigResponseBody {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetUpdateTimestamp(v int64) *DescribeDnsGtmMonitorConfigResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetMonitorExtendInfo(v string) *DescribeDnsGtmMonitorConfigResponseBody {
	s.MonitorExtendInfo = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetMonitorConfigId(v string) *DescribeDnsGtmMonitorConfigResponseBody {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetCreateTimestamp(v int64) *DescribeDnsGtmMonitorConfigResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetInterval(v int32) *DescribeDnsGtmMonitorConfigResponseBody {
	s.Interval = &v
	return s
}

type DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes struct {
	CityCode    *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	IspCode     *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	CityName    *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	IspName     *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) SetCityCode(v string) *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes {
	s.CityCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) SetCountryName(v string) *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes {
	s.CountryName = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) SetIspCode(v string) *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes {
	s.IspCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) SetCityName(v string) *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes {
	s.CityName = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) SetCountryCode(v string) *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes {
	s.CountryCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) SetIspName(v string) *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes {
	s.IspName = &v
	return s
}

type DescribeDnsGtmMonitorConfigResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsGtmMonitorConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsGtmMonitorConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsGtmMonitorConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorConfigResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmMonitorConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponse) SetBody(v *DescribeDnsGtmMonitorConfigResponseBody) *DescribeDnsGtmMonitorConfigResponse {
	s.Body = v
	return s
}

type DescribeDnsProductInstanceRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDnsProductInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsProductInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstanceRequest) SetLang(v string) *DescribeDnsProductInstanceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsProductInstanceRequest) SetUserClientIp(v string) *DescribeDnsProductInstanceRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsProductInstanceRequest) SetInstanceId(v string) *DescribeDnsProductInstanceRequest {
	s.InstanceId = &v
	return s
}

type DescribeDnsProductInstanceResponseBody struct {
	MonitorNodeCount      *int64    `json:"MonitorNodeCount,omitempty" xml:"MonitorNodeCount,omitempty"`
	InBlackHole           *bool     `json:"InBlackHole,omitempty" xml:"InBlackHole,omitempty"`
	BindDomainCount       *int64    `json:"BindDomainCount,omitempty" xml:"BindDomainCount,omitempty"`
	RegionLines           *bool     `json:"RegionLines,omitempty" xml:"RegionLines,omitempty"`
	BindCount             *int64    `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	EndTime               *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTimestamp        *int64    `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	ISPLines              *string   `json:"ISPLines,omitempty" xml:"ISPLines,omitempty"`
	EndTimestamp          *int64    `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	DnsServers            []*string `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Repeated"`
	DDosDefendQuery       *int64    `json:"DDosDefendQuery,omitempty" xml:"DDosDefendQuery,omitempty"`
	DnsSecurity           *string   `json:"DnsSecurity,omitempty" xml:"DnsSecurity,omitempty"`
	DomainType            *string   `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	URLForwardCount       *int64    `json:"URLForwardCount,omitempty" xml:"URLForwardCount,omitempty"`
	TTLMinValue           *int64    `json:"TTLMinValue,omitempty" xml:"TTLMinValue,omitempty"`
	PaymentType           *string   `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	VersionName           *string   `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	OverseaLine           *string   `json:"OverseaLine,omitempty" xml:"OverseaLine,omitempty"`
	ISPRegionLines        *string   `json:"ISPRegionLines,omitempty" xml:"ISPRegionLines,omitempty"`
	Gslb                  *bool     `json:"Gslb,omitempty" xml:"Gslb,omitempty"`
	BindUsedCount         *int64    `json:"BindUsedCount,omitempty" xml:"BindUsedCount,omitempty"`
	RequestId             *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DnsSLBCount           *int64    `json:"DnsSLBCount,omitempty" xml:"DnsSLBCount,omitempty"`
	InstanceId            *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MonitorTaskCount      *int64    `json:"MonitorTaskCount,omitempty" xml:"MonitorTaskCount,omitempty"`
	StartTime             *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DDosDefendFlow        *int64    `json:"DDosDefendFlow,omitempty" xml:"DDosDefendFlow,omitempty"`
	MonitorFrequency      *int64    `json:"MonitorFrequency,omitempty" xml:"MonitorFrequency,omitempty"`
	SearchEngineLines     *string   `json:"SearchEngineLines,omitempty" xml:"SearchEngineLines,omitempty"`
	BindDomainUsedCount   *int64    `json:"BindDomainUsedCount,omitempty" xml:"BindDomainUsedCount,omitempty"`
	VersionCode           *string   `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	OverseaDDosDefendFlow *int64    `json:"OverseaDDosDefendFlow,omitempty" xml:"OverseaDDosDefendFlow,omitempty"`
	InClean               *bool     `json:"InClean,omitempty" xml:"InClean,omitempty"`
	SubDomainLevel        *int64    `json:"SubDomainLevel,omitempty" xml:"SubDomainLevel,omitempty"`
	Domain                *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDnsProductInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsProductInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstanceResponseBody) SetMonitorNodeCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.MonitorNodeCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetInBlackHole(v bool) *DescribeDnsProductInstanceResponseBody {
	s.InBlackHole = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetBindDomainCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.BindDomainCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetRegionLines(v bool) *DescribeDnsProductInstanceResponseBody {
	s.RegionLines = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetBindCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.BindCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetEndTime(v string) *DescribeDnsProductInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetStartTimestamp(v int64) *DescribeDnsProductInstanceResponseBody {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetISPLines(v string) *DescribeDnsProductInstanceResponseBody {
	s.ISPLines = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetEndTimestamp(v int64) *DescribeDnsProductInstanceResponseBody {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetDnsServers(v []*string) *DescribeDnsProductInstanceResponseBody {
	s.DnsServers = v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetDDosDefendQuery(v int64) *DescribeDnsProductInstanceResponseBody {
	s.DDosDefendQuery = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetDnsSecurity(v string) *DescribeDnsProductInstanceResponseBody {
	s.DnsSecurity = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetDomainType(v string) *DescribeDnsProductInstanceResponseBody {
	s.DomainType = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetURLForwardCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.URLForwardCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetTTLMinValue(v int64) *DescribeDnsProductInstanceResponseBody {
	s.TTLMinValue = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetPaymentType(v string) *DescribeDnsProductInstanceResponseBody {
	s.PaymentType = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetVersionName(v string) *DescribeDnsProductInstanceResponseBody {
	s.VersionName = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetOverseaLine(v string) *DescribeDnsProductInstanceResponseBody {
	s.OverseaLine = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetISPRegionLines(v string) *DescribeDnsProductInstanceResponseBody {
	s.ISPRegionLines = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetGslb(v bool) *DescribeDnsProductInstanceResponseBody {
	s.Gslb = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetBindUsedCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.BindUsedCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetRequestId(v string) *DescribeDnsProductInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetDnsSLBCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.DnsSLBCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetInstanceId(v string) *DescribeDnsProductInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetMonitorTaskCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.MonitorTaskCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetStartTime(v string) *DescribeDnsProductInstanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetDDosDefendFlow(v int64) *DescribeDnsProductInstanceResponseBody {
	s.DDosDefendFlow = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetMonitorFrequency(v int64) *DescribeDnsProductInstanceResponseBody {
	s.MonitorFrequency = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetSearchEngineLines(v string) *DescribeDnsProductInstanceResponseBody {
	s.SearchEngineLines = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetBindDomainUsedCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.BindDomainUsedCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetVersionCode(v string) *DescribeDnsProductInstanceResponseBody {
	s.VersionCode = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetOverseaDDosDefendFlow(v int64) *DescribeDnsProductInstanceResponseBody {
	s.OverseaDDosDefendFlow = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetInClean(v bool) *DescribeDnsProductInstanceResponseBody {
	s.InClean = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetSubDomainLevel(v int64) *DescribeDnsProductInstanceResponseBody {
	s.SubDomainLevel = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetDomain(v string) *DescribeDnsProductInstanceResponseBody {
	s.Domain = &v
	return s
}

type DescribeDnsProductInstanceResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsProductInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsProductInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsProductInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstanceResponse) SetHeaders(v map[string]*string) *DescribeDnsProductInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsProductInstanceResponse) SetBody(v *DescribeDnsProductInstanceResponseBody) *DescribeDnsProductInstanceResponse {
	s.Body = v
	return s
}

type DescribeDnsProductInstancesRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VersionCode  *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	DomainType   *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
}

func (s DescribeDnsProductInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsProductInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstancesRequest) SetLang(v string) *DescribeDnsProductInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsProductInstancesRequest) SetUserClientIp(v string) *DescribeDnsProductInstancesRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsProductInstancesRequest) SetPageNumber(v int64) *DescribeDnsProductInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsProductInstancesRequest) SetPageSize(v int64) *DescribeDnsProductInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsProductInstancesRequest) SetVersionCode(v string) *DescribeDnsProductInstancesRequest {
	s.VersionCode = &v
	return s
}

func (s *DescribeDnsProductInstancesRequest) SetDomainType(v string) *DescribeDnsProductInstancesRequest {
	s.DomainType = &v
	return s
}

type DescribeDnsProductInstancesResponseBody struct {
	TotalCount  *int64                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	DomainType  *string                                               `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	PageSize    *int64                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber  *int64                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DnsProducts []*DescribeDnsProductInstancesResponseBodyDnsProducts `json:"DnsProducts,omitempty" xml:"DnsProducts,omitempty" type:"Repeated"`
}

func (s DescribeDnsProductInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsProductInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstancesResponseBody) SetTotalCount(v int64) *DescribeDnsProductInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBody) SetDomainType(v string) *DescribeDnsProductInstancesResponseBody {
	s.DomainType = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBody) SetPageSize(v int64) *DescribeDnsProductInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBody) SetRequestId(v string) *DescribeDnsProductInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBody) SetPageNumber(v int64) *DescribeDnsProductInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBody) SetDnsProducts(v []*DescribeDnsProductInstancesResponseBodyDnsProducts) *DescribeDnsProductInstancesResponseBody {
	s.DnsProducts = v
	return s
}

type DescribeDnsProductInstancesResponseBodyDnsProducts struct {
	OverseaLine           *string `json:"OverseaLine,omitempty" xml:"OverseaLine,omitempty"`
	PaymentType           *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	MonitorNodeCount      *int64  `json:"MonitorNodeCount,omitempty" xml:"MonitorNodeCount,omitempty"`
	InBlackHole           *bool   `json:"InBlackHole,omitempty" xml:"InBlackHole,omitempty"`
	BindDomainUsedCount   *int64  `json:"BindDomainUsedCount,omitempty" xml:"BindDomainUsedCount,omitempty"`
	ISPRegionLines        *string `json:"ISPRegionLines,omitempty" xml:"ISPRegionLines,omitempty"`
	TTLMinValue           *int64  `json:"TTLMinValue,omitempty" xml:"TTLMinValue,omitempty"`
	ISPLines              *string `json:"ISPLines,omitempty" xml:"ISPLines,omitempty"`
	SearchEngineLines     *string `json:"SearchEngineLines,omitempty" xml:"SearchEngineLines,omitempty"`
	EndTimestamp          *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	VersionName           *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	VersionCode           *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	MonitorTaskCount      *int64  `json:"MonitorTaskCount,omitempty" xml:"MonitorTaskCount,omitempty"`
	BindUsedCount         *int64  `json:"BindUsedCount,omitempty" xml:"BindUsedCount,omitempty"`
	Domain                *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	MonitorFrequency      *int64  `json:"MonitorFrequency,omitempty" xml:"MonitorFrequency,omitempty"`
	InClean               *bool   `json:"InClean,omitempty" xml:"InClean,omitempty"`
	URLForwardCount       *int64  `json:"URLForwardCount,omitempty" xml:"URLForwardCount,omitempty"`
	StartTimestamp        *int64  `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	DDosDefendQuery       *int64  `json:"DDosDefendQuery,omitempty" xml:"DDosDefendQuery,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DDosDefendFlow        *int64  `json:"DDosDefendFlow,omitempty" xml:"DDosDefendFlow,omitempty"`
	BindCount             *int64  `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	SubDomainLevel        *int64  `json:"SubDomainLevel,omitempty" xml:"SubDomainLevel,omitempty"`
	BindDomainCount       *int64  `json:"BindDomainCount,omitempty" xml:"BindDomainCount,omitempty"`
	EndTime               *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime             *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	OverseaDDosDefendFlow *int64  `json:"OverseaDDosDefendFlow,omitempty" xml:"OverseaDDosDefendFlow,omitempty"`
	RegionLines           *bool   `json:"RegionLines,omitempty" xml:"RegionLines,omitempty"`
	Gslb                  *bool   `json:"Gslb,omitempty" xml:"Gslb,omitempty"`
	DnsSecurity           *string `json:"DnsSecurity,omitempty" xml:"DnsSecurity,omitempty"`
	DnsSLBCount           *int64  `json:"DnsSLBCount,omitempty" xml:"DnsSLBCount,omitempty"`
}

func (s DescribeDnsProductInstancesResponseBodyDnsProducts) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsProductInstancesResponseBodyDnsProducts) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetOverseaLine(v string) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.OverseaLine = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetPaymentType(v string) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.PaymentType = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetMonitorNodeCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.MonitorNodeCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetInBlackHole(v bool) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.InBlackHole = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetBindDomainUsedCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.BindDomainUsedCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetISPRegionLines(v string) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.ISPRegionLines = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetTTLMinValue(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.TTLMinValue = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetISPLines(v string) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.ISPLines = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetSearchEngineLines(v string) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.SearchEngineLines = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetEndTimestamp(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetVersionName(v string) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.VersionName = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetVersionCode(v string) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.VersionCode = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetMonitorTaskCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.MonitorTaskCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetBindUsedCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.BindUsedCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetDomain(v string) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.Domain = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetMonitorFrequency(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.MonitorFrequency = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetInClean(v bool) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.InClean = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetURLForwardCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.URLForwardCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetStartTimestamp(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetDDosDefendQuery(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.DDosDefendQuery = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetInstanceId(v string) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetDDosDefendFlow(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.DDosDefendFlow = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetBindCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.BindCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetSubDomainLevel(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.SubDomainLevel = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetBindDomainCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.BindDomainCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetEndTime(v string) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.EndTime = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetStartTime(v string) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.StartTime = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetOverseaDDosDefendFlow(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.OverseaDDosDefendFlow = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetRegionLines(v bool) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.RegionLines = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetGslb(v bool) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.Gslb = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetDnsSecurity(v string) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.DnsSecurity = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetDnsSLBCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.DnsSLBCount = &v
	return s
}

type DescribeDnsProductInstancesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnsProductInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnsProductInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnsProductInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstancesResponse) SetHeaders(v map[string]*string) *DescribeDnsProductInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsProductInstancesResponse) SetBody(v *DescribeDnsProductInstancesResponseBody) *DescribeDnsProductInstancesResponse {
	s.Body = v
	return s
}

type DescribeDNSSLBSubDomainsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDNSSLBSubDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDNSSLBSubDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDNSSLBSubDomainsRequest) SetLang(v string) *DescribeDNSSLBSubDomainsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsRequest) SetUserClientIp(v string) *DescribeDNSSLBSubDomainsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsRequest) SetDomainName(v string) *DescribeDNSSLBSubDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsRequest) SetPageNumber(v int64) *DescribeDNSSLBSubDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsRequest) SetPageSize(v int64) *DescribeDNSSLBSubDomainsRequest {
	s.PageSize = &v
	return s
}

type DescribeDNSSLBSubDomainsResponseBody struct {
	TotalCount    *int64                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize      *int64                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber    *int64                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	SlbSubDomains []*DescribeDNSSLBSubDomainsResponseBodySlbSubDomains `json:"SlbSubDomains,omitempty" xml:"SlbSubDomains,omitempty" type:"Repeated"`
}

func (s DescribeDNSSLBSubDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDNSSLBSubDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDNSSLBSubDomainsResponseBody) SetTotalCount(v int64) *DescribeDNSSLBSubDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBody) SetPageSize(v int64) *DescribeDNSSLBSubDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBody) SetRequestId(v string) *DescribeDNSSLBSubDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBody) SetPageNumber(v int64) *DescribeDNSSLBSubDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBody) SetSlbSubDomains(v []*DescribeDNSSLBSubDomainsResponseBodySlbSubDomains) *DescribeDNSSLBSubDomainsResponseBody {
	s.SlbSubDomains = v
	return s
}

type DescribeDNSSLBSubDomainsResponseBodySlbSubDomains struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	RecordCount *int64  `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	Open        *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	SubDomain   *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s DescribeDNSSLBSubDomainsResponseBodySlbSubDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDNSSLBSubDomainsResponseBodySlbSubDomains) GoString() string {
	return s.String()
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains) SetType(v string) *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains {
	s.Type = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains) SetRecordCount(v int64) *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains {
	s.RecordCount = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains) SetOpen(v bool) *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains {
	s.Open = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains) SetSubDomain(v string) *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains {
	s.SubDomain = &v
	return s
}

type DescribeDNSSLBSubDomainsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDNSSLBSubDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDNSSLBSubDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDNSSLBSubDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDNSSLBSubDomainsResponse) SetHeaders(v map[string]*string) *DescribeDNSSLBSubDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponse) SetBody(v *DescribeDNSSLBSubDomainsResponseBody) *DescribeDNSSLBSubDomainsResponse {
	s.Body = v
	return s
}

type DescribeDohAccountStatisticsRequest struct {
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s DescribeDohAccountStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohAccountStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDohAccountStatisticsRequest) SetLang(v string) *DescribeDohAccountStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDohAccountStatisticsRequest) SetStartDate(v string) *DescribeDohAccountStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDohAccountStatisticsRequest) SetEndDate(v string) *DescribeDohAccountStatisticsRequest {
	s.EndDate = &v
	return s
}

type DescribeDohAccountStatisticsResponseBody struct {
	RequestId  *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistics []*DescribeDohAccountStatisticsResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeDohAccountStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohAccountStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDohAccountStatisticsResponseBody) SetRequestId(v string) *DescribeDohAccountStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDohAccountStatisticsResponseBody) SetStatistics(v []*DescribeDohAccountStatisticsResponseBodyStatistics) *DescribeDohAccountStatisticsResponseBody {
	s.Statistics = v
	return s
}

type DescribeDohAccountStatisticsResponseBodyStatistics struct {
	V6HttpCount  *int64 `json:"V6HttpCount,omitempty" xml:"V6HttpCount,omitempty"`
	V4HttpsCount *int64 `json:"V4HttpsCount,omitempty" xml:"V4HttpsCount,omitempty"`
	Timestamp    *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TotalCount   *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	V4HttpCount  *int64 `json:"V4HttpCount,omitempty" xml:"V4HttpCount,omitempty"`
	V6HttpsCount *int64 `json:"V6HttpsCount,omitempty" xml:"V6HttpsCount,omitempty"`
}

func (s DescribeDohAccountStatisticsResponseBodyStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohAccountStatisticsResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) SetV6HttpCount(v int64) *DescribeDohAccountStatisticsResponseBodyStatistics {
	s.V6HttpCount = &v
	return s
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) SetV4HttpsCount(v int64) *DescribeDohAccountStatisticsResponseBodyStatistics {
	s.V4HttpsCount = &v
	return s
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) SetTimestamp(v int64) *DescribeDohAccountStatisticsResponseBodyStatistics {
	s.Timestamp = &v
	return s
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) SetTotalCount(v int64) *DescribeDohAccountStatisticsResponseBodyStatistics {
	s.TotalCount = &v
	return s
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) SetV4HttpCount(v int64) *DescribeDohAccountStatisticsResponseBodyStatistics {
	s.V4HttpCount = &v
	return s
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) SetV6HttpsCount(v int64) *DescribeDohAccountStatisticsResponseBodyStatistics {
	s.V6HttpsCount = &v
	return s
}

type DescribeDohAccountStatisticsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDohAccountStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDohAccountStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohAccountStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDohAccountStatisticsResponse) SetHeaders(v map[string]*string) *DescribeDohAccountStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDohAccountStatisticsResponse) SetBody(v *DescribeDohAccountStatisticsResponseBody) *DescribeDohAccountStatisticsResponse {
	s.Body = v
	return s
}

type DescribeDohDomainStatisticsRequest struct {
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s DescribeDohDomainStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohDomainStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsRequest) SetLang(v string) *DescribeDohDomainStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDohDomainStatisticsRequest) SetDomainName(v string) *DescribeDohDomainStatisticsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDohDomainStatisticsRequest) SetStartDate(v string) *DescribeDohDomainStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDohDomainStatisticsRequest) SetEndDate(v string) *DescribeDohDomainStatisticsRequest {
	s.EndDate = &v
	return s
}

type DescribeDohDomainStatisticsResponseBody struct {
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistics []*DescribeDohDomainStatisticsResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeDohDomainStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohDomainStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsResponseBody) SetRequestId(v string) *DescribeDohDomainStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDohDomainStatisticsResponseBody) SetStatistics(v []*DescribeDohDomainStatisticsResponseBodyStatistics) *DescribeDohDomainStatisticsResponseBody {
	s.Statistics = v
	return s
}

type DescribeDohDomainStatisticsResponseBodyStatistics struct {
	V6HttpCount  *int64 `json:"V6HttpCount,omitempty" xml:"V6HttpCount,omitempty"`
	V4HttpsCount *int64 `json:"V4HttpsCount,omitempty" xml:"V4HttpsCount,omitempty"`
	Timestamp    *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TotalCount   *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	V4HttpCount  *int64 `json:"V4HttpCount,omitempty" xml:"V4HttpCount,omitempty"`
	V6HttpsCount *int64 `json:"V6HttpsCount,omitempty" xml:"V6HttpsCount,omitempty"`
}

func (s DescribeDohDomainStatisticsResponseBodyStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohDomainStatisticsResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) SetV6HttpCount(v int64) *DescribeDohDomainStatisticsResponseBodyStatistics {
	s.V6HttpCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) SetV4HttpsCount(v int64) *DescribeDohDomainStatisticsResponseBodyStatistics {
	s.V4HttpsCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) SetTimestamp(v int64) *DescribeDohDomainStatisticsResponseBodyStatistics {
	s.Timestamp = &v
	return s
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) SetTotalCount(v int64) *DescribeDohDomainStatisticsResponseBodyStatistics {
	s.TotalCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) SetV4HttpCount(v int64) *DescribeDohDomainStatisticsResponseBodyStatistics {
	s.V4HttpCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) SetV6HttpsCount(v int64) *DescribeDohDomainStatisticsResponseBodyStatistics {
	s.V6HttpsCount = &v
	return s
}

type DescribeDohDomainStatisticsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDohDomainStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDohDomainStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohDomainStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsResponse) SetHeaders(v map[string]*string) *DescribeDohDomainStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDohDomainStatisticsResponse) SetBody(v *DescribeDohDomainStatisticsResponseBody) *DescribeDohDomainStatisticsResponse {
	s.Body = v
	return s
}

type DescribeDohDomainStatisticsSummaryRequest struct {
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Direction  *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDohDomainStatisticsSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohDomainStatisticsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsSummaryRequest) SetLang(v string) *DescribeDohDomainStatisticsSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryRequest) SetPageNumber(v int32) *DescribeDohDomainStatisticsSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryRequest) SetPageSize(v int32) *DescribeDohDomainStatisticsSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryRequest) SetStartDate(v string) *DescribeDohDomainStatisticsSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryRequest) SetEndDate(v string) *DescribeDohDomainStatisticsSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryRequest) SetOrderBy(v string) *DescribeDohDomainStatisticsSummaryRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryRequest) SetDirection(v string) *DescribeDohDomainStatisticsSummaryRequest {
	s.Direction = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryRequest) SetDomainName(v string) *DescribeDohDomainStatisticsSummaryRequest {
	s.DomainName = &v
	return s
}

type DescribeDohDomainStatisticsSummaryResponseBody struct {
	PageSize   *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalPages *int32                                                      `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems *int32                                                      `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	Statistics []*DescribeDohDomainStatisticsSummaryResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeDohDomainStatisticsSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohDomainStatisticsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) SetPageSize(v int32) *DescribeDohDomainStatisticsSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) SetRequestId(v string) *DescribeDohDomainStatisticsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) SetPageNumber(v int32) *DescribeDohDomainStatisticsSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) SetTotalPages(v int32) *DescribeDohDomainStatisticsSummaryResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) SetTotalItems(v int32) *DescribeDohDomainStatisticsSummaryResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) SetStatistics(v []*DescribeDohDomainStatisticsSummaryResponseBodyStatistics) *DescribeDohDomainStatisticsSummaryResponseBody {
	s.Statistics = v
	return s
}

type DescribeDohDomainStatisticsSummaryResponseBodyStatistics struct {
	V6HttpCount  *int64  `json:"V6HttpCount,omitempty" xml:"V6HttpCount,omitempty"`
	V4HttpsCount *int64  `json:"V4HttpsCount,omitempty" xml:"V4HttpsCount,omitempty"`
	IpCount      *int64  `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	TotalCount   *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	HttpCount    *int64  `json:"HttpCount,omitempty" xml:"HttpCount,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	HttpsCount   *int64  `json:"HttpsCount,omitempty" xml:"HttpsCount,omitempty"`
	V4HttpCount  *int64  `json:"V4HttpCount,omitempty" xml:"V4HttpCount,omitempty"`
	V6HttpsCount *int64  `json:"V6HttpsCount,omitempty" xml:"V6HttpsCount,omitempty"`
}

func (s DescribeDohDomainStatisticsSummaryResponseBodyStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohDomainStatisticsSummaryResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetV6HttpCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.V6HttpCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetV4HttpsCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.V4HttpsCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetIpCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.IpCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetTotalCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.TotalCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetHttpCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.HttpCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetDomainName(v string) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.DomainName = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetHttpsCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.HttpsCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetV4HttpCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.V4HttpCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetV6HttpsCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.V6HttpsCount = &v
	return s
}

type DescribeDohDomainStatisticsSummaryResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDohDomainStatisticsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDohDomainStatisticsSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohDomainStatisticsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsSummaryResponse) SetHeaders(v map[string]*string) *DescribeDohDomainStatisticsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponse) SetBody(v *DescribeDohDomainStatisticsSummaryResponseBody) *DescribeDohDomainStatisticsSummaryResponse {
	s.Body = v
	return s
}

type DescribeDohSubDomainStatisticsRequest struct {
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s DescribeDohSubDomainStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsRequest) SetLang(v string) *DescribeDohSubDomainStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsRequest) SetSubDomain(v string) *DescribeDohSubDomainStatisticsRequest {
	s.SubDomain = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsRequest) SetStartDate(v string) *DescribeDohSubDomainStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsRequest) SetEndDate(v string) *DescribeDohSubDomainStatisticsRequest {
	s.EndDate = &v
	return s
}

type DescribeDohSubDomainStatisticsResponseBody struct {
	RequestId  *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistics []*DescribeDohSubDomainStatisticsResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeDohSubDomainStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsResponseBody) SetRequestId(v string) *DescribeDohSubDomainStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponseBody) SetStatistics(v []*DescribeDohSubDomainStatisticsResponseBodyStatistics) *DescribeDohSubDomainStatisticsResponseBody {
	s.Statistics = v
	return s
}

type DescribeDohSubDomainStatisticsResponseBodyStatistics struct {
	V6HttpCount  *int64 `json:"V6HttpCount,omitempty" xml:"V6HttpCount,omitempty"`
	V4HttpsCount *int64 `json:"V4HttpsCount,omitempty" xml:"V4HttpsCount,omitempty"`
	Timestamp    *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TotalCount   *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	V4HttpCount  *int64 `json:"V4HttpCount,omitempty" xml:"V4HttpCount,omitempty"`
	V6HttpsCount *int64 `json:"V6HttpsCount,omitempty" xml:"V6HttpsCount,omitempty"`
}

func (s DescribeDohSubDomainStatisticsResponseBodyStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) SetV6HttpCount(v int64) *DescribeDohSubDomainStatisticsResponseBodyStatistics {
	s.V6HttpCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) SetV4HttpsCount(v int64) *DescribeDohSubDomainStatisticsResponseBodyStatistics {
	s.V4HttpsCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) SetTimestamp(v int64) *DescribeDohSubDomainStatisticsResponseBodyStatistics {
	s.Timestamp = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) SetTotalCount(v int64) *DescribeDohSubDomainStatisticsResponseBodyStatistics {
	s.TotalCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) SetV4HttpCount(v int64) *DescribeDohSubDomainStatisticsResponseBodyStatistics {
	s.V4HttpCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) SetV6HttpsCount(v int64) *DescribeDohSubDomainStatisticsResponseBodyStatistics {
	s.V6HttpsCount = &v
	return s
}

type DescribeDohSubDomainStatisticsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDohSubDomainStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDohSubDomainStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsResponse) SetHeaders(v map[string]*string) *DescribeDohSubDomainStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponse) SetBody(v *DescribeDohSubDomainStatisticsResponseBody) *DescribeDohSubDomainStatisticsResponse {
	s.Body = v
	return s
}

type DescribeDohSubDomainStatisticsSummaryRequest struct {
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Direction  *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	SubDomain  *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDohSubDomainStatisticsSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetLang(v string) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetPageNumber(v int32) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetPageSize(v int32) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetStartDate(v string) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetEndDate(v string) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetOrderBy(v string) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetDirection(v string) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.Direction = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetSubDomain(v string) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.SubDomain = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetDomainName(v string) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.DomainName = &v
	return s
}

type DescribeDohSubDomainStatisticsSummaryResponseBody struct {
	PageSize   *int32                                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalPages *int32                                                         `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems *int32                                                         `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	Statistics []*DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeDohSubDomainStatisticsSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) SetPageSize(v int32) *DescribeDohSubDomainStatisticsSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) SetRequestId(v string) *DescribeDohSubDomainStatisticsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) SetPageNumber(v int32) *DescribeDohSubDomainStatisticsSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) SetTotalPages(v int32) *DescribeDohSubDomainStatisticsSummaryResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) SetTotalItems(v int32) *DescribeDohSubDomainStatisticsSummaryResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) SetStatistics(v []*DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) *DescribeDohSubDomainStatisticsSummaryResponseBody {
	s.Statistics = v
	return s
}

type DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics struct {
	V6HttpCount  *int64  `json:"V6HttpCount,omitempty" xml:"V6HttpCount,omitempty"`
	V4HttpsCount *int64  `json:"V4HttpsCount,omitempty" xml:"V4HttpsCount,omitempty"`
	IpCount      *int64  `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	SubDomain    *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	TotalCount   *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	HttpCount    *int64  `json:"HttpCount,omitempty" xml:"HttpCount,omitempty"`
	HttpsCount   *int64  `json:"HttpsCount,omitempty" xml:"HttpsCount,omitempty"`
	V4HttpCount  *int64  `json:"V4HttpCount,omitempty" xml:"V4HttpCount,omitempty"`
	V6HttpsCount *int64  `json:"V6HttpsCount,omitempty" xml:"V6HttpsCount,omitempty"`
}

func (s DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetV6HttpCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.V6HttpCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetV4HttpsCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.V4HttpsCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetIpCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.IpCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetSubDomain(v string) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.SubDomain = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetTotalCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.TotalCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetHttpCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.HttpCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetHttpsCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.HttpsCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetV4HttpCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.V4HttpCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetV6HttpsCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.V6HttpsCount = &v
	return s
}

type DescribeDohSubDomainStatisticsSummaryResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDohSubDomainStatisticsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDohSubDomainStatisticsSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsSummaryResponse) SetHeaders(v map[string]*string) *DescribeDohSubDomainStatisticsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponse) SetBody(v *DescribeDohSubDomainStatisticsSummaryResponseBody) *DescribeDohSubDomainStatisticsSummaryResponse {
	s.Body = v
	return s
}

type DescribeDohUserInfoRequest struct {
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s DescribeDohUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohUserInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDohUserInfoRequest) SetLang(v string) *DescribeDohUserInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDohUserInfoRequest) SetStartDate(v string) *DescribeDohUserInfoRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDohUserInfoRequest) SetEndDate(v string) *DescribeDohUserInfoRequest {
	s.EndDate = &v
	return s
}

type DescribeDohUserInfoResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubDomainCount *int32  `json:"SubDomainCount,omitempty" xml:"SubDomainCount,omitempty"`
	PdnsId         *int64  `json:"PdnsId,omitempty" xml:"PdnsId,omitempty"`
	DomainCount    *int32  `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
}

func (s DescribeDohUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDohUserInfoResponseBody) SetRequestId(v string) *DescribeDohUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDohUserInfoResponseBody) SetSubDomainCount(v int32) *DescribeDohUserInfoResponseBody {
	s.SubDomainCount = &v
	return s
}

func (s *DescribeDohUserInfoResponseBody) SetPdnsId(v int64) *DescribeDohUserInfoResponseBody {
	s.PdnsId = &v
	return s
}

func (s *DescribeDohUserInfoResponseBody) SetDomainCount(v int32) *DescribeDohUserInfoResponseBody {
	s.DomainCount = &v
	return s
}

type DescribeDohUserInfoResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDohUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDohUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDohUserInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDohUserInfoResponse) SetHeaders(v map[string]*string) *DescribeDohUserInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDohUserInfoResponse) SetBody(v *DescribeDohUserInfoResponseBody) *DescribeDohUserInfoResponse {
	s.Body = v
	return s
}

type DescribeDomainDnssecInfoRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDomainDnssecInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDnssecInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainDnssecInfoRequest) SetLang(v string) *DescribeDomainDnssecInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainDnssecInfoRequest) SetUserClientIp(v string) *DescribeDomainDnssecInfoRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDomainDnssecInfoRequest) SetDomainName(v string) *DescribeDomainDnssecInfoRequest {
	s.DomainName = &v
	return s
}

type DescribeDomainDnssecInfoResponseBody struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Digest     *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	PublicKey  *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	DigestType *string `json:"DigestType,omitempty" xml:"DigestType,omitempty"`
	DsRecord   *string `json:"DsRecord,omitempty" xml:"DsRecord,omitempty"`
	KeyTag     *string `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	Flags      *string `json:"Flags,omitempty" xml:"Flags,omitempty"`
	Algorithm  *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
}

func (s DescribeDomainDnssecInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDnssecInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainDnssecInfoResponseBody) SetStatus(v string) *DescribeDomainDnssecInfoResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetRequestId(v string) *DescribeDomainDnssecInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetDigest(v string) *DescribeDomainDnssecInfoResponseBody {
	s.Digest = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetDomainName(v string) *DescribeDomainDnssecInfoResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetPublicKey(v string) *DescribeDomainDnssecInfoResponseBody {
	s.PublicKey = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetDigestType(v string) *DescribeDomainDnssecInfoResponseBody {
	s.DigestType = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetDsRecord(v string) *DescribeDomainDnssecInfoResponseBody {
	s.DsRecord = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetKeyTag(v string) *DescribeDomainDnssecInfoResponseBody {
	s.KeyTag = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetFlags(v string) *DescribeDomainDnssecInfoResponseBody {
	s.Flags = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetAlgorithm(v string) *DescribeDomainDnssecInfoResponseBody {
	s.Algorithm = &v
	return s
}

type DescribeDomainDnssecInfoResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainDnssecInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainDnssecInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDnssecInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainDnssecInfoResponse) SetHeaders(v map[string]*string) *DescribeDomainDnssecInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainDnssecInfoResponse) SetBody(v *DescribeDomainDnssecInfoResponseBody) *DescribeDomainDnssecInfoResponse {
	s.Body = v
	return s
}

type DescribeDomainGroupsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	KeyWord      *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDomainGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainGroupsRequest) SetLang(v string) *DescribeDomainGroupsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainGroupsRequest) SetUserClientIp(v string) *DescribeDomainGroupsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDomainGroupsRequest) SetKeyWord(v string) *DescribeDomainGroupsRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeDomainGroupsRequest) SetPageNumber(v int64) *DescribeDomainGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainGroupsRequest) SetPageSize(v int64) *DescribeDomainGroupsRequest {
	s.PageSize = &v
	return s
}

type DescribeDomainGroupsResponseBody struct {
	DomainGroups []*DescribeDomainGroupsResponseBodyDomainGroups `json:"DomainGroups,omitempty" xml:"DomainGroups,omitempty" type:"Repeated"`
	TotalCount   *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize     *int64                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber   *int64                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeDomainGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainGroupsResponseBody) SetDomainGroups(v []*DescribeDomainGroupsResponseBodyDomainGroups) *DescribeDomainGroupsResponseBody {
	s.DomainGroups = v
	return s
}

func (s *DescribeDomainGroupsResponseBody) SetTotalCount(v int64) *DescribeDomainGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainGroupsResponseBody) SetPageSize(v int64) *DescribeDomainGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainGroupsResponseBody) SetRequestId(v string) *DescribeDomainGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainGroupsResponseBody) SetPageNumber(v int64) *DescribeDomainGroupsResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeDomainGroupsResponseBodyDomainGroups struct {
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	DomainCount *int64  `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
}

func (s DescribeDomainGroupsResponseBodyDomainGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainGroupsResponseBodyDomainGroups) GoString() string {
	return s.String()
}

func (s *DescribeDomainGroupsResponseBodyDomainGroups) SetGroupId(v string) *DescribeDomainGroupsResponseBodyDomainGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainGroupsResponseBodyDomainGroups) SetGroupName(v string) *DescribeDomainGroupsResponseBodyDomainGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeDomainGroupsResponseBodyDomainGroups) SetDomainCount(v int64) *DescribeDomainGroupsResponseBodyDomainGroups {
	s.DomainCount = &v
	return s
}

type DescribeDomainGroupsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainGroupsResponse) SetHeaders(v map[string]*string) *DescribeDomainGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainGroupsResponse) SetBody(v *DescribeDomainGroupsResponseBody) *DescribeDomainGroupsResponse {
	s.Body = v
	return s
}

type DescribeDomainInfoRequest struct {
	Lang                 *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp         *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName           *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	NeedDetailAttributes *bool   `json:"NeedDetailAttributes,omitempty" xml:"NeedDetailAttributes,omitempty"`
}

func (s DescribeDomainInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainInfoRequest) SetLang(v string) *DescribeDomainInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainInfoRequest) SetUserClientIp(v string) *DescribeDomainInfoRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDomainInfoRequest) SetDomainName(v string) *DescribeDomainInfoRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainInfoRequest) SetNeedDetailAttributes(v bool) *DescribeDomainInfoRequest {
	s.NeedDetailAttributes = &v
	return s
}

type DescribeDomainInfoResponseBody struct {
	RecordLineTreeJson *string                                      `json:"RecordLineTreeJson,omitempty" xml:"RecordLineTreeJson,omitempty"`
	GroupName          *string                                      `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InBlackHole        *bool                                        `json:"InBlackHole,omitempty" xml:"InBlackHole,omitempty"`
	RegionLines        *bool                                        `json:"RegionLines,omitempty" xml:"RegionLines,omitempty"`
	SlaveDns           *bool                                        `json:"SlaveDns,omitempty" xml:"SlaveDns,omitempty"`
	AliDomain          *bool                                        `json:"AliDomain,omitempty" xml:"AliDomain,omitempty"`
	RequestId          *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId    *string                                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceId         *string                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DomainName         *string                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	CreateTime         *string                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PunyCode           *string                                      `json:"PunyCode,omitempty" xml:"PunyCode,omitempty"`
	DnsServers         []*string                                    `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Repeated"`
	Remark             *string                                      `json:"Remark,omitempty" xml:"Remark,omitempty"`
	GroupId            *string                                      `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	VersionCode        *string                                      `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	RecordLines        []*DescribeDomainInfoResponseBodyRecordLines `json:"RecordLines,omitempty" xml:"RecordLines,omitempty" type:"Repeated"`
	DomainId           *string                                      `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	AvailableTtls      []*string                                    `json:"AvailableTtls,omitempty" xml:"AvailableTtls,omitempty" type:"Repeated"`
	MinTtl             *int64                                       `json:"MinTtl,omitempty" xml:"MinTtl,omitempty"`
	InClean            *bool                                        `json:"InClean,omitempty" xml:"InClean,omitempty"`
	VersionName        *string                                      `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	LineType           *string                                      `json:"LineType,omitempty" xml:"LineType,omitempty"`
}

func (s DescribeDomainInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainInfoResponseBody) SetRecordLineTreeJson(v string) *DescribeDomainInfoResponseBody {
	s.RecordLineTreeJson = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetGroupName(v string) *DescribeDomainInfoResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetInBlackHole(v bool) *DescribeDomainInfoResponseBody {
	s.InBlackHole = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetRegionLines(v bool) *DescribeDomainInfoResponseBody {
	s.RegionLines = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetSlaveDns(v bool) *DescribeDomainInfoResponseBody {
	s.SlaveDns = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetAliDomain(v bool) *DescribeDomainInfoResponseBody {
	s.AliDomain = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetRequestId(v string) *DescribeDomainInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetResourceGroupId(v string) *DescribeDomainInfoResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetInstanceId(v string) *DescribeDomainInfoResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetDomainName(v string) *DescribeDomainInfoResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetCreateTime(v string) *DescribeDomainInfoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetPunyCode(v string) *DescribeDomainInfoResponseBody {
	s.PunyCode = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetDnsServers(v []*string) *DescribeDomainInfoResponseBody {
	s.DnsServers = v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetRemark(v string) *DescribeDomainInfoResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetGroupId(v string) *DescribeDomainInfoResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetVersionCode(v string) *DescribeDomainInfoResponseBody {
	s.VersionCode = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetRecordLines(v []*DescribeDomainInfoResponseBodyRecordLines) *DescribeDomainInfoResponseBody {
	s.RecordLines = v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetDomainId(v string) *DescribeDomainInfoResponseBody {
	s.DomainId = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetAvailableTtls(v []*string) *DescribeDomainInfoResponseBody {
	s.AvailableTtls = v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetMinTtl(v int64) *DescribeDomainInfoResponseBody {
	s.MinTtl = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetInClean(v bool) *DescribeDomainInfoResponseBody {
	s.InClean = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetVersionName(v string) *DescribeDomainInfoResponseBody {
	s.VersionName = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetLineType(v string) *DescribeDomainInfoResponseBody {
	s.LineType = &v
	return s
}

type DescribeDomainInfoResponseBodyRecordLines struct {
	FatherCode      *string `json:"FatherCode,omitempty" xml:"FatherCode,omitempty"`
	LineDisplayName *string `json:"LineDisplayName,omitempty" xml:"LineDisplayName,omitempty"`
	LineCode        *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	LineName        *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
}

func (s DescribeDomainInfoResponseBodyRecordLines) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainInfoResponseBodyRecordLines) GoString() string {
	return s.String()
}

func (s *DescribeDomainInfoResponseBodyRecordLines) SetFatherCode(v string) *DescribeDomainInfoResponseBodyRecordLines {
	s.FatherCode = &v
	return s
}

func (s *DescribeDomainInfoResponseBodyRecordLines) SetLineDisplayName(v string) *DescribeDomainInfoResponseBodyRecordLines {
	s.LineDisplayName = &v
	return s
}

func (s *DescribeDomainInfoResponseBodyRecordLines) SetLineCode(v string) *DescribeDomainInfoResponseBodyRecordLines {
	s.LineCode = &v
	return s
}

func (s *DescribeDomainInfoResponseBodyRecordLines) SetLineName(v string) *DescribeDomainInfoResponseBodyRecordLines {
	s.LineName = &v
	return s
}

type DescribeDomainInfoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainInfoResponse) SetHeaders(v map[string]*string) *DescribeDomainInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainInfoResponse) SetBody(v *DescribeDomainInfoResponseBody) *DescribeDomainInfoResponse {
	s.Body = v
	return s
}

type DescribeDomainLogsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	KeyWord      *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate      *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDomainLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainLogsRequest) SetLang(v string) *DescribeDomainLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetUserClientIp(v string) *DescribeDomainLogsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetKeyWord(v string) *DescribeDomainLogsRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetGroupId(v string) *DescribeDomainLogsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetPageNumber(v int64) *DescribeDomainLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetPageSize(v int64) *DescribeDomainLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetStartDate(v string) *DescribeDomainLogsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetEndDate(v string) *DescribeDomainLogsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetType(v string) *DescribeDomainLogsRequest {
	s.Type = &v
	return s
}

type DescribeDomainLogsResponseBody struct {
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int64                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int64                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DomainLogs []*DescribeDomainLogsResponseBodyDomainLogs `json:"DomainLogs,omitempty" xml:"DomainLogs,omitempty" type:"Repeated"`
}

func (s DescribeDomainLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainLogsResponseBody) SetTotalCount(v int64) *DescribeDomainLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainLogsResponseBody) SetPageSize(v int64) *DescribeDomainLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainLogsResponseBody) SetRequestId(v string) *DescribeDomainLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainLogsResponseBody) SetPageNumber(v int64) *DescribeDomainLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainLogsResponseBody) SetDomainLogs(v []*DescribeDomainLogsResponseBodyDomainLogs) *DescribeDomainLogsResponseBody {
	s.DomainLogs = v
	return s
}

type DescribeDomainLogsResponseBodyDomainLogs struct {
	Action          *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionTimestamp *int64  `json:"ActionTimestamp,omitempty" xml:"ActionTimestamp,omitempty"`
	ZoneId          *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ClientIp        *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ActionTime      *string `json:"ActionTime,omitempty" xml:"ActionTime,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDomainLogsResponseBodyDomainLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainLogsResponseBodyDomainLogs) GoString() string {
	return s.String()
}

func (s *DescribeDomainLogsResponseBodyDomainLogs) SetAction(v string) *DescribeDomainLogsResponseBodyDomainLogs {
	s.Action = &v
	return s
}

func (s *DescribeDomainLogsResponseBodyDomainLogs) SetActionTimestamp(v int64) *DescribeDomainLogsResponseBodyDomainLogs {
	s.ActionTimestamp = &v
	return s
}

func (s *DescribeDomainLogsResponseBodyDomainLogs) SetZoneId(v string) *DescribeDomainLogsResponseBodyDomainLogs {
	s.ZoneId = &v
	return s
}

func (s *DescribeDomainLogsResponseBodyDomainLogs) SetClientIp(v string) *DescribeDomainLogsResponseBodyDomainLogs {
	s.ClientIp = &v
	return s
}

func (s *DescribeDomainLogsResponseBodyDomainLogs) SetMessage(v string) *DescribeDomainLogsResponseBodyDomainLogs {
	s.Message = &v
	return s
}

func (s *DescribeDomainLogsResponseBodyDomainLogs) SetActionTime(v string) *DescribeDomainLogsResponseBodyDomainLogs {
	s.ActionTime = &v
	return s
}

func (s *DescribeDomainLogsResponseBodyDomainLogs) SetDomainName(v string) *DescribeDomainLogsResponseBodyDomainLogs {
	s.DomainName = &v
	return s
}

type DescribeDomainLogsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainLogsResponse) SetHeaders(v map[string]*string) *DescribeDomainLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainLogsResponse) SetBody(v *DescribeDomainLogsResponseBody) *DescribeDomainLogsResponse {
	s.Body = v
	return s
}

type DescribeDomainNsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainType   *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
}

func (s DescribeDomainNsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainNsRequest) SetLang(v string) *DescribeDomainNsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainNsRequest) SetUserClientIp(v string) *DescribeDomainNsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDomainNsRequest) SetDomainName(v string) *DescribeDomainNsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainNsRequest) SetDomainType(v string) *DescribeDomainNsRequest {
	s.DomainType = &v
	return s
}

type DescribeDomainNsResponseBody struct {
	AllAliDns        *bool     `json:"AllAliDns,omitempty" xml:"AllAliDns,omitempty"`
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExpectDnsServers []*string `json:"ExpectDnsServers,omitempty" xml:"ExpectDnsServers,omitempty" type:"Repeated"`
	DnsServers       []*string `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Repeated"`
	IncludeAliDns    *bool     `json:"IncludeAliDns,omitempty" xml:"IncludeAliDns,omitempty"`
}

func (s DescribeDomainNsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainNsResponseBody) SetAllAliDns(v bool) *DescribeDomainNsResponseBody {
	s.AllAliDns = &v
	return s
}

func (s *DescribeDomainNsResponseBody) SetRequestId(v string) *DescribeDomainNsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainNsResponseBody) SetExpectDnsServers(v []*string) *DescribeDomainNsResponseBody {
	s.ExpectDnsServers = v
	return s
}

func (s *DescribeDomainNsResponseBody) SetDnsServers(v []*string) *DescribeDomainNsResponseBody {
	s.DnsServers = v
	return s
}

func (s *DescribeDomainNsResponseBody) SetIncludeAliDns(v bool) *DescribeDomainNsResponseBody {
	s.IncludeAliDns = &v
	return s
}

type DescribeDomainNsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainNsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainNsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainNsResponse) SetHeaders(v map[string]*string) *DescribeDomainNsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainNsResponse) SetBody(v *DescribeDomainNsResponseBody) *DescribeDomainNsResponse {
	s.Body = v
	return s
}

type DescribeDomainRecordInfoRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	RecordId     *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s DescribeDomainRecordInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRecordInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordInfoRequest) SetLang(v string) *DescribeDomainRecordInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainRecordInfoRequest) SetUserClientIp(v string) *DescribeDomainRecordInfoRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDomainRecordInfoRequest) SetRecordId(v string) *DescribeDomainRecordInfoRequest {
	s.RecordId = &v
	return s
}

type DescribeDomainRecordInfoResponseBody struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RR         *string `json:"RR,omitempty" xml:"RR,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Priority   *int64  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	PunyCode   *string `json:"PunyCode,omitempty" xml:"PunyCode,omitempty"`
	TTL        *int64  `json:"TTL,omitempty" xml:"TTL,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Line       *string `json:"Line,omitempty" xml:"Line,omitempty"`
	Locked     *bool   `json:"Locked,omitempty" xml:"Locked,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	DomainId   *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
	RecordId   *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s DescribeDomainRecordInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRecordInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordInfoResponseBody) SetStatus(v string) *DescribeDomainRecordInfoResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetRR(v string) *DescribeDomainRecordInfoResponseBody {
	s.RR = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetGroupName(v string) *DescribeDomainRecordInfoResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetRequestId(v string) *DescribeDomainRecordInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetDomainName(v string) *DescribeDomainRecordInfoResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetPriority(v int64) *DescribeDomainRecordInfoResponseBody {
	s.Priority = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetPunyCode(v string) *DescribeDomainRecordInfoResponseBody {
	s.PunyCode = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetTTL(v int64) *DescribeDomainRecordInfoResponseBody {
	s.TTL = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetGroupId(v string) *DescribeDomainRecordInfoResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetLine(v string) *DescribeDomainRecordInfoResponseBody {
	s.Line = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetLocked(v bool) *DescribeDomainRecordInfoResponseBody {
	s.Locked = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetType(v string) *DescribeDomainRecordInfoResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetDomainId(v string) *DescribeDomainRecordInfoResponseBody {
	s.DomainId = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetValue(v string) *DescribeDomainRecordInfoResponseBody {
	s.Value = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetRecordId(v string) *DescribeDomainRecordInfoResponseBody {
	s.RecordId = &v
	return s
}

type DescribeDomainRecordInfoResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRecordInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRecordInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRecordInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordInfoResponse) SetHeaders(v map[string]*string) *DescribeDomainRecordInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRecordInfoResponse) SetBody(v *DescribeDomainRecordInfoResponseBody) *DescribeDomainRecordInfoResponse {
	s.Body = v
	return s
}

type DescribeDomainRecordsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	KeyWord      *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	RRKeyWord    *string `json:"RRKeyWord,omitempty" xml:"RRKeyWord,omitempty"`
	TypeKeyWord  *string `json:"TypeKeyWord,omitempty" xml:"TypeKeyWord,omitempty"`
	ValueKeyWord *string `json:"ValueKeyWord,omitempty" xml:"ValueKeyWord,omitempty"`
	OrderBy      *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Direction    *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	SearchMode   *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	GroupId      *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Line         *string `json:"Line,omitempty" xml:"Line,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDomainRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordsRequest) SetLang(v string) *DescribeDomainRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetUserClientIp(v string) *DescribeDomainRecordsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetDomainName(v string) *DescribeDomainRecordsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetPageNumber(v int64) *DescribeDomainRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetPageSize(v int64) *DescribeDomainRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetKeyWord(v string) *DescribeDomainRecordsRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetRRKeyWord(v string) *DescribeDomainRecordsRequest {
	s.RRKeyWord = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetTypeKeyWord(v string) *DescribeDomainRecordsRequest {
	s.TypeKeyWord = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetValueKeyWord(v string) *DescribeDomainRecordsRequest {
	s.ValueKeyWord = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetOrderBy(v string) *DescribeDomainRecordsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetDirection(v string) *DescribeDomainRecordsRequest {
	s.Direction = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetSearchMode(v string) *DescribeDomainRecordsRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetGroupId(v int64) *DescribeDomainRecordsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetType(v string) *DescribeDomainRecordsRequest {
	s.Type = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetLine(v string) *DescribeDomainRecordsRequest {
	s.Line = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetStatus(v string) *DescribeDomainRecordsRequest {
	s.Status = &v
	return s
}

type DescribeDomainRecordsResponseBody struct {
	TotalCount    *int64                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize      *int64                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainRecords []*DescribeDomainRecordsResponseBodyDomainRecords `json:"DomainRecords,omitempty" xml:"DomainRecords,omitempty" type:"Repeated"`
	PageNumber    *int64                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeDomainRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordsResponseBody) SetTotalCount(v int64) *DescribeDomainRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainRecordsResponseBody) SetPageSize(v int64) *DescribeDomainRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainRecordsResponseBody) SetRequestId(v string) *DescribeDomainRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRecordsResponseBody) SetDomainRecords(v []*DescribeDomainRecordsResponseBodyDomainRecords) *DescribeDomainRecordsResponseBody {
	s.DomainRecords = v
	return s
}

func (s *DescribeDomainRecordsResponseBody) SetPageNumber(v int64) *DescribeDomainRecordsResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeDomainRecordsResponseBodyDomainRecords struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TTL        *int64  `json:"TTL,omitempty" xml:"TTL,omitempty"`
	RecordId   *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Priority   *int64  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RR         *string `json:"RR,omitempty" xml:"RR,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Weight     *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Line       *string `json:"Line,omitempty" xml:"Line,omitempty"`
	Locked     *bool   `json:"Locked,omitempty" xml:"Locked,omitempty"`
}

func (s DescribeDomainRecordsResponseBodyDomainRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRecordsResponseBodyDomainRecords) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) SetStatus(v string) *DescribeDomainRecordsResponseBodyDomainRecords {
	s.Status = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) SetType(v string) *DescribeDomainRecordsResponseBodyDomainRecords {
	s.Type = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) SetRemark(v string) *DescribeDomainRecordsResponseBodyDomainRecords {
	s.Remark = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) SetTTL(v int64) *DescribeDomainRecordsResponseBodyDomainRecords {
	s.TTL = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) SetRecordId(v string) *DescribeDomainRecordsResponseBodyDomainRecords {
	s.RecordId = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) SetPriority(v int64) *DescribeDomainRecordsResponseBodyDomainRecords {
	s.Priority = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) SetRR(v string) *DescribeDomainRecordsResponseBodyDomainRecords {
	s.RR = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) SetDomainName(v string) *DescribeDomainRecordsResponseBodyDomainRecords {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) SetWeight(v int32) *DescribeDomainRecordsResponseBodyDomainRecords {
	s.Weight = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) SetValue(v string) *DescribeDomainRecordsResponseBodyDomainRecords {
	s.Value = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) SetLine(v string) *DescribeDomainRecordsResponseBodyDomainRecords {
	s.Line = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) SetLocked(v bool) *DescribeDomainRecordsResponseBodyDomainRecords {
	s.Locked = &v
	return s
}

type DescribeDomainRecordsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordsResponse) SetHeaders(v map[string]*string) *DescribeDomainRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRecordsResponse) SetBody(v *DescribeDomainRecordsResponseBody) *DescribeDomainRecordsResponse {
	s.Body = v
	return s
}

type DescribeDomainsRequest struct {
	Lang            *string                      `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp    *string                      `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	KeyWord         *string                      `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	GroupId         *string                      `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber      *int64                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int64                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchMode      *string                      `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	ResourceGroupId *string                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	OrderBy         *string                      `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Direction       *string                      `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Starmark        *bool                        `json:"Starmark,omitempty" xml:"Starmark,omitempty"`
	StartDate       *string                      `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate         *string                      `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Tag             []*DescribeDomainsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequest) SetLang(v string) *DescribeDomainsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainsRequest) SetUserClientIp(v string) *DescribeDomainsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDomainsRequest) SetKeyWord(v string) *DescribeDomainsRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeDomainsRequest) SetGroupId(v string) *DescribeDomainsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageNumber(v int64) *DescribeDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageSize(v int64) *DescribeDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainsRequest) SetSearchMode(v string) *DescribeDomainsRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeDomainsRequest) SetResourceGroupId(v string) *DescribeDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainsRequest) SetOrderBy(v string) *DescribeDomainsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeDomainsRequest) SetDirection(v string) *DescribeDomainsRequest {
	s.Direction = &v
	return s
}

func (s *DescribeDomainsRequest) SetStarmark(v bool) *DescribeDomainsRequest {
	s.Starmark = &v
	return s
}

func (s *DescribeDomainsRequest) SetStartDate(v string) *DescribeDomainsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDomainsRequest) SetEndDate(v string) *DescribeDomainsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDomainsRequest) SetTag(v []*DescribeDomainsRequestTag) *DescribeDomainsRequest {
	s.Tag = v
	return s
}

type DescribeDomainsRequestTag struct {
}

func (s DescribeDomainsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsRequestTag) GoString() string {
	return s.String()
}

type DescribeDomainsResponseBody struct {
	Domains    []*DescribeDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int64                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int64                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBody) SetDomains(v []*DescribeDomainsResponseBodyDomains) *DescribeDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDomainsResponseBody) SetTotalCount(v int64) *DescribeDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetPageSize(v int64) *DescribeDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetRequestId(v string) *DescribeDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetPageNumber(v int64) *DescribeDomainsResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeDomainsResponseBodyDomains struct {
	Remark          *string                                   `json:"Remark,omitempty" xml:"Remark,omitempty"`
	CreateTime      *string                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RecordCount     *int64                                    `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	Tags            []*DescribeDomainsResponseBodyDomainsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	InstanceId      *string                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DomainName      *string                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainId        *string                                   `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	AliDomain       *bool                                     `json:"AliDomain,omitempty" xml:"AliDomain,omitempty"`
	GroupId         *string                                   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName       *string                                   `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ResourceGroupId *string                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceEndTime *string                                   `json:"InstanceEndTime,omitempty" xml:"InstanceEndTime,omitempty"`
	InstanceExpired *bool                                     `json:"InstanceExpired,omitempty" xml:"InstanceExpired,omitempty"`
	VersionName     *string                                   `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	DnsServers      []*string                                 `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Repeated"`
	VersionCode     *string                                   `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	PunyCode        *string                                   `json:"PunyCode,omitempty" xml:"PunyCode,omitempty"`
	RegistrantEmail *string                                   `json:"RegistrantEmail,omitempty" xml:"RegistrantEmail,omitempty"`
	CreateTimestamp *int64                                    `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Starmark        *bool                                     `json:"Starmark,omitempty" xml:"Starmark,omitempty"`
}

func (s DescribeDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomains) SetRemark(v string) *DescribeDomainsResponseBodyDomains {
	s.Remark = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCreateTime(v string) *DescribeDomainsResponseBodyDomains {
	s.CreateTime = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetRecordCount(v int64) *DescribeDomainsResponseBodyDomains {
	s.RecordCount = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetTags(v []*DescribeDomainsResponseBodyDomainsTags) *DescribeDomainsResponseBodyDomains {
	s.Tags = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetInstanceId(v string) *DescribeDomainsResponseBodyDomains {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetDomainName(v string) *DescribeDomainsResponseBodyDomains {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetDomainId(v string) *DescribeDomainsResponseBodyDomains {
	s.DomainId = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetAliDomain(v bool) *DescribeDomainsResponseBodyDomains {
	s.AliDomain = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetGroupId(v string) *DescribeDomainsResponseBodyDomains {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetGroupName(v string) *DescribeDomainsResponseBodyDomains {
	s.GroupName = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetResourceGroupId(v string) *DescribeDomainsResponseBodyDomains {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetInstanceEndTime(v string) *DescribeDomainsResponseBodyDomains {
	s.InstanceEndTime = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetInstanceExpired(v bool) *DescribeDomainsResponseBodyDomains {
	s.InstanceExpired = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetVersionName(v string) *DescribeDomainsResponseBodyDomains {
	s.VersionName = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetDnsServers(v []*string) *DescribeDomainsResponseBodyDomains {
	s.DnsServers = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetVersionCode(v string) *DescribeDomainsResponseBodyDomains {
	s.VersionCode = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetPunyCode(v string) *DescribeDomainsResponseBodyDomains {
	s.PunyCode = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetRegistrantEmail(v string) *DescribeDomainsResponseBodyDomains {
	s.RegistrantEmail = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCreateTimestamp(v int64) *DescribeDomainsResponseBodyDomains {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetStarmark(v bool) *DescribeDomainsResponseBodyDomains {
	s.Starmark = &v
	return s
}

type DescribeDomainsResponseBodyDomainsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsTags) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsTags) SetKey(v string) *DescribeDomainsResponseBodyDomainsTags {
	s.Key = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsTags) SetValue(v string) *DescribeDomainsResponseBodyDomainsTags {
	s.Value = &v
	return s
}

type DescribeDomainsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponse) SetHeaders(v map[string]*string) *DescribeDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainsResponse) SetBody(v *DescribeDomainsResponseBody) *DescribeDomainsResponse {
	s.Body = v
	return s
}

type DescribeDomainStatisticsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	DomainType   *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
}

func (s DescribeDomainStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsRequest) SetLang(v string) *DescribeDomainStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainStatisticsRequest) SetUserClientIp(v string) *DescribeDomainStatisticsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDomainStatisticsRequest) SetDomainName(v string) *DescribeDomainStatisticsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainStatisticsRequest) SetStartDate(v string) *DescribeDomainStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDomainStatisticsRequest) SetEndDate(v string) *DescribeDomainStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDomainStatisticsRequest) SetDomainType(v string) *DescribeDomainStatisticsRequest {
	s.DomainType = &v
	return s
}

type DescribeDomainStatisticsResponseBody struct {
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistics []*DescribeDomainStatisticsResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeDomainStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsResponseBody) SetRequestId(v string) *DescribeDomainStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainStatisticsResponseBody) SetStatistics(v []*DescribeDomainStatisticsResponseBodyStatistics) *DescribeDomainStatisticsResponseBody {
	s.Statistics = v
	return s
}

type DescribeDomainStatisticsResponseBodyStatistics struct {
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Count     *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDomainStatisticsResponseBodyStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatisticsResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsResponseBodyStatistics) SetTimestamp(v int64) *DescribeDomainStatisticsResponseBodyStatistics {
	s.Timestamp = &v
	return s
}

func (s *DescribeDomainStatisticsResponseBodyStatistics) SetCount(v int64) *DescribeDomainStatisticsResponseBodyStatistics {
	s.Count = &v
	return s
}

type DescribeDomainStatisticsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsResponse) SetHeaders(v map[string]*string) *DescribeDomainStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainStatisticsResponse) SetBody(v *DescribeDomainStatisticsResponseBody) *DescribeDomainStatisticsResponse {
	s.Body = v
	return s
}

type DescribeDomainStatisticsSummaryRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OrderBy      *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Direction    *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	SearchMode   *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Threshold    *int64  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeDomainStatisticsSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatisticsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsSummaryRequest) SetLang(v string) *DescribeDomainStatisticsSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetUserClientIp(v string) *DescribeDomainStatisticsSummaryRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetPageNumber(v int64) *DescribeDomainStatisticsSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetPageSize(v int64) *DescribeDomainStatisticsSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetStartDate(v string) *DescribeDomainStatisticsSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetEndDate(v string) *DescribeDomainStatisticsSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetOrderBy(v string) *DescribeDomainStatisticsSummaryRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetDirection(v string) *DescribeDomainStatisticsSummaryRequest {
	s.Direction = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetSearchMode(v string) *DescribeDomainStatisticsSummaryRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetKeyword(v string) *DescribeDomainStatisticsSummaryRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetThreshold(v int64) *DescribeDomainStatisticsSummaryRequest {
	s.Threshold = &v
	return s
}

type DescribeDomainStatisticsSummaryResponseBody struct {
	PageSize   *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalPages *int32                                                   `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems *int32                                                   `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	Statistics []*DescribeDomainStatisticsSummaryResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeDomainStatisticsSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatisticsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsSummaryResponseBody) SetPageSize(v int32) *DescribeDomainStatisticsSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBody) SetRequestId(v string) *DescribeDomainStatisticsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBody) SetPageNumber(v int32) *DescribeDomainStatisticsSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBody) SetTotalPages(v int32) *DescribeDomainStatisticsSummaryResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBody) SetTotalItems(v int32) *DescribeDomainStatisticsSummaryResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBody) SetStatistics(v []*DescribeDomainStatisticsSummaryResponseBodyStatistics) *DescribeDomainStatisticsSummaryResponseBody {
	s.Statistics = v
	return s
}

type DescribeDomainStatisticsSummaryResponseBodyStatistics struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Count      *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
}

func (s DescribeDomainStatisticsSummaryResponseBodyStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatisticsSummaryResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatistics) SetDomainName(v string) *DescribeDomainStatisticsSummaryResponseBodyStatistics {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatistics) SetCount(v int64) *DescribeDomainStatisticsSummaryResponseBodyStatistics {
	s.Count = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatistics) SetDomainType(v string) *DescribeDomainStatisticsSummaryResponseBodyStatistics {
	s.DomainType = &v
	return s
}

type DescribeDomainStatisticsSummaryResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainStatisticsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainStatisticsSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatisticsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsSummaryResponse) SetHeaders(v map[string]*string) *DescribeDomainStatisticsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponse) SetBody(v *DescribeDomainStatisticsSummaryResponseBody) *DescribeDomainStatisticsSummaryResponse {
	s.Body = v
	return s
}

type DescribeGtmAccessStrategiesRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGtmAccessStrategiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAccessStrategiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategiesRequest) SetLang(v string) *DescribeGtmAccessStrategiesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmAccessStrategiesRequest) SetUserClientIp(v string) *DescribeGtmAccessStrategiesRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGtmAccessStrategiesRequest) SetInstanceId(v string) *DescribeGtmAccessStrategiesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmAccessStrategiesRequest) SetPageNumber(v int32) *DescribeGtmAccessStrategiesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmAccessStrategiesRequest) SetPageSize(v int32) *DescribeGtmAccessStrategiesRequest {
	s.PageSize = &v
	return s
}

type DescribeGtmAccessStrategiesResponseBody struct {
	PageSize   *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Strategies []*DescribeGtmAccessStrategiesResponseBodyStrategies `json:"Strategies,omitempty" xml:"Strategies,omitempty" type:"Repeated"`
	TotalPages *int32                                               `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems *int32                                               `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
}

func (s DescribeGtmAccessStrategiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAccessStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategiesResponseBody) SetPageSize(v int32) *DescribeGtmAccessStrategiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBody) SetRequestId(v string) *DescribeGtmAccessStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBody) SetPageNumber(v int32) *DescribeGtmAccessStrategiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBody) SetStrategies(v []*DescribeGtmAccessStrategiesResponseBodyStrategies) *DescribeGtmAccessStrategiesResponseBody {
	s.Strategies = v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBody) SetTotalPages(v int32) *DescribeGtmAccessStrategiesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBody) SetTotalItems(v int32) *DescribeGtmAccessStrategiesResponseBody {
	s.TotalItems = &v
	return s
}

type DescribeGtmAccessStrategiesResponseBodyStrategies struct {
	AccessMode                    *string                                                   `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	StrategyName                  *string                                                   `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	DefaultAddrPoolMonitorStatus  *string                                                   `json:"DefaultAddrPoolMonitorStatus,omitempty" xml:"DefaultAddrPoolMonitorStatus,omitempty"`
	StrategyMode                  *string                                                   `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
	CreateTime                    *string                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DefaultAddrPoolStatus         *string                                                   `json:"DefaultAddrPoolStatus,omitempty" xml:"DefaultAddrPoolStatus,omitempty"`
	InstanceId                    *string                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lines                         []*DescribeGtmAccessStrategiesResponseBodyStrategiesLines `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Repeated"`
	FailoverAddrPoolId            *string                                                   `json:"FailoverAddrPoolId,omitempty" xml:"FailoverAddrPoolId,omitempty"`
	DefaultAddrPoolId             *string                                                   `json:"DefaultAddrPoolId,omitempty" xml:"DefaultAddrPoolId,omitempty"`
	StrategyId                    *string                                                   `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	FailoverAddrPoolStatus        *string                                                   `json:"FailoverAddrPoolStatus,omitempty" xml:"FailoverAddrPoolStatus,omitempty"`
	AccessStatus                  *string                                                   `json:"AccessStatus,omitempty" xml:"AccessStatus,omitempty"`
	FailoverAddrPoolMonitorStatus *string                                                   `json:"FailoverAddrPoolMonitorStatus,omitempty" xml:"FailoverAddrPoolMonitorStatus,omitempty"`
	DefaultAddrPoolName           *string                                                   `json:"DefaultAddrPoolName,omitempty" xml:"DefaultAddrPoolName,omitempty"`
	FailoverAddrPoolName          *string                                                   `json:"FailoverAddrPoolName,omitempty" xml:"FailoverAddrPoolName,omitempty"`
	CreateTimestamp               *int64                                                    `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeGtmAccessStrategiesResponseBodyStrategies) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAccessStrategiesResponseBodyStrategies) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetAccessMode(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.AccessMode = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetStrategyName(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.StrategyName = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetDefaultAddrPoolMonitorStatus(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.DefaultAddrPoolMonitorStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetStrategyMode(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.StrategyMode = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetCreateTime(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetDefaultAddrPoolStatus(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.DefaultAddrPoolStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetInstanceId(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetLines(v []*DescribeGtmAccessStrategiesResponseBodyStrategiesLines) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.Lines = v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetFailoverAddrPoolId(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.FailoverAddrPoolId = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetDefaultAddrPoolId(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.DefaultAddrPoolId = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetStrategyId(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.StrategyId = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetFailoverAddrPoolStatus(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.FailoverAddrPoolStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetAccessStatus(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.AccessStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetFailoverAddrPoolMonitorStatus(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.FailoverAddrPoolMonitorStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetDefaultAddrPoolName(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.DefaultAddrPoolName = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetFailoverAddrPoolName(v string) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.FailoverAddrPoolName = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetCreateTimestamp(v int64) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.CreateTimestamp = &v
	return s
}

type DescribeGtmAccessStrategiesResponseBodyStrategiesLines struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	LineCode  *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	LineName  *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
}

func (s DescribeGtmAccessStrategiesResponseBodyStrategiesLines) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAccessStrategiesResponseBodyStrategiesLines) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesLines) SetGroupName(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesLines {
	s.GroupName = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesLines) SetLineCode(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesLines {
	s.LineCode = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesLines) SetLineName(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesLines {
	s.LineName = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesLines) SetGroupCode(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesLines {
	s.GroupCode = &v
	return s
}

type DescribeGtmAccessStrategiesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmAccessStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmAccessStrategiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAccessStrategiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategiesResponse) SetHeaders(v map[string]*string) *DescribeGtmAccessStrategiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmAccessStrategiesResponse) SetBody(v *DescribeGtmAccessStrategiesResponseBody) *DescribeGtmAccessStrategiesResponse {
	s.Body = v
	return s
}

type DescribeGtmAccessStrategyRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	StrategyId   *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DescribeGtmAccessStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyRequest) SetLang(v string) *DescribeGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmAccessStrategyRequest) SetUserClientIp(v string) *DescribeGtmAccessStrategyRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGtmAccessStrategyRequest) SetStrategyId(v string) *DescribeGtmAccessStrategyRequest {
	s.StrategyId = &v
	return s
}

type DescribeGtmAccessStrategyResponseBody struct {
	RequestId                     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId                    *string                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StrategyId                    *string                                       `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	DefaultAddrPoolStatus         *string                                       `json:"DefaultAddrPoolStatus,omitempty" xml:"DefaultAddrPoolStatus,omitempty"`
	FailoverAddrPoolId            *string                                       `json:"FailoverAddrPoolId,omitempty" xml:"FailoverAddrPoolId,omitempty"`
	AccessStatus                  *string                                       `json:"AccessStatus,omitempty" xml:"AccessStatus,omitempty"`
	DefaultAddrPoolMonitorStatus  *string                                       `json:"DefaultAddrPoolMonitorStatus,omitempty" xml:"DefaultAddrPoolMonitorStatus,omitempty"`
	DefaultAddrPoolName           *string                                       `json:"DefaultAddrPoolName,omitempty" xml:"DefaultAddrPoolName,omitempty"`
	DefultAddrPoolId              *string                                       `json:"DefultAddrPoolId,omitempty" xml:"DefultAddrPoolId,omitempty"`
	StrategyName                  *string                                       `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	FailoverAddrPoolStatus        *string                                       `json:"FailoverAddrPoolStatus,omitempty" xml:"FailoverAddrPoolStatus,omitempty"`
	AccessMode                    *string                                       `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	StrategyMode                  *string                                       `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
	FailoverAddrPoolMonitorStatus *string                                       `json:"FailoverAddrPoolMonitorStatus,omitempty" xml:"FailoverAddrPoolMonitorStatus,omitempty"`
	FailoverAddrPoolName          *string                                       `json:"FailoverAddrPoolName,omitempty" xml:"FailoverAddrPoolName,omitempty"`
	Lines                         []*DescribeGtmAccessStrategyResponseBodyLines `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Repeated"`
}

func (s DescribeGtmAccessStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyResponseBody) SetRequestId(v string) *DescribeGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetInstanceId(v string) *DescribeGtmAccessStrategyResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetStrategyId(v string) *DescribeGtmAccessStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetDefaultAddrPoolStatus(v string) *DescribeGtmAccessStrategyResponseBody {
	s.DefaultAddrPoolStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetFailoverAddrPoolId(v string) *DescribeGtmAccessStrategyResponseBody {
	s.FailoverAddrPoolId = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetAccessStatus(v string) *DescribeGtmAccessStrategyResponseBody {
	s.AccessStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetDefaultAddrPoolMonitorStatus(v string) *DescribeGtmAccessStrategyResponseBody {
	s.DefaultAddrPoolMonitorStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetDefaultAddrPoolName(v string) *DescribeGtmAccessStrategyResponseBody {
	s.DefaultAddrPoolName = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetDefultAddrPoolId(v string) *DescribeGtmAccessStrategyResponseBody {
	s.DefultAddrPoolId = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetStrategyName(v string) *DescribeGtmAccessStrategyResponseBody {
	s.StrategyName = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetFailoverAddrPoolStatus(v string) *DescribeGtmAccessStrategyResponseBody {
	s.FailoverAddrPoolStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetAccessMode(v string) *DescribeGtmAccessStrategyResponseBody {
	s.AccessMode = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetStrategyMode(v string) *DescribeGtmAccessStrategyResponseBody {
	s.StrategyMode = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetFailoverAddrPoolMonitorStatus(v string) *DescribeGtmAccessStrategyResponseBody {
	s.FailoverAddrPoolMonitorStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetFailoverAddrPoolName(v string) *DescribeGtmAccessStrategyResponseBody {
	s.FailoverAddrPoolName = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetLines(v []*DescribeGtmAccessStrategyResponseBodyLines) *DescribeGtmAccessStrategyResponseBody {
	s.Lines = v
	return s
}

type DescribeGtmAccessStrategyResponseBodyLines struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	LineCode  *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	LineName  *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
}

func (s DescribeGtmAccessStrategyResponseBodyLines) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAccessStrategyResponseBodyLines) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyResponseBodyLines) SetGroupName(v string) *DescribeGtmAccessStrategyResponseBodyLines {
	s.GroupName = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBodyLines) SetLineCode(v string) *DescribeGtmAccessStrategyResponseBodyLines {
	s.LineCode = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBodyLines) SetLineName(v string) *DescribeGtmAccessStrategyResponseBodyLines {
	s.LineName = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBodyLines) SetGroupCode(v string) *DescribeGtmAccessStrategyResponseBodyLines {
	s.GroupCode = &v
	return s
}

type DescribeGtmAccessStrategyResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmAccessStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *DescribeGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmAccessStrategyResponse) SetBody(v *DescribeGtmAccessStrategyResponseBody) *DescribeGtmAccessStrategyResponse {
	s.Body = v
	return s
}

type DescribeGtmAccessStrategyAvailableConfigRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeGtmAccessStrategyAvailableConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAccessStrategyAvailableConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyAvailableConfigRequest) SetLang(v string) *DescribeGtmAccessStrategyAvailableConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigRequest) SetUserClientIp(v string) *DescribeGtmAccessStrategyAvailableConfigRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigRequest) SetInstanceId(v string) *DescribeGtmAccessStrategyAvailableConfigRequest {
	s.InstanceId = &v
	return s
}

type DescribeGtmAccessStrategyAvailableConfigResponseBody struct {
	RequestId *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AddrPools []*DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools `json:"AddrPools,omitempty" xml:"AddrPools,omitempty" type:"Repeated"`
	Lines     []*DescribeGtmAccessStrategyAvailableConfigResponseBodyLines     `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Repeated"`
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBody) SetRequestId(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBody) SetAddrPools(v []*DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools) *DescribeGtmAccessStrategyAvailableConfigResponseBody {
	s.AddrPools = v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBody) SetLines(v []*DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) *DescribeGtmAccessStrategyAvailableConfigResponseBody {
	s.Lines = v
	return s
}

type DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools struct {
	AddrPoolId   *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	AddrPoolName *string `json:"AddrPoolName,omitempty" xml:"AddrPoolName,omitempty"`
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools) SetAddrPoolId(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools) SetAddrPoolName(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools {
	s.AddrPoolName = &v
	return s
}

type DescribeGtmAccessStrategyAvailableConfigResponseBodyLines struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	FatherCode *string `json:"FatherCode,omitempty" xml:"FatherCode,omitempty"`
	LineCode   *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	LineName   *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
	GroupCode  *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) SetStatus(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines {
	s.Status = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) SetFatherCode(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines {
	s.FatherCode = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) SetLineCode(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines {
	s.LineCode = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) SetGroupName(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines {
	s.GroupName = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) SetLineName(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines {
	s.LineName = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) SetGroupCode(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines {
	s.GroupCode = &v
	return s
}

type DescribeGtmAccessStrategyAvailableConfigResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmAccessStrategyAvailableConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmAccessStrategyAvailableConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAccessStrategyAvailableConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponse) SetHeaders(v map[string]*string) *DescribeGtmAccessStrategyAvailableConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponse) SetBody(v *DescribeGtmAccessStrategyAvailableConfigResponseBody) *DescribeGtmAccessStrategyAvailableConfigResponse {
	s.Body = v
	return s
}

type DescribeGtmAvailableAlertGroupRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeGtmAvailableAlertGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAvailableAlertGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmAvailableAlertGroupRequest) SetLang(v string) *DescribeGtmAvailableAlertGroupRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmAvailableAlertGroupRequest) SetUserClientIp(v string) *DescribeGtmAvailableAlertGroupRequest {
	s.UserClientIp = &v
	return s
}

type DescribeGtmAvailableAlertGroupResponseBody struct {
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AvailableAlertGroup *string `json:"AvailableAlertGroup,omitempty" xml:"AvailableAlertGroup,omitempty"`
}

func (s DescribeGtmAvailableAlertGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAvailableAlertGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmAvailableAlertGroupResponseBody) SetRequestId(v string) *DescribeGtmAvailableAlertGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmAvailableAlertGroupResponseBody) SetAvailableAlertGroup(v string) *DescribeGtmAvailableAlertGroupResponseBody {
	s.AvailableAlertGroup = &v
	return s
}

type DescribeGtmAvailableAlertGroupResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmAvailableAlertGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmAvailableAlertGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmAvailableAlertGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmAvailableAlertGroupResponse) SetHeaders(v map[string]*string) *DescribeGtmAvailableAlertGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmAvailableAlertGroupResponse) SetBody(v *DescribeGtmAvailableAlertGroupResponseBody) *DescribeGtmAvailableAlertGroupResponse {
	s.Body = v
	return s
}

type DescribeGtmInstanceRequest struct {
	Lang                 *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp         *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NeedDetailAttributes *bool   `json:"NeedDetailAttributes,omitempty" xml:"NeedDetailAttributes,omitempty"`
}

func (s DescribeGtmInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceRequest) SetLang(v string) *DescribeGtmInstanceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmInstanceRequest) SetUserClientIp(v string) *DescribeGtmInstanceRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGtmInstanceRequest) SetInstanceId(v string) *DescribeGtmInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmInstanceRequest) SetNeedDetailAttributes(v bool) *DescribeGtmInstanceRequest {
	s.NeedDetailAttributes = &v
	return s
}

type DescribeGtmInstanceResponseBody struct {
	ExpireTimestamp   *int64  `json:"ExpireTimestamp,omitempty" xml:"ExpireTimestamp,omitempty"`
	UserDomainName    *string `json:"UserDomainName,omitempty" xml:"UserDomainName,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LbaStrategy       *string `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CnameMode         *string `json:"CnameMode,omitempty" xml:"CnameMode,omitempty"`
	Ttl               *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Cname             *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	InstanceName      *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	VersionCode       *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	AlertGroup        *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	AddressPoolNum    *int32  `json:"AddressPoolNum,omitempty" xml:"AddressPoolNum,omitempty"`
	AccessStrategyNum *int32  `json:"AccessStrategyNum,omitempty" xml:"AccessStrategyNum,omitempty"`
	ExpireTime        *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CreateTimestamp   *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeGtmInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceResponseBody) SetExpireTimestamp(v int64) *DescribeGtmInstanceResponseBody {
	s.ExpireTimestamp = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetUserDomainName(v string) *DescribeGtmInstanceResponseBody {
	s.UserDomainName = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetRequestId(v string) *DescribeGtmInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetLbaStrategy(v string) *DescribeGtmInstanceResponseBody {
	s.LbaStrategy = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetInstanceId(v string) *DescribeGtmInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetCreateTime(v string) *DescribeGtmInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetCnameMode(v string) *DescribeGtmInstanceResponseBody {
	s.CnameMode = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetTtl(v int32) *DescribeGtmInstanceResponseBody {
	s.Ttl = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetCname(v string) *DescribeGtmInstanceResponseBody {
	s.Cname = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetInstanceName(v string) *DescribeGtmInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetVersionCode(v string) *DescribeGtmInstanceResponseBody {
	s.VersionCode = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetAlertGroup(v string) *DescribeGtmInstanceResponseBody {
	s.AlertGroup = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetAddressPoolNum(v int32) *DescribeGtmInstanceResponseBody {
	s.AddressPoolNum = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetAccessStrategyNum(v int32) *DescribeGtmInstanceResponseBody {
	s.AccessStrategyNum = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetExpireTime(v string) *DescribeGtmInstanceResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetCreateTimestamp(v int64) *DescribeGtmInstanceResponseBody {
	s.CreateTimestamp = &v
	return s
}

type DescribeGtmInstanceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceResponse) SetHeaders(v map[string]*string) *DescribeGtmInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmInstanceResponse) SetBody(v *DescribeGtmInstanceResponseBody) *DescribeGtmInstanceResponse {
	s.Body = v
	return s
}

type DescribeGtmInstanceAddressPoolRequest struct {
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AddrPoolId   *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
}

func (s DescribeGtmInstanceAddressPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolRequest) SetUserClientIp(v string) *DescribeGtmInstanceAddressPoolRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolRequest) SetLang(v string) *DescribeGtmInstanceAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolRequest) SetAddrPoolId(v string) *DescribeGtmInstanceAddressPoolRequest {
	s.AddrPoolId = &v
	return s
}

type DescribeGtmInstanceAddressPoolResponseBody struct {
	Status              *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Addrs               []*DescribeGtmInstanceAddressPoolResponseBodyAddrs `json:"Addrs,omitempty" xml:"Addrs,omitempty" type:"Repeated"`
	RequestId           *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime          *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AddrCount           *int32                                             `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	Name                *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Type                *string                                            `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime          *string                                            `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	AddrPoolId          *string                                            `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	UpdateTimestamp     *int64                                             `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	MonitorConfigId     *string                                            `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	MinAvailableAddrNum *int32                                             `json:"MinAvailableAddrNum,omitempty" xml:"MinAvailableAddrNum,omitempty"`
	MonitorStatus       *string                                            `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	CreateTimestamp     *int64                                             `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeGtmInstanceAddressPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetStatus(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetAddrs(v []*DescribeGtmInstanceAddressPoolResponseBodyAddrs) *DescribeGtmInstanceAddressPoolResponseBody {
	s.Addrs = v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetRequestId(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetCreateTime(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetAddrCount(v int32) *DescribeGtmInstanceAddressPoolResponseBody {
	s.AddrCount = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetName(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetType(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetUpdateTime(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetAddrPoolId(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetUpdateTimestamp(v int64) *DescribeGtmInstanceAddressPoolResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetMonitorConfigId(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetMinAvailableAddrNum(v int32) *DescribeGtmInstanceAddressPoolResponseBody {
	s.MinAvailableAddrNum = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetMonitorStatus(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.MonitorStatus = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetCreateTimestamp(v int64) *DescribeGtmInstanceAddressPoolResponseBody {
	s.CreateTimestamp = &v
	return s
}

type DescribeGtmInstanceAddressPoolResponseBodyAddrs struct {
	Value           *string `json:"Value,omitempty" xml:"Value,omitempty"`
	UpdateTimestamp *int64  `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	AlertStatus     *string `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	LbaWeight       *int32  `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AddrId          *int64  `json:"AddrId,omitempty" xml:"AddrId,omitempty"`
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeGtmInstanceAddressPoolResponseBodyAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolResponseBodyAddrs) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrs) SetValue(v string) *DescribeGtmInstanceAddressPoolResponseBodyAddrs {
	s.Value = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrs) SetUpdateTimestamp(v int64) *DescribeGtmInstanceAddressPoolResponseBodyAddrs {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrs) SetUpdateTime(v string) *DescribeGtmInstanceAddressPoolResponseBodyAddrs {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrs) SetAlertStatus(v string) *DescribeGtmInstanceAddressPoolResponseBodyAddrs {
	s.AlertStatus = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrs) SetLbaWeight(v int32) *DescribeGtmInstanceAddressPoolResponseBodyAddrs {
	s.LbaWeight = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrs) SetCreateTime(v string) *DescribeGtmInstanceAddressPoolResponseBodyAddrs {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrs) SetAddrId(v int64) *DescribeGtmInstanceAddressPoolResponseBodyAddrs {
	s.AddrId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrs) SetMode(v string) *DescribeGtmInstanceAddressPoolResponseBodyAddrs {
	s.Mode = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrs) SetCreateTimestamp(v int64) *DescribeGtmInstanceAddressPoolResponseBodyAddrs {
	s.CreateTimestamp = &v
	return s
}

type DescribeGtmInstanceAddressPoolResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmInstanceAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmInstanceAddressPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolResponse) SetHeaders(v map[string]*string) *DescribeGtmInstanceAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponse) SetBody(v *DescribeGtmInstanceAddressPoolResponseBody) *DescribeGtmInstanceAddressPoolResponse {
	s.Body = v
	return s
}

type DescribeGtmInstanceAddressPoolsRequest struct {
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGtmInstanceAddressPoolsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolsRequest) SetUserClientIp(v string) *DescribeGtmInstanceAddressPoolsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsRequest) SetLang(v string) *DescribeGtmInstanceAddressPoolsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsRequest) SetInstanceId(v string) *DescribeGtmInstanceAddressPoolsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsRequest) SetPageNumber(v int32) *DescribeGtmInstanceAddressPoolsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsRequest) SetPageSize(v int32) *DescribeGtmInstanceAddressPoolsRequest {
	s.PageSize = &v
	return s
}

type DescribeGtmInstanceAddressPoolsResponseBody struct {
	PageSize   *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalPages *int32                                                  `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems *int32                                                  `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	AddrPools  []*DescribeGtmInstanceAddressPoolsResponseBodyAddrPools `json:"AddrPools,omitempty" xml:"AddrPools,omitempty" type:"Repeated"`
}

func (s DescribeGtmInstanceAddressPoolsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) SetPageSize(v int32) *DescribeGtmInstanceAddressPoolsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) SetRequestId(v string) *DescribeGtmInstanceAddressPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) SetPageNumber(v int32) *DescribeGtmInstanceAddressPoolsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) SetTotalPages(v int32) *DescribeGtmInstanceAddressPoolsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) SetTotalItems(v int32) *DescribeGtmInstanceAddressPoolsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) SetAddrPools(v []*DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) *DescribeGtmInstanceAddressPoolsResponseBody {
	s.AddrPools = v
	return s
}

type DescribeGtmInstanceAddressPoolsResponseBodyAddrPools struct {
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MonitorConfigId     *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	MinAvailableAddrNum *int32  `json:"MinAvailableAddrNum,omitempty" xml:"MinAvailableAddrNum,omitempty"`
	UpdateTimestamp     *int64  `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	MonitorStatus       *string `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	AddrPoolId          *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AddrCount           *int32  `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	CreateTimestamp     *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) SetType(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.Type = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) SetStatus(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.Status = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) SetUpdateTime(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) SetCreateTime(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) SetMonitorConfigId(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) SetMinAvailableAddrNum(v int32) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.MinAvailableAddrNum = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) SetUpdateTimestamp(v int64) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) SetMonitorStatus(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.MonitorStatus = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) SetAddrPoolId(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) SetName(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.Name = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) SetAddrCount(v int32) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.AddrCount = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) SetCreateTimestamp(v int64) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.CreateTimestamp = &v
	return s
}

type DescribeGtmInstanceAddressPoolsResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmInstanceAddressPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmInstanceAddressPoolsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolsResponse) SetHeaders(v map[string]*string) *DescribeGtmInstanceAddressPoolsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponse) SetBody(v *DescribeGtmInstanceAddressPoolsResponseBody) *DescribeGtmInstanceAddressPoolsResponse {
	s.Body = v
	return s
}

type DescribeGtmInstancesRequest struct {
	Lang                 *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp         *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Keyword              *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	NeedDetailAttributes *bool   `json:"NeedDetailAttributes,omitempty" xml:"NeedDetailAttributes,omitempty"`
}

func (s DescribeGtmInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstancesRequest) SetLang(v string) *DescribeGtmInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmInstancesRequest) SetUserClientIp(v string) *DescribeGtmInstancesRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGtmInstancesRequest) SetPageNumber(v int32) *DescribeGtmInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmInstancesRequest) SetPageSize(v int32) *DescribeGtmInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmInstancesRequest) SetKeyword(v string) *DescribeGtmInstancesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeGtmInstancesRequest) SetResourceGroupId(v string) *DescribeGtmInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGtmInstancesRequest) SetNeedDetailAttributes(v bool) *DescribeGtmInstancesRequest {
	s.NeedDetailAttributes = &v
	return s
}

type DescribeGtmInstancesResponseBody struct {
	PageSize     *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber   *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	GtmInstances []*DescribeGtmInstancesResponseBodyGtmInstances `json:"GtmInstances,omitempty" xml:"GtmInstances,omitempty" type:"Repeated"`
	TotalPages   *int32                                          `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems   *int32                                          `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
}

func (s DescribeGtmInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstancesResponseBody) SetPageSize(v int32) *DescribeGtmInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmInstancesResponseBody) SetRequestId(v string) *DescribeGtmInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmInstancesResponseBody) SetPageNumber(v int32) *DescribeGtmInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmInstancesResponseBody) SetGtmInstances(v []*DescribeGtmInstancesResponseBodyGtmInstances) *DescribeGtmInstancesResponseBody {
	s.GtmInstances = v
	return s
}

func (s *DescribeGtmInstancesResponseBody) SetTotalPages(v int32) *DescribeGtmInstancesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeGtmInstancesResponseBody) SetTotalItems(v int32) *DescribeGtmInstancesResponseBody {
	s.TotalItems = &v
	return s
}

type DescribeGtmInstancesResponseBodyGtmInstances struct {
	ExpireTime        *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	AccessStrategyNum *int32  `json:"AccessStrategyNum,omitempty" xml:"AccessStrategyNum,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CnameMode         *string `json:"CnameMode,omitempty" xml:"CnameMode,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ExpireTimestamp   *int64  `json:"ExpireTimestamp,omitempty" xml:"ExpireTimestamp,omitempty"`
	Ttl               *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	AlertGroup        *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	AddressPoolNum    *int32  `json:"AddressPoolNum,omitempty" xml:"AddressPoolNum,omitempty"`
	InstanceName      *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	LbaStrategy       *string `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	Cname             *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	VersionCode       *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	UserDomainName    *string `json:"UserDomainName,omitempty" xml:"UserDomainName,omitempty"`
	CreateTimestamp   *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeGtmInstancesResponseBodyGtmInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstancesResponseBodyGtmInstances) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetExpireTime(v string) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetAccessStrategyNum(v int32) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.AccessStrategyNum = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetCreateTime(v string) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetCnameMode(v string) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.CnameMode = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetInstanceId(v string) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetExpireTimestamp(v int64) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.ExpireTimestamp = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetTtl(v int32) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.Ttl = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetAlertGroup(v string) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.AlertGroup = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetAddressPoolNum(v int32) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.AddressPoolNum = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetInstanceName(v string) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetLbaStrategy(v string) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.LbaStrategy = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetCname(v string) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.Cname = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetVersionCode(v string) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.VersionCode = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetUserDomainName(v string) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.UserDomainName = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetCreateTimestamp(v int64) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.CreateTimestamp = &v
	return s
}

type DescribeGtmInstancesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstancesResponse) SetHeaders(v map[string]*string) *DescribeGtmInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmInstancesResponse) SetBody(v *DescribeGtmInstancesResponseBody) *DescribeGtmInstancesResponse {
	s.Body = v
	return s
}

type DescribeGtmInstanceStatusRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeGtmInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceStatusRequest) SetLang(v string) *DescribeGtmInstanceStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmInstanceStatusRequest) SetUserClientIp(v string) *DescribeGtmInstanceStatusRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGtmInstanceStatusRequest) SetInstanceId(v string) *DescribeGtmInstanceStatusRequest {
	s.InstanceId = &v
	return s
}

type DescribeGtmInstanceStatusResponseBody struct {
	Status                      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StrategyNotAvailableNum     *int32  `json:"StrategyNotAvailableNum,omitempty" xml:"StrategyNotAvailableNum,omitempty"`
	RequestId                   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SwitchToFailoverStrategyNum *int32  `json:"SwitchToFailoverStrategyNum,omitempty" xml:"SwitchToFailoverStrategyNum,omitempty"`
	StatusReason                *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	AddrNotAvailableNum         *int32  `json:"AddrNotAvailableNum,omitempty" xml:"AddrNotAvailableNum,omitempty"`
	AddrPoolNotAvailableNum     *int32  `json:"AddrPoolNotAvailableNum,omitempty" xml:"AddrPoolNotAvailableNum,omitempty"`
}

func (s DescribeGtmInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceStatusResponseBody) SetStatus(v string) *DescribeGtmInstanceStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGtmInstanceStatusResponseBody) SetStrategyNotAvailableNum(v int32) *DescribeGtmInstanceStatusResponseBody {
	s.StrategyNotAvailableNum = &v
	return s
}

func (s *DescribeGtmInstanceStatusResponseBody) SetRequestId(v string) *DescribeGtmInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmInstanceStatusResponseBody) SetSwitchToFailoverStrategyNum(v int32) *DescribeGtmInstanceStatusResponseBody {
	s.SwitchToFailoverStrategyNum = &v
	return s
}

func (s *DescribeGtmInstanceStatusResponseBody) SetStatusReason(v string) *DescribeGtmInstanceStatusResponseBody {
	s.StatusReason = &v
	return s
}

func (s *DescribeGtmInstanceStatusResponseBody) SetAddrNotAvailableNum(v int32) *DescribeGtmInstanceStatusResponseBody {
	s.AddrNotAvailableNum = &v
	return s
}

func (s *DescribeGtmInstanceStatusResponseBody) SetAddrPoolNotAvailableNum(v int32) *DescribeGtmInstanceStatusResponseBody {
	s.AddrPoolNotAvailableNum = &v
	return s
}

type DescribeGtmInstanceStatusResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceStatusResponse) SetHeaders(v map[string]*string) *DescribeGtmInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmInstanceStatusResponse) SetBody(v *DescribeGtmInstanceStatusResponseBody) *DescribeGtmInstanceStatusResponse {
	s.Body = v
	return s
}

type DescribeGtmInstanceSystemCnameRequest struct {
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeGtmInstanceSystemCnameRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceSystemCnameRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceSystemCnameRequest) SetUserClientIp(v string) *DescribeGtmInstanceSystemCnameRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGtmInstanceSystemCnameRequest) SetLang(v string) *DescribeGtmInstanceSystemCnameRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmInstanceSystemCnameRequest) SetInstanceId(v string) *DescribeGtmInstanceSystemCnameRequest {
	s.InstanceId = &v
	return s
}

type DescribeGtmInstanceSystemCnameResponseBody struct {
	SystemCname *string `json:"SystemCname,omitempty" xml:"SystemCname,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGtmInstanceSystemCnameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceSystemCnameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceSystemCnameResponseBody) SetSystemCname(v string) *DescribeGtmInstanceSystemCnameResponseBody {
	s.SystemCname = &v
	return s
}

func (s *DescribeGtmInstanceSystemCnameResponseBody) SetRequestId(v string) *DescribeGtmInstanceSystemCnameResponseBody {
	s.RequestId = &v
	return s
}

type DescribeGtmInstanceSystemCnameResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmInstanceSystemCnameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmInstanceSystemCnameResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmInstanceSystemCnameResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceSystemCnameResponse) SetHeaders(v map[string]*string) *DescribeGtmInstanceSystemCnameResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmInstanceSystemCnameResponse) SetBody(v *DescribeGtmInstanceSystemCnameResponseBody) *DescribeGtmInstanceSystemCnameResponse {
	s.Body = v
	return s
}

type DescribeGtmLogsRequest struct {
	UserClientIp   *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Keyword        *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTimestamp *int64  `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	EndTimestamp   *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
}

func (s DescribeGtmLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmLogsRequest) SetUserClientIp(v string) *DescribeGtmLogsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGtmLogsRequest) SetLang(v string) *DescribeGtmLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmLogsRequest) SetInstanceId(v string) *DescribeGtmLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmLogsRequest) SetKeyword(v string) *DescribeGtmLogsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeGtmLogsRequest) SetPageNumber(v int32) *DescribeGtmLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmLogsRequest) SetPageSize(v int32) *DescribeGtmLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmLogsRequest) SetStartTimestamp(v int64) *DescribeGtmLogsRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeGtmLogsRequest) SetEndTimestamp(v int64) *DescribeGtmLogsRequest {
	s.EndTimestamp = &v
	return s
}

type DescribeGtmLogsResponseBody struct {
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalPages *int32                             `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Logs       []*DescribeGtmLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	TotalItems *int32                             `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
}

func (s DescribeGtmLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmLogsResponseBody) SetPageSize(v int32) *DescribeGtmLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmLogsResponseBody) SetRequestId(v string) *DescribeGtmLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmLogsResponseBody) SetPageNumber(v int32) *DescribeGtmLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmLogsResponseBody) SetTotalPages(v int32) *DescribeGtmLogsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeGtmLogsResponseBody) SetLogs(v []*DescribeGtmLogsResponseBodyLogs) *DescribeGtmLogsResponseBody {
	s.Logs = v
	return s
}

func (s *DescribeGtmLogsResponseBody) SetTotalItems(v int32) *DescribeGtmLogsResponseBody {
	s.TotalItems = &v
	return s
}

type DescribeGtmLogsResponseBodyLogs struct {
	OperTimestamp *int64  `json:"OperTimestamp,omitempty" xml:"OperTimestamp,omitempty"`
	EntityId      *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType    *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	OperTime      *string `json:"OperTime,omitempty" xml:"OperTime,omitempty"`
	OperIp        *string `json:"OperIp,omitempty" xml:"OperIp,omitempty"`
	OperAction    *string `json:"OperAction,omitempty" xml:"OperAction,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	EntityName    *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeGtmLogsResponseBodyLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *DescribeGtmLogsResponseBodyLogs) SetOperTimestamp(v int64) *DescribeGtmLogsResponseBodyLogs {
	s.OperTimestamp = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogs) SetEntityId(v string) *DescribeGtmLogsResponseBodyLogs {
	s.EntityId = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogs) SetEntityType(v string) *DescribeGtmLogsResponseBodyLogs {
	s.EntityType = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogs) SetOperTime(v string) *DescribeGtmLogsResponseBodyLogs {
	s.OperTime = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogs) SetOperIp(v string) *DescribeGtmLogsResponseBodyLogs {
	s.OperIp = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogs) SetOperAction(v string) *DescribeGtmLogsResponseBodyLogs {
	s.OperAction = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogs) SetContent(v string) *DescribeGtmLogsResponseBodyLogs {
	s.Content = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogs) SetEntityName(v string) *DescribeGtmLogsResponseBodyLogs {
	s.EntityName = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogs) SetId(v int64) *DescribeGtmLogsResponseBodyLogs {
	s.Id = &v
	return s
}

type DescribeGtmLogsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmLogsResponse) SetHeaders(v map[string]*string) *DescribeGtmLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmLogsResponse) SetBody(v *DescribeGtmLogsResponseBody) *DescribeGtmLogsResponse {
	s.Body = v
	return s
}

type DescribeGtmMonitorAvailableConfigRequest struct {
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeGtmMonitorAvailableConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmMonitorAvailableConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorAvailableConfigRequest) SetUserClientIp(v string) *DescribeGtmMonitorAvailableConfigRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigRequest) SetLang(v string) *DescribeGtmMonitorAvailableConfigRequest {
	s.Lang = &v
	return s
}

type DescribeGtmMonitorAvailableConfigResponseBody struct {
	RequestId    *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IspCityNodes []*DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Repeated"`
}

func (s DescribeGtmMonitorAvailableConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmMonitorAvailableConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorAvailableConfigResponseBody) SetRequestId(v string) *DescribeGtmMonitorAvailableConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBody) SetIspCityNodes(v []*DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) *DescribeGtmMonitorAvailableConfigResponseBody {
	s.IspCityNodes = v
	return s
}

type DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes struct {
	CityCode        *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	Mainland        *bool   `json:"Mainland,omitempty" xml:"Mainland,omitempty"`
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IspCode         *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	CityName        *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	IspName         *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	GroupType       *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	DefaultSelected *bool   `json:"DefaultSelected,omitempty" xml:"DefaultSelected,omitempty"`
}

func (s DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) SetCityCode(v string) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes {
	s.CityCode = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) SetMainland(v bool) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes {
	s.Mainland = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) SetGroupName(v string) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes {
	s.GroupName = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) SetIspCode(v string) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes {
	s.IspCode = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) SetCityName(v string) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes {
	s.CityName = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) SetIspName(v string) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes {
	s.IspName = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) SetGroupType(v string) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes {
	s.GroupType = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) SetDefaultSelected(v bool) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes {
	s.DefaultSelected = &v
	return s
}

type DescribeGtmMonitorAvailableConfigResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmMonitorAvailableConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmMonitorAvailableConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmMonitorAvailableConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorAvailableConfigResponse) SetHeaders(v map[string]*string) *DescribeGtmMonitorAvailableConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponse) SetBody(v *DescribeGtmMonitorAvailableConfigResponseBody) *DescribeGtmMonitorAvailableConfigResponse {
	s.Body = v
	return s
}

type DescribeGtmMonitorConfigRequest struct {
	UserClientIp    *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
}

func (s DescribeGtmMonitorConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmMonitorConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorConfigRequest) SetUserClientIp(v string) *DescribeGtmMonitorConfigRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGtmMonitorConfigRequest) SetLang(v string) *DescribeGtmMonitorConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmMonitorConfigRequest) SetMonitorConfigId(v string) *DescribeGtmMonitorConfigRequest {
	s.MonitorConfigId = &v
	return s
}

type DescribeGtmMonitorConfigResponseBody struct {
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Timeout           *int32                                              `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	ProtocolType      *string                                             `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	IspCityNodes      []*DescribeGtmMonitorConfigResponseBodyIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Repeated"`
	CreateTime        *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UpdateTime        *string                                             `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	EvaluationCount   *int32                                              `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	UpdateTimestamp   *int64                                              `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	MonitorExtendInfo *string                                             `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
	MonitorConfigId   *string                                             `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	CreateTimestamp   *int64                                              `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Interval          *int32                                              `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeGtmMonitorConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmMonitorConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorConfigResponseBody) SetRequestId(v string) *DescribeGtmMonitorConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetTimeout(v int32) *DescribeGtmMonitorConfigResponseBody {
	s.Timeout = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetProtocolType(v string) *DescribeGtmMonitorConfigResponseBody {
	s.ProtocolType = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetIspCityNodes(v []*DescribeGtmMonitorConfigResponseBodyIspCityNodes) *DescribeGtmMonitorConfigResponseBody {
	s.IspCityNodes = v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetCreateTime(v string) *DescribeGtmMonitorConfigResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetUpdateTime(v string) *DescribeGtmMonitorConfigResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetEvaluationCount(v int32) *DescribeGtmMonitorConfigResponseBody {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetUpdateTimestamp(v int64) *DescribeGtmMonitorConfigResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetMonitorExtendInfo(v string) *DescribeGtmMonitorConfigResponseBody {
	s.MonitorExtendInfo = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetMonitorConfigId(v string) *DescribeGtmMonitorConfigResponseBody {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetCreateTimestamp(v int64) *DescribeGtmMonitorConfigResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetInterval(v int32) *DescribeGtmMonitorConfigResponseBody {
	s.Interval = &v
	return s
}

type DescribeGtmMonitorConfigResponseBodyIspCityNodes struct {
	CityCode    *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	IspCode     *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	CityName    *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	IspName     *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s DescribeGtmMonitorConfigResponseBodyIspCityNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmMonitorConfigResponseBodyIspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodes) SetCityCode(v string) *DescribeGtmMonitorConfigResponseBodyIspCityNodes {
	s.CityCode = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodes) SetCountryName(v string) *DescribeGtmMonitorConfigResponseBodyIspCityNodes {
	s.CountryName = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodes) SetIspCode(v string) *DescribeGtmMonitorConfigResponseBodyIspCityNodes {
	s.IspCode = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodes) SetCityName(v string) *DescribeGtmMonitorConfigResponseBodyIspCityNodes {
	s.CityName = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodes) SetCountryCode(v string) *DescribeGtmMonitorConfigResponseBodyIspCityNodes {
	s.CountryCode = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodes) SetIspName(v string) *DescribeGtmMonitorConfigResponseBodyIspCityNodes {
	s.IspName = &v
	return s
}

type DescribeGtmMonitorConfigResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmMonitorConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmMonitorConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmMonitorConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorConfigResponse) SetHeaders(v map[string]*string) *DescribeGtmMonitorConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmMonitorConfigResponse) SetBody(v *DescribeGtmMonitorConfigResponseBody) *DescribeGtmMonitorConfigResponse {
	s.Body = v
	return s
}

type DescribeGtmRecoveryPlanRequest struct {
	UserClientIp   *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RecoveryPlanId *int64  `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
}

func (s DescribeGtmRecoveryPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanRequest) SetUserClientIp(v string) *DescribeGtmRecoveryPlanRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGtmRecoveryPlanRequest) SetLang(v string) *DescribeGtmRecoveryPlanRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmRecoveryPlanRequest) SetRecoveryPlanId(v int64) *DescribeGtmRecoveryPlanRequest {
	s.RecoveryPlanId = &v
	return s
}

type DescribeGtmRecoveryPlanResponseBody struct {
	Status                *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	LastRollbackTime      *string                                              `json:"LastRollbackTime,omitempty" xml:"LastRollbackTime,omitempty"`
	FaultAddrPoolNum      *int32                                               `json:"FaultAddrPoolNum,omitempty" xml:"FaultAddrPoolNum,omitempty"`
	FaultAddrPools        []*DescribeGtmRecoveryPlanResponseBodyFaultAddrPools `json:"FaultAddrPools,omitempty" xml:"FaultAddrPools,omitempty" type:"Repeated"`
	LastExecuteTime       *string                                              `json:"LastExecuteTime,omitempty" xml:"LastExecuteTime,omitempty"`
	RequestId             *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime            *string                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LastExecuteTimestamp  *int64                                               `json:"LastExecuteTimestamp,omitempty" xml:"LastExecuteTimestamp,omitempty"`
	Remark                *string                                              `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Name                  *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	RecoveryPlanId        *int64                                               `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
	UpdateTime            *string                                              `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp       *int64                                               `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	LastRollbackTimestamp *int64                                               `json:"LastRollbackTimestamp,omitempty" xml:"LastRollbackTimestamp,omitempty"`
	CreateTimestamp       *int64                                               `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeGtmRecoveryPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetStatus(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetLastRollbackTime(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.LastRollbackTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetFaultAddrPoolNum(v int32) *DescribeGtmRecoveryPlanResponseBody {
	s.FaultAddrPoolNum = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetFaultAddrPools(v []*DescribeGtmRecoveryPlanResponseBodyFaultAddrPools) *DescribeGtmRecoveryPlanResponseBody {
	s.FaultAddrPools = v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetLastExecuteTime(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.LastExecuteTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetRequestId(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetCreateTime(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetLastExecuteTimestamp(v int64) *DescribeGtmRecoveryPlanResponseBody {
	s.LastExecuteTimestamp = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetRemark(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetName(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetRecoveryPlanId(v int64) *DescribeGtmRecoveryPlanResponseBody {
	s.RecoveryPlanId = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetUpdateTime(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetUpdateTimestamp(v int64) *DescribeGtmRecoveryPlanResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetLastRollbackTimestamp(v int64) *DescribeGtmRecoveryPlanResponseBody {
	s.LastRollbackTimestamp = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetCreateTimestamp(v int64) *DescribeGtmRecoveryPlanResponseBody {
	s.CreateTimestamp = &v
	return s
}

type DescribeGtmRecoveryPlanResponseBodyFaultAddrPools struct {
	Addrs        []*DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsAddrs `json:"Addrs,omitempty" xml:"Addrs,omitempty" type:"Repeated"`
	AddrPoolId   *string                                                   `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	InstanceId   *string                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AddrPoolName *string                                                   `json:"AddrPoolName,omitempty" xml:"AddrPoolName,omitempty"`
}

func (s DescribeGtmRecoveryPlanResponseBodyFaultAddrPools) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmRecoveryPlanResponseBodyFaultAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools) SetAddrs(v []*DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsAddrs) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools {
	s.Addrs = v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools) SetAddrPoolId(v string) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools) SetInstanceId(v string) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools) SetAddrPoolName(v string) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools {
	s.AddrPoolName = &v
	return s
}

type DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsAddrs struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Mode  *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsAddrs) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsAddrs) SetValue(v string) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsAddrs {
	s.Value = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsAddrs) SetMode(v string) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsAddrs {
	s.Mode = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsAddrs) SetId(v int64) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsAddrs {
	s.Id = &v
	return s
}

type DescribeGtmRecoveryPlanResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmRecoveryPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanResponse) SetHeaders(v map[string]*string) *DescribeGtmRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmRecoveryPlanResponse) SetBody(v *DescribeGtmRecoveryPlanResponseBody) *DescribeGtmRecoveryPlanResponse {
	s.Body = v
	return s
}

type DescribeGtmRecoveryPlanAvailableConfigRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeGtmRecoveryPlanAvailableConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmRecoveryPlanAvailableConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanAvailableConfigRequest) SetLang(v string) *DescribeGtmRecoveryPlanAvailableConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigRequest) SetUserClientIp(v string) *DescribeGtmRecoveryPlanAvailableConfigRequest {
	s.UserClientIp = &v
	return s
}

type DescribeGtmRecoveryPlanAvailableConfigResponseBody struct {
	Instances []*DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	RequestId *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBody) SetInstances(v []*DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances) *DescribeGtmRecoveryPlanAvailableConfigResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBody) SetRequestId(v string) *DescribeGtmRecoveryPlanAvailableConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances struct {
	AddrPools    []*DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesAddrPools `json:"AddrPools,omitempty" xml:"AddrPools,omitempty" type:"Repeated"`
	InstanceName *string                                                                 `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceId   *string                                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances) SetAddrPools(v []*DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesAddrPools) *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances {
	s.AddrPools = v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances) SetInstanceName(v string) *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances) SetInstanceId(v string) *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstances {
	s.InstanceId = &v
	return s
}

type DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesAddrPools struct {
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesAddrPools) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesAddrPools) SetAddrPoolId(v string) *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesAddrPools {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesAddrPools) SetName(v string) *DescribeGtmRecoveryPlanAvailableConfigResponseBodyInstancesAddrPools {
	s.Name = &v
	return s
}

type DescribeGtmRecoveryPlanAvailableConfigResponse struct {
	Headers map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmRecoveryPlanAvailableConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponse) SetHeaders(v map[string]*string) *DescribeGtmRecoveryPlanAvailableConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponse) SetBody(v *DescribeGtmRecoveryPlanAvailableConfigResponseBody) *DescribeGtmRecoveryPlanAvailableConfigResponse {
	s.Body = v
	return s
}

type DescribeGtmRecoveryPlansRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGtmRecoveryPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmRecoveryPlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlansRequest) SetLang(v string) *DescribeGtmRecoveryPlansRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmRecoveryPlansRequest) SetUserClientIp(v string) *DescribeGtmRecoveryPlansRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGtmRecoveryPlansRequest) SetKeyword(v string) *DescribeGtmRecoveryPlansRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeGtmRecoveryPlansRequest) SetPageNumber(v int32) *DescribeGtmRecoveryPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmRecoveryPlansRequest) SetPageSize(v int32) *DescribeGtmRecoveryPlansRequest {
	s.PageSize = &v
	return s
}

type DescribeGtmRecoveryPlansResponseBody struct {
	RecoveryPlans []*DescribeGtmRecoveryPlansResponseBodyRecoveryPlans `json:"RecoveryPlans,omitempty" xml:"RecoveryPlans,omitempty" type:"Repeated"`
	PageSize      *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber    *int32                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalPages    *int32                                               `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems    *int32                                               `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
}

func (s DescribeGtmRecoveryPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmRecoveryPlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlansResponseBody) SetRecoveryPlans(v []*DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) *DescribeGtmRecoveryPlansResponseBody {
	s.RecoveryPlans = v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBody) SetPageSize(v int32) *DescribeGtmRecoveryPlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBody) SetRequestId(v string) *DescribeGtmRecoveryPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBody) SetPageNumber(v int32) *DescribeGtmRecoveryPlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBody) SetTotalPages(v int32) *DescribeGtmRecoveryPlansResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBody) SetTotalItems(v int32) *DescribeGtmRecoveryPlansResponseBody {
	s.TotalItems = &v
	return s
}

type DescribeGtmRecoveryPlansResponseBodyRecoveryPlans struct {
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	LastRollbackTimestamp *int64  `json:"LastRollbackTimestamp,omitempty" xml:"LastRollbackTimestamp,omitempty"`
	UpdateTime            *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Remark                *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	CreateTime            *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RecoveryPlanId        *int64  `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
	UpdateTimestamp       *int64  `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	LastExecuteTimestamp  *int64  `json:"LastExecuteTimestamp,omitempty" xml:"LastExecuteTimestamp,omitempty"`
	LastExecuteTime       *string `json:"LastExecuteTime,omitempty" xml:"LastExecuteTime,omitempty"`
	LastRollbackTime      *string `json:"LastRollbackTime,omitempty" xml:"LastRollbackTime,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	FaultAddrPoolNum      *int32  `json:"FaultAddrPoolNum,omitempty" xml:"FaultAddrPoolNum,omitempty"`
	CreateTimestamp       *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) SetStatus(v string) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	s.Status = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) SetLastRollbackTimestamp(v int64) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	s.LastRollbackTimestamp = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) SetUpdateTime(v string) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) SetRemark(v string) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	s.Remark = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) SetCreateTime(v string) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) SetRecoveryPlanId(v int64) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	s.RecoveryPlanId = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) SetUpdateTimestamp(v int64) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) SetLastExecuteTimestamp(v int64) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	s.LastExecuteTimestamp = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) SetLastExecuteTime(v string) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	s.LastExecuteTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) SetLastRollbackTime(v string) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	s.LastRollbackTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) SetName(v string) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	s.Name = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) SetFaultAddrPoolNum(v int32) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	s.FaultAddrPoolNum = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) SetCreateTimestamp(v int64) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	s.CreateTimestamp = &v
	return s
}

type DescribeGtmRecoveryPlansResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGtmRecoveryPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGtmRecoveryPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGtmRecoveryPlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlansResponse) SetHeaders(v map[string]*string) *DescribeGtmRecoveryPlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmRecoveryPlansResponse) SetBody(v *DescribeGtmRecoveryPlansResponseBody) *DescribeGtmRecoveryPlansResponse {
	s.Body = v
	return s
}

type DescribeInstanceDomainsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDomainsRequest) SetLang(v string) *DescribeInstanceDomainsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceDomainsRequest) SetUserClientIp(v string) *DescribeInstanceDomainsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeInstanceDomainsRequest) SetPageNumber(v int64) *DescribeInstanceDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceDomainsRequest) SetPageSize(v int64) *DescribeInstanceDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceDomainsRequest) SetInstanceId(v string) *DescribeInstanceDomainsRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceDomainsResponseBody struct {
	PageSize        *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	InstanceDomains []*DescribeInstanceDomainsResponseBodyInstanceDomains `json:"InstanceDomains,omitempty" xml:"InstanceDomains,omitempty" type:"Repeated"`
	TotalPages      *int32                                                `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems      *int32                                                `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
}

func (s DescribeInstanceDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDomainsResponseBody) SetPageSize(v int32) *DescribeInstanceDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceDomainsResponseBody) SetRequestId(v string) *DescribeInstanceDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceDomainsResponseBody) SetPageNumber(v int32) *DescribeInstanceDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceDomainsResponseBody) SetInstanceDomains(v []*DescribeInstanceDomainsResponseBodyInstanceDomains) *DescribeInstanceDomainsResponseBody {
	s.InstanceDomains = v
	return s
}

func (s *DescribeInstanceDomainsResponseBody) SetTotalPages(v int32) *DescribeInstanceDomainsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeInstanceDomainsResponseBody) SetTotalItems(v int32) *DescribeInstanceDomainsResponseBody {
	s.TotalItems = &v
	return s
}

type DescribeInstanceDomainsResponseBodyInstanceDomains struct {
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeInstanceDomainsResponseBodyInstanceDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDomainsResponseBodyInstanceDomains) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDomainsResponseBodyInstanceDomains) SetCreateTime(v string) *DescribeInstanceDomainsResponseBodyInstanceDomains {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstanceDomainsResponseBodyInstanceDomains) SetDomainName(v string) *DescribeInstanceDomainsResponseBodyInstanceDomains {
	s.DomainName = &v
	return s
}

func (s *DescribeInstanceDomainsResponseBodyInstanceDomains) SetCreateTimestamp(v int64) *DescribeInstanceDomainsResponseBodyInstanceDomains {
	s.CreateTimestamp = &v
	return s
}

type DescribeInstanceDomainsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDomainsResponse) SetHeaders(v map[string]*string) *DescribeInstanceDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceDomainsResponse) SetBody(v *DescribeInstanceDomainsResponseBody) *DescribeInstanceDomainsResponse {
	s.Body = v
	return s
}

type DescribeRecordLogsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	KeyWord      *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate      *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
}

func (s DescribeRecordLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordLogsRequest) SetLang(v string) *DescribeRecordLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRecordLogsRequest) SetUserClientIp(v string) *DescribeRecordLogsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeRecordLogsRequest) SetDomainName(v string) *DescribeRecordLogsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeRecordLogsRequest) SetPageNumber(v int64) *DescribeRecordLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRecordLogsRequest) SetPageSize(v int64) *DescribeRecordLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordLogsRequest) SetKeyWord(v string) *DescribeRecordLogsRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeRecordLogsRequest) SetStartDate(v string) *DescribeRecordLogsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeRecordLogsRequest) SetEndDate(v string) *DescribeRecordLogsRequest {
	s.EndDate = &v
	return s
}

type DescribeRecordLogsResponseBody struct {
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int64                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int64                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RecordLogs []*DescribeRecordLogsResponseBodyRecordLogs `json:"RecordLogs,omitempty" xml:"RecordLogs,omitempty" type:"Repeated"`
}

func (s DescribeRecordLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordLogsResponseBody) SetTotalCount(v int64) *DescribeRecordLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRecordLogsResponseBody) SetPageSize(v int64) *DescribeRecordLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordLogsResponseBody) SetRequestId(v string) *DescribeRecordLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordLogsResponseBody) SetPageNumber(v int64) *DescribeRecordLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRecordLogsResponseBody) SetRecordLogs(v []*DescribeRecordLogsResponseBodyRecordLogs) *DescribeRecordLogsResponseBody {
	s.RecordLogs = v
	return s
}

type DescribeRecordLogsResponseBodyRecordLogs struct {
	Action          *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionTimestamp *int64  `json:"ActionTimestamp,omitempty" xml:"ActionTimestamp,omitempty"`
	ClientIp        *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ActionTime      *string `json:"ActionTime,omitempty" xml:"ActionTime,omitempty"`
}

func (s DescribeRecordLogsResponseBodyRecordLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordLogsResponseBodyRecordLogs) GoString() string {
	return s.String()
}

func (s *DescribeRecordLogsResponseBodyRecordLogs) SetAction(v string) *DescribeRecordLogsResponseBodyRecordLogs {
	s.Action = &v
	return s
}

func (s *DescribeRecordLogsResponseBodyRecordLogs) SetActionTimestamp(v int64) *DescribeRecordLogsResponseBodyRecordLogs {
	s.ActionTimestamp = &v
	return s
}

func (s *DescribeRecordLogsResponseBodyRecordLogs) SetClientIp(v string) *DescribeRecordLogsResponseBodyRecordLogs {
	s.ClientIp = &v
	return s
}

func (s *DescribeRecordLogsResponseBodyRecordLogs) SetMessage(v string) *DescribeRecordLogsResponseBodyRecordLogs {
	s.Message = &v
	return s
}

func (s *DescribeRecordLogsResponseBodyRecordLogs) SetActionTime(v string) *DescribeRecordLogsResponseBodyRecordLogs {
	s.ActionTime = &v
	return s
}

type DescribeRecordLogsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRecordLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecordLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordLogsResponse) SetHeaders(v map[string]*string) *DescribeRecordLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordLogsResponse) SetBody(v *DescribeRecordLogsResponseBody) *DescribeRecordLogsResponse {
	s.Body = v
	return s
}

type DescribeRecordStatisticsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Rr           *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	DomainType   *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
}

func (s DescribeRecordStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsRequest) SetLang(v string) *DescribeRecordStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRecordStatisticsRequest) SetUserClientIp(v string) *DescribeRecordStatisticsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeRecordStatisticsRequest) SetStartDate(v string) *DescribeRecordStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeRecordStatisticsRequest) SetEndDate(v string) *DescribeRecordStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeRecordStatisticsRequest) SetDomainName(v string) *DescribeRecordStatisticsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeRecordStatisticsRequest) SetRr(v string) *DescribeRecordStatisticsRequest {
	s.Rr = &v
	return s
}

func (s *DescribeRecordStatisticsRequest) SetDomainType(v string) *DescribeRecordStatisticsRequest {
	s.DomainType = &v
	return s
}

type DescribeRecordStatisticsResponseBody struct {
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistics []*DescribeRecordStatisticsResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeRecordStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsResponseBody) SetRequestId(v string) *DescribeRecordStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordStatisticsResponseBody) SetStatistics(v []*DescribeRecordStatisticsResponseBodyStatistics) *DescribeRecordStatisticsResponseBody {
	s.Statistics = v
	return s
}

type DescribeRecordStatisticsResponseBodyStatistics struct {
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Count     *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeRecordStatisticsResponseBodyStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordStatisticsResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsResponseBodyStatistics) SetTimestamp(v int64) *DescribeRecordStatisticsResponseBodyStatistics {
	s.Timestamp = &v
	return s
}

func (s *DescribeRecordStatisticsResponseBodyStatistics) SetCount(v int64) *DescribeRecordStatisticsResponseBodyStatistics {
	s.Count = &v
	return s
}

type DescribeRecordStatisticsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRecordStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecordStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsResponse) SetHeaders(v map[string]*string) *DescribeRecordStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordStatisticsResponse) SetBody(v *DescribeRecordStatisticsResponseBody) *DescribeRecordStatisticsResponse {
	s.Body = v
	return s
}

type DescribeRecordStatisticsSummaryRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OrderBy      *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Direction    *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	SearchMode   *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Threshold    *int64  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	DomainType   *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
}

func (s DescribeRecordStatisticsSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordStatisticsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsSummaryRequest) SetLang(v string) *DescribeRecordStatisticsSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetUserClientIp(v string) *DescribeRecordStatisticsSummaryRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetPageNumber(v int64) *DescribeRecordStatisticsSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetPageSize(v int64) *DescribeRecordStatisticsSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetStartDate(v string) *DescribeRecordStatisticsSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetEndDate(v string) *DescribeRecordStatisticsSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetOrderBy(v string) *DescribeRecordStatisticsSummaryRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetDirection(v string) *DescribeRecordStatisticsSummaryRequest {
	s.Direction = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetDomainName(v string) *DescribeRecordStatisticsSummaryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetSearchMode(v string) *DescribeRecordStatisticsSummaryRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetKeyword(v string) *DescribeRecordStatisticsSummaryRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetThreshold(v int64) *DescribeRecordStatisticsSummaryRequest {
	s.Threshold = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetDomainType(v string) *DescribeRecordStatisticsSummaryRequest {
	s.DomainType = &v
	return s
}

type DescribeRecordStatisticsSummaryResponseBody struct {
	PageSize   *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalPages *int32                                                   `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems *int32                                                   `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	Statistics []*DescribeRecordStatisticsSummaryResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeRecordStatisticsSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordStatisticsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsSummaryResponseBody) SetPageSize(v int32) *DescribeRecordStatisticsSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBody) SetRequestId(v string) *DescribeRecordStatisticsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBody) SetPageNumber(v int32) *DescribeRecordStatisticsSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBody) SetTotalPages(v int32) *DescribeRecordStatisticsSummaryResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBody) SetTotalItems(v int32) *DescribeRecordStatisticsSummaryResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBody) SetStatistics(v []*DescribeRecordStatisticsSummaryResponseBodyStatistics) *DescribeRecordStatisticsSummaryResponseBody {
	s.Statistics = v
	return s
}

type DescribeRecordStatisticsSummaryResponseBodyStatistics struct {
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	Count     *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeRecordStatisticsSummaryResponseBodyStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordStatisticsSummaryResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsSummaryResponseBodyStatistics) SetSubDomain(v string) *DescribeRecordStatisticsSummaryResponseBodyStatistics {
	s.SubDomain = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBodyStatistics) SetCount(v int64) *DescribeRecordStatisticsSummaryResponseBodyStatistics {
	s.Count = &v
	return s
}

type DescribeRecordStatisticsSummaryResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRecordStatisticsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecordStatisticsSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordStatisticsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsSummaryResponse) SetHeaders(v map[string]*string) *DescribeRecordStatisticsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponse) SetBody(v *DescribeRecordStatisticsSummaryResponseBody) *DescribeRecordStatisticsSummaryResponse {
	s.Body = v
	return s
}

type DescribeSubDomainRecordsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	SubDomain    *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Line         *string `json:"Line,omitempty" xml:"Line,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeSubDomainRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubDomainRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainRecordsRequest) SetLang(v string) *DescribeSubDomainRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) SetUserClientIp(v string) *DescribeSubDomainRecordsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) SetSubDomain(v string) *DescribeSubDomainRecordsRequest {
	s.SubDomain = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) SetPageNumber(v int64) *DescribeSubDomainRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) SetPageSize(v int64) *DescribeSubDomainRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) SetType(v string) *DescribeSubDomainRecordsRequest {
	s.Type = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) SetLine(v string) *DescribeSubDomainRecordsRequest {
	s.Line = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) SetDomainName(v string) *DescribeSubDomainRecordsRequest {
	s.DomainName = &v
	return s
}

type DescribeSubDomainRecordsResponseBody struct {
	TotalCount    *int64                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize      *int64                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainRecords []*DescribeSubDomainRecordsResponseBodyDomainRecords `json:"DomainRecords,omitempty" xml:"DomainRecords,omitempty" type:"Repeated"`
	PageNumber    *int64                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeSubDomainRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubDomainRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainRecordsResponseBody) SetTotalCount(v int64) *DescribeSubDomainRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBody) SetPageSize(v int64) *DescribeSubDomainRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBody) SetRequestId(v string) *DescribeSubDomainRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBody) SetDomainRecords(v []*DescribeSubDomainRecordsResponseBodyDomainRecords) *DescribeSubDomainRecordsResponseBody {
	s.DomainRecords = v
	return s
}

func (s *DescribeSubDomainRecordsResponseBody) SetPageNumber(v int64) *DescribeSubDomainRecordsResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeSubDomainRecordsResponseBodyDomainRecords struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight     *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TTL        *int64  `json:"TTL,omitempty" xml:"TTL,omitempty"`
	Line       *string `json:"Line,omitempty" xml:"Line,omitempty"`
	RecordId   *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Priority   *int64  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RR         *string `json:"RR,omitempty" xml:"RR,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Locked     *bool   `json:"Locked,omitempty" xml:"Locked,omitempty"`
}

func (s DescribeSubDomainRecordsResponseBodyDomainRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubDomainRecordsResponseBodyDomainRecords) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecords) SetStatus(v string) *DescribeSubDomainRecordsResponseBodyDomainRecords {
	s.Status = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecords) SetType(v string) *DescribeSubDomainRecordsResponseBodyDomainRecords {
	s.Type = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecords) SetWeight(v int32) *DescribeSubDomainRecordsResponseBodyDomainRecords {
	s.Weight = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecords) SetValue(v string) *DescribeSubDomainRecordsResponseBodyDomainRecords {
	s.Value = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecords) SetTTL(v int64) *DescribeSubDomainRecordsResponseBodyDomainRecords {
	s.TTL = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecords) SetLine(v string) *DescribeSubDomainRecordsResponseBodyDomainRecords {
	s.Line = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecords) SetRecordId(v string) *DescribeSubDomainRecordsResponseBodyDomainRecords {
	s.RecordId = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecords) SetPriority(v int64) *DescribeSubDomainRecordsResponseBodyDomainRecords {
	s.Priority = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecords) SetRR(v string) *DescribeSubDomainRecordsResponseBodyDomainRecords {
	s.RR = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecords) SetDomainName(v string) *DescribeSubDomainRecordsResponseBodyDomainRecords {
	s.DomainName = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecords) SetLocked(v bool) *DescribeSubDomainRecordsResponseBodyDomainRecords {
	s.Locked = &v
	return s
}

type DescribeSubDomainRecordsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSubDomainRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSubDomainRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubDomainRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainRecordsResponse) SetHeaders(v map[string]*string) *DescribeSubDomainRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubDomainRecordsResponse) SetBody(v *DescribeSubDomainRecordsResponseBody) *DescribeSubDomainRecordsResponse {
	s.Body = v
	return s
}

type DescribeSupportLinesRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeSupportLinesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportLinesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupportLinesRequest) SetLang(v string) *DescribeSupportLinesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSupportLinesRequest) SetUserClientIp(v string) *DescribeSupportLinesRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeSupportLinesRequest) SetDomainName(v string) *DescribeSupportLinesRequest {
	s.DomainName = &v
	return s
}

type DescribeSupportLinesResponseBody struct {
	RecordLines []*DescribeSupportLinesResponseBodyRecordLines `json:"RecordLines,omitempty" xml:"RecordLines,omitempty" type:"Repeated"`
	RequestId   *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSupportLinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportLinesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportLinesResponseBody) SetRecordLines(v []*DescribeSupportLinesResponseBodyRecordLines) *DescribeSupportLinesResponseBody {
	s.RecordLines = v
	return s
}

func (s *DescribeSupportLinesResponseBody) SetRequestId(v string) *DescribeSupportLinesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSupportLinesResponseBodyRecordLines struct {
	FatherCode      *string `json:"FatherCode,omitempty" xml:"FatherCode,omitempty"`
	LineDisplayName *string `json:"LineDisplayName,omitempty" xml:"LineDisplayName,omitempty"`
	LineCode        *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	LineName        *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
}

func (s DescribeSupportLinesResponseBodyRecordLines) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportLinesResponseBodyRecordLines) GoString() string {
	return s.String()
}

func (s *DescribeSupportLinesResponseBodyRecordLines) SetFatherCode(v string) *DescribeSupportLinesResponseBodyRecordLines {
	s.FatherCode = &v
	return s
}

func (s *DescribeSupportLinesResponseBodyRecordLines) SetLineDisplayName(v string) *DescribeSupportLinesResponseBodyRecordLines {
	s.LineDisplayName = &v
	return s
}

func (s *DescribeSupportLinesResponseBodyRecordLines) SetLineCode(v string) *DescribeSupportLinesResponseBodyRecordLines {
	s.LineCode = &v
	return s
}

func (s *DescribeSupportLinesResponseBodyRecordLines) SetLineName(v string) *DescribeSupportLinesResponseBodyRecordLines {
	s.LineName = &v
	return s
}

type DescribeSupportLinesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSupportLinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSupportLinesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportLinesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupportLinesResponse) SetHeaders(v map[string]*string) *DescribeSupportLinesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupportLinesResponse) SetBody(v *DescribeSupportLinesResponseBody) *DescribeSupportLinesResponse {
	s.Body = v
	return s
}

type DescribeTagsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) SetLang(v string) *DescribeTagsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTagsRequest) SetUserClientIp(v string) *DescribeTagsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceType(v string) *DescribeTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagsRequest) SetPageNumber(v int64) *DescribeTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsRequest) SetPageSize(v int64) *DescribeTagsRequest {
	s.PageSize = &v
	return s
}

type DescribeTagsResponseBody struct {
	TotalCount *int64                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int64                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int64                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Tags       []*DescribeTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) SetTotalCount(v int64) *DescribeTagsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTagsResponseBody) SetPageSize(v int64) *DescribeTagsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetPageNumber(v int64) *DescribeTagsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTags(v []*DescribeTagsResponseBodyTags) *DescribeTagsResponseBody {
	s.Tags = v
	return s
}

type DescribeTagsResponseBodyTags struct {
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTags) SetKey(v string) *DescribeTagsResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeTagsResponseBodyTags) SetValues(v []*string) *DescribeTagsResponseBodyTags {
	s.Values = v
	return s
}

type DescribeTagsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponse) SetHeaders(v map[string]*string) *DescribeTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagsResponse) SetBody(v *DescribeTagsResponseBody) *DescribeTagsResponse {
	s.Body = v
	return s
}

type DescribeTransferDomainsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TransferType *string `json:"TransferType,omitempty" xml:"TransferType,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FromUserId   *int64  `json:"FromUserId,omitempty" xml:"FromUserId,omitempty"`
	TargetUserId *int64  `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s DescribeTransferDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransferDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTransferDomainsRequest) SetLang(v string) *DescribeTransferDomainsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTransferDomainsRequest) SetUserClientIp(v string) *DescribeTransferDomainsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeTransferDomainsRequest) SetPageNumber(v int64) *DescribeTransferDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTransferDomainsRequest) SetPageSize(v int64) *DescribeTransferDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTransferDomainsRequest) SetTransferType(v string) *DescribeTransferDomainsRequest {
	s.TransferType = &v
	return s
}

func (s *DescribeTransferDomainsRequest) SetDomainName(v string) *DescribeTransferDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeTransferDomainsRequest) SetFromUserId(v int64) *DescribeTransferDomainsRequest {
	s.FromUserId = &v
	return s
}

func (s *DescribeTransferDomainsRequest) SetTargetUserId(v int64) *DescribeTransferDomainsRequest {
	s.TargetUserId = &v
	return s
}

type DescribeTransferDomainsResponseBody struct {
	TotalCount      *int64                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize        *int64                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int64                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DomainTransfers []*DescribeTransferDomainsResponseBodyDomainTransfers `json:"DomainTransfers,omitempty" xml:"DomainTransfers,omitempty" type:"Repeated"`
}

func (s DescribeTransferDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransferDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTransferDomainsResponseBody) SetTotalCount(v int64) *DescribeTransferDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTransferDomainsResponseBody) SetPageSize(v int64) *DescribeTransferDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTransferDomainsResponseBody) SetRequestId(v string) *DescribeTransferDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTransferDomainsResponseBody) SetPageNumber(v int64) *DescribeTransferDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTransferDomainsResponseBody) SetDomainTransfers(v []*DescribeTransferDomainsResponseBodyDomainTransfers) *DescribeTransferDomainsResponseBody {
	s.DomainTransfers = v
	return s
}

type DescribeTransferDomainsResponseBodyDomainTransfers struct {
	FromUserId      *int64  `json:"FromUserId,omitempty" xml:"FromUserId,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TargetUserId    *int64  `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeTransferDomainsResponseBodyDomainTransfers) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransferDomainsResponseBodyDomainTransfers) GoString() string {
	return s.String()
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfers) SetFromUserId(v int64) *DescribeTransferDomainsResponseBodyDomainTransfers {
	s.FromUserId = &v
	return s
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfers) SetCreateTime(v string) *DescribeTransferDomainsResponseBodyDomainTransfers {
	s.CreateTime = &v
	return s
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfers) SetTargetUserId(v int64) *DescribeTransferDomainsResponseBodyDomainTransfers {
	s.TargetUserId = &v
	return s
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfers) SetDomainName(v string) *DescribeTransferDomainsResponseBodyDomainTransfers {
	s.DomainName = &v
	return s
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfers) SetId(v int64) *DescribeTransferDomainsResponseBodyDomainTransfers {
	s.Id = &v
	return s
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfers) SetCreateTimestamp(v int64) *DescribeTransferDomainsResponseBodyDomainTransfers {
	s.CreateTimestamp = &v
	return s
}

type DescribeTransferDomainsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTransferDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTransferDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransferDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTransferDomainsResponse) SetHeaders(v map[string]*string) *DescribeTransferDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTransferDomainsResponse) SetBody(v *DescribeTransferDomainsResponseBody) *DescribeTransferDomainsResponse {
	s.Body = v
	return s
}

type ExecuteGtmRecoveryPlanRequest struct {
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp   *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	RecoveryPlanId *int64  `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
}

func (s ExecuteGtmRecoveryPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteGtmRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *ExecuteGtmRecoveryPlanRequest) SetLang(v string) *ExecuteGtmRecoveryPlanRequest {
	s.Lang = &v
	return s
}

func (s *ExecuteGtmRecoveryPlanRequest) SetUserClientIp(v string) *ExecuteGtmRecoveryPlanRequest {
	s.UserClientIp = &v
	return s
}

func (s *ExecuteGtmRecoveryPlanRequest) SetRecoveryPlanId(v int64) *ExecuteGtmRecoveryPlanRequest {
	s.RecoveryPlanId = &v
	return s
}

type ExecuteGtmRecoveryPlanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteGtmRecoveryPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteGtmRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteGtmRecoveryPlanResponseBody) SetRequestId(v string) *ExecuteGtmRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteGtmRecoveryPlanResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteGtmRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteGtmRecoveryPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteGtmRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *ExecuteGtmRecoveryPlanResponse) SetHeaders(v map[string]*string) *ExecuteGtmRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *ExecuteGtmRecoveryPlanResponse) SetBody(v *ExecuteGtmRecoveryPlanResponseBody) *ExecuteGtmRecoveryPlanResponse {
	s.Body = v
	return s
}

type GetMainDomainNameRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InputString  *string `json:"InputString,omitempty" xml:"InputString,omitempty"`
}

func (s GetMainDomainNameRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMainDomainNameRequest) GoString() string {
	return s.String()
}

func (s *GetMainDomainNameRequest) SetLang(v string) *GetMainDomainNameRequest {
	s.Lang = &v
	return s
}

func (s *GetMainDomainNameRequest) SetUserClientIp(v string) *GetMainDomainNameRequest {
	s.UserClientIp = &v
	return s
}

func (s *GetMainDomainNameRequest) SetInputString(v string) *GetMainDomainNameRequest {
	s.InputString = &v
	return s
}

type GetMainDomainNameResponseBody struct {
	RR          *string `json:"RR,omitempty" xml:"RR,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainLevel *int64  `json:"DomainLevel,omitempty" xml:"DomainLevel,omitempty"`
}

func (s GetMainDomainNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMainDomainNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetMainDomainNameResponseBody) SetRR(v string) *GetMainDomainNameResponseBody {
	s.RR = &v
	return s
}

func (s *GetMainDomainNameResponseBody) SetRequestId(v string) *GetMainDomainNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMainDomainNameResponseBody) SetDomainName(v string) *GetMainDomainNameResponseBody {
	s.DomainName = &v
	return s
}

func (s *GetMainDomainNameResponseBody) SetDomainLevel(v int64) *GetMainDomainNameResponseBody {
	s.DomainLevel = &v
	return s
}

type GetMainDomainNameResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMainDomainNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMainDomainNameResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMainDomainNameResponse) GoString() string {
	return s.String()
}

func (s *GetMainDomainNameResponse) SetHeaders(v map[string]*string) *GetMainDomainNameResponse {
	s.Headers = v
	return s
}

func (s *GetMainDomainNameResponse) SetBody(v *GetMainDomainNameResponseBody) *GetMainDomainNameResponse {
	s.Body = v
	return s
}

type GetTxtRecordForVerifyRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s GetTxtRecordForVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTxtRecordForVerifyRequest) GoString() string {
	return s.String()
}

func (s *GetTxtRecordForVerifyRequest) SetLang(v string) *GetTxtRecordForVerifyRequest {
	s.Lang = &v
	return s
}

func (s *GetTxtRecordForVerifyRequest) SetDomainName(v string) *GetTxtRecordForVerifyRequest {
	s.DomainName = &v
	return s
}

func (s *GetTxtRecordForVerifyRequest) SetType(v string) *GetTxtRecordForVerifyRequest {
	s.Type = &v
	return s
}

func (s *GetTxtRecordForVerifyRequest) SetUserClientIp(v string) *GetTxtRecordForVerifyRequest {
	s.UserClientIp = &v
	return s
}

type GetTxtRecordForVerifyResponseBody struct {
	RR         *string `json:"RR,omitempty" xml:"RR,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTxtRecordForVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTxtRecordForVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *GetTxtRecordForVerifyResponseBody) SetRR(v string) *GetTxtRecordForVerifyResponseBody {
	s.RR = &v
	return s
}

func (s *GetTxtRecordForVerifyResponseBody) SetRequestId(v string) *GetTxtRecordForVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTxtRecordForVerifyResponseBody) SetDomainName(v string) *GetTxtRecordForVerifyResponseBody {
	s.DomainName = &v
	return s
}

func (s *GetTxtRecordForVerifyResponseBody) SetValue(v string) *GetTxtRecordForVerifyResponseBody {
	s.Value = &v
	return s
}

type GetTxtRecordForVerifyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTxtRecordForVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTxtRecordForVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTxtRecordForVerifyResponse) GoString() string {
	return s.String()
}

func (s *GetTxtRecordForVerifyResponse) SetHeaders(v map[string]*string) *GetTxtRecordForVerifyResponse {
	s.Headers = v
	return s
}

func (s *GetTxtRecordForVerifyResponse) SetBody(v *GetTxtRecordForVerifyResponseBody) *GetTxtRecordForVerifyResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	Lang         *string                       `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string                       `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Size         *int32                        `json:"Size,omitempty" xml:"Size,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetLang(v string) *ListTagResourcesRequest {
	s.Lang = &v
	return s
}

func (s *ListTagResourcesRequest) SetUserClientIp(v string) *ListTagResourcesRequest {
	s.UserClientIp = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetSize(v int32) *ListTagResourcesRequest {
	s.Size = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyHichinaDomainDNSRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s ModifyHichinaDomainDNSRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHichinaDomainDNSRequest) GoString() string {
	return s.String()
}

func (s *ModifyHichinaDomainDNSRequest) SetLang(v string) *ModifyHichinaDomainDNSRequest {
	s.Lang = &v
	return s
}

func (s *ModifyHichinaDomainDNSRequest) SetUserClientIp(v string) *ModifyHichinaDomainDNSRequest {
	s.UserClientIp = &v
	return s
}

func (s *ModifyHichinaDomainDNSRequest) SetDomainName(v string) *ModifyHichinaDomainDNSRequest {
	s.DomainName = &v
	return s
}

type ModifyHichinaDomainDNSResponseBody struct {
	RequestId          *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NewDnsServers      []*string `json:"NewDnsServers,omitempty" xml:"NewDnsServers,omitempty" type:"Repeated"`
	OriginalDnsServers []*string `json:"OriginalDnsServers,omitempty" xml:"OriginalDnsServers,omitempty" type:"Repeated"`
}

func (s ModifyHichinaDomainDNSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHichinaDomainDNSResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHichinaDomainDNSResponseBody) SetRequestId(v string) *ModifyHichinaDomainDNSResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHichinaDomainDNSResponseBody) SetNewDnsServers(v []*string) *ModifyHichinaDomainDNSResponseBody {
	s.NewDnsServers = v
	return s
}

func (s *ModifyHichinaDomainDNSResponseBody) SetOriginalDnsServers(v []*string) *ModifyHichinaDomainDNSResponseBody {
	s.OriginalDnsServers = v
	return s
}

type ModifyHichinaDomainDNSResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyHichinaDomainDNSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHichinaDomainDNSResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHichinaDomainDNSResponse) GoString() string {
	return s.String()
}

func (s *ModifyHichinaDomainDNSResponse) SetHeaders(v map[string]*string) *ModifyHichinaDomainDNSResponse {
	s.Headers = v
	return s
}

func (s *ModifyHichinaDomainDNSResponse) SetBody(v *ModifyHichinaDomainDNSResponseBody) *ModifyHichinaDomainDNSResponse {
	s.Body = v
	return s
}

type MoveDomainResourceGroupRequest struct {
	Lang               *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	UserClientIp       *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s MoveDomainResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveDomainResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveDomainResourceGroupRequest) SetLang(v string) *MoveDomainResourceGroupRequest {
	s.Lang = &v
	return s
}

func (s *MoveDomainResourceGroupRequest) SetResourceId(v string) *MoveDomainResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveDomainResourceGroupRequest) SetNewResourceGroupId(v string) *MoveDomainResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveDomainResourceGroupRequest) SetUserClientIp(v string) *MoveDomainResourceGroupRequest {
	s.UserClientIp = &v
	return s
}

type MoveDomainResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveDomainResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveDomainResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveDomainResourceGroupResponseBody) SetRequestId(v string) *MoveDomainResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveDomainResourceGroupResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveDomainResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveDomainResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveDomainResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveDomainResourceGroupResponse) SetHeaders(v map[string]*string) *MoveDomainResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveDomainResourceGroupResponse) SetBody(v *MoveDomainResourceGroupResponseBody) *MoveDomainResourceGroupResponse {
	s.Body = v
	return s
}

type MoveGtmResourceGroupRequest struct {
	Lang               *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	UserClientIp       *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s MoveGtmResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveGtmResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveGtmResourceGroupRequest) SetLang(v string) *MoveGtmResourceGroupRequest {
	s.Lang = &v
	return s
}

func (s *MoveGtmResourceGroupRequest) SetResourceId(v string) *MoveGtmResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveGtmResourceGroupRequest) SetNewResourceGroupId(v string) *MoveGtmResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveGtmResourceGroupRequest) SetUserClientIp(v string) *MoveGtmResourceGroupRequest {
	s.UserClientIp = &v
	return s
}

type MoveGtmResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveGtmResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveGtmResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveGtmResourceGroupResponseBody) SetRequestId(v string) *MoveGtmResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveGtmResourceGroupResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveGtmResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveGtmResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveGtmResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveGtmResourceGroupResponse) SetHeaders(v map[string]*string) *MoveGtmResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveGtmResourceGroupResponse) SetBody(v *MoveGtmResourceGroupResponseBody) *MoveGtmResourceGroupResponse {
	s.Body = v
	return s
}

type OperateBatchDomainRequest struct {
	Lang             *string                                      `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp     *string                                      `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Type             *string                                      `json:"Type,omitempty" xml:"Type,omitempty"`
	DomainRecordInfo []*OperateBatchDomainRequestDomainRecordInfo `json:"DomainRecordInfo,omitempty" xml:"DomainRecordInfo,omitempty" type:"Repeated"`
}

func (s OperateBatchDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateBatchDomainRequest) GoString() string {
	return s.String()
}

func (s *OperateBatchDomainRequest) SetLang(v string) *OperateBatchDomainRequest {
	s.Lang = &v
	return s
}

func (s *OperateBatchDomainRequest) SetUserClientIp(v string) *OperateBatchDomainRequest {
	s.UserClientIp = &v
	return s
}

func (s *OperateBatchDomainRequest) SetType(v string) *OperateBatchDomainRequest {
	s.Type = &v
	return s
}

func (s *OperateBatchDomainRequest) SetDomainRecordInfo(v []*OperateBatchDomainRequestDomainRecordInfo) *OperateBatchDomainRequest {
	s.DomainRecordInfo = v
	return s
}

type OperateBatchDomainRequestDomainRecordInfo struct {
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Ttl      *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Line     *string `json:"Line,omitempty" xml:"Line,omitempty"`
	NewRr    *string `json:"NewRr,omitempty" xml:"NewRr,omitempty"`
	Rr       *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	Priority *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	NewType  *string `json:"NewType,omitempty" xml:"NewType,omitempty"`
	NewValue *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
}

func (s OperateBatchDomainRequestDomainRecordInfo) String() string {
	return tea.Prettify(s)
}

func (s OperateBatchDomainRequestDomainRecordInfo) GoString() string {
	return s.String()
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetType(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.Type = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetValue(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.Value = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetTtl(v int32) *OperateBatchDomainRequestDomainRecordInfo {
	s.Ttl = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetDomain(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.Domain = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetLine(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.Line = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetNewRr(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.NewRr = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetRr(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.Rr = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetPriority(v int32) *OperateBatchDomainRequestDomainRecordInfo {
	s.Priority = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetNewType(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.NewType = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetNewValue(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.NewValue = &v
	return s
}

type OperateBatchDomainResponseBody struct {
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateBatchDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateBatchDomainResponseBody) GoString() string {
	return s.String()
}

func (s *OperateBatchDomainResponseBody) SetTaskId(v int64) *OperateBatchDomainResponseBody {
	s.TaskId = &v
	return s
}

func (s *OperateBatchDomainResponseBody) SetRequestId(v string) *OperateBatchDomainResponseBody {
	s.RequestId = &v
	return s
}

type OperateBatchDomainResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OperateBatchDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OperateBatchDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateBatchDomainResponse) GoString() string {
	return s.String()
}

func (s *OperateBatchDomainResponse) SetHeaders(v map[string]*string) *OperateBatchDomainResponse {
	s.Headers = v
	return s
}

func (s *OperateBatchDomainResponse) SetBody(v *OperateBatchDomainResponseBody) *OperateBatchDomainResponse {
	s.Body = v
	return s
}

type PreviewGtmRecoveryPlanRequest struct {
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp   *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	RecoveryPlanId *int64  `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s PreviewGtmRecoveryPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s PreviewGtmRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *PreviewGtmRecoveryPlanRequest) SetLang(v string) *PreviewGtmRecoveryPlanRequest {
	s.Lang = &v
	return s
}

func (s *PreviewGtmRecoveryPlanRequest) SetUserClientIp(v string) *PreviewGtmRecoveryPlanRequest {
	s.UserClientIp = &v
	return s
}

func (s *PreviewGtmRecoveryPlanRequest) SetRecoveryPlanId(v int64) *PreviewGtmRecoveryPlanRequest {
	s.RecoveryPlanId = &v
	return s
}

func (s *PreviewGtmRecoveryPlanRequest) SetPageNumber(v int32) *PreviewGtmRecoveryPlanRequest {
	s.PageNumber = &v
	return s
}

func (s *PreviewGtmRecoveryPlanRequest) SetPageSize(v int32) *PreviewGtmRecoveryPlanRequest {
	s.PageSize = &v
	return s
}

type PreviewGtmRecoveryPlanResponseBody struct {
	PageSize   *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Previews   []*PreviewGtmRecoveryPlanResponseBodyPreviews `json:"Previews,omitempty" xml:"Previews,omitempty" type:"Repeated"`
	TotalPages *int32                                        `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems *int32                                        `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
}

func (s PreviewGtmRecoveryPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreviewGtmRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewGtmRecoveryPlanResponseBody) SetPageSize(v int32) *PreviewGtmRecoveryPlanResponseBody {
	s.PageSize = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBody) SetRequestId(v string) *PreviewGtmRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBody) SetPageNumber(v int32) *PreviewGtmRecoveryPlanResponseBody {
	s.PageNumber = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBody) SetPreviews(v []*PreviewGtmRecoveryPlanResponseBodyPreviews) *PreviewGtmRecoveryPlanResponseBody {
	s.Previews = v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBody) SetTotalPages(v int32) *PreviewGtmRecoveryPlanResponseBody {
	s.TotalPages = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBody) SetTotalItems(v int32) *PreviewGtmRecoveryPlanResponseBody {
	s.TotalItems = &v
	return s
}

type PreviewGtmRecoveryPlanResponseBodyPreviews struct {
	InstanceId     *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SwitchInfos    []*PreviewGtmRecoveryPlanResponseBodyPreviewsSwitchInfos `json:"SwitchInfos,omitempty" xml:"SwitchInfos,omitempty" type:"Repeated"`
	Name           *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	UserDomainName *string                                                  `json:"UserDomainName,omitempty" xml:"UserDomainName,omitempty"`
}

func (s PreviewGtmRecoveryPlanResponseBodyPreviews) String() string {
	return tea.Prettify(s)
}

func (s PreviewGtmRecoveryPlanResponseBodyPreviews) GoString() string {
	return s.String()
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviews) SetInstanceId(v string) *PreviewGtmRecoveryPlanResponseBodyPreviews {
	s.InstanceId = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviews) SetSwitchInfos(v []*PreviewGtmRecoveryPlanResponseBodyPreviewsSwitchInfos) *PreviewGtmRecoveryPlanResponseBodyPreviews {
	s.SwitchInfos = v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviews) SetName(v string) *PreviewGtmRecoveryPlanResponseBodyPreviews {
	s.Name = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviews) SetUserDomainName(v string) *PreviewGtmRecoveryPlanResponseBodyPreviews {
	s.UserDomainName = &v
	return s
}

type PreviewGtmRecoveryPlanResponseBodyPreviewsSwitchInfos struct {
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s PreviewGtmRecoveryPlanResponseBodyPreviewsSwitchInfos) String() string {
	return tea.Prettify(s)
}

func (s PreviewGtmRecoveryPlanResponseBodyPreviewsSwitchInfos) GoString() string {
	return s.String()
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsSwitchInfos) SetStrategyName(v string) *PreviewGtmRecoveryPlanResponseBodyPreviewsSwitchInfos {
	s.StrategyName = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsSwitchInfos) SetContent(v string) *PreviewGtmRecoveryPlanResponseBodyPreviewsSwitchInfos {
	s.Content = &v
	return s
}

type PreviewGtmRecoveryPlanResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PreviewGtmRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PreviewGtmRecoveryPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s PreviewGtmRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *PreviewGtmRecoveryPlanResponse) SetHeaders(v map[string]*string) *PreviewGtmRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *PreviewGtmRecoveryPlanResponse) SetBody(v *PreviewGtmRecoveryPlanResponseBody) *PreviewGtmRecoveryPlanResponse {
	s.Body = v
	return s
}

type RetrieveDomainRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s RetrieveDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s RetrieveDomainRequest) GoString() string {
	return s.String()
}

func (s *RetrieveDomainRequest) SetLang(v string) *RetrieveDomainRequest {
	s.Lang = &v
	return s
}

func (s *RetrieveDomainRequest) SetUserClientIp(v string) *RetrieveDomainRequest {
	s.UserClientIp = &v
	return s
}

func (s *RetrieveDomainRequest) SetDomainName(v string) *RetrieveDomainRequest {
	s.DomainName = &v
	return s
}

type RetrieveDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetrieveDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetrieveDomainResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveDomainResponseBody) SetRequestId(v string) *RetrieveDomainResponseBody {
	s.RequestId = &v
	return s
}

type RetrieveDomainResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RetrieveDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RetrieveDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s RetrieveDomainResponse) GoString() string {
	return s.String()
}

func (s *RetrieveDomainResponse) SetHeaders(v map[string]*string) *RetrieveDomainResponse {
	s.Headers = v
	return s
}

func (s *RetrieveDomainResponse) SetBody(v *RetrieveDomainResponseBody) *RetrieveDomainResponse {
	s.Body = v
	return s
}

type RollbackGtmRecoveryPlanRequest struct {
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp   *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	RecoveryPlanId *int64  `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
}

func (s RollbackGtmRecoveryPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackGtmRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *RollbackGtmRecoveryPlanRequest) SetLang(v string) *RollbackGtmRecoveryPlanRequest {
	s.Lang = &v
	return s
}

func (s *RollbackGtmRecoveryPlanRequest) SetUserClientIp(v string) *RollbackGtmRecoveryPlanRequest {
	s.UserClientIp = &v
	return s
}

func (s *RollbackGtmRecoveryPlanRequest) SetRecoveryPlanId(v int64) *RollbackGtmRecoveryPlanRequest {
	s.RecoveryPlanId = &v
	return s
}

type RollbackGtmRecoveryPlanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackGtmRecoveryPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackGtmRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackGtmRecoveryPlanResponseBody) SetRequestId(v string) *RollbackGtmRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

type RollbackGtmRecoveryPlanResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RollbackGtmRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackGtmRecoveryPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackGtmRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *RollbackGtmRecoveryPlanResponse) SetHeaders(v map[string]*string) *RollbackGtmRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *RollbackGtmRecoveryPlanResponse) SetBody(v *RollbackGtmRecoveryPlanResponseBody) *RollbackGtmRecoveryPlanResponse {
	s.Body = v
	return s
}

type SetDnsGtmAccessModeRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	StrategyId   *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	AccessMode   *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
}

func (s SetDnsGtmAccessModeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDnsGtmAccessModeRequest) GoString() string {
	return s.String()
}

func (s *SetDnsGtmAccessModeRequest) SetLang(v string) *SetDnsGtmAccessModeRequest {
	s.Lang = &v
	return s
}

func (s *SetDnsGtmAccessModeRequest) SetUserClientIp(v string) *SetDnsGtmAccessModeRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetDnsGtmAccessModeRequest) SetStrategyId(v string) *SetDnsGtmAccessModeRequest {
	s.StrategyId = &v
	return s
}

func (s *SetDnsGtmAccessModeRequest) SetAccessMode(v string) *SetDnsGtmAccessModeRequest {
	s.AccessMode = &v
	return s
}

type SetDnsGtmAccessModeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDnsGtmAccessModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDnsGtmAccessModeResponseBody) GoString() string {
	return s.String()
}

func (s *SetDnsGtmAccessModeResponseBody) SetRequestId(v string) *SetDnsGtmAccessModeResponseBody {
	s.RequestId = &v
	return s
}

type SetDnsGtmAccessModeResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDnsGtmAccessModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDnsGtmAccessModeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDnsGtmAccessModeResponse) GoString() string {
	return s.String()
}

func (s *SetDnsGtmAccessModeResponse) SetHeaders(v map[string]*string) *SetDnsGtmAccessModeResponse {
	s.Headers = v
	return s
}

func (s *SetDnsGtmAccessModeResponse) SetBody(v *SetDnsGtmAccessModeResponseBody) *SetDnsGtmAccessModeResponse {
	s.Body = v
	return s
}

type SetDnsGtmMonitorStatusRequest struct {
	UserClientIp    *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDnsGtmMonitorStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDnsGtmMonitorStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDnsGtmMonitorStatusRequest) SetUserClientIp(v string) *SetDnsGtmMonitorStatusRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetDnsGtmMonitorStatusRequest) SetLang(v string) *SetDnsGtmMonitorStatusRequest {
	s.Lang = &v
	return s
}

func (s *SetDnsGtmMonitorStatusRequest) SetMonitorConfigId(v string) *SetDnsGtmMonitorStatusRequest {
	s.MonitorConfigId = &v
	return s
}

func (s *SetDnsGtmMonitorStatusRequest) SetStatus(v string) *SetDnsGtmMonitorStatusRequest {
	s.Status = &v
	return s
}

type SetDnsGtmMonitorStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDnsGtmMonitorStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDnsGtmMonitorStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDnsGtmMonitorStatusResponseBody) SetRequestId(v string) *SetDnsGtmMonitorStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetDnsGtmMonitorStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDnsGtmMonitorStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDnsGtmMonitorStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDnsGtmMonitorStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDnsGtmMonitorStatusResponse) SetHeaders(v map[string]*string) *SetDnsGtmMonitorStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDnsGtmMonitorStatusResponse) SetBody(v *SetDnsGtmMonitorStatusResponseBody) *SetDnsGtmMonitorStatusResponse {
	s.Body = v
	return s
}

type SetDNSSLBStatusRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	SubDomain    *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	Open         *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SetDNSSLBStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDNSSLBStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDNSSLBStatusRequest) SetLang(v string) *SetDNSSLBStatusRequest {
	s.Lang = &v
	return s
}

func (s *SetDNSSLBStatusRequest) SetUserClientIp(v string) *SetDNSSLBStatusRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetDNSSLBStatusRequest) SetSubDomain(v string) *SetDNSSLBStatusRequest {
	s.SubDomain = &v
	return s
}

func (s *SetDNSSLBStatusRequest) SetOpen(v bool) *SetDNSSLBStatusRequest {
	s.Open = &v
	return s
}

func (s *SetDNSSLBStatusRequest) SetDomainName(v string) *SetDNSSLBStatusRequest {
	s.DomainName = &v
	return s
}

func (s *SetDNSSLBStatusRequest) SetType(v string) *SetDNSSLBStatusRequest {
	s.Type = &v
	return s
}

type SetDNSSLBStatusResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordCount *int64  `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	Open        *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s SetDNSSLBStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDNSSLBStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDNSSLBStatusResponseBody) SetRequestId(v string) *SetDNSSLBStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDNSSLBStatusResponseBody) SetRecordCount(v int64) *SetDNSSLBStatusResponseBody {
	s.RecordCount = &v
	return s
}

func (s *SetDNSSLBStatusResponseBody) SetOpen(v bool) *SetDNSSLBStatusResponseBody {
	s.Open = &v
	return s
}

type SetDNSSLBStatusResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDNSSLBStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDNSSLBStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDNSSLBStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDNSSLBStatusResponse) SetHeaders(v map[string]*string) *SetDNSSLBStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDNSSLBStatusResponse) SetBody(v *SetDNSSLBStatusResponseBody) *SetDNSSLBStatusResponse {
	s.Body = v
	return s
}

type SetDomainDnssecStatusRequest struct {
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDomainDnssecStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDomainDnssecStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDomainDnssecStatusRequest) SetUserClientIp(v string) *SetDomainDnssecStatusRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetDomainDnssecStatusRequest) SetLang(v string) *SetDomainDnssecStatusRequest {
	s.Lang = &v
	return s
}

func (s *SetDomainDnssecStatusRequest) SetDomainName(v string) *SetDomainDnssecStatusRequest {
	s.DomainName = &v
	return s
}

func (s *SetDomainDnssecStatusRequest) SetStatus(v string) *SetDomainDnssecStatusRequest {
	s.Status = &v
	return s
}

type SetDomainDnssecStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDomainDnssecStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDomainDnssecStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainDnssecStatusResponseBody) SetRequestId(v string) *SetDomainDnssecStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetDomainDnssecStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDomainDnssecStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDomainDnssecStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDomainDnssecStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDomainDnssecStatusResponse) SetHeaders(v map[string]*string) *SetDomainDnssecStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDomainDnssecStatusResponse) SetBody(v *SetDomainDnssecStatusResponseBody) *SetDomainDnssecStatusResponse {
	s.Body = v
	return s
}

type SetDomainRecordStatusRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	RecordId     *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDomainRecordStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDomainRecordStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDomainRecordStatusRequest) SetLang(v string) *SetDomainRecordStatusRequest {
	s.Lang = &v
	return s
}

func (s *SetDomainRecordStatusRequest) SetUserClientIp(v string) *SetDomainRecordStatusRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetDomainRecordStatusRequest) SetRecordId(v string) *SetDomainRecordStatusRequest {
	s.RecordId = &v
	return s
}

func (s *SetDomainRecordStatusRequest) SetStatus(v string) *SetDomainRecordStatusRequest {
	s.Status = &v
	return s
}

type SetDomainRecordStatusResponseBody struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s SetDomainRecordStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDomainRecordStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainRecordStatusResponseBody) SetStatus(v string) *SetDomainRecordStatusResponseBody {
	s.Status = &v
	return s
}

func (s *SetDomainRecordStatusResponseBody) SetRequestId(v string) *SetDomainRecordStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDomainRecordStatusResponseBody) SetRecordId(v string) *SetDomainRecordStatusResponseBody {
	s.RecordId = &v
	return s
}

type SetDomainRecordStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDomainRecordStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDomainRecordStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDomainRecordStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDomainRecordStatusResponse) SetHeaders(v map[string]*string) *SetDomainRecordStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDomainRecordStatusResponse) SetBody(v *SetDomainRecordStatusResponseBody) *SetDomainRecordStatusResponse {
	s.Body = v
	return s
}

type SetGtmAccessModeRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	StrategyId   *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	AccessMode   *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
}

func (s SetGtmAccessModeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGtmAccessModeRequest) GoString() string {
	return s.String()
}

func (s *SetGtmAccessModeRequest) SetLang(v string) *SetGtmAccessModeRequest {
	s.Lang = &v
	return s
}

func (s *SetGtmAccessModeRequest) SetUserClientIp(v string) *SetGtmAccessModeRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetGtmAccessModeRequest) SetStrategyId(v string) *SetGtmAccessModeRequest {
	s.StrategyId = &v
	return s
}

func (s *SetGtmAccessModeRequest) SetAccessMode(v string) *SetGtmAccessModeRequest {
	s.AccessMode = &v
	return s
}

type SetGtmAccessModeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGtmAccessModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGtmAccessModeResponseBody) GoString() string {
	return s.String()
}

func (s *SetGtmAccessModeResponseBody) SetRequestId(v string) *SetGtmAccessModeResponseBody {
	s.RequestId = &v
	return s
}

type SetGtmAccessModeResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetGtmAccessModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGtmAccessModeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGtmAccessModeResponse) GoString() string {
	return s.String()
}

func (s *SetGtmAccessModeResponse) SetHeaders(v map[string]*string) *SetGtmAccessModeResponse {
	s.Headers = v
	return s
}

func (s *SetGtmAccessModeResponse) SetBody(v *SetGtmAccessModeResponseBody) *SetGtmAccessModeResponse {
	s.Body = v
	return s
}

type SetGtmMonitorStatusRequest struct {
	UserClientIp    *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetGtmMonitorStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGtmMonitorStatusRequest) GoString() string {
	return s.String()
}

func (s *SetGtmMonitorStatusRequest) SetUserClientIp(v string) *SetGtmMonitorStatusRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetGtmMonitorStatusRequest) SetLang(v string) *SetGtmMonitorStatusRequest {
	s.Lang = &v
	return s
}

func (s *SetGtmMonitorStatusRequest) SetMonitorConfigId(v string) *SetGtmMonitorStatusRequest {
	s.MonitorConfigId = &v
	return s
}

func (s *SetGtmMonitorStatusRequest) SetStatus(v string) *SetGtmMonitorStatusRequest {
	s.Status = &v
	return s
}

type SetGtmMonitorStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGtmMonitorStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGtmMonitorStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetGtmMonitorStatusResponseBody) SetRequestId(v string) *SetGtmMonitorStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetGtmMonitorStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetGtmMonitorStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGtmMonitorStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGtmMonitorStatusResponse) GoString() string {
	return s.String()
}

func (s *SetGtmMonitorStatusResponse) SetHeaders(v map[string]*string) *SetGtmMonitorStatusResponse {
	s.Headers = v
	return s
}

func (s *SetGtmMonitorStatusResponse) SetBody(v *SetGtmMonitorStatusResponseBody) *SetGtmMonitorStatusResponse {
	s.Body = v
	return s
}

type SwitchDnsGtmInstanceStrategyModeRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StrategyMode *string `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
}

func (s SwitchDnsGtmInstanceStrategyModeRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchDnsGtmInstanceStrategyModeRequest) GoString() string {
	return s.String()
}

func (s *SwitchDnsGtmInstanceStrategyModeRequest) SetLang(v string) *SwitchDnsGtmInstanceStrategyModeRequest {
	s.Lang = &v
	return s
}

func (s *SwitchDnsGtmInstanceStrategyModeRequest) SetUserClientIp(v string) *SwitchDnsGtmInstanceStrategyModeRequest {
	s.UserClientIp = &v
	return s
}

func (s *SwitchDnsGtmInstanceStrategyModeRequest) SetInstanceId(v string) *SwitchDnsGtmInstanceStrategyModeRequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchDnsGtmInstanceStrategyModeRequest) SetStrategyMode(v string) *SwitchDnsGtmInstanceStrategyModeRequest {
	s.StrategyMode = &v
	return s
}

type SwitchDnsGtmInstanceStrategyModeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchDnsGtmInstanceStrategyModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchDnsGtmInstanceStrategyModeResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchDnsGtmInstanceStrategyModeResponseBody) SetRequestId(v string) *SwitchDnsGtmInstanceStrategyModeResponseBody {
	s.RequestId = &v
	return s
}

type SwitchDnsGtmInstanceStrategyModeResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SwitchDnsGtmInstanceStrategyModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchDnsGtmInstanceStrategyModeResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchDnsGtmInstanceStrategyModeResponse) GoString() string {
	return s.String()
}

func (s *SwitchDnsGtmInstanceStrategyModeResponse) SetHeaders(v map[string]*string) *SwitchDnsGtmInstanceStrategyModeResponse {
	s.Headers = v
	return s
}

func (s *SwitchDnsGtmInstanceStrategyModeResponse) SetBody(v *SwitchDnsGtmInstanceStrategyModeResponseBody) *SwitchDnsGtmInstanceStrategyModeResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	Lang         *string                   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string                   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	OverWrite    *bool                     `json:"OverWrite,omitempty" xml:"OverWrite,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetLang(v string) *TagResourcesRequest {
	s.Lang = &v
	return s
}

func (s *TagResourcesRequest) SetUserClientIp(v string) *TagResourcesRequest {
	s.UserClientIp = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetOverWrite(v bool) *TagResourcesRequest {
	s.OverWrite = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type TransferDomainRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DomainNames  *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TargetUserId *int64  `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s TransferDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferDomainRequest) GoString() string {
	return s.String()
}

func (s *TransferDomainRequest) SetLang(v string) *TransferDomainRequest {
	s.Lang = &v
	return s
}

func (s *TransferDomainRequest) SetDomainNames(v string) *TransferDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *TransferDomainRequest) SetRemark(v string) *TransferDomainRequest {
	s.Remark = &v
	return s
}

func (s *TransferDomainRequest) SetTargetUserId(v int64) *TransferDomainRequest {
	s.TargetUserId = &v
	return s
}

func (s *TransferDomainRequest) SetUserClientIp(v string) *TransferDomainRequest {
	s.UserClientIp = &v
	return s
}

type TransferDomainResponseBody struct {
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransferDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferDomainResponseBody) GoString() string {
	return s.String()
}

func (s *TransferDomainResponseBody) SetTaskId(v int64) *TransferDomainResponseBody {
	s.TaskId = &v
	return s
}

func (s *TransferDomainResponseBody) SetRequestId(v string) *TransferDomainResponseBody {
	s.RequestId = &v
	return s
}

type TransferDomainResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransferDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferDomainResponse) GoString() string {
	return s.String()
}

func (s *TransferDomainResponse) SetHeaders(v map[string]*string) *TransferDomainResponse {
	s.Headers = v
	return s
}

func (s *TransferDomainResponse) SetBody(v *TransferDomainResponseBody) *TransferDomainResponse {
	s.Body = v
	return s
}

type UnbindInstanceDomainsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainNames  *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UnbindInstanceDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindInstanceDomainsRequest) GoString() string {
	return s.String()
}

func (s *UnbindInstanceDomainsRequest) SetLang(v string) *UnbindInstanceDomainsRequest {
	s.Lang = &v
	return s
}

func (s *UnbindInstanceDomainsRequest) SetUserClientIp(v string) *UnbindInstanceDomainsRequest {
	s.UserClientIp = &v
	return s
}

func (s *UnbindInstanceDomainsRequest) SetDomainNames(v string) *UnbindInstanceDomainsRequest {
	s.DomainNames = &v
	return s
}

func (s *UnbindInstanceDomainsRequest) SetInstanceId(v string) *UnbindInstanceDomainsRequest {
	s.InstanceId = &v
	return s
}

type UnbindInstanceDomainsResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FailedCount  *int32  `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	SuccessCount *int32  `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s UnbindInstanceDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindInstanceDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindInstanceDomainsResponseBody) SetRequestId(v string) *UnbindInstanceDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindInstanceDomainsResponseBody) SetFailedCount(v int32) *UnbindInstanceDomainsResponseBody {
	s.FailedCount = &v
	return s
}

func (s *UnbindInstanceDomainsResponseBody) SetSuccessCount(v int32) *UnbindInstanceDomainsResponseBody {
	s.SuccessCount = &v
	return s
}

type UnbindInstanceDomainsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindInstanceDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindInstanceDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindInstanceDomainsResponse) GoString() string {
	return s.String()
}

func (s *UnbindInstanceDomainsResponse) SetHeaders(v map[string]*string) *UnbindInstanceDomainsResponse {
	s.Headers = v
	return s
}

func (s *UnbindInstanceDomainsResponse) SetBody(v *UnbindInstanceDomainsResponseBody) *UnbindInstanceDomainsResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	Lang         *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetLang(v string) *UntagResourcesRequest {
	s.Lang = &v
	return s
}

func (s *UntagResourcesRequest) SetUserClientIp(v string) *UntagResourcesRequest {
	s.UserClientIp = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateCustomLineRequest struct {
	Lang         *string                             `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string                             `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	LineName     *string                             `json:"LineName,omitempty" xml:"LineName,omitempty"`
	LineId       *int64                              `json:"LineId,omitempty" xml:"LineId,omitempty"`
	IpSegment    []*UpdateCustomLineRequestIpSegment `json:"IpSegment,omitempty" xml:"IpSegment,omitempty" type:"Repeated"`
}

func (s UpdateCustomLineRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomLineRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomLineRequest) SetLang(v string) *UpdateCustomLineRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCustomLineRequest) SetUserClientIp(v string) *UpdateCustomLineRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateCustomLineRequest) SetLineName(v string) *UpdateCustomLineRequest {
	s.LineName = &v
	return s
}

func (s *UpdateCustomLineRequest) SetLineId(v int64) *UpdateCustomLineRequest {
	s.LineId = &v
	return s
}

func (s *UpdateCustomLineRequest) SetIpSegment(v []*UpdateCustomLineRequestIpSegment) *UpdateCustomLineRequest {
	s.IpSegment = v
	return s
}

type UpdateCustomLineRequestIpSegment struct {
	EndIp   *string `json:"EndIp,omitempty" xml:"EndIp,omitempty"`
	StartIp *string `json:"StartIp,omitempty" xml:"StartIp,omitempty"`
}

func (s UpdateCustomLineRequestIpSegment) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomLineRequestIpSegment) GoString() string {
	return s.String()
}

func (s *UpdateCustomLineRequestIpSegment) SetEndIp(v string) *UpdateCustomLineRequestIpSegment {
	s.EndIp = &v
	return s
}

func (s *UpdateCustomLineRequestIpSegment) SetStartIp(v string) *UpdateCustomLineRequestIpSegment {
	s.StartIp = &v
	return s
}

type UpdateCustomLineResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomLineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomLineResponseBody) SetRequestId(v string) *UpdateCustomLineResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCustomLineResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCustomLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCustomLineResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomLineResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomLineResponse) SetHeaders(v map[string]*string) *UpdateCustomLineResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomLineResponse) SetBody(v *UpdateCustomLineResponseBody) *UpdateCustomLineResponse {
	s.Body = v
	return s
}

type UpdateDnsGtmAccessStrategyRequest struct {
	Lang                        *string                                              `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp                *string                                              `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	StrategyId                  *string                                              `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	StrategyName                *string                                              `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	Lines                       *string                                              `json:"Lines,omitempty" xml:"Lines,omitempty"`
	DefaultAddrPoolType         *string                                              `json:"DefaultAddrPoolType,omitempty" xml:"DefaultAddrPoolType,omitempty"`
	DefaultLbaStrategy          *string                                              `json:"DefaultLbaStrategy,omitempty" xml:"DefaultLbaStrategy,omitempty"`
	DefaultMinAvailableAddrNum  *int32                                               `json:"DefaultMinAvailableAddrNum,omitempty" xml:"DefaultMinAvailableAddrNum,omitempty"`
	DefaultMaxReturnAddrNum     *int32                                               `json:"DefaultMaxReturnAddrNum,omitempty" xml:"DefaultMaxReturnAddrNum,omitempty"`
	DefaultLatencyOptimization  *string                                              `json:"DefaultLatencyOptimization,omitempty" xml:"DefaultLatencyOptimization,omitempty"`
	FailoverAddrPoolType        *string                                              `json:"FailoverAddrPoolType,omitempty" xml:"FailoverAddrPoolType,omitempty"`
	FailoverLbaStrategy         *string                                              `json:"FailoverLbaStrategy,omitempty" xml:"FailoverLbaStrategy,omitempty"`
	FailoverMinAvailableAddrNum *int32                                               `json:"FailoverMinAvailableAddrNum,omitempty" xml:"FailoverMinAvailableAddrNum,omitempty"`
	FailoverMaxReturnAddrNum    *int32                                               `json:"FailoverMaxReturnAddrNum,omitempty" xml:"FailoverMaxReturnAddrNum,omitempty"`
	FailoverLatencyOptimization *string                                              `json:"FailoverLatencyOptimization,omitempty" xml:"FailoverLatencyOptimization,omitempty"`
	DefaultAddrPool             []*UpdateDnsGtmAccessStrategyRequestDefaultAddrPool  `json:"DefaultAddrPool,omitempty" xml:"DefaultAddrPool,omitempty" type:"Repeated"`
	FailoverAddrPool            []*UpdateDnsGtmAccessStrategyRequestFailoverAddrPool `json:"FailoverAddrPool,omitempty" xml:"FailoverAddrPool,omitempty" type:"Repeated"`
}

func (s UpdateDnsGtmAccessStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetLang(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetUserClientIp(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetStrategyId(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetStrategyName(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.StrategyName = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetLines(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.Lines = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetDefaultAddrPoolType(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.DefaultAddrPoolType = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetDefaultLbaStrategy(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.DefaultLbaStrategy = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetDefaultMinAvailableAddrNum(v int32) *UpdateDnsGtmAccessStrategyRequest {
	s.DefaultMinAvailableAddrNum = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetDefaultMaxReturnAddrNum(v int32) *UpdateDnsGtmAccessStrategyRequest {
	s.DefaultMaxReturnAddrNum = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetDefaultLatencyOptimization(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.DefaultLatencyOptimization = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetFailoverAddrPoolType(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.FailoverAddrPoolType = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetFailoverLbaStrategy(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.FailoverLbaStrategy = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetFailoverMinAvailableAddrNum(v int32) *UpdateDnsGtmAccessStrategyRequest {
	s.FailoverMinAvailableAddrNum = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetFailoverMaxReturnAddrNum(v int32) *UpdateDnsGtmAccessStrategyRequest {
	s.FailoverMaxReturnAddrNum = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetFailoverLatencyOptimization(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.FailoverLatencyOptimization = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetDefaultAddrPool(v []*UpdateDnsGtmAccessStrategyRequestDefaultAddrPool) *UpdateDnsGtmAccessStrategyRequest {
	s.DefaultAddrPool = v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetFailoverAddrPool(v []*UpdateDnsGtmAccessStrategyRequestFailoverAddrPool) *UpdateDnsGtmAccessStrategyRequest {
	s.FailoverAddrPool = v
	return s
}

type UpdateDnsGtmAccessStrategyRequestDefaultAddrPool struct {
	LbaWeight *int32  `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateDnsGtmAccessStrategyRequestDefaultAddrPool) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmAccessStrategyRequestDefaultAddrPool) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAccessStrategyRequestDefaultAddrPool) SetLbaWeight(v int32) *UpdateDnsGtmAccessStrategyRequestDefaultAddrPool {
	s.LbaWeight = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequestDefaultAddrPool) SetId(v string) *UpdateDnsGtmAccessStrategyRequestDefaultAddrPool {
	s.Id = &v
	return s
}

type UpdateDnsGtmAccessStrategyRequestFailoverAddrPool struct {
	LbaWeight *int32  `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateDnsGtmAccessStrategyRequestFailoverAddrPool) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmAccessStrategyRequestFailoverAddrPool) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAccessStrategyRequestFailoverAddrPool) SetLbaWeight(v int32) *UpdateDnsGtmAccessStrategyRequestFailoverAddrPool {
	s.LbaWeight = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequestFailoverAddrPool) SetId(v string) *UpdateDnsGtmAccessStrategyRequestFailoverAddrPool {
	s.Id = &v
	return s
}

type UpdateDnsGtmAccessStrategyResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s UpdateDnsGtmAccessStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAccessStrategyResponseBody) SetRequestId(v string) *UpdateDnsGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyResponseBody) SetStrategyId(v string) *UpdateDnsGtmAccessStrategyResponseBody {
	s.StrategyId = &v
	return s
}

type UpdateDnsGtmAccessStrategyResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDnsGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDnsGtmAccessStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *UpdateDnsGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateDnsGtmAccessStrategyResponse) SetBody(v *UpdateDnsGtmAccessStrategyResponseBody) *UpdateDnsGtmAccessStrategyResponse {
	s.Body = v
	return s
}

type UpdateDnsGtmAddressPoolRequest struct {
	UserClientIp *string                               `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang         *string                               `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AddrPoolId   *string                               `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	Name         *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	LbaStrategy  *string                               `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	Addr         []*UpdateDnsGtmAddressPoolRequestAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Repeated"`
}

func (s UpdateDnsGtmAddressPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAddressPoolRequest) SetUserClientIp(v string) *UpdateDnsGtmAddressPoolRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequest) SetLang(v string) *UpdateDnsGtmAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequest) SetAddrPoolId(v string) *UpdateDnsGtmAddressPoolRequest {
	s.AddrPoolId = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequest) SetName(v string) *UpdateDnsGtmAddressPoolRequest {
	s.Name = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequest) SetLbaStrategy(v string) *UpdateDnsGtmAddressPoolRequest {
	s.LbaStrategy = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequest) SetAddr(v []*UpdateDnsGtmAddressPoolRequestAddr) *UpdateDnsGtmAddressPoolRequest {
	s.Addr = v
	return s
}

type UpdateDnsGtmAddressPoolRequestAddr struct {
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	LbaWeight     *int32  `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	Addr          *string `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Mode          *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s UpdateDnsGtmAddressPoolRequestAddr) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmAddressPoolRequestAddr) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) SetAttributeInfo(v string) *UpdateDnsGtmAddressPoolRequestAddr {
	s.AttributeInfo = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) SetRemark(v string) *UpdateDnsGtmAddressPoolRequestAddr {
	s.Remark = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) SetLbaWeight(v int32) *UpdateDnsGtmAddressPoolRequestAddr {
	s.LbaWeight = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) SetAddr(v string) *UpdateDnsGtmAddressPoolRequestAddr {
	s.Addr = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) SetMode(v string) *UpdateDnsGtmAddressPoolRequestAddr {
	s.Mode = &v
	return s
}

type UpdateDnsGtmAddressPoolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDnsGtmAddressPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAddressPoolResponseBody) SetRequestId(v string) *UpdateDnsGtmAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDnsGtmAddressPoolResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDnsGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDnsGtmAddressPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAddressPoolResponse) SetHeaders(v map[string]*string) *UpdateDnsGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *UpdateDnsGtmAddressPoolResponse) SetBody(v *UpdateDnsGtmAddressPoolResponseBody) *UpdateDnsGtmAddressPoolResponse {
	s.Body = v
	return s
}

type UpdateDnsGtmInstanceGlobalConfigRequest struct {
	Lang                 *string                                               `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp         *string                                               `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId           *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName         *string                                               `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Ttl                  *int32                                                `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	PublicCnameMode      *string                                               `json:"PublicCnameMode,omitempty" xml:"PublicCnameMode,omitempty"`
	PublicUserDomainName *string                                               `json:"PublicUserDomainName,omitempty" xml:"PublicUserDomainName,omitempty"`
	PublicZoneName       *string                                               `json:"PublicZoneName,omitempty" xml:"PublicZoneName,omitempty"`
	AlertGroup           *string                                               `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	CnameType            *string                                               `json:"CnameType,omitempty" xml:"CnameType,omitempty"`
	AlertConfig          []*UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
}

func (s UpdateDnsGtmInstanceGlobalConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmInstanceGlobalConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetLang(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetUserClientIp(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetInstanceId(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetInstanceName(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetTtl(v int32) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetPublicCnameMode(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.PublicCnameMode = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetPublicUserDomainName(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.PublicUserDomainName = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetPublicZoneName(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.PublicZoneName = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetAlertGroup(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.AlertGroup = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetCnameType(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.CnameType = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetAlertConfig(v []*UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.AlertConfig = v
	return s
}

type UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig struct {
	SmsNotice   *bool   `json:"SmsNotice,omitempty" xml:"SmsNotice,omitempty"`
	NoticeType  *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	EmailNotice *bool   `json:"EmailNotice,omitempty" xml:"EmailNotice,omitempty"`
}

func (s UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) SetSmsNotice(v bool) *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig {
	s.SmsNotice = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) SetNoticeType(v string) *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig {
	s.NoticeType = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) SetEmailNotice(v bool) *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig {
	s.EmailNotice = &v
	return s
}

type UpdateDnsGtmInstanceGlobalConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDnsGtmInstanceGlobalConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmInstanceGlobalConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmInstanceGlobalConfigResponseBody) SetRequestId(v string) *UpdateDnsGtmInstanceGlobalConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDnsGtmInstanceGlobalConfigResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDnsGtmInstanceGlobalConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDnsGtmInstanceGlobalConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmInstanceGlobalConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmInstanceGlobalConfigResponse) SetHeaders(v map[string]*string) *UpdateDnsGtmInstanceGlobalConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigResponse) SetBody(v *UpdateDnsGtmInstanceGlobalConfigResponseBody) *UpdateDnsGtmInstanceGlobalConfigResponse {
	s.Body = v
	return s
}

type UpdateDnsGtmMonitorRequest struct {
	UserClientIp      *string                                  `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang              *string                                  `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MonitorConfigId   *string                                  `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	ProtocolType      *string                                  `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Interval          *int32                                   `json:"Interval,omitempty" xml:"Interval,omitempty"`
	EvaluationCount   *int32                                   `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Timeout           *int32                                   `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	MonitorExtendInfo *string                                  `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
	IspCityNode       []*UpdateDnsGtmMonitorRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
}

func (s UpdateDnsGtmMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmMonitorRequest) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmMonitorRequest) SetUserClientIp(v string) *UpdateDnsGtmMonitorRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetLang(v string) *UpdateDnsGtmMonitorRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetMonitorConfigId(v string) *UpdateDnsGtmMonitorRequest {
	s.MonitorConfigId = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetProtocolType(v string) *UpdateDnsGtmMonitorRequest {
	s.ProtocolType = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetInterval(v int32) *UpdateDnsGtmMonitorRequest {
	s.Interval = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetEvaluationCount(v int32) *UpdateDnsGtmMonitorRequest {
	s.EvaluationCount = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetTimeout(v int32) *UpdateDnsGtmMonitorRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetMonitorExtendInfo(v string) *UpdateDnsGtmMonitorRequest {
	s.MonitorExtendInfo = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetIspCityNode(v []*UpdateDnsGtmMonitorRequestIspCityNode) *UpdateDnsGtmMonitorRequest {
	s.IspCityNode = v
	return s
}

type UpdateDnsGtmMonitorRequestIspCityNode struct {
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	IspCode  *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
}

func (s UpdateDnsGtmMonitorRequestIspCityNode) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmMonitorRequestIspCityNode) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmMonitorRequestIspCityNode) SetCityCode(v string) *UpdateDnsGtmMonitorRequestIspCityNode {
	s.CityCode = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequestIspCityNode) SetIspCode(v string) *UpdateDnsGtmMonitorRequestIspCityNode {
	s.IspCode = &v
	return s
}

type UpdateDnsGtmMonitorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDnsGtmMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmMonitorResponseBody) SetRequestId(v string) *UpdateDnsGtmMonitorResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDnsGtmMonitorResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDnsGtmMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDnsGtmMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDnsGtmMonitorResponse) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmMonitorResponse) SetHeaders(v map[string]*string) *UpdateDnsGtmMonitorResponse {
	s.Headers = v
	return s
}

func (s *UpdateDnsGtmMonitorResponse) SetBody(v *UpdateDnsGtmMonitorResponseBody) *UpdateDnsGtmMonitorResponse {
	s.Body = v
	return s
}

type UpdateDNSSLBWeightRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	RecordId     *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Weight       *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateDNSSLBWeightRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDNSSLBWeightRequest) GoString() string {
	return s.String()
}

func (s *UpdateDNSSLBWeightRequest) SetLang(v string) *UpdateDNSSLBWeightRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDNSSLBWeightRequest) SetUserClientIp(v string) *UpdateDNSSLBWeightRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateDNSSLBWeightRequest) SetRecordId(v string) *UpdateDNSSLBWeightRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateDNSSLBWeightRequest) SetWeight(v int32) *UpdateDNSSLBWeightRequest {
	s.Weight = &v
	return s
}

type UpdateDNSSLBWeightResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Weight    *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateDNSSLBWeightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDNSSLBWeightResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDNSSLBWeightResponseBody) SetRequestId(v string) *UpdateDNSSLBWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDNSSLBWeightResponseBody) SetRecordId(v string) *UpdateDNSSLBWeightResponseBody {
	s.RecordId = &v
	return s
}

func (s *UpdateDNSSLBWeightResponseBody) SetWeight(v int32) *UpdateDNSSLBWeightResponseBody {
	s.Weight = &v
	return s
}

type UpdateDNSSLBWeightResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDNSSLBWeightResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDNSSLBWeightResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDNSSLBWeightResponse) GoString() string {
	return s.String()
}

func (s *UpdateDNSSLBWeightResponse) SetHeaders(v map[string]*string) *UpdateDNSSLBWeightResponse {
	s.Headers = v
	return s
}

func (s *UpdateDNSSLBWeightResponse) SetBody(v *UpdateDNSSLBWeightResponseBody) *UpdateDNSSLBWeightResponse {
	s.Body = v
	return s
}

type UpdateDomainGroupRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s UpdateDomainGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainGroupRequest) SetLang(v string) *UpdateDomainGroupRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDomainGroupRequest) SetUserClientIp(v string) *UpdateDomainGroupRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateDomainGroupRequest) SetGroupId(v string) *UpdateDomainGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateDomainGroupRequest) SetGroupName(v string) *UpdateDomainGroupRequest {
	s.GroupName = &v
	return s
}

type UpdateDomainGroupResponseBody struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s UpdateDomainGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainGroupResponseBody) SetGroupName(v string) *UpdateDomainGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *UpdateDomainGroupResponseBody) SetRequestId(v string) *UpdateDomainGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDomainGroupResponseBody) SetGroupId(v string) *UpdateDomainGroupResponseBody {
	s.GroupId = &v
	return s
}

type UpdateDomainGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDomainGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDomainGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainGroupResponse) SetHeaders(v map[string]*string) *UpdateDomainGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainGroupResponse) SetBody(v *UpdateDomainGroupResponseBody) *UpdateDomainGroupResponse {
	s.Body = v
	return s
}

type UpdateDomainRecordRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	RecordId     *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RR           *string `json:"RR,omitempty" xml:"RR,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TTL          *int64  `json:"TTL,omitempty" xml:"TTL,omitempty"`
	Priority     *int64  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Line         *string `json:"Line,omitempty" xml:"Line,omitempty"`
}

func (s UpdateDomainRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainRecordRequest) SetLang(v string) *UpdateDomainRecordRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetUserClientIp(v string) *UpdateDomainRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetRecordId(v string) *UpdateDomainRecordRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetRR(v string) *UpdateDomainRecordRequest {
	s.RR = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetType(v string) *UpdateDomainRecordRequest {
	s.Type = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetValue(v string) *UpdateDomainRecordRequest {
	s.Value = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetTTL(v int64) *UpdateDomainRecordRequest {
	s.TTL = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetPriority(v int64) *UpdateDomainRecordRequest {
	s.Priority = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetLine(v string) *UpdateDomainRecordRequest {
	s.Line = &v
	return s
}

type UpdateDomainRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s UpdateDomainRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainRecordResponseBody) SetRequestId(v string) *UpdateDomainRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDomainRecordResponseBody) SetRecordId(v string) *UpdateDomainRecordResponseBody {
	s.RecordId = &v
	return s
}

type UpdateDomainRecordResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDomainRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDomainRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainRecordResponse) SetHeaders(v map[string]*string) *UpdateDomainRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainRecordResponse) SetBody(v *UpdateDomainRecordResponseBody) *UpdateDomainRecordResponse {
	s.Body = v
	return s
}

type UpdateDomainRecordRemarkRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	RecordId     *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateDomainRecordRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainRecordRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainRecordRemarkRequest) SetLang(v string) *UpdateDomainRecordRemarkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDomainRecordRemarkRequest) SetUserClientIp(v string) *UpdateDomainRecordRemarkRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateDomainRecordRemarkRequest) SetRecordId(v string) *UpdateDomainRecordRemarkRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateDomainRecordRemarkRequest) SetRemark(v string) *UpdateDomainRecordRemarkRequest {
	s.Remark = &v
	return s
}

type UpdateDomainRecordRemarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDomainRecordRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainRecordRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainRecordRemarkResponseBody) SetRequestId(v string) *UpdateDomainRecordRemarkResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDomainRecordRemarkResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDomainRecordRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDomainRecordRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainRecordRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainRecordRemarkResponse) SetHeaders(v map[string]*string) *UpdateDomainRecordRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainRecordRemarkResponse) SetBody(v *UpdateDomainRecordRemarkResponseBody) *UpdateDomainRecordRemarkResponse {
	s.Body = v
	return s
}

type UpdateDomainRemarkRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateDomainRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainRemarkRequest) SetLang(v string) *UpdateDomainRemarkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDomainRemarkRequest) SetUserClientIp(v string) *UpdateDomainRemarkRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateDomainRemarkRequest) SetDomainName(v string) *UpdateDomainRemarkRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDomainRemarkRequest) SetRemark(v string) *UpdateDomainRemarkRequest {
	s.Remark = &v
	return s
}

type UpdateDomainRemarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDomainRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainRemarkResponseBody) SetRequestId(v string) *UpdateDomainRemarkResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDomainRemarkResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDomainRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDomainRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainRemarkResponse) SetHeaders(v map[string]*string) *UpdateDomainRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainRemarkResponse) SetBody(v *UpdateDomainRemarkResponseBody) *UpdateDomainRemarkResponse {
	s.Body = v
	return s
}

type UpdateGtmAccessStrategyRequest struct {
	Lang               *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp       *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	StrategyId         *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	StrategyName       *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	DefaultAddrPoolId  *string `json:"DefaultAddrPoolId,omitempty" xml:"DefaultAddrPoolId,omitempty"`
	FailoverAddrPoolId *string `json:"FailoverAddrPoolId,omitempty" xml:"FailoverAddrPoolId,omitempty"`
	AccessLines        *string `json:"AccessLines,omitempty" xml:"AccessLines,omitempty"`
}

func (s UpdateGtmAccessStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateGtmAccessStrategyRequest) SetLang(v string) *UpdateGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *UpdateGtmAccessStrategyRequest) SetUserClientIp(v string) *UpdateGtmAccessStrategyRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateGtmAccessStrategyRequest) SetStrategyId(v string) *UpdateGtmAccessStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *UpdateGtmAccessStrategyRequest) SetStrategyName(v string) *UpdateGtmAccessStrategyRequest {
	s.StrategyName = &v
	return s
}

func (s *UpdateGtmAccessStrategyRequest) SetDefaultAddrPoolId(v string) *UpdateGtmAccessStrategyRequest {
	s.DefaultAddrPoolId = &v
	return s
}

func (s *UpdateGtmAccessStrategyRequest) SetFailoverAddrPoolId(v string) *UpdateGtmAccessStrategyRequest {
	s.FailoverAddrPoolId = &v
	return s
}

func (s *UpdateGtmAccessStrategyRequest) SetAccessLines(v string) *UpdateGtmAccessStrategyRequest {
	s.AccessLines = &v
	return s
}

type UpdateGtmAccessStrategyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGtmAccessStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGtmAccessStrategyResponseBody) SetRequestId(v string) *UpdateGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGtmAccessStrategyResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGtmAccessStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *UpdateGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateGtmAccessStrategyResponse) SetBody(v *UpdateGtmAccessStrategyResponseBody) *UpdateGtmAccessStrategyResponse {
	s.Body = v
	return s
}

type UpdateGtmAddressPoolRequest struct {
	UserClientIp        *string                            `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang                *string                            `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AddrPoolId          *string                            `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	Name                *string                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Type                *string                            `json:"Type,omitempty" xml:"Type,omitempty"`
	MinAvailableAddrNum *int32                             `json:"MinAvailableAddrNum,omitempty" xml:"MinAvailableAddrNum,omitempty"`
	Addr                []*UpdateGtmAddressPoolRequestAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Repeated"`
}

func (s UpdateGtmAddressPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *UpdateGtmAddressPoolRequest) SetUserClientIp(v string) *UpdateGtmAddressPoolRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateGtmAddressPoolRequest) SetLang(v string) *UpdateGtmAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *UpdateGtmAddressPoolRequest) SetAddrPoolId(v string) *UpdateGtmAddressPoolRequest {
	s.AddrPoolId = &v
	return s
}

func (s *UpdateGtmAddressPoolRequest) SetName(v string) *UpdateGtmAddressPoolRequest {
	s.Name = &v
	return s
}

func (s *UpdateGtmAddressPoolRequest) SetType(v string) *UpdateGtmAddressPoolRequest {
	s.Type = &v
	return s
}

func (s *UpdateGtmAddressPoolRequest) SetMinAvailableAddrNum(v int32) *UpdateGtmAddressPoolRequest {
	s.MinAvailableAddrNum = &v
	return s
}

func (s *UpdateGtmAddressPoolRequest) SetAddr(v []*UpdateGtmAddressPoolRequestAddr) *UpdateGtmAddressPoolRequest {
	s.Addr = v
	return s
}

type UpdateGtmAddressPoolRequestAddr struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	LbaWeight *int32  `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	Mode      *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s UpdateGtmAddressPoolRequestAddr) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmAddressPoolRequestAddr) GoString() string {
	return s.String()
}

func (s *UpdateGtmAddressPoolRequestAddr) SetValue(v string) *UpdateGtmAddressPoolRequestAddr {
	s.Value = &v
	return s
}

func (s *UpdateGtmAddressPoolRequestAddr) SetLbaWeight(v int32) *UpdateGtmAddressPoolRequestAddr {
	s.LbaWeight = &v
	return s
}

func (s *UpdateGtmAddressPoolRequestAddr) SetMode(v string) *UpdateGtmAddressPoolRequestAddr {
	s.Mode = &v
	return s
}

type UpdateGtmAddressPoolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGtmAddressPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGtmAddressPoolResponseBody) SetRequestId(v string) *UpdateGtmAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGtmAddressPoolResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGtmAddressPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *UpdateGtmAddressPoolResponse) SetHeaders(v map[string]*string) *UpdateGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *UpdateGtmAddressPoolResponse) SetBody(v *UpdateGtmAddressPoolResponseBody) *UpdateGtmAddressPoolResponse {
	s.Body = v
	return s
}

type UpdateGtmInstanceGlobalConfigRequest struct {
	Lang                  *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp          *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName          *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Ttl                   *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	UserDomainName        *string `json:"UserDomainName,omitempty" xml:"UserDomainName,omitempty"`
	LbaStrategy           *string `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	AlertGroup            *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	CnameMode             *string `json:"CnameMode,omitempty" xml:"CnameMode,omitempty"`
	CnameCustomDomainName *string `json:"CnameCustomDomainName,omitempty" xml:"CnameCustomDomainName,omitempty"`
}

func (s UpdateGtmInstanceGlobalConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmInstanceGlobalConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetLang(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.Lang = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetUserClientIp(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetInstanceId(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetInstanceName(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetTtl(v int32) *UpdateGtmInstanceGlobalConfigRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetUserDomainName(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.UserDomainName = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetLbaStrategy(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.LbaStrategy = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetAlertGroup(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.AlertGroup = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetCnameMode(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.CnameMode = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetCnameCustomDomainName(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.CnameCustomDomainName = &v
	return s
}

type UpdateGtmInstanceGlobalConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGtmInstanceGlobalConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmInstanceGlobalConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGtmInstanceGlobalConfigResponseBody) SetRequestId(v string) *UpdateGtmInstanceGlobalConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGtmInstanceGlobalConfigResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGtmInstanceGlobalConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGtmInstanceGlobalConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmInstanceGlobalConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateGtmInstanceGlobalConfigResponse) SetHeaders(v map[string]*string) *UpdateGtmInstanceGlobalConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigResponse) SetBody(v *UpdateGtmInstanceGlobalConfigResponseBody) *UpdateGtmInstanceGlobalConfigResponse {
	s.Body = v
	return s
}

type UpdateGtmMonitorRequest struct {
	UserClientIp      *string                               `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Lang              *string                               `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MonitorConfigId   *string                               `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	ProtocolType      *string                               `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Interval          *int32                                `json:"Interval,omitempty" xml:"Interval,omitempty"`
	EvaluationCount   *int32                                `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Timeout           *int32                                `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	MonitorExtendInfo *string                               `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
	IspCityNode       []*UpdateGtmMonitorRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
}

func (s UpdateGtmMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmMonitorRequest) GoString() string {
	return s.String()
}

func (s *UpdateGtmMonitorRequest) SetUserClientIp(v string) *UpdateGtmMonitorRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateGtmMonitorRequest) SetLang(v string) *UpdateGtmMonitorRequest {
	s.Lang = &v
	return s
}

func (s *UpdateGtmMonitorRequest) SetMonitorConfigId(v string) *UpdateGtmMonitorRequest {
	s.MonitorConfigId = &v
	return s
}

func (s *UpdateGtmMonitorRequest) SetProtocolType(v string) *UpdateGtmMonitorRequest {
	s.ProtocolType = &v
	return s
}

func (s *UpdateGtmMonitorRequest) SetInterval(v int32) *UpdateGtmMonitorRequest {
	s.Interval = &v
	return s
}

func (s *UpdateGtmMonitorRequest) SetEvaluationCount(v int32) *UpdateGtmMonitorRequest {
	s.EvaluationCount = &v
	return s
}

func (s *UpdateGtmMonitorRequest) SetTimeout(v int32) *UpdateGtmMonitorRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateGtmMonitorRequest) SetMonitorExtendInfo(v string) *UpdateGtmMonitorRequest {
	s.MonitorExtendInfo = &v
	return s
}

func (s *UpdateGtmMonitorRequest) SetIspCityNode(v []*UpdateGtmMonitorRequestIspCityNode) *UpdateGtmMonitorRequest {
	s.IspCityNode = v
	return s
}

type UpdateGtmMonitorRequestIspCityNode struct {
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	IspCode  *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
}

func (s UpdateGtmMonitorRequestIspCityNode) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmMonitorRequestIspCityNode) GoString() string {
	return s.String()
}

func (s *UpdateGtmMonitorRequestIspCityNode) SetCityCode(v string) *UpdateGtmMonitorRequestIspCityNode {
	s.CityCode = &v
	return s
}

func (s *UpdateGtmMonitorRequestIspCityNode) SetIspCode(v string) *UpdateGtmMonitorRequestIspCityNode {
	s.IspCode = &v
	return s
}

type UpdateGtmMonitorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGtmMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGtmMonitorResponseBody) SetRequestId(v string) *UpdateGtmMonitorResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGtmMonitorResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGtmMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGtmMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmMonitorResponse) GoString() string {
	return s.String()
}

func (s *UpdateGtmMonitorResponse) SetHeaders(v map[string]*string) *UpdateGtmMonitorResponse {
	s.Headers = v
	return s
}

func (s *UpdateGtmMonitorResponse) SetBody(v *UpdateGtmMonitorResponseBody) *UpdateGtmMonitorResponse {
	s.Body = v
	return s
}

type UpdateGtmRecoveryPlanRequest struct {
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp   *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	RecoveryPlanId *int64  `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark         *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	FaultAddrPool  *string `json:"FaultAddrPool,omitempty" xml:"FaultAddrPool,omitempty"`
}

func (s UpdateGtmRecoveryPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateGtmRecoveryPlanRequest) SetLang(v string) *UpdateGtmRecoveryPlanRequest {
	s.Lang = &v
	return s
}

func (s *UpdateGtmRecoveryPlanRequest) SetUserClientIp(v string) *UpdateGtmRecoveryPlanRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateGtmRecoveryPlanRequest) SetRecoveryPlanId(v int64) *UpdateGtmRecoveryPlanRequest {
	s.RecoveryPlanId = &v
	return s
}

func (s *UpdateGtmRecoveryPlanRequest) SetName(v string) *UpdateGtmRecoveryPlanRequest {
	s.Name = &v
	return s
}

func (s *UpdateGtmRecoveryPlanRequest) SetRemark(v string) *UpdateGtmRecoveryPlanRequest {
	s.Remark = &v
	return s
}

func (s *UpdateGtmRecoveryPlanRequest) SetFaultAddrPool(v string) *UpdateGtmRecoveryPlanRequest {
	s.FaultAddrPool = &v
	return s
}

type UpdateGtmRecoveryPlanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGtmRecoveryPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGtmRecoveryPlanResponseBody) SetRequestId(v string) *UpdateGtmRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGtmRecoveryPlanResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGtmRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGtmRecoveryPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGtmRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateGtmRecoveryPlanResponse) SetHeaders(v map[string]*string) *UpdateGtmRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateGtmRecoveryPlanResponse) SetBody(v *UpdateGtmRecoveryPlanResponseBody) *UpdateGtmRecoveryPlanResponse {
	s.Body = v
	return s
}

type ValidateDnsGtmAttributeInfoRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	LineCode     *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
}

func (s ValidateDnsGtmAttributeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateDnsGtmAttributeInfoRequest) GoString() string {
	return s.String()
}

func (s *ValidateDnsGtmAttributeInfoRequest) SetLang(v string) *ValidateDnsGtmAttributeInfoRequest {
	s.Lang = &v
	return s
}

func (s *ValidateDnsGtmAttributeInfoRequest) SetUserClientIp(v string) *ValidateDnsGtmAttributeInfoRequest {
	s.UserClientIp = &v
	return s
}

func (s *ValidateDnsGtmAttributeInfoRequest) SetLineCode(v string) *ValidateDnsGtmAttributeInfoRequest {
	s.LineCode = &v
	return s
}

type ValidateDnsGtmAttributeInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValidateDnsGtmAttributeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateDnsGtmAttributeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateDnsGtmAttributeInfoResponseBody) SetRequestId(v string) *ValidateDnsGtmAttributeInfoResponseBody {
	s.RequestId = &v
	return s
}

type ValidateDnsGtmAttributeInfoResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateDnsGtmAttributeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateDnsGtmAttributeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateDnsGtmAttributeInfoResponse) GoString() string {
	return s.String()
}

func (s *ValidateDnsGtmAttributeInfoResponse) SetHeaders(v map[string]*string) *ValidateDnsGtmAttributeInfoResponse {
	s.Headers = v
	return s
}

func (s *ValidateDnsGtmAttributeInfoResponse) SetBody(v *ValidateDnsGtmAttributeInfoResponseBody) *ValidateDnsGtmAttributeInfoResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("alidns"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddCustomLineWithOptions(request *AddCustomLineRequest, runtime *util.RuntimeOptions) (_result *AddCustomLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddCustomLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddCustomLine"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCustomLine(request *AddCustomLineRequest) (_result *AddCustomLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCustomLineResponse{}
	_body, _err := client.AddCustomLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDnsGtmAccessStrategyWithOptions(request *AddDnsGtmAccessStrategyRequest, runtime *util.RuntimeOptions) (_result *AddDnsGtmAccessStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddDnsGtmAccessStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddDnsGtmAccessStrategy"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDnsGtmAccessStrategy(request *AddDnsGtmAccessStrategyRequest) (_result *AddDnsGtmAccessStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDnsGtmAccessStrategyResponse{}
	_body, _err := client.AddDnsGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDnsGtmAddressPoolWithOptions(request *AddDnsGtmAddressPoolRequest, runtime *util.RuntimeOptions) (_result *AddDnsGtmAddressPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddDnsGtmAddressPoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddDnsGtmAddressPool"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDnsGtmAddressPool(request *AddDnsGtmAddressPoolRequest) (_result *AddDnsGtmAddressPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDnsGtmAddressPoolResponse{}
	_body, _err := client.AddDnsGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDnsGtmMonitorWithOptions(request *AddDnsGtmMonitorRequest, runtime *util.RuntimeOptions) (_result *AddDnsGtmMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddDnsGtmMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddDnsGtmMonitor"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDnsGtmMonitor(request *AddDnsGtmMonitorRequest) (_result *AddDnsGtmMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDnsGtmMonitorResponse{}
	_body, _err := client.AddDnsGtmMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDomainWithOptions(request *AddDomainRequest, runtime *util.RuntimeOptions) (_result *AddDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddDomain"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDomain(request *AddDomainRequest) (_result *AddDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDomainResponse{}
	_body, _err := client.AddDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDomainBackupWithOptions(request *AddDomainBackupRequest, runtime *util.RuntimeOptions) (_result *AddDomainBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddDomainBackupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddDomainBackup"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDomainBackup(request *AddDomainBackupRequest) (_result *AddDomainBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDomainBackupResponse{}
	_body, _err := client.AddDomainBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDomainGroupWithOptions(request *AddDomainGroupRequest, runtime *util.RuntimeOptions) (_result *AddDomainGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddDomainGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddDomainGroup"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDomainGroup(request *AddDomainGroupRequest) (_result *AddDomainGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDomainGroupResponse{}
	_body, _err := client.AddDomainGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDomainRecordWithOptions(request *AddDomainRecordRequest, runtime *util.RuntimeOptions) (_result *AddDomainRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddDomainRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddDomainRecord"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDomainRecord(request *AddDomainRecordRequest) (_result *AddDomainRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDomainRecordResponse{}
	_body, _err := client.AddDomainRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddGtmAccessStrategyWithOptions(request *AddGtmAccessStrategyRequest, runtime *util.RuntimeOptions) (_result *AddGtmAccessStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddGtmAccessStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddGtmAccessStrategy"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddGtmAccessStrategy(request *AddGtmAccessStrategyRequest) (_result *AddGtmAccessStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddGtmAccessStrategyResponse{}
	_body, _err := client.AddGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddGtmAddressPoolWithOptions(request *AddGtmAddressPoolRequest, runtime *util.RuntimeOptions) (_result *AddGtmAddressPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddGtmAddressPoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddGtmAddressPool"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddGtmAddressPool(request *AddGtmAddressPoolRequest) (_result *AddGtmAddressPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddGtmAddressPoolResponse{}
	_body, _err := client.AddGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddGtmMonitorWithOptions(request *AddGtmMonitorRequest, runtime *util.RuntimeOptions) (_result *AddGtmMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddGtmMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddGtmMonitor"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddGtmMonitor(request *AddGtmMonitorRequest) (_result *AddGtmMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddGtmMonitorResponse{}
	_body, _err := client.AddGtmMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddGtmRecoveryPlanWithOptions(request *AddGtmRecoveryPlanRequest, runtime *util.RuntimeOptions) (_result *AddGtmRecoveryPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddGtmRecoveryPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddGtmRecoveryPlan"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddGtmRecoveryPlan(request *AddGtmRecoveryPlanRequest) (_result *AddGtmRecoveryPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddGtmRecoveryPlanResponse{}
	_body, _err := client.AddGtmRecoveryPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindInstanceDomainsWithOptions(request *BindInstanceDomainsRequest, runtime *util.RuntimeOptions) (_result *BindInstanceDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindInstanceDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindInstanceDomains"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindInstanceDomains(request *BindInstanceDomainsRequest) (_result *BindInstanceDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindInstanceDomainsResponse{}
	_body, _err := client.BindInstanceDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeDomainGroupWithOptions(request *ChangeDomainGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeDomainGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChangeDomainGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ChangeDomainGroup"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeDomainGroup(request *ChangeDomainGroupRequest) (_result *ChangeDomainGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeDomainGroupResponse{}
	_body, _err := client.ChangeDomainGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeDomainOfDnsProductWithOptions(request *ChangeDomainOfDnsProductRequest, runtime *util.RuntimeOptions) (_result *ChangeDomainOfDnsProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChangeDomainOfDnsProductResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ChangeDomainOfDnsProduct"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeDomainOfDnsProduct(request *ChangeDomainOfDnsProductRequest) (_result *ChangeDomainOfDnsProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeDomainOfDnsProductResponse{}
	_body, _err := client.ChangeDomainOfDnsProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopyGtmConfigWithOptions(request *CopyGtmConfigRequest, runtime *util.RuntimeOptions) (_result *CopyGtmConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CopyGtmConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CopyGtmConfig"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopyGtmConfig(request *CopyGtmConfigRequest) (_result *CopyGtmConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopyGtmConfigResponse{}
	_body, _err := client.CopyGtmConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCustomLinesWithOptions(request *DeleteCustomLinesRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomLinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCustomLinesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCustomLines"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCustomLines(request *DeleteCustomLinesRequest) (_result *DeleteCustomLinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomLinesResponse{}
	_body, _err := client.DeleteCustomLinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDnsGtmAccessStrategyWithOptions(request *DeleteDnsGtmAccessStrategyRequest, runtime *util.RuntimeOptions) (_result *DeleteDnsGtmAccessStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDnsGtmAccessStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDnsGtmAccessStrategy"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDnsGtmAccessStrategy(request *DeleteDnsGtmAccessStrategyRequest) (_result *DeleteDnsGtmAccessStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDnsGtmAccessStrategyResponse{}
	_body, _err := client.DeleteDnsGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDnsGtmAddressPoolWithOptions(request *DeleteDnsGtmAddressPoolRequest, runtime *util.RuntimeOptions) (_result *DeleteDnsGtmAddressPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDnsGtmAddressPoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDnsGtmAddressPool"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDnsGtmAddressPool(request *DeleteDnsGtmAddressPoolRequest) (_result *DeleteDnsGtmAddressPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDnsGtmAddressPoolResponse{}
	_body, _err := client.DeleteDnsGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDomain"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomain(request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainGroupWithOptions(request *DeleteDomainGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDomainGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDomainGroup"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomainGroup(request *DeleteDomainGroupRequest) (_result *DeleteDomainGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainGroupResponse{}
	_body, _err := client.DeleteDomainGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainRecordWithOptions(request *DeleteDomainRecordRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDomainRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDomainRecord"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomainRecord(request *DeleteDomainRecordRequest) (_result *DeleteDomainRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainRecordResponse{}
	_body, _err := client.DeleteDomainRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGtmAccessStrategyWithOptions(request *DeleteGtmAccessStrategyRequest, runtime *util.RuntimeOptions) (_result *DeleteGtmAccessStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteGtmAccessStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteGtmAccessStrategy"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGtmAccessStrategy(request *DeleteGtmAccessStrategyRequest) (_result *DeleteGtmAccessStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGtmAccessStrategyResponse{}
	_body, _err := client.DeleteGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGtmAddressPoolWithOptions(request *DeleteGtmAddressPoolRequest, runtime *util.RuntimeOptions) (_result *DeleteGtmAddressPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteGtmAddressPoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteGtmAddressPool"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGtmAddressPool(request *DeleteGtmAddressPoolRequest) (_result *DeleteGtmAddressPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGtmAddressPoolResponse{}
	_body, _err := client.DeleteGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGtmRecoveryPlanWithOptions(request *DeleteGtmRecoveryPlanRequest, runtime *util.RuntimeOptions) (_result *DeleteGtmRecoveryPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteGtmRecoveryPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteGtmRecoveryPlan"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGtmRecoveryPlan(request *DeleteGtmRecoveryPlanRequest) (_result *DeleteGtmRecoveryPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGtmRecoveryPlanResponse{}
	_body, _err := client.DeleteGtmRecoveryPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSubDomainRecordsWithOptions(request *DeleteSubDomainRecordsRequest, runtime *util.RuntimeOptions) (_result *DeleteSubDomainRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSubDomainRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSubDomainRecords"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSubDomainRecords(request *DeleteSubDomainRecordsRequest) (_result *DeleteSubDomainRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSubDomainRecordsResponse{}
	_body, _err := client.DeleteSubDomainRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBatchResultCountWithOptions(request *DescribeBatchResultCountRequest, runtime *util.RuntimeOptions) (_result *DescribeBatchResultCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBatchResultCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBatchResultCount"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBatchResultCount(request *DescribeBatchResultCountRequest) (_result *DescribeBatchResultCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBatchResultCountResponse{}
	_body, _err := client.DescribeBatchResultCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBatchResultDetailWithOptions(request *DescribeBatchResultDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeBatchResultDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBatchResultDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBatchResultDetail"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBatchResultDetail(request *DescribeBatchResultDetailRequest) (_result *DescribeBatchResultDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBatchResultDetailResponse{}
	_body, _err := client.DescribeBatchResultDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomLineWithOptions(request *DescribeCustomLineRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCustomLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCustomLine"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomLine(request *DescribeCustomLineRequest) (_result *DescribeCustomLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomLineResponse{}
	_body, _err := client.DescribeCustomLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomLinesWithOptions(request *DescribeCustomLinesRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomLinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCustomLinesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCustomLines"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomLines(request *DescribeCustomLinesRequest) (_result *DescribeCustomLinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomLinesResponse{}
	_body, _err := client.DescribeCustomLinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmAccessStrategiesWithOptions(request *DescribeDnsGtmAccessStrategiesRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmAccessStrategiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmAccessStrategiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmAccessStrategies"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmAccessStrategies(request *DescribeDnsGtmAccessStrategiesRequest) (_result *DescribeDnsGtmAccessStrategiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmAccessStrategiesResponse{}
	_body, _err := client.DescribeDnsGtmAccessStrategiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmAccessStrategyWithOptions(request *DescribeDnsGtmAccessStrategyRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmAccessStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmAccessStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmAccessStrategy"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmAccessStrategy(request *DescribeDnsGtmAccessStrategyRequest) (_result *DescribeDnsGtmAccessStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmAccessStrategyResponse{}
	_body, _err := client.DescribeDnsGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmAccessStrategyAvailableConfigWithOptions(request *DescribeDnsGtmAccessStrategyAvailableConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmAccessStrategyAvailableConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmAccessStrategyAvailableConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmAccessStrategyAvailableConfig"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmAccessStrategyAvailableConfig(request *DescribeDnsGtmAccessStrategyAvailableConfigRequest) (_result *DescribeDnsGtmAccessStrategyAvailableConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmAccessStrategyAvailableConfigResponse{}
	_body, _err := client.DescribeDnsGtmAccessStrategyAvailableConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmAddrAttributeInfoWithOptions(request *DescribeDnsGtmAddrAttributeInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmAddrAttributeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmAddrAttributeInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmAddrAttributeInfo"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmAddrAttributeInfo(request *DescribeDnsGtmAddrAttributeInfoRequest) (_result *DescribeDnsGtmAddrAttributeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmAddrAttributeInfoResponse{}
	_body, _err := client.DescribeDnsGtmAddrAttributeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmAddressPoolAvailableConfigWithOptions(request *DescribeDnsGtmAddressPoolAvailableConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmAddressPoolAvailableConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmAddressPoolAvailableConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmAddressPoolAvailableConfig"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmAddressPoolAvailableConfig(request *DescribeDnsGtmAddressPoolAvailableConfigRequest) (_result *DescribeDnsGtmAddressPoolAvailableConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmAddressPoolAvailableConfigResponse{}
	_body, _err := client.DescribeDnsGtmAddressPoolAvailableConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmAvailableAlertGroupWithOptions(request *DescribeDnsGtmAvailableAlertGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmAvailableAlertGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmAvailableAlertGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmAvailableAlertGroup"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmAvailableAlertGroup(request *DescribeDnsGtmAvailableAlertGroupRequest) (_result *DescribeDnsGtmAvailableAlertGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmAvailableAlertGroupResponse{}
	_body, _err := client.DescribeDnsGtmAvailableAlertGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmInstanceWithOptions(request *DescribeDnsGtmInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmInstance"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmInstance(request *DescribeDnsGtmInstanceRequest) (_result *DescribeDnsGtmInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceResponse{}
	_body, _err := client.DescribeDnsGtmInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmInstanceAddressPoolWithOptions(request *DescribeDnsGtmInstanceAddressPoolRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmInstanceAddressPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmInstanceAddressPoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmInstanceAddressPool"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmInstanceAddressPool(request *DescribeDnsGtmInstanceAddressPoolRequest) (_result *DescribeDnsGtmInstanceAddressPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceAddressPoolResponse{}
	_body, _err := client.DescribeDnsGtmInstanceAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmInstanceAddressPoolsWithOptions(request *DescribeDnsGtmInstanceAddressPoolsRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmInstanceAddressPoolsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmInstanceAddressPoolsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmInstanceAddressPools"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmInstanceAddressPools(request *DescribeDnsGtmInstanceAddressPoolsRequest) (_result *DescribeDnsGtmInstanceAddressPoolsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceAddressPoolsResponse{}
	_body, _err := client.DescribeDnsGtmInstanceAddressPoolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmInstancesWithOptions(request *DescribeDnsGtmInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmInstances"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmInstances(request *DescribeDnsGtmInstancesRequest) (_result *DescribeDnsGtmInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmInstancesResponse{}
	_body, _err := client.DescribeDnsGtmInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmInstanceStatusWithOptions(request *DescribeDnsGtmInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmInstanceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmInstanceStatus"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmInstanceStatus(request *DescribeDnsGtmInstanceStatusRequest) (_result *DescribeDnsGtmInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceStatusResponse{}
	_body, _err := client.DescribeDnsGtmInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmInstanceSystemCnameWithOptions(request *DescribeDnsGtmInstanceSystemCnameRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmInstanceSystemCnameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmInstanceSystemCnameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmInstanceSystemCname"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmInstanceSystemCname(request *DescribeDnsGtmInstanceSystemCnameRequest) (_result *DescribeDnsGtmInstanceSystemCnameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceSystemCnameResponse{}
	_body, _err := client.DescribeDnsGtmInstanceSystemCnameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmLogsWithOptions(request *DescribeDnsGtmLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmLogs"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmLogs(request *DescribeDnsGtmLogsRequest) (_result *DescribeDnsGtmLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmLogsResponse{}
	_body, _err := client.DescribeDnsGtmLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmMonitorAvailableConfigWithOptions(request *DescribeDnsGtmMonitorAvailableConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmMonitorAvailableConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmMonitorAvailableConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmMonitorAvailableConfig"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmMonitorAvailableConfig(request *DescribeDnsGtmMonitorAvailableConfigRequest) (_result *DescribeDnsGtmMonitorAvailableConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmMonitorAvailableConfigResponse{}
	_body, _err := client.DescribeDnsGtmMonitorAvailableConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsGtmMonitorConfigWithOptions(request *DescribeDnsGtmMonitorConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsGtmMonitorConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsGtmMonitorConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsGtmMonitorConfig"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsGtmMonitorConfig(request *DescribeDnsGtmMonitorConfigRequest) (_result *DescribeDnsGtmMonitorConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsGtmMonitorConfigResponse{}
	_body, _err := client.DescribeDnsGtmMonitorConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsProductInstanceWithOptions(request *DescribeDnsProductInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsProductInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsProductInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsProductInstance"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsProductInstance(request *DescribeDnsProductInstanceRequest) (_result *DescribeDnsProductInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsProductInstanceResponse{}
	_body, _err := client.DescribeDnsProductInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnsProductInstancesWithOptions(request *DescribeDnsProductInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDnsProductInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnsProductInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnsProductInstances"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnsProductInstances(request *DescribeDnsProductInstancesRequest) (_result *DescribeDnsProductInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnsProductInstancesResponse{}
	_body, _err := client.DescribeDnsProductInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDNSSLBSubDomainsWithOptions(request *DescribeDNSSLBSubDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeDNSSLBSubDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDNSSLBSubDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDNSSLBSubDomains"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDNSSLBSubDomains(request *DescribeDNSSLBSubDomainsRequest) (_result *DescribeDNSSLBSubDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDNSSLBSubDomainsResponse{}
	_body, _err := client.DescribeDNSSLBSubDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDohAccountStatisticsWithOptions(request *DescribeDohAccountStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeDohAccountStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDohAccountStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDohAccountStatistics"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDohAccountStatistics(request *DescribeDohAccountStatisticsRequest) (_result *DescribeDohAccountStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDohAccountStatisticsResponse{}
	_body, _err := client.DescribeDohAccountStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDohDomainStatisticsWithOptions(request *DescribeDohDomainStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeDohDomainStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDohDomainStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDohDomainStatistics"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDohDomainStatistics(request *DescribeDohDomainStatisticsRequest) (_result *DescribeDohDomainStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDohDomainStatisticsResponse{}
	_body, _err := client.DescribeDohDomainStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDohDomainStatisticsSummaryWithOptions(request *DescribeDohDomainStatisticsSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeDohDomainStatisticsSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDohDomainStatisticsSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDohDomainStatisticsSummary"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDohDomainStatisticsSummary(request *DescribeDohDomainStatisticsSummaryRequest) (_result *DescribeDohDomainStatisticsSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDohDomainStatisticsSummaryResponse{}
	_body, _err := client.DescribeDohDomainStatisticsSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDohSubDomainStatisticsWithOptions(request *DescribeDohSubDomainStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeDohSubDomainStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDohSubDomainStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDohSubDomainStatistics"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDohSubDomainStatistics(request *DescribeDohSubDomainStatisticsRequest) (_result *DescribeDohSubDomainStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDohSubDomainStatisticsResponse{}
	_body, _err := client.DescribeDohSubDomainStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDohSubDomainStatisticsSummaryWithOptions(request *DescribeDohSubDomainStatisticsSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeDohSubDomainStatisticsSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDohSubDomainStatisticsSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDohSubDomainStatisticsSummary"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDohSubDomainStatisticsSummary(request *DescribeDohSubDomainStatisticsSummaryRequest) (_result *DescribeDohSubDomainStatisticsSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDohSubDomainStatisticsSummaryResponse{}
	_body, _err := client.DescribeDohSubDomainStatisticsSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDohUserInfoWithOptions(request *DescribeDohUserInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDohUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDohUserInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDohUserInfo"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDohUserInfo(request *DescribeDohUserInfoRequest) (_result *DescribeDohUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDohUserInfoResponse{}
	_body, _err := client.DescribeDohUserInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainDnssecInfoWithOptions(request *DescribeDomainDnssecInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainDnssecInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainDnssecInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainDnssecInfo"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainDnssecInfo(request *DescribeDomainDnssecInfoRequest) (_result *DescribeDomainDnssecInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainDnssecInfoResponse{}
	_body, _err := client.DescribeDomainDnssecInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainGroupsWithOptions(request *DescribeDomainGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainGroups"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainGroups(request *DescribeDomainGroupsRequest) (_result *DescribeDomainGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainGroupsResponse{}
	_body, _err := client.DescribeDomainGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainInfoWithOptions(request *DescribeDomainInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainInfo"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainInfo(request *DescribeDomainInfoRequest) (_result *DescribeDomainInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainInfoResponse{}
	_body, _err := client.DescribeDomainInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainLogsWithOptions(request *DescribeDomainLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainLogs"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainLogs(request *DescribeDomainLogsRequest) (_result *DescribeDomainLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainLogsResponse{}
	_body, _err := client.DescribeDomainLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainNsWithOptions(request *DescribeDomainNsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainNsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainNsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainNs"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainNs(request *DescribeDomainNsRequest) (_result *DescribeDomainNsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainNsResponse{}
	_body, _err := client.DescribeDomainNsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainRecordInfoWithOptions(request *DescribeDomainRecordInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRecordInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainRecordInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainRecordInfo"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRecordInfo(request *DescribeDomainRecordInfoRequest) (_result *DescribeDomainRecordInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRecordInfoResponse{}
	_body, _err := client.DescribeDomainRecordInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainRecordsWithOptions(request *DescribeDomainRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainRecords"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRecords(request *DescribeDomainRecordsRequest) (_result *DescribeDomainRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRecordsResponse{}
	_body, _err := client.DescribeDomainRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainsWithOptions(request *DescribeDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomains"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomains(request *DescribeDomainsRequest) (_result *DescribeDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.DescribeDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainStatisticsWithOptions(request *DescribeDomainStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainStatistics"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainStatistics(request *DescribeDomainStatisticsRequest) (_result *DescribeDomainStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainStatisticsResponse{}
	_body, _err := client.DescribeDomainStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainStatisticsSummaryWithOptions(request *DescribeDomainStatisticsSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainStatisticsSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainStatisticsSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainStatisticsSummary"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainStatisticsSummary(request *DescribeDomainStatisticsSummaryRequest) (_result *DescribeDomainStatisticsSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainStatisticsSummaryResponse{}
	_body, _err := client.DescribeDomainStatisticsSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmAccessStrategiesWithOptions(request *DescribeGtmAccessStrategiesRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmAccessStrategiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmAccessStrategiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmAccessStrategies"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmAccessStrategies(request *DescribeGtmAccessStrategiesRequest) (_result *DescribeGtmAccessStrategiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmAccessStrategiesResponse{}
	_body, _err := client.DescribeGtmAccessStrategiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmAccessStrategyWithOptions(request *DescribeGtmAccessStrategyRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmAccessStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmAccessStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmAccessStrategy"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmAccessStrategy(request *DescribeGtmAccessStrategyRequest) (_result *DescribeGtmAccessStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmAccessStrategyResponse{}
	_body, _err := client.DescribeGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmAccessStrategyAvailableConfigWithOptions(request *DescribeGtmAccessStrategyAvailableConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmAccessStrategyAvailableConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmAccessStrategyAvailableConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmAccessStrategyAvailableConfig"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmAccessStrategyAvailableConfig(request *DescribeGtmAccessStrategyAvailableConfigRequest) (_result *DescribeGtmAccessStrategyAvailableConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmAccessStrategyAvailableConfigResponse{}
	_body, _err := client.DescribeGtmAccessStrategyAvailableConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmAvailableAlertGroupWithOptions(request *DescribeGtmAvailableAlertGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmAvailableAlertGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmAvailableAlertGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmAvailableAlertGroup"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmAvailableAlertGroup(request *DescribeGtmAvailableAlertGroupRequest) (_result *DescribeGtmAvailableAlertGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmAvailableAlertGroupResponse{}
	_body, _err := client.DescribeGtmAvailableAlertGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmInstanceWithOptions(request *DescribeGtmInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmInstance"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmInstance(request *DescribeGtmInstanceRequest) (_result *DescribeGtmInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmInstanceResponse{}
	_body, _err := client.DescribeGtmInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmInstanceAddressPoolWithOptions(request *DescribeGtmInstanceAddressPoolRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmInstanceAddressPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmInstanceAddressPoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmInstanceAddressPool"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmInstanceAddressPool(request *DescribeGtmInstanceAddressPoolRequest) (_result *DescribeGtmInstanceAddressPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmInstanceAddressPoolResponse{}
	_body, _err := client.DescribeGtmInstanceAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmInstanceAddressPoolsWithOptions(request *DescribeGtmInstanceAddressPoolsRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmInstanceAddressPoolsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmInstanceAddressPoolsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmInstanceAddressPools"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmInstanceAddressPools(request *DescribeGtmInstanceAddressPoolsRequest) (_result *DescribeGtmInstanceAddressPoolsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmInstanceAddressPoolsResponse{}
	_body, _err := client.DescribeGtmInstanceAddressPoolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmInstancesWithOptions(request *DescribeGtmInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmInstances"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmInstances(request *DescribeGtmInstancesRequest) (_result *DescribeGtmInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmInstancesResponse{}
	_body, _err := client.DescribeGtmInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmInstanceStatusWithOptions(request *DescribeGtmInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmInstanceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmInstanceStatus"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmInstanceStatus(request *DescribeGtmInstanceStatusRequest) (_result *DescribeGtmInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmInstanceStatusResponse{}
	_body, _err := client.DescribeGtmInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmInstanceSystemCnameWithOptions(request *DescribeGtmInstanceSystemCnameRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmInstanceSystemCnameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmInstanceSystemCnameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmInstanceSystemCname"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmInstanceSystemCname(request *DescribeGtmInstanceSystemCnameRequest) (_result *DescribeGtmInstanceSystemCnameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmInstanceSystemCnameResponse{}
	_body, _err := client.DescribeGtmInstanceSystemCnameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmLogsWithOptions(request *DescribeGtmLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmLogs"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmLogs(request *DescribeGtmLogsRequest) (_result *DescribeGtmLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmLogsResponse{}
	_body, _err := client.DescribeGtmLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmMonitorAvailableConfigWithOptions(request *DescribeGtmMonitorAvailableConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmMonitorAvailableConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmMonitorAvailableConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmMonitorAvailableConfig"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmMonitorAvailableConfig(request *DescribeGtmMonitorAvailableConfigRequest) (_result *DescribeGtmMonitorAvailableConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmMonitorAvailableConfigResponse{}
	_body, _err := client.DescribeGtmMonitorAvailableConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmMonitorConfigWithOptions(request *DescribeGtmMonitorConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmMonitorConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmMonitorConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmMonitorConfig"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmMonitorConfig(request *DescribeGtmMonitorConfigRequest) (_result *DescribeGtmMonitorConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmMonitorConfigResponse{}
	_body, _err := client.DescribeGtmMonitorConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmRecoveryPlanWithOptions(request *DescribeGtmRecoveryPlanRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmRecoveryPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmRecoveryPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmRecoveryPlan"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmRecoveryPlan(request *DescribeGtmRecoveryPlanRequest) (_result *DescribeGtmRecoveryPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmRecoveryPlanResponse{}
	_body, _err := client.DescribeGtmRecoveryPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmRecoveryPlanAvailableConfigWithOptions(request *DescribeGtmRecoveryPlanAvailableConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmRecoveryPlanAvailableConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmRecoveryPlanAvailableConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmRecoveryPlanAvailableConfig"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmRecoveryPlanAvailableConfig(request *DescribeGtmRecoveryPlanAvailableConfigRequest) (_result *DescribeGtmRecoveryPlanAvailableConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmRecoveryPlanAvailableConfigResponse{}
	_body, _err := client.DescribeGtmRecoveryPlanAvailableConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGtmRecoveryPlansWithOptions(request *DescribeGtmRecoveryPlansRequest, runtime *util.RuntimeOptions) (_result *DescribeGtmRecoveryPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGtmRecoveryPlansResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGtmRecoveryPlans"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGtmRecoveryPlans(request *DescribeGtmRecoveryPlansRequest) (_result *DescribeGtmRecoveryPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGtmRecoveryPlansResponse{}
	_body, _err := client.DescribeGtmRecoveryPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceDomainsWithOptions(request *DescribeInstanceDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceDomains"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceDomains(request *DescribeInstanceDomainsRequest) (_result *DescribeInstanceDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceDomainsResponse{}
	_body, _err := client.DescribeInstanceDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecordLogsWithOptions(request *DescribeRecordLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRecordLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRecordLogs"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecordLogs(request *DescribeRecordLogsRequest) (_result *DescribeRecordLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecordLogsResponse{}
	_body, _err := client.DescribeRecordLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecordStatisticsWithOptions(request *DescribeRecordStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRecordStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRecordStatistics"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecordStatistics(request *DescribeRecordStatisticsRequest) (_result *DescribeRecordStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecordStatisticsResponse{}
	_body, _err := client.DescribeRecordStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecordStatisticsSummaryWithOptions(request *DescribeRecordStatisticsSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordStatisticsSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRecordStatisticsSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRecordStatisticsSummary"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecordStatisticsSummary(request *DescribeRecordStatisticsSummaryRequest) (_result *DescribeRecordStatisticsSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecordStatisticsSummaryResponse{}
	_body, _err := client.DescribeRecordStatisticsSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSubDomainRecordsWithOptions(request *DescribeSubDomainRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSubDomainRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSubDomainRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSubDomainRecords"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubDomainRecords(request *DescribeSubDomainRecordsRequest) (_result *DescribeSubDomainRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubDomainRecordsResponse{}
	_body, _err := client.DescribeSubDomainRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSupportLinesWithOptions(request *DescribeSupportLinesRequest, runtime *util.RuntimeOptions) (_result *DescribeSupportLinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSupportLinesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSupportLines"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSupportLines(request *DescribeSupportLinesRequest) (_result *DescribeSupportLinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSupportLinesResponse{}
	_body, _err := client.DescribeSupportLinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTagsWithOptions(request *DescribeTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTags"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTags(request *DescribeTagsRequest) (_result *DescribeTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DescribeTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTransferDomainsWithOptions(request *DescribeTransferDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeTransferDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTransferDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTransferDomains"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTransferDomains(request *DescribeTransferDomainsRequest) (_result *DescribeTransferDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTransferDomainsResponse{}
	_body, _err := client.DescribeTransferDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteGtmRecoveryPlanWithOptions(request *ExecuteGtmRecoveryPlanRequest, runtime *util.RuntimeOptions) (_result *ExecuteGtmRecoveryPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteGtmRecoveryPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteGtmRecoveryPlan"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteGtmRecoveryPlan(request *ExecuteGtmRecoveryPlanRequest) (_result *ExecuteGtmRecoveryPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteGtmRecoveryPlanResponse{}
	_body, _err := client.ExecuteGtmRecoveryPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMainDomainNameWithOptions(request *GetMainDomainNameRequest, runtime *util.RuntimeOptions) (_result *GetMainDomainNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMainDomainNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMainDomainName"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMainDomainName(request *GetMainDomainNameRequest) (_result *GetMainDomainNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMainDomainNameResponse{}
	_body, _err := client.GetMainDomainNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTxtRecordForVerifyWithOptions(request *GetTxtRecordForVerifyRequest, runtime *util.RuntimeOptions) (_result *GetTxtRecordForVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTxtRecordForVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTxtRecordForVerify"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTxtRecordForVerify(request *GetTxtRecordForVerifyRequest) (_result *GetTxtRecordForVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTxtRecordForVerifyResponse{}
	_body, _err := client.GetTxtRecordForVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHichinaDomainDNSWithOptions(request *ModifyHichinaDomainDNSRequest, runtime *util.RuntimeOptions) (_result *ModifyHichinaDomainDNSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyHichinaDomainDNSResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyHichinaDomainDNS"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHichinaDomainDNS(request *ModifyHichinaDomainDNSRequest) (_result *ModifyHichinaDomainDNSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHichinaDomainDNSResponse{}
	_body, _err := client.ModifyHichinaDomainDNSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveDomainResourceGroupWithOptions(request *MoveDomainResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveDomainResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MoveDomainResourceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MoveDomainResourceGroup"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveDomainResourceGroup(request *MoveDomainResourceGroupRequest) (_result *MoveDomainResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveDomainResourceGroupResponse{}
	_body, _err := client.MoveDomainResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveGtmResourceGroupWithOptions(request *MoveGtmResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveGtmResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MoveGtmResourceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MoveGtmResourceGroup"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveGtmResourceGroup(request *MoveGtmResourceGroupRequest) (_result *MoveGtmResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveGtmResourceGroupResponse{}
	_body, _err := client.MoveGtmResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OperateBatchDomainWithOptions(request *OperateBatchDomainRequest, runtime *util.RuntimeOptions) (_result *OperateBatchDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OperateBatchDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OperateBatchDomain"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OperateBatchDomain(request *OperateBatchDomainRequest) (_result *OperateBatchDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateBatchDomainResponse{}
	_body, _err := client.OperateBatchDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreviewGtmRecoveryPlanWithOptions(request *PreviewGtmRecoveryPlanRequest, runtime *util.RuntimeOptions) (_result *PreviewGtmRecoveryPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PreviewGtmRecoveryPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PreviewGtmRecoveryPlan"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreviewGtmRecoveryPlan(request *PreviewGtmRecoveryPlanRequest) (_result *PreviewGtmRecoveryPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PreviewGtmRecoveryPlanResponse{}
	_body, _err := client.PreviewGtmRecoveryPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RetrieveDomainWithOptions(request *RetrieveDomainRequest, runtime *util.RuntimeOptions) (_result *RetrieveDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RetrieveDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RetrieveDomain"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetrieveDomain(request *RetrieveDomainRequest) (_result *RetrieveDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetrieveDomainResponse{}
	_body, _err := client.RetrieveDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackGtmRecoveryPlanWithOptions(request *RollbackGtmRecoveryPlanRequest, runtime *util.RuntimeOptions) (_result *RollbackGtmRecoveryPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RollbackGtmRecoveryPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RollbackGtmRecoveryPlan"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackGtmRecoveryPlan(request *RollbackGtmRecoveryPlanRequest) (_result *RollbackGtmRecoveryPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackGtmRecoveryPlanResponse{}
	_body, _err := client.RollbackGtmRecoveryPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDnsGtmAccessModeWithOptions(request *SetDnsGtmAccessModeRequest, runtime *util.RuntimeOptions) (_result *SetDnsGtmAccessModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDnsGtmAccessModeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDnsGtmAccessMode"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDnsGtmAccessMode(request *SetDnsGtmAccessModeRequest) (_result *SetDnsGtmAccessModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDnsGtmAccessModeResponse{}
	_body, _err := client.SetDnsGtmAccessModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDnsGtmMonitorStatusWithOptions(request *SetDnsGtmMonitorStatusRequest, runtime *util.RuntimeOptions) (_result *SetDnsGtmMonitorStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDnsGtmMonitorStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDnsGtmMonitorStatus"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDnsGtmMonitorStatus(request *SetDnsGtmMonitorStatusRequest) (_result *SetDnsGtmMonitorStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDnsGtmMonitorStatusResponse{}
	_body, _err := client.SetDnsGtmMonitorStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDNSSLBStatusWithOptions(request *SetDNSSLBStatusRequest, runtime *util.RuntimeOptions) (_result *SetDNSSLBStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDNSSLBStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDNSSLBStatus"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDNSSLBStatus(request *SetDNSSLBStatusRequest) (_result *SetDNSSLBStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDNSSLBStatusResponse{}
	_body, _err := client.SetDNSSLBStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDomainDnssecStatusWithOptions(request *SetDomainDnssecStatusRequest, runtime *util.RuntimeOptions) (_result *SetDomainDnssecStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDomainDnssecStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDomainDnssecStatus"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDomainDnssecStatus(request *SetDomainDnssecStatusRequest) (_result *SetDomainDnssecStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDomainDnssecStatusResponse{}
	_body, _err := client.SetDomainDnssecStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDomainRecordStatusWithOptions(request *SetDomainRecordStatusRequest, runtime *util.RuntimeOptions) (_result *SetDomainRecordStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDomainRecordStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDomainRecordStatus"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDomainRecordStatus(request *SetDomainRecordStatusRequest) (_result *SetDomainRecordStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDomainRecordStatusResponse{}
	_body, _err := client.SetDomainRecordStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGtmAccessModeWithOptions(request *SetGtmAccessModeRequest, runtime *util.RuntimeOptions) (_result *SetGtmAccessModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetGtmAccessModeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetGtmAccessMode"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGtmAccessMode(request *SetGtmAccessModeRequest) (_result *SetGtmAccessModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGtmAccessModeResponse{}
	_body, _err := client.SetGtmAccessModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGtmMonitorStatusWithOptions(request *SetGtmMonitorStatusRequest, runtime *util.RuntimeOptions) (_result *SetGtmMonitorStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetGtmMonitorStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetGtmMonitorStatus"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGtmMonitorStatus(request *SetGtmMonitorStatusRequest) (_result *SetGtmMonitorStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGtmMonitorStatusResponse{}
	_body, _err := client.SetGtmMonitorStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchDnsGtmInstanceStrategyModeWithOptions(request *SwitchDnsGtmInstanceStrategyModeRequest, runtime *util.RuntimeOptions) (_result *SwitchDnsGtmInstanceStrategyModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SwitchDnsGtmInstanceStrategyModeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SwitchDnsGtmInstanceStrategyMode"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchDnsGtmInstanceStrategyMode(request *SwitchDnsGtmInstanceStrategyModeRequest) (_result *SwitchDnsGtmInstanceStrategyModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchDnsGtmInstanceStrategyModeResponse{}
	_body, _err := client.SwitchDnsGtmInstanceStrategyModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferDomainWithOptions(request *TransferDomainRequest, runtime *util.RuntimeOptions) (_result *TransferDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TransferDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TransferDomain"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferDomain(request *TransferDomainRequest) (_result *TransferDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferDomainResponse{}
	_body, _err := client.TransferDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindInstanceDomainsWithOptions(request *UnbindInstanceDomainsRequest, runtime *util.RuntimeOptions) (_result *UnbindInstanceDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnbindInstanceDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindInstanceDomains"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindInstanceDomains(request *UnbindInstanceDomainsRequest) (_result *UnbindInstanceDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindInstanceDomainsResponse{}
	_body, _err := client.UnbindInstanceDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCustomLineWithOptions(request *UpdateCustomLineRequest, runtime *util.RuntimeOptions) (_result *UpdateCustomLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCustomLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCustomLine"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCustomLine(request *UpdateCustomLineRequest) (_result *UpdateCustomLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCustomLineResponse{}
	_body, _err := client.UpdateCustomLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDnsGtmAccessStrategyWithOptions(request *UpdateDnsGtmAccessStrategyRequest, runtime *util.RuntimeOptions) (_result *UpdateDnsGtmAccessStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDnsGtmAccessStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDnsGtmAccessStrategy"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDnsGtmAccessStrategy(request *UpdateDnsGtmAccessStrategyRequest) (_result *UpdateDnsGtmAccessStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDnsGtmAccessStrategyResponse{}
	_body, _err := client.UpdateDnsGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDnsGtmAddressPoolWithOptions(request *UpdateDnsGtmAddressPoolRequest, runtime *util.RuntimeOptions) (_result *UpdateDnsGtmAddressPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDnsGtmAddressPoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDnsGtmAddressPool"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDnsGtmAddressPool(request *UpdateDnsGtmAddressPoolRequest) (_result *UpdateDnsGtmAddressPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDnsGtmAddressPoolResponse{}
	_body, _err := client.UpdateDnsGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDnsGtmInstanceGlobalConfigWithOptions(request *UpdateDnsGtmInstanceGlobalConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateDnsGtmInstanceGlobalConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDnsGtmInstanceGlobalConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDnsGtmInstanceGlobalConfig"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDnsGtmInstanceGlobalConfig(request *UpdateDnsGtmInstanceGlobalConfigRequest) (_result *UpdateDnsGtmInstanceGlobalConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDnsGtmInstanceGlobalConfigResponse{}
	_body, _err := client.UpdateDnsGtmInstanceGlobalConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDnsGtmMonitorWithOptions(request *UpdateDnsGtmMonitorRequest, runtime *util.RuntimeOptions) (_result *UpdateDnsGtmMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDnsGtmMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDnsGtmMonitor"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDnsGtmMonitor(request *UpdateDnsGtmMonitorRequest) (_result *UpdateDnsGtmMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDnsGtmMonitorResponse{}
	_body, _err := client.UpdateDnsGtmMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDNSSLBWeightWithOptions(request *UpdateDNSSLBWeightRequest, runtime *util.RuntimeOptions) (_result *UpdateDNSSLBWeightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDNSSLBWeightResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDNSSLBWeight"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDNSSLBWeight(request *UpdateDNSSLBWeightRequest) (_result *UpdateDNSSLBWeightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDNSSLBWeightResponse{}
	_body, _err := client.UpdateDNSSLBWeightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDomainGroupWithOptions(request *UpdateDomainGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateDomainGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDomainGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDomainGroup"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDomainGroup(request *UpdateDomainGroupRequest) (_result *UpdateDomainGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDomainGroupResponse{}
	_body, _err := client.UpdateDomainGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDomainRecordWithOptions(request *UpdateDomainRecordRequest, runtime *util.RuntimeOptions) (_result *UpdateDomainRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDomainRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDomainRecord"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDomainRecord(request *UpdateDomainRecordRequest) (_result *UpdateDomainRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDomainRecordResponse{}
	_body, _err := client.UpdateDomainRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDomainRecordRemarkWithOptions(request *UpdateDomainRecordRemarkRequest, runtime *util.RuntimeOptions) (_result *UpdateDomainRecordRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDomainRecordRemarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDomainRecordRemark"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDomainRecordRemark(request *UpdateDomainRecordRemarkRequest) (_result *UpdateDomainRecordRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDomainRecordRemarkResponse{}
	_body, _err := client.UpdateDomainRecordRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDomainRemarkWithOptions(request *UpdateDomainRemarkRequest, runtime *util.RuntimeOptions) (_result *UpdateDomainRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDomainRemarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDomainRemark"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDomainRemark(request *UpdateDomainRemarkRequest) (_result *UpdateDomainRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDomainRemarkResponse{}
	_body, _err := client.UpdateDomainRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGtmAccessStrategyWithOptions(request *UpdateGtmAccessStrategyRequest, runtime *util.RuntimeOptions) (_result *UpdateGtmAccessStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateGtmAccessStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateGtmAccessStrategy"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGtmAccessStrategy(request *UpdateGtmAccessStrategyRequest) (_result *UpdateGtmAccessStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGtmAccessStrategyResponse{}
	_body, _err := client.UpdateGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGtmAddressPoolWithOptions(request *UpdateGtmAddressPoolRequest, runtime *util.RuntimeOptions) (_result *UpdateGtmAddressPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateGtmAddressPoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateGtmAddressPool"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGtmAddressPool(request *UpdateGtmAddressPoolRequest) (_result *UpdateGtmAddressPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGtmAddressPoolResponse{}
	_body, _err := client.UpdateGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGtmInstanceGlobalConfigWithOptions(request *UpdateGtmInstanceGlobalConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateGtmInstanceGlobalConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateGtmInstanceGlobalConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateGtmInstanceGlobalConfig"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGtmInstanceGlobalConfig(request *UpdateGtmInstanceGlobalConfigRequest) (_result *UpdateGtmInstanceGlobalConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGtmInstanceGlobalConfigResponse{}
	_body, _err := client.UpdateGtmInstanceGlobalConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGtmMonitorWithOptions(request *UpdateGtmMonitorRequest, runtime *util.RuntimeOptions) (_result *UpdateGtmMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateGtmMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateGtmMonitor"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGtmMonitor(request *UpdateGtmMonitorRequest) (_result *UpdateGtmMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGtmMonitorResponse{}
	_body, _err := client.UpdateGtmMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGtmRecoveryPlanWithOptions(request *UpdateGtmRecoveryPlanRequest, runtime *util.RuntimeOptions) (_result *UpdateGtmRecoveryPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateGtmRecoveryPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateGtmRecoveryPlan"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGtmRecoveryPlan(request *UpdateGtmRecoveryPlanRequest) (_result *UpdateGtmRecoveryPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGtmRecoveryPlanResponse{}
	_body, _err := client.UpdateGtmRecoveryPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateDnsGtmAttributeInfoWithOptions(request *ValidateDnsGtmAttributeInfoRequest, runtime *util.RuntimeOptions) (_result *ValidateDnsGtmAttributeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ValidateDnsGtmAttributeInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ValidateDnsGtmAttributeInfo"), tea.String("2015-01-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateDnsGtmAttributeInfo(request *ValidateDnsGtmAttributeInfoRequest) (_result *ValidateDnsGtmAttributeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateDnsGtmAttributeInfoResponse{}
	_body, _err := client.ValidateDnsGtmAttributeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
