// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DataDisk struct {
	// 数据盘
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DataDisk) String() string {
	return tea.Prettify(s)
}

func (s DataDisk) GoString() string {
	return s.String()
}

func (s *DataDisk) SetSize(v int64) *DataDisk {
	s.Size = &v
	return s
}

type TcpConfig struct {
	// 调度算法。取值：wrr（默认值）：权重值越高的后端服务器，被轮询到的次数（概率）也越高。wlc：除了根据每台后端服务器设定的权重值来进行轮询，同时还考虑后端服务器的实际负载（即连接数）。当权重值相同时，当前连接数越小的后端服务器被轮询到的次数（概率）也越高。rr：按照访问顺序依次将外部请求依序分发到后端服务器。sch：基于源IP地址的一致性hash，相同的源地址会调度到相同的后端服务器。
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// 会话保持的超时时间。取值：0~3600（秒）。默认值：0，表示关闭会话保持。
	PersistenceTimeout *int32 `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	// 连接超时时间。取值：10~900（秒）。
	EstablishedTimeout *int32 `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
}

func (s TcpConfig) String() string {
	return tea.Prettify(s)
}

func (s TcpConfig) GoString() string {
	return s.String()
}

func (s *TcpConfig) SetScheduler(v string) *TcpConfig {
	s.Scheduler = &v
	return s
}

func (s *TcpConfig) SetPersistenceTimeout(v int32) *TcpConfig {
	s.PersistenceTimeout = &v
	return s
}

func (s *TcpConfig) SetEstablishedTimeout(v int32) *TcpConfig {
	s.EstablishedTimeout = &v
	return s
}

type UdpCheck struct {
	// 健康检查连续成功多少次后，将后端服务器的健康检查状态由fail判定为success。  取值：2-10。
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// 健康检查连续失败多少次后，将后端服务器的健康检查状态由success判定为fail。  取值：2-10。
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// 接收来自运行状况检查的响应需要等待的时间。  如果后端ENS在指定的时间内没有正确响应，则判定为健康检查失败。  取值：1-300（秒）。默认为5秒
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// 健康检查的时间间隔。  取值：1-50（秒）。
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// 健康检查使用的端口。取值：1-65535  不设置此参数时，表示使用后端服务端口（BackendServerPort）。
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
}

func (s UdpCheck) String() string {
	return tea.Prettify(s)
}

func (s UdpCheck) GoString() string {
	return s.String()
}

func (s *UdpCheck) SetHealthyThreshold(v int32) *UdpCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *UdpCheck) SetUnhealthyThreshold(v int32) *UdpCheck {
	s.UnhealthyThreshold = &v
	return s
}

func (s *UdpCheck) SetHealthCheckConnectTimeout(v int32) *UdpCheck {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *UdpCheck) SetHealthCheckInterval(v int32) *UdpCheck {
	s.HealthCheckInterval = &v
	return s
}

func (s *UdpCheck) SetHealthCheckConnectPort(v int32) *UdpCheck {
	s.HealthCheckConnectPort = &v
	return s
}

type HealthCheck struct {
	// 健康检查连续成功多少次后，将后端服务器的健康检查状态由fail判定为success。  取值：2~10。    说明 在HealthCheck值为on时才会有效。
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// 是否开启健康检查。  取值：on | off。
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// 健康检查连续失败多少次后，将后端服务器的健康检查状态由success判定为fail。  取值：2~10。   说明 在HealthCheck值为on时才会有效。
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// 每次健康检查响应的最大超时时间。  取值：1~300（秒）。  默认值：5。
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// 接收来自运行状况检查的响应需要等待的时间。如果后端ECS在指定的时间内没有正确响应，则判定为健康检查失败。在HealthCheck值为on时才会有效。  取值：1~300（秒）。   说明 如果HealthCHeckTimeout的值小于HealthCheckInterval的值，则HealthCHeckTimeout无效，超时时间为HealthCheckInterval的值。
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// 健康检查的时间间隔。  取值： 1~50（秒）。   说明 在HealthCheck值为on时才会有效。
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// 健康检查的后端服务器的端口。  取值： 1~65535。    说明 在HealthCheck值为on时才会有效。
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// 健康检查类型。  取值：tcp（默认值） | http。
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// 用于健康检查的域名，取值：  $_ip： 后端服务器的私网IP。当指定了IP或该参数未指定时，负载均衡会使用各后端服务器的私网IP当做健康检查使用的域名。是否要支持？ domain：域名长度为1-80字符，只能包含字母、数字、点号（.）和连字符（-）。   说明 在HealthCheck值为on时才会有效。
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// 用于健康检查的URI。  长度限制为1~80，只能使用字母、数字和”-/.%?#&amp;“这些字符。 URL不能只为”/“，但必须以”/“开头。    说明 在HealthCheck值为on时才会有效。
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// 健康检查正常的HTTP状态码，多个状态码用逗号分隔。  默认值为http_2xx。  取值：http_2xx | http_3xx | http_4xx | http_5xx。   说明 在HealthCheck值为on时才会有效。
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// 健康检查的method
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
}

func (s HealthCheck) String() string {
	return tea.Prettify(s)
}

func (s HealthCheck) GoString() string {
	return s.String()
}

func (s *HealthCheck) SetHealthyThreshold(v int32) *HealthCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *HealthCheck) SetHealthCheck(v string) *HealthCheck {
	s.HealthCheck = &v
	return s
}

func (s *HealthCheck) SetUnhealthyThreshold(v int32) *HealthCheck {
	s.UnhealthyThreshold = &v
	return s
}

func (s *HealthCheck) SetHealthCheckConnectTimeout(v int32) *HealthCheck {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *HealthCheck) SetHealthCheckTimeout(v int32) *HealthCheck {
	s.HealthCheckTimeout = &v
	return s
}

func (s *HealthCheck) SetHealthCheckInterval(v int32) *HealthCheck {
	s.HealthCheckInterval = &v
	return s
}

func (s *HealthCheck) SetHealthCheckConnectPort(v int32) *HealthCheck {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *HealthCheck) SetHealthCheckType(v string) *HealthCheck {
	s.HealthCheckType = &v
	return s
}

func (s *HealthCheck) SetHealthCheckDomain(v string) *HealthCheck {
	s.HealthCheckDomain = &v
	return s
}

func (s *HealthCheck) SetHealthCheckURI(v string) *HealthCheck {
	s.HealthCheckURI = &v
	return s
}

func (s *HealthCheck) SetHealthCheckHttpCode(v string) *HealthCheck {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *HealthCheck) SetHealthCheckMethod(v string) *HealthCheck {
	s.HealthCheckMethod = &v
	return s
}

type HttpConfig struct {
	// 调度算法。取值：  wrr（默认值）：权重值越高的后端服务器，被轮询到的次数（概率）也越高。 wlc：除了根据每台后端服务器设定的权重值来进行轮询，同时还考虑后端服务器的实际负载（即连接数）。当权重值相同时，当前连接数越小的后端服务器被轮询到的次数（概率）也越高。 rr：按照访问顺序依次将外部请求依序分发到后端服务器。
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// 是否开启会话保持。  取值：on | off。
	StickySession *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	// 是否开启通过X-Forwarded-For头字段获取来访者真实 IP。  取值为on。
	XForwardedFor *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
	// cookie的处理方式。取值：  insert：植入Cookie。  客户端第一次访问时，负载均衡会在返回请求中植入Cookie（即在HTTP/HTTPS响应报文中插入SERVERID），下次客户端携带此Cookie访问，负载均衡服务会将请求定向转发给之前记录到的后端服务器上。  server：重写Cookie。  负载均衡发现用户自定义了Cookie，将会对原来的Cookie进行重写，下次客户端携带新的Cookie访问，负载均衡服务会将请求定向转发给之前记录到的后端服务器。   说明 当StickySession的值为on时，必须指定该参数。
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	// 服务器上配置的Cookie。 长度为1-200，只能包含ASCII英文字母和数字字符，不能包含逗号、分号或空格，也不能以$开头。 说明 当StickySession为on且StickySessionType为server时，该参数必选。
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// Cookie超时时间。  取值：1~86400（秒）。   说明 当StickySession为on且StickySessionType为insert时，该参数必选。
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// 指定连接空闲超时时间，取值范围为1~60秒，默认值为15秒。  在超时时间内一直没有访问请求，负载均衡会暂时中断当前连接，直到一下次请求来临时重新建立新的连接。
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// 指定请求超时时间，取值范围为1~180秒，默认值为60秒。  在超时时间内后端服务器一直没有响应，负载均衡将放弃等待，给客户端返回 HTTP 504 错误码。
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// 服务器证书的ID。
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
}

func (s HttpConfig) String() string {
	return tea.Prettify(s)
}

func (s HttpConfig) GoString() string {
	return s.String()
}

func (s *HttpConfig) SetScheduler(v string) *HttpConfig {
	s.Scheduler = &v
	return s
}

func (s *HttpConfig) SetStickySession(v string) *HttpConfig {
	s.StickySession = &v
	return s
}

func (s *HttpConfig) SetXForwardedFor(v string) *HttpConfig {
	s.XForwardedFor = &v
	return s
}

func (s *HttpConfig) SetStickySessionType(v string) *HttpConfig {
	s.StickySessionType = &v
	return s
}

func (s *HttpConfig) SetCookie(v string) *HttpConfig {
	s.Cookie = &v
	return s
}

func (s *HttpConfig) SetCookieTimeout(v int32) *HttpConfig {
	s.CookieTimeout = &v
	return s
}

func (s *HttpConfig) SetIdleTimeout(v int32) *HttpConfig {
	s.IdleTimeout = &v
	return s
}

func (s *HttpConfig) SetRequestTimeout(v int32) *HttpConfig {
	s.RequestTimeout = &v
	return s
}

func (s *HttpConfig) SetServerCertificateId(v string) *HttpConfig {
	s.ServerCertificateId = &v
	return s
}

type SecurityGroupRule struct {
	// 方向
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// 协议
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// 目的端口
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// 源端口
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// 授权策略
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// 目标网段
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// 源网段
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 优先级
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
}

func (s SecurityGroupRule) String() string {
	return tea.Prettify(s)
}

func (s SecurityGroupRule) GoString() string {
	return s.String()
}

func (s *SecurityGroupRule) SetDirection(v string) *SecurityGroupRule {
	s.Direction = &v
	return s
}

func (s *SecurityGroupRule) SetIpProtocol(v string) *SecurityGroupRule {
	s.IpProtocol = &v
	return s
}

func (s *SecurityGroupRule) SetPortRange(v string) *SecurityGroupRule {
	s.PortRange = &v
	return s
}

func (s *SecurityGroupRule) SetSourcePortRange(v string) *SecurityGroupRule {
	s.SourcePortRange = &v
	return s
}

func (s *SecurityGroupRule) SetPolicy(v string) *SecurityGroupRule {
	s.Policy = &v
	return s
}

func (s *SecurityGroupRule) SetDestCidrIp(v string) *SecurityGroupRule {
	s.DestCidrIp = &v
	return s
}

func (s *SecurityGroupRule) SetSourceCidrIp(v string) *SecurityGroupRule {
	s.SourceCidrIp = &v
	return s
}

func (s *SecurityGroupRule) SetDescription(v string) *SecurityGroupRule {
	s.Description = &v
	return s
}

func (s *SecurityGroupRule) SetPriority(v int32) *SecurityGroupRule {
	s.Priority = &v
	return s
}

type UdpConfig struct {
	// 调度算法。取值：  wrr（默认值）：权重值越高的后端服务器，被轮询到的次数（概率）也越高。 wlc：除了根据每台后端服务器设定的权重值来进行轮询，同时还考虑后端服务器的实际负载（即连接数）。当权重值相同时，当前连接数越小的后端服务器被轮询到的次数（概率）也越高。 rr：按照访问顺序依次将外部请求依序分发到后端服务器。 sch：基于源IP地址的一致性hash，相同的源地址会调度到相同的后端服务器。
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// hash key
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
}

func (s UdpConfig) String() string {
	return tea.Prettify(s)
}

func (s UdpConfig) GoString() string {
	return s.String()
}

func (s *UdpConfig) SetScheduler(v string) *UdpConfig {
	s.Scheduler = &v
	return s
}

func (s *UdpConfig) SetHashKey(v string) *UdpConfig {
	s.HashKey = &v
	return s
}

type DescribeServcieScheduleRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	PodConfigName *string `json:"PodConfigName,omitempty" xml:"PodConfigName,omitempty"`
}

func (s DescribeServcieScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServcieScheduleRequest) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleRequest) SetAppId(v string) *DescribeServcieScheduleRequest {
	s.AppId = &v
	return s
}

func (s *DescribeServcieScheduleRequest) SetUuid(v string) *DescribeServcieScheduleRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeServcieScheduleRequest) SetPodConfigName(v string) *DescribeServcieScheduleRequest {
	s.PodConfigName = &v
	return s
}

type DescribeServcieScheduleResponseBody struct {
	Index           *int32                                              `json:"Index,omitempty" xml:"Index,omitempty"`
	InstanceId      *string                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIp      *string                                             `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	InstancePort    *int32                                              `json:"InstancePort,omitempty" xml:"InstancePort,omitempty"`
	PodAbstractInfo *DescribeServcieScheduleResponseBodyPodAbstractInfo `json:"PodAbstractInfo,omitempty" xml:"PodAbstractInfo,omitempty" type:"Struct"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestRepeated *bool                                               `json:"RequestRepeated,omitempty" xml:"RequestRepeated,omitempty"`
	TcpPorts        *string                                             `json:"TcpPorts,omitempty" xml:"TcpPorts,omitempty"`
}

func (s DescribeServcieScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServcieScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleResponseBody) SetIndex(v int32) *DescribeServcieScheduleResponseBody {
	s.Index = &v
	return s
}

func (s *DescribeServcieScheduleResponseBody) SetInstanceId(v string) *DescribeServcieScheduleResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeServcieScheduleResponseBody) SetInstanceIp(v string) *DescribeServcieScheduleResponseBody {
	s.InstanceIp = &v
	return s
}

func (s *DescribeServcieScheduleResponseBody) SetInstancePort(v int32) *DescribeServcieScheduleResponseBody {
	s.InstancePort = &v
	return s
}

func (s *DescribeServcieScheduleResponseBody) SetPodAbstractInfo(v *DescribeServcieScheduleResponseBodyPodAbstractInfo) *DescribeServcieScheduleResponseBody {
	s.PodAbstractInfo = v
	return s
}

func (s *DescribeServcieScheduleResponseBody) SetRequestId(v string) *DescribeServcieScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServcieScheduleResponseBody) SetRequestRepeated(v bool) *DescribeServcieScheduleResponseBody {
	s.RequestRepeated = &v
	return s
}

func (s *DescribeServcieScheduleResponseBody) SetTcpPorts(v string) *DescribeServcieScheduleResponseBody {
	s.TcpPorts = &v
	return s
}

type DescribeServcieScheduleResponseBodyPodAbstractInfo struct {
	ContainerService  *bool                                                                `json:"ContainerService,omitempty" xml:"ContainerService,omitempty"`
	ContainerStatuses *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses `json:"ContainerStatuses,omitempty" xml:"ContainerStatuses,omitempty" type:"Struct"`
	Name              *bool                                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace         *bool                                                                `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ResourceScope     *bool                                                                `json:"ResourceScope,omitempty" xml:"ResourceScope,omitempty"`
	Status            *bool                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeServcieScheduleResponseBodyPodAbstractInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeServcieScheduleResponseBodyPodAbstractInfo) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) SetContainerService(v bool) *DescribeServcieScheduleResponseBodyPodAbstractInfo {
	s.ContainerService = &v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) SetContainerStatuses(v *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses) *DescribeServcieScheduleResponseBodyPodAbstractInfo {
	s.ContainerStatuses = v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) SetName(v bool) *DescribeServcieScheduleResponseBodyPodAbstractInfo {
	s.Name = &v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) SetNamespace(v bool) *DescribeServcieScheduleResponseBodyPodAbstractInfo {
	s.Namespace = &v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) SetResourceScope(v bool) *DescribeServcieScheduleResponseBodyPodAbstractInfo {
	s.ResourceScope = &v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) SetStatus(v bool) *DescribeServcieScheduleResponseBodyPodAbstractInfo {
	s.Status = &v
	return s
}

type DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses struct {
	ContainerStatus []*DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus `json:"ContainerStatus,omitempty" xml:"ContainerStatus,omitempty" type:"Repeated"`
}

func (s DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses) String() string {
	return tea.Prettify(s)
}

func (s DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses) SetContainerStatus(v []*DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus) *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses {
	s.ContainerStatus = v
	return s
}

type DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus struct {
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus) SetContainerId(v string) *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus {
	s.ContainerId = &v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus) SetName(v string) *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus {
	s.Name = &v
	return s
}

type DescribeServcieScheduleResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeServcieScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServcieScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServcieScheduleResponse) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleResponse) SetHeaders(v map[string]*string) *DescribeServcieScheduleResponse {
	s.Headers = v
	return s
}

func (s *DescribeServcieScheduleResponse) SetBody(v *DescribeServcieScheduleResponseBody) *DescribeServcieScheduleResponse {
	s.Body = v
	return s
}

type JoinSecurityGroupRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s JoinSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *JoinSecurityGroupRequest) SetVersion(v string) *JoinSecurityGroupRequest {
	s.Version = &v
	return s
}

func (s *JoinSecurityGroupRequest) SetSecurityGroupId(v string) *JoinSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *JoinSecurityGroupRequest) SetInstanceId(v string) *JoinSecurityGroupRequest {
	s.InstanceId = &v
	return s
}

type JoinSecurityGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JoinSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *JoinSecurityGroupResponseBody) SetRequestId(v string) *JoinSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

type JoinSecurityGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *JoinSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *JoinSecurityGroupResponse) SetHeaders(v map[string]*string) *JoinSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *JoinSecurityGroupResponse) SetBody(v *JoinSecurityGroupResponseBody) *JoinSecurityGroupResponse {
	s.Body = v
	return s
}

type RestartDeviceInstanceRequest struct {
	// App ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RestartDeviceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDeviceInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDeviceInstanceRequest) SetAppId(v string) *RestartDeviceInstanceRequest {
	s.AppId = &v
	return s
}

func (s *RestartDeviceInstanceRequest) SetInstanceId(v string) *RestartDeviceInstanceRequest {
	s.InstanceId = &v
	return s
}

type RestartDeviceInstanceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDeviceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartDeviceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDeviceInstanceResponseBody) SetRequestId(v string) *RestartDeviceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartDeviceInstanceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartDeviceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartDeviceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartDeviceInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartDeviceInstanceResponse) SetHeaders(v map[string]*string) *RestartDeviceInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartDeviceInstanceResponse) SetBody(v *RestartDeviceInstanceResponseBody) *RestartDeviceInstanceResponse {
	s.Body = v
	return s
}

type ReleasePrePaidInstanceRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleasePrePaidInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleasePrePaidInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleasePrePaidInstanceRequest) SetVersion(v string) *ReleasePrePaidInstanceRequest {
	s.Version = &v
	return s
}

func (s *ReleasePrePaidInstanceRequest) SetInstanceId(v string) *ReleasePrePaidInstanceRequest {
	s.InstanceId = &v
	return s
}

type ReleasePrePaidInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleasePrePaidInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleasePrePaidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePrePaidInstanceResponseBody) SetRequestId(v string) *ReleasePrePaidInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ReleasePrePaidInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleasePrePaidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleasePrePaidInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleasePrePaidInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleasePrePaidInstanceResponse) SetHeaders(v map[string]*string) *ReleasePrePaidInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleasePrePaidInstanceResponse) SetBody(v *ReleasePrePaidInstanceResponseBody) *ReleasePrePaidInstanceResponse {
	s.Body = v
	return s
}

type CreateApplicationRequest struct {
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Timeout  *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) SetTemplate(v string) *CreateApplicationRequest {
	s.Template = &v
	return s
}

func (s *CreateApplicationRequest) SetTimeout(v int32) *CreateApplicationRequest {
	s.Timeout = &v
	return s
}

type CreateApplicationResponseBody struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) SetAppId(v string) *CreateApplicationResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

type CreateApplicationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponse) SetHeaders(v map[string]*string) *CreateApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationResponse) SetBody(v *CreateApplicationResponseBody) *CreateApplicationResponse {
	s.Body = v
	return s
}

type CreateNetworkRequest struct {
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	NetworkName *string `json:"NetworkName,omitempty" xml:"NetworkName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkRequest) SetEnsRegionId(v string) *CreateNetworkRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateNetworkRequest) SetCidrBlock(v string) *CreateNetworkRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateNetworkRequest) SetNetworkName(v string) *CreateNetworkRequest {
	s.NetworkName = &v
	return s
}

func (s *CreateNetworkRequest) SetDescription(v string) *CreateNetworkRequest {
	s.Description = &v
	return s
}

type CreateNetworkResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
}

func (s CreateNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkResponseBody) SetRequestId(v string) *CreateNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkResponseBody) SetNetworkId(v string) *CreateNetworkResponseBody {
	s.NetworkId = &v
	return s
}

type CreateNetworkResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkResponse) SetHeaders(v map[string]*string) *CreateNetworkResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkResponse) SetBody(v *CreateNetworkResponseBody) *CreateNetworkResponse {
	s.Body = v
	return s
}

type ReInitDiskRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	DiskId  *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s ReInitDiskRequest) String() string {
	return tea.Prettify(s)
}

func (s ReInitDiskRequest) GoString() string {
	return s.String()
}

func (s *ReInitDiskRequest) SetVersion(v string) *ReInitDiskRequest {
	s.Version = &v
	return s
}

func (s *ReInitDiskRequest) SetDiskId(v string) *ReInitDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ReInitDiskRequest) SetImageId(v string) *ReInitDiskRequest {
	s.ImageId = &v
	return s
}

type ReInitDiskResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReInitDiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReInitDiskResponseBody) GoString() string {
	return s.String()
}

func (s *ReInitDiskResponseBody) SetCode(v int32) *ReInitDiskResponseBody {
	s.Code = &v
	return s
}

func (s *ReInitDiskResponseBody) SetRequestId(v string) *ReInitDiskResponseBody {
	s.RequestId = &v
	return s
}

type ReInitDiskResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReInitDiskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReInitDiskResponse) String() string {
	return tea.Prettify(s)
}

func (s ReInitDiskResponse) GoString() string {
	return s.String()
}

func (s *ReInitDiskResponse) SetHeaders(v map[string]*string) *ReInitDiskResponse {
	s.Headers = v
	return s
}

func (s *ReInitDiskResponse) SetBody(v *ReInitDiskResponseBody) *ReInitDiskResponse {
	s.Body = v
	return s
}

type AddNetworkInterfaceToInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Networks   *string `json:"Networks,omitempty" xml:"Networks,omitempty"`
}

func (s AddNetworkInterfaceToInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddNetworkInterfaceToInstanceRequest) GoString() string {
	return s.String()
}

func (s *AddNetworkInterfaceToInstanceRequest) SetInstanceId(v string) *AddNetworkInterfaceToInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AddNetworkInterfaceToInstanceRequest) SetNetworks(v string) *AddNetworkInterfaceToInstanceRequest {
	s.Networks = &v
	return s
}

type AddNetworkInterfaceToInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddNetworkInterfaceToInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddNetworkInterfaceToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AddNetworkInterfaceToInstanceResponseBody) SetRequestId(v string) *AddNetworkInterfaceToInstanceResponseBody {
	s.RequestId = &v
	return s
}

type AddNetworkInterfaceToInstanceResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddNetworkInterfaceToInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddNetworkInterfaceToInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddNetworkInterfaceToInstanceResponse) GoString() string {
	return s.String()
}

func (s *AddNetworkInterfaceToInstanceResponse) SetHeaders(v map[string]*string) *AddNetworkInterfaceToInstanceResponse {
	s.Headers = v
	return s
}

func (s *AddNetworkInterfaceToInstanceResponse) SetBody(v *AddNetworkInterfaceToInstanceResponseBody) *AddNetworkInterfaceToInstanceResponse {
	s.Body = v
	return s
}

type AssociateEipAddressRequest struct {
	Version              *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsRegionId          *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	Eip                  *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	InstanceIdInternetIp *string `json:"InstanceIdInternetIp,omitempty" xml:"InstanceIdInternetIp,omitempty"`
}

func (s AssociateEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateEipAddressRequest) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressRequest) SetVersion(v string) *AssociateEipAddressRequest {
	s.Version = &v
	return s
}

func (s *AssociateEipAddressRequest) SetEnsRegionId(v string) *AssociateEipAddressRequest {
	s.EnsRegionId = &v
	return s
}

func (s *AssociateEipAddressRequest) SetEip(v string) *AssociateEipAddressRequest {
	s.Eip = &v
	return s
}

func (s *AssociateEipAddressRequest) SetInstanceIdInternetIp(v string) *AssociateEipAddressRequest {
	s.InstanceIdInternetIp = &v
	return s
}

type AssociateEipAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateEipAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressResponseBody) SetRequestId(v string) *AssociateEipAddressResponseBody {
	s.RequestId = &v
	return s
}

type AssociateEipAddressResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssociateEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateEipAddressResponse) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressResponse) SetHeaders(v map[string]*string) *AssociateEipAddressResponse {
	s.Headers = v
	return s
}

func (s *AssociateEipAddressResponse) SetBody(v *AssociateEipAddressResponseBody) *AssociateEipAddressResponse {
	s.Body = v
	return s
}

type StopEpnInstanceRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
}

func (s StopEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopEpnInstanceRequest) SetEPNInstanceId(v string) *StopEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

type StopEpnInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopEpnInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopEpnInstanceResponseBody) SetRequestId(v string) *StopEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopEpnInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopEpnInstanceResponse) SetHeaders(v map[string]*string) *StopEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopEpnInstanceResponse) SetBody(v *StopEpnInstanceResponseBody) *StopEpnInstanceResponse {
	s.Body = v
	return s
}

type DescribeEipAddressesRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	Eips        *string `json:"Eips,omitempty" xml:"Eips,omitempty"`
}

func (s DescribeEipAddressesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEipAddressesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesRequest) SetVersion(v string) *DescribeEipAddressesRequest {
	s.Version = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetEnsRegionId(v string) *DescribeEipAddressesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetEips(v string) *DescribeEipAddressesRequest {
	s.Eips = &v
	return s
}

type DescribeEipAddressesResponseBody struct {
	EipAddresses *DescribeEipAddressesResponseBodyEipAddresses `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty" type:"Struct"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEipAddressesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEipAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBody) SetEipAddresses(v *DescribeEipAddressesResponseBodyEipAddresses) *DescribeEipAddressesResponseBody {
	s.EipAddresses = v
	return s
}

func (s *DescribeEipAddressesResponseBody) SetRequestId(v string) *DescribeEipAddressesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEipAddressesResponseBodyEipAddresses struct {
	EipAddress []*DescribeEipAddressesResponseBodyEipAddressesEipAddress `json:"EipAddress,omitempty" xml:"EipAddress,omitempty" type:"Repeated"`
}

func (s DescribeEipAddressesResponseBodyEipAddresses) String() string {
	return tea.Prettify(s)
}

func (s DescribeEipAddressesResponseBodyEipAddresses) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBodyEipAddresses) SetEipAddress(v []*DescribeEipAddressesResponseBodyEipAddressesEipAddress) *DescribeEipAddressesResponseBodyEipAddresses {
	s.EipAddress = v
	return s
}

type DescribeEipAddressesResponseBodyEipAddressesEipAddress struct {
	Eip                  *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	InstanceIdInternetIp *string `json:"InstanceIdInternetIp,omitempty" xml:"InstanceIdInternetIp,omitempty"`
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddress) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetEip(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.Eip = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetInstanceIdInternetIp(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.InstanceIdInternetIp = &v
	return s
}

type DescribeEipAddressesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEipAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEipAddressesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEipAddressesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponse) SetHeaders(v map[string]*string) *DescribeEipAddressesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEipAddressesResponse) SetBody(v *DescribeEipAddressesResponseBody) *DescribeEipAddressesResponse {
	s.Body = v
	return s
}

type RevokeSecurityGroupEgressRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Priority        *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	DestCidrIp      *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s RevokeSecurityGroupEgressRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeSecurityGroupEgressRequest) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupEgressRequest) SetVersion(v string) *RevokeSecurityGroupEgressRequest {
	s.Version = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetIpProtocol(v string) *RevokeSecurityGroupEgressRequest {
	s.IpProtocol = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetPortRange(v string) *RevokeSecurityGroupEgressRequest {
	s.PortRange = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetSecurityGroupId(v string) *RevokeSecurityGroupEgressRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetPolicy(v string) *RevokeSecurityGroupEgressRequest {
	s.Policy = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetPriority(v int32) *RevokeSecurityGroupEgressRequest {
	s.Priority = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetDestCidrIp(v string) *RevokeSecurityGroupEgressRequest {
	s.DestCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetSourcePortRange(v string) *RevokeSecurityGroupEgressRequest {
	s.SourcePortRange = &v
	return s
}

type RevokeSecurityGroupEgressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeSecurityGroupEgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeSecurityGroupEgressResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupEgressResponseBody) SetRequestId(v string) *RevokeSecurityGroupEgressResponseBody {
	s.RequestId = &v
	return s
}

type RevokeSecurityGroupEgressResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RevokeSecurityGroupEgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeSecurityGroupEgressResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeSecurityGroupEgressResponse) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupEgressResponse) SetHeaders(v map[string]*string) *RevokeSecurityGroupEgressResponse {
	s.Headers = v
	return s
}

func (s *RevokeSecurityGroupEgressResponse) SetBody(v *RevokeSecurityGroupEgressResponseBody) *RevokeSecurityGroupEgressResponse {
	s.Body = v
	return s
}

type ConfigureSecurityGroupPermissionsRequest struct {
	SecurityGroupId      *string                                                         `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	AuthorizePermissions []*ConfigureSecurityGroupPermissionsRequestAuthorizePermissions `json:"AuthorizePermissions,omitempty" xml:"AuthorizePermissions,omitempty" type:"Repeated"`
	RevokePermissions    []*ConfigureSecurityGroupPermissionsRequestRevokePermissions    `json:"RevokePermissions,omitempty" xml:"RevokePermissions,omitempty" type:"Repeated"`
}

func (s ConfigureSecurityGroupPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSecurityGroupPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSecurityGroupPermissionsRequest) SetSecurityGroupId(v string) *ConfigureSecurityGroupPermissionsRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequest) SetAuthorizePermissions(v []*ConfigureSecurityGroupPermissionsRequestAuthorizePermissions) *ConfigureSecurityGroupPermissionsRequest {
	s.AuthorizePermissions = v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequest) SetRevokePermissions(v []*ConfigureSecurityGroupPermissionsRequestRevokePermissions) *ConfigureSecurityGroupPermissionsRequest {
	s.RevokePermissions = v
	return s
}

type ConfigureSecurityGroupPermissionsRequestAuthorizePermissions struct {
	Direction       *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	DestCidrIp      *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	SourceCidrIp    *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Priority        *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s ConfigureSecurityGroupPermissionsRequestAuthorizePermissions) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSecurityGroupPermissionsRequestAuthorizePermissions) GoString() string {
	return s.String()
}

func (s *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions) SetDirection(v string) *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions {
	s.Direction = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions) SetIpProtocol(v string) *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions {
	s.IpProtocol = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions) SetPortRange(v string) *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions {
	s.PortRange = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions) SetSourcePortRange(v string) *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions {
	s.SourcePortRange = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions) SetPolicy(v string) *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions {
	s.Policy = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions) SetDestCidrIp(v string) *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions {
	s.DestCidrIp = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions) SetSourceCidrIp(v string) *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions {
	s.SourceCidrIp = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions) SetDescription(v string) *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions {
	s.Description = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions) SetPriority(v int32) *ConfigureSecurityGroupPermissionsRequestAuthorizePermissions {
	s.Priority = &v
	return s
}

type ConfigureSecurityGroupPermissionsRequestRevokePermissions struct {
	Direction       *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	DestCidrIp      *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	SourceCidrIp    *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	Priority        *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s ConfigureSecurityGroupPermissionsRequestRevokePermissions) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSecurityGroupPermissionsRequestRevokePermissions) GoString() string {
	return s.String()
}

func (s *ConfigureSecurityGroupPermissionsRequestRevokePermissions) SetDirection(v string) *ConfigureSecurityGroupPermissionsRequestRevokePermissions {
	s.Direction = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestRevokePermissions) SetIpProtocol(v string) *ConfigureSecurityGroupPermissionsRequestRevokePermissions {
	s.IpProtocol = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestRevokePermissions) SetPortRange(v string) *ConfigureSecurityGroupPermissionsRequestRevokePermissions {
	s.PortRange = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestRevokePermissions) SetSourcePortRange(v string) *ConfigureSecurityGroupPermissionsRequestRevokePermissions {
	s.SourcePortRange = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestRevokePermissions) SetPolicy(v string) *ConfigureSecurityGroupPermissionsRequestRevokePermissions {
	s.Policy = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestRevokePermissions) SetDestCidrIp(v string) *ConfigureSecurityGroupPermissionsRequestRevokePermissions {
	s.DestCidrIp = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestRevokePermissions) SetSourceCidrIp(v string) *ConfigureSecurityGroupPermissionsRequestRevokePermissions {
	s.SourceCidrIp = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsRequestRevokePermissions) SetPriority(v int32) *ConfigureSecurityGroupPermissionsRequestRevokePermissions {
	s.Priority = &v
	return s
}

type ConfigureSecurityGroupPermissionsShrinkRequest struct {
	SecurityGroupId            *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	AuthorizePermissionsShrink *string `json:"AuthorizePermissions,omitempty" xml:"AuthorizePermissions,omitempty"`
	RevokePermissionsShrink    *string `json:"RevokePermissions,omitempty" xml:"RevokePermissions,omitempty"`
}

func (s ConfigureSecurityGroupPermissionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSecurityGroupPermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSecurityGroupPermissionsShrinkRequest) SetSecurityGroupId(v string) *ConfigureSecurityGroupPermissionsShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsShrinkRequest) SetAuthorizePermissionsShrink(v string) *ConfigureSecurityGroupPermissionsShrinkRequest {
	s.AuthorizePermissionsShrink = &v
	return s
}

func (s *ConfigureSecurityGroupPermissionsShrinkRequest) SetRevokePermissionsShrink(v string) *ConfigureSecurityGroupPermissionsShrinkRequest {
	s.RevokePermissionsShrink = &v
	return s
}

type ConfigureSecurityGroupPermissionsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigureSecurityGroupPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSecurityGroupPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSecurityGroupPermissionsResponseBody) SetRequestId(v string) *ConfigureSecurityGroupPermissionsResponseBody {
	s.RequestId = &v
	return s
}

type ConfigureSecurityGroupPermissionsResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigureSecurityGroupPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigureSecurityGroupPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSecurityGroupPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ConfigureSecurityGroupPermissionsResponse) SetHeaders(v map[string]*string) *ConfigureSecurityGroupPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ConfigureSecurityGroupPermissionsResponse) SetBody(v *ConfigureSecurityGroupPermissionsResponseBody) *ConfigureSecurityGroupPermissionsResponse {
	s.Body = v
	return s
}

type DescribeEnsNetLevelRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeEnsNetLevelRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetLevelRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetLevelRequest) SetVersion(v string) *DescribeEnsNetLevelRequest {
	s.Version = &v
	return s
}

type DescribeEnsNetLevelResponseBody struct {
	Code         *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	EnsNetLevels *DescribeEnsNetLevelResponseBodyEnsNetLevels `json:"EnsNetLevels,omitempty" xml:"EnsNetLevels,omitempty" type:"Struct"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnsNetLevelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetLevelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetLevelResponseBody) SetCode(v int32) *DescribeEnsNetLevelResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnsNetLevelResponseBody) SetEnsNetLevels(v *DescribeEnsNetLevelResponseBodyEnsNetLevels) *DescribeEnsNetLevelResponseBody {
	s.EnsNetLevels = v
	return s
}

func (s *DescribeEnsNetLevelResponseBody) SetRequestId(v string) *DescribeEnsNetLevelResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEnsNetLevelResponseBodyEnsNetLevels struct {
	EnsNetLevel []*DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel `json:"EnsNetLevel,omitempty" xml:"EnsNetLevel,omitempty" type:"Repeated"`
}

func (s DescribeEnsNetLevelResponseBodyEnsNetLevels) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetLevelResponseBodyEnsNetLevels) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetLevelResponseBodyEnsNetLevels) SetEnsNetLevel(v []*DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel) *DescribeEnsNetLevelResponseBodyEnsNetLevels {
	s.EnsNetLevel = v
	return s
}

type DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel struct {
	EnsNetLevelCode *string `json:"EnsNetLevelCode,omitempty" xml:"EnsNetLevelCode,omitempty"`
}

func (s DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel) SetEnsNetLevelCode(v string) *DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel {
	s.EnsNetLevelCode = &v
	return s
}

type DescribeEnsNetLevelResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEnsNetLevelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEnsNetLevelResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetLevelResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetLevelResponse) SetHeaders(v map[string]*string) *DescribeEnsNetLevelResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsNetLevelResponse) SetBody(v *DescribeEnsNetLevelResponseBody) *DescribeEnsNetLevelResponse {
	s.Body = v
	return s
}

type DescribeNetworkAttributeRequest struct {
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
}

func (s DescribeNetworkAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeRequest) SetNetworkId(v string) *DescribeNetworkAttributeRequest {
	s.NetworkId = &v
	return s
}

type DescribeNetworkAttributeResponseBody struct {
	// Id of the request
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EnsRegionId    *string                                             `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	NetworkId      *string                                             `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	NetworkName    *string                                             `json:"NetworkName,omitempty" xml:"NetworkName,omitempty"`
	CidrBlock      *string                                             `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Status         *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Description    *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedTime    *string                                             `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	VSwitchIds     *DescribeNetworkAttributeResponseBodyVSwitchIds     `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	CloudResources *DescribeNetworkAttributeResponseBodyCloudResources `json:"CloudResources,omitempty" xml:"CloudResources,omitempty" type:"Struct"`
}

