// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddDatabaseRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxx
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DbDescription *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// 1234!2#%A***
	DbPassword *string `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// test_usr
	DbUserName *string `json:"DbUserName,omitempty" xml:"DbUserName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dg-8d9bqu1030m****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 127.0.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDatabaseRequest) GoString() string {
	return s.String()
}

func (s *AddDatabaseRequest) SetClientToken(v string) *AddDatabaseRequest {
	s.ClientToken = &v
	return s
}

func (s *AddDatabaseRequest) SetDbDescription(v string) *AddDatabaseRequest {
	s.DbDescription = &v
	return s
}

func (s *AddDatabaseRequest) SetDbName(v string) *AddDatabaseRequest {
	s.DbName = &v
	return s
}

func (s *AddDatabaseRequest) SetDbPassword(v string) *AddDatabaseRequest {
	s.DbPassword = &v
	return s
}

func (s *AddDatabaseRequest) SetDbType(v string) *AddDatabaseRequest {
	s.DbType = &v
	return s
}

func (s *AddDatabaseRequest) SetDbUserName(v string) *AddDatabaseRequest {
	s.DbUserName = &v
	return s
}

func (s *AddDatabaseRequest) SetGatewayId(v string) *AddDatabaseRequest {
	s.GatewayId = &v
	return s
}

func (s *AddDatabaseRequest) SetHost(v string) *AddDatabaseRequest {
	s.Host = &v
	return s
}

func (s *AddDatabaseRequest) SetPort(v int32) *AddDatabaseRequest {
	s.Port = &v
	return s
}

func (s *AddDatabaseRequest) SetRegionId(v string) *AddDatabaseRequest {
	s.RegionId = &v
	return s
}

type AddDatabaseResponseBody struct {
	// example:
	//
	// 200
	Code       *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	DbInstance *AddDatabaseResponseBodyDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Struct"`
	// example:
	//
	// SYSTEM_ERR
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// A8B2EED2-70EF-51F1-8820-914F9AC9BAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *AddDatabaseResponseBody) SetCode(v string) *AddDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *AddDatabaseResponseBody) SetDbInstance(v *AddDatabaseResponseBodyDbInstance) *AddDatabaseResponseBody {
	s.DbInstance = v
	return s
}

func (s *AddDatabaseResponseBody) SetErrorMsg(v string) *AddDatabaseResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AddDatabaseResponseBody) SetRequestId(v string) *AddDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDatabaseResponseBody) SetSuccess(v bool) *AddDatabaseResponseBody {
	s.Success = &v
	return s
}

type AddDatabaseResponseBodyDbInstance struct {
	// 连接使用的主机
	//
	// example:
	//
	// 10.0.0.1
	ConnectHost *string `json:"ConnectHost,omitempty" xml:"ConnectHost,omitempty"`
	// example:
	//
	// 32875
	ConnectPort *int32 `json:"ConnectPort,omitempty" xml:"ConnectPort,omitempty"`
	// 备注信息
	//
	// example:
	//
	// test
	DbDescription *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	// 数据库类型。
	//
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// 数据库所在网关ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-8d9bqu1030mhk0ix
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// 网关名称
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// 通过网关所在宿主机去访问数据库的地址。
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 数据库实例ID
	//
	// example:
	//
	// dg-db-rgfg9p4586o7y79b
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 当前实例的状态
	//
	// example:
	//
	// SUCCESS
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// 网络类型
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// 节点的ID
	//
	// example:
	//
	// dg-node-r0SR-E90lsIRNgj6B_9m
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// 归属主账号ID
	//
	// example:
	//
	// 167xxxxxxxxxx
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// 通过网关所在宿主机去访问数据库的端口。
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// 所在的地域编号
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 服务类型
	//
	// example:
	//
	// ECS
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// 用户ID
	//
	// example:
	//
	// 167xxxxxxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// VpcId
	//
	// example:
	//
	// vpc-xxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// VpcInstanceId
	//
	// example:
	//
	// i-xxxxxxx
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s AddDatabaseResponseBodyDbInstance) String() string {
	return tea.Prettify(s)
}

func (s AddDatabaseResponseBodyDbInstance) GoString() string {
	return s.String()
}

func (s *AddDatabaseResponseBodyDbInstance) SetConnectHost(v string) *AddDatabaseResponseBodyDbInstance {
	s.ConnectHost = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetConnectPort(v int32) *AddDatabaseResponseBodyDbInstance {
	s.ConnectPort = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetDbDescription(v string) *AddDatabaseResponseBodyDbInstance {
	s.DbDescription = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetDbType(v string) *AddDatabaseResponseBodyDbInstance {
	s.DbType = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetGatewayId(v string) *AddDatabaseResponseBodyDbInstance {
	s.GatewayId = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetGatewayName(v string) *AddDatabaseResponseBodyDbInstance {
	s.GatewayName = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetHost(v string) *AddDatabaseResponseBodyDbInstance {
	s.Host = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetInstanceId(v string) *AddDatabaseResponseBodyDbInstance {
	s.InstanceId = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetInstanceStatus(v string) *AddDatabaseResponseBodyDbInstance {
	s.InstanceStatus = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetNetworkType(v string) *AddDatabaseResponseBodyDbInstance {
	s.NetworkType = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetNodeId(v string) *AddDatabaseResponseBodyDbInstance {
	s.NodeId = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetParentId(v string) *AddDatabaseResponseBodyDbInstance {
	s.ParentId = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetPort(v int32) *AddDatabaseResponseBodyDbInstance {
	s.Port = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetRegionId(v string) *AddDatabaseResponseBodyDbInstance {
	s.RegionId = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetServiceType(v string) *AddDatabaseResponseBodyDbInstance {
	s.ServiceType = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetUserId(v string) *AddDatabaseResponseBodyDbInstance {
	s.UserId = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetVpcId(v string) *AddDatabaseResponseBodyDbInstance {
	s.VpcId = &v
	return s
}

func (s *AddDatabaseResponseBodyDbInstance) SetVpcInstanceId(v string) *AddDatabaseResponseBodyDbInstance {
	s.VpcInstanceId = &v
	return s
}

type AddDatabaseResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDatabaseResponse) GoString() string {
	return s.String()
}

func (s *AddDatabaseResponse) SetHeaders(v map[string]*string) *AddDatabaseResponse {
	s.Headers = v
	return s
}

func (s *AddDatabaseResponse) SetStatusCode(v int32) *AddDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDatabaseResponse) SetBody(v *AddDatabaseResponseBody) *AddDatabaseResponse {
	s.Body = v
	return s
}

type AddDatabaseListRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"host":"127.0.0.1","port":"3306","gatewayId":"dg-xsdasdasdas****"}]
	DatabaseString *string `json:"DatabaseString,omitempty" xml:"DatabaseString,omitempty"`
}

func (s AddDatabaseListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDatabaseListRequest) GoString() string {
	return s.String()
}

func (s *AddDatabaseListRequest) SetClientToken(v string) *AddDatabaseListRequest {
	s.ClientToken = &v
	return s
}

func (s *AddDatabaseListRequest) SetDatabaseString(v string) *AddDatabaseListRequest {
	s.DatabaseString = &v
	return s
}

type AddDatabaseListResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESS
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// ERROR
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 514F794F-7E30-5DAA-97C0-0B0D75DA0259
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDatabaseListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDatabaseListResponseBody) GoString() string {
	return s.String()
}

func (s *AddDatabaseListResponseBody) SetCode(v string) *AddDatabaseListResponseBody {
	s.Code = &v
	return s
}

func (s *AddDatabaseListResponseBody) SetData(v string) *AddDatabaseListResponseBody {
	s.Data = &v
	return s
}

func (s *AddDatabaseListResponseBody) SetErrorMsg(v string) *AddDatabaseListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AddDatabaseListResponseBody) SetRequestId(v string) *AddDatabaseListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDatabaseListResponseBody) SetSuccess(v bool) *AddDatabaseListResponseBody {
	s.Success = &v
	return s
}

type AddDatabaseListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDatabaseListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDatabaseListResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDatabaseListResponse) GoString() string {
	return s.String()
}

func (s *AddDatabaseListResponse) SetHeaders(v map[string]*string) *AddDatabaseListResponse {
	s.Headers = v
	return s
}

func (s *AddDatabaseListResponse) SetStatusCode(v int32) *AddDatabaseListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDatabaseListResponse) SetBody(v *AddDatabaseListResponseBody) *AddDatabaseListResponse {
	s.Body = v
	return s
}

type CheckDGEnabledResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data     *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 8F29E3E9-2847-53BF-ADF0-130E3CEA5C63
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckDGEnabledResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDGEnabledResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDGEnabledResponseBody) SetCode(v string) *CheckDGEnabledResponseBody {
	s.Code = &v
	return s
}

func (s *CheckDGEnabledResponseBody) SetData(v string) *CheckDGEnabledResponseBody {
	s.Data = &v
	return s
}

func (s *CheckDGEnabledResponseBody) SetErrorMsg(v string) *CheckDGEnabledResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CheckDGEnabledResponseBody) SetRequestId(v string) *CheckDGEnabledResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDGEnabledResponseBody) SetSuccess(v bool) *CheckDGEnabledResponseBody {
	s.Success = &v
	return s
}

type CheckDGEnabledResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDGEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDGEnabledResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDGEnabledResponse) GoString() string {
	return s.String()
}

func (s *CheckDGEnabledResponse) SetHeaders(v map[string]*string) *CheckDGEnabledResponse {
	s.Headers = v
	return s
}

func (s *CheckDGEnabledResponse) SetStatusCode(v int32) *CheckDGEnabledResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDGEnabledResponse) SetBody(v *CheckDGEnabledResponseBody) *CheckDGEnabledResponse {
	s.Body = v
	return s
}

type ConnectDatabaseRequest struct {
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// 1234!2#%A****
	DbPassword *string `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// test_usr
	DbUserName *string `json:"DbUserName,omitempty" xml:"DbUserName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dg-58c36y906675****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 127.0.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ConnectDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ConnectDatabaseRequest) GoString() string {
	return s.String()
}

func (s *ConnectDatabaseRequest) SetDbName(v string) *ConnectDatabaseRequest {
	s.DbName = &v
	return s
}

func (s *ConnectDatabaseRequest) SetDbPassword(v string) *ConnectDatabaseRequest {
	s.DbPassword = &v
	return s
}

func (s *ConnectDatabaseRequest) SetDbType(v string) *ConnectDatabaseRequest {
	s.DbType = &v
	return s
}

func (s *ConnectDatabaseRequest) SetDbUserName(v string) *ConnectDatabaseRequest {
	s.DbUserName = &v
	return s
}

func (s *ConnectDatabaseRequest) SetGatewayId(v string) *ConnectDatabaseRequest {
	s.GatewayId = &v
	return s
}

func (s *ConnectDatabaseRequest) SetHost(v string) *ConnectDatabaseRequest {
	s.Host = &v
	return s
}

func (s *ConnectDatabaseRequest) SetPort(v int32) *ConnectDatabaseRequest {
	s.Port = &v
	return s
}

type ConnectDatabaseResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// F4EFCDC5-E69C-5A6F-B170-C5379D9D6811
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConnectDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConnectDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *ConnectDatabaseResponseBody) SetCode(v string) *ConnectDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *ConnectDatabaseResponseBody) SetData(v string) *ConnectDatabaseResponseBody {
	s.Data = &v
	return s
}

