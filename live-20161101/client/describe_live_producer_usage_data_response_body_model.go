// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveProducerUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBillProducerData(v *DescribeLiveProducerUsageDataResponseBodyBillProducerData) *DescribeLiveProducerUsageDataResponseBody
	GetBillProducerData() *DescribeLiveProducerUsageDataResponseBodyBillProducerData
	SetEndTime(v string) *DescribeLiveProducerUsageDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeLiveProducerUsageDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveProducerUsageDataResponseBody
	GetStartTime() *string
}

type DescribeLiveProducerUsageDataResponseBody struct {
	// The production studio usage data.
	BillProducerData *DescribeLiveProducerUsageDataResponseBodyBillProducerData `json:"BillProducerData,omitempty" xml:"BillProducerData,omitempty" type:"Struct"`
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2018-09-30T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range for which the data was queried.
	//
	// example:
	//
	// 2018-10-31T15:59:59Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveProducerUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveProducerUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveProducerUsageDataResponseBody) GetBillProducerData() *DescribeLiveProducerUsageDataResponseBodyBillProducerData {
	return s.BillProducerData
}

func (s *DescribeLiveProducerUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveProducerUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveProducerUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveProducerUsageDataResponseBody) SetBillProducerData(v *DescribeLiveProducerUsageDataResponseBodyBillProducerData) *DescribeLiveProducerUsageDataResponseBody {
	s.BillProducerData = v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBody) SetEndTime(v string) *DescribeLiveProducerUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBody) SetRequestId(v string) *DescribeLiveProducerUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBody) SetStartTime(v string) *DescribeLiveProducerUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBody) Validate() error {
	if s.BillProducerData != nil {
		if err := s.BillProducerData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveProducerUsageDataResponseBodyBillProducerData struct {
	BillProducerDataItem []*DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem `json:"BillProducerDataItem,omitempty" xml:"BillProducerDataItem,omitempty" type:"Repeated"`
}

func (s DescribeLiveProducerUsageDataResponseBodyBillProducerData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveProducerUsageDataResponseBodyBillProducerData) GoString() string {
	return s.String()
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerData) GetBillProducerDataItem() []*DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem {
	return s.BillProducerDataItem
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerData) SetBillProducerDataItem(v []*DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) *DescribeLiveProducerUsageDataResponseBodyBillProducerData {
	s.BillProducerDataItem = v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerData) Validate() error {
	if s.BillProducerDataItem != nil {
		for _, item := range s.BillProducerDataItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem struct {
	// The domain name. If domain is specified by the SplitBy parameter, the usage data is returned based on different domain names.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The production studio instance. If instance is specified by the SplitBy parameter, the usage data is returned based on different production studio instances.
	//
	// example:
	//
	// a17d0184-462d-4630-b2a6-8c26dde2****
	Instance *string `json:"Instance,omitempty" xml:"Instance,omitempty"`
	// The duration of high definition streaming. Unit: minutes.
	//
	// example:
	//
	// 6000
	OutputHdDuration *int64 `json:"OutputHdDuration,omitempty" xml:"OutputHdDuration,omitempty"`
	// The duration of low definition streaming. Unit: minutes.
	//
	// example:
	//
	// 1001
	OutputLdDuration *int64 `json:"OutputLdDuration,omitempty" xml:"OutputLdDuration,omitempty"`
	// The duration of standard definition streaming. Unit: minutes.
	//
	// example:
	//
	// 500
	OutputSdDuration *int64 `json:"OutputSdDuration,omitempty" xml:"OutputSdDuration,omitempty"`
	// The region. If region is specified by the SplitBy parameter, the usage data is returned based on different regions.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2018-09-30T00:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The duration of high definition transcoding. Unit: minutes.
	//
	// example:
	//
	// 6777
	TranHdDuration *int64 `json:"TranHdDuration,omitempty" xml:"TranHdDuration,omitempty"`
	// The duration of low definition transcoding. Unit: minutes.
	//
	// example:
	//
	// 111
	TranLdDuration *int64 `json:"TranLdDuration,omitempty" xml:"TranLdDuration,omitempty"`
	// The duration of standard definition transcoding. Unit: minutes.
	//
	// example:
	//
	// 666
	TranSdDuration *int64 `json:"TranSdDuration,omitempty" xml:"TranSdDuration,omitempty"`
	// The type of the production studio. If type is specified by the SplitBy parameter, the usage data is returned based on different types of production studios.
	//
	// example:
	//
	// slidelive
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) GoString() string {
	return s.String()
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) GetInstance() *string {
	return s.Instance
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) GetOutputHdDuration() *int64 {
	return s.OutputHdDuration
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) GetOutputLdDuration() *int64 {
	return s.OutputLdDuration
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) GetOutputSdDuration() *int64 {
	return s.OutputSdDuration
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) GetTranHdDuration() *int64 {
	return s.TranHdDuration
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) GetTranLdDuration() *int64 {
	return s.TranLdDuration
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) GetTranSdDuration() *int64 {
	return s.TranSdDuration
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) GetType() *string {
	return s.Type
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) SetDomainName(v string) *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) SetInstance(v string) *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem {
	s.Instance = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) SetOutputHdDuration(v int64) *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem {
	s.OutputHdDuration = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) SetOutputLdDuration(v int64) *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem {
	s.OutputLdDuration = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) SetOutputSdDuration(v int64) *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem {
	s.OutputSdDuration = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) SetRegion(v string) *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem {
	s.Region = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) SetTimeStamp(v string) *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) SetTranHdDuration(v int64) *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem {
	s.TranHdDuration = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) SetTranLdDuration(v int64) *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem {
	s.TranLdDuration = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) SetTranSdDuration(v int64) *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem {
	s.TranSdDuration = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) SetType(v string) *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem {
	s.Type = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponseBodyBillProducerDataBillProducerDataItem) Validate() error {
	return dara.Validate(s)
}