func (s DescribeNetworkAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponseBody) SetRequestId(v string) *DescribeNetworkAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetEnsRegionId(v string) *DescribeNetworkAttributeResponseBody {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetNetworkId(v string) *DescribeNetworkAttributeResponseBody {
	s.NetworkId = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetNetworkName(v string) *DescribeNetworkAttributeResponseBody {
	s.NetworkName = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetCidrBlock(v string) *DescribeNetworkAttributeResponseBody {
	s.CidrBlock = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetStatus(v string) *DescribeNetworkAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetDescription(v string) *DescribeNetworkAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetCreatedTime(v string) *DescribeNetworkAttributeResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetVSwitchIds(v *DescribeNetworkAttributeResponseBodyVSwitchIds) *DescribeNetworkAttributeResponseBody {
	s.VSwitchIds = v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetCloudResources(v *DescribeNetworkAttributeResponseBodyCloudResources) *DescribeNetworkAttributeResponseBody {
	s.CloudResources = v
	return s
}

type DescribeNetworkAttributeResponseBodyVSwitchIds struct {
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAttributeResponseBodyVSwitchIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkAttributeResponseBodyVSwitchIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponseBodyVSwitchIds) SetVSwitchId(v []*string) *DescribeNetworkAttributeResponseBodyVSwitchIds {
	s.VSwitchId = v
	return s
}

type DescribeNetworkAttributeResponseBodyCloudResources struct {
	CloudResourceSetType []*DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType `json:"CloudResourceSetType,omitempty" xml:"CloudResourceSetType,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAttributeResponseBodyCloudResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkAttributeResponseBodyCloudResources) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponseBodyCloudResources) SetCloudResourceSetType(v []*DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType) *DescribeNetworkAttributeResponseBodyCloudResources {
	s.CloudResourceSetType = v
	return s
}

type DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType struct {
	ResourceCount *int32  `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType) SetResourceCount(v int32) *DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType {
	s.ResourceCount = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType) SetResourceType(v string) *DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType {
	s.ResourceType = &v
	return s
}

type DescribeNetworkAttributeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNetworkAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNetworkAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponse) SetHeaders(v map[string]*string) *DescribeNetworkAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkAttributeResponse) SetBody(v *DescribeNetworkAttributeResponseBody) *DescribeNetworkAttributeResponse {
	s.Body = v
	return s
}

type JoinPublicIpsToEpnInstanceRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	InstanceInfos *string `json:"InstanceInfos,omitempty" xml:"InstanceInfos,omitempty"`
}

func (s JoinPublicIpsToEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinPublicIpsToEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *JoinPublicIpsToEpnInstanceRequest) SetEPNInstanceId(v string) *JoinPublicIpsToEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *JoinPublicIpsToEpnInstanceRequest) SetInstanceInfos(v string) *JoinPublicIpsToEpnInstanceRequest {
	s.InstanceInfos = &v
	return s
}

type JoinPublicIpsToEpnInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JoinPublicIpsToEpnInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinPublicIpsToEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *JoinPublicIpsToEpnInstanceResponseBody) SetRequestId(v string) *JoinPublicIpsToEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

type JoinPublicIpsToEpnInstanceResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *JoinPublicIpsToEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinPublicIpsToEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinPublicIpsToEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *JoinPublicIpsToEpnInstanceResponse) SetHeaders(v map[string]*string) *JoinPublicIpsToEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *JoinPublicIpsToEpnInstanceResponse) SetBody(v *JoinPublicIpsToEpnInstanceResponseBody) *JoinPublicIpsToEpnInstanceResponse {
	s.Body = v
	return s
}

type StopInstanceRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ForceStop  *string `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) SetVersion(v string) *StopInstanceRequest {
	s.Version = &v
	return s
}

func (s *StopInstanceRequest) SetInstanceId(v string) *StopInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceRequest) SetForceStop(v string) *StopInstanceRequest {
	s.ForceStop = &v
	return s
}

type StopInstanceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) SetCode(v int32) *StopInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopInstanceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopInstanceResponse) SetHeaders(v map[string]*string) *StopInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopInstanceResponse) SetBody(v *StopInstanceResponseBody) *StopInstanceResponse {
	s.Body = v
	return s
}

type AttachEnsInstancesRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Scripts    *string `json:"Scripts,omitempty" xml:"Scripts,omitempty"`
}

func (s AttachEnsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachEnsInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachEnsInstancesRequest) SetVersion(v string) *AttachEnsInstancesRequest {
	s.Version = &v
	return s
}

func (s *AttachEnsInstancesRequest) SetInstanceId(v string) *AttachEnsInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachEnsInstancesRequest) SetScripts(v string) *AttachEnsInstancesRequest {
	s.Scripts = &v
	return s
}

type AttachEnsInstancesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachEnsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachEnsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachEnsInstancesResponseBody) SetRequestId(v string) *AttachEnsInstancesResponseBody {
	s.RequestId = &v
	return s
}

type AttachEnsInstancesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachEnsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachEnsInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachEnsInstancesResponse) GoString() string {
	return s.String()
}

func (s *AttachEnsInstancesResponse) SetHeaders(v map[string]*string) *AttachEnsInstancesResponse {
	s.Headers = v
	return s
}

func (s *AttachEnsInstancesResponse) SetBody(v *AttachEnsInstancesResponseBody) *AttachEnsInstancesResponse {
	s.Body = v
	return s
}

type DescribeEpnBandwitdhByInternetChargeTypeRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Isp             *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	EnsRegionId     *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	NetworkingModel *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
}

func (s DescribeEpnBandwitdhByInternetChargeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnBandwitdhByInternetChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetVersion(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.Version = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetStartTime(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetEndTime(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetIsp(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.Isp = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetEnsRegionId(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetNetworkingModel(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.NetworkingModel = &v
	return s
}

type DescribeEpnBandwitdhByInternetChargeTypeResponseBody struct {
	BandwidthValue     *int64  `json:"BandwidthValue,omitempty" xml:"BandwidthValue,omitempty"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeStamp          *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeEpnBandwitdhByInternetChargeTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnBandwitdhByInternetChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) SetBandwidthValue(v int64) *DescribeEpnBandwitdhByInternetChargeTypeResponseBody {
	s.BandwidthValue = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) SetInternetChargeType(v string) *DescribeEpnBandwitdhByInternetChargeTypeResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) SetRequestId(v string) *DescribeEpnBandwitdhByInternetChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) SetTimeStamp(v string) *DescribeEpnBandwitdhByInternetChargeTypeResponseBody {
	s.TimeStamp = &v
	return s
}

type DescribeEpnBandwitdhByInternetChargeTypeResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEpnBandwitdhByInternetChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEpnBandwitdhByInternetChargeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnBandwitdhByInternetChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponse) SetHeaders(v map[string]*string) *DescribeEpnBandwitdhByInternetChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponse) SetBody(v *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) *DescribeEpnBandwitdhByInternetChargeTypeResponse {
	s.Body = v
	return s
}

type RevokeSecurityGroupRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Priority        *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SourceCidrIp    *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s RevokeSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupRequest) SetVersion(v string) *RevokeSecurityGroupRequest {
	s.Version = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetIpProtocol(v string) *RevokeSecurityGroupRequest {
	s.IpProtocol = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetPortRange(v string) *RevokeSecurityGroupRequest {
	s.PortRange = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSecurityGroupId(v string) *RevokeSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetPolicy(v string) *RevokeSecurityGroupRequest {
	s.Policy = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetPriority(v int32) *RevokeSecurityGroupRequest {
	s.Priority = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSourceCidrIp(v string) *RevokeSecurityGroupRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSourcePortRange(v string) *RevokeSecurityGroupRequest {
	s.SourcePortRange = &v
	return s
}

type RevokeSecurityGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupResponseBody) SetRequestId(v string) *RevokeSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

type RevokeSecurityGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RevokeSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupResponse) SetHeaders(v map[string]*string) *RevokeSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *RevokeSecurityGroupResponse) SetBody(v *RevokeSecurityGroupResponseBody) *RevokeSecurityGroupResponse {
	s.Body = v
	return s
}

type RemovePublicIpsFromEpnInstanceRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	InstanceInfos *string `json:"InstanceInfos,omitempty" xml:"InstanceInfos,omitempty"`
}

func (s RemovePublicIpsFromEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemovePublicIpsFromEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *RemovePublicIpsFromEpnInstanceRequest) SetEPNInstanceId(v string) *RemovePublicIpsFromEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *RemovePublicIpsFromEpnInstanceRequest) SetInstanceInfos(v string) *RemovePublicIpsFromEpnInstanceRequest {
	s.InstanceInfos = &v
	return s
}

type RemovePublicIpsFromEpnInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePublicIpsFromEpnInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemovePublicIpsFromEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePublicIpsFromEpnInstanceResponseBody) SetRequestId(v string) *RemovePublicIpsFromEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RemovePublicIpsFromEpnInstanceResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemovePublicIpsFromEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemovePublicIpsFromEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RemovePublicIpsFromEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *RemovePublicIpsFromEpnInstanceResponse) SetHeaders(v map[string]*string) *RemovePublicIpsFromEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *RemovePublicIpsFromEpnInstanceResponse) SetBody(v *RemovePublicIpsFromEpnInstanceResponseBody) *RemovePublicIpsFromEpnInstanceResponse {
	s.Body = v
	return s
}

type UpgradeApplicationRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Timeout  *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpgradeApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationRequest) SetAppId(v string) *UpgradeApplicationRequest {
	s.AppId = &v
	return s
}

func (s *UpgradeApplicationRequest) SetTemplate(v string) *UpgradeApplicationRequest {
	s.Template = &v
	return s
}

func (s *UpgradeApplicationRequest) SetTimeout(v int32) *UpgradeApplicationRequest {
	s.Timeout = &v
	return s
}

type UpgradeApplicationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationResponseBody) SetRequestId(v string) *UpgradeApplicationResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeApplicationResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationResponse) SetHeaders(v map[string]*string) *UpgradeApplicationResponse {
	s.Headers = v
	return s
}

func (s *UpgradeApplicationResponse) SetBody(v *UpgradeApplicationResponseBody) *UpgradeApplicationResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	Version              *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsRegionId          *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	EnsRegionIds         *string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty"`
	InstanceIds          *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	InstanceName         *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ImageId              *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	OrderByParams        *string `json:"OrderByParams,omitempty" xml:"OrderByParams,omitempty"`
	EnsServiceId         *string `json:"EnsServiceId,omitempty" xml:"EnsServiceId,omitempty"`
	InstanceResourceType *string `json:"InstanceResourceType,omitempty" xml:"InstanceResourceType,omitempty"`
	SearchKey            *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	NetworkId            *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	SecurityGroupId      *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetVersion(v string) *DescribeInstancesRequest {
	s.Version = &v
	return s
}

func (s *DescribeInstancesRequest) SetEnsRegionId(v string) *DescribeInstancesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceId(v string) *DescribeInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesRequest) SetEnsRegionIds(v string) *DescribeInstancesRequest {
	s.EnsRegionIds = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceIds(v string) *DescribeInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceName(v string) *DescribeInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesRequest) SetImageId(v string) *DescribeInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int32) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v string) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetStatus(v string) *DescribeInstancesRequest {
	s.Status = &v
	return s
}

func (s *DescribeInstancesRequest) SetOrderByParams(v string) *DescribeInstancesRequest {
	s.OrderByParams = &v
	return s
}

func (s *DescribeInstancesRequest) SetEnsServiceId(v string) *DescribeInstancesRequest {
	s.EnsServiceId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceResourceType(v string) *DescribeInstancesRequest {
	s.InstanceResourceType = &v
	return s
}

func (s *DescribeInstancesRequest) SetSearchKey(v string) *DescribeInstancesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeInstancesRequest) SetNetworkId(v string) *DescribeInstancesRequest {
	s.NetworkId = &v
	return s
}

func (s *DescribeInstancesRequest) SetVSwitchId(v string) *DescribeInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesRequest) SetSecurityGroupId(v string) *DescribeInstancesRequest {
	s.SecurityGroupId = &v
	return s
}

type DescribeInstancesResponseBody struct {
	Code       *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Instances  *DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetCode(v int32) *DescribeInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageNumber(v int32) *DescribeInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int32) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	Instance []*DescribeInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetInstance(v []*DescribeInstancesResponseBodyInstancesInstance) *DescribeInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

type DescribeInstancesResponseBodyInstancesInstance struct {
	CreationTime            *string                                                           `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Status                  *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	SpecName                *string                                                           `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	InstanceResourceType    *string                                                           `json:"InstanceResourceType,omitempty" xml:"InstanceResourceType,omitempty"`
	HostName                *string                                                           `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceId              *string                                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetMaxBandwidthIn  *int32                                                            `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut *int32                                                            `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	EnsRegionId             *string                                                           `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	Cpu                     *string                                                           `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	ExpiredTime             *string                                                           `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceName            *string                                                           `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Disk                    *int32                                                            `json:"Disk,omitempty" xml:"Disk,omitempty"`
	Memory                  *int32                                                            `json:"Memory,omitempty" xml:"Memory,omitempty"`
	OSName                  *string                                                           `json:"OSName,omitempty" xml:"OSName,omitempty"`
	ImageId                 *string                                                           `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	DataDisk                *DescribeInstancesResponseBodyInstancesInstanceDataDisk           `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Struct"`
	PublicIpAddresses       *DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses  `json:"PublicIpAddresses,omitempty" xml:"PublicIpAddresses,omitempty" type:"Struct"`
	PrivateIpAddresses      *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses `json:"PrivateIpAddresses,omitempty" xml:"PrivateIpAddresses,omitempty" type:"Struct"`
	SecurityGroupIds        *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds   `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
	InnerIpAddress          *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress     `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty" type:"Struct"`
	PublicIpAddress         *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress    `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty" type:"Struct"`
	SystemDisk              *DescribeInstancesResponseBodyInstancesInstanceSystemDisk         `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	NetworkAttributes       *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes  `json:"NetworkAttributes,omitempty" xml:"NetworkAttributes,omitempty" type:"Struct"`
}

func (s DescribeInstancesResponseBodyInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCreationTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetStatus(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetSpecName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.SpecName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceResourceType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceResourceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetHostName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.HostName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInternetMaxBandwidthIn(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInternetMaxBandwidthOut(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetEnsRegionId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCpu(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.Cpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetExpiredTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDisk(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.Disk = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMemory(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.Memory = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetOSName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.OSName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetImageId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ImageId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDataDisk(v *DescribeInstancesResponseBodyInstancesInstanceDataDisk) *DescribeInstancesResponseBodyInstancesInstance {
	s.DataDisk = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetPublicIpAddresses(v *DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses) *DescribeInstancesResponseBodyInstancesInstance {
	s.PublicIpAddresses = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetPrivateIpAddresses(v *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses) *DescribeInstancesResponseBodyInstancesInstance {
	s.PrivateIpAddresses = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetSecurityGroupIds(v *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) *DescribeInstancesResponseBodyInstancesInstance {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInnerIpAddress(v *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) *DescribeInstancesResponseBodyInstancesInstance {
	s.InnerIpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetPublicIpAddress(v *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) *DescribeInstancesResponseBodyInstancesInstance {
	s.PublicIpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetSystemDisk(v *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) *DescribeInstancesResponseBodyInstancesInstance {
	s.SystemDisk = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetNetworkAttributes(v *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) *DescribeInstancesResponseBodyInstancesInstance {
	s.NetworkAttributes = v
	return s
}

type DescribeInstancesResponseBodyInstancesInstanceDataDisk struct {
	DataDisk []*DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceDataDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDisk) SetDataDisk(v []*DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) *DescribeInstancesResponseBodyInstancesInstanceDataDisk {
	s.DataDisk = v
	return s
}

type DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk struct {
	DeviceType *string `json:"device_type,omitempty" xml:"device_type,omitempty"`
	DiskType   *string `json:"disk_type,omitempty" xml:"disk_type,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	DiskName   *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Uuid       *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Storage    *int32  `json:"storage,omitempty" xml:"storage,omitempty"`
	DiskId     *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	Category   *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetDeviceType(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.DeviceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetDiskType(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.DiskType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetSize(v int32) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.Size = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetDiskName(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.DiskName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetUuid(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.Uuid = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetStorage(v int32) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.Storage = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetDiskId(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.DiskId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetCategory(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.Category = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetName(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.Name = &v
	return s
}

type DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses struct {
	PublicIpAddress []*DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses) SetPublicIpAddress(v []*DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) *DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses {
	s.PublicIpAddress = v
	return s
}

type DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress struct {
	Ip      *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	GateWay *string `json:"GateWay,omitempty" xml:"GateWay,omitempty"`
	Isp     *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) SetIp(v string) *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress {
	s.Ip = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) SetGateWay(v string) *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress {
	s.GateWay = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) SetIsp(v string) *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress {
	s.Isp = &v
	return s
}

type DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses struct {
	PrivateIpAddress []*DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses) SetPrivateIpAddress(v []*DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses {
	s.PrivateIpAddress = v
	return s
}

type DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress struct {
	Ip      *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	GateWay *string `json:"GateWay,omitempty" xml:"GateWay,omitempty"`
	Isp     *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) SetIp(v string) *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress {
	s.Ip = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) SetGateWay(v string) *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress {
	s.GateWay = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) SetIsp(v string) *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress {
	s.Isp = &v
	return s
}

type DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds struct {
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) SetSecurityGroupId(v []*string) *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

type DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) SetIpAddress(v []*string) *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress {
	s.IpAddress = v
	return s
}

type DescribeInstancesResponseBodyInstancesInstancePublicIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) SetIpAddress(v []*string) *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress {
	s.IpAddress = v
	return s
}

type DescribeInstancesResponseBodyInstancesInstanceSystemDisk struct {
	DeviceType *string `json:"device_type,omitempty" xml:"device_type,omitempty"`
	DiskType   *string `json:"disk_type,omitempty" xml:"disk_type,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	DiskName   *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Uuid       *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Storage    *int32  `json:"storage,omitempty" xml:"storage,omitempty"`
	DiskId     *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	Category   *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceSystemDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetDeviceType(v string) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.DeviceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetDiskType(v string) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.DiskType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetSize(v int32) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.Size = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetDiskName(v string) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.DiskName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetUuid(v string) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.Uuid = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetStorage(v int32) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.Storage = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetDiskId(v string) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.DiskId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetCategory(v string) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.Category = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetName(v string) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.Name = &v
	return s
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes struct {
	NetworkId        *string                                                                          `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	VSwitchId        *string                                                                          `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	PrivateIpAddress *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Struct"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) SetNetworkId(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes {
	s.NetworkId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) SetVSwitchId(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) SetPrivateIpAddress(v *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress) *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes {
	s.PrivateIpAddress = v
	return s
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress) SetIpAddress(v []*string) *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress {
	s.IpAddress = v
	return s
}

type DescribeInstancesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type RescaleDeviceServiceRequest struct {
	RescaleType      *string `json:"RescaleType,omitempty" xml:"RescaleType,omitempty"`
	RescaleLevel     *string `json:"RescaleLevel,omitempty" xml:"RescaleLevel,omitempty"`
	Timeout          *int64  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ResourceSelector *string `json:"ResourceSelector,omitempty" xml:"ResourceSelector,omitempty"`
}

func (s RescaleDeviceServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RescaleDeviceServiceRequest) GoString() string {
	return s.String()
}

func (s *RescaleDeviceServiceRequest) SetRescaleType(v string) *RescaleDeviceServiceRequest {
	s.RescaleType = &v
	return s
}

func (s *RescaleDeviceServiceRequest) SetRescaleLevel(v string) *RescaleDeviceServiceRequest {
	s.RescaleLevel = &v
	return s
}

func (s *RescaleDeviceServiceRequest) SetTimeout(v int64) *RescaleDeviceServiceRequest {
	s.Timeout = &v
	return s
}

func (s *RescaleDeviceServiceRequest) SetAppId(v string) *RescaleDeviceServiceRequest {
	s.AppId = &v
	return s
}

func (s *RescaleDeviceServiceRequest) SetResourceSelector(v string) *RescaleDeviceServiceRequest {
	s.ResourceSelector = &v
	return s
}

type RescaleDeviceServiceResponseBody struct {
	// Id of the request
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceDetailInfos []*RescaleDeviceServiceResponseBodyResourceDetailInfos `json:"ResourceDetailInfos,omitempty" xml:"ResourceDetailInfos,omitempty" type:"Repeated"`
}

func (s RescaleDeviceServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RescaleDeviceServiceResponseBody) GoString() string {
	return s.String()
}

func (s *RescaleDeviceServiceResponseBody) SetRequestId(v string) *RescaleDeviceServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RescaleDeviceServiceResponseBody) SetResourceDetailInfos(v []*RescaleDeviceServiceResponseBodyResourceDetailInfos) *RescaleDeviceServiceResponseBody {
	s.ResourceDetailInfos = v
	return s
}

type RescaleDeviceServiceResponseBodyResourceDetailInfos struct {
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	ID       *string `json:"ID,omitempty" xml:"ID,omitempty"`
	IP       *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s RescaleDeviceServiceResponseBodyResourceDetailInfos) String() string {
	return tea.Prettify(s)
}

func (s RescaleDeviceServiceResponseBodyResourceDetailInfos) GoString() string {
	return s.String()
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) SetRegionID(v string) *RescaleDeviceServiceResponseBodyResourceDetailInfos {
	s.RegionID = &v
	return s
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) SetID(v string) *RescaleDeviceServiceResponseBodyResourceDetailInfos {
	s.ID = &v
	return s
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) SetIP(v string) *RescaleDeviceServiceResponseBodyResourceDetailInfos {
	s.IP = &v
	return s
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) SetServer(v string) *RescaleDeviceServiceResponseBodyResourceDetailInfos {
	s.Server = &v
	return s
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) SetType(v string) *RescaleDeviceServiceResponseBodyResourceDetailInfos {
	s.Type = &v
	return s
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) SetStatus(v string) *RescaleDeviceServiceResponseBodyResourceDetailInfos {
	s.Status = &v
	return s
}

type RescaleDeviceServiceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RescaleDeviceServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RescaleDeviceServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RescaleDeviceServiceResponse) GoString() string {
	return s.String()
}

func (s *RescaleDeviceServiceResponse) SetHeaders(v map[string]*string) *RescaleDeviceServiceResponse {
	s.Headers = v
	return s
}

func (s *RescaleDeviceServiceResponse) SetBody(v *RescaleDeviceServiceResponseBody) *RescaleDeviceServiceResponse {
	s.Body = v
	return s
}

type CreateEnsServiceRequest struct {
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsServiceId *string `json:"EnsServiceId,omitempty" xml:"EnsServiceId,omitempty"`
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s CreateEnsServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnsServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateEnsServiceRequest) SetVersion(v string) *CreateEnsServiceRequest {
	s.Version = &v
	return s
}

func (s *CreateEnsServiceRequest) SetEnsServiceId(v string) *CreateEnsServiceRequest {
	s.EnsServiceId = &v
	return s
}

func (s *CreateEnsServiceRequest) SetOrderType(v string) *CreateEnsServiceRequest {
	s.OrderType = &v
	return s
}

type CreateEnsServiceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEnsServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEnsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnsServiceResponseBody) SetCode(v int32) *CreateEnsServiceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEnsServiceResponseBody) SetRequestId(v string) *CreateEnsServiceResponseBody {
	s.RequestId = &v
	return s
}

type CreateEnsServiceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEnsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEnsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEnsServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateEnsServiceResponse) SetHeaders(v map[string]*string) *CreateEnsServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateEnsServiceResponse) SetBody(v *CreateEnsServiceResponseBody) *CreateEnsServiceResponse {
	s.Body = v
	return s
}

type ExportBillDetailDataRequest struct {
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s ExportBillDetailDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportBillDetailDataRequest) GoString() string {
	return s.String()
}

func (s *ExportBillDetailDataRequest) SetVersion(v string) *ExportBillDetailDataRequest {
	s.Version = &v
	return s
}

func (s *ExportBillDetailDataRequest) SetStartDate(v string) *ExportBillDetailDataRequest {
	s.StartDate = &v
	return s
}

func (s *ExportBillDetailDataRequest) SetEndDate(v string) *ExportBillDetailDataRequest {
	s.EndDate = &v
	return s
}

type ExportBillDetailDataResponseBody struct {
	FilePath  *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportBillDetailDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportBillDetailDataResponseBody) GoString() string {
	return s.String()
}

func (s *ExportBillDetailDataResponseBody) SetFilePath(v string) *ExportBillDetailDataResponseBody {
	s.FilePath = &v
	return s
}

func (s *ExportBillDetailDataResponseBody) SetRequestId(v string) *ExportBillDetailDataResponseBody {
	s.RequestId = &v
	return s
}

type ExportBillDetailDataResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExportBillDetailDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExportBillDetailDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportBillDetailDataResponse) GoString() string {
	return s.String()
}

func (s *ExportBillDetailDataResponse) SetHeaders(v map[string]*string) *ExportBillDetailDataResponse {
	s.Headers = v
	return s
}

func (s *ExportBillDetailDataResponse) SetBody(v *ExportBillDetailDataResponseBody) *ExportBillDetailDataResponse {
	s.Body = v
	return s
}

type DescribeImageInfosRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	OsType  *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
}

func (s DescribeImageInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageInfosRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosRequest) SetVersion(v string) *DescribeImageInfosRequest {
	s.Version = &v
	return s
}

func (s *DescribeImageInfosRequest) SetOsType(v string) *DescribeImageInfosRequest {
	s.OsType = &v
	return s
}

type DescribeImageInfosResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Images    *DescribeImageInfosResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosResponseBody) SetCode(v int32) *DescribeImageInfosResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageInfosResponseBody) SetImages(v *DescribeImageInfosResponseBodyImages) *DescribeImageInfosResponseBody {
	s.Images = v
	return s
}

func (s *DescribeImageInfosResponseBody) SetRequestId(v string) *DescribeImageInfosResponseBody {
	s.RequestId = &v
	return s
}

type DescribeImageInfosResponseBodyImages struct {
	Image []*DescribeImageInfosResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeImageInfosResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageInfosResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosResponseBodyImages) SetImage(v []*DescribeImageInfosResponseBodyImagesImage) *DescribeImageInfosResponseBodyImages {
	s.Image = v
	return s
}

type DescribeImageInfosResponseBodyImagesImage struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageSize    *string `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	OSName       *string `json:"OSName,omitempty" xml:"OSName,omitempty"`
	OSType       *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
}

func (s DescribeImageInfosResponseBodyImagesImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageInfosResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetDescription(v string) *DescribeImageInfosResponseBodyImagesImage {
	s.Description = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetImageId(v string) *DescribeImageInfosResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetImageSize(v string) *DescribeImageInfosResponseBodyImagesImage {
	s.ImageSize = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetImageVersion(v string) *DescribeImageInfosResponseBodyImagesImage {
	s.ImageVersion = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetOSName(v string) *DescribeImageInfosResponseBodyImagesImage {
	s.OSName = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetOSType(v string) *DescribeImageInfosResponseBodyImagesImage {
	s.OSType = &v
	return s
}

type DescribeImageInfosResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImageInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImageInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageInfosResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosResponse) SetHeaders(v map[string]*string) *DescribeImageInfosResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageInfosResponse) SetBody(v *DescribeImageInfosResponseBody) *DescribeImageInfosResponse {
	s.Body = v
	return s
}

type ModifyNetworkAttributeRequest struct {
	NetworkId   *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	NetworkName *string `json:"NetworkName,omitempty" xml:"NetworkName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyNetworkAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkAttributeRequest) SetNetworkId(v string) *ModifyNetworkAttributeRequest {
	s.NetworkId = &v
	return s
}

func (s *ModifyNetworkAttributeRequest) SetNetworkName(v string) *ModifyNetworkAttributeRequest {
	s.NetworkName = &v
	return s
}

func (s *ModifyNetworkAttributeRequest) SetDescription(v string) *ModifyNetworkAttributeRequest {
	s.Description = &v
	return s
}

type ModifyNetworkAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNetworkAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNetworkAttributeResponseBody) SetRequestId(v string) *ModifyNetworkAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNetworkAttributeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyNetworkAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNetworkAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkAttributeResponse) SetHeaders(v map[string]*string) *ModifyNetworkAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyNetworkAttributeResponse) SetBody(v *ModifyNetworkAttributeResponseBody) *ModifyNetworkAttributeResponse {
	s.Body = v
	return s
}

type DeleteKeyPairsRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
}

func (s DeleteKeyPairsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsRequest) SetVersion(v string) *DeleteKeyPairsRequest {
	s.Version = &v
	return s
}

func (s *DeleteKeyPairsRequest) SetKeyPairName(v string) *DeleteKeyPairsRequest {
	s.KeyPairName = &v
	return s
}

type DeleteKeyPairsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKeyPairsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyPairsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsResponseBody) SetRequestId(v string) *DeleteKeyPairsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteKeyPairsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteKeyPairsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteKeyPairsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyPairsResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsResponse) SetHeaders(v map[string]*string) *DeleteKeyPairsResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeyPairsResponse) SetBody(v *DeleteKeyPairsResponseBody) *DeleteKeyPairsResponse {
	s.Body = v
	return s
}

type AuthorizeSecurityGroupRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Priority        *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SourceCidrIp    *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s AuthorizeSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupRequest) SetVersion(v string) *AuthorizeSecurityGroupRequest {
	s.Version = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetIpProtocol(v string) *AuthorizeSecurityGroupRequest {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPortRange(v string) *AuthorizeSecurityGroupRequest {
	s.PortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSecurityGroupId(v string) *AuthorizeSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPolicy(v string) *AuthorizeSecurityGroupRequest {
	s.Policy = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPriority(v int32) *AuthorizeSecurityGroupRequest {
	s.Priority = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourceCidrIp(v string) *AuthorizeSecurityGroupRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourcePortRange(v string) *AuthorizeSecurityGroupRequest {
	s.SourcePortRange = &v
	return s
}

type AuthorizeSecurityGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupResponseBody) SetRequestId(v string) *AuthorizeSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

type AuthorizeSecurityGroupResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AuthorizeSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthorizeSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupResponse) SetHeaders(v map[string]*string) *AuthorizeSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeSecurityGroupResponse) SetBody(v *AuthorizeSecurityGroupResponseBody) *AuthorizeSecurityGroupResponse {
	s.Body = v
	return s
}

type DeleteSecurityGroupRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DeleteSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupRequest) SetVersion(v string) *DeleteSecurityGroupRequest {
	s.Version = &v
	return s
}

func (s *DeleteSecurityGroupRequest) SetSecurityGroupId(v string) *DeleteSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

type DeleteSecurityGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupResponseBody) SetRequestId(v string) *DeleteSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSecurityGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupResponse) SetHeaders(v map[string]*string) *DeleteSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityGroupResponse) SetBody(v *DeleteSecurityGroupResponseBody) *DeleteSecurityGroupResponse {
	s.Body = v
	return s
}

type ModifyVSwitchAttributeRequest struct {
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyVSwitchAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVSwitchAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVSwitchAttributeRequest) SetVSwitchId(v string) *ModifyVSwitchAttributeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetVSwitchName(v string) *ModifyVSwitchAttributeRequest {
	s.VSwitchName = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetDescription(v string) *ModifyVSwitchAttributeRequest {
	s.Description = &v
	return s
}

type ModifyVSwitchAttributeResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVSwitchAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVSwitchAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVSwitchAttributeResponseBody) SetRequestId(v string) *ModifyVSwitchAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVSwitchAttributeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyVSwitchAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVSwitchAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVSwitchAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyVSwitchAttributeResponse) SetHeaders(v map[string]*string) *ModifyVSwitchAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyVSwitchAttributeResponse) SetBody(v *ModifyVSwitchAttributeResponseBody) *ModifyVSwitchAttributeResponse {
	s.Body = v
	return s
}

type ModifyInstanceAttributeRequest struct {
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ModifyInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) SetVersion(v string) *ModifyInstanceAttributeRequest {
	s.Version = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetPassword(v string) *ModifyInstanceAttributeRequest {
	s.Password = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceName(v string) *ModifyInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

type ModifyInstanceAttributeResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponseBody) SetCode(v int32) *ModifyInstanceAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyInstanceAttributeResponseBody) SetRequestId(v string) *ModifyInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceAttributeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponse) SetHeaders(v map[string]*string) *ModifyInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAttributeResponse) SetBody(v *ModifyInstanceAttributeResponseBody) *ModifyInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeEnsEipAddressesRequest struct {
	// 要查询的EIP实例的ID。  最多支持输入50个EIP实例ID，实例ID之间用逗号（,）分隔。
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// 要查询的EIP的IP地址。  最多支持输入50个EIP的IP地址，IP地址之间用逗号（,）分隔。
	EipAddress             *string `json:"EipAddress,omitempty" xml:"EipAddress,omitempty"`
	AssociatedInstanceId   *string `json:"AssociatedInstanceId,omitempty" xml:"AssociatedInstanceId,omitempty"`
	AssociatedInstanceType *string `json:"AssociatedInstanceType,omitempty" xml:"AssociatedInstanceType,omitempty"`
	// 列表的页码，默认值为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时每页的行数，最大值为100，默认值为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeEnsEipAddressesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsEipAddressesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsEipAddressesRequest) SetAllocationId(v string) *DescribeEnsEipAddressesRequest {
	s.AllocationId = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetEipAddress(v string) *DescribeEnsEipAddressesRequest {
	s.EipAddress = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetAssociatedInstanceId(v string) *DescribeEnsEipAddressesRequest {
	s.AssociatedInstanceId = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetAssociatedInstanceType(v string) *DescribeEnsEipAddressesRequest {
	s.AssociatedInstanceType = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetPageNumber(v int32) *DescribeEnsEipAddressesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetPageSize(v int32) *DescribeEnsEipAddressesRequest {
	s.PageSize = &v
	return s
}

type DescribeEnsEipAddressesResponseBody struct {
	// Id of the request
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EipAddresses *DescribeEnsEipAddressesResponseBodyEipAddresses `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty" type:"Struct"`
	PageNumber   *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount   *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEnsEipAddressesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsEipAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsEipAddressesResponseBody) SetRequestId(v string) *DescribeEnsEipAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBody) SetEipAddresses(v *DescribeEnsEipAddressesResponseBodyEipAddresses) *DescribeEnsEipAddressesResponseBody {
	s.EipAddresses = v
	return s
}

func (s *DescribeEnsEipAddressesResponseBody) SetPageNumber(v int32) *DescribeEnsEipAddressesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBody) SetPageSize(v int32) *DescribeEnsEipAddressesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBody) SetTotalCount(v int32) *DescribeEnsEipAddressesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEnsEipAddressesResponseBodyEipAddresses struct {
	EipAddress []*DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress `json:"EipAddress,omitempty" xml:"EipAddress,omitempty" type:"Repeated"`
}

func (s DescribeEnsEipAddressesResponseBodyEipAddresses) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsEipAddressesResponseBodyEipAddresses) GoString() string {
	return s.String()
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddresses) SetEipAddress(v []*DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) *DescribeEnsEipAddressesResponseBodyEipAddresses {
	s.EipAddress = v
	return s
}

type DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress struct {
	AllocationId       *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	AllocationTime     *string `json:"AllocationTime,omitempty" xml:"AllocationTime,omitempty"`
	Bandwidth          *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ChargeType         *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType       *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	IpAddress          *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	EnsRegionId        *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Isp                *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GoString() string {
	return s.String()
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetAllocationId(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.AllocationId = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetAllocationTime(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.AllocationTime = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetBandwidth(v int32) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.Bandwidth = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetChargeType(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.ChargeType = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetDescription(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.Description = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetInstanceId(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.InstanceId = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetInstanceType(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.InstanceType = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetInternetChargeType(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetIpAddress(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.IpAddress = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetName(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.Name = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetEnsRegionId(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetStatus(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.Status = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetIsp(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.Isp = &v
	return s
}

type DescribeEnsEipAddressesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEnsEipAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEnsEipAddressesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsEipAddressesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsEipAddressesResponse) SetHeaders(v map[string]*string) *DescribeEnsEipAddressesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsEipAddressesResponse) SetBody(v *DescribeEnsEipAddressesResponseBody) *DescribeEnsEipAddressesResponse {
	s.Body = v
	return s
}

type RescaleApplicationRequest struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RescaleType      *string `json:"RescaleType,omitempty" xml:"RescaleType,omitempty"`
	RescaleLevel     *string `json:"RescaleLevel,omitempty" xml:"RescaleLevel,omitempty"`
	ResourceSelector *string `json:"ResourceSelector,omitempty" xml:"ResourceSelector,omitempty"`
	ToAppVersion     *string `json:"ToAppVersion,omitempty" xml:"ToAppVersion,omitempty"`
	Timeout          *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s RescaleApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s RescaleApplicationRequest) GoString() string {
	return s.String()
}

func (s *RescaleApplicationRequest) SetAppId(v string) *RescaleApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RescaleApplicationRequest) SetRescaleType(v string) *RescaleApplicationRequest {
	s.RescaleType = &v
	return s
}

func (s *RescaleApplicationRequest) SetRescaleLevel(v string) *RescaleApplicationRequest {
	s.RescaleLevel = &v
	return s
}

func (s *RescaleApplicationRequest) SetResourceSelector(v string) *RescaleApplicationRequest {
	s.ResourceSelector = &v
	return s
}

func (s *RescaleApplicationRequest) SetToAppVersion(v string) *RescaleApplicationRequest {
	s.ToAppVersion = &v
	return s
}

func (s *RescaleApplicationRequest) SetTimeout(v int32) *RescaleApplicationRequest {
	s.Timeout = &v
	return s
}

type RescaleApplicationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RescaleApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RescaleApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RescaleApplicationResponseBody) SetRequestId(v string) *RescaleApplicationResponseBody {
	s.RequestId = &v
	return s
}

type RescaleApplicationResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RescaleApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RescaleApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s RescaleApplicationResponse) GoString() string {
	return s.String()
}

func (s *RescaleApplicationResponse) SetHeaders(v map[string]*string) *RescaleApplicationResponse {
	s.Headers = v
	return s
}

func (s *RescaleApplicationResponse) SetBody(v *RescaleApplicationResponseBody) *RescaleApplicationResponse {
	s.Body = v
	return s
}

type DescribeSecurityGroupsRequest struct {
	Version           *string `json:"Version,omitempty" xml:"Version,omitempty"`
	SecurityGroupId   *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s DescribeSecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsRequest) SetVersion(v string) *DescribeSecurityGroupsRequest {
	s.Version = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetSecurityGroupId(v string) *DescribeSecurityGroupsRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetPageNumber(v int32) *DescribeSecurityGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetPageSize(v int32) *DescribeSecurityGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetSecurityGroupName(v string) *DescribeSecurityGroupsRequest {
	s.SecurityGroupName = &v
	return s
}

type DescribeSecurityGroupsResponseBody struct {
	PageSize       *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber     *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount     *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SecurityGroups *DescribeSecurityGroupsResponseBodySecurityGroups `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Struct"`
}

func (s DescribeSecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBody) SetPageSize(v int32) *DescribeSecurityGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetRequestId(v string) *DescribeSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetPageNumber(v int32) *DescribeSecurityGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetTotalCount(v int32) *DescribeSecurityGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetSecurityGroups(v *DescribeSecurityGroupsResponseBodySecurityGroups) *DescribeSecurityGroupsResponseBody {
	s.SecurityGroups = v
	return s
}

type DescribeSecurityGroupsResponseBodySecurityGroups struct {
	SecurityGroup []*DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupsResponseBodySecurityGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBodySecurityGroups) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroups) SetSecurityGroup(v []*DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) *DescribeSecurityGroupsResponseBodySecurityGroups {
	s.SecurityGroup = v
	return s
}

type DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup struct {
	CreationTime      *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	InstanceCount     *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityGroupId   *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetCreationTime(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.CreationTime = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetInstanceCount(v int32) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.InstanceCount = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetDescription(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.Description = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetSecurityGroupId(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetSecurityGroupName(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.SecurityGroupName = &v
	return s
}

type DescribeSecurityGroupsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponse) SetHeaders(v map[string]*string) *DescribeSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityGroupsResponse) SetBody(v *DescribeSecurityGroupsResponseBody) *DescribeSecurityGroupsResponse {
	s.Body = v
	return s
}

type DescribeUserBandWidthDataRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Period      *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Isp         *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeUserBandWidthDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBandWidthDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserBandWidthDataRequest) SetVersion(v string) *DescribeUserBandWidthDataRequest {
	s.Version = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetEnsRegionId(v string) *DescribeUserBandWidthDataRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetInstanceId(v string) *DescribeUserBandWidthDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetStartTime(v string) *DescribeUserBandWidthDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetEndTime(v string) *DescribeUserBandWidthDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetPeriod(v string) *DescribeUserBandWidthDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetIsp(v string) *DescribeUserBandWidthDataRequest {
	s.Isp = &v
	return s
}

type DescribeUserBandWidthDataResponseBody struct {
	Code        *int32                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	MonitorData *DescribeUserBandWidthDataResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Struct"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserBandWidthDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBandWidthDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserBandWidthDataResponseBody) SetCode(v int32) *DescribeUserBandWidthDataResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBody) SetMonitorData(v *DescribeUserBandWidthDataResponseBodyMonitorData) *DescribeUserBandWidthDataResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeUserBandWidthDataResponseBody) SetRequestId(v string) *DescribeUserBandWidthDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserBandWidthDataResponseBodyMonitorData struct {
	BandWidthMonitorData []*DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData `json:"BandWidthMonitorData,omitempty" xml:"BandWidthMonitorData,omitempty" type:"Repeated"`
	MaxDownBandWidth     *string                                                                 `json:"MaxDownBandWidth,omitempty" xml:"MaxDownBandWidth,omitempty"`
	MaxUpBandWidth       *string                                                                 `json:"MaxUpBandWidth,omitempty" xml:"MaxUpBandWidth,omitempty"`
}

func (s DescribeUserBandWidthDataResponseBodyMonitorData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBandWidthDataResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorData) SetBandWidthMonitorData(v []*DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) *DescribeUserBandWidthDataResponseBodyMonitorData {
	s.BandWidthMonitorData = v
	return s
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorData) SetMaxDownBandWidth(v string) *DescribeUserBandWidthDataResponseBodyMonitorData {
	s.MaxDownBandWidth = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorData) SetMaxUpBandWidth(v string) *DescribeUserBandWidthDataResponseBodyMonitorData {
	s.MaxUpBandWidth = &v
	return s
}

type DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData struct {
	DownBandWidth *int32  `json:"DownBandWidth,omitempty" xml:"DownBandWidth,omitempty"`
	InternetRX    *int32  `json:"InternetRX,omitempty" xml:"InternetRX,omitempty"`
	InternetTX    *int32  `json:"InternetTX,omitempty" xml:"InternetTX,omitempty"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	UpBandWidth   *int32  `json:"UpBandWidth,omitempty" xml:"UpBandWidth,omitempty"`
}

func (s DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetDownBandWidth(v int32) *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.DownBandWidth = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetInternetRX(v int32) *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.InternetRX = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetInternetTX(v int32) *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.InternetTX = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetTimeStamp(v string) *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetUpBandWidth(v int32) *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.UpBandWidth = &v
	return s
}

type DescribeUserBandWidthDataResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserBandWidthDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserBandWidthDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBandWidthDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserBandWidthDataResponse) SetHeaders(v map[string]*string) *DescribeUserBandWidthDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserBandWidthDataResponse) SetBody(v *DescribeUserBandWidthDataResponseBody) *DescribeUserBandWidthDataResponse {
	s.Body = v
	return s
}

type StartEpnInstanceRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
}

func (s StartEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartEpnInstanceRequest) SetEPNInstanceId(v string) *StartEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

type StartEpnInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartEpnInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartEpnInstanceResponseBody) SetRequestId(v string) *StartEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartEpnInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartEpnInstanceResponse) SetHeaders(v map[string]*string) *StartEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartEpnInstanceResponse) SetBody(v *StartEpnInstanceResponseBody) *StartEpnInstanceResponse {
	s.Body = v
	return s
}

type DescribeEpnBandWidthDataRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsRegionId     *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Period          *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Isp             *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	NetworkingModel *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
	EPNInstanceId   *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
}

func (s DescribeEpnBandWidthDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnBandWidthDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandWidthDataRequest) SetVersion(v string) *DescribeEpnBandWidthDataRequest {
	s.Version = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetEnsRegionId(v string) *DescribeEpnBandWidthDataRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetInstanceId(v string) *DescribeEpnBandWidthDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetStartTime(v string) *DescribeEpnBandWidthDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetEndTime(v string) *DescribeEpnBandWidthDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetPeriod(v string) *DescribeEpnBandWidthDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetIsp(v string) *DescribeEpnBandWidthDataRequest {
	s.Isp = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetNetworkingModel(v string) *DescribeEpnBandWidthDataRequest {
	s.NetworkingModel = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetEPNInstanceId(v string) *DescribeEpnBandWidthDataRequest {
	s.EPNInstanceId = &v
	return s
}

type DescribeEpnBandWidthDataResponseBody struct {
	MonitorData *DescribeEpnBandWidthDataResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Struct"`
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEpnBandWidthDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnBandWidthDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandWidthDataResponseBody) SetMonitorData(v *DescribeEpnBandWidthDataResponseBodyMonitorData) *DescribeEpnBandWidthDataResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBody) SetRequestId(v string) *DescribeEpnBandWidthDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEpnBandWidthDataResponseBodyMonitorData struct {
	BandWidthMonitorData []*DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData `json:"BandWidthMonitorData,omitempty" xml:"BandWidthMonitorData,omitempty" type:"Repeated"`
	MaxDownBandWidth     *int64                                                                 `json:"MaxDownBandWidth,omitempty" xml:"MaxDownBandWidth,omitempty"`
	MaxUpBandWidth       *int64                                                                 `json:"MaxUpBandWidth,omitempty" xml:"MaxUpBandWidth,omitempty"`
}

func (s DescribeEpnBandWidthDataResponseBodyMonitorData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnBandWidthDataResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorData) SetBandWidthMonitorData(v []*DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) *DescribeEpnBandWidthDataResponseBodyMonitorData {
	s.BandWidthMonitorData = v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorData) SetMaxDownBandWidth(v int64) *DescribeEpnBandWidthDataResponseBodyMonitorData {
	s.MaxDownBandWidth = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorData) SetMaxUpBandWidth(v int64) *DescribeEpnBandWidthDataResponseBodyMonitorData {
	s.MaxUpBandWidth = &v
	return s
}

type DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData struct {
	DownBandWidth *int64  `json:"DownBandWidth,omitempty" xml:"DownBandWidth,omitempty"`
	InternetRX    *int64  `json:"InternetRX,omitempty" xml:"InternetRX,omitempty"`
	InternetTX    *int64  `json:"InternetTX,omitempty" xml:"InternetTX,omitempty"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	UpBandWidth   *int64  `json:"UpBandWidth,omitempty" xml:"UpBandWidth,omitempty"`
}

func (s DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetDownBandWidth(v int64) *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.DownBandWidth = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetInternetRX(v int64) *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.InternetRX = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetInternetTX(v int64) *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.InternetTX = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetTimeStamp(v string) *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetUpBandWidth(v int64) *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.UpBandWidth = &v
	return s
}

type DescribeEpnBandWidthDataResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEpnBandWidthDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEpnBandWidthDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnBandWidthDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandWidthDataResponse) SetHeaders(v map[string]*string) *DescribeEpnBandWidthDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeEpnBandWidthDataResponse) SetBody(v *DescribeEpnBandWidthDataResponseBody) *DescribeEpnBandWidthDataResponse {
	s.Body = v
	return s
}

type DescribeMeasurementDataRequest struct {
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s DescribeMeasurementDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataRequest) SetVersion(v string) *DescribeMeasurementDataRequest {
	s.Version = &v
	return s
}

func (s *DescribeMeasurementDataRequest) SetStartDate(v string) *DescribeMeasurementDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeMeasurementDataRequest) SetEndDate(v string) *DescribeMeasurementDataRequest {
	s.EndDate = &v
	return s
}

type DescribeMeasurementDataResponseBody struct {
	MeasurementDatas *DescribeMeasurementDataResponseBodyMeasurementDatas `json:"MeasurementDatas,omitempty" xml:"MeasurementDatas,omitempty" type:"Struct"`
	RequestId        *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeasurementDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBody) SetMeasurementDatas(v *DescribeMeasurementDataResponseBodyMeasurementDatas) *DescribeMeasurementDataResponseBody {
	s.MeasurementDatas = v
	return s
}

func (s *DescribeMeasurementDataResponseBody) SetRequestId(v string) *DescribeMeasurementDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMeasurementDataResponseBodyMeasurementDatas struct {
	MeasurementData []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData `json:"MeasurementData,omitempty" xml:"MeasurementData,omitempty" type:"Repeated"`
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatas) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatas) SetMeasurementData(v []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) *DescribeMeasurementDataResponseBodyMeasurementDatas {
	s.MeasurementData = v
	return s
}

type DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData struct {
	BandWidthFeeDatas      *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas      `json:"BandWidthFeeDatas,omitempty" xml:"BandWidthFeeDatas,omitempty" type:"Struct"`
	ChargeModel            *string                                                                                   `json:"ChargeModel,omitempty" xml:"ChargeModel,omitempty"`
	CostCycle              *string                                                                                   `json:"CostCycle,omitempty" xml:"CostCycle,omitempty"`
	CostEndTime            *string                                                                                   `json:"CostEndTime,omitempty" xml:"CostEndTime,omitempty"`
	CostStartTime          *string                                                                                   `json:"CostStartTime,omitempty" xml:"CostStartTime,omitempty"`
	ResourceFeeData        *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData        `json:"ResourceFeeData,omitempty" xml:"ResourceFeeData,omitempty" type:"Struct"`
	ResourceFeeDataDetails *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails `json:"ResourceFeeDataDetails,omitempty" xml:"ResourceFeeDataDetails,omitempty" type:"Struct"`
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetBandWidthFeeDatas(v *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.BandWidthFeeDatas = v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetChargeModel(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.ChargeModel = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetCostCycle(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.CostCycle = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetCostEndTime(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.CostEndTime = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetCostStartTime(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.CostStartTime = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetResourceFeeData(v *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.ResourceFeeData = v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetResourceFeeDataDetails(v *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.ResourceFeeDataDetails = v
	return s
}

type DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas struct {
	BandWidthFeeData []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData `json:"BandWidthFeeData,omitempty" xml:"BandWidthFeeData,omitempty" type:"Repeated"`
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) SetBandWidthFeeData(v []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas {
	s.BandWidthFeeData = v
	return s
}

type DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData struct {
	CostCode *string `json:"CostCode,omitempty" xml:"CostCode,omitempty"`
	CostName *string `json:"CostName,omitempty" xml:"CostName,omitempty"`
	CostVal  *int32  `json:"CostVal,omitempty" xml:"CostVal,omitempty"`
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostCode(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostCode = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostName(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostName = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostVal(v int32) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostVal = &v
	return s
}

type DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData struct {
	Memory  *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	Vcpu    *int32 `json:"Vcpu,omitempty" xml:"Vcpu,omitempty"`
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) SetMemory(v int32) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData {
	s.Memory = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) SetStorage(v int32) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData {
	s.Storage = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) SetVcpu(v int32) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData {
	s.Vcpu = &v
	return s
}

type DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails struct {
	ResourceFeeDataDetail []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail `json:"ResourceFeeDataDetail,omitempty" xml:"ResourceFeeDataDetail,omitempty" type:"Repeated"`
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails) SetResourceFeeDataDetail(v []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails {
	s.ResourceFeeDataDetail = v
	return s
}

type DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail struct {
	CostCode     *string `json:"CostCode,omitempty" xml:"CostCode,omitempty"`
	CostName     *string `json:"CostName,omitempty" xml:"CostName,omitempty"`
	CostVal      *int32  `json:"CostVal,omitempty" xml:"CostVal,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) SetCostCode(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail {
	s.CostCode = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) SetCostName(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail {
	s.CostName = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) SetCostVal(v int32) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail {
	s.CostVal = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) SetResourceType(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail {
	s.ResourceType = &v
	return s
}

type DescribeMeasurementDataResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMeasurementDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMeasurementDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponse) SetHeaders(v map[string]*string) *DescribeMeasurementDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeasurementDataResponse) SetBody(v *DescribeMeasurementDataResponseBody) *DescribeMeasurementDataResponse {
	s.Body = v
	return s
}

type CreateVSwitchRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NetworkId   *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
}

func (s CreateVSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVSwitchRequest) GoString() string {
	return s.String()
}

func (s *CreateVSwitchRequest) SetVersion(v string) *CreateVSwitchRequest {
	s.Version = &v
	return s
}

func (s *CreateVSwitchRequest) SetEnsRegionId(v string) *CreateVSwitchRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateVSwitchRequest) SetCidrBlock(v string) *CreateVSwitchRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateVSwitchRequest) SetVSwitchName(v string) *CreateVSwitchRequest {
	s.VSwitchName = &v
	return s
}

func (s *CreateVSwitchRequest) SetDescription(v string) *CreateVSwitchRequest {
	s.Description = &v
	return s
}

func (s *CreateVSwitchRequest) SetNetworkId(v string) *CreateVSwitchRequest {
	s.NetworkId = &v
	return s
}

type CreateVSwitchResponseBody struct {
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVSwitchResponseBody) SetVSwitchId(v string) *CreateVSwitchResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *CreateVSwitchResponseBody) SetRequestId(v string) *CreateVSwitchResponseBody {
	s.RequestId = &v
	return s
}

type CreateVSwitchResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVSwitchResponse) GoString() string {
	return s.String()
}

func (s *CreateVSwitchResponse) SetHeaders(v map[string]*string) *CreateVSwitchResponse {
	s.Headers = v
	return s
}

func (s *CreateVSwitchResponse) SetBody(v *CreateVSwitchResponseBody) *CreateVSwitchResponse {
	s.Body = v
	return s
}

type CreateEpnInstanceRequest struct {
	EPNInstanceType         *string `json:"EPNInstanceType,omitempty" xml:"EPNInstanceType,omitempty"`
	EPNInstanceName         *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	InternetChargeType      *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	NetworkingModel         *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
	InternetMaxBandwidthOut *int32  `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
}

func (s CreateEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateEpnInstanceRequest) SetEPNInstanceType(v string) *CreateEpnInstanceRequest {
	s.EPNInstanceType = &v
	return s
}

func (s *CreateEpnInstanceRequest) SetEPNInstanceName(v string) *CreateEpnInstanceRequest {
	s.EPNInstanceName = &v
	return s
}

func (s *CreateEpnInstanceRequest) SetInternetChargeType(v string) *CreateEpnInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateEpnInstanceRequest) SetNetworkingModel(v string) *CreateEpnInstanceRequest {
	s.NetworkingModel = &v
	return s
}

func (s *CreateEpnInstanceRequest) SetInternetMaxBandwidthOut(v int32) *CreateEpnInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

type CreateEpnInstanceResponseBody struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEpnInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEpnInstanceResponseBody) SetEPNInstanceId(v string) *CreateEpnInstanceResponseBody {
	s.EPNInstanceId = &v
	return s
}

func (s *CreateEpnInstanceResponseBody) SetRequestId(v string) *CreateEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateEpnInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateEpnInstanceResponse) SetHeaders(v map[string]*string) *CreateEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateEpnInstanceResponse) SetBody(v *CreateEpnInstanceResponseBody) *CreateEpnInstanceResponse {
	s.Body = v
	return s
}

type ModifySecurityGroupAttributeRequest struct {
	SecurityGroupId   *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifySecurityGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupAttributeRequest) SetSecurityGroupId(v string) *ModifySecurityGroupAttributeRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetSecurityGroupName(v string) *ModifySecurityGroupAttributeRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetDescription(v string) *ModifySecurityGroupAttributeRequest {
	s.Description = &v
	return s
}

type ModifySecurityGroupAttributeResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupAttributeResponseBody) SetRequestId(v string) *ModifySecurityGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifySecurityGroupAttributeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySecurityGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySecurityGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupAttributeResponse) SetHeaders(v map[string]*string) *ModifySecurityGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityGroupAttributeResponse) SetBody(v *ModifySecurityGroupAttributeResponseBody) *ModifySecurityGroupAttributeResponse {
	s.Body = v
	return s
}

type DescribeAvailableResourceRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) SetVersion(v string) *DescribeAvailableResourceRequest {
	s.Version = &v
	return s
}

type DescribeAvailableResourceResponseBody struct {
	Code             *int32                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Images           *DescribeAvailableResourceResponseBodyImages           `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupportResources *DescribeAvailableResourceResponseBodySupportResources `json:"SupportResources,omitempty" xml:"SupportResources,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) SetCode(v int32) *DescribeAvailableResourceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetImages(v *DescribeAvailableResourceResponseBodyImages) *DescribeAvailableResourceResponseBody {
	s.Images = v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetSupportResources(v *DescribeAvailableResourceResponseBodySupportResources) *DescribeAvailableResourceResponseBody {
	s.SupportResources = v
	return s
}

type DescribeAvailableResourceResponseBodyImages struct {
	Image []*DescribeAvailableResourceResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyImages) SetImage(v []*DescribeAvailableResourceResponseBodyImagesImage) *DescribeAvailableResourceResponseBodyImages {
	s.Image = v
	return s
}

type DescribeAvailableResourceResponseBodyImagesImage struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyImagesImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyImagesImage) SetImageId(v string) *DescribeAvailableResourceResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyImagesImage) SetImageName(v string) *DescribeAvailableResourceResponseBodyImagesImage {
	s.ImageName = &v
	return s
}

type DescribeAvailableResourceResponseBodySupportResources struct {
	SupportResource []*DescribeAvailableResourceResponseBodySupportResourcesSupportResource `json:"SupportResource,omitempty" xml:"SupportResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodySupportResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportResources) SetSupportResource(v []*DescribeAvailableResourceResponseBodySupportResourcesSupportResource) *DescribeAvailableResourceResponseBodySupportResources {
	s.SupportResource = v
	return s
}

type DescribeAvailableResourceResponseBodySupportResourcesSupportResource struct {
	DataDiskSize          *string `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	EnsRegionId           *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	InstanceSpec          *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	SupportResourcesCount *string `json:"SupportResourcesCount,omitempty" xml:"SupportResourcesCount,omitempty"`
	SystemDiskSize        *string `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeAvailableResourceResponseBodySupportResourcesSupportResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportResourcesSupportResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) SetDataDiskSize(v string) *DescribeAvailableResourceResponseBodySupportResourcesSupportResource {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) SetEnsRegionId(v string) *DescribeAvailableResourceResponseBodySupportResourcesSupportResource {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) SetInstanceSpec(v string) *DescribeAvailableResourceResponseBodySupportResourcesSupportResource {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) SetSupportResourcesCount(v string) *DescribeAvailableResourceResponseBodySupportResourcesSupportResource {
	s.SupportResourcesCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportResourcesSupportResource) SetSystemDiskSize(v string) *DescribeAvailableResourceResponseBodySupportResourcesSupportResource {
	s.SystemDiskSize = &v
	return s
}

type DescribeAvailableResourceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponse) SetHeaders(v map[string]*string) *DescribeAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableResourceResponse) SetBody(v *DescribeAvailableResourceResponseBody) *DescribeAvailableResourceResponse {
	s.Body = v
	return s
}

type DeleteNetworkRequest struct {
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
}

func (s DeleteNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkRequest) SetNetworkId(v string) *DeleteNetworkRequest {
	s.NetworkId = &v
	return s
}

type DeleteNetworkResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkResponseBody) SetRequestId(v string) *DeleteNetworkResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNetworkResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkResponse) SetHeaders(v map[string]*string) *DeleteNetworkResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkResponse) SetBody(v *DeleteNetworkResponseBody) *DeleteNetworkResponse {
	s.Body = v
	return s
}

type ReleaseEipAddressRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	Eips        *string `json:"Eips,omitempty" xml:"Eips,omitempty"`
}

func (s ReleaseEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseEipAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleaseEipAddressRequest) SetVersion(v string) *ReleaseEipAddressRequest {
	s.Version = &v
	return s
}

func (s *ReleaseEipAddressRequest) SetEnsRegionId(v string) *ReleaseEipAddressRequest {
	s.EnsRegionId = &v
	return s
}

func (s *ReleaseEipAddressRequest) SetEips(v string) *ReleaseEipAddressRequest {
	s.Eips = &v
	return s
}

type ReleaseEipAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseEipAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseEipAddressResponseBody) SetRequestId(v string) *ReleaseEipAddressResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseEipAddressResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseEipAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleaseEipAddressResponse) SetHeaders(v map[string]*string) *ReleaseEipAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleaseEipAddressResponse) SetBody(v *ReleaseEipAddressResponseBody) *ReleaseEipAddressResponse {
	s.Body = v
	return s
}

