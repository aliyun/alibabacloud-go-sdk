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

type AddBuDBInstanceRelationRequest struct {
	BusinessUnit *string `json:"BusinessUnit,omitempty" xml:"BusinessUnit,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s AddBuDBInstanceRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s AddBuDBInstanceRelationRequest) GoString() string {
	return s.String()
}

func (s *AddBuDBInstanceRelationRequest) SetBusinessUnit(v string) *AddBuDBInstanceRelationRequest {
	s.BusinessUnit = &v
	return s
}

func (s *AddBuDBInstanceRelationRequest) SetDBInstanceId(v string) *AddBuDBInstanceRelationRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AddBuDBInstanceRelationRequest) SetOwnerId(v int64) *AddBuDBInstanceRelationRequest {
	s.OwnerId = &v
	return s
}

type AddBuDBInstanceRelationResponseBody struct {
	BusinessUnit   *string `json:"BusinessUnit,omitempty" xml:"BusinessUnit,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddBuDBInstanceRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddBuDBInstanceRelationResponseBody) GoString() string {
	return s.String()
}

func (s *AddBuDBInstanceRelationResponseBody) SetBusinessUnit(v string) *AddBuDBInstanceRelationResponseBody {
	s.BusinessUnit = &v
	return s
}

func (s *AddBuDBInstanceRelationResponseBody) SetDBInstanceName(v string) *AddBuDBInstanceRelationResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *AddBuDBInstanceRelationResponseBody) SetRequestId(v string) *AddBuDBInstanceRelationResponseBody {
	s.RequestId = &v
	return s
}

type AddBuDBInstanceRelationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddBuDBInstanceRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddBuDBInstanceRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s AddBuDBInstanceRelationResponse) GoString() string {
	return s.String()
}

func (s *AddBuDBInstanceRelationResponse) SetHeaders(v map[string]*string) *AddBuDBInstanceRelationResponse {
	s.Headers = v
	return s
}

func (s *AddBuDBInstanceRelationResponse) SetStatusCode(v int32) *AddBuDBInstanceRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBuDBInstanceRelationResponse) SetBody(v *AddBuDBInstanceRelationResponseBody) *AddBuDBInstanceRelationResponse {
	s.Body = v
	return s
}

type AllocateInstancePublicConnectionRequest struct {
	AddressType            *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Port                   *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionRequest) SetAddressType(v string) *AllocateInstancePublicConnectionRequest {
	s.AddressType = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetDBInstanceId(v string) *AllocateInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetPort(v string) *AllocateInstancePublicConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

type AllocateInstancePublicConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateInstancePublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionResponseBody) SetRequestId(v string) *AllocateInstancePublicConnectionResponseBody {
	s.RequestId = &v
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

type CheckServiceLinkedRoleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleRequest) SetRegionId(v string) *CheckServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

type CheckServiceLinkedRoleResponseBody struct {
	HasServiceLinkedRole *string `json:"HasServiceLinkedRole,omitempty" xml:"HasServiceLinkedRole,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponseBody) SetHasServiceLinkedRole(v string) *CheckServiceLinkedRoleResponseBody {
	s.HasServiceLinkedRole = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetRegionId(v string) *CheckServiceLinkedRoleResponseBody {
	s.RegionId = &v
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

type CreateAccountRequest struct {
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword    *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DatabaseName       *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId    *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetDBInstanceId(v string) *CreateAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateAccountRequest) SetDatabaseName(v string) *CreateAccountRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerId(v int64) *CreateAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetResourceGroupId(v string) *CreateAccountRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountResponse) SetHeaders(v map[string]*string) *CreateAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountResponse) SetStatusCode(v int32) *CreateAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccountResponse) SetBody(v *CreateAccountResponseBody) *CreateAccountResponse {
	s.Body = v
	return s
}

type CreateDBInstanceRequest struct {
	ClientToken           *string                       `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CreateSampleData      *bool                         `json:"CreateSampleData,omitempty" xml:"CreateSampleData,omitempty"`
	DBInstanceCategory    *string                       `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
	DBInstanceClass       *string                       `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceDescription *string                       `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceGroupCount  *string                       `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	DBInstanceMode        *string                       `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	Engine                *string                       `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string                       `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	IdleTime              *int32                        `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	InstanceNetworkType   *string                       `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	InstanceSpec          *string                       `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	MasterNodeNum         *string                       `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	OwnerId               *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayType               *string                       `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period                *string                       `json:"Period,omitempty" xml:"Period,omitempty"`
	PrivateIpAddress      *string                       `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	RegionId              *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityIPList        *string                       `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	SegNodeNum            *string                       `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	SegStorageType        *string                       `json:"SegStorageType,omitempty" xml:"SegStorageType,omitempty"`
	ServerlessMode        *string                       `json:"ServerlessMode,omitempty" xml:"ServerlessMode,omitempty"`
	ServerlessResource    *int32                        `json:"ServerlessResource,omitempty" xml:"ServerlessResource,omitempty"`
	StorageSize           *int64                        `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	StorageType           *string                       `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tag                   []*CreateDBInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UsedTime              *string                       `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VPCId                 *string                       `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId             *string                       `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                *string                       `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCreateSampleData(v bool) *CreateDBInstanceRequest {
	s.CreateSampleData = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceCategory(v string) *CreateDBInstanceRequest {
	s.DBInstanceCategory = &v
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

func (s *CreateDBInstanceRequest) SetDBInstanceGroupCount(v string) *CreateDBInstanceRequest {
	s.DBInstanceGroupCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceMode(v string) *CreateDBInstanceRequest {
	s.DBInstanceMode = &v
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

func (s *CreateDBInstanceRequest) SetIdleTime(v int32) *CreateDBInstanceRequest {
	s.IdleTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetInstanceNetworkType(v string) *CreateDBInstanceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetInstanceSpec(v string) *CreateDBInstanceRequest {
	s.InstanceSpec = &v
	return s
}

func (s *CreateDBInstanceRequest) SetMasterNodeNum(v string) *CreateDBInstanceRequest {
	s.MasterNodeNum = &v
	return s
}

func (s *CreateDBInstanceRequest) SetOwnerId(v int64) *CreateDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPayType(v string) *CreateDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v string) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPrivateIpAddress(v string) *CreateDBInstanceRequest {
	s.PrivateIpAddress = &v
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

func (s *CreateDBInstanceRequest) SetSecurityIPList(v string) *CreateDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSegNodeNum(v string) *CreateDBInstanceRequest {
	s.SegNodeNum = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSegStorageType(v string) *CreateDBInstanceRequest {
	s.SegStorageType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetServerlessMode(v string) *CreateDBInstanceRequest {
	s.ServerlessMode = &v
	return s
}

func (s *CreateDBInstanceRequest) SetServerlessResource(v int32) *CreateDBInstanceRequest {
	s.ServerlessResource = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageSize(v int64) *CreateDBInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageType(v string) *CreateDBInstanceRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTag(v []*CreateDBInstanceRequestTag) *CreateDBInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateDBInstanceRequest) SetUsedTime(v string) *CreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVPCId(v string) *CreateDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

type CreateDBInstanceRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDBInstanceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequestTag) SetKey(v string) *CreateDBInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDBInstanceRequestTag) SetValue(v string) *CreateDBInstanceRequestTag {
	s.Value = &v
	return s
}

type CreateDBInstanceResponseBody struct {
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBody) SetConnectionString(v string) *CreateDBInstanceResponseBody {
	s.ConnectionString = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetDBInstanceId(v string) *CreateDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetOrderId(v string) *CreateDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetPort(v string) *CreateDBInstanceResponseBody {
	s.Port = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetRequestId(v string) *CreateDBInstanceResponseBody {
	s.RequestId = &v
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

type CreateDBInstancePlanRequest struct {
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlanConfig       *string `json:"PlanConfig,omitempty" xml:"PlanConfig,omitempty"`
	PlanDesc         *string `json:"PlanDesc,omitempty" xml:"PlanDesc,omitempty"`
	PlanEndDate      *string `json:"PlanEndDate,omitempty" xml:"PlanEndDate,omitempty"`
	PlanName         *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	PlanScheduleType *string `json:"PlanScheduleType,omitempty" xml:"PlanScheduleType,omitempty"`
	PlanStartDate    *string `json:"PlanStartDate,omitempty" xml:"PlanStartDate,omitempty"`
	PlanType         *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
}

func (s CreateDBInstancePlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstancePlanRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstancePlanRequest) SetDBInstanceId(v string) *CreateDBInstancePlanRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetOwnerId(v int64) *CreateDBInstancePlanRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanConfig(v string) *CreateDBInstancePlanRequest {
	s.PlanConfig = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanDesc(v string) *CreateDBInstancePlanRequest {
	s.PlanDesc = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanEndDate(v string) *CreateDBInstancePlanRequest {
	s.PlanEndDate = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanName(v string) *CreateDBInstancePlanRequest {
	s.PlanName = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanScheduleType(v string) *CreateDBInstancePlanRequest {
	s.PlanScheduleType = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanStartDate(v string) *CreateDBInstancePlanRequest {
	s.PlanStartDate = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanType(v string) *CreateDBInstancePlanRequest {
	s.PlanType = &v
	return s
}

type CreateDBInstancePlanResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PlanId       *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateDBInstancePlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstancePlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstancePlanResponseBody) SetDBInstanceId(v string) *CreateDBInstancePlanResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstancePlanResponseBody) SetErrorMessage(v string) *CreateDBInstancePlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDBInstancePlanResponseBody) SetPlanId(v string) *CreateDBInstancePlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *CreateDBInstancePlanResponseBody) SetRequestId(v string) *CreateDBInstancePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstancePlanResponseBody) SetStatus(v string) *CreateDBInstancePlanResponseBody {
	s.Status = &v
	return s
}

type CreateDBInstancePlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDBInstancePlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBInstancePlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstancePlanResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstancePlanResponse) SetHeaders(v map[string]*string) *CreateDBInstancePlanResponse {
	s.Headers = v
	return s
}

func (s *CreateDBInstancePlanResponse) SetStatusCode(v int32) *CreateDBInstancePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBInstancePlanResponse) SetBody(v *CreateDBInstancePlanResponseBody) *CreateDBInstancePlanResponse {
	s.Body = v
	return s
}

type CreateECSDBInstanceRequest struct {
	BackupId              *string                          `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	ClientToken           *string                          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceCategory    *string                          `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
	DBInstanceDescription *string                          `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	EncryptionKey         *string                          `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	EncryptionType        *string                          `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	Engine                *string                          `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string                          `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	InstanceNetworkType   *string                          `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	InstanceSpec          *string                          `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	MasterNodeNum         *int32                           `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	OwnerId               *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayType               *string                          `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period                *string                          `json:"Period,omitempty" xml:"Period,omitempty"`
	PrivateIpAddress      *string                          `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	RegionId              *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityIPList        *string                          `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	SegNodeNum            *int32                           `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	SegStorageType        *string                          `json:"SegStorageType,omitempty" xml:"SegStorageType,omitempty"`
	SrcDbInstanceName     *string                          `json:"SrcDbInstanceName,omitempty" xml:"SrcDbInstanceName,omitempty"`
	StorageSize           *int32                           `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	Tag                   []*CreateECSDBInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UsedTime              *string                          `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VPCId                 *string                          `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId             *string                          `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                *string                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateECSDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateECSDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateECSDBInstanceRequest) SetBackupId(v string) *CreateECSDBInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetClientToken(v string) *CreateECSDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetDBInstanceCategory(v string) *CreateECSDBInstanceRequest {
	s.DBInstanceCategory = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetDBInstanceDescription(v string) *CreateECSDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetEncryptionKey(v string) *CreateECSDBInstanceRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetEncryptionType(v string) *CreateECSDBInstanceRequest {
	s.EncryptionType = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetEngine(v string) *CreateECSDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetEngineVersion(v string) *CreateECSDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetInstanceNetworkType(v string) *CreateECSDBInstanceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetInstanceSpec(v string) *CreateECSDBInstanceRequest {
	s.InstanceSpec = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetMasterNodeNum(v int32) *CreateECSDBInstanceRequest {
	s.MasterNodeNum = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetOwnerId(v int64) *CreateECSDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetPayType(v string) *CreateECSDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetPeriod(v string) *CreateECSDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetPrivateIpAddress(v string) *CreateECSDBInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetRegionId(v string) *CreateECSDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetResourceGroupId(v string) *CreateECSDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetSecurityIPList(v string) *CreateECSDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetSegNodeNum(v int32) *CreateECSDBInstanceRequest {
	s.SegNodeNum = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetSegStorageType(v string) *CreateECSDBInstanceRequest {
	s.SegStorageType = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetSrcDbInstanceName(v string) *CreateECSDBInstanceRequest {
	s.SrcDbInstanceName = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetStorageSize(v int32) *CreateECSDBInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetTag(v []*CreateECSDBInstanceRequestTag) *CreateECSDBInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateECSDBInstanceRequest) SetUsedTime(v string) *CreateECSDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetVPCId(v string) *CreateECSDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetVSwitchId(v string) *CreateECSDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetZoneId(v string) *CreateECSDBInstanceRequest {
	s.ZoneId = &v
	return s
}

type CreateECSDBInstanceRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateECSDBInstanceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateECSDBInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateECSDBInstanceRequestTag) SetKey(v string) *CreateECSDBInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateECSDBInstanceRequestTag) SetValue(v string) *CreateECSDBInstanceRequestTag {
	s.Value = &v
	return s
}

type CreateECSDBInstanceResponseBody struct {
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateECSDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateECSDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateECSDBInstanceResponseBody) SetConnectionString(v string) *CreateECSDBInstanceResponseBody {
	s.ConnectionString = &v
	return s
}

func (s *CreateECSDBInstanceResponseBody) SetDBInstanceId(v string) *CreateECSDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateECSDBInstanceResponseBody) SetOrderId(v string) *CreateECSDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateECSDBInstanceResponseBody) SetPort(v string) *CreateECSDBInstanceResponseBody {
	s.Port = &v
	return s
}

func (s *CreateECSDBInstanceResponseBody) SetRequestId(v string) *CreateECSDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateECSDBInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateECSDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateECSDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateECSDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateECSDBInstanceResponse) SetHeaders(v map[string]*string) *CreateECSDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateECSDBInstanceResponse) SetStatusCode(v int32) *CreateECSDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateECSDBInstanceResponse) SetBody(v *CreateECSDBInstanceResponseBody) *CreateECSDBInstanceResponse {
	s.Body = v
	return s
}

type CreateSampleDataRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CreateSampleDataRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSampleDataRequest) GoString() string {
	return s.String()
}

func (s *CreateSampleDataRequest) SetDBInstanceId(v string) *CreateSampleDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateSampleDataRequest) SetOwnerId(v int64) *CreateSampleDataRequest {
	s.OwnerId = &v
	return s
}

type CreateSampleDataResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateSampleDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSampleDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSampleDataResponseBody) SetDBInstanceId(v string) *CreateSampleDataResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateSampleDataResponseBody) SetErrorMessage(v string) *CreateSampleDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSampleDataResponseBody) SetRequestId(v string) *CreateSampleDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSampleDataResponseBody) SetStatus(v bool) *CreateSampleDataResponseBody {
	s.Status = &v
	return s
}

type CreateSampleDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSampleDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSampleDataResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSampleDataResponse) GoString() string {
	return s.String()
}

func (s *CreateSampleDataResponse) SetHeaders(v map[string]*string) *CreateSampleDataResponse {
	s.Headers = v
	return s
}

func (s *CreateSampleDataResponse) SetStatusCode(v int32) *CreateSampleDataResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSampleDataResponse) SetBody(v *CreateSampleDataResponseBody) *CreateSampleDataResponse {
	s.Body = v
	return s
}

type CreateServiceLinkedRoleRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) SetOwnerId(v int64) *CreateServiceLinkedRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetRegionId(v string) *CreateServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

type CreateServiceLinkedRoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceLinkedRoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetStatusCode(v int32) *CreateServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetBody(v *CreateServiceLinkedRoleResponseBody) *CreateServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type DeleteDBInstanceRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceRequest) SetClientToken(v string) *DeleteDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetDBInstanceId(v string) *DeleteDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetOwnerId(v int64) *DeleteDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetResourceGroupId(v string) *DeleteDBInstanceRequest {
	s.ResourceGroupId = &v
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

type DeleteDBInstancePlanRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlanId       *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s DeleteDBInstancePlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstancePlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstancePlanRequest) SetDBInstanceId(v string) *DeleteDBInstancePlanRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstancePlanRequest) SetOwnerId(v int64) *DeleteDBInstancePlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBInstancePlanRequest) SetPlanId(v string) *DeleteDBInstancePlanRequest {
	s.PlanId = &v
	return s
}

type DeleteDBInstancePlanResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PlanId       *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteDBInstancePlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstancePlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstancePlanResponseBody) SetDBInstanceId(v string) *DeleteDBInstancePlanResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstancePlanResponseBody) SetErrorMessage(v string) *DeleteDBInstancePlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDBInstancePlanResponseBody) SetPlanId(v string) *DeleteDBInstancePlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *DeleteDBInstancePlanResponseBody) SetRequestId(v string) *DeleteDBInstancePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBInstancePlanResponseBody) SetStatus(v string) *DeleteDBInstancePlanResponseBody {
	s.Status = &v
	return s
}

type DeleteDBInstancePlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDBInstancePlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBInstancePlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstancePlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBInstancePlanResponse) SetHeaders(v map[string]*string) *DeleteDBInstancePlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBInstancePlanResponse) SetStatusCode(v int32) *DeleteDBInstancePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBInstancePlanResponse) SetBody(v *DeleteDBInstancePlanResponseBody) *DeleteDBInstancePlanResponse {
	s.Body = v
	return s
}

type DeleteDatabaseRequest struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBName          *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseRequest) SetDBInstanceId(v string) *DeleteDatabaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetDBName(v string) *DeleteDatabaseRequest {
	s.DBName = &v
	return s
}

func (s *DeleteDatabaseRequest) SetResourceGroupId(v string) *DeleteDatabaseRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteDatabaseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponseBody) SetRequestId(v string) *DeleteDatabaseResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDatabaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeAccountsRequest struct {
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsRequest) SetDBInstanceId(v string) *DescribeAccountsRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeAccountsResponseBody struct {
	Accounts  *DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) SetAccounts(v *DescribeAccountsResponseBodyAccounts) *DescribeAccountsResponseBody {
	s.Accounts = v
	return s
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAccountsResponseBodyAccounts struct {
	DBInstanceAccount []*DescribeAccountsResponseBodyAccountsDBInstanceAccount `json:"DBInstanceAccount,omitempty" xml:"DBInstanceAccount,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccounts) SetDBInstanceAccount(v []*DescribeAccountsResponseBodyAccountsDBInstanceAccount) *DescribeAccountsResponseBodyAccounts {
	s.DBInstanceAccount = v
	return s
}

type DescribeAccountsResponseBodyAccountsDBInstanceAccount struct {
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountStatus      *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountsDBInstanceAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsDBInstanceAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountName(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetDBInstanceId(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.DBInstanceId = &v
	return s
}

type DescribeAccountsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponse) SetHeaders(v map[string]*string) *DescribeAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountsResponse) SetStatusCode(v int32) *DescribeAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountsResponse) SetBody(v *DescribeAccountsResponseBody) *DescribeAccountsResponse {
	s.Body = v
	return s
}

type DescribeAvailableResourcesRequest struct {
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesRequest) SetChargeType(v string) *DescribeAvailableResourcesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeAvailableResourcesRequest) SetRegion(v string) *DescribeAvailableResourcesRequest {
	s.Region = &v
	return s
}

func (s *DescribeAvailableResourcesRequest) SetZoneId(v string) *DescribeAvailableResourcesRequest {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourcesResponseBody struct {
	RegionId  *string                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*DescribeAvailableResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBody) SetRegionId(v string) *DescribeAvailableResourcesResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBody) SetRequestId(v string) *DescribeAvailableResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBody) SetResources(v []*DescribeAvailableResourcesResponseBodyResources) *DescribeAvailableResourcesResponseBody {
	s.Resources = v
	return s
}

