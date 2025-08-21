// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceHistoryEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListInstanceHistoryEventsResponseBodyHeaders) *ListInstanceHistoryEventsResponseBody
	GetHeaders() *ListInstanceHistoryEventsResponseBodyHeaders
	SetRequestId(v string) *ListInstanceHistoryEventsResponseBody
	GetRequestId() *string
	SetResult(v []*ListInstanceHistoryEventsResponseBodyResult) *ListInstanceHistoryEventsResponseBody
	GetResult() []*ListInstanceHistoryEventsResponseBodyResult
}

type ListInstanceHistoryEventsResponseBody struct {
	Headers *ListInstanceHistoryEventsResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// example:
	//
	// D1A6830A-F59B-4E05-BFAC-9496C21DBBA7
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListInstanceHistoryEventsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListInstanceHistoryEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHistoryEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceHistoryEventsResponseBody) GetHeaders() *ListInstanceHistoryEventsResponseBodyHeaders {
	return s.Headers
}

func (s *ListInstanceHistoryEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceHistoryEventsResponseBody) GetResult() []*ListInstanceHistoryEventsResponseBodyResult {
	return s.Result
}

func (s *ListInstanceHistoryEventsResponseBody) SetHeaders(v *ListInstanceHistoryEventsResponseBodyHeaders) *ListInstanceHistoryEventsResponseBody {
	s.Headers = v
	return s
}

func (s *ListInstanceHistoryEventsResponseBody) SetRequestId(v string) *ListInstanceHistoryEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceHistoryEventsResponseBody) SetResult(v []*ListInstanceHistoryEventsResponseBodyResult) *ListInstanceHistoryEventsResponseBody {
	s.Result = v
	return s
}

func (s *ListInstanceHistoryEventsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceHistoryEventsResponseBodyHeaders struct {
	// example:
	//
	// 15
	XTotalCount *int64 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
	// example:
	//
	// 2
	XTotalFailed *int64 `json:"X-Total-Failed,omitempty" xml:"X-Total-Failed,omitempty"`
	// example:
	//
	// 13
	XTotalSuccess *int64 `json:"X-Total-Success,omitempty" xml:"X-Total-Success,omitempty"`
}

func (s ListInstanceHistoryEventsResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHistoryEventsResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListInstanceHistoryEventsResponseBodyHeaders) GetXTotalCount() *int64 {
	return s.XTotalCount
}

func (s *ListInstanceHistoryEventsResponseBodyHeaders) GetXTotalFailed() *int64 {
	return s.XTotalFailed
}

func (s *ListInstanceHistoryEventsResponseBodyHeaders) GetXTotalSuccess() *int64 {
	return s.XTotalSuccess
}

func (s *ListInstanceHistoryEventsResponseBodyHeaders) SetXTotalCount(v int64) *ListInstanceHistoryEventsResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListInstanceHistoryEventsResponseBodyHeaders) SetXTotalFailed(v int64) *ListInstanceHistoryEventsResponseBodyHeaders {
	s.XTotalFailed = &v
	return s
}

func (s *ListInstanceHistoryEventsResponseBodyHeaders) SetXTotalSuccess(v int64) *ListInstanceHistoryEventsResponseBodyHeaders {
	s.XTotalSuccess = &v
	return s
}

func (s *ListInstanceHistoryEventsResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListInstanceHistoryEventsResponseBodyResult struct {
	// example:
	//
	// i-2ze8s9cjdf2cv969****
	EcsId *string `json:"ecsId,omitempty" xml:"ecsId,omitempty"`
	// example:
	//
	// 2017-12-07T00:00:00Z
	EventCreateTime *string `json:"eventCreateTime,omitempty" xml:"eventCreateTime,omitempty"`
	// example:
	//
	// EXECUTED
	EventCycleStatus *string `json:"eventCycleStatus,omitempty" xml:"eventCycleStatus,omitempty"`
	// example:
	//
	// 2017-12-07T00:00:00Z
	EventExecuteTime *string `json:"eventExecuteTime,omitempty" xml:"eventExecuteTime,omitempty"`
	// example:
	//
	// 2017-12-07T00:00:00Z
	EventFinashTime *string `json:"eventFinashTime,omitempty" xml:"eventFinashTime,omitempty"`
	// example:
	//
	// INFO
	EventLevel *string `json:"eventLevel,omitempty" xml:"eventLevel,omitempty"`
	// example:
	//
	// ECS:AUTO_RESTART
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// example:
	//
	// es-cn-2r42l7a740005****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 10.1.xx.xx
	NodeIP *string `json:"nodeIP,omitempty" xml:"nodeIP,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListInstanceHistoryEventsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHistoryEventsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInstanceHistoryEventsResponseBodyResult) GetEcsId() *string {
	return s.EcsId
}

func (s *ListInstanceHistoryEventsResponseBodyResult) GetEventCreateTime() *string {
	return s.EventCreateTime
}

func (s *ListInstanceHistoryEventsResponseBodyResult) GetEventCycleStatus() *string {
	return s.EventCycleStatus
}

func (s *ListInstanceHistoryEventsResponseBodyResult) GetEventExecuteTime() *string {
	return s.EventExecuteTime
}

func (s *ListInstanceHistoryEventsResponseBodyResult) GetEventFinashTime() *string {
	return s.EventFinashTime
}

func (s *ListInstanceHistoryEventsResponseBodyResult) GetEventLevel() *string {
	return s.EventLevel
}

func (s *ListInstanceHistoryEventsResponseBodyResult) GetEventType() *string {
	return s.EventType
}

func (s *ListInstanceHistoryEventsResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceHistoryEventsResponseBodyResult) GetNodeIP() *string {
	return s.NodeIP
}

func (s *ListInstanceHistoryEventsResponseBodyResult) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceHistoryEventsResponseBodyResult) SetEcsId(v string) *ListInstanceHistoryEventsResponseBodyResult {
	s.EcsId = &v
	return s
}

func (s *ListInstanceHistoryEventsResponseBodyResult) SetEventCreateTime(v string) *ListInstanceHistoryEventsResponseBodyResult {
	s.EventCreateTime = &v
	return s
}

func (s *ListInstanceHistoryEventsResponseBodyResult) SetEventCycleStatus(v string) *ListInstanceHistoryEventsResponseBodyResult {
	s.EventCycleStatus = &v
	return s
}

func (s *ListInstanceHistoryEventsResponseBodyResult) SetEventExecuteTime(v string) *ListInstanceHistoryEventsResponseBodyResult {
	s.EventExecuteTime = &v
	return s
}

func (s *ListInstanceHistoryEventsResponseBodyResult) SetEventFinashTime(v string) *ListInstanceHistoryEventsResponseBodyResult {
	s.EventFinashTime = &v
	return s
}

func (s *ListInstanceHistoryEventsResponseBodyResult) SetEventLevel(v string) *ListInstanceHistoryEventsResponseBodyResult {
	s.EventLevel = &v
	return s
}

func (s *ListInstanceHistoryEventsResponseBodyResult) SetEventType(v string) *ListInstanceHistoryEventsResponseBodyResult {
	s.EventType = &v
	return s
}

func (s *ListInstanceHistoryEventsResponseBodyResult) SetInstanceId(v string) *ListInstanceHistoryEventsResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceHistoryEventsResponseBodyResult) SetNodeIP(v string) *ListInstanceHistoryEventsResponseBodyResult {
	s.NodeIP = &v
	return s
}

func (s *ListInstanceHistoryEventsResponseBodyResult) SetRegionId(v string) *ListInstanceHistoryEventsResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *ListInstanceHistoryEventsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
