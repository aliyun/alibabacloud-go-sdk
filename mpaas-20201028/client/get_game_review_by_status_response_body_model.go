// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGameReviewByStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetGameReviewByStatusResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetGameReviewByStatusResponseBody
	GetResultCode() *string
	SetResultContent(v *GetGameReviewByStatusResponseBodyResultContent) *GetGameReviewByStatusResponseBody
	GetResultContent() *GetGameReviewByStatusResponseBodyResultContent
	SetResultMessage(v string) *GetGameReviewByStatusResponseBody
	GetResultMessage() *string
}

type GetGameReviewByStatusResponseBody struct {
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                         `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *GetGameReviewByStatusResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                         `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetGameReviewByStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGameReviewByStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetGameReviewByStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGameReviewByStatusResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetGameReviewByStatusResponseBody) GetResultContent() *GetGameReviewByStatusResponseBodyResultContent {
	return s.ResultContent
}

func (s *GetGameReviewByStatusResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *GetGameReviewByStatusResponseBody) SetRequestId(v string) *GetGameReviewByStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGameReviewByStatusResponseBody) SetResultCode(v string) *GetGameReviewByStatusResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetGameReviewByStatusResponseBody) SetResultContent(v *GetGameReviewByStatusResponseBodyResultContent) *GetGameReviewByStatusResponseBody {
	s.ResultContent = v
	return s
}

func (s *GetGameReviewByStatusResponseBody) SetResultMessage(v string) *GetGameReviewByStatusResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetGameReviewByStatusResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGameReviewByStatusResponseBodyResultContent struct {
	Content    []*GetGameReviewByStatusResponseBodyResultContentContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	ErrorCode  *string                                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	PageNum    *int32                                                   `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg  *string                                                  `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success    *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetGameReviewByStatusResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s GetGameReviewByStatusResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *GetGameReviewByStatusResponseBodyResultContent) GetContent() []*GetGameReviewByStatusResponseBodyResultContentContent {
	return s.Content
}