type DescribeVSwitchesRequest struct {
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsRegionId   *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	VSwitchId     *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName   *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	NetworkId     *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OrderByParams *string `json:"OrderByParams,omitempty" xml:"OrderByParams,omitempty"`
}

func (s DescribeVSwitchesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesRequest) SetVersion(v string) *DescribeVSwitchesRequest {
	s.Version = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetEnsRegionId(v string) *DescribeVSwitchesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVSwitchId(v string) *DescribeVSwitchesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVSwitchName(v string) *DescribeVSwitchesRequest {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetNetworkId(v string) *DescribeVSwitchesRequest {
	s.NetworkId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetPageNumber(v int32) *DescribeVSwitchesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetPageSize(v int32) *DescribeVSwitchesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetOrderByParams(v string) *DescribeVSwitchesRequest {
	s.OrderByParams = &v
	return s
}

type DescribeVSwitchesResponseBody struct {
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VSwitches  *DescribeVSwitchesResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Struct"`
}

func (s DescribeVSwitchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBody) SetPageSize(v int32) *DescribeVSwitchesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetRequestId(v string) *DescribeVSwitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetPageNumber(v int32) *DescribeVSwitchesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetTotalCount(v int32) *DescribeVSwitchesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetVSwitches(v *DescribeVSwitchesResponseBodyVSwitches) *DescribeVSwitchesResponseBody {
	s.VSwitches = v
	return s
}

type DescribeVSwitchesResponseBodyVSwitches struct {
	VSwitch []*DescribeVSwitchesResponseBodyVSwitchesVSwitch `json:"VSwitch,omitempty" xml:"VSwitch,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchesResponseBodyVSwitches) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetVSwitch(v []*DescribeVSwitchesResponseBodyVSwitchesVSwitch) *DescribeVSwitchesResponseBodyVSwitches {
	s.VSwitch = v
	return s
}

type DescribeVSwitchesResponseBodyVSwitchesVSwitch struct {
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	FreeIpCount *int64  `json:"FreeIpCount,omitempty" xml:"FreeIpCount,omitempty"`
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NetworkId   *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitch) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetEnsRegionId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetStatus(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetFreeIpCount(v int64) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.FreeIpCount = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetVSwitchId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetCidrBlock(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetDescription(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Description = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetNetworkId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.NetworkId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetCreatedTime(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.CreatedTime = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetVSwitchName(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.VSwitchName = &v
	return s
}

type DescribeVSwitchesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVSwitchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVSwitchesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponse) SetHeaders(v map[string]*string) *DescribeVSwitchesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVSwitchesResponse) SetBody(v *DescribeVSwitchesResponseBody) *DescribeVSwitchesResponse {
	s.Body = v
	return s
}

type ExportImageRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	OSSBucket   *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	OSSRegionId *string `json:"OSSRegionId,omitempty" xml:"OSSRegionId,omitempty"`
	OSSPrefix   *string `json:"OSSPrefix,omitempty" xml:"OSSPrefix,omitempty"`
	RoleName    *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ExportImageRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportImageRequest) GoString() string {
	return s.String()
}

func (s *ExportImageRequest) SetVersion(v string) *ExportImageRequest {
	s.Version = &v
	return s
}

func (s *ExportImageRequest) SetImageId(v string) *ExportImageRequest {
	s.ImageId = &v
	return s
}

func (s *ExportImageRequest) SetOSSBucket(v string) *ExportImageRequest {
	s.OSSBucket = &v
	return s
}

func (s *ExportImageRequest) SetOSSRegionId(v string) *ExportImageRequest {
	s.OSSRegionId = &v
	return s
}

func (s *ExportImageRequest) SetOSSPrefix(v string) *ExportImageRequest {
	s.OSSPrefix = &v
	return s
}

func (s *ExportImageRequest) SetRoleName(v string) *ExportImageRequest {
	s.RoleName = &v
	return s
}

type ExportImageResponseBody struct {
	ExportedImageURL *string `json:"ExportedImageURL,omitempty" xml:"ExportedImageURL,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportImageResponseBody) GoString() string {
	return s.String()
}

func (s *ExportImageResponseBody) SetExportedImageURL(v string) *ExportImageResponseBody {
	s.ExportedImageURL = &v
	return s
}

func (s *ExportImageResponseBody) SetRequestId(v string) *ExportImageResponseBody {
	s.RequestId = &v
	return s
}

type ExportImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExportImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExportImageResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportImageResponse) GoString() string {
	return s.String()
}

func (s *ExportImageResponse) SetHeaders(v map[string]*string) *ExportImageResponse {
	s.Headers = v
	return s
}

func (s *ExportImageResponse) SetBody(v *ExportImageResponseBody) *ExportImageResponse {
	s.Body = v
	return s
}

type DescribeInstanceSpecRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecRequest) SetVersion(v string) *DescribeInstanceSpecRequest {
	s.Version = &v
	return s
}

type DescribeInstanceSpecResponseBody struct {
	BandwidthLimit    *int32                                         `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	Code              *int32                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	DataDiskMaxSize   *int32                                         `json:"DataDiskMaxSize,omitempty" xml:"DataDiskMaxSize,omitempty"`
	DataDiskMinSize   *int32                                         `json:"DataDiskMinSize,omitempty" xml:"DataDiskMinSize,omitempty"`
	InstanceSpecs     *DescribeInstanceSpecResponseBodyInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" type:"Struct"`
	RequestId         *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SystemDiskMaxSize *int32                                         `json:"SystemDiskMaxSize,omitempty" xml:"SystemDiskMaxSize,omitempty"`
}

func (s DescribeInstanceSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecResponseBody) SetBandwidthLimit(v int32) *DescribeInstanceSpecResponseBody {
	s.BandwidthLimit = &v
	return s
}

func (s *DescribeInstanceSpecResponseBody) SetCode(v int32) *DescribeInstanceSpecResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceSpecResponseBody) SetDataDiskMaxSize(v int32) *DescribeInstanceSpecResponseBody {
	s.DataDiskMaxSize = &v
	return s
}

func (s *DescribeInstanceSpecResponseBody) SetDataDiskMinSize(v int32) *DescribeInstanceSpecResponseBody {
	s.DataDiskMinSize = &v
	return s
}

func (s *DescribeInstanceSpecResponseBody) SetInstanceSpecs(v *DescribeInstanceSpecResponseBodyInstanceSpecs) *DescribeInstanceSpecResponseBody {
	s.InstanceSpecs = v
	return s
}

func (s *DescribeInstanceSpecResponseBody) SetRequestId(v string) *DescribeInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSpecResponseBody) SetSystemDiskMaxSize(v int32) *DescribeInstanceSpecResponseBody {
	s.SystemDiskMaxSize = &v
	return s
}

type DescribeInstanceSpecResponseBodyInstanceSpecs struct {
	InstanceSpec []*DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSpecResponseBodyInstanceSpecs) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecResponseBodyInstanceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecs) SetInstanceSpec(v []*DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) *DescribeInstanceSpecResponseBodyInstanceSpecs {
	s.InstanceSpec = v
	return s
}

type DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec struct {
	Core         *string `json:"Core,omitempty" xml:"Core,omitempty"`
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Memory       *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) SetCore(v string) *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec {
	s.Core = &v
	return s
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) SetDisplayName(v string) *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec {
	s.DisplayName = &v
	return s
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) SetInstanceType(v string) *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) SetMemory(v string) *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec {
	s.Memory = &v
	return s
}

type DescribeInstanceSpecResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecResponse) SetHeaders(v map[string]*string) *DescribeInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSpecResponse) SetBody(v *DescribeInstanceSpecResponseBody) *DescribeInstanceSpecResponse {
	s.Body = v
	return s
}

type StartInstanceRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) SetVersion(v string) *StartInstanceRequest {
	s.Version = &v
	return s
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

type StartInstanceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetCode(v int32) *StartInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetHeaders(v map[string]*string) *StartInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type DescribeInstanceTypesRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeInstanceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesRequest) SetVersion(v string) *DescribeInstanceTypesRequest {
	s.Version = &v
	return s
}

type DescribeInstanceTypesResponseBody struct {
	Code          *int32                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceTypes *DescribeInstanceTypesResponseBodyInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBody) SetCode(v int32) *DescribeInstanceTypesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceTypesResponseBody) SetInstanceTypes(v *DescribeInstanceTypesResponseBodyInstanceTypes) *DescribeInstanceTypesResponseBody {
	s.InstanceTypes = v
	return s
}

func (s *DescribeInstanceTypesResponseBody) SetRequestId(v string) *DescribeInstanceTypesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceTypesResponseBodyInstanceTypes struct {
	InstanceType []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) SetInstanceType(v []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) *DescribeInstanceTypesResponseBodyInstanceTypes {
	s.InstanceType = v
	return s
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceType struct {
	CpuCoreCount     *int32  `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	InstanceTypeId   *string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty"`
	InstanceTypeName *string `json:"InstanceTypeName,omitempty" xml:"InstanceTypeName,omitempty"`
	MemorySize       *int32  `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetCpuCoreCount(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.CpuCoreCount = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceTypeId(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceTypeId = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceTypeName(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceTypeName = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetMemorySize(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.MemorySize = &v
	return s
}

type DescribeInstanceTypesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponse) SetHeaders(v map[string]*string) *DescribeInstanceTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTypesResponse) SetBody(v *DescribeInstanceTypesResponseBody) *DescribeInstanceTypesResponse {
	s.Body = v
	return s
}

type UnAssociateEnsEipAddressRequest struct {
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
}

func (s UnAssociateEnsEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s UnAssociateEnsEipAddressRequest) GoString() string {
	return s.String()
}

func (s *UnAssociateEnsEipAddressRequest) SetAllocationId(v string) *UnAssociateEnsEipAddressRequest {
	s.AllocationId = &v
	return s
}

type UnAssociateEnsEipAddressResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnAssociateEnsEipAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnAssociateEnsEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UnAssociateEnsEipAddressResponseBody) SetRequestId(v string) *UnAssociateEnsEipAddressResponseBody {
	s.RequestId = &v
	return s
}

type UnAssociateEnsEipAddressResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnAssociateEnsEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnAssociateEnsEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s UnAssociateEnsEipAddressResponse) GoString() string {
	return s.String()
}

func (s *UnAssociateEnsEipAddressResponse) SetHeaders(v map[string]*string) *UnAssociateEnsEipAddressResponse {
	s.Headers = v
	return s
}

func (s *UnAssociateEnsEipAddressResponse) SetBody(v *UnAssociateEnsEipAddressResponseBody) *UnAssociateEnsEipAddressResponse {
	s.Body = v
	return s
}

type GetVmListRequest struct {
	AliUid     *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	GroupUuid  *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetVmListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVmListRequest) GoString() string {
	return s.String()
}

func (s *GetVmListRequest) SetAliUid(v int64) *GetVmListRequest {
	s.AliUid = &v
	return s
}

func (s *GetVmListRequest) SetGroupUuid(v string) *GetVmListRequest {
	s.GroupUuid = &v
	return s
}

func (s *GetVmListRequest) SetPageNumber(v int32) *GetVmListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetVmListRequest) SetPageSize(v int32) *GetVmListRequest {
	s.PageSize = &v
	return s
}

type GetVmListResponseBody struct {
	// 业务状态码
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回信息
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// 业务数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s GetVmListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVmListResponseBody) GoString() string {
	return s.String()
}

func (s *GetVmListResponseBody) SetCode(v int32) *GetVmListResponseBody {
	s.Code = &v
	return s
}

func (s *GetVmListResponseBody) SetRequestId(v string) *GetVmListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVmListResponseBody) SetMsg(v string) *GetVmListResponseBody {
	s.Msg = &v
	return s
}

func (s *GetVmListResponseBody) SetData(v string) *GetVmListResponseBody {
	s.Data = &v
	return s
}

func (s *GetVmListResponseBody) SetDesc(v string) *GetVmListResponseBody {
	s.Desc = &v
	return s
}

type GetVmListResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVmListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVmListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVmListResponse) GoString() string {
	return s.String()
}

func (s *GetVmListResponse) SetHeaders(v map[string]*string) *GetVmListResponse {
	s.Headers = v
	return s
}

func (s *GetVmListResponse) SetBody(v *GetVmListResponseBody) *GetVmListResponse {
	s.Body = v
	return s
}

type RebootInstanceRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ForceStop  *string `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
}

func (s RebootInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebootInstanceRequest) SetVersion(v string) *RebootInstanceRequest {
	s.Version = &v
	return s
}

func (s *RebootInstanceRequest) SetInstanceId(v string) *RebootInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RebootInstanceRequest) SetForceStop(v string) *RebootInstanceRequest {
	s.ForceStop = &v
	return s
}

type RebootInstanceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebootInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RebootInstanceResponseBody) SetCode(v int32) *RebootInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RebootInstanceResponseBody) SetRequestId(v string) *RebootInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RebootInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RebootInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebootInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebootInstanceResponse) SetHeaders(v map[string]*string) *RebootInstanceResponse {
	s.Headers = v
	return s
}

func (s *RebootInstanceResponse) SetBody(v *RebootInstanceResponseBody) *RebootInstanceResponse {
	s.Body = v
	return s
}

type DescribePrePaidInstanceStockRequest struct {
	Version        *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsRegionId    *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	SystemDiskSize *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	DataDiskSize   *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	InstanceSpec   *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
}

func (s DescribePrePaidInstanceStockRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePrePaidInstanceStockRequest) GoString() string {
	return s.String()
}

func (s *DescribePrePaidInstanceStockRequest) SetVersion(v string) *DescribePrePaidInstanceStockRequest {
	s.Version = &v
	return s
}

func (s *DescribePrePaidInstanceStockRequest) SetEnsRegionId(v string) *DescribePrePaidInstanceStockRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribePrePaidInstanceStockRequest) SetSystemDiskSize(v int32) *DescribePrePaidInstanceStockRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribePrePaidInstanceStockRequest) SetDataDiskSize(v int32) *DescribePrePaidInstanceStockRequest {
	s.DataDiskSize = &v
	return s
}

func (s *DescribePrePaidInstanceStockRequest) SetInstanceSpec(v string) *DescribePrePaidInstanceStockRequest {
	s.InstanceSpec = &v
	return s
}

type DescribePrePaidInstanceStockResponseBody struct {
	AvaliableCount *int32  `json:"AvaliableCount,omitempty" xml:"AvaliableCount,omitempty"`
	Cores          *int32  `json:"Cores,omitempty" xml:"Cores,omitempty"`
	DataDiskSize   *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	EnsRegionId    *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	InstanceSpec   *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	Memory         *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SystemDiskSize *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribePrePaidInstanceStockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePrePaidInstanceStockResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrePaidInstanceStockResponseBody) SetAvaliableCount(v int32) *DescribePrePaidInstanceStockResponseBody {
	s.AvaliableCount = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetCores(v int32) *DescribePrePaidInstanceStockResponseBody {
	s.Cores = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetDataDiskSize(v int32) *DescribePrePaidInstanceStockResponseBody {
	s.DataDiskSize = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetEnsRegionId(v string) *DescribePrePaidInstanceStockResponseBody {
	s.EnsRegionId = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetInstanceSpec(v string) *DescribePrePaidInstanceStockResponseBody {
	s.InstanceSpec = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetMemory(v int32) *DescribePrePaidInstanceStockResponseBody {
	s.Memory = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetRequestId(v string) *DescribePrePaidInstanceStockResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetSystemDiskSize(v int32) *DescribePrePaidInstanceStockResponseBody {
	s.SystemDiskSize = &v
	return s
}

type DescribePrePaidInstanceStockResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePrePaidInstanceStockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePrePaidInstanceStockResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePrePaidInstanceStockResponse) GoString() string {
	return s.String()
}

func (s *DescribePrePaidInstanceStockResponse) SetHeaders(v map[string]*string) *DescribePrePaidInstanceStockResponse {
	s.Headers = v
	return s
}

func (s *DescribePrePaidInstanceStockResponse) SetBody(v *DescribePrePaidInstanceStockResponseBody) *DescribePrePaidInstanceStockResponse {
	s.Body = v
	return s
}

type DescribeEnsRegionIdIpv6InfoRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
}

func (s DescribeEnsRegionIdIpv6InfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdIpv6InfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdIpv6InfoRequest) SetVersion(v string) *DescribeEnsRegionIdIpv6InfoRequest {
	s.Version = &v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoRequest) SetEnsRegionId(v string) *DescribeEnsRegionIdIpv6InfoRequest {
	s.EnsRegionId = &v
	return s
}

type DescribeEnsRegionIdIpv6InfoResponseBody struct {
	RequestId       *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupportIpv6Info *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info `json:"SupportIpv6Info,omitempty" xml:"SupportIpv6Info,omitempty" type:"Struct"`
}

func (s DescribeEnsRegionIdIpv6InfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdIpv6InfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdIpv6InfoResponseBody) SetRequestId(v string) *DescribeEnsRegionIdIpv6InfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoResponseBody) SetSupportIpv6Info(v *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info) *DescribeEnsRegionIdIpv6InfoResponseBody {
	s.SupportIpv6Info = v
	return s
}

type DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info struct {
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	SupportIpv6 *bool   `json:"SupportIpv6,omitempty" xml:"SupportIpv6,omitempty"`
}

func (s DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info) SetEnsRegionId(v string) *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info) SetSupportIpv6(v bool) *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info {
	s.SupportIpv6 = &v
	return s
}

type DescribeEnsRegionIdIpv6InfoResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEnsRegionIdIpv6InfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEnsRegionIdIpv6InfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdIpv6InfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdIpv6InfoResponse) SetHeaders(v map[string]*string) *DescribeEnsRegionIdIpv6InfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoResponse) SetBody(v *DescribeEnsRegionIdIpv6InfoResponseBody) *DescribeEnsRegionIdIpv6InfoResponse {
	s.Body = v
	return s
}

