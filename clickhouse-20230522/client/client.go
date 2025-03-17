// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateAccountRequest struct {
	// The name of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The type of the database account. Valid values:
	//
	// 	- **NormalAccount**: standard account
	//
	// 	- **SuperAccount**: privileged account
	//
	// This parameter is required.
	//
	// example:
	//
	// NormalAccount
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The description of the account.
	//
	// example:
	//
	// Used for account
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about permissions.
	DmlAuthSetting *CreateAccountRequestDmlAuthSetting `json:"DmlAuthSetting,omitempty" xml:"DmlAuthSetting,omitempty" type:"Struct"`
	// The password of the database account. The password must meet the following requirements:
	//
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - The following special characters are supported: ! @ # $ % ^ & 	- ( ) _ + - =
	//
	// - The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// a1b2c3d4@
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// clickhouse
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetAccount(v string) *CreateAccountRequest {
	s.Account = &v
	return s
}

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountRequest) SetDBInstanceId(v string) *CreateAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateAccountRequest) SetDescription(v string) *CreateAccountRequest {
	s.Description = &v
	return s
}

func (s *CreateAccountRequest) SetDmlAuthSetting(v *CreateAccountRequestDmlAuthSetting) *CreateAccountRequest {
	s.DmlAuthSetting = v
	return s
}

func (s *CreateAccountRequest) SetPassword(v string) *CreateAccountRequest {
	s.Password = &v
	return s
}

func (s *CreateAccountRequest) SetProduct(v string) *CreateAccountRequest {
	s.Product = &v
	return s
}

func (s *CreateAccountRequest) SetRegionId(v string) *CreateAccountRequest {
	s.RegionId = &v
	return s
}

type CreateAccountRequestDmlAuthSetting struct {
	// The databases on which you want to grant permissions. Separate multiple databases with commas (,).
	AllowDatabases []*string `json:"AllowDatabases,omitempty" xml:"AllowDatabases,omitempty" type:"Repeated"`
	// The dictionaries on which you want to grant permissions. Separate multiple dictionaries with commas (,).
	AllowDictionaries []*string `json:"AllowDictionaries,omitempty" xml:"AllowDictionaries,omitempty" type:"Repeated"`
	// Specifies whether to grant the DDL permissions to the database account. Valid values:
	//
	// 	- **true**: The account has the permissions to execute DDL statements.
	//
	// 	- **false**: The account does not have the permissions to execute DDL statements.
	//
	// example:
	//
	// true
	DdlAuthority *bool `json:"DdlAuthority,omitempty" xml:"DdlAuthority,omitempty"`
	// Specifies whether to grant the DML permissions to the database account. Valid values:
	//
	// 	- **0**: The account has the permissions to read data from the database, write data to the database, and modify the settings of the database.
	//
	// 	- **1**: The account only has the permissions to read data from the database.
	//
	// 	- **2**: The account only has the permissions to read data from the database and modify the settings of the database.
	//
	// example:
	//
	// 0
	DmlAuthority *int32 `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
}

func (s CreateAccountRequestDmlAuthSetting) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequestDmlAuthSetting) GoString() string {
	return s.String()
}

func (s *CreateAccountRequestDmlAuthSetting) SetAllowDatabases(v []*string) *CreateAccountRequestDmlAuthSetting {
	s.AllowDatabases = v
	return s
}

func (s *CreateAccountRequestDmlAuthSetting) SetAllowDictionaries(v []*string) *CreateAccountRequestDmlAuthSetting {
	s.AllowDictionaries = v
	return s
}

func (s *CreateAccountRequestDmlAuthSetting) SetDdlAuthority(v bool) *CreateAccountRequestDmlAuthSetting {
	s.DdlAuthority = &v
	return s
}

func (s *CreateAccountRequestDmlAuthSetting) SetDmlAuthority(v int32) *CreateAccountRequestDmlAuthSetting {
	s.DmlAuthority = &v
	return s
}

type CreateAccountShrinkRequest struct {
	// The name of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The type of the database account. Valid values:
	//
	// 	- **NormalAccount**: standard account
	//
	// 	- **SuperAccount**: privileged account
	//
	// This parameter is required.
	//
	// example:
	//
	// NormalAccount
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The description of the account.
	//
	// example:
	//
	// Used for account
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about permissions.
	DmlAuthSettingShrink *string `json:"DmlAuthSetting,omitempty" xml:"DmlAuthSetting,omitempty"`
	// The password of the database account. The password must meet the following requirements:
	//
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - The following special characters are supported: ! @ # $ % ^ & 	- ( ) _ + - =
	//
	// - The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// a1b2c3d4@
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// clickhouse
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateAccountShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountShrinkRequest) SetAccount(v string) *CreateAccountShrinkRequest {
	s.Account = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetAccountType(v string) *CreateAccountShrinkRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetDBInstanceId(v string) *CreateAccountShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetDescription(v string) *CreateAccountShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetDmlAuthSettingShrink(v string) *CreateAccountShrinkRequest {
	s.DmlAuthSettingShrink = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetPassword(v string) *CreateAccountShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetProduct(v string) *CreateAccountShrinkRequest {
	s.Product = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetRegionId(v string) *CreateAccountShrinkRequest {
	s.RegionId = &v
	return s
}

type CreateAccountResponseBody struct {
	// The data returned.
	Data *CreateAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) SetData(v *CreateAccountResponseBodyData) *CreateAccountResponseBody {
	s.Data = v
	return s
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccountResponseBodyData struct {
	// The name of the database account.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s CreateAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBodyData) SetAccount(v string) *CreateAccountResponseBodyData {
	s.Account = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetDBInstanceId(v string) *CreateAccountResponseBodyData {
	s.DBInstanceId = &v
	return s
}

type CreateAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateDBRequest struct {
	// Database remark information.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database name. The name must meet the following requirements:
	//
	// 	- The name can contain lowercase letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a lowercase letter and end with a lowercase letter or digit.
	//
	// 	- The name can be up to 64 characters in length.
	//
	// >  An underscore (_) is counted as two characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdb001
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateDBRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBRequest) GoString() string {
	return s.String()
}

func (s *CreateDBRequest) SetComment(v string) *CreateDBRequest {
	s.Comment = &v
	return s
}

func (s *CreateDBRequest) SetDBInstanceId(v string) *CreateDBRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBRequest) SetDBName(v string) *CreateDBRequest {
	s.DBName = &v
	return s
}

func (s *CreateDBRequest) SetRegionId(v string) *CreateDBRequest {
	s.RegionId = &v
	return s
}

type CreateDBResponseBody struct {
	// The data returned.
	Data *CreateDBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 94F92113-FF63-5E57-8401-6FE123AD11DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBResponseBody) SetData(v *CreateDBResponseBodyData) *CreateDBResponseBody {
	s.Data = v
	return s
}

func (s *CreateDBResponseBody) SetRequestId(v string) *CreateDBResponseBody {
	s.RequestId = &v
	return s
}

type CreateDBResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testdb001
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s CreateDBResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDBResponseBodyData) SetDBInstanceId(v string) *CreateDBResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBResponseBodyData) SetDBName(v string) *CreateDBResponseBodyData {
	s.DBName = &v
	return s
}

type CreateDBResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResponse) GoString() string {
	return s.String()
}

func (s *CreateDBResponse) SetHeaders(v map[string]*string) *CreateDBResponse {
	s.Headers = v
	return s
}

func (s *CreateDBResponse) SetStatusCode(v int32) *CreateDBResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBResponse) SetBody(v *CreateDBResponseBody) *CreateDBResponse {
	s.Body = v
	return s
}

type CreateDBInstanceRequest struct {
	// The backup set ID.
	//
	// example:
	//
	// 1
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token. Make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// AB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// Used for test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The deployment status of the cluster.
	//
	// example:
	//
	// multi_az
	DeploySchema *string `json:"DeploySchema,omitempty" xml:"DeploySchema,omitempty"`
	// The engine type.
	//
	// Valid values:
	//
	// 	- clickhouse
	//
	// example:
	//
	// clickhouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version.
	//
	// example:
	//
	// 23.8
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The configurations of multi-zone deployment.
	MultiZone []*CreateDBInstanceRequestMultiZone `json:"MultiZone,omitempty" xml:"MultiZone,omitempty" type:"Repeated"`
	// The region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maximum capacity for auto scaling.
	//
	// example:
	//
	// 32
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum capacity for auto scaling.
	//
	// example:
	//
	// 8
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-2ze1*********
	SourceDBInstanceId *string `json:"SourceDBInstanceId,omitempty" xml:"SourceDBInstanceId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-uf6xmupdn7v6ui9f****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf632qye9oqt4x4sr****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) SetBackupSetId(v string) *CreateDBInstanceRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceDescription(v string) *CreateDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDeploySchema(v string) *CreateDBInstanceRequest {
	s.DeploySchema = &v
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

func (s *CreateDBInstanceRequest) SetMultiZone(v []*CreateDBInstanceRequestMultiZone) *CreateDBInstanceRequest {
	s.MultiZone = v
	return s
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetScaleMax(v string) *CreateDBInstanceRequest {
	s.ScaleMax = &v
	return s
}

func (s *CreateDBInstanceRequest) SetScaleMin(v string) *CreateDBInstanceRequest {
	s.ScaleMin = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSourceDBInstanceId(v string) *CreateDBInstanceRequest {
	s.SourceDBInstanceId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVpcId(v string) *CreateDBInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVswitchId(v string) *CreateDBInstanceRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

type CreateDBInstanceRequestMultiZone struct {
	// The vSwitch IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceRequestMultiZone) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceRequestMultiZone) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequestMultiZone) SetVSwitchIds(v []*string) *CreateDBInstanceRequestMultiZone {
	s.VSwitchIds = v
	return s
}

func (s *CreateDBInstanceRequestMultiZone) SetZoneId(v string) *CreateDBInstanceRequestMultiZone {
	s.ZoneId = &v
	return s
}

type CreateDBInstanceShrinkRequest struct {
	// The backup set ID.
	//
	// example:
	//
	// 1
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token. Make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// AB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// Used for test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The deployment status of the cluster.
	//
	// example:
	//
	// multi_az
	DeploySchema *string `json:"DeploySchema,omitempty" xml:"DeploySchema,omitempty"`
	// The engine type.
	//
	// Valid values:
	//
	// 	- clickhouse
	//
	// example:
	//
	// clickhouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version.
	//
	// example:
	//
	// 23.8
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The configurations of multi-zone deployment.
	MultiZoneShrink *string `json:"MultiZone,omitempty" xml:"MultiZone,omitempty"`
	// The region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maximum capacity for auto scaling.
	//
	// example:
	//
	// 32
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum capacity for auto scaling.
	//
	// example:
	//
	// 8
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-2ze1*********
	SourceDBInstanceId *string `json:"SourceDBInstanceId,omitempty" xml:"SourceDBInstanceId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-uf6xmupdn7v6ui9f****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf632qye9oqt4x4sr****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceShrinkRequest) SetBackupSetId(v string) *CreateDBInstanceShrinkRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetClientToken(v string) *CreateDBInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBInstanceDescription(v string) *CreateDBInstanceShrinkRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDeploySchema(v string) *CreateDBInstanceShrinkRequest {
	s.DeploySchema = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetEngine(v string) *CreateDBInstanceShrinkRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetEngineVersion(v string) *CreateDBInstanceShrinkRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetMultiZoneShrink(v string) *CreateDBInstanceShrinkRequest {
	s.MultiZoneShrink = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetRegionId(v string) *CreateDBInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetScaleMax(v string) *CreateDBInstanceShrinkRequest {
	s.ScaleMax = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetScaleMin(v string) *CreateDBInstanceShrinkRequest {
	s.ScaleMin = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetSourceDBInstanceId(v string) *CreateDBInstanceShrinkRequest {
	s.SourceDBInstanceId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetVpcId(v string) *CreateDBInstanceShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetVswitchId(v string) *CreateDBInstanceShrinkRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetZoneId(v string) *CreateDBInstanceShrinkRequest {
	s.ZoneId = &v
	return s
}

type CreateDBInstanceResponseBody struct {
	// The response parameters.
	Data *CreateDBInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
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
	// The endpoint.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****-clickhouse.clickhouseserver.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 21154955706****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDBInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBodyData) SetConnectionString(v string) *CreateDBInstanceResponseBodyData {
	s.ConnectionString = &v
	return s
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

type CreateEndpointRequest struct {
	// The prefix of the new endpoint. The prefix of the ConnectionString parameter.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****-clickhouse.clickhouseserver.rds.aliyuncs.com
	ConnectionPrefix *string `json:"ConnectionPrefix,omitempty" xml:"ConnectionPrefix,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The network type.
	//
	// Valid values:
	//
	// 	- Public
	//
	// example:
	//
	// Public
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateEndpointRequest) SetConnectionPrefix(v string) *CreateEndpointRequest {
	s.ConnectionPrefix = &v
	return s
}

func (s *CreateEndpointRequest) SetDBInstanceId(v string) *CreateEndpointRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateEndpointRequest) SetDBInstanceNetType(v string) *CreateEndpointRequest {
	s.DBInstanceNetType = &v
	return s
}

func (s *CreateEndpointRequest) SetRegionId(v string) *CreateEndpointRequest {
	s.RegionId = &v
	return s
}

type CreateEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEndpointResponseBody) SetRequestId(v string) *CreateEndpointResponseBody {
	s.RequestId = &v
	return s
}

type CreateEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateEndpointResponse) SetHeaders(v map[string]*string) *CreateEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateEndpointResponse) SetStatusCode(v int32) *CreateEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEndpointResponse) SetBody(v *CreateEndpointResponseBody) *CreateEndpointResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	// The destination database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// clickhouse
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) SetAccount(v string) *DeleteAccountRequest {
	s.Account = &v
	return s
}

