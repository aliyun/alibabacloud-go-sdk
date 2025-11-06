// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListInstanceAlarmResponseBody
	GetCode() *int32
	SetData(v *ListInstanceAlarmResponseBodyData) *ListInstanceAlarmResponseBody
	GetData() *ListInstanceAlarmResponseBodyData
	SetMessage(v string) *ListInstanceAlarmResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstanceAlarmResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstanceAlarmResponseBody
	GetSuccess() *bool
}

type ListInstanceAlarmResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListInstanceAlarmResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstanceAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceAlarmResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListInstanceAlarmResponseBody) GetData() *ListInstanceAlarmResponseBodyData {
	return s.Data
}

func (s *ListInstanceAlarmResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceAlarmResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceAlarmResponseBody) SetCode(v int32) *ListInstanceAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceAlarmResponseBody) SetData(v *ListInstanceAlarmResponseBodyData) *ListInstanceAlarmResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceAlarmResponseBody) SetMessage(v string) *ListInstanceAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceAlarmResponseBody) SetRequestId(v string) *ListInstanceAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceAlarmResponseBody) SetSuccess(v bool) *ListInstanceAlarmResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceAlarmResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceAlarmResponseBodyData struct {
	CurrentPage *int32                                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      *ListInstanceAlarmResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Struct"`
}

func (s ListInstanceAlarmResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAlarmResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceAlarmResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListInstanceAlarmResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceAlarmResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstanceAlarmResponseBodyData) GetVoList() *ListInstanceAlarmResponseBodyDataVoList {
	return s.VoList
}

func (s *ListInstanceAlarmResponseBodyData) SetCurrentPage(v int32) *ListInstanceAlarmResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyData) SetPageSize(v int32) *ListInstanceAlarmResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyData) SetTotalCount(v int64) *ListInstanceAlarmResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyData) SetVoList(v *ListInstanceAlarmResponseBodyDataVoList) *ListInstanceAlarmResponseBodyData {
	s.VoList = v
	return s
}