func (s *ConnectDatabaseResponseBody) SetErrorMsg(v string) *ConnectDatabaseResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ConnectDatabaseResponseBody) SetRequestId(v string) *ConnectDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConnectDatabaseResponseBody) SetSuccess(v bool) *ConnectDatabaseResponseBody {
	s.Success = &v
	return s
}

type ConnectDatabaseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConnectDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConnectDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ConnectDatabaseResponse) GoString() string {
	return s.String()
}

func (s *ConnectDatabaseResponse) SetHeaders(v map[string]*string) *ConnectDatabaseResponse {
	s.Headers = v
	return s
}

func (s *ConnectDatabaseResponse) SetStatusCode(v int32) *ConnectDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ConnectDatabaseResponse) SetBody(v *ConnectDatabaseResponseBody) *ConnectDatabaseResponse {
	s.Body = v
	return s
}

type CreateGatewayRequest struct {
	GatewayDesc *string `json:"GatewayDesc,omitempty" xml:"GatewayDesc,omitempty"`
	// This parameter is required.
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequest) SetGatewayDesc(v string) *CreateGatewayRequest {
	s.GatewayDesc = &v
	return s
}

func (s *CreateGatewayRequest) SetGatewayName(v string) *CreateGatewayRequest {
	s.GatewayName = &v
	return s
}

func (s *CreateGatewayRequest) SetRegionId(v string) *CreateGatewayRequest {
	s.RegionId = &v
	return s
}

type CreateGatewayResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// dg-nmz841r7b681****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// A9A8885B-3626-592E-9149-8D2971A545AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponseBody) SetCode(v string) *CreateGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayResponseBody) SetData(v string) *CreateGatewayResponseBody {
	s.Data = &v
	return s
}

func (s *CreateGatewayResponseBody) SetErrorMsg(v string) *CreateGatewayResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateGatewayResponseBody) SetRequestId(v string) *CreateGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayResponseBody) SetSuccess(v bool) *CreateGatewayResponseBody {
	s.Success = &v
	return s
}

type CreateGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponse) SetHeaders(v map[string]*string) *CreateGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayResponse) SetStatusCode(v int32) *CreateGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayResponse) SetBody(v *CreateGatewayResponseBody) *CreateGatewayResponse {
	s.Body = v
	return s
}

type CreateGatewayVerifyCodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dg-8e0j08630s08****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
}

func (s CreateGatewayVerifyCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayVerifyCodeRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayVerifyCodeRequest) SetGatewayId(v string) *CreateGatewayVerifyCodeRequest {
	s.GatewayId = &v
	return s
}

type CreateGatewayVerifyCodeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 0N3ufVIG43RRAs7diEoep6WHVOHPKj3a
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// C9ADD2AA-27E3-5D62-A676-092EDC5303C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGatewayVerifyCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayVerifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayVerifyCodeResponseBody) SetCode(v string) *CreateGatewayVerifyCodeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayVerifyCodeResponseBody) SetData(v string) *CreateGatewayVerifyCodeResponseBody {
	s.Data = &v
	return s
}

func (s *CreateGatewayVerifyCodeResponseBody) SetErrorMsg(v string) *CreateGatewayVerifyCodeResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateGatewayVerifyCodeResponseBody) SetRequestId(v string) *CreateGatewayVerifyCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayVerifyCodeResponseBody) SetSuccess(v bool) *CreateGatewayVerifyCodeResponseBody {
	s.Success = &v
	return s
}

type CreateGatewayVerifyCodeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayVerifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayVerifyCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayVerifyCodeResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayVerifyCodeResponse) SetHeaders(v map[string]*string) *CreateGatewayVerifyCodeResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayVerifyCodeResponse) SetStatusCode(v int32) *CreateGatewayVerifyCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayVerifyCodeResponse) SetBody(v *CreateGatewayVerifyCodeResponseBody) *CreateGatewayVerifyCodeResponse {
	s.Body = v
	return s
}

type DeleteDatabaseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// db-22h1qa9d452f****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseRequest) SetInstanceId(v string) *DeleteDatabaseRequest {
	s.InstanceId = &v
	return s
}

type DeleteDatabaseResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESS
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// B9FB545B-03E3-5AE3-9D9E-2FE26EE2C48F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponseBody) SetCode(v string) *DeleteDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDatabaseResponseBody) SetData(v string) *DeleteDatabaseResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDatabaseResponseBody) SetErrorMsg(v string) *DeleteDatabaseResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteDatabaseResponseBody) SetRequestId(v string) *DeleteDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatabaseResponseBody) SetSuccess(v bool) *DeleteDatabaseResponseBody {
	s.Success = &v
	return s
}

type DeleteDatabaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponse) SetHeaders(v map[string]*string) *DeleteDatabaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatabaseResponse) SetStatusCode(v int32) *DeleteDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatabaseResponse) SetBody(v *DeleteDatabaseResponseBody) *DeleteDatabaseResponse {
	s.Body = v
	return s
}

type DeleteGatewayRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dg-nmz841r7b681****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
}

func (s DeleteGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRequest) SetGatewayId(v string) *DeleteGatewayRequest {
	s.GatewayId = &v
	return s
}

type DeleteGatewayResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESS
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 8F29E3E9-2847-53BF-ADF0-130E3CEA5C63
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponseBody) SetCode(v string) *DeleteGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetData(v string) *DeleteGatewayResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetErrorMsg(v string) *DeleteGatewayResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetRequestId(v string) *DeleteGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetSuccess(v bool) *DeleteGatewayResponseBody {
	s.Success = &v
	return s
}

type DeleteGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponse) SetHeaders(v map[string]*string) *DeleteGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayResponse) SetStatusCode(v int32) *DeleteGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayResponse) SetBody(v *DeleteGatewayResponseBody) *DeleteGatewayResponse {
	s.Body = v
	return s
}

type DeleteGatewayInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dg-nmz841r7b681****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dg-node-QeH6VfT8GRnPrYWX****
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" xml:"GatewayInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteGatewayInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayInstanceRequest) SetGatewayId(v string) *DeleteGatewayInstanceRequest {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayInstanceRequest) SetGatewayInstanceId(v string) *DeleteGatewayInstanceRequest {
	s.GatewayInstanceId = &v
	return s
}

func (s *DeleteGatewayInstanceRequest) SetRegionId(v string) *DeleteGatewayInstanceRequest {
	s.RegionId = &v
	return s
}

type DeleteGatewayInstanceResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESS
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 8F29E3E9-2847-53BF-ADF0-130E3CEA5C63
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGatewayInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayInstanceResponseBody) SetCode(v string) *DeleteGatewayInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayInstanceResponseBody) SetData(v string) *DeleteGatewayInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGatewayInstanceResponseBody) SetErrorMsg(v string) *DeleteGatewayInstanceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteGatewayInstanceResponseBody) SetRequestId(v string) *DeleteGatewayInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayInstanceResponseBody) SetSuccess(v bool) *DeleteGatewayInstanceResponseBody {
	s.Success = &v
	return s
}

type DeleteGatewayInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayInstanceResponse) SetHeaders(v map[string]*string) *DeleteGatewayInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayInstanceResponse) SetStatusCode(v int32) *DeleteGatewayInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayInstanceResponse) SetBody(v *DeleteGatewayInstanceResponseBody) *DeleteGatewayInstanceResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string                             `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Regions  *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// example:
	//
	// B0557F7A-62C3-50DC-9E09-77CAE658F776
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetCode(v string) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetErrorMsg(v string) *DescribeRegionsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// example:
	//
	// dg.cn-hangzhou.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DownloadGatewayProgramRequest struct {
	// example:
	//
	// 3.0
	DgVersion *string `json:"DgVersion,omitempty" xml:"DgVersion,omitempty"`
	// example:
	//
	// LINUX
	UserOS *string `json:"UserOS,omitempty" xml:"UserOS,omitempty"`
}

func (s DownloadGatewayProgramRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadGatewayProgramRequest) GoString() string {
	return s.String()
}

func (s *DownloadGatewayProgramRequest) SetDgVersion(v string) *DownloadGatewayProgramRequest {
	s.DgVersion = &v
	return s
}

func (s *DownloadGatewayProgramRequest) SetUserOS(v string) *DownloadGatewayProgramRequest {
	s.UserOS = &v
	return s
}

type DownloadGatewayProgramResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://dg-prod-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/dg3-pkgs/DBGateway_linux?Expires=170833****&OSSAccessKeyId=LTAIfHvNGC8y****&Signature=wa8KqWuyZdB0hwx%2BKvbgZeSTW****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 15D856B8-A95C-5DA5-B0FC-67246286EA7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DownloadGatewayProgramResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadGatewayProgramResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadGatewayProgramResponseBody) SetCode(v string) *DownloadGatewayProgramResponseBody {
	s.Code = &v
	return s
}

func (s *DownloadGatewayProgramResponseBody) SetData(v string) *DownloadGatewayProgramResponseBody {
	s.Data = &v
	return s
}

func (s *DownloadGatewayProgramResponseBody) SetErrorMsg(v string) *DownloadGatewayProgramResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DownloadGatewayProgramResponseBody) SetRequestId(v string) *DownloadGatewayProgramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadGatewayProgramResponseBody) SetSuccess(v bool) *DownloadGatewayProgramResponseBody {
	s.Success = &v
	return s
}

type DownloadGatewayProgramResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadGatewayProgramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadGatewayProgramResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadGatewayProgramResponse) GoString() string {
	return s.String()
}

func (s *DownloadGatewayProgramResponse) SetHeaders(v map[string]*string) *DownloadGatewayProgramResponse {
	s.Headers = v
	return s
}

func (s *DownloadGatewayProgramResponse) SetStatusCode(v int32) *DownloadGatewayProgramResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadGatewayProgramResponse) SetBody(v *DownloadGatewayProgramResponseBody) *DownloadGatewayProgramResponse {
	s.Body = v
	return s
}

type FindUserGatewayByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dg-nmz841r7b681****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
}

func (s FindUserGatewayByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s FindUserGatewayByIdRequest) GoString() string {
	return s.String()
}

func (s *FindUserGatewayByIdRequest) SetGatewayId(v string) *FindUserGatewayByIdRequest {
	s.GatewayId = &v
	return s
}

type FindUserGatewayByIdResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string                                 `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Gateway  *FindUserGatewayByIdResponseBodyGateway `json:"Gateway,omitempty" xml:"Gateway,omitempty" type:"Struct"`
	// example:
	//
	// 41FC4DFE-EA8A-5A56-A16C-F607C3409C79
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindUserGatewayByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindUserGatewayByIdResponseBody) GoString() string {
	return s.String()
}

func (s *FindUserGatewayByIdResponseBody) SetCode(v string) *FindUserGatewayByIdResponseBody {
	s.Code = &v
	return s
}

func (s *FindUserGatewayByIdResponseBody) SetErrorMsg(v string) *FindUserGatewayByIdResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *FindUserGatewayByIdResponseBody) SetGateway(v *FindUserGatewayByIdResponseBodyGateway) *FindUserGatewayByIdResponseBody {
	s.Gateway = v
	return s
}

func (s *FindUserGatewayByIdResponseBody) SetRequestId(v string) *FindUserGatewayByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindUserGatewayByIdResponseBody) SetSuccess(v bool) *FindUserGatewayByIdResponseBody {
	s.Success = &v
	return s
}

type FindUserGatewayByIdResponseBodyGateway struct {
	// example:
	//
	// test_user
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// 3.0
	DgVersion *string `json:"DgVersion,omitempty" xml:"DgVersion,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ExceptionMsg *string `json:"ExceptionMsg,omitempty" xml:"ExceptionMsg,omitempty"`
	// 网关的描述
	GatewayDesc *string `json:"GatewayDesc,omitempty" xml:"GatewayDesc,omitempty"`
	// 网关的编号
	//
	// example:
	//
	// dg-pv33g51gw69h****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// 网关的名称
	//
	// This parameter is required.
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// example:
	//
	// 3
	NumOfExceptionInstance *int32 `json:"NumOfExceptionInstance,omitempty" xml:"NumOfExceptionInstance,omitempty"`
	// example:
	//
	// 3
	NumOfRunningInstance *int32 `json:"NumOfRunningInstance,omitempty" xml:"NumOfRunningInstance,omitempty"`
	// 地域的编号
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 网关的状态
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 用户的编号
	//
	// example:
	//
	// 100****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s FindUserGatewayByIdResponseBodyGateway) String() string {
	return tea.Prettify(s)
}

func (s FindUserGatewayByIdResponseBodyGateway) GoString() string {
	return s.String()
}

func (s *FindUserGatewayByIdResponseBodyGateway) SetCreatorId(v string) *FindUserGatewayByIdResponseBodyGateway {
	s.CreatorId = &v
	return s
}

func (s *FindUserGatewayByIdResponseBodyGateway) SetDgVersion(v string) *FindUserGatewayByIdResponseBodyGateway {
	s.DgVersion = &v
	return s
}

func (s *FindUserGatewayByIdResponseBodyGateway) SetExceptionMsg(v string) *FindUserGatewayByIdResponseBodyGateway {
	s.ExceptionMsg = &v
	return s
}

func (s *FindUserGatewayByIdResponseBodyGateway) SetGatewayDesc(v string) *FindUserGatewayByIdResponseBodyGateway {
	s.GatewayDesc = &v
	return s
}

func (s *FindUserGatewayByIdResponseBodyGateway) SetGatewayId(v string) *FindUserGatewayByIdResponseBodyGateway {
	s.GatewayId = &v
	return s
}

func (s *FindUserGatewayByIdResponseBodyGateway) SetGatewayName(v string) *FindUserGatewayByIdResponseBodyGateway {
	s.GatewayName = &v
	return s
}

func (s *FindUserGatewayByIdResponseBodyGateway) SetNumOfExceptionInstance(v int32) *FindUserGatewayByIdResponseBodyGateway {
	s.NumOfExceptionInstance = &v
	return s
}

func (s *FindUserGatewayByIdResponseBodyGateway) SetNumOfRunningInstance(v int32) *FindUserGatewayByIdResponseBodyGateway {
	s.NumOfRunningInstance = &v
	return s
}

func (s *FindUserGatewayByIdResponseBodyGateway) SetRegionId(v string) *FindUserGatewayByIdResponseBodyGateway {
	s.RegionId = &v
	return s
}

func (s *FindUserGatewayByIdResponseBodyGateway) SetStatus(v string) *FindUserGatewayByIdResponseBodyGateway {
	s.Status = &v
	return s
}

func (s *FindUserGatewayByIdResponseBodyGateway) SetUserId(v string) *FindUserGatewayByIdResponseBodyGateway {
	s.UserId = &v
	return s
}

type FindUserGatewayByIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindUserGatewayByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindUserGatewayByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s FindUserGatewayByIdResponse) GoString() string {
	return s.String()
}

func (s *FindUserGatewayByIdResponse) SetHeaders(v map[string]*string) *FindUserGatewayByIdResponse {
	s.Headers = v
	return s
}

func (s *FindUserGatewayByIdResponse) SetStatusCode(v int32) *FindUserGatewayByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *FindUserGatewayByIdResponse) SetBody(v *FindUserGatewayByIdResponseBody) *FindUserGatewayByIdResponse {
	s.Body = v
	return s
}

type GetUserDatabasesRequest struct {
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// dg-nmz841r7b681****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// 127.0.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// db-22h1qa9d452f****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetUserDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserDatabasesRequest) GoString() string {
	return s.String()
}

func (s *GetUserDatabasesRequest) SetDbType(v string) *GetUserDatabasesRequest {
	s.DbType = &v
	return s
}

func (s *GetUserDatabasesRequest) SetGatewayId(v string) *GetUserDatabasesRequest {
	s.GatewayId = &v
	return s
}

func (s *GetUserDatabasesRequest) SetHost(v string) *GetUserDatabasesRequest {
	s.Host = &v
	return s
}

func (s *GetUserDatabasesRequest) SetInstanceId(v string) *GetUserDatabasesRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserDatabasesRequest) SetPageNumber(v string) *GetUserDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserDatabasesRequest) SetPageSize(v string) *GetUserDatabasesRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserDatabasesRequest) SetPort(v int32) *GetUserDatabasesRequest {
	s.Port = &v
	return s
}

func (s *GetUserDatabasesRequest) SetRegionId(v string) *GetUserDatabasesRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserDatabasesRequest) SetSearchKey(v string) *GetUserDatabasesRequest {
	s.SearchKey = &v
	return s
}

type GetUserDatabasesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 100
	Count          *int32                                      `json:"Count,omitempty" xml:"Count,omitempty"`
	DbInstanceList *GetUserDatabasesResponseBodyDbInstanceList `json:"DbInstanceList,omitempty" xml:"DbInstanceList,omitempty" type:"Struct"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 514F794F-7E30-5DAA-97C0-0B0D75DA0259
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserDatabasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserDatabasesResponseBody) SetCode(v string) *GetUserDatabasesResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserDatabasesResponseBody) SetCount(v int32) *GetUserDatabasesResponseBody {
	s.Count = &v
	return s
}

