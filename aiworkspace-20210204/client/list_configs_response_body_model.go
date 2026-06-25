// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListConfigsResponseBodyConfigs) *ListConfigsResponseBody
	GetConfigs() []*ListConfigsResponseBodyConfigs
	SetRequestId(v string) *ListConfigsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListConfigsResponseBody
	GetTotalCount() *int64
}

type ListConfigsResponseBody struct {
	// The list of configuration items.
	Configs []*ListConfigsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A******C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigsResponseBody) GetConfigs() []*ListConfigsResponseBodyConfigs {
	return s.Configs
}

func (s *ListConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConfigsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListConfigsResponseBody) SetConfigs(v []*ListConfigsResponseBodyConfigs) *ListConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *ListConfigsResponseBody) SetRequestId(v string) *ListConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigsResponseBody) SetTotalCount(v int64) *ListConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListConfigsResponseBody) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConfigsResponseBodyConfigs struct {
	// The key of the configuration item. The following keys are supported:
	//
	// - tempStoragePath: The path for temporary storage. This key is valid only when CategoryName is set to CommonResourceConfig.
	//
	// - isAutoRecycle: The automatic recycling configuration. This key is valid only when CategoryName is set to DLCAutoRecycle.
	//
	// - priorityConfig: The priority configuration. This key is valid only when CategoryName is set to DLCPriorityConfig or DSWPriorityConfig.
	//
	// - quotaMaximumDuration: The configuration for the maximum runtime of a DLC task in a quota. This key is valid only when CategoryName is set to QuotaMaximumDuration.
	//
	// - predefinedTags: The predefined labels for the workspace. Resources that you create must have these labels.
	//
	// example:
	//
	// tempTableLifecycle
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// oss://***
	ConfigValue     *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The list of labels for the configuration item.
	Labels []*ListConfigsResponseBodyConfigsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s ListConfigsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListConfigsResponseBodyConfigs) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *ListConfigsResponseBodyConfigs) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *ListConfigsResponseBodyConfigs) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListConfigsResponseBodyConfigs) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListConfigsResponseBodyConfigs) GetLabels() []*ListConfigsResponseBodyConfigsLabels {
	return s.Labels
}

func (s *ListConfigsResponseBodyConfigs) SetConfigKey(v string) *ListConfigsResponseBodyConfigs {
	s.ConfigKey = &v
	return s
}

func (s *ListConfigsResponseBodyConfigs) SetConfigValue(v string) *ListConfigsResponseBodyConfigs {
	s.ConfigValue = &v
	return s
}

func (s *ListConfigsResponseBodyConfigs) SetGmtCreateTime(v string) *ListConfigsResponseBodyConfigs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListConfigsResponseBodyConfigs) SetGmtModifiedTime(v string) *ListConfigsResponseBodyConfigs {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListConfigsResponseBodyConfigs) SetLabels(v []*ListConfigsResponseBodyConfigsLabels) *ListConfigsResponseBodyConfigs {
	s.Labels = v
	return s
}

func (s *ListConfigsResponseBodyConfigs) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConfigsResponseBodyConfigsLabels struct {
	// The key of the label.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the label.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListConfigsResponseBodyConfigsLabels) String() string {
	return dara.Prettify(s)
}

func (s ListConfigsResponseBodyConfigsLabels) GoString() string {
	return s.String()
}

func (s *ListConfigsResponseBodyConfigsLabels) GetKey() *string {
	return s.Key
}

func (s *ListConfigsResponseBodyConfigsLabels) GetValue() *string {
	return s.Value
}

func (s *ListConfigsResponseBodyConfigsLabels) SetKey(v string) *ListConfigsResponseBodyConfigsLabels {
	s.Key = &v
	return s
}

func (s *ListConfigsResponseBodyConfigsLabels) SetValue(v string) *ListConfigsResponseBodyConfigsLabels {
	s.Value = &v
	return s
}

func (s *ListConfigsResponseBodyConfigsLabels) Validate() error {
	return dara.Validate(s)
}
