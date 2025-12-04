// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCallTaskDetailResponseBody
	GetCode() *string
	SetData(v []*ListCallTaskDetailResponseBodyData) *ListCallTaskDetailResponseBody
	GetData() []*ListCallTaskDetailResponseBodyData
	SetPageNumber(v int64) *ListCallTaskDetailResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCallTaskDetailResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListCallTaskDetailResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListCallTaskDetailResponseBody
	GetTotal() *int64
	SetTotalPage(v int64) *ListCallTaskDetailResponseBody
	GetTotalPage() *int64
}

type ListCallTaskDetailResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the task.
	Data []*ListCallTaskDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D692AC3D-CBA8-417F-BEB9-5B73718922D4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of called numbers.
	//
	// example:
	//
	// 1000
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListCallTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCallTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListCallTaskDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCallTaskDetailResponseBody) GetData() []*ListCallTaskDetailResponseBodyData {
	return s.Data
}

func (s *ListCallTaskDetailResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCallTaskDetailResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCallTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCallTaskDetailResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListCallTaskDetailResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *ListCallTaskDetailResponseBody) SetCode(v string) *ListCallTaskDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListCallTaskDetailResponseBody) SetData(v []*ListCallTaskDetailResponseBodyData) *ListCallTaskDetailResponseBody {
	s.Data = v
	return s
}

func (s *ListCallTaskDetailResponseBody) SetPageNumber(v int64) *ListCallTaskDetailResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCallTaskDetailResponseBody) SetPageSize(v int64) *ListCallTaskDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCallTaskDetailResponseBody) SetRequestId(v string) *ListCallTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCallTaskDetailResponseBody) SetTotal(v int64) *ListCallTaskDetailResponseBody {
	s.Total = &v
	return s
}

func (s *ListCallTaskDetailResponseBody) SetTotalPage(v int64) *ListCallTaskDetailResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListCallTaskDetailResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCallTaskDetailResponseBodyData struct {
	// The called number.
	//
	// example:
	//
	// 1300000****
	CalledNum *string `json:"CalledNum,omitempty" xml:"CalledNum,omitempty"`
	// The calling number.
	//
	// example:
	//
	// 0571000****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// The call duration. Unit: seconds.
	//
	// example:
	//
	// 200
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is unavailable.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The task state. Valid values:
	//
	// 	- **SUCCESS**: The task was successful.
	//
	// 	- **FAIL**: The task failed.
	//
	// 	- **INIT**: The task was not started.
	//
	// example:
	//
	// FAIL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCallTaskDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCallTaskDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCallTaskDetailResponseBodyData) GetCalledNum() *string {
	return s.CalledNum
}

func (s *ListCallTaskDetailResponseBodyData) GetCaller() *string {
	return s.Caller
}

func (s *ListCallTaskDetailResponseBodyData) GetDuration() *int64 {
	return s.Duration
}

func (s *ListCallTaskDetailResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListCallTaskDetailResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListCallTaskDetailResponseBodyData) SetCalledNum(v string) *ListCallTaskDetailResponseBodyData {
	s.CalledNum = &v
	return s
}

func (s *ListCallTaskDetailResponseBodyData) SetCaller(v string) *ListCallTaskDetailResponseBodyData {
	s.Caller = &v
	return s
}

func (s *ListCallTaskDetailResponseBodyData) SetDuration(v int64) *ListCallTaskDetailResponseBodyData {
	s.Duration = &v
	return s
}

func (s *ListCallTaskDetailResponseBodyData) SetId(v int64) *ListCallTaskDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCallTaskDetailResponseBodyData) SetStatus(v string) *ListCallTaskDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListCallTaskDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
