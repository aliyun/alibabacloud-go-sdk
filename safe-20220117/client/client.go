// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ChangeCancelRequest struct {
	AuthKey       *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign      *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	ReqTimestamp  *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	SourceOrderId *string `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
}

func (s ChangeCancelRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeCancelRequest) GoString() string {
	return s.String()
}

func (s *ChangeCancelRequest) SetAuthKey(v string) *ChangeCancelRequest {
	s.AuthKey = &v
	return s
}

func (s *ChangeCancelRequest) SetAuthSign(v string) *ChangeCancelRequest {
	s.AuthSign = &v
	return s
}

func (s *ChangeCancelRequest) SetReqTimestamp(v int64) *ChangeCancelRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *ChangeCancelRequest) SetSourceOrderId(v string) *ChangeCancelRequest {
	s.SourceOrderId = &v
	return s
}

type ChangeCancelResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeCancelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeCancelResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeCancelResponseBody) SetCode(v int32) *ChangeCancelResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeCancelResponseBody) SetData(v string) *ChangeCancelResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeCancelResponseBody) SetMessage(v string) *ChangeCancelResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeCancelResponseBody) SetRequestId(v string) *ChangeCancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeCancelResponseBody) SetSuccess(v bool) *ChangeCancelResponseBody {
	s.Success = &v
	return s
}

type ChangeCancelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeCancelResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeCancelResponse) GoString() string {
	return s.String()
}

func (s *ChangeCancelResponse) SetHeaders(v map[string]*string) *ChangeCancelResponse {
	s.Headers = v
	return s
}

func (s *ChangeCancelResponse) SetStatusCode(v int32) *ChangeCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeCancelResponse) SetBody(v *ChangeCancelResponseBody) *ChangeCancelResponse {
	s.Body = v
	return s
}

type ChangeCheckRequest struct {
	AffectCustomer           *string                                     `json:"AffectCustomer,omitempty" xml:"AffectCustomer,omitempty"`
	ApproveFlowParam         *ChangeCheckRequestApproveFlowParam         `json:"ApproveFlowParam,omitempty" xml:"ApproveFlowParam,omitempty" type:"Struct"`
	AuthKey                  *string                                     `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign                 *string                                     `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	BgCustomTemplateExtraDTO *ChangeCheckRequestBgCustomTemplateExtraDTO `json:"BgCustomTemplateExtraDTO,omitempty" xml:"BgCustomTemplateExtraDTO,omitempty" type:"Struct"`
	BgVid                    *string                                     `json:"BgVid,omitempty" xml:"BgVid,omitempty"`
	BlockInfos               []*ChangeCheckRequestBlockInfos             `json:"BlockInfos,omitempty" xml:"BlockInfos,omitempty" type:"Repeated"`
	CallBackInfo             *ChangeCheckRequestCallBackInfo             `json:"CallBackInfo,omitempty" xml:"CallBackInfo,omitempty" type:"Struct"`
	ChangeDataType           *string                                     `json:"ChangeDataType,omitempty" xml:"ChangeDataType,omitempty"`
	ChangeDesc               *string                                     `json:"ChangeDesc,omitempty" xml:"ChangeDesc,omitempty"`
	ChangeEndTime            *int64                                      `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	ChangeEnv                *string                                     `json:"ChangeEnv,omitempty" xml:"ChangeEnv,omitempty"`
	ChangeItems              *string                                     `json:"ChangeItems,omitempty" xml:"ChangeItems,omitempty"`
	ChangeObject             *string                                     `json:"ChangeObject,omitempty" xml:"ChangeObject,omitempty"`
	ChangeOptSubType         *string                                     `json:"ChangeOptSubType,omitempty" xml:"ChangeOptSubType,omitempty"`
	ChangeOptType            *string                                     `json:"ChangeOptType,omitempty" xml:"ChangeOptType,omitempty"`
	ChangeReason             *string                                     `json:"ChangeReason,omitempty" xml:"ChangeReason,omitempty"`
	ChangeRmarks             *string                                     `json:"ChangeRmarks,omitempty" xml:"ChangeRmarks,omitempty"`
	ChangeSchemes            *string                                     `json:"ChangeSchemes,omitempty" xml:"ChangeSchemes,omitempty"`
	ChangeStartTime          *int64                                      `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
	ChangeSubTypeDesc        *string                                     `json:"ChangeSubTypeDesc,omitempty" xml:"ChangeSubTypeDesc,omitempty"`
	ChangeSystem             *string                                     `json:"ChangeSystem,omitempty" xml:"ChangeSystem,omitempty"`
	ChangeTimes              []*ChangeCheckRequestChangeTimes            `json:"ChangeTimes,omitempty" xml:"ChangeTimes,omitempty" type:"Repeated"`
	ChangeTitle              *string                                     `json:"ChangeTitle,omitempty" xml:"ChangeTitle,omitempty"`
	ChangeValidation         *string                                     `json:"ChangeValidation,omitempty" xml:"ChangeValidation,omitempty"`
	CreatorEmpId             *string                                     `json:"CreatorEmpId,omitempty" xml:"CreatorEmpId,omitempty"`
	DamagedChangeNotices     []*ChangeCheckRequestDamagedChangeNotices   `json:"DamagedChangeNotices,omitempty" xml:"DamagedChangeNotices,omitempty" type:"Repeated"`
	ExecutorEmpId            *string                                     `json:"ExecutorEmpId,omitempty" xml:"ExecutorEmpId,omitempty"`
	ExtraInfo                *string                                     `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Follower                 []*string                                   `json:"Follower,omitempty" xml:"Follower,omitempty" type:"Repeated"`
	GrayStatus               *string                                     `json:"GrayStatus,omitempty" xml:"GrayStatus,omitempty"`
	HarmChangeNoticeEnum     *string                                     `json:"HarmChangeNoticeEnum,omitempty" xml:"HarmChangeNoticeEnum,omitempty"`
	Incidence                *string                                     `json:"Incidence,omitempty" xml:"Incidence,omitempty"`
	InfluenceInfo            *ChangeCheckRequestInfluenceInfo            `json:"InfluenceInfo,omitempty" xml:"InfluenceInfo,omitempty" type:"Struct"`
	Instance                 *ChangeCheckRequestInstance                 `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	NeedModifyDoc            *string                                     `json:"NeedModifyDoc,omitempty" xml:"NeedModifyDoc,omitempty"`
	Product                  []*ChangeCheckRequestProduct                `json:"Product,omitempty" xml:"Product,omitempty" type:"Repeated"`
	ReleasePackageInfos      []*ChangeCheckRequestReleasePackageInfos    `json:"ReleasePackageInfos,omitempty" xml:"ReleasePackageInfos,omitempty" type:"Repeated"`
	ReqTimestamp             *int64                                      `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	ReuseSourceOrderId       *string                                     `json:"ReuseSourceOrderId,omitempty" xml:"ReuseSourceOrderId,omitempty"`
	RiskLevel                *string                                     `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Rollback                 *string                                     `json:"Rollback,omitempty" xml:"Rollback,omitempty"`
	SourceName               *string                                     `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	SourceOrderId            *string                                     `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
	SourceUrl                *string                                     `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	WhiteType                *int32                                      `json:"WhiteType,omitempty" xml:"WhiteType,omitempty"`
}

func (s ChangeCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequest) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequest) SetAffectCustomer(v string) *ChangeCheckRequest {
	s.AffectCustomer = &v
	return s
}

func (s *ChangeCheckRequest) SetApproveFlowParam(v *ChangeCheckRequestApproveFlowParam) *ChangeCheckRequest {
	s.ApproveFlowParam = v
	return s
}

func (s *ChangeCheckRequest) SetAuthKey(v string) *ChangeCheckRequest {
	s.AuthKey = &v
	return s
}

func (s *ChangeCheckRequest) SetAuthSign(v string) *ChangeCheckRequest {
	s.AuthSign = &v
	return s
}

func (s *ChangeCheckRequest) SetBgCustomTemplateExtraDTO(v *ChangeCheckRequestBgCustomTemplateExtraDTO) *ChangeCheckRequest {
	s.BgCustomTemplateExtraDTO = v
	return s
}

func (s *ChangeCheckRequest) SetBgVid(v string) *ChangeCheckRequest {
	s.BgVid = &v
	return s
}

func (s *ChangeCheckRequest) SetBlockInfos(v []*ChangeCheckRequestBlockInfos) *ChangeCheckRequest {
	s.BlockInfos = v
	return s
}

func (s *ChangeCheckRequest) SetCallBackInfo(v *ChangeCheckRequestCallBackInfo) *ChangeCheckRequest {
	s.CallBackInfo = v
	return s
}

func (s *ChangeCheckRequest) SetChangeDataType(v string) *ChangeCheckRequest {
	s.ChangeDataType = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeDesc(v string) *ChangeCheckRequest {
	s.ChangeDesc = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeEndTime(v int64) *ChangeCheckRequest {
	s.ChangeEndTime = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeEnv(v string) *ChangeCheckRequest {
	s.ChangeEnv = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeItems(v string) *ChangeCheckRequest {
	s.ChangeItems = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeObject(v string) *ChangeCheckRequest {
	s.ChangeObject = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeOptSubType(v string) *ChangeCheckRequest {
	s.ChangeOptSubType = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeOptType(v string) *ChangeCheckRequest {
	s.ChangeOptType = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeReason(v string) *ChangeCheckRequest {
	s.ChangeReason = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeRmarks(v string) *ChangeCheckRequest {
	s.ChangeRmarks = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeSchemes(v string) *ChangeCheckRequest {
	s.ChangeSchemes = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeStartTime(v int64) *ChangeCheckRequest {
	s.ChangeStartTime = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeSubTypeDesc(v string) *ChangeCheckRequest {
	s.ChangeSubTypeDesc = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeSystem(v string) *ChangeCheckRequest {
	s.ChangeSystem = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeTimes(v []*ChangeCheckRequestChangeTimes) *ChangeCheckRequest {
	s.ChangeTimes = v
	return s
}

func (s *ChangeCheckRequest) SetChangeTitle(v string) *ChangeCheckRequest {
	s.ChangeTitle = &v
	return s
}

func (s *ChangeCheckRequest) SetChangeValidation(v string) *ChangeCheckRequest {
	s.ChangeValidation = &v
	return s
}

func (s *ChangeCheckRequest) SetCreatorEmpId(v string) *ChangeCheckRequest {
	s.CreatorEmpId = &v
	return s
}

func (s *ChangeCheckRequest) SetDamagedChangeNotices(v []*ChangeCheckRequestDamagedChangeNotices) *ChangeCheckRequest {
	s.DamagedChangeNotices = v
	return s
}

func (s *ChangeCheckRequest) SetExecutorEmpId(v string) *ChangeCheckRequest {
	s.ExecutorEmpId = &v
	return s
}

func (s *ChangeCheckRequest) SetExtraInfo(v string) *ChangeCheckRequest {
	s.ExtraInfo = &v
	return s
}

func (s *ChangeCheckRequest) SetFollower(v []*string) *ChangeCheckRequest {
	s.Follower = v
	return s
}

func (s *ChangeCheckRequest) SetGrayStatus(v string) *ChangeCheckRequest {
	s.GrayStatus = &v
	return s
}

func (s *ChangeCheckRequest) SetHarmChangeNoticeEnum(v string) *ChangeCheckRequest {
	s.HarmChangeNoticeEnum = &v
	return s
}

func (s *ChangeCheckRequest) SetIncidence(v string) *ChangeCheckRequest {
	s.Incidence = &v
	return s
}

func (s *ChangeCheckRequest) SetInfluenceInfo(v *ChangeCheckRequestInfluenceInfo) *ChangeCheckRequest {
	s.InfluenceInfo = v
	return s
}

func (s *ChangeCheckRequest) SetInstance(v *ChangeCheckRequestInstance) *ChangeCheckRequest {
	s.Instance = v
	return s
}

func (s *ChangeCheckRequest) SetNeedModifyDoc(v string) *ChangeCheckRequest {
	s.NeedModifyDoc = &v
	return s
}

func (s *ChangeCheckRequest) SetProduct(v []*ChangeCheckRequestProduct) *ChangeCheckRequest {
	s.Product = v
	return s
}

func (s *ChangeCheckRequest) SetReleasePackageInfos(v []*ChangeCheckRequestReleasePackageInfos) *ChangeCheckRequest {
	s.ReleasePackageInfos = v
	return s
}

func (s *ChangeCheckRequest) SetReqTimestamp(v int64) *ChangeCheckRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *ChangeCheckRequest) SetReuseSourceOrderId(v string) *ChangeCheckRequest {
	s.ReuseSourceOrderId = &v
	return s
}

func (s *ChangeCheckRequest) SetRiskLevel(v string) *ChangeCheckRequest {
	s.RiskLevel = &v
	return s
}

func (s *ChangeCheckRequest) SetRollback(v string) *ChangeCheckRequest {
	s.Rollback = &v
	return s
}

func (s *ChangeCheckRequest) SetSourceName(v string) *ChangeCheckRequest {
	s.SourceName = &v
	return s
}

func (s *ChangeCheckRequest) SetSourceOrderId(v string) *ChangeCheckRequest {
	s.SourceOrderId = &v
	return s
}

func (s *ChangeCheckRequest) SetSourceUrl(v string) *ChangeCheckRequest {
	s.SourceUrl = &v
	return s
}

func (s *ChangeCheckRequest) SetWhiteType(v int32) *ChangeCheckRequest {
	s.WhiteType = &v
	return s
}

type ChangeCheckRequestApproveFlowParam struct {
	ApproveNodes []*ChangeCheckRequestApproveFlowParamApproveNodes `json:"ApproveNodes,omitempty" xml:"ApproveNodes,omitempty" type:"Repeated"`
	AuthKey      *string                                           `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign     *string                                           `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	BgVid        *string                                           `json:"BgVid,omitempty" xml:"BgVid,omitempty"`
	FlowStatus   *int32                                            `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty"`
	Timestamp    *int64                                            `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ChangeCheckRequestApproveFlowParam) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestApproveFlowParam) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestApproveFlowParam) SetApproveNodes(v []*ChangeCheckRequestApproveFlowParamApproveNodes) *ChangeCheckRequestApproveFlowParam {
	s.ApproveNodes = v
	return s
}

func (s *ChangeCheckRequestApproveFlowParam) SetAuthKey(v string) *ChangeCheckRequestApproveFlowParam {
	s.AuthKey = &v
	return s
}

func (s *ChangeCheckRequestApproveFlowParam) SetAuthSign(v string) *ChangeCheckRequestApproveFlowParam {
	s.AuthSign = &v
	return s
}

func (s *ChangeCheckRequestApproveFlowParam) SetBgVid(v string) *ChangeCheckRequestApproveFlowParam {
	s.BgVid = &v
	return s
}

func (s *ChangeCheckRequestApproveFlowParam) SetFlowStatus(v int32) *ChangeCheckRequestApproveFlowParam {
	s.FlowStatus = &v
	return s
}

func (s *ChangeCheckRequestApproveFlowParam) SetTimestamp(v int64) *ChangeCheckRequestApproveFlowParam {
	s.Timestamp = &v
	return s
}

type ChangeCheckRequestApproveFlowParamApproveNodes struct {
	ApproverDTO      []*ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO `json:"ApproverDTO,omitempty" xml:"ApproverDTO,omitempty" type:"Repeated"`
	NodeStatus       *int32                                                       `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	ProcessName      *string                                                      `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	ProcessNodeOrder *int32                                                       `json:"ProcessNodeOrder,omitempty" xml:"ProcessNodeOrder,omitempty"`
	Strategy         *int32                                                       `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s ChangeCheckRequestApproveFlowParamApproveNodes) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestApproveFlowParamApproveNodes) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestApproveFlowParamApproveNodes) SetApproverDTO(v []*ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) *ChangeCheckRequestApproveFlowParamApproveNodes {
	s.ApproverDTO = v
	return s
}

func (s *ChangeCheckRequestApproveFlowParamApproveNodes) SetNodeStatus(v int32) *ChangeCheckRequestApproveFlowParamApproveNodes {
	s.NodeStatus = &v
	return s
}

func (s *ChangeCheckRequestApproveFlowParamApproveNodes) SetProcessName(v string) *ChangeCheckRequestApproveFlowParamApproveNodes {
	s.ProcessName = &v
	return s
}

func (s *ChangeCheckRequestApproveFlowParamApproveNodes) SetProcessNodeOrder(v int32) *ChangeCheckRequestApproveFlowParamApproveNodes {
	s.ProcessNodeOrder = &v
	return s
}

func (s *ChangeCheckRequestApproveFlowParamApproveNodes) SetStrategy(v int32) *ChangeCheckRequestApproveFlowParamApproveNodes {
	s.Strategy = &v
	return s
}

type ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO struct {
	ApproveDesc  *string `json:"ApproveDesc,omitempty" xml:"ApproveDesc,omitempty"`
	ApproveTime  *string `json:"ApproveTime,omitempty" xml:"ApproveTime,omitempty"`
	ApproverId   *string `json:"ApproverId,omitempty" xml:"ApproverId,omitempty"`
	ApproverName *string `json:"ApproverName,omitempty" xml:"ApproverName,omitempty"`
	Opinion      *int32  `json:"Opinion,omitempty" xml:"Opinion,omitempty"`
}

func (s ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) SetApproveDesc(v string) *ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO {
	s.ApproveDesc = &v
	return s
}

func (s *ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) SetApproveTime(v string) *ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO {
	s.ApproveTime = &v
	return s
}

func (s *ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) SetApproverId(v string) *ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO {
	s.ApproverId = &v
	return s
}

func (s *ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) SetApproverName(v string) *ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO {
	s.ApproverName = &v
	return s
}

func (s *ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) SetOpinion(v int32) *ChangeCheckRequestApproveFlowParamApproveNodesApproverDTO {
	s.Opinion = &v
	return s
}

type ChangeCheckRequestBgCustomTemplateExtraDTO struct {
	BgCustomTemplate      *string `json:"BgCustomTemplate,omitempty" xml:"BgCustomTemplate,omitempty"`
	BgCustomTemplateId    *int64  `json:"BgCustomTemplateId,omitempty" xml:"BgCustomTemplateId,omitempty"`
	BgCustomTemplateInfo  *string `json:"BgCustomTemplateInfo,omitempty" xml:"BgCustomTemplateInfo,omitempty"`
	BgCustomTemplateTitle *string `json:"BgCustomTemplateTitle,omitempty" xml:"BgCustomTemplateTitle,omitempty"`
	BgVid                 *string `json:"BgVid,omitempty" xml:"BgVid,omitempty"`
	ExtraInfo             *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
}

func (s ChangeCheckRequestBgCustomTemplateExtraDTO) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestBgCustomTemplateExtraDTO) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestBgCustomTemplateExtraDTO) SetBgCustomTemplate(v string) *ChangeCheckRequestBgCustomTemplateExtraDTO {
	s.BgCustomTemplate = &v
	return s
}

func (s *ChangeCheckRequestBgCustomTemplateExtraDTO) SetBgCustomTemplateId(v int64) *ChangeCheckRequestBgCustomTemplateExtraDTO {
	s.BgCustomTemplateId = &v
	return s
}

func (s *ChangeCheckRequestBgCustomTemplateExtraDTO) SetBgCustomTemplateInfo(v string) *ChangeCheckRequestBgCustomTemplateExtraDTO {
	s.BgCustomTemplateInfo = &v
	return s
}

func (s *ChangeCheckRequestBgCustomTemplateExtraDTO) SetBgCustomTemplateTitle(v string) *ChangeCheckRequestBgCustomTemplateExtraDTO {
	s.BgCustomTemplateTitle = &v
	return s
}

func (s *ChangeCheckRequestBgCustomTemplateExtraDTO) SetBgVid(v string) *ChangeCheckRequestBgCustomTemplateExtraDTO {
	s.BgVid = &v
	return s
}

func (s *ChangeCheckRequestBgCustomTemplateExtraDTO) SetExtraInfo(v string) *ChangeCheckRequestBgCustomTemplateExtraDTO {
	s.ExtraInfo = &v
	return s
}

type ChangeCheckRequestBlockInfos struct {
	HitInfos []*ChangeCheckRequestBlockInfosHitInfos `json:"HitInfos,omitempty" xml:"HitInfos,omitempty" type:"Repeated"`
	Id       *int64                                  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ChangeCheckRequestBlockInfos) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestBlockInfos) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestBlockInfos) SetHitInfos(v []*ChangeCheckRequestBlockInfosHitInfos) *ChangeCheckRequestBlockInfos {
	s.HitInfos = v
	return s
}

func (s *ChangeCheckRequestBlockInfos) SetId(v int64) *ChangeCheckRequestBlockInfos {
	s.Id = &v
	return s
}

type ChangeCheckRequestBlockInfosHitInfos struct {
	HitInfo   *string `json:"HitInfo,omitempty" xml:"HitInfo,omitempty"`
	HitObject *string `json:"HitObject,omitempty" xml:"HitObject,omitempty"`
	Scope     *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ChangeCheckRequestBlockInfosHitInfos) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestBlockInfosHitInfos) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestBlockInfosHitInfos) SetHitInfo(v string) *ChangeCheckRequestBlockInfosHitInfos {
	s.HitInfo = &v
	return s
}

func (s *ChangeCheckRequestBlockInfosHitInfos) SetHitObject(v string) *ChangeCheckRequestBlockInfosHitInfos {
	s.HitObject = &v
	return s
}

func (s *ChangeCheckRequestBlockInfosHitInfos) SetScope(v string) *ChangeCheckRequestBlockInfosHitInfos {
	s.Scope = &v
	return s
}

type ChangeCheckRequestCallBackInfo struct {
	Api        *string `json:"Api,omitempty" xml:"Api,omitempty"`
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	EndPoint   *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	ExtraInfo  *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	PopProduct *string `json:"PopProduct,omitempty" xml:"PopProduct,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url        *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ChangeCheckRequestCallBackInfo) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestCallBackInfo) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestCallBackInfo) SetApi(v string) *ChangeCheckRequestCallBackInfo {
	s.Api = &v
	return s
}

func (s *ChangeCheckRequestCallBackInfo) SetApiVersion(v string) *ChangeCheckRequestCallBackInfo {
	s.ApiVersion = &v
	return s
}

func (s *ChangeCheckRequestCallBackInfo) SetEndPoint(v string) *ChangeCheckRequestCallBackInfo {
	s.EndPoint = &v
	return s
}

func (s *ChangeCheckRequestCallBackInfo) SetExtraInfo(v string) *ChangeCheckRequestCallBackInfo {
	s.ExtraInfo = &v
	return s
}

func (s *ChangeCheckRequestCallBackInfo) SetPopProduct(v string) *ChangeCheckRequestCallBackInfo {
	s.PopProduct = &v
	return s
}

func (s *ChangeCheckRequestCallBackInfo) SetRegionId(v string) *ChangeCheckRequestCallBackInfo {
	s.RegionId = &v
	return s
}

func (s *ChangeCheckRequestCallBackInfo) SetType(v string) *ChangeCheckRequestCallBackInfo {
	s.Type = &v
	return s
}

func (s *ChangeCheckRequestCallBackInfo) SetUrl(v string) *ChangeCheckRequestCallBackInfo {
	s.Url = &v
	return s
}

type ChangeCheckRequestChangeTimes struct {
	ChangeEndTime   *int64 `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	ChangeStartTime *int64 `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
}

func (s ChangeCheckRequestChangeTimes) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestChangeTimes) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestChangeTimes) SetChangeEndTime(v int64) *ChangeCheckRequestChangeTimes {
	s.ChangeEndTime = &v
	return s
}

func (s *ChangeCheckRequestChangeTimes) SetChangeStartTime(v int64) *ChangeCheckRequestChangeTimes {
	s.ChangeStartTime = &v
	return s
}

type ChangeCheckRequestDamagedChangeNotices struct {
	BgCancelNoticeContent *string                                                     `json:"BgCancelNoticeContent,omitempty" xml:"BgCancelNoticeContent,omitempty"`
	BgCancelNoticeEventId *string                                                     `json:"BgCancelNoticeEventId,omitempty" xml:"BgCancelNoticeEventId,omitempty"`
	Channel               []*string                                                   `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Repeated"`
	Content               *string                                                     `json:"Content,omitempty" xml:"Content,omitempty"`
	EventId               *string                                                     `json:"EventId,omitempty" xml:"EventId,omitempty"`
	SensitiveCustomers    []*ChangeCheckRequestDamagedChangeNoticesSensitiveCustomers `json:"SensitiveCustomers,omitempty" xml:"SensitiveCustomers,omitempty" type:"Repeated"`
	Type                  *string                                                     `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ChangeCheckRequestDamagedChangeNotices) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestDamagedChangeNotices) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestDamagedChangeNotices) SetBgCancelNoticeContent(v string) *ChangeCheckRequestDamagedChangeNotices {
	s.BgCancelNoticeContent = &v
	return s
}

func (s *ChangeCheckRequestDamagedChangeNotices) SetBgCancelNoticeEventId(v string) *ChangeCheckRequestDamagedChangeNotices {
	s.BgCancelNoticeEventId = &v
	return s
}

func (s *ChangeCheckRequestDamagedChangeNotices) SetChannel(v []*string) *ChangeCheckRequestDamagedChangeNotices {
	s.Channel = v
	return s
}

func (s *ChangeCheckRequestDamagedChangeNotices) SetContent(v string) *ChangeCheckRequestDamagedChangeNotices {
	s.Content = &v
	return s
}

func (s *ChangeCheckRequestDamagedChangeNotices) SetEventId(v string) *ChangeCheckRequestDamagedChangeNotices {
	s.EventId = &v
	return s
}

func (s *ChangeCheckRequestDamagedChangeNotices) SetSensitiveCustomers(v []*ChangeCheckRequestDamagedChangeNoticesSensitiveCustomers) *ChangeCheckRequestDamagedChangeNotices {
	s.SensitiveCustomers = v
	return s
}

func (s *ChangeCheckRequestDamagedChangeNotices) SetType(v string) *ChangeCheckRequestDamagedChangeNotices {
	s.Type = &v
	return s
}

type ChangeCheckRequestDamagedChangeNoticesSensitiveCustomers struct {
	CustomerInfo []*ChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo `json:"CustomerInfo,omitempty" xml:"CustomerInfo,omitempty" type:"Repeated"`
	ProductCode  *string                                                                 `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ChangeCheckRequestDamagedChangeNoticesSensitiveCustomers) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestDamagedChangeNoticesSensitiveCustomers) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestDamagedChangeNoticesSensitiveCustomers) SetCustomerInfo(v []*ChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo) *ChangeCheckRequestDamagedChangeNoticesSensitiveCustomers {
	s.CustomerInfo = v
	return s
}

func (s *ChangeCheckRequestDamagedChangeNoticesSensitiveCustomers) SetProductCode(v string) *ChangeCheckRequestDamagedChangeNoticesSensitiveCustomers {
	s.ProductCode = &v
	return s
}

type ChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo struct {
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Type      *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid       *string                `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo) SetExtraInfo(v map[string]interface{}) *ChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo {
	s.ExtraInfo = v
	return s
}

func (s *ChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo) SetType(v string) *ChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo {
	s.Type = &v
	return s
}

func (s *ChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo) SetUid(v string) *ChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo {
	s.Uid = &v
	return s
}

type ChangeCheckRequestInfluenceInfo struct {
	NoticeInfos        []*ChangeCheckRequestInfluenceInfoNoticeInfos        `json:"NoticeInfos,omitempty" xml:"NoticeInfos,omitempty" type:"Repeated"`
	SensitiveCustomers []*ChangeCheckRequestInfluenceInfoSensitiveCustomers `json:"SensitiveCustomers,omitempty" xml:"SensitiveCustomers,omitempty" type:"Repeated"`
}

func (s ChangeCheckRequestInfluenceInfo) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestInfluenceInfo) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestInfluenceInfo) SetNoticeInfos(v []*ChangeCheckRequestInfluenceInfoNoticeInfos) *ChangeCheckRequestInfluenceInfo {
	s.NoticeInfos = v
	return s
}

func (s *ChangeCheckRequestInfluenceInfo) SetSensitiveCustomers(v []*ChangeCheckRequestInfluenceInfoSensitiveCustomers) *ChangeCheckRequestInfluenceInfo {
	s.SensitiveCustomers = v
	return s
}

type ChangeCheckRequestInfluenceInfoNoticeInfos struct {
	Channel []*string `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Repeated"`
	Content *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	EventId *string   `json:"EventId,omitempty" xml:"EventId,omitempty"`
}

func (s ChangeCheckRequestInfluenceInfoNoticeInfos) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestInfluenceInfoNoticeInfos) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestInfluenceInfoNoticeInfos) SetChannel(v []*string) *ChangeCheckRequestInfluenceInfoNoticeInfos {
	s.Channel = v
	return s
}

func (s *ChangeCheckRequestInfluenceInfoNoticeInfos) SetContent(v string) *ChangeCheckRequestInfluenceInfoNoticeInfos {
	s.Content = &v
	return s
}

func (s *ChangeCheckRequestInfluenceInfoNoticeInfos) SetEventId(v string) *ChangeCheckRequestInfluenceInfoNoticeInfos {
	s.EventId = &v
	return s
}

type ChangeCheckRequestInfluenceInfoSensitiveCustomers struct {
	CustomerInfo []*ChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo `json:"CustomerInfo,omitempty" xml:"CustomerInfo,omitempty" type:"Repeated"`
	ProductCode  *string                                                          `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ChangeCheckRequestInfluenceInfoSensitiveCustomers) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestInfluenceInfoSensitiveCustomers) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestInfluenceInfoSensitiveCustomers) SetCustomerInfo(v []*ChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo) *ChangeCheckRequestInfluenceInfoSensitiveCustomers {
	s.CustomerInfo = v
	return s
}

func (s *ChangeCheckRequestInfluenceInfoSensitiveCustomers) SetProductCode(v string) *ChangeCheckRequestInfluenceInfoSensitiveCustomers {
	s.ProductCode = &v
	return s
}

type ChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo struct {
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Type      *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid       *string                `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo) SetExtraInfo(v map[string]interface{}) *ChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo {
	s.ExtraInfo = v
	return s
}

func (s *ChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo) SetType(v string) *ChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo {
	s.Type = &v
	return s
}

func (s *ChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo) SetUid(v string) *ChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo {
	s.Uid = &v
	return s
}

type ChangeCheckRequestInstance struct {
	AttributionApp []*string `json:"AttributionApp,omitempty" xml:"AttributionApp,omitempty" type:"Repeated"`
	InfluenceApp   []*string `json:"InfluenceApp,omitempty" xml:"InfluenceApp,omitempty" type:"Repeated"`
	Instance       []*string `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	Nc             []*string `json:"Nc,omitempty" xml:"Nc,omitempty" type:"Repeated"`
	Uids           []*string `json:"Uids,omitempty" xml:"Uids,omitempty" type:"Repeated"`
}

func (s ChangeCheckRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestInstance) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestInstance) SetAttributionApp(v []*string) *ChangeCheckRequestInstance {
	s.AttributionApp = v
	return s
}

func (s *ChangeCheckRequestInstance) SetInfluenceApp(v []*string) *ChangeCheckRequestInstance {
	s.InfluenceApp = v
	return s
}

func (s *ChangeCheckRequestInstance) SetInstance(v []*string) *ChangeCheckRequestInstance {
	s.Instance = v
	return s
}

func (s *ChangeCheckRequestInstance) SetNc(v []*string) *ChangeCheckRequestInstance {
	s.Nc = v
	return s
}

func (s *ChangeCheckRequestInstance) SetUids(v []*string) *ChangeCheckRequestInstance {
	s.Uids = v
	return s
}

type ChangeCheckRequestProduct struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ChangeCheckRequestProduct) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestProduct) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestProduct) SetCode(v string) *ChangeCheckRequestProduct {
	s.Code = &v
	return s
}

func (s *ChangeCheckRequestProduct) SetName(v string) *ChangeCheckRequestProduct {
	s.Name = &v
	return s
}

type ChangeCheckRequestReleasePackageInfos struct {
	ProductCode    *string   `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ReleasePackage []*string `json:"ReleasePackage,omitempty" xml:"ReleasePackage,omitempty" type:"Repeated"`
}

func (s ChangeCheckRequestReleasePackageInfos) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckRequestReleasePackageInfos) GoString() string {
	return s.String()
}

