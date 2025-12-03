// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterServiceConfigHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigHistoryList(v *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) *ListClusterServiceConfigHistoryResponseBody
	GetConfigHistoryList() *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList
	SetPageNumber(v int32) *ListClusterServiceConfigHistoryResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *ListClusterServiceConfigHistoryResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *ListClusterServiceConfigHistoryResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *ListClusterServiceConfigHistoryResponseBody
	GetTotalRecordCount() *int32
}

type ListClusterServiceConfigHistoryResponseBody struct {
	ConfigHistoryList *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList `json:"ConfigHistoryList,omitempty" xml:"ConfigHistoryList,omitempty" type:"Struct"`
	PageNumber        *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount   *int32                                                        `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId         *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount  *int32                                                        `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListClusterServiceConfigHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterServiceConfigHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryResponseBody) GetConfigHistoryList() *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList {
	return s.ConfigHistoryList
}

func (s *ListClusterServiceConfigHistoryResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClusterServiceConfigHistoryResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *ListClusterServiceConfigHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterServiceConfigHistoryResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetConfigHistoryList(v *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) *ListClusterServiceConfigHistoryResponseBody {
	s.ConfigHistoryList = v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetPageNumber(v int32) *ListClusterServiceConfigHistoryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetPageRecordCount(v int32) *ListClusterServiceConfigHistoryResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetRequestId(v string) *ListClusterServiceConfigHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetTotalRecordCount(v int32) *ListClusterServiceConfigHistoryResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBody) Validate() error {
	if s.ConfigHistoryList != nil {
		if err := s.ConfigHistoryList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListClusterServiceConfigHistoryResponseBodyConfigHistoryList struct {
	ConfigHistory []*ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory `json:"ConfigHistory,omitempty" xml:"ConfigHistory,omitempty" type:"Repeated"`
}

func (s ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) String() string {
	return dara.Prettify(s)
}

func (s ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) GetConfigHistory() []*ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	return s.ConfigHistory
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) SetConfigHistory(v []*ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList {
	s.ConfigHistory = v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) Validate() error {
	if s.ConfigHistory != nil {
		for _, item := range s.ConfigHistory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Effective  *string `json:"Effective,omitempty" xml:"Effective,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NewValue   *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	OldValue   *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
}

func (s ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) String() string {
	return dara.Prettify(s)
}

func (s ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) GetEffective() *string {
	return s.Effective
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) GetName() *string {
	return s.Name
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) GetNewValue() *string {
	return s.NewValue
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) GetOldValue() *string {
	return s.OldValue
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetCreateTime(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.CreateTime = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetEffective(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.Effective = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetName(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.Name = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetNewValue(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.NewValue = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetOldValue(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.OldValue = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) Validate() error {
	return dara.Validate(s)
}
