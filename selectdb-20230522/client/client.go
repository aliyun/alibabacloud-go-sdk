// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AllocateInstancePublicConnectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-h033cn****-pub-i3
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Public
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 5ED62C81-9948-5612-81E1-EA3853752306
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 498115273
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AllocateInstancePublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionResponseBody) SetInstanceName(v string) *AllocateInstancePublicConnectionResponseBody {
	s.InstanceName = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponseBody) SetRequestId(v string) *AllocateInstancePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponseBody) SetTaskId(v int64) *AllocateInstancePublicConnectionResponseBody {
	s.TaskId = &v
	return s
}

type AllocateInstancePublicConnectionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateInstancePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 200
	CacheSize *int32 `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// AB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// selectdb-cn-7213c8y****-public.selectdbfe.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb.xlarge
	DBInstanceClass       *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// example:
	//
	// SelectDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 172.16.XX.XX/12,192.168.XX.XX/22
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-i
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
	// example:
	//
	// ADF42B18-43FD-5100-83A9-BE81AB70C863
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckCreateDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckCreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCreateDBInstanceResponseBody) SetRequestId(v string) *CheckCreateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CheckCreateDBInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCreateDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// False
	HasServiceLinkedRole *bool `json:"HasServiceLinkedRole,omitempty" xml:"HasServiceLinkedRole,omitempty"`
	// example:
	//
	// 1F455133-981E-5AD0-80EB-26EA1EF3C65F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 200
	CacheSize *string `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb.2xlarge
	DBClusterClass *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testdb
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// selectdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 代表资源组的资源属性字段
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-i
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
	Data *CreateDBClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// F8900A96-67F7-5274-A41B-7722E1ECF8C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 219543646290345
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 200GB
	CacheSize *int32 `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// AB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// selectdb-cn-7213c8y****-public.selectdbfe.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb.xlarge
	DBInstanceClass       *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// example:
	//
	// selectdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 代表资源组的资源属性字段
	//
	// example:
	//
	// rg-aekzt2zaluvuvqa_fake
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 192.168.1.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-e
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
	Data *CreateDBInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 4773E4EC-025D-509F-AEA9-D53123FDFB0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 21137950671****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// F203FA74-3041-589F-BE66-E570793A0C91
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceLinkedRoleForSelectDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// selectdb-xxxb9f2w-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 代表资源组的资源属性字段
	//
	// example:
	//
	// rg-4690g37929****
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
	Data *DeleteDBClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// F203FA74-3041-589F-BE66-E570793A0C91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 220088764060782
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// example:
	//
	// BD0D0B17-C145-5B91-BFC2-6791927EE973
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeDBClusterConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// be.conf
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213c8yvv09-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBClusterConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigRequest) SetConfigKey(v string) *DescribeDBClusterConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetDBClusterId(v string) *DescribeDBClusterConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetDBInstanceId(v string) *DescribeDBClusterConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetRegionId(v string) *DescribeDBClusterConfigRequest {
	s.RegionId = &v
	return s
}

type DescribeDBClusterConfigResponseBody struct {
	// example:
	//
	// failed
	AccessDeniedDetail *string                                  `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *DescribeDBClusterConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// An error occurred while processing your request.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// ADF42B18-43FD-5100-83A9-BE81AB70C863
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigResponseBody) SetAccessDeniedDetail(v string) *DescribeDBClusterConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBody) SetData(v *DescribeDBClusterConfigResponseBodyData) *DescribeDBClusterConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBClusterConfigResponseBody) SetDynamicCode(v string) *DescribeDBClusterConfigResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBody) SetDynamicMessage(v string) *DescribeDBClusterConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBody) SetRequestId(v string) *DescribeDBClusterConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterConfigResponseBodyData struct {
	// example:
	//
	// selectdb-cn-wny3li00g02-be
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// example:
	//
	// 6585
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// example:
	//
	// selectdb-cn-wny3li00g02
	DbInstanceName *string                                          `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	Params         []*DescribeDBClusterConfigResponseBodyDataParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 107841167
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDBClusterConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigResponseBodyData) SetDbClusterId(v string) *DescribeDBClusterConfigResponseBodyData {
	s.DbClusterId = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyData) SetDbInstanceId(v string) *DescribeDBClusterConfigResponseBodyData {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyData) SetDbInstanceName(v string) *DescribeDBClusterConfigResponseBodyData {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyData) SetParams(v []*DescribeDBClusterConfigResponseBodyDataParams) *DescribeDBClusterConfigResponseBodyData {
	s.Params = v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyData) SetTaskId(v int32) *DescribeDBClusterConfigResponseBodyData {
	s.TaskId = &v
	return s
}

type DescribeDBClusterConfigResponseBodyDataParams struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 15
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// true
	IsDynamic *int32 `json:"IsDynamic,omitempty" xml:"IsDynamic,omitempty"`
	// example:
	//
	// true
	IsUserModifiable *int32 `json:"IsUserModifiable,omitempty" xml:"IsUserModifiable,omitempty"`
	// example:
	//
	// doris_scanner_thread_pool_thread_num
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// [0-20000]
	Optional      *string `json:"Optional,omitempty" xml:"Optional,omitempty"`
	ParamCategory *string `json:"ParamCategory,omitempty" xml:"ParamCategory,omitempty"`
	// example:
	//
	// 10
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClusterConfigResponseBodyDataParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigResponseBodyDataParams) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetComment(v string) *DescribeDBClusterConfigResponseBodyDataParams {
	s.Comment = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetDefaultValue(v string) *DescribeDBClusterConfigResponseBodyDataParams {
	s.DefaultValue = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetIsDynamic(v int32) *DescribeDBClusterConfigResponseBodyDataParams {
	s.IsDynamic = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetIsUserModifiable(v int32) *DescribeDBClusterConfigResponseBodyDataParams {
	s.IsUserModifiable = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetName(v string) *DescribeDBClusterConfigResponseBodyDataParams {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetOptional(v string) *DescribeDBClusterConfigResponseBodyDataParams {
	s.Optional = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetParamCategory(v string) *DescribeDBClusterConfigResponseBodyDataParams {
	s.ParamCategory = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBodyDataParams) SetValue(v string) *DescribeDBClusterConfigResponseBodyDataParams {
	s.Value = &v
	return s
}

type DescribeDBClusterConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigResponse) SetHeaders(v map[string]*string) *DescribeDBClusterConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterConfigResponse) SetStatusCode(v int32) *DescribeDBClusterConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterConfigResponse) SetBody(v *DescribeDBClusterConfigResponseBody) *DescribeDBClusterConfigResponse {
	s.Body = v
	return s
}

type DescribeDBClusterConfigChangeLogsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// be.conf
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213c8yvv09-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-jia3ma3b003
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-08T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-25T09:48:23Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterConfigChangeLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigChangeLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigChangeLogsRequest) SetConfigKey(v string) *DescribeDBClusterConfigChangeLogsRequest {
	s.ConfigKey = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsRequest) SetDBClusterId(v string) *DescribeDBClusterConfigChangeLogsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsRequest) SetDBInstanceId(v string) *DescribeDBClusterConfigChangeLogsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsRequest) SetEndTime(v string) *DescribeDBClusterConfigChangeLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsRequest) SetRegionId(v string) *DescribeDBClusterConfigChangeLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsRequest) SetStartTime(v string) *DescribeDBClusterConfigChangeLogsRequest {
	s.StartTime = &v
	return s
}

type DescribeDBClusterConfigChangeLogsResponseBody struct {
	// example:
	//
	// failed
	AccessDeniedDetail *string                                            `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *DescribeDBClusterConfigChangeLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// An error occurred while processing your request.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// F8900A96-67F7-5274-A41B-7722E1ECF8C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterConfigChangeLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigChangeLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) SetAccessDeniedDetail(v string) *DescribeDBClusterConfigChangeLogsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) SetData(v *DescribeDBClusterConfigChangeLogsResponseBodyData) *DescribeDBClusterConfigChangeLogsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) SetDynamicCode(v string) *DescribeDBClusterConfigChangeLogsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) SetDynamicMessage(v string) *DescribeDBClusterConfigChangeLogsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBody) SetRequestId(v string) *DescribeDBClusterConfigChangeLogsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterConfigChangeLogsResponseBodyData struct {
	// example:
	//
	// selectdb-cn-wny3li00g02-be
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// example:
	//
	// 6585
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// example:
	//
	// selectdb-cn-wny3li00g02
	DbInstanceName  *string                                                             `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	ParamChangeLogs []*DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs `json:"ParamChangeLogs,omitempty" xml:"ParamChangeLogs,omitempty" type:"Repeated"`
	// example:
	//
	// 107841167
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDBClusterConfigChangeLogsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigChangeLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) SetDbClusterId(v string) *DescribeDBClusterConfigChangeLogsResponseBodyData {
	s.DbClusterId = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) SetDbInstanceId(v string) *DescribeDBClusterConfigChangeLogsResponseBodyData {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) SetDbInstanceName(v string) *DescribeDBClusterConfigChangeLogsResponseBodyData {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) SetParamChangeLogs(v []*DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) *DescribeDBClusterConfigChangeLogsResponseBodyData {
	s.ParamChangeLogs = v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyData) SetTaskId(v int32) *DescribeDBClusterConfigChangeLogsResponseBodyData {
	s.TaskId = &v
	return s
}

type DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs struct {
	// example:
	//
	// 2022-10-11T08:53:32Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2024-03-08T10:08Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 617975
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IsApplied *bool `json:"IsApplied,omitempty" xml:"IsApplied,omitempty"`
	// example:
	//
	// cumulative_compaction_rounds_for_each_base_compaction_round
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12
	NewValue *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	// example:
	//
	// 10
	OldValue *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
}

func (s DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) SetGmtCreated(v string) *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) SetGmtModified(v string) *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	s.GmtModified = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) SetId(v int64) *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	s.Id = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) SetIsApplied(v bool) *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	s.IsApplied = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) SetName(v string) *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) SetNewValue(v string) *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	s.NewValue = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs) SetOldValue(v string) *DescribeDBClusterConfigChangeLogsResponseBodyDataParamChangeLogs {
	s.OldValue = &v
	return s
}

type DescribeDBClusterConfigChangeLogsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterConfigChangeLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterConfigChangeLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigChangeLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigChangeLogsResponse) SetHeaders(v map[string]*string) *DescribeDBClusterConfigChangeLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponse) SetStatusCode(v int32) *DescribeDBClusterConfigChangeLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponse) SetBody(v *DescribeDBClusterConfigChangeLogsResponseBody) *DescribeDBClusterConfigChangeLogsResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	CanUpgradeVersions []*string `json:"CanUpgradeVersions,omitempty" xml:"CanUpgradeVersions,omitempty" type:"Repeated"`
	// example:
	//
	// Prepaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 2023-08-14T03:00:42Z
	CreateTime    *string                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBClusterList []*DescribeDBInstanceAttributeResponseBodyDBClusterList `json:"DBClusterList,omitempty" xml:"DBClusterList,omitempty" type:"Repeated"`
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// selectdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 3.0.1
	EngineMinorVersion *string `json:"EngineMinorVersion,omitempty" xml:"EngineMinorVersion,omitempty"`
	// example:
	//
	// 2.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// 2023-09-17T00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2023-08-17T09:58Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// lock
	LockMode *int64 `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// nolock
	LockReason        *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MaintainEndtime   *string `json:"MaintainEndtime,omitempty" xml:"MaintainEndtime,omitempty"`
	MaintainStarttime *string `json:"MaintainStarttime,omitempty" xml:"MaintainStarttime,omitempty"`
	// example:
	//
	// 0
	ObjectStoreSize *int64 `json:"ObjectStoreSize,omitempty" xml:"ObjectStoreSize,omitempty"`
	// example:
	//
	// 06758CAB-1204-5852-A471-29C87D5C1D0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8
	ResourceCpu     *int64  `json:"ResourceCpu,omitempty" xml:"ResourceCpu,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 400
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// example:
	//
	// cn-beijing-h-aliyun
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
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

func (s *DescribeDBInstanceAttributeResponseBody) SetResourceGroupId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.ResourceGroupId = &v
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
	// example:
	//
	// 200
	CacheStorageSizeGB *string `json:"CacheStorageSizeGB,omitempty" xml:"CacheStorageSizeGB,omitempty"`
	// example:
	//
	// cloud_essd
	CacheStorageType *string `json:"CacheStorageType,omitempty" xml:"CacheStorageType,omitempty"`
	// example:
	//
	// Prepaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 8
	CpuCores *int64 `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	// example:
	//
	// 2023-08-14T09:24:13Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// selectdb.2xlarge
	DbClusterClass *string `json:"DbClusterClass,omitempty" xml:"DbClusterClass,omitempty"`
	// example:
	//
	// selectdb-cn-h033cjs****-be
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// example:
	//
	// test01
	DbClusterName  *string `json:"DbClusterName,omitempty" xml:"DbClusterName,omitempty"`
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// example:
	//
	// 64
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// example:
	//
	// 2023-08-14T09:24:13Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	DBClustersNetInfos []*DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos `json:"DBClustersNetInfos,omitempty" xml:"DBClustersNetInfos,omitempty" type:"Repeated"`
	DBInstanceNetInfos []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos `json:"DBInstanceNetInfos,omitempty" xml:"DBInstanceNetInfos,omitempty" type:"Repeated"`
	// example:
	//
	// ADF42B18-43FD-5100-83A9-BE81AB70C863
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetDBClustersNetInfos(v []*DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) *DescribeDBInstanceNetInfoResponseBody {
	s.DBClustersNetInfos = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetDBInstanceNetInfos(v []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) *DescribeDBInstanceNetInfoResponseBody {
	s.DBInstanceNetInfos = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos struct {
	ClusterId        *string                                                            `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ConnectionString *string                                                            `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	Ip               *string                                                            `json:"Ip,omitempty" xml:"Ip,omitempty"`
	NetType          *string                                                            `json:"NetType,omitempty" xml:"NetType,omitempty"`
	PortList         []*DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList `json:"PortList,omitempty" xml:"PortList,omitempty" type:"Repeated"`
	UserVisible      *bool                                                              `json:"UserVisible,omitempty" xml:"UserVisible,omitempty"`
	VpcId            *string                                                            `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcInstanceId    *string                                                            `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
	VswitchId        *string                                                            `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetClusterId(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.ClusterId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetConnectionString(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetIp(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.Ip = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetNetType(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.NetType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetPortList(v []*DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.PortList = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetUserVisible(v bool) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.UserVisible = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetVpcId(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetVpcInstanceId(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetVswitchId(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.VswitchId = &v
	return s
}

type DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList) SetPort(v int32) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList) SetProtocol(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList {
	s.Protocol = &v
	return s
}

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// selectdb-cn-h033cnd****-fe.selectdbfe.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// 172.16.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// VPC
	NetType  *string                                                            `json:"NetType,omitempty" xml:"NetType,omitempty"`
	PortList []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList `json:"PortList,omitempty" xml:"PortList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	UserVisible *bool `json:"UserVisible,omitempty" xml:"UserVisible,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-wz90scxq6ods388ft****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// selectdb-cn-h033cnd****-fe-20230816101006
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
	// example:
	//
	// vsw-uf6mlqti065rer6m0****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetClusterId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.ClusterId = &v
	return s
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
	// example:
	//
	// MySQLPort
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 9030
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceNetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceIds *string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	// example:
	//
	// ACTIVATION
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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
	Items []*DescribeDBInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 30
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// BC854513-E85E-54F3-9842-B9CCD3308CDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
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
	// example:
	//
	// basic
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 1
	ClusterCount *int32 `json:"ClusterCount,omitempty" xml:"ClusterCount,omitempty"`
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// selectdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// 2024-03-29T03:47:05Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2023-08-12T04:14Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2023-08-12T19:05Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// Instance
	InstanceUsedType *string `json:"InstanceUsedType,omitempty" xml:"InstanceUsedType,omitempty"`
	// example:
	//
	// false
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// example:
	//
	// 0
	LockMode *int64 `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// nolock
	LockReason           *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MaintainEndTimeStr   *string `json:"MaintainEndTimeStr,omitempty" xml:"MaintainEndTimeStr,omitempty"`
	MaintainEndtime      *string `json:"MaintainEndtime,omitempty" xml:"MaintainEndtime,omitempty"`
	MaintainStartTimeStr *string `json:"MaintainStartTimeStr,omitempty" xml:"MaintainStartTimeStr,omitempty"`
	MaintainStarttime    *string `json:"MaintainStarttime,omitempty" xml:"MaintainStarttime,omitempty"`
	// example:
	//
	// 200
	ObjectStoreSize *int64  `json:"ObjectStoreSize,omitempty" xml:"ObjectStoreSize,omitempty"`
	ParentInstance  *string `json:"ParentInstance,omitempty" xml:"ParentInstance,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 8
	ResourceCpu     *int64  `json:"ResourceCpu,omitempty" xml:"ResourceCpu,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 64
	ResourceMemory *int64 `json:"ResourceMemory,omitempty" xml:"ResourceMemory,omitempty"`
	// example:
	//
	// 0
	ScaleMax *int64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// example:
	//
	// 0
	ScaleMin     *int64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	ScaleReplica *int64 `json:"ScaleReplica,omitempty" xml:"ScaleReplica,omitempty"`
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 100
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// example:
	//
	// cloud_essd
	StorageType     *string                                     `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tags            []*DescribeDBInstancesResponseBodyItemsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TenantClusterId *string                                     `json:"TenantClusterId,omitempty" xml:"TenantClusterId,omitempty"`
	TenantToken     *string                                     `json:"TenantToken,omitempty" xml:"TenantToken,omitempty"`
	TenantUserId    *string                                     `json:"TenantUserId,omitempty" xml:"TenantUserId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-i
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

type DescribeDBInstancesResponseBodyItemsTags struct {
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// it
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceName *string                                         `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	GroupItems     []*DescribeSecurityIPListResponseBodyGroupItems `json:"GroupItems,omitempty" xml:"GroupItems,omitempty" type:"Repeated"`
	// example:
	//
	// 5CBE044D-4594-525D-AC65-E988553D853E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSecurityIPListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIPListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListResponseBody) SetDBInstanceName(v string) *DescribeSecurityIPListResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSecurityIPListResponseBody) SetGroupItems(v []*DescribeSecurityIPListResponseBodyGroupItems) *DescribeSecurityIPListResponseBody {
	s.GroupItems = v
	return s
}

func (s *DescribeSecurityIPListResponseBody) SetRequestId(v string) *DescribeSecurityIPListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSecurityIPListResponseBodyGroupItems struct {
	// example:
	//
	// ipv4
	AecurityIPType *string `json:"AecurityIPType,omitempty" xml:"AecurityIPType,omitempty"`
	// example:
	//
	// group1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// ""
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// example:
	//
	// 127.0.XX.XX
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// mix
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityIPListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetCreateBEClusterInquiryRequest struct {
	// example:
	//
	// 200
	CacheSize *int64 `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// selectdb_go_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 4
	ComputeSize *int64 `json:"ComputeSize,omitempty" xml:"ComputeSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-xxx
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// example:
	//
	// 200
	PreCacheSize *int64 `json:"PreCacheSize,omitempty" xml:"PreCacheSize,omitempty"`
	// example:
	//
	// 4
	PreComputeSize *int64 `json:"PreComputeSize,omitempty" xml:"PreComputeSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hour
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetCreateBEClusterInquiryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCreateBEClusterInquiryRequest) GoString() string {
	return s.String()
}

func (s *GetCreateBEClusterInquiryRequest) SetCacheSize(v int64) *GetCreateBEClusterInquiryRequest {
	s.CacheSize = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetChargeType(v string) *GetCreateBEClusterInquiryRequest {
	s.ChargeType = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetCommodityCode(v string) *GetCreateBEClusterInquiryRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetComputeSize(v int64) *GetCreateBEClusterInquiryRequest {
	s.ComputeSize = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetDbInstanceId(v string) *GetCreateBEClusterInquiryRequest {
	s.DbInstanceId = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetPreCacheSize(v int64) *GetCreateBEClusterInquiryRequest {
	s.PreCacheSize = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetPreComputeSize(v int64) *GetCreateBEClusterInquiryRequest {
	s.PreComputeSize = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetPricingCycle(v string) *GetCreateBEClusterInquiryRequest {
	s.PricingCycle = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetQuantity(v int64) *GetCreateBEClusterInquiryRequest {
	s.Quantity = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetRegionId(v string) *GetCreateBEClusterInquiryRequest {
	s.RegionId = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetResourceOwnerId(v int64) *GetCreateBEClusterInquiryRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetCreateBEClusterInquiryResponseBody struct {
	Data *GetCreateBEClusterInquiryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 06758CAB-1204-5852-A471-29C87D5C1D0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCreateBEClusterInquiryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCreateBEClusterInquiryResponseBody) GoString() string {
	return s.String()
}

func (s *GetCreateBEClusterInquiryResponseBody) SetData(v *GetCreateBEClusterInquiryResponseBodyData) *GetCreateBEClusterInquiryResponseBody {
	s.Data = v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBody) SetRequestId(v string) *GetCreateBEClusterInquiryResponseBody {
	s.RequestId = &v
	return s
}

type GetCreateBEClusterInquiryResponseBodyData struct {
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 1.76
	TradeAmount *string `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s GetCreateBEClusterInquiryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCreateBEClusterInquiryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCreateBEClusterInquiryResponseBodyData) SetCurrency(v string) *GetCreateBEClusterInquiryResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBodyData) SetTradeAmount(v string) *GetCreateBEClusterInquiryResponseBodyData {
	s.TradeAmount = &v
	return s
}

type GetCreateBEClusterInquiryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCreateBEClusterInquiryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCreateBEClusterInquiryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCreateBEClusterInquiryResponse) GoString() string {
	return s.String()
}

func (s *GetCreateBEClusterInquiryResponse) SetHeaders(v map[string]*string) *GetCreateBEClusterInquiryResponse {
	s.Headers = v
	return s
}

func (s *GetCreateBEClusterInquiryResponse) SetStatusCode(v int32) *GetCreateBEClusterInquiryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCreateBEClusterInquiryResponse) SetBody(v *GetCreateBEClusterInquiryResponseBody) *GetCreateBEClusterInquiryResponse {
	s.Body = v
	return s
}

type GetModifyBEClusterInquiryRequest struct {
	// example:
	//
	// 200
	CacheSize *int64 `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// selectdb-xxx-be
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb_go_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 4
	ComputeSize *int64 `json:"ComputeSize,omitempty" xml:"ComputeSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-xxx
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// example:
	//
	// 200
	PreCacheSize *int64 `json:"PreCacheSize,omitempty" xml:"PreCacheSize,omitempty"`
	// example:
	//
	// 4
	PreComputeSize *int64 `json:"PreComputeSize,omitempty" xml:"PreComputeSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hour
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetModifyBEClusterInquiryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetModifyBEClusterInquiryRequest) GoString() string {
	return s.String()
}

func (s *GetModifyBEClusterInquiryRequest) SetCacheSize(v int64) *GetModifyBEClusterInquiryRequest {
	s.CacheSize = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetChargeType(v string) *GetModifyBEClusterInquiryRequest {
	s.ChargeType = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetClusterId(v string) *GetModifyBEClusterInquiryRequest {
	s.ClusterId = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetCommodityCode(v string) *GetModifyBEClusterInquiryRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetComputeSize(v int64) *GetModifyBEClusterInquiryRequest {
	s.ComputeSize = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetDbInstanceId(v string) *GetModifyBEClusterInquiryRequest {
	s.DbInstanceId = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetPreCacheSize(v int64) *GetModifyBEClusterInquiryRequest {
	s.PreCacheSize = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetPreComputeSize(v int64) *GetModifyBEClusterInquiryRequest {
	s.PreComputeSize = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetPricingCycle(v string) *GetModifyBEClusterInquiryRequest {
	s.PricingCycle = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetQuantity(v int64) *GetModifyBEClusterInquiryRequest {
	s.Quantity = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetRegionId(v string) *GetModifyBEClusterInquiryRequest {
	s.RegionId = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetResourceOwnerId(v int64) *GetModifyBEClusterInquiryRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetModifyBEClusterInquiryResponseBody struct {
	Data *GetModifyBEClusterInquiryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 06758CAB-1204-5852-A471-29C87D5C1D0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetModifyBEClusterInquiryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModifyBEClusterInquiryResponseBody) GoString() string {
	return s.String()
}

func (s *GetModifyBEClusterInquiryResponseBody) SetData(v *GetModifyBEClusterInquiryResponseBodyData) *GetModifyBEClusterInquiryResponseBody {
	s.Data = v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBody) SetRequestId(v string) *GetModifyBEClusterInquiryResponseBody {
	s.RequestId = &v
	return s
}

type GetModifyBEClusterInquiryResponseBodyData struct {
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 1.76
	TradeAmount *string `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s GetModifyBEClusterInquiryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetModifyBEClusterInquiryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModifyBEClusterInquiryResponseBodyData) SetCurrency(v string) *GetModifyBEClusterInquiryResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyData) SetTradeAmount(v string) *GetModifyBEClusterInquiryResponseBodyData {
	s.TradeAmount = &v
	return s
}

type GetModifyBEClusterInquiryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModifyBEClusterInquiryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModifyBEClusterInquiryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModifyBEClusterInquiryResponse) GoString() string {
	return s.String()
}

func (s *GetModifyBEClusterInquiryResponse) SetHeaders(v map[string]*string) *GetModifyBEClusterInquiryResponse {
	s.Headers = v
	return s
}

func (s *GetModifyBEClusterInquiryResponse) SetStatusCode(v int32) *GetModifyBEClusterInquiryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponse) SetBody(v *GetModifyBEClusterInquiryResponseBody) *GetModifyBEClusterInquiryResponse {
	s.Body = v
	return s
}

type ModifyBEClusterAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-xxxb9f2w-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DBInstanceDescription
	InstanceAttributeType *string `json:"InstanceAttributeType,omitempty" xml:"InstanceAttributeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testdb
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	// example:
	//
	// 58E21E11-90FF-50F8-A615-8DEB193676E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBEClusterAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBEClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBEClusterAttributeResponseBody) SetRequestId(v string) *ModifyBEClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBEClusterAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBEClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// selectdb.2xlarge
	DBClusterClass *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-xxxb9f2w-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// selectdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string                          `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Data         *ModifyDBClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0BF61F90-ACED-55DF-A6FE-56714B2BFCF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// selectdb-cn-7213cjv****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 219396937240838
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ModifyDBClusterConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// be.conf
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213c8yvv09-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-xxxd8a5h60y
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{\\"name\\":\\"param1\\",\\"value\\":\\"1234577777\\"},{\\"name\\":\\"param2\\",\\"value\\":\\"${yyyyMMdd}\\"}]
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s ModifyDBClusterConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigRequest) SetConfigKey(v string) *ModifyDBClusterConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetDBClusterId(v string) *ModifyDBClusterConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetDBInstanceId(v string) *ModifyDBClusterConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetParameters(v string) *ModifyDBClusterConfigRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetRegionId(v string) *ModifyDBClusterConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetSwitchTimeMode(v string) *ModifyDBClusterConfigRequest {
	s.SwitchTimeMode = &v
	return s
}

