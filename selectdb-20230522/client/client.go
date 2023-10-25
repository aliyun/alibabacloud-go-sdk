// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AllocateInstancePublicConnectionRequest struct {
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NetType                *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetDBInstanceId(v string) *AllocateInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetNetType(v string) *AllocateInstancePublicConnectionRequest {
	s.NetType = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetRegionId(v string) *AllocateInstancePublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

type AllocateInstancePublicConnectionResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AllocateInstancePublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionResponseBody) SetCode(v string) *AllocateInstancePublicConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponseBody) SetInstanceName(v string) *AllocateInstancePublicConnectionResponseBody {
	s.InstanceName = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponseBody) SetMessage(v string) *AllocateInstancePublicConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponseBody) SetRequestId(v string) *AllocateInstancePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponseBody) SetSuccess(v bool) *AllocateInstancePublicConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponseBody) SetTaskId(v int64) *AllocateInstancePublicConnectionResponseBody {
	s.TaskId = &v
	return s
}

type AllocateInstancePublicConnectionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AllocateInstancePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocateInstancePublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionResponse) SetHeaders(v map[string]*string) *AllocateInstancePublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocateInstancePublicConnectionResponse) SetStatusCode(v int32) *AllocateInstancePublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponse) SetBody(v *AllocateInstancePublicConnectionResponseBody) *AllocateInstancePublicConnectionResponse {
	s.Body = v
	return s
}

type CheckCreateDBInstanceRequest struct {
	CacheSize             *int32  `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	ChargeType            *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionString      *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	DBInstanceClass       *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	Engine                *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	Period                *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityIPList        *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	StorageSize           *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	UsedTime              *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	VpcId  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CheckCreateDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckCreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CheckCreateDBInstanceRequest) SetCacheSize(v int32) *CheckCreateDBInstanceRequest {
	s.CacheSize = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetChargeType(v string) *CheckCreateDBInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetClientToken(v string) *CheckCreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetConnectionString(v string) *CheckCreateDBInstanceRequest {
	s.ConnectionString = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetDBInstanceClass(v string) *CheckCreateDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetDBInstanceDescription(v string) *CheckCreateDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetEngine(v string) *CheckCreateDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetEngineVersion(v string) *CheckCreateDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetPeriod(v string) *CheckCreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetRegionId(v string) *CheckCreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetResourceGroupId(v string) *CheckCreateDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetResourceOwnerId(v int64) *CheckCreateDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetSecurityIPList(v string) *CheckCreateDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetStorageSize(v string) *CheckCreateDBInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetUsedTime(v int32) *CheckCreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetVSwitchId(v string) *CheckCreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetVpcId(v string) *CheckCreateDBInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetZoneId(v string) *CheckCreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

type CheckCreateDBInstanceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckCreateDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckCreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCreateDBInstanceResponseBody) SetCode(v string) *CheckCreateDBInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CheckCreateDBInstanceResponseBody) SetMessage(v string) *CheckCreateDBInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CheckCreateDBInstanceResponseBody) SetRequestId(v string) *CheckCreateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCreateDBInstanceResponseBody) SetSuccess(v string) *CheckCreateDBInstanceResponseBody {
	s.Success = &v
	return s
}

type CheckCreateDBInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckCreateDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckCreateDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckCreateDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CheckCreateDBInstanceResponse) SetHeaders(v map[string]*string) *CheckCreateDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CheckCreateDBInstanceResponse) SetStatusCode(v int32) *CheckCreateDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCreateDBInstanceResponse) SetBody(v *CheckCreateDBInstanceResponseBody) *CheckCreateDBInstanceResponse {
	s.Body = v
	return s
}

type CheckServiceLinkedRoleRequest struct {
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleRequest) SetOwnerAccount(v string) *CheckServiceLinkedRoleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) SetResourceOwnerId(v int64) *CheckServiceLinkedRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

type CheckServiceLinkedRoleResponseBody struct {
	HasServiceLinkedRole *bool   `json:"HasServiceLinkedRole,omitempty" xml:"HasServiceLinkedRole,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponseBody) SetHasServiceLinkedRole(v bool) *CheckServiceLinkedRoleResponseBody {
	s.HasServiceLinkedRole = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type CheckServiceLinkedRoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CheckServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceLinkedRoleResponse) SetStatusCode(v int32) *CheckServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceLinkedRoleResponse) SetBody(v *CheckServiceLinkedRoleResponseBody) *CheckServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type CreateDBClusterRequest struct {
	CacheSize            *string `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	ChargeType           *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DBClusterClass       *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// 代表资源一级ID的资源属性字段
	DBInstanceId  *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Engine        *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	Period        *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 代表资源组的资源属性字段
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StorageSize     *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	UsedTime        *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VSwitchId       *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	VpcId  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterRequest) SetCacheSize(v string) *CreateDBClusterRequest {
	s.CacheSize = &v
	return s
}

func (s *CreateDBClusterRequest) SetChargeType(v string) *CreateDBClusterRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterClass(v string) *CreateDBClusterRequest {
	s.DBClusterClass = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterDescription(v string) *CreateDBClusterRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBInstanceId(v string) *CreateDBClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBClusterRequest) SetEngine(v string) *CreateDBClusterRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBClusterRequest) SetEngineVersion(v string) *CreateDBClusterRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBClusterRequest) SetPeriod(v string) *CreateDBClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateDBClusterRequest) SetRegionId(v string) *CreateDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceGroupId(v string) *CreateDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceOwnerId(v int64) *CreateDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageSize(v string) *CreateDBClusterRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateDBClusterRequest) SetUsedTime(v string) *CreateDBClusterRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetVSwitchId(v string) *CreateDBClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBClusterRequest) SetVpcId(v string) *CreateDBClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBClusterRequest) SetZoneId(v string) *CreateDBClusterRequest {
	s.ZoneId = &v
	return s
}

type CreateDBClusterResponseBody struct {
	Data      *CreateDBClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponseBody) SetData(v *CreateDBClusterResponseBodyData) *CreateDBClusterResponseBody {
	s.Data = v
	return s
}

func (s *CreateDBClusterResponseBody) SetRequestId(v string) *CreateDBClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateDBClusterResponseBodyData struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId      *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDBClusterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponseBodyData) SetDBInstanceId(v string) *CreateDBClusterResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBClusterResponseBodyData) SetOrderId(v int64) *CreateDBClusterResponseBodyData {
	s.OrderId = &v
	return s
}

type CreateDBClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponse) SetHeaders(v map[string]*string) *CreateDBClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateDBClusterResponse) SetStatusCode(v int32) *CreateDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBClusterResponse) SetBody(v *CreateDBClusterResponseBody) *CreateDBClusterResponse {
	s.Body = v
	return s
}

type CreateDBInstanceRequest struct {
	CacheSize             *int32  `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	ChargeType            *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionString      *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	DBInstanceClass       *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	Engine                *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	Period                *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 代表资源组的资源属性字段
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityIPList  *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	StorageSize     *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	UsedTime        *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VSwitchId       *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	VpcId  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) SetCacheSize(v int32) *CreateDBInstanceRequest {
	s.CacheSize = &v
	return s
}

