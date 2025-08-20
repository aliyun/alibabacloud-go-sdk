// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsDatasourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApsDatasource(v *DescribeApsDatasourceResponseBodyApsDatasource) *DescribeApsDatasourceResponseBody
	GetApsDatasource() *DescribeApsDatasourceResponseBodyApsDatasource
	SetRequestId(v string) *DescribeApsDatasourceResponseBody
	GetRequestId() *string
}

type DescribeApsDatasourceResponseBody struct {
	// The queried APS data source.
	ApsDatasource *DescribeApsDatasourceResponseBodyApsDatasource `json:"ApsDatasource,omitempty" xml:"ApsDatasource,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******-**D8-5***-A***-****587
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApsDatasourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApsDatasourceResponseBody) GetApsDatasource() *DescribeApsDatasourceResponseBodyApsDatasource {
	return s.ApsDatasource
}

func (s *DescribeApsDatasourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApsDatasourceResponseBody) SetApsDatasource(v *DescribeApsDatasourceResponseBodyApsDatasource) *DescribeApsDatasourceResponseBody {
	s.ApsDatasource = v
	return s
}

func (s *DescribeApsDatasourceResponseBody) SetRequestId(v string) *DescribeApsDatasourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApsDatasourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApsDatasourceResponseBodyApsDatasource struct {
	// The time when the data source was created.
	//
	// example:
	//
	// 2024-04-12T15:03:38Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The information about Databricks.
	//
	// example:
	//
	// -
	DatabricksInfo *DescribeApsDatasourceResponseBodyApsDatasourceDatabricksInfo `json:"DatabricksInfo,omitempty" xml:"DatabricksInfo,omitempty" type:"Struct"`
	// The description of the data source.
	//
	// example:
	//
	// test
	DatasourceDescription *string `json:"DatasourceDescription,omitempty" xml:"DatasourceDescription,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// kafka-2024***
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// The type of the data source.
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
	HiveInfo *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo `json:"HiveInfo,omitempty" xml:"HiveInfo,omitempty" type:"Struct"`
	// The information about the Kafka instance.
	//
	// example:
	//
	// -
	KafkaInfo *DescribeApsDatasourceResponseBodyApsDatasourceKafkaInfo `json:"KafkaInfo,omitempty" xml:"KafkaInfo,omitempty" type:"Struct"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	PolarDBMysqlInfo *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo `json:"PolarDBMysqlInfo,omitempty" xml:"PolarDBMysqlInfo,omitempty" type:"Struct"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	RdsMysqlInfo *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo `json:"RdsMysqlInfo,omitempty" xml:"RdsMysqlInfo,omitempty" type:"Struct"`
	// The Simple Log Service (SLS) project.
	//
	// example:
	//
	// -
	SlsInfo *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
}