type DescribeAvailableResourcesResponseBodyResources struct {
	SupportedEngines []*DescribeAvailableResourcesResponseBodyResourcesSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" type:"Repeated"`
	ZoneId           *string                                                            `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResources) SetSupportedEngines(v []*DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) *DescribeAvailableResourcesResponseBodyResources {
	s.SupportedEngines = v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResources) SetZoneId(v string) *DescribeAvailableResourcesResponseBodyResources {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourcesResponseBodyResourcesSupportedEngines struct {
	Mode                     *string                                                                                    `json:"Mode,omitempty" xml:"Mode,omitempty"`
	SupportedEngineVersion   *string                                                                                    `json:"SupportedEngineVersion,omitempty" xml:"SupportedEngineVersion,omitempty"`
	SupportedInstanceClasses []*DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses `json:"SupportedInstanceClasses,omitempty" xml:"SupportedInstanceClasses,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) SetMode(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines {
	s.Mode = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) SetSupportedEngineVersion(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines {
	s.SupportedEngineVersion = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) SetSupportedInstanceClasses(v []*DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines {
	s.SupportedInstanceClasses = v
	return s
}

type DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses struct {
	Category      *string                                                                                             `json:"Category,omitempty" xml:"Category,omitempty"`
	Description   *string                                                                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayClass  *string                                                                                             `json:"DisplayClass,omitempty" xml:"DisplayClass,omitempty"`
	InstanceClass *string                                                                                             `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	NodeCount     *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount   `json:"NodeCount,omitempty" xml:"NodeCount,omitempty" type:"Struct"`
	StorageSize   *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize `json:"StorageSize,omitempty" xml:"StorageSize,omitempty" type:"Struct"`
	StorageType   *string                                                                                             `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetCategory(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.Category = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetDescription(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.Description = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetDisplayClass(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.DisplayClass = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetInstanceClass(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.InstanceClass = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetNodeCount(v *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.NodeCount = v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetStorageSize(v *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.StorageSize = v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetStorageType(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.StorageType = &v
	return s
}

type DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount struct {
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	Step     *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetMaxCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetMinCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetStep(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.Step = &v
	return s
}

type DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize struct {
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	Step     *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetMaxCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetMinCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetStep(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.Step = &v
	return s
}

type DescribeAvailableResourcesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAvailableResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponse) SetHeaders(v map[string]*string) *DescribeAvailableResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableResourcesResponse) SetStatusCode(v int32) *DescribeAvailableResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableResourcesResponse) SetBody(v *DescribeAvailableResourcesResponseBody) *DescribeAvailableResourcesResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetDBInstanceId(v string) *DescribeBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	BackupRetentionPeriod *int32  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	EnableRecoveryPoint   *bool   `json:"EnableRecoveryPoint,omitempty" xml:"EnableRecoveryPoint,omitempty"`
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupTime   *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	RecoveryPointPeriod   *string `json:"RecoveryPointPeriod,omitempty" xml:"RecoveryPointPeriod,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetEnableRecoveryPoint(v bool) *DescribeBackupPolicyResponseBody {
	s.EnableRecoveryPoint = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRecoveryPointPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.RecoveryPointPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPolicyResponse) SetStatusCode(v int32) *DescribeBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeDBClusterNodeRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s DescribeDBClusterNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeRequest) SetDBInstanceId(v string) *DescribeDBClusterNodeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBClusterNodeRequest) SetNodeType(v string) *DescribeDBClusterNodeRequest {
	s.NodeType = &v
	return s
}

type DescribeDBClusterNodeResponseBody struct {
	DBClusterId *string                                   `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Nodes       []*DescribeDBClusterNodeResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeResponseBody) SetDBClusterId(v string) *DescribeDBClusterNodeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterNodeResponseBody) SetNodes(v []*DescribeDBClusterNodeResponseBodyNodes) *DescribeDBClusterNodeResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribeDBClusterNodeResponseBody) SetRequestId(v string) *DescribeDBClusterNodeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterNodeResponseBodyNodes struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDBClusterNodeResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNodeResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeResponseBodyNodes) SetName(v string) *DescribeDBClusterNodeResponseBodyNodes {
	s.Name = &v
	return s
}

type DescribeDBClusterNodeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBClusterNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeResponse) SetHeaders(v map[string]*string) *DescribeDBClusterNodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterNodeResponse) SetStatusCode(v int32) *DescribeDBClusterNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterNodeResponse) SetBody(v *DescribeDBClusterNodeResponseBody) *DescribeDBClusterNodeResponse {
	s.Body = v
	return s
}

type DescribeDBClusterPerformanceRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceRequest) SetDBInstanceId(v string) *DescribeDBClusterPerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetEndTime(v string) *DescribeDBClusterPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetKey(v string) *DescribeDBClusterPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetNodeType(v string) *DescribeDBClusterPerformanceRequest {
	s.NodeType = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetNodes(v string) *DescribeDBClusterPerformanceRequest {
	s.Nodes = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeDBClusterPerformanceResponseBody struct {
	DBClusterId     *string                                                    `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime         *string                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PerformanceKeys []*DescribeDBClusterPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Repeated"`
	RequestId       *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetPerformanceKeys(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeys) *DescribeDBClusterPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetRequestId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetStartTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeys struct {
	Name   *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	Series []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	Unit   *string                                                          `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) SetSeries(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) *DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	s.Series = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) SetUnit(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	s.Unit = &v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries struct {
	Name   *string                                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Role   *string                                                                `json:"Role,omitempty" xml:"Role,omitempty"`
	Values []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) SetRole(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries {
	s.Role = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) SetValues(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries {
	s.Values = v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues struct {
	Point []*string `json:"Point,omitempty" xml:"Point,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) SetPoint(v []*string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues {
	s.Point = v
	return s
}

type DescribeDBClusterPerformanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBClusterPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBClusterPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetStatusCode(v int32) *DescribeDBClusterPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetBody(v *DescribeDBClusterPerformanceResponseBody) *DescribeDBClusterPerformanceResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceAttributeRequest struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DescribeDBInstanceAttributeRequest) SetOwnerId(v int64) *DescribeDBInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetResourceGroupId(v string) *DescribeDBInstanceAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBody struct {
	Items     *DescribeDBInstanceAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBody) SetItems(v *DescribeDBInstanceAttributeResponseBodyItems) *DescribeDBInstanceAttributeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyItems struct {
	DBInstanceAttribute []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute `json:"DBInstanceAttribute,omitempty" xml:"DBInstanceAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItems) SetDBInstanceAttribute(v []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) *DescribeDBInstanceAttributeResponseBodyItems {
	s.DBInstanceAttribute = v
	return s
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute struct {
	AvailabilityValue     *string                                                              `json:"AvailabilityValue,omitempty" xml:"AvailabilityValue,omitempty"`
	ConnectionMode        *string                                                              `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	ConnectionString      *string                                                              `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	CoreVersion           *string                                                              `json:"CoreVersion,omitempty" xml:"CoreVersion,omitempty"`
	CpuCores              *int32                                                               `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	CpuCoresPerNode       *int32                                                               `json:"CpuCoresPerNode,omitempty" xml:"CpuCoresPerNode,omitempty"`
	CreationTime          *string                                                              `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DBInstanceCategory    *string                                                              `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
	DBInstanceClass       *string                                                              `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceClassType   *string                                                              `json:"DBInstanceClassType,omitempty" xml:"DBInstanceClassType,omitempty"`
	DBInstanceCpuCores    *int32                                                               `json:"DBInstanceCpuCores,omitempty" xml:"DBInstanceCpuCores,omitempty"`
	DBInstanceDescription *string                                                              `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceDiskMBPS    *int64                                                               `json:"DBInstanceDiskMBPS,omitempty" xml:"DBInstanceDiskMBPS,omitempty"`
	DBInstanceGroupCount  *string                                                              `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	DBInstanceId          *string                                                              `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceMemory      *int64                                                               `json:"DBInstanceMemory,omitempty" xml:"DBInstanceMemory,omitempty"`
	DBInstanceMode        *string                                                              `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	DBInstanceNetType     *string                                                              `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	DBInstanceStatus      *string                                                              `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceStorage     *int64                                                               `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	EncryptionKey         *string                                                              `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	EncryptionType        *string                                                              `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	Engine                *string                                                              `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string                                                              `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime            *string                                                              `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HostType              *string                                                              `json:"HostType,omitempty" xml:"HostType,omitempty"`
	IdleTime              *int32                                                               `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	InstanceNetworkType   *string                                                              `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	LockMode              *string                                                              `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason            *string                                                              `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MaintainEndTime       *string                                                              `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime     *string                                                              `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	MasterNodeNum         *int32                                                               `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	MaxConnections        *int32                                                               `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	MemoryPerNode         *int32                                                               `json:"MemoryPerNode,omitempty" xml:"MemoryPerNode,omitempty"`
	MemorySize            *int64                                                               `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	MemoryUnit            *string                                                              `json:"MemoryUnit,omitempty" xml:"MemoryUnit,omitempty"`
	MinorVersion          *string                                                              `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	PayType               *string                                                              `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port                  *string                                                              `json:"Port,omitempty" xml:"Port,omitempty"`
	ReadDelayTime         *string                                                              `json:"ReadDelayTime,omitempty" xml:"ReadDelayTime,omitempty"`
	RegionId              *string                                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string                                                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RunningTime           *string                                                              `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	SecurityIPList        *string                                                              `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	SegNodeNum            *int32                                                               `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	SegmentCounts         *int32                                                               `json:"SegmentCounts,omitempty" xml:"SegmentCounts,omitempty"`
	ServerlessMode        *string                                                              `json:"ServerlessMode,omitempty" xml:"ServerlessMode,omitempty"`
	ServerlessResource    *int32                                                               `json:"ServerlessResource,omitempty" xml:"ServerlessResource,omitempty"`
	StartTime             *string                                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StoragePerNode        *int32                                                               `json:"StoragePerNode,omitempty" xml:"StoragePerNode,omitempty"`
	StorageSize           *int64                                                               `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	StorageType           *string                                                              `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	StorageUnit           *string                                                              `json:"StorageUnit,omitempty" xml:"StorageUnit,omitempty"`
	SupportRestore        *bool                                                                `json:"SupportRestore,omitempty" xml:"SupportRestore,omitempty"`
	Tags                  *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// vSwitch ID。
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	VpcId  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetAvailabilityValue(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.AvailabilityValue = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetConnectionMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetConnectionString(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCoreVersion(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CoreVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCpuCores(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCpuCoresPerNode(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CpuCoresPerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCreationTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceCategory(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceCategory = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceClass(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceClassType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceClassType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceCpuCores(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceCpuCores = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceDescription(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceDiskMBPS(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceDiskMBPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceGroupCount(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceGroupCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceMemory(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceMemory = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceNetType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceStatus(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceStorage(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetEncryptionKey(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetEncryptionType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.EncryptionType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetEngine(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetEngineVersion(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetExpireTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetHostType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.HostType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetIdleTime(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.IdleTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetInstanceNetworkType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetLockMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetLockReason(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaintainEndTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaintainStartTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMasterNodeNum(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MasterNodeNum = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaxConnections(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMemoryPerNode(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MemoryPerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMemorySize(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MemorySize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMemoryUnit(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MemoryUnit = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMinorVersion(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetPayType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetPort(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetReadDelayTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ReadDelayTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetResourceGroupId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetRunningTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.RunningTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSecurityIPList(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSegNodeNum(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SegNodeNum = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSegmentCounts(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SegmentCounts = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetServerlessMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ServerlessMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetServerlessResource(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ServerlessResource = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStartTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStoragePerNode(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StoragePerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStorageSize(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStorageType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStorageUnit(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageUnit = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSupportRestore(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SupportRestore = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetTags(v *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Tags = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetVpcId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ZoneId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags struct {
	Tag []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) SetTag(v []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags {
	s.Tag = v
	return s
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) SetKey(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) SetValue(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag {
	s.Value = &v
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

type DescribeDBInstanceDataBloatRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDBInstanceDataBloatRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataBloatRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataBloatRequest) SetDBInstanceId(v string) *DescribeDBInstanceDataBloatRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDataBloatRequest) SetPageNumber(v int32) *DescribeDBInstanceDataBloatRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDataBloatRequest) SetPageSize(v int32) *DescribeDBInstanceDataBloatRequest {
	s.PageSize = &v
	return s
}

type DescribeDBInstanceDataBloatResponseBody struct {
	Items      []*DescribeDBInstanceDataBloatResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceDataBloatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataBloatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataBloatResponseBody) SetItems(v []*DescribeDBInstanceDataBloatResponseBodyItems) *DescribeDBInstanceDataBloatResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBody) SetPageNumber(v int32) *DescribeDBInstanceDataBloatResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBody) SetRequestId(v string) *DescribeDBInstanceDataBloatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBody) SetTotalCount(v int32) *DescribeDBInstanceDataBloatResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDBInstanceDataBloatResponseBodyItems struct {
	BloatCeoff       *string `json:"BloatCeoff,omitempty" xml:"BloatCeoff,omitempty"`
	BloatSize        *string `json:"BloatSize,omitempty" xml:"BloatSize,omitempty"`
	DatabaseName     *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	ExpectTableSize  *string `json:"ExpectTableSize,omitempty" xml:"ExpectTableSize,omitempty"`
	RealTableSize    *string `json:"RealTableSize,omitempty" xml:"RealTableSize,omitempty"`
	SchemaName       *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	Sequence         *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	StorageType      *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	SuggestedAction  *string `json:"SuggestedAction,omitempty" xml:"SuggestedAction,omitempty"`
	TableName        *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TimeLastUpdated  *string `json:"TimeLastUpdated,omitempty" xml:"TimeLastUpdated,omitempty"`
	TimeLastVacuumed *string `json:"TimeLastVacuumed,omitempty" xml:"TimeLastVacuumed,omitempty"`
}

func (s DescribeDBInstanceDataBloatResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataBloatResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetBloatCeoff(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.BloatCeoff = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetBloatSize(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.BloatSize = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetDatabaseName(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetExpectTableSize(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.ExpectTableSize = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetRealTableSize(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.RealTableSize = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetSchemaName(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetSequence(v int32) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.Sequence = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetStorageType(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetSuggestedAction(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.SuggestedAction = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetTableName(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetTimeLastUpdated(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.TimeLastUpdated = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetTimeLastVacuumed(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.TimeLastVacuumed = &v
	return s
}

type DescribeDBInstanceDataBloatResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceDataBloatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceDataBloatResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataBloatResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataBloatResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceDataBloatResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceDataBloatResponse) SetStatusCode(v int32) *DescribeDBInstanceDataBloatResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponse) SetBody(v *DescribeDBInstanceDataBloatResponseBody) *DescribeDBInstanceDataBloatResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceDataSkewRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDBInstanceDataSkewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataSkewRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSkewRequest) SetDBInstanceId(v string) *DescribeDBInstanceDataSkewRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDataSkewRequest) SetPageNumber(v int32) *DescribeDBInstanceDataSkewRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDataSkewRequest) SetPageSize(v int32) *DescribeDBInstanceDataSkewRequest {
	s.PageSize = &v
	return s
}

type DescribeDBInstanceDataSkewResponseBody struct {
	Items      []*DescribeDBInstanceDataSkewResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceDataSkewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataSkewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSkewResponseBody) SetItems(v []*DescribeDBInstanceDataSkewResponseBodyItems) *DescribeDBInstanceDataSkewResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBody) SetPageNumber(v int32) *DescribeDBInstanceDataSkewResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBody) SetRequestId(v string) *DescribeDBInstanceDataSkewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBody) SetTotalCount(v int32) *DescribeDBInstanceDataSkewResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDBInstanceDataSkewResponseBodyItems struct {
	DatabaseName    *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	DistributeKey   *string `json:"DistributeKey,omitempty" xml:"DistributeKey,omitempty"`
	Owner           *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	SchemaName      *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	Sequence        *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	TableName       *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableSize       *string `json:"TableSize,omitempty" xml:"TableSize,omitempty"`
	TableSkew       *string `json:"TableSkew,omitempty" xml:"TableSkew,omitempty"`
	TimeLastUpdated *string `json:"TimeLastUpdated,omitempty" xml:"TimeLastUpdated,omitempty"`
}

func (s DescribeDBInstanceDataSkewResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataSkewResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetDatabaseName(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetDistributeKey(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.DistributeKey = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetOwner(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetSchemaName(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetSequence(v int32) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.Sequence = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetTableName(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetTableSize(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.TableSize = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetTableSkew(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.TableSkew = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetTimeLastUpdated(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.TimeLastUpdated = &v
	return s
}

type DescribeDBInstanceDataSkewResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceDataSkewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceDataSkewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataSkewResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSkewResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceDataSkewResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceDataSkewResponse) SetStatusCode(v int32) *DescribeDBInstanceDataSkewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponse) SetBody(v *DescribeDBInstanceDataSkewResponseBody) *DescribeDBInstanceDataSkewResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceDiagnosisSummaryRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RolePreferd  *string `json:"RolePreferd,omitempty" xml:"RolePreferd,omitempty"`
	StartStatus  *string `json:"StartStatus,omitempty" xml:"StartStatus,omitempty"`
	SyncMode     *string `json:"SyncMode,omitempty" xml:"SyncMode,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetDBInstanceId(v string) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetPageNumber(v int32) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetPageSize(v int32) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetRolePreferd(v string) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.RolePreferd = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetStartStatus(v string) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.StartStatus = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetSyncMode(v string) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.SyncMode = &v
	return s
}

type DescribeDBInstanceDiagnosisSummaryResponseBody struct {
	Items             []*DescribeDBInstanceDiagnosisSummaryResponseBodyItems           `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MasterStatusInfo  *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo  `json:"MasterStatusInfo,omitempty" xml:"MasterStatusInfo,omitempty" type:"Struct"`
	PageNumber        *string                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId         *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SegmentStatusInfo *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo `json:"SegmentStatusInfo,omitempty" xml:"SegmentStatusInfo,omitempty" type:"Struct"`
	TotalCount        *string                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetItems(v []*DescribeDBInstanceDiagnosisSummaryResponseBodyItems) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetMasterStatusInfo(v *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.MasterStatusInfo = v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetPageNumber(v string) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetRequestId(v string) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetSegmentStatusInfo(v *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.SegmentStatusInfo = v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetTotalCount(v string) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDBInstanceDiagnosisSummaryResponseBodyItems struct {
	Hostname            *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	NodeAddress         *string `json:"NodeAddress,omitempty" xml:"NodeAddress,omitempty"`
	NodeCID             *string `json:"NodeCID,omitempty" xml:"NodeCID,omitempty"`
	NodeID              *string `json:"NodeID,omitempty" xml:"NodeID,omitempty"`
	NodeName            *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodePort            *string `json:"NodePort,omitempty" xml:"NodePort,omitempty"`
	NodePreferredRole   *string `json:"NodePreferredRole,omitempty" xml:"NodePreferredRole,omitempty"`
	NodeReplicationMode *string `json:"NodeReplicationMode,omitempty" xml:"NodeReplicationMode,omitempty"`
	NodeRole            *string `json:"NodeRole,omitempty" xml:"NodeRole,omitempty"`
	NodeStatus          *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	NodeType            *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetHostname(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.Hostname = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeAddress(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeAddress = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeCID(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeCID = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeID(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeID = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeName(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeName = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodePort(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodePort = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodePreferredRole(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodePreferredRole = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeReplicationMode(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeReplicationMode = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeRole(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeRole = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeStatus(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeStatus = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeType(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeType = &v
	return s
}

type DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo struct {
	ExceptionNodeNum    *int32 `json:"ExceptionNodeNum,omitempty" xml:"ExceptionNodeNum,omitempty"`
	NormalNodeNum       *int32 `json:"NormalNodeNum,omitempty" xml:"NormalNodeNum,omitempty"`
	NotPreferredNodeNum *int32 `json:"NotPreferredNodeNum,omitempty" xml:"NotPreferredNodeNum,omitempty"`
	NotSyncingNodeNum   *int32 `json:"NotSyncingNodeNum,omitempty" xml:"NotSyncingNodeNum,omitempty"`
	PreferredNodeNum    *int32 `json:"PreferredNodeNum,omitempty" xml:"PreferredNodeNum,omitempty"`
	SyncedNodeNum       *int32 `json:"SyncedNodeNum,omitempty" xml:"SyncedNodeNum,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetExceptionNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.ExceptionNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetNormalNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.NormalNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetNotPreferredNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.NotPreferredNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetNotSyncingNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.NotSyncingNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetPreferredNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.PreferredNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetSyncedNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.SyncedNodeNum = &v
	return s
}

type DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo struct {
	ExceptionNodeNum    *int32 `json:"ExceptionNodeNum,omitempty" xml:"ExceptionNodeNum,omitempty"`
	NormalNodeNum       *int32 `json:"NormalNodeNum,omitempty" xml:"NormalNodeNum,omitempty"`
	NotPreferredNodeNum *int32 `json:"NotPreferredNodeNum,omitempty" xml:"NotPreferredNodeNum,omitempty"`
	NotSyncingNodeNum   *int32 `json:"NotSyncingNodeNum,omitempty" xml:"NotSyncingNodeNum,omitempty"`
	PreferredNodeNum    *int32 `json:"PreferredNodeNum,omitempty" xml:"PreferredNodeNum,omitempty"`
	SyncedNodeNum       *int32 `json:"SyncedNodeNum,omitempty" xml:"SyncedNodeNum,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetExceptionNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.ExceptionNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetNormalNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.NormalNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetNotPreferredNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.NotPreferredNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetNotSyncingNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.NotSyncingNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetPreferredNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.PreferredNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetSyncedNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.SyncedNodeNum = &v
	return s
}

type DescribeDBInstanceDiagnosisSummaryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceDiagnosisSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceDiagnosisSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponse) SetStatusCode(v int32) *DescribeDBInstanceDiagnosisSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponse) SetBody(v *DescribeDBInstanceDiagnosisSummaryResponseBody) *DescribeDBInstanceDiagnosisSummaryResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceErrorLogRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database     *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Host         *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Keywords     *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	LogLevel     *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	User         *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDBInstanceErrorLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceErrorLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceErrorLogRequest) SetDBInstanceId(v string) *DescribeDBInstanceErrorLogRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetDatabase(v string) *DescribeDBInstanceErrorLogRequest {
	s.Database = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetEndTime(v string) *DescribeDBInstanceErrorLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetHost(v string) *DescribeDBInstanceErrorLogRequest {
	s.Host = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetKeywords(v string) *DescribeDBInstanceErrorLogRequest {
	s.Keywords = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetLogLevel(v string) *DescribeDBInstanceErrorLogRequest {
	s.LogLevel = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetPageNumber(v int32) *DescribeDBInstanceErrorLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetPageSize(v int32) *DescribeDBInstanceErrorLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetStartTime(v string) *DescribeDBInstanceErrorLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetUser(v string) *DescribeDBInstanceErrorLogRequest {
	s.User = &v
	return s
}

type DescribeDBInstanceErrorLogResponseBody struct {
	Items      []*DescribeDBInstanceErrorLogResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceErrorLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceErrorLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceErrorLogResponseBody) SetItems(v []*DescribeDBInstanceErrorLogResponseBodyItems) *DescribeDBInstanceErrorLogResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBody) SetPageNumber(v int32) *DescribeDBInstanceErrorLogResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBody) SetRequestId(v string) *DescribeDBInstanceErrorLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBody) SetTotalCount(v int32) *DescribeDBInstanceErrorLogResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDBInstanceErrorLogResponseBodyItems struct {
	Database   *string `json:"Database,omitempty" xml:"Database,omitempty"`
	Host       *string `json:"Host,omitempty" xml:"Host,omitempty"`
	LogContext *string `json:"LogContext,omitempty" xml:"LogContext,omitempty"`
	LogLevel   *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	Time       *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	User       *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDBInstanceErrorLogResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceErrorLogResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetDatabase(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.Database = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetHost(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.Host = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetLogContext(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.LogContext = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetLogLevel(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.LogLevel = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetTime(v int64) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.Time = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetUser(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.User = &v
	return s
}

type DescribeDBInstanceErrorLogResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceErrorLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceErrorLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceErrorLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceErrorLogResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceErrorLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceErrorLogResponse) SetStatusCode(v int32) *DescribeDBInstanceErrorLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponse) SetBody(v *DescribeDBInstanceErrorLogResponseBody) *DescribeDBInstanceErrorLogResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceIPArrayListRequest struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDBInstanceIPArrayListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListRequest) SetDBInstanceId(v string) *DescribeDBInstanceIPArrayListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListRequest) SetResourceGroupId(v string) *DescribeDBInstanceIPArrayListRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDBInstanceIPArrayListResponseBody struct {
	Items     *DescribeDBInstanceIPArrayListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceIPArrayListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponseBody) SetItems(v *DescribeDBInstanceIPArrayListResponseBodyItems) *DescribeDBInstanceIPArrayListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseBody) SetRequestId(v string) *DescribeDBInstanceIPArrayListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceIPArrayListResponseBodyItems struct {
	DBInstanceIPArray []*DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray `json:"DBInstanceIPArray,omitempty" xml:"DBInstanceIPArray,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceIPArrayListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItems) SetDBInstanceIPArray(v []*DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) *DescribeDBInstanceIPArrayListResponseBodyItems {
	s.DBInstanceIPArray = v
	return s
}

type DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray struct {
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty"`
	DBInstanceIPArrayName      *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	SecurityIPList             *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) SetDBInstanceIPArrayAttribute(v string) *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray {
	s.DBInstanceIPArrayAttribute = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) SetDBInstanceIPArrayName(v string) *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray {
	s.DBInstanceIPArrayName = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) SetSecurityIPList(v string) *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray {
	s.SecurityIPList = &v
	return s
}

type DescribeDBInstanceIPArrayListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceIPArrayListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceIPArrayListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceIPArrayListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponse) SetStatusCode(v int32) *DescribeDBInstanceIPArrayListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponse) SetBody(v *DescribeDBInstanceIPArrayListResponseBody) *DescribeDBInstanceIPArrayListResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceIndexUsageRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDBInstanceIndexUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIndexUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIndexUsageRequest) SetDBInstanceId(v string) *DescribeDBInstanceIndexUsageRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageRequest) SetPageNumber(v int32) *DescribeDBInstanceIndexUsageRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageRequest) SetPageSize(v int32) *DescribeDBInstanceIndexUsageRequest {
	s.PageSize = &v
	return s
}

type DescribeDBInstanceIndexUsageResponseBody struct {
	Items      []*DescribeDBInstanceIndexUsageResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceIndexUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIndexUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIndexUsageResponseBody) SetItems(v []*DescribeDBInstanceIndexUsageResponseBodyItems) *DescribeDBInstanceIndexUsageResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBody) SetPageNumber(v int32) *DescribeDBInstanceIndexUsageResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBody) SetRequestId(v string) *DescribeDBInstanceIndexUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBody) SetTotalCount(v int32) *DescribeDBInstanceIndexUsageResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDBInstanceIndexUsageResponseBodyItems struct {
	DatabaseName     *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	IndexDef         *string `json:"IndexDef,omitempty" xml:"IndexDef,omitempty"`
	IndexName        *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	IndexScanTimes   *int32  `json:"IndexScanTimes,omitempty" xml:"IndexScanTimes,omitempty"`
	IndexSize        *string `json:"IndexSize,omitempty" xml:"IndexSize,omitempty"`
	IsPartitionTable *bool   `json:"IsPartitionTable,omitempty" xml:"IsPartitionTable,omitempty"`
	ParentTableName  *string `json:"ParentTableName,omitempty" xml:"ParentTableName,omitempty"`
	SchemaName       *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName        *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TimeLastUpdated  *string `json:"TimeLastUpdated,omitempty" xml:"TimeLastUpdated,omitempty"`
}

func (s DescribeDBInstanceIndexUsageResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIndexUsageResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetDatabaseName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIndexDef(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IndexDef = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIndexName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IndexName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIndexScanTimes(v int32) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IndexScanTimes = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIndexSize(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IndexSize = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIsPartitionTable(v bool) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IsPartitionTable = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetParentTableName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.ParentTableName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetSchemaName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetTableName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetTimeLastUpdated(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.TimeLastUpdated = &v
	return s
}

type DescribeDBInstanceIndexUsageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceIndexUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceIndexUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIndexUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIndexUsageResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceIndexUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponse) SetStatusCode(v int32) *DescribeDBInstanceIndexUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponse) SetBody(v *DescribeDBInstanceIndexUsageResponseBody) *DescribeDBInstanceIndexUsageResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceNetInfoRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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

type DescribeDBInstanceNetInfoResponseBody struct {
	DBInstanceNetInfos  *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos `json:"DBInstanceNetInfos,omitempty" xml:"DBInstanceNetInfos,omitempty" type:"Struct"`
	InstanceNetworkType *string                                                  `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	RequestId           *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetDBInstanceNetInfos(v *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) *DescribeDBInstanceNetInfoResponseBody {
	s.DBInstanceNetInfos = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos struct {
	DBInstanceNetInfo []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" xml:"DBInstanceNetInfo,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetDBInstanceNetInfo(v []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.DBInstanceNetInfo = v
	return s
}

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo struct {
	AddressType      *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	IPAddress        *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	IPType           *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// VPC ID。
	VPCId         *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId     *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetAddressType(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.AddressType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetConnectionString(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetIPAddress(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetIPType(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.IPType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetPort(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVPCId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVSwitchId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVpcInstanceId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VpcInstanceId = &v
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

type DescribeDBInstanceOnECSAttributeRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDBInstanceOnECSAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeRequest) SetDBInstanceId(v string) *DescribeDBInstanceOnECSAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeRequest) SetOwnerId(v int64) *DescribeDBInstanceOnECSAttributeRequest {
	s.OwnerId = &v
	return s
}

type DescribeDBInstanceOnECSAttributeResponseBody struct {
	Items     *DescribeDBInstanceOnECSAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceOnECSAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponseBody) SetItems(v *DescribeDBInstanceOnECSAttributeResponseBodyItems) *DescribeDBInstanceOnECSAttributeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBody) SetRequestId(v string) *DescribeDBInstanceOnECSAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceOnECSAttributeResponseBodyItems struct {
	DBInstanceAttribute []*DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute `json:"DBInstanceAttribute,omitempty" xml:"DBInstanceAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItems) SetDBInstanceAttribute(v []*DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) *DescribeDBInstanceOnECSAttributeResponseBodyItems {
	s.DBInstanceAttribute = v
	return s
}

type DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute struct {
	ConnectionMode        *string                                                                   `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	ConnectionString      *string                                                                   `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	CpuCores              *int32                                                                    `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	CreationTime          *string                                                                   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DBInstanceCategory    *string                                                                   `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
	DBInstanceClass       *string                                                                   `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceDescription *string                                                                   `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId          *string                                                                   `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStatus      *string                                                                   `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	EncryptionKey         *string                                                                   `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	EncryptionType        *string                                                                   `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	Engine                *string                                                                   `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string                                                                   `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime            *string                                                                   `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceDeployType    *string                                                                   `json:"InstanceDeployType,omitempty" xml:"InstanceDeployType,omitempty"`
	InstanceNetworkType   *string                                                                   `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	LockMode              *string                                                                   `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MasterNodeNum         *int32                                                                    `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	MemorySize            *int32                                                                    `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	MinorVersion          *string                                                                   `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	PayType               *string                                                                   `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port                  *string                                                                   `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId              *string                                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SegNodeNum            *int32                                                                    `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	StorageSize           *int32                                                                    `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	StorageType           *string                                                                   `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	SupportRestore        *bool                                                                     `json:"SupportRestore,omitempty" xml:"SupportRestore,omitempty"`
	Tags                  *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// vSwitch ID。
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	VpcId  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetConnectionMode(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetConnectionString(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetCpuCores(v int32) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetCreationTime(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceCategory(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceCategory = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceClass(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceDescription(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceId(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceStatus(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetEncryptionKey(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetEncryptionType(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.EncryptionType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetEngine(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetEngineVersion(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetExpireTime(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetInstanceDeployType(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.InstanceDeployType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetInstanceNetworkType(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetLockMode(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetMasterNodeNum(v int32) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.MasterNodeNum = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetMemorySize(v int32) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.MemorySize = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetMinorVersion(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetPayType(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetPort(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetRegionId(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetSegNodeNum(v int32) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.SegNodeNum = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetStorageSize(v int32) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetStorageType(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetSupportRestore(v bool) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.SupportRestore = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetTags(v *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTags) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.Tags = v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetVSwitchId(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetVpcId(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetZoneId(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.ZoneId = &v
	return s
}

type DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTags struct {
	Tag []*DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTags) SetTag(v []*DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTags {
	s.Tag = v
	return s
}

type DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag) SetKey(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag) SetValue(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag {
	s.Value = &v
	return s
}

type DescribeDBInstanceOnECSAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceOnECSAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceOnECSAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceOnECSAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponse) SetStatusCode(v int32) *DescribeDBInstanceOnECSAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponse) SetBody(v *DescribeDBInstanceOnECSAttributeResponseBody) *DescribeDBInstanceOnECSAttributeResponse {
	s.Body = v
	return s
}

type DescribeDBInstancePerformanceRequest struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Key             *string `json:"Key,omitempty" xml:"Key,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstancePerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceRequest) SetDBInstanceId(v string) *DescribeDBInstancePerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetEndTime(v string) *DescribeDBInstancePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetKey(v string) *DescribeDBInstancePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetResourceGroupId(v string) *DescribeDBInstancePerformanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetStartTime(v string) *DescribeDBInstancePerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeDBInstancePerformanceResponseBody struct {
	DBInstanceId    *string   `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EndTime         *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Engine          *string   `json:"Engine,omitempty" xml:"Engine,omitempty"`
	PerformanceKeys []*string `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Repeated"`
	RequestId       *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstancePerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBody) SetDBInstanceId(v string) *DescribeDBInstancePerformanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetEndTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetEngine(v string) *DescribeDBInstancePerformanceResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetPerformanceKeys(v []*string) *DescribeDBInstancePerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetRequestId(v string) *DescribeDBInstancePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetStartTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDBInstancePerformanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstancePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstancePerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBInstancePerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancePerformanceResponse) SetStatusCode(v int32) *DescribeDBInstancePerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponse) SetBody(v *DescribeDBInstancePerformanceResponseBody) *DescribeDBInstancePerformanceResponse {
	s.Body = v
	return s
}

