// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPoint(v string) *UpdateSubscriptionRequest
	GetAccessPoint() *string
	SetAliyunUid(v int64) *UpdateSubscriptionRequest
	GetAliyunUid() *int64
	SetDefaultTopic(v string) *UpdateSubscriptionRequest
	GetDefaultTopic() *string
	SetEventSubscriptionsJson(v string) *UpdateSubscriptionRequest
	GetEventSubscriptionsJson() *string
	SetInstanceId(v string) *UpdateSubscriptionRequest
	GetInstanceId() *string
	SetMqInstanceId(v string) *UpdateSubscriptionRequest
	GetMqInstanceId() *string
	SetMqType(v string) *UpdateSubscriptionRequest
	GetMqType() *string
	SetPassword(v string) *UpdateSubscriptionRequest
	GetPassword() *string
	SetProducerId(v string) *UpdateSubscriptionRequest
	GetProducerId() *string
	SetUsername(v string) *UpdateSubscriptionRequest
	GetUsername() *string
}

type UpdateSubscriptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rmq-cn-****.cn-shanghai.rmq.aliyuncs.com:8080
	AccessPoint *string `json:"AccessPoint,omitempty" xml:"AccessPoint,omitempty"`
	AliyunUid   *int64  `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// example:
	//
	// ccc-event
	DefaultTopic *string `json:"DefaultTopic,omitempty" xml:"DefaultTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"Name ":"StopConsultant ","Disabled ":true},{"Name ":"QueueingRerouted ","Disabled ":true},{"Name ":"IvrTracking ","Disabled ":true},{"Name ":"DualTrackRecordingReady ","Disabled ":true},{"Name ":"SatisfactionSurveyResponse ","Disabled ":true},{"Name ":"Dialing ","Disabled ":false},{"Name ":"Abandoned ","Disabled ":true},{"Name ":"QueueingCancelled ","Disabled ":true},{"Name ":"RecordingReady ","Disabled ":true},{"Name ":"StopCoach ","Disabled ":true},{"Name ":"Established ","Disabled ":true},{"Name ":"InitiateConsultant ","Disabled ":true},{"Name ":"Route2IVR ","Disabled ":false},{"Name ":"Retrieved ","Disabled ":true},{"Name ":"CampaignPaused ","Disabled ":true},{"Name ":"TextStream ","Disabled ":true},{"Name ":"AgentRelease ","Disabled ":true},{"Name ":"VoicemailReady ","Disabled ":true},{"Name ":"Released ","Disabled ":false},{"Name ":"CDRReady ","Disabled ":true},{"Name ":"CaseAttempted ","Disabled ":true},{"Name ":"AgentBreak ","Disabled ":true},{"Name ":"AgentRingingTimeout ","Disabled ":true},{"Name ":"AgentReady ","Disabled ":true},{"Name ":"CampaignResumed ","Disabled ":true},{"Name ":"AgentDialing ","Disabled ":true},{"Name ":"Ringing ","Disabled ":true},{"Name ":"StartConsultant ","Disabled ":false},{"Name ":"QueueingTimeout ","Disabled ":true},{"Name ":"AgentTalk ","Disabled ":true},{"Name ":"AgentRinging ","Disabled ":true},{"Name ":"StartMonitor ","Disabled ":true},{"Name ":"BlindTransfer ","Disabled ":true},{"Name ":"Intercept ","Disabled ":true},{"Name ":"RingingTimeout ","Disabled ":true},{"Name ":"StartConference ","Disabled ":true},{"Name ":"QueueingOverflow ","Disabled ":true},{"Name ":"BargeIn ","Disabled ":true},{"Name ":"StopConference ","Disabled ":true},{"Name ":"AgentCheckOut ","Disabled ":true},{"Name ":"CampaignSubmitted ","Disabled ":true},{"Name ":"AgentCheckIn ","Disabled ":true},{"Name ":"Enqueue ","Disabled ":false},{"Name ":"AttendedTransfer ","Disabled ":true},{"Name ":"StopMonitor ","Disabled ":true},{"Name ":"DispatchingFailure ","Disabled ":true},{"Name ":"SatisfactionSurveyOffer ","Disabled ":true},{"Name ":"CampaignCompleted ","Disabled ":true},{"Name ":"CampaignAborted ","Disabled ":true},{"Name ":"AssignAgent ","Disabled ":true},{"Name ":"Held ","Disabled ":true},{"Name ":"StartCoach ","Disabled ":true}]
	EventSubscriptionsJson *string `json:"EventSubscriptionsJson,omitempty" xml:"EventSubscriptionsJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// rmq-cn-****
	MqInstanceId *string `json:"MqInstanceId,omitempty" xml:"MqInstanceId,omitempty"`
	// example:
	//
	// rocketmq5
	MqType *string `json:"MqType,omitempty" xml:"MqType,omitempty"`
	// example:
	//
	// password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// GID_xxx
	ProducerId *string `json:"ProducerId,omitempty" xml:"ProducerId,omitempty"`
	// example:
	//
	// username
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionRequest) GetAccessPoint() *string {
	return s.AccessPoint
}

func (s *UpdateSubscriptionRequest) GetAliyunUid() *int64 {
	return s.AliyunUid
}

func (s *UpdateSubscriptionRequest) GetDefaultTopic() *string {
	return s.DefaultTopic
}

func (s *UpdateSubscriptionRequest) GetEventSubscriptionsJson() *string {
	return s.EventSubscriptionsJson
}

func (s *UpdateSubscriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateSubscriptionRequest) GetMqInstanceId() *string {
	return s.MqInstanceId
}

func (s *UpdateSubscriptionRequest) GetMqType() *string {
	return s.MqType
}

func (s *UpdateSubscriptionRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateSubscriptionRequest) GetProducerId() *string {
	return s.ProducerId
}

func (s *UpdateSubscriptionRequest) GetUsername() *string {
	return s.Username
}

func (s *UpdateSubscriptionRequest) SetAccessPoint(v string) *UpdateSubscriptionRequest {
	s.AccessPoint = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetAliyunUid(v int64) *UpdateSubscriptionRequest {
	s.AliyunUid = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetDefaultTopic(v string) *UpdateSubscriptionRequest {
	s.DefaultTopic = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetEventSubscriptionsJson(v string) *UpdateSubscriptionRequest {
	s.EventSubscriptionsJson = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetInstanceId(v string) *UpdateSubscriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetMqInstanceId(v string) *UpdateSubscriptionRequest {
	s.MqInstanceId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetMqType(v string) *UpdateSubscriptionRequest {
	s.MqType = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetPassword(v string) *UpdateSubscriptionRequest {
	s.Password = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetProducerId(v string) *UpdateSubscriptionRequest {
	s.ProducerId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetUsername(v string) *UpdateSubscriptionRequest {
	s.Username = &v
	return s
}

func (s *UpdateSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
