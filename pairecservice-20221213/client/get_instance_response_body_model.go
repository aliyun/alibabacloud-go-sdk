// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *GetInstanceResponseBody
	GetChargeType() *string
	SetCommodityCode(v string) *GetInstanceResponseBody
	GetCommodityCode() *string
	SetConfig(v *GetInstanceResponseBodyConfig) *GetInstanceResponseBody
	GetConfig() *GetInstanceResponseBodyConfig
	SetExpiredTime(v string) *GetInstanceResponseBody
	GetExpiredTime() *string
	SetFeatureStoreInfo(v *GetInstanceResponseBodyFeatureStoreInfo) *GetInstanceResponseBody
	GetFeatureStoreInfo() *GetInstanceResponseBodyFeatureStoreInfo
	SetGmtCreateTime(v string) *GetInstanceResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetInstanceResponseBody
	GetGmtModifiedTime() *string
	SetInstanceId(v string) *GetInstanceResponseBody
	GetInstanceId() *string
	SetOperatingTool(v *GetInstanceResponseBodyOperatingTool) *GetInstanceResponseBody
	GetOperatingTool() *GetInstanceResponseBodyOperatingTool
	SetRecommendCustomization(v *GetInstanceResponseBodyRecommendCustomization) *GetInstanceResponseBody
	GetRecommendCustomization() *GetInstanceResponseBodyRecommendCustomization
	SetRegionId(v string) *GetInstanceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetInstanceResponseBody
	GetStatus() *string
	SetType(v string) *GetInstanceResponseBody
	GetType() *string
}

type GetInstanceResponseBody struct {
	// The billing type of the instance. Currently, only Subscription (prepayment) is supported.
	//
	// example:
	//
	// Subscription
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The commodity code of the instance.
	//
	// example:
	//
	// airec_developers_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The instance configuration.
	Config *GetInstanceResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2022-12-14 00:00:00.0
	ExpiredTime      *string                                  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FeatureStoreInfo *GetInstanceResponseBodyFeatureStoreInfo `json:"FeatureStoreInfo,omitempty" xml:"FeatureStoreInfo,omitempty" type:"Struct"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2022-10-13 17:34:52.0
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the instance was last updated.
	//
	// example:
	//
	// 2022-11-05 09:02:30.0
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The configuration of the operations tool.
	OperatingTool          *GetInstanceResponseBodyOperatingTool          `json:"OperatingTool,omitempty" xml:"OperatingTool,omitempty" type:"Struct"`
	RecommendCustomization *GetInstanceResponseBodyRecommendCustomization `json:"RecommendCustomization,omitempty" xml:"RecommendCustomization,omitempty" type:"Struct"`
	// The region ID. Valid values:
	//
	// - cn-shenzhen: China (Shenzhen).
	//
	// - cn-hangzhou: China (Hangzhou).
	//
	// - cn-beijing: China (Beijing).
	//
	// - cn-shanghai: China (Shanghai).
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The instance status. Valid values:
	//
	// - Initializing: The instance is being initialized.
	//
	// - Stopped: The instance is stopped.
	//
	// - Running: The instance is running.
	//
	// example:
	//
	// Initializing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The instance type. Valid values:
	//
	// - basic: Basic Edition.
	//
	// - highleve: Upgraded Edition.
	//
	// - advance: Advanced Edition.
	//
	// - standard: Standard Edition.
	//
	// example:
	//
	// basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetInstanceResponseBody) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetInstanceResponseBody) GetConfig() *GetInstanceResponseBodyConfig {
	return s.Config
}

func (s *GetInstanceResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *GetInstanceResponseBody) GetFeatureStoreInfo() *GetInstanceResponseBodyFeatureStoreInfo {
	return s.FeatureStoreInfo
}

func (s *GetInstanceResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetInstanceResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBody) GetOperatingTool() *GetInstanceResponseBodyOperatingTool {
	return s.OperatingTool
}

func (s *GetInstanceResponseBody) GetRecommendCustomization() *GetInstanceResponseBodyRecommendCustomization {
	return s.RecommendCustomization
}

func (s *GetInstanceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBody) GetType() *string {
	return s.Type
}

func (s *GetInstanceResponseBody) SetChargeType(v string) *GetInstanceResponseBody {
	s.ChargeType = &v
	return s
}

func (s *GetInstanceResponseBody) SetCommodityCode(v string) *GetInstanceResponseBody {
	s.CommodityCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetConfig(v *GetInstanceResponseBodyConfig) *GetInstanceResponseBody {
	s.Config = v
	return s
}

func (s *GetInstanceResponseBody) SetExpiredTime(v string) *GetInstanceResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetFeatureStoreInfo(v *GetInstanceResponseBodyFeatureStoreInfo) *GetInstanceResponseBody {
	s.FeatureStoreInfo = v
	return s
}

func (s *GetInstanceResponseBody) SetGmtCreateTime(v string) *GetInstanceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetGmtModifiedTime(v string) *GetInstanceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceId(v string) *GetInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBody) SetOperatingTool(v *GetInstanceResponseBodyOperatingTool) *GetInstanceResponseBody {
	s.OperatingTool = v
	return s
}

func (s *GetInstanceResponseBody) SetRecommendCustomization(v *GetInstanceResponseBodyRecommendCustomization) *GetInstanceResponseBody {
	s.RecommendCustomization = v
	return s
}

func (s *GetInstanceResponseBody) SetRegionId(v string) *GetInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetStatus(v string) *GetInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBody) SetType(v string) *GetInstanceResponseBody {
	s.Type = &v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.FeatureStoreInfo != nil {
		if err := s.FeatureStoreInfo.Validate(); err != nil {
			return err
		}
	}
	if s.OperatingTool != nil {
		if err := s.OperatingTool.Validate(); err != nil {
			return err
		}
	}
	if s.RecommendCustomization != nil {
		if err := s.RecommendCustomization.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceResponseBodyConfig struct {
	// The list of modeling environments.
	DataManagements []*GetInstanceResponseBodyConfigDataManagements `json:"DataManagements,omitempty" xml:"DataManagements,omitempty" type:"Repeated"`
	// The list of service engines.
	Engines []*GetInstanceResponseBodyConfigEngines `json:"Engines,omitempty" xml:"Engines,omitempty" type:"Repeated"`
	// The list of supporting features.
	Monitors []*GetInstanceResponseBodyConfigMonitors `json:"Monitors,omitempty" xml:"Monitors,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyConfig) GetDataManagements() []*GetInstanceResponseBodyConfigDataManagements {
	return s.DataManagements
}

