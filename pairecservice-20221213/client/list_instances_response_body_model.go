// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody
	GetInstances() []*ListInstancesResponseBodyInstances
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListInstancesResponseBody
	GetTotalCount() *int32
}

type ListInstancesResponseBody struct {
	// A list of instances.
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BDB621CB-A81E-5D39-8793-39A365CBCC74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned instances.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetInstances() []*ListInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyInstances struct {
	// The billing method of the instance. Only `Subscription` (prepaid) is supported.
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
	Config *ListInstancesResponseBodyInstancesConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2022-12-14 00:00:00.0
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2022-10-13 17:34:52.0
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the instance was last modified.
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
	// The configuration of the operating tool.
	OperatingTool *ListInstancesResponseBodyInstancesOperatingTool `json:"OperatingTool,omitempty" xml:"OperatingTool,omitempty" type:"Struct"`
	// The region ID. Valid values:
	//
	// - `cn-shenzhen`: China (Shenzhen)
	//
	// - `cn-hangzhou`: China (Hangzhou)
	//
	// - `cn-beijing`: China (Beijing)
	//
	// - `cn-shanghai`: China (Shanghai)
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance status. Valid values:
	//
	// - `Initializing`: The instance is initializing.
	//
	// - `Stopped`: The instance is stopped.
	//
	// - `Running`: The instance is running.
	//
	// example:
	//
	// Initializing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The instance type. Valid values:
	//
	// - `basic`: Basic Edition
	//
	// - `high-level`: High-level Edition
	//
	// - `advanced`: Advanced Edition
	//
	// - `standard`: Standard Edition
	//
	// example:
	//
	// basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListInstancesResponseBodyInstances) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListInstancesResponseBodyInstances) GetConfig() *ListInstancesResponseBodyInstancesConfig {
	return s.Config
}

func (s *ListInstancesResponseBodyInstances) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListInstancesResponseBodyInstances) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListInstancesResponseBodyInstances) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyInstances) GetOperatingTool() *ListInstancesResponseBodyInstancesOperatingTool {
	return s.OperatingTool
}

func (s *ListInstancesResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstances) GetType() *string {
	return s.Type
}