type DescribeApplicationRequest struct {
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersions         *string `json:"AppVersions,omitempty" xml:"AppVersions,omitempty"`
	Level               *string `json:"Level,omitempty" xml:"Level,omitempty"`
	OutDetailStatParams *string `json:"OutDetailStatParams,omitempty" xml:"OutDetailStatParams,omitempty"`
}

func (s DescribeApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationRequest) SetAppId(v string) *DescribeApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationRequest) SetAppVersions(v string) *DescribeApplicationRequest {
	s.AppVersions = &v
	return s
}

func (s *DescribeApplicationRequest) SetLevel(v string) *DescribeApplicationRequest {
	s.Level = &v
	return s
}

func (s *DescribeApplicationRequest) SetOutDetailStatParams(v string) *DescribeApplicationRequest {
	s.OutDetailStatParams = &v
	return s
}

type DescribeApplicationResponseBody struct {
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationResponseBody) SetApplication(v string) *DescribeApplicationResponseBody {
	s.Application = &v
	return s
}

func (s *DescribeApplicationResponseBody) SetRequestId(v string) *DescribeApplicationResponseBody {
	s.RequestId = &v
	return s
}

type DescribeApplicationResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationResponse) SetHeaders(v map[string]*string) *DescribeApplicationResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationResponse) SetBody(v *DescribeApplicationResponseBody) *DescribeApplicationResponse {
	s.Body = v
	return s
}

type DescribeEnsRegionIdResourceRequest struct {
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OrderByParams *string `json:"OrderByParams,omitempty" xml:"OrderByParams,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Isp           *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeEnsRegionIdResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdResourceRequest) SetVersion(v string) *DescribeEnsRegionIdResourceRequest {
	s.Version = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetStartTime(v string) *DescribeEnsRegionIdResourceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetEndTime(v string) *DescribeEnsRegionIdResourceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetOrderByParams(v string) *DescribeEnsRegionIdResourceRequest {
	s.OrderByParams = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetPageNumber(v int32) *DescribeEnsRegionIdResourceRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetPageSize(v string) *DescribeEnsRegionIdResourceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetIsp(v string) *DescribeEnsRegionIdResourceRequest {
	s.Isp = &v
	return s
}

type DescribeEnsRegionIdResourceResponseBody struct {
	EnsRegionIdResources *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources `json:"EnsRegionIdResources,omitempty" xml:"EnsRegionIdResources,omitempty" type:"Struct"`
	PageNumber           *int32                                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount           *int32                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEnsRegionIdResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdResourceResponseBody) SetEnsRegionIdResources(v *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources) *DescribeEnsRegionIdResourceResponseBody {
	s.EnsRegionIdResources = v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBody) SetPageNumber(v int32) *DescribeEnsRegionIdResourceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBody) SetPageSize(v int32) *DescribeEnsRegionIdResourceResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBody) SetRequestId(v string) *DescribeEnsRegionIdResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBody) SetTotalCount(v int32) *DescribeEnsRegionIdResourceResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources struct {
	EnsRegionIdResource []*DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource `json:"EnsRegionIdResource,omitempty" xml:"EnsRegionIdResource,omitempty" type:"Repeated"`
}

func (s DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources) SetEnsRegionIdResource(v []*DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources {
	s.EnsRegionIdResource = v
	return s
}

type DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource struct {
	Area              *string `json:"Area,omitempty" xml:"Area,omitempty"`
	AreaCode          *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	BizDate           *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	EnsRegionId       *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	EnsRegionIdName   *string `json:"EnsRegionIdName,omitempty" xml:"EnsRegionIdName,omitempty"`
	InstanceCount     *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	InternetBandwidth *int32  `json:"InternetBandwidth,omitempty" xml:"InternetBandwidth,omitempty"`
	Isp               *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	VCpu              *int32  `json:"VCpu,omitempty" xml:"VCpu,omitempty"`
}

func (s DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetArea(v string) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.Area = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetAreaCode(v string) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.AreaCode = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetBizDate(v string) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.BizDate = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetEnsRegionId(v string) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetEnsRegionIdName(v string) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.EnsRegionIdName = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetInstanceCount(v int32) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.InstanceCount = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetInternetBandwidth(v int32) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.InternetBandwidth = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetIsp(v string) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.Isp = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetVCpu(v int32) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.VCpu = &v
	return s
}

type DescribeEnsRegionIdResourceResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEnsRegionIdResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEnsRegionIdResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdResourceResponse) SetHeaders(v map[string]*string) *DescribeEnsRegionIdResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsRegionIdResourceResponse) SetBody(v *DescribeEnsRegionIdResourceResponseBody) *DescribeEnsRegionIdResourceResponse {
	s.Body = v
	return s
}

type DescribePriceRequest struct {
	SystemDisk         *DescribePriceRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	DataDisk           []*DescribePriceRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	Version            *string                         `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceType       *string                         `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	EnsRegionId        *string                         `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	Period             *int32                          `json:"Period,omitempty" xml:"Period,omitempty"`
	Quantity           *int32                          `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	InternetChargeType *string                         `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) SetSystemDisk(v *DescribePriceRequestSystemDisk) *DescribePriceRequest {
	s.SystemDisk = v
	return s
}

func (s *DescribePriceRequest) SetDataDisk(v []*DescribePriceRequestDataDisk) *DescribePriceRequest {
	s.DataDisk = v
	return s
}

func (s *DescribePriceRequest) SetVersion(v string) *DescribePriceRequest {
	s.Version = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceType(v string) *DescribePriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribePriceRequest) SetEnsRegionId(v string) *DescribePriceRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribePriceRequest) SetPeriod(v int32) *DescribePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceRequest) SetQuantity(v int32) *DescribePriceRequest {
	s.Quantity = &v
	return s
}

func (s *DescribePriceRequest) SetInternetChargeType(v string) *DescribePriceRequest {
	s.InternetChargeType = &v
	return s
}

type DescribePriceRequestSystemDisk struct {
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribePriceRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestSystemDisk) SetSize(v int32) *DescribePriceRequestSystemDisk {
	s.Size = &v
	return s
}

type DescribePriceRequestDataDisk struct {
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribePriceRequestDataDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequestDataDisk) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestDataDisk) SetSize(v int32) *DescribePriceRequestDataDisk {
	s.Size = &v
	return s
}

type DescribePriceResponseBody struct {
	PriceInfo *DescribePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) SetPriceInfo(v *DescribePriceResponseBodyPriceInfo) *DescribePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribePriceResponseBody) SetRequestId(v string) *DescribePriceResponseBody {
	s.RequestId = &v
	return s
}

type DescribePriceResponseBodyPriceInfo struct {
	Price *DescribePriceResponseBodyPriceInfoPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
}

func (s DescribePriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfo) SetPrice(v *DescribePriceResponseBodyPriceInfoPrice) *DescribePriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

type DescribePriceResponseBodyPriceInfoPrice struct {
	Currency      *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TradePrice    *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribePriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

type DescribePriceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribePriceResponse) SetHeaders(v map[string]*string) *DescribePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribePriceResponse) SetBody(v *DescribePriceResponseBody) *DescribePriceResponse {
	s.Body = v
	return s
}

type DescribeDataDownloadURLRequest struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DataName             *string `json:"DataName,omitempty" xml:"DataName,omitempty"`
	DataVersion          *string `json:"DataVersion,omitempty" xml:"DataVersion,omitempty"`
	ServerFilterStrategy *string `json:"ServerFilterStrategy,omitempty" xml:"ServerFilterStrategy,omitempty"`
	ExpireTimeout        *int64  `json:"ExpireTimeout,omitempty" xml:"ExpireTimeout,omitempty"`
}

func (s DescribeDataDownloadURLRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDownloadURLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataDownloadURLRequest) SetAppId(v string) *DescribeDataDownloadURLRequest {
	s.AppId = &v
	return s
}

func (s *DescribeDataDownloadURLRequest) SetDataName(v string) *DescribeDataDownloadURLRequest {
	s.DataName = &v
	return s
}

func (s *DescribeDataDownloadURLRequest) SetDataVersion(v string) *DescribeDataDownloadURLRequest {
	s.DataVersion = &v
	return s
}

func (s *DescribeDataDownloadURLRequest) SetServerFilterStrategy(v string) *DescribeDataDownloadURLRequest {
	s.ServerFilterStrategy = &v
	return s
}

func (s *DescribeDataDownloadURLRequest) SetExpireTimeout(v int64) *DescribeDataDownloadURLRequest {
	s.ExpireTimeout = &v
	return s
}

type DescribeDataDownloadURLResponseBody struct {
	Data      *DescribeDataDownloadURLResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int64                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeDataDownloadURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDownloadURLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataDownloadURLResponseBody) SetData(v *DescribeDataDownloadURLResponseBodyData) *DescribeDataDownloadURLResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDataDownloadURLResponseBody) SetRequestId(v string) *DescribeDataDownloadURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataDownloadURLResponseBody) SetCode(v int64) *DescribeDataDownloadURLResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDataDownloadURLResponseBody) SetMessage(v string) *DescribeDataDownloadURLResponseBody {
	s.Message = &v
	return s
}

type DescribeDataDownloadURLResponseBodyData struct {
	ServerList []*DescribeDataDownloadURLResponseBodyDataServerList `json:"ServerList,omitempty" xml:"ServerList,omitempty" type:"Repeated"`
	Url        *string                                              `json:"Url,omitempty" xml:"Url,omitempty"`
	ExpireTime *string                                              `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
}

func (s DescribeDataDownloadURLResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDownloadURLResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDataDownloadURLResponseBodyData) SetServerList(v []*DescribeDataDownloadURLResponseBodyDataServerList) *DescribeDataDownloadURLResponseBodyData {
	s.ServerList = v
	return s
}

func (s *DescribeDataDownloadURLResponseBodyData) SetUrl(v string) *DescribeDataDownloadURLResponseBodyData {
	s.Url = &v
	return s
}

func (s *DescribeDataDownloadURLResponseBodyData) SetExpireTime(v string) *DescribeDataDownloadURLResponseBodyData {
	s.ExpireTime = &v
	return s
}

type DescribeDataDownloadURLResponseBodyDataServerList struct {
	Host     *string `json:"Host,omitempty" xml:"Host,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDataDownloadURLResponseBodyDataServerList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDownloadURLResponseBodyDataServerList) GoString() string {
	return s.String()
}

func (s *DescribeDataDownloadURLResponseBodyDataServerList) SetHost(v string) *DescribeDataDownloadURLResponseBodyDataServerList {
	s.Host = &v
	return s
}

func (s *DescribeDataDownloadURLResponseBodyDataServerList) SetRegionId(v string) *DescribeDataDownloadURLResponseBodyDataServerList {
	s.RegionId = &v
	return s
}

type DescribeDataDownloadURLResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataDownloadURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataDownloadURLResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDownloadURLResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataDownloadURLResponse) SetHeaders(v map[string]*string) *DescribeDataDownloadURLResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataDownloadURLResponse) SetBody(v *DescribeDataDownloadURLResponseBody) *DescribeDataDownloadURLResponse {
	s.Body = v
	return s
}

type ListApplicationsRequest struct {
	ClusterNames     *string `json:"ClusterNames,omitempty" xml:"ClusterNames,omitempty"`
	AppVersions      *string `json:"AppVersions,omitempty" xml:"AppVersions,omitempty"`
	Level            *string `json:"Level,omitempty" xml:"Level,omitempty"`
	OutAppInfoParams *string `json:"OutAppInfoParams,omitempty" xml:"OutAppInfoParams,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	MinDate          *string `json:"MinDate,omitempty" xml:"MinDate,omitempty"`
	MaxDate          *string `json:"MaxDate,omitempty" xml:"MaxDate,omitempty"`
}

func (s ListApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) SetClusterNames(v string) *ListApplicationsRequest {
	s.ClusterNames = &v
	return s
}

func (s *ListApplicationsRequest) SetAppVersions(v string) *ListApplicationsRequest {
	s.AppVersions = &v
	return s
}

func (s *ListApplicationsRequest) SetLevel(v string) *ListApplicationsRequest {
	s.Level = &v
	return s
}

func (s *ListApplicationsRequest) SetOutAppInfoParams(v string) *ListApplicationsRequest {
	s.OutAppInfoParams = &v
	return s
}

func (s *ListApplicationsRequest) SetPageNumber(v int32) *ListApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsRequest) SetPageSize(v int32) *ListApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsRequest) SetMinDate(v string) *ListApplicationsRequest {
	s.MinDate = &v
	return s
}

func (s *ListApplicationsRequest) SetMaxDate(v string) *ListApplicationsRequest {
	s.MaxDate = &v
	return s
}

type ListApplicationsResponseBody struct {
	Applications *ListApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	PageNumber   *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) SetApplications(v *ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsResponseBody) SetPageNumber(v int32) *ListApplicationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsResponseBody) SetPageSize(v int32) *ListApplicationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponseBody) SetTotalCount(v int32) *ListApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicationsResponseBodyApplications struct {
	Application []*ListApplicationsResponseBodyApplicationsApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplications) SetApplication(v []*ListApplicationsResponseBodyApplicationsApplication) *ListApplicationsResponseBodyApplications {
	s.Application = v
	return s
}

type ListApplicationsResponseBodyApplicationsApplication struct {
	AppList     *ListApplicationsResponseBodyApplicationsApplicationAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Struct"`
	ClusterName *string                                                     `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
}

func (s ListApplicationsResponseBodyApplicationsApplication) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplication) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetAppList(v *ListApplicationsResponseBodyApplicationsApplicationAppList) *ListApplicationsResponseBodyApplicationsApplication {
	s.AppList = v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetClusterName(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.ClusterName = &v
	return s
}

type ListApplicationsResponseBodyApplicationsApplicationAppList struct {
	App []*ListApplicationsResponseBodyApplicationsApplicationAppListApp `json:"App,omitempty" xml:"App,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyApplicationsApplicationAppList) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplicationAppList) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppList) SetApp(v []*ListApplicationsResponseBodyApplicationsApplicationAppListApp) *ListApplicationsResponseBodyApplicationsApplicationAppList {
	s.App = v
	return s
}

type ListApplicationsResponseBodyApplicationsApplicationAppListApp struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppInfo *string `json:"AppInfo,omitempty" xml:"AppInfo,omitempty"`
}

func (s ListApplicationsResponseBodyApplicationsApplicationAppListApp) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplicationAppListApp) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppListApp) SetAppId(v string) *ListApplicationsResponseBodyApplicationsApplicationAppListApp {
	s.AppId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppListApp) SetAppInfo(v string) *ListApplicationsResponseBodyApplicationsApplicationAppListApp {
	s.AppInfo = &v
	return s
}

type ListApplicationsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponse) SetHeaders(v map[string]*string) *ListApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsResponse) SetBody(v *ListApplicationsResponseBody) *ListApplicationsResponse {
	s.Body = v
	return s
}

type PreCreateEnsServiceRequest struct {
	Version                 *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsServiceName          *string `json:"EnsServiceName,omitempty" xml:"EnsServiceName,omitempty"`
	ImageId                 *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceSpec            *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	SystemDiskSize          *string `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	DataDiskSize            *string `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	BandwidthType           *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	InstanceBandwithdLimit  *string `json:"InstanceBandwithdLimit,omitempty" xml:"InstanceBandwithdLimit,omitempty"`
	Password                *string `json:"Password,omitempty" xml:"Password,omitempty"`
	KeyPairName             *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	UserData                *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	NetLevel                *string `json:"NetLevel,omitempty" xml:"NetLevel,omitempty"`
	SchedulingStrategy      *string `json:"SchedulingStrategy,omitempty" xml:"SchedulingStrategy,omitempty"`
	SchedulingPriceStrategy *string `json:"SchedulingPriceStrategy,omitempty" xml:"SchedulingPriceStrategy,omitempty"`
	BuyResourcesDetail      *string `json:"BuyResourcesDetail,omitempty" xml:"BuyResourcesDetail,omitempty"`
}

func (s PreCreateEnsServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s PreCreateEnsServiceRequest) GoString() string {
	return s.String()
}

func (s *PreCreateEnsServiceRequest) SetVersion(v string) *PreCreateEnsServiceRequest {
	s.Version = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetEnsServiceName(v string) *PreCreateEnsServiceRequest {
	s.EnsServiceName = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetImageId(v string) *PreCreateEnsServiceRequest {
	s.ImageId = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetInstanceSpec(v string) *PreCreateEnsServiceRequest {
	s.InstanceSpec = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetSystemDiskSize(v string) *PreCreateEnsServiceRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetDataDiskSize(v string) *PreCreateEnsServiceRequest {
	s.DataDiskSize = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetBandwidthType(v string) *PreCreateEnsServiceRequest {
	s.BandwidthType = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetInstanceBandwithdLimit(v string) *PreCreateEnsServiceRequest {
	s.InstanceBandwithdLimit = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetPassword(v string) *PreCreateEnsServiceRequest {
	s.Password = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetKeyPairName(v string) *PreCreateEnsServiceRequest {
	s.KeyPairName = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetUserData(v string) *PreCreateEnsServiceRequest {
	s.UserData = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetNetLevel(v string) *PreCreateEnsServiceRequest {
	s.NetLevel = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetSchedulingStrategy(v string) *PreCreateEnsServiceRequest {
	s.SchedulingStrategy = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetSchedulingPriceStrategy(v string) *PreCreateEnsServiceRequest {
	s.SchedulingPriceStrategy = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetBuyResourcesDetail(v string) *PreCreateEnsServiceRequest {
	s.BuyResourcesDetail = &v
	return s
}

type PreCreateEnsServiceResponseBody struct {
	BuyResourcesDetail *string `json:"BuyResourcesDetail,omitempty" xml:"BuyResourcesDetail,omitempty"`
	Code               *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	EnsServiceId       *string `json:"EnsServiceId,omitempty" xml:"EnsServiceId,omitempty"`
	NetLevel           *string `json:"NetLevel,omitempty" xml:"NetLevel,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreCreateEnsServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreCreateEnsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *PreCreateEnsServiceResponseBody) SetBuyResourcesDetail(v string) *PreCreateEnsServiceResponseBody {
	s.BuyResourcesDetail = &v
	return s
}

func (s *PreCreateEnsServiceResponseBody) SetCode(v int32) *PreCreateEnsServiceResponseBody {
	s.Code = &v
	return s
}

func (s *PreCreateEnsServiceResponseBody) SetEnsServiceId(v string) *PreCreateEnsServiceResponseBody {
	s.EnsServiceId = &v
	return s
}

func (s *PreCreateEnsServiceResponseBody) SetNetLevel(v string) *PreCreateEnsServiceResponseBody {
	s.NetLevel = &v
	return s
}

func (s *PreCreateEnsServiceResponseBody) SetRequestId(v string) *PreCreateEnsServiceResponseBody {
	s.RequestId = &v
	return s
}

type PreCreateEnsServiceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PreCreateEnsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PreCreateEnsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s PreCreateEnsServiceResponse) GoString() string {
	return s.String()
}

func (s *PreCreateEnsServiceResponse) SetHeaders(v map[string]*string) *PreCreateEnsServiceResponse {
	s.Headers = v
	return s
}

func (s *PreCreateEnsServiceResponse) SetBody(v *PreCreateEnsServiceResponseBody) *PreCreateEnsServiceResponse {
	s.Body = v
	return s
}

type DescribeBandWithdChargeTypeRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeBandWithdChargeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandWithdChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeBandWithdChargeTypeRequest) SetVersion(v string) *DescribeBandWithdChargeTypeRequest {
	s.Version = &v
	return s
}

type DescribeBandWithdChargeTypeResponseBody struct {
	BandWithTypeInfo   *string `json:"BandWithTypeInfo,omitempty" xml:"BandWithTypeInfo,omitempty"`
	ChargeContractType *string `json:"ChargeContractType,omitempty" xml:"ChargeContractType,omitempty"`
	ChargeCycleInfo    *string `json:"ChargeCycleInfo,omitempty" xml:"ChargeCycleInfo,omitempty"`
	Code               *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBandWithdChargeTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandWithdChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBandWithdChargeTypeResponseBody) SetBandWithTypeInfo(v string) *DescribeBandWithdChargeTypeResponseBody {
	s.BandWithTypeInfo = &v
	return s
}

func (s *DescribeBandWithdChargeTypeResponseBody) SetChargeContractType(v string) *DescribeBandWithdChargeTypeResponseBody {
	s.ChargeContractType = &v
	return s
}

func (s *DescribeBandWithdChargeTypeResponseBody) SetChargeCycleInfo(v string) *DescribeBandWithdChargeTypeResponseBody {
	s.ChargeCycleInfo = &v
	return s
}

func (s *DescribeBandWithdChargeTypeResponseBody) SetCode(v int32) *DescribeBandWithdChargeTypeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBandWithdChargeTypeResponseBody) SetRequestId(v string) *DescribeBandWithdChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBandWithdChargeTypeResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBandWithdChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBandWithdChargeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandWithdChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeBandWithdChargeTypeResponse) SetHeaders(v map[string]*string) *DescribeBandWithdChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeBandWithdChargeTypeResponse) SetBody(v *DescribeBandWithdChargeTypeResponseBody) *DescribeBandWithdChargeTypeResponse {
	s.Body = v
	return s
}

type DescribeReservedResourceRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeReservedResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceRequest) SetVersion(v string) *DescribeReservedResourceRequest {
	s.Version = &v
	return s
}

type DescribeReservedResourceResponseBody struct {
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Images           *DescribeReservedResourceResponseBodyImages           `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	SupportResources *DescribeReservedResourceResponseBodySupportResources `json:"SupportResources,omitempty" xml:"SupportResources,omitempty" type:"Struct"`
	Code             *int32                                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeReservedResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseBody) SetRequestId(v string) *DescribeReservedResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReservedResourceResponseBody) SetImages(v *DescribeReservedResourceResponseBodyImages) *DescribeReservedResourceResponseBody {
	s.Images = v
	return s
}

func (s *DescribeReservedResourceResponseBody) SetSupportResources(v *DescribeReservedResourceResponseBodySupportResources) *DescribeReservedResourceResponseBody {
	s.SupportResources = v
	return s
}

func (s *DescribeReservedResourceResponseBody) SetCode(v int32) *DescribeReservedResourceResponseBody {
	s.Code = &v
	return s
}

type DescribeReservedResourceResponseBodyImages struct {
	Image []*DescribeReservedResourceResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeReservedResourceResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseBodyImages) SetImage(v []*DescribeReservedResourceResponseBodyImagesImage) *DescribeReservedResourceResponseBodyImages {
	s.Image = v
	return s
}

type DescribeReservedResourceResponseBodyImagesImage struct {
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s DescribeReservedResourceResponseBodyImagesImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseBodyImagesImage) SetImageName(v string) *DescribeReservedResourceResponseBodyImagesImage {
	s.ImageName = &v
	return s
}

func (s *DescribeReservedResourceResponseBodyImagesImage) SetImageId(v string) *DescribeReservedResourceResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

type DescribeReservedResourceResponseBodySupportResources struct {
	SupportResource []*DescribeReservedResourceResponseBodySupportResourcesSupportResource `json:"SupportResource,omitempty" xml:"SupportResource,omitempty" type:"Repeated"`
}

func (s DescribeReservedResourceResponseBodySupportResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponseBodySupportResources) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseBodySupportResources) SetSupportResource(v []*DescribeReservedResourceResponseBodySupportResourcesSupportResource) *DescribeReservedResourceResponseBodySupportResources {
	s.SupportResource = v
	return s
}

type DescribeReservedResourceResponseBodySupportResourcesSupportResource struct {
	EnsRegionId           *string                                                                             `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	DataDiskSizes         *DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes   `json:"DataDiskSizes,omitempty" xml:"DataDiskSizes,omitempty" type:"Struct"`
	SupportResourcesCount *string                                                                             `json:"SupportResourcesCount,omitempty" xml:"SupportResourcesCount,omitempty"`
	SystemDiskSizes       *DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes `json:"SystemDiskSizes,omitempty" xml:"SystemDiskSizes,omitempty" type:"Struct"`
	InstanceSpec          *string                                                                             `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
}

func (s DescribeReservedResourceResponseBodySupportResourcesSupportResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponseBodySupportResourcesSupportResource) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) SetEnsRegionId(v string) *DescribeReservedResourceResponseBodySupportResourcesSupportResource {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) SetDataDiskSizes(v *DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes) *DescribeReservedResourceResponseBodySupportResourcesSupportResource {
	s.DataDiskSizes = v
	return s
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) SetSupportResourcesCount(v string) *DescribeReservedResourceResponseBodySupportResourcesSupportResource {
	s.SupportResourcesCount = &v
	return s
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) SetSystemDiskSizes(v *DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes) *DescribeReservedResourceResponseBodySupportResourcesSupportResource {
	s.SystemDiskSizes = v
	return s
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResource) SetInstanceSpec(v string) *DescribeReservedResourceResponseBodySupportResourcesSupportResource {
	s.InstanceSpec = &v
	return s
}

type DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes struct {
	DataDiskSize []*string `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" type:"Repeated"`
}

func (s DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes) SetDataDiskSize(v []*string) *DescribeReservedResourceResponseBodySupportResourcesSupportResourceDataDiskSizes {
	s.DataDiskSize = v
	return s
}

type DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes struct {
	SystemDiskSize []*string `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" type:"Repeated"`
}

func (s DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes) SetSystemDiskSize(v []*string) *DescribeReservedResourceResponseBodySupportResourcesSupportResourceSystemDiskSizes {
	s.SystemDiskSize = v
	return s
}

type DescribeReservedResourceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeReservedResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeReservedResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponse) SetHeaders(v map[string]*string) *DescribeReservedResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeReservedResourceResponse) SetBody(v *DescribeReservedResourceResponseBody) *DescribeReservedResourceResponse {
	s.Body = v
	return s
}

type DescribeCreatePrePaidInstanceResultRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeCreatePrePaidInstanceResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCreatePrePaidInstanceResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeCreatePrePaidInstanceResultRequest) SetVersion(v string) *DescribeCreatePrePaidInstanceResultRequest {
	s.Version = &v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultRequest) SetInstanceId(v string) *DescribeCreatePrePaidInstanceResultRequest {
	s.InstanceId = &v
	return s
}

type DescribeCreatePrePaidInstanceResultResponseBody struct {
	InstanceCreateResult *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult `json:"InstanceCreateResult,omitempty" xml:"InstanceCreateResult,omitempty" type:"Struct"`
	RequestId            *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCreatePrePaidInstanceResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCreatePrePaidInstanceResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCreatePrePaidInstanceResultResponseBody) SetInstanceCreateResult(v *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult) *DescribeCreatePrePaidInstanceResultResponseBody {
	s.InstanceCreateResult = v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultResponseBody) SetRequestId(v string) *DescribeCreatePrePaidInstanceResultResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult struct {
	InstanceCreateStatus *string `json:"InstanceCreateStatus,omitempty" xml:"InstanceCreateStatus,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult) GoString() string {
	return s.String()
}

func (s *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult) SetInstanceCreateStatus(v string) *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult {
	s.InstanceCreateStatus = &v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult) SetInstanceId(v string) *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult {
	s.InstanceId = &v
	return s
}

type DescribeCreatePrePaidInstanceResultResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCreatePrePaidInstanceResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCreatePrePaidInstanceResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCreatePrePaidInstanceResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeCreatePrePaidInstanceResultResponse) SetHeaders(v map[string]*string) *DescribeCreatePrePaidInstanceResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultResponse) SetBody(v *DescribeCreatePrePaidInstanceResultResponseBody) *DescribeCreatePrePaidInstanceResultResponse {
	s.Body = v
	return s
}

type MigrateVmRequest struct {
	Tenant      *string `json:"Tenant,omitempty" xml:"Tenant,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	GroupUuid   *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
}

func (s MigrateVmRequest) String() string {
	return tea.Prettify(s)
}

func (s MigrateVmRequest) GoString() string {
	return s.String()
}

func (s *MigrateVmRequest) SetTenant(v string) *MigrateVmRequest {
	s.Tenant = &v
	return s
}

func (s *MigrateVmRequest) SetInstanceIds(v string) *MigrateVmRequest {
	s.InstanceIds = &v
	return s
}

func (s *MigrateVmRequest) SetGroupUuid(v string) *MigrateVmRequest {
	s.GroupUuid = &v
	return s
}

type MigrateVmResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Desc      *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s MigrateVmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MigrateVmResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateVmResponseBody) SetRequestId(v string) *MigrateVmResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateVmResponseBody) SetCode(v int32) *MigrateVmResponseBody {
	s.Code = &v
	return s
}

func (s *MigrateVmResponseBody) SetData(v string) *MigrateVmResponseBody {
	s.Data = &v
	return s
}

func (s *MigrateVmResponseBody) SetMsg(v string) *MigrateVmResponseBody {
	s.Msg = &v
	return s
}

func (s *MigrateVmResponseBody) SetDesc(v string) *MigrateVmResponseBody {
	s.Desc = &v
	return s
}

type MigrateVmResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MigrateVmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MigrateVmResponse) String() string {
	return tea.Prettify(s)
}

func (s MigrateVmResponse) GoString() string {
	return s.String()
}

func (s *MigrateVmResponse) SetHeaders(v map[string]*string) *MigrateVmResponse {
	s.Headers = v
	return s
}

func (s *MigrateVmResponse) SetBody(v *MigrateVmResponseBody) *MigrateVmResponse {
	s.Body = v
	return s
}

type JoinVSwitchesToEpnInstanceRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	VSwitchesInfo *string `json:"VSwitchesInfo,omitempty" xml:"VSwitchesInfo,omitempty"`
}

func (s JoinVSwitchesToEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinVSwitchesToEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *JoinVSwitchesToEpnInstanceRequest) SetEPNInstanceId(v string) *JoinVSwitchesToEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *JoinVSwitchesToEpnInstanceRequest) SetVSwitchesInfo(v string) *JoinVSwitchesToEpnInstanceRequest {
	s.VSwitchesInfo = &v
	return s
}

type JoinVSwitchesToEpnInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JoinVSwitchesToEpnInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinVSwitchesToEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *JoinVSwitchesToEpnInstanceResponseBody) SetRequestId(v string) *JoinVSwitchesToEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

type JoinVSwitchesToEpnInstanceResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *JoinVSwitchesToEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinVSwitchesToEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinVSwitchesToEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *JoinVSwitchesToEpnInstanceResponse) SetHeaders(v map[string]*string) *JoinVSwitchesToEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *JoinVSwitchesToEpnInstanceResponse) SetBody(v *JoinVSwitchesToEpnInstanceResponseBody) *JoinVSwitchesToEpnInstanceResponse {
	s.Body = v
	return s
}

type DeleteVSwitchRequest struct {
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DeleteVSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVSwitchRequest) GoString() string {
	return s.String()
}

func (s *DeleteVSwitchRequest) SetVersion(v string) *DeleteVSwitchRequest {
	s.Version = &v
	return s
}

func (s *DeleteVSwitchRequest) SetVSwitchId(v string) *DeleteVSwitchRequest {
	s.VSwitchId = &v
	return s
}

type DeleteVSwitchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVSwitchResponseBody) SetRequestId(v string) *DeleteVSwitchResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVSwitchResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVSwitchResponse) GoString() string {
	return s.String()
}

func (s *DeleteVSwitchResponse) SetHeaders(v map[string]*string) *DeleteVSwitchResponse {
	s.Headers = v
	return s
}

func (s *DeleteVSwitchResponse) SetBody(v *DeleteVSwitchResponseBody) *DeleteVSwitchResponse {
	s.Body = v
	return s
}

type DescribeInstanceVncUrlRequest struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceVncUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceVncUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVncUrlRequest) SetInstanceId(v string) *DescribeInstanceVncUrlRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceVncUrlResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 管理终端Url。
	VncUrl *string `json:"VncUrl,omitempty" xml:"VncUrl,omitempty"`
}

func (s DescribeInstanceVncUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceVncUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVncUrlResponseBody) SetRequestId(v string) *DescribeInstanceVncUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceVncUrlResponseBody) SetVncUrl(v string) *DescribeInstanceVncUrlResponseBody {
	s.VncUrl = &v
	return s
}

type DescribeInstanceVncUrlResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceVncUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceVncUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceVncUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVncUrlResponse) SetHeaders(v map[string]*string) *DescribeInstanceVncUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceVncUrlResponse) SetBody(v *DescribeInstanceVncUrlResponseBody) *DescribeInstanceVncUrlResponse {
	s.Body = v
	return s
}

type DescribeImagesRequest struct {
	Product     *string `json:"product,omitempty" xml:"product,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ImageName   *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	PageNumber  *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagesRequest) SetProduct(v string) *DescribeImagesRequest {
	s.Product = &v
	return s
}

func (s *DescribeImagesRequest) SetVersion(v string) *DescribeImagesRequest {
	s.Version = &v
	return s
}

func (s *DescribeImagesRequest) SetEnsRegionId(v string) *DescribeImagesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeImagesRequest) SetImageId(v string) *DescribeImagesRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesRequest) SetStatus(v string) *DescribeImagesRequest {
	s.Status = &v
	return s
}

func (s *DescribeImagesRequest) SetImageName(v string) *DescribeImagesRequest {
	s.ImageName = &v
	return s
}

func (s *DescribeImagesRequest) SetPageNumber(v string) *DescribeImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeImagesRequest) SetPageSize(v string) *DescribeImagesRequest {
	s.PageSize = &v
	return s
}

type DescribeImagesResponseBody struct {
	Code       *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Images     *DescribeImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBody) SetCode(v int32) *DescribeImagesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImagesResponseBody) SetImages(v *DescribeImagesResponseBodyImages) *DescribeImagesResponseBody {
	s.Images = v
	return s
}

func (s *DescribeImagesResponseBody) SetPageNumber(v int32) *DescribeImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeImagesResponseBody) SetPageSize(v int32) *DescribeImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeImagesResponseBody) SetRequestId(v string) *DescribeImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImagesResponseBody) SetTotalCount(v int32) *DescribeImagesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeImagesResponseBodyImages struct {
	Image []*DescribeImagesResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImages) SetImage(v []*DescribeImagesResponseBodyImagesImage) *DescribeImagesResponseBodyImages {
	s.Image = v
	return s
}

type DescribeImagesResponseBodyImagesImage struct {
	Architecture    *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	CreationTime    *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName       *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	ImageSize       *string `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	Platform        *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s DescribeImagesResponseBodyImagesImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImagesImage) SetArchitecture(v string) *DescribeImagesResponseBodyImagesImage {
	s.Architecture = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetCreationTime(v string) *DescribeImagesResponseBodyImagesImage {
	s.CreationTime = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetImageId(v string) *DescribeImagesResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetImageName(v string) *DescribeImagesResponseBodyImagesImage {
	s.ImageName = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetImageOwnerAlias(v string) *DescribeImagesResponseBodyImagesImage {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetImageSize(v string) *DescribeImagesResponseBodyImagesImage {
	s.ImageSize = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetPlatform(v string) *DescribeImagesResponseBodyImagesImage {
	s.Platform = &v
	return s
}

type DescribeImagesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponse) SetHeaders(v map[string]*string) *DescribeImagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeImagesResponse) SetBody(v *DescribeImagesResponseBody) *DescribeImagesResponse {
	s.Body = v
	return s
}

type AuthorizeSecurityGroupEgressRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Priority        *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	DestCidrIp      *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s AuthorizeSecurityGroupEgressRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeSecurityGroupEgressRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupEgressRequest) SetVersion(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Version = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetIpProtocol(v string) *AuthorizeSecurityGroupEgressRequest {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPortRange(v string) *AuthorizeSecurityGroupEgressRequest {
	s.PortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetSecurityGroupId(v string) *AuthorizeSecurityGroupEgressRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPolicy(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Policy = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPriority(v int32) *AuthorizeSecurityGroupEgressRequest {
	s.Priority = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDestCidrIp(v string) *AuthorizeSecurityGroupEgressRequest {
	s.DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetSourcePortRange(v string) *AuthorizeSecurityGroupEgressRequest {
	s.SourcePortRange = &v
	return s
}

type AuthorizeSecurityGroupEgressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeSecurityGroupEgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeSecurityGroupEgressResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupEgressResponseBody) SetRequestId(v string) *AuthorizeSecurityGroupEgressResponseBody {
	s.RequestId = &v
	return s
}

type AuthorizeSecurityGroupEgressResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AuthorizeSecurityGroupEgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthorizeSecurityGroupEgressResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeSecurityGroupEgressResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupEgressResponse) SetHeaders(v map[string]*string) *AuthorizeSecurityGroupEgressResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeSecurityGroupEgressResponse) SetBody(v *AuthorizeSecurityGroupEgressResponseBody) *AuthorizeSecurityGroupEgressResponse {
	s.Body = v
	return s
}

type DescribeEpnInstanceAttributeRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
}

func (s DescribeEpnInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeRequest) SetEPNInstanceId(v string) *DescribeEpnInstanceAttributeRequest {
	s.EPNInstanceId = &v
	return s
}

type DescribeEpnInstanceAttributeResponseBody struct {
	ConfVersions    []*DescribeEpnInstanceAttributeResponseBodyConfVersions `json:"ConfVersions,omitempty" xml:"ConfVersions,omitempty" type:"Repeated"`
	EPNInstanceId   *string                                                 `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	EPNInstanceName *string                                                 `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	Instances       []*DescribeEpnInstanceAttributeResponseBodyInstances    `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	NetworkingModel *string                                                 `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
	RequestId       *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VSwitches       []*DescribeEpnInstanceAttributeResponseBodyVSwitches    `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
}

func (s DescribeEpnInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeResponseBody) SetConfVersions(v []*DescribeEpnInstanceAttributeResponseBodyConfVersions) *DescribeEpnInstanceAttributeResponseBody {
	s.ConfVersions = v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBody) SetEPNInstanceId(v string) *DescribeEpnInstanceAttributeResponseBody {
	s.EPNInstanceId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBody) SetEPNInstanceName(v string) *DescribeEpnInstanceAttributeResponseBody {
	s.EPNInstanceName = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBody) SetInstances(v []*DescribeEpnInstanceAttributeResponseBodyInstances) *DescribeEpnInstanceAttributeResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBody) SetNetworkingModel(v string) *DescribeEpnInstanceAttributeResponseBody {
	s.NetworkingModel = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBody) SetRequestId(v string) *DescribeEpnInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBody) SetVSwitches(v []*DescribeEpnInstanceAttributeResponseBodyVSwitches) *DescribeEpnInstanceAttributeResponseBody {
	s.VSwitches = v
	return s
}

type DescribeEpnInstanceAttributeResponseBodyConfVersions struct {
	ConfVersion *string `json:"ConfVersion,omitempty" xml:"ConfVersion,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
}

func (s DescribeEpnInstanceAttributeResponseBodyConfVersions) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstanceAttributeResponseBodyConfVersions) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeResponseBodyConfVersions) SetConfVersion(v string) *DescribeEpnInstanceAttributeResponseBodyConfVersions {
	s.ConfVersion = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyConfVersions) SetEnsRegionId(v string) *DescribeEpnInstanceAttributeResponseBodyConfVersions {
	s.EnsRegionId = &v
	return s
}

type DescribeEpnInstanceAttributeResponseBodyInstances struct {
	EnsRegionId      *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Isp              *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	PublicIpAddress  *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEpnInstanceAttributeResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstanceAttributeResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) SetEnsRegionId(v string) *DescribeEpnInstanceAttributeResponseBodyInstances {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) SetInstanceId(v string) *DescribeEpnInstanceAttributeResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) SetInstanceName(v string) *DescribeEpnInstanceAttributeResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) SetIsp(v string) *DescribeEpnInstanceAttributeResponseBodyInstances {
	s.Isp = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) SetPrivateIpAddress(v string) *DescribeEpnInstanceAttributeResponseBodyInstances {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) SetPublicIpAddress(v string) *DescribeEpnInstanceAttributeResponseBodyInstances {
	s.PublicIpAddress = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) SetStatus(v string) *DescribeEpnInstanceAttributeResponseBodyInstances {
	s.Status = &v
	return s
}