func (s *ChangeCheckRequestReleasePackageInfos) SetProductCode(v string) *ChangeCheckRequestReleasePackageInfos {
	s.ProductCode = &v
	return s
}

func (s *ChangeCheckRequestReleasePackageInfos) SetReleasePackage(v []*string) *ChangeCheckRequestReleasePackageInfos {
	s.ReleasePackage = v
	return s
}

type ChangeCheckShrinkRequest struct {
	AffectCustomer             *string                                           `json:"AffectCustomer,omitempty" xml:"AffectCustomer,omitempty"`
	ApproveFlowParam           *ChangeCheckShrinkRequestApproveFlowParam         `json:"ApproveFlowParam,omitempty" xml:"ApproveFlowParam,omitempty" type:"Struct"`
	AuthKey                    *string                                           `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign                   *string                                           `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	BgCustomTemplateExtraDTO   *ChangeCheckShrinkRequestBgCustomTemplateExtraDTO `json:"BgCustomTemplateExtraDTO,omitempty" xml:"BgCustomTemplateExtraDTO,omitempty" type:"Struct"`
	BgVid                      *string                                           `json:"BgVid,omitempty" xml:"BgVid,omitempty"`
	BlockInfos                 []*ChangeCheckShrinkRequestBlockInfos             `json:"BlockInfos,omitempty" xml:"BlockInfos,omitempty" type:"Repeated"`
	CallBackInfo               *ChangeCheckShrinkRequestCallBackInfo             `json:"CallBackInfo,omitempty" xml:"CallBackInfo,omitempty" type:"Struct"`
	ChangeDataType             *string                                           `json:"ChangeDataType,omitempty" xml:"ChangeDataType,omitempty"`
	ChangeDesc                 *string                                           `json:"ChangeDesc,omitempty" xml:"ChangeDesc,omitempty"`
	ChangeEndTime              *int64                                            `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	ChangeEnv                  *string                                           `json:"ChangeEnv,omitempty" xml:"ChangeEnv,omitempty"`
	ChangeItems                *string                                           `json:"ChangeItems,omitempty" xml:"ChangeItems,omitempty"`
	ChangeObject               *string                                           `json:"ChangeObject,omitempty" xml:"ChangeObject,omitempty"`
	ChangeOptSubType           *string                                           `json:"ChangeOptSubType,omitempty" xml:"ChangeOptSubType,omitempty"`
	ChangeOptType              *string                                           `json:"ChangeOptType,omitempty" xml:"ChangeOptType,omitempty"`
	ChangeReason               *string                                           `json:"ChangeReason,omitempty" xml:"ChangeReason,omitempty"`
	ChangeRmarks               *string                                           `json:"ChangeRmarks,omitempty" xml:"ChangeRmarks,omitempty"`
	ChangeSchemes              *string                                           `json:"ChangeSchemes,omitempty" xml:"ChangeSchemes,omitempty"`
	ChangeStartTime            *int64                                            `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
	ChangeSubTypeDesc          *string                                           `json:"ChangeSubTypeDesc,omitempty" xml:"ChangeSubTypeDesc,omitempty"`
	ChangeSystem               *string                                           `json:"ChangeSystem,omitempty" xml:"ChangeSystem,omitempty"`
	ChangeTimes                []*ChangeCheckShrinkRequestChangeTimes            `json:"ChangeTimes,omitempty" xml:"ChangeTimes,omitempty" type:"Repeated"`
	ChangeTitle                *string                                           `json:"ChangeTitle,omitempty" xml:"ChangeTitle,omitempty"`
	ChangeValidation           *string                                           `json:"ChangeValidation,omitempty" xml:"ChangeValidation,omitempty"`
	CreatorEmpId               *string                                           `json:"CreatorEmpId,omitempty" xml:"CreatorEmpId,omitempty"`
	DamagedChangeNoticesShrink *string                                           `json:"DamagedChangeNotices,omitempty" xml:"DamagedChangeNotices,omitempty"`
	ExecutorEmpId              *string                                           `json:"ExecutorEmpId,omitempty" xml:"ExecutorEmpId,omitempty"`
	ExtraInfo                  *string                                           `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Follower                   []*string                                         `json:"Follower,omitempty" xml:"Follower,omitempty" type:"Repeated"`
	GrayStatus                 *string                                           `json:"GrayStatus,omitempty" xml:"GrayStatus,omitempty"`
	HarmChangeNoticeEnum       *string                                           `json:"HarmChangeNoticeEnum,omitempty" xml:"HarmChangeNoticeEnum,omitempty"`
	Incidence                  *string                                           `json:"Incidence,omitempty" xml:"Incidence,omitempty"`
	InfluenceInfo              *ChangeCheckShrinkRequestInfluenceInfo            `json:"InfluenceInfo,omitempty" xml:"InfluenceInfo,omitempty" type:"Struct"`
	Instance                   *ChangeCheckShrinkRequestInstance                 `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	NeedModifyDoc              *string                                           `json:"NeedModifyDoc,omitempty" xml:"NeedModifyDoc,omitempty"`
	Product                    []*ChangeCheckShrinkRequestProduct                `json:"Product,omitempty" xml:"Product,omitempty" type:"Repeated"`
	ReleasePackageInfos        []*ChangeCheckShrinkRequestReleasePackageInfos    `json:"ReleasePackageInfos,omitempty" xml:"ReleasePackageInfos,omitempty" type:"Repeated"`
	ReqTimestamp               *int64                                            `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	ReuseSourceOrderId         *string                                           `json:"ReuseSourceOrderId,omitempty" xml:"ReuseSourceOrderId,omitempty"`
	RiskLevel                  *string                                           `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Rollback                   *string                                           `json:"Rollback,omitempty" xml:"Rollback,omitempty"`
	SourceName                 *string                                           `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	SourceOrderId              *string                                           `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
	SourceUrl                  *string                                           `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	WhiteType                  *int32                                            `json:"WhiteType,omitempty" xml:"WhiteType,omitempty"`
}

func (s ChangeCheckShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequest) SetAffectCustomer(v string) *ChangeCheckShrinkRequest {
	s.AffectCustomer = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetApproveFlowParam(v *ChangeCheckShrinkRequestApproveFlowParam) *ChangeCheckShrinkRequest {
	s.ApproveFlowParam = v
	return s
}

func (s *ChangeCheckShrinkRequest) SetAuthKey(v string) *ChangeCheckShrinkRequest {
	s.AuthKey = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetAuthSign(v string) *ChangeCheckShrinkRequest {
	s.AuthSign = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetBgCustomTemplateExtraDTO(v *ChangeCheckShrinkRequestBgCustomTemplateExtraDTO) *ChangeCheckShrinkRequest {
	s.BgCustomTemplateExtraDTO = v
	return s
}

func (s *ChangeCheckShrinkRequest) SetBgVid(v string) *ChangeCheckShrinkRequest {
	s.BgVid = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetBlockInfos(v []*ChangeCheckShrinkRequestBlockInfos) *ChangeCheckShrinkRequest {
	s.BlockInfos = v
	return s
}

func (s *ChangeCheckShrinkRequest) SetCallBackInfo(v *ChangeCheckShrinkRequestCallBackInfo) *ChangeCheckShrinkRequest {
	s.CallBackInfo = v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeDataType(v string) *ChangeCheckShrinkRequest {
	s.ChangeDataType = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeDesc(v string) *ChangeCheckShrinkRequest {
	s.ChangeDesc = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeEndTime(v int64) *ChangeCheckShrinkRequest {
	s.ChangeEndTime = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeEnv(v string) *ChangeCheckShrinkRequest {
	s.ChangeEnv = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeItems(v string) *ChangeCheckShrinkRequest {
	s.ChangeItems = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeObject(v string) *ChangeCheckShrinkRequest {
	s.ChangeObject = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeOptSubType(v string) *ChangeCheckShrinkRequest {
	s.ChangeOptSubType = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeOptType(v string) *ChangeCheckShrinkRequest {
	s.ChangeOptType = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeReason(v string) *ChangeCheckShrinkRequest {
	s.ChangeReason = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeRmarks(v string) *ChangeCheckShrinkRequest {
	s.ChangeRmarks = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeSchemes(v string) *ChangeCheckShrinkRequest {
	s.ChangeSchemes = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeStartTime(v int64) *ChangeCheckShrinkRequest {
	s.ChangeStartTime = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeSubTypeDesc(v string) *ChangeCheckShrinkRequest {
	s.ChangeSubTypeDesc = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeSystem(v string) *ChangeCheckShrinkRequest {
	s.ChangeSystem = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeTimes(v []*ChangeCheckShrinkRequestChangeTimes) *ChangeCheckShrinkRequest {
	s.ChangeTimes = v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeTitle(v string) *ChangeCheckShrinkRequest {
	s.ChangeTitle = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetChangeValidation(v string) *ChangeCheckShrinkRequest {
	s.ChangeValidation = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetCreatorEmpId(v string) *ChangeCheckShrinkRequest {
	s.CreatorEmpId = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetDamagedChangeNoticesShrink(v string) *ChangeCheckShrinkRequest {
	s.DamagedChangeNoticesShrink = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetExecutorEmpId(v string) *ChangeCheckShrinkRequest {
	s.ExecutorEmpId = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetExtraInfo(v string) *ChangeCheckShrinkRequest {
	s.ExtraInfo = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetFollower(v []*string) *ChangeCheckShrinkRequest {
	s.Follower = v
	return s
}

func (s *ChangeCheckShrinkRequest) SetGrayStatus(v string) *ChangeCheckShrinkRequest {
	s.GrayStatus = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetHarmChangeNoticeEnum(v string) *ChangeCheckShrinkRequest {
	s.HarmChangeNoticeEnum = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetIncidence(v string) *ChangeCheckShrinkRequest {
	s.Incidence = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetInfluenceInfo(v *ChangeCheckShrinkRequestInfluenceInfo) *ChangeCheckShrinkRequest {
	s.InfluenceInfo = v
	return s
}

func (s *ChangeCheckShrinkRequest) SetInstance(v *ChangeCheckShrinkRequestInstance) *ChangeCheckShrinkRequest {
	s.Instance = v
	return s
}

func (s *ChangeCheckShrinkRequest) SetNeedModifyDoc(v string) *ChangeCheckShrinkRequest {
	s.NeedModifyDoc = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetProduct(v []*ChangeCheckShrinkRequestProduct) *ChangeCheckShrinkRequest {
	s.Product = v
	return s
}

func (s *ChangeCheckShrinkRequest) SetReleasePackageInfos(v []*ChangeCheckShrinkRequestReleasePackageInfos) *ChangeCheckShrinkRequest {
	s.ReleasePackageInfos = v
	return s
}

func (s *ChangeCheckShrinkRequest) SetReqTimestamp(v int64) *ChangeCheckShrinkRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetReuseSourceOrderId(v string) *ChangeCheckShrinkRequest {
	s.ReuseSourceOrderId = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetRiskLevel(v string) *ChangeCheckShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetRollback(v string) *ChangeCheckShrinkRequest {
	s.Rollback = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetSourceName(v string) *ChangeCheckShrinkRequest {
	s.SourceName = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetSourceOrderId(v string) *ChangeCheckShrinkRequest {
	s.SourceOrderId = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetSourceUrl(v string) *ChangeCheckShrinkRequest {
	s.SourceUrl = &v
	return s
}

func (s *ChangeCheckShrinkRequest) SetWhiteType(v int32) *ChangeCheckShrinkRequest {
	s.WhiteType = &v
	return s
}

type ChangeCheckShrinkRequestApproveFlowParam struct {
	ApproveNodes []*ChangeCheckShrinkRequestApproveFlowParamApproveNodes `json:"ApproveNodes,omitempty" xml:"ApproveNodes,omitempty" type:"Repeated"`
	AuthKey      *string                                                 `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign     *string                                                 `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	BgVid        *string                                                 `json:"BgVid,omitempty" xml:"BgVid,omitempty"`
	FlowStatus   *int32                                                  `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty"`
	Timestamp    *int64                                                  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ChangeCheckShrinkRequestApproveFlowParam) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestApproveFlowParam) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestApproveFlowParam) SetApproveNodes(v []*ChangeCheckShrinkRequestApproveFlowParamApproveNodes) *ChangeCheckShrinkRequestApproveFlowParam {
	s.ApproveNodes = v
	return s
}

func (s *ChangeCheckShrinkRequestApproveFlowParam) SetAuthKey(v string) *ChangeCheckShrinkRequestApproveFlowParam {
	s.AuthKey = &v
	return s
}

func (s *ChangeCheckShrinkRequestApproveFlowParam) SetAuthSign(v string) *ChangeCheckShrinkRequestApproveFlowParam {
	s.AuthSign = &v
	return s
}

func (s *ChangeCheckShrinkRequestApproveFlowParam) SetBgVid(v string) *ChangeCheckShrinkRequestApproveFlowParam {
	s.BgVid = &v
	return s
}

func (s *ChangeCheckShrinkRequestApproveFlowParam) SetFlowStatus(v int32) *ChangeCheckShrinkRequestApproveFlowParam {
	s.FlowStatus = &v
	return s
}

func (s *ChangeCheckShrinkRequestApproveFlowParam) SetTimestamp(v int64) *ChangeCheckShrinkRequestApproveFlowParam {
	s.Timestamp = &v
	return s
}

type ChangeCheckShrinkRequestApproveFlowParamApproveNodes struct {
	ApproverDTO      []*ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO `json:"ApproverDTO,omitempty" xml:"ApproverDTO,omitempty" type:"Repeated"`
	NodeStatus       *int32                                                             `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	ProcessName      *string                                                            `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	ProcessNodeOrder *int32                                                             `json:"ProcessNodeOrder,omitempty" xml:"ProcessNodeOrder,omitempty"`
	Strategy         *int32                                                             `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s ChangeCheckShrinkRequestApproveFlowParamApproveNodes) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestApproveFlowParamApproveNodes) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestApproveFlowParamApproveNodes) SetApproverDTO(v []*ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO) *ChangeCheckShrinkRequestApproveFlowParamApproveNodes {
	s.ApproverDTO = v
	return s
}

func (s *ChangeCheckShrinkRequestApproveFlowParamApproveNodes) SetNodeStatus(v int32) *ChangeCheckShrinkRequestApproveFlowParamApproveNodes {
	s.NodeStatus = &v
	return s
}

func (s *ChangeCheckShrinkRequestApproveFlowParamApproveNodes) SetProcessName(v string) *ChangeCheckShrinkRequestApproveFlowParamApproveNodes {
	s.ProcessName = &v
	return s
}

func (s *ChangeCheckShrinkRequestApproveFlowParamApproveNodes) SetProcessNodeOrder(v int32) *ChangeCheckShrinkRequestApproveFlowParamApproveNodes {
	s.ProcessNodeOrder = &v
	return s
}

func (s *ChangeCheckShrinkRequestApproveFlowParamApproveNodes) SetStrategy(v int32) *ChangeCheckShrinkRequestApproveFlowParamApproveNodes {
	s.Strategy = &v
	return s
}

type ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO struct {
	ApproveDesc  *string `json:"ApproveDesc,omitempty" xml:"ApproveDesc,omitempty"`
	ApproveTime  *string `json:"ApproveTime,omitempty" xml:"ApproveTime,omitempty"`
	ApproverId   *string `json:"ApproverId,omitempty" xml:"ApproverId,omitempty"`
	ApproverName *string `json:"ApproverName,omitempty" xml:"ApproverName,omitempty"`
	Opinion      *int32  `json:"Opinion,omitempty" xml:"Opinion,omitempty"`
}

func (s ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO) SetApproveDesc(v string) *ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO {
	s.ApproveDesc = &v
	return s
}

func (s *ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO) SetApproveTime(v string) *ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO {
	s.ApproveTime = &v
	return s
}

func (s *ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO) SetApproverId(v string) *ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO {
	s.ApproverId = &v
	return s
}

func (s *ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO) SetApproverName(v string) *ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO {
	s.ApproverName = &v
	return s
}

func (s *ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO) SetOpinion(v int32) *ChangeCheckShrinkRequestApproveFlowParamApproveNodesApproverDTO {
	s.Opinion = &v
	return s
}

type ChangeCheckShrinkRequestBgCustomTemplateExtraDTO struct {
	BgCustomTemplate      *string `json:"BgCustomTemplate,omitempty" xml:"BgCustomTemplate,omitempty"`
	BgCustomTemplateId    *int64  `json:"BgCustomTemplateId,omitempty" xml:"BgCustomTemplateId,omitempty"`
	BgCustomTemplateInfo  *string `json:"BgCustomTemplateInfo,omitempty" xml:"BgCustomTemplateInfo,omitempty"`
	BgCustomTemplateTitle *string `json:"BgCustomTemplateTitle,omitempty" xml:"BgCustomTemplateTitle,omitempty"`
	BgVid                 *string `json:"BgVid,omitempty" xml:"BgVid,omitempty"`
	ExtraInfo             *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
}

func (s ChangeCheckShrinkRequestBgCustomTemplateExtraDTO) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestBgCustomTemplateExtraDTO) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestBgCustomTemplateExtraDTO) SetBgCustomTemplate(v string) *ChangeCheckShrinkRequestBgCustomTemplateExtraDTO {
	s.BgCustomTemplate = &v
	return s
}

func (s *ChangeCheckShrinkRequestBgCustomTemplateExtraDTO) SetBgCustomTemplateId(v int64) *ChangeCheckShrinkRequestBgCustomTemplateExtraDTO {
	s.BgCustomTemplateId = &v
	return s
}

func (s *ChangeCheckShrinkRequestBgCustomTemplateExtraDTO) SetBgCustomTemplateInfo(v string) *ChangeCheckShrinkRequestBgCustomTemplateExtraDTO {
	s.BgCustomTemplateInfo = &v
	return s
}

func (s *ChangeCheckShrinkRequestBgCustomTemplateExtraDTO) SetBgCustomTemplateTitle(v string) *ChangeCheckShrinkRequestBgCustomTemplateExtraDTO {
	s.BgCustomTemplateTitle = &v
	return s
}

func (s *ChangeCheckShrinkRequestBgCustomTemplateExtraDTO) SetBgVid(v string) *ChangeCheckShrinkRequestBgCustomTemplateExtraDTO {
	s.BgVid = &v
	return s
}

func (s *ChangeCheckShrinkRequestBgCustomTemplateExtraDTO) SetExtraInfo(v string) *ChangeCheckShrinkRequestBgCustomTemplateExtraDTO {
	s.ExtraInfo = &v
	return s
}

type ChangeCheckShrinkRequestBlockInfos struct {
	HitInfos []*ChangeCheckShrinkRequestBlockInfosHitInfos `json:"HitInfos,omitempty" xml:"HitInfos,omitempty" type:"Repeated"`
	Id       *int64                                        `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ChangeCheckShrinkRequestBlockInfos) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestBlockInfos) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestBlockInfos) SetHitInfos(v []*ChangeCheckShrinkRequestBlockInfosHitInfos) *ChangeCheckShrinkRequestBlockInfos {
	s.HitInfos = v
	return s
}

func (s *ChangeCheckShrinkRequestBlockInfos) SetId(v int64) *ChangeCheckShrinkRequestBlockInfos {
	s.Id = &v
	return s
}

type ChangeCheckShrinkRequestBlockInfosHitInfos struct {
	HitInfo   *string `json:"HitInfo,omitempty" xml:"HitInfo,omitempty"`
	HitObject *string `json:"HitObject,omitempty" xml:"HitObject,omitempty"`
	Scope     *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ChangeCheckShrinkRequestBlockInfosHitInfos) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestBlockInfosHitInfos) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestBlockInfosHitInfos) SetHitInfo(v string) *ChangeCheckShrinkRequestBlockInfosHitInfos {
	s.HitInfo = &v
	return s
}

func (s *ChangeCheckShrinkRequestBlockInfosHitInfos) SetHitObject(v string) *ChangeCheckShrinkRequestBlockInfosHitInfos {
	s.HitObject = &v
	return s
}

func (s *ChangeCheckShrinkRequestBlockInfosHitInfos) SetScope(v string) *ChangeCheckShrinkRequestBlockInfosHitInfos {
	s.Scope = &v
	return s
}

type ChangeCheckShrinkRequestCallBackInfo struct {
	Api        *string `json:"Api,omitempty" xml:"Api,omitempty"`
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	EndPoint   *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	ExtraInfo  *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	PopProduct *string `json:"PopProduct,omitempty" xml:"PopProduct,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url        *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ChangeCheckShrinkRequestCallBackInfo) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestCallBackInfo) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestCallBackInfo) SetApi(v string) *ChangeCheckShrinkRequestCallBackInfo {
	s.Api = &v
	return s
}

func (s *ChangeCheckShrinkRequestCallBackInfo) SetApiVersion(v string) *ChangeCheckShrinkRequestCallBackInfo {
	s.ApiVersion = &v
	return s
}

func (s *ChangeCheckShrinkRequestCallBackInfo) SetEndPoint(v string) *ChangeCheckShrinkRequestCallBackInfo {
	s.EndPoint = &v
	return s
}

func (s *ChangeCheckShrinkRequestCallBackInfo) SetExtraInfo(v string) *ChangeCheckShrinkRequestCallBackInfo {
	s.ExtraInfo = &v
	return s
}

func (s *ChangeCheckShrinkRequestCallBackInfo) SetPopProduct(v string) *ChangeCheckShrinkRequestCallBackInfo {
	s.PopProduct = &v
	return s
}

func (s *ChangeCheckShrinkRequestCallBackInfo) SetRegionId(v string) *ChangeCheckShrinkRequestCallBackInfo {
	s.RegionId = &v
	return s
}

func (s *ChangeCheckShrinkRequestCallBackInfo) SetType(v string) *ChangeCheckShrinkRequestCallBackInfo {
	s.Type = &v
	return s
}

func (s *ChangeCheckShrinkRequestCallBackInfo) SetUrl(v string) *ChangeCheckShrinkRequestCallBackInfo {
	s.Url = &v
	return s
}

type ChangeCheckShrinkRequestChangeTimes struct {
	ChangeEndTime   *int64 `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	ChangeStartTime *int64 `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
}

func (s ChangeCheckShrinkRequestChangeTimes) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestChangeTimes) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestChangeTimes) SetChangeEndTime(v int64) *ChangeCheckShrinkRequestChangeTimes {
	s.ChangeEndTime = &v
	return s
}

func (s *ChangeCheckShrinkRequestChangeTimes) SetChangeStartTime(v int64) *ChangeCheckShrinkRequestChangeTimes {
	s.ChangeStartTime = &v
	return s
}

type ChangeCheckShrinkRequestInfluenceInfo struct {
	NoticeInfos        []*ChangeCheckShrinkRequestInfluenceInfoNoticeInfos        `json:"NoticeInfos,omitempty" xml:"NoticeInfos,omitempty" type:"Repeated"`
	SensitiveCustomers []*ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomers `json:"SensitiveCustomers,omitempty" xml:"SensitiveCustomers,omitempty" type:"Repeated"`
}

func (s ChangeCheckShrinkRequestInfluenceInfo) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestInfluenceInfo) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestInfluenceInfo) SetNoticeInfos(v []*ChangeCheckShrinkRequestInfluenceInfoNoticeInfos) *ChangeCheckShrinkRequestInfluenceInfo {
	s.NoticeInfos = v
	return s
}

func (s *ChangeCheckShrinkRequestInfluenceInfo) SetSensitiveCustomers(v []*ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomers) *ChangeCheckShrinkRequestInfluenceInfo {
	s.SensitiveCustomers = v
	return s
}

type ChangeCheckShrinkRequestInfluenceInfoNoticeInfos struct {
	Channel []*string `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Repeated"`
	Content *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	EventId *string   `json:"EventId,omitempty" xml:"EventId,omitempty"`
}

func (s ChangeCheckShrinkRequestInfluenceInfoNoticeInfos) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestInfluenceInfoNoticeInfos) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestInfluenceInfoNoticeInfos) SetChannel(v []*string) *ChangeCheckShrinkRequestInfluenceInfoNoticeInfos {
	s.Channel = v
	return s
}

func (s *ChangeCheckShrinkRequestInfluenceInfoNoticeInfos) SetContent(v string) *ChangeCheckShrinkRequestInfluenceInfoNoticeInfos {
	s.Content = &v
	return s
}

func (s *ChangeCheckShrinkRequestInfluenceInfoNoticeInfos) SetEventId(v string) *ChangeCheckShrinkRequestInfluenceInfoNoticeInfos {
	s.EventId = &v
	return s
}

type ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomers struct {
	CustomerInfo []*ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomersCustomerInfo `json:"CustomerInfo,omitempty" xml:"CustomerInfo,omitempty" type:"Repeated"`
	ProductCode  *string                                                                `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomers) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomers) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomers) SetCustomerInfo(v []*ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomersCustomerInfo) *ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomers {
	s.CustomerInfo = v
	return s
}

func (s *ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomers) SetProductCode(v string) *ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomers {
	s.ProductCode = &v
	return s
}

type ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomersCustomerInfo struct {
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Type      *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid       *string                `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomersCustomerInfo) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomersCustomerInfo) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomersCustomerInfo) SetExtraInfo(v map[string]interface{}) *ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomersCustomerInfo {
	s.ExtraInfo = v
	return s
}

func (s *ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomersCustomerInfo) SetType(v string) *ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomersCustomerInfo {
	s.Type = &v
	return s
}

func (s *ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomersCustomerInfo) SetUid(v string) *ChangeCheckShrinkRequestInfluenceInfoSensitiveCustomersCustomerInfo {
	s.Uid = &v
	return s
}

type ChangeCheckShrinkRequestInstance struct {
	AttributionApp []*string `json:"AttributionApp,omitempty" xml:"AttributionApp,omitempty" type:"Repeated"`
	InfluenceApp   []*string `json:"InfluenceApp,omitempty" xml:"InfluenceApp,omitempty" type:"Repeated"`
	Instance       []*string `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	Nc             []*string `json:"Nc,omitempty" xml:"Nc,omitempty" type:"Repeated"`
	Uids           []*string `json:"Uids,omitempty" xml:"Uids,omitempty" type:"Repeated"`
}

func (s ChangeCheckShrinkRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestInstance) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestInstance) SetAttributionApp(v []*string) *ChangeCheckShrinkRequestInstance {
	s.AttributionApp = v
	return s
}

func (s *ChangeCheckShrinkRequestInstance) SetInfluenceApp(v []*string) *ChangeCheckShrinkRequestInstance {
	s.InfluenceApp = v
	return s
}

func (s *ChangeCheckShrinkRequestInstance) SetInstance(v []*string) *ChangeCheckShrinkRequestInstance {
	s.Instance = v
	return s
}

func (s *ChangeCheckShrinkRequestInstance) SetNc(v []*string) *ChangeCheckShrinkRequestInstance {
	s.Nc = v
	return s
}

func (s *ChangeCheckShrinkRequestInstance) SetUids(v []*string) *ChangeCheckShrinkRequestInstance {
	s.Uids = v
	return s
}

type ChangeCheckShrinkRequestProduct struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ChangeCheckShrinkRequestProduct) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestProduct) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestProduct) SetCode(v string) *ChangeCheckShrinkRequestProduct {
	s.Code = &v
	return s
}

func (s *ChangeCheckShrinkRequestProduct) SetName(v string) *ChangeCheckShrinkRequestProduct {
	s.Name = &v
	return s
}

type ChangeCheckShrinkRequestReleasePackageInfos struct {
	ProductCode    *string   `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ReleasePackage []*string `json:"ReleasePackage,omitempty" xml:"ReleasePackage,omitempty" type:"Repeated"`
}

func (s ChangeCheckShrinkRequestReleasePackageInfos) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckShrinkRequestReleasePackageInfos) GoString() string {
	return s.String()
}

func (s *ChangeCheckShrinkRequestReleasePackageInfos) SetProductCode(v string) *ChangeCheckShrinkRequestReleasePackageInfos {
	s.ProductCode = &v
	return s
}

func (s *ChangeCheckShrinkRequestReleasePackageInfos) SetReleasePackage(v []*string) *ChangeCheckShrinkRequestReleasePackageInfos {
	s.ReleasePackage = v
	return s
}

type ChangeCheckResponseBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ChangeCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeCheckResponseBody) SetCode(v int32) *ChangeCheckResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeCheckResponseBody) SetData(v *ChangeCheckResponseBodyData) *ChangeCheckResponseBody {
	s.Data = v
	return s
}

func (s *ChangeCheckResponseBody) SetMessage(v string) *ChangeCheckResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeCheckResponseBody) SetRequestId(v string) *ChangeCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeCheckResponseBody) SetSuccess(v bool) *ChangeCheckResponseBody {
	s.Success = &v
	return s
}

type ChangeCheckResponseBodyData struct {
	ApproveResultUrl  *string                                         `json:"ApproveResultUrl,omitempty" xml:"ApproveResultUrl,omitempty"`
	BgCheckStatus     *string                                         `json:"BgCheckStatus,omitempty" xml:"BgCheckStatus,omitempty"`
	BgVid             *string                                         `json:"BgVid,omitempty" xml:"BgVid,omitempty"`
	ChangeStatus      *string                                         `json:"ChangeStatus,omitempty" xml:"ChangeStatus,omitempty"`
	CheckResultUrl    *string                                         `json:"CheckResultUrl,omitempty" xml:"CheckResultUrl,omitempty"`
	CheckStatus       *string                                         `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckholdReason   []*string                                       `json:"CheckholdReason,omitempty" xml:"CheckholdReason,omitempty" type:"Repeated"`
	RuleDetailUrlList []*ChangeCheckResponseBodyDataRuleDetailUrlList `json:"RuleDetailUrlList,omitempty" xml:"RuleDetailUrlList,omitempty" type:"Repeated"`
	SourceOrderId     *string                                         `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
}

func (s ChangeCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeCheckResponseBodyData) SetApproveResultUrl(v string) *ChangeCheckResponseBodyData {
	s.ApproveResultUrl = &v
	return s
}

func (s *ChangeCheckResponseBodyData) SetBgCheckStatus(v string) *ChangeCheckResponseBodyData {
	s.BgCheckStatus = &v
	return s
}

