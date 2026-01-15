// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDumpMetaListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DumpMetaListResponseBodyData) *DumpMetaListResponseBody
	GetData() *DumpMetaListResponseBodyData
	SetRequestId(v string) *DumpMetaListResponseBody
	GetRequestId() *string
}

type DumpMetaListResponseBody struct {
	// The information about the task that is used to export metadata.
	Data *DumpMetaListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B3137727-7D6E-488C-BA21-0E034C38A879
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DumpMetaListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DumpMetaListResponseBody) GoString() string {
	return s.String()
}

func (s *DumpMetaListResponseBody) GetData() *DumpMetaListResponseBodyData {
	return s.Data
}

func (s *DumpMetaListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DumpMetaListResponseBody) SetData(v *DumpMetaListResponseBodyData) *DumpMetaListResponseBody {
	s.Data = v
	return s
}

func (s *DumpMetaListResponseBody) SetRequestId(v string) *DumpMetaListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DumpMetaListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DumpMetaListResponseBodyData struct {
	// A list of tasks that are used to export metadata.
	DumpMetaList []*DumpMetaListResponseBodyDataDumpMetaList `json:"DumpMetaList,omitempty" xml:"DumpMetaList,omitempty" type:"Repeated"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DumpMetaListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DumpMetaListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DumpMetaListResponseBodyData) GetDumpMetaList() []*DumpMetaListResponseBodyDataDumpMetaList {
	return s.DumpMetaList
}

func (s *DumpMetaListResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DumpMetaListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DumpMetaListResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DumpMetaListResponseBodyData) SetDumpMetaList(v []*DumpMetaListResponseBodyDataDumpMetaList) *DumpMetaListResponseBodyData {
	s.DumpMetaList = v
	return s
}

func (s *DumpMetaListResponseBodyData) SetPageNumber(v int32) *DumpMetaListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DumpMetaListResponseBodyData) SetPageSize(v int32) *DumpMetaListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DumpMetaListResponseBodyData) SetTotalCount(v int64) *DumpMetaListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DumpMetaListResponseBodyData) Validate() error {
	if s.DumpMetaList != nil {
		for _, item := range s.DumpMetaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DumpMetaListResponseBodyDataDumpMetaList struct {
	// The error code returned.
	//
	// 	- A value of 0 indicates that the operation is successful.
	//
	// 	- Values other than 0 indicate errors.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 500
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The address where you can download the metadata. The address is valid for 2 hours.
	//
	// example:
	//
	// https://imagesearchname.oss-cn-shanghai.aliyuncs.com/xxx
	MetaUrl *string `json:"MetaUrl,omitempty" xml:"MetaUrl,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The status of the export task.
	//
	// 	- PROCESSING: in progress
	//
	// 	- FAIL: failed
	//
	// 	- SUCCESS: successful
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the task was created. Unit: milliseconds.
	//
	// example:
	//
	// 1629095713000
	UtcCreate *string `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
	// The time when the task was updated. Unit: milliseconds.
	//
	// example:
	//
	// 1629095760000
	UtcModified *int64 `json:"UtcModified,omitempty" xml:"UtcModified,omitempty"`
}

func (s DumpMetaListResponseBodyDataDumpMetaList) String() string {
	return dara.Prettify(s)
}

func (s DumpMetaListResponseBodyDataDumpMetaList) GoString() string {
	return s.String()
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) GetCode() *string {
	return s.Code
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) GetId() *int64 {
	return s.Id
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) GetMetaUrl() *string {
	return s.MetaUrl
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) GetMsg() *string {
	return s.Msg
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) GetStatus() *string {
	return s.Status
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) GetUtcCreate() *string {
	return s.UtcCreate
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) GetUtcModified() *int64 {
	return s.UtcModified
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetCode(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.Code = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetId(v int64) *DumpMetaListResponseBodyDataDumpMetaList {
	s.Id = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetMetaUrl(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.MetaUrl = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetMsg(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.Msg = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetStatus(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.Status = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetUtcCreate(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.UtcCreate = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetUtcModified(v int64) *DumpMetaListResponseBodyDataDumpMetaList {
	s.UtcModified = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) Validate() error {
	return dara.Validate(s)
}
