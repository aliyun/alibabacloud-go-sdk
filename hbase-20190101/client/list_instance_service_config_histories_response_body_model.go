// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceServiceConfigHistoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigureHistoryList(v *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) *ListInstanceServiceConfigHistoriesResponseBody
	GetConfigureHistoryList() *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList
	SetPageNumber(v int32) *ListInstanceServiceConfigHistoriesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *ListInstanceServiceConfigHistoriesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *ListInstanceServiceConfigHistoriesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int64) *ListInstanceServiceConfigHistoriesResponseBody
	GetTotalRecordCount() *int64
}

type ListInstanceServiceConfigHistoriesResponseBody struct {
	ConfigureHistoryList *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList `json:"ConfigureHistoryList,omitempty" xml:"ConfigureHistoryList,omitempty" type:"Struct"`
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
	// 658C1549-2C02-4FD9-9490-EB3B285F9DCA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListInstanceServiceConfigHistoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) GetConfigureHistoryList() *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList {
	return s.ConfigureHistoryList
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetConfigureHistoryList(v *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) *ListInstanceServiceConfigHistoriesResponseBody {
	s.ConfigureHistoryList = v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetPageNumber(v int32) *ListInstanceServiceConfigHistoriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetPageRecordCount(v int32) *ListInstanceServiceConfigHistoriesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetRequestId(v string) *ListInstanceServiceConfigHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetTotalRecordCount(v int64) *ListInstanceServiceConfigHistoriesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) Validate() error {
	if s.ConfigureHistoryList != nil {
		if err := s.ConfigureHistoryList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList struct {
	Config []*ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
}

func (s ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) GetConfig() []*ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	return s.Config
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) SetConfig(v []*ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList {
	s.Config = v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) Validate() error {
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

type ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig struct {
	// example:
	//
	// hbase#hbase-site.xml#hbase.client.keyvalue.maxsize
	ConfigureName *string `json:"ConfigureName,omitempty" xml:"ConfigureName,omitempty"`
	// example:
	//
	// 1608708923000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// false
	Effective *string `json:"Effective,omitempty" xml:"Effective,omitempty"`
	// example:
	//
	// 10485770
	NewValue *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	// example:
	//
	// 10485760
	OldValue *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
}

func (s ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) GetConfigureName() *string {
	return s.ConfigureName
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) GetEffective() *string {
	return s.Effective
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) GetNewValue() *string {
	return s.NewValue
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) GetOldValue() *string {
	return s.OldValue
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetConfigureName(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.ConfigureName = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetCreateTime(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.CreateTime = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetEffective(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.Effective = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetNewValue(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.NewValue = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetOldValue(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.OldValue = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) Validate() error {
	return dara.Validate(s)
}
