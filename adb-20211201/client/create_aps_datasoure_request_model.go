// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsDatasoureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateApsDatasoureRequest
	GetDBClusterId() *string
	SetDatabricksInfo(v *CreateApsDatasoureRequestDatabricksInfo) *CreateApsDatasoureRequest
	GetDatabricksInfo() *CreateApsDatasoureRequestDatabricksInfo
	SetDatasourceDescription(v string) *CreateApsDatasoureRequest
	GetDatasourceDescription() *string
	SetDatasourceName(v string) *CreateApsDatasoureRequest
	GetDatasourceName() *string
	SetDatasourceType(v string) *CreateApsDatasoureRequest
	GetDatasourceType() *string
	SetHiveInfo(v *CreateApsDatasoureRequestHiveInfo) *CreateApsDatasoureRequest
	GetHiveInfo() *CreateApsDatasoureRequestHiveInfo
	SetKafkaInfo(v *CreateApsDatasoureRequestKafkaInfo) *CreateApsDatasoureRequest
	GetKafkaInfo() *CreateApsDatasoureRequestKafkaInfo
	SetMode(v string) *CreateApsDatasoureRequest
	GetMode() *string
	SetPolarDBMysqlInfo(v *CreateApsDatasoureRequestPolarDBMysqlInfo) *CreateApsDatasoureRequest
	GetPolarDBMysqlInfo() *CreateApsDatasoureRequestPolarDBMysqlInfo
	SetPolarDBXInfo(v *CreateApsDatasoureRequestPolarDBXInfo) *CreateApsDatasoureRequest
	GetPolarDBXInfo() *CreateApsDatasoureRequestPolarDBXInfo
	SetRdsMysqlInfo(v *CreateApsDatasoureRequestRdsMysqlInfo) *CreateApsDatasoureRequest
	GetRdsMysqlInfo() *CreateApsDatasoureRequestRdsMysqlInfo
	SetRegionId(v string) *CreateApsDatasoureRequest
	GetRegionId() *string
	SetSlsInfo(v *CreateApsDatasoureRequestSlsInfo) *CreateApsDatasoureRequest
	GetSlsInfo() *CreateApsDatasoureRequestSlsInfo
}

type CreateApsDatasoureRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-xxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The information about the Databricks data source.
	//
	// example:
	//
	// -
	DatabricksInfo *CreateApsDatasoureRequestDatabricksInfo `json:"DatabricksInfo,omitempty" xml:"DatabricksInfo,omitempty" type:"Struct"`
	// The description of the data source.
	//
	// example:
	//
	// description
	DatasourceDescription *string `json:"DatasourceDescription,omitempty" xml:"DatasourceDescription,omitempty"`
	// The name of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls-******
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// The type of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// KAFKA
	DatasourceType *string `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	// The information about the Hive data source.
	//
	// example:
	//
	// -
	HiveInfo *CreateApsDatasoureRequestHiveInfo `json:"HiveInfo,omitempty" xml:"HiveInfo,omitempty" type:"Struct"`
	// The information about the source Apache Kafka instance.
	//
	// example:
	//
	// -
	KafkaInfo *CreateApsDatasoureRequestKafkaInfo `json:"KafkaInfo,omitempty" xml:"KafkaInfo,omitempty" type:"Struct"`
	// The mode.
	//
	// example:
	//
	// ALI_CLOUD_INSTANCE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The information about the source PolarDB for MySQL cluster.
	//
	// example:
	//
	// -
	PolarDBMysqlInfo *CreateApsDatasoureRequestPolarDBMysqlInfo `json:"PolarDBMysqlInfo,omitempty" xml:"PolarDBMysqlInfo,omitempty" type:"Struct"`
	// The information about the source PolarDB-X instance.
	//
	// example:
	//
	// -
	PolarDBXInfo *CreateApsDatasoureRequestPolarDBXInfo `json:"PolarDBXInfo,omitempty" xml:"PolarDBXInfo,omitempty" type:"Struct"`
	// The information about the source ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// -
	RdsMysqlInfo *CreateApsDatasoureRequestRdsMysqlInfo `json:"RdsMysqlInfo,omitempty" xml:"RdsMysqlInfo,omitempty" type:"Struct"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about the source Simple Log Service (SLS) instance or cluster.
	//
	// example:
	//
	// -
	SlsInfo *CreateApsDatasoureRequestSlsInfo `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
}

func (s CreateApsDatasoureRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApsDatasoureRequest) GoString() string {
	return s.String()
}

func (s *CreateApsDatasoureRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApsDatasoureRequest) GetDatabricksInfo() *CreateApsDatasoureRequestDatabricksInfo {
	return s.DatabricksInfo
}

func (s *CreateApsDatasoureRequest) GetDatasourceDescription() *string {
	return s.DatasourceDescription
}

func (s *CreateApsDatasoureRequest) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *CreateApsDatasoureRequest) GetDatasourceType() *string {
	return s.DatasourceType
}

func (s *CreateApsDatasoureRequest) GetHiveInfo() *CreateApsDatasoureRequestHiveInfo {
	return s.HiveInfo
}

func (s *CreateApsDatasoureRequest) GetKafkaInfo() *CreateApsDatasoureRequestKafkaInfo {
	return s.KafkaInfo
}

func (s *CreateApsDatasoureRequest) GetMode() *string {
	return s.Mode
}

func (s *CreateApsDatasoureRequest) GetPolarDBMysqlInfo() *CreateApsDatasoureRequestPolarDBMysqlInfo {
	return s.PolarDBMysqlInfo
}

func (s *CreateApsDatasoureRequest) GetPolarDBXInfo() *CreateApsDatasoureRequestPolarDBXInfo {
	return s.PolarDBXInfo
}

func (s *CreateApsDatasoureRequest) GetRdsMysqlInfo() *CreateApsDatasoureRequestRdsMysqlInfo {
	return s.RdsMysqlInfo
}

func (s *CreateApsDatasoureRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApsDatasoureRequest) GetSlsInfo() *CreateApsDatasoureRequestSlsInfo {
	return s.SlsInfo
}

func (s *CreateApsDatasoureRequest) SetDBClusterId(v string) *CreateApsDatasoureRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateApsDatasoureRequest) SetDatabricksInfo(v *CreateApsDatasoureRequestDatabricksInfo) *CreateApsDatasoureRequest {
	s.DatabricksInfo = v
	return s
}

func (s *CreateApsDatasoureRequest) SetDatasourceDescription(v string) *CreateApsDatasoureRequest {
	s.DatasourceDescription = &v
	return s
}

func (s *CreateApsDatasoureRequest) SetDatasourceName(v string) *CreateApsDatasoureRequest {
	s.DatasourceName = &v
	return s
}

func (s *CreateApsDatasoureRequest) SetDatasourceType(v string) *CreateApsDatasoureRequest {
	s.DatasourceType = &v
	return s
}

func (s *CreateApsDatasoureRequest) SetHiveInfo(v *CreateApsDatasoureRequestHiveInfo) *CreateApsDatasoureRequest {
	s.HiveInfo = v
	return s
}

func (s *CreateApsDatasoureRequest) SetKafkaInfo(v *CreateApsDatasoureRequestKafkaInfo) *CreateApsDatasoureRequest {
	s.KafkaInfo = v
	return s
}

func (s *CreateApsDatasoureRequest) SetMode(v string) *CreateApsDatasoureRequest {
	s.Mode = &v
	return s
}

func (s *CreateApsDatasoureRequest) SetPolarDBMysqlInfo(v *CreateApsDatasoureRequestPolarDBMysqlInfo) *CreateApsDatasoureRequest {
	s.PolarDBMysqlInfo = v
	return s
}

func (s *CreateApsDatasoureRequest) SetPolarDBXInfo(v *CreateApsDatasoureRequestPolarDBXInfo) *CreateApsDatasoureRequest {
	s.PolarDBXInfo = v
	return s
}

func (s *CreateApsDatasoureRequest) SetRdsMysqlInfo(v *CreateApsDatasoureRequestRdsMysqlInfo) *CreateApsDatasoureRequest {
	s.RdsMysqlInfo = v
	return s
}

func (s *CreateApsDatasoureRequest) SetRegionId(v string) *CreateApsDatasoureRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApsDatasoureRequest) SetSlsInfo(v *CreateApsDatasoureRequestSlsInfo) *CreateApsDatasoureRequest {
	s.SlsInfo = v
	return s
}

func (s *CreateApsDatasoureRequest) Validate() error {
	return dara.Validate(s)
}

type CreateApsDatasoureRequestDatabricksInfo struct {
	// The token that is used to access Databricks.
	//
	// example:
	//
	// ******
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The URL of the workspace.
	//
	// example:
	//
	// xxxxx
	WorkspaceURL *string `json:"WorkspaceURL,omitempty" xml:"WorkspaceURL,omitempty"`
}

func (s CreateApsDatasoureRequestDatabricksInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateApsDatasoureRequestDatabricksInfo) GoString() string {
	return s.String()
}

func (s *CreateApsDatasoureRequestDatabricksInfo) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateApsDatasoureRequestDatabricksInfo) GetWorkspaceURL() *string {
	return s.WorkspaceURL
}

func (s *CreateApsDatasoureRequestDatabricksInfo) SetAccessToken(v string) *CreateApsDatasoureRequestDatabricksInfo {
	s.AccessToken = &v
	return s
}

func (s *CreateApsDatasoureRequestDatabricksInfo) SetWorkspaceURL(v string) *CreateApsDatasoureRequestDatabricksInfo {
	s.WorkspaceURL = &v
	return s
}

func (s *CreateApsDatasoureRequestDatabricksInfo) Validate() error {
	return dara.Validate(s)
}

type CreateApsDatasoureRequestHiveInfo struct {
	// The cluster ID.
	//
	// example:
	//
	// ******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The configuration of the host.
	//
	// example:
	//
	// ******
	HostConfig *string `json:"HostConfig,omitempty" xml:"HostConfig,omitempty"`
	// The URL of the Hive Metastore.
	//
	// example:
	//
	// ******
	MetaStoreUri *string `json:"MetaStoreUri,omitempty" xml:"MetaStoreUri,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-uf*******h
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1*****k
	Vswitch *string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty"`
}

func (s CreateApsDatasoureRequestHiveInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateApsDatasoureRequestHiveInfo) GoString() string {
	return s.String()
}

func (s *CreateApsDatasoureRequestHiveInfo) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateApsDatasoureRequestHiveInfo) GetHostConfig() *string {
	return s.HostConfig
}

func (s *CreateApsDatasoureRequestHiveInfo) GetMetaStoreUri() *string {
	return s.MetaStoreUri
}

func (s *CreateApsDatasoureRequestHiveInfo) GetSecurityGroup() *string {
	return s.SecurityGroup
}

func (s *CreateApsDatasoureRequestHiveInfo) GetVswitch() *string {
	return s.Vswitch
}

func (s *CreateApsDatasoureRequestHiveInfo) SetClusterId(v string) *CreateApsDatasoureRequestHiveInfo {
	s.ClusterId = &v
	return s
}

func (s *CreateApsDatasoureRequestHiveInfo) SetHostConfig(v string) *CreateApsDatasoureRequestHiveInfo {
	s.HostConfig = &v
	return s
}

func (s *CreateApsDatasoureRequestHiveInfo) SetMetaStoreUri(v string) *CreateApsDatasoureRequestHiveInfo {
	s.MetaStoreUri = &v
	return s
}

func (s *CreateApsDatasoureRequestHiveInfo) SetSecurityGroup(v string) *CreateApsDatasoureRequestHiveInfo {
	s.SecurityGroup = &v
	return s
}

func (s *CreateApsDatasoureRequestHiveInfo) SetVswitch(v string) *CreateApsDatasoureRequestHiveInfo {
	s.Vswitch = &v
	return s
}

func (s *CreateApsDatasoureRequestHiveInfo) Validate() error {
	return dara.Validate(s)
}

type CreateApsDatasoureRequestKafkaInfo struct {
	// The ID of the Apache Kafka instance.
	//
	// example:
	//
	// ******
	KafkaClusterId *string `json:"KafkaClusterId,omitempty" xml:"KafkaClusterId,omitempty"`
	// The topic of the Apache Kafka instance.
	//
	// example:
	//
	// test
	KafkaTopic *string `json:"KafkaTopic,omitempty" xml:"KafkaTopic,omitempty"`
}

func (s CreateApsDatasoureRequestKafkaInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateApsDatasoureRequestKafkaInfo) GoString() string {
	return s.String()
}

func (s *CreateApsDatasoureRequestKafkaInfo) GetKafkaClusterId() *string {
	return s.KafkaClusterId
}

func (s *CreateApsDatasoureRequestKafkaInfo) GetKafkaTopic() *string {
	return s.KafkaTopic
}

func (s *CreateApsDatasoureRequestKafkaInfo) SetKafkaClusterId(v string) *CreateApsDatasoureRequestKafkaInfo {
	s.KafkaClusterId = &v
	return s
}

func (s *CreateApsDatasoureRequestKafkaInfo) SetKafkaTopic(v string) *CreateApsDatasoureRequestKafkaInfo {
	s.KafkaTopic = &v
	return s
}