func (s *GetUserDatabasesResponseBody) SetDbInstanceList(v *GetUserDatabasesResponseBodyDbInstanceList) *GetUserDatabasesResponseBody {
	s.DbInstanceList = v
	return s
}

func (s *GetUserDatabasesResponseBody) SetErrorMsg(v string) *GetUserDatabasesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetUserDatabasesResponseBody) SetRequestId(v string) *GetUserDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserDatabasesResponseBody) SetSuccess(v string) *GetUserDatabasesResponseBody {
	s.Success = &v
	return s
}

type GetUserDatabasesResponseBodyDbInstanceList struct {
	DbInstance []*GetUserDatabasesResponseBodyDbInstanceListDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Repeated"`
}

func (s GetUserDatabasesResponseBodyDbInstanceList) String() string {
	return tea.Prettify(s)
}

func (s GetUserDatabasesResponseBodyDbInstanceList) GoString() string {
	return s.String()
}

func (s *GetUserDatabasesResponseBodyDbInstanceList) SetDbInstance(v []*GetUserDatabasesResponseBodyDbInstanceListDbInstance) *GetUserDatabasesResponseBodyDbInstanceList {
	s.DbInstance = v
	return s
}

type GetUserDatabasesResponseBodyDbInstanceListDbInstance struct {
	// 连接使用的主机
	//
	// example:
	//
	// 10.0.0.0
	ConnectHost *string `json:"ConnectHost,omitempty" xml:"ConnectHost,omitempty"`
	// example:
	//
	// 10001
	ConnectPort *int32 `json:"ConnectPort,omitempty" xml:"ConnectPort,omitempty"`
	// 备注信息
	DbDescription *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	// 数据库类型。
	//
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// 数据库所在网关ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-pil582nbfe6p****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// 网关名称
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 2023-05-03 00:00:00
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 通过网关所在宿主机去访问数据库的地址。
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 数据库实例ID
	//
	// example:
	//
	// db-22h1qa9d452f****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 当前实例的状态
	//
	// example:
	//
	// SUCCESS
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// 网络类型
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// 节点的ID
	//
	// example:
	//
	// dg-node-xxxxx
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// 归属主账号ID
	//
	// example:
	//
	// 100XXXXX
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// 通过网关所在宿主机去访问数据库的端口。
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// 所在的地域编号
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 服务类型
	//
	// example:
	//
	// ECS
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// 用户ID
	//
	// example:
	//
	// 100XXXXX
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// VpcId
	//
	// example:
	//
	// vpc-bp1alpkpdb8fh3avx****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// VpcInstanceId
	//
	// example:
	//
	// i-xxxxxxxxxx
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s GetUserDatabasesResponseBodyDbInstanceListDbInstance) String() string {
	return tea.Prettify(s)
}