func (s *DeleteAccountRequest) SetDBInstanceId(v string) *DeleteAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteAccountRequest) SetProduct(v string) *DeleteAccountRequest {
	s.Product = &v
	return s
}

func (s *DeleteAccountRequest) SetRegionId(v string) *DeleteAccountRequest {
	s.RegionId = &v
	return s
}

type DeleteAccountResponseBody struct {
	// The data returned.
	Data *DeleteAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 21D06907-CEA5-561D-B6B1-198BCCE99ED1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) SetData(v *DeleteAccountResponseBodyData) *DeleteAccountResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccountResponseBodyData struct {
	// The name of the account.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DeleteAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBodyData) SetAccount(v string) *DeleteAccountResponseBodyData {
	s.Account = &v
	return s
}

func (s *DeleteAccountResponseBodyData) SetDBInstanceId(v string) *DeleteAccountResponseBodyData {
	s.DBInstanceId = &v
	return s
}

type DeleteAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponse) SetHeaders(v map[string]*string) *DeleteAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccountResponse) SetStatusCode(v int32) *DeleteAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccountResponse) SetBody(v *DeleteAccountResponseBody) *DeleteAccountResponse {
	s.Body = v
	return s
}

type DeleteDBRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the destination database.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdb001
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDBRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBRequest) SetDBInstanceId(v string) *DeleteDBRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBRequest) SetDBName(v string) *DeleteDBRequest {
	s.DBName = &v
	return s
}

func (s *DeleteDBRequest) SetRegionId(v string) *DeleteDBRequest {
	s.RegionId = &v
	return s
}

type DeleteDBResponseBody struct {
	// The data returned.
	Data *DeleteDBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 06798FEE-BEF2-5FAF-A30D-728973BBE97C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBResponseBody) SetData(v *DeleteDBResponseBodyData) *DeleteDBResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDBResponseBody) SetRequestId(v string) *DeleteDBResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testdb001
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DeleteDBResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDBResponseBodyData) SetDBInstanceId(v string) *DeleteDBResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBResponseBodyData) SetDBName(v string) *DeleteDBResponseBodyData {
	s.DBName = &v
	return s
}

type DeleteDBResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBResponse) SetHeaders(v map[string]*string) *DeleteDBResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBResponse) SetStatusCode(v int32) *DeleteDBResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBResponse) SetBody(v *DeleteDBResponseBody) *DeleteDBResponse {
	s.Body = v
	return s
}

type DeleteDBInstanceRequest struct {
	// The ID of the destination cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

type DeleteDBInstanceResponseBody struct {
	// The data returned.
	Data *DeleteDBInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceResponseBody) SetData(v *DeleteDBInstanceResponseBodyData) *DeleteDBInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDBInstanceResponseBody) SetRequestId(v string) *DeleteDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBInstanceResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DeleteDBInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceResponseBodyData) SetDBInstanceId(v string) *DeleteDBInstanceResponseBodyData {
	s.DBInstanceId = &v
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

type DeleteEndpointRequest struct {
	// The prefix of the endpoint, which indicates the prefix of the value of the ConnectionString parameter.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****-clickhouse.clickhouseserver.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId      *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteEndpointRequest) SetConnectionString(v string) *DeleteEndpointRequest {
	s.ConnectionString = &v
	return s
}

func (s *DeleteEndpointRequest) SetDBInstanceId(v string) *DeleteEndpointRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteEndpointRequest) SetDBInstanceNetType(v string) *DeleteEndpointRequest {
	s.DBInstanceNetType = &v
	return s
}

func (s *DeleteEndpointRequest) SetRegionId(v string) *DeleteEndpointRequest {
	s.RegionId = &v
	return s
}

type DeleteEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEndpointResponseBody) SetRequestId(v string) *DeleteEndpointResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteEndpointResponse) SetHeaders(v map[string]*string) *DeleteEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteEndpointResponse) SetStatusCode(v int32) *DeleteEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEndpointResponse) SetBody(v *DeleteEndpointResponseBody) *DeleteEndpointResponse {
	s.Body = v
	return s
}

type DescribeAccountAuthorityRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAccountAuthorityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAuthorityRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountAuthorityRequest) SetAccount(v string) *DescribeAccountAuthorityRequest {
	s.Account = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetDBInstanceId(v string) *DescribeAccountAuthorityRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetRegionId(v string) *DescribeAccountAuthorityRequest {
	s.RegionId = &v
	return s
}

type DescribeAccountAuthorityResponseBody struct {
	// The returned result.
	Data *DescribeAccountAuthorityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountAuthorityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountAuthorityResponseBody) SetData(v *DescribeAccountAuthorityResponseBodyData) *DescribeAccountAuthorityResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetRequestId(v string) *DescribeAccountAuthorityResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAccountAuthorityResponseBodyData struct {
	// The name of the database account.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The databases on which permissions are granted.
	AllowDatabases []*string `json:"AllowDatabases,omitempty" xml:"AllowDatabases,omitempty" type:"Repeated"`
	// The dictionaries on which permissions are granted.
	AllowDictionaries []*string `json:"AllowDictionaries,omitempty" xml:"AllowDictionaries,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Indicates whether the DDL permissions are granted to the database account. Valid values:
	//
	// 	- **true**: The account has the permissions to execute DDL statements.
	//
	// 	- **false**: The account does not have the permissions to execute DDL statements.
	//
	// example:
	//
	// true
	DdlAuthority *bool `json:"DdlAuthority,omitempty" xml:"DdlAuthority,omitempty"`
	// Indicates whether the DML permissions are granted to the database account. Valid values:
	//
	// 	- 0: The account has the permissions to read data from the database, write data to the database, and modify the settings of the database.
	//
	// 	- 1: The account only has the permissions to read data from the database.
	//
	// 	- 2: The account only has the permissions to read data from the database and modify the settings of the database.
	//
	// example:
	//
	// 0
	DmlAuthority *int32 `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
	// All databases.
	TotalDatabases []*string `json:"TotalDatabases,omitempty" xml:"TotalDatabases,omitempty" type:"Repeated"`
	// The database.
	TotalDictionaries []*string `json:"TotalDictionaries,omitempty" xml:"TotalDictionaries,omitempty" type:"Repeated"`
}

func (s DescribeAccountAuthorityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAuthorityResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountAuthorityResponseBodyData) SetAccount(v string) *DescribeAccountAuthorityResponseBodyData {
	s.Account = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) SetAllowDatabases(v []*string) *DescribeAccountAuthorityResponseBodyData {
	s.AllowDatabases = v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) SetAllowDictionaries(v []*string) *DescribeAccountAuthorityResponseBodyData {
	s.AllowDictionaries = v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) SetDBInstanceId(v string) *DescribeAccountAuthorityResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) SetDdlAuthority(v bool) *DescribeAccountAuthorityResponseBodyData {
	s.DdlAuthority = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) SetDmlAuthority(v int32) *DescribeAccountAuthorityResponseBodyData {
	s.DmlAuthority = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) SetTotalDatabases(v []*string) *DescribeAccountAuthorityResponseBodyData {
	s.TotalDatabases = v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) SetTotalDictionaries(v []*string) *DescribeAccountAuthorityResponseBodyData {
	s.TotalDictionaries = v
	return s
}

type DescribeAccountAuthorityResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountAuthorityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAuthorityResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountAuthorityResponse) SetHeaders(v map[string]*string) *DescribeAccountAuthorityResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountAuthorityResponse) SetStatusCode(v int32) *DescribeAccountAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountAuthorityResponse) SetBody(v *DescribeAccountAuthorityResponseBody) *DescribeAccountAuthorityResponse {
	s.Body = v
	return s
}

type DescribeAccountsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// clickhouse
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) SetDBInstanceId(v string) *DescribeAccountsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountsRequest) SetPageNumber(v string) *DescribeAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsRequest) SetPageSize(v string) *DescribeAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountsRequest) SetProduct(v string) *DescribeAccountsRequest {
	s.Product = &v
	return s
}

func (s *DescribeAccountsRequest) SetRegionId(v string) *DescribeAccountsRequest {
	s.RegionId = &v
	return s
}

type DescribeAccountsResponseBody struct {
	// The result returned.
	Data *DescribeAccountsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) SetData(v *DescribeAccountsResponseBodyData) *DescribeAccountsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAccountsResponseBodyData struct {
	// The database accounts.
	Accounts []*DescribeAccountsResponseBodyDataAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccountsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyData) SetAccounts(v []*DescribeAccountsResponseBodyDataAccounts) *DescribeAccountsResponseBodyData {
	s.Accounts = v
	return s
}

func (s *DescribeAccountsResponseBodyData) SetPageNumber(v int32) *DescribeAccountsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsResponseBodyData) SetPageSize(v int32) *DescribeAccountsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountsResponseBodyData) SetTotalCount(v int32) *DescribeAccountsResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeAccountsResponseBodyDataAccounts struct {
	// The username of the database account.
	//
	// example:
	//
	// test
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The type of the database account. Valid values:
	//
	// 	- **1**: standard account
	//
	// 	- **6**: privileged account
	//
	// example:
	//
	// NormalAccount
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The description.
	//
	// example:
	//
	// Used for test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The state of the database account. Valid values:
	//
	// 	- **0**: The database account is being created.
	//
	// 	- **1**: The database account is in use.
	//
	// 	- **3**: The database account is being deleted.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAccountsResponseBodyDataAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyDataAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyDataAccounts) SetAccount(v string) *DescribeAccountsResponseBodyDataAccounts {
	s.Account = &v
	return s
}

func (s *DescribeAccountsResponseBodyDataAccounts) SetAccountType(v string) *DescribeAccountsResponseBodyDataAccounts {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsResponseBodyDataAccounts) SetDescription(v string) *DescribeAccountsResponseBodyDataAccounts {
	s.Description = &v
	return s
}

func (s *DescribeAccountsResponseBodyDataAccounts) SetStatus(v string) *DescribeAccountsResponseBodyDataAccounts {
	s.Status = &v
	return s
}

type DescribeAccountsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeDBInstanceAttributeRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// Valid values:
	//
	// 	- cn-beijing
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

