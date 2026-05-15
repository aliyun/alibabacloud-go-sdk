// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTouchListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryTouchListResponseBody
	GetCode() *string
	SetMessage(v string) *QueryTouchListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTouchListResponseBody
	GetRequestId() *string
	SetResultData(v *QueryTouchListResponseBodyResultData) *QueryTouchListResponseBody
	GetResultData() *QueryTouchListResponseBodyResultData
	SetSuccess(v bool) *QueryTouchListResponseBody
	GetSuccess() *bool
}

type QueryTouchListResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 865658FD-80DE-5D49-ABEB-F3CC9863F4F1
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultData *QueryTouchListResponseBodyResultData `json:"ResultData,omitempty" xml:"ResultData,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTouchListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTouchListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTouchListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryTouchListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTouchListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTouchListResponseBody) GetResultData() *QueryTouchListResponseBodyResultData {
	return s.ResultData
}

func (s *QueryTouchListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryTouchListResponseBody) SetCode(v string) *QueryTouchListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTouchListResponseBody) SetMessage(v string) *QueryTouchListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTouchListResponseBody) SetRequestId(v string) *QueryTouchListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTouchListResponseBody) SetResultData(v *QueryTouchListResponseBodyResultData) *QueryTouchListResponseBody {
	s.ResultData = v
	return s
}

func (s *QueryTouchListResponseBody) SetSuccess(v bool) *QueryTouchListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTouchListResponseBody) Validate() error {
	if s.ResultData != nil {
		if err := s.ResultData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTouchListResponseBodyResultData struct {
	// example:
	//
	// 1
	CurrentPage *int32                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*QueryTouchListResponseBodyResultDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// false
	Empty *bool `json:"Empty,omitempty" xml:"Empty,omitempty"`
	// example:
	//
	// 2
	NextPage *int32 `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 2
	OnePageSize *int32 `json:"OnePageSize,omitempty" xml:"OnePageSize,omitempty"`
	// example:
	//
	// 2
	PreviousPage *int32 `json:"PreviousPage,omitempty" xml:"PreviousPage,omitempty"`
	// example:
	//
	// 4
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// example:
	//
	// 100
	TotalResults *int32 `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
}

func (s QueryTouchListResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s QueryTouchListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryTouchListResponseBodyResultData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryTouchListResponseBodyResultData) GetData() []*QueryTouchListResponseBodyResultDataData {
	return s.Data
}

func (s *QueryTouchListResponseBodyResultData) GetEmpty() *bool {
	return s.Empty
}

func (s *QueryTouchListResponseBodyResultData) GetNextPage() *int32 {
	return s.NextPage
}

func (s *QueryTouchListResponseBodyResultData) GetOnePageSize() *int32 {
	return s.OnePageSize
}

func (s *QueryTouchListResponseBodyResultData) GetPreviousPage() *int32 {
	return s.PreviousPage
}

func (s *QueryTouchListResponseBodyResultData) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *QueryTouchListResponseBodyResultData) GetTotalResults() *int32 {
	return s.TotalResults
}

func (s *QueryTouchListResponseBodyResultData) SetCurrentPage(v int32) *QueryTouchListResponseBodyResultData {
	s.CurrentPage = &v
	return s
}

func (s *QueryTouchListResponseBodyResultData) SetData(v []*QueryTouchListResponseBodyResultDataData) *QueryTouchListResponseBodyResultData {
	s.Data = v
	return s
}

func (s *QueryTouchListResponseBodyResultData) SetEmpty(v bool) *QueryTouchListResponseBodyResultData {
	s.Empty = &v
	return s
}

func (s *QueryTouchListResponseBodyResultData) SetNextPage(v int32) *QueryTouchListResponseBodyResultData {
	s.NextPage = &v
	return s
}

func (s *QueryTouchListResponseBodyResultData) SetOnePageSize(v int32) *QueryTouchListResponseBodyResultData {
	s.OnePageSize = &v
	return s
}

func (s *QueryTouchListResponseBodyResultData) SetPreviousPage(v int32) *QueryTouchListResponseBodyResultData {
	s.PreviousPage = &v
	return s
}

func (s *QueryTouchListResponseBodyResultData) SetTotalPage(v int32) *QueryTouchListResponseBodyResultData {
	s.TotalPage = &v
	return s
}

func (s *QueryTouchListResponseBodyResultData) SetTotalResults(v int32) *QueryTouchListResponseBodyResultData {
	s.TotalResults = &v
	return s
}

func (s *QueryTouchListResponseBodyResultData) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryTouchListResponseBodyResultDataData struct {
	// example:
	//
	// 905
	BuId *int64 `json:"BuId,omitempty" xml:"BuId,omitempty"`
	// example:
	//
	// 4f8807a9de024507a3090b5b66a8****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1
	ChannelType *int32 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// 1611207976000
	CloseTime       *int64  `json:"CloseTime,omitempty" xml:"CloseTime,omitempty"`
	CommonQueueName *string `json:"CommonQueueName,omitempty" xml:"CommonQueueName,omitempty"`
	// example:
	//
	// 100
	DepId          *int64                                            `json:"DepId,omitempty" xml:"DepId,omitempty"`
	ExtAttrs       *QueryTouchListResponseBodyResultDataDataExtAttrs `json:"ExtAttrs,omitempty" xml:"ExtAttrs,omitempty" type:"Struct"`
	ExtAttrsString map[string]interface{}                            `json:"ExtAttrsString,omitempty" xml:"ExtAttrsString,omitempty"`
	// example:
	//
	// xxxx
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// example:
	//
	// 1611209971000
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// example:
	//
	// 6400665****
	FromId *int64 `json:"FromId,omitempty" xml:"FromId,omitempty"`
	// example:
	//
	// 1611209971000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1611207979000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 6400665****
	MemberId   *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// example:
	//
	// 0
	ParentTouchId *int64 `json:"ParentTouchId,omitempty" xml:"ParentTouchId,omitempty"`
	// example:
	//
	// 111
	QueueId *int64 `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// example:
	//
	// 67****
	ServicerId *int64 `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	// example:
	//
	// 13900001234
	ServicerName *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// null
	SwitchUser *string `json:"SwitchUser,omitempty" xml:"SwitchUser,omitempty"`
	// example:
	//
	// 678026
	ToId *int64 `json:"ToId,omitempty" xml:"ToId,omitempty"`
	// example:
	//
	// 1
	TouchContent *string `json:"TouchContent,omitempty" xml:"TouchContent,omitempty"`
	// example:
	//
	// 2
	TouchEndReason *int32 `json:"TouchEndReason,omitempty" xml:"TouchEndReason,omitempty"`
	// example:
	//
	// 1386****
	TouchId *string `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
	// example:
	//
	// 111
	TouchTime *string `json:"TouchTime,omitempty" xml:"TouchTime,omitempty"`
	// example:
	//
	// 2
	TouchType *int32 `json:"TouchType,omitempty" xml:"TouchType,omitempty"`
	// example:
	//
	// 1386****
	UserTouchId *int64 `json:"UserTouchId,omitempty" xml:"UserTouchId,omitempty"`
}

