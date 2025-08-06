// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeVodDomainRealTimeBpsDataResponseBodyData) *DescribeVodDomainRealTimeBpsDataResponseBody
	GetData() *DescribeVodDomainRealTimeBpsDataResponseBodyData
	SetRequestId(v string) *DescribeVodDomainRealTimeBpsDataResponseBody
	GetRequestId() *string
}

type DescribeVodDomainRealTimeBpsDataResponseBody struct {
	// The returned data.
	Data *DescribeVodDomainRealTimeBpsDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B49E6DDA-F413-422B-B58E-2FA23F286726
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainRealTimeBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeBpsDataResponseBody) GetData() *DescribeVodDomainRealTimeBpsDataResponseBodyData {
	return s.Data
}

func (s *DescribeVodDomainRealTimeBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainRealTimeBpsDataResponseBody) SetData(v *DescribeVodDomainRealTimeBpsDataResponseBodyData) *DescribeVodDomainRealTimeBpsDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeVodDomainRealTimeBpsDataResponseBody) SetRequestId(v string) *DescribeVodDomainRealTimeBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainRealTimeBpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainRealTimeBpsDataResponseBodyData struct {
	BpsModel []*DescribeVodDomainRealTimeBpsDataResponseBodyDataBpsModel `json:"BpsModel,omitempty" xml:"BpsModel,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainRealTimeBpsDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeBpsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeBpsDataResponseBodyData) GetBpsModel() []*DescribeVodDomainRealTimeBpsDataResponseBodyDataBpsModel {
	return s.BpsModel
}

func (s *DescribeVodDomainRealTimeBpsDataResponseBodyData) SetBpsModel(v []*DescribeVodDomainRealTimeBpsDataResponseBodyDataBpsModel) *DescribeVodDomainRealTimeBpsDataResponseBodyData {
	s.BpsModel = v
	return s
}

func (s *DescribeVodDomainRealTimeBpsDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainRealTimeBpsDataResponseBodyDataBpsModel struct {
	// The bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 16710625.733333332
	Bps *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-11-30T05:41:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodDomainRealTimeBpsDataResponseBodyDataBpsModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeBpsDataResponseBodyDataBpsModel) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeBpsDataResponseBodyDataBpsModel) GetBps() *float32 {
	return s.Bps
}

func (s *DescribeVodDomainRealTimeBpsDataResponseBodyDataBpsModel) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainRealTimeBpsDataResponseBodyDataBpsModel) SetBps(v float32) *DescribeVodDomainRealTimeBpsDataResponseBodyDataBpsModel {
	s.Bps = &v
	return s
}

func (s *DescribeVodDomainRealTimeBpsDataResponseBodyDataBpsModel) SetTimeStamp(v string) *DescribeVodDomainRealTimeBpsDataResponseBodyDataBpsModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainRealTimeBpsDataResponseBodyDataBpsModel) Validate() error {
	return dara.Validate(s)
}
