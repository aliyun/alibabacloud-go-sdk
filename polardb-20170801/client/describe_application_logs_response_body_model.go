// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationLogsResponseBody
	GetApplicationId() *string
	SetItems(v *DescribeApplicationLogsResponseBodyItems) *DescribeApplicationLogsResponseBody
	GetItems() *DescribeApplicationLogsResponseBodyItems
	SetPageNumber(v int32) *DescribeApplicationLogsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeApplicationLogsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeApplicationLogsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeApplicationLogsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeApplicationLogsResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string                                   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Items         *DescribeApplicationLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 6
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// CED079B7-A408-41A1-BFF1-EC608E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 9
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeApplicationLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationLogsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationLogsResponseBody) GetItems() *DescribeApplicationLogsResponseBodyItems {
	return s.Items
}

func (s *DescribeApplicationLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApplicationLogsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeApplicationLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationLogsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeApplicationLogsResponseBody) SetApplicationId(v string) *DescribeApplicationLogsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationLogsResponseBody) SetItems(v *DescribeApplicationLogsResponseBodyItems) *DescribeApplicationLogsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeApplicationLogsResponseBody) SetPageNumber(v int32) *DescribeApplicationLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApplicationLogsResponseBody) SetPageRecordCount(v int32) *DescribeApplicationLogsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeApplicationLogsResponseBody) SetRequestId(v string) *DescribeApplicationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationLogsResponseBody) SetTotalRecordCount(v int32) *DescribeApplicationLogsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeApplicationLogsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationLogsResponseBodyItems struct {
	LogRecords []*DescribeApplicationLogsResponseBodyItemsLogRecords `json:"LogRecords,omitempty" xml:"LogRecords,omitempty" type:"Repeated"`
}

func (s DescribeApplicationLogsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeApplicationLogsResponseBodyItems) GetLogRecords() []*DescribeApplicationLogsResponseBodyItemsLogRecords {
	return s.LogRecords
}

func (s *DescribeApplicationLogsResponseBodyItems) SetLogRecords(v []*DescribeApplicationLogsResponseBodyItemsLogRecords) *DescribeApplicationLogsResponseBodyItems {
	s.LogRecords = v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItems) Validate() error {
	if s.LogRecords != nil {
		for _, item := range s.LogRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationLogsResponseBodyItemsLogRecords struct {
	ComponentName  *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	ContainerName  *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Date           *string `json:"Date,omitempty" xml:"Date,omitempty"`
	FileLine       *string `json:"FileLine,omitempty" xml:"FileLine,omitempty"`
	FileName       *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FullFilePath   *string `json:"FullFilePath,omitempty" xml:"FullFilePath,omitempty"`
	Hostname       *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	LogLevelId     *int64  `json:"LogLevelId,omitempty" xml:"LogLevelId,omitempty"`
	LogLevelName   *string `json:"LogLevelName,omitempty" xml:"LogLevelName,omitempty"`
	Method         *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Runtime        *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	RuntimeVersion *string `json:"RuntimeVersion,omitempty" xml:"RuntimeVersion,omitempty"`
	Time           *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeApplicationLogsResponseBodyItemsLogRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationLogsResponseBodyItemsLogRecords) GoString() string {
	return s.String()
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetComponentName() *string {
	return s.ComponentName
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetContainerName() *string {
	return s.ContainerName
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetContent() *string {
	return s.Content
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetDate() *string {
	return s.Date
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetFileLine() *string {
	return s.FileLine
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetFileName() *string {
	return s.FileName
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetFullFilePath() *string {
	return s.FullFilePath
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetLogLevelId() *int64 {
	return s.LogLevelId
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetLogLevelName() *string {
	return s.LogLevelName
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetMethod() *string {
	return s.Method
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetRuntime() *string {
	return s.Runtime
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetRuntimeVersion() *string {
	return s.RuntimeVersion
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) GetTime() *string {
	return s.Time
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetComponentName(v string) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.ComponentName = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetContainerName(v string) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.ContainerName = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetContent(v string) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.Content = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetDate(v string) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.Date = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetFileLine(v string) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.FileLine = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetFileName(v string) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.FileName = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetFullFilePath(v string) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.FullFilePath = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetHostname(v string) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.Hostname = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetLogLevelId(v int64) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.LogLevelId = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetLogLevelName(v string) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.LogLevelName = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetMethod(v string) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.Method = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetName(v string) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.Name = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetRuntime(v string) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.Runtime = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetRuntimeVersion(v string) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.RuntimeVersion = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) SetTime(v string) *DescribeApplicationLogsResponseBodyItemsLogRecords {
	s.Time = &v
	return s
}

func (s *DescribeApplicationLogsResponseBodyItemsLogRecords) Validate() error {
	return dara.Validate(s)
}
