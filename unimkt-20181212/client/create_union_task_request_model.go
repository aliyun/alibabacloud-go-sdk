// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUnionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnchorId(v string) *CreateUnionTaskRequest
	GetAnchorId() *string
	SetBrandUserId(v int64) *CreateUnionTaskRequest
	GetBrandUserId() *int64
	SetBrandUserNick(v string) *CreateUnionTaskRequest
	GetBrandUserNick() *string
	SetChannel(v string) *CreateUnionTaskRequest
	GetChannel() *string
	SetChannelId(v string) *CreateUnionTaskRequest
	GetChannelId() *string
	SetChargePloy(v int64) *CreateUnionTaskRequest
	GetChargePloy() *int64
	SetChargeType(v int64) *CreateUnionTaskRequest
	GetChargeType() *int64
	SetClientToken(v string) *CreateUnionTaskRequest
	GetClientToken() *string
	SetContentId(v int64) *CreateUnionTaskRequest
	GetContentId() *int64
	SetContentUrl(v string) *CreateUnionTaskRequest
	GetContentUrl() *string
	SetCustomCreativeType(v string) *CreateUnionTaskRequest
	GetCustomCreativeType() *string
	SetEndTime(v string) *CreateUnionTaskRequest
	GetEndTime() *string
	SetIndustryLabelBagId(v int32) *CreateUnionTaskRequest
	GetIndustryLabelBagId() *int32
	SetMediaIdBlackList(v []*string) *CreateUnionTaskRequest
	GetMediaIdBlackList() []*string
	SetMediaIdWhiteList(v []*string) *CreateUnionTaskRequest
	GetMediaIdWhiteList() []*string
	SetMediaIndustry(v string) *CreateUnionTaskRequest
	GetMediaIndustry() *string
	SetName(v string) *CreateUnionTaskRequest
	GetName() *string
	SetOptimizationSwitch(v int64) *CreateUnionTaskRequest
	GetOptimizationSwitch() *int64
	SetProxyUserId(v int64) *CreateUnionTaskRequest
	GetProxyUserId() *int64
	SetQuota(v int64) *CreateUnionTaskRequest
	GetQuota() *int64
	SetQuotaDay(v int64) *CreateUnionTaskRequest
	GetQuotaDay() *int64
	SetStartTime(v string) *CreateUnionTaskRequest
	GetStartTime() *string
	SetTaskBizType(v string) *CreateUnionTaskRequest
	GetTaskBizType() *string
	SetTaskRuleType(v string) *CreateUnionTaskRequest
	GetTaskRuleType() *string
	SetTaskType(v string) *CreateUnionTaskRequest
	GetTaskType() *string
}

type CreateUnionTaskRequest struct {
	AnchorId           *string   `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	BrandUserId        *int64    `json:"BrandUserId,omitempty" xml:"BrandUserId,omitempty"`
	BrandUserNick      *string   `json:"BrandUserNick,omitempty" xml:"BrandUserNick,omitempty"`
	Channel            *string   `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ChannelId          *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ChargePloy         *int64    `json:"ChargePloy,omitempty" xml:"ChargePloy,omitempty"`
	ChargeType         *int64    `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClientToken        *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContentId          *int64    `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	ContentUrl         *string   `json:"ContentUrl,omitempty" xml:"ContentUrl,omitempty"`
	CustomCreativeType *string   `json:"CustomCreativeType,omitempty" xml:"CustomCreativeType,omitempty"`
	EndTime            *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IndustryLabelBagId *int32    `json:"IndustryLabelBagId,omitempty" xml:"IndustryLabelBagId,omitempty"`
	MediaIdBlackList   []*string `json:"MediaIdBlackList,omitempty" xml:"MediaIdBlackList,omitempty" type:"Repeated"`
	MediaIdWhiteList   []*string `json:"MediaIdWhiteList,omitempty" xml:"MediaIdWhiteList,omitempty" type:"Repeated"`
	MediaIndustry      *string   `json:"MediaIndustry,omitempty" xml:"MediaIndustry,omitempty"`
	Name               *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OptimizationSwitch *int64    `json:"OptimizationSwitch,omitempty" xml:"OptimizationSwitch,omitempty"`
	ProxyUserId        *int64    `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	Quota              *int64    `json:"Quota,omitempty" xml:"Quota,omitempty"`
	QuotaDay           *int64    `json:"QuotaDay,omitempty" xml:"QuotaDay,omitempty"`
	StartTime          *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskBizType        *string   `json:"TaskBizType,omitempty" xml:"TaskBizType,omitempty"`
	TaskRuleType       *string   `json:"TaskRuleType,omitempty" xml:"TaskRuleType,omitempty"`
	TaskType           *string   `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateUnionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUnionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateUnionTaskRequest) GetAnchorId() *string {
	return s.AnchorId
}

func (s *CreateUnionTaskRequest) GetBrandUserId() *int64 {
	return s.BrandUserId
}

func (s *CreateUnionTaskRequest) GetBrandUserNick() *string {
	return s.BrandUserNick
}

func (s *CreateUnionTaskRequest) GetChannel() *string {
	return s.Channel
}

func (s *CreateUnionTaskRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *CreateUnionTaskRequest) GetChargePloy() *int64 {
	return s.ChargePloy
}

