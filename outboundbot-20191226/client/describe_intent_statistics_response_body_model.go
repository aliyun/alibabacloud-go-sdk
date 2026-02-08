// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntentStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeIntentStatisticsResponseBody
	GetCode() *string
	SetGlobalIntentNum(v int32) *DescribeIntentStatisticsResponseBody
	GetGlobalIntentNum() *int32
	SetGlobalIntents(v []*DescribeIntentStatisticsResponseBodyGlobalIntents) *DescribeIntentStatisticsResponseBody
	GetGlobalIntents() []*DescribeIntentStatisticsResponseBodyGlobalIntents
	SetGroupId(v string) *DescribeIntentStatisticsResponseBody
	GetGroupId() *string
	SetHttpStatusCode(v int32) *DescribeIntentStatisticsResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *DescribeIntentStatisticsResponseBody
	GetInstanceId() *string
	SetIntentsAfterNoAnswer(v []*DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) *DescribeIntentStatisticsResponseBody
	GetIntentsAfterNoAnswer() []*DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer
	SetMessage(v string) *DescribeIntentStatisticsResponseBody
	GetMessage() *string
	SetProcessIntentNum(v int32) *DescribeIntentStatisticsResponseBody
	GetProcessIntentNum() *int32
	SetProcessIntents(v []*DescribeIntentStatisticsResponseBodyProcessIntents) *DescribeIntentStatisticsResponseBody
	GetProcessIntents() []*DescribeIntentStatisticsResponseBodyProcessIntents
	SetRequestId(v string) *DescribeIntentStatisticsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeIntentStatisticsResponseBody
	GetSuccess() *bool
}

type DescribeIntentStatisticsResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Total number of global intents
	//
	// example:
	//
	// 100
	GlobalIntentNum *int32 `json:"GlobalIntentNum,omitempty" xml:"GlobalIntentNum,omitempty"`
	// Array of global events
	//
	// example:
	//
	// []
	GlobalIntents []*DescribeIntentStatisticsResponseBodyGlobalIntents `json:"GlobalIntents,omitempty" xml:"GlobalIntents,omitempty" type:"Repeated"`
	// Group ID to which the instance belongs.
	//
	// example:
	//
	// 0c3f352f-d045-491d-9ce7-11f2d2b7775d
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Instance ID
	//
	// example:
	//
	// a4274627-265f-4e14-b2d6-4ee7d4f8593e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// List of unrecognized intent information
	IntentsAfterNoAnswer []*DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer `json:"IntentsAfterNoAnswer,omitempty" xml:"IntentsAfterNoAnswer,omitempty" type:"Repeated"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Total count of flow intents
	//
	// example:
	//
	// 123
	ProcessIntentNum *int32 `json:"ProcessIntentNum,omitempty" xml:"ProcessIntentNum,omitempty"`
	// List of intent information
	//
	// example:
	//
	// []
	ProcessIntents []*DescribeIntentStatisticsResponseBodyProcessIntents `json:"ProcessIntents,omitempty" xml:"ProcessIntents,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeIntentStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntentStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIntentStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeIntentStatisticsResponseBody) GetGlobalIntentNum() *int32 {
	return s.GlobalIntentNum
}

func (s *DescribeIntentStatisticsResponseBody) GetGlobalIntents() []*DescribeIntentStatisticsResponseBodyGlobalIntents {
	return s.GlobalIntents
}

func (s *DescribeIntentStatisticsResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeIntentStatisticsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeIntentStatisticsResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIntentStatisticsResponseBody) GetIntentsAfterNoAnswer() []*DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer {
	return s.IntentsAfterNoAnswer
}

func (s *DescribeIntentStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeIntentStatisticsResponseBody) GetProcessIntentNum() *int32 {
	return s.ProcessIntentNum
}

func (s *DescribeIntentStatisticsResponseBody) GetProcessIntents() []*DescribeIntentStatisticsResponseBodyProcessIntents {
	return s.ProcessIntents
}

