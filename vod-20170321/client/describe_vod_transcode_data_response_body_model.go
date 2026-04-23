// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTranscodeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodTranscodeDataResponseBody
	GetDataInterval() *string
	SetRequestId(v string) *DescribeVodTranscodeDataResponseBody
	GetRequestId() *string
	SetTranscodeData(v *DescribeVodTranscodeDataResponseBodyTranscodeData) *DescribeVodTranscodeDataResponseBody
	GetTranscodeData() *DescribeVodTranscodeDataResponseBodyTranscodeData
}

type DescribeVodTranscodeDataResponseBody struct {
	// The interval at which the data was queried. Valid values:
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
	RequestId     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranscodeData *DescribeVodTranscodeDataResponseBodyTranscodeData `json:"TranscodeData,omitempty" xml:"TranscodeData,omitempty" type:"Struct"`
}

func (s DescribeVodTranscodeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTranscodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodTranscodeDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodTranscodeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodTranscodeDataResponseBody) GetTranscodeData() *DescribeVodTranscodeDataResponseBodyTranscodeData {
	return s.TranscodeData
}

func (s *DescribeVodTranscodeDataResponseBody) SetDataInterval(v string) *DescribeVodTranscodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodTranscodeDataResponseBody) SetRequestId(v string) *DescribeVodTranscodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodTranscodeDataResponseBody) SetTranscodeData(v *DescribeVodTranscodeDataResponseBodyTranscodeData) *DescribeVodTranscodeDataResponseBody {
	s.TranscodeData = v
	return s
}

func (s *DescribeVodTranscodeDataResponseBody) Validate() error {
	if s.TranscodeData != nil {
		if err := s.TranscodeData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodTranscodeDataResponseBodyTranscodeData struct {
	TranscodeDataItem []*DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem `json:"TranscodeDataItem,omitempty" xml:"TranscodeDataItem,omitempty" type:"Repeated"`
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeData) GoString() string {
	return s.String()
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeData) GetTranscodeDataItem() []*DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem {
	return s.TranscodeDataItem
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeData) SetTranscodeDataItem(v []*DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem) *DescribeVodTranscodeDataResponseBodyTranscodeData {
	s.TranscodeDataItem = v
	return s
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeData) Validate() error {
	if s.TranscodeDataItem != nil {
		for _, item := range s.TranscodeDataItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem struct {
	Data      *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	TimeStamp *string                                                                 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem) GoString() string {
	return s.String()
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem) GetData() *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData {
	return s.Data
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem) SetData(v *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData) *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem {
	s.Data = v
	return s
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem) SetTimeStamp(v string) *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData struct {
	DataItem []*DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem `json:"DataItem,omitempty" xml:"DataItem,omitempty" type:"Repeated"`
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData) GoString() string {
	return s.String()
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData) GetDataItem() []*DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem {
	return s.DataItem
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData) SetDataItem(v []*DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem) *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData {
	s.DataItem = v
	return s
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData) Validate() error {
	if s.DataItem != nil {
		for _, item := range s.DataItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem) GoString() string {
	return s.String()
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem) GetName() *string {
	return s.Name
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem) GetValue() *string {
	return s.Value
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem) SetName(v string) *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem {
	s.Name = &v
	return s
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem) SetValue(v string) *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem {
	s.Value = &v
	return s
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem) Validate() error {
	return dara.Validate(s)
}
