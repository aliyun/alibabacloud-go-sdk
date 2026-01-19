// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterTaskLogFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeAIDBClusterTaskLogFilesResponseBody
	GetEndTime() *string
	SetItems(v *DescribeAIDBClusterTaskLogFilesResponseBodyItems) *DescribeAIDBClusterTaskLogFilesResponseBody
	GetItems() *DescribeAIDBClusterTaskLogFilesResponseBodyItems
	SetPageNumber(v int64) *DescribeAIDBClusterTaskLogFilesResponseBody
	GetPageNumber() *int64
	SetPageRecordCount(v int32) *DescribeAIDBClusterTaskLogFilesResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int64) *DescribeAIDBClusterTaskLogFilesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeAIDBClusterTaskLogFilesResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeAIDBClusterTaskLogFilesResponseBody
	GetStartTime() *string
}

type DescribeAIDBClusterTaskLogFilesResponseBody struct {
	// example:
	//
	// 2026-01-15T15:00:00Z
	EndTime *string                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Items   *DescribeAIDBClusterTaskLogFilesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 7
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7E2FE3BB-C677-5FF9-9FC5-9CF364BD6BE5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2026-01-15T14:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAIDBClusterTaskLogFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskLogFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) GetItems() *DescribeAIDBClusterTaskLogFilesResponseBodyItems {
	return s.Items
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) SetEndTime(v string) *DescribeAIDBClusterTaskLogFilesResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) SetItems(v *DescribeAIDBClusterTaskLogFilesResponseBodyItems) *DescribeAIDBClusterTaskLogFilesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) SetPageNumber(v int64) *DescribeAIDBClusterTaskLogFilesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) SetPageRecordCount(v int32) *DescribeAIDBClusterTaskLogFilesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) SetPageSize(v int64) *DescribeAIDBClusterTaskLogFilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) SetRequestId(v string) *DescribeAIDBClusterTaskLogFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) SetStartTime(v string) *DescribeAIDBClusterTaskLogFilesResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAIDBClusterTaskLogFilesResponseBodyItems struct {
	SlsLogItems []*DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems `json:"SlsLogItems,omitempty" xml:"SlsLogItems,omitempty" type:"Repeated"`
}

func (s DescribeAIDBClusterTaskLogFilesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskLogFilesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBodyItems) GetSlsLogItems() []*DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems {
	return s.SlsLogItems
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBodyItems) SetSlsLogItems(v []*DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems) *DescribeAIDBClusterTaskLogFilesResponseBodyItems {
	s.SlsLogItems = v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBodyItems) Validate() error {
	if s.SlsLogItems != nil {
		for _, item := range s.SlsLogItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems struct {
	// example:
	//
	// 2026-01-15T14:13:50.830295892Z
	LogTime *string `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1765677660
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems) GetLogTime() *string {
	return s.LogTime
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems) GetMessage() *string {
	return s.Message
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems) GetTimestamp() *string {
	return s.Timestamp
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems) SetLogTime(v string) *DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems {
	s.LogTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems) SetMessage(v string) *DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems {
	s.Message = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems) SetTimestamp(v string) *DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems {
	s.Timestamp = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesResponseBodyItemsSlsLogItems) Validate() error {
	return dara.Validate(s)
}