type ModifyDBClusterConfigResponseBody struct {
	// example:
	//
	// failed
	AccessDeniedDetail *string                                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *ModifyDBClusterConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// An error occurred while processing your request.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// BC854513-E85E-54F3-9842-B9CCD3308CDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigResponseBody) SetAccessDeniedDetail(v string) *ModifyDBClusterConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBody) SetData(v *ModifyDBClusterConfigResponseBodyData) *ModifyDBClusterConfigResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDBClusterConfigResponseBody) SetDynamicCode(v string) *ModifyDBClusterConfigResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBody) SetDynamicMessage(v string) *ModifyDBClusterConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBody) SetRequestId(v string) *ModifyDBClusterConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterConfigResponseBodyData struct {
	// example:
	//
	// selectdb-cn-wny3li00g02-be
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// example:
	//
	// 6585
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// example:
	//
	// selectdb-cn-wny3li00g02
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// example:
	//
	// 107878719
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDBClusterConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigResponseBodyData) SetDbClusterId(v string) *ModifyDBClusterConfigResponseBodyData {
	s.DbClusterId = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBodyData) SetDbInstanceId(v string) *ModifyDBClusterConfigResponseBodyData {
	s.DbInstanceId = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBodyData) SetDbInstanceName(v string) *ModifyDBClusterConfigResponseBodyData {
	s.DbInstanceName = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBodyData) SetTaskId(v int32) *ModifyDBClusterConfigResponseBodyData {
	s.TaskId = &v
	return s
}

type ModifyDBClusterConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigResponse) SetHeaders(v map[string]*string) *ModifyDBClusterConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterConfigResponse) SetStatusCode(v int32) *ModifyDBClusterConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterConfigResponse) SetBody(v *ModifyDBClusterConfigResponseBody) *ModifyDBClusterConfigResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DBInstanceDescription
	InstanceAttributeType *string `json:"InstanceAttributeType,omitempty" xml:"InstanceAttributeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testdb01
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	// example:
	//
	// 2DB29DEE-52E4-57EE-BF68-2C95C20E6658
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAttributeResponseBody) SetRequestId(v string) *ModifyDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 127.0.XX.XX,127.2.XX.XX
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// group1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// grouptag
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// example:
	//
	// 195F64C2-8F11-532B-A436-FC08A221D756
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 127.0.XX.XX,127.2.XX.XX
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// ipv4
	SecurityIPType *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
	// example:
	//
	// 479095561
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// mix
	WhitelistNetType *string `json:"WhitelistNetType,omitempty" xml:"WhitelistNetType,omitempty"`
}

func (s ModifySecurityIPListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIPListResponseBody) GoString() string {
	return s.String()
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

func (s *ModifySecurityIPListResponseBody) SetTaskId(v int64) *ModifySecurityIPListResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetWhitelistNetType(v string) *ModifySecurityIPListResponseBody {
	s.WhitelistNetType = &v
	return s
}

type ModifySecurityIPListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityIPListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213c8y****-public.selectdbfe.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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
	// example:
	//
	// 4773E4EC-025D-509F-AEA9-D53123FDFB0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseInstancePublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionResponseBody) SetRequestId(v string) *ReleaseInstancePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseInstancePublicConnectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseInstancePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// admin
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a1b2c3d4@
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// example:
	//
	// 58E21E11-90FF-50F8-A615-8DEB193676E0
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213c8yvv09-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// 地域
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源组id
	//
	// example:
	//
	// rg-4690g37929****
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
	Data *RestartDBClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// BD0D0B17-C145-5B91-BFC2-6791927EE973
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// selectdb-cn-7213c8y****-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// selectdb-cn-7213cjv****
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213c8yvv09-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// example:
	//
	// F203FA74-3041-589F-BE66-E570793A0C91
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartBEClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// example:
	//
	// BC854513-E85E-54F3-9842-B9CCD3308CDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopBEClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopBEClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StopBEClusterResponseBody) SetRequestId(v string) *StopBEClusterResponseBody {
	s.RequestId = &v
	return s
}

type StopBEClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopBEClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// example:
	//
	// 4773E4EC-025D-509F-AEA9-D53123FDFB0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeDBInstanceEngineVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceEngineVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceEngineVersionResponseBody) SetRequestId(v string) *UpgradeDBInstanceEngineVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeDBInstanceEngineVersionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeDBInstanceEngineVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

// Summary:
//
// 申请公网地址
//
// @param request - AllocateInstancePublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateInstancePublicConnectionResponse
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

// Summary:
//
// 申请公网地址
//
// @param request - AllocateInstancePublicConnectionRequest
//
// @return AllocateInstancePublicConnectionResponse
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

// Summary:
//
// SelectDB实例创建前检查
//
// @param request - CheckCreateDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckCreateDBInstanceResponse
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

// Summary:
//
// SelectDB实例创建前检查
//
// @param request - CheckCreateDBInstanceRequest
//
// @return CheckCreateDBInstanceResponse
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

// Summary:
//
// 检查服务关联角色
//
// @param request - CheckServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckServiceLinkedRoleResponse
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

// Summary:
//
// 检查服务关联角色
//
// @param request - CheckServiceLinkedRoleRequest
//
// @return CheckServiceLinkedRoleResponse
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

// Summary:
//
// SelectDB实例下创建集群
//
// @param request - CreateDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBClusterResponse
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

// Summary:
//
// SelectDB实例下创建集群
//
// @param request - CreateDBClusterRequest
//
// @return CreateDBClusterResponse
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

// Summary:
//
// 创建SelectDB实例
//
// @param request - CreateDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstanceResponse
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

// Summary:
//
// 创建SelectDB实例
//
// @param request - CreateDBInstanceRequest
//
// @return CreateDBInstanceResponse
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

// Summary:
//
// 创建服务关联角色
//
// @param request - CreateServiceLinkedRoleForSelectDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceLinkedRoleForSelectDBResponse
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

// Summary:
//
// 创建服务关联角色
//
// @param request - CreateServiceLinkedRoleForSelectDBRequest
//
// @return CreateServiceLinkedRoleForSelectDBResponse
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

// Summary:
//
// 释放实例下集群
//
// @param request - DeleteDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBClusterResponse
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

// Summary:
//
// 释放实例下集群
//
// @param request - DeleteDBClusterRequest
//
// @return DeleteDBClusterResponse
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

// Summary:
//
// 删除DB实例
//
// @param request - DeleteDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBInstanceResponse
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

// Summary:
//
// 删除DB实例
//
// @param request - DeleteDBInstanceRequest
//
// @return DeleteDBInstanceResponse
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

// Summary:
//
// 查看集群配置
//
// @param request - DescribeDBClusterConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterConfigResponse
func (client *Client) DescribeDBClusterConfigWithOptions(request *DescribeDBClusterConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigKey)) {
		query["ConfigKey"] = request.ConfigKey
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterConfig"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看集群配置
//
// @param request - DescribeDBClusterConfigRequest
//
// @return DescribeDBClusterConfigResponse
func (client *Client) DescribeDBClusterConfig(request *DescribeDBClusterConfigRequest) (_result *DescribeDBClusterConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterConfigResponse{}
	_body, _err := client.DescribeDBClusterConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看集群配置变更记录
//
// @param request - DescribeDBClusterConfigChangeLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterConfigChangeLogsResponse
func (client *Client) DescribeDBClusterConfigChangeLogsWithOptions(request *DescribeDBClusterConfigChangeLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterConfigChangeLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigKey)) {
		query["ConfigKey"] = request.ConfigKey
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterConfigChangeLogs"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterConfigChangeLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看集群配置变更记录
//
// @param request - DescribeDBClusterConfigChangeLogsRequest
//
// @return DescribeDBClusterConfigChangeLogsResponse
func (client *Client) DescribeDBClusterConfigChangeLogs(request *DescribeDBClusterConfigChangeLogsRequest) (_result *DescribeDBClusterConfigChangeLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterConfigChangeLogsResponse{}
	_body, _err := client.DescribeDBClusterConfigChangeLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例详情
//
// @param request - DescribeDBInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceAttributeResponse
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

// Summary:
//
// 查询实例详情
//
// @param request - DescribeDBInstanceAttributeRequest
//
// @return DescribeDBInstanceAttributeResponse
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

// Summary:
//
// 查询实例网络链接
//
// @param request - DescribeDBInstanceNetInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceNetInfoResponse
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

// Summary:
//
// 查询实例网络链接
//
// @param request - DescribeDBInstanceNetInfoRequest
//
// @return DescribeDBInstanceNetInfoResponse
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

// Summary:
//
// 查询实例列表
//
// @param request - DescribeDBInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancesResponse
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

// Summary:
//
// 查询实例列表
//
// @param request - DescribeDBInstancesRequest
//
// @return DescribeDBInstancesResponse
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

// Summary:
//
// 查看白名单
//
// @param request - DescribeSecurityIPListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityIPListResponse
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

// Summary:
//
// 查看白名单
//
// @param request - DescribeSecurityIPListRequest
//
// @return DescribeSecurityIPListResponse
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

// Summary:
//
// 新建集群询价
//
// @param request - GetCreateBEClusterInquiryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreateBEClusterInquiryResponse
func (client *Client) GetCreateBEClusterInquiryWithOptions(request *GetCreateBEClusterInquiryRequest, runtime *util.RuntimeOptions) (_result *GetCreateBEClusterInquiryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCreateBEClusterInquiry"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCreateBEClusterInquiryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建集群询价
//
// @param request - GetCreateBEClusterInquiryRequest
//
// @return GetCreateBEClusterInquiryResponse
func (client *Client) GetCreateBEClusterInquiry(request *GetCreateBEClusterInquiryRequest) (_result *GetCreateBEClusterInquiryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCreateBEClusterInquiryResponse{}
	_body, _err := client.GetCreateBEClusterInquiryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 集群变配询价
//
// @param request - GetModifyBEClusterInquiryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModifyBEClusterInquiryResponse
func (client *Client) GetModifyBEClusterInquiryWithOptions(request *GetModifyBEClusterInquiryRequest, runtime *util.RuntimeOptions) (_result *GetModifyBEClusterInquiryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetModifyBEClusterInquiry"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetModifyBEClusterInquiryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 集群变配询价
//
// @param request - GetModifyBEClusterInquiryRequest
//
// @return GetModifyBEClusterInquiryResponse
func (client *Client) GetModifyBEClusterInquiry(request *GetModifyBEClusterInquiryRequest) (_result *GetModifyBEClusterInquiryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetModifyBEClusterInquiryResponse{}
	_body, _err := client.GetModifyBEClusterInquiryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改BE集群名称、属性、设置
//
// @param request - ModifyBEClusterAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBEClusterAttributeResponse
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

// Summary:
//
// 修改BE集群名称、属性、设置
//
// @param request - ModifyBEClusterAttributeRequest
//
// @return ModifyBEClusterAttributeResponse
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

// Summary:
//
// 集群变配
//
// @param request - ModifyDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterResponse
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

// Summary:
//
// 集群变配
//
// @param request - ModifyDBClusterRequest
//
// @return ModifyDBClusterResponse
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

// Summary:
//
// 修改集群配置
//
// @param request - ModifyDBClusterConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterConfigResponse
func (client *Client) ModifyDBClusterConfigWithOptions(request *ModifyDBClusterConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigKey)) {
		query["ConfigKey"] = request.ConfigKey
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTimeMode)) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterConfig"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改集群配置
//
// @param request - ModifyDBClusterConfigRequest
//
// @return ModifyDBClusterConfigResponse
func (client *Client) ModifyDBClusterConfig(request *ModifyDBClusterConfigRequest) (_result *ModifyDBClusterConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterConfigResponse{}
	_body, _err := client.ModifyDBClusterConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改实例属性
//
// @param request - ModifyDBInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceAttributeResponse
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

// Summary:
//
// 修改实例属性
//
// @param request - ModifyDBInstanceAttributeRequest
//
// @return ModifyDBInstanceAttributeResponse
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

// Summary:
//
// 变更白名单
//
// @param request - ModifySecurityIPListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecurityIPListResponse
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

// Summary:
//
// 变更白名单
//
// @param request - ModifySecurityIPListRequest
//
// @return ModifySecurityIPListResponse
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

// Summary:
//
// 释放公网地址
//
// @param request - ReleaseInstancePublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstancePublicConnectionResponse
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

// Summary:
//
// 释放公网地址
//
// @param request - ReleaseInstancePublicConnectionRequest
//
// @return ReleaseInstancePublicConnectionResponse
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

// Summary:
//
// 修改 Admin 账号的密码。
//
// @param request - ResetAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAccountPasswordResponse
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

// Summary:
//
// 修改 Admin 账号的密码。
//
// @param request - ResetAccountPasswordRequest
//
// @return ResetAccountPasswordResponse
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

// Summary:
//
// 重启BE集群
//
// @param request - RestartDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDBClusterResponse
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

// Summary:
//
// 重启BE集群
//
// @param request - RestartDBClusterRequest
//
// @return RestartDBClusterResponse
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

// Summary:
//
// 暂停后恢复集群
//
// @param request - StartBEClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartBEClusterResponse
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

// Summary:
//
// 暂停后恢复集群
//
// @param request - StartBEClusterRequest
//
// @return StartBEClusterResponse
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

// Summary:
//
// 暂停BE集群
//
// @param request - StopBEClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopBEClusterResponse
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

// Summary:
//
// 暂停BE集群
//
// @param request - StopBEClusterRequest
//
// @return StopBEClusterResponse
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

// Summary:
//
// 实例内核版本升级
//
// @param request - UpgradeDBInstanceEngineVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeDBInstanceEngineVersionResponse
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

// Summary:
//
// 实例内核版本升级
//
// @param request - UpgradeDBInstanceEngineVersionRequest
//
// @return UpgradeDBInstanceEngineVersionResponse
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