type DescribeEpnInstanceAttributeResponseBodyVSwitches struct {
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeEpnInstanceAttributeResponseBodyVSwitches) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstanceAttributeResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeResponseBodyVSwitches) SetCidrBlock(v string) *DescribeEpnInstanceAttributeResponseBodyVSwitches {
	s.CidrBlock = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyVSwitches) SetEnsRegionId(v string) *DescribeEpnInstanceAttributeResponseBodyVSwitches {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyVSwitches) SetVSwitchId(v string) *DescribeEpnInstanceAttributeResponseBodyVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyVSwitches) SetVSwitchName(v string) *DescribeEpnInstanceAttributeResponseBodyVSwitches {
	s.VSwitchName = &v
	return s
}

type DescribeEpnInstanceAttributeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEpnInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEpnInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeEpnInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEpnInstanceAttributeResponse) SetBody(v *DescribeEpnInstanceAttributeResponseBody) *DescribeEpnInstanceAttributeResponse {
	s.Body = v
	return s
}

type ReleaseInstanceRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceRequest) SetInstanceId(v string) *ReleaseInstanceRequest {
	s.InstanceId = &v
	return s
}

type ReleaseInstanceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponseBody) SetRequestId(v string) *ReleaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponse) SetHeaders(v map[string]*string) *ReleaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstanceResponse) SetBody(v *ReleaseInstanceResponseBody) *ReleaseInstanceResponse {
	s.Body = v
	return s
}

type AssociateEnsEipAddressRequest struct {
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s AssociateEnsEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateEnsEipAddressRequest) GoString() string {
	return s.String()
}

func (s *AssociateEnsEipAddressRequest) SetAllocationId(v string) *AssociateEnsEipAddressRequest {
	s.AllocationId = &v
	return s
}

func (s *AssociateEnsEipAddressRequest) SetInstanceId(v string) *AssociateEnsEipAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateEnsEipAddressRequest) SetInstanceType(v string) *AssociateEnsEipAddressRequest {
	s.InstanceType = &v
	return s
}

type AssociateEnsEipAddressResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateEnsEipAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateEnsEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateEnsEipAddressResponseBody) SetRequestId(v string) *AssociateEnsEipAddressResponseBody {
	s.RequestId = &v
	return s
}

type AssociateEnsEipAddressResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssociateEnsEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateEnsEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateEnsEipAddressResponse) GoString() string {
	return s.String()
}

func (s *AssociateEnsEipAddressResponse) SetHeaders(v map[string]*string) *AssociateEnsEipAddressResponse {
	s.Headers = v
	return s
}

func (s *AssociateEnsEipAddressResponse) SetBody(v *AssociateEnsEipAddressResponseBody) *AssociateEnsEipAddressResponse {
	s.Body = v
	return s
}

type PushApplicationDataRequest struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Timeout      *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	PushStrategy *string `json:"PushStrategy,omitempty" xml:"PushStrategy,omitempty"`
}

func (s PushApplicationDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PushApplicationDataRequest) GoString() string {
	return s.String()
}

func (s *PushApplicationDataRequest) SetData(v string) *PushApplicationDataRequest {
	s.Data = &v
	return s
}

func (s *PushApplicationDataRequest) SetAppId(v string) *PushApplicationDataRequest {
	s.AppId = &v
	return s
}

func (s *PushApplicationDataRequest) SetTimeout(v int32) *PushApplicationDataRequest {
	s.Timeout = &v
	return s
}

func (s *PushApplicationDataRequest) SetPushStrategy(v string) *PushApplicationDataRequest {
	s.PushStrategy = &v
	return s
}

type PushApplicationDataResponseBody struct {
	PushResults *PushApplicationDataResponseBodyPushResults `json:"PushResults,omitempty" xml:"PushResults,omitempty" type:"Struct"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushApplicationDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushApplicationDataResponseBody) GoString() string {
	return s.String()
}

func (s *PushApplicationDataResponseBody) SetPushResults(v *PushApplicationDataResponseBodyPushResults) *PushApplicationDataResponseBody {
	s.PushResults = v
	return s
}

func (s *PushApplicationDataResponseBody) SetRequestId(v string) *PushApplicationDataResponseBody {
	s.RequestId = &v
	return s
}

type PushApplicationDataResponseBodyPushResults struct {
	PushResult []*PushApplicationDataResponseBodyPushResultsPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" type:"Repeated"`
}

func (s PushApplicationDataResponseBodyPushResults) String() string {
	return tea.Prettify(s)
}

func (s PushApplicationDataResponseBodyPushResults) GoString() string {
	return s.String()
}

func (s *PushApplicationDataResponseBodyPushResults) SetPushResult(v []*PushApplicationDataResponseBodyPushResultsPushResult) *PushApplicationDataResponseBodyPushResults {
	s.PushResult = v
	return s
}

type PushApplicationDataResponseBodyPushResultsPushResult struct {
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResultCode    *int32  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultDescrip *string `json:"ResultDescrip,omitempty" xml:"ResultDescrip,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PushApplicationDataResponseBodyPushResultsPushResult) String() string {
	return tea.Prettify(s)
}

func (s PushApplicationDataResponseBodyPushResultsPushResult) GoString() string {
	return s.String()
}

func (s *PushApplicationDataResponseBodyPushResultsPushResult) SetName(v string) *PushApplicationDataResponseBodyPushResultsPushResult {
	s.Name = &v
	return s
}

func (s *PushApplicationDataResponseBodyPushResultsPushResult) SetResultCode(v int32) *PushApplicationDataResponseBodyPushResultsPushResult {
	s.ResultCode = &v
	return s
}

func (s *PushApplicationDataResponseBodyPushResultsPushResult) SetResultDescrip(v string) *PushApplicationDataResponseBodyPushResultsPushResult {
	s.ResultDescrip = &v
	return s
}

func (s *PushApplicationDataResponseBodyPushResultsPushResult) SetVersion(v string) *PushApplicationDataResponseBodyPushResultsPushResult {
	s.Version = &v
	return s
}

type PushApplicationDataResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PushApplicationDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushApplicationDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PushApplicationDataResponse) GoString() string {
	return s.String()
}

func (s *PushApplicationDataResponse) SetHeaders(v map[string]*string) *PushApplicationDataResponse {
	s.Headers = v
	return s
}

func (s *PushApplicationDataResponse) SetBody(v *PushApplicationDataResponseBody) *PushApplicationDataResponse {
	s.Body = v
	return s
}

type DescribeEnsRegionsRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
}

func (s DescribeEnsRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionsRequest) SetVersion(v string) *DescribeEnsRegionsRequest {
	s.Version = &v
	return s
}

func (s *DescribeEnsRegionsRequest) SetEnsRegionId(v string) *DescribeEnsRegionsRequest {
	s.EnsRegionId = &v
	return s
}

type DescribeEnsRegionsResponseBody struct {
	Code       *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	EnsRegions *DescribeEnsRegionsResponseBodyEnsRegions `json:"EnsRegions,omitempty" xml:"EnsRegions,omitempty" type:"Struct"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnsRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionsResponseBody) SetCode(v int32) *DescribeEnsRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnsRegionsResponseBody) SetEnsRegions(v *DescribeEnsRegionsResponseBodyEnsRegions) *DescribeEnsRegionsResponseBody {
	s.EnsRegions = v
	return s
}

func (s *DescribeEnsRegionsResponseBody) SetRequestId(v string) *DescribeEnsRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEnsRegionsResponseBodyEnsRegions struct {
	EnsRegions []*DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions `json:"EnsRegions,omitempty" xml:"EnsRegions,omitempty" type:"Repeated"`
}

func (s DescribeEnsRegionsResponseBodyEnsRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionsResponseBodyEnsRegions) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionsResponseBodyEnsRegions) SetEnsRegions(v []*DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) *DescribeEnsRegionsResponseBodyEnsRegions {
	s.EnsRegions = v
	return s
}

type DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions struct {
	Area        *string `json:"Area,omitempty" xml:"Area,omitempty"`
	EnName      *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) SetArea(v string) *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions {
	s.Area = &v
	return s
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) SetEnName(v string) *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions {
	s.EnName = &v
	return s
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) SetEnsRegionId(v string) *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) SetName(v string) *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions {
	s.Name = &v
	return s
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) SetProvince(v string) *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions {
	s.Province = &v
	return s
}

type DescribeEnsRegionsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEnsRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEnsRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionsResponse) SetHeaders(v map[string]*string) *DescribeEnsRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsRegionsResponse) SetBody(v *DescribeEnsRegionsResponseBody) *DescribeEnsRegionsResponse {
	s.Body = v
	return s
}

type DescribeInstanceMonitorDataRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s DescribeInstanceMonitorDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataRequest) SetVersion(v string) *DescribeInstanceMonitorDataRequest {
	s.Version = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetInstanceId(v string) *DescribeInstanceMonitorDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetStartTime(v string) *DescribeInstanceMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetEndTime(v string) *DescribeInstanceMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetPeriod(v string) *DescribeInstanceMonitorDataRequest {
	s.Period = &v
	return s
}

type DescribeInstanceMonitorDataResponseBody struct {
	Code        *int32                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	MonitorData *DescribeInstanceMonitorDataResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Struct"`
	RequestId   *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceMonitorDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataResponseBody) SetCode(v int32) *DescribeInstanceMonitorDataResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBody) SetMonitorData(v *DescribeInstanceMonitorDataResponseBodyMonitorData) *DescribeInstanceMonitorDataResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBody) SetRequestId(v string) *DescribeInstanceMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceMonitorDataResponseBodyMonitorData struct {
	InstanceMonitorData []*DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData `json:"InstanceMonitorData,omitempty" xml:"InstanceMonitorData,omitempty" type:"Repeated"`
}

func (s DescribeInstanceMonitorDataResponseBodyMonitorData) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMonitorDataResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorData) SetInstanceMonitorData(v []*DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) *DescribeInstanceMonitorDataResponseBodyMonitorData {
	s.InstanceMonitorData = v
	return s
}

type DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData struct {
	CPU        *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Memory     *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetCPU(v string) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.CPU = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetInstanceId(v string) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetMemory(v string) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.Memory = &v
	return s
}

type DescribeInstanceMonitorDataResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceMonitorDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataResponse) SetHeaders(v map[string]*string) *DescribeInstanceMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceMonitorDataResponse) SetBody(v *DescribeInstanceMonitorDataResponseBody) *DescribeInstanceMonitorDataResponse {
	s.Body = v
	return s
}

type UnassociateEipAddressRequest struct {
	Version              *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsRegionId          *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	Eip                  *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	InstanceIdInternetIp *string `json:"InstanceIdInternetIp,omitempty" xml:"InstanceIdInternetIp,omitempty"`
}

func (s UnassociateEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s UnassociateEipAddressRequest) GoString() string {
	return s.String()
}

func (s *UnassociateEipAddressRequest) SetVersion(v string) *UnassociateEipAddressRequest {
	s.Version = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetEnsRegionId(v string) *UnassociateEipAddressRequest {
	s.EnsRegionId = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetEip(v string) *UnassociateEipAddressRequest {
	s.Eip = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetInstanceIdInternetIp(v string) *UnassociateEipAddressRequest {
	s.InstanceIdInternetIp = &v
	return s
}

type UnassociateEipAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassociateEipAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnassociateEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UnassociateEipAddressResponseBody) SetRequestId(v string) *UnassociateEipAddressResponseBody {
	s.RequestId = &v
	return s
}

type UnassociateEipAddressResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnassociateEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnassociateEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s UnassociateEipAddressResponse) GoString() string {
	return s.String()
}

func (s *UnassociateEipAddressResponse) SetHeaders(v map[string]*string) *UnassociateEipAddressResponse {
	s.Headers = v
	return s
}

func (s *UnassociateEipAddressResponse) SetBody(v *UnassociateEipAddressResponseBody) *UnassociateEipAddressResponse {
	s.Body = v
	return s
}

type DescribeDataPushResultRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DataNames    *string `json:"DataNames,omitempty" xml:"DataNames,omitempty"`
	DataVersions *string `json:"DataVersions,omitempty" xml:"DataVersions,omitempty"`
	MinDate      *string `json:"MinDate,omitempty" xml:"MinDate,omitempty"`
	MaxDate      *string `json:"MaxDate,omitempty" xml:"MaxDate,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionIds    *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
}

func (s DescribeDataPushResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultRequest) SetAppId(v string) *DescribeDataPushResultRequest {
	s.AppId = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetDataNames(v string) *DescribeDataPushResultRequest {
	s.DataNames = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetDataVersions(v string) *DescribeDataPushResultRequest {
	s.DataVersions = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetMinDate(v string) *DescribeDataPushResultRequest {
	s.MinDate = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetMaxDate(v string) *DescribeDataPushResultRequest {
	s.MaxDate = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetPageNumber(v int32) *DescribeDataPushResultRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetPageSize(v int32) *DescribeDataPushResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetRegionIds(v string) *DescribeDataPushResultRequest {
	s.RegionIds = &v
	return s
}

type DescribeDataPushResultResponseBody struct {
	PageNumber  *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PushResults *DescribeDataPushResultResponseBodyPushResults `json:"PushResults,omitempty" xml:"PushResults,omitempty" type:"Struct"`
	RequestId   *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataPushResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponseBody) SetPageNumber(v int32) *DescribeDataPushResultResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataPushResultResponseBody) SetPageSize(v int32) *DescribeDataPushResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataPushResultResponseBody) SetPushResults(v *DescribeDataPushResultResponseBodyPushResults) *DescribeDataPushResultResponseBody {
	s.PushResults = v
	return s
}

func (s *DescribeDataPushResultResponseBody) SetRequestId(v string) *DescribeDataPushResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataPushResultResponseBody) SetTotalCount(v int32) *DescribeDataPushResultResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDataPushResultResponseBodyPushResults struct {
	PushResult []*DescribeDataPushResultResponseBodyPushResultsPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" type:"Repeated"`
}

func (s DescribeDataPushResultResponseBodyPushResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponseBodyPushResults) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponseBodyPushResults) SetPushResult(v []*DescribeDataPushResultResponseBodyPushResultsPushResult) *DescribeDataPushResultResponseBodyPushResults {
	s.PushResult = v
	return s
}

type DescribeDataPushResultResponseBodyPushResultsPushResult struct {
	Name        *string                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	StatusStatS *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS `json:"StatusStatS,omitempty" xml:"StatusStatS,omitempty" type:"Struct"`
	Version     *string                                                             `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResult) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResult) SetName(v string) *DescribeDataPushResultResponseBodyPushResultsPushResult {
	s.Name = &v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResult) SetStatusStatS(v *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS) *DescribeDataPushResultResponseBodyPushResultsPushResult {
	s.StatusStatS = v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResult) SetVersion(v string) *DescribeDataPushResultResponseBodyPushResultsPushResult {
	s.Version = &v
	return s
}

type DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS struct {
	StatusStat []*DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat `json:"StatusStat,omitempty" xml:"StatusStat,omitempty" type:"Repeated"`
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS) SetStatusStat(v []*DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatS {
	s.StatusStat = v
	return s
}

type DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat struct {
	RegionIdCount *int32                                                                                 `json:"RegionIdCount,omitempty" xml:"RegionIdCount,omitempty"`
	RegionIds     *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Struct"`
	Status        *string                                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) SetRegionIdCount(v int32) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat {
	s.RegionIdCount = &v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) SetRegionIds(v *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat {
	s.RegionIds = v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat) SetStatus(v string) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStat {
	s.Status = &v
	return s
}

type DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds struct {
	RegionId []*DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId `json:"RegionId,omitempty" xml:"RegionId,omitempty" type:"Repeated"`
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds) SetRegionId(v []*DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIds {
	s.RegionId = v
	return s
}

type DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatusDescrip *string `json:"StatusDescrip,omitempty" xml:"StatusDescrip,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) SetRegionId(v string) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId {
	s.RegionId = &v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) SetStartTime(v string) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId {
	s.StartTime = &v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) SetStatusDescrip(v string) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId {
	s.StatusDescrip = &v
	return s
}

func (s *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) SetUpdateTime(v string) *DescribeDataPushResultResponseBodyPushResultsPushResultStatusStatSStatusStatRegionIdsRegionId {
	s.UpdateTime = &v
	return s
}

type DescribeDataPushResultResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataPushResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataPushResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponse) SetHeaders(v map[string]*string) *DescribeDataPushResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataPushResultResponse) SetBody(v *DescribeDataPushResultResponseBody) *DescribeDataPushResultResponse {
	s.Body = v
	return s
}

type LeaveSecurityGroupRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s LeaveSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s LeaveSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *LeaveSecurityGroupRequest) SetVersion(v string) *LeaveSecurityGroupRequest {
	s.Version = &v
	return s
}

func (s *LeaveSecurityGroupRequest) SetSecurityGroupId(v string) *LeaveSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *LeaveSecurityGroupRequest) SetInstanceId(v string) *LeaveSecurityGroupRequest {
	s.InstanceId = &v
	return s
}

type LeaveSecurityGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LeaveSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LeaveSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *LeaveSecurityGroupResponseBody) SetRequestId(v string) *LeaveSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

type LeaveSecurityGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LeaveSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LeaveSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s LeaveSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *LeaveSecurityGroupResponse) SetHeaders(v map[string]*string) *LeaveSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *LeaveSecurityGroupResponse) SetBody(v *LeaveSecurityGroupResponseBody) *LeaveSecurityGroupResponse {
	s.Body = v
	return s
}

type DescribeInstanceAutoRenewAttributeRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeInstanceAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetVersion(v string) *DescribeInstanceAutoRenewAttributeRequest {
	s.Version = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetInstanceIds(v string) *DescribeInstanceAutoRenewAttributeRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetOwnerId(v int64) *DescribeInstanceAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

type DescribeInstanceAutoRenewAttributeResponseBody struct {
	Code                    *int32                                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceRenewAttributes *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes `json:"InstanceRenewAttributes,omitempty" xml:"InstanceRenewAttributes,omitempty" type:"Struct"`
	RequestId               *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceAutoRenewAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) SetCode(v int32) *DescribeInstanceAutoRenewAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) SetInstanceRenewAttributes(v *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) *DescribeInstanceAutoRenewAttributeResponseBody {
	s.InstanceRenewAttributes = v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes struct {
	InstanceRenewAttribute []*DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute `json:"InstanceRenewAttribute,omitempty" xml:"InstanceRenewAttribute,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) SetInstanceRenewAttribute(v []*DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes {
	s.InstanceRenewAttribute = v
	return s
}

type DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute struct {
	AutoRenewal *bool   `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	Duration    *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) SetAutoRenewal(v bool) *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) SetDuration(v string) *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute {
	s.Duration = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) SetInstanceId(v string) *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute {
	s.InstanceId = &v
	return s
}

type DescribeInstanceAutoRenewAttributeResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceAutoRenewAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *DescribeInstanceAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponse) SetBody(v *DescribeInstanceAutoRenewAttributeResponseBody) *DescribeInstanceAutoRenewAttributeResponse {
	s.Body = v
	return s
}

type ImportKeyPairRequest struct {
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	KeyPairName   *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	PublicKeyBody *string `json:"PublicKeyBody,omitempty" xml:"PublicKeyBody,omitempty"`
}

func (s ImportKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairRequest) GoString() string {
	return s.String()
}

func (s *ImportKeyPairRequest) SetVersion(v string) *ImportKeyPairRequest {
	s.Version = &v
	return s
}

func (s *ImportKeyPairRequest) SetKeyPairName(v string) *ImportKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairRequest) SetPublicKeyBody(v string) *ImportKeyPairRequest {
	s.PublicKeyBody = &v
	return s
}

type ImportKeyPairResponseBody struct {
	KeyPairFingerPrint *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty"`
	KeyPairName        *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponseBody) SetKeyPairFingerPrint(v string) *ImportKeyPairResponseBody {
	s.KeyPairFingerPrint = &v
	return s
}

func (s *ImportKeyPairResponseBody) SetKeyPairName(v string) *ImportKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairResponseBody) SetRequestId(v string) *ImportKeyPairResponseBody {
	s.RequestId = &v
	return s
}

type ImportKeyPairResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairResponse) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponse) SetHeaders(v map[string]*string) *ImportKeyPairResponse {
	s.Headers = v
	return s
}

func (s *ImportKeyPairResponse) SetBody(v *ImportKeyPairResponseBody) *ImportKeyPairResponse {
	s.Body = v
	return s
}

type DeleteVmRequest struct {
	AliUid       *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	WorkloadUuid *string `json:"WorkloadUuid,omitempty" xml:"WorkloadUuid,omitempty"`
}

func (s DeleteVmRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVmRequest) GoString() string {
	return s.String()
}

func (s *DeleteVmRequest) SetAliUid(v int64) *DeleteVmRequest {
	s.AliUid = &v
	return s
}

func (s *DeleteVmRequest) SetWorkloadUuid(v string) *DeleteVmRequest {
	s.WorkloadUuid = &v
	return s
}

type DeleteVmResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Desc      *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVmResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVmResponseBody) SetCode(v int32) *DeleteVmResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVmResponseBody) SetMsg(v string) *DeleteVmResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteVmResponseBody) SetDesc(v string) *DeleteVmResponseBody {
	s.Desc = &v
	return s
}

func (s *DeleteVmResponseBody) SetData(v string) *DeleteVmResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteVmResponseBody) SetRequestId(v string) *DeleteVmResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVmResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVmResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVmResponse) GoString() string {
	return s.String()
}

func (s *DeleteVmResponse) SetHeaders(v map[string]*string) *DeleteVmResponse {
	s.Headers = v
	return s
}

func (s *DeleteVmResponse) SetBody(v *DeleteVmResponseBody) *DeleteVmResponse {
	s.Body = v
	return s
}

type CheckQuotaRequest struct {
	AliUid            *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ResourceAttribute *string `json:"ResourceAttribute,omitempty" xml:"ResourceAttribute,omitempty"`
	GroupUuid         *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
}

func (s CheckQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckQuotaRequest) GoString() string {
	return s.String()
}

func (s *CheckQuotaRequest) SetAliUid(v int64) *CheckQuotaRequest {
	s.AliUid = &v
	return s
}

func (s *CheckQuotaRequest) SetResourceAttribute(v string) *CheckQuotaRequest {
	s.ResourceAttribute = &v
	return s
}

func (s *CheckQuotaRequest) SetGroupUuid(v string) *CheckQuotaRequest {
	s.GroupUuid = &v
	return s
}

type CheckQuotaResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Desc      *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s CheckQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *CheckQuotaResponseBody) SetRequestId(v string) *CheckQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckQuotaResponseBody) SetCode(v int32) *CheckQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *CheckQuotaResponseBody) SetData(v string) *CheckQuotaResponseBody {
	s.Data = &v
	return s
}

func (s *CheckQuotaResponseBody) SetMsg(v string) *CheckQuotaResponseBody {
	s.Msg = &v
	return s
}

func (s *CheckQuotaResponseBody) SetDesc(v string) *CheckQuotaResponseBody {
	s.Desc = &v
	return s
}

type CheckQuotaResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckQuotaResponse) GoString() string {
	return s.String()
}

func (s *CheckQuotaResponse) SetHeaders(v map[string]*string) *CheckQuotaResponse {
	s.Headers = v
	return s
}

func (s *CheckQuotaResponse) SetBody(v *CheckQuotaResponseBody) *CheckQuotaResponse {
	s.Body = v
	return s
}

type CreateImageRequest struct {
	Product                *string `json:"product,omitempty" xml:"product,omitempty"`
	Version                *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ImageName              *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	DeleteAfterImageUpload *string `json:"DeleteAfterImageUpload,omitempty" xml:"DeleteAfterImageUpload,omitempty"`
}

func (s CreateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) SetProduct(v string) *CreateImageRequest {
	s.Product = &v
	return s
}

func (s *CreateImageRequest) SetVersion(v string) *CreateImageRequest {
	s.Version = &v
	return s
}

func (s *CreateImageRequest) SetInstanceId(v string) *CreateImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateImageRequest) SetImageName(v string) *CreateImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateImageRequest) SetDeleteAfterImageUpload(v string) *CreateImageRequest {
	s.DeleteAfterImageUpload = &v
	return s
}

type CreateImageResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageResponseBody) SetCode(v int32) *CreateImageResponseBody {
	s.Code = &v
	return s
}

func (s *CreateImageResponseBody) SetRequestId(v string) *CreateImageResponseBody {
	s.RequestId = &v
	return s
}

type CreateImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponse) GoString() string {
	return s.String()
}

func (s *CreateImageResponse) SetHeaders(v map[string]*string) *CreateImageResponse {
	s.Headers = v
	return s
}

func (s *CreateImageResponse) SetBody(v *CreateImageResponseBody) *CreateImageResponse {
	s.Body = v
	return s
}

type CreateEipInstanceRequest struct {
	// ENS节点ID
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// EIP的带宽峰值
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// EIP的计费方式，取值：  PrePaid：包年包月。 PostPaid（默认值）：按量计费。 当InstanceChargeType取值为PostPaid时，InternetChargeType不能为PayByBandwidth
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// EIP的计量方式，取值：  PayByBandwidth（默认值）：按带宽计费。 取值：95BandwidthByMonth：月95。
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// EIP实例名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 运营商信息
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s CreateEipInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEipInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateEipInstanceRequest) SetEnsRegionId(v string) *CreateEipInstanceRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateEipInstanceRequest) SetBandwidth(v int64) *CreateEipInstanceRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateEipInstanceRequest) SetInstanceChargeType(v string) *CreateEipInstanceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateEipInstanceRequest) SetInternetChargeType(v string) *CreateEipInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateEipInstanceRequest) SetName(v string) *CreateEipInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateEipInstanceRequest) SetIsp(v string) *CreateEipInstanceRequest {
	s.Isp = &v
	return s
}

type CreateEipInstanceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// EIP的ID。
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
}

func (s CreateEipInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEipInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEipInstanceResponseBody) SetRequestId(v string) *CreateEipInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEipInstanceResponseBody) SetAllocationId(v string) *CreateEipInstanceResponseBody {
	s.AllocationId = &v
	return s
}

type CreateEipInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEipInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEipInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEipInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateEipInstanceResponse) SetHeaders(v map[string]*string) *CreateEipInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateEipInstanceResponse) SetBody(v *CreateEipInstanceResponseBody) *CreateEipInstanceResponse {
	s.Body = v
	return s
}

type DeleteEpnInstanceRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
}

func (s DeleteEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteEpnInstanceRequest) SetEPNInstanceId(v string) *DeleteEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

type DeleteEpnInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEpnInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEpnInstanceResponseBody) SetRequestId(v string) *DeleteEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEpnInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteEpnInstanceResponse) SetHeaders(v map[string]*string) *DeleteEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteEpnInstanceResponse) SetBody(v *DeleteEpnInstanceResponseBody) *DeleteEpnInstanceResponse {
	s.Body = v
	return s
}

type ModifyImageAttributeRequest struct {
	Product   *string `json:"product,omitempty" xml:"product,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
}

func (s ModifyImageAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeRequest) SetProduct(v string) *ModifyImageAttributeRequest {
	s.Product = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetVersion(v string) *ModifyImageAttributeRequest {
	s.Version = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageId(v string) *ModifyImageAttributeRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageName(v string) *ModifyImageAttributeRequest {
	s.ImageName = &v
	return s
}

type ModifyImageAttributeResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeResponseBody) SetCode(v int32) *ModifyImageAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyImageAttributeResponseBody) SetRequestId(v string) *ModifyImageAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyImageAttributeResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyImageAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyImageAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeResponse) SetHeaders(v map[string]*string) *ModifyImageAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageAttributeResponse) SetBody(v *ModifyImageAttributeResponseBody) *ModifyImageAttributeResponse {
	s.Body = v
	return s
}

type DescribeBandwitdhByInternetChargeTypeRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Isp         *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
}

func (s DescribeBandwitdhByInternetChargeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandwitdhByInternetChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) SetVersion(v string) *DescribeBandwitdhByInternetChargeTypeRequest {
	s.Version = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) SetStartTime(v string) *DescribeBandwitdhByInternetChargeTypeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) SetEndTime(v string) *DescribeBandwitdhByInternetChargeTypeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) SetIsp(v string) *DescribeBandwitdhByInternetChargeTypeRequest {
	s.Isp = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) SetEnsRegionId(v string) *DescribeBandwitdhByInternetChargeTypeRequest {
	s.EnsRegionId = &v
	return s
}

type DescribeBandwitdhByInternetChargeTypeResponseBody struct {
	BandwidthValue     *int64  `json:"BandwidthValue,omitempty" xml:"BandwidthValue,omitempty"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeStamp          *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeBandwitdhByInternetChargeTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandwitdhByInternetChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBandwitdhByInternetChargeTypeResponseBody) SetBandwidthValue(v int64) *DescribeBandwitdhByInternetChargeTypeResponseBody {
	s.BandwidthValue = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeResponseBody) SetInternetChargeType(v string) *DescribeBandwitdhByInternetChargeTypeResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeResponseBody) SetRequestId(v string) *DescribeBandwitdhByInternetChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeResponseBody) SetTimeStamp(v string) *DescribeBandwitdhByInternetChargeTypeResponseBody {
	s.TimeStamp = &v
	return s
}

type DescribeBandwitdhByInternetChargeTypeResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBandwitdhByInternetChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBandwitdhByInternetChargeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandwitdhByInternetChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeBandwitdhByInternetChargeTypeResponse) SetHeaders(v map[string]*string) *DescribeBandwitdhByInternetChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeResponse) SetBody(v *DescribeBandwitdhByInternetChargeTypeResponseBody) *DescribeBandwitdhByInternetChargeTypeResponse {
	s.Body = v
	return s
}

