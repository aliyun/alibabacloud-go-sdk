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

type CreateDatabaseRequest struct {
	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 数据库名称。 不能使用某些预留关键字，如 test、mysql。
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// 数据库描述信息。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 数据库编码方式。 详细参见DescribeCharset返回的Charset字段。
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CreateDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseRequest) SetClientToken(v string) *CreateDatabaseRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDatabaseRequest) SetDatabaseName(v string) *CreateDatabaseRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateDatabaseRequest) SetDescription(v string) *CreateDatabaseRequest {
	s.Description = &v
	return s
}

func (s *CreateDatabaseRequest) SetEncoding(v string) *CreateDatabaseRequest {
	s.Encoding = &v
	return s
}

func (s *CreateDatabaseRequest) SetInstanceId(v string) *CreateDatabaseRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDatabaseRequest) SetTenantId(v string) *CreateDatabaseRequest {
	s.TenantId = &v
	return s
}

type CreateDatabaseResponseBody struct {
	// 数据库名称。
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatabaseResponseBody) SetDatabaseName(v string) *CreateDatabaseResponseBody {
	s.DatabaseName = &v
	return s
}

func (s *CreateDatabaseResponseBody) SetRequestId(v string) *CreateDatabaseResponseBody {
	s.RequestId = &v
	return s
}

type CreateDatabaseResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseResponse) GoString() string {
	return s.String()
}

func (s *CreateDatabaseResponse) SetHeaders(v map[string]*string) *CreateDatabaseResponse {
	s.Headers = v
	return s
}

func (s *CreateDatabaseResponse) SetBody(v *CreateDatabaseResponseBody) *CreateDatabaseResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	// 是否要自动续费。当参数ChargeType取值PrePaid时生效。取值范围：  true：自动续费。 false（默认）：不自动续费。
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// 每次自动续费的时长，当参数AutoRenew取值True时，该参数为必填参数。  PeriodUnit为Week时，AutoRenewPeriod取值范围为{"1", "2", "3"}。  PeriodUnit为Month时，AutoRenewPeriod取值范围为{"1", "2", "3", "6", "12"}。
	AutoRenewPeriod *int64 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// 实例的付费方式。取值范围：  PrePay：包年包月。选择该类付费方式时，您必须确认自己的账号支持余额支付/信用支付，否则将返回 InvalidPayMethod的错误提示。 PostPay（默认）：按量付费。其默认按小时来计费
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// 存储空间大小，单位GB。  存储空间的限制根据集群规格不同而不同，具体如下：  - 8C32GB：100GB~10000GB  - 14C70GB：200GB~10000GB  - 30C180GB：400GB~10000GB  - 62C400G：800GB-10000GB。  各套餐的存储空间默认值为其最小值。
	DiskSize *int64 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// 集群规格信息。  当前支持四种套餐：  - 8C32GB：8核 32GB  - 14C70GB（默认）：14核 70GB  - 30C180GB：30核 180GB  - 62C400GB：62核 400GB
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// 集群名称。
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 购买资源的时长，单位由PeriodUnit指定。当参数InstanceChargeType取值为PrePaid时才生效且为必选值。一旦指定了DedicatedHostId，则取值范围不能超过专有宿主机的订阅时长。取值范围：  PeriodUnit=Week时，Period取值：{“1”, “2”, “3”, “4”}。 PeriodUnit=Month时，Period取值：{“1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”, ”48”, ”60”}。
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// 购买资源的时长。  包年包月取值范围： Month。 默认值：包年包月为Month，按量计费，默认周期为Hour。
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// 实例所在的企业资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Oceanbase集群的系列  - normal（默认）：高可用版本  - basic：基础版本（API暂不支持，请页面购买）
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// 实例所属的可用区ID。更多详情，请参见DescribeZones获取可用区列表。
	Zones *string `json:"Zones,omitempty" xml:"Zones,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v int64) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetDiskSize(v int64) *CreateInstanceRequest {
	s.DiskSize = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceClass(v string) *CreateInstanceRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v int64) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriodUnit(v string) *CreateInstanceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetSeries(v string) *CreateInstanceRequest {
	s.Series = &v
	return s
}

func (s *CreateInstanceRequest) SetZones(v string) *CreateInstanceRequest {
	s.Zones = &v
	return s
}

type CreateInstanceResponseBody struct {
	// 返回数据
	Data []*CreateInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetData(v []*CreateInstanceResponseBodyData) *CreateInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceResponseBodyData struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 订单ID。该参数只有创建包年包月ECS实例（请求参数InstanceChargeType=PrePaid）时有返回值。
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// 资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyData) SetInstanceId(v string) *CreateInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetOrderId(v string) *CreateInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetResourceGroupId(v string) *CreateInstanceResponseBodyData {
	s.ResourceGroupId = &v
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

type CreateSecurityIpGroupRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// IP安全白名单组的组名。 由小写英文字符开头，由小写英文字符或者数字结尾，只能包含小写英文字符，数字和下划线，长度在 2-32 个字符之间。
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	// IP安全白名单列表。 其为一个Json格式的数组，数组中每个对象为一个IP字符串或者IP段。最多可设置 40 个。
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s CreateSecurityIpGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityIpGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityIpGroupRequest) SetInstanceId(v string) *CreateSecurityIpGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSecurityIpGroupRequest) SetSecurityIpGroupName(v string) *CreateSecurityIpGroupRequest {
	s.SecurityIpGroupName = &v
	return s
}

func (s *CreateSecurityIpGroupRequest) SetSecurityIps(v string) *CreateSecurityIpGroupRequest {
	s.SecurityIps = &v
	return s
}

type CreateSecurityIpGroupResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 白名单分组信息。
	SecurityIpGroup *CreateSecurityIpGroupResponseBodySecurityIpGroup `json:"SecurityIpGroup,omitempty" xml:"SecurityIpGroup,omitempty" type:"Struct"`
}

func (s CreateSecurityIpGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityIpGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityIpGroupResponseBody) SetRequestId(v string) *CreateSecurityIpGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityIpGroupResponseBody) SetSecurityIpGroup(v *CreateSecurityIpGroupResponseBodySecurityIpGroup) *CreateSecurityIpGroupResponseBody {
	s.SecurityIpGroup = v
	return s
}

type CreateSecurityIpGroupResponseBodySecurityIpGroup struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 白名单分组名称。
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	// IP白名单分组中的IP地址或地址段。其为一个Json格式的数组，数组中每个对象为一个IP字符串或者IP段。
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s CreateSecurityIpGroupResponseBodySecurityIpGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityIpGroupResponseBodySecurityIpGroup) GoString() string {
	return s.String()
}

func (s *CreateSecurityIpGroupResponseBodySecurityIpGroup) SetInstanceId(v string) *CreateSecurityIpGroupResponseBodySecurityIpGroup {
	s.InstanceId = &v
	return s
}

func (s *CreateSecurityIpGroupResponseBodySecurityIpGroup) SetSecurityIpGroupName(v string) *CreateSecurityIpGroupResponseBodySecurityIpGroup {
	s.SecurityIpGroupName = &v
	return s
}

func (s *CreateSecurityIpGroupResponseBodySecurityIpGroup) SetSecurityIps(v string) *CreateSecurityIpGroupResponseBodySecurityIpGroup {
	s.SecurityIps = &v
	return s
}

type CreateSecurityIpGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSecurityIpGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSecurityIpGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityIpGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityIpGroupResponse) SetHeaders(v map[string]*string) *CreateSecurityIpGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityIpGroupResponse) SetBody(v *CreateSecurityIpGroupResponseBody) *CreateSecurityIpGroupResponse {
	s.Body = v
	return s
}

type CreateTenantRequest struct {
	// 字符集。 详细参见：DescribeCharset。
	Charset *string `json:"Charset,omitempty" xml:"Charset,omitempty"`
	// 租户的CPU大小，单位：核数（Cores）
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// 租户描述信息。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户内存大小，单位GB。
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// 租户的主可用区。 其为集群部署可用区中的一个。
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// 租户模式。 当前支持Oracle模式（Oracle）、MySQL模式（MySQL） 详细参见：DescribeInstanceTenantModes。
	TenantMode *string `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
	// 租户名称。 长度为2~20 个字符，支持英文字母、数字和下划线，区分大小写，必须以字母或下划线开头。 不可设置为 sys。
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// 租户所在时区。 详细参见：DescribeTimeZones。
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// 租户的资源分布节点数。 其与集群的部署模式相耦合，如集群模式为2-2-2，则最后分布节点数最多为2个。
	UnitNum *int32 `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
	// 虚拟交换机（VSwitch） ID。 如果没有合适的 VSwitch，请根据页面提示创建一个，详情参见 使用交换机。
	UserVSwitchId *string `json:"UserVSwitchId,omitempty" xml:"UserVSwitchId,omitempty"`
	// 专有网络（VPC） ID。 如果没有合适的 VPC，请根据页面提示创建一个 VPC，详情参见 什么是专有网络
	UserVpcId *string `json:"UserVpcId,omitempty" xml:"UserVpcId,omitempty"`
}

func (s CreateTenantRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantRequest) SetCharset(v string) *CreateTenantRequest {
	s.Charset = &v
	return s
}

func (s *CreateTenantRequest) SetCpu(v int32) *CreateTenantRequest {
	s.Cpu = &v
	return s
}

func (s *CreateTenantRequest) SetDescription(v string) *CreateTenantRequest {
	s.Description = &v
	return s
}

func (s *CreateTenantRequest) SetInstanceId(v string) *CreateTenantRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTenantRequest) SetMemory(v int32) *CreateTenantRequest {
	s.Memory = &v
	return s
}

func (s *CreateTenantRequest) SetPrimaryZone(v string) *CreateTenantRequest {
	s.PrimaryZone = &v
	return s
}

func (s *CreateTenantRequest) SetTenantMode(v string) *CreateTenantRequest {
	s.TenantMode = &v
	return s
}

func (s *CreateTenantRequest) SetTenantName(v string) *CreateTenantRequest {
	s.TenantName = &v
	return s
}

func (s *CreateTenantRequest) SetTimeZone(v string) *CreateTenantRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateTenantRequest) SetUnitNum(v int32) *CreateTenantRequest {
	s.UnitNum = &v
	return s
}

func (s *CreateTenantRequest) SetUserVSwitchId(v string) *CreateTenantRequest {
	s.UserVSwitchId = &v
	return s
}

func (s *CreateTenantRequest) SetUserVpcId(v string) *CreateTenantRequest {
	s.UserVpcId = &v
	return s
}

type CreateTenantResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CreateTenantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTenantResponseBody) SetRequestId(v string) *CreateTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTenantResponseBody) SetTenantId(v string) *CreateTenantResponseBody {
	s.TenantId = &v
	return s
}

type CreateTenantResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTenantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTenantResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantResponse) GoString() string {
	return s.String()
}

func (s *CreateTenantResponse) SetHeaders(v map[string]*string) *CreateTenantResponse {
	s.Headers = v
	return s
}

func (s *CreateTenantResponse) SetBody(v *CreateTenantResponseBody) *CreateTenantResponse {
	s.Body = v
	return s
}

type CreateTenantReadOnlyConnectionRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 可用区的ZoneId。 详细参见DescribeInstance中的AvailableZones。
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTenantReadOnlyConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantReadOnlyConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantReadOnlyConnectionRequest) SetClientToken(v string) *CreateTenantReadOnlyConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTenantReadOnlyConnectionRequest) SetInstanceId(v string) *CreateTenantReadOnlyConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTenantReadOnlyConnectionRequest) SetTenantId(v string) *CreateTenantReadOnlyConnectionRequest {
	s.TenantId = &v
	return s
}

func (s *CreateTenantReadOnlyConnectionRequest) SetZoneId(v string) *CreateTenantReadOnlyConnectionRequest {
	s.ZoneId = &v
	return s
}

type CreateTenantReadOnlyConnectionResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTenantReadOnlyConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantReadOnlyConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTenantReadOnlyConnectionResponseBody) SetRequestId(v string) *CreateTenantReadOnlyConnectionResponseBody {
	s.RequestId = &v
	return s
}

type CreateTenantReadOnlyConnectionResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTenantReadOnlyConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTenantReadOnlyConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantReadOnlyConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateTenantReadOnlyConnectionResponse) SetHeaders(v map[string]*string) *CreateTenantReadOnlyConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateTenantReadOnlyConnectionResponse) SetBody(v *CreateTenantReadOnlyConnectionResponseBody) *CreateTenantReadOnlyConnectionResponse {
	s.Body = v
	return s
}

type CreateTenantUserRequest struct {
	// 数据库描述信息。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 用户的角色信息。 对于Oracle模式：无需填写。 对于MySQL模式：超级账户默认为ALL PRIVILEGES，无需填写。 普通用户填写账户的信息，其为Json格式的字符串，默认为一个数组，数组内包含schema（Oracle模式）或Database（MySQL模式)信息和角色信息（Role）。 角色包含以下几类： 读写权限（ReadWrite）：ALL PRIVILEGES ； 只读权限（ReadOnly）：SELECT DDL权限（DDL）：CREATE,DROP,ALTER,SHOW VIEW,CREATE VIEW DML权限：SELECT,INSERT,UPDATE,DELETE,SHOW VIEW； DML权限（DML）：SELECT,INSERT,UPDATE,DELETE,SHOW VIEW。
	Roles *string `json:"Roles,omitempty" xml:"Roles,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 数据库账号名称。 不能使用某些预留关键字，如 SYS、root等。
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// 数据库账号访问密码。 必须包含大写英文字符、小写英文字符、数字、特殊字符占三种，长度为 10-32 位； 特殊字符为：!@#$%^&* ()_ +-=
	UserPassword *string `json:"UserPassword,omitempty" xml:"UserPassword,omitempty"`
	// 数据库账户的类型 Admin：超级账户 Normal：普通账户
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s CreateTenantUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantUserRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantUserRequest) SetDescription(v string) *CreateTenantUserRequest {
	s.Description = &v
	return s
}

func (s *CreateTenantUserRequest) SetInstanceId(v string) *CreateTenantUserRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTenantUserRequest) SetRoles(v string) *CreateTenantUserRequest {
	s.Roles = &v
	return s
}

func (s *CreateTenantUserRequest) SetTenantId(v string) *CreateTenantUserRequest {
	s.TenantId = &v
	return s
}

func (s *CreateTenantUserRequest) SetUserName(v string) *CreateTenantUserRequest {
	s.UserName = &v
	return s
}

func (s *CreateTenantUserRequest) SetUserPassword(v string) *CreateTenantUserRequest {
	s.UserPassword = &v
	return s
}

func (s *CreateTenantUserRequest) SetUserType(v string) *CreateTenantUserRequest {
	s.UserType = &v
	return s
}

type CreateTenantUserResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 租户中的数据库账号信息
	TenantUser []*CreateTenantUserResponseBodyTenantUser `json:"TenantUser,omitempty" xml:"TenantUser,omitempty" type:"Repeated"`
}

func (s CreateTenantUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTenantUserResponseBody) SetRequestId(v string) *CreateTenantUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTenantUserResponseBody) SetTenantUser(v []*CreateTenantUserResponseBodyTenantUser) *CreateTenantUserResponseBody {
	s.TenantUser = v
	return s
}

type CreateTenantUserResponseBodyTenantUser struct {
	// 账户对应的角色权限信息
	Roles []*CreateTenantUserResponseBodyTenantUserRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// 数据库账号名称。
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// 数据库账号的状态。 - Locked：锁定 - ONLINE：解锁 创建完的账户默认为ONLINE状态
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
	// 数据库账号类型。 - Admin：超级账户 - Normal：普通账户
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s CreateTenantUserResponseBodyTenantUser) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantUserResponseBodyTenantUser) GoString() string {
	return s.String()
}

func (s *CreateTenantUserResponseBodyTenantUser) SetRoles(v []*CreateTenantUserResponseBodyTenantUserRoles) *CreateTenantUserResponseBodyTenantUser {
	s.Roles = v
	return s
}

func (s *CreateTenantUserResponseBodyTenantUser) SetUserName(v string) *CreateTenantUserResponseBodyTenantUser {
	s.UserName = &v
	return s
}

func (s *CreateTenantUserResponseBodyTenantUser) SetUserStatus(v string) *CreateTenantUserResponseBodyTenantUser {
	s.UserStatus = &v
	return s
}

func (s *CreateTenantUserResponseBodyTenantUser) SetUserType(v string) *CreateTenantUserResponseBodyTenantUser {
	s.UserType = &v
	return s
}

type CreateTenantUserResponseBodyTenantUserRoles struct {
	// 数据库名称
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// 拥有的角色。 对于Oracle模式，角色为schema级别，其可分为 - ReadWrite：读写权限，包括CREATE TABLE CREATE VIEW CREATE PROCEDURE CREATE SYNONYM CREATE SEQUENCE CREATE TRIGGER CREATE TYPE CREATE SESSION EXECUTE ANY PROCEDURE CREATE ANY OUTLINE ALTER ANY OUTLINE DROP ANY OUTLINE CREATE ANY PROCEDURE ALTER ANY PROCEDURE DROP ANY PROCEDURE CREATE ANY SEQUENCE ALTER ANY SEQUENCE DROP ANY SEQUENCE CREATE ANY TYPE ALTER ANY TYPE DROP ANY TYPE SYSKM CREATE ANY TRIGGER ALTER ANY TRIGGER DROP ANY TRIGGER CREATE PROFILE ALTER PROFILE DROP PROFILE； - ReadOnly：只读权限，SELECT
	// 对于MySQL模式，角色为数据库（Database）级别，其有以下几类： - ReadWrite：读写权限，包括ALL PRIVILEGES； - ReadOnly：只读权限，包括SELECT - DDL: DDL权限，包括CREATE,DROP,ALTER,SHOW VIEW,CREATE VIEW - DML: DML权限，包括SELECT,INSERT,UPDATE,DELETE,SHOW VIEW。
	// * 另外Oracle账户对自己的schema有默认的读写权限，这里不会列出。
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s CreateTenantUserResponseBodyTenantUserRoles) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantUserResponseBodyTenantUserRoles) GoString() string {
	return s.String()
}

func (s *CreateTenantUserResponseBodyTenantUserRoles) SetDatabase(v string) *CreateTenantUserResponseBodyTenantUserRoles {
	s.Database = &v
	return s
}

func (s *CreateTenantUserResponseBodyTenantUserRoles) SetRole(v string) *CreateTenantUserResponseBodyTenantUserRoles {
	s.Role = &v
	return s
}

type CreateTenantUserResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTenantUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTenantUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantUserResponse) GoString() string {
	return s.String()
}

func (s *CreateTenantUserResponse) SetHeaders(v map[string]*string) *CreateTenantUserResponse {
	s.Headers = v
	return s
}

func (s *CreateTenantUserResponse) SetBody(v *CreateTenantUserResponseBody) *CreateTenantUserResponse {
	s.Body = v
	return s
}

type DeleteDatabasesRequest struct {
	// 数据库名称列表。 其为Json格式的数组，数组中每个对象都为数据库名称的字符串。
	DatabaseNames *string `json:"DatabaseNames,omitempty" xml:"DatabaseNames,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeleteDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabasesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabasesRequest) SetDatabaseNames(v string) *DeleteDatabasesRequest {
	s.DatabaseNames = &v
	return s
}

func (s *DeleteDatabasesRequest) SetInstanceId(v string) *DeleteDatabasesRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDatabasesRequest) SetTenantId(v string) *DeleteDatabasesRequest {
	s.TenantId = &v
	return s
}

type DeleteDatabasesResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatabasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatabasesResponseBody) SetRequestId(v string) *DeleteDatabasesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDatabasesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabasesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatabasesResponse) SetHeaders(v map[string]*string) *DeleteDatabasesResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatabasesResponse) SetBody(v *DeleteDatabasesResponseBody) *DeleteDatabasesResponse {
	s.Body = v
	return s
}

type DeleteInstancesRequest struct {
	// 集群删除后的备份保留策略。取值如下： - receive_all：保留全部备份集; - delete_all：删除全部备份集； - receive_last：保留最后一个备份集。 默认值为delete_all。
	BackupRetainMode *string `json:"BackupRetainMode,omitempty" xml:"BackupRetainMode,omitempty"`
	// 要删除的集群ID。格式为son数组的字符串。
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s DeleteInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstancesRequest) SetBackupRetainMode(v string) *DeleteInstancesRequest {
	s.BackupRetainMode = &v
	return s
}

func (s *DeleteInstancesRequest) SetInstanceIds(v string) *DeleteInstancesRequest {
	s.InstanceIds = &v
	return s
}

type DeleteInstancesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstancesResponseBody) SetRequestId(v string) *DeleteInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstancesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstancesResponse) SetHeaders(v map[string]*string) *DeleteInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstancesResponse) SetBody(v *DeleteInstancesResponseBody) *DeleteInstancesResponse {
	s.Body = v
	return s
}

type DeleteSecurityIpGroupRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// IP安全白名单组的组名。 由小写英文字符开头，由小写英文字符或者数字结尾，只能包含小写英文字符，数字和下划线，长度在 2-32 个字符之间。
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
}

func (s DeleteSecurityIpGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityIpGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIpGroupRequest) SetInstanceId(v string) *DeleteSecurityIpGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSecurityIpGroupRequest) SetSecurityIpGroupName(v string) *DeleteSecurityIpGroupRequest {
	s.SecurityIpGroupName = &v
	return s
}

type DeleteSecurityIpGroupResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 删除的IP白名单分组信息。
	SecurityIpGroup *DeleteSecurityIpGroupResponseBodySecurityIpGroup `json:"SecurityIpGroup,omitempty" xml:"SecurityIpGroup,omitempty" type:"Struct"`
}

func (s DeleteSecurityIpGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityIpGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIpGroupResponseBody) SetRequestId(v string) *DeleteSecurityIpGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityIpGroupResponseBody) SetSecurityIpGroup(v *DeleteSecurityIpGroupResponseBodySecurityIpGroup) *DeleteSecurityIpGroupResponseBody {
	s.SecurityIpGroup = v
	return s
}

type DeleteSecurityIpGroupResponseBodySecurityIpGroup struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 删除的IP安全白名单组的组名。
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
}

func (s DeleteSecurityIpGroupResponseBodySecurityIpGroup) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityIpGroupResponseBodySecurityIpGroup) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIpGroupResponseBodySecurityIpGroup) SetInstanceId(v string) *DeleteSecurityIpGroupResponseBodySecurityIpGroup {
	s.InstanceId = &v
	return s
}

func (s *DeleteSecurityIpGroupResponseBodySecurityIpGroup) SetSecurityIpGroupName(v string) *DeleteSecurityIpGroupResponseBodySecurityIpGroup {
	s.SecurityIpGroupName = &v
	return s
}

type DeleteSecurityIpGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSecurityIpGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSecurityIpGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityIpGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIpGroupResponse) SetHeaders(v map[string]*string) *DeleteSecurityIpGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityIpGroupResponse) SetBody(v *DeleteSecurityIpGroupResponseBody) *DeleteSecurityIpGroupResponse {
	s.Body = v
	return s
}

type DeleteTenantUsersRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 用户名及具备的角色列表。
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s DeleteTenantUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTenantUsersRequest) GoString() string {
	return s.String()
}

func (s *DeleteTenantUsersRequest) SetInstanceId(v string) *DeleteTenantUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteTenantUsersRequest) SetTenantId(v string) *DeleteTenantUsersRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteTenantUsersRequest) SetUsers(v string) *DeleteTenantUsersRequest {
	s.Users = &v
	return s
}

type DeleteTenantUsersResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTenantUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTenantUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTenantUsersResponseBody) SetRequestId(v string) *DeleteTenantUsersResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTenantUsersResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTenantUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTenantUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTenantUsersResponse) GoString() string {
	return s.String()
}

func (s *DeleteTenantUsersResponse) SetHeaders(v map[string]*string) *DeleteTenantUsersResponse {
	s.Headers = v
	return s
}

func (s *DeleteTenantUsersResponse) SetBody(v *DeleteTenantUsersResponseBody) *DeleteTenantUsersResponse {
	s.Body = v
	return s
}

type DeleteTenantsRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户的ID列表。 其为Json格式的数组，数组中每个对象都为租户名称的字符串。
	TenantIds *string `json:"TenantIds,omitempty" xml:"TenantIds,omitempty"`
}

func (s DeleteTenantsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTenantsRequest) GoString() string {
	return s.String()
}

func (s *DeleteTenantsRequest) SetInstanceId(v string) *DeleteTenantsRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteTenantsRequest) SetTenantIds(v string) *DeleteTenantsRequest {
	s.TenantIds = &v
	return s
}

type DeleteTenantsResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 租户的ID列表。
	TenantIds []*string `json:"TenantIds,omitempty" xml:"TenantIds,omitempty" type:"Repeated"`
}

func (s DeleteTenantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTenantsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTenantsResponseBody) SetRequestId(v string) *DeleteTenantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTenantsResponseBody) SetTenantIds(v []*string) *DeleteTenantsResponseBody {
	s.TenantIds = v
	return s
}

type DeleteTenantsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTenantsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTenantsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTenantsResponse) GoString() string {
	return s.String()
}

func (s *DeleteTenantsResponse) SetHeaders(v map[string]*string) *DeleteTenantsResponse {
	s.Headers = v
	return s
}

func (s *DeleteTenantsResponse) SetBody(v *DeleteTenantsResponseBody) *DeleteTenantsResponse {
	s.Body = v
	return s
}

type DescribeAnomalySQLListRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// 数据库名称
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 过滤条件
	FilterCondition map[string]interface{} `json:"FilterCondition,omitempty" xml:"FilterCondition,omitempty"`
	// 节点ip
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// 分页查询的页码。 起始值：1 默认值：1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时设置的每页行数。 最大值：100 默认值：10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// SQLID
	SQLId *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// 关键字查询
	SearchKeyWord *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	// 参数查询
	SearchParameter *string `json:"SearchParameter,omitempty" xml:"SearchParameter,omitempty"`
	// 查询规则
	SearchRule *string `json:"SearchRule,omitempty" xml:"SearchRule,omitempty"`
	// 查询值
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	// 排序列
	SortColumn *string `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	// 排序规则
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 租户ID
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAnomalySQLListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnomalySQLListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnomalySQLListRequest) SetAcceptLanguage(v string) *DescribeAnomalySQLListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetDbName(v string) *DescribeAnomalySQLListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetEndTime(v string) *DescribeAnomalySQLListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetFilterCondition(v map[string]interface{}) *DescribeAnomalySQLListRequest {
	s.FilterCondition = v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetNodeIp(v string) *DescribeAnomalySQLListRequest {
	s.NodeIp = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetPageNumber(v int32) *DescribeAnomalySQLListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetPageSize(v int32) *DescribeAnomalySQLListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetSQLId(v string) *DescribeAnomalySQLListRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetSearchKeyWord(v string) *DescribeAnomalySQLListRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetSearchParameter(v string) *DescribeAnomalySQLListRequest {
	s.SearchParameter = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetSearchRule(v string) *DescribeAnomalySQLListRequest {
	s.SearchRule = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetSearchValue(v string) *DescribeAnomalySQLListRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetSortColumn(v string) *DescribeAnomalySQLListRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetSortOrder(v string) *DescribeAnomalySQLListRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetStartTime(v string) *DescribeAnomalySQLListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetTenantId(v string) *DescribeAnomalySQLListRequest {
	s.TenantId = &v
	return s
}

type DescribeAnomalySQLListShrinkRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// 数据库名称
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 过滤条件
	FilterConditionShrink *string `json:"FilterCondition,omitempty" xml:"FilterCondition,omitempty"`
	// 节点ip
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// 分页查询的页码。 起始值：1 默认值：1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时设置的每页行数。 最大值：100 默认值：10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// SQLID
	SQLId *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// 关键字查询
	SearchKeyWord *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	// 参数查询
	SearchParameter *string `json:"SearchParameter,omitempty" xml:"SearchParameter,omitempty"`
	// 查询规则
	SearchRule *string `json:"SearchRule,omitempty" xml:"SearchRule,omitempty"`
	// 查询值
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	// 排序列
	SortColumn *string `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	// 排序规则
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 租户ID
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAnomalySQLListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnomalySQLListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnomalySQLListShrinkRequest) SetAcceptLanguage(v string) *DescribeAnomalySQLListShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetDbName(v string) *DescribeAnomalySQLListShrinkRequest {
	s.DbName = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetEndTime(v string) *DescribeAnomalySQLListShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetFilterConditionShrink(v string) *DescribeAnomalySQLListShrinkRequest {
	s.FilterConditionShrink = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetNodeIp(v string) *DescribeAnomalySQLListShrinkRequest {
	s.NodeIp = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetPageNumber(v int32) *DescribeAnomalySQLListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetPageSize(v int32) *DescribeAnomalySQLListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetSQLId(v string) *DescribeAnomalySQLListShrinkRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetSearchKeyWord(v string) *DescribeAnomalySQLListShrinkRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetSearchParameter(v string) *DescribeAnomalySQLListShrinkRequest {
	s.SearchParameter = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetSearchRule(v string) *DescribeAnomalySQLListShrinkRequest {
	s.SearchRule = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetSearchValue(v string) *DescribeAnomalySQLListShrinkRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetSortColumn(v string) *DescribeAnomalySQLListShrinkRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetSortOrder(v string) *DescribeAnomalySQLListShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetStartTime(v string) *DescribeAnomalySQLListShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetTenantId(v string) *DescribeAnomalySQLListShrinkRequest {
	s.TenantId = &v
	return s
}

type DescribeAnomalySQLListResponseBody struct {
	// 可疑SQL列表
	AnomalySQLList []*DescribeAnomalySQLListResponseBodyAnomalySQLList `json:"AnomalySQLList,omitempty" xml:"AnomalySQLList,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAnomalySQLListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnomalySQLListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAnomalySQLListResponseBody) SetAnomalySQLList(v []*DescribeAnomalySQLListResponseBodyAnomalySQLList) *DescribeAnomalySQLListResponseBody {
	s.AnomalySQLList = v
	return s
}

func (s *DescribeAnomalySQLListResponseBody) SetRequestId(v string) *DescribeAnomalySQLListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBody) SetTotalCount(v int64) *DescribeAnomalySQLListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAnomalySQLListResponseBodyAnomalySQLList struct {
	// 平均CPU时间
	CpuTime *int64 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	// 数据库名
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// 诊断信息
	Diagnosis *string `json:"Diagnosis,omitempty" xml:"Diagnosis,omitempty"`
	// 诊断规则
	DiagnosisRule *string `json:"DiagnosisRule,omitempty" xml:"DiagnosisRule,omitempty"`
	// 执行次数
	Executions *int64 `json:"Executions,omitempty" xml:"Executions,omitempty"`
	// 返回数据序号
	Key *int64 `json:"Key,omitempty" xml:"Key,omitempty"`
	// 请求时间
	RequestTime *int64 `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	// 请求时间（零时区）
	RequestTimeUTCString *string `json:"RequestTimeUTCString,omitempty" xml:"RequestTimeUTCString,omitempty"`
	// SQLID
	SQLId *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// sql文本
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// 建议
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// 用户名
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeAnomalySQLListResponseBodyAnomalySQLList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnomalySQLListResponseBodyAnomalySQLList) GoString() string {
	return s.String()
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetCpuTime(v int64) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.CpuTime = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetDbName(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.DbName = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetDiagnosis(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.Diagnosis = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetDiagnosisRule(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.DiagnosisRule = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetExecutions(v int64) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.Executions = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetKey(v int64) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.Key = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetRequestTime(v int64) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.RequestTime = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetRequestTimeUTCString(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.RequestTimeUTCString = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetSQLId(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.SQLId = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetSQLText(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.SQLText = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetSuggestion(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.Suggestion = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetUserName(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.UserName = &v
	return s
}

type DescribeAnomalySQLListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAnomalySQLListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAnomalySQLListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnomalySQLListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAnomalySQLListResponse) SetHeaders(v map[string]*string) *DescribeAnomalySQLListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAnomalySQLListResponse) SetBody(v *DescribeAnomalySQLListResponseBody) *DescribeAnomalySQLListResponse {
	s.Body = v
	return s
}

type DescribeAvailableCpuResourceRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 用户权限修改类型。 可支持以下几种方式： update：全量更新权限，默认值； add：新增权限； delete：删除权限。 默认值：update。
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAvailableCpuResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableCpuResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCpuResourceRequest) SetInstanceId(v string) *DescribeAvailableCpuResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAvailableCpuResourceRequest) SetModifyType(v string) *DescribeAvailableCpuResourceRequest {
	s.ModifyType = &v
	return s
}

func (s *DescribeAvailableCpuResourceRequest) SetTenantId(v string) *DescribeAvailableCpuResourceRequest {
	s.TenantId = &v
	return s
}

type DescribeAvailableCpuResourceResponseBody struct {
	// 可用的CPU资源信息。
	Data []*DescribeAvailableCpuResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableCpuResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableCpuResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCpuResourceResponseBody) SetData(v []*DescribeAvailableCpuResourceResponseBodyData) *DescribeAvailableCpuResourceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAvailableCpuResourceResponseBody) SetRequestId(v string) *DescribeAvailableCpuResourceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAvailableCpuResourceResponseBodyData struct {
	// 单个unit最大可申请的CPU核数。
	MaxCpu *int64 `json:"MaxCpu,omitempty" xml:"MaxCpu,omitempty"`
	// 单个unit最小需要申请的CPU核数。
	MinCpu *int64 `json:"MinCpu,omitempty" xml:"MinCpu,omitempty"`
	// 租户的unit个数。
	UnitNum *int64 `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s DescribeAvailableCpuResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableCpuResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCpuResourceResponseBodyData) SetMaxCpu(v int64) *DescribeAvailableCpuResourceResponseBodyData {
	s.MaxCpu = &v
	return s
}

func (s *DescribeAvailableCpuResourceResponseBodyData) SetMinCpu(v int64) *DescribeAvailableCpuResourceResponseBodyData {
	s.MinCpu = &v
	return s
}

func (s *DescribeAvailableCpuResourceResponseBodyData) SetUnitNum(v int64) *DescribeAvailableCpuResourceResponseBodyData {
	s.UnitNum = &v
	return s
}

type DescribeAvailableCpuResourceResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAvailableCpuResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableCpuResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableCpuResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCpuResourceResponse) SetHeaders(v map[string]*string) *DescribeAvailableCpuResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableCpuResourceResponse) SetBody(v *DescribeAvailableCpuResourceResponseBody) *DescribeAvailableCpuResourceResponse {
	s.Body = v
	return s
}

