// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTodoTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizTag(v string) *GetTodoTaskResponseBody
	GetBizTag() *string
	SetCardTypeId(v string) *GetTodoTaskResponseBody
	GetCardTypeId() *string
	SetCreatedTime(v int64) *GetTodoTaskResponseBody
	GetCreatedTime() *int64
	SetCreatorId(v string) *GetTodoTaskResponseBody
	GetCreatorId() *string
	SetDescription(v string) *GetTodoTaskResponseBody
	GetDescription() *string
	SetDetailUrl(v *GetTodoTaskResponseBodyDetailUrl) *GetTodoTaskResponseBody
	GetDetailUrl() *GetTodoTaskResponseBodyDetailUrl
	SetDone(v bool) *GetTodoTaskResponseBody
	GetDone() *bool
	SetDueTime(v int64) *GetTodoTaskResponseBody
	GetDueTime() *int64
	SetExecutorIds(v []*string) *GetTodoTaskResponseBody
	GetExecutorIds() []*string
	SetFinishTime(v int64) *GetTodoTaskResponseBody
	GetFinishTime() *int64
	SetId(v string) *GetTodoTaskResponseBody
	GetId() *string
	SetIsOnlyShowExecutor(v bool) *GetTodoTaskResponseBody
	GetIsOnlyShowExecutor() *bool
	SetModifiedTime(v int64) *GetTodoTaskResponseBody
	GetModifiedTime() *int64
	SetModifierId(v string) *GetTodoTaskResponseBody
	GetModifierId() *string
	SetParticipantIds(v []*string) *GetTodoTaskResponseBody
	GetParticipantIds() []*string
	SetPriority(v int32) *GetTodoTaskResponseBody
	GetPriority() *int32
	SetRequestId(v string) *GetTodoTaskResponseBody
	GetRequestId() *string
	SetSource(v string) *GetTodoTaskResponseBody
	GetSource() *string
	SetSourceId(v string) *GetTodoTaskResponseBody
	GetSourceId() *string
	SetStartTime(v int64) *GetTodoTaskResponseBody
	GetStartTime() *int64
	SetSubject(v string) *GetTodoTaskResponseBody
	GetSubject() *string
	SetTenantId(v string) *GetTodoTaskResponseBody
	GetTenantId() *string
	SetTenantType(v string) *GetTodoTaskResponseBody
	GetTenantType() *string
	SetVendorRequestId(v string) *GetTodoTaskResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetTodoTaskResponseBody
	GetVendorType() *string
}

