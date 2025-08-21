// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsUpPeakPublishStreamDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescribeVsUpPeakPublishStreamDatas(v *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas) *DescribeVsUpPeakPublishStreamDataResponseBody
	GetDescribeVsUpPeakPublishStreamDatas() *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas
	SetRequestId(v string) *DescribeVsUpPeakPublishStreamDataResponseBody
	GetRequestId() *string
}

type DescribeVsUpPeakPublishStreamDataResponseBody struct {
	DescribeVsUpPeakPublishStreamDatas *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas `json:"DescribeVsUpPeakPublishStreamDatas,omitempty" xml:"DescribeVsUpPeakPublishStreamDatas,omitempty" type:"Struct"`
	// example:
	//
	// 27A3C548-A699-48F9-97CD-F35D81075AF3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsUpPeakPublishStreamDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsUpPeakPublishStreamDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBody) GetDescribeVsUpPeakPublishStreamDatas() *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas {
	return s.DescribeVsUpPeakPublishStreamDatas
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBody) SetDescribeVsUpPeakPublishStreamDatas(v *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas) *DescribeVsUpPeakPublishStreamDataResponseBody {
	s.DescribeVsUpPeakPublishStreamDatas = v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBody) SetRequestId(v string) *DescribeVsUpPeakPublishStreamDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas struct {
	DescribeVsUpPeakPublishStreamData []*DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData `json:"DescribeVsUpPeakPublishStreamData,omitempty" xml:"DescribeVsUpPeakPublishStreamData,omitempty" type:"Repeated"`
}

func (s DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas) GoString() string {
	return s.String()
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas) GetDescribeVsUpPeakPublishStreamData() []*DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData {
	return s.DescribeVsUpPeakPublishStreamData
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas) SetDescribeVsUpPeakPublishStreamData(v []*DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas {
	s.DescribeVsUpPeakPublishStreamData = v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas) Validate() error {
	return dara.Validate(s)
}

type DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData struct {
	BandWidth *string `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	// example:
	//
	// 1522252320000
	PeakTime *string `json:"PeakTime,omitempty" xml:"PeakTime,omitempty"`
	// example:
	//
	// 31
	PublishStreamNum *int32 `json:"PublishStreamNum,omitempty" xml:"PublishStreamNum,omitempty"`
	// example:
	//
	// 1522166400000
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// example:
	//
	// example.com
	StatName *string `json:"StatName,omitempty" xml:"StatName,omitempty"`
}

func (s DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) GoString() string {
	return s.String()
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) GetBandWidth() *string {
	return s.BandWidth
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) GetPeakTime() *string {
	return s.PeakTime
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) GetPublishStreamNum() *int32 {
	return s.PublishStreamNum
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) GetQueryTime() *string {
	return s.QueryTime
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) GetStatName() *string {
	return s.StatName
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) SetBandWidth(v string) *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData {
	s.BandWidth = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) SetPeakTime(v string) *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData {
	s.PeakTime = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) SetPublishStreamNum(v int32) *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData {
	s.PublishStreamNum = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) SetQueryTime(v string) *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData {
	s.QueryTime = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) SetStatName(v string) *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData {
	s.StatName = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) Validate() error {
	return dara.Validate(s)
}