type DescribeAvailableMemResourceRequest struct {
	CpuNum *int64 `json:"CpuNum,omitempty" xml:"CpuNum,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 租户的资源分布节点数。 其与集群的部署模式相耦合，如集群模式为2-2-2，则最后分布节点数最多为2个。
	UnitNum *int64 `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s DescribeAvailableMemResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableMemResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableMemResourceRequest) SetCpuNum(v int64) *DescribeAvailableMemResourceRequest {
	s.CpuNum = &v
	return s
}

func (s *DescribeAvailableMemResourceRequest) SetInstanceId(v string) *DescribeAvailableMemResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAvailableMemResourceRequest) SetTenantId(v string) *DescribeAvailableMemResourceRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeAvailableMemResourceRequest) SetUnitNum(v int64) *DescribeAvailableMemResourceRequest {
	s.UnitNum = &v
	return s
}

type DescribeAvailableMemResourceResponseBody struct {
	// 可用的内存资源信息。
	Data *DescribeAvailableMemResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableMemResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableMemResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableMemResourceResponseBody) SetData(v *DescribeAvailableMemResourceResponseBodyData) *DescribeAvailableMemResourceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAvailableMemResourceResponseBody) SetRequestId(v string) *DescribeAvailableMemResourceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAvailableMemResourceResponseBodyData struct {
	// 单个unit最大可申请的内存大小，单位：GB。
	MaxMem *int64 `json:"MaxMem,omitempty" xml:"MaxMem,omitempty"`
	// 单个unit最小需要申请的内存大小，单位：GB。
	MinMem *int64 `json:"MinMem,omitempty" xml:"MinMem,omitempty"`
	// 租户的unit个数。
	UsedMem *int64 `json:"UsedMem,omitempty" xml:"UsedMem,omitempty"`
}

func (s DescribeAvailableMemResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableMemResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAvailableMemResourceResponseBodyData) SetMaxMem(v int64) *DescribeAvailableMemResourceResponseBodyData {
	s.MaxMem = &v
	return s
}

func (s *DescribeAvailableMemResourceResponseBodyData) SetMinMem(v int64) *DescribeAvailableMemResourceResponseBodyData {
	s.MinMem = &v
	return s
}

func (s *DescribeAvailableMemResourceResponseBodyData) SetUsedMem(v int64) *DescribeAvailableMemResourceResponseBodyData {
	s.UsedMem = &v
	return s
}

type DescribeAvailableMemResourceResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAvailableMemResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableMemResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableMemResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableMemResourceResponse) SetHeaders(v map[string]*string) *DescribeAvailableMemResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableMemResourceResponse) SetBody(v *DescribeAvailableMemResourceResponseBody) *DescribeAvailableMemResourceResponse {
	s.Body = v
	return s
}

type DescribeCharsetRequest struct {
	// 租户模式。 当前支持Oracle模式（Oracle）、MySQL模式（MySQL） 详细参见：DescribeInstanceTenantModes。
	TenantMode *string `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
}

func (s DescribeCharsetRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharsetRequest) GoString() string {
	return s.String()
}

func (s *DescribeCharsetRequest) SetTenantMode(v string) *DescribeCharsetRequest {
	s.TenantMode = &v
	return s
}

type DescribeCharsetResponseBody struct {
	// 字符集列表。
	Charset []*DescribeCharsetResponseBodyCharset `json:"Charset,omitempty" xml:"Charset,omitempty" type:"Repeated"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCharsetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharsetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCharsetResponseBody) SetCharset(v []*DescribeCharsetResponseBodyCharset) *DescribeCharsetResponseBody {
	s.Charset = v
	return s
}

func (s *DescribeCharsetResponseBody) SetRequestId(v string) *DescribeCharsetResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCharsetResponseBodyCharset struct {
	// 字符集名称。
	Charset *string `json:"Charset,omitempty" xml:"Charset,omitempty"`
}

func (s DescribeCharsetResponseBodyCharset) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharsetResponseBodyCharset) GoString() string {
	return s.String()
}

func (s *DescribeCharsetResponseBodyCharset) SetCharset(v string) *DescribeCharsetResponseBodyCharset {
	s.Charset = &v
	return s
}

type DescribeCharsetResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCharsetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCharsetResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharsetResponse) GoString() string {
	return s.String()
}

func (s *DescribeCharsetResponse) SetHeaders(v map[string]*string) *DescribeCharsetResponse {
	s.Headers = v
	return s
}

func (s *DescribeCharsetResponse) SetBody(v *DescribeCharsetResponseBody) *DescribeCharsetResponse {
	s.Body = v
	return s
}

type DescribeDatabasesRequest struct {
	// 数据库名称。 不能使用某些预留关键字，如 test、mysql。
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// 分页查询的页码。 起始值：1 默认值：1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时设置的每页行数。 最大值：100 默认值：10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 查询列表的删选关键字。
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 是否返回数据库中的表信息。 Oracle模式使用true，MySQL模式使用false。
	WithTables *bool `json:"WithTables,omitempty" xml:"WithTables,omitempty"`
}

func (s DescribeDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesRequest) SetDatabaseName(v string) *DescribeDatabasesRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDatabasesRequest) SetPageNumber(v int32) *DescribeDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabasesRequest) SetPageSize(v int32) *DescribeDatabasesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabasesRequest) SetSearchKey(v string) *DescribeDatabasesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeDatabasesRequest) SetTenantId(v string) *DescribeDatabasesRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeDatabasesRequest) SetWithTables(v bool) *DescribeDatabasesRequest {
	s.WithTables = &v
	return s
}

type DescribeDatabasesResponseBody struct {
	// 租户中的数据库列表。
	Databases []*DescribeDatabasesResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 租户中的数据库总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDatabasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBody) SetDatabases(v []*DescribeDatabasesResponseBodyDatabases) *DescribeDatabasesResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDatabasesResponseBody) SetRequestId(v string) *DescribeDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabasesResponseBody) SetTotalCount(v int32) *DescribeDatabasesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDatabasesResponseBodyDatabases struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 实际数据大小
	DataSize *float64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// 数据库名称。
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// 数据库类型
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// 数据库的描述信息。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 数据库的编码。目前支持utf8mb4、gbk等编码。
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// 所需容量
	RequiredSize *float64 `json:"RequiredSize,omitempty" xml:"RequiredSize,omitempty"`
	// 数据库的状态。 - ONLINE: 运行中 - DELETING: 删除中
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 数据库表信息
	Tables []*DescribeDatabasesResponseBodyDatabasesTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 对该数据库赋权的账号信息。
	Users []*DescribeDatabasesResponseBodyDatabasesUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s DescribeDatabasesResponseBodyDatabases) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabases) SetCreateTime(v string) *DescribeDatabasesResponseBodyDatabases {
	s.CreateTime = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetDataSize(v float64) *DescribeDatabasesResponseBodyDatabases {
	s.DataSize = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetDatabaseName(v string) *DescribeDatabasesResponseBodyDatabases {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetDbType(v string) *DescribeDatabasesResponseBodyDatabases {
	s.DbType = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetDescription(v string) *DescribeDatabasesResponseBodyDatabases {
	s.Description = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetEncoding(v string) *DescribeDatabasesResponseBodyDatabases {
	s.Encoding = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetRequiredSize(v float64) *DescribeDatabasesResponseBodyDatabases {
	s.RequiredSize = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetStatus(v string) *DescribeDatabasesResponseBodyDatabases {
	s.Status = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetTables(v []*DescribeDatabasesResponseBodyDatabasesTables) *DescribeDatabasesResponseBodyDatabases {
	s.Tables = v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetTenantId(v string) *DescribeDatabasesResponseBodyDatabases {
	s.TenantId = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetUsers(v []*DescribeDatabasesResponseBodyDatabasesUsers) *DescribeDatabasesResponseBodyDatabases {
	s.Users = v
	return s
}

type DescribeDatabasesResponseBodyDatabasesTables struct {
	// 数据库表名
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeDatabasesResponseBodyDatabasesTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesTables) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesTables) SetTableName(v string) *DescribeDatabasesResponseBodyDatabasesTables {
	s.TableName = &v
	return s
}

type DescribeDatabasesResponseBodyDatabasesUsers struct {
	// 账号赋予该库的角色权限。 对于MySQL模式，角色为数据库（Database）级别，其有以下几类： - ReadWrite：读写权限，包括ALL PRIVILEGES； - ReadOnly：只读权限，包括SELECT - DDL: DDL权限，包括CREATE,DROP,ALTER,SHOW VIEW,CREATE VIEW - DML: DML权限，包括SELECT,INSERT,UPDATE,DELETE,SHOW VIEW。
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 账号名称。
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// 账号类型。 - Admin：超级账户 - Normal：普通账户
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeDatabasesResponseBodyDatabasesUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesUsers) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesUsers) SetRole(v string) *DescribeDatabasesResponseBodyDatabasesUsers {
	s.Role = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesUsers) SetUserName(v string) *DescribeDatabasesResponseBodyDatabasesUsers {
	s.UserName = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesUsers) SetUserType(v string) *DescribeDatabasesResponseBodyDatabasesUsers {
	s.UserType = &v
	return s
}

type DescribeDatabasesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponse) SetHeaders(v map[string]*string) *DescribeDatabasesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabasesResponse) SetBody(v *DescribeDatabasesResponseBody) *DescribeDatabasesResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
	// OceanBase 集群 ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 分页查询的页码。<br>起始值：1， 默认值：1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时设置的每页行数。<br>最大值：100，默认值：10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) SetInstanceId(v string) *DescribeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceRequest) SetPageNumber(v int32) *DescribeInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceRequest) SetPageSize(v int32) *DescribeInstanceRequest {
	s.PageSize = &v
	return s
}

type DescribeInstanceResponseBody struct {
	// OceanBase 集群信息。
	Instance *DescribeInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// 请求 ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetInstance(v *DescribeInstanceResponseBodyInstance) *DescribeInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceResponseBodyInstance struct {
	// 是否开启自动续费。该参数只在预付费（PREPAY）集群有意义。
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// 是否开启自动升级 OBServer 版本。
	AutoUpgradeObVersion *bool `json:"AutoUpgradeObVersion,omitempty" xml:"AutoUpgradeObVersion,omitempty"`
	// 可用区列表。
	AvailableZones []*string `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	// 集群的创建时间（UTC时间）。
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 集群的数据合并时间。
	DataMergeTime *string `json:"DataMergeTime,omitempty" xml:"DataMergeTime,omitempty"`
	// 集群的数据副本模式。 单机房为n，双机房为n-n，多机房为n-n-n，其中n为各机房的observer节点数。
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// 集群的部署类型。<br> - multiple：多机房 <br>- single：单机房 <br>- dual：双机房
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// 集群部署的存储类型。默认为cloud_essd_pl1：ESSD云盘。
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// 集群过期时间（UTC格式）。
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// 集群规格信息。<br> 当前支持四种套餐：<br> - 8C32G：8核 32GB <br>- 14C70G：14核 70GB<br> - 30C180G：30核 180GB<br> - 62C400G：62核 400GB。
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// OceanBase 集群 ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// OceanBase 集群名称。
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// OBServer版本是否为最新版本。
	IsLatestObVersion *bool `json:"IsLatestObVersion,omitempty" xml:"IsLatestObVersion,omitempty"`
	// 是否使用可信ecs
	IsTrustEcs *bool `json:"IsTrustEcs,omitempty" xml:"IsTrustEcs,omitempty"`
	// 集群的每天例行维护时间，UTC时间。
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	// OBServer 详细版本信息。
	ObRpmVersion *string `json:"ObRpmVersion,omitempty" xml:"ObRpmVersion,omitempty"`
	// OceanBase 集群的付费类型 <br>- PREPAY：预付费 <br>- POSTPAY：按量付费
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// 集群的资源信息
	Resource *DescribeInstanceResponseBodyInstanceResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// OceanBase 集群的系列 <br>- NORMAL：高可用版本 <br>- BASIC：基础版本
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// 集群状态。 <br>- PENDING_CREATE: 创建中 <br>- ONLINE: 运行中 <br>- TENANT_CREATING：租户创建中 <br>- TENANT_SPEC_MODIFYING：租户规格修改中 <br>- EXPANDING: 节点扩容中 <br>- REDUCING: 节点缩容中 <br>- SPEC_UPGRADING:套餐规格扩容中 <br>- DISK_UPGRADING:存储规格扩容中 <br>- WHITE_LIST_MODIFYING: 修改白名单中 <br>- PARAMETER_MODIFYING: 修改参数中 <br>- SSL_MODIFYING: SSL变更中 <br>- PREPAID_EXPIRE_CLOSED: 预付费集群欠费中 <br>- ARREARS_CLOSED: 后付费集群欠费中 <br>- PENDING_DELETE: 删除中。 集群一般为运行中的状态（ONLINE）。
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// OBServer 版本信息。
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeInstanceResponseBodyInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstance) SetAutoRenewal(v bool) *DescribeInstanceResponseBodyInstance {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetAutoUpgradeObVersion(v bool) *DescribeInstanceResponseBodyInstance {
	s.AutoUpgradeObVersion = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetAvailableZones(v []*string) *DescribeInstanceResponseBodyInstance {
	s.AvailableZones = v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetCreateTime(v string) *DescribeInstanceResponseBodyInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetDataMergeTime(v string) *DescribeInstanceResponseBodyInstance {
	s.DataMergeTime = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetDeployMode(v string) *DescribeInstanceResponseBodyInstance {
	s.DeployMode = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetDeployType(v string) *DescribeInstanceResponseBodyInstance {
	s.DeployType = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetDiskType(v string) *DescribeInstanceResponseBodyInstance {
	s.DiskType = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetExpireTime(v string) *DescribeInstanceResponseBodyInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetInstanceClass(v string) *DescribeInstanceResponseBodyInstance {
	s.InstanceClass = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetInstanceId(v string) *DescribeInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetInstanceName(v string) *DescribeInstanceResponseBodyInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetIsLatestObVersion(v bool) *DescribeInstanceResponseBodyInstance {
	s.IsLatestObVersion = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetIsTrustEcs(v bool) *DescribeInstanceResponseBodyInstance {
	s.IsTrustEcs = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetMaintainTime(v string) *DescribeInstanceResponseBodyInstance {
	s.MaintainTime = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetObRpmVersion(v string) *DescribeInstanceResponseBodyInstance {
	s.ObRpmVersion = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetPayType(v string) *DescribeInstanceResponseBodyInstance {
	s.PayType = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetResource(v *DescribeInstanceResponseBodyInstanceResource) *DescribeInstanceResponseBodyInstance {
	s.Resource = v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetSeries(v string) *DescribeInstanceResponseBodyInstance {
	s.Series = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetStatus(v string) *DescribeInstanceResponseBodyInstance {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetVersion(v string) *DescribeInstanceResponseBodyInstance {
	s.Version = &v
	return s
}

type DescribeInstanceResponseBodyInstanceResource struct {
	// 集群的CPU资源信息
	Cpu *DescribeInstanceResponseBodyInstanceResourceCpu `json:"Cpu,omitempty" xml:"Cpu,omitempty" type:"Struct"`
	// 集群的存储资源信息
	DiskSize *DescribeInstanceResponseBodyInstanceResourceDiskSize `json:"DiskSize,omitempty" xml:"DiskSize,omitempty" type:"Struct"`
	// 集群的内存资源信息
	Memory *DescribeInstanceResponseBodyInstanceResourceMemory `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	// 集群的资源Unit数量。
	UnitCount *int64 `json:"UnitCount,omitempty" xml:"UnitCount,omitempty"`
}

func (s DescribeInstanceResponseBodyInstanceResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceResource) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceResource) SetCpu(v *DescribeInstanceResponseBodyInstanceResourceCpu) *DescribeInstanceResponseBodyInstanceResource {
	s.Cpu = v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResource) SetDiskSize(v *DescribeInstanceResponseBodyInstanceResourceDiskSize) *DescribeInstanceResponseBodyInstanceResource {
	s.DiskSize = v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResource) SetMemory(v *DescribeInstanceResponseBodyInstanceResourceMemory) *DescribeInstanceResponseBodyInstanceResource {
	s.Memory = v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResource) SetUnitCount(v int64) *DescribeInstanceResponseBodyInstanceResource {
	s.UnitCount = &v
	return s
}

type DescribeInstanceResponseBodyInstanceResourceCpu struct {
	// 集群总CPU，单位：核数
	TotalCpu *int64 `json:"TotalCpu,omitempty" xml:"TotalCpu,omitempty"`
	// 集群中每个副本节点的CPU，单位：核数
	UnitCpu *int64 `json:"UnitCpu,omitempty" xml:"UnitCpu,omitempty"`
	// 集群已使用的CPU，单位：核数
	UsedCpu *int64 `json:"UsedCpu,omitempty" xml:"UsedCpu,omitempty"`
}

func (s DescribeInstanceResponseBodyInstanceResourceCpu) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceResourceCpu) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceResourceCpu) SetTotalCpu(v int64) *DescribeInstanceResponseBodyInstanceResourceCpu {
	s.TotalCpu = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceCpu) SetUnitCpu(v int64) *DescribeInstanceResponseBodyInstanceResourceCpu {
	s.UnitCpu = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceCpu) SetUsedCpu(v int64) *DescribeInstanceResponseBodyInstanceResourceCpu {
	s.UsedCpu = &v
	return s
}

type DescribeInstanceResponseBodyInstanceResourceDiskSize struct {
	// 集群总存储空间，单位：GB
	TotalDiskSize *int64 `json:"TotalDiskSize,omitempty" xml:"TotalDiskSize,omitempty"`
	// 集群每个副本的存储空间，单位：GB
	UnitDiskSize *int64 `json:"UnitDiskSize,omitempty" xml:"UnitDiskSize,omitempty"`
	// 集群已使用的存储空间，单位：GB
	UsedDiskSize *int64 `json:"UsedDiskSize,omitempty" xml:"UsedDiskSize,omitempty"`
}

func (s DescribeInstanceResponseBodyInstanceResourceDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceResourceDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceResourceDiskSize) SetTotalDiskSize(v int64) *DescribeInstanceResponseBodyInstanceResourceDiskSize {
	s.TotalDiskSize = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceDiskSize) SetUnitDiskSize(v int64) *DescribeInstanceResponseBodyInstanceResourceDiskSize {
	s.UnitDiskSize = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceDiskSize) SetUsedDiskSize(v int64) *DescribeInstanceResponseBodyInstanceResourceDiskSize {
	s.UsedDiskSize = &v
	return s
}

type DescribeInstanceResponseBodyInstanceResourceMemory struct {
	// 集群总内存，单位：GB
	TotalMemory *int64 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	// 集群中每个副本的内存，单位：GB
	UnitMemory *int64 `json:"UnitMemory,omitempty" xml:"UnitMemory,omitempty"`
	// 集群已使用的内存，单位：GB
	UsedMemory *int64 `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
}

func (s DescribeInstanceResponseBodyInstanceResourceMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceResourceMemory) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceResourceMemory) SetTotalMemory(v int64) *DescribeInstanceResponseBodyInstanceResourceMemory {
	s.TotalMemory = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceMemory) SetUnitMemory(v int64) *DescribeInstanceResponseBodyInstanceResourceMemory {
	s.UnitMemory = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceMemory) SetUsedMemory(v int64) *DescribeInstanceResponseBodyInstanceResourceMemory {
	s.UsedMemory = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeInstanceCreatableZoneRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceCreatableZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceCreatableZoneRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCreatableZoneRequest) SetInstanceId(v string) *DescribeInstanceCreatableZoneRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceCreatableZoneResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 可用区列表信息。
	ZoneList []*DescribeInstanceCreatableZoneResponseBodyZoneList `json:"ZoneList,omitempty" xml:"ZoneList,omitempty" type:"Repeated"`
}

func (s DescribeInstanceCreatableZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceCreatableZoneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCreatableZoneResponseBody) SetRequestId(v string) *DescribeInstanceCreatableZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceCreatableZoneResponseBody) SetZoneList(v []*DescribeInstanceCreatableZoneResponseBodyZoneList) *DescribeInstanceCreatableZoneResponseBody {
	s.ZoneList = v
	return s
}

type DescribeInstanceCreatableZoneResponseBodyZoneList struct {
	// 是否是集群部署的可用区。
	IsInCluster *bool `json:"IsInCluster,omitempty" xml:"IsInCluster,omitempty"`
	// 可用区ID。
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s DescribeInstanceCreatableZoneResponseBodyZoneList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceCreatableZoneResponseBodyZoneList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCreatableZoneResponseBodyZoneList) SetIsInCluster(v bool) *DescribeInstanceCreatableZoneResponseBodyZoneList {
	s.IsInCluster = &v
	return s
}

func (s *DescribeInstanceCreatableZoneResponseBodyZoneList) SetZone(v string) *DescribeInstanceCreatableZoneResponseBodyZoneList {
	s.Zone = &v
	return s
}

type DescribeInstanceCreatableZoneResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceCreatableZoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceCreatableZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceCreatableZoneResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCreatableZoneResponse) SetHeaders(v map[string]*string) *DescribeInstanceCreatableZoneResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceCreatableZoneResponse) SetBody(v *DescribeInstanceCreatableZoneResponseBody) *DescribeInstanceCreatableZoneResponse {
	s.Body = v
	return s
}

type DescribeInstanceTenantModesRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceTenantModesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTenantModesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTenantModesRequest) SetInstanceId(v string) *DescribeInstanceTenantModesRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceTenantModesResponseBody struct {
	// 租户模式列表。
	InstanceModes []*string `json:"InstanceModes,omitempty" xml:"InstanceModes,omitempty" type:"Repeated"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceTenantModesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTenantModesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTenantModesResponseBody) SetInstanceModes(v []*string) *DescribeInstanceTenantModesResponseBody {
	s.InstanceModes = v
	return s
}

func (s *DescribeInstanceTenantModesResponseBody) SetRequestId(v string) *DescribeInstanceTenantModesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceTenantModesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceTenantModesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceTenantModesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTenantModesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTenantModesResponse) SetHeaders(v map[string]*string) *DescribeInstanceTenantModesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTenantModesResponse) SetBody(v *DescribeInstanceTenantModesResponseBody) *DescribeInstanceTenantModesResponse {
	s.Body = v
	return s
}

type DescribeInstanceTopologyRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceTopologyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyRequest) SetInstanceId(v string) *DescribeInstanceTopologyRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceTopologyResponseBody struct {
	// 集群的拓扑信息。
	InstanceTopology *DescribeInstanceTopologyResponseBodyInstanceTopology `json:"InstanceTopology,omitempty" xml:"InstanceTopology,omitempty" type:"Struct"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceTopologyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBody) SetInstanceTopology(v *DescribeInstanceTopologyResponseBodyInstanceTopology) *DescribeInstanceTopologyResponseBody {
	s.InstanceTopology = v
	return s
}

func (s *DescribeInstanceTopologyResponseBody) SetRequestId(v string) *DescribeInstanceTopologyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopology struct {
	// 租户信息。
	Tenants []*DescribeInstanceTopologyResponseBodyInstanceTopologyTenants `json:"Tenants,omitempty" xml:"Tenants,omitempty" type:"Repeated"`
	// 集群的可用区信息。
	Zones []*DescribeInstanceTopologyResponseBodyInstanceTopologyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopology) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopology) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopology) SetTenants(v []*DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) *DescribeInstanceTopologyResponseBodyInstanceTopology {
	s.Tenants = v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopology) SetZones(v []*DescribeInstanceTopologyResponseBodyInstanceTopologyZones) *DescribeInstanceTopologyResponseBodyInstanceTopology {
	s.Zones = v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyTenants struct {
	// 租户CPU，单位：核。
	TenantCpu *float32 `json:"TenantCpu,omitempty" xml:"TenantCpu,omitempty"`
	// 租户的部署类型。 - multiple：多机房 - single：单机房 - dual：双机房。
	TenantDeployType *string `json:"TenantDeployType,omitempty" xml:"TenantDeployType,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 租户内存大小，单位：GB。
	TenantMemory *float32 `json:"TenantMemory,omitempty" xml:"TenantMemory,omitempty"`
	// 租户模式。 - Oracle：Oracle模式 - MySQL：MySQL模式。
	TenantMode *string `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
	// 租户名称。
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// 租户状态。 - PENDING_CREATE: 创建中 - RESTORE: 恢复中 - ONLINE: 运行中 - SPEC_MODIFYING: 规格修改中 - ALLOCATING_INTERNET_ADDRESS: 公网地址分配中 - PENDING_OFFLINE_INTERNET_ADDRESS: 公网地址关闭中 - PRIMARY_ZONE_MODIFYING: 切主可用区中 - PARAMETER_MODIFYING: 参数修改中 - WHITE_LIST_MODIFING: 白名单修改中
	TenantStatus *string `json:"TenantStatus,omitempty" xml:"TenantStatus,omitempty"`
	// 租户的unit个数。
	TenantUnitNum *int32 `json:"TenantUnitNum,omitempty" xml:"TenantUnitNum,omitempty"`
	// 租户的可用区信息。
	TenantZones []*DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones `json:"TenantZones,omitempty" xml:"TenantZones,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantCpu(v float32) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantCpu = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantDeployType(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantDeployType = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantId(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantMemory(v float32) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantMemory = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantMode(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantMode = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantName(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantName = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantStatus(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantStatus = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantUnitNum(v int32) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantUnitNum = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantZones(v []*DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantZones = v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones struct {
	// 是否为主可用区
	IsPrimaryTenantZone *string `json:"IsPrimaryTenantZone,omitempty" xml:"IsPrimaryTenantZone,omitempty"`
	// 可用区ID。
	TenantZoneId *string `json:"TenantZoneId,omitempty" xml:"TenantZoneId,omitempty"`
	// 可用区的访问角色。 - ReadWrite：可读可写 - ReadOnly：只读。
	TenantZoneRole *string `json:"TenantZoneRole,omitempty" xml:"TenantZoneRole,omitempty"`
	// 资源节点信息列表
	Units []*DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits `json:"Units,omitempty" xml:"Units,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones) SetIsPrimaryTenantZone(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones {
	s.IsPrimaryTenantZone = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones) SetTenantZoneId(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones {
	s.TenantZoneId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones) SetTenantZoneRole(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones {
	s.TenantZoneRole = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones) SetUnits(v []*DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones {
	s.Units = v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits struct {
	// 是否可以取消迁移。该字段只对正在手动迁入中或手动迁出中的unit有效。
	EnableCancelMigrateUnit *bool `json:"EnableCancelMigrateUnit,omitempty" xml:"EnableCancelMigrateUnit,omitempty"`
	// 是否可做迁移。
	EnableMigrateUnit *bool `json:"EnableMigrateUnit,omitempty" xml:"EnableMigrateUnit,omitempty"`
	// 是否为手动迁移。
	ManualMigrate *bool `json:"ManualMigrate,omitempty" xml:"ManualMigrate,omitempty"`
	// 资源节点所在的observer节点ID
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// 资源节点CPU，单位：核。
	UnitCpu *float32 `json:"UnitCpu,omitempty" xml:"UnitCpu,omitempty"`
	// 资源节点ID
	UnitId *string `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
	// 资源节点内存大小，单位：GB。
	UnitMemory *float32 `json:"UnitMemory,omitempty" xml:"UnitMemory,omitempty"`
	// 资源节点的状态。 ONLINE: 运行中 IMMIGRATING 迁入中 EMIGRATING: 迁出中 CANCEL_EMIGRATING: 取消迁入中 CANCEL_EMIGRATING：取消迁出中 DELETING：删除中
	UnitStatus *string `json:"UnitStatus,omitempty" xml:"UnitStatus,omitempty"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetEnableCancelMigrateUnit(v bool) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.EnableCancelMigrateUnit = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetEnableMigrateUnit(v bool) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.EnableMigrateUnit = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetManualMigrate(v bool) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.ManualMigrate = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetNodeId(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.NodeId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetUnitCpu(v float32) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.UnitCpu = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetUnitId(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.UnitId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetUnitMemory(v float32) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.UnitMemory = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetUnitStatus(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.UnitStatus = &v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyZones struct {
	// 节点信息。
	Nodes []*DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// 地域ID。
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 可用区的存储容量。
	ZoneDisk *string `json:"ZoneDisk,omitempty" xml:"ZoneDisk,omitempty"`
	// 可用区ID。
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZones) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZones) SetNodes(v []*DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes) *DescribeInstanceTopologyResponseBodyInstanceTopologyZones {
	s.Nodes = v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZones) SetRegion(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyZones {
	s.Region = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZones) SetZoneDisk(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyZones {
	s.ZoneDisk = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZones) SetZoneId(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyZones {
	s.ZoneId = &v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes struct {
	// 节点归属的副本ID。
	NodeCopyId *int64 `json:"NodeCopyId,omitempty" xml:"NodeCopyId,omitempty"`
	// 节点ID。
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// 节点资源信息。
	NodeResource []*DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource `json:"NodeResource,omitempty" xml:"NodeResource,omitempty" type:"Repeated"`
	// 节点运行状态。
	NodeStatus *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes) SetNodeCopyId(v int64) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes {
	s.NodeCopyId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes) SetNodeId(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes {
	s.NodeId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes) SetNodeResource(v []*DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes {
	s.NodeResource = v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes) SetNodeStatus(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes {
	s.NodeStatus = &v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource struct {
	// 节点CPU资源信息。
	Cpu *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu `json:"Cpu,omitempty" xml:"Cpu,omitempty" type:"Struct"`
	// 节点存储资源信息。
	DiskSize *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize `json:"DiskSize,omitempty" xml:"DiskSize,omitempty" type:"Struct"`
	// 节点内存资源信息。
	Memory *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource) SetCpu(v *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource {
	s.Cpu = v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource) SetDiskSize(v *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource {
	s.DiskSize = v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource) SetMemory(v *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource {
	s.Memory = v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu struct {
	// 节点总的CPU，单位：核。
	TotalCpu *int32 `json:"TotalCpu,omitempty" xml:"TotalCpu,omitempty"`
	// 节点已使用的CPU，单位：核。
	UsedCpu *float32 `json:"UsedCpu,omitempty" xml:"UsedCpu,omitempty"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu) SetTotalCpu(v int32) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu {
	s.TotalCpu = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu) SetUsedCpu(v float32) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu {
	s.UsedCpu = &v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize struct {
	// 节点总存储空间，单位：GB。
	TotalDiskSize *int64 `json:"TotalDiskSize,omitempty" xml:"TotalDiskSize,omitempty"`
	// 节点已使用的存储空间，单位：GB。
	UsedDiskSize *int64 `json:"UsedDiskSize,omitempty" xml:"UsedDiskSize,omitempty"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize) SetTotalDiskSize(v int64) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize {
	s.TotalDiskSize = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize) SetUsedDiskSize(v int64) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize {
	s.UsedDiskSize = &v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory struct {
	// 节点的总内存，单位：GB。
	TotalMemory *int64 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	// 节点已使用的内存，单位：GB。
	UsedMemory *float32 `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory) SetTotalMemory(v int64) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory {
	s.TotalMemory = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory) SetUsedMemory(v float32) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory {
	s.UsedMemory = &v
	return s
}

type DescribeInstanceTopologyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceTopologyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponse) SetHeaders(v map[string]*string) *DescribeInstanceTopologyResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTopologyResponse) SetBody(v *DescribeInstanceTopologyResponseBody) *DescribeInstanceTopologyResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Oceanbase集群名称。 长度为1~20个英文或中文字符。如果没有指定该参数，默认值为集群的InstanceId。
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 分页查询的页码。 起始值：1 默认值：1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时设置的每页行数。 最大值：100 默认值：10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 资源组ID信息。如果不填写，返回全部资源。
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 查询列表的删选关键字。
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetInstanceId(v string) *DescribeInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceName(v string) *DescribeInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int32) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetSearchKey(v string) *DescribeInstancesRequest {
	s.SearchKey = &v
	return s
}

