// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApsDatasoureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyApsDatasoureRequest
	GetDBClusterId() *string
	SetDatasourceDescription(v string) *ModifyApsDatasoureRequest
	GetDatasourceDescription() *string
	SetDatasourceId(v int64) *ModifyApsDatasoureRequest
	GetDatasourceId() *int64
	SetDatasourceName(v string) *ModifyApsDatasoureRequest
	GetDatasourceName() *string
	SetKafkaInfo(v *ModifyApsDatasoureRequestKafkaInfo) *ModifyApsDatasoureRequest
	GetKafkaInfo() *ModifyApsDatasoureRequestKafkaInfo
	SetLakehouseId(v *ModifyApsDatasoureRequestLakehouseId) *ModifyApsDatasoureRequest
	GetLakehouseId() *ModifyApsDatasoureRequestLakehouseId
	SetPolarDBMysqlInfo(v *ModifyApsDatasoureRequestPolarDBMysqlInfo) *ModifyApsDatasoureRequest
	GetPolarDBMysqlInfo() *ModifyApsDatasoureRequestPolarDBMysqlInfo
	SetRdsMysqlInfo(v *ModifyApsDatasoureRequestRdsMysqlInfo) *ModifyApsDatasoureRequest
	GetRdsMysqlInfo() *ModifyApsDatasoureRequestRdsMysqlInfo
	SetRegionId(v string) *ModifyApsDatasoureRequest
	GetRegionId() *string
	SetSlsInfo(v *ModifyApsDatasoureRequestSlsInfo) *ModifyApsDatasoureRequest
	GetSlsInfo() *ModifyApsDatasoureRequestSlsInfo
}

type ModifyApsDatasoureRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// test
	DatasourceDescription *string `json:"DatasourceDescription,omitempty" xml:"DatasourceDescription,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// sls-******
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// The information about the Kafka instance.
	//
	// example:
	//
	// -
	KafkaInfo *ModifyApsDatasoureRequestKafkaInfo `json:"KafkaInfo,omitempty" xml:"KafkaInfo,omitempty" type:"Struct"`
	// The lakehouse ID.
	//
	// example:
	//
	// 123
	LakehouseId *ModifyApsDatasoureRequestLakehouseId `json:"LakehouseId,omitempty" xml:"LakehouseId,omitempty" type:"Struct"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	PolarDBMysqlInfo *ModifyApsDatasoureRequestPolarDBMysqlInfo `json:"PolarDBMysqlInfo,omitempty" xml:"PolarDBMysqlInfo,omitempty" type:"Struct"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	RdsMysqlInfo *ModifyApsDatasoureRequestRdsMysqlInfo `json:"RdsMysqlInfo,omitempty" xml:"RdsMysqlInfo,omitempty" type:"Struct"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about Simple Log Service (SLS).
	//
	// example:
	//
	// -
	SlsInfo *ModifyApsDatasoureRequestSlsInfo `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
}

func (s ModifyApsDatasoureRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsDatasoureRequest) GoString() string {
	return s.String()
}

func (s *ModifyApsDatasoureRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyApsDatasoureRequest) GetDatasourceDescription() *string {
	return s.DatasourceDescription
}

func (s *ModifyApsDatasoureRequest) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *ModifyApsDatasoureRequest) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *ModifyApsDatasoureRequest) GetKafkaInfo() *ModifyApsDatasoureRequestKafkaInfo {
	return s.KafkaInfo
}

func (s *ModifyApsDatasoureRequest) GetLakehouseId() *ModifyApsDatasoureRequestLakehouseId {
	return s.LakehouseId
}

func (s *ModifyApsDatasoureRequest) GetPolarDBMysqlInfo() *ModifyApsDatasoureRequestPolarDBMysqlInfo {
	return s.PolarDBMysqlInfo
}

func (s *ModifyApsDatasoureRequest) GetRdsMysqlInfo() *ModifyApsDatasoureRequestRdsMysqlInfo {
	return s.RdsMysqlInfo
}

func (s *ModifyApsDatasoureRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApsDatasoureRequest) GetSlsInfo() *ModifyApsDatasoureRequestSlsInfo {
	return s.SlsInfo
}

func (s *ModifyApsDatasoureRequest) SetDBClusterId(v string) *ModifyApsDatasoureRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyApsDatasoureRequest) SetDatasourceDescription(v string) *ModifyApsDatasoureRequest {
	s.DatasourceDescription = &v
	return s
}

func (s *ModifyApsDatasoureRequest) SetDatasourceId(v int64) *ModifyApsDatasoureRequest {
	s.DatasourceId = &v
	return s
}

func (s *ModifyApsDatasoureRequest) SetDatasourceName(v string) *ModifyApsDatasoureRequest {
	s.DatasourceName = &v
	return s
}

func (s *ModifyApsDatasoureRequest) SetKafkaInfo(v *ModifyApsDatasoureRequestKafkaInfo) *ModifyApsDatasoureRequest {
	s.KafkaInfo = v
	return s
}

func (s *ModifyApsDatasoureRequest) SetLakehouseId(v *ModifyApsDatasoureRequestLakehouseId) *ModifyApsDatasoureRequest {
	s.LakehouseId = v
	return s
}

func (s *ModifyApsDatasoureRequest) SetPolarDBMysqlInfo(v *ModifyApsDatasoureRequestPolarDBMysqlInfo) *ModifyApsDatasoureRequest {
	s.PolarDBMysqlInfo = v
	return s
}

func (s *ModifyApsDatasoureRequest) SetRdsMysqlInfo(v *ModifyApsDatasoureRequestRdsMysqlInfo) *ModifyApsDatasoureRequest {
	s.RdsMysqlInfo = v
	return s
}

func (s *ModifyApsDatasoureRequest) SetRegionId(v string) *ModifyApsDatasoureRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyApsDatasoureRequest) SetSlsInfo(v *ModifyApsDatasoureRequestSlsInfo) *ModifyApsDatasoureRequest {
	s.SlsInfo = v
	return s
}

func (s *ModifyApsDatasoureRequest) Validate() error {
	if s.KafkaInfo != nil {
		if err := s.KafkaInfo.Validate(); err != nil {
			return err
		}
	}
	if s.LakehouseId != nil {
		if err := s.LakehouseId.Validate(); err != nil {
			return err
		}
	}
	if s.PolarDBMysqlInfo != nil {
		if err := s.PolarDBMysqlInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RdsMysqlInfo != nil {
		if err := s.RdsMysqlInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SlsInfo != nil {
		if err := s.SlsInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyApsDatasoureRequestKafkaInfo struct {
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
	// test-topic
	KafkaTopic *string `json:"KafkaTopic,omitempty" xml:"KafkaTopic,omitempty"`
}

func (s ModifyApsDatasoureRequestKafkaInfo) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsDatasoureRequestKafkaInfo) GoString() string {
	return s.String()
}

func (s *ModifyApsDatasoureRequestKafkaInfo) GetKafkaClusterId() *string {
	return s.KafkaClusterId
}

func (s *ModifyApsDatasoureRequestKafkaInfo) GetKafkaTopic() *string {
	return s.KafkaTopic
}

func (s *ModifyApsDatasoureRequestKafkaInfo) SetKafkaClusterId(v string) *ModifyApsDatasoureRequestKafkaInfo {
	s.KafkaClusterId = &v
	return s
}

func (s *ModifyApsDatasoureRequestKafkaInfo) SetKafkaTopic(v string) *ModifyApsDatasoureRequestKafkaInfo {
	s.KafkaTopic = &v
	return s
}

func (s *ModifyApsDatasoureRequestKafkaInfo) Validate() error {
	return dara.Validate(s)
}

type ModifyApsDatasoureRequestLakehouseId struct {
	// The name of the security group.
	//
	// example:
	//
	// sg-******
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the vSwitch.
	//
	// example:
	//
	// vsw-******
	Vswitch *string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty"`
}

func (s ModifyApsDatasoureRequestLakehouseId) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsDatasoureRequestLakehouseId) GoString() string {
	return s.String()
}

func (s *ModifyApsDatasoureRequestLakehouseId) GetSecurityGroup() *string {
	return s.SecurityGroup
}

func (s *ModifyApsDatasoureRequestLakehouseId) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyApsDatasoureRequestLakehouseId) GetVswitch() *string {
	return s.Vswitch
}

func (s *ModifyApsDatasoureRequestLakehouseId) SetSecurityGroup(v string) *ModifyApsDatasoureRequestLakehouseId {
	s.SecurityGroup = &v
	return s
}

func (s *ModifyApsDatasoureRequestLakehouseId) SetVpcId(v string) *ModifyApsDatasoureRequestLakehouseId {
	s.VpcId = &v
	return s
}

func (s *ModifyApsDatasoureRequestLakehouseId) SetVswitch(v string) *ModifyApsDatasoureRequestLakehouseId {
	s.Vswitch = &v
	return s
}

func (s *ModifyApsDatasoureRequestLakehouseId) Validate() error {
	return dara.Validate(s)
}

type ModifyApsDatasoureRequestPolarDBMysqlInfo struct {
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
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
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
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ModifyApsDatasoureRequestPolarDBMysqlInfo) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsDatasoureRequestPolarDBMysqlInfo) GoString() string {
	return s.String()
}

func (s *ModifyApsDatasoureRequestPolarDBMysqlInfo) GetConnectUrl() *string {
	return s.ConnectUrl
}

