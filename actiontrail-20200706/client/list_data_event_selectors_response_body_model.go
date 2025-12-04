// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataEventSelectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataEventSelectorsResponseBodyData) *ListDataEventSelectorsResponseBody
	GetData() *ListDataEventSelectorsResponseBodyData
	SetRequestId(v string) *ListDataEventSelectorsResponseBody
	GetRequestId() *string
}

type ListDataEventSelectorsResponseBody struct {
	Data *ListDataEventSelectorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 8A74FD2E-A9B9-461C-BCE9-D9668DF1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataEventSelectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataEventSelectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataEventSelectorsResponseBody) GetData() *ListDataEventSelectorsResponseBodyData {
	return s.Data
}

func (s *ListDataEventSelectorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataEventSelectorsResponseBody) SetData(v *ListDataEventSelectorsResponseBodyData) *ListDataEventSelectorsResponseBody {
	s.Data = v
	return s
}

func (s *ListDataEventSelectorsResponseBody) SetRequestId(v string) *ListDataEventSelectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataEventSelectorsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataEventSelectorsResponseBodyData struct {
	DataEventSelectorInfos []*ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos `json:"DataEventSelectorInfos,omitempty" xml:"DataEventSelectorInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// VjE6bHJlTGoxdm1M****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListDataEventSelectorsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataEventSelectorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataEventSelectorsResponseBodyData) GetDataEventSelectorInfos() []*ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos {
	return s.DataEventSelectorInfos
}

func (s *ListDataEventSelectorsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataEventSelectorsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataEventSelectorsResponseBodyData) SetDataEventSelectorInfos(v []*ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos) *ListDataEventSelectorsResponseBodyData {
	s.DataEventSelectorInfos = v
	return s
}

func (s *ListDataEventSelectorsResponseBodyData) SetMaxResults(v int32) *ListDataEventSelectorsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListDataEventSelectorsResponseBodyData) SetNextToken(v string) *ListDataEventSelectorsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListDataEventSelectorsResponseBodyData) Validate() error {
	if s.DataEventSelectorInfos != nil {
		for _, item := range s.DataEventSelectorInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos struct {
	// example:
	//
	// [{"EventName":{"Equals":["GetObject","CopyObject","AppendObject"]},"ReadWriteType":"All","ServiceName":"Oss"}]
	EventSelectors *string `json:"EventSelectors,omitempty" xml:"EventSelectors,omitempty"`
	// example:
	//
	// true
	IsTrailAllRegion   *bool                                                                             `json:"IsTrailAllRegion,omitempty" xml:"IsTrailAllRegion,omitempty"`
	SlsDeliveryConfigs []*ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs `json:"SlsDeliveryConfigs,omitempty" xml:"SlsDeliveryConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// acs:actiontrail:cn-shanghai:159498693826****:trail/trail-name
	TrailArn *string `json:"TrailArn,omitempty" xml:"TrailArn,omitempty"`
	// example:
	//
	// trail-name
	TrailName *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
}

func (s ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos) String() string {
	return dara.Prettify(s)
}

func (s ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos) GoString() string {
	return s.String()
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos) GetEventSelectors() *string {
	return s.EventSelectors
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos) GetIsTrailAllRegion() *bool {
	return s.IsTrailAllRegion
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos) GetSlsDeliveryConfigs() []*ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs {
	return s.SlsDeliveryConfigs
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos) GetTrailArn() *string {
	return s.TrailArn
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos) GetTrailName() *string {
	return s.TrailName
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos) SetEventSelectors(v string) *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos {
	s.EventSelectors = &v
	return s
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos) SetIsTrailAllRegion(v bool) *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos {
	s.IsTrailAllRegion = &v
	return s
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos) SetSlsDeliveryConfigs(v []*ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos {
	s.SlsDeliveryConfigs = v
	return s
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos) SetTrailArn(v string) *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos {
	s.TrailArn = &v
	return s
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos) SetTrailName(v string) *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos {
	s.TrailName = &v
	return s
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfos) Validate() error {
	if s.SlsDeliveryConfigs != nil {
		for _, item := range s.SlsDeliveryConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs struct {
	// example:
	//
	// 2023-09-30T16:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// LogServiceException
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// acs:log:cn-shanghai:159498693826****:project/actiontrail-log-159498693826****-cn-shanghai
	RegionSlsProjectArn *string `json:"RegionSlsProjectArn,omitempty" xml:"RegionSlsProjectArn,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cn-shanghai
	TrailRegion *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
}

func (s ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) GoString() string {
	return s.String()
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) GetRegionSlsProjectArn() *string {
	return s.RegionSlsProjectArn
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) GetStatus() *string {
	return s.Status
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) GetTrailRegion() *string {
	return s.TrailRegion
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) SetCreateTime(v string) *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs {
	s.CreateTime = &v
	return s
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) SetErrorCode(v string) *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs {
	s.ErrorCode = &v
	return s
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) SetErrorMessage(v string) *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) SetRegionSlsProjectArn(v string) *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs {
	s.RegionSlsProjectArn = &v
	return s
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) SetStatus(v string) *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs {
	s.Status = &v
	return s
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) SetTrailRegion(v string) *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs {
	s.TrailRegion = &v
	return s
}

func (s *ListDataEventSelectorsResponseBodyDataDataEventSelectorInfosSlsDeliveryConfigs) Validate() error {
	return dara.Validate(s)
}