type DescribeInstancesResponseBody struct {
	// Oceanbase集群列表。
	Instances []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 查询到的Oceanbase集群个数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
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

type DescribeInstancesResponseBodyInstances struct {
	// 集群部署所在的可用区信息。
	AvailableZones []*string `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	// Oceanbase集群的产品码。 - oceanbase_oceanbasepre_public_cn：国内站云数据库包年包月套餐。 - oceanbase_oceanbasepost_public_cn：国内站云数据库按小时付费套餐。 - oceanbase_obpre_public_intl：国际站云数据库包年包月套餐。
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// 集群的CPU核数。
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// 集群的创建时间，UTC时间。
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 集群的数据副本模式。 单机房为n，双机房为n-n，多机房为n-n-n，其中n为各机房的observer节点数。
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// 集群的部署类型。 - multiple：多机房 - single：单机房 - dual：双机房
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// 存储空间大小，单位GB。
	DiskSize *string `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// 集群部署的存储类型。默认为cloud_essd_pl1：ESSD云盘。
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// 是否允许新增节点。
	EnableUpgradeNodes *bool `json:"EnableUpgradeNodes,omitempty" xml:"EnableUpgradeNodes,omitempty"`
	// 集群已过期时间，单位：秒(s)。
	ExpireSeconds *int32 `json:"ExpireSeconds,omitempty" xml:"ExpireSeconds,omitempty"`
	// 集群过期时间（UTC格式）。
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// 集群规格信息。 当前支持四种套餐： - 8C32G：8核 32GB - 14C70G：14核 70GB - 30C180G：30核 180GB - 62C400G：62核 400GB。
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Oceanbase集群名称。
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 集群的每天例行维护时间，UTC时间。
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	// 实例的内存大小，单位GB。
	Mem *int64 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// Oceanbase集群的付费类型 - PREPAY：预付费 - POSTPAY：按量付费
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// 集群的资源信息
	Resource *DescribeInstancesResponseBodyInstancesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// 资源组ID信息。
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 集群的白名单信息。（待废弃）
	SecurityIps []*string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty" type:"Repeated"`
	// Oceanbase集群的系列 - NORMAL：高可用版本 - BASIC：基础版本
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// 集群状态。 - PENDING_CREATE: 创建中 - ONLINE: 运行中 - TENANT_CREATING：租户创建中 - TENANT_SPEC_MODIFYING：租户规格修改中 - EXPANDING: 节点扩容中 - REDUCING: 节点缩容中 - SPEC_UPGRADING:套餐规格扩容中 - DISK_UPGRADING:存储规格扩容中 - WHITE_LIST_MODIFYING: 修改白名单中 - PARAMETER_MODIFYING: 修改参数中 - SSL_MODIFYING: SSL变更中 - PREPAID_EXPIRE_CLOSED: 预付费集群欠费中 - ARREARS_CLOSED: 后付费集群欠费中 - PENDING_DELETE: 删除中。 集群一般为运行中的状态（ONLINE）。
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// 已使用的存储空间，单位GB。
	UsedDiskSize *int64 `json:"UsedDiskSize,omitempty" xml:"UsedDiskSize,omitempty"`
	// Observer版本信息。
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// vpcId
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetAvailableZones(v []*string) *DescribeInstancesResponseBodyInstances {
	s.AvailableZones = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetCommodityCode(v string) *DescribeInstancesResponseBodyInstances {
	s.CommodityCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetCpu(v int32) *DescribeInstancesResponseBodyInstances {
	s.Cpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetCreateTime(v string) *DescribeInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDeployMode(v string) *DescribeInstancesResponseBodyInstances {
	s.DeployMode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDeployType(v string) *DescribeInstancesResponseBodyInstances {
	s.DeployType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDiskSize(v string) *DescribeInstancesResponseBodyInstances {
	s.DiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDiskType(v string) *DescribeInstancesResponseBodyInstances {
	s.DiskType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetEnableUpgradeNodes(v bool) *DescribeInstancesResponseBodyInstances {
	s.EnableUpgradeNodes = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireSeconds(v int32) *DescribeInstancesResponseBodyInstances {
	s.ExpireSeconds = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v string) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceClass(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceClass = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceName(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetMaintainTime(v string) *DescribeInstancesResponseBodyInstances {
	s.MaintainTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetMem(v int64) *DescribeInstancesResponseBodyInstances {
	s.Mem = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetPayType(v string) *DescribeInstancesResponseBodyInstances {
	s.PayType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResource(v *DescribeInstancesResponseBodyInstancesResource) *DescribeInstancesResponseBodyInstances {
	s.Resource = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceGroupId(v string) *DescribeInstancesResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetSecurityIps(v []*string) *DescribeInstancesResponseBodyInstances {
	s.SecurityIps = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetSeries(v string) *DescribeInstancesResponseBodyInstances {
	s.Series = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetState(v string) *DescribeInstancesResponseBodyInstances {
	s.State = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetUsedDiskSize(v int64) *DescribeInstancesResponseBodyInstances {
	s.UsedDiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVersion(v string) *DescribeInstancesResponseBodyInstances {
	s.Version = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVpcId(v string) *DescribeInstancesResponseBodyInstances {
	s.VpcId = &v
	return s
}

type DescribeInstancesResponseBodyInstancesResource struct {
	// 集群的CPU资源信息
	Cpu *DescribeInstancesResponseBodyInstancesResourceCpu `json:"Cpu,omitempty" xml:"Cpu,omitempty" type:"Struct"`
	// 集群的存储资源信息
	DiskSize *DescribeInstancesResponseBodyInstancesResourceDiskSize `json:"DiskSize,omitempty" xml:"DiskSize,omitempty" type:"Struct"`
	// 集群的内存资源信息
	Memory *DescribeInstancesResponseBodyInstancesResourceMemory `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	// 集群的资源Unit数量。
	UnitCount *int64 `json:"UnitCount,omitempty" xml:"UnitCount,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesResource) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesResource) SetCpu(v *DescribeInstancesResponseBodyInstancesResourceCpu) *DescribeInstancesResponseBodyInstancesResource {
	s.Cpu = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResource) SetDiskSize(v *DescribeInstancesResponseBodyInstancesResourceDiskSize) *DescribeInstancesResponseBodyInstancesResource {
	s.DiskSize = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResource) SetMemory(v *DescribeInstancesResponseBodyInstancesResourceMemory) *DescribeInstancesResponseBodyInstancesResource {
	s.Memory = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResource) SetUnitCount(v int64) *DescribeInstancesResponseBodyInstancesResource {
	s.UnitCount = &v
	return s
}

type DescribeInstancesResponseBodyInstancesResourceCpu struct {
	// 集群总CPU，单位：核数
	TotalCpu *int64 `json:"TotalCpu,omitempty" xml:"TotalCpu,omitempty"`
	// 集群中每个副本节点的CPU，单位：核数
	UnitCpu *int64 `json:"UnitCpu,omitempty" xml:"UnitCpu,omitempty"`
	// 集群已使用的CPU，单位：核数
	UsedCpu *int64 `json:"UsedCpu,omitempty" xml:"UsedCpu,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesResourceCpu) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesResourceCpu) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesResourceCpu) SetTotalCpu(v int64) *DescribeInstancesResponseBodyInstancesResourceCpu {
	s.TotalCpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceCpu) SetUnitCpu(v int64) *DescribeInstancesResponseBodyInstancesResourceCpu {
	s.UnitCpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceCpu) SetUsedCpu(v int64) *DescribeInstancesResponseBodyInstancesResourceCpu {
	s.UsedCpu = &v
	return s
}

type DescribeInstancesResponseBodyInstancesResourceDiskSize struct {
	// 集群总存储空间，单位：GB
	TotalDiskSize *int64 `json:"TotalDiskSize,omitempty" xml:"TotalDiskSize,omitempty"`
	// 集群每个副本的存储空间，单位：GB
	UnitDiskSize *int64 `json:"UnitDiskSize,omitempty" xml:"UnitDiskSize,omitempty"`
	// 集群已使用的存储空间，单位：GB
	UsedDiskSize *int64 `json:"UsedDiskSize,omitempty" xml:"UsedDiskSize,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesResourceDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesResourceDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesResourceDiskSize) SetTotalDiskSize(v int64) *DescribeInstancesResponseBodyInstancesResourceDiskSize {
	s.TotalDiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceDiskSize) SetUnitDiskSize(v int64) *DescribeInstancesResponseBodyInstancesResourceDiskSize {
	s.UnitDiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceDiskSize) SetUsedDiskSize(v int64) *DescribeInstancesResponseBodyInstancesResourceDiskSize {
	s.UsedDiskSize = &v
	return s
}

type DescribeInstancesResponseBodyInstancesResourceMemory struct {
	// 集群总内存，单位：GB
	TotalMemory *int64 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	// 集群中每个副本的内存，单位：GB
	UnitMemory *int64 `json:"UnitMemory,omitempty" xml:"UnitMemory,omitempty"`
	// 集群已使用的内存，单位：GB
	UsedMemory *int64 `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesResourceMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesResourceMemory) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesResourceMemory) SetTotalMemory(v int64) *DescribeInstancesResponseBodyInstancesResourceMemory {
	s.TotalMemory = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceMemory) SetUnitMemory(v int64) *DescribeInstancesResponseBodyInstancesResourceMemory {
	s.UnitMemory = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceMemory) SetUsedMemory(v int64) *DescribeInstancesResponseBodyInstancesResourceMemory {
	s.UsedMemory = &v
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

type DescribeNodeMetricsRequest struct {
	// 参数历史查看的结束时间。
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 监控指标
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// 节点列表
	NodeIdList *string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty"`
	// 节点名称
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// 分页查询的页码。 起始值：1 默认值：1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时设置的每页行数。 最大值：100 默认值：10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 参数历史查看的起始时间。
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeNodeMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeMetricsRequest) SetEndTime(v string) *DescribeNodeMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetInstanceId(v string) *DescribeNodeMetricsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetMetrics(v string) *DescribeNodeMetricsRequest {
	s.Metrics = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetNodeIdList(v string) *DescribeNodeMetricsRequest {
	s.NodeIdList = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetNodeName(v string) *DescribeNodeMetricsRequest {
	s.NodeName = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetPageNumber(v int32) *DescribeNodeMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetPageSize(v int32) *DescribeNodeMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetStartTime(v string) *DescribeNodeMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetTenantId(v string) *DescribeNodeMetricsRequest {
	s.TenantId = &v
	return s
}

type DescribeNodeMetricsResponseBody struct {
	// 节点指标信息
	NodeMetrics *string `json:"NodeMetrics,omitempty" xml:"NodeMetrics,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总数量
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNodeMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeMetricsResponseBody) SetNodeMetrics(v string) *DescribeNodeMetricsResponseBody {
	s.NodeMetrics = &v
	return s
}

func (s *DescribeNodeMetricsResponseBody) SetRequestId(v string) *DescribeNodeMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeMetricsResponseBody) SetTotalCount(v int32) *DescribeNodeMetricsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeNodeMetricsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNodeMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNodeMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeMetricsResponse) SetHeaders(v map[string]*string) *DescribeNodeMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeMetricsResponse) SetBody(v *DescribeNodeMetricsResponseBody) *DescribeNodeMetricsResponse {
	s.Body = v
	return s
}

type DescribeOutlineBindingRequest struct {
	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 当值为True时，查询SQLID在数据库中的限流信息
	IsConcurrentLimit *bool `json:"IsConcurrentLimit,omitempty" xml:"IsConcurrentLimit,omitempty"`
	// SQLID
	SQLId *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// 租户名称。 长度为2~20 个字符，支持英文字母、数字和下划线，区分大小写，必须以字母或下划线开头。 不可设置为 sys。
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeOutlineBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutlineBindingRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutlineBindingRequest) SetDatabaseName(v string) *DescribeOutlineBindingRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeOutlineBindingRequest) SetInstanceId(v string) *DescribeOutlineBindingRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeOutlineBindingRequest) SetIsConcurrentLimit(v bool) *DescribeOutlineBindingRequest {
	s.IsConcurrentLimit = &v
	return s
}

func (s *DescribeOutlineBindingRequest) SetSQLId(v string) *DescribeOutlineBindingRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeOutlineBindingRequest) SetTableName(v string) *DescribeOutlineBindingRequest {
	s.TableName = &v
	return s
}

func (s *DescribeOutlineBindingRequest) SetTenantId(v string) *DescribeOutlineBindingRequest {
	s.TenantId = &v
	return s
}

type DescribeOutlineBindingResponseBody struct {
	// 绑定信息
	OutlineBinding *DescribeOutlineBindingResponseBodyOutlineBinding `json:"OutlineBinding,omitempty" xml:"OutlineBinding,omitempty" type:"Struct"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOutlineBindingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutlineBindingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutlineBindingResponseBody) SetOutlineBinding(v *DescribeOutlineBindingResponseBodyOutlineBinding) *DescribeOutlineBindingResponseBody {
	s.OutlineBinding = v
	return s
}

func (s *DescribeOutlineBindingResponseBody) SetRequestId(v string) *DescribeOutlineBindingResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOutlineBindingResponseBodyOutlineBinding struct {
	// 绑定索引
	BindIndex *string `json:"BindIndex,omitempty" xml:"BindIndex,omitempty"`
	// 绑定计划
	BindPlan *string `json:"BindPlan,omitempty" xml:"BindPlan,omitempty"`
	// 最大并发
	MaxConcurrent *int32 `json:"MaxConcurrent,omitempty" xml:"MaxConcurrent,omitempty"`
	// OutlineID
	OutlineId *int64 `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
}

func (s DescribeOutlineBindingResponseBodyOutlineBinding) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutlineBindingResponseBodyOutlineBinding) GoString() string {
	return s.String()
}

func (s *DescribeOutlineBindingResponseBodyOutlineBinding) SetBindIndex(v string) *DescribeOutlineBindingResponseBodyOutlineBinding {
	s.BindIndex = &v
	return s
}

func (s *DescribeOutlineBindingResponseBodyOutlineBinding) SetBindPlan(v string) *DescribeOutlineBindingResponseBodyOutlineBinding {
	s.BindPlan = &v
	return s
}

func (s *DescribeOutlineBindingResponseBodyOutlineBinding) SetMaxConcurrent(v int32) *DescribeOutlineBindingResponseBodyOutlineBinding {
	s.MaxConcurrent = &v
	return s
}

func (s *DescribeOutlineBindingResponseBodyOutlineBinding) SetOutlineId(v int64) *DescribeOutlineBindingResponseBodyOutlineBinding {
	s.OutlineId = &v
	return s
}

type DescribeOutlineBindingResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOutlineBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOutlineBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutlineBindingResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutlineBindingResponse) SetHeaders(v map[string]*string) *DescribeOutlineBindingResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutlineBindingResponse) SetBody(v *DescribeOutlineBindingResponseBody) *DescribeOutlineBindingResponse {
	s.Body = v
	return s
}

type DescribeParametersRequest struct {
	// 参数类型。 当前支持集群（CLUSTER)和租户（TENANT）
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// 参数类型的资源标识。 如果为集群参数可以不填，若为租户的参数，则传入租户的TenantId。
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeParametersRequest) SetDimension(v string) *DescribeParametersRequest {
	s.Dimension = &v
	return s
}

func (s *DescribeParametersRequest) SetDimensionValue(v string) *DescribeParametersRequest {
	s.DimensionValue = &v
	return s
}

func (s *DescribeParametersRequest) SetInstanceId(v string) *DescribeParametersRequest {
	s.InstanceId = &v
	return s
}

type DescribeParametersResponseBody struct {
	// 参数列表信息
	Parameters []*DescribeParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBody) SetParameters(v []*DescribeParametersResponseBodyParameters) *DescribeParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeParametersResponseBody) SetRequestId(v string) *DescribeParametersResponseBody {
	s.RequestId = &v
	return s
}

type DescribeParametersResponseBodyParameters struct {
	// 参数的可接受取值范围。 其格式为具备两个字符串元素的数组类型，表示一个范围值，第一个元素为最小值，第二个元素为最大值。
	AcceptableValue []*string `json:"AcceptableValue,omitempty" xml:"AcceptableValue,omitempty" type:"Repeated"`
	// 参数的当前取值。
	CurrentValue *string `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
	// 参数的默认取值。
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// 参数的描述信息。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 参数名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 修改此参数是否需要重启 - true：需要重启 - false：不需要重启
	NeedReboot *bool `json:"NeedReboot,omitempty" xml:"NeedReboot,omitempty"`
	// 参数的不允许取值范围。 其格式为具备两个字符串元素的数组类型，表示一个范围值，第一个元素为最小值，第二个元素为最大值。
	RejectedValue []*string `json:"RejectedValue,omitempty" xml:"RejectedValue,omitempty" type:"Repeated"`
	// 参数取值的类型。 其支持： - ENUM: 数值枚举 - RANGE: 数值范围 - TIME: 时间 - CAPACITY：存储容量值（K，M，G）
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s DescribeParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyParameters) SetAcceptableValue(v []*string) *DescribeParametersResponseBodyParameters {
	s.AcceptableValue = v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetCurrentValue(v string) *DescribeParametersResponseBodyParameters {
	s.CurrentValue = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetDefaultValue(v string) *DescribeParametersResponseBodyParameters {
	s.DefaultValue = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetDescription(v string) *DescribeParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetName(v string) *DescribeParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetNeedReboot(v bool) *DescribeParametersResponseBodyParameters {
	s.NeedReboot = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetRejectedValue(v []*string) *DescribeParametersResponseBodyParameters {
	s.RejectedValue = v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetValueType(v string) *DescribeParametersResponseBodyParameters {
	s.ValueType = &v
	return s
}

type DescribeParametersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponse) SetHeaders(v map[string]*string) *DescribeParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeParametersResponse) SetBody(v *DescribeParametersResponseBody) *DescribeParametersResponse {
	s.Body = v
	return s
}

type DescribeParametersHistoryRequest struct {
	// 参数类型。 当前支持集群（CLUSTER)和租户（TENANT）
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// 参数类型的资源标识。 如果为集群参数可以不填，若为租户的参数，则传入租户的TenantId。
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
	// 参数历史查看的结束时间。
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 分页查询的页码。 起始值：1 默认值：1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时设置的每页行数。 最大值：100 默认值：10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 参数历史查看的起始时间。
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeParametersHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeParametersHistoryRequest) SetDimension(v string) *DescribeParametersHistoryRequest {
	s.Dimension = &v
	return s
}

func (s *DescribeParametersHistoryRequest) SetDimensionValue(v string) *DescribeParametersHistoryRequest {
	s.DimensionValue = &v
	return s
}

func (s *DescribeParametersHistoryRequest) SetEndTime(v string) *DescribeParametersHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeParametersHistoryRequest) SetInstanceId(v string) *DescribeParametersHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeParametersHistoryRequest) SetPageNumber(v int32) *DescribeParametersHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeParametersHistoryRequest) SetPageSize(v int32) *DescribeParametersHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeParametersHistoryRequest) SetStartTime(v string) *DescribeParametersHistoryRequest {
	s.StartTime = &v
	return s
}

type DescribeParametersHistoryResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 参数修改历史列表。
	Respond []*DescribeParametersHistoryResponseBodyRespond `json:"Respond,omitempty" xml:"Respond,omitempty" type:"Repeated"`
}

func (s DescribeParametersHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParametersHistoryResponseBody) SetRequestId(v string) *DescribeParametersHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParametersHistoryResponseBody) SetRespond(v []*DescribeParametersHistoryResponseBodyRespond) *DescribeParametersHistoryResponseBody {
	s.Respond = v
	return s
}

type DescribeParametersHistoryResponseBodyRespond struct {
	// 每页记录数。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 参数修改历史信息。
	Parameters []*DescribeParametersHistoryResponseBodyRespondParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// 查询到的参数修改历史记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeParametersHistoryResponseBodyRespond) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersHistoryResponseBodyRespond) GoString() string {
	return s.String()
}

func (s *DescribeParametersHistoryResponseBodyRespond) SetPageNumber(v int32) *DescribeParametersHistoryResponseBodyRespond {
	s.PageNumber = &v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespond) SetParameters(v []*DescribeParametersHistoryResponseBodyRespondParameters) *DescribeParametersHistoryResponseBodyRespond {
	s.Parameters = v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespond) SetTotalCount(v int32) *DescribeParametersHistoryResponseBodyRespond {
	s.TotalCount = &v
	return s
}

type DescribeParametersHistoryResponseBodyRespondParameters struct {
	// 参数修改的发起时间。
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 参数类型的资源标识。 如果为集群则为DEFAULT_DIMENSION_VALUE，若为租户的参数，则传入租户的TenantId。
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
	// 参数名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数修改后的值.
	NewValue *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	// 参数修改前的值。
	OldValue *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
	// 修改状态。 - APPLIED：成功 - SCHEDULING：待修改。
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 参数修改的生效时间。
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeParametersHistoryResponseBodyRespondParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersHistoryResponseBodyRespondParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersHistoryResponseBodyRespondParameters) SetCreateTime(v string) *DescribeParametersHistoryResponseBodyRespondParameters {
	s.CreateTime = &v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespondParameters) SetDimensionValue(v string) *DescribeParametersHistoryResponseBodyRespondParameters {
	s.DimensionValue = &v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespondParameters) SetName(v string) *DescribeParametersHistoryResponseBodyRespondParameters {
	s.Name = &v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespondParameters) SetNewValue(v string) *DescribeParametersHistoryResponseBodyRespondParameters {
	s.NewValue = &v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespondParameters) SetOldValue(v string) *DescribeParametersHistoryResponseBodyRespondParameters {
	s.OldValue = &v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespondParameters) SetStatus(v string) *DescribeParametersHistoryResponseBodyRespondParameters {
	s.Status = &v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespondParameters) SetUpdateTime(v string) *DescribeParametersHistoryResponseBodyRespondParameters {
	s.UpdateTime = &v
	return s
}

type DescribeParametersHistoryResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeParametersHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParametersHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeParametersHistoryResponse) SetHeaders(v map[string]*string) *DescribeParametersHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeParametersHistoryResponse) SetBody(v *DescribeParametersHistoryResponseBody) *DescribeParametersHistoryResponse {
	s.Body = v
	return s
}

type DescribeRecommendIndexRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// SQLID
	SQLId *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// 租户ID
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeRecommendIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecommendIndexRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecommendIndexRequest) SetInstanceId(v string) *DescribeRecommendIndexRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRecommendIndexRequest) SetSQLId(v string) *DescribeRecommendIndexRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeRecommendIndexRequest) SetTenantId(v string) *DescribeRecommendIndexRequest {
	s.TenantId = &v
	return s
}

type DescribeRecommendIndexResponseBody struct {
	// 推荐索引信息
	RecommendIndex *DescribeRecommendIndexResponseBodyRecommendIndex `json:"RecommendIndex,omitempty" xml:"RecommendIndex,omitempty" type:"Struct"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRecommendIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecommendIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecommendIndexResponseBody) SetRecommendIndex(v *DescribeRecommendIndexResponseBodyRecommendIndex) *DescribeRecommendIndexResponseBody {
	s.RecommendIndex = v
	return s
}

func (s *DescribeRecommendIndexResponseBody) SetRequestId(v string) *DescribeRecommendIndexResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRecommendIndexResponseBodyRecommendIndex struct {
	// 建议索引，如果是主键，就是 PRIMARY，如果不是主键，以用户取名为准
	SuggestIndex *string `json:"SuggestIndex,omitempty" xml:"SuggestIndex,omitempty"`
	// 表信息
	TableList *string `json:"TableList,omitempty" xml:"TableList,omitempty"`
	// 租户模式
	TenantMode *string `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
}

func (s DescribeRecommendIndexResponseBodyRecommendIndex) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecommendIndexResponseBodyRecommendIndex) GoString() string {
	return s.String()
}

func (s *DescribeRecommendIndexResponseBodyRecommendIndex) SetSuggestIndex(v string) *DescribeRecommendIndexResponseBodyRecommendIndex {
	s.SuggestIndex = &v
	return s
}

func (s *DescribeRecommendIndexResponseBodyRecommendIndex) SetTableList(v string) *DescribeRecommendIndexResponseBodyRecommendIndex {
	s.TableList = &v
	return s
}

func (s *DescribeRecommendIndexResponseBodyRecommendIndex) SetTenantMode(v string) *DescribeRecommendIndexResponseBodyRecommendIndex {
	s.TenantMode = &v
	return s
}

type DescribeRecommendIndexResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRecommendIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecommendIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecommendIndexResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecommendIndexResponse) SetHeaders(v map[string]*string) *DescribeRecommendIndexResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecommendIndexResponse) SetBody(v *DescribeRecommendIndexResponseBody) *DescribeRecommendIndexResponse {
	s.Body = v
	return s
}

type DescribeSQLDetailsRequest struct {
	// SQLID
	SQLId *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// 租户ID
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeSQLDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLDetailsRequest) SetSQLId(v string) *DescribeSQLDetailsRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSQLDetailsRequest) SetTenantId(v string) *DescribeSQLDetailsRequest {
	s.TenantId = &v
	return s
}

type DescribeSQLDetailsResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// SQL详情
	SQLDetails []*DescribeSQLDetailsResponseBodySQLDetails `json:"SQLDetails,omitempty" xml:"SQLDetails,omitempty" type:"Repeated"`
}

func (s DescribeSQLDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLDetailsResponseBody) SetRequestId(v string) *DescribeSQLDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLDetailsResponseBody) SetSQLDetails(v []*DescribeSQLDetailsResponseBodySQLDetails) *DescribeSQLDetailsResponseBody {
	s.SQLDetails = v
	return s
}

type DescribeSQLDetailsResponseBodySQLDetails struct {
	// 数据库名
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// SQL文本
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// 用户名
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSQLDetailsResponseBodySQLDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLDetailsResponseBodySQLDetails) GoString() string {
	return s.String()
}

func (s *DescribeSQLDetailsResponseBodySQLDetails) SetDbName(v string) *DescribeSQLDetailsResponseBodySQLDetails {
	s.DbName = &v
	return s
}

func (s *DescribeSQLDetailsResponseBodySQLDetails) SetSQLText(v string) *DescribeSQLDetailsResponseBodySQLDetails {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLDetailsResponseBodySQLDetails) SetUserName(v string) *DescribeSQLDetailsResponseBodySQLDetails {
	s.UserName = &v
	return s
}

type DescribeSQLDetailsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSQLDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLDetailsResponse) SetHeaders(v map[string]*string) *DescribeSQLDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLDetailsResponse) SetBody(v *DescribeSQLDetailsResponseBody) *DescribeSQLDetailsResponse {
	s.Body = v
	return s
}

type DescribeSQLHistoryListRequest struct {
	// 参数历史查看的结束时间。
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 页码
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页数量
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// SQLID
	SQLId *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// 参数历史查看的起始时间。
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 租户ID
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeSQLHistoryListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLHistoryListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLHistoryListRequest) SetEndTime(v string) *DescribeSQLHistoryListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLHistoryListRequest) SetPageNumber(v int32) *DescribeSQLHistoryListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLHistoryListRequest) SetPageSize(v int32) *DescribeSQLHistoryListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLHistoryListRequest) SetSQLId(v string) *DescribeSQLHistoryListRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSQLHistoryListRequest) SetStartTime(v string) *DescribeSQLHistoryListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLHistoryListRequest) SetTenantId(v string) *DescribeSQLHistoryListRequest {
	s.TenantId = &v
	return s
}

type DescribeSQLHistoryListResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// SQL历史信息
	SQLHistoryList *DescribeSQLHistoryListResponseBodySQLHistoryList `json:"SQLHistoryList,omitempty" xml:"SQLHistoryList,omitempty" type:"Struct"`
}

func (s DescribeSQLHistoryListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLHistoryListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLHistoryListResponseBody) SetRequestId(v string) *DescribeSQLHistoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBody) SetSQLHistoryList(v *DescribeSQLHistoryListResponseBodySQLHistoryList) *DescribeSQLHistoryListResponseBody {
	s.SQLHistoryList = v
	return s
}

type DescribeSQLHistoryListResponseBodySQLHistoryList struct {
	// 数量
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 列表
	List []*DescribeSQLHistoryListResponseBodySQLHistoryListList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeSQLHistoryListResponseBodySQLHistoryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLHistoryListResponseBodySQLHistoryList) GoString() string {
	return s.String()
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryList) SetCount(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryList {
	s.Count = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryList) SetList(v []*DescribeSQLHistoryListResponseBodySQLHistoryListList) *DescribeSQLHistoryListResponseBodySQLHistoryList {
	s.List = v
	return s
}

type DescribeSQLHistoryListResponseBodySQLHistoryListList struct {
	// 影响行数
	AffectedRows *int64 `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
	// 客户端等待
	AppWaitTime *float32 `json:"AppWaitTime,omitempty" xml:"AppWaitTime,omitempty"`
	// BlockCache命中次数
	BlockCacheHit *int64 `json:"BlockCacheHit,omitempty" xml:"BlockCacheHit,omitempty"`
	// BlockIndexCache命中次数
	BlockIndexCacheHit *int64 `json:"BlockIndexCacheHit,omitempty" xml:"BlockIndexCacheHit,omitempty"`
	// BloomFilterCache命中次数
	BloomFilterCacheHit *int64 `json:"BloomFilterCacheHit,omitempty" xml:"BloomFilterCacheHit,omitempty"`
	// 客户端IP
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// 并发等待
	ConcurrencyWaitTime *float32 `json:"ConcurrencyWaitTime,omitempty" xml:"ConcurrencyWaitTime,omitempty"`
	// 平均CPU时间
	CpuTime *float32 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	// 数据库名
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// 解码等待
	DecodeTime *float32 `json:"DecodeTime,omitempty" xml:"DecodeTime,omitempty"`
	// 物理读
	DiskRead *int64 `json:"DiskRead,omitempty" xml:"DiskRead,omitempty"`
	// 平均响应时间
	ElapsedTime *float32 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 结束时间（零时区）
	EndTimeUTCString *string `json:"EndTimeUTCString,omitempty" xml:"EndTimeUTCString,omitempty"`
	// 等待事件
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// 每秒次数
	ExecPerSecond *int64 `json:"ExecPerSecond,omitempty" xml:"ExecPerSecond,omitempty"`
	// 内部执行时间
	ExecuteTime *float32 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// 执行次数
	Executions *int64 `json:"Executions,omitempty" xml:"Executions,omitempty"`
	// 失败次数
	FailTimes *int64 `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	// 硬解析时间
	GetPlanTime *float32 `json:"GetPlanTime,omitempty" xml:"GetPlanTime,omitempty"`
	// IO等待
	IOWaitTime *float32 `json:"IOWaitTime,omitempty" xml:"IOWaitTime,omitempty"`
	// 逻辑读
	LogicalRead *int64 `json:"LogicalRead,omitempty" xml:"LogicalRead,omitempty"`
	// 最大CPU时间
	MaxCpuTime *float32 `json:"MaxCpuTime,omitempty" xml:"MaxCpuTime,omitempty"`
	// 最大响应时间
	MaxElapsedTime *float32 `json:"MaxElapsedTime,omitempty" xml:"MaxElapsedTime,omitempty"`
	// 读内存行数
	MemstoreReadRowCount *int64 `json:"MemstoreReadRowCount,omitempty" xml:"MemstoreReadRowCount,omitempty"`
	// 未命中计划次数
	MissPlans *int64 `json:"MissPlans,omitempty" xml:"MissPlans,omitempty"`
	// 网络等待
	NetWaitTime *float32 `json:"NetWaitTime,omitempty" xml:"NetWaitTime,omitempty"`
	// 节点IP
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// 队列等待
	QueueTime *float32 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	// RPC次数
	RPCCount *int64 `json:"RPCCount,omitempty" xml:"RPCCount,omitempty"`
	// 远程计划数
	RemotePlans *int64 `json:"RemotePlans,omitempty" xml:"RemotePlans,omitempty"`
	// 重试次数
	RetryCount *int64 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// 返回行数
	ReturnRows *int64 `json:"ReturnRows,omitempty" xml:"ReturnRows,omitempty"`
	// RowCache命中次数
	RowCacheHit *int64 `json:"RowCacheHit,omitempty" xml:"RowCacheHit,omitempty"`
	// 调度时间
	ScheduleTime *float32 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// 读磁盘行数
	SsstoreReadRowCount *int64 `json:"SsstoreReadRowCount,omitempty" xml:"SsstoreReadRowCount,omitempty"`
	// 内部等待
	TotalWaitTime *float32 `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	// 用户名
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSQLHistoryListResponseBodySQLHistoryListList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLHistoryListResponseBodySQLHistoryListList) GoString() string {
	return s.String()
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetAffectedRows(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.AffectedRows = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetAppWaitTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.AppWaitTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetBlockCacheHit(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.BlockCacheHit = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetBlockIndexCacheHit(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.BlockIndexCacheHit = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetBloomFilterCacheHit(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.BloomFilterCacheHit = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetClientIp(v string) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.ClientIp = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetConcurrencyWaitTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.ConcurrencyWaitTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetCpuTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.CpuTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetDbName(v string) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.DbName = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetDecodeTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.DecodeTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetDiskRead(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.DiskRead = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetElapsedTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetEndTime(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetEndTimeUTCString(v string) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.EndTimeUTCString = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetEvent(v string) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.Event = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetExecPerSecond(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.ExecPerSecond = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetExecuteTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetExecutions(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.Executions = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetFailTimes(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.FailTimes = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetGetPlanTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.GetPlanTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetIOWaitTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.IOWaitTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetLogicalRead(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.LogicalRead = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetMaxCpuTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.MaxCpuTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetMaxElapsedTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.MaxElapsedTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetMemstoreReadRowCount(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.MemstoreReadRowCount = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetMissPlans(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.MissPlans = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetNetWaitTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.NetWaitTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetNodeIp(v string) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.NodeIp = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetQueueTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.QueueTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetRPCCount(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.RPCCount = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetRemotePlans(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.RemotePlans = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetRetryCount(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.RetryCount = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetReturnRows(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.ReturnRows = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetRowCacheHit(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.RowCacheHit = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetScheduleTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.ScheduleTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetSsstoreReadRowCount(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.SsstoreReadRowCount = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetTotalWaitTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.TotalWaitTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetUserName(v string) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.UserName = &v
	return s
}

type DescribeSQLHistoryListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSQLHistoryListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLHistoryListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLHistoryListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLHistoryListResponse) SetHeaders(v map[string]*string) *DescribeSQLHistoryListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLHistoryListResponse) SetBody(v *DescribeSQLHistoryListResponseBody) *DescribeSQLHistoryListResponse {
	s.Body = v
	return s
}

type DescribeSQLPlansRequest struct {
	// SQLID
	SQLId *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// 租户ID
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeSQLPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlansRequest) SetSQLId(v string) *DescribeSQLPlansRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSQLPlansRequest) SetTenantId(v string) *DescribeSQLPlansRequest {
	s.TenantId = &v
	return s
}

type DescribeSQLPlansResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 执行计划信息
	SQLPlans []*DescribeSQLPlansResponseBodySQLPlans `json:"SQLPlans,omitempty" xml:"SQLPlans,omitempty" type:"Repeated"`
}

func (s DescribeSQLPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlansResponseBody) SetRequestId(v string) *DescribeSQLPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLPlansResponseBody) SetSQLPlans(v []*DescribeSQLPlansResponseBodySQLPlans) *DescribeSQLPlansResponseBody {
	s.SQLPlans = v
	return s
}

type DescribeSQLPlansResponseBodySQLPlans struct {
	// 平均执行时间 (ms)
	AvgExecutionMS *float32 `json:"AvgExecutionMS,omitempty" xml:"AvgExecutionMS,omitempty"`
	// 平均执行时间
	AvgExecutionTimeMS *int64 `json:"AvgExecutionTimeMS,omitempty" xml:"AvgExecutionTimeMS,omitempty"`
	// 首次加载时间
	FirstLoadTime *int64 `json:"FirstLoadTime,omitempty" xml:"FirstLoadTime,omitempty"`
	// 首次加载时间(零时区)
	FirstLoadTimeUTCString *string `json:"FirstLoadTimeUTCString,omitempty" xml:"FirstLoadTimeUTCString,omitempty"`
	// 命中次数
	HitCount *int32 `json:"HitCount,omitempty" xml:"HitCount,omitempty"`
	// 合并版本
	MergedVersion *int32 `json:"MergedVersion,omitempty" xml:"MergedVersion,omitempty"`
	// 节点ip
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// Outline数据
	OutlineData *string `json:"OutlineData,omitempty" xml:"OutlineData,omitempty"`
	// OutlineID
	OutlineId *int64 `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
	// 绑定时间
	OutlineTime *int64 `json:"OutlineTime,omitempty" xml:"OutlineTime,omitempty"`
	// 绑定时间(零时区)
	OutlineTimeUTCString *string `json:"OutlineTimeUTCString,omitempty" xml:"OutlineTimeUTCString,omitempty"`
	// SQL的完整执行计划
	PlanFull *string `json:"PlanFull,omitempty" xml:"PlanFull,omitempty"`
	// SQL执行计划在数据库内部的ID
	PlanId *int32 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// 计划信息
	PlanInfo *string `json:"PlanInfo,omitempty" xml:"PlanInfo,omitempty"`
	// SQL执行计划在诊断系统内部的唯一标识
	PlanUnionHash *string `json:"PlanUnionHash,omitempty" xml:"PlanUnionHash,omitempty"`
	// 查询sql
	QuerySQL *string `json:"QuerySQL,omitempty" xml:"QuerySQL,omitempty"`
}

func (s DescribeSQLPlansResponseBodySQLPlans) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlansResponseBodySQLPlans) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetAvgExecutionMS(v float32) *DescribeSQLPlansResponseBodySQLPlans {
	s.AvgExecutionMS = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetAvgExecutionTimeMS(v int64) *DescribeSQLPlansResponseBodySQLPlans {
	s.AvgExecutionTimeMS = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetFirstLoadTime(v int64) *DescribeSQLPlansResponseBodySQLPlans {
	s.FirstLoadTime = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetFirstLoadTimeUTCString(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.FirstLoadTimeUTCString = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetHitCount(v int32) *DescribeSQLPlansResponseBodySQLPlans {
	s.HitCount = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetMergedVersion(v int32) *DescribeSQLPlansResponseBodySQLPlans {
	s.MergedVersion = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetNodeIp(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.NodeIp = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetOutlineData(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.OutlineData = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetOutlineId(v int64) *DescribeSQLPlansResponseBodySQLPlans {
	s.OutlineId = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetOutlineTime(v int64) *DescribeSQLPlansResponseBodySQLPlans {
	s.OutlineTime = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetOutlineTimeUTCString(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.OutlineTimeUTCString = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetPlanFull(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.PlanFull = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetPlanId(v int32) *DescribeSQLPlansResponseBodySQLPlans {
	s.PlanId = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetPlanInfo(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.PlanInfo = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetPlanUnionHash(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.PlanUnionHash = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetQuerySQL(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.QuerySQL = &v
	return s
}

type DescribeSQLPlansResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSQLPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlansResponse) SetHeaders(v map[string]*string) *DescribeSQLPlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLPlansResponse) SetBody(v *DescribeSQLPlansResponseBody) *DescribeSQLPlansResponse {
	s.Body = v
	return s
}

type DescribeSecurityIpGroupsRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeSecurityIpGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpGroupsRequest) SetInstanceId(v string) *DescribeSecurityIpGroupsRequest {
	s.InstanceId = &v
	return s
}

type DescribeSecurityIpGroupsResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// IP白名单分组信息。
	SecurityIpGroups []*DescribeSecurityIpGroupsResponseBodySecurityIpGroups `json:"SecurityIpGroups,omitempty" xml:"SecurityIpGroups,omitempty" type:"Repeated"`
	// 查询到的白名单分组个数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSecurityIpGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpGroupsResponseBody) SetRequestId(v string) *DescribeSecurityIpGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityIpGroupsResponseBody) SetSecurityIpGroups(v []*DescribeSecurityIpGroupsResponseBodySecurityIpGroups) *DescribeSecurityIpGroupsResponseBody {
	s.SecurityIpGroups = v
	return s
}

func (s *DescribeSecurityIpGroupsResponseBody) SetTotalCount(v int32) *DescribeSecurityIpGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSecurityIpGroupsResponseBodySecurityIpGroups struct {
	// 安全组名称。
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	// IP安全白名单列表。其为一个Json格式的数组，数组中每个对象为一个IP字符串或者IP段。
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s DescribeSecurityIpGroupsResponseBodySecurityIpGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpGroupsResponseBodySecurityIpGroups) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpGroupsResponseBodySecurityIpGroups) SetSecurityIpGroupName(v string) *DescribeSecurityIpGroupsResponseBodySecurityIpGroups {
	s.SecurityIpGroupName = &v
	return s
}

func (s *DescribeSecurityIpGroupsResponseBodySecurityIpGroups) SetSecurityIps(v string) *DescribeSecurityIpGroupsResponseBodySecurityIpGroups {
	s.SecurityIps = &v
	return s
}

type DescribeSecurityIpGroupsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityIpGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityIpGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpGroupsResponse) SetHeaders(v map[string]*string) *DescribeSecurityIpGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityIpGroupsResponse) SetBody(v *DescribeSecurityIpGroupsResponseBody) *DescribeSecurityIpGroupsResponse {
	s.Body = v
	return s
}

type DescribeSlowSQLHistoryListRequest struct {
	// 结束时间
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// SQL唯一标识
	SQLId *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 租户名
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeSlowSQLHistoryListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLHistoryListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLHistoryListRequest) SetEndTime(v string) *DescribeSlowSQLHistoryListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListRequest) SetPageNumber(v int32) *DescribeSlowSQLHistoryListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowSQLHistoryListRequest) SetPageSize(v int32) *DescribeSlowSQLHistoryListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowSQLHistoryListRequest) SetSQLId(v string) *DescribeSlowSQLHistoryListRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSlowSQLHistoryListRequest) SetStartTime(v string) *DescribeSlowSQLHistoryListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListRequest) SetTenantId(v string) *DescribeSlowSQLHistoryListRequest {
	s.TenantId = &v
	return s
}

type DescribeSlowSQLHistoryListResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 慢SQL历史列表
	SlowSQLHistoryList *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList `json:"SlowSQLHistoryList,omitempty" xml:"SlowSQLHistoryList,omitempty" type:"Struct"`
}

func (s DescribeSlowSQLHistoryListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLHistoryListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLHistoryListResponseBody) SetRequestId(v string) *DescribeSlowSQLHistoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBody) SetSlowSQLHistoryList(v *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList) *DescribeSlowSQLHistoryListResponseBody {
	s.SlowSQLHistoryList = v
	return s
}

type DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList struct {
	Count *int64                                                          `json:"Count,omitempty" xml:"Count,omitempty"`
	List  []*DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList) SetCount(v int64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList {
	s.Count = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList) SetList(v []*DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList {
	s.List = v
	return s
}

type DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList struct {
	AffectedRows         *float64 `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
	AppWaitTime          *float64 `json:"AppWaitTime,omitempty" xml:"AppWaitTime,omitempty"`
	BlockCacheHit        *float64 `json:"BlockCacheHit,omitempty" xml:"BlockCacheHit,omitempty"`
	BlockIndexCacheHit   *float64 `json:"BlockIndexCacheHit,omitempty" xml:"BlockIndexCacheHit,omitempty"`
	BloomFilterCacheHit  *float64 `json:"BloomFilterCacheHit,omitempty" xml:"BloomFilterCacheHit,omitempty"`
	ClientIp             *string  `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ConcurrencyWaitTime  *float64 `json:"ConcurrencyWaitTime,omitempty" xml:"ConcurrencyWaitTime,omitempty"`
	CpuTime              *float64 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	DbName               *string  `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DecodeTime           *float64 `json:"DecodeTime,omitempty" xml:"DecodeTime,omitempty"`
	DiskRead             *float64 `json:"DiskRead,omitempty" xml:"DiskRead,omitempty"`
	ElapsedTime          *float64 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	EndTimeUTCString     *string  `json:"EndTimeUTCString,omitempty" xml:"EndTimeUTCString,omitempty"`
	Event                *string  `json:"Event,omitempty" xml:"Event,omitempty"`
	ExecPerSecond        *float64 `json:"ExecPerSecond,omitempty" xml:"ExecPerSecond,omitempty"`
	ExecuteTime          *float64 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	Executions           *float64 `json:"Executions,omitempty" xml:"Executions,omitempty"`
	FailTimes            *float64 `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	GetplanTime          *float64 `json:"GetplanTime,omitempty" xml:"GetplanTime,omitempty"`
	IOWaitTime           *float64 `json:"IOWaitTime,omitempty" xml:"IOWaitTime,omitempty"`
	LogicalRead          *float64 `json:"LogicalRead,omitempty" xml:"LogicalRead,omitempty"`
	MaxCpuTime           *float64 `json:"MaxCpuTime,omitempty" xml:"MaxCpuTime,omitempty"`
	MaxElapsedTime       *float64 `json:"MaxElapsedTime,omitempty" xml:"MaxElapsedTime,omitempty"`
	MemstoreReadRowCount *float64 `json:"MemstoreReadRowCount,omitempty" xml:"MemstoreReadRowCount,omitempty"`
	MissPlans            *float64 `json:"MissPlans,omitempty" xml:"MissPlans,omitempty"`
	NetwaitTime          *float64 `json:"NetwaitTime,omitempty" xml:"NetwaitTime,omitempty"`
	NodeIp               *string  `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	QueueTime            *float64 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	RPCCount             *float64 `json:"RPCCount,omitempty" xml:"RPCCount,omitempty"`
	RemotePlans          *float64 `json:"RemotePlans,omitempty" xml:"RemotePlans,omitempty"`
	RetryCount           *float64 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	ReturnRows           *float64 `json:"ReturnRows,omitempty" xml:"ReturnRows,omitempty"`
	RowCacheHit          *float64 `json:"RowCacheHit,omitempty" xml:"RowCacheHit,omitempty"`
	ScheduleTime         *float64 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	SqlId                *string  `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	SqlType              *string  `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	SsstoreReadRowCount  *float64 `json:"SsstoreReadRowCount,omitempty" xml:"SsstoreReadRowCount,omitempty"`
	TenantName           *string  `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	TotalWaitTime        *float64 `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	UserName             *string  `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetAffectedRows(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.AffectedRows = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetAppWaitTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.AppWaitTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetBlockCacheHit(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.BlockCacheHit = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetBlockIndexCacheHit(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.BlockIndexCacheHit = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetBloomFilterCacheHit(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.BloomFilterCacheHit = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetClientIp(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.ClientIp = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetConcurrencyWaitTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.ConcurrencyWaitTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetCpuTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.CpuTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetDbName(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.DbName = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetDecodeTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.DecodeTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetDiskRead(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.DiskRead = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetElapsedTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetEndTimeUTCString(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.EndTimeUTCString = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetEvent(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.Event = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetExecPerSecond(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.ExecPerSecond = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetExecuteTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetExecutions(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.Executions = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetFailTimes(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.FailTimes = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetGetplanTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.GetplanTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetIOWaitTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.IOWaitTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetLogicalRead(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.LogicalRead = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetMaxCpuTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.MaxCpuTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetMaxElapsedTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.MaxElapsedTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetMemstoreReadRowCount(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.MemstoreReadRowCount = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetMissPlans(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.MissPlans = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetNetwaitTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.NetwaitTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetNodeIp(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.NodeIp = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetQueueTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.QueueTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetRPCCount(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.RPCCount = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetRemotePlans(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.RemotePlans = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetRetryCount(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.RetryCount = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetReturnRows(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.ReturnRows = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetRowCacheHit(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.RowCacheHit = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetScheduleTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.ScheduleTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetSqlId(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.SqlId = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetSqlType(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.SqlType = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetSsstoreReadRowCount(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.SsstoreReadRowCount = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetTenantName(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.TenantName = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetTotalWaitTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.TotalWaitTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetUserName(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.UserName = &v
	return s
}

type DescribeSlowSQLHistoryListResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlowSQLHistoryListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowSQLHistoryListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLHistoryListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLHistoryListResponse) SetHeaders(v map[string]*string) *DescribeSlowSQLHistoryListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowSQLHistoryListResponse) SetBody(v *DescribeSlowSQLHistoryListResponseBody) *DescribeSlowSQLHistoryListResponse {
	s.Body = v
	return s
}

type DescribeSlowSQLListRequest struct {
	// 数据库名称
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 过滤条件
	FilterCondition map[string]interface{} `json:"FilterCondition,omitempty" xml:"FilterCondition,omitempty"`
	// 节点ip
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// 页码
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// SQL唯一标识
	SQLId *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// 查询关键字
	SearchKeyWord *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	// 查询参数
	SearchParameter *string `json:"SearchParameter,omitempty" xml:"SearchParameter,omitempty"`
	// 查询规则
	SearchRule *string `json:"SearchRule,omitempty" xml:"SearchRule,omitempty"`
	// 查询值
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	// 排序列
	SortColumn *string `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	// 排序顺序
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 租户ID
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeSlowSQLListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLListRequest) SetDbName(v string) *DescribeSlowSQLListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetEndTime(v string) *DescribeSlowSQLListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetFilterCondition(v map[string]interface{}) *DescribeSlowSQLListRequest {
	s.FilterCondition = v
	return s
}

func (s *DescribeSlowSQLListRequest) SetNodeIp(v string) *DescribeSlowSQLListRequest {
	s.NodeIp = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetPageNumber(v int32) *DescribeSlowSQLListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetPageSize(v int32) *DescribeSlowSQLListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetSQLId(v string) *DescribeSlowSQLListRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetSearchKeyWord(v string) *DescribeSlowSQLListRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetSearchParameter(v string) *DescribeSlowSQLListRequest {
	s.SearchParameter = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetSearchRule(v string) *DescribeSlowSQLListRequest {
	s.SearchRule = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetSearchValue(v string) *DescribeSlowSQLListRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetSortColumn(v string) *DescribeSlowSQLListRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetSortOrder(v string) *DescribeSlowSQLListRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetStartTime(v string) *DescribeSlowSQLListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetTenantId(v string) *DescribeSlowSQLListRequest {
	s.TenantId = &v
	return s
}

type DescribeSlowSQLListShrinkRequest struct {
	// 数据库名称
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 过滤条件
	FilterConditionShrink *string `json:"FilterCondition,omitempty" xml:"FilterCondition,omitempty"`
	// 节点ip
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// 页码
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// SQL唯一标识
	SQLId *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// 查询关键字
	SearchKeyWord *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	// 查询参数
	SearchParameter *string `json:"SearchParameter,omitempty" xml:"SearchParameter,omitempty"`
	// 查询规则
	SearchRule *string `json:"SearchRule,omitempty" xml:"SearchRule,omitempty"`
	// 查询值
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	// 排序列
	SortColumn *string `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	// 排序顺序
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 租户ID
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeSlowSQLListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLListShrinkRequest) SetDbName(v string) *DescribeSlowSQLListShrinkRequest {
	s.DbName = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetEndTime(v string) *DescribeSlowSQLListShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetFilterConditionShrink(v string) *DescribeSlowSQLListShrinkRequest {
	s.FilterConditionShrink = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetNodeIp(v string) *DescribeSlowSQLListShrinkRequest {
	s.NodeIp = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetPageNumber(v int32) *DescribeSlowSQLListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetPageSize(v int32) *DescribeSlowSQLListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetSQLId(v string) *DescribeSlowSQLListShrinkRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetSearchKeyWord(v string) *DescribeSlowSQLListShrinkRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetSearchParameter(v string) *DescribeSlowSQLListShrinkRequest {
	s.SearchParameter = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetSearchRule(v string) *DescribeSlowSQLListShrinkRequest {
	s.SearchRule = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetSearchValue(v string) *DescribeSlowSQLListShrinkRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetSortColumn(v string) *DescribeSlowSQLListShrinkRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetSortOrder(v string) *DescribeSlowSQLListShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetStartTime(v string) *DescribeSlowSQLListShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetTenantId(v string) *DescribeSlowSQLListShrinkRequest {
	s.TenantId = &v
	return s
}

type DescribeSlowSQLListResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 慢SQL列表信息
	SlowSQLList []*DescribeSlowSQLListResponseBodySlowSQLList `json:"SlowSQLList,omitempty" xml:"SlowSQLList,omitempty" type:"Repeated"`
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSlowSQLListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLListResponseBody) SetRequestId(v string) *DescribeSlowSQLListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowSQLListResponseBody) SetSlowSQLList(v []*DescribeSlowSQLListResponseBodySlowSQLList) *DescribeSlowSQLListResponseBody {
	s.SlowSQLList = v
	return s
}

func (s *DescribeSlowSQLListResponseBody) SetTotalCount(v int64) *DescribeSlowSQLListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSlowSQLListResponseBodySlowSQLList struct {
	AffectedRows         *int64   `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
	AppWaitTime          *float32 `json:"AppWaitTime,omitempty" xml:"AppWaitTime,omitempty"`
	BlockCacheHit        *int64   `json:"BlockCacheHit,omitempty" xml:"BlockCacheHit,omitempty"`
	BlockIndexCacheHit   *int64   `json:"BlockIndexCacheHit,omitempty" xml:"BlockIndexCacheHit,omitempty"`
	BloomFilterCacheHit  *int64   `json:"BloomFilterCacheHit,omitempty" xml:"BloomFilterCacheHit,omitempty"`
	ClientIp             *string  `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ConcurrencyWaitTime  *float32 `json:"ConcurrencyWaitTime,omitempty" xml:"ConcurrencyWaitTime,omitempty"`
	CpuTime              *float32 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	DbName               *string  `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DecodeTime           *float32 `json:"DecodeTime,omitempty" xml:"DecodeTime,omitempty"`
	DiskRead             *int64   `json:"DiskRead,omitempty" xml:"DiskRead,omitempty"`
	ElapsedTime          *float32 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	Event                *string  `json:"Event,omitempty" xml:"Event,omitempty"`
	ExecPerSecond        *float32 `json:"ExecPerSecond,omitempty" xml:"ExecPerSecond,omitempty"`
	ExecuteTime          *float32 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	Executions           *int64   `json:"Executions,omitempty" xml:"Executions,omitempty"`
	FailTimes            *int64   `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	GetPlanTime          *float32 `json:"GetPlanTime,omitempty" xml:"GetPlanTime,omitempty"`
	IOWaitTime           *float32 `json:"IOWaitTime,omitempty" xml:"IOWaitTime,omitempty"`
	Key                  *int64   `json:"Key,omitempty" xml:"Key,omitempty"`
	LogicalRead          *int64   `json:"LogicalRead,omitempty" xml:"LogicalRead,omitempty"`
	MaxCpuTime           *float32 `json:"MaxCpuTime,omitempty" xml:"MaxCpuTime,omitempty"`
	MaxElapsedTime       *float32 `json:"MaxElapsedTime,omitempty" xml:"MaxElapsedTime,omitempty"`
	MemstoreReadRowCount *int64   `json:"MemstoreReadRowCount,omitempty" xml:"MemstoreReadRowCount,omitempty"`
	MissPlans            *int64   `json:"MissPlans,omitempty" xml:"MissPlans,omitempty"`
	NetWaitTime          *float32 `json:"NetWaitTime,omitempty" xml:"NetWaitTime,omitempty"`
	NodeIp               *string  `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	QueueTime            *float32 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	RPCCount             *int64   `json:"RPCCount,omitempty" xml:"RPCCount,omitempty"`
	RemotePlans          *int64   `json:"RemotePlans,omitempty" xml:"RemotePlans,omitempty"`
	RetryCount           *int64   `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	ReturnRows           *int64   `json:"ReturnRows,omitempty" xml:"ReturnRows,omitempty"`
	RowCacheHit          *int64   `json:"RowCacheHit,omitempty" xml:"RowCacheHit,omitempty"`
	SQLId                *string  `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	SQLText              *string  `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	SQLType              *int64   `json:"SQLType,omitempty" xml:"SQLType,omitempty"`
	ScheduleTime         *float32 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	SsstoreReadRowCount  *int64   `json:"SsstoreReadRowCount,omitempty" xml:"SsstoreReadRowCount,omitempty"`
	TotalWaitTime        *float32 `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	UserName             *string  `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSlowSQLListResponseBodySlowSQLList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLListResponseBodySlowSQLList) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetAffectedRows(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.AffectedRows = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetAppWaitTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.AppWaitTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetBlockCacheHit(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.BlockCacheHit = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetBlockIndexCacheHit(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.BlockIndexCacheHit = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetBloomFilterCacheHit(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.BloomFilterCacheHit = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetClientIp(v string) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.ClientIp = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetConcurrencyWaitTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.ConcurrencyWaitTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetCpuTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.CpuTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetDbName(v string) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.DbName = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetDecodeTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.DecodeTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetDiskRead(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.DiskRead = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetElapsedTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetEvent(v string) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.Event = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetExecPerSecond(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.ExecPerSecond = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetExecuteTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetExecutions(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.Executions = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetFailTimes(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.FailTimes = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetGetPlanTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.GetPlanTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetIOWaitTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.IOWaitTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetKey(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.Key = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetLogicalRead(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.LogicalRead = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetMaxCpuTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.MaxCpuTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetMaxElapsedTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.MaxElapsedTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetMemstoreReadRowCount(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.MemstoreReadRowCount = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetMissPlans(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.MissPlans = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetNetWaitTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.NetWaitTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetNodeIp(v string) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.NodeIp = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetQueueTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.QueueTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetRPCCount(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.RPCCount = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetRemotePlans(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.RemotePlans = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetRetryCount(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.RetryCount = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetReturnRows(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.ReturnRows = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetRowCacheHit(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.RowCacheHit = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetSQLId(v string) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.SQLId = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetSQLText(v string) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetSQLType(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.SQLType = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetScheduleTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.ScheduleTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetSsstoreReadRowCount(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.SsstoreReadRowCount = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetTotalWaitTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.TotalWaitTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetUserName(v string) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.UserName = &v
	return s
}

type DescribeSlowSQLListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlowSQLListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowSQLListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLListResponse) SetHeaders(v map[string]*string) *DescribeSlowSQLListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowSQLListResponse) SetBody(v *DescribeSlowSQLListResponseBody) *DescribeSlowSQLListResponse {
	s.Body = v
	return s
}

type DescribeSqlAuditsRequest struct {
	// 客户端IP
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// 数据库
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 执行耗时最大值
	ExecuteTimeMax *int64 `json:"ExecuteTimeMax,omitempty" xml:"ExecuteTimeMax,omitempty"`
	// 执行耗时最小值
	ExecuteTimeMin *int64  `json:"ExecuteTimeMin,omitempty" xml:"ExecuteTimeMin,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 返回的SQL文本是否需要脱敏
	NeedMasking *bool `json:"NeedMasking,omitempty" xml:"NeedMasking,omitempty"`
	// 节点，用逗号分隔
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// 操作类型，用逗号分隔
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// 分页页码
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 扫描记录数最大值
	ScanRowsMax *int64 `json:"ScanRowsMax,omitempty" xml:"ScanRowsMax,omitempty"`
	// 扫描记录数最小值
	ScanRowsMin *int64 `json:"ScanRowsMin,omitempty" xml:"ScanRowsMin,omitempty"`
	// 关键词
	SearchKeyWord *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	// 关键词关系 and/or
	SearchKeyWordMethod *string `json:"SearchKeyWordMethod,omitempty" xml:"SearchKeyWordMethod,omitempty"`
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantId  *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 用户名
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSqlAuditsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlAuditsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlAuditsRequest) SetClientIp(v string) *DescribeSqlAuditsRequest {
	s.ClientIp = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetDbName(v string) *DescribeSqlAuditsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetEndTime(v string) *DescribeSqlAuditsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetExecuteTimeMax(v int64) *DescribeSqlAuditsRequest {
	s.ExecuteTimeMax = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetExecuteTimeMin(v int64) *DescribeSqlAuditsRequest {
	s.ExecuteTimeMin = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetInstanceId(v string) *DescribeSqlAuditsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetNeedMasking(v bool) *DescribeSqlAuditsRequest {
	s.NeedMasking = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetNodeIp(v string) *DescribeSqlAuditsRequest {
	s.NodeIp = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetOperatorType(v string) *DescribeSqlAuditsRequest {
	s.OperatorType = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetPageNumber(v int64) *DescribeSqlAuditsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetPageSize(v int64) *DescribeSqlAuditsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetScanRowsMax(v int64) *DescribeSqlAuditsRequest {
	s.ScanRowsMax = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetScanRowsMin(v int64) *DescribeSqlAuditsRequest {
	s.ScanRowsMin = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetSearchKeyWord(v string) *DescribeSqlAuditsRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetSearchKeyWordMethod(v string) *DescribeSqlAuditsRequest {
	s.SearchKeyWordMethod = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetStartTime(v string) *DescribeSqlAuditsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetTenantId(v string) *DescribeSqlAuditsRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeSqlAuditsRequest) SetUserName(v string) *DescribeSqlAuditsRequest {
	s.UserName = &v
	return s
}

type DescribeSqlAuditsResponseBody struct {
	Data []*DescribeSqlAuditsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSqlAuditsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlAuditsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlAuditsResponseBody) SetData(v []*DescribeSqlAuditsResponseBodyData) *DescribeSqlAuditsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSqlAuditsResponseBody) SetRequestId(v string) *DescribeSqlAuditsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlAuditsResponseBody) SetTotalCount(v int64) *DescribeSqlAuditsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSqlAuditsResponseBodyData struct {
	// 更新行数
	AffectedRows *int64 `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
	// 客户端IP
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// 数据库
	DataBaseName *string `json:"DataBaseName,omitempty" xml:"DataBaseName,omitempty"`
	// 执行耗时 (返回的毫秒，无需转换)
	ExecuteTime *float64 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	Key         *int64   `json:"Key,omitempty" xml:"Key,omitempty"`
	// 操作类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// 扫描行数
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// SQLID
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// SQL语句(数据已经过脱敏)
	SqlText *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
	// 用户
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSqlAuditsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlAuditsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSqlAuditsResponseBodyData) SetAffectedRows(v int64) *DescribeSqlAuditsResponseBodyData {
	s.AffectedRows = &v
	return s
}

func (s *DescribeSqlAuditsResponseBodyData) SetClientIp(v string) *DescribeSqlAuditsResponseBodyData {
	s.ClientIp = &v
	return s
}

func (s *DescribeSqlAuditsResponseBodyData) SetDataBaseName(v string) *DescribeSqlAuditsResponseBodyData {
	s.DataBaseName = &v
	return s
}

func (s *DescribeSqlAuditsResponseBodyData) SetExecuteTime(v float64) *DescribeSqlAuditsResponseBodyData {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSqlAuditsResponseBodyData) SetKey(v int64) *DescribeSqlAuditsResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeSqlAuditsResponseBodyData) SetOperatorType(v string) *DescribeSqlAuditsResponseBodyData {
	s.OperatorType = &v
	return s
}

func (s *DescribeSqlAuditsResponseBodyData) SetScanRows(v int64) *DescribeSqlAuditsResponseBodyData {
	s.ScanRows = &v
	return s
}

func (s *DescribeSqlAuditsResponseBodyData) SetSqlId(v string) *DescribeSqlAuditsResponseBodyData {
	s.SqlId = &v
	return s
}

func (s *DescribeSqlAuditsResponseBodyData) SetSqlText(v string) *DescribeSqlAuditsResponseBodyData {
	s.SqlText = &v
	return s
}

func (s *DescribeSqlAuditsResponseBodyData) SetUserName(v string) *DescribeSqlAuditsResponseBodyData {
	s.UserName = &v
	return s
}

type DescribeSqlAuditsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSqlAuditsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSqlAuditsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlAuditsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlAuditsResponse) SetHeaders(v map[string]*string) *DescribeSqlAuditsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlAuditsResponse) SetBody(v *DescribeSqlAuditsResponseBody) *DescribeSqlAuditsResponse {
	s.Body = v
	return s
}

type DescribeTenantRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeTenantRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantRequest) GoString() string {
	return s.String()
}

func (s *DescribeTenantRequest) SetInstanceId(v string) *DescribeTenantRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTenantRequest) SetTenantId(v string) *DescribeTenantRequest {
	s.TenantId = &v
	return s
}

type DescribeTenantResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 租户信息。
	Tenant *DescribeTenantResponseBodyTenant `json:"Tenant,omitempty" xml:"Tenant,omitempty" type:"Struct"`
}

func (s DescribeTenantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBody) SetRequestId(v string) *DescribeTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTenantResponseBody) SetTenant(v *DescribeTenantResponseBodyTenant) *DescribeTenantResponseBody {
	s.Tenant = v
	return s
}

type DescribeTenantResponseBodyTenant struct {
	// Clog服务开启状态。 CLOSED: 关闭 ONLINE：服务中
	ClogServiceStatus *string `json:"ClogServiceStatus,omitempty" xml:"ClogServiceStatus,omitempty"`
	// 租户的创建时间。
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 租户的数据副本模式。 对于高可用版本：三地域为N-N-N，两可用区和单可用区均为N-N； 对于基础版为N。 其中，N为单可用区内的节点数。
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// 集群的部署类型。 multiple：多机房 single：单机房 dual：双机房
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// 租户描述信息。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 是否可使用Clog服务，如需开启，请工单联系。
	EnableClogService *bool `json:"EnableClogService,omitempty" xml:"EnableClogService,omitempty"`
	// 是否可开启租户公网地址。
	EnableInternetAddressService *bool `json:"EnableInternetAddressService,omitempty" xml:"EnableInternetAddressService,omitempty"`
	// 租户的主可用区。
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// 租户的状态。 PENDING_CREATE: 创建中 RESTORE: 恢复中 ONLINE: 运行中 SPEC_MODIFYING: 规格修改中 ALLOCATING_INTERNET_ADDRESS: 公网地址分配中 PENDING_OFFLINE_INTERNET_ADDRESS: 公网地址关闭中 PRIMARY_ZONE_MODIFYING: 切主可用区中 PARAMETER_MODIFYING: 参数修改中 WHITE_LIST_MODIFING: 白名单修改中
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 租户的连接访问信息。
	TenantConnections []*DescribeTenantResponseBodyTenantTenantConnections `json:"TenantConnections,omitempty" xml:"TenantConnections,omitempty" type:"Repeated"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 租户模式。 当前支持： Oracle：Oracle模式 MySQL: MySQL模式
	TenantMode *string `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
	// 租户名称。
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// 租户的资源信息。
	TenantResource *DescribeTenantResponseBodyTenantTenantResource `json:"TenantResource,omitempty" xml:"TenantResource,omitempty" type:"Struct"`
	// 租户可用区信息。
	TenantZones []*DescribeTenantResponseBodyTenantTenantZones `json:"TenantZones,omitempty" xml:"TenantZones,omitempty" type:"Repeated"`
	// 专有网络（VPC） ID。 如果没有合适的 VPC，请根据页面提示创建一个 VPC，详情参见 什么是专有网络
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeTenantResponseBodyTenant) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBodyTenant) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBodyTenant) SetClogServiceStatus(v string) *DescribeTenantResponseBodyTenant {
	s.ClogServiceStatus = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetCreateTime(v string) *DescribeTenantResponseBodyTenant {
	s.CreateTime = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetDeployMode(v string) *DescribeTenantResponseBodyTenant {
	s.DeployMode = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetDeployType(v string) *DescribeTenantResponseBodyTenant {
	s.DeployType = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetDescription(v string) *DescribeTenantResponseBodyTenant {
	s.Description = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetEnableClogService(v bool) *DescribeTenantResponseBodyTenant {
	s.EnableClogService = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetEnableInternetAddressService(v bool) *DescribeTenantResponseBodyTenant {
	s.EnableInternetAddressService = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetPrimaryZone(v string) *DescribeTenantResponseBodyTenant {
	s.PrimaryZone = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetStatus(v string) *DescribeTenantResponseBodyTenant {
	s.Status = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetTenantConnections(v []*DescribeTenantResponseBodyTenantTenantConnections) *DescribeTenantResponseBodyTenant {
	s.TenantConnections = v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetTenantId(v string) *DescribeTenantResponseBodyTenant {
	s.TenantId = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetTenantMode(v string) *DescribeTenantResponseBodyTenant {
	s.TenantMode = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetTenantName(v string) *DescribeTenantResponseBodyTenant {
	s.TenantName = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetTenantResource(v *DescribeTenantResponseBodyTenantTenantResource) *DescribeTenantResponseBodyTenant {
	s.TenantResource = v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetTenantZones(v []*DescribeTenantResponseBodyTenantTenantZones) *DescribeTenantResponseBodyTenant {
	s.TenantZones = v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetVpcId(v string) *DescribeTenantResponseBodyTenant {
	s.VpcId = &v
	return s
}

type DescribeTenantResponseBodyTenantTenantConnections struct {
	// 连接地址的服务模式。 ReadWrite：可读可写，提供强一致读写服务 ReadOnly：只读，保证数据最终一致性 Clog：事务日志服务
	ConnectionRole *string `json:"ConnectionRole,omitempty" xml:"ConnectionRole,omitempty"`
	// 租户连接对应的访问可用区列表。
	ConnectionZones []*string `json:"ConnectionZones,omitempty" xml:"ConnectionZones,omitempty" type:"Repeated"`
	// 访问租户的公网连接地址。
	InternetAddress *string `json:"InternetAddress,omitempty" xml:"InternetAddress,omitempty"`
	// 访问租户的公网连接开通状态。 CLOSED：关闭 ALLOCATING_INTERNET_ADDRESS：申请中 PENDING_OFFLINE_INTERNET_ADDRESS: 公网地址关闭中 ONLINE：服务中
	InternetAddressStatus *string `json:"InternetAddressStatus,omitempty" xml:"InternetAddressStatus,omitempty"`
	// 访问租户的公网连接端口。
	InternetPort *int32 `json:"InternetPort,omitempty" xml:"InternetPort,omitempty"`
	// 访问租户的私网连接地址。
	IntranetAddress *string `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	// 访问租户的连接地址对应的主可用区。
	IntranetAddressMasterZoneId *string `json:"IntranetAddressMasterZoneId,omitempty" xml:"IntranetAddressMasterZoneId,omitempty"`
	// 访问租户的连接地址对应的备可用区。
	IntranetAddressSlaveZoneId *string `json:"IntranetAddressSlaveZoneId,omitempty" xml:"IntranetAddressSlaveZoneId,omitempty"`
	// 访问租户的私网连接地址状态。 ONLINE: 服务中
	IntranetAddressStatus *string `json:"IntranetAddressStatus,omitempty" xml:"IntranetAddressStatus,omitempty"`
	// 访问租户的私网连接端口。
	IntranetPort *int32 `json:"IntranetPort,omitempty" xml:"IntranetPort,omitempty"`
	// 虚拟交换机（VSwitch） ID。
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// 专有网络（VPC） ID。
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeTenantResponseBodyTenantTenantConnections) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBodyTenantTenantConnections) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetConnectionRole(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.ConnectionRole = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetConnectionZones(v []*string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.ConnectionZones = v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetInternetAddress(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.InternetAddress = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetInternetAddressStatus(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.InternetAddressStatus = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetInternetPort(v int32) *DescribeTenantResponseBodyTenantTenantConnections {
	s.InternetPort = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetIntranetAddress(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.IntranetAddress = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetIntranetAddressMasterZoneId(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.IntranetAddressMasterZoneId = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetIntranetAddressSlaveZoneId(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.IntranetAddressSlaveZoneId = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetIntranetAddressStatus(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.IntranetAddressStatus = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetIntranetPort(v int32) *DescribeTenantResponseBodyTenantTenantConnections {
	s.IntranetPort = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetVSwitchId(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.VSwitchId = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetVpcId(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.VpcId = &v
	return s
}

type DescribeTenantResponseBodyTenantTenantResource struct {
	// 租户的CPU资源信息。
	Cpu *DescribeTenantResponseBodyTenantTenantResourceCpu `json:"Cpu,omitempty" xml:"Cpu,omitempty" type:"Struct"`
	// 租户磁盘资源信息。
	DiskSize *DescribeTenantResponseBodyTenantTenantResourceDiskSize `json:"DiskSize,omitempty" xml:"DiskSize,omitempty" type:"Struct"`
	// 租户内存资源信息。
	Memory *DescribeTenantResponseBodyTenantTenantResourceMemory `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	// 租户的unit个数。
	UnitNum *int32 `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s DescribeTenantResponseBodyTenantTenantResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBodyTenantTenantResource) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBodyTenantTenantResource) SetCpu(v *DescribeTenantResponseBodyTenantTenantResourceCpu) *DescribeTenantResponseBodyTenantTenantResource {
	s.Cpu = v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantResource) SetDiskSize(v *DescribeTenantResponseBodyTenantTenantResourceDiskSize) *DescribeTenantResponseBodyTenantTenantResource {
	s.DiskSize = v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantResource) SetMemory(v *DescribeTenantResponseBodyTenantTenantResourceMemory) *DescribeTenantResponseBodyTenantTenantResource {
	s.Memory = v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantResource) SetUnitNum(v int32) *DescribeTenantResponseBodyTenantTenantResource {
	s.UnitNum = &v
	return s
}

type DescribeTenantResponseBodyTenantTenantResourceCpu struct {
	// 租户总的CPU核数，单位：Core。
	TotalCpu *float32 `json:"TotalCpu,omitempty" xml:"TotalCpu,omitempty"`
	// 租户每个Unit的CPU核数，单位：Core。
	UnitCpu *float32 `json:"UnitCpu,omitempty" xml:"UnitCpu,omitempty"`
	// 租户已使用的CPU核数，单位：Core。
	UsedCpu *float32 `json:"UsedCpu,omitempty" xml:"UsedCpu,omitempty"`
}

func (s DescribeTenantResponseBodyTenantTenantResourceCpu) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBodyTenantTenantResourceCpu) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBodyTenantTenantResourceCpu) SetTotalCpu(v float32) *DescribeTenantResponseBodyTenantTenantResourceCpu {
	s.TotalCpu = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantResourceCpu) SetUnitCpu(v float32) *DescribeTenantResponseBodyTenantTenantResourceCpu {
	s.UnitCpu = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantResourceCpu) SetUsedCpu(v float32) *DescribeTenantResponseBodyTenantTenantResourceCpu {
	s.UsedCpu = &v
	return s
}

type DescribeTenantResponseBodyTenantTenantResourceDiskSize struct {
	// 租户已使用的磁盘大小，单位：GB。
	UsedDiskSize *float32 `json:"UsedDiskSize,omitempty" xml:"UsedDiskSize,omitempty"`
}

func (s DescribeTenantResponseBodyTenantTenantResourceDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBodyTenantTenantResourceDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBodyTenantTenantResourceDiskSize) SetUsedDiskSize(v float32) *DescribeTenantResponseBodyTenantTenantResourceDiskSize {
	s.UsedDiskSize = &v
	return s
}

type DescribeTenantResponseBodyTenantTenantResourceMemory struct {
	// 租户总的内存大小，单位：GB。
	TotalMemory *float32 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	// 租户每个Unit的内存大小，单位：GB。
	UnitMemory *float32 `json:"UnitMemory,omitempty" xml:"UnitMemory,omitempty"`
	// 租户已使用的内存，单位：GB。
	UsedMemory *float32 `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
}

func (s DescribeTenantResponseBodyTenantTenantResourceMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBodyTenantTenantResourceMemory) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBodyTenantTenantResourceMemory) SetTotalMemory(v float32) *DescribeTenantResponseBodyTenantTenantResourceMemory {
	s.TotalMemory = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantResourceMemory) SetUnitMemory(v float32) *DescribeTenantResponseBodyTenantTenantResourceMemory {
	s.UnitMemory = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantResourceMemory) SetUsedMemory(v float32) *DescribeTenantResponseBodyTenantTenantResourceMemory {
	s.UsedMemory = &v
	return s
}

type DescribeTenantResponseBodyTenantTenantZones struct {
	// 租户可用区所属的地域。
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 租户可用区ID。
	TenantZoneId *string `json:"TenantZoneId,omitempty" xml:"TenantZoneId,omitempty"`
	// 租户可用区的请求类型。 ReadWrite：可读可写 ReadOnly：只读 对于多机房高可用集群，其主可用区为ReadWrite。备可用区为ReadOnly; 对于单机房高可用集群，其所有可用区均为ReadWrite。
	TenantZoneRole *string `json:"TenantZoneRole,omitempty" xml:"TenantZoneRole,omitempty"`
}

func (s DescribeTenantResponseBodyTenantTenantZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBodyTenantTenantZones) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBodyTenantTenantZones) SetRegion(v string) *DescribeTenantResponseBodyTenantTenantZones {
	s.Region = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantZones) SetTenantZoneId(v string) *DescribeTenantResponseBodyTenantTenantZones {
	s.TenantZoneId = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantZones) SetTenantZoneRole(v string) *DescribeTenantResponseBodyTenantTenantZones {
	s.TenantZoneRole = &v
	return s
}

type DescribeTenantResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTenantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTenantResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponse) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponse) SetHeaders(v map[string]*string) *DescribeTenantResponse {
	s.Headers = v
	return s
}

func (s *DescribeTenantResponse) SetBody(v *DescribeTenantResponseBody) *DescribeTenantResponse {
	s.Body = v
	return s
}

type DescribeTenantMetricsRequest struct {
	// 参数历史查看的结束时间。
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 监控指标
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// 分页查询的页码。 起始值：1 默认值：1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时设置的每页行数。 最大值：100 默认值：10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 参数历史查看的起始时间。
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 租户ID列表
	TenantIdList *string `json:"TenantIdList,omitempty" xml:"TenantIdList,omitempty"`
	// 租户名称。 长度为2~20 个字符，支持英文字母、数字和下划线，区分大小写，必须以字母或下划线开头。 不可设置为 sys。
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s DescribeTenantMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTenantMetricsRequest) SetEndTime(v string) *DescribeTenantMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetInstanceId(v string) *DescribeTenantMetricsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetMetrics(v string) *DescribeTenantMetricsRequest {
	s.Metrics = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetPageNumber(v int32) *DescribeTenantMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetPageSize(v int32) *DescribeTenantMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetStartTime(v string) *DescribeTenantMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetTenantId(v string) *DescribeTenantMetricsRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetTenantIdList(v string) *DescribeTenantMetricsRequest {
	s.TenantIdList = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetTenantName(v string) *DescribeTenantMetricsRequest {
	s.TenantName = &v
	return s
}

type DescribeTenantMetricsResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 租户指标信息。
	TenantMetrics *string `json:"TenantMetrics,omitempty" xml:"TenantMetrics,omitempty"`
	// 总数量
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTenantMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTenantMetricsResponseBody) SetRequestId(v string) *DescribeTenantMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTenantMetricsResponseBody) SetTenantMetrics(v string) *DescribeTenantMetricsResponseBody {
	s.TenantMetrics = &v
	return s
}

func (s *DescribeTenantMetricsResponseBody) SetTotalCount(v int32) *DescribeTenantMetricsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTenantMetricsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTenantMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTenantMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTenantMetricsResponse) SetHeaders(v map[string]*string) *DescribeTenantMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTenantMetricsResponse) SetBody(v *DescribeTenantMetricsResponseBody) *DescribeTenantMetricsResponse {
	s.Body = v
	return s
}

type DescribeTenantUserRolesResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 用户的角色列表。角色包含以下几类： 读写权限（ReadWrite）：ALL PRIVILEGES ； 只读权限（ReadOnly）：SELECT  DDL权限（DDL）：CREATE,DROP,ALTER,SHOW VIEW,CREATE VIEW  DML权限（DML）：SELECT,INSERT,UPDATE,DELETE,SHOW VIEW。
	Role []*string `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s DescribeTenantUserRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTenantUserRolesResponseBody) SetRequestId(v string) *DescribeTenantUserRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTenantUserRolesResponseBody) SetRole(v []*string) *DescribeTenantUserRolesResponseBody {
	s.Role = v
	return s
}

type DescribeTenantUserRolesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTenantUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTenantUserRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantUserRolesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTenantUserRolesResponse) SetHeaders(v map[string]*string) *DescribeTenantUserRolesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTenantUserRolesResponse) SetBody(v *DescribeTenantUserRolesResponseBody) *DescribeTenantUserRolesResponse {
	s.Body = v
	return s
}

type DescribeTenantUsersRequest struct {
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// 分页查询的页码。 起始值：1 默认值：1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时设置的每页行数。 最大值：100 默认值：10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 查询列表的删选关键字。
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 数据库账号名称。 不能使用某些预留关键字，如 SYS、root等。
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeTenantUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeTenantUsersRequest) SetAddressType(v string) *DescribeTenantUsersRequest {
	s.AddressType = &v
	return s
}

func (s *DescribeTenantUsersRequest) SetPageNumber(v int32) *DescribeTenantUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTenantUsersRequest) SetPageSize(v int32) *DescribeTenantUsersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTenantUsersRequest) SetSearchKey(v string) *DescribeTenantUsersRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeTenantUsersRequest) SetTenantId(v string) *DescribeTenantUsersRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeTenantUsersRequest) SetUserName(v string) *DescribeTenantUsersRequest {
	s.UserName = &v
	return s
}

type DescribeTenantUsersResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 租户中的数据库账号信息
	TenantUsers []*DescribeTenantUsersResponseBodyTenantUsers `json:"TenantUsers,omitempty" xml:"TenantUsers,omitempty" type:"Repeated"`
	// 租户中的数据库账号总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTenantUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTenantUsersResponseBody) SetRequestId(v string) *DescribeTenantUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTenantUsersResponseBody) SetTenantUsers(v []*DescribeTenantUsersResponseBodyTenantUsers) *DescribeTenantUsersResponseBody {
	s.TenantUsers = v
	return s
}

func (s *DescribeTenantUsersResponseBody) SetTotalCount(v int32) *DescribeTenantUsersResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTenantUsersResponseBodyTenantUsers struct {
	// 账号具备的数据库权限信息。
	Databases []*DescribeTenantUsersResponseBodyTenantUsersDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// 数据库账号的描述信息。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 数据库账号名称。
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// 数据库账号的状态。 - LOCKED：锁定 - ONLINE：解锁
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
	// 数据库账号类型。 - Admin：超级账户 - Normal：普通账户
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeTenantUsersResponseBodyTenantUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantUsersResponseBodyTenantUsers) GoString() string {
	return s.String()
}

func (s *DescribeTenantUsersResponseBodyTenantUsers) SetDatabases(v []*DescribeTenantUsersResponseBodyTenantUsersDatabases) *DescribeTenantUsersResponseBodyTenantUsers {
	s.Databases = v
	return s
}

func (s *DescribeTenantUsersResponseBodyTenantUsers) SetDescription(v string) *DescribeTenantUsersResponseBodyTenantUsers {
	s.Description = &v
	return s
}

func (s *DescribeTenantUsersResponseBodyTenantUsers) SetUserName(v string) *DescribeTenantUsersResponseBodyTenantUsers {
	s.UserName = &v
	return s
}

func (s *DescribeTenantUsersResponseBodyTenantUsers) SetUserStatus(v string) *DescribeTenantUsersResponseBodyTenantUsers {
	s.UserStatus = &v
	return s
}

func (s *DescribeTenantUsersResponseBodyTenantUsers) SetUserType(v string) *DescribeTenantUsersResponseBodyTenantUsers {
	s.UserType = &v
	return s
}

type DescribeTenantUsersResponseBodyTenantUsersDatabases struct {
	// 数据库(schema)名称
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// 拥有的角色。 对于Oracle模式，角色为schema级别，其可分为 - ReadWrite：读写权限，包括CREATE TABLE CREATE VIEW CREATE PROCEDURE CREATE SYNONYM CREATE SEQUENCE CREATE TRIGGER CREATE TYPE CREATE SESSION EXECUTE ANY PROCEDURE CREATE ANY OUTLINE ALTER ANY OUTLINE DROP ANY OUTLINE CREATE ANY PROCEDURE ALTER ANY PROCEDURE DROP ANY PROCEDURE CREATE ANY SEQUENCE ALTER ANY SEQUENCE DROP ANY SEQUENCE CREATE ANY TYPE ALTER ANY TYPE DROP ANY TYPE SYSKM CREATE ANY TRIGGER ALTER ANY TRIGGER DROP ANY TRIGGER CREATE PROFILE ALTER PROFILE DROP PROFILE； - ReadOnly：只读权限，SELECT
	// 对于MySQL模式，角色为数据库（Database）级别，其有以下几类： - ReadWrite：读写权限，包括ALL PRIVILEGES； - ReadOnly：只读权限，包括SELECT - DDL: DDL权限，包括CREATE,DROP,ALTER,SHOW VIEW,CREATE VIEW - DML: DML权限，包括SELECT,INSERT,UPDATE,DELETE,SHOW VIEW。
	// * 另外Oracle账户对自己的schema有默认的读写权限，这里不会列出。
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 表的名称。
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeTenantUsersResponseBodyTenantUsersDatabases) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantUsersResponseBodyTenantUsersDatabases) GoString() string {
	return s.String()
}

func (s *DescribeTenantUsersResponseBodyTenantUsersDatabases) SetDatabase(v string) *DescribeTenantUsersResponseBodyTenantUsersDatabases {
	s.Database = &v
	return s
}

func (s *DescribeTenantUsersResponseBodyTenantUsersDatabases) SetRole(v string) *DescribeTenantUsersResponseBodyTenantUsersDatabases {
	s.Role = &v
	return s
}

func (s *DescribeTenantUsersResponseBodyTenantUsersDatabases) SetTable(v string) *DescribeTenantUsersResponseBodyTenantUsersDatabases {
	s.Table = &v
	return s
}

type DescribeTenantUsersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTenantUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTenantUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeTenantUsersResponse) SetHeaders(v map[string]*string) *DescribeTenantUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeTenantUsersResponse) SetBody(v *DescribeTenantUsersResponseBody) *DescribeTenantUsersResponse {
	s.Body = v
	return s
}

type DescribeTenantZonesReadRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeTenantZonesReadRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantZonesReadRequest) GoString() string {
	return s.String()
}

func (s *DescribeTenantZonesReadRequest) SetInstanceId(v string) *DescribeTenantZonesReadRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTenantZonesReadRequest) SetTenantId(v string) *DescribeTenantZonesReadRequest {
	s.TenantId = &v
	return s
}

type DescribeTenantZonesReadResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 租户可用区信息。
	TenantZones []*DescribeTenantZonesReadResponseBodyTenantZones `json:"TenantZones,omitempty" xml:"TenantZones,omitempty" type:"Repeated"`
}

func (s DescribeTenantZonesReadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantZonesReadResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTenantZonesReadResponseBody) SetRequestId(v string) *DescribeTenantZonesReadResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTenantZonesReadResponseBody) SetTenantZones(v []*DescribeTenantZonesReadResponseBodyTenantZones) *DescribeTenantZonesReadResponseBody {
	s.TenantZones = v
	return s
}

type DescribeTenantZonesReadResponseBodyTenantZones struct {
	// 是否可选为主库。
	IsElectable *bool `json:"IsElectable,omitempty" xml:"IsElectable,omitempty"`
	// 是否为主可用区。
	IsPrimary *bool `json:"IsPrimary,omitempty" xml:"IsPrimary,omitempty"`
	// 是否已经创建只读连接。
	IsReadOnlyAddressMaster *bool `json:"IsReadOnlyAddressMaster,omitempty" xml:"IsReadOnlyAddressMaster,omitempty"`
	// 是否可以设置为只读。
	IsReadable *string `json:"IsReadable,omitempty" xml:"IsReadable,omitempty"`
	// 可用区ID。
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s DescribeTenantZonesReadResponseBodyTenantZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantZonesReadResponseBodyTenantZones) GoString() string {
	return s.String()
}

func (s *DescribeTenantZonesReadResponseBodyTenantZones) SetIsElectable(v bool) *DescribeTenantZonesReadResponseBodyTenantZones {
	s.IsElectable = &v
	return s
}

func (s *DescribeTenantZonesReadResponseBodyTenantZones) SetIsPrimary(v bool) *DescribeTenantZonesReadResponseBodyTenantZones {
	s.IsPrimary = &v
	return s
}

func (s *DescribeTenantZonesReadResponseBodyTenantZones) SetIsReadOnlyAddressMaster(v bool) *DescribeTenantZonesReadResponseBodyTenantZones {
	s.IsReadOnlyAddressMaster = &v
	return s
}

func (s *DescribeTenantZonesReadResponseBodyTenantZones) SetIsReadable(v string) *DescribeTenantZonesReadResponseBodyTenantZones {
	s.IsReadable = &v
	return s
}

func (s *DescribeTenantZonesReadResponseBodyTenantZones) SetZone(v string) *DescribeTenantZonesReadResponseBodyTenantZones {
	s.Zone = &v
	return s
}

type DescribeTenantZonesReadResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTenantZonesReadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTenantZonesReadResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantZonesReadResponse) GoString() string {
	return s.String()
}

func (s *DescribeTenantZonesReadResponse) SetHeaders(v map[string]*string) *DescribeTenantZonesReadResponse {
	s.Headers = v
	return s
}

func (s *DescribeTenantZonesReadResponse) SetBody(v *DescribeTenantZonesReadResponseBody) *DescribeTenantZonesReadResponse {
	s.Body = v
	return s
}

type DescribeTenantsRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 分页查询的页码。 起始值：1 默认值：1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时设置的每页行数。 最大值：100 默认值：10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 查询列表的删选关键字。
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 租户名称。 长度为2~20 个字符，支持英文字母、数字和下划线，区分大小写，必须以字母或下划线开头。 不可设置为 sys。
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s DescribeTenantsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTenantsRequest) SetInstanceId(v string) *DescribeTenantsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTenantsRequest) SetPageNumber(v int32) *DescribeTenantsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTenantsRequest) SetPageSize(v int32) *DescribeTenantsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTenantsRequest) SetSearchKey(v string) *DescribeTenantsRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeTenantsRequest) SetTenantId(v string) *DescribeTenantsRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeTenantsRequest) SetTenantName(v string) *DescribeTenantsRequest {
	s.TenantName = &v
	return s
}

type DescribeTenantsResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 租户的列表信息。
	Tenants []*DescribeTenantsResponseBodyTenants `json:"Tenants,omitempty" xml:"Tenants,omitempty" type:"Repeated"`
	// 查询到的指定Oceanbase集群的租户总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTenantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTenantsResponseBody) SetRequestId(v string) *DescribeTenantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTenantsResponseBody) SetTenants(v []*DescribeTenantsResponseBodyTenants) *DescribeTenantsResponseBody {
	s.Tenants = v
	return s
}

func (s *DescribeTenantsResponseBody) SetTotalCount(v int32) *DescribeTenantsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTenantsResponseBodyTenants struct {
	// 租户总的CPU核数，单位：Core。
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// 租户的创建时间。
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 租户的数据副本模式。 对于高可用版本：三地域为N-N-N，两可用区和单可用区均为N-N； 对于基础版为N。 其中，N为单可用区内的节点数。
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// 租户的部署类型。 multiple：多机房 single：单机房 dual：双机房
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// 租户描述信息。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 租户总的内存大小，单位：GB。
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// 租户的主可用区。
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// 租户的状态。 PENDING_CREATE: 创建中 RESTORE: 恢复中 ONLINE: 运行中 SPEC_MODIFYING: 规格修改中 ALLOCATING_INTERNET_ADDRESS: 公网地址分配中 PENDING_OFFLINE_INTERNET_ADDRESS: 公网地址关闭中 PRIMARY_ZONE_MODIFYING: 切主可用区中 PARAMETER_MODIFYING: 参数修改中 WHITE_LIST_MODIFING: 白名单修改中
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 租户模式。 当前支持： Oracle：Oracle模式 MySQL: MySQL模式
	TenantMode *string `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
	// 租户名称。
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// 租户每个Unit的CPU核数，单位：Core。
	UnitCpu *int32 `json:"UnitCpu,omitempty" xml:"UnitCpu,omitempty"`
	// 租户每个Unit的内存大小，单位：GB。
	UnitMem *int32 `json:"UnitMem,omitempty" xml:"UnitMem,omitempty"`
	// 租户的unit个数。
	UnitNum *int32 `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
	// 租户已使用磁盘量
	UsedDiskSize *float64 `json:"UsedDiskSize,omitempty" xml:"UsedDiskSize,omitempty"`
	// 专有网络（VPC） ID。 如果没有合适的 VPC，请根据页面提示创建一个 VPC，详情参见 什么是专有网络
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeTenantsResponseBodyTenants) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantsResponseBodyTenants) GoString() string {
	return s.String()
}

func (s *DescribeTenantsResponseBodyTenants) SetCpu(v int32) *DescribeTenantsResponseBodyTenants {
	s.Cpu = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetCreateTime(v string) *DescribeTenantsResponseBodyTenants {
	s.CreateTime = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetDeployMode(v string) *DescribeTenantsResponseBodyTenants {
	s.DeployMode = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetDeployType(v string) *DescribeTenantsResponseBodyTenants {
	s.DeployType = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetDescription(v string) *DescribeTenantsResponseBodyTenants {
	s.Description = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetMem(v int32) *DescribeTenantsResponseBodyTenants {
	s.Mem = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetPrimaryZone(v string) *DescribeTenantsResponseBodyTenants {
	s.PrimaryZone = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetStatus(v string) *DescribeTenantsResponseBodyTenants {
	s.Status = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetTenantId(v string) *DescribeTenantsResponseBodyTenants {
	s.TenantId = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetTenantMode(v string) *DescribeTenantsResponseBodyTenants {
	s.TenantMode = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetTenantName(v string) *DescribeTenantsResponseBodyTenants {
	s.TenantName = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetUnitCpu(v int32) *DescribeTenantsResponseBodyTenants {
	s.UnitCpu = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetUnitMem(v int32) *DescribeTenantsResponseBodyTenants {
	s.UnitMem = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetUnitNum(v int32) *DescribeTenantsResponseBodyTenants {
	s.UnitNum = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetUsedDiskSize(v float64) *DescribeTenantsResponseBodyTenants {
	s.UsedDiskSize = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetVpcId(v string) *DescribeTenantsResponseBodyTenants {
	s.VpcId = &v
	return s
}

type DescribeTenantsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTenantsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTenantsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTenantsResponse) SetHeaders(v map[string]*string) *DescribeTenantsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTenantsResponse) SetBody(v *DescribeTenantsResponseBody) *DescribeTenantsResponse {
	s.Body = v
	return s
}

type DescribeTimeZonesResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 租户允许的时区信息。
	TimeZones *DescribeTimeZonesResponseBodyTimeZones `json:"TimeZones,omitempty" xml:"TimeZones,omitempty" type:"Struct"`
}

func (s DescribeTimeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTimeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTimeZonesResponseBody) SetRequestId(v string) *DescribeTimeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTimeZonesResponseBody) SetTimeZones(v *DescribeTimeZonesResponseBodyTimeZones) *DescribeTimeZonesResponseBody {
	s.TimeZones = v
	return s
}

type DescribeTimeZonesResponseBodyTimeZones struct {
	// 默认时区。
	Default *string `json:"Default,omitempty" xml:"Default,omitempty"`
	// 时区列表信息。
	List []*DescribeTimeZonesResponseBodyTimeZonesList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeTimeZonesResponseBodyTimeZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeTimeZonesResponseBodyTimeZones) GoString() string {
	return s.String()
}

func (s *DescribeTimeZonesResponseBodyTimeZones) SetDefault(v string) *DescribeTimeZonesResponseBodyTimeZones {
	s.Default = &v
	return s
}

func (s *DescribeTimeZonesResponseBodyTimeZones) SetList(v []*DescribeTimeZonesResponseBodyTimeZonesList) *DescribeTimeZonesResponseBodyTimeZones {
	s.List = v
	return s
}

type DescribeTimeZonesResponseBodyTimeZonesList struct {
	// 时区描述。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 时区名称。
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s DescribeTimeZonesResponseBodyTimeZonesList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTimeZonesResponseBodyTimeZonesList) GoString() string {
	return s.String()
}

func (s *DescribeTimeZonesResponseBodyTimeZonesList) SetDescription(v string) *DescribeTimeZonesResponseBodyTimeZonesList {
	s.Description = &v
	return s
}

func (s *DescribeTimeZonesResponseBodyTimeZonesList) SetTimeZone(v string) *DescribeTimeZonesResponseBodyTimeZonesList {
	s.TimeZone = &v
	return s
}

type DescribeTimeZonesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTimeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTimeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTimeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTimeZonesResponse) SetHeaders(v map[string]*string) *DescribeTimeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTimeZonesResponse) SetBody(v *DescribeTimeZonesResponseBody) *DescribeTimeZonesResponse {
	s.Body = v
	return s
}

type DescribeTopSQLListRequest struct {
	// 数据库名称
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 过滤条件
	FilterCondition map[string]interface{} `json:"FilterCondition,omitempty" xml:"FilterCondition,omitempty"`
	// 节点ip
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// 分页查询的页码。 起始值：1 默认值：1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时设置的每页行数。 最大值：100 默认值：10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// SQLID
	SQLId *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// 关键字查询
	SearchKeyWord *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	// 参数查询
	SearchParameter *string `json:"SearchParameter,omitempty" xml:"SearchParameter,omitempty"`
	// 查询规则
	SearchRule *string `json:"SearchRule,omitempty" xml:"SearchRule,omitempty"`
	// 查询值
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	// 排序列
	SortColumn *string `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	// 排序规则
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// 起始时间。
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 租户ID
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeTopSQLListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopSQLListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopSQLListRequest) SetDbName(v string) *DescribeTopSQLListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetEndTime(v string) *DescribeTopSQLListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetFilterCondition(v map[string]interface{}) *DescribeTopSQLListRequest {
	s.FilterCondition = v
	return s
}

func (s *DescribeTopSQLListRequest) SetNodeIp(v string) *DescribeTopSQLListRequest {
	s.NodeIp = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetPageNumber(v int32) *DescribeTopSQLListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetPageSize(v int32) *DescribeTopSQLListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetSQLId(v string) *DescribeTopSQLListRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetSearchKeyWord(v string) *DescribeTopSQLListRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetSearchParameter(v string) *DescribeTopSQLListRequest {
	s.SearchParameter = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetSearchRule(v string) *DescribeTopSQLListRequest {
	s.SearchRule = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetSearchValue(v string) *DescribeTopSQLListRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetSortColumn(v string) *DescribeTopSQLListRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetSortOrder(v string) *DescribeTopSQLListRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetStartTime(v string) *DescribeTopSQLListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetTenantId(v string) *DescribeTopSQLListRequest {
	s.TenantId = &v
	return s
}

type DescribeTopSQLListShrinkRequest struct {
	// 数据库名称
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 过滤条件
	FilterConditionShrink *string `json:"FilterCondition,omitempty" xml:"FilterCondition,omitempty"`
	// 节点ip
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// 分页查询的页码。 起始值：1 默认值：1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时设置的每页行数。 最大值：100 默认值：10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// SQLID
	SQLId *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// 关键字查询
	SearchKeyWord *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	// 参数查询
	SearchParameter *string `json:"SearchParameter,omitempty" xml:"SearchParameter,omitempty"`
	// 查询规则
	SearchRule *string `json:"SearchRule,omitempty" xml:"SearchRule,omitempty"`
	// 查询值
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	// 排序列
	SortColumn *string `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	// 排序规则
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// 起始时间。
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 租户ID
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeTopSQLListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopSQLListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopSQLListShrinkRequest) SetDbName(v string) *DescribeTopSQLListShrinkRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetEndTime(v string) *DescribeTopSQLListShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetFilterConditionShrink(v string) *DescribeTopSQLListShrinkRequest {
	s.FilterConditionShrink = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetNodeIp(v string) *DescribeTopSQLListShrinkRequest {
	s.NodeIp = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetPageNumber(v int32) *DescribeTopSQLListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetPageSize(v int32) *DescribeTopSQLListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetSQLId(v string) *DescribeTopSQLListShrinkRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetSearchKeyWord(v string) *DescribeTopSQLListShrinkRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetSearchParameter(v string) *DescribeTopSQLListShrinkRequest {
	s.SearchParameter = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetSearchRule(v string) *DescribeTopSQLListShrinkRequest {
	s.SearchRule = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetSearchValue(v string) *DescribeTopSQLListShrinkRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetSortColumn(v string) *DescribeTopSQLListShrinkRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetSortOrder(v string) *DescribeTopSQLListShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetStartTime(v string) *DescribeTopSQLListShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetTenantId(v string) *DescribeTopSQLListShrinkRequest {
	s.TenantId = &v
	return s
}

type DescribeTopSQLListResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// TOPSQL列表
	TopSQLList []*DescribeTopSQLListResponseBodyTopSQLList `json:"TopSQLList,omitempty" xml:"TopSQLList,omitempty" type:"Repeated"`
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTopSQLListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopSQLListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTopSQLListResponseBody) SetRequestId(v string) *DescribeTopSQLListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTopSQLListResponseBody) SetTopSQLList(v []*DescribeTopSQLListResponseBodyTopSQLList) *DescribeTopSQLListResponseBody {
	s.TopSQLList = v
	return s
}

func (s *DescribeTopSQLListResponseBody) SetTotalCount(v int64) *DescribeTopSQLListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTopSQLListResponseBodyTopSQLList struct {
	// 影响行数
	AffectedRows *int64 `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
	// 客户端等待
	AppWaitTime *float32 `json:"AppWaitTime,omitempty" xml:"AppWaitTime,omitempty"`
	// BlockCache命中次数
	BlockCacheHit *int64 `json:"BlockCacheHit,omitempty" xml:"BlockCacheHit,omitempty"`
	// BlockIndexCache命中次数
	BlockIndexCacheHit *int64 `json:"BlockIndexCacheHit,omitempty" xml:"BlockIndexCacheHit,omitempty"`
	// BloomFilterCache命中次数
	BloomFilterCacheHit *int64 `json:"BloomFilterCacheHit,omitempty" xml:"BloomFilterCacheHit,omitempty"`
	// 客户端IP
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// 并发等待
	ConcurrencyWaitTime *float32 `json:"ConcurrencyWaitTime,omitempty" xml:"ConcurrencyWaitTime,omitempty"`
	// 平均CPU时间
	CpuTime *float32 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	// 数据库名
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// 解码等待
	DecodeTime *float32 `json:"DecodeTime,omitempty" xml:"DecodeTime,omitempty"`
	// 物理读
	DiskRead *int64 `json:"DiskRead,omitempty" xml:"DiskRead,omitempty"`
	// 平均响应时间
	ElapsedTime *float32 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// 等待事件
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// 每秒次数
	ExecPerSecond *float32 `json:"ExecPerSecond,omitempty" xml:"ExecPerSecond,omitempty"`
	// 内部执行时间
	ExecuteTime *float32 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// 执行次数
	Executions *int64 `json:"Executions,omitempty" xml:"Executions,omitempty"`
	// 失败次数
	FailTimes *int64 `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	// 硬解析时间
	GetPlanTime *float32 `json:"GetPlanTime,omitempty" xml:"GetPlanTime,omitempty"`
	// IO等待
	IOWaitTime *float32 `json:"IOWaitTime,omitempty" xml:"IOWaitTime,omitempty"`
	// 返回数据序号
	Key *int64 `json:"Key,omitempty" xml:"Key,omitempty"`
	// 逻辑读
	LogicalRead *int64 `json:"LogicalRead,omitempty" xml:"LogicalRead,omitempty"`
	// 最大CPU时间
	MaxCpuTime *float32 `json:"MaxCpuTime,omitempty" xml:"MaxCpuTime,omitempty"`
	// 最大响应时间
	MaxElapsedTime *float32 `json:"MaxElapsedTime,omitempty" xml:"MaxElapsedTime,omitempty"`
	// 读内存行数
	MemstoreReadRowCount *int64 `json:"MemstoreReadRowCount,omitempty" xml:"MemstoreReadRowCount,omitempty"`
	// 未命中计划次数
	MissPlans *int64 `json:"MissPlans,omitempty" xml:"MissPlans,omitempty"`
	// 网络等待
	NetWaitTime *float32 `json:"NetWaitTime,omitempty" xml:"NetWaitTime,omitempty"`
	// 节点IP
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// 队列等待
	QueueTime *float32 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	// RPC次数
	RPCCount *int64 `json:"RPCCount,omitempty" xml:"RPCCount,omitempty"`
	// 远程计划数
	RemotePlans *int64 `json:"RemotePlans,omitempty" xml:"RemotePlans,omitempty"`
	// 重试次数
	RetryCount *int64 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// 返回行数
	ReturnRows *int64 `json:"ReturnRows,omitempty" xml:"ReturnRows,omitempty"`
	// RowCache命中次数
	RowCacheHit *int64 `json:"RowCacheHit,omitempty" xml:"RowCacheHit,omitempty"`
	// SQLID
	SQLId *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// sql文本
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// sql类型
	SQLType *int64 `json:"SQLType,omitempty" xml:"SQLType,omitempty"`
	// 调度时间
	ScheduleTime *float32 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// 读磁盘行数
	SsstoreReadRowCount *int64 `json:"SsstoreReadRowCount,omitempty" xml:"SsstoreReadRowCount,omitempty"`
	// 内部等待
	TotalWaitTime *float32 `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	// 用户名
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeTopSQLListResponseBodyTopSQLList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopSQLListResponseBodyTopSQLList) GoString() string {
	return s.String()
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetAffectedRows(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.AffectedRows = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetAppWaitTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.AppWaitTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetBlockCacheHit(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.BlockCacheHit = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetBlockIndexCacheHit(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.BlockIndexCacheHit = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetBloomFilterCacheHit(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.BloomFilterCacheHit = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetClientIp(v string) *DescribeTopSQLListResponseBodyTopSQLList {
	s.ClientIp = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetConcurrencyWaitTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.ConcurrencyWaitTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetCpuTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.CpuTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetDbName(v string) *DescribeTopSQLListResponseBodyTopSQLList {
	s.DbName = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetDecodeTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.DecodeTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetDiskRead(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.DiskRead = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetElapsedTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetEvent(v string) *DescribeTopSQLListResponseBodyTopSQLList {
	s.Event = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetExecPerSecond(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.ExecPerSecond = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetExecuteTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetExecutions(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.Executions = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetFailTimes(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.FailTimes = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetGetPlanTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.GetPlanTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetIOWaitTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.IOWaitTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetKey(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.Key = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetLogicalRead(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.LogicalRead = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetMaxCpuTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.MaxCpuTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetMaxElapsedTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.MaxElapsedTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetMemstoreReadRowCount(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.MemstoreReadRowCount = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetMissPlans(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.MissPlans = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetNetWaitTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.NetWaitTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetNodeIp(v string) *DescribeTopSQLListResponseBodyTopSQLList {
	s.NodeIp = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetQueueTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.QueueTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetRPCCount(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.RPCCount = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetRemotePlans(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.RemotePlans = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetRetryCount(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.RetryCount = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetReturnRows(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.ReturnRows = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetRowCacheHit(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.RowCacheHit = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetSQLId(v string) *DescribeTopSQLListResponseBodyTopSQLList {
	s.SQLId = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetSQLText(v string) *DescribeTopSQLListResponseBodyTopSQLList {
	s.SQLText = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetSQLType(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.SQLType = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetScheduleTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.ScheduleTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetSsstoreReadRowCount(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.SsstoreReadRowCount = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetTotalWaitTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.TotalWaitTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetUserName(v string) *DescribeTopSQLListResponseBodyTopSQLList {
	s.UserName = &v
	return s
}

type DescribeTopSQLListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTopSQLListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTopSQLListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopSQLListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTopSQLListResponse) SetHeaders(v map[string]*string) *DescribeTopSQLListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTopSQLListResponse) SetBody(v *DescribeTopSQLListResponseBody) *DescribeTopSQLListResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	// 集群的部署类型。- multiple：多机房 - single：单机房 - dual：双机房
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// Oceanbase集群的系列 - normal（默认）：高可用版本 - basic：基础版本
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetDeployType(v string) *DescribeZonesRequest {
	s.DeployType = &v
	return s
}

func (s *DescribeZonesRequest) SetSeries(v string) *DescribeZonesRequest {
	s.Series = &v
	return s
}

type DescribeZonesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 可用区列表信息
	Zones []*DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	// 部署模式
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// Oceanbase集群的系列  - NORMAL（默认）：高可用版本  - BASIC：基础版本
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// 在各可用区的库存状态。其取值为：  WithStock：库存充足。 ClosedWithStock：库存供应保障能力低，建议选用WithStock状态的产品规格。 WithoutStock：库存售罄，将会补充资源，建议选用WithStock状态的产品规格
	StockStatus *string `json:"StockStatus,omitempty" xml:"StockStatus,omitempty"`
	// 可用区ID列表。对于多可用区集群，其可用区名称以逗号连接
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// 可用区名称
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetDeployType(v string) *DescribeZonesResponseBodyZones {
	s.DeployType = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetSeries(v string) *DescribeZonesResponseBodyZones {
	s.Series = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetStockStatus(v string) *DescribeZonesResponseBodyZones {
	s.StockStatus = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetZoneId(v string) *DescribeZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetZoneName(v string) *DescribeZonesResponseBodyZones {
	s.ZoneName = &v
	return s
}

type DescribeZonesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type ModifyDatabaseDescriptionRequest struct {
	// 数据库名称。 不能使用某些预留关键字，如 test、mysql。
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// 数据库描述信息。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ModifyDatabaseDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseDescriptionRequest) SetDatabaseName(v string) *ModifyDatabaseDescriptionRequest {
	s.DatabaseName = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetDescription(v string) *ModifyDatabaseDescriptionRequest {
	s.Description = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetInstanceId(v string) *ModifyDatabaseDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetTenantId(v string) *ModifyDatabaseDescriptionRequest {
	s.TenantId = &v
	return s
}

type ModifyDatabaseDescriptionResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDatabaseDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseDescriptionResponseBody) SetRequestId(v string) *ModifyDatabaseDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDatabaseDescriptionResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDatabaseDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDatabaseDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDatabaseDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseDescriptionResponse) SetBody(v *ModifyDatabaseDescriptionResponseBody) *ModifyDatabaseDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDatabaseUserRolesRequest struct {
	// 数据库名称。 不能使用某些预留关键字，如 test、mysql。
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 用户名及具备的角色列表。
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s ModifyDatabaseUserRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseUserRolesRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseUserRolesRequest) SetDatabaseName(v string) *ModifyDatabaseUserRolesRequest {
	s.DatabaseName = &v
	return s
}

func (s *ModifyDatabaseUserRolesRequest) SetInstanceId(v string) *ModifyDatabaseUserRolesRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDatabaseUserRolesRequest) SetTenantId(v string) *ModifyDatabaseUserRolesRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyDatabaseUserRolesRequest) SetUsers(v string) *ModifyDatabaseUserRolesRequest {
	s.Users = &v
	return s
}

type ModifyDatabaseUserRolesResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 账号信息。
	TenantUser *ModifyDatabaseUserRolesResponseBodyTenantUser `json:"TenantUser,omitempty" xml:"TenantUser,omitempty" type:"Struct"`
}

func (s ModifyDatabaseUserRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseUserRolesResponseBody) SetRequestId(v string) *ModifyDatabaseUserRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDatabaseUserRolesResponseBody) SetTenantUser(v *ModifyDatabaseUserRolesResponseBodyTenantUser) *ModifyDatabaseUserRolesResponseBody {
	s.TenantUser = v
	return s
}

type ModifyDatabaseUserRolesResponseBodyTenantUser struct {
	// 数据库名称。
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 对数据库有赋权的账户信息
	Users []*ModifyDatabaseUserRolesResponseBodyTenantUserUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ModifyDatabaseUserRolesResponseBodyTenantUser) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseUserRolesResponseBodyTenantUser) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseUserRolesResponseBodyTenantUser) SetDatabaseName(v string) *ModifyDatabaseUserRolesResponseBodyTenantUser {
	s.DatabaseName = &v
	return s
}

func (s *ModifyDatabaseUserRolesResponseBodyTenantUser) SetTenantId(v string) *ModifyDatabaseUserRolesResponseBodyTenantUser {
	s.TenantId = &v
	return s
}

func (s *ModifyDatabaseUserRolesResponseBodyTenantUser) SetUsers(v []*ModifyDatabaseUserRolesResponseBodyTenantUserUsers) *ModifyDatabaseUserRolesResponseBodyTenantUser {
	s.Users = v
	return s
}

type ModifyDatabaseUserRolesResponseBodyTenantUserUsers struct {
	// 账号赋予该库的角色权限。 对于MySQL模式，角色为数据库（Database）级别，其有以下几类： - ReadWrite：读写权限，包括ALL PRIVILEGES； - ReadOnly：只读权限，包括SELECT - DDL: DDL权限，包括CREATE,DROP,ALTER,SHOW VIEW,CREATE VIEW - DML: DML权限，包括SELECT,INSERT,UPDATE,DELETE,SHOW VIEW。
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 账号名称。
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ModifyDatabaseUserRolesResponseBodyTenantUserUsers) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseUserRolesResponseBodyTenantUserUsers) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseUserRolesResponseBodyTenantUserUsers) SetRole(v string) *ModifyDatabaseUserRolesResponseBodyTenantUserUsers {
	s.Role = &v
	return s
}

func (s *ModifyDatabaseUserRolesResponseBodyTenantUserUsers) SetUserName(v string) *ModifyDatabaseUserRolesResponseBodyTenantUserUsers {
	s.UserName = &v
	return s
}

type ModifyDatabaseUserRolesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDatabaseUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDatabaseUserRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseUserRolesResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseUserRolesResponse) SetHeaders(v map[string]*string) *ModifyDatabaseUserRolesResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseUserRolesResponse) SetBody(v *ModifyDatabaseUserRolesResponseBody) *ModifyDatabaseUserRolesResponse {
	s.Body = v
	return s
}

type ModifyInstanceNameRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Oceanbase集群名称。 长度为1~20个英文或中文字符。如果没有指定该参数，默认值为集群的InstanceId。
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ModifyInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameRequest) SetInstanceId(v string) *ModifyInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetInstanceName(v string) *ModifyInstanceNameRequest {
	s.InstanceName = &v
	return s
}

type ModifyInstanceNameResponseBody struct {
	// Oceanbase集群名称。
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameResponseBody) SetInstanceName(v string) *ModifyInstanceNameResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceNameResponseBody) SetRequestId(v string) *ModifyInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceNameResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameResponse) SetHeaders(v map[string]*string) *ModifyInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceNameResponse) SetBody(v *ModifyInstanceNameResponseBody) *ModifyInstanceNameResponse {
	s.Body = v
	return s
}

type ModifyParametersRequest struct {
	// 参数类型。 当前支持集群（CLUSTER)和租户（TENANT）
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// 参数类型的资源标识。 如果为集群参数可以不填，若为租户的参数，则传入租户的TenantId。
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 参数信息。 其为一个Json格式的数组，数组中每个对象包括两个元素：参数名称（Name）和参数值（Value）。 注意：集群和租户可修改的参数名称和参数值的范围不同，详见DescribeParameters。
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s ModifyParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyParametersRequest) SetDimension(v string) *ModifyParametersRequest {
	s.Dimension = &v
	return s
}

func (s *ModifyParametersRequest) SetDimensionValue(v string) *ModifyParametersRequest {
	s.DimensionValue = &v
	return s
}

func (s *ModifyParametersRequest) SetInstanceId(v string) *ModifyParametersRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyParametersRequest) SetParameters(v string) *ModifyParametersRequest {
	s.Parameters = &v
	return s
}

type ModifyParametersResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 修改结果信息。
	Results *ModifyParametersResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
}

func (s ModifyParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyParametersResponseBody) SetRequestId(v string) *ModifyParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyParametersResponseBody) SetResults(v *ModifyParametersResponseBodyResults) *ModifyParametersResponseBody {
	s.Results = v
	return s
}

type ModifyParametersResponseBodyResults struct {
	// 修改失败信息。
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 是否修改成功。 - true：修改成功 - false：修改失败
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyParametersResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ModifyParametersResponseBodyResults) SetMessage(v string) *ModifyParametersResponseBodyResults {
	s.Message = &v
	return s
}

func (s *ModifyParametersResponseBodyResults) SetSuccess(v bool) *ModifyParametersResponseBodyResults {
	s.Success = &v
	return s
}

type ModifyParametersResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersResponse) GoString() string {
	return s.String()
}

func (s *ModifyParametersResponse) SetHeaders(v map[string]*string) *ModifyParametersResponse {
	s.Headers = v
	return s
}

func (s *ModifyParametersResponse) SetBody(v *ModifyParametersResponseBody) *ModifyParametersResponse {
	s.Body = v
	return s
}

type ModifySecurityIpsRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// IP安全白名单组的组名。 由小写英文字符开头，由小写英文字符或者数字结尾，只能包含小写英文字符，数字和下划线，长度在 2-32 个字符之间。
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	// IP安全白名单列表。 其为一个Json格式的数组，数组中每个对象为一个IP字符串或者IP段。最多可设置 40 个。
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) SetInstanceId(v string) *ModifySecurityIpsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIpGroupName(v string) *ModifySecurityIpsRequest {
	s.SecurityIpGroupName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIps(v string) *ModifySecurityIpsRequest {
	s.SecurityIps = &v
	return s
}

type ModifySecurityIpsResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// IP白名单分组信息。
	SecurityIpGroup *ModifySecurityIpsResponseBodySecurityIpGroup `json:"SecurityIpGroup,omitempty" xml:"SecurityIpGroup,omitempty" type:"Struct"`
}

func (s ModifySecurityIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponseBody) SetRequestId(v string) *ModifySecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityIpsResponseBody) SetSecurityIpGroup(v *ModifySecurityIpsResponseBodySecurityIpGroup) *ModifySecurityIpsResponseBody {
	s.SecurityIpGroup = v
	return s
}

type ModifySecurityIpsResponseBodySecurityIpGroup struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 安全组名称。
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	// IP安全白名单列表。其为一个Json格式的数组，数组中每个对象为一个IP字符串或者IP段。
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s ModifySecurityIpsResponseBodySecurityIpGroup) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsResponseBodySecurityIpGroup) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponseBodySecurityIpGroup) SetInstanceId(v string) *ModifySecurityIpsResponseBodySecurityIpGroup {
	s.InstanceId = &v
	return s
}

func (s *ModifySecurityIpsResponseBodySecurityIpGroup) SetSecurityIpGroupName(v string) *ModifySecurityIpsResponseBodySecurityIpGroup {
	s.SecurityIpGroupName = &v
	return s
}

func (s *ModifySecurityIpsResponseBodySecurityIpGroup) SetSecurityIps(v string) *ModifySecurityIpsResponseBodySecurityIpGroup {
	s.SecurityIps = &v
	return s
}

type ModifySecurityIpsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySecurityIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponse) SetHeaders(v map[string]*string) *ModifySecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityIpsResponse) SetBody(v *ModifySecurityIpsResponseBody) *ModifySecurityIpsResponse {
	s.Body = v
	return s
}

type ModifyTenantPrimaryZoneRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户的主可用区。 其为集群部署可用区中的一个。
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ModifyTenantPrimaryZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantPrimaryZoneRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantPrimaryZoneRequest) SetInstanceId(v string) *ModifyTenantPrimaryZoneRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTenantPrimaryZoneRequest) SetPrimaryZone(v string) *ModifyTenantPrimaryZoneRequest {
	s.PrimaryZone = &v
	return s
}

func (s *ModifyTenantPrimaryZoneRequest) SetTenantId(v string) *ModifyTenantPrimaryZoneRequest {
	s.TenantId = &v
	return s
}

type ModifyTenantPrimaryZoneResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTenantPrimaryZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantPrimaryZoneResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantPrimaryZoneResponseBody) SetRequestId(v string) *ModifyTenantPrimaryZoneResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTenantPrimaryZoneResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyTenantPrimaryZoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTenantPrimaryZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantPrimaryZoneResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantPrimaryZoneResponse) SetHeaders(v map[string]*string) *ModifyTenantPrimaryZoneResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantPrimaryZoneResponse) SetBody(v *ModifyTenantPrimaryZoneResponseBody) *ModifyTenantPrimaryZoneResponse {
	s.Body = v
	return s
}

type ModifyTenantResourceRequest struct {
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户内存大小，单位GB。
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ModifyTenantResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantResourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantResourceRequest) SetCpu(v int32) *ModifyTenantResourceRequest {
	s.Cpu = &v
	return s
}

func (s *ModifyTenantResourceRequest) SetInstanceId(v string) *ModifyTenantResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTenantResourceRequest) SetMemory(v int32) *ModifyTenantResourceRequest {
	s.Memory = &v
	return s
}

func (s *ModifyTenantResourceRequest) SetTenantId(v string) *ModifyTenantResourceRequest {
	s.TenantId = &v
	return s
}

type ModifyTenantResourceResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ModifyTenantResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantResourceResponseBody) SetRequestId(v string) *ModifyTenantResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTenantResourceResponseBody) SetTenantId(v string) *ModifyTenantResourceResponseBody {
	s.TenantId = &v
	return s
}

type ModifyTenantResourceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyTenantResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTenantResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantResourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantResourceResponse) SetHeaders(v map[string]*string) *ModifyTenantResourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantResourceResponse) SetBody(v *ModifyTenantResourceResponseBody) *ModifyTenantResourceResponse {
	s.Body = v
	return s
}

type ModifyTenantUserDescriptionRequest struct {
	// 数据库描述信息。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 数据库账号名称。 不能使用某些预留关键字，如 SYS、root等。
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ModifyTenantUserDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserDescriptionRequest) SetDescription(v string) *ModifyTenantUserDescriptionRequest {
	s.Description = &v
	return s
}

func (s *ModifyTenantUserDescriptionRequest) SetInstanceId(v string) *ModifyTenantUserDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTenantUserDescriptionRequest) SetTenantId(v string) *ModifyTenantUserDescriptionRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyTenantUserDescriptionRequest) SetUserName(v string) *ModifyTenantUserDescriptionRequest {
	s.UserName = &v
	return s
}

type ModifyTenantUserDescriptionResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTenantUserDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserDescriptionResponseBody) SetRequestId(v string) *ModifyTenantUserDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTenantUserDescriptionResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyTenantUserDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTenantUserDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserDescriptionResponse) SetHeaders(v map[string]*string) *ModifyTenantUserDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantUserDescriptionResponse) SetBody(v *ModifyTenantUserDescriptionResponseBody) *ModifyTenantUserDescriptionResponse {
	s.Body = v
	return s
}

type ModifyTenantUserPasswordRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 数据库账号名称。 不能使用某些预留关键字，如 SYS、root等。
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// 数据库账号访问密码。 必须包含大写英文字符、小写英文字符、数字、特殊字符占三种，长度为 10-32 位； 特殊字符为：!@#$%^&* ()_ +-=
	UserPassword *string `json:"UserPassword,omitempty" xml:"UserPassword,omitempty"`
}

func (s ModifyTenantUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserPasswordRequest) SetInstanceId(v string) *ModifyTenantUserPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTenantUserPasswordRequest) SetTenantId(v string) *ModifyTenantUserPasswordRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyTenantUserPasswordRequest) SetUserName(v string) *ModifyTenantUserPasswordRequest {
	s.UserName = &v
	return s
}

func (s *ModifyTenantUserPasswordRequest) SetUserPassword(v string) *ModifyTenantUserPasswordRequest {
	s.UserPassword = &v
	return s
}

type ModifyTenantUserPasswordResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTenantUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserPasswordResponseBody) SetRequestId(v string) *ModifyTenantUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTenantUserPasswordResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyTenantUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTenantUserPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserPasswordResponse) SetHeaders(v map[string]*string) *ModifyTenantUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantUserPasswordResponse) SetBody(v *ModifyTenantUserPasswordResponseBody) *ModifyTenantUserPasswordResponse {
	s.Body = v
	return s
}

type ModifyTenantUserRolesRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 用户权限修改类型。 可支持以下几种方式： update：全量更新权限，默认值； add：新增权限； delete：删除权限。 默认值：update。
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 数据库账号名称。 不能使用某些预留关键字，如 SYS、root等。
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// 数据库的账号角色信息。
	UserRole *string `json:"UserRole,omitempty" xml:"UserRole,omitempty"`
}

func (s ModifyTenantUserRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserRolesRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserRolesRequest) SetInstanceId(v string) *ModifyTenantUserRolesRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTenantUserRolesRequest) SetModifyType(v string) *ModifyTenantUserRolesRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyTenantUserRolesRequest) SetTenantId(v string) *ModifyTenantUserRolesRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyTenantUserRolesRequest) SetUserName(v string) *ModifyTenantUserRolesRequest {
	s.UserName = &v
	return s
}

func (s *ModifyTenantUserRolesRequest) SetUserRole(v string) *ModifyTenantUserRolesRequest {
	s.UserRole = &v
	return s
}

type ModifyTenantUserRolesResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 租户中的数据库账号信息。
	TenantUser *ModifyTenantUserRolesResponseBodyTenantUser `json:"TenantUser,omitempty" xml:"TenantUser,omitempty" type:"Struct"`
}

func (s ModifyTenantUserRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserRolesResponseBody) SetRequestId(v string) *ModifyTenantUserRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTenantUserRolesResponseBody) SetTenantUser(v *ModifyTenantUserRolesResponseBodyTenantUser) *ModifyTenantUserRolesResponseBody {
	s.TenantUser = v
	return s
}

type ModifyTenantUserRolesResponseBodyTenantUser struct {
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 账号名称。
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// 账号类型。 Admin：超级账户 Normal：普通账户
	UserRole []*ModifyTenantUserRolesResponseBodyTenantUserUserRole `json:"UserRole,omitempty" xml:"UserRole,omitempty" type:"Repeated"`
}

func (s ModifyTenantUserRolesResponseBodyTenantUser) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserRolesResponseBodyTenantUser) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserRolesResponseBodyTenantUser) SetTenantId(v string) *ModifyTenantUserRolesResponseBodyTenantUser {
	s.TenantId = &v
	return s
}

func (s *ModifyTenantUserRolesResponseBodyTenantUser) SetUserName(v string) *ModifyTenantUserRolesResponseBodyTenantUser {
	s.UserName = &v
	return s
}

func (s *ModifyTenantUserRolesResponseBodyTenantUser) SetUserRole(v []*ModifyTenantUserRolesResponseBodyTenantUserUserRole) *ModifyTenantUserRolesResponseBodyTenantUser {
	s.UserRole = v
	return s
}

type ModifyTenantUserRolesResponseBodyTenantUserUserRole struct {
	// 数据库(schema)名称
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// 是否授权成功。
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// 拥有的角色。
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 表的名称。
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s ModifyTenantUserRolesResponseBodyTenantUserUserRole) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserRolesResponseBodyTenantUserUserRole) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserRolesResponseBodyTenantUserUserRole) SetDatabase(v string) *ModifyTenantUserRolesResponseBodyTenantUserUserRole {
	s.Database = &v
	return s
}

func (s *ModifyTenantUserRolesResponseBodyTenantUserUserRole) SetIsSuccess(v bool) *ModifyTenantUserRolesResponseBodyTenantUserUserRole {
	s.IsSuccess = &v
	return s
}

func (s *ModifyTenantUserRolesResponseBodyTenantUserUserRole) SetRole(v string) *ModifyTenantUserRolesResponseBodyTenantUserUserRole {
	s.Role = &v
	return s
}

func (s *ModifyTenantUserRolesResponseBodyTenantUserUserRole) SetTable(v string) *ModifyTenantUserRolesResponseBodyTenantUserUserRole {
	s.Table = &v
	return s
}

type ModifyTenantUserRolesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyTenantUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTenantUserRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserRolesResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserRolesResponse) SetHeaders(v map[string]*string) *ModifyTenantUserRolesResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantUserRolesResponse) SetBody(v *ModifyTenantUserRolesResponseBody) *ModifyTenantUserRolesResponse {
	s.Body = v
	return s
}

type ModifyTenantUserStatusRequest struct {
	// Oceanbase集群ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 数据库账号名称。 不能使用某些预留关键字，如 SYS、root等。
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// 数据库账号的状态。 Locked：锁定 Normal：解锁
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s ModifyTenantUserStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserStatusRequest) SetInstanceId(v string) *ModifyTenantUserStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTenantUserStatusRequest) SetTenantId(v string) *ModifyTenantUserStatusRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyTenantUserStatusRequest) SetUserName(v string) *ModifyTenantUserStatusRequest {
	s.UserName = &v
	return s
}

func (s *ModifyTenantUserStatusRequest) SetUserStatus(v string) *ModifyTenantUserStatusRequest {
	s.UserStatus = &v
	return s
}

type ModifyTenantUserStatusResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 租户中的数据库账号信息
	TenantUser []*ModifyTenantUserStatusResponseBodyTenantUser `json:"TenantUser,omitempty" xml:"TenantUser,omitempty" type:"Repeated"`
}

func (s ModifyTenantUserStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserStatusResponseBody) SetRequestId(v string) *ModifyTenantUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTenantUserStatusResponseBody) SetTenantUser(v []*ModifyTenantUserStatusResponseBodyTenantUser) *ModifyTenantUserStatusResponseBody {
	s.TenantUser = v
	return s
}

type ModifyTenantUserStatusResponseBodyTenantUser struct {
	// 租户ID。
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 账号名称。
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// 数据库账号的状态。 - LOCKED：锁定 - ONLINE：解锁
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s ModifyTenantUserStatusResponseBodyTenantUser) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserStatusResponseBodyTenantUser) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserStatusResponseBodyTenantUser) SetTenantId(v string) *ModifyTenantUserStatusResponseBodyTenantUser {
	s.TenantId = &v
	return s
}

func (s *ModifyTenantUserStatusResponseBodyTenantUser) SetUserName(v string) *ModifyTenantUserStatusResponseBodyTenantUser {
	s.UserName = &v
	return s
}

func (s *ModifyTenantUserStatusResponseBodyTenantUser) SetUserStatus(v string) *ModifyTenantUserStatusResponseBodyTenantUser {
	s.UserStatus = &v
	return s
}

type ModifyTenantUserStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyTenantUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTenantUserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserStatusResponse) SetHeaders(v map[string]*string) *ModifyTenantUserStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantUserStatusResponse) SetBody(v *ModifyTenantUserStatusResponseBody) *ModifyTenantUserStatusResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("oceanbasepro"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateDatabaseWithOptions(request *CreateDatabaseRequest, runtime *util.RuntimeOptions) (_result *CreateDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Encoding)) {
		body["Encoding"] = request.Encoding
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDatabase"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDatabase(request *CreateDatabaseRequest) (_result *CreateDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CreateDatabaseWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		body["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.DiskSize)) {
		body["DiskSize"] = request.DiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceClass)) {
		body["InstanceClass"] = request.InstanceClass
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		body["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Series)) {
		body["Series"] = request.Series
	}

	if !tea.BoolValue(util.IsUnset(request.Zones)) {
		body["Zones"] = request.Zones
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateSecurityIpGroupWithOptions(request *CreateSecurityIpGroupRequest, runtime *util.RuntimeOptions) (_result *CreateSecurityIpGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIpGroupName)) {
		body["SecurityIpGroupName"] = request.SecurityIpGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIps)) {
		body["SecurityIps"] = request.SecurityIps
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSecurityIpGroup"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSecurityIpGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSecurityIpGroup(request *CreateSecurityIpGroupRequest) (_result *CreateSecurityIpGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSecurityIpGroupResponse{}
	_body, _err := client.CreateSecurityIpGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTenantWithOptions(request *CreateTenantRequest, runtime *util.RuntimeOptions) (_result *CreateTenantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Charset)) {
		body["Charset"] = request.Charset
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		body["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		body["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryZone)) {
		body["PrimaryZone"] = request.PrimaryZone
	}

	if !tea.BoolValue(util.IsUnset(request.TenantMode)) {
		body["TenantMode"] = request.TenantMode
	}

	if !tea.BoolValue(util.IsUnset(request.TenantName)) {
		body["TenantName"] = request.TenantName
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		body["TimeZone"] = request.TimeZone
	}

	if !tea.BoolValue(util.IsUnset(request.UnitNum)) {
		body["UnitNum"] = request.UnitNum
	}

	if !tea.BoolValue(util.IsUnset(request.UserVSwitchId)) {
		body["UserVSwitchId"] = request.UserVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.UserVpcId)) {
		body["UserVpcId"] = request.UserVpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTenant"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTenantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTenant(request *CreateTenantRequest) (_result *CreateTenantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTenantResponse{}
	_body, _err := client.CreateTenantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTenantReadOnlyConnectionWithOptions(request *CreateTenantReadOnlyConnectionRequest, runtime *util.RuntimeOptions) (_result *CreateTenantReadOnlyConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTenantReadOnlyConnection"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTenantReadOnlyConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTenantReadOnlyConnection(request *CreateTenantReadOnlyConnectionRequest) (_result *CreateTenantReadOnlyConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTenantReadOnlyConnectionResponse{}
	_body, _err := client.CreateTenantReadOnlyConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTenantUserWithOptions(request *CreateTenantUserRequest, runtime *util.RuntimeOptions) (_result *CreateTenantUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Roles)) {
		body["Roles"] = request.Roles
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserPassword)) {
		body["UserPassword"] = request.UserPassword
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		body["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTenantUser"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTenantUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTenantUser(request *CreateTenantUserRequest) (_result *CreateTenantUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTenantUserResponse{}
	_body, _err := client.CreateTenantUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDatabasesWithOptions(request *DeleteDatabasesRequest, runtime *util.RuntimeOptions) (_result *DeleteDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseNames)) {
		body["DatabaseNames"] = request.DatabaseNames
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDatabases"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDatabases(request *DeleteDatabasesRequest) (_result *DeleteDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDatabasesResponse{}
	_body, _err := client.DeleteDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstancesWithOptions(request *DeleteInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupRetainMode)) {
		body["BackupRetainMode"] = request.BackupRetainMode
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		body["InstanceIds"] = request.InstanceIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstances"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstances(request *DeleteInstancesRequest) (_result *DeleteInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstancesResponse{}
	_body, _err := client.DeleteInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSecurityIpGroupWithOptions(request *DeleteSecurityIpGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteSecurityIpGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIpGroupName)) {
		body["SecurityIpGroupName"] = request.SecurityIpGroupName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSecurityIpGroup"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSecurityIpGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecurityIpGroup(request *DeleteSecurityIpGroupRequest) (_result *DeleteSecurityIpGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecurityIpGroupResponse{}
	_body, _err := client.DeleteSecurityIpGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTenantUsersWithOptions(request *DeleteTenantUsersRequest, runtime *util.RuntimeOptions) (_result *DeleteTenantUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTenantUsers"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTenantUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTenantUsers(request *DeleteTenantUsersRequest) (_result *DeleteTenantUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTenantUsersResponse{}
	_body, _err := client.DeleteTenantUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTenantsWithOptions(request *DeleteTenantsRequest, runtime *util.RuntimeOptions) (_result *DeleteTenantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantIds)) {
		body["TenantIds"] = request.TenantIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTenants"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTenantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTenants(request *DeleteTenantsRequest) (_result *DeleteTenantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTenantsResponse{}
	_body, _err := client.DeleteTenantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAnomalySQLListWithOptions(tmpReq *DescribeAnomalySQLListRequest, runtime *util.RuntimeOptions) (_result *DescribeAnomalySQLListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeAnomalySQLListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FilterCondition)) {
		request.FilterConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterCondition, tea.String("FilterCondition"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FilterConditionShrink)) {
		body["FilterCondition"] = request.FilterConditionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIp)) {
		body["NodeIp"] = request.NodeIp
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKeyWord)) {
		body["SearchKeyWord"] = request.SearchKeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.SearchParameter)) {
		body["SearchParameter"] = request.SearchParameter
	}

	if !tea.BoolValue(util.IsUnset(request.SearchRule)) {
		body["SearchRule"] = request.SearchRule
	}

	if !tea.BoolValue(util.IsUnset(request.SearchValue)) {
		body["SearchValue"] = request.SearchValue
	}

	if !tea.BoolValue(util.IsUnset(request.SortColumn)) {
		body["SortColumn"] = request.SortColumn
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		body["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAnomalySQLList"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAnomalySQLListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAnomalySQLList(request *DescribeAnomalySQLListRequest) (_result *DescribeAnomalySQLListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAnomalySQLListResponse{}
	_body, _err := client.DescribeAnomalySQLListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableCpuResourceWithOptions(request *DescribeAvailableCpuResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableCpuResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyType)) {
		body["ModifyType"] = request.ModifyType
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableCpuResource"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAvailableCpuResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableCpuResource(request *DescribeAvailableCpuResourceRequest) (_result *DescribeAvailableCpuResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableCpuResourceResponse{}
	_body, _err := client.DescribeAvailableCpuResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableMemResourceWithOptions(request *DescribeAvailableMemResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableMemResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CpuNum)) {
		body["CpuNum"] = request.CpuNum
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UnitNum)) {
		body["UnitNum"] = request.UnitNum
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableMemResource"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAvailableMemResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableMemResource(request *DescribeAvailableMemResourceRequest) (_result *DescribeAvailableMemResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableMemResourceResponse{}
	_body, _err := client.DescribeAvailableMemResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCharsetWithOptions(request *DescribeCharsetRequest, runtime *util.RuntimeOptions) (_result *DescribeCharsetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantMode)) {
		body["TenantMode"] = request.TenantMode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCharset"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCharsetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCharset(request *DescribeCharsetRequest) (_result *DescribeCharsetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCharsetResponse{}
	_body, _err := client.DescribeCharsetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDatabasesWithOptions(request *DescribeDatabasesRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.WithTables)) {
		body["WithTables"] = request.WithTables
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDatabases"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDatabases(request *DescribeDatabasesRequest) (_result *DescribeDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabasesResponse{}
	_body, _err := client.DescribeDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstance"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceCreatableZoneWithOptions(request *DescribeInstanceCreatableZoneRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceCreatableZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceCreatableZone"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceCreatableZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceCreatableZone(request *DescribeInstanceCreatableZoneRequest) (_result *DescribeInstanceCreatableZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceCreatableZoneResponse{}
	_body, _err := client.DescribeInstanceCreatableZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceTenantModesWithOptions(request *DescribeInstanceTenantModesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceTenantModesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceTenantModes"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceTenantModesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceTenantModes(request *DescribeInstanceTenantModesRequest) (_result *DescribeInstanceTenantModesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceTenantModesResponse{}
	_body, _err := client.DescribeInstanceTenantModesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceTopologyWithOptions(request *DescribeInstanceTopologyRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceTopologyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceTopology"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceTopologyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceTopology(request *DescribeInstanceTopologyRequest) (_result *DescribeInstanceTopologyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceTopologyResponse{}
	_body, _err := client.DescribeInstanceTopologyWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstances"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeNodeMetricsWithOptions(request *DescribeNodeMetricsRequest, runtime *util.RuntimeOptions) (_result *DescribeNodeMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Metrics)) {
		body["Metrics"] = request.Metrics
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIdList)) {
		body["NodeIdList"] = request.NodeIdList
	}

	if !tea.BoolValue(util.IsUnset(request.NodeName)) {
		body["NodeName"] = request.NodeName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNodeMetrics"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNodeMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNodeMetrics(request *DescribeNodeMetricsRequest) (_result *DescribeNodeMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNodeMetricsResponse{}
	_body, _err := client.DescribeNodeMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOutlineBindingWithOptions(request *DescribeOutlineBindingRequest, runtime *util.RuntimeOptions) (_result *DescribeOutlineBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsConcurrentLimit)) {
		body["IsConcurrentLimit"] = request.IsConcurrentLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOutlineBinding"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOutlineBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOutlineBinding(request *DescribeOutlineBindingRequest) (_result *DescribeOutlineBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOutlineBindingResponse{}
	_body, _err := client.DescribeOutlineBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeParametersWithOptions(request *DescribeParametersRequest, runtime *util.RuntimeOptions) (_result *DescribeParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimension)) {
		body["Dimension"] = request.Dimension
	}

	if !tea.BoolValue(util.IsUnset(request.DimensionValue)) {
		body["DimensionValue"] = request.DimensionValue
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeParameters"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParameters(request *DescribeParametersRequest) (_result *DescribeParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParametersResponse{}
	_body, _err := client.DescribeParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeParametersHistoryWithOptions(request *DescribeParametersHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeParametersHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimension)) {
		body["Dimension"] = request.Dimension
	}

	if !tea.BoolValue(util.IsUnset(request.DimensionValue)) {
		body["DimensionValue"] = request.DimensionValue
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeParametersHistory"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeParametersHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParametersHistory(request *DescribeParametersHistoryRequest) (_result *DescribeParametersHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParametersHistoryResponse{}
	_body, _err := client.DescribeParametersHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecommendIndexWithOptions(request *DescribeRecommendIndexRequest, runtime *util.RuntimeOptions) (_result *DescribeRecommendIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRecommendIndex"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRecommendIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecommendIndex(request *DescribeRecommendIndexRequest) (_result *DescribeRecommendIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecommendIndexResponse{}
	_body, _err := client.DescribeRecommendIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLDetailsWithOptions(request *DescribeSQLDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLDetails"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLDetails(request *DescribeSQLDetailsRequest) (_result *DescribeSQLDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLDetailsResponse{}
	_body, _err := client.DescribeSQLDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLHistoryListWithOptions(request *DescribeSQLHistoryListRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLHistoryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLHistoryList"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLHistoryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLHistoryList(request *DescribeSQLHistoryListRequest) (_result *DescribeSQLHistoryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLHistoryListResponse{}
	_body, _err := client.DescribeSQLHistoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLPlansWithOptions(request *DescribeSQLPlansRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLPlans"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLPlans(request *DescribeSQLPlansRequest) (_result *DescribeSQLPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLPlansResponse{}
	_body, _err := client.DescribeSQLPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityIpGroupsWithOptions(request *DescribeSecurityIpGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityIpGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSecurityIpGroups"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSecurityIpGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityIpGroups(request *DescribeSecurityIpGroupsRequest) (_result *DescribeSecurityIpGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityIpGroupsResponse{}
	_body, _err := client.DescribeSecurityIpGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowSQLHistoryListWithOptions(request *DescribeSlowSQLHistoryListRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowSQLHistoryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowSQLHistoryList"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlowSQLHistoryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowSQLHistoryList(request *DescribeSlowSQLHistoryListRequest) (_result *DescribeSlowSQLHistoryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowSQLHistoryListResponse{}
	_body, _err := client.DescribeSlowSQLHistoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowSQLListWithOptions(tmpReq *DescribeSlowSQLListRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowSQLListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeSlowSQLListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FilterCondition)) {
		request.FilterConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterCondition, tea.String("FilterCondition"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FilterConditionShrink)) {
		body["FilterCondition"] = request.FilterConditionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIp)) {
		body["NodeIp"] = request.NodeIp
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKeyWord)) {
		body["SearchKeyWord"] = request.SearchKeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.SearchParameter)) {
		body["SearchParameter"] = request.SearchParameter
	}

	if !tea.BoolValue(util.IsUnset(request.SearchRule)) {
		body["SearchRule"] = request.SearchRule
	}

	if !tea.BoolValue(util.IsUnset(request.SearchValue)) {
		body["SearchValue"] = request.SearchValue
	}

	if !tea.BoolValue(util.IsUnset(request.SortColumn)) {
		body["SortColumn"] = request.SortColumn
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		body["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowSQLList"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlowSQLListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowSQLList(request *DescribeSlowSQLListRequest) (_result *DescribeSlowSQLListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowSQLListResponse{}
	_body, _err := client.DescribeSlowSQLListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSqlAuditsWithOptions(request *DescribeSqlAuditsRequest, runtime *util.RuntimeOptions) (_result *DescribeSqlAuditsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		body["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteTimeMax)) {
		body["ExecuteTimeMax"] = request.ExecuteTimeMax
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteTimeMin)) {
		body["ExecuteTimeMin"] = request.ExecuteTimeMin
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NeedMasking)) {
		body["NeedMasking"] = request.NeedMasking
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIp)) {
		body["NodeIp"] = request.NodeIp
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorType)) {
		body["OperatorType"] = request.OperatorType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ScanRowsMax)) {
		body["ScanRowsMax"] = request.ScanRowsMax
	}

	if !tea.BoolValue(util.IsUnset(request.ScanRowsMin)) {
		body["ScanRowsMin"] = request.ScanRowsMin
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKeyWord)) {
		body["SearchKeyWord"] = request.SearchKeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKeyWordMethod)) {
		body["SearchKeyWordMethod"] = request.SearchKeyWordMethod
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSqlAudits"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSqlAuditsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSqlAudits(request *DescribeSqlAuditsRequest) (_result *DescribeSqlAuditsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSqlAuditsResponse{}
	_body, _err := client.DescribeSqlAuditsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTenantWithOptions(request *DescribeTenantRequest, runtime *util.RuntimeOptions) (_result *DescribeTenantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTenant"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTenantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTenant(request *DescribeTenantRequest) (_result *DescribeTenantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTenantResponse{}
	_body, _err := client.DescribeTenantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTenantMetricsWithOptions(request *DescribeTenantMetricsRequest, runtime *util.RuntimeOptions) (_result *DescribeTenantMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Metrics)) {
		body["Metrics"] = request.Metrics
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantIdList)) {
		body["TenantIdList"] = request.TenantIdList
	}

	if !tea.BoolValue(util.IsUnset(request.TenantName)) {
		body["TenantName"] = request.TenantName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTenantMetrics"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTenantMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTenantMetrics(request *DescribeTenantMetricsRequest) (_result *DescribeTenantMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTenantMetricsResponse{}
	_body, _err := client.DescribeTenantMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTenantUserRolesWithOptions(runtime *util.RuntimeOptions) (_result *DescribeTenantUserRolesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeTenantUserRoles"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTenantUserRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTenantUserRoles() (_result *DescribeTenantUserRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTenantUserRolesResponse{}
	_body, _err := client.DescribeTenantUserRolesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTenantUsersWithOptions(request *DescribeTenantUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeTenantUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		body["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTenantUsers"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTenantUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTenantUsers(request *DescribeTenantUsersRequest) (_result *DescribeTenantUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTenantUsersResponse{}
	_body, _err := client.DescribeTenantUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTenantZonesReadWithOptions(request *DescribeTenantZonesReadRequest, runtime *util.RuntimeOptions) (_result *DescribeTenantZonesReadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTenantZonesRead"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTenantZonesReadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTenantZonesRead(request *DescribeTenantZonesReadRequest) (_result *DescribeTenantZonesReadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTenantZonesReadResponse{}
	_body, _err := client.DescribeTenantZonesReadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTenantsWithOptions(request *DescribeTenantsRequest, runtime *util.RuntimeOptions) (_result *DescribeTenantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantName)) {
		body["TenantName"] = request.TenantName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTenants"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTenantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTenants(request *DescribeTenantsRequest) (_result *DescribeTenantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTenantsResponse{}
	_body, _err := client.DescribeTenantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTimeZonesWithOptions(runtime *util.RuntimeOptions) (_result *DescribeTimeZonesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeTimeZones"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTimeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTimeZones() (_result *DescribeTimeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTimeZonesResponse{}
	_body, _err := client.DescribeTimeZonesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTopSQLListWithOptions(tmpReq *DescribeTopSQLListRequest, runtime *util.RuntimeOptions) (_result *DescribeTopSQLListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeTopSQLListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FilterCondition)) {
		request.FilterConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterCondition, tea.String("FilterCondition"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FilterConditionShrink)) {
		body["FilterCondition"] = request.FilterConditionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIp)) {
		body["NodeIp"] = request.NodeIp
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKeyWord)) {
		body["SearchKeyWord"] = request.SearchKeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.SearchParameter)) {
		body["SearchParameter"] = request.SearchParameter
	}

	if !tea.BoolValue(util.IsUnset(request.SearchRule)) {
		body["SearchRule"] = request.SearchRule
	}

	if !tea.BoolValue(util.IsUnset(request.SearchValue)) {
		body["SearchValue"] = request.SearchValue
	}

	if !tea.BoolValue(util.IsUnset(request.SortColumn)) {
		body["SortColumn"] = request.SortColumn
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		body["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTopSQLList"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTopSQLListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTopSQLList(request *DescribeTopSQLListRequest) (_result *DescribeTopSQLListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTopSQLListResponse{}
	_body, _err := client.DescribeTopSQLListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeployType)) {
		body["DeployType"] = request.DeployType
	}

	if !tea.BoolValue(util.IsUnset(request.Series)) {
		body["Series"] = request.Series
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDatabaseDescriptionWithOptions(request *ModifyDatabaseDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDatabaseDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDatabaseDescription"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDatabaseDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDatabaseDescription(request *ModifyDatabaseDescriptionRequest) (_result *ModifyDatabaseDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDatabaseDescriptionResponse{}
	_body, _err := client.ModifyDatabaseDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDatabaseUserRolesWithOptions(request *ModifyDatabaseUserRolesRequest, runtime *util.RuntimeOptions) (_result *ModifyDatabaseUserRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDatabaseUserRoles"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDatabaseUserRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDatabaseUserRoles(request *ModifyDatabaseUserRolesRequest) (_result *ModifyDatabaseUserRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDatabaseUserRolesResponse{}
	_body, _err := client.ModifyDatabaseUserRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceNameWithOptions(request *ModifyInstanceNameRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceName"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (_result *ModifyInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceNameResponse{}
	_body, _err := client.ModifyInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyParametersWithOptions(request *ModifyParametersRequest, runtime *util.RuntimeOptions) (_result *ModifyParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimension)) {
		body["Dimension"] = request.Dimension
	}

	if !tea.BoolValue(util.IsUnset(request.DimensionValue)) {
		body["DimensionValue"] = request.DimensionValue
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["Parameters"] = request.Parameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyParameters"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyParameters(request *ModifyParametersRequest) (_result *ModifyParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyParametersResponse{}
	_body, _err := client.ModifyParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySecurityIpsWithOptions(request *ModifySecurityIpsRequest, runtime *util.RuntimeOptions) (_result *ModifySecurityIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIpGroupName)) {
		body["SecurityIpGroupName"] = request.SecurityIpGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIps)) {
		body["SecurityIps"] = request.SecurityIps
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySecurityIps"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySecurityIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySecurityIps(request *ModifySecurityIpsRequest) (_result *ModifySecurityIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySecurityIpsResponse{}
	_body, _err := client.ModifySecurityIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTenantPrimaryZoneWithOptions(request *ModifyTenantPrimaryZoneRequest, runtime *util.RuntimeOptions) (_result *ModifyTenantPrimaryZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryZone)) {
		body["PrimaryZone"] = request.PrimaryZone
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTenantPrimaryZone"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTenantPrimaryZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTenantPrimaryZone(request *ModifyTenantPrimaryZoneRequest) (_result *ModifyTenantPrimaryZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTenantPrimaryZoneResponse{}
	_body, _err := client.ModifyTenantPrimaryZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTenantResourceWithOptions(request *ModifyTenantResourceRequest, runtime *util.RuntimeOptions) (_result *ModifyTenantResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		body["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		body["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTenantResource"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTenantResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTenantResource(request *ModifyTenantResourceRequest) (_result *ModifyTenantResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTenantResourceResponse{}
	_body, _err := client.ModifyTenantResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTenantUserDescriptionWithOptions(request *ModifyTenantUserDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyTenantUserDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTenantUserDescription"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTenantUserDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTenantUserDescription(request *ModifyTenantUserDescriptionRequest) (_result *ModifyTenantUserDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTenantUserDescriptionResponse{}
	_body, _err := client.ModifyTenantUserDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTenantUserPasswordWithOptions(request *ModifyTenantUserPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyTenantUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserPassword)) {
		body["UserPassword"] = request.UserPassword
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTenantUserPassword"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTenantUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTenantUserPassword(request *ModifyTenantUserPasswordRequest) (_result *ModifyTenantUserPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTenantUserPasswordResponse{}
	_body, _err := client.ModifyTenantUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTenantUserRolesWithOptions(request *ModifyTenantUserRolesRequest, runtime *util.RuntimeOptions) (_result *ModifyTenantUserRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyType)) {
		body["ModifyType"] = request.ModifyType
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserRole)) {
		body["UserRole"] = request.UserRole
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTenantUserRoles"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTenantUserRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTenantUserRoles(request *ModifyTenantUserRolesRequest) (_result *ModifyTenantUserRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTenantUserRolesResponse{}
	_body, _err := client.ModifyTenantUserRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTenantUserStatusWithOptions(request *ModifyTenantUserStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyTenantUserStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserStatus)) {
		body["UserStatus"] = request.UserStatus
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTenantUserStatus"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTenantUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTenantUserStatus(request *ModifyTenantUserStatusRequest) (_result *ModifyTenantUserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTenantUserStatusResponse{}
	_body, _err := client.ModifyTenantUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
