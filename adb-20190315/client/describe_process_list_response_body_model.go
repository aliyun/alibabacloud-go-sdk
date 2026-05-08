// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeProcessListResponseBodyItems) *DescribeProcessListResponseBody
	GetItems() *DescribeProcessListResponseBodyItems
	SetPageNumber(v string) *DescribeProcessListResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeProcessListResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeProcessListResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeProcessListResponseBody
	GetTotalCount() *string
}

type DescribeProcessListResponseBody struct {
	Items *DescribeProcessListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProcessListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBody) GetItems() *DescribeProcessListResponseBodyItems {
	return s.Items
}

func (s *DescribeProcessListResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeProcessListResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeProcessListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProcessListResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeProcessListResponseBody) SetItems(v *DescribeProcessListResponseBodyItems) *DescribeProcessListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeProcessListResponseBody) SetPageNumber(v string) *DescribeProcessListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessListResponseBody) SetPageSize(v string) *DescribeProcessListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessListResponseBody) SetRequestId(v string) *DescribeProcessListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProcessListResponseBody) SetTotalCount(v string) *DescribeProcessListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeProcessListResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeProcessListResponseBodyItems struct {
	Process []*DescribeProcessListResponseBodyItemsProcess `json:"Process,omitempty" xml:"Process,omitempty" type:"Repeated"`
}

func (s DescribeProcessListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyItems) GetProcess() []*DescribeProcessListResponseBodyItemsProcess {
	return s.Process
}

func (s *DescribeProcessListResponseBodyItems) SetProcess(v []*DescribeProcessListResponseBodyItemsProcess) *DescribeProcessListResponseBodyItems {
	s.Process = v
	return s
}

func (s *DescribeProcessListResponseBodyItems) Validate() error {
	if s.Process != nil {
		for _, item := range s.Process {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeProcessListResponseBodyItemsProcess struct {
	Command   *string `json:"Command,omitempty" xml:"Command,omitempty"`
	DB        *string `json:"DB,omitempty" xml:"DB,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Info      *string `json:"Info,omitempty" xml:"Info,omitempty"`
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Time      *int32  `json:"Time,omitempty" xml:"Time,omitempty"`
	User      *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeProcessListResponseBodyItemsProcess) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListResponseBodyItemsProcess) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyItemsProcess) GetCommand() *string {
	return s.Command
}

func (s *DescribeProcessListResponseBodyItemsProcess) GetDB() *string {
	return s.DB
}

func (s *DescribeProcessListResponseBodyItemsProcess) GetHost() *string {
	return s.Host
}

func (s *DescribeProcessListResponseBodyItemsProcess) GetId() *int32 {
	return s.Id
}

func (s *DescribeProcessListResponseBodyItemsProcess) GetInfo() *string {
	return s.Info
}

func (s *DescribeProcessListResponseBodyItemsProcess) GetProcessId() *string {
	return s.ProcessId
}

func (s *DescribeProcessListResponseBodyItemsProcess) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeProcessListResponseBodyItemsProcess) GetTime() *int32 {
	return s.Time
}

func (s *DescribeProcessListResponseBodyItemsProcess) GetUser() *string {
	return s.User
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetCommand(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.Command = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetDB(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.DB = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetHost(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.Host = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetId(v int32) *DescribeProcessListResponseBodyItemsProcess {
	s.Id = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetInfo(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.Info = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetProcessId(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.ProcessId = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetStartTime(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.StartTime = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetTime(v int32) *DescribeProcessListResponseBodyItemsProcess {
	s.Time = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetUser(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.User = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) Validate() error {
	return dara.Validate(s)
}
