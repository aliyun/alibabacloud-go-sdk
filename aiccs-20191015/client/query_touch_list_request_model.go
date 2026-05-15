// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTouchListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v []*string) *QueryTouchListRequest
	GetChannelId() []*string
	SetChannelType(v []*int32) *QueryTouchListRequest
	GetChannelType() []*int32
	SetCloseTimeEnd(v int64) *QueryTouchListRequest
	GetCloseTimeEnd() *int64
	SetCloseTimeStart(v int64) *QueryTouchListRequest
	GetCloseTimeStart() *int64
	SetCurrentPage(v int32) *QueryTouchListRequest
	GetCurrentPage() *int32
	SetEvaluationLevel(v []*int32) *QueryTouchListRequest
	GetEvaluationLevel() []*int32
	SetEvaluationScore(v []*int32) *QueryTouchListRequest
	GetEvaluationScore() []*int32
	SetEvaluationStatus(v []*int32) *QueryTouchListRequest
	GetEvaluationStatus() []*int32
	SetFirstTimeEnd(v int64) *QueryTouchListRequest
	GetFirstTimeEnd() *int64
	SetFirstTimeStart(v int64) *QueryTouchListRequest
	GetFirstTimeStart() *int64
	SetInstanceId(v string) *QueryTouchListRequest
	GetInstanceId() *string
	SetMemberId(v []*int64) *QueryTouchListRequest
	GetMemberId() []*int64
	SetMemberName(v []*string) *QueryTouchListRequest
	GetMemberName() []*string
	SetPageSize(v int32) *QueryTouchListRequest
	GetPageSize() *int32
	SetQueueId(v []*int64) *QueryTouchListRequest
	GetQueueId() []*int64
	SetServicerId(v []*int64) *QueryTouchListRequest
	GetServicerId() []*int64
	SetServicerName(v []*string) *QueryTouchListRequest
	GetServicerName() []*string
	SetTouchId(v []*int64) *QueryTouchListRequest
	GetTouchId() []*int64
	SetTouchType(v []*int32) *QueryTouchListRequest
	GetTouchType() []*int32
}

type QueryTouchListRequest struct {
	ChannelId   []*string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty" type:"Repeated"`
	ChannelType []*int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty" type:"Repeated"`
	// example:
	//
	// 1614600500000
	CloseTimeEnd *int64 `json:"CloseTimeEnd,omitempty" xml:"CloseTimeEnd,omitempty"`
	// example:
	//
	// 1614600400000
	CloseTimeStart *int64 `json:"CloseTimeStart,omitempty" xml:"CloseTimeStart,omitempty"`
	// example:
	//
	// 1
	CurrentPage      *int32   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EvaluationLevel  []*int32 `json:"EvaluationLevel,omitempty" xml:"EvaluationLevel,omitempty" type:"Repeated"`
	EvaluationScore  []*int32 `json:"EvaluationScore,omitempty" xml:"EvaluationScore,omitempty" type:"Repeated"`
	EvaluationStatus []*int32 `json:"EvaluationStatus,omitempty" xml:"EvaluationStatus,omitempty" type:"Repeated"`
	// example:
	//
	// 1614599400000
	FirstTimeEnd *int64 `json:"FirstTimeEnd,omitempty" xml:"FirstTimeEnd,omitempty"`
	// example:
	//
	// 1614596400000
	FirstTimeStart *int64 `json:"FirstTimeStart,omitempty" xml:"FirstTimeStart,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemberId   []*int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty" type:"Repeated"`
	MemberName []*string `json:"MemberName,omitempty" xml:"MemberName,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueueId      []*int64  `json:"QueueId,omitempty" xml:"QueueId,omitempty" type:"Repeated"`
	ServicerId   []*int64  `json:"ServicerId,omitempty" xml:"ServicerId,omitempty" type:"Repeated"`
	ServicerName []*string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty" type:"Repeated"`
	TouchId      []*int64  `json:"TouchId,omitempty" xml:"TouchId,omitempty" type:"Repeated"`
	TouchType    []*int32  `json:"TouchType,omitempty" xml:"TouchType,omitempty" type:"Repeated"`
}

func (s QueryTouchListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTouchListRequest) GoString() string {
	return s.String()
}

func (s *QueryTouchListRequest) GetChannelId() []*string {
	return s.ChannelId
}

func (s *QueryTouchListRequest) GetChannelType() []*int32 {
	return s.ChannelType
}

