// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChangeLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeLogs(v *DescribeChangeLogsResponseBodyChangeLogs) *DescribeChangeLogsResponseBody
	GetChangeLogs() *DescribeChangeLogsResponseBodyChangeLogs
	SetPageNumber(v int32) *DescribeChangeLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeChangeLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeChangeLogsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeChangeLogsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeChangeLogsResponseBody
	GetTotalPages() *int32
}

type DescribeChangeLogsResponseBody struct {
	ChangeLogs *DescribeChangeLogsResponseBodyChangeLogs `json:"ChangeLogs,omitempty" xml:"ChangeLogs,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F0FCB52A-D512-41A0-8595-40234EDCFD8B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 100
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeChangeLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChangeLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsResponseBody) GetChangeLogs() *DescribeChangeLogsResponseBodyChangeLogs {
	return s.ChangeLogs
}

func (s *DescribeChangeLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeChangeLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeChangeLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChangeLogsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeChangeLogsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeChangeLogsResponseBody) SetChangeLogs(v *DescribeChangeLogsResponseBodyChangeLogs) *DescribeChangeLogsResponseBody {
	s.ChangeLogs = v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetPageNumber(v int32) *DescribeChangeLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetPageSize(v int32) *DescribeChangeLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetRequestId(v string) *DescribeChangeLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetTotalItems(v int32) *DescribeChangeLogsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetTotalPages(v int32) *DescribeChangeLogsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeChangeLogsResponseBody) Validate() error {
	if s.ChangeLogs != nil {
		if err := s.ChangeLogs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeChangeLogsResponseBodyChangeLogs struct {
	ChangeLog []*DescribeChangeLogsResponseBodyChangeLogsChangeLog `json:"ChangeLog,omitempty" xml:"ChangeLog,omitempty" type:"Repeated"`
}

func (s DescribeChangeLogsResponseBodyChangeLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeChangeLogsResponseBodyChangeLogs) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsResponseBodyChangeLogs) GetChangeLog() []*DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	return s.ChangeLog
}

func (s *DescribeChangeLogsResponseBodyChangeLogs) SetChangeLog(v []*DescribeChangeLogsResponseBodyChangeLogsChangeLog) *DescribeChangeLogsResponseBodyChangeLogs {
	s.ChangeLog = v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogs) Validate() error {
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

type DescribeChangeLogsResponseBodyChangeLogsChangeLog struct {
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreatorId      *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorSubType *string `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	CreatorType    *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	CreatorUserId  *string `json:"CreatorUserId,omitempty" xml:"CreatorUserId,omitempty"`
	EntityId       *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName     *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OperAction     *string `json:"OperAction,omitempty" xml:"OperAction,omitempty"`
	OperIp         *string `json:"OperIp,omitempty" xml:"OperIp,omitempty"`
	OperObject     *string `json:"OperObject,omitempty" xml:"OperObject,omitempty"`
	OperTime       *string `json:"OperTime,omitempty" xml:"OperTime,omitempty"`
	OperTimestamp  *int64  `json:"OperTimestamp,omitempty" xml:"OperTimestamp,omitempty"`
}

func (s DescribeChangeLogsResponseBodyChangeLogsChangeLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeChangeLogsResponseBodyChangeLogsChangeLog) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) GetContent() *string {
	return s.Content
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) GetCreatorId() *string {
	return s.CreatorId
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) GetCreatorSubType() *string {
	return s.CreatorSubType
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) GetCreatorType() *string {
	return s.CreatorType
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) GetCreatorUserId() *string {
	return s.CreatorUserId
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) GetEntityId() *string {
	return s.EntityId
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) GetEntityName() *string {
	return s.EntityName
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) GetId() *int64 {
	return s.Id
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) GetOperAction() *string {
	return s.OperAction
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) GetOperIp() *string {
	return s.OperIp
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) GetOperObject() *string {
	return s.OperObject
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) GetOperTime() *string {
	return s.OperTime
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) GetOperTimestamp() *int64 {
	return s.OperTimestamp
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetContent(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.Content = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetCreatorId(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.CreatorId = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetCreatorSubType(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.CreatorSubType = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetCreatorType(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.CreatorType = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetCreatorUserId(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.CreatorUserId = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetEntityId(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.EntityId = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetEntityName(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.EntityName = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetId(v int64) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.Id = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperAction(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperAction = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperIp(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperIp = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperObject(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperObject = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperTime(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperTime = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperTimestamp(v int64) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperTimestamp = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) Validate() error {
	return dara.Validate(s)
}