type DescribeDBInstancePlansRequest struct {
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlanCreateDate   *string `json:"PlanCreateDate,omitempty" xml:"PlanCreateDate,omitempty"`
	PlanDesc         *string `json:"PlanDesc,omitempty" xml:"PlanDesc,omitempty"`
	PlanId           *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PlanScheduleType *string `json:"PlanScheduleType,omitempty" xml:"PlanScheduleType,omitempty"`
	PlanType         *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
}

func (s DescribeDBInstancePlansRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansRequest) SetDBInstanceId(v string) *DescribeDBInstancePlansRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetOwnerId(v int64) *DescribeDBInstancePlansRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanCreateDate(v string) *DescribeDBInstancePlansRequest {
	s.PlanCreateDate = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanDesc(v string) *DescribeDBInstancePlansRequest {
	s.PlanDesc = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanId(v string) *DescribeDBInstancePlansRequest {
	s.PlanId = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanScheduleType(v string) *DescribeDBInstancePlansRequest {
	s.PlanScheduleType = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanType(v string) *DescribeDBInstancePlansRequest {
	s.PlanType = &v
	return s
}

type DescribeDBInstancePlansResponseBody struct {
	ErrorMessage     *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Items            *DescribeDBInstancePlansResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber       *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                    `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status           *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalRecordCount *int32                                    `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstancePlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansResponseBody) SetErrorMessage(v string) *DescribeDBInstancePlansResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetItems(v *DescribeDBInstancePlansResponseBodyItems) *DescribeDBInstancePlansResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetPageNumber(v int32) *DescribeDBInstancePlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetPageRecordCount(v int32) *DescribeDBInstancePlansResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetRequestId(v string) *DescribeDBInstancePlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetStatus(v string) *DescribeDBInstancePlansResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetTotalRecordCount(v int32) *DescribeDBInstancePlansResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeDBInstancePlansResponseBodyItems struct {
	PlanList []*DescribeDBInstancePlansResponseBodyItemsPlanList `json:"PlanList,omitempty" xml:"PlanList,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancePlansResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePlansResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansResponseBodyItems) SetPlanList(v []*DescribeDBInstancePlansResponseBodyItemsPlanList) *DescribeDBInstancePlansResponseBodyItems {
	s.PlanList = v
	return s
}

type DescribeDBInstancePlansResponseBodyItemsPlanList struct {
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	PlanConfig       *string `json:"PlanConfig,omitempty" xml:"PlanConfig,omitempty"`
	PlanDesc         *string `json:"PlanDesc,omitempty" xml:"PlanDesc,omitempty"`
	PlanEndDate      *string `json:"PlanEndDate,omitempty" xml:"PlanEndDate,omitempty"`
	PlanId           *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PlanName         *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	PlanScheduleType *string `json:"PlanScheduleType,omitempty" xml:"PlanScheduleType,omitempty"`
	PlanStartDate    *string `json:"PlanStartDate,omitempty" xml:"PlanStartDate,omitempty"`
	PlanStatus       *string `json:"PlanStatus,omitempty" xml:"PlanStatus,omitempty"`
	PlanType         *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
}

func (s DescribeDBInstancePlansResponseBodyItemsPlanList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePlansResponseBodyItemsPlanList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetDBInstanceId(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanConfig(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanConfig = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanDesc(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanDesc = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanEndDate(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanEndDate = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanId(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanId = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanName(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanName = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanScheduleType(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanScheduleType = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanStartDate(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanStartDate = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanStatus(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanStatus = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanType(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanType = &v
	return s
}

type DescribeDBInstancePlansResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstancePlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstancePlansResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansResponse) SetHeaders(v map[string]*string) *DescribeDBInstancePlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancePlansResponse) SetStatusCode(v int32) *DescribeDBInstancePlansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancePlansResponse) SetBody(v *DescribeDBInstancePlansResponseBody) *DescribeDBInstancePlansResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceSQLPatternsRequest struct {
	DBInstanceId  *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database      *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	QueryKeywords *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	SourceIP      *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	User          *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDBInstanceSQLPatternsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSQLPatternsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSQLPatternsRequest) SetDBInstanceId(v string) *DescribeDBInstanceSQLPatternsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsRequest) SetDatabase(v string) *DescribeDBInstanceSQLPatternsRequest {
	s.Database = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsRequest) SetEndTime(v string) *DescribeDBInstanceSQLPatternsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsRequest) SetQueryKeywords(v string) *DescribeDBInstanceSQLPatternsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsRequest) SetSourceIP(v string) *DescribeDBInstanceSQLPatternsRequest {
	s.SourceIP = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsRequest) SetStartTime(v string) *DescribeDBInstanceSQLPatternsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsRequest) SetUser(v string) *DescribeDBInstanceSQLPatternsRequest {
	s.User = &v
	return s
}

type DescribeDBInstanceSQLPatternsResponseBody struct {
	DBClusterId *string                                              `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime     *string                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Patterns    []*DescribeDBInstanceSQLPatternsResponseBodyPatterns `json:"Patterns,omitempty" xml:"Patterns,omitempty" type:"Repeated"`
	RequestId   *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime   *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstanceSQLPatternsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSQLPatternsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSQLPatternsResponseBody) SetDBClusterId(v string) *DescribeDBInstanceSQLPatternsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsResponseBody) SetEndTime(v string) *DescribeDBInstanceSQLPatternsResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsResponseBody) SetPatterns(v []*DescribeDBInstanceSQLPatternsResponseBodyPatterns) *DescribeDBInstanceSQLPatternsResponseBody {
	s.Patterns = v
	return s
}

func (s *DescribeDBInstanceSQLPatternsResponseBody) SetRequestId(v string) *DescribeDBInstanceSQLPatternsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsResponseBody) SetStartTime(v string) *DescribeDBInstanceSQLPatternsResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDBInstanceSQLPatternsResponseBodyPatterns struct {
	Name   *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Values map[string]interface{} `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeDBInstanceSQLPatternsResponseBodyPatterns) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSQLPatternsResponseBodyPatterns) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSQLPatternsResponseBodyPatterns) SetName(v string) *DescribeDBInstanceSQLPatternsResponseBodyPatterns {
	s.Name = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsResponseBodyPatterns) SetValues(v map[string]interface{}) *DescribeDBInstanceSQLPatternsResponseBodyPatterns {
	s.Values = v
	return s
}

type DescribeDBInstanceSQLPatternsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceSQLPatternsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceSQLPatternsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSQLPatternsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSQLPatternsResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceSQLPatternsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceSQLPatternsResponse) SetStatusCode(v int32) *DescribeDBInstanceSQLPatternsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsResponse) SetBody(v *DescribeDBInstanceSQLPatternsResponseBody) *DescribeDBInstanceSQLPatternsResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceSSLRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDBInstanceSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLRequest) SetDBInstanceId(v string) *DescribeDBInstanceSSLRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDBInstanceSSLResponseBody struct {
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SSLEnabled     *bool   `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	SSLExpiredTime *string `json:"SSLExpiredTime,omitempty" xml:"SSLExpiredTime,omitempty"`
}

func (s DescribeDBInstanceSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponseBody) SetCertCommonName(v string) *DescribeDBInstanceSSLResponseBody {
	s.CertCommonName = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetDBInstanceId(v string) *DescribeDBInstanceSSLResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetDBInstanceName(v string) *DescribeDBInstanceSSLResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetRequestId(v string) *DescribeDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetSSLEnabled(v bool) *DescribeDBInstanceSSLResponseBody {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetSSLExpiredTime(v string) *DescribeDBInstanceSSLResponseBody {
	s.SSLExpiredTime = &v
	return s
}

type DescribeDBInstanceSSLResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceSSLResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceSSLResponse) SetStatusCode(v int32) *DescribeDBInstanceSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceSSLResponse) SetBody(v *DescribeDBInstanceSSLResponseBody) *DescribeDBInstanceSSLResponse {
	s.Body = v
	return s
}

type DescribeDBInstancesRequest struct {
	DBInstanceCategories  []*string                        `json:"DBInstanceCategories,omitempty" xml:"DBInstanceCategories,omitempty" type:"Repeated"`
	DBInstanceDescription *string                          `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceIds         *string                          `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	DBInstanceModes       []*string                        `json:"DBInstanceModes,omitempty" xml:"DBInstanceModes,omitempty" type:"Repeated"`
	DBInstanceStatuses    []*string                        `json:"DBInstanceStatuses,omitempty" xml:"DBInstanceStatuses,omitempty" type:"Repeated"`
	InstanceDeployTypes   []*string                        `json:"InstanceDeployTypes,omitempty" xml:"InstanceDeployTypes,omitempty" type:"Repeated"`
	InstanceNetworkType   *string                          `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerId               *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber            *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId              *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag                   []*DescribeDBInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) SetDBInstanceCategories(v []*string) *DescribeDBInstancesRequest {
	s.DBInstanceCategories = v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceDescription(v string) *DescribeDBInstancesRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceIds(v string) *DescribeDBInstancesRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceModes(v []*string) *DescribeDBInstancesRequest {
	s.DBInstanceModes = v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceStatuses(v []*string) *DescribeDBInstancesRequest {
	s.DBInstanceStatuses = v
	return s
}

func (s *DescribeDBInstancesRequest) SetInstanceDeployTypes(v []*string) *DescribeDBInstancesRequest {
	s.InstanceDeployTypes = v
	return s
}

func (s *DescribeDBInstancesRequest) SetInstanceNetworkType(v string) *DescribeDBInstancesRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetOwnerId(v int64) *DescribeDBInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageNumber(v int32) *DescribeDBInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageSize(v int32) *DescribeDBInstancesRequest {
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

func (s *DescribeDBInstancesRequest) SetTag(v []*DescribeDBInstancesRequestTag) *DescribeDBInstancesRequest {
	s.Tag = v
	return s
}

type DescribeDBInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequestTag) SetKey(v string) *DescribeDBInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesRequestTag) SetValue(v string) *DescribeDBInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeDBInstancesShrinkRequest struct {
	DBInstanceCategoriesShrink *string                                `json:"DBInstanceCategories,omitempty" xml:"DBInstanceCategories,omitempty"`
	DBInstanceDescription      *string                                `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceIds              *string                                `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	DBInstanceModesShrink      *string                                `json:"DBInstanceModes,omitempty" xml:"DBInstanceModes,omitempty"`
	DBInstanceStatusesShrink   *string                                `json:"DBInstanceStatuses,omitempty" xml:"DBInstanceStatuses,omitempty"`
	InstanceDeployTypesShrink  *string                                `json:"InstanceDeployTypes,omitempty" xml:"InstanceDeployTypes,omitempty"`
	InstanceNetworkType        *string                                `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerId                    *int64                                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber                 *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId                   *string                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId            *string                                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag                        []*DescribeDBInstancesShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceCategoriesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceCategoriesShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceDescription(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceIds(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceModesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceModesShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceStatusesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceStatusesShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetInstanceDeployTypesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.InstanceDeployTypesShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetInstanceNetworkType(v string) *DescribeDBInstancesShrinkRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetOwnerId(v int64) *DescribeDBInstancesShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetPageNumber(v int32) *DescribeDBInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetPageSize(v int32) *DescribeDBInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetRegionId(v string) *DescribeDBInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetResourceGroupId(v string) *DescribeDBInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetTag(v []*DescribeDBInstancesShrinkRequestTag) *DescribeDBInstancesShrinkRequest {
	s.Tag = v
	return s
}

type DescribeDBInstancesShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesShrinkRequestTag) SetKey(v string) *DescribeDBInstancesShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequestTag) SetValue(v string) *DescribeDBInstancesShrinkRequestTag {
	s.Value = &v
	return s
}

type DescribeDBInstancesResponseBody struct {
	Items            *DescribeDBInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber       *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                                `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBody) SetItems(v *DescribeDBInstancesResponseBodyItems) *DescribeDBInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageNumber(v int32) *DescribeDBInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageRecordCount(v int32) *DescribeDBInstancesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetRequestId(v string) *DescribeDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetTotalRecordCount(v int32) *DescribeDBInstancesResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeDBInstancesResponseBodyItems struct {
	DBInstance []*DescribeDBInstancesResponseBodyItemsDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItems) SetDBInstance(v []*DescribeDBInstancesResponseBodyItemsDBInstance) *DescribeDBInstancesResponseBodyItems {
	s.DBInstance = v
	return s
}

type DescribeDBInstancesResponseBodyItemsDBInstance struct {
	ConnectionMode        *string                                             `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	CreateTime            *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBInstanceCategory    *string                                             `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
	DBInstanceDescription *string                                             `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId          *string                                             `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceMode        *string                                             `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	DBInstanceNetType     *string                                             `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	DBInstanceStatus      *string                                             `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	Engine                *string                                             `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string                                             `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime            *string                                             `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceDeployType    *string                                             `json:"InstanceDeployType,omitempty" xml:"InstanceDeployType,omitempty"`
	InstanceNetworkType   *string                                             `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	LockMode              *string                                             `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason            *string                                             `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MasterNodeNum         *int32                                              `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	PayType               *string                                             `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId              *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string                                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SegNodeNum            *string                                             `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	ServerlessMode        *string                                             `json:"ServerlessMode,omitempty" xml:"ServerlessMode,omitempty"`
	StorageSize           *string                                             `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	StorageType           *string                                             `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tags                  *DescribeDBInstancesResponseBodyItemsDBInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// vSwitch ID。
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	VpcId  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItemsDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetConnectionMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetCreateTime(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceCategory(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceCategory = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceDescription(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceNetType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceStatus(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetEngine(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetEngineVersion(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetExpireTime(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetInstanceDeployType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.InstanceDeployType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetInstanceNetworkType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetLockMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetLockReason(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetMasterNodeNum(v int32) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.MasterNodeNum = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetPayType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetRegionId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetResourceGroupId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetSegNodeNum(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.SegNodeNum = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetServerlessMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ServerlessMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetStorageSize(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetStorageType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetTags(v *DescribeDBInstancesResponseBodyItemsDBInstanceTags) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.Tags = v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetVSwitchId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetVpcId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetZoneId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ZoneId = &v
	return s
}

type DescribeDBInstancesResponseBodyItemsDBInstanceTags struct {
	Tag []*DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceTags) SetTag(v []*DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) *DescribeDBInstancesResponseBodyItemsDBInstanceTags {
	s.Tag = v
	return s
}

type DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) SetKey(v string) *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) SetValue(v string) *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag {
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

type DescribeDataBackupsRequest struct {
	BackupId     *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupMode   *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DataType     *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsRequest) SetBackupId(v string) *DescribeDataBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetBackupMode(v string) *DescribeDataBackupsRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetBackupStatus(v string) *DescribeDataBackupsRequest {
	s.BackupStatus = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetDBInstanceId(v string) *DescribeDataBackupsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetDataType(v string) *DescribeDataBackupsRequest {
	s.DataType = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetEndTime(v string) *DescribeDataBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetPageNumber(v int32) *DescribeDataBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetPageSize(v int32) *DescribeDataBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetStartTime(v string) *DescribeDataBackupsRequest {
	s.StartTime = &v
	return s
}

type DescribeDataBackupsResponseBody struct {
	Items      []*DescribeDataBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsResponseBody) SetItems(v []*DescribeDataBackupsResponseBodyItems) *DescribeDataBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetPageNumber(v int32) *DescribeDataBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetPageSize(v int32) *DescribeDataBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetRequestId(v string) *DescribeDataBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetTotalCount(v int32) *DescribeDataBackupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDataBackupsResponseBodyItems struct {
	BackupEndTime        *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupEndTimeLocal   *string `json:"BackupEndTimeLocal,omitempty" xml:"BackupEndTimeLocal,omitempty"`
	BackupMode           *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupSetId          *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	BackupSize           *int64  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	BackupStartTime      *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupStartTimeLocal *string `json:"BackupStartTimeLocal,omitempty" xml:"BackupStartTimeLocal,omitempty"`
	BackupStatus         *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	BaksetName           *string `json:"BaksetName,omitempty" xml:"BaksetName,omitempty"`
	ConsistentTime       *int64  `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DataType             *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
}

func (s DescribeDataBackupsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupEndTime(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupEndTimeLocal(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupEndTimeLocal = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupMode(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupMode = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupSetId(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupSetId = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupSize(v int64) *DescribeDataBackupsResponseBodyItems {
	s.BackupSize = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupStartTime(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupStartTimeLocal(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupStartTimeLocal = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupStatus(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupStatus = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBaksetName(v string) *DescribeDataBackupsResponseBodyItems {
	s.BaksetName = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetConsistentTime(v int64) *DescribeDataBackupsResponseBodyItems {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetDBInstanceId(v string) *DescribeDataBackupsResponseBodyItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetDataType(v string) *DescribeDataBackupsResponseBodyItems {
	s.DataType = &v
	return s
}

type DescribeDataBackupsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDataBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsResponse) SetHeaders(v map[string]*string) *DescribeDataBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataBackupsResponse) SetStatusCode(v int32) *DescribeDataBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataBackupsResponse) SetBody(v *DescribeDataBackupsResponseBody) *DescribeDataBackupsResponse {
	s.Body = v
	return s
}

type DescribeDataShareInstancesRequest struct {
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SearchValue     *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
}

func (s DescribeDataShareInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataShareInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesRequest) SetOwnerId(v int64) *DescribeDataShareInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetPageNumber(v int32) *DescribeDataShareInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetPageSize(v int32) *DescribeDataShareInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetRegionId(v string) *DescribeDataShareInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetResourceGroupId(v string) *DescribeDataShareInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetSearchValue(v string) *DescribeDataShareInstancesRequest {
	s.SearchValue = &v
	return s
}

type DescribeDataShareInstancesResponseBody struct {
	Items            *DescribeDataShareInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber       *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                       `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                                       `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDataShareInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataShareInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesResponseBody) SetItems(v *DescribeDataShareInstancesResponseBodyItems) *DescribeDataShareInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataShareInstancesResponseBody) SetPageNumber(v int32) *DescribeDataShareInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBody) SetPageRecordCount(v int32) *DescribeDataShareInstancesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBody) SetRequestId(v string) *DescribeDataShareInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBody) SetTotalRecordCount(v int32) *DescribeDataShareInstancesResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeDataShareInstancesResponseBodyItems struct {
	DBInstance []*DescribeDataShareInstancesResponseBodyItemsDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeDataShareInstancesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataShareInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesResponseBodyItems) SetDBInstance(v []*DescribeDataShareInstancesResponseBodyItemsDBInstance) *DescribeDataShareInstancesResponseBodyItems {
	s.DBInstance = v
	return s
}

type DescribeDataShareInstancesResponseBodyItemsDBInstance struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceMode  *string `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	DataShareStatus *string `json:"DataShareStatus,omitempty" xml:"DataShareStatus,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId          *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDataShareInstancesResponseBodyItemsDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataShareInstancesResponseBodyItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetDBInstanceId(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetDBInstanceMode(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.DBInstanceMode = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetDataShareStatus(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.DataShareStatus = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetDescription(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.Description = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetRegionId(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetZoneId(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.ZoneId = &v
	return s
}

type DescribeDataShareInstancesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDataShareInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataShareInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataShareInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesResponse) SetHeaders(v map[string]*string) *DescribeDataShareInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataShareInstancesResponse) SetStatusCode(v int32) *DescribeDataShareInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataShareInstancesResponse) SetBody(v *DescribeDataShareInstancesResponseBody) *DescribeDataShareInstancesResponse {
	s.Body = v
	return s
}

type DescribeDataSharePerformanceRequest struct {
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Key             *string `json:"Key,omitempty" xml:"Key,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataSharePerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSharePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceRequest) SetEndTime(v string) *DescribeDataSharePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataSharePerformanceRequest) SetKey(v string) *DescribeDataSharePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDataSharePerformanceRequest) SetRegionId(v string) *DescribeDataSharePerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDataSharePerformanceRequest) SetResourceGroupId(v string) *DescribeDataSharePerformanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDataSharePerformanceRequest) SetStartTime(v string) *DescribeDataSharePerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeDataSharePerformanceResponseBody struct {
	DBClusterId     *string                                                    `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime         *string                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PerformanceKeys []*DescribeDataSharePerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Repeated"`
	RequestId       *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataSharePerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSharePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponseBody) SetDBClusterId(v string) *DescribeDataSharePerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBody) SetEndTime(v string) *DescribeDataSharePerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBody) SetPerformanceKeys(v []*DescribeDataSharePerformanceResponseBodyPerformanceKeys) *DescribeDataSharePerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDataSharePerformanceResponseBody) SetRequestId(v string) *DescribeDataSharePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBody) SetStartTime(v string) *DescribeDataSharePerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDataSharePerformanceResponseBodyPerformanceKeys struct {
	Name   *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	Series []*DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	Unit   *string                                                          `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeys) SetName(v string) *DescribeDataSharePerformanceResponseBodyPerformanceKeys {
	s.Name = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeys) SetSeries(v []*DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) *DescribeDataSharePerformanceResponseBodyPerformanceKeys {
	s.Series = v
	return s
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeys) SetUnit(v string) *DescribeDataSharePerformanceResponseBodyPerformanceKeys {
	s.Unit = &v
	return s
}

type DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries struct {
	Name   *string                                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Values []*DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) SetName(v string) *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries {
	s.Name = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) SetValues(v []*DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues) *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries {
	s.Values = v
	return s
}

type DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues struct {
	Point []*string `json:"Point,omitempty" xml:"Point,omitempty" type:"Repeated"`
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues) SetPoint(v []*string) *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues {
	s.Point = v
	return s
}

type DescribeDataSharePerformanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDataSharePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataSharePerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSharePerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponse) SetHeaders(v map[string]*string) *DescribeDataSharePerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataSharePerformanceResponse) SetStatusCode(v int32) *DescribeDataSharePerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataSharePerformanceResponse) SetBody(v *DescribeDataSharePerformanceResponseBody) *DescribeDataSharePerformanceResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisDimensionsRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDiagnosisDimensionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsRequest) SetDBInstanceId(v string) *DescribeDiagnosisDimensionsRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDiagnosisDimensionsResponseBody struct {
	Databases []*string `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserNames []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosisDimensionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetDatabases(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetRequestId(v string) *DescribeDiagnosisDimensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetUserNames(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.UserNames = v
	return s
}

type DescribeDiagnosisDimensionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiagnosisDimensionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisDimensionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisDimensionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponse) SetStatusCode(v int32) *DescribeDiagnosisDimensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisDimensionsResponse) SetBody(v *DescribeDiagnosisDimensionsResponseBody) *DescribeDiagnosisDimensionsResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisMonitorPerformanceRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database       *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	User           *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetDBInstanceId(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetDatabase(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetEndTime(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetQueryCondition(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetStartTime(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetUser(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.User = &v
	return s
}

type DescribeDiagnosisMonitorPerformanceResponseBody struct {
	Performances          []*DescribeDiagnosisMonitorPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	PerformancesThreshold *int32                                                         `json:"PerformancesThreshold,omitempty" xml:"PerformancesThreshold,omitempty"`
	PerformancesTruncated *bool                                                          `json:"PerformancesTruncated,omitempty" xml:"PerformancesTruncated,omitempty"`
	RequestId             *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetPerformances(v []*DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetPerformancesThreshold(v int32) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.PerformancesThreshold = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetPerformancesTruncated(v bool) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.PerformancesTruncated = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetRequestId(v string) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDiagnosisMonitorPerformanceResponseBodyPerformances struct {
	Cost      *int32  `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Database  *string `json:"Database,omitempty" xml:"Database,omitempty"`
	QueryID   *string `json:"QueryID,omitempty" xml:"QueryID,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	User      *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetCost(v int32) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.Cost = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetDatabase(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetQueryID(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.QueryID = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetStartTime(v int64) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetStatus(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetUser(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.User = &v
	return s
}

type DescribeDiagnosisMonitorPerformanceResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiagnosisMonitorPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisMonitorPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisMonitorPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) SetStatusCode(v int32) *DescribeDiagnosisMonitorPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) SetBody(v *DescribeDiagnosisMonitorPerformanceResponseBody) *DescribeDiagnosisMonitorPerformanceResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisRecordsRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database       *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Keyword        *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Order          *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	User           *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDiagnosisRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsRequest) SetDBInstanceId(v string) *DescribeDiagnosisRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetDatabase(v string) *DescribeDiagnosisRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetEndTime(v string) *DescribeDiagnosisRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetKeyword(v string) *DescribeDiagnosisRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetOrder(v string) *DescribeDiagnosisRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetPageNumber(v int32) *DescribeDiagnosisRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetPageSize(v int32) *DescribeDiagnosisRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetQueryCondition(v string) *DescribeDiagnosisRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetStartTime(v string) *DescribeDiagnosisRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetUser(v string) *DescribeDiagnosisRecordsRequest {
	s.User = &v
	return s
}

type DescribeDiagnosisRecordsResponseBody struct {
	Items      []*DescribeDiagnosisRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBody) SetItems(v []*DescribeDiagnosisRecordsResponseBodyItems) *DescribeDiagnosisRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetPageNumber(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetRequestId(v string) *DescribeDiagnosisRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetTotalCount(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDiagnosisRecordsResponseBodyItems struct {
	Database              *string `json:"Database,omitempty" xml:"Database,omitempty"`
	Duration              *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	QueryID               *string `json:"QueryID,omitempty" xml:"QueryID,omitempty"`
	SQLStmt               *string `json:"SQLStmt,omitempty" xml:"SQLStmt,omitempty"`
	SQLTruncated          *bool   `json:"SQLTruncated,omitempty" xml:"SQLTruncated,omitempty"`
	SQLTruncatedThreshold *int32  `json:"SQLTruncatedThreshold,omitempty" xml:"SQLTruncatedThreshold,omitempty"`
	SessionID             *string `json:"SessionID,omitempty" xml:"SessionID,omitempty"`
	StartTime             *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	User                  *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetDatabase(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetDuration(v int32) *DescribeDiagnosisRecordsResponseBodyItems {
	s.Duration = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetQueryID(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.QueryID = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetSQLStmt(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.SQLStmt = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetSQLTruncated(v bool) *DescribeDiagnosisRecordsResponseBodyItems {
	s.SQLTruncated = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetSQLTruncatedThreshold(v int32) *DescribeDiagnosisRecordsResponseBodyItems {
	s.SQLTruncatedThreshold = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetSessionID(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.SessionID = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetStartTime(v int64) *DescribeDiagnosisRecordsResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetStatus(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetUser(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.User = &v
	return s
}

type DescribeDiagnosisRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiagnosisRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisRecordsResponse) SetStatusCode(v int32) *DescribeDiagnosisRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponse) SetBody(v *DescribeDiagnosisRecordsResponseBody) *DescribeDiagnosisRecordsResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisSQLInfoRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database     *string `json:"Database,omitempty" xml:"Database,omitempty"`
	QueryID      *string `json:"QueryID,omitempty" xml:"QueryID,omitempty"`
}

func (s DescribeDiagnosisSQLInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoRequest) SetDBInstanceId(v string) *DescribeDiagnosisSQLInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetDatabase(v string) *DescribeDiagnosisSQLInfoRequest {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetQueryID(v string) *DescribeDiagnosisSQLInfoRequest {
	s.QueryID = &v
	return s
}

type DescribeDiagnosisSQLInfoResponseBody struct {
	Database      *string `json:"Database,omitempty" xml:"Database,omitempty"`
	Duration      *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	MaxOutputRows *string `json:"MaxOutputRows,omitempty" xml:"MaxOutputRows,omitempty"`
	QueryID       *string `json:"QueryID,omitempty" xml:"QueryID,omitempty"`
	QueryPlan     *string `json:"QueryPlan,omitempty" xml:"QueryPlan,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SQLStmt       *string `json:"SQLStmt,omitempty" xml:"SQLStmt,omitempty"`
	SessionID     *string `json:"SessionID,omitempty" xml:"SessionID,omitempty"`
	SortedMetrics *string `json:"SortedMetrics,omitempty" xml:"SortedMetrics,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TextPlan      *string `json:"TextPlan,omitempty" xml:"TextPlan,omitempty"`
	User          *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDiagnosisSQLInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetDatabase(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetDuration(v int32) *DescribeDiagnosisSQLInfoResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetMaxOutputRows(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.MaxOutputRows = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetQueryID(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.QueryID = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetQueryPlan(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.QueryPlan = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetRequestId(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetSQLStmt(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.SQLStmt = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetSessionID(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.SessionID = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetSortedMetrics(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.SortedMetrics = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetStartTime(v int64) *DescribeDiagnosisSQLInfoResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetStatus(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetTextPlan(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.TextPlan = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetUser(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.User = &v
	return s
}

type DescribeDiagnosisSQLInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiagnosisSQLInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisSQLInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisSQLInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponse) SetStatusCode(v int32) *DescribeDiagnosisSQLInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponse) SetBody(v *DescribeDiagnosisSQLInfoResponseBody) *DescribeDiagnosisSQLInfoResponse {
	s.Body = v
	return s
}

type DescribeDownloadRecordsRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDownloadRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsRequest) SetDBInstanceId(v string) *DescribeDownloadRecordsRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDownloadRecordsResponseBody struct {
	Records   []*DescribeDownloadRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDownloadRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponseBody) SetRecords(v []*DescribeDownloadRecordsResponseBodyRecords) *DescribeDownloadRecordsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeDownloadRecordsResponseBody) SetRequestId(v string) *DescribeDownloadRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDownloadRecordsResponseBodyRecords struct {
	DownloadId   *int64  `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	DownloadUrl  *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ExceptionMsg *string `json:"ExceptionMsg,omitempty" xml:"ExceptionMsg,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDownloadRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetDownloadId(v int64) *DescribeDownloadRecordsResponseBodyRecords {
	s.DownloadId = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetDownloadUrl(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetExceptionMsg(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.ExceptionMsg = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetFileName(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.FileName = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetStatus(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.Status = &v
	return s
}

type DescribeDownloadRecordsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDownloadRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDownloadRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponse) SetHeaders(v map[string]*string) *DescribeDownloadRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadRecordsResponse) SetStatusCode(v int32) *DescribeDownloadRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadRecordsResponse) SetBody(v *DescribeDownloadRecordsResponseBody) *DescribeDownloadRecordsResponse {
	s.Body = v
	return s
}

type DescribeHealthStatusRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s DescribeHealthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusRequest) SetDBInstanceId(v string) *DescribeHealthStatusRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetKey(v string) *DescribeHealthStatusRequest {
	s.Key = &v
	return s
}

type DescribeHealthStatusResponseBody struct {
	DBClusterId *string                                 `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status      *DescribeHealthStatusResponseBodyStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s DescribeHealthStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBody) SetDBClusterId(v string) *DescribeHealthStatusResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeHealthStatusResponseBody) SetRequestId(v string) *DescribeHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHealthStatusResponseBody) SetStatus(v *DescribeHealthStatusResponseBodyStatus) *DescribeHealthStatusResponseBody {
	s.Status = v
	return s
}

type DescribeHealthStatusResponseBodyStatus struct {
	AdbgpSegmentDiskUsagePercentMax *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax `json:"adbgp_segment_disk_usage_percent_max,omitempty" xml:"adbgp_segment_disk_usage_percent_max,omitempty" type:"Struct"`
	AdbpgConnectionStatus           *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus           `json:"adbpg_connection_status,omitempty" xml:"adbpg_connection_status,omitempty" type:"Struct"`
	AdbpgDiskStatus                 *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus                 `json:"adbpg_disk_status,omitempty" xml:"adbpg_disk_status,omitempty" type:"Struct"`
	AdbpgDiskUsagePercent           *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent           `json:"adbpg_disk_usage_percent,omitempty" xml:"adbpg_disk_usage_percent,omitempty" type:"Struct"`
	AdbpgMasterDiskUsagePercentMax  *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax  `json:"adbpg_master_disk_usage_percent_max,omitempty" xml:"adbpg_master_disk_usage_percent_max,omitempty" type:"Struct"`
	AdbpgMasterStatus               *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus               `json:"adbpg_master_status,omitempty" xml:"adbpg_master_status,omitempty" type:"Struct"`
	AdbpgSegmentStatus              *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus              `json:"adbpg_segment_status,omitempty" xml:"adbpg_segment_status,omitempty" type:"Struct"`
	AdbpgStatus                     *DescribeHealthStatusResponseBodyStatusAdbpgStatus                     `json:"adbpg_status,omitempty" xml:"adbpg_status,omitempty" type:"Struct"`
	NodeMasterConnectionStatus      *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus      `json:"node_master_connection_status,omitempty" xml:"node_master_connection_status,omitempty" type:"Struct"`
	NodeMasterStatus                *DescribeHealthStatusResponseBodyStatusNodeMasterStatus                `json:"node_master_status,omitempty" xml:"node_master_status,omitempty" type:"Struct"`
	NodeSegmentConnectionStatus     *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus     `json:"node_segment_connection_status,omitempty" xml:"node_segment_connection_status,omitempty" type:"Struct"`
	NodeSegmentDiskStatus           *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus           `json:"node_segment_disk_status,omitempty" xml:"node_segment_disk_status,omitempty" type:"Struct"`
}

func (s DescribeHealthStatusResponseBodyStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbgpSegmentDiskUsagePercentMax(v *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) *DescribeHealthStatusResponseBodyStatus {
	s.AdbgpSegmentDiskUsagePercentMax = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgConnectionStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgConnectionStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgDiskStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgDiskStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgDiskUsagePercent(v *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgDiskUsagePercent = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgMasterDiskUsagePercentMax(v *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgMasterDiskUsagePercentMax = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgMasterStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgMasterStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgSegmentStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgSegmentStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetNodeMasterConnectionStatus(v *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) *DescribeHealthStatusResponseBodyStatus {
	s.NodeMasterConnectionStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetNodeMasterStatus(v *DescribeHealthStatusResponseBodyStatusNodeMasterStatus) *DescribeHealthStatusResponseBodyStatus {
	s.NodeMasterStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetNodeSegmentConnectionStatus(v *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) *DescribeHealthStatusResponseBodyStatus {
	s.NodeSegmentConnectionStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetNodeSegmentDiskStatus(v *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) *DescribeHealthStatusResponseBodyStatus {
	s.NodeSegmentDiskStatus = v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax struct {
	Status *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Value  *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus struct {
	Status *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Value  *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus struct {
	Status *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Value  *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent struct {
	Status *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Value  *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax struct {
	Status *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Value  *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus struct {
	Status *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Value  *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus struct {
	Status *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Value  *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgStatus struct {
	Status *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Value  *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus struct {
	Status *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Value  *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusNodeMasterStatus struct {
	Status *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Value  *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusNodeMasterStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusNodeMasterStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusNodeMasterStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusNodeMasterStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus struct {
	Status *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Value  *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus struct {
	Status *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Value  *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHealthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHealthStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponse) SetHeaders(v map[string]*string) *DescribeHealthStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeHealthStatusResponse) SetStatusCode(v int32) *DescribeHealthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHealthStatusResponse) SetBody(v *DescribeHealthStatusResponseBody) *DescribeHealthStatusResponse {
	s.Body = v
	return s
}

type DescribeLogBackupsRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLogBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupsRequest) SetDBInstanceId(v string) *DescribeLogBackupsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetEndTime(v string) *DescribeLogBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetPageNumber(v int32) *DescribeLogBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetPageSize(v int32) *DescribeLogBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetStartTime(v string) *DescribeLogBackupsRequest {
	s.StartTime = &v
	return s
}

type DescribeLogBackupsResponseBody struct {
	Items        []*DescribeLogBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber   *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalLogSize *int64                                 `json:"TotalLogSize,omitempty" xml:"TotalLogSize,omitempty"`
}

func (s DescribeLogBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupsResponseBody) SetItems(v []*DescribeLogBackupsResponseBodyItems) *DescribeLogBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetPageNumber(v int32) *DescribeLogBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetPageSize(v int32) *DescribeLogBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetRequestId(v string) *DescribeLogBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetTotalCount(v int32) *DescribeLogBackupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetTotalLogSize(v int64) *DescribeLogBackupsResponseBody {
	s.TotalLogSize = &v
	return s
}

type DescribeLogBackupsResponseBodyItems struct {
	BackupId     *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	LogFileName  *string `json:"LogFileName,omitempty" xml:"LogFileName,omitempty"`
	LogFileSize  *int64  `json:"LogFileSize,omitempty" xml:"LogFileSize,omitempty"`
	LogTime      *string `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	SegmentName  *string `json:"SegmentName,omitempty" xml:"SegmentName,omitempty"`
}

func (s DescribeLogBackupsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupsResponseBodyItems) SetBackupId(v string) *DescribeLogBackupsResponseBodyItems {
	s.BackupId = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetDBInstanceId(v string) *DescribeLogBackupsResponseBodyItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetLogFileName(v string) *DescribeLogBackupsResponseBodyItems {
	s.LogFileName = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetLogFileSize(v int64) *DescribeLogBackupsResponseBodyItems {
	s.LogFileSize = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetLogTime(v string) *DescribeLogBackupsResponseBodyItems {
	s.LogTime = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetSegmentName(v string) *DescribeLogBackupsResponseBodyItems {
	s.SegmentName = &v
	return s
}

type DescribeLogBackupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLogBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupsResponse) SetHeaders(v map[string]*string) *DescribeLogBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogBackupsResponse) SetStatusCode(v int32) *DescribeLogBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogBackupsResponse) SetBody(v *DescribeLogBackupsResponseBody) *DescribeLogBackupsResponse {
	s.Body = v
	return s
}

type DescribeModifyParameterLogRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeModifyParameterLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeModifyParameterLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogRequest) SetDBInstanceId(v string) *DescribeModifyParameterLogRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetEndTime(v string) *DescribeModifyParameterLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetStartTime(v string) *DescribeModifyParameterLogRequest {
	s.StartTime = &v
	return s
}

type DescribeModifyParameterLogResponseBody struct {
	Changelogs []*DescribeModifyParameterLogResponseBodyChangelogs `json:"Changelogs,omitempty" xml:"Changelogs,omitempty" type:"Repeated"`
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeModifyParameterLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeModifyParameterLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponseBody) SetChangelogs(v []*DescribeModifyParameterLogResponseBodyChangelogs) *DescribeModifyParameterLogResponseBody {
	s.Changelogs = v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) SetRequestId(v string) *DescribeModifyParameterLogResponseBody {
	s.RequestId = &v
	return s
}

type DescribeModifyParameterLogResponseBodyChangelogs struct {
	EffectTime           *string `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValid       *string `json:"ParameterValid,omitempty" xml:"ParameterValid,omitempty"`
	ParameterValueAfter  *string `json:"ParameterValueAfter,omitempty" xml:"ParameterValueAfter,omitempty"`
	ParameterValueBefore *string `json:"ParameterValueBefore,omitempty" xml:"ParameterValueBefore,omitempty"`
}

func (s DescribeModifyParameterLogResponseBodyChangelogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeModifyParameterLogResponseBodyChangelogs) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetEffectTime(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.EffectTime = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetParameterName(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.ParameterName = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetParameterValid(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.ParameterValid = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetParameterValueAfter(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.ParameterValueAfter = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetParameterValueBefore(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.ParameterValueBefore = &v
	return s
}

type DescribeModifyParameterLogResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeModifyParameterLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeModifyParameterLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeModifyParameterLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponse) SetHeaders(v map[string]*string) *DescribeModifyParameterLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeModifyParameterLogResponse) SetStatusCode(v int32) *DescribeModifyParameterLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModifyParameterLogResponse) SetBody(v *DescribeModifyParameterLogResponseBody) *DescribeModifyParameterLogResponse {
	s.Body = v
	return s
}

type DescribeParametersRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeParametersRequest) SetDBInstanceId(v string) *DescribeParametersRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeParametersResponseBody struct {
	Parameters []*DescribeParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CurrentValue         *string `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
	ForceRestartInstance *string `json:"ForceRestartInstance,omitempty" xml:"ForceRestartInstance,omitempty"`
	IsChangeableConfig   *string `json:"IsChangeableConfig,omitempty" xml:"IsChangeableConfig,omitempty"`
	OptionalRange        *string `json:"OptionalRange,omitempty" xml:"OptionalRange,omitempty"`
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyParameters) SetCurrentValue(v string) *DescribeParametersResponseBodyParameters {
	s.CurrentValue = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetForceRestartInstance(v string) *DescribeParametersResponseBodyParameters {
	s.ForceRestartInstance = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetIsChangeableConfig(v string) *DescribeParametersResponseBodyParameters {
	s.IsChangeableConfig = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetOptionalRange(v string) *DescribeParametersResponseBodyParameters {
	s.OptionalRange = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetParameterDescription(v string) *DescribeParametersResponseBodyParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetParameterName(v string) *DescribeParametersResponseBodyParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetParameterValue(v string) *DescribeParametersResponseBodyParameters {
	s.ParameterValue = &v
	return s
}

type DescribeParametersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeParametersResponse) SetStatusCode(v int32) *DescribeParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParametersResponse) SetBody(v *DescribeParametersResponseBody) *DescribeParametersResponse {
	s.Body = v
	return s
}

type DescribeRdsVSwitchsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRdsVSwitchsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsRequest) SetOwnerAccount(v string) *DescribeRdsVSwitchsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetOwnerId(v int64) *DescribeRdsVSwitchsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetRegionId(v string) *DescribeRdsVSwitchsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetResourceGroupId(v string) *DescribeRdsVSwitchsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetResourceOwnerAccount(v string) *DescribeRdsVSwitchsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetResourceOwnerId(v int64) *DescribeRdsVSwitchsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetSecurityToken(v string) *DescribeRdsVSwitchsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetVpcId(v string) *DescribeRdsVSwitchsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetZoneId(v string) *DescribeRdsVSwitchsRequest {
	s.ZoneId = &v
	return s
}

type DescribeRdsVSwitchsResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VSwitches *DescribeRdsVSwitchsResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Struct"`
}

func (s DescribeRdsVSwitchsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseBody) SetRequestId(v string) *DescribeRdsVSwitchsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBody) SetVSwitches(v *DescribeRdsVSwitchsResponseBodyVSwitches) *DescribeRdsVSwitchsResponseBody {
	s.VSwitches = v
	return s
}

type DescribeRdsVSwitchsResponseBodyVSwitches struct {
	VSwitch []*DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch `json:"VSwitch,omitempty" xml:"VSwitch,omitempty" type:"Repeated"`
}

func (s DescribeRdsVSwitchsResponseBodyVSwitches) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitches) SetVSwitch(v []*DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) *DescribeRdsVSwitchsResponseBodyVSwitches {
	s.VSwitch = v
	return s
}

type DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch struct {
	AliUid      *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Bid         *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	IsDefault   *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	IzNo        *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	RegionNo    *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// vSwitch ID。
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetAliUid(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.AliUid = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetBid(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.Bid = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetCidrBlock(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetGmtCreate(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetGmtModified(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.GmtModified = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetIsDefault(v bool) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetIzNo(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.IzNo = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetRegionNo(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.RegionNo = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetStatus(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.Status = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetVSwitchId(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetVSwitchName(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.VSwitchName = &v
	return s
}

type DescribeRdsVSwitchsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRdsVSwitchsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRdsVSwitchsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponse) SetHeaders(v map[string]*string) *DescribeRdsVSwitchsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsVSwitchsResponse) SetStatusCode(v int32) *DescribeRdsVSwitchsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsVSwitchsResponse) SetBody(v *DescribeRdsVSwitchsResponseBody) *DescribeRdsVSwitchsResponse {
	s.Body = v
	return s
}

type DescribeRdsVpcsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRdsVpcsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsRequest) SetOwnerAccount(v string) *DescribeRdsVpcsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetOwnerId(v int64) *DescribeRdsVpcsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetRegionId(v string) *DescribeRdsVpcsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetResourceGroupId(v string) *DescribeRdsVpcsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetResourceOwnerAccount(v string) *DescribeRdsVpcsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetResourceOwnerId(v int64) *DescribeRdsVpcsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetSecurityToken(v string) *DescribeRdsVpcsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetZoneId(v string) *DescribeRdsVpcsRequest {
	s.ZoneId = &v
	return s
}

type DescribeRdsVpcsResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Vpcs      *DescribeRdsVpcsResponseBodyVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Struct"`
}

func (s DescribeRdsVpcsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBody) SetRequestId(v string) *DescribeRdsVpcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsVpcsResponseBody) SetVpcs(v *DescribeRdsVpcsResponseBodyVpcs) *DescribeRdsVpcsResponseBody {
	s.Vpcs = v
	return s
}

type DescribeRdsVpcsResponseBodyVpcs struct {
	Vpc []*DescribeRdsVpcsResponseBodyVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s DescribeRdsVpcsResponseBodyVpcs) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponseBodyVpcs) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBodyVpcs) SetVpc(v []*DescribeRdsVpcsResponseBodyVpcsVpc) *DescribeRdsVpcsResponseBodyVpcs {
	s.Vpc = v
	return s
}

type DescribeRdsVpcsResponseBodyVpcsVpc struct {
	AliUid      *string                                       `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Bid         *string                                       `json:"Bid,omitempty" xml:"Bid,omitempty"`
	CidrBlock   *string                                       `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	GmtCreate   *string                                       `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string                                       `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	IsDefault   *bool                                         `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	RegionNo    *string                                       `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	Status      *string                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchs    []*DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" type:"Repeated"`
	// VPC ID。
	VpcId   *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeRdsVpcsResponseBodyVpcsVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponseBodyVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetAliUid(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.AliUid = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetBid(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.Bid = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetCidrBlock(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetGmtCreate(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetGmtModified(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.GmtModified = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetIsDefault(v bool) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetRegionNo(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetStatus(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.Status = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetVSwitchs(v []*DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.VSwitchs = v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetVpcId(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetVpcName(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.VpcName = &v
	return s
}

type DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs struct {
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	IsDefault   *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	IzNo        *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// vSwitch ID。
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetCidrBlock(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetGmtCreate(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetGmtModified(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.GmtModified = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetIsDefault(v bool) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetIzNo(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.IzNo = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetStatus(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.Status = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetVSwitchId(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetVSwitchName(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.VSwitchName = &v
	return s
}

type DescribeRdsVpcsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRdsVpcsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRdsVpcsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponse) SetHeaders(v map[string]*string) *DescribeRdsVpcsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsVpcsResponse) SetStatusCode(v int32) *DescribeRdsVpcsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsVpcsResponse) SetBody(v *DescribeRdsVpcsResponseBody) *DescribeRdsVpcsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegion(v string) *DescribeRegionsRequest {
	s.Region = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
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
	RegionId *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Zones    *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetZones(v *DescribeRegionsResponseBodyRegionsRegionZones) *DescribeRegionsResponseBodyRegionsRegion {
	s.Zones = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionZones struct {
	Zone []*DescribeRegionsResponseBodyRegionsRegionZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZones) SetZone(v []*DescribeRegionsResponseBodyRegionsRegionZonesZone) *DescribeRegionsResponseBodyRegionsRegionZones {
	s.Zone = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionZonesZone struct {
	VpcEnabled *bool   `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetVpcEnabled(v bool) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.VpcEnabled = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.ZoneId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeResourceUsageRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeResourceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageRequest) SetDBInstanceId(v string) *DescribeResourceUsageRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeResourceUsageResponseBody struct {
	BackupSize   *int64  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DataSize     *int64  `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	DiskUsed     *int64  `json:"DiskUsed,omitempty" xml:"DiskUsed,omitempty"`
	Engine       *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	LogSize      *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeResourceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponseBody) SetBackupSize(v int64) *DescribeResourceUsageResponseBody {
	s.BackupSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetDBInstanceId(v string) *DescribeResourceUsageResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetDataSize(v int64) *DescribeResourceUsageResponseBody {
	s.DataSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetDiskUsed(v int64) *DescribeResourceUsageResponseBody {
	s.DiskUsed = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetEngine(v string) *DescribeResourceUsageResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetLogSize(v int64) *DescribeResourceUsageResponseBody {
	s.LogSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetRequestId(v string) *DescribeResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeResourceUsageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceUsageResponse) SetStatusCode(v int32) *DescribeResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceUsageResponse) SetBody(v *DescribeResourceUsageResponseBody) *DescribeResourceUsageResponse {
	s.Body = v
	return s
}

type DescribeSQLCollectorPolicyRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeSQLCollectorPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLCollectorPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLCollectorPolicyRequest) SetDBInstanceId(v string) *DescribeSQLCollectorPolicyRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeSQLCollectorPolicyResponseBody struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SQLCollectorStatus *string `json:"SQLCollectorStatus,omitempty" xml:"SQLCollectorStatus,omitempty"`
}

func (s DescribeSQLCollectorPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLCollectorPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLCollectorPolicyResponseBody) SetRequestId(v string) *DescribeSQLCollectorPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLCollectorPolicyResponseBody) SetSQLCollectorStatus(v string) *DescribeSQLCollectorPolicyResponseBody {
	s.SQLCollectorStatus = &v
	return s
}

type DescribeSQLCollectorPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLCollectorPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLCollectorPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLCollectorPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLCollectorPolicyResponse) SetHeaders(v map[string]*string) *DescribeSQLCollectorPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLCollectorPolicyResponse) SetStatusCode(v int32) *DescribeSQLCollectorPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLCollectorPolicyResponse) SetBody(v *DescribeSQLCollectorPolicyResponseBody) *DescribeSQLCollectorPolicyResponse {
	s.Body = v
	return s
}

type DescribeSQLLogByQueryIdRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	QueryId      *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s DescribeSQLLogByQueryIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogByQueryIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogByQueryIdRequest) SetDBInstanceId(v string) *DescribeSQLLogByQueryIdRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogByQueryIdRequest) SetQueryId(v string) *DescribeSQLLogByQueryIdRequest {
	s.QueryId = &v
	return s
}

type DescribeSQLLogByQueryIdResponseBody struct {
	Items     []*DescribeSQLLogByQueryIdResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSQLLogByQueryIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogByQueryIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogByQueryIdResponseBody) SetItems(v []*DescribeSQLLogByQueryIdResponseBodyItems) *DescribeSQLLogByQueryIdResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBody) SetRequestId(v string) *DescribeSQLLogByQueryIdResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSQLLogByQueryIdResponseBodyItems struct {
	AccountName          *string   `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBName               *string   `json:"DBName,omitempty" xml:"DBName,omitempty"`
	DBRole               *string   `json:"DBRole,omitempty" xml:"DBRole,omitempty"`
	ExecuteCost          *float32  `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	ExecuteState         *string   `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	OperationClass       *string   `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	OperationExecuteTime *string   `json:"OperationExecuteTime,omitempty" xml:"OperationExecuteTime,omitempty"`
	OperationType        *string   `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	QueryId              *string   `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	ReturnRowCounts      *int64    `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	SQLPlan              *string   `json:"SQLPlan,omitempty" xml:"SQLPlan,omitempty"`
	SQLText              *string   `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	ScanRowCounts        *int64    `json:"ScanRowCounts,omitempty" xml:"ScanRowCounts,omitempty"`
	SliceIds             []*string `json:"SliceIds,omitempty" xml:"SliceIds,omitempty" type:"Repeated"`
	SourceIP             *string   `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	SourcePort           *int32    `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
}

func (s DescribeSQLLogByQueryIdResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogByQueryIdResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetAccountName(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.AccountName = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetDBName(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetDBRole(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.DBRole = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetExecuteCost(v float32) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetExecuteState(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetOperationClass(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetOperationExecuteTime(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.OperationExecuteTime = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetOperationType(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetQueryId(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.QueryId = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetReturnRowCounts(v int64) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetSQLPlan(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.SQLPlan = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetSQLText(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetScanRowCounts(v int64) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.ScanRowCounts = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetSliceIds(v []*string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.SliceIds = v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetSourceIP(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetSourcePort(v int32) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.SourcePort = &v
	return s
}

type DescribeSQLLogByQueryIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLLogByQueryIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLLogByQueryIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogByQueryIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogByQueryIdResponse) SetHeaders(v map[string]*string) *DescribeSQLLogByQueryIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogByQueryIdResponse) SetStatusCode(v int32) *DescribeSQLLogByQueryIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponse) SetBody(v *DescribeSQLLogByQueryIdResponseBody) *DescribeSQLLogByQueryIdResponse {
	s.Body = v
	return s
}

type DescribeSQLLogCountRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database       *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecuteCost    *string `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	ExecuteState   *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	MaxExecuteCost *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
	MinExecuteCost *string `json:"MinExecuteCost,omitempty" xml:"MinExecuteCost,omitempty"`
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	OperationType  *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	QueryKeywords  *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	SourceIP       *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	User           *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLLogCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountRequest) SetDBInstanceId(v string) *DescribeSQLLogCountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetDatabase(v string) *DescribeSQLLogCountRequest {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetEndTime(v string) *DescribeSQLLogCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetExecuteCost(v string) *DescribeSQLLogCountRequest {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetExecuteState(v string) *DescribeSQLLogCountRequest {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetMaxExecuteCost(v string) *DescribeSQLLogCountRequest {
	s.MaxExecuteCost = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetMinExecuteCost(v string) *DescribeSQLLogCountRequest {
	s.MinExecuteCost = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetOperationClass(v string) *DescribeSQLLogCountRequest {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetOperationType(v string) *DescribeSQLLogCountRequest {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetQueryKeywords(v string) *DescribeSQLLogCountRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetSourceIP(v string) *DescribeSQLLogCountRequest {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetStartTime(v string) *DescribeSQLLogCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetUser(v string) *DescribeSQLLogCountRequest {
	s.User = &v
	return s
}

type DescribeSQLLogCountResponseBody struct {
	DBClusterId *string                                 `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime     *string                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Items       []*DescribeSQLLogCountResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime   *string                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSQLLogCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseBody) SetDBClusterId(v string) *DescribeSQLLogCountResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLLogCountResponseBody) SetEndTime(v string) *DescribeSQLLogCountResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogCountResponseBody) SetItems(v []*DescribeSQLLogCountResponseBodyItems) *DescribeSQLLogCountResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLLogCountResponseBody) SetRequestId(v string) *DescribeSQLLogCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogCountResponseBody) SetStartTime(v string) *DescribeSQLLogCountResponseBody {
	s.StartTime = &v
	return s
}

type DescribeSQLLogCountResponseBodyItems struct {
	Name   *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Series []*DescribeSQLLogCountResponseBodyItemsSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogCountResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseBodyItems) SetName(v string) *DescribeSQLLogCountResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeSQLLogCountResponseBodyItems) SetSeries(v []*DescribeSQLLogCountResponseBodyItemsSeries) *DescribeSQLLogCountResponseBodyItems {
	s.Series = v
	return s
}

type DescribeSQLLogCountResponseBodyItemsSeries struct {
	Values []*DescribeSQLLogCountResponseBodyItemsSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogCountResponseBodyItemsSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponseBodyItemsSeries) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseBodyItemsSeries) SetValues(v []*DescribeSQLLogCountResponseBodyItemsSeriesValues) *DescribeSQLLogCountResponseBodyItemsSeries {
	s.Values = v
	return s
}

type DescribeSQLLogCountResponseBodyItemsSeriesValues struct {
	Point []*string `json:"Point,omitempty" xml:"Point,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogCountResponseBodyItemsSeriesValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponseBodyItemsSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseBodyItemsSeriesValues) SetPoint(v []*string) *DescribeSQLLogCountResponseBodyItemsSeriesValues {
	s.Point = v
	return s
}

type DescribeSQLLogCountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLLogCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLLogCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponse) SetHeaders(v map[string]*string) *DescribeSQLLogCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogCountResponse) SetStatusCode(v int32) *DescribeSQLLogCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLLogCountResponse) SetBody(v *DescribeSQLLogCountResponseBody) *DescribeSQLLogCountResponse {
	s.Body = v
	return s
}

type DescribeSQLLogFilesRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSQLLogFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesRequest) SetDBInstanceId(v string) *DescribeSQLLogFilesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetFileName(v string) *DescribeSQLLogFilesRequest {
	s.FileName = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetPageNumber(v int32) *DescribeSQLLogFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetPageSize(v int32) *DescribeSQLLogFilesRequest {
	s.PageSize = &v
	return s
}

type DescribeSQLLogFilesResponseBody struct {
	Items            *DescribeSQLLogFilesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber       *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                                `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSQLLogFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponseBody) SetItems(v *DescribeSQLLogFilesResponseBodyItems) *DescribeSQLLogFilesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLLogFilesResponseBody) SetPageNumber(v int32) *DescribeSQLLogFilesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogFilesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBody) SetRequestId(v string) *DescribeSQLLogFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBody) SetTotalRecordCount(v int32) *DescribeSQLLogFilesResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeSQLLogFilesResponseBodyItems struct {
	LogFile []*DescribeSQLLogFilesResponseBodyItemsLogFile `json:"LogFile,omitempty" xml:"LogFile,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogFilesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogFilesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponseBodyItems) SetLogFile(v []*DescribeSQLLogFilesResponseBodyItemsLogFile) *DescribeSQLLogFilesResponseBodyItems {
	s.LogFile = v
	return s
}

type DescribeSQLLogFilesResponseBodyItemsLogFile struct {
	FileID         *string `json:"FileID,omitempty" xml:"FileID,omitempty"`
	LogDownloadURL *string `json:"LogDownloadURL,omitempty" xml:"LogDownloadURL,omitempty"`
	LogEndTime     *string `json:"LogEndTime,omitempty" xml:"LogEndTime,omitempty"`
	LogSize        *string `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	LogStartTime   *string `json:"LogStartTime,omitempty" xml:"LogStartTime,omitempty"`
	LogStatus      *string `json:"LogStatus,omitempty" xml:"LogStatus,omitempty"`
}

func (s DescribeSQLLogFilesResponseBodyItemsLogFile) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogFilesResponseBodyItemsLogFile) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetFileID(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.FileID = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogDownloadURL(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogDownloadURL = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogEndTime(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogEndTime = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogSize(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogSize = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogStartTime(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogStartTime = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogStatus(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogStatus = &v
	return s
}

type DescribeSQLLogFilesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLLogFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLLogFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponse) SetHeaders(v map[string]*string) *DescribeSQLLogFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogFilesResponse) SetStatusCode(v int32) *DescribeSQLLogFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLLogFilesResponse) SetBody(v *DescribeSQLLogFilesResponseBody) *DescribeSQLLogFilesResponse {
	s.Body = v
	return s
}

type DescribeSQLLogRecordsRequest struct {
	DBInstanceId  *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database      *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Form          *string `json:"Form,omitempty" xml:"Form,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryKeywords *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	User          *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsRequest) SetDBInstanceId(v string) *DescribeSQLLogRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetDatabase(v string) *DescribeSQLLogRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetEndTime(v string) *DescribeSQLLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetForm(v string) *DescribeSQLLogRecordsRequest {
	s.Form = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetPageNumber(v int32) *DescribeSQLLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetPageSize(v int32) *DescribeSQLLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetQueryKeywords(v string) *DescribeSQLLogRecordsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetStartTime(v string) *DescribeSQLLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetUser(v string) *DescribeSQLLogRecordsRequest {
	s.User = &v
	return s
}

type DescribeSQLLogRecordsResponseBody struct {
	Items            *DescribeSQLLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber       *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                  `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                                  `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSQLLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsResponseBody) SetItems(v *DescribeSQLLogRecordsResponseBodyItems) *DescribeSQLLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLLogRecordsResponseBody) SetPageNumber(v int32) *DescribeSQLLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBody) SetRequestId(v string) *DescribeSQLLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeSQLLogRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeSQLLogRecordsResponseBodyItems struct {
	SQLRecord []*DescribeSQLLogRecordsResponseBodyItemsSQLRecord `json:"SQLRecord,omitempty" xml:"SQLRecord,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsResponseBodyItems) SetSQLRecord(v []*DescribeSQLLogRecordsResponseBodyItemsSQLRecord) *DescribeSQLLogRecordsResponseBodyItems {
	s.SQLRecord = v
	return s
}

type DescribeSQLLogRecordsResponseBodyItemsSQLRecord struct {
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBName              *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ExecuteTime         *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	HostAddress         *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	ReturnRowCounts     *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	SQLText             *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	ThreadID            *string `json:"ThreadID,omitempty" xml:"ThreadID,omitempty"`
	TotalExecutionTimes *int64  `json:"TotalExecutionTimes,omitempty" xml:"TotalExecutionTimes,omitempty"`
}

func (s DescribeSQLLogRecordsResponseBodyItemsSQLRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogRecordsResponseBodyItemsSQLRecord) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetAccountName(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.AccountName = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetDBName(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetExecuteTime(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetHostAddress(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetReturnRowCounts(v int64) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetSQLText(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetThreadID(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.ThreadID = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetTotalExecutionTimes(v int64) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.TotalExecutionTimes = &v
	return s
}

type DescribeSQLLogRecordsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeSQLLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogRecordsResponse) SetStatusCode(v int32) *DescribeSQLLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLLogRecordsResponse) SetBody(v *DescribeSQLLogRecordsResponseBody) *DescribeSQLLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeSQLLogsRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database       *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecuteCost    *string `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	ExecuteState   *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	MaxExecuteCost *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
	MinExecuteCost *string `json:"MinExecuteCost,omitempty" xml:"MinExecuteCost,omitempty"`
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	OperationType  *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryKeywords  *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	SourceIP       *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	User           *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsRequest) SetDBInstanceId(v string) *DescribeSQLLogsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetDatabase(v string) *DescribeSQLLogsRequest {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetEndTime(v string) *DescribeSQLLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetExecuteCost(v string) *DescribeSQLLogsRequest {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetExecuteState(v string) *DescribeSQLLogsRequest {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetMaxExecuteCost(v string) *DescribeSQLLogsRequest {
	s.MaxExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetMinExecuteCost(v string) *DescribeSQLLogsRequest {
	s.MinExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetOperationClass(v string) *DescribeSQLLogsRequest {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetOperationType(v string) *DescribeSQLLogsRequest {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetPageNumber(v int32) *DescribeSQLLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetPageSize(v int32) *DescribeSQLLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetQueryKeywords(v string) *DescribeSQLLogsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetSourceIP(v string) *DescribeSQLLogsRequest {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetStartTime(v string) *DescribeSQLLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetUser(v string) *DescribeSQLLogsRequest {
	s.User = &v
	return s
}

type DescribeSQLLogsResponseBody struct {
	Items           []*DescribeSQLLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber      *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount *int32                              `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId       *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSQLLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsResponseBody) SetItems(v []*DescribeSQLLogsResponseBodyItems) *DescribeSQLLogsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLLogsResponseBody) SetPageNumber(v int32) *DescribeSQLLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogsResponseBody) SetRequestId(v string) *DescribeSQLLogsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSQLLogsResponseBodyItems struct {
	AccountName          *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBName               *string  `json:"DBName,omitempty" xml:"DBName,omitempty"`
	DBRole               *string  `json:"DBRole,omitempty" xml:"DBRole,omitempty"`
	ExecuteCost          *float32 `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	ExecuteState         *string  `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	OperationClass       *string  `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	OperationExecuteTime *string  `json:"OperationExecuteTime,omitempty" xml:"OperationExecuteTime,omitempty"`
	OperationType        *string  `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	ReturnRowCounts      *int64   `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	SQLPlan              *string  `json:"SQLPlan,omitempty" xml:"SQLPlan,omitempty"`
	SQLText              *string  `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	ScanRowCounts        *int64   `json:"ScanRowCounts,omitempty" xml:"ScanRowCounts,omitempty"`
	SourceIP             *string  `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	SourcePort           *int32   `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
}

func (s DescribeSQLLogsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsResponseBodyItems) SetAccountName(v string) *DescribeSQLLogsResponseBodyItems {
	s.AccountName = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetDBName(v string) *DescribeSQLLogsResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetDBRole(v string) *DescribeSQLLogsResponseBodyItems {
	s.DBRole = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetExecuteCost(v float32) *DescribeSQLLogsResponseBodyItems {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetExecuteState(v string) *DescribeSQLLogsResponseBodyItems {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetOperationClass(v string) *DescribeSQLLogsResponseBodyItems {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetOperationExecuteTime(v string) *DescribeSQLLogsResponseBodyItems {
	s.OperationExecuteTime = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetOperationType(v string) *DescribeSQLLogsResponseBodyItems {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetReturnRowCounts(v int64) *DescribeSQLLogsResponseBodyItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetSQLPlan(v string) *DescribeSQLLogsResponseBodyItems {
	s.SQLPlan = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetSQLText(v string) *DescribeSQLLogsResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetScanRowCounts(v int64) *DescribeSQLLogsResponseBodyItems {
	s.ScanRowCounts = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetSourceIP(v string) *DescribeSQLLogsResponseBodyItems {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetSourcePort(v int32) *DescribeSQLLogsResponseBodyItems {
	s.SourcePort = &v
	return s
}

type DescribeSQLLogsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsResponse) SetHeaders(v map[string]*string) *DescribeSQLLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogsResponse) SetStatusCode(v int32) *DescribeSQLLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLLogsResponse) SetBody(v *DescribeSQLLogsResponseBody) *DescribeSQLLogsResponse {
	s.Body = v
	return s
}

type DescribeSQLLogsOnSliceRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ExecuteState   *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	MaxExecuteCost *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
	MinExecuteCost *string `json:"MinExecuteCost,omitempty" xml:"MinExecuteCost,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryId        *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	SliceId        *string `json:"SliceId,omitempty" xml:"SliceId,omitempty"`
}

func (s DescribeSQLLogsOnSliceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsOnSliceRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsOnSliceRequest) SetDBInstanceId(v string) *DescribeSQLLogsOnSliceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogsOnSliceRequest) SetExecuteState(v string) *DescribeSQLLogsOnSliceRequest {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsOnSliceRequest) SetMaxExecuteCost(v string) *DescribeSQLLogsOnSliceRequest {
	s.MaxExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsOnSliceRequest) SetMinExecuteCost(v string) *DescribeSQLLogsOnSliceRequest {
	s.MinExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsOnSliceRequest) SetPageNumber(v int32) *DescribeSQLLogsOnSliceRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsOnSliceRequest) SetPageSize(v int32) *DescribeSQLLogsOnSliceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogsOnSliceRequest) SetQueryId(v string) *DescribeSQLLogsOnSliceRequest {
	s.QueryId = &v
	return s
}

func (s *DescribeSQLLogsOnSliceRequest) SetSliceId(v string) *DescribeSQLLogsOnSliceRequest {
	s.SliceId = &v
	return s
}

type DescribeSQLLogsOnSliceResponseBody struct {
	PageNumber      *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount *int32                                             `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SliceLogItems   []*DescribeSQLLogsOnSliceResponseBodySliceLogItems `json:"SliceLogItems,omitempty" xml:"SliceLogItems,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogsOnSliceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsOnSliceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsOnSliceResponseBody) SetPageNumber(v int32) *DescribeSQLLogsOnSliceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogsOnSliceResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBody) SetRequestId(v string) *DescribeSQLLogsOnSliceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBody) SetSliceLogItems(v []*DescribeSQLLogsOnSliceResponseBodySliceLogItems) *DescribeSQLLogsOnSliceResponseBody {
	s.SliceLogItems = v
	return s
}

type DescribeSQLLogsOnSliceResponseBodySliceLogItems struct {
	ExecuteCost             *float32 `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	ExecuteStatus           *string  `json:"ExecuteStatus,omitempty" xml:"ExecuteStatus,omitempty"`
	OperationExecuteEndTime *string  `json:"OperationExecuteEndTime,omitempty" xml:"OperationExecuteEndTime,omitempty"`
	OperationExecuteTime    *string  `json:"OperationExecuteTime,omitempty" xml:"OperationExecuteTime,omitempty"`
	PeakMemory              *float32 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	ReturnRowCounts         *int64   `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	// Segment ID。
	SegmentId   *string `json:"SegmentId,omitempty" xml:"SegmentId,omitempty"`
	SegmentName *string `json:"SegmentName,omitempty" xml:"SegmentName,omitempty"`
}

func (s DescribeSQLLogsOnSliceResponseBodySliceLogItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsOnSliceResponseBodySliceLogItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetExecuteCost(v float32) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetExecuteStatus(v string) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.ExecuteStatus = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetOperationExecuteEndTime(v string) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.OperationExecuteEndTime = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetOperationExecuteTime(v string) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.OperationExecuteTime = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetPeakMemory(v float32) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.PeakMemory = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetReturnRowCounts(v int64) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetSegmentId(v string) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.SegmentId = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetSegmentName(v string) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.SegmentName = &v
	return s
}

type DescribeSQLLogsOnSliceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLLogsOnSliceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLLogsOnSliceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsOnSliceResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsOnSliceResponse) SetHeaders(v map[string]*string) *DescribeSQLLogsOnSliceResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogsOnSliceResponse) SetStatusCode(v int32) *DescribeSQLLogsOnSliceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponse) SetBody(v *DescribeSQLLogsOnSliceResponseBody) *DescribeSQLLogsOnSliceResponse {
	s.Body = v
	return s
}

type DescribeSQLLogsV2Request struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database        *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecuteCost     *string `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	ExecuteState    *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	MaxExecuteCost  *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
	MinExecuteCost  *string `json:"MinExecuteCost,omitempty" xml:"MinExecuteCost,omitempty"`
	OperationClass  *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	OperationType   *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	PageNumber      *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryKeywords   *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SourceIP        *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	User            *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLLogsV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsV2Request) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsV2Request) SetDBInstanceId(v string) *DescribeSQLLogsV2Request {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetDatabase(v string) *DescribeSQLLogsV2Request {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetEndTime(v string) *DescribeSQLLogsV2Request {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetExecuteCost(v string) *DescribeSQLLogsV2Request {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetExecuteState(v string) *DescribeSQLLogsV2Request {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetMaxExecuteCost(v string) *DescribeSQLLogsV2Request {
	s.MaxExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetMinExecuteCost(v string) *DescribeSQLLogsV2Request {
	s.MinExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetOperationClass(v string) *DescribeSQLLogsV2Request {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetOperationType(v string) *DescribeSQLLogsV2Request {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetPageNumber(v string) *DescribeSQLLogsV2Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetPageSize(v string) *DescribeSQLLogsV2Request {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetQueryKeywords(v string) *DescribeSQLLogsV2Request {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetRegionId(v string) *DescribeSQLLogsV2Request {
	s.RegionId = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetResourceGroupId(v string) *DescribeSQLLogsV2Request {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetSourceIP(v string) *DescribeSQLLogsV2Request {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetStartTime(v string) *DescribeSQLLogsV2Request {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetUser(v string) *DescribeSQLLogsV2Request {
	s.User = &v
	return s
}

type DescribeSQLLogsV2ResponseBody struct {
	Items           []*DescribeSQLLogsV2ResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber      *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount *int32                                `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId       *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSQLLogsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsV2ResponseBody) SetItems(v []*DescribeSQLLogsV2ResponseBodyItems) *DescribeSQLLogsV2ResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLLogsV2ResponseBody) SetPageNumber(v int32) *DescribeSQLLogsV2ResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogsV2ResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBody) SetRequestId(v string) *DescribeSQLLogsV2ResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSQLLogsV2ResponseBodyItems struct {
	AccountName          *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBName               *string  `json:"DBName,omitempty" xml:"DBName,omitempty"`
	DBRole               *string  `json:"DBRole,omitempty" xml:"DBRole,omitempty"`
	ExecuteCost          *float32 `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	ExecuteState         *string  `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	OperationClass       *string  `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	OperationExecuteTime *string  `json:"OperationExecuteTime,omitempty" xml:"OperationExecuteTime,omitempty"`
	OperationType        *string  `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	ReturnRowCounts      *int64   `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	SQLText              *string  `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	ScanRowCounts        *int64   `json:"ScanRowCounts,omitempty" xml:"ScanRowCounts,omitempty"`
	SourceIP             *string  `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	SourcePort           *int32   `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
}

func (s DescribeSQLLogsV2ResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsV2ResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetAccountName(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.AccountName = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetDBName(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetDBRole(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.DBRole = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetExecuteCost(v float32) *DescribeSQLLogsV2ResponseBodyItems {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetExecuteState(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetOperationClass(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetOperationExecuteTime(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.OperationExecuteTime = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetOperationType(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetReturnRowCounts(v int64) *DescribeSQLLogsV2ResponseBodyItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetSQLText(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetScanRowCounts(v int64) *DescribeSQLLogsV2ResponseBodyItems {
	s.ScanRowCounts = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetSourceIP(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetSourcePort(v int32) *DescribeSQLLogsV2ResponseBodyItems {
	s.SourcePort = &v
	return s
}

type DescribeSQLLogsV2Response struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLLogsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLLogsV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsV2Response) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsV2Response) SetHeaders(v map[string]*string) *DescribeSQLLogsV2Response {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogsV2Response) SetStatusCode(v int32) *DescribeSQLLogsV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLLogsV2Response) SetBody(v *DescribeSQLLogsV2ResponseBody) *DescribeSQLLogsV2Response {
	s.Body = v
	return s
}

type DescribeSampleDataRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeSampleDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSampleDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataRequest) SetDBInstanceId(v string) *DescribeSampleDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSampleDataRequest) SetOwnerId(v int64) *DescribeSampleDataRequest {
	s.OwnerId = &v
	return s
}

type DescribeSampleDataResponseBody struct {
	DBInstanceId  *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ErrorMessage  *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HasSampleData *bool   `json:"HasSampleData,omitempty" xml:"HasSampleData,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSampleDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSampleDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataResponseBody) SetDBInstanceId(v string) *DescribeSampleDataResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSampleDataResponseBody) SetErrorMessage(v string) *DescribeSampleDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSampleDataResponseBody) SetHasSampleData(v bool) *DescribeSampleDataResponseBody {
	s.HasSampleData = &v
	return s
}

func (s *DescribeSampleDataResponseBody) SetRequestId(v string) *DescribeSampleDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSampleDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSampleDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSampleDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSampleDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataResponse) SetHeaders(v map[string]*string) *DescribeSampleDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeSampleDataResponse) SetStatusCode(v int32) *DescribeSampleDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSampleDataResponse) SetBody(v *DescribeSampleDataResponseBody) *DescribeSampleDataResponse {
	s.Body = v
	return s
}

type DescribeSlowLogRecordsRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBName       *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// SQL ID。
	SQLId     *int64  `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequest) SetDBInstanceId(v string) *DescribeSlowLogRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBName(v string) *DescribeSlowLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetEndTime(v string) *DescribeSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageNumber(v int32) *DescribeSlowLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageSize(v int32) *DescribeSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetSQLId(v int64) *DescribeSlowLogRecordsRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetStartTime(v string) *DescribeSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

type DescribeSlowLogRecordsResponseBody struct {
	Engine           *string                                  `json:"Engine,omitempty" xml:"Engine,omitempty"`
	Items            *DescribeSlowLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber       *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                   `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                                   `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) SetEngine(v string) *DescribeSlowLogRecordsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetItems(v *DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageNumber(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageRecordCount(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeSlowLogRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeSlowLogRecordsResponseBodyItems struct {
	SQLSlowRecord []*DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord `json:"SQLSlowRecord,omitempty" xml:"SQLSlowRecord,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSQLSlowRecord(v []*DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) *DescribeSlowLogRecordsResponseBodyItems {
	s.SQLSlowRecord = v
	return s
}

type DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord struct {
	DBName             *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	HostAddress        *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	LockTimes          *int64  `json:"LockTimes,omitempty" xml:"LockTimes,omitempty"`
	ParseRowCounts     *int64  `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	QueryTimes         *int64  `json:"QueryTimes,omitempty" xml:"QueryTimes,omitempty"`
	ReturnRowCounts    *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	SQLText            *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetExecutionStartTime(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ExecutionStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetHostAddress(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetLockTimes(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.LockTimes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetParseRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetQueryTimes(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.QueryTimes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetReturnRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetSQLText(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.SQLText = &v
	return s
}

type DescribeSlowLogRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSlowLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeSlowLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetStatusCode(v int32) *DescribeSlowLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetBody(v *DescribeSlowLogRecordsResponseBody) *DescribeSlowLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeSlowSQLLogsRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database       *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecuteState   *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	MaxExecuteCost *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
	MinExecuteCost *string `json:"MinExecuteCost,omitempty" xml:"MinExecuteCost,omitempty"`
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	OperationType  *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryKeywords  *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	SourceIP       *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	User           *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSlowSQLLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLLogsRequest) SetDBInstanceId(v string) *DescribeSlowSQLLogsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetDatabase(v string) *DescribeSlowSQLLogsRequest {
	s.Database = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetEndTime(v string) *DescribeSlowSQLLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetExecuteState(v string) *DescribeSlowSQLLogsRequest {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetMaxExecuteCost(v string) *DescribeSlowSQLLogsRequest {
	s.MaxExecuteCost = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetMinExecuteCost(v string) *DescribeSlowSQLLogsRequest {
	s.MinExecuteCost = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetOperationClass(v string) *DescribeSlowSQLLogsRequest {
	s.OperationClass = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetOperationType(v string) *DescribeSlowSQLLogsRequest {
	s.OperationType = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetPageNumber(v int32) *DescribeSlowSQLLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetPageSize(v int32) *DescribeSlowSQLLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetQueryKeywords(v string) *DescribeSlowSQLLogsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetSourceIP(v string) *DescribeSlowSQLLogsRequest {
	s.SourceIP = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetStartTime(v string) *DescribeSlowSQLLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetUser(v string) *DescribeSlowSQLLogsRequest {
	s.User = &v
	return s
}

type DescribeSlowSQLLogsResponseBody struct {
	Items           []*DescribeSlowSQLLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber      *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount *int32                                  `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId       *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSlowSQLLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLLogsResponseBody) SetItems(v []*DescribeSlowSQLLogsResponseBodyItems) *DescribeSlowSQLLogsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowSQLLogsResponseBody) SetPageNumber(v int32) *DescribeSlowSQLLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBody) SetPageRecordCount(v int32) *DescribeSlowSQLLogsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBody) SetRequestId(v string) *DescribeSlowSQLLogsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSlowSQLLogsResponseBodyItems struct {
	AccountName          *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBName               *string  `json:"DBName,omitempty" xml:"DBName,omitempty"`
	DBRole               *string  `json:"DBRole,omitempty" xml:"DBRole,omitempty"`
	ExecuteCost          *float32 `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	ExecuteState         *string  `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	OperationClass       *string  `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	OperationExecuteTime *string  `json:"OperationExecuteTime,omitempty" xml:"OperationExecuteTime,omitempty"`
	OperationType        *string  `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	QueryId              *string  `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	ReturnRowCounts      *int64   `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	SQLPlan              *string  `json:"SQLPlan,omitempty" xml:"SQLPlan,omitempty"`
	SQLText              *string  `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	ScanRowCounts        *int64   `json:"ScanRowCounts,omitempty" xml:"ScanRowCounts,omitempty"`
	SourceIP             *string  `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	SourcePort           *int32   `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
}

func (s DescribeSlowSQLLogsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetAccountName(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.AccountName = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetDBName(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetDBRole(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.DBRole = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetExecuteCost(v float32) *DescribeSlowSQLLogsResponseBodyItems {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetExecuteState(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetOperationClass(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.OperationClass = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetOperationExecuteTime(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.OperationExecuteTime = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetOperationType(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.OperationType = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetQueryId(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.QueryId = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetReturnRowCounts(v int64) *DescribeSlowSQLLogsResponseBodyItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetSQLPlan(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.SQLPlan = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetSQLText(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetScanRowCounts(v int64) *DescribeSlowSQLLogsResponseBodyItems {
	s.ScanRowCounts = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetSourceIP(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.SourceIP = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetSourcePort(v int32) *DescribeSlowSQLLogsResponseBodyItems {
	s.SourcePort = &v
	return s
}

type DescribeSlowSQLLogsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSlowSQLLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowSQLLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLLogsResponse) SetHeaders(v map[string]*string) *DescribeSlowSQLLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowSQLLogsResponse) SetStatusCode(v int32) *DescribeSlowSQLLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowSQLLogsResponse) SetBody(v *DescribeSlowSQLLogsResponseBody) *DescribeSlowSQLLogsResponse {
	s.Body = v
	return s
}

type DescribeSpecificationRequest struct {
	CpuCores     *int32  `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StorageType  *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	TotalNodeNum *int32  `json:"TotalNodeNum,omitempty" xml:"TotalNodeNum,omitempty"`
}

func (s DescribeSpecificationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationRequest) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationRequest) SetCpuCores(v int32) *DescribeSpecificationRequest {
	s.CpuCores = &v
	return s
}

func (s *DescribeSpecificationRequest) SetDBInstanceId(v string) *DescribeSpecificationRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSpecificationRequest) SetOwnerId(v int64) *DescribeSpecificationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSpecificationRequest) SetStorageType(v string) *DescribeSpecificationRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeSpecificationRequest) SetTotalNodeNum(v int32) *DescribeSpecificationRequest {
	s.TotalNodeNum = &v
	return s
}

type DescribeSpecificationResponseBody struct {
	DBInstanceClass      []*DescribeSpecificationResponseBodyDBInstanceClass      `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty" type:"Repeated"`
	DBInstanceGroupCount []*DescribeSpecificationResponseBodyDBInstanceGroupCount `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty" type:"Repeated"`
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageNotice        []*DescribeSpecificationResponseBodyStorageNotice        `json:"StorageNotice,omitempty" xml:"StorageNotice,omitempty" type:"Repeated"`
}

func (s DescribeSpecificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationResponseBody) SetDBInstanceClass(v []*DescribeSpecificationResponseBodyDBInstanceClass) *DescribeSpecificationResponseBody {
	s.DBInstanceClass = v
	return s
}

func (s *DescribeSpecificationResponseBody) SetDBInstanceGroupCount(v []*DescribeSpecificationResponseBodyDBInstanceGroupCount) *DescribeSpecificationResponseBody {
	s.DBInstanceGroupCount = v
	return s
}

func (s *DescribeSpecificationResponseBody) SetRequestId(v string) *DescribeSpecificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSpecificationResponseBody) SetStorageNotice(v []*DescribeSpecificationResponseBodyStorageNotice) *DescribeSpecificationResponseBody {
	s.StorageNotice = v
	return s
}

type DescribeSpecificationResponseBodyDBInstanceClass struct {
	Text  *string `json:"Text,omitempty" xml:"Text,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSpecificationResponseBodyDBInstanceClass) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationResponseBodyDBInstanceClass) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationResponseBodyDBInstanceClass) SetText(v string) *DescribeSpecificationResponseBodyDBInstanceClass {
	s.Text = &v
	return s
}

func (s *DescribeSpecificationResponseBodyDBInstanceClass) SetValue(v string) *DescribeSpecificationResponseBodyDBInstanceClass {
	s.Value = &v
	return s
}

type DescribeSpecificationResponseBodyDBInstanceGroupCount struct {
	Text  *string `json:"Text,omitempty" xml:"Text,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSpecificationResponseBodyDBInstanceGroupCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationResponseBodyDBInstanceGroupCount) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationResponseBodyDBInstanceGroupCount) SetText(v string) *DescribeSpecificationResponseBodyDBInstanceGroupCount {
	s.Text = &v
	return s
}

func (s *DescribeSpecificationResponseBodyDBInstanceGroupCount) SetValue(v string) *DescribeSpecificationResponseBodyDBInstanceGroupCount {
	s.Value = &v
	return s
}

type DescribeSpecificationResponseBodyStorageNotice struct {
	Text  *string `json:"Text,omitempty" xml:"Text,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSpecificationResponseBodyStorageNotice) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationResponseBodyStorageNotice) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationResponseBodyStorageNotice) SetText(v string) *DescribeSpecificationResponseBodyStorageNotice {
	s.Text = &v
	return s
}

func (s *DescribeSpecificationResponseBodyStorageNotice) SetValue(v string) *DescribeSpecificationResponseBodyStorageNotice {
	s.Value = &v
	return s
}

type DescribeSpecificationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSpecificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSpecificationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationResponse) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationResponse) SetHeaders(v map[string]*string) *DescribeSpecificationResponse {
	s.Headers = v
	return s
}

func (s *DescribeSpecificationResponse) SetStatusCode(v int32) *DescribeSpecificationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSpecificationResponse) SetBody(v *DescribeSpecificationResponseBody) *DescribeSpecificationResponse {
	s.Body = v
	return s
}

type DescribeSupportFeaturesRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeSupportFeaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportFeaturesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupportFeaturesRequest) SetDBInstanceId(v string) *DescribeSupportFeaturesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSupportFeaturesRequest) SetOwnerId(v int64) *DescribeSupportFeaturesRequest {
	s.OwnerId = &v
	return s
}

type DescribeSupportFeaturesResponseBody struct {
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupportFeatureList *string `json:"SupportFeatureList,omitempty" xml:"SupportFeatureList,omitempty"`
}

func (s DescribeSupportFeaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportFeaturesResponseBody) SetDBInstanceId(v string) *DescribeSupportFeaturesResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSupportFeaturesResponseBody) SetRequestId(v string) *DescribeSupportFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupportFeaturesResponseBody) SetSupportFeatureList(v string) *DescribeSupportFeaturesResponseBody {
	s.SupportFeatureList = &v
	return s
}

type DescribeSupportFeaturesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSupportFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSupportFeaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportFeaturesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupportFeaturesResponse) SetHeaders(v map[string]*string) *DescribeSupportFeaturesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupportFeaturesResponse) SetStatusCode(v int32) *DescribeSupportFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupportFeaturesResponse) SetBody(v *DescribeSupportFeaturesResponseBody) *DescribeSupportFeaturesResponse {
	s.Body = v
	return s
}

type DescribeTagsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) SetOwnerAccount(v string) *DescribeTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetOwnerId(v int64) *DescribeTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetRegionId(v string) *DescribeTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceGroupId(v string) *DescribeTagsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerAccount(v string) *DescribeTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerId(v int64) *DescribeTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceType(v string) *DescribeTagsRequest {
	s.ResourceType = &v
	return s
}

type DescribeTagsResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*DescribeTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTags(v []*DescribeTagsResponseBodyTags) *DescribeTagsResponseBody {
	s.Tags = v
	return s
}

type DescribeTagsResponseBodyTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTags) SetTagKey(v string) *DescribeTagsResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *DescribeTagsResponseBodyTags) SetTagValue(v string) *DescribeTagsResponseBodyTags {
	s.TagValue = &v
	return s
}

type DescribeTagsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeTagsResponse) SetStatusCode(v int32) *DescribeTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagsResponse) SetBody(v *DescribeTagsResponseBody) *DescribeTagsResponse {
	s.Body = v
	return s
}

type DescribeUserEncryptionKeyListRequest struct {
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUserEncryptionKeyListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListRequest) SetPageNumber(v string) *DescribeUserEncryptionKeyListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetPageSize(v string) *DescribeUserEncryptionKeyListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetRegionId(v string) *DescribeUserEncryptionKeyListRequest {
	s.RegionId = &v
	return s
}

type DescribeUserEncryptionKeyListResponseBody struct {
	KmsKeys   []*DescribeUserEncryptionKeyListResponseBodyKmsKeys `json:"KmsKeys,omitempty" xml:"KmsKeys,omitempty" type:"Repeated"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserEncryptionKeyListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetKmsKeys(v []*DescribeUserEncryptionKeyListResponseBodyKmsKeys) *DescribeUserEncryptionKeyListResponseBody {
	s.KmsKeys = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserEncryptionKeyListResponseBodyKmsKeys struct {
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s DescribeUserEncryptionKeyListResponseBodyKmsKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBodyKmsKeys) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBodyKmsKeys) SetKeyId(v string) *DescribeUserEncryptionKeyListResponseBodyKmsKeys {
	s.KeyId = &v
	return s
}

type DescribeUserEncryptionKeyListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeUserEncryptionKeyListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserEncryptionKeyListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponse) SetHeaders(v map[string]*string) *DescribeUserEncryptionKeyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponse) SetStatusCode(v int32) *DescribeUserEncryptionKeyListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponse) SetBody(v *DescribeUserEncryptionKeyListResponseBody) *DescribeUserEncryptionKeyListResponse {
	s.Body = v
	return s
}

type DescribeWaitingSQLInfoRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database     *string `json:"Database,omitempty" xml:"Database,omitempty"`
	PID          *string `json:"PID,omitempty" xml:"PID,omitempty"`
}

func (s DescribeWaitingSQLInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLInfoRequest) SetDBInstanceId(v string) *DescribeWaitingSQLInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeWaitingSQLInfoRequest) SetDatabase(v string) *DescribeWaitingSQLInfoRequest {
	s.Database = &v
	return s
}

func (s *DescribeWaitingSQLInfoRequest) SetPID(v string) *DescribeWaitingSQLInfoRequest {
	s.PID = &v
	return s
}

type DescribeWaitingSQLInfoResponseBody struct {
	Database  *string                                    `json:"Database,omitempty" xml:"Database,omitempty"`
	Items     []*DescribeWaitingSQLInfoResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWaitingSQLInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLInfoResponseBody) SetDatabase(v string) *DescribeWaitingSQLInfoResponseBody {
	s.Database = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBody) SetItems(v []*DescribeWaitingSQLInfoResponseBodyItems) *DescribeWaitingSQLInfoResponseBody {
	s.Items = v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBody) SetRequestId(v string) *DescribeWaitingSQLInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeWaitingSQLInfoResponseBodyItems struct {
	Application          *string `json:"Application,omitempty" xml:"Application,omitempty"`
	BlockedByApplication *string `json:"BlockedByApplication,omitempty" xml:"BlockedByApplication,omitempty"`
	BlockedByPID         *string `json:"BlockedByPID,omitempty" xml:"BlockedByPID,omitempty"`
	BlockedBySQLStmt     *string `json:"BlockedBySQLStmt,omitempty" xml:"BlockedBySQLStmt,omitempty"`
	BlockedByUser        *string `json:"BlockedByUser,omitempty" xml:"BlockedByUser,omitempty"`
	GrantLocks           *string `json:"GrantLocks,omitempty" xml:"GrantLocks,omitempty"`
	NotGrantLocks        *string `json:"NotGrantLocks,omitempty" xml:"NotGrantLocks,omitempty"`
	PID                  *string `json:"PID,omitempty" xml:"PID,omitempty"`
	SQLStmt              *string `json:"SQLStmt,omitempty" xml:"SQLStmt,omitempty"`
	User                 *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeWaitingSQLInfoResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetApplication(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.Application = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetBlockedByApplication(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.BlockedByApplication = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetBlockedByPID(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.BlockedByPID = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetBlockedBySQLStmt(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.BlockedBySQLStmt = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetBlockedByUser(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.BlockedByUser = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetGrantLocks(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.GrantLocks = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetNotGrantLocks(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.NotGrantLocks = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetPID(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.PID = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetSQLStmt(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.SQLStmt = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetUser(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.User = &v
	return s
}

type DescribeWaitingSQLInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeWaitingSQLInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWaitingSQLInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLInfoResponse) SetHeaders(v map[string]*string) *DescribeWaitingSQLInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeWaitingSQLInfoResponse) SetStatusCode(v int32) *DescribeWaitingSQLInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponse) SetBody(v *DescribeWaitingSQLInfoResponseBody) *DescribeWaitingSQLInfoResponse {
	s.Body = v
	return s
}

type DescribeWaitingSQLRecordsRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database       *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Keyword        *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Order          *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	User           *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeWaitingSQLRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLRecordsRequest) SetDBInstanceId(v string) *DescribeWaitingSQLRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetDatabase(v string) *DescribeWaitingSQLRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetEndTime(v string) *DescribeWaitingSQLRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetKeyword(v string) *DescribeWaitingSQLRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetOrder(v string) *DescribeWaitingSQLRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetPageNumber(v int32) *DescribeWaitingSQLRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetPageSize(v int32) *DescribeWaitingSQLRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetQueryCondition(v string) *DescribeWaitingSQLRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetStartTime(v string) *DescribeWaitingSQLRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetUser(v string) *DescribeWaitingSQLRecordsRequest {
	s.User = &v
	return s
}

type DescribeWaitingSQLRecordsResponseBody struct {
	Items      []*DescribeWaitingSQLRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWaitingSQLRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLRecordsResponseBody) SetItems(v []*DescribeWaitingSQLRecordsResponseBodyItems) *DescribeWaitingSQLRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBody) SetPageNumber(v int32) *DescribeWaitingSQLRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBody) SetRequestId(v string) *DescribeWaitingSQLRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBody) SetTotalCount(v int32) *DescribeWaitingSQLRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeWaitingSQLRecordsResponseBodyItems struct {
	Database    *string `json:"Database,omitempty" xml:"Database,omitempty"`
	PID         *string `json:"PID,omitempty" xml:"PID,omitempty"`
	SQLStmt     *string `json:"SQLStmt,omitempty" xml:"SQLStmt,omitempty"`
	SessionID   *string `json:"SessionID,omitempty" xml:"SessionID,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	User        *string `json:"User,omitempty" xml:"User,omitempty"`
	WaitingTime *int64  `json:"WaitingTime,omitempty" xml:"WaitingTime,omitempty"`
}

func (s DescribeWaitingSQLRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetDatabase(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.Database = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetPID(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.PID = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetSQLStmt(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.SQLStmt = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetSessionID(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.SessionID = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetStartTime(v int64) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetStatus(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetUser(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.User = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetWaitingTime(v int64) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.WaitingTime = &v
	return s
}

type DescribeWaitingSQLRecordsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeWaitingSQLRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWaitingSQLRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLRecordsResponse) SetHeaders(v map[string]*string) *DescribeWaitingSQLRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWaitingSQLRecordsResponse) SetStatusCode(v int32) *DescribeWaitingSQLRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponse) SetBody(v *DescribeWaitingSQLRecordsResponseBody) *DescribeWaitingSQLRecordsResponse {
	s.Body = v
	return s
}

type DownloadDiagnosisRecordsRequest struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database        *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	QueryCondition  *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	User            *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DownloadDiagnosisRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsRequest) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsRequest) SetDBInstanceId(v string) *DownloadDiagnosisRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetDatabase(v string) *DownloadDiagnosisRecordsRequest {
	s.Database = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetEndTime(v string) *DownloadDiagnosisRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetLang(v string) *DownloadDiagnosisRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetQueryCondition(v string) *DownloadDiagnosisRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetResourceGroupId(v string) *DownloadDiagnosisRecordsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetStartTime(v string) *DownloadDiagnosisRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetUser(v string) *DownloadDiagnosisRecordsRequest {
	s.User = &v
	return s
}

type DownloadDiagnosisRecordsResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DownloadId   *string `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadDiagnosisRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsResponseBody) SetDBInstanceId(v string) *DownloadDiagnosisRecordsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponseBody) SetDownloadId(v string) *DownloadDiagnosisRecordsResponseBody {
	s.DownloadId = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponseBody) SetRequestId(v string) *DownloadDiagnosisRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DownloadDiagnosisRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DownloadDiagnosisRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadDiagnosisRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsResponse) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsResponse) SetHeaders(v map[string]*string) *DownloadDiagnosisRecordsResponse {
	s.Headers = v
	return s
}

func (s *DownloadDiagnosisRecordsResponse) SetStatusCode(v int32) *DownloadDiagnosisRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponse) SetBody(v *DownloadDiagnosisRecordsResponseBody) *DownloadDiagnosisRecordsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string                       `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
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
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
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

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyAccountDescriptionRequest struct {
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDBInstanceId(v string) *ModifyAccountDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

type ModifyAccountDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponseBody) SetRequestId(v string) *ModifyAccountDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccountDescriptionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAccountDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccountDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponse) SetHeaders(v map[string]*string) *ModifyAccountDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetStatusCode(v int32) *ModifyAccountDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetBody(v *ModifyAccountDescriptionResponseBody) *ModifyAccountDescriptionResponse {
	s.Body = v
	return s
}

type ModifyBackupPolicyRequest struct {
	BackupRetentionPeriod *int32  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EnableRecoveryPoint   *bool   `json:"EnableRecoveryPoint,omitempty" xml:"EnableRecoveryPoint,omitempty"`
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupTime   *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	RecoveryPointPeriod   *string `json:"RecoveryPointPeriod,omitempty" xml:"RecoveryPointPeriod,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPeriod(v int32) *ModifyBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDBInstanceId(v string) *ModifyBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableRecoveryPoint(v bool) *ModifyBackupPolicyRequest {
	s.EnableRecoveryPoint = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetRecoveryPointPeriod(v string) *ModifyBackupPolicyRequest {
	s.RecoveryPointPeriod = &v
	return s
}

type ModifyBackupPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponseBody) SetRequestId(v string) *ModifyBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBackupPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPolicyResponse) SetStatusCode(v int32) *ModifyBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupPolicyResponse) SetBody(v *ModifyBackupPolicyResponseBody) *ModifyBackupPolicyResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceConnectionModeRequest struct {
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ModifyDBInstanceConnectionModeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionModeRequest) SetConnectionMode(v string) *ModifyDBInstanceConnectionModeRequest {
	s.ConnectionMode = &v
	return s
}

func (s *ModifyDBInstanceConnectionModeRequest) SetDBInstanceId(v string) *ModifyDBInstanceConnectionModeRequest {
	s.DBInstanceId = &v
	return s
}

type ModifyDBInstanceConnectionModeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceConnectionModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionModeResponseBody) SetRequestId(v string) *ModifyDBInstanceConnectionModeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceConnectionModeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBInstanceConnectionModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceConnectionModeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionModeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceConnectionModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceConnectionModeResponse) SetStatusCode(v int32) *ModifyDBInstanceConnectionModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceConnectionModeResponse) SetBody(v *ModifyDBInstanceConnectionModeResponseBody) *ModifyDBInstanceConnectionModeResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceConnectionStringRequest struct {
	ConnectionStringPrefix  *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	DBInstanceId            *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Port                    *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyDBInstanceConnectionStringRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringRequest) SetConnectionStringPrefix(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetPort(v string) *ModifyDBInstanceConnectionStringRequest {
	s.Port = &v
	return s
}

type ModifyDBInstanceConnectionStringResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceConnectionStringResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetRequestId(v string) *ModifyDBInstanceConnectionStringResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceConnectionStringResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBInstanceConnectionStringResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceConnectionStringResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceConnectionStringResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponse) SetStatusCode(v int32) *ModifyDBInstanceConnectionStringResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponse) SetBody(v *ModifyDBInstanceConnectionStringResponseBody) *ModifyDBInstanceConnectionStringResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceDescriptionRequest struct {
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyDBInstanceDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceDescription(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceId(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetResourceGroupId(v string) *ModifyDBInstanceDescriptionRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyDBInstanceDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionResponseBody) SetRequestId(v string) *ModifyDBInstanceDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceDescriptionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceDescriptionResponse) SetStatusCode(v int32) *ModifyDBInstanceDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceDescriptionResponse) SetBody(v *ModifyDBInstanceDescriptionResponseBody) *ModifyDBInstanceDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceMaintainTimeRequest struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyDBInstanceMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetDBInstanceId(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetEndTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetResourceGroupId(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetStartTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.StartTime = &v
	return s
}

type ModifyDBInstanceMaintainTimeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceMaintainTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeResponseBody) SetRequestId(v string) *ModifyDBInstanceMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceMaintainTimeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBInstanceMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceMaintainTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceMaintainTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceMaintainTimeResponse) SetStatusCode(v int32) *ModifyDBInstanceMaintainTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeResponse) SetBody(v *ModifyDBInstanceMaintainTimeResponseBody) *ModifyDBInstanceMaintainTimeResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceNetworkTypeRequest struct {
	DBInstanceId        *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	PrivateIpAddress    *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// VPC ID。
	VPCId     *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ModifyDBInstanceNetworkTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetDBInstanceId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetInstanceNetworkType(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetPrivateIpAddress(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetVPCId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.VPCId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetVSwitchId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.VSwitchId = &v
	return s
}

type ModifyDBInstanceNetworkTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceNetworkTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeResponseBody) SetRequestId(v string) *ModifyDBInstanceNetworkTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceNetworkTypeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBInstanceNetworkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceNetworkTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceNetworkTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceNetworkTypeResponse) SetStatusCode(v int32) *ModifyDBInstanceNetworkTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeResponse) SetBody(v *ModifyDBInstanceNetworkTypeResponseBody) *ModifyDBInstanceNetworkTypeResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceResourceGroupRequest struct {
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NewResourceGroupId   *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceResourceGroupRequest) SetDBInstanceId(v string) *ModifyDBInstanceResourceGroupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetNewResourceGroupId(v string) *ModifyDBInstanceResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetOwnerAccount(v string) *ModifyDBInstanceResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetOwnerId(v int64) *ModifyDBInstanceResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetResourceGroupId(v string) *ModifyDBInstanceResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyDBInstanceResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceResourceGroupResponseBody) SetRequestId(v string) *ModifyDBInstanceResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceResourceGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBInstanceResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceResourceGroupResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceResourceGroupResponse) SetStatusCode(v int32) *ModifyDBInstanceResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupResponse) SetBody(v *ModifyDBInstanceResourceGroupResponseBody) *ModifyDBInstanceResourceGroupResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceSSLRequest struct {
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	SSLEnabled       *int32  `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
}

func (s ModifyDBInstanceSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLRequest) SetConnectionString(v string) *ModifyDBInstanceSSLRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetDBInstanceId(v string) *ModifyDBInstanceSSLRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetSSLEnabled(v int32) *ModifyDBInstanceSSLRequest {
	s.SSLEnabled = &v
	return s
}

type ModifyDBInstanceSSLResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLResponseBody) SetRequestId(v string) *ModifyDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceSSLResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceSSLResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceSSLResponse) SetStatusCode(v int32) *ModifyDBInstanceSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceSSLResponse) SetBody(v *ModifyDBInstanceSSLResponseBody) *ModifyDBInstanceSSLResponse {
	s.Body = v
	return s
}

type ModifyParametersRequest struct {
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ForceRestartInstance *bool   `json:"ForceRestartInstance,omitempty" xml:"ForceRestartInstance,omitempty"`
	Parameters           *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s ModifyParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyParametersRequest) SetDBInstanceId(v string) *ModifyParametersRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyParametersRequest) SetForceRestartInstance(v bool) *ModifyParametersRequest {
	s.ForceRestartInstance = &v
	return s
}

func (s *ModifyParametersRequest) SetParameters(v string) *ModifyParametersRequest {
	s.Parameters = &v
	return s
}

type ModifyParametersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type ModifyParametersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyParametersResponse) SetStatusCode(v int32) *ModifyParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyParametersResponse) SetBody(v *ModifyParametersResponseBody) *ModifyParametersResponse {
	s.Body = v
	return s
}

type ModifySQLCollectorPolicyRequest struct {
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	SQLCollectorStatus *string `json:"SQLCollectorStatus,omitempty" xml:"SQLCollectorStatus,omitempty"`
}

func (s ModifySQLCollectorPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySQLCollectorPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorPolicyRequest) SetDBInstanceId(v string) *ModifySQLCollectorPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySQLCollectorPolicyRequest) SetSQLCollectorStatus(v string) *ModifySQLCollectorPolicyRequest {
	s.SQLCollectorStatus = &v
	return s
}

type ModifySQLCollectorPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySQLCollectorPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySQLCollectorPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorPolicyResponseBody) SetRequestId(v string) *ModifySQLCollectorPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifySQLCollectorPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySQLCollectorPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySQLCollectorPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySQLCollectorPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorPolicyResponse) SetHeaders(v map[string]*string) *ModifySQLCollectorPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifySQLCollectorPolicyResponse) SetStatusCode(v int32) *ModifySQLCollectorPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySQLCollectorPolicyResponse) SetBody(v *ModifySQLCollectorPolicyResponseBody) *ModifySQLCollectorPolicyResponse {
	s.Body = v
	return s
}

type ModifySecurityIpsRequest struct {
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty"`
	DBInstanceIPArrayName      *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	DBInstanceId               *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ResourceGroupId            *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityIPList             *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) SetDBInstanceIPArrayAttribute(v string) *ModifySecurityIpsRequest {
	s.DBInstanceIPArrayAttribute = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceIPArrayName(v string) *ModifySecurityIpsRequest {
	s.DBInstanceIPArrayName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceId(v string) *ModifySecurityIpsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetResourceGroupId(v string) *ModifySecurityIpsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIPList(v string) *ModifySecurityIpsRequest {
	s.SecurityIPList = &v
	return s
}

type ModifySecurityIpsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type ModifySecurityIpsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifySecurityIpsResponse) SetStatusCode(v int32) *ModifySecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityIpsResponse) SetBody(v *ModifySecurityIpsResponseBody) *ModifySecurityIpsResponse {
	s.Body = v
	return s
}

type PauseInstanceRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s PauseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s PauseInstanceRequest) GoString() string {
	return s.String()
}

func (s *PauseInstanceRequest) SetDBInstanceId(v string) *PauseInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *PauseInstanceRequest) SetOwnerId(v int64) *PauseInstanceRequest {
	s.OwnerId = &v
	return s
}

type PauseInstanceResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PauseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PauseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PauseInstanceResponseBody) SetDBInstanceId(v string) *PauseInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *PauseInstanceResponseBody) SetErrorMessage(v string) *PauseInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PauseInstanceResponseBody) SetRequestId(v string) *PauseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseInstanceResponseBody) SetStatus(v bool) *PauseInstanceResponseBody {
	s.Status = &v
	return s
}

type PauseInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PauseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PauseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s PauseInstanceResponse) GoString() string {
	return s.String()
}

func (s *PauseInstanceResponse) SetHeaders(v map[string]*string) *PauseInstanceResponse {
	s.Headers = v
	return s
}

func (s *PauseInstanceResponse) SetStatusCode(v int32) *PauseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseInstanceResponse) SetBody(v *PauseInstanceResponseBody) *PauseInstanceResponse {
	s.Body = v
	return s
}

type RebalanceDBInstanceRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s RebalanceDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RebalanceDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebalanceDBInstanceRequest) SetClientToken(v string) *RebalanceDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RebalanceDBInstanceRequest) SetDBInstanceId(v string) *RebalanceDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

type RebalanceDBInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebalanceDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebalanceDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RebalanceDBInstanceResponseBody) SetRequestId(v string) *RebalanceDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RebalanceDBInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RebalanceDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebalanceDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RebalanceDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebalanceDBInstanceResponse) SetHeaders(v map[string]*string) *RebalanceDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RebalanceDBInstanceResponse) SetStatusCode(v int32) *RebalanceDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RebalanceDBInstanceResponse) SetBody(v *RebalanceDBInstanceResponseBody) *RebalanceDBInstanceResponse {
	s.Body = v
	return s
}

type ReleaseInstancePublicConnectionRequest struct {
	AddressType             *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	DBInstanceId            *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ReleaseInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionRequest) SetAddressType(v string) *ReleaseInstancePublicConnectionRequest {
	s.AddressType = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetCurrentConnectionString(v string) *ReleaseInstancePublicConnectionRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetDBInstanceId(v string) *ReleaseInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

type ReleaseInstancePublicConnectionResponseBody struct {
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

type RestartDBInstanceRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s RestartDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceRequest) SetClientToken(v string) *RestartDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartDBInstanceRequest) SetDBInstanceId(v string) *RestartDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

type RestartDBInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceResponseBody) SetRequestId(v string) *RestartDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartDBInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceResponse) SetHeaders(v map[string]*string) *RestartDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartDBInstanceResponse) SetStatusCode(v int32) *RestartDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDBInstanceResponse) SetBody(v *RestartDBInstanceResponseBody) *RestartDBInstanceResponse {
	s.Body = v
	return s
}

type ResumeInstanceRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ResumeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeInstanceRequest) GoString() string {
	return s.String()
}

func (s *ResumeInstanceRequest) SetDBInstanceId(v string) *ResumeInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ResumeInstanceRequest) SetOwnerId(v int64) *ResumeInstanceRequest {
	s.OwnerId = &v
	return s
}

type ResumeInstanceResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ResumeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeInstanceResponseBody) SetDBInstanceId(v string) *ResumeInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetErrorMessage(v string) *ResumeInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetRequestId(v string) *ResumeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetStatus(v bool) *ResumeInstanceResponseBody {
	s.Status = &v
	return s
}

type ResumeInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeInstanceResponse) GoString() string {
	return s.String()
}

func (s *ResumeInstanceResponse) SetHeaders(v map[string]*string) *ResumeInstanceResponse {
	s.Headers = v
	return s
}

func (s *ResumeInstanceResponse) SetStatusCode(v int32) *ResumeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeInstanceResponse) SetBody(v *ResumeInstanceResponseBody) *ResumeInstanceResponse {
	s.Body = v
	return s
}

type SetDBInstancePlanStatusRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlanId       *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PlanStatus   *string `json:"PlanStatus,omitempty" xml:"PlanStatus,omitempty"`
}

func (s SetDBInstancePlanStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDBInstancePlanStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDBInstancePlanStatusRequest) SetDBInstanceId(v string) *SetDBInstancePlanStatusRequest {
	s.DBInstanceId = &v
	return s
}

func (s *SetDBInstancePlanStatusRequest) SetOwnerId(v int64) *SetDBInstancePlanStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDBInstancePlanStatusRequest) SetPlanId(v string) *SetDBInstancePlanStatusRequest {
	s.PlanId = &v
	return s
}