func (s *DescribeIntentStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIntentStatisticsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeIntentStatisticsResponseBody) SetCode(v string) *DescribeIntentStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBody) SetGlobalIntentNum(v int32) *DescribeIntentStatisticsResponseBody {
	s.GlobalIntentNum = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBody) SetGlobalIntents(v []*DescribeIntentStatisticsResponseBodyGlobalIntents) *DescribeIntentStatisticsResponseBody {
	s.GlobalIntents = v
	return s
}

func (s *DescribeIntentStatisticsResponseBody) SetGroupId(v string) *DescribeIntentStatisticsResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBody) SetHttpStatusCode(v int32) *DescribeIntentStatisticsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBody) SetInstanceId(v string) *DescribeIntentStatisticsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBody) SetIntentsAfterNoAnswer(v []*DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) *DescribeIntentStatisticsResponseBody {
	s.IntentsAfterNoAnswer = v
	return s
}

func (s *DescribeIntentStatisticsResponseBody) SetMessage(v string) *DescribeIntentStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBody) SetProcessIntentNum(v int32) *DescribeIntentStatisticsResponseBody {
	s.ProcessIntentNum = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBody) SetProcessIntents(v []*DescribeIntentStatisticsResponseBodyProcessIntents) *DescribeIntentStatisticsResponseBody {
	s.ProcessIntents = v
	return s
}