func (s *ChangeCheckResponseBodyData) SetBgVid(v string) *ChangeCheckResponseBodyData {
	s.BgVid = &v
	return s
}

func (s *ChangeCheckResponseBodyData) SetChangeStatus(v string) *ChangeCheckResponseBodyData {
	s.ChangeStatus = &v
	return s
}

func (s *ChangeCheckResponseBodyData) SetCheckResultUrl(v string) *ChangeCheckResponseBodyData {
	s.CheckResultUrl = &v
	return s
}

func (s *ChangeCheckResponseBodyData) SetCheckStatus(v string) *ChangeCheckResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *ChangeCheckResponseBodyData) SetCheckholdReason(v []*string) *ChangeCheckResponseBodyData {
	s.CheckholdReason = v
	return s
}

func (s *ChangeCheckResponseBodyData) SetRuleDetailUrlList(v []*ChangeCheckResponseBodyDataRuleDetailUrlList) *ChangeCheckResponseBodyData {
	s.RuleDetailUrlList = v
	return s
}

func (s *ChangeCheckResponseBodyData) SetSourceOrderId(v string) *ChangeCheckResponseBodyData {
	s.SourceOrderId = &v
	return s
}

type ChangeCheckResponseBodyDataRuleDetailUrlList struct {
	SceneEnum *string `json:"SceneEnum,omitempty" xml:"SceneEnum,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ChangeCheckResponseBodyDataRuleDetailUrlList) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckResponseBodyDataRuleDetailUrlList) GoString() string {
	return s.String()
}

func (s *ChangeCheckResponseBodyDataRuleDetailUrlList) SetSceneEnum(v string) *ChangeCheckResponseBodyDataRuleDetailUrlList {
	s.SceneEnum = &v
	return s
}

func (s *ChangeCheckResponseBodyDataRuleDetailUrlList) SetTitle(v string) *ChangeCheckResponseBodyDataRuleDetailUrlList {
	s.Title = &v
	return s
}

func (s *ChangeCheckResponseBodyDataRuleDetailUrlList) SetUrl(v string) *ChangeCheckResponseBodyDataRuleDetailUrlList {
	s.Url = &v
	return s
}

type ChangeCheckResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeCheckResponse) GoString() string {
	return s.String()
}

func (s *ChangeCheckResponse) SetHeaders(v map[string]*string) *ChangeCheckResponse {
	s.Headers = v
	return s
}

func (s *ChangeCheckResponse) SetStatusCode(v int32) *ChangeCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeCheckResponse) SetBody(v *ChangeCheckResponseBody) *ChangeCheckResponse {
	s.Body = v
	return s
}

type ChangeEndRequest struct {
	AuthKey       *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign      *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	ChangeEndTime *int64  `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	ChangeResult  *string `json:"ChangeResult,omitempty" xml:"ChangeResult,omitempty"`
	CurBatchNo    *int32  `json:"CurBatchNo,omitempty" xml:"CurBatchNo,omitempty"`
	ExecutorEmpId *string `json:"ExecutorEmpId,omitempty" xml:"ExecutorEmpId,omitempty"`
	ReqTimestamp  *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	SourceOrderId *string `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
	TotalBatchNo  *int32  `json:"TotalBatchNo,omitempty" xml:"TotalBatchNo,omitempty"`
}

func (s ChangeEndRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeEndRequest) GoString() string {
	return s.String()
}

func (s *ChangeEndRequest) SetAuthKey(v string) *ChangeEndRequest {
	s.AuthKey = &v
	return s
}

func (s *ChangeEndRequest) SetAuthSign(v string) *ChangeEndRequest {
	s.AuthSign = &v
	return s
}

func (s *ChangeEndRequest) SetChangeEndTime(v int64) *ChangeEndRequest {
	s.ChangeEndTime = &v
	return s
}

func (s *ChangeEndRequest) SetChangeResult(v string) *ChangeEndRequest {
	s.ChangeResult = &v
	return s
}

func (s *ChangeEndRequest) SetCurBatchNo(v int32) *ChangeEndRequest {
	s.CurBatchNo = &v
	return s
}

func (s *ChangeEndRequest) SetExecutorEmpId(v string) *ChangeEndRequest {
	s.ExecutorEmpId = &v
	return s
}

func (s *ChangeEndRequest) SetReqTimestamp(v int64) *ChangeEndRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *ChangeEndRequest) SetSourceOrderId(v string) *ChangeEndRequest {
	s.SourceOrderId = &v
	return s
}

func (s *ChangeEndRequest) SetTotalBatchNo(v int32) *ChangeEndRequest {
	s.TotalBatchNo = &v
	return s
}

type ChangeEndResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeEndResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeEndResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeEndResponseBody) SetCode(v int32) *ChangeEndResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeEndResponseBody) SetData(v string) *ChangeEndResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeEndResponseBody) SetMessage(v string) *ChangeEndResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeEndResponseBody) SetRequestId(v string) *ChangeEndResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeEndResponseBody) SetSuccess(v bool) *ChangeEndResponseBody {
	s.Success = &v
	return s
}

type ChangeEndResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeEndResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeEndResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeEndResponse) GoString() string {
	return s.String()
}

func (s *ChangeEndResponse) SetHeaders(v map[string]*string) *ChangeEndResponse {
	s.Headers = v
	return s
}

func (s *ChangeEndResponse) SetStatusCode(v int32) *ChangeEndResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeEndResponse) SetBody(v *ChangeEndResponseBody) *ChangeEndResponse {
	s.Body = v
	return s
}

type ChangeStartRequest struct {
	AuthKey         *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign        *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	ChangeEndTime   *int64  `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	ChangeObject    *string `json:"ChangeObject,omitempty" xml:"ChangeObject,omitempty"`
	ChangeOptType   *string `json:"ChangeOptType,omitempty" xml:"ChangeOptType,omitempty"`
	ChangeStartTime *int64  `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
	ChangeTitle     *string `json:"ChangeTitle,omitempty" xml:"ChangeTitle,omitempty"`
	CreatorEmpId    *string `json:"CreatorEmpId,omitempty" xml:"CreatorEmpId,omitempty"`
	CurBatchNo      *int32  `json:"CurBatchNo,omitempty" xml:"CurBatchNo,omitempty"`
	ExecutorEmpId   *string `json:"ExecutorEmpId,omitempty" xml:"ExecutorEmpId,omitempty"`
	ReqTimestamp    *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	SourceOrderId   *string `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
	TotalBatchNo    *int32  `json:"TotalBatchNo,omitempty" xml:"TotalBatchNo,omitempty"`
}

func (s ChangeStartRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeStartRequest) GoString() string {
	return s.String()
}

func (s *ChangeStartRequest) SetAuthKey(v string) *ChangeStartRequest {
	s.AuthKey = &v
	return s
}

func (s *ChangeStartRequest) SetAuthSign(v string) *ChangeStartRequest {
	s.AuthSign = &v
	return s
}

func (s *ChangeStartRequest) SetChangeEndTime(v int64) *ChangeStartRequest {
	s.ChangeEndTime = &v
	return s
}

func (s *ChangeStartRequest) SetChangeObject(v string) *ChangeStartRequest {
	s.ChangeObject = &v
	return s
}

func (s *ChangeStartRequest) SetChangeOptType(v string) *ChangeStartRequest {
	s.ChangeOptType = &v
	return s
}

func (s *ChangeStartRequest) SetChangeStartTime(v int64) *ChangeStartRequest {
	s.ChangeStartTime = &v
	return s
}

func (s *ChangeStartRequest) SetChangeTitle(v string) *ChangeStartRequest {
	s.ChangeTitle = &v
	return s
}

func (s *ChangeStartRequest) SetCreatorEmpId(v string) *ChangeStartRequest {
	s.CreatorEmpId = &v
	return s
}

func (s *ChangeStartRequest) SetCurBatchNo(v int32) *ChangeStartRequest {
	s.CurBatchNo = &v
	return s
}

func (s *ChangeStartRequest) SetExecutorEmpId(v string) *ChangeStartRequest {
	s.ExecutorEmpId = &v
	return s
}

func (s *ChangeStartRequest) SetReqTimestamp(v int64) *ChangeStartRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *ChangeStartRequest) SetSourceOrderId(v string) *ChangeStartRequest {
	s.SourceOrderId = &v
	return s
}

func (s *ChangeStartRequest) SetTotalBatchNo(v int32) *ChangeStartRequest {
	s.TotalBatchNo = &v
	return s
}

type ChangeStartResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeStartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeStartResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeStartResponseBody) SetCode(v int32) *ChangeStartResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeStartResponseBody) SetData(v string) *ChangeStartResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeStartResponseBody) SetMessage(v string) *ChangeStartResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeStartResponseBody) SetRequestId(v string) *ChangeStartResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeStartResponseBody) SetSuccess(v bool) *ChangeStartResponseBody {
	s.Success = &v
	return s
}

type ChangeStartResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeStartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeStartResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeStartResponse) GoString() string {
	return s.String()
}

func (s *ChangeStartResponse) SetHeaders(v map[string]*string) *ChangeStartResponse {
	s.Headers = v
	return s
}

func (s *ChangeStartResponse) SetStatusCode(v int32) *ChangeStartResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeStartResponse) SetBody(v *ChangeStartResponseBody) *ChangeStartResponse {
	s.Body = v
	return s
}

type CreateBlockRequest struct {
	ApproveStrategyNodes []*CreateBlockRequestApproveStrategyNodes `json:"ApproveStrategyNodes,omitempty" xml:"ApproveStrategyNodes,omitempty" type:"Repeated"`
	BlockId              *int64                                    `json:"BlockId,omitempty" xml:"BlockId,omitempty"`
	Director             *string                                   `json:"Director,omitempty" xml:"Director,omitempty"`
	IsNeedApprove        *int32                                    `json:"IsNeedApprove,omitempty" xml:"IsNeedApprove,omitempty"`
	IsRecall             *int32                                    `json:"IsRecall,omitempty" xml:"IsRecall,omitempty"`
	IsTemplate           *int32                                    `json:"IsTemplate,omitempty" xml:"IsTemplate,omitempty"`
	LabelName            *string                                   `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	NoticeDesc           *string                                   `json:"NoticeDesc,omitempty" xml:"NoticeDesc,omitempty"`
	NoticeEnclosureInfos []*CreateBlockRequestNoticeEnclosureInfos `json:"NoticeEnclosureInfos,omitempty" xml:"NoticeEnclosureInfos,omitempty" type:"Repeated"`
	NoticeRequestLink    *string                                   `json:"NoticeRequestLink,omitempty" xml:"NoticeRequestLink,omitempty"`
	NoticeType           *int32                                    `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	Reason               *string                                   `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Scene                *int32                                    `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Scopes               []*CreateBlockRequestScopes               `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Repeated"`
	Status               *int32                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Title                *string                                   `json:"Title,omitempty" xml:"Title,omitempty"`
	Type                 *string                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	VersionId            *int64                                    `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateBlockRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockRequest) GoString() string {
	return s.String()
}

func (s *CreateBlockRequest) SetApproveStrategyNodes(v []*CreateBlockRequestApproveStrategyNodes) *CreateBlockRequest {
	s.ApproveStrategyNodes = v
	return s
}

func (s *CreateBlockRequest) SetBlockId(v int64) *CreateBlockRequest {
	s.BlockId = &v
	return s
}

func (s *CreateBlockRequest) SetDirector(v string) *CreateBlockRequest {
	s.Director = &v
	return s
}

func (s *CreateBlockRequest) SetIsNeedApprove(v int32) *CreateBlockRequest {
	s.IsNeedApprove = &v
	return s
}

func (s *CreateBlockRequest) SetIsRecall(v int32) *CreateBlockRequest {
	s.IsRecall = &v
	return s
}

func (s *CreateBlockRequest) SetIsTemplate(v int32) *CreateBlockRequest {
	s.IsTemplate = &v
	return s
}

func (s *CreateBlockRequest) SetLabelName(v string) *CreateBlockRequest {
	s.LabelName = &v
	return s
}

func (s *CreateBlockRequest) SetNoticeDesc(v string) *CreateBlockRequest {
	s.NoticeDesc = &v
	return s
}

func (s *CreateBlockRequest) SetNoticeEnclosureInfos(v []*CreateBlockRequestNoticeEnclosureInfos) *CreateBlockRequest {
	s.NoticeEnclosureInfos = v
	return s
}

func (s *CreateBlockRequest) SetNoticeRequestLink(v string) *CreateBlockRequest {
	s.NoticeRequestLink = &v
	return s
}

func (s *CreateBlockRequest) SetNoticeType(v int32) *CreateBlockRequest {
	s.NoticeType = &v
	return s
}

func (s *CreateBlockRequest) SetReason(v string) *CreateBlockRequest {
	s.Reason = &v
	return s
}

func (s *CreateBlockRequest) SetScene(v int32) *CreateBlockRequest {
	s.Scene = &v
	return s
}

func (s *CreateBlockRequest) SetScopes(v []*CreateBlockRequestScopes) *CreateBlockRequest {
	s.Scopes = v
	return s
}

func (s *CreateBlockRequest) SetStatus(v int32) *CreateBlockRequest {
	s.Status = &v
	return s
}

func (s *CreateBlockRequest) SetTitle(v string) *CreateBlockRequest {
	s.Title = &v
	return s
}

func (s *CreateBlockRequest) SetType(v string) *CreateBlockRequest {
	s.Type = &v
	return s
}

func (s *CreateBlockRequest) SetVersionId(v int64) *CreateBlockRequest {
	s.VersionId = &v
	return s
}

type CreateBlockRequestApproveStrategyNodes struct {
	ApproveRuleType *int32    `json:"ApproveRuleType,omitempty" xml:"ApproveRuleType,omitempty"`
	ApproveType     *int32    `json:"ApproveType,omitempty" xml:"ApproveType,omitempty"`
	Id              *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeCode        *string   `json:"NodeCode,omitempty" xml:"NodeCode,omitempty"`
	PriorityOrder   *int32    `json:"PriorityOrder,omitempty" xml:"PriorityOrder,omitempty"`
	RoleCode        *int32    `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	RoleValue       []*string `json:"RoleValue,omitempty" xml:"RoleValue,omitempty" type:"Repeated"`
	TemplateId      *int64    `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateBlockRequestApproveStrategyNodes) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockRequestApproveStrategyNodes) GoString() string {
	return s.String()
}

func (s *CreateBlockRequestApproveStrategyNodes) SetApproveRuleType(v int32) *CreateBlockRequestApproveStrategyNodes {
	s.ApproveRuleType = &v
	return s
}

func (s *CreateBlockRequestApproveStrategyNodes) SetApproveType(v int32) *CreateBlockRequestApproveStrategyNodes {
	s.ApproveType = &v
	return s
}

func (s *CreateBlockRequestApproveStrategyNodes) SetId(v int64) *CreateBlockRequestApproveStrategyNodes {
	s.Id = &v
	return s
}

func (s *CreateBlockRequestApproveStrategyNodes) SetName(v string) *CreateBlockRequestApproveStrategyNodes {
	s.Name = &v
	return s
}

func (s *CreateBlockRequestApproveStrategyNodes) SetNodeCode(v string) *CreateBlockRequestApproveStrategyNodes {
	s.NodeCode = &v
	return s
}

func (s *CreateBlockRequestApproveStrategyNodes) SetPriorityOrder(v int32) *CreateBlockRequestApproveStrategyNodes {
	s.PriorityOrder = &v
	return s
}

func (s *CreateBlockRequestApproveStrategyNodes) SetRoleCode(v int32) *CreateBlockRequestApproveStrategyNodes {
	s.RoleCode = &v
	return s
}

func (s *CreateBlockRequestApproveStrategyNodes) SetRoleValue(v []*string) *CreateBlockRequestApproveStrategyNodes {
	s.RoleValue = v
	return s
}

func (s *CreateBlockRequestApproveStrategyNodes) SetTemplateId(v int64) *CreateBlockRequestApproveStrategyNodes {
	s.TemplateId = &v
	return s
}

type CreateBlockRequestNoticeEnclosureInfos struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Url  *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateBlockRequestNoticeEnclosureInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockRequestNoticeEnclosureInfos) GoString() string {
	return s.String()
}

func (s *CreateBlockRequestNoticeEnclosureInfos) SetName(v string) *CreateBlockRequestNoticeEnclosureInfos {
	s.Name = &v
	return s
}

func (s *CreateBlockRequestNoticeEnclosureInfos) SetUrl(v string) *CreateBlockRequestNoticeEnclosureInfos {
	s.Url = &v
	return s
}

type CreateBlockRequestScopes struct {
	BlockHarm  []*int32                            `json:"BlockHarm,omitempty" xml:"BlockHarm,omitempty" type:"Repeated"`
	BlockScope *CreateBlockRequestScopesBlockScope `json:"BlockScope,omitempty" xml:"BlockScope,omitempty" type:"Struct"`
	EffectTime []*int64                            `json:"EffectTime,omitempty" xml:"EffectTime,omitempty" type:"Repeated"`
	ScopeRule  *string                             `json:"ScopeRule,omitempty" xml:"ScopeRule,omitempty"`
}

func (s CreateBlockRequestScopes) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockRequestScopes) GoString() string {
	return s.String()
}

func (s *CreateBlockRequestScopes) SetBlockHarm(v []*int32) *CreateBlockRequestScopes {
	s.BlockHarm = v
	return s
}

func (s *CreateBlockRequestScopes) SetBlockScope(v *CreateBlockRequestScopesBlockScope) *CreateBlockRequestScopes {
	s.BlockScope = v
	return s
}

func (s *CreateBlockRequestScopes) SetEffectTime(v []*int64) *CreateBlockRequestScopes {
	s.EffectTime = v
	return s
}

func (s *CreateBlockRequestScopes) SetScopeRule(v string) *CreateBlockRequestScopes {
	s.ScopeRule = &v
	return s
}

type CreateBlockRequestScopesBlockScope struct {
	App            *CreateBlockRequestScopesBlockScopeApp        `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	BgSystem       []*CreateBlockRequestScopesBlockScopeBgSystem `json:"BgSystem,omitempty" xml:"BgSystem,omitempty" type:"Repeated"`
	Cluster        *CreateBlockRequestScopesBlockScopeCluster    `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Struct"`
	Customer       []*CreateBlockRequestScopesBlockScopeCustomer `json:"Customer,omitempty" xml:"Customer,omitempty" type:"Repeated"`
	Dept           []*string                                     `json:"Dept,omitempty" xml:"Dept,omitempty" type:"Repeated"`
	Express        *string                                       `json:"Express,omitempty" xml:"Express,omitempty"`
	Infrastructure []*string                                     `json:"Infrastructure,omitempty" xml:"Infrastructure,omitempty" type:"Repeated"`
	Product        []*CreateBlockRequestScopesBlockScopeProduct  `json:"Product,omitempty" xml:"Product,omitempty" type:"Repeated"`
}

func (s CreateBlockRequestScopesBlockScope) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockRequestScopesBlockScope) GoString() string {
	return s.String()
}

func (s *CreateBlockRequestScopesBlockScope) SetApp(v *CreateBlockRequestScopesBlockScopeApp) *CreateBlockRequestScopesBlockScope {
	s.App = v
	return s
}

func (s *CreateBlockRequestScopesBlockScope) SetBgSystem(v []*CreateBlockRequestScopesBlockScopeBgSystem) *CreateBlockRequestScopesBlockScope {
	s.BgSystem = v
	return s
}

func (s *CreateBlockRequestScopesBlockScope) SetCluster(v *CreateBlockRequestScopesBlockScopeCluster) *CreateBlockRequestScopesBlockScope {
	s.Cluster = v
	return s
}

func (s *CreateBlockRequestScopesBlockScope) SetCustomer(v []*CreateBlockRequestScopesBlockScopeCustomer) *CreateBlockRequestScopesBlockScope {
	s.Customer = v
	return s
}

func (s *CreateBlockRequestScopesBlockScope) SetDept(v []*string) *CreateBlockRequestScopesBlockScope {
	s.Dept = v
	return s
}

func (s *CreateBlockRequestScopesBlockScope) SetExpress(v string) *CreateBlockRequestScopesBlockScope {
	s.Express = &v
	return s
}

func (s *CreateBlockRequestScopesBlockScope) SetInfrastructure(v []*string) *CreateBlockRequestScopesBlockScope {
	s.Infrastructure = v
	return s
}

func (s *CreateBlockRequestScopesBlockScope) SetProduct(v []*CreateBlockRequestScopesBlockScopeProduct) *CreateBlockRequestScopesBlockScope {
	s.Product = v
	return s
}

type CreateBlockRequestScopesBlockScopeApp struct {
	AppName []*string `json:"AppName,omitempty" xml:"AppName,omitempty" type:"Repeated"`
	Type    *int32    `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateBlockRequestScopesBlockScopeApp) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockRequestScopesBlockScopeApp) GoString() string {
	return s.String()
}

func (s *CreateBlockRequestScopesBlockScopeApp) SetAppName(v []*string) *CreateBlockRequestScopesBlockScopeApp {
	s.AppName = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeApp) SetType(v int32) *CreateBlockRequestScopesBlockScopeApp {
	s.Type = &v
	return s
}

type CreateBlockRequestScopesBlockScopeBgSystem struct {
	RelateCodes  []*string `json:"RelateCodes,omitempty" xml:"RelateCodes,omitempty" type:"Repeated"`
	SelfCodeName *string   `json:"SelfCodeName,omitempty" xml:"SelfCodeName,omitempty"`
}

func (s CreateBlockRequestScopesBlockScopeBgSystem) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockRequestScopesBlockScopeBgSystem) GoString() string {
	return s.String()
}

func (s *CreateBlockRequestScopesBlockScopeBgSystem) SetRelateCodes(v []*string) *CreateBlockRequestScopesBlockScopeBgSystem {
	s.RelateCodes = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeBgSystem) SetSelfCodeName(v string) *CreateBlockRequestScopesBlockScopeBgSystem {
	s.SelfCodeName = &v
	return s
}

type CreateBlockRequestScopesBlockScopeCluster struct {
	CodeNames []*string                                             `json:"CodeNames,omitempty" xml:"CodeNames,omitempty" type:"Repeated"`
	Relations []*CreateBlockRequestScopesBlockScopeClusterRelations `json:"Relations,omitempty" xml:"Relations,omitempty" type:"Repeated"`
}

func (s CreateBlockRequestScopesBlockScopeCluster) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockRequestScopesBlockScopeCluster) GoString() string {
	return s.String()
}

func (s *CreateBlockRequestScopesBlockScopeCluster) SetCodeNames(v []*string) *CreateBlockRequestScopesBlockScopeCluster {
	s.CodeNames = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeCluster) SetRelations(v []*CreateBlockRequestScopesBlockScopeClusterRelations) *CreateBlockRequestScopesBlockScopeCluster {
	s.Relations = v
	return s
}

type CreateBlockRequestScopesBlockScopeClusterRelations struct {
	AppCodes    []*string `json:"AppCodes,omitempty" xml:"AppCodes,omitempty" type:"Repeated"`
	LabelCodes  []*string `json:"LabelCodes,omitempty" xml:"LabelCodes,omitempty" type:"Repeated"`
	RelateCodes []*string `json:"RelateCodes,omitempty" xml:"RelateCodes,omitempty" type:"Repeated"`
	SelfCode    *string   `json:"SelfCode,omitempty" xml:"SelfCode,omitempty"`
}

func (s CreateBlockRequestScopesBlockScopeClusterRelations) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockRequestScopesBlockScopeClusterRelations) GoString() string {
	return s.String()
}

func (s *CreateBlockRequestScopesBlockScopeClusterRelations) SetAppCodes(v []*string) *CreateBlockRequestScopesBlockScopeClusterRelations {
	s.AppCodes = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeClusterRelations) SetLabelCodes(v []*string) *CreateBlockRequestScopesBlockScopeClusterRelations {
	s.LabelCodes = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeClusterRelations) SetRelateCodes(v []*string) *CreateBlockRequestScopesBlockScopeClusterRelations {
	s.RelateCodes = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeClusterRelations) SetSelfCode(v string) *CreateBlockRequestScopesBlockScopeClusterRelations {
	s.SelfCode = &v
	return s
}

type CreateBlockRequestScopesBlockScopeCustomer struct {
	CodeNames []*string                                              `json:"CodeNames,omitempty" xml:"CodeNames,omitempty" type:"Repeated"`
	Relations []*CreateBlockRequestScopesBlockScopeCustomerRelations `json:"Relations,omitempty" xml:"Relations,omitempty" type:"Repeated"`
	Uid       *string                                                `json:"Uid,omitempty" xml:"Uid,omitempty"`
	ViewCodes []*int32                                               `json:"ViewCodes,omitempty" xml:"ViewCodes,omitempty" type:"Repeated"`
}

func (s CreateBlockRequestScopesBlockScopeCustomer) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockRequestScopesBlockScopeCustomer) GoString() string {
	return s.String()
}

func (s *CreateBlockRequestScopesBlockScopeCustomer) SetCodeNames(v []*string) *CreateBlockRequestScopesBlockScopeCustomer {
	s.CodeNames = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeCustomer) SetRelations(v []*CreateBlockRequestScopesBlockScopeCustomerRelations) *CreateBlockRequestScopesBlockScopeCustomer {
	s.Relations = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeCustomer) SetUid(v string) *CreateBlockRequestScopesBlockScopeCustomer {
	s.Uid = &v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeCustomer) SetViewCodes(v []*int32) *CreateBlockRequestScopesBlockScopeCustomer {
	s.ViewCodes = v
	return s
}

type CreateBlockRequestScopesBlockScopeCustomerRelations struct {
	AppCodes    []*string `json:"AppCodes,omitempty" xml:"AppCodes,omitempty" type:"Repeated"`
	LabelCodes  []*string `json:"LabelCodes,omitempty" xml:"LabelCodes,omitempty" type:"Repeated"`
	RelateCodes []*string `json:"RelateCodes,omitempty" xml:"RelateCodes,omitempty" type:"Repeated"`
	SelfCode    *string   `json:"SelfCode,omitempty" xml:"SelfCode,omitempty"`
}

func (s CreateBlockRequestScopesBlockScopeCustomerRelations) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockRequestScopesBlockScopeCustomerRelations) GoString() string {
	return s.String()
}

func (s *CreateBlockRequestScopesBlockScopeCustomerRelations) SetAppCodes(v []*string) *CreateBlockRequestScopesBlockScopeCustomerRelations {
	s.AppCodes = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeCustomerRelations) SetLabelCodes(v []*string) *CreateBlockRequestScopesBlockScopeCustomerRelations {
	s.LabelCodes = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeCustomerRelations) SetRelateCodes(v []*string) *CreateBlockRequestScopesBlockScopeCustomerRelations {
	s.RelateCodes = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeCustomerRelations) SetSelfCode(v string) *CreateBlockRequestScopesBlockScopeCustomerRelations {
	s.SelfCode = &v
	return s
}

type CreateBlockRequestScopesBlockScopeProduct struct {
	CodeNames []*string                                             `json:"CodeNames,omitempty" xml:"CodeNames,omitempty" type:"Repeated"`
	Key       *string                                               `json:"Key,omitempty" xml:"Key,omitempty"`
	Relations []*CreateBlockRequestScopesBlockScopeProductRelations `json:"Relations,omitempty" xml:"Relations,omitempty" type:"Repeated"`
	ViewCode  []*string                                             `json:"ViewCode,omitempty" xml:"ViewCode,omitempty" type:"Repeated"`
}

func (s CreateBlockRequestScopesBlockScopeProduct) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockRequestScopesBlockScopeProduct) GoString() string {
	return s.String()
}

func (s *CreateBlockRequestScopesBlockScopeProduct) SetCodeNames(v []*string) *CreateBlockRequestScopesBlockScopeProduct {
	s.CodeNames = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeProduct) SetKey(v string) *CreateBlockRequestScopesBlockScopeProduct {
	s.Key = &v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeProduct) SetRelations(v []*CreateBlockRequestScopesBlockScopeProductRelations) *CreateBlockRequestScopesBlockScopeProduct {
	s.Relations = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeProduct) SetViewCode(v []*string) *CreateBlockRequestScopesBlockScopeProduct {
	s.ViewCode = v
	return s
}

type CreateBlockRequestScopesBlockScopeProductRelations struct {
	AppCodes    []*string `json:"AppCodes,omitempty" xml:"AppCodes,omitempty" type:"Repeated"`
	LabelCodes  []*string `json:"LabelCodes,omitempty" xml:"LabelCodes,omitempty" type:"Repeated"`
	RelateCodes []*string `json:"RelateCodes,omitempty" xml:"RelateCodes,omitempty" type:"Repeated"`
	SelfCode    *string   `json:"SelfCode,omitempty" xml:"SelfCode,omitempty"`
}

func (s CreateBlockRequestScopesBlockScopeProductRelations) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockRequestScopesBlockScopeProductRelations) GoString() string {
	return s.String()
}

func (s *CreateBlockRequestScopesBlockScopeProductRelations) SetAppCodes(v []*string) *CreateBlockRequestScopesBlockScopeProductRelations {
	s.AppCodes = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeProductRelations) SetLabelCodes(v []*string) *CreateBlockRequestScopesBlockScopeProductRelations {
	s.LabelCodes = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeProductRelations) SetRelateCodes(v []*string) *CreateBlockRequestScopesBlockScopeProductRelations {
	s.RelateCodes = v
	return s
}

func (s *CreateBlockRequestScopesBlockScopeProductRelations) SetSelfCode(v string) *CreateBlockRequestScopesBlockScopeProductRelations {
	s.SelfCode = &v
	return s
}

type CreateBlockResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBlockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBlockResponseBody) SetCode(v int32) *CreateBlockResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBlockResponseBody) SetData(v int64) *CreateBlockResponseBody {
	s.Data = &v
	return s
}

func (s *CreateBlockResponseBody) SetMessage(v string) *CreateBlockResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBlockResponseBody) SetRequestId(v string) *CreateBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBlockResponseBody) SetSuccess(v bool) *CreateBlockResponseBody {
	s.Success = &v
	return s
}

type CreateBlockResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBlockResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBlockResponse) GoString() string {
	return s.String()
}

func (s *CreateBlockResponse) SetHeaders(v map[string]*string) *CreateBlockResponse {
	s.Headers = v
	return s
}

func (s *CreateBlockResponse) SetStatusCode(v int32) *CreateBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBlockResponse) SetBody(v *CreateBlockResponseBody) *CreateBlockResponse {
	s.Body = v
	return s
}

type CreateOperatorRequest struct {
	AuthKey      *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign     *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	BgObject     *string `json:"BgObject,omitempty" xml:"BgObject,omitempty"`
	BgSystem     *string `json:"BgSystem,omitempty" xml:"BgSystem,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CurEmpId     *string `json:"CurEmpId,omitempty" xml:"CurEmpId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NoCheck      *bool   `json:"NoCheck,omitempty" xml:"NoCheck,omitempty"`
	NoRisk       *bool   `json:"NoRisk,omitempty" xml:"NoRisk,omitempty"`
	ReqTimestamp *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
}