func (s *CreateDBInstanceRequest) SetChargeType(v string) *CreateDBInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetConnectionString(v string) *CreateDBInstanceRequest {
	s.ConnectionString = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceClass(v string) *CreateDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceDescription(v string) *CreateDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngine(v string) *CreateDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngineVersion(v string) *CreateDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v string) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceGroupId(v string) *CreateDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceOwnerId(v int64) *CreateDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSecurityIPList(v string) *CreateDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageSize(v string) *CreateDBInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateDBInstanceRequest) SetUsedTime(v int32) *CreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVpcId(v string) *CreateDBInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

type CreateDBInstanceResponseBody struct {
	Data      *CreateDBInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBody) SetData(v *CreateDBInstanceResponseBodyData) *CreateDBInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateDBInstanceResponseBody) SetRequestId(v string) *CreateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateDBInstanceResponseBodyData struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId      *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDBInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBodyData) SetDBInstanceId(v string) *CreateDBInstanceResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceResponseBodyData) SetOrderId(v int64) *CreateDBInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

type CreateDBInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponse) SetHeaders(v map[string]*string) *CreateDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDBInstanceResponse) SetStatusCode(v int32) *CreateDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBInstanceResponse) SetBody(v *CreateDBInstanceResponseBody) *CreateDBInstanceResponse {
	s.Body = v
	return s
}

type CreateServiceLinkedRoleForSelectDBRequest struct {
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateServiceLinkedRoleForSelectDBRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleForSelectDBRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForSelectDBRequest) SetOwnerAccount(v string) *CreateServiceLinkedRoleForSelectDBRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateServiceLinkedRoleForSelectDBRequest) SetResourceOwnerId(v int64) *CreateServiceLinkedRoleForSelectDBRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateServiceLinkedRoleForSelectDBResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceLinkedRoleForSelectDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleForSelectDBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForSelectDBResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleForSelectDBResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceLinkedRoleForSelectDBResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceLinkedRoleForSelectDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleForSelectDBResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleForSelectDBResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForSelectDBResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleForSelectDBResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleForSelectDBResponse) SetStatusCode(v int32) *CreateServiceLinkedRoleForSelectDBResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceLinkedRoleForSelectDBResponse) SetBody(v *CreateServiceLinkedRoleForSelectDBResponseBody) *CreateServiceLinkedRoleForSelectDBResponse {
	s.Body = v
	return s
}

type DeleteDBClusterRequest struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// 代表资源一级ID的资源属性字段
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 代表资源组的资源属性字段
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterRequest) SetDBClusterId(v string) *DeleteDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetDBInstanceId(v string) *DeleteDBClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetRegionId(v string) *DeleteDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetResourceGroupId(v string) *DeleteDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetResourceOwnerId(v int64) *DeleteDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteDBClusterResponseBody struct {
	Data      *DeleteDBClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponseBody) SetData(v *DeleteDBClusterResponseBodyData) *DeleteDBClusterResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDBClusterResponseBody) SetRequestId(v string) *DeleteDBClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBClusterResponseBodyData struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s DeleteDBClusterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponseBodyData) SetDBInstanceId(v string) *DeleteDBClusterResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBClusterResponseBodyData) SetOrderId(v string) *DeleteDBClusterResponseBodyData {
	s.OrderId = &v
	return s
}

type DeleteDBClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponse) SetHeaders(v map[string]*string) *DeleteDBClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBClusterResponse) SetStatusCode(v int32) *DeleteDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBClusterResponse) SetBody(v *DeleteDBClusterResponseBody) *DeleteDBClusterResponse {
	s.Body = v
	return s
}