func (s QueryTouchListResponseBodyResultDataData) String() string {
	return dara.Prettify(s)
}

func (s QueryTouchListResponseBodyResultDataData) GoString() string {
	return s.String()
}

func (s *QueryTouchListResponseBodyResultDataData) GetBuId() *int64 {
	return s.BuId
}

func (s *QueryTouchListResponseBodyResultDataData) GetChannelId() *string {
	return s.ChannelId
}

func (s *QueryTouchListResponseBodyResultDataData) GetChannelType() *int32 {
	return s.ChannelType
}

func (s *QueryTouchListResponseBodyResultDataData) GetCloseTime() *int64 {
	return s.CloseTime
}

func (s *QueryTouchListResponseBodyResultDataData) GetCommonQueueName() *string {
	return s.CommonQueueName
}

func (s *QueryTouchListResponseBodyResultDataData) GetDepId() *int64 {
	return s.DepId
}

func (s *QueryTouchListResponseBodyResultDataData) GetExtAttrs() *QueryTouchListResponseBodyResultDataDataExtAttrs {
	return s.ExtAttrs
}

func (s *QueryTouchListResponseBodyResultDataData) GetExtAttrsString() map[string]interface{} {
	return s.ExtAttrsString
}

func (s *QueryTouchListResponseBodyResultDataData) GetFeedback() *string {
	return s.Feedback
}