func (s DescribeApsDatasourceResponseBodyApsDatasource) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsDatasourceResponseBodyApsDatasource) GoString() string {
	return s.String()
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) GetDatabricksInfo() *DescribeApsDatasourceResponseBodyApsDatasourceDatabricksInfo {
	return s.DatabricksInfo
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) GetDatasourceDescription() *string {
	return s.DatasourceDescription
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) GetDatasourceType() *string {
	return s.DatasourceType
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) GetHiveInfo() *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo {
	return s.HiveInfo
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) GetKafkaInfo() *DescribeApsDatasourceResponseBodyApsDatasourceKafkaInfo {
	return s.KafkaInfo
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) GetPolarDBMysqlInfo() *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo {
	return s.PolarDBMysqlInfo
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) GetRdsMysqlInfo() *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo {
	return s.RdsMysqlInfo
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) GetSlsInfo() *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo {
	return s.SlsInfo
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) SetCreateTime(v string) *DescribeApsDatasourceResponseBodyApsDatasource {
	s.CreateTime = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) SetDBClusterId(v string) *DescribeApsDatasourceResponseBodyApsDatasource {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) SetDatabricksInfo(v *DescribeApsDatasourceResponseBodyApsDatasourceDatabricksInfo) *DescribeApsDatasourceResponseBodyApsDatasource {
	s.DatabricksInfo = v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) SetDatasourceDescription(v string) *DescribeApsDatasourceResponseBodyApsDatasource {
	s.DatasourceDescription = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) SetDatasourceName(v string) *DescribeApsDatasourceResponseBodyApsDatasource {
	s.DatasourceName = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) SetDatasourceType(v string) *DescribeApsDatasourceResponseBodyApsDatasource {
	s.DatasourceType = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) SetHiveInfo(v *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo) *DescribeApsDatasourceResponseBodyApsDatasource {
	s.HiveInfo = v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) SetKafkaInfo(v *DescribeApsDatasourceResponseBodyApsDatasourceKafkaInfo) *DescribeApsDatasourceResponseBodyApsDatasource {
	s.KafkaInfo = v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) SetPolarDBMysqlInfo(v *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) *DescribeApsDatasourceResponseBodyApsDatasource {
	s.PolarDBMysqlInfo = v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) SetRdsMysqlInfo(v *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo) *DescribeApsDatasourceResponseBodyApsDatasource {
	s.RdsMysqlInfo = v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) SetSlsInfo(v *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) *DescribeApsDatasourceResponseBodyApsDatasource {
	s.SlsInfo = v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasource) Validate() error {
	return dara.Validate(s)
}

type DescribeApsDatasourceResponseBodyApsDatasourceDatabricksInfo struct {
	// The token that is used to access Databricks.
	//
	// example:
	//
	// ******
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// The URL of the workspace.
	//
	// example:
	//
	// -
	WorkspaceURL *string `json:"workspaceURL,omitempty" xml:"workspaceURL,omitempty"`
}