type DeleteDBInstanceRequest struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceRequest) SetDBInstanceId(v string) *DeleteDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetRegionId(v string) *DeleteDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetResourceOwnerId(v int64) *DeleteDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteDBInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceResponseBody) SetRequestId(v string) *DeleteDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceResponse) SetHeaders(v map[string]*string) *DeleteDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBInstanceResponse) SetStatusCode(v int32) *DeleteDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBInstanceResponse) SetBody(v *DeleteDBInstanceResponseBody) *DeleteDBInstanceResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceAttributeRequest struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeRequest) SetDBInstanceId(v string) *DescribeDBInstanceAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetRegionId(v string) *DescribeDBInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBody struct {
	CanUpgradeVersions []*string                                               `json:"CanUpgradeVersions,omitempty" xml:"CanUpgradeVersions,omitempty" type:"Repeated"`
	ChargeType         *string                                                 `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime         *string                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBClusterList      []*DescribeDBInstanceAttributeResponseBodyDBClusterList `json:"DBClusterList,omitempty" xml:"DBClusterList,omitempty" type:"Repeated"`
	DBInstanceId       *string                                                 `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Description        *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	Engine             *string                                                 `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineMinorVersion *string                                                 `json:"EngineMinorVersion,omitempty" xml:"EngineMinorVersion,omitempty"`
	EngineVersion      *string                                                 `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime         *string                                                 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GmtModified        *string                                                 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LockMode           *int64                                                  `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason         *string                                                 `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MaintainEndtime    *string                                                 `json:"MaintainEndtime,omitempty" xml:"MaintainEndtime,omitempty"`
	MaintainStarttime  *string                                                 `json:"MaintainStarttime,omitempty" xml:"MaintainStarttime,omitempty"`
	ObjectStoreSize    *int64                                                  `json:"ObjectStoreSize,omitempty" xml:"ObjectStoreSize,omitempty"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceCpu        *int64                                                  `json:"ResourceCpu,omitempty" xml:"ResourceCpu,omitempty"`
	Status             *string                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageSize        *int64                                                  `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	SubDomain          *string                                                 `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBody) SetCanUpgradeVersions(v []*string) *DescribeDBInstanceAttributeResponseBody {
	s.CanUpgradeVersions = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetChargeType(v string) *DescribeDBInstanceAttributeResponseBody {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetCreateTime(v string) *DescribeDBInstanceAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetDBClusterList(v []*DescribeDBInstanceAttributeResponseBodyDBClusterList) *DescribeDBInstanceAttributeResponseBody {
	s.DBClusterList = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetDescription(v string) *DescribeDBInstanceAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetEngine(v string) *DescribeDBInstanceAttributeResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetEngineMinorVersion(v string) *DescribeDBInstanceAttributeResponseBody {
	s.EngineMinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetEngineVersion(v string) *DescribeDBInstanceAttributeResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetExpireTime(v string) *DescribeDBInstanceAttributeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetGmtModified(v string) *DescribeDBInstanceAttributeResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetLockMode(v int64) *DescribeDBInstanceAttributeResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetLockReason(v string) *DescribeDBInstanceAttributeResponseBody {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetMaintainEndtime(v string) *DescribeDBInstanceAttributeResponseBody {
	s.MaintainEndtime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetMaintainStarttime(v string) *DescribeDBInstanceAttributeResponseBody {
	s.MaintainStarttime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetObjectStoreSize(v int64) *DescribeDBInstanceAttributeResponseBody {
	s.ObjectStoreSize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetResourceCpu(v int64) *DescribeDBInstanceAttributeResponseBody {
	s.ResourceCpu = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetStatus(v string) *DescribeDBInstanceAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetStorageSize(v int64) *DescribeDBInstanceAttributeResponseBody {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetSubDomain(v string) *DescribeDBInstanceAttributeResponseBody {
	s.SubDomain = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBClusterList struct {
	CacheStorageSizeGB  *string `json:"CacheStorageSizeGB,omitempty" xml:"CacheStorageSizeGB,omitempty"`
	CacheStorageSizeGiB *int64  `json:"CacheStorageSizeGiB,omitempty" xml:"CacheStorageSizeGiB,omitempty"`
	CacheStorageType    *string `json:"CacheStorageType,omitempty" xml:"CacheStorageType,omitempty"`
	ChargeType          *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CpuCores            *int64  `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	CreatedTime         *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DbClusterClass      *string `json:"DbClusterClass,omitempty" xml:"DbClusterClass,omitempty"`
	DbClusterId         *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	DbClusterName       *string `json:"DbClusterName,omitempty" xml:"DbClusterName,omitempty"`
	DbInstanceName      *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	Memory              *int64  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	ObjectStoreSizeGiB  *int64  `json:"ObjectStoreSizeGiB,omitempty" xml:"ObjectStoreSizeGiB,omitempty"`
	PerformanceLevel    *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBClusterList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBClusterList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetCacheStorageSizeGB(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.CacheStorageSizeGB = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetCacheStorageSizeGiB(v int64) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.CacheStorageSizeGiB = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetCacheStorageType(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.CacheStorageType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetChargeType(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetCpuCores(v int64) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetCreatedTime(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetDbClusterClass(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.DbClusterClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetDbClusterId(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.DbClusterId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetDbClusterName(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.DbClusterName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetDbInstanceName(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetMemory(v int64) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.Memory = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetObjectStoreSizeGiB(v int64) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.ObjectStoreSizeGiB = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetPerformanceLevel(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetStartTime(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.Status = &v
	return s
}

type DescribeDBInstanceAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceAttributeResponse) SetStatusCode(v int32) *DescribeDBInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponse) SetBody(v *DescribeDBInstanceAttributeResponseBody) *DescribeDBInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceNetInfoRequest struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceNetInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoRequest) SetDBInstanceId(v string) *DescribeDBInstanceNetInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetRegionId(v string) *DescribeDBInstanceNetInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceNetInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBInstanceNetInfoResponseBody struct {
	DBInstanceNetInfos []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos `json:"DBInstanceNetInfos,omitempty" xml:"DBInstanceNetInfos,omitempty" type:"Repeated"`
	RequestId          *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetDBInstanceNetInfos(v []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) *DescribeDBInstanceNetInfoResponseBody {
	s.DBInstanceNetInfos = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos struct {
	ConnectionString *string                                                            `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	Ip               *string                                                            `json:"Ip,omitempty" xml:"Ip,omitempty"`
	NetType          *string                                                            `json:"NetType,omitempty" xml:"NetType,omitempty"`
	PortList         []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList `json:"PortList,omitempty" xml:"PortList,omitempty" type:"Repeated"`
	UserVisible      *bool                                                              `json:"UserVisible,omitempty" xml:"UserVisible,omitempty"`
	// VPC ID。
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
	VswitchId     *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetConnectionString(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetIp(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.Ip = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetNetType(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.NetType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetPortList(v []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.PortList = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetUserVisible(v bool) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.UserVisible = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetVpcId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetVpcInstanceId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetVswitchId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.VswitchId = &v
	return s
}

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList) SetPort(v int32) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList) SetProtocol(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList {
	s.Protocol = &v
	return s
}

type DescribeDBInstanceNetInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceNetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceNetInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceNetInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponse) SetStatusCode(v int32) *DescribeDBInstanceNetInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponse) SetBody(v *DescribeDBInstanceNetInfoResponseBody) *DescribeDBInstanceNetInfoResponse {
	s.Body = v
	return s
}

type DescribeDBInstancesRequest struct {
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceIds         *string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	DBInstanceStatus      *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	PageNumber            *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) SetDBInstanceDescription(v string) *DescribeDBInstancesRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceIds(v string) *DescribeDBInstancesRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceStatus(v string) *DescribeDBInstancesRequest {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageNumber(v int64) *DescribeDBInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageSize(v int64) *DescribeDBInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetRegionId(v string) *DescribeDBInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceGroupId(v string) *DescribeDBInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceOwnerId(v int64) *DescribeDBInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBInstancesResponseBody struct {
	Items            []*DescribeDBInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber       *int64                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int64                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int64                                  `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBody) SetItems(v []*DescribeDBInstancesResponseBodyItems) *DescribeDBInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageNumber(v int64) *DescribeDBInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageSize(v int64) *DescribeDBInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetRequestId(v string) *DescribeDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetTotalRecordCount(v int64) *DescribeDBInstancesResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeDBInstancesResponseBodyItems struct {
	Category             *string                                              `json:"Category,omitempty" xml:"Category,omitempty"`
	ChargeType           *string                                              `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClusterCount         *int32                                               `json:"ClusterCount,omitempty" xml:"ClusterCount,omitempty"`
	CreateTime           *string                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBClusterList        []*DescribeDBInstancesResponseBodyItemsDBClusterList `json:"DBClusterList,omitempty" xml:"DBClusterList,omitempty" type:"Repeated"`
	DBInstanceId         *string                                              `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Description          *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Engine               *string                                              `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion        *string                                              `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime           *string                                              `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GmtCreated           *string                                              `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified          *string                                              `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InstanceClass        *string                                              `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	InstanceUsedType     *string                                              `json:"InstanceUsedType,omitempty" xml:"InstanceUsedType,omitempty"`
	IsDeleted            *bool                                                `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	LockMode             *int64                                               `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason           *string                                              `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MaintainEndTimeStr   *string                                              `json:"MaintainEndTimeStr,omitempty" xml:"MaintainEndTimeStr,omitempty"`
	MaintainEndtime      *string                                              `json:"MaintainEndtime,omitempty" xml:"MaintainEndtime,omitempty"`
	MaintainStartTimeStr *string                                              `json:"MaintainStartTimeStr,omitempty" xml:"MaintainStartTimeStr,omitempty"`
	MaintainStarttime    *string                                              `json:"MaintainStarttime,omitempty" xml:"MaintainStarttime,omitempty"`
	ObjectStoreSize      *int64                                               `json:"ObjectStoreSize,omitempty" xml:"ObjectStoreSize,omitempty"`
	ParentInstance       *string                                              `json:"ParentInstance,omitempty" xml:"ParentInstance,omitempty"`
	RegionId             *string                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceCpu          *int64                                               `json:"ResourceCpu,omitempty" xml:"ResourceCpu,omitempty"`
	ResourceGroupId      *string                                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceMemory       *int64                                               `json:"ResourceMemory,omitempty" xml:"ResourceMemory,omitempty"`
	ScaleMax             *int64                                               `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	ScaleMin             *int64                                               `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	ScaleReplica         *int64                                               `json:"ScaleReplica,omitempty" xml:"ScaleReplica,omitempty"`
	Status               *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageSize          *int64                                               `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	StorageType          *string                                              `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tags                 []*DescribeDBInstancesResponseBodyItemsTags          `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TenantClusterId      *string                                              `json:"TenantClusterId,omitempty" xml:"TenantClusterId,omitempty"`
	TenantToken          *string                                              `json:"TenantToken,omitempty" xml:"TenantToken,omitempty"`
	TenantUserId         *string                                              `json:"TenantUserId,omitempty" xml:"TenantUserId,omitempty"`
	// VPC ID。
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId        *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ZoneId           *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ConnectionString *string `json:"connectionString,omitempty" xml:"connectionString,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItems) SetCategory(v string) *DescribeDBInstancesResponseBodyItems {
	s.Category = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetChargeType(v string) *DescribeDBInstancesResponseBodyItems {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetClusterCount(v int32) *DescribeDBInstancesResponseBodyItems {
	s.ClusterCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetCreateTime(v string) *DescribeDBInstancesResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetDBClusterList(v []*DescribeDBInstancesResponseBodyItemsDBClusterList) *DescribeDBInstancesResponseBodyItems {
	s.DBClusterList = v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetDBInstanceId(v string) *DescribeDBInstancesResponseBodyItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetDescription(v string) *DescribeDBInstancesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetEngine(v string) *DescribeDBInstancesResponseBodyItems {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetEngineVersion(v string) *DescribeDBInstancesResponseBodyItems {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetExpireTime(v string) *DescribeDBInstancesResponseBodyItems {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetGmtCreated(v string) *DescribeDBInstancesResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetGmtModified(v string) *DescribeDBInstancesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetInstanceClass(v string) *DescribeDBInstancesResponseBodyItems {
	s.InstanceClass = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetInstanceUsedType(v string) *DescribeDBInstancesResponseBodyItems {
	s.InstanceUsedType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetIsDeleted(v bool) *DescribeDBInstancesResponseBodyItems {
	s.IsDeleted = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetLockMode(v int64) *DescribeDBInstancesResponseBodyItems {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetLockReason(v string) *DescribeDBInstancesResponseBodyItems {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetMaintainEndTimeStr(v string) *DescribeDBInstancesResponseBodyItems {
	s.MaintainEndTimeStr = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetMaintainEndtime(v string) *DescribeDBInstancesResponseBodyItems {
	s.MaintainEndtime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetMaintainStartTimeStr(v string) *DescribeDBInstancesResponseBodyItems {
	s.MaintainStartTimeStr = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetMaintainStarttime(v string) *DescribeDBInstancesResponseBodyItems {
	s.MaintainStarttime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetObjectStoreSize(v int64) *DescribeDBInstancesResponseBodyItems {
	s.ObjectStoreSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetParentInstance(v string) *DescribeDBInstancesResponseBodyItems {
	s.ParentInstance = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetRegionId(v string) *DescribeDBInstancesResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetResourceCpu(v int64) *DescribeDBInstancesResponseBodyItems {
	s.ResourceCpu = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetResourceGroupId(v string) *DescribeDBInstancesResponseBodyItems {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetResourceMemory(v int64) *DescribeDBInstancesResponseBodyItems {
	s.ResourceMemory = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetScaleMax(v int64) *DescribeDBInstancesResponseBodyItems {
	s.ScaleMax = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetScaleMin(v int64) *DescribeDBInstancesResponseBodyItems {
	s.ScaleMin = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetScaleReplica(v int64) *DescribeDBInstancesResponseBodyItems {
	s.ScaleReplica = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetStatus(v string) *DescribeDBInstancesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetStorageSize(v int64) *DescribeDBInstancesResponseBodyItems {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetStorageType(v string) *DescribeDBInstancesResponseBodyItems {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetTags(v []*DescribeDBInstancesResponseBodyItemsTags) *DescribeDBInstancesResponseBodyItems {
	s.Tags = v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetTenantClusterId(v string) *DescribeDBInstancesResponseBodyItems {
	s.TenantClusterId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetTenantToken(v string) *DescribeDBInstancesResponseBodyItems {
	s.TenantToken = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetTenantUserId(v string) *DescribeDBInstancesResponseBodyItems {
	s.TenantUserId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetVpcId(v string) *DescribeDBInstancesResponseBodyItems {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetVswitchId(v string) *DescribeDBInstancesResponseBodyItems {
	s.VswitchId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetZoneId(v string) *DescribeDBInstancesResponseBodyItems {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetConnectionString(v string) *DescribeDBInstancesResponseBodyItems {
	s.ConnectionString = &v
	return s
}

type DescribeDBInstancesResponseBodyItemsDBClusterList struct {
	CacheStorageSizeGiB *int64  `json:"CacheStorageSizeGiB,omitempty" xml:"CacheStorageSizeGiB,omitempty"`
	CacheStorageType    *string `json:"CacheStorageType,omitempty" xml:"CacheStorageType,omitempty"`
	CpuCores            *int64  `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	CreatedTime         *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DBInstanceId        *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DbClusterClass      *string `json:"DbClusterClass,omitempty" xml:"DbClusterClass,omitempty"`
	DbClusterId         *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	DbClusterName       *string `json:"DbClusterName,omitempty" xml:"DbClusterName,omitempty"`
	Memory              *int64  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	PerformanceLevel    *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ObjectStoreSizeGiB  *int64  `json:"objectStoreSizeGiB,omitempty" xml:"objectStoreSizeGiB,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItemsDBClusterList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBClusterList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBClusterList) SetCacheStorageSizeGiB(v int64) *DescribeDBInstancesResponseBodyItemsDBClusterList {
	s.CacheStorageSizeGiB = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBClusterList) SetCacheStorageType(v string) *DescribeDBInstancesResponseBodyItemsDBClusterList {
	s.CacheStorageType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBClusterList) SetCpuCores(v int64) *DescribeDBInstancesResponseBodyItemsDBClusterList {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBClusterList) SetCreatedTime(v string) *DescribeDBInstancesResponseBodyItemsDBClusterList {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBClusterList) SetDBInstanceId(v string) *DescribeDBInstancesResponseBodyItemsDBClusterList {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBClusterList) SetDbClusterClass(v string) *DescribeDBInstancesResponseBodyItemsDBClusterList {
	s.DbClusterClass = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBClusterList) SetDbClusterId(v string) *DescribeDBInstancesResponseBodyItemsDBClusterList {
	s.DbClusterId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBClusterList) SetDbClusterName(v string) *DescribeDBInstancesResponseBodyItemsDBClusterList {
	s.DbClusterName = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBClusterList) SetMemory(v int64) *DescribeDBInstancesResponseBodyItemsDBClusterList {
	s.Memory = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBClusterList) SetPerformanceLevel(v string) *DescribeDBInstancesResponseBodyItemsDBClusterList {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBClusterList) SetStatus(v string) *DescribeDBInstancesResponseBodyItemsDBClusterList {
	s.Status = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBClusterList) SetObjectStoreSizeGiB(v int64) *DescribeDBInstancesResponseBodyItemsDBClusterList {
	s.ObjectStoreSizeGiB = &v
	return s
}

type DescribeDBInstancesResponseBodyItemsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItemsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsTags) SetKey(v string) *DescribeDBInstancesResponseBodyItemsTags {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsTags) SetValue(v string) *DescribeDBInstancesResponseBodyItemsTags {
	s.Value = &v
	return s
}

type DescribeDBInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponse) SetHeaders(v map[string]*string) *DescribeDBInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancesResponse) SetStatusCode(v int32) *DescribeDBInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancesResponse) SetBody(v *DescribeDBInstancesResponseBody) *DescribeDBInstancesResponse {
	s.Body = v
	return s
}

type DescribeSecurityIPListRequest struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSecurityIPListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIPListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListRequest) SetDBInstanceId(v string) *DescribeSecurityIPListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSecurityIPListRequest) SetRegionId(v string) *DescribeSecurityIPListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityIPListRequest) SetResourceOwnerId(v int64) *DescribeSecurityIPListRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeSecurityIPListResponseBody struct {
	Code           *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	DBInstanceName *string                                         `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	GroupItems     []*DescribeSecurityIPListResponseBodyGroupItems `json:"GroupItems,omitempty" xml:"GroupItems,omitempty" type:"Repeated"`
	Message        *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSecurityIPListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIPListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListResponseBody) SetCode(v string) *DescribeSecurityIPListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSecurityIPListResponseBody) SetDBInstanceName(v string) *DescribeSecurityIPListResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSecurityIPListResponseBody) SetGroupItems(v []*DescribeSecurityIPListResponseBodyGroupItems) *DescribeSecurityIPListResponseBody {
	s.GroupItems = v
	return s
}

func (s *DescribeSecurityIPListResponseBody) SetMessage(v string) *DescribeSecurityIPListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSecurityIPListResponseBody) SetRequestId(v string) *DescribeSecurityIPListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityIPListResponseBody) SetSuccess(v bool) *DescribeSecurityIPListResponseBody {
	s.Success = &v
	return s
}

type DescribeSecurityIPListResponseBodyGroupItems struct {
	AecurityIPType   *string `json:"AecurityIPType,omitempty" xml:"AecurityIPType,omitempty"`
	GroupName        *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupTag         *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	SecurityIPList   *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	WhitelistNetType *string `json:"WhitelistNetType,omitempty" xml:"WhitelistNetType,omitempty"`
}

func (s DescribeSecurityIPListResponseBodyGroupItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIPListResponseBodyGroupItems) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) SetAecurityIPType(v string) *DescribeSecurityIPListResponseBodyGroupItems {
	s.AecurityIPType = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) SetGroupName(v string) *DescribeSecurityIPListResponseBodyGroupItems {
	s.GroupName = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) SetGroupTag(v string) *DescribeSecurityIPListResponseBodyGroupItems {
	s.GroupTag = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) SetSecurityIPList(v string) *DescribeSecurityIPListResponseBodyGroupItems {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) SetWhitelistNetType(v string) *DescribeSecurityIPListResponseBodyGroupItems {
	s.WhitelistNetType = &v
	return s
}

type DescribeSecurityIPListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSecurityIPListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityIPListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIPListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListResponse) SetHeaders(v map[string]*string) *DescribeSecurityIPListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityIPListResponse) SetStatusCode(v int32) *DescribeSecurityIPListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityIPListResponse) SetBody(v *DescribeSecurityIPListResponseBody) *DescribeSecurityIPListResponse {
	s.Body = v
	return s
}

type ModifyBEClusterAttributeRequest struct {
	DBClusterId           *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	InstanceAttributeType *string `json:"InstanceAttributeType,omitempty" xml:"InstanceAttributeType,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Value                 *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyBEClusterAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBEClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyBEClusterAttributeRequest) SetDBClusterId(v string) *ModifyBEClusterAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyBEClusterAttributeRequest) SetDBInstanceId(v string) *ModifyBEClusterAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyBEClusterAttributeRequest) SetInstanceAttributeType(v string) *ModifyBEClusterAttributeRequest {
	s.InstanceAttributeType = &v
	return s
}

func (s *ModifyBEClusterAttributeRequest) SetRegionId(v string) *ModifyBEClusterAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBEClusterAttributeRequest) SetResourceOwnerId(v int64) *ModifyBEClusterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBEClusterAttributeRequest) SetValue(v string) *ModifyBEClusterAttributeRequest {
	s.Value = &v
	return s
}

type ModifyBEClusterAttributeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyBEClusterAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBEClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBEClusterAttributeResponseBody) SetCode(v string) *ModifyBEClusterAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyBEClusterAttributeResponseBody) SetMessage(v string) *ModifyBEClusterAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyBEClusterAttributeResponseBody) SetRequestId(v string) *ModifyBEClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBEClusterAttributeResponseBody) SetSuccess(v string) *ModifyBEClusterAttributeResponseBody {
	s.Success = &v
	return s
}

type ModifyBEClusterAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyBEClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBEClusterAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBEClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyBEClusterAttributeResponse) SetHeaders(v map[string]*string) *ModifyBEClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyBEClusterAttributeResponse) SetStatusCode(v int32) *ModifyBEClusterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBEClusterAttributeResponse) SetBody(v *ModifyBEClusterAttributeResponseBody) *ModifyBEClusterAttributeResponse {
	s.Body = v
	return s
}

type ModifyDBClusterRequest struct {
	DBClusterClass  *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Engine          *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterRequest) SetDBClusterClass(v string) *ModifyDBClusterRequest {
	s.DBClusterClass = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterId(v string) *ModifyDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBInstanceId(v string) *ModifyDBClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetEngine(v string) *ModifyDBClusterRequest {
	s.Engine = &v
	return s
}

func (s *ModifyDBClusterRequest) SetRegionId(v string) *ModifyDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetResourceOwnerId(v int64) *ModifyDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyDBClusterResponseBody struct {
	// 代表资源一级ID的资源属性字段
	DBInstanceId *string                          `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Data         *ModifyDBClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponseBody) SetDBInstanceId(v string) *ModifyDBClusterResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) SetData(v *ModifyDBClusterResponseBodyData) *ModifyDBClusterResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDBClusterResponseBody) SetRequestId(v string) *ModifyDBClusterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterResponseBodyData struct {
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId      *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyDBClusterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponseBodyData) SetDBClusterId(v string) *ModifyDBClusterResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterResponseBodyData) SetDBInstanceId(v string) *ModifyDBClusterResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBClusterResponseBodyData) SetOrderId(v int64) *ModifyDBClusterResponseBodyData {
	s.OrderId = &v
	return s
}

type ModifyDBClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponse) SetHeaders(v map[string]*string) *ModifyDBClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterResponse) SetStatusCode(v int32) *ModifyDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterResponse) SetBody(v *ModifyDBClusterResponseBody) *ModifyDBClusterResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceAttributeRequest struct {
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	InstanceAttributeType *string `json:"InstanceAttributeType,omitempty" xml:"InstanceAttributeType,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Value                 *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyDBInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAttributeRequest) SetDBInstanceId(v string) *ModifyDBInstanceAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetInstanceAttributeType(v string) *ModifyDBInstanceAttributeRequest {
	s.InstanceAttributeType = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetRegionId(v string) *ModifyDBInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetValue(v string) *ModifyDBInstanceAttributeRequest {
	s.Value = &v
	return s
}

type ModifyDBInstanceAttributeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDBInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAttributeResponseBody) SetCode(v string) *ModifyDBInstanceAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDBInstanceAttributeResponseBody) SetMessage(v string) *ModifyDBInstanceAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDBInstanceAttributeResponseBody) SetRequestId(v string) *ModifyDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceAttributeResponseBody) SetSuccess(v bool) *ModifyDBInstanceAttributeResponseBody {
	s.Success = &v
	return s
}

type ModifyDBInstanceAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAttributeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceAttributeResponse) SetStatusCode(v int32) *ModifyDBInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceAttributeResponse) SetBody(v *ModifyDBInstanceAttributeResponseBody) *ModifyDBInstanceAttributeResponse {
	s.Body = v
	return s
}

type ModifySecurityIPListRequest struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ModifyMode      *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityIPList  *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s ModifySecurityIPListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIPListRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPListRequest) SetDBInstanceId(v string) *ModifySecurityIPListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySecurityIPListRequest) SetGroupName(v string) *ModifySecurityIPListRequest {
	s.GroupName = &v
	return s
}

func (s *ModifySecurityIPListRequest) SetModifyMode(v string) *ModifySecurityIPListRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifySecurityIPListRequest) SetRegionId(v string) *ModifySecurityIPListRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityIPListRequest) SetResourceOwnerId(v int64) *ModifySecurityIPListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityIPListRequest) SetSecurityIPList(v string) *ModifySecurityIPListRequest {
	s.SecurityIPList = &v
	return s
}

type ModifySecurityIPListResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DBInstanceName   *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	GroupName        *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupTag         *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityIPList   *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	SecurityIPType   *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId           *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	WhitelistNetType *string `json:"WhitelistNetType,omitempty" xml:"WhitelistNetType,omitempty"`
}

func (s ModifySecurityIPListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIPListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPListResponseBody) SetCode(v string) *ModifySecurityIPListResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetDBInstanceName(v string) *ModifySecurityIPListResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetGroupName(v string) *ModifySecurityIPListResponseBody {
	s.GroupName = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetGroupTag(v string) *ModifySecurityIPListResponseBody {
	s.GroupTag = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetMessage(v string) *ModifySecurityIPListResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetRequestId(v string) *ModifySecurityIPListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetSecurityIPList(v string) *ModifySecurityIPListResponseBody {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetSecurityIPType(v string) *ModifySecurityIPListResponseBody {
	s.SecurityIPType = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetSuccess(v bool) *ModifySecurityIPListResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetTaskId(v int64) *ModifySecurityIPListResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetWhitelistNetType(v string) *ModifySecurityIPListResponseBody {
	s.WhitelistNetType = &v
	return s
}

type ModifySecurityIPListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySecurityIPListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySecurityIPListResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIPListResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPListResponse) SetHeaders(v map[string]*string) *ModifySecurityIPListResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityIPListResponse) SetStatusCode(v int32) *ModifySecurityIPListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityIPListResponse) SetBody(v *ModifySecurityIPListResponseBody) *ModifySecurityIPListResponse {
	s.Body = v
	return s
}

type ReleaseInstancePublicConnectionRequest struct {
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId  *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleaseInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionRequest) SetConnectionString(v string) *ReleaseInstancePublicConnectionRequest {
	s.ConnectionString = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetDBInstanceId(v string) *ReleaseInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetRegionId(v string) *ReleaseInstancePublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *ReleaseInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

type ReleaseInstancePublicConnectionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseInstancePublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionResponseBody) SetCode(v string) *ReleaseInstancePublicConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseInstancePublicConnectionResponseBody) SetMessage(v string) *ReleaseInstancePublicConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseInstancePublicConnectionResponseBody) SetRequestId(v string) *ReleaseInstancePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionResponseBody) SetSuccess(v bool) *ReleaseInstancePublicConnectionResponseBody {
	s.Success = &v
	return s
}

type ReleaseInstancePublicConnectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseInstancePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseInstancePublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionResponse) SetHeaders(v map[string]*string) *ReleaseInstancePublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstancePublicConnectionResponse) SetStatusCode(v int32) *ReleaseInstancePublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseInstancePublicConnectionResponse) SetBody(v *ReleaseInstancePublicConnectionResponseBody) *ReleaseInstancePublicConnectionResponse {
	s.Body = v
	return s
}

type ResetAccountPasswordRequest struct {
	AccountName     *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ResetAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRequest) SetAccountName(v string) *ResetAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountPassword(v string) *ResetAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetDBInstanceId(v string) *ResetAccountPasswordRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetRegionId(v string) *ResetAccountPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetResourceOwnerId(v int64) *ResetAccountPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

type ResetAccountPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ResetAccountPasswordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponse) SetHeaders(v map[string]*string) *ResetAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetAccountPasswordResponse) SetStatusCode(v int32) *ResetAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAccountPasswordResponse) SetBody(v *ResetAccountPasswordResponseBody) *ResetAccountPasswordResponse {
	s.Body = v
	return s
}

type RestartDBClusterRequest struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// 代表资源一级ID的资源属性字段
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// 地域
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RestartDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDBClusterRequest) GoString() string {
	return s.String()
}

func (s *RestartDBClusterRequest) SetDBClusterId(v string) *RestartDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *RestartDBClusterRequest) SetDBInstanceId(v string) *RestartDBClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RestartDBClusterRequest) SetRegionId(v string) *RestartDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *RestartDBClusterRequest) SetResourceGroupId(v string) *RestartDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RestartDBClusterRequest) SetResourceOwnerId(v int64) *RestartDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

type RestartDBClusterResponseBody struct {
	Data      *RestartDBClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDBClusterResponseBody) SetData(v *RestartDBClusterResponseBodyData) *RestartDBClusterResponseBody {
	s.Data = v
	return s
}

func (s *RestartDBClusterResponseBody) SetRequestId(v string) *RestartDBClusterResponseBody {
	s.RequestId = &v
	return s
}

type RestartDBClusterResponseBodyData struct {
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s RestartDBClusterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RestartDBClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestartDBClusterResponseBodyData) SetDBClusterId(v string) *RestartDBClusterResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *RestartDBClusterResponseBodyData) SetDBInstanceId(v string) *RestartDBClusterResponseBodyData {
	s.DBInstanceId = &v
	return s
}

type RestartDBClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartDBClusterResponse) GoString() string {
	return s.String()
}

func (s *RestartDBClusterResponse) SetHeaders(v map[string]*string) *RestartDBClusterResponse {
	s.Headers = v
	return s
}

func (s *RestartDBClusterResponse) SetStatusCode(v int32) *RestartDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDBClusterResponse) SetBody(v *RestartDBClusterResponseBody) *RestartDBClusterResponse {
	s.Body = v
	return s
}

type StartBEClusterRequest struct {
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StartBEClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s StartBEClusterRequest) GoString() string {
	return s.String()
}

func (s *StartBEClusterRequest) SetDBClusterId(v string) *StartBEClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *StartBEClusterRequest) SetDBInstanceId(v string) *StartBEClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *StartBEClusterRequest) SetRegionId(v string) *StartBEClusterRequest {
	s.RegionId = &v
	return s
}

func (s *StartBEClusterRequest) SetResourceOwnerId(v int64) *StartBEClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

type StartBEClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartBEClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartBEClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StartBEClusterResponseBody) SetRequestId(v string) *StartBEClusterResponseBody {
	s.RequestId = &v
	return s
}

type StartBEClusterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartBEClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartBEClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s StartBEClusterResponse) GoString() string {
	return s.String()
}

func (s *StartBEClusterResponse) SetHeaders(v map[string]*string) *StartBEClusterResponse {
	s.Headers = v
	return s
}

func (s *StartBEClusterResponse) SetStatusCode(v int32) *StartBEClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StartBEClusterResponse) SetBody(v *StartBEClusterResponseBody) *StartBEClusterResponse {
	s.Body = v
	return s
}

type StopBEClusterRequest struct {
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StopBEClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s StopBEClusterRequest) GoString() string {
	return s.String()
}

func (s *StopBEClusterRequest) SetDBClusterId(v string) *StopBEClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *StopBEClusterRequest) SetDBInstanceId(v string) *StopBEClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *StopBEClusterRequest) SetRegionId(v string) *StopBEClusterRequest {
	s.RegionId = &v
	return s
}

func (s *StopBEClusterRequest) SetResourceOwnerId(v int64) *StopBEClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

type StopBEClusterResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopBEClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopBEClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StopBEClusterResponseBody) SetCode(v string) *StopBEClusterResponseBody {
	s.Code = &v
	return s
}

func (s *StopBEClusterResponseBody) SetMessage(v string) *StopBEClusterResponseBody {
	s.Message = &v
	return s
}

func (s *StopBEClusterResponseBody) SetRequestId(v string) *StopBEClusterResponseBody {
	s.RequestId = &v
	return s
}

type StopBEClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopBEClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopBEClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s StopBEClusterResponse) GoString() string {
	return s.String()
}

func (s *StopBEClusterResponse) SetHeaders(v map[string]*string) *StopBEClusterResponse {
	s.Headers = v
	return s
}

func (s *StopBEClusterResponse) SetStatusCode(v int32) *StopBEClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StopBEClusterResponse) SetBody(v *StopBEClusterResponseBody) *StopBEClusterResponse {
	s.Body = v
	return s
}

type UpgradeDBInstanceEngineVersionRequest struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EngineVersion   *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SwitchTimeMode  *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s UpgradeDBInstanceEngineVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceEngineVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetDBInstanceId(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetEngineVersion(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.EngineVersion = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetRegionId(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBInstanceEngineVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetSwitchTimeMode(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.SwitchTimeMode = &v
	return s
}

type UpgradeDBInstanceEngineVersionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeDBInstanceEngineVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceEngineVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceEngineVersionResponseBody) SetCode(v string) *UpgradeDBInstanceEngineVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionResponseBody) SetMessage(v string) *UpgradeDBInstanceEngineVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionResponseBody) SetRequestId(v string) *UpgradeDBInstanceEngineVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeDBInstanceEngineVersionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeDBInstanceEngineVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeDBInstanceEngineVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceEngineVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceEngineVersionResponse) SetHeaders(v map[string]*string) *UpgradeDBInstanceEngineVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBInstanceEngineVersionResponse) SetStatusCode(v int32) *UpgradeDBInstanceEngineVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionResponse) SetBody(v *UpgradeDBInstanceEngineVersionResponseBody) *UpgradeDBInstanceEngineVersionResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("selectdb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AllocateInstancePublicConnectionWithOptions(request *AllocateInstancePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionStringPrefix)) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NetType)) {
		query["NetType"] = request.NetType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AllocateInstancePublicConnection"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocateInstancePublicConnection(request *AllocateInstancePublicConnectionRequest) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.AllocateInstancePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckCreateDBInstanceWithOptions(request *CheckCreateDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CheckCreateDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CacheSize)) {
		query["CacheSize"] = request.CacheSize
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionString)) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceClass)) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceDescription)) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIPList)) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSize)) {
		query["StorageSize"] = request.StorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UsedTime)) {
		query["UsedTime"] = request.UsedTime
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckCreateDBInstance"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckCreateDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckCreateDBInstance(request *CheckCreateDBInstanceRequest) (_result *CheckCreateDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckCreateDBInstanceResponse{}
	_body, _err := client.CheckCreateDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckServiceLinkedRoleWithOptions(request *CheckServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CheckServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckServiceLinkedRole"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckServiceLinkedRole(request *CheckServiceLinkedRoleRequest) (_result *CheckServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CheckServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBClusterWithOptions(request *CreateDBClusterRequest, runtime *util.RuntimeOptions) (_result *CreateDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CacheSize)) {
		query["CacheSize"] = request.CacheSize
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterClass)) {
		query["DBClusterClass"] = request.DBClusterClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterDescription)) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSize)) {
		query["StorageSize"] = request.StorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UsedTime)) {
		query["UsedTime"] = request.UsedTime
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		body["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBCluster"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBCluster(request *CreateDBClusterRequest) (_result *CreateDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CreateDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBInstanceWithOptions(request *CreateDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CacheSize)) {
		query["CacheSize"] = request.CacheSize
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionString)) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceClass)) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceDescription)) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIPList)) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSize)) {
		query["StorageSize"] = request.StorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UsedTime)) {
		query["UsedTime"] = request.UsedTime
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBInstance"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBInstance(request *CreateDBInstanceRequest) (_result *CreateDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.CreateDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceLinkedRoleForSelectDBWithOptions(request *CreateServiceLinkedRoleForSelectDBRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleForSelectDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceLinkedRoleForSelectDB"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceLinkedRoleForSelectDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceLinkedRoleForSelectDB(request *CreateServiceLinkedRoleForSelectDBRequest) (_result *CreateServiceLinkedRoleForSelectDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleForSelectDBResponse{}
	_body, _err := client.CreateServiceLinkedRoleForSelectDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBClusterWithOptions(request *DeleteDBClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		body["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBCluster"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBCluster(request *DeleteDBClusterRequest) (_result *DeleteDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.DeleteDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBInstanceWithOptions(request *DeleteDBInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		body["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBInstance"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) (_result *DeleteDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBInstanceResponse{}
	_body, _err := client.DeleteDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceAttributeWithOptions(request *DescribeDBInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceAttribute"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceAttribute(request *DescribeDBInstanceAttributeRequest) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceAttributeResponse{}
	_body, _err := client.DescribeDBInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceNetInfoWithOptions(request *DescribeDBInstanceNetInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceNetInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceNetInfo"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceNetInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceNetInfo(request *DescribeDBInstanceNetInfoRequest) (_result *DescribeDBInstanceNetInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceNetInfoResponse{}
	_body, _err := client.DescribeDBInstanceNetInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstancesWithOptions(request *DescribeDBInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceDescription)) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceIds)) {
		query["DBInstanceIds"] = request.DBInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceStatus)) {
		query["DBInstanceStatus"] = request.DBInstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstances"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (_result *DescribeDBInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.DescribeDBInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityIPListWithOptions(request *DescribeSecurityIPListRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityIPListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSecurityIPList"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSecurityIPListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityIPList(request *DescribeSecurityIPListRequest) (_result *DescribeSecurityIPListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityIPListResponse{}
	_body, _err := client.DescribeSecurityIPListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBEClusterAttributeWithOptions(request *ModifyBEClusterAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyBEClusterAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceAttributeType)) {
		query["InstanceAttributeType"] = request.InstanceAttributeType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyBEClusterAttribute"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyBEClusterAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBEClusterAttribute(request *ModifyBEClusterAttributeRequest) (_result *ModifyBEClusterAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBEClusterAttributeResponse{}
	_body, _err := client.ModifyBEClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterWithOptions(request *ModifyDBClusterRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterClass)) {
		query["DBClusterClass"] = request.DBClusterClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBCluster"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBCluster(request *ModifyDBClusterRequest) (_result *ModifyDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.ModifyDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceAttributeWithOptions(request *ModifyDBInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceAttributeType)) {
		query["InstanceAttributeType"] = request.InstanceAttributeType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceAttribute"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceAttribute(request *ModifyDBInstanceAttributeRequest) (_result *ModifyDBInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceAttributeResponse{}
	_body, _err := client.ModifyDBInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySecurityIPListWithOptions(request *ModifySecurityIPListRequest, runtime *util.RuntimeOptions) (_result *ModifySecurityIPListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySecurityIPList"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySecurityIPListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySecurityIPList(request *ModifySecurityIPListRequest) (_result *ModifySecurityIPListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySecurityIPListResponse{}
	_body, _err := client.ModifySecurityIPListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseInstancePublicConnectionWithOptions(request *ReleaseInstancePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionString)) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseInstancePublicConnection"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseInstancePublicConnection(request *ReleaseInstancePublicConnectionRequest) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.ReleaseInstancePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAccountPasswordWithOptions(request *ResetAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAccountPassword"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (_result *ResetAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.ResetAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartDBClusterWithOptions(request *RestartDBClusterRequest, runtime *util.RuntimeOptions) (_result *RestartDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		body["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartDBCluster"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartDBCluster(request *RestartDBClusterRequest) (_result *RestartDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartDBClusterResponse{}
	_body, _err := client.RestartDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartBEClusterWithOptions(request *StartBEClusterRequest, runtime *util.RuntimeOptions) (_result *StartBEClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartBECluster"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartBEClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartBECluster(request *StartBEClusterRequest) (_result *StartBEClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartBEClusterResponse{}
	_body, _err := client.StartBEClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopBEClusterWithOptions(request *StopBEClusterRequest, runtime *util.RuntimeOptions) (_result *StopBEClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopBECluster"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopBEClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopBECluster(request *StopBEClusterRequest) (_result *StopBEClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopBEClusterResponse{}
	_body, _err := client.StopBEClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeDBInstanceEngineVersionWithOptions(request *UpgradeDBInstanceEngineVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeDBInstanceEngineVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeDBInstanceEngineVersion"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeDBInstanceEngineVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeDBInstanceEngineVersion(request *UpgradeDBInstanceEngineVersionRequest) (_result *UpgradeDBInstanceEngineVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeDBInstanceEngineVersionResponse{}
	_body, _err := client.UpgradeDBInstanceEngineVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