func (s CreateOperatorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOperatorRequest) GoString() string {
	return s.String()
}

func (s *CreateOperatorRequest) SetAuthKey(v string) *CreateOperatorRequest {
	s.AuthKey = &v
	return s
}

func (s *CreateOperatorRequest) SetAuthSign(v string) *CreateOperatorRequest {
	s.AuthSign = &v
	return s
}

func (s *CreateOperatorRequest) SetBgObject(v string) *CreateOperatorRequest {
	s.BgObject = &v
	return s
}

func (s *CreateOperatorRequest) SetBgSystem(v string) *CreateOperatorRequest {
	s.BgSystem = &v
	return s
}

func (s *CreateOperatorRequest) SetCode(v string) *CreateOperatorRequest {
	s.Code = &v
	return s
}

func (s *CreateOperatorRequest) SetCurEmpId(v string) *CreateOperatorRequest {
	s.CurEmpId = &v
	return s
}

func (s *CreateOperatorRequest) SetName(v string) *CreateOperatorRequest {
	s.Name = &v
	return s
}

func (s *CreateOperatorRequest) SetNoCheck(v bool) *CreateOperatorRequest {
	s.NoCheck = &v
	return s
}

func (s *CreateOperatorRequest) SetNoRisk(v bool) *CreateOperatorRequest {
	s.NoRisk = &v
	return s
}

func (s *CreateOperatorRequest) SetReqTimestamp(v int64) *CreateOperatorRequest {
	s.ReqTimestamp = &v
	return s
}

type CreateOperatorResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateOperatorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOperatorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOperatorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOperatorResponseBody) SetCode(v int32) *CreateOperatorResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOperatorResponseBody) SetData(v *CreateOperatorResponseBodyData) *CreateOperatorResponseBody {
	s.Data = v
	return s
}

func (s *CreateOperatorResponseBody) SetMessage(v string) *CreateOperatorResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOperatorResponseBody) SetRequestId(v string) *CreateOperatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOperatorResponseBody) SetSuccess(v bool) *CreateOperatorResponseBody {
	s.Success = &v
	return s
}

type CreateOperatorResponseBodyData struct {
	ApproveStrategyId *int64 `json:"ApproveStrategyId,omitempty" xml:"ApproveStrategyId,omitempty"`
	RuleId            *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateOperatorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateOperatorResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateOperatorResponseBodyData) SetApproveStrategyId(v int64) *CreateOperatorResponseBodyData {
	s.ApproveStrategyId = &v
	return s
}

func (s *CreateOperatorResponseBodyData) SetRuleId(v int64) *CreateOperatorResponseBodyData {
	s.RuleId = &v
	return s
}

type CreateOperatorResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOperatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOperatorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOperatorResponse) GoString() string {
	return s.String()
}

func (s *CreateOperatorResponse) SetHeaders(v map[string]*string) *CreateOperatorResponse {
	s.Headers = v
	return s
}

func (s *CreateOperatorResponse) SetStatusCode(v int32) *CreateOperatorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOperatorResponse) SetBody(v *CreateOperatorResponseBody) *CreateOperatorResponse {
	s.Body = v
	return s
}

type QueryRequest struct {
	AuthKey       *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign      *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	NeedValidate  *bool   `json:"NeedValidate,omitempty" xml:"NeedValidate,omitempty"`
	QueryType     *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ReqTimestamp  *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	SourceOrderId *string `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
}

func (s QueryRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRequest) GoString() string {
	return s.String()
}

func (s *QueryRequest) SetAuthKey(v string) *QueryRequest {
	s.AuthKey = &v
	return s
}

func (s *QueryRequest) SetAuthSign(v string) *QueryRequest {
	s.AuthSign = &v
	return s
}

func (s *QueryRequest) SetNeedValidate(v bool) *QueryRequest {
	s.NeedValidate = &v
	return s
}

func (s *QueryRequest) SetQueryType(v string) *QueryRequest {
	s.QueryType = &v
	return s
}

func (s *QueryRequest) SetReqTimestamp(v int64) *QueryRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *QueryRequest) SetSourceOrderId(v string) *QueryRequest {
	s.SourceOrderId = &v
	return s
}

type QueryResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryResponseBody) SetCode(v int32) *QueryResponseBody {
	s.Code = &v
	return s
}

func (s *QueryResponseBody) SetData(v string) *QueryResponseBody {
	s.Data = &v
	return s
}

func (s *QueryResponseBody) SetMessage(v string) *QueryResponseBody {
	s.Message = &v
	return s
}

func (s *QueryResponseBody) SetRequestId(v string) *QueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryResponseBody) SetSuccess(v bool) *QueryResponseBody {
	s.Success = &v
	return s
}

type QueryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryResponse) GoString() string {
	return s.String()
}

func (s *QueryResponse) SetHeaders(v map[string]*string) *QueryResponse {
	s.Headers = v
	return s
}

func (s *QueryResponse) SetStatusCode(v int32) *QueryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryResponse) SetBody(v *QueryResponseBody) *QueryResponse {
	s.Body = v
	return s
}

type QueryApproveFlowRequest struct {
	AuthKey       *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign      *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	ReqTimestamp  *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	SourceOrderId *string `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
	Stage         *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
}

func (s QueryApproveFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryApproveFlowRequest) GoString() string {
	return s.String()
}

func (s *QueryApproveFlowRequest) SetAuthKey(v string) *QueryApproveFlowRequest {
	s.AuthKey = &v
	return s
}

func (s *QueryApproveFlowRequest) SetAuthSign(v string) *QueryApproveFlowRequest {
	s.AuthSign = &v
	return s
}

func (s *QueryApproveFlowRequest) SetReqTimestamp(v int64) *QueryApproveFlowRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *QueryApproveFlowRequest) SetSourceOrderId(v string) *QueryApproveFlowRequest {
	s.SourceOrderId = &v
	return s
}

func (s *QueryApproveFlowRequest) SetStage(v string) *QueryApproveFlowRequest {
	s.Stage = &v
	return s
}

type QueryApproveFlowResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryApproveFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryApproveFlowResponseBody) GoString() string {
	return s.String()
}

func (s *QueryApproveFlowResponseBody) SetCode(v int32) *QueryApproveFlowResponseBody {
	s.Code = &v
	return s
}

func (s *QueryApproveFlowResponseBody) SetData(v string) *QueryApproveFlowResponseBody {
	s.Data = &v
	return s
}

func (s *QueryApproveFlowResponseBody) SetMessage(v string) *QueryApproveFlowResponseBody {
	s.Message = &v
	return s
}

func (s *QueryApproveFlowResponseBody) SetRequestId(v string) *QueryApproveFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryApproveFlowResponseBody) SetSuccess(v bool) *QueryApproveFlowResponseBody {
	s.Success = &v
	return s
}

type QueryApproveFlowResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryApproveFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryApproveFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryApproveFlowResponse) GoString() string {
	return s.String()
}

func (s *QueryApproveFlowResponse) SetHeaders(v map[string]*string) *QueryApproveFlowResponse {
	s.Headers = v
	return s
}

func (s *QueryApproveFlowResponse) SetStatusCode(v int32) *QueryApproveFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryApproveFlowResponse) SetBody(v *QueryApproveFlowResponseBody) *QueryApproveFlowResponse {
	s.Body = v
	return s
}

type QueryBlockEventRequest struct {
	AuthKey      *string                             `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign     *string                             `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	BgSystemName *string                             `json:"BgSystemName,omitempty" xml:"BgSystemName,omitempty"`
	BlockHarm    *string                             `json:"BlockHarm,omitempty" xml:"BlockHarm,omitempty"`
	Category     *string                             `json:"Category,omitempty" xml:"Category,omitempty"`
	DeptNo       *string                             `json:"DeptNo,omitempty" xml:"DeptNo,omitempty"`
	EndTime      *int64                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit        *int32                              `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NeedRule     *bool                               `json:"NeedRule,omitempty" xml:"NeedRule,omitempty"`
	Page         *int32                              `json:"Page,omitempty" xml:"Page,omitempty"`
	ProductCodes []*string                           `json:"ProductCodes,omitempty" xml:"ProductCodes,omitempty" type:"Repeated"`
	RegionReqs   []*QueryBlockEventRequestRegionReqs `json:"RegionReqs,omitempty" xml:"RegionReqs,omitempty" type:"Repeated"`
	ReqTimestamp *int64                              `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	Scope        []*string                           `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
	StartTime    *int64                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryBlockEventRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBlockEventRequest) GoString() string {
	return s.String()
}

func (s *QueryBlockEventRequest) SetAuthKey(v string) *QueryBlockEventRequest {
	s.AuthKey = &v
	return s
}

func (s *QueryBlockEventRequest) SetAuthSign(v string) *QueryBlockEventRequest {
	s.AuthSign = &v
	return s
}

func (s *QueryBlockEventRequest) SetBgSystemName(v string) *QueryBlockEventRequest {
	s.BgSystemName = &v
	return s
}

func (s *QueryBlockEventRequest) SetBlockHarm(v string) *QueryBlockEventRequest {
	s.BlockHarm = &v
	return s
}

func (s *QueryBlockEventRequest) SetCategory(v string) *QueryBlockEventRequest {
	s.Category = &v
	return s
}

func (s *QueryBlockEventRequest) SetDeptNo(v string) *QueryBlockEventRequest {
	s.DeptNo = &v
	return s
}

func (s *QueryBlockEventRequest) SetEndTime(v int64) *QueryBlockEventRequest {
	s.EndTime = &v
	return s
}

func (s *QueryBlockEventRequest) SetLimit(v int32) *QueryBlockEventRequest {
	s.Limit = &v
	return s
}

func (s *QueryBlockEventRequest) SetNeedRule(v bool) *QueryBlockEventRequest {
	s.NeedRule = &v
	return s
}

func (s *QueryBlockEventRequest) SetPage(v int32) *QueryBlockEventRequest {
	s.Page = &v
	return s
}

func (s *QueryBlockEventRequest) SetProductCodes(v []*string) *QueryBlockEventRequest {
	s.ProductCodes = v
	return s
}

func (s *QueryBlockEventRequest) SetRegionReqs(v []*QueryBlockEventRequestRegionReqs) *QueryBlockEventRequest {
	s.RegionReqs = v
	return s
}

func (s *QueryBlockEventRequest) SetReqTimestamp(v int64) *QueryBlockEventRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *QueryBlockEventRequest) SetScope(v []*string) *QueryBlockEventRequest {
	s.Scope = v
	return s
}

func (s *QueryBlockEventRequest) SetStartTime(v int64) *QueryBlockEventRequest {
	s.StartTime = &v
	return s
}

type QueryBlockEventRequestRegionReqs struct {
	ProductCode *string   `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Regions     []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s QueryBlockEventRequestRegionReqs) String() string {
	return tea.Prettify(s)
}

func (s QueryBlockEventRequestRegionReqs) GoString() string {
	return s.String()
}

func (s *QueryBlockEventRequestRegionReqs) SetProductCode(v string) *QueryBlockEventRequestRegionReqs {
	s.ProductCode = &v
	return s
}

func (s *QueryBlockEventRequestRegionReqs) SetRegions(v []*string) *QueryBlockEventRequestRegionReqs {
	s.Regions = v
	return s
}

type QueryBlockEventResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryBlockEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryBlockEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBlockEventResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBlockEventResponseBody) SetCode(v int32) *QueryBlockEventResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBlockEventResponseBody) SetData(v *QueryBlockEventResponseBodyData) *QueryBlockEventResponseBody {
	s.Data = v
	return s
}

func (s *QueryBlockEventResponseBody) SetMessage(v string) *QueryBlockEventResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBlockEventResponseBody) SetRequestId(v string) *QueryBlockEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBlockEventResponseBody) SetSuccess(v bool) *QueryBlockEventResponseBody {
	s.Success = &v
	return s
}

type QueryBlockEventResponseBodyData struct {
	DataInfo   []*QueryBlockEventResponseBodyDataDataInfo `json:"DataInfo,omitempty" xml:"DataInfo,omitempty" type:"Repeated"`
	ExtraInfo  map[string]*string                         `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Pagination *QueryBlockEventResponseBodyDataPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	Total      *int64                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryBlockEventResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryBlockEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBlockEventResponseBodyData) SetDataInfo(v []*QueryBlockEventResponseBodyDataDataInfo) *QueryBlockEventResponseBodyData {
	s.DataInfo = v
	return s
}

func (s *QueryBlockEventResponseBodyData) SetExtraInfo(v map[string]*string) *QueryBlockEventResponseBodyData {
	s.ExtraInfo = v
	return s
}

func (s *QueryBlockEventResponseBodyData) SetPagination(v *QueryBlockEventResponseBodyDataPagination) *QueryBlockEventResponseBodyData {
	s.Pagination = v
	return s
}

func (s *QueryBlockEventResponseBodyData) SetTotal(v int64) *QueryBlockEventResponseBodyData {
	s.Total = &v
	return s
}

type QueryBlockEventResponseBodyDataDataInfo struct {
	EmpId       *string                                              `json:"EmpId,omitempty" xml:"EmpId,omitempty"`
	EndTime     *int64                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventTimes  []*QueryBlockEventResponseBodyDataDataInfoEventTimes `json:"EventTimes,omitempty" xml:"EventTimes,omitempty" type:"Repeated"`
	GmtCreate   *int64                                               `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64                                               `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	Reason      *string                                              `json:"Reason,omitempty" xml:"Reason,omitempty"`
	StartTime   *int64                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Title       *string                                              `json:"Title,omitempty" xml:"Title,omitempty"`
	Url         *string                                              `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryBlockEventResponseBodyDataDataInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryBlockEventResponseBodyDataDataInfo) GoString() string {
	return s.String()
}

func (s *QueryBlockEventResponseBodyDataDataInfo) SetEmpId(v string) *QueryBlockEventResponseBodyDataDataInfo {
	s.EmpId = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfo) SetEndTime(v int64) *QueryBlockEventResponseBodyDataDataInfo {
	s.EndTime = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfo) SetEventTimes(v []*QueryBlockEventResponseBodyDataDataInfoEventTimes) *QueryBlockEventResponseBodyDataDataInfo {
	s.EventTimes = v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfo) SetGmtCreate(v int64) *QueryBlockEventResponseBodyDataDataInfo {
	s.GmtCreate = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfo) SetGmtModified(v int64) *QueryBlockEventResponseBodyDataDataInfo {
	s.GmtModified = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfo) SetId(v int64) *QueryBlockEventResponseBodyDataDataInfo {
	s.Id = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfo) SetReason(v string) *QueryBlockEventResponseBodyDataDataInfo {
	s.Reason = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfo) SetStartTime(v int64) *QueryBlockEventResponseBodyDataDataInfo {
	s.StartTime = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfo) SetTitle(v string) *QueryBlockEventResponseBodyDataDataInfo {
	s.Title = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfo) SetUrl(v string) *QueryBlockEventResponseBodyDataDataInfo {
	s.Url = &v
	return s
}

type QueryBlockEventResponseBodyDataDataInfoEventTimes struct {
	EndTime   *int64                                                   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Express   *string                                                  `json:"Express,omitempty" xml:"Express,omitempty"`
	Rule      []*QueryBlockEventResponseBodyDataDataInfoEventTimesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
	StartTime *int64                                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryBlockEventResponseBodyDataDataInfoEventTimes) String() string {
	return tea.Prettify(s)
}

func (s QueryBlockEventResponseBodyDataDataInfoEventTimes) GoString() string {
	return s.String()
}

func (s *QueryBlockEventResponseBodyDataDataInfoEventTimes) SetEndTime(v int64) *QueryBlockEventResponseBodyDataDataInfoEventTimes {
	s.EndTime = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfoEventTimes) SetExpress(v string) *QueryBlockEventResponseBodyDataDataInfoEventTimes {
	s.Express = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfoEventTimes) SetRule(v []*QueryBlockEventResponseBodyDataDataInfoEventTimesRule) *QueryBlockEventResponseBodyDataDataInfoEventTimes {
	s.Rule = v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfoEventTimes) SetStartTime(v int64) *QueryBlockEventResponseBodyDataDataInfoEventTimes {
	s.StartTime = &v
	return s
}

type QueryBlockEventResponseBodyDataDataInfoEventTimesRule struct {
	Level1 *string `json:"Level1,omitempty" xml:"Level1,omitempty"`
	Level2 *string `json:"Level2,omitempty" xml:"Level2,omitempty"`
	Level3 *string `json:"Level3,omitempty" xml:"Level3,omitempty"`
	Level4 *string `json:"Level4,omitempty" xml:"Level4,omitempty"`
	Level5 *string `json:"Level5,omitempty" xml:"Level5,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryBlockEventResponseBodyDataDataInfoEventTimesRule) String() string {
	return tea.Prettify(s)
}

func (s QueryBlockEventResponseBodyDataDataInfoEventTimesRule) GoString() string {
	return s.String()
}

func (s *QueryBlockEventResponseBodyDataDataInfoEventTimesRule) SetLevel1(v string) *QueryBlockEventResponseBodyDataDataInfoEventTimesRule {
	s.Level1 = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfoEventTimesRule) SetLevel2(v string) *QueryBlockEventResponseBodyDataDataInfoEventTimesRule {
	s.Level2 = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfoEventTimesRule) SetLevel3(v string) *QueryBlockEventResponseBodyDataDataInfoEventTimesRule {
	s.Level3 = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfoEventTimesRule) SetLevel4(v string) *QueryBlockEventResponseBodyDataDataInfoEventTimesRule {
	s.Level4 = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfoEventTimesRule) SetLevel5(v string) *QueryBlockEventResponseBodyDataDataInfoEventTimesRule {
	s.Level5 = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataDataInfoEventTimesRule) SetType(v string) *QueryBlockEventResponseBodyDataDataInfoEventTimesRule {
	s.Type = &v
	return s
}

type QueryBlockEventResponseBodyDataPagination struct {
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Page  *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s QueryBlockEventResponseBodyDataPagination) String() string {
	return tea.Prettify(s)
}

func (s QueryBlockEventResponseBodyDataPagination) GoString() string {
	return s.String()
}

func (s *QueryBlockEventResponseBodyDataPagination) SetLimit(v int32) *QueryBlockEventResponseBodyDataPagination {
	s.Limit = &v
	return s
}

func (s *QueryBlockEventResponseBodyDataPagination) SetPage(v int32) *QueryBlockEventResponseBodyDataPagination {
	s.Page = &v
	return s
}

type QueryBlockEventResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBlockEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBlockEventResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBlockEventResponse) GoString() string {
	return s.String()
}

func (s *QueryBlockEventResponse) SetHeaders(v map[string]*string) *QueryBlockEventResponse {
	s.Headers = v
	return s
}

func (s *QueryBlockEventResponse) SetStatusCode(v int32) *QueryBlockEventResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBlockEventResponse) SetBody(v *QueryBlockEventResponseBody) *QueryBlockEventResponse {
	s.Body = v
	return s
}

type QueryChangeInfoRequest struct {
	AuthKey       *string                          `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign      *string                          `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	Az            []*string                        `json:"Az,omitempty" xml:"Az,omitempty" type:"Repeated"`
	BgVid         *string                          `json:"BgVid,omitempty" xml:"BgVid,omitempty"`
	BuId          *string                          `json:"BuId,omitempty" xml:"BuId,omitempty"`
	ChangeSystem  *string                          `json:"ChangeSystem,omitempty" xml:"ChangeSystem,omitempty"`
	EndTime       *int64                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Keyword       *string                          `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	LevelTree     *QueryChangeInfoRequestLevelTree `json:"LevelTree,omitempty" xml:"LevelTree,omitempty" type:"Struct"`
	Limit         *int32                           `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Page          *int32                           `json:"Page,omitempty" xml:"Page,omitempty"`
	Product       []*string                        `json:"Product,omitempty" xml:"Product,omitempty" type:"Repeated"`
	Region        []*string                        `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
	ReqTimestamp  *int64                           `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	Source        *string                          `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceOrderId *string                          `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
	StartTime     *int64                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type          *string                          `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryChangeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryChangeInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryChangeInfoRequest) SetAuthKey(v string) *QueryChangeInfoRequest {
	s.AuthKey = &v
	return s
}

func (s *QueryChangeInfoRequest) SetAuthSign(v string) *QueryChangeInfoRequest {
	s.AuthSign = &v
	return s
}

func (s *QueryChangeInfoRequest) SetAz(v []*string) *QueryChangeInfoRequest {
	s.Az = v
	return s
}

func (s *QueryChangeInfoRequest) SetBgVid(v string) *QueryChangeInfoRequest {
	s.BgVid = &v
	return s
}

func (s *QueryChangeInfoRequest) SetBuId(v string) *QueryChangeInfoRequest {
	s.BuId = &v
	return s
}

func (s *QueryChangeInfoRequest) SetChangeSystem(v string) *QueryChangeInfoRequest {
	s.ChangeSystem = &v
	return s
}

func (s *QueryChangeInfoRequest) SetEndTime(v int64) *QueryChangeInfoRequest {
	s.EndTime = &v
	return s
}

func (s *QueryChangeInfoRequest) SetKeyword(v string) *QueryChangeInfoRequest {
	s.Keyword = &v
	return s
}

func (s *QueryChangeInfoRequest) SetLevelTree(v *QueryChangeInfoRequestLevelTree) *QueryChangeInfoRequest {
	s.LevelTree = v
	return s
}

func (s *QueryChangeInfoRequest) SetLimit(v int32) *QueryChangeInfoRequest {
	s.Limit = &v
	return s
}

func (s *QueryChangeInfoRequest) SetPage(v int32) *QueryChangeInfoRequest {
	s.Page = &v
	return s
}

func (s *QueryChangeInfoRequest) SetProduct(v []*string) *QueryChangeInfoRequest {
	s.Product = v
	return s
}

func (s *QueryChangeInfoRequest) SetRegion(v []*string) *QueryChangeInfoRequest {
	s.Region = v
	return s
}

func (s *QueryChangeInfoRequest) SetReqTimestamp(v int64) *QueryChangeInfoRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *QueryChangeInfoRequest) SetSource(v string) *QueryChangeInfoRequest {
	s.Source = &v
	return s
}

func (s *QueryChangeInfoRequest) SetSourceOrderId(v string) *QueryChangeInfoRequest {
	s.SourceOrderId = &v
	return s
}

func (s *QueryChangeInfoRequest) SetStartTime(v int64) *QueryChangeInfoRequest {
	s.StartTime = &v
	return s
}

func (s *QueryChangeInfoRequest) SetType(v string) *QueryChangeInfoRequest {
	s.Type = &v
	return s
}

type QueryChangeInfoRequestLevelTree struct {
	DataType *string                                    `json:"DataType,omitempty" xml:"DataType,omitempty"`
	TreeData []*QueryChangeInfoRequestLevelTreeTreeData `json:"TreeData,omitempty" xml:"TreeData,omitempty" type:"Repeated"`
}

func (s QueryChangeInfoRequestLevelTree) String() string {
	return tea.Prettify(s)
}

func (s QueryChangeInfoRequestLevelTree) GoString() string {
	return s.String()
}

func (s *QueryChangeInfoRequestLevelTree) SetDataType(v string) *QueryChangeInfoRequestLevelTree {
	s.DataType = &v
	return s
}

func (s *QueryChangeInfoRequestLevelTree) SetTreeData(v []*QueryChangeInfoRequestLevelTreeTreeData) *QueryChangeInfoRequestLevelTree {
	s.TreeData = v
	return s
}

type QueryChangeInfoRequestLevelTreeTreeData struct {
	Data        []interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	DataSubType *string       `json:"DataSubType,omitempty" xml:"DataSubType,omitempty"`
	Value       []*string     `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s QueryChangeInfoRequestLevelTreeTreeData) String() string {
	return tea.Prettify(s)
}

func (s QueryChangeInfoRequestLevelTreeTreeData) GoString() string {
	return s.String()
}

func (s *QueryChangeInfoRequestLevelTreeTreeData) SetData(v []interface{}) *QueryChangeInfoRequestLevelTreeTreeData {
	s.Data = v
	return s
}

func (s *QueryChangeInfoRequestLevelTreeTreeData) SetDataSubType(v string) *QueryChangeInfoRequestLevelTreeTreeData {
	s.DataSubType = &v
	return s
}

func (s *QueryChangeInfoRequestLevelTreeTreeData) SetValue(v []*string) *QueryChangeInfoRequestLevelTreeTreeData {
	s.Value = v
	return s
}

type QueryChangeInfoResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryChangeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryChangeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryChangeInfoResponseBody) SetCode(v int32) *QueryChangeInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryChangeInfoResponseBody) SetData(v string) *QueryChangeInfoResponseBody {
	s.Data = &v
	return s
}

func (s *QueryChangeInfoResponseBody) SetMessage(v string) *QueryChangeInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryChangeInfoResponseBody) SetRequestId(v string) *QueryChangeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryChangeInfoResponseBody) SetSuccess(v bool) *QueryChangeInfoResponseBody {
	s.Success = &v
	return s
}

type QueryChangeInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryChangeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryChangeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryChangeInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryChangeInfoResponse) SetHeaders(v map[string]*string) *QueryChangeInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryChangeInfoResponse) SetStatusCode(v int32) *QueryChangeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryChangeInfoResponse) SetBody(v *QueryChangeInfoResponseBody) *QueryChangeInfoResponse {
	s.Body = v
	return s
}

type QueryCheckInfoRequest struct {
	AuthKey       *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign      *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	ReqTimestamp  *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	SourceOrderId *string `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
}

func (s QueryCheckInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCheckInfoRequest) SetAuthKey(v string) *QueryCheckInfoRequest {
	s.AuthKey = &v
	return s
}

func (s *QueryCheckInfoRequest) SetAuthSign(v string) *QueryCheckInfoRequest {
	s.AuthSign = &v
	return s
}

func (s *QueryCheckInfoRequest) SetReqTimestamp(v int64) *QueryCheckInfoRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *QueryCheckInfoRequest) SetSourceOrderId(v string) *QueryCheckInfoRequest {
	s.SourceOrderId = &v
	return s
}

type QueryCheckInfoResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryCheckInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCheckInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCheckInfoResponseBody) SetCode(v int32) *QueryCheckInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCheckInfoResponseBody) SetData(v *QueryCheckInfoResponseBodyData) *QueryCheckInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryCheckInfoResponseBody) SetMessage(v string) *QueryCheckInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCheckInfoResponseBody) SetRequestId(v string) *QueryCheckInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCheckInfoResponseBody) SetSuccess(v bool) *QueryCheckInfoResponseBody {
	s.Success = &v
	return s
}

type QueryCheckInfoResponseBodyData struct {
	CheckDetailList []*QueryCheckInfoResponseBodyDataCheckDetailList `json:"CheckDetailList,omitempty" xml:"CheckDetailList,omitempty" type:"Repeated"`
	CheckResultUrl  *string                                          `json:"CheckResultUrl,omitempty" xml:"CheckResultUrl,omitempty"`
}

func (s QueryCheckInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCheckInfoResponseBodyData) SetCheckDetailList(v []*QueryCheckInfoResponseBodyDataCheckDetailList) *QueryCheckInfoResponseBodyData {
	s.CheckDetailList = v
	return s
}

func (s *QueryCheckInfoResponseBodyData) SetCheckResultUrl(v string) *QueryCheckInfoResponseBodyData {
	s.CheckResultUrl = &v
	return s
}

type QueryCheckInfoResponseBodyDataCheckDetailList struct {
	BlockRule       []*QueryCheckInfoResponseBodyDataCheckDetailListBlockRule `json:"BlockRule,omitempty" xml:"BlockRule,omitempty" type:"Repeated"`
	CheckholdReason *string                                                   `json:"CheckholdReason,omitempty" xml:"CheckholdReason,omitempty"`
	Desc            *string                                                   `json:"Desc,omitempty" xml:"Desc,omitempty"`
	PicInfo         *string                                                   `json:"PicInfo,omitempty" xml:"PicInfo,omitempty"`
	RiskExplain     *string                                                   `json:"RiskExplain,omitempty" xml:"RiskExplain,omitempty"`
	Title           *string                                                   `json:"Title,omitempty" xml:"Title,omitempty"`
	Url             *string                                                   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryCheckInfoResponseBodyDataCheckDetailList) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckInfoResponseBodyDataCheckDetailList) GoString() string {
	return s.String()
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailList) SetBlockRule(v []*QueryCheckInfoResponseBodyDataCheckDetailListBlockRule) *QueryCheckInfoResponseBodyDataCheckDetailList {
	s.BlockRule = v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailList) SetCheckholdReason(v string) *QueryCheckInfoResponseBodyDataCheckDetailList {
	s.CheckholdReason = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailList) SetDesc(v string) *QueryCheckInfoResponseBodyDataCheckDetailList {
	s.Desc = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailList) SetPicInfo(v string) *QueryCheckInfoResponseBodyDataCheckDetailList {
	s.PicInfo = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailList) SetRiskExplain(v string) *QueryCheckInfoResponseBodyDataCheckDetailList {
	s.RiskExplain = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailList) SetTitle(v string) *QueryCheckInfoResponseBodyDataCheckDetailList {
	s.Title = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailList) SetUrl(v string) *QueryCheckInfoResponseBodyDataCheckDetailList {
	s.Url = &v
	return s
}

type QueryCheckInfoResponseBodyDataCheckDetailListBlockRule struct {
	BlockHarm      *string                                                                `json:"BlockHarm,omitempty" xml:"BlockHarm,omitempty"`
	BlockId        *int64                                                                 `json:"BlockId,omitempty" xml:"BlockId,omitempty"`
	Express        *string                                                                `json:"Express,omitempty" xml:"Express,omitempty"`
	ScopeEndTime   *int64                                                                 `json:"ScopeEndTime,omitempty" xml:"ScopeEndTime,omitempty"`
	ScopeNodeList  []*QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList `json:"ScopeNodeList,omitempty" xml:"ScopeNodeList,omitempty" type:"Repeated"`
	ScopeRuleId    *int64                                                                 `json:"ScopeRuleId,omitempty" xml:"ScopeRuleId,omitempty"`
	ScopeStartTime *int64                                                                 `json:"ScopeStartTime,omitempty" xml:"ScopeStartTime,omitempty"`
}

