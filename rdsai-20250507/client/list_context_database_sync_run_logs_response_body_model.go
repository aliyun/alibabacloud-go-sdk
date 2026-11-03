// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncRunLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListContextDatabaseSyncRunLogsResponseBody
	GetCode() *string
	SetData(v *ListContextDatabaseSyncRunLogsResponseBodyData) *ListContextDatabaseSyncRunLogsResponseBody
	GetData() *ListContextDatabaseSyncRunLogsResponseBodyData
	SetHttpStatusCode(v int64) *ListContextDatabaseSyncRunLogsResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListContextDatabaseSyncRunLogsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListContextDatabaseSyncRunLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListContextDatabaseSyncRunLogsResponseBody
	GetSuccess() *bool
}

type ListContextDatabaseSyncRunLogsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Sample value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListContextDatabaseSyncRunLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message.
	//
	// example:
	//
	// Sample value
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// id-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListContextDatabaseSyncRunLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListContextDatabaseSyncRunLogsResponseBody) GetData() *ListContextDatabaseSyncRunLogsResponseBodyData {
	return s.Data
}

func (s *ListContextDatabaseSyncRunLogsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListContextDatabaseSyncRunLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContextDatabaseSyncRunLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseSyncRunLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListContextDatabaseSyncRunLogsResponseBody) SetCode(v string) *ListContextDatabaseSyncRunLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBody) SetData(v *ListContextDatabaseSyncRunLogsResponseBodyData) *ListContextDatabaseSyncRunLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBody) SetHttpStatusCode(v int64) *ListContextDatabaseSyncRunLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBody) SetMessage(v string) *ListContextDatabaseSyncRunLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBody) SetRequestId(v string) *ListContextDatabaseSyncRunLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBody) SetSuccess(v bool) *ListContextDatabaseSyncRunLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListContextDatabaseSyncRunLogsResponseBodyData struct {
	// The list of logs.
	//
	// example:
	//
	// [{"Id":1000000000000000, "Level":"info", "Message":"Operation successful", "Detail":"Synchronization completed. 12 documents uploaded."}]
	Logs []*ListContextDatabaseSyncRunLogsResponseBodyDataLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListContextDatabaseSyncRunLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyData) GetLogs() []*ListContextDatabaseSyncRunLogsResponseBodyDataLogs {
	return s.Logs
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyData) SetLogs(v []*ListContextDatabaseSyncRunLogsResponseBodyDataLogs) *ListContextDatabaseSyncRunLogsResponseBodyData {
	s.Logs = v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyData) SetPage(v int32) *ListContextDatabaseSyncRunLogsResponseBodyData {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyData) SetPageSize(v int32) *ListContextDatabaseSyncRunLogsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyData) SetTotal(v int32) *ListContextDatabaseSyncRunLogsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyData) Validate() error {
	if s.Logs != nil {
		for _, item := range s.Logs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextDatabaseSyncRunLogsResponseBodyDataLogs struct {
	// The creation time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The details.
	//
	// example:
	//
	// Synchronization completed. 12 documents uploaded
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The run ID.
	//
	// example:
	//
	// 1000000000000000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The log level. Valid values: debug, info, warn, and error.
	//
	// example:
	//
	// info
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The response message.
	//
	// example:
	//
	// Operation successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListContextDatabaseSyncRunLogsResponseBodyDataLogs) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunLogsResponseBodyDataLogs) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyDataLogs) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyDataLogs) GetDetail() *string {
	return s.Detail
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyDataLogs) GetId() *int64 {
	return s.Id
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyDataLogs) GetLevel() *string {
	return s.Level
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyDataLogs) GetMessage() *string {
	return s.Message
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyDataLogs) SetCreatedAt(v string) *ListContextDatabaseSyncRunLogsResponseBodyDataLogs {
	s.CreatedAt = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyDataLogs) SetDetail(v string) *ListContextDatabaseSyncRunLogsResponseBodyDataLogs {
	s.Detail = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyDataLogs) SetId(v int64) *ListContextDatabaseSyncRunLogsResponseBodyDataLogs {
	s.Id = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyDataLogs) SetLevel(v string) *ListContextDatabaseSyncRunLogsResponseBodyDataLogs {
	s.Level = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyDataLogs) SetMessage(v string) *ListContextDatabaseSyncRunLogsResponseBodyDataLogs {
	s.Message = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponseBodyDataLogs) Validate() error {
	return dara.Validate(s)
}