func (s *CreateApsDatasoureRequestKafkaInfo) Validate() error {
	return dara.Validate(s)
}

type CreateApsDatasoureRequestPolarDBMysqlInfo struct {
	// Specifies whether the data source is a cross-account resource. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Across *bool `json:"Across,omitempty" xml:"Across,omitempty"`
	// The name of the cross-account role.
	//
	// example:
	//
	// test-role
	AcrossRole *string `json:"AcrossRole,omitempty" xml:"AcrossRole,omitempty"`
	// The cross-account UID.
	//
	// example:
	//
	// 123456789*
	AcrossUid *string `json:"AcrossUid,omitempty" xml:"AcrossUid,omitempty"`
	// The URL used to connect to the custom ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// ****
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pc-bp*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The password.
	//
	// example:
	//
	// ***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-******
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	// The username used to access the instance.
	//
	// example:
	//
	// test-user-name
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateApsDatasoureRequestPolarDBMysqlInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateApsDatasoureRequestPolarDBMysqlInfo) GoString() string {
	return s.String()
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) GetAcross() *bool {
	return s.Across
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) GetAcrossRole() *string {
	return s.AcrossRole
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) GetAcrossUid() *string {
	return s.AcrossUid
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) GetConnectUrl() *string {
	return s.ConnectUrl
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) GetPassword() *string {
	return s.Password
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) GetSecurityGroup() *string {
	return s.SecurityGroup
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) GetUserName() *string {
	return s.UserName
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) SetAcross(v bool) *CreateApsDatasoureRequestPolarDBMysqlInfo {
	s.Across = &v
	return s
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) SetAcrossRole(v string) *CreateApsDatasoureRequestPolarDBMysqlInfo {
	s.AcrossRole = &v
	return s
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) SetAcrossUid(v string) *CreateApsDatasoureRequestPolarDBMysqlInfo {
	s.AcrossUid = &v
	return s
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) SetConnectUrl(v string) *CreateApsDatasoureRequestPolarDBMysqlInfo {
	s.ConnectUrl = &v
	return s
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) SetInstanceId(v string) *CreateApsDatasoureRequestPolarDBMysqlInfo {
	s.InstanceId = &v
	return s
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) SetPassword(v string) *CreateApsDatasoureRequestPolarDBMysqlInfo {
	s.Password = &v
	return s
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) SetRegionId(v string) *CreateApsDatasoureRequestPolarDBMysqlInfo {
	s.RegionId = &v
	return s
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) SetSecurityGroup(v string) *CreateApsDatasoureRequestPolarDBMysqlInfo {
	s.SecurityGroup = &v
	return s
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) SetUserName(v string) *CreateApsDatasoureRequestPolarDBMysqlInfo {
	s.UserName = &v
	return s
}

func (s *CreateApsDatasoureRequestPolarDBMysqlInfo) Validate() error {
	return dara.Validate(s)
}

type CreateApsDatasoureRequestPolarDBXInfo struct {
	// The instance ID.
	//
	// example:
	//
	// -
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateApsDatasoureRequestPolarDBXInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateApsDatasoureRequestPolarDBXInfo) GoString() string {
	return s.String()
}

func (s *CreateApsDatasoureRequestPolarDBXInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateApsDatasoureRequestPolarDBXInfo) SetInstanceId(v string) *CreateApsDatasoureRequestPolarDBXInfo {
	s.InstanceId = &v
	return s
}

func (s *CreateApsDatasoureRequestPolarDBXInfo) Validate() error {
	return dara.Validate(s)
}

type CreateApsDatasoureRequestRdsMysqlInfo struct {
	// The URL used to connect to the read-only instance.
	//
	// example:
	//
	// ******
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-xxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The password of the database account of the instance.
	//
	// example:
	//
	// ******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-******
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	// The name of the database account of the instance.
	//
	// example:
	//
	// user
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateApsDatasoureRequestRdsMysqlInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateApsDatasoureRequestRdsMysqlInfo) GoString() string {
	return s.String()
}

func (s *CreateApsDatasoureRequestRdsMysqlInfo) GetConnectUrl() *string {
	return s.ConnectUrl
}

func (s *CreateApsDatasoureRequestRdsMysqlInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateApsDatasoureRequestRdsMysqlInfo) GetPassword() *string {
	return s.Password
}