func (s *CreateUnionTaskRequest) GetChargeType() *int64 {
	return s.ChargeType
}

func (s *CreateUnionTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateUnionTaskRequest) GetContentId() *int64 {
	return s.ContentId
}

func (s *CreateUnionTaskRequest) GetContentUrl() *string {
	return s.ContentUrl
}

func (s *CreateUnionTaskRequest) GetCustomCreativeType() *string {
	return s.CustomCreativeType
}

func (s *CreateUnionTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateUnionTaskRequest) GetIndustryLabelBagId() *int32 {
	return s.IndustryLabelBagId
}

func (s *CreateUnionTaskRequest) GetMediaIdBlackList() []*string {
	return s.MediaIdBlackList
}

func (s *CreateUnionTaskRequest) GetMediaIdWhiteList() []*string {
	return s.MediaIdWhiteList
}

func (s *CreateUnionTaskRequest) GetMediaIndustry() *string {
	return s.MediaIndustry
}

func (s *CreateUnionTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateUnionTaskRequest) GetOptimizationSwitch() *int64 {
	return s.OptimizationSwitch
}

func (s *CreateUnionTaskRequest) GetProxyUserId() *int64 {
	return s.ProxyUserId
}

func (s *CreateUnionTaskRequest) GetQuota() *int64 {
	return s.Quota
}

func (s *CreateUnionTaskRequest) GetQuotaDay() *int64 {
	return s.QuotaDay
}

func (s *CreateUnionTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateUnionTaskRequest) GetTaskBizType() *string {
	return s.TaskBizType
}

func (s *CreateUnionTaskRequest) GetTaskRuleType() *string {
	return s.TaskRuleType
}

func (s *CreateUnionTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateUnionTaskRequest) SetAnchorId(v string) *CreateUnionTaskRequest {
	s.AnchorId = &v
	return s
}

func (s *CreateUnionTaskRequest) SetBrandUserId(v int64) *CreateUnionTaskRequest {
	s.BrandUserId = &v
	return s
}

func (s *CreateUnionTaskRequest) SetBrandUserNick(v string) *CreateUnionTaskRequest {
	s.BrandUserNick = &v
	return s
}

func (s *CreateUnionTaskRequest) SetChannel(v string) *CreateUnionTaskRequest {
	s.Channel = &v
	return s
}

func (s *CreateUnionTaskRequest) SetChannelId(v string) *CreateUnionTaskRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateUnionTaskRequest) SetChargePloy(v int64) *CreateUnionTaskRequest {
	s.ChargePloy = &v
	return s
}

func (s *CreateUnionTaskRequest) SetChargeType(v int64) *CreateUnionTaskRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateUnionTaskRequest) SetClientToken(v string) *CreateUnionTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateUnionTaskRequest) SetContentId(v int64) *CreateUnionTaskRequest {
	s.ContentId = &v
	return s
}

func (s *CreateUnionTaskRequest) SetContentUrl(v string) *CreateUnionTaskRequest {
	s.ContentUrl = &v
	return s
}

func (s *CreateUnionTaskRequest) SetCustomCreativeType(v string) *CreateUnionTaskRequest {
	s.CustomCreativeType = &v
	return s
}

func (s *CreateUnionTaskRequest) SetEndTime(v string) *CreateUnionTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateUnionTaskRequest) SetIndustryLabelBagId(v int32) *CreateUnionTaskRequest {
	s.IndustryLabelBagId = &v
	return s
}

func (s *CreateUnionTaskRequest) SetMediaIdBlackList(v []*string) *CreateUnionTaskRequest {
	s.MediaIdBlackList = v
	return s
}

func (s *CreateUnionTaskRequest) SetMediaIdWhiteList(v []*string) *CreateUnionTaskRequest {
	s.MediaIdWhiteList = v
	return s
}

func (s *CreateUnionTaskRequest) SetMediaIndustry(v string) *CreateUnionTaskRequest {
	s.MediaIndustry = &v
	return s
}

func (s *CreateUnionTaskRequest) SetName(v string) *CreateUnionTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateUnionTaskRequest) SetOptimizationSwitch(v int64) *CreateUnionTaskRequest {
	s.OptimizationSwitch = &v
	return s
}

func (s *CreateUnionTaskRequest) SetProxyUserId(v int64) *CreateUnionTaskRequest {
	s.ProxyUserId = &v
	return s
}

func (s *CreateUnionTaskRequest) SetQuota(v int64) *CreateUnionTaskRequest {
	s.Quota = &v
	return s
}

func (s *CreateUnionTaskRequest) SetQuotaDay(v int64) *CreateUnionTaskRequest {
	s.QuotaDay = &v
	return s
}

func (s *CreateUnionTaskRequest) SetStartTime(v string) *CreateUnionTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateUnionTaskRequest) SetTaskBizType(v string) *CreateUnionTaskRequest {
	s.TaskBizType = &v
	return s
}

func (s *CreateUnionTaskRequest) SetTaskRuleType(v string) *CreateUnionTaskRequest {
	s.TaskRuleType = &v
	return s
}

func (s *CreateUnionTaskRequest) SetTaskType(v string) *CreateUnionTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateUnionTaskRequest) Validate() error {
	return dara.Validate(s)
}
