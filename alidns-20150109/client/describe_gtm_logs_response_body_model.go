// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v *DescribeGtmLogsResponseBodyLogs) *DescribeGtmLogsResponseBody
	GetLogs() *DescribeGtmLogsResponseBodyLogs
	SetPageNumber(v int32) *DescribeGtmLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGtmLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeGtmLogsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeGtmLogsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeGtmLogsResponseBody
	GetTotalPages() *int32
}

type DescribeGtmLogsResponseBody struct {
	// The list of logs returned.
	Logs *DescribeGtmLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 50C60A29-2E93-425A-ABA8-068686E28873
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned on all pages.
	//
	// example:
	//
	// 224
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 224
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeGtmLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmLogsResponseBody) GetLogs() *DescribeGtmLogsResponseBodyLogs {
	return s.Logs
}

func (s *DescribeGtmLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGtmLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGtmLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmLogsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeGtmLogsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeGtmLogsResponseBody) SetLogs(v *DescribeGtmLogsResponseBodyLogs) *DescribeGtmLogsResponseBody {
	s.Logs = v
	return s
}

func (s *DescribeGtmLogsResponseBody) SetPageNumber(v int32) *DescribeGtmLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmLogsResponseBody) SetPageSize(v int32) *DescribeGtmLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmLogsResponseBody) SetRequestId(v string) *DescribeGtmLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmLogsResponseBody) SetTotalItems(v int32) *DescribeGtmLogsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeGtmLogsResponseBody) SetTotalPages(v int32) *DescribeGtmLogsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeGtmLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmLogsResponseBodyLogs struct {
	Log []*DescribeGtmLogsResponseBodyLogsLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Repeated"`
}

func (s DescribeGtmLogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *DescribeGtmLogsResponseBodyLogs) GetLog() []*DescribeGtmLogsResponseBodyLogsLog {
	return s.Log
}

func (s *DescribeGtmLogsResponseBodyLogs) SetLog(v []*DescribeGtmLogsResponseBodyLogsLog) *DescribeGtmLogsResponseBodyLogs {
	s.Log = v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogs) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmLogsResponseBodyLogsLog struct {
	// The formatted message content.
	//
	// example:
	//
	// addtest-pool-1
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the object that was operated on.
	//
	// example:
	//
	// 121212
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The name of the object that was operated on.
	//
	// example:
	//
	// test-pool-1
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The type of the object that was operated on.
	//
	// example:
	//
	// POOL
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The ID of the log record.
	//
	// example:
	//
	// 6726
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The operation performed.
	//
	// example:
	//
	// add
	OperAction *string `json:"OperAction,omitempty" xml:"OperAction,omitempty"`
	// The IP address subject to the operation.
	//
	// example:
	//
	// 106.11.34.X
	OperIp *string `json:"OperIp,omitempty" xml:"OperIp,omitempty"`
	// The time when the operation was performed.
	//
	// example:
	//
	// 2018-01-24T07:35Z
	OperTime *string `json:"OperTime,omitempty" xml:"OperTime,omitempty"`
	// A timestamp that indicates the time when the operation was performed.
	//
	// example:
	//
	// 1516779348000
	OperTimestamp *int64 `json:"OperTimestamp,omitempty" xml:"OperTimestamp,omitempty"`
}

func (s DescribeGtmLogsResponseBodyLogsLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmLogsResponseBodyLogsLog) GoString() string {
	return s.String()
}

func (s *DescribeGtmLogsResponseBodyLogsLog) GetContent() *string {
	return s.Content
}

func (s *DescribeGtmLogsResponseBodyLogsLog) GetEntityId() *string {
	return s.EntityId
}

func (s *DescribeGtmLogsResponseBodyLogsLog) GetEntityName() *string {
	return s.EntityName
}

func (s *DescribeGtmLogsResponseBodyLogsLog) GetEntityType() *string {
	return s.EntityType
}

func (s *DescribeGtmLogsResponseBodyLogsLog) GetId() *int64 {
	return s.Id
}

func (s *DescribeGtmLogsResponseBodyLogsLog) GetOperAction() *string {
	return s.OperAction
}

func (s *DescribeGtmLogsResponseBodyLogsLog) GetOperIp() *string {
	return s.OperIp
}

func (s *DescribeGtmLogsResponseBodyLogsLog) GetOperTime() *string {
	return s.OperTime
}

func (s *DescribeGtmLogsResponseBodyLogsLog) GetOperTimestamp() *int64 {
	return s.OperTimestamp
}

func (s *DescribeGtmLogsResponseBodyLogsLog) SetContent(v string) *DescribeGtmLogsResponseBodyLogsLog {
	s.Content = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogsLog) SetEntityId(v string) *DescribeGtmLogsResponseBodyLogsLog {
	s.EntityId = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogsLog) SetEntityName(v string) *DescribeGtmLogsResponseBodyLogsLog {
	s.EntityName = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogsLog) SetEntityType(v string) *DescribeGtmLogsResponseBodyLogsLog {
	s.EntityType = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogsLog) SetId(v int64) *DescribeGtmLogsResponseBodyLogsLog {
	s.Id = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogsLog) SetOperAction(v string) *DescribeGtmLogsResponseBodyLogsLog {
	s.OperAction = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogsLog) SetOperIp(v string) *DescribeGtmLogsResponseBodyLogsLog {
	s.OperIp = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogsLog) SetOperTime(v string) *DescribeGtmLogsResponseBodyLogsLog {
	s.OperTime = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogsLog) SetOperTimestamp(v int64) *DescribeGtmLogsResponseBodyLogsLog {
	s.OperTimestamp = &v
	return s
}

func (s *DescribeGtmLogsResponseBodyLogsLog) Validate() error {
	return dara.Validate(s)
}
