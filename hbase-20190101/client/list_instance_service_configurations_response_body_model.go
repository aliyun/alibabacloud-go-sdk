// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceServiceConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigureList(v *ListInstanceServiceConfigurationsResponseBodyConfigureList) *ListInstanceServiceConfigurationsResponseBody
	GetConfigureList() *ListInstanceServiceConfigurationsResponseBodyConfigureList
	SetPageNumber(v int32) *ListInstanceServiceConfigurationsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *ListInstanceServiceConfigurationsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *ListInstanceServiceConfigurationsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int64) *ListInstanceServiceConfigurationsResponseBody
	GetTotalRecordCount() *int64
}

type ListInstanceServiceConfigurationsResponseBody struct {
	ConfigureList *ListInstanceServiceConfigurationsResponseBodyConfigureList `json:"ConfigureList,omitempty" xml:"ConfigureList,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 5B381E36-BCA3-4377-8638-B65C236617D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 42
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListInstanceServiceConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceServiceConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsResponseBody) GetConfigureList() *ListInstanceServiceConfigurationsResponseBodyConfigureList {
	return s.ConfigureList
}

func (s *ListInstanceServiceConfigurationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceServiceConfigurationsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *ListInstanceServiceConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceServiceConfigurationsResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetConfigureList(v *ListInstanceServiceConfigurationsResponseBodyConfigureList) *ListInstanceServiceConfigurationsResponseBody {
	s.ConfigureList = v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetPageNumber(v int32) *ListInstanceServiceConfigurationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetPageRecordCount(v int32) *ListInstanceServiceConfigurationsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetRequestId(v string) *ListInstanceServiceConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetTotalRecordCount(v int64) *ListInstanceServiceConfigurationsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBody) Validate() error {
	if s.ConfigureList != nil {
		if err := s.ConfigureList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceServiceConfigurationsResponseBodyConfigureList struct {
	Config []*ListInstanceServiceConfigurationsResponseBodyConfigureListConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
}

func (s ListInstanceServiceConfigurationsResponseBodyConfigureList) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceServiceConfigurationsResponseBodyConfigureList) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureList) GetConfig() []*ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	return s.Config
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureList) SetConfig(v []*ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) *ListInstanceServiceConfigurationsResponseBodyConfigureList {
	s.Config = v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureList) Validate() error {
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

type ListInstanceServiceConfigurationsResponseBodyConfigureListConfig struct {
	// example:
	//
	// hbase#hbase-site.xml#hbase.client.keyvalue.maxsize
	ConfigureName *string `json:"ConfigureName,omitempty" xml:"ConfigureName,omitempty"`
	// example:
	//
	// INT
	ConfigureUnit *string `json:"ConfigureUnit,omitempty" xml:"ConfigureUnit,omitempty"`
	// example:
	//
	// 10485760
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// hbase client keyvalue maxsize
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	NeedRestart *string `json:"NeedRestart,omitempty" xml:"NeedRestart,omitempty"`
	// example:
	//
	// 10485760
	RunningValue *string `json:"RunningValue,omitempty" xml:"RunningValue,omitempty"`
	// example:
	//
	// R[10485760,52428800]
	ValueRange *string `json:"ValueRange,omitempty" xml:"ValueRange,omitempty"`
}

func (s ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) GetConfigureName() *string {
	return s.ConfigureName
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) GetConfigureUnit() *string {
	return s.ConfigureUnit
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) GetDescription() *string {
	return s.Description
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) GetNeedRestart() *string {
	return s.NeedRestart
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) GetRunningValue() *string {
	return s.RunningValue
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) GetValueRange() *string {
	return s.ValueRange
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetConfigureName(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.ConfigureName = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetConfigureUnit(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.ConfigureUnit = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetDefaultValue(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.DefaultValue = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetDescription(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.Description = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetNeedRestart(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.NeedRestart = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetRunningValue(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.RunningValue = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetValueRange(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.ValueRange = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) Validate() error {
	return dara.Validate(s)
}
