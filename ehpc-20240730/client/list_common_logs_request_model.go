// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionName(v []*string) *ListCommonLogsRequest
	GetActionName() []*string
	SetActionStatus(v string) *ListCommonLogsRequest
	GetActionStatus() *string
	SetClusterId(v string) *ListCommonLogsRequest
	GetClusterId() *string
	SetFrom(v int64) *ListCommonLogsRequest
	GetFrom() *int64
	SetIsReverse(v bool) *ListCommonLogsRequest
	GetIsReverse() *bool
	SetLogRequestId(v string) *ListCommonLogsRequest
	GetLogRequestId() *string
	SetLogType(v string) *ListCommonLogsRequest
	GetLogType() *string
	SetOperatorUid(v string) *ListCommonLogsRequest
	GetOperatorUid() *string
	SetPageNumber(v int32) *ListCommonLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCommonLogsRequest
	GetPageSize() *int32
	SetResource(v string) *ListCommonLogsRequest
	GetResource() *string
	SetTo(v int64) *ListCommonLogsRequest
	GetTo() *int64
}

type ListCommonLogsRequest struct {
	// The action types.
	ActionName []*string `json:"ActionName,omitempty" xml:"ActionName,omitempty" type:"Repeated"`
	// The action status. Logs associated with the specific action status are returned.
	//
	// Valid values:
	//
	// 	- Finished: The action is completed.
	//
	// 	- Failed: The action failed.
	//
	// 	- InProgress: The action is being performed.
	//
	// example:
	//
	// Finished
	ActionStatus *string `json:"ActionStatus,omitempty" xml:"ActionStatus,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The start time of the time range. The time is a timestamp. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1703821542
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// Specifies whether to display results in reverse order.
	//
	// Default value: true
	//
	// example:
	//
	// true
	IsReverse *bool `json:"IsReverse,omitempty" xml:"IsReverse,omitempty"`
	// The request ID of the action. Logs associated with the specific request ID are returned.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	// The log type. Logs of the specific type are returned.
	//
	// example:
	//
	// Operation
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The account ID of the operator.
	//
	// example:
	//
	// 137***
	OperatorUid *string `json:"OperatorUid,omitempty" xml:"OperatorUid,omitempty"`
	// The page number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the resource involved in the action. Logs associated with the specific resource are returned. This parameter is not recommended.
	//
	// example:
	//
	// i-abc***
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The end time of the time range. The time is a timestamp. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1703821666
	To *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s ListCommonLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCommonLogsRequest) GoString() string {
	return s.String()
}

func (s *ListCommonLogsRequest) GetActionName() []*string {
	return s.ActionName
}

func (s *ListCommonLogsRequest) GetActionStatus() *string {
	return s.ActionStatus
}

func (s *ListCommonLogsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListCommonLogsRequest) GetFrom() *int64 {
	return s.From
}

func (s *ListCommonLogsRequest) GetIsReverse() *bool {
	return s.IsReverse
}

func (s *ListCommonLogsRequest) GetLogRequestId() *string {
	return s.LogRequestId
}

func (s *ListCommonLogsRequest) GetLogType() *string {
	return s.LogType
}

func (s *ListCommonLogsRequest) GetOperatorUid() *string {
	return s.OperatorUid
}

func (s *ListCommonLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCommonLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCommonLogsRequest) GetResource() *string {
	return s.Resource
}

func (s *ListCommonLogsRequest) GetTo() *int64 {
	return s.To
}

func (s *ListCommonLogsRequest) SetActionName(v []*string) *ListCommonLogsRequest {
	s.ActionName = v
	return s
}

func (s *ListCommonLogsRequest) SetActionStatus(v string) *ListCommonLogsRequest {
	s.ActionStatus = &v
	return s
}

func (s *ListCommonLogsRequest) SetClusterId(v string) *ListCommonLogsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCommonLogsRequest) SetFrom(v int64) *ListCommonLogsRequest {
	s.From = &v
	return s
}

func (s *ListCommonLogsRequest) SetIsReverse(v bool) *ListCommonLogsRequest {
	s.IsReverse = &v
	return s
}

func (s *ListCommonLogsRequest) SetLogRequestId(v string) *ListCommonLogsRequest {
	s.LogRequestId = &v
	return s
}

func (s *ListCommonLogsRequest) SetLogType(v string) *ListCommonLogsRequest {
	s.LogType = &v
	return s
}

func (s *ListCommonLogsRequest) SetOperatorUid(v string) *ListCommonLogsRequest {
	s.OperatorUid = &v
	return s
}

func (s *ListCommonLogsRequest) SetPageNumber(v int32) *ListCommonLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCommonLogsRequest) SetPageSize(v int32) *ListCommonLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCommonLogsRequest) SetResource(v string) *ListCommonLogsRequest {
	s.Resource = &v
	return s
}

func (s *ListCommonLogsRequest) SetTo(v int64) *ListCommonLogsRequest {
	s.To = &v
	return s
}

func (s *ListCommonLogsRequest) Validate() error {
	return dara.Validate(s)
}