func (s *ModifyApsDatasoureRequestPolarDBMysqlInfo) GetPassword() *string {
	return s.Password
}

func (s *ModifyApsDatasoureRequestPolarDBMysqlInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApsDatasoureRequestPolarDBMysqlInfo) GetUserName() *string {
	return s.UserName
}

func (s *ModifyApsDatasoureRequestPolarDBMysqlInfo) SetConnectUrl(v string) *ModifyApsDatasoureRequestPolarDBMysqlInfo {
	s.ConnectUrl = &v
	return s
}

func (s *ModifyApsDatasoureRequestPolarDBMysqlInfo) SetPassword(v string) *ModifyApsDatasoureRequestPolarDBMysqlInfo {
	s.Password = &v
	return s
}

func (s *ModifyApsDatasoureRequestPolarDBMysqlInfo) SetRegionId(v string) *ModifyApsDatasoureRequestPolarDBMysqlInfo {
	s.RegionId = &v
	return s
}

func (s *ModifyApsDatasoureRequestPolarDBMysqlInfo) SetUserName(v string) *ModifyApsDatasoureRequestPolarDBMysqlInfo {
	s.UserName = &v
	return s
}

func (s *ModifyApsDatasoureRequestPolarDBMysqlInfo) Validate() error {
	return dara.Validate(s)
}

type ModifyApsDatasoureRequestRdsMysqlInfo struct {
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
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
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
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ModifyApsDatasoureRequestRdsMysqlInfo) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsDatasoureRequestRdsMysqlInfo) GoString() string {
	return s.String()
}

func (s *ModifyApsDatasoureRequestRdsMysqlInfo) GetConnectUrl() *string {
	return s.ConnectUrl
}

func (s *ModifyApsDatasoureRequestRdsMysqlInfo) GetPassword() *string {
	return s.Password
}

func (s *ModifyApsDatasoureRequestRdsMysqlInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApsDatasoureRequestRdsMysqlInfo) GetUserName() *string {
	return s.UserName
}

func (s *ModifyApsDatasoureRequestRdsMysqlInfo) SetConnectUrl(v string) *ModifyApsDatasoureRequestRdsMysqlInfo {
	s.ConnectUrl = &v
	return s
}

func (s *ModifyApsDatasoureRequestRdsMysqlInfo) SetPassword(v string) *ModifyApsDatasoureRequestRdsMysqlInfo {
	s.Password = &v
	return s
}

func (s *ModifyApsDatasoureRequestRdsMysqlInfo) SetRegionId(v string) *ModifyApsDatasoureRequestRdsMysqlInfo {
	s.RegionId = &v
	return s
}

func (s *ModifyApsDatasoureRequestRdsMysqlInfo) SetUserName(v string) *ModifyApsDatasoureRequestRdsMysqlInfo {
	s.UserName = &v
	return s
}

func (s *ModifyApsDatasoureRequestRdsMysqlInfo) Validate() error {
	return dara.Validate(s)
}

type ModifyApsDatasoureRequestSlsInfo struct {
	// Specifies whether to use a cross-account resource as the data source. Valid values:
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
	// 123456
	AcrossUid *string `json:"AcrossUid,omitempty" xml:"AcrossUid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
}

func (s ModifyApsDatasoureRequestSlsInfo) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsDatasoureRequestSlsInfo) GoString() string {
	return s.String()
}

func (s *ModifyApsDatasoureRequestSlsInfo) GetAcross() *bool {
	return s.Across
}

func (s *ModifyApsDatasoureRequestSlsInfo) GetAcrossRole() *string {
	return s.AcrossRole
}

func (s *ModifyApsDatasoureRequestSlsInfo) GetAcrossUid() *string {
	return s.AcrossUid
}

func (s *ModifyApsDatasoureRequestSlsInfo) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *ModifyApsDatasoureRequestSlsInfo) SetAcross(v bool) *ModifyApsDatasoureRequestSlsInfo {
	s.Across = &v
	return s
}

func (s *ModifyApsDatasoureRequestSlsInfo) SetAcrossRole(v string) *ModifyApsDatasoureRequestSlsInfo {
	s.AcrossRole = &v
	return s
}

func (s *ModifyApsDatasoureRequestSlsInfo) SetAcrossUid(v string) *ModifyApsDatasoureRequestSlsInfo {
	s.AcrossUid = &v
	return s
}

func (s *ModifyApsDatasoureRequestSlsInfo) SetSourceRegionId(v string) *ModifyApsDatasoureRequestSlsInfo {
	s.SourceRegionId = &v
	return s
}

func (s *ModifyApsDatasoureRequestSlsInfo) Validate() error {
	return dara.Validate(s)
}
