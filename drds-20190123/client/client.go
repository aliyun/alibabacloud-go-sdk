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

type ChangeAccountPasswordRequest struct {
	// The name of the member account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The new password.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ChangeAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ChangeAccountPasswordRequest) SetAccountName(v string) *ChangeAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ChangeAccountPasswordRequest) SetDrdsInstanceId(v string) *ChangeAccountPasswordRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ChangeAccountPasswordRequest) SetPassword(v string) *ChangeAccountPasswordRequest {
	s.Password = &v
	return s
}

type ChangeAccountPasswordResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeAccountPasswordResponseBody) SetRequestId(v string) *ChangeAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeAccountPasswordResponseBody) SetSuccess(v bool) *ChangeAccountPasswordResponseBody {
	s.Success = &v
	return s
}

type ChangeAccountPasswordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ChangeAccountPasswordResponse) SetHeaders(v map[string]*string) *ChangeAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ChangeAccountPasswordResponse) SetStatusCode(v int32) *ChangeAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeAccountPasswordResponse) SetBody(v *ChangeAccountPasswordResponseBody) *ChangeAccountPasswordResponse {
	s.Body = v
	return s
}

type ChangeInstanceAzoneRequest struct {
	ChangeVSwitch *bool `json:"ChangeVSwitch,omitempty" xml:"ChangeVSwitch,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	DrdsRegionId *string `json:"DrdsRegionId,omitempty" xml:"DrdsRegionId,omitempty"`
	NewVSwitch   *string `json:"NewVSwitch,omitempty" xml:"NewVSwitch,omitempty"`
	// The source zone of the PolarDB-X 1.0 instance.
	OriginAzoneId *string `json:"OriginAzoneId,omitempty" xml:"OriginAzoneId,omitempty"`
	// The destination zone to which you want to modify
	TargetAzoneId *string `json:"TargetAzoneId,omitempty" xml:"TargetAzoneId,omitempty"`
}

func (s ChangeInstanceAzoneRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeInstanceAzoneRequest) GoString() string {
	return s.String()
}

func (s *ChangeInstanceAzoneRequest) SetChangeVSwitch(v bool) *ChangeInstanceAzoneRequest {
	s.ChangeVSwitch = &v
	return s
}

func (s *ChangeInstanceAzoneRequest) SetDrdsInstanceId(v string) *ChangeInstanceAzoneRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ChangeInstanceAzoneRequest) SetDrdsRegionId(v string) *ChangeInstanceAzoneRequest {
	s.DrdsRegionId = &v
	return s
}

func (s *ChangeInstanceAzoneRequest) SetNewVSwitch(v string) *ChangeInstanceAzoneRequest {
	s.NewVSwitch = &v
	return s
}

func (s *ChangeInstanceAzoneRequest) SetOriginAzoneId(v string) *ChangeInstanceAzoneRequest {
	s.OriginAzoneId = &v
	return s
}

func (s *ChangeInstanceAzoneRequest) SetTargetAzoneId(v string) *ChangeInstanceAzoneRequest {
	s.TargetAzoneId = &v
	return s
}

type ChangeInstanceAzoneResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeInstanceAzoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeInstanceAzoneResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeInstanceAzoneResponseBody) SetRequestId(v string) *ChangeInstanceAzoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeInstanceAzoneResponseBody) SetSuccess(v bool) *ChangeInstanceAzoneResponseBody {
	s.Success = &v
	return s
}

type ChangeInstanceAzoneResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeInstanceAzoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeInstanceAzoneResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeInstanceAzoneResponse) GoString() string {
	return s.String()
}

func (s *ChangeInstanceAzoneResponse) SetHeaders(v map[string]*string) *ChangeInstanceAzoneResponse {
	s.Headers = v
	return s
}

func (s *ChangeInstanceAzoneResponse) SetStatusCode(v int32) *ChangeInstanceAzoneResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeInstanceAzoneResponse) SetBody(v *ChangeInstanceAzoneResponseBody) *ChangeInstanceAzoneResponse {
	s.Body = v
	return s
}

type CheckDrdsDbNameRequest struct {
	// DRDS database name
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// DRDS instance ID
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s CheckDrdsDbNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDrdsDbNameRequest) GoString() string {
	return s.String()
}

func (s *CheckDrdsDbNameRequest) SetDbName(v string) *CheckDrdsDbNameRequest {
	s.DbName = &v
	return s
}

func (s *CheckDrdsDbNameRequest) SetDrdsInstanceId(v string) *CheckDrdsDbNameRequest {
	s.DrdsInstanceId = &v
	return s
}

type CheckDrdsDbNameResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the DRDS database name is valid. Valid values: true: The database name is valid. false: the database name is invalid.
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the call is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckDrdsDbNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDrdsDbNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDrdsDbNameResponseBody) SetRequestId(v string) *CheckDrdsDbNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDrdsDbNameResponseBody) SetResult(v bool) *CheckDrdsDbNameResponseBody {
	s.Result = &v
	return s
}

func (s *CheckDrdsDbNameResponseBody) SetSuccess(v bool) *CheckDrdsDbNameResponseBody {
	s.Success = &v
	return s
}

type CheckDrdsDbNameResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDrdsDbNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDrdsDbNameResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDrdsDbNameResponse) GoString() string {
	return s.String()
}

func (s *CheckDrdsDbNameResponse) SetHeaders(v map[string]*string) *CheckDrdsDbNameResponse {
	s.Headers = v
	return s
}

func (s *CheckDrdsDbNameResponse) SetStatusCode(v int32) *CheckDrdsDbNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDrdsDbNameResponse) SetBody(v *CheckDrdsDbNameResponseBody) *CheckDrdsDbNameResponse {
	s.Body = v
	return s
}

type CheckExpandStatusRequest struct {
	// The name of the PolarDB-X database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s CheckExpandStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckExpandStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusRequest) SetDbName(v string) *CheckExpandStatusRequest {
	s.DbName = &v
	return s
}

func (s *CheckExpandStatusRequest) SetDrdsInstanceId(v string) *CheckExpandStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

type CheckExpandStatusResponseBody struct {
	// The result of the verification.
	Data *CheckExpandStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckExpandStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckExpandStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusResponseBody) SetData(v *CheckExpandStatusResponseBodyData) *CheckExpandStatusResponseBody {
	s.Data = v
	return s
}

func (s *CheckExpandStatusResponseBody) SetRequestId(v string) *CheckExpandStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckExpandStatusResponseBody) SetSuccess(v bool) *CheckExpandStatusResponseBody {
	s.Success = &v
	return s
}

type CheckExpandStatusResponseBodyData struct {
	// Indicates whether scale-out operations can be performed on the database.
	IsActive *bool `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	// The additional information.
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
}

func (s CheckExpandStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckExpandStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusResponseBodyData) SetIsActive(v bool) *CheckExpandStatusResponseBodyData {
	s.IsActive = &v
	return s
}

func (s *CheckExpandStatusResponseBodyData) SetMsg(v string) *CheckExpandStatusResponseBodyData {
	s.Msg = &v
	return s
}

type CheckExpandStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckExpandStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckExpandStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckExpandStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusResponse) SetHeaders(v map[string]*string) *CheckExpandStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckExpandStatusResponse) SetStatusCode(v int32) *CheckExpandStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckExpandStatusResponse) SetBody(v *CheckExpandStatusResponseBody) *CheckExpandStatusResponse {
	s.Body = v
	return s
}

type CheckSqlAuditEnableStatusRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s CheckSqlAuditEnableStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckSqlAuditEnableStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckSqlAuditEnableStatusRequest) SetDbName(v string) *CheckSqlAuditEnableStatusRequest {
	s.DbName = &v
	return s
}

func (s *CheckSqlAuditEnableStatusRequest) SetDrdsInstanceId(v string) *CheckSqlAuditEnableStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

type CheckSqlAuditEnableStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the SQL audit feature. Valid values:
	//
	// *   enabled: The SQL audit feature is enabled.
	// *   disabled: The SQL audit feature is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckSqlAuditEnableStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckSqlAuditEnableStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSqlAuditEnableStatusResponseBody) SetRequestId(v string) *CheckSqlAuditEnableStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSqlAuditEnableStatusResponseBody) SetStatus(v string) *CheckSqlAuditEnableStatusResponseBody {
	s.Status = &v
	return s
}

func (s *CheckSqlAuditEnableStatusResponseBody) SetSuccess(v bool) *CheckSqlAuditEnableStatusResponseBody {
	s.Success = &v
	return s
}

type CheckSqlAuditEnableStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSqlAuditEnableStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSqlAuditEnableStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckSqlAuditEnableStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckSqlAuditEnableStatusResponse) SetHeaders(v map[string]*string) *CheckSqlAuditEnableStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckSqlAuditEnableStatusResponse) SetStatusCode(v int32) *CheckSqlAuditEnableStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSqlAuditEnableStatusResponse) SetBody(v *CheckSqlAuditEnableStatusResponseBody) *CheckSqlAuditEnableStatusResponse {
	s.Body = v
	return s
}

type CreateDrdsDBRequest struct {
	// The name of the account that has permissions to access all databases on the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required only when the Type parameter is set to VERTICAL.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The type of the storage instances that are used by the PolarDB-X 1.0 database. Set the value to RDS.
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// Specifies whether the required ApsaraDB RDS for MySQL instance is being created.
	DbInstanceIsCreating *bool `json:"DbInstanceIsCreating,omitempty" xml:"DbInstanceIsCreating,omitempty"`
	// The name of the PolarDB-X 1.0 database you want to create.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance on which you want to create the database.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The encoding method that is used by the database.
	Encode     *string                          `json:"Encode,omitempty" xml:"Encode,omitempty"`
	InstDbName []*CreateDrdsDBRequestInstDbName `json:"InstDbName,omitempty" xml:"InstDbName,omitempty" type:"Repeated"`
	// The password that is used to log on to the database.
	Password        *string                               `json:"Password,omitempty" xml:"Password,omitempty"`
	RdsInstance     []*string                             `json:"RdsInstance,omitempty" xml:"RdsInstance,omitempty" type:"Repeated"`
	RdsSuperAccount []*CreateDrdsDBRequestRdsSuperAccount `json:"RdsSuperAccount,omitempty" xml:"RdsSuperAccount,omitempty" type:"Repeated"`
	// The partitioning mode of the database. Valid values:
	//
	// *   **HORIZONTAL**: The database is horizontally partitioned (sharded).
	// *   **VERTICAL**: The database is vertically partitioned.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDrdsDBRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsDBRequest) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBRequest) SetAccountName(v string) *CreateDrdsDBRequest {
	s.AccountName = &v
	return s
}

func (s *CreateDrdsDBRequest) SetDbInstType(v string) *CreateDrdsDBRequest {
	s.DbInstType = &v
	return s
}

func (s *CreateDrdsDBRequest) SetDbInstanceIsCreating(v bool) *CreateDrdsDBRequest {
	s.DbInstanceIsCreating = &v
	return s
}

func (s *CreateDrdsDBRequest) SetDbName(v string) *CreateDrdsDBRequest {
	s.DbName = &v
	return s
}

func (s *CreateDrdsDBRequest) SetDrdsInstanceId(v string) *CreateDrdsDBRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateDrdsDBRequest) SetEncode(v string) *CreateDrdsDBRequest {
	s.Encode = &v
	return s
}

func (s *CreateDrdsDBRequest) SetInstDbName(v []*CreateDrdsDBRequestInstDbName) *CreateDrdsDBRequest {
	s.InstDbName = v
	return s
}

func (s *CreateDrdsDBRequest) SetPassword(v string) *CreateDrdsDBRequest {
	s.Password = &v
	return s
}

func (s *CreateDrdsDBRequest) SetRdsInstance(v []*string) *CreateDrdsDBRequest {
	s.RdsInstance = v
	return s
}

func (s *CreateDrdsDBRequest) SetRdsSuperAccount(v []*CreateDrdsDBRequestRdsSuperAccount) *CreateDrdsDBRequest {
	s.RdsSuperAccount = v
	return s
}

func (s *CreateDrdsDBRequest) SetType(v string) *CreateDrdsDBRequest {
	s.Type = &v
	return s
}

type CreateDrdsDBRequestInstDbName struct {
	// The ID of the ApsaraDB RDS for MySQL instance on which the databases need to be vertically partitioned. This parameter is required only when the Type parameter is set to VERTICAL.
	DbInstanceId *string   `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	ShardDbName  []*string `json:"ShardDbName,omitempty" xml:"ShardDbName,omitempty" type:"Repeated"`
}

func (s CreateDrdsDBRequestInstDbName) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsDBRequestInstDbName) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBRequestInstDbName) SetDbInstanceId(v string) *CreateDrdsDBRequestInstDbName {
	s.DbInstanceId = &v
	return s
}

func (s *CreateDrdsDBRequestInstDbName) SetShardDbName(v []*string) *CreateDrdsDBRequestInstDbName {
	s.ShardDbName = v
	return s
}

type CreateDrdsDBRequestRdsSuperAccount struct {
	// The account name of the super administrator that is used to connect to the ApsaraDB RDS for MySQL instance.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of ApsaraDB RDS instance.
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The password of the super administrator account that is used to connect to the ApsaraDB RDS instance.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s CreateDrdsDBRequestRdsSuperAccount) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsDBRequestRdsSuperAccount) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBRequestRdsSuperAccount) SetAccountName(v string) *CreateDrdsDBRequestRdsSuperAccount {
	s.AccountName = &v
	return s
}

func (s *CreateDrdsDBRequestRdsSuperAccount) SetDbInstanceId(v string) *CreateDrdsDBRequestRdsSuperAccount {
	s.DbInstanceId = &v
	return s
}

func (s *CreateDrdsDBRequestRdsSuperAccount) SetPassword(v string) *CreateDrdsDBRequestRdsSuperAccount {
	s.Password = &v
	return s
}

type CreateDrdsDBResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDrdsDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsDBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBResponseBody) SetRequestId(v string) *CreateDrdsDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDrdsDBResponseBody) SetSuccess(v bool) *CreateDrdsDBResponseBody {
	s.Success = &v
	return s
}

type CreateDrdsDBResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDrdsDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDrdsDBResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsDBResponse) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBResponse) SetHeaders(v map[string]*string) *CreateDrdsDBResponse {
	s.Headers = v
	return s
}

func (s *CreateDrdsDBResponse) SetStatusCode(v int32) *CreateDrdsDBResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDrdsDBResponse) SetBody(v *CreateDrdsDBResponseBody) *CreateDrdsDBResponse {
	s.Body = v
	return s
}

type CreateDrdsInstanceRequest struct {
	// Specifies the client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies the description of the instance. The description must meet the following requirements:
	//
	// *   The description cannot contain the prefix http:// or https://.
	// *   The description must start with a letter or a Chinese character, and can contain uppercase and lowercase letters, Chinese characters, digits, underscores (\_), and hyphens (-).
	// *   The description must be 2 to 256 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies the purchase duration of the subscription instance.
	//
	// *   If the PricingCycle parameter is set to year, the value range of the Duration parameter is 1 to 3.
	// *   If the PricingCycle parameter is set to month, the value range of the Duration parameter is 1 to 9.
	//
	// >  This parameter only takes effect when the PayType parameter is set to drdsPre.
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies the instance type of the instance. Valid values:
	//
	// *   **drds.sn2.4c16g**: The instance is of the Starter Edition.
	// *   **drds.sn2.8c32g**: The instance is of the Standard Edition
	// *   **drds.sn2.16c64g**: The instance is of the Enterprise Edition.
	InstanceSeries *string `json:"InstanceSeries,omitempty" xml:"InstanceSeries,omitempty"`
	// Specifies whether to enable automatic renewal. Valid values:
	//
	// *   **true**: If the PricingCycle parameter is set to month, the subscription is automatically renewed for one month. If the PricingCycle parameter is set to year, the subscription is automatically renewed for one year.
	// *   **false**: The auto-renewal feature is disabled for the instance.
	//
	// >  This parameter only takes effect when the PayType parameter is set to drdsPre.
	IsAutoRenew *bool `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	// Specifies the ID of the primary instance. This parameter is only required when you create a read-only instance.
	MasterInstId *string `json:"MasterInstId,omitempty" xml:"MasterInstId,omitempty"`
	// Specifies the MySQL version that is supported by the instance. Valid values:
	//
	// *   **5**: The instance is fully compatible with MySQL 5.x. This value is the default value.
	// *   **8**: The instance is fully compatible with MySQL 8.0.
	//
	// >  This parameter only takes effect when you create a primary instance. By default, the MySQL version of the read-only instance is the same as that of the primary instance.
	MySQLVersion *int32 `json:"MySQLVersion,omitempty" xml:"MySQLVersion,omitempty"`
	// Specifies the billing method of the instance. Valid values:
	//
	// *   **drdsPre**: The instance uses the subscription billing method.
	// *   **drdsPost**: The instance uses the pay-as-you-go billing method.
	// *   **drdsRo**: By default, the pay-as-you-go billing method is used when you create read-only instances.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Specifies the unit of the subscription duration of the subscription instance. Valid values:
	//
	// *   **year**: The unit of the subscription duration is year.
	// *   **month**: The unit of the subscription duration is month.
	//
	// >  This parameter is required if you set the PayType parameter to drdsPre.
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// Specifies the number of instances to be created. You can set the value only to 1. The value specifies that you can create one instance each time.
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// Specifies the region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies the ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies the specification code of the instance. The value consists of the instance type and the specified instance specification. For example, you can set the value to drds.sn2.4c16g.8c32g.
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// Specifies the type of the instance. Set the value to PRIVATE. The value PRIVATE specifies that the instance is a dedicated instance.
	//
	// >  You can also set the value to 1 to specify that the instance is a dedicated instance.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Specifies the ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Specifies the ID of the vSwitch.
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// Specifies the zone ID of the instance.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// Specifies whether the instance is a high-availability instance.
	IsHa *bool `json:"isHa,omitempty" xml:"isHa,omitempty"`
}

func (s CreateDrdsInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceRequest) SetClientToken(v string) *CreateDrdsInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetDescription(v string) *CreateDrdsInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetDuration(v int32) *CreateDrdsInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetInstanceSeries(v string) *CreateDrdsInstanceRequest {
	s.InstanceSeries = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetIsAutoRenew(v bool) *CreateDrdsInstanceRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetMasterInstId(v string) *CreateDrdsInstanceRequest {
	s.MasterInstId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetMySQLVersion(v int32) *CreateDrdsInstanceRequest {
	s.MySQLVersion = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetPayType(v string) *CreateDrdsInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetPricingCycle(v string) *CreateDrdsInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetQuantity(v int32) *CreateDrdsInstanceRequest {
	s.Quantity = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetRegionId(v string) *CreateDrdsInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetResourceGroupId(v string) *CreateDrdsInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetSpecification(v string) *CreateDrdsInstanceRequest {
	s.Specification = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetType(v string) *CreateDrdsInstanceRequest {
	s.Type = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetVpcId(v string) *CreateDrdsInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetVswitchId(v string) *CreateDrdsInstanceRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetZoneId(v string) *CreateDrdsInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetIsHa(v bool) *CreateDrdsInstanceRequest {
	s.IsHa = &v
	return s
}

type CreateDrdsInstanceResponseBody struct {
	// Indicates the details of the result.
	Data *CreateDrdsInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates the ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDrdsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponseBody) SetData(v *CreateDrdsInstanceResponseBodyData) *CreateDrdsInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateDrdsInstanceResponseBody) SetRequestId(v string) *CreateDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDrdsInstanceResponseBody) SetSuccess(v bool) *CreateDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

type CreateDrdsInstanceResponseBodyData struct {
	// Indicates the ID of the instance.
	DrdsInstanceIdList *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList `json:"DrdsInstanceIdList,omitempty" xml:"DrdsInstanceIdList,omitempty" type:"Struct"`
	// Indicates the ID of the order.
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDrdsInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponseBodyData) SetDrdsInstanceIdList(v *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) *CreateDrdsInstanceResponseBodyData {
	s.DrdsInstanceIdList = v
	return s
}

func (s *CreateDrdsInstanceResponseBodyData) SetOrderId(v int64) *CreateDrdsInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

type CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList struct {
	DrdsInstanceIdList []*string `json:"drdsInstanceIdList,omitempty" xml:"drdsInstanceIdList,omitempty" type:"Repeated"`
}

func (s CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) SetDrdsInstanceIdList(v []*string) *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList {
	s.DrdsInstanceIdList = v
	return s
}

type CreateDrdsInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDrdsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponse) SetHeaders(v map[string]*string) *CreateDrdsInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDrdsInstanceResponse) SetStatusCode(v int32) *CreateDrdsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDrdsInstanceResponse) SetBody(v *CreateDrdsInstanceResponseBody) *CreateDrdsInstanceResponse {
	s.Body = v
	return s
}

type CreateInstanceAccountRequest struct {
	// The username of the account you want to create.
	AccountName *string                                    `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbPrivilege []*CreateInstanceAccountRequestDbPrivilege `json:"DbPrivilege,omitempty" xml:"DbPrivilege,omitempty" type:"Repeated"`
	// The ID of the PolarDB-X 1.0 instance for which you want to create the account.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The password of the account you want to create.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s CreateInstanceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountRequest) SetAccountName(v string) *CreateInstanceAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateInstanceAccountRequest) SetDbPrivilege(v []*CreateInstanceAccountRequestDbPrivilege) *CreateInstanceAccountRequest {
	s.DbPrivilege = v
	return s
}

func (s *CreateInstanceAccountRequest) SetDrdsInstanceId(v string) *CreateInstanceAccountRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateInstanceAccountRequest) SetPassword(v string) *CreateInstanceAccountRequest {
	s.Password = &v
	return s
}

type CreateInstanceAccountRequestDbPrivilege struct {
	// The name of the database that you want to manage by using the account to create.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The permissions that you want to grant to the account to manage the database.
	Privilege *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
}

func (s CreateInstanceAccountRequestDbPrivilege) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAccountRequestDbPrivilege) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountRequestDbPrivilege) SetDbName(v string) *CreateInstanceAccountRequestDbPrivilege {
	s.DbName = &v
	return s
}

func (s *CreateInstanceAccountRequestDbPrivilege) SetPrivilege(v string) *CreateInstanceAccountRequestDbPrivilege {
	s.Privilege = &v
	return s
}

type CreateInstanceAccountResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountResponseBody) SetRequestId(v string) *CreateInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetSuccess(v bool) *CreateInstanceAccountResponseBody {
	s.Success = &v
	return s
}

type CreateInstanceAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountResponse) SetHeaders(v map[string]*string) *CreateInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceAccountResponse) SetStatusCode(v int32) *CreateInstanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceAccountResponse) SetBody(v *CreateInstanceAccountResponseBody) *CreateInstanceAccountResponse {
	s.Body = v
	return s
}

type CreateInstanceInternetAddressRequest struct {
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region to which the DRDS instance belongs.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateInstanceInternetAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceInternetAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceInternetAddressRequest) SetDrdsInstanceId(v string) *CreateInstanceInternetAddressRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateInstanceInternetAddressRequest) SetRegionId(v string) *CreateInstanceInternetAddressRequest {
	s.RegionId = &v
	return s
}

type CreateInstanceInternetAddressResponseBody struct {
	// The error code returned when the activity fails.
	//
	// >  This parameter appears only when an error occurs during the API call.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the public IP address was created.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceInternetAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceInternetAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceInternetAddressResponseBody) SetCode(v int32) *CreateInstanceInternetAddressResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceInternetAddressResponseBody) SetData(v bool) *CreateInstanceInternetAddressResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInstanceInternetAddressResponseBody) SetRequestId(v string) *CreateInstanceInternetAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceInternetAddressResponseBody) SetSuccess(v bool) *CreateInstanceInternetAddressResponseBody {
	s.Success = &v
	return s
}

type CreateInstanceInternetAddressResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceInternetAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceInternetAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceInternetAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceInternetAddressResponse) SetHeaders(v map[string]*string) *CreateInstanceInternetAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceInternetAddressResponse) SetStatusCode(v int32) *CreateInstanceInternetAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceInternetAddressResponse) SetBody(v *CreateInstanceInternetAddressResponseBody) *CreateInstanceInternetAddressResponse {
	s.Body = v
	return s
}

type CreateOrderForRdsRequest struct {
	// The JSON string that contains the order details. For more information, see [CreateDBInstance](~~26228~~).
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateOrderForRdsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderForRdsRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderForRdsRequest) SetParams(v string) *CreateOrderForRdsRequest {
	s.Params = &v
	return s
}

func (s *CreateOrderForRdsRequest) SetRegionId(v string) *CreateOrderForRdsRequest {
	s.RegionId = &v
	return s
}

type CreateOrderForRdsResponseBody struct {
	// The ID of the purchased RDS instance.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOrderForRdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderForRdsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderForRdsResponseBody) SetData(v string) *CreateOrderForRdsResponseBody {
	s.Data = &v
	return s
}

func (s *CreateOrderForRdsResponseBody) SetRequestId(v string) *CreateOrderForRdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderForRdsResponseBody) SetSuccess(v bool) *CreateOrderForRdsResponseBody {
	s.Success = &v
	return s
}

type CreateOrderForRdsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrderForRdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrderForRdsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderForRdsResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderForRdsResponse) SetHeaders(v map[string]*string) *CreateOrderForRdsResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderForRdsResponse) SetStatusCode(v int32) *CreateOrderForRdsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrderForRdsResponse) SetBody(v *CreateOrderForRdsResponseBody) *CreateOrderForRdsResponse {
	s.Body = v
	return s
}

type CreateShardTaskRequest struct {
	// The name of the DRDS database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region where the resource group resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the source table.
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	// The name of the destination table.
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	// The type of the task. Valid values:`  SHARD_TO_SINGLE `,`  SINGLE_TO_SHARD `,`  SHARD_TO_SHARD `.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateShardTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateShardTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateShardTaskRequest) SetDbName(v string) *CreateShardTaskRequest {
	s.DbName = &v
	return s
}

func (s *CreateShardTaskRequest) SetDrdsInstanceId(v string) *CreateShardTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateShardTaskRequest) SetRegionId(v string) *CreateShardTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateShardTaskRequest) SetSourceTableName(v string) *CreateShardTaskRequest {
	s.SourceTableName = &v
	return s
}

func (s *CreateShardTaskRequest) SetTargetTableName(v string) *CreateShardTaskRequest {
	s.TargetTableName = &v
	return s
}

func (s *CreateShardTaskRequest) SetTaskType(v string) *CreateShardTaskRequest {
	s.TaskType = &v
	return s
}

type CreateShardTaskResponseBody struct {
	// Task creation result
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the operation.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateShardTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateShardTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateShardTaskResponseBody) SetData(v bool) *CreateShardTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateShardTaskResponseBody) SetRequestId(v string) *CreateShardTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateShardTaskResponseBody) SetSuccess(v bool) *CreateShardTaskResponseBody {
	s.Success = &v
	return s
}

type CreateShardTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateShardTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateShardTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateShardTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateShardTaskResponse) SetHeaders(v map[string]*string) *CreateShardTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateShardTaskResponse) SetStatusCode(v int32) *CreateShardTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateShardTaskResponse) SetBody(v *CreateShardTaskResponseBody) *CreateShardTaskResponse {
	s.Body = v
	return s
}

type DescribeBackMenuRequest struct {
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeBackMenuRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackMenuRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuRequest) SetDrdsInstanceId(v string) *DescribeBackMenuRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeBackMenuResponseBody struct {
	// The backup information list.
	List *DescribeBackMenuResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of request.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackMenuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackMenuResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuResponseBody) SetList(v *DescribeBackMenuResponseBodyList) *DescribeBackMenuResponseBody {
	s.List = v
	return s
}

func (s *DescribeBackMenuResponseBody) SetRequestId(v string) *DescribeBackMenuResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackMenuResponseBody) SetSuccess(v bool) *DescribeBackMenuResponseBody {
	s.Success = &v
	return s
}

type DescribeBackMenuResponseBodyList struct {
	List []*DescribeBackMenuResponseBodyListList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s DescribeBackMenuResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackMenuResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuResponseBodyList) SetList(v []*DescribeBackMenuResponseBodyListList) *DescribeBackMenuResponseBodyList {
	s.List = v
	return s
}

type DescribeBackMenuResponseBodyListList struct {
	// The backup method. Valid values:
	//
	// *   **Logic **: logical backup
	// *   **phy**: physical backup
	MenuName *string `json:"MenuName,omitempty" xml:"MenuName,omitempty"`
	// Indicates whether backup recovery is supported.
	Support *bool `json:"Support,omitempty" xml:"Support,omitempty"`
}

func (s DescribeBackMenuResponseBodyListList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackMenuResponseBodyListList) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuResponseBodyListList) SetMenuName(v string) *DescribeBackMenuResponseBodyListList {
	s.MenuName = &v
	return s
}

func (s *DescribeBackMenuResponseBodyListList) SetSupport(v bool) *DescribeBackMenuResponseBodyListList {
	s.Support = &v
	return s
}

type DescribeBackMenuResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackMenuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackMenuResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackMenuResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuResponse) SetHeaders(v map[string]*string) *DescribeBackMenuResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackMenuResponse) SetStatusCode(v int32) *DescribeBackMenuResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackMenuResponse) SetBody(v *DescribeBackMenuResponseBody) *DescribeBackMenuResponse {
	s.Body = v
	return s
}

type DescribeBackupDbsRequest struct {
	// Query by backup set ID
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of a DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// Query by restoration time.
	PreferredRestoreTime *string `json:"PreferredRestoreTime,omitempty" xml:"PreferredRestoreTime,omitempty"`
}

func (s DescribeBackupDbsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupDbsRequest) SetBackupId(v string) *DescribeBackupDbsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupDbsRequest) SetDrdsInstanceId(v string) *DescribeBackupDbsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBackupDbsRequest) SetPreferredRestoreTime(v string) *DescribeBackupDbsRequest {
	s.PreferredRestoreTime = &v
	return s
}

type DescribeBackupDbsResponseBody struct {
	// The details about a database.
	DbNames *DescribeBackupDbsResponseBodyDbNames `json:"DbNames,omitempty" xml:"DbNames,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of request.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupDbsResponseBody) SetDbNames(v *DescribeBackupDbsResponseBodyDbNames) *DescribeBackupDbsResponseBody {
	s.DbNames = v
	return s
}

func (s *DescribeBackupDbsResponseBody) SetRequestId(v string) *DescribeBackupDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupDbsResponseBody) SetSuccess(v bool) *DescribeBackupDbsResponseBody {
	s.Success = &v
	return s
}

type DescribeBackupDbsResponseBodyDbNames struct {
	DbName []*string `json:"dbName,omitempty" xml:"dbName,omitempty" type:"Repeated"`
}

func (s DescribeBackupDbsResponseBodyDbNames) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDbsResponseBodyDbNames) GoString() string {
	return s.String()
}

func (s *DescribeBackupDbsResponseBodyDbNames) SetDbName(v []*string) *DescribeBackupDbsResponseBodyDbNames {
	s.DbName = v
	return s
}

type DescribeBackupDbsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupDbsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupDbsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDbsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupDbsResponse) SetHeaders(v map[string]*string) *DescribeBackupDbsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupDbsResponse) SetStatusCode(v int32) *DescribeBackupDbsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupDbsResponse) SetBody(v *DescribeBackupDbsResponseBody) *DescribeBackupDbsResponse {
	s.Body = v
	return s
}

type DescribeBackupLocalRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeBackupLocalRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLocalRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupLocalRequest) SetDrdsInstanceId(v string) *DescribeBackupLocalRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeBackupLocalResponseBody struct {
	// The information about the backup policy.
	BackupPolicyDO *DescribeBackupLocalResponseBodyBackupPolicyDO `json:"BackupPolicyDO,omitempty" xml:"BackupPolicyDO,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupLocalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLocalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupLocalResponseBody) SetBackupPolicyDO(v *DescribeBackupLocalResponseBodyBackupPolicyDO) *DescribeBackupLocalResponseBody {
	s.BackupPolicyDO = v
	return s
}

func (s *DescribeBackupLocalResponseBody) SetRequestId(v string) *DescribeBackupLocalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupLocalResponseBody) SetSuccess(v bool) *DescribeBackupLocalResponseBody {
	s.Success = &v
	return s
}

type DescribeBackupLocalResponseBodyBackupPolicyDO struct {
	// No value is returned.
	BackupAppName *string `json:"BackupAppName,omitempty" xml:"BackupAppName,omitempty"`
	// No value is returned.
	BackupDbName *string `json:"BackupDbName,omitempty" xml:"BackupDbName,omitempty"`
	// No value is returned.
	BackupLevel *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	// No value is returned.
	BackupLog *string `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	// No value is returned.
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// No value is returned.
	BackupPolicyMode *string `json:"BackupPolicyMode,omitempty" xml:"BackupPolicyMode,omitempty"`
	// No value is returned.
	BackupRetentionPeriod *int64 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// No value is returned.
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// No value is returned.
	DataBackupRetentionPeriod *int64 `json:"DataBackupRetentionPeriod,omitempty" xml:"DataBackupRetentionPeriod,omitempty"`
	// No value is returned.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// No value is returned.
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the feature is enabled to forcibly delete binary log files if the used storage space of the instance exceeds 90% of the total storage space or the remaining storage space is less than 5 GB. Valid values:
	//
	// *   1: The feature is enabled.
	// *   0: The feature is disabled.
	HighSpaceUsageProtection *int64 `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	// The number of hours for which log backup files are retained on the instance. Valid values: 0 to 168. Default value: **18**. The value **0** indicates that log backup files are not retained.
	LocalLogRetentionHours *int64 `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	// The maximum storage usage that is allowed for local log files. Valid values: 0 to 50. Default value: 30.
	LocalLogRetentionSpace *int64 `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	// No value is returned.
	LogBackupRetentionPeriod *int64 `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	// No value is returned.
	NextBackupActuallyTime *string `json:"NextBackupActuallyTime,omitempty" xml:"NextBackupActuallyTime,omitempty"`
	// No value is returned.
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// No value is returned.
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
}

func (s DescribeBackupLocalResponseBodyBackupPolicyDO) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLocalResponseBodyBackupPolicyDO) GoString() string {
	return s.String()
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupAppName(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupAppName = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupDbName(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupDbName = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupLevel(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupLevel = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupLog(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupLog = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupMode(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupPolicyMode(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupPolicyMode = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupRetentionPeriod(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupType(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetDataBackupRetentionPeriod(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.DataBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetGmtCreate(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.GmtCreate = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetGmtModified(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.GmtModified = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetHighSpaceUsageProtection(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetLocalLogRetentionHours(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetLocalLogRetentionSpace(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetLogBackupRetentionPeriod(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetNextBackupActuallyTime(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.NextBackupActuallyTime = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetPreferredBackupPeriod(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetPreferredBackupTime(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.PreferredBackupTime = &v
	return s
}

type DescribeBackupLocalResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupLocalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupLocalResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLocalResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupLocalResponse) SetHeaders(v map[string]*string) *DescribeBackupLocalResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupLocalResponse) SetStatusCode(v int32) *DescribeBackupLocalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupLocalResponse) SetBody(v *DescribeBackupLocalResponseBody) *DescribeBackupLocalResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	// The ID of the instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetDrdsInstanceId(v string) *DescribeBackupPolicyRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	// The information about the backup policy.
	BackupPolicyDO *DescribeBackupPolicyResponseBodyBackupPolicyDO `json:"BackupPolicyDO,omitempty" xml:"BackupPolicyDO,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetBackupPolicyDO(v *DescribeBackupPolicyResponseBodyBackupPolicyDO) *DescribeBackupPolicyResponseBody {
	s.BackupPolicyDO = v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetSuccess(v bool) *DescribeBackupPolicyResponseBody {
	s.Success = &v
	return s
}

type DescribeBackupPolicyResponseBodyBackupPolicyDO struct {
	// No value is returned.
	BackupAppName *string `json:"BackupAppName,omitempty" xml:"BackupAppName,omitempty"`
	// No value is returned.
	BackupDbName *string `json:"BackupDbName,omitempty" xml:"BackupDbName,omitempty"`
	// The backup level. Valid values:
	//
	// *   **db**: database backup
	// *   **instance**: instance backup
	BackupLevel *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	// Indicates whether the log backup feature is enabled. Valid values:
	//
	// *   **1**: The log backup feature is enabled.
	// *   **0**: The log backup feature is disabled.
	BackupLog *string `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	// The backup mode. Valid values:
	//
	// *   **logic**: logical backup
	// *   **phy**: fast backup
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The type of the backup policy. Valid values:
	//
	// *   **DataBackupPolicy**: a data backup policy
	// *   **LogBackupPolicy**: a log backup policy
	BackupPolicyMode *string `json:"BackupPolicyMode,omitempty" xml:"BackupPolicyMode,omitempty"`
	// The retention period of backup files. Unit: days.
	BackupRetentionPeriod *int64 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// No value is returned.
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The retention period of data backup files. Unit: days.
	DataBackupRetentionPeriod *int64 `json:"DataBackupRetentionPeriod,omitempty" xml:"DataBackupRetentionPeriod,omitempty"`
	// No value is returned.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// No value is returned.
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// No value is returned.
	HighSpaceUsageProtection *int64 `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	// No value is returned.
	LocalLogRetentionHours *int64 `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	// No value is returned.
	LocalLogRetentionSpace *int64 `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	// The retention period of log backup files. Unit: days.
	LogBackupRetentionPeriod *int64 `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	// No value is returned.
	NextBackupActuallyTime *string `json:"NextBackupActuallyTime,omitempty" xml:"NextBackupActuallyTime,omitempty"`
	// The backup cycle. You can specify multiple backup cycles. Separate multiple backup cycles with commas (,). Valid values:
	//
	// *   **0**: every Monday
	// *   **1**: every Tuesday
	// *   **2**: every Wednesday
	// *   **3**: every Thursday
	// *   **4**: every Friday
	// *   **5**: every Saturday
	// *   **6**: every Sunday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The time range in which a backup is performed. The time is displayed in UTC.
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyBackupPolicyDO) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyBackupPolicyDO) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupAppName(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupAppName = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupDbName(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupDbName = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupLevel(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupLevel = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupLog(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupLog = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupMode(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupPolicyMode(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupPolicyMode = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupRetentionPeriod(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupType(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetDataBackupRetentionPeriod(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.DataBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetGmtCreate(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.GmtCreate = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetGmtModified(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.GmtModified = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetHighSpaceUsageProtection(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetLocalLogRetentionHours(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetLocalLogRetentionSpace(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetLogBackupRetentionPeriod(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetNextBackupActuallyTime(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.NextBackupActuallyTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.PreferredBackupTime = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeBackupSetsRequest struct {
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The end of the query time which is in timestamp format (measured in millisecond) .
	//
	// >  The end time must be later than the start time.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the query time which is in timestamp format (measured in millisecond).
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsRequest) SetDrdsInstanceId(v string) *DescribeBackupSetsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBackupSetsRequest) SetEndTime(v string) *DescribeBackupSetsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupSetsRequest) SetStartTime(v string) *DescribeBackupSetsRequest {
	s.StartTime = &v
	return s
}

type DescribeBackupSetsResponseBody struct {
	// The list of backup sets.
	BackupSets *DescribeBackupSetsResponseBodyBackupSets `json:"BackupSets,omitempty" xml:"BackupSets,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponseBody) SetBackupSets(v *DescribeBackupSetsResponseBodyBackupSets) *DescribeBackupSetsResponseBody {
	s.BackupSets = v
	return s
}

func (s *DescribeBackupSetsResponseBody) SetRequestId(v string) *DescribeBackupSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupSetsResponseBody) SetSuccess(v bool) *DescribeBackupSetsResponseBody {
	s.Success = &v
	return s
}

type DescribeBackupSetsResponseBodyBackupSets struct {
	BackupSet []*DescribeBackupSetsResponseBodyBackupSetsBackupSet `json:"backupSet,omitempty" xml:"backupSet,omitempty" type:"Repeated"`
}

func (s DescribeBackupSetsResponseBodyBackupSets) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetsResponseBodyBackupSets) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponseBodyBackupSets) SetBackupSet(v []*DescribeBackupSetsResponseBodyBackupSetsBackupSet) *DescribeBackupSetsResponseBodyBackupSets {
	s.BackupSet = v
	return s
}

type DescribeBackupSetsResponseBodyBackupSetsBackupSet struct {
	// Backup Recovery duration.
	BackupConsitentTime *string `json:"BackupConsitentTime,omitempty" xml:"BackupConsitentTime,omitempty"`
	// The list of backup databases.
	BackupDbs *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs `json:"BackupDbs,omitempty" xml:"BackupDbs,omitempty" type:"Struct"`
	// The end of the backup time which is in timestamp format (measured in millisecond).
	//
	// >  0 indicates not finished.
	BackupEndTime *int64 `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The level of the backup. Valid values:
	//
	// *   db: The database level.
	// *   instance: the instance level.
	BackupLevel *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	// The backup method. Valid values:
	//
	// *   logic: the logical backup.
	// *   phy: fast backup
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The beginning of the backup time which is in timestamp format (measured in millisecond).
	BackupStartTime *int64 `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The size of the backup set. Unit: MB.
	BackupTotalSize *string `json:"BackupTotalSize,omitempty" xml:"BackupTotalSize,omitempty"`
	// The type of the backup. Valid values:
	//
	// *   manual: indicates a manual backup.
	// *   auto: indicates an automatic backup.
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// Indicates whether the backup set can be restored. Valid values:
	EnableRecovery *bool `json:"EnableRecovery,omitempty" xml:"EnableRecovery,omitempty"`
	// The ID of the data backup file you want to use.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status of the backup instance. Valid values:
	//
	// *   \-1: Failed
	// *   0: Not started
	// *   1: The storage instance is running.
	// *   2: Success
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBackupSetsResponseBodyBackupSetsBackupSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetsResponseBodyBackupSetsBackupSet) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupConsitentTime(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupConsitentTime = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupDbs(v *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupDbs = v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupEndTime(v int64) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupLevel(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupLevel = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupMode(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupStartTime(v int64) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupTotalSize(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupTotalSize = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupType(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetEnableRecovery(v bool) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.EnableRecovery = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetId(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.Id = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetStatus(v int64) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.Status = &v
	return s
}

type DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs struct {
	BackupDb []*string `json:"backupDb,omitempty" xml:"backupDb,omitempty" type:"Repeated"`
}

func (s DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) SetBackupDb(v []*string) *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs {
	s.BackupDb = v
	return s
}

type DescribeBackupSetsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupSetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponse) SetHeaders(v map[string]*string) *DescribeBackupSetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupSetsResponse) SetStatusCode(v int32) *DescribeBackupSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupSetsResponse) SetBody(v *DescribeBackupSetsResponseBody) *DescribeBackupSetsResponse {
	s.Body = v
	return s
}

type DescribeBackupTimesRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeBackupTimesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTimesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupTimesRequest) SetDrdsInstanceId(v string) *DescribeBackupTimesRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeBackupTimesResponseBody struct {
	// Indicates the ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates the information about the time range within which the data of the instance can be restored to a point in time.
	RestoreTime *DescribeBackupTimesResponseBodyRestoreTime `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupTimesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTimesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupTimesResponseBody) SetRequestId(v string) *DescribeBackupTimesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupTimesResponseBody) SetRestoreTime(v *DescribeBackupTimesResponseBodyRestoreTime) *DescribeBackupTimesResponseBody {
	s.RestoreTime = v
	return s
}

func (s *DescribeBackupTimesResponseBody) SetSuccess(v bool) *DescribeBackupTimesResponseBody {
	s.Success = &v
	return s
}

type DescribeBackupTimesResponseBodyRestoreTime struct {
	// Indicates the end time. The time is in the UNIX timestamp format. The time is in UTC. Unit: ms.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates the start time. The time is in the UNIX timestamp format. The time must be in UTC. Unit: ms.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupTimesResponseBodyRestoreTime) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTimesResponseBodyRestoreTime) GoString() string {
	return s.String()
}

func (s *DescribeBackupTimesResponseBodyRestoreTime) SetEndTime(v string) *DescribeBackupTimesResponseBodyRestoreTime {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupTimesResponseBodyRestoreTime) SetStartTime(v string) *DescribeBackupTimesResponseBodyRestoreTime {
	s.StartTime = &v
	return s
}

type DescribeBackupTimesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupTimesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupTimesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTimesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupTimesResponse) SetHeaders(v map[string]*string) *DescribeBackupTimesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupTimesResponse) SetStatusCode(v int32) *DescribeBackupTimesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupTimesResponse) SetBody(v *DescribeBackupTimesResponseBody) *DescribeBackupTimesResponse {
	s.Body = v
	return s
}

type DescribeBroadcastTablesRequest struct {
	// The number of the page to return.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The content of the query.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBroadcastTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBroadcastTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesRequest) SetCurrentPage(v int32) *DescribeBroadcastTablesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetDbName(v string) *DescribeBroadcastTablesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetDrdsInstanceId(v string) *DescribeBroadcastTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetPageSize(v int32) *DescribeBroadcastTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetQuery(v string) *DescribeBroadcastTablesRequest {
	s.Query = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetRegionId(v string) *DescribeBroadcastTablesRequest {
	s.RegionId = &v
	return s
}

type DescribeBroadcastTablesResponseBody struct {
	// Indicates whether the database is sharded.
	IsShard *bool `json:"IsShard,omitempty" xml:"IsShard,omitempty"`
	// Indicates information about broadcast tables.
	List []*DescribeBroadcastTablesResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Indicates the page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Indicates the number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates the ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates the total number of entries returned.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeBroadcastTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBroadcastTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesResponseBody) SetIsShard(v bool) *DescribeBroadcastTablesResponseBody {
	s.IsShard = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetList(v []*DescribeBroadcastTablesResponseBodyList) *DescribeBroadcastTablesResponseBody {
	s.List = v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetPageNumber(v int32) *DescribeBroadcastTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetPageSize(v int32) *DescribeBroadcastTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetRequestId(v string) *DescribeBroadcastTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetSuccess(v bool) *DescribeBroadcastTablesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetTotal(v int32) *DescribeBroadcastTablesResponseBody {
	s.Total = &v
	return s
}

type DescribeBroadcastTablesResponseBodyList struct {
	// Indicates whether a table is a broadcast table.
	Broadcast *bool `json:"Broadcast,omitempty" xml:"Broadcast,omitempty"`
	// Indicates the type of the broadcast table. Valid values:
	//
	// *   **1**: multi-write mode
	// *   **2**: synchronous mode
	BroadcastType *string `json:"BroadcastType,omitempty" xml:"BroadcastType,omitempty"`
	// Indicates the storage type of the database. Valid values:
	//
	// *   **0**: RDS
	// *   **4**: PolarDB
	DbInstType *int32 `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// Indicates whether the broadcast table was sharded.
	IsShard *bool `json:"IsShard,omitempty" xml:"IsShard,omitempty"`
	// Indicates the activation state of the broadcast table. Valid values:
	//
	// *   **1**: The broadcast table is activated.
	// *   **2**: The broadcast table is being activated.
	// *   **3**: An exception occurs when the broadcast table is being activated.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates the name of the table.
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeBroadcastTablesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBroadcastTablesResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesResponseBodyList) SetBroadcast(v bool) *DescribeBroadcastTablesResponseBodyList {
	s.Broadcast = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetBroadcastType(v string) *DescribeBroadcastTablesResponseBodyList {
	s.BroadcastType = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetDbInstType(v int32) *DescribeBroadcastTablesResponseBodyList {
	s.DbInstType = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetIsShard(v bool) *DescribeBroadcastTablesResponseBodyList {
	s.IsShard = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetStatus(v int32) *DescribeBroadcastTablesResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetTable(v string) *DescribeBroadcastTablesResponseBodyList {
	s.Table = &v
	return s
}

type DescribeBroadcastTablesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBroadcastTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBroadcastTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBroadcastTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesResponse) SetHeaders(v map[string]*string) *DescribeBroadcastTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBroadcastTablesResponse) SetStatusCode(v int32) *DescribeBroadcastTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBroadcastTablesResponse) SetBody(v *DescribeBroadcastTablesResponseBody) *DescribeBroadcastTablesResponse {
	s.Body = v
	return s
}

type DescribeDbInstanceDbsRequest struct {
	// The name of the privileged account of the PolarDB-X 1.0 instance. You do not need to specify this parameter if you have no privileged account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The engine type of the storage-layer databases. Valid values: **POLARDB** and **RDS**.
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The ID of the instance in which the storage-layer databases are deployed.
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The password of the privileged account. You do not need to specify this parameter if you have no privileged account.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s DescribeDbInstanceDbsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstanceDbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsRequest) SetAccountName(v string) *DescribeDbInstanceDbsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetDbInstType(v string) *DescribeDbInstanceDbsRequest {
	s.DbInstType = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetDbInstanceId(v string) *DescribeDbInstanceDbsRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetDrdsInstanceId(v string) *DescribeDbInstanceDbsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetPassword(v string) *DescribeDbInstanceDbsRequest {
	s.Password = &v
	return s
}

type DescribeDbInstanceDbsResponseBody struct {
	// Indicates the information about the storage-layer databases.
	Databases *DescribeDbInstanceDbsResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates the total number of storage-layer databases.
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDbInstanceDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstanceDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsResponseBody) SetDatabases(v *DescribeDbInstanceDbsResponseBodyDatabases) *DescribeDbInstanceDbsResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDbInstanceDbsResponseBody) SetRequestId(v string) *DescribeDbInstanceDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBody) SetSuccess(v bool) *DescribeDbInstanceDbsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBody) SetTotal(v string) *DescribeDbInstanceDbsResponseBody {
	s.Total = &v
	return s
}

type DescribeDbInstanceDbsResponseBodyDatabases struct {
	Database []*DescribeDbInstanceDbsResponseBodyDatabasesDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Repeated"`
}

func (s DescribeDbInstanceDbsResponseBodyDatabases) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstanceDbsResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsResponseBodyDatabases) SetDatabase(v []*DescribeDbInstanceDbsResponseBodyDatabasesDatabase) *DescribeDbInstanceDbsResponseBodyDatabases {
	s.Database = v
	return s
}

type DescribeDbInstanceDbsResponseBodyDatabasesDatabase struct {
	// Indicates the name of a storage-layer database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates the description of the storage-layer database.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates the state of the storage-layer database. Valid values:
	//
	// *   **0**: The database is being created.
	// *   **1**: The database is available.
	// *   **3**: The database is being deleted.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDbInstanceDbsResponseBodyDatabasesDatabase) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstanceDbsResponseBodyDatabasesDatabase) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) SetDbName(v string) *DescribeDbInstanceDbsResponseBodyDatabasesDatabase {
	s.DbName = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) SetDescription(v string) *DescribeDbInstanceDbsResponseBodyDatabasesDatabase {
	s.Description = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) SetStatus(v int32) *DescribeDbInstanceDbsResponseBodyDatabasesDatabase {
	s.Status = &v
	return s
}

type DescribeDbInstanceDbsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDbInstanceDbsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDbInstanceDbsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstanceDbsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsResponse) SetHeaders(v map[string]*string) *DescribeDbInstanceDbsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDbInstanceDbsResponse) SetStatusCode(v int32) *DescribeDbInstanceDbsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDbInstanceDbsResponse) SetBody(v *DescribeDbInstanceDbsResponseBody) *DescribeDbInstanceDbsResponse {
	s.Body = v
	return s
}

type DescribeDbInstancesRequest struct {
	// Storage layer type. Valid values: **POLARDB** or **RDS**.
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The ID of a DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the storage or cluster.
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
}

func (s DescribeDbInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesRequest) SetDbInstType(v string) *DescribeDbInstancesRequest {
	s.DbInstType = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetDrdsInstanceId(v string) *DescribeDbInstancesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetPageNumber(v int32) *DescribeDbInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetPageSize(v int32) *DescribeDbInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetRegionId(v string) *DescribeDbInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetSearch(v string) *DescribeDbInstancesRequest {
	s.Search = &v
	return s
}

type DescribeDbInstancesResponseBody struct {
	// The details of the instance.
	Items *DescribeDbInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDbInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponseBody) SetItems(v *DescribeDbInstancesResponseBodyItems) *DescribeDbInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDbInstancesResponseBody) SetRequestId(v string) *DescribeDbInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDbInstancesResponseBodyItems struct {
	DBInstance []*DescribeDbInstancesResponseBodyItemsDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeDbInstancesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponseBodyItems) SetDBInstance(v []*DescribeDbInstancesResponseBodyItemsDBInstance) *DescribeDbInstancesResponseBodyItems {
	s.DBInstance = v
	return s
}

type DescribeDbInstancesResponseBodyItemsDBInstance struct {
	AllowAllCategory *bool `json:"AllowAllCategory,omitempty" xml:"AllowAllCategory,omitempty"`
	// The description of the storage instance.
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The ID of the storage instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Storage layer instance status. Valid values:
	//
	// *   **0**: creating
	// *   **1**: In use
	// *   **3**: Deleting
	// *   **5**: restarting
	// *   **6**: upgrading /Downgrading
	// *   **7**: Recovering
	// *   **8**: switching the Internet and intranet
	DBInstanceStatus *int32 `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The storage layer instance type.
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The engine of the storage instance.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the engine for the storage instance.
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The network type of the storage layer. Valid values:
	//
	// *   **VPC**: VPC
	// *   **CLASSIC **: Classic Network
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The details about a read-only storage instance.
	ReadOnlyDBInstanceId *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Struct"`
	// The ID of the region where the storage instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the zone where the storage instance resides.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDbInstancesResponseBodyItemsDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesResponseBodyItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetAllowAllCategory(v bool) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.AllowAllCategory = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceDescription(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceId(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceStatus(v int32) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceType(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetEngine(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetEngineVersion(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetInstanceNetworkType(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetReadOnlyDBInstanceId(v *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.ReadOnlyDBInstanceId = v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetRegionId(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetZoneId(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.ZoneId = &v
	return s
}

type DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId struct {
	ReadOnlyDBInstanceId []*string `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) SetReadOnlyDBInstanceId(v []*string) *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId {
	s.ReadOnlyDBInstanceId = v
	return s
}

type DescribeDbInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDbInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDbInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponse) SetHeaders(v map[string]*string) *DescribeDbInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDbInstancesResponse) SetStatusCode(v int32) *DescribeDbInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDbInstancesResponse) SetBody(v *DescribeDbInstancesResponseBody) *DescribeDbInstancesResponse {
	s.Body = v
	return s
}

type DescribeDrdsDBRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsDBRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBRequest) SetDbName(v string) *DescribeDrdsDBRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsDBResponseBody struct {
	// Indicates the details about the database.
	Data *DescribeDrdsDBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates the ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBResponseBody) SetData(v *DescribeDrdsDBResponseBodyData) *DescribeDrdsDBResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsDBResponseBody) SetRequestId(v string) *DescribeDrdsDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBResponseBody) SetSuccess(v bool) *DescribeDrdsDBResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsDBResponseBodyData struct {
	// Indicates the time when the database was created. The value is in the UNIX timestamp format. Unit: ms.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates the storage type of the database.
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// Indicates the name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates the type of the instance in which the database is deployed. Valid values:
	//
	// *   **MASTER**: The instance is a primary instance.
	// *   **SLAVE**: The instance is a read-only instance.
	InstRole *string `json:"InstRole,omitempty" xml:"InstRole,omitempty"`
	// Indicates the database sharding method.
	//
	// *   **HORIZONTAL**: The database is horizontally sharded.
	// *   **VERTICAL**: The database is vertically sharded.
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// Indicates the schema name of the database.
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// Indicates the state of the database. Valid values:
	//
	// *   **TO_BE_INIT**: The database is being created.
	// *   **NORMAL**: The database is running.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDrdsDBResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBResponseBodyData) SetCreateTime(v string) *DescribeDrdsDBResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetDbInstType(v string) *DescribeDrdsDBResponseBodyData {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetDbName(v string) *DescribeDrdsDBResponseBodyData {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetInstRole(v string) *DescribeDrdsDBResponseBodyData {
	s.InstRole = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetMode(v string) *DescribeDrdsDBResponseBodyData {
	s.Mode = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetSchema(v string) *DescribeDrdsDBResponseBodyData {
	s.Schema = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetStatus(v string) *DescribeDrdsDBResponseBodyData {
	s.Status = &v
	return s
}

type DescribeDrdsDBResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsDBResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBResponse) SetHeaders(v map[string]*string) *DescribeDrdsDBResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDBResponse) SetStatusCode(v int32) *DescribeDrdsDBResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDBResponse) SetBody(v *DescribeDrdsDBResponseBody) *DescribeDrdsDBResponse {
	s.Body = v
	return s
}

type DescribeDrdsDBClusterRequest struct {
	// The ID of the PolarDB cluster.
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The name of the DRDS database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of a DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterRequest) SetDbInstanceId(v string) *DescribeDrdsDBClusterRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDrdsDBClusterRequest) SetDbName(v string) *DescribeDrdsDBClusterRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBClusterRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBClusterRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsDBClusterResponseBody struct {
	// The details of each PolarDB cluster.
	DbInstance *DescribeDrdsDBClusterResponseBodyDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBody) SetDbInstance(v *DescribeDrdsDBClusterResponseBodyDbInstance) *DescribeDrdsDBClusterResponseBody {
	s.DbInstance = v
	return s
}

func (s *DescribeDrdsDBClusterResponseBody) SetRequestId(v string) *DescribeDrdsDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBody) SetSuccess(v bool) *DescribeDrdsDBClusterResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsDBClusterResponseBodyDbInstance struct {
	// The ID of the PolarDB cluster.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The status of the PolarDB instance.
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The information about the nodes in the PolarDB Cluster.
	DBNodes *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Struct"`
	// The type of storage used by the DRDS database.
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The endpoint of the PolarDB read /write splitting endpoint
	Endpoints *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	// The type of the DRDS database storage engine.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the DRDS database storage engine.
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the PolarDB cluster expires.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The network type of the PolarDB cluster.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the PolarDB cluster.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The PolarDB access port.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of RDS instance. PolarDB cluster does not support this parameter.
	RdsInstType *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	// This parameter specifies the Read mode when the database storage type is PolarDB.
	//
	// Valid values:
	//
	// *   **DEFAULT**: the default mode (that is, all read traffic is sent to the PolarDB read /write node).
	// *   **CUSTOM**: Custom mode (you can customize the ratio of traffic sent to read /write nodes and read-only nodes).
	// *   **BALANCE**: read balancing mode (the read traffic is automatically distributed by the read load module of the PolarDB cluster, which can also be understood as the read traffic being evenly distributed to each node).
	ReadMode *string `json:"ReadMode,omitempty" xml:"ReadMode,omitempty"`
	// The number of days remaining on the PolarDB for MySQL instance.
	RemainDays *string `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDBInstanceId(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDBNodes(v *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DBNodes = v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDbInstType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetEndpoints(v *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.Endpoints = v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetEngine(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetEngineVersion(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetExpireTime(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetNetworkType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetPayType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetPort(v int32) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetRdsInstType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetReadMode(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.ReadMode = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetRemainDays(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.RemainDays = &v
	return s
}

type DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes struct {
	DBNode []*DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) SetDBNode(v []*DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes {
	s.DBNode = v
	return s
}

type DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode struct {
	// The ID of the node in the apsaradb for PolarDB cluster.
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The role of a node in the apsaradb for PolarDB cluster. Valid values:
	//
	// *   **Reader**
	// *   **Writer**
	DBNodeRole *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	// The status of the nodes in the PolarDB cluster.
	DBNodeStatus *string `json:"DBNodeStatus,omitempty" xml:"DBNodeStatus,omitempty"`
	// The ID of the zone where the node of the PolarDB cluster resides.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetDBNodeId(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetDBNodeRole(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetDBNodeStatus(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.DBNodeStatus = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetZoneId(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.ZoneId = &v
	return s
}

type DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints struct {
	Endpoint []*DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) SetEndpoint(v []*DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints {
	s.Endpoint = v
	return s
}

type DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint struct {
	// The ID of the PolarDB connection address.
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID list of the nodes in the PolarDB connection string. Separate multiple nodes with commas (,).
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The read ratio of this connection address managed by the DRDS database.
	ReadWeight *int32 `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) SetEndpointId(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint {
	s.EndpointId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) SetNodeIds(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint {
	s.NodeIds = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) SetReadWeight(v int32) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint {
	s.ReadWeight = &v
	return s
}

type DescribeDrdsDBClusterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponse) SetHeaders(v map[string]*string) *DescribeDrdsDBClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDBClusterResponse) SetStatusCode(v int32) *DescribeDrdsDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDBClusterResponse) SetBody(v *DescribeDrdsDBClusterResponseBody) *DescribeDrdsDBClusterResponse {
	s.Body = v
	return s
}

type DescribeDrdsDBIpWhiteListRequest struct {
	// The database name.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The instance ID.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The name of the whitelist group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsDBIpWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetDbName(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetGroupName(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetRegionId(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.RegionId = &v
	return s
}

type DescribeDrdsDBIpWhiteListResponseBody struct {
	// The IP address whitelist.
	IpWhiteList *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList `json:"IpWhiteList,omitempty" xml:"IpWhiteList,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDBIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) SetIpWhiteList(v *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) *DescribeDrdsDBIpWhiteListResponseBody {
	s.IpWhiteList = v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) SetRequestId(v string) *DescribeDrdsDBIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) SetSuccess(v bool) *DescribeDrdsDBIpWhiteListResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList struct {
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) SetIp(v []*string) *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList {
	s.Ip = v
	return s
}

type DescribeDrdsDBIpWhiteListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsDBIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsDBIpWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListResponse) SetHeaders(v map[string]*string) *DescribeDrdsDBIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponse) SetStatusCode(v int32) *DescribeDrdsDBIpWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponse) SetBody(v *DescribeDrdsDBIpWhiteListResponseBody) *DescribeDrdsDBIpWhiteListResponse {
	s.Body = v
	return s
}

type DescribeDrdsDBsRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The number of the page to return. The value of this parameter must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of databases to return on each page. Valid values: **30**, **50**, and **100**.
	//
	// Default value: **30**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the PolarDB-X 1.0 instance is created.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsDBsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDBsRequest) SetPageNumber(v int32) *DescribeDrdsDBsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsDBsRequest) SetPageSize(v int32) *DescribeDrdsDBsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsDBsRequest) SetRegionId(v string) *DescribeDrdsDBsRequest {
	s.RegionId = &v
	return s
}

type DescribeDrdsDBsResponseBody struct {
	// The list of returned databases.
	Data *DescribeDrdsDBsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of databases returned on each page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of returned databases.
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsDBsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsResponseBody) SetData(v *DescribeDrdsDBsResponseBodyData) *DescribeDrdsDBsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetPageNumber(v string) *DescribeDrdsDBsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetPageSize(v string) *DescribeDrdsDBsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetRequestId(v string) *DescribeDrdsDBsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetSuccess(v bool) *DescribeDrdsDBsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetTotal(v string) *DescribeDrdsDBsResponseBody {
	s.Total = &v
	return s
}

type DescribeDrdsDBsResponseBodyData struct {
	Db []*DescribeDrdsDBsResponseBodyDataDb `json:"Db,omitempty" xml:"Db,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDBsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsResponseBodyData) SetDb(v []*DescribeDrdsDBsResponseBodyDataDb) *DescribeDrdsDBsResponseBodyData {
	s.Db = v
	return s
}

type DescribeDrdsDBsResponseBodyDataDb struct {
	// The time when the database is created. The value of this parameter is a UNIX timestamp. Unit: ms.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of the database. Valid values: **RDS** and **POLARDB**.
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The partitioning mode of the database. Valid values:
	//
	// *   **HORIZONTAL**: The database is horizontally partitioned.
	// *   **VERTICAL**: The database is vertically partitioned.
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The schema ID that is assigned to the partitioned database.
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The state of the database.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDrdsDBsResponseBodyDataDb) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBsResponseBodyDataDb) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetCreateTime(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetDbInstType(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetDbName(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetMode(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.Mode = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetSchema(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.Schema = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetStatus(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.Status = &v
	return s
}

type DescribeDrdsDBsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsDBsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsDBsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsResponse) SetHeaders(v map[string]*string) *DescribeDrdsDBsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDBsResponse) SetStatusCode(v int32) *DescribeDrdsDBsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDBsResponse) SetBody(v *DescribeDrdsDBsResponseBody) *DescribeDrdsDBsResponse {
	s.Body = v
	return s
}

type DescribeDrdsDbInstanceRequest struct {
	// The ID of the custom ApsaraDB RDS for MySQL instance that you want to query.
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The name of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsDbInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceRequest) SetDbInstanceId(v string) *DescribeDrdsDbInstanceRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceRequest) SetDbName(v string) *DescribeDrdsDbInstanceRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDbInstanceRequest) SetDrdsInstanceId(v string) *DescribeDrdsDbInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsDbInstanceResponseBody struct {
	// The detailed information about the returned custom ApsaraDB RDS for MySQL instance.
	DbInstance *DescribeDrdsDbInstanceResponseBodyDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDbInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBody) SetDbInstance(v *DescribeDrdsDbInstanceResponseBodyDbInstance) *DescribeDrdsDbInstanceResponseBody {
	s.DbInstance = v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBody) SetRequestId(v string) *DescribeDrdsDbInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBody) SetSuccess(v bool) *DescribeDrdsDbInstanceResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsDbInstanceResponseBodyDbInstance struct {
	// The URL used to connect to the custom ApsaraDB RDS for MySQL instance.
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// The ID of the ApsaraDB RDS for MySQL instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The state of the instance.
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The role of the instance. Valid values:
	//
	// *   **Primary**: The instance is a primary instance.
	// *   **ReadOnly**: The instance is a read-only instance.
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The ID of the resource.
	DmInstanceId *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	// The engine of the database that is run on the instance. Valid value: **MySQL**.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version of the database that is run on the instance. Valid values: **5.7**.
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the custom ApsaraDB RDS for MySQL instance expires. The value of this parameter is a UNIX timestamp. Unit: seconds.
	//
	// >  This parameter is returned only when the custom ApsaraDB RDS for MySQL instance is a subscription instance.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The type of the network. Valid values: **VPC**.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the custom ApsaraDB RDS for MySQL instance. Valid values:
	//
	// *   **Prepaid**: subscription
	// *   **Postaid**: pay-as-you-go
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port used to connect to the custom ApsaraDB RDS for MySQL instance.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the instance.
	RdsInstType *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	// The list of read-only ApsaraDB RDS for MySQL instances.
	ReadOnlyInstances *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances `json:"ReadOnlyInstances,omitempty" xml:"ReadOnlyInstances,omitempty" type:"Struct"`
	// The read ratio of the instance.
	ReadWeight *int32 `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	// The number of remaining days before the instance expires.
	RemainDays *string `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetConnectUrl(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDBInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDbInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetEngine(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetEngineVersion(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetExpireTime(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetNetworkType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetPayType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetPort(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetRdsInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetReadOnlyInstances(v *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ReadOnlyInstances = v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetReadWeight(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetRemainDays(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.RemainDays = &v
	return s
}

type DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances struct {
	ReadOnlyInstance []*DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance `json:"ReadOnlyInstance,omitempty" xml:"ReadOnlyInstance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) SetReadOnlyInstance(v []*DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances {
	s.ReadOnlyInstance = v
	return s
}

type DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance struct {
	// The URL used to connect to the read-only instance.
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// The ID of the read-only instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The state of the read-only instance.
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The role of the read-only instance.
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The ID of the resource.
	DmInstanceId *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	// The engine of the database that is run on the read-only instance.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version of the database that is run on the read-only instance.
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The timestamp that indicates when the read-only instance expires.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The network type of the read-only instance.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the read-only instance.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port used to connect to the read-only instance.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the ApsaraDB RDS for MySQL instance.
	RdsInstType *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	// The read ratio of the read-only instance.
	ReadWeight *int32 `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	// The number of remaining days before the read-only instance expires.
	RemainDays *string `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
	// This parameter is unavailable for read-only instances.
	VersionAction *int32 `json:"VersionAction,omitempty" xml:"VersionAction,omitempty"`
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetConnectUrl(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDBInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDbInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngine(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngineVersion(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetExpireTime(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetNetworkType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetPayType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetPort(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetRdsInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetReadWeight(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetRemainDays(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RemainDays = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetVersionAction(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.VersionAction = &v
	return s
}

type DescribeDrdsDbInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsDbInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsDbInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponse) SetHeaders(v map[string]*string) *DescribeDrdsDbInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDbInstanceResponse) SetStatusCode(v int32) *DescribeDrdsDbInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponse) SetBody(v *DescribeDrdsDbInstanceResponseBody) *DescribeDrdsDbInstanceResponse {
	s.Body = v
	return s
}

type DescribeDrdsDbInstancesRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The number of the page to return.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDrdsDbInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesRequest) SetDbName(v string) *DescribeDrdsDbInstancesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDbInstancesRequest) SetDrdsInstanceId(v string) *DescribeDrdsDbInstancesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesRequest) SetPageNumber(v int32) *DescribeDrdsDbInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsDbInstancesRequest) SetPageSize(v int32) *DescribeDrdsDbInstancesRequest {
	s.PageSize = &v
	return s
}

type DescribeDrdsDbInstancesResponseBody struct {
	// Indicates information about the ApsaraDB RDS for MySQL instances that are used to store the data of the specified database.
	DbInstances *DescribeDrdsDbInstancesResponseBodyDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Struct"`
	// Indicates the page number of the returned page.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Indicates the number of entries returned per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates the ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates the number of primary ApsaraDB RDS for MySQL instances.
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsDbInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBody) SetDbInstances(v *DescribeDrdsDbInstancesResponseBodyDbInstances) *DescribeDrdsDbInstancesResponseBody {
	s.DbInstances = v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetPageNumber(v string) *DescribeDrdsDbInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetPageSize(v string) *DescribeDrdsDbInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetRequestId(v string) *DescribeDrdsDbInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetSuccess(v bool) *DescribeDrdsDbInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetTotal(v string) *DescribeDrdsDbInstancesResponseBody {
	s.Total = &v
	return s
}

type DescribeDrdsDbInstancesResponseBodyDbInstances struct {
	DbInstance []*DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstances) SetDbInstance(v []*DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) *DescribeDrdsDbInstancesResponseBodyDbInstances {
	s.DbInstance = v
	return s
}

type DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance struct {
	// Indicates the endpoint that is used to connect to an ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// Indicates the ID of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Indicates the state of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database. Valid values:
	//
	// *   **0**: The ApsaraDB RDS for MySQL instance is being created.
	// *   **1**: The ApsaraDB RDS for MySQL instance is running.
	// *   **3**: The ApsaraDB RDS for MySQL instance is being deleted.
	// *   **5**: The ApsaraDB RDS for MySQL instance is being restarted.
	// *   **6**: The ApsaraDB RDS for MySQL instance is being upgraded or downgraded.
	// *   **7**: The ApsaraDB RDS for MySQL instance is being backed up.
	// *   **8**: The network type of the ApsaraDB RDS for MySQL instance is being changed.
	// *   **9**: The ApsaraDB RDS for MySQL instance is being migrated.
	// *   **11**: The data of the ApsaraDB RDS for MySQL instance is being migrated.
	// *   **12**: A disaster-recovery instance is being generated.
	// *   **13**: Data is being imported to the ApsaraDB RDS for MySQL instance.
	// *   **14**: Data is being imported to the ApsaraDB RDS for MySQL instance from an another ApsaraDB RDS for MySQL instance.
	// *   **15**: A failover is being performed.
	// *   **16**: A temporary instance is being created.
	// *   **17**: A network is being created for the ApsaraDB RDS for MySQL instance.
	// *   **18**: The ApsaraDB RDS for MySQL instance is being cloned.
	// *   **19**: The link is being changed.
	// *   **20**: The read-only instances of the ApsaraDB RDS for MySQL instance are being migrated.
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// Indicates the type of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database. The value is set to RDS.
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// Indicates the ID of a resource.
	DmInstanceId *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	// Indicates the engine of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Indicates the engine version of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Indicates the point in time when the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database expires.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates the network type of the ApsaraDB RDS for MySQL instance.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// Indicates the billing method of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database. Valid values:
	//
	// *   **drdsPre**: The instance uses the subscription billing method.
	// *   **drdsPost**: The instance uses the pay-as-you-go billing method.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Indicates the port that is used to connect to the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Indicates whether the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database is a primary instance or a read-only instance.
	//
	// *   **Primary**: The instance is a primary instance.
	// *   **Readonly**: The instance is a read-only instance.
	RdsInstType *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	// Indicates information about the read-only instances of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	ReadOnlyInstances *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances `json:"ReadOnlyInstances,omitempty" xml:"ReadOnlyInstances,omitempty" type:"Struct"`
	// Indicates the read weight of the read-only instance.
	ReadWeight *int32 `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	// Indicates the number of remaining days before a subscription instance expires.
	RemainDays *int32 `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetConnectUrl(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceId(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDbInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetEngine(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetEngineVersion(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetExpireTime(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetNetworkType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetPayType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetPort(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetRdsInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetReadOnlyInstances(v *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ReadOnlyInstances = v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetReadWeight(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetRemainDays(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.RemainDays = &v
	return s
}

type DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances struct {
	ReadOnlyInstance []*DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance `json:"ReadOnlyInstance,omitempty" xml:"ReadOnlyInstance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) SetReadOnlyInstance(v []*DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances {
	s.ReadOnlyInstance = v
	return s
}

type DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance struct {
	// Indicates the endpoint that is used to connect to the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// Indicates the state of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database. Valid values:
	//
	// *   **0**: The ApsaraDB RDS for MySQL instance is being created.
	// *   **1**: The ApsaraDB RDS for MySQL instance is running.
	// *   **3**: The ApsaraDB RDS for MySQL instance is being deleted.
	// *   **5**: The ApsaraDB RDS for MySQL instance is being restarted.
	// *   **6**: The ApsaraDB RDS for MySQL instance is being upgraded or downgraded.
	// *   **7**: The ApsaraDB RDS for MySQL instance is being backed up.
	// *   **8**: The network type of the ApsaraDB RDS for MySQL instance is being changed.
	// *   **9**: The ApsaraDB RDS for MySQL instance is being migrated.
	// *   **11**: The data of the ApsaraDB RDS for MySQL instance is being migrated.
	// *   **12**: A disaster-recovery instance is being generated.
	// *   **13**: Data is being imported to the ApsaraDB RDS for MySQL instance.
	// *   **14**: Data is being imported to the ApsaraDB RDS for MySQL instance from an another ApsaraDB RDS for MySQL instance.
	// *   **15**: A failover is being performed.
	// *   **16**: A temporary instance is being created.
	// *   **17**: A network is being created for the ApsaraDB RDS for MySQL instance.
	// *   **18**: The ApsaraDB RDS for MySQL instance is being cloned.
	// *   **19**: The link is being changed.
	// *   **20**: The read-only instances of the ApsaraDB RDS for MySQL instance are being migrated.
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// Indicates the type of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database. The value is set to RDS.
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// Indicates the ID of a resource.
	DmInstanceId *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	// Indicates the engine of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Indicates the engine version of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Indicates the timestamp when the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database expires.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates the name of a read-only instance.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Indicates the network type of the read-only instance.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// Indicates the billing method of the read-only instance.
	//
	// *   **drdsPre**: The instance uses the subscription billing method.
	// *   **drdsPost**: The instance uses the pay-as-you-go billing method.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Indicates the port that is used to connect to the read-only instance.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Indicates the type of the read-only instance.
	RdsInstType *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	// Indicates the read weight of the read-only instance.
	ReadWeight *int32 `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	// Indicates the number of remaining days before the read-only instance expires.
	RemainDays *int32 `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetConnectUrl(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetDbInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngine(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngineVersion(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetExpireTime(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetInstanceName(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetNetworkType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetPayType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetPort(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetRdsInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetReadWeight(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetRemainDays(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RemainDays = &v
	return s
}

type DescribeDrdsDbInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsDbInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsDbInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponse) SetHeaders(v map[string]*string) *DescribeDrdsDbInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDbInstancesResponse) SetStatusCode(v int32) *DescribeDrdsDbInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponse) SetBody(v *DescribeDrdsDbInstancesResponseBody) *DescribeDrdsDbInstancesResponse {
	s.Body = v
	return s
}

type DescribeDrdsDbRdsNameListRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsDbRdsNameListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbRdsNameListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbRdsNameListRequest) SetDbName(v string) *DescribeDrdsDbRdsNameListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListRequest) SetDrdsInstanceId(v string) *DescribeDrdsDbRdsNameListRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsDbRdsNameListResponseBody struct {
	// Indicates the instances that are used to store the data of a database.
	InstanceNameList *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList `json:"InstanceNameList,omitempty" xml:"InstanceNameList,omitempty" type:"Struct"`
	// Indicates the ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDbRdsNameListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbRdsNameListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbRdsNameListResponseBody) SetInstanceNameList(v *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) *DescribeDrdsDbRdsNameListResponseBody {
	s.InstanceNameList = v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponseBody) SetRequestId(v string) *DescribeDrdsDbRdsNameListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponseBody) SetSuccess(v bool) *DescribeDrdsDbRdsNameListResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsDbRdsNameListResponseBodyInstanceNameList struct {
	InstanceName []*string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) SetInstanceName(v []*string) *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList {
	s.InstanceName = v
	return s
}

type DescribeDrdsDbRdsNameListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsDbRdsNameListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsDbRdsNameListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbRdsNameListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbRdsNameListResponse) SetHeaders(v map[string]*string) *DescribeDrdsDbRdsNameListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponse) SetStatusCode(v int32) *DescribeDrdsDbRdsNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponse) SetBody(v *DescribeDrdsDbRdsNameListResponseBody) *DescribeDrdsDbRdsNameListResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstanceRequest struct {
	// The ID of the instance that you want to query.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region in which the instance is created.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceRequest) SetRegionId(v string) *DescribeDrdsInstanceRequest {
	s.RegionId = &v
	return s
}

type DescribeDrdsInstanceResponseBody struct {
	// The details of the instance.
	Data *DescribeDrdsInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBody) SetData(v *DescribeDrdsInstanceResponseBodyData) *DescribeDrdsInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsInstanceResponseBody) SetRequestId(v string) *DescribeDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsInstanceResponseBodyData struct {
	// The commodity code of the instance.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The timestamp that indicates when the instance is created.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The timestamp that indicates when the instance expires.
	ExpireDate *int64 `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The role of the instance. Valid values:
	//
	// *   **MASTER**: The instance is a primary instance.
	// *   **SLAVE**: The instance is a read-only instance to analyze complex queries
	// *   **SLAVE_FLOW**: The instance is a read-only instance for high-concurrency scenarios
	InstRole *string `json:"InstRole,omitempty" xml:"InstRole,omitempty"`
	// The instance series of the instance.
	InstanceSeries *string `json:"InstanceSeries,omitempty" xml:"InstanceSeries,omitempty"`
	// The specification of the instance.
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The tag of the instance. Valid values:
	//
	// *   **NORMAL**: The instance is a standard instance.
	// *   **HA**: The instance is a high-availability (HA) instance.
	// *   **VPC**: The instance is a VPC-based instance.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The machine type of the instance. The value of this parameter is **ecs**.
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The ID of the primary instance.
	//
	// >  This parameter is returned only when the instance is a primary instance.
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	// The MySQL version that is supported by the instance.
	MysqlVersion *int32 `json:"MysqlVersion,omitempty" xml:"MysqlVersion,omitempty"`
	// The network type of the instance. Valid values: CLASSIC and VPC.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the purchased instance.
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	// The version of .
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// The details about each read-only instance that is associated with the instance.
	ReadOnlyDBInstanceIds *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds `json:"ReadOnlyDBInstanceIds,omitempty" xml:"ReadOnlyDBInstanceIds,omitempty" type:"Struct"`
	// The ID of the region in which the instance is created.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. The value of this parameter can be null.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The state of the instance.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the instance used for storage.
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The type of the instance. Valid values: PRIVATE and PUBLIC. The value of PRIVATE indicates that the instance is a dedicated instance. The value of PUBLIC indicates that the instance is a shared instance.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the instance. The value of this parameter is 0.
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
	// Indicates whether the version of the instance can be upgraded.
	VersionAction *string `json:"VersionAction,omitempty" xml:"VersionAction,omitempty"`
	// The list of returned virtual IP addresses (VIPs).
	Vips *DescribeDrdsInstanceResponseBodyDataVips `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
	// The ID of the instance that is deployed in the VPC.
	VpcCloudInstanceId *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	// The ID of the zone in which the instance is located.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDrdsInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyData) SetCommodityCode(v string) *DescribeDrdsInstanceResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetCreateTime(v int64) *DescribeDrdsInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetDescription(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetDrdsInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetExpireDate(v int64) *DescribeDrdsInstanceResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetInstRole(v string) *DescribeDrdsInstanceResponseBodyData {
	s.InstRole = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetInstanceSeries(v string) *DescribeDrdsInstanceResponseBodyData {
	s.InstanceSeries = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetInstanceSpec(v string) *DescribeDrdsInstanceResponseBodyData {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetLabel(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Label = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetMachineType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.MachineType = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetMasterInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.MasterInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetMysqlVersion(v int32) *DescribeDrdsInstanceResponseBodyData {
	s.MysqlVersion = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetNetworkType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetOrderInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.OrderInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetProductVersion(v string) *DescribeDrdsInstanceResponseBodyData {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetReadOnlyDBInstanceIds(v *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) *DescribeDrdsInstanceResponseBodyData {
	s.ReadOnlyDBInstanceIds = v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetRegionId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetResourceGroupId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetStatus(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetStorageType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.StorageType = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVersion(v int64) *DescribeDrdsInstanceResponseBodyData {
	s.Version = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVersionAction(v string) *DescribeDrdsInstanceResponseBodyData {
	s.VersionAction = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVips(v *DescribeDrdsInstanceResponseBodyDataVips) *DescribeDrdsInstanceResponseBodyData {
	s.Vips = v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVpcCloudInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetZoneId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.ZoneId = &v
	return s
}

type DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds struct {
	ReadOnlyDBInstanceId []*string `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) SetReadOnlyDBInstanceId(v []*string) *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds {
	s.ReadOnlyDBInstanceId = v
	return s
}

type DescribeDrdsInstanceResponseBodyDataVips struct {
	Vip []*DescribeDrdsInstanceResponseBodyDataVipsVip `json:"Vip,omitempty" xml:"Vip,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceResponseBodyDataVips) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyDataVips) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyDataVips) SetVip(v []*DescribeDrdsInstanceResponseBodyDataVipsVip) *DescribeDrdsInstanceResponseBodyDataVips {
	s.Vip = v
	return s
}

type DescribeDrdsInstanceResponseBodyDataVipsVip struct {
	// The domain name that is mapped to the VIP.
	Dns *string `json:"Dns,omitempty" xml:"Dns,omitempty"`
	// The number of remaining days before the VIP expires.
	ExpireDays *int64 `json:"ExpireDays,omitempty" xml:"ExpireDays,omitempty"`
	// The ports that are opened on the VIP.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the VIP. Valid values: intranet and internet.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeDrdsInstanceResponseBodyDataVipsVip) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyDataVipsVip) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetDns(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.Dns = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetExpireDays(v int64) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.ExpireDays = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetPort(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.Port = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetType(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetVpcId(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.VpcId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetVswitchId(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.VswitchId = &v
	return s
}

type DescribeDrdsInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceResponse) SetStatusCode(v int32) *DescribeDrdsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstanceResponse) SetBody(v *DescribeDrdsInstanceResponseBody) *DescribeDrdsInstanceResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstanceDbMonitorRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the Distributed Relational Database Service (DRDS) instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The end time. Specify the time in the UNIX timestamp format. The time must be in UTC. Unit: ms.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance monitoring metrics. You can specify one or more metrics for a query at a time. Separate multiple metric parameters with commas (,).
	//
	// >  For more information about the details of performance monitoring metrics, see [Database monitoring](~~186704~~).
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. Specify the time in the UNIX timestamp format. The time must be in UTC. Unit: ms.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDrdsInstanceDbMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetDbName(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetEndTime(v int64) *DescribeDrdsInstanceDbMonitorRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetKey(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetRegionId(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetStartTime(v int64) *DescribeDrdsInstanceDbMonitorRequest {
	s.StartTime = &v
	return s
}

type DescribeDrdsInstanceDbMonitorResponseBody struct {
	// The list of monitoring data.
	Data []*DescribeDrdsInstanceDbMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) SetData(v []*DescribeDrdsInstanceDbMonitorResponseBodyData) *DescribeDrdsInstanceDbMonitorResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) SetRequestId(v string) *DescribeDrdsInstanceDbMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceDbMonitorResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsInstanceDbMonitorResponseBodyData struct {
	// The name of the monitoring metric.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The unit of the monitoring metric.
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The details about the value of monitoring data.
	Values []*DescribeDrdsInstanceDbMonitorResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyData) SetKey(v string) *DescribeDrdsInstanceDbMonitorResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyData) SetUnit(v string) *DescribeDrdsInstanceDbMonitorResponseBodyData {
	s.Unit = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyData) SetValues(v []*DescribeDrdsInstanceDbMonitorResponseBodyDataValues) *DescribeDrdsInstanceDbMonitorResponseBodyData {
	s.Values = v
	return s
}

type DescribeDrdsInstanceDbMonitorResponseBodyDataValues struct {
	// The time point when the value of monitoring data was obtained. The value is in the UNIX timestamp format. Unit: ms.
	Date *int64 `json:"Date,omitempty" xml:"Date,omitempty"`
	// The data value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyDataValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataValues) SetDate(v int64) *DescribeDrdsInstanceDbMonitorResponseBodyDataValues {
	s.Date = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataValues) SetValue(v string) *DescribeDrdsInstanceDbMonitorResponseBodyDataValues {
	s.Value = &v
	return s
}

type DescribeDrdsInstanceDbMonitorResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsInstanceDbMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsInstanceDbMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceDbMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponse) SetStatusCode(v int32) *DescribeDrdsInstanceDbMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponse) SetBody(v *DescribeDrdsInstanceDbMonitorResponseBody) *DescribeDrdsInstanceDbMonitorResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstanceLevelTasksRequest struct {
	// The ID of the PolarDB-X 1.0 instance of which the unfinished tasks you want to query.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsInstanceLevelTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceLevelTasksRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsInstanceLevelTasksResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The list of returned unfinished tasks.
	Tasks *DescribeDrdsInstanceLevelTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s DescribeDrdsInstanceLevelTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) SetRequestId(v string) *DescribeDrdsInstanceLevelTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceLevelTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) SetTasks(v *DescribeDrdsInstanceLevelTasksResponseBodyTasks) *DescribeDrdsInstanceLevelTasksResponseBody {
	s.Tasks = v
	return s
}

type DescribeDrdsInstanceLevelTasksResponseBodyTasks struct {
	Task []*DescribeDrdsInstanceLevelTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceLevelTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasks) SetTask(v []*DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) *DescribeDrdsInstanceLevelTasksResponseBodyTasks {
	s.Task = v
	return s
}

type DescribeDrdsInstanceLevelTasksResponseBodyTasksTask struct {
	// Indicates whether the task can be canceled.
	AllowCancel *bool `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// The error message returned for the task.
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The timestamp when the task is created.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The progress of the task. Valid values: 0 to 100.
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The description of the task progress.
	ProgressDescription *string `json:"ProgressDescription,omitempty" xml:"ProgressDescription,omitempty"`
	// Indicates whether the progress of the task is displayed.
	ShowProgress *bool `json:"ShowProgress,omitempty" xml:"ShowProgress,omitempty"`
	// The ID of the task.
	TargetId *int64 `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The phase of the task.
	TaskPhase *string `json:"TaskPhase,omitempty" xml:"TaskPhase,omitempty"`
	// The state of the task. Valid values:
	//
	// *   0: The task is being executed.
	// *   1: The task is executed.
	// *   2: The task failed to be executed.
	// *   3: The task is canceled.
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The type of the task.
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetAllowCancel(v bool) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.AllowCancel = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetErrMsg(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.ErrMsg = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetGmtCreate(v int64) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetProgress(v int32) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.Progress = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetProgressDescription(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.ProgressDescription = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetShowProgress(v bool) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.ShowProgress = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTargetId(v int64) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TargetId = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskName(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskName = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskPhase(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskPhase = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskStatus(v int32) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskStatus = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskType(v int32) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskType = &v
	return s
}

type DescribeDrdsInstanceLevelTasksResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsInstanceLevelTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsInstanceLevelTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceLevelTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponse) SetStatusCode(v int32) *DescribeDrdsInstanceLevelTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponse) SetBody(v *DescribeDrdsInstanceLevelTasksResponseBody) *DescribeDrdsInstanceLevelTasksResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstanceMonitorRequest struct {
	// The ID of the instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The end time of the query. Specify the value in the UNIX timestamp format. The timestamp must be in UTC. Unit: ms.
	//
	// >  If the time range that you specify is less than 1 hour, the monitoring data that is collected in a 1-hour period before the end time is returned.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance monitoring metrics. You can specify one or more metrics. Separate multiple metric names with commas (,).
	//
	// >  For more information about performance monitoring metrics, see [Monitor instances](~~186703~~).
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The multiple of the default time interval that you want to use to collect monitoring data. By default, the system collects monitoring data of resources at an interval of 1 minute. If you set the value of this parameter to 2, the system collects monitoring data of the instance at an interval of 2 minutes.
	PeriodMultiple *int32 `json:"PeriodMultiple,omitempty" xml:"PeriodMultiple,omitempty"`
	// The ID of the region where the instance is deployed.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the query. Specify the value in the UNIX timestamp format. The timestamp must be in UTC. Unit: ms.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDrdsInstanceMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceMonitorRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetEndTime(v int64) *DescribeDrdsInstanceMonitorRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetKey(v string) *DescribeDrdsInstanceMonitorRequest {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetPeriodMultiple(v int32) *DescribeDrdsInstanceMonitorRequest {
	s.PeriodMultiple = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetRegionId(v string) *DescribeDrdsInstanceMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetStartTime(v int64) *DescribeDrdsInstanceMonitorRequest {
	s.StartTime = &v
	return s
}

type DescribeDrdsInstanceMonitorResponseBody struct {
	// The result set of the query.
	Data []*DescribeDrdsInstanceMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDrdsInstanceMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBody) SetData(v []*DescribeDrdsInstanceMonitorResponseBodyData) *DescribeDrdsInstanceMonitorResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBody) SetRequestId(v string) *DescribeDrdsInstanceMonitorResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDrdsInstanceMonitorResponseBodyData struct {
	// The name of the metric.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The number of nodes.
	NodeNum *int32 `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// The unit of the metric value.
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The details of the monitoring data of the metric.
	Values []*DescribeDrdsInstanceMonitorResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceMonitorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetKey(v string) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetNodeNum(v int32) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.NodeNum = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetUnit(v string) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.Unit = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetValues(v []*DescribeDrdsInstanceMonitorResponseBodyDataValues) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.Values = v
	return s
}

type DescribeDrdsInstanceMonitorResponseBodyDataValues struct {
	// The point in time when the value of the metric was collected. The value is in the UNIX timestamp format. The timestamp is displayed in UTC. Unit: ms.
	Date *int64 `json:"Date,omitempty" xml:"Date,omitempty"`
	// The value of the metric.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDrdsInstanceMonitorResponseBodyDataValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataValues) SetDate(v int64) *DescribeDrdsInstanceMonitorResponseBodyDataValues {
	s.Date = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataValues) SetValue(v string) *DescribeDrdsInstanceMonitorResponseBodyDataValues {
	s.Value = &v
	return s
}

type DescribeDrdsInstanceMonitorResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsInstanceMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsInstanceMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponse) SetStatusCode(v int32) *DescribeDrdsInstanceMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponse) SetBody(v *DescribeDrdsInstanceMonitorResponseBody) *DescribeDrdsInstanceMonitorResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstanceVersionRequest struct {
	// The ID of the PolarDB-X 1.0 instance whose version you want to query.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsInstanceVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceVersionRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceVersionRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceVersionRequest) SetRegionId(v string) *DescribeDrdsInstanceVersionRequest {
	s.RegionId = &v
	return s
}

type DescribeDrdsInstanceVersionResponseBody struct {
	// The details about the instance version.
	Data *DescribeDrdsInstanceVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsInstanceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceVersionResponseBody) SetData(v *DescribeDrdsInstanceVersionResponseBodyData) *DescribeDrdsInstanceVersionResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsInstanceVersionResponseBody) SetRequestId(v string) *DescribeDrdsInstanceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceVersionResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceVersionResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsInstanceVersionResponseBodyData struct {
	// The current version of the instance.
	InstanceVersion *string `json:"InstanceVersion,omitempty" xml:"InstanceVersion,omitempty"`
	// The latest version of the instance.
	NewestVersion *string `json:"NewestVersion,omitempty" xml:"NewestVersion,omitempty"`
}

func (s DescribeDrdsInstanceVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceVersionResponseBodyData) SetInstanceVersion(v string) *DescribeDrdsInstanceVersionResponseBodyData {
	s.InstanceVersion = &v
	return s
}

func (s *DescribeDrdsInstanceVersionResponseBodyData) SetNewestVersion(v string) *DescribeDrdsInstanceVersionResponseBodyData {
	s.NewestVersion = &v
	return s
}

type DescribeDrdsInstanceVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsInstanceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsInstanceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceVersionResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceVersionResponse) SetStatusCode(v int32) *DescribeDrdsInstanceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstanceVersionResponse) SetBody(v *DescribeDrdsInstanceVersionResponseBody) *DescribeDrdsInstanceVersionResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstancesRequest struct {
	// The description of the instances.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the instances that you want to query expire.
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Specifies whether hybrid queries are supported.
	Mix *bool `json:"Mix,omitempty" xml:"Mix,omitempty"`
	// The number of the page to return.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of instances returned on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The version of the service.
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instances you want to query belong. The value of this parameter can be NULL.
	ResourceGroupId *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*DescribeDrdsInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of the instances that you want to query. Valid values:
	//
	// *   **0**: shared instances
	// *   **1**: dedicated instances
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDrdsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesRequest) SetDescription(v string) *DescribeDrdsInstancesRequest {
	s.Description = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetExpired(v bool) *DescribeDrdsInstancesRequest {
	s.Expired = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetMix(v bool) *DescribeDrdsInstancesRequest {
	s.Mix = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetPageNumber(v int32) *DescribeDrdsInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetPageSize(v int32) *DescribeDrdsInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetProductVersion(v string) *DescribeDrdsInstancesRequest {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetRegionId(v string) *DescribeDrdsInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetResourceGroupId(v string) *DescribeDrdsInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetTag(v []*DescribeDrdsInstancesRequestTag) *DescribeDrdsInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetType(v string) *DescribeDrdsInstancesRequest {
	s.Type = &v
	return s
}

type DescribeDrdsInstancesRequestTag struct {
	// The key of the tag configured for the instances you want to query.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag configured for the instances you want to query.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDrdsInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesRequestTag) SetKey(v string) *DescribeDrdsInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstancesRequestTag) SetValue(v string) *DescribeDrdsInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeDrdsInstancesResponseBody struct {
	// The list of returned instances.
	Instances *DescribeDrdsInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of instances returned on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances returned.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBody) SetInstances(v *DescribeDrdsInstancesResponseBodyInstances) *DescribeDrdsInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetPageNumber(v int32) *DescribeDrdsInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetPageSize(v int32) *DescribeDrdsInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetRequestId(v string) *DescribeDrdsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetTotal(v int32) *DescribeDrdsInstancesResponseBody {
	s.Total = &v
	return s
}

type DescribeDrdsInstancesResponseBodyInstances struct {
	Instance []*DescribeDrdsInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstances) SetInstance(v []*DescribeDrdsInstancesResponseBodyInstancesInstance) *DescribeDrdsInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

type DescribeDrdsInstancesResponseBodyInstancesInstance struct {
	// The commodity code of the service.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The timestamp that indicates when the instance is created.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The timestamp that indicates when the instance expires.
	ExpireDate *int64 `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The role of the instance. Valid values:
	//
	// *   MASTER: The instance is a primary instance.
	// *   SLAVE: The instance is a read-only instance to analyze complex queries.
	// *   SLAVE_FLOW: The instance is a read-only instance for high-concurrency scenarios.
	InstRole *string `json:"InstRole,omitempty" xml:"InstRole,omitempty"`
	// The instance series.
	InstanceSeries *string `json:"InstanceSeries,omitempty" xml:"InstanceSeries,omitempty"`
	// The specification of the instance.
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The tag of the instance. Valid values:
	//
	// *   **NORMAL**: The instance is a standard instance.
	// *   **HA**: The instance is a high-availability (HA) instance.
	// *   **VPC**: The instance is a VPC-based instance.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The machine type of the instance. Valid value: ecs.
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The ID of the primary instance.
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	// The network type of the instance. Valid values:
	//
	// *   **CLASSIC**
	// *   **VPC**
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the purchased instance.
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	// The version of the service.
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// The IDs of read-only instances that are associated with the instance.
	ReadOnlyDBInstanceIds *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds `json:"ReadOnlyDBInstanceIds,omitempty" xml:"ReadOnlyDBInstanceIds,omitempty" type:"Struct"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the instance.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the instance. Valid values:
	//
	// *   **PUBLIC**: The returned instance is a shared instance.
	// *   **PRIVATE**: The returned instance is a dedicated instance.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the instance.
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
	// Indicates whether the version of the instance can be upgraded.
	VersionAction *string `json:"VersionAction,omitempty" xml:"VersionAction,omitempty"`
	// The list of returned virtual IP addresses (VIPs).
	Vips *DescribeDrdsInstancesResponseBodyInstancesInstanceVips `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
	// The ID of the instance that is deployed in the VPC.
	VpcCloudInstanceId *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	// The ID of the VPC to which the instance belongs.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone in which the resource is located.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The edition of the instance. Valid values:
	//
	// *   **starter**: Starter Edition
	// *   **enterprise**: Enterprise Edition
	// *   **standard**: Standard Edition
	Series *string `json:"series,omitempty" xml:"series,omitempty"`
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetCommodityCode(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetCreateTime(v int64) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetDescription(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Description = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetDrdsInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetExpireDate(v int64) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ExpireDate = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetInstRole(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.InstRole = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetInstanceSeries(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.InstanceSeries = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetInstanceSpec(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetLabel(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Label = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetMachineType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.MachineType = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetMasterInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.MasterInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetNetworkType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetOrderInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.OrderInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetProductVersion(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetReadOnlyDBInstanceIds(v *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ReadOnlyDBInstanceIds = v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetRegionId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetResourceGroupId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetStatus(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVersion(v int64) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Version = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVersionAction(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.VersionAction = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVips(v *DescribeDrdsInstancesResponseBodyInstancesInstanceVips) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Vips = v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVpcCloudInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVpcId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetZoneId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetSeries(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Series = &v
	return s
}

type DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds struct {
	ReadOnlyDBInstanceId []*string `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) SetReadOnlyDBInstanceId(v []*string) *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds {
	s.ReadOnlyDBInstanceId = v
	return s
}

type DescribeDrdsInstancesResponseBodyInstancesInstanceVips struct {
	Vip []*DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip `json:"Vip,omitempty" xml:"Vip,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceVips) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceVips) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVips) SetVip(v []*DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) *DescribeDrdsInstancesResponseBodyInstancesInstanceVips {
	s.Vip = v
	return s
}

type DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip struct {
	// The virtual IP address.
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ports that are opened on the VIP.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the VIP. Valid values:
	//
	// *   intranet: a private IP address
	// *   internet: a public IP address
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The domain name that is mapped to the VIP.
	Dns *string `json:"dns,omitempty" xml:"dns,omitempty"`
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetIP(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.IP = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetPort(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.Port = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetVpcId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.VpcId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetVswitchId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.VswitchId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetDns(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.Dns = &v
	return s
}

type DescribeDrdsInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstancesResponse) SetStatusCode(v int32) *DescribeDrdsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstancesResponse) SetBody(v *DescribeDrdsInstancesResponseBody) *DescribeDrdsInstancesResponse {
	s.Body = v
	return s
}

type DescribeDrdsParamsRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The type of nodes whose parameters you want to query. Valid values:
	//
	// *   **INSTANCE: the instance level.**
	// *   **DB**: the database level.
	ParamLevel *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	// The ID of the region where the PolarDB-X 1.0 instance is created.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsParamsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsParamsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsRequest) SetDbName(v string) *DescribeDrdsParamsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsParamsRequest) SetDrdsInstanceId(v string) *DescribeDrdsParamsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsParamsRequest) SetParamLevel(v string) *DescribeDrdsParamsRequest {
	s.ParamLevel = &v
	return s
}

func (s *DescribeDrdsParamsRequest) SetRegionId(v string) *DescribeDrdsParamsRequest {
	s.RegionId = &v
	return s
}

type DescribeDrdsParamsResponseBody struct {
	// Indicates information about parameters.
	List []*DescribeDrdsParamsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Indicates the ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsParamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsParamsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsResponseBody) SetList(v []*DescribeDrdsParamsResponseBodyList) *DescribeDrdsParamsResponseBody {
	s.List = v
	return s
}

func (s *DescribeDrdsParamsResponseBody) SetRequestId(v string) *DescribeDrdsParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsParamsResponseBody) SetSuccess(v bool) *DescribeDrdsParamsResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsParamsResponseBodyList struct {
	// Indicates the name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates whether a restart is required.
	NeedRestart *bool `json:"NeedRestart,omitempty" xml:"NeedRestart,omitempty"`
	// Indicates the default value of a parameter.
	ParamDefaultValue *string `json:"ParamDefaultValue,omitempty" xml:"ParamDefaultValue,omitempty"`
	// Indicates the description of the parameter.
	ParamDesc *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	// Indicates the name of the parameter.
	ParamEnglishName *string `json:"ParamEnglishName,omitempty" xml:"ParamEnglishName,omitempty"`
	// Indicates the parameter level.
	ParamLevel *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	// Indicates the name of the parameter.
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// Indicates the value range of the parameter.
	ParamRanges *string `json:"ParamRanges,omitempty" xml:"ParamRanges,omitempty"`
	// Indicates the type of the parameter.
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Indicates the value of the parameter.
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
	// Indicates the name of the variable.
	ParamVariableName *string `json:"ParamVariableName,omitempty" xml:"ParamVariableName,omitempty"`
}

func (s DescribeDrdsParamsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsParamsResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsResponseBodyList) SetDbName(v string) *DescribeDrdsParamsResponseBodyList {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetNeedRestart(v bool) *DescribeDrdsParamsResponseBodyList {
	s.NeedRestart = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamDefaultValue(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamDefaultValue = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamDesc(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamDesc = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamEnglishName(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamEnglishName = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamLevel(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamLevel = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamName(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamName = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamRanges(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamRanges = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamType(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamType = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamValue(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamValue = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamVariableName(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamVariableName = &v
	return s
}

type DescribeDrdsParamsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsParamsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsParamsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsResponse) SetHeaders(v map[string]*string) *DescribeDrdsParamsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsParamsResponse) SetStatusCode(v int32) *DescribeDrdsParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsParamsResponse) SetBody(v *DescribeDrdsParamsResponseBody) *DescribeDrdsParamsResponse {
	s.Body = v
	return s
}

type DescribeDrdsRdsInstancesRequest struct {
	// The ID of the PolarDB-X instance.
	//
	// > You can call the [DescribeDrdsInstances](~~139284~~) operation to query the information about instances in the specified account, such as the IDs of the instances.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsRdsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsRdsInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesRequest) SetDrdsInstanceId(v string) *DescribeDrdsRdsInstancesRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsRdsInstancesResponseBody struct {
	// The information about the custom ApsaraDB RDS for MySQL instances at the storage layer.
	DbInstances *DescribeDrdsRdsInstancesResponseBodyDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsRdsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsRdsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesResponseBody) SetDbInstances(v *DescribeDrdsRdsInstancesResponseBodyDbInstances) *DescribeDrdsRdsInstancesResponseBody {
	s.DbInstances = v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBody) SetRequestId(v string) *DescribeDrdsRdsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBody) SetSuccess(v bool) *DescribeDrdsRdsInstancesResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsRdsInstancesResponseBodyDbInstances struct {
	DbInstance []*DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsRdsInstancesResponseBodyDbInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsRdsInstancesResponseBodyDbInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstances) SetDbInstance(v []*DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) *DescribeDrdsRdsInstancesResponseBodyDbInstances {
	s.DbInstance = v
	return s
}

type DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance struct {
	// The internal endpoint of the custom ApsaraDB RDS for MySQL instance at the storage layer.
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// The number of CPU cores of the custom ApsaraDB RDS for MySQL instance at the storage layer.
	DBInstanceCPU *string `json:"DBInstanceCPU,omitempty" xml:"DBInstanceCPU,omitempty"`
	// The instance family of the custom ApsaraDB RDS for MySQL instance at the storage layer. Valid values:
	//
	// *   **x**: general-purpose instance family
	// *   **d**: dedicated instance family
	// *   **h**: dedicated host instance family
	DBInstanceClassType *string `json:"DBInstanceClassType,omitempty" xml:"DBInstanceClassType,omitempty"`
	// The ID of the custom ApsaraDB RDS for MySQL instance at the storage layer.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The memory size of the custom ApsaraDB RDS for MySQL instance at the storage layer. Unit: MB.
	DBInstanceMemory *int64 `json:"DBInstanceMemory,omitempty" xml:"DBInstanceMemory,omitempty"`
	// The status of the custom ApsaraDB RDS for MySQL instance at the storage layer. Valid values:
	//
	// *   0: The instance is being created.
	// *   1: The instance is running.
	// *   3: The instance is being deleted.
	// *   5: The instance is being restarted.
	// *   6: The instance is being upgraded or downgraded.
	// *   7: The instance is being backed up.
	// *   8: The network type of the instance is being changed.
	// *   9: The instance is being migrated.
	// *   11: The data stored on the instance is being migrated.
	// *   12: A disaster recovery instance is being generated.
	// *   13: Data is being imported to the instance.
	// *   14: Data is being imported from another RDS instance to the instance.
	// *   15: A switchover is being performed.
	// *   16: A temporary instance is being created.
	// *   17: The network of the instance is being created.
	// *   18: The instance is being cloned.
	// *   19: The link is being changed.
	// *   20: The read-only RDS instances of the instance are being migrated.
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The storage space of the custom ApsaraDB RDS for MySQL instance at the storage layer. Unit: GB.
	DBInstanceStorage *int64 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The type of the instance at the storage layer. The value is RDS.
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The ID of the resource.
	DmInstanceId *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	// The engine type of the custom ApsaraDB RDS for MySQL instance at the storage layer. The value is MySQL.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version of the custom ApsaraDB RDS for MySQL instance at the storage layer. The value is 8.0.
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The lock mode of the RDS instance. Valid values:
	//
	// 0: The instance is not locked.
	//
	// 1: The instance is manually locked.
	//
	// 2: The instance is automatically locked if the instance expires.
	//
	// 3: The instance is automatically locked if the instance is rolled back.
	//
	// 4: The instance is automatically locked if the storage space of the instance reaches the upper limit.
	//
	// 5: The instance is automatically locked if the storage space of the read-only instances reaches the upper limit.
	LockMode *int32 `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the RDS instance is locked.
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The network type of the custom ApsaraDB RDS for MySQL instance at the storage layer. The value is VPC.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the custom ApsaraDB RDS for MySQL instance at the storage layer. Valid values:
	//
	// *   Postpaid: pay-as-you-go
	// *   Prepaid: subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port used to connect to the instance over an internal network.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the custom ApsaraDB RDS for MySQL instance at the storage layer. Valid values:
	//
	// *   Primary: primary instance
	// *   Readonly: read-only instance
	RdsInstType *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	// The read and write weights of the custom ApsaraDB RDS for MySQL instance at the storage layer.
	ReadWeight *int32 `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
}

func (s DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetConnectUrl(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceCPU(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceCPU = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceClassType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceClassType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceId(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceMemory(v int64) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceMemory = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceStorage(v int64) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDbInstType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDmInstanceId(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetEngine(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetEngineVersion(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetLockMode(v int32) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetLockReason(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.LockReason = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetNetworkType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetPayType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetPort(v int32) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetRdsInstType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetReadWeight(v int32) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.ReadWeight = &v
	return s
}

type DescribeDrdsRdsInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsRdsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsRdsInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsRdsInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesResponse) SetHeaders(v map[string]*string) *DescribeDrdsRdsInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsRdsInstancesResponse) SetStatusCode(v int32) *DescribeDrdsRdsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponse) SetBody(v *DescribeDrdsRdsInstancesResponseBody) *DescribeDrdsRdsInstancesResponse {
	s.Body = v
	return s
}

type DescribeDrdsShardingDbsRequest struct {
	// The name of the database whose shards you want to query.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The matching pattern of the database name.
	DbNamePattern *string `json:"DbNamePattern,omitempty" xml:"DbNamePattern,omitempty"`
	// The ID of the PolarDB-X 1.0 instance whose database shards you want to query.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The page number of the returned page.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of database shards returned on each page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDrdsShardingDbsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsShardingDbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsRequest) SetDbName(v string) *DescribeDrdsShardingDbsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) SetDbNamePattern(v string) *DescribeDrdsShardingDbsRequest {
	s.DbNamePattern = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) SetDrdsInstanceId(v string) *DescribeDrdsShardingDbsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) SetPageNumber(v int64) *DescribeDrdsShardingDbsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) SetPageSize(v int64) *DescribeDrdsShardingDbsRequest {
	s.PageSize = &v
	return s
}

type DescribeDrdsShardingDbsResponseBody struct {
	// The page number of the returned page.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of database shards returned per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of returned database shards.
	ShardingDbs *DescribeDrdsShardingDbsResponseBodyShardingDbs `json:"ShardingDbs,omitempty" xml:"ShardingDbs,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of returned database shards.
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsShardingDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsShardingDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsResponseBody) SetPageNumber(v string) *DescribeDrdsShardingDbsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) SetPageSize(v string) *DescribeDrdsShardingDbsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) SetRequestId(v string) *DescribeDrdsShardingDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) SetShardingDbs(v *DescribeDrdsShardingDbsResponseBodyShardingDbs) *DescribeDrdsShardingDbsResponseBody {
	s.ShardingDbs = v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) SetSuccess(v bool) *DescribeDrdsShardingDbsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) SetTotal(v string) *DescribeDrdsShardingDbsResponseBody {
	s.Total = &v
	return s
}

type DescribeDrdsShardingDbsResponseBodyShardingDbs struct {
	ShardingDb []*DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb `json:"ShardingDb,omitempty" xml:"ShardingDb,omitempty" type:"Repeated"`
}

func (s DescribeDrdsShardingDbsResponseBodyShardingDbs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsShardingDbsResponseBodyShardingDbs) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbs) SetShardingDb(v []*DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) *DescribeDrdsShardingDbsResponseBodyShardingDbs {
	s.ShardingDb = v
	return s
}

type DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb struct {
	// The timeout period for a transaction to wait for the release of the data lock.
	BlockingTimeout *int32 `json:"BlockingTimeout,omitempty" xml:"BlockingTimeout,omitempty"`
	// The URL that is used to access the Apsara RDS for MySQL instance.
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// The properties of the connection string.
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// The ID of the Apsara RDS for MySQL instance that is used as the storage of the database shard.
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The status of the database.
	DbStatus *string `json:"DbStatus,omitempty" xml:"DbStatus,omitempty"`
	// The engine of the database.
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The name of group on which the database shard is stored.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The timeout period of an idle connection.
	IdleTimeOut *int32 `json:"IdleTimeOut,omitempty" xml:"IdleTimeOut,omitempty"`
	// The maximum size of the connection pool.
	MaxPoolSize *int32 `json:"MaxPoolSize,omitempty" xml:"MaxPoolSize,omitempty"`
	// The minimum size of the connection pool.
	MinPoolSize *int32 `json:"MinPoolSize,omitempty" xml:"MinPoolSize,omitempty"`
	// The size of cache for the returned results.
	PreparedStatementCacheSize *int32 `json:"PreparedStatementCacheSize,omitempty" xml:"PreparedStatementCacheSize,omitempty"`
	// The name of the database shard.
	ShardingDbName *string `json:"ShardingDbName,omitempty" xml:"ShardingDbName,omitempty"`
	// The username that is used to connect to the ApsaraDB RDS for MySQL instance.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetBlockingTimeout(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.BlockingTimeout = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetConnectUrl(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetConnectionProperties(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.ConnectionProperties = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetDbInstanceId(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetDbStatus(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.DbStatus = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetDbType(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.DbType = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetGroupName(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.GroupName = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetIdleTimeOut(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.IdleTimeOut = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetMaxPoolSize(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.MaxPoolSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetMinPoolSize(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.MinPoolSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetPreparedStatementCacheSize(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.PreparedStatementCacheSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetShardingDbName(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.ShardingDbName = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetUserName(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.UserName = &v
	return s
}

type DescribeDrdsShardingDbsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsShardingDbsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsShardingDbsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsShardingDbsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsResponse) SetHeaders(v map[string]*string) *DescribeDrdsShardingDbsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsShardingDbsResponse) SetStatusCode(v int32) *DescribeDrdsShardingDbsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponse) SetBody(v *DescribeDrdsShardingDbsResponseBody) *DescribeDrdsShardingDbsResponse {
	s.Body = v
	return s
}

type DescribeDrdsSlowSqlsRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The start time of the SQL query. Specify the time in the UNIX timestamp format. The time must be in UTC. Unit: ms.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The SQL execution time. Unit: ms.
	ExeTime *int64 `json:"ExeTime,omitempty" xml:"ExeTime,omitempty"`
	// The number of the page to return.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The end time of the SQL query. Specify the time in the UNIX timestamp format. The time must be in UTC. Unit: ms.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDrdsSlowSqlsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSlowSqlsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsRequest) SetDbName(v string) *DescribeDrdsSlowSqlsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetDrdsInstanceId(v string) *DescribeDrdsSlowSqlsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetEndTime(v int64) *DescribeDrdsSlowSqlsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetExeTime(v int64) *DescribeDrdsSlowSqlsRequest {
	s.ExeTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetPageNumber(v int32) *DescribeDrdsSlowSqlsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetPageSize(v int32) *DescribeDrdsSlowSqlsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetStartTime(v int64) *DescribeDrdsSlowSqlsRequest {
	s.StartTime = &v
	return s
}

type DescribeDrdsSlowSqlsResponseBody struct {
	// Indicates the details of the slow SQL query.
	Items *DescribeDrdsSlowSqlsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// Indicates the page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Indicates the number of entries returned on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates the ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates the total number of entries.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsSlowSqlsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSlowSqlsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetItems(v *DescribeDrdsSlowSqlsResponseBodyItems) *DescribeDrdsSlowSqlsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetPageNumber(v int32) *DescribeDrdsSlowSqlsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetPageSize(v int32) *DescribeDrdsSlowSqlsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetRequestId(v string) *DescribeDrdsSlowSqlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetSuccess(v bool) *DescribeDrdsSlowSqlsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetTotal(v int32) *DescribeDrdsSlowSqlsResponseBody {
	s.Total = &v
	return s
}

type DescribeDrdsSlowSqlsResponseBodyItems struct {
	Item []*DescribeDrdsSlowSqlsResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s DescribeDrdsSlowSqlsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSlowSqlsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsResponseBodyItems) SetItem(v []*DescribeDrdsSlowSqlsResponseBodyItemsItem) *DescribeDrdsSlowSqlsResponseBodyItems {
	s.Item = v
	return s
}

type DescribeDrdsSlowSqlsResponseBodyItemsItem struct {
	// Indicates the IP address of the execution machine.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Indicates the response time. Unit: ms.
	ResponseTime *int64 `json:"ResponseTime,omitempty" xml:"ResponseTime,omitempty"`
	// Indicates the name of the database.
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// Indicates the time when the slow SQL query was sent. Unit: ms.
	SendTime *int64 `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	// Indicates the content of the slow SQL query.
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s DescribeDrdsSlowSqlsResponseBodyItemsItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSlowSqlsResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetHost(v string) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.Host = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetResponseTime(v int64) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.ResponseTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetSchema(v string) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.Schema = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetSendTime(v int64) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.SendTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetSql(v string) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.Sql = &v
	return s
}

type DescribeDrdsSlowSqlsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsSlowSqlsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsSlowSqlsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSlowSqlsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsResponse) SetHeaders(v map[string]*string) *DescribeDrdsSlowSqlsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsSlowSqlsResponse) SetStatusCode(v int32) *DescribeDrdsSlowSqlsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponse) SetBody(v *DescribeDrdsSlowSqlsResponseBody) *DescribeDrdsSlowSqlsResponse {
	s.Body = v
	return s
}

type DescribeDrdsSqlAuditStatusRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsSqlAuditStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusRequest) SetDrdsInstanceId(v string) *DescribeDrdsSqlAuditStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsSqlAuditStatusResponseBody struct {
	// The returned data set.
	Data *DescribeDrdsSqlAuditStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsSqlAuditStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) SetData(v *DescribeDrdsSqlAuditStatusResponseBodyData) *DescribeDrdsSqlAuditStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) SetRequestId(v string) *DescribeDrdsSqlAuditStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) SetSuccess(v bool) *DescribeDrdsSqlAuditStatusResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsSqlAuditStatusResponseBodyData struct {
	Data []*DescribeDrdsSqlAuditStatusResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeDrdsSqlAuditStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyData) SetData(v []*DescribeDrdsSqlAuditStatusResponseBodyDataData) *DescribeDrdsSqlAuditStatusResponseBodyData {
	s.Data = v
	return s
}

type DescribeDrdsSqlAuditStatusResponseBodyDataData struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates whether the complete report of the SQL audit is supported. Valid values: true and false.
	Detailed *string `json:"Detailed,omitempty" xml:"Detailed,omitempty"`
	// Indicates whether the SQL audit feature is enabled for the database. Valid values: true and false.
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The UID of the external delivery.
	//
	// > This parameter is returned only if external log delivery is enabled.
	ExtraAliUid *int64 `json:"ExtraAliUid,omitempty" xml:"ExtraAliUid,omitempty"`
	// The Log Service Logstore from which logs are delivered.
	//
	// > This parameter is returned only if external log delivery is enabled.
	ExtraSlsLogStore *string `json:"ExtraSlsLogStore,omitempty" xml:"ExtraSlsLogStore,omitempty"`
	// The Log Service project from which logs are delivered.
	//
	// > This parameter is returned only if external log delivery is enabled.
	ExtraSlsProject *string `json:"ExtraSlsProject,omitempty" xml:"ExtraSlsProject,omitempty"`
	// Indicates whether external log delivery is enabled. Valid values: true and false.
	ExtraWriteEnabled *bool `json:"ExtraWriteEnabled,omitempty" xml:"ExtraWriteEnabled,omitempty"`
}

func (s DescribeDrdsSqlAuditStatusResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetDbName(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetDetailed(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.Detailed = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetEnabled(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.Enabled = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraAliUid(v int64) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraAliUid = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraSlsLogStore(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraSlsLogStore = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraSlsProject(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraSlsProject = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraWriteEnabled(v bool) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraWriteEnabled = &v
	return s
}

type DescribeDrdsSqlAuditStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsSqlAuditStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsSqlAuditStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusResponse) SetHeaders(v map[string]*string) *DescribeDrdsSqlAuditStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponse) SetStatusCode(v int32) *DescribeDrdsSqlAuditStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponse) SetBody(v *DescribeDrdsSqlAuditStatusResponseBody) *DescribeDrdsSqlAuditStatusResponse {
	s.Body = v
	return s
}

type DescribeDrdsTasksRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The type of tasks.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDrdsTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksRequest) SetDbName(v string) *DescribeDrdsTasksRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsTasksRequest) SetDrdsInstanceId(v string) *DescribeDrdsTasksRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsTasksRequest) SetTaskType(v string) *DescribeDrdsTasksRequest {
	s.TaskType = &v
	return s
}

type DescribeDrdsTasksResponseBody struct {
	// Indicates the ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates information about the tasks.
	Tasks *DescribeDrdsTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s DescribeDrdsTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksResponseBody) SetRequestId(v string) *DescribeDrdsTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsTasksResponseBody) SetSuccess(v bool) *DescribeDrdsTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsTasksResponseBody) SetTasks(v *DescribeDrdsTasksResponseBodyTasks) *DescribeDrdsTasksResponseBody {
	s.Tasks = v
	return s
}

type DescribeDrdsTasksResponseBodyTasks struct {
	Task []*DescribeDrdsTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeDrdsTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksResponseBodyTasks) SetTask(v []*DescribeDrdsTasksResponseBodyTasksTask) *DescribeDrdsTasksResponseBodyTasks {
	s.Task = v
	return s
}

type DescribeDrdsTasksResponseBodyTasksTask struct {
	// Indicates the content of a task.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Indicates the ID of the task.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates the state of the task.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeDrdsTasksResponseBodyTasksTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) SetContent(v string) *DescribeDrdsTasksResponseBodyTasksTask {
	s.Content = &v
	return s
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) SetId(v int64) *DescribeDrdsTasksResponseBodyTasksTask {
	s.Id = &v
	return s
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) SetState(v string) *DescribeDrdsTasksResponseBodyTasksTask {
	s.State = &v
	return s
}

type DescribeDrdsTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksResponse) SetHeaders(v map[string]*string) *DescribeDrdsTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsTasksResponse) SetStatusCode(v int32) *DescribeDrdsTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsTasksResponse) SetBody(v *DescribeDrdsTasksResponseBody) *DescribeDrdsTasksResponse {
	s.Body = v
	return s
}

type DescribeExpandLogicTableInfoListRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeExpandLogicTableInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListRequest) SetDbName(v string) *DescribeExpandLogicTableInfoListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListRequest) SetDrdsInstanceId(v string) *DescribeExpandLogicTableInfoListRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeExpandLogicTableInfoListResponseBody struct {
	// Indicates the result that is returned.
	Data *DescribeExpandLogicTableInfoListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExpandLogicTableInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListResponseBody) SetData(v *DescribeExpandLogicTableInfoListResponseBodyData) *DescribeExpandLogicTableInfoListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBody) SetRequestId(v string) *DescribeExpandLogicTableInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBody) SetSuccess(v bool) *DescribeExpandLogicTableInfoListResponseBody {
	s.Success = &v
	return s
}

type DescribeExpandLogicTableInfoListResponseBodyData struct {
	Data []*DescribeExpandLogicTableInfoListResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeExpandLogicTableInfoListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListResponseBodyData) SetData(v []*DescribeExpandLogicTableInfoListResponseBodyDataData) *DescribeExpandLogicTableInfoListResponseBodyData {
	s.Data = v
	return s
}

type DescribeExpandLogicTableInfoListResponseBodyDataData struct {
	// Indicates the database sharding key.
	ShardDbKey *string `json:"ShardDbKey,omitempty" xml:"ShardDbKey,omitempty"`
	// Indicates the table sharding key.
	ShardTbKey *string `json:"ShardTbKey,omitempty" xml:"ShardTbKey,omitempty"`
	// Indicates the name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeExpandLogicTableInfoListResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) SetShardDbKey(v string) *DescribeExpandLogicTableInfoListResponseBodyDataData {
	s.ShardDbKey = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) SetShardTbKey(v string) *DescribeExpandLogicTableInfoListResponseBodyDataData {
	s.ShardTbKey = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) SetTableName(v string) *DescribeExpandLogicTableInfoListResponseBodyDataData {
	s.TableName = &v
	return s
}

type DescribeExpandLogicTableInfoListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpandLogicTableInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpandLogicTableInfoListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListResponse) SetHeaders(v map[string]*string) *DescribeExpandLogicTableInfoListResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponse) SetStatusCode(v int32) *DescribeExpandLogicTableInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponse) SetBody(v *DescribeExpandLogicTableInfoListResponseBody) *DescribeExpandLogicTableInfoListResponse {
	s.Body = v
	return s
}

type DescribeHotDbListRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeHotDbListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListRequest) SetDbName(v string) *DescribeHotDbListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeHotDbListRequest) SetDrdsInstanceId(v string) *DescribeHotDbListRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeHotDbListResponseBody struct {
	// The result that is returned.
	Data *DescribeHotDbListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned.
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeHotDbListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBody) SetData(v *DescribeHotDbListResponseBodyData) *DescribeHotDbListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeHotDbListResponseBody) SetMsg(v string) *DescribeHotDbListResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeHotDbListResponseBody) SetRequestId(v string) *DescribeHotDbListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHotDbListResponseBody) SetSuccess(v bool) *DescribeHotDbListResponseBody {
	s.Success = &v
	return s
}

type DescribeHotDbListResponseBodyData struct {
	// The information about the databases on which hot-spot scale-out is performed.
	List *DescribeHotDbListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// The random number.
	RandomCode *string `json:"RandomCode,omitempty" xml:"RandomCode,omitempty"`
}

func (s DescribeHotDbListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBodyData) SetList(v *DescribeHotDbListResponseBodyDataList) *DescribeHotDbListResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeHotDbListResponseBodyData) SetRandomCode(v string) *DescribeHotDbListResponseBodyData {
	s.RandomCode = &v
	return s
}

type DescribeHotDbListResponseBodyDataList struct {
	InstanceDb []*DescribeHotDbListResponseBodyDataListInstanceDb `json:"InstanceDb,omitempty" xml:"InstanceDb,omitempty" type:"Repeated"`
}

func (s DescribeHotDbListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBodyDataList) SetInstanceDb(v []*DescribeHotDbListResponseBodyDataListInstanceDb) *DescribeHotDbListResponseBodyDataList {
	s.InstanceDb = v
	return s
}

type DescribeHotDbListResponseBodyDataListInstanceDb struct {
	HotDbList *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList `json:"HotDbList,omitempty" xml:"HotDbList,omitempty" type:"Struct"`
	// The name of the instance.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DescribeHotDbListResponseBodyDataListInstanceDb) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponseBodyDataListInstanceDb) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDb) SetHotDbList(v *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) *DescribeHotDbListResponseBodyDataListInstanceDb {
	s.HotDbList = v
	return s
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDb) SetInstanceName(v string) *DescribeHotDbListResponseBodyDataListInstanceDb {
	s.InstanceName = &v
	return s
}

type DescribeHotDbListResponseBodyDataListInstanceDbHotDbList struct {
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) SetData(v []*string) *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList {
	s.Data = v
	return s
}

type DescribeHotDbListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHotDbListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHotDbListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponse) SetHeaders(v map[string]*string) *DescribeHotDbListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHotDbListResponse) SetStatusCode(v int32) *DescribeHotDbListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHotDbListResponse) SetBody(v *DescribeHotDbListResponseBody) *DescribeHotDbListResponse {
	s.Body = v
	return s
}

type DescribeInstDbLogInfoRequest struct {
	// The name of the DRDS database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstDbLogInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbLogInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoRequest) SetDbName(v string) *DescribeInstDbLogInfoRequest {
	s.DbName = &v
	return s
}

func (s *DescribeInstDbLogInfoRequest) SetDrdsInstanceId(v string) *DescribeInstDbLogInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeInstDbLogInfoResponseBody struct {
	// The time range for log query.
	LogTimeRange *DescribeInstDbLogInfoResponseBodyLogTimeRange `json:"LogTimeRange,omitempty" xml:"LogTimeRange,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstDbLogInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbLogInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoResponseBody) SetLogTimeRange(v *DescribeInstDbLogInfoResponseBodyLogTimeRange) *DescribeInstDbLogInfoResponseBody {
	s.LogTimeRange = v
	return s
}

func (s *DescribeInstDbLogInfoResponseBody) SetRequestId(v string) *DescribeInstDbLogInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstDbLogInfoResponseBody) SetSuccess(v bool) *DescribeInstDbLogInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeInstDbLogInfoResponseBodyLogTimeRange struct {
	// The start time of the query time range.
	SupportLatestTime *int64 `json:"SupportLatestTime,omitempty" xml:"SupportLatestTime,omitempty"`
	// The end time of the task.
	SupportOldestTime *int64 `json:"SupportOldestTime,omitempty" xml:"SupportOldestTime,omitempty"`
}

func (s DescribeInstDbLogInfoResponseBodyLogTimeRange) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbLogInfoResponseBodyLogTimeRange) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoResponseBodyLogTimeRange) SetSupportLatestTime(v int64) *DescribeInstDbLogInfoResponseBodyLogTimeRange {
	s.SupportLatestTime = &v
	return s
}

func (s *DescribeInstDbLogInfoResponseBodyLogTimeRange) SetSupportOldestTime(v int64) *DescribeInstDbLogInfoResponseBodyLogTimeRange {
	s.SupportOldestTime = &v
	return s
}

type DescribeInstDbLogInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstDbLogInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstDbLogInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbLogInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoResponse) SetHeaders(v map[string]*string) *DescribeInstDbLogInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstDbLogInfoResponse) SetStatusCode(v int32) *DescribeInstDbLogInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstDbLogInfoResponse) SetBody(v *DescribeInstDbLogInfoResponseBody) *DescribeInstDbLogInfoResponse {
	s.Body = v
	return s
}

type DescribeInstDbSlsInfoRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstDbSlsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbSlsInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstDbSlsInfoRequest) SetDbName(v string) *DescribeInstDbSlsInfoRequest {
	s.DbName = &v
	return s
}

func (s *DescribeInstDbSlsInfoRequest) SetDrdsInstanceId(v string) *DescribeInstDbSlsInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeInstDbSlsInfoResponseBody struct {
	// The details of the SQL audit.
	AuditInfo *DescribeInstDbSlsInfoResponseBodyAuditInfo `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstDbSlsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbSlsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstDbSlsInfoResponseBody) SetAuditInfo(v *DescribeInstDbSlsInfoResponseBodyAuditInfo) *DescribeInstDbSlsInfoResponseBody {
	s.AuditInfo = v
	return s
}

func (s *DescribeInstDbSlsInfoResponseBody) SetRequestId(v string) *DescribeInstDbSlsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstDbSlsInfoResponseBody) SetSuccess(v bool) *DescribeInstDbSlsInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeInstDbSlsInfoResponseBodyAuditInfo struct {
	// The name of the LogStore.
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The name of the Log Service project.
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DescribeInstDbSlsInfoResponseBodyAuditInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbSlsInfoResponseBodyAuditInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstDbSlsInfoResponseBodyAuditInfo) SetLogStore(v string) *DescribeInstDbSlsInfoResponseBodyAuditInfo {
	s.LogStore = &v
	return s
}

func (s *DescribeInstDbSlsInfoResponseBodyAuditInfo) SetProject(v string) *DescribeInstDbSlsInfoResponseBodyAuditInfo {
	s.Project = &v
	return s
}

type DescribeInstDbSlsInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstDbSlsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstDbSlsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbSlsInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstDbSlsInfoResponse) SetHeaders(v map[string]*string) *DescribeInstDbSlsInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstDbSlsInfoResponse) SetStatusCode(v int32) *DescribeInstDbSlsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstDbSlsInfoResponse) SetBody(v *DescribeInstDbSlsInfoResponseBody) *DescribeInstDbSlsInfoResponse {
	s.Body = v
	return s
}

type DescribeInstanceAccountsRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstanceAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsRequest) SetDrdsInstanceId(v string) *DescribeInstanceAccountsRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeInstanceAccountsResponseBody struct {
	// Indicates the information about the instance accounts.
	InstanceAccounts *DescribeInstanceAccountsResponseBodyInstanceAccounts `json:"InstanceAccounts,omitempty" xml:"InstanceAccounts,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstanceAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBody) SetInstanceAccounts(v *DescribeInstanceAccountsResponseBodyInstanceAccounts) *DescribeInstanceAccountsResponseBody {
	s.InstanceAccounts = v
	return s
}

func (s *DescribeInstanceAccountsResponseBody) SetRequestId(v string) *DescribeInstanceAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBody) SetSuccess(v bool) *DescribeInstanceAccountsResponseBody {
	s.Success = &v
	return s
}

type DescribeInstanceAccountsResponseBodyInstanceAccounts struct {
	InstanceAccount []*DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount `json:"InstanceAccount,omitempty" xml:"InstanceAccount,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccounts) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccounts) SetInstanceAccount(v []*DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) *DescribeInstanceAccountsResponseBodyInstanceAccounts {
	s.InstanceAccount = v
	return s
}

type DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount struct {
	// Indicates the username of an instance account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Indicates the type of an instance account. Valid values:
	//
	// *   **0**: The instance account is a privileged account.
	// *   **1**: The instance account is a standard account.
	AccountType *int32 `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// Indicates the information about the permissions of an account on a database.
	DbPrivileges *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges `json:"DbPrivileges,omitempty" xml:"DbPrivileges,omitempty" type:"Struct"`
	// Indicates the description of an account. By default, if 0 is the value of the AccountType parameter, **Created by DRDS** is returned as the value of the Description parameter. If 1 is the value of the AccountType parameter, an empty string is returned as the value of the Description parameter. You can modify the description of an account on the Accounts page in the PolarDB-X console.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates an IP address that is allowed to access the database. The value **%** indicates that each IP address is allowed to access the database. \</note>
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetAccountName(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetAccountType(v int32) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.AccountType = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetDbPrivileges(v *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.DbPrivileges = v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetDescription(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetHost(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.Host = &v
	return s
}

type DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges struct {
	DbPrivilege []*DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege `json:"DbPrivilege,omitempty" xml:"DbPrivilege,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) SetDbPrivilege(v []*DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges {
	s.DbPrivilege = v
	return s
}

type DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege struct {
	// Indicates the name of a database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates the permissions that an account is granted on the database. Valid values:
	//
	// *   **R**: The account is granted the permissions that are required to read the data of the database.
	// *   **W**: The account is granted the permissions that are required to write data to the database.
	// *   **DDL**: The account is granted the permissions that are required to perform DDL operations on the database.
	// *   **DML**: The account is granted the permissions that are required to perform DML operations on the database.
	Privilege *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) SetDbName(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege {
	s.DbName = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) SetPrivilege(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege {
	s.Privilege = &v
	return s
}

type DescribeInstanceAccountsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponse) SetHeaders(v map[string]*string) *DescribeInstanceAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAccountsResponse) SetStatusCode(v int32) *DescribeInstanceAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceAccountsResponse) SetBody(v *DescribeInstanceAccountsResponseBody) *DescribeInstanceAccountsResponse {
	s.Body = v
	return s
}

type DescribeInstanceSwitchAzoneRequest struct {
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstanceSwitchAzoneRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneRequest) SetDrdsInstanceId(v string) *DescribeInstanceSwitchAzoneRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeInstanceSwitchAzoneResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the operation.
	Result *DescribeInstanceSwitchAzoneResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstanceSwitchAzoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneResponseBody) SetRequestId(v string) *DescribeInstanceSwitchAzoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBody) SetResult(v *DescribeInstanceSwitchAzoneResponseBodyResult) *DescribeInstanceSwitchAzoneResponseBody {
	s.Result = v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBody) SetSuccess(v bool) *DescribeInstanceSwitchAzoneResponseBody {
	s.Success = &v
	return s
}

type DescribeInstanceSwitchAzoneResponseBodyResult struct {
	// The ID of the source azoneId.
	OriginAzoneId *string `json:"OriginAzoneId,omitempty" xml:"OriginAzoneId,omitempty"`
	// regionId.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates whether the job can be switched.
	SwitchAble *bool `json:"SwitchAble,omitempty" xml:"SwitchAble,omitempty"`
	// Target azones.
	TargetAzones *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones `json:"TargetAzones,omitempty" xml:"TargetAzones,omitempty" type:"Struct"`
}

func (s DescribeInstanceSwitchAzoneResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) SetOriginAzoneId(v string) *DescribeInstanceSwitchAzoneResponseBodyResult {
	s.OriginAzoneId = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) SetRegionId(v string) *DescribeInstanceSwitchAzoneResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) SetSwitchAble(v bool) *DescribeInstanceSwitchAzoneResponseBodyResult {
	s.SwitchAble = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) SetTargetAzones(v *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones) *DescribeInstanceSwitchAzoneResponseBodyResult {
	s.TargetAzones = v
	return s
}

type DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones struct {
	TargetAzone []*string `json:"TargetAzone,omitempty" xml:"TargetAzone,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones) SetTargetAzone(v []*string) *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones {
	s.TargetAzone = v
	return s
}

type DescribeInstanceSwitchAzoneResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSwitchAzoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceSwitchAzoneResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneResponse) SetHeaders(v map[string]*string) *DescribeInstanceSwitchAzoneResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponse) SetStatusCode(v int32) *DescribeInstanceSwitchAzoneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponse) SetBody(v *DescribeInstanceSwitchAzoneResponseBody) *DescribeInstanceSwitchAzoneResponse {
	s.Body = v
	return s
}

type DescribeInstanceSwitchNetworkRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstanceSwitchNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkRequest) SetDrdsInstanceId(v string) *DescribeInstanceSwitchNetworkRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeInstanceSwitchNetworkResponseBody struct {
	// Indicates the ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates the information about the virtual private cloud (VPC) in which the instance is deployed.
	VpcInfos *DescribeInstanceSwitchNetworkResponseBodyVpcInfos `json:"VpcInfos,omitempty" xml:"VpcInfos,omitempty" type:"Struct"`
}

func (s DescribeInstanceSwitchNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBody) SetRequestId(v string) *DescribeInstanceSwitchNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBody) SetSuccess(v bool) *DescribeInstanceSwitchNetworkResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBody) SetVpcInfos(v *DescribeInstanceSwitchNetworkResponseBodyVpcInfos) *DescribeInstanceSwitchNetworkResponseBody {
	s.VpcInfos = v
	return s
}

type DescribeInstanceSwitchNetworkResponseBodyVpcInfos struct {
	VpcInfo []*DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo `json:"VpcInfo,omitempty" xml:"VpcInfo,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfos) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfos) SetVpcInfo(v []*DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) *DescribeInstanceSwitchNetworkResponseBodyVpcInfos {
	s.VpcInfo = v
	return s
}

type DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo struct {
	// Indicates the ID of the region in which the instance is deployed.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates the ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Indicates the name of the VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// Indicates information about the vSwitch to which the instance is connected.
	VswitchInfos *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos `json:"VswitchInfos,omitempty" xml:"VswitchInfos,omitempty" type:"Struct"`
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetRegionId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetVpcId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetVpcName(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.VpcName = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetVswitchInfos(v *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.VswitchInfos = v
	return s
}

type DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos struct {
	VswitchInfo []*DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo `json:"VswitchInfo,omitempty" xml:"VswitchInfo,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos) SetVswitchInfo(v []*DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos {
	s.VswitchInfo = v
	return s
}

type DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo struct {
	// Indicates the ID of the zone in which the instance is deployed.
	AzoneId *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	// Indicates whether you can change the network type of the instance.
	DrdsSupported *bool `json:"DrdsSupported,omitempty" xml:"DrdsSupported,omitempty"`
	// Indicates the ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Indicates the ID of the vSwitch.
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// Indicates the name of the vSwitch.
	VswitchName *string `json:"VswitchName,omitempty" xml:"VswitchName,omitempty"`
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetAzoneId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.AzoneId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetDrdsSupported(v bool) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.DrdsSupported = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetVpcId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetVswitchId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetVswitchName(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.VswitchName = &v
	return s
}

type DescribeInstanceSwitchNetworkResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSwitchNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceSwitchNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponse) SetHeaders(v map[string]*string) *DescribeInstanceSwitchNetworkResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponse) SetStatusCode(v int32) *DescribeInstanceSwitchNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponse) SetBody(v *DescribeInstanceSwitchNetworkResponseBody) *DescribeInstanceSwitchNetworkResponse {
	s.Body = v
	return s
}

type DescribePreCheckResultRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the precheck task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribePreCheckResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckResultRequest) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultRequest) SetDrdsInstanceId(v string) *DescribePreCheckResultRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribePreCheckResultRequest) SetRegionId(v string) *DescribePreCheckResultRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePreCheckResultRequest) SetTaskId(v string) *DescribePreCheckResultRequest {
	s.TaskId = &v
	return s
}

type DescribePreCheckResultResponseBody struct {
	// Indicates the result of the precheck task.
	PreCheckResult *DescribePreCheckResultResponseBodyPreCheckResult `json:"PreCheckResult,omitempty" xml:"PreCheckResult,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePreCheckResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultResponseBody) SetPreCheckResult(v *DescribePreCheckResultResponseBodyPreCheckResult) *DescribePreCheckResultResponseBody {
	s.PreCheckResult = v
	return s
}

func (s *DescribePreCheckResultResponseBody) SetRequestId(v string) *DescribePreCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePreCheckResultResponseBody) SetSuccess(v bool) *DescribePreCheckResultResponseBody {
	s.Success = &v
	return s
}

type DescribePreCheckResultResponseBodyPreCheckResult struct {
	// Indicates the name of the precheck task.
	PreCheckName *string `json:"PreCheckName,omitempty" xml:"PreCheckName,omitempty"`
	// Indicates the state of the precheck task.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Indicates the details about the subtasks of the precheck task.
	SubCheckItems []*DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems `json:"SubCheckItems,omitempty" xml:"SubCheckItems,omitempty" type:"Repeated"`
}

func (s DescribePreCheckResultResponseBodyPreCheckResult) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckResultResponseBodyPreCheckResult) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultResponseBodyPreCheckResult) SetPreCheckName(v string) *DescribePreCheckResultResponseBodyPreCheckResult {
	s.PreCheckName = &v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResult) SetState(v string) *DescribePreCheckResultResponseBodyPreCheckResult {
	s.State = &v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResult) SetSubCheckItems(v []*DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) *DescribePreCheckResultResponseBodyPreCheckResult {
	s.SubCheckItems = v
	return s
}

type DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems struct {
	// Indicates the error code that is returned by a subtask.
	ErrorMsgCode *string `json:"ErrorMsgCode,omitempty" xml:"ErrorMsgCode,omitempty"`
	// Indicates an error message.
	ErrorMsgParams []*string `json:"ErrorMsgParams,omitempty" xml:"ErrorMsgParams,omitempty" type:"Repeated"`
	// Indicates the name of the subtask.
	PreCheckItemName *string `json:"PreCheckItemName,omitempty" xml:"PreCheckItemName,omitempty"`
	// Indicates the state of the subtask.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) SetErrorMsgCode(v string) *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems {
	s.ErrorMsgCode = &v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) SetErrorMsgParams(v []*string) *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems {
	s.ErrorMsgParams = v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) SetPreCheckItemName(v string) *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems {
	s.PreCheckItemName = &v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) SetState(v string) *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems {
	s.State = &v
	return s
}

type DescribePreCheckResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePreCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePreCheckResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckResultResponse) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultResponse) SetHeaders(v map[string]*string) *DescribePreCheckResultResponse {
	s.Headers = v
	return s
}

func (s *DescribePreCheckResultResponse) SetStatusCode(v int32) *DescribePreCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePreCheckResultResponse) SetBody(v *DescribePreCheckResultResponseBody) *DescribePreCheckResultResponse {
	s.Body = v
	return s
}

type DescribeRDSPerformanceRequest struct {
	// The type of the database engine.
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The ID of the Distributed Relational Database Service (DRDS) instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The end time of the query. Specify the time in the UNIX timestamp format. The time must be in UTC. Unit: ms.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance monitoring metrics. You can specify one or more metrics for a query at a time. Separate multiple metric parameters with commas (,).
	//
	// >  For more information about the details of performance monitoring metrics, see [Storage monitoring](~~186705~~).
	Keys *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	// The ID of the storage-layer ApsaraDB RDS for MySQL instance.
	RdsInstanceId *string `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	// The start time of the query. Specify the time in the UNIX timestamp format. The time must be in UTC. Unit: ms.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRDSPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceRequest) SetDbInstType(v string) *DescribeRDSPerformanceRequest {
	s.DbInstType = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetDrdsInstanceId(v string) *DescribeRDSPerformanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetEndTime(v int64) *DescribeRDSPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetKeys(v string) *DescribeRDSPerformanceRequest {
	s.Keys = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetRdsInstanceId(v string) *DescribeRDSPerformanceRequest {
	s.RdsInstanceId = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetStartTime(v int64) *DescribeRDSPerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeRDSPerformanceResponseBody struct {
	// The result set of the query.
	Data []*DescribeRDSPerformanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRDSPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceResponseBody) SetData(v []*DescribeRDSPerformanceResponseBodyData) *DescribeRDSPerformanceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRDSPerformanceResponseBody) SetRequestId(v string) *DescribeRDSPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBody) SetSuccess(v bool) *DescribeRDSPerformanceResponseBody {
	s.Success = &v
	return s
}

type DescribeRDSPerformanceResponseBodyData struct {
	// The name of the monitoring metric.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the node.
	//
	// >  This parameter is returned only when the storage type of the database is PolarDB for MySQL. If the storage type of the database is ApsaraDB RDS for MySQL, this parameter is not returned.
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The number of nodes.
	NodeNum *int32 `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// The unit of the monitoring metric.
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The details of the monitoring metric data.
	Values []*DescribeRDSPerformanceResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeRDSPerformanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSPerformanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceResponseBodyData) SetKey(v string) *DescribeRDSPerformanceResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyData) SetNodeName(v string) *DescribeRDSPerformanceResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyData) SetNodeNum(v int32) *DescribeRDSPerformanceResponseBodyData {
	s.NodeNum = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyData) SetUnit(v string) *DescribeRDSPerformanceResponseBodyData {
	s.Unit = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyData) SetValues(v []*DescribeRDSPerformanceResponseBodyDataValues) *DescribeRDSPerformanceResponseBodyData {
	s.Values = v
	return s
}

type DescribeRDSPerformanceResponseBodyDataValues struct {
	// The time point when the value of the monitoring metric was obtained. The value is in the UNIX timestamp format. The time is displayed in UTC. Unit: ms.
	Date *int64 `json:"Date,omitempty" xml:"Date,omitempty"`
	// The value of the monitoring metric.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRDSPerformanceResponseBodyDataValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSPerformanceResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceResponseBodyDataValues) SetDate(v int64) *DescribeRDSPerformanceResponseBodyDataValues {
	s.Date = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyDataValues) SetValue(v string) *DescribeRDSPerformanceResponseBodyDataValues {
	s.Value = &v
	return s
}

type DescribeRDSPerformanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRDSPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRDSPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceResponse) SetHeaders(v map[string]*string) *DescribeRDSPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRDSPerformanceResponse) SetStatusCode(v int32) *DescribeRDSPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRDSPerformanceResponse) SetBody(v *DescribeRDSPerformanceResponseBody) *DescribeRDSPerformanceResponse {
	s.Body = v
	return s
}

type DescribeRdsCommodityRequest struct {
	// The commodity code of the service.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The type of the order.
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DescribeRdsCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsCommodityRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsCommodityRequest) SetCommodityCode(v string) *DescribeRdsCommodityRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeRdsCommodityRequest) SetDrdsInstanceId(v string) *DescribeRdsCommodityRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRdsCommodityRequest) SetOrderType(v string) *DescribeRdsCommodityRequest {
	s.OrderType = &v
	return s
}

type DescribeRdsCommodityResponseBody struct {
	// Indicates the returned result.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Indicates the ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRdsCommodityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsCommodityResponseBody) SetData(v string) *DescribeRdsCommodityResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeRdsCommodityResponseBody) SetRequestId(v string) *DescribeRdsCommodityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsCommodityResponseBody) SetSuccess(v bool) *DescribeRdsCommodityResponseBody {
	s.Success = &v
	return s
}

type DescribeRdsCommodityResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdsCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdsCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsCommodityResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsCommodityResponse) SetHeaders(v map[string]*string) *DescribeRdsCommodityResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsCommodityResponse) SetStatusCode(v int32) *DescribeRdsCommodityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsCommodityResponse) SetBody(v *DescribeRdsCommodityResponseBody) *DescribeRdsCommodityResponse {
	s.Body = v
	return s
}

type DescribeRdsPerformanceSummaryRequest struct {
	// The ID of a DRDS instance.
	DrdsInstanceId *string   `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RdsInstanceId  []*string `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty" type:"Repeated"`
	// The ID of the region where the streaming domain resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRdsPerformanceSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryRequest) SetDrdsInstanceId(v string) *DescribeRdsPerformanceSummaryRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryRequest) SetRdsInstanceId(v []*string) *DescribeRdsPerformanceSummaryRequest {
	s.RdsInstanceId = v
	return s
}

func (s *DescribeRdsPerformanceSummaryRequest) SetRegionId(v string) *DescribeRdsPerformanceSummaryRequest {
	s.RegionId = &v
	return s
}

type DescribeRdsPerformanceSummaryResponseBody struct {
	// A collection of objects.
	RdsPerformanceInfos []*DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos `json:"RdsPerformanceInfos,omitempty" xml:"RdsPerformanceInfos,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRdsPerformanceSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryResponseBody) SetRdsPerformanceInfos(v []*DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) *DescribeRdsPerformanceSummaryResponseBody {
	s.RdsPerformanceInfos = v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBody) SetRequestId(v string) *DescribeRdsPerformanceSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBody) SetSuccess(v bool) *DescribeRdsPerformanceSummaryResponseBody {
	s.Success = &v
	return s
}

type DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos struct {
	// The number of active sessions of the RDS instance.
	ActiveSessions *int32 `json:"ActiveSessions,omitempty" xml:"ActiveSessions,omitempty"`
	// The CPU utilization of an RDS instance.
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The IOPS of the RDS instance.
	Iops *float32 `json:"Iops,omitempty" xml:"Iops,omitempty"`
	// The ID of an RDS instance.
	RdsId *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	// The disk usage of apsaradb for RDS. Unit: MB.
	SpaceUsage *int64 `json:"SpaceUsage,omitempty" xml:"SpaceUsage,omitempty"`
	// The total number of current RDS sessions.
	TotalSessions *int32 `json:"TotalSessions,omitempty" xml:"TotalSessions,omitempty"`
}

func (s DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetActiveSessions(v int32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.ActiveSessions = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetCpu(v float32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.Cpu = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetIops(v float32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.Iops = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetRdsId(v string) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.RdsId = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetSpaceUsage(v int64) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.SpaceUsage = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetTotalSessions(v int32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.TotalSessions = &v
	return s
}

type DescribeRdsPerformanceSummaryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdsPerformanceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdsPerformanceSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryResponse) SetHeaders(v map[string]*string) *DescribeRdsPerformanceSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponse) SetStatusCode(v int32) *DescribeRdsPerformanceSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponse) SetBody(v *DescribeRdsPerformanceSummaryResponseBody) *DescribeRdsPerformanceSummaryResponse {
	s.Body = v
	return s
}

type DescribeRdsSuperAccountInstancesRequest struct {
	// The type of the ApsaraDB RDS for MySQL instances. Default value: **RDS**.
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string   `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RdsInstance    []*string `json:"RdsInstance,omitempty" xml:"RdsInstance,omitempty" type:"Repeated"`
}

func (s DescribeRdsSuperAccountInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsSuperAccountInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsSuperAccountInstancesRequest) SetDbInstType(v string) *DescribeRdsSuperAccountInstancesRequest {
	s.DbInstType = &v
	return s
}

func (s *DescribeRdsSuperAccountInstancesRequest) SetDrdsInstanceId(v string) *DescribeRdsSuperAccountInstancesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRdsSuperAccountInstancesRequest) SetRdsInstance(v []*string) *DescribeRdsSuperAccountInstancesRequest {
	s.RdsInstance = v
	return s
}

type DescribeRdsSuperAccountInstancesResponseBody struct {
	// The privileged accounts.
	DbInstances *DescribeRdsSuperAccountInstancesResponseBodyDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRdsSuperAccountInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsSuperAccountInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsSuperAccountInstancesResponseBody) SetDbInstances(v *DescribeRdsSuperAccountInstancesResponseBodyDbInstances) *DescribeRdsSuperAccountInstancesResponseBody {
	s.DbInstances = v
	return s
}

func (s *DescribeRdsSuperAccountInstancesResponseBody) SetRequestId(v string) *DescribeRdsSuperAccountInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRdsSuperAccountInstancesResponseBodyDbInstances struct {
	DbInstance []*string `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Repeated"`
}

func (s DescribeRdsSuperAccountInstancesResponseBodyDbInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsSuperAccountInstancesResponseBodyDbInstances) GoString() string {
	return s.String()
}

func (s *DescribeRdsSuperAccountInstancesResponseBodyDbInstances) SetDbInstance(v []*string) *DescribeRdsSuperAccountInstancesResponseBodyDbInstances {
	s.DbInstance = v
	return s
}

type DescribeRdsSuperAccountInstancesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdsSuperAccountInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdsSuperAccountInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsSuperAccountInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsSuperAccountInstancesResponse) SetHeaders(v map[string]*string) *DescribeRdsSuperAccountInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsSuperAccountInstancesResponse) SetStatusCode(v int32) *DescribeRdsSuperAccountInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsSuperAccountInstancesResponse) SetBody(v *DescribeRdsSuperAccountInstancesResponseBody) *DescribeRdsSuperAccountInstancesResponse {
	s.Body = v
	return s
}

type DescribeRecycleBinStatusRequest struct {
	// The name of the database that is created in the PolarDB-X 1.0 instance.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRecycleBinStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecycleBinStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinStatusRequest) SetDbName(v string) *DescribeRecycleBinStatusRequest {
	s.DbName = &v
	return s
}

func (s *DescribeRecycleBinStatusRequest) SetDrdsInstanceId(v string) *DescribeRecycleBinStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRecycleBinStatusRequest) SetRegionId(v string) *DescribeRecycleBinStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeRecycleBinStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the table recycle bin. Valid values:
	//
	// *   disable: The table recycle bin is enabled.
	// *   enable: The table recycle bin is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The result of the request.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRecycleBinStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecycleBinStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinStatusResponseBody) SetRequestId(v string) *DescribeRecycleBinStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecycleBinStatusResponseBody) SetStatus(v string) *DescribeRecycleBinStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeRecycleBinStatusResponseBody) SetSuccess(v bool) *DescribeRecycleBinStatusResponseBody {
	s.Success = &v
	return s
}

type DescribeRecycleBinStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecycleBinStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecycleBinStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecycleBinStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinStatusResponse) SetHeaders(v map[string]*string) *DescribeRecycleBinStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecycleBinStatusResponse) SetStatusCode(v int32) *DescribeRecycleBinStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecycleBinStatusResponse) SetBody(v *DescribeRecycleBinStatusResponseBody) *DescribeRecycleBinStatusResponse {
	s.Body = v
	return s
}

type DescribeRecycleBinTablesRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRecycleBinTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecycleBinTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinTablesRequest) SetDbName(v string) *DescribeRecycleBinTablesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeRecycleBinTablesRequest) SetDrdsInstanceId(v string) *DescribeRecycleBinTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRecycleBinTablesRequest) SetRegionId(v string) *DescribeRecycleBinTablesRequest {
	s.RegionId = &v
	return s
}

type DescribeRecycleBinTablesResponseBody struct {
	// The data object returned.
	Data []*DescribeRecycleBinTablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRecycleBinTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecycleBinTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinTablesResponseBody) SetData(v []*DescribeRecycleBinTablesResponseBodyData) *DescribeRecycleBinTablesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRecycleBinTablesResponseBody) SetRequestId(v string) *DescribeRecycleBinTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecycleBinTablesResponseBody) SetSuccess(v bool) *DescribeRecycleBinTablesResponseBody {
	s.Success = &v
	return s
}

type DescribeRecycleBinTablesResponseBodyData struct {
	// The time when the table was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The original name of the table.
	OriginalTableName *string `json:"OriginalTableName,omitempty" xml:"OriginalTableName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeRecycleBinTablesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecycleBinTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinTablesResponseBodyData) SetCreateTime(v string) *DescribeRecycleBinTablesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeRecycleBinTablesResponseBodyData) SetOriginalTableName(v string) *DescribeRecycleBinTablesResponseBodyData {
	s.OriginalTableName = &v
	return s
}

func (s *DescribeRecycleBinTablesResponseBodyData) SetTableName(v string) *DescribeRecycleBinTablesResponseBodyData {
	s.TableName = &v
	return s
}

type DescribeRecycleBinTablesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecycleBinTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecycleBinTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecycleBinTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinTablesResponse) SetHeaders(v map[string]*string) *DescribeRecycleBinTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecycleBinTablesResponse) SetStatusCode(v int32) *DescribeRecycleBinTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecycleBinTablesResponse) SetBody(v *DescribeRecycleBinTablesResponseBody) *DescribeRecycleBinTablesResponse {
	s.Body = v
	return s
}

type DescribeRestoreOrderRequest struct {
	// The name of the database involved in the backup.
	BackupDbNames *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	// The ID of the backup set.
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The level of the backup. Valid values:
	//
	// *   **DB**: The database Level
	// *   **instance **: instance level
	BackupLevel *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	// The backup mode. Valid values: **logic** or **phy**.
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The ID of the instance for which to modify the backup policy.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The preferred backup time.
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
}

func (s DescribeRestoreOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderRequest) SetBackupDbNames(v string) *DescribeRestoreOrderRequest {
	s.BackupDbNames = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetBackupId(v string) *DescribeRestoreOrderRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetBackupLevel(v string) *DescribeRestoreOrderRequest {
	s.BackupLevel = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetBackupMode(v string) *DescribeRestoreOrderRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetDrdsInstanceId(v string) *DescribeRestoreOrderRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetPreferredBackupTime(v string) *DescribeRestoreOrderRequest {
	s.PreferredBackupTime = &v
	return s
}

type DescribeRestoreOrderResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned data object.
	RestoreOrderDO *DescribeRestoreOrderResponseBodyRestoreOrderDO `json:"RestoreOrderDO,omitempty" xml:"RestoreOrderDO,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRestoreOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBody) SetRequestId(v string) *DescribeRestoreOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBody) SetRestoreOrderDO(v *DescribeRestoreOrderResponseBodyRestoreOrderDO) *DescribeRestoreOrderResponseBody {
	s.RestoreOrderDO = v
	return s
}

func (s *DescribeRestoreOrderResponseBody) SetSuccess(v bool) *DescribeRestoreOrderResponseBody {
	s.Success = &v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDO struct {
	// The information of the restored order.
	DrdsOrderDOList *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList `json:"DrdsOrderDOList,omitempty" xml:"DrdsOrderDOList,omitempty" type:"Struct"`
	// The ID of the restored apsaradb for PolarDB cluster.
	PolarOrderDOList *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList `json:"PolarOrderDOList,omitempty" xml:"PolarOrderDOList,omitempty" type:"Struct"`
	// The information of the restored RDS order.
	RdsOrderDOList *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList `json:"RdsOrderDOList,omitempty" xml:"RdsOrderDOList,omitempty" type:"Struct"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDO) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDO) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) SetDrdsOrderDOList(v *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDO {
	s.DrdsOrderDOList = v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) SetPolarOrderDOList(v *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDO {
	s.PolarOrderDOList = v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) SetRdsOrderDOList(v *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDO {
	s.RdsOrderDOList = v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList struct {
	DrdsOrderDOList []*DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList `json:"DrdsOrderDOList,omitempty" xml:"DrdsOrderDOList,omitempty" type:"Repeated"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList) SetDrdsOrderDOList(v []*DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList {
	s.DrdsOrderDOList = v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList struct {
	// The ID of the zone for which to query resources.
	AzoneId *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	// The instance type of the instance.
	InstSpec *string `json:"InstSpec,omitempty" xml:"InstSpec,omitempty"`
	// The network type. Valid values:
	//
	// *   **Classic **: Classic Network
	// *   **vpc**: VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the vSwitch in the VPC.
	VSwtichId *string `json:"VSwtichId,omitempty" xml:"VSwtichId,omitempty"`
	// The ID of the VPC network.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetAzoneId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.AzoneId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetInstSpec(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.InstSpec = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetNetwork(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.Network = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetRegionId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.RegionId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetVSwtichId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.VSwtichId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetVpcId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.VpcId = &v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList struct {
	PolarOrderDOList []*DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList `json:"PolarOrderDOList,omitempty" xml:"PolarOrderDOList,omitempty" type:"Repeated"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) SetPolarOrderDOList(v []*DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList {
	s.PolarOrderDOList = v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList struct {
	// The zone ID of the node.
	AzoneId *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	// The capacity of disk.
	DbInstanceStorage *string `json:"DbInstanceStorage,omitempty" xml:"DbInstanceStorage,omitempty"`
	// The storage engine of PolarDB.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The type of the instance.
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The network type. Valid values:
	//
	// *   **Classic**: Classic Network
	// *   **vpc**: VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The number of streams that were returned.
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The version of the operating system.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetAzoneId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.AzoneId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetDbInstanceStorage(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.DbInstanceStorage = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetEngine(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Engine = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetInstanceClass(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.InstanceClass = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetNetwork(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Network = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetNum(v int64) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Num = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetRegionId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.RegionId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetVersion(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Version = &v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList struct {
	RdsOrderDOList []*DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList `json:"RdsOrderDOList,omitempty" xml:"RdsOrderDOList,omitempty" type:"Repeated"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) SetRdsOrderDOList(v []*DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList {
	s.RdsOrderDOList = v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList struct {
	// The zone ID of the node.
	AzoneId *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	// The capacity of disk.
	DbInstanceStorage *string `json:"DbInstanceStorage,omitempty" xml:"DbInstanceStorage,omitempty"`
	// The storage engine of the instance.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The instance type of the instance.
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The network type. Valid values: - **Classic **: Classic Network
	// - **vpc**: VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The number of streams that were returned.
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The version of the operating system.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetAzoneId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.AzoneId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetDbInstanceStorage(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.DbInstanceStorage = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetEngine(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Engine = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetInstanceClass(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.InstanceClass = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetNetwork(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Network = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetNum(v int64) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Num = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetRegionId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.RegionId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetVersion(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Version = &v
	return s
}

type DescribeRestoreOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponse) SetHeaders(v map[string]*string) *DescribeRestoreOrderResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreOrderResponse) SetStatusCode(v int32) *DescribeRestoreOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreOrderResponse) SetBody(v *DescribeRestoreOrderResponseBody) *DescribeRestoreOrderResponse {
	s.Body = v
	return s
}

type DescribeShardTaskInfoRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the table that you want to convert or shard.
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	// The name of the table that is generated after you convert or shard the table.
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
}

func (s DescribeShardTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoRequest) SetDbName(v string) *DescribeShardTaskInfoRequest {
	s.DbName = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) SetDrdsInstanceId(v string) *DescribeShardTaskInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) SetRegionId(v string) *DescribeShardTaskInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) SetSourceTableName(v string) *DescribeShardTaskInfoRequest {
	s.SourceTableName = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) SetTargetTableName(v string) *DescribeShardTaskInfoRequest {
	s.TargetTableName = &v
	return s
}

type DescribeShardTaskInfoResponseBody struct {
	// Indicates the data that is returned.
	Data *DescribeShardTaskInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates the unique ID of the request. If the request fails, provide this ID for technical support to troubleshoot the failure.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeShardTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBody) SetData(v *DescribeShardTaskInfoResponseBodyData) *DescribeShardTaskInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeShardTaskInfoResponseBody) SetRequestId(v string) *DescribeShardTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBody) SetSuccess(v bool) *DescribeShardTaskInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeShardTaskInfoResponseBodyData struct {
	// Indicates the number of remaining days before the tasks to shard tables or convert tables expire.
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Indicates information about full migration tasks.
	Full *DescribeShardTaskInfoResponseBodyDataFull `json:"Full,omitempty" xml:"Full,omitempty" type:"Struct"`
	// Indicates information about full check tasks.
	FullCheck *DescribeShardTaskInfoResponseBodyDataFullCheck `json:"FullCheck,omitempty" xml:"FullCheck,omitempty" type:"Struct"`
	// Indicates information about full correction tasks.
	FullRevise *DescribeShardTaskInfoResponseBodyDataFullRevise `json:"FullRevise,omitempty" xml:"FullRevise,omitempty" type:"Struct"`
	// Indicates information about incremental data synchronization.
	Increment *DescribeShardTaskInfoResponseBodyDataIncrement `json:"Increment,omitempty" xml:"Increment,omitempty" type:"Struct"`
	// Indicates the incremental data synchronization progress.
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Indicates check tasks.
	Review *DescribeShardTaskInfoResponseBodyDataReview `json:"Review,omitempty" xml:"Review,omitempty" type:"Struct"`
	// Indicates the name of the table that you convert or shard.
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	// Indicates the current stage of the task.
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// Indicates the state of the tasks to shard tables or convert tables.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates the name of the table after you convert or shard the table.
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyData) SetExpired(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetFull(v *DescribeShardTaskInfoResponseBodyDataFull) *DescribeShardTaskInfoResponseBodyData {
	s.Full = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetFullCheck(v *DescribeShardTaskInfoResponseBodyDataFullCheck) *DescribeShardTaskInfoResponseBodyData {
	s.FullCheck = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetFullRevise(v *DescribeShardTaskInfoResponseBodyDataFullRevise) *DescribeShardTaskInfoResponseBodyData {
	s.FullRevise = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetIncrement(v *DescribeShardTaskInfoResponseBodyDataIncrement) *DescribeShardTaskInfoResponseBodyData {
	s.Increment = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetProgress(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetReview(v *DescribeShardTaskInfoResponseBodyDataReview) *DescribeShardTaskInfoResponseBodyData {
	s.Review = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetSourceTableName(v string) *DescribeShardTaskInfoResponseBodyData {
	s.SourceTableName = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetStage(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Stage = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetStatus(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetTargetTableName(v string) *DescribeShardTaskInfoResponseBodyData {
	s.TargetTableName = &v
	return s
}

type DescribeShardTaskInfoResponseBodyDataFull struct {
	// Indicates the number of remaining days before the tasks expire.
	Expired *int32 `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Indicates the progress of the tasks.
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Indicates the start time when the tasks are performed.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates the number of tasks.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// Indicates the number of transactions processed by the database per second.
	Tps *int32 `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataFull) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataFull) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataFull {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Tps = &v
	return s
}

type DescribeShardTaskInfoResponseBodyDataFullCheck struct {
	// Indicates the number of remaining days before the tasks expire.
	Expired *int32 `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Indicates the progress of the tasks.
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Indicates the start time when the tasks are performed.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates the number of tasks.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// Indicates the number of transactions processed by the database per second.
	Tps *int32 `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataFullCheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataFullCheck) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Tps = &v
	return s
}

type DescribeShardTaskInfoResponseBodyDataFullRevise struct {
	// Indicates the number of remaining days before the tasks expire.
	Expired *int32 `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Indicates the progress of the tasks.
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Indicates the start time when the tasks are performed.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates the number of tasks.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// Indicates the number of transactions processed by the database per second.
	Tps *int32 `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataFullRevise) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataFullRevise) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Tps = &v
	return s
}

type DescribeShardTaskInfoResponseBodyDataIncrement struct {
	// Indicates the latency of the incremental data synchronization.
	Delay *int32 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// Indicates the start time when the incremental data synchronization is performed.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates the number of transactions processed by the database per second.
	Tps *int32 `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataIncrement) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataIncrement) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataIncrement) SetDelay(v int32) *DescribeShardTaskInfoResponseBodyDataIncrement {
	s.Delay = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataIncrement) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataIncrement {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataIncrement) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataIncrement {
	s.Tps = &v
	return s
}

type DescribeShardTaskInfoResponseBodyDataReview struct {
	// Indicates the number of remaining days before the tasks expire.
	Expired *int32 `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Indicates the progress of the tasks.
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Indicates the start time when the tasks are performed.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates the number of tasks.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// Indicates the number of transactions processed by the database per second.
	Tps *int32 `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataReview) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataReview) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataReview {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Tps = &v
	return s
}

type DescribeShardTaskInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeShardTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeShardTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponse) SetHeaders(v map[string]*string) *DescribeShardTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeShardTaskInfoResponse) SetStatusCode(v int32) *DescribeShardTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeShardTaskInfoResponse) SetBody(v *DescribeShardTaskInfoResponseBody) *DescribeShardTaskInfoResponse {
	s.Body = v
	return s
}

type DescribeSqlFlashbakTaskRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeSqlFlashbakTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlFlashbakTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskRequest) SetDrdsInstanceId(v string) *DescribeSqlFlashbakTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeSqlFlashbakTaskResponseBody struct {
	// Indicates the ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates the information about flashback tasks.
	SqlFlashbackTasks *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks `json:"SqlFlashbackTasks,omitempty" xml:"SqlFlashbackTasks,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSqlFlashbakTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlFlashbakTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskResponseBody) SetRequestId(v string) *DescribeSqlFlashbakTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBody) SetSqlFlashbackTasks(v *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) *DescribeSqlFlashbakTaskResponseBody {
	s.SqlFlashbackTasks = v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBody) SetSuccess(v bool) *DescribeSqlFlashbakTaskResponseBody {
	s.Success = &v
	return s
}

type DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks struct {
	SqlFlashbackTask []*DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask `json:"SqlFlashbackTask,omitempty" xml:"SqlFlashbackTask,omitempty" type:"Repeated"`
}

func (s DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) SetSqlFlashbackTask(v []*DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks {
	s.SqlFlashbackTask = v
	return s
}

type DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask struct {
	// Indicates the name of the database on which a flashback task is performed.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates the download URL.
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// Indicates the time when the download URL expires.
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates the point in time when the instance was created.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Indicates the point in time when the flashback task is performed.
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates the ID of the primary key that corresponds to a table used in the flashback task.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates the ID of the instance.
	InstId *string `json:"InstId,omitempty" xml:"InstId,omitempty"`
	// Indicates the progress of the reverse call.
	RecallProgress *int32 `json:"RecallProgress,omitempty" xml:"RecallProgress,omitempty"`
	// Indicates the type of the flashback task. Valid values:
	//
	// *   **1**: image restoration
	// *   **2**: reverse restoration
	RecallRestoreType *int32 `json:"RecallRestoreType,omitempty" xml:"RecallRestoreType,omitempty"`
	// Indicates the status of the data recall task.
	RecallStatus *int32 `json:"RecallStatus,omitempty" xml:"RecallStatus,omitempty"`
	// Indicates the type of the reverse call. Valid values:
	//
	// *   **0**: exact search
	// *   **1**: fuzzy search
	RecallType *int32 `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// Indicates the start time of the reverse call.
	SearchEndTime *int64 `json:"SearchEndTime,omitempty" xml:"SearchEndTime,omitempty"`
	// Indicates the end time of the reverse call.
	SearchStartTime *int64 `json:"SearchStartTime,omitempty" xml:"SearchStartTime,omitempty"`
	// Indicates the number of data rows that are flashed back.
	SqlCounter *int64 `json:"SqlCounter,omitempty" xml:"SqlCounter,omitempty"`
	// Indicates the primary key specified in the SQL statements.
	SqlPk *string `json:"SqlPk,omitempty" xml:"SqlPk,omitempty"`
	// Indicates the types of the SQL statements.
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// Indicates the name of the table that contains the data that are flashed back.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// Indicates the ID of the trace of the SQL query.
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetDbName(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.DbName = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetDownloadUrl(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetExpireTime(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.ExpireTime = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetGmtCreate(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetGmtModified(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.GmtModified = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetId(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.Id = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetInstId(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.InstId = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallProgress(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallProgress = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallRestoreType(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallRestoreType = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallStatus(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallStatus = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallType(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallType = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSearchEndTime(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SearchEndTime = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSearchStartTime(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SearchStartTime = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSqlCounter(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SqlCounter = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSqlPk(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SqlPk = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSqlType(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SqlType = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetTableName(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.TableName = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetTraceId(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.TraceId = &v
	return s
}

type DescribeSqlFlashbakTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSqlFlashbakTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSqlFlashbakTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlFlashbakTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskResponse) SetHeaders(v map[string]*string) *DescribeSqlFlashbakTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlFlashbakTaskResponse) SetStatusCode(v int32) *DescribeSqlFlashbakTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponse) SetBody(v *DescribeSqlFlashbakTaskResponseBody) *DescribeSqlFlashbakTaskResponse {
	s.Body = v
	return s
}

type DescribeTableRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region where the PolarDB-X 1.0 instance is created.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableRequest) SetDbName(v string) *DescribeTableRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTableRequest) SetDrdsInstanceId(v string) *DescribeTableRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeTableRequest) SetRegionId(v string) *DescribeTableRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableRequest) SetTableName(v string) *DescribeTableRequest {
	s.TableName = &v
	return s
}

type DescribeTableResponseBody struct {
	// Indicates the returned data.
	Data *DescribeTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates the unique ID of the request. If the request fails, provide this ID for technical support to troubleshoot the failure.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableResponseBody) SetData(v *DescribeTableResponseBodyData) *DescribeTableResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTableResponseBody) SetRequestId(v string) *DescribeTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableResponseBody) SetSuccess(v bool) *DescribeTableResponseBody {
	s.Success = &v
	return s
}

type DescribeTableResponseBodyData struct {
	// Indicates the details about the table schema.
	List []*DescribeTableResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeTableResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTableResponseBodyData) SetList(v []*DescribeTableResponseBodyDataList) *DescribeTableResponseBodyData {
	s.List = v
	return s
}

type DescribeTableResponseBodyDataList struct {
	// Indicates the name of a column.
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// Indicates the type of the column.
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// Extra
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// Indicates the primary key of the table.
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// Indicates whether the column can be empty.
	IsAllowNull *string `json:"IsAllowNull,omitempty" xml:"IsAllowNull,omitempty"`
	// Indicates whether the column is the primary key column of the table.
	IsPk *string `json:"IsPk,omitempty" xml:"IsPk,omitempty"`
}

func (s DescribeTableResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeTableResponseBodyDataList) SetColumnName(v string) *DescribeTableResponseBodyDataList {
	s.ColumnName = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetColumnType(v string) *DescribeTableResponseBodyDataList {
	s.ColumnType = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetExtra(v string) *DescribeTableResponseBodyDataList {
	s.Extra = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetIndex(v string) *DescribeTableResponseBodyDataList {
	s.Index = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetIsAllowNull(v string) *DescribeTableResponseBodyDataList {
	s.IsAllowNull = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetIsPk(v string) *DescribeTableResponseBodyDataList {
	s.IsPk = &v
	return s
}

type DescribeTableResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTableResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableResponse) SetHeaders(v map[string]*string) *DescribeTableResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableResponse) SetStatusCode(v int32) *DescribeTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTableResponse) SetBody(v *DescribeTableResponseBody) *DescribeTableResponse {
	s.Body = v
	return s
}

type DescribeTableListByTypeRequest struct {
	// The number of the page to return.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field that you specify for your query.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of tables. Valid values:
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s DescribeTableListByTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableListByTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableListByTypeRequest) SetCurrentPage(v int32) *DescribeTableListByTypeRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetDbName(v string) *DescribeTableListByTypeRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetDrdsInstanceId(v string) *DescribeTableListByTypeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetPageSize(v int32) *DescribeTableListByTypeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetQuery(v string) *DescribeTableListByTypeRequest {
	s.Query = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetRegionId(v string) *DescribeTableListByTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetTableType(v string) *DescribeTableListByTypeRequest {
	s.TableType = &v
	return s
}

type DescribeTableListByTypeResponseBody struct {
	// Indicates the information about tables.
	List []*DescribeTableListByTypeResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Indicates the page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Indicates the number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates the unique ID of the request. If the request fails, provide this ID for technical support to troubleshoot the failure.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates the total number of returned tables.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeTableListByTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableListByTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableListByTypeResponseBody) SetList(v []*DescribeTableListByTypeResponseBodyList) *DescribeTableListByTypeResponseBody {
	s.List = v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetPageNumber(v int32) *DescribeTableListByTypeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetPageSize(v int32) *DescribeTableListByTypeResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetRequestId(v string) *DescribeTableListByTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetSuccess(v bool) *DescribeTableListByTypeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetTotal(v int32) *DescribeTableListByTypeResponseBody {
	s.Total = &v
	return s
}

type DescribeTableListByTypeResponseBodyList struct {
	// Indicates the property of a table.
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// Indicates the name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableListByTypeResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableListByTypeResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeTableListByTypeResponseBodyList) SetProperty(v string) *DescribeTableListByTypeResponseBodyList {
	s.Property = &v
	return s
}

func (s *DescribeTableListByTypeResponseBodyList) SetTableName(v string) *DescribeTableListByTypeResponseBodyList {
	s.TableName = &v
	return s
}

type DescribeTableListByTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTableListByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTableListByTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableListByTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableListByTypeResponse) SetHeaders(v map[string]*string) *DescribeTableListByTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableListByTypeResponse) SetStatusCode(v int32) *DescribeTableListByTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTableListByTypeResponse) SetBody(v *DescribeTableListByTypeResponseBody) *DescribeTableListByTypeResponse {
	s.Body = v
	return s
}

type DescribeTablesRequest struct {
	// The page number of the returned page.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the database whose tables you want to query.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The number of tables returned on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query condition. The value of this parameter is the ID of the PolarDB-X 1.0 instance.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablesRequest) SetCurrentPage(v int32) *DescribeTablesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTablesRequest) SetDbName(v string) *DescribeTablesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTablesRequest) SetDrdsInstanceId(v string) *DescribeTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeTablesRequest) SetPageSize(v int32) *DescribeTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTablesRequest) SetQuery(v string) *DescribeTablesRequest {
	s.Query = &v
	return s
}

func (s *DescribeTablesRequest) SetRegionId(v string) *DescribeTablesRequest {
	s.RegionId = &v
	return s
}

type DescribeTablesResponseBody struct {
	// The list of returned tables.
	List []*DescribeTablesResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The number of returned pages.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of tables returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned tables.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBody) SetList(v []*DescribeTablesResponseBodyList) *DescribeTablesResponseBody {
	s.List = v
	return s
}

func (s *DescribeTablesResponseBody) SetPageNumber(v int32) *DescribeTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTablesResponseBody) SetPageSize(v int32) *DescribeTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTablesResponseBody) SetRequestId(v string) *DescribeTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTablesResponseBody) SetSuccess(v bool) *DescribeTablesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTablesResponseBody) SetTotal(v int32) *DescribeTablesResponseBody {
	s.Total = &v
	return s
}

type DescribeTablesResponseBodyList struct {
	// Indicates whether full table scanning is allowed.
	AllowFullTableScan *bool `json:"AllowFullTableScan,omitempty" xml:"AllowFullTableScan,omitempty"`
	// Indicates whether the table is a replicated table.
	Broadcast *bool `json:"Broadcast,omitempty" xml:"Broadcast,omitempty"`
	// The type of the PolarDB-X 1.0 instance. Valid values:
	//
	// *   0: The instance is a dedicated instance.
	// *   1: The instance is a shard instance.
	DbInstType *int32 `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// Indicates whether the table is locked.
	IsLocked *bool `json:"IsLocked,omitempty" xml:"IsLocked,omitempty"`
	// Indicates whether the table is sharded.
	IsShard *bool `json:"IsShard,omitempty" xml:"IsShard,omitempty"`
	// The shard key of the table.
	ShardKey *string `json:"ShardKey,omitempty" xml:"ShardKey,omitempty"`
	// Indicates whether sharding tasks are performed on the table. Valid values:
	//
	// *   0: No sharding task is performed on the table.
	// *   1: Sharding tasks are performed on the table.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the table.
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeTablesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyList) SetAllowFullTableScan(v bool) *DescribeTablesResponseBodyList {
	s.AllowFullTableScan = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetBroadcast(v bool) *DescribeTablesResponseBodyList {
	s.Broadcast = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetDbInstType(v int32) *DescribeTablesResponseBodyList {
	s.DbInstType = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetIsLocked(v bool) *DescribeTablesResponseBodyList {
	s.IsLocked = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetIsShard(v bool) *DescribeTablesResponseBodyList {
	s.IsShard = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetShardKey(v string) *DescribeTablesResponseBodyList {
	s.ShardKey = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetStatus(v int32) *DescribeTablesResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetTable(v string) *DescribeTablesResponseBodyList {
	s.Table = &v
	return s
}

type DescribeTablesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponse) SetHeaders(v map[string]*string) *DescribeTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTablesResponse) SetStatusCode(v int32) *DescribeTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTablesResponse) SetBody(v *DescribeTablesResponseBody) *DescribeTablesResponse {
	s.Body = v
	return s
}

type DisableSqlAuditRequest struct {
	// The name of the database for which you want to disable the SQL audit feature.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DisableSqlAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableSqlAuditRequest) GoString() string {
	return s.String()
}

func (s *DisableSqlAuditRequest) SetDbName(v string) *DisableSqlAuditRequest {
	s.DbName = &v
	return s
}

func (s *DisableSqlAuditRequest) SetDrdsInstanceId(v string) *DisableSqlAuditRequest {
	s.DrdsInstanceId = &v
	return s
}

type DisableSqlAuditResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return result.
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableSqlAuditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableSqlAuditResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSqlAuditResponseBody) SetRequestId(v string) *DisableSqlAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSqlAuditResponseBody) SetResult(v bool) *DisableSqlAuditResponseBody {
	s.Result = &v
	return s
}

func (s *DisableSqlAuditResponseBody) SetSuccess(v bool) *DisableSqlAuditResponseBody {
	s.Success = &v
	return s
}

type DisableSqlAuditResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSqlAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSqlAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableSqlAuditResponse) GoString() string {
	return s.String()
}

func (s *DisableSqlAuditResponse) SetHeaders(v map[string]*string) *DisableSqlAuditResponse {
	s.Headers = v
	return s
}

func (s *DisableSqlAuditResponse) SetStatusCode(v int32) *DisableSqlAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSqlAuditResponse) SetBody(v *DisableSqlAuditResponseBody) *DisableSqlAuditResponse {
	s.Body = v
	return s
}

type EnableInstanceIpv6AddressRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region in which the instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableInstanceIpv6AddressRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableInstanceIpv6AddressRequest) GoString() string {
	return s.String()
}

func (s *EnableInstanceIpv6AddressRequest) SetDrdsInstanceId(v string) *EnableInstanceIpv6AddressRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *EnableInstanceIpv6AddressRequest) SetRegionId(v string) *EnableInstanceIpv6AddressRequest {
	s.RegionId = &v
	return s
}

type EnableInstanceIpv6AddressResponseBody struct {
	// The result of the request.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableInstanceIpv6AddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableInstanceIpv6AddressResponseBody) GoString() string {
	return s.String()
}

func (s *EnableInstanceIpv6AddressResponseBody) SetData(v bool) *EnableInstanceIpv6AddressResponseBody {
	s.Data = &v
	return s
}

func (s *EnableInstanceIpv6AddressResponseBody) SetRequestId(v string) *EnableInstanceIpv6AddressResponseBody {
	s.RequestId = &v
	return s
}

type EnableInstanceIpv6AddressResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableInstanceIpv6AddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableInstanceIpv6AddressResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableInstanceIpv6AddressResponse) GoString() string {
	return s.String()
}

func (s *EnableInstanceIpv6AddressResponse) SetHeaders(v map[string]*string) *EnableInstanceIpv6AddressResponse {
	s.Headers = v
	return s
}

func (s *EnableInstanceIpv6AddressResponse) SetStatusCode(v int32) *EnableInstanceIpv6AddressResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableInstanceIpv6AddressResponse) SetBody(v *EnableInstanceIpv6AddressResponseBody) *EnableInstanceIpv6AddressResponse {
	s.Body = v
	return s
}

type EnableSqlAuditRequest struct {
	// The name of the database for which you want to enable the SQL audit feature.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// Specifies whether to backtrack historical SQL statements for auditing.
	IsRecall *bool `json:"IsRecall,omitempty" xml:"IsRecall,omitempty"`
	// The timestamp that indicates when the backtracking ends. Unit: milliseconds.
	//
	// > The end time of the backtracking must be later than the start time of the backtracking.
	RecallEndTimestamp *string `json:"RecallEndTimestamp,omitempty" xml:"RecallEndTimestamp,omitempty"`
	// The timestamp that indicates when the backtracking starts. Unit: milliseconds.
	RecallStartTimestamp *string `json:"RecallStartTimestamp,omitempty" xml:"RecallStartTimestamp,omitempty"`
}

func (s EnableSqlAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlAuditRequest) GoString() string {
	return s.String()
}

func (s *EnableSqlAuditRequest) SetDbName(v string) *EnableSqlAuditRequest {
	s.DbName = &v
	return s
}

func (s *EnableSqlAuditRequest) SetDrdsInstanceId(v string) *EnableSqlAuditRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *EnableSqlAuditRequest) SetIsRecall(v bool) *EnableSqlAuditRequest {
	s.IsRecall = &v
	return s
}

func (s *EnableSqlAuditRequest) SetRecallEndTimestamp(v string) *EnableSqlAuditRequest {
	s.RecallEndTimestamp = &v
	return s
}

func (s *EnableSqlAuditRequest) SetRecallStartTimestamp(v string) *EnableSqlAuditRequest {
	s.RecallStartTimestamp = &v
	return s
}

type EnableSqlAuditResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indices whether the SQL audit feature is enabled.
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSqlAuditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlAuditResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSqlAuditResponseBody) SetRequestId(v string) *EnableSqlAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableSqlAuditResponseBody) SetResult(v bool) *EnableSqlAuditResponseBody {
	s.Result = &v
	return s
}

func (s *EnableSqlAuditResponseBody) SetSuccess(v bool) *EnableSqlAuditResponseBody {
	s.Success = &v
	return s
}

type EnableSqlAuditResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableSqlAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSqlAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlAuditResponse) GoString() string {
	return s.String()
}

func (s *EnableSqlAuditResponse) SetHeaders(v map[string]*string) *EnableSqlAuditResponse {
	s.Headers = v
	return s
}

func (s *EnableSqlAuditResponse) SetStatusCode(v int32) *EnableSqlAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableSqlAuditResponse) SetBody(v *EnableSqlAuditResponseBody) *EnableSqlAuditResponse {
	s.Body = v
	return s
}

type EnableSqlFlashbackMatchSwitchRequest struct {
	// The name of the database you want to back up.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the ApsaraDB RDS for PostgreSQL instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s EnableSqlFlashbackMatchSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlFlashbackMatchSwitchRequest) GoString() string {
	return s.String()
}

func (s *EnableSqlFlashbackMatchSwitchRequest) SetDbName(v string) *EnableSqlFlashbackMatchSwitchRequest {
	s.DbName = &v
	return s
}

func (s *EnableSqlFlashbackMatchSwitchRequest) SetDrdsInstanceId(v string) *EnableSqlFlashbackMatchSwitchRequest {
	s.DrdsInstanceId = &v
	return s
}

type EnableSqlFlashbackMatchSwitchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether SqlFlashbackMatchSwitch is enabled or not.
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was sent successfully or not.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSqlFlashbackMatchSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlFlashbackMatchSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSqlFlashbackMatchSwitchResponseBody) SetRequestId(v string) *EnableSqlFlashbackMatchSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableSqlFlashbackMatchSwitchResponseBody) SetResult(v bool) *EnableSqlFlashbackMatchSwitchResponseBody {
	s.Result = &v
	return s
}

func (s *EnableSqlFlashbackMatchSwitchResponseBody) SetSuccess(v bool) *EnableSqlFlashbackMatchSwitchResponseBody {
	s.Success = &v
	return s
}

type EnableSqlFlashbackMatchSwitchResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableSqlFlashbackMatchSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSqlFlashbackMatchSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlFlashbackMatchSwitchResponse) GoString() string {
	return s.String()
}

func (s *EnableSqlFlashbackMatchSwitchResponse) SetHeaders(v map[string]*string) *EnableSqlFlashbackMatchSwitchResponse {
	s.Headers = v
	return s
}

func (s *EnableSqlFlashbackMatchSwitchResponse) SetStatusCode(v int32) *EnableSqlFlashbackMatchSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableSqlFlashbackMatchSwitchResponse) SetBody(v *EnableSqlFlashbackMatchSwitchResponseBody) *EnableSqlFlashbackMatchSwitchResponse {
	s.Body = v
	return s
}

type FlashbackRecycleBinTableRequest struct {
	// The name of the database to which the table belongs.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the instance to which the table belongs.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the logical table to be restored.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s FlashbackRecycleBinTableRequest) String() string {
	return tea.Prettify(s)
}

func (s FlashbackRecycleBinTableRequest) GoString() string {
	return s.String()
}

func (s *FlashbackRecycleBinTableRequest) SetDbName(v string) *FlashbackRecycleBinTableRequest {
	s.DbName = &v
	return s
}

func (s *FlashbackRecycleBinTableRequest) SetDrdsInstanceId(v string) *FlashbackRecycleBinTableRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *FlashbackRecycleBinTableRequest) SetRegionId(v string) *FlashbackRecycleBinTableRequest {
	s.RegionId = &v
	return s
}

func (s *FlashbackRecycleBinTableRequest) SetTableName(v string) *FlashbackRecycleBinTableRequest {
	s.TableName = &v
	return s
}

type FlashbackRecycleBinTableResponseBody struct {
	// Indicates whether the deleted logical table is restored.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FlashbackRecycleBinTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FlashbackRecycleBinTableResponseBody) GoString() string {
	return s.String()
}

func (s *FlashbackRecycleBinTableResponseBody) SetData(v bool) *FlashbackRecycleBinTableResponseBody {
	s.Data = &v
	return s
}

func (s *FlashbackRecycleBinTableResponseBody) SetRequestId(v string) *FlashbackRecycleBinTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlashbackRecycleBinTableResponseBody) SetSuccess(v bool) *FlashbackRecycleBinTableResponseBody {
	s.Success = &v
	return s
}

type FlashbackRecycleBinTableResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlashbackRecycleBinTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlashbackRecycleBinTableResponse) String() string {
	return tea.Prettify(s)
}

func (s FlashbackRecycleBinTableResponse) GoString() string {
	return s.String()
}

func (s *FlashbackRecycleBinTableResponse) SetHeaders(v map[string]*string) *FlashbackRecycleBinTableResponse {
	s.Headers = v
	return s
}

func (s *FlashbackRecycleBinTableResponse) SetStatusCode(v int32) *FlashbackRecycleBinTableResponse {
	s.StatusCode = &v
	return s
}

func (s *FlashbackRecycleBinTableResponse) SetBody(v *FlashbackRecycleBinTableResponseBody) *FlashbackRecycleBinTableResponse {
	s.Body = v
	return s
}

type GetDrdsDbRdsRelationInfoRequest struct {
	// The name of the DRDS database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s GetDrdsDbRdsRelationInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDrdsDbRdsRelationInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDrdsDbRdsRelationInfoRequest) SetDbName(v string) *GetDrdsDbRdsRelationInfoRequest {
	s.DbName = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoRequest) SetDrdsInstanceId(v string) *GetDrdsDbRdsRelationInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

type GetDrdsDbRdsRelationInfoResponseBody struct {
	// The structure information about the storage instances of the DRDS database. Each entry corresponds to a primary storage instance.
	Data []*GetDrdsDbRdsRelationInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDrdsDbRdsRelationInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDrdsDbRdsRelationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDrdsDbRdsRelationInfoResponseBody) SetData(v []*GetDrdsDbRdsRelationInfoResponseBodyData) *GetDrdsDbRdsRelationInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBody) SetRequestId(v string) *GetDrdsDbRdsRelationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBody) SetSuccess(v bool) *GetDrdsDbRdsRelationInfoResponseBody {
	s.Success = &v
	return s
}

type GetDrdsDbRdsRelationInfoResponseBodyData struct {
	// The ID of the storage instance.
	RdsInstanceId *string `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	// The IDs of the read-only storage instances.
	ReadOnlyInstanceInfo []*string `json:"ReadOnlyInstanceInfo,omitempty" xml:"ReadOnlyInstanceInfo,omitempty" type:"Repeated"`
	// The ID of the storage instance that is in use. If the specified instance in the request is a primary DRDS instance, the value of this parameter is the ID of the primary storage instance. If the specified instance in the request is a read-only DRDS instance, the value of this parameter is the ID of the secondary storage instance.
	UsedInstanceId *string `json:"UsedInstanceId,omitempty" xml:"UsedInstanceId,omitempty"`
	// The type of the storage instance that is in use.
	UsedInstanceType *string `json:"UsedInstanceType,omitempty" xml:"UsedInstanceType,omitempty"`
}

func (s GetDrdsDbRdsRelationInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDrdsDbRdsRelationInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) SetRdsInstanceId(v string) *GetDrdsDbRdsRelationInfoResponseBodyData {
	s.RdsInstanceId = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) SetReadOnlyInstanceInfo(v []*string) *GetDrdsDbRdsRelationInfoResponseBodyData {
	s.ReadOnlyInstanceInfo = v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) SetUsedInstanceId(v string) *GetDrdsDbRdsRelationInfoResponseBodyData {
	s.UsedInstanceId = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) SetUsedInstanceType(v string) *GetDrdsDbRdsRelationInfoResponseBodyData {
	s.UsedInstanceType = &v
	return s
}

type GetDrdsDbRdsRelationInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDrdsDbRdsRelationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDrdsDbRdsRelationInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDrdsDbRdsRelationInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDrdsDbRdsRelationInfoResponse) SetHeaders(v map[string]*string) *GetDrdsDbRdsRelationInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponse) SetStatusCode(v int32) *GetDrdsDbRdsRelationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponse) SetBody(v *GetDrdsDbRdsRelationInfoResponseBody) *GetDrdsDbRdsRelationInfoResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// Specify the token that is used to display the returned tags on multiple pages.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which the resource is located.
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Set the value to INSTANCE.
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
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
	// The key of the tag that you want to query.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that you want to query.
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
	// The token that is used to display the returned tags on multiple pages.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The list of returned tags.
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

func (s *ListTagResourcesResponseBody) SetSuccess(v bool) *ListTagResourcesResponseBody {
	s.Success = &v
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
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. The value of this parameter is fixed to INSTANCE.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ManagePrivateRdsRequest struct {
	// The ID of the custom ApsaraDB RDS instance at the storage layer.
	//
	// > You can call the [DescribeDrdsRdsInstances](~~215526~~) operation to query the details of all ApsaraDB RDS instances, including the ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// > You can call the [DescribeDrdsInstances](~~139284~~) operation to query the details of all PolarDB-X 1.0 instances within an Alibaba Cloud account, including the IDs of the instances.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The JSON string that consists of request parameters and the values of the request parameters of an operation that you need to call for the custom ApsaraDB RDS instance. The value of a request parameter is of the STRING type. Example: `{NodeId:"1797****"}`.
	//
	// For more information about the request parameters and valid values of the request parameters of each operation, see the request parameter sections in the following topics:
	//
	// *   [DescribeDBInstanceAttribute](~~26231~~)
	// *   [DescribeAvailableClasses](~~196546~~)
	// *   [DescribeSQLCollectorPolicy](~~26292~~)
	// *   [ModifySQLCollectorPolicy](~~26293~~)
	// *   [DescribeParameters](~~26285~~)
	// *   [ModifyParameter](~~26286~~)
	// *   [DescribeDBInstanceHAConfig](~~26244~~)
	// *   [SwitchDBInstanceHA](~~26251~~)
	//
	// > Among the required request parameters of the preceding operations, you do not need to specify the `Action` and `DBInstanceId` parameters. You must specify all the other required request parameters.
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The operation that you want to perform on the custom ApsaraDB RDS instance. Valid values:
	//
	// *   **DescribeDBInstanceAttribute**: queries the details of the custom ApsaraDB RDS instance.
	// *   **DescribeAvailableClasses**: queries the specifications that are supported for a custom ApsaraDB RDS instance. The specifications include the instance type and the storage capacity.
	// *   **DescribeSQLCollectorPolicy**: queries whether the SQL Explorer (SQL Audit) feature is enabled for custom ApsaraDB RDS instance.
	// *   **ModifySQLCollectorPolicy**: enables or disables the SQL Explorer (SQL Audit) feature for the custom ApsaraDB RDS instance.
	// *   **DescribeParameters**: queries the parameter settings of the custom ApsaraDB RDS instance.
	// *   **ModifyParameter**: modifies the parameters of the custom ApsaraDB RDS instance.
	// *   **DescribeDBInstanceHAConfig**: queries the high availability mode and data replication mode of the custom ApsaraDB RDS instance.
	// *   **SwitchDBInstanceHA**: switches workloads between the primary and secondary custom ApsaraDB RDS instances.
	RdsAction *string `json:"RdsAction,omitempty" xml:"RdsAction,omitempty"`
	// The ID of the region in which the PolarDB-X 1.0 instance resides.
	//
	// > You can call the [DescribeDrdsInstances](~~139284~~) operation to query the details of all PolarDB-X 1.0 instances within an Alibaba Cloud account, including the IDs of regions in which the instances reside.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ManagePrivateRdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ManagePrivateRdsRequest) GoString() string {
	return s.String()
}

func (s *ManagePrivateRdsRequest) SetDBInstanceId(v string) *ManagePrivateRdsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetDrdsInstanceId(v string) *ManagePrivateRdsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetParams(v string) *ManagePrivateRdsRequest {
	s.Params = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetRdsAction(v string) *ManagePrivateRdsRequest {
	s.RdsAction = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetRegionId(v string) *ManagePrivateRdsRequest {
	s.RegionId = &v
	return s
}

type ManagePrivateRdsResponseBody struct {
	// The parameter result set returned for the operation that is called for the custom ApsaraDB RDS instance.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ManagePrivateRdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManagePrivateRdsResponseBody) GoString() string {
	return s.String()
}

func (s *ManagePrivateRdsResponseBody) SetData(v string) *ManagePrivateRdsResponseBody {
	s.Data = &v
	return s
}

func (s *ManagePrivateRdsResponseBody) SetRequestId(v string) *ManagePrivateRdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManagePrivateRdsResponseBody) SetSuccess(v bool) *ManagePrivateRdsResponseBody {
	s.Success = &v
	return s
}

type ManagePrivateRdsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManagePrivateRdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManagePrivateRdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ManagePrivateRdsResponse) GoString() string {
	return s.String()
}

func (s *ManagePrivateRdsResponse) SetHeaders(v map[string]*string) *ManagePrivateRdsResponse {
	s.Headers = v
	return s
}

func (s *ManagePrivateRdsResponse) SetStatusCode(v int32) *ManagePrivateRdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ManagePrivateRdsResponse) SetBody(v *ManagePrivateRdsResponseBody) *ManagePrivateRdsResponse {
	s.Body = v
	return s
}

type ModifyAccountDescriptionRequest struct {
	// The name of the member account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The description of the account.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the ApsaraDB RDS for PostgreSQL instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDescription(v string) *ModifyAccountDescriptionRequest {
	s.Description = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDrdsInstanceId(v string) *ModifyAccountDescriptionRequest {
	s.DrdsInstanceId = &v
	return s
}

type ModifyAccountDescriptionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was sent successfully or not.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *ModifyAccountDescriptionResponseBody) SetSuccess(v bool) *ModifyAccountDescriptionResponseBody {
	s.Success = &v
	return s
}

type ModifyAccountDescriptionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ModifyAccountPrivilegeRequest struct {
	// The username of the account that you want to modify.
	AccountName *string                                     `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbPrivilege []*ModifyAccountPrivilegeRequestDbPrivilege `json:"DbPrivilege,omitempty" xml:"DbPrivilege,omitempty" type:"Repeated"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region in which the PolarDB-X 1.0 instance is located.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAccountPrivilegeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeRequest) SetAccountName(v string) *ModifyAccountPrivilegeRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetDbPrivilege(v []*ModifyAccountPrivilegeRequestDbPrivilege) *ModifyAccountPrivilegeRequest {
	s.DbPrivilege = v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetDrdsInstanceId(v string) *ModifyAccountPrivilegeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetRegionId(v string) *ModifyAccountPrivilegeRequest {
	s.RegionId = &v
	return s
}

type ModifyAccountPrivilegeRequestDbPrivilege struct {
	// The name of the database that you want to manage by using the account to modify.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The permissions that you want to grant to the account.
	Privilege *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
}

func (s ModifyAccountPrivilegeRequestDbPrivilege) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegeRequestDbPrivilege) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeRequestDbPrivilege) SetDbName(v string) *ModifyAccountPrivilegeRequestDbPrivilege {
	s.DbName = &v
	return s
}

func (s *ModifyAccountPrivilegeRequestDbPrivilege) SetPrivilege(v string) *ModifyAccountPrivilegeRequestDbPrivilege {
	s.Privilege = &v
	return s
}

type ModifyAccountPrivilegeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAccountPrivilegeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeResponseBody) SetRequestId(v string) *ModifyAccountPrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountPrivilegeResponseBody) SetSuccess(v bool) *ModifyAccountPrivilegeResponseBody {
	s.Success = &v
	return s
}

type ModifyAccountPrivilegeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountPrivilegeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeResponse) SetHeaders(v map[string]*string) *ModifyAccountPrivilegeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountPrivilegeResponse) SetStatusCode(v int32) *ModifyAccountPrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountPrivilegeResponse) SetBody(v *ModifyAccountPrivilegeResponseBody) *ModifyAccountPrivilegeResponse {
	s.Body = v
	return s
}

type ModifyDrdsInstanceDescriptionRequest struct {
	// The description of the DRDS instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s ModifyDrdsInstanceDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDrdsInstanceDescriptionRequest) SetDescription(v string) *ModifyDrdsInstanceDescriptionRequest {
	s.Description = &v
	return s
}

func (s *ModifyDrdsInstanceDescriptionRequest) SetDrdsInstanceId(v string) *ModifyDrdsInstanceDescriptionRequest {
	s.DrdsInstanceId = &v
	return s
}

type ModifyDrdsInstanceDescriptionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDrdsInstanceDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsInstanceDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDrdsInstanceDescriptionResponseBody) SetRequestId(v string) *ModifyDrdsInstanceDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDrdsInstanceDescriptionResponseBody) SetSuccess(v bool) *ModifyDrdsInstanceDescriptionResponseBody {
	s.Success = &v
	return s
}

type ModifyDrdsInstanceDescriptionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDrdsInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDrdsInstanceDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDrdsInstanceDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDrdsInstanceDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDrdsInstanceDescriptionResponse) SetStatusCode(v int32) *ModifyDrdsInstanceDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDrdsInstanceDescriptionResponse) SetBody(v *ModifyDrdsInstanceDescriptionResponseBody) *ModifyDrdsInstanceDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDrdsIpWhiteListRequest struct {
	// The name of the DRDS database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the Message Queue for Apache Kafka instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The attribute of the IP address whitelist group.
	GroupAttribute *string `json:"GroupAttribute,omitempty" xml:"GroupAttribute,omitempty"`
	// The name of the IP address whitelist group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The modified whitelist. Separate multiple IP addresses with commas (,).
	IpWhiteList *string `json:"IpWhiteList,omitempty" xml:"IpWhiteList,omitempty"`
	// Specifies the mode. Valid values:
	//
	// *   `True`: append modifications
	// *   `False`: overwrite modification
	Mode *bool `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ModifyDrdsIpWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyDrdsIpWhiteListRequest) SetDbName(v string) *ModifyDrdsIpWhiteListRequest {
	s.DbName = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetDrdsInstanceId(v string) *ModifyDrdsIpWhiteListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetGroupAttribute(v string) *ModifyDrdsIpWhiteListRequest {
	s.GroupAttribute = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetGroupName(v string) *ModifyDrdsIpWhiteListRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetIpWhiteList(v string) *ModifyDrdsIpWhiteListRequest {
	s.IpWhiteList = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetMode(v bool) *ModifyDrdsIpWhiteListRequest {
	s.Mode = &v
	return s
}

type ModifyDrdsIpWhiteListResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDrdsIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDrdsIpWhiteListResponseBody) SetRequestId(v string) *ModifyDrdsIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDrdsIpWhiteListResponseBody) SetSuccess(v bool) *ModifyDrdsIpWhiteListResponseBody {
	s.Success = &v
	return s
}

type ModifyDrdsIpWhiteListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDrdsIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDrdsIpWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ModifyDrdsIpWhiteListResponse) SetHeaders(v map[string]*string) *ModifyDrdsIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ModifyDrdsIpWhiteListResponse) SetStatusCode(v int32) *ModifyDrdsIpWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDrdsIpWhiteListResponse) SetBody(v *ModifyDrdsIpWhiteListResponseBody) *ModifyDrdsIpWhiteListResponse {
	s.Body = v
	return s
}

type ModifyPolarDbReadWeightRequest struct {
	// Polar cluster ID.
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The node list in the destination apsaradb for PolarDB cluster. The nodes in each cluster are separated with commas (,) and colons (:).
	DbNodeIds *string `json:"DbNodeIds,omitempty" xml:"DbNodeIds,omitempty"`
	// The ID of a DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The weight of the PolarDB cluster. Separate multiple weights with commas (,).
	Weights *string `json:"Weights,omitempty" xml:"Weights,omitempty"`
}

func (s ModifyPolarDbReadWeightRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolarDbReadWeightRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolarDbReadWeightRequest) SetDbInstanceId(v string) *ModifyPolarDbReadWeightRequest {
	s.DbInstanceId = &v
	return s
}

func (s *ModifyPolarDbReadWeightRequest) SetDbName(v string) *ModifyPolarDbReadWeightRequest {
	s.DbName = &v
	return s
}

func (s *ModifyPolarDbReadWeightRequest) SetDbNodeIds(v string) *ModifyPolarDbReadWeightRequest {
	s.DbNodeIds = &v
	return s
}

func (s *ModifyPolarDbReadWeightRequest) SetDrdsInstanceId(v string) *ModifyPolarDbReadWeightRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyPolarDbReadWeightRequest) SetWeights(v string) *ModifyPolarDbReadWeightRequest {
	s.Weights = &v
	return s
}

type ModifyPolarDbReadWeightResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the database creation failure records were removed from the PolarDB-X instance.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyPolarDbReadWeightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolarDbReadWeightResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolarDbReadWeightResponseBody) SetRequestId(v string) *ModifyPolarDbReadWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPolarDbReadWeightResponseBody) SetSuccess(v bool) *ModifyPolarDbReadWeightResponseBody {
	s.Success = &v
	return s
}

type ModifyPolarDbReadWeightResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolarDbReadWeightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolarDbReadWeightResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolarDbReadWeightResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolarDbReadWeightResponse) SetHeaders(v map[string]*string) *ModifyPolarDbReadWeightResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolarDbReadWeightResponse) SetStatusCode(v int32) *ModifyPolarDbReadWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolarDbReadWeightResponse) SetBody(v *ModifyPolarDbReadWeightResponseBody) *ModifyPolarDbReadWeightResponse {
	s.Body = v
	return s
}

type ModifyRdsReadWeightRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The names of the ApsaraDB RDS for MySQL instances. Separate the names with commas (,).
	InstanceNames *string `json:"InstanceNames,omitempty" xml:"InstanceNames,omitempty"`
	// The weights of the ApsaraDB RDS for MySQL instances. Separate the weights with commas (,).
	Weights *string `json:"Weights,omitempty" xml:"Weights,omitempty"`
}

func (s ModifyRdsReadWeightRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRdsReadWeightRequest) GoString() string {
	return s.String()
}

func (s *ModifyRdsReadWeightRequest) SetDbName(v string) *ModifyRdsReadWeightRequest {
	s.DbName = &v
	return s
}

func (s *ModifyRdsReadWeightRequest) SetDrdsInstanceId(v string) *ModifyRdsReadWeightRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyRdsReadWeightRequest) SetInstanceNames(v string) *ModifyRdsReadWeightRequest {
	s.InstanceNames = &v
	return s
}

func (s *ModifyRdsReadWeightRequest) SetWeights(v string) *ModifyRdsReadWeightRequest {
	s.Weights = &v
	return s
}

type ModifyRdsReadWeightResponseBody struct {
	// Indicates the ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyRdsReadWeightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRdsReadWeightResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRdsReadWeightResponseBody) SetRequestId(v string) *ModifyRdsReadWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRdsReadWeightResponseBody) SetSuccess(v bool) *ModifyRdsReadWeightResponseBody {
	s.Success = &v
	return s
}

type ModifyRdsReadWeightResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRdsReadWeightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRdsReadWeightResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRdsReadWeightResponse) GoString() string {
	return s.String()
}

func (s *ModifyRdsReadWeightResponse) SetHeaders(v map[string]*string) *ModifyRdsReadWeightResponse {
	s.Headers = v
	return s
}

func (s *ModifyRdsReadWeightResponse) SetStatusCode(v int32) *ModifyRdsReadWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRdsReadWeightResponse) SetBody(v *ModifyRdsReadWeightResponseBody) *ModifyRdsReadWeightResponse {
	s.Body = v
	return s
}

type PutStartBackupRequest struct {
	// If you need to back up data at the database level, you must specify the list of databases to be backed up, and separate multiple databases with commas (,).
	BackupDbNames *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	// The backup level. Valid values:
	//
	// *   instance: instance
	// *   db: The database type.
	BackupLevel *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	// The backup mode. For more information, see [backup mode](~~108631~~) and the valid values are as follows:
	//
	// *   phy: fast backup
	// *   logic: Consistent backup
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s PutStartBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s PutStartBackupRequest) GoString() string {
	return s.String()
}

func (s *PutStartBackupRequest) SetBackupDbNames(v string) *PutStartBackupRequest {
	s.BackupDbNames = &v
	return s
}

func (s *PutStartBackupRequest) SetBackupLevel(v string) *PutStartBackupRequest {
	s.BackupLevel = &v
	return s
}

func (s *PutStartBackupRequest) SetBackupMode(v string) *PutStartBackupRequest {
	s.BackupMode = &v
	return s
}

func (s *PutStartBackupRequest) SetDrdsInstanceId(v string) *PutStartBackupRequest {
	s.DrdsInstanceId = &v
	return s
}

type PutStartBackupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the backup task was submitted.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutStartBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutStartBackupResponseBody) GoString() string {
	return s.String()
}

func (s *PutStartBackupResponseBody) SetRequestId(v string) *PutStartBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutStartBackupResponseBody) SetResult(v string) *PutStartBackupResponseBody {
	s.Result = &v
	return s
}

func (s *PutStartBackupResponseBody) SetSuccess(v bool) *PutStartBackupResponseBody {
	s.Success = &v
	return s
}

type PutStartBackupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutStartBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutStartBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s PutStartBackupResponse) GoString() string {
	return s.String()
}

func (s *PutStartBackupResponse) SetHeaders(v map[string]*string) *PutStartBackupResponse {
	s.Headers = v
	return s
}

func (s *PutStartBackupResponse) SetStatusCode(v int32) *PutStartBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *PutStartBackupResponse) SetBody(v *PutStartBackupResponseBody) *PutStartBackupResponse {
	s.Body = v
	return s
}

type RefreshDrdsAtomUrlRequest struct {
	// The name of the DRDS database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RefreshDrdsAtomUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshDrdsAtomUrlRequest) GoString() string {
	return s.String()
}

func (s *RefreshDrdsAtomUrlRequest) SetDbName(v string) *RefreshDrdsAtomUrlRequest {
	s.DbName = &v
	return s
}

func (s *RefreshDrdsAtomUrlRequest) SetDrdsInstanceId(v string) *RefreshDrdsAtomUrlRequest {
	s.DrdsInstanceId = &v
	return s
}

type RefreshDrdsAtomUrlResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the connection after refresh was successful.
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefreshDrdsAtomUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshDrdsAtomUrlResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshDrdsAtomUrlResponseBody) SetRequestId(v string) *RefreshDrdsAtomUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshDrdsAtomUrlResponseBody) SetResult(v bool) *RefreshDrdsAtomUrlResponseBody {
	s.Result = &v
	return s
}

func (s *RefreshDrdsAtomUrlResponseBody) SetSuccess(v bool) *RefreshDrdsAtomUrlResponseBody {
	s.Success = &v
	return s
}

type RefreshDrdsAtomUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshDrdsAtomUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshDrdsAtomUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshDrdsAtomUrlResponse) GoString() string {
	return s.String()
}

func (s *RefreshDrdsAtomUrlResponse) SetHeaders(v map[string]*string) *RefreshDrdsAtomUrlResponse {
	s.Headers = v
	return s
}

func (s *RefreshDrdsAtomUrlResponse) SetStatusCode(v int32) *RefreshDrdsAtomUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshDrdsAtomUrlResponse) SetBody(v *RefreshDrdsAtomUrlResponseBody) *RefreshDrdsAtomUrlResponse {
	s.Body = v
	return s
}

type ReleaseInstanceInternetAddressRequest struct {
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The region where the instance is located.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleaseInstanceInternetAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceInternetAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceInternetAddressRequest) SetDrdsInstanceId(v string) *ReleaseInstanceInternetAddressRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ReleaseInstanceInternetAddressRequest) SetRegionId(v string) *ReleaseInstanceInternetAddressRequest {
	s.RegionId = &v
	return s
}

type ReleaseInstanceInternetAddressResponseBody struct {
	// The result returned by the current API.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseInstanceInternetAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceInternetAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceInternetAddressResponseBody) SetData(v bool) *ReleaseInstanceInternetAddressResponseBody {
	s.Data = &v
	return s
}

func (s *ReleaseInstanceInternetAddressResponseBody) SetRequestId(v string) *ReleaseInstanceInternetAddressResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseInstanceInternetAddressResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseInstanceInternetAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseInstanceInternetAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceInternetAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceInternetAddressResponse) SetHeaders(v map[string]*string) *ReleaseInstanceInternetAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstanceInternetAddressResponse) SetStatusCode(v int32) *ReleaseInstanceInternetAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseInstanceInternetAddressResponse) SetBody(v *ReleaseInstanceInternetAddressResponseBody) *ReleaseInstanceInternetAddressResponse {
	s.Body = v
	return s
}

type RemoveBackupsSetRequest struct {
	// The ID of the backup set. You can call the [DescribeBackupSets](~~139331~~) interface to query the ID of a backup set.
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveBackupsSetRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveBackupsSetRequest) GoString() string {
	return s.String()
}

func (s *RemoveBackupsSetRequest) SetBackupId(v string) *RemoveBackupsSetRequest {
	s.BackupId = &v
	return s
}

func (s *RemoveBackupsSetRequest) SetDrdsInstanceId(v string) *RemoveBackupsSetRequest {
	s.DrdsInstanceId = &v
	return s
}

type RemoveBackupsSetResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether SQL audit was disabled for the DRDS database.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveBackupsSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveBackupsSetResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveBackupsSetResponseBody) SetRequestId(v string) *RemoveBackupsSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveBackupsSetResponseBody) SetResult(v string) *RemoveBackupsSetResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveBackupsSetResponseBody) SetSuccess(v bool) *RemoveBackupsSetResponseBody {
	s.Success = &v
	return s
}

type RemoveBackupsSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveBackupsSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveBackupsSetResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveBackupsSetResponse) GoString() string {
	return s.String()
}

func (s *RemoveBackupsSetResponse) SetHeaders(v map[string]*string) *RemoveBackupsSetResponse {
	s.Headers = v
	return s
}

func (s *RemoveBackupsSetResponse) SetStatusCode(v int32) *RemoveBackupsSetResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveBackupsSetResponse) SetBody(v *RemoveBackupsSetResponseBody) *RemoveBackupsSetResponse {
	s.Body = v
	return s
}

type RemoveDrdsDbRequest struct {
	// The name of the database you want to back up.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the DRDS instance to which the destination database belongs.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveDrdsDbRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbRequest) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbRequest) SetDbName(v string) *RemoveDrdsDbRequest {
	s.DbName = &v
	return s
}

func (s *RemoveDrdsDbRequest) SetDrdsInstanceId(v string) *RemoveDrdsDbRequest {
	s.DrdsInstanceId = &v
	return s
}

type RemoveDrdsDbResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveDrdsDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbResponseBody) SetRequestId(v string) *RemoveDrdsDbResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDrdsDbResponseBody) SetSuccess(v bool) *RemoveDrdsDbResponseBody {
	s.Success = &v
	return s
}

type RemoveDrdsDbResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDrdsDbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDrdsDbResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbResponse) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbResponse) SetHeaders(v map[string]*string) *RemoveDrdsDbResponse {
	s.Headers = v
	return s
}

func (s *RemoveDrdsDbResponse) SetStatusCode(v int32) *RemoveDrdsDbResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDrdsDbResponse) SetBody(v *RemoveDrdsDbResponseBody) *RemoveDrdsDbResponse {
	s.Body = v
	return s
}

type RemoveDrdsDbFailedRecordRequest struct {
	// The name of the DRDS database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the ApsaraDB RDS for PostgreSQL instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveDrdsDbFailedRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbFailedRecordRequest) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbFailedRecordRequest) SetDbName(v string) *RemoveDrdsDbFailedRecordRequest {
	s.DbName = &v
	return s
}

func (s *RemoveDrdsDbFailedRecordRequest) SetDrdsInstanceId(v string) *RemoveDrdsDbFailedRecordRequest {
	s.DrdsInstanceId = &v
	return s
}

type RemoveDrdsDbFailedRecordResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the database creation failure records were deleted from the DRDS instance.
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveDrdsDbFailedRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbFailedRecordResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbFailedRecordResponseBody) SetRequestId(v string) *RemoveDrdsDbFailedRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDrdsDbFailedRecordResponseBody) SetResult(v bool) *RemoveDrdsDbFailedRecordResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveDrdsDbFailedRecordResponseBody) SetSuccess(v bool) *RemoveDrdsDbFailedRecordResponseBody {
	s.Success = &v
	return s
}

type RemoveDrdsDbFailedRecordResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDrdsDbFailedRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDrdsDbFailedRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbFailedRecordResponse) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbFailedRecordResponse) SetHeaders(v map[string]*string) *RemoveDrdsDbFailedRecordResponse {
	s.Headers = v
	return s
}

func (s *RemoveDrdsDbFailedRecordResponse) SetStatusCode(v int32) *RemoveDrdsDbFailedRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDrdsDbFailedRecordResponse) SetBody(v *RemoveDrdsDbFailedRecordResponseBody) *RemoveDrdsDbFailedRecordResponse {
	s.Body = v
	return s
}

type RemoveDrdsInstanceRequest struct {
	// The ID of the instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveDrdsInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *RemoveDrdsInstanceRequest) SetDrdsInstanceId(v string) *RemoveDrdsInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

type RemoveDrdsInstanceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveDrdsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDrdsInstanceResponseBody) SetRequestId(v string) *RemoveDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDrdsInstanceResponseBody) SetSuccess(v bool) *RemoveDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

type RemoveDrdsInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDrdsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsInstanceResponse) GoString() string {
	return s.String()
}

func (s *RemoveDrdsInstanceResponse) SetHeaders(v map[string]*string) *RemoveDrdsInstanceResponse {
	s.Headers = v
	return s
}

func (s *RemoveDrdsInstanceResponse) SetStatusCode(v int32) *RemoveDrdsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDrdsInstanceResponse) SetBody(v *RemoveDrdsInstanceResponseBody) *RemoveDrdsInstanceResponse {
	s.Body = v
	return s
}

type RemoveInstanceAccountRequest struct {
	// The name of the member account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveInstanceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *RemoveInstanceAccountRequest) SetAccountName(v string) *RemoveInstanceAccountRequest {
	s.AccountName = &v
	return s
}

func (s *RemoveInstanceAccountRequest) SetDrdsInstanceId(v string) *RemoveInstanceAccountRequest {
	s.DrdsInstanceId = &v
	return s
}

type RemoveInstanceAccountResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveInstanceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInstanceAccountResponseBody) SetRequestId(v string) *RemoveInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveInstanceAccountResponseBody) SetSuccess(v bool) *RemoveInstanceAccountResponseBody {
	s.Success = &v
	return s
}

type RemoveInstanceAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveInstanceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *RemoveInstanceAccountResponse) SetHeaders(v map[string]*string) *RemoveInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *RemoveInstanceAccountResponse) SetStatusCode(v int32) *RemoveInstanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveInstanceAccountResponse) SetBody(v *RemoveInstanceAccountResponseBody) *RemoveInstanceAccountResponse {
	s.Body = v
	return s
}

type RemoveRecycleBinTableRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the logical table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s RemoveRecycleBinTableRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveRecycleBinTableRequest) GoString() string {
	return s.String()
}

func (s *RemoveRecycleBinTableRequest) SetDbName(v string) *RemoveRecycleBinTableRequest {
	s.DbName = &v
	return s
}

func (s *RemoveRecycleBinTableRequest) SetDrdsInstanceId(v string) *RemoveRecycleBinTableRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RemoveRecycleBinTableRequest) SetRegionId(v string) *RemoveRecycleBinTableRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveRecycleBinTableRequest) SetTableName(v string) *RemoveRecycleBinTableRequest {
	s.TableName = &v
	return s
}

type RemoveRecycleBinTableResponseBody struct {
	// Indicates whether the table in the recycle bin is deleted.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveRecycleBinTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveRecycleBinTableResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveRecycleBinTableResponseBody) SetData(v bool) *RemoveRecycleBinTableResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveRecycleBinTableResponseBody) SetRequestId(v string) *RemoveRecycleBinTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveRecycleBinTableResponseBody) SetSuccess(v bool) *RemoveRecycleBinTableResponseBody {
	s.Success = &v
	return s
}

type RemoveRecycleBinTableResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveRecycleBinTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveRecycleBinTableResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveRecycleBinTableResponse) GoString() string {
	return s.String()
}

func (s *RemoveRecycleBinTableResponse) SetHeaders(v map[string]*string) *RemoveRecycleBinTableResponse {
	s.Headers = v
	return s
}

func (s *RemoveRecycleBinTableResponse) SetStatusCode(v int32) *RemoveRecycleBinTableResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveRecycleBinTableResponse) SetBody(v *RemoveRecycleBinTableResponseBody) *RemoveRecycleBinTableResponse {
	s.Body = v
	return s
}

type RestartDrdsInstanceRequest struct {
	// The ID of a DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RestartDrdsInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDrdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDrdsInstanceRequest) SetDrdsInstanceId(v string) *RestartDrdsInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

type RestartDrdsInstanceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the database creation failure records were removed from the PolarDB-X instance.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the task.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RestartDrdsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDrdsInstanceResponseBody) SetRequestId(v string) *RestartDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDrdsInstanceResponseBody) SetSuccess(v bool) *RestartDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RestartDrdsInstanceResponseBody) SetTaskId(v int64) *RestartDrdsInstanceResponseBody {
	s.TaskId = &v
	return s
}

type RestartDrdsInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartDrdsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartDrdsInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartDrdsInstanceResponse) SetHeaders(v map[string]*string) *RestartDrdsInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartDrdsInstanceResponse) SetStatusCode(v int32) *RestartDrdsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDrdsInstanceResponse) SetBody(v *RestartDrdsInstanceResponseBody) *RestartDrdsInstanceResponse {
	s.Body = v
	return s
}

type RollbackInstanceVersionRequest struct {
	// The instance ID.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RollbackInstanceVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackInstanceVersionRequest) GoString() string {
	return s.String()
}

func (s *RollbackInstanceVersionRequest) SetDrdsInstanceId(v string) *RollbackInstanceVersionRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RollbackInstanceVersionRequest) SetRegionId(v string) *RollbackInstanceVersionRequest {
	s.RegionId = &v
	return s
}

type RollbackInstanceVersionResponseBody struct {
	// Indicates whether the instance version was rolled back.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackInstanceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackInstanceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackInstanceVersionResponseBody) SetData(v string) *RollbackInstanceVersionResponseBody {
	s.Data = &v
	return s
}

func (s *RollbackInstanceVersionResponseBody) SetRequestId(v string) *RollbackInstanceVersionResponseBody {
	s.RequestId = &v
	return s
}

type RollbackInstanceVersionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackInstanceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackInstanceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackInstanceVersionResponse) GoString() string {
	return s.String()
}

func (s *RollbackInstanceVersionResponse) SetHeaders(v map[string]*string) *RollbackInstanceVersionResponse {
	s.Headers = v
	return s
}

func (s *RollbackInstanceVersionResponse) SetStatusCode(v int32) *RollbackInstanceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackInstanceVersionResponse) SetBody(v *RollbackInstanceVersionResponseBody) *RollbackInstanceVersionResponse {
	s.Body = v
	return s
}

type SetBackupLocalRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// Specifies whether to enable the feature to forcibly delete binary log files if the used storage space reaches 90% of the total storage space or the remaining storage space is less than 5 GB. Valid values: 1 and 0. A value of 1 specifies to enable this feature. A value of 0 specifies not to enable this feature.
	HighSpaceUsageProtection *string `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	// The number of hours for which log backup files are retained on the instance. Valid values: 0 to 168. Default value: 18. A value of 0 indicates that log backup files are not retained.
	LocalLogRetentionHours *string `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	// The maximum storage space usage that is allowed for log files on the instance. Valid values: 0 to 50. Default value: 30.
	LocalLogRetentionSpace *string `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
}

func (s SetBackupLocalRequest) String() string {
	return tea.Prettify(s)
}

func (s SetBackupLocalRequest) GoString() string {
	return s.String()
}

func (s *SetBackupLocalRequest) SetDrdsInstanceId(v string) *SetBackupLocalRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetBackupLocalRequest) SetHighSpaceUsageProtection(v string) *SetBackupLocalRequest {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *SetBackupLocalRequest) SetLocalLogRetentionHours(v string) *SetBackupLocalRequest {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *SetBackupLocalRequest) SetLocalLogRetentionSpace(v string) *SetBackupLocalRequest {
	s.LocalLogRetentionSpace = &v
	return s
}

type SetBackupLocalResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result returned.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetBackupLocalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetBackupLocalResponseBody) GoString() string {
	return s.String()
}

func (s *SetBackupLocalResponseBody) SetRequestId(v string) *SetBackupLocalResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetBackupLocalResponseBody) SetResult(v string) *SetBackupLocalResponseBody {
	s.Result = &v
	return s
}

func (s *SetBackupLocalResponseBody) SetSuccess(v bool) *SetBackupLocalResponseBody {
	s.Success = &v
	return s
}

type SetBackupLocalResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetBackupLocalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetBackupLocalResponse) String() string {
	return tea.Prettify(s)
}

func (s SetBackupLocalResponse) GoString() string {
	return s.String()
}

func (s *SetBackupLocalResponse) SetHeaders(v map[string]*string) *SetBackupLocalResponse {
	s.Headers = v
	return s
}

func (s *SetBackupLocalResponse) SetStatusCode(v int32) *SetBackupLocalResponse {
	s.StatusCode = &v
	return s
}

func (s *SetBackupLocalResponse) SetBody(v *SetBackupLocalResponseBody) *SetBackupLocalResponse {
	s.Body = v
	return s
}

type SetBackupPolicyRequest struct {
	// The databases to be backed up. Separate multiple databases with commas (,).
	//
	// >  This parameter takes effect only when the backup level is database level.
	BackupDbNames *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	// The level of the backup. Valid values:
	//
	// *   db: The database type.
	// *   instance: instance
	BackupLevel *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	// Specifies whether to enable log Backup. Valid values:
	//
	// *   1: enabled.
	// *   0: disabled.
	BackupLog *string `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	// The backup mode. Valid values:
	//
	// *   **Logic **: logical backup
	// *   **phy**: physical backup
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The retention period of the backup data. Value range: 7 to 730.
	DataBackupRetentionPeriod *string `json:"DataBackupRetentionPeriod,omitempty" xml:"DataBackupRetentionPeriod,omitempty"`
	// The ID of the instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The log retention period. Valid values: 7 to 730. This value must be less than or equal to the number of data backup days.
	LogBackupRetentionPeriod *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	// The end time of the backup.
	PreferredBackupEndTime *string `json:"PreferredBackupEndTime,omitempty" xml:"PreferredBackupEndTime,omitempty"`
	// The backup cycle. Valid values:
	//
	// *   0: Monday
	// *   1: Tuesday
	// *   2: Wednesday
	// *   3: Thursday
	// *   4: Friday
	// *   5: Saturday
	// *   6: Sunday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The start time of the backup.
	PreferredBackupStartTime *string `json:"PreferredBackupStartTime,omitempty" xml:"PreferredBackupStartTime,omitempty"`
}

func (s SetBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetBackupPolicyRequest) SetBackupDbNames(v string) *SetBackupPolicyRequest {
	s.BackupDbNames = &v
	return s
}

func (s *SetBackupPolicyRequest) SetBackupLevel(v string) *SetBackupPolicyRequest {
	s.BackupLevel = &v
	return s
}

func (s *SetBackupPolicyRequest) SetBackupLog(v string) *SetBackupPolicyRequest {
	s.BackupLog = &v
	return s
}

func (s *SetBackupPolicyRequest) SetBackupMode(v string) *SetBackupPolicyRequest {
	s.BackupMode = &v
	return s
}

func (s *SetBackupPolicyRequest) SetDataBackupRetentionPeriod(v string) *SetBackupPolicyRequest {
	s.DataBackupRetentionPeriod = &v
	return s
}

func (s *SetBackupPolicyRequest) SetDrdsInstanceId(v string) *SetBackupPolicyRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetBackupPolicyRequest) SetLogBackupRetentionPeriod(v string) *SetBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *SetBackupPolicyRequest) SetPreferredBackupEndTime(v string) *SetBackupPolicyRequest {
	s.PreferredBackupEndTime = &v
	return s
}

func (s *SetBackupPolicyRequest) SetPreferredBackupPeriod(v string) *SetBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *SetBackupPolicyRequest) SetPreferredBackupStartTime(v string) *SetBackupPolicyRequest {
	s.PreferredBackupStartTime = &v
	return s
}

type SetBackupPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the backup policy was successfully configured.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the database creation failure records were removed from the DRDS instance.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetBackupPolicyResponseBody) SetRequestId(v string) *SetBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetBackupPolicyResponseBody) SetResult(v string) *SetBackupPolicyResponseBody {
	s.Result = &v
	return s
}

func (s *SetBackupPolicyResponseBody) SetSuccess(v bool) *SetBackupPolicyResponseBody {
	s.Success = &v
	return s
}

type SetBackupPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetBackupPolicyResponse) SetHeaders(v map[string]*string) *SetBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetBackupPolicyResponse) SetStatusCode(v int32) *SetBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetBackupPolicyResponse) SetBody(v *SetBackupPolicyResponseBody) *SetBackupPolicyResponse {
	s.Body = v
	return s
}

type SetupBroadcastTablesRequest struct {
	// Specifies whether to activate a broadcast table for the database.
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The name of the database for which you want to configure a broadcast table.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region in which the PolarDB-X 1.0 instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the table.
	TableName []*string `json:"TableName,omitempty" xml:"TableName,omitempty" type:"Repeated"`
}

func (s SetupBroadcastTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s SetupBroadcastTablesRequest) GoString() string {
	return s.String()
}

func (s *SetupBroadcastTablesRequest) SetActive(v bool) *SetupBroadcastTablesRequest {
	s.Active = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetDbName(v string) *SetupBroadcastTablesRequest {
	s.DbName = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetDrdsInstanceId(v string) *SetupBroadcastTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetRegionId(v string) *SetupBroadcastTablesRequest {
	s.RegionId = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetTableName(v []*string) *SetupBroadcastTablesRequest {
	s.TableName = v
	return s
}

type SetupBroadcastTablesResponseBody struct {
	// Indicates whether the broadcast table is configured.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetupBroadcastTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetupBroadcastTablesResponseBody) GoString() string {
	return s.String()
}

func (s *SetupBroadcastTablesResponseBody) SetData(v bool) *SetupBroadcastTablesResponseBody {
	s.Data = &v
	return s
}

func (s *SetupBroadcastTablesResponseBody) SetRequestId(v string) *SetupBroadcastTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetupBroadcastTablesResponseBody) SetSuccess(v bool) *SetupBroadcastTablesResponseBody {
	s.Success = &v
	return s
}

type SetupBroadcastTablesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetupBroadcastTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetupBroadcastTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s SetupBroadcastTablesResponse) GoString() string {
	return s.String()
}

func (s *SetupBroadcastTablesResponse) SetHeaders(v map[string]*string) *SetupBroadcastTablesResponse {
	s.Headers = v
	return s
}

func (s *SetupBroadcastTablesResponse) SetStatusCode(v int32) *SetupBroadcastTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetupBroadcastTablesResponse) SetBody(v *SetupBroadcastTablesResponseBody) *SetupBroadcastTablesResponse {
	s.Body = v
	return s
}

type SetupDrdsParamsRequest struct {
	Data []*SetupDrdsParamsRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the PolarDB-X 1.0 instance for which you want to configure parameters.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The resource for which you want to configure parameters. Valid values:
	//
	// *   **INSTANCE**: Configure parameters for the instance.
	// *   **DB**: Configure parameters for the databases of the instance.
	ParamLevel *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	// The ID of the region in which the PolarDB-X 1.0 instance is located.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetupDrdsParamsRequest) String() string {
	return tea.Prettify(s)
}

func (s SetupDrdsParamsRequest) GoString() string {
	return s.String()
}

func (s *SetupDrdsParamsRequest) SetData(v []*SetupDrdsParamsRequestData) *SetupDrdsParamsRequest {
	s.Data = v
	return s
}

func (s *SetupDrdsParamsRequest) SetDrdsInstanceId(v string) *SetupDrdsParamsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetupDrdsParamsRequest) SetParamLevel(v string) *SetupDrdsParamsRequest {
	s.ParamLevel = &v
	return s
}

func (s *SetupDrdsParamsRequest) SetRegionId(v string) *SetupDrdsParamsRequest {
	s.RegionId = &v
	return s
}

type SetupDrdsParamsRequestData struct {
	// The name of the parameter that you want to configure for a database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The valid values of the parameter.
	ParamRanges *string `json:"ParamRanges,omitempty" xml:"ParamRanges,omitempty"`
	// The type of the parameter that you want to configure. Valid values:
	//
	// *   **ATOM**: the configuration item in the layer-3 data source.
	// *   **CONFIG**: the configuration item in ConfigServer.
	// *   **DIAMOND**: the configuration item in Diamond.
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The value of parameter that you want to configure.
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
	// The name of the parameter that you want to configure.
	ParamVariableName *string `json:"ParamVariableName,omitempty" xml:"ParamVariableName,omitempty"`
}

func (s SetupDrdsParamsRequestData) String() string {
	return tea.Prettify(s)
}

func (s SetupDrdsParamsRequestData) GoString() string {
	return s.String()
}

func (s *SetupDrdsParamsRequestData) SetDbName(v string) *SetupDrdsParamsRequestData {
	s.DbName = &v
	return s
}

func (s *SetupDrdsParamsRequestData) SetParamRanges(v string) *SetupDrdsParamsRequestData {
	s.ParamRanges = &v
	return s
}

func (s *SetupDrdsParamsRequestData) SetParamType(v string) *SetupDrdsParamsRequestData {
	s.ParamType = &v
	return s
}

func (s *SetupDrdsParamsRequestData) SetParamValue(v string) *SetupDrdsParamsRequestData {
	s.ParamValue = &v
	return s
}

func (s *SetupDrdsParamsRequestData) SetParamVariableName(v string) *SetupDrdsParamsRequestData {
	s.ParamVariableName = &v
	return s
}

type SetupDrdsParamsResponseBody struct {
	// The returned results.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetupDrdsParamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetupDrdsParamsResponseBody) GoString() string {
	return s.String()
}

func (s *SetupDrdsParamsResponseBody) SetData(v bool) *SetupDrdsParamsResponseBody {
	s.Data = &v
	return s
}

func (s *SetupDrdsParamsResponseBody) SetRequestId(v string) *SetupDrdsParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetupDrdsParamsResponseBody) SetSuccess(v bool) *SetupDrdsParamsResponseBody {
	s.Success = &v
	return s
}

type SetupDrdsParamsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetupDrdsParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetupDrdsParamsResponse) String() string {
	return tea.Prettify(s)
}

func (s SetupDrdsParamsResponse) GoString() string {
	return s.String()
}

func (s *SetupDrdsParamsResponse) SetHeaders(v map[string]*string) *SetupDrdsParamsResponse {
	s.Headers = v
	return s
}

func (s *SetupDrdsParamsResponse) SetStatusCode(v int32) *SetupDrdsParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetupDrdsParamsResponse) SetBody(v *SetupDrdsParamsResponseBody) *SetupDrdsParamsResponse {
	s.Body = v
	return s
}

type SetupRecycleBinStatusRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies the status of the table recycle bin. Valid values:
	//
	// *   enable: The table recycle bin is enabled.
	// *   disable: The table recycle bin is disabled.
	StatusAction *string `json:"StatusAction,omitempty" xml:"StatusAction,omitempty"`
}

func (s SetupRecycleBinStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetupRecycleBinStatusRequest) GoString() string {
	return s.String()
}

func (s *SetupRecycleBinStatusRequest) SetDbName(v string) *SetupRecycleBinStatusRequest {
	s.DbName = &v
	return s
}

func (s *SetupRecycleBinStatusRequest) SetDrdsInstanceId(v string) *SetupRecycleBinStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetupRecycleBinStatusRequest) SetRegionId(v string) *SetupRecycleBinStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetupRecycleBinStatusRequest) SetStatusAction(v string) *SetupRecycleBinStatusRequest {
	s.StatusAction = &v
	return s
}

type SetupRecycleBinStatusResponseBody struct {
	// Indicates whether the table recycle bin is enabled.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetupRecycleBinStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetupRecycleBinStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetupRecycleBinStatusResponseBody) SetData(v bool) *SetupRecycleBinStatusResponseBody {
	s.Data = &v
	return s
}

func (s *SetupRecycleBinStatusResponseBody) SetRequestId(v string) *SetupRecycleBinStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetupRecycleBinStatusResponseBody) SetSuccess(v bool) *SetupRecycleBinStatusResponseBody {
	s.Success = &v
	return s
}

type SetupRecycleBinStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetupRecycleBinStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetupRecycleBinStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetupRecycleBinStatusResponse) GoString() string {
	return s.String()
}

func (s *SetupRecycleBinStatusResponse) SetHeaders(v map[string]*string) *SetupRecycleBinStatusResponse {
	s.Headers = v
	return s
}

func (s *SetupRecycleBinStatusResponse) SetStatusCode(v int32) *SetupRecycleBinStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetupRecycleBinStatusResponse) SetBody(v *SetupRecycleBinStatusResponseBody) *SetupRecycleBinStatusResponse {
	s.Body = v
	return s
}

type SetupTableRequest struct {
	// Specifies whether to enable full table scan.
	AllowFullTableScan *bool `json:"AllowFullTableScan,omitempty" xml:"AllowFullTableScan,omitempty"`
	// The name of the database in which the table resides.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region where the streaming domain resides.
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TableName []*string `json:"TableName,omitempty" xml:"TableName,omitempty" type:"Repeated"`
}

func (s SetupTableRequest) String() string {
	return tea.Prettify(s)
}

func (s SetupTableRequest) GoString() string {
	return s.String()
}

func (s *SetupTableRequest) SetAllowFullTableScan(v bool) *SetupTableRequest {
	s.AllowFullTableScan = &v
	return s
}

func (s *SetupTableRequest) SetDbName(v string) *SetupTableRequest {
	s.DbName = &v
	return s
}

func (s *SetupTableRequest) SetDrdsInstanceId(v string) *SetupTableRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetupTableRequest) SetRegionId(v string) *SetupTableRequest {
	s.RegionId = &v
	return s
}

func (s *SetupTableRequest) SetTableName(v []*string) *SetupTableRequest {
	s.TableName = v
	return s
}

type SetupTableResponseBody struct {
	// Specifies whether to use a full table scan.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetupTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetupTableResponseBody) GoString() string {
	return s.String()
}

func (s *SetupTableResponseBody) SetData(v bool) *SetupTableResponseBody {
	s.Data = &v
	return s
}

func (s *SetupTableResponseBody) SetRequestId(v string) *SetupTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetupTableResponseBody) SetSuccess(v bool) *SetupTableResponseBody {
	s.Success = &v
	return s
}

type SetupTableResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetupTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetupTableResponse) String() string {
	return tea.Prettify(s)
}

func (s SetupTableResponse) GoString() string {
	return s.String()
}

func (s *SetupTableResponse) SetHeaders(v map[string]*string) *SetupTableResponse {
	s.Headers = v
	return s
}

func (s *SetupTableResponse) SetStatusCode(v int32) *SetupTableResponse {
	s.StatusCode = &v
	return s
}

func (s *SetupTableResponse) SetBody(v *SetupTableResponseBody) *SetupTableResponse {
	s.Body = v
	return s
}

type StartRestoreRequest struct {
	// The name of the database to be restored. Separate multiple databases with commas (,).
	//
	// >  If you do not specify any database name, all databases in the instance are restored by default.
	BackupDbNames *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	// The ID of the DRDS backup set.
	//
	// >  If you do not specify this parameter, the system restores data by time (PreferredBackupTime).
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The level of the backup. Valid values:
	//
	// *   db: The database level.
	// *   instance: the instance level.
	BackupLevel *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	// The backup method. Valid values:
	//
	// *   logic: the logical backup.
	// *   phy: fast backup
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The ID of the DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The restoration time of the instance, in the format of`  yyyy-MM-dd HH:mm:ss `.
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
}

func (s StartRestoreRequest) String() string {
	return tea.Prettify(s)
}

func (s StartRestoreRequest) GoString() string {
	return s.String()
}

func (s *StartRestoreRequest) SetBackupDbNames(v string) *StartRestoreRequest {
	s.BackupDbNames = &v
	return s
}

func (s *StartRestoreRequest) SetBackupId(v string) *StartRestoreRequest {
	s.BackupId = &v
	return s
}

func (s *StartRestoreRequest) SetBackupLevel(v string) *StartRestoreRequest {
	s.BackupLevel = &v
	return s
}

func (s *StartRestoreRequest) SetBackupMode(v string) *StartRestoreRequest {
	s.BackupMode = &v
	return s
}

func (s *StartRestoreRequest) SetDrdsInstanceId(v string) *StartRestoreRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *StartRestoreRequest) SetPreferredBackupTime(v string) *StartRestoreRequest {
	s.PreferredBackupTime = &v
	return s
}

type StartRestoreResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether SQL audit was disabled for the DRDS database.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartRestoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartRestoreResponseBody) GoString() string {
	return s.String()
}

func (s *StartRestoreResponseBody) SetRequestId(v string) *StartRestoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRestoreResponseBody) SetResult(v string) *StartRestoreResponseBody {
	s.Result = &v
	return s
}

func (s *StartRestoreResponseBody) SetSuccess(v bool) *StartRestoreResponseBody {
	s.Success = &v
	return s
}

type StartRestoreResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRestoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRestoreResponse) String() string {
	return tea.Prettify(s)
}

func (s StartRestoreResponse) GoString() string {
	return s.String()
}

func (s *StartRestoreResponse) SetHeaders(v map[string]*string) *StartRestoreResponse {
	s.Headers = v
	return s
}

func (s *StartRestoreResponse) SetStatusCode(v int32) *StartRestoreResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRestoreResponse) SetBody(v *StartRestoreResponseBody) *StartRestoreResponse {
	s.Body = v
	return s
}

type SubmitCleanTaskRequest struct {
	// The name of the database that is scaled out.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The scale-out type. Valid values:
	//
	// *   smooth_expand: smooth scale-out
	// *   hot_expand: hot-spot scale-out
	ExpandType *string `json:"ExpandType,omitempty" xml:"ExpandType,omitempty"`
	// The job ID of the scale-out task. The value of this parameter is the same as that of the ParentJobId parameter.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the scale-out task. This parameter is returned if you send a request for the smooth scale-out task.
	ParentJobId *string `json:"ParentJobId,omitempty" xml:"ParentJobId,omitempty"`
}

func (s SubmitCleanTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitCleanTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitCleanTaskRequest) SetDbName(v string) *SubmitCleanTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetDrdsInstanceId(v string) *SubmitCleanTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetExpandType(v string) *SubmitCleanTaskRequest {
	s.ExpandType = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetJobId(v string) *SubmitCleanTaskRequest {
	s.JobId = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetParentJobId(v string) *SubmitCleanTaskRequest {
	s.ParentJobId = &v
	return s
}

type SubmitCleanTaskResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitCleanTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitCleanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCleanTaskResponseBody) SetRequestId(v string) *SubmitCleanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCleanTaskResponseBody) SetSuccess(v bool) *SubmitCleanTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitCleanTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitCleanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitCleanTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitCleanTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitCleanTaskResponse) SetHeaders(v map[string]*string) *SubmitCleanTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitCleanTaskResponse) SetStatusCode(v int32) *SubmitCleanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCleanTaskResponse) SetBody(v *SubmitCleanTaskResponseBody) *SubmitCleanTaskResponse {
	s.Body = v
	return s
}

type SubmitHotExpandPreCheckTaskRequest struct {
	// The type of the database. Valid values:
	//
	// *   RDS
	// *   PolarDB
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The name of the PolarDB-X database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The name of the table.
	TableList []*string `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Repeated"`
}

func (s SubmitHotExpandPreCheckTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandPreCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetDbInstType(v string) *SubmitHotExpandPreCheckTaskRequest {
	s.DbInstType = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetDbName(v string) *SubmitHotExpandPreCheckTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetDrdsInstanceId(v string) *SubmitHotExpandPreCheckTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetTableList(v []*string) *SubmitHotExpandPreCheckTaskRequest {
	s.TableList = v
	return s
}

type SubmitHotExpandPreCheckTaskResponseBody struct {
	// The result of the task.
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the task.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitHotExpandPreCheckTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandPreCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) SetMsg(v string) *SubmitHotExpandPreCheckTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) SetRequestId(v string) *SubmitHotExpandPreCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) SetSuccess(v bool) *SubmitHotExpandPreCheckTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) SetTaskId(v int64) *SubmitHotExpandPreCheckTaskResponseBody {
	s.TaskId = &v
	return s
}

type SubmitHotExpandPreCheckTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitHotExpandPreCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitHotExpandPreCheckTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandPreCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandPreCheckTaskResponse) SetHeaders(v map[string]*string) *SubmitHotExpandPreCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponse) SetStatusCode(v int32) *SubmitHotExpandPreCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponse) SetBody(v *SubmitHotExpandPreCheckTaskResponseBody) *SubmitHotExpandPreCheckTaskResponse {
	s.Body = v
	return s
}

type SubmitHotExpandTaskRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The information about the database on which you want to perform hot-spot scale-out.
	ExtendedMapping []*SubmitHotExpandTaskRequestExtendedMapping `json:"ExtendedMapping,omitempty" xml:"ExtendedMapping,omitempty" type:"Repeated"`
	// The information about the instance to which the hot-spot database belongs.
	InstanceDbMapping []*SubmitHotExpandTaskRequestInstanceDbMapping `json:"InstanceDbMapping,omitempty" xml:"InstanceDbMapping,omitempty" type:"Repeated"`
	// The information about the hot-spot database.
	Mapping []*SubmitHotExpandTaskRequestMapping `json:"Mapping,omitempty" xml:"Mapping,omitempty" type:"Repeated"`
	// The information about the privileged account.
	SupperAccountMapping []*SubmitHotExpandTaskRequestSupperAccountMapping `json:"SupperAccountMapping,omitempty" xml:"SupperAccountMapping,omitempty" type:"Repeated"`
	// The description of the task.
	TaskDesc *string `json:"TaskDesc,omitempty" xml:"TaskDesc,omitempty"`
	// The name of the task.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s SubmitHotExpandTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequest) SetDbName(v string) *SubmitHotExpandTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetDrdsInstanceId(v string) *SubmitHotExpandTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetExtendedMapping(v []*SubmitHotExpandTaskRequestExtendedMapping) *SubmitHotExpandTaskRequest {
	s.ExtendedMapping = v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetInstanceDbMapping(v []*SubmitHotExpandTaskRequestInstanceDbMapping) *SubmitHotExpandTaskRequest {
	s.InstanceDbMapping = v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetMapping(v []*SubmitHotExpandTaskRequestMapping) *SubmitHotExpandTaskRequest {
	s.Mapping = v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetSupperAccountMapping(v []*SubmitHotExpandTaskRequestSupperAccountMapping) *SubmitHotExpandTaskRequest {
	s.SupperAccountMapping = v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetTaskDesc(v string) *SubmitHotExpandTaskRequest {
	s.TaskDesc = &v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetTaskName(v string) *SubmitHotExpandTaskRequest {
	s.TaskName = &v
	return s
}

type SubmitHotExpandTaskRequestExtendedMapping struct {
	// The name of the source physical database.
	SrcDb *string `json:"SrcDb,omitempty" xml:"SrcDb,omitempty"`
	// The ID of the ApsaraDB RDS instance to which the source physical database belongs.
	SrcInstanceId *string `json:"SrcInstanceId,omitempty" xml:"SrcInstanceId,omitempty"`
}

func (s SubmitHotExpandTaskRequestExtendedMapping) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskRequestExtendedMapping) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequestExtendedMapping) SetSrcDb(v string) *SubmitHotExpandTaskRequestExtendedMapping {
	s.SrcDb = &v
	return s
}

func (s *SubmitHotExpandTaskRequestExtendedMapping) SetSrcInstanceId(v string) *SubmitHotExpandTaskRequestExtendedMapping {
	s.SrcInstanceId = &v
	return s
}

type SubmitHotExpandTaskRequestInstanceDbMapping struct {
	// The name of the hot-spot database.
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The name of the ApsaraDB RDS instance to which the hot-spot database belongs.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s SubmitHotExpandTaskRequestInstanceDbMapping) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskRequestInstanceDbMapping) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequestInstanceDbMapping) SetDbList(v string) *SubmitHotExpandTaskRequestInstanceDbMapping {
	s.DbList = &v
	return s
}

func (s *SubmitHotExpandTaskRequestInstanceDbMapping) SetInstanceName(v string) *SubmitHotExpandTaskRequestInstanceDbMapping {
	s.InstanceName = &v
	return s
}

type SubmitHotExpandTaskRequestMapping struct {
	// The shard key used to split the database to which the associated table belongs.
	DbShardColumn *string `json:"DbShardColumn,omitempty" xml:"DbShardColumn,omitempty"`
	// The name of the hot-spot database.
	HotDbName *string `json:"HotDbName,omitempty" xml:"HotDbName,omitempty"`
	// The name of the hot-spot table. The name must be prefixed with the name of a logical table.
	HotTableName *string `json:"HotTableName,omitempty" xml:"HotTableName,omitempty"`
	// The name of the logical table on which you want to perform hot-spot scale-out.
	LogicTable *string `json:"LogicTable,omitempty" xml:"LogicTable,omitempty"`
	// The value of the shard key used to split a database.
	ShardDbValue *string `json:"ShardDbValue,omitempty" xml:"ShardDbValue,omitempty"`
	// The value of the shard key used to split a table.
	ShardTbValue *string `json:"ShardTbValue,omitempty" xml:"ShardTbValue,omitempty"`
	// The shard key used to split an associated table.
	TbShardColumn *string `json:"TbShardColumn,omitempty" xml:"TbShardColumn,omitempty"`
}

func (s SubmitHotExpandTaskRequestMapping) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskRequestMapping) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequestMapping) SetDbShardColumn(v string) *SubmitHotExpandTaskRequestMapping {
	s.DbShardColumn = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetHotDbName(v string) *SubmitHotExpandTaskRequestMapping {
	s.HotDbName = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetHotTableName(v string) *SubmitHotExpandTaskRequestMapping {
	s.HotTableName = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetLogicTable(v string) *SubmitHotExpandTaskRequestMapping {
	s.LogicTable = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetShardDbValue(v string) *SubmitHotExpandTaskRequestMapping {
	s.ShardDbValue = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetShardTbValue(v string) *SubmitHotExpandTaskRequestMapping {
	s.ShardTbValue = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetTbShardColumn(v string) *SubmitHotExpandTaskRequestMapping {
	s.TbShardColumn = &v
	return s
}

type SubmitHotExpandTaskRequestSupperAccountMapping struct {
	// The ID of the ApsaraDB RDS instance that has the privileged account.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The name of the privileged account of the ApsaraDB RDS instance.
	SupperAccount *string `json:"SupperAccount,omitempty" xml:"SupperAccount,omitempty"`
	// The password of the privileged account of the ApsaraDB RDS instance.
	SupperPassword *string `json:"SupperPassword,omitempty" xml:"SupperPassword,omitempty"`
}

func (s SubmitHotExpandTaskRequestSupperAccountMapping) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskRequestSupperAccountMapping) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequestSupperAccountMapping) SetInstanceName(v string) *SubmitHotExpandTaskRequestSupperAccountMapping {
	s.InstanceName = &v
	return s
}

func (s *SubmitHotExpandTaskRequestSupperAccountMapping) SetSupperAccount(v string) *SubmitHotExpandTaskRequestSupperAccountMapping {
	s.SupperAccount = &v
	return s
}

func (s *SubmitHotExpandTaskRequestSupperAccountMapping) SetSupperPassword(v string) *SubmitHotExpandTaskRequestSupperAccountMapping {
	s.SupperPassword = &v
	return s
}

type SubmitHotExpandTaskResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitHotExpandTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskResponseBody) SetRequestId(v string) *SubmitHotExpandTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitHotExpandTaskResponseBody) SetSuccess(v bool) *SubmitHotExpandTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitHotExpandTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitHotExpandTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitHotExpandTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskResponse) SetHeaders(v map[string]*string) *SubmitHotExpandTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitHotExpandTaskResponse) SetStatusCode(v int32) *SubmitHotExpandTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitHotExpandTaskResponse) SetBody(v *SubmitHotExpandTaskResponseBody) *SubmitHotExpandTaskResponse {
	s.Body = v
	return s
}

type SubmitSmoothExpandPreCheckRequest struct {
	// The type of the database. Valid values:
	//
	// *   RDS
	// *   POLARDB
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The name of the PolarDB-X database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s SubmitSmoothExpandPreCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckRequest) SetDbInstType(v string) *SubmitSmoothExpandPreCheckRequest {
	s.DbInstType = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckRequest) SetDbName(v string) *SubmitSmoothExpandPreCheckRequest {
	s.DbName = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckRequest) SetDrdsInstanceId(v string) *SubmitSmoothExpandPreCheckRequest {
	s.DrdsInstanceId = &v
	return s
}

type SubmitSmoothExpandPreCheckResponseBody struct {
	// The result of the precheck task.
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the precheck task.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitSmoothExpandPreCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckResponseBody) SetMsg(v string) *SubmitSmoothExpandPreCheckResponseBody {
	s.Msg = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponseBody) SetRequestId(v string) *SubmitSmoothExpandPreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponseBody) SetSuccess(v bool) *SubmitSmoothExpandPreCheckResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponseBody) SetTaskId(v int64) *SubmitSmoothExpandPreCheckResponseBody {
	s.TaskId = &v
	return s
}

type SubmitSmoothExpandPreCheckResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSmoothExpandPreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSmoothExpandPreCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckResponse) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckResponse) SetHeaders(v map[string]*string) *SubmitSmoothExpandPreCheckResponse {
	s.Headers = v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponse) SetStatusCode(v int32) *SubmitSmoothExpandPreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponse) SetBody(v *SubmitSmoothExpandPreCheckResponseBody) *SubmitSmoothExpandPreCheckResponse {
	s.Body = v
	return s
}

type SubmitSmoothExpandPreCheckTaskRequest struct {
	// The name of the PolarDB-X 1.0 database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s SubmitSmoothExpandPreCheckTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckTaskRequest) SetDbName(v string) *SubmitSmoothExpandPreCheckTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskRequest) SetDrdsInstanceId(v string) *SubmitSmoothExpandPreCheckTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

type SubmitSmoothExpandPreCheckTaskResponseBody struct {
	// Indicates whether the precheck task was submitted.
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the task.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitSmoothExpandPreCheckTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) SetMsg(v string) *SubmitSmoothExpandPreCheckTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) SetRequestId(v string) *SubmitSmoothExpandPreCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) SetSuccess(v bool) *SubmitSmoothExpandPreCheckTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) SetTaskId(v int64) *SubmitSmoothExpandPreCheckTaskResponseBody {
	s.TaskId = &v
	return s
}

type SubmitSmoothExpandPreCheckTaskResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSmoothExpandPreCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSmoothExpandPreCheckTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckTaskResponse) SetHeaders(v map[string]*string) *SubmitSmoothExpandPreCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponse) SetStatusCode(v int32) *SubmitSmoothExpandPreCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponse) SetBody(v *SubmitSmoothExpandPreCheckTaskResponseBody) *SubmitSmoothExpandPreCheckTaskResponse {
	s.Body = v
	return s
}

type SubmitSqlFlashbackTaskRequest struct {
	// The name of the DRDS database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of a DRDS instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The time when the SQL flashback task ends.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The restoration type. Valid values:
	//
	// *   1: Image restoration
	// *   0: reverse recovery
	RecallRestoreType *int32 `json:"RecallRestoreType,omitempty" xml:"RecallRestoreType,omitempty"`
	// Exact match or fuzzy match. Valid values:
	//
	// *   0: the exact match.
	// *   1: the fuzzy match.
	RecallType *int32 `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// The primary key of flashback SQL.
	SqlPk *string `json:"SqlPk,omitempty" xml:"SqlPk,omitempty"`
	// The type of the SQL statement. Valid values: INSERT, UPDATE, and DELETE. Separate multiple types with commas (,).
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The start time of the flashback SQL statement.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the table where the flashback SQL operation was performed.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The Trace ID of the flashback SQL.
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s SubmitSqlFlashbackTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSqlFlashbackTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitSqlFlashbackTaskRequest) SetDbName(v string) *SubmitSqlFlashbackTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetDrdsInstanceId(v string) *SubmitSqlFlashbackTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetEndTime(v string) *SubmitSqlFlashbackTaskRequest {
	s.EndTime = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetRecallRestoreType(v int32) *SubmitSqlFlashbackTaskRequest {
	s.RecallRestoreType = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetRecallType(v int32) *SubmitSqlFlashbackTaskRequest {
	s.RecallType = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetSqlPk(v string) *SubmitSqlFlashbackTaskRequest {
	s.SqlPk = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetSqlType(v string) *SubmitSqlFlashbackTaskRequest {
	s.SqlType = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetStartTime(v string) *SubmitSqlFlashbackTaskRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetTableName(v string) *SubmitSqlFlashbackTaskRequest {
	s.TableName = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetTraceId(v string) *SubmitSqlFlashbackTaskRequest {
	s.TraceId = &v
	return s
}

type SubmitSqlFlashbackTaskResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the database creation failure records were removed from the DRDS instance.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the replication task.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitSqlFlashbackTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSqlFlashbackTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSqlFlashbackTaskResponseBody) SetRequestId(v string) *SubmitSqlFlashbackTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSqlFlashbackTaskResponseBody) SetSuccess(v bool) *SubmitSqlFlashbackTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSqlFlashbackTaskResponseBody) SetTaskId(v int64) *SubmitSqlFlashbackTaskResponseBody {
	s.TaskId = &v
	return s
}

type SubmitSqlFlashbackTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSqlFlashbackTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSqlFlashbackTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSqlFlashbackTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitSqlFlashbackTaskResponse) SetHeaders(v map[string]*string) *SubmitSqlFlashbackTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitSqlFlashbackTaskResponse) SetStatusCode(v int32) *SubmitSqlFlashbackTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSqlFlashbackTaskResponse) SetBody(v *SubmitSqlFlashbackTaskResponseBody) *SubmitSqlFlashbackTaskResponse {
	s.Body = v
	return s
}

type SwitchGlobalBroadcastTypeRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SwitchGlobalBroadcastTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchGlobalBroadcastTypeRequest) GoString() string {
	return s.String()
}

func (s *SwitchGlobalBroadcastTypeRequest) SetDbName(v string) *SwitchGlobalBroadcastTypeRequest {
	s.DbName = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeRequest) SetDrdsInstanceId(v string) *SwitchGlobalBroadcastTypeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeRequest) SetRegionId(v string) *SwitchGlobalBroadcastTypeRequest {
	s.RegionId = &v
	return s
}

type SwitchGlobalBroadcastTypeResponseBody struct {
	// Indicates whether the mode of broadcast tables was switched from the multi-write mode to the asynchronous link mode.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SwitchGlobalBroadcastTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchGlobalBroadcastTypeResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchGlobalBroadcastTypeResponseBody) SetData(v bool) *SwitchGlobalBroadcastTypeResponseBody {
	s.Data = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeResponseBody) SetRequestId(v string) *SwitchGlobalBroadcastTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeResponseBody) SetSuccess(v bool) *SwitchGlobalBroadcastTypeResponseBody {
	s.Success = &v
	return s
}

type SwitchGlobalBroadcastTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchGlobalBroadcastTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchGlobalBroadcastTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchGlobalBroadcastTypeResponse) GoString() string {
	return s.String()
}

func (s *SwitchGlobalBroadcastTypeResponse) SetHeaders(v map[string]*string) *SwitchGlobalBroadcastTypeResponse {
	s.Headers = v
	return s
}

func (s *SwitchGlobalBroadcastTypeResponse) SetStatusCode(v int32) *SwitchGlobalBroadcastTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeResponse) SetBody(v *SwitchGlobalBroadcastTypeResponseBody) *SwitchGlobalBroadcastTypeResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The ID of the region in which the resource is located.
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Set the value to INSTANCE.
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
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
	// The key of the tag that you want to add.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that you want to add.
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UntagResourcesRequest struct {
	// Specifies whether to delete all tags of the resource.
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region ID of the instance.
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to INSTANCE.
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the database creation failure records were removed from the DRDS instance.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UpdateInstanceNetworkRequest struct {
	// Specifies the retention period of the classic network endpoint. Unit: days.
	ClassicExpiredDays *int32 `json:"ClassicExpiredDays,omitempty" xml:"ClassicExpiredDays,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// Specifies whether to retain the classic network endpoint.
	RetainClassic *bool `json:"RetainClassic,omitempty" xml:"RetainClassic,omitempty"`
	// The network type of the PolarDB-X 1.0 instance. Valid values:
	//
	// *   vpc: Virtual Private Cloud (VPC)
	// *   classic: classic network
	SrcInstanceNetworkType *string `json:"SrcInstanceNetworkType,omitempty" xml:"SrcInstanceNetworkType,omitempty"`
}

func (s UpdateInstanceNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNetworkRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkRequest) SetClassicExpiredDays(v int32) *UpdateInstanceNetworkRequest {
	s.ClassicExpiredDays = &v
	return s
}

func (s *UpdateInstanceNetworkRequest) SetDrdsInstanceId(v string) *UpdateInstanceNetworkRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpdateInstanceNetworkRequest) SetRetainClassic(v bool) *UpdateInstanceNetworkRequest {
	s.RetainClassic = &v
	return s
}

func (s *UpdateInstanceNetworkRequest) SetSrcInstanceNetworkType(v string) *UpdateInstanceNetworkRequest {
	s.SrcInstanceNetworkType = &v
	return s
}

type UpdateInstanceNetworkResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkResponseBody) SetRequestId(v string) *UpdateInstanceNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceNetworkResponseBody) SetSuccess(v bool) *UpdateInstanceNetworkResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceNetworkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNetworkResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkResponse) SetHeaders(v map[string]*string) *UpdateInstanceNetworkResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceNetworkResponse) SetStatusCode(v int32) *UpdateInstanceNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceNetworkResponse) SetBody(v *UpdateInstanceNetworkResponseBody) *UpdateInstanceNetworkResponse {
	s.Body = v
	return s
}

type UpdatePrivateRdsClassRequest struct {
	// Specifies whether to use vouchers to offset the purchase fees. Valid values: **true** and **false**. Default value: false.
	//
	// > If you downgrade the specifications of an instance after you use the vouchers, the vouchers used for the purchase cannot be refunded.
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The ID of the custom ApsaraDB RDS instance at the storage layer.
	//
	// > You can call the [DescribeDrdsRdsInstances](~~xxxx~~) operation to query the details of all ApsaraDB RDS instances at the storage layer of a PolarDB-X 1.0 instance, including the IDs of the ApsaraDB RDS instances.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// > You can call the [DescribeDrdsInstances](~~139284~~) operation to query the details of all PolarDB-X 1.0 instances within an Alibaba Cloud account, including the IDs of the instances.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// This parameter is discontinued.
	PrePayDuration *int32 `json:"PrePayDuration,omitempty" xml:"PrePayDuration,omitempty"`
	// The new instance type of the custom ApsaraDB RDS instance at the storage layer.
	//
	// > You can call the [DescribeAvailableClasses](~~196546~~) operation to view the specifications that are supported for a custom ApsaraDB RDS instance. The specifications include the instance type and the storage capacity.
	RdsClass *string `json:"RdsClass,omitempty" xml:"RdsClass,omitempty"`
	// The new storage capacity of the custom ApsaraDB RDS instance at the storage layer.
	//
	// > You can call the [DescribeAvailableClasses](~~196546~~) operation to view the specifications that are supported for a custom ApsaraDB RDS instance. The specifications include the instance type and the storage capacity.
	Storage *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s UpdatePrivateRdsClassRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateRdsClassRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateRdsClassRequest) SetAutoUseCoupon(v bool) *UpdatePrivateRdsClassRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetDBInstanceId(v string) *UpdatePrivateRdsClassRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetDrdsInstanceId(v string) *UpdatePrivateRdsClassRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetPrePayDuration(v int32) *UpdatePrivateRdsClassRequest {
	s.PrePayDuration = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetRdsClass(v string) *UpdatePrivateRdsClassRequest {
	s.RdsClass = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetStorage(v string) *UpdatePrivateRdsClassRequest {
	s.Storage = &v
	return s
}

type UpdatePrivateRdsClassResponseBody struct {
	// The ID of the order.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePrivateRdsClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateRdsClassResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivateRdsClassResponseBody) SetData(v string) *UpdatePrivateRdsClassResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePrivateRdsClassResponseBody) SetRequestId(v string) *UpdatePrivateRdsClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrivateRdsClassResponseBody) SetSuccess(v bool) *UpdatePrivateRdsClassResponseBody {
	s.Success = &v
	return s
}

type UpdatePrivateRdsClassResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrivateRdsClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrivateRdsClassResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateRdsClassResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivateRdsClassResponse) SetHeaders(v map[string]*string) *UpdatePrivateRdsClassResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivateRdsClassResponse) SetStatusCode(v int32) *UpdatePrivateRdsClassResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrivateRdsClassResponse) SetBody(v *UpdatePrivateRdsClassResponseBody) *UpdatePrivateRdsClassResponse {
	s.Body = v
	return s
}

type UpdateResourceGroupAttributeRequest struct {
	// The ID of the instance that you want to transfer.
	//
	// >  You can call the [DescribeDrdsInstances](~~139284~~) operation to view the details of the instances under the account, including the instance IDs.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the resource group that you want to specify.
	//
	// >  You can call the [ListResourceGroups](~~158855~~) operation to view the details of the resource groups, including the resource group IDs.
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The ID of the region where the instance you want to transfer is located.
	//
	// >  You can call the [DescribeDrdsInstances](~~139284~~) operation to view the details of the instances under the account, including the region IDs.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateResourceGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupAttributeRequest) SetDrdsInstanceId(v string) *UpdateResourceGroupAttributeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpdateResourceGroupAttributeRequest) SetNewResourceGroupId(v string) *UpdateResourceGroupAttributeRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *UpdateResourceGroupAttributeRequest) SetRegionId(v string) *UpdateResourceGroupAttributeRequest {
	s.RegionId = &v
	return s
}

type UpdateResourceGroupAttributeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupAttributeResponseBody) SetRequestId(v string) *UpdateResourceGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResourceGroupAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupAttributeResponse) SetHeaders(v map[string]*string) *UpdateResourceGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceGroupAttributeResponse) SetStatusCode(v int32) *UpdateResourceGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceGroupAttributeResponse) SetBody(v *UpdateResourceGroupAttributeResponseBody) *UpdateResourceGroupAttributeResponse {
	s.Body = v
	return s
}

type UpgradeHiStoreInstanceRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the column-oriented storage instance.
	HistoreInstanceId *string `json:"HistoreInstanceId,omitempty" xml:"HistoreInstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpgradeHiStoreInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeHiStoreInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeHiStoreInstanceRequest) SetDrdsInstanceId(v string) *UpgradeHiStoreInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpgradeHiStoreInstanceRequest) SetHistoreInstanceId(v string) *UpgradeHiStoreInstanceRequest {
	s.HistoreInstanceId = &v
	return s
}

func (s *UpgradeHiStoreInstanceRequest) SetRegionId(v string) *UpgradeHiStoreInstanceRequest {
	s.RegionId = &v
	return s
}

type UpgradeHiStoreInstanceResponseBody struct {
	// Indicates whether the request was successful. A value of true indicates that the request was successful. An error message was returned if the request failed.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeHiStoreInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeHiStoreInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeHiStoreInstanceResponseBody) SetData(v string) *UpgradeHiStoreInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *UpgradeHiStoreInstanceResponseBody) SetRequestId(v string) *UpgradeHiStoreInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeHiStoreInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeHiStoreInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeHiStoreInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeHiStoreInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeHiStoreInstanceResponse) SetHeaders(v map[string]*string) *UpgradeHiStoreInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeHiStoreInstanceResponse) SetStatusCode(v int32) *UpgradeHiStoreInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeHiStoreInstanceResponse) SetBody(v *UpgradeHiStoreInstanceResponseBody) *UpgradeHiStoreInstanceResponse {
	s.Body = v
	return s
}

type UpgradeInstanceVersionRequest struct {
	// The ID of the PolarDB-X 1.0 instance that you want to upgrade.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The version number of the PolarDB-X 1.0 instance. You can leave this parameter unspecified.
	Rpm *string `json:"Rpm,omitempty" xml:"Rpm,omitempty"`
}

func (s UpgradeInstanceVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionRequest) SetDrdsInstanceId(v string) *UpgradeInstanceVersionRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpgradeInstanceVersionRequest) SetRegionId(v string) *UpgradeInstanceVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeInstanceVersionRequest) SetRpm(v string) *UpgradeInstanceVersionRequest {
	s.Rpm = &v
	return s
}

type UpgradeInstanceVersionResponseBody struct {
	// The result of the request.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeInstanceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionResponseBody) SetData(v string) *UpgradeInstanceVersionResponseBody {
	s.Data = &v
	return s
}

func (s *UpgradeInstanceVersionResponseBody) SetRequestId(v string) *UpgradeInstanceVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeInstanceVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeInstanceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeInstanceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionResponse) SetHeaders(v map[string]*string) *UpgradeInstanceVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeInstanceVersionResponse) SetStatusCode(v int32) *UpgradeInstanceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeInstanceVersionResponse) SetBody(v *UpgradeInstanceVersionResponseBody) *UpgradeInstanceVersionResponse {
	s.Body = v
	return s
}

type ValidateShardTaskRequest struct {
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region where the PolarDB-X 1.0 instance is created.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the table or table shard on which you want to perform the task.
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	// The name of the table or table shard on which you perform the task.
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	// The type of the task. Valid values:
	//
	// *   **SINGLE_TO_SHARD**: converts a single table to a table shard.
	// *   **SHARD_TO_SINGLE**: converts a table shard to a single table.
	// *   **SHARD_TO_SHARD**: converts a table shard to another table shard.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ValidateShardTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateShardTaskRequest) GoString() string {
	return s.String()
}

func (s *ValidateShardTaskRequest) SetDbName(v string) *ValidateShardTaskRequest {
	s.DbName = &v
	return s
}

func (s *ValidateShardTaskRequest) SetDrdsInstanceId(v string) *ValidateShardTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ValidateShardTaskRequest) SetRegionId(v string) *ValidateShardTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ValidateShardTaskRequest) SetSourceTableName(v string) *ValidateShardTaskRequest {
	s.SourceTableName = &v
	return s
}

func (s *ValidateShardTaskRequest) SetTargetTableName(v string) *ValidateShardTaskRequest {
	s.TargetTableName = &v
	return s
}

func (s *ValidateShardTaskRequest) SetTaskType(v string) *ValidateShardTaskRequest {
	s.TaskType = &v
	return s
}

type ValidateShardTaskResponseBody struct {
	// Indicates the check results.
	List []*ValidateShardTaskResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Indicates the ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ValidateShardTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateShardTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateShardTaskResponseBody) SetList(v []*ValidateShardTaskResponseBodyList) *ValidateShardTaskResponseBody {
	s.List = v
	return s
}

func (s *ValidateShardTaskResponseBody) SetRequestId(v string) *ValidateShardTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateShardTaskResponseBody) SetSuccess(v bool) *ValidateShardTaskResponseBody {
	s.Success = &v
	return s
}

type ValidateShardTaskResponseBodyList struct {
	// Indicates the name of a check item.
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// Indicates the result of the check item. Valid values:
	//
	// *   **0**: indicates the task is valid.
	// *   **1**: indicates the task is invalid.
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ValidateShardTaskResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ValidateShardTaskResponseBodyList) GoString() string {
	return s.String()
}

func (s *ValidateShardTaskResponseBodyList) SetItem(v string) *ValidateShardTaskResponseBodyList {
	s.Item = &v
	return s
}

func (s *ValidateShardTaskResponseBodyList) SetResult(v int32) *ValidateShardTaskResponseBodyList {
	s.Result = &v
	return s
}

type ValidateShardTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateShardTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateShardTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateShardTaskResponse) GoString() string {
	return s.String()
}

func (s *ValidateShardTaskResponse) SetHeaders(v map[string]*string) *ValidateShardTaskResponse {
	s.Headers = v
	return s
}

func (s *ValidateShardTaskResponse) SetStatusCode(v int32) *ValidateShardTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateShardTaskResponse) SetBody(v *ValidateShardTaskResponseBody) *ValidateShardTaskResponse {
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
		"ap-northeast-1":              tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":                  tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":              tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":              tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":              tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("drds.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("drds.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("drds.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("drds.aliyuncs.com"),
		"cn-chengdu":                  tea.String("drds.aliyuncs.com"),
		"cn-edge-1":                   tea.String("drds.aliyuncs.com"),
		"cn-fujian":                   tea.String("drds.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("drds.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("drds.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("drds.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("drds.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("drds.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("drds.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("drds.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("drds.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("drds.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("drds.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("drds.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("drds.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("drds.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("drds.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("drds.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("drds.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("drds.aliyuncs.com"),
		"cn-wuhan":                    tea.String("drds.aliyuncs.com"),
		"cn-yushanfang":               tea.String("drds.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("drds.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("drds.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("drds.aliyuncs.com"),
		"eu-central-1":                tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":                   tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"me-east-1":                   tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"us-west-1":                   tea.String("drds.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("drds"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ChangeAccountPasswordWithOptions(request *ChangeAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ChangeAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeAccountPassword"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeAccountPassword(request *ChangeAccountPasswordRequest) (_result *ChangeAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeAccountPasswordResponse{}
	_body, _err := client.ChangeAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeInstanceAzoneWithOptions(request *ChangeInstanceAzoneRequest, runtime *util.RuntimeOptions) (_result *ChangeInstanceAzoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeVSwitch)) {
		query["ChangeVSwitch"] = request.ChangeVSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsRegionId)) {
		query["DrdsRegionId"] = request.DrdsRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.NewVSwitch)) {
		query["NewVSwitch"] = request.NewVSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.OriginAzoneId)) {
		query["OriginAzoneId"] = request.OriginAzoneId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetAzoneId)) {
		query["TargetAzoneId"] = request.TargetAzoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeInstanceAzone"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeInstanceAzoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeInstanceAzone(request *ChangeInstanceAzoneRequest) (_result *ChangeInstanceAzoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeInstanceAzoneResponse{}
	_body, _err := client.ChangeInstanceAzoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckDrdsDbNameWithOptions(request *CheckDrdsDbNameRequest, runtime *util.RuntimeOptions) (_result *CheckDrdsDbNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDrdsDbName"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDrdsDbNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckDrdsDbName(request *CheckDrdsDbNameRequest) (_result *CheckDrdsDbNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDrdsDbNameResponse{}
	_body, _err := client.CheckDrdsDbNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckExpandStatusWithOptions(request *CheckExpandStatusRequest, runtime *util.RuntimeOptions) (_result *CheckExpandStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckExpandStatus"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckExpandStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckExpandStatus(request *CheckExpandStatusRequest) (_result *CheckExpandStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckExpandStatusResponse{}
	_body, _err := client.CheckExpandStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckSqlAuditEnableStatusWithOptions(request *CheckSqlAuditEnableStatusRequest, runtime *util.RuntimeOptions) (_result *CheckSqlAuditEnableStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckSqlAuditEnableStatus"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckSqlAuditEnableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckSqlAuditEnableStatus(request *CheckSqlAuditEnableStatusRequest) (_result *CheckSqlAuditEnableStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckSqlAuditEnableStatusResponse{}
	_body, _err := client.CheckSqlAuditEnableStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDrdsDBWithOptions(request *CreateDrdsDBRequest, runtime *util.RuntimeOptions) (_result *CreateDrdsDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstType)) {
		query["DbInstType"] = request.DbInstType
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstanceIsCreating)) {
		query["DbInstanceIsCreating"] = request.DbInstanceIsCreating
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Encode)) {
		query["Encode"] = request.Encode
	}

	if !tea.BoolValue(util.IsUnset(request.InstDbName)) {
		query["InstDbName"] = request.InstDbName
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RdsInstance)) {
		query["RdsInstance"] = request.RdsInstance
	}

	if !tea.BoolValue(util.IsUnset(request.RdsSuperAccount)) {
		query["RdsSuperAccount"] = request.RdsSuperAccount
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDrdsDB"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDrdsDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDrdsDB(request *CreateDrdsDBRequest) (_result *CreateDrdsDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDrdsDBResponse{}
	_body, _err := client.CreateDrdsDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDrdsInstanceWithOptions(request *CreateDrdsInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateDrdsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceSeries)) {
		query["InstanceSeries"] = request.InstanceSeries
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoRenew)) {
		query["IsAutoRenew"] = request.IsAutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.MasterInstId)) {
		query["MasterInstId"] = request.MasterInstId
	}

	if !tea.BoolValue(util.IsUnset(request.MySQLVersion)) {
		query["MySQLVersion"] = request.MySQLVersion
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.Quantity)) {
		query["Quantity"] = request.Quantity
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Specification)) {
		query["Specification"] = request.Specification
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VswitchId)) {
		query["VswitchId"] = request.VswitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.IsHa)) {
		query["isHa"] = request.IsHa
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDrdsInstance"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDrdsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDrdsInstance(request *CreateDrdsInstanceRequest) (_result *CreateDrdsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDrdsInstanceResponse{}
	_body, _err := client.CreateDrdsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceAccountWithOptions(request *CreateInstanceAccountRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DbPrivilege)) {
		query["DbPrivilege"] = request.DbPrivilege
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceAccount"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstanceAccount(request *CreateInstanceAccountRequest) (_result *CreateInstanceAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceAccountResponse{}
	_body, _err := client.CreateInstanceAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceInternetAddressWithOptions(request *CreateInstanceInternetAddressRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceInternetAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceInternetAddress"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceInternetAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstanceInternetAddress(request *CreateInstanceInternetAddressRequest) (_result *CreateInstanceInternetAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceInternetAddressResponse{}
	_body, _err := client.CreateInstanceInternetAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you call this operation, make sure that you understand the billing methods and pricing of PolarDB-X 1.0. For more information, visit the [pricing page](https://www.aliyun.com/price/product#/rds/detail).
 *
 * @param request CreateOrderForRdsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateOrderForRdsResponse
 */
func (client *Client) CreateOrderForRdsWithOptions(request *CreateOrderForRdsRequest, runtime *util.RuntimeOptions) (_result *CreateOrderForRdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrderForRds"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrderForRdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you call this operation, make sure that you understand the billing methods and pricing of PolarDB-X 1.0. For more information, visit the [pricing page](https://www.aliyun.com/price/product#/rds/detail).
 *
 * @param request CreateOrderForRdsRequest
 * @return CreateOrderForRdsResponse
 */
func (client *Client) CreateOrderForRds(request *CreateOrderForRdsRequest) (_result *CreateOrderForRdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrderForRdsResponse{}
	_body, _err := client.CreateOrderForRdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateShardTaskWithOptions(request *CreateShardTaskRequest, runtime *util.RuntimeOptions) (_result *CreateShardTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTableName)) {
		query["SourceTableName"] = request.SourceTableName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetTableName)) {
		query["TargetTableName"] = request.TargetTableName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateShardTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateShardTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateShardTask(request *CreateShardTaskRequest) (_result *CreateShardTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateShardTaskResponse{}
	_body, _err := client.CreateShardTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackMenuWithOptions(request *DescribeBackMenuRequest, runtime *util.RuntimeOptions) (_result *DescribeBackMenuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackMenu"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackMenuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackMenu(request *DescribeBackMenuRequest) (_result *DescribeBackMenuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackMenuResponse{}
	_body, _err := client.DescribeBackMenuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupDbsWithOptions(request *DescribeBackupDbsRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupDbsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredRestoreTime)) {
		query["PreferredRestoreTime"] = request.PreferredRestoreTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupDbs"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupDbsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupDbs(request *DescribeBackupDbsRequest) (_result *DescribeBackupDbsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupDbsResponse{}
	_body, _err := client.DescribeBackupDbsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupLocalWithOptions(request *DescribeBackupLocalRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupLocalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupLocal"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupLocalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupLocal(request *DescribeBackupLocalRequest) (_result *DescribeBackupLocalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupLocalResponse{}
	_body, _err := client.DescribeBackupLocalWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupPolicy"),
		Version:     tea.String("2019-01-23"),
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

func (client *Client) DescribeBackupSetsWithOptions(request *DescribeBackupSetsRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupSetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
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
		Action:      tea.String("DescribeBackupSets"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupSetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupSets(request *DescribeBackupSetsRequest) (_result *DescribeBackupSetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupSetsResponse{}
	_body, _err := client.DescribeBackupSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupTimesWithOptions(request *DescribeBackupTimesRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupTimesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupTimes"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupTimesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupTimes(request *DescribeBackupTimesRequest) (_result *DescribeBackupTimesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupTimesResponse{}
	_body, _err := client.DescribeBackupTimesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBroadcastTablesWithOptions(request *DescribeBroadcastTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeBroadcastTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBroadcastTables"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBroadcastTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBroadcastTables(request *DescribeBroadcastTablesRequest) (_result *DescribeBroadcastTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBroadcastTablesResponse{}
	_body, _err := client.DescribeBroadcastTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDbInstanceDbsWithOptions(request *DescribeDbInstanceDbsRequest, runtime *util.RuntimeOptions) (_result *DescribeDbInstanceDbsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstType)) {
		query["DbInstType"] = request.DbInstType
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDbInstanceDbs"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDbInstanceDbsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDbInstanceDbs(request *DescribeDbInstanceDbsRequest) (_result *DescribeDbInstanceDbsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDbInstanceDbsResponse{}
	_body, _err := client.DescribeDbInstanceDbsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDbInstancesWithOptions(request *DescribeDbInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDbInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstType)) {
		query["DbInstType"] = request.DbInstType
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
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

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDbInstances"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDbInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDbInstances(request *DescribeDbInstancesRequest) (_result *DescribeDbInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDbInstancesResponse{}
	_body, _err := client.DescribeDbInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDBWithOptions(request *DescribeDrdsDBRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDB"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDB(request *DescribeDrdsDBRequest) (_result *DescribeDrdsDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDBResponse{}
	_body, _err := client.DescribeDrdsDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDBClusterWithOptions(request *DescribeDrdsDBClusterRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDBCluster"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDBCluster(request *DescribeDrdsDBClusterRequest) (_result *DescribeDrdsDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDBClusterResponse{}
	_body, _err := client.DescribeDrdsDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDBIpWhiteListWithOptions(request *DescribeDrdsDBIpWhiteListRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDBIpWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDBIpWhiteList"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsDBIpWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDBIpWhiteList(request *DescribeDrdsDBIpWhiteListRequest) (_result *DescribeDrdsDBIpWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDBIpWhiteListResponse{}
	_body, _err := client.DescribeDrdsDBIpWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDBsWithOptions(request *DescribeDrdsDBsRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDBsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDBs"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsDBsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDBs(request *DescribeDrdsDBsRequest) (_result *DescribeDrdsDBsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDBsResponse{}
	_body, _err := client.DescribeDrdsDBsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDbInstanceWithOptions(request *DescribeDrdsDbInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDbInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDbInstance"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsDbInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDbInstance(request *DescribeDrdsDbInstanceRequest) (_result *DescribeDrdsDbInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDbInstanceResponse{}
	_body, _err := client.DescribeDrdsDbInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDbInstancesWithOptions(request *DescribeDrdsDbInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDbInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
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
		Action:      tea.String("DescribeDrdsDbInstances"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsDbInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDbInstances(request *DescribeDrdsDbInstancesRequest) (_result *DescribeDrdsDbInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDbInstancesResponse{}
	_body, _err := client.DescribeDrdsDbInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDbRdsNameListWithOptions(request *DescribeDrdsDbRdsNameListRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDbRdsNameListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDbRdsNameList"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsDbRdsNameListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDbRdsNameList(request *DescribeDrdsDbRdsNameListRequest) (_result *DescribeDrdsDbRdsNameListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDbRdsNameListResponse{}
	_body, _err := client.DescribeDrdsDbRdsNameListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceWithOptions(request *DescribeDrdsInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstance"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsInstance(request *DescribeDrdsInstanceRequest) (_result *DescribeDrdsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsInstanceResponse{}
	_body, _err := client.DescribeDrdsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceDbMonitorWithOptions(request *DescribeDrdsInstanceDbMonitorRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceDbMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
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
		Action:      tea.String("DescribeDrdsInstanceDbMonitor"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsInstanceDbMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceDbMonitor(request *DescribeDrdsInstanceDbMonitorRequest) (_result *DescribeDrdsInstanceDbMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsInstanceDbMonitorResponse{}
	_body, _err := client.DescribeDrdsInstanceDbMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceLevelTasksWithOptions(request *DescribeDrdsInstanceLevelTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceLevelTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstanceLevelTasks"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsInstanceLevelTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceLevelTasks(request *DescribeDrdsInstanceLevelTasksRequest) (_result *DescribeDrdsInstanceLevelTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsInstanceLevelTasksResponse{}
	_body, _err := client.DescribeDrdsInstanceLevelTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceMonitorWithOptions(request *DescribeDrdsInstanceMonitorRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodMultiple)) {
		query["PeriodMultiple"] = request.PeriodMultiple
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
		Action:      tea.String("DescribeDrdsInstanceMonitor"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsInstanceMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceMonitor(request *DescribeDrdsInstanceMonitorRequest) (_result *DescribeDrdsInstanceMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsInstanceMonitorResponse{}
	_body, _err := client.DescribeDrdsInstanceMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceVersionWithOptions(request *DescribeDrdsInstanceVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstanceVersion"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsInstanceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceVersion(request *DescribeDrdsInstanceVersionRequest) (_result *DescribeDrdsInstanceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsInstanceVersionResponse{}
	_body, _err := client.DescribeDrdsInstanceVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsInstancesWithOptions(request *DescribeDrdsInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Expired)) {
		query["Expired"] = request.Expired
	}

	if !tea.BoolValue(util.IsUnset(request.Mix)) {
		query["Mix"] = request.Mix
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersion)) {
		query["ProductVersion"] = request.ProductVersion
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

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstances"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsInstances(request *DescribeDrdsInstancesRequest) (_result *DescribeDrdsInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsInstancesResponse{}
	_body, _err := client.DescribeDrdsInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsParamsWithOptions(request *DescribeDrdsParamsRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsParamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ParamLevel)) {
		query["ParamLevel"] = request.ParamLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsParams"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsParamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsParams(request *DescribeDrdsParamsRequest) (_result *DescribeDrdsParamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsParamsResponse{}
	_body, _err := client.DescribeDrdsParamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsRdsInstancesWithOptions(request *DescribeDrdsRdsInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsRdsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsRdsInstances"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsRdsInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsRdsInstances(request *DescribeDrdsRdsInstancesRequest) (_result *DescribeDrdsRdsInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsRdsInstancesResponse{}
	_body, _err := client.DescribeDrdsRdsInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsShardingDbsWithOptions(request *DescribeDrdsShardingDbsRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsShardingDbsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DbNamePattern)) {
		query["DbNamePattern"] = request.DbNamePattern
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
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
		Action:      tea.String("DescribeDrdsShardingDbs"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsShardingDbsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsShardingDbs(request *DescribeDrdsShardingDbsRequest) (_result *DescribeDrdsShardingDbsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsShardingDbsResponse{}
	_body, _err := client.DescribeDrdsShardingDbsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsSlowSqlsWithOptions(request *DescribeDrdsSlowSqlsRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsSlowSqlsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExeTime)) {
		query["ExeTime"] = request.ExeTime
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
		Action:      tea.String("DescribeDrdsSlowSqls"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsSlowSqlsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsSlowSqls(request *DescribeDrdsSlowSqlsRequest) (_result *DescribeDrdsSlowSqlsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsSlowSqlsResponse{}
	_body, _err := client.DescribeDrdsSlowSqlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsSqlAuditStatusWithOptions(request *DescribeDrdsSqlAuditStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsSqlAuditStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsSqlAuditStatus"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsSqlAuditStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsSqlAuditStatus(request *DescribeDrdsSqlAuditStatusRequest) (_result *DescribeDrdsSqlAuditStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsSqlAuditStatusResponse{}
	_body, _err := client.DescribeDrdsSqlAuditStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsTasksWithOptions(request *DescribeDrdsTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsTasks"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsTasks(request *DescribeDrdsTasksRequest) (_result *DescribeDrdsTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsTasksResponse{}
	_body, _err := client.DescribeDrdsTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExpandLogicTableInfoListWithOptions(request *DescribeExpandLogicTableInfoListRequest, runtime *util.RuntimeOptions) (_result *DescribeExpandLogicTableInfoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpandLogicTableInfoList"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExpandLogicTableInfoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExpandLogicTableInfoList(request *DescribeExpandLogicTableInfoListRequest) (_result *DescribeExpandLogicTableInfoListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpandLogicTableInfoListResponse{}
	_body, _err := client.DescribeExpandLogicTableInfoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHotDbListWithOptions(request *DescribeHotDbListRequest, runtime *util.RuntimeOptions) (_result *DescribeHotDbListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHotDbList"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHotDbListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHotDbList(request *DescribeHotDbListRequest) (_result *DescribeHotDbListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHotDbListResponse{}
	_body, _err := client.DescribeHotDbListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstDbLogInfoWithOptions(request *DescribeInstDbLogInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeInstDbLogInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstDbLogInfo"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstDbLogInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstDbLogInfo(request *DescribeInstDbLogInfoRequest) (_result *DescribeInstDbLogInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstDbLogInfoResponse{}
	_body, _err := client.DescribeInstDbLogInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstDbSlsInfoWithOptions(request *DescribeInstDbSlsInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeInstDbSlsInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstDbSlsInfo"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstDbSlsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstDbSlsInfo(request *DescribeInstDbSlsInfoRequest) (_result *DescribeInstDbSlsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstDbSlsInfoResponse{}
	_body, _err := client.DescribeInstDbSlsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceAccountsWithOptions(request *DescribeInstanceAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceAccounts"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceAccounts(request *DescribeInstanceAccountsRequest) (_result *DescribeInstanceAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceAccountsResponse{}
	_body, _err := client.DescribeInstanceAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceSwitchAzoneWithOptions(request *DescribeInstanceSwitchAzoneRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSwitchAzoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceSwitchAzone"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceSwitchAzoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceSwitchAzone(request *DescribeInstanceSwitchAzoneRequest) (_result *DescribeInstanceSwitchAzoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSwitchAzoneResponse{}
	_body, _err := client.DescribeInstanceSwitchAzoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ****
 *
 * @param request DescribeInstanceSwitchNetworkRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeInstanceSwitchNetworkResponse
 */
func (client *Client) DescribeInstanceSwitchNetworkWithOptions(request *DescribeInstanceSwitchNetworkRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSwitchNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceSwitchNetwork"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceSwitchNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ****
 *
 * @param request DescribeInstanceSwitchNetworkRequest
 * @return DescribeInstanceSwitchNetworkResponse
 */
func (client *Client) DescribeInstanceSwitchNetwork(request *DescribeInstanceSwitchNetworkRequest) (_result *DescribeInstanceSwitchNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSwitchNetworkResponse{}
	_body, _err := client.DescribeInstanceSwitchNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePreCheckResultWithOptions(request *DescribePreCheckResultRequest, runtime *util.RuntimeOptions) (_result *DescribePreCheckResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePreCheckResult"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePreCheckResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePreCheckResult(request *DescribePreCheckResultRequest) (_result *DescribePreCheckResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePreCheckResultResponse{}
	_body, _err := client.DescribePreCheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRDSPerformanceWithOptions(request *DescribeRDSPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeRDSPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstType)) {
		query["DbInstType"] = request.DbInstType
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keys)) {
		query["Keys"] = request.Keys
	}

	if !tea.BoolValue(util.IsUnset(request.RdsInstanceId)) {
		query["RdsInstanceId"] = request.RdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRDSPerformance"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRDSPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRDSPerformance(request *DescribeRDSPerformanceRequest) (_result *DescribeRDSPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRDSPerformanceResponse{}
	_body, _err := client.DescribeRDSPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRdsCommodityWithOptions(request *DescribeRdsCommodityRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsCommodityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRdsCommodity"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdsCommodityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRdsCommodity(request *DescribeRdsCommodityRequest) (_result *DescribeRdsCommodityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsCommodityResponse{}
	_body, _err := client.DescribeRdsCommodityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRdsPerformanceSummaryWithOptions(request *DescribeRdsPerformanceSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsPerformanceSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RdsInstanceId)) {
		query["RdsInstanceId"] = request.RdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRdsPerformanceSummary"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdsPerformanceSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRdsPerformanceSummary(request *DescribeRdsPerformanceSummaryRequest) (_result *DescribeRdsPerformanceSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsPerformanceSummaryResponse{}
	_body, _err := client.DescribeRdsPerformanceSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRdsSuperAccountInstancesWithOptions(request *DescribeRdsSuperAccountInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsSuperAccountInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstType)) {
		query["DbInstType"] = request.DbInstType
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RdsInstance)) {
		query["RdsInstance"] = request.RdsInstance
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRdsSuperAccountInstances"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdsSuperAccountInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRdsSuperAccountInstances(request *DescribeRdsSuperAccountInstancesRequest) (_result *DescribeRdsSuperAccountInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsSuperAccountInstancesResponse{}
	_body, _err := client.DescribeRdsSuperAccountInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecycleBinStatusWithOptions(request *DescribeRecycleBinStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeRecycleBinStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRecycleBinStatus"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRecycleBinStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecycleBinStatus(request *DescribeRecycleBinStatusRequest) (_result *DescribeRecycleBinStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecycleBinStatusResponse{}
	_body, _err := client.DescribeRecycleBinStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecycleBinTablesWithOptions(request *DescribeRecycleBinTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeRecycleBinTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRecycleBinTables"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRecycleBinTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecycleBinTables(request *DescribeRecycleBinTablesRequest) (_result *DescribeRecycleBinTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecycleBinTablesResponse{}
	_body, _err := client.DescribeRecycleBinTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRestoreOrderWithOptions(request *DescribeRestoreOrderRequest, runtime *util.RuntimeOptions) (_result *DescribeRestoreOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupDbNames)) {
		query["BackupDbNames"] = request.BackupDbNames
	}

	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupLevel)) {
		query["BackupLevel"] = request.BackupLevel
	}

	if !tea.BoolValue(util.IsUnset(request.BackupMode)) {
		query["BackupMode"] = request.BackupMode
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupTime)) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRestoreOrder"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRestoreOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRestoreOrder(request *DescribeRestoreOrderRequest) (_result *DescribeRestoreOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreOrderResponse{}
	_body, _err := client.DescribeRestoreOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeShardTaskInfoWithOptions(request *DescribeShardTaskInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeShardTaskInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTableName)) {
		query["SourceTableName"] = request.SourceTableName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetTableName)) {
		query["TargetTableName"] = request.TargetTableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeShardTaskInfo"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeShardTaskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeShardTaskInfo(request *DescribeShardTaskInfoRequest) (_result *DescribeShardTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeShardTaskInfoResponse{}
	_body, _err := client.DescribeShardTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSqlFlashbakTaskWithOptions(request *DescribeSqlFlashbakTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeSqlFlashbakTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSqlFlashbakTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSqlFlashbakTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSqlFlashbakTask(request *DescribeSqlFlashbakTaskRequest) (_result *DescribeSqlFlashbakTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSqlFlashbakTaskResponse{}
	_body, _err := client.DescribeSqlFlashbakTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTableWithOptions(request *DescribeTableRequest, runtime *util.RuntimeOptions) (_result *DescribeTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTable"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTable(request *DescribeTableRequest) (_result *DescribeTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTableResponse{}
	_body, _err := client.DescribeTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTableListByTypeWithOptions(request *DescribeTableListByTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeTableListByTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TableType)) {
		query["TableType"] = request.TableType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTableListByType"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTableListByTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTableListByType(request *DescribeTableListByTypeRequest) (_result *DescribeTableListByTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTableListByTypeResponse{}
	_body, _err := client.DescribeTableListByTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTablesWithOptions(request *DescribeTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTables"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTables(request *DescribeTablesRequest) (_result *DescribeTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTablesResponse{}
	_body, _err := client.DescribeTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableSqlAuditWithOptions(request *DisableSqlAuditRequest, runtime *util.RuntimeOptions) (_result *DisableSqlAuditResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableSqlAudit"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableSqlAuditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableSqlAudit(request *DisableSqlAuditRequest) (_result *DisableSqlAuditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableSqlAuditResponse{}
	_body, _err := client.DisableSqlAuditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableInstanceIpv6AddressWithOptions(request *EnableInstanceIpv6AddressRequest, runtime *util.RuntimeOptions) (_result *EnableInstanceIpv6AddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableInstanceIpv6Address"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableInstanceIpv6AddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableInstanceIpv6Address(request *EnableInstanceIpv6AddressRequest) (_result *EnableInstanceIpv6AddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableInstanceIpv6AddressResponse{}
	_body, _err := client.EnableInstanceIpv6AddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableSqlAuditWithOptions(request *EnableSqlAuditRequest, runtime *util.RuntimeOptions) (_result *EnableSqlAuditResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecall)) {
		query["IsRecall"] = request.IsRecall
	}

	if !tea.BoolValue(util.IsUnset(request.RecallEndTimestamp)) {
		query["RecallEndTimestamp"] = request.RecallEndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.RecallStartTimestamp)) {
		query["RecallStartTimestamp"] = request.RecallStartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableSqlAudit"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableSqlAuditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSqlAudit(request *EnableSqlAuditRequest) (_result *EnableSqlAuditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableSqlAuditResponse{}
	_body, _err := client.EnableSqlAuditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableSqlFlashbackMatchSwitchWithOptions(request *EnableSqlFlashbackMatchSwitchRequest, runtime *util.RuntimeOptions) (_result *EnableSqlFlashbackMatchSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableSqlFlashbackMatchSwitch"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableSqlFlashbackMatchSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSqlFlashbackMatchSwitch(request *EnableSqlFlashbackMatchSwitchRequest) (_result *EnableSqlFlashbackMatchSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableSqlFlashbackMatchSwitchResponse{}
	_body, _err := client.EnableSqlFlashbackMatchSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FlashbackRecycleBinTableWithOptions(request *FlashbackRecycleBinTableRequest, runtime *util.RuntimeOptions) (_result *FlashbackRecycleBinTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FlashbackRecycleBinTable"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FlashbackRecycleBinTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FlashbackRecycleBinTable(request *FlashbackRecycleBinTableRequest) (_result *FlashbackRecycleBinTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FlashbackRecycleBinTableResponse{}
	_body, _err := client.FlashbackRecycleBinTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDrdsDbRdsRelationInfoWithOptions(request *GetDrdsDbRdsRelationInfoRequest, runtime *util.RuntimeOptions) (_result *GetDrdsDbRdsRelationInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDrdsDbRdsRelationInfo"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDrdsDbRdsRelationInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDrdsDbRdsRelationInfo(request *GetDrdsDbRdsRelationInfoRequest) (_result *GetDrdsDbRdsRelationInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDrdsDbRdsRelationInfoResponse{}
	_body, _err := client.GetDrdsDbRdsRelationInfoWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
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
		Version:     tea.String("2019-01-23"),
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

func (client *Client) ManagePrivateRdsWithOptions(request *ManagePrivateRdsRequest, runtime *util.RuntimeOptions) (_result *ManagePrivateRdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.RdsAction)) {
		query["RdsAction"] = request.RdsAction
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ManagePrivateRds"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ManagePrivateRdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ManagePrivateRds(request *ManagePrivateRdsRequest) (_result *ManagePrivateRdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ManagePrivateRdsResponse{}
	_body, _err := client.ManagePrivateRdsWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountDescription"),
		Version:     tea.String("2019-01-23"),
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

func (client *Client) ModifyAccountPrivilegeWithOptions(request *ModifyAccountPrivilegeRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountPrivilegeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DbPrivilege)) {
		query["DbPrivilege"] = request.DbPrivilege
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountPrivilege"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccountPrivilegeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountPrivilege(request *ModifyAccountPrivilegeRequest) (_result *ModifyAccountPrivilegeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountPrivilegeResponse{}
	_body, _err := client.ModifyAccountPrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDrdsInstanceDescriptionWithOptions(request *ModifyDrdsInstanceDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDrdsInstanceDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDrdsInstanceDescription"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDrdsInstanceDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDrdsInstanceDescription(request *ModifyDrdsInstanceDescriptionRequest) (_result *ModifyDrdsInstanceDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDrdsInstanceDescriptionResponse{}
	_body, _err := client.ModifyDrdsInstanceDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDrdsIpWhiteListWithOptions(request *ModifyDrdsIpWhiteListRequest, runtime *util.RuntimeOptions) (_result *ModifyDrdsIpWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupAttribute)) {
		query["GroupAttribute"] = request.GroupAttribute
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.IpWhiteList)) {
		query["IpWhiteList"] = request.IpWhiteList
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDrdsIpWhiteList"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDrdsIpWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDrdsIpWhiteList(request *ModifyDrdsIpWhiteListRequest) (_result *ModifyDrdsIpWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDrdsIpWhiteListResponse{}
	_body, _err := client.ModifyDrdsIpWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPolarDbReadWeightWithOptions(request *ModifyPolarDbReadWeightRequest, runtime *util.RuntimeOptions) (_result *ModifyPolarDbReadWeightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DbNodeIds)) {
		query["DbNodeIds"] = request.DbNodeIds
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Weights)) {
		query["Weights"] = request.Weights
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPolarDbReadWeight"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPolarDbReadWeightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPolarDbReadWeight(request *ModifyPolarDbReadWeightRequest) (_result *ModifyPolarDbReadWeightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPolarDbReadWeightResponse{}
	_body, _err := client.ModifyPolarDbReadWeightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRdsReadWeightWithOptions(request *ModifyRdsReadWeightRequest, runtime *util.RuntimeOptions) (_result *ModifyRdsReadWeightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceNames)) {
		query["InstanceNames"] = request.InstanceNames
	}

	if !tea.BoolValue(util.IsUnset(request.Weights)) {
		query["Weights"] = request.Weights
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyRdsReadWeight"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyRdsReadWeightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRdsReadWeight(request *ModifyRdsReadWeightRequest) (_result *ModifyRdsReadWeightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRdsReadWeightResponse{}
	_body, _err := client.ModifyRdsReadWeightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutStartBackupWithOptions(request *PutStartBackupRequest, runtime *util.RuntimeOptions) (_result *PutStartBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupDbNames)) {
		query["BackupDbNames"] = request.BackupDbNames
	}

	if !tea.BoolValue(util.IsUnset(request.BackupLevel)) {
		query["BackupLevel"] = request.BackupLevel
	}

	if !tea.BoolValue(util.IsUnset(request.BackupMode)) {
		query["BackupMode"] = request.BackupMode
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutStartBackup"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutStartBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutStartBackup(request *PutStartBackupRequest) (_result *PutStartBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutStartBackupResponse{}
	_body, _err := client.PutStartBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshDrdsAtomUrlWithOptions(request *RefreshDrdsAtomUrlRequest, runtime *util.RuntimeOptions) (_result *RefreshDrdsAtomUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshDrdsAtomUrl"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshDrdsAtomUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshDrdsAtomUrl(request *RefreshDrdsAtomUrlRequest) (_result *RefreshDrdsAtomUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshDrdsAtomUrlResponse{}
	_body, _err := client.RefreshDrdsAtomUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseInstanceInternetAddressWithOptions(request *ReleaseInstanceInternetAddressRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstanceInternetAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseInstanceInternetAddress"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseInstanceInternetAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseInstanceInternetAddress(request *ReleaseInstanceInternetAddressRequest) (_result *ReleaseInstanceInternetAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstanceInternetAddressResponse{}
	_body, _err := client.ReleaseInstanceInternetAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveBackupsSetWithOptions(request *RemoveBackupsSetRequest, runtime *util.RuntimeOptions) (_result *RemoveBackupsSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveBackupsSet"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveBackupsSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveBackupsSet(request *RemoveBackupsSetRequest) (_result *RemoveBackupsSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveBackupsSetResponse{}
	_body, _err := client.RemoveBackupsSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveDrdsDbWithOptions(request *RemoveDrdsDbRequest, runtime *util.RuntimeOptions) (_result *RemoveDrdsDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveDrdsDb"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveDrdsDbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveDrdsDb(request *RemoveDrdsDbRequest) (_result *RemoveDrdsDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveDrdsDbResponse{}
	_body, _err := client.RemoveDrdsDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveDrdsDbFailedRecordWithOptions(request *RemoveDrdsDbFailedRecordRequest, runtime *util.RuntimeOptions) (_result *RemoveDrdsDbFailedRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveDrdsDbFailedRecord"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveDrdsDbFailedRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveDrdsDbFailedRecord(request *RemoveDrdsDbFailedRecordRequest) (_result *RemoveDrdsDbFailedRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveDrdsDbFailedRecordResponse{}
	_body, _err := client.RemoveDrdsDbFailedRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * > *   You can call this operation to release an instance that is charged based on only the pay-as-you-go billing method.
 * >*   If the specifications of the instance are being changed, or one or more databases exist in the instance, you cannot call this operation to release the instance.
 *
 * @param request RemoveDrdsInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveDrdsInstanceResponse
 */
func (client *Client) RemoveDrdsInstanceWithOptions(request *RemoveDrdsInstanceRequest, runtime *util.RuntimeOptions) (_result *RemoveDrdsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveDrdsInstance"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveDrdsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * > *   You can call this operation to release an instance that is charged based on only the pay-as-you-go billing method.
 * >*   If the specifications of the instance are being changed, or one or more databases exist in the instance, you cannot call this operation to release the instance.
 *
 * @param request RemoveDrdsInstanceRequest
 * @return RemoveDrdsInstanceResponse
 */
func (client *Client) RemoveDrdsInstance(request *RemoveDrdsInstanceRequest) (_result *RemoveDrdsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveDrdsInstanceResponse{}
	_body, _err := client.RemoveDrdsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveInstanceAccountWithOptions(request *RemoveInstanceAccountRequest, runtime *util.RuntimeOptions) (_result *RemoveInstanceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveInstanceAccount"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveInstanceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveInstanceAccount(request *RemoveInstanceAccountRequest) (_result *RemoveInstanceAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveInstanceAccountResponse{}
	_body, _err := client.RemoveInstanceAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveRecycleBinTableWithOptions(request *RemoveRecycleBinTableRequest, runtime *util.RuntimeOptions) (_result *RemoveRecycleBinTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveRecycleBinTable"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveRecycleBinTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveRecycleBinTable(request *RemoveRecycleBinTableRequest) (_result *RemoveRecycleBinTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveRecycleBinTableResponse{}
	_body, _err := client.RemoveRecycleBinTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartDrdsInstanceWithOptions(request *RestartDrdsInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartDrdsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartDrdsInstance"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartDrdsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartDrdsInstance(request *RestartDrdsInstanceRequest) (_result *RestartDrdsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartDrdsInstanceResponse{}
	_body, _err := client.RestartDrdsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackInstanceVersionWithOptions(request *RollbackInstanceVersionRequest, runtime *util.RuntimeOptions) (_result *RollbackInstanceVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RollbackInstanceVersion"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RollbackInstanceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackInstanceVersion(request *RollbackInstanceVersionRequest) (_result *RollbackInstanceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackInstanceVersionResponse{}
	_body, _err := client.RollbackInstanceVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetBackupLocalWithOptions(request *SetBackupLocalRequest, runtime *util.RuntimeOptions) (_result *SetBackupLocalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.HighSpaceUsageProtection)) {
		query["HighSpaceUsageProtection"] = request.HighSpaceUsageProtection
	}

	if !tea.BoolValue(util.IsUnset(request.LocalLogRetentionHours)) {
		query["LocalLogRetentionHours"] = request.LocalLogRetentionHours
	}

	if !tea.BoolValue(util.IsUnset(request.LocalLogRetentionSpace)) {
		query["LocalLogRetentionSpace"] = request.LocalLogRetentionSpace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetBackupLocal"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetBackupLocalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetBackupLocal(request *SetBackupLocalRequest) (_result *SetBackupLocalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetBackupLocalResponse{}
	_body, _err := client.SetBackupLocalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetBackupPolicyWithOptions(request *SetBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *SetBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupDbNames)) {
		query["BackupDbNames"] = request.BackupDbNames
	}

	if !tea.BoolValue(util.IsUnset(request.BackupLevel)) {
		query["BackupLevel"] = request.BackupLevel
	}

	if !tea.BoolValue(util.IsUnset(request.BackupLog)) {
		query["BackupLog"] = request.BackupLog
	}

	if !tea.BoolValue(util.IsUnset(request.BackupMode)) {
		query["BackupMode"] = request.BackupMode
	}

	if !tea.BoolValue(util.IsUnset(request.DataBackupRetentionPeriod)) {
		query["DataBackupRetentionPeriod"] = request.DataBackupRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LogBackupRetentionPeriod)) {
		query["LogBackupRetentionPeriod"] = request.LogBackupRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupEndTime)) {
		query["PreferredBackupEndTime"] = request.PreferredBackupEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupPeriod)) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupStartTime)) {
		query["PreferredBackupStartTime"] = request.PreferredBackupStartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetBackupPolicy"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetBackupPolicy(request *SetBackupPolicyRequest) (_result *SetBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetBackupPolicyResponse{}
	_body, _err := client.SetBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetupBroadcastTablesWithOptions(request *SetupBroadcastTablesRequest, runtime *util.RuntimeOptions) (_result *SetupBroadcastTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		query["Active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetupBroadcastTables"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetupBroadcastTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetupBroadcastTables(request *SetupBroadcastTablesRequest) (_result *SetupBroadcastTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetupBroadcastTablesResponse{}
	_body, _err := client.SetupBroadcastTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetupDrdsParamsWithOptions(request *SetupDrdsParamsRequest, runtime *util.RuntimeOptions) (_result *SetupDrdsParamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ParamLevel)) {
		query["ParamLevel"] = request.ParamLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetupDrdsParams"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetupDrdsParamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetupDrdsParams(request *SetupDrdsParamsRequest) (_result *SetupDrdsParamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetupDrdsParamsResponse{}
	_body, _err := client.SetupDrdsParamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetupRecycleBinStatusWithOptions(request *SetupRecycleBinStatusRequest, runtime *util.RuntimeOptions) (_result *SetupRecycleBinStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StatusAction)) {
		query["StatusAction"] = request.StatusAction
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetupRecycleBinStatus"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetupRecycleBinStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetupRecycleBinStatus(request *SetupRecycleBinStatusRequest) (_result *SetupRecycleBinStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetupRecycleBinStatusResponse{}
	_body, _err := client.SetupRecycleBinStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetupTableWithOptions(request *SetupTableRequest, runtime *util.RuntimeOptions) (_result *SetupTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowFullTableScan)) {
		query["AllowFullTableScan"] = request.AllowFullTableScan
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetupTable"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetupTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetupTable(request *SetupTableRequest) (_result *SetupTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetupTableResponse{}
	_body, _err := client.SetupTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartRestoreWithOptions(request *StartRestoreRequest, runtime *util.RuntimeOptions) (_result *StartRestoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupDbNames)) {
		query["BackupDbNames"] = request.BackupDbNames
	}

	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupLevel)) {
		query["BackupLevel"] = request.BackupLevel
	}

	if !tea.BoolValue(util.IsUnset(request.BackupMode)) {
		query["BackupMode"] = request.BackupMode
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupTime)) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartRestore"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartRestoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartRestore(request *StartRestoreRequest) (_result *StartRestoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartRestoreResponse{}
	_body, _err := client.StartRestoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitCleanTaskWithOptions(request *SubmitCleanTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitCleanTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpandType)) {
		query["ExpandType"] = request.ExpandType
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentJobId)) {
		query["ParentJobId"] = request.ParentJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitCleanTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitCleanTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitCleanTask(request *SubmitCleanTaskRequest) (_result *SubmitCleanTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitCleanTaskResponse{}
	_body, _err := client.SubmitCleanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitHotExpandPreCheckTaskWithOptions(request *SubmitHotExpandPreCheckTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitHotExpandPreCheckTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstType)) {
		query["DbInstType"] = request.DbInstType
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TableList)) {
		query["TableList"] = request.TableList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitHotExpandPreCheckTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitHotExpandPreCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitHotExpandPreCheckTask(request *SubmitHotExpandPreCheckTaskRequest) (_result *SubmitHotExpandPreCheckTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitHotExpandPreCheckTaskResponse{}
	_body, _err := client.SubmitHotExpandPreCheckTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitHotExpandTaskWithOptions(request *SubmitHotExpandTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitHotExpandTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendedMapping)) {
		query["ExtendedMapping"] = request.ExtendedMapping
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceDbMapping)) {
		query["InstanceDbMapping"] = request.InstanceDbMapping
	}

	if !tea.BoolValue(util.IsUnset(request.Mapping)) {
		query["Mapping"] = request.Mapping
	}

	if !tea.BoolValue(util.IsUnset(request.SupperAccountMapping)) {
		query["SupperAccountMapping"] = request.SupperAccountMapping
	}

	if !tea.BoolValue(util.IsUnset(request.TaskDesc)) {
		query["TaskDesc"] = request.TaskDesc
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitHotExpandTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitHotExpandTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitHotExpandTask(request *SubmitHotExpandTaskRequest) (_result *SubmitHotExpandTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitHotExpandTaskResponse{}
	_body, _err := client.SubmitHotExpandTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSmoothExpandPreCheckWithOptions(request *SubmitSmoothExpandPreCheckRequest, runtime *util.RuntimeOptions) (_result *SubmitSmoothExpandPreCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstType)) {
		query["DbInstType"] = request.DbInstType
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitSmoothExpandPreCheck"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSmoothExpandPreCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSmoothExpandPreCheck(request *SubmitSmoothExpandPreCheckRequest) (_result *SubmitSmoothExpandPreCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSmoothExpandPreCheckResponse{}
	_body, _err := client.SubmitSmoothExpandPreCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSmoothExpandPreCheckTaskWithOptions(request *SubmitSmoothExpandPreCheckTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitSmoothExpandPreCheckTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitSmoothExpandPreCheckTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSmoothExpandPreCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSmoothExpandPreCheckTask(request *SubmitSmoothExpandPreCheckTaskRequest) (_result *SubmitSmoothExpandPreCheckTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSmoothExpandPreCheckTaskResponse{}
	_body, _err := client.SubmitSmoothExpandPreCheckTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSqlFlashbackTaskWithOptions(request *SubmitSqlFlashbackTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitSqlFlashbackTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RecallRestoreType)) {
		query["RecallRestoreType"] = request.RecallRestoreType
	}

	if !tea.BoolValue(util.IsUnset(request.RecallType)) {
		query["RecallType"] = request.RecallType
	}

	if !tea.BoolValue(util.IsUnset(request.SqlPk)) {
		query["SqlPk"] = request.SqlPk
	}

	if !tea.BoolValue(util.IsUnset(request.SqlType)) {
		query["SqlType"] = request.SqlType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	if !tea.BoolValue(util.IsUnset(request.TraceId)) {
		query["TraceId"] = request.TraceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitSqlFlashbackTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSqlFlashbackTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSqlFlashbackTask(request *SubmitSqlFlashbackTaskRequest) (_result *SubmitSqlFlashbackTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSqlFlashbackTaskResponse{}
	_body, _err := client.SubmitSqlFlashbackTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchGlobalBroadcastTypeWithOptions(request *SwitchGlobalBroadcastTypeRequest, runtime *util.RuntimeOptions) (_result *SwitchGlobalBroadcastTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchGlobalBroadcastType"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwitchGlobalBroadcastTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchGlobalBroadcastType(request *SwitchGlobalBroadcastTypeRequest) (_result *SwitchGlobalBroadcastTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchGlobalBroadcastTypeResponse{}
	_body, _err := client.SwitchGlobalBroadcastTypeWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
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
		Version:     tea.String("2019-01-23"),
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

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
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
		Version:     tea.String("2019-01-23"),
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

func (client *Client) UpdateInstanceNetworkWithOptions(request *UpdateInstanceNetworkRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassicExpiredDays)) {
		query["ClassicExpiredDays"] = request.ClassicExpiredDays
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RetainClassic)) {
		query["RetainClassic"] = request.RetainClassic
	}

	if !tea.BoolValue(util.IsUnset(request.SrcInstanceNetworkType)) {
		query["SrcInstanceNetworkType"] = request.SrcInstanceNetworkType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceNetwork"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstanceNetwork(request *UpdateInstanceNetworkRequest) (_result *UpdateInstanceNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceNetworkResponse{}
	_body, _err := client.UpdateInstanceNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePrivateRdsClassWithOptions(request *UpdatePrivateRdsClassRequest, runtime *util.RuntimeOptions) (_result *UpdatePrivateRdsClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoUseCoupon)) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PrePayDuration)) {
		query["PrePayDuration"] = request.PrePayDuration
	}

	if !tea.BoolValue(util.IsUnset(request.RdsClass)) {
		query["RdsClass"] = request.RdsClass
	}

	if !tea.BoolValue(util.IsUnset(request.Storage)) {
		query["Storage"] = request.Storage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePrivateRdsClass"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePrivateRdsClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePrivateRdsClass(request *UpdatePrivateRdsClassRequest) (_result *UpdatePrivateRdsClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePrivateRdsClassResponse{}
	_body, _err := client.UpdatePrivateRdsClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResourceGroupAttributeWithOptions(request *UpdateResourceGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceGroupAttribute"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResourceGroupAttribute(request *UpdateResourceGroupAttributeRequest) (_result *UpdateResourceGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResourceGroupAttributeResponse{}
	_body, _err := client.UpdateResourceGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeHiStoreInstanceWithOptions(request *UpgradeHiStoreInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeHiStoreInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.HistoreInstanceId)) {
		query["HistoreInstanceId"] = request.HistoreInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeHiStoreInstance"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeHiStoreInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeHiStoreInstance(request *UpgradeHiStoreInstanceRequest) (_result *UpgradeHiStoreInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeHiStoreInstanceResponse{}
	_body, _err := client.UpgradeHiStoreInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeInstanceVersionWithOptions(request *UpgradeInstanceVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeInstanceVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Rpm)) {
		query["Rpm"] = request.Rpm
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeInstanceVersion"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeInstanceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeInstanceVersion(request *UpgradeInstanceVersionRequest) (_result *UpgradeInstanceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeInstanceVersionResponse{}
	_body, _err := client.UpgradeInstanceVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateShardTaskWithOptions(request *ValidateShardTaskRequest, runtime *util.RuntimeOptions) (_result *ValidateShardTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTableName)) {
		query["SourceTableName"] = request.SourceTableName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetTableName)) {
		query["TargetTableName"] = request.TargetTableName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateShardTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateShardTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateShardTask(request *ValidateShardTaskRequest) (_result *ValidateShardTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateShardTaskResponse{}
	_body, _err := client.ValidateShardTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