func (s *SetDBInstancePlanStatusRequest) SetPlanStatus(v string) *SetDBInstancePlanStatusRequest {
	s.PlanStatus = &v
	return s
}

type SetDBInstancePlanStatusResponseBody struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PlanId       *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDBInstancePlanStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDBInstancePlanStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDBInstancePlanStatusResponseBody) SetErrorMessage(v string) *SetDBInstancePlanStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetDBInstancePlanStatusResponseBody) SetPlanId(v string) *SetDBInstancePlanStatusResponseBody {
	s.PlanId = &v
	return s
}

func (s *SetDBInstancePlanStatusResponseBody) SetRequestId(v string) *SetDBInstancePlanStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDBInstancePlanStatusResponseBody) SetStatus(v string) *SetDBInstancePlanStatusResponseBody {
	s.Status = &v
	return s
}

type SetDBInstancePlanStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDBInstancePlanStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDBInstancePlanStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDBInstancePlanStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDBInstancePlanStatusResponse) SetHeaders(v map[string]*string) *SetDBInstancePlanStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDBInstancePlanStatusResponse) SetStatusCode(v int32) *SetDBInstancePlanStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDBInstancePlanStatusResponse) SetBody(v *SetDBInstancePlanStatusResponseBody) *SetDBInstancePlanStatusResponse {
	s.Body = v
	return s
}

