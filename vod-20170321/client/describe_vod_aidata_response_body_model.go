// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodAIDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIData(v *DescribeVodAIDataResponseBodyAIData) *DescribeVodAIDataResponseBody
	GetAIData() *DescribeVodAIDataResponseBodyAIData
	SetDataInterval(v string) *DescribeVodAIDataResponseBody
	GetDataInterval() *string
	SetRequestId(v string) *DescribeVodAIDataResponseBody
	GetRequestId() *string
}

type DescribeVodAIDataResponseBody struct {
	// The statistics on video AI.
	AIData *DescribeVodAIDataResponseBodyAIData `json:"AIData,omitempty" xml:"AIData,omitempty" type:"Struct"`
	// The time granularity at which the data was queried. Valid values:
	//
	// 	- **hour**
	//
	// 	- **day**
	//
	// example:
	//
	// day
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C370DAF1-C838-4288-****-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodAIDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodAIDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodAIDataResponseBody) GetAIData() *DescribeVodAIDataResponseBodyAIData {
	return s.AIData
}

func (s *DescribeVodAIDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodAIDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodAIDataResponseBody) SetAIData(v *DescribeVodAIDataResponseBodyAIData) *DescribeVodAIDataResponseBody {
	s.AIData = v
	return s
}

func (s *DescribeVodAIDataResponseBody) SetDataInterval(v string) *DescribeVodAIDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodAIDataResponseBody) SetRequestId(v string) *DescribeVodAIDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodAIDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodAIDataResponseBodyAIData struct {
	AIDataItem []*DescribeVodAIDataResponseBodyAIDataAIDataItem `json:"AIDataItem,omitempty" xml:"AIDataItem,omitempty" type:"Repeated"`
}

func (s DescribeVodAIDataResponseBodyAIData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodAIDataResponseBodyAIData) GoString() string {
	return s.String()
}

func (s *DescribeVodAIDataResponseBodyAIData) GetAIDataItem() []*DescribeVodAIDataResponseBodyAIDataAIDataItem {
	return s.AIDataItem
}

func (s *DescribeVodAIDataResponseBodyAIData) SetAIDataItem(v []*DescribeVodAIDataResponseBodyAIDataAIDataItem) *DescribeVodAIDataResponseBodyAIData {
	s.AIDataItem = v
	return s
}

func (s *DescribeVodAIDataResponseBodyAIData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodAIDataResponseBodyAIDataAIDataItem struct {
	// The statistics on video AI of each type.
	Data *DescribeVodAIDataResponseBodyAIDataAIDataItemData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The timestamp of the returned data. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-02-01T13:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodAIDataResponseBodyAIDataAIDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodAIDataResponseBodyAIDataAIDataItem) GoString() string {
	return s.String()
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItem) GetData() *DescribeVodAIDataResponseBodyAIDataAIDataItemData {
	return s.Data
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItem) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItem) SetData(v *DescribeVodAIDataResponseBodyAIDataAIDataItemData) *DescribeVodAIDataResponseBodyAIDataAIDataItem {
	s.Data = v
	return s
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItem) SetTimeStamp(v string) *DescribeVodAIDataResponseBodyAIDataAIDataItem {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItem) Validate() error {
	return dara.Validate(s)
}

type DescribeVodAIDataResponseBodyAIDataAIDataItemData struct {
	DataItem []*DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem `json:"DataItem,omitempty" xml:"DataItem,omitempty" type:"Repeated"`
}

func (s DescribeVodAIDataResponseBodyAIDataAIDataItemData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodAIDataResponseBodyAIDataAIDataItemData) GoString() string {
	return s.String()
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItemData) GetDataItem() []*DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem {
	return s.DataItem
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItemData) SetDataItem(v []*DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem) *DescribeVodAIDataResponseBodyAIDataAIDataItemData {
	s.DataItem = v
	return s
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItemData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem struct {
	// The type of video AI. Valid values:
	//
	// 	- **AIVideoCensor**: automated review
	//
	// 	- **AIVideoFPShot**: media fingerprinting
	//
	// 	- **AIVideoTag**: smart tagging
	//
	// example:
	//
	// AIVideoCensor
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The processing duration. Unit: seconds.
	//
	// example:
	//
	// 111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem) GoString() string {
	return s.String()
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem) GetName() *string {
	return s.Name
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem) GetValue() *string {
	return s.Value
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem) SetName(v string) *DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem {
	s.Name = &v
	return s
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem) SetValue(v string) *DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem {
	s.Value = &v
	return s
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem) Validate() error {
	return dara.Validate(s)
}