func (s GetUserDatabasesResponseBodyDbInstanceListDbInstance) GoString() string {
	return s.String()
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetConnectHost(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.ConnectHost = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetConnectPort(v int32) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.ConnectPort = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetDbDescription(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.DbDescription = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetDbType(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.DbType = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetGatewayId(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.GatewayId = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetGatewayName(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.GatewayName = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetGmtCreate(v int64) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.GmtCreate = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetHost(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.Host = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetInstanceId(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.InstanceId = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetInstanceStatus(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.InstanceStatus = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetNetworkType(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.NetworkType = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetNodeId(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.NodeId = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetParentId(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.ParentId = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetPort(v int32) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.Port = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetRegionId(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.RegionId = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetServiceType(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.ServiceType = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetUserId(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.UserId = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetVpcId(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.VpcId = &v
	return s
}

func (s *GetUserDatabasesResponseBodyDbInstanceListDbInstance) SetVpcInstanceId(v string) *GetUserDatabasesResponseBodyDbInstanceListDbInstance {
	s.VpcInstanceId = &v
	return s
}

type GetUserDatabasesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserDatabasesResponse) GoString() string {
	return s.String()
}

func (s *GetUserDatabasesResponse) SetHeaders(v map[string]*string) *GetUserDatabasesResponse {
	s.Headers = v
	return s
}

func (s *GetUserDatabasesResponse) SetStatusCode(v int32) *GetUserDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserDatabasesResponse) SetBody(v *GetUserDatabasesResponseBody) *GetUserDatabasesResponse {
	s.Body = v
	return s
}

type GetUserGatewayInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dg-nmz841r7b681****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
}

func (s GetUserGatewayInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewayInstancesRequest) GoString() string {
	return s.String()
}

func (s *GetUserGatewayInstancesRequest) SetGatewayId(v string) *GetUserGatewayInstancesRequest {
	s.GatewayId = &v
	return s
}

type GetUserGatewayInstancesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg            *string                                                   `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	GatewayInstanceList []*GetUserGatewayInstancesResponseBodyGatewayInstanceList `json:"GatewayInstanceList,omitempty" xml:"GatewayInstanceList,omitempty" type:"Repeated"`
	// example:
	//
	// 514F794F-7E30-5DAA-97C0-0B0D75DA0259
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserGatewayInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewayInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserGatewayInstancesResponseBody) SetCode(v string) *GetUserGatewayInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBody) SetErrorMsg(v string) *GetUserGatewayInstancesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBody) SetGatewayInstanceList(v []*GetUserGatewayInstancesResponseBodyGatewayInstanceList) *GetUserGatewayInstancesResponseBody {
	s.GatewayInstanceList = v
	return s
}

func (s *GetUserGatewayInstancesResponseBody) SetRequestId(v string) *GetUserGatewayInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBody) SetSuccess(v bool) *GetUserGatewayInstancesResponseBody {
	s.Success = &v
	return s
}

type GetUserGatewayInstancesResponseBodyGatewayInstanceList struct {
	// 连接类型
	//
	// example:
	//
	// internet
	ConnectEndpointType *string `json:"ConnectEndpointType,omitempty" xml:"ConnectEndpointType,omitempty"`
	// 进程的版本号
	//
	// example:
	//
	// 3.0
	CurrentDaemonVersion *string `json:"CurrentDaemonVersion,omitempty" xml:"CurrentDaemonVersion,omitempty"`
	// 版本号
	//
	// example:
	//
	// 3.0
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// 端点地址
	//
	// example:
	//
	// 127.0.XX.XX
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// 网关ID
	//
	// example:
	//
	// dg-159t17m19ps0****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// dg-node-wJOb0tO-gaaWFCug****
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" xml:"GatewayInstanceId,omitempty"`
	// example:
	//
	// STOPPED
	GatewayInstanceStatus *string `json:"GatewayInstanceStatus,omitempty" xml:"GatewayInstanceStatus,omitempty"`
	// 上次更新时间戳
	//
	// example:
	//
	// 1699330233000
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// 本地IP地址
	//
	// example:
	//
	// 127.0.XX.XX
	LocalIP *string `json:"LocalIP,omitempty" xml:"LocalIP,omitempty"`
	// 提示信息
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 主机
	//
	// example:
	//
	// 127.0.0.1
	OutputIP *string `json:"OutputIP,omitempty" xml:"OutputIP,omitempty"`
	// 代表region的资源属性字段
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetUserGatewayInstancesResponseBodyGatewayInstanceList) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewayInstancesResponseBodyGatewayInstanceList) GoString() string {
	return s.String()
}

func (s *GetUserGatewayInstancesResponseBodyGatewayInstanceList) SetConnectEndpointType(v string) *GetUserGatewayInstancesResponseBodyGatewayInstanceList {
	s.ConnectEndpointType = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBodyGatewayInstanceList) SetCurrentDaemonVersion(v string) *GetUserGatewayInstancesResponseBodyGatewayInstanceList {
	s.CurrentDaemonVersion = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBodyGatewayInstanceList) SetCurrentVersion(v string) *GetUserGatewayInstancesResponseBodyGatewayInstanceList {
	s.CurrentVersion = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBodyGatewayInstanceList) SetEndPoint(v string) *GetUserGatewayInstancesResponseBodyGatewayInstanceList {
	s.EndPoint = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBodyGatewayInstanceList) SetGatewayId(v string) *GetUserGatewayInstancesResponseBodyGatewayInstanceList {
	s.GatewayId = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBodyGatewayInstanceList) SetGatewayInstanceId(v string) *GetUserGatewayInstancesResponseBodyGatewayInstanceList {
	s.GatewayInstanceId = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBodyGatewayInstanceList) SetGatewayInstanceStatus(v string) *GetUserGatewayInstancesResponseBodyGatewayInstanceList {
	s.GatewayInstanceStatus = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBodyGatewayInstanceList) SetLastUpdateTime(v int64) *GetUserGatewayInstancesResponseBodyGatewayInstanceList {
	s.LastUpdateTime = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBodyGatewayInstanceList) SetLocalIP(v string) *GetUserGatewayInstancesResponseBodyGatewayInstanceList {
	s.LocalIP = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBodyGatewayInstanceList) SetMessage(v string) *GetUserGatewayInstancesResponseBodyGatewayInstanceList {
	s.Message = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBodyGatewayInstanceList) SetOutputIP(v string) *GetUserGatewayInstancesResponseBodyGatewayInstanceList {
	s.OutputIP = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBodyGatewayInstanceList) SetRegionId(v string) *GetUserGatewayInstancesResponseBodyGatewayInstanceList {
	s.RegionId = &v
	return s
}

type GetUserGatewayInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserGatewayInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserGatewayInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewayInstancesResponse) GoString() string {
	return s.String()
}

func (s *GetUserGatewayInstancesResponse) SetHeaders(v map[string]*string) *GetUserGatewayInstancesResponse {
	s.Headers = v
	return s
}

func (s *GetUserGatewayInstancesResponse) SetStatusCode(v int32) *GetUserGatewayInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserGatewayInstancesResponse) SetBody(v *GetUserGatewayInstancesResponseBody) *GetUserGatewayInstancesResponse {
	s.Body = v
	return s
}

type GetUserGatewaysRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetUserGatewaysRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewaysRequest) GoString() string {
	return s.String()
}

func (s *GetUserGatewaysRequest) SetPageNumber(v int32) *GetUserGatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserGatewaysRequest) SetPageSize(v int32) *GetUserGatewaysRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserGatewaysRequest) SetRegionId(v string) *GetUserGatewaysRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserGatewaysRequest) SetSearchKey(v string) *GetUserGatewaysRequest {
	s.SearchKey = &v
	return s
}

type GetUserGatewaysResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg    *string                                 `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	GatewayList *GetUserGatewaysResponseBodyGatewayList `json:"GatewayList,omitempty" xml:"GatewayList,omitempty" type:"Struct"`
	// example:
	//
	// 41FC4DFE-EA8A-5A56-A16C-F607C3409C79
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserGatewaysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserGatewaysResponseBody) SetCode(v string) *GetUserGatewaysResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserGatewaysResponseBody) SetCount(v int32) *GetUserGatewaysResponseBody {
	s.Count = &v
	return s
}

func (s *GetUserGatewaysResponseBody) SetErrorMsg(v string) *GetUserGatewaysResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetUserGatewaysResponseBody) SetGatewayList(v *GetUserGatewaysResponseBodyGatewayList) *GetUserGatewaysResponseBody {
	s.GatewayList = v
	return s
}

func (s *GetUserGatewaysResponseBody) SetRequestId(v string) *GetUserGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserGatewaysResponseBody) SetSuccess(v bool) *GetUserGatewaysResponseBody {
	s.Success = &v
	return s
}

type GetUserGatewaysResponseBodyGatewayList struct {
	Gateway []*GetUserGatewaysResponseBodyGatewayListGateway `json:"Gateway,omitempty" xml:"Gateway,omitempty" type:"Repeated"`
}

func (s GetUserGatewaysResponseBodyGatewayList) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewaysResponseBodyGatewayList) GoString() string {
	return s.String()
}

func (s *GetUserGatewaysResponseBodyGatewayList) SetGateway(v []*GetUserGatewaysResponseBodyGatewayListGateway) *GetUserGatewaysResponseBodyGatewayList {
	s.Gateway = v
	return s
}

type GetUserGatewaysResponseBodyGatewayListGateway struct {
	// 创建用户id
	//
	// example:
	//
	// 100****
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// 网关版本
	//
	// example:
	//
	// 3.0
	DgVersion *string `json:"DgVersion,omitempty" xml:"DgVersion,omitempty"`
	// 网关异常信息
	//
	// example:
	//
	// exception
	ExceptionMsg *string `json:"ExceptionMsg,omitempty" xml:"ExceptionMsg,omitempty"`
	// 网关的描述
	GatewayDesc *string `json:"GatewayDesc,omitempty" xml:"GatewayDesc,omitempty"`
	// 网关的编号
	//
	// example:
	//
	// dg-lch384wg5701****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// 网关的名称
	//
	// This parameter is required.
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// 异常实例数量
	//
	// example:
	//
	// 3
	NumOfExceptionInstance *int32 `json:"NumOfExceptionInstance,omitempty" xml:"NumOfExceptionInstance,omitempty"`
	// 运行中实例数量
	//
	// example:
	//
	// 3
	NumOfRunningInstance *int32 `json:"NumOfRunningInstance,omitempty" xml:"NumOfRunningInstance,omitempty"`
	// 地域的编号
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 网关的状态
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 用户的编号
	//
	// example:
	//
	// 100****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserGatewaysResponseBodyGatewayListGateway) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewaysResponseBodyGatewayListGateway) GoString() string {
	return s.String()
}

func (s *GetUserGatewaysResponseBodyGatewayListGateway) SetCreatorId(v string) *GetUserGatewaysResponseBodyGatewayListGateway {
	s.CreatorId = &v
	return s
}

func (s *GetUserGatewaysResponseBodyGatewayListGateway) SetDgVersion(v string) *GetUserGatewaysResponseBodyGatewayListGateway {
	s.DgVersion = &v
	return s
}

func (s *GetUserGatewaysResponseBodyGatewayListGateway) SetExceptionMsg(v string) *GetUserGatewaysResponseBodyGatewayListGateway {
	s.ExceptionMsg = &v
	return s
}

func (s *GetUserGatewaysResponseBodyGatewayListGateway) SetGatewayDesc(v string) *GetUserGatewaysResponseBodyGatewayListGateway {
	s.GatewayDesc = &v
	return s
}

func (s *GetUserGatewaysResponseBodyGatewayListGateway) SetGatewayId(v string) *GetUserGatewaysResponseBodyGatewayListGateway {
	s.GatewayId = &v
	return s
}

func (s *GetUserGatewaysResponseBodyGatewayListGateway) SetGatewayName(v string) *GetUserGatewaysResponseBodyGatewayListGateway {
	s.GatewayName = &v
	return s
}

func (s *GetUserGatewaysResponseBodyGatewayListGateway) SetNumOfExceptionInstance(v int32) *GetUserGatewaysResponseBodyGatewayListGateway {
	s.NumOfExceptionInstance = &v
	return s
}

func (s *GetUserGatewaysResponseBodyGatewayListGateway) SetNumOfRunningInstance(v int32) *GetUserGatewaysResponseBodyGatewayListGateway {
	s.NumOfRunningInstance = &v
	return s
}

func (s *GetUserGatewaysResponseBodyGatewayListGateway) SetRegionId(v string) *GetUserGatewaysResponseBodyGatewayListGateway {
	s.RegionId = &v
	return s
}

func (s *GetUserGatewaysResponseBodyGatewayListGateway) SetStatus(v string) *GetUserGatewaysResponseBodyGatewayListGateway {
	s.Status = &v
	return s
}

func (s *GetUserGatewaysResponseBodyGatewayListGateway) SetUserId(v string) *GetUserGatewaysResponseBodyGatewayListGateway {
	s.UserId = &v
	return s
}

type GetUserGatewaysResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserGatewaysResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewaysResponse) GoString() string {
	return s.String()
}

func (s *GetUserGatewaysResponse) SetHeaders(v map[string]*string) *GetUserGatewaysResponse {
	s.Headers = v
	return s
}

func (s *GetUserGatewaysResponse) SetStatusCode(v int32) *GetUserGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserGatewaysResponse) SetBody(v *GetUserGatewaysResponseBody) *GetUserGatewaysResponse {
	s.Body = v
	return s
}

type ListDatabaseAccessPointRequest struct {
	// example:
	//
	// dg-db-n2a285spnpy3****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// example:
	//
	// dg-nmz841r7b681****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// 127.0.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// example:
	//
	// vpc-wz9c473cmu2gg7i7l****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListDatabaseAccessPointRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatabaseAccessPointRequest) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccessPointRequest) SetDbInstanceId(v string) *ListDatabaseAccessPointRequest {
	s.DbInstanceId = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetGatewayId(v string) *ListDatabaseAccessPointRequest {
	s.GatewayId = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetHost(v string) *ListDatabaseAccessPointRequest {
	s.Host = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetPageNumber(v string) *ListDatabaseAccessPointRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetPageSize(v string) *ListDatabaseAccessPointRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetPort(v int32) *ListDatabaseAccessPointRequest {
	s.Port = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetRegionId(v string) *ListDatabaseAccessPointRequest {
	s.RegionId = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetSearchKey(v string) *ListDatabaseAccessPointRequest {
	s.SearchKey = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetVpcId(v string) *ListDatabaseAccessPointRequest {
	s.VpcId = &v
	return s
}

type ListDatabaseAccessPointResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Count                     *int32                                                          `json:"Count,omitempty" xml:"Count,omitempty"`
	DbInstanceAccessPointList []*ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList `json:"DbInstanceAccessPointList,omitempty" xml:"DbInstanceAccessPointList,omitempty" type:"Repeated"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// C9ADD2AA-27E3-5D62-A676-092EDC5303C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDatabaseAccessPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatabaseAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccessPointResponseBody) SetCode(v string) *ListDatabaseAccessPointResponseBody {
	s.Code = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBody) SetCount(v int32) *ListDatabaseAccessPointResponseBody {
	s.Count = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBody) SetDbInstanceAccessPointList(v []*ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList) *ListDatabaseAccessPointResponseBody {
	s.DbInstanceAccessPointList = v
	return s
}

func (s *ListDatabaseAccessPointResponseBody) SetErrorMsg(v string) *ListDatabaseAccessPointResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBody) SetRequestId(v string) *ListDatabaseAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBody) SetSuccess(v string) *ListDatabaseAccessPointResponseBody {
	s.Success = &v
	return s
}

type ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList struct {
	// example:
	//
	// 10.0.0.0
	AccessAddr *string `json:"AccessAddr,omitempty" xml:"AccessAddr,omitempty"`
	// example:
	//
	// 33306
	AccessPort *int32 `json:"AccessPort,omitempty" xml:"AccessPort,omitempty"`
	// example:
	//
	// dg-db-n2a285spnpy3****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// example:
	//
	// 2023-03-09 14:20:04.0
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2023-03-09 14:20:04.0
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// dg-nmz841r7b681****
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// example:
	//
	// cn-hangzhou-j
	VpcAzoneId *string `json:"VpcAzoneId,omitempty" xml:"VpcAzoneId,omitempty"`
	// example:
	//
	// vpc-2ze0612ts436tn0sh****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-8vbgi74rgel72rax4****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList) String() string {
	return tea.Prettify(s)
}

func (s ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList) SetAccessAddr(v string) *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList {
	s.AccessAddr = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList) SetAccessPort(v int32) *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList {
	s.AccessPort = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList) SetDbInstanceId(v string) *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList {
	s.DbInstanceId = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList) SetGmtCreate(v int64) *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList {
	s.GmtCreate = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList) SetGmtModified(v int64) *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList {
	s.GmtModified = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList) SetRouterId(v string) *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList {
	s.RouterId = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList) SetVpcAzoneId(v string) *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList {
	s.VpcAzoneId = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList) SetVpcId(v string) *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList {
	s.VpcId = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList) SetVswitchId(v string) *ListDatabaseAccessPointResponseBodyDbInstanceAccessPointList {
	s.VswitchId = &v
	return s
}

type ListDatabaseAccessPointResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatabaseAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatabaseAccessPointResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatabaseAccessPointResponse) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccessPointResponse) SetHeaders(v map[string]*string) *ListDatabaseAccessPointResponse {
	s.Headers = v
	return s
}