type SetDataShareInstanceRequest struct {
	InstanceList  []*string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	OperationType *string   `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerId       *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetDataShareInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDataShareInstanceRequest) GoString() string {
	return s.String()
}

func (s *SetDataShareInstanceRequest) SetInstanceList(v []*string) *SetDataShareInstanceRequest {
	s.InstanceList = v
	return s
}

func (s *SetDataShareInstanceRequest) SetOperationType(v string) *SetDataShareInstanceRequest {
	s.OperationType = &v
	return s
}

func (s *SetDataShareInstanceRequest) SetOwnerId(v int64) *SetDataShareInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDataShareInstanceRequest) SetRegionId(v string) *SetDataShareInstanceRequest {
	s.RegionId = &v
	return s
}

type SetDataShareInstanceShrinkRequest struct {
	InstanceListShrink *string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty"`
	OperationType      *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetDataShareInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDataShareInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetDataShareInstanceShrinkRequest) SetInstanceListShrink(v string) *SetDataShareInstanceShrinkRequest {
	s.InstanceListShrink = &v
	return s
}

func (s *SetDataShareInstanceShrinkRequest) SetOperationType(v string) *SetDataShareInstanceShrinkRequest {
	s.OperationType = &v
	return s
}

func (s *SetDataShareInstanceShrinkRequest) SetOwnerId(v int64) *SetDataShareInstanceShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDataShareInstanceShrinkRequest) SetRegionId(v string) *SetDataShareInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

