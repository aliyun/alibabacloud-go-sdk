// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodMultiUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeVodMultiUsageDataResponseBody
	GetEndTime() *string
	SetMultiUsageData(v *DescribeVodMultiUsageDataResponseBodyMultiUsageData) *DescribeVodMultiUsageDataResponseBody
	GetMultiUsageData() *DescribeVodMultiUsageDataResponseBodyMultiUsageData
	SetRequestId(v string) *DescribeVodMultiUsageDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodMultiUsageDataResponseBody
	GetStartTime() *string
}

type DescribeVodMultiUsageDataResponseBody struct {
	EndTime        *string                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MultiUsageData *DescribeVodMultiUsageDataResponseBodyMultiUsageData `json:"MultiUsageData,omitempty" xml:"MultiUsageData,omitempty" type:"Struct"`
	RequestId      *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime      *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodMultiUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodMultiUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodMultiUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodMultiUsageDataResponseBody) GetMultiUsageData() *DescribeVodMultiUsageDataResponseBodyMultiUsageData {
	return s.MultiUsageData
}

func (s *DescribeVodMultiUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodMultiUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodMultiUsageDataResponseBody) SetEndTime(v string) *DescribeVodMultiUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodMultiUsageDataResponseBody) SetMultiUsageData(v *DescribeVodMultiUsageDataResponseBodyMultiUsageData) *DescribeVodMultiUsageDataResponseBody {
	s.MultiUsageData = v
	return s
}

func (s *DescribeVodMultiUsageDataResponseBody) SetRequestId(v string) *DescribeVodMultiUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodMultiUsageDataResponseBody) SetStartTime(v string) *DescribeVodMultiUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodMultiUsageDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodMultiUsageDataResponseBodyMultiUsageData struct {
	DataModule []*DescribeVodMultiUsageDataResponseBodyMultiUsageDataDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodMultiUsageDataResponseBodyMultiUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodMultiUsageDataResponseBodyMultiUsageData) GoString() string {
	return s.String()
}

func (s *DescribeVodMultiUsageDataResponseBodyMultiUsageData) GetDataModule() []*DescribeVodMultiUsageDataResponseBodyMultiUsageDataDataModule {
	return s.DataModule
}

func (s *DescribeVodMultiUsageDataResponseBodyMultiUsageData) SetDataModule(v []*DescribeVodMultiUsageDataResponseBodyMultiUsageDataDataModule) *DescribeVodMultiUsageDataResponseBodyMultiUsageData {
	s.DataModule = v
	return s
}

func (s *DescribeVodMultiUsageDataResponseBodyMultiUsageData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodMultiUsageDataResponseBodyMultiUsageDataDataModule struct {
	Value     *int64  `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s DescribeVodMultiUsageDataResponseBodyMultiUsageDataDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodMultiUsageDataResponseBodyMultiUsageDataDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodMultiUsageDataResponseBodyMultiUsageDataDataModule) GetValue() *int64 {
	return s.Value
}

func (s *DescribeVodMultiUsageDataResponseBodyMultiUsageDataDataModule) GetValueType() *string {
	return s.ValueType
}

func (s *DescribeVodMultiUsageDataResponseBodyMultiUsageDataDataModule) SetValue(v int64) *DescribeVodMultiUsageDataResponseBodyMultiUsageDataDataModule {
	s.Value = &v
	return s
}

func (s *DescribeVodMultiUsageDataResponseBodyMultiUsageDataDataModule) SetValueType(v string) *DescribeVodMultiUsageDataResponseBodyMultiUsageDataDataModule {
	s.ValueType = &v
	return s
}

func (s *DescribeVodMultiUsageDataResponseBodyMultiUsageDataDataModule) Validate() error {
	return dara.Validate(s)
}