func (s *ListDatabaseAccessPointResponse) SetStatusCode(v int32) *ListDatabaseAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatabaseAccessPointResponse) SetBody(v *ListDatabaseAccessPointResponseBody) *ListDatabaseAccessPointResponse {
	s.Body = v
	return s
}

type ModifyDatabaseRequest struct {
	DbDescription *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// 817130****
	DbPassword *string `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// test
	DbUserName *string `json:"DbUserName,omitempty" xml:"DbUserName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 127.0.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dg-db-n2a285spnpy3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseRequest) SetDbDescription(v string) *ModifyDatabaseRequest {
	s.DbDescription = &v
	return s
}

func (s *ModifyDatabaseRequest) SetDbName(v string) *ModifyDatabaseRequest {
	s.DbName = &v
	return s
}

func (s *ModifyDatabaseRequest) SetDbPassword(v string) *ModifyDatabaseRequest {
	s.DbPassword = &v
	return s
}

func (s *ModifyDatabaseRequest) SetDbType(v string) *ModifyDatabaseRequest {
	s.DbType = &v
	return s
}

func (s *ModifyDatabaseRequest) SetDbUserName(v string) *ModifyDatabaseRequest {
	s.DbUserName = &v
	return s
}

func (s *ModifyDatabaseRequest) SetHost(v string) *ModifyDatabaseRequest {
	s.Host = &v
	return s
}

func (s *ModifyDatabaseRequest) SetInstanceId(v string) *ModifyDatabaseRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDatabaseRequest) SetPort(v int32) *ModifyDatabaseRequest {
	s.Port = &v
	return s
}

type ModifyDatabaseResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESS
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 8F29E3E9-2847-53BF-ADF0-130E3CEA5C63
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseResponseBody) SetCode(v string) *ModifyDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDatabaseResponseBody) SetData(v string) *ModifyDatabaseResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyDatabaseResponseBody) SetErrorMsg(v string) *ModifyDatabaseResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ModifyDatabaseResponseBody) SetRequestId(v string) *ModifyDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDatabaseResponseBody) SetSuccess(v bool) *ModifyDatabaseResponseBody {
	s.Success = &v
	return s
}

type ModifyDatabaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseResponse) SetHeaders(v map[string]*string) *ModifyDatabaseResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseResponse) SetStatusCode(v int32) *ModifyDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseResponse) SetBody(v *ModifyDatabaseResponseBody) *ModifyDatabaseResponse {
	s.Body = v
	return s
}

