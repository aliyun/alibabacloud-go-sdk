// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeQpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeVodDomainRealTimeQpsDataResponseBodyData) *DescribeVodDomainRealTimeQpsDataResponseBody
	GetData() *DescribeVodDomainRealTimeQpsDataResponseBodyData
	SetRequestId(v string) *DescribeVodDomainRealTimeQpsDataResponseBody
	GetRequestId() *string
}

type DescribeVodDomainRealTimeQpsDataResponseBody struct {
	// The returned data.
	Data *DescribeVodDomainRealTimeQpsDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 32DC9806-E9F9-4490-BBDC-B3A9E32FCC1D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainRealTimeQpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeQpsDataResponseBody) GetData() *DescribeVodDomainRealTimeQpsDataResponseBodyData {
	return s.Data
}

func (s *DescribeVodDomainRealTimeQpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainRealTimeQpsDataResponseBody) SetData(v *DescribeVodDomainRealTimeQpsDataResponseBodyData) *DescribeVodDomainRealTimeQpsDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeVodDomainRealTimeQpsDataResponseBody) SetRequestId(v string) *DescribeVodDomainRealTimeQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainRealTimeQpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainRealTimeQpsDataResponseBodyData struct {
	QpsModel []*DescribeVodDomainRealTimeQpsDataResponseBodyDataQpsModel `json:"QpsModel,omitempty" xml:"QpsModel,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainRealTimeQpsDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeQpsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeQpsDataResponseBodyData) GetQpsModel() []*DescribeVodDomainRealTimeQpsDataResponseBodyDataQpsModel {
	return s.QpsModel
}

func (s *DescribeVodDomainRealTimeQpsDataResponseBodyData) SetQpsModel(v []*DescribeVodDomainRealTimeQpsDataResponseBodyDataQpsModel) *DescribeVodDomainRealTimeQpsDataResponseBodyData {
	s.QpsModel = v
	return s
}

func (s *DescribeVodDomainRealTimeQpsDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainRealTimeQpsDataResponseBodyDataQpsModel struct {
	// The number of queries per second.
	//
	// example:
	//
	// 1851.25
	Qps *float32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The timestamp of the returned data. The time follows the ISO 8601 standard. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-12-02T11:25:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodDomainRealTimeQpsDataResponseBodyDataQpsModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeQpsDataResponseBodyDataQpsModel) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeQpsDataResponseBodyDataQpsModel) GetQps() *float32 {
	return s.Qps
}

func (s *DescribeVodDomainRealTimeQpsDataResponseBodyDataQpsModel) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainRealTimeQpsDataResponseBodyDataQpsModel) SetQps(v float32) *DescribeVodDomainRealTimeQpsDataResponseBodyDataQpsModel {
	s.Qps = &v
	return s
}

func (s *DescribeVodDomainRealTimeQpsDataResponseBodyDataQpsModel) SetTimeStamp(v string) *DescribeVodDomainRealTimeQpsDataResponseBodyDataQpsModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainRealTimeQpsDataResponseBodyDataQpsModel) Validate() error {
	return dara.Validate(s)
}
