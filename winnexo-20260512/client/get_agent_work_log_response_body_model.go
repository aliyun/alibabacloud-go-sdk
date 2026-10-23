// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentWorkLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentWorkLogResponseBody
	GetCode() *string
	SetLogs(v []*GetAgentWorkLogResponseBodyLogs) *GetAgentWorkLogResponseBody
	GetLogs() []*GetAgentWorkLogResponseBodyLogs
	SetMessage(v string) *GetAgentWorkLogResponseBody
	GetMessage() *string
	SetPage(v int64) *GetAgentWorkLogResponseBody
	GetPage() *int64
	SetPageSize(v int64) *GetAgentWorkLogResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *GetAgentWorkLogResponseBody
	GetRequestId() *string
	SetTotal(v int64) *GetAgentWorkLogResponseBody
	GetTotal() *int64
}

type GetAgentWorkLogResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The agent operation records in reverse chronological order. If no records are saved, an empty list is returned.
	//
	// example:
	//
	// []
	Logs []*GetAgentWorkLogResponseBodyLogs `json:"logs,omitempty" xml:"logs,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The current page number, starting from 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of filtered operation log entries.
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetAgentWorkLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentWorkLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentWorkLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentWorkLogResponseBody) GetLogs() []*GetAgentWorkLogResponseBodyLogs {
	return s.Logs
}

func (s *GetAgentWorkLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentWorkLogResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *GetAgentWorkLogResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetAgentWorkLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentWorkLogResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetAgentWorkLogResponseBody) SetCode(v string) *GetAgentWorkLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentWorkLogResponseBody) SetLogs(v []*GetAgentWorkLogResponseBodyLogs) *GetAgentWorkLogResponseBody {
	s.Logs = v
	return s
}

func (s *GetAgentWorkLogResponseBody) SetMessage(v string) *GetAgentWorkLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentWorkLogResponseBody) SetPage(v int64) *GetAgentWorkLogResponseBody {
	s.Page = &v
	return s
}

func (s *GetAgentWorkLogResponseBody) SetPageSize(v int64) *GetAgentWorkLogResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetAgentWorkLogResponseBody) SetRequestId(v string) *GetAgentWorkLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentWorkLogResponseBody) SetTotal(v int64) *GetAgentWorkLogResponseBody {
	s.Total = &v
	return s
}

func (s *GetAgentWorkLogResponseBody) Validate() error {
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

type GetAgentWorkLogResponseBodyLogs struct {
	// The record type. Valid values: THINK, REPLY, TOOL_CALL, SKILL_CALL, and ARTIFACT.
	//
	// example:
	//
	// REPLY
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The saved raw record content. Tool invocations and artifact metadata are represented as JSON strings.
	//
	// example:
	//
	// The data query for the current round is completed
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// The duration of this operation record, in milliseconds. The value is 0 for artifact records.
	//
	// example:
	//
	// 1200
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The key of the node to which the record belongs. This is not the session ID.
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The time when the event was recorded, in ISO 8601 format. If the source record has no time, an empty string is returned.
	//
	// example:
	//
	// 2026-10-14T08:00:00Z
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetAgentWorkLogResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s GetAgentWorkLogResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *GetAgentWorkLogResponseBodyLogs) GetAction() *string {
	return s.Action
}

func (s *GetAgentWorkLogResponseBodyLogs) GetDetail() *string {
	return s.Detail
}

func (s *GetAgentWorkLogResponseBodyLogs) GetDuration() *int64 {
	return s.Duration
}

func (s *GetAgentWorkLogResponseBodyLogs) GetNodeId() *string {
	return s.NodeId
}

func (s *GetAgentWorkLogResponseBodyLogs) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetAgentWorkLogResponseBodyLogs) SetAction(v string) *GetAgentWorkLogResponseBodyLogs {
	s.Action = &v
	return s
}

func (s *GetAgentWorkLogResponseBodyLogs) SetDetail(v string) *GetAgentWorkLogResponseBodyLogs {
	s.Detail = &v
	return s
}

func (s *GetAgentWorkLogResponseBodyLogs) SetDuration(v int64) *GetAgentWorkLogResponseBodyLogs {
	s.Duration = &v
	return s
}

func (s *GetAgentWorkLogResponseBodyLogs) SetNodeId(v string) *GetAgentWorkLogResponseBodyLogs {
	s.NodeId = &v
	return s
}

func (s *GetAgentWorkLogResponseBodyLogs) SetTimestamp(v string) *GetAgentWorkLogResponseBodyLogs {
	s.Timestamp = &v
	return s
}

func (s *GetAgentWorkLogResponseBodyLogs) Validate() error {
	return dara.Validate(s)
}