func (s *GetGameReviewByStatusResponseBodyResultContent) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetGameReviewByStatusResponseBodyResultContent) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetGameReviewByStatusResponseBodyResultContent) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetGameReviewByStatusResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGameReviewByStatusResponseBodyResultContent) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *GetGameReviewByStatusResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *GetGameReviewByStatusResponseBodyResultContent) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetGameReviewByStatusResponseBodyResultContent) SetContent(v []*GetGameReviewByStatusResponseBodyResultContentContent) *GetGameReviewByStatusResponseBodyResultContent {
	s.Content = v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContent) SetErrorCode(v string) *GetGameReviewByStatusResponseBodyResultContent {
	s.ErrorCode = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContent) SetPageNum(v int32) *GetGameReviewByStatusResponseBodyResultContent {
	s.PageNum = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContent) SetPageSize(v int32) *GetGameReviewByStatusResponseBodyResultContent {
	s.PageSize = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContent) SetRequestId(v string) *GetGameReviewByStatusResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContent) SetResultMsg(v string) *GetGameReviewByStatusResponseBodyResultContent {
	s.ResultMsg = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContent) SetSuccess(v bool) *GetGameReviewByStatusResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContent) SetTotalCount(v int64) *GetGameReviewByStatusResponseBodyResultContent {
	s.TotalCount = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContent) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGameReviewByStatusResponseBodyResultContentContent struct {
	Appendix            *string                                                                `json:"Appendix,omitempty" xml:"Appendix,omitempty"`
	AutoOnline          *bool                                                                  `json:"AutoOnline,omitempty" xml:"AutoOnline,omitempty"`
	Creator             *string                                                                `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Extension           *string                                                                `json:"Extension,omitempty" xml:"Extension,omitempty"`
	FlowId              *int64                                                                 `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	GmtCreate           *string                                                                `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified         *string                                                                `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Icons               []*string                                                              `json:"Icons,omitempty" xml:"Icons,omitempty" type:"Repeated"`
	Id                  *int64                                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	ReviewChannel       *string                                                                `json:"ReviewChannel,omitempty" xml:"ReviewChannel,omitempty"`
	ReviewIntegerStatus *int32                                                                 `json:"ReviewIntegerStatus,omitempty" xml:"ReviewIntegerStatus,omitempty"`
	ReviewProgress      []*GetGameReviewByStatusResponseBodyResultContentContentReviewProgress `json:"ReviewProgress,omitempty" xml:"ReviewProgress,omitempty" type:"Repeated"`
	ReviewStatus        *string                                                                `json:"ReviewStatus,omitempty" xml:"ReviewStatus,omitempty"`
	ReviewTarget        *int64                                                                 `json:"ReviewTarget,omitempty" xml:"ReviewTarget,omitempty"`
	TargetDetail        *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail     `json:"TargetDetail,omitempty" xml:"TargetDetail,omitempty" type:"Struct"`
	TargetType          *string                                                                `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetGameReviewByStatusResponseBodyResultContentContent) String() string {
	return dara.Prettify(s)
}

func (s GetGameReviewByStatusResponseBodyResultContentContent) GoString() string {
	return s.String()
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetAppendix() *string {
	return s.Appendix
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetAutoOnline() *bool {
	return s.AutoOnline
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetCreator() *string {
	return s.Creator
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetExtension() *string {
	return s.Extension
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetFlowId() *int64 {
	return s.FlowId
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetIcons() []*string {
	return s.Icons
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetId() *int64 {
	return s.Id
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetReviewChannel() *string {
	return s.ReviewChannel
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetReviewIntegerStatus() *int32 {
	return s.ReviewIntegerStatus
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetReviewProgress() []*GetGameReviewByStatusResponseBodyResultContentContentReviewProgress {
	return s.ReviewProgress
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetReviewStatus() *string {
	return s.ReviewStatus
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetReviewTarget() *int64 {
	return s.ReviewTarget
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetTargetDetail() *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	return s.TargetDetail
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) GetTargetType() *string {
	return s.TargetType
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetAppendix(v string) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.Appendix = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetAutoOnline(v bool) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.AutoOnline = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetCreator(v string) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.Creator = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetExtension(v string) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.Extension = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetFlowId(v int64) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.FlowId = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetGmtCreate(v string) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.GmtCreate = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetGmtModified(v string) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.GmtModified = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetIcons(v []*string) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.Icons = v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetId(v int64) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.Id = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetReviewChannel(v string) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.ReviewChannel = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetReviewIntegerStatus(v int32) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.ReviewIntegerStatus = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetReviewProgress(v []*GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.ReviewProgress = v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetReviewStatus(v string) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.ReviewStatus = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetReviewTarget(v int64) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.ReviewTarget = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetTargetDetail(v *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.TargetDetail = v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) SetTargetType(v string) *GetGameReviewByStatusResponseBodyResultContentContent {
	s.TargetType = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContent) Validate() error {
	if s.ReviewProgress != nil {
		for _, item := range s.ReviewProgress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TargetDetail != nil {
		if err := s.TargetDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGameReviewByStatusResponseBodyResultContentContentReviewProgress struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FlowNodeId   *int64  `json:"FlowNodeId,omitempty" xml:"FlowNodeId,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	NodeIndex    *int32  `json:"NodeIndex,omitempty" xml:"NodeIndex,omitempty"`
	NodeName     *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeStatus   *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	NodeStatusId *int32  `json:"NodeStatusId,omitempty" xml:"NodeStatusId,omitempty"`
	Operator     *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RoleId       *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) String() string {
	return dara.Prettify(s)
}