func (s *QueryTouchListRequest) GetCloseTimeEnd() *int64 {
	return s.CloseTimeEnd
}

func (s *QueryTouchListRequest) GetCloseTimeStart() *int64 {
	return s.CloseTimeStart
}

func (s *QueryTouchListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryTouchListRequest) GetEvaluationLevel() []*int32 {
	return s.EvaluationLevel
}

func (s *QueryTouchListRequest) GetEvaluationScore() []*int32 {
	return s.EvaluationScore
}

func (s *QueryTouchListRequest) GetEvaluationStatus() []*int32 {
	return s.EvaluationStatus
}

func (s *QueryTouchListRequest) GetFirstTimeEnd() *int64 {
	return s.FirstTimeEnd
}

func (s *QueryTouchListRequest) GetFirstTimeStart() *int64 {
	return s.FirstTimeStart
}

func (s *QueryTouchListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTouchListRequest) GetMemberId() []*int64 {
	return s.MemberId
}

func (s *QueryTouchListRequest) GetMemberName() []*string {
	return s.MemberName
}

func (s *QueryTouchListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTouchListRequest) GetQueueId() []*int64 {
	return s.QueueId
}

func (s *QueryTouchListRequest) GetServicerId() []*int64 {
	return s.ServicerId
}

func (s *QueryTouchListRequest) GetServicerName() []*string {
	return s.ServicerName
}

func (s *QueryTouchListRequest) GetTouchId() []*int64 {
	return s.TouchId
}

func (s *QueryTouchListRequest) GetTouchType() []*int32 {
	return s.TouchType
}

func (s *QueryTouchListRequest) SetChannelId(v []*string) *QueryTouchListRequest {
	s.ChannelId = v
	return s
}

func (s *QueryTouchListRequest) SetChannelType(v []*int32) *QueryTouchListRequest {
	s.ChannelType = v
	return s
}

func (s *QueryTouchListRequest) SetCloseTimeEnd(v int64) *QueryTouchListRequest {
	s.CloseTimeEnd = &v
	return s
}

func (s *QueryTouchListRequest) SetCloseTimeStart(v int64) *QueryTouchListRequest {
	s.CloseTimeStart = &v
	return s
}

func (s *QueryTouchListRequest) SetCurrentPage(v int32) *QueryTouchListRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryTouchListRequest) SetEvaluationLevel(v []*int32) *QueryTouchListRequest {
	s.EvaluationLevel = v
	return s
}

func (s *QueryTouchListRequest) SetEvaluationScore(v []*int32) *QueryTouchListRequest {
	s.EvaluationScore = v
	return s
}

func (s *QueryTouchListRequest) SetEvaluationStatus(v []*int32) *QueryTouchListRequest {
	s.EvaluationStatus = v
	return s
}

func (s *QueryTouchListRequest) SetFirstTimeEnd(v int64) *QueryTouchListRequest {
	s.FirstTimeEnd = &v
	return s
}

func (s *QueryTouchListRequest) SetFirstTimeStart(v int64) *QueryTouchListRequest {
	s.FirstTimeStart = &v
	return s
}

func (s *QueryTouchListRequest) SetInstanceId(v string) *QueryTouchListRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTouchListRequest) SetMemberId(v []*int64) *QueryTouchListRequest {
	s.MemberId = v
	return s
}

func (s *QueryTouchListRequest) SetMemberName(v []*string) *QueryTouchListRequest {
	s.MemberName = v
	return s
}

func (s *QueryTouchListRequest) SetPageSize(v int32) *QueryTouchListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTouchListRequest) SetQueueId(v []*int64) *QueryTouchListRequest {
	s.QueueId = v
	return s
}

func (s *QueryTouchListRequest) SetServicerId(v []*int64) *QueryTouchListRequest {
	s.ServicerId = v
	return s
}

func (s *QueryTouchListRequest) SetServicerName(v []*string) *QueryTouchListRequest {
	s.ServicerName = v
	return s
}

func (s *QueryTouchListRequest) SetTouchId(v []*int64) *QueryTouchListRequest {
	s.TouchId = v
	return s
}

func (s *QueryTouchListRequest) SetTouchType(v []*int32) *QueryTouchListRequest {
	s.TouchType = v
	return s
}

func (s *QueryTouchListRequest) Validate() error {
	return dara.Validate(s)
}
