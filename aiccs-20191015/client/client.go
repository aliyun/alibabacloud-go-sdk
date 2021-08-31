// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddHotlineNumberRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 号码
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// 号码描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 是否用于入呼
	EnableInbound *bool `json:"EnableInbound,omitempty" xml:"EnableInbound,omitempty"`
	// 入呼ivr流程id
	InboundFlowId *int64 `json:"InboundFlowId,omitempty" xml:"InboundFlowId,omitempty"`
	// 是否用于外呼
	EnableOutbound *bool `json:"EnableOutbound,omitempty" xml:"EnableOutbound,omitempty"`
	// 外呼是否对所有部门生效
	OutboundAllDepart *bool `json:"OutboundAllDepart,omitempty" xml:"OutboundAllDepart,omitempty"`
	// 是否开启入呼满意度
	EnableInboundEvaluation *bool `json:"EnableInboundEvaluation,omitempty" xml:"EnableInboundEvaluation,omitempty"`
	// 是否开启外呼满意度
	EnableOutboundEvaluation *bool `json:"EnableOutboundEvaluation,omitempty" xml:"EnableOutboundEvaluation,omitempty"`
	// 满意度等级
	EvaluationLevel *int32 `json:"EvaluationLevel,omitempty" xml:"EvaluationLevel,omitempty"`
	// 外呼生效范围
	OutboundRangeList []*AddHotlineNumberRequestOutboundRangeList `json:"OutboundRangeList,omitempty" xml:"OutboundRangeList,omitempty" type:"Repeated"`
}

func (s AddHotlineNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddHotlineNumberRequest) GoString() string {
	return s.String()
}

func (s *AddHotlineNumberRequest) SetInstanceId(v string) *AddHotlineNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *AddHotlineNumberRequest) SetHotlineNumber(v string) *AddHotlineNumberRequest {
	s.HotlineNumber = &v
	return s
}

func (s *AddHotlineNumberRequest) SetDescription(v string) *AddHotlineNumberRequest {
	s.Description = &v
	return s
}

func (s *AddHotlineNumberRequest) SetEnableInbound(v bool) *AddHotlineNumberRequest {
	s.EnableInbound = &v
	return s
}

func (s *AddHotlineNumberRequest) SetInboundFlowId(v int64) *AddHotlineNumberRequest {
	s.InboundFlowId = &v
	return s
}

func (s *AddHotlineNumberRequest) SetEnableOutbound(v bool) *AddHotlineNumberRequest {
	s.EnableOutbound = &v
	return s
}

func (s *AddHotlineNumberRequest) SetOutboundAllDepart(v bool) *AddHotlineNumberRequest {
	s.OutboundAllDepart = &v
	return s
}

func (s *AddHotlineNumberRequest) SetEnableInboundEvaluation(v bool) *AddHotlineNumberRequest {
	s.EnableInboundEvaluation = &v
	return s
}

func (s *AddHotlineNumberRequest) SetEnableOutboundEvaluation(v bool) *AddHotlineNumberRequest {
	s.EnableOutboundEvaluation = &v
	return s
}

func (s *AddHotlineNumberRequest) SetEvaluationLevel(v int32) *AddHotlineNumberRequest {
	s.EvaluationLevel = &v
	return s
}

func (s *AddHotlineNumberRequest) SetOutboundRangeList(v []*AddHotlineNumberRequestOutboundRangeList) *AddHotlineNumberRequest {
	s.OutboundRangeList = v
	return s
}

type AddHotlineNumberRequestOutboundRangeList struct {
	// 生效部门id
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// 生效技能组列表（部门123下）
	GroupIdList []*int64 `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty" type:"Repeated"`
}

func (s AddHotlineNumberRequestOutboundRangeList) String() string {
	return tea.Prettify(s)
}

func (s AddHotlineNumberRequestOutboundRangeList) GoString() string {
	return s.String()
}

func (s *AddHotlineNumberRequestOutboundRangeList) SetDepartmentId(v int64) *AddHotlineNumberRequestOutboundRangeList {
	s.DepartmentId = &v
	return s
}

func (s *AddHotlineNumberRequestOutboundRangeList) SetGroupIdList(v []*int64) *AddHotlineNumberRequestOutboundRangeList {
	s.GroupIdList = v
	return s
}

type AddHotlineNumberShrinkRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 号码
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// 号码描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 是否用于入呼
	EnableInbound *bool `json:"EnableInbound,omitempty" xml:"EnableInbound,omitempty"`
	// 入呼ivr流程id
	InboundFlowId *int64 `json:"InboundFlowId,omitempty" xml:"InboundFlowId,omitempty"`
	// 是否用于外呼
	EnableOutbound *bool `json:"EnableOutbound,omitempty" xml:"EnableOutbound,omitempty"`
	// 外呼是否对所有部门生效
	OutboundAllDepart *bool `json:"OutboundAllDepart,omitempty" xml:"OutboundAllDepart,omitempty"`
	// 是否开启入呼满意度
	EnableInboundEvaluation *bool `json:"EnableInboundEvaluation,omitempty" xml:"EnableInboundEvaluation,omitempty"`
	// 是否开启外呼满意度
	EnableOutboundEvaluation *bool `json:"EnableOutboundEvaluation,omitempty" xml:"EnableOutboundEvaluation,omitempty"`
	// 满意度等级
	EvaluationLevel *int32 `json:"EvaluationLevel,omitempty" xml:"EvaluationLevel,omitempty"`
	// 外呼生效范围
	OutboundRangeListShrink *string `json:"OutboundRangeList,omitempty" xml:"OutboundRangeList,omitempty"`
}

func (s AddHotlineNumberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddHotlineNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddHotlineNumberShrinkRequest) SetInstanceId(v string) *AddHotlineNumberShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetHotlineNumber(v string) *AddHotlineNumberShrinkRequest {
	s.HotlineNumber = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetDescription(v string) *AddHotlineNumberShrinkRequest {
	s.Description = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetEnableInbound(v bool) *AddHotlineNumberShrinkRequest {
	s.EnableInbound = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetInboundFlowId(v int64) *AddHotlineNumberShrinkRequest {
	s.InboundFlowId = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetEnableOutbound(v bool) *AddHotlineNumberShrinkRequest {
	s.EnableOutbound = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetOutboundAllDepart(v bool) *AddHotlineNumberShrinkRequest {
	s.OutboundAllDepart = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetEnableInboundEvaluation(v bool) *AddHotlineNumberShrinkRequest {
	s.EnableInboundEvaluation = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetEnableOutboundEvaluation(v bool) *AddHotlineNumberShrinkRequest {
	s.EnableOutboundEvaluation = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetEvaluationLevel(v int32) *AddHotlineNumberShrinkRequest {
	s.EvaluationLevel = &v
	return s
}

func (s *AddHotlineNumberShrinkRequest) SetOutboundRangeListShrink(v string) *AddHotlineNumberShrinkRequest {
	s.OutboundRangeListShrink = &v
	return s
}

type AddHotlineNumberResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 接口调用是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// http状态码
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s AddHotlineNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddHotlineNumberResponseBody) GoString() string {
	return s.String()
}

func (s *AddHotlineNumberResponseBody) SetRequestId(v string) *AddHotlineNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddHotlineNumberResponseBody) SetSuccess(v bool) *AddHotlineNumberResponseBody {
	s.Success = &v
	return s
}

func (s *AddHotlineNumberResponseBody) SetCode(v string) *AddHotlineNumberResponseBody {
	s.Code = &v
	return s
}

func (s *AddHotlineNumberResponseBody) SetMessage(v string) *AddHotlineNumberResponseBody {
	s.Message = &v
	return s
}

func (s *AddHotlineNumberResponseBody) SetHttpStatusCode(v int64) *AddHotlineNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

type AddHotlineNumberResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddHotlineNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddHotlineNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddHotlineNumberResponse) GoString() string {
	return s.String()
}

func (s *AddHotlineNumberResponse) SetHeaders(v map[string]*string) *AddHotlineNumberResponse {
	s.Headers = v
	return s
}

func (s *AddHotlineNumberResponse) SetBody(v *AddHotlineNumberResponseBody) *AddHotlineNumberResponse {
	s.Body = v
	return s
}

type AddOuterAccountRequest struct {
	OuterAccountId      *string `json:"OuterAccountId,omitempty" xml:"OuterAccountId,omitempty"`
	OuterAccountType    *string `json:"OuterAccountType,omitempty" xml:"OuterAccountType,omitempty"`
	OuterAccountName    *string `json:"OuterAccountName,omitempty" xml:"OuterAccountName,omitempty"`
	Avatar              *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	RealName            *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	OuterDepartmentId   *string `json:"OuterDepartmentId,omitempty" xml:"OuterDepartmentId,omitempty"`
	OuterGroupIds       *string `json:"OuterGroupIds,omitempty" xml:"OuterGroupIds,omitempty"`
	Ext                 *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	OuterDepartmentType *string `json:"OuterDepartmentType,omitempty" xml:"OuterDepartmentType,omitempty"`
	OuterGroupType      *string `json:"OuterGroupType,omitempty" xml:"OuterGroupType,omitempty"`
}

func (s AddOuterAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOuterAccountRequest) GoString() string {
	return s.String()
}

func (s *AddOuterAccountRequest) SetOuterAccountId(v string) *AddOuterAccountRequest {
	s.OuterAccountId = &v
	return s
}

func (s *AddOuterAccountRequest) SetOuterAccountType(v string) *AddOuterAccountRequest {
	s.OuterAccountType = &v
	return s
}

func (s *AddOuterAccountRequest) SetOuterAccountName(v string) *AddOuterAccountRequest {
	s.OuterAccountName = &v
	return s
}

func (s *AddOuterAccountRequest) SetAvatar(v string) *AddOuterAccountRequest {
	s.Avatar = &v
	return s
}

func (s *AddOuterAccountRequest) SetRealName(v string) *AddOuterAccountRequest {
	s.RealName = &v
	return s
}

func (s *AddOuterAccountRequest) SetOuterDepartmentId(v string) *AddOuterAccountRequest {
	s.OuterDepartmentId = &v
	return s
}

func (s *AddOuterAccountRequest) SetOuterGroupIds(v string) *AddOuterAccountRequest {
	s.OuterGroupIds = &v
	return s
}

func (s *AddOuterAccountRequest) SetExt(v string) *AddOuterAccountRequest {
	s.Ext = &v
	return s
}

func (s *AddOuterAccountRequest) SetOuterDepartmentType(v string) *AddOuterAccountRequest {
	s.OuterDepartmentType = &v
	return s
}

func (s *AddOuterAccountRequest) SetOuterGroupType(v string) *AddOuterAccountRequest {
	s.OuterGroupType = &v
	return s
}

type AddOuterAccountResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddOuterAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOuterAccountResponseBody) GoString() string {
	return s.String()
}

func (s *AddOuterAccountResponseBody) SetMessage(v string) *AddOuterAccountResponseBody {
	s.Message = &v
	return s
}

func (s *AddOuterAccountResponseBody) SetRequestId(v string) *AddOuterAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddOuterAccountResponseBody) SetData(v string) *AddOuterAccountResponseBody {
	s.Data = &v
	return s
}

func (s *AddOuterAccountResponseBody) SetCode(v string) *AddOuterAccountResponseBody {
	s.Code = &v
	return s
}

func (s *AddOuterAccountResponseBody) SetSuccess(v bool) *AddOuterAccountResponseBody {
	s.Success = &v
	return s
}

type AddOuterAccountResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddOuterAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddOuterAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOuterAccountResponse) GoString() string {
	return s.String()
}

func (s *AddOuterAccountResponse) SetHeaders(v map[string]*string) *AddOuterAccountResponse {
	s.Headers = v
	return s
}

func (s *AddOuterAccountResponse) SetBody(v *AddOuterAccountResponseBody) *AddOuterAccountResponse {
	s.Body = v
	return s
}

type AddSkillGroupRequest struct {
	OuterGroupId        *string `json:"OuterGroupId,omitempty" xml:"OuterGroupId,omitempty"`
	OuterGroupName      *string `json:"OuterGroupName,omitempty" xml:"OuterGroupName,omitempty"`
	OuterGroupType      *string `json:"OuterGroupType,omitempty" xml:"OuterGroupType,omitempty"`
	OuterDepartmentId   *string `json:"OuterDepartmentId,omitempty" xml:"OuterDepartmentId,omitempty"`
	OuterDepartmentType *string `json:"OuterDepartmentType,omitempty" xml:"OuterDepartmentType,omitempty"`
}

func (s AddSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *AddSkillGroupRequest) SetOuterGroupId(v string) *AddSkillGroupRequest {
	s.OuterGroupId = &v
	return s
}

func (s *AddSkillGroupRequest) SetOuterGroupName(v string) *AddSkillGroupRequest {
	s.OuterGroupName = &v
	return s
}

func (s *AddSkillGroupRequest) SetOuterGroupType(v string) *AddSkillGroupRequest {
	s.OuterGroupType = &v
	return s
}

func (s *AddSkillGroupRequest) SetOuterDepartmentId(v string) *AddSkillGroupRequest {
	s.OuterDepartmentId = &v
	return s
}

func (s *AddSkillGroupRequest) SetOuterDepartmentType(v string) *AddSkillGroupRequest {
	s.OuterDepartmentType = &v
	return s
}

type AddSkillGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddSkillGroupResponseBody) SetMessage(v string) *AddSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AddSkillGroupResponseBody) SetRequestId(v string) *AddSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSkillGroupResponseBody) SetData(v string) *AddSkillGroupResponseBody {
	s.Data = &v
	return s
}

func (s *AddSkillGroupResponseBody) SetCode(v string) *AddSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddSkillGroupResponseBody) SetSuccess(v bool) *AddSkillGroupResponseBody {
	s.Success = &v
	return s
}

type AddSkillGroupResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *AddSkillGroupResponse) SetHeaders(v map[string]*string) *AddSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *AddSkillGroupResponse) SetBody(v *AddSkillGroupResponseBody) *AddSkillGroupResponse {
	s.Body = v
	return s
}

type AiccsSmartCallRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	CalledShowNumber     *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	VoiceCode            *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	PlayTimes            *int32  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	Volume               *int32  `json:"Volume,omitempty" xml:"Volume,omitempty"`
	Speed                *int32  `json:"Speed,omitempty" xml:"Speed,omitempty"`
	AsrModelId           *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	AsrBaseId            *string `json:"AsrBaseId,omitempty" xml:"AsrBaseId,omitempty"`
	AsrAlsAmId           *string `json:"AsrAlsAmId,omitempty" xml:"AsrAlsAmId,omitempty"`
	AsrVocabularyId      *string `json:"AsrVocabularyId,omitempty" xml:"AsrVocabularyId,omitempty"`
	RecordFlag           *bool   `json:"RecordFlag,omitempty" xml:"RecordFlag,omitempty"`
	PauseTime            *int32  `json:"PauseTime,omitempty" xml:"PauseTime,omitempty"`
	MuteTime             *int32  `json:"MuteTime,omitempty" xml:"MuteTime,omitempty"`
	ActionCodeBreak      *bool   `json:"ActionCodeBreak,omitempty" xml:"ActionCodeBreak,omitempty"`
	DynamicId            *string `json:"DynamicId,omitempty" xml:"DynamicId,omitempty"`
	EarlyMediaAsr        *bool   `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	VoiceCodeParam       *string `json:"VoiceCodeParam,omitempty" xml:"VoiceCodeParam,omitempty"`
	SessionTimeout       *int32  `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	ActionCodeTimeBreak  *int32  `json:"ActionCodeTimeBreak,omitempty" xml:"ActionCodeTimeBreak,omitempty"`
	TtsConf              *bool   `json:"TtsConf,omitempty" xml:"TtsConf,omitempty"`
	TtsStyle             *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	TtsVolume            *int32  `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	TtsSpeed             *int32  `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
}

func (s AiccsSmartCallRequest) String() string {
	return tea.Prettify(s)
}

func (s AiccsSmartCallRequest) GoString() string {
	return s.String()
}

func (s *AiccsSmartCallRequest) SetOwnerId(v int64) *AiccsSmartCallRequest {
	s.OwnerId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetResourceOwnerAccount(v string) *AiccsSmartCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AiccsSmartCallRequest) SetResourceOwnerId(v int64) *AiccsSmartCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetProdCode(v string) *AiccsSmartCallRequest {
	s.ProdCode = &v
	return s
}

func (s *AiccsSmartCallRequest) SetCalledShowNumber(v string) *AiccsSmartCallRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *AiccsSmartCallRequest) SetCalledNumber(v string) *AiccsSmartCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *AiccsSmartCallRequest) SetVoiceCode(v string) *AiccsSmartCallRequest {
	s.VoiceCode = &v
	return s
}

func (s *AiccsSmartCallRequest) SetOutId(v string) *AiccsSmartCallRequest {
	s.OutId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetPlayTimes(v int32) *AiccsSmartCallRequest {
	s.PlayTimes = &v
	return s
}

func (s *AiccsSmartCallRequest) SetVolume(v int32) *AiccsSmartCallRequest {
	s.Volume = &v
	return s
}

func (s *AiccsSmartCallRequest) SetSpeed(v int32) *AiccsSmartCallRequest {
	s.Speed = &v
	return s
}

func (s *AiccsSmartCallRequest) SetAsrModelId(v string) *AiccsSmartCallRequest {
	s.AsrModelId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetAsrBaseId(v string) *AiccsSmartCallRequest {
	s.AsrBaseId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetAsrAlsAmId(v string) *AiccsSmartCallRequest {
	s.AsrAlsAmId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetAsrVocabularyId(v string) *AiccsSmartCallRequest {
	s.AsrVocabularyId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetRecordFlag(v bool) *AiccsSmartCallRequest {
	s.RecordFlag = &v
	return s
}

func (s *AiccsSmartCallRequest) SetPauseTime(v int32) *AiccsSmartCallRequest {
	s.PauseTime = &v
	return s
}

func (s *AiccsSmartCallRequest) SetMuteTime(v int32) *AiccsSmartCallRequest {
	s.MuteTime = &v
	return s
}

func (s *AiccsSmartCallRequest) SetActionCodeBreak(v bool) *AiccsSmartCallRequest {
	s.ActionCodeBreak = &v
	return s
}

func (s *AiccsSmartCallRequest) SetDynamicId(v string) *AiccsSmartCallRequest {
	s.DynamicId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetEarlyMediaAsr(v bool) *AiccsSmartCallRequest {
	s.EarlyMediaAsr = &v
	return s
}

func (s *AiccsSmartCallRequest) SetVoiceCodeParam(v string) *AiccsSmartCallRequest {
	s.VoiceCodeParam = &v
	return s
}

func (s *AiccsSmartCallRequest) SetSessionTimeout(v int32) *AiccsSmartCallRequest {
	s.SessionTimeout = &v
	return s
}

func (s *AiccsSmartCallRequest) SetActionCodeTimeBreak(v int32) *AiccsSmartCallRequest {
	s.ActionCodeTimeBreak = &v
	return s
}

func (s *AiccsSmartCallRequest) SetTtsConf(v bool) *AiccsSmartCallRequest {
	s.TtsConf = &v
	return s
}

func (s *AiccsSmartCallRequest) SetTtsStyle(v string) *AiccsSmartCallRequest {
	s.TtsStyle = &v
	return s
}

func (s *AiccsSmartCallRequest) SetTtsVolume(v int32) *AiccsSmartCallRequest {
	s.TtsVolume = &v
	return s
}

func (s *AiccsSmartCallRequest) SetTtsSpeed(v int32) *AiccsSmartCallRequest {
	s.TtsSpeed = &v
	return s
}

type AiccsSmartCallResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AiccsSmartCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AiccsSmartCallResponseBody) GoString() string {
	return s.String()
}

func (s *AiccsSmartCallResponseBody) SetMessage(v string) *AiccsSmartCallResponseBody {
	s.Message = &v
	return s
}

func (s *AiccsSmartCallResponseBody) SetRequestId(v string) *AiccsSmartCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *AiccsSmartCallResponseBody) SetData(v string) *AiccsSmartCallResponseBody {
	s.Data = &v
	return s
}

func (s *AiccsSmartCallResponseBody) SetCode(v string) *AiccsSmartCallResponseBody {
	s.Code = &v
	return s
}

type AiccsSmartCallResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AiccsSmartCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AiccsSmartCallResponse) String() string {
	return tea.Prettify(s)
}

func (s AiccsSmartCallResponse) GoString() string {
	return s.String()
}

func (s *AiccsSmartCallResponse) SetHeaders(v map[string]*string) *AiccsSmartCallResponse {
	s.Headers = v
	return s
}

func (s *AiccsSmartCallResponse) SetBody(v *AiccsSmartCallResponseBody) *AiccsSmartCallResponse {
	s.Body = v
	return s
}

type AiccsSmartCallOperateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Command              *string `json:"Command,omitempty" xml:"Command,omitempty"`
	Param                *string `json:"Param,omitempty" xml:"Param,omitempty"`
}

func (s AiccsSmartCallOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s AiccsSmartCallOperateRequest) GoString() string {
	return s.String()
}

func (s *AiccsSmartCallOperateRequest) SetOwnerId(v int64) *AiccsSmartCallOperateRequest {
	s.OwnerId = &v
	return s
}

func (s *AiccsSmartCallOperateRequest) SetResourceOwnerAccount(v string) *AiccsSmartCallOperateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AiccsSmartCallOperateRequest) SetResourceOwnerId(v int64) *AiccsSmartCallOperateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AiccsSmartCallOperateRequest) SetProdCode(v string) *AiccsSmartCallOperateRequest {
	s.ProdCode = &v
	return s
}

func (s *AiccsSmartCallOperateRequest) SetCallId(v string) *AiccsSmartCallOperateRequest {
	s.CallId = &v
	return s
}

func (s *AiccsSmartCallOperateRequest) SetCommand(v string) *AiccsSmartCallOperateRequest {
	s.Command = &v
	return s
}

func (s *AiccsSmartCallOperateRequest) SetParam(v string) *AiccsSmartCallOperateRequest {
	s.Param = &v
	return s
}

type AiccsSmartCallOperateResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AiccsSmartCallOperateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AiccsSmartCallOperateResponseBody) GoString() string {
	return s.String()
}

func (s *AiccsSmartCallOperateResponseBody) SetMessage(v string) *AiccsSmartCallOperateResponseBody {
	s.Message = &v
	return s
}

func (s *AiccsSmartCallOperateResponseBody) SetRequestId(v string) *AiccsSmartCallOperateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AiccsSmartCallOperateResponseBody) SetData(v string) *AiccsSmartCallOperateResponseBody {
	s.Data = &v
	return s
}

func (s *AiccsSmartCallOperateResponseBody) SetCode(v string) *AiccsSmartCallOperateResponseBody {
	s.Code = &v
	return s
}

type AiccsSmartCallOperateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AiccsSmartCallOperateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AiccsSmartCallOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s AiccsSmartCallOperateResponse) GoString() string {
	return s.String()
}

func (s *AiccsSmartCallOperateResponse) SetHeaders(v map[string]*string) *AiccsSmartCallOperateResponse {
	s.Headers = v
	return s
}

func (s *AiccsSmartCallOperateResponse) SetBody(v *AiccsSmartCallOperateResponseBody) *AiccsSmartCallOperateResponse {
	s.Body = v
	return s
}

type AnswerCallRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
}

func (s AnswerCallRequest) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallRequest) GoString() string {
	return s.String()
}

func (s *AnswerCallRequest) SetClientToken(v string) *AnswerCallRequest {
	s.ClientToken = &v
	return s
}

func (s *AnswerCallRequest) SetInstanceId(v string) *AnswerCallRequest {
	s.InstanceId = &v
	return s
}

func (s *AnswerCallRequest) SetAccountName(v string) *AnswerCallRequest {
	s.AccountName = &v
	return s
}

func (s *AnswerCallRequest) SetCallId(v string) *AnswerCallRequest {
	s.CallId = &v
	return s
}

func (s *AnswerCallRequest) SetJobId(v string) *AnswerCallRequest {
	s.JobId = &v
	return s
}

func (s *AnswerCallRequest) SetConnectionId(v string) *AnswerCallRequest {
	s.ConnectionId = &v
	return s
}

type AnswerCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AnswerCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallResponseBody) GoString() string {
	return s.String()
}

func (s *AnswerCallResponseBody) SetCode(v string) *AnswerCallResponseBody {
	s.Code = &v
	return s
}

func (s *AnswerCallResponseBody) SetMessage(v string) *AnswerCallResponseBody {
	s.Message = &v
	return s
}

func (s *AnswerCallResponseBody) SetRequestId(v string) *AnswerCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnswerCallResponseBody) SetSuccess(v bool) *AnswerCallResponseBody {
	s.Success = &v
	return s
}

type AnswerCallResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AnswerCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AnswerCallResponse) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallResponse) GoString() string {
	return s.String()
}

func (s *AnswerCallResponse) SetHeaders(v map[string]*string) *AnswerCallResponse {
	s.Headers = v
	return s
}

func (s *AnswerCallResponse) SetBody(v *AnswerCallResponseBody) *AnswerCallResponse {
	s.Body = v
	return s
}

type AttachTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	CallString           *string `json:"CallString,omitempty" xml:"CallString,omitempty"`
}

func (s AttachTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachTaskRequest) GoString() string {
	return s.String()
}

func (s *AttachTaskRequest) SetOwnerId(v int64) *AttachTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachTaskRequest) SetResourceOwnerAccount(v string) *AttachTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachTaskRequest) SetResourceOwnerId(v int64) *AttachTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachTaskRequest) SetTaskId(v int64) *AttachTaskRequest {
	s.TaskId = &v
	return s
}

func (s *AttachTaskRequest) SetCallString(v string) *AttachTaskRequest {
	s.CallString = &v
	return s
}

type AttachTaskResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AttachTaskResponseBody) SetRequestId(v string) *AttachTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachTaskResponseBody) SetData(v int64) *AttachTaskResponseBody {
	s.Data = &v
	return s
}

func (s *AttachTaskResponseBody) SetCode(v string) *AttachTaskResponseBody {
	s.Code = &v
	return s
}

func (s *AttachTaskResponseBody) SetMessage(v string) *AttachTaskResponseBody {
	s.Message = &v
	return s
}

func (s *AttachTaskResponseBody) SetSuccess(v bool) *AttachTaskResponseBody {
	s.Success = &v
	return s
}

type AttachTaskResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachTaskResponse) GoString() string {
	return s.String()
}

func (s *AttachTaskResponse) SetHeaders(v map[string]*string) *AttachTaskResponse {
	s.Headers = v
	return s
}

func (s *AttachTaskResponse) SetBody(v *AttachTaskResponseBody) *AttachTaskResponse {
	s.Body = v
	return s
}

type BatchCreateQualityProjectsRequest struct {
	ProjectName      *string   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	CheckFreqType    *int32    `json:"CheckFreqType,omitempty" xml:"CheckFreqType,omitempty"`
	TimeRangeStart   *string   `json:"TimeRangeStart,omitempty" xml:"TimeRangeStart,omitempty"`
	TimeRangeEnd     *string   `json:"TimeRangeEnd,omitempty" xml:"TimeRangeEnd,omitempty"`
	AnalysisIds      []*int    `json:"AnalysisIds,omitempty" xml:"AnalysisIds,omitempty" type:"Repeated"`
	InstanceList     []*string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	ChannelTouchType []*int    `json:"ChannelTouchType,omitempty" xml:"ChannelTouchType,omitempty" type:"Repeated"`
}

func (s BatchCreateQualityProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateQualityProjectsRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateQualityProjectsRequest) SetProjectName(v string) *BatchCreateQualityProjectsRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchCreateQualityProjectsRequest) SetCheckFreqType(v int32) *BatchCreateQualityProjectsRequest {
	s.CheckFreqType = &v
	return s
}

func (s *BatchCreateQualityProjectsRequest) SetTimeRangeStart(v string) *BatchCreateQualityProjectsRequest {
	s.TimeRangeStart = &v
	return s
}

func (s *BatchCreateQualityProjectsRequest) SetTimeRangeEnd(v string) *BatchCreateQualityProjectsRequest {
	s.TimeRangeEnd = &v
	return s
}

func (s *BatchCreateQualityProjectsRequest) SetAnalysisIds(v []*int) *BatchCreateQualityProjectsRequest {
	s.AnalysisIds = v
	return s
}

func (s *BatchCreateQualityProjectsRequest) SetInstanceList(v []*string) *BatchCreateQualityProjectsRequest {
	s.InstanceList = v
	return s
}

func (s *BatchCreateQualityProjectsRequest) SetChannelTouchType(v []*int) *BatchCreateQualityProjectsRequest {
	s.ChannelTouchType = v
	return s
}

type BatchCreateQualityProjectsResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      []*BatchCreateQualityProjectsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s BatchCreateQualityProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateQualityProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateQualityProjectsResponseBody) SetCode(v string) *BatchCreateQualityProjectsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchCreateQualityProjectsResponseBody) SetMessage(v string) *BatchCreateQualityProjectsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchCreateQualityProjectsResponseBody) SetRequestId(v string) *BatchCreateQualityProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateQualityProjectsResponseBody) SetSuccess(v bool) *BatchCreateQualityProjectsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchCreateQualityProjectsResponseBody) SetData(v []*BatchCreateQualityProjectsResponseBodyData) *BatchCreateQualityProjectsResponseBody {
	s.Data = v
	return s
}

type BatchCreateQualityProjectsResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Version    *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
	ProjectId  *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s BatchCreateQualityProjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateQualityProjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchCreateQualityProjectsResponseBodyData) SetInstanceId(v string) *BatchCreateQualityProjectsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *BatchCreateQualityProjectsResponseBodyData) SetVersion(v int32) *BatchCreateQualityProjectsResponseBodyData {
	s.Version = &v
	return s
}

func (s *BatchCreateQualityProjectsResponseBodyData) SetProjectId(v int64) *BatchCreateQualityProjectsResponseBodyData {
	s.ProjectId = &v
	return s
}

type BatchCreateQualityProjectsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchCreateQualityProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchCreateQualityProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateQualityProjectsResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateQualityProjectsResponse) SetHeaders(v map[string]*string) *BatchCreateQualityProjectsResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateQualityProjectsResponse) SetBody(v *BatchCreateQualityProjectsResponseBody) *BatchCreateQualityProjectsResponse {
	s.Body = v
	return s
}

type CancelTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelTaskRequest) SetOwnerId(v int64) *CancelTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelTaskRequest) SetResourceOwnerAccount(v string) *CancelTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelTaskRequest) SetResourceOwnerId(v int64) *CancelTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelTaskRequest) SetTaskId(v int64) *CancelTaskRequest {
	s.TaskId = &v
	return s
}

type CancelTaskResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelTaskResponseBody) SetRequestId(v string) *CancelTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelTaskResponseBody) SetData(v bool) *CancelTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelTaskResponseBody) SetCode(v string) *CancelTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelTaskResponseBody) SetMessage(v string) *CancelTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelTaskResponseBody) SetSuccess(v bool) *CancelTaskResponseBody {
	s.Success = &v
	return s
}

type CancelTaskResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelTaskResponse) SetHeaders(v map[string]*string) *CancelTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelTaskResponse) SetBody(v *CancelTaskResponseBody) *CancelTaskResponse {
	s.Body = v
	return s
}

type ChangeChatAgentStatusRequest struct {
	// clientToken
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 示例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 账户名称
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// 修改到的状态类型
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s ChangeChatAgentStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeChatAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *ChangeChatAgentStatusRequest) SetClientToken(v string) *ChangeChatAgentStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *ChangeChatAgentStatusRequest) SetInstanceId(v string) *ChangeChatAgentStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeChatAgentStatusRequest) SetAccountName(v string) *ChangeChatAgentStatusRequest {
	s.AccountName = &v
	return s
}

func (s *ChangeChatAgentStatusRequest) SetMethod(v string) *ChangeChatAgentStatusRequest {
	s.Method = &v
	return s
}

type ChangeChatAgentStatusResponseBody struct {
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeChatAgentStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeChatAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeChatAgentStatusResponseBody) SetMessage(v string) *ChangeChatAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetRequestId(v string) *ChangeChatAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetHttpStatusCode(v int32) *ChangeChatAgentStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetData(v string) *ChangeChatAgentStatusResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetCode(v string) *ChangeChatAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetSuccess(v bool) *ChangeChatAgentStatusResponseBody {
	s.Success = &v
	return s
}

type ChangeChatAgentStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeChatAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeChatAgentStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeChatAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *ChangeChatAgentStatusResponse) SetHeaders(v map[string]*string) *ChangeChatAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *ChangeChatAgentStatusResponse) SetBody(v *ChangeChatAgentStatusResponseBody) *ChangeChatAgentStatusResponse {
	s.Body = v
	return s
}

type ChangeQualityProjectStatusRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ProjectId  *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ChangeQualityProjectStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeQualityProjectStatusRequest) GoString() string {
	return s.String()
}

func (s *ChangeQualityProjectStatusRequest) SetInstanceId(v string) *ChangeQualityProjectStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeQualityProjectStatusRequest) SetStatus(v int32) *ChangeQualityProjectStatusRequest {
	s.Status = &v
	return s
}

func (s *ChangeQualityProjectStatusRequest) SetProjectId(v int64) *ChangeQualityProjectStatusRequest {
	s.ProjectId = &v
	return s
}

type ChangeQualityProjectStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeQualityProjectStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeQualityProjectStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeQualityProjectStatusResponseBody) SetCode(v string) *ChangeQualityProjectStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeQualityProjectStatusResponseBody) SetMessage(v string) *ChangeQualityProjectStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeQualityProjectStatusResponseBody) SetData(v string) *ChangeQualityProjectStatusResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeQualityProjectStatusResponseBody) SetRequestId(v string) *ChangeQualityProjectStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeQualityProjectStatusResponseBody) SetSuccess(v bool) *ChangeQualityProjectStatusResponseBody {
	s.Success = &v
	return s
}

type ChangeQualityProjectStatusResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeQualityProjectStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeQualityProjectStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeQualityProjectStatusResponse) GoString() string {
	return s.String()
}

func (s *ChangeQualityProjectStatusResponse) SetHeaders(v map[string]*string) *ChangeQualityProjectStatusResponse {
	s.Headers = v
	return s
}

func (s *ChangeQualityProjectStatusResponse) SetBody(v *ChangeQualityProjectStatusResponseBody) *ChangeQualityProjectStatusResponse {
	s.Body = v
	return s
}

type CreateAgentRequest struct {
	// js sdk中自动生成的鉴权token
	ClientToken      *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId       *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName      *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName      *string  `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	SkillGroupId     []*int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty" type:"Repeated"`
	SkillGroupIdList []*int64 `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty" type:"Repeated"`
}

func (s CreateAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRequest) SetClientToken(v string) *CreateAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAgentRequest) SetInstanceId(v string) *CreateAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAgentRequest) SetAccountName(v string) *CreateAgentRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAgentRequest) SetDisplayName(v string) *CreateAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateAgentRequest) SetSkillGroupId(v []*int64) *CreateAgentRequest {
	s.SkillGroupId = v
	return s
}

func (s *CreateAgentRequest) SetSkillGroupIdList(v []*int64) *CreateAgentRequest {
	s.SkillGroupIdList = v
	return s
}

type CreateAgentResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s CreateAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBody) SetMessage(v string) *CreateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgentResponseBody) SetRequestId(v string) *CreateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentResponseBody) SetData(v int64) *CreateAgentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAgentResponseBody) SetCode(v string) *CreateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgentResponseBody) SetSuccess(v bool) *CreateAgentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAgentResponseBody) SetHttpStatusCode(v int64) *CreateAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

type CreateAgentResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentResponse) SetHeaders(v map[string]*string) *CreateAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentResponse) SetBody(v *CreateAgentResponseBody) *CreateAgentResponse {
	s.Body = v
	return s
}

type CreateDepartmentRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 部门名称
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
}

func (s CreateDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentRequest) GoString() string {
	return s.String()
}

func (s *CreateDepartmentRequest) SetInstanceId(v string) *CreateDepartmentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDepartmentRequest) SetDepartmentName(v string) *CreateDepartmentRequest {
	s.DepartmentName = &v
	return s
}

type CreateDepartmentResponseBody struct {
	// Id of the request
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s CreateDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDepartmentResponseBody) SetRequestId(v string) *CreateDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDepartmentResponseBody) SetSuccess(v bool) *CreateDepartmentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDepartmentResponseBody) SetCode(v string) *CreateDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDepartmentResponseBody) SetMessage(v string) *CreateDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDepartmentResponseBody) SetHttpStatusCode(v int32) *CreateDepartmentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDepartmentResponseBody) SetData(v int64) *CreateDepartmentResponseBody {
	s.Data = &v
	return s
}

type CreateDepartmentResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentResponse) GoString() string {
	return s.String()
}

func (s *CreateDepartmentResponse) SetHeaders(v map[string]*string) *CreateDepartmentResponse {
	s.Headers = v
	return s
}

func (s *CreateDepartmentResponse) SetBody(v *CreateDepartmentResponseBody) *CreateDepartmentResponse {
	s.Body = v
	return s
}

type CreateOutboundTaskRequest struct {
	TaskType      *int32  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskName      *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RetryTime     *int32  `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
	RetryInterval *int32  `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	SkillGroup    *int64  `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
	Ani           *string `json:"Ani,omitempty" xml:"Ani,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Model         *int32  `json:"Model,omitempty" xml:"Model,omitempty"`
	DepartmentId  *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	ExtAttrs      *string `json:"ExtAttrs,omitempty" xml:"ExtAttrs,omitempty"`
	CallInfos     *string `json:"CallInfos,omitempty" xml:"CallInfos,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateOutboundTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateOutboundTaskRequest) SetTaskType(v int32) *CreateOutboundTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetTaskName(v string) *CreateOutboundTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetDescription(v string) *CreateOutboundTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetStartDate(v string) *CreateOutboundTaskRequest {
	s.StartDate = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetEndDate(v string) *CreateOutboundTaskRequest {
	s.EndDate = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetStartTime(v string) *CreateOutboundTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetEndTime(v string) *CreateOutboundTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetRetryTime(v int32) *CreateOutboundTaskRequest {
	s.RetryTime = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetRetryInterval(v int32) *CreateOutboundTaskRequest {
	s.RetryInterval = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetSkillGroup(v int64) *CreateOutboundTaskRequest {
	s.SkillGroup = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetAni(v string) *CreateOutboundTaskRequest {
	s.Ani = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetGroupName(v string) *CreateOutboundTaskRequest {
	s.GroupName = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetModel(v int32) *CreateOutboundTaskRequest {
	s.Model = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetDepartmentId(v int64) *CreateOutboundTaskRequest {
	s.DepartmentId = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetExtAttrs(v string) *CreateOutboundTaskRequest {
	s.ExtAttrs = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetCallInfos(v string) *CreateOutboundTaskRequest {
	s.CallInfos = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetInstanceId(v string) *CreateOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

type CreateOutboundTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOutboundTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOutboundTaskResponseBody) SetCode(v string) *CreateOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOutboundTaskResponseBody) SetMessage(v string) *CreateOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOutboundTaskResponseBody) SetData(v string) *CreateOutboundTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateOutboundTaskResponseBody) SetRequestId(v string) *CreateOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOutboundTaskResponseBody) SetSuccess(v bool) *CreateOutboundTaskResponseBody {
	s.Success = &v
	return s
}

type CreateOutboundTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOutboundTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateOutboundTaskResponse) SetHeaders(v map[string]*string) *CreateOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateOutboundTaskResponse) SetBody(v *CreateOutboundTaskResponseBody) *CreateOutboundTaskResponse {
	s.Body = v
	return s
}

type CreateQualityProjectRequest struct {
	ProjectName      *string   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	CheckFreqType    *int32    `json:"CheckFreqType,omitempty" xml:"CheckFreqType,omitempty"`
	ScopeType        *int32    `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	TimeRangeStart   *string   `json:"TimeRangeStart,omitempty" xml:"TimeRangeStart,omitempty"`
	TimeRangeEnd     *string   `json:"TimeRangeEnd,omitempty" xml:"TimeRangeEnd,omitempty"`
	InstanceId       *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ChannelTouchType []*int    `json:"ChannelTouchType,omitempty" xml:"ChannelTouchType,omitempty" type:"Repeated"`
	DepList          []*int    `json:"DepList,omitempty" xml:"DepList,omitempty" type:"Repeated"`
	GroupList        []*int    `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	ServicerList     []*string `json:"ServicerList,omitempty" xml:"ServicerList,omitempty" type:"Repeated"`
	AnalysisIds      []*int    `json:"AnalysisIds,omitempty" xml:"AnalysisIds,omitempty" type:"Repeated"`
}

func (s CreateQualityProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateQualityProjectRequest) SetProjectName(v string) *CreateQualityProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateQualityProjectRequest) SetCheckFreqType(v int32) *CreateQualityProjectRequest {
	s.CheckFreqType = &v
	return s
}

func (s *CreateQualityProjectRequest) SetScopeType(v int32) *CreateQualityProjectRequest {
	s.ScopeType = &v
	return s
}

func (s *CreateQualityProjectRequest) SetTimeRangeStart(v string) *CreateQualityProjectRequest {
	s.TimeRangeStart = &v
	return s
}

func (s *CreateQualityProjectRequest) SetTimeRangeEnd(v string) *CreateQualityProjectRequest {
	s.TimeRangeEnd = &v
	return s
}

func (s *CreateQualityProjectRequest) SetInstanceId(v string) *CreateQualityProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateQualityProjectRequest) SetChannelTouchType(v []*int) *CreateQualityProjectRequest {
	s.ChannelTouchType = v
	return s
}

func (s *CreateQualityProjectRequest) SetDepList(v []*int) *CreateQualityProjectRequest {
	s.DepList = v
	return s
}

func (s *CreateQualityProjectRequest) SetGroupList(v []*int) *CreateQualityProjectRequest {
	s.GroupList = v
	return s
}

func (s *CreateQualityProjectRequest) SetServicerList(v []*string) *CreateQualityProjectRequest {
	s.ServicerList = v
	return s
}

func (s *CreateQualityProjectRequest) SetAnalysisIds(v []*int) *CreateQualityProjectRequest {
	s.AnalysisIds = v
	return s
}

type CreateQualityProjectResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *CreateQualityProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CreateQualityProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQualityProjectResponseBody) SetCode(v string) *CreateQualityProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateQualityProjectResponseBody) SetMessage(v string) *CreateQualityProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateQualityProjectResponseBody) SetRequestId(v string) *CreateQualityProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQualityProjectResponseBody) SetSuccess(v bool) *CreateQualityProjectResponseBody {
	s.Success = &v
	return s
}

func (s *CreateQualityProjectResponseBody) SetData(v *CreateQualityProjectResponseBodyData) *CreateQualityProjectResponseBody {
	s.Data = v
	return s
}

type CreateQualityProjectResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Version    *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
	ProjectId  *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateQualityProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateQualityProjectResponseBodyData) SetInstanceId(v string) *CreateQualityProjectResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateQualityProjectResponseBodyData) SetVersion(v int32) *CreateQualityProjectResponseBodyData {
	s.Version = &v
	return s
}

func (s *CreateQualityProjectResponseBodyData) SetProjectId(v int64) *CreateQualityProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

type CreateQualityProjectResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateQualityProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateQualityProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateQualityProjectResponse) SetHeaders(v map[string]*string) *CreateQualityProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateQualityProjectResponse) SetBody(v *CreateQualityProjectResponseBody) *CreateQualityProjectResponse {
	s.Body = v
	return s
}

type CreateQualityRuleRequest struct {
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleTag    *int32    `json:"RuleTag,omitempty" xml:"RuleTag,omitempty"`
	MatchType  *int32    `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	KeyWords   []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
}

func (s CreateQualityRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateQualityRuleRequest) SetInstanceId(v string) *CreateQualityRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateQualityRuleRequest) SetName(v string) *CreateQualityRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateQualityRuleRequest) SetRuleTag(v int32) *CreateQualityRuleRequest {
	s.RuleTag = &v
	return s
}

func (s *CreateQualityRuleRequest) SetMatchType(v int32) *CreateQualityRuleRequest {
	s.MatchType = &v
	return s
}

func (s *CreateQualityRuleRequest) SetKeyWords(v []*string) *CreateQualityRuleRequest {
	s.KeyWords = v
	return s
}

type CreateQualityRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateQualityRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQualityRuleResponseBody) SetCode(v string) *CreateQualityRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateQualityRuleResponseBody) SetMessage(v string) *CreateQualityRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateQualityRuleResponseBody) SetRequestId(v string) *CreateQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQualityRuleResponseBody) SetSuccess(v bool) *CreateQualityRuleResponseBody {
	s.Success = &v
	return s
}

type CreateQualityRuleResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateQualityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateQualityRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateQualityRuleResponse) SetHeaders(v map[string]*string) *CreateQualityRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateQualityRuleResponse) SetBody(v *CreateQualityRuleResponseBody) *CreateQualityRuleResponse {
	s.Body = v
	return s
}

type CreateSkillGroupRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ChannelType    *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 部门ID
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
}

func (s CreateSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupRequest) SetInstanceId(v string) *CreateSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSkillGroupRequest) SetSkillGroupName(v string) *CreateSkillGroupRequest {
	s.SkillGroupName = &v
	return s
}

func (s *CreateSkillGroupRequest) SetDescription(v string) *CreateSkillGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateSkillGroupRequest) SetDisplayName(v string) *CreateSkillGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateSkillGroupRequest) SetChannelType(v int32) *CreateSkillGroupRequest {
	s.ChannelType = &v
	return s
}

func (s *CreateSkillGroupRequest) SetClientToken(v string) *CreateSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSkillGroupRequest) SetDepartmentId(v int64) *CreateSkillGroupRequest {
	s.DepartmentId = &v
	return s
}

type CreateSkillGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupResponseBody) SetMessage(v string) *CreateSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetRequestId(v string) *CreateSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetData(v int64) *CreateSkillGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetCode(v string) *CreateSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetSuccess(v bool) *CreateSkillGroupResponseBody {
	s.Success = &v
	return s
}

type CreateSkillGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupResponse) SetHeaders(v map[string]*string) *CreateSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateSkillGroupResponse) SetBody(v *CreateSkillGroupResponseBody) *CreateSkillGroupResponse {
	s.Body = v
	return s
}

type CreateTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	RobotId              *string `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	Caller               *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	CallString           *string `json:"CallString,omitempty" xml:"CallString,omitempty"`
	CallStringType       *string `json:"CallStringType,omitempty" xml:"CallStringType,omitempty"`
	RetryFlag            *int32  `json:"RetryFlag,omitempty" xml:"RetryFlag,omitempty"`
	RetryCount           *int32  `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	RetryInterval        *int32  `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	RetryStatusCode      *string `json:"RetryStatusCode,omitempty" xml:"RetryStatusCode,omitempty"`
	StartNow             *bool   `json:"StartNow,omitempty" xml:"StartNow,omitempty"`
	WorkTimeList         *string `json:"WorkTimeList,omitempty" xml:"WorkTimeList,omitempty"`
	WorkDay              *string `json:"WorkDay,omitempty" xml:"WorkDay,omitempty"`
	SeatCount            *string `json:"SeatCount,omitempty" xml:"SeatCount,omitempty"`
}

func (s CreateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) SetOwnerId(v int64) *CreateTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTaskRequest) SetResourceOwnerAccount(v string) *CreateTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTaskRequest) SetResourceOwnerId(v int64) *CreateTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTaskRequest) SetTaskName(v string) *CreateTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateTaskRequest) SetRobotId(v string) *CreateTaskRequest {
	s.RobotId = &v
	return s
}

func (s *CreateTaskRequest) SetCaller(v string) *CreateTaskRequest {
	s.Caller = &v
	return s
}

func (s *CreateTaskRequest) SetCallString(v string) *CreateTaskRequest {
	s.CallString = &v
	return s
}

func (s *CreateTaskRequest) SetCallStringType(v string) *CreateTaskRequest {
	s.CallStringType = &v
	return s
}

func (s *CreateTaskRequest) SetRetryFlag(v int32) *CreateTaskRequest {
	s.RetryFlag = &v
	return s
}

func (s *CreateTaskRequest) SetRetryCount(v int32) *CreateTaskRequest {
	s.RetryCount = &v
	return s
}

func (s *CreateTaskRequest) SetRetryInterval(v int32) *CreateTaskRequest {
	s.RetryInterval = &v
	return s
}

func (s *CreateTaskRequest) SetRetryStatusCode(v string) *CreateTaskRequest {
	s.RetryStatusCode = &v
	return s
}

func (s *CreateTaskRequest) SetStartNow(v bool) *CreateTaskRequest {
	s.StartNow = &v
	return s
}

func (s *CreateTaskRequest) SetWorkTimeList(v string) *CreateTaskRequest {
	s.WorkTimeList = &v
	return s
}

func (s *CreateTaskRequest) SetWorkDay(v string) *CreateTaskRequest {
	s.WorkDay = &v
	return s
}

func (s *CreateTaskRequest) SetSeatCount(v string) *CreateTaskRequest {
	s.SeatCount = &v
	return s
}

type CreateTaskResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBody) SetRequestId(v string) *CreateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskResponseBody) SetData(v int64) *CreateTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTaskResponseBody) SetCode(v string) *CreateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskResponseBody) SetMessage(v string) *CreateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskResponseBody) SetSuccess(v bool) *CreateTaskResponseBody {
	s.Success = &v
	return s
}

type CreateTaskResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskResponse) SetHeaders(v map[string]*string) *CreateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskResponse) SetBody(v *CreateTaskResponseBody) *CreateTaskResponse {
	s.Body = v
	return s
}

type CreateThirdSsoAgentRequest struct {
	// clientToken
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// param1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// param2
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// param3
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// param4
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// param5
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// param6
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitempty" xml:"SkillGroupIds,omitempty" type:"Repeated"`
	// param7
	RoleIds []*int64 `json:"RoleIds,omitempty" xml:"RoleIds,omitempty" type:"Repeated"`
}

func (s CreateThirdSsoAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateThirdSsoAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateThirdSsoAgentRequest) SetClientToken(v string) *CreateThirdSsoAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetInstanceId(v string) *CreateThirdSsoAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetClientId(v string) *CreateThirdSsoAgentRequest {
	s.ClientId = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetAccountId(v string) *CreateThirdSsoAgentRequest {
	s.AccountId = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetAccountName(v string) *CreateThirdSsoAgentRequest {
	s.AccountName = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetDisplayName(v string) *CreateThirdSsoAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetSkillGroupIds(v []*int64) *CreateThirdSsoAgentRequest {
	s.SkillGroupIds = v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetRoleIds(v []*int64) *CreateThirdSsoAgentRequest {
	s.RoleIds = v
	return s
}

type CreateThirdSsoAgentResponseBody struct {
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 新创建的坐席id
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateThirdSsoAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateThirdSsoAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateThirdSsoAgentResponseBody) SetMessage(v string) *CreateThirdSsoAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetRequestId(v string) *CreateThirdSsoAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetHttpStatusCode(v int64) *CreateThirdSsoAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetData(v int64) *CreateThirdSsoAgentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetCode(v string) *CreateThirdSsoAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetSuccess(v bool) *CreateThirdSsoAgentResponseBody {
	s.Success = &v
	return s
}

type CreateThirdSsoAgentResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateThirdSsoAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateThirdSsoAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateThirdSsoAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateThirdSsoAgentResponse) SetHeaders(v map[string]*string) *CreateThirdSsoAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateThirdSsoAgentResponse) SetBody(v *CreateThirdSsoAgentResponseBody) *CreateThirdSsoAgentResponse {
	s.Body = v
	return s
}

type DeleteAgentRequest struct {
	// js sdk中自动生成的鉴权token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DeleteAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentRequest) SetClientToken(v string) *DeleteAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAgentRequest) SetInstanceId(v string) *DeleteAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAgentRequest) SetAccountName(v string) *DeleteAgentRequest {
	s.AccountName = &v
	return s
}

type DeleteAgentResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s DeleteAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentResponseBody) SetMessage(v string) *DeleteAgentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAgentResponseBody) SetRequestId(v string) *DeleteAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentResponseBody) SetCode(v string) *DeleteAgentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAgentResponseBody) SetSuccess(v bool) *DeleteAgentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAgentResponseBody) SetHttpStatusCode(v int64) *DeleteAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

type DeleteAgentResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAgentResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentResponse) SetHeaders(v map[string]*string) *DeleteAgentResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentResponse) SetBody(v *DeleteAgentResponseBody) *DeleteAgentResponse {
	s.Body = v
	return s
}

type DeleteDepartmentRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 部门id
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
}

func (s DeleteDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDepartmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentRequest) SetInstanceId(v string) *DeleteDepartmentRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDepartmentRequest) SetDepartmentId(v int64) *DeleteDepartmentRequest {
	s.DepartmentId = &v
	return s
}

type DeleteDepartmentResponseBody struct {
	// Id of the request
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s DeleteDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentResponseBody) SetRequestId(v string) *DeleteDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetSuccess(v bool) *DeleteDepartmentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetCode(v string) *DeleteDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetMessage(v string) *DeleteDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetHttpStatusCode(v int64) *DeleteDepartmentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetData(v bool) *DeleteDepartmentResponseBody {
	s.Data = &v
	return s
}

type DeleteDepartmentResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDepartmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentResponse) SetHeaders(v map[string]*string) *DeleteDepartmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteDepartmentResponse) SetBody(v *DeleteDepartmentResponseBody) *DeleteDepartmentResponse {
	s.Body = v
	return s
}

type DeleteHotlineNumberRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 号码
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
}

func (s DeleteHotlineNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotlineNumberRequest) GoString() string {
	return s.String()
}

func (s *DeleteHotlineNumberRequest) SetInstanceId(v string) *DeleteHotlineNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteHotlineNumberRequest) SetHotlineNumber(v string) *DeleteHotlineNumberRequest {
	s.HotlineNumber = &v
	return s
}

type DeleteHotlineNumberResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 接口调用是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// http状态码
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s DeleteHotlineNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotlineNumberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHotlineNumberResponseBody) SetRequestId(v string) *DeleteHotlineNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHotlineNumberResponseBody) SetSuccess(v bool) *DeleteHotlineNumberResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHotlineNumberResponseBody) SetCode(v string) *DeleteHotlineNumberResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHotlineNumberResponseBody) SetMessage(v string) *DeleteHotlineNumberResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHotlineNumberResponseBody) SetHttpStatusCode(v int64) *DeleteHotlineNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

type DeleteHotlineNumberResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteHotlineNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHotlineNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotlineNumberResponse) GoString() string {
	return s.String()
}

func (s *DeleteHotlineNumberResponse) SetHeaders(v map[string]*string) *DeleteHotlineNumberResponse {
	s.Headers = v
	return s
}

func (s *DeleteHotlineNumberResponse) SetBody(v *DeleteHotlineNumberResponseBody) *DeleteHotlineNumberResponse {
	s.Body = v
	return s
}

type DeleteOutboundTaskRequest struct {
	OutboundTaskId *int64  `json:"OutboundTaskId,omitempty" xml:"OutboundTaskId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteOutboundTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteOutboundTaskRequest) SetOutboundTaskId(v int64) *DeleteOutboundTaskRequest {
	s.OutboundTaskId = &v
	return s
}

func (s *DeleteOutboundTaskRequest) SetInstanceId(v string) *DeleteOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

type DeleteOutboundTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteOutboundTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOutboundTaskResponseBody) SetCode(v string) *DeleteOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteOutboundTaskResponseBody) SetMessage(v string) *DeleteOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteOutboundTaskResponseBody) SetData(v string) *DeleteOutboundTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteOutboundTaskResponseBody) SetRequestId(v string) *DeleteOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOutboundTaskResponseBody) SetSuccess(v bool) *DeleteOutboundTaskResponseBody {
	s.Success = &v
	return s
}

type DeleteOutboundTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteOutboundTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteOutboundTaskResponse) SetHeaders(v map[string]*string) *DeleteOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteOutboundTaskResponse) SetBody(v *DeleteOutboundTaskResponseBody) *DeleteOutboundTaskResponse {
	s.Body = v
	return s
}

type DeleteOuterAccountRequest struct {
	OuterAccountId   *string `json:"OuterAccountId,omitempty" xml:"OuterAccountId,omitempty"`
	OuterAccountType *string `json:"OuterAccountType,omitempty" xml:"OuterAccountType,omitempty"`
}

func (s DeleteOuterAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOuterAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteOuterAccountRequest) SetOuterAccountId(v string) *DeleteOuterAccountRequest {
	s.OuterAccountId = &v
	return s
}

func (s *DeleteOuterAccountRequest) SetOuterAccountType(v string) *DeleteOuterAccountRequest {
	s.OuterAccountType = &v
	return s
}

type DeleteOuterAccountResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteOuterAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOuterAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOuterAccountResponseBody) SetMessage(v string) *DeleteOuterAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteOuterAccountResponseBody) SetRequestId(v string) *DeleteOuterAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOuterAccountResponseBody) SetData(v bool) *DeleteOuterAccountResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteOuterAccountResponseBody) SetCode(v string) *DeleteOuterAccountResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteOuterAccountResponseBody) SetSuccess(v bool) *DeleteOuterAccountResponseBody {
	s.Success = &v
	return s
}

type DeleteOuterAccountResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteOuterAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteOuterAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOuterAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteOuterAccountResponse) SetHeaders(v map[string]*string) *DeleteOuterAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteOuterAccountResponse) SetBody(v *DeleteOuterAccountResponseBody) *DeleteOuterAccountResponse {
	s.Body = v
	return s
}

type DeleteQualityProjectRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProjectId  *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteQualityProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQualityProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityProjectRequest) SetInstanceId(v string) *DeleteQualityProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteQualityProjectRequest) SetProjectId(v int64) *DeleteQualityProjectRequest {
	s.ProjectId = &v
	return s
}

type DeleteQualityProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQualityProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQualityProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualityProjectResponseBody) SetCode(v string) *DeleteQualityProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQualityProjectResponseBody) SetMessage(v string) *DeleteQualityProjectResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQualityProjectResponseBody) SetData(v string) *DeleteQualityProjectResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQualityProjectResponseBody) SetRequestId(v string) *DeleteQualityProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQualityProjectResponseBody) SetSuccess(v bool) *DeleteQualityProjectResponseBody {
	s.Success = &v
	return s
}

type DeleteQualityProjectResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteQualityProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteQualityProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQualityProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteQualityProjectResponse) SetHeaders(v map[string]*string) *DeleteQualityProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteQualityProjectResponse) SetBody(v *DeleteQualityProjectResponseBody) *DeleteQualityProjectResponse {
	s.Body = v
	return s
}

type DeleteQualityRuleRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteQualityRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityRuleRequest) SetInstanceId(v string) *DeleteQualityRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteQualityRuleRequest) SetId(v int64) *DeleteQualityRuleRequest {
	s.Id = &v
	return s
}

type DeleteQualityRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQualityRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualityRuleResponseBody) SetCode(v string) *DeleteQualityRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) SetMessage(v string) *DeleteQualityRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) SetData(v string) *DeleteQualityRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) SetRequestId(v string) *DeleteQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) SetSuccess(v bool) *DeleteQualityRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteQualityRuleResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteQualityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteQualityRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQualityRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteQualityRuleResponse) SetHeaders(v map[string]*string) *DeleteQualityRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteQualityRuleResponse) SetBody(v *DeleteQualityRuleResponseBody) *DeleteQualityRuleResponse {
	s.Body = v
	return s
}

type DeleteSkillGroupRequest struct {
	OuterGroupType *string `json:"OuterGroupType,omitempty" xml:"OuterGroupType,omitempty"`
	OuterGroupId   *string `json:"OuterGroupId,omitempty" xml:"OuterGroupId,omitempty"`
}

func (s DeleteSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupRequest) SetOuterGroupType(v string) *DeleteSkillGroupRequest {
	s.OuterGroupType = &v
	return s
}

func (s *DeleteSkillGroupRequest) SetOuterGroupId(v string) *DeleteSkillGroupRequest {
	s.OuterGroupId = &v
	return s
}

type DeleteSkillGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupResponseBody) SetMessage(v string) *DeleteSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSkillGroupResponseBody) SetRequestId(v string) *DeleteSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSkillGroupResponseBody) SetData(v bool) *DeleteSkillGroupResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSkillGroupResponseBody) SetCode(v string) *DeleteSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSkillGroupResponseBody) SetSuccess(v bool) *DeleteSkillGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteSkillGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupResponse) SetHeaders(v map[string]*string) *DeleteSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSkillGroupResponse) SetBody(v *DeleteSkillGroupResponseBody) *DeleteSkillGroupResponse {
	s.Body = v
	return s
}

type DescribeRecordDataRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountId            *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Acid                 *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	SecLevel             *int32  `json:"SecLevel,omitempty" xml:"SecLevel,omitempty"`
}

func (s DescribeRecordDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordDataRequest) SetOwnerId(v int64) *DescribeRecordDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRecordDataRequest) SetResourceOwnerAccount(v string) *DescribeRecordDataRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRecordDataRequest) SetResourceOwnerId(v int64) *DescribeRecordDataRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRecordDataRequest) SetProdCode(v string) *DescribeRecordDataRequest {
	s.ProdCode = &v
	return s
}

func (s *DescribeRecordDataRequest) SetAccountType(v string) *DescribeRecordDataRequest {
	s.AccountType = &v
	return s
}

func (s *DescribeRecordDataRequest) SetAccountId(v string) *DescribeRecordDataRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeRecordDataRequest) SetAcid(v string) *DescribeRecordDataRequest {
	s.Acid = &v
	return s
}

func (s *DescribeRecordDataRequest) SetSecLevel(v int32) *DescribeRecordDataRequest {
	s.SecLevel = &v
	return s
}

type DescribeRecordDataResponseBody struct {
	OssLink   *string `json:"OssLink,omitempty" xml:"OssLink,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AgentId   *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	Acid      *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeRecordDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordDataResponseBody) SetOssLink(v string) *DescribeRecordDataResponseBody {
	s.OssLink = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetRequestId(v string) *DescribeRecordDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetAgentId(v string) *DescribeRecordDataResponseBody {
	s.AgentId = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetAcid(v string) *DescribeRecordDataResponseBody {
	s.Acid = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetCode(v string) *DescribeRecordDataResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetMessage(v string) *DescribeRecordDataResponseBody {
	s.Message = &v
	return s
}

type DescribeRecordDataResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRecordDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecordDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordDataResponse) SetHeaders(v map[string]*string) *DescribeRecordDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordDataResponse) SetBody(v *DescribeRecordDataResponseBody) *DescribeRecordDataResponse {
	s.Body = v
	return s
}

type EditQualityProjectRequest struct {
	InstanceId       *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProjectId        *int64    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName      *string   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	CheckFreqType    *int32    `json:"CheckFreqType,omitempty" xml:"CheckFreqType,omitempty"`
	ProjectVersion   *int32    `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
	ScopeType        *int32    `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	TimeRangeStart   *string   `json:"TimeRangeStart,omitempty" xml:"TimeRangeStart,omitempty"`
	TimeRangeEnd     *string   `json:"TimeRangeEnd,omitempty" xml:"TimeRangeEnd,omitempty"`
	AnalysisIds      []*int    `json:"AnalysisIds,omitempty" xml:"AnalysisIds,omitempty" type:"Repeated"`
	DepList          []*int    `json:"DepList,omitempty" xml:"DepList,omitempty" type:"Repeated"`
	GroupList        []*int    `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	ServicerList     []*string `json:"ServicerList,omitempty" xml:"ServicerList,omitempty" type:"Repeated"`
	ChannelTouchType []*int    `json:"ChannelTouchType,omitempty" xml:"ChannelTouchType,omitempty" type:"Repeated"`
}

func (s EditQualityProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s EditQualityProjectRequest) GoString() string {
	return s.String()
}

func (s *EditQualityProjectRequest) SetInstanceId(v string) *EditQualityProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *EditQualityProjectRequest) SetProjectId(v int64) *EditQualityProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *EditQualityProjectRequest) SetProjectName(v string) *EditQualityProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *EditQualityProjectRequest) SetCheckFreqType(v int32) *EditQualityProjectRequest {
	s.CheckFreqType = &v
	return s
}

func (s *EditQualityProjectRequest) SetProjectVersion(v int32) *EditQualityProjectRequest {
	s.ProjectVersion = &v
	return s
}

func (s *EditQualityProjectRequest) SetScopeType(v int32) *EditQualityProjectRequest {
	s.ScopeType = &v
	return s
}

func (s *EditQualityProjectRequest) SetTimeRangeStart(v string) *EditQualityProjectRequest {
	s.TimeRangeStart = &v
	return s
}

func (s *EditQualityProjectRequest) SetTimeRangeEnd(v string) *EditQualityProjectRequest {
	s.TimeRangeEnd = &v
	return s
}

func (s *EditQualityProjectRequest) SetAnalysisIds(v []*int) *EditQualityProjectRequest {
	s.AnalysisIds = v
	return s
}

func (s *EditQualityProjectRequest) SetDepList(v []*int) *EditQualityProjectRequest {
	s.DepList = v
	return s
}

func (s *EditQualityProjectRequest) SetGroupList(v []*int) *EditQualityProjectRequest {
	s.GroupList = v
	return s
}

func (s *EditQualityProjectRequest) SetServicerList(v []*string) *EditQualityProjectRequest {
	s.ServicerList = v
	return s
}

func (s *EditQualityProjectRequest) SetChannelTouchType(v []*int) *EditQualityProjectRequest {
	s.ChannelTouchType = v
	return s
}

type EditQualityProjectResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      []*EditQualityProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s EditQualityProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditQualityProjectResponseBody) GoString() string {
	return s.String()
}

func (s *EditQualityProjectResponseBody) SetCode(v string) *EditQualityProjectResponseBody {
	s.Code = &v
	return s
}

func (s *EditQualityProjectResponseBody) SetMessage(v string) *EditQualityProjectResponseBody {
	s.Message = &v
	return s
}

func (s *EditQualityProjectResponseBody) SetRequestId(v string) *EditQualityProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditQualityProjectResponseBody) SetSuccess(v bool) *EditQualityProjectResponseBody {
	s.Success = &v
	return s
}

func (s *EditQualityProjectResponseBody) SetData(v []*EditQualityProjectResponseBodyData) *EditQualityProjectResponseBody {
	s.Data = v
	return s
}

type EditQualityProjectResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Version    *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
	ProjectId  *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s EditQualityProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EditQualityProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *EditQualityProjectResponseBodyData) SetInstanceId(v string) *EditQualityProjectResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *EditQualityProjectResponseBodyData) SetVersion(v int32) *EditQualityProjectResponseBodyData {
	s.Version = &v
	return s
}

func (s *EditQualityProjectResponseBodyData) SetProjectId(v int64) *EditQualityProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

type EditQualityProjectResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditQualityProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditQualityProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s EditQualityProjectResponse) GoString() string {
	return s.String()
}

func (s *EditQualityProjectResponse) SetHeaders(v map[string]*string) *EditQualityProjectResponse {
	s.Headers = v
	return s
}

func (s *EditQualityProjectResponse) SetBody(v *EditQualityProjectResponseBody) *EditQualityProjectResponse {
	s.Body = v
	return s
}

type EditQualityRuleRequest struct {
	InstanceId    *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name          *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleTag       *int32    `json:"RuleTag,omitempty" xml:"RuleTag,omitempty"`
	MatchType     *int32    `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	QualityRuleId *int64    `json:"QualityRuleId,omitempty" xml:"QualityRuleId,omitempty"`
	KeyWords      []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
}

func (s EditQualityRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EditQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *EditQualityRuleRequest) SetInstanceId(v string) *EditQualityRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *EditQualityRuleRequest) SetName(v string) *EditQualityRuleRequest {
	s.Name = &v
	return s
}

func (s *EditQualityRuleRequest) SetRuleTag(v int32) *EditQualityRuleRequest {
	s.RuleTag = &v
	return s
}

func (s *EditQualityRuleRequest) SetMatchType(v int32) *EditQualityRuleRequest {
	s.MatchType = &v
	return s
}

func (s *EditQualityRuleRequest) SetQualityRuleId(v int64) *EditQualityRuleRequest {
	s.QualityRuleId = &v
	return s
}

func (s *EditQualityRuleRequest) SetKeyWords(v []*string) *EditQualityRuleRequest {
	s.KeyWords = v
	return s
}

type EditQualityRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditQualityRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EditQualityRuleResponseBody) SetCode(v string) *EditQualityRuleResponseBody {
	s.Code = &v
	return s
}

func (s *EditQualityRuleResponseBody) SetMessage(v string) *EditQualityRuleResponseBody {
	s.Message = &v
	return s
}

func (s *EditQualityRuleResponseBody) SetRequestId(v string) *EditQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditQualityRuleResponseBody) SetSuccess(v bool) *EditQualityRuleResponseBody {
	s.Success = &v
	return s
}

type EditQualityRuleResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditQualityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditQualityRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EditQualityRuleResponse) GoString() string {
	return s.String()
}

func (s *EditQualityRuleResponse) SetHeaders(v map[string]*string) *EditQualityRuleResponse {
	s.Headers = v
	return s
}

func (s *EditQualityRuleResponse) SetBody(v *EditQualityRuleResponseBody) *EditQualityRuleResponse {
	s.Body = v
	return s
}

type EditQualityRuleTagRequest struct {
	InstanceId    *string                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AnalysisTypes []*EditQualityRuleTagRequestAnalysisTypes `json:"AnalysisTypes,omitempty" xml:"AnalysisTypes,omitempty" type:"Repeated"`
}

func (s EditQualityRuleTagRequest) String() string {
	return tea.Prettify(s)
}

func (s EditQualityRuleTagRequest) GoString() string {
	return s.String()
}

func (s *EditQualityRuleTagRequest) SetInstanceId(v string) *EditQualityRuleTagRequest {
	s.InstanceId = &v
	return s
}

func (s *EditQualityRuleTagRequest) SetAnalysisTypes(v []*EditQualityRuleTagRequestAnalysisTypes) *EditQualityRuleTagRequest {
	s.AnalysisTypes = v
	return s
}

type EditQualityRuleTagRequestAnalysisTypes struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s EditQualityRuleTagRequestAnalysisTypes) String() string {
	return tea.Prettify(s)
}

func (s EditQualityRuleTagRequestAnalysisTypes) GoString() string {
	return s.String()
}

func (s *EditQualityRuleTagRequestAnalysisTypes) SetName(v string) *EditQualityRuleTagRequestAnalysisTypes {
	s.Name = &v
	return s
}

func (s *EditQualityRuleTagRequestAnalysisTypes) SetId(v int64) *EditQualityRuleTagRequestAnalysisTypes {
	s.Id = &v
	return s
}

type EditQualityRuleTagResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditQualityRuleTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditQualityRuleTagResponseBody) GoString() string {
	return s.String()
}

func (s *EditQualityRuleTagResponseBody) SetCode(v string) *EditQualityRuleTagResponseBody {
	s.Code = &v
	return s
}

func (s *EditQualityRuleTagResponseBody) SetMessage(v string) *EditQualityRuleTagResponseBody {
	s.Message = &v
	return s
}

func (s *EditQualityRuleTagResponseBody) SetData(v string) *EditQualityRuleTagResponseBody {
	s.Data = &v
	return s
}

func (s *EditQualityRuleTagResponseBody) SetRequestId(v string) *EditQualityRuleTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditQualityRuleTagResponseBody) SetSuccess(v bool) *EditQualityRuleTagResponseBody {
	s.Success = &v
	return s
}

type EditQualityRuleTagResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditQualityRuleTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditQualityRuleTagResponse) String() string {
	return tea.Prettify(s)
}

func (s EditQualityRuleTagResponse) GoString() string {
	return s.String()
}

func (s *EditQualityRuleTagResponse) SetHeaders(v map[string]*string) *EditQualityRuleTagResponse {
	s.Headers = v
	return s
}

func (s *EditQualityRuleTagResponse) SetBody(v *EditQualityRuleTagResponseBody) *EditQualityRuleTagResponse {
	s.Body = v
	return s
}

type EncryptPhoneNumRequest struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 号码明文
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
}

func (s EncryptPhoneNumRequest) String() string {
	return tea.Prettify(s)
}

func (s EncryptPhoneNumRequest) GoString() string {
	return s.String()
}

func (s *EncryptPhoneNumRequest) SetInstanceId(v string) *EncryptPhoneNumRequest {
	s.InstanceId = &v
	return s
}

func (s *EncryptPhoneNumRequest) SetPhoneNum(v string) *EncryptPhoneNumRequest {
	s.PhoneNum = &v
	return s
}

type EncryptPhoneNumResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 接口调用是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 加密后密文
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s EncryptPhoneNumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EncryptPhoneNumResponseBody) GoString() string {
	return s.String()
}

func (s *EncryptPhoneNumResponseBody) SetRequestId(v string) *EncryptPhoneNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *EncryptPhoneNumResponseBody) SetSuccess(v bool) *EncryptPhoneNumResponseBody {
	s.Success = &v
	return s
}

func (s *EncryptPhoneNumResponseBody) SetCode(v string) *EncryptPhoneNumResponseBody {
	s.Code = &v
	return s
}

func (s *EncryptPhoneNumResponseBody) SetMessage(v string) *EncryptPhoneNumResponseBody {
	s.Message = &v
	return s
}

func (s *EncryptPhoneNumResponseBody) SetData(v string) *EncryptPhoneNumResponseBody {
	s.Data = &v
	return s
}

type EncryptPhoneNumResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EncryptPhoneNumResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EncryptPhoneNumResponse) String() string {
	return tea.Prettify(s)
}

func (s EncryptPhoneNumResponse) GoString() string {
	return s.String()
}

func (s *EncryptPhoneNumResponse) SetHeaders(v map[string]*string) *EncryptPhoneNumResponse {
	s.Headers = v
	return s
}

func (s *EncryptPhoneNumResponse) SetBody(v *EncryptPhoneNumResponseBody) *EncryptPhoneNumResponse {
	s.Body = v
	return s
}

type FetchCallRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId           *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId     *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
}

func (s FetchCallRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchCallRequest) GoString() string {
	return s.String()
}

func (s *FetchCallRequest) SetClientToken(v string) *FetchCallRequest {
	s.ClientToken = &v
	return s
}

func (s *FetchCallRequest) SetInstanceId(v string) *FetchCallRequest {
	s.InstanceId = &v
	return s
}

func (s *FetchCallRequest) SetAccountName(v string) *FetchCallRequest {
	s.AccountName = &v
	return s
}

func (s *FetchCallRequest) SetCallId(v string) *FetchCallRequest {
	s.CallId = &v
	return s
}

func (s *FetchCallRequest) SetJobId(v string) *FetchCallRequest {
	s.JobId = &v
	return s
}

func (s *FetchCallRequest) SetConnectionId(v string) *FetchCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *FetchCallRequest) SetHoldConnectionId(v string) *FetchCallRequest {
	s.HoldConnectionId = &v
	return s
}

type FetchCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FetchCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FetchCallResponseBody) GoString() string {
	return s.String()
}

func (s *FetchCallResponseBody) SetCode(v string) *FetchCallResponseBody {
	s.Code = &v
	return s
}

func (s *FetchCallResponseBody) SetMessage(v string) *FetchCallResponseBody {
	s.Message = &v
	return s
}

func (s *FetchCallResponseBody) SetRequestId(v string) *FetchCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchCallResponseBody) SetSuccess(v bool) *FetchCallResponseBody {
	s.Success = &v
	return s
}

type FetchCallResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FetchCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FetchCallResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchCallResponse) GoString() string {
	return s.String()
}

func (s *FetchCallResponse) SetHeaders(v map[string]*string) *FetchCallResponse {
	s.Headers = v
	return s
}

func (s *FetchCallResponse) SetBody(v *FetchCallResponseBody) *FetchCallResponse {
	s.Body = v
	return s
}

type FinishHotlineServiceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s FinishHotlineServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishHotlineServiceRequest) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceRequest) SetClientToken(v string) *FinishHotlineServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *FinishHotlineServiceRequest) SetInstanceId(v string) *FinishHotlineServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *FinishHotlineServiceRequest) SetAccountName(v string) *FinishHotlineServiceRequest {
	s.AccountName = &v
	return s
}

type FinishHotlineServiceResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s FinishHotlineServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FinishHotlineServiceResponseBody) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceResponseBody) SetMessage(v string) *FinishHotlineServiceResponseBody {
	s.Message = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) SetRequestId(v string) *FinishHotlineServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) SetCode(v string) *FinishHotlineServiceResponseBody {
	s.Code = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) SetSuccess(v bool) *FinishHotlineServiceResponseBody {
	s.Success = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) SetHttpStatusCode(v int64) *FinishHotlineServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

type FinishHotlineServiceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FinishHotlineServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FinishHotlineServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishHotlineServiceResponse) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceResponse) SetHeaders(v map[string]*string) *FinishHotlineServiceResponse {
	s.Headers = v
	return s
}

func (s *FinishHotlineServiceResponse) SetBody(v *FinishHotlineServiceResponseBody) *FinishHotlineServiceResponse {
	s.Body = v
	return s
}

type GenerateWebSocketSignRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GenerateWebSocketSignRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateWebSocketSignRequest) GoString() string {
	return s.String()
}

func (s *GenerateWebSocketSignRequest) SetClientToken(v string) *GenerateWebSocketSignRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateWebSocketSignRequest) SetInstanceId(v string) *GenerateWebSocketSignRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateWebSocketSignRequest) SetAccountName(v string) *GenerateWebSocketSignRequest {
	s.AccountName = &v
	return s
}

type GenerateWebSocketSignResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GenerateWebSocketSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateWebSocketSignResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateWebSocketSignResponseBody) SetMessage(v string) *GenerateWebSocketSignResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetRequestId(v string) *GenerateWebSocketSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetData(v string) *GenerateWebSocketSignResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetCode(v string) *GenerateWebSocketSignResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetSuccess(v bool) *GenerateWebSocketSignResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetHttpStatusCode(v int64) *GenerateWebSocketSignResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GenerateWebSocketSignResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateWebSocketSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateWebSocketSignResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateWebSocketSignResponse) GoString() string {
	return s.String()
}

func (s *GenerateWebSocketSignResponse) SetHeaders(v map[string]*string) *GenerateWebSocketSignResponse {
	s.Headers = v
	return s
}

func (s *GenerateWebSocketSignResponse) SetBody(v *GenerateWebSocketSignResponseBody) *GenerateWebSocketSignResponse {
	s.Body = v
	return s
}

type GetAgentRequest struct {
	// js sdk中自动生成的鉴权token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRequest) SetClientToken(v string) *GetAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *GetAgentRequest) SetInstanceId(v string) *GetAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentRequest) SetAccountName(v string) *GetAgentRequest {
	s.AccountName = &v
	return s
}

type GetAgentResponseBody struct {
	RequestId      *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code           *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GetAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBody) SetRequestId(v string) *GetAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentResponseBody) SetData(v *GetAgentResponseBodyData) *GetAgentResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentResponseBody) SetCode(v string) *GetAgentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentResponseBody) SetSuccess(v bool) *GetAgentResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentResponseBody) SetHttpStatusCode(v int64) *GetAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GetAgentResponseBodyData struct {
	Status      *int32                               `json:"Status,omitempty" xml:"Status,omitempty"`
	DisplayName *string                              `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	AgentId     *int64                               `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	GroupList   []*GetAgentResponseBodyDataGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	AccountName *string                              `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	TenantId    *int64                               `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetAgentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyData) SetStatus(v int32) *GetAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAgentResponseBodyData) SetDisplayName(v string) *GetAgentResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetAgentResponseBodyData) SetAgentId(v int64) *GetAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetAgentResponseBodyData) SetGroupList(v []*GetAgentResponseBodyDataGroupList) *GetAgentResponseBodyData {
	s.GroupList = v
	return s
}

func (s *GetAgentResponseBodyData) SetAccountName(v string) *GetAgentResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *GetAgentResponseBodyData) SetTenantId(v int64) *GetAgentResponseBodyData {
	s.TenantId = &v
	return s
}

type GetAgentResponseBodyDataGroupList struct {
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ChannelType  *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	SkillGroupId *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAgentResponseBodyDataGroupList) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponseBodyDataGroupList) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyDataGroupList) SetDisplayName(v string) *GetAgentResponseBodyDataGroupList {
	s.DisplayName = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetDescription(v string) *GetAgentResponseBodyDataGroupList {
	s.Description = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetChannelType(v int32) *GetAgentResponseBodyDataGroupList {
	s.ChannelType = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetSkillGroupId(v int64) *GetAgentResponseBodyDataGroupList {
	s.SkillGroupId = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetName(v string) *GetAgentResponseBodyDataGroupList {
	s.Name = &v
	return s
}

type GetAgentResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponse) GoString() string {
	return s.String()
}

func (s *GetAgentResponse) SetHeaders(v map[string]*string) *GetAgentResponse {
	s.Headers = v
	return s
}

func (s *GetAgentResponse) SetBody(v *GetAgentResponseBody) *GetAgentResponse {
	s.Body = v
	return s
}

type GetAgentBasisStatusRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// 部门id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
}

func (s GetAgentBasisStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentBasisStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAgentBasisStatusRequest) SetInstanceId(v string) *GetAgentBasisStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentBasisStatusRequest) SetStartDate(v int64) *GetAgentBasisStatusRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentBasisStatusRequest) SetEndDate(v int64) *GetAgentBasisStatusRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentBasisStatusRequest) SetPageSize(v int32) *GetAgentBasisStatusRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentBasisStatusRequest) SetCurrentPage(v int32) *GetAgentBasisStatusRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentBasisStatusRequest) SetAgentIds(v []*int64) *GetAgentBasisStatusRequest {
	s.AgentIds = v
	return s
}

func (s *GetAgentBasisStatusRequest) SetDepIds(v []*int64) *GetAgentBasisStatusRequest {
	s.DepIds = v
	return s
}

type GetAgentBasisStatusShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
}

func (s GetAgentBasisStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentBasisStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAgentBasisStatusShrinkRequest) SetInstanceId(v string) *GetAgentBasisStatusShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentBasisStatusShrinkRequest) SetStartDate(v int64) *GetAgentBasisStatusShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentBasisStatusShrinkRequest) SetEndDate(v int64) *GetAgentBasisStatusShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentBasisStatusShrinkRequest) SetPageSize(v int32) *GetAgentBasisStatusShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentBasisStatusShrinkRequest) SetCurrentPage(v int32) *GetAgentBasisStatusShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentBasisStatusShrinkRequest) SetAgentIdsShrink(v string) *GetAgentBasisStatusShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetAgentBasisStatusShrinkRequest) SetDepIdsShrink(v string) *GetAgentBasisStatusShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

type GetAgentBasisStatusResponseBody struct {
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 调用接口是否成功
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// data
	Data *GetAgentBasisStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetAgentBasisStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentBasisStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentBasisStatusResponseBody) SetRequestId(v string) *GetAgentBasisStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentBasisStatusResponseBody) SetMessage(v string) *GetAgentBasisStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentBasisStatusResponseBody) SetCode(v string) *GetAgentBasisStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentBasisStatusResponseBody) SetSuccess(v string) *GetAgentBasisStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentBasisStatusResponseBody) SetData(v *GetAgentBasisStatusResponseBodyData) *GetAgentBasisStatusResponseBody {
	s.Data = v
	return s
}

type GetAgentBasisStatusResponseBodyData struct {
	// 当前页数
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总记录数
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s GetAgentBasisStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgentBasisStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentBasisStatusResponseBodyData) SetPageNum(v int32) *GetAgentBasisStatusResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetAgentBasisStatusResponseBodyData) SetPageSize(v int32) *GetAgentBasisStatusResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAgentBasisStatusResponseBodyData) SetTotalNum(v int32) *GetAgentBasisStatusResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetAgentBasisStatusResponseBodyData) SetRows(v string) *GetAgentBasisStatusResponseBodyData {
	s.Rows = &v
	return s
}

type GetAgentBasisStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAgentBasisStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentBasisStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentBasisStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAgentBasisStatusResponse) SetHeaders(v map[string]*string) *GetAgentBasisStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAgentBasisStatusResponse) SetBody(v *GetAgentBasisStatusResponseBody) *GetAgentBasisStatusResponse {
	s.Body = v
	return s
}

type GetAgentByIdRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AgentId    *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
}

func (s GetAgentByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentByIdRequest) GoString() string {
	return s.String()
}

func (s *GetAgentByIdRequest) SetInstanceId(v string) *GetAgentByIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentByIdRequest) SetAgentId(v int64) *GetAgentByIdRequest {
	s.AgentId = &v
	return s
}

type GetAgentByIdResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *GetAgentByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetAgentByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentByIdResponseBody) SetCode(v string) *GetAgentByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentByIdResponseBody) SetMessage(v string) *GetAgentByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentByIdResponseBody) SetRequestId(v string) *GetAgentByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentByIdResponseBody) SetSuccess(v bool) *GetAgentByIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentByIdResponseBody) SetData(v *GetAgentByIdResponseBodyData) *GetAgentByIdResponseBody {
	s.Data = v
	return s
}

type GetAgentByIdResponseBodyData struct {
	ShowName       *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	ForeignKey     *string `json:"ForeignKey,omitempty" xml:"ForeignKey,omitempty"`
	ServicerType   *int32  `json:"ServicerType,omitempty" xml:"ServicerType,omitempty"`
	RealName       *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	AgentId        *int32  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	ForeignNick    *string `json:"ForeignNick,omitempty" xml:"ForeignNick,omitempty"`
}

func (s GetAgentByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgentByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentByIdResponseBodyData) SetShowName(v string) *GetAgentByIdResponseBodyData {
	s.ShowName = &v
	return s
}

func (s *GetAgentByIdResponseBodyData) SetForeignKey(v string) *GetAgentByIdResponseBodyData {
	s.ForeignKey = &v
	return s
}

func (s *GetAgentByIdResponseBodyData) SetServicerType(v int32) *GetAgentByIdResponseBodyData {
	s.ServicerType = &v
	return s
}

func (s *GetAgentByIdResponseBodyData) SetRealName(v string) *GetAgentByIdResponseBodyData {
	s.RealName = &v
	return s
}

func (s *GetAgentByIdResponseBodyData) SetCreateUserName(v string) *GetAgentByIdResponseBodyData {
	s.CreateUserName = &v
	return s
}

func (s *GetAgentByIdResponseBodyData) SetAgentId(v int32) *GetAgentByIdResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetAgentByIdResponseBodyData) SetForeignNick(v string) *GetAgentByIdResponseBodyData {
	s.ForeignNick = &v
	return s
}

type GetAgentByIdResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAgentByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentByIdResponse) GoString() string {
	return s.String()
}

func (s *GetAgentByIdResponse) SetHeaders(v map[string]*string) *GetAgentByIdResponse {
	s.Headers = v
	return s
}

func (s *GetAgentByIdResponse) SetBody(v *GetAgentByIdResponseBody) *GetAgentByIdResponse {
	s.Body = v
	return s
}

type GetAgentDetailReportRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// 部门id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// 时间纬度类型
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
	// 是否根据坐席分组显示
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// 是否根据部门分组显示
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
}

func (s GetAgentDetailReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentDetailReportRequest) GoString() string {
	return s.String()
}

func (s *GetAgentDetailReportRequest) SetInstanceId(v string) *GetAgentDetailReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentDetailReportRequest) SetStartDate(v int64) *GetAgentDetailReportRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentDetailReportRequest) SetEndDate(v int64) *GetAgentDetailReportRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentDetailReportRequest) SetPageSize(v int32) *GetAgentDetailReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentDetailReportRequest) SetCurrentPage(v int32) *GetAgentDetailReportRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentDetailReportRequest) SetAgentIds(v []*int64) *GetAgentDetailReportRequest {
	s.AgentIds = v
	return s
}

func (s *GetAgentDetailReportRequest) SetDepIds(v []*int64) *GetAgentDetailReportRequest {
	s.DepIds = v
	return s
}

func (s *GetAgentDetailReportRequest) SetTimeLatitudeType(v string) *GetAgentDetailReportRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetAgentDetailReportRequest) SetExistAgentGrouping(v bool) *GetAgentDetailReportRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetAgentDetailReportRequest) SetExistDepartmentGrouping(v bool) *GetAgentDetailReportRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

type GetAgentDetailReportShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// 时间纬度类型
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
	// 是否根据坐席分组显示
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// 是否根据部门分组显示
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
}

func (s GetAgentDetailReportShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentDetailReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAgentDetailReportShrinkRequest) SetInstanceId(v string) *GetAgentDetailReportShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetStartDate(v int64) *GetAgentDetailReportShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetEndDate(v int64) *GetAgentDetailReportShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetPageSize(v int32) *GetAgentDetailReportShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetCurrentPage(v int32) *GetAgentDetailReportShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetAgentIdsShrink(v string) *GetAgentDetailReportShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetDepIdsShrink(v string) *GetAgentDetailReportShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetTimeLatitudeType(v string) *GetAgentDetailReportShrinkRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetExistAgentGrouping(v bool) *GetAgentDetailReportShrinkRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetExistDepartmentGrouping(v bool) *GetAgentDetailReportShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

type GetAgentDetailReportResponseBody struct {
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 调用接口是否成功
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// data
	Data *GetAgentDetailReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetAgentDetailReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentDetailReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentDetailReportResponseBody) SetRequestId(v string) *GetAgentDetailReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentDetailReportResponseBody) SetMessage(v string) *GetAgentDetailReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentDetailReportResponseBody) SetCode(v string) *GetAgentDetailReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentDetailReportResponseBody) SetSuccess(v string) *GetAgentDetailReportResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentDetailReportResponseBody) SetData(v *GetAgentDetailReportResponseBodyData) *GetAgentDetailReportResponseBody {
	s.Data = v
	return s
}

type GetAgentDetailReportResponseBodyData struct {
	// 当前页数
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 页大小
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总记录数
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s GetAgentDetailReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgentDetailReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentDetailReportResponseBodyData) SetPageNum(v int64) *GetAgentDetailReportResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetAgentDetailReportResponseBodyData) SetPageSize(v int64) *GetAgentDetailReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAgentDetailReportResponseBodyData) SetTotalNum(v int64) *GetAgentDetailReportResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetAgentDetailReportResponseBodyData) SetRows(v string) *GetAgentDetailReportResponseBodyData {
	s.Rows = &v
	return s
}

type GetAgentDetailReportResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAgentDetailReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentDetailReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentDetailReportResponse) GoString() string {
	return s.String()
}

func (s *GetAgentDetailReportResponse) SetHeaders(v map[string]*string) *GetAgentDetailReportResponse {
	s.Headers = v
	return s
}

func (s *GetAgentDetailReportResponse) SetBody(v *GetAgentDetailReportResponseBody) *GetAgentDetailReportResponse {
	s.Body = v
	return s
}

type GetAgentIndexRealTimeRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DepIds      []*int  `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	GroupIds    []*int  `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
}

func (s GetAgentIndexRealTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentIndexRealTimeRequest) GoString() string {
	return s.String()
}

func (s *GetAgentIndexRealTimeRequest) SetInstanceId(v string) *GetAgentIndexRealTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentIndexRealTimeRequest) SetPageSize(v int32) *GetAgentIndexRealTimeRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentIndexRealTimeRequest) SetCurrentPage(v int32) *GetAgentIndexRealTimeRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentIndexRealTimeRequest) SetDepIds(v []*int) *GetAgentIndexRealTimeRequest {
	s.DepIds = v
	return s
}

func (s *GetAgentIndexRealTimeRequest) SetGroupIds(v []*int) *GetAgentIndexRealTimeRequest {
	s.GroupIds = v
	return s
}

type GetAgentIndexRealTimeResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *GetAgentIndexRealTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetAgentIndexRealTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentIndexRealTimeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentIndexRealTimeResponseBody) SetCode(v string) *GetAgentIndexRealTimeResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBody) SetMessage(v string) *GetAgentIndexRealTimeResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBody) SetRequestId(v string) *GetAgentIndexRealTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBody) SetSuccess(v bool) *GetAgentIndexRealTimeResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBody) SetData(v *GetAgentIndexRealTimeResponseBodyData) *GetAgentIndexRealTimeResponseBody {
	s.Data = v
	return s
}

type GetAgentIndexRealTimeResponseBodyData struct {
	PageSize *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                          `json:"Total,omitempty" xml:"Total,omitempty"`
	Page     *int32                                          `json:"Page,omitempty" xml:"Page,omitempty"`
	Columns  []*GetAgentIndexRealTimeResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Rows     []map[string]interface{}                        `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s GetAgentIndexRealTimeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgentIndexRealTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentIndexRealTimeResponseBodyData) SetPageSize(v int32) *GetAgentIndexRealTimeResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBodyData) SetTotal(v int32) *GetAgentIndexRealTimeResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBodyData) SetPage(v int32) *GetAgentIndexRealTimeResponseBodyData {
	s.Page = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBodyData) SetColumns(v []*GetAgentIndexRealTimeResponseBodyDataColumns) *GetAgentIndexRealTimeResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetAgentIndexRealTimeResponseBodyData) SetRows(v []map[string]interface{}) *GetAgentIndexRealTimeResponseBodyData {
	s.Rows = v
	return s
}

type GetAgentIndexRealTimeResponseBodyDataColumns struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetAgentIndexRealTimeResponseBodyDataColumns) String() string {
	return tea.Prettify(s)
}

func (s GetAgentIndexRealTimeResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *GetAgentIndexRealTimeResponseBodyDataColumns) SetKey(v string) *GetAgentIndexRealTimeResponseBodyDataColumns {
	s.Key = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBodyDataColumns) SetTitle(v string) *GetAgentIndexRealTimeResponseBodyDataColumns {
	s.Title = &v
	return s
}

type GetAgentIndexRealTimeResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAgentIndexRealTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentIndexRealTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentIndexRealTimeResponse) GoString() string {
	return s.String()
}

func (s *GetAgentIndexRealTimeResponse) SetHeaders(v map[string]*string) *GetAgentIndexRealTimeResponse {
	s.Headers = v
	return s
}

func (s *GetAgentIndexRealTimeResponse) SetBody(v *GetAgentIndexRealTimeResponseBody) *GetAgentIndexRealTimeResponse {
	s.Body = v
	return s
}

type GetAgentServiceStatusRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// 部门id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// 时间纬度类型
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据坐席分组
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
}

func (s GetAgentServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAgentServiceStatusRequest) SetInstanceId(v string) *GetAgentServiceStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentServiceStatusRequest) SetStartDate(v int64) *GetAgentServiceStatusRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentServiceStatusRequest) SetEndDate(v int64) *GetAgentServiceStatusRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentServiceStatusRequest) SetPageSize(v int32) *GetAgentServiceStatusRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentServiceStatusRequest) SetCurrentPage(v int32) *GetAgentServiceStatusRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentServiceStatusRequest) SetAgentIds(v []*int64) *GetAgentServiceStatusRequest {
	s.AgentIds = v
	return s
}

func (s *GetAgentServiceStatusRequest) SetDepIds(v []*int64) *GetAgentServiceStatusRequest {
	s.DepIds = v
	return s
}

func (s *GetAgentServiceStatusRequest) SetTimeLatitudeType(v string) *GetAgentServiceStatusRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetAgentServiceStatusRequest) SetExistDepartmentGrouping(v bool) *GetAgentServiceStatusRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetAgentServiceStatusRequest) SetExistAgentGrouping(v bool) *GetAgentServiceStatusRequest {
	s.ExistAgentGrouping = &v
	return s
}

type GetAgentServiceStatusShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// 时间纬度类型
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据坐席分组
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
}

func (s GetAgentServiceStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentServiceStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAgentServiceStatusShrinkRequest) SetInstanceId(v string) *GetAgentServiceStatusShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetStartDate(v int64) *GetAgentServiceStatusShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetEndDate(v int64) *GetAgentServiceStatusShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetPageSize(v int32) *GetAgentServiceStatusShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetCurrentPage(v int32) *GetAgentServiceStatusShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetAgentIdsShrink(v string) *GetAgentServiceStatusShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetDepIdsShrink(v string) *GetAgentServiceStatusShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetTimeLatitudeType(v string) *GetAgentServiceStatusShrinkRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetExistDepartmentGrouping(v bool) *GetAgentServiceStatusShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetAgentServiceStatusShrinkRequest) SetExistAgentGrouping(v bool) *GetAgentServiceStatusShrinkRequest {
	s.ExistAgentGrouping = &v
	return s
}

type GetAgentServiceStatusResponseBody struct {
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *GetAgentServiceStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 调用接口是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentServiceStatusResponseBody) SetMessage(v string) *GetAgentServiceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentServiceStatusResponseBody) SetRequestId(v string) *GetAgentServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentServiceStatusResponseBody) SetData(v *GetAgentServiceStatusResponseBodyData) *GetAgentServiceStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentServiceStatusResponseBody) SetCode(v string) *GetAgentServiceStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentServiceStatusResponseBody) SetSuccess(v bool) *GetAgentServiceStatusResponseBody {
	s.Success = &v
	return s
}

type GetAgentServiceStatusResponseBodyData struct {
	// 当前页数
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 页大小
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总记录数
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s GetAgentServiceStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgentServiceStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentServiceStatusResponseBodyData) SetPageNum(v int64) *GetAgentServiceStatusResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetAgentServiceStatusResponseBodyData) SetPageSize(v int64) *GetAgentServiceStatusResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAgentServiceStatusResponseBodyData) SetTotalNum(v int64) *GetAgentServiceStatusResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetAgentServiceStatusResponseBodyData) SetRows(v string) *GetAgentServiceStatusResponseBodyData {
	s.Rows = &v
	return s
}

type GetAgentServiceStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAgentServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAgentServiceStatusResponse) SetHeaders(v map[string]*string) *GetAgentServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAgentServiceStatusResponse) SetBody(v *GetAgentServiceStatusResponseBody) *GetAgentServiceStatusResponse {
	s.Body = v
	return s
}

type GetAgentStatisticsRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// 部门id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// 时间纬度类型
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据坐席分组
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
}

func (s GetAgentStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetAgentStatisticsRequest) SetInstanceId(v string) *GetAgentStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentStatisticsRequest) SetStartDate(v int64) *GetAgentStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentStatisticsRequest) SetEndDate(v int64) *GetAgentStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentStatisticsRequest) SetPageSize(v int32) *GetAgentStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentStatisticsRequest) SetCurrentPage(v int32) *GetAgentStatisticsRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentStatisticsRequest) SetAgentIds(v []*int64) *GetAgentStatisticsRequest {
	s.AgentIds = v
	return s
}

func (s *GetAgentStatisticsRequest) SetDepIds(v []*int64) *GetAgentStatisticsRequest {
	s.DepIds = v
	return s
}

func (s *GetAgentStatisticsRequest) SetTimeLatitudeType(v string) *GetAgentStatisticsRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetAgentStatisticsRequest) SetExistDepartmentGrouping(v bool) *GetAgentStatisticsRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetAgentStatisticsRequest) SetExistAgentGrouping(v bool) *GetAgentStatisticsRequest {
	s.ExistAgentGrouping = &v
	return s
}

type GetAgentStatisticsShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// 时间纬度类型
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据坐席分组
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
}

func (s GetAgentStatisticsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentStatisticsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAgentStatisticsShrinkRequest) SetInstanceId(v string) *GetAgentStatisticsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetStartDate(v int64) *GetAgentStatisticsShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetEndDate(v int64) *GetAgentStatisticsShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetPageSize(v int32) *GetAgentStatisticsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetCurrentPage(v int32) *GetAgentStatisticsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetAgentIdsShrink(v string) *GetAgentStatisticsShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetDepIdsShrink(v string) *GetAgentStatisticsShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetTimeLatitudeType(v string) *GetAgentStatisticsShrinkRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetExistDepartmentGrouping(v bool) *GetAgentStatisticsShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetExistAgentGrouping(v bool) *GetAgentStatisticsShrinkRequest {
	s.ExistAgentGrouping = &v
	return s
}

type GetAgentStatisticsResponseBody struct {
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 调用接口是否成功
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// data
	Data *GetAgentStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetAgentStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentStatisticsResponseBody) SetRequestId(v string) *GetAgentStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentStatisticsResponseBody) SetMessage(v string) *GetAgentStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentStatisticsResponseBody) SetCode(v string) *GetAgentStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentStatisticsResponseBody) SetSuccess(v string) *GetAgentStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentStatisticsResponseBody) SetData(v *GetAgentStatisticsResponseBodyData) *GetAgentStatisticsResponseBody {
	s.Data = v
	return s
}

type GetAgentStatisticsResponseBodyData struct {
	// 当前页数
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总记录数
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s GetAgentStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgentStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentStatisticsResponseBodyData) SetPageNum(v int32) *GetAgentStatisticsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetAgentStatisticsResponseBodyData) SetPageSize(v int32) *GetAgentStatisticsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAgentStatisticsResponseBodyData) SetTotalNum(v int32) *GetAgentStatisticsResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetAgentStatisticsResponseBodyData) SetRows(v string) *GetAgentStatisticsResponseBodyData {
	s.Rows = &v
	return s
}

type GetAgentStatisticsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAgentStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetAgentStatisticsResponse) SetHeaders(v map[string]*string) *GetAgentStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetAgentStatisticsResponse) SetBody(v *GetAgentStatisticsResponseBody) *GetAgentStatisticsResponse {
	s.Body = v
	return s
}

type GetAllDepartmentRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAllDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllDepartmentRequest) GoString() string {
	return s.String()
}

func (s *GetAllDepartmentRequest) SetInstanceId(v string) *GetAllDepartmentRequest {
	s.InstanceId = &v
	return s
}

type GetAllDepartmentResponseBody struct {
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// data
	Data []*GetAllDepartmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAllDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllDepartmentResponseBody) SetMessage(v string) *GetAllDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetAllDepartmentResponseBody) SetRequestId(v string) *GetAllDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllDepartmentResponseBody) SetHttpStatusCode(v int32) *GetAllDepartmentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAllDepartmentResponseBody) SetData(v []*GetAllDepartmentResponseBodyData) *GetAllDepartmentResponseBody {
	s.Data = v
	return s
}

func (s *GetAllDepartmentResponseBody) SetCode(v string) *GetAllDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAllDepartmentResponseBody) SetSuccess(v bool) *GetAllDepartmentResponseBody {
	s.Success = &v
	return s
}

type GetAllDepartmentResponseBodyData struct {
	DepartmentId   *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
}

func (s GetAllDepartmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAllDepartmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAllDepartmentResponseBodyData) SetDepartmentId(v int64) *GetAllDepartmentResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetAllDepartmentResponseBodyData) SetDepartmentName(v string) *GetAllDepartmentResponseBodyData {
	s.DepartmentName = &v
	return s
}

type GetAllDepartmentResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAllDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllDepartmentResponse) GoString() string {
	return s.String()
}

func (s *GetAllDepartmentResponse) SetHeaders(v map[string]*string) *GetAllDepartmentResponse {
	s.Headers = v
	return s
}

func (s *GetAllDepartmentResponse) SetBody(v *GetAllDepartmentResponseBody) *GetAllDepartmentResponse {
	s.Body = v
	return s
}

type GetConfigNumListRequest struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 部门ID
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// 账号名称
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetConfigNumListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConfigNumListRequest) GoString() string {
	return s.String()
}

func (s *GetConfigNumListRequest) SetInstanceId(v string) *GetConfigNumListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConfigNumListRequest) SetDepartmentId(v int64) *GetConfigNumListRequest {
	s.DepartmentId = &v
	return s
}

func (s *GetConfigNumListRequest) SetAccountName(v string) *GetConfigNumListRequest {
	s.AccountName = &v
	return s
}

type GetConfigNumListResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 接口调用是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 号码列表
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetConfigNumListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConfigNumListResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigNumListResponseBody) SetRequestId(v string) *GetConfigNumListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigNumListResponseBody) SetSuccess(v bool) *GetConfigNumListResponseBody {
	s.Success = &v
	return s
}

func (s *GetConfigNumListResponseBody) SetCode(v string) *GetConfigNumListResponseBody {
	s.Code = &v
	return s
}

func (s *GetConfigNumListResponseBody) SetMessage(v string) *GetConfigNumListResponseBody {
	s.Message = &v
	return s
}

func (s *GetConfigNumListResponseBody) SetData(v []*string) *GetConfigNumListResponseBody {
	s.Data = v
	return s
}

type GetConfigNumListResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConfigNumListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConfigNumListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConfigNumListResponse) GoString() string {
	return s.String()
}

func (s *GetConfigNumListResponse) SetHeaders(v map[string]*string) *GetConfigNumListResponse {
	s.Headers = v
	return s
}

func (s *GetConfigNumListResponse) SetBody(v *GetConfigNumListResponseBody) *GetConfigNumListResponse {
	s.Body = v
	return s
}

type GetCustomerInfoRequest struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 会员ID
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s GetCustomerInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerInfoRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerInfoRequest) SetInstanceId(v string) *GetCustomerInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCustomerInfoRequest) SetMemberId(v int64) *GetCustomerInfoRequest {
	s.MemberId = &v
	return s
}

type GetCustomerInfoResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 会员信息
	Data *GetCustomerInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetCustomerInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerInfoResponseBody) SetRequestId(v string) *GetCustomerInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerInfoResponseBody) SetMessage(v string) *GetCustomerInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerInfoResponseBody) SetCode(v string) *GetCustomerInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerInfoResponseBody) SetSuccess(v bool) *GetCustomerInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomerInfoResponseBody) SetData(v *GetCustomerInfoResponseBodyData) *GetCustomerInfoResponseBody {
	s.Data = v
	return s
}

type GetCustomerInfoResponseBodyData struct {
	// 昵称
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	// 头像
	Photo *string `json:"Photo,omitempty" xml:"Photo,omitempty"`
	// 会员ID
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 真实姓名
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	// 外部ID
	OuterId *string `json:"OuterId,omitempty" xml:"OuterId,omitempty"`
	// 自定义字段
	CustomizeFields map[string]interface{} `json:"CustomizeFields,omitempty" xml:"CustomizeFields,omitempty"`
}

func (s GetCustomerInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomerInfoResponseBodyData) SetNick(v string) *GetCustomerInfoResponseBodyData {
	s.Nick = &v
	return s
}

func (s *GetCustomerInfoResponseBodyData) SetPhoto(v string) *GetCustomerInfoResponseBodyData {
	s.Photo = &v
	return s
}

func (s *GetCustomerInfoResponseBodyData) SetUserId(v int64) *GetCustomerInfoResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetCustomerInfoResponseBodyData) SetRealName(v string) *GetCustomerInfoResponseBodyData {
	s.RealName = &v
	return s
}

func (s *GetCustomerInfoResponseBodyData) SetOuterId(v string) *GetCustomerInfoResponseBodyData {
	s.OuterId = &v
	return s
}

func (s *GetCustomerInfoResponseBodyData) SetCustomizeFields(v map[string]interface{}) *GetCustomerInfoResponseBodyData {
	s.CustomizeFields = v
	return s
}

type GetCustomerInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCustomerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCustomerInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerInfoResponse) SetHeaders(v map[string]*string) *GetCustomerInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerInfoResponse) SetBody(v *GetCustomerInfoResponseBody) *GetCustomerInfoResponse {
	s.Body = v
	return s
}

type GetDepartmentalLatitudeAgentStatusRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 技能组分组id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// 是否根据技能组分组id分组显示
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
}

func (s GetDepartmentalLatitudeAgentStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDepartmentalLatitudeAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) SetInstanceId(v string) *GetDepartmentalLatitudeAgentStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) SetStartDate(v int64) *GetDepartmentalLatitudeAgentStatusRequest {
	s.StartDate = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) SetEndDate(v int64) *GetDepartmentalLatitudeAgentStatusRequest {
	s.EndDate = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) SetPageSize(v int64) *GetDepartmentalLatitudeAgentStatusRequest {
	s.PageSize = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) SetCurrentPage(v int64) *GetDepartmentalLatitudeAgentStatusRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) SetDepIds(v []*int64) *GetDepartmentalLatitudeAgentStatusRequest {
	s.DepIds = v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) SetExistDepartmentGrouping(v bool) *GetDepartmentalLatitudeAgentStatusRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

type GetDepartmentalLatitudeAgentStatusShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 技能组分组id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// 是否根据技能组分组id分组显示
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
}

func (s GetDepartmentalLatitudeAgentStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDepartmentalLatitudeAgentStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) SetInstanceId(v string) *GetDepartmentalLatitudeAgentStatusShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) SetStartDate(v int64) *GetDepartmentalLatitudeAgentStatusShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) SetEndDate(v int64) *GetDepartmentalLatitudeAgentStatusShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) SetPageSize(v int64) *GetDepartmentalLatitudeAgentStatusShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) SetCurrentPage(v int64) *GetDepartmentalLatitudeAgentStatusShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) SetDepIdsShrink(v string) *GetDepartmentalLatitudeAgentStatusShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) SetExistDepartmentGrouping(v bool) *GetDepartmentalLatitudeAgentStatusShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

type GetDepartmentalLatitudeAgentStatusResponseBody struct {
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 数据
	Data *GetDepartmentalLatitudeAgentStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 接口调用是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDepartmentalLatitudeAgentStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDepartmentalLatitudeAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) SetMessage(v string) *GetDepartmentalLatitudeAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) SetRequestId(v string) *GetDepartmentalLatitudeAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) SetData(v *GetDepartmentalLatitudeAgentStatusResponseBodyData) *GetDepartmentalLatitudeAgentStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) SetCode(v string) *GetDepartmentalLatitudeAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) SetSuccess(v bool) *GetDepartmentalLatitudeAgentStatusResponseBody {
	s.Success = &v
	return s
}

type GetDepartmentalLatitudeAgentStatusResponseBodyData struct {
	// 当前页数
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 每页的数量
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总共记录数
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s GetDepartmentalLatitudeAgentStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDepartmentalLatitudeAgentStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBodyData) SetPageNum(v int32) *GetDepartmentalLatitudeAgentStatusResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBodyData) SetPageSize(v int32) *GetDepartmentalLatitudeAgentStatusResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBodyData) SetTotalNum(v int32) *GetDepartmentalLatitudeAgentStatusResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBodyData) SetRows(v string) *GetDepartmentalLatitudeAgentStatusResponseBodyData {
	s.Rows = &v
	return s
}

type GetDepartmentalLatitudeAgentStatusResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDepartmentalLatitudeAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDepartmentalLatitudeAgentStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDepartmentalLatitudeAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDepartmentalLatitudeAgentStatusResponse) SetHeaders(v map[string]*string) *GetDepartmentalLatitudeAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponse) SetBody(v *GetDepartmentalLatitudeAgentStatusResponseBody) *GetDepartmentalLatitudeAgentStatusResponse {
	s.Body = v
	return s
}

type GetDepGroupTreeDataRequest struct {
	// 租户实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 坐席ID
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
}

func (s GetDepGroupTreeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDepGroupTreeDataRequest) GoString() string {
	return s.String()
}

func (s *GetDepGroupTreeDataRequest) SetInstanceId(v string) *GetDepGroupTreeDataRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDepGroupTreeDataRequest) SetAgentId(v int64) *GetDepGroupTreeDataRequest {
	s.AgentId = &v
	return s
}

type GetDepGroupTreeDataResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Success
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// Data
	Data []*GetDepGroupTreeDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetDepGroupTreeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDepGroupTreeDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetDepGroupTreeDataResponseBody) SetRequestId(v string) *GetDepGroupTreeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBody) SetMessage(v string) *GetDepGroupTreeDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBody) SetCode(v string) *GetDepGroupTreeDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBody) SetSuccess(v string) *GetDepGroupTreeDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBody) SetData(v []*GetDepGroupTreeDataResponseBodyData) *GetDepGroupTreeDataResponseBody {
	s.Data = v
	return s
}

type GetDepGroupTreeDataResponseBodyData struct {
	// depGroupName
	DepGroupName *string `json:"DepGroupName,omitempty" xml:"DepGroupName,omitempty"`
	// depGroupId
	DepGroupId *string `json:"DepGroupId,omitempty" xml:"DepGroupId,omitempty"`
	// groupDTOS
	GroupDTOS []*GetDepGroupTreeDataResponseBodyDataGroupDTOS `json:"GroupDTOS,omitempty" xml:"GroupDTOS,omitempty" type:"Repeated"`
}

func (s GetDepGroupTreeDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDepGroupTreeDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDepGroupTreeDataResponseBodyData) SetDepGroupName(v string) *GetDepGroupTreeDataResponseBodyData {
	s.DepGroupName = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBodyData) SetDepGroupId(v string) *GetDepGroupTreeDataResponseBodyData {
	s.DepGroupId = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBodyData) SetGroupDTOS(v []*GetDepGroupTreeDataResponseBodyDataGroupDTOS) *GetDepGroupTreeDataResponseBodyData {
	s.GroupDTOS = v
	return s
}

type GetDepGroupTreeDataResponseBodyDataGroupDTOS struct {
	// skillGroupId
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetDepGroupTreeDataResponseBodyDataGroupDTOS) String() string {
	return tea.Prettify(s)
}

func (s GetDepGroupTreeDataResponseBodyDataGroupDTOS) GoString() string {
	return s.String()
}

func (s *GetDepGroupTreeDataResponseBodyDataGroupDTOS) SetSkillGroupId(v int64) *GetDepGroupTreeDataResponseBodyDataGroupDTOS {
	s.SkillGroupId = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBodyDataGroupDTOS) SetName(v string) *GetDepGroupTreeDataResponseBodyDataGroupDTOS {
	s.Name = &v
	return s
}

type GetDepGroupTreeDataResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDepGroupTreeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDepGroupTreeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDepGroupTreeDataResponse) GoString() string {
	return s.String()
}

func (s *GetDepGroupTreeDataResponse) SetHeaders(v map[string]*string) *GetDepGroupTreeDataResponse {
	s.Headers = v
	return s
}

func (s *GetDepGroupTreeDataResponse) SetBody(v *GetDepGroupTreeDataResponseBody) *GetDepGroupTreeDataResponse {
	s.Body = v
	return s
}

type GetHotlineAgentDetailRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetHotlineAgentDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailRequest) SetClientToken(v string) *GetHotlineAgentDetailRequest {
	s.ClientToken = &v
	return s
}

func (s *GetHotlineAgentDetailRequest) SetInstanceId(v string) *GetHotlineAgentDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineAgentDetailRequest) SetAccountName(v string) *GetHotlineAgentDetailRequest {
	s.AccountName = &v
	return s
}

type GetHotlineAgentDetailResponseBody struct {
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetHotlineAgentDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GetHotlineAgentDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailResponseBody) SetMessage(v string) *GetHotlineAgentDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetRequestId(v string) *GetHotlineAgentDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetData(v *GetHotlineAgentDetailResponseBodyData) *GetHotlineAgentDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetCode(v string) *GetHotlineAgentDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetSuccess(v bool) *GetHotlineAgentDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetHttpStatusCode(v int64) *GetHotlineAgentDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GetHotlineAgentDetailResponseBodyData struct {
	AgentStatusCode *string `json:"AgentStatusCode,omitempty" xml:"AgentStatusCode,omitempty"`
	Token           *string `json:"Token,omitempty" xml:"Token,omitempty"`
	AgentId         *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	Assigned        *bool   `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	RestType        *int32  `json:"RestType,omitempty" xml:"RestType,omitempty"`
	AgentStatus     *int32  `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	TenantId        *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetHotlineAgentDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAgentStatusCode(v string) *GetHotlineAgentDetailResponseBodyData {
	s.AgentStatusCode = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetToken(v string) *GetHotlineAgentDetailResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAgentId(v int64) *GetHotlineAgentDetailResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAssigned(v bool) *GetHotlineAgentDetailResponseBodyData {
	s.Assigned = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetRestType(v int32) *GetHotlineAgentDetailResponseBodyData {
	s.RestType = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAgentStatus(v int32) *GetHotlineAgentDetailResponseBodyData {
	s.AgentStatus = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetTenantId(v int64) *GetHotlineAgentDetailResponseBodyData {
	s.TenantId = &v
	return s
}

type GetHotlineAgentDetailResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineAgentDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineAgentDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailResponse) SetHeaders(v map[string]*string) *GetHotlineAgentDetailResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineAgentDetailResponse) SetBody(v *GetHotlineAgentDetailResponseBody) *GetHotlineAgentDetailResponse {
	s.Body = v
	return s
}

type GetHotlineAgentDetailReportRequest struct {
	CurrentPage *int32   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate   *int64   `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate     *int64   `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	InstanceId  *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DepIds      []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	GroupIds    []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
}

func (s GetHotlineAgentDetailReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportRequest) SetCurrentPage(v int32) *GetHotlineAgentDetailReportRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetPageSize(v int32) *GetHotlineAgentDetailReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetStartDate(v int64) *GetHotlineAgentDetailReportRequest {
	s.StartDate = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetEndDate(v int64) *GetHotlineAgentDetailReportRequest {
	s.EndDate = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetInstanceId(v string) *GetHotlineAgentDetailReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetDepIds(v []*int64) *GetHotlineAgentDetailReportRequest {
	s.DepIds = v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetGroupIds(v []*int64) *GetHotlineAgentDetailReportRequest {
	s.GroupIds = v
	return s
}

type GetHotlineAgentDetailReportResponseBody struct {
	Message        *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetHotlineAgentDetailReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code           *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *string                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GetHotlineAgentDetailReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponseBody) SetMessage(v string) *GetHotlineAgentDetailReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetRequestId(v string) *GetHotlineAgentDetailReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetData(v *GetHotlineAgentDetailReportResponseBodyData) *GetHotlineAgentDetailReportResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetCode(v string) *GetHotlineAgentDetailReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetSuccess(v string) *GetHotlineAgentDetailReportResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetHttpStatusCode(v int64) *GetHotlineAgentDetailReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GetHotlineAgentDetailReportResponseBodyData struct {
	Rows     []map[string]interface{}                              `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	PageSize *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                                `json:"Total,omitempty" xml:"Total,omitempty"`
	Columns  []*GetHotlineAgentDetailReportResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Page     *int32                                                `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s GetHotlineAgentDetailReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetRows(v []map[string]interface{}) *GetHotlineAgentDetailReportResponseBodyData {
	s.Rows = v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetPageSize(v int32) *GetHotlineAgentDetailReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetTotal(v int32) *GetHotlineAgentDetailReportResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetColumns(v []*GetHotlineAgentDetailReportResponseBodyDataColumns) *GetHotlineAgentDetailReportResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetPage(v int32) *GetHotlineAgentDetailReportResponseBodyData {
	s.Page = &v
	return s
}

type GetHotlineAgentDetailReportResponseBodyDataColumns struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetHotlineAgentDetailReportResponseBodyDataColumns) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponseBodyDataColumns) SetKey(v string) *GetHotlineAgentDetailReportResponseBodyDataColumns {
	s.Key = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyDataColumns) SetTitle(v string) *GetHotlineAgentDetailReportResponseBodyDataColumns {
	s.Title = &v
	return s
}

type GetHotlineAgentDetailReportResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineAgentDetailReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineAgentDetailReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponse) SetHeaders(v map[string]*string) *GetHotlineAgentDetailReportResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineAgentDetailReportResponse) SetBody(v *GetHotlineAgentDetailReportResponseBody) *GetHotlineAgentDetailReportResponse {
	s.Body = v
	return s
}

type GetHotlineAgentStatusRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetHotlineAgentStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentStatusRequest) SetInstanceId(v string) *GetHotlineAgentStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineAgentStatusRequest) SetAccountName(v string) *GetHotlineAgentStatusRequest {
	s.AccountName = &v
	return s
}

type GetHotlineAgentStatusResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GetHotlineAgentStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentStatusResponseBody) SetMessage(v string) *GetHotlineAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetRequestId(v string) *GetHotlineAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetData(v string) *GetHotlineAgentStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetCode(v string) *GetHotlineAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetSuccess(v bool) *GetHotlineAgentStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetHttpStatusCode(v int64) *GetHotlineAgentStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GetHotlineAgentStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineAgentStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentStatusResponse) SetHeaders(v map[string]*string) *GetHotlineAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineAgentStatusResponse) SetBody(v *GetHotlineAgentStatusResponseBody) *GetHotlineAgentStatusResponse {
	s.Body = v
	return s
}

type GetHotlineCallActionRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Act         *int32  `json:"Act,omitempty" xml:"Act,omitempty"`
	FromSource  *string `json:"FromSource,omitempty" xml:"FromSource,omitempty"`
	Biz         *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	Acc         *string `json:"Acc,omitempty" xml:"Acc,omitempty"`
}

func (s GetHotlineCallActionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineCallActionRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineCallActionRequest) SetClientToken(v string) *GetHotlineCallActionRequest {
	s.ClientToken = &v
	return s
}

func (s *GetHotlineCallActionRequest) SetInstanceId(v string) *GetHotlineCallActionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineCallActionRequest) SetAccountName(v string) *GetHotlineCallActionRequest {
	s.AccountName = &v
	return s
}

func (s *GetHotlineCallActionRequest) SetAct(v int32) *GetHotlineCallActionRequest {
	s.Act = &v
	return s
}

func (s *GetHotlineCallActionRequest) SetFromSource(v string) *GetHotlineCallActionRequest {
	s.FromSource = &v
	return s
}

func (s *GetHotlineCallActionRequest) SetBiz(v string) *GetHotlineCallActionRequest {
	s.Biz = &v
	return s
}

func (s *GetHotlineCallActionRequest) SetAcc(v string) *GetHotlineCallActionRequest {
	s.Acc = &v
	return s
}

type GetHotlineCallActionResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *GetHotlineCallActionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetHotlineCallActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineCallActionResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineCallActionResponseBody) SetCode(v string) *GetHotlineCallActionResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineCallActionResponseBody) SetMessage(v string) *GetHotlineCallActionResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineCallActionResponseBody) SetRequestId(v string) *GetHotlineCallActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineCallActionResponseBody) SetSuccess(v bool) *GetHotlineCallActionResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineCallActionResponseBody) SetData(v *GetHotlineCallActionResponseBodyData) *GetHotlineCallActionResponseBody {
	s.Data = v
	return s
}

type GetHotlineCallActionResponseBodyData struct {
	TouchId      *int64  `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
	DepId        *int64  `json:"DepId,omitempty" xml:"DepId,omitempty"`
	ServicerName *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	MemberName   *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	ChannelType  *int64  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	SubTouchId   *int64  `json:"SubTouchId,omitempty" xml:"SubTouchId,omitempty"`
	CalloutName  *string `json:"CalloutName,omitempty" xml:"CalloutName,omitempty"`
	ActionId     *int64  `json:"ActionId,omitempty" xml:"ActionId,omitempty"`
	ServicerId   *int64  `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	BuId         *int64  `json:"BuId,omitempty" xml:"BuId,omitempty"`
	CalloutId    *int64  `json:"CalloutId,omitempty" xml:"CalloutId,omitempty"`
	CaseId       *int64  `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	ChannelId    *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	IsTransfer   *string `json:"IsTransfer,omitempty" xml:"IsTransfer,omitempty"`
	MemberId     *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	MemberList   *string `json:"MemberList,omitempty" xml:"MemberList,omitempty"`
}

func (s GetHotlineCallActionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineCallActionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineCallActionResponseBodyData) SetTouchId(v int64) *GetHotlineCallActionResponseBodyData {
	s.TouchId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetDepId(v int64) *GetHotlineCallActionResponseBodyData {
	s.DepId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetServicerName(v string) *GetHotlineCallActionResponseBodyData {
	s.ServicerName = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetMemberName(v string) *GetHotlineCallActionResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetChannelType(v int64) *GetHotlineCallActionResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetSubTouchId(v int64) *GetHotlineCallActionResponseBodyData {
	s.SubTouchId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetCalloutName(v string) *GetHotlineCallActionResponseBodyData {
	s.CalloutName = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetActionId(v int64) *GetHotlineCallActionResponseBodyData {
	s.ActionId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetServicerId(v int64) *GetHotlineCallActionResponseBodyData {
	s.ServicerId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetBuId(v int64) *GetHotlineCallActionResponseBodyData {
	s.BuId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetCalloutId(v int64) *GetHotlineCallActionResponseBodyData {
	s.CalloutId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetCaseId(v int64) *GetHotlineCallActionResponseBodyData {
	s.CaseId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetChannelId(v string) *GetHotlineCallActionResponseBodyData {
	s.ChannelId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetIsTransfer(v string) *GetHotlineCallActionResponseBodyData {
	s.IsTransfer = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetMemberId(v int64) *GetHotlineCallActionResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetTaskId(v int64) *GetHotlineCallActionResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetMemberList(v string) *GetHotlineCallActionResponseBodyData {
	s.MemberList = &v
	return s
}

type GetHotlineCallActionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineCallActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineCallActionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineCallActionResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineCallActionResponse) SetHeaders(v map[string]*string) *GetHotlineCallActionResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineCallActionResponse) SetBody(v *GetHotlineCallActionResponseBody) *GetHotlineCallActionResponse {
	s.Body = v
	return s
}

type GetHotlineGroupDetailReportRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate   *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate     *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DepIds      []*int  `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	GroupIds    []*int  `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
}

func (s GetHotlineGroupDetailReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportRequest) SetCurrentPage(v int32) *GetHotlineGroupDetailReportRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetPageSize(v int32) *GetHotlineGroupDetailReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetStartDate(v int64) *GetHotlineGroupDetailReportRequest {
	s.StartDate = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetEndDate(v int64) *GetHotlineGroupDetailReportRequest {
	s.EndDate = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetInstanceId(v string) *GetHotlineGroupDetailReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetDepIds(v []*int) *GetHotlineGroupDetailReportRequest {
	s.DepIds = v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetGroupIds(v []*int) *GetHotlineGroupDetailReportRequest {
	s.GroupIds = v
	return s
}

type GetHotlineGroupDetailReportResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *GetHotlineGroupDetailReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetHotlineGroupDetailReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponseBody) SetCode(v string) *GetHotlineGroupDetailReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetMessage(v string) *GetHotlineGroupDetailReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetRequestId(v string) *GetHotlineGroupDetailReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetSuccess(v string) *GetHotlineGroupDetailReportResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetData(v *GetHotlineGroupDetailReportResponseBodyData) *GetHotlineGroupDetailReportResponseBody {
	s.Data = v
	return s
}

type GetHotlineGroupDetailReportResponseBodyData struct {
	PageSize *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                                `json:"Total,omitempty" xml:"Total,omitempty"`
	Page     *int32                                                `json:"Page,omitempty" xml:"Page,omitempty"`
	Columns  []*GetHotlineGroupDetailReportResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Rows     []map[string]interface{}                              `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s GetHotlineGroupDetailReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetPageSize(v int32) *GetHotlineGroupDetailReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetTotal(v int32) *GetHotlineGroupDetailReportResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetPage(v int32) *GetHotlineGroupDetailReportResponseBodyData {
	s.Page = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetColumns(v []*GetHotlineGroupDetailReportResponseBodyDataColumns) *GetHotlineGroupDetailReportResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetRows(v []map[string]interface{}) *GetHotlineGroupDetailReportResponseBodyData {
	s.Rows = v
	return s
}

type GetHotlineGroupDetailReportResponseBodyDataColumns struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetHotlineGroupDetailReportResponseBodyDataColumns) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponseBodyDataColumns) SetKey(v string) *GetHotlineGroupDetailReportResponseBodyDataColumns {
	s.Key = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyDataColumns) SetTitle(v string) *GetHotlineGroupDetailReportResponseBodyDataColumns {
	s.Title = &v
	return s
}

type GetHotlineGroupDetailReportResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineGroupDetailReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineGroupDetailReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponse) SetHeaders(v map[string]*string) *GetHotlineGroupDetailReportResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineGroupDetailReportResponse) SetBody(v *GetHotlineGroupDetailReportResponseBody) *GetHotlineGroupDetailReportResponse {
	s.Body = v
	return s
}

type GetHotlineMessageLogRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 会话id
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
}

func (s GetHotlineMessageLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineMessageLogRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineMessageLogRequest) SetInstanceId(v string) *GetHotlineMessageLogRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineMessageLogRequest) SetAcid(v string) *GetHotlineMessageLogRequest {
	s.Acid = &v
	return s
}

type GetHotlineMessageLogResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	Data    []*GetHotlineMessageLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetHotlineMessageLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineMessageLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineMessageLogResponseBody) SetRequestId(v string) *GetHotlineMessageLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineMessageLogResponseBody) SetSuccess(v bool) *GetHotlineMessageLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineMessageLogResponseBody) SetCode(v string) *GetHotlineMessageLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineMessageLogResponseBody) SetMessage(v string) *GetHotlineMessageLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineMessageLogResponseBody) SetData(v []*GetHotlineMessageLogResponseBodyData) *GetHotlineMessageLogResponseBody {
	s.Data = v
	return s
}

type GetHotlineMessageLogResponseBodyData struct {
	// 会话ID
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	// 发送方类型（1：会员，2：坐席）
	SenderType *int32 `json:"SenderType,omitempty" xml:"SenderType,omitempty"`
	// 开始时间
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 记录id
	Mid *string `json:"Mid,omitempty" xml:"Mid,omitempty"`
	// 会话内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s GetHotlineMessageLogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineMessageLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineMessageLogResponseBodyData) SetAcid(v string) *GetHotlineMessageLogResponseBodyData {
	s.Acid = &v
	return s
}

func (s *GetHotlineMessageLogResponseBodyData) SetSenderType(v int32) *GetHotlineMessageLogResponseBodyData {
	s.SenderType = &v
	return s
}

func (s *GetHotlineMessageLogResponseBodyData) SetStartTime(v int64) *GetHotlineMessageLogResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetHotlineMessageLogResponseBodyData) SetEndTime(v int64) *GetHotlineMessageLogResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetHotlineMessageLogResponseBodyData) SetMid(v string) *GetHotlineMessageLogResponseBodyData {
	s.Mid = &v
	return s
}

func (s *GetHotlineMessageLogResponseBodyData) SetContent(v string) *GetHotlineMessageLogResponseBodyData {
	s.Content = &v
	return s
}

type GetHotlineMessageLogResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineMessageLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineMessageLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineMessageLogResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineMessageLogResponse) SetHeaders(v map[string]*string) *GetHotlineMessageLogResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineMessageLogResponse) SetBody(v *GetHotlineMessageLogResponseBody) *GetHotlineMessageLogResponse {
	s.Body = v
	return s
}

type GetHotlineRuntimeInfoRequest struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 账号名
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetHotlineRuntimeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineRuntimeInfoRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineRuntimeInfoRequest) SetInstanceId(v string) *GetHotlineRuntimeInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineRuntimeInfoRequest) SetAccountName(v string) *GetHotlineRuntimeInfoRequest {
	s.AccountName = &v
	return s
}

type GetHotlineRuntimeInfoResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 接口调用是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 数据结果
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetHotlineRuntimeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineRuntimeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineRuntimeInfoResponseBody) SetRequestId(v string) *GetHotlineRuntimeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineRuntimeInfoResponseBody) SetSuccess(v bool) *GetHotlineRuntimeInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineRuntimeInfoResponseBody) SetCode(v string) *GetHotlineRuntimeInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineRuntimeInfoResponseBody) SetMessage(v string) *GetHotlineRuntimeInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineRuntimeInfoResponseBody) SetData(v map[string]interface{}) *GetHotlineRuntimeInfoResponseBody {
	s.Data = v
	return s
}

type GetHotlineRuntimeInfoResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineRuntimeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineRuntimeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineRuntimeInfoResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineRuntimeInfoResponse) SetHeaders(v map[string]*string) *GetHotlineRuntimeInfoResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineRuntimeInfoResponse) SetBody(v *GetHotlineRuntimeInfoResponseBody) *GetHotlineRuntimeInfoResponse {
	s.Body = v
	return s
}

type GetHotlineServiceStatisticsRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// 部门id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// 技能组id列表
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// 时间纬度类型
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
	// 是否根据坐席分组
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetHotlineServiceStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineServiceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineServiceStatisticsRequest) SetInstanceId(v string) *GetHotlineServiceStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetStartDate(v int64) *GetHotlineServiceStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetEndDate(v int64) *GetHotlineServiceStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetPageSize(v int32) *GetHotlineServiceStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetCurrentPage(v int32) *GetHotlineServiceStatisticsRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetAgentIds(v []*int64) *GetHotlineServiceStatisticsRequest {
	s.AgentIds = v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetDepIds(v []*int64) *GetHotlineServiceStatisticsRequest {
	s.DepIds = v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetGroupIds(v []*int64) *GetHotlineServiceStatisticsRequest {
	s.GroupIds = v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetTimeLatitudeType(v string) *GetHotlineServiceStatisticsRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetExistAgentGrouping(v bool) *GetHotlineServiceStatisticsRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetExistDepartmentGrouping(v bool) *GetHotlineServiceStatisticsRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetExistSkillGroupGrouping(v bool) *GetHotlineServiceStatisticsRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetHotlineServiceStatisticsShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// 技能组id列表
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// 时间纬度类型
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
	// 是否根据坐席分组
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetHotlineServiceStatisticsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineServiceStatisticsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetInstanceId(v string) *GetHotlineServiceStatisticsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetStartDate(v int64) *GetHotlineServiceStatisticsShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetEndDate(v int64) *GetHotlineServiceStatisticsShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetPageSize(v int32) *GetHotlineServiceStatisticsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetCurrentPage(v int32) *GetHotlineServiceStatisticsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetAgentIdsShrink(v string) *GetHotlineServiceStatisticsShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetDepIdsShrink(v string) *GetHotlineServiceStatisticsShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetGroupIdsShrink(v string) *GetHotlineServiceStatisticsShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetTimeLatitudeType(v string) *GetHotlineServiceStatisticsShrinkRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetExistAgentGrouping(v bool) *GetHotlineServiceStatisticsShrinkRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetExistDepartmentGrouping(v bool) *GetHotlineServiceStatisticsShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetHotlineServiceStatisticsShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetHotlineServiceStatisticsShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetHotlineServiceStatisticsResponseBody struct {
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 调用接口是否成功
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// data
	Data *GetHotlineServiceStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetHotlineServiceStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineServiceStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineServiceStatisticsResponseBody) SetRequestId(v string) *GetHotlineServiceStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBody) SetMessage(v string) *GetHotlineServiceStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBody) SetCode(v string) *GetHotlineServiceStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBody) SetSuccess(v string) *GetHotlineServiceStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBody) SetData(v *GetHotlineServiceStatisticsResponseBodyData) *GetHotlineServiceStatisticsResponseBody {
	s.Data = v
	return s
}

type GetHotlineServiceStatisticsResponseBodyData struct {
	// 当前页数
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总记录数
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s GetHotlineServiceStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineServiceStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineServiceStatisticsResponseBodyData) SetPageNum(v int32) *GetHotlineServiceStatisticsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBodyData) SetPageSize(v int32) *GetHotlineServiceStatisticsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBodyData) SetTotalNum(v int32) *GetHotlineServiceStatisticsResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBodyData) SetRows(v string) *GetHotlineServiceStatisticsResponseBodyData {
	s.Rows = &v
	return s
}

type GetHotlineServiceStatisticsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineServiceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineServiceStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineServiceStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineServiceStatisticsResponse) SetHeaders(v map[string]*string) *GetHotlineServiceStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineServiceStatisticsResponse) SetBody(v *GetHotlineServiceStatisticsResponseBody) *GetHotlineServiceStatisticsResponse {
	s.Body = v
	return s
}

type GetHotlineWaitingNumberRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetHotlineWaitingNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineWaitingNumberRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineWaitingNumberRequest) SetClientToken(v string) *GetHotlineWaitingNumberRequest {
	s.ClientToken = &v
	return s
}

func (s *GetHotlineWaitingNumberRequest) SetInstanceId(v string) *GetHotlineWaitingNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineWaitingNumberRequest) SetAccountName(v string) *GetHotlineWaitingNumberRequest {
	s.AccountName = &v
	return s
}

type GetHotlineWaitingNumberResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineWaitingNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineWaitingNumberResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineWaitingNumberResponseBody) SetCode(v string) *GetHotlineWaitingNumberResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetMessage(v string) *GetHotlineWaitingNumberResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetData(v int64) *GetHotlineWaitingNumberResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetRequestId(v string) *GetHotlineWaitingNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetSuccess(v string) *GetHotlineWaitingNumberResponseBody {
	s.Success = &v
	return s
}

type GetHotlineWaitingNumberResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineWaitingNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineWaitingNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineWaitingNumberResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineWaitingNumberResponse) SetHeaders(v map[string]*string) *GetHotlineWaitingNumberResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineWaitingNumberResponse) SetBody(v *GetHotlineWaitingNumberResponseBody) *GetHotlineWaitingNumberResponse {
	s.Body = v
	return s
}

type GetIndexCurrentValueRequest struct {
	DepIds     *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	GroupIds   *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetIndexCurrentValueRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIndexCurrentValueRequest) GoString() string {
	return s.String()
}

func (s *GetIndexCurrentValueRequest) SetDepIds(v string) *GetIndexCurrentValueRequest {
	s.DepIds = &v
	return s
}

func (s *GetIndexCurrentValueRequest) SetGroupIds(v string) *GetIndexCurrentValueRequest {
	s.GroupIds = &v
	return s
}

func (s *GetIndexCurrentValueRequest) SetInstanceId(v string) *GetIndexCurrentValueRequest {
	s.InstanceId = &v
	return s
}

type GetIndexCurrentValueResponseBody struct {
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetIndexCurrentValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIndexCurrentValueResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexCurrentValueResponseBody) SetMessage(v string) *GetIndexCurrentValueResponseBody {
	s.Message = &v
	return s
}

func (s *GetIndexCurrentValueResponseBody) SetRequestId(v string) *GetIndexCurrentValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexCurrentValueResponseBody) SetData(v []map[string]interface{}) *GetIndexCurrentValueResponseBody {
	s.Data = v
	return s
}

func (s *GetIndexCurrentValueResponseBody) SetCode(v string) *GetIndexCurrentValueResponseBody {
	s.Code = &v
	return s
}

func (s *GetIndexCurrentValueResponseBody) SetSuccess(v bool) *GetIndexCurrentValueResponseBody {
	s.Success = &v
	return s
}

type GetIndexCurrentValueResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIndexCurrentValueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIndexCurrentValueResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIndexCurrentValueResponse) GoString() string {
	return s.String()
}

func (s *GetIndexCurrentValueResponse) SetHeaders(v map[string]*string) *GetIndexCurrentValueResponse {
	s.Headers = v
	return s
}

func (s *GetIndexCurrentValueResponse) SetBody(v *GetIndexCurrentValueResponseBody) *GetIndexCurrentValueResponse {
	s.Body = v
	return s
}

type GetInstanceListRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceListRequest) SetPageNumber(v int32) *GetInstanceListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetInstanceListRequest) SetPageSize(v int32) *GetInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *GetInstanceListRequest) SetName(v string) *GetInstanceListRequest {
	s.Name = &v
	return s
}

type GetInstanceListResponseBody struct {
	HttpStatusCode     *int32                                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId          *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Code               *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message            *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize           *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber         *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount         *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	CommodityInstances []*GetInstanceListResponseBodyCommodityInstances `json:"CommodityInstances,omitempty" xml:"CommodityInstances,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBody) SetHttpStatusCode(v int32) *GetInstanceListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceListResponseBody) SetRequestId(v string) *GetInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceListResponseBody) SetSuccess(v bool) *GetInstanceListResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceListResponseBody) SetCode(v string) *GetInstanceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceListResponseBody) SetMessage(v string) *GetInstanceListResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceListResponseBody) SetPageSize(v int32) *GetInstanceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetInstanceListResponseBody) SetPageNumber(v int32) *GetInstanceListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetInstanceListResponseBody) SetTotalCount(v int32) *GetInstanceListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetInstanceListResponseBody) SetCommodityInstances(v []*GetInstanceListResponseBodyCommodityInstances) *GetInstanceListResponseBody {
	s.CommodityInstances = v
	return s
}

type GetInstanceListResponseBodyCommodityInstances struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceListResponseBodyCommodityInstances) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBodyCommodityInstances) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyCommodityInstances) SetName(v string) *GetInstanceListResponseBodyCommodityInstances {
	s.Name = &v
	return s
}

func (s *GetInstanceListResponseBodyCommodityInstances) SetInstanceId(v string) *GetInstanceListResponseBodyCommodityInstances {
	s.InstanceId = &v
	return s
}

type GetInstanceListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponse) SetHeaders(v map[string]*string) *GetInstanceListResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceListResponse) SetBody(v *GetInstanceListResponseBody) *GetInstanceListResponse {
	s.Body = v
	return s
}

type GetMcuLvsIpRequest struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetMcuLvsIpRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMcuLvsIpRequest) GoString() string {
	return s.String()
}

func (s *GetMcuLvsIpRequest) SetInstanceId(v string) *GetMcuLvsIpRequest {
	s.InstanceId = &v
	return s
}

type GetMcuLvsIpResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// error code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// error message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// result data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetMcuLvsIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMcuLvsIpResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcuLvsIpResponseBody) SetRequestId(v string) *GetMcuLvsIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcuLvsIpResponseBody) SetSuccess(v bool) *GetMcuLvsIpResponseBody {
	s.Success = &v
	return s
}

func (s *GetMcuLvsIpResponseBody) SetCode(v string) *GetMcuLvsIpResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcuLvsIpResponseBody) SetMessage(v string) *GetMcuLvsIpResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcuLvsIpResponseBody) SetData(v string) *GetMcuLvsIpResponseBody {
	s.Data = &v
	return s
}

type GetMcuLvsIpResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMcuLvsIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMcuLvsIpResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMcuLvsIpResponse) GoString() string {
	return s.String()
}

func (s *GetMcuLvsIpResponse) SetHeaders(v map[string]*string) *GetMcuLvsIpResponse {
	s.Headers = v
	return s
}

func (s *GetMcuLvsIpResponse) SetBody(v *GetMcuLvsIpResponseBody) *GetMcuLvsIpResponse {
	s.Body = v
	return s
}

type GetNumLocationRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PhoneNum    *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
}

func (s GetNumLocationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNumLocationRequest) GoString() string {
	return s.String()
}

func (s *GetNumLocationRequest) SetClientToken(v string) *GetNumLocationRequest {
	s.ClientToken = &v
	return s
}

func (s *GetNumLocationRequest) SetInstanceId(v string) *GetNumLocationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNumLocationRequest) SetPhoneNum(v string) *GetNumLocationRequest {
	s.PhoneNum = &v
	return s
}

type GetNumLocationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNumLocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNumLocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetNumLocationResponseBody) SetCode(v string) *GetNumLocationResponseBody {
	s.Code = &v
	return s
}

func (s *GetNumLocationResponseBody) SetMessage(v string) *GetNumLocationResponseBody {
	s.Message = &v
	return s
}

func (s *GetNumLocationResponseBody) SetData(v string) *GetNumLocationResponseBody {
	s.Data = &v
	return s
}

func (s *GetNumLocationResponseBody) SetRequestId(v string) *GetNumLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNumLocationResponseBody) SetSuccess(v bool) *GetNumLocationResponseBody {
	s.Success = &v
	return s
}

type GetNumLocationResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetNumLocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNumLocationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNumLocationResponse) GoString() string {
	return s.String()
}

func (s *GetNumLocationResponse) SetHeaders(v map[string]*string) *GetNumLocationResponse {
	s.Headers = v
	return s
}

func (s *GetNumLocationResponse) SetBody(v *GetNumLocationResponseBody) *GetNumLocationResponse {
	s.Body = v
	return s
}

type GetOnlineSeatInformationRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// 部门id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
}

func (s GetOnlineSeatInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOnlineSeatInformationRequest) GoString() string {
	return s.String()
}

func (s *GetOnlineSeatInformationRequest) SetInstanceId(v string) *GetOnlineSeatInformationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOnlineSeatInformationRequest) SetStartDate(v int64) *GetOnlineSeatInformationRequest {
	s.StartDate = &v
	return s
}

func (s *GetOnlineSeatInformationRequest) SetEndDate(v int64) *GetOnlineSeatInformationRequest {
	s.EndDate = &v
	return s
}

func (s *GetOnlineSeatInformationRequest) SetPageSize(v int32) *GetOnlineSeatInformationRequest {
	s.PageSize = &v
	return s
}

func (s *GetOnlineSeatInformationRequest) SetCurrentPage(v int32) *GetOnlineSeatInformationRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetOnlineSeatInformationRequest) SetAgentIds(v []*int64) *GetOnlineSeatInformationRequest {
	s.AgentIds = v
	return s
}

func (s *GetOnlineSeatInformationRequest) SetDepIds(v []*int64) *GetOnlineSeatInformationRequest {
	s.DepIds = v
	return s
}

type GetOnlineSeatInformationShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
}

func (s GetOnlineSeatInformationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOnlineSeatInformationShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetOnlineSeatInformationShrinkRequest) SetInstanceId(v string) *GetOnlineSeatInformationShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOnlineSeatInformationShrinkRequest) SetStartDate(v int64) *GetOnlineSeatInformationShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetOnlineSeatInformationShrinkRequest) SetEndDate(v int64) *GetOnlineSeatInformationShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetOnlineSeatInformationShrinkRequest) SetPageSize(v int32) *GetOnlineSeatInformationShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetOnlineSeatInformationShrinkRequest) SetCurrentPage(v int32) *GetOnlineSeatInformationShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetOnlineSeatInformationShrinkRequest) SetAgentIdsShrink(v string) *GetOnlineSeatInformationShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetOnlineSeatInformationShrinkRequest) SetDepIdsShrink(v string) *GetOnlineSeatInformationShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

type GetOnlineSeatInformationResponseBody struct {
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 调用接口是否成功
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// data
	Data *GetOnlineSeatInformationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetOnlineSeatInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOnlineSeatInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetOnlineSeatInformationResponseBody) SetRequestId(v string) *GetOnlineSeatInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBody) SetMessage(v string) *GetOnlineSeatInformationResponseBody {
	s.Message = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBody) SetCode(v string) *GetOnlineSeatInformationResponseBody {
	s.Code = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBody) SetSuccess(v string) *GetOnlineSeatInformationResponseBody {
	s.Success = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBody) SetData(v *GetOnlineSeatInformationResponseBodyData) *GetOnlineSeatInformationResponseBody {
	s.Data = v
	return s
}

type GetOnlineSeatInformationResponseBodyData struct {
	// 当前页数
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总记录数
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s GetOnlineSeatInformationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOnlineSeatInformationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOnlineSeatInformationResponseBodyData) SetPageNum(v int32) *GetOnlineSeatInformationResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBodyData) SetPageSize(v int32) *GetOnlineSeatInformationResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBodyData) SetTotalNum(v int32) *GetOnlineSeatInformationResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBodyData) SetRows(v string) *GetOnlineSeatInformationResponseBodyData {
	s.Rows = &v
	return s
}

type GetOnlineSeatInformationResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOnlineSeatInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOnlineSeatInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOnlineSeatInformationResponse) GoString() string {
	return s.String()
}

func (s *GetOnlineSeatInformationResponse) SetHeaders(v map[string]*string) *GetOnlineSeatInformationResponse {
	s.Headers = v
	return s
}

func (s *GetOnlineSeatInformationResponse) SetBody(v *GetOnlineSeatInformationResponseBody) *GetOnlineSeatInformationResponse {
	s.Body = v
	return s
}

type GetOnlineServiceVolumeRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// 部门id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// 技能组id列表
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// 时间纬度类型
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
	// 是否根据坐席分组
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetOnlineServiceVolumeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOnlineServiceVolumeRequest) GoString() string {
	return s.String()
}

func (s *GetOnlineServiceVolumeRequest) SetInstanceId(v string) *GetOnlineServiceVolumeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetStartDate(v int64) *GetOnlineServiceVolumeRequest {
	s.StartDate = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetEndDate(v int64) *GetOnlineServiceVolumeRequest {
	s.EndDate = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetPageSize(v int32) *GetOnlineServiceVolumeRequest {
	s.PageSize = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetCurrentPage(v int32) *GetOnlineServiceVolumeRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetAgentIds(v []*int64) *GetOnlineServiceVolumeRequest {
	s.AgentIds = v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetDepIds(v []*int64) *GetOnlineServiceVolumeRequest {
	s.DepIds = v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetGroupIds(v []*int64) *GetOnlineServiceVolumeRequest {
	s.GroupIds = v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetTimeLatitudeType(v string) *GetOnlineServiceVolumeRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetExistAgentGrouping(v bool) *GetOnlineServiceVolumeRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetExistDepartmentGrouping(v bool) *GetOnlineServiceVolumeRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetOnlineServiceVolumeRequest) SetExistSkillGroupGrouping(v bool) *GetOnlineServiceVolumeRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetOnlineServiceVolumeShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// 技能组id列表
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// 时间纬度类型
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
	// 是否根据坐席分组
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetOnlineServiceVolumeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOnlineServiceVolumeShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetInstanceId(v string) *GetOnlineServiceVolumeShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetStartDate(v int64) *GetOnlineServiceVolumeShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetEndDate(v int64) *GetOnlineServiceVolumeShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetPageSize(v int32) *GetOnlineServiceVolumeShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetCurrentPage(v int32) *GetOnlineServiceVolumeShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetAgentIdsShrink(v string) *GetOnlineServiceVolumeShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetDepIdsShrink(v string) *GetOnlineServiceVolumeShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetGroupIdsShrink(v string) *GetOnlineServiceVolumeShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetTimeLatitudeType(v string) *GetOnlineServiceVolumeShrinkRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetExistAgentGrouping(v bool) *GetOnlineServiceVolumeShrinkRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetExistDepartmentGrouping(v bool) *GetOnlineServiceVolumeShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetOnlineServiceVolumeShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetOnlineServiceVolumeShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetOnlineServiceVolumeResponseBody struct {
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 调用接口是否成功
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// data
	Data *GetOnlineServiceVolumeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOnlineServiceVolumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOnlineServiceVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *GetOnlineServiceVolumeResponseBody) SetMessage(v string) *GetOnlineServiceVolumeResponseBody {
	s.Message = &v
	return s
}

func (s *GetOnlineServiceVolumeResponseBody) SetCode(v string) *GetOnlineServiceVolumeResponseBody {
	s.Code = &v
	return s
}

func (s *GetOnlineServiceVolumeResponseBody) SetSuccess(v string) *GetOnlineServiceVolumeResponseBody {
	s.Success = &v
	return s
}

func (s *GetOnlineServiceVolumeResponseBody) SetData(v *GetOnlineServiceVolumeResponseBodyData) *GetOnlineServiceVolumeResponseBody {
	s.Data = v
	return s
}

func (s *GetOnlineServiceVolumeResponseBody) SetRequestId(v string) *GetOnlineServiceVolumeResponseBody {
	s.RequestId = &v
	return s
}

type GetOnlineServiceVolumeResponseBodyData struct {
	// 当前页数
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总记录数
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s GetOnlineServiceVolumeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOnlineServiceVolumeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOnlineServiceVolumeResponseBodyData) SetPageNum(v int32) *GetOnlineServiceVolumeResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetOnlineServiceVolumeResponseBodyData) SetPageSize(v int32) *GetOnlineServiceVolumeResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetOnlineServiceVolumeResponseBodyData) SetTotalNum(v int32) *GetOnlineServiceVolumeResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetOnlineServiceVolumeResponseBodyData) SetRows(v string) *GetOnlineServiceVolumeResponseBodyData {
	s.Rows = &v
	return s
}

type GetOnlineServiceVolumeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOnlineServiceVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOnlineServiceVolumeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOnlineServiceVolumeResponse) GoString() string {
	return s.String()
}

func (s *GetOnlineServiceVolumeResponse) SetHeaders(v map[string]*string) *GetOnlineServiceVolumeResponse {
	s.Headers = v
	return s
}

func (s *GetOnlineServiceVolumeResponse) SetBody(v *GetOnlineServiceVolumeResponseBody) *GetOnlineServiceVolumeResponse {
	s.Body = v
	return s
}

type GetOutbounNumListRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetOutbounNumListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListRequest) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListRequest) SetClientToken(v string) *GetOutbounNumListRequest {
	s.ClientToken = &v
	return s
}

func (s *GetOutbounNumListRequest) SetInstanceId(v string) *GetOutbounNumListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOutbounNumListRequest) SetAccountName(v string) *GetOutbounNumListRequest {
	s.AccountName = &v
	return s
}

type GetOutbounNumListResponseBody struct {
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetOutbounNumListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GetOutbounNumListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBody) SetMessage(v string) *GetOutbounNumListResponseBody {
	s.Message = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetRequestId(v string) *GetOutbounNumListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetData(v *GetOutbounNumListResponseBodyData) *GetOutbounNumListResponseBody {
	s.Data = v
	return s
}

func (s *GetOutbounNumListResponseBody) SetCode(v string) *GetOutbounNumListResponseBody {
	s.Code = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetSuccess(v bool) *GetOutbounNumListResponseBody {
	s.Success = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetHttpStatusCode(v int64) *GetOutbounNumListResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GetOutbounNumListResponseBodyData struct {
	NumGroup []*GetOutbounNumListResponseBodyDataNumGroup `json:"NumGroup,omitempty" xml:"NumGroup,omitempty" type:"Repeated"`
	Num      []*GetOutbounNumListResponseBodyDataNum      `json:"Num,omitempty" xml:"Num,omitempty" type:"Repeated"`
}

func (s GetOutbounNumListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBodyData) SetNumGroup(v []*GetOutbounNumListResponseBodyDataNumGroup) *GetOutbounNumListResponseBodyData {
	s.NumGroup = v
	return s
}

func (s *GetOutbounNumListResponseBodyData) SetNum(v []*GetOutbounNumListResponseBodyDataNum) *GetOutbounNumListResponseBodyData {
	s.Num = v
	return s
}

type GetOutbounNumListResponseBodyDataNumGroup struct {
	Type        *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetOutbounNumListResponseBodyDataNumGroup) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponseBodyDataNumGroup) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) SetType(v int32) *GetOutbounNumListResponseBodyDataNumGroup {
	s.Type = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) SetValue(v string) *GetOutbounNumListResponseBodyDataNumGroup {
	s.Value = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) SetDescription(v string) *GetOutbounNumListResponseBodyDataNumGroup {
	s.Description = &v
	return s
}

type GetOutbounNumListResponseBodyDataNum struct {
	Type        *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetOutbounNumListResponseBodyDataNum) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponseBodyDataNum) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBodyDataNum) SetType(v int32) *GetOutbounNumListResponseBodyDataNum {
	s.Type = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNum) SetValue(v string) *GetOutbounNumListResponseBodyDataNum {
	s.Value = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNum) SetDescription(v string) *GetOutbounNumListResponseBodyDataNum {
	s.Description = &v
	return s
}

type GetOutbounNumListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOutbounNumListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOutbounNumListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponse) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponse) SetHeaders(v map[string]*string) *GetOutbounNumListResponse {
	s.Headers = v
	return s
}

func (s *GetOutbounNumListResponse) SetBody(v *GetOutbounNumListResponseBody) *GetOutbounNumListResponse {
	s.Body = v
	return s
}

type GetQualityProjectDetailRequest struct {
	// 租户实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 质检任务ID
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetQualityProjectDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQualityProjectDetailRequest) GoString() string {
	return s.String()
}

func (s *GetQualityProjectDetailRequest) SetInstanceId(v string) *GetQualityProjectDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQualityProjectDetailRequest) SetProjectId(v int64) *GetQualityProjectDetailRequest {
	s.ProjectId = &v
	return s
}

type GetQualityProjectDetailResponseBody struct {
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Data
	Data *GetQualityProjectDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Success
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityProjectDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQualityProjectDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityProjectDetailResponseBody) SetMessage(v string) *GetQualityProjectDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityProjectDetailResponseBody) SetRequestId(v string) *GetQualityProjectDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityProjectDetailResponseBody) SetData(v *GetQualityProjectDetailResponseBodyData) *GetQualityProjectDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityProjectDetailResponseBody) SetCode(v string) *GetQualityProjectDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityProjectDetailResponseBody) SetSuccess(v string) *GetQualityProjectDetailResponseBody {
	s.Success = &v
	return s
}

type GetQualityProjectDetailResponseBodyData struct {
	// 质检任务状态
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 质检类型
	QualityType *int32 `json:"QualityType,omitempty" xml:"QualityType,omitempty"`
	// 质检规则ID
	QualityRuleIds []*int64 `json:"QualityRuleIds,omitempty" xml:"QualityRuleIds,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 质检任务名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 质检周期
	CheckFreqType *int32 `json:"CheckFreqType,omitempty" xml:"CheckFreqType,omitempty"`
	// 技能组分组
	DepList []*int64 `json:"DepList,omitempty" xml:"DepList,omitempty" type:"Repeated"`
	// 坐席列表
	ServicerList []*int64 `json:"ServicerList,omitempty" xml:"ServicerList,omitempty" type:"Repeated"`
	// Version
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// 技能组列表
	GroupList []*int64 `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	// 质检任务ID
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s GetQualityProjectDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQualityProjectDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityProjectDetailResponseBodyData) SetStatus(v int32) *GetQualityProjectDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetQualityType(v int32) *GetQualityProjectDetailResponseBodyData {
	s.QualityType = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetQualityRuleIds(v []*int64) *GetQualityProjectDetailResponseBodyData {
	s.QualityRuleIds = v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetCreateTime(v string) *GetQualityProjectDetailResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetProjectName(v string) *GetQualityProjectDetailResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetCheckFreqType(v int32) *GetQualityProjectDetailResponseBodyData {
	s.CheckFreqType = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetDepList(v []*int64) *GetQualityProjectDetailResponseBodyData {
	s.DepList = v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetServicerList(v []*int64) *GetQualityProjectDetailResponseBodyData {
	s.ServicerList = v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetVersion(v int32) *GetQualityProjectDetailResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetGroupList(v []*int64) *GetQualityProjectDetailResponseBodyData {
	s.GroupList = v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetId(v int64) *GetQualityProjectDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetModifyTime(v string) *GetQualityProjectDetailResponseBodyData {
	s.ModifyTime = &v
	return s
}

type GetQualityProjectDetailResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQualityProjectDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQualityProjectDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQualityProjectDetailResponse) GoString() string {
	return s.String()
}

func (s *GetQualityProjectDetailResponse) SetHeaders(v map[string]*string) *GetQualityProjectDetailResponse {
	s.Headers = v
	return s
}

func (s *GetQualityProjectDetailResponse) SetBody(v *GetQualityProjectDetailResponseBody) *GetQualityProjectDetailResponse {
	s.Body = v
	return s
}

type GetQualityProjectListRequest struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 质检项ID
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 质检项名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 质检项状态
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// PageNo
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// PageSize
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 质检频率
	CheckFreqType *int64 `json:"checkFreqType,omitempty" xml:"checkFreqType,omitempty"`
}

func (s GetQualityProjectListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQualityProjectListRequest) GoString() string {
	return s.String()
}

func (s *GetQualityProjectListRequest) SetInstanceId(v string) *GetQualityProjectListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQualityProjectListRequest) SetProjectId(v int64) *GetQualityProjectListRequest {
	s.ProjectId = &v
	return s
}

func (s *GetQualityProjectListRequest) SetProjectName(v string) *GetQualityProjectListRequest {
	s.ProjectName = &v
	return s
}

func (s *GetQualityProjectListRequest) SetStatus(v int32) *GetQualityProjectListRequest {
	s.Status = &v
	return s
}

func (s *GetQualityProjectListRequest) SetPageNo(v int32) *GetQualityProjectListRequest {
	s.PageNo = &v
	return s
}

func (s *GetQualityProjectListRequest) SetPageSize(v int32) *GetQualityProjectListRequest {
	s.PageSize = &v
	return s
}

func (s *GetQualityProjectListRequest) SetCheckFreqType(v int64) *GetQualityProjectListRequest {
	s.CheckFreqType = &v
	return s
}

type GetQualityProjectListResponseBody struct {
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Data
	Data *GetQualityProjectListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityProjectListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQualityProjectListResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityProjectListResponseBody) SetMessage(v string) *GetQualityProjectListResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityProjectListResponseBody) SetRequestId(v string) *GetQualityProjectListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityProjectListResponseBody) SetData(v *GetQualityProjectListResponseBodyData) *GetQualityProjectListResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityProjectListResponseBody) SetCode(v string) *GetQualityProjectListResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityProjectListResponseBody) SetSuccess(v bool) *GetQualityProjectListResponseBody {
	s.Success = &v
	return s
}

type GetQualityProjectListResponseBodyData struct {
	// 质检项列表
	QualityProjectList []*GetQualityProjectListResponseBodyDataQualityProjectList `json:"QualityProjectList,omitempty" xml:"QualityProjectList,omitempty" type:"Repeated"`
	// PageNo
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// PageSize
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQualityProjectListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQualityProjectListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityProjectListResponseBodyData) SetQualityProjectList(v []*GetQualityProjectListResponseBodyDataQualityProjectList) *GetQualityProjectListResponseBodyData {
	s.QualityProjectList = v
	return s
}

func (s *GetQualityProjectListResponseBodyData) SetPageNo(v int32) *GetQualityProjectListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQualityProjectListResponseBodyData) SetPageSize(v int32) *GetQualityProjectListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQualityProjectListResponseBodyData) SetTotal(v int64) *GetQualityProjectListResponseBodyData {
	s.Total = &v
	return s
}

type GetQualityProjectListResponseBodyDataQualityProjectList struct {
	// 质检任务状态
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 质检任务类型
	QualityType *int32 `json:"QualityType,omitempty" xml:"QualityType,omitempty"`
	// 质检规则列表
	QualityRuleIds []*int64 `json:"QualityRuleIds,omitempty" xml:"QualityRuleIds,omitempty" type:"Repeated"`
	// CreateTime
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 质检任务名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 质检任务频率
	CheckFreqType *int32 `json:"CheckFreqType,omitempty" xml:"CheckFreqType,omitempty"`
	// 技能组分组列表
	DepList []*int64 `json:"DepList,omitempty" xml:"DepList,omitempty" type:"Repeated"`
	// 坐席列表
	ServicerList []*int64 `json:"ServicerList,omitempty" xml:"ServicerList,omitempty" type:"Repeated"`
	// 版本
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// 技能组列表
	GroupList []*int64 `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	// 质检任务Id
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s GetQualityProjectListResponseBodyDataQualityProjectList) String() string {
	return tea.Prettify(s)
}

func (s GetQualityProjectListResponseBodyDataQualityProjectList) GoString() string {
	return s.String()
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetStatus(v int32) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.Status = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetQualityType(v int32) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.QualityType = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetQualityRuleIds(v []*int64) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.QualityRuleIds = v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetCreateTime(v string) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.CreateTime = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetProjectName(v string) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.ProjectName = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetCheckFreqType(v int32) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.CheckFreqType = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetDepList(v []*int64) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.DepList = v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetServicerList(v []*int64) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.ServicerList = v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetVersion(v int32) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.Version = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetGroupList(v []*int64) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.GroupList = v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetId(v int64) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.Id = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetModifyTime(v string) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.ModifyTime = &v
	return s
}

type GetQualityProjectListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQualityProjectListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQualityProjectListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQualityProjectListResponse) GoString() string {
	return s.String()
}

func (s *GetQualityProjectListResponse) SetHeaders(v map[string]*string) *GetQualityProjectListResponse {
	s.Headers = v
	return s
}

func (s *GetQualityProjectListResponse) SetBody(v *GetQualityProjectListResponseBody) *GetQualityProjectListResponse {
	s.Body = v
	return s
}

type GetQualityProjectLogRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProjectId  *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetQualityProjectLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQualityProjectLogRequest) GoString() string {
	return s.String()
}

func (s *GetQualityProjectLogRequest) SetInstanceId(v string) *GetQualityProjectLogRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQualityProjectLogRequest) SetProjectId(v int64) *GetQualityProjectLogRequest {
	s.ProjectId = &v
	return s
}

type GetQualityProjectLogResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      []*GetQualityProjectLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetQualityProjectLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQualityProjectLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityProjectLogResponseBody) SetCode(v string) *GetQualityProjectLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityProjectLogResponseBody) SetMessage(v string) *GetQualityProjectLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityProjectLogResponseBody) SetRequestId(v string) *GetQualityProjectLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityProjectLogResponseBody) SetSuccess(v bool) *GetQualityProjectLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityProjectLogResponseBody) SetData(v []*GetQualityProjectLogResponseBodyData) *GetQualityProjectLogResponseBody {
	s.Data = v
	return s
}

type GetQualityProjectLogResponseBodyData struct {
	ProjectCreateTime *string `json:"ProjectCreateTime,omitempty" xml:"ProjectCreateTime,omitempty"`
	ActionType        *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	ActionTime        *string `json:"ActionTime,omitempty" xml:"ActionTime,omitempty"`
	ActionData        *string `json:"ActionData,omitempty" xml:"ActionData,omitempty"`
	ProjectId         *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetQualityProjectLogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQualityProjectLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityProjectLogResponseBodyData) SetProjectCreateTime(v string) *GetQualityProjectLogResponseBodyData {
	s.ProjectCreateTime = &v
	return s
}

func (s *GetQualityProjectLogResponseBodyData) SetActionType(v string) *GetQualityProjectLogResponseBodyData {
	s.ActionType = &v
	return s
}

func (s *GetQualityProjectLogResponseBodyData) SetActionTime(v string) *GetQualityProjectLogResponseBodyData {
	s.ActionTime = &v
	return s
}

func (s *GetQualityProjectLogResponseBodyData) SetActionData(v string) *GetQualityProjectLogResponseBodyData {
	s.ActionData = &v
	return s
}

func (s *GetQualityProjectLogResponseBodyData) SetProjectId(v int64) *GetQualityProjectLogResponseBodyData {
	s.ProjectId = &v
	return s
}

type GetQualityProjectLogResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQualityProjectLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQualityProjectLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQualityProjectLogResponse) GoString() string {
	return s.String()
}

func (s *GetQualityProjectLogResponse) SetHeaders(v map[string]*string) *GetQualityProjectLogResponse {
	s.Headers = v
	return s
}

func (s *GetQualityProjectLogResponse) SetBody(v *GetQualityProjectLogResponseBody) *GetQualityProjectLogResponse {
	s.Body = v
	return s
}

type GetQualityResultRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo         *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TouchStartTime *string `json:"TouchStartTime,omitempty" xml:"TouchStartTime,omitempty"`
	TouchEndTime   *string `json:"TouchEndTime,omitempty" xml:"TouchEndTime,omitempty"`
	ChannelType    *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	HitStatus      *int32  `json:"HitStatus,omitempty" xml:"HitStatus,omitempty"`
	GroupIds       []*int  `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	QualityRuleIds []*int  `json:"QualityRuleIds,omitempty" xml:"QualityRuleIds,omitempty" type:"Repeated"`
	ProjectIds     []*int  `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
}

func (s GetQualityResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQualityResultRequest) GoString() string {
	return s.String()
}

func (s *GetQualityResultRequest) SetInstanceId(v string) *GetQualityResultRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQualityResultRequest) SetPageNo(v int32) *GetQualityResultRequest {
	s.PageNo = &v
	return s
}

func (s *GetQualityResultRequest) SetPageSize(v int32) *GetQualityResultRequest {
	s.PageSize = &v
	return s
}

func (s *GetQualityResultRequest) SetTouchStartTime(v string) *GetQualityResultRequest {
	s.TouchStartTime = &v
	return s
}

func (s *GetQualityResultRequest) SetTouchEndTime(v string) *GetQualityResultRequest {
	s.TouchEndTime = &v
	return s
}

func (s *GetQualityResultRequest) SetChannelType(v string) *GetQualityResultRequest {
	s.ChannelType = &v
	return s
}

func (s *GetQualityResultRequest) SetHitStatus(v int32) *GetQualityResultRequest {
	s.HitStatus = &v
	return s
}

func (s *GetQualityResultRequest) SetGroupIds(v []*int) *GetQualityResultRequest {
	s.GroupIds = v
	return s
}

func (s *GetQualityResultRequest) SetQualityRuleIds(v []*int) *GetQualityResultRequest {
	s.QualityRuleIds = v
	return s
}

func (s *GetQualityResultRequest) SetProjectIds(v []*int) *GetQualityResultRequest {
	s.ProjectIds = v
	return s
}

type GetQualityResultResponseBody struct {
	Code            *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	ChannelTypeName *string                           `json:"ChannelTypeName,omitempty" xml:"ChannelTypeName,omitempty"`
	RequestId       *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Data            *GetQualityResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetQualityResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQualityResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityResultResponseBody) SetCode(v string) *GetQualityResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityResultResponseBody) SetMessage(v string) *GetQualityResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityResultResponseBody) SetChannelTypeName(v string) *GetQualityResultResponseBody {
	s.ChannelTypeName = &v
	return s
}

func (s *GetQualityResultResponseBody) SetRequestId(v string) *GetQualityResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityResultResponseBody) SetSuccess(v bool) *GetQualityResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityResultResponseBody) SetData(v *GetQualityResultResponseBodyData) *GetQualityResultResponseBody {
	s.Data = v
	return s
}

type GetQualityResultResponseBodyData struct {
	PageNo                    *int32                                                       `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize                  *int32                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNum                  *int32                                                       `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	QualityResultResponseList []*GetQualityResultResponseBodyDataQualityResultResponseList `json:"QualityResultResponseList,omitempty" xml:"QualityResultResponseList,omitempty" type:"Repeated"`
}

func (s GetQualityResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQualityResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityResultResponseBodyData) SetPageNo(v int32) *GetQualityResultResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQualityResultResponseBodyData) SetPageSize(v int32) *GetQualityResultResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQualityResultResponseBodyData) SetTotalNum(v int32) *GetQualityResultResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetQualityResultResponseBodyData) SetQualityResultResponseList(v []*GetQualityResultResponseBodyDataQualityResultResponseList) *GetQualityResultResponseBodyData {
	s.QualityResultResponseList = v
	return s
}

type GetQualityResultResponseBodyDataQualityResultResponseList struct {
	TouchId         *string `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
	MemberName      *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	ServicerName    *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	ChannelType     *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ChannelTypeName *string `json:"ChannelTypeName,omitempty" xml:"ChannelTypeName,omitempty"`
	TouchStartTime  *string `json:"TouchStartTime,omitempty" xml:"TouchStartTime,omitempty"`
	ServicerId      *string `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	RuleId          *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	HitStatus       *bool   `json:"HitStatus,omitempty" xml:"HitStatus,omitempty"`
	HitDetail       *string `json:"HitDetail,omitempty" xml:"HitDetail,omitempty"`
}

func (s GetQualityResultResponseBodyDataQualityResultResponseList) String() string {
	return tea.Prettify(s)
}

func (s GetQualityResultResponseBodyDataQualityResultResponseList) GoString() string {
	return s.String()
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetTouchId(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.TouchId = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetMemberName(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.MemberName = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetServicerName(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.ServicerName = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetChannelType(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.ChannelType = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetProjectId(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.ProjectId = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetProjectName(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.ProjectName = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetChannelTypeName(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.ChannelTypeName = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetTouchStartTime(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.TouchStartTime = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetServicerId(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.ServicerId = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetRuleId(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.RuleId = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetRuleName(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.RuleName = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetGroupName(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.GroupName = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetGroupId(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.GroupId = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetInstanceName(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.InstanceName = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetHitStatus(v bool) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.HitStatus = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetHitDetail(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.HitDetail = &v
	return s
}

type GetQualityResultResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQualityResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQualityResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQualityResultResponse) GoString() string {
	return s.String()
}

func (s *GetQualityResultResponse) SetHeaders(v map[string]*string) *GetQualityResultResponse {
	s.Headers = v
	return s
}

func (s *GetQualityResultResponse) SetBody(v *GetQualityResultResponseBody) *GetQualityResultResponse {
	s.Body = v
	return s
}

type GetQualityRuleDetailRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QualityRuleId *int64  `json:"QualityRuleId,omitempty" xml:"QualityRuleId,omitempty"`
}

func (s GetQualityRuleDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQualityRuleDetailRequest) GoString() string {
	return s.String()
}

func (s *GetQualityRuleDetailRequest) SetInstanceId(v string) *GetQualityRuleDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQualityRuleDetailRequest) SetQualityRuleId(v int64) *GetQualityRuleDetailRequest {
	s.QualityRuleId = &v
	return s
}

type GetQualityRuleDetailResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *GetQualityRuleDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetQualityRuleDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQualityRuleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityRuleDetailResponseBody) SetCode(v string) *GetQualityRuleDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityRuleDetailResponseBody) SetMessage(v string) *GetQualityRuleDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityRuleDetailResponseBody) SetRequestId(v string) *GetQualityRuleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityRuleDetailResponseBody) SetSuccess(v bool) *GetQualityRuleDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityRuleDetailResponseBody) SetData(v *GetQualityRuleDetailResponseBodyData) *GetQualityRuleDetailResponseBody {
	s.Data = v
	return s
}

type GetQualityRuleDetailResponseBodyData struct {
	RuleTag        *int32    `json:"RuleTag,omitempty" xml:"RuleTag,omitempty"`
	MatchType      *int32    `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleCreateTime *string   `json:"RuleCreateTime,omitempty" xml:"RuleCreateTime,omitempty"`
	RuleId         *int64    `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	KeyWords       []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
}

func (s GetQualityRuleDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQualityRuleDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityRuleDetailResponseBodyData) SetRuleTag(v int32) *GetQualityRuleDetailResponseBodyData {
	s.RuleTag = &v
	return s
}

func (s *GetQualityRuleDetailResponseBodyData) SetMatchType(v int32) *GetQualityRuleDetailResponseBodyData {
	s.MatchType = &v
	return s
}

func (s *GetQualityRuleDetailResponseBodyData) SetName(v string) *GetQualityRuleDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetQualityRuleDetailResponseBodyData) SetRuleCreateTime(v string) *GetQualityRuleDetailResponseBodyData {
	s.RuleCreateTime = &v
	return s
}

func (s *GetQualityRuleDetailResponseBodyData) SetRuleId(v int64) *GetQualityRuleDetailResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *GetQualityRuleDetailResponseBodyData) SetKeyWords(v []*string) *GetQualityRuleDetailResponseBodyData {
	s.KeyWords = v
	return s
}

type GetQualityRuleDetailResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQualityRuleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQualityRuleDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQualityRuleDetailResponse) GoString() string {
	return s.String()
}

func (s *GetQualityRuleDetailResponse) SetHeaders(v map[string]*string) *GetQualityRuleDetailResponse {
	s.Headers = v
	return s
}

func (s *GetQualityRuleDetailResponse) SetBody(v *GetQualityRuleDetailResponseBody) *GetQualityRuleDetailResponse {
	s.Body = v
	return s
}

type GetQualityRuleListRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo     *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetQualityRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQualityRuleListRequest) GoString() string {
	return s.String()
}

func (s *GetQualityRuleListRequest) SetInstanceId(v string) *GetQualityRuleListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQualityRuleListRequest) SetPageNo(v int32) *GetQualityRuleListRequest {
	s.PageNo = &v
	return s
}

func (s *GetQualityRuleListRequest) SetPageSize(v int32) *GetQualityRuleListRequest {
	s.PageSize = &v
	return s
}

type GetQualityRuleListResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *GetQualityRuleListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetQualityRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQualityRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityRuleListResponseBody) SetCode(v string) *GetQualityRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityRuleListResponseBody) SetMessage(v string) *GetQualityRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityRuleListResponseBody) SetRequestId(v string) *GetQualityRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityRuleListResponseBody) SetSuccess(v bool) *GetQualityRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityRuleListResponseBody) SetData(v *GetQualityRuleListResponseBodyData) *GetQualityRuleListResponseBody {
	s.Data = v
	return s
}

type GetQualityRuleListResponseBodyData struct {
	PageNo          *int32                                               `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total           *int64                                               `json:"Total,omitempty" xml:"Total,omitempty"`
	QualityRuleList []*GetQualityRuleListResponseBodyDataQualityRuleList `json:"QualityRuleList,omitempty" xml:"QualityRuleList,omitempty" type:"Repeated"`
}

func (s GetQualityRuleListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQualityRuleListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityRuleListResponseBodyData) SetPageNo(v int32) *GetQualityRuleListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQualityRuleListResponseBodyData) SetPageSize(v int32) *GetQualityRuleListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQualityRuleListResponseBodyData) SetTotal(v int64) *GetQualityRuleListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetQualityRuleListResponseBodyData) SetQualityRuleList(v []*GetQualityRuleListResponseBodyDataQualityRuleList) *GetQualityRuleListResponseBodyData {
	s.QualityRuleList = v
	return s
}

type GetQualityRuleListResponseBodyDataQualityRuleList struct {
	RuleTag        *int32    `json:"RuleTag,omitempty" xml:"RuleTag,omitempty"`
	MatchType      *int32    `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleCreateTime *string   `json:"RuleCreateTime,omitempty" xml:"RuleCreateTime,omitempty"`
	RuleId         *int64    `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	KeyWords       []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
}

func (s GetQualityRuleListResponseBodyDataQualityRuleList) String() string {
	return tea.Prettify(s)
}

func (s GetQualityRuleListResponseBodyDataQualityRuleList) GoString() string {
	return s.String()
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) SetRuleTag(v int32) *GetQualityRuleListResponseBodyDataQualityRuleList {
	s.RuleTag = &v
	return s
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) SetMatchType(v int32) *GetQualityRuleListResponseBodyDataQualityRuleList {
	s.MatchType = &v
	return s
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) SetName(v string) *GetQualityRuleListResponseBodyDataQualityRuleList {
	s.Name = &v
	return s
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) SetRuleCreateTime(v string) *GetQualityRuleListResponseBodyDataQualityRuleList {
	s.RuleCreateTime = &v
	return s
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) SetRuleId(v int64) *GetQualityRuleListResponseBodyDataQualityRuleList {
	s.RuleId = &v
	return s
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) SetKeyWords(v []*string) *GetQualityRuleListResponseBodyDataQualityRuleList {
	s.KeyWords = v
	return s
}

type GetQualityRuleListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQualityRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQualityRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQualityRuleListResponse) GoString() string {
	return s.String()
}

func (s *GetQualityRuleListResponse) SetHeaders(v map[string]*string) *GetQualityRuleListResponse {
	s.Headers = v
	return s
}

func (s *GetQualityRuleListResponse) SetBody(v *GetQualityRuleListResponseBody) *GetQualityRuleListResponse {
	s.Body = v
	return s
}

type GetQualityRuleTagListRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetQualityRuleTagListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQualityRuleTagListRequest) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTagListRequest) SetInstanceId(v string) *GetQualityRuleTagListRequest {
	s.InstanceId = &v
	return s
}

type GetQualityRuleTagListResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      []*GetQualityRuleTagListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetQualityRuleTagListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQualityRuleTagListResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTagListResponseBody) SetCode(v string) *GetQualityRuleTagListResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityRuleTagListResponseBody) SetMessage(v string) *GetQualityRuleTagListResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityRuleTagListResponseBody) SetRequestId(v string) *GetQualityRuleTagListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityRuleTagListResponseBody) SetSuccess(v bool) *GetQualityRuleTagListResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityRuleTagListResponseBody) SetData(v []*GetQualityRuleTagListResponseBodyData) *GetQualityRuleTagListResponseBody {
	s.Data = v
	return s
}

type GetQualityRuleTagListResponseBodyData struct {
	RuleTagName *string `json:"RuleTagName,omitempty" xml:"RuleTagName,omitempty"`
	RuleTagId   *int64  `json:"RuleTagId,omitempty" xml:"RuleTagId,omitempty"`
}

func (s GetQualityRuleTagListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQualityRuleTagListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTagListResponseBodyData) SetRuleTagName(v string) *GetQualityRuleTagListResponseBodyData {
	s.RuleTagName = &v
	return s
}

func (s *GetQualityRuleTagListResponseBodyData) SetRuleTagId(v int64) *GetQualityRuleTagListResponseBodyData {
	s.RuleTagId = &v
	return s
}

type GetQualityRuleTagListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQualityRuleTagListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQualityRuleTagListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQualityRuleTagListResponse) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTagListResponse) SetHeaders(v map[string]*string) *GetQualityRuleTagListResponse {
	s.Headers = v
	return s
}

func (s *GetQualityRuleTagListResponse) SetBody(v *GetQualityRuleTagListResponseBody) *GetQualityRuleTagListResponse {
	s.Body = v
	return s
}

type GetQueueInformationRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 部门id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// 技能组id列表
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetQueueInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQueueInformationRequest) GoString() string {
	return s.String()
}

func (s *GetQueueInformationRequest) SetInstanceId(v string) *GetQueueInformationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQueueInformationRequest) SetStartDate(v int64) *GetQueueInformationRequest {
	s.StartDate = &v
	return s
}

func (s *GetQueueInformationRequest) SetEndDate(v int64) *GetQueueInformationRequest {
	s.EndDate = &v
	return s
}

func (s *GetQueueInformationRequest) SetPageSize(v int32) *GetQueueInformationRequest {
	s.PageSize = &v
	return s
}

func (s *GetQueueInformationRequest) SetCurrentPage(v int32) *GetQueueInformationRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetQueueInformationRequest) SetDepIds(v []*int64) *GetQueueInformationRequest {
	s.DepIds = v
	return s
}

func (s *GetQueueInformationRequest) SetGroupIds(v []*int64) *GetQueueInformationRequest {
	s.GroupIds = v
	return s
}

func (s *GetQueueInformationRequest) SetExistDepartmentGrouping(v bool) *GetQueueInformationRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetQueueInformationRequest) SetExistSkillGroupGrouping(v bool) *GetQueueInformationRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetQueueInformationShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// 技能组id列表
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetQueueInformationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQueueInformationShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetQueueInformationShrinkRequest) SetInstanceId(v string) *GetQueueInformationShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetStartDate(v int64) *GetQueueInformationShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetEndDate(v int64) *GetQueueInformationShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetPageSize(v int32) *GetQueueInformationShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetCurrentPage(v int32) *GetQueueInformationShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetDepIdsShrink(v string) *GetQueueInformationShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetGroupIdsShrink(v string) *GetQueueInformationShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetExistDepartmentGrouping(v bool) *GetQueueInformationShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetQueueInformationShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetQueueInformationShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetQueueInformationResponseBody struct {
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 调用接口是否成功
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// data
	Data *GetQueueInformationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetQueueInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQueueInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueueInformationResponseBody) SetRequestId(v string) *GetQueueInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueueInformationResponseBody) SetMessage(v string) *GetQueueInformationResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueueInformationResponseBody) SetCode(v string) *GetQueueInformationResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueueInformationResponseBody) SetSuccess(v string) *GetQueueInformationResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueueInformationResponseBody) SetData(v *GetQueueInformationResponseBodyData) *GetQueueInformationResponseBody {
	s.Data = v
	return s
}

type GetQueueInformationResponseBodyData struct {
	// 当前页数
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总记录数
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s GetQueueInformationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQueueInformationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueueInformationResponseBodyData) SetPageNum(v int32) *GetQueueInformationResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetQueueInformationResponseBodyData) SetPageSize(v int32) *GetQueueInformationResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueueInformationResponseBodyData) SetTotalNum(v int32) *GetQueueInformationResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetQueueInformationResponseBodyData) SetRows(v string) *GetQueueInformationResponseBodyData {
	s.Rows = &v
	return s
}

type GetQueueInformationResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQueueInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQueueInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQueueInformationResponse) GoString() string {
	return s.String()
}

func (s *GetQueueInformationResponse) SetHeaders(v map[string]*string) *GetQueueInformationResponse {
	s.Headers = v
	return s
}

func (s *GetQueueInformationResponse) SetBody(v *GetQueueInformationResponseBody) *GetQueueInformationResponse {
	s.Body = v
	return s
}

type GetRecordDataRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Acid       *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
}

func (s GetRecordDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecordDataRequest) GoString() string {
	return s.String()
}

func (s *GetRecordDataRequest) SetInstanceId(v string) *GetRecordDataRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRecordDataRequest) SetAcid(v string) *GetRecordDataRequest {
	s.Acid = &v
	return s
}

type GetRecordDataResponseBody struct {
	Message        *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetRecordDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code           *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GetRecordDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecordDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordDataResponseBody) SetMessage(v string) *GetRecordDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetRecordDataResponseBody) SetRequestId(v string) *GetRecordDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecordDataResponseBody) SetData(v *GetRecordDataResponseBodyData) *GetRecordDataResponseBody {
	s.Data = v
	return s
}

func (s *GetRecordDataResponseBody) SetCode(v string) *GetRecordDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetRecordDataResponseBody) SetSuccess(v bool) *GetRecordDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetRecordDataResponseBody) SetHttpStatusCode(v int64) *GetRecordDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GetRecordDataResponseBodyData struct {
	OssLink *string `json:"OssLink,omitempty" xml:"OssLink,omitempty"`
	Acid    *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
}

func (s GetRecordDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRecordDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRecordDataResponseBodyData) SetOssLink(v string) *GetRecordDataResponseBodyData {
	s.OssLink = &v
	return s
}

func (s *GetRecordDataResponseBodyData) SetAcid(v string) *GetRecordDataResponseBodyData {
	s.Acid = &v
	return s
}

type GetRecordDataResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRecordDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRecordDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecordDataResponse) GoString() string {
	return s.String()
}

func (s *GetRecordDataResponse) SetHeaders(v map[string]*string) *GetRecordDataResponse {
	s.Headers = v
	return s
}

func (s *GetRecordDataResponse) SetBody(v *GetRecordDataResponseBody) *GetRecordDataResponse {
	s.Body = v
	return s
}

type GetRtcTokenRequest struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 账号名称
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetRtcTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRtcTokenRequest) GoString() string {
	return s.String()
}

func (s *GetRtcTokenRequest) SetInstanceId(v string) *GetRtcTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRtcTokenRequest) SetAccountName(v string) *GetRtcTokenRequest {
	s.AccountName = &v
	return s
}

type GetRtcTokenResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// data
	Data *GetRtcTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetRtcTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRtcTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetRtcTokenResponseBody) SetRequestId(v string) *GetRtcTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRtcTokenResponseBody) SetSuccess(v bool) *GetRtcTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetRtcTokenResponseBody) SetCode(v string) *GetRtcTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetRtcTokenResponseBody) SetMessage(v string) *GetRtcTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetRtcTokenResponseBody) SetData(v *GetRtcTokenResponseBodyData) *GetRtcTokenResponseBody {
	s.Data = v
	return s
}

type GetRtcTokenResponseBodyData struct {
	// token信息
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// rtcId
	RtcId *string `json:"RtcId,omitempty" xml:"RtcId,omitempty"`
	// 账号名
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetRtcTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRtcTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRtcTokenResponseBodyData) SetToken(v string) *GetRtcTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetRtcTokenResponseBodyData) SetRtcId(v string) *GetRtcTokenResponseBodyData {
	s.RtcId = &v
	return s
}

func (s *GetRtcTokenResponseBodyData) SetAccountName(v string) *GetRtcTokenResponseBodyData {
	s.AccountName = &v
	return s
}

type GetRtcTokenResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRtcTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRtcTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRtcTokenResponse) GoString() string {
	return s.String()
}

func (s *GetRtcTokenResponse) SetHeaders(v map[string]*string) *GetRtcTokenResponse {
	s.Headers = v
	return s
}

func (s *GetRtcTokenResponse) SetBody(v *GetRtcTokenResponseBody) *GetRtcTokenResponse {
	s.Body = v
	return s
}

type GetSeatInformationRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 部门id列表
	DepIds []*int64 `json:"depIds,omitempty" xml:"depIds,omitempty" type:"Repeated"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"existDepartmentGrouping,omitempty" xml:"existDepartmentGrouping,omitempty"`
}

func (s GetSeatInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSeatInformationRequest) GoString() string {
	return s.String()
}

func (s *GetSeatInformationRequest) SetInstanceId(v string) *GetSeatInformationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSeatInformationRequest) SetStartDate(v int64) *GetSeatInformationRequest {
	s.StartDate = &v
	return s
}

func (s *GetSeatInformationRequest) SetEndDate(v int64) *GetSeatInformationRequest {
	s.EndDate = &v
	return s
}

func (s *GetSeatInformationRequest) SetPageSize(v int32) *GetSeatInformationRequest {
	s.PageSize = &v
	return s
}

func (s *GetSeatInformationRequest) SetCurrentPage(v int32) *GetSeatInformationRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSeatInformationRequest) SetDepIds(v []*int64) *GetSeatInformationRequest {
	s.DepIds = v
	return s
}

func (s *GetSeatInformationRequest) SetExistDepartmentGrouping(v bool) *GetSeatInformationRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

type GetSeatInformationShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"depIds,omitempty" xml:"depIds,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"existDepartmentGrouping,omitempty" xml:"existDepartmentGrouping,omitempty"`
}

func (s GetSeatInformationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSeatInformationShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSeatInformationShrinkRequest) SetInstanceId(v string) *GetSeatInformationShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSeatInformationShrinkRequest) SetStartDate(v int64) *GetSeatInformationShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetSeatInformationShrinkRequest) SetEndDate(v int64) *GetSeatInformationShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetSeatInformationShrinkRequest) SetPageSize(v int32) *GetSeatInformationShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSeatInformationShrinkRequest) SetCurrentPage(v int32) *GetSeatInformationShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSeatInformationShrinkRequest) SetDepIdsShrink(v string) *GetSeatInformationShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetSeatInformationShrinkRequest) SetExistDepartmentGrouping(v bool) *GetSeatInformationShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

type GetSeatInformationResponseBody struct {
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 调用接口是否成功
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// data
	Data *GetSeatInformationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetSeatInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSeatInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetSeatInformationResponseBody) SetRequestId(v string) *GetSeatInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSeatInformationResponseBody) SetMessage(v string) *GetSeatInformationResponseBody {
	s.Message = &v
	return s
}

func (s *GetSeatInformationResponseBody) SetCode(v string) *GetSeatInformationResponseBody {
	s.Code = &v
	return s
}

func (s *GetSeatInformationResponseBody) SetSuccess(v string) *GetSeatInformationResponseBody {
	s.Success = &v
	return s
}

func (s *GetSeatInformationResponseBody) SetData(v *GetSeatInformationResponseBodyData) *GetSeatInformationResponseBody {
	s.Data = v
	return s
}

type GetSeatInformationResponseBodyData struct {
	// 当前页数
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总记录数
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rowr *string `json:"Rowr,omitempty" xml:"Rowr,omitempty"`
}

func (s GetSeatInformationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSeatInformationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSeatInformationResponseBodyData) SetPageNum(v int32) *GetSeatInformationResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetSeatInformationResponseBodyData) SetPageSize(v int32) *GetSeatInformationResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSeatInformationResponseBodyData) SetTotalNum(v int32) *GetSeatInformationResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetSeatInformationResponseBodyData) SetRowr(v string) *GetSeatInformationResponseBodyData {
	s.Rowr = &v
	return s
}

type GetSeatInformationResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSeatInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSeatInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSeatInformationResponse) GoString() string {
	return s.String()
}

func (s *GetSeatInformationResponse) SetHeaders(v map[string]*string) *GetSeatInformationResponse {
	s.Headers = v
	return s
}

func (s *GetSeatInformationResponse) SetBody(v *GetSeatInformationResponseBody) *GetSeatInformationResponse {
	s.Body = v
	return s
}

type GetSkillGroupAgentStatusDetailsRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 部门id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// 技能组id列表
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetSkillGroupAgentStatusDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupAgentStatusDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetInstanceId(v string) *GetSkillGroupAgentStatusDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetStartDate(v int64) *GetSkillGroupAgentStatusDetailsRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetEndDate(v int64) *GetSkillGroupAgentStatusDetailsRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetPageSize(v int32) *GetSkillGroupAgentStatusDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetCurrentPage(v int32) *GetSkillGroupAgentStatusDetailsRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetDepIds(v []*int64) *GetSkillGroupAgentStatusDetailsRequest {
	s.DepIds = v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetGroupIds(v []*int64) *GetSkillGroupAgentStatusDetailsRequest {
	s.GroupIds = v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupAgentStatusDetailsRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupAgentStatusDetailsRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetSkillGroupAgentStatusDetailsShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// 技能组id列表
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetSkillGroupAgentStatusDetailsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupAgentStatusDetailsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetInstanceId(v string) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetStartDate(v int64) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetEndDate(v int64) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetPageSize(v int32) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetCurrentPage(v int32) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetDepIdsShrink(v string) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetGroupIdsShrink(v string) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetSkillGroupAgentStatusDetailsResponseBody struct {
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 接口调用是否成功
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// data
	Data *GetSkillGroupAgentStatusDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetSkillGroupAgentStatusDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupAgentStatusDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) SetRequestId(v string) *GetSkillGroupAgentStatusDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) SetMessage(v string) *GetSkillGroupAgentStatusDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) SetCode(v string) *GetSkillGroupAgentStatusDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) SetSuccess(v string) *GetSkillGroupAgentStatusDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) SetData(v *GetSkillGroupAgentStatusDetailsResponseBodyData) *GetSkillGroupAgentStatusDetailsResponseBody {
	s.Data = v
	return s
}

type GetSkillGroupAgentStatusDetailsResponseBodyData struct {
	// 当前页数
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 每页的数量
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总记录数
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s GetSkillGroupAgentStatusDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupAgentStatusDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAgentStatusDetailsResponseBodyData) SetPageNum(v int64) *GetSkillGroupAgentStatusDetailsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBodyData) SetPageSize(v int64) *GetSkillGroupAgentStatusDetailsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBodyData) SetTotalNum(v int64) *GetSkillGroupAgentStatusDetailsResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBodyData) SetRows(v string) *GetSkillGroupAgentStatusDetailsResponseBodyData {
	s.Rows = &v
	return s
}

type GetSkillGroupAgentStatusDetailsResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSkillGroupAgentStatusDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSkillGroupAgentStatusDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupAgentStatusDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAgentStatusDetailsResponse) SetHeaders(v map[string]*string) *GetSkillGroupAgentStatusDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponse) SetBody(v *GetSkillGroupAgentStatusDetailsResponseBody) *GetSkillGroupAgentStatusDetailsResponse {
	s.Body = v
	return s
}

type GetSkillGroupAndAgentStatusSummaryRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 部门id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// 技能组id列表
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetSkillGroupAndAgentStatusSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupAndAgentStatusSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetInstanceId(v string) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetStartDate(v int64) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetEndDate(v int64) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetPageSize(v int32) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetCurrentPage(v int32) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetDepIds(v []*int64) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.DepIds = v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetGroupIds(v []*int64) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.GroupIds = v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupAndAgentStatusSummaryRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetSkillGroupAndAgentStatusSummaryShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// 技能组id列表
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetSkillGroupAndAgentStatusSummaryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupAndAgentStatusSummaryShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetInstanceId(v string) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetStartDate(v int64) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetEndDate(v int64) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetPageSize(v int32) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetCurrentPage(v int32) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetDepIdsShrink(v string) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetGroupIdsShrink(v string) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupAndAgentStatusSummaryShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetSkillGroupAndAgentStatusSummaryResponseBody struct {
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 接口调用是否成功
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// data
	Data *GetSkillGroupAndAgentStatusSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetSkillGroupAndAgentStatusSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupAndAgentStatusSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) SetRequestId(v string) *GetSkillGroupAndAgentStatusSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) SetMessage(v string) *GetSkillGroupAndAgentStatusSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) SetCode(v string) *GetSkillGroupAndAgentStatusSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) SetSuccess(v string) *GetSkillGroupAndAgentStatusSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) SetData(v *GetSkillGroupAndAgentStatusSummaryResponseBodyData) *GetSkillGroupAndAgentStatusSummaryResponseBody {
	s.Data = v
	return s
}

type GetSkillGroupAndAgentStatusSummaryResponseBodyData struct {
	// 当前页数
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 每页的数量
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总记录数
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s GetSkillGroupAndAgentStatusSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupAndAgentStatusSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBodyData) SetPageNum(v int32) *GetSkillGroupAndAgentStatusSummaryResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBodyData) SetPageSize(v int32) *GetSkillGroupAndAgentStatusSummaryResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBodyData) SetTotalNum(v int32) *GetSkillGroupAndAgentStatusSummaryResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBodyData) SetRows(v string) *GetSkillGroupAndAgentStatusSummaryResponseBodyData {
	s.Rows = &v
	return s
}

type GetSkillGroupAndAgentStatusSummaryResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSkillGroupAndAgentStatusSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSkillGroupAndAgentStatusSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupAndAgentStatusSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAndAgentStatusSummaryResponse) SetHeaders(v map[string]*string) *GetSkillGroupAndAgentStatusSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponse) SetBody(v *GetSkillGroupAndAgentStatusSummaryResponseBody) *GetSkillGroupAndAgentStatusSummaryResponse {
	s.Body = v
	return s
}

type GetSkillGroupLatitudeStateRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 部门id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// 技能组id列表
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetSkillGroupLatitudeStateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupLatitudeStateRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupLatitudeStateRequest) SetInstanceId(v string) *GetSkillGroupLatitudeStateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetStartDate(v int64) *GetSkillGroupLatitudeStateRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetEndDate(v int64) *GetSkillGroupLatitudeStateRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetPageSize(v int32) *GetSkillGroupLatitudeStateRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetCurrentPage(v int32) *GetSkillGroupLatitudeStateRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetDepIds(v []*int64) *GetSkillGroupLatitudeStateRequest {
	s.DepIds = v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetGroupIds(v []*int64) *GetSkillGroupLatitudeStateRequest {
	s.GroupIds = v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupLatitudeStateRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupLatitudeStateRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupLatitudeStateRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetSkillGroupLatitudeStateShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// 技能组id列表
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetSkillGroupLatitudeStateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupLatitudeStateShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetInstanceId(v string) *GetSkillGroupLatitudeStateShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetStartDate(v int64) *GetSkillGroupLatitudeStateShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetEndDate(v int64) *GetSkillGroupLatitudeStateShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetPageSize(v int32) *GetSkillGroupLatitudeStateShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetCurrentPage(v int32) *GetSkillGroupLatitudeStateShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetDepIdsShrink(v string) *GetSkillGroupLatitudeStateShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetGroupIdsShrink(v string) *GetSkillGroupLatitudeStateShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupLatitudeStateShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupLatitudeStateShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupLatitudeStateShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetSkillGroupLatitudeStateResponseBody struct {
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 接口调用是否成功
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// data
	Data *GetSkillGroupLatitudeStateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetSkillGroupLatitudeStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupLatitudeStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupLatitudeStateResponseBody) SetRequestId(v string) *GetSkillGroupLatitudeStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBody) SetMessage(v string) *GetSkillGroupLatitudeStateResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBody) SetCode(v string) *GetSkillGroupLatitudeStateResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBody) SetSuccess(v string) *GetSkillGroupLatitudeStateResponseBody {
	s.Success = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBody) SetData(v *GetSkillGroupLatitudeStateResponseBodyData) *GetSkillGroupLatitudeStateResponseBody {
	s.Data = v
	return s
}

type GetSkillGroupLatitudeStateResponseBodyData struct {
	// 当前页数
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 每页的数量
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总共记录数
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s GetSkillGroupLatitudeStateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupLatitudeStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupLatitudeStateResponseBodyData) SetPageNum(v int32) *GetSkillGroupLatitudeStateResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBodyData) SetPageSize(v int32) *GetSkillGroupLatitudeStateResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBodyData) SetTotalNum(v int32) *GetSkillGroupLatitudeStateResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBodyData) SetRows(v string) *GetSkillGroupLatitudeStateResponseBodyData {
	s.Rows = &v
	return s
}

type GetSkillGroupLatitudeStateResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSkillGroupLatitudeStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSkillGroupLatitudeStateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupLatitudeStateResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupLatitudeStateResponse) SetHeaders(v map[string]*string) *GetSkillGroupLatitudeStateResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupLatitudeStateResponse) SetBody(v *GetSkillGroupLatitudeStateResponseBody) *GetSkillGroupLatitudeStateResponse {
	s.Body = v
	return s
}

type GetSkillGroupServiceCapabilityRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 部门id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// 技能组id列表
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetSkillGroupServiceCapabilityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupServiceCapabilityRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceCapabilityRequest) SetInstanceId(v string) *GetSkillGroupServiceCapabilityRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetStartDate(v int64) *GetSkillGroupServiceCapabilityRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetEndDate(v int64) *GetSkillGroupServiceCapabilityRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetPageSize(v int32) *GetSkillGroupServiceCapabilityRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetCurrentPage(v int32) *GetSkillGroupServiceCapabilityRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetDepIds(v []*int64) *GetSkillGroupServiceCapabilityRequest {
	s.DepIds = v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetGroupIds(v []*int64) *GetSkillGroupServiceCapabilityRequest {
	s.GroupIds = v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupServiceCapabilityRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupServiceCapabilityRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetSkillGroupServiceCapabilityShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// 技能组id列表
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetSkillGroupServiceCapabilityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupServiceCapabilityShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetInstanceId(v string) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetStartDate(v int64) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetEndDate(v int64) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetPageSize(v int32) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetCurrentPage(v int32) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetDepIdsShrink(v string) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetGroupIdsShrink(v string) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupServiceCapabilityShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetSkillGroupServiceCapabilityResponseBody struct {
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 调用接口是否成功
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// data
	Data *GetSkillGroupServiceCapabilityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetSkillGroupServiceCapabilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupServiceCapabilityResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceCapabilityResponseBody) SetRequestId(v string) *GetSkillGroupServiceCapabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBody) SetMessage(v string) *GetSkillGroupServiceCapabilityResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBody) SetCode(v string) *GetSkillGroupServiceCapabilityResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBody) SetSuccess(v string) *GetSkillGroupServiceCapabilityResponseBody {
	s.Success = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBody) SetData(v *GetSkillGroupServiceCapabilityResponseBodyData) *GetSkillGroupServiceCapabilityResponseBody {
	s.Data = v
	return s
}

type GetSkillGroupServiceCapabilityResponseBodyData struct {
	// 当前页数
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总记录数
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s GetSkillGroupServiceCapabilityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupServiceCapabilityResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceCapabilityResponseBodyData) SetPageNum(v int32) *GetSkillGroupServiceCapabilityResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBodyData) SetPageSize(v int32) *GetSkillGroupServiceCapabilityResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBodyData) SetTotalNum(v int64) *GetSkillGroupServiceCapabilityResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBodyData) SetRows(v string) *GetSkillGroupServiceCapabilityResponseBodyData {
	s.Rows = &v
	return s
}

type GetSkillGroupServiceCapabilityResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSkillGroupServiceCapabilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSkillGroupServiceCapabilityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupServiceCapabilityResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceCapabilityResponse) SetHeaders(v map[string]*string) *GetSkillGroupServiceCapabilityResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponse) SetBody(v *GetSkillGroupServiceCapabilityResponseBody) *GetSkillGroupServiceCapabilityResponse {
	s.Body = v
	return s
}

type GetSkillGroupServiceStatusRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 技能组id列表
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// 部门id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// 技能组id列表
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// 时间纬度类型
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
	// 是否根据技能组分组
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
	// 是否根据机器实例分组
	ExistRobotInstanceGrouping *bool `json:"ExistRobotInstanceGrouping,omitempty" xml:"ExistRobotInstanceGrouping,omitempty"`
	// 是否根据渠道实例分组
	ExistChannelInstanceGrouping *bool `json:"ExistChannelInstanceGrouping,omitempty" xml:"ExistChannelInstanceGrouping,omitempty"`
}

func (s GetSkillGroupServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceStatusRequest) SetInstanceId(v string) *GetSkillGroupServiceStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetStartDate(v int64) *GetSkillGroupServiceStatusRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetEndDate(v int64) *GetSkillGroupServiceStatusRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetPageSize(v int32) *GetSkillGroupServiceStatusRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetCurrentPage(v int32) *GetSkillGroupServiceStatusRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetAgentIds(v []*int64) *GetSkillGroupServiceStatusRequest {
	s.AgentIds = v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetDepIds(v []*int64) *GetSkillGroupServiceStatusRequest {
	s.DepIds = v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetGroupIds(v []*int64) *GetSkillGroupServiceStatusRequest {
	s.GroupIds = v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetTimeLatitudeType(v string) *GetSkillGroupServiceStatusRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetExistAgentGrouping(v bool) *GetSkillGroupServiceStatusRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupServiceStatusRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupServiceStatusRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetExistRobotInstanceGrouping(v bool) *GetSkillGroupServiceStatusRequest {
	s.ExistRobotInstanceGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusRequest) SetExistChannelInstanceGrouping(v bool) *GetSkillGroupServiceStatusRequest {
	s.ExistChannelInstanceGrouping = &v
	return s
}

type GetSkillGroupServiceStatusShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 技能组id列表
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// 技能组id列表
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// 时间纬度类型
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
	// 是否根据技能组分组
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
	// 是否根据机器实例分组
	ExistRobotInstanceGrouping *bool `json:"ExistRobotInstanceGrouping,omitempty" xml:"ExistRobotInstanceGrouping,omitempty"`
	// 是否根据渠道实例分组
	ExistChannelInstanceGrouping *bool `json:"ExistChannelInstanceGrouping,omitempty" xml:"ExistChannelInstanceGrouping,omitempty"`
}

func (s GetSkillGroupServiceStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupServiceStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetInstanceId(v string) *GetSkillGroupServiceStatusShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetStartDate(v int64) *GetSkillGroupServiceStatusShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetEndDate(v int64) *GetSkillGroupServiceStatusShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetPageSize(v int32) *GetSkillGroupServiceStatusShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetCurrentPage(v int32) *GetSkillGroupServiceStatusShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetAgentIdsShrink(v string) *GetSkillGroupServiceStatusShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetDepIdsShrink(v string) *GetSkillGroupServiceStatusShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetGroupIdsShrink(v string) *GetSkillGroupServiceStatusShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetTimeLatitudeType(v string) *GetSkillGroupServiceStatusShrinkRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetExistAgentGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetExistRobotInstanceGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest {
	s.ExistRobotInstanceGrouping = &v
	return s
}

func (s *GetSkillGroupServiceStatusShrinkRequest) SetExistChannelInstanceGrouping(v bool) *GetSkillGroupServiceStatusShrinkRequest {
	s.ExistChannelInstanceGrouping = &v
	return s
}

type GetSkillGroupServiceStatusResponseBody struct {
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 调用接口是否成功
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// data
	Data *GetSkillGroupServiceStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetSkillGroupServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceStatusResponseBody) SetRequestId(v string) *GetSkillGroupServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBody) SetMessage(v string) *GetSkillGroupServiceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBody) SetCode(v string) *GetSkillGroupServiceStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBody) SetSuccess(v string) *GetSkillGroupServiceStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBody) SetData(v *GetSkillGroupServiceStatusResponseBodyData) *GetSkillGroupServiceStatusResponseBody {
	s.Data = v
	return s
}

type GetSkillGroupServiceStatusResponseBodyData struct {
	// 页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总记录数
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// 当前页数
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
}

func (s GetSkillGroupServiceStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupServiceStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceStatusResponseBodyData) SetPageSize(v int32) *GetSkillGroupServiceStatusResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBodyData) SetTotalNum(v int32) *GetSkillGroupServiceStatusResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBodyData) SetRows(v string) *GetSkillGroupServiceStatusResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBodyData) SetPageNum(v int32) *GetSkillGroupServiceStatusResponseBodyData {
	s.PageNum = &v
	return s
}

type GetSkillGroupServiceStatusResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSkillGroupServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSkillGroupServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceStatusResponse) SetHeaders(v map[string]*string) *GetSkillGroupServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupServiceStatusResponse) SetBody(v *GetSkillGroupServiceStatusResponseBody) *GetSkillGroupServiceStatusResponse {
	s.Body = v
	return s
}

type GetSkillGroupStatusTotalRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// 部门id列表
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// 技能组id列表
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// 时间纬度类型
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
	// 是否根据坐席分组
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetSkillGroupStatusTotalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupStatusTotalRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupStatusTotalRequest) SetInstanceId(v string) *GetSkillGroupStatusTotalRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetStartDate(v int64) *GetSkillGroupStatusTotalRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetEndDate(v int64) *GetSkillGroupStatusTotalRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetPageSize(v int32) *GetSkillGroupStatusTotalRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetCurrentPage(v int32) *GetSkillGroupStatusTotalRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetAgentIds(v []*int64) *GetSkillGroupStatusTotalRequest {
	s.AgentIds = v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetDepIds(v []*int64) *GetSkillGroupStatusTotalRequest {
	s.DepIds = v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetGroupIds(v []*int64) *GetSkillGroupStatusTotalRequest {
	s.GroupIds = v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetTimeLatitudeType(v string) *GetSkillGroupStatusTotalRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetExistAgentGrouping(v bool) *GetSkillGroupStatusTotalRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupStatusTotalRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupStatusTotalRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupStatusTotalRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetSkillGroupStatusTotalShrinkRequest struct {
	// AICCS实例ID，在智能联络中心控制台上可以看到
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始日期时间戳（毫秒）
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束日期时间戳（毫秒）
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 每页大小（默认为10)
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 当前页（默认为1）
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 坐席id列表
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// 部门id列表
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// 技能组id列表
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// 时间纬度类型
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
	// 是否根据坐席分组
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// 是否根据部门分组
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// 是否根据技能组分组
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
}

func (s GetSkillGroupStatusTotalShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupStatusTotalShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetInstanceId(v string) *GetSkillGroupStatusTotalShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetStartDate(v int64) *GetSkillGroupStatusTotalShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetEndDate(v int64) *GetSkillGroupStatusTotalShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetPageSize(v int32) *GetSkillGroupStatusTotalShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetCurrentPage(v int32) *GetSkillGroupStatusTotalShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetAgentIdsShrink(v string) *GetSkillGroupStatusTotalShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetDepIdsShrink(v string) *GetSkillGroupStatusTotalShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetGroupIdsShrink(v string) *GetSkillGroupStatusTotalShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetTimeLatitudeType(v string) *GetSkillGroupStatusTotalShrinkRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetExistAgentGrouping(v bool) *GetSkillGroupStatusTotalShrinkRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupStatusTotalShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupStatusTotalShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupStatusTotalShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

type GetSkillGroupStatusTotalResponseBody struct {
	// 请求ID，用于跟踪错误原因
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误编码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 接口调用是否成功
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// data
	Data *GetSkillGroupStatusTotalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetSkillGroupStatusTotalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupStatusTotalResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupStatusTotalResponseBody) SetRequestId(v string) *GetSkillGroupStatusTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBody) SetMessage(v string) *GetSkillGroupStatusTotalResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBody) SetCode(v string) *GetSkillGroupStatusTotalResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBody) SetSuccess(v string) *GetSkillGroupStatusTotalResponseBody {
	s.Success = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBody) SetData(v *GetSkillGroupStatusTotalResponseBodyData) *GetSkillGroupStatusTotalResponseBody {
	s.Data = v
	return s
}

type GetSkillGroupStatusTotalResponseBodyData struct {
	// 当前页数
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 页大小
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总记录数
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 信息为list<map>类型的json字符串
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s GetSkillGroupStatusTotalResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupStatusTotalResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupStatusTotalResponseBodyData) SetPageNum(v int64) *GetSkillGroupStatusTotalResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBodyData) SetPageSize(v int64) *GetSkillGroupStatusTotalResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBodyData) SetTotalNum(v int64) *GetSkillGroupStatusTotalResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBodyData) SetRows(v string) *GetSkillGroupStatusTotalResponseBodyData {
	s.Rows = &v
	return s
}

type GetSkillGroupStatusTotalResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSkillGroupStatusTotalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSkillGroupStatusTotalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupStatusTotalResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupStatusTotalResponse) SetHeaders(v map[string]*string) *GetSkillGroupStatusTotalResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupStatusTotalResponse) SetBody(v *GetSkillGroupStatusTotalResponseBody) *GetSkillGroupStatusTotalResponse {
	s.Body = v
	return s
}

type HangupCallRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
}

func (s HangupCallRequest) String() string {
	return tea.Prettify(s)
}

func (s HangupCallRequest) GoString() string {
	return s.String()
}

func (s *HangupCallRequest) SetClientToken(v string) *HangupCallRequest {
	s.ClientToken = &v
	return s
}

func (s *HangupCallRequest) SetInstanceId(v string) *HangupCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HangupCallRequest) SetAccountName(v string) *HangupCallRequest {
	s.AccountName = &v
	return s
}

func (s *HangupCallRequest) SetCallId(v string) *HangupCallRequest {
	s.CallId = &v
	return s
}

func (s *HangupCallRequest) SetJobId(v string) *HangupCallRequest {
	s.JobId = &v
	return s
}

func (s *HangupCallRequest) SetConnectionId(v string) *HangupCallRequest {
	s.ConnectionId = &v
	return s
}

type HangupCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HangupCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HangupCallResponseBody) GoString() string {
	return s.String()
}

func (s *HangupCallResponseBody) SetCode(v string) *HangupCallResponseBody {
	s.Code = &v
	return s
}

func (s *HangupCallResponseBody) SetMessage(v string) *HangupCallResponseBody {
	s.Message = &v
	return s
}

func (s *HangupCallResponseBody) SetRequestId(v string) *HangupCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HangupCallResponseBody) SetSuccess(v bool) *HangupCallResponseBody {
	s.Success = &v
	return s
}

type HangupCallResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HangupCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HangupCallResponse) String() string {
	return tea.Prettify(s)
}

func (s HangupCallResponse) GoString() string {
	return s.String()
}

func (s *HangupCallResponse) SetHeaders(v map[string]*string) *HangupCallResponse {
	s.Headers = v
	return s
}

func (s *HangupCallResponse) SetBody(v *HangupCallResponseBody) *HangupCallResponse {
	s.Body = v
	return s
}

type HangUpDoubleCallRequest struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 会话ID
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
}

func (s HangUpDoubleCallRequest) String() string {
	return tea.Prettify(s)
}

func (s HangUpDoubleCallRequest) GoString() string {
	return s.String()
}

func (s *HangUpDoubleCallRequest) SetInstanceId(v string) *HangUpDoubleCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HangUpDoubleCallRequest) SetAcid(v string) *HangUpDoubleCallRequest {
	s.Acid = &v
	return s
}

type HangUpDoubleCallResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否调用成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s HangUpDoubleCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HangUpDoubleCallResponseBody) GoString() string {
	return s.String()
}

func (s *HangUpDoubleCallResponseBody) SetRequestId(v string) *HangUpDoubleCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HangUpDoubleCallResponseBody) SetSuccess(v bool) *HangUpDoubleCallResponseBody {
	s.Success = &v
	return s
}

func (s *HangUpDoubleCallResponseBody) SetCode(v string) *HangUpDoubleCallResponseBody {
	s.Code = &v
	return s
}

func (s *HangUpDoubleCallResponseBody) SetMessage(v string) *HangUpDoubleCallResponseBody {
	s.Message = &v
	return s
}

type HangUpDoubleCallResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HangUpDoubleCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HangUpDoubleCallResponse) String() string {
	return tea.Prettify(s)
}

func (s HangUpDoubleCallResponse) GoString() string {
	return s.String()
}

func (s *HangUpDoubleCallResponse) SetHeaders(v map[string]*string) *HangUpDoubleCallResponse {
	s.Headers = v
	return s
}

func (s *HangUpDoubleCallResponse) SetBody(v *HangUpDoubleCallResponseBody) *HangUpDoubleCallResponse {
	s.Body = v
	return s
}

type HangupThirdCallRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
}

func (s HangupThirdCallRequest) String() string {
	return tea.Prettify(s)
}

func (s HangupThirdCallRequest) GoString() string {
	return s.String()
}

func (s *HangupThirdCallRequest) SetClientToken(v string) *HangupThirdCallRequest {
	s.ClientToken = &v
	return s
}

func (s *HangupThirdCallRequest) SetInstanceId(v string) *HangupThirdCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HangupThirdCallRequest) SetAccountName(v string) *HangupThirdCallRequest {
	s.AccountName = &v
	return s
}

func (s *HangupThirdCallRequest) SetCallId(v string) *HangupThirdCallRequest {
	s.CallId = &v
	return s
}

func (s *HangupThirdCallRequest) SetJobId(v string) *HangupThirdCallRequest {
	s.JobId = &v
	return s
}

func (s *HangupThirdCallRequest) SetConnectionId(v string) *HangupThirdCallRequest {
	s.ConnectionId = &v
	return s
}

type HangupThirdCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HangupThirdCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HangupThirdCallResponseBody) GoString() string {
	return s.String()
}

func (s *HangupThirdCallResponseBody) SetCode(v string) *HangupThirdCallResponseBody {
	s.Code = &v
	return s
}

func (s *HangupThirdCallResponseBody) SetMessage(v string) *HangupThirdCallResponseBody {
	s.Message = &v
	return s
}

func (s *HangupThirdCallResponseBody) SetRequestId(v string) *HangupThirdCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HangupThirdCallResponseBody) SetSuccess(v bool) *HangupThirdCallResponseBody {
	s.Success = &v
	return s
}

type HangupThirdCallResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HangupThirdCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HangupThirdCallResponse) String() string {
	return tea.Prettify(s)
}

func (s HangupThirdCallResponse) GoString() string {
	return s.String()
}

func (s *HangupThirdCallResponse) SetHeaders(v map[string]*string) *HangupThirdCallResponse {
	s.Headers = v
	return s
}

func (s *HangupThirdCallResponse) SetBody(v *HangupThirdCallResponseBody) *HangupThirdCallResponse {
	s.Body = v
	return s
}

type HoldCallRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
}

func (s HoldCallRequest) String() string {
	return tea.Prettify(s)
}

func (s HoldCallRequest) GoString() string {
	return s.String()
}

func (s *HoldCallRequest) SetClientToken(v string) *HoldCallRequest {
	s.ClientToken = &v
	return s
}

func (s *HoldCallRequest) SetInstanceId(v string) *HoldCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HoldCallRequest) SetAccountName(v string) *HoldCallRequest {
	s.AccountName = &v
	return s
}

func (s *HoldCallRequest) SetCallId(v string) *HoldCallRequest {
	s.CallId = &v
	return s
}

func (s *HoldCallRequest) SetJobId(v string) *HoldCallRequest {
	s.JobId = &v
	return s
}

func (s *HoldCallRequest) SetConnectionId(v string) *HoldCallRequest {
	s.ConnectionId = &v
	return s
}

type HoldCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HoldCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HoldCallResponseBody) GoString() string {
	return s.String()
}

func (s *HoldCallResponseBody) SetCode(v string) *HoldCallResponseBody {
	s.Code = &v
	return s
}

func (s *HoldCallResponseBody) SetMessage(v string) *HoldCallResponseBody {
	s.Message = &v
	return s
}

func (s *HoldCallResponseBody) SetRequestId(v string) *HoldCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HoldCallResponseBody) SetSuccess(v bool) *HoldCallResponseBody {
	s.Success = &v
	return s
}

type HoldCallResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HoldCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HoldCallResponse) String() string {
	return tea.Prettify(s)
}

func (s HoldCallResponse) GoString() string {
	return s.String()
}

func (s *HoldCallResponse) SetHeaders(v map[string]*string) *HoldCallResponse {
	s.Headers = v
	return s
}

func (s *HoldCallResponse) SetBody(v *HoldCallResponseBody) *HoldCallResponse {
	s.Body = v
	return s
}

type HotlineSessionQueryRequest struct {
	InstanceId        *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Acid              *string   `json:"Acid,omitempty" xml:"Acid,omitempty"`
	CallType          *int32    `json:"CallType,omitempty" xml:"CallType,omitempty"`
	CalledNumber      *string   `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber     *string   `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	GroupId           *int64    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName         *string   `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	MemberId          *string   `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName        *string   `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	QueryEndTime      *int64    `json:"QueryEndTime,omitempty" xml:"QueryEndTime,omitempty"`
	QueryStartTime    *int64    `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	RequestId         *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServicerName      *string   `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	ServicerId        *string   `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	Params            *string   `json:"Params,omitempty" xml:"Params,omitempty"`
	PageSize          *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo            *int32    `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	CallResult        *string   `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	Id                *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	AcidList          []*string `json:"AcidList,omitempty" xml:"AcidList,omitempty" type:"Repeated"`
	CallTypeList      []*int32  `json:"CallTypeList,omitempty" xml:"CallTypeList,omitempty" type:"Repeated"`
	GroupIdList       []*int64  `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty" type:"Repeated"`
	CallingNumberList []*string `json:"CallingNumberList,omitempty" xml:"CallingNumberList,omitempty" type:"Repeated"`
	CalledNumberList  []*string `json:"CalledNumberList,omitempty" xml:"CalledNumberList,omitempty" type:"Repeated"`
	MemberIdList      []*string `json:"MemberIdList,omitempty" xml:"MemberIdList,omitempty" type:"Repeated"`
	ServicerIdList    []*string `json:"ServicerIdList,omitempty" xml:"ServicerIdList,omitempty" type:"Repeated"`
	CallResultList    []*string `json:"CallResultList,omitempty" xml:"CallResultList,omitempty" type:"Repeated"`
}

func (s HotlineSessionQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s HotlineSessionQueryRequest) GoString() string {
	return s.String()
}

func (s *HotlineSessionQueryRequest) SetInstanceId(v string) *HotlineSessionQueryRequest {
	s.InstanceId = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetAcid(v string) *HotlineSessionQueryRequest {
	s.Acid = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetCallType(v int32) *HotlineSessionQueryRequest {
	s.CallType = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetCalledNumber(v string) *HotlineSessionQueryRequest {
	s.CalledNumber = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetCallingNumber(v string) *HotlineSessionQueryRequest {
	s.CallingNumber = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetGroupId(v int64) *HotlineSessionQueryRequest {
	s.GroupId = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetGroupName(v string) *HotlineSessionQueryRequest {
	s.GroupName = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetMemberId(v string) *HotlineSessionQueryRequest {
	s.MemberId = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetMemberName(v string) *HotlineSessionQueryRequest {
	s.MemberName = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetQueryEndTime(v int64) *HotlineSessionQueryRequest {
	s.QueryEndTime = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetQueryStartTime(v int64) *HotlineSessionQueryRequest {
	s.QueryStartTime = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetRequestId(v string) *HotlineSessionQueryRequest {
	s.RequestId = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetServicerName(v string) *HotlineSessionQueryRequest {
	s.ServicerName = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetServicerId(v string) *HotlineSessionQueryRequest {
	s.ServicerId = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetParams(v string) *HotlineSessionQueryRequest {
	s.Params = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetPageSize(v int32) *HotlineSessionQueryRequest {
	s.PageSize = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetPageNo(v int32) *HotlineSessionQueryRequest {
	s.PageNo = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetCallResult(v string) *HotlineSessionQueryRequest {
	s.CallResult = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetId(v string) *HotlineSessionQueryRequest {
	s.Id = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetAcidList(v []*string) *HotlineSessionQueryRequest {
	s.AcidList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetCallTypeList(v []*int32) *HotlineSessionQueryRequest {
	s.CallTypeList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetGroupIdList(v []*int64) *HotlineSessionQueryRequest {
	s.GroupIdList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetCallingNumberList(v []*string) *HotlineSessionQueryRequest {
	s.CallingNumberList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetCalledNumberList(v []*string) *HotlineSessionQueryRequest {
	s.CalledNumberList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetMemberIdList(v []*string) *HotlineSessionQueryRequest {
	s.MemberIdList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetServicerIdList(v []*string) *HotlineSessionQueryRequest {
	s.ServicerIdList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetCallResultList(v []*string) *HotlineSessionQueryRequest {
	s.CallResultList = v
	return s
}

type HotlineSessionQueryResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *HotlineSessionQueryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HotlineSessionQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HotlineSessionQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HotlineSessionQueryResponseBody) SetMessage(v string) *HotlineSessionQueryResponseBody {
	s.Message = &v
	return s
}

func (s *HotlineSessionQueryResponseBody) SetRequestId(v string) *HotlineSessionQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotlineSessionQueryResponseBody) SetData(v *HotlineSessionQueryResponseBodyData) *HotlineSessionQueryResponseBody {
	s.Data = v
	return s
}

func (s *HotlineSessionQueryResponseBody) SetCode(v string) *HotlineSessionQueryResponseBody {
	s.Code = &v
	return s
}

func (s *HotlineSessionQueryResponseBody) SetSuccess(v bool) *HotlineSessionQueryResponseBody {
	s.Success = &v
	return s
}

type HotlineSessionQueryResponseBodyData struct {
	CallDetailRecord []*HotlineSessionQueryResponseBodyDataCallDetailRecord `json:"CallDetailRecord,omitempty" xml:"CallDetailRecord,omitempty" type:"Repeated"`
	PageSize         *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber       *int32                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount       *int32                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s HotlineSessionQueryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s HotlineSessionQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *HotlineSessionQueryResponseBodyData) SetCallDetailRecord(v []*HotlineSessionQueryResponseBodyDataCallDetailRecord) *HotlineSessionQueryResponseBodyData {
	s.CallDetailRecord = v
	return s
}

func (s *HotlineSessionQueryResponseBodyData) SetPageSize(v int32) *HotlineSessionQueryResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyData) SetPageNumber(v int32) *HotlineSessionQueryResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyData) SetTotalCount(v int32) *HotlineSessionQueryResponseBodyData {
	s.TotalCount = &v
	return s
}

type HotlineSessionQueryResponseBodyDataCallDetailRecord struct {
	CallResult            *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	TrunkCall             *string `json:"TrunkCall,omitempty" xml:"TrunkCall,omitempty"`
	ServicerName          *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	OutQueueTime          *string `json:"OutQueueTime,omitempty" xml:"OutQueueTime,omitempty"`
	CallContinueTime      *int32  `json:"CallContinueTime,omitempty" xml:"CallContinueTime,omitempty"`
	CreateTime            *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PickUpTime            *string `json:"PickUpTime,omitempty" xml:"PickUpTime,omitempty"`
	RingContinueTime      *int32  `json:"RingContinueTime,omitempty" xml:"RingContinueTime,omitempty"`
	CalledNumber          *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	ServicerId            *string `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	HangUpTime            *string `json:"HangUpTime,omitempty" xml:"HangUpTime,omitempty"`
	EvaluationLevel       *int32  `json:"EvaluationLevel,omitempty" xml:"EvaluationLevel,omitempty"`
	PassiveTransferId     *string `json:"PassiveTransferId,omitempty" xml:"PassiveTransferId,omitempty"`
	ActiveTransferId      *string `json:"ActiveTransferId,omitempty" xml:"ActiveTransferId,omitempty"`
	HangUpRole            *string `json:"HangUpRole,omitempty" xml:"HangUpRole,omitempty"`
	PassiveTransferIdType *string `json:"PassiveTransferIdType,omitempty" xml:"PassiveTransferIdType,omitempty"`
	MemberName            *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	EvaluationScore       *int32  `json:"EvaluationScore,omitempty" xml:"EvaluationScore,omitempty"`
	Acid                  *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	RingStartTime         *string `json:"RingStartTime,omitempty" xml:"RingStartTime,omitempty"`
	CallType              *int32  `json:"CallType,omitempty" xml:"CallType,omitempty"`
	GroupName             *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupId               *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RingEndTime           *string `json:"RingEndTime,omitempty" xml:"RingEndTime,omitempty"`
	CallingNumber         *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	InQueueTime           *string `json:"InQueueTime,omitempty" xml:"InQueueTime,omitempty"`
	MemberId              *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Id                    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	QueueUpContinueTime   *int32  `json:"QueueUpContinueTime,omitempty" xml:"QueueUpContinueTime,omitempty"`
}

func (s HotlineSessionQueryResponseBodyDataCallDetailRecord) String() string {
	return tea.Prettify(s)
}

func (s HotlineSessionQueryResponseBodyDataCallDetailRecord) GoString() string {
	return s.String()
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetCallResult(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.CallResult = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetTrunkCall(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.TrunkCall = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetServicerName(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.ServicerName = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetOutQueueTime(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.OutQueueTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetCallContinueTime(v int32) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.CallContinueTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetCreateTime(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.CreateTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetPickUpTime(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.PickUpTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetRingContinueTime(v int32) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.RingContinueTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetCalledNumber(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.CalledNumber = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetServicerId(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.ServicerId = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetHangUpTime(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.HangUpTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetEvaluationLevel(v int32) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.EvaluationLevel = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetPassiveTransferId(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.PassiveTransferId = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetActiveTransferId(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.ActiveTransferId = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetHangUpRole(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.HangUpRole = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetPassiveTransferIdType(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.PassiveTransferIdType = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetMemberName(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.MemberName = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetEvaluationScore(v int32) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.EvaluationScore = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetAcid(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.Acid = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetRingStartTime(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.RingStartTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetCallType(v int32) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.CallType = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetGroupName(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.GroupName = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetGroupId(v int64) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.GroupId = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetRingEndTime(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.RingEndTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetCallingNumber(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.CallingNumber = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetInQueueTime(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.InQueueTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetMemberId(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.MemberId = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetId(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.Id = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetQueueUpContinueTime(v int32) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.QueueUpContinueTime = &v
	return s
}

type HotlineSessionQueryResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HotlineSessionQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HotlineSessionQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s HotlineSessionQueryResponse) GoString() string {
	return s.String()
}

func (s *HotlineSessionQueryResponse) SetHeaders(v map[string]*string) *HotlineSessionQueryResponse {
	s.Headers = v
	return s
}

func (s *HotlineSessionQueryResponse) SetBody(v *HotlineSessionQueryResponseBody) *HotlineSessionQueryResponse {
	s.Body = v
	return s
}

type InsertTaskDetailRequest struct {
	OutboundTaskId *int64  `json:"OutboundTaskId,omitempty" xml:"OutboundTaskId,omitempty"`
	CallInfos      *string `json:"CallInfos,omitempty" xml:"CallInfos,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s InsertTaskDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *InsertTaskDetailRequest) SetOutboundTaskId(v int64) *InsertTaskDetailRequest {
	s.OutboundTaskId = &v
	return s
}

func (s *InsertTaskDetailRequest) SetCallInfos(v string) *InsertTaskDetailRequest {
	s.CallInfos = &v
	return s
}

func (s *InsertTaskDetailRequest) SetInstanceId(v string) *InsertTaskDetailRequest {
	s.InstanceId = &v
	return s
}

type InsertTaskDetailResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertTaskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *InsertTaskDetailResponseBody) SetCode(v string) *InsertTaskDetailResponseBody {
	s.Code = &v
	return s
}

func (s *InsertTaskDetailResponseBody) SetMessage(v string) *InsertTaskDetailResponseBody {
	s.Message = &v
	return s
}

func (s *InsertTaskDetailResponseBody) SetData(v string) *InsertTaskDetailResponseBody {
	s.Data = &v
	return s
}

func (s *InsertTaskDetailResponseBody) SetRequestId(v string) *InsertTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertTaskDetailResponseBody) SetSuccess(v bool) *InsertTaskDetailResponseBody {
	s.Success = &v
	return s
}

type InsertTaskDetailResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertTaskDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *InsertTaskDetailResponse) SetHeaders(v map[string]*string) *InsertTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *InsertTaskDetailResponse) SetBody(v *InsertTaskDetailResponseBody) *InsertTaskDetailResponse {
	s.Body = v
	return s
}

type JoinThirdCallRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId           *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId     *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
}

func (s JoinThirdCallRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinThirdCallRequest) GoString() string {
	return s.String()
}

func (s *JoinThirdCallRequest) SetClientToken(v string) *JoinThirdCallRequest {
	s.ClientToken = &v
	return s
}

func (s *JoinThirdCallRequest) SetInstanceId(v string) *JoinThirdCallRequest {
	s.InstanceId = &v
	return s
}

func (s *JoinThirdCallRequest) SetAccountName(v string) *JoinThirdCallRequest {
	s.AccountName = &v
	return s
}

func (s *JoinThirdCallRequest) SetCallId(v string) *JoinThirdCallRequest {
	s.CallId = &v
	return s
}

func (s *JoinThirdCallRequest) SetJobId(v string) *JoinThirdCallRequest {
	s.JobId = &v
	return s
}

func (s *JoinThirdCallRequest) SetConnectionId(v string) *JoinThirdCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *JoinThirdCallRequest) SetHoldConnectionId(v string) *JoinThirdCallRequest {
	s.HoldConnectionId = &v
	return s
}

type JoinThirdCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s JoinThirdCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinThirdCallResponseBody) GoString() string {
	return s.String()
}

func (s *JoinThirdCallResponseBody) SetCode(v string) *JoinThirdCallResponseBody {
	s.Code = &v
	return s
}

func (s *JoinThirdCallResponseBody) SetMessage(v string) *JoinThirdCallResponseBody {
	s.Message = &v
	return s
}

func (s *JoinThirdCallResponseBody) SetRequestId(v string) *JoinThirdCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinThirdCallResponseBody) SetSuccess(v bool) *JoinThirdCallResponseBody {
	s.Success = &v
	return s
}

type JoinThirdCallResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *JoinThirdCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinThirdCallResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinThirdCallResponse) GoString() string {
	return s.String()
}

func (s *JoinThirdCallResponse) SetHeaders(v map[string]*string) *JoinThirdCallResponse {
	s.Headers = v
	return s
}

func (s *JoinThirdCallResponse) SetBody(v *JoinThirdCallResponseBody) *JoinThirdCallResponse {
	s.Body = v
	return s
}

type ListAgentBySkillGroupIdRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ListAgentBySkillGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAgentBySkillGroupIdRequest) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdRequest) SetClientToken(v string) *ListAgentBySkillGroupIdRequest {
	s.ClientToken = &v
	return s
}

func (s *ListAgentBySkillGroupIdRequest) SetInstanceId(v string) *ListAgentBySkillGroupIdRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAgentBySkillGroupIdRequest) SetSkillGroupId(v int64) *ListAgentBySkillGroupIdRequest {
	s.SkillGroupId = &v
	return s
}

type ListAgentBySkillGroupIdResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      []*ListAgentBySkillGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListAgentBySkillGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAgentBySkillGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdResponseBody) SetCode(v string) *ListAgentBySkillGroupIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetMessage(v string) *ListAgentBySkillGroupIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetRequestId(v string) *ListAgentBySkillGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetSuccess(v bool) *ListAgentBySkillGroupIdResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetData(v []*ListAgentBySkillGroupIdResponseBodyData) *ListAgentBySkillGroupIdResponseBody {
	s.Data = v
	return s
}

type ListAgentBySkillGroupIdResponseBodyData struct {
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	AgentId     *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	TenantId    *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListAgentBySkillGroupIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAgentBySkillGroupIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetStatus(v int32) *ListAgentBySkillGroupIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetDisplayName(v string) *ListAgentBySkillGroupIdResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetAgentId(v int64) *ListAgentBySkillGroupIdResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetAccountName(v string) *ListAgentBySkillGroupIdResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetTenantId(v int64) *ListAgentBySkillGroupIdResponseBodyData {
	s.TenantId = &v
	return s
}

type ListAgentBySkillGroupIdResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAgentBySkillGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAgentBySkillGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAgentBySkillGroupIdResponse) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdResponse) SetHeaders(v map[string]*string) *ListAgentBySkillGroupIdResponse {
	s.Headers = v
	return s
}

func (s *ListAgentBySkillGroupIdResponse) SetBody(v *ListAgentBySkillGroupIdResponseBody) *ListAgentBySkillGroupIdResponse {
	s.Body = v
	return s
}

type ListAiccsRobotRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RobotName            *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
}

func (s ListAiccsRobotRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAiccsRobotRequest) GoString() string {
	return s.String()
}

func (s *ListAiccsRobotRequest) SetOwnerId(v int64) *ListAiccsRobotRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAiccsRobotRequest) SetResourceOwnerAccount(v string) *ListAiccsRobotRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAiccsRobotRequest) SetResourceOwnerId(v int64) *ListAiccsRobotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAiccsRobotRequest) SetRobotName(v string) *ListAiccsRobotRequest {
	s.RobotName = &v
	return s
}

type ListAiccsRobotResponseBody struct {
	// Id of the request
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*ListAiccsRobotResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAiccsRobotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAiccsRobotResponseBody) GoString() string {
	return s.String()
}

func (s *ListAiccsRobotResponseBody) SetRequestId(v string) *ListAiccsRobotResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAiccsRobotResponseBody) SetData(v []*ListAiccsRobotResponseBodyData) *ListAiccsRobotResponseBody {
	s.Data = v
	return s
}

func (s *ListAiccsRobotResponseBody) SetCode(v string) *ListAiccsRobotResponseBody {
	s.Code = &v
	return s
}

func (s *ListAiccsRobotResponseBody) SetMessage(v string) *ListAiccsRobotResponseBody {
	s.Message = &v
	return s
}

func (s *ListAiccsRobotResponseBody) SetSuccess(v bool) *ListAiccsRobotResponseBody {
	s.Success = &v
	return s
}

type ListAiccsRobotResponseBodyData struct {
	RobotType    *string `json:"RobotType,omitempty" xml:"RobotType,omitempty"`
	AtSence      *string `json:"AtSence,omitempty" xml:"AtSence,omitempty"`
	AtProfession *string `json:"AtProfession,omitempty" xml:"AtProfession,omitempty"`
	RobotName    *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListAiccsRobotResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAiccsRobotResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAiccsRobotResponseBodyData) SetRobotType(v string) *ListAiccsRobotResponseBodyData {
	s.RobotType = &v
	return s
}

func (s *ListAiccsRobotResponseBodyData) SetAtSence(v string) *ListAiccsRobotResponseBodyData {
	s.AtSence = &v
	return s
}

func (s *ListAiccsRobotResponseBodyData) SetAtProfession(v string) *ListAiccsRobotResponseBodyData {
	s.AtProfession = &v
	return s
}

func (s *ListAiccsRobotResponseBodyData) SetRobotName(v string) *ListAiccsRobotResponseBodyData {
	s.RobotName = &v
	return s
}

func (s *ListAiccsRobotResponseBodyData) SetId(v int64) *ListAiccsRobotResponseBodyData {
	s.Id = &v
	return s
}

type ListAiccsRobotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAiccsRobotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAiccsRobotResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAiccsRobotResponse) GoString() string {
	return s.String()
}

func (s *ListAiccsRobotResponse) SetHeaders(v map[string]*string) *ListAiccsRobotResponse {
	s.Headers = v
	return s
}

func (s *ListAiccsRobotResponse) SetBody(v *ListAiccsRobotResponseBody) *ListAiccsRobotResponse {
	s.Body = v
	return s
}

type ListChatRecordDetailRequest struct {
	// clientToken
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 在线挂断的时间范围
	CloseTimeEnd *int64 `json:"CloseTimeEnd,omitempty" xml:"CloseTimeEnd,omitempty"`
	// 当前页
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 每页数据量
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 在线挂断的时间范围
	CloseTimeStart *int64 `json:"CloseTimeStart,omitempty" xml:"CloseTimeStart,omitempty"`
}

func (s ListChatRecordDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChatRecordDetailRequest) GoString() string {
	return s.String()
}

func (s *ListChatRecordDetailRequest) SetClientToken(v string) *ListChatRecordDetailRequest {
	s.ClientToken = &v
	return s
}

func (s *ListChatRecordDetailRequest) SetInstanceId(v string) *ListChatRecordDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *ListChatRecordDetailRequest) SetCloseTimeEnd(v int64) *ListChatRecordDetailRequest {
	s.CloseTimeEnd = &v
	return s
}

func (s *ListChatRecordDetailRequest) SetCurrentPage(v int32) *ListChatRecordDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListChatRecordDetailRequest) SetPageSize(v int32) *ListChatRecordDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListChatRecordDetailRequest) SetCloseTimeStart(v int64) *ListChatRecordDetailRequest {
	s.CloseTimeStart = &v
	return s
}

type ListChatRecordDetailResponseBody struct {
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// data
	ResultData *ListChatRecordDetailResponseBodyResultData `json:"ResultData,omitempty" xml:"ResultData,omitempty" type:"Struct"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListChatRecordDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChatRecordDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatRecordDetailResponseBody) SetMessage(v string) *ListChatRecordDetailResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatRecordDetailResponseBody) SetRequestId(v string) *ListChatRecordDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatRecordDetailResponseBody) SetHttpStatusCode(v int32) *ListChatRecordDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListChatRecordDetailResponseBody) SetResultData(v *ListChatRecordDetailResponseBodyResultData) *ListChatRecordDetailResponseBody {
	s.ResultData = v
	return s
}

func (s *ListChatRecordDetailResponseBody) SetCode(v string) *ListChatRecordDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatRecordDetailResponseBody) SetSuccess(v bool) *ListChatRecordDetailResponseBody {
	s.Success = &v
	return s
}

type ListChatRecordDetailResponseBodyResultData struct {
	CurrentPage  *int64                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	TotalResults *int64                                            `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
	TotalPage    *int64                                            `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	OnePageSize  *int64                                            `json:"OnePageSize,omitempty" xml:"OnePageSize,omitempty"`
	Data         []*ListChatRecordDetailResponseBodyResultDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListChatRecordDetailResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s ListChatRecordDetailResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListChatRecordDetailResponseBodyResultData) SetCurrentPage(v int64) *ListChatRecordDetailResponseBodyResultData {
	s.CurrentPage = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultData) SetTotalResults(v int64) *ListChatRecordDetailResponseBodyResultData {
	s.TotalResults = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultData) SetTotalPage(v int64) *ListChatRecordDetailResponseBodyResultData {
	s.TotalPage = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultData) SetOnePageSize(v int64) *ListChatRecordDetailResponseBodyResultData {
	s.OnePageSize = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultData) SetData(v []*ListChatRecordDetailResponseBodyResultDataData) *ListChatRecordDetailResponseBodyResultData {
	s.Data = v
	return s
}

type ListChatRecordDetailResponseBodyResultDataData struct {
	ServicerName *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	// 在线开始时间
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 在线结束时间
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 在线会话详细信息
	MessageList []*ListChatRecordDetailResponseBodyResultDataDataMessageList `json:"MessageList,omitempty" xml:"MessageList,omitempty" type:"Repeated"`
}

func (s ListChatRecordDetailResponseBodyResultDataData) String() string {
	return tea.Prettify(s)
}

func (s ListChatRecordDetailResponseBodyResultDataData) GoString() string {
	return s.String()
}

func (s *ListChatRecordDetailResponseBodyResultDataData) SetServicerName(v string) *ListChatRecordDetailResponseBodyResultDataData {
	s.ServicerName = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataData) SetStartTime(v int64) *ListChatRecordDetailResponseBodyResultDataData {
	s.StartTime = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataData) SetEndTime(v int64) *ListChatRecordDetailResponseBodyResultDataData {
	s.EndTime = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataData) SetMessageList(v []*ListChatRecordDetailResponseBodyResultDataDataMessageList) *ListChatRecordDetailResponseBodyResultDataData {
	s.MessageList = v
	return s
}

type ListChatRecordDetailResponseBodyResultDataDataMessageList struct {
	SenderName *string `json:"SenderName,omitempty" xml:"SenderName,omitempty"`
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	SenderType *int64  `json:"SenderType,omitempty" xml:"SenderType,omitempty"`
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MsgType    *string `json:"MsgType,omitempty" xml:"MsgType,omitempty"`
}

func (s ListChatRecordDetailResponseBodyResultDataDataMessageList) String() string {
	return tea.Prettify(s)
}

func (s ListChatRecordDetailResponseBodyResultDataDataMessageList) GoString() string {
	return s.String()
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) SetSenderName(v string) *ListChatRecordDetailResponseBodyResultDataDataMessageList {
	s.SenderName = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) SetContent(v string) *ListChatRecordDetailResponseBodyResultDataDataMessageList {
	s.Content = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) SetSenderType(v int64) *ListChatRecordDetailResponseBodyResultDataDataMessageList {
	s.SenderType = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) SetCreateTime(v int64) *ListChatRecordDetailResponseBodyResultDataDataMessageList {
	s.CreateTime = &v
	return s
}

func (s *ListChatRecordDetailResponseBodyResultDataDataMessageList) SetMsgType(v string) *ListChatRecordDetailResponseBodyResultDataDataMessageList {
	s.MsgType = &v
	return s
}

type ListChatRecordDetailResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListChatRecordDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListChatRecordDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChatRecordDetailResponse) GoString() string {
	return s.String()
}

func (s *ListChatRecordDetailResponse) SetHeaders(v map[string]*string) *ListChatRecordDetailResponse {
	s.Headers = v
	return s
}

func (s *ListChatRecordDetailResponse) SetBody(v *ListChatRecordDetailResponseBody) *ListChatRecordDetailResponse {
	s.Body = v
	return s
}

type ListDialogRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Called               *string `json:"Called,omitempty" xml:"Called,omitempty"`
}

func (s ListDialogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDialogRequest) GoString() string {
	return s.String()
}

func (s *ListDialogRequest) SetOwnerId(v int64) *ListDialogRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDialogRequest) SetResourceOwnerAccount(v string) *ListDialogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListDialogRequest) SetResourceOwnerId(v int64) *ListDialogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListDialogRequest) SetTaskId(v int64) *ListDialogRequest {
	s.TaskId = &v
	return s
}

func (s *ListDialogRequest) SetCalled(v string) *ListDialogRequest {
	s.Called = &v
	return s
}

type ListDialogResponseBody struct {
	// Id of the request
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*ListDialogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDialogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDialogResponseBody) GoString() string {
	return s.String()
}

func (s *ListDialogResponseBody) SetRequestId(v string) *ListDialogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDialogResponseBody) SetData(v []*ListDialogResponseBodyData) *ListDialogResponseBody {
	s.Data = v
	return s
}

func (s *ListDialogResponseBody) SetCode(v string) *ListDialogResponseBody {
	s.Code = &v
	return s
}

func (s *ListDialogResponseBody) SetMessage(v string) *ListDialogResponseBody {
	s.Message = &v
	return s
}

func (s *ListDialogResponseBody) SetSuccess(v bool) *ListDialogResponseBody {
	s.Success = &v
	return s
}

type ListDialogResponseBodyData struct {
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Time     *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Tag      *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListDialogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDialogResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDialogResponseBodyData) SetRole(v string) *ListDialogResponseBodyData {
	s.Role = &v
	return s
}

func (s *ListDialogResponseBodyData) SetContent(v string) *ListDialogResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListDialogResponseBodyData) SetNodeType(v string) *ListDialogResponseBodyData {
	s.NodeType = &v
	return s
}

func (s *ListDialogResponseBodyData) SetTime(v string) *ListDialogResponseBodyData {
	s.Time = &v
	return s
}

func (s *ListDialogResponseBodyData) SetTag(v string) *ListDialogResponseBodyData {
	s.Tag = &v
	return s
}

type ListDialogResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDialogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDialogResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDialogResponse) GoString() string {
	return s.String()
}

func (s *ListDialogResponse) SetHeaders(v map[string]*string) *ListDialogResponse {
	s.Headers = v
	return s
}

func (s *ListDialogResponse) SetBody(v *ListDialogResponseBody) *ListDialogResponse {
	s.Body = v
	return s
}

type ListHotlineRecordRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CallId      *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
}

func (s ListHotlineRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordRequest) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordRequest) SetClientToken(v string) *ListHotlineRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *ListHotlineRecordRequest) SetInstanceId(v string) *ListHotlineRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHotlineRecordRequest) SetCallId(v string) *ListHotlineRecordRequest {
	s.CallId = &v
	return s
}

type ListHotlineRecordResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      []*ListHotlineRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListHotlineRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordResponseBody) SetCode(v string) *ListHotlineRecordResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetMessage(v string) *ListHotlineRecordResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetRequestId(v string) *ListHotlineRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetSuccess(v bool) *ListHotlineRecordResponseBody {
	s.Success = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetData(v []*ListHotlineRecordResponseBodyData) *ListHotlineRecordResponseBody {
	s.Data = v
	return s
}

type ListHotlineRecordResponseBodyData struct {
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
}

func (s ListHotlineRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordResponseBodyData) SetEndTime(v int64) *ListHotlineRecordResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetStartTime(v int64) *ListHotlineRecordResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetUrl(v string) *ListHotlineRecordResponseBodyData {
	s.Url = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetConnectionId(v string) *ListHotlineRecordResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetCallId(v string) *ListHotlineRecordResponseBodyData {
	s.CallId = &v
	return s
}

type ListHotlineRecordResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHotlineRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHotlineRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordResponse) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordResponse) SetHeaders(v map[string]*string) *ListHotlineRecordResponse {
	s.Headers = v
	return s
}

func (s *ListHotlineRecordResponse) SetBody(v *ListHotlineRecordResponseBody) *ListHotlineRecordResponse {
	s.Body = v
	return s
}

type ListHotlineRecordDetailRequest struct {
	// clientToken
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 热线挂断的时间范围
	CloseTimeEnd *int64 `json:"CloseTimeEnd,omitempty" xml:"CloseTimeEnd,omitempty"`
	// 当前页
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 每页数据量
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 热线挂断的时间范围
	CloseTimeStart *int64 `json:"CloseTimeStart,omitempty" xml:"CloseTimeStart,omitempty"`
}

func (s ListHotlineRecordDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordDetailRequest) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordDetailRequest) SetClientToken(v string) *ListHotlineRecordDetailRequest {
	s.ClientToken = &v
	return s
}

func (s *ListHotlineRecordDetailRequest) SetInstanceId(v string) *ListHotlineRecordDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHotlineRecordDetailRequest) SetCloseTimeEnd(v int64) *ListHotlineRecordDetailRequest {
	s.CloseTimeEnd = &v
	return s
}

func (s *ListHotlineRecordDetailRequest) SetCurrentPage(v int32) *ListHotlineRecordDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListHotlineRecordDetailRequest) SetPageSize(v int32) *ListHotlineRecordDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListHotlineRecordDetailRequest) SetCloseTimeStart(v int64) *ListHotlineRecordDetailRequest {
	s.CloseTimeStart = &v
	return s
}

type ListHotlineRecordDetailResponseBody struct {
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// data
	ResultData *ListHotlineRecordDetailResponseBodyResultData `json:"ResultData,omitempty" xml:"ResultData,omitempty" type:"Struct"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListHotlineRecordDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordDetailResponseBody) SetMessage(v string) *ListHotlineRecordDetailResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBody) SetRequestId(v string) *ListHotlineRecordDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBody) SetHttpStatusCode(v int32) *ListHotlineRecordDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBody) SetResultData(v *ListHotlineRecordDetailResponseBodyResultData) *ListHotlineRecordDetailResponseBody {
	s.ResultData = v
	return s
}

func (s *ListHotlineRecordDetailResponseBody) SetCode(v string) *ListHotlineRecordDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBody) SetSuccess(v bool) *ListHotlineRecordDetailResponseBody {
	s.Success = &v
	return s
}

type ListHotlineRecordDetailResponseBodyResultData struct {
	CurrentPage  *int64                                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	TotalResults *int64                                               `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
	TotalPage    *int64                                               `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	OnePageSize  *int64                                               `json:"OnePageSize,omitempty" xml:"OnePageSize,omitempty"`
	Data         []*ListHotlineRecordDetailResponseBodyResultDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListHotlineRecordDetailResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordDetailResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordDetailResponseBodyResultData) SetCurrentPage(v int64) *ListHotlineRecordDetailResponseBodyResultData {
	s.CurrentPage = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultData) SetTotalResults(v int64) *ListHotlineRecordDetailResponseBodyResultData {
	s.TotalResults = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultData) SetTotalPage(v int64) *ListHotlineRecordDetailResponseBodyResultData {
	s.TotalPage = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultData) SetOnePageSize(v int64) *ListHotlineRecordDetailResponseBodyResultData {
	s.OnePageSize = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultData) SetData(v []*ListHotlineRecordDetailResponseBodyResultDataData) *ListHotlineRecordDetailResponseBodyResultData {
	s.Data = v
	return s
}

type ListHotlineRecordDetailResponseBodyResultDataData struct {
	ServicerName *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	// 热线开始时间
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 热线结束时间
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 热线通话录音地址
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
}

func (s ListHotlineRecordDetailResponseBodyResultDataData) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordDetailResponseBodyResultDataData) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordDetailResponseBodyResultDataData) SetServicerName(v string) *ListHotlineRecordDetailResponseBodyResultDataData {
	s.ServicerName = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultDataData) SetStartTime(v int64) *ListHotlineRecordDetailResponseBodyResultDataData {
	s.StartTime = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultDataData) SetEndTime(v int64) *ListHotlineRecordDetailResponseBodyResultDataData {
	s.EndTime = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultDataData) SetOssUrl(v string) *ListHotlineRecordDetailResponseBodyResultDataData {
	s.OssUrl = &v
	return s
}

type ListHotlineRecordDetailResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHotlineRecordDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHotlineRecordDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordDetailResponse) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordDetailResponse) SetHeaders(v map[string]*string) *ListHotlineRecordDetailResponse {
	s.Headers = v
	return s
}

func (s *ListHotlineRecordDetailResponse) SetBody(v *ListHotlineRecordDetailResponseBody) *ListHotlineRecordDetailResponse {
	s.Body = v
	return s
}

type ListOutboundPhoneNumberRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s ListOutboundPhoneNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *ListOutboundPhoneNumberRequest) SetClientToken(v string) *ListOutboundPhoneNumberRequest {
	s.ClientToken = &v
	return s
}

func (s *ListOutboundPhoneNumberRequest) SetInstanceId(v string) *ListOutboundPhoneNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOutboundPhoneNumberRequest) SetAccountName(v string) *ListOutboundPhoneNumberRequest {
	s.AccountName = &v
	return s
}

type ListOutboundPhoneNumberResponseBody struct {
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s ListOutboundPhoneNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ListOutboundPhoneNumberResponseBody) SetMessage(v string) *ListOutboundPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetRequestId(v string) *ListOutboundPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetData(v []*string) *ListOutboundPhoneNumberResponseBody {
	s.Data = v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetCode(v string) *ListOutboundPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetSuccess(v bool) *ListOutboundPhoneNumberResponseBody {
	s.Success = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetHttpStatusCode(v int64) *ListOutboundPhoneNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

type ListOutboundPhoneNumberResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOutboundPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOutboundPhoneNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *ListOutboundPhoneNumberResponse) SetHeaders(v map[string]*string) *ListOutboundPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *ListOutboundPhoneNumberResponse) SetBody(v *ListOutboundPhoneNumberResponseBody) *ListOutboundPhoneNumberResponse {
	s.Body = v
	return s
}

type ListOutboundStrategiesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	BuId                 *int64  `json:"BuId,omitempty" xml:"BuId,omitempty"`
	Keyword              *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s ListOutboundStrategiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundStrategiesRequest) GoString() string {
	return s.String()
}

func (s *ListOutboundStrategiesRequest) SetOwnerId(v int64) *ListOutboundStrategiesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListOutboundStrategiesRequest) SetResourceOwnerAccount(v string) *ListOutboundStrategiesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListOutboundStrategiesRequest) SetResourceOwnerId(v int64) *ListOutboundStrategiesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListOutboundStrategiesRequest) SetProdCode(v string) *ListOutboundStrategiesRequest {
	s.ProdCode = &v
	return s
}

func (s *ListOutboundStrategiesRequest) SetBuId(v int64) *ListOutboundStrategiesRequest {
	s.BuId = &v
	return s
}

func (s *ListOutboundStrategiesRequest) SetKeyword(v string) *ListOutboundStrategiesRequest {
	s.Keyword = &v
	return s
}

type ListOutboundStrategiesResponseBody struct {
	Code               *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message            *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OutboundStrategies []*ListOutboundStrategiesResponseBodyOutboundStrategies `json:"OutboundStrategies,omitempty" xml:"OutboundStrategies,omitempty" type:"Repeated"`
}

func (s ListOutboundStrategiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOutboundStrategiesResponseBody) SetCode(v string) *ListOutboundStrategiesResponseBody {
	s.Code = &v
	return s
}

func (s *ListOutboundStrategiesResponseBody) SetMessage(v string) *ListOutboundStrategiesResponseBody {
	s.Message = &v
	return s
}

func (s *ListOutboundStrategiesResponseBody) SetRequestId(v string) *ListOutboundStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOutboundStrategiesResponseBody) SetOutboundStrategies(v []*ListOutboundStrategiesResponseBodyOutboundStrategies) *ListOutboundStrategiesResponseBody {
	s.OutboundStrategies = v
	return s
}

type ListOutboundStrategiesResponseBodyOutboundStrategies struct {
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SuccessRate          *int32  `json:"SuccessRate,omitempty" xml:"SuccessRate,omitempty"`
	Process              *int32  `json:"Process,omitempty" xml:"Process,omitempty"`
	GmtModifiedStr       *string `json:"GmtModifiedStr,omitempty" xml:"GmtModifiedStr,omitempty"`
	OutboundNum          *string `json:"OutboundNum,omitempty" xml:"OutboundNum,omitempty"`
	OutboundStrategyName *string `json:"OutboundStrategyName,omitempty" xml:"OutboundStrategyName,omitempty"`
	ModifierId           *int64  `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	SceneName            *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	OutboundStrategyId   *int64  `json:"OutboundStrategyId,omitempty" xml:"OutboundStrategyId,omitempty"`
	CreatorId            *int64  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	RobotName            *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	ModifierName         *string `json:"ModifierName,omitempty" xml:"ModifierName,omitempty"`
	ResourceAllocation   *int32  `json:"ResourceAllocation,omitempty" xml:"ResourceAllocation,omitempty"`
	ExtAttr              *string `json:"ExtAttr,omitempty" xml:"ExtAttr,omitempty"`
	NumType              *int32  `json:"NumType,omitempty" xml:"NumType,omitempty"`
	BuId                 *int64  `json:"BuId,omitempty" xml:"BuId,omitempty"`
	RobotId              *string `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	CreatorName          *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	DepartmentId         *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	RobotType            *int32  `json:"RobotType,omitempty" xml:"RobotType,omitempty"`
	RuleCode             *int64  `json:"RuleCode,omitempty" xml:"RuleCode,omitempty"`
	GmtCreateStr         *string `json:"GmtCreateStr,omitempty" xml:"GmtCreateStr,omitempty"`
}

func (s ListOutboundStrategiesResponseBodyOutboundStrategies) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundStrategiesResponseBodyOutboundStrategies) GoString() string {
	return s.String()
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetStatus(v int32) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.Status = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetSuccessRate(v int32) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.SuccessRate = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetProcess(v int32) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.Process = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetGmtModifiedStr(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.GmtModifiedStr = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetOutboundNum(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.OutboundNum = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetOutboundStrategyName(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.OutboundStrategyName = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetModifierId(v int64) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.ModifierId = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetSceneName(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.SceneName = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetOutboundStrategyId(v int64) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.OutboundStrategyId = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetCreatorId(v int64) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.CreatorId = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetRobotName(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.RobotName = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetModifierName(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.ModifierName = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetResourceAllocation(v int32) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.ResourceAllocation = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetExtAttr(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.ExtAttr = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetNumType(v int32) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.NumType = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetBuId(v int64) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.BuId = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetRobotId(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.RobotId = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetCreatorName(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.CreatorName = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetDepartmentId(v int64) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.DepartmentId = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetRobotType(v int32) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.RobotType = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetRuleCode(v int64) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.RuleCode = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetGmtCreateStr(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.GmtCreateStr = &v
	return s
}

type ListOutboundStrategiesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOutboundStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOutboundStrategiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundStrategiesResponse) GoString() string {
	return s.String()
}

func (s *ListOutboundStrategiesResponse) SetHeaders(v map[string]*string) *ListOutboundStrategiesResponse {
	s.Headers = v
	return s
}

func (s *ListOutboundStrategiesResponse) SetBody(v *ListOutboundStrategiesResponseBody) *ListOutboundStrategiesResponse {
	s.Body = v
	return s
}

type ListOuterOrderedNumbersRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
}

func (s ListOuterOrderedNumbersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOuterOrderedNumbersRequest) GoString() string {
	return s.String()
}

func (s *ListOuterOrderedNumbersRequest) SetOwnerId(v int64) *ListOuterOrderedNumbersRequest {
	s.OwnerId = &v
	return s
}

func (s *ListOuterOrderedNumbersRequest) SetResourceOwnerAccount(v string) *ListOuterOrderedNumbersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListOuterOrderedNumbersRequest) SetResourceOwnerId(v int64) *ListOuterOrderedNumbersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListOuterOrderedNumbersRequest) SetProdCode(v string) *ListOuterOrderedNumbersRequest {
	s.ProdCode = &v
	return s
}

type ListOuterOrderedNumbersResponseBody struct {
	Code      *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Numbers   []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
}

func (s ListOuterOrderedNumbersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOuterOrderedNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOuterOrderedNumbersResponseBody) SetCode(v string) *ListOuterOrderedNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *ListOuterOrderedNumbersResponseBody) SetMessage(v string) *ListOuterOrderedNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ListOuterOrderedNumbersResponseBody) SetRequestId(v string) *ListOuterOrderedNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOuterOrderedNumbersResponseBody) SetNumbers(v []*string) *ListOuterOrderedNumbersResponseBody {
	s.Numbers = v
	return s
}

type ListOuterOrderedNumbersResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOuterOrderedNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOuterOrderedNumbersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOuterOrderedNumbersResponse) GoString() string {
	return s.String()
}

func (s *ListOuterOrderedNumbersResponse) SetHeaders(v map[string]*string) *ListOuterOrderedNumbersResponse {
	s.Headers = v
	return s
}

func (s *ListOuterOrderedNumbersResponse) SetBody(v *ListOuterOrderedNumbersResponseBody) *ListOuterOrderedNumbersResponse {
	s.Body = v
	return s
}

type ListRobotCallDialogRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s ListRobotCallDialogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRobotCallDialogRequest) GoString() string {
	return s.String()
}

func (s *ListRobotCallDialogRequest) SetOwnerId(v int64) *ListRobotCallDialogRequest {
	s.OwnerId = &v
	return s
}

func (s *ListRobotCallDialogRequest) SetResourceOwnerAccount(v string) *ListRobotCallDialogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListRobotCallDialogRequest) SetResourceOwnerId(v int64) *ListRobotCallDialogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListRobotCallDialogRequest) SetCallId(v string) *ListRobotCallDialogRequest {
	s.CallId = &v
	return s
}

func (s *ListRobotCallDialogRequest) SetCreateTime(v string) *ListRobotCallDialogRequest {
	s.CreateTime = &v
	return s
}

type ListRobotCallDialogResponseBody struct {
	// Id of the request
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*ListRobotCallDialogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRobotCallDialogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRobotCallDialogResponseBody) GoString() string {
	return s.String()
}

func (s *ListRobotCallDialogResponseBody) SetRequestId(v string) *ListRobotCallDialogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRobotCallDialogResponseBody) SetData(v []*ListRobotCallDialogResponseBodyData) *ListRobotCallDialogResponseBody {
	s.Data = v
	return s
}

func (s *ListRobotCallDialogResponseBody) SetCode(v string) *ListRobotCallDialogResponseBody {
	s.Code = &v
	return s
}

func (s *ListRobotCallDialogResponseBody) SetMessage(v string) *ListRobotCallDialogResponseBody {
	s.Message = &v
	return s
}

func (s *ListRobotCallDialogResponseBody) SetSuccess(v bool) *ListRobotCallDialogResponseBody {
	s.Success = &v
	return s
}

type ListRobotCallDialogResponseBodyData struct {
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Time     *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Tag      *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListRobotCallDialogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRobotCallDialogResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRobotCallDialogResponseBodyData) SetRole(v string) *ListRobotCallDialogResponseBodyData {
	s.Role = &v
	return s
}

func (s *ListRobotCallDialogResponseBodyData) SetContent(v string) *ListRobotCallDialogResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListRobotCallDialogResponseBodyData) SetNodeType(v string) *ListRobotCallDialogResponseBodyData {
	s.NodeType = &v
	return s
}

func (s *ListRobotCallDialogResponseBodyData) SetTime(v string) *ListRobotCallDialogResponseBodyData {
	s.Time = &v
	return s
}

func (s *ListRobotCallDialogResponseBodyData) SetTag(v string) *ListRobotCallDialogResponseBodyData {
	s.Tag = &v
	return s
}

type ListRobotCallDialogResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRobotCallDialogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRobotCallDialogResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRobotCallDialogResponse) GoString() string {
	return s.String()
}

func (s *ListRobotCallDialogResponse) SetHeaders(v map[string]*string) *ListRobotCallDialogResponse {
	s.Headers = v
	return s
}

func (s *ListRobotCallDialogResponse) SetBody(v *ListRobotCallDialogResponseBody) *ListRobotCallDialogResponse {
	s.Body = v
	return s
}

type ListRolesRequest struct {
	// clientToken
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 租户实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) SetClientToken(v string) *ListRolesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListRolesRequest) SetInstanceId(v string) *ListRolesRequest {
	s.InstanceId = &v
	return s
}

type ListRolesResponseBody struct {
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// data
	Data []*ListRolesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) SetMessage(v string) *ListRolesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRolesResponseBody) SetHttpStatusCode(v int32) *ListRolesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRolesResponseBody) SetData(v []*ListRolesResponseBodyData) *ListRolesResponseBody {
	s.Data = v
	return s
}

func (s *ListRolesResponseBody) SetSuccess(v bool) *ListRolesResponseBody {
	s.Success = &v
	return s
}

type ListRolesResponseBodyData struct {
	// 角色id
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 租户id
	BuId *int64 `json:"BuId,omitempty" xml:"BuId,omitempty"`
	// 角色名称
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 角色code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 角色描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 所属角色组id
	RoleGroupId *int64 `json:"RoleGroupId,omitempty" xml:"RoleGroupId,omitempty"`
	// 所属角色组名称
	RoleGroupName *string `json:"RoleGroupName,omitempty" xml:"RoleGroupName,omitempty"`
}

func (s ListRolesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyData) SetRoleId(v int64) *ListRolesResponseBodyData {
	s.RoleId = &v
	return s
}

func (s *ListRolesResponseBodyData) SetCreateTime(v string) *ListRolesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListRolesResponseBodyData) SetBuId(v int64) *ListRolesResponseBodyData {
	s.BuId = &v
	return s
}

func (s *ListRolesResponseBodyData) SetTitle(v string) *ListRolesResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListRolesResponseBodyData) SetCode(v string) *ListRolesResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListRolesResponseBodyData) SetDescription(v string) *ListRolesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListRolesResponseBodyData) SetRoleGroupId(v int64) *ListRolesResponseBodyData {
	s.RoleGroupId = &v
	return s
}

func (s *ListRolesResponseBodyData) SetRoleGroupName(v string) *ListRolesResponseBodyData {
	s.RoleGroupName = &v
	return s
}

type ListRolesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponse) GoString() string {
	return s.String()
}

func (s *ListRolesResponse) SetHeaders(v map[string]*string) *ListRolesResponse {
	s.Headers = v
	return s
}

func (s *ListRolesResponse) SetBody(v *ListRolesResponseBody) *ListRolesResponse {
	s.Body = v
	return s
}

type ListSkillGroupRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ChannelType *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
}

func (s ListSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *ListSkillGroupRequest) SetClientToken(v string) *ListSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ListSkillGroupRequest) SetInstanceId(v string) *ListSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSkillGroupRequest) SetChannelType(v int32) *ListSkillGroupRequest {
	s.ChannelType = &v
	return s
}

type ListSkillGroupResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      []*ListSkillGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillGroupResponseBody) SetCode(v string) *ListSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListSkillGroupResponseBody) SetMessage(v string) *ListSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListSkillGroupResponseBody) SetRequestId(v string) *ListSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillGroupResponseBody) SetSuccess(v bool) *ListSkillGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListSkillGroupResponseBody) SetData(v []*ListSkillGroupResponseBodyData) *ListSkillGroupResponseBody {
	s.Data = v
	return s
}

type ListSkillGroupResponseBodyData struct {
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ChannelType  *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	SkillGroupId *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ListSkillGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillGroupResponseBodyData) SetDisplayName(v string) *ListSkillGroupResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetDescription(v string) *ListSkillGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetName(v string) *ListSkillGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetChannelType(v int32) *ListSkillGroupResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetSkillGroupId(v int64) *ListSkillGroupResponseBodyData {
	s.SkillGroupId = &v
	return s
}

type ListSkillGroupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *ListSkillGroupResponse) SetHeaders(v map[string]*string) *ListSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *ListSkillGroupResponse) SetBody(v *ListSkillGroupResponseBody) *ListSkillGroupResponse {
	s.Body = v
	return s
}

type ListTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RobotName            *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
}

func (s ListTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskRequest) GoString() string {
	return s.String()
}

func (s *ListTaskRequest) SetOwnerId(v int64) *ListTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTaskRequest) SetResourceOwnerAccount(v string) *ListTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTaskRequest) SetResourceOwnerId(v int64) *ListTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTaskRequest) SetTaskName(v string) *ListTaskRequest {
	s.TaskName = &v
	return s
}

func (s *ListTaskRequest) SetPageNo(v int32) *ListTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListTaskRequest) SetPageSize(v int32) *ListTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskRequest) SetStatus(v string) *ListTaskRequest {
	s.Status = &v
	return s
}

func (s *ListTaskRequest) SetTaskId(v int64) *ListTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ListTaskRequest) SetRobotName(v string) *ListTaskRequest {
	s.RobotName = &v
	return s
}

type ListTaskResponseBody struct {
	// Id of the request
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskResponseBody) SetRequestId(v string) *ListTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskResponseBody) SetData(v *ListTaskResponseBodyData) *ListTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListTaskResponseBody) SetCode(v string) *ListTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListTaskResponseBody) SetMessage(v string) *ListTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ListTaskResponseBody) SetSuccess(v bool) *ListTaskResponseBody {
	s.Success = &v
	return s
}

type ListTaskResponseBodyData struct {
	PageNo   *int64                            `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int64                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                            `json:"Total,omitempty" xml:"Total,omitempty"`
	Record   []*ListTaskResponseBodyDataRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s ListTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTaskResponseBodyData) SetPageNo(v int64) *ListTaskResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListTaskResponseBodyData) SetPageSize(v int64) *ListTaskResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTaskResponseBodyData) SetTotal(v int64) *ListTaskResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListTaskResponseBodyData) SetRecord(v []*ListTaskResponseBodyDataRecord) *ListTaskResponseBodyData {
	s.Record = v
	return s
}

type ListTaskResponseBodyDataRecord struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreate     *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	TotalCount    *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	FireTime      *string `json:"FireTime,omitempty" xml:"FireTime,omitempty"`
	TaskName      *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	RobotId       *int64  `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	RobotName     *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	CompleteCount *int32  `json:"CompleteCount,omitempty" xml:"CompleteCount,omitempty"`
}

func (s ListTaskResponseBodyDataRecord) String() string {
	return tea.Prettify(s)
}

func (s ListTaskResponseBodyDataRecord) GoString() string {
	return s.String()
}

func (s *ListTaskResponseBodyDataRecord) SetStatus(v string) *ListTaskResponseBodyDataRecord {
	s.Status = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetGmtCreate(v string) *ListTaskResponseBodyDataRecord {
	s.GmtCreate = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetTotalCount(v int32) *ListTaskResponseBodyDataRecord {
	s.TotalCount = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetFireTime(v string) *ListTaskResponseBodyDataRecord {
	s.FireTime = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetTaskName(v string) *ListTaskResponseBodyDataRecord {
	s.TaskName = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetRobotId(v int64) *ListTaskResponseBodyDataRecord {
	s.RobotId = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetRobotName(v string) *ListTaskResponseBodyDataRecord {
	s.RobotName = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetId(v int64) *ListTaskResponseBodyDataRecord {
	s.Id = &v
	return s
}

func (s *ListTaskResponseBodyDataRecord) SetCompleteCount(v int32) *ListTaskResponseBodyDataRecord {
	s.CompleteCount = &v
	return s
}

type ListTaskResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaskResponse) GoString() string {
	return s.String()
}

func (s *ListTaskResponse) SetHeaders(v map[string]*string) *ListTaskResponse {
	s.Headers = v
	return s
}

func (s *ListTaskResponse) SetBody(v *ListTaskResponseBody) *ListTaskResponse {
	s.Body = v
	return s
}

type ListTaskDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StatusCode           *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Called               *string `json:"Called,omitempty" xml:"Called,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListTaskDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *ListTaskDetailRequest) SetOwnerId(v int64) *ListTaskDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTaskDetailRequest) SetResourceOwnerAccount(v string) *ListTaskDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTaskDetailRequest) SetResourceOwnerId(v int64) *ListTaskDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTaskDetailRequest) SetStatusCode(v string) *ListTaskDetailRequest {
	s.StatusCode = &v
	return s
}

func (s *ListTaskDetailRequest) SetStatus(v string) *ListTaskDetailRequest {
	s.Status = &v
	return s
}

func (s *ListTaskDetailRequest) SetPageNo(v int32) *ListTaskDetailRequest {
	s.PageNo = &v
	return s
}

func (s *ListTaskDetailRequest) SetPageSize(v int32) *ListTaskDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskDetailRequest) SetCalled(v string) *ListTaskDetailRequest {
	s.Called = &v
	return s
}

func (s *ListTaskDetailRequest) SetTaskId(v int64) *ListTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *ListTaskDetailRequest) SetId(v int64) *ListTaskDetailRequest {
	s.Id = &v
	return s
}

type ListTaskDetailResponseBody struct {
	// Id of the request
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListTaskDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTaskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskDetailResponseBody) SetRequestId(v string) *ListTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskDetailResponseBody) SetData(v *ListTaskDetailResponseBodyData) *ListTaskDetailResponseBody {
	s.Data = v
	return s
}

func (s *ListTaskDetailResponseBody) SetCode(v string) *ListTaskDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListTaskDetailResponseBody) SetMessage(v string) *ListTaskDetailResponseBody {
	s.Message = &v
	return s
}

func (s *ListTaskDetailResponseBody) SetSuccess(v bool) *ListTaskDetailResponseBody {
	s.Success = &v
	return s
}

type ListTaskDetailResponseBodyData struct {
	PageNo   *int64                                  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int64                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                  `json:"Total,omitempty" xml:"Total,omitempty"`
	Record   []*ListTaskDetailResponseBodyDataRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s ListTaskDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTaskDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTaskDetailResponseBodyData) SetPageNo(v int64) *ListTaskDetailResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListTaskDetailResponseBodyData) SetPageSize(v int64) *ListTaskDetailResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTaskDetailResponseBodyData) SetTotal(v int64) *ListTaskDetailResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListTaskDetailResponseBodyData) SetRecord(v []*ListTaskDetailResponseBodyDataRecord) *ListTaskDetailResponseBodyData {
	s.Record = v
	return s
}

type ListTaskDetailResponseBodyDataRecord struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RetryCurTimes  *int32  `json:"RetryCurTimes,omitempty" xml:"RetryCurTimes,omitempty"`
	Called         *string `json:"Called,omitempty" xml:"Called,omitempty"`
	Caller         *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Duration       *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	StatusCode     *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusCodeDesc *string `json:"StatusCodeDesc,omitempty" xml:"StatusCodeDesc,omitempty"`
	RetryTimes     *int32  `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Direction      *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Tags           *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListTaskDetailResponseBodyDataRecord) String() string {
	return tea.Prettify(s)
}

func (s ListTaskDetailResponseBodyDataRecord) GoString() string {
	return s.String()
}

func (s *ListTaskDetailResponseBodyDataRecord) SetStatus(v string) *ListTaskDetailResponseBodyDataRecord {
	s.Status = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetRetryCurTimes(v int32) *ListTaskDetailResponseBodyDataRecord {
	s.RetryCurTimes = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetCalled(v string) *ListTaskDetailResponseBodyDataRecord {
	s.Called = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetCaller(v string) *ListTaskDetailResponseBodyDataRecord {
	s.Caller = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetDuration(v int32) *ListTaskDetailResponseBodyDataRecord {
	s.Duration = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetId(v int64) *ListTaskDetailResponseBodyDataRecord {
	s.Id = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetStatusCode(v string) *ListTaskDetailResponseBodyDataRecord {
	s.StatusCode = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetStatusCodeDesc(v string) *ListTaskDetailResponseBodyDataRecord {
	s.StatusCodeDesc = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetRetryTimes(v int32) *ListTaskDetailResponseBodyDataRecord {
	s.RetryTimes = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetStartTime(v string) *ListTaskDetailResponseBodyDataRecord {
	s.StartTime = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetEndTime(v string) *ListTaskDetailResponseBodyDataRecord {
	s.EndTime = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetDirection(v string) *ListTaskDetailResponseBodyDataRecord {
	s.Direction = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetTags(v string) *ListTaskDetailResponseBodyDataRecord {
	s.Tags = &v
	return s
}

type ListTaskDetailResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTaskDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *ListTaskDetailResponse) SetHeaders(v map[string]*string) *ListTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *ListTaskDetailResponse) SetBody(v *ListTaskDetailResponseBody) *ListTaskDetailResponse {
	s.Body = v
	return s
}

type MakeCallRequest struct {
	OuterAccountId   *string `json:"OuterAccountId,omitempty" xml:"OuterAccountId,omitempty"`
	OuterAccountType *string `json:"OuterAccountType,omitempty" xml:"OuterAccountType,omitempty"`
	CommandCode      *string `json:"CommandCode,omitempty" xml:"CommandCode,omitempty"`
	CallingNumber    *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	CalledNumber     *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	ExtInfo          *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s MakeCallRequest) String() string {
	return tea.Prettify(s)
}

func (s MakeCallRequest) GoString() string {
	return s.String()
}

func (s *MakeCallRequest) SetOuterAccountId(v string) *MakeCallRequest {
	s.OuterAccountId = &v
	return s
}

func (s *MakeCallRequest) SetOuterAccountType(v string) *MakeCallRequest {
	s.OuterAccountType = &v
	return s
}

func (s *MakeCallRequest) SetCommandCode(v string) *MakeCallRequest {
	s.CommandCode = &v
	return s
}

func (s *MakeCallRequest) SetCallingNumber(v string) *MakeCallRequest {
	s.CallingNumber = &v
	return s
}

func (s *MakeCallRequest) SetCalledNumber(v string) *MakeCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *MakeCallRequest) SetExtInfo(v string) *MakeCallRequest {
	s.ExtInfo = &v
	return s
}

type MakeCallResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MakeCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MakeCallResponseBody) GoString() string {
	return s.String()
}

func (s *MakeCallResponseBody) SetMessage(v string) *MakeCallResponseBody {
	s.Message = &v
	return s
}

func (s *MakeCallResponseBody) SetRequestId(v string) *MakeCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *MakeCallResponseBody) SetData(v string) *MakeCallResponseBody {
	s.Data = &v
	return s
}

func (s *MakeCallResponseBody) SetCode(v string) *MakeCallResponseBody {
	s.Code = &v
	return s
}

func (s *MakeCallResponseBody) SetSuccess(v bool) *MakeCallResponseBody {
	s.Success = &v
	return s
}

type MakeCallResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MakeCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MakeCallResponse) String() string {
	return tea.Prettify(s)
}

func (s MakeCallResponse) GoString() string {
	return s.String()
}

func (s *MakeCallResponse) SetHeaders(v map[string]*string) *MakeCallResponse {
	s.Headers = v
	return s
}

func (s *MakeCallResponse) SetBody(v *MakeCallResponseBody) *MakeCallResponse {
	s.Body = v
	return s
}

type MakeDoubleCallRequest struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 账号名称
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// 坐席手机号（需要通过坐席手机呼叫才需要填写）
	ServicerPhone *string `json:"ServicerPhone,omitempty" xml:"ServicerPhone,omitempty"`
	// 用户手机号
	MemberPhone *string `json:"MemberPhone,omitempty" xml:"MemberPhone,omitempty"`
	// 外呼主叫号码
	OutboundCallNumber *string `json:"OutboundCallNumber,omitempty" xml:"OutboundCallNumber,omitempty"`
	// 业务携带数据（JsonString）
	BizData *string `json:"BizData,omitempty" xml:"BizData,omitempty"`
}

func (s MakeDoubleCallRequest) String() string {
	return tea.Prettify(s)
}

func (s MakeDoubleCallRequest) GoString() string {
	return s.String()
}

func (s *MakeDoubleCallRequest) SetInstanceId(v string) *MakeDoubleCallRequest {
	s.InstanceId = &v
	return s
}

func (s *MakeDoubleCallRequest) SetAccountName(v string) *MakeDoubleCallRequest {
	s.AccountName = &v
	return s
}

func (s *MakeDoubleCallRequest) SetServicerPhone(v string) *MakeDoubleCallRequest {
	s.ServicerPhone = &v
	return s
}

func (s *MakeDoubleCallRequest) SetMemberPhone(v string) *MakeDoubleCallRequest {
	s.MemberPhone = &v
	return s
}

func (s *MakeDoubleCallRequest) SetOutboundCallNumber(v string) *MakeDoubleCallRequest {
	s.OutboundCallNumber = &v
	return s
}

func (s *MakeDoubleCallRequest) SetBizData(v string) *MakeDoubleCallRequest {
	s.BizData = &v
	return s
}

type MakeDoubleCallResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回数据
	Data *MakeDoubleCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s MakeDoubleCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MakeDoubleCallResponseBody) GoString() string {
	return s.String()
}

func (s *MakeDoubleCallResponseBody) SetRequestId(v string) *MakeDoubleCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *MakeDoubleCallResponseBody) SetSuccess(v bool) *MakeDoubleCallResponseBody {
	s.Success = &v
	return s
}

func (s *MakeDoubleCallResponseBody) SetCode(v string) *MakeDoubleCallResponseBody {
	s.Code = &v
	return s
}

func (s *MakeDoubleCallResponseBody) SetMessage(v string) *MakeDoubleCallResponseBody {
	s.Message = &v
	return s
}

func (s *MakeDoubleCallResponseBody) SetData(v *MakeDoubleCallResponseBodyData) *MakeDoubleCallResponseBody {
	s.Data = v
	return s
}

type MakeDoubleCallResponseBodyData struct {
	// 会话id
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
}

func (s MakeDoubleCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MakeDoubleCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *MakeDoubleCallResponseBodyData) SetAcid(v string) *MakeDoubleCallResponseBodyData {
	s.Acid = &v
	return s
}

type MakeDoubleCallResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MakeDoubleCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MakeDoubleCallResponse) String() string {
	return tea.Prettify(s)
}

func (s MakeDoubleCallResponse) GoString() string {
	return s.String()
}

func (s *MakeDoubleCallResponse) SetHeaders(v map[string]*string) *MakeDoubleCallResponse {
	s.Headers = v
	return s
}

func (s *MakeDoubleCallResponse) SetBody(v *MakeDoubleCallResponseBody) *MakeDoubleCallResponse {
	s.Body = v
	return s
}

type QueryHotlineInQueueRequest struct {
	OuterGroupId   *string `json:"OuterGroupId,omitempty" xml:"OuterGroupId,omitempty"`
	OuterGroupType *string `json:"OuterGroupType,omitempty" xml:"OuterGroupType,omitempty"`
}

func (s QueryHotlineInQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineInQueueRequest) GoString() string {
	return s.String()
}

func (s *QueryHotlineInQueueRequest) SetOuterGroupId(v string) *QueryHotlineInQueueRequest {
	s.OuterGroupId = &v
	return s
}

func (s *QueryHotlineInQueueRequest) SetOuterGroupType(v string) *QueryHotlineInQueueRequest {
	s.OuterGroupType = &v
	return s
}

type QueryHotlineInQueueResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryHotlineInQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineInQueueResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHotlineInQueueResponseBody) SetMessage(v string) *QueryHotlineInQueueResponseBody {
	s.Message = &v
	return s
}

func (s *QueryHotlineInQueueResponseBody) SetRequestId(v string) *QueryHotlineInQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHotlineInQueueResponseBody) SetData(v string) *QueryHotlineInQueueResponseBody {
	s.Data = &v
	return s
}

func (s *QueryHotlineInQueueResponseBody) SetCode(v string) *QueryHotlineInQueueResponseBody {
	s.Code = &v
	return s
}

func (s *QueryHotlineInQueueResponseBody) SetSuccess(v bool) *QueryHotlineInQueueResponseBody {
	s.Success = &v
	return s
}

type QueryHotlineInQueueResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryHotlineInQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHotlineInQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineInQueueResponse) GoString() string {
	return s.String()
}

func (s *QueryHotlineInQueueResponse) SetHeaders(v map[string]*string) *QueryHotlineInQueueResponse {
	s.Headers = v
	return s
}

func (s *QueryHotlineInQueueResponse) SetBody(v *QueryHotlineInQueueResponseBody) *QueryHotlineInQueueResponse {
	s.Body = v
	return s
}

type QueryHotlineNumberRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 当前页码
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 每页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 号码（支持模糊查询）
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// 部门id（技能组分组）
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// 技能组列表
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
}

func (s QueryHotlineNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineNumberRequest) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberRequest) SetInstanceId(v string) *QueryHotlineNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryHotlineNumberRequest) SetCurrentPage(v int32) *QueryHotlineNumberRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryHotlineNumberRequest) SetPageSize(v int32) *QueryHotlineNumberRequest {
	s.PageSize = &v
	return s
}

func (s *QueryHotlineNumberRequest) SetHotlineNumber(v string) *QueryHotlineNumberRequest {
	s.HotlineNumber = &v
	return s
}

func (s *QueryHotlineNumberRequest) SetDepartmentId(v int64) *QueryHotlineNumberRequest {
	s.DepartmentId = &v
	return s
}

func (s *QueryHotlineNumberRequest) SetGroupIds(v []*int64) *QueryHotlineNumberRequest {
	s.GroupIds = v
	return s
}

type QueryHotlineNumberShrinkRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 当前页码
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 每页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 号码（支持模糊查询）
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// 部门id（技能组分组）
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// 技能组列表
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
}

func (s QueryHotlineNumberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberShrinkRequest) SetInstanceId(v string) *QueryHotlineNumberShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryHotlineNumberShrinkRequest) SetCurrentPage(v int32) *QueryHotlineNumberShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryHotlineNumberShrinkRequest) SetPageSize(v int32) *QueryHotlineNumberShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryHotlineNumberShrinkRequest) SetHotlineNumber(v string) *QueryHotlineNumberShrinkRequest {
	s.HotlineNumber = &v
	return s
}

func (s *QueryHotlineNumberShrinkRequest) SetDepartmentId(v int64) *QueryHotlineNumberShrinkRequest {
	s.DepartmentId = &v
	return s
}

func (s *QueryHotlineNumberShrinkRequest) SetGroupIdsShrink(v string) *QueryHotlineNumberShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

type QueryHotlineNumberResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回结果数据
	Data *QueryHotlineNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryHotlineNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineNumberResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberResponseBody) SetRequestId(v string) *QueryHotlineNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHotlineNumberResponseBody) SetSuccess(v bool) *QueryHotlineNumberResponseBody {
	s.Success = &v
	return s
}

func (s *QueryHotlineNumberResponseBody) SetCode(v string) *QueryHotlineNumberResponseBody {
	s.Code = &v
	return s
}

func (s *QueryHotlineNumberResponseBody) SetMessage(v string) *QueryHotlineNumberResponseBody {
	s.Message = &v
	return s
}

func (s *QueryHotlineNumberResponseBody) SetData(v *QueryHotlineNumberResponseBodyData) *QueryHotlineNumberResponseBody {
	s.Data = v
	return s
}

type QueryHotlineNumberResponseBodyData struct {
	// 数据总量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 当前页面
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 每页大小
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 号码列表
	HotlineNumList []*QueryHotlineNumberResponseBodyDataHotlineNumList `json:"HotlineNumList,omitempty" xml:"HotlineNumList,omitempty" type:"Repeated"`
}

func (s QueryHotlineNumberResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineNumberResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberResponseBodyData) SetTotalCount(v int64) *QueryHotlineNumberResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyData) SetCurrentPage(v int64) *QueryHotlineNumberResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyData) SetPageSize(v int64) *QueryHotlineNumberResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyData) SetHotlineNumList(v []*QueryHotlineNumberResponseBodyDataHotlineNumList) *QueryHotlineNumberResponseBodyData {
	s.HotlineNumList = v
	return s
}

type QueryHotlineNumberResponseBodyDataHotlineNumList struct {
	// 号码
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// 号码描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 归属地
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// 运营商
	Sp *string `json:"Sp,omitempty" xml:"Sp,omitempty"`
	// 是否用于入呼
	InBoundEnabled *bool `json:"InBoundEnabled,omitempty" xml:"InBoundEnabled,omitempty"`
	// 入呼流程id
	FlowId *int64 `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// 入呼流程名称
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// 是否用于外呼
	OutboundEnabled *bool `json:"OutboundEnabled,omitempty" xml:"OutboundEnabled,omitempty"`
	// 外呼针对所有部门生效
	CalloutAllDepartment *bool `json:"CalloutAllDepartment,omitempty" xml:"CalloutAllDepartment,omitempty"`
	// 外呼生效列表
	CalloutRangeList []*QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList `json:"CalloutRangeList,omitempty" xml:"CalloutRangeList,omitempty" type:"Repeated"`
	// 满意度状态
	EvaluationStatus *int32 `json:"EvaluationStatus,omitempty" xml:"EvaluationStatus,omitempty"`
}

func (s QueryHotlineNumberResponseBodyDataHotlineNumList) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineNumberResponseBodyDataHotlineNumList) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetHotlineNumber(v string) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.HotlineNumber = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetDescription(v string) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.Description = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetLocation(v string) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.Location = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetSp(v string) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.Sp = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetInBoundEnabled(v bool) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.InBoundEnabled = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetFlowId(v int64) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.FlowId = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetFlowName(v string) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.FlowName = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetOutboundEnabled(v bool) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.OutboundEnabled = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetCalloutAllDepartment(v bool) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.CalloutAllDepartment = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetCalloutRangeList(v []*QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.CalloutRangeList = v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumList) SetEvaluationStatus(v int32) *QueryHotlineNumberResponseBodyDataHotlineNumList {
	s.EvaluationStatus = &v
	return s
}

type QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList struct {
	// 部门id
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// 部门名称
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// 技能组列表
	GroupDOList []*QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList `json:"GroupDOList,omitempty" xml:"GroupDOList,omitempty" type:"Repeated"`
}

func (s QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) SetDepartmentId(v int64) *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList {
	s.DepartmentId = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) SetDepartmentName(v string) *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList {
	s.DepartmentName = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList) SetGroupDOList(v []*QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList) *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeList {
	s.GroupDOList = v
	return s
}

type QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList struct {
	// 技能组id
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 技能组名称
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList) SetGroupId(v int64) *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList {
	s.GroupId = &v
	return s
}

func (s *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList) SetGroupName(v string) *QueryHotlineNumberResponseBodyDataHotlineNumListCalloutRangeListGroupDOList {
	s.GroupName = &v
	return s
}

type QueryHotlineNumberResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryHotlineNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHotlineNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineNumberResponse) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberResponse) SetHeaders(v map[string]*string) *QueryHotlineNumberResponse {
	s.Headers = v
	return s
}

func (s *QueryHotlineNumberResponse) SetBody(v *QueryHotlineNumberResponseBody) *QueryHotlineNumberResponse {
	s.Body = v
	return s
}

type QueryOutboundTaskRequest struct {
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType     *int32  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskName     *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SkillGroup   *int64  `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
	Ani          *string `json:"Ani,omitempty" xml:"Ani,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryOutboundTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryOutboundTaskRequest) SetTaskId(v int64) *QueryOutboundTaskRequest {
	s.TaskId = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetTaskType(v int32) *QueryOutboundTaskRequest {
	s.TaskType = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetTaskName(v string) *QueryOutboundTaskRequest {
	s.TaskName = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetStartDate(v string) *QueryOutboundTaskRequest {
	s.StartDate = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetEndDate(v string) *QueryOutboundTaskRequest {
	s.EndDate = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetStartTime(v string) *QueryOutboundTaskRequest {
	s.StartTime = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetEndTime(v string) *QueryOutboundTaskRequest {
	s.EndTime = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetSkillGroup(v int64) *QueryOutboundTaskRequest {
	s.SkillGroup = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetAni(v string) *QueryOutboundTaskRequest {
	s.Ani = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetStatus(v string) *QueryOutboundTaskRequest {
	s.Status = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetGroupName(v string) *QueryOutboundTaskRequest {
	s.GroupName = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetDepartmentId(v string) *QueryOutboundTaskRequest {
	s.DepartmentId = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetInstanceId(v string) *QueryOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetPageSize(v int32) *QueryOutboundTaskRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetCurrentPage(v int32) *QueryOutboundTaskRequest {
	s.CurrentPage = &v
	return s
}

type QueryOutboundTaskResponseBody struct {
	HttpStatusCode *string                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *string                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Data           *QueryOutboundTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryOutboundTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOutboundTaskResponseBody) SetHttpStatusCode(v string) *QueryOutboundTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryOutboundTaskResponseBody) SetCode(v string) *QueryOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOutboundTaskResponseBody) SetMessage(v string) *QueryOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOutboundTaskResponseBody) SetRequestId(v string) *QueryOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOutboundTaskResponseBody) SetSuccess(v string) *QueryOutboundTaskResponseBody {
	s.Success = &v
	return s
}

func (s *QueryOutboundTaskResponseBody) SetData(v *QueryOutboundTaskResponseBodyData) *QueryOutboundTaskResponseBody {
	s.Data = v
	return s
}

type QueryOutboundTaskResponseBodyData struct {
	TotalResults *string                                  `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
	CurrentPage  *string                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *string                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	List         []*QueryOutboundTaskResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s QueryOutboundTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryOutboundTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOutboundTaskResponseBodyData) SetTotalResults(v string) *QueryOutboundTaskResponseBodyData {
	s.TotalResults = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyData) SetCurrentPage(v string) *QueryOutboundTaskResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyData) SetPageSize(v string) *QueryOutboundTaskResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyData) SetList(v []*QueryOutboundTaskResponseBodyDataList) *QueryOutboundTaskResponseBodyData {
	s.List = v
	return s
}

type QueryOutboundTaskResponseBodyDataList struct {
	Type          *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	RetryInterval *int32  `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	RetryTime     *int32  `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Creator       *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	BuId          *int64  `json:"BuId,omitempty" xml:"BuId,omitempty"`
	Model         *int32  `json:"Model,omitempty" xml:"Model,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Modifier      *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DepartmentId  *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	SkillGroup    *int64  `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ExtAttrs      *string `json:"ExtAttrs,omitempty" xml:"ExtAttrs,omitempty"`
	CallerNum     *string `json:"CallerNum,omitempty" xml:"CallerNum,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s QueryOutboundTaskResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryOutboundTaskResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryOutboundTaskResponseBodyDataList) SetType(v int32) *QueryOutboundTaskResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetStatus(v int32) *QueryOutboundTaskResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetEndDate(v string) *QueryOutboundTaskResponseBodyDataList {
	s.EndDate = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetRetryInterval(v int32) *QueryOutboundTaskResponseBodyDataList {
	s.RetryInterval = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetRetryTime(v int32) *QueryOutboundTaskResponseBodyDataList {
	s.RetryTime = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetStartDate(v string) *QueryOutboundTaskResponseBodyDataList {
	s.StartDate = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetGmtModified(v int64) *QueryOutboundTaskResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetCreator(v string) *QueryOutboundTaskResponseBodyDataList {
	s.Creator = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetEndTime(v string) *QueryOutboundTaskResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetBuId(v int64) *QueryOutboundTaskResponseBodyDataList {
	s.BuId = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetModel(v int32) *QueryOutboundTaskResponseBodyDataList {
	s.Model = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetStartTime(v string) *QueryOutboundTaskResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetModifier(v string) *QueryOutboundTaskResponseBodyDataList {
	s.Modifier = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetGroupName(v string) *QueryOutboundTaskResponseBodyDataList {
	s.GroupName = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetDescription(v string) *QueryOutboundTaskResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetDepartmentId(v int64) *QueryOutboundTaskResponseBodyDataList {
	s.DepartmentId = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetGmtCreate(v int64) *QueryOutboundTaskResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetSkillGroup(v int64) *QueryOutboundTaskResponseBodyDataList {
	s.SkillGroup = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetName(v string) *QueryOutboundTaskResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetExtAttrs(v string) *QueryOutboundTaskResponseBodyDataList {
	s.ExtAttrs = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetCallerNum(v string) *QueryOutboundTaskResponseBodyDataList {
	s.CallerNum = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetId(v int64) *QueryOutboundTaskResponseBodyDataList {
	s.Id = &v
	return s
}

type QueryOutboundTaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOutboundTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryOutboundTaskResponse) SetHeaders(v map[string]*string) *QueryOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryOutboundTaskResponse) SetBody(v *QueryOutboundTaskResponseBody) *QueryOutboundTaskResponse {
	s.Body = v
	return s
}

type QuerySkillGroupsRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 技能组类型（1：热线，2：在线，4：工单）
	ChannelType *int32 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// 部门ID
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
}

func (s QuerySkillGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySkillGroupsRequest) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsRequest) SetInstanceId(v string) *QuerySkillGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetPageNo(v int32) *QuerySkillGroupsRequest {
	s.PageNo = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetPageSize(v int32) *QuerySkillGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetClientToken(v string) *QuerySkillGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetChannelType(v int32) *QuerySkillGroupsRequest {
	s.ChannelType = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetDepartmentId(v int64) *QuerySkillGroupsRequest {
	s.DepartmentId = &v
	return s
}

type QuerySkillGroupsResponseBody struct {
	OnePageSize  *int32                              `json:"OnePageSize,omitempty" xml:"OnePageSize,omitempty"`
	TotalPage    *int32                              `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage  *int32                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	TotalResults *int32                              `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
	Data         []*QuerySkillGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s QuerySkillGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySkillGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsResponseBody) SetOnePageSize(v int32) *QuerySkillGroupsResponseBody {
	s.OnePageSize = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetTotalPage(v int32) *QuerySkillGroupsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetRequestId(v string) *QuerySkillGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetCurrentPage(v int32) *QuerySkillGroupsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetTotalResults(v int32) *QuerySkillGroupsResponseBody {
	s.TotalResults = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetData(v []*QuerySkillGroupsResponseBodyData) *QuerySkillGroupsResponseBody {
	s.Data = v
	return s
}

type QuerySkillGroupsResponseBodyData struct {
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ChannelType    *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	SkillGroupId   *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s QuerySkillGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySkillGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsResponseBodyData) SetDisplayName(v string) *QuerySkillGroupsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetDescription(v string) *QuerySkillGroupsResponseBodyData {
	s.Description = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetChannelType(v int32) *QuerySkillGroupsResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetSkillGroupName(v string) *QuerySkillGroupsResponseBodyData {
	s.SkillGroupName = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetSkillGroupId(v int64) *QuerySkillGroupsResponseBodyData {
	s.SkillGroupId = &v
	return s
}

type QuerySkillGroupsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySkillGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySkillGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySkillGroupsResponse) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsResponse) SetHeaders(v map[string]*string) *QuerySkillGroupsResponse {
	s.Headers = v
	return s
}

func (s *QuerySkillGroupsResponse) SetBody(v *QuerySkillGroupsResponseBody) *QuerySkillGroupsResponse {
	s.Body = v
	return s
}

type QueryTaskDetailRequest struct {
	OutboundTaskId   *string `json:"OutboundTaskId,omitempty" xml:"OutboundTaskId,omitempty"`
	StatusList       *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	EndReasonList    *string `json:"EndReasonList,omitempty" xml:"EndReasonList,omitempty"`
	SkillGroup       *string `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
	ServicerName     *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	ServicerId       *string `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	PriorityList     *string `json:"PriorityList,omitempty" xml:"PriorityList,omitempty"`
	TaskId           *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Ani              *string `json:"Ani,omitempty" xml:"Ani,omitempty"`
	Dnis             *string `json:"Dnis,omitempty" xml:"Dnis,omitempty"`
	Sid              *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	DepartmentIdList *string `json:"DepartmentIdList,omitempty" xml:"DepartmentIdList,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage      *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryTaskDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailRequest) SetOutboundTaskId(v string) *QueryTaskDetailRequest {
	s.OutboundTaskId = &v
	return s
}

func (s *QueryTaskDetailRequest) SetStatusList(v string) *QueryTaskDetailRequest {
	s.StatusList = &v
	return s
}

func (s *QueryTaskDetailRequest) SetEndReasonList(v string) *QueryTaskDetailRequest {
	s.EndReasonList = &v
	return s
}

func (s *QueryTaskDetailRequest) SetSkillGroup(v string) *QueryTaskDetailRequest {
	s.SkillGroup = &v
	return s
}

func (s *QueryTaskDetailRequest) SetServicerName(v string) *QueryTaskDetailRequest {
	s.ServicerName = &v
	return s
}

func (s *QueryTaskDetailRequest) SetServicerId(v string) *QueryTaskDetailRequest {
	s.ServicerId = &v
	return s
}

func (s *QueryTaskDetailRequest) SetPriorityList(v string) *QueryTaskDetailRequest {
	s.PriorityList = &v
	return s
}

func (s *QueryTaskDetailRequest) SetTaskId(v int64) *QueryTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *QueryTaskDetailRequest) SetAni(v string) *QueryTaskDetailRequest {
	s.Ani = &v
	return s
}

func (s *QueryTaskDetailRequest) SetDnis(v string) *QueryTaskDetailRequest {
	s.Dnis = &v
	return s
}

func (s *QueryTaskDetailRequest) SetSid(v string) *QueryTaskDetailRequest {
	s.Sid = &v
	return s
}

func (s *QueryTaskDetailRequest) SetDepartmentIdList(v string) *QueryTaskDetailRequest {
	s.DepartmentIdList = &v
	return s
}

func (s *QueryTaskDetailRequest) SetInstanceId(v string) *QueryTaskDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTaskDetailRequest) SetPageSize(v int32) *QueryTaskDetailRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskDetailRequest) SetCurrentPage(v int32) *QueryTaskDetailRequest {
	s.CurrentPage = &v
	return s
}

type QueryTaskDetailResponseBody struct {
	HttpStatusCode *string                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *string                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Data           *QueryTaskDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryTaskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailResponseBody) SetHttpStatusCode(v string) *QueryTaskDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryTaskDetailResponseBody) SetCode(v string) *QueryTaskDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTaskDetailResponseBody) SetMessage(v string) *QueryTaskDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTaskDetailResponseBody) SetRequestId(v string) *QueryTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskDetailResponseBody) SetSuccess(v string) *QueryTaskDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTaskDetailResponseBody) SetData(v *QueryTaskDetailResponseBodyData) *QueryTaskDetailResponseBody {
	s.Data = v
	return s
}

type QueryTaskDetailResponseBodyData struct {
	TotalResults *string                                `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
	CurrentPage  *string                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *string                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	List         []*QueryTaskDetailResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s QueryTaskDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailResponseBodyData) SetTotalResults(v string) *QueryTaskDetailResponseBodyData {
	s.TotalResults = &v
	return s
}

func (s *QueryTaskDetailResponseBodyData) SetCurrentPage(v string) *QueryTaskDetailResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryTaskDetailResponseBodyData) SetPageSize(v string) *QueryTaskDetailResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryTaskDetailResponseBodyData) SetList(v []*QueryTaskDetailResponseBodyDataList) *QueryTaskDetailResponseBodyData {
	s.List = v
	return s
}

type QueryTaskDetailResponseBodyDataList struct {
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ServicerName   *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	MemberName     *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	OutboundNum    *int32  `json:"OutboundNum,omitempty" xml:"OutboundNum,omitempty"`
	RetryTime      *string `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
	Priority       *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	GmtModified    *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Dnis           *string `json:"Dnis,omitempty" xml:"Dnis,omitempty"`
	ServicerId     *int64  `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	OutboundTaskId *int64  `json:"OutboundTaskId,omitempty" xml:"OutboundTaskId,omitempty"`
	BuId           *int64  `json:"BuId,omitempty" xml:"BuId,omitempty"`
	EndReason      *int32  `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
	GmtCreate      *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	DepartmentId   *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Ani            *string `json:"Ani,omitempty" xml:"Ani,omitempty"`
	MemberId       *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	SkillGroup     *int32  `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
	ExtAttrs       *string `json:"ExtAttrs,omitempty" xml:"ExtAttrs,omitempty"`
	Id             *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s QueryTaskDetailResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailResponseBodyDataList) SetStatus(v int32) *QueryTaskDetailResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetServicerName(v string) *QueryTaskDetailResponseBodyDataList {
	s.ServicerName = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetMemberName(v string) *QueryTaskDetailResponseBodyDataList {
	s.MemberName = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetOutboundNum(v int32) *QueryTaskDetailResponseBodyDataList {
	s.OutboundNum = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetRetryTime(v string) *QueryTaskDetailResponseBodyDataList {
	s.RetryTime = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetPriority(v int32) *QueryTaskDetailResponseBodyDataList {
	s.Priority = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetGmtModified(v int64) *QueryTaskDetailResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetDnis(v string) *QueryTaskDetailResponseBodyDataList {
	s.Dnis = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetServicerId(v int64) *QueryTaskDetailResponseBodyDataList {
	s.ServicerId = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetOutboundTaskId(v int64) *QueryTaskDetailResponseBodyDataList {
	s.OutboundTaskId = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetBuId(v int64) *QueryTaskDetailResponseBodyDataList {
	s.BuId = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetEndReason(v int32) *QueryTaskDetailResponseBodyDataList {
	s.EndReason = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetGmtCreate(v int64) *QueryTaskDetailResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetDepartmentId(v int64) *QueryTaskDetailResponseBodyDataList {
	s.DepartmentId = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetAni(v string) *QueryTaskDetailResponseBodyDataList {
	s.Ani = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetMemberId(v int64) *QueryTaskDetailResponseBodyDataList {
	s.MemberId = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetSkillGroup(v int32) *QueryTaskDetailResponseBodyDataList {
	s.SkillGroup = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetExtAttrs(v string) *QueryTaskDetailResponseBodyDataList {
	s.ExtAttrs = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetId(v int32) *QueryTaskDetailResponseBodyDataList {
	s.Id = &v
	return s
}

type QueryTaskDetailResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTaskDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailResponse) SetHeaders(v map[string]*string) *QueryTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskDetailResponse) SetBody(v *QueryTaskDetailResponseBody) *QueryTaskDetailResponse {
	s.Body = v
	return s
}

type QueryTaskResultRequest struct {
	// app
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// operator
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// token
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// param1
	Param1 *string `json:"Param1,omitempty" xml:"Param1,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskResultRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskResultRequest) SetApp(v string) *QueryTaskResultRequest {
	s.App = &v
	return s
}

func (s *QueryTaskResultRequest) SetOperator(v string) *QueryTaskResultRequest {
	s.Operator = &v
	return s
}

func (s *QueryTaskResultRequest) SetToken(v string) *QueryTaskResultRequest {
	s.Token = &v
	return s
}

func (s *QueryTaskResultRequest) SetParam1(v string) *QueryTaskResultRequest {
	s.Param1 = &v
	return s
}

func (s *QueryTaskResultRequest) SetRequestId(v string) *QueryTaskResultRequest {
	s.RequestId = &v
	return s
}

type QueryTaskResultResponseBody struct {
	// status
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// buildId
	BuildId *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	// passNumber
	PassNumber *int64 `json:"PassNumber,omitempty" xml:"PassNumber,omitempty"`
	// totalNumber
	TotalNumber *int64 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
	// url
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// testSetRecordId
	TestSetRecordId *int64 `json:"TestSetRecordId,omitempty" xml:"TestSetRecordId,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskResultResponseBody) SetStatus(v string) *QueryTaskResultResponseBody {
	s.Status = &v
	return s
}

func (s *QueryTaskResultResponseBody) SetMessage(v string) *QueryTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTaskResultResponseBody) SetBuildId(v string) *QueryTaskResultResponseBody {
	s.BuildId = &v
	return s
}

func (s *QueryTaskResultResponseBody) SetPassNumber(v int64) *QueryTaskResultResponseBody {
	s.PassNumber = &v
	return s
}

func (s *QueryTaskResultResponseBody) SetTotalNumber(v int64) *QueryTaskResultResponseBody {
	s.TotalNumber = &v
	return s
}

func (s *QueryTaskResultResponseBody) SetUrl(v string) *QueryTaskResultResponseBody {
	s.Url = &v
	return s
}

func (s *QueryTaskResultResponseBody) SetSuccess(v bool) *QueryTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTaskResultResponseBody) SetTestSetRecordId(v int64) *QueryTaskResultResponseBody {
	s.TestSetRecordId = &v
	return s
}

func (s *QueryTaskResultResponseBody) SetRequestId(v string) *QueryTaskResultResponseBody {
	s.RequestId = &v
	return s
}

type QueryTaskResultResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskResultResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskResultResponse) SetHeaders(v map[string]*string) *QueryTaskResultResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskResultResponse) SetBody(v *QueryTaskResultResponseBody) *QueryTaskResultResponse {
	s.Body = v
	return s
}

type QueryTicketsRequest struct {
	InstanceId  *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CaseId      *int64                 `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	CaseType    *int32                 `json:"CaseType,omitempty" xml:"CaseType,omitempty"`
	CaseStatus  *int32                 `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	SrType      *int64                 `json:"SrType,omitempty" xml:"SrType,omitempty"`
	TaskStatus  *int32                 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	ChannelId   *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ChannelType *int32                 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	TouchId     *int64                 `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
	DealId      *int64                 `json:"DealId,omitempty" xml:"DealId,omitempty"`
	Extra       map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	PageSize    *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryTicketsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketsRequest) GoString() string {
	return s.String()
}

func (s *QueryTicketsRequest) SetInstanceId(v string) *QueryTicketsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTicketsRequest) SetCaseId(v int64) *QueryTicketsRequest {
	s.CaseId = &v
	return s
}

func (s *QueryTicketsRequest) SetCaseType(v int32) *QueryTicketsRequest {
	s.CaseType = &v
	return s
}

func (s *QueryTicketsRequest) SetCaseStatus(v int32) *QueryTicketsRequest {
	s.CaseStatus = &v
	return s
}

func (s *QueryTicketsRequest) SetSrType(v int64) *QueryTicketsRequest {
	s.SrType = &v
	return s
}

func (s *QueryTicketsRequest) SetTaskStatus(v int32) *QueryTicketsRequest {
	s.TaskStatus = &v
	return s
}

func (s *QueryTicketsRequest) SetChannelId(v string) *QueryTicketsRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryTicketsRequest) SetChannelType(v int32) *QueryTicketsRequest {
	s.ChannelType = &v
	return s
}

func (s *QueryTicketsRequest) SetTouchId(v int64) *QueryTicketsRequest {
	s.TouchId = &v
	return s
}

func (s *QueryTicketsRequest) SetDealId(v int64) *QueryTicketsRequest {
	s.DealId = &v
	return s
}

func (s *QueryTicketsRequest) SetExtra(v map[string]interface{}) *QueryTicketsRequest {
	s.Extra = v
	return s
}

func (s *QueryTicketsRequest) SetPageSize(v int32) *QueryTicketsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTicketsRequest) SetCurrentPage(v int32) *QueryTicketsRequest {
	s.CurrentPage = &v
	return s
}

type QueryTicketsShrinkRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CaseId      *int64  `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	CaseType    *int32  `json:"CaseType,omitempty" xml:"CaseType,omitempty"`
	CaseStatus  *int32  `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	SrType      *int64  `json:"SrType,omitempty" xml:"SrType,omitempty"`
	TaskStatus  *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ChannelType *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	TouchId     *int64  `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
	DealId      *int64  `json:"DealId,omitempty" xml:"DealId,omitempty"`
	ExtraShrink *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryTicketsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryTicketsShrinkRequest) SetInstanceId(v string) *QueryTicketsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCaseId(v int64) *QueryTicketsShrinkRequest {
	s.CaseId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCaseType(v int32) *QueryTicketsShrinkRequest {
	s.CaseType = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCaseStatus(v int32) *QueryTicketsShrinkRequest {
	s.CaseStatus = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetSrType(v int64) *QueryTicketsShrinkRequest {
	s.SrType = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetTaskStatus(v int32) *QueryTicketsShrinkRequest {
	s.TaskStatus = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetChannelId(v string) *QueryTicketsShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetChannelType(v int32) *QueryTicketsShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetTouchId(v int64) *QueryTicketsShrinkRequest {
	s.TouchId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetDealId(v int64) *QueryTicketsShrinkRequest {
	s.DealId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetExtraShrink(v string) *QueryTicketsShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetPageSize(v int32) *QueryTicketsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCurrentPage(v int32) *QueryTicketsShrinkRequest {
	s.CurrentPage = &v
	return s
}

type QueryTicketsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTicketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTicketsResponseBody) SetCode(v string) *QueryTicketsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTicketsResponseBody) SetMessage(v string) *QueryTicketsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTicketsResponseBody) SetData(v string) *QueryTicketsResponseBody {
	s.Data = &v
	return s
}

func (s *QueryTicketsResponseBody) SetRequestId(v string) *QueryTicketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTicketsResponseBody) SetSuccess(v bool) *QueryTicketsResponseBody {
	s.Success = &v
	return s
}

type QueryTicketsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTicketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTicketsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketsResponse) GoString() string {
	return s.String()
}

func (s *QueryTicketsResponse) SetHeaders(v map[string]*string) *QueryTicketsResponse {
	s.Headers = v
	return s
}

func (s *QueryTicketsResponse) SetBody(v *QueryTicketsResponseBody) *QueryTicketsResponse {
	s.Body = v
	return s
}

type QueryTouchListRequest struct {
	InstanceId       *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	FirstTimeStart   *int64    `json:"FirstTimeStart,omitempty" xml:"FirstTimeStart,omitempty"`
	FirstTimeEnd     *int64    `json:"FirstTimeEnd,omitempty" xml:"FirstTimeEnd,omitempty"`
	CloseTimeStart   *int64    `json:"CloseTimeStart,omitempty" xml:"CloseTimeStart,omitempty"`
	CloseTimeEnd     *int64    `json:"CloseTimeEnd,omitempty" xml:"CloseTimeEnd,omitempty"`
	TouchId          []*int64  `json:"TouchId,omitempty" xml:"TouchId,omitempty" type:"Repeated"`
	ChannelId        []*string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty" type:"Repeated"`
	ChannelType      []*int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty" type:"Repeated"`
	TouchType        []*int32  `json:"TouchType,omitempty" xml:"TouchType,omitempty" type:"Repeated"`
	MemberId         []*int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty" type:"Repeated"`
	MemberName       []*string `json:"MemberName,omitempty" xml:"MemberName,omitempty" type:"Repeated"`
	ServicerId       []*int64  `json:"ServicerId,omitempty" xml:"ServicerId,omitempty" type:"Repeated"`
	ServicerName     []*string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty" type:"Repeated"`
	QueueId          []*int64  `json:"QueueId,omitempty" xml:"QueueId,omitempty" type:"Repeated"`
	EvaluationStatus []*int32  `json:"EvaluationStatus,omitempty" xml:"EvaluationStatus,omitempty" type:"Repeated"`
	EvaluationLevel  []*int32  `json:"EvaluationLevel,omitempty" xml:"EvaluationLevel,omitempty" type:"Repeated"`
	EvaluationScore  []*int32  `json:"EvaluationScore,omitempty" xml:"EvaluationScore,omitempty" type:"Repeated"`
	PageSize         *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage      *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryTouchListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTouchListRequest) GoString() string {
	return s.String()
}

func (s *QueryTouchListRequest) SetInstanceId(v string) *QueryTouchListRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTouchListRequest) SetFirstTimeStart(v int64) *QueryTouchListRequest {
	s.FirstTimeStart = &v
	return s
}

func (s *QueryTouchListRequest) SetFirstTimeEnd(v int64) *QueryTouchListRequest {
	s.FirstTimeEnd = &v
	return s
}

func (s *QueryTouchListRequest) SetCloseTimeStart(v int64) *QueryTouchListRequest {
	s.CloseTimeStart = &v
	return s
}

func (s *QueryTouchListRequest) SetCloseTimeEnd(v int64) *QueryTouchListRequest {
	s.CloseTimeEnd = &v
	return s
}

func (s *QueryTouchListRequest) SetTouchId(v []*int64) *QueryTouchListRequest {
	s.TouchId = v
	return s
}

func (s *QueryTouchListRequest) SetChannelId(v []*string) *QueryTouchListRequest {
	s.ChannelId = v
	return s
}

func (s *QueryTouchListRequest) SetChannelType(v []*int32) *QueryTouchListRequest {
	s.ChannelType = v
	return s
}

func (s *QueryTouchListRequest) SetTouchType(v []*int32) *QueryTouchListRequest {
	s.TouchType = v
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

func (s *QueryTouchListRequest) SetServicerId(v []*int64) *QueryTouchListRequest {
	s.ServicerId = v
	return s
}

func (s *QueryTouchListRequest) SetServicerName(v []*string) *QueryTouchListRequest {
	s.ServicerName = v
	return s
}

func (s *QueryTouchListRequest) SetQueueId(v []*int64) *QueryTouchListRequest {
	s.QueueId = v
	return s
}

func (s *QueryTouchListRequest) SetEvaluationStatus(v []*int32) *QueryTouchListRequest {
	s.EvaluationStatus = v
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

func (s *QueryTouchListRequest) SetPageSize(v int32) *QueryTouchListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTouchListRequest) SetCurrentPage(v int32) *QueryTouchListRequest {
	s.CurrentPage = &v
	return s
}

type QueryTouchListResponseBody struct {
	ResultData *QueryTouchListResponseBodyResultData `json:"ResultData,omitempty" xml:"ResultData,omitempty" type:"Struct"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code       *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTouchListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTouchListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTouchListResponseBody) SetResultData(v *QueryTouchListResponseBodyResultData) *QueryTouchListResponseBody {
	s.ResultData = v
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

func (s *QueryTouchListResponseBody) SetCode(v string) *QueryTouchListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTouchListResponseBody) SetSuccess(v bool) *QueryTouchListResponseBody {
	s.Success = &v
	return s
}

type QueryTouchListResponseBodyResultData struct {
	TotalResults *int32                                      `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
	NextPage     *int32                                      `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	CurrentPage  *int32                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data         []*QueryTouchListResponseBodyResultDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	TotalPage    *int32                                      `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PreviousPage *int32                                      `json:"PreviousPage,omitempty" xml:"PreviousPage,omitempty"`
	OnePageSize  *int32                                      `json:"OnePageSize,omitempty" xml:"OnePageSize,omitempty"`
	Empty        *bool                                       `json:"Empty,omitempty" xml:"Empty,omitempty"`
}

func (s QueryTouchListResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryTouchListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryTouchListResponseBodyResultData) SetTotalResults(v int32) *QueryTouchListResponseBodyResultData {
	s.TotalResults = &v
	return s
}

func (s *QueryTouchListResponseBodyResultData) SetNextPage(v int32) *QueryTouchListResponseBodyResultData {
	s.NextPage = &v
	return s
}

func (s *QueryTouchListResponseBodyResultData) SetCurrentPage(v int32) *QueryTouchListResponseBodyResultData {
	s.CurrentPage = &v
	return s
}

func (s *QueryTouchListResponseBodyResultData) SetData(v []*QueryTouchListResponseBodyResultDataData) *QueryTouchListResponseBodyResultData {
	s.Data = v
	return s
}

func (s *QueryTouchListResponseBodyResultData) SetTotalPage(v int32) *QueryTouchListResponseBodyResultData {
	s.TotalPage = &v
	return s
}

func (s *QueryTouchListResponseBodyResultData) SetPreviousPage(v int32) *QueryTouchListResponseBodyResultData {
	s.PreviousPage = &v
	return s
}

func (s *QueryTouchListResponseBodyResultData) SetOnePageSize(v int32) *QueryTouchListResponseBodyResultData {
	s.OnePageSize = &v
	return s
}

func (s *QueryTouchListResponseBodyResultData) SetEmpty(v bool) *QueryTouchListResponseBodyResultData {
	s.Empty = &v
	return s
}

type QueryTouchListResponseBodyResultDataData struct {
	Status          *int32                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	ToId            *int64                                            `json:"ToId,omitempty" xml:"ToId,omitempty"`
	ParentTouchId   *int64                                            `json:"ParentTouchId,omitempty" xml:"ParentTouchId,omitempty"`
	ServicerName    *string                                           `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	ChannelType     *int32                                            `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	CloseTime       *int64                                            `json:"CloseTime,omitempty" xml:"CloseTime,omitempty"`
	GmtModified     *int64                                            `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ServicerId      *int64                                            `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	SwitchUser      *string                                           `json:"SwitchUser,omitempty" xml:"SwitchUser,omitempty"`
	BuId            *int64                                            `json:"BuId,omitempty" xml:"BuId,omitempty"`
	FromId          *int64                                            `json:"FromId,omitempty" xml:"FromId,omitempty"`
	UserTouchId     *int64                                            `json:"UserTouchId,omitempty" xml:"UserTouchId,omitempty"`
	TouchTime       *string                                           `json:"TouchTime,omitempty" xml:"TouchTime,omitempty"`
	TouchContent    *string                                           `json:"TouchContent,omitempty" xml:"TouchContent,omitempty"`
	Feedback        *string                                           `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	TouchId         *string                                           `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
	QueueId         *int64                                            `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	DepId           *int64                                            `json:"DepId,omitempty" xml:"DepId,omitempty"`
	TouchEndReason  *int32                                            `json:"TouchEndReason,omitempty" xml:"TouchEndReason,omitempty"`
	MemberName      *string                                           `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	CommonQueueName *string                                           `json:"CommonQueueName,omitempty" xml:"CommonQueueName,omitempty"`
	FirstTime       *int64                                            `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	TouchType       *int32                                            `json:"TouchType,omitempty" xml:"TouchType,omitempty"`
	ChannelId       *string                                           `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	GmtCreate       *int64                                            `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	MemberId        *int64                                            `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	ExtAttrs        *QueryTouchListResponseBodyResultDataDataExtAttrs `json:"ExtAttrs,omitempty" xml:"ExtAttrs,omitempty" type:"Struct"`
}

func (s QueryTouchListResponseBodyResultDataData) String() string {
	return tea.Prettify(s)
}

func (s QueryTouchListResponseBodyResultDataData) GoString() string {
	return s.String()
}

func (s *QueryTouchListResponseBodyResultDataData) SetStatus(v int32) *QueryTouchListResponseBodyResultDataData {
	s.Status = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetToId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.ToId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetParentTouchId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.ParentTouchId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetServicerName(v string) *QueryTouchListResponseBodyResultDataData {
	s.ServicerName = &v
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

func (s *QueryTouchListResponseBodyResultDataData) SetGmtModified(v int64) *QueryTouchListResponseBodyResultDataData {
	s.GmtModified = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetServicerId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.ServicerId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetSwitchUser(v string) *QueryTouchListResponseBodyResultDataData {
	s.SwitchUser = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetBuId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.BuId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetFromId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.FromId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetUserTouchId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.UserTouchId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetTouchTime(v string) *QueryTouchListResponseBodyResultDataData {
	s.TouchTime = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetTouchContent(v string) *QueryTouchListResponseBodyResultDataData {
	s.TouchContent = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetFeedback(v string) *QueryTouchListResponseBodyResultDataData {
	s.Feedback = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetTouchId(v string) *QueryTouchListResponseBodyResultDataData {
	s.TouchId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetQueueId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.QueueId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetDepId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.DepId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetTouchEndReason(v int32) *QueryTouchListResponseBodyResultDataData {
	s.TouchEndReason = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetMemberName(v string) *QueryTouchListResponseBodyResultDataData {
	s.MemberName = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetCommonQueueName(v string) *QueryTouchListResponseBodyResultDataData {
	s.CommonQueueName = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetFirstTime(v int64) *QueryTouchListResponseBodyResultDataData {
	s.FirstTime = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetTouchType(v int32) *QueryTouchListResponseBodyResultDataData {
	s.TouchType = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetChannelId(v string) *QueryTouchListResponseBodyResultDataData {
	s.ChannelId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetGmtCreate(v int64) *QueryTouchListResponseBodyResultDataData {
	s.GmtCreate = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetMemberId(v int64) *QueryTouchListResponseBodyResultDataData {
	s.MemberId = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataData) SetExtAttrs(v *QueryTouchListResponseBodyResultDataDataExtAttrs) *QueryTouchListResponseBodyResultDataData {
	s.ExtAttrs = v
	return s
}

type QueryTouchListResponseBodyResultDataDataExtAttrs struct {
	EvaluationScore        *int32 `json:"EvaluationScore,omitempty" xml:"EvaluationScore,omitempty"`
	EvaluationLevel        *int32 `json:"EvaluationLevel,omitempty" xml:"EvaluationLevel,omitempty"`
	EvaluationSolution     *int32 `json:"EvaluationSolution,omitempty" xml:"EvaluationSolution,omitempty"`
	OnlineSessionSource    *int32 `json:"OnlineSessionSource,omitempty" xml:"OnlineSessionSource,omitempty"`
	OnlineJoinRespInterval *int32 `json:"OnlineJoinRespInterval,omitempty" xml:"OnlineJoinRespInterval,omitempty"`
	EvaluationStatus       *int32 `json:"EvaluationStatus,omitempty" xml:"EvaluationStatus,omitempty"`
	// 外呼为主叫号码
	OutCallRouteNumber *string `json:"OutCallRouteNumber,omitempty" xml:"OutCallRouteNumber,omitempty"`
	// 外呼为被叫号码,入呼为被叫号码
	Dnis *string `json:"Dnis,omitempty" xml:"Dnis,omitempty"`
	// 入呼为主叫号码
	Ani *string `json:"Ani,omitempty" xml:"Ani,omitempty"`
}

func (s QueryTouchListResponseBodyResultDataDataExtAttrs) String() string {
	return tea.Prettify(s)
}

func (s QueryTouchListResponseBodyResultDataDataExtAttrs) GoString() string {
	return s.String()
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetEvaluationScore(v int32) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.EvaluationScore = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetEvaluationLevel(v int32) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.EvaluationLevel = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetEvaluationSolution(v int32) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.EvaluationSolution = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetOnlineSessionSource(v int32) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.OnlineSessionSource = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetOnlineJoinRespInterval(v int32) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.OnlineJoinRespInterval = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetEvaluationStatus(v int32) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.EvaluationStatus = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetOutCallRouteNumber(v string) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.OutCallRouteNumber = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetDnis(v string) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.Dnis = &v
	return s
}

func (s *QueryTouchListResponseBodyResultDataDataExtAttrs) SetAni(v string) *QueryTouchListResponseBodyResultDataDataExtAttrs {
	s.Ani = &v
	return s
}

type QueryTouchListResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTouchListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTouchListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTouchListResponse) GoString() string {
	return s.String()
}

func (s *QueryTouchListResponse) SetHeaders(v map[string]*string) *QueryTouchListResponse {
	s.Headers = v
	return s
}

func (s *QueryTouchListResponse) SetBody(v *QueryTouchListResponseBody) *QueryTouchListResponse {
	s.Body = v
	return s
}

type RemoveSkillGroupRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s RemoveSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupRequest) SetInstanceId(v string) *RemoveSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveSkillGroupRequest) SetSkillGroupId(v string) *RemoveSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *RemoveSkillGroupRequest) SetClientToken(v string) *RemoveSkillGroupRequest {
	s.ClientToken = &v
	return s
}

type RemoveSkillGroupResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupResponseBody) SetCode(v string) *RemoveSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveSkillGroupResponseBody) SetMessage(v string) *RemoveSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveSkillGroupResponseBody) SetRequestId(v string) *RemoveSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSkillGroupResponseBody) SetSuccess(v bool) *RemoveSkillGroupResponseBody {
	s.Success = &v
	return s
}

type RemoveSkillGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupResponse) SetHeaders(v map[string]*string) *RemoveSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveSkillGroupResponse) SetBody(v *RemoveSkillGroupResponseBody) *RemoveSkillGroupResponse {
	s.Body = v
	return s
}

type ResetHotlineNumberRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 号码
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// 号码描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 是否用于入呼
	EnableInbound *bool `json:"EnableInbound,omitempty" xml:"EnableInbound,omitempty"`
	// 入呼ivr流程id
	InboundFlowId *int64 `json:"InboundFlowId,omitempty" xml:"InboundFlowId,omitempty"`
	// 是否用于外呼
	EnableOutbound *bool `json:"EnableOutbound,omitempty" xml:"EnableOutbound,omitempty"`
	// 外呼是否对所有部门生效
	OutboundAllDepart *bool `json:"OutboundAllDepart,omitempty" xml:"OutboundAllDepart,omitempty"`
	// 是否开启入呼满意度
	EnableInboundEvaluation *bool `json:"EnableInboundEvaluation,omitempty" xml:"EnableInboundEvaluation,omitempty"`
	// 是否开启外呼满意度
	EnableOutboundEvaluation *bool `json:"EnableOutboundEvaluation,omitempty" xml:"EnableOutboundEvaluation,omitempty"`
	// 满意度等级
	EvaluationLevel *int32 `json:"EvaluationLevel,omitempty" xml:"EvaluationLevel,omitempty"`
	// 外呼生效范围
	OutboundRangeList []*ResetHotlineNumberRequestOutboundRangeList `json:"OutboundRangeList,omitempty" xml:"OutboundRangeList,omitempty" type:"Repeated"`
}

func (s ResetHotlineNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetHotlineNumberRequest) GoString() string {
	return s.String()
}

func (s *ResetHotlineNumberRequest) SetInstanceId(v string) *ResetHotlineNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetHotlineNumber(v string) *ResetHotlineNumberRequest {
	s.HotlineNumber = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetDescription(v string) *ResetHotlineNumberRequest {
	s.Description = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetEnableInbound(v bool) *ResetHotlineNumberRequest {
	s.EnableInbound = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetInboundFlowId(v int64) *ResetHotlineNumberRequest {
	s.InboundFlowId = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetEnableOutbound(v bool) *ResetHotlineNumberRequest {
	s.EnableOutbound = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetOutboundAllDepart(v bool) *ResetHotlineNumberRequest {
	s.OutboundAllDepart = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetEnableInboundEvaluation(v bool) *ResetHotlineNumberRequest {
	s.EnableInboundEvaluation = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetEnableOutboundEvaluation(v bool) *ResetHotlineNumberRequest {
	s.EnableOutboundEvaluation = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetEvaluationLevel(v int32) *ResetHotlineNumberRequest {
	s.EvaluationLevel = &v
	return s
}

func (s *ResetHotlineNumberRequest) SetOutboundRangeList(v []*ResetHotlineNumberRequestOutboundRangeList) *ResetHotlineNumberRequest {
	s.OutboundRangeList = v
	return s
}

type ResetHotlineNumberRequestOutboundRangeList struct {
	// 生效部门id
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// 生效技能组列表（部门123下）
	GroupIdList []*int64 `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty" type:"Repeated"`
}

func (s ResetHotlineNumberRequestOutboundRangeList) String() string {
	return tea.Prettify(s)
}

func (s ResetHotlineNumberRequestOutboundRangeList) GoString() string {
	return s.String()
}

func (s *ResetHotlineNumberRequestOutboundRangeList) SetDepartmentId(v int64) *ResetHotlineNumberRequestOutboundRangeList {
	s.DepartmentId = &v
	return s
}

func (s *ResetHotlineNumberRequestOutboundRangeList) SetGroupIdList(v []*int64) *ResetHotlineNumberRequestOutboundRangeList {
	s.GroupIdList = v
	return s
}

type ResetHotlineNumberShrinkRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 号码
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// 号码描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 是否用于入呼
	EnableInbound *bool `json:"EnableInbound,omitempty" xml:"EnableInbound,omitempty"`
	// 入呼ivr流程id
	InboundFlowId *int64 `json:"InboundFlowId,omitempty" xml:"InboundFlowId,omitempty"`
	// 是否用于外呼
	EnableOutbound *bool `json:"EnableOutbound,omitempty" xml:"EnableOutbound,omitempty"`
	// 外呼是否对所有部门生效
	OutboundAllDepart *bool `json:"OutboundAllDepart,omitempty" xml:"OutboundAllDepart,omitempty"`
	// 是否开启入呼满意度
	EnableInboundEvaluation *bool `json:"EnableInboundEvaluation,omitempty" xml:"EnableInboundEvaluation,omitempty"`
	// 是否开启外呼满意度
	EnableOutboundEvaluation *bool `json:"EnableOutboundEvaluation,omitempty" xml:"EnableOutboundEvaluation,omitempty"`
	// 满意度等级
	EvaluationLevel *int32 `json:"EvaluationLevel,omitempty" xml:"EvaluationLevel,omitempty"`
	// 外呼生效范围
	OutboundRangeListShrink *string `json:"OutboundRangeList,omitempty" xml:"OutboundRangeList,omitempty"`
}

func (s ResetHotlineNumberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetHotlineNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *ResetHotlineNumberShrinkRequest) SetInstanceId(v string) *ResetHotlineNumberShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetHotlineNumber(v string) *ResetHotlineNumberShrinkRequest {
	s.HotlineNumber = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetDescription(v string) *ResetHotlineNumberShrinkRequest {
	s.Description = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetEnableInbound(v bool) *ResetHotlineNumberShrinkRequest {
	s.EnableInbound = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetInboundFlowId(v int64) *ResetHotlineNumberShrinkRequest {
	s.InboundFlowId = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetEnableOutbound(v bool) *ResetHotlineNumberShrinkRequest {
	s.EnableOutbound = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetOutboundAllDepart(v bool) *ResetHotlineNumberShrinkRequest {
	s.OutboundAllDepart = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetEnableInboundEvaluation(v bool) *ResetHotlineNumberShrinkRequest {
	s.EnableInboundEvaluation = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetEnableOutboundEvaluation(v bool) *ResetHotlineNumberShrinkRequest {
	s.EnableOutboundEvaluation = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetEvaluationLevel(v int32) *ResetHotlineNumberShrinkRequest {
	s.EvaluationLevel = &v
	return s
}

func (s *ResetHotlineNumberShrinkRequest) SetOutboundRangeListShrink(v string) *ResetHotlineNumberShrinkRequest {
	s.OutboundRangeListShrink = &v
	return s
}

type ResetHotlineNumberResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 接口调用是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// http状态码
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s ResetHotlineNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetHotlineNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ResetHotlineNumberResponseBody) SetRequestId(v string) *ResetHotlineNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetHotlineNumberResponseBody) SetSuccess(v bool) *ResetHotlineNumberResponseBody {
	s.Success = &v
	return s
}

func (s *ResetHotlineNumberResponseBody) SetCode(v string) *ResetHotlineNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ResetHotlineNumberResponseBody) SetMessage(v string) *ResetHotlineNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ResetHotlineNumberResponseBody) SetHttpStatusCode(v int64) *ResetHotlineNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

type ResetHotlineNumberResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetHotlineNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetHotlineNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetHotlineNumberResponse) GoString() string {
	return s.String()
}

func (s *ResetHotlineNumberResponse) SetHeaders(v map[string]*string) *ResetHotlineNumberResponse {
	s.Headers = v
	return s
}

func (s *ResetHotlineNumberResponse) SetBody(v *ResetHotlineNumberResponseBody) *ResetHotlineNumberResponse {
	s.Body = v
	return s
}

type RestartOutboundTaskRequest struct {
	OutboundTaskId *int64  `json:"OutboundTaskId,omitempty" xml:"OutboundTaskId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RestartOutboundTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *RestartOutboundTaskRequest) SetOutboundTaskId(v int64) *RestartOutboundTaskRequest {
	s.OutboundTaskId = &v
	return s
}

func (s *RestartOutboundTaskRequest) SetInstanceId(v string) *RestartOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

type RestartOutboundTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartOutboundTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RestartOutboundTaskResponseBody) SetCode(v string) *RestartOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *RestartOutboundTaskResponseBody) SetMessage(v string) *RestartOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *RestartOutboundTaskResponseBody) SetData(v string) *RestartOutboundTaskResponseBody {
	s.Data = &v
	return s
}

func (s *RestartOutboundTaskResponseBody) SetRequestId(v string) *RestartOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartOutboundTaskResponseBody) SetSuccess(v bool) *RestartOutboundTaskResponseBody {
	s.Success = &v
	return s
}

type RestartOutboundTaskResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartOutboundTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *RestartOutboundTaskResponse) SetHeaders(v map[string]*string) *RestartOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *RestartOutboundTaskResponse) SetBody(v *RestartOutboundTaskResponseBody) *RestartOutboundTaskResponse {
	s.Body = v
	return s
}

type RobotCallRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CalledShowNumber     *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	RobotId              *int64  `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	RecordFlag           *bool   `json:"RecordFlag,omitempty" xml:"RecordFlag,omitempty"`
	EarlyMediaAsr        *bool   `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	Params               *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s RobotCallRequest) String() string {
	return tea.Prettify(s)
}

func (s RobotCallRequest) GoString() string {
	return s.String()
}

func (s *RobotCallRequest) SetOwnerId(v int64) *RobotCallRequest {
	s.OwnerId = &v
	return s
}

func (s *RobotCallRequest) SetResourceOwnerAccount(v string) *RobotCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RobotCallRequest) SetResourceOwnerId(v int64) *RobotCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RobotCallRequest) SetCalledShowNumber(v string) *RobotCallRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *RobotCallRequest) SetCalledNumber(v string) *RobotCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *RobotCallRequest) SetRobotId(v int64) *RobotCallRequest {
	s.RobotId = &v
	return s
}

func (s *RobotCallRequest) SetOutId(v string) *RobotCallRequest {
	s.OutId = &v
	return s
}

func (s *RobotCallRequest) SetRecordFlag(v bool) *RobotCallRequest {
	s.RecordFlag = &v
	return s
}

func (s *RobotCallRequest) SetEarlyMediaAsr(v bool) *RobotCallRequest {
	s.EarlyMediaAsr = &v
	return s
}

func (s *RobotCallRequest) SetParams(v string) *RobotCallRequest {
	s.Params = &v
	return s
}

type RobotCallResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RobotCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RobotCallResponseBody) GoString() string {
	return s.String()
}

func (s *RobotCallResponseBody) SetMessage(v string) *RobotCallResponseBody {
	s.Message = &v
	return s
}

func (s *RobotCallResponseBody) SetRequestId(v string) *RobotCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *RobotCallResponseBody) SetData(v string) *RobotCallResponseBody {
	s.Data = &v
	return s
}

func (s *RobotCallResponseBody) SetCode(v string) *RobotCallResponseBody {
	s.Code = &v
	return s
}

type RobotCallResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RobotCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RobotCallResponse) String() string {
	return tea.Prettify(s)
}

func (s RobotCallResponse) GoString() string {
	return s.String()
}

func (s *RobotCallResponse) SetHeaders(v map[string]*string) *RobotCallResponse {
	s.Headers = v
	return s
}

func (s *RobotCallResponse) SetBody(v *RobotCallResponseBody) *RobotCallResponse {
	s.Body = v
	return s
}

type SendCcoSmartCallRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	CalledShowNumber     *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	VoiceCode            *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	PlayTimes            *int32  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	Volume               *int32  `json:"Volume,omitempty" xml:"Volume,omitempty"`
	Speed                *int32  `json:"Speed,omitempty" xml:"Speed,omitempty"`
	AsrModelId           *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	AsrBaseId            *string `json:"AsrBaseId,omitempty" xml:"AsrBaseId,omitempty"`
	AsrAlsAmId           *string `json:"AsrAlsAmId,omitempty" xml:"AsrAlsAmId,omitempty"`
	AsrVocabularyId      *string `json:"AsrVocabularyId,omitempty" xml:"AsrVocabularyId,omitempty"`
	RecordFlag           *bool   `json:"RecordFlag,omitempty" xml:"RecordFlag,omitempty"`
	PauseTime            *int32  `json:"PauseTime,omitempty" xml:"PauseTime,omitempty"`
	MuteTime             *int32  `json:"MuteTime,omitempty" xml:"MuteTime,omitempty"`
	ActionCodeBreak      *bool   `json:"ActionCodeBreak,omitempty" xml:"ActionCodeBreak,omitempty"`
	DynamicId            *string `json:"DynamicId,omitempty" xml:"DynamicId,omitempty"`
	EarlyMediaAsr        *bool   `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	VoiceCodeParam       *string `json:"VoiceCodeParam,omitempty" xml:"VoiceCodeParam,omitempty"`
	SessionTimeout       *int32  `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	ActionCodeTimeBreak  *int32  `json:"ActionCodeTimeBreak,omitempty" xml:"ActionCodeTimeBreak,omitempty"`
	TtsConf              *bool   `json:"TtsConf,omitempty" xml:"TtsConf,omitempty"`
	TtsStyle             *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	TtsVolume            *int32  `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	TtsSpeed             *int32  `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
}

func (s SendCcoSmartCallRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCcoSmartCallRequest) GoString() string {
	return s.String()
}

func (s *SendCcoSmartCallRequest) SetOwnerId(v int64) *SendCcoSmartCallRequest {
	s.OwnerId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetResourceOwnerAccount(v string) *SendCcoSmartCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetResourceOwnerId(v int64) *SendCcoSmartCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetProdCode(v string) *SendCcoSmartCallRequest {
	s.ProdCode = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetCalledShowNumber(v string) *SendCcoSmartCallRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetCalledNumber(v string) *SendCcoSmartCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetVoiceCode(v string) *SendCcoSmartCallRequest {
	s.VoiceCode = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetOutId(v string) *SendCcoSmartCallRequest {
	s.OutId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetPlayTimes(v int32) *SendCcoSmartCallRequest {
	s.PlayTimes = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetVolume(v int32) *SendCcoSmartCallRequest {
	s.Volume = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetSpeed(v int32) *SendCcoSmartCallRequest {
	s.Speed = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetAsrModelId(v string) *SendCcoSmartCallRequest {
	s.AsrModelId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetAsrBaseId(v string) *SendCcoSmartCallRequest {
	s.AsrBaseId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetAsrAlsAmId(v string) *SendCcoSmartCallRequest {
	s.AsrAlsAmId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetAsrVocabularyId(v string) *SendCcoSmartCallRequest {
	s.AsrVocabularyId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetRecordFlag(v bool) *SendCcoSmartCallRequest {
	s.RecordFlag = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetPauseTime(v int32) *SendCcoSmartCallRequest {
	s.PauseTime = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetMuteTime(v int32) *SendCcoSmartCallRequest {
	s.MuteTime = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetActionCodeBreak(v bool) *SendCcoSmartCallRequest {
	s.ActionCodeBreak = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetDynamicId(v string) *SendCcoSmartCallRequest {
	s.DynamicId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetEarlyMediaAsr(v bool) *SendCcoSmartCallRequest {
	s.EarlyMediaAsr = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetVoiceCodeParam(v string) *SendCcoSmartCallRequest {
	s.VoiceCodeParam = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetSessionTimeout(v int32) *SendCcoSmartCallRequest {
	s.SessionTimeout = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetActionCodeTimeBreak(v int32) *SendCcoSmartCallRequest {
	s.ActionCodeTimeBreak = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetTtsConf(v bool) *SendCcoSmartCallRequest {
	s.TtsConf = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetTtsStyle(v string) *SendCcoSmartCallRequest {
	s.TtsStyle = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetTtsVolume(v int32) *SendCcoSmartCallRequest {
	s.TtsVolume = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetTtsSpeed(v int32) *SendCcoSmartCallRequest {
	s.TtsSpeed = &v
	return s
}

type SendCcoSmartCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendCcoSmartCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCcoSmartCallResponseBody) GoString() string {
	return s.String()
}

func (s *SendCcoSmartCallResponseBody) SetCode(v string) *SendCcoSmartCallResponseBody {
	s.Code = &v
	return s
}

func (s *SendCcoSmartCallResponseBody) SetMessage(v string) *SendCcoSmartCallResponseBody {
	s.Message = &v
	return s
}

func (s *SendCcoSmartCallResponseBody) SetData(v string) *SendCcoSmartCallResponseBody {
	s.Data = &v
	return s
}

func (s *SendCcoSmartCallResponseBody) SetRequestId(v string) *SendCcoSmartCallResponseBody {
	s.RequestId = &v
	return s
}

type SendCcoSmartCallResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendCcoSmartCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCcoSmartCallResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCcoSmartCallResponse) GoString() string {
	return s.String()
}

func (s *SendCcoSmartCallResponse) SetHeaders(v map[string]*string) *SendCcoSmartCallResponse {
	s.Headers = v
	return s
}

func (s *SendCcoSmartCallResponse) SetBody(v *SendCcoSmartCallResponseBody) *SendCcoSmartCallResponse {
	s.Body = v
	return s
}

type SendCcoSmartCallOperateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Command              *string `json:"Command,omitempty" xml:"Command,omitempty"`
	Param                *string `json:"Param,omitempty" xml:"Param,omitempty"`
}

func (s SendCcoSmartCallOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCcoSmartCallOperateRequest) GoString() string {
	return s.String()
}

func (s *SendCcoSmartCallOperateRequest) SetOwnerId(v int64) *SendCcoSmartCallOperateRequest {
	s.OwnerId = &v
	return s
}

func (s *SendCcoSmartCallOperateRequest) SetResourceOwnerAccount(v string) *SendCcoSmartCallOperateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendCcoSmartCallOperateRequest) SetResourceOwnerId(v int64) *SendCcoSmartCallOperateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendCcoSmartCallOperateRequest) SetProdCode(v string) *SendCcoSmartCallOperateRequest {
	s.ProdCode = &v
	return s
}

func (s *SendCcoSmartCallOperateRequest) SetCallId(v string) *SendCcoSmartCallOperateRequest {
	s.CallId = &v
	return s
}

func (s *SendCcoSmartCallOperateRequest) SetCommand(v string) *SendCcoSmartCallOperateRequest {
	s.Command = &v
	return s
}

func (s *SendCcoSmartCallOperateRequest) SetParam(v string) *SendCcoSmartCallOperateRequest {
	s.Param = &v
	return s
}

type SendCcoSmartCallOperateResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendCcoSmartCallOperateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCcoSmartCallOperateResponseBody) GoString() string {
	return s.String()
}

func (s *SendCcoSmartCallOperateResponseBody) SetCode(v string) *SendCcoSmartCallOperateResponseBody {
	s.Code = &v
	return s
}

func (s *SendCcoSmartCallOperateResponseBody) SetMessage(v string) *SendCcoSmartCallOperateResponseBody {
	s.Message = &v
	return s
}

func (s *SendCcoSmartCallOperateResponseBody) SetData(v string) *SendCcoSmartCallOperateResponseBody {
	s.Data = &v
	return s
}

func (s *SendCcoSmartCallOperateResponseBody) SetRequestId(v string) *SendCcoSmartCallOperateResponseBody {
	s.RequestId = &v
	return s
}

type SendCcoSmartCallOperateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendCcoSmartCallOperateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCcoSmartCallOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCcoSmartCallOperateResponse) GoString() string {
	return s.String()
}

func (s *SendCcoSmartCallOperateResponse) SetHeaders(v map[string]*string) *SendCcoSmartCallOperateResponse {
	s.Headers = v
	return s
}

func (s *SendCcoSmartCallOperateResponse) SetBody(v *SendCcoSmartCallOperateResponseBody) *SendCcoSmartCallOperateResponse {
	s.Body = v
	return s
}

type SendHotlineHeartBeatRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Token       *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s SendHotlineHeartBeatRequest) String() string {
	return tea.Prettify(s)
}

func (s SendHotlineHeartBeatRequest) GoString() string {
	return s.String()
}

func (s *SendHotlineHeartBeatRequest) SetClientToken(v string) *SendHotlineHeartBeatRequest {
	s.ClientToken = &v
	return s
}

func (s *SendHotlineHeartBeatRequest) SetInstanceId(v string) *SendHotlineHeartBeatRequest {
	s.InstanceId = &v
	return s
}

func (s *SendHotlineHeartBeatRequest) SetAccountName(v string) *SendHotlineHeartBeatRequest {
	s.AccountName = &v
	return s
}

func (s *SendHotlineHeartBeatRequest) SetToken(v string) *SendHotlineHeartBeatRequest {
	s.Token = &v
	return s
}

type SendHotlineHeartBeatResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendHotlineHeartBeatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendHotlineHeartBeatResponseBody) GoString() string {
	return s.String()
}

func (s *SendHotlineHeartBeatResponseBody) SetCode(v string) *SendHotlineHeartBeatResponseBody {
	s.Code = &v
	return s
}

func (s *SendHotlineHeartBeatResponseBody) SetMessage(v string) *SendHotlineHeartBeatResponseBody {
	s.Message = &v
	return s
}

func (s *SendHotlineHeartBeatResponseBody) SetRequestId(v string) *SendHotlineHeartBeatResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendHotlineHeartBeatResponseBody) SetSuccess(v bool) *SendHotlineHeartBeatResponseBody {
	s.Success = &v
	return s
}

type SendHotlineHeartBeatResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendHotlineHeartBeatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendHotlineHeartBeatResponse) String() string {
	return tea.Prettify(s)
}

func (s SendHotlineHeartBeatResponse) GoString() string {
	return s.String()
}

func (s *SendHotlineHeartBeatResponse) SetHeaders(v map[string]*string) *SendHotlineHeartBeatResponse {
	s.Headers = v
	return s
}

func (s *SendHotlineHeartBeatResponse) SetBody(v *SendHotlineHeartBeatResponseBody) *SendHotlineHeartBeatResponse {
	s.Body = v
	return s
}

type StartCallRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Caller      *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Callee      *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
}

func (s StartCallRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCallRequest) GoString() string {
	return s.String()
}

func (s *StartCallRequest) SetClientToken(v string) *StartCallRequest {
	s.ClientToken = &v
	return s
}

func (s *StartCallRequest) SetInstanceId(v string) *StartCallRequest {
	s.InstanceId = &v
	return s
}

func (s *StartCallRequest) SetAccountName(v string) *StartCallRequest {
	s.AccountName = &v
	return s
}

func (s *StartCallRequest) SetCaller(v string) *StartCallRequest {
	s.Caller = &v
	return s
}

func (s *StartCallRequest) SetCallee(v string) *StartCallRequest {
	s.Callee = &v
	return s
}

type StartCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCallResponseBody) GoString() string {
	return s.String()
}

func (s *StartCallResponseBody) SetCode(v string) *StartCallResponseBody {
	s.Code = &v
	return s
}

func (s *StartCallResponseBody) SetMessage(v string) *StartCallResponseBody {
	s.Message = &v
	return s
}

func (s *StartCallResponseBody) SetRequestId(v string) *StartCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCallResponseBody) SetSuccess(v bool) *StartCallResponseBody {
	s.Success = &v
	return s
}

type StartCallResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCallResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCallResponse) GoString() string {
	return s.String()
}

func (s *StartCallResponse) SetHeaders(v map[string]*string) *StartCallResponse {
	s.Headers = v
	return s
}

func (s *StartCallResponse) SetBody(v *StartCallResponseBody) *StartCallResponse {
	s.Body = v
	return s
}

type StartCallV2Request struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Caller      *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	CallerType  *int32  `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	Callee      *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
}

func (s StartCallV2Request) String() string {
	return tea.Prettify(s)
}

func (s StartCallV2Request) GoString() string {
	return s.String()
}

func (s *StartCallV2Request) SetClientToken(v string) *StartCallV2Request {
	s.ClientToken = &v
	return s
}

func (s *StartCallV2Request) SetInstanceId(v string) *StartCallV2Request {
	s.InstanceId = &v
	return s
}

func (s *StartCallV2Request) SetAccountName(v string) *StartCallV2Request {
	s.AccountName = &v
	return s
}

func (s *StartCallV2Request) SetCaller(v string) *StartCallV2Request {
	s.Caller = &v
	return s
}

func (s *StartCallV2Request) SetCallerType(v int32) *StartCallV2Request {
	s.CallerType = &v
	return s
}

func (s *StartCallV2Request) SetCallee(v string) *StartCallV2Request {
	s.Callee = &v
	return s
}

type StartCallV2ResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartCallV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCallV2ResponseBody) GoString() string {
	return s.String()
}

func (s *StartCallV2ResponseBody) SetCode(v string) *StartCallV2ResponseBody {
	s.Code = &v
	return s
}

func (s *StartCallV2ResponseBody) SetMessage(v string) *StartCallV2ResponseBody {
	s.Message = &v
	return s
}

func (s *StartCallV2ResponseBody) SetRequestId(v string) *StartCallV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCallV2ResponseBody) SetSuccess(v bool) *StartCallV2ResponseBody {
	s.Success = &v
	return s
}

type StartCallV2Response struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartCallV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCallV2Response) String() string {
	return tea.Prettify(s)
}

func (s StartCallV2Response) GoString() string {
	return s.String()
}

func (s *StartCallV2Response) SetHeaders(v map[string]*string) *StartCallV2Response {
	s.Headers = v
	return s
}

func (s *StartCallV2Response) SetBody(v *StartCallV2ResponseBody) *StartCallV2Response {
	s.Body = v
	return s
}

type StartChatWorkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// accountName
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s StartChatWorkRequest) String() string {
	return tea.Prettify(s)
}

func (s StartChatWorkRequest) GoString() string {
	return s.String()
}

func (s *StartChatWorkRequest) SetInstanceId(v string) *StartChatWorkRequest {
	s.InstanceId = &v
	return s
}

func (s *StartChatWorkRequest) SetAccountName(v string) *StartChatWorkRequest {
	s.AccountName = &v
	return s
}

type StartChatWorkResponseBody struct {
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartChatWorkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartChatWorkResponseBody) GoString() string {
	return s.String()
}

func (s *StartChatWorkResponseBody) SetMessage(v string) *StartChatWorkResponseBody {
	s.Message = &v
	return s
}

func (s *StartChatWorkResponseBody) SetRequestId(v string) *StartChatWorkResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartChatWorkResponseBody) SetHttpStatusCode(v int32) *StartChatWorkResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartChatWorkResponseBody) SetData(v string) *StartChatWorkResponseBody {
	s.Data = &v
	return s
}

func (s *StartChatWorkResponseBody) SetCode(v string) *StartChatWorkResponseBody {
	s.Code = &v
	return s
}

func (s *StartChatWorkResponseBody) SetSuccess(v bool) *StartChatWorkResponseBody {
	s.Success = &v
	return s
}

type StartChatWorkResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartChatWorkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartChatWorkResponse) String() string {
	return tea.Prettify(s)
}

func (s StartChatWorkResponse) GoString() string {
	return s.String()
}

func (s *StartChatWorkResponse) SetHeaders(v map[string]*string) *StartChatWorkResponse {
	s.Headers = v
	return s
}

func (s *StartChatWorkResponse) SetBody(v *StartChatWorkResponseBody) *StartChatWorkResponse {
	s.Body = v
	return s
}

type StartHotlineServiceRequest struct {
	// js sdk中自动生成的鉴权token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s StartHotlineServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartHotlineServiceRequest) GoString() string {
	return s.String()
}

func (s *StartHotlineServiceRequest) SetClientToken(v string) *StartHotlineServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartHotlineServiceRequest) SetInstanceId(v string) *StartHotlineServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartHotlineServiceRequest) SetAccountName(v string) *StartHotlineServiceRequest {
	s.AccountName = &v
	return s
}

type StartHotlineServiceResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s StartHotlineServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartHotlineServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StartHotlineServiceResponseBody) SetMessage(v string) *StartHotlineServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetRequestId(v string) *StartHotlineServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetData(v string) *StartHotlineServiceResponseBody {
	s.Data = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetCode(v string) *StartHotlineServiceResponseBody {
	s.Code = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetSuccess(v bool) *StartHotlineServiceResponseBody {
	s.Success = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetHttpStatusCode(v int64) *StartHotlineServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

type StartHotlineServiceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartHotlineServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartHotlineServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartHotlineServiceResponse) GoString() string {
	return s.String()
}

func (s *StartHotlineServiceResponse) SetHeaders(v map[string]*string) *StartHotlineServiceResponse {
	s.Headers = v
	return s
}

func (s *StartHotlineServiceResponse) SetBody(v *StartHotlineServiceResponseBody) *StartHotlineServiceResponse {
	s.Body = v
	return s
}

type StartMicroOutboundRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountId            *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CommandCode          *string `json:"CommandCode,omitempty" xml:"CommandCode,omitempty"`
	CallingNumber        *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	ExtInfo              *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	AppName              *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s StartMicroOutboundRequest) String() string {
	return tea.Prettify(s)
}

func (s StartMicroOutboundRequest) GoString() string {
	return s.String()
}

func (s *StartMicroOutboundRequest) SetOwnerId(v int64) *StartMicroOutboundRequest {
	s.OwnerId = &v
	return s
}

func (s *StartMicroOutboundRequest) SetResourceOwnerAccount(v string) *StartMicroOutboundRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartMicroOutboundRequest) SetResourceOwnerId(v int64) *StartMicroOutboundRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartMicroOutboundRequest) SetProdCode(v string) *StartMicroOutboundRequest {
	s.ProdCode = &v
	return s
}

func (s *StartMicroOutboundRequest) SetAccountType(v string) *StartMicroOutboundRequest {
	s.AccountType = &v
	return s
}

func (s *StartMicroOutboundRequest) SetAccountId(v string) *StartMicroOutboundRequest {
	s.AccountId = &v
	return s
}

func (s *StartMicroOutboundRequest) SetCommandCode(v string) *StartMicroOutboundRequest {
	s.CommandCode = &v
	return s
}

func (s *StartMicroOutboundRequest) SetCallingNumber(v string) *StartMicroOutboundRequest {
	s.CallingNumber = &v
	return s
}

func (s *StartMicroOutboundRequest) SetCalledNumber(v string) *StartMicroOutboundRequest {
	s.CalledNumber = &v
	return s
}

func (s *StartMicroOutboundRequest) SetExtInfo(v string) *StartMicroOutboundRequest {
	s.ExtInfo = &v
	return s
}

func (s *StartMicroOutboundRequest) SetAppName(v string) *StartMicroOutboundRequest {
	s.AppName = &v
	return s
}

type StartMicroOutboundResponseBody struct {
	CustomerInfo     *string `json:"CustomerInfo,omitempty" xml:"CustomerInfo,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InvokeCmdId      *string `json:"InvokeCmdId,omitempty" xml:"InvokeCmdId,omitempty"`
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	InvokeCreateTime *string `json:"InvokeCreateTime,omitempty" xml:"InvokeCreateTime,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s StartMicroOutboundResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartMicroOutboundResponseBody) GoString() string {
	return s.String()
}

func (s *StartMicroOutboundResponseBody) SetCustomerInfo(v string) *StartMicroOutboundResponseBody {
	s.CustomerInfo = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetRequestId(v string) *StartMicroOutboundResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetInvokeCmdId(v string) *StartMicroOutboundResponseBody {
	s.InvokeCmdId = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetCode(v string) *StartMicroOutboundResponseBody {
	s.Code = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetInvokeCreateTime(v string) *StartMicroOutboundResponseBody {
	s.InvokeCreateTime = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetMessage(v string) *StartMicroOutboundResponseBody {
	s.Message = &v
	return s
}

type StartMicroOutboundResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartMicroOutboundResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartMicroOutboundResponse) String() string {
	return tea.Prettify(s)
}

func (s StartMicroOutboundResponse) GoString() string {
	return s.String()
}

func (s *StartMicroOutboundResponse) SetHeaders(v map[string]*string) *StartMicroOutboundResponse {
	s.Headers = v
	return s
}

func (s *StartMicroOutboundResponse) SetBody(v *StartMicroOutboundResponseBody) *StartMicroOutboundResponse {
	s.Body = v
	return s
}

type StartTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	StartNow             *bool   `json:"StartNow,omitempty" xml:"StartNow,omitempty"`
}

func (s StartTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StartTaskRequest) GoString() string {
	return s.String()
}

func (s *StartTaskRequest) SetOwnerId(v int64) *StartTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StartTaskRequest) SetResourceOwnerAccount(v string) *StartTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartTaskRequest) SetResourceOwnerId(v int64) *StartTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartTaskRequest) SetTaskId(v int64) *StartTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StartTaskRequest) SetStartNow(v bool) *StartTaskRequest {
	s.StartNow = &v
	return s
}

type StartTaskResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartTaskResponseBody) SetRequestId(v string) *StartTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTaskResponseBody) SetData(v bool) *StartTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StartTaskResponseBody) SetCode(v string) *StartTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StartTaskResponseBody) SetMessage(v string) *StartTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StartTaskResponseBody) SetSuccess(v bool) *StartTaskResponseBody {
	s.Success = &v
	return s
}

type StartTaskResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartTaskResponse) GoString() string {
	return s.String()
}

func (s *StartTaskResponse) SetHeaders(v map[string]*string) *StartTaskResponse {
	s.Headers = v
	return s
}

func (s *StartTaskResponse) SetBody(v *StartTaskResponseBody) *StartTaskResponse {
	s.Body = v
	return s
}

type StartTaskByAppRequest struct {
	// app
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// operator
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// token
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// envTypes
	EnvTypes *string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty"`
	// appName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// appId
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// buildId
	BuildId *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	// siteType
	SiteType *string `json:"SiteType,omitempty" xml:"SiteType,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartTaskByAppRequest) String() string {
	return tea.Prettify(s)
}

func (s StartTaskByAppRequest) GoString() string {
	return s.String()
}

func (s *StartTaskByAppRequest) SetApp(v string) *StartTaskByAppRequest {
	s.App = &v
	return s
}

func (s *StartTaskByAppRequest) SetOperator(v string) *StartTaskByAppRequest {
	s.Operator = &v
	return s
}

func (s *StartTaskByAppRequest) SetToken(v string) *StartTaskByAppRequest {
	s.Token = &v
	return s
}

func (s *StartTaskByAppRequest) SetEnvTypes(v string) *StartTaskByAppRequest {
	s.EnvTypes = &v
	return s
}

func (s *StartTaskByAppRequest) SetAppName(v string) *StartTaskByAppRequest {
	s.AppName = &v
	return s
}

func (s *StartTaskByAppRequest) SetAppId(v int32) *StartTaskByAppRequest {
	s.AppId = &v
	return s
}

func (s *StartTaskByAppRequest) SetBuildId(v string) *StartTaskByAppRequest {
	s.BuildId = &v
	return s
}

func (s *StartTaskByAppRequest) SetSiteType(v string) *StartTaskByAppRequest {
	s.SiteType = &v
	return s
}

func (s *StartTaskByAppRequest) SetRequestId(v string) *StartTaskByAppRequest {
	s.RequestId = &v
	return s
}

type StartTaskByAppResponseBody struct {
	// status
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// buildId
	BuildId *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	// passNumber
	PassNumber *int64 `json:"PassNumber,omitempty" xml:"PassNumber,omitempty"`
	// totalNumber
	TotalNumber *int64 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
	// url
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// testSetRecordId
	TestSetRecordId *int64 `json:"TestSetRecordId,omitempty" xml:"TestSetRecordId,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartTaskByAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartTaskByAppResponseBody) GoString() string {
	return s.String()
}

func (s *StartTaskByAppResponseBody) SetStatus(v string) *StartTaskByAppResponseBody {
	s.Status = &v
	return s
}

func (s *StartTaskByAppResponseBody) SetMessage(v string) *StartTaskByAppResponseBody {
	s.Message = &v
	return s
}

func (s *StartTaskByAppResponseBody) SetBuildId(v string) *StartTaskByAppResponseBody {
	s.BuildId = &v
	return s
}

func (s *StartTaskByAppResponseBody) SetPassNumber(v int64) *StartTaskByAppResponseBody {
	s.PassNumber = &v
	return s
}

func (s *StartTaskByAppResponseBody) SetTotalNumber(v int64) *StartTaskByAppResponseBody {
	s.TotalNumber = &v
	return s
}

func (s *StartTaskByAppResponseBody) SetUrl(v string) *StartTaskByAppResponseBody {
	s.Url = &v
	return s
}

func (s *StartTaskByAppResponseBody) SetSuccess(v bool) *StartTaskByAppResponseBody {
	s.Success = &v
	return s
}

func (s *StartTaskByAppResponseBody) SetTestSetRecordId(v int64) *StartTaskByAppResponseBody {
	s.TestSetRecordId = &v
	return s
}

func (s *StartTaskByAppResponseBody) SetRequestId(v string) *StartTaskByAppResponseBody {
	s.RequestId = &v
	return s
}

type StartTaskByAppResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartTaskByAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartTaskByAppResponse) String() string {
	return tea.Prettify(s)
}

func (s StartTaskByAppResponse) GoString() string {
	return s.String()
}

func (s *StartTaskByAppResponse) SetHeaders(v map[string]*string) *StartTaskByAppResponse {
	s.Headers = v
	return s
}

func (s *StartTaskByAppResponse) SetBody(v *StartTaskByAppResponseBody) *StartTaskByAppResponse {
	s.Body = v
	return s
}

type StopTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StopTaskRequest) GoString() string {
	return s.String()
}

func (s *StopTaskRequest) SetOwnerId(v int64) *StopTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StopTaskRequest) SetResourceOwnerAccount(v string) *StopTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopTaskRequest) SetResourceOwnerId(v int64) *StopTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopTaskRequest) SetTaskId(v int64) *StopTaskRequest {
	s.TaskId = &v
	return s
}

type StopTaskResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopTaskResponseBody) SetRequestId(v string) *StopTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTaskResponseBody) SetData(v bool) *StopTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StopTaskResponseBody) SetCode(v string) *StopTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopTaskResponseBody) SetMessage(v string) *StopTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopTaskResponseBody) SetSuccess(v bool) *StopTaskResponseBody {
	s.Success = &v
	return s
}

type StopTaskResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopTaskResponse) GoString() string {
	return s.String()
}

func (s *StopTaskResponse) SetHeaders(v map[string]*string) *StopTaskResponse {
	s.Headers = v
	return s
}

func (s *StopTaskResponse) SetBody(v *StopTaskResponseBody) *StopTaskResponse {
	s.Body = v
	return s
}

type SuspendHotlineServiceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Type        *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SuspendHotlineServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendHotlineServiceRequest) GoString() string {
	return s.String()
}

func (s *SuspendHotlineServiceRequest) SetClientToken(v string) *SuspendHotlineServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *SuspendHotlineServiceRequest) SetInstanceId(v string) *SuspendHotlineServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *SuspendHotlineServiceRequest) SetAccountName(v string) *SuspendHotlineServiceRequest {
	s.AccountName = &v
	return s
}

func (s *SuspendHotlineServiceRequest) SetType(v int32) *SuspendHotlineServiceRequest {
	s.Type = &v
	return s
}

type SuspendHotlineServiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendHotlineServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendHotlineServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendHotlineServiceResponseBody) SetCode(v string) *SuspendHotlineServiceResponseBody {
	s.Code = &v
	return s
}

func (s *SuspendHotlineServiceResponseBody) SetMessage(v string) *SuspendHotlineServiceResponseBody {
	s.Message = &v
	return s
}

func (s *SuspendHotlineServiceResponseBody) SetRequestId(v string) *SuspendHotlineServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendHotlineServiceResponseBody) SetSuccess(v bool) *SuspendHotlineServiceResponseBody {
	s.Success = &v
	return s
}

type SuspendHotlineServiceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SuspendHotlineServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SuspendHotlineServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendHotlineServiceResponse) GoString() string {
	return s.String()
}

func (s *SuspendHotlineServiceResponse) SetHeaders(v map[string]*string) *SuspendHotlineServiceResponse {
	s.Headers = v
	return s
}

func (s *SuspendHotlineServiceResponse) SetBody(v *SuspendHotlineServiceResponseBody) *SuspendHotlineServiceResponse {
	s.Body = v
	return s
}

type SuspendOutboundTaskRequest struct {
	OutboundTaskId *int64  `json:"OutboundTaskId,omitempty" xml:"OutboundTaskId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SuspendOutboundTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *SuspendOutboundTaskRequest) SetOutboundTaskId(v int64) *SuspendOutboundTaskRequest {
	s.OutboundTaskId = &v
	return s
}

func (s *SuspendOutboundTaskRequest) SetInstanceId(v string) *SuspendOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

type SuspendOutboundTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendOutboundTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendOutboundTaskResponseBody) SetCode(v string) *SuspendOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SuspendOutboundTaskResponseBody) SetMessage(v string) *SuspendOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SuspendOutboundTaskResponseBody) SetData(v string) *SuspendOutboundTaskResponseBody {
	s.Data = &v
	return s
}

func (s *SuspendOutboundTaskResponseBody) SetRequestId(v string) *SuspendOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendOutboundTaskResponseBody) SetSuccess(v bool) *SuspendOutboundTaskResponseBody {
	s.Success = &v
	return s
}

type SuspendOutboundTaskResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SuspendOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SuspendOutboundTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *SuspendOutboundTaskResponse) SetHeaders(v map[string]*string) *SuspendOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *SuspendOutboundTaskResponse) SetBody(v *SuspendOutboundTaskResponseBody) *SuspendOutboundTaskResponse {
	s.Body = v
	return s
}

type TransferCallToSkillGroupRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	SkillGroupId     *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	CallId           *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId     *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	IsSingleTransfer *bool   `json:"IsSingleTransfer,omitempty" xml:"IsSingleTransfer,omitempty"`
}

func (s TransferCallToSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *TransferCallToSkillGroupRequest) SetClientToken(v string) *TransferCallToSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetInstanceId(v string) *TransferCallToSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetAccountName(v string) *TransferCallToSkillGroupRequest {
	s.AccountName = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetSkillGroupId(v int64) *TransferCallToSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetCallId(v string) *TransferCallToSkillGroupRequest {
	s.CallId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetJobId(v string) *TransferCallToSkillGroupRequest {
	s.JobId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetConnectionId(v string) *TransferCallToSkillGroupRequest {
	s.ConnectionId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetHoldConnectionId(v string) *TransferCallToSkillGroupRequest {
	s.HoldConnectionId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetType(v int32) *TransferCallToSkillGroupRequest {
	s.Type = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetIsSingleTransfer(v bool) *TransferCallToSkillGroupRequest {
	s.IsSingleTransfer = &v
	return s
}

type TransferCallToSkillGroupResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferCallToSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *TransferCallToSkillGroupResponseBody) SetCode(v string) *TransferCallToSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *TransferCallToSkillGroupResponseBody) SetMessage(v string) *TransferCallToSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *TransferCallToSkillGroupResponseBody) SetRequestId(v string) *TransferCallToSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferCallToSkillGroupResponseBody) SetSuccess(v bool) *TransferCallToSkillGroupResponseBody {
	s.Success = &v
	return s
}

type TransferCallToSkillGroupResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransferCallToSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferCallToSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *TransferCallToSkillGroupResponse) SetHeaders(v map[string]*string) *TransferCallToSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *TransferCallToSkillGroupResponse) SetBody(v *TransferCallToSkillGroupResponseBody) *TransferCallToSkillGroupResponse {
	s.Body = v
	return s
}

type UpdateAgentRequest struct {
	ClientToken      *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId       *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName      *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName      *string  `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	SkillGroupId     []*int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty" type:"Repeated"`
	SkillGroupIdList []*int64 `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty" type:"Repeated"`
}

func (s UpdateAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequest) SetClientToken(v string) *UpdateAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAgentRequest) SetInstanceId(v string) *UpdateAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAgentRequest) SetAccountName(v string) *UpdateAgentRequest {
	s.AccountName = &v
	return s
}

func (s *UpdateAgentRequest) SetDisplayName(v string) *UpdateAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateAgentRequest) SetSkillGroupId(v []*int64) *UpdateAgentRequest {
	s.SkillGroupId = v
	return s
}

func (s *UpdateAgentRequest) SetSkillGroupIdList(v []*int64) *UpdateAgentRequest {
	s.SkillGroupIdList = v
	return s
}

type UpdateAgentResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s UpdateAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponseBody) SetMessage(v string) *UpdateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgentResponseBody) SetRequestId(v string) *UpdateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentResponseBody) SetCode(v string) *UpdateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgentResponseBody) SetSuccess(v bool) *UpdateAgentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAgentResponseBody) SetHttpStatusCode(v int64) *UpdateAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

type UpdateAgentResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponse) SetHeaders(v map[string]*string) *UpdateAgentResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentResponse) SetBody(v *UpdateAgentResponseBody) *UpdateAgentResponse {
	s.Body = v
	return s
}

type UpdateDepartmentRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 部门名称
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// 部门id
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
}

func (s UpdateDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDepartmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentRequest) SetInstanceId(v string) *UpdateDepartmentRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDepartmentRequest) SetDepartmentName(v string) *UpdateDepartmentRequest {
	s.DepartmentName = &v
	return s
}

func (s *UpdateDepartmentRequest) SetDepartmentId(v int64) *UpdateDepartmentRequest {
	s.DepartmentId = &v
	return s
}

type UpdateDepartmentResponseBody struct {
	// Id of the request
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s UpdateDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentResponseBody) SetRequestId(v string) *UpdateDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetSuccess(v bool) *UpdateDepartmentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetCode(v string) *UpdateDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetMessage(v string) *UpdateDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetHttpStatusCode(v int64) *UpdateDepartmentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetData(v bool) *UpdateDepartmentResponseBody {
	s.Data = &v
	return s
}

type UpdateDepartmentResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDepartmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentResponse) SetHeaders(v map[string]*string) *UpdateDepartmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateDepartmentResponse) SetBody(v *UpdateDepartmentResponseBody) *UpdateDepartmentResponse {
	s.Body = v
	return s
}

type UpdateOuterAccountRequest struct {
	OuterAccountId      *string `json:"OuterAccountId,omitempty" xml:"OuterAccountId,omitempty"`
	OuterAccountType    *string `json:"OuterAccountType,omitempty" xml:"OuterAccountType,omitempty"`
	OuterAccountName    *string `json:"OuterAccountName,omitempty" xml:"OuterAccountName,omitempty"`
	Avatar              *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	RealName            *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	OuterDepartmentId   *string `json:"OuterDepartmentId,omitempty" xml:"OuterDepartmentId,omitempty"`
	OuterGroupIds       *string `json:"OuterGroupIds,omitempty" xml:"OuterGroupIds,omitempty"`
	Ext                 *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	OuterGroupType      *string `json:"OuterGroupType,omitempty" xml:"OuterGroupType,omitempty"`
	OuterDepartmentType *string `json:"OuterDepartmentType,omitempty" xml:"OuterDepartmentType,omitempty"`
}

func (s UpdateOuterAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOuterAccountRequest) GoString() string {
	return s.String()
}

func (s *UpdateOuterAccountRequest) SetOuterAccountId(v string) *UpdateOuterAccountRequest {
	s.OuterAccountId = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetOuterAccountType(v string) *UpdateOuterAccountRequest {
	s.OuterAccountType = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetOuterAccountName(v string) *UpdateOuterAccountRequest {
	s.OuterAccountName = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetAvatar(v string) *UpdateOuterAccountRequest {
	s.Avatar = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetRealName(v string) *UpdateOuterAccountRequest {
	s.RealName = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetOuterDepartmentId(v string) *UpdateOuterAccountRequest {
	s.OuterDepartmentId = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetOuterGroupIds(v string) *UpdateOuterAccountRequest {
	s.OuterGroupIds = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetExt(v string) *UpdateOuterAccountRequest {
	s.Ext = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetOuterGroupType(v string) *UpdateOuterAccountRequest {
	s.OuterGroupType = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetOuterDepartmentType(v string) *UpdateOuterAccountRequest {
	s.OuterDepartmentType = &v
	return s
}

type UpdateOuterAccountResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateOuterAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOuterAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOuterAccountResponseBody) SetMessage(v string) *UpdateOuterAccountResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateOuterAccountResponseBody) SetRequestId(v string) *UpdateOuterAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOuterAccountResponseBody) SetData(v string) *UpdateOuterAccountResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateOuterAccountResponseBody) SetCode(v string) *UpdateOuterAccountResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateOuterAccountResponseBody) SetSuccess(v bool) *UpdateOuterAccountResponseBody {
	s.Success = &v
	return s
}

type UpdateOuterAccountResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOuterAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOuterAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOuterAccountResponse) GoString() string {
	return s.String()
}

func (s *UpdateOuterAccountResponse) SetHeaders(v map[string]*string) *UpdateOuterAccountResponse {
	s.Headers = v
	return s
}

func (s *UpdateOuterAccountResponse) SetBody(v *UpdateOuterAccountResponseBody) *UpdateOuterAccountResponse {
	s.Body = v
	return s
}

type UpdateSkillGroupRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId   *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdateSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupRequest) SetInstanceId(v string) *UpdateSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetSkillGroupId(v int64) *UpdateSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetSkillGroupName(v string) *UpdateSkillGroupRequest {
	s.SkillGroupName = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetDescription(v string) *UpdateSkillGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetDisplayName(v string) *UpdateSkillGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetClientToken(v string) *UpdateSkillGroupRequest {
	s.ClientToken = &v
	return s
}

type UpdateSkillGroupResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupResponseBody) SetCode(v string) *UpdateSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) SetMessage(v string) *UpdateSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) SetRequestId(v string) *UpdateSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) SetSuccess(v bool) *UpdateSkillGroupResponseBody {
	s.Success = &v
	return s
}

type UpdateSkillGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupResponse) SetHeaders(v map[string]*string) *UpdateSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateSkillGroupResponse) SetBody(v *UpdateSkillGroupResponseBody) *UpdateSkillGroupResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("aiccs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddHotlineNumberWithOptions(tmpReq *AddHotlineNumberRequest, runtime *util.RuntimeOptions) (_result *AddHotlineNumberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddHotlineNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OutboundRangeList)) {
		request.OutboundRangeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutboundRangeList, tea.String("OutboundRangeList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddHotlineNumberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddHotlineNumber"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddHotlineNumber(request *AddHotlineNumberRequest) (_result *AddHotlineNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddHotlineNumberResponse{}
	_body, _err := client.AddHotlineNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddOuterAccountWithOptions(request *AddOuterAccountRequest, runtime *util.RuntimeOptions) (_result *AddOuterAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &AddOuterAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddOuterAccount"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddOuterAccount(request *AddOuterAccountRequest) (_result *AddOuterAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddOuterAccountResponse{}
	_body, _err := client.AddOuterAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSkillGroupWithOptions(request *AddSkillGroupRequest, runtime *util.RuntimeOptions) (_result *AddSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &AddSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddSkillGroup"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSkillGroup(request *AddSkillGroupRequest) (_result *AddSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSkillGroupResponse{}
	_body, _err := client.AddSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AiccsSmartCallWithOptions(request *AiccsSmartCallRequest, runtime *util.RuntimeOptions) (_result *AiccsSmartCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AiccsSmartCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AiccsSmartCall"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AiccsSmartCall(request *AiccsSmartCallRequest) (_result *AiccsSmartCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AiccsSmartCallResponse{}
	_body, _err := client.AiccsSmartCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AiccsSmartCallOperateWithOptions(request *AiccsSmartCallOperateRequest, runtime *util.RuntimeOptions) (_result *AiccsSmartCallOperateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AiccsSmartCallOperateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AiccsSmartCallOperate"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AiccsSmartCallOperate(request *AiccsSmartCallOperateRequest) (_result *AiccsSmartCallOperateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AiccsSmartCallOperateResponse{}
	_body, _err := client.AiccsSmartCallOperateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AnswerCallWithOptions(request *AnswerCallRequest, runtime *util.RuntimeOptions) (_result *AnswerCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AnswerCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AnswerCall"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AnswerCall(request *AnswerCallRequest) (_result *AnswerCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AnswerCallResponse{}
	_body, _err := client.AnswerCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachTaskWithOptions(request *AttachTaskRequest, runtime *util.RuntimeOptions) (_result *AttachTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachTask"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachTask(request *AttachTaskRequest) (_result *AttachTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachTaskResponse{}
	_body, _err := client.AttachTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchCreateQualityProjectsWithOptions(request *BatchCreateQualityProjectsRequest, runtime *util.RuntimeOptions) (_result *BatchCreateQualityProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchCreateQualityProjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchCreateQualityProjects"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchCreateQualityProjects(request *BatchCreateQualityProjectsRequest) (_result *BatchCreateQualityProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchCreateQualityProjectsResponse{}
	_body, _err := client.BatchCreateQualityProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelTaskWithOptions(request *CancelTaskRequest, runtime *util.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelTask"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelTask(request *CancelTaskRequest) (_result *CancelTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelTaskResponse{}
	_body, _err := client.CancelTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeChatAgentStatusWithOptions(request *ChangeChatAgentStatusRequest, runtime *util.RuntimeOptions) (_result *ChangeChatAgentStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChangeChatAgentStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ChangeChatAgentStatus"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeChatAgentStatus(request *ChangeChatAgentStatusRequest) (_result *ChangeChatAgentStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeChatAgentStatusResponse{}
	_body, _err := client.ChangeChatAgentStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeQualityProjectStatusWithOptions(request *ChangeQualityProjectStatusRequest, runtime *util.RuntimeOptions) (_result *ChangeQualityProjectStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChangeQualityProjectStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ChangeQualityProjectStatus"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeQualityProjectStatus(request *ChangeQualityProjectStatusRequest) (_result *ChangeQualityProjectStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeQualityProjectStatusResponse{}
	_body, _err := client.ChangeQualityProjectStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAgentWithOptions(request *CreateAgentRequest, runtime *util.RuntimeOptions) (_result *CreateAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAgentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAgent"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAgent(request *CreateAgentRequest) (_result *CreateAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAgentResponse{}
	_body, _err := client.CreateAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDepartmentWithOptions(request *CreateDepartmentRequest, runtime *util.RuntimeOptions) (_result *CreateDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDepartmentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDepartment"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDepartment(request *CreateDepartmentRequest) (_result *CreateDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDepartmentResponse{}
	_body, _err := client.CreateDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOutboundTaskWithOptions(request *CreateOutboundTaskRequest, runtime *util.RuntimeOptions) (_result *CreateOutboundTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateOutboundTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateOutboundTask"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOutboundTask(request *CreateOutboundTaskRequest) (_result *CreateOutboundTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOutboundTaskResponse{}
	_body, _err := client.CreateOutboundTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateQualityProjectWithOptions(request *CreateQualityProjectRequest, runtime *util.RuntimeOptions) (_result *CreateQualityProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateQualityProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateQualityProject"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateQualityProject(request *CreateQualityProjectRequest) (_result *CreateQualityProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateQualityProjectResponse{}
	_body, _err := client.CreateQualityProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateQualityRuleWithOptions(request *CreateQualityRuleRequest, runtime *util.RuntimeOptions) (_result *CreateQualityRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateQualityRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateQualityRule"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateQualityRule(request *CreateQualityRuleRequest) (_result *CreateQualityRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateQualityRuleResponse{}
	_body, _err := client.CreateQualityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSkillGroupWithOptions(request *CreateSkillGroupRequest, runtime *util.RuntimeOptions) (_result *CreateSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSkillGroup"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSkillGroup(request *CreateSkillGroupRequest) (_result *CreateSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSkillGroupResponse{}
	_body, _err := client.CreateSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTaskWithOptions(request *CreateTaskRequest, runtime *util.RuntimeOptions) (_result *CreateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTask"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTask(request *CreateTaskRequest) (_result *CreateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTaskResponse{}
	_body, _err := client.CreateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateThirdSsoAgentWithOptions(request *CreateThirdSsoAgentRequest, runtime *util.RuntimeOptions) (_result *CreateThirdSsoAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateThirdSsoAgentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateThirdSsoAgent"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateThirdSsoAgent(request *CreateThirdSsoAgentRequest) (_result *CreateThirdSsoAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateThirdSsoAgentResponse{}
	_body, _err := client.CreateThirdSsoAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAgentWithOptions(request *DeleteAgentRequest, runtime *util.RuntimeOptions) (_result *DeleteAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteAgentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAgent"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAgent(request *DeleteAgentRequest) (_result *DeleteAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAgentResponse{}
	_body, _err := client.DeleteAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDepartmentWithOptions(request *DeleteDepartmentRequest, runtime *util.RuntimeOptions) (_result *DeleteDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDepartmentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDepartment"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDepartment(request *DeleteDepartmentRequest) (_result *DeleteDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDepartmentResponse{}
	_body, _err := client.DeleteDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHotlineNumberWithOptions(request *DeleteHotlineNumberRequest, runtime *util.RuntimeOptions) (_result *DeleteHotlineNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteHotlineNumberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteHotlineNumber"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHotlineNumber(request *DeleteHotlineNumberRequest) (_result *DeleteHotlineNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHotlineNumberResponse{}
	_body, _err := client.DeleteHotlineNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOutboundTaskWithOptions(request *DeleteOutboundTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteOutboundTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteOutboundTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteOutboundTask"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOutboundTask(request *DeleteOutboundTaskRequest) (_result *DeleteOutboundTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteOutboundTaskResponse{}
	_body, _err := client.DeleteOutboundTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOuterAccountWithOptions(request *DeleteOuterAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteOuterAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteOuterAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteOuterAccount"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOuterAccount(request *DeleteOuterAccountRequest) (_result *DeleteOuterAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteOuterAccountResponse{}
	_body, _err := client.DeleteOuterAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteQualityProjectWithOptions(request *DeleteQualityProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteQualityProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteQualityProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteQualityProject"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteQualityProject(request *DeleteQualityProjectRequest) (_result *DeleteQualityProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteQualityProjectResponse{}
	_body, _err := client.DeleteQualityProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteQualityRuleWithOptions(request *DeleteQualityRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteQualityRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteQualityRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteQualityRule"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteQualityRule(request *DeleteQualityRuleRequest) (_result *DeleteQualityRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteQualityRuleResponse{}
	_body, _err := client.DeleteQualityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSkillGroupWithOptions(request *DeleteSkillGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSkillGroup"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSkillGroup(request *DeleteSkillGroupRequest) (_result *DeleteSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSkillGroupResponse{}
	_body, _err := client.DeleteSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecordDataWithOptions(request *DescribeRecordDataRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRecordDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRecordData"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecordData(request *DescribeRecordDataRequest) (_result *DescribeRecordDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecordDataResponse{}
	_body, _err := client.DescribeRecordDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditQualityProjectWithOptions(request *EditQualityProjectRequest, runtime *util.RuntimeOptions) (_result *EditQualityProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EditQualityProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EditQualityProject"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditQualityProject(request *EditQualityProjectRequest) (_result *EditQualityProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditQualityProjectResponse{}
	_body, _err := client.EditQualityProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditQualityRuleWithOptions(request *EditQualityRuleRequest, runtime *util.RuntimeOptions) (_result *EditQualityRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EditQualityRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EditQualityRule"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditQualityRule(request *EditQualityRuleRequest) (_result *EditQualityRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditQualityRuleResponse{}
	_body, _err := client.EditQualityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditQualityRuleTagWithOptions(request *EditQualityRuleTagRequest, runtime *util.RuntimeOptions) (_result *EditQualityRuleTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EditQualityRuleTagResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EditQualityRuleTag"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditQualityRuleTag(request *EditQualityRuleTagRequest) (_result *EditQualityRuleTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditQualityRuleTagResponse{}
	_body, _err := client.EditQualityRuleTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EncryptPhoneNumWithOptions(request *EncryptPhoneNumRequest, runtime *util.RuntimeOptions) (_result *EncryptPhoneNumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &EncryptPhoneNumResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EncryptPhoneNum"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EncryptPhoneNum(request *EncryptPhoneNumRequest) (_result *EncryptPhoneNumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EncryptPhoneNumResponse{}
	_body, _err := client.EncryptPhoneNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FetchCallWithOptions(request *FetchCallRequest, runtime *util.RuntimeOptions) (_result *FetchCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FetchCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FetchCall"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FetchCall(request *FetchCallRequest) (_result *FetchCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FetchCallResponse{}
	_body, _err := client.FetchCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FinishHotlineServiceWithOptions(request *FinishHotlineServiceRequest, runtime *util.RuntimeOptions) (_result *FinishHotlineServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FinishHotlineServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FinishHotlineService"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FinishHotlineService(request *FinishHotlineServiceRequest) (_result *FinishHotlineServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FinishHotlineServiceResponse{}
	_body, _err := client.FinishHotlineServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateWebSocketSignWithOptions(request *GenerateWebSocketSignRequest, runtime *util.RuntimeOptions) (_result *GenerateWebSocketSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateWebSocketSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateWebSocketSign"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateWebSocketSign(request *GenerateWebSocketSignRequest) (_result *GenerateWebSocketSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateWebSocketSignResponse{}
	_body, _err := client.GenerateWebSocketSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentWithOptions(request *GetAgentRequest, runtime *util.RuntimeOptions) (_result *GetAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAgentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAgent"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgent(request *GetAgentRequest) (_result *GetAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentResponse{}
	_body, _err := client.GetAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentBasisStatusWithOptions(tmpReq *GetAgentBasisStatusRequest, runtime *util.RuntimeOptions) (_result *GetAgentBasisStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetAgentBasisStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AgentIds)) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, tea.String("AgentIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAgentBasisStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAgentBasisStatus"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgentBasisStatus(request *GetAgentBasisStatusRequest) (_result *GetAgentBasisStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentBasisStatusResponse{}
	_body, _err := client.GetAgentBasisStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentByIdWithOptions(request *GetAgentByIdRequest, runtime *util.RuntimeOptions) (_result *GetAgentByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAgentByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAgentById"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgentById(request *GetAgentByIdRequest) (_result *GetAgentByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentByIdResponse{}
	_body, _err := client.GetAgentByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentDetailReportWithOptions(tmpReq *GetAgentDetailReportRequest, runtime *util.RuntimeOptions) (_result *GetAgentDetailReportResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetAgentDetailReportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AgentIds)) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, tea.String("AgentIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAgentDetailReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAgentDetailReport"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgentDetailReport(request *GetAgentDetailReportRequest) (_result *GetAgentDetailReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentDetailReportResponse{}
	_body, _err := client.GetAgentDetailReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentIndexRealTimeWithOptions(request *GetAgentIndexRealTimeRequest, runtime *util.RuntimeOptions) (_result *GetAgentIndexRealTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAgentIndexRealTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAgentIndexRealTime"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgentIndexRealTime(request *GetAgentIndexRealTimeRequest) (_result *GetAgentIndexRealTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentIndexRealTimeResponse{}
	_body, _err := client.GetAgentIndexRealTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentServiceStatusWithOptions(tmpReq *GetAgentServiceStatusRequest, runtime *util.RuntimeOptions) (_result *GetAgentServiceStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetAgentServiceStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AgentIds)) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, tea.String("AgentIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAgentServiceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAgentServiceStatus"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgentServiceStatus(request *GetAgentServiceStatusRequest) (_result *GetAgentServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentServiceStatusResponse{}
	_body, _err := client.GetAgentServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentStatisticsWithOptions(tmpReq *GetAgentStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetAgentStatisticsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetAgentStatisticsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AgentIds)) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, tea.String("AgentIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAgentStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAgentStatistics"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgentStatistics(request *GetAgentStatisticsRequest) (_result *GetAgentStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentStatisticsResponse{}
	_body, _err := client.GetAgentStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllDepartmentWithOptions(request *GetAllDepartmentRequest, runtime *util.RuntimeOptions) (_result *GetAllDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAllDepartmentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAllDepartment"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllDepartment(request *GetAllDepartmentRequest) (_result *GetAllDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAllDepartmentResponse{}
	_body, _err := client.GetAllDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConfigNumListWithOptions(request *GetConfigNumListRequest, runtime *util.RuntimeOptions) (_result *GetConfigNumListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetConfigNumListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetConfigNumList"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConfigNumList(request *GetConfigNumListRequest) (_result *GetConfigNumListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConfigNumListResponse{}
	_body, _err := client.GetConfigNumListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCustomerInfoWithOptions(request *GetCustomerInfoRequest, runtime *util.RuntimeOptions) (_result *GetCustomerInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetCustomerInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCustomerInfo"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCustomerInfo(request *GetCustomerInfoRequest) (_result *GetCustomerInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCustomerInfoResponse{}
	_body, _err := client.GetCustomerInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDepartmentalLatitudeAgentStatusWithOptions(tmpReq *GetDepartmentalLatitudeAgentStatusRequest, runtime *util.RuntimeOptions) (_result *GetDepartmentalLatitudeAgentStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetDepartmentalLatitudeAgentStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDepartmentalLatitudeAgentStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDepartmentalLatitudeAgentStatus"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDepartmentalLatitudeAgentStatus(request *GetDepartmentalLatitudeAgentStatusRequest) (_result *GetDepartmentalLatitudeAgentStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDepartmentalLatitudeAgentStatusResponse{}
	_body, _err := client.GetDepartmentalLatitudeAgentStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDepGroupTreeDataWithOptions(request *GetDepGroupTreeDataRequest, runtime *util.RuntimeOptions) (_result *GetDepGroupTreeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDepGroupTreeDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDepGroupTreeData"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDepGroupTreeData(request *GetDepGroupTreeDataRequest) (_result *GetDepGroupTreeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDepGroupTreeDataResponse{}
	_body, _err := client.GetDepGroupTreeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineAgentDetailWithOptions(request *GetHotlineAgentDetailRequest, runtime *util.RuntimeOptions) (_result *GetHotlineAgentDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetHotlineAgentDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotlineAgentDetail"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineAgentDetail(request *GetHotlineAgentDetailRequest) (_result *GetHotlineAgentDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineAgentDetailResponse{}
	_body, _err := client.GetHotlineAgentDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineAgentDetailReportWithOptions(request *GetHotlineAgentDetailReportRequest, runtime *util.RuntimeOptions) (_result *GetHotlineAgentDetailReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHotlineAgentDetailReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotlineAgentDetailReport"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineAgentDetailReport(request *GetHotlineAgentDetailReportRequest) (_result *GetHotlineAgentDetailReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineAgentDetailReportResponse{}
	_body, _err := client.GetHotlineAgentDetailReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineAgentStatusWithOptions(request *GetHotlineAgentStatusRequest, runtime *util.RuntimeOptions) (_result *GetHotlineAgentStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHotlineAgentStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotlineAgentStatus"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineAgentStatus(request *GetHotlineAgentStatusRequest) (_result *GetHotlineAgentStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineAgentStatusResponse{}
	_body, _err := client.GetHotlineAgentStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineCallActionWithOptions(request *GetHotlineCallActionRequest, runtime *util.RuntimeOptions) (_result *GetHotlineCallActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHotlineCallActionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotlineCallAction"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineCallAction(request *GetHotlineCallActionRequest) (_result *GetHotlineCallActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineCallActionResponse{}
	_body, _err := client.GetHotlineCallActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineGroupDetailReportWithOptions(request *GetHotlineGroupDetailReportRequest, runtime *util.RuntimeOptions) (_result *GetHotlineGroupDetailReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHotlineGroupDetailReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotlineGroupDetailReport"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineGroupDetailReport(request *GetHotlineGroupDetailReportRequest) (_result *GetHotlineGroupDetailReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineGroupDetailReportResponse{}
	_body, _err := client.GetHotlineGroupDetailReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineMessageLogWithOptions(request *GetHotlineMessageLogRequest, runtime *util.RuntimeOptions) (_result *GetHotlineMessageLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetHotlineMessageLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotlineMessageLog"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineMessageLog(request *GetHotlineMessageLogRequest) (_result *GetHotlineMessageLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineMessageLogResponse{}
	_body, _err := client.GetHotlineMessageLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineRuntimeInfoWithOptions(request *GetHotlineRuntimeInfoRequest, runtime *util.RuntimeOptions) (_result *GetHotlineRuntimeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetHotlineRuntimeInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotlineRuntimeInfo"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineRuntimeInfo(request *GetHotlineRuntimeInfoRequest) (_result *GetHotlineRuntimeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineRuntimeInfoResponse{}
	_body, _err := client.GetHotlineRuntimeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineServiceStatisticsWithOptions(tmpReq *GetHotlineServiceStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetHotlineServiceStatisticsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetHotlineServiceStatisticsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AgentIds)) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, tea.String("AgentIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.GroupIds)) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, tea.String("GroupIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetHotlineServiceStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotlineServiceStatistics"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineServiceStatistics(request *GetHotlineServiceStatisticsRequest) (_result *GetHotlineServiceStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineServiceStatisticsResponse{}
	_body, _err := client.GetHotlineServiceStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineWaitingNumberWithOptions(request *GetHotlineWaitingNumberRequest, runtime *util.RuntimeOptions) (_result *GetHotlineWaitingNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetHotlineWaitingNumberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotlineWaitingNumber"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineWaitingNumber(request *GetHotlineWaitingNumberRequest) (_result *GetHotlineWaitingNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineWaitingNumberResponse{}
	_body, _err := client.GetHotlineWaitingNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIndexCurrentValueWithOptions(request *GetIndexCurrentValueRequest, runtime *util.RuntimeOptions) (_result *GetIndexCurrentValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetIndexCurrentValueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetIndexCurrentValue"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIndexCurrentValue(request *GetIndexCurrentValueRequest) (_result *GetIndexCurrentValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIndexCurrentValueResponse{}
	_body, _err := client.GetIndexCurrentValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceListWithOptions(request *GetInstanceListRequest, runtime *util.RuntimeOptions) (_result *GetInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetInstanceListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetInstanceList"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceList(request *GetInstanceListRequest) (_result *GetInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceListResponse{}
	_body, _err := client.GetInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMcuLvsIpWithOptions(request *GetMcuLvsIpRequest, runtime *util.RuntimeOptions) (_result *GetMcuLvsIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetMcuLvsIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMcuLvsIp"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMcuLvsIp(request *GetMcuLvsIpRequest) (_result *GetMcuLvsIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMcuLvsIpResponse{}
	_body, _err := client.GetMcuLvsIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNumLocationWithOptions(request *GetNumLocationRequest, runtime *util.RuntimeOptions) (_result *GetNumLocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetNumLocationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetNumLocation"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNumLocation(request *GetNumLocationRequest) (_result *GetNumLocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNumLocationResponse{}
	_body, _err := client.GetNumLocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOnlineSeatInformationWithOptions(tmpReq *GetOnlineSeatInformationRequest, runtime *util.RuntimeOptions) (_result *GetOnlineSeatInformationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetOnlineSeatInformationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AgentIds)) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, tea.String("AgentIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetOnlineSeatInformationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOnlineSeatInformation"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOnlineSeatInformation(request *GetOnlineSeatInformationRequest) (_result *GetOnlineSeatInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOnlineSeatInformationResponse{}
	_body, _err := client.GetOnlineSeatInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOnlineServiceVolumeWithOptions(tmpReq *GetOnlineServiceVolumeRequest, runtime *util.RuntimeOptions) (_result *GetOnlineServiceVolumeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetOnlineServiceVolumeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AgentIds)) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, tea.String("AgentIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.GroupIds)) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, tea.String("GroupIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetOnlineServiceVolumeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOnlineServiceVolume"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOnlineServiceVolume(request *GetOnlineServiceVolumeRequest) (_result *GetOnlineServiceVolumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOnlineServiceVolumeResponse{}
	_body, _err := client.GetOnlineServiceVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOutbounNumListWithOptions(request *GetOutbounNumListRequest, runtime *util.RuntimeOptions) (_result *GetOutbounNumListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOutbounNumListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOutbounNumList"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOutbounNumList(request *GetOutbounNumListRequest) (_result *GetOutbounNumListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOutbounNumListResponse{}
	_body, _err := client.GetOutbounNumListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQualityProjectDetailWithOptions(request *GetQualityProjectDetailRequest, runtime *util.RuntimeOptions) (_result *GetQualityProjectDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetQualityProjectDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetQualityProjectDetail"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQualityProjectDetail(request *GetQualityProjectDetailRequest) (_result *GetQualityProjectDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQualityProjectDetailResponse{}
	_body, _err := client.GetQualityProjectDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQualityProjectListWithOptions(request *GetQualityProjectListRequest, runtime *util.RuntimeOptions) (_result *GetQualityProjectListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetQualityProjectListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetQualityProjectList"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQualityProjectList(request *GetQualityProjectListRequest) (_result *GetQualityProjectListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQualityProjectListResponse{}
	_body, _err := client.GetQualityProjectListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQualityProjectLogWithOptions(request *GetQualityProjectLogRequest, runtime *util.RuntimeOptions) (_result *GetQualityProjectLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetQualityProjectLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetQualityProjectLog"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQualityProjectLog(request *GetQualityProjectLogRequest) (_result *GetQualityProjectLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQualityProjectLogResponse{}
	_body, _err := client.GetQualityProjectLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQualityResultWithOptions(request *GetQualityResultRequest, runtime *util.RuntimeOptions) (_result *GetQualityResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetQualityResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetQualityResult"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQualityResult(request *GetQualityResultRequest) (_result *GetQualityResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQualityResultResponse{}
	_body, _err := client.GetQualityResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQualityRuleDetailWithOptions(request *GetQualityRuleDetailRequest, runtime *util.RuntimeOptions) (_result *GetQualityRuleDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetQualityRuleDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetQualityRuleDetail"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQualityRuleDetail(request *GetQualityRuleDetailRequest) (_result *GetQualityRuleDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQualityRuleDetailResponse{}
	_body, _err := client.GetQualityRuleDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQualityRuleListWithOptions(request *GetQualityRuleListRequest, runtime *util.RuntimeOptions) (_result *GetQualityRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetQualityRuleListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetQualityRuleList"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQualityRuleList(request *GetQualityRuleListRequest) (_result *GetQualityRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQualityRuleListResponse{}
	_body, _err := client.GetQualityRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQualityRuleTagListWithOptions(request *GetQualityRuleTagListRequest, runtime *util.RuntimeOptions) (_result *GetQualityRuleTagListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetQualityRuleTagListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetQualityRuleTagList"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQualityRuleTagList(request *GetQualityRuleTagListRequest) (_result *GetQualityRuleTagListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQualityRuleTagListResponse{}
	_body, _err := client.GetQualityRuleTagListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQueueInformationWithOptions(tmpReq *GetQueueInformationRequest, runtime *util.RuntimeOptions) (_result *GetQueueInformationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetQueueInformationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.GroupIds)) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, tea.String("GroupIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetQueueInformationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetQueueInformation"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQueueInformation(request *GetQueueInformationRequest) (_result *GetQueueInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQueueInformationResponse{}
	_body, _err := client.GetQueueInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRecordDataWithOptions(request *GetRecordDataRequest, runtime *util.RuntimeOptions) (_result *GetRecordDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRecordDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRecordData"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRecordData(request *GetRecordDataRequest) (_result *GetRecordDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRecordDataResponse{}
	_body, _err := client.GetRecordDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRtcTokenWithOptions(request *GetRtcTokenRequest, runtime *util.RuntimeOptions) (_result *GetRtcTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetRtcTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRtcToken"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRtcToken(request *GetRtcTokenRequest) (_result *GetRtcTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRtcTokenResponse{}
	_body, _err := client.GetRtcTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSeatInformationWithOptions(tmpReq *GetSeatInformationRequest, runtime *util.RuntimeOptions) (_result *GetSeatInformationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetSeatInformationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("depIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetSeatInformationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSeatInformation"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSeatInformation(request *GetSeatInformationRequest) (_result *GetSeatInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSeatInformationResponse{}
	_body, _err := client.GetSeatInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSkillGroupAgentStatusDetailsWithOptions(tmpReq *GetSkillGroupAgentStatusDetailsRequest, runtime *util.RuntimeOptions) (_result *GetSkillGroupAgentStatusDetailsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetSkillGroupAgentStatusDetailsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.GroupIds)) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, tea.String("GroupIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetSkillGroupAgentStatusDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSkillGroupAgentStatusDetails"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSkillGroupAgentStatusDetails(request *GetSkillGroupAgentStatusDetailsRequest) (_result *GetSkillGroupAgentStatusDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSkillGroupAgentStatusDetailsResponse{}
	_body, _err := client.GetSkillGroupAgentStatusDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSkillGroupAndAgentStatusSummaryWithOptions(tmpReq *GetSkillGroupAndAgentStatusSummaryRequest, runtime *util.RuntimeOptions) (_result *GetSkillGroupAndAgentStatusSummaryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetSkillGroupAndAgentStatusSummaryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.GroupIds)) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, tea.String("GroupIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetSkillGroupAndAgentStatusSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSkillGroupAndAgentStatusSummary"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSkillGroupAndAgentStatusSummary(request *GetSkillGroupAndAgentStatusSummaryRequest) (_result *GetSkillGroupAndAgentStatusSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSkillGroupAndAgentStatusSummaryResponse{}
	_body, _err := client.GetSkillGroupAndAgentStatusSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSkillGroupLatitudeStateWithOptions(tmpReq *GetSkillGroupLatitudeStateRequest, runtime *util.RuntimeOptions) (_result *GetSkillGroupLatitudeStateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetSkillGroupLatitudeStateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.GroupIds)) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, tea.String("GroupIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetSkillGroupLatitudeStateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSkillGroupLatitudeState"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSkillGroupLatitudeState(request *GetSkillGroupLatitudeStateRequest) (_result *GetSkillGroupLatitudeStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSkillGroupLatitudeStateResponse{}
	_body, _err := client.GetSkillGroupLatitudeStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSkillGroupServiceCapabilityWithOptions(tmpReq *GetSkillGroupServiceCapabilityRequest, runtime *util.RuntimeOptions) (_result *GetSkillGroupServiceCapabilityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetSkillGroupServiceCapabilityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.GroupIds)) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, tea.String("GroupIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetSkillGroupServiceCapabilityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSkillGroupServiceCapability"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSkillGroupServiceCapability(request *GetSkillGroupServiceCapabilityRequest) (_result *GetSkillGroupServiceCapabilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSkillGroupServiceCapabilityResponse{}
	_body, _err := client.GetSkillGroupServiceCapabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSkillGroupServiceStatusWithOptions(tmpReq *GetSkillGroupServiceStatusRequest, runtime *util.RuntimeOptions) (_result *GetSkillGroupServiceStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetSkillGroupServiceStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AgentIds)) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, tea.String("AgentIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.GroupIds)) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, tea.String("GroupIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetSkillGroupServiceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSkillGroupServiceStatus"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSkillGroupServiceStatus(request *GetSkillGroupServiceStatusRequest) (_result *GetSkillGroupServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSkillGroupServiceStatusResponse{}
	_body, _err := client.GetSkillGroupServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSkillGroupStatusTotalWithOptions(tmpReq *GetSkillGroupStatusTotalRequest, runtime *util.RuntimeOptions) (_result *GetSkillGroupStatusTotalResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetSkillGroupStatusTotalShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AgentIds)) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, tea.String("AgentIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DepIds)) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, tea.String("DepIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.GroupIds)) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, tea.String("GroupIds"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetSkillGroupStatusTotalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSkillGroupStatusTotal"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSkillGroupStatusTotal(request *GetSkillGroupStatusTotalRequest) (_result *GetSkillGroupStatusTotalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSkillGroupStatusTotalResponse{}
	_body, _err := client.GetSkillGroupStatusTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HangupCallWithOptions(request *HangupCallRequest, runtime *util.RuntimeOptions) (_result *HangupCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &HangupCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("HangupCall"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HangupCall(request *HangupCallRequest) (_result *HangupCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HangupCallResponse{}
	_body, _err := client.HangupCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HangUpDoubleCallWithOptions(request *HangUpDoubleCallRequest, runtime *util.RuntimeOptions) (_result *HangUpDoubleCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &HangUpDoubleCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("HangUpDoubleCall"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HangUpDoubleCall(request *HangUpDoubleCallRequest) (_result *HangUpDoubleCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HangUpDoubleCallResponse{}
	_body, _err := client.HangUpDoubleCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HangupThirdCallWithOptions(request *HangupThirdCallRequest, runtime *util.RuntimeOptions) (_result *HangupThirdCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &HangupThirdCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("HangupThirdCall"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HangupThirdCall(request *HangupThirdCallRequest) (_result *HangupThirdCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HangupThirdCallResponse{}
	_body, _err := client.HangupThirdCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HoldCallWithOptions(request *HoldCallRequest, runtime *util.RuntimeOptions) (_result *HoldCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &HoldCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("HoldCall"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HoldCall(request *HoldCallRequest) (_result *HoldCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HoldCallResponse{}
	_body, _err := client.HoldCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HotlineSessionQueryWithOptions(request *HotlineSessionQueryRequest, runtime *util.RuntimeOptions) (_result *HotlineSessionQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &HotlineSessionQueryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("HotlineSessionQuery"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HotlineSessionQuery(request *HotlineSessionQueryRequest) (_result *HotlineSessionQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HotlineSessionQueryResponse{}
	_body, _err := client.HotlineSessionQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertTaskDetailWithOptions(request *InsertTaskDetailRequest, runtime *util.RuntimeOptions) (_result *InsertTaskDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InsertTaskDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InsertTaskDetail"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertTaskDetail(request *InsertTaskDetailRequest) (_result *InsertTaskDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertTaskDetailResponse{}
	_body, _err := client.InsertTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinThirdCallWithOptions(request *JoinThirdCallRequest, runtime *util.RuntimeOptions) (_result *JoinThirdCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &JoinThirdCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("JoinThirdCall"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinThirdCall(request *JoinThirdCallRequest) (_result *JoinThirdCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinThirdCallResponse{}
	_body, _err := client.JoinThirdCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAgentBySkillGroupIdWithOptions(request *ListAgentBySkillGroupIdRequest, runtime *util.RuntimeOptions) (_result *ListAgentBySkillGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListAgentBySkillGroupIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAgentBySkillGroupId"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAgentBySkillGroupId(request *ListAgentBySkillGroupIdRequest) (_result *ListAgentBySkillGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAgentBySkillGroupIdResponse{}
	_body, _err := client.ListAgentBySkillGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAiccsRobotWithOptions(request *ListAiccsRobotRequest, runtime *util.RuntimeOptions) (_result *ListAiccsRobotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAiccsRobotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAiccsRobot"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAiccsRobot(request *ListAiccsRobotRequest) (_result *ListAiccsRobotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAiccsRobotResponse{}
	_body, _err := client.ListAiccsRobotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListChatRecordDetailWithOptions(request *ListChatRecordDetailRequest, runtime *util.RuntimeOptions) (_result *ListChatRecordDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListChatRecordDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListChatRecordDetail"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListChatRecordDetail(request *ListChatRecordDetailRequest) (_result *ListChatRecordDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListChatRecordDetailResponse{}
	_body, _err := client.ListChatRecordDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDialogWithOptions(request *ListDialogRequest, runtime *util.RuntimeOptions) (_result *ListDialogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDialogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDialog"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDialog(request *ListDialogRequest) (_result *ListDialogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDialogResponse{}
	_body, _err := client.ListDialogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotlineRecordWithOptions(request *ListHotlineRecordRequest, runtime *util.RuntimeOptions) (_result *ListHotlineRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListHotlineRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHotlineRecord"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotlineRecord(request *ListHotlineRecordRequest) (_result *ListHotlineRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHotlineRecordResponse{}
	_body, _err := client.ListHotlineRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotlineRecordDetailWithOptions(request *ListHotlineRecordDetailRequest, runtime *util.RuntimeOptions) (_result *ListHotlineRecordDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListHotlineRecordDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHotlineRecordDetail"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotlineRecordDetail(request *ListHotlineRecordDetailRequest) (_result *ListHotlineRecordDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHotlineRecordDetailResponse{}
	_body, _err := client.ListHotlineRecordDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOutboundPhoneNumberWithOptions(request *ListOutboundPhoneNumberRequest, runtime *util.RuntimeOptions) (_result *ListOutboundPhoneNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListOutboundPhoneNumberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOutboundPhoneNumber"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOutboundPhoneNumber(request *ListOutboundPhoneNumberRequest) (_result *ListOutboundPhoneNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOutboundPhoneNumberResponse{}
	_body, _err := client.ListOutboundPhoneNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOutboundStrategiesWithOptions(request *ListOutboundStrategiesRequest, runtime *util.RuntimeOptions) (_result *ListOutboundStrategiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOutboundStrategiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOutboundStrategies"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOutboundStrategies(request *ListOutboundStrategiesRequest) (_result *ListOutboundStrategiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOutboundStrategiesResponse{}
	_body, _err := client.ListOutboundStrategiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOuterOrderedNumbersWithOptions(request *ListOuterOrderedNumbersRequest, runtime *util.RuntimeOptions) (_result *ListOuterOrderedNumbersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOuterOrderedNumbersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOuterOrderedNumbers"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOuterOrderedNumbers(request *ListOuterOrderedNumbersRequest) (_result *ListOuterOrderedNumbersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOuterOrderedNumbersResponse{}
	_body, _err := client.ListOuterOrderedNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRobotCallDialogWithOptions(request *ListRobotCallDialogRequest, runtime *util.RuntimeOptions) (_result *ListRobotCallDialogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRobotCallDialogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRobotCallDialog"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRobotCallDialog(request *ListRobotCallDialogRequest) (_result *ListRobotCallDialogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRobotCallDialogResponse{}
	_body, _err := client.ListRobotCallDialogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRolesWithOptions(request *ListRolesRequest, runtime *util.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListRolesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRoles"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRoles(request *ListRolesRequest) (_result *ListRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRolesResponse{}
	_body, _err := client.ListRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSkillGroupWithOptions(request *ListSkillGroupRequest, runtime *util.RuntimeOptions) (_result *ListSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSkillGroup"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSkillGroup(request *ListSkillGroupRequest) (_result *ListSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSkillGroupResponse{}
	_body, _err := client.ListSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTaskWithOptions(request *ListTaskRequest, runtime *util.RuntimeOptions) (_result *ListTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTask"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTask(request *ListTaskRequest) (_result *ListTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTaskResponse{}
	_body, _err := client.ListTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTaskDetailWithOptions(request *ListTaskDetailRequest, runtime *util.RuntimeOptions) (_result *ListTaskDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTaskDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTaskDetail"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTaskDetail(request *ListTaskDetailRequest) (_result *ListTaskDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTaskDetailResponse{}
	_body, _err := client.ListTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MakeCallWithOptions(request *MakeCallRequest, runtime *util.RuntimeOptions) (_result *MakeCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MakeCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MakeCall"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MakeCall(request *MakeCallRequest) (_result *MakeCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MakeCallResponse{}
	_body, _err := client.MakeCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MakeDoubleCallWithOptions(request *MakeDoubleCallRequest, runtime *util.RuntimeOptions) (_result *MakeDoubleCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MakeDoubleCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MakeDoubleCall"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MakeDoubleCall(request *MakeDoubleCallRequest) (_result *MakeDoubleCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MakeDoubleCallResponse{}
	_body, _err := client.MakeDoubleCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHotlineInQueueWithOptions(request *QueryHotlineInQueueRequest, runtime *util.RuntimeOptions) (_result *QueryHotlineInQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &QueryHotlineInQueueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryHotlineInQueue"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHotlineInQueue(request *QueryHotlineInQueueRequest) (_result *QueryHotlineInQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryHotlineInQueueResponse{}
	_body, _err := client.QueryHotlineInQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHotlineNumberWithOptions(tmpReq *QueryHotlineNumberRequest, runtime *util.RuntimeOptions) (_result *QueryHotlineNumberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryHotlineNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.GroupIds)) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, tea.String("GroupIds"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &QueryHotlineNumberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryHotlineNumber"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHotlineNumber(request *QueryHotlineNumberRequest) (_result *QueryHotlineNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryHotlineNumberResponse{}
	_body, _err := client.QueryHotlineNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOutboundTaskWithOptions(request *QueryOutboundTaskRequest, runtime *util.RuntimeOptions) (_result *QueryOutboundTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryOutboundTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryOutboundTask"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOutboundTask(request *QueryOutboundTaskRequest) (_result *QueryOutboundTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOutboundTaskResponse{}
	_body, _err := client.QueryOutboundTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySkillGroupsWithOptions(request *QuerySkillGroupsRequest, runtime *util.RuntimeOptions) (_result *QuerySkillGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySkillGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySkillGroups"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySkillGroups(request *QuerySkillGroupsRequest) (_result *QuerySkillGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySkillGroupsResponse{}
	_body, _err := client.QuerySkillGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTaskDetailWithOptions(request *QueryTaskDetailRequest, runtime *util.RuntimeOptions) (_result *QueryTaskDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTaskDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTaskDetail"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTaskDetail(request *QueryTaskDetailRequest) (_result *QueryTaskDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTaskDetailResponse{}
	_body, _err := client.QueryTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTaskResultWithOptions(request *QueryTaskResultRequest, runtime *util.RuntimeOptions) (_result *QueryTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &QueryTaskResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTaskResult"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTaskResult(request *QueryTaskResultRequest) (_result *QueryTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTaskResultResponse{}
	_body, _err := client.QueryTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTicketsWithOptions(tmpReq *QueryTicketsRequest, runtime *util.RuntimeOptions) (_result *QueryTicketsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryTicketsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extra)) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, tea.String("Extra"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTicketsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTickets"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTickets(request *QueryTicketsRequest) (_result *QueryTicketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTicketsResponse{}
	_body, _err := client.QueryTicketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTouchListWithOptions(request *QueryTouchListRequest, runtime *util.RuntimeOptions) (_result *QueryTouchListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTouchListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTouchList"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTouchList(request *QueryTouchListRequest) (_result *QueryTouchListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTouchListResponse{}
	_body, _err := client.QueryTouchListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveSkillGroupWithOptions(request *RemoveSkillGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveSkillGroup"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSkillGroup(request *RemoveSkillGroupRequest) (_result *RemoveSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveSkillGroupResponse{}
	_body, _err := client.RemoveSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetHotlineNumberWithOptions(tmpReq *ResetHotlineNumberRequest, runtime *util.RuntimeOptions) (_result *ResetHotlineNumberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ResetHotlineNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OutboundRangeList)) {
		request.OutboundRangeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutboundRangeList, tea.String("OutboundRangeList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetHotlineNumberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetHotlineNumber"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetHotlineNumber(request *ResetHotlineNumberRequest) (_result *ResetHotlineNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetHotlineNumberResponse{}
	_body, _err := client.ResetHotlineNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartOutboundTaskWithOptions(request *RestartOutboundTaskRequest, runtime *util.RuntimeOptions) (_result *RestartOutboundTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestartOutboundTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestartOutboundTask"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartOutboundTask(request *RestartOutboundTaskRequest) (_result *RestartOutboundTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartOutboundTaskResponse{}
	_body, _err := client.RestartOutboundTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RobotCallWithOptions(request *RobotCallRequest, runtime *util.RuntimeOptions) (_result *RobotCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RobotCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RobotCall"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RobotCall(request *RobotCallRequest) (_result *RobotCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RobotCallResponse{}
	_body, _err := client.RobotCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendCcoSmartCallWithOptions(request *SendCcoSmartCallRequest, runtime *util.RuntimeOptions) (_result *SendCcoSmartCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendCcoSmartCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendCcoSmartCall"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendCcoSmartCall(request *SendCcoSmartCallRequest) (_result *SendCcoSmartCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCcoSmartCallResponse{}
	_body, _err := client.SendCcoSmartCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendCcoSmartCallOperateWithOptions(request *SendCcoSmartCallOperateRequest, runtime *util.RuntimeOptions) (_result *SendCcoSmartCallOperateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendCcoSmartCallOperateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendCcoSmartCallOperate"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendCcoSmartCallOperate(request *SendCcoSmartCallOperateRequest) (_result *SendCcoSmartCallOperateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCcoSmartCallOperateResponse{}
	_body, _err := client.SendCcoSmartCallOperateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendHotlineHeartBeatWithOptions(request *SendHotlineHeartBeatRequest, runtime *util.RuntimeOptions) (_result *SendHotlineHeartBeatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendHotlineHeartBeatResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendHotlineHeartBeat"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendHotlineHeartBeat(request *SendHotlineHeartBeatRequest) (_result *SendHotlineHeartBeatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendHotlineHeartBeatResponse{}
	_body, _err := client.SendHotlineHeartBeatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCallWithOptions(request *StartCallRequest, runtime *util.RuntimeOptions) (_result *StartCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartCall"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCall(request *StartCallRequest) (_result *StartCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartCallResponse{}
	_body, _err := client.StartCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCallV2WithOptions(request *StartCallV2Request, runtime *util.RuntimeOptions) (_result *StartCallV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartCallV2Response{}
	_body, _err := client.DoRPCRequest(tea.String("StartCallV2"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCallV2(request *StartCallV2Request) (_result *StartCallV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartCallV2Response{}
	_body, _err := client.StartCallV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartChatWorkWithOptions(request *StartChatWorkRequest, runtime *util.RuntimeOptions) (_result *StartChatWorkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartChatWorkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartChatWork"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartChatWork(request *StartChatWorkRequest) (_result *StartChatWorkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartChatWorkResponse{}
	_body, _err := client.StartChatWorkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartHotlineServiceWithOptions(request *StartHotlineServiceRequest, runtime *util.RuntimeOptions) (_result *StartHotlineServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartHotlineServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartHotlineService"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartHotlineService(request *StartHotlineServiceRequest) (_result *StartHotlineServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartHotlineServiceResponse{}
	_body, _err := client.StartHotlineServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartMicroOutboundWithOptions(request *StartMicroOutboundRequest, runtime *util.RuntimeOptions) (_result *StartMicroOutboundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartMicroOutboundResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartMicroOutbound"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartMicroOutbound(request *StartMicroOutboundRequest) (_result *StartMicroOutboundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartMicroOutboundResponse{}
	_body, _err := client.StartMicroOutboundWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartTaskWithOptions(request *StartTaskRequest, runtime *util.RuntimeOptions) (_result *StartTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartTask"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartTask(request *StartTaskRequest) (_result *StartTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartTaskResponse{}
	_body, _err := client.StartTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartTaskByAppWithOptions(request *StartTaskByAppRequest, runtime *util.RuntimeOptions) (_result *StartTaskByAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartTaskByAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartTaskByApp"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartTaskByApp(request *StartTaskByAppRequest) (_result *StartTaskByAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartTaskByAppResponse{}
	_body, _err := client.StartTaskByAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopTaskWithOptions(request *StopTaskRequest, runtime *util.RuntimeOptions) (_result *StopTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopTask"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopTask(request *StopTaskRequest) (_result *StopTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopTaskResponse{}
	_body, _err := client.StopTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SuspendHotlineServiceWithOptions(request *SuspendHotlineServiceRequest, runtime *util.RuntimeOptions) (_result *SuspendHotlineServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SuspendHotlineServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SuspendHotlineService"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuspendHotlineService(request *SuspendHotlineServiceRequest) (_result *SuspendHotlineServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuspendHotlineServiceResponse{}
	_body, _err := client.SuspendHotlineServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SuspendOutboundTaskWithOptions(request *SuspendOutboundTaskRequest, runtime *util.RuntimeOptions) (_result *SuspendOutboundTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SuspendOutboundTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SuspendOutboundTask"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuspendOutboundTask(request *SuspendOutboundTaskRequest) (_result *SuspendOutboundTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuspendOutboundTaskResponse{}
	_body, _err := client.SuspendOutboundTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferCallToSkillGroupWithOptions(request *TransferCallToSkillGroupRequest, runtime *util.RuntimeOptions) (_result *TransferCallToSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TransferCallToSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TransferCallToSkillGroup"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferCallToSkillGroup(request *TransferCallToSkillGroupRequest) (_result *TransferCallToSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferCallToSkillGroupResponse{}
	_body, _err := client.TransferCallToSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAgentWithOptions(request *UpdateAgentRequest, runtime *util.RuntimeOptions) (_result *UpdateAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAgentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAgent"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAgent(request *UpdateAgentRequest) (_result *UpdateAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAgentResponse{}
	_body, _err := client.UpdateAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDepartmentWithOptions(request *UpdateDepartmentRequest, runtime *util.RuntimeOptions) (_result *UpdateDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDepartmentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDepartment"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDepartment(request *UpdateDepartmentRequest) (_result *UpdateDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDepartmentResponse{}
	_body, _err := client.UpdateDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOuterAccountWithOptions(request *UpdateOuterAccountRequest, runtime *util.RuntimeOptions) (_result *UpdateOuterAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &UpdateOuterAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateOuterAccount"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOuterAccount(request *UpdateOuterAccountRequest) (_result *UpdateOuterAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOuterAccountResponse{}
	_body, _err := client.UpdateOuterAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSkillGroupWithOptions(request *UpdateSkillGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSkillGroup"), tea.String("2019-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSkillGroup(request *UpdateSkillGroupRequest) (_result *UpdateSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSkillGroupResponse{}
	_body, _err := client.UpdateSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