type DescribeDBInstanceAttributeResponseBody struct {
	// The result returned.
	Data *DescribeDBInstanceAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBody) SetData(v *DescribeDBInstanceAttributeResponseBodyData) *DescribeDBInstanceAttributeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyData struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 140692647406****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The channel ID.
	//
	// example:
	//
	// PD39050615820269****
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The billing method. Enterprise Edition clusters use the pay-as-you-go billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the cluster was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2023-09-14T08:14:48Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Indicates whether the release protection feature is enabled for the cluster.
	//
	// example:
	//
	// 0/1
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The deployment mode of the cluster. Valid values: single_az and multi_az.
	//
	// 	- single_az: indicates that the server nodes are deployed in the primary zone. The ID of the primary zone is specified by the ZoneID parameter.
	//
	// 	- multi_az: indicates that the server nodes are deployed in multiple zones. The information about the zones is specified by the MultiZones parameter.
	//
	// The keeper nodes are deployed in multiple zones.
	//
	// example:
	//
	// single_az
	DeploySchema *string `json:"DeploySchema,omitempty" xml:"DeploySchema,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// Used for test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The disabled database ports. Multiple database ports are separated by commas (,).
	//
	// example:
	//
	// 9001,8123
	DisabledPorts *string `json:"DisabledPorts,omitempty" xml:"DisabledPorts,omitempty"`
	// The engine type.
	//
	// example:
	//
	// clickhouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The minor engine version of the cluster.
	//
	// example:
	//
	// 23.8.1.41495_6
	EngineMinorVersion *string `json:"EngineMinorVersion,omitempty" xml:"EngineMinorVersion,omitempty"`
	// The engine version.
	//
	// example:
	//
	// 23.8
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the cluster expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// >  Pay-as-you-go clusters never expire. If the cluster is a pay-as-you-go cluster, an empty string is returned for this parameter.
	//
	// example:
	//
	// 2024-04-17T08:14:48Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The latest minor engine version.
	//
	// example:
	//
	// 23.8.1.41495_6
	LatestEngineMinorVersion *string `json:"LatestEngineMinorVersion,omitempty" xml:"LatestEngineMinorVersion,omitempty"`
	// The lock mode of the cluster.
	//
	// example:
	//
	// 0
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the cluster was locked.
	//
	// example:
	//
	// nolock
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The end time of the maintenance window.
	//
	// example:
	//
	// 21:00
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window.
	//
	// example:
	//
	// 12:00
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The information about the zones.
	MultiZones []*DescribeDBInstanceAttributeResponseBodyDataMultiZones `json:"MultiZones,omitempty" xml:"MultiZones,omitempty" type:"Repeated"`
	// The nodes.
	Nodes []*DescribeDBInstanceAttributeResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The size of the object storage space.
	//
	// example:
	//
	// 13
	ObjectStoreSize *string `json:"ObjectStoreSize,omitempty" xml:"ObjectStoreSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// rg-acfmzygvt54****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The maximum capacity for elastic scaling.
	//
	// example:
	//
	// 32
	ScaleMax *int32 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum capacity for elastic scaling.
	//
	// example:
	//
	// 8
	ScaleMin *int32 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The cluster status.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The size of the storage space. Unit: GB.
	//
	// example:
	//
	// 12
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The storage type.
	//
	// example:
	//
	// 100
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The details of the tags.
	Tags []*DescribeDBInstanceAttributeResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf67ij56zm9x4uc6hmilg
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-wz9duj8xd6r1gzhsg*****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetAliUid(v int64) *DescribeDBInstanceAttributeResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetBid(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.Bid = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetChargeType(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetCreateTime(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetDeletionProtection(v bool) *DescribeDBInstanceAttributeResponseBodyData {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetDeploySchema(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.DeploySchema = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetDescription(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetDisabledPorts(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.DisabledPorts = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetEngine(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetEngineMinorVersion(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.EngineMinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetEngineVersion(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetExpireTime(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetLatestEngineMinorVersion(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.LatestEngineMinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetLockMode(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetLockReason(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetMaintainEndTime(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetMaintainStartTime(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetMultiZones(v []*DescribeDBInstanceAttributeResponseBodyDataMultiZones) *DescribeDBInstanceAttributeResponseBodyData {
	s.MultiZones = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetNodes(v []*DescribeDBInstanceAttributeResponseBodyDataNodes) *DescribeDBInstanceAttributeResponseBodyData {
	s.Nodes = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetObjectStoreSize(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.ObjectStoreSize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetResourceGroupId(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetScaleMax(v int32) *DescribeDBInstanceAttributeResponseBodyData {
	s.ScaleMax = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetScaleMin(v int32) *DescribeDBInstanceAttributeResponseBodyData {
	s.ScaleMin = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetStorageSize(v int32) *DescribeDBInstanceAttributeResponseBodyData {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetStorageType(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetTags(v []*DescribeDBInstanceAttributeResponseBodyDataTags) *DescribeDBInstanceAttributeResponseBodyData {
	s.Tags = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetVpcId(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.ZoneId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDataMultiZones struct {
	// The vSwitch IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDataMultiZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDataMultiZones) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDataMultiZones) SetVSwitchIds(v []*string) *DescribeDBInstanceAttributeResponseBodyDataMultiZones {
	s.VSwitchIds = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDataMultiZones) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyDataMultiZones {
	s.ZoneId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDataNodes struct {
	// The node status.
	//
	// example:
	//
	// active
	NodeStatus *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDataNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDataNodes) SetNodeStatus(v string) *DescribeDBInstanceAttributeResponseBodyDataNodes {
	s.NodeStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDataNodes) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyDataNodes {
	s.ZoneId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDataTags struct {
	// The key of the tag.
	//
	// example:
	//
	// id
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// ck
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDataTags) SetKey(v string) *DescribeDBInstanceAttributeResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDataTags) SetValue(v string) *DescribeDBInstanceAttributeResponseBodyDataTags {
	s.Value = &v
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

type DescribeDBInstanceDataSourcesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database name.
	//
	// example:
	//
	// dbtest
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The table name.
	//
	// example:
	//
	// tableTest
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeDBInstanceDataSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSourcesRequest) SetDBInstanceId(v string) *DescribeDBInstanceDataSourcesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesRequest) SetDBName(v string) *DescribeDBInstanceDataSourcesRequest {
	s.DBName = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesRequest) SetRegionId(v string) *DescribeDBInstanceDataSourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesRequest) SetTableName(v string) *DescribeDBInstanceDataSourcesRequest {
	s.TableName = &v
	return s
}

type DescribeDBInstanceDataSourcesResponseBody struct {
	// The returned result.
	Data *DescribeDBInstanceDataSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F543E6CC-6868-523D-8D28-0E92CF977ED2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceDataSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSourcesResponseBody) SetData(v *DescribeDBInstanceDataSourcesResponseBodyData) *DescribeDBInstanceDataSourcesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBody) SetRequestId(v string) *DescribeDBInstanceDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceDataSourcesResponseBodyData struct {
	// The columns.
	Columns []*DescribeDBInstanceDataSourcesResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The account.
	//
	// example:
	//
	// default
	Schemas *string `json:"Schemas,omitempty" xml:"Schemas,omitempty"`
	// The tables.
	Tables []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceDataSourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSourcesResponseBodyData) SetColumns(v []*DescribeDBInstanceDataSourcesResponseBodyDataColumns) *DescribeDBInstanceDataSourcesResponseBodyData {
	s.Columns = v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyData) SetDBInstanceId(v string) *DescribeDBInstanceDataSourcesResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyData) SetSchemas(v string) *DescribeDBInstanceDataSourcesResponseBodyData {
	s.Schemas = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyData) SetTables(v []*string) *DescribeDBInstanceDataSourcesResponseBodyData {
	s.Tables = v
	return s
}

type DescribeDBInstanceDataSourcesResponseBodyDataColumns struct {
	// The column name.
	//
	// example:
	//
	// c31
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The description of the database account.
	//
	// example:
	//
	// Used for test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The database name.
	//
	// example:
	//
	// dbtest
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// Indicates whether the column is the primary key of the table. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The table name.
	//
	// example:
	//
	// tableTest
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The type of the stored data.
	//
	// example:
	//
	// UInt64
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDBInstanceDataSourcesResponseBodyDataColumns) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataSourcesResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) SetColumnName(v string) *DescribeDBInstanceDataSourcesResponseBodyDataColumns {
	s.ColumnName = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) SetComment(v string) *DescribeDBInstanceDataSourcesResponseBodyDataColumns {
	s.Comment = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) SetDBName(v string) *DescribeDBInstanceDataSourcesResponseBodyDataColumns {
	s.DBName = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) SetPrimaryKey(v string) *DescribeDBInstanceDataSourcesResponseBodyDataColumns {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) SetTableName(v string) *DescribeDBInstanceDataSourcesResponseBodyDataColumns {
	s.TableName = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) SetType(v string) *DescribeDBInstanceDataSourcesResponseBodyDataColumns {
	s.Type = &v
	return s
}

type DescribeDBInstanceDataSourcesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceDataSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSourcesResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponse) SetStatusCode(v int32) *DescribeDBInstanceDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponse) SetBody(v *DescribeDBInstanceDataSourcesResponseBody) *DescribeDBInstanceDataSourcesResponse {
	s.Body = v
	return s
}

type DescribeDBInstancesRequest struct {
	// The cluster IDs. Separate multiple cluster IDs with commas (,).
	//
	// example:
	//
	// cc-xxxxx,cx-xxxx
	DBInstanceIds *string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	// The cluster status.
	//
	// example:
	//
	// active
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDBInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) SetDBInstanceIds(v string) *DescribeDBInstancesRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceStatus(v string) *DescribeDBInstancesRequest {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDescription(v string) *DescribeDBInstancesRequest {
	s.Description = &v
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

type DescribeDBInstancesResponseBody struct {
	// The returned result.
	Data *DescribeDBInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBody) SetData(v *DescribeDBInstancesResponseBodyData) *DescribeDBInstancesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetRequestId(v string) *DescribeDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstancesResponseBodyData struct {
	// The clusters.
	DBInstances []*DescribeDBInstancesResponseBodyDataDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyData) SetDBInstances(v []*DescribeDBInstancesResponseBodyDataDBInstances) *DescribeDBInstancesResponseBodyData {
	s.DBInstances = v
	return s
}

func (s *DescribeDBInstancesResponseBodyData) SetPageNumber(v int32) *DescribeDBInstancesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyData) SetPageSize(v int32) *DescribeDBInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyData) SetTotalCount(v string) *DescribeDBInstancesResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeDBInstancesResponseBodyDataDBInstances struct {
	// The user ID.
	//
	// example:
	//
	// 1294****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The channel ID.
	//
	// example:
	//
	// 186681****
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PrePaid: subscription
	//
	// 	- PostPaid: pay-as-you-go
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2022-12-04 21:16:15
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Indicates whether the release protection feature is enabled for the cluster.
	//
	// example:
	//
	// False
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// test_desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The engine type.
	//
	// example:
	//
	// clickhouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version.
	//
	// example:
	//
	// 22.8
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the cluster expires.
	//
	// example:
	//
	// 2024-02-16 11:51:06
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The lock mode.
	//
	// example:
	//
	// 0
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the cluster was locked.
	//
	// example:
	//
	// null
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The end time of the maintenance window.
	//
	// example:
	//
	// 04:00:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window.
	//
	// example:
	//
	// 00:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The maximum capacity for elastic scaling.
	//
	// example:
	//
	// 13
	ScaleMax *int32 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum capacity for elastic scaling.
	//
	// example:
	//
	// 1
	ScaleMin *int32 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The cluster status.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*DescribeDBInstancesResponseBodyDataDBInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-8vb5mw****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-uf6kg****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDataDBInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDataDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetAliUid(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.AliUid = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetBid(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.Bid = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetChargeType(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetCreateTime(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetDBInstanceId(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetDeletionProtection(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetDescription(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.Description = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetEngine(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetEngineVersion(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetExpireTime(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetLockMode(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetLockReason(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetMaintainEndTime(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetMaintainStartTime(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetRegionId(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetResourceGroupId(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetScaleMax(v int32) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.ScaleMax = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetScaleMin(v int32) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.ScaleMin = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetStatus(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.Status = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetTags(v []*DescribeDBInstancesResponseBodyDataDBInstancesTags) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.Tags = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetVSwitchId(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetVpcId(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetZoneId(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.ZoneId = &v
	return s
}

type DescribeDBInstancesResponseBodyDataDBInstancesTags struct {
	// The tag key.
	//
	// example:
	//
	// tag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDataDBInstancesTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDataDBInstancesTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDataDBInstancesTags) SetKey(v string) *DescribeDBInstancesResponseBodyDataDBInstancesTags {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstancesTags) SetValue(v string) *DescribeDBInstancesResponseBodyDataDBInstancesTags {
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

type DescribeEndpointsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsRequest) SetDBInstanceId(v string) *DescribeEndpointsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeEndpointsRequest) SetRegionId(v string) *DescribeEndpointsRequest {
	s.RegionId = &v
	return s
}

type DescribeEndpointsResponseBody struct {
	// The returned result.
	Data *DescribeEndpointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBody) SetData(v *DescribeEndpointsResponseBodyData) *DescribeEndpointsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEndpointsResponseBody) SetRequestId(v string) *DescribeEndpointsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEndpointsResponseBodyData struct {
	// The details of the endpoints.
	Endpoints []*DescribeEndpointsResponseBodyDataEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The network type of the cluster. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **PUBLIC**
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
}

func (s DescribeEndpointsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyData) SetEndpoints(v []*DescribeEndpointsResponseBodyDataEndpoints) *DescribeEndpointsResponseBodyData {
	s.Endpoints = v
	return s
}

func (s *DescribeEndpointsResponseBodyData) SetInstanceNetworkType(v string) *DescribeEndpointsResponseBodyData {
	s.InstanceNetworkType = &v
	return s
}

type DescribeEndpointsResponseBodyDataEndpoints struct {
	// The endpoint of the cluster.
	//
	// example:
	//
	// cc-****-clickhouse.clickhouseserver.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 172.30.XX.XX
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// 	- VPC
	//
	// 	- PUBLIC
	//
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The details of the ports.
	Ports []*DescribeEndpointsResponseBodyDataEndpointsPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The state of the cluster.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-0xi8829****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-uf61z****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf61z****
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeEndpointsResponseBodyDataEndpoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponseBodyDataEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetConnectionString(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.ConnectionString = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetIPAddress(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.IPAddress = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetNetType(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.NetType = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetPorts(v []*DescribeEndpointsResponseBodyDataEndpointsPorts) *DescribeEndpointsResponseBodyDataEndpoints {
	s.Ports = v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetStatus(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.Status = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetVSwitchId(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.VSwitchId = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetVpcId(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.VpcId = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetVpcInstanceId(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.VpcInstanceId = &v
	return s
}

type DescribeEndpointsResponseBodyDataEndpointsPorts struct {
	// The port used to connect to the cluster. Valid values:
	//
	// 	- 8123: This value is returned when the value of Protocol is HttpPort.
	//
	// 	- 8443: This value is returned when the value of Protocol is HttpsPort.
	//
	// 	- 9000: This value is returned when the value of Protocol is TcpPort.
	//
	// example:
	//
	// 8123
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- HttpPort
	//
	// 	- HttpsPort
	//
	// 	- TcpPort
	//
	// example:
	//
	// HttpPort
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeEndpointsResponseBodyDataEndpointsPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponseBodyDataEndpointsPorts) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyDataEndpointsPorts) SetPort(v int32) *DescribeEndpointsResponseBodyDataEndpointsPorts {
	s.Port = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpointsPorts) SetProtocol(v string) *DescribeEndpointsResponseBodyDataEndpointsPorts {
	s.Protocol = &v
	return s
}

type DescribeEndpointsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponse) SetHeaders(v map[string]*string) *DescribeEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndpointsResponse) SetStatusCode(v int32) *DescribeEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndpointsResponse) SetBody(v *DescribeEndpointsResponseBody) *DescribeEndpointsResponse {
	s.Body = v
	return s
}

type DescribeProcessListRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The query ID.
	//
	// example:
	//
	// 1
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	// The user who executes the query statement.
	//
	// example:
	//
	// testuser
	InitialUser *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	// The keyword of the query statement.
	//
	// example:
	//
	// SELECT
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The execution duration of slow SQL queries. Minimum value: 1000. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	QueryDurationMs *string `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// Specifies the columns by which the query results are sorted in descending order.
	//
	// 	- 0: The query results are sorted by the query_duration_ms column.
	//
	// 	- 1: The query results are sorted by the query_duration_ms and query_start_time columns.
	//
	// 	- 2: The query results are sorted by the query_duration_ms, query_start_time, and user columns.
	//
	// example:
	//
	// id
	QueryOrder *int64 `json:"QueryOrder,omitempty" xml:"QueryOrder,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeProcessListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListRequest) GoString() string {
	return s.String()
}

func (s *DescribeProcessListRequest) SetDBInstanceId(v string) *DescribeProcessListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeProcessListRequest) SetInitialQueryId(v string) *DescribeProcessListRequest {
	s.InitialQueryId = &v
	return s
}

func (s *DescribeProcessListRequest) SetInitialUser(v string) *DescribeProcessListRequest {
	s.InitialUser = &v
	return s
}

func (s *DescribeProcessListRequest) SetKeyword(v string) *DescribeProcessListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeProcessListRequest) SetPageNumber(v int32) *DescribeProcessListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessListRequest) SetPageSize(v int32) *DescribeProcessListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessListRequest) SetQueryDurationMs(v string) *DescribeProcessListRequest {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeProcessListRequest) SetQueryOrder(v int64) *DescribeProcessListRequest {
	s.QueryOrder = &v
	return s
}

func (s *DescribeProcessListRequest) SetRegionId(v string) *DescribeProcessListRequest {
	s.RegionId = &v
	return s
}

type DescribeProcessListResponseBody struct {
	// The data returned.
	Data *DescribeProcessListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProcessListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBody) SetData(v *DescribeProcessListResponseBodyData) *DescribeProcessListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeProcessListResponseBody) SetRequestId(v string) *DescribeProcessListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProcessListResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-xxxx
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The result sets.
	ResultSet []*DescribeProcessListResponseBodyDataResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProcessListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyData) SetDBInstanceID(v int32) *DescribeProcessListResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *DescribeProcessListResponseBodyData) SetDBInstanceName(v string) *DescribeProcessListResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeProcessListResponseBodyData) SetResultSet(v []*DescribeProcessListResponseBodyDataResultSet) *DescribeProcessListResponseBodyData {
	s.ResultSet = v
	return s
}

func (s *DescribeProcessListResponseBodyData) SetTotalCount(v int32) *DescribeProcessListResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeProcessListResponseBodyDataResultSet struct {
	// The address to which the query statement is sent.
	//
	// example:
	//
	// 0:0:0:0:0:ffff:1edd65ea
	InitialAddress *string `json:"InitialAddress,omitempty" xml:"InitialAddress,omitempty"`
	// The query ID.
	//
	// example:
	//
	// \\"79f7e40b-87e2-4ef4-b6df-21889a3a030e\\"
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	// The user who executes the query statement.
	//
	// example:
	//
	// bany
	InitialUser *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	// The query statement that is running.
	//
	// example:
	//
	// select 	- from test
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The minimum query duration. Minimum value: **1000**. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	QueryDurationMs *int64 `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// The beginning of the time range to query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-07-23T10:13:23Z
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
}

func (s DescribeProcessListResponseBodyDataResultSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBodyDataResultSet) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyDataResultSet) SetInitialAddress(v string) *DescribeProcessListResponseBodyDataResultSet {
	s.InitialAddress = &v
	return s
}

func (s *DescribeProcessListResponseBodyDataResultSet) SetInitialQueryId(v string) *DescribeProcessListResponseBodyDataResultSet {
	s.InitialQueryId = &v
	return s
}

func (s *DescribeProcessListResponseBodyDataResultSet) SetInitialUser(v string) *DescribeProcessListResponseBodyDataResultSet {
	s.InitialUser = &v
	return s
}

func (s *DescribeProcessListResponseBodyDataResultSet) SetQuery(v string) *DescribeProcessListResponseBodyDataResultSet {
	s.Query = &v
	return s
}

func (s *DescribeProcessListResponseBodyDataResultSet) SetQueryDurationMs(v int64) *DescribeProcessListResponseBodyDataResultSet {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeProcessListResponseBodyDataResultSet) SetQueryStartTime(v string) *DescribeProcessListResponseBodyDataResultSet {
	s.QueryStartTime = &v
	return s
}

type DescribeProcessListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProcessListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProcessListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponse) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponse) SetHeaders(v map[string]*string) *DescribeProcessListResponse {
	s.Headers = v
	return s
}

func (s *DescribeProcessListResponse) SetStatusCode(v int32) *DescribeProcessListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProcessListResponse) SetBody(v *DescribeProcessListResponseBody) *DescribeProcessListResponse {
	s.Body = v
	return s
}

type DescribeSecurityIPListRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

type DescribeSecurityIPListResponseBody struct {
	// The data returned.
	Data *DescribeSecurityIPListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSecurityIPListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIPListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListResponseBody) SetData(v *DescribeSecurityIPListResponseBodyData) *DescribeSecurityIPListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSecurityIPListResponseBody) SetRequestId(v string) *DescribeSecurityIPListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSecurityIPListResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// TestCluster
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The details about the whitelists.
	GroupItems []*DescribeSecurityIPListResponseBodyDataGroupItems `json:"GroupItems,omitempty" xml:"GroupItems,omitempty" type:"Repeated"`
}

func (s DescribeSecurityIPListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIPListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListResponseBodyData) SetDBInstanceID(v int32) *DescribeSecurityIPListResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyData) SetDBInstanceName(v string) *DescribeSecurityIPListResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyData) SetGroupItems(v []*DescribeSecurityIPListResponseBodyDataGroupItems) *DescribeSecurityIPListResponseBodyData {
	s.GroupItems = v
	return s
}

type DescribeSecurityIPListResponseBodyDataGroupItems struct {
	// The name of the whitelist.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tag of the whitelist.
	//
	// example:
	//
	// test
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// The IP addresses and CIDR blocks in the whitelist.
	//
	// example:
	//
	// 127.0.XX.XX
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The IP address type.
	//
	// example:
	//
	// ipv4
	SecurityIPType *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
	// The network type of the whitelist.
	//
	// example:
	//
	// mix
	WhitelistNetType *string `json:"WhitelistNetType,omitempty" xml:"WhitelistNetType,omitempty"`
}

func (s DescribeSecurityIPListResponseBodyDataGroupItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIPListResponseBodyDataGroupItems) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) SetGroupName(v string) *DescribeSecurityIPListResponseBodyDataGroupItems {
	s.GroupName = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) SetGroupTag(v string) *DescribeSecurityIPListResponseBodyDataGroupItems {
	s.GroupTag = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) SetSecurityIPList(v string) *DescribeSecurityIPListResponseBodyDataGroupItems {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) SetSecurityIPType(v string) *DescribeSecurityIPListResponseBodyDataGroupItems {
	s.SecurityIPType = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) SetWhitelistNetType(v string) *DescribeSecurityIPListResponseBodyDataGroupItems {
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

type DescribeSlowLogRecordsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-dd hh:mm:ss format. The time must be in UTC.
	//
	// example:
	//
	// 2023-09-15 16:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- 30 (default)
	//
	// 	- 50
	//
	// 	- 100
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The execution duration of slow SQL queries. Minimum value: **1000**. Unit: milliseconds.
	//
	// example:
	//
	// 3000
	QueryDurationMs *string `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-dd hh:mm:ss format. The time must be in UTC.
	//
	// example:
	//
	// 2023-09-11 16:00:00
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

func (s *DescribeSlowLogRecordsRequest) SetQueryDurationMs(v string) *DescribeSlowLogRecordsRequest {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetRegionId(v string) *DescribeSlowLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetStartTime(v string) *DescribeSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

type DescribeSlowLogRecordsResponseBody struct {
	// The data returned.
	Data *DescribeSlowLogRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DF203CC8-5F68-5E3F-8050-3C77DD65731A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) SetData(v *DescribeSlowLogRecordsResponseBodyData) *DescribeSlowLogRecordsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSlowLogRecordsResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z32****
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// TestCluster
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The result sets.
	ResultSet []*DescribeSlowLogRecordsResponseBodyDataResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetDBInstanceID(v int32) *DescribeSlowLogRecordsResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetDBInstanceName(v string) *DescribeSlowLogRecordsResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetResultSet(v []*DescribeSlowLogRecordsResponseBodyDataResultSet) *DescribeSlowLogRecordsResponseBodyData {
	s.ResultSet = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetTotalCount(v int32) *DescribeSlowLogRecordsResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeSlowLogRecordsResponseBodyDataResultSet struct {
	// The address to which the query statement is sent.
	//
	// example:
	//
	// 0:0:0:0:0:ffff:1edd65ea
	InitialAddress *string `json:"InitialAddress,omitempty" xml:"InitialAddress,omitempty"`
	// The query ID.
	//
	// example:
	//
	// \\"ae915a3ad30e77e67a7215d05b658cc6\\"
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	// The user who executes the query statement.
	//
	// example:
	//
	// bany
	InitialUser *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	// The peak memory usage for the query. Unit: bytes.
	//
	// example:
	//
	// 4941696
	MemoryUsage *int64 `json:"MemoryUsage,omitempty" xml:"MemoryUsage,omitempty"`
	// The query statement that is running.
	//
	// example:
	//
	// select 	- from test
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The execution duration of slow SQL queries. Minimum value: **1000**. Unit: milliseconds.
	//
	// example:
	//
	// 3000
	QueryDurationMs *int64 `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// The beginning of the time range to query. The time is in the yyyy-MM-dd hh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-09-11 16:00:00
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	// The size of the data that is scanned. Unit: bytes.
	//
	// example:
	//
	// 4507128020832
	ReadBytes *int64 `json:"ReadBytes,omitempty" xml:"ReadBytes,omitempty"`
	// The number of read rows.
	//
	// example:
	//
	// 10
	ReadRows *int64 `json:"ReadRows,omitempty" xml:"ReadRows,omitempty"`
	// The size of the result data. Unit: bytes.
	//
	// example:
	//
	// 10
	ResultBytes *int64 `json:"ResultBytes,omitempty" xml:"ResultBytes,omitempty"`
	// The type of the slow query logs.
	//
	// example:
	//
	// ExceptionWhileProcessing
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyDataResultSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyDataResultSet) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetInitialAddress(v string) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.InitialAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetInitialQueryId(v string) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.InitialQueryId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetInitialUser(v string) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.InitialUser = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetMemoryUsage(v int64) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.MemoryUsage = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetQuery(v string) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.Query = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetQueryDurationMs(v int64) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetQueryStartTime(v string) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.QueryStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetReadBytes(v int64) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.ReadBytes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetReadRows(v int64) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.ReadRows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetResultBytes(v int64) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.ResultBytes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetType(v string) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.Type = &v
	return s
}

type DescribeSlowLogRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlowLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeSlowLogTrendRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-dd hh:mm:ss format. The time must be in UTC.
	//
	// example:
	//
	// 2023-06-07 10:03:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// clickhouse
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The execution duration of slow SQL queries. Minimum value: **1000**. Unit: milliseconds.
	//
	// example:
	//
	// 3000
	QueryDurationMs *string `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start of the time range to query. Specify the time in the yyyy-MM-dd hh:mm:ss format. The time must be in UTC.
	//
	// example:
	//
	// 2023-04-13 17:48:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendRequest) SetDBInstanceId(v string) *DescribeSlowLogTrendRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetEndTime(v string) *DescribeSlowLogTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetProduct(v string) *DescribeSlowLogTrendRequest {
	s.Product = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetQueryDurationMs(v string) *DescribeSlowLogTrendRequest {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetRegionId(v string) *DescribeSlowLogTrendRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetStartTime(v string) *DescribeSlowLogTrendRequest {
	s.StartTime = &v
	return s
}

type DescribeSlowLogTrendResponseBody struct {
	// The returned result.
	Data *DescribeSlowLogTrendResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7D3ECB0E-98CA-5E08-A9CA-F70C5A1E9BDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSlowLogTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBody) SetData(v *DescribeSlowLogTrendResponseBodyData) *DescribeSlowLogTrendResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) SetRequestId(v string) *DescribeSlowLogTrendResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSlowLogTrendResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// clusterTest
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The result sets.
	ResultSet []*DescribeSlowLogTrendResponseBodyDataResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodyData) SetDBInstanceID(v int32) *DescribeSlowLogTrendResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyData) SetDBInstanceName(v string) *DescribeSlowLogTrendResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyData) SetResultSet(v []*DescribeSlowLogTrendResponseBodyDataResultSet) *DescribeSlowLogTrendResponseBodyData {
	s.ResultSet = v
	return s
}

type DescribeSlowLogTrendResponseBodyDataResultSet struct {
	// The average execution duration of slow SQL queries. Minimum value: **1000**. Unit: milliseconds.
	//
	// example:
	//
	// 2000
	AvgQueryDurationMs *int64 `json:"AvgQueryDurationMs,omitempty" xml:"AvgQueryDurationMs,omitempty"`
	// The total number of SQL queries within the specified time range.
	//
	// example:
	//
	// 1
	Cnt *int64 `json:"Cnt,omitempty" xml:"Cnt,omitempty"`
	// The maximum execution duration of slow SQL queries. Minimum value: **1000**. Unit: milliseconds.
	//
	// example:
	//
	// 3000
	MaxQueryDurationMs *int64 `json:"MaxQueryDurationMs,omitempty" xml:"MaxQueryDurationMs,omitempty"`
	// The minimum execution duration of slow SQL queries. Minimum value: **1000**. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	MinQueryDurationMs *int64 `json:"MinQueryDurationMs,omitempty" xml:"MinQueryDurationMs,omitempty"`
	// The beginning of the time range to query. The time is in the yyyy-MM-dd hh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-04-13 17:48:00
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
}

func (s DescribeSlowLogTrendResponseBodyDataResultSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodyDataResultSet) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) SetAvgQueryDurationMs(v int64) *DescribeSlowLogTrendResponseBodyDataResultSet {
	s.AvgQueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) SetCnt(v int64) *DescribeSlowLogTrendResponseBodyDataResultSet {
	s.Cnt = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) SetMaxQueryDurationMs(v int64) *DescribeSlowLogTrendResponseBodyDataResultSet {
	s.MaxQueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) SetMinQueryDurationMs(v int64) *DescribeSlowLogTrendResponseBodyDataResultSet {
	s.MinQueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) SetQueryStartTime(v string) *DescribeSlowLogTrendResponseBodyDataResultSet {
	s.QueryStartTime = &v
	return s
}

type DescribeSlowLogTrendResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlowLogTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlowLogTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponse) SetHeaders(v map[string]*string) *DescribeSlowLogTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowLogTrendResponse) SetStatusCode(v int32) *DescribeSlowLogTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowLogTrendResponse) SetBody(v *DescribeSlowLogTrendResponseBody) *DescribeSlowLogTrendResponse {
	s.Body = v
	return s
}

type KillProcessRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The query ID.
	//
	// example:
	//
	// 1
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s KillProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s KillProcessRequest) GoString() string {
	return s.String()
}

func (s *KillProcessRequest) SetDBInstanceId(v string) *KillProcessRequest {
	s.DBInstanceId = &v
	return s
}

func (s *KillProcessRequest) SetInitialQueryId(v string) *KillProcessRequest {
	s.InitialQueryId = &v
	return s
}

func (s *KillProcessRequest) SetRegionId(v string) *KillProcessRequest {
	s.RegionId = &v
	return s
}

type KillProcessResponseBody struct {
	// The data returned.
	Data *KillProcessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KillProcessResponseBody) GoString() string {
	return s.String()
}

func (s *KillProcessResponseBody) SetData(v *KillProcessResponseBodyData) *KillProcessResponseBody {
	s.Data = v
	return s
}

func (s *KillProcessResponseBody) SetRequestId(v string) *KillProcessResponseBody {
	s.RequestId = &v
	return s
}

type KillProcessResponseBodyData struct {
	// The number of queries that are terminated.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s KillProcessResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s KillProcessResponseBodyData) GoString() string {
	return s.String()
}

func (s *KillProcessResponseBodyData) SetCount(v int64) *KillProcessResponseBodyData {
	s.Count = &v
	return s
}

func (s *KillProcessResponseBodyData) SetDBInstanceID(v int32) *KillProcessResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *KillProcessResponseBodyData) SetDBInstanceName(v string) *KillProcessResponseBodyData {
	s.DBInstanceName = &v
	return s
}

type KillProcessResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KillProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s KillProcessResponse) GoString() string {
	return s.String()
}

func (s *KillProcessResponse) SetHeaders(v map[string]*string) *KillProcessResponse {
	s.Headers = v
	return s
}

func (s *KillProcessResponse) SetStatusCode(v int32) *KillProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *KillProcessResponse) SetBody(v *KillProcessResponseBody) *KillProcessResponse {
	s.Body = v
	return s
}

type ModifyAccountAuthorityRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The information about permissions.
	//
	// This parameter is required.
	DmlAuthSetting *ModifyAccountAuthorityRequestDmlAuthSetting `json:"DmlAuthSetting,omitempty" xml:"DmlAuthSetting,omitempty" type:"Struct"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAccountAuthorityRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountAuthorityRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityRequest) SetAccount(v string) *ModifyAccountAuthorityRequest {
	s.Account = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetDBInstanceId(v string) *ModifyAccountAuthorityRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetDmlAuthSetting(v *ModifyAccountAuthorityRequestDmlAuthSetting) *ModifyAccountAuthorityRequest {
	s.DmlAuthSetting = v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetRegionId(v string) *ModifyAccountAuthorityRequest {
	s.RegionId = &v
	return s
}

type ModifyAccountAuthorityRequestDmlAuthSetting struct {
	// The databases on which you want to grant permissions. Separate multiple databases with commas (,).
	AllowDatabases []*string `json:"AllowDatabases,omitempty" xml:"AllowDatabases,omitempty" type:"Repeated"`
	// The dictionaries on which you want to grant permissions. Separate multiple dictionaries with commas (,).
	AllowDictionaries []*string `json:"AllowDictionaries,omitempty" xml:"AllowDictionaries,omitempty" type:"Repeated"`
	// Specifies whether to grant the DDL permissions to the database account. Valid values:
	//
	// 	- **true**: The account has the permissions to execute DDL statements.
	//
	// 	- **false**: The account does not have the permissions to execute DDL statements.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DdlAuthority *bool `json:"DdlAuthority,omitempty" xml:"DdlAuthority,omitempty"`
	// Specifies whether to grant the DML permissions to the database account. Valid values:
	//
	// 	- **0**: The account has the permissions to read data from the database, write data to the database, and modify the settings of the database.
	//
	// 	- **1**: The account only has the permissions to read data from the database.
	//
	// 	- **2**: The account only has the permissions to read data from the database and modify the settings of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	DmlAuthority *int32 `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
}

func (s ModifyAccountAuthorityRequestDmlAuthSetting) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountAuthorityRequestDmlAuthSetting) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityRequestDmlAuthSetting) SetAllowDatabases(v []*string) *ModifyAccountAuthorityRequestDmlAuthSetting {
	s.AllowDatabases = v
	return s
}

func (s *ModifyAccountAuthorityRequestDmlAuthSetting) SetAllowDictionaries(v []*string) *ModifyAccountAuthorityRequestDmlAuthSetting {
	s.AllowDictionaries = v
	return s
}

func (s *ModifyAccountAuthorityRequestDmlAuthSetting) SetDdlAuthority(v bool) *ModifyAccountAuthorityRequestDmlAuthSetting {
	s.DdlAuthority = &v
	return s
}

func (s *ModifyAccountAuthorityRequestDmlAuthSetting) SetDmlAuthority(v int32) *ModifyAccountAuthorityRequestDmlAuthSetting {
	s.DmlAuthority = &v
	return s
}

type ModifyAccountAuthorityShrinkRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The information about permissions.
	//
	// This parameter is required.
	DmlAuthSettingShrink *string `json:"DmlAuthSetting,omitempty" xml:"DmlAuthSetting,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAccountAuthorityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountAuthorityShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityShrinkRequest) SetAccount(v string) *ModifyAccountAuthorityShrinkRequest {
	s.Account = &v
	return s
}

func (s *ModifyAccountAuthorityShrinkRequest) SetDBInstanceId(v string) *ModifyAccountAuthorityShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAccountAuthorityShrinkRequest) SetDmlAuthSettingShrink(v string) *ModifyAccountAuthorityShrinkRequest {
	s.DmlAuthSettingShrink = &v
	return s
}

func (s *ModifyAccountAuthorityShrinkRequest) SetRegionId(v string) *ModifyAccountAuthorityShrinkRequest {
	s.RegionId = &v
	return s
}

type ModifyAccountAuthorityResponseBody struct {
	// The result returned.
	Data *ModifyAccountAuthorityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountAuthorityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityResponseBody) SetData(v *ModifyAccountAuthorityResponseBodyData) *ModifyAccountAuthorityResponseBody {
	s.Data = v
	return s
}

func (s *ModifyAccountAuthorityResponseBody) SetRequestId(v string) *ModifyAccountAuthorityResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccountAuthorityResponseBodyData struct {
	// The name of the database account.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ModifyAccountAuthorityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountAuthorityResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityResponseBodyData) SetAccount(v string) *ModifyAccountAuthorityResponseBodyData {
	s.Account = &v
	return s
}

func (s *ModifyAccountAuthorityResponseBodyData) SetDBInstanceId(v string) *ModifyAccountAuthorityResponseBodyData {
	s.DBInstanceId = &v
	return s
}

type ModifyAccountAuthorityResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountAuthorityResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountAuthorityResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityResponse) SetHeaders(v map[string]*string) *ModifyAccountAuthorityResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountAuthorityResponse) SetStatusCode(v int32) *ModifyAccountAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountAuthorityResponse) SetBody(v *ModifyAccountAuthorityResponseBody) *ModifyAccountAuthorityResponse {
	s.Body = v
	return s
}

type ModifyAccountDescriptionRequest struct {
	// The name of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testuser
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The description of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) SetAccount(v string) *ModifyAccountDescriptionRequest {
	s.Account = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDBInstanceId(v string) *ModifyAccountDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDescription(v string) *ModifyAccountDescriptionRequest {
	s.Description = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetRegionId(v string) *ModifyAccountDescriptionRequest {
	s.RegionId = &v
	return s
}

type ModifyAccountDescriptionResponseBody struct {
	// The returned data.
	Data *ModifyAccountDescriptionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponseBody) SetData(v *ModifyAccountDescriptionResponseBodyData) *ModifyAccountDescriptionResponseBody {
	s.Data = v
	return s
}

func (s *ModifyAccountDescriptionResponseBody) SetRequestId(v string) *ModifyAccountDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccountDescriptionResponseBodyData struct {
	// The name of the database account.
	//
	// example:
	//
	// testuser
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ModifyAccountDescriptionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponseBodyData) SetAccount(v string) *ModifyAccountDescriptionResponseBodyData {
	s.Account = &v
	return s
}

func (s *ModifyAccountDescriptionResponseBodyData) SetDBInstanceId(v string) *ModifyAccountDescriptionResponseBodyData {
	s.DBInstanceId = &v
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

type ModifyDBInstanceAttributeRequest struct {
	// The configuration that you want to modify.
	//
	// 	- MaintainTime: the O\\&M time
	//
	// 	- DBInstanceDescription: the cluster name
	//
	// This parameter is required.
	//
	// example:
	//
	// DBInstanceDescription
	AttributeType *string `json:"AttributeType,omitempty" xml:"AttributeType,omitempty"`
	// The new value of the configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// clickhouse
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDBInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAttributeRequest) SetAttributeType(v string) *ModifyDBInstanceAttributeRequest {
	s.AttributeType = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetAttributeValue(v string) *ModifyDBInstanceAttributeRequest {
	s.AttributeValue = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetDBInstanceId(v string) *ModifyDBInstanceAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetProduct(v string) *ModifyDBInstanceAttributeRequest {
	s.Product = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetRegionId(v string) *ModifyDBInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

type ModifyDBInstanceAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
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

type ModifyDBInstanceClassRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maximum capacity for elastic scaling.
	//
	// example:
	//
	// 32
	ScaleMax *int64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum capacity for elastic scaling.
	//
	// example:
	//
	// 2
	ScaleMin *int64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
}

func (s ModifyDBInstanceClassRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassRequest) SetDBInstanceId(v string) *ModifyDBInstanceClassRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetRegionId(v string) *ModifyDBInstanceClassRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetScaleMax(v int64) *ModifyDBInstanceClassRequest {
	s.ScaleMax = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetScaleMin(v int64) *ModifyDBInstanceClassRequest {
	s.ScaleMin = &v
	return s
}

type ModifyDBInstanceClassResponseBody struct {
	// The returned result.
	Data *ModifyDBInstanceClassResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceClassResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassResponseBody) SetData(v *ModifyDBInstanceClassResponseBodyData) *ModifyDBInstanceClassResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDBInstanceClassResponseBody) SetRequestId(v string) *ModifyDBInstanceClassResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceClassResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceID *int64 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The maximum capacity for elastic scaling.
	//
	// example:
	//
	// 32
	ScaleMax *int64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum capacity for elastic scaling.
	//
	// example:
	//
	// 2
	ScaleMin *int64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 10000****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDBInstanceClassResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceClassResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassResponseBodyData) SetDBInstanceID(v int64) *ModifyDBInstanceClassResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *ModifyDBInstanceClassResponseBodyData) SetDBInstanceName(v string) *ModifyDBInstanceClassResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceClassResponseBodyData) SetScaleMax(v int64) *ModifyDBInstanceClassResponseBodyData {
	s.ScaleMax = &v
	return s
}

func (s *ModifyDBInstanceClassResponseBodyData) SetScaleMin(v int64) *ModifyDBInstanceClassResponseBodyData {
	s.ScaleMin = &v
	return s
}

func (s *ModifyDBInstanceClassResponseBodyData) SetTaskId(v int64) *ModifyDBInstanceClassResponseBodyData {
	s.TaskId = &v
	return s
}

type ModifyDBInstanceClassResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceClassResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceClassResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceClassResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceClassResponse) SetStatusCode(v int32) *ModifyDBInstanceClassResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceClassResponse) SetBody(v *ModifyDBInstanceClassResponseBody) *ModifyDBInstanceClassResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceConnectionStringRequest struct {
	// The endpoint of the cluster.
	//
	// example:
	//
	// cc-2ze34****-clickhouse..clickhouseserver.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The prefix of the endpoint that is used to connect to the database.
	//
	// example:
	//
	// cc-****-clickhouse
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceId      *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// 	- The database ports that you want to disable. Separate multiple ports with commas (,).
	//
	// 	- This parameter is supported only for clusters whose minor engine version is 24.10.1.11098_1 or later.
	//
	//     **
	//
	//     **Note*	- If you create a cluster whose minor engine version is earlier than 24.10.1.11098_1 and you update the minor engine version to 24.10.1.11098_1 or later, the cluster still does not support this parameter.
	//
	// example:
	//
	// 9001,8123
	DisablePorts *string `json:"DisablePorts,omitempty" xml:"DisablePorts,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDBInstanceConnectionStringRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringRequest) SetConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetConnectionStringPrefix(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceNetType(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceNetType = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDisablePorts(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DisablePorts = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetRegionId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.RegionId = &v
	return s
}

type ModifyDBInstanceConnectionStringResponseBody struct {
	// The data returned.
	Data *ModifyDBInstanceConnectionStringResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceConnectionStringResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetData(v *ModifyDBInstanceConnectionStringResponseBodyData) *ModifyDBInstanceConnectionStringResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetRequestId(v string) *ModifyDBInstanceConnectionStringResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceConnectionStringResponseBodyData struct {
	// The endpoint of the cluster.
	//
	// example:
	//
	// cc-2ze34****-clickhouse..clickhouseserver.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The disabled database ports.
	//
	// example:
	//
	// 9001,8123
	DisabledPorts *string `json:"DisabledPorts,omitempty" xml:"DisabledPorts,omitempty"`
}

func (s ModifyDBInstanceConnectionStringResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetConnectionString(v string) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetDBInstanceID(v int32) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetDBInstanceName(v string) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetDisabledPorts(v string) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.DisabledPorts = &v
	return s
}

type ModifyDBInstanceConnectionStringResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceConnectionStringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ModifySecurityIPListRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the whitelist whose settings you want to modify.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The modification mode.
	//
	// 	- 0: overwrites the original IP addresses and CIDR blocks in the whitelist.
	//
	// 	- 1: adds the IP addresses and CIDR blocks to the whitelist.
	//
	// 	- 2: removes the IP addresses and CIDR blocks from the whitelist.
	//
	// >  We recommend that you set the value to 0.
	//
	// example:
	//
	// 0
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP addresses and CIDR blocks in the whitelist.
	//
	// example:
	//
	// 192.168.0.0/24,172.16.0.0/24
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

func (s *ModifySecurityIPListRequest) SetSecurityIPList(v string) *ModifySecurityIPListRequest {
	s.SecurityIPList = &v
	return s
}

type ModifySecurityIPListResponseBody struct {
	// The returned result.
	Data *ModifySecurityIPListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityIPListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIPListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPListResponseBody) SetData(v *ModifySecurityIPListResponseBodyData) *ModifySecurityIPListResponseBody {
	s.Data = v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetRequestId(v string) *ModifySecurityIPListResponseBody {
	s.RequestId = &v
	return s
}

type ModifySecurityIPListResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-xxxx
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cc-xxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the whitelist.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tag of the whitelist.
	//
	// example:
	//
	// test
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// The IP addresses and CIDR blocks in the whitelist.
	//
	// example:
	//
	// 192.168.0.0/24,172.16.0.0/24
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The IP address type.
	//
	// example:
	//
	// ipv4
	SecurityIPType *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The network type of the whitelist.
	//
	// example:
	//
	// mix
	WhitelistNetType *string `json:"WhitelistNetType,omitempty" xml:"WhitelistNetType,omitempty"`
}

func (s ModifySecurityIPListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIPListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPListResponseBodyData) SetDBInstanceID(v int32) *ModifySecurityIPListResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) SetDBInstanceName(v string) *ModifySecurityIPListResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) SetGroupName(v string) *ModifySecurityIPListResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) SetGroupTag(v string) *ModifySecurityIPListResponseBodyData {
	s.GroupTag = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) SetSecurityIPList(v string) *ModifySecurityIPListResponseBodyData {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) SetSecurityIPType(v string) *ModifySecurityIPListResponseBodyData {
	s.SecurityIPType = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) SetTaskId(v int32) *ModifySecurityIPListResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) SetWhitelistNetType(v string) *ModifySecurityIPListResponseBodyData {
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

type ResetAccountPasswordRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The password of the database account. The password must meet the following requirements:
	//
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - The following special characters are supported: ! @ # $ % ^ & 	- ( ) _ + - =
	//
	// - The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Aa
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The service name.
	//
	// example:
	//
	// clickhouse
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRequest) SetAccount(v string) *ResetAccountPasswordRequest {
	s.Account = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetDBInstanceId(v string) *ResetAccountPasswordRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetPassword(v string) *ResetAccountPasswordRequest {
	s.Password = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetProduct(v string) *ResetAccountPasswordRequest {
	s.Product = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetRegionId(v string) *ResetAccountPasswordRequest {
	s.RegionId = &v
	return s
}

type ResetAccountPasswordResponseBody struct {
	// The result returned.
	Data *ResetAccountPasswordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5A6A077A-577C-536E-AC13-8E715D7A34C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) SetData(v *ResetAccountPasswordResponseBodyData) *ResetAccountPasswordResponseBody {
	s.Data = v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ResetAccountPasswordResponseBodyData struct {
	// The name of the account.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ResetAccountPasswordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBodyData) SetAccount(v string) *ResetAccountPasswordResponseBodyData {
	s.Account = &v
	return s
}

func (s *ResetAccountPasswordResponseBodyData) SetDBInstanceId(v string) *ResetAccountPasswordResponseBodyData {
	s.DBInstanceId = &v
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

type RestartDBInstanceRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RestartDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceRequest) SetDBInstanceId(v string) *RestartDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RestartDBInstanceRequest) SetRegionId(v string) *RestartDBInstanceRequest {
	s.RegionId = &v
	return s
}

type RestartDBInstanceResponseBody struct {
	// The data returned.
	Data *RestartDBInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceResponseBody) SetData(v *RestartDBInstanceResponseBodyData) *RestartDBInstanceResponseBody {
	s.Data = v
	return s
}

func (s *RestartDBInstanceResponseBody) SetRequestId(v string) *RestartDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartDBInstanceResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceID *int64 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test1
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 100001080
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RestartDBInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceResponseBodyData) SetDBInstanceID(v int64) *RestartDBInstanceResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *RestartDBInstanceResponseBodyData) SetDBInstanceName(v string) *RestartDBInstanceResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *RestartDBInstanceResponseBodyData) SetTaskId(v int64) *RestartDBInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

type RestartDBInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type StartDBInstanceRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartDBInstanceRequest) SetDBInstanceId(v string) *StartDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *StartDBInstanceRequest) SetRegionId(v string) *StartDBInstanceRequest {
	s.RegionId = &v
	return s
}

type StartDBInstanceResponseBody struct {
	// The data returned.
	Data *StartDBInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartDBInstanceResponseBody) SetData(v *StartDBInstanceResponseBodyData) *StartDBInstanceResponseBody {
	s.Data = v
	return s
}

func (s *StartDBInstanceResponseBody) SetRequestId(v string) *StartDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartDBInstanceResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceID *int64 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test1
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 100000837
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartDBInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StartDBInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartDBInstanceResponseBodyData) SetDBInstanceID(v int64) *StartDBInstanceResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *StartDBInstanceResponseBodyData) SetDBInstanceName(v string) *StartDBInstanceResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *StartDBInstanceResponseBodyData) SetTaskId(v int64) *StartDBInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

type StartDBInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartDBInstanceResponse) SetHeaders(v map[string]*string) *StartDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartDBInstanceResponse) SetStatusCode(v int32) *StartDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDBInstanceResponse) SetBody(v *StartDBInstanceResponseBody) *StartDBInstanceResponse {
	s.Body = v
	return s
}

type StopDBInstanceRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopDBInstanceRequest) SetDBInstanceId(v string) *StopDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *StopDBInstanceRequest) SetRegionId(v string) *StopDBInstanceRequest {
	s.RegionId = &v
	return s
}

type StopDBInstanceResponseBody struct {
	// The data returned.
	Data *StopDBInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopDBInstanceResponseBody) SetData(v *StopDBInstanceResponseBodyData) *StopDBInstanceResponseBody {
	s.Data = v
	return s
}

func (s *StopDBInstanceResponseBody) SetRequestId(v string) *StopDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopDBInstanceResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceID *int64 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test1
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 100000785
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopDBInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StopDBInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *StopDBInstanceResponseBodyData) SetDBInstanceID(v int64) *StopDBInstanceResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *StopDBInstanceResponseBodyData) SetDBInstanceName(v string) *StopDBInstanceResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *StopDBInstanceResponseBodyData) SetTaskId(v int64) *StopDBInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

type StopDBInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopDBInstanceResponse) SetHeaders(v map[string]*string) *StopDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopDBInstanceResponse) SetStatusCode(v int32) *StopDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDBInstanceResponse) SetBody(v *StopDBInstanceResponseBody) *StopDBInstanceResponse {
	s.Body = v
	return s
}

type UpgradeMinorVersionRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1jyis8p15we****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The update time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  If you set SwitchTimeMode to SpecifyTime, you must configure this parameter to specify the update time.
	//
	// example:
	//
	// 2023-01-09T05:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// Specifies whether to update the minor engine version of the cluster immediately. Valid values:
	//
	// 	- **Immediate**: The system immediately performs the update.
	//
	// 	- **MaintainTime**: The system performs the update during the specified maintenance window.
	//
	// 	- **SpecifyTime**: The system performs the update at a specified time.
	//
	// example:
	//
	// Immediate
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	// The minor engine version to which you want to update.
	//
	// >  By default, TargetMinorVersion is not set and the minor engine version of the cluster is updated to the latest version.
	//
	// example:
	//
	// 23.8.1.41495_6
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
}

func (s UpgradeMinorVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionRequest) SetDBInstanceId(v string) *UpgradeMinorVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetRegionId(v string) *UpgradeMinorVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetSwitchTime(v string) *UpgradeMinorVersionRequest {
	s.SwitchTime = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetSwitchTimeMode(v string) *UpgradeMinorVersionRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetTargetMinorVersion(v string) *UpgradeMinorVersionRequest {
	s.TargetMinorVersion = &v
	return s
}

type UpgradeMinorVersionResponseBody struct {
	// The returned result.
	Data *UpgradeMinorVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FE242962-6DA3-5FC8-9691-37B62A3210F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeMinorVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMinorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionResponseBody) SetData(v *UpgradeMinorVersionResponseBodyData) *UpgradeMinorVersionResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeMinorVersionResponseBody) SetRequestId(v string) *UpgradeMinorVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeMinorVersionResponseBodyData struct {
	// The instance ID.
	//
	// example:
	//
	// cc-uf6x229yeq166****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s UpgradeMinorVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMinorVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionResponseBodyData) SetDBInstanceName(v string) *UpgradeMinorVersionResponseBodyData {
	s.DBInstanceName = &v
	return s
}

type UpgradeMinorVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeMinorVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeMinorVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMinorVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionResponse) SetHeaders(v map[string]*string) *UpgradeMinorVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeMinorVersionResponse) SetStatusCode(v int32) *UpgradeMinorVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeMinorVersionResponse) SetBody(v *UpgradeMinorVersionResponseBody) *UpgradeMinorVersionResponse {
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
		"ap-northeast-2-pop":          tea.String("clickhouse.aliyuncs.com"),
		"ap-southeast-1":              tea.String("clickhouse.aliyuncs.com"),
		"cn-beijing":                  tea.String("clickhouse.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("clickhouse.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("clickhouse.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("clickhouse.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("clickhouse.aliyuncs.com"),
		"cn-edge-1":                   tea.String("clickhouse.aliyuncs.com"),
		"cn-fujian":                   tea.String("clickhouse.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("clickhouse.aliyuncs.com"),
		"cn-hongkong":                 tea.String("clickhouse.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("clickhouse.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("clickhouse.aliyuncs.com"),
		"cn-qingdao":                  tea.String("clickhouse.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("clickhouse.aliyuncs.com"),
		"cn-shanghai":                 tea.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("clickhouse.aliyuncs.com"),
		"cn-wuhan":                    tea.String("clickhouse.aliyuncs.com"),
		"cn-yushanfang":               tea.String("clickhouse.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("clickhouse.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("clickhouse.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("clickhouse.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("clickhouse.aliyuncs.com"),
		"me-east-1":                   tea.String("clickhouse.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("clickhouse.aliyuncs.com"),
		"us-east-1":                   tea.String("clickhouse.aliyuncs.com"),
		"us-west-1":                   tea.String("clickhouse.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("clickhouse"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates a database account for an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param tmpReq - CreateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountResponse
func (client *Client) CreateAccountWithOptions(tmpReq *CreateAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAccountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DmlAuthSetting)) {
		request.DmlAuthSettingShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DmlAuthSetting, tea.String("DmlAuthSetting"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DmlAuthSettingShrink)) {
		query["DmlAuthSetting"] = request.DmlAuthSettingShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccount"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAccountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAccountResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a database account for an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - CreateAccountRequest
//
// @return CreateAccountResponse
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

// Summary:
//
// Creates an ApsaraDB for ClickHouse database.
//
// @param request - CreateDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBResponse
func (client *Client) CreateDBWithOptions(request *CreateDBRequest, runtime *util.RuntimeOptions) (_result *CreateDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDB"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateDBResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateDBResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates an ApsaraDB for ClickHouse database.
//
// @param request - CreateDBRequest
//
// @return CreateDBResponse
func (client *Client) CreateDB(request *CreateDBRequest) (_result *CreateDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBResponse{}
	_body, _err := client.CreateDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an ApsaraDB for ClickHouse cluster that runs Enterprise Edition.
//
// @param tmpReq - CreateDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstanceResponse
func (client *Client) CreateDBInstanceWithOptions(tmpReq *CreateDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDBInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MultiZone)) {
		request.MultiZoneShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MultiZone, tea.String("MultiZone"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupSetId)) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceDescription)) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DeploySchema)) {
		query["DeploySchema"] = request.DeploySchema
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.MultiZoneShrink)) {
		query["MultiZone"] = request.MultiZoneShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleMax)) {
		query["ScaleMax"] = request.ScaleMax
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleMin)) {
		query["ScaleMin"] = request.ScaleMin
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDBInstanceId)) {
		query["SourceDBInstanceId"] = request.SourceDBInstanceId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateDBInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateDBInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates an ApsaraDB for ClickHouse cluster that runs Enterprise Edition.
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
// Applies for a public endpoint.
//
// @param request - CreateEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEndpointResponse
func (client *Client) CreateEndpointWithOptions(request *CreateEndpointRequest, runtime *util.RuntimeOptions) (_result *CreateEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionPrefix)) {
		query["ConnectionPrefix"] = request.ConnectionPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceNetType)) {
		query["DBInstanceNetType"] = request.DBInstanceNetType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEndpoint"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateEndpointResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateEndpointResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Applies for a public endpoint.
//
// @param request - CreateEndpointRequest
//
// @return CreateEndpointResponse
func (client *Client) CreateEndpoint(request *CreateEndpointRequest) (_result *CreateEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEndpointResponse{}
	_body, _err := client.CreateEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a database account from an ApsaraDB for ClickHouse cluster.
//
// @param request - DeleteAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccount"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAccountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAccountResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a database account from an ApsaraDB for ClickHouse cluster.
//
// @param request - DeleteAccountRequest
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an ApsaraDB for ClickHouse database.
//
// @param request - DeleteDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBResponse
func (client *Client) DeleteDBWithOptions(request *DeleteDBRequest, runtime *util.RuntimeOptions) (_result *DeleteDBResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDB"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteDBResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteDBResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes an ApsaraDB for ClickHouse database.
//
// @param request - DeleteDBRequest
//
// @return DeleteDBResponse
func (client *Client) DeleteDB(request *DeleteDBRequest) (_result *DeleteDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBResponse{}
	_body, _err := client.DeleteDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases an ApsaraDB for ClickHouse Enterprise Edition cluster.
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteDBInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteDBInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Releases an ApsaraDB for ClickHouse Enterprise Edition cluster.
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
// Releases a public endpoint.
//
// @param request - DeleteEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEndpointResponse
func (client *Client) DeleteEndpointWithOptions(request *DeleteEndpointRequest, runtime *util.RuntimeOptions) (_result *DeleteEndpointResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.DBInstanceNetType)) {
		query["DBInstanceNetType"] = request.DBInstanceNetType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEndpoint"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteEndpointResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteEndpointResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Releases a public endpoint.
//
// @param request - DeleteEndpointRequest
//
// @return DeleteEndpointResponse
func (client *Client) DeleteEndpoint(request *DeleteEndpointRequest) (_result *DeleteEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEndpointResponse{}
	_body, _err := client.DeleteEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the permissions of a database account.
//
// @param request - DescribeAccountAuthorityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountAuthorityResponse
func (client *Client) DescribeAccountAuthorityWithOptions(request *DescribeAccountAuthorityRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountAuthorityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
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
		Action:      tea.String("DescribeAccountAuthority"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeAccountAuthorityResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeAccountAuthorityResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the permissions of a database account.
//
// @param request - DescribeAccountAuthorityRequest
//
// @return DescribeAccountAuthorityResponse
func (client *Client) DescribeAccountAuthority(request *DescribeAccountAuthorityRequest) (_result *DescribeAccountAuthorityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountAuthorityResponse{}
	_body, _err := client.DescribeAccountAuthorityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries database accounts for an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountsResponse
func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccounts"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeAccountsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeAccountsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries database accounts for an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeAccountsRequest
//
// @return DescribeAccountsResponse
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

// Summary:
//
// Queries the details of an ApsaraDB for ClickHouse cluster that runs Enterprise Edition.
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDBInstanceAttributeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDBInstanceAttributeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of an ApsaraDB for ClickHouse cluster that runs Enterprise Edition.
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
// Queries the schema of a database or a table.
//
// @param request - DescribeDBInstanceDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceDataSourcesResponse
func (client *Client) DescribeDBInstanceDataSourcesWithOptions(request *DescribeDBInstanceDataSourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceDataSourcesResponse, _err error) {
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
		Action:      tea.String("DescribeDBInstanceDataSources"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDBInstanceDataSourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDBInstanceDataSourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the schema of a database or a table.
//
// @param request - DescribeDBInstanceDataSourcesRequest
//
// @return DescribeDBInstanceDataSourcesResponse
func (client *Client) DescribeDBInstanceDataSources(request *DescribeDBInstanceDataSourcesRequest) (_result *DescribeDBInstanceDataSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceDataSourcesResponse{}
	_body, _err := client.DescribeDBInstanceDataSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of ApsaraDB for ClickHouse clusters.
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
	if !tea.BoolValue(util.IsUnset(request.DBInstanceIds)) {
		query["DBInstanceIds"] = request.DBInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceStatus)) {
		query["DBInstanceStatus"] = request.DBInstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDBInstancesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDBInstancesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of ApsaraDB for ClickHouse clusters.
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
// Queries the endpoint of an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndpointsResponse
func (client *Client) DescribeEndpointsWithOptions(request *DescribeEndpointsRequest, runtime *util.RuntimeOptions) (_result *DescribeEndpointsResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEndpoints"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeEndpointsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeEndpointsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the endpoint of an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeEndpointsRequest
//
// @return DescribeEndpointsResponse
func (client *Client) DescribeEndpoints(request *DescribeEndpointsRequest) (_result *DescribeEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEndpointsResponse{}
	_body, _err := client.DescribeEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Views running queries.
//
// @param request - DescribeProcessListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProcessListResponse
func (client *Client) DescribeProcessListWithOptions(request *DescribeProcessListRequest, runtime *util.RuntimeOptions) (_result *DescribeProcessListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InitialQueryId)) {
		query["InitialQueryId"] = request.InitialQueryId
	}

	if !tea.BoolValue(util.IsUnset(request.InitialUser)) {
		query["InitialUser"] = request.InitialUser
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryDurationMs)) {
		query["QueryDurationMs"] = request.QueryDurationMs
	}

	if !tea.BoolValue(util.IsUnset(request.QueryOrder)) {
		query["QueryOrder"] = request.QueryOrder
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProcessList"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeProcessListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeProcessListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Views running queries.
//
// @param request - DescribeProcessListRequest
//
// @return DescribeProcessListResponse
func (client *Client) DescribeProcessList(request *DescribeProcessListRequest) (_result *DescribeProcessListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProcessListResponse{}
	_body, _err := client.DescribeProcessListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the whitelist of an ApsaraDB for ClickHouse cluster.
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSecurityIPListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSecurityIPListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the whitelist of an ApsaraDB for ClickHouse cluster.
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
// Queries the details of slow query logs.
//
// @param request - DescribeSlowLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowLogRecordsResponse
func (client *Client) DescribeSlowLogRecordsWithOptions(request *DescribeSlowLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.QueryDurationMs)) {
		query["QueryDurationMs"] = request.QueryDurationMs
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
		Action:      tea.String("DescribeSlowLogRecords"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSlowLogRecordsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSlowLogRecordsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of slow query logs.
//
// @param request - DescribeSlowLogRecordsRequest
//
// @return DescribeSlowLogRecordsResponse
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

// Summary:
//
// Queries the trend of slow query logs.
//
// @param request - DescribeSlowLogTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowLogTrendResponse
func (client *Client) DescribeSlowLogTrendWithOptions(request *DescribeSlowLogTrendRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowLogTrendResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.QueryDurationMs)) {
		query["QueryDurationMs"] = request.QueryDurationMs
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
		Action:      tea.String("DescribeSlowLogTrend"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSlowLogTrendResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSlowLogTrendResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the trend of slow query logs.
//
// @param request - DescribeSlowLogTrendRequest
//
// @return DescribeSlowLogTrendResponse
func (client *Client) DescribeSlowLogTrend(request *DescribeSlowLogTrendRequest) (_result *DescribeSlowLogTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowLogTrendResponse{}
	_body, _err := client.DescribeSlowLogTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates an ongoing query.
//
// @param request - KillProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillProcessResponse
func (client *Client) KillProcessWithOptions(request *KillProcessRequest, runtime *util.RuntimeOptions) (_result *KillProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InitialQueryId)) {
		query["InitialQueryId"] = request.InitialQueryId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("KillProcess"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &KillProcessResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &KillProcessResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Terminates an ongoing query.
//
// @param request - KillProcessRequest
//
// @return KillProcessResponse
func (client *Client) KillProcess(request *KillProcessRequest) (_result *KillProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KillProcessResponse{}
	_body, _err := client.KillProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the permissions of a database account.
//
// @param tmpReq - ModifyAccountAuthorityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountAuthorityResponse
func (client *Client) ModifyAccountAuthorityWithOptions(tmpReq *ModifyAccountAuthorityRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountAuthorityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyAccountAuthorityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DmlAuthSetting)) {
		request.DmlAuthSettingShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DmlAuthSetting, tea.String("DmlAuthSetting"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DmlAuthSettingShrink)) {
		query["DmlAuthSetting"] = request.DmlAuthSettingShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountAuthority"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyAccountAuthorityResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyAccountAuthorityResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the permissions of a database account.
//
// @param request - ModifyAccountAuthorityRequest
//
// @return ModifyAccountAuthorityResponse
func (client *Client) ModifyAccountAuthority(request *ModifyAccountAuthorityRequest) (_result *ModifyAccountAuthorityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountAuthorityResponse{}
	_body, _err := client.ModifyAccountAuthorityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of a database account.
//
// @param request - ModifyAccountDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountDescriptionResponse
func (client *Client) ModifyAccountDescriptionWithOptions(request *ModifyAccountDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountDescription"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyAccountDescriptionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyAccountDescriptionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the description of a database account.
//
// @param request - ModifyAccountDescriptionRequest
//
// @return ModifyAccountDescriptionResponse
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

// Summary:
//
// Modifies the configurations of an ApsaraDB for ClickHouse cluster.
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
	if !tea.BoolValue(util.IsUnset(request.AttributeType)) {
		query["AttributeType"] = request.AttributeType
	}

	if !tea.BoolValue(util.IsUnset(request.AttributeValue)) {
		query["AttributeValue"] = request.AttributeValue
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyDBInstanceAttributeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyDBInstanceAttributeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the configurations of an ApsaraDB for ClickHouse cluster.
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
// Modifies the elastic scaling settings of an ApsaraDB for ClickHouse cluster.
//
// @param request - ModifyDBInstanceClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceClassResponse
func (client *Client) ModifyDBInstanceClassWithOptions(request *ModifyDBInstanceClassRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceClassResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ScaleMax)) {
		query["ScaleMax"] = request.ScaleMax
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleMin)) {
		query["ScaleMin"] = request.ScaleMin
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceClass"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyDBInstanceClassResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyDBInstanceClassResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the elastic scaling settings of an ApsaraDB for ClickHouse cluster.
//
// @param request - ModifyDBInstanceClassRequest
//
// @return ModifyDBInstanceClassResponse
func (client *Client) ModifyDBInstanceClass(request *ModifyDBInstanceClassRequest) (_result *ModifyDBInstanceClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceClassResponse{}
	_body, _err := client.ModifyDBInstanceClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the endpoint of an ApsaraDB for ClickHouse cluster.
//
// @param request - ModifyDBInstanceConnectionStringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceConnectionStringResponse
func (client *Client) ModifyDBInstanceConnectionStringWithOptions(request *ModifyDBInstanceConnectionStringRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceConnectionStringResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionString)) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionStringPrefix)) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceNetType)) {
		query["DBInstanceNetType"] = request.DBInstanceNetType
	}

	if !tea.BoolValue(util.IsUnset(request.DisablePorts)) {
		query["DisablePorts"] = request.DisablePorts
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceConnectionString"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyDBInstanceConnectionStringResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyDBInstanceConnectionStringResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the endpoint of an ApsaraDB for ClickHouse cluster.
//
// @param request - ModifyDBInstanceConnectionStringRequest
//
// @return ModifyDBInstanceConnectionStringResponse
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

// Summary:
//
// Modifies the whitelist settings of an ApsaraDB for ClickHouse cluster.
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyMode)) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIPList)) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySecurityIPList"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifySecurityIPListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifySecurityIPListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the whitelist settings of an ApsaraDB for ClickHouse cluster.
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
// Resets the password of a database account for an ApsaraDB for ClickHouse Enterprise Edition cluster.
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAccountPassword"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ResetAccountPasswordResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ResetAccountPasswordResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Resets the password of a database account for an ApsaraDB for ClickHouse Enterprise Edition cluster.
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
// Restarts an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - RestartDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDBInstanceResponse
func (client *Client) RestartDBInstanceWithOptions(request *RestartDBInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartDBInstanceResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartDBInstance"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RestartDBInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RestartDBInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Restarts an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - RestartDBInstanceRequest
//
// @return RestartDBInstanceResponse
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

// Summary:
//
// Starts an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - StartDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDBInstanceResponse
func (client *Client) StartDBInstanceWithOptions(request *StartDBInstanceRequest, runtime *util.RuntimeOptions) (_result *StartDBInstanceResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDBInstance"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StartDBInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StartDBInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Starts an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - StartDBInstanceRequest
//
// @return StartDBInstanceResponse
func (client *Client) StartDBInstance(request *StartDBInstanceRequest) (_result *StartDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDBInstanceResponse{}
	_body, _err := client.StartDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - StopDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDBInstanceResponse
func (client *Client) StopDBInstanceWithOptions(request *StopDBInstanceRequest, runtime *util.RuntimeOptions) (_result *StopDBInstanceResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDBInstance"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StopDBInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StopDBInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Stops an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - StopDBInstanceRequest
//
// @return StopDBInstanceResponse
func (client *Client) StopDBInstance(request *StopDBInstanceRequest) (_result *StopDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDBInstanceResponse{}
	_body, _err := client.StopDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the minor engine version of an ApsaraDB for ClickHouse cluster that runs Enterprise Edition.
//
// @param request - UpgradeMinorVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeMinorVersionResponse
func (client *Client) UpgradeMinorVersionWithOptions(request *UpgradeMinorVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeMinorVersionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SwitchTime)) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTimeMode)) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	if !tea.BoolValue(util.IsUnset(request.TargetMinorVersion)) {
		query["TargetMinorVersion"] = request.TargetMinorVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeMinorVersion"),
		Version:     tea.String("2023-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpgradeMinorVersionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpgradeMinorVersionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates the minor engine version of an ApsaraDB for ClickHouse cluster that runs Enterprise Edition.
//
// @param request - UpgradeMinorVersionRequest
//
// @return UpgradeMinorVersionResponse
func (client *Client) UpgradeMinorVersion(request *UpgradeMinorVersionRequest) (_result *UpgradeMinorVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeMinorVersionResponse{}
	_body, _err := client.UpgradeMinorVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
