// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditStatus(v int32) *CreateDataLimitRequest
	GetAuditStatus() *int32
	SetAutoScan(v int32) *CreateDataLimitRequest
	GetAutoScan() *int32
	SetCertificatePermission(v string) *CreateDataLimitRequest
	GetCertificatePermission() *string
	SetEnable(v int32) *CreateDataLimitRequest
	GetEnable() *int32
	SetEngineType(v string) *CreateDataLimitRequest
	GetEngineType() *string
	SetEventStatus(v int32) *CreateDataLimitRequest
	GetEventStatus() *int32
	SetFeatureType(v int32) *CreateDataLimitRequest
	GetFeatureType() *int32
	SetInstantlyScan(v bool) *CreateDataLimitRequest
	GetInstantlyScan() *bool
	SetLang(v string) *CreateDataLimitRequest
	GetLang() *string
	SetLogStoreDay(v int32) *CreateDataLimitRequest
	GetLogStoreDay() *int32
	SetOcrStatus(v int32) *CreateDataLimitRequest
	GetOcrStatus() *int32
	SetParentId(v string) *CreateDataLimitRequest
	GetParentId() *string
	SetPassword(v string) *CreateDataLimitRequest
	GetPassword() *string
	SetPort(v int32) *CreateDataLimitRequest
	GetPort() *int32
	SetResourceType(v int32) *CreateDataLimitRequest
	GetResourceType() *int32
	SetSamplingSize(v int32) *CreateDataLimitRequest
	GetSamplingSize() *int32
	SetServiceRegionId(v string) *CreateDataLimitRequest
	GetServiceRegionId() *string
	SetSourceIp(v string) *CreateDataLimitRequest
	GetSourceIp() *string
	SetUserName(v string) *CreateDataLimitRequest
	GetUserName() *string
}

