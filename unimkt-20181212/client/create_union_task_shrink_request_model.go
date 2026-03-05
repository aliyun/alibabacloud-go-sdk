// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUnionTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnchorId(v string) *CreateUnionTaskShrinkRequest
	GetAnchorId() *string
	SetBrandUserId(v int64) *CreateUnionTaskShrinkRequest
	GetBrandUserId() *int64
	SetBrandUserNick(v string) *CreateUnionTaskShrinkRequest
	GetBrandUserNick() *string
	SetChannel(v string) *CreateUnionTaskShrinkRequest
	GetChannel() *string
	SetChannelId(v string) *CreateUnionTaskShrinkRequest
	GetChannelId() *string
	SetChargePloy(v int64) *CreateUnionTaskShrinkRequest
	GetChargePloy() *int64
	SetChargeType(v int64) *CreateUnionTaskShrinkRequest
	GetChargeType() *int64
	SetClientToken(v string) *CreateUnionTaskShrinkRequest
	GetClientToken() *string
	SetContentId(v int64) *CreateUnionTaskShrinkRequest
	GetContentId() *int64
	SetContentUrl(v string) *CreateUnionTaskShrinkRequest
	GetContentUrl() *string
	SetCustomCreativeType(v string) *CreateUnionTaskShrinkRequest
	GetCustomCreativeType() *string
	SetEndTime(v string) *CreateUnionTaskShrinkRequest
	GetEndTime() *string
	SetIndustryLabelBagId(v int32) *CreateUnionTaskShrinkRequest
	GetIndustryLabelBagId() *int32
	SetMediaIdBlackListShrink(v string) *CreateUnionTaskShrinkRequest
	GetMediaIdBlackListShrink() *string
	SetMediaIdWhiteListShrink(v string) *CreateUnionTaskShrinkRequest
	GetMediaIdWhiteListShrink() *string
	SetMediaIndustry(v string) *CreateUnionTaskShrinkRequest
	GetMediaIndustry() *string
	SetName(v string) *CreateUnionTaskShrinkRequest
	GetName() *string
	SetOptimizationSwitch(v int64) *CreateUnionTaskShrinkRequest
	GetOptimizationSwitch() *int64
	SetProxyUserId(v int64) *CreateUnionTaskShrinkRequest
	GetProxyUserId() *int64
	SetQuota(v int64) *CreateUnionTaskShrinkRequest
	GetQuota() *int64
	SetQuotaDay(v int64) *CreateUnionTaskShrinkRequest
	GetQuotaDay() *int64
	SetStartTime(v string) *CreateUnionTaskShrinkRequest
	GetStartTime() *string
	SetTaskBizType(v string) *CreateUnionTaskShrinkRequest
	GetTaskBizType() *string
	SetTaskRuleType(v string) *CreateUnionTaskShrinkRequest
	GetTaskRuleType() *string
	SetTaskType(v string) *CreateUnionTaskShrinkRequest
	GetTaskType() *string
}

type CreateUnionTaskShrinkRequest struct {
	AnchorId               *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	BrandUserId            *int64  `json:"BrandUserId,omitempty" xml:"BrandUserId,omitempty"`
	BrandUserNick          *string `json:"BrandUserNick,omitempty" xml:"BrandUserNick,omitempty"`
	Channel                *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ChannelId              *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ChargePloy             *int64  `json:"ChargePloy,omitempty" xml:"ChargePloy,omitempty"`
	ChargeType             *int64  `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContentId              *int64  `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	ContentUrl             *string `json:"ContentUrl,omitempty" xml:"ContentUrl,omitempty"`
	CustomCreativeType     *string `json:"CustomCreativeType,omitempty" xml:"CustomCreativeType,omitempty"`
	EndTime                *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IndustryLabelBagId     *int32  `json:"IndustryLabelBagId,omitempty" xml:"IndustryLabelBagId,omitempty"`
	MediaIdBlackListShrink *string `json:"MediaIdBlackList,omitempty" xml:"MediaIdBlackList,omitempty"`
	MediaIdWhiteListShrink *string `json:"MediaIdWhiteList,omitempty" xml:"MediaIdWhiteList,omitempty"`
	MediaIndustry          *string `json:"MediaIndustry,omitempty" xml:"MediaIndustry,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OptimizationSwitch     *int64  `json:"OptimizationSwitch,omitempty" xml:"OptimizationSwitch,omitempty"`
	ProxyUserId            *int64  `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	Quota                  *int64  `json:"Quota,omitempty" xml:"Quota,omitempty"`
	QuotaDay               *int64  `json:"QuotaDay,omitempty" xml:"QuotaDay,omitempty"`
	StartTime              *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskBizType            *string `json:"TaskBizType,omitempty" xml:"TaskBizType,omitempty"`
	TaskRuleType           *string `json:"TaskRuleType,omitempty" xml:"TaskRuleType,omitempty"`
	TaskType               *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateUnionTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUnionTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUnionTaskShrinkRequest) GetAnchorId() *string {
	return s.AnchorId
}

