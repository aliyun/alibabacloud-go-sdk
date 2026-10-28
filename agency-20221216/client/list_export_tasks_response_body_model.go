// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExportTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListExportTasksResponseBody
	GetCode() *string
	SetData(v []*ListExportTasksResponseBodyData) *ListExportTasksResponseBody
	GetData() []*ListExportTasksResponseBodyData
	SetMessage(v string) *ListExportTasksResponseBody
	GetMessage() *string
	SetPageNo(v int32) *ListExportTasksResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListExportTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListExportTasksResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListExportTasksResponseBody
	GetTotal() *int32
}

type ListExportTasksResponseBody struct {
	// The status code of the returned result. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned list.
	Data []*ListExportTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The description of the returned result.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// asda1231as
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 12
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListExportTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExportTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListExportTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListExportTasksResponseBody) GetData() []*ListExportTasksResponseBodyData {
	return s.Data
}

func (s *ListExportTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListExportTasksResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListExportTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExportTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExportTasksResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListExportTasksResponseBody) SetCode(v string) *ListExportTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListExportTasksResponseBody) SetData(v []*ListExportTasksResponseBodyData) *ListExportTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListExportTasksResponseBody) SetMessage(v string) *ListExportTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListExportTasksResponseBody) SetPageNo(v int32) *ListExportTasksResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListExportTasksResponseBody) SetPageSize(v int32) *ListExportTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListExportTasksResponseBody) SetRequestId(v string) *ListExportTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExportTasksResponseBody) SetTotal(v int32) *ListExportTasksResponseBody {
	s.Total = &v
	return s
}

func (s *ListExportTasksResponseBody) Validate() error {
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

type ListExportTasksResponseBodyData struct {
	// The creation time. The format is YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2024-11-01 10:22:11
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// 511376624869_quota_2023-06-22_2023-12-21_all_2023122121310057
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The unique ID of the export task.
	//
	// example:
	//
	// 12355
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The display status of the task.
	//
	// example:
	//
	// Export Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status code of the task. Valid values:
	//
	// - 2: Exporting.
	//
	// - 3: Export successful.
	//
	// - -1: Export failed.
	//
	// example:
	//
	// 3
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The download URL of the file.
	//
	// example:
	//
	// //aliyun-eco-market-servic-singapore.oss-ap-southeast-1.aliyuncs.com/511376624869_quota_2023-06-22_2023-12-21_all_2023122121310057
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListExportTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListExportTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExportTasksResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListExportTasksResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *ListExportTasksResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListExportTasksResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ListExportTasksResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListExportTasksResponseBodyData) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ListExportTasksResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *ListExportTasksResponseBodyData) SetCreateTime(v string) *ListExportTasksResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListExportTasksResponseBodyData) SetFileName(v string) *ListExportTasksResponseBodyData {
	s.FileName = &v
	return s
}

func (s *ListExportTasksResponseBodyData) SetId(v int64) *ListExportTasksResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListExportTasksResponseBodyData) SetMessage(v string) *ListExportTasksResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListExportTasksResponseBodyData) SetStatus(v string) *ListExportTasksResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListExportTasksResponseBodyData) SetStatusCode(v string) *ListExportTasksResponseBodyData {
	s.StatusCode = &v
	return s
}

func (s *ListExportTasksResponseBodyData) SetUrl(v string) *ListExportTasksResponseBodyData {
	s.Url = &v
	return s
}

func (s *ListExportTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