type ModifyGatewayRequest struct {
	GatewayDesc *string `json:"GatewayDesc,omitempty" xml:"GatewayDesc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dg-nmz841r7b681****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// This parameter is required.
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
}

func (s ModifyGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayRequest) GoString() string {
	return s.String()
}

func (s *ModifyGatewayRequest) SetGatewayDesc(v string) *ModifyGatewayRequest {
	s.GatewayDesc = &v
	return s
}

func (s *ModifyGatewayRequest) SetGatewayId(v string) *ModifyGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *ModifyGatewayRequest) SetGatewayName(v string) *ModifyGatewayRequest {
	s.GatewayName = &v
	return s
}

type ModifyGatewayResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESS
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 18071187-5EA1-5DD4-AAD9-F27C5713CD1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGatewayResponseBody) SetCode(v string) *ModifyGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyGatewayResponseBody) SetData(v string) *ModifyGatewayResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyGatewayResponseBody) SetErrorMsg(v string) *ModifyGatewayResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ModifyGatewayResponseBody) SetRequestId(v string) *ModifyGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGatewayResponseBody) SetSuccess(v bool) *ModifyGatewayResponseBody {
	s.Success = &v
	return s
}

type ModifyGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayResponse) GoString() string {
	return s.String()
}

func (s *ModifyGatewayResponse) SetHeaders(v map[string]*string) *ModifyGatewayResponse {
	s.Headers = v
	return s
}

func (s *ModifyGatewayResponse) SetStatusCode(v int32) *ModifyGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGatewayResponse) SetBody(v *ModifyGatewayResponseBody) *ModifyGatewayResponse {
	s.Body = v
	return s
}

type RetryDatabaseRequest struct {
	// example:
	//
	// XXXXX
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DbDescription *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// HongRui****
	DbPassword *string `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// test
	DbUserName *string `json:"DbUserName,omitempty" xml:"DbUserName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dg-nmz841r7b681****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 127.0.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RetryDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s RetryDatabaseRequest) GoString() string {
	return s.String()
}

func (s *RetryDatabaseRequest) SetClientToken(v string) *RetryDatabaseRequest {
	s.ClientToken = &v
	return s
}

func (s *RetryDatabaseRequest) SetDbDescription(v string) *RetryDatabaseRequest {
	s.DbDescription = &v
	return s
}

func (s *RetryDatabaseRequest) SetDbName(v string) *RetryDatabaseRequest {
	s.DbName = &v
	return s
}

func (s *RetryDatabaseRequest) SetDbPassword(v string) *RetryDatabaseRequest {
	s.DbPassword = &v
	return s
}

func (s *RetryDatabaseRequest) SetDbType(v string) *RetryDatabaseRequest {
	s.DbType = &v
	return s
}

func (s *RetryDatabaseRequest) SetDbUserName(v string) *RetryDatabaseRequest {
	s.DbUserName = &v
	return s
}

func (s *RetryDatabaseRequest) SetGatewayId(v string) *RetryDatabaseRequest {
	s.GatewayId = &v
	return s
}

func (s *RetryDatabaseRequest) SetHost(v string) *RetryDatabaseRequest {
	s.Host = &v
	return s
}

func (s *RetryDatabaseRequest) SetPort(v int32) *RetryDatabaseRequest {
	s.Port = &v
	return s
}

func (s *RetryDatabaseRequest) SetRegionId(v string) *RetryDatabaseRequest {
	s.RegionId = &v
	return s
}

type RetryDatabaseResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESS
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 19EDB8E2-FCE6-5797-91F4-80F832C90371
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RetryDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *RetryDatabaseResponseBody) SetCode(v string) *RetryDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *RetryDatabaseResponseBody) SetData(v string) *RetryDatabaseResponseBody {
	s.Data = &v
	return s
}

func (s *RetryDatabaseResponseBody) SetErrorMsg(v string) *RetryDatabaseResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RetryDatabaseResponseBody) SetRequestId(v string) *RetryDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryDatabaseResponseBody) SetSuccess(v bool) *RetryDatabaseResponseBody {
	s.Success = &v
	return s
}

type RetryDatabaseResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryDatabaseResponse) GoString() string {
	return s.String()
}

func (s *RetryDatabaseResponse) SetHeaders(v map[string]*string) *RetryDatabaseResponse {
	s.Headers = v
	return s
}

func (s *RetryDatabaseResponse) SetStatusCode(v int32) *RetryDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryDatabaseResponse) SetBody(v *RetryDatabaseResponseBody) *RetryDatabaseResponse {
	s.Body = v
	return s
}

type StopGatewayRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dg-nmz841r7b681****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// dg-node-hvsGB7qVCaaW****-v-0
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" xml:"GatewayInstanceId,omitempty"`
}

func (s StopGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s StopGatewayRequest) GoString() string {
	return s.String()
}

func (s *StopGatewayRequest) SetGatewayId(v string) *StopGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *StopGatewayRequest) SetGatewayInstanceId(v string) *StopGatewayRequest {
	s.GatewayInstanceId = &v
	return s
}

type StopGatewayResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESS
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Gateway exception, please launch local dg first
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// DE3CC21B-E317-5ED7-A212-A62517EA0022
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *StopGatewayResponseBody) SetCode(v string) *StopGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *StopGatewayResponseBody) SetData(v string) *StopGatewayResponseBody {
	s.Data = &v
	return s
}

func (s *StopGatewayResponseBody) SetErrorMsg(v string) *StopGatewayResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *StopGatewayResponseBody) SetRequestId(v string) *StopGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopGatewayResponseBody) SetSuccess(v bool) *StopGatewayResponseBody {
	s.Success = &v
	return s
}

type StopGatewayResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s StopGatewayResponse) GoString() string {
	return s.String()
}

func (s *StopGatewayResponse) SetHeaders(v map[string]*string) *StopGatewayResponse {
	s.Headers = v
	return s
}

