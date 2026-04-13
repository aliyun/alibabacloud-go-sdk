// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDetectConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetectConfigs(v []*ListDetectConfigsResponseBodyDetectConfigs) *ListDetectConfigsResponseBody
	GetDetectConfigs() []*ListDetectConfigsResponseBodyDetectConfigs
	SetMaxResults(v int32) *ListDetectConfigsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDetectConfigsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDetectConfigsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDetectConfigsResponseBody
	GetTotalCount() *int32
}

type ListDetectConfigsResponseBody struct {
	DetectConfigs []*ListDetectConfigsResponseBodyDetectConfigs `json:"detectConfigs,omitempty" xml:"detectConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 30BaZ9ekYWXJdqshYecA++coNg7qT1Zbm3RfLyFIZeY=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 79284133-D4BA-56B3-954C-D538256F7EAA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 82
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDetectConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDetectConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDetectConfigsResponseBody) GetDetectConfigs() []*ListDetectConfigsResponseBodyDetectConfigs {
	return s.DetectConfigs
}

func (s *ListDetectConfigsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDetectConfigsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDetectConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDetectConfigsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDetectConfigsResponseBody) SetDetectConfigs(v []*ListDetectConfigsResponseBodyDetectConfigs) *ListDetectConfigsResponseBody {
	s.DetectConfigs = v
	return s
}

func (s *ListDetectConfigsResponseBody) SetMaxResults(v int32) *ListDetectConfigsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDetectConfigsResponseBody) SetNextToken(v string) *ListDetectConfigsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDetectConfigsResponseBody) SetRequestId(v string) *ListDetectConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDetectConfigsResponseBody) SetTotalCount(v int32) *ListDetectConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDetectConfigsResponseBody) Validate() error {
	if s.DetectConfigs != nil {
		for _, item := range s.DetectConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDetectConfigsResponseBodyDetectConfigs struct {
	AlarmConfigs []*ListDetectConfigsResponseBodyDetectConfigsAlarmConfigs `json:"alarmConfigs,omitempty" xml:"alarmConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-04-10T02:30:04Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 0 0 0 ? 	- 1
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// example:
	//
	// this is a description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// dc-xxxx
	DetectConfigId *string `json:"detectConfigId,omitempty" xml:"detectConfigId,omitempty"`
	// example:
	//
	// test
	DetectConfigName *string `json:"detectConfigName,omitempty" xml:"detectConfigName,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// Cron
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s ListDetectConfigsResponseBodyDetectConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListDetectConfigsResponseBodyDetectConfigs) GoString() string {
	return s.String()
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) GetAlarmConfigs() []*ListDetectConfigsResponseBodyDetectConfigsAlarmConfigs {
	return s.AlarmConfigs
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) GetDescription() *string {
	return s.Description
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) GetDetectConfigId() *string {
	return s.DetectConfigId
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) GetDetectConfigName() *string {
	return s.DetectConfigName
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) SetAlarmConfigs(v []*ListDetectConfigsResponseBodyDetectConfigsAlarmConfigs) *ListDetectConfigsResponseBodyDetectConfigs {
	s.AlarmConfigs = v
	return s
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) SetCreateTime(v string) *ListDetectConfigsResponseBodyDetectConfigs {
	s.CreateTime = &v
	return s
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) SetCronExpression(v string) *ListDetectConfigsResponseBodyDetectConfigs {
	s.CronExpression = &v
	return s
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) SetDescription(v string) *ListDetectConfigsResponseBodyDetectConfigs {
	s.Description = &v
	return s
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) SetDetectConfigId(v string) *ListDetectConfigsResponseBodyDetectConfigs {
	s.DetectConfigId = &v
	return s
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) SetDetectConfigName(v string) *ListDetectConfigsResponseBodyDetectConfigs {
	s.DetectConfigName = &v
	return s
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) SetEnabled(v bool) *ListDetectConfigsResponseBodyDetectConfigs {
	s.Enabled = &v
	return s
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) SetTriggerType(v string) *ListDetectConfigsResponseBodyDetectConfigs {
	s.TriggerType = &v
	return s
}

func (s *ListDetectConfigsResponseBodyDetectConfigs) Validate() error {
	if s.AlarmConfigs != nil {
		for _, item := range s.AlarmConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDetectConfigsResponseBodyDetectConfigsAlarmConfigs struct {
	// example:
	//
	// https://metrichub-cms-cn-hangzhou.aliyuncs.com/event/notify?xxxxx
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// cms
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDetectConfigsResponseBodyDetectConfigsAlarmConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListDetectConfigsResponseBodyDetectConfigsAlarmConfigs) GoString() string {
	return s.String()
}

func (s *ListDetectConfigsResponseBodyDetectConfigsAlarmConfigs) GetAddress() *string {
	return s.Address
}

func (s *ListDetectConfigsResponseBodyDetectConfigsAlarmConfigs) GetType() *string {
	return s.Type
}

func (s *ListDetectConfigsResponseBodyDetectConfigsAlarmConfigs) SetAddress(v string) *ListDetectConfigsResponseBodyDetectConfigsAlarmConfigs {
	s.Address = &v
	return s
}

func (s *ListDetectConfigsResponseBodyDetectConfigsAlarmConfigs) SetType(v string) *ListDetectConfigsResponseBodyDetectConfigsAlarmConfigs {
	s.Type = &v
	return s
}

func (s *ListDetectConfigsResponseBodyDetectConfigsAlarmConfigs) Validate() error {
	return dara.Validate(s)
}
