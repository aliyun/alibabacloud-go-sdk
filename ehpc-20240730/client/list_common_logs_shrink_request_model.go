// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonLogsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionNameShrink(v string) *ListCommonLogsShrinkRequest
	GetActionNameShrink() *string
	SetActionStatus(v string) *ListCommonLogsShrinkRequest
	GetActionStatus() *string
	SetClusterId(v string) *ListCommonLogsShrinkRequest
	GetClusterId() *string
	SetFrom(v int64) *ListCommonLogsShrinkRequest
	GetFrom() *int64
	SetIsReverse(v bool) *ListCommonLogsShrinkRequest
	GetIsReverse() *bool
	SetLogRequestId(v string) *ListCommonLogsShrinkRequest
	GetLogRequestId() *string
	SetLogType(v string) *ListCommonLogsShrinkRequest
	GetLogType() *string
	SetOperatorUid(v string) *ListCommonLogsShrinkRequest
	GetOperatorUid() *string
	SetPageNumber(v int32) *ListCommonLogsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCommonLogsShrinkRequest
	GetPageSize() *int32
	SetResource(v string) *ListCommonLogsShrinkRequest
	GetResource() *string
	SetTo(v int64) *ListCommonLogsShrinkRequest
	GetTo() *int64
}

type ListCommonLogsShrinkRequest struct {
	// The action types.
	ActionNameShrink *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
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

func (s ListCommonLogsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCommonLogsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCommonLogsShrinkRequest) GetActionNameShrink() *string {
	return s.ActionNameShrink
}

func (s *ListCommonLogsShrinkRequest) GetActionStatus() *string {
	return s.ActionStatus
}

func (s *ListCommonLogsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListCommonLogsShrinkRequest) GetFrom() *int64 {
	return s.From
}

func (s *ListCommonLogsShrinkRequest) GetIsReverse() *bool {
	return s.IsReverse
}

func (s *ListCommonLogsShrinkRequest) GetLogRequestId() *string {
	return s.LogRequestId
}

func (s *ListCommonLogsShrinkRequest) GetLogType() *string {
	return s.LogType
}

func (s *ListCommonLogsShrinkRequest) GetOperatorUid() *string {
	return s.OperatorUid
}

func (s *ListCommonLogsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCommonLogsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCommonLogsShrinkRequest) GetResource() *string {
	return s.Resource
}

func (s *ListCommonLogsShrinkRequest) GetTo() *int64 {
	return s.To
}

func (s *ListCommonLogsShrinkRequest) SetActionNameShrink(v string) *ListCommonLogsShrinkRequest {
	s.ActionNameShrink = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetActionStatus(v string) *ListCommonLogsShrinkRequest {
	s.ActionStatus = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetClusterId(v string) *ListCommonLogsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetFrom(v int64) *ListCommonLogsShrinkRequest {
	s.From = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetIsReverse(v bool) *ListCommonLogsShrinkRequest {
	s.IsReverse = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetLogRequestId(v string) *ListCommonLogsShrinkRequest {
	s.LogRequestId = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetLogType(v string) *ListCommonLogsShrinkRequest {
	s.LogType = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetOperatorUid(v string) *ListCommonLogsShrinkRequest {
	s.OperatorUid = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetPageNumber(v int32) *ListCommonLogsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetPageSize(v int32) *ListCommonLogsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetResource(v string) *ListCommonLogsShrinkRequest {
	s.Resource = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetTo(v int64) *ListCommonLogsShrinkRequest {
	s.To = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
