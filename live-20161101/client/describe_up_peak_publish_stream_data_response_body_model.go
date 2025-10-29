// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpPeakPublishStreamDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescribeUpPeakPublishStreamDatas(v *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatas) *DescribeUpPeakPublishStreamDataResponseBody
	GetDescribeUpPeakPublishStreamDatas() *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatas
	SetRequestId(v string) *DescribeUpPeakPublishStreamDataResponseBody
	GetRequestId() *string
}

type DescribeUpPeakPublishStreamDataResponseBody struct {
	// The information about the peak number of concurrently ingested streams on each day.
	DescribeUpPeakPublishStreamDatas *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatas `json:"DescribeUpPeakPublishStreamDatas,omitempty" xml:"DescribeUpPeakPublishStreamDatas,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6CFDE7AB-571A-14EA-B072-989FF753****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUpPeakPublishStreamDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpPeakPublishStreamDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUpPeakPublishStreamDataResponseBody) GetDescribeUpPeakPublishStreamDatas() *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatas {
	return s.DescribeUpPeakPublishStreamDatas
}

func (s *DescribeUpPeakPublishStreamDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUpPeakPublishStreamDataResponseBody) SetDescribeUpPeakPublishStreamDatas(v *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatas) *DescribeUpPeakPublishStreamDataResponseBody {
	s.DescribeUpPeakPublishStreamDatas = v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponseBody) SetRequestId(v string) *DescribeUpPeakPublishStreamDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponseBody) Validate() error {
	if s.DescribeUpPeakPublishStreamDatas != nil {
		if err := s.DescribeUpPeakPublishStreamDatas.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatas struct {
	DescribeUpPeakPublishStreamData []*DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData `json:"DescribeUpPeakPublishStreamData,omitempty" xml:"DescribeUpPeakPublishStreamData,omitempty" type:"Repeated"`
}

func (s DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatas) GoString() string {
	return s.String()
}

func (s *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatas) GetDescribeUpPeakPublishStreamData() []*DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData {
	return s.DescribeUpPeakPublishStreamData
}

func (s *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatas) SetDescribeUpPeakPublishStreamData(v []*DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatas {
	s.DescribeUpPeakPublishStreamData = v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatas) Validate() error {
	if s.DescribeUpPeakPublishStreamData != nil {
		for _, item := range s.DescribeUpPeakPublishStreamData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData struct {
	// The daily peak inbound bandwidth.
	//
	// example:
	//
	// 777.2727083333333
	BandWidth *string `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	// The time when the daily peak number of concurrently ingested streams is reached.
	//
	// example:
	//
	// 1522180000000
	PeakTime *string `json:"PeakTime,omitempty" xml:"PeakTime,omitempty"`
	// The daily peak number of concurrently ingested streams.
	//
	// example:
	//
	// 36
	PublishStreamNum *int32 `json:"PublishStreamNum,omitempty" xml:"PublishStreamNum,omitempty"`
	// The time queried on the day.
	//
	// example:
	//
	// 1522080000000
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// The category of the statistical data. If the DomainSwitch parameter is set to on, the value of this parameter is the domain name. If the DomainSwitch parameter is set to off or not specified, the value of this parameter is the user ID.
	//
	// example:
	//
	// push-live.aliyuncs.com
	StatName *string `json:"StatName,omitempty" xml:"StatName,omitempty"`
}

func (s DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) GoString() string {
	return s.String()
}

func (s *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) GetBandWidth() *string {
	return s.BandWidth
}

func (s *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) GetPeakTime() *string {
	return s.PeakTime
}

func (s *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) GetPublishStreamNum() *int32 {
	return s.PublishStreamNum
}

func (s *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) GetQueryTime() *string {
	return s.QueryTime
}

func (s *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) GetStatName() *string {
	return s.StatName
}

func (s *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) SetBandWidth(v string) *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData {
	s.BandWidth = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) SetPeakTime(v string) *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData {
	s.PeakTime = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) SetPublishStreamNum(v int32) *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData {
	s.PublishStreamNum = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) SetQueryTime(v string) *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData {
	s.QueryTime = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) SetStatName(v string) *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData {
	s.StatName = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponseBodyDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) Validate() error {
	return dara.Validate(s)
}