func (s *QueryTouchListResponseBodyResultDataData) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *QueryTouchListResponseBodyResultDataData) GetFromId() *int64 {
	return s.FromId
}

func (s *QueryTouchListResponseBodyResultDataData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryTouchListResponseBodyResultDataData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *QueryTouchListResponseBodyResultDataData) GetMemberId() *int64 {
	return s.MemberId
}

func (s *QueryTouchListResponseBodyResultDataData) GetMemberName() *string {
	return s.MemberName
}

func (s *QueryTouchListResponseBodyResultDataData) GetParentTouchId() *int64 {
	return s.ParentTouchId
}

func (s *QueryTouchListResponseBodyResultDataData) GetQueueId() *int64 {
	return s.QueueId
}

func (s *QueryTouchListResponseBodyResultDataData) GetServicerId() *int64 {
	return s.ServicerId
}

func (s *QueryTouchListResponseBodyResultDataData) GetServicerName() *string {
	return s.ServicerName
}

func (s *QueryTouchListResponseBodyResultDataData) GetStatus() *int32 {
	return s.Status
}

func (s *QueryTouchListResponseBodyResultDataData) GetSwitchUser() *string {
	return s.SwitchUser
}

func (s *QueryTouchListResponseBodyResultDataData) GetToId() *int64 {
	return s.ToId
}

func (s *QueryTouchListResponseBodyResultDataData) GetTouchContent() *string {
	return s.TouchContent
}

func (s *QueryTouchListResponseBodyResultDataData) GetTouchEndReason() *int32 {
	return s.TouchEndReason
}

func (s *QueryTouchListResponseBodyResultDataData) GetTouchId() *string {
	return s.TouchId
}

func (s *QueryTouchListResponseBodyResultDataData) GetTouchTime() *string {
	return s.TouchTime
}

func (s *QueryTouchListResponseBodyResultDataData) GetTouchType() *int32 {
	return s.TouchType
}

func (s *QueryTouchListResponseBodyResultDataData) GetUserTouchId() *int64 {
	return s.UserTouchId
}