func (s QueryCheckInfoResponseBodyDataCheckDetailListBlockRule) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckInfoResponseBodyDataCheckDetailListBlockRule) GoString() string {
	return s.String()
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRule) SetBlockHarm(v string) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRule {
	s.BlockHarm = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRule) SetBlockId(v int64) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRule {
	s.BlockId = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRule) SetExpress(v string) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRule {
	s.Express = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRule) SetScopeEndTime(v int64) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRule {
	s.ScopeEndTime = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRule) SetScopeNodeList(v []*QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRule {
	s.ScopeNodeList = v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRule) SetScopeRuleId(v int64) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRule {
	s.ScopeRuleId = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRule) SetScopeStartTime(v int64) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRule {
	s.ScopeStartTime = &v
	return s
}

type QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList struct {
	LeafLevel *string `json:"LeafLevel,omitempty" xml:"LeafLevel,omitempty"`
	Level1    *string `json:"Level1,omitempty" xml:"Level1,omitempty"`
	Level2    *string `json:"Level2,omitempty" xml:"Level2,omitempty"`
	Level3    *string `json:"Level3,omitempty" xml:"Level3,omitempty"`
	Level4    *string `json:"Level4,omitempty" xml:"Level4,omitempty"`
	Level5    *string `json:"Level5,omitempty" xml:"Level5,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
	RuleId    *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList) GoString() string {
	return s.String()
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList) SetLeafLevel(v string) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList {
	s.LeafLevel = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList) SetLevel1(v string) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList {
	s.Level1 = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList) SetLevel2(v string) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList {
	s.Level2 = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList) SetLevel3(v string) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList {
	s.Level3 = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList) SetLevel4(v string) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList {
	s.Level4 = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList) SetLevel5(v string) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList {
	s.Level5 = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList) SetPath(v string) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList {
	s.Path = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList) SetRuleId(v int64) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList {
	s.RuleId = &v
	return s
}

func (s *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList) SetType(v string) *QueryCheckInfoResponseBodyDataCheckDetailListBlockRuleScopeNodeList {
	s.Type = &v
	return s
}

type QueryCheckInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCheckInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCheckInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCheckInfoResponse) SetHeaders(v map[string]*string) *QueryCheckInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCheckInfoResponse) SetStatusCode(v int32) *QueryCheckInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCheckInfoResponse) SetBody(v *QueryCheckInfoResponseBody) *QueryCheckInfoResponse {
	s.Body = v
	return s
}

type QueryCustomerRequest struct {
	AuthKey      *string   `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign     *string   `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	Product      []*string `json:"Product,omitempty" xml:"Product,omitempty" type:"Repeated"`
	ReqTimestamp *int64    `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	Type         *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomerRequest) SetAuthKey(v string) *QueryCustomerRequest {
	s.AuthKey = &v
	return s
}

func (s *QueryCustomerRequest) SetAuthSign(v string) *QueryCustomerRequest {
	s.AuthSign = &v
	return s
}

func (s *QueryCustomerRequest) SetProduct(v []*string) *QueryCustomerRequest {
	s.Product = v
	return s
}

func (s *QueryCustomerRequest) SetReqTimestamp(v int64) *QueryCustomerRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *QueryCustomerRequest) SetType(v string) *QueryCustomerRequest {
	s.Type = &v
	return s
}

type QueryCustomerResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryCustomerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomerResponseBody) SetCode(v int32) *QueryCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCustomerResponseBody) SetData(v []*QueryCustomerResponseBodyData) *QueryCustomerResponseBody {
	s.Data = v
	return s
}

func (s *QueryCustomerResponseBody) SetMessage(v string) *QueryCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCustomerResponseBody) SetRequestId(v string) *QueryCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCustomerResponseBody) SetSuccess(v bool) *QueryCustomerResponseBody {
	s.Success = &v
	return s
}

type QueryCustomerResponseBodyData struct {
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid     *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryCustomerResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCustomerResponseBodyData) SetProduct(v string) *QueryCustomerResponseBodyData {
	s.Product = &v
	return s
}

func (s *QueryCustomerResponseBodyData) SetType(v string) *QueryCustomerResponseBodyData {
	s.Type = &v
	return s
}

func (s *QueryCustomerResponseBodyData) SetUid(v string) *QueryCustomerResponseBodyData {
	s.Uid = &v
	return s
}

type QueryCustomerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomerResponse) SetHeaders(v map[string]*string) *QueryCustomerResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomerResponse) SetStatusCode(v int32) *QueryCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomerResponse) SetBody(v *QueryCustomerResponseBody) *QueryCustomerResponse {
	s.Body = v
	return s
}

type QueryExecuteInfoRequest struct {
	AuthKey       *string                           `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign      *string                           `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	Az            []*string                         `json:"Az,omitempty" xml:"Az,omitempty" type:"Repeated"`
	BgVid         *string                           `json:"BgVid,omitempty" xml:"BgVid,omitempty"`
	BuId          *string                           `json:"BuId,omitempty" xml:"BuId,omitempty"`
	EndTime       *int64                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExVid         *string                           `json:"ExVid,omitempty" xml:"ExVid,omitempty"`
	Keyword       *string                           `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	LevelTree     *QueryExecuteInfoRequestLevelTree `json:"LevelTree,omitempty" xml:"LevelTree,omitempty" type:"Struct"`
	Limit         *int32                            `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Page          *int32                            `json:"Page,omitempty" xml:"Page,omitempty"`
	Product       []*string                         `json:"Product,omitempty" xml:"Product,omitempty" type:"Repeated"`
	Region        []*string                         `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
	ReqTimestamp  *int64                            `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	Source        *string                           `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceOrderId *string                           `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
	StartTime     *int64                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryExecuteInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryExecuteInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryExecuteInfoRequest) SetAuthKey(v string) *QueryExecuteInfoRequest {
	s.AuthKey = &v
	return s
}

func (s *QueryExecuteInfoRequest) SetAuthSign(v string) *QueryExecuteInfoRequest {
	s.AuthSign = &v
	return s
}

func (s *QueryExecuteInfoRequest) SetAz(v []*string) *QueryExecuteInfoRequest {
	s.Az = v
	return s
}

func (s *QueryExecuteInfoRequest) SetBgVid(v string) *QueryExecuteInfoRequest {
	s.BgVid = &v
	return s
}

func (s *QueryExecuteInfoRequest) SetBuId(v string) *QueryExecuteInfoRequest {
	s.BuId = &v
	return s
}

func (s *QueryExecuteInfoRequest) SetEndTime(v int64) *QueryExecuteInfoRequest {
	s.EndTime = &v
	return s
}

func (s *QueryExecuteInfoRequest) SetExVid(v string) *QueryExecuteInfoRequest {
	s.ExVid = &v
	return s
}

func (s *QueryExecuteInfoRequest) SetKeyword(v string) *QueryExecuteInfoRequest {
	s.Keyword = &v
	return s
}

func (s *QueryExecuteInfoRequest) SetLevelTree(v *QueryExecuteInfoRequestLevelTree) *QueryExecuteInfoRequest {
	s.LevelTree = v
	return s
}

func (s *QueryExecuteInfoRequest) SetLimit(v int32) *QueryExecuteInfoRequest {
	s.Limit = &v
	return s
}

func (s *QueryExecuteInfoRequest) SetPage(v int32) *QueryExecuteInfoRequest {
	s.Page = &v
	return s
}

func (s *QueryExecuteInfoRequest) SetProduct(v []*string) *QueryExecuteInfoRequest {
	s.Product = v
	return s
}

func (s *QueryExecuteInfoRequest) SetRegion(v []*string) *QueryExecuteInfoRequest {
	s.Region = v
	return s
}

func (s *QueryExecuteInfoRequest) SetReqTimestamp(v int64) *QueryExecuteInfoRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *QueryExecuteInfoRequest) SetSource(v string) *QueryExecuteInfoRequest {
	s.Source = &v
	return s
}

func (s *QueryExecuteInfoRequest) SetSourceOrderId(v string) *QueryExecuteInfoRequest {
	s.SourceOrderId = &v
	return s
}

func (s *QueryExecuteInfoRequest) SetStartTime(v int64) *QueryExecuteInfoRequest {
	s.StartTime = &v
	return s
}

type QueryExecuteInfoRequestLevelTree struct {
	DataType *string                                     `json:"DataType,omitempty" xml:"DataType,omitempty"`
	TreeData []*QueryExecuteInfoRequestLevelTreeTreeData `json:"TreeData,omitempty" xml:"TreeData,omitempty" type:"Repeated"`
}

func (s QueryExecuteInfoRequestLevelTree) String() string {
	return tea.Prettify(s)
}

func (s QueryExecuteInfoRequestLevelTree) GoString() string {
	return s.String()
}

func (s *QueryExecuteInfoRequestLevelTree) SetDataType(v string) *QueryExecuteInfoRequestLevelTree {
	s.DataType = &v
	return s
}

func (s *QueryExecuteInfoRequestLevelTree) SetTreeData(v []*QueryExecuteInfoRequestLevelTreeTreeData) *QueryExecuteInfoRequestLevelTree {
	s.TreeData = v
	return s
}

type QueryExecuteInfoRequestLevelTreeTreeData struct {
	Data        []interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	DataSubType *string       `json:"DataSubType,omitempty" xml:"DataSubType,omitempty"`
	Value       []*string     `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s QueryExecuteInfoRequestLevelTreeTreeData) String() string {
	return tea.Prettify(s)
}

func (s QueryExecuteInfoRequestLevelTreeTreeData) GoString() string {
	return s.String()
}

func (s *QueryExecuteInfoRequestLevelTreeTreeData) SetData(v []interface{}) *QueryExecuteInfoRequestLevelTreeTreeData {
	s.Data = v
	return s
}

func (s *QueryExecuteInfoRequestLevelTreeTreeData) SetDataSubType(v string) *QueryExecuteInfoRequestLevelTreeTreeData {
	s.DataSubType = &v
	return s
}

func (s *QueryExecuteInfoRequestLevelTreeTreeData) SetValue(v []*string) *QueryExecuteInfoRequestLevelTreeTreeData {
	s.Value = v
	return s
}

type QueryExecuteInfoResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryExecuteInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryExecuteInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryExecuteInfoResponseBody) SetCode(v int32) *QueryExecuteInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryExecuteInfoResponseBody) SetData(v string) *QueryExecuteInfoResponseBody {
	s.Data = &v
	return s
}

func (s *QueryExecuteInfoResponseBody) SetMessage(v string) *QueryExecuteInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryExecuteInfoResponseBody) SetRequestId(v string) *QueryExecuteInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryExecuteInfoResponseBody) SetSuccess(v bool) *QueryExecuteInfoResponseBody {
	s.Success = &v
	return s
}

type QueryExecuteInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryExecuteInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryExecuteInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryExecuteInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryExecuteInfoResponse) SetHeaders(v map[string]*string) *QueryExecuteInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryExecuteInfoResponse) SetStatusCode(v int32) *QueryExecuteInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryExecuteInfoResponse) SetBody(v *QueryExecuteInfoResponseBody) *QueryExecuteInfoResponse {
	s.Body = v
	return s
}

type QueryInnerProductInfoRequest struct {
	AuthKey      *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign     *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	Limit        *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Page         *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	ProductCode  *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ReqTimestamp *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
}

func (s QueryInnerProductInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerProductInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryInnerProductInfoRequest) SetAuthKey(v string) *QueryInnerProductInfoRequest {
	s.AuthKey = &v
	return s
}

func (s *QueryInnerProductInfoRequest) SetAuthSign(v string) *QueryInnerProductInfoRequest {
	s.AuthSign = &v
	return s
}

func (s *QueryInnerProductInfoRequest) SetLimit(v int32) *QueryInnerProductInfoRequest {
	s.Limit = &v
	return s
}

func (s *QueryInnerProductInfoRequest) SetPage(v int32) *QueryInnerProductInfoRequest {
	s.Page = &v
	return s
}

func (s *QueryInnerProductInfoRequest) SetProductCode(v string) *QueryInnerProductInfoRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryInnerProductInfoRequest) SetReqTimestamp(v int64) *QueryInnerProductInfoRequest {
	s.ReqTimestamp = &v
	return s
}

type QueryInnerProductInfoResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryInnerProductInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryInnerProductInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerProductInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInnerProductInfoResponseBody) SetCode(v int32) *QueryInnerProductInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInnerProductInfoResponseBody) SetData(v *QueryInnerProductInfoResponseBodyData) *QueryInnerProductInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryInnerProductInfoResponseBody) SetMessage(v string) *QueryInnerProductInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInnerProductInfoResponseBody) SetRequestId(v string) *QueryInnerProductInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInnerProductInfoResponseBody) SetSuccess(v bool) *QueryInnerProductInfoResponseBody {
	s.Success = &v
	return s
}

type QueryInnerProductInfoResponseBodyData struct {
	DataInfo   []*QueryInnerProductInfoResponseBodyDataDataInfo `json:"DataInfo,omitempty" xml:"DataInfo,omitempty" type:"Repeated"`
	Pagination *QueryInnerProductInfoResponseBodyDataPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	Total      *int64                                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryInnerProductInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerProductInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInnerProductInfoResponseBodyData) SetDataInfo(v []*QueryInnerProductInfoResponseBodyDataDataInfo) *QueryInnerProductInfoResponseBodyData {
	s.DataInfo = v
	return s
}

func (s *QueryInnerProductInfoResponseBodyData) SetPagination(v *QueryInnerProductInfoResponseBodyDataPagination) *QueryInnerProductInfoResponseBodyData {
	s.Pagination = v
	return s
}

func (s *QueryInnerProductInfoResponseBodyData) SetTotal(v int64) *QueryInnerProductInfoResponseBodyData {
	s.Total = &v
	return s
}

type QueryInnerProductInfoResponseBodyDataDataInfo struct {
	InnerProductCode *string `json:"InnerProductCode,omitempty" xml:"InnerProductCode,omitempty"`
	InnerProductName *string `json:"InnerProductName,omitempty" xml:"InnerProductName,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName      *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s QueryInnerProductInfoResponseBodyDataDataInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerProductInfoResponseBodyDataDataInfo) GoString() string {
	return s.String()
}

func (s *QueryInnerProductInfoResponseBodyDataDataInfo) SetInnerProductCode(v string) *QueryInnerProductInfoResponseBodyDataDataInfo {
	s.InnerProductCode = &v
	return s
}

func (s *QueryInnerProductInfoResponseBodyDataDataInfo) SetInnerProductName(v string) *QueryInnerProductInfoResponseBodyDataDataInfo {
	s.InnerProductName = &v
	return s
}

func (s *QueryInnerProductInfoResponseBodyDataDataInfo) SetProductCode(v string) *QueryInnerProductInfoResponseBodyDataDataInfo {
	s.ProductCode = &v
	return s
}

func (s *QueryInnerProductInfoResponseBodyDataDataInfo) SetProductName(v string) *QueryInnerProductInfoResponseBodyDataDataInfo {
	s.ProductName = &v
	return s
}

type QueryInnerProductInfoResponseBodyDataPagination struct {
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Page  *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s QueryInnerProductInfoResponseBodyDataPagination) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerProductInfoResponseBodyDataPagination) GoString() string {
	return s.String()
}

func (s *QueryInnerProductInfoResponseBodyDataPagination) SetLimit(v int32) *QueryInnerProductInfoResponseBodyDataPagination {
	s.Limit = &v
	return s
}

func (s *QueryInnerProductInfoResponseBodyDataPagination) SetPage(v int32) *QueryInnerProductInfoResponseBodyDataPagination {
	s.Page = &v
	return s
}

type QueryInnerProductInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInnerProductInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInnerProductInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerProductInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryInnerProductInfoResponse) SetHeaders(v map[string]*string) *QueryInnerProductInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryInnerProductInfoResponse) SetStatusCode(v int32) *QueryInnerProductInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInnerProductInfoResponse) SetBody(v *QueryInnerProductInfoResponseBody) *QueryInnerProductInfoResponse {
	s.Body = v
	return s
}

type QueryRegionAzRequest struct {
	AuthKey      *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign     *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	Limit        *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Page         *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ReqTimestamp *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
}

func (s QueryRegionAzRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRegionAzRequest) GoString() string {
	return s.String()
}

func (s *QueryRegionAzRequest) SetAuthKey(v string) *QueryRegionAzRequest {
	s.AuthKey = &v
	return s
}

func (s *QueryRegionAzRequest) SetAuthSign(v string) *QueryRegionAzRequest {
	s.AuthSign = &v
	return s
}

func (s *QueryRegionAzRequest) SetLimit(v int32) *QueryRegionAzRequest {
	s.Limit = &v
	return s
}

func (s *QueryRegionAzRequest) SetPage(v int32) *QueryRegionAzRequest {
	s.Page = &v
	return s
}

func (s *QueryRegionAzRequest) SetProduct(v string) *QueryRegionAzRequest {
	s.Product = &v
	return s
}

func (s *QueryRegionAzRequest) SetReqTimestamp(v int64) *QueryRegionAzRequest {
	s.ReqTimestamp = &v
	return s
}

type QueryRegionAzResponseBody struct {
	Code      *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryRegionAzResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRegionAzResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRegionAzResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRegionAzResponseBody) SetCode(v int32) *QueryRegionAzResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRegionAzResponseBody) SetData(v *QueryRegionAzResponseBodyData) *QueryRegionAzResponseBody {
	s.Data = v
	return s
}

func (s *QueryRegionAzResponseBody) SetMessage(v string) *QueryRegionAzResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRegionAzResponseBody) SetRequestId(v string) *QueryRegionAzResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRegionAzResponseBody) SetSuccess(v bool) *QueryRegionAzResponseBody {
	s.Success = &v
	return s
}

type QueryRegionAzResponseBodyData struct {
	DataInfo   []*QueryRegionAzResponseBodyDataDataInfo `json:"DataInfo,omitempty" xml:"DataInfo,omitempty" type:"Repeated"`
	Pagination *QueryRegionAzResponseBodyDataPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	Total      *int64                                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryRegionAzResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRegionAzResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRegionAzResponseBodyData) SetDataInfo(v []*QueryRegionAzResponseBodyDataDataInfo) *QueryRegionAzResponseBodyData {
	s.DataInfo = v
	return s
}

func (s *QueryRegionAzResponseBodyData) SetPagination(v *QueryRegionAzResponseBodyDataPagination) *QueryRegionAzResponseBodyData {
	s.Pagination = v
	return s
}

func (s *QueryRegionAzResponseBodyData) SetTotal(v int64) *QueryRegionAzResponseBodyData {
	s.Total = &v
	return s
}

type QueryRegionAzResponseBodyDataDataInfo struct {
	AzList     []*string `json:"AzList,omitempty" xml:"AzList,omitempty" type:"Repeated"`
	RegionCode *string   `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	RegionName *string   `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s QueryRegionAzResponseBodyDataDataInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryRegionAzResponseBodyDataDataInfo) GoString() string {
	return s.String()
}

func (s *QueryRegionAzResponseBodyDataDataInfo) SetAzList(v []*string) *QueryRegionAzResponseBodyDataDataInfo {
	s.AzList = v
	return s
}

func (s *QueryRegionAzResponseBodyDataDataInfo) SetRegionCode(v string) *QueryRegionAzResponseBodyDataDataInfo {
	s.RegionCode = &v
	return s
}

func (s *QueryRegionAzResponseBodyDataDataInfo) SetRegionName(v string) *QueryRegionAzResponseBodyDataDataInfo {
	s.RegionName = &v
	return s
}

type QueryRegionAzResponseBodyDataPagination struct {
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Page  *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s QueryRegionAzResponseBodyDataPagination) String() string {
	return tea.Prettify(s)
}

func (s QueryRegionAzResponseBodyDataPagination) GoString() string {
	return s.String()
}

func (s *QueryRegionAzResponseBodyDataPagination) SetLimit(v int32) *QueryRegionAzResponseBodyDataPagination {
	s.Limit = &v
	return s
}

func (s *QueryRegionAzResponseBodyDataPagination) SetPage(v int32) *QueryRegionAzResponseBodyDataPagination {
	s.Page = &v
	return s
}

type QueryRegionAzResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRegionAzResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRegionAzResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRegionAzResponse) GoString() string {
	return s.String()
}

func (s *QueryRegionAzResponse) SetHeaders(v map[string]*string) *QueryRegionAzResponse {
	s.Headers = v
	return s
}

func (s *QueryRegionAzResponse) SetStatusCode(v int32) *QueryRegionAzResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRegionAzResponse) SetBody(v *QueryRegionAzResponseBody) *QueryRegionAzResponse {
	s.Body = v
	return s
}

type SafeChangeCancelRequest struct {
	AuthKey       *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign      *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	OperateEmpNo  *string `json:"OperateEmpNo,omitempty" xml:"OperateEmpNo,omitempty"`
	ReqTimestamp  *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	SourceOrderId *string `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
}

func (s SafeChangeCancelRequest) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCancelRequest) GoString() string {
	return s.String()
}

func (s *SafeChangeCancelRequest) SetAuthKey(v string) *SafeChangeCancelRequest {
	s.AuthKey = &v
	return s
}

func (s *SafeChangeCancelRequest) SetAuthSign(v string) *SafeChangeCancelRequest {
	s.AuthSign = &v
	return s
}

func (s *SafeChangeCancelRequest) SetOperateEmpNo(v string) *SafeChangeCancelRequest {
	s.OperateEmpNo = &v
	return s
}

func (s *SafeChangeCancelRequest) SetReqTimestamp(v int64) *SafeChangeCancelRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *SafeChangeCancelRequest) SetSourceOrderId(v string) *SafeChangeCancelRequest {
	s.SourceOrderId = &v
	return s
}

type SafeChangeCancelResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SafeChangeCancelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SafeChangeCancelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCancelResponseBody) GoString() string {
	return s.String()
}

func (s *SafeChangeCancelResponseBody) SetCode(v int32) *SafeChangeCancelResponseBody {
	s.Code = &v
	return s
}

func (s *SafeChangeCancelResponseBody) SetData(v *SafeChangeCancelResponseBodyData) *SafeChangeCancelResponseBody {
	s.Data = v
	return s
}

func (s *SafeChangeCancelResponseBody) SetMessage(v string) *SafeChangeCancelResponseBody {
	s.Message = &v
	return s
}

func (s *SafeChangeCancelResponseBody) SetRequestId(v string) *SafeChangeCancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *SafeChangeCancelResponseBody) SetSuccess(v bool) *SafeChangeCancelResponseBody {
	s.Success = &v
	return s
}

type SafeChangeCancelResponseBodyData struct {
	SourceOrderId *string `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
}

func (s SafeChangeCancelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCancelResponseBodyData) GoString() string {
	return s.String()
}

func (s *SafeChangeCancelResponseBodyData) SetSourceOrderId(v string) *SafeChangeCancelResponseBodyData {
	s.SourceOrderId = &v
	return s
}

type SafeChangeCancelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SafeChangeCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SafeChangeCancelResponse) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCancelResponse) GoString() string {
	return s.String()
}

func (s *SafeChangeCancelResponse) SetHeaders(v map[string]*string) *SafeChangeCancelResponse {
	s.Headers = v
	return s
}

func (s *SafeChangeCancelResponse) SetStatusCode(v int32) *SafeChangeCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *SafeChangeCancelResponse) SetBody(v *SafeChangeCancelResponseBody) *SafeChangeCancelResponse {
	s.Body = v
	return s
}

type SafeChangeCheckRequest struct {
	AffectCustomer           *string                                         `json:"AffectCustomer,omitempty" xml:"AffectCustomer,omitempty"`
	ApproveFlowParam         *SafeChangeCheckRequestApproveFlowParam         `json:"ApproveFlowParam,omitempty" xml:"ApproveFlowParam,omitempty" type:"Struct"`
	AuthKey                  *string                                         `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign                 *string                                         `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	BgCustomTemplateExtraDTO *SafeChangeCheckRequestBgCustomTemplateExtraDTO `json:"BgCustomTemplateExtraDTO,omitempty" xml:"BgCustomTemplateExtraDTO,omitempty" type:"Struct"`
	BlockInfos               []*SafeChangeCheckRequestBlockInfos             `json:"BlockInfos,omitempty" xml:"BlockInfos,omitempty" type:"Repeated"`
	CallBackInfo             *SafeChangeCheckRequestCallBackInfo             `json:"CallBackInfo,omitempty" xml:"CallBackInfo,omitempty" type:"Struct"`
	ChangeDataType           *string                                         `json:"ChangeDataType,omitempty" xml:"ChangeDataType,omitempty"`
	ChangeDesc               *string                                         `json:"ChangeDesc,omitempty" xml:"ChangeDesc,omitempty"`
	ChangeEndTime            *int64                                          `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	ChangeEnv                *string                                         `json:"ChangeEnv,omitempty" xml:"ChangeEnv,omitempty"`
	ChangeItems              *string                                         `json:"ChangeItems,omitempty" xml:"ChangeItems,omitempty"`
	ChangeObject             *string                                         `json:"ChangeObject,omitempty" xml:"ChangeObject,omitempty"`
	ChangeOptSubType         *string                                         `json:"ChangeOptSubType,omitempty" xml:"ChangeOptSubType,omitempty"`
	ChangeOptType            *string                                         `json:"ChangeOptType,omitempty" xml:"ChangeOptType,omitempty"`
	ChangeReason             *string                                         `json:"ChangeReason,omitempty" xml:"ChangeReason,omitempty"`
	ChangeRmarks             *string                                         `json:"ChangeRmarks,omitempty" xml:"ChangeRmarks,omitempty"`
	ChangeSchemes            *string                                         `json:"ChangeSchemes,omitempty" xml:"ChangeSchemes,omitempty"`
	ChangeStartTime          *int64                                          `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
	ChangeSubTypeDesc        *string                                         `json:"ChangeSubTypeDesc,omitempty" xml:"ChangeSubTypeDesc,omitempty"`
	ChangeSystem             *string                                         `json:"ChangeSystem,omitempty" xml:"ChangeSystem,omitempty"`
	ChangeTimes              []*SafeChangeCheckRequestChangeTimes            `json:"ChangeTimes,omitempty" xml:"ChangeTimes,omitempty" type:"Repeated"`
	ChangeTitle              *string                                         `json:"ChangeTitle,omitempty" xml:"ChangeTitle,omitempty"`
	ChangeValidation         *string                                         `json:"ChangeValidation,omitempty" xml:"ChangeValidation,omitempty"`
	Checker                  []*string                                       `json:"Checker,omitempty" xml:"Checker,omitempty" type:"Repeated"`
	CreatorEmpId             *string                                         `json:"CreatorEmpId,omitempty" xml:"CreatorEmpId,omitempty"`
	DamagedChangeNotices     []*SafeChangeCheckRequestDamagedChangeNotices   `json:"DamagedChangeNotices,omitempty" xml:"DamagedChangeNotices,omitempty" type:"Repeated"`
	ExecutorEmpId            *string                                         `json:"ExecutorEmpId,omitempty" xml:"ExecutorEmpId,omitempty"`
	ExtraInfo                *string                                         `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Follower                 []*string                                       `json:"Follower,omitempty" xml:"Follower,omitempty" type:"Repeated"`
	GrayStatus               *string                                         `json:"GrayStatus,omitempty" xml:"GrayStatus,omitempty"`
	HarmChangeNoticeEnum     *string                                         `json:"HarmChangeNoticeEnum,omitempty" xml:"HarmChangeNoticeEnum,omitempty"`
	Incidence                *string                                         `json:"Incidence,omitempty" xml:"Incidence,omitempty"`
	InfluenceInfo            *SafeChangeCheckRequestInfluenceInfo            `json:"InfluenceInfo,omitempty" xml:"InfluenceInfo,omitempty" type:"Struct"`
	Instance                 *SafeChangeCheckRequestInstance                 `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	NeedModifyDoc            *string                                         `json:"NeedModifyDoc,omitempty" xml:"NeedModifyDoc,omitempty"`
	OperateEmpNo             *string                                         `json:"OperateEmpNo,omitempty" xml:"OperateEmpNo,omitempty"`
	Product                  []*SafeChangeCheckRequestProduct                `json:"Product,omitempty" xml:"Product,omitempty" type:"Repeated"`
	ReleasePackageInfos      []*SafeChangeCheckRequestReleasePackageInfos    `json:"ReleasePackageInfos,omitempty" xml:"ReleasePackageInfos,omitempty" type:"Repeated"`
	ReqTimestamp             *int64                                          `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	ReuseSourceOrderId       *string                                         `json:"ReuseSourceOrderId,omitempty" xml:"ReuseSourceOrderId,omitempty"`
	RiskLevel                *string                                         `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Rollback                 *string                                         `json:"Rollback,omitempty" xml:"Rollback,omitempty"`
	SourceName               *string                                         `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	SourceOrderId            *string                                         `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
	SourceUrl                *string                                         `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	WhiteType                *int32                                          `json:"whiteType,omitempty" xml:"whiteType,omitempty"`
}

func (s SafeChangeCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequest) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequest) SetAffectCustomer(v string) *SafeChangeCheckRequest {
	s.AffectCustomer = &v
	return s
}

func (s *SafeChangeCheckRequest) SetApproveFlowParam(v *SafeChangeCheckRequestApproveFlowParam) *SafeChangeCheckRequest {
	s.ApproveFlowParam = v
	return s
}

func (s *SafeChangeCheckRequest) SetAuthKey(v string) *SafeChangeCheckRequest {
	s.AuthKey = &v
	return s
}

func (s *SafeChangeCheckRequest) SetAuthSign(v string) *SafeChangeCheckRequest {
	s.AuthSign = &v
	return s
}

func (s *SafeChangeCheckRequest) SetBgCustomTemplateExtraDTO(v *SafeChangeCheckRequestBgCustomTemplateExtraDTO) *SafeChangeCheckRequest {
	s.BgCustomTemplateExtraDTO = v
	return s
}

func (s *SafeChangeCheckRequest) SetBlockInfos(v []*SafeChangeCheckRequestBlockInfos) *SafeChangeCheckRequest {
	s.BlockInfos = v
	return s
}

func (s *SafeChangeCheckRequest) SetCallBackInfo(v *SafeChangeCheckRequestCallBackInfo) *SafeChangeCheckRequest {
	s.CallBackInfo = v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeDataType(v string) *SafeChangeCheckRequest {
	s.ChangeDataType = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeDesc(v string) *SafeChangeCheckRequest {
	s.ChangeDesc = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeEndTime(v int64) *SafeChangeCheckRequest {
	s.ChangeEndTime = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeEnv(v string) *SafeChangeCheckRequest {
	s.ChangeEnv = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeItems(v string) *SafeChangeCheckRequest {
	s.ChangeItems = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeObject(v string) *SafeChangeCheckRequest {
	s.ChangeObject = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeOptSubType(v string) *SafeChangeCheckRequest {
	s.ChangeOptSubType = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeOptType(v string) *SafeChangeCheckRequest {
	s.ChangeOptType = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeReason(v string) *SafeChangeCheckRequest {
	s.ChangeReason = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeRmarks(v string) *SafeChangeCheckRequest {
	s.ChangeRmarks = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeSchemes(v string) *SafeChangeCheckRequest {
	s.ChangeSchemes = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeStartTime(v int64) *SafeChangeCheckRequest {
	s.ChangeStartTime = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeSubTypeDesc(v string) *SafeChangeCheckRequest {
	s.ChangeSubTypeDesc = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeSystem(v string) *SafeChangeCheckRequest {
	s.ChangeSystem = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeTimes(v []*SafeChangeCheckRequestChangeTimes) *SafeChangeCheckRequest {
	s.ChangeTimes = v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeTitle(v string) *SafeChangeCheckRequest {
	s.ChangeTitle = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChangeValidation(v string) *SafeChangeCheckRequest {
	s.ChangeValidation = &v
	return s
}

func (s *SafeChangeCheckRequest) SetChecker(v []*string) *SafeChangeCheckRequest {
	s.Checker = v
	return s
}

func (s *SafeChangeCheckRequest) SetCreatorEmpId(v string) *SafeChangeCheckRequest {
	s.CreatorEmpId = &v
	return s
}

func (s *SafeChangeCheckRequest) SetDamagedChangeNotices(v []*SafeChangeCheckRequestDamagedChangeNotices) *SafeChangeCheckRequest {
	s.DamagedChangeNotices = v
	return s
}

func (s *SafeChangeCheckRequest) SetExecutorEmpId(v string) *SafeChangeCheckRequest {
	s.ExecutorEmpId = &v
	return s
}

func (s *SafeChangeCheckRequest) SetExtraInfo(v string) *SafeChangeCheckRequest {
	s.ExtraInfo = &v
	return s
}

func (s *SafeChangeCheckRequest) SetFollower(v []*string) *SafeChangeCheckRequest {
	s.Follower = v
	return s
}

func (s *SafeChangeCheckRequest) SetGrayStatus(v string) *SafeChangeCheckRequest {
	s.GrayStatus = &v
	return s
}

func (s *SafeChangeCheckRequest) SetHarmChangeNoticeEnum(v string) *SafeChangeCheckRequest {
	s.HarmChangeNoticeEnum = &v
	return s
}

func (s *SafeChangeCheckRequest) SetIncidence(v string) *SafeChangeCheckRequest {
	s.Incidence = &v
	return s
}

func (s *SafeChangeCheckRequest) SetInfluenceInfo(v *SafeChangeCheckRequestInfluenceInfo) *SafeChangeCheckRequest {
	s.InfluenceInfo = v
	return s
}

func (s *SafeChangeCheckRequest) SetInstance(v *SafeChangeCheckRequestInstance) *SafeChangeCheckRequest {
	s.Instance = v
	return s
}

func (s *SafeChangeCheckRequest) SetNeedModifyDoc(v string) *SafeChangeCheckRequest {
	s.NeedModifyDoc = &v
	return s
}

func (s *SafeChangeCheckRequest) SetOperateEmpNo(v string) *SafeChangeCheckRequest {
	s.OperateEmpNo = &v
	return s
}

func (s *SafeChangeCheckRequest) SetProduct(v []*SafeChangeCheckRequestProduct) *SafeChangeCheckRequest {
	s.Product = v
	return s
}

func (s *SafeChangeCheckRequest) SetReleasePackageInfos(v []*SafeChangeCheckRequestReleasePackageInfos) *SafeChangeCheckRequest {
	s.ReleasePackageInfos = v
	return s
}

func (s *SafeChangeCheckRequest) SetReqTimestamp(v int64) *SafeChangeCheckRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *SafeChangeCheckRequest) SetReuseSourceOrderId(v string) *SafeChangeCheckRequest {
	s.ReuseSourceOrderId = &v
	return s
}

func (s *SafeChangeCheckRequest) SetRiskLevel(v string) *SafeChangeCheckRequest {
	s.RiskLevel = &v
	return s
}

func (s *SafeChangeCheckRequest) SetRollback(v string) *SafeChangeCheckRequest {
	s.Rollback = &v
	return s
}

func (s *SafeChangeCheckRequest) SetSourceName(v string) *SafeChangeCheckRequest {
	s.SourceName = &v
	return s
}

func (s *SafeChangeCheckRequest) SetSourceOrderId(v string) *SafeChangeCheckRequest {
	s.SourceOrderId = &v
	return s
}

func (s *SafeChangeCheckRequest) SetSourceUrl(v string) *SafeChangeCheckRequest {
	s.SourceUrl = &v
	return s
}

func (s *SafeChangeCheckRequest) SetWhiteType(v int32) *SafeChangeCheckRequest {
	s.WhiteType = &v
	return s
}

type SafeChangeCheckRequestApproveFlowParam struct {
	ApproveNodes []*SafeChangeCheckRequestApproveFlowParamApproveNodes `json:"ApproveNodes,omitempty" xml:"ApproveNodes,omitempty" type:"Repeated"`
	FlowStatus   *int32                                                `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty"`
}

