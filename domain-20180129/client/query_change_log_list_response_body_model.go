// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChangeLogListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *QueryChangeLogListResponseBody
	GetCurrentPageNum() *int32
	SetData(v *QueryChangeLogListResponseBodyData) *QueryChangeLogListResponseBody
	GetData() *QueryChangeLogListResponseBodyData
	SetNextPage(v bool) *QueryChangeLogListResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *QueryChangeLogListResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *QueryChangeLogListResponseBody
	GetPrePage() *bool
	SetRequestId(v string) *QueryChangeLogListResponseBody
	GetRequestId() *string
	SetResultLimit(v bool) *QueryChangeLogListResponseBody
	GetResultLimit() *bool
	SetTotalItemNum(v int32) *QueryChangeLogListResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *QueryChangeLogListResponseBody
	GetTotalPageNum() *int32
}

type QueryChangeLogListResponseBody struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                              `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryChangeLogListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// true
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// false
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// example:
	//
	// 2DEDFF32-7827-46B1-BE90-3DB8ABD91A58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	ResultLimit *bool `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// example:
	//
	// 1000
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1000
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryChangeLogListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryChangeLogListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryChangeLogListResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryChangeLogListResponseBody) GetData() *QueryChangeLogListResponseBodyData {
	return s.Data
}

func (s *QueryChangeLogListResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryChangeLogListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryChangeLogListResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryChangeLogListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryChangeLogListResponseBody) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *QueryChangeLogListResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryChangeLogListResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryChangeLogListResponseBody) SetCurrentPageNum(v int32) *QueryChangeLogListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryChangeLogListResponseBody) SetData(v *QueryChangeLogListResponseBodyData) *QueryChangeLogListResponseBody {
	s.Data = v
	return s
}

func (s *QueryChangeLogListResponseBody) SetNextPage(v bool) *QueryChangeLogListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryChangeLogListResponseBody) SetPageSize(v int32) *QueryChangeLogListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryChangeLogListResponseBody) SetPrePage(v bool) *QueryChangeLogListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryChangeLogListResponseBody) SetRequestId(v string) *QueryChangeLogListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryChangeLogListResponseBody) SetResultLimit(v bool) *QueryChangeLogListResponseBody {
	s.ResultLimit = &v
	return s
}

func (s *QueryChangeLogListResponseBody) SetTotalItemNum(v int32) *QueryChangeLogListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryChangeLogListResponseBody) SetTotalPageNum(v int32) *QueryChangeLogListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryChangeLogListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryChangeLogListResponseBodyData struct {
	ChangeLog []*QueryChangeLogListResponseBodyDataChangeLog `json:"ChangeLog,omitempty" xml:"ChangeLog,omitempty" type:"Repeated"`
}

func (s QueryChangeLogListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryChangeLogListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryChangeLogListResponseBodyData) GetChangeLog() []*QueryChangeLogListResponseBodyDataChangeLog {
	return s.ChangeLog
}

func (s *QueryChangeLogListResponseBodyData) SetChangeLog(v []*QueryChangeLogListResponseBodyDataChangeLog) *QueryChangeLogListResponseBodyData {
	s.ChangeLog = v
	return s
}

func (s *QueryChangeLogListResponseBodyData) Validate() error {
	if s.ChangeLog != nil {
		for _, item := range s.ChangeLog {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryChangeLogListResponseBodyDataChangeLog struct {
	// example:
	//
	// dns1;dns2 -> dns3;dns4
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// DNS modification
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// 127.0.0.1
	OperationIPAddress *string `json:"OperationIPAddress,omitempty" xml:"OperationIPAddress,omitempty"`
	Remark             *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// Failed
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 2017-12-26 12:00:00
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryChangeLogListResponseBodyDataChangeLog) String() string {
	return dara.Prettify(s)
}

func (s QueryChangeLogListResponseBodyDataChangeLog) GoString() string {
	return s.String()
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) GetDetails() *string {
	return s.Details
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) GetOperation() *string {
	return s.Operation
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) GetOperationIPAddress() *string {
	return s.OperationIPAddress
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) GetRemark() *string {
	return s.Remark
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) GetResult() *string {
	return s.Result
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) GetTime() *string {
	return s.Time
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) SetDetails(v string) *QueryChangeLogListResponseBodyDataChangeLog {
	s.Details = &v
	return s
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) SetDomainName(v string) *QueryChangeLogListResponseBodyDataChangeLog {
	s.DomainName = &v
	return s
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) SetOperation(v string) *QueryChangeLogListResponseBodyDataChangeLog {
	s.Operation = &v
	return s
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) SetOperationIPAddress(v string) *QueryChangeLogListResponseBodyDataChangeLog {
	s.OperationIPAddress = &v
	return s
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) SetRemark(v string) *QueryChangeLogListResponseBodyDataChangeLog {
	s.Remark = &v
	return s
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) SetResult(v string) *QueryChangeLogListResponseBodyDataChangeLog {
	s.Result = &v
	return s
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) SetTime(v string) *QueryChangeLogListResponseBodyDataChangeLog {
	s.Time = &v
	return s
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) Validate() error {
	return dara.Validate(s)
}