func (s *DescribeIntentStatisticsResponseBody) SetRequestId(v string) *DescribeIntentStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBody) SetSuccess(v bool) *DescribeIntentStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBody) Validate() error {
	if s.GlobalIntents != nil {
		for _, item := range s.GlobalIntents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IntentsAfterNoAnswer != nil {
		for _, item := range s.IntentsAfterNoAnswer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ProcessIntents != nil {
		for _, item := range s.ProcessIntents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIntentStatisticsResponseBodyGlobalIntents struct {
	// Task group ID
	//
	// example:
	//
	// 0c3f352f-d045-491d-9ce7-11f2d2b7775d
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Hits followed by no acknowledgement
	//
	// example:
	//
	// 1
	HitAfterNoAnswer *int32 `json:"HitAfterNoAnswer,omitempty" xml:"HitAfterNoAnswer,omitempty"`
	// Hit quantity
	//
	// example:
	//
	// 11
	HitNum *int32 `json:"HitNum,omitempty" xml:"HitNum,omitempty"`
	// Instance ID
	//
	// example:
	//
	// a4274627-265f-4e14-b2d6-4ee7d4f8593e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Intent ID
	//
	// example:
	//
	// 21343425
	IntentId *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// Intent name
	//
	// example:
	//
	// 请假意图
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// Type
	//
	// GlobalIntent(0): global intent
	//
	// ProcessIntent(1): flow intent
	//
	// example:
	//
	// GlobalIntent
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeIntentStatisticsResponseBodyGlobalIntents) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntentStatisticsResponseBodyGlobalIntents) GoString() string {
	return s.String()
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) GetHitAfterNoAnswer() *int32 {
	return s.HitAfterNoAnswer
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) GetHitNum() *int32 {
	return s.HitNum
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) GetIntentId() *string {
	return s.IntentId
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) GetIntentName() *string {
	return s.IntentName
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) GetType() *string {
	return s.Type
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) SetGroupId(v string) *DescribeIntentStatisticsResponseBodyGlobalIntents {
	s.GroupId = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) SetHitAfterNoAnswer(v int32) *DescribeIntentStatisticsResponseBodyGlobalIntents {
	s.HitAfterNoAnswer = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) SetHitNum(v int32) *DescribeIntentStatisticsResponseBodyGlobalIntents {
	s.HitNum = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) SetInstanceId(v string) *DescribeIntentStatisticsResponseBodyGlobalIntents {
	s.InstanceId = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) SetIntentId(v string) *DescribeIntentStatisticsResponseBodyGlobalIntents {
	s.IntentId = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) SetIntentName(v string) *DescribeIntentStatisticsResponseBodyGlobalIntents {
	s.IntentName = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) SetType(v string) *DescribeIntentStatisticsResponseBodyGlobalIntents {
	s.Type = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyGlobalIntents) Validate() error {
	return dara.Validate(s)
}

type DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer struct {
	// Task group ID
	//
	// example:
	//
	// b24d321a-2a74-4dd1-a0ba-4ab09cef6652
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Hits with no acknowledgement after matching
	//
	// example:
	//
	// 1
	HitAfterNoAnswer *int32 `json:"HitAfterNoAnswer,omitempty" xml:"HitAfterNoAnswer,omitempty"`
	// Instance ID
	//
	// example:
	//
	// a4274627-265f-4e14-b2d6-4ee7d4f8593e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Intent ID
	//
	// example:
	//
	// 10398032
	IntentId *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// Intent name
	//
	// example:
	//
	// 请假意图
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
}

func (s DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) GoString() string {
	return s.String()
}

func (s *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) GetHitAfterNoAnswer() *int32 {
	return s.HitAfterNoAnswer
}

func (s *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) GetIntentId() *string {
	return s.IntentId
}

func (s *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) GetIntentName() *string {
	return s.IntentName
}

func (s *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) SetGroupId(v string) *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer {
	s.GroupId = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) SetHitAfterNoAnswer(v int32) *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer {
	s.HitAfterNoAnswer = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) SetInstanceId(v string) *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer {
	s.InstanceId = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) SetIntentId(v string) *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer {
	s.IntentId = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) SetIntentName(v string) *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer {
	s.IntentName = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyIntentsAfterNoAnswer) Validate() error {
	return dara.Validate(s)
}

type DescribeIntentStatisticsResponseBodyProcessIntents struct {
	// Task group ID
	//
	// example:
	//
	// 0c3f352f-d045-491d-9ce7-11f2d2b7775d
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Hits followed by no acknowledgement
	//
	// example:
	//
	// 1
	HitAfterNoAnswer *int32 `json:"HitAfterNoAnswer,omitempty" xml:"HitAfterNoAnswer,omitempty"`
	// Hit quantity
	//
	// example:
	//
	// 10
	HitNum *int32 `json:"HitNum,omitempty" xml:"HitNum,omitempty"`
	// Instance ID
	//
	// example:
	//
	// a4274627-265f-4e14-b2d6-4ee7d4f8593e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Intent ID
	//
	// example:
	//
	// 12343
	IntentId *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// Intent name
	//
	// example:
	//
	// 请假意图
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// Quantity percentage
	//
	// example:
	//
	// 33.33%
	RateDisplay *string `json:"RateDisplay,omitempty" xml:"RateDisplay,omitempty"`
	// Type
	//
	// - GlobalIntent(0): global intent
	//
	// - ProcessIntent(1): flow intent
	//
	// example:
	//
	// ProcessIntent
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeIntentStatisticsResponseBodyProcessIntents) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntentStatisticsResponseBodyProcessIntents) GoString() string {
	return s.String()
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) GetHitAfterNoAnswer() *int32 {
	return s.HitAfterNoAnswer
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) GetHitNum() *int32 {
	return s.HitNum
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) GetIntentId() *string {
	return s.IntentId
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) GetIntentName() *string {
	return s.IntentName
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) GetRateDisplay() *string {
	return s.RateDisplay
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) GetType() *string {
	return s.Type
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) SetGroupId(v string) *DescribeIntentStatisticsResponseBodyProcessIntents {
	s.GroupId = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) SetHitAfterNoAnswer(v int32) *DescribeIntentStatisticsResponseBodyProcessIntents {
	s.HitAfterNoAnswer = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) SetHitNum(v int32) *DescribeIntentStatisticsResponseBodyProcessIntents {
	s.HitNum = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) SetInstanceId(v string) *DescribeIntentStatisticsResponseBodyProcessIntents {
	s.InstanceId = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) SetIntentId(v string) *DescribeIntentStatisticsResponseBodyProcessIntents {
	s.IntentId = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) SetIntentName(v string) *DescribeIntentStatisticsResponseBodyProcessIntents {
	s.IntentName = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) SetRateDisplay(v string) *DescribeIntentStatisticsResponseBodyProcessIntents {
	s.RateDisplay = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) SetType(v string) *DescribeIntentStatisticsResponseBodyProcessIntents {
	s.Type = &v
	return s
}

func (s *DescribeIntentStatisticsResponseBodyProcessIntents) Validate() error {
	return dara.Validate(s)
}