type SetDataShareInstanceResponseBody struct {
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDataShareInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDataShareInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *SetDataShareInstanceResponseBody) SetErrMessage(v string) *SetDataShareInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SetDataShareInstanceResponseBody) SetRequestId(v string) *SetDataShareInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDataShareInstanceResponseBody) SetStatus(v string) *SetDataShareInstanceResponseBody {
	s.Status = &v
	return s
}

type SetDataShareInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDataShareInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDataShareInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDataShareInstanceResponse) GoString() string {
	return s.String()
}

func (s *SetDataShareInstanceResponse) SetHeaders(v map[string]*string) *SetDataShareInstanceResponse {
	s.Headers = v
	return s
}

func (s *SetDataShareInstanceResponse) SetStatusCode(v int32) *SetDataShareInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDataShareInstanceResponse) SetBody(v *SetDataShareInstanceResponseBody) *SetDataShareInstanceResponse {
	s.Body = v
	return s
}

type SwitchDBInstanceNetTypeRequest struct {
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Port                   *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s SwitchDBInstanceNetTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchDBInstanceNetTypeRequest) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceNetTypeRequest) SetConnectionStringPrefix(v string) *SwitchDBInstanceNetTypeRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetDBInstanceId(v string) *SwitchDBInstanceNetTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetPort(v string) *SwitchDBInstanceNetTypeRequest {
	s.Port = &v
	return s
}

type SwitchDBInstanceNetTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchDBInstanceNetTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchDBInstanceNetTypeResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceNetTypeResponseBody) SetRequestId(v string) *SwitchDBInstanceNetTypeResponseBody {
	s.RequestId = &v
	return s
}

type SwitchDBInstanceNetTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SwitchDBInstanceNetTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchDBInstanceNetTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchDBInstanceNetTypeResponse) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceNetTypeResponse) SetHeaders(v map[string]*string) *SwitchDBInstanceNetTypeResponse {
	s.Headers = v
	return s
}

func (s *SwitchDBInstanceNetTypeResponse) SetStatusCode(v int32) *SwitchDBInstanceNetTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchDBInstanceNetTypeResponse) SetBody(v *SwitchDBInstanceNetTypeResponseBody) *SwitchDBInstanceNetTypeResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerAccount         *string                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnloadSampleDataRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UnloadSampleDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UnloadSampleDataRequest) GoString() string {
	return s.String()
}

func (s *UnloadSampleDataRequest) SetDBInstanceId(v string) *UnloadSampleDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UnloadSampleDataRequest) SetOwnerId(v int64) *UnloadSampleDataRequest {
	s.OwnerId = &v
	return s
}

type UnloadSampleDataResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UnloadSampleDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnloadSampleDataResponseBody) GoString() string {
	return s.String()
}

func (s *UnloadSampleDataResponseBody) SetDBInstanceId(v string) *UnloadSampleDataResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *UnloadSampleDataResponseBody) SetErrorMessage(v string) *UnloadSampleDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnloadSampleDataResponseBody) SetRequestId(v string) *UnloadSampleDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnloadSampleDataResponseBody) SetStatus(v bool) *UnloadSampleDataResponseBody {
	s.Status = &v
	return s
}

type UnloadSampleDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnloadSampleDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnloadSampleDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UnloadSampleDataResponse) GoString() string {
	return s.String()
}

func (s *UnloadSampleDataResponse) SetHeaders(v map[string]*string) *UnloadSampleDataResponse {
	s.Headers = v
	return s
}

func (s *UnloadSampleDataResponse) SetStatusCode(v int32) *UnloadSampleDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UnloadSampleDataResponse) SetBody(v *UnloadSampleDataResponseBody) *UnloadSampleDataResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateDBInstancePlanRequest struct {
	DBInstanceId  *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlanConfig    *string `json:"PlanConfig,omitempty" xml:"PlanConfig,omitempty"`
	PlanDesc      *string `json:"PlanDesc,omitempty" xml:"PlanDesc,omitempty"`
	PlanEndDate   *string `json:"PlanEndDate,omitempty" xml:"PlanEndDate,omitempty"`
	PlanId        *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PlanName      *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	PlanStartDate *string `json:"PlanStartDate,omitempty" xml:"PlanStartDate,omitempty"`
}

func (s UpdateDBInstancePlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstancePlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateDBInstancePlanRequest) SetDBInstanceId(v string) *UpdateDBInstancePlanRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetOwnerId(v int64) *UpdateDBInstancePlanRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanConfig(v string) *UpdateDBInstancePlanRequest {
	s.PlanConfig = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanDesc(v string) *UpdateDBInstancePlanRequest {
	s.PlanDesc = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanEndDate(v string) *UpdateDBInstancePlanRequest {
	s.PlanEndDate = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanId(v string) *UpdateDBInstancePlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanName(v string) *UpdateDBInstancePlanRequest {
	s.PlanName = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanStartDate(v string) *UpdateDBInstancePlanRequest {
	s.PlanStartDate = &v
	return s
}

type UpdateDBInstancePlanResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PlanId       *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateDBInstancePlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstancePlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDBInstancePlanResponseBody) SetDBInstanceId(v string) *UpdateDBInstancePlanResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *UpdateDBInstancePlanResponseBody) SetErrorMessage(v string) *UpdateDBInstancePlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDBInstancePlanResponseBody) SetPlanId(v string) *UpdateDBInstancePlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *UpdateDBInstancePlanResponseBody) SetRequestId(v string) *UpdateDBInstancePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDBInstancePlanResponseBody) SetStatus(v string) *UpdateDBInstancePlanResponseBody {
	s.Status = &v
	return s
}

type UpdateDBInstancePlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDBInstancePlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDBInstancePlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstancePlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateDBInstancePlanResponse) SetHeaders(v map[string]*string) *UpdateDBInstancePlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateDBInstancePlanResponse) SetStatusCode(v int32) *UpdateDBInstancePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDBInstancePlanResponse) SetBody(v *UpdateDBInstancePlanResponseBody) *UpdateDBInstancePlanResponse {
	s.Body = v
	return s
}

type UpgradeDBInstanceRequest struct {
	DBInstanceClass      *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceGroupCount *string `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	InstanceSpec         *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	MasterNodeNum        *string `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SegNodeNum           *string `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	StorageSize          *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	UpgradeType          *int64  `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
}

func (s UpgradeDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceRequest) SetDBInstanceClass(v string) *UpgradeDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetDBInstanceGroupCount(v string) *UpgradeDBInstanceRequest {
	s.DBInstanceGroupCount = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetDBInstanceId(v string) *UpgradeDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetInstanceSpec(v string) *UpgradeDBInstanceRequest {
	s.InstanceSpec = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetMasterNodeNum(v string) *UpgradeDBInstanceRequest {
	s.MasterNodeNum = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetOwnerId(v int64) *UpgradeDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetPayType(v string) *UpgradeDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetRegionId(v string) *UpgradeDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetResourceGroupId(v string) *UpgradeDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetSegNodeNum(v string) *UpgradeDBInstanceRequest {
	s.SegNodeNum = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetStorageSize(v string) *UpgradeDBInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetUpgradeType(v int64) *UpgradeDBInstanceRequest {
	s.UpgradeType = &v
	return s
}

type UpgradeDBInstanceResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceResponseBody) SetDBInstanceId(v string) *UpgradeDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceResponseBody) SetOrderId(v string) *UpgradeDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeDBInstanceResponseBody) SetRequestId(v string) *UpgradeDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeDBInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceResponse) SetHeaders(v map[string]*string) *UpgradeDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBInstanceResponse) SetStatusCode(v int32) *UpgradeDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBInstanceResponse) SetBody(v *UpgradeDBInstanceResponseBody) *UpgradeDBInstanceResponse {
	s.Body = v
	return s
}

type UpgradeDBVersionRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	MajorVersion   *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	MinorVersion   *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SwitchTime     *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s UpgradeDBVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBVersionRequest) SetDBInstanceId(v string) *UpgradeDBVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetMajorVersion(v string) *UpgradeDBVersionRequest {
	s.MajorVersion = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetMinorVersion(v string) *UpgradeDBVersionRequest {
	s.MinorVersion = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetOwnerId(v int64) *UpgradeDBVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetRegionId(v string) *UpgradeDBVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetSwitchTime(v string) *UpgradeDBVersionRequest {
	s.SwitchTime = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetSwitchTimeMode(v string) *UpgradeDBVersionRequest {
	s.SwitchTimeMode = &v
	return s
}

type UpgradeDBVersionResponseBody struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeDBVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBVersionResponseBody) SetDBInstanceId(v string) *UpgradeDBVersionResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBVersionResponseBody) SetDBInstanceName(v string) *UpgradeDBVersionResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeDBVersionResponseBody) SetRequestId(v string) *UpgradeDBVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBVersionResponseBody) SetTaskId(v string) *UpgradeDBVersionResponseBody {
	s.TaskId = &v
	return s
}

type UpgradeDBVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeDBVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeDBVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBVersionResponse) SetHeaders(v map[string]*string) *UpgradeDBVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBVersionResponse) SetStatusCode(v int32) *UpgradeDBVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBVersionResponse) SetBody(v *UpgradeDBVersionResponseBody) *UpgradeDBVersionResponse {
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
		"cn-beijing":            tea.String("gpdb.aliyuncs.com"),
		"cn-hangzhou":           tea.String("gpdb.aliyuncs.com"),
		"cn-shanghai":           tea.String("gpdb.aliyuncs.com"),
		"cn-shenzhen":           tea.String("gpdb.aliyuncs.com"),
		"cn-hongkong":           tea.String("gpdb.aliyuncs.com"),
		"ap-southeast-1":        tea.String("gpdb.aliyuncs.com"),
		"us-west-1":             tea.String("gpdb.aliyuncs.com"),
		"us-east-1":             tea.String("gpdb.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("gpdb.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("gpdb.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("gpdb.aliyuncs.com"),
		"cn-qingdao":            tea.String("gpdb.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("gpdb.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("gpdb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

/**
 * @deprecated
 *
 * @param request AddBuDBInstanceRelationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddBuDBInstanceRelationResponse
 */
// Deprecated
func (client *Client) AddBuDBInstanceRelationWithOptions(request *AddBuDBInstanceRelationRequest, runtime *util.RuntimeOptions) (_result *AddBuDBInstanceRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessUnit)) {
		query["BusinessUnit"] = request.BusinessUnit
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddBuDBInstanceRelation"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddBuDBInstanceRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request AddBuDBInstanceRelationRequest
 * @return AddBuDBInstanceRelationResponse
 */
// Deprecated
func (client *Client) AddBuDBInstanceRelation(request *AddBuDBInstanceRelationRequest) (_result *AddBuDBInstanceRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddBuDBInstanceRelationResponse{}
	_body, _err := client.AddBuDBInstanceRelationWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionStringPrefix)) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AllocateInstancePublicConnection"),
		Version:     tea.String("2016-05-03"),
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

func (client *Client) CheckServiceLinkedRoleWithOptions(request *CheckServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CheckServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckServiceLinkedRole"),
		Version:     tea.String("2016-05-03"),
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

func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountDescription)) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccount"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccount(request *CreateAccountRequest) (_result *CreateAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccountResponse{}
	_body, _err := client.CreateAccountWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CreateSampleData)) {
		query["CreateSampleData"] = request.CreateSampleData
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceCategory)) {
		query["DBInstanceCategory"] = request.DBInstanceCategory
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceClass)) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceDescription)) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceGroupCount)) {
		query["DBInstanceGroupCount"] = request.DBInstanceGroupCount
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceMode)) {
		query["DBInstanceMode"] = request.DBInstanceMode
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTime)) {
		query["IdleTime"] = request.IdleTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceNetworkType)) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceSpec)) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !tea.BoolValue(util.IsUnset(request.MasterNodeNum)) {
		query["MasterNodeNum"] = request.MasterNodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIPList)) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !tea.BoolValue(util.IsUnset(request.SegNodeNum)) {
		query["SegNodeNum"] = request.SegNodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.SegStorageType)) {
		query["SegStorageType"] = request.SegStorageType
	}

	if !tea.BoolValue(util.IsUnset(request.ServerlessMode)) {
		query["ServerlessMode"] = request.ServerlessMode
	}

	if !tea.BoolValue(util.IsUnset(request.ServerlessResource)) {
		query["ServerlessResource"] = request.ServerlessResource
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSize)) {
		query["StorageSize"] = request.StorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.UsedTime)) {
		query["UsedTime"] = request.UsedTime
	}

	if !tea.BoolValue(util.IsUnset(request.VPCId)) {
		query["VPCId"] = request.VPCId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBInstance"),
		Version:     tea.String("2016-05-03"),
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

func (client *Client) CreateDBInstancePlanWithOptions(request *CreateDBInstancePlanRequest, runtime *util.RuntimeOptions) (_result *CreateDBInstancePlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanConfig)) {
		query["PlanConfig"] = request.PlanConfig
	}

	if !tea.BoolValue(util.IsUnset(request.PlanDesc)) {
		query["PlanDesc"] = request.PlanDesc
	}

	if !tea.BoolValue(util.IsUnset(request.PlanEndDate)) {
		query["PlanEndDate"] = request.PlanEndDate
	}

	if !tea.BoolValue(util.IsUnset(request.PlanName)) {
		query["PlanName"] = request.PlanName
	}

	if !tea.BoolValue(util.IsUnset(request.PlanScheduleType)) {
		query["PlanScheduleType"] = request.PlanScheduleType
	}

	if !tea.BoolValue(util.IsUnset(request.PlanStartDate)) {
		query["PlanStartDate"] = request.PlanStartDate
	}

	if !tea.BoolValue(util.IsUnset(request.PlanType)) {
		query["PlanType"] = request.PlanType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBInstancePlan"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBInstancePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBInstancePlan(request *CreateDBInstancePlanRequest) (_result *CreateDBInstancePlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBInstancePlanResponse{}
	_body, _err := client.CreateDBInstancePlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateECSDBInstanceWithOptions(request *CreateECSDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateECSDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceCategory)) {
		query["DBInstanceCategory"] = request.DBInstanceCategory
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceDescription)) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptionKey)) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptionType)) {
		query["EncryptionType"] = request.EncryptionType
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceNetworkType)) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceSpec)) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !tea.BoolValue(util.IsUnset(request.MasterNodeNum)) {
		query["MasterNodeNum"] = request.MasterNodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIPList)) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !tea.BoolValue(util.IsUnset(request.SegNodeNum)) {
		query["SegNodeNum"] = request.SegNodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.SegStorageType)) {
		query["SegStorageType"] = request.SegStorageType
	}

	if !tea.BoolValue(util.IsUnset(request.SrcDbInstanceName)) {
		query["SrcDbInstanceName"] = request.SrcDbInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSize)) {
		query["StorageSize"] = request.StorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.UsedTime)) {
		query["UsedTime"] = request.UsedTime
	}

	if !tea.BoolValue(util.IsUnset(request.VPCId)) {
		query["VPCId"] = request.VPCId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateECSDBInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateECSDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateECSDBInstance(request *CreateECSDBInstanceRequest) (_result *CreateECSDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateECSDBInstanceResponse{}
	_body, _err := client.CreateECSDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSampleDataWithOptions(request *CreateSampleDataRequest, runtime *util.RuntimeOptions) (_result *CreateSampleDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSampleData"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSampleDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSampleData(request *CreateSampleDataRequest) (_result *CreateSampleDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSampleDataResponse{}
	_body, _err := client.CreateSampleDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceLinkedRole"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBInstance"),
		Version:     tea.String("2016-05-03"),
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

func (client *Client) DeleteDBInstancePlanWithOptions(request *DeleteDBInstancePlanRequest, runtime *util.RuntimeOptions) (_result *DeleteDBInstancePlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBInstancePlan"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDBInstancePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBInstancePlan(request *DeleteDBInstancePlanRequest) (_result *DeleteDBInstancePlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBInstancePlanResponse{}
	_body, _err := client.DeleteDBInstancePlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request DeleteDatabaseRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDatabaseResponse
 */
// Deprecated
func (client *Client) DeleteDatabaseWithOptions(request *DeleteDatabaseRequest, runtime *util.RuntimeOptions) (_result *DeleteDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDatabase"),
		Version:     tea.String("2016-05-03"),
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

/**
 * @deprecated
 *
 * @param request DeleteDatabaseRequest
 * @return DeleteDatabaseResponse
 */
// Deprecated
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

func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccounts"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccounts(request *DescribeAccountsRequest) (_result *DescribeAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DescribeAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableResourcesWithOptions(request *DescribeAvailableResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableResources"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAvailableResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableResources(request *DescribeAvailableResourcesRequest) (_result *DescribeAvailableResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourcesResponse{}
	_body, _err := client.DescribeAvailableResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupPolicy"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPolicy(request *DescribeBackupPolicyRequest) (_result *DescribeBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DescribeBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterNodeWithOptions(request *DescribeDBClusterNodeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeType)) {
		query["NodeType"] = request.NodeType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterNode"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterNode(request *DescribeDBClusterNodeRequest) (_result *DescribeDBClusterNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterNodeResponse{}
	_body, _err := client.DescribeDBClusterNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterPerformanceWithOptions(request *DescribeDBClusterPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.NodeType)) {
		query["NodeType"] = request.NodeType
	}

	if !tea.BoolValue(util.IsUnset(request.Nodes)) {
		query["Nodes"] = request.Nodes
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterPerformance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterPerformance(request *DescribeDBClusterPerformanceRequest) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DescribeDBClusterPerformanceWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceAttribute"),
		Version:     tea.String("2016-05-03"),
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

func (client *Client) DescribeDBInstanceDataBloatWithOptions(request *DescribeDBInstanceDataBloatRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceDataBloatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceDataBloat"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceDataBloatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceDataBloat(request *DescribeDBInstanceDataBloatRequest) (_result *DescribeDBInstanceDataBloatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceDataBloatResponse{}
	_body, _err := client.DescribeDBInstanceDataBloatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceDataSkewWithOptions(request *DescribeDBInstanceDataSkewRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceDataSkewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceDataSkew"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceDataSkewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceDataSkew(request *DescribeDBInstanceDataSkewRequest) (_result *DescribeDBInstanceDataSkewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceDataSkewResponse{}
	_body, _err := client.DescribeDBInstanceDataSkewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceDiagnosisSummaryWithOptions(request *DescribeDBInstanceDiagnosisSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceDiagnosisSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RolePreferd)) {
		query["RolePreferd"] = request.RolePreferd
	}

	if !tea.BoolValue(util.IsUnset(request.StartStatus)) {
		query["StartStatus"] = request.StartStatus
	}

	if !tea.BoolValue(util.IsUnset(request.SyncMode)) {
		query["SyncMode"] = request.SyncMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceDiagnosisSummary"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceDiagnosisSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceDiagnosisSummary(request *DescribeDBInstanceDiagnosisSummaryRequest) (_result *DescribeDBInstanceDiagnosisSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceDiagnosisSummaryResponse{}
	_body, _err := client.DescribeDBInstanceDiagnosisSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceErrorLogWithOptions(request *DescribeDBInstanceErrorLogRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceErrorLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		query["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		query["Keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.LogLevel)) {
		query["LogLevel"] = request.LogLevel
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceErrorLog"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceErrorLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceErrorLog(request *DescribeDBInstanceErrorLogRequest) (_result *DescribeDBInstanceErrorLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceErrorLogResponse{}
	_body, _err := client.DescribeDBInstanceErrorLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceIPArrayListWithOptions(request *DescribeDBInstanceIPArrayListRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceIPArrayListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceIPArrayList"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceIPArrayListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceIPArrayList(request *DescribeDBInstanceIPArrayListRequest) (_result *DescribeDBInstanceIPArrayListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceIPArrayListResponse{}
	_body, _err := client.DescribeDBInstanceIPArrayListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceIndexUsageWithOptions(request *DescribeDBInstanceIndexUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceIndexUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceIndexUsage"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceIndexUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceIndexUsage(request *DescribeDBInstanceIndexUsageRequest) (_result *DescribeDBInstanceIndexUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceIndexUsageResponse{}
	_body, _err := client.DescribeDBInstanceIndexUsageWithOptions(request, runtime)
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceNetInfo"),
		Version:     tea.String("2016-05-03"),
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

func (client *Client) DescribeDBInstanceOnECSAttributeWithOptions(request *DescribeDBInstanceOnECSAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceOnECSAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceOnECSAttribute"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceOnECSAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceOnECSAttribute(request *DescribeDBInstanceOnECSAttributeRequest) (_result *DescribeDBInstanceOnECSAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceOnECSAttributeResponse{}
	_body, _err := client.DescribeDBInstanceOnECSAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstancePerformanceWithOptions(request *DescribeDBInstancePerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancePerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstancePerformance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstancePerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstancePerformance(request *DescribeDBInstancePerformanceRequest) (_result *DescribeDBInstancePerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstancePerformanceResponse{}
	_body, _err := client.DescribeDBInstancePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstancePlansWithOptions(request *DescribeDBInstancePlansRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancePlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanCreateDate)) {
		query["PlanCreateDate"] = request.PlanCreateDate
	}

	if !tea.BoolValue(util.IsUnset(request.PlanDesc)) {
		query["PlanDesc"] = request.PlanDesc
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanScheduleType)) {
		query["PlanScheduleType"] = request.PlanScheduleType
	}

	if !tea.BoolValue(util.IsUnset(request.PlanType)) {
		query["PlanType"] = request.PlanType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstancePlans"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstancePlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstancePlans(request *DescribeDBInstancePlansRequest) (_result *DescribeDBInstancePlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstancePlansResponse{}
	_body, _err := client.DescribeDBInstancePlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceSQLPatternsWithOptions(request *DescribeDBInstanceSQLPatternsRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceSQLPatternsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeywords)) {
		query["QueryKeywords"] = request.QueryKeywords
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIP)) {
		query["SourceIP"] = request.SourceIP
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceSQLPatterns"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceSQLPatternsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceSQLPatterns(request *DescribeDBInstanceSQLPatternsRequest) (_result *DescribeDBInstanceSQLPatternsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceSQLPatternsResponse{}
	_body, _err := client.DescribeDBInstanceSQLPatternsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceSSLWithOptions(request *DescribeDBInstanceSSLRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceSSL"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceSSL(request *DescribeDBInstanceSSLRequest) (_result *DescribeDBInstanceSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceSSLResponse{}
	_body, _err := client.DescribeDBInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstancesWithOptions(tmpReq *DescribeDBInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeDBInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DBInstanceCategories)) {
		request.DBInstanceCategoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBInstanceCategories, tea.String("DBInstanceCategories"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DBInstanceModes)) {
		request.DBInstanceModesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBInstanceModes, tea.String("DBInstanceModes"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DBInstanceStatuses)) {
		request.DBInstanceStatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBInstanceStatuses, tea.String("DBInstanceStatuses"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceDeployTypes)) {
		request.InstanceDeployTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceDeployTypes, tea.String("InstanceDeployTypes"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceCategoriesShrink)) {
		query["DBInstanceCategories"] = request.DBInstanceCategoriesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceDescription)) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceIds)) {
		query["DBInstanceIds"] = request.DBInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceModesShrink)) {
		query["DBInstanceModes"] = request.DBInstanceModesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceStatusesShrink)) {
		query["DBInstanceStatuses"] = request.DBInstanceStatusesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceDeployTypesShrink)) {
		query["InstanceDeployTypes"] = request.InstanceDeployTypesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceNetworkType)) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstances"),
		Version:     tea.String("2016-05-03"),
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

func (client *Client) DescribeDataBackupsWithOptions(request *DescribeDataBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDataBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupMode)) {
		query["BackupMode"] = request.BackupMode
	}

	if !tea.BoolValue(util.IsUnset(request.BackupStatus)) {
		query["BackupStatus"] = request.BackupStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataBackups"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataBackups(request *DescribeDataBackupsRequest) (_result *DescribeDataBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataBackupsResponse{}
	_body, _err := client.DescribeDataBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataShareInstancesWithOptions(request *DescribeDataShareInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDataShareInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.SearchValue)) {
		query["SearchValue"] = request.SearchValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataShareInstances"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataShareInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataShareInstances(request *DescribeDataShareInstancesRequest) (_result *DescribeDataShareInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataShareInstancesResponse{}
	_body, _err := client.DescribeDataShareInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataSharePerformanceWithOptions(request *DescribeDataSharePerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDataSharePerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataSharePerformance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataSharePerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataSharePerformance(request *DescribeDataSharePerformanceRequest) (_result *DescribeDataSharePerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataSharePerformanceResponse{}
	_body, _err := client.DescribeDataSharePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisDimensionsWithOptions(request *DescribeDiagnosisDimensionsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisDimensionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisDimensions"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisDimensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisDimensions(request *DescribeDiagnosisDimensionsRequest) (_result *DescribeDiagnosisDimensionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisDimensionsResponse{}
	_body, _err := client.DescribeDiagnosisDimensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisMonitorPerformanceWithOptions(request *DescribeDiagnosisMonitorPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisMonitorPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisMonitorPerformance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisMonitorPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisMonitorPerformance(request *DescribeDiagnosisMonitorPerformanceRequest) (_result *DescribeDiagnosisMonitorPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisMonitorPerformanceResponse{}
	_body, _err := client.DescribeDiagnosisMonitorPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisRecordsWithOptions(request *DescribeDiagnosisRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisRecords"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisRecords(request *DescribeDiagnosisRecordsRequest) (_result *DescribeDiagnosisRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisRecordsResponse{}
	_body, _err := client.DescribeDiagnosisRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisSQLInfoWithOptions(request *DescribeDiagnosisSQLInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisSQLInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.QueryID)) {
		query["QueryID"] = request.QueryID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisSQLInfo"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisSQLInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisSQLInfo(request *DescribeDiagnosisSQLInfoRequest) (_result *DescribeDiagnosisSQLInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisSQLInfoResponse{}
	_body, _err := client.DescribeDiagnosisSQLInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDownloadRecordsWithOptions(request *DescribeDownloadRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDownloadRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDownloadRecords"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDownloadRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDownloadRecords(request *DescribeDownloadRecordsRequest) (_result *DescribeDownloadRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDownloadRecordsResponse{}
	_body, _err := client.DescribeDownloadRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHealthStatusWithOptions(request *DescribeHealthStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeHealthStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHealthStatus"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHealthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHealthStatus(request *DescribeHealthStatusRequest) (_result *DescribeHealthStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHealthStatusResponse{}
	_body, _err := client.DescribeHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogBackupsWithOptions(request *DescribeLogBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeLogBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogBackups"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogBackups(request *DescribeLogBackupsRequest) (_result *DescribeLogBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogBackupsResponse{}
	_body, _err := client.DescribeLogBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeModifyParameterLogWithOptions(request *DescribeModifyParameterLogRequest, runtime *util.RuntimeOptions) (_result *DescribeModifyParameterLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeModifyParameterLog"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeModifyParameterLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeModifyParameterLog(request *DescribeModifyParameterLogRequest) (_result *DescribeModifyParameterLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeModifyParameterLogResponse{}
	_body, _err := client.DescribeModifyParameterLogWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeParameters"),
		Version:     tea.String("2016-05-03"),
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

func (client *Client) DescribeRdsVSwitchsWithOptions(request *DescribeRdsVSwitchsRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsVSwitchsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
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
		Action:      tea.String("DescribeRdsVSwitchs"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdsVSwitchsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRdsVSwitchs(request *DescribeRdsVSwitchsRequest) (_result *DescribeRdsVSwitchsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsVSwitchsResponse{}
	_body, _err := client.DescribeRdsVSwitchsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRdsVpcsWithOptions(request *DescribeRdsVpcsRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsVpcsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRdsVpcs"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdsVpcsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRdsVpcs(request *DescribeRdsVpcsRequest) (_result *DescribeRdsVpcsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsVpcsResponse{}
	_body, _err := client.DescribeRdsVpcsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2016-05-03"),
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

func (client *Client) DescribeResourceUsageWithOptions(request *DescribeResourceUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeResourceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceUsage"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceUsage(request *DescribeResourceUsageRequest) (_result *DescribeResourceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourceUsageResponse{}
	_body, _err := client.DescribeResourceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLCollectorPolicyWithOptions(request *DescribeSQLCollectorPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLCollectorPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLCollectorPolicy"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLCollectorPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLCollectorPolicy(request *DescribeSQLCollectorPolicyRequest) (_result *DescribeSQLCollectorPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLCollectorPolicyResponse{}
	_body, _err := client.DescribeSQLCollectorPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogByQueryIdWithOptions(request *DescribeSQLLogByQueryIdRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogByQueryIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLLogByQueryId"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLLogByQueryIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogByQueryId(request *DescribeSQLLogByQueryIdRequest) (_result *DescribeSQLLogByQueryIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogByQueryIdResponse{}
	_body, _err := client.DescribeSQLLogByQueryIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogCountWithOptions(request *DescribeSQLLogCountRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteCost)) {
		query["ExecuteCost"] = request.ExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteState)) {
		query["ExecuteState"] = request.ExecuteState
	}

	if !tea.BoolValue(util.IsUnset(request.MaxExecuteCost)) {
		query["MaxExecuteCost"] = request.MaxExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.MinExecuteCost)) {
		query["MinExecuteCost"] = request.MinExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.OperationClass)) {
		query["OperationClass"] = request.OperationClass
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeywords)) {
		query["QueryKeywords"] = request.QueryKeywords
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIP)) {
		query["SourceIP"] = request.SourceIP
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLLogCount"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLLogCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogCount(request *DescribeSQLLogCountRequest) (_result *DescribeSQLLogCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogCountResponse{}
	_body, _err := client.DescribeSQLLogCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogFilesWithOptions(request *DescribeSQLLogFilesRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLLogFiles"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLLogFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogFiles(request *DescribeSQLLogFilesRequest) (_result *DescribeSQLLogFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogFilesResponse{}
	_body, _err := client.DescribeSQLLogFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogRecordsWithOptions(request *DescribeSQLLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Form)) {
		query["Form"] = request.Form
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeywords)) {
		query["QueryKeywords"] = request.QueryKeywords
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLLogRecords"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogRecords(request *DescribeSQLLogRecordsRequest) (_result *DescribeSQLLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogRecordsResponse{}
	_body, _err := client.DescribeSQLLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogsWithOptions(request *DescribeSQLLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteCost)) {
		query["ExecuteCost"] = request.ExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteState)) {
		query["ExecuteState"] = request.ExecuteState
	}

	if !tea.BoolValue(util.IsUnset(request.MaxExecuteCost)) {
		query["MaxExecuteCost"] = request.MaxExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.MinExecuteCost)) {
		query["MinExecuteCost"] = request.MinExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.OperationClass)) {
		query["OperationClass"] = request.OperationClass
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeywords)) {
		query["QueryKeywords"] = request.QueryKeywords
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIP)) {
		query["SourceIP"] = request.SourceIP
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLLogs"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogs(request *DescribeSQLLogsRequest) (_result *DescribeSQLLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogsResponse{}
	_body, _err := client.DescribeSQLLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogsOnSliceWithOptions(request *DescribeSQLLogsOnSliceRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogsOnSliceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteState)) {
		query["ExecuteState"] = request.ExecuteState
	}

	if !tea.BoolValue(util.IsUnset(request.MaxExecuteCost)) {
		query["MaxExecuteCost"] = request.MaxExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.MinExecuteCost)) {
		query["MinExecuteCost"] = request.MinExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
	}

	if !tea.BoolValue(util.IsUnset(request.SliceId)) {
		query["SliceId"] = request.SliceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLLogsOnSlice"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLLogsOnSliceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogsOnSlice(request *DescribeSQLLogsOnSliceRequest) (_result *DescribeSQLLogsOnSliceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogsOnSliceResponse{}
	_body, _err := client.DescribeSQLLogsOnSliceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogsV2WithOptions(request *DescribeSQLLogsV2Request, runtime *util.RuntimeOptions) (_result *DescribeSQLLogsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteCost)) {
		query["ExecuteCost"] = request.ExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteState)) {
		query["ExecuteState"] = request.ExecuteState
	}

	if !tea.BoolValue(util.IsUnset(request.MaxExecuteCost)) {
		query["MaxExecuteCost"] = request.MaxExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.MinExecuteCost)) {
		query["MinExecuteCost"] = request.MinExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.OperationClass)) {
		query["OperationClass"] = request.OperationClass
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeywords)) {
		query["QueryKeywords"] = request.QueryKeywords
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIP)) {
		query["SourceIP"] = request.SourceIP
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLLogsV2"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLLogsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogsV2(request *DescribeSQLLogsV2Request) (_result *DescribeSQLLogsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogsV2Response{}
	_body, _err := client.DescribeSQLLogsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSampleDataWithOptions(request *DescribeSampleDataRequest, runtime *util.RuntimeOptions) (_result *DescribeSampleDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSampleData"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSampleDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSampleData(request *DescribeSampleDataRequest) (_result *DescribeSampleDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSampleDataResponse{}
	_body, _err := client.DescribeSampleDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowLogRecordsWithOptions(request *DescribeSlowLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		query["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowLogRecords"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowLogRecords(request *DescribeSlowLogRecordsRequest) (_result *DescribeSlowLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DescribeSlowLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowSQLLogsWithOptions(request *DescribeSlowSQLLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowSQLLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteState)) {
		query["ExecuteState"] = request.ExecuteState
	}

	if !tea.BoolValue(util.IsUnset(request.MaxExecuteCost)) {
		query["MaxExecuteCost"] = request.MaxExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.MinExecuteCost)) {
		query["MinExecuteCost"] = request.MinExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.OperationClass)) {
		query["OperationClass"] = request.OperationClass
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeywords)) {
		query["QueryKeywords"] = request.QueryKeywords
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIP)) {
		query["SourceIP"] = request.SourceIP
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowSQLLogs"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlowSQLLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowSQLLogs(request *DescribeSlowSQLLogsRequest) (_result *DescribeSlowSQLLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowSQLLogsResponse{}
	_body, _err := client.DescribeSlowSQLLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSpecificationWithOptions(request *DescribeSpecificationRequest, runtime *util.RuntimeOptions) (_result *DescribeSpecificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CpuCores)) {
		query["CpuCores"] = request.CpuCores
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	if !tea.BoolValue(util.IsUnset(request.TotalNodeNum)) {
		query["TotalNodeNum"] = request.TotalNodeNum
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSpecification"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSpecificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSpecification(request *DescribeSpecificationRequest) (_result *DescribeSpecificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSpecificationResponse{}
	_body, _err := client.DescribeSpecificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSupportFeaturesWithOptions(request *DescribeSupportFeaturesRequest, runtime *util.RuntimeOptions) (_result *DescribeSupportFeaturesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSupportFeatures"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSupportFeaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSupportFeatures(request *DescribeSupportFeaturesRequest) (_result *DescribeSupportFeaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSupportFeaturesResponse{}
	_body, _err := client.DescribeSupportFeaturesWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTags"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeUserEncryptionKeyListWithOptions(request *DescribeUserEncryptionKeyListRequest, runtime *util.RuntimeOptions) (_result *DescribeUserEncryptionKeyListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserEncryptionKeyList"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserEncryptionKeyListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserEncryptionKeyList(request *DescribeUserEncryptionKeyListRequest) (_result *DescribeUserEncryptionKeyListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserEncryptionKeyListResponse{}
	_body, _err := client.DescribeUserEncryptionKeyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWaitingSQLInfoWithOptions(request *DescribeWaitingSQLInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeWaitingSQLInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.PID)) {
		query["PID"] = request.PID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWaitingSQLInfo"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWaitingSQLInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWaitingSQLInfo(request *DescribeWaitingSQLInfoRequest) (_result *DescribeWaitingSQLInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWaitingSQLInfoResponse{}
	_body, _err := client.DescribeWaitingSQLInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWaitingSQLRecordsWithOptions(request *DescribeWaitingSQLRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeWaitingSQLRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWaitingSQLRecords"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWaitingSQLRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWaitingSQLRecords(request *DescribeWaitingSQLRecordsRequest) (_result *DescribeWaitingSQLRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWaitingSQLRecordsResponse{}
	_body, _err := client.DescribeWaitingSQLRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadDiagnosisRecordsWithOptions(request *DownloadDiagnosisRecordsRequest, runtime *util.RuntimeOptions) (_result *DownloadDiagnosisRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DownloadDiagnosisRecords"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadDiagnosisRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadDiagnosisRecords(request *DownloadDiagnosisRecordsRequest) (_result *DownloadDiagnosisRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadDiagnosisRecordsResponse{}
	_body, _err := client.DownloadDiagnosisRecordsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyAccountDescriptionWithOptions(request *ModifyAccountDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountDescription)) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountDescription"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (_result *ModifyAccountDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.ModifyAccountDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupPolicyWithOptions(request *ModifyBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupRetentionPeriod)) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableRecoveryPoint)) {
		query["EnableRecoveryPoint"] = request.EnableRecoveryPoint
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupPeriod)) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupTime)) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	if !tea.BoolValue(util.IsUnset(request.RecoveryPointPeriod)) {
		query["RecoveryPointPeriod"] = request.RecoveryPointPeriod
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyBackupPolicy"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupPolicy(request *ModifyBackupPolicyRequest) (_result *ModifyBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.ModifyBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceConnectionModeWithOptions(request *ModifyDBInstanceConnectionModeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceConnectionModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionMode)) {
		query["ConnectionMode"] = request.ConnectionMode
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceConnectionMode"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceConnectionModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceConnectionMode(request *ModifyDBInstanceConnectionModeRequest) (_result *ModifyDBInstanceConnectionModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceConnectionModeResponse{}
	_body, _err := client.ModifyDBInstanceConnectionModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceConnectionStringWithOptions(request *ModifyDBInstanceConnectionStringRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceConnectionStringResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionStringPrefix)) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentConnectionString)) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceConnectionString"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceConnectionStringResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceConnectionString(request *ModifyDBInstanceConnectionStringRequest) (_result *ModifyDBInstanceConnectionStringResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceConnectionStringResponse{}
	_body, _err := client.ModifyDBInstanceConnectionStringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceDescriptionWithOptions(request *ModifyDBInstanceDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceDescription)) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceDescription"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceDescription(request *ModifyDBInstanceDescriptionRequest) (_result *ModifyDBInstanceDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceDescriptionResponse{}
	_body, _err := client.ModifyDBInstanceDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceMaintainTimeWithOptions(request *ModifyDBInstanceMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceMaintainTime"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceMaintainTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceMaintainTime(request *ModifyDBInstanceMaintainTimeRequest) (_result *ModifyDBInstanceMaintainTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceMaintainTimeResponse{}
	_body, _err := client.ModifyDBInstanceMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceNetworkTypeWithOptions(request *ModifyDBInstanceNetworkTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceNetworkTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceNetworkType)) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.VPCId)) {
		query["VPCId"] = request.VPCId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceNetworkType"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceNetworkTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceNetworkType(request *ModifyDBInstanceNetworkTypeRequest) (_result *ModifyDBInstanceNetworkTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceNetworkTypeResponse{}
	_body, _err := client.ModifyDBInstanceNetworkTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceResourceGroupWithOptions(request *ModifyDBInstanceResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceResourceGroup"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceResourceGroup(request *ModifyDBInstanceResourceGroupRequest) (_result *ModifyDBInstanceResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceResourceGroupResponse{}
	_body, _err := client.ModifyDBInstanceResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceSSLWithOptions(request *ModifyDBInstanceSSLRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceSSLResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SSLEnabled)) {
		query["SSLEnabled"] = request.SSLEnabled
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceSSL"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceSSL(request *ModifyDBInstanceSSLRequest) (_result *ModifyDBInstanceSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceSSLResponse{}
	_body, _err := client.ModifyDBInstanceSSLWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ForceRestartInstance)) {
		query["ForceRestartInstance"] = request.ForceRestartInstance
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyParameters"),
		Version:     tea.String("2016-05-03"),
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

func (client *Client) ModifySQLCollectorPolicyWithOptions(request *ModifySQLCollectorPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifySQLCollectorPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SQLCollectorStatus)) {
		query["SQLCollectorStatus"] = request.SQLCollectorStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySQLCollectorPolicy"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySQLCollectorPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySQLCollectorPolicy(request *ModifySQLCollectorPolicyRequest) (_result *ModifySQLCollectorPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySQLCollectorPolicyResponse{}
	_body, _err := client.ModifySQLCollectorPolicyWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceIPArrayAttribute)) {
		query["DBInstanceIPArrayAttribute"] = request.DBInstanceIPArrayAttribute
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceIPArrayName)) {
		query["DBInstanceIPArrayName"] = request.DBInstanceIPArrayName
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIPList)) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySecurityIps"),
		Version:     tea.String("2016-05-03"),
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

func (client *Client) PauseInstanceWithOptions(request *PauseInstanceRequest, runtime *util.RuntimeOptions) (_result *PauseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PauseInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PauseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PauseInstance(request *PauseInstanceRequest) (_result *PauseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PauseInstanceResponse{}
	_body, _err := client.PauseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebalanceDBInstanceWithOptions(request *RebalanceDBInstanceRequest, runtime *util.RuntimeOptions) (_result *RebalanceDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebalanceDBInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebalanceDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebalanceDBInstance(request *RebalanceDBInstanceRequest) (_result *RebalanceDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebalanceDBInstanceResponse{}
	_body, _err := client.RebalanceDBInstanceWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentConnectionString)) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseInstancePublicConnection"),
		Version:     tea.String("2016-05-03"),
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAccountPassword"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

func (client *Client) RestartDBInstanceWithOptions(request *RestartDBInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartDBInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartDBInstance(request *RestartDBInstanceRequest) (_result *RestartDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartDBInstanceResponse{}
	_body, _err := client.RestartDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeInstanceWithOptions(request *ResumeInstanceRequest, runtime *util.RuntimeOptions) (_result *ResumeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeInstance(request *ResumeInstanceRequest) (_result *ResumeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeInstanceResponse{}
	_body, _err := client.ResumeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDBInstancePlanStatusWithOptions(request *SetDBInstancePlanStatusRequest, runtime *util.RuntimeOptions) (_result *SetDBInstancePlanStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanStatus)) {
		query["PlanStatus"] = request.PlanStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDBInstancePlanStatus"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDBInstancePlanStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDBInstancePlanStatus(request *SetDBInstancePlanStatusRequest) (_result *SetDBInstancePlanStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDBInstancePlanStatusResponse{}
	_body, _err := client.SetDBInstancePlanStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDataShareInstanceWithOptions(tmpReq *SetDataShareInstanceRequest, runtime *util.RuntimeOptions) (_result *SetDataShareInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetDataShareInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceList)) {
		request.InstanceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceList, tea.String("InstanceList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceListShrink)) {
		query["InstanceList"] = request.InstanceListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDataShareInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDataShareInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDataShareInstance(request *SetDataShareInstanceRequest) (_result *SetDataShareInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDataShareInstanceResponse{}
	_body, _err := client.SetDataShareInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchDBInstanceNetTypeWithOptions(request *SwitchDBInstanceNetTypeRequest, runtime *util.RuntimeOptions) (_result *SwitchDBInstanceNetTypeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchDBInstanceNetType"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwitchDBInstanceNetTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchDBInstanceNetType(request *SwitchDBInstanceNetTypeRequest) (_result *SwitchDBInstanceNetTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchDBInstanceNetTypeResponse{}
	_body, _err := client.SwitchDBInstanceNetTypeWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UnloadSampleDataWithOptions(request *UnloadSampleDataRequest, runtime *util.RuntimeOptions) (_result *UnloadSampleDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnloadSampleData"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnloadSampleDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnloadSampleData(request *UnloadSampleDataRequest) (_result *UnloadSampleDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnloadSampleDataResponse{}
	_body, _err := client.UnloadSampleDataWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateDBInstancePlanWithOptions(request *UpdateDBInstancePlanRequest, runtime *util.RuntimeOptions) (_result *UpdateDBInstancePlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanConfig)) {
		query["PlanConfig"] = request.PlanConfig
	}

	if !tea.BoolValue(util.IsUnset(request.PlanDesc)) {
		query["PlanDesc"] = request.PlanDesc
	}

	if !tea.BoolValue(util.IsUnset(request.PlanEndDate)) {
		query["PlanEndDate"] = request.PlanEndDate
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanName)) {
		query["PlanName"] = request.PlanName
	}

	if !tea.BoolValue(util.IsUnset(request.PlanStartDate)) {
		query["PlanStartDate"] = request.PlanStartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDBInstancePlan"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDBInstancePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDBInstancePlan(request *UpdateDBInstancePlanRequest) (_result *UpdateDBInstancePlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDBInstancePlanResponse{}
	_body, _err := client.UpdateDBInstancePlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeDBInstanceWithOptions(request *UpgradeDBInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceClass)) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceGroupCount)) {
		query["DBInstanceGroupCount"] = request.DBInstanceGroupCount
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceSpec)) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !tea.BoolValue(util.IsUnset(request.MasterNodeNum)) {
		query["MasterNodeNum"] = request.MasterNodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SegNodeNum)) {
		query["SegNodeNum"] = request.SegNodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSize)) {
		query["StorageSize"] = request.StorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeType)) {
		query["UpgradeType"] = request.UpgradeType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeDBInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (_result *UpgradeDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeDBInstanceResponse{}
	_body, _err := client.UpgradeDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeDBVersionWithOptions(request *UpgradeDBVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeDBVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MajorVersion)) {
		query["MajorVersion"] = request.MajorVersion
	}

	if !tea.BoolValue(util.IsUnset(request.MinorVersion)) {
		query["MinorVersion"] = request.MinorVersion
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTime)) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTimeMode)) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeDBVersion"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeDBVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeDBVersion(request *UpgradeDBVersionRequest) (_result *UpgradeDBVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeDBVersionResponse{}
	_body, _err := client.UpgradeDBVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