type GetTodoTaskResponseBody struct {
	// example:
	//
	// isv_dingtalkTodo
	BizTag *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	// example:
	//
	// 此参数禁止发布
	CardTypeId *string `json:"cardTypeId,omitempty" xml:"cardTypeId,omitempty"`
	// example:
	//
	// 1617675000000
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// 012345
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// 应用可以调用该接口获取钉钉待办任务详情信息及状态。
	Description *string                           `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl   *GetTodoTaskResponseBodyDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	// example:
	//
	// true
	Done *bool `json:"done,omitempty" xml:"done,omitempty"`
	// example:
	//
	// 1617675000000
	DueTime     *int64    `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1617675000000
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// example:
	//
	// OPJpwtxxxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IsOnlyShowExecutor *bool `json:"isOnlyShowExecutor,omitempty" xml:"isOnlyShowExecutor,omitempty"`
	// example:
	//
	// 1617675000000
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// 012345
	ModifierId     *string   `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	ParticipantIds []*string `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// example:
	//
	// PUoiinWIxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// isv_dingtalkTodo
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// isv_dingxxx
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// 1617675000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 接入钉钉待办
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// example:
	//
	// orgId1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// example:
	//
	// org
	TenantType *string `json:"tenantType,omitempty" xml:"tenantType,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetTodoTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTodoTaskResponseBody) GetBizTag() *string {
	return s.BizTag
}

func (s *GetTodoTaskResponseBody) GetCardTypeId() *string {
	return s.CardTypeId
}

func (s *GetTodoTaskResponseBody) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetTodoTaskResponseBody) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetTodoTaskResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetTodoTaskResponseBody) GetDetailUrl() *GetTodoTaskResponseBodyDetailUrl {
	return s.DetailUrl
}

func (s *GetTodoTaskResponseBody) GetDone() *bool {
	return s.Done
}

func (s *GetTodoTaskResponseBody) GetDueTime() *int64 {
	return s.DueTime
}

func (s *GetTodoTaskResponseBody) GetExecutorIds() []*string {
	return s.ExecutorIds
}

func (s *GetTodoTaskResponseBody) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetTodoTaskResponseBody) GetId() *string {
	return s.Id
}

func (s *GetTodoTaskResponseBody) GetIsOnlyShowExecutor() *bool {
	return s.IsOnlyShowExecutor
}

func (s *GetTodoTaskResponseBody) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetTodoTaskResponseBody) GetModifierId() *string {
	return s.ModifierId
}

func (s *GetTodoTaskResponseBody) GetParticipantIds() []*string {
	return s.ParticipantIds
}

func (s *GetTodoTaskResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *GetTodoTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTodoTaskResponseBody) GetSource() *string {
	return s.Source
}

func (s *GetTodoTaskResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *GetTodoTaskResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetTodoTaskResponseBody) GetSubject() *string {
	return s.Subject
}

func (s *GetTodoTaskResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *GetTodoTaskResponseBody) GetTenantType() *string {
	return s.TenantType
}

func (s *GetTodoTaskResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetTodoTaskResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetTodoTaskResponseBody) SetBizTag(v string) *GetTodoTaskResponseBody {
	s.BizTag = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetCardTypeId(v string) *GetTodoTaskResponseBody {
	s.CardTypeId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetCreatedTime(v int64) *GetTodoTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetCreatorId(v string) *GetTodoTaskResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetDescription(v string) *GetTodoTaskResponseBody {
	s.Description = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetDetailUrl(v *GetTodoTaskResponseBodyDetailUrl) *GetTodoTaskResponseBody {
	s.DetailUrl = v
	return s
}

func (s *GetTodoTaskResponseBody) SetDone(v bool) *GetTodoTaskResponseBody {
	s.Done = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetDueTime(v int64) *GetTodoTaskResponseBody {
	s.DueTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetExecutorIds(v []*string) *GetTodoTaskResponseBody {
	s.ExecutorIds = v
	return s
}

func (s *GetTodoTaskResponseBody) SetFinishTime(v int64) *GetTodoTaskResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetId(v string) *GetTodoTaskResponseBody {
	s.Id = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetIsOnlyShowExecutor(v bool) *GetTodoTaskResponseBody {
	s.IsOnlyShowExecutor = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetModifiedTime(v int64) *GetTodoTaskResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetModifierId(v string) *GetTodoTaskResponseBody {
	s.ModifierId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetParticipantIds(v []*string) *GetTodoTaskResponseBody {
	s.ParticipantIds = v
	return s
}

func (s *GetTodoTaskResponseBody) SetPriority(v int32) *GetTodoTaskResponseBody {
	s.Priority = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetRequestId(v string) *GetTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetSource(v string) *GetTodoTaskResponseBody {
	s.Source = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetSourceId(v string) *GetTodoTaskResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetStartTime(v int64) *GetTodoTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetSubject(v string) *GetTodoTaskResponseBody {
	s.Subject = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetTenantId(v string) *GetTodoTaskResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetTenantType(v string) *GetTodoTaskResponseBody {
	s.TenantType = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetVendorRequestId(v string) *GetTodoTaskResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetVendorType(v string) *GetTodoTaskResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetTodoTaskResponseBody) Validate() error {
	if s.DetailUrl != nil {
		if err := s.DetailUrl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTodoTaskResponseBodyDetailUrl struct {
	// example:
	//
	// dingtalk://dingtalkclient/action/open_mini_app?miniAppId={0}&ddMode=push&page=pages%2ftask-detail%2ftask-detail%3ftaskId%3d{1}
	AppUrl *string `json:"AppUrl,omitempty" xml:"AppUrl,omitempty"`
	// example:
	//
	// https://todo.dingtalk.com/ding-portal/detail/task/{0}
	PcUrl *string `json:"PcUrl,omitempty" xml:"PcUrl,omitempty"`
}

func (s GetTodoTaskResponseBodyDetailUrl) String() string {
	return dara.Prettify(s)
}

func (s GetTodoTaskResponseBodyDetailUrl) GoString() string {
	return s.String()
}

func (s *GetTodoTaskResponseBodyDetailUrl) GetAppUrl() *string {
	return s.AppUrl
}

func (s *GetTodoTaskResponseBodyDetailUrl) GetPcUrl() *string {
	return s.PcUrl
}

func (s *GetTodoTaskResponseBodyDetailUrl) SetAppUrl(v string) *GetTodoTaskResponseBodyDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *GetTodoTaskResponseBodyDetailUrl) SetPcUrl(v string) *GetTodoTaskResponseBodyDetailUrl {
	s.PcUrl = &v
	return s
}

func (s *GetTodoTaskResponseBodyDetailUrl) Validate() error {
	return dara.Validate(s)
}
