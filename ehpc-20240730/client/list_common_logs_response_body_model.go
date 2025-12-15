// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*ListCommonLogsResponseBodyLogs) *ListCommonLogsResponseBody
	GetLogs() []*ListCommonLogsResponseBodyLogs
	SetPageNumber(v int32) *ListCommonLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCommonLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCommonLogsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCommonLogsResponseBody
	GetTotalCount() *int32
	SetUid(v string) *ListCommonLogsResponseBody
	GetUid() *string
}

type ListCommonLogsResponseBody struct {
	// The brief information of operation logs.
	Logs []*ListCommonLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 464E9919-D04F-4D1D-B375-15989492****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 137***
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListCommonLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCommonLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommonLogsResponseBody) GetLogs() []*ListCommonLogsResponseBodyLogs {
	return s.Logs
}

func (s *ListCommonLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCommonLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCommonLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCommonLogsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCommonLogsResponseBody) GetUid() *string {
	return s.Uid
}

func (s *ListCommonLogsResponseBody) SetLogs(v []*ListCommonLogsResponseBodyLogs) *ListCommonLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListCommonLogsResponseBody) SetPageNumber(v int32) *ListCommonLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCommonLogsResponseBody) SetPageSize(v int32) *ListCommonLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCommonLogsResponseBody) SetRequestId(v string) *ListCommonLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommonLogsResponseBody) SetTotalCount(v int32) *ListCommonLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCommonLogsResponseBody) SetUid(v string) *ListCommonLogsResponseBody {
	s.Uid = &v
	return s
}

func (s *ListCommonLogsResponseBody) Validate() error {
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

type ListCommonLogsResponseBodyLogs struct {
	// The name of the action corresponding to the log.
	//
	// example:
	//
	// CreaterCluster
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-9T3xPNezoS
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The log type.
	//
	// example:
	//
	// Operation
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The message of the log.
	//
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The account ID of the operator.
	//
	// example:
	//
	// 137***
	OperatorUid *string `json:"OperatorUid,omitempty" xml:"OperatorUid,omitempty"`
	// The request ID associated with the action that generated the log.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The action state in the log. Valid values:
	//
	// 	- InProgress: The action is being performed.
	//
	// 	- Finished: The action is completed.
	//
	// 	- Failed: The action failed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The involved resource.
	//
	// example:
	//
	// i-abc***
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The time when the log was generated.
	//
	// example:
	//
	// 2024-08-22 14:21:54
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s ListCommonLogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s ListCommonLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListCommonLogsResponseBodyLogs) GetAction() *string {
	return s.Action
}

func (s *ListCommonLogsResponseBodyLogs) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListCommonLogsResponseBodyLogs) GetLogType() *string {
	return s.LogType
}

func (s *ListCommonLogsResponseBodyLogs) GetMessage() *string {
	return s.Message
}

func (s *ListCommonLogsResponseBodyLogs) GetOperatorUid() *string {
	return s.OperatorUid
}

func (s *ListCommonLogsResponseBodyLogs) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCommonLogsResponseBodyLogs) GetStatus() *string {
	return s.Status
}

func (s *ListCommonLogsResponseBodyLogs) GetTarget() *string {
	return s.Target
}

func (s *ListCommonLogsResponseBodyLogs) GetTime() *string {
	return s.Time
}

func (s *ListCommonLogsResponseBodyLogs) SetAction(v string) *ListCommonLogsResponseBodyLogs {
	s.Action = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetClusterId(v string) *ListCommonLogsResponseBodyLogs {
	s.ClusterId = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetLogType(v string) *ListCommonLogsResponseBodyLogs {
	s.LogType = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetMessage(v string) *ListCommonLogsResponseBodyLogs {
	s.Message = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetOperatorUid(v string) *ListCommonLogsResponseBodyLogs {
	s.OperatorUid = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetRequestId(v string) *ListCommonLogsResponseBodyLogs {
	s.RequestId = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetStatus(v string) *ListCommonLogsResponseBodyLogs {
	s.Status = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetTarget(v string) *ListCommonLogsResponseBodyLogs {
	s.Target = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetTime(v string) *ListCommonLogsResponseBodyLogs {
	s.Time = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) Validate() error {
	return dara.Validate(s)
}