type CreateSecurityGroupRequest struct {
	Version           *string `json:"Version,omitempty" xml:"Version,omitempty"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupRequest) SetVersion(v string) *CreateSecurityGroupRequest {
	s.Version = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetSecurityGroupName(v string) *CreateSecurityGroupRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetDescription(v string) *CreateSecurityGroupRequest {
	s.Description = &v
	return s
}

type CreateSecurityGroupResponseBody struct {
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupResponseBody) SetSecurityGroupId(v string) *CreateSecurityGroupResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateSecurityGroupResponseBody) SetRequestId(v string) *CreateSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateSecurityGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupResponse) SetHeaders(v map[string]*string) *CreateSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityGroupResponse) SetBody(v *CreateSecurityGroupResponseBody) *CreateSecurityGroupResponse {
	s.Body = v
	return s
}

type DescribeEpnInstancesRequest struct {
	EPNInstanceId   *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	EPNInstanceName *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeEpnInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstancesRequest) SetEPNInstanceId(v string) *DescribeEpnInstancesRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *DescribeEpnInstancesRequest) SetEPNInstanceName(v string) *DescribeEpnInstancesRequest {
	s.EPNInstanceName = &v
	return s
}

func (s *DescribeEpnInstancesRequest) SetPageNumber(v int32) *DescribeEpnInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEpnInstancesRequest) SetPageSize(v int32) *DescribeEpnInstancesRequest {
	s.PageSize = &v
	return s
}

type DescribeEpnInstancesResponseBody struct {
	EPNInstances *DescribeEpnInstancesResponseBodyEPNInstances `json:"EPNInstances,omitempty" xml:"EPNInstances,omitempty" type:"Struct"`
	PageNumber   *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEpnInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstancesResponseBody) SetEPNInstances(v *DescribeEpnInstancesResponseBodyEPNInstances) *DescribeEpnInstancesResponseBody {
	s.EPNInstances = v
	return s
}

func (s *DescribeEpnInstancesResponseBody) SetPageNumber(v int32) *DescribeEpnInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEpnInstancesResponseBody) SetPageSize(v int32) *DescribeEpnInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEpnInstancesResponseBody) SetRequestId(v string) *DescribeEpnInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEpnInstancesResponseBody) SetTotalCount(v int32) *DescribeEpnInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEpnInstancesResponseBodyEPNInstances struct {
	EPNInstance []*DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance `json:"EPNInstance,omitempty" xml:"EPNInstance,omitempty" type:"Repeated"`
}

func (s DescribeEpnInstancesResponseBodyEPNInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstancesResponseBodyEPNInstances) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstancesResponseBodyEPNInstances) SetEPNInstance(v []*DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) *DescribeEpnInstancesResponseBodyEPNInstances {
	s.EPNInstance = v
	return s
}

type DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance struct {
	CreationTime            *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	EPNInstanceId           *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	EPNInstanceName         *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	EPNInstanceType         *string `json:"EPNInstanceType,omitempty" xml:"EPNInstanceType,omitempty"`
	EndTime                 *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InternetMaxBandwidthOut *int32  `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	ModifyTime              *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	NetworkingModel         *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
	StartTime               *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetCreationTime(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetEPNInstanceId(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.EPNInstanceId = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetEPNInstanceName(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.EPNInstanceName = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetEPNInstanceType(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.EPNInstanceType = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetEndTime(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.EndTime = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetInternetMaxBandwidthOut(v int32) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetModifyTime(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.ModifyTime = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetNetworkingModel(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.NetworkingModel = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetStartTime(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.StartTime = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetStatus(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.Status = &v
	return s
}

type DescribeEpnInstancesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEpnInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEpnInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstancesResponse) SetHeaders(v map[string]*string) *DescribeEpnInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEpnInstancesResponse) SetBody(v *DescribeEpnInstancesResponseBody) *DescribeEpnInstancesResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	SystemDisk         *CreateInstanceRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	DataDisk           []*CreateInstanceRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	InstanceType       *string                          `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	EnsRegionId        *string                          `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	Password           *string                          `json:"Password,omitempty" xml:"Password,omitempty"`
	Period             *string                          `json:"Period,omitempty" xml:"Period,omitempty"`
	ImageId            *string                          `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Quantity           *string                          `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	InternetChargeType *string                          `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	AutoRenewPeriod    *string                          `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	AutoRenew          *string                          `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	IpType             *string                          `json:"IpType,omitempty" xml:"IpType,omitempty"`
	OwnerId            *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	KeyPairName        *string                          `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	UserData           *string                          `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VSwitchId          *string                          `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	PrivateIpAddress   *string                          `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	PaymentType        *string                          `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	InstanceName       *string                          `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	HostName           *string                          `json:"HostName,omitempty" xml:"HostName,omitempty"`
	UniqueSuffix       *bool                            `json:"UniqueSuffix,omitempty" xml:"UniqueSuffix,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetSystemDisk(v *CreateInstanceRequestSystemDisk) *CreateInstanceRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateInstanceRequest) SetDataDisk(v []*CreateInstanceRequestDataDisk) *CreateInstanceRequest {
	s.DataDisk = v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetEnsRegionId(v string) *CreateInstanceRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateInstanceRequest) SetPassword(v string) *CreateInstanceRequest {
	s.Password = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v string) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetImageId(v string) *CreateInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateInstanceRequest) SetQuantity(v string) *CreateInstanceRequest {
	s.Quantity = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetChargeType(v string) *CreateInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v string) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v string) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetIpType(v string) *CreateInstanceRequest {
	s.IpType = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerId(v int64) *CreateInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetKeyPairName(v string) *CreateInstanceRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateInstanceRequest) SetUserData(v string) *CreateInstanceRequest {
	s.UserData = &v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchId(v string) *CreateInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequest) SetPrivateIpAddress(v string) *CreateInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentType(v string) *CreateInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetHostName(v string) *CreateInstanceRequest {
	s.HostName = &v
	return s
}

func (s *CreateInstanceRequest) SetUniqueSuffix(v bool) *CreateInstanceRequest {
	s.UniqueSuffix = &v
	return s
}

type CreateInstanceRequestSystemDisk struct {
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateInstanceRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestSystemDisk) SetSize(v string) *CreateInstanceRequestSystemDisk {
	s.Size = &v
	return s
}

type CreateInstanceRequestDataDisk struct {
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateInstanceRequestDataDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestDataDisk) SetSize(v string) *CreateInstanceRequestDataDisk {
	s.Size = &v
	return s
}

type CreateInstanceResponseBody struct {
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceIds *CreateInstanceResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	Code        *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceIds(v *CreateInstanceResponseBodyInstanceIds) *CreateInstanceResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateInstanceResponseBody) SetCode(v int32) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

type CreateInstanceResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s CreateInstanceResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyInstanceIds) SetInstanceId(v []*string) *CreateInstanceResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

type CreateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type DescribeDataDistResultRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DataNames    *string `json:"DataNames,omitempty" xml:"DataNames,omitempty"`
	DataVersions *string `json:"DataVersions,omitempty" xml:"DataVersions,omitempty"`
	InstanceIds  *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	MinDate      *string `json:"MinDate,omitempty" xml:"MinDate,omitempty"`
	MaxDate      *string `json:"MaxDate,omitempty" xml:"MaxDate,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDataDistResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultRequest) SetAppId(v string) *DescribeDataDistResultRequest {
	s.AppId = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetDataNames(v string) *DescribeDataDistResultRequest {
	s.DataNames = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetDataVersions(v string) *DescribeDataDistResultRequest {
	s.DataVersions = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetInstanceIds(v string) *DescribeDataDistResultRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetMinDate(v string) *DescribeDataDistResultRequest {
	s.MinDate = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetMaxDate(v string) *DescribeDataDistResultRequest {
	s.MaxDate = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetPageNumber(v int32) *DescribeDataDistResultRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetPageSize(v int32) *DescribeDataDistResultRequest {
	s.PageSize = &v
	return s
}

type DescribeDataDistResultResponseBody struct {
	DistResults *DescribeDataDistResultResponseBodyDistResults `json:"DistResults,omitempty" xml:"DistResults,omitempty" type:"Struct"`
	PageNumber  *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataDistResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseBody) SetDistResults(v *DescribeDataDistResultResponseBodyDistResults) *DescribeDataDistResultResponseBody {
	s.DistResults = v
	return s
}

func (s *DescribeDataDistResultResponseBody) SetPageNumber(v int32) *DescribeDataDistResultResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataDistResultResponseBody) SetPageSize(v int32) *DescribeDataDistResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataDistResultResponseBody) SetRequestId(v string) *DescribeDataDistResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataDistResultResponseBody) SetTotalCount(v int32) *DescribeDataDistResultResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDataDistResultResponseBodyDistResults struct {
	DistResult []*DescribeDataDistResultResponseBodyDistResultsDistResult `json:"DistResult,omitempty" xml:"DistResult,omitempty" type:"Repeated"`
}

func (s DescribeDataDistResultResponseBodyDistResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponseBodyDistResults) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseBodyDistResults) SetDistResult(v []*DescribeDataDistResultResponseBodyDistResultsDistResult) *DescribeDataDistResultResponseBodyDistResults {
	s.DistResult = v
	return s
}

type DescribeDataDistResultResponseBodyDistResultsDistResult struct {
	Name        *string                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	StatusStats *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats `json:"StatusStats,omitempty" xml:"StatusStats,omitempty" type:"Struct"`
	Version     *string                                                             `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResult) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResult) SetName(v string) *DescribeDataDistResultResponseBodyDistResultsDistResult {
	s.Name = &v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResult) SetStatusStats(v *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats) *DescribeDataDistResultResponseBodyDistResultsDistResult {
	s.StatusStats = v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResult) SetVersion(v string) *DescribeDataDistResultResponseBodyDistResultsDistResult {
	s.Version = &v
	return s
}

type DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats struct {
	StatusStat []*DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat `json:"StatusStat,omitempty" xml:"StatusStat,omitempty" type:"Repeated"`
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats) SetStatusStat(v []*DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats {
	s.StatusStat = v
	return s
}

type DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat struct {
	InstanceCount *string                                                                                `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	Instances     *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	Status        *string                                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) SetInstanceCount(v string) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat {
	s.InstanceCount = &v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) SetInstances(v *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat {
	s.Instances = v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) SetStatus(v string) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat {
	s.Status = &v
	return s
}

type DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances struct {
	Instance []*DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances) SetInstance(v []*DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances {
	s.Instance = v
	return s
}

type DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatusDescrip *string `json:"StatusDescrip,omitempty" xml:"StatusDescrip,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) SetInstanceId(v string) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) SetStartTime(v string) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance {
	s.StartTime = &v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) SetStatusDescrip(v string) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance {
	s.StatusDescrip = &v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) SetUpdateTime(v string) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance {
	s.UpdateTime = &v
	return s
}

type DescribeDataDistResultResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataDistResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataDistResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponse) SetHeaders(v map[string]*string) *DescribeDataDistResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataDistResultResponse) SetBody(v *DescribeDataDistResultResponseBody) *DescribeDataDistResultResponse {
	s.Body = v
	return s
}

type ReleasePostPaidInstanceRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleasePostPaidInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleasePostPaidInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleasePostPaidInstanceRequest) SetVersion(v string) *ReleasePostPaidInstanceRequest {
	s.Version = &v
	return s
}

func (s *ReleasePostPaidInstanceRequest) SetInstanceId(v string) *ReleasePostPaidInstanceRequest {
	s.InstanceId = &v
	return s
}

type ReleasePostPaidInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleasePostPaidInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleasePostPaidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePostPaidInstanceResponseBody) SetRequestId(v string) *ReleasePostPaidInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ReleasePostPaidInstanceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleasePostPaidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleasePostPaidInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleasePostPaidInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleasePostPaidInstanceResponse) SetHeaders(v map[string]*string) *ReleasePostPaidInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleasePostPaidInstanceResponse) SetBody(v *ReleasePostPaidInstanceResponseBody) *ReleasePostPaidInstanceResponse {
	s.Body = v
	return s
}

type CreateEPInstanceRequest struct {
	EPNInstanceType         *string `json:"EPNInstanceType,omitempty" xml:"EPNInstanceType,omitempty"`
	EPNInstanceName         *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	InternetChargeType      *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	NetworkingModel         *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
	InternetMaxBandwidthOut *int32  `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
}

func (s CreateEPInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEPInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateEPInstanceRequest) SetEPNInstanceType(v string) *CreateEPInstanceRequest {
	s.EPNInstanceType = &v
	return s
}

func (s *CreateEPInstanceRequest) SetEPNInstanceName(v string) *CreateEPInstanceRequest {
	s.EPNInstanceName = &v
	return s
}

func (s *CreateEPInstanceRequest) SetInternetChargeType(v string) *CreateEPInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateEPInstanceRequest) SetNetworkingModel(v string) *CreateEPInstanceRequest {
	s.NetworkingModel = &v
	return s
}

func (s *CreateEPInstanceRequest) SetInternetMaxBandwidthOut(v int32) *CreateEPInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

type CreateEPInstanceResponseBody struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEPInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEPInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEPInstanceResponseBody) SetEPNInstanceId(v string) *CreateEPInstanceResponseBody {
	s.EPNInstanceId = &v
	return s
}

func (s *CreateEPInstanceResponseBody) SetRequestId(v string) *CreateEPInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateEPInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEPInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEPInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEPInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateEPInstanceResponse) SetHeaders(v map[string]*string) *CreateEPInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateEPInstanceResponse) SetBody(v *CreateEPInstanceResponseBody) *CreateEPInstanceResponse {
	s.Body = v
	return s
}

type DescribeEpnMeasurementDataRequest struct {
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s DescribeEpnMeasurementDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnMeasurementDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataRequest) SetVersion(v string) *DescribeEpnMeasurementDataRequest {
	s.Version = &v
	return s
}

func (s *DescribeEpnMeasurementDataRequest) SetStartDate(v string) *DescribeEpnMeasurementDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeEpnMeasurementDataRequest) SetEndDate(v string) *DescribeEpnMeasurementDataRequest {
	s.EndDate = &v
	return s
}

type DescribeEpnMeasurementDataResponseBody struct {
	MeasurementDatas *DescribeEpnMeasurementDataResponseBodyMeasurementDatas `json:"MeasurementDatas,omitempty" xml:"MeasurementDatas,omitempty" type:"Struct"`
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEpnMeasurementDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponseBody) SetMeasurementDatas(v *DescribeEpnMeasurementDataResponseBodyMeasurementDatas) *DescribeEpnMeasurementDataResponseBody {
	s.MeasurementDatas = v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBody) SetRequestId(v string) *DescribeEpnMeasurementDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEpnMeasurementDataResponseBodyMeasurementDatas struct {
	MeasurementData []*DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData `json:"MeasurementData,omitempty" xml:"MeasurementData,omitempty" type:"Repeated"`
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatas) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatas) SetMeasurementData(v []*DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) *DescribeEpnMeasurementDataResponseBodyMeasurementDatas {
	s.MeasurementData = v
	return s
}

type DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData struct {
	BandWidthFeeDatas *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas `json:"BandWidthFeeDatas,omitempty" xml:"BandWidthFeeDatas,omitempty" type:"Struct"`
	ChargeModel       *string                                                                                 `json:"ChargeModel,omitempty" xml:"ChargeModel,omitempty"`
	CostCycle         *string                                                                                 `json:"CostCycle,omitempty" xml:"CostCycle,omitempty"`
	CostEndTime       *string                                                                                 `json:"CostEndTime,omitempty" xml:"CostEndTime,omitempty"`
	CostStartTime     *string                                                                                 `json:"CostStartTime,omitempty" xml:"CostStartTime,omitempty"`
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetBandWidthFeeDatas(v *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.BandWidthFeeDatas = v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetChargeModel(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.ChargeModel = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetCostCycle(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.CostCycle = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetCostEndTime(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.CostEndTime = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetCostStartTime(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.CostStartTime = &v
	return s
}

type DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas struct {
	BandWidthFeeData []*DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData `json:"BandWidthFeeData,omitempty" xml:"BandWidthFeeData,omitempty" type:"Repeated"`
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) SetBandWidthFeeData(v []*DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas {
	s.BandWidthFeeData = v
	return s
}

type DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData struct {
	CostCode *string `json:"CostCode,omitempty" xml:"CostCode,omitempty"`
	CostName *string `json:"CostName,omitempty" xml:"CostName,omitempty"`
	CostType *string `json:"CostType,omitempty" xml:"CostType,omitempty"`
	CostVal  *int32  `json:"CostVal,omitempty" xml:"CostVal,omitempty"`
	IspLine  *string `json:"IspLine,omitempty" xml:"IspLine,omitempty"`
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostCode(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostCode = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostName(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostName = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostType(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostType = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostVal(v int32) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostVal = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetIspLine(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.IspLine = &v
	return s
}

type DescribeEpnMeasurementDataResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEpnMeasurementDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEpnMeasurementDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponse) SetHeaders(v map[string]*string) *DescribeEpnMeasurementDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeEpnMeasurementDataResponse) SetBody(v *DescribeEpnMeasurementDataResponseBody) *DescribeEpnMeasurementDataResponse {
	s.Body = v
	return s
}

type DescribeApplicationResourceSummaryRequest struct {
	Level        *string `json:"Level,omitempty" xml:"Level,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeApplicationResourceSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationResourceSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationResourceSummaryRequest) SetLevel(v string) *DescribeApplicationResourceSummaryRequest {
	s.Level = &v
	return s
}

func (s *DescribeApplicationResourceSummaryRequest) SetResourceType(v string) *DescribeApplicationResourceSummaryRequest {
	s.ResourceType = &v
	return s
}

type DescribeApplicationResourceSummaryResponseBody struct {
	ApplicationResource *string `json:"ApplicationResource,omitempty" xml:"ApplicationResource,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApplicationResourceSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationResourceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationResourceSummaryResponseBody) SetApplicationResource(v string) *DescribeApplicationResourceSummaryResponseBody {
	s.ApplicationResource = &v
	return s
}

func (s *DescribeApplicationResourceSummaryResponseBody) SetRequestId(v string) *DescribeApplicationResourceSummaryResponseBody {
	s.RequestId = &v
	return s
}

type DescribeApplicationResourceSummaryResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApplicationResourceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApplicationResourceSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationResourceSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationResourceSummaryResponse) SetHeaders(v map[string]*string) *DescribeApplicationResourceSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationResourceSummaryResponse) SetBody(v *DescribeApplicationResourceSummaryResponseBody) *DescribeApplicationResourceSummaryResponse {
	s.Body = v
	return s
}

type RollbackApplicationRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	FromAppVersion *string `json:"FromAppVersion,omitempty" xml:"FromAppVersion,omitempty"`
	ToAppVersion   *string `json:"ToAppVersion,omitempty" xml:"ToAppVersion,omitempty"`
	Timeout        *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s RollbackApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackApplicationRequest) GoString() string {
	return s.String()
}

func (s *RollbackApplicationRequest) SetAppId(v string) *RollbackApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RollbackApplicationRequest) SetFromAppVersion(v string) *RollbackApplicationRequest {
	s.FromAppVersion = &v
	return s
}

func (s *RollbackApplicationRequest) SetToAppVersion(v string) *RollbackApplicationRequest {
	s.ToAppVersion = &v
	return s
}

func (s *RollbackApplicationRequest) SetTimeout(v int32) *RollbackApplicationRequest {
	s.Timeout = &v
	return s
}

type RollbackApplicationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackApplicationResponseBody) SetRequestId(v string) *RollbackApplicationResponseBody {
	s.RequestId = &v
	return s
}

type RollbackApplicationResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RollbackApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackApplicationResponse) GoString() string {
	return s.String()
}

func (s *RollbackApplicationResponse) SetHeaders(v map[string]*string) *RollbackApplicationResponse {
	s.Headers = v
	return s
}

func (s *RollbackApplicationResponse) SetBody(v *RollbackApplicationResponseBody) *RollbackApplicationResponse {
	s.Body = v
	return s
}

type ExportMeasurementDataRequest struct {
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s ExportMeasurementDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportMeasurementDataRequest) GoString() string {
	return s.String()
}

func (s *ExportMeasurementDataRequest) SetVersion(v string) *ExportMeasurementDataRequest {
	s.Version = &v
	return s
}

func (s *ExportMeasurementDataRequest) SetStartDate(v string) *ExportMeasurementDataRequest {
	s.StartDate = &v
	return s
}

func (s *ExportMeasurementDataRequest) SetEndDate(v string) *ExportMeasurementDataRequest {
	s.EndDate = &v
	return s
}

type ExportMeasurementDataResponseBody struct {
	FilePath  *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportMeasurementDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportMeasurementDataResponseBody) GoString() string {
	return s.String()
}

func (s *ExportMeasurementDataResponseBody) SetFilePath(v string) *ExportMeasurementDataResponseBody {
	s.FilePath = &v
	return s
}

func (s *ExportMeasurementDataResponseBody) SetRequestId(v string) *ExportMeasurementDataResponseBody {
	s.RequestId = &v
	return s
}

type ExportMeasurementDataResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExportMeasurementDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExportMeasurementDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportMeasurementDataResponse) GoString() string {
	return s.String()
}

func (s *ExportMeasurementDataResponse) SetHeaders(v map[string]*string) *ExportMeasurementDataResponse {
	s.Headers = v
	return s
}

func (s *ExportMeasurementDataResponse) SetBody(v *ExportMeasurementDataResponseBody) *ExportMeasurementDataResponse {
	s.Body = v
	return s
}

type DeleteApplicationRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Timeout *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s DeleteApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) SetAppId(v string) *DeleteApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DeleteApplicationRequest) SetTimeout(v int32) *DeleteApplicationRequest {
	s.Timeout = &v
	return s
}

type DeleteApplicationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponseBody) SetRequestId(v string) *DeleteApplicationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApplicationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponse) SetHeaders(v map[string]*string) *DeleteApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationResponse) SetBody(v *DeleteApplicationResponseBody) *DeleteApplicationResponse {
	s.Body = v
	return s
}

type DescribeKeyPairsRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	PageNumber  *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeKeyPairsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsRequest) SetVersion(v string) *DescribeKeyPairsRequest {
	s.Version = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetKeyPairName(v string) *DescribeKeyPairsRequest {
	s.KeyPairName = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetPageNumber(v string) *DescribeKeyPairsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetPageSize(v string) *DescribeKeyPairsRequest {
	s.PageSize = &v
	return s
}

type DescribeKeyPairsResponseBody struct {
	KeyPairs   *DescribeKeyPairsResponseBodyKeyPairs `json:"KeyPairs,omitempty" xml:"KeyPairs,omitempty" type:"Struct"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeKeyPairsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyPairsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBody) SetKeyPairs(v *DescribeKeyPairsResponseBodyKeyPairs) *DescribeKeyPairsResponseBody {
	s.KeyPairs = v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetPageNumber(v int32) *DescribeKeyPairsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetPageSize(v int32) *DescribeKeyPairsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetRequestId(v string) *DescribeKeyPairsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetTotalCount(v int32) *DescribeKeyPairsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeKeyPairsResponseBodyKeyPairs struct {
	KeyPair []*DescribeKeyPairsResponseBodyKeyPairsKeyPair `json:"KeyPair,omitempty" xml:"KeyPair,omitempty" type:"Repeated"`
}

func (s DescribeKeyPairsResponseBodyKeyPairs) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyPairsResponseBodyKeyPairs) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBodyKeyPairs) SetKeyPair(v []*DescribeKeyPairsResponseBodyKeyPairsKeyPair) *DescribeKeyPairsResponseBodyKeyPairs {
	s.KeyPair = v
	return s
}

type DescribeKeyPairsResponseBodyKeyPairsKeyPair struct {
	CreationTime       *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	KeyPairFingerPrint *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty"`
	KeyPairName        *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
}

func (s DescribeKeyPairsResponseBodyKeyPairsKeyPair) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyPairsResponseBodyKeyPairsKeyPair) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) SetCreationTime(v string) *DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	s.CreationTime = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) SetKeyPairFingerPrint(v string) *DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	s.KeyPairFingerPrint = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) SetKeyPairName(v string) *DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	s.KeyPairName = &v
	return s
}

type DescribeKeyPairsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeKeyPairsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeKeyPairsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyPairsResponse) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponse) SetHeaders(v map[string]*string) *DescribeKeyPairsResponse {
	s.Headers = v
	return s
}

func (s *DescribeKeyPairsResponse) SetBody(v *DescribeKeyPairsResponseBody) *DescribeKeyPairsResponse {
	s.Body = v
	return s
}

type ModifyEnsEipAddressAttributeRequest struct {
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Bandwidth    *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
}

func (s ModifyEnsEipAddressAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyEnsEipAddressAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyEnsEipAddressAttributeRequest) SetAllocationId(v string) *ModifyEnsEipAddressAttributeRequest {
	s.AllocationId = &v
	return s
}

func (s *ModifyEnsEipAddressAttributeRequest) SetName(v string) *ModifyEnsEipAddressAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyEnsEipAddressAttributeRequest) SetDescription(v string) *ModifyEnsEipAddressAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyEnsEipAddressAttributeRequest) SetBandwidth(v int32) *ModifyEnsEipAddressAttributeRequest {
	s.Bandwidth = &v
	return s
}

type ModifyEnsEipAddressAttributeResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEnsEipAddressAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyEnsEipAddressAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEnsEipAddressAttributeResponseBody) SetRequestId(v string) *ModifyEnsEipAddressAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyEnsEipAddressAttributeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyEnsEipAddressAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyEnsEipAddressAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyEnsEipAddressAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyEnsEipAddressAttributeResponse) SetHeaders(v map[string]*string) *ModifyEnsEipAddressAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyEnsEipAddressAttributeResponse) SetBody(v *ModifyEnsEipAddressAttributeResponseBody) *ModifyEnsEipAddressAttributeResponse {
	s.Body = v
	return s
}

type DescribeEnsResourceUsageRequest struct {
	// vm实例使用结束时间查询开始范围，格式： yyyy-MM-dd或yyyy-MM-dd HH:mm:ss
	ExpiredStartTime *string `json:"ExpiredStartTime,omitempty" xml:"ExpiredStartTime,omitempty"`
	// vm实例使用结束时间查询结束范围，，格式： yyyy-MM-dd或yyyy-MM-dd HH:mm:ss
	ExpiredEndTime *string `json:"ExpiredEndTime,omitempty" xml:"ExpiredEndTime,omitempty"`
}

func (s DescribeEnsResourceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsResourceUsageRequest) SetExpiredStartTime(v string) *DescribeEnsResourceUsageRequest {
	s.ExpiredStartTime = &v
	return s
}

func (s *DescribeEnsResourceUsageRequest) SetExpiredEndTime(v string) *DescribeEnsResourceUsageRequest {
	s.ExpiredEndTime = &v
	return s
}

type DescribeEnsResourceUsageResponseBody struct {
	// Id of the request
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EnsResourceUsage []*DescribeEnsResourceUsageResponseBodyEnsResourceUsage `json:"EnsResourceUsage,omitempty" xml:"EnsResourceUsage,omitempty" type:"Repeated"`
}

func (s DescribeEnsResourceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsResourceUsageResponseBody) SetRequestId(v string) *DescribeEnsResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBody) SetEnsResourceUsage(v []*DescribeEnsResourceUsageResponseBodyEnsResourceUsage) *DescribeEnsResourceUsageResponseBody {
	s.EnsResourceUsage = v
	return s
}

type DescribeEnsResourceUsageResponseBodyEnsResourceUsage struct {
	ServiceType          *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	InstanceCount        *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	CpuSum               *int64  `json:"CpuSum,omitempty" xml:"CpuSum,omitempty"`
	GpuSum               *int64  `json:"GpuSum,omitempty" xml:"GpuSum,omitempty"`
	DownCount            *int32  `json:"DownCount,omitempty" xml:"DownCount,omitempty"`
	ExpiredCount         *int32  `json:"ExpiredCount,omitempty" xml:"ExpiredCount,omitempty"`
	ExpiringCount        *int32  `json:"ExpiringCount,omitempty" xml:"ExpiringCount,omitempty"`
	RunningCount         *int32  `json:"RunningCount,omitempty" xml:"RunningCount,omitempty"`
	DiskCount            *int32  `json:"DiskCount,omitempty" xml:"DiskCount,omitempty"`
	StorageSum           *int64  `json:"StorageSum,omitempty" xml:"StorageSum,omitempty"`
	ComputeResourceCount *int32  `json:"ComputeResourceCount,omitempty" xml:"ComputeResourceCount,omitempty"`
}

func (s DescribeEnsResourceUsageResponseBodyEnsResourceUsage) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsResourceUsageResponseBodyEnsResourceUsage) GoString() string {
	return s.String()
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetServiceType(v string) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.ServiceType = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetInstanceCount(v int32) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.InstanceCount = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetCpuSum(v int64) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.CpuSum = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetGpuSum(v int64) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.GpuSum = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetDownCount(v int32) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.DownCount = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetExpiredCount(v int32) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.ExpiredCount = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetExpiringCount(v int32) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.ExpiringCount = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetRunningCount(v int32) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.RunningCount = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetDiskCount(v int32) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.DiskCount = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetStorageSum(v int64) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.StorageSum = &v
	return s
}

func (s *DescribeEnsResourceUsageResponseBodyEnsResourceUsage) SetComputeResourceCount(v int32) *DescribeEnsResourceUsageResponseBodyEnsResourceUsage {
	s.ComputeResourceCount = &v
	return s
}

type DescribeEnsResourceUsageResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEnsResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEnsResourceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeEnsResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsResourceUsageResponse) SetBody(v *DescribeEnsResourceUsageResponseBody) *DescribeEnsResourceUsageResponse {
	s.Body = v
	return s
}

type DescribeExportImageStatusRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s DescribeExportImageStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportImageStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeExportImageStatusRequest) SetVersion(v string) *DescribeExportImageStatusRequest {
	s.Version = &v
	return s
}

func (s *DescribeExportImageStatusRequest) SetImageId(v string) *DescribeExportImageStatusRequest {
	s.ImageId = &v
	return s
}

type DescribeExportImageStatusResponseBody struct {
	ImageExportStatus *string `json:"ImageExportStatus,omitempty" xml:"ImageExportStatus,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExportImageStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportImageStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExportImageStatusResponseBody) SetImageExportStatus(v string) *DescribeExportImageStatusResponseBody {
	s.ImageExportStatus = &v
	return s
}

func (s *DescribeExportImageStatusResponseBody) SetRequestId(v string) *DescribeExportImageStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeExportImageStatusResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExportImageStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExportImageStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportImageStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeExportImageStatusResponse) SetHeaders(v map[string]*string) *DescribeExportImageStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeExportImageStatusResponse) SetBody(v *DescribeExportImageStatusResponseBody) *DescribeExportImageStatusResponse {
	s.Body = v
	return s
}

type DescribeExportImageInfoRequest struct {
	ImageId    *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName  *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeExportImageInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportImageInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeExportImageInfoRequest) SetImageId(v string) *DescribeExportImageInfoRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeExportImageInfoRequest) SetImageName(v string) *DescribeExportImageInfoRequest {
	s.ImageName = &v
	return s
}

func (s *DescribeExportImageInfoRequest) SetPageNumber(v int32) *DescribeExportImageInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeExportImageInfoRequest) SetPageSize(v int32) *DescribeExportImageInfoRequest {
	s.PageSize = &v
	return s
}

type DescribeExportImageInfoResponseBody struct {
	Images     *DescribeExportImageInfoResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExportImageInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportImageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExportImageInfoResponseBody) SetImages(v *DescribeExportImageInfoResponseBodyImages) *DescribeExportImageInfoResponseBody {
	s.Images = v
	return s
}

func (s *DescribeExportImageInfoResponseBody) SetPageNumber(v int32) *DescribeExportImageInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeExportImageInfoResponseBody) SetPageSize(v int32) *DescribeExportImageInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeExportImageInfoResponseBody) SetRequestId(v string) *DescribeExportImageInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExportImageInfoResponseBody) SetTotalCount(v int32) *DescribeExportImageInfoResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeExportImageInfoResponseBodyImages struct {
	Image []*DescribeExportImageInfoResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeExportImageInfoResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportImageInfoResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeExportImageInfoResponseBodyImages) SetImage(v []*DescribeExportImageInfoResponseBodyImagesImage) *DescribeExportImageInfoResponseBodyImages {
	s.Image = v
	return s
}

type DescribeExportImageInfoResponseBodyImagesImage struct {
	Architecture      *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	CreationTime      *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ExportedImageURL  *string `json:"ExportedImageURL,omitempty" xml:"ExportedImageURL,omitempty"`
	ImageExportStatus *string `json:"ImageExportStatus,omitempty" xml:"ImageExportStatus,omitempty"`
	ImageId           *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName         *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageOwnerAlias   *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	Platform          *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s DescribeExportImageInfoResponseBodyImagesImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportImageInfoResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetArchitecture(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.Architecture = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetCreationTime(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.CreationTime = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetExportedImageURL(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.ExportedImageURL = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetImageExportStatus(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.ImageExportStatus = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetImageId(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetImageName(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.ImageName = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetImageOwnerAlias(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetPlatform(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.Platform = &v
	return s
}

type DescribeExportImageInfoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExportImageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExportImageInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportImageInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeExportImageInfoResponse) SetHeaders(v map[string]*string) *DescribeExportImageInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeExportImageInfoResponse) SetBody(v *DescribeExportImageInfoResponseBody) *DescribeExportImageInfoResponse {
	s.Body = v
	return s
}

type DistApplicationDataRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DistStrategy *string `json:"DistStrategy,omitempty" xml:"DistStrategy,omitempty"`
}

func (s DistApplicationDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DistApplicationDataRequest) GoString() string {
	return s.String()
}

func (s *DistApplicationDataRequest) SetAppId(v string) *DistApplicationDataRequest {
	s.AppId = &v
	return s
}

func (s *DistApplicationDataRequest) SetData(v string) *DistApplicationDataRequest {
	s.Data = &v
	return s
}

func (s *DistApplicationDataRequest) SetDistStrategy(v string) *DistApplicationDataRequest {
	s.DistStrategy = &v
	return s
}

type DistApplicationDataResponseBody struct {
	RequestId              *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DistInstanceTotalCount *int32                                          `json:"DistInstanceTotalCount,omitempty" xml:"DistInstanceTotalCount,omitempty"`
	DistInstanceIds        *DistApplicationDataResponseBodyDistInstanceIds `json:"DistInstanceIds,omitempty" xml:"DistInstanceIds,omitempty" type:"Struct"`
	DistResults            *DistApplicationDataResponseBodyDistResults     `json:"DistResults,omitempty" xml:"DistResults,omitempty" type:"Struct"`
}

func (s DistApplicationDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DistApplicationDataResponseBody) GoString() string {
	return s.String()
}

func (s *DistApplicationDataResponseBody) SetRequestId(v string) *DistApplicationDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DistApplicationDataResponseBody) SetDistInstanceTotalCount(v int32) *DistApplicationDataResponseBody {
	s.DistInstanceTotalCount = &v
	return s
}

func (s *DistApplicationDataResponseBody) SetDistInstanceIds(v *DistApplicationDataResponseBodyDistInstanceIds) *DistApplicationDataResponseBody {
	s.DistInstanceIds = v
	return s
}

func (s *DistApplicationDataResponseBody) SetDistResults(v *DistApplicationDataResponseBodyDistResults) *DistApplicationDataResponseBody {
	s.DistResults = v
	return s
}

type DistApplicationDataResponseBodyDistInstanceIds struct {
	DistInstanceId []*string `json:"DistInstanceId,omitempty" xml:"DistInstanceId,omitempty" type:"Repeated"`
}

func (s DistApplicationDataResponseBodyDistInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s DistApplicationDataResponseBodyDistInstanceIds) GoString() string {
	return s.String()
}

func (s *DistApplicationDataResponseBodyDistInstanceIds) SetDistInstanceId(v []*string) *DistApplicationDataResponseBodyDistInstanceIds {
	s.DistInstanceId = v
	return s
}

type DistApplicationDataResponseBodyDistResults struct {
	DistResult []*DistApplicationDataResponseBodyDistResultsDistResult `json:"DistResult,omitempty" xml:"DistResult,omitempty" type:"Repeated"`
}

func (s DistApplicationDataResponseBodyDistResults) String() string {
	return tea.Prettify(s)
}

func (s DistApplicationDataResponseBodyDistResults) GoString() string {
	return s.String()
}

func (s *DistApplicationDataResponseBodyDistResults) SetDistResult(v []*DistApplicationDataResponseBodyDistResultsDistResult) *DistApplicationDataResponseBodyDistResults {
	s.DistResult = v
	return s
}

type DistApplicationDataResponseBodyDistResultsDistResult struct {
	ResultDescrip *string `json:"ResultDescrip,omitempty" xml:"ResultDescrip,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	ResultCode    *int32  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DistApplicationDataResponseBodyDistResultsDistResult) String() string {
	return tea.Prettify(s)
}

func (s DistApplicationDataResponseBodyDistResultsDistResult) GoString() string {
	return s.String()
}

func (s *DistApplicationDataResponseBodyDistResultsDistResult) SetResultDescrip(v string) *DistApplicationDataResponseBodyDistResultsDistResult {
	s.ResultDescrip = &v
	return s
}

func (s *DistApplicationDataResponseBodyDistResultsDistResult) SetVersion(v string) *DistApplicationDataResponseBodyDistResultsDistResult {
	s.Version = &v
	return s
}

func (s *DistApplicationDataResponseBodyDistResultsDistResult) SetResultCode(v int32) *DistApplicationDataResponseBodyDistResultsDistResult {
	s.ResultCode = &v
	return s
}

func (s *DistApplicationDataResponseBodyDistResultsDistResult) SetName(v string) *DistApplicationDataResponseBodyDistResultsDistResult {
	s.Name = &v
	return s
}

type DistApplicationDataResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DistApplicationDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DistApplicationDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DistApplicationDataResponse) GoString() string {
	return s.String()
}

func (s *DistApplicationDataResponse) SetHeaders(v map[string]*string) *DistApplicationDataResponse {
	s.Headers = v
	return s
}

func (s *DistApplicationDataResponse) SetBody(v *DistApplicationDataResponseBody) *DistApplicationDataResponse {
	s.Body = v
	return s
}

type ModifyImageSharePermissionRequest struct {
	ImageId        *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	AddAccounts    *string `json:"AddAccounts,omitempty" xml:"AddAccounts,omitempty"`
	RemoveAccounts *string `json:"RemoveAccounts,omitempty" xml:"RemoveAccounts,omitempty"`
}

func (s ModifyImageSharePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageSharePermissionRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageSharePermissionRequest) SetImageId(v string) *ModifyImageSharePermissionRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetAddAccounts(v string) *ModifyImageSharePermissionRequest {
	s.AddAccounts = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetRemoveAccounts(v string) *ModifyImageSharePermissionRequest {
	s.RemoveAccounts = &v
	return s
}

type ModifyImageSharePermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageSharePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageSharePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageSharePermissionResponseBody) SetRequestId(v string) *ModifyImageSharePermissionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyImageSharePermissionResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyImageSharePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyImageSharePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageSharePermissionResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageSharePermissionResponse) SetHeaders(v map[string]*string) *ModifyImageSharePermissionResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageSharePermissionResponse) SetBody(v *ModifyImageSharePermissionResponseBody) *ModifyImageSharePermissionResponse {
	s.Body = v
	return s
}

type DescribeSecurityGroupAttributeRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DescribeSecurityGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeRequest) SetVersion(v string) *DescribeSecurityGroupAttributeRequest {
	s.Version = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) SetSecurityGroupId(v string) *DescribeSecurityGroupAttributeRequest {
	s.SecurityGroupId = &v
	return s
}

type DescribeSecurityGroupAttributeResponseBody struct {
	RequestId         *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Description       *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityGroupId   *string                                                `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupName *string                                                `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	Permissions       *DescribeSecurityGroupAttributeResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Struct"`
}