func (s SafeChangeCheckRequestApproveFlowParam) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestApproveFlowParam) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestApproveFlowParam) SetApproveNodes(v []*SafeChangeCheckRequestApproveFlowParamApproveNodes) *SafeChangeCheckRequestApproveFlowParam {
	s.ApproveNodes = v
	return s
}

func (s *SafeChangeCheckRequestApproveFlowParam) SetFlowStatus(v int32) *SafeChangeCheckRequestApproveFlowParam {
	s.FlowStatus = &v
	return s
}

type SafeChangeCheckRequestApproveFlowParamApproveNodes struct {
	ApproverDTO      []*SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO `json:"ApproverDTO,omitempty" xml:"ApproverDTO,omitempty" type:"Repeated"`
	NodeStatus       *int32                                                           `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	ProcessName      *string                                                          `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	ProcessNodeOrder *int32                                                           `json:"ProcessNodeOrder,omitempty" xml:"ProcessNodeOrder,omitempty"`
	Strategy         *int32                                                           `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s SafeChangeCheckRequestApproveFlowParamApproveNodes) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestApproveFlowParamApproveNodes) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestApproveFlowParamApproveNodes) SetApproverDTO(v []*SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) *SafeChangeCheckRequestApproveFlowParamApproveNodes {
	s.ApproverDTO = v
	return s
}

func (s *SafeChangeCheckRequestApproveFlowParamApproveNodes) SetNodeStatus(v int32) *SafeChangeCheckRequestApproveFlowParamApproveNodes {
	s.NodeStatus = &v
	return s
}

func (s *SafeChangeCheckRequestApproveFlowParamApproveNodes) SetProcessName(v string) *SafeChangeCheckRequestApproveFlowParamApproveNodes {
	s.ProcessName = &v
	return s
}

func (s *SafeChangeCheckRequestApproveFlowParamApproveNodes) SetProcessNodeOrder(v int32) *SafeChangeCheckRequestApproveFlowParamApproveNodes {
	s.ProcessNodeOrder = &v
	return s
}

func (s *SafeChangeCheckRequestApproveFlowParamApproveNodes) SetStrategy(v int32) *SafeChangeCheckRequestApproveFlowParamApproveNodes {
	s.Strategy = &v
	return s
}

type SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO struct {
	ApproveDesc  *string `json:"ApproveDesc,omitempty" xml:"ApproveDesc,omitempty"`
	ApproveTime  *int64  `json:"ApproveTime,omitempty" xml:"ApproveTime,omitempty"`
	ApproverId   *string `json:"ApproverId,omitempty" xml:"ApproverId,omitempty"`
	ApproverName *string `json:"ApproverName,omitempty" xml:"ApproverName,omitempty"`
	Opinion      *int32  `json:"Opinion,omitempty" xml:"Opinion,omitempty"`
}

func (s SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) SetApproveDesc(v string) *SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO {
	s.ApproveDesc = &v
	return s
}

func (s *SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) SetApproveTime(v int64) *SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO {
	s.ApproveTime = &v
	return s
}

func (s *SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) SetApproverId(v string) *SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO {
	s.ApproverId = &v
	return s
}

func (s *SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) SetApproverName(v string) *SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO {
	s.ApproverName = &v
	return s
}

func (s *SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO) SetOpinion(v int32) *SafeChangeCheckRequestApproveFlowParamApproveNodesApproverDTO {
	s.Opinion = &v
	return s
}

type SafeChangeCheckRequestBgCustomTemplateExtraDTO struct {
	BgCustomTemplateInfo *string `json:"BgCustomTemplateInfo,omitempty" xml:"BgCustomTemplateInfo,omitempty"`
}

func (s SafeChangeCheckRequestBgCustomTemplateExtraDTO) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestBgCustomTemplateExtraDTO) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestBgCustomTemplateExtraDTO) SetBgCustomTemplateInfo(v string) *SafeChangeCheckRequestBgCustomTemplateExtraDTO {
	s.BgCustomTemplateInfo = &v
	return s
}

type SafeChangeCheckRequestBlockInfos struct {
	HitInfos []*SafeChangeCheckRequestBlockInfosHitInfos `json:"HitInfos,omitempty" xml:"HitInfos,omitempty" type:"Repeated"`
	Id       *int64                                      `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SafeChangeCheckRequestBlockInfos) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestBlockInfos) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestBlockInfos) SetHitInfos(v []*SafeChangeCheckRequestBlockInfosHitInfos) *SafeChangeCheckRequestBlockInfos {
	s.HitInfos = v
	return s
}

func (s *SafeChangeCheckRequestBlockInfos) SetId(v int64) *SafeChangeCheckRequestBlockInfos {
	s.Id = &v
	return s
}

type SafeChangeCheckRequestBlockInfosHitInfos struct {
	HitInfo   *string `json:"HitInfo,omitempty" xml:"HitInfo,omitempty"`
	HitObject *string `json:"HitObject,omitempty" xml:"HitObject,omitempty"`
	Scope     *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s SafeChangeCheckRequestBlockInfosHitInfos) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestBlockInfosHitInfos) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestBlockInfosHitInfos) SetHitInfo(v string) *SafeChangeCheckRequestBlockInfosHitInfos {
	s.HitInfo = &v
	return s
}

func (s *SafeChangeCheckRequestBlockInfosHitInfos) SetHitObject(v string) *SafeChangeCheckRequestBlockInfosHitInfos {
	s.HitObject = &v
	return s
}

func (s *SafeChangeCheckRequestBlockInfosHitInfos) SetScope(v string) *SafeChangeCheckRequestBlockInfosHitInfos {
	s.Scope = &v
	return s
}

type SafeChangeCheckRequestCallBackInfo struct {
	Api        *string `json:"Api,omitempty" xml:"Api,omitempty"`
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	EndPoint   *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	PopProduct *string `json:"PopProduct,omitempty" xml:"PopProduct,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url        *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SafeChangeCheckRequestCallBackInfo) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestCallBackInfo) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestCallBackInfo) SetApi(v string) *SafeChangeCheckRequestCallBackInfo {
	s.Api = &v
	return s
}

func (s *SafeChangeCheckRequestCallBackInfo) SetApiVersion(v string) *SafeChangeCheckRequestCallBackInfo {
	s.ApiVersion = &v
	return s
}

func (s *SafeChangeCheckRequestCallBackInfo) SetEndPoint(v string) *SafeChangeCheckRequestCallBackInfo {
	s.EndPoint = &v
	return s
}

func (s *SafeChangeCheckRequestCallBackInfo) SetPopProduct(v string) *SafeChangeCheckRequestCallBackInfo {
	s.PopProduct = &v
	return s
}

func (s *SafeChangeCheckRequestCallBackInfo) SetRegionId(v string) *SafeChangeCheckRequestCallBackInfo {
	s.RegionId = &v
	return s
}

func (s *SafeChangeCheckRequestCallBackInfo) SetType(v string) *SafeChangeCheckRequestCallBackInfo {
	s.Type = &v
	return s
}

func (s *SafeChangeCheckRequestCallBackInfo) SetUrl(v string) *SafeChangeCheckRequestCallBackInfo {
	s.Url = &v
	return s
}

type SafeChangeCheckRequestChangeTimes struct {
	ChangeEndTime   *int64 `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	ChangeStartTime *int64 `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
}

func (s SafeChangeCheckRequestChangeTimes) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestChangeTimes) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestChangeTimes) SetChangeEndTime(v int64) *SafeChangeCheckRequestChangeTimes {
	s.ChangeEndTime = &v
	return s
}

func (s *SafeChangeCheckRequestChangeTimes) SetChangeStartTime(v int64) *SafeChangeCheckRequestChangeTimes {
	s.ChangeStartTime = &v
	return s
}

type SafeChangeCheckRequestDamagedChangeNotices struct {
	BgCancelNoticeContent *string                                                         `json:"BgCancelNoticeContent,omitempty" xml:"BgCancelNoticeContent,omitempty"`
	BgCancelNoticeEventId *string                                                         `json:"BgCancelNoticeEventId,omitempty" xml:"BgCancelNoticeEventId,omitempty"`
	Channel               []*string                                                       `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Repeated"`
	Content               *string                                                         `json:"Content,omitempty" xml:"Content,omitempty"`
	EventId               *string                                                         `json:"EventId,omitempty" xml:"EventId,omitempty"`
	SensitiveCustomers    []*SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomers `json:"SensitiveCustomers,omitempty" xml:"SensitiveCustomers,omitempty" type:"Repeated"`
	Type                  *string                                                         `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SafeChangeCheckRequestDamagedChangeNotices) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestDamagedChangeNotices) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestDamagedChangeNotices) SetBgCancelNoticeContent(v string) *SafeChangeCheckRequestDamagedChangeNotices {
	s.BgCancelNoticeContent = &v
	return s
}

func (s *SafeChangeCheckRequestDamagedChangeNotices) SetBgCancelNoticeEventId(v string) *SafeChangeCheckRequestDamagedChangeNotices {
	s.BgCancelNoticeEventId = &v
	return s
}

func (s *SafeChangeCheckRequestDamagedChangeNotices) SetChannel(v []*string) *SafeChangeCheckRequestDamagedChangeNotices {
	s.Channel = v
	return s
}

func (s *SafeChangeCheckRequestDamagedChangeNotices) SetContent(v string) *SafeChangeCheckRequestDamagedChangeNotices {
	s.Content = &v
	return s
}

func (s *SafeChangeCheckRequestDamagedChangeNotices) SetEventId(v string) *SafeChangeCheckRequestDamagedChangeNotices {
	s.EventId = &v
	return s
}

func (s *SafeChangeCheckRequestDamagedChangeNotices) SetSensitiveCustomers(v []*SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomers) *SafeChangeCheckRequestDamagedChangeNotices {
	s.SensitiveCustomers = v
	return s
}

func (s *SafeChangeCheckRequestDamagedChangeNotices) SetType(v string) *SafeChangeCheckRequestDamagedChangeNotices {
	s.Type = &v
	return s
}

type SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomers struct {
	CustomerInfo []*SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo `json:"CustomerInfo,omitempty" xml:"CustomerInfo,omitempty" type:"Repeated"`
	ProductCode  *string                                                                     `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomers) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomers) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomers) SetCustomerInfo(v []*SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo) *SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomers {
	s.CustomerInfo = v
	return s
}

func (s *SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomers) SetProductCode(v string) *SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomers {
	s.ProductCode = &v
	return s
}

type SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo struct {
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Type      *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid       *string                `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo) SetExtraInfo(v map[string]interface{}) *SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo {
	s.ExtraInfo = v
	return s
}

func (s *SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo) SetType(v string) *SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo {
	s.Type = &v
	return s
}

func (s *SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo) SetUid(v string) *SafeChangeCheckRequestDamagedChangeNoticesSensitiveCustomersCustomerInfo {
	s.Uid = &v
	return s
}

type SafeChangeCheckRequestInfluenceInfo struct {
	NoticeInfos        []*SafeChangeCheckRequestInfluenceInfoNoticeInfos        `json:"NoticeInfos,omitempty" xml:"NoticeInfos,omitempty" type:"Repeated"`
	SensitiveCustomers []*SafeChangeCheckRequestInfluenceInfoSensitiveCustomers `json:"SensitiveCustomers,omitempty" xml:"SensitiveCustomers,omitempty" type:"Repeated"`
}

func (s SafeChangeCheckRequestInfluenceInfo) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestInfluenceInfo) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestInfluenceInfo) SetNoticeInfos(v []*SafeChangeCheckRequestInfluenceInfoNoticeInfos) *SafeChangeCheckRequestInfluenceInfo {
	s.NoticeInfos = v
	return s
}

func (s *SafeChangeCheckRequestInfluenceInfo) SetSensitiveCustomers(v []*SafeChangeCheckRequestInfluenceInfoSensitiveCustomers) *SafeChangeCheckRequestInfluenceInfo {
	s.SensitiveCustomers = v
	return s
}

type SafeChangeCheckRequestInfluenceInfoNoticeInfos struct {
	Channel []*string `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Repeated"`
	Content *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	EventId *string   `json:"EventId,omitempty" xml:"EventId,omitempty"`
}

func (s SafeChangeCheckRequestInfluenceInfoNoticeInfos) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestInfluenceInfoNoticeInfos) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestInfluenceInfoNoticeInfos) SetChannel(v []*string) *SafeChangeCheckRequestInfluenceInfoNoticeInfos {
	s.Channel = v
	return s
}

func (s *SafeChangeCheckRequestInfluenceInfoNoticeInfos) SetContent(v string) *SafeChangeCheckRequestInfluenceInfoNoticeInfos {
	s.Content = &v
	return s
}

func (s *SafeChangeCheckRequestInfluenceInfoNoticeInfos) SetEventId(v string) *SafeChangeCheckRequestInfluenceInfoNoticeInfos {
	s.EventId = &v
	return s
}

type SafeChangeCheckRequestInfluenceInfoSensitiveCustomers struct {
	CustomerInfo []*SafeChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo `json:"CustomerInfo,omitempty" xml:"CustomerInfo,omitempty" type:"Repeated"`
	ProductCode  *string                                                              `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s SafeChangeCheckRequestInfluenceInfoSensitiveCustomers) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestInfluenceInfoSensitiveCustomers) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestInfluenceInfoSensitiveCustomers) SetCustomerInfo(v []*SafeChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo) *SafeChangeCheckRequestInfluenceInfoSensitiveCustomers {
	s.CustomerInfo = v
	return s
}

func (s *SafeChangeCheckRequestInfluenceInfoSensitiveCustomers) SetProductCode(v string) *SafeChangeCheckRequestInfluenceInfoSensitiveCustomers {
	s.ProductCode = &v
	return s
}

type SafeChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo struct {
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Type      *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid       *string                `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s SafeChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo) SetExtraInfo(v map[string]interface{}) *SafeChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo {
	s.ExtraInfo = v
	return s
}

func (s *SafeChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo) SetType(v string) *SafeChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo {
	s.Type = &v
	return s
}

func (s *SafeChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo) SetUid(v string) *SafeChangeCheckRequestInfluenceInfoSensitiveCustomersCustomerInfo {
	s.Uid = &v
	return s
}

type SafeChangeCheckRequestInstance struct {
	Nc             []*string `json:"Nc,omitempty" xml:"Nc,omitempty" type:"Repeated"`
	Uids           []*string `json:"Uids,omitempty" xml:"Uids,omitempty" type:"Repeated"`
	AttributionApp []*string `json:"attributionApp,omitempty" xml:"attributionApp,omitempty" type:"Repeated"`
	InfluenceApp   []*string `json:"influenceApp,omitempty" xml:"influenceApp,omitempty" type:"Repeated"`
	Instance       []*string `json:"instance,omitempty" xml:"instance,omitempty" type:"Repeated"`
}

func (s SafeChangeCheckRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestInstance) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestInstance) SetNc(v []*string) *SafeChangeCheckRequestInstance {
	s.Nc = v
	return s
}

func (s *SafeChangeCheckRequestInstance) SetUids(v []*string) *SafeChangeCheckRequestInstance {
	s.Uids = v
	return s
}

func (s *SafeChangeCheckRequestInstance) SetAttributionApp(v []*string) *SafeChangeCheckRequestInstance {
	s.AttributionApp = v
	return s
}

func (s *SafeChangeCheckRequestInstance) SetInfluenceApp(v []*string) *SafeChangeCheckRequestInstance {
	s.InfluenceApp = v
	return s
}

func (s *SafeChangeCheckRequestInstance) SetInstance(v []*string) *SafeChangeCheckRequestInstance {
	s.Instance = v
	return s
}

type SafeChangeCheckRequestProduct struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SafeChangeCheckRequestProduct) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestProduct) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestProduct) SetCode(v string) *SafeChangeCheckRequestProduct {
	s.Code = &v
	return s
}

func (s *SafeChangeCheckRequestProduct) SetName(v string) *SafeChangeCheckRequestProduct {
	s.Name = &v
	return s
}

type SafeChangeCheckRequestReleasePackageInfos struct {
	ProductCode    *string   `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ReleasePackage []*string `json:"ReleasePackage,omitempty" xml:"ReleasePackage,omitempty" type:"Repeated"`
}

func (s SafeChangeCheckRequestReleasePackageInfos) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckRequestReleasePackageInfos) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckRequestReleasePackageInfos) SetProductCode(v string) *SafeChangeCheckRequestReleasePackageInfos {
	s.ProductCode = &v
	return s
}

func (s *SafeChangeCheckRequestReleasePackageInfos) SetReleasePackage(v []*string) *SafeChangeCheckRequestReleasePackageInfos {
	s.ReleasePackage = v
	return s
}

type SafeChangeCheckResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SafeChangeCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SafeChangeCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckResponseBody) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckResponseBody) SetCode(v int32) *SafeChangeCheckResponseBody {
	s.Code = &v
	return s
}

func (s *SafeChangeCheckResponseBody) SetData(v *SafeChangeCheckResponseBodyData) *SafeChangeCheckResponseBody {
	s.Data = v
	return s
}

func (s *SafeChangeCheckResponseBody) SetMessage(v string) *SafeChangeCheckResponseBody {
	s.Message = &v
	return s
}

func (s *SafeChangeCheckResponseBody) SetRequestId(v string) *SafeChangeCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *SafeChangeCheckResponseBody) SetSuccess(v bool) *SafeChangeCheckResponseBody {
	s.Success = &v
	return s
}

type SafeChangeCheckResponseBodyData struct {
	ApproveResultUrl  *string                                             `json:"ApproveResultUrl,omitempty" xml:"ApproveResultUrl,omitempty"`
	BgCheckStatus     *string                                             `json:"BgCheckStatus,omitempty" xml:"BgCheckStatus,omitempty"`
	BgVid             *string                                             `json:"BgVid,omitempty" xml:"BgVid,omitempty"`
	CheckResultUrl    *string                                             `json:"CheckResultUrl,omitempty" xml:"CheckResultUrl,omitempty"`
	CheckStatus       *string                                             `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckholdReason   []*string                                           `json:"CheckholdReason,omitempty" xml:"CheckholdReason,omitempty" type:"Repeated"`
	RuleDetailUrlList []*SafeChangeCheckResponseBodyDataRuleDetailUrlList `json:"RuleDetailUrlList,omitempty" xml:"RuleDetailUrlList,omitempty" type:"Repeated"`
	SourceOrderId     *string                                             `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
}

func (s SafeChangeCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckResponseBodyData) SetApproveResultUrl(v string) *SafeChangeCheckResponseBodyData {
	s.ApproveResultUrl = &v
	return s
}

func (s *SafeChangeCheckResponseBodyData) SetBgCheckStatus(v string) *SafeChangeCheckResponseBodyData {
	s.BgCheckStatus = &v
	return s
}

func (s *SafeChangeCheckResponseBodyData) SetBgVid(v string) *SafeChangeCheckResponseBodyData {
	s.BgVid = &v
	return s
}

func (s *SafeChangeCheckResponseBodyData) SetCheckResultUrl(v string) *SafeChangeCheckResponseBodyData {
	s.CheckResultUrl = &v
	return s
}

func (s *SafeChangeCheckResponseBodyData) SetCheckStatus(v string) *SafeChangeCheckResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *SafeChangeCheckResponseBodyData) SetCheckholdReason(v []*string) *SafeChangeCheckResponseBodyData {
	s.CheckholdReason = v
	return s
}

func (s *SafeChangeCheckResponseBodyData) SetRuleDetailUrlList(v []*SafeChangeCheckResponseBodyDataRuleDetailUrlList) *SafeChangeCheckResponseBodyData {
	s.RuleDetailUrlList = v
	return s
}

func (s *SafeChangeCheckResponseBodyData) SetSourceOrderId(v string) *SafeChangeCheckResponseBodyData {
	s.SourceOrderId = &v
	return s
}

type SafeChangeCheckResponseBodyDataRuleDetailUrlList struct {
	SceneEnum *string `json:"SceneEnum,omitempty" xml:"SceneEnum,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SafeChangeCheckResponseBodyDataRuleDetailUrlList) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckResponseBodyDataRuleDetailUrlList) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckResponseBodyDataRuleDetailUrlList) SetSceneEnum(v string) *SafeChangeCheckResponseBodyDataRuleDetailUrlList {
	s.SceneEnum = &v
	return s
}

func (s *SafeChangeCheckResponseBodyDataRuleDetailUrlList) SetTitle(v string) *SafeChangeCheckResponseBodyDataRuleDetailUrlList {
	s.Title = &v
	return s
}

func (s *SafeChangeCheckResponseBodyDataRuleDetailUrlList) SetUrl(v string) *SafeChangeCheckResponseBodyDataRuleDetailUrlList {
	s.Url = &v
	return s
}

type SafeChangeCheckResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SafeChangeCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SafeChangeCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeCheckResponse) GoString() string {
	return s.String()
}

func (s *SafeChangeCheckResponse) SetHeaders(v map[string]*string) *SafeChangeCheckResponse {
	s.Headers = v
	return s
}

func (s *SafeChangeCheckResponse) SetStatusCode(v int32) *SafeChangeCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *SafeChangeCheckResponse) SetBody(v *SafeChangeCheckResponseBody) *SafeChangeCheckResponse {
	s.Body = v
	return s
}

type SafeChangeEndRequest struct {
	AuthKey       *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign      *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	ChangeEndTime *int64  `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	ChangeResult  *string `json:"ChangeResult,omitempty" xml:"ChangeResult,omitempty"`
	CurBatchNo    *int32  `json:"CurBatchNo,omitempty" xml:"CurBatchNo,omitempty"`
	ExecutorEmpId *string `json:"ExecutorEmpId,omitempty" xml:"ExecutorEmpId,omitempty"`
	ReqTimestamp  *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	SourceOrderId *string `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
	TotalBatchNo  *int32  `json:"TotalBatchNo,omitempty" xml:"TotalBatchNo,omitempty"`
}

func (s SafeChangeEndRequest) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeEndRequest) GoString() string {
	return s.String()
}

func (s *SafeChangeEndRequest) SetAuthKey(v string) *SafeChangeEndRequest {
	s.AuthKey = &v
	return s
}

func (s *SafeChangeEndRequest) SetAuthSign(v string) *SafeChangeEndRequest {
	s.AuthSign = &v
	return s
}

func (s *SafeChangeEndRequest) SetChangeEndTime(v int64) *SafeChangeEndRequest {
	s.ChangeEndTime = &v
	return s
}

func (s *SafeChangeEndRequest) SetChangeResult(v string) *SafeChangeEndRequest {
	s.ChangeResult = &v
	return s
}

func (s *SafeChangeEndRequest) SetCurBatchNo(v int32) *SafeChangeEndRequest {
	s.CurBatchNo = &v
	return s
}

func (s *SafeChangeEndRequest) SetExecutorEmpId(v string) *SafeChangeEndRequest {
	s.ExecutorEmpId = &v
	return s
}

func (s *SafeChangeEndRequest) SetReqTimestamp(v int64) *SafeChangeEndRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *SafeChangeEndRequest) SetSourceOrderId(v string) *SafeChangeEndRequest {
	s.SourceOrderId = &v
	return s
}

func (s *SafeChangeEndRequest) SetTotalBatchNo(v int32) *SafeChangeEndRequest {
	s.TotalBatchNo = &v
	return s
}

type SafeChangeEndResponseBody struct {
	Code      *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SafeChangeEndResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SafeChangeEndResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeEndResponseBody) GoString() string {
	return s.String()
}

func (s *SafeChangeEndResponseBody) SetCode(v int32) *SafeChangeEndResponseBody {
	s.Code = &v
	return s
}

func (s *SafeChangeEndResponseBody) SetData(v *SafeChangeEndResponseBodyData) *SafeChangeEndResponseBody {
	s.Data = v
	return s
}

func (s *SafeChangeEndResponseBody) SetMessage(v string) *SafeChangeEndResponseBody {
	s.Message = &v
	return s
}

func (s *SafeChangeEndResponseBody) SetRequestId(v string) *SafeChangeEndResponseBody {
	s.RequestId = &v
	return s
}

func (s *SafeChangeEndResponseBody) SetSuccess(v bool) *SafeChangeEndResponseBody {
	s.Success = &v
	return s
}

type SafeChangeEndResponseBodyData struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SafeChangeEndResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeEndResponseBodyData) GoString() string {
	return s.String()
}

func (s *SafeChangeEndResponseBodyData) SetStatus(v string) *SafeChangeEndResponseBodyData {
	s.Status = &v
	return s
}

type SafeChangeEndResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SafeChangeEndResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SafeChangeEndResponse) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeEndResponse) GoString() string {
	return s.String()
}

func (s *SafeChangeEndResponse) SetHeaders(v map[string]*string) *SafeChangeEndResponse {
	s.Headers = v
	return s
}

func (s *SafeChangeEndResponse) SetStatusCode(v int32) *SafeChangeEndResponse {
	s.StatusCode = &v
	return s
}

func (s *SafeChangeEndResponse) SetBody(v *SafeChangeEndResponseBody) *SafeChangeEndResponse {
	s.Body = v
	return s
}