func (s *ListInstancesResponseBodyInstances) SetChargeType(v string) *ListInstancesResponseBodyInstances {
	s.ChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCommodityCode(v string) *ListInstancesResponseBodyInstances {
	s.CommodityCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetConfig(v *ListInstancesResponseBodyInstancesConfig) *ListInstancesResponseBodyInstances {
	s.Config = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetExpiredTime(v string) *ListInstancesResponseBodyInstances {
	s.ExpiredTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtCreateTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtModifiedTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetOperatingTool(v *ListInstancesResponseBodyInstancesOperatingTool) *ListInstancesResponseBodyInstances {
	s.OperatingTool = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetRegionId(v string) *ListInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetType(v string) *ListInstancesResponseBodyInstances {
	s.Type = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.OperatingTool != nil {
		if err := s.OperatingTool.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesResponseBodyInstancesConfig struct {
	// A list of data management configurations.
	DataManagements []*ListInstancesResponseBodyInstancesConfigDataManagements `json:"DataManagements,omitempty" xml:"DataManagements,omitempty" type:"Repeated"`
	// A list of service engines.
	Engines []*ListInstancesResponseBodyInstancesConfigEngines `json:"Engines,omitempty" xml:"Engines,omitempty" type:"Repeated"`
	// A list of monitoring components.
	Monitors []*ListInstancesResponseBodyInstancesConfigMonitors `json:"Monitors,omitempty" xml:"Monitors,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyInstancesConfig) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesConfig) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesConfig) GetDataManagements() []*ListInstancesResponseBodyInstancesConfigDataManagements {
	return s.DataManagements
}

func (s *ListInstancesResponseBodyInstancesConfig) GetEngines() []*ListInstancesResponseBodyInstancesConfigEngines {
	return s.Engines
}

func (s *ListInstancesResponseBodyInstancesConfig) GetMonitors() []*ListInstancesResponseBodyInstancesConfigMonitors {
	return s.Monitors
}

func (s *ListInstancesResponseBodyInstancesConfig) SetDataManagements(v []*ListInstancesResponseBodyInstancesConfigDataManagements) *ListInstancesResponseBodyInstancesConfig {
	s.DataManagements = v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfig) SetEngines(v []*ListInstancesResponseBodyInstancesConfigEngines) *ListInstancesResponseBodyInstancesConfig {
	s.Engines = v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfig) SetMonitors(v []*ListInstancesResponseBodyInstancesConfigMonitors) *ListInstancesResponseBodyInstancesConfig {
	s.Monitors = v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfig) Validate() error {
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

type ListInstancesResponseBodyInstancesConfigDataManagements struct {
	// The component code.
	//
	// example:
	//
	// storage
	ComponentCode *string `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	// The metadata of the component.
	Meta map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// The component type.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesResponseBodyInstancesConfigDataManagements) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesConfigDataManagements) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesConfigDataManagements) GetComponentCode() *string {
	return s.ComponentCode
}

func (s *ListInstancesResponseBodyInstancesConfigDataManagements) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *ListInstancesResponseBodyInstancesConfigDataManagements) GetType() *string {
	return s.Type
}

func (s *ListInstancesResponseBodyInstancesConfigDataManagements) SetComponentCode(v string) *ListInstancesResponseBodyInstancesConfigDataManagements {
	s.ComponentCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigDataManagements) SetMeta(v map[string]interface{}) *ListInstancesResponseBodyInstancesConfigDataManagements {
	s.Meta = v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigDataManagements) SetType(v string) *ListInstancesResponseBodyInstancesConfigDataManagements {
	s.Type = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigDataManagements) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesConfigEngines struct {
	// The component code.
	//
	// example:
	//
	// feature
	ComponentCode *string `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	// The metadata of the component.
	Meta map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// The component type.
	//
	// example:
	//
	// Hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesResponseBodyInstancesConfigEngines) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesConfigEngines) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesConfigEngines) GetComponentCode() *string {
	return s.ComponentCode
}

func (s *ListInstancesResponseBodyInstancesConfigEngines) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *ListInstancesResponseBodyInstancesConfigEngines) GetType() *string {
	return s.Type
}

func (s *ListInstancesResponseBodyInstancesConfigEngines) SetComponentCode(v string) *ListInstancesResponseBodyInstancesConfigEngines {
	s.ComponentCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigEngines) SetMeta(v map[string]interface{}) *ListInstancesResponseBodyInstancesConfigEngines {
	s.Meta = v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigEngines) SetType(v string) *ListInstancesResponseBodyInstancesConfigEngines {
	s.Type = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigEngines) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesConfigMonitors struct {
	// The component code.
	//
	// example:
	//
	// featuresets
	ComponentCode *string `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	// The metadata of the component.
	Meta map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// The component type.
	//
	// example:
	//
	// Platform
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesResponseBodyInstancesConfigMonitors) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesConfigMonitors) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesConfigMonitors) GetComponentCode() *string {
	return s.ComponentCode
}

func (s *ListInstancesResponseBodyInstancesConfigMonitors) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *ListInstancesResponseBodyInstancesConfigMonitors) GetType() *string {
	return s.Type
}

func (s *ListInstancesResponseBodyInstancesConfigMonitors) SetComponentCode(v string) *ListInstancesResponseBodyInstancesConfigMonitors {
	s.ComponentCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigMonitors) SetMeta(v map[string]interface{}) *ListInstancesResponseBodyInstancesConfigMonitors {
	s.Meta = v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigMonitors) SetType(v string) *ListInstancesResponseBodyInstancesConfigMonitors {
	s.Type = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigMonitors) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesOperatingTool struct {
	// Specifies whether the operating tool is enabled for the instance. Valid values:
	//
	// - `true`: The tool is enabled.
	//
	// - `false`: The tool is disabled.
	//
	// example:
	//
	// true
	IsEnable *bool `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
}

func (s ListInstancesResponseBodyInstancesOperatingTool) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesOperatingTool) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesOperatingTool) GetIsEnable() *bool {
	return s.IsEnable
}

func (s *ListInstancesResponseBodyInstancesOperatingTool) SetIsEnable(v bool) *ListInstancesResponseBodyInstancesOperatingTool {
	s.IsEnable = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesOperatingTool) Validate() error {
	return dara.Validate(s)
}