func (s DescribeSecurityGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetRequestId(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetDescription(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetSecurityGroupId(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetSecurityGroupName(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetPermissions(v *DescribeSecurityGroupAttributeResponseBodyPermissions) *DescribeSecurityGroupAttributeResponseBody {
	s.Permissions = v
	return s
}

type DescribeSecurityGroupAttributeResponseBodyPermissions struct {
	Permission []*DescribeSecurityGroupAttributeResponseBodyPermissionsPermission `json:"Permission,omitempty" xml:"Permission,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupAttributeResponseBodyPermissions) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissions) SetPermission(v []*DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) *DescribeSecurityGroupAttributeResponseBodyPermissions {
	s.Permission = v
	return s
}

type DescribeSecurityGroupAttributeResponseBodyPermissionsPermission struct {
	CreationTime    *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Direction       *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	SourceCidrIp    *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	DestCidrIp      *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	Priority        *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetCreationTime(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.CreationTime = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetDirection(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Direction = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetPolicy(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Policy = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetPortRange(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.PortRange = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetSourceCidrIp(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetIpProtocol(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.IpProtocol = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetDestCidrIp(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.DestCidrIp = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetPriority(v int32) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Priority = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetSourcePortRange(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.SourcePortRange = &v
	return s
}

type DescribeSecurityGroupAttributeResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponse) SetHeaders(v map[string]*string) *DescribeSecurityGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityGroupAttributeResponse) SetBody(v *DescribeSecurityGroupAttributeResponseBody) *DescribeSecurityGroupAttributeResponse {
	s.Body = v
	return s
}

type DescribeDeviceServiceRequest struct {
	// APP ID
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDeviceServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceRequest) SetAppId(v string) *DescribeDeviceServiceRequest {
	s.AppId = &v
	return s
}

func (s *DescribeDeviceServiceRequest) SetInstanceId(v string) *DescribeDeviceServiceRequest {
	s.InstanceId = &v
	return s
}

type DescribeDeviceServiceResponseBody struct {
	// Id of the request
	RequestId           *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceDetailInfos []*DescribeDeviceServiceResponseBodyResourceDetailInfos `json:"ResourceDetailInfos,omitempty" xml:"ResourceDetailInfos,omitempty" type:"Repeated"`
}

func (s DescribeDeviceServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceResponseBody) SetRequestId(v string) *DescribeDeviceServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceServiceResponseBody) SetResourceDetailInfos(v []*DescribeDeviceServiceResponseBodyResourceDetailInfos) *DescribeDeviceServiceResponseBody {
	s.ResourceDetailInfos = v
	return s
}

type DescribeDeviceServiceResponseBodyResourceDetailInfos struct {
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	ID       *string `json:"ID,omitempty" xml:"ID,omitempty"`
	IP       *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDeviceServiceResponseBodyResourceDetailInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceServiceResponseBodyResourceDetailInfos) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetRegionID(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.RegionID = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetID(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.ID = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetIP(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.IP = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetServer(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.Server = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetStatus(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.Status = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetType(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.Type = &v
	return s
}

type DescribeDeviceServiceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeviceServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceResponse) SetHeaders(v map[string]*string) *DescribeDeviceServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceServiceResponse) SetBody(v *DescribeDeviceServiceResponseBody) *DescribeDeviceServiceResponse {
	s.Body = v
	return s
}

type DescribeEnsNetSaleDistrictRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	NetLevelCode    *string `json:"NetLevelCode,omitempty" xml:"NetLevelCode,omitempty"`
	NetDistrictCode *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty"`
}

func (s DescribeEnsNetSaleDistrictRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetSaleDistrictRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetSaleDistrictRequest) SetVersion(v string) *DescribeEnsNetSaleDistrictRequest {
	s.Version = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictRequest) SetNetLevelCode(v string) *DescribeEnsNetSaleDistrictRequest {
	s.NetLevelCode = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictRequest) SetNetDistrictCode(v string) *DescribeEnsNetSaleDistrictRequest {
	s.NetDistrictCode = &v
	return s
}

type DescribeEnsNetSaleDistrictResponseBody struct {
	Code            *int32                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	EnsNetDistricts *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts `json:"EnsNetDistricts,omitempty" xml:"EnsNetDistricts,omitempty" type:"Struct"`
	RequestId       *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnsNetSaleDistrictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetSaleDistrictResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetSaleDistrictResponseBody) SetCode(v int32) *DescribeEnsNetSaleDistrictResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBody) SetEnsNetDistricts(v *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts) *DescribeEnsNetSaleDistrictResponseBody {
	s.EnsNetDistricts = v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBody) SetRequestId(v string) *DescribeEnsNetSaleDistrictResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts struct {
	EnsNetDistrict []*DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict `json:"EnsNetDistrict,omitempty" xml:"EnsNetDistrict,omitempty" type:"Repeated"`
}

func (s DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts) SetEnsNetDistrict(v []*DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts {
	s.EnsNetDistrict = v
	return s
}

type DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict struct {
	EnsRegionIdCount      *string `json:"EnsRegionIdCount,omitempty" xml:"EnsRegionIdCount,omitempty"`
	InstanceCount         *string `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	NetDistrictCode       *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty"`
	NetDistrictEnName     *string `json:"NetDistrictEnName,omitempty" xml:"NetDistrictEnName,omitempty"`
	NetDistrictFatherCode *string `json:"NetDistrictFatherCode,omitempty" xml:"NetDistrictFatherCode,omitempty"`
	NetDistrictLevel      *string `json:"NetDistrictLevel,omitempty" xml:"NetDistrictLevel,omitempty"`
	NetDistrictName       *string `json:"NetDistrictName,omitempty" xml:"NetDistrictName,omitempty"`
}

func (s DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetEnsRegionIdCount(v string) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.EnsRegionIdCount = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetInstanceCount(v string) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.InstanceCount = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictCode(v string) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictCode = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictEnName(v string) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictEnName = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictFatherCode(v string) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictFatherCode = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictLevel(v string) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictLevel = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictName(v string) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictName = &v
	return s
}

type DescribeEnsNetSaleDistrictResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEnsNetSaleDistrictResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEnsNetSaleDistrictResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetSaleDistrictResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetSaleDistrictResponse) SetHeaders(v map[string]*string) *DescribeEnsNetSaleDistrictResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponse) SetBody(v *DescribeEnsNetSaleDistrictResponseBody) *DescribeEnsNetSaleDistrictResponse {
	s.Body = v
	return s
}

type DescribeImageSharePermissionRequest struct {
	ImageId    *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	AliyunId   *int64  `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
}

func (s DescribeImageSharePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageSharePermissionRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionRequest) SetImageId(v string) *DescribeImageSharePermissionRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetPageNumber(v string) *DescribeImageSharePermissionRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetPageSize(v string) *DescribeImageSharePermissionRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetAliyunId(v int64) *DescribeImageSharePermissionRequest {
	s.AliyunId = &v
	return s
}

type DescribeImageSharePermissionResponseBody struct {
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ImageId    *string                                           `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Accounts   *DescribeImageSharePermissionResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
}

func (s DescribeImageSharePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageSharePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionResponseBody) SetTotalCount(v int32) *DescribeImageSharePermissionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetPageSize(v int32) *DescribeImageSharePermissionResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetRequestId(v string) *DescribeImageSharePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetPageNumber(v int32) *DescribeImageSharePermissionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetImageId(v string) *DescribeImageSharePermissionResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetAccounts(v *DescribeImageSharePermissionResponseBodyAccounts) *DescribeImageSharePermissionResponseBody {
	s.Accounts = v
	return s
}

type DescribeImageSharePermissionResponseBodyAccounts struct {
	Account []*string `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s DescribeImageSharePermissionResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageSharePermissionResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionResponseBodyAccounts) SetAccount(v []*string) *DescribeImageSharePermissionResponseBodyAccounts {
	s.Account = v
	return s
}

type DescribeImageSharePermissionResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImageSharePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImageSharePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageSharePermissionResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionResponse) SetHeaders(v map[string]*string) *DescribeImageSharePermissionResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageSharePermissionResponse) SetBody(v *DescribeImageSharePermissionResponseBody) *DescribeImageSharePermissionResponse {
	s.Body = v
	return s
}

type ModifyEpnInstanceRequest struct {
	EPNInstanceId           *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	EPNInstanceName         *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	NetworkingModel         *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
	InternetMaxBandwidthOut *int32  `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
}

func (s ModifyEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyEpnInstanceRequest) SetEPNInstanceId(v string) *ModifyEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *ModifyEpnInstanceRequest) SetEPNInstanceName(v string) *ModifyEpnInstanceRequest {
	s.EPNInstanceName = &v
	return s
}

func (s *ModifyEpnInstanceRequest) SetNetworkingModel(v string) *ModifyEpnInstanceRequest {
	s.NetworkingModel = &v
	return s
}

func (s *ModifyEpnInstanceRequest) SetInternetMaxBandwidthOut(v int32) *ModifyEpnInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

type ModifyEpnInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEpnInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEpnInstanceResponseBody) SetRequestId(v string) *ModifyEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ModifyEpnInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyEpnInstanceResponse) SetHeaders(v map[string]*string) *ModifyEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyEpnInstanceResponse) SetBody(v *ModifyEpnInstanceResponseBody) *ModifyEpnInstanceResponse {
	s.Body = v
	return s
}

type DescribeEnsNetDistrictRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	NetLevelCode    *string `json:"NetLevelCode,omitempty" xml:"NetLevelCode,omitempty"`
	NetDistrictCode *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty"`
}

func (s DescribeEnsNetDistrictRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetDistrictRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetDistrictRequest) SetVersion(v string) *DescribeEnsNetDistrictRequest {
	s.Version = &v
	return s
}

func (s *DescribeEnsNetDistrictRequest) SetNetLevelCode(v string) *DescribeEnsNetDistrictRequest {
	s.NetLevelCode = &v
	return s
}

func (s *DescribeEnsNetDistrictRequest) SetNetDistrictCode(v string) *DescribeEnsNetDistrictRequest {
	s.NetDistrictCode = &v
	return s
}

type DescribeEnsNetDistrictResponseBody struct {
	Code            *int32                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	EnsNetDistricts *DescribeEnsNetDistrictResponseBodyEnsNetDistricts `json:"EnsNetDistricts,omitempty" xml:"EnsNetDistricts,omitempty" type:"Struct"`
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnsNetDistrictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetDistrictResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetDistrictResponseBody) SetCode(v int32) *DescribeEnsNetDistrictResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseBody) SetEnsNetDistricts(v *DescribeEnsNetDistrictResponseBodyEnsNetDistricts) *DescribeEnsNetDistrictResponseBody {
	s.EnsNetDistricts = v
	return s
}

func (s *DescribeEnsNetDistrictResponseBody) SetRequestId(v string) *DescribeEnsNetDistrictResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEnsNetDistrictResponseBodyEnsNetDistricts struct {
	EnsNetDistrict []*DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict `json:"EnsNetDistrict,omitempty" xml:"EnsNetDistrict,omitempty" type:"Repeated"`
}

func (s DescribeEnsNetDistrictResponseBodyEnsNetDistricts) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetDistrictResponseBodyEnsNetDistricts) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistricts) SetEnsNetDistrict(v []*DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) *DescribeEnsNetDistrictResponseBodyEnsNetDistricts {
	s.EnsNetDistrict = v
	return s
}

type DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict struct {
	EnsRegionIdCount      *string `json:"EnsRegionIdCount,omitempty" xml:"EnsRegionIdCount,omitempty"`
	NetDistrictCode       *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty"`
	NetDistrictEnName     *string `json:"NetDistrictEnName,omitempty" xml:"NetDistrictEnName,omitempty"`
	NetDistrictFatherCode *string `json:"NetDistrictFatherCode,omitempty" xml:"NetDistrictFatherCode,omitempty"`
	NetDistrictLevel      *string `json:"NetDistrictLevel,omitempty" xml:"NetDistrictLevel,omitempty"`
	NetDistrictName       *string `json:"NetDistrictName,omitempty" xml:"NetDistrictName,omitempty"`
}

func (s DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetEnsRegionIdCount(v string) *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.EnsRegionIdCount = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictCode(v string) *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictCode = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictEnName(v string) *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictEnName = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictFatherCode(v string) *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictFatherCode = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictLevel(v string) *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictLevel = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictName(v string) *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictName = &v
	return s
}

type DescribeEnsNetDistrictResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEnsNetDistrictResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEnsNetDistrictResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetDistrictResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetDistrictResponse) SetHeaders(v map[string]*string) *DescribeEnsNetDistrictResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsNetDistrictResponse) SetBody(v *DescribeEnsNetDistrictResponseBody) *DescribeEnsNetDistrictResponse {
	s.Body = v
	return s
}

type DescribeNetworksRequest struct {
	NetworkId   *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	NetworkName *string `json:"NetworkName,omitempty" xml:"NetworkName,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeNetworksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworksRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworksRequest) SetNetworkId(v string) *DescribeNetworksRequest {
	s.NetworkId = &v
	return s
}

func (s *DescribeNetworksRequest) SetEnsRegionId(v string) *DescribeNetworksRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworksRequest) SetNetworkName(v string) *DescribeNetworksRequest {
	s.NetworkName = &v
	return s
}

func (s *DescribeNetworksRequest) SetPageNumber(v int32) *DescribeNetworksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworksRequest) SetPageSize(v int32) *DescribeNetworksRequest {
	s.PageSize = &v
	return s
}

type DescribeNetworksResponseBody struct {
	// Id of the request
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Networks   *DescribeNetworksResponseBodyNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Struct"`
}

func (s DescribeNetworksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworksResponseBody) SetRequestId(v string) *DescribeNetworksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworksResponseBody) SetTotalCount(v int32) *DescribeNetworksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNetworksResponseBody) SetPageSize(v int32) *DescribeNetworksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworksResponseBody) SetPageNumber(v int32) *DescribeNetworksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworksResponseBody) SetNetworks(v *DescribeNetworksResponseBodyNetworks) *DescribeNetworksResponseBody {
	s.Networks = v
	return s
}

type DescribeNetworksResponseBodyNetworks struct {
	Network []*DescribeNetworksResponseBodyNetworksNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Repeated"`
}

func (s DescribeNetworksResponseBodyNetworks) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworksResponseBodyNetworks) GoString() string {
	return s.String()
}

func (s *DescribeNetworksResponseBodyNetworks) SetNetwork(v []*DescribeNetworksResponseBodyNetworksNetwork) *DescribeNetworksResponseBodyNetworks {
	s.Network = v
	return s
}

type DescribeNetworksResponseBodyNetworksNetwork struct {
	EnsRegionId *string                                                `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	NetworkId   *string                                                `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	NetworkName *string                                                `json:"NetworkName,omitempty" xml:"NetworkName,omitempty"`
	CidrBlock   *string                                                `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Status      *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Description *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedTime *string                                                `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	VSwitchIds  *DescribeNetworksResponseBodyNetworksNetworkVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
}

func (s DescribeNetworksResponseBodyNetworksNetwork) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworksResponseBodyNetworksNetwork) GoString() string {
	return s.String()
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetEnsRegionId(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetNetworkId(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.NetworkId = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetNetworkName(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.NetworkName = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetCidrBlock(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.CidrBlock = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetStatus(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.Status = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetDescription(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.Description = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetCreatedTime(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.CreatedTime = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetVSwitchIds(v *DescribeNetworksResponseBodyNetworksNetworkVSwitchIds) *DescribeNetworksResponseBodyNetworksNetwork {
	s.VSwitchIds = v
	return s
}

type DescribeNetworksResponseBodyNetworksNetworkVSwitchIds struct {
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s DescribeNetworksResponseBodyNetworksNetworkVSwitchIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworksResponseBodyNetworksNetworkVSwitchIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworksResponseBodyNetworksNetworkVSwitchIds) SetVSwitchId(v []*string) *DescribeNetworksResponseBodyNetworksNetworkVSwitchIds {
	s.VSwitchId = v
	return s
}

type DescribeNetworksResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNetworksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNetworksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworksResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworksResponse) SetHeaders(v map[string]*string) *DescribeNetworksResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworksResponse) SetBody(v *DescribeNetworksResponseBody) *DescribeNetworksResponse {
	s.Body = v
	return s
}

type RunServiceScheduleRequest struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Uuid             *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	ClientIp         *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ServiceAction    *string `json:"ServiceAction,omitempty" xml:"ServiceAction,omitempty"`
	PodConfigName    *string `json:"PodConfigName,omitempty" xml:"PodConfigName,omitempty"`
	PreLockedTimeout *int32  `json:"PreLockedTimeout,omitempty" xml:"PreLockedTimeout,omitempty"`
	Directorys       *string `json:"Directorys,omitempty" xml:"Directorys,omitempty"`
	ServiceCommands  *string `json:"ServiceCommands,omitempty" xml:"ServiceCommands,omitempty"`
	ScheduleStrategy *string `json:"ScheduleStrategy,omitempty" xml:"ScheduleStrategy,omitempty"`
}

func (s RunServiceScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s RunServiceScheduleRequest) GoString() string {
	return s.String()
}

func (s *RunServiceScheduleRequest) SetAppId(v string) *RunServiceScheduleRequest {
	s.AppId = &v
	return s
}

func (s *RunServiceScheduleRequest) SetUuid(v string) *RunServiceScheduleRequest {
	s.Uuid = &v
	return s
}

func (s *RunServiceScheduleRequest) SetClientIp(v string) *RunServiceScheduleRequest {
	s.ClientIp = &v
	return s
}

func (s *RunServiceScheduleRequest) SetServiceAction(v string) *RunServiceScheduleRequest {
	s.ServiceAction = &v
	return s
}

func (s *RunServiceScheduleRequest) SetPodConfigName(v string) *RunServiceScheduleRequest {
	s.PodConfigName = &v
	return s
}

func (s *RunServiceScheduleRequest) SetPreLockedTimeout(v int32) *RunServiceScheduleRequest {
	s.PreLockedTimeout = &v
	return s
}

func (s *RunServiceScheduleRequest) SetDirectorys(v string) *RunServiceScheduleRequest {
	s.Directorys = &v
	return s
}

func (s *RunServiceScheduleRequest) SetServiceCommands(v string) *RunServiceScheduleRequest {
	s.ServiceCommands = &v
	return s
}

func (s *RunServiceScheduleRequest) SetScheduleStrategy(v string) *RunServiceScheduleRequest {
	s.ScheduleStrategy = &v
	return s
}

type RunServiceScheduleResponseBody struct {
	CommandResults  *RunServiceScheduleResponseBodyCommandResults `json:"CommandResults,omitempty" xml:"CommandResults,omitempty" type:"Struct"`
	Index           *int32                                        `json:"Index,omitempty" xml:"Index,omitempty"`
	InstanceId      *string                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIp      *string                                       `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	InstancePort    *int32                                        `json:"InstancePort,omitempty" xml:"InstancePort,omitempty"`
	RequestId       *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestRepeated *string                                       `json:"RequestRepeated,omitempty" xml:"RequestRepeated,omitempty"`
	TcpPorts        *bool                                         `json:"TcpPorts,omitempty" xml:"TcpPorts,omitempty"`
}

func (s RunServiceScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunServiceScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *RunServiceScheduleResponseBody) SetCommandResults(v *RunServiceScheduleResponseBodyCommandResults) *RunServiceScheduleResponseBody {
	s.CommandResults = v
	return s
}

func (s *RunServiceScheduleResponseBody) SetIndex(v int32) *RunServiceScheduleResponseBody {
	s.Index = &v
	return s
}

func (s *RunServiceScheduleResponseBody) SetInstanceId(v string) *RunServiceScheduleResponseBody {
	s.InstanceId = &v
	return s
}

func (s *RunServiceScheduleResponseBody) SetInstanceIp(v string) *RunServiceScheduleResponseBody {
	s.InstanceIp = &v
	return s
}

func (s *RunServiceScheduleResponseBody) SetInstancePort(v int32) *RunServiceScheduleResponseBody {
	s.InstancePort = &v
	return s
}

func (s *RunServiceScheduleResponseBody) SetRequestId(v string) *RunServiceScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunServiceScheduleResponseBody) SetRequestRepeated(v string) *RunServiceScheduleResponseBody {
	s.RequestRepeated = &v
	return s
}

func (s *RunServiceScheduleResponseBody) SetTcpPorts(v bool) *RunServiceScheduleResponseBody {
	s.TcpPorts = &v
	return s
}

type RunServiceScheduleResponseBodyCommandResults struct {
	CommandResult []*RunServiceScheduleResponseBodyCommandResultsCommandResult `json:"CommandResult,omitempty" xml:"CommandResult,omitempty" type:"Repeated"`
}

func (s RunServiceScheduleResponseBodyCommandResults) String() string {
	return tea.Prettify(s)
}

func (s RunServiceScheduleResponseBodyCommandResults) GoString() string {
	return s.String()
}

func (s *RunServiceScheduleResponseBodyCommandResults) SetCommandResult(v []*RunServiceScheduleResponseBodyCommandResultsCommandResult) *RunServiceScheduleResponseBodyCommandResults {
	s.CommandResult = v
	return s
}

type RunServiceScheduleResponseBodyCommandResultsCommandResult struct {
	Command       *string `json:"Command,omitempty" xml:"Command,omitempty"`
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	ResultMsg     *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
}

func (s RunServiceScheduleResponseBodyCommandResultsCommandResult) String() string {
	return tea.Prettify(s)
}

func (s RunServiceScheduleResponseBodyCommandResultsCommandResult) GoString() string {
	return s.String()
}

func (s *RunServiceScheduleResponseBodyCommandResultsCommandResult) SetCommand(v string) *RunServiceScheduleResponseBodyCommandResultsCommandResult {
	s.Command = &v
	return s
}

func (s *RunServiceScheduleResponseBodyCommandResultsCommandResult) SetContainerName(v string) *RunServiceScheduleResponseBodyCommandResultsCommandResult {
	s.ContainerName = &v
	return s
}

func (s *RunServiceScheduleResponseBodyCommandResultsCommandResult) SetResultMsg(v string) *RunServiceScheduleResponseBodyCommandResultsCommandResult {
	s.ResultMsg = &v
	return s
}

type RunServiceScheduleResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunServiceScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunServiceScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s RunServiceScheduleResponse) GoString() string {
	return s.String()
}

func (s *RunServiceScheduleResponse) SetHeaders(v map[string]*string) *RunServiceScheduleResponse {
	s.Headers = v
	return s
}

func (s *RunServiceScheduleResponse) SetBody(v *RunServiceScheduleResponseBody) *RunServiceScheduleResponse {
	s.Body = v
	return s
}

type RemoveVSwitchesFromEpnInstanceRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	VSwitchesInfo *string `json:"VSwitchesInfo,omitempty" xml:"VSwitchesInfo,omitempty"`
}

func (s RemoveVSwitchesFromEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveVSwitchesFromEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *RemoveVSwitchesFromEpnInstanceRequest) SetEPNInstanceId(v string) *RemoveVSwitchesFromEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *RemoveVSwitchesFromEpnInstanceRequest) SetVSwitchesInfo(v string) *RemoveVSwitchesFromEpnInstanceRequest {
	s.VSwitchesInfo = &v
	return s
}

type RemoveVSwitchesFromEpnInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveVSwitchesFromEpnInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveVSwitchesFromEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveVSwitchesFromEpnInstanceResponseBody) SetRequestId(v string) *RemoveVSwitchesFromEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RemoveVSwitchesFromEpnInstanceResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveVSwitchesFromEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveVSwitchesFromEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveVSwitchesFromEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *RemoveVSwitchesFromEpnInstanceResponse) SetHeaders(v map[string]*string) *RemoveVSwitchesFromEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *RemoveVSwitchesFromEpnInstanceResponse) SetBody(v *RemoveVSwitchesFromEpnInstanceResponseBody) *RemoveVSwitchesFromEpnInstanceResponse {
	s.Body = v
	return s
}

type SchedulePodRequest struct {
	AliUid            *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	GroupUuid         *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	WorkloadUuid      *string `json:"WorkloadUuid,omitempty" xml:"WorkloadUuid,omitempty"`
	Tenant            *string `json:"Tenant,omitempty" xml:"Tenant,omitempty"`
	Regions           *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	AreaCodes         *string `json:"AreaCodes,omitempty" xml:"AreaCodes,omitempty"`
	Isps              *string `json:"Isps,omitempty" xml:"Isps,omitempty"`
	Requirements      *string `json:"Requirements,omitempty" xml:"Requirements,omitempty"`
	Labels            *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	ResourceAttribute *string `json:"ResourceAttribute,omitempty" xml:"ResourceAttribute,omitempty"`
}

func (s SchedulePodRequest) String() string {
	return tea.Prettify(s)
}

func (s SchedulePodRequest) GoString() string {
	return s.String()
}

func (s *SchedulePodRequest) SetAliUid(v int64) *SchedulePodRequest {
	s.AliUid = &v
	return s
}

func (s *SchedulePodRequest) SetGroupUuid(v string) *SchedulePodRequest {
	s.GroupUuid = &v
	return s
}

func (s *SchedulePodRequest) SetWorkloadUuid(v string) *SchedulePodRequest {
	s.WorkloadUuid = &v
	return s
}

func (s *SchedulePodRequest) SetTenant(v string) *SchedulePodRequest {
	s.Tenant = &v
	return s
}

func (s *SchedulePodRequest) SetRegions(v string) *SchedulePodRequest {
	s.Regions = &v
	return s
}

func (s *SchedulePodRequest) SetAreaCodes(v string) *SchedulePodRequest {
	s.AreaCodes = &v
	return s
}

func (s *SchedulePodRequest) SetIsps(v string) *SchedulePodRequest {
	s.Isps = &v
	return s
}

func (s *SchedulePodRequest) SetRequirements(v string) *SchedulePodRequest {
	s.Requirements = &v
	return s
}

func (s *SchedulePodRequest) SetLabels(v string) *SchedulePodRequest {
	s.Labels = &v
	return s
}

func (s *SchedulePodRequest) SetResourceAttribute(v string) *SchedulePodRequest {
	s.ResourceAttribute = &v
	return s
}

type SchedulePodResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Desc      *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s SchedulePodResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SchedulePodResponseBody) GoString() string {
	return s.String()
}

func (s *SchedulePodResponseBody) SetRequestId(v string) *SchedulePodResponseBody {
	s.RequestId = &v
	return s
}

func (s *SchedulePodResponseBody) SetCode(v int64) *SchedulePodResponseBody {
	s.Code = &v
	return s
}

func (s *SchedulePodResponseBody) SetMsg(v string) *SchedulePodResponseBody {
	s.Msg = &v
	return s
}

func (s *SchedulePodResponseBody) SetDesc(v string) *SchedulePodResponseBody {
	s.Desc = &v
	return s
}

func (s *SchedulePodResponseBody) SetData(v string) *SchedulePodResponseBody {
	s.Data = &v
	return s
}

type SchedulePodResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SchedulePodResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SchedulePodResponse) String() string {
	return tea.Prettify(s)
}

func (s SchedulePodResponse) GoString() string {
	return s.String()
}

func (s *SchedulePodResponse) SetHeaders(v map[string]*string) *SchedulePodResponse {
	s.Headers = v
	return s
}

func (s *SchedulePodResponse) SetBody(v *SchedulePodResponseBody) *SchedulePodResponse {
	s.Body = v
	return s
}

type DescribeNetworkInterfacesRequest struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	EnsRegionId      *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	PageNumber       *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeNetworkInterfacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkInterfacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesRequest) SetInstanceId(v string) *DescribeNetworkInterfacesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetVSwitchId(v string) *DescribeNetworkInterfacesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetEnsRegionId(v string) *DescribeNetworkInterfacesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetPrimaryIpAddress(v string) *DescribeNetworkInterfacesRequest {
	s.PrimaryIpAddress = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetPageNumber(v string) *DescribeNetworkInterfacesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetPageSize(v string) *DescribeNetworkInterfacesRequest {
	s.PageSize = &v
	return s
}

type DescribeNetworkInterfacesResponseBody struct {
	NetworkInterfaceSets *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets `json:"NetworkInterfaceSets,omitempty" xml:"NetworkInterfaceSets,omitempty" type:"Struct"`
	PageNumber           *int32                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount           *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNetworkInterfacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBody) SetNetworkInterfaceSets(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) *DescribeNetworkInterfacesResponseBody {
	s.NetworkInterfaceSets = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) SetPageNumber(v int32) *DescribeNetworkInterfacesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) SetPageSize(v int32) *DescribeNetworkInterfacesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) SetRequestId(v string) *DescribeNetworkInterfacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) SetTotalCount(v int32) *DescribeNetworkInterfacesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets struct {
	NetworkInterfaceSet []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet `json:"NetworkInterfaceSet,omitempty" xml:"NetworkInterfaceSet,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) SetNetworkInterfaceSet(v []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.NetworkInterfaceSet = v
	return s
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet struct {
	CreationTime       *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	EnsRegionId        *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MacAddress         *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	PrimaryIp          *string `json:"PrimaryIp,omitempty" xml:"PrimaryIp,omitempty"`
	PrimaryIpType      *string `json:"PrimaryIpType,omitempty" xml:"PrimaryIpType,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetCreationTime(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.CreationTime = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetEnsRegionId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetInstanceId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetMacAddress(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.MacAddress = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetNetworkInterfaceId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetPrimaryIp(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.PrimaryIp = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetPrimaryIpType(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.PrimaryIpType = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetStatus(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.Status = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetVSwitchId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.VSwitchId = &v
	return s
}

type DescribeNetworkInterfacesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNetworkInterfacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNetworkInterfacesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkInterfacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponse) SetHeaders(v map[string]*string) *DescribeNetworkInterfacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkInterfacesResponse) SetBody(v *DescribeNetworkInterfacesResponseBody) *DescribeNetworkInterfacesResponse {
	s.Body = v
	return s
}

type AllocateEipAddressRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	Count       *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	MinCount    *int32  `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
}

func (s AllocateEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateEipAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressRequest) SetVersion(v string) *AllocateEipAddressRequest {
	s.Version = &v
	return s
}

func (s *AllocateEipAddressRequest) SetEnsRegionId(v string) *AllocateEipAddressRequest {
	s.EnsRegionId = &v
	return s
}

func (s *AllocateEipAddressRequest) SetCount(v int32) *AllocateEipAddressRequest {
	s.Count = &v
	return s
}

func (s *AllocateEipAddressRequest) SetMinCount(v int32) *AllocateEipAddressRequest {
	s.MinCount = &v
	return s
}

type AllocateEipAddressResponseBody struct {
	BizStatusCode *string                                     `json:"BizStatusCode,omitempty" xml:"BizStatusCode,omitempty"`
	EipAddresses  *AllocateEipAddressResponseBodyEipAddresses `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty" type:"Struct"`
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateEipAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressResponseBody) SetBizStatusCode(v string) *AllocateEipAddressResponseBody {
	s.BizStatusCode = &v
	return s
}

func (s *AllocateEipAddressResponseBody) SetEipAddresses(v *AllocateEipAddressResponseBodyEipAddresses) *AllocateEipAddressResponseBody {
	s.EipAddresses = v
	return s
}

func (s *AllocateEipAddressResponseBody) SetRequestId(v string) *AllocateEipAddressResponseBody {
	s.RequestId = &v
	return s
}

type AllocateEipAddressResponseBodyEipAddresses struct {
	EipAddress []*AllocateEipAddressResponseBodyEipAddressesEipAddress `json:"EipAddress,omitempty" xml:"EipAddress,omitempty" type:"Repeated"`
}

func (s AllocateEipAddressResponseBodyEipAddresses) String() string {
	return tea.Prettify(s)
}

func (s AllocateEipAddressResponseBodyEipAddresses) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressResponseBodyEipAddresses) SetEipAddress(v []*AllocateEipAddressResponseBodyEipAddressesEipAddress) *AllocateEipAddressResponseBodyEipAddresses {
	s.EipAddress = v
	return s
}

type AllocateEipAddressResponseBodyEipAddressesEipAddress struct {
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
}

func (s AllocateEipAddressResponseBodyEipAddressesEipAddress) String() string {
	return tea.Prettify(s)
}

func (s AllocateEipAddressResponseBodyEipAddressesEipAddress) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressResponseBodyEipAddressesEipAddress) SetEip(v string) *AllocateEipAddressResponseBodyEipAddressesEipAddress {
	s.Eip = &v
	return s
}

type AllocateEipAddressResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AllocateEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocateEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateEipAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressResponse) SetHeaders(v map[string]*string) *AllocateEipAddressResponse {
	s.Headers = v
	return s
}

func (s *AllocateEipAddressResponse) SetBody(v *AllocateEipAddressResponseBody) *AllocateEipAddressResponse {
	s.Body = v
	return s
}

type CreateKeyPairRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
}

func (s CreateKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyPairRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyPairRequest) SetVersion(v string) *CreateKeyPairRequest {
	s.Version = &v
	return s
}

func (s *CreateKeyPairRequest) SetKeyPairName(v string) *CreateKeyPairRequest {
	s.KeyPairName = &v
	return s
}

type CreateKeyPairResponseBody struct {
	KeyPairFingerPrint *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty"`
	KeyPairId          *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	KeyPairName        *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	PrivateKeyBody     *string `json:"PrivateKeyBody,omitempty" xml:"PrivateKeyBody,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyPairResponseBody) SetKeyPairFingerPrint(v string) *CreateKeyPairResponseBody {
	s.KeyPairFingerPrint = &v
	return s
}

func (s *CreateKeyPairResponseBody) SetKeyPairId(v string) *CreateKeyPairResponseBody {
	s.KeyPairId = &v
	return s
}

func (s *CreateKeyPairResponseBody) SetKeyPairName(v string) *CreateKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *CreateKeyPairResponseBody) SetPrivateKeyBody(v string) *CreateKeyPairResponseBody {
	s.PrivateKeyBody = &v
	return s
}

func (s *CreateKeyPairResponseBody) SetRequestId(v string) *CreateKeyPairResponseBody {
	s.RequestId = &v
	return s
}

type CreateKeyPairResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyPairResponse) GoString() string {
	return s.String()
}

func (s *CreateKeyPairResponse) SetHeaders(v map[string]*string) *CreateKeyPairResponse {
	s.Headers = v
	return s
}

func (s *CreateKeyPairResponse) SetBody(v *CreateKeyPairResponseBody) *CreateKeyPairResponse {
	s.Body = v
	return s
}

type DescribeAvailableResourceInfoRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAvailableResourceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoRequest) SetVersion(v string) *DescribeAvailableResourceInfoRequest {
	s.Version = &v
	return s
}

type DescribeAvailableResourceInfoResponseBody struct {
	RequestId        *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Images           *DescribeAvailableResourceInfoResponseBodyImages           `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	SupportResources *DescribeAvailableResourceInfoResponseBodySupportResources `json:"SupportResources,omitempty" xml:"SupportResources,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBody) SetRequestId(v string) *DescribeAvailableResourceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBody) SetImages(v *DescribeAvailableResourceInfoResponseBodyImages) *DescribeAvailableResourceInfoResponseBody {
	s.Images = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBody) SetSupportResources(v *DescribeAvailableResourceInfoResponseBodySupportResources) *DescribeAvailableResourceInfoResponseBody {
	s.SupportResources = v
	return s
}

type DescribeAvailableResourceInfoResponseBodyImages struct {
	Image []*DescribeAvailableResourceInfoResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodyImages) SetImage(v []*DescribeAvailableResourceInfoResponseBodyImagesImage) *DescribeAvailableResourceInfoResponseBodyImages {
	s.Image = v
	return s
}

type DescribeAvailableResourceInfoResponseBodyImagesImage struct {
	ImageSize *int32  `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s DescribeAvailableResourceInfoResponseBodyImagesImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodyImagesImage) SetImageSize(v int32) *DescribeAvailableResourceInfoResponseBodyImagesImage {
	s.ImageSize = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodyImagesImage) SetImageName(v string) *DescribeAvailableResourceInfoResponseBodyImagesImage {
	s.ImageName = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodyImagesImage) SetImageId(v string) *DescribeAvailableResourceInfoResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

type DescribeAvailableResourceInfoResponseBodySupportResources struct {
	SupportResource []*DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource `json:"SupportResource,omitempty" xml:"SupportResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResources) SetSupportResource(v []*DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) *DescribeAvailableResourceInfoResponseBodySupportResources {
	s.SupportResource = v
	return s
}

type DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource struct {
	DataDiskMaxSize     *int32                                                                                       `json:"DataDiskMaxSize,omitempty" xml:"DataDiskMaxSize,omitempty"`
	BandwidthTypes      *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes      `json:"BandwidthTypes,omitempty" xml:"BandwidthTypes,omitempty" type:"Struct"`
	SystemDiskMinSize   *int32                                                                                       `json:"SystemDiskMinSize,omitempty" xml:"SystemDiskMinSize,omitempty"`
	EnsRegionIds        *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds        `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" type:"Struct"`
	EnsRegionIdsExtends *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends `json:"EnsRegionIdsExtends,omitempty" xml:"EnsRegionIdsExtends,omitempty" type:"Struct"`
	InstanceSpeces      *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces      `json:"InstanceSpeces,omitempty" xml:"InstanceSpeces,omitempty" type:"Struct"`
	SystemDiskMaxSize   *int32                                                                                       `json:"SystemDiskMaxSize,omitempty" xml:"SystemDiskMaxSize,omitempty"`
	DataDiskMinSize     *int32                                                                                       `json:"DataDiskMinSize,omitempty" xml:"DataDiskMinSize,omitempty"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetDataDiskMaxSize(v int32) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.DataDiskMaxSize = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetBandwidthTypes(v *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.BandwidthTypes = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetSystemDiskMinSize(v int32) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.SystemDiskMinSize = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetEnsRegionIds(v *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetEnsRegionIdsExtends(v *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.EnsRegionIdsExtends = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetInstanceSpeces(v *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.InstanceSpeces = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetSystemDiskMaxSize(v int32) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.SystemDiskMaxSize = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource) SetDataDiskMinSize(v int32) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResource {
	s.DataDiskMinSize = &v
	return s
}

type DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes struct {
	BandwidthType []*string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes) SetBandwidthType(v []*string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceBandwidthTypes {
	s.BandwidthType = v
	return s
}

type DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds struct {
	EnsRegionId []*string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds) SetEnsRegionId(v []*string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIds {
	s.EnsRegionId = v
	return s
}

type DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends struct {
	EnsRegionId []*DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends) SetEnsRegionId(v []*DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtends {
	s.EnsRegionId = v
	return s
}

type DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId struct {
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	EnName      *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	Area        *string `json:"Area,omitempty" xml:"Area,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetEnsRegionId(v string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetEnName(v string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.EnName = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetArea(v string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.Area = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetName(v string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.Name = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetProvince(v string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.Province = &v
	return s
}

type DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces struct {
	InstanceSpec []*string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces) SetInstanceSpec(v []*string) *DescribeAvailableResourceInfoResponseBodySupportResourcesSupportResourceInstanceSpeces {
	s.InstanceSpec = v
	return s
}

type DescribeAvailableResourceInfoResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAvailableResourceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableResourceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponse) SetHeaders(v map[string]*string) *DescribeAvailableResourceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableResourceInfoResponse) SetBody(v *DescribeAvailableResourceInfoResponseBody) *DescribeAvailableResourceInfoResponse {
	s.Body = v
	return s
}

type CreateVmAndSaveStockRequest struct {
	AliUid            *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Tenant            *string `json:"Tenant,omitempty" xml:"Tenant,omitempty"`
	WorkloadUuid      *string `json:"WorkloadUuid,omitempty" xml:"WorkloadUuid,omitempty"`
	GroupUuid         *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	ResourceAttribute *string `json:"ResourceAttribute,omitempty" xml:"ResourceAttribute,omitempty"`
}

func (s CreateVmAndSaveStockRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVmAndSaveStockRequest) GoString() string {
	return s.String()
}

func (s *CreateVmAndSaveStockRequest) SetAliUid(v int64) *CreateVmAndSaveStockRequest {
	s.AliUid = &v
	return s
}

func (s *CreateVmAndSaveStockRequest) SetTenant(v string) *CreateVmAndSaveStockRequest {
	s.Tenant = &v
	return s
}

func (s *CreateVmAndSaveStockRequest) SetWorkloadUuid(v string) *CreateVmAndSaveStockRequest {
	s.WorkloadUuid = &v
	return s
}

func (s *CreateVmAndSaveStockRequest) SetGroupUuid(v string) *CreateVmAndSaveStockRequest {
	s.GroupUuid = &v
	return s
}

func (s *CreateVmAndSaveStockRequest) SetResourceAttribute(v string) *CreateVmAndSaveStockRequest {
	s.ResourceAttribute = &v
	return s
}

type CreateVmAndSaveStockResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Desc      *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s CreateVmAndSaveStockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVmAndSaveStockResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVmAndSaveStockResponseBody) SetRequestId(v string) *CreateVmAndSaveStockResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVmAndSaveStockResponseBody) SetCode(v int32) *CreateVmAndSaveStockResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVmAndSaveStockResponseBody) SetMsg(v string) *CreateVmAndSaveStockResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateVmAndSaveStockResponseBody) SetDesc(v string) *CreateVmAndSaveStockResponseBody {
	s.Desc = &v
	return s
}

func (s *CreateVmAndSaveStockResponseBody) SetData(v string) *CreateVmAndSaveStockResponseBody {
	s.Data = &v
	return s
}

type CreateVmAndSaveStockResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVmAndSaveStockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVmAndSaveStockResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVmAndSaveStockResponse) GoString() string {
	return s.String()
}

func (s *CreateVmAndSaveStockResponse) SetHeaders(v map[string]*string) *CreateVmAndSaveStockResponse {
	s.Headers = v
	return s
}