func (s *ListInstanceAlarmResponseBodyData) Validate() error {
	if s.VoList != nil {
		if err := s.VoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceAlarmResponseBodyDataVoList struct {
	CommodityInstanceAlarmVO []*ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO `json:"CommodityInstanceAlarmVO,omitempty" xml:"CommodityInstanceAlarmVO,omitempty" type:"Repeated"`
}

func (s ListInstanceAlarmResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAlarmResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *ListInstanceAlarmResponseBodyDataVoList) GetCommodityInstanceAlarmVO() []*ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO {
	return s.CommodityInstanceAlarmVO
}

func (s *ListInstanceAlarmResponseBodyDataVoList) SetCommodityInstanceAlarmVO(v []*ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO) *ListInstanceAlarmResponseBodyDataVoList {
	s.CommodityInstanceAlarmVO = v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoList) Validate() error {
	if s.CommodityInstanceAlarmVO != nil {
		for _, item := range s.CommodityInstanceAlarmVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO struct {
	AlarmVO      *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO `json:"AlarmVO,omitempty" xml:"AlarmVO,omitempty" type:"Struct"`
	InstanceId   *string                                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string                                                                 `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO) GoString() string {
	return s.String()
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO) GetAlarmVO() *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO {
	return s.AlarmVO
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO) SetAlarmVO(v *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO {
	s.AlarmVO = v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO) SetInstanceId(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO) SetInstanceName(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO {
	s.InstanceName = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVO) Validate() error {
	if s.AlarmVO != nil {
		if err := s.AlarmVO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO struct {
	AlarmCount     *int32                                                                              `json:"AlarmCount,omitempty" xml:"AlarmCount,omitempty"`
	AlarmDetails   *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetails `json:"AlarmDetails,omitempty" xml:"AlarmDetails,omitempty" type:"Struct"`
	HasConfigAlarm *bool                                                                               `json:"HasConfigAlarm,omitempty" xml:"HasConfigAlarm,omitempty"`
}

func (s ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO) GoString() string {
	return s.String()
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO) GetAlarmCount() *int32 {
	return s.AlarmCount
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO) GetAlarmDetails() *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetails {
	return s.AlarmDetails
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO) GetHasConfigAlarm() *bool {
	return s.HasConfigAlarm
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO) SetAlarmCount(v int32) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO {
	s.AlarmCount = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO) SetAlarmDetails(v *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetails) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO {
	s.AlarmDetails = v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO) SetHasConfigAlarm(v bool) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO {
	s.HasConfigAlarm = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVO) Validate() error {
	if s.AlarmDetails != nil {
		if err := s.AlarmDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetails struct {
	AlarmDetail []*ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail `json:"AlarmDetail,omitempty" xml:"AlarmDetail,omitempty" type:"Repeated"`
}

func (s ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetails) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetails) GoString() string {
	return s.String()
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetails) GetAlarmDetail() []*ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	return s.AlarmDetail
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetails) SetAlarmDetail(v []*ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetails {
	s.AlarmDetail = v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetails) Validate() error {
	if s.AlarmDetail != nil {
		for _, item := range s.AlarmDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail struct {
	AlertState          *string `json:"AlertState,omitempty" xml:"AlertState,omitempty"`
	ComparisonOperator  *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	ContactGroups       *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	Dimensions          *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EffectiveInterval   *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	EnableState         *bool   `json:"EnableState,omitempty" xml:"EnableState,omitempty"`
	GroupId             *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	MailSubject         *string `json:"MailSubject,omitempty" xml:"MailSubject,omitempty"`
	MetricName          *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Namespace           *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	Period              *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Resources           *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	RuleId              *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName            *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	SilenceTime         *string `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	Statistics          *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold           *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times               *string `json:"Times,omitempty" xml:"Times,omitempty"`
	Webhook             *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GoString() string {
	return s.String()
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetAlertState() *string {
	return s.AlertState
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetDimensions() *string {
	return s.Dimensions
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetEnableState() *bool {
	return s.EnableState
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetGroupId() *string {
	return s.GroupId
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetGroupName() *string {
	return s.GroupName
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetMailSubject() *string {
	return s.MailSubject
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetMetricName() *string {
	return s.MetricName
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetNamespace() *string {
	return s.Namespace
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetNoEffectiveInterval() *string {
	return s.NoEffectiveInterval
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetPeriod() *string {
	return s.Period
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetResources() *string {
	return s.Resources
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetRuleId() *string {
	return s.RuleId
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetRuleName() *string {
	return s.RuleName
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetSilenceTime() *string {
	return s.SilenceTime
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetStatistics() *string {
	return s.Statistics
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetThreshold() *string {
	return s.Threshold
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetTimes() *string {
	return s.Times
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) GetWebhook() *string {
	return s.Webhook
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetAlertState(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.AlertState = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetComparisonOperator(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.ComparisonOperator = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetContactGroups(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.ContactGroups = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetDimensions(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.Dimensions = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetEffectiveInterval(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.EffectiveInterval = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetEnableState(v bool) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.EnableState = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetGroupId(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.GroupId = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetGroupName(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.GroupName = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetMailSubject(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.MailSubject = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetMetricName(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.MetricName = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetNamespace(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.Namespace = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetNoEffectiveInterval(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.NoEffectiveInterval = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetPeriod(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.Period = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetResources(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.Resources = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetRuleId(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.RuleId = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetRuleName(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.RuleName = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetSilenceTime(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.SilenceTime = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetStatistics(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.Statistics = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetThreshold(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.Threshold = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetTimes(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.Times = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) SetWebhook(v string) *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail {
	s.Webhook = &v
	return s
}

func (s *ListInstanceAlarmResponseBodyDataVoListCommodityInstanceAlarmVOAlarmVOAlarmDetailsAlarmDetail) Validate() error {
	return dara.Validate(s)
}
