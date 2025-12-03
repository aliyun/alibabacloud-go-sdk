// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterServiceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListClusterServiceConfigResponseBody
	GetClusterId() *string
	SetClusterName(v string) *ListClusterServiceConfigResponseBody
	GetClusterName() *string
	SetConfigList(v *ListClusterServiceConfigResponseBodyConfigList) *ListClusterServiceConfigResponseBody
	GetConfigList() *ListClusterServiceConfigResponseBodyConfigList
	SetPageNumber(v int32) *ListClusterServiceConfigResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *ListClusterServiceConfigResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *ListClusterServiceConfigResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *ListClusterServiceConfigResponseBody
	GetTotalRecordCount() *int32
}

type ListClusterServiceConfigResponseBody struct {
	ClusterId        *string                                         `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName      *string                                         `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ConfigList       *ListClusterServiceConfigResponseBodyConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Struct"`
	PageNumber       *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                          `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                                          `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListClusterServiceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterServiceConfigResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListClusterServiceConfigResponseBody) GetConfigList() *ListClusterServiceConfigResponseBodyConfigList {
	return s.ConfigList
}

func (s *ListClusterServiceConfigResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClusterServiceConfigResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *ListClusterServiceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterServiceConfigResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListClusterServiceConfigResponseBody) SetClusterId(v string) *ListClusterServiceConfigResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ListClusterServiceConfigResponseBody) SetClusterName(v string) *ListClusterServiceConfigResponseBody {
	s.ClusterName = &v
	return s
}

func (s *ListClusterServiceConfigResponseBody) SetConfigList(v *ListClusterServiceConfigResponseBodyConfigList) *ListClusterServiceConfigResponseBody {
	s.ConfigList = v
	return s
}

func (s *ListClusterServiceConfigResponseBody) SetPageNumber(v int32) *ListClusterServiceConfigResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterServiceConfigResponseBody) SetPageRecordCount(v int32) *ListClusterServiceConfigResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListClusterServiceConfigResponseBody) SetRequestId(v string) *ListClusterServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterServiceConfigResponseBody) SetTotalRecordCount(v int32) *ListClusterServiceConfigResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListClusterServiceConfigResponseBody) Validate() error {
	if s.ConfigList != nil {
		if err := s.ConfigList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListClusterServiceConfigResponseBodyConfigList struct {
	Config []*ListClusterServiceConfigResponseBodyConfigListConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
}

func (s ListClusterServiceConfigResponseBodyConfigList) String() string {
	return dara.Prettify(s)
}

func (s ListClusterServiceConfigResponseBodyConfigList) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigResponseBodyConfigList) GetConfig() []*ListClusterServiceConfigResponseBodyConfigListConfig {
	return s.Config
}

func (s *ListClusterServiceConfigResponseBodyConfigList) SetConfig(v []*ListClusterServiceConfigResponseBodyConfigListConfig) *ListClusterServiceConfigResponseBodyConfigList {
	s.Config = v
	return s
}

func (s *ListClusterServiceConfigResponseBodyConfigList) Validate() error {
	if s.Config != nil {
		for _, item := range s.Config {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterServiceConfigResponseBodyConfigListConfig struct {
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedRestart  *string `json:"NeedRestart,omitempty" xml:"NeedRestart,omitempty"`
	RunningValue *string `json:"RunningValue,omitempty" xml:"RunningValue,omitempty"`
	Unit         *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	ValueRange   *string `json:"ValueRange,omitempty" xml:"ValueRange,omitempty"`
}

func (s ListClusterServiceConfigResponseBodyConfigListConfig) String() string {
	return dara.Prettify(s)
}

func (s ListClusterServiceConfigResponseBodyConfigListConfig) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) GetDescription() *string {
	return s.Description
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) GetName() *string {
	return s.Name
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) GetNeedRestart() *string {
	return s.NeedRestart
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) GetRunningValue() *string {
	return s.RunningValue
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) GetUnit() *string {
	return s.Unit
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) GetValueRange() *string {
	return s.ValueRange
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) SetDefaultValue(v string) *ListClusterServiceConfigResponseBodyConfigListConfig {
	s.DefaultValue = &v
	return s
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) SetDescription(v string) *ListClusterServiceConfigResponseBodyConfigListConfig {
	s.Description = &v
	return s
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) SetName(v string) *ListClusterServiceConfigResponseBodyConfigListConfig {
	s.Name = &v
	return s
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) SetNeedRestart(v string) *ListClusterServiceConfigResponseBodyConfigListConfig {
	s.NeedRestart = &v
	return s
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) SetRunningValue(v string) *ListClusterServiceConfigResponseBodyConfigListConfig {
	s.RunningValue = &v
	return s
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) SetUnit(v string) *ListClusterServiceConfigResponseBodyConfigListConfig {
	s.Unit = &v
	return s
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) SetValueRange(v string) *ListClusterServiceConfigResponseBodyConfigListConfig {
	s.ValueRange = &v
	return s
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) Validate() error {
	return dara.Validate(s)
}