type CreateDataLimitRequest struct {
	// Specifies whether to enable the security audit feature. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Specifies whether to automatically trigger a re-scan after a rule is modified. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// > When a re-scan is triggered, DSC scans all data in your data asset.
	//
	// example:
	//
	// 1
	AutoScan *int32 `json:"AutoScan,omitempty" xml:"AutoScan,omitempty"`
	// The permissions. Valid values:
	//
	// 	- **ReadOnly**: read-only permissions
	//
	// 	- **ReadWrite**: read and write permissions
	//
	// example:
	//
	// ReadOnly
	CertificatePermission *string `json:"CertificatePermission,omitempty" xml:"CertificatePermission,omitempty"`
	// Specifies whether to enable sensitive data detection. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// > If this is your first time to authorize DSC to access the data asset, the default value is 1. If this is not your first time to authorize DSC to access the data asset, the default value is the same as that used in the last authorization operation. Both 1 and 0 are possible.
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The database engine that is run by the instance. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// Specifies whether to enable anomalous event detection. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes (default)
	//
	// example:
	//
	// 1
	EventStatus *int32 `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// Specifies whether to immediately scan the authorized asset. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// false
	InstantlyScan *bool `json:"InstantlyScan,omitempty" xml:"InstantlyScan,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The retention period of raw logs after you enable the security audit feature. Unit: days. Valid values:
	//
	// 	- **30**
	//
	// 	- **90**
	//
	// 	- **180**
	//
	// 	- **365**
	//
	// example:
	//
	// 30
	LogStoreDay *int32 `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	// Specifies whether to enable optical character recognition (OCR). Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 0
	OcrStatus *int32 `json:"OcrStatus,omitempty" xml:"OcrStatus,omitempty"`
	// The name of the asset. The value is a connection string. It consists of an instance ID and a database name, which are separated by a comma (,). This parameter is required.
	//
	// example:
	//
	// test-11**
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The password that is used to access the database.
	//
	// example:
	//
	// passwd
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port that is used to connect to the database.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of service to which the data asset belongs. Valid values:
	//
	// 	- **1*	- :MaxCompute
	//
	// 	- **2**: Object Storage Service (OSS)
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4*	- :Tablestore
	//
	// 	- **5**: ApsaraDB RDS
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The number of sensitive data samples that are collected after sensitive data detection is enabled. Valid values:
	//
	// 	- **0**
	//
	// 	- **5**
	//
	// 	- **10**
	//
	// example:
	//
	// 0
	SamplingSize *int32 `json:"SamplingSize,omitempty" xml:"SamplingSize,omitempty"`
	// The region in which the data asset resides. Valid values:
	//
	// 	- **cn-beijing**: China (Beijing).
	//
	// 	- **cn-zhangjiakou**: China (Zhangjiakou)
	//
	// 	- **cn-huhehaote**: China (Hohhot)
	//
	// 	- **cn-hangzhou**: China (Hangzhou)
	//
	// 	- **cn-shanghai**: China (Shanghai)
	//
	// 	- **cn-shenzhen**: China (Shenzhen)
	//
	// 	- **cn-hongkong**: China (Hong Kong)
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The username that is used to access the database.
	//
	// example:
	//
	// yhm
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateDataLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLimitRequest) GoString() string {
	return s.String()
}

func (s *CreateDataLimitRequest) GetAuditStatus() *int32 {
	return s.AuditStatus
}

func (s *CreateDataLimitRequest) GetAutoScan() *int32 {
	return s.AutoScan
}

func (s *CreateDataLimitRequest) GetCertificatePermission() *string {
	return s.CertificatePermission
}

func (s *CreateDataLimitRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *CreateDataLimitRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *CreateDataLimitRequest) GetEventStatus() *int32 {
	return s.EventStatus
}

func (s *CreateDataLimitRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *CreateDataLimitRequest) GetInstantlyScan() *bool {
	return s.InstantlyScan
}

func (s *CreateDataLimitRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateDataLimitRequest) GetLogStoreDay() *int32 {
	return s.LogStoreDay
}

func (s *CreateDataLimitRequest) GetOcrStatus() *int32 {
	return s.OcrStatus
}

func (s *CreateDataLimitRequest) GetParentId() *string {
	return s.ParentId
}

func (s *CreateDataLimitRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateDataLimitRequest) GetPort() *int32 {
	return s.Port
}

func (s *CreateDataLimitRequest) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *CreateDataLimitRequest) GetSamplingSize() *int32 {
	return s.SamplingSize
}

func (s *CreateDataLimitRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *CreateDataLimitRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateDataLimitRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateDataLimitRequest) SetAuditStatus(v int32) *CreateDataLimitRequest {
	s.AuditStatus = &v
	return s
}

func (s *CreateDataLimitRequest) SetAutoScan(v int32) *CreateDataLimitRequest {
	s.AutoScan = &v
	return s
}

func (s *CreateDataLimitRequest) SetCertificatePermission(v string) *CreateDataLimitRequest {
	s.CertificatePermission = &v
	return s
}

func (s *CreateDataLimitRequest) SetEnable(v int32) *CreateDataLimitRequest {
	s.Enable = &v
	return s
}

func (s *CreateDataLimitRequest) SetEngineType(v string) *CreateDataLimitRequest {
	s.EngineType = &v
	return s
}

func (s *CreateDataLimitRequest) SetEventStatus(v int32) *CreateDataLimitRequest {
	s.EventStatus = &v
	return s
}

func (s *CreateDataLimitRequest) SetFeatureType(v int32) *CreateDataLimitRequest {
	s.FeatureType = &v
	return s
}

func (s *CreateDataLimitRequest) SetInstantlyScan(v bool) *CreateDataLimitRequest {
	s.InstantlyScan = &v
	return s
}

func (s *CreateDataLimitRequest) SetLang(v string) *CreateDataLimitRequest {
	s.Lang = &v
	return s
}

func (s *CreateDataLimitRequest) SetLogStoreDay(v int32) *CreateDataLimitRequest {
	s.LogStoreDay = &v
	return s
}

func (s *CreateDataLimitRequest) SetOcrStatus(v int32) *CreateDataLimitRequest {
	s.OcrStatus = &v
	return s
}

func (s *CreateDataLimitRequest) SetParentId(v string) *CreateDataLimitRequest {
	s.ParentId = &v
	return s
}

func (s *CreateDataLimitRequest) SetPassword(v string) *CreateDataLimitRequest {
	s.Password = &v
	return s
}

func (s *CreateDataLimitRequest) SetPort(v int32) *CreateDataLimitRequest {
	s.Port = &v
	return s
}

func (s *CreateDataLimitRequest) SetResourceType(v int32) *CreateDataLimitRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateDataLimitRequest) SetSamplingSize(v int32) *CreateDataLimitRequest {
	s.SamplingSize = &v
	return s
}

func (s *CreateDataLimitRequest) SetServiceRegionId(v string) *CreateDataLimitRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *CreateDataLimitRequest) SetSourceIp(v string) *CreateDataLimitRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateDataLimitRequest) SetUserName(v string) *CreateDataLimitRequest {
	s.UserName = &v
	return s
}

func (s *CreateDataLimitRequest) Validate() error {
	return dara.Validate(s)
}
