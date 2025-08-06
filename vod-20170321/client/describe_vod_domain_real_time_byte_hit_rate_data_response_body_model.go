// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeByteHitRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeVodDomainRealTimeByteHitRateDataResponseBodyData) *DescribeVodDomainRealTimeByteHitRateDataResponseBody
	GetData() *DescribeVodDomainRealTimeByteHitRateDataResponseBodyData
	SetRequestId(v string) *DescribeVodDomainRealTimeByteHitRateDataResponseBody
	GetRequestId() *string
}

type DescribeVodDomainRealTimeByteHitRateDataResponseBody struct {
	// The returned data.
	Data *DescribeVodDomainRealTimeByteHitRateDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 70A26B11-3673-479C-AEA8-E03FC5D3496D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainRealTimeByteHitRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeByteHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponseBody) GetData() *DescribeVodDomainRealTimeByteHitRateDataResponseBodyData {
	return s.Data
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponseBody) SetData(v *DescribeVodDomainRealTimeByteHitRateDataResponseBodyData) *DescribeVodDomainRealTimeByteHitRateDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponseBody) SetRequestId(v string) *DescribeVodDomainRealTimeByteHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainRealTimeByteHitRateDataResponseBodyData struct {
	ByteHitRateDataModel []*DescribeVodDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel `json:"ByteHitRateDataModel,omitempty" xml:"ByteHitRateDataModel,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainRealTimeByteHitRateDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeByteHitRateDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponseBodyData) GetByteHitRateDataModel() []*DescribeVodDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	return s.ByteHitRateDataModel
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponseBodyData) SetByteHitRateDataModel(v []*DescribeVodDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) *DescribeVodDomainRealTimeByteHitRateDataResponseBodyData {
	s.ByteHitRateDataModel = v
	return s
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel struct {
	// The byte hit ratio in percentage.
	//
	// example:
	//
	// 0.8956940476262277
	ByteHitRate *float32 `json:"ByteHitRate,omitempty" xml:"ByteHitRate,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-05-15T09:13:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) GetByteHitRate() *float32 {
	return s.ByteHitRate
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) SetByteHitRate(v float32) *DescribeVodDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	s.ByteHitRate = &v
	return s
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) SetTimeStamp(v string) *DescribeVodDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) Validate() error {
	return dara.Validate(s)
}