func (s *StopGatewayResponse) SetStatusCode(v int32) *StopGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *StopGatewayResponse) SetBody(v *StopGatewayResponseBody) *StopGatewayResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dms-dg"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AddDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDatabaseResponse
func (client *Client) AddDatabaseWithOptions(request *AddDatabaseRequest, runtime *util.RuntimeOptions) (_result *AddDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DbDescription)) {
		body["DbDescription"] = request.DbDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DbPassword)) {
		body["DbPassword"] = request.DbPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		body["DbType"] = request.DbType
	}

	if !tea.BoolValue(util.IsUnset(request.DbUserName)) {
		body["DbUserName"] = request.DbUserName
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		body["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDatabase"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddDatabaseRequest
//
// @return AddDatabaseResponse
func (client *Client) AddDatabase(request *AddDatabaseRequest) (_result *AddDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDatabaseResponse{}
	_body, _err := client.AddDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddDatabaseListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDatabaseListResponse
func (client *Client) AddDatabaseListWithOptions(request *AddDatabaseListRequest, runtime *util.RuntimeOptions) (_result *AddDatabaseListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseString)) {
		body["DatabaseString"] = request.DatabaseString
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDatabaseList"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDatabaseListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddDatabaseListRequest
//
// @return AddDatabaseListResponse
func (client *Client) AddDatabaseList(request *AddDatabaseListRequest) (_result *AddDatabaseListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDatabaseListResponse{}
	_body, _err := client.AddDatabaseListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CheckDGEnabledRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDGEnabledResponse
func (client *Client) CheckDGEnabledWithOptions(runtime *util.RuntimeOptions) (_result *CheckDGEnabledResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("CheckDGEnabled"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDGEnabledResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return CheckDGEnabledResponse
func (client *Client) CheckDGEnabled() (_result *CheckDGEnabledResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDGEnabledResponse{}
	_body, _err := client.CheckDGEnabledWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ConnectDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConnectDatabaseResponse
func (client *Client) ConnectDatabaseWithOptions(request *ConnectDatabaseRequest, runtime *util.RuntimeOptions) (_result *ConnectDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DbPassword)) {
		body["DbPassword"] = request.DbPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		body["DbType"] = request.DbType
	}

	if !tea.BoolValue(util.IsUnset(request.DbUserName)) {
		body["DbUserName"] = request.DbUserName
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		body["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConnectDatabase"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConnectDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ConnectDatabaseRequest
//
// @return ConnectDatabaseResponse
func (client *Client) ConnectDatabase(request *ConnectDatabaseRequest) (_result *ConnectDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConnectDatabaseResponse{}
	_body, _err := client.ConnectDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建网关
//
// @param request - CreateGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayResponse
func (client *Client) CreateGatewayWithOptions(request *CreateGatewayRequest, runtime *util.RuntimeOptions) (_result *CreateGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayDesc)) {
		body["GatewayDesc"] = request.GatewayDesc
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayName)) {
		body["GatewayName"] = request.GatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGateway"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建网关
//
// @param request - CreateGatewayRequest
//
// @return CreateGatewayResponse
func (client *Client) CreateGateway(request *CreateGatewayRequest) (_result *CreateGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGatewayResponse{}
	_body, _err := client.CreateGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建网关的随机验证码
//
// @param request - CreateGatewayVerifyCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayVerifyCodeResponse
func (client *Client) CreateGatewayVerifyCodeWithOptions(request *CreateGatewayVerifyCodeRequest, runtime *util.RuntimeOptions) (_result *CreateGatewayVerifyCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGatewayVerifyCode"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayVerifyCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建网关的随机验证码
//
// @param request - CreateGatewayVerifyCodeRequest
//
// @return CreateGatewayVerifyCodeResponse
func (client *Client) CreateGatewayVerifyCode(request *CreateGatewayVerifyCodeRequest) (_result *CreateGatewayVerifyCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGatewayVerifyCodeResponse{}
	_body, _err := client.CreateGatewayVerifyCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatabaseResponse
func (client *Client) DeleteDatabaseWithOptions(request *DeleteDatabaseRequest, runtime *util.RuntimeOptions) (_result *DeleteDatabaseResponse, _err error) {
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
		Action:      tea.String("DeleteDatabase"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteDatabaseRequest
//
// @return DeleteDatabaseResponse
func (client *Client) DeleteDatabase(request *DeleteDatabaseRequest) (_result *DeleteDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DeleteDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除网关
//
// @param request - DeleteGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGatewayWithOptions(request *DeleteGatewayRequest, runtime *util.RuntimeOptions) (_result *DeleteGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGateway"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除网关
//
// @param request - DeleteGatewayRequest
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGateway(request *DeleteGatewayRequest) (_result *DeleteGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.DeleteGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteGatewayInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayInstanceResponse
func (client *Client) DeleteGatewayInstanceWithOptions(request *DeleteGatewayInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteGatewayInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayInstanceId)) {
		body["GatewayInstanceId"] = request.GatewayInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGatewayInstance"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteGatewayInstanceRequest
//
// @return DeleteGatewayInstanceResponse
func (client *Client) DeleteGatewayInstance(request *DeleteGatewayInstanceRequest) (_result *DeleteGatewayInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGatewayInstanceResponse{}
	_body, _err := client.DeleteGatewayInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DownloadGatewayProgramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadGatewayProgramResponse
func (client *Client) DownloadGatewayProgramWithOptions(request *DownloadGatewayProgramRequest, runtime *util.RuntimeOptions) (_result *DownloadGatewayProgramResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DgVersion)) {
		body["DgVersion"] = request.DgVersion
	}

	if !tea.BoolValue(util.IsUnset(request.UserOS)) {
		body["UserOS"] = request.UserOS
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DownloadGatewayProgram"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadGatewayProgramResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DownloadGatewayProgramRequest
//
// @return DownloadGatewayProgramResponse
func (client *Client) DownloadGatewayProgram(request *DownloadGatewayProgramRequest) (_result *DownloadGatewayProgramResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadGatewayProgramResponse{}
	_body, _err := client.DownloadGatewayProgramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据GatewayId查找网关详情
//
// @param request - FindUserGatewayByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindUserGatewayByIdResponse
func (client *Client) FindUserGatewayByIdWithOptions(request *FindUserGatewayByIdRequest, runtime *util.RuntimeOptions) (_result *FindUserGatewayByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FindUserGatewayById"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindUserGatewayByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据GatewayId查找网关详情
//
// @param request - FindUserGatewayByIdRequest
//
// @return FindUserGatewayByIdResponse
func (client *Client) FindUserGatewayById(request *FindUserGatewayByIdRequest) (_result *FindUserGatewayByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindUserGatewayByIdResponse{}
	_body, _err := client.FindUserGatewayByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetUserDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserDatabasesResponse
func (client *Client) GetUserDatabasesWithOptions(request *GetUserDatabasesRequest, runtime *util.RuntimeOptions) (_result *GetUserDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		body["DbType"] = request.DbType
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		body["Host"] = request.Host
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

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserDatabases"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetUserDatabasesRequest
//
// @return GetUserDatabasesResponse
func (client *Client) GetUserDatabases(request *GetUserDatabasesRequest) (_result *GetUserDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserDatabasesResponse{}
	_body, _err := client.GetUserDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetUserGatewayInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGatewayInstancesResponse
func (client *Client) GetUserGatewayInstancesWithOptions(request *GetUserGatewayInstancesRequest, runtime *util.RuntimeOptions) (_result *GetUserGatewayInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserGatewayInstances"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserGatewayInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetUserGatewayInstancesRequest
//
// @return GetUserGatewayInstancesResponse
func (client *Client) GetUserGatewayInstances(request *GetUserGatewayInstancesRequest) (_result *GetUserGatewayInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserGatewayInstancesResponse{}
	_body, _err := client.GetUserGatewayInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户DG网关列表
//
// @param request - GetUserGatewaysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGatewaysResponse
func (client *Client) GetUserGatewaysWithOptions(request *GetUserGatewaysRequest, runtime *util.RuntimeOptions) (_result *GetUserGatewaysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserGateways"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserGatewaysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户DG网关列表
//
// @param request - GetUserGatewaysRequest
//
// @return GetUserGatewaysResponse
func (client *Client) GetUserGateways(request *GetUserGatewaysRequest) (_result *GetUserGatewaysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserGatewaysResponse{}
	_body, _err := client.GetUserGatewaysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListDatabaseAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabaseAccessPointResponse
func (client *Client) ListDatabaseAccessPointWithOptions(request *ListDatabaseAccessPointRequest, runtime *util.RuntimeOptions) (_result *ListDatabaseAccessPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		body["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		body["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDatabaseAccessPoint"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDatabaseAccessPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListDatabaseAccessPointRequest
//
// @return ListDatabaseAccessPointResponse
func (client *Client) ListDatabaseAccessPoint(request *ListDatabaseAccessPointRequest) (_result *ListDatabaseAccessPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDatabaseAccessPointResponse{}
	_body, _err := client.ListDatabaseAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatabaseResponse
func (client *Client) ModifyDatabaseWithOptions(request *ModifyDatabaseRequest, runtime *util.RuntimeOptions) (_result *ModifyDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbDescription)) {
		body["DbDescription"] = request.DbDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DbPassword)) {
		body["DbPassword"] = request.DbPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		body["DbType"] = request.DbType
	}

	if !tea.BoolValue(util.IsUnset(request.DbUserName)) {
		body["DbUserName"] = request.DbUserName
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		body["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDatabase"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyDatabaseRequest
//
// @return ModifyDatabaseResponse
func (client *Client) ModifyDatabase(request *ModifyDatabaseRequest) (_result *ModifyDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDatabaseResponse{}
	_body, _err := client.ModifyDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改网关信息
//
// @param request - ModifyGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGatewayResponse
func (client *Client) ModifyGatewayWithOptions(request *ModifyGatewayRequest, runtime *util.RuntimeOptions) (_result *ModifyGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayDesc)) {
		body["GatewayDesc"] = request.GatewayDesc
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayName)) {
		body["GatewayName"] = request.GatewayName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyGateway"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改网关信息
//
// @param request - ModifyGatewayRequest
//
// @return ModifyGatewayResponse
func (client *Client) ModifyGateway(request *ModifyGatewayRequest) (_result *ModifyGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGatewayResponse{}
	_body, _err := client.ModifyGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RetryDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryDatabaseResponse
func (client *Client) RetryDatabaseWithOptions(request *RetryDatabaseRequest, runtime *util.RuntimeOptions) (_result *RetryDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DbDescription)) {
		body["DbDescription"] = request.DbDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DbPassword)) {
		body["DbPassword"] = request.DbPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		body["DbType"] = request.DbType
	}

	if !tea.BoolValue(util.IsUnset(request.DbUserName)) {
		body["DbUserName"] = request.DbUserName
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		body["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RetryDatabase"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RetryDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - RetryDatabaseRequest
//
// @return RetryDatabaseResponse
func (client *Client) RetryDatabase(request *RetryDatabaseRequest) (_result *RetryDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetryDatabaseResponse{}
	_body, _err := client.RetryDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止网关实例
//
// @param request - StopGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopGatewayResponse
func (client *Client) StopGatewayWithOptions(request *StopGatewayRequest, runtime *util.RuntimeOptions) (_result *StopGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayInstanceId)) {
		body["GatewayInstanceId"] = request.GatewayInstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopGateway"),
		Version:     tea.String("2023-09-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止网关实例
//
// @param request - StopGatewayRequest
//
// @return StopGatewayResponse
func (client *Client) StopGateway(request *StopGatewayRequest) (_result *StopGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopGatewayResponse{}
	_body, _err := client.StopGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