func (s *CreateApsDatasoureRequestRdsMysqlInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApsDatasoureRequestRdsMysqlInfo) GetSecurityGroup() *string {
	return s.SecurityGroup
}

func (s *CreateApsDatasoureRequestRdsMysqlInfo) GetUserName() *string {
	return s.UserName
}

func (s *CreateApsDatasoureRequestRdsMysqlInfo) SetConnectUrl(v string) *CreateApsDatasoureRequestRdsMysqlInfo {
	s.ConnectUrl = &v
	return s
}

func (s *CreateApsDatasoureRequestRdsMysqlInfo) SetInstanceId(v string) *CreateApsDatasoureRequestRdsMysqlInfo {
	s.InstanceId = &v
	return s
}

func (s *CreateApsDatasoureRequestRdsMysqlInfo) SetPassword(v string) *CreateApsDatasoureRequestRdsMysqlInfo {
	s.Password = &v
	return s
}

func (s *CreateApsDatasoureRequestRdsMysqlInfo) SetRegionId(v string) *CreateApsDatasoureRequestRdsMysqlInfo {
	s.RegionId = &v
	return s
}

func (s *CreateApsDatasoureRequestRdsMysqlInfo) SetSecurityGroup(v string) *CreateApsDatasoureRequestRdsMysqlInfo {
	s.SecurityGroup = &v
	return s
}

func (s *CreateApsDatasoureRequestRdsMysqlInfo) SetUserName(v string) *CreateApsDatasoureRequestRdsMysqlInfo {
	s.UserName = &v
	return s
}

func (s *CreateApsDatasoureRequestRdsMysqlInfo) Validate() error {
	return dara.Validate(s)
}

type CreateApsDatasoureRequestSlsInfo struct {
	// Specifies whether the data source is a cross-account resource.
	//
	// example:
	//
	// false
	Across *bool `json:"Across,omitempty" xml:"Across,omitempty"`
	// The name of the cross-account role.
	//
	// example:
	//
	// yyy
	AcrossRole *string `json:"AcrossRole,omitempty" xml:"AcrossRole,omitempty"`
	// The cross-account UID.
	//
	// example:
	//
	// xxxx
	AcrossUid *string `json:"AcrossUid,omitempty" xml:"AcrossUid,omitempty"`
	// The SLS project.
	//
	// example:
	//
	// test-project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// The name of the SLS Logstore.
	//
	// example:
	//
	// test-store
	Store *string `json:"Store,omitempty" xml:"Store,omitempty"`
}

func (s CreateApsDatasoureRequestSlsInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateApsDatasoureRequestSlsInfo) GoString() string {
	return s.String()
}

func (s *CreateApsDatasoureRequestSlsInfo) GetAcross() *bool {
	return s.Across
}

func (s *CreateApsDatasoureRequestSlsInfo) GetAcrossRole() *string {
	return s.AcrossRole
}

func (s *CreateApsDatasoureRequestSlsInfo) GetAcrossUid() *string {
	return s.AcrossUid
}

func (s *CreateApsDatasoureRequestSlsInfo) GetProject() *string {
	return s.Project
}

func (s *CreateApsDatasoureRequestSlsInfo) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *CreateApsDatasoureRequestSlsInfo) GetStore() *string {
	return s.Store
}

func (s *CreateApsDatasoureRequestSlsInfo) SetAcross(v bool) *CreateApsDatasoureRequestSlsInfo {
	s.Across = &v
	return s
}

func (s *CreateApsDatasoureRequestSlsInfo) SetAcrossRole(v string) *CreateApsDatasoureRequestSlsInfo {
	s.AcrossRole = &v
	return s
}

func (s *CreateApsDatasoureRequestSlsInfo) SetAcrossUid(v string) *CreateApsDatasoureRequestSlsInfo {
	s.AcrossUid = &v
	return s
}

func (s *CreateApsDatasoureRequestSlsInfo) SetProject(v string) *CreateApsDatasoureRequestSlsInfo {
	s.Project = &v
	return s
}

func (s *CreateApsDatasoureRequestSlsInfo) SetSourceRegionId(v string) *CreateApsDatasoureRequestSlsInfo {
	s.SourceRegionId = &v
	return s
}

func (s *CreateApsDatasoureRequestSlsInfo) SetStore(v string) *CreateApsDatasoureRequestSlsInfo {
	s.Store = &v
	return s
}

func (s *CreateApsDatasoureRequestSlsInfo) Validate() error {
	return dara.Validate(s)
}