type SafeChangeQueryRequest struct {
	AuthKey       *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign      *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	NeedValidate  *bool   `json:"NeedValidate,omitempty" xml:"NeedValidate,omitempty"`
	QueryType     *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ReqTimestamp  *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	ReturnType    *bool   `json:"ReturnType,omitempty" xml:"ReturnType,omitempty"`
	SourceOrderId *string `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
}

func (s SafeChangeQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeQueryRequest) GoString() string {
	return s.String()
}

func (s *SafeChangeQueryRequest) SetAuthKey(v string) *SafeChangeQueryRequest {
	s.AuthKey = &v
	return s
}

func (s *SafeChangeQueryRequest) SetAuthSign(v string) *SafeChangeQueryRequest {
	s.AuthSign = &v
	return s
}

func (s *SafeChangeQueryRequest) SetNeedValidate(v bool) *SafeChangeQueryRequest {
	s.NeedValidate = &v
	return s
}

func (s *SafeChangeQueryRequest) SetQueryType(v string) *SafeChangeQueryRequest {
	s.QueryType = &v
	return s
}

func (s *SafeChangeQueryRequest) SetReqTimestamp(v int64) *SafeChangeQueryRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *SafeChangeQueryRequest) SetReturnType(v bool) *SafeChangeQueryRequest {
	s.ReturnType = &v
	return s
}

func (s *SafeChangeQueryRequest) SetSourceOrderId(v string) *SafeChangeQueryRequest {
	s.SourceOrderId = &v
	return s
}

type SafeChangeQueryResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SafeChangeQueryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SafeChangeQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeQueryResponseBody) GoString() string {
	return s.String()
}

func (s *SafeChangeQueryResponseBody) SetCode(v int32) *SafeChangeQueryResponseBody {
	s.Code = &v
	return s
}

func (s *SafeChangeQueryResponseBody) SetData(v *SafeChangeQueryResponseBodyData) *SafeChangeQueryResponseBody {
	s.Data = v
	return s
}

func (s *SafeChangeQueryResponseBody) SetMessage(v string) *SafeChangeQueryResponseBody {
	s.Message = &v
	return s
}

func (s *SafeChangeQueryResponseBody) SetRequestId(v string) *SafeChangeQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SafeChangeQueryResponseBody) SetSuccess(v bool) *SafeChangeQueryResponseBody {
	s.Success = &v
	return s
}

type SafeChangeQueryResponseBodyData struct {
	ApproveResultUrl *string                                       `json:"ApproveResultUrl,omitempty" xml:"ApproveResultUrl,omitempty"`
	ApproveStatus    *string                                       `json:"ApproveStatus,omitempty" xml:"ApproveStatus,omitempty"`
	ChangeCancel     *string                                       `json:"ChangeCancel,omitempty" xml:"ChangeCancel,omitempty"`
	ChangeEndTime    *int64                                        `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	ChangeObject     *string                                       `json:"ChangeObject,omitempty" xml:"ChangeObject,omitempty"`
	ChangeOptType    *string                                       `json:"ChangeOptType,omitempty" xml:"ChangeOptType,omitempty"`
	ChangeResult     *string                                       `json:"ChangeResult,omitempty" xml:"ChangeResult,omitempty"`
	ChangeStartTime  *int64                                        `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
	ChangeStatus     *string                                       `json:"ChangeStatus,omitempty" xml:"ChangeStatus,omitempty"`
	ChangeSystem     *string                                       `json:"ChangeSystem,omitempty" xml:"ChangeSystem,omitempty"`
	ChangeTimes      []*SafeChangeQueryResponseBodyDataChangeTimes `json:"ChangeTimes,omitempty" xml:"ChangeTimes,omitempty" type:"Repeated"`
	ChangeTitle      *string                                       `json:"ChangeTitle,omitempty" xml:"ChangeTitle,omitempty"`
	CheckResultUrl   *string                                       `json:"CheckResultUrl,omitempty" xml:"CheckResultUrl,omitempty"`
	CheckStatus      *string                                       `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckholdReason  []*string                                     `json:"CheckholdReason,omitempty" xml:"CheckholdReason,omitempty" type:"Repeated"`
	ExecutorEmpId    *string                                       `json:"ExecutorEmpId,omitempty" xml:"ExecutorEmpId,omitempty"`
	ExecutorEmpName  *string                                       `json:"ExecutorEmpName,omitempty" xml:"ExecutorEmpName,omitempty"`
	SourceOrderId    *string                                       `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
}

func (s SafeChangeQueryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *SafeChangeQueryResponseBodyData) SetApproveResultUrl(v string) *SafeChangeQueryResponseBodyData {
	s.ApproveResultUrl = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetApproveStatus(v string) *SafeChangeQueryResponseBodyData {
	s.ApproveStatus = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetChangeCancel(v string) *SafeChangeQueryResponseBodyData {
	s.ChangeCancel = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetChangeEndTime(v int64) *SafeChangeQueryResponseBodyData {
	s.ChangeEndTime = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetChangeObject(v string) *SafeChangeQueryResponseBodyData {
	s.ChangeObject = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetChangeOptType(v string) *SafeChangeQueryResponseBodyData {
	s.ChangeOptType = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetChangeResult(v string) *SafeChangeQueryResponseBodyData {
	s.ChangeResult = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetChangeStartTime(v int64) *SafeChangeQueryResponseBodyData {
	s.ChangeStartTime = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetChangeStatus(v string) *SafeChangeQueryResponseBodyData {
	s.ChangeStatus = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetChangeSystem(v string) *SafeChangeQueryResponseBodyData {
	s.ChangeSystem = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetChangeTimes(v []*SafeChangeQueryResponseBodyDataChangeTimes) *SafeChangeQueryResponseBodyData {
	s.ChangeTimes = v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetChangeTitle(v string) *SafeChangeQueryResponseBodyData {
	s.ChangeTitle = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetCheckResultUrl(v string) *SafeChangeQueryResponseBodyData {
	s.CheckResultUrl = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetCheckStatus(v string) *SafeChangeQueryResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetCheckholdReason(v []*string) *SafeChangeQueryResponseBodyData {
	s.CheckholdReason = v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetExecutorEmpId(v string) *SafeChangeQueryResponseBodyData {
	s.ExecutorEmpId = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetExecutorEmpName(v string) *SafeChangeQueryResponseBodyData {
	s.ExecutorEmpName = &v
	return s
}

func (s *SafeChangeQueryResponseBodyData) SetSourceOrderId(v string) *SafeChangeQueryResponseBodyData {
	s.SourceOrderId = &v
	return s
}

type SafeChangeQueryResponseBodyDataChangeTimes struct {
	ChangeEndTime   *int64 `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	ChangeStartTime *int64 `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
}

func (s SafeChangeQueryResponseBodyDataChangeTimes) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeQueryResponseBodyDataChangeTimes) GoString() string {
	return s.String()
}

func (s *SafeChangeQueryResponseBodyDataChangeTimes) SetChangeEndTime(v int64) *SafeChangeQueryResponseBodyDataChangeTimes {
	s.ChangeEndTime = &v
	return s
}

func (s *SafeChangeQueryResponseBodyDataChangeTimes) SetChangeStartTime(v int64) *SafeChangeQueryResponseBodyDataChangeTimes {
	s.ChangeStartTime = &v
	return s
}

type SafeChangeQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SafeChangeQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SafeChangeQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeQueryResponse) GoString() string {
	return s.String()
}

func (s *SafeChangeQueryResponse) SetHeaders(v map[string]*string) *SafeChangeQueryResponse {
	s.Headers = v
	return s
}

func (s *SafeChangeQueryResponse) SetStatusCode(v int32) *SafeChangeQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *SafeChangeQueryResponse) SetBody(v *SafeChangeQueryResponseBody) *SafeChangeQueryResponse {
	s.Body = v
	return s
}

type SafeChangeQueryApproveFlowRequest struct {
	AuthKey       *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign      *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	ReqTimestamp  *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	SourceOrderId *string `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
	Stage         *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
}

func (s SafeChangeQueryApproveFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeQueryApproveFlowRequest) GoString() string {
	return s.String()
}

func (s *SafeChangeQueryApproveFlowRequest) SetAuthKey(v string) *SafeChangeQueryApproveFlowRequest {
	s.AuthKey = &v
	return s
}

func (s *SafeChangeQueryApproveFlowRequest) SetAuthSign(v string) *SafeChangeQueryApproveFlowRequest {
	s.AuthSign = &v
	return s
}

func (s *SafeChangeQueryApproveFlowRequest) SetReqTimestamp(v int64) *SafeChangeQueryApproveFlowRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *SafeChangeQueryApproveFlowRequest) SetSourceOrderId(v string) *SafeChangeQueryApproveFlowRequest {
	s.SourceOrderId = &v
	return s
}

func (s *SafeChangeQueryApproveFlowRequest) SetStage(v string) *SafeChangeQueryApproveFlowRequest {
	s.Stage = &v
	return s
}

type SafeChangeQueryApproveFlowResponseBody struct {
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*SafeChangeQueryApproveFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SafeChangeQueryApproveFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeQueryApproveFlowResponseBody) GoString() string {
	return s.String()
}

func (s *SafeChangeQueryApproveFlowResponseBody) SetCode(v int32) *SafeChangeQueryApproveFlowResponseBody {
	s.Code = &v
	return s
}

func (s *SafeChangeQueryApproveFlowResponseBody) SetData(v []*SafeChangeQueryApproveFlowResponseBodyData) *SafeChangeQueryApproveFlowResponseBody {
	s.Data = v
	return s
}

func (s *SafeChangeQueryApproveFlowResponseBody) SetMessage(v string) *SafeChangeQueryApproveFlowResponseBody {
	s.Message = &v
	return s
}

func (s *SafeChangeQueryApproveFlowResponseBody) SetRequestId(v string) *SafeChangeQueryApproveFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *SafeChangeQueryApproveFlowResponseBody) SetSuccess(v bool) *SafeChangeQueryApproveFlowResponseBody {
	s.Success = &v
	return s
}

type SafeChangeQueryApproveFlowResponseBodyData struct {
	ApproveStrategy *string `json:"ApproveStrategy,omitempty" xml:"ApproveStrategy,omitempty"`
	Approver        *string `json:"Approver,omitempty" xml:"Approver,omitempty"`
	NodeName        *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeStatus      *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
}

func (s SafeChangeQueryApproveFlowResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeQueryApproveFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *SafeChangeQueryApproveFlowResponseBodyData) SetApproveStrategy(v string) *SafeChangeQueryApproveFlowResponseBodyData {
	s.ApproveStrategy = &v
	return s
}

func (s *SafeChangeQueryApproveFlowResponseBodyData) SetApprover(v string) *SafeChangeQueryApproveFlowResponseBodyData {
	s.Approver = &v
	return s
}

func (s *SafeChangeQueryApproveFlowResponseBodyData) SetNodeName(v string) *SafeChangeQueryApproveFlowResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *SafeChangeQueryApproveFlowResponseBodyData) SetNodeStatus(v string) *SafeChangeQueryApproveFlowResponseBodyData {
	s.NodeStatus = &v
	return s
}

type SafeChangeQueryApproveFlowResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SafeChangeQueryApproveFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SafeChangeQueryApproveFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeQueryApproveFlowResponse) GoString() string {
	return s.String()
}

func (s *SafeChangeQueryApproveFlowResponse) SetHeaders(v map[string]*string) *SafeChangeQueryApproveFlowResponse {
	s.Headers = v
	return s
}

func (s *SafeChangeQueryApproveFlowResponse) SetStatusCode(v int32) *SafeChangeQueryApproveFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *SafeChangeQueryApproveFlowResponse) SetBody(v *SafeChangeQueryApproveFlowResponseBody) *SafeChangeQueryApproveFlowResponse {
	s.Body = v
	return s
}

type SafeChangeStartRequest struct {
	AuthKey         *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign        *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	ChangeEndTime   *int64  `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	ChangeObject    *string `json:"ChangeObject,omitempty" xml:"ChangeObject,omitempty"`
	ChangeOptType   *string `json:"ChangeOptType,omitempty" xml:"ChangeOptType,omitempty"`
	ChangeStartTime *int64  `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
	ChangeTitle     *string `json:"ChangeTitle,omitempty" xml:"ChangeTitle,omitempty"`
	CreatorEmpId    *string `json:"CreatorEmpId,omitempty" xml:"CreatorEmpId,omitempty"`
	CurBatchNo      *int32  `json:"CurBatchNo,omitempty" xml:"CurBatchNo,omitempty"`
	ExecutorEmpId   *string `json:"ExecutorEmpId,omitempty" xml:"ExecutorEmpId,omitempty"`
	ReqTimestamp    *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	SourceOrderId   *string `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
	TotalBatchNo    *int32  `json:"TotalBatchNo,omitempty" xml:"TotalBatchNo,omitempty"`
}

func (s SafeChangeStartRequest) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeStartRequest) GoString() string {
	return s.String()
}

func (s *SafeChangeStartRequest) SetAuthKey(v string) *SafeChangeStartRequest {
	s.AuthKey = &v
	return s
}

func (s *SafeChangeStartRequest) SetAuthSign(v string) *SafeChangeStartRequest {
	s.AuthSign = &v
	return s
}

func (s *SafeChangeStartRequest) SetChangeEndTime(v int64) *SafeChangeStartRequest {
	s.ChangeEndTime = &v
	return s
}

func (s *SafeChangeStartRequest) SetChangeObject(v string) *SafeChangeStartRequest {
	s.ChangeObject = &v
	return s
}

func (s *SafeChangeStartRequest) SetChangeOptType(v string) *SafeChangeStartRequest {
	s.ChangeOptType = &v
	return s
}

func (s *SafeChangeStartRequest) SetChangeStartTime(v int64) *SafeChangeStartRequest {
	s.ChangeStartTime = &v
	return s
}

func (s *SafeChangeStartRequest) SetChangeTitle(v string) *SafeChangeStartRequest {
	s.ChangeTitle = &v
	return s
}

func (s *SafeChangeStartRequest) SetCreatorEmpId(v string) *SafeChangeStartRequest {
	s.CreatorEmpId = &v
	return s
}

func (s *SafeChangeStartRequest) SetCurBatchNo(v int32) *SafeChangeStartRequest {
	s.CurBatchNo = &v
	return s
}

func (s *SafeChangeStartRequest) SetExecutorEmpId(v string) *SafeChangeStartRequest {
	s.ExecutorEmpId = &v
	return s
}

func (s *SafeChangeStartRequest) SetReqTimestamp(v int64) *SafeChangeStartRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *SafeChangeStartRequest) SetSourceOrderId(v string) *SafeChangeStartRequest {
	s.SourceOrderId = &v
	return s
}

func (s *SafeChangeStartRequest) SetTotalBatchNo(v int32) *SafeChangeStartRequest {
	s.TotalBatchNo = &v
	return s
}

type SafeChangeStartResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SafeChangeStartResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SafeChangeStartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeStartResponseBody) GoString() string {
	return s.String()
}

func (s *SafeChangeStartResponseBody) SetCode(v int32) *SafeChangeStartResponseBody {
	s.Code = &v
	return s
}

func (s *SafeChangeStartResponseBody) SetData(v *SafeChangeStartResponseBodyData) *SafeChangeStartResponseBody {
	s.Data = v
	return s
}

func (s *SafeChangeStartResponseBody) SetMessage(v string) *SafeChangeStartResponseBody {
	s.Message = &v
	return s
}

func (s *SafeChangeStartResponseBody) SetRequestId(v string) *SafeChangeStartResponseBody {
	s.RequestId = &v
	return s
}

func (s *SafeChangeStartResponseBody) SetSuccess(v bool) *SafeChangeStartResponseBody {
	s.Success = &v
	return s
}

type SafeChangeStartResponseBodyData struct {
	ApproveResultUrl *string `json:"ApproveResultUrl,omitempty" xml:"ApproveResultUrl,omitempty"`
	CheckResultUrl   *string `json:"CheckResultUrl,omitempty" xml:"CheckResultUrl,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubSatus         *string `json:"SubSatus,omitempty" xml:"SubSatus,omitempty"`
}

func (s SafeChangeStartResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeStartResponseBodyData) GoString() string {
	return s.String()
}

func (s *SafeChangeStartResponseBodyData) SetApproveResultUrl(v string) *SafeChangeStartResponseBodyData {
	s.ApproveResultUrl = &v
	return s
}

func (s *SafeChangeStartResponseBodyData) SetCheckResultUrl(v string) *SafeChangeStartResponseBodyData {
	s.CheckResultUrl = &v
	return s
}

func (s *SafeChangeStartResponseBodyData) SetStatus(v string) *SafeChangeStartResponseBodyData {
	s.Status = &v
	return s
}

func (s *SafeChangeStartResponseBodyData) SetSubSatus(v string) *SafeChangeStartResponseBodyData {
	s.SubSatus = &v
	return s
}

type SafeChangeStartResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SafeChangeStartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SafeChangeStartResponse) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeStartResponse) GoString() string {
	return s.String()
}

func (s *SafeChangeStartResponse) SetHeaders(v map[string]*string) *SafeChangeStartResponse {
	s.Headers = v
	return s
}

func (s *SafeChangeStartResponse) SetStatusCode(v int32) *SafeChangeStartResponse {
	s.StatusCode = &v
	return s
}

func (s *SafeChangeStartResponse) SetBody(v *SafeChangeStartResponseBody) *SafeChangeStartResponse {
	s.Body = v
	return s
}

type SafeChangeStartApproveRequest struct {
	AuthKey       *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign      *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	CreatorEmpId  *string `json:"CreatorEmpId,omitempty" xml:"CreatorEmpId,omitempty"`
	ExtraInfo     *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	ReqTimestamp  *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	SourceOrderId *string `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
}

func (s SafeChangeStartApproveRequest) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeStartApproveRequest) GoString() string {
	return s.String()
}

func (s *SafeChangeStartApproveRequest) SetAuthKey(v string) *SafeChangeStartApproveRequest {
	s.AuthKey = &v
	return s
}

func (s *SafeChangeStartApproveRequest) SetAuthSign(v string) *SafeChangeStartApproveRequest {
	s.AuthSign = &v
	return s
}

func (s *SafeChangeStartApproveRequest) SetCreatorEmpId(v string) *SafeChangeStartApproveRequest {
	s.CreatorEmpId = &v
	return s
}

func (s *SafeChangeStartApproveRequest) SetExtraInfo(v string) *SafeChangeStartApproveRequest {
	s.ExtraInfo = &v
	return s
}

func (s *SafeChangeStartApproveRequest) SetReqTimestamp(v int64) *SafeChangeStartApproveRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *SafeChangeStartApproveRequest) SetSourceOrderId(v string) *SafeChangeStartApproveRequest {
	s.SourceOrderId = &v
	return s
}

type SafeChangeStartApproveResponseBody struct {
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SafeChangeStartApproveResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SafeChangeStartApproveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeStartApproveResponseBody) GoString() string {
	return s.String()
}

func (s *SafeChangeStartApproveResponseBody) SetCode(v int32) *SafeChangeStartApproveResponseBody {
	s.Code = &v
	return s
}

func (s *SafeChangeStartApproveResponseBody) SetData(v *SafeChangeStartApproveResponseBodyData) *SafeChangeStartApproveResponseBody {
	s.Data = v
	return s
}

func (s *SafeChangeStartApproveResponseBody) SetMessage(v string) *SafeChangeStartApproveResponseBody {
	s.Message = &v
	return s
}

func (s *SafeChangeStartApproveResponseBody) SetRequestId(v string) *SafeChangeStartApproveResponseBody {
	s.RequestId = &v
	return s
}

func (s *SafeChangeStartApproveResponseBody) SetSuccess(v bool) *SafeChangeStartApproveResponseBody {
	s.Success = &v
	return s
}

type SafeChangeStartApproveResponseBodyData struct {
	ApproveStatus *string `json:"ApproveStatus,omitempty" xml:"ApproveStatus,omitempty"`
}

func (s SafeChangeStartApproveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeStartApproveResponseBodyData) GoString() string {
	return s.String()
}

func (s *SafeChangeStartApproveResponseBodyData) SetApproveStatus(v string) *SafeChangeStartApproveResponseBodyData {
	s.ApproveStatus = &v
	return s
}

type SafeChangeStartApproveResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SafeChangeStartApproveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SafeChangeStartApproveResponse) String() string {
	return tea.Prettify(s)
}

func (s SafeChangeStartApproveResponse) GoString() string {
	return s.String()
}

func (s *SafeChangeStartApproveResponse) SetHeaders(v map[string]*string) *SafeChangeStartApproveResponse {
	s.Headers = v
	return s
}

func (s *SafeChangeStartApproveResponse) SetStatusCode(v int32) *SafeChangeStartApproveResponse {
	s.StatusCode = &v
	return s
}

func (s *SafeChangeStartApproveResponse) SetBody(v *SafeChangeStartApproveResponseBody) *SafeChangeStartApproveResponse {
	s.Body = v
	return s
}

type SafeScopeDataRequest struct {
	AuthKey        *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign       *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	Category       *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CodeList       *string `json:"CodeList,omitempty" xml:"CodeList,omitempty"`
	Factor         *string `json:"Factor,omitempty" xml:"Factor,omitempty"`
	GroupBy        *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	IdList         *string `json:"IdList,omitempty" xml:"IdList,omitempty"`
	Item           *string `json:"Item,omitempty" xml:"Item,omitempty"`
	Limit          *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NeedTotalCount *bool   `json:"NeedTotalCount,omitempty" xml:"NeedTotalCount,omitempty"`
	OrderBy        *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderDirection *string `json:"OrderDirection,omitempty" xml:"OrderDirection,omitempty"`
	Page           *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	ParentCode     *string `json:"ParentCode,omitempty" xml:"ParentCode,omitempty"`
	ParentId       *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	ProductCode    *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductId      *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RegionNameEn   *string `json:"RegionNameEn,omitempty" xml:"RegionNameEn,omitempty"`
	ReqTimestamp   *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	SearchValue    *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	Type           *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid            *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s SafeScopeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SafeScopeDataRequest) GoString() string {
	return s.String()
}

func (s *SafeScopeDataRequest) SetAuthKey(v string) *SafeScopeDataRequest {
	s.AuthKey = &v
	return s
}

func (s *SafeScopeDataRequest) SetAuthSign(v string) *SafeScopeDataRequest {
	s.AuthSign = &v
	return s
}

func (s *SafeScopeDataRequest) SetCategory(v string) *SafeScopeDataRequest {
	s.Category = &v
	return s
}

func (s *SafeScopeDataRequest) SetCodeList(v string) *SafeScopeDataRequest {
	s.CodeList = &v
	return s
}

func (s *SafeScopeDataRequest) SetFactor(v string) *SafeScopeDataRequest {
	s.Factor = &v
	return s
}

func (s *SafeScopeDataRequest) SetGroupBy(v string) *SafeScopeDataRequest {
	s.GroupBy = &v
	return s
}

func (s *SafeScopeDataRequest) SetIdList(v string) *SafeScopeDataRequest {
	s.IdList = &v
	return s
}

func (s *SafeScopeDataRequest) SetItem(v string) *SafeScopeDataRequest {
	s.Item = &v
	return s
}

func (s *SafeScopeDataRequest) SetLimit(v int32) *SafeScopeDataRequest {
	s.Limit = &v
	return s
}

func (s *SafeScopeDataRequest) SetNeedTotalCount(v bool) *SafeScopeDataRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *SafeScopeDataRequest) SetOrderBy(v string) *SafeScopeDataRequest {
	s.OrderBy = &v
	return s
}

func (s *SafeScopeDataRequest) SetOrderDirection(v string) *SafeScopeDataRequest {
	s.OrderDirection = &v
	return s
}

func (s *SafeScopeDataRequest) SetPage(v int32) *SafeScopeDataRequest {
	s.Page = &v
	return s
}

func (s *SafeScopeDataRequest) SetParentCode(v string) *SafeScopeDataRequest {
	s.ParentCode = &v
	return s
}

func (s *SafeScopeDataRequest) SetParentId(v int64) *SafeScopeDataRequest {
	s.ParentId = &v
	return s
}

func (s *SafeScopeDataRequest) SetProductCode(v string) *SafeScopeDataRequest {
	s.ProductCode = &v
	return s
}

func (s *SafeScopeDataRequest) SetProductId(v int64) *SafeScopeDataRequest {
	s.ProductId = &v
	return s
}

func (s *SafeScopeDataRequest) SetRegionNameEn(v string) *SafeScopeDataRequest {
	s.RegionNameEn = &v
	return s
}

func (s *SafeScopeDataRequest) SetReqTimestamp(v int64) *SafeScopeDataRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *SafeScopeDataRequest) SetSearchValue(v string) *SafeScopeDataRequest {
	s.SearchValue = &v
	return s
}

func (s *SafeScopeDataRequest) SetType(v int32) *SafeScopeDataRequest {
	s.Type = &v
	return s
}

func (s *SafeScopeDataRequest) SetUid(v string) *SafeScopeDataRequest {
	s.Uid = &v
	return s
}

type SafeScopeDataResponseBody struct {
	Code      *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SafeScopeDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SafeScopeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SafeScopeDataResponseBody) GoString() string {
	return s.String()
}

func (s *SafeScopeDataResponseBody) SetCode(v int32) *SafeScopeDataResponseBody {
	s.Code = &v
	return s
}

func (s *SafeScopeDataResponseBody) SetData(v *SafeScopeDataResponseBodyData) *SafeScopeDataResponseBody {
	s.Data = v
	return s
}

func (s *SafeScopeDataResponseBody) SetMessage(v string) *SafeScopeDataResponseBody {
	s.Message = &v
	return s
}

func (s *SafeScopeDataResponseBody) SetRequestId(v string) *SafeScopeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *SafeScopeDataResponseBody) SetSuccess(v bool) *SafeScopeDataResponseBody {
	s.Success = &v
	return s
}

type SafeScopeDataResponseBodyData struct {
	Data       []interface{}                            `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ExtraInfo  map[string]*string                       `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Pagination *SafeScopeDataResponseBodyDataPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	Total      *int64                                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SafeScopeDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SafeScopeDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *SafeScopeDataResponseBodyData) SetData(v []interface{}) *SafeScopeDataResponseBodyData {
	s.Data = v
	return s
}

func (s *SafeScopeDataResponseBodyData) SetExtraInfo(v map[string]*string) *SafeScopeDataResponseBodyData {
	s.ExtraInfo = v
	return s
}

func (s *SafeScopeDataResponseBodyData) SetPagination(v *SafeScopeDataResponseBodyDataPagination) *SafeScopeDataResponseBodyData {
	s.Pagination = v
	return s
}

func (s *SafeScopeDataResponseBodyData) SetTotal(v int64) *SafeScopeDataResponseBodyData {
	s.Total = &v
	return s
}

type SafeScopeDataResponseBodyDataPagination struct {
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Page  *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s SafeScopeDataResponseBodyDataPagination) String() string {
	return tea.Prettify(s)
}

func (s SafeScopeDataResponseBodyDataPagination) GoString() string {
	return s.String()
}

func (s *SafeScopeDataResponseBodyDataPagination) SetLimit(v int32) *SafeScopeDataResponseBodyDataPagination {
	s.Limit = &v
	return s
}

func (s *SafeScopeDataResponseBodyDataPagination) SetPage(v int32) *SafeScopeDataResponseBodyDataPagination {
	s.Page = &v
	return s
}

type SafeScopeDataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SafeScopeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SafeScopeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SafeScopeDataResponse) GoString() string {
	return s.String()
}

func (s *SafeScopeDataResponse) SetHeaders(v map[string]*string) *SafeScopeDataResponse {
	s.Headers = v
	return s
}

func (s *SafeScopeDataResponse) SetStatusCode(v int32) *SafeScopeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *SafeScopeDataResponse) SetBody(v *SafeScopeDataResponseBody) *SafeScopeDataResponse {
	s.Body = v
	return s
}

type StartApproveRequest struct {
	AuthKey       *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign      *string `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	CreatorEmpId  *string `json:"CreatorEmpId,omitempty" xml:"CreatorEmpId,omitempty"`
	ExtraInfo     *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	ReqTimestamp  *int64  `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	SourceOrderId *string `json:"SourceOrderId,omitempty" xml:"SourceOrderId,omitempty"`
}

func (s StartApproveRequest) String() string {
	return tea.Prettify(s)
}

func (s StartApproveRequest) GoString() string {
	return s.String()
}

func (s *StartApproveRequest) SetAuthKey(v string) *StartApproveRequest {
	s.AuthKey = &v
	return s
}

func (s *StartApproveRequest) SetAuthSign(v string) *StartApproveRequest {
	s.AuthSign = &v
	return s
}

func (s *StartApproveRequest) SetCreatorEmpId(v string) *StartApproveRequest {
	s.CreatorEmpId = &v
	return s
}

func (s *StartApproveRequest) SetExtraInfo(v string) *StartApproveRequest {
	s.ExtraInfo = &v
	return s
}

func (s *StartApproveRequest) SetReqTimestamp(v int64) *StartApproveRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *StartApproveRequest) SetSourceOrderId(v string) *StartApproveRequest {
	s.SourceOrderId = &v
	return s
}

type StartApproveResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartApproveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartApproveResponseBody) GoString() string {
	return s.String()
}

func (s *StartApproveResponseBody) SetCode(v int32) *StartApproveResponseBody {
	s.Code = &v
	return s
}

func (s *StartApproveResponseBody) SetData(v string) *StartApproveResponseBody {
	s.Data = &v
	return s
}

func (s *StartApproveResponseBody) SetMessage(v string) *StartApproveResponseBody {
	s.Message = &v
	return s
}

func (s *StartApproveResponseBody) SetRequestId(v string) *StartApproveResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartApproveResponseBody) SetSuccess(v bool) *StartApproveResponseBody {
	s.Success = &v
	return s
}

type StartApproveResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartApproveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartApproveResponse) String() string {
	return tea.Prettify(s)
}

func (s StartApproveResponse) GoString() string {
	return s.String()
}

func (s *StartApproveResponse) SetHeaders(v map[string]*string) *StartApproveResponse {
	s.Headers = v
	return s
}

func (s *StartApproveResponse) SetStatusCode(v int32) *StartApproveResponse {
	s.StatusCode = &v
	return s
}

func (s *StartApproveResponse) SetBody(v *StartApproveResponseBody) *StartApproveResponse {
	s.Body = v
	return s
}

type SyncProductRequest struct {
	AuthKey         *string                              `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSign        *string                              `json:"AuthSign,omitempty" xml:"AuthSign,omitempty"`
	ReqTimestamp    *int64                               `json:"ReqTimestamp,omitempty" xml:"ReqTimestamp,omitempty"`
	SyncProductList []*SyncProductRequestSyncProductList `json:"SyncProductList,omitempty" xml:"SyncProductList,omitempty" type:"Repeated"`
}

func (s SyncProductRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncProductRequest) GoString() string {
	return s.String()
}

func (s *SyncProductRequest) SetAuthKey(v string) *SyncProductRequest {
	s.AuthKey = &v
	return s
}

func (s *SyncProductRequest) SetAuthSign(v string) *SyncProductRequest {
	s.AuthSign = &v
	return s
}

func (s *SyncProductRequest) SetReqTimestamp(v int64) *SyncProductRequest {
	s.ReqTimestamp = &v
	return s
}

func (s *SyncProductRequest) SetSyncProductList(v []*SyncProductRequestSyncProductList) *SyncProductRequest {
	s.SyncProductList = v
	return s
}

type SyncProductRequestSyncProductList struct {
	Code             *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	InnerProductList []*SyncProductRequestSyncProductListInnerProductList `json:"InnerProductList,omitempty" xml:"InnerProductList,omitempty" type:"Repeated"`
	Name             *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SyncProductRequestSyncProductList) String() string {
	return tea.Prettify(s)
}

func (s SyncProductRequestSyncProductList) GoString() string {
	return s.String()
}

func (s *SyncProductRequestSyncProductList) SetCode(v string) *SyncProductRequestSyncProductList {
	s.Code = &v
	return s
}

func (s *SyncProductRequestSyncProductList) SetInnerProductList(v []*SyncProductRequestSyncProductListInnerProductList) *SyncProductRequestSyncProductList {
	s.InnerProductList = v
	return s
}

func (s *SyncProductRequestSyncProductList) SetName(v string) *SyncProductRequestSyncProductList {
	s.Name = &v
	return s
}

type SyncProductRequestSyncProductListInnerProductList struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SyncProductRequestSyncProductListInnerProductList) String() string {
	return tea.Prettify(s)
}

func (s SyncProductRequestSyncProductListInnerProductList) GoString() string {
	return s.String()
}

func (s *SyncProductRequestSyncProductListInnerProductList) SetCode(v string) *SyncProductRequestSyncProductListInnerProductList {
	s.Code = &v
	return s
}

func (s *SyncProductRequestSyncProductListInnerProductList) SetName(v string) *SyncProductRequestSyncProductListInnerProductList {
	s.Name = &v
	return s
}

type SyncProductResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncProductResponseBody) GoString() string {
	return s.String()
}

func (s *SyncProductResponseBody) SetCode(v int32) *SyncProductResponseBody {
	s.Code = &v
	return s
}