func (s GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) GoString() string {
	return s.String()
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) GetDescription() *string {
	return s.Description
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) GetFlowNodeId() *int64 {
	return s.FlowNodeId
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) GetNodeIndex() *int32 {
	return s.NodeIndex
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) GetNodeName() *string {
	return s.NodeName
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) GetNodeStatus() *string {
	return s.NodeStatus
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) GetNodeStatusId() *int32 {
	return s.NodeStatusId
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) GetOperator() *string {
	return s.Operator
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) GetRemark() *string {
	return s.Remark
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) GetRoleId() *int64 {
	return s.RoleId
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) SetDescription(v string) *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress {
	s.Description = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) SetFlowNodeId(v int64) *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress {
	s.FlowNodeId = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) SetGmtModified(v string) *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress {
	s.GmtModified = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) SetNodeIndex(v int32) *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress {
	s.NodeIndex = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) SetNodeName(v string) *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress {
	s.NodeName = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) SetNodeStatus(v string) *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress {
	s.NodeStatus = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) SetNodeStatusId(v int32) *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress {
	s.NodeStatusId = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) SetOperator(v string) *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress {
	s.Operator = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) SetRemark(v string) *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress {
	s.Remark = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) SetRoleId(v int64) *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress {
	s.RoleId = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentReviewProgress) Validate() error {
	return dara.Validate(s)
}

type GetGameReviewByStatusResponseBodyResultContentContentTargetDetail struct {
	AutoOnline        *bool   `json:"AutoOnline,omitempty" xml:"AutoOnline,omitempty"`
	Category          *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Detail            *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	GameMaker         *string `json:"GameMaker,omitempty" xml:"GameMaker,omitempty"`
	IconUrl           *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	Introduction      *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	MiniProgramCode   *string `json:"MiniProgramCode,omitempty" xml:"MiniProgramCode,omitempty"`
	MiniProgramInfoId *int64  `json:"MiniProgramInfoId,omitempty" xml:"MiniProgramInfoId,omitempty"`
	MiniProgramName   *string `json:"MiniProgramName,omitempty" xml:"MiniProgramName,omitempty"`
	MiniResourceId    *int64  `json:"MiniResourceId,omitempty" xml:"MiniResourceId,omitempty"`
	PublishStatus     *int32  `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	QrCodeUrl         *string `json:"QrCodeUrl,omitempty" xml:"QrCodeUrl,omitempty"`
	ReviewTargetType  *string `json:"ReviewTargetType,omitempty" xml:"ReviewTargetType,omitempty"`
	SubType           *int32  `json:"SubType,omitempty" xml:"SubType,omitempty"`
	TargetId          *int64  `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	Version           *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) String() string {
	return dara.Prettify(s)
}

func (s GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GoString() string {
	return s.String()
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetAutoOnline() *bool {
	return s.AutoOnline
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetCategory() *string {
	return s.Category
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetDetail() *string {
	return s.Detail
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetGameMaker() *string {
	return s.GameMaker
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetIconUrl() *string {
	return s.IconUrl
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetIntroduction() *string {
	return s.Introduction
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetMiniProgramCode() *string {
	return s.MiniProgramCode
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetMiniProgramInfoId() *int64 {
	return s.MiniProgramInfoId
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetMiniProgramName() *string {
	return s.MiniProgramName
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetMiniResourceId() *int64 {
	return s.MiniResourceId
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetPublishStatus() *int32 {
	return s.PublishStatus
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetQrCodeUrl() *string {
	return s.QrCodeUrl
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetReviewTargetType() *string {
	return s.ReviewTargetType
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetSubType() *int32 {
	return s.SubType
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetTargetId() *int64 {
	return s.TargetId
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) GetVersion() *string {
	return s.Version
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetAutoOnline(v bool) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.AutoOnline = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetCategory(v string) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.Category = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetDetail(v string) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.Detail = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetGameMaker(v string) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.GameMaker = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetIconUrl(v string) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.IconUrl = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetIntroduction(v string) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.Introduction = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetMiniProgramCode(v string) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.MiniProgramCode = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetMiniProgramInfoId(v int64) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.MiniProgramInfoId = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetMiniProgramName(v string) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.MiniProgramName = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetMiniResourceId(v int64) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.MiniResourceId = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetPublishStatus(v int32) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.PublishStatus = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetQrCodeUrl(v string) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.QrCodeUrl = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetReviewTargetType(v string) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.ReviewTargetType = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetSubType(v int32) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.SubType = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetTargetId(v int64) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.TargetId = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) SetVersion(v string) *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail {
	s.Version = &v
	return s
}

func (s *GetGameReviewByStatusResponseBodyResultContentContentTargetDetail) Validate() error {
	return dara.Validate(s)
}