func (s *CreateVmAndSaveStockResponse) SetBody(v *CreateVmAndSaveStockResponseBody) *CreateVmAndSaveStockResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("ens"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DescribeServcieScheduleWithOptions(request *DescribeServcieScheduleRequest, runtime *util.RuntimeOptions) (_result *DescribeServcieScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeServcieScheduleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeServcieSchedule"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServcieSchedule(request *DescribeServcieScheduleRequest) (_result *DescribeServcieScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServcieScheduleResponse{}
	_body, _err := client.DescribeServcieScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinSecurityGroupWithOptions(request *JoinSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *JoinSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &JoinSecurityGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("JoinSecurityGroup"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinSecurityGroup(request *JoinSecurityGroupRequest) (_result *JoinSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinSecurityGroupResponse{}
	_body, _err := client.JoinSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartDeviceInstanceWithOptions(request *RestartDeviceInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartDeviceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RestartDeviceInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestartDeviceInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartDeviceInstance(request *RestartDeviceInstanceRequest) (_result *RestartDeviceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartDeviceInstanceResponse{}
	_body, _err := client.RestartDeviceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleasePrePaidInstanceWithOptions(request *ReleasePrePaidInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleasePrePaidInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleasePrePaidInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleasePrePaidInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleasePrePaidInstance(request *ReleasePrePaidInstanceRequest) (_result *ReleasePrePaidInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleasePrePaidInstanceResponse{}
	_body, _err := client.ReleasePrePaidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApplicationWithOptions(request *CreateApplicationRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateApplication"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNetworkWithOptions(request *CreateNetworkRequest, runtime *util.RuntimeOptions) (_result *CreateNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateNetworkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateNetwork"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNetwork(request *CreateNetworkRequest) (_result *CreateNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNetworkResponse{}
	_body, _err := client.CreateNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReInitDiskWithOptions(request *ReInitDiskRequest, runtime *util.RuntimeOptions) (_result *ReInitDiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReInitDiskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReInitDisk"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReInitDisk(request *ReInitDiskRequest) (_result *ReInitDiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReInitDiskResponse{}
	_body, _err := client.ReInitDiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddNetworkInterfaceToInstanceWithOptions(request *AddNetworkInterfaceToInstanceRequest, runtime *util.RuntimeOptions) (_result *AddNetworkInterfaceToInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddNetworkInterfaceToInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddNetworkInterfaceToInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddNetworkInterfaceToInstance(request *AddNetworkInterfaceToInstanceRequest) (_result *AddNetworkInterfaceToInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddNetworkInterfaceToInstanceResponse{}
	_body, _err := client.AddNetworkInterfaceToInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateEipAddressWithOptions(request *AssociateEipAddressRequest, runtime *util.RuntimeOptions) (_result *AssociateEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssociateEipAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssociateEipAddress"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateEipAddress(request *AssociateEipAddressRequest) (_result *AssociateEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateEipAddressResponse{}
	_body, _err := client.AssociateEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopEpnInstanceWithOptions(request *StopEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *StopEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopEpnInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopEpnInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopEpnInstance(request *StopEpnInstanceRequest) (_result *StopEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopEpnInstanceResponse{}
	_body, _err := client.StopEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEipAddressesWithOptions(request *DescribeEipAddressesRequest, runtime *util.RuntimeOptions) (_result *DescribeEipAddressesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEipAddressesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEipAddresses"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEipAddresses(request *DescribeEipAddressesRequest) (_result *DescribeEipAddressesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEipAddressesResponse{}
	_body, _err := client.DescribeEipAddressesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeSecurityGroupEgressWithOptions(request *RevokeSecurityGroupEgressRequest, runtime *util.RuntimeOptions) (_result *RevokeSecurityGroupEgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RevokeSecurityGroupEgressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RevokeSecurityGroupEgress"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeSecurityGroupEgress(request *RevokeSecurityGroupEgressRequest) (_result *RevokeSecurityGroupEgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeSecurityGroupEgressResponse{}
	_body, _err := client.RevokeSecurityGroupEgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigureSecurityGroupPermissionsWithOptions(tmpReq *ConfigureSecurityGroupPermissionsRequest, runtime *util.RuntimeOptions) (_result *ConfigureSecurityGroupPermissionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ConfigureSecurityGroupPermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AuthorizePermissions)) {
		request.AuthorizePermissionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthorizePermissions, tea.String("AuthorizePermissions"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RevokePermissions)) {
		request.RevokePermissionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RevokePermissions, tea.String("RevokePermissions"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigureSecurityGroupPermissionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigureSecurityGroupPermissions"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigureSecurityGroupPermissions(request *ConfigureSecurityGroupPermissionsRequest) (_result *ConfigureSecurityGroupPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigureSecurityGroupPermissionsResponse{}
	_body, _err := client.ConfigureSecurityGroupPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnsNetLevelWithOptions(request *DescribeEnsNetLevelRequest, runtime *util.RuntimeOptions) (_result *DescribeEnsNetLevelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEnsNetLevelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEnsNetLevel"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnsNetLevel(request *DescribeEnsNetLevelRequest) (_result *DescribeEnsNetLevelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnsNetLevelResponse{}
	_body, _err := client.DescribeEnsNetLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNetworkAttributeWithOptions(request *DescribeNetworkAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeNetworkAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNetworkAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNetworkAttribute"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNetworkAttribute(request *DescribeNetworkAttributeRequest) (_result *DescribeNetworkAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNetworkAttributeResponse{}
	_body, _err := client.DescribeNetworkAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinPublicIpsToEpnInstanceWithOptions(request *JoinPublicIpsToEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *JoinPublicIpsToEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &JoinPublicIpsToEpnInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("JoinPublicIpsToEpnInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinPublicIpsToEpnInstance(request *JoinPublicIpsToEpnInstanceRequest) (_result *JoinPublicIpsToEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinPublicIpsToEpnInstanceResponse{}
	_body, _err := client.JoinPublicIpsToEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInstanceWithOptions(request *StopInstanceRequest, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInstance(request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachEnsInstancesWithOptions(request *AttachEnsInstancesRequest, runtime *util.RuntimeOptions) (_result *AttachEnsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachEnsInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachEnsInstances"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachEnsInstances(request *AttachEnsInstancesRequest) (_result *AttachEnsInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachEnsInstancesResponse{}
	_body, _err := client.AttachEnsInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEpnBandwitdhByInternetChargeTypeWithOptions(request *DescribeEpnBandwitdhByInternetChargeTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeEpnBandwitdhByInternetChargeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEpnBandwitdhByInternetChargeTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEpnBandwitdhByInternetChargeType"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEpnBandwitdhByInternetChargeType(request *DescribeEpnBandwitdhByInternetChargeTypeRequest) (_result *DescribeEpnBandwitdhByInternetChargeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEpnBandwitdhByInternetChargeTypeResponse{}
	_body, _err := client.DescribeEpnBandwitdhByInternetChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeSecurityGroupWithOptions(request *RevokeSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *RevokeSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RevokeSecurityGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RevokeSecurityGroup"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeSecurityGroup(request *RevokeSecurityGroupRequest) (_result *RevokeSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeSecurityGroupResponse{}
	_body, _err := client.RevokeSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemovePublicIpsFromEpnInstanceWithOptions(request *RemovePublicIpsFromEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *RemovePublicIpsFromEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemovePublicIpsFromEpnInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemovePublicIpsFromEpnInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemovePublicIpsFromEpnInstance(request *RemovePublicIpsFromEpnInstanceRequest) (_result *RemovePublicIpsFromEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemovePublicIpsFromEpnInstanceResponse{}
	_body, _err := client.RemovePublicIpsFromEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeApplicationWithOptions(request *UpgradeApplicationRequest, runtime *util.RuntimeOptions) (_result *UpgradeApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeApplicationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeApplication"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeApplication(request *UpgradeApplicationRequest) (_result *UpgradeApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeApplicationResponse{}
	_body, _err := client.UpgradeApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstances"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RescaleDeviceServiceWithOptions(request *RescaleDeviceServiceRequest, runtime *util.RuntimeOptions) (_result *RescaleDeviceServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RescaleDeviceServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RescaleDeviceService"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RescaleDeviceService(request *RescaleDeviceServiceRequest) (_result *RescaleDeviceServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RescaleDeviceServiceResponse{}
	_body, _err := client.RescaleDeviceServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEnsServiceWithOptions(request *CreateEnsServiceRequest, runtime *util.RuntimeOptions) (_result *CreateEnsServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateEnsServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateEnsService"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEnsService(request *CreateEnsServiceRequest) (_result *CreateEnsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEnsServiceResponse{}
	_body, _err := client.CreateEnsServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportBillDetailDataWithOptions(request *ExportBillDetailDataRequest, runtime *util.RuntimeOptions) (_result *ExportBillDetailDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExportBillDetailDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExportBillDetailData"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportBillDetailData(request *ExportBillDetailDataRequest) (_result *ExportBillDetailDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportBillDetailDataResponse{}
	_body, _err := client.ExportBillDetailDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageInfosWithOptions(request *DescribeImageInfosRequest, runtime *util.RuntimeOptions) (_result *DescribeImageInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeImageInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeImageInfos"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageInfos(request *DescribeImageInfosRequest) (_result *DescribeImageInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageInfosResponse{}
	_body, _err := client.DescribeImageInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNetworkAttributeWithOptions(request *ModifyNetworkAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyNetworkAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyNetworkAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyNetworkAttribute"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNetworkAttribute(request *ModifyNetworkAttributeRequest) (_result *ModifyNetworkAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNetworkAttributeResponse{}
	_body, _err := client.ModifyNetworkAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteKeyPairsWithOptions(request *DeleteKeyPairsRequest, runtime *util.RuntimeOptions) (_result *DeleteKeyPairsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteKeyPairsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteKeyPairs"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteKeyPairs(request *DeleteKeyPairsRequest) (_result *DeleteKeyPairsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteKeyPairsResponse{}
	_body, _err := client.DeleteKeyPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthorizeSecurityGroupWithOptions(request *AuthorizeSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *AuthorizeSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AuthorizeSecurityGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AuthorizeSecurityGroup"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthorizeSecurityGroup(request *AuthorizeSecurityGroupRequest) (_result *AuthorizeSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeSecurityGroupResponse{}
	_body, _err := client.AuthorizeSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSecurityGroupWithOptions(request *DeleteSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSecurityGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSecurityGroup"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (_result *DeleteSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecurityGroupResponse{}
	_body, _err := client.DeleteSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyVSwitchAttributeWithOptions(request *ModifyVSwitchAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyVSwitchAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyVSwitchAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyVSwitchAttribute"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyVSwitchAttribute(request *ModifyVSwitchAttributeRequest) (_result *ModifyVSwitchAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVSwitchAttributeResponse{}
	_body, _err := client.ModifyVSwitchAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceAttributeWithOptions(request *ModifyInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceAttribute"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (_result *ModifyInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.ModifyInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnsEipAddressesWithOptions(request *DescribeEnsEipAddressesRequest, runtime *util.RuntimeOptions) (_result *DescribeEnsEipAddressesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEnsEipAddressesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEnsEipAddresses"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnsEipAddresses(request *DescribeEnsEipAddressesRequest) (_result *DescribeEnsEipAddressesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnsEipAddressesResponse{}
	_body, _err := client.DescribeEnsEipAddressesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RescaleApplicationWithOptions(request *RescaleApplicationRequest, runtime *util.RuntimeOptions) (_result *RescaleApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RescaleApplicationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RescaleApplication"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RescaleApplication(request *RescaleApplicationRequest) (_result *RescaleApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RescaleApplicationResponse{}
	_body, _err := client.RescaleApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityGroupsWithOptions(request *DescribeSecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecurityGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityGroups"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (_result *DescribeSecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityGroupsResponse{}
	_body, _err := client.DescribeSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserBandWidthDataWithOptions(request *DescribeUserBandWidthDataRequest, runtime *util.RuntimeOptions) (_result *DescribeUserBandWidthDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserBandWidthDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserBandWidthData"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserBandWidthData(request *DescribeUserBandWidthDataRequest) (_result *DescribeUserBandWidthDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserBandWidthDataResponse{}
	_body, _err := client.DescribeUserBandWidthDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartEpnInstanceWithOptions(request *StartEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *StartEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartEpnInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartEpnInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartEpnInstance(request *StartEpnInstanceRequest) (_result *StartEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartEpnInstanceResponse{}
	_body, _err := client.StartEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEpnBandWidthDataWithOptions(request *DescribeEpnBandWidthDataRequest, runtime *util.RuntimeOptions) (_result *DescribeEpnBandWidthDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEpnBandWidthDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEpnBandWidthData"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEpnBandWidthData(request *DescribeEpnBandWidthDataRequest) (_result *DescribeEpnBandWidthDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEpnBandWidthDataResponse{}
	_body, _err := client.DescribeEpnBandWidthDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeasurementDataWithOptions(request *DescribeMeasurementDataRequest, runtime *util.RuntimeOptions) (_result *DescribeMeasurementDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMeasurementDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMeasurementData"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeasurementData(request *DescribeMeasurementDataRequest) (_result *DescribeMeasurementDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeasurementDataResponse{}
	_body, _err := client.DescribeMeasurementDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVSwitchWithOptions(request *CreateVSwitchRequest, runtime *util.RuntimeOptions) (_result *CreateVSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVSwitch"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVSwitch(request *CreateVSwitchRequest) (_result *CreateVSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVSwitchResponse{}
	_body, _err := client.CreateVSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEpnInstanceWithOptions(request *CreateEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateEpnInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateEpnInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEpnInstance(request *CreateEpnInstanceRequest) (_result *CreateEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEpnInstanceResponse{}
	_body, _err := client.CreateEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySecurityGroupAttributeWithOptions(request *ModifySecurityGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifySecurityGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySecurityGroupAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySecurityGroupAttribute"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySecurityGroupAttribute(request *ModifySecurityGroupAttributeRequest) (_result *ModifySecurityGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySecurityGroupAttributeResponse{}
	_body, _err := client.ModifySecurityGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableResourceWithOptions(request *DescribeAvailableResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAvailableResource"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableResource(request *DescribeAvailableResourceRequest) (_result *DescribeAvailableResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DescribeAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNetworkWithOptions(request *DeleteNetworkRequest, runtime *util.RuntimeOptions) (_result *DeleteNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteNetworkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteNetwork"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNetwork(request *DeleteNetworkRequest) (_result *DeleteNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNetworkResponse{}
	_body, _err := client.DeleteNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseEipAddressWithOptions(request *ReleaseEipAddressRequest, runtime *util.RuntimeOptions) (_result *ReleaseEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseEipAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseEipAddress"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseEipAddress(request *ReleaseEipAddressRequest) (_result *ReleaseEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseEipAddressResponse{}
	_body, _err := client.ReleaseEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVSwitchesWithOptions(request *DescribeVSwitchesRequest, runtime *util.RuntimeOptions) (_result *DescribeVSwitchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVSwitches"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVSwitches(request *DescribeVSwitchesRequest) (_result *DescribeVSwitchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.DescribeVSwitchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportImageWithOptions(request *ExportImageRequest, runtime *util.RuntimeOptions) (_result *ExportImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExportImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExportImage"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportImage(request *ExportImageRequest) (_result *ExportImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportImageResponse{}
	_body, _err := client.ExportImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceSpecWithOptions(request *DescribeInstanceSpecRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceSpecResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceSpec"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceSpec(request *DescribeInstanceSpecRequest) (_result *DescribeInstanceSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSpecResponse{}
	_body, _err := client.DescribeInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceTypesWithOptions(request *DescribeInstanceTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceTypes"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (_result *DescribeInstanceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceTypesResponse{}
	_body, _err := client.DescribeInstanceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnAssociateEnsEipAddressWithOptions(request *UnAssociateEnsEipAddressRequest, runtime *util.RuntimeOptions) (_result *UnAssociateEnsEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnAssociateEnsEipAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnAssociateEnsEipAddress"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnAssociateEnsEipAddress(request *UnAssociateEnsEipAddressRequest) (_result *UnAssociateEnsEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnAssociateEnsEipAddressResponse{}
	_body, _err := client.UnAssociateEnsEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVmListWithOptions(request *GetVmListRequest, runtime *util.RuntimeOptions) (_result *GetVmListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetVmListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVmList"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVmList(request *GetVmListRequest) (_result *GetVmListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVmListResponse{}
	_body, _err := client.GetVmListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebootInstanceWithOptions(request *RebootInstanceRequest, runtime *util.RuntimeOptions) (_result *RebootInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RebootInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RebootInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebootInstance(request *RebootInstanceRequest) (_result *RebootInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootInstanceResponse{}
	_body, _err := client.RebootInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePrePaidInstanceStockWithOptions(request *DescribePrePaidInstanceStockRequest, runtime *util.RuntimeOptions) (_result *DescribePrePaidInstanceStockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePrePaidInstanceStockResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePrePaidInstanceStock"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePrePaidInstanceStock(request *DescribePrePaidInstanceStockRequest) (_result *DescribePrePaidInstanceStockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePrePaidInstanceStockResponse{}
	_body, _err := client.DescribePrePaidInstanceStockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnsRegionIdIpv6InfoWithOptions(request *DescribeEnsRegionIdIpv6InfoRequest, runtime *util.RuntimeOptions) (_result *DescribeEnsRegionIdIpv6InfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEnsRegionIdIpv6InfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEnsRegionIdIpv6Info"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnsRegionIdIpv6Info(request *DescribeEnsRegionIdIpv6InfoRequest) (_result *DescribeEnsRegionIdIpv6InfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnsRegionIdIpv6InfoResponse{}
	_body, _err := client.DescribeEnsRegionIdIpv6InfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicationWithOptions(request *DescribeApplicationRequest, runtime *util.RuntimeOptions) (_result *DescribeApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeApplicationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeApplication"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplication(request *DescribeApplicationRequest) (_result *DescribeApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApplicationResponse{}
	_body, _err := client.DescribeApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnsRegionIdResourceWithOptions(request *DescribeEnsRegionIdResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeEnsRegionIdResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEnsRegionIdResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEnsRegionIdResource"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnsRegionIdResource(request *DescribeEnsRegionIdResourceRequest) (_result *DescribeEnsRegionIdResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnsRegionIdResourceResponse{}
	_body, _err := client.DescribeEnsRegionIdResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePriceWithOptions(request *DescribePriceRequest, runtime *util.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePrice"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePrice(request *DescribePriceRequest) (_result *DescribePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePriceResponse{}
	_body, _err := client.DescribePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataDownloadURLWithOptions(request *DescribeDataDownloadURLRequest, runtime *util.RuntimeOptions) (_result *DescribeDataDownloadURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeDataDownloadURLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDataDownloadURL"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataDownloadURL(request *DescribeDataDownloadURLRequest) (_result *DescribeDataDownloadURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataDownloadURLResponse{}
	_body, _err := client.DescribeDataDownloadURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationsWithOptions(request *ListApplicationsRequest, runtime *util.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListApplicationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListApplications"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplications(request *ListApplicationsRequest) (_result *ListApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationsResponse{}
	_body, _err := client.ListApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreCreateEnsServiceWithOptions(request *PreCreateEnsServiceRequest, runtime *util.RuntimeOptions) (_result *PreCreateEnsServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PreCreateEnsServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PreCreateEnsService"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreCreateEnsService(request *PreCreateEnsServiceRequest) (_result *PreCreateEnsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PreCreateEnsServiceResponse{}
	_body, _err := client.PreCreateEnsServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBandWithdChargeTypeWithOptions(request *DescribeBandWithdChargeTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeBandWithdChargeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBandWithdChargeTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBandWithdChargeType"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBandWithdChargeType(request *DescribeBandWithdChargeTypeRequest) (_result *DescribeBandWithdChargeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBandWithdChargeTypeResponse{}
	_body, _err := client.DescribeBandWithdChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeReservedResourceWithOptions(request *DescribeReservedResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeReservedResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeReservedResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeReservedResource"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeReservedResource(request *DescribeReservedResourceRequest) (_result *DescribeReservedResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeReservedResourceResponse{}
	_body, _err := client.DescribeReservedResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCreatePrePaidInstanceResultWithOptions(request *DescribeCreatePrePaidInstanceResultRequest, runtime *util.RuntimeOptions) (_result *DescribeCreatePrePaidInstanceResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCreatePrePaidInstanceResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCreatePrePaidInstanceResult"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCreatePrePaidInstanceResult(request *DescribeCreatePrePaidInstanceResultRequest) (_result *DescribeCreatePrePaidInstanceResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCreatePrePaidInstanceResultResponse{}
	_body, _err := client.DescribeCreatePrePaidInstanceResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MigrateVmWithOptions(request *MigrateVmRequest, runtime *util.RuntimeOptions) (_result *MigrateVmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MigrateVmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MigrateVm"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MigrateVm(request *MigrateVmRequest) (_result *MigrateVmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MigrateVmResponse{}
	_body, _err := client.MigrateVmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinVSwitchesToEpnInstanceWithOptions(request *JoinVSwitchesToEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *JoinVSwitchesToEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &JoinVSwitchesToEpnInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("JoinVSwitchesToEpnInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinVSwitchesToEpnInstance(request *JoinVSwitchesToEpnInstanceRequest) (_result *JoinVSwitchesToEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinVSwitchesToEpnInstanceResponse{}
	_body, _err := client.JoinVSwitchesToEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVSwitchWithOptions(request *DeleteVSwitchRequest, runtime *util.RuntimeOptions) (_result *DeleteVSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVSwitch"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVSwitch(request *DeleteVSwitchRequest) (_result *DeleteVSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVSwitchResponse{}
	_body, _err := client.DeleteVSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceVncUrlWithOptions(request *DescribeInstanceVncUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceVncUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceVncUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceVncUrl"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceVncUrl(request *DescribeInstanceVncUrlRequest) (_result *DescribeInstanceVncUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceVncUrlResponse{}
	_body, _err := client.DescribeInstanceVncUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImagesWithOptions(request *DescribeImagesRequest, runtime *util.RuntimeOptions) (_result *DescribeImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeImages"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImages(request *DescribeImagesRequest) (_result *DescribeImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DescribeImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthorizeSecurityGroupEgressWithOptions(request *AuthorizeSecurityGroupEgressRequest, runtime *util.RuntimeOptions) (_result *AuthorizeSecurityGroupEgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AuthorizeSecurityGroupEgressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AuthorizeSecurityGroupEgress"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthorizeSecurityGroupEgress(request *AuthorizeSecurityGroupEgressRequest) (_result *AuthorizeSecurityGroupEgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeSecurityGroupEgressResponse{}
	_body, _err := client.AuthorizeSecurityGroupEgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEpnInstanceAttributeWithOptions(request *DescribeEpnInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeEpnInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEpnInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEpnInstanceAttribute"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEpnInstanceAttribute(request *DescribeEpnInstanceAttributeRequest) (_result *DescribeEpnInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEpnInstanceAttributeResponse{}
	_body, _err := client.DescribeEpnInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseInstance(request *ReleaseInstanceRequest) (_result *ReleaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.ReleaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateEnsEipAddressWithOptions(request *AssociateEnsEipAddressRequest, runtime *util.RuntimeOptions) (_result *AssociateEnsEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssociateEnsEipAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssociateEnsEipAddress"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateEnsEipAddress(request *AssociateEnsEipAddressRequest) (_result *AssociateEnsEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateEnsEipAddressResponse{}
	_body, _err := client.AssociateEnsEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushApplicationDataWithOptions(request *PushApplicationDataRequest, runtime *util.RuntimeOptions) (_result *PushApplicationDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PushApplicationDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PushApplicationData"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushApplicationData(request *PushApplicationDataRequest) (_result *PushApplicationDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushApplicationDataResponse{}
	_body, _err := client.PushApplicationDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnsRegionsWithOptions(request *DescribeEnsRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeEnsRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEnsRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEnsRegions"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnsRegions(request *DescribeEnsRegionsRequest) (_result *DescribeEnsRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnsRegionsResponse{}
	_body, _err := client.DescribeEnsRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceMonitorDataWithOptions(request *DescribeInstanceMonitorDataRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceMonitorDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceMonitorDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceMonitorData"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceMonitorData(request *DescribeInstanceMonitorDataRequest) (_result *DescribeInstanceMonitorDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceMonitorDataResponse{}
	_body, _err := client.DescribeInstanceMonitorDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnassociateEipAddressWithOptions(request *UnassociateEipAddressRequest, runtime *util.RuntimeOptions) (_result *UnassociateEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnassociateEipAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnassociateEipAddress"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnassociateEipAddress(request *UnassociateEipAddressRequest) (_result *UnassociateEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnassociateEipAddressResponse{}
	_body, _err := client.UnassociateEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataPushResultWithOptions(request *DescribeDataPushResultRequest, runtime *util.RuntimeOptions) (_result *DescribeDataPushResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDataPushResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDataPushResult"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataPushResult(request *DescribeDataPushResultRequest) (_result *DescribeDataPushResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataPushResultResponse{}
	_body, _err := client.DescribeDataPushResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LeaveSecurityGroupWithOptions(request *LeaveSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *LeaveSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &LeaveSecurityGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("LeaveSecurityGroup"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LeaveSecurityGroup(request *LeaveSecurityGroupRequest) (_result *LeaveSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LeaveSecurityGroupResponse{}
	_body, _err := client.LeaveSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceAutoRenewAttributeWithOptions(request *DescribeInstanceAutoRenewAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceAutoRenewAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceAutoRenewAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceAutoRenewAttribute"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceAutoRenewAttribute(request *DescribeInstanceAutoRenewAttributeRequest) (_result *DescribeInstanceAutoRenewAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceAutoRenewAttributeResponse{}
	_body, _err := client.DescribeInstanceAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportKeyPairWithOptions(request *ImportKeyPairRequest, runtime *util.RuntimeOptions) (_result *ImportKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImportKeyPairResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImportKeyPair"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportKeyPair(request *ImportKeyPairRequest) (_result *ImportKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportKeyPairResponse{}
	_body, _err := client.ImportKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVmWithOptions(request *DeleteVmRequest, runtime *util.RuntimeOptions) (_result *DeleteVmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVm"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVm(request *DeleteVmRequest) (_result *DeleteVmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVmResponse{}
	_body, _err := client.DeleteVmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckQuotaWithOptions(request *CheckQuotaRequest, runtime *util.RuntimeOptions) (_result *CheckQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckQuotaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckQuota"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckQuota(request *CheckQuotaRequest) (_result *CheckQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckQuotaResponse{}
	_body, _err := client.CheckQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageWithOptions(request *CreateImageRequest, runtime *util.RuntimeOptions) (_result *CreateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateImage"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImage(request *CreateImageRequest) (_result *CreateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageResponse{}
	_body, _err := client.CreateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEipInstanceWithOptions(request *CreateEipInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateEipInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateEipInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateEipInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEipInstance(request *CreateEipInstanceRequest) (_result *CreateEipInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEipInstanceResponse{}
	_body, _err := client.CreateEipInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEpnInstanceWithOptions(request *DeleteEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEpnInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEpnInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEpnInstance(request *DeleteEpnInstanceRequest) (_result *DeleteEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEpnInstanceResponse{}
	_body, _err := client.DeleteEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyImageAttributeWithOptions(request *ModifyImageAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyImageAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyImageAttribute"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) (_result *ModifyImageAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.ModifyImageAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBandwitdhByInternetChargeTypeWithOptions(request *DescribeBandwitdhByInternetChargeTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeBandwitdhByInternetChargeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBandwitdhByInternetChargeTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBandwitdhByInternetChargeType"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBandwitdhByInternetChargeType(request *DescribeBandwitdhByInternetChargeTypeRequest) (_result *DescribeBandwitdhByInternetChargeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBandwitdhByInternetChargeTypeResponse{}
	_body, _err := client.DescribeBandwitdhByInternetChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSecurityGroupWithOptions(request *CreateSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *CreateSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSecurityGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSecurityGroup"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (_result *CreateSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSecurityGroupResponse{}
	_body, _err := client.CreateSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEpnInstancesWithOptions(request *DescribeEpnInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeEpnInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEpnInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEpnInstances"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEpnInstances(request *DescribeEpnInstancesRequest) (_result *DescribeEpnInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEpnInstancesResponse{}
	_body, _err := client.DescribeEpnInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataDistResultWithOptions(request *DescribeDataDistResultRequest, runtime *util.RuntimeOptions) (_result *DescribeDataDistResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDataDistResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDataDistResult"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataDistResult(request *DescribeDataDistResultRequest) (_result *DescribeDataDistResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataDistResultResponse{}
	_body, _err := client.DescribeDataDistResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleasePostPaidInstanceWithOptions(request *ReleasePostPaidInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleasePostPaidInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleasePostPaidInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleasePostPaidInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleasePostPaidInstance(request *ReleasePostPaidInstanceRequest) (_result *ReleasePostPaidInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleasePostPaidInstanceResponse{}
	_body, _err := client.ReleasePostPaidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEPInstanceWithOptions(request *CreateEPInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateEPInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateEPInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateEPInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEPInstance(request *CreateEPInstanceRequest) (_result *CreateEPInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEPInstanceResponse{}
	_body, _err := client.CreateEPInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEpnMeasurementDataWithOptions(request *DescribeEpnMeasurementDataRequest, runtime *util.RuntimeOptions) (_result *DescribeEpnMeasurementDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEpnMeasurementDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEpnMeasurementData"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEpnMeasurementData(request *DescribeEpnMeasurementDataRequest) (_result *DescribeEpnMeasurementDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEpnMeasurementDataResponse{}
	_body, _err := client.DescribeEpnMeasurementDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicationResourceSummaryWithOptions(request *DescribeApplicationResourceSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeApplicationResourceSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeApplicationResourceSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeApplicationResourceSummary"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplicationResourceSummary(request *DescribeApplicationResourceSummaryRequest) (_result *DescribeApplicationResourceSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApplicationResourceSummaryResponse{}
	_body, _err := client.DescribeApplicationResourceSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackApplicationWithOptions(request *RollbackApplicationRequest, runtime *util.RuntimeOptions) (_result *RollbackApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RollbackApplicationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RollbackApplication"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackApplication(request *RollbackApplicationRequest) (_result *RollbackApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackApplicationResponse{}
	_body, _err := client.RollbackApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportMeasurementDataWithOptions(request *ExportMeasurementDataRequest, runtime *util.RuntimeOptions) (_result *ExportMeasurementDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExportMeasurementDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExportMeasurementData"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportMeasurementData(request *ExportMeasurementDataRequest) (_result *ExportMeasurementDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportMeasurementDataResponse{}
	_body, _err := client.ExportMeasurementDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, runtime *util.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteApplication"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (_result *DeleteApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeKeyPairsWithOptions(request *DescribeKeyPairsRequest, runtime *util.RuntimeOptions) (_result *DescribeKeyPairsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeKeyPairsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeKeyPairs"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeKeyPairs(request *DescribeKeyPairsRequest) (_result *DescribeKeyPairsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKeyPairsResponse{}
	_body, _err := client.DescribeKeyPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyEnsEipAddressAttributeWithOptions(request *ModifyEnsEipAddressAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyEnsEipAddressAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyEnsEipAddressAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyEnsEipAddressAttribute"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyEnsEipAddressAttribute(request *ModifyEnsEipAddressAttributeRequest) (_result *ModifyEnsEipAddressAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyEnsEipAddressAttributeResponse{}
	_body, _err := client.ModifyEnsEipAddressAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnsResourceUsageWithOptions(request *DescribeEnsResourceUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeEnsResourceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeEnsResourceUsageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEnsResourceUsage"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnsResourceUsage(request *DescribeEnsResourceUsageRequest) (_result *DescribeEnsResourceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnsResourceUsageResponse{}
	_body, _err := client.DescribeEnsResourceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExportImageStatusWithOptions(request *DescribeExportImageStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeExportImageStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeExportImageStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeExportImageStatus"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExportImageStatus(request *DescribeExportImageStatusRequest) (_result *DescribeExportImageStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExportImageStatusResponse{}
	_body, _err := client.DescribeExportImageStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExportImageInfoWithOptions(request *DescribeExportImageInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeExportImageInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeExportImageInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeExportImageInfo"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExportImageInfo(request *DescribeExportImageInfoRequest) (_result *DescribeExportImageInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExportImageInfoResponse{}
	_body, _err := client.DescribeExportImageInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DistApplicationDataWithOptions(request *DistApplicationDataRequest, runtime *util.RuntimeOptions) (_result *DistApplicationDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DistApplicationDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DistApplicationData"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DistApplicationData(request *DistApplicationDataRequest) (_result *DistApplicationDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DistApplicationDataResponse{}
	_body, _err := client.DistApplicationDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyImageSharePermissionWithOptions(request *ModifyImageSharePermissionRequest, runtime *util.RuntimeOptions) (_result *ModifyImageSharePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyImageSharePermissionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyImageSharePermission"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyImageSharePermission(request *ModifyImageSharePermissionRequest) (_result *ModifyImageSharePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyImageSharePermissionResponse{}
	_body, _err := client.ModifyImageSharePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityGroupAttributeWithOptions(request *DescribeSecurityGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecurityGroupAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityGroupAttribute"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityGroupAttribute(request *DescribeSecurityGroupAttributeRequest) (_result *DescribeSecurityGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityGroupAttributeResponse{}
	_body, _err := client.DescribeSecurityGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceServiceWithOptions(request *DescribeDeviceServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeDeviceServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDeviceService"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceService(request *DescribeDeviceServiceRequest) (_result *DescribeDeviceServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceServiceResponse{}
	_body, _err := client.DescribeDeviceServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnsNetSaleDistrictWithOptions(request *DescribeEnsNetSaleDistrictRequest, runtime *util.RuntimeOptions) (_result *DescribeEnsNetSaleDistrictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEnsNetSaleDistrictResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEnsNetSaleDistrict"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnsNetSaleDistrict(request *DescribeEnsNetSaleDistrictRequest) (_result *DescribeEnsNetSaleDistrictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnsNetSaleDistrictResponse{}
	_body, _err := client.DescribeEnsNetSaleDistrictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageSharePermissionWithOptions(request *DescribeImageSharePermissionRequest, runtime *util.RuntimeOptions) (_result *DescribeImageSharePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeImageSharePermissionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeImageSharePermission"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageSharePermission(request *DescribeImageSharePermissionRequest) (_result *DescribeImageSharePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageSharePermissionResponse{}
	_body, _err := client.DescribeImageSharePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyEpnInstanceWithOptions(request *ModifyEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *ModifyEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyEpnInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyEpnInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyEpnInstance(request *ModifyEpnInstanceRequest) (_result *ModifyEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyEpnInstanceResponse{}
	_body, _err := client.ModifyEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnsNetDistrictWithOptions(request *DescribeEnsNetDistrictRequest, runtime *util.RuntimeOptions) (_result *DescribeEnsNetDistrictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEnsNetDistrictResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEnsNetDistrict"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnsNetDistrict(request *DescribeEnsNetDistrictRequest) (_result *DescribeEnsNetDistrictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnsNetDistrictResponse{}
	_body, _err := client.DescribeEnsNetDistrictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNetworksWithOptions(request *DescribeNetworksRequest, runtime *util.RuntimeOptions) (_result *DescribeNetworksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNetworksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNetworks"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNetworks(request *DescribeNetworksRequest) (_result *DescribeNetworksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNetworksResponse{}
	_body, _err := client.DescribeNetworksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunServiceScheduleWithOptions(request *RunServiceScheduleRequest, runtime *util.RuntimeOptions) (_result *RunServiceScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RunServiceScheduleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RunServiceSchedule"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunServiceSchedule(request *RunServiceScheduleRequest) (_result *RunServiceScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunServiceScheduleResponse{}
	_body, _err := client.RunServiceScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveVSwitchesFromEpnInstanceWithOptions(request *RemoveVSwitchesFromEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *RemoveVSwitchesFromEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveVSwitchesFromEpnInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveVSwitchesFromEpnInstance"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveVSwitchesFromEpnInstance(request *RemoveVSwitchesFromEpnInstanceRequest) (_result *RemoveVSwitchesFromEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveVSwitchesFromEpnInstanceResponse{}
	_body, _err := client.RemoveVSwitchesFromEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SchedulePodWithOptions(request *SchedulePodRequest, runtime *util.RuntimeOptions) (_result *SchedulePodResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SchedulePodResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SchedulePod"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SchedulePod(request *SchedulePodRequest) (_result *SchedulePodResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SchedulePodResponse{}
	_body, _err := client.SchedulePodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNetworkInterfacesWithOptions(request *DescribeNetworkInterfacesRequest, runtime *util.RuntimeOptions) (_result *DescribeNetworkInterfacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNetworkInterfacesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNetworkInterfaces"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNetworkInterfaces(request *DescribeNetworkInterfacesRequest) (_result *DescribeNetworkInterfacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNetworkInterfacesResponse{}
	_body, _err := client.DescribeNetworkInterfacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AllocateEipAddressWithOptions(request *AllocateEipAddressRequest, runtime *util.RuntimeOptions) (_result *AllocateEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AllocateEipAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AllocateEipAddress"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocateEipAddress(request *AllocateEipAddressRequest) (_result *AllocateEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateEipAddressResponse{}
	_body, _err := client.AllocateEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateKeyPairWithOptions(request *CreateKeyPairRequest, runtime *util.RuntimeOptions) (_result *CreateKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateKeyPairResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateKeyPair"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateKeyPair(request *CreateKeyPairRequest) (_result *CreateKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateKeyPairResponse{}
	_body, _err := client.CreateKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableResourceInfoWithOptions(request *DescribeAvailableResourceInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAvailableResourceInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAvailableResourceInfo"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableResourceInfo(request *DescribeAvailableResourceInfoRequest) (_result *DescribeAvailableResourceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourceInfoResponse{}
	_body, _err := client.DescribeAvailableResourceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVmAndSaveStockWithOptions(request *CreateVmAndSaveStockRequest, runtime *util.RuntimeOptions) (_result *CreateVmAndSaveStockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVmAndSaveStockResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVmAndSaveStock"), tea.String("2017-11-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVmAndSaveStock(request *CreateVmAndSaveStockRequest) (_result *CreateVmAndSaveStockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVmAndSaveStockResponse{}
	_body, _err := client.CreateVmAndSaveStockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