func (s DescribeApsDatasourceResponseBodyApsDatasourceDatabricksInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsDatasourceResponseBodyApsDatasourceDatabricksInfo) GoString() string {
	return s.String()
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceDatabricksInfo) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceDatabricksInfo) GetWorkspaceURL() *string {
	return s.WorkspaceURL
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceDatabricksInfo) SetAccessToken(v string) *DescribeApsDatasourceResponseBodyApsDatasourceDatabricksInfo {
	s.AccessToken = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceDatabricksInfo) SetWorkspaceURL(v string) *DescribeApsDatasourceResponseBodyApsDatasourceDatabricksInfo {
	s.WorkspaceURL = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceDatabricksInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo struct {
	// The ID of the E-MapReduce (EMR) cluster.
	//
	// example:
	//
	// -
	EmrClusterId *string `json:"EmrClusterId,omitempty" xml:"EmrClusterId,omitempty"`
	// The URL of the Hive Metastore.
	//
	// example:
	//
	// -
	MetaStoreUri *string `json:"MetaStoreUri,omitempty" xml:"MetaStoreUri,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-******
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-******
	Vswitch *string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty"`
}

func (s DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo) GoString() string {
	return s.String()
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo) GetEmrClusterId() *string {
	return s.EmrClusterId
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo) GetMetaStoreUri() *string {
	return s.MetaStoreUri
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo) GetSecurityGroup() *string {
	return s.SecurityGroup
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo) GetVswitch() *string {
	return s.Vswitch
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo) SetEmrClusterId(v string) *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo {
	s.EmrClusterId = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo) SetMetaStoreUri(v string) *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo {
	s.MetaStoreUri = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo) SetSecurityGroup(v string) *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo {
	s.SecurityGroup = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo) SetVswitch(v string) *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo {
	s.Vswitch = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceHiveInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeApsDatasourceResponseBodyApsDatasourceKafkaInfo struct {
	// The ID of the Kafka instance.
	//
	// example:
	//
	// -
	KafkaClusterId *string `json:"KafkaClusterId,omitempty" xml:"KafkaClusterId,omitempty"`
	// The topic of the Kafka instance.
	//
	// example:
	//
	// [{\\"value\\": \\"hongxian_test\\"}]
	KafkaTopic *string `json:"KafkaTopic,omitempty" xml:"KafkaTopic,omitempty"`
}

func (s DescribeApsDatasourceResponseBodyApsDatasourceKafkaInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsDatasourceResponseBodyApsDatasourceKafkaInfo) GoString() string {
	return s.String()
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceKafkaInfo) GetKafkaClusterId() *string {
	return s.KafkaClusterId
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceKafkaInfo) GetKafkaTopic() *string {
	return s.KafkaTopic
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceKafkaInfo) SetKafkaClusterId(v string) *DescribeApsDatasourceResponseBodyApsDatasourceKafkaInfo {
	s.KafkaClusterId = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceKafkaInfo) SetKafkaTopic(v string) *DescribeApsDatasourceResponseBodyApsDatasourceKafkaInfo {
	s.KafkaTopic = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceKafkaInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo struct {
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	Across *bool `json:"Across,omitempty" xml:"Across,omitempty"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	AcrossRole *string `json:"AcrossRole,omitempty" xml:"AcrossRole,omitempty"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	AcrossUid *string `json:"AcrossUid,omitempty" xml:"AcrossUid,omitempty"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) GoString() string {
	return s.String()
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) GetAcross() *bool {
	return s.Across
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) GetAcrossRole() *string {
	return s.AcrossRole
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) GetAcrossUid() *string {
	return s.AcrossUid
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) GetConnectUrl() *string {
	return s.ConnectUrl
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) GetSecurityGroup() *string {
	return s.SecurityGroup
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) GetUserName() *string {
	return s.UserName
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) SetAcross(v bool) *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo {
	s.Across = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) SetAcrossRole(v string) *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo {
	s.AcrossRole = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) SetAcrossUid(v string) *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo {
	s.AcrossUid = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) SetConnectUrl(v string) *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) SetInstanceId(v string) *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) SetRegionId(v string) *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) SetSecurityGroup(v string) *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo {
	s.SecurityGroup = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) SetUserName(v string) *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo {
	s.UserName = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourcePolarDBMysqlInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo struct {
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo) GoString() string {
	return s.String()
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo) GetConnectUrl() *string {
	return s.ConnectUrl
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo) GetSecurityGroup() *string {
	return s.SecurityGroup
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo) GetUserName() *string {
	return s.UserName
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo) SetConnectUrl(v string) *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo) SetInstanceId(v string) *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo) SetRegionId(v string) *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo) SetSecurityGroup(v string) *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo {
	s.SecurityGroup = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo) SetUserName(v string) *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo {
	s.UserName = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceRdsMysqlInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo struct {
	// Indicates whether the data source is a cross-account resource. Valid values:
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
	// test
	AcrossRole *string `json:"AcrossRole,omitempty" xml:"AcrossRole,omitempty"`
	// The cross-account UID.
	//
	// example:
	//
	// 123456
	AcrossUid *string `json:"AcrossUid,omitempty" xml:"AcrossUid,omitempty"`
	// The name of the SLS project.
	//
	// example:
	//
	// ***
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// The name of the SLS Logstore.
	//
	// example:
	//
	// ***
	Store *string `json:"Store,omitempty" xml:"Store,omitempty"`
}

func (s DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) GoString() string {
	return s.String()
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) GetAcross() *bool {
	return s.Across
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) GetAcrossRole() *string {
	return s.AcrossRole
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) GetAcrossUid() *string {
	return s.AcrossUid
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) GetProject() *string {
	return s.Project
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) GetStore() *string {
	return s.Store
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) SetAcross(v bool) *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo {
	s.Across = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) SetAcrossRole(v string) *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo {
	s.AcrossRole = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) SetAcrossUid(v string) *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo {
	s.AcrossUid = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) SetProject(v string) *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo {
	s.Project = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) SetSourceRegionId(v string) *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo {
	s.SourceRegionId = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) SetStore(v string) *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo {
	s.Store = &v
	return s
}

func (s *DescribeApsDatasourceResponseBodyApsDatasourceSlsInfo) Validate() error {
	return dara.Validate(s)
}