func (s *QueryTouchListResponseBodyResultDataData) SetBuId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.BuId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetChannelId(v string) *QueryTouchListResponseBodyResultDataData {
	s.ChannelId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetChannelType(v int32) *QueryTouchListResponseBodyResultDataData {
	s.ChannelType = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetCloseTime(v int64) *QueryTouchListResponseBodyResultDataData {
	s.CloseTime = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetCommonQueueName(v string) *QueryTouchListResponseBodyResultDataData {
	s.CommonQueueName = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetDepId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.DepId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetExtAttrs(v *QueryTouchListResponseBodyResultDataDataExtAttrs) *QueryTouchListResponseBodyResultDataData {
	s.ExtAttrs = v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetExtAttrsString(v map[string]interface{}) *QueryTouchListResponseBodyResultDataData {
	s.ExtAttrsString = v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetFeedback(v string) *QueryTouchListResponseBodyResultDataData {
	s.Feedback = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetFirstTime(v int64) *QueryTouchListResponseBodyResultDataData {
	s.FirstTime = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetFromId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.FromId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetGmtCreate(v int64) *QueryTouchListResponseBodyResultDataData {
	s.GmtCreate = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetGmtModified(v int64) *QueryTouchListResponseBodyResultDataData {
	s.GmtModified = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetMemberId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.MemberId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetMemberName(v string) *QueryTouchListResponseBodyResultDataData {
	s.MemberName = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetParentTouchId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.ParentTouchId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetQueueId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.QueueId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetServicerId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.ServicerId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetServicerName(v string) *QueryTouchListResponseBodyResultDataData {
	s.ServicerName = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetStatus(v int32) *QueryTouchListResponseBodyResultDataData {
	s.Status = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetSwitchUser(v string) *QueryTouchListResponseBodyResultDataData {
	s.SwitchUser = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetToId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.ToId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetTouchContent(v string) *QueryTouchListResponseBodyResultDataData {
	s.TouchContent = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetTouchEndReason(v int32) *QueryTouchListResponseBodyResultDataData {
	s.TouchEndReason = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetTouchId(v string) *QueryTouchListResponseBodyResultDataData {
	s.TouchId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetTouchTime(v string) *QueryTouchListResponseBodyResultDataData {
	s.TouchTime = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetTouchType(v int32) *QueryTouchListResponseBodyResultDataData {
	s.TouchType = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetUserTouchId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.UserTouchId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) Validate() error {
	if s.ExtAttrs != nil {
		if err := s.ExtAttrs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTouchListResponseBodyResultDataDataExtAttrs struct {
	// example:
	//
	// 187****0000
	Ani *string `json:"Ani,omitempty" xml:"Ani,omitempty"`
	// example:
	//
	// 05712688****
	Dnis *string `json:"Dnis,omitempty" xml:"Dnis,omitempty"`
	// example:
	//
	// 2
	EvaluationLevel *int32 `json:"EvaluationLevel,omitempty" xml:"EvaluationLevel,omitempty"`
	// example:
	//
	// 1
	EvaluationScore *int32 `json:"EvaluationScore,omitempty" xml:"EvaluationScore,omitempty"`
	// example:
	//
	// 0
	EvaluationSolution *int32 `json:"EvaluationSolution,omitempty" xml:"EvaluationSolution,omitempty"`
	// example:
	//
	// 1
	EvaluationStatus *int32 `json:"EvaluationStatus,omitempty" xml:"EvaluationStatus,omitempty"`
	// example:
	//
	// 1
	OnlineJoinRespInterval *int32 `json:"OnlineJoinRespInterval,omitempty" xml:"OnlineJoinRespInterval,omitempty"`
	// example:
	//
	// 0
	OnlineSessionSource *int32 `json:"OnlineSessionSource,omitempty" xml:"OnlineSessionSource,omitempty"`
	// example:
	//
	// 05712688****
	OutCallRouteNumber *string `json:"OutCallRouteNumber,omitempty" xml:"OutCallRouteNumber,omitempty"`
}

func (s QueryTouchListResponseBodyResultDataDataExtAttrs) String() string {
	return dara.Prettify(s)
}

func (s QueryTouchListResponseBodyResultDataDataExtAttrs) GoString() string {
	return s.String()
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) GetAni() *string {
	return s.Ani
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) GetDnis() *string {
	return s.Dnis
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) GetEvaluationLevel() *int32 {
	return s.EvaluationLevel
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) GetEvaluationScore() *int32 {
	return s.EvaluationScore
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) GetEvaluationSolution() *int32 {
	return s.EvaluationSolution
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) GetEvaluationStatus() *int32 {
	return s.EvaluationStatus
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) GetOnlineJoinRespInterval() *int32 {
	return s.OnlineJoinRespInterval
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) GetOnlineSessionSource() *int32 {
	return s.OnlineSessionSource
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) GetOutCallRouteNumber() *string {
	return s.OutCallRouteNumber
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetAni(v string) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.Ani = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetDnis(v string) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.Dnis = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetEvaluationLevel(v int32) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.EvaluationLevel = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetEvaluationScore(v int32) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.EvaluationScore = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetEvaluationSolution(v int32) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.EvaluationSolution = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetEvaluationStatus(v int32) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.EvaluationStatus = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetOnlineJoinRespInterval(v int32) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.OnlineJoinRespInterval = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetOnlineSessionSource(v int32) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.OnlineSessionSource = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetOutCallRouteNumber(v string) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.OutCallRouteNumber = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) Validate() error {
	return dara.Validate(s)
}
