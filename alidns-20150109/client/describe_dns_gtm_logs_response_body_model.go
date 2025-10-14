// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v *DescribeDnsGtmLogsResponseBodyLogs) *DescribeDnsGtmLogsResponseBody
	GetLogs() *DescribeDnsGtmLogsResponseBodyLogs
	SetPageNumber(v int32) *DescribeDnsGtmLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDnsGtmLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDnsGtmLogsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeDnsGtmLogsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeDnsGtmLogsResponseBody
	GetTotalPages() *int32
}

type DescribeDnsGtmLogsResponseBody struct {
	// The returned logs.
	Logs *DescribeDnsGtmLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Struct"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
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
	// 1
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeDnsGtmLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmLogsResponseBody) GetLogs() *DescribeDnsGtmLogsResponseBodyLogs {
	return s.Logs
}

func (s *DescribeDnsGtmLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDnsGtmLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDnsGtmLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmLogsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeDnsGtmLogsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeDnsGtmLogsResponseBody) SetLogs(v *DescribeDnsGtmLogsResponseBodyLogs) *DescribeDnsGtmLogsResponseBody {
	s.Logs = v
	return s
}

func (s *DescribeDnsGtmLogsResponseBody) SetPageNumber(v int32) *DescribeDnsGtmLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBody) SetPageSize(v int32) *DescribeDnsGtmLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBody) SetRequestId(v string) *DescribeDnsGtmLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBody) SetTotalItems(v int32) *DescribeDnsGtmLogsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBody) SetTotalPages(v int32) *DescribeDnsGtmLogsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBody) Validate() error {
	if s.Logs != nil {
		if err := s.Logs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmLogsResponseBodyLogs struct {
	Log []*DescribeDnsGtmLogsResponseBodyLogsLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmLogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmLogsResponseBodyLogs) GetLog() []*DescribeDnsGtmLogsResponseBodyLogsLog {
	return s.Log
}

func (s *DescribeDnsGtmLogsResponseBodyLogs) SetLog(v []*DescribeDnsGtmLogsResponseBodyLogsLog) *DescribeDnsGtmLogsResponseBodyLogs {
	s.Log = v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogs) Validate() error {
	if s.Log != nil {
		for _, item := range s.Log {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmLogsResponseBodyLogsLog struct {
	// The formatted message content.
	//
	// example:
	//
	// addtest-pool-1
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the object on which the operation was performed.
	//
	// example:
	//
	// 121212
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The name of the object on which the operation was performed.
	//
	// example:
	//
	// test-pool-1
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The type of the object on which the operation was performed.
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The ID of the record.
	//
	// example:
	//
	// 6726
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The operation performed.
	OperAction *string `json:"OperAction,omitempty" xml:"OperAction,omitempty"`
	// The time when the operation was performed.
	//
	// example:
	//
	// 2018-01-24T07:35Z
	OperTime *string `json:"OperTime,omitempty" xml:"OperTime,omitempty"`
	// The timestamp of the operation.
	//
	// example:
	//
	// 1516779348000
	OperTimestamp *int64 `json:"OperTimestamp,omitempty" xml:"OperTimestamp,omitempty"`
}

func (s DescribeDnsGtmLogsResponseBodyLogsLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmLogsResponseBodyLogsLog) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) GetContent() *string {
	return s.Content
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) GetEntityId() *string {
	return s.EntityId
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) GetEntityName() *string {
	return s.EntityName
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) GetEntityType() *string {
	return s.EntityType
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) GetId() *int64 {
	return s.Id
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) GetOperAction() *string {
	return s.OperAction
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) GetOperTime() *string {
	return s.OperTime
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) GetOperTimestamp() *int64 {
	return s.OperTimestamp
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) SetContent(v string) *DescribeDnsGtmLogsResponseBodyLogsLog {
	s.Content = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) SetEntityId(v string) *DescribeDnsGtmLogsResponseBodyLogsLog {
	s.EntityId = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) SetEntityName(v string) *DescribeDnsGtmLogsResponseBodyLogsLog {
	s.EntityName = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) SetEntityType(v string) *DescribeDnsGtmLogsResponseBodyLogsLog {
	s.EntityType = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) SetId(v int64) *DescribeDnsGtmLogsResponseBodyLogsLog {
	s.Id = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) SetOperAction(v string) *DescribeDnsGtmLogsResponseBodyLogsLog {
	s.OperAction = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) SetOperTime(v string) *DescribeDnsGtmLogsResponseBodyLogsLog {
	s.OperTime = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) SetOperTimestamp(v int64) *DescribeDnsGtmLogsResponseBodyLogsLog {
	s.OperTimestamp = &v
	return s
}

func (s *DescribeDnsGtmLogsResponseBodyLogsLog) Validate() error {
	return dara.Validate(s)
}