func (s *GetInstanceResponseBodyConfig) GetEngines() []*GetInstanceResponseBodyConfigEngines {
	return s.Engines
}

func (s *GetInstanceResponseBodyConfig) GetMonitors() []*GetInstanceResponseBodyConfigMonitors {
	return s.Monitors
}

func (s *GetInstanceResponseBodyConfig) SetDataManagements(v []*GetInstanceResponseBodyConfigDataManagements) *GetInstanceResponseBodyConfig {
	s.DataManagements = v
	return s
}

func (s *GetInstanceResponseBodyConfig) SetEngines(v []*GetInstanceResponseBodyConfigEngines) *GetInstanceResponseBodyConfig {
	s.Engines = v
	return s
}

func (s *GetInstanceResponseBodyConfig) SetMonitors(v []*GetInstanceResponseBodyConfigMonitors) *GetInstanceResponseBodyConfig {
	s.Monitors = v
	return s
}

func (s *GetInstanceResponseBodyConfig) Validate() error {
	if s.DataManagements != nil {
		for _, item := range s.DataManagements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Engines != nil {
		for _, item := range s.Engines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Monitors != nil {
		for _, item := range s.Monitors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceResponseBodyConfigDataManagements struct {
	// The component code.
	//
	// example:
	//
	// storage
	ComponentCode *string `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	// The metadata.
	Meta map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// The component type.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceResponseBodyConfigDataManagements) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyConfigDataManagements) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyConfigDataManagements) GetComponentCode() *string {
	return s.ComponentCode
}

func (s *GetInstanceResponseBodyConfigDataManagements) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *GetInstanceResponseBodyConfigDataManagements) GetType() *string {
	return s.Type
}

func (s *GetInstanceResponseBodyConfigDataManagements) SetComponentCode(v string) *GetInstanceResponseBodyConfigDataManagements {
	s.ComponentCode = &v
	return s
}

func (s *GetInstanceResponseBodyConfigDataManagements) SetMeta(v map[string]interface{}) *GetInstanceResponseBodyConfigDataManagements {
	s.Meta = v
	return s
}

func (s *GetInstanceResponseBodyConfigDataManagements) SetType(v string) *GetInstanceResponseBodyConfigDataManagements {
	s.Type = &v
	return s
}

func (s *GetInstanceResponseBodyConfigDataManagements) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyConfigEngines struct {
	// The component code.
	//
	// example:
	//
	// feature
	ComponentCode *string `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	// The metadata.
	Meta map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// The component type.
	//
	// example:
	//
	// Hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceResponseBodyConfigEngines) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyConfigEngines) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyConfigEngines) GetComponentCode() *string {
	return s.ComponentCode
}

func (s *GetInstanceResponseBodyConfigEngines) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *GetInstanceResponseBodyConfigEngines) GetType() *string {
	return s.Type
}

func (s *GetInstanceResponseBodyConfigEngines) SetComponentCode(v string) *GetInstanceResponseBodyConfigEngines {
	s.ComponentCode = &v
	return s
}

func (s *GetInstanceResponseBodyConfigEngines) SetMeta(v map[string]interface{}) *GetInstanceResponseBodyConfigEngines {
	s.Meta = v
	return s
}

func (s *GetInstanceResponseBodyConfigEngines) SetType(v string) *GetInstanceResponseBodyConfigEngines {
	s.Type = &v
	return s
}

func (s *GetInstanceResponseBodyConfigEngines) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyConfigMonitors struct {
	// The component code.
	//
	// example:
	//
	// featuresets
	ComponentCode *string `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	// The metadata.
	Meta map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// The component type.
	//
	// example:
	//
	// Platform
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceResponseBodyConfigMonitors) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyConfigMonitors) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyConfigMonitors) GetComponentCode() *string {
	return s.ComponentCode
}

func (s *GetInstanceResponseBodyConfigMonitors) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *GetInstanceResponseBodyConfigMonitors) GetType() *string {
	return s.Type
}

func (s *GetInstanceResponseBodyConfigMonitors) SetComponentCode(v string) *GetInstanceResponseBodyConfigMonitors {
	s.ComponentCode = &v
	return s
}

func (s *GetInstanceResponseBodyConfigMonitors) SetMeta(v map[string]interface{}) *GetInstanceResponseBodyConfigMonitors {
	s.Meta = v
	return s
}

func (s *GetInstanceResponseBodyConfigMonitors) SetType(v string) *GetInstanceResponseBodyConfigMonitors {
	s.Type = &v
	return s
}

func (s *GetInstanceResponseBodyConfigMonitors) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyFeatureStoreInfo struct {
	FeatureDBStatus *string `json:"FeatureDBStatus,omitempty" xml:"FeatureDBStatus,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceResponseBodyFeatureStoreInfo) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyFeatureStoreInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyFeatureStoreInfo) GetFeatureDBStatus() *string {
	return s.FeatureDBStatus
}

func (s *GetInstanceResponseBodyFeatureStoreInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyFeatureStoreInfo) SetFeatureDBStatus(v string) *GetInstanceResponseBodyFeatureStoreInfo {
	s.FeatureDBStatus = &v
	return s
}

func (s *GetInstanceResponseBodyFeatureStoreInfo) SetInstanceId(v string) *GetInstanceResponseBodyFeatureStoreInfo {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyFeatureStoreInfo) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyOperatingTool struct {
	// Indicates whether the operations tool is enabled for the instance. Valid values:
	//
	// - True: Enabled.
	//
	// - False: Not enabled.
	//
	// example:
	//
	// True
	IsEnable *bool `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
}

func (s GetInstanceResponseBodyOperatingTool) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyOperatingTool) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyOperatingTool) GetIsEnable() *bool {
	return s.IsEnable
}

func (s *GetInstanceResponseBodyOperatingTool) SetIsEnable(v bool) *GetInstanceResponseBodyOperatingTool {
	s.IsEnable = &v
	return s
}

func (s *GetInstanceResponseBodyOperatingTool) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyRecommendCustomization struct {
	IsEnable *bool `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
}

func (s GetInstanceResponseBodyRecommendCustomization) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyRecommendCustomization) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyRecommendCustomization) GetIsEnable() *bool {
	return s.IsEnable
}

func (s *GetInstanceResponseBodyRecommendCustomization) SetIsEnable(v bool) *GetInstanceResponseBodyRecommendCustomization {
	s.IsEnable = &v
	return s
}

func (s *GetInstanceResponseBodyRecommendCustomization) Validate() error {
	return dara.Validate(s)
}