func (s *SyncProductResponseBody) SetData(v bool) *SyncProductResponseBody {
	s.Data = &v
	return s
}

func (s *SyncProductResponseBody) SetMessage(v string) *SyncProductResponseBody {
	s.Message = &v
	return s
}

func (s *SyncProductResponseBody) SetRequestId(v string) *SyncProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncProductResponseBody) SetSuccess(v bool) *SyncProductResponseBody {
	s.Success = &v
	return s
}

type SyncProductResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncProductResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncProductResponse) GoString() string {
	return s.String()
}

func (s *SyncProductResponse) SetHeaders(v map[string]*string) *SyncProductResponse {
	s.Headers = v
	return s
}

func (s *SyncProductResponse) SetStatusCode(v int32) *SyncProductResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncProductResponse) SetBody(v *SyncProductResponseBody) *SyncProductResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("safe"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 变更取消
//
// @param request - ChangeCancelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeCancelResponse
func (client *Client) ChangeCancelWithOptions(request *ChangeCancelRequest, runtime *util.RuntimeOptions) (_result *ChangeCancelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		query["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		query["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		query["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		query["SourceOrderId"] = request.SourceOrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeCancel"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeCancelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更取消
//
// @param request - ChangeCancelRequest
//
// @return ChangeCancelResponse
func (client *Client) ChangeCancel(request *ChangeCancelRequest) (_result *ChangeCancelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeCancelResponse{}
	_body, _err := client.ChangeCancelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 安全变更check
//
// @param tmpReq - ChangeCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeCheckResponse
func (client *Client) ChangeCheckWithOptions(tmpReq *ChangeCheckRequest, runtime *util.RuntimeOptions) (_result *ChangeCheckResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ChangeCheckShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DamagedChangeNotices)) {
		request.DamagedChangeNoticesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DamagedChangeNotices, tea.String("DamagedChangeNotices"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AffectCustomer)) {
		body["AffectCustomer"] = request.AffectCustomer
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApproveFlowParam)) {
		bodyFlat["ApproveFlowParam"] = request.ApproveFlowParam
	}

	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.BgCustomTemplateExtraDTO)) {
		bodyFlat["BgCustomTemplateExtraDTO"] = request.BgCustomTemplateExtraDTO
	}

	if !tea.BoolValue(util.IsUnset(request.BgVid)) {
		body["BgVid"] = request.BgVid
	}

	if !tea.BoolValue(util.IsUnset(request.BlockInfos)) {
		body["BlockInfos"] = request.BlockInfos
	}

	if !tea.BoolValue(util.IsUnset(request.CallBackInfo)) {
		bodyFlat["CallBackInfo"] = request.CallBackInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeDataType)) {
		body["ChangeDataType"] = request.ChangeDataType
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeDesc)) {
		body["ChangeDesc"] = request.ChangeDesc
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeEndTime)) {
		body["ChangeEndTime"] = request.ChangeEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeEnv)) {
		body["ChangeEnv"] = request.ChangeEnv
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeItems)) {
		body["ChangeItems"] = request.ChangeItems
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeObject)) {
		body["ChangeObject"] = request.ChangeObject
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeOptSubType)) {
		body["ChangeOptSubType"] = request.ChangeOptSubType
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeOptType)) {
		body["ChangeOptType"] = request.ChangeOptType
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeReason)) {
		body["ChangeReason"] = request.ChangeReason
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeRmarks)) {
		body["ChangeRmarks"] = request.ChangeRmarks
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeSchemes)) {
		body["ChangeSchemes"] = request.ChangeSchemes
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeStartTime)) {
		body["ChangeStartTime"] = request.ChangeStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeSubTypeDesc)) {
		body["ChangeSubTypeDesc"] = request.ChangeSubTypeDesc
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeSystem)) {
		body["ChangeSystem"] = request.ChangeSystem
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeTimes)) {
		body["ChangeTimes"] = request.ChangeTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeTitle)) {
		body["ChangeTitle"] = request.ChangeTitle
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeValidation)) {
		body["ChangeValidation"] = request.ChangeValidation
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorEmpId)) {
		body["CreatorEmpId"] = request.CreatorEmpId
	}

	if !tea.BoolValue(util.IsUnset(request.DamagedChangeNoticesShrink)) {
		body["DamagedChangeNotices"] = request.DamagedChangeNoticesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorEmpId)) {
		body["ExecutorEmpId"] = request.ExecutorEmpId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraInfo)) {
		body["ExtraInfo"] = request.ExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Follower)) {
		body["Follower"] = request.Follower
	}

	if !tea.BoolValue(util.IsUnset(request.GrayStatus)) {
		body["GrayStatus"] = request.GrayStatus
	}

	if !tea.BoolValue(util.IsUnset(request.HarmChangeNoticeEnum)) {
		body["HarmChangeNoticeEnum"] = request.HarmChangeNoticeEnum
	}

	if !tea.BoolValue(util.IsUnset(request.Incidence)) {
		body["Incidence"] = request.Incidence
	}

	if !tea.BoolValue(util.IsUnset(request.InfluenceInfo)) {
		bodyFlat["InfluenceInfo"] = request.InfluenceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		bodyFlat["Instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.NeedModifyDoc)) {
		body["NeedModifyDoc"] = request.NeedModifyDoc
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		body["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.ReleasePackageInfos)) {
		body["ReleasePackageInfos"] = request.ReleasePackageInfos
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.ReuseSourceOrderId)) {
		body["ReuseSourceOrderId"] = request.ReuseSourceOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !tea.BoolValue(util.IsUnset(request.Rollback)) {
		body["Rollback"] = request.Rollback
	}

	if !tea.BoolValue(util.IsUnset(request.SourceName)) {
		body["SourceName"] = request.SourceName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		body["SourceOrderId"] = request.SourceOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUrl)) {
		body["SourceUrl"] = request.SourceUrl
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteType)) {
		body["WhiteType"] = request.WhiteType
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeCheck"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 安全变更check
//
// @param request - ChangeCheckRequest
//
// @return ChangeCheckResponse
func (client *Client) ChangeCheck(request *ChangeCheckRequest) (_result *ChangeCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeCheckResponse{}
	_body, _err := client.ChangeCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更执行end
//
// @param request - ChangeEndRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeEndResponse
func (client *Client) ChangeEndWithOptions(request *ChangeEndRequest, runtime *util.RuntimeOptions) (_result *ChangeEndResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		query["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		query["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeEndTime)) {
		query["ChangeEndTime"] = request.ChangeEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeResult)) {
		query["ChangeResult"] = request.ChangeResult
	}

	if !tea.BoolValue(util.IsUnset(request.CurBatchNo)) {
		query["CurBatchNo"] = request.CurBatchNo
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorEmpId)) {
		query["ExecutorEmpId"] = request.ExecutorEmpId
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		query["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		query["SourceOrderId"] = request.SourceOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.TotalBatchNo)) {
		query["TotalBatchNo"] = request.TotalBatchNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeEnd"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeEndResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更执行end
//
// @param request - ChangeEndRequest
//
// @return ChangeEndResponse
func (client *Client) ChangeEnd(request *ChangeEndRequest) (_result *ChangeEndResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeEndResponse{}
	_body, _err := client.ChangeEndWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更执行start
//
// @param request - ChangeStartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeStartResponse
func (client *Client) ChangeStartWithOptions(request *ChangeStartRequest, runtime *util.RuntimeOptions) (_result *ChangeStartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		query["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		query["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeEndTime)) {
		query["ChangeEndTime"] = request.ChangeEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeObject)) {
		query["ChangeObject"] = request.ChangeObject
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeOptType)) {
		query["ChangeOptType"] = request.ChangeOptType
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeStartTime)) {
		query["ChangeStartTime"] = request.ChangeStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeTitle)) {
		query["ChangeTitle"] = request.ChangeTitle
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorEmpId)) {
		query["CreatorEmpId"] = request.CreatorEmpId
	}

	if !tea.BoolValue(util.IsUnset(request.CurBatchNo)) {
		query["CurBatchNo"] = request.CurBatchNo
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorEmpId)) {
		query["ExecutorEmpId"] = request.ExecutorEmpId
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		query["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		query["SourceOrderId"] = request.SourceOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.TotalBatchNo)) {
		query["TotalBatchNo"] = request.TotalBatchNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeStart"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeStartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更执行start
//
// @param request - ChangeStartRequest
//
// @return ChangeStartResponse
func (client *Client) ChangeStart(request *ChangeStartRequest) (_result *ChangeStartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeStartResponse{}
	_body, _err := client.ChangeStartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 三方创建封网接口
//
// @param request - CreateBlockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBlockResponse
func (client *Client) CreateBlockWithOptions(request *CreateBlockRequest, runtime *util.RuntimeOptions) (_result *CreateBlockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApproveStrategyNodes)) {
		bodyFlat["ApproveStrategyNodes"] = request.ApproveStrategyNodes
	}

	if !tea.BoolValue(util.IsUnset(request.BlockId)) {
		body["BlockId"] = request.BlockId
	}

	if !tea.BoolValue(util.IsUnset(request.Director)) {
		body["Director"] = request.Director
	}

	if !tea.BoolValue(util.IsUnset(request.IsNeedApprove)) {
		body["IsNeedApprove"] = request.IsNeedApprove
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecall)) {
		body["IsRecall"] = request.IsRecall
	}

	if !tea.BoolValue(util.IsUnset(request.IsTemplate)) {
		body["IsTemplate"] = request.IsTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.LabelName)) {
		body["LabelName"] = request.LabelName
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeDesc)) {
		body["NoticeDesc"] = request.NoticeDesc
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeEnclosureInfos)) {
		bodyFlat["NoticeEnclosureInfos"] = request.NoticeEnclosureInfos
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeRequestLink)) {
		body["NoticeRequestLink"] = request.NoticeRequestLink
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeType)) {
		body["NoticeType"] = request.NoticeType
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		body["Reason"] = request.Reason
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.Scopes)) {
		bodyFlat["Scopes"] = request.Scopes
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		body["VersionId"] = request.VersionId
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBlock"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBlockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 三方创建封网接口
//
// @param request - CreateBlockRequest
//
// @return CreateBlockResponse
func (client *Client) CreateBlock(request *CreateBlockRequest) (_result *CreateBlockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBlockResponse{}
	_body, _err := client.CreateBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建操作类型
//
// @param request - CreateOperatorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOperatorResponse
func (client *Client) CreateOperatorWithOptions(request *CreateOperatorRequest, runtime *util.RuntimeOptions) (_result *CreateOperatorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.BgObject)) {
		body["BgObject"] = request.BgObject
	}

	if !tea.BoolValue(util.IsUnset(request.BgSystem)) {
		body["BgSystem"] = request.BgSystem
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.CurEmpId)) {
		body["CurEmpId"] = request.CurEmpId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NoCheck)) {
		body["NoCheck"] = request.NoCheck
	}

	if !tea.BoolValue(util.IsUnset(request.NoRisk)) {
		body["NoRisk"] = request.NoRisk
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOperator"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOperatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建操作类型
//
// @param request - CreateOperatorRequest
//
// @return CreateOperatorResponse
func (client *Client) CreateOperator(request *CreateOperatorRequest) (_result *CreateOperatorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOperatorResponse{}
	_body, _err := client.CreateOperatorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更状态查询
//
// @param request - QueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryResponse
func (client *Client) QueryWithOptions(request *QueryRequest, runtime *util.RuntimeOptions) (_result *QueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		query["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		query["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.NeedValidate)) {
		query["NeedValidate"] = request.NeedValidate
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		query["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		query["SourceOrderId"] = request.SourceOrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Query"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更状态查询
//
// @param request - QueryRequest
//
// @return QueryResponse
func (client *Client) Query(request *QueryRequest) (_result *QueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryResponse{}
	_body, _err := client.QueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询审批实例信息
//
// @param request - QueryApproveFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryApproveFlowResponse
func (client *Client) QueryApproveFlowWithOptions(request *QueryApproveFlowRequest, runtime *util.RuntimeOptions) (_result *QueryApproveFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		query["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		query["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		query["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		query["SourceOrderId"] = request.SourceOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.Stage)) {
		query["Stage"] = request.Stage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryApproveFlow"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryApproveFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审批实例信息
//
// @param request - QueryApproveFlowRequest
//
// @return QueryApproveFlowResponse
func (client *Client) QueryApproveFlow(request *QueryApproveFlowRequest) (_result *QueryApproveFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryApproveFlowResponse{}
	_body, _err := client.QueryApproveFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查封网事件
//
// @param request - QueryBlockEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBlockEventResponse
func (client *Client) QueryBlockEventWithOptions(request *QueryBlockEventRequest, runtime *util.RuntimeOptions) (_result *QueryBlockEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.BgSystemName)) {
		body["BgSystemName"] = request.BgSystemName
	}

	if !tea.BoolValue(util.IsUnset(request.BlockHarm)) {
		body["BlockHarm"] = request.BlockHarm
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.DeptNo)) {
		body["DeptNo"] = request.DeptNo
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NeedRule)) {
		body["NeedRule"] = request.NeedRule
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCodes)) {
		body["ProductCodes"] = request.ProductCodes
	}

	if !tea.BoolValue(util.IsUnset(request.RegionReqs)) {
		body["RegionReqs"] = request.RegionReqs
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryBlockEvent"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBlockEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查封网事件
//
// @param request - QueryBlockEventRequest
//
// @return QueryBlockEventResponse
func (client *Client) QueryBlockEvent(request *QueryBlockEventRequest) (_result *QueryBlockEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBlockEventResponse{}
	_body, _err := client.QueryBlockEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更单详情
//
// @param request - QueryChangeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryChangeInfoResponse
func (client *Client) QueryChangeInfoWithOptions(request *QueryChangeInfoRequest, runtime *util.RuntimeOptions) (_result *QueryChangeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.Az)) {
		body["Az"] = request.Az
	}

	if !tea.BoolValue(util.IsUnset(request.BgVid)) {
		body["BgVid"] = request.BgVid
	}

	if !tea.BoolValue(util.IsUnset(request.BuId)) {
		body["BuId"] = request.BuId
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeSystem)) {
		body["ChangeSystem"] = request.ChangeSystem
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LevelTree)) {
		bodyFlat["LevelTree"] = request.LevelTree
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		body["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		body["SourceOrderId"] = request.SourceOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryChangeInfo"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryChangeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更单详情
//
// @param request - QueryChangeInfoRequest
//
// @return QueryChangeInfoResponse
func (client *Client) QueryChangeInfo(request *QueryChangeInfoRequest) (_result *QueryChangeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryChangeInfoResponse{}
	_body, _err := client.QueryChangeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询检测详情接口
//
// @param request - QueryCheckInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCheckInfoResponse
func (client *Client) QueryCheckInfoWithOptions(request *QueryCheckInfoRequest, runtime *util.RuntimeOptions) (_result *QueryCheckInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		body["SourceOrderId"] = request.SourceOrderId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCheckInfo"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCheckInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询检测详情接口
//
// @param request - QueryCheckInfoRequest
//
// @return QueryCheckInfoResponse
func (client *Client) QueryCheckInfo(request *QueryCheckInfoRequest) (_result *QueryCheckInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCheckInfoResponse{}
	_body, _err := client.QueryCheckInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询敏感客户
//
// @param request - QueryCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomerResponse
func (client *Client) QueryCustomerWithOptions(request *QueryCustomerRequest, runtime *util.RuntimeOptions) (_result *QueryCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		body["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCustomer"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询敏感客户
//
// @param request - QueryCustomerRequest
//
// @return QueryCustomerResponse
func (client *Client) QueryCustomer(request *QueryCustomerRequest) (_result *QueryCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCustomerResponse{}
	_body, _err := client.QueryCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行单详情
//
// @param request - QueryExecuteInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryExecuteInfoResponse
func (client *Client) QueryExecuteInfoWithOptions(request *QueryExecuteInfoRequest, runtime *util.RuntimeOptions) (_result *QueryExecuteInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.Az)) {
		body["Az"] = request.Az
	}

	if !tea.BoolValue(util.IsUnset(request.BgVid)) {
		body["BgVid"] = request.BgVid
	}

	if !tea.BoolValue(util.IsUnset(request.BuId)) {
		body["BuId"] = request.BuId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExVid)) {
		body["ExVid"] = request.ExVid
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LevelTree)) {
		bodyFlat["LevelTree"] = request.LevelTree
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		body["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		body["SourceOrderId"] = request.SourceOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryExecuteInfo"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryExecuteInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行单详情
//
// @param request - QueryExecuteInfoRequest
//
// @return QueryExecuteInfoResponse
func (client *Client) QueryExecuteInfo(request *QueryExecuteInfoRequest) (_result *QueryExecuteInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryExecuteInfoResponse{}
	_body, _err := client.QueryExecuteInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询内部产品接口
//
// @param request - QueryInnerProductInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInnerProductInfoResponse
func (client *Client) QueryInnerProductInfoWithOptions(request *QueryInnerProductInfoRequest, runtime *util.RuntimeOptions) (_result *QueryInnerProductInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryInnerProductInfo"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryInnerProductInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询内部产品接口
//
// @param request - QueryInnerProductInfoRequest
//
// @return QueryInnerProductInfoResponse
func (client *Client) QueryInnerProductInfo(request *QueryInnerProductInfoRequest) (_result *QueryInnerProductInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryInnerProductInfoResponse{}
	_body, _err := client.QueryInnerProductInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryRegionAzRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRegionAzResponse
func (client *Client) QueryRegionAzWithOptions(request *QueryRegionAzRequest, runtime *util.RuntimeOptions) (_result *QueryRegionAzResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		body["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRegionAz"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRegionAzResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryRegionAzRequest
//
// @return QueryRegionAzResponse
func (client *Client) QueryRegionAz(request *QueryRegionAzRequest) (_result *QueryRegionAzResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRegionAzResponse{}
	_body, _err := client.QueryRegionAzWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更取消接口
//
// @param request - SafeChangeCancelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SafeChangeCancelResponse
func (client *Client) SafeChangeCancelWithOptions(request *SafeChangeCancelRequest, runtime *util.RuntimeOptions) (_result *SafeChangeCancelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.OperateEmpNo)) {
		body["OperateEmpNo"] = request.OperateEmpNo
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		body["SourceOrderId"] = request.SourceOrderId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SafeChangeCancel"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SafeChangeCancelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更取消接口
//
// @param request - SafeChangeCancelRequest
//
// @return SafeChangeCancelResponse
func (client *Client) SafeChangeCancel(request *SafeChangeCancelRequest) (_result *SafeChangeCancelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SafeChangeCancelResponse{}
	_body, _err := client.SafeChangeCancelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更check接口
//
// @param request - SafeChangeCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SafeChangeCheckResponse
func (client *Client) SafeChangeCheckWithOptions(request *SafeChangeCheckRequest, runtime *util.RuntimeOptions) (_result *SafeChangeCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Checker)) {
		query["Checker"] = request.Checker
	}

	if !tea.BoolValue(util.IsUnset(request.HarmChangeNoticeEnum)) {
		query["HarmChangeNoticeEnum"] = request.HarmChangeNoticeEnum
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AffectCustomer)) {
		body["AffectCustomer"] = request.AffectCustomer
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApproveFlowParam)) {
		bodyFlat["ApproveFlowParam"] = request.ApproveFlowParam
	}

	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.BgCustomTemplateExtraDTO)) {
		bodyFlat["BgCustomTemplateExtraDTO"] = request.BgCustomTemplateExtraDTO
	}

	if !tea.BoolValue(util.IsUnset(request.BlockInfos)) {
		body["BlockInfos"] = request.BlockInfos
	}

	if !tea.BoolValue(util.IsUnset(request.CallBackInfo)) {
		bodyFlat["CallBackInfo"] = request.CallBackInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeDataType)) {
		body["ChangeDataType"] = request.ChangeDataType
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeDesc)) {
		body["ChangeDesc"] = request.ChangeDesc
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeEndTime)) {
		body["ChangeEndTime"] = request.ChangeEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeEnv)) {
		body["ChangeEnv"] = request.ChangeEnv
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeItems)) {
		body["ChangeItems"] = request.ChangeItems
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeObject)) {
		body["ChangeObject"] = request.ChangeObject
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeOptSubType)) {
		body["ChangeOptSubType"] = request.ChangeOptSubType
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeOptType)) {
		body["ChangeOptType"] = request.ChangeOptType
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeReason)) {
		body["ChangeReason"] = request.ChangeReason
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeRmarks)) {
		body["ChangeRmarks"] = request.ChangeRmarks
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeSchemes)) {
		body["ChangeSchemes"] = request.ChangeSchemes
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeStartTime)) {
		body["ChangeStartTime"] = request.ChangeStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeSubTypeDesc)) {
		body["ChangeSubTypeDesc"] = request.ChangeSubTypeDesc
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeSystem)) {
		body["ChangeSystem"] = request.ChangeSystem
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeTimes)) {
		body["ChangeTimes"] = request.ChangeTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeTitle)) {
		body["ChangeTitle"] = request.ChangeTitle
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeValidation)) {
		body["ChangeValidation"] = request.ChangeValidation
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorEmpId)) {
		body["CreatorEmpId"] = request.CreatorEmpId
	}

	if !tea.BoolValue(util.IsUnset(request.DamagedChangeNotices)) {
		bodyFlat["DamagedChangeNotices"] = request.DamagedChangeNotices
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorEmpId)) {
		body["ExecutorEmpId"] = request.ExecutorEmpId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraInfo)) {
		body["ExtraInfo"] = request.ExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Follower)) {
		body["Follower"] = request.Follower
	}

	if !tea.BoolValue(util.IsUnset(request.GrayStatus)) {
		body["GrayStatus"] = request.GrayStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Incidence)) {
		body["Incidence"] = request.Incidence
	}

	if !tea.BoolValue(util.IsUnset(request.InfluenceInfo)) {
		bodyFlat["InfluenceInfo"] = request.InfluenceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		bodyFlat["Instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.NeedModifyDoc)) {
		body["NeedModifyDoc"] = request.NeedModifyDoc
	}

	if !tea.BoolValue(util.IsUnset(request.OperateEmpNo)) {
		body["OperateEmpNo"] = request.OperateEmpNo
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		body["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.ReleasePackageInfos)) {
		body["ReleasePackageInfos"] = request.ReleasePackageInfos
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.ReuseSourceOrderId)) {
		body["ReuseSourceOrderId"] = request.ReuseSourceOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !tea.BoolValue(util.IsUnset(request.Rollback)) {
		body["Rollback"] = request.Rollback
	}

	if !tea.BoolValue(util.IsUnset(request.SourceName)) {
		body["SourceName"] = request.SourceName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		body["SourceOrderId"] = request.SourceOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUrl)) {
		body["SourceUrl"] = request.SourceUrl
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteType)) {
		body["whiteType"] = request.WhiteType
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SafeChangeCheck"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SafeChangeCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更check接口
//
// @param request - SafeChangeCheckRequest
//
// @return SafeChangeCheckResponse
func (client *Client) SafeChangeCheck(request *SafeChangeCheckRequest) (_result *SafeChangeCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SafeChangeCheckResponse{}
	_body, _err := client.SafeChangeCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更End接口
//
// @param request - SafeChangeEndRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SafeChangeEndResponse
func (client *Client) SafeChangeEndWithOptions(request *SafeChangeEndRequest, runtime *util.RuntimeOptions) (_result *SafeChangeEndResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeEndTime)) {
		body["ChangeEndTime"] = request.ChangeEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeResult)) {
		body["ChangeResult"] = request.ChangeResult
	}

	if !tea.BoolValue(util.IsUnset(request.CurBatchNo)) {
		body["CurBatchNo"] = request.CurBatchNo
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorEmpId)) {
		body["ExecutorEmpId"] = request.ExecutorEmpId
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		body["SourceOrderId"] = request.SourceOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.TotalBatchNo)) {
		body["TotalBatchNo"] = request.TotalBatchNo
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SafeChangeEnd"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SafeChangeEndResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更End接口
//
// @param request - SafeChangeEndRequest
//
// @return SafeChangeEndResponse
func (client *Client) SafeChangeEnd(request *SafeChangeEndRequest) (_result *SafeChangeEndResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SafeChangeEndResponse{}
	_body, _err := client.SafeChangeEndWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更单查询
//
// @param request - SafeChangeQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SafeChangeQueryResponse
func (client *Client) SafeChangeQueryWithOptions(request *SafeChangeQueryRequest, runtime *util.RuntimeOptions) (_result *SafeChangeQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReturnType)) {
		query["ReturnType"] = request.ReturnType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.NeedValidate)) {
		body["NeedValidate"] = request.NeedValidate
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		body["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		body["SourceOrderId"] = request.SourceOrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SafeChangeQuery"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SafeChangeQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更单查询
//
// @param request - SafeChangeQueryRequest
//
// @return SafeChangeQueryResponse
func (client *Client) SafeChangeQuery(request *SafeChangeQueryRequest) (_result *SafeChangeQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SafeChangeQueryResponse{}
	_body, _err := client.SafeChangeQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询审批实例信息
//
// @param request - SafeChangeQueryApproveFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SafeChangeQueryApproveFlowResponse
func (client *Client) SafeChangeQueryApproveFlowWithOptions(request *SafeChangeQueryApproveFlowRequest, runtime *util.RuntimeOptions) (_result *SafeChangeQueryApproveFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		body["SourceOrderId"] = request.SourceOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.Stage)) {
		body["Stage"] = request.Stage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SafeChangeQueryApproveFlow"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SafeChangeQueryApproveFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审批实例信息
//
// @param request - SafeChangeQueryApproveFlowRequest
//
// @return SafeChangeQueryApproveFlowResponse
func (client *Client) SafeChangeQueryApproveFlow(request *SafeChangeQueryApproveFlowRequest) (_result *SafeChangeQueryApproveFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SafeChangeQueryApproveFlowResponse{}
	_body, _err := client.SafeChangeQueryApproveFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更Start接口
//
// @param request - SafeChangeStartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SafeChangeStartResponse
func (client *Client) SafeChangeStartWithOptions(request *SafeChangeStartRequest, runtime *util.RuntimeOptions) (_result *SafeChangeStartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeEndTime)) {
		body["ChangeEndTime"] = request.ChangeEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeObject)) {
		body["ChangeObject"] = request.ChangeObject
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeOptType)) {
		body["ChangeOptType"] = request.ChangeOptType
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeStartTime)) {
		body["ChangeStartTime"] = request.ChangeStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeTitle)) {
		body["ChangeTitle"] = request.ChangeTitle
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorEmpId)) {
		body["CreatorEmpId"] = request.CreatorEmpId
	}

	if !tea.BoolValue(util.IsUnset(request.CurBatchNo)) {
		body["CurBatchNo"] = request.CurBatchNo
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorEmpId)) {
		body["ExecutorEmpId"] = request.ExecutorEmpId
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		body["SourceOrderId"] = request.SourceOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.TotalBatchNo)) {
		body["TotalBatchNo"] = request.TotalBatchNo
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SafeChangeStart"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SafeChangeStartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更Start接口
//
// @param request - SafeChangeStartRequest
//
// @return SafeChangeStartResponse
func (client *Client) SafeChangeStart(request *SafeChangeStartRequest) (_result *SafeChangeStartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SafeChangeStartResponse{}
	_body, _err := client.SafeChangeStartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交审批
//
// @param request - SafeChangeStartApproveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SafeChangeStartApproveResponse
func (client *Client) SafeChangeStartApproveWithOptions(request *SafeChangeStartApproveRequest, runtime *util.RuntimeOptions) (_result *SafeChangeStartApproveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorEmpId)) {
		body["CreatorEmpId"] = request.CreatorEmpId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraInfo)) {
		body["ExtraInfo"] = request.ExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		body["SourceOrderId"] = request.SourceOrderId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SafeChangeStartApprove"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SafeChangeStartApproveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交审批
//
// @param request - SafeChangeStartApproveRequest
//
// @return SafeChangeStartApproveResponse
func (client *Client) SafeChangeStartApprove(request *SafeChangeStartApproveRequest) (_result *SafeChangeStartApproveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SafeChangeStartApproveResponse{}
	_body, _err := client.SafeChangeStartApproveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 封网范围数据查询
//
// @param request - SafeScopeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SafeScopeDataResponse
func (client *Client) SafeScopeDataWithOptions(request *SafeScopeDataRequest, runtime *util.RuntimeOptions) (_result *SafeScopeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.CodeList)) {
		body["CodeList"] = request.CodeList
	}

	if !tea.BoolValue(util.IsUnset(request.Factor)) {
		body["Factor"] = request.Factor
	}

	if !tea.BoolValue(util.IsUnset(request.GroupBy)) {
		body["GroupBy"] = request.GroupBy
	}

	if !tea.BoolValue(util.IsUnset(request.IdList)) {
		body["IdList"] = request.IdList
	}

	if !tea.BoolValue(util.IsUnset(request.Item)) {
		body["Item"] = request.Item
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NeedTotalCount)) {
		body["NeedTotalCount"] = request.NeedTotalCount
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDirection)) {
		body["OrderDirection"] = request.OrderDirection
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.ParentCode)) {
		body["ParentCode"] = request.ParentCode
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		body["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNameEn)) {
		body["RegionNameEn"] = request.RegionNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SearchValue)) {
		body["SearchValue"] = request.SearchValue
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		body["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SafeScopeData"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SafeScopeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 封网范围数据查询
//
// @param request - SafeScopeDataRequest
//
// @return SafeScopeDataResponse
func (client *Client) SafeScopeData(request *SafeScopeDataRequest) (_result *SafeScopeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SafeScopeDataResponse{}
	_body, _err := client.SafeScopeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交审批
//
// @param request - StartApproveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartApproveResponse
func (client *Client) StartApproveWithOptions(request *StartApproveRequest, runtime *util.RuntimeOptions) (_result *StartApproveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		query["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		query["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorEmpId)) {
		query["CreatorEmpId"] = request.CreatorEmpId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraInfo)) {
		query["ExtraInfo"] = request.ExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		query["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOrderId)) {
		query["SourceOrderId"] = request.SourceOrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartApprove"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartApproveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交审批
//
// @param request - StartApproveRequest
//
// @return StartApproveResponse
func (client *Client) StartApprove(request *StartApproveRequest) (_result *StartApproveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartApproveResponse{}
	_body, _err := client.StartApproveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步产品接口
//
// @param request - SyncProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncProductResponse
func (client *Client) SyncProductWithOptions(request *SyncProductRequest, runtime *util.RuntimeOptions) (_result *SyncProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		body["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthSign)) {
		body["AuthSign"] = request.AuthSign
	}

	if !tea.BoolValue(util.IsUnset(request.ReqTimestamp)) {
		body["ReqTimestamp"] = request.ReqTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.SyncProductList)) {
		body["SyncProductList"] = request.SyncProductList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncProduct"),
		Version:     tea.String("2022-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步产品接口
//
// @param request - SyncProductRequest
//
// @return SyncProductResponse
func (client *Client) SyncProduct(request *SyncProductRequest) (_result *SyncProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncProductResponse{}
	_body, _err := client.SyncProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