func (s *CreateUnionTaskShrinkRequest) GetBrandUserId() *int64 {
	return s.BrandUserId
}

func (s *CreateUnionTaskShrinkRequest) GetBrandUserNick() *string {
	return s.BrandUserNick
}

func (s *CreateUnionTaskShrinkRequest) GetChannel() *string {
	return s.Channel
}

func (s *CreateUnionTaskShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *CreateUnionTaskShrinkRequest) GetChargePloy() *int64 {
	return s.ChargePloy
}

func (s *CreateUnionTaskShrinkRequest) GetChargeType() *int64 {
	return s.ChargeType
}

func (s *CreateUnionTaskShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateUnionTaskShrinkRequest) GetContentId() *int64 {
	return s.ContentId
}

func (s *CreateUnionTaskShrinkRequest) GetContentUrl() *string {
	return s.ContentUrl
}

func (s *CreateUnionTaskShrinkRequest) GetCustomCreativeType() *string {
	return s.CustomCreativeType
}

func (s *CreateUnionTaskShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateUnionTaskShrinkRequest) GetIndustryLabelBagId() *int32 {
	return s.IndustryLabelBagId
}

func (s *CreateUnionTaskShrinkRequest) GetMediaIdBlackListShrink() *string {
	return s.MediaIdBlackListShrink
}

func (s *CreateUnionTaskShrinkRequest) GetMediaIdWhiteListShrink() *string {
	return s.MediaIdWhiteListShrink
}

func (s *CreateUnionTaskShrinkRequest) GetMediaIndustry() *string {
	return s.MediaIndustry
}

func (s *CreateUnionTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateUnionTaskShrinkRequest) GetOptimizationSwitch() *int64 {
	return s.OptimizationSwitch
}

func (s *CreateUnionTaskShrinkRequest) GetProxyUserId() *int64 {
	return s.ProxyUserId
}

func (s *CreateUnionTaskShrinkRequest) GetQuota() *int64 {
	return s.Quota
}

func (s *CreateUnionTaskShrinkRequest) GetQuotaDay() *int64 {
	return s.QuotaDay
}

func (s *CreateUnionTaskShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateUnionTaskShrinkRequest) GetTaskBizType() *string {
	return s.TaskBizType
}

func (s *CreateUnionTaskShrinkRequest) GetTaskRuleType() *string {
	return s.TaskRuleType
}

func (s *CreateUnionTaskShrinkRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateUnionTaskShrinkRequest) SetAnchorId(v string) *CreateUnionTaskShrinkRequest {
	s.AnchorId = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetBrandUserId(v int64) *CreateUnionTaskShrinkRequest {
	s.BrandUserId = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetBrandUserNick(v string) *CreateUnionTaskShrinkRequest {
	s.BrandUserNick = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetChannel(v string) *CreateUnionTaskShrinkRequest {
	s.Channel = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetChannelId(v string) *CreateUnionTaskShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetChargePloy(v int64) *CreateUnionTaskShrinkRequest {
	s.ChargePloy = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetChargeType(v int64) *CreateUnionTaskShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetClientToken(v string) *CreateUnionTaskShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetContentId(v int64) *CreateUnionTaskShrinkRequest {
	s.ContentId = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetContentUrl(v string) *CreateUnionTaskShrinkRequest {
	s.ContentUrl = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetCustomCreativeType(v string) *CreateUnionTaskShrinkRequest {
	s.CustomCreativeType = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetEndTime(v string) *CreateUnionTaskShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetIndustryLabelBagId(v int32) *CreateUnionTaskShrinkRequest {
	s.IndustryLabelBagId = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetMediaIdBlackListShrink(v string) *CreateUnionTaskShrinkRequest {
	s.MediaIdBlackListShrink = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetMediaIdWhiteListShrink(v string) *CreateUnionTaskShrinkRequest {
	s.MediaIdWhiteListShrink = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetMediaIndustry(v string) *CreateUnionTaskShrinkRequest {
	s.MediaIndustry = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetName(v string) *CreateUnionTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetOptimizationSwitch(v int64) *CreateUnionTaskShrinkRequest {
	s.OptimizationSwitch = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetProxyUserId(v int64) *CreateUnionTaskShrinkRequest {
	s.ProxyUserId = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetQuota(v int64) *CreateUnionTaskShrinkRequest {
	s.Quota = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetQuotaDay(v int64) *CreateUnionTaskShrinkRequest {
	s.QuotaDay = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetStartTime(v string) *CreateUnionTaskShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetTaskBizType(v string) *CreateUnionTaskShrinkRequest {
	s.TaskBizType = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetTaskRuleType(v string) *CreateUnionTaskShrinkRequest {
	s.TaskRuleType = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) SetTaskType(v string) *CreateUnionTaskShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *CreateUnionTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
