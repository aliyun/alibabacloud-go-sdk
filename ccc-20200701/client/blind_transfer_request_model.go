// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlindTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallPriority(v int32) *BlindTransferRequest
	GetCallPriority() *int32
	SetContactFlowVariables(v string) *BlindTransferRequest
	GetContactFlowVariables() *string
	SetDeviceId(v string) *BlindTransferRequest
	GetDeviceId() *string
	SetInstanceId(v string) *BlindTransferRequest
	GetInstanceId() *string
	SetJobId(v string) *BlindTransferRequest
	GetJobId() *string
	SetQueuingOverflowThreshold(v int64) *BlindTransferRequest
	GetQueuingOverflowThreshold() *int64
	SetQueuingTimeoutSeconds(v int64) *BlindTransferRequest
	GetQueuingTimeoutSeconds() *int64
	SetRoutingType(v string) *BlindTransferRequest
	GetRoutingType() *string
	SetSkillGroupId(v string) *BlindTransferRequest
	GetSkillGroupId() *string
	SetStrategyName(v string) *BlindTransferRequest
	GetStrategyName() *string
	SetStrategyParams(v string) *BlindTransferRequest
	GetStrategyParams() *string
	SetTags(v string) *BlindTransferRequest
	GetTags() *string
	SetTimeoutSeconds(v int32) *BlindTransferRequest
	GetTimeoutSeconds() *int32
	SetTransferee(v string) *BlindTransferRequest
	GetTransferee() *string
	SetTransfereeType(v string) *BlindTransferRequest
	GetTransfereeType() *string
	SetTransferor(v string) *BlindTransferRequest
	GetTransferor() *string
	SetUserId(v string) *BlindTransferRequest
	GetUserId() *string
}

type BlindTransferRequest struct {
	// The queue priority when transferring to a skill group. Valid values are 0–9, where 0 is the highest priority and 9 is the lowest.
	//
	// example:
	//
	// 5
	CallPriority *int32 `json:"CallPriority,omitempty" xml:"CallPriority,omitempty"`
	// Variables passed to the contact flow. This field is optional. The variables configured here can be retrieved and used in the IVR flow. The format is a JSON string representing a set of key-value pairs.
	//
	// example:
	//
	// {"name":"王先生","time":"19点20分","address":"某某中心"}
	ContactFlowVariables *string `json:"ContactFlowVariables,omitempty" xml:"ContactFlowVariables,omitempty"`
	// Device ID. This parameter is meaningless and can be filled with any value.
	//
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The call ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The queuing overflow threshold when the transfer target is a skill group queue. The default value is 0, which means no overflow occurs.
	//
	// example:
	//
	// 0
	QueuingOverflowThreshold *int64 `json:"QueuingOverflowThreshold,omitempty" xml:"QueuingOverflowThreshold,omitempty"`
	// The queuing timeout duration in seconds when the transfer target is a skill group queue.
	//
	// example:
	//
	// 10
	QueuingTimeoutSeconds *int64 `json:"QueuingTimeoutSeconds,omitempty" xml:"QueuingTimeoutSeconds,omitempty"`
	// The call routing type. Valid values are Automatic or Manual. If this parameter is empty, the system defaults to Automatic routing, which is also the current default behavior of the system. When Manual routing is selected, you must invoke APIs such as ClaimCall to assign the call to a specific agent.
	//
	// example:
	//
	// Manual
	RoutingType *string `json:"RoutingType,omitempty" xml:"RoutingType,omitempty"`
	// Skill group ID.
	//
	// example:
	//
	// ee914df4-82bf-4919-bcb3-9cb8aa437f35
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// The policy name for agent assignment when transferring to a skill group queue.
	//
	// example:
	//
	// MOST_IDLE，MOST_SKILLED，MOST_ACQUAINTED，CUSTOMIZED等
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The parameters for the agent assignment policy when transferring to a skill group queue.
	//
	// example:
	//
	// 当分配策略为CUSTOMIZED时，本参数的内容为如下格式：
	//
	//  {
	//
	//   "functionId": "512fed64-e379-400f-a1a5-14d5730xxxxx",
	//
	//   "functionName": "routing-strategy-test-2"
	//
	// }
	StrategyParams *string `json:"StrategyParams,omitempty" xml:"StrategyParams,omitempty"`
	// Ingest endpoint data, primarily used for extension purposes. Regular users do not need to concern themselves with this field.
	//
	// example:
	//
	// 5295578135#WAEtqY5U&Biz_Package_Rexian_Zhuanjieanquanyungaojie_2527
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Timeout duration for the direct transfer, in seconds. If the transferee does not answer within the specified time, the call is disconnected. This field is optional and defaults to 30 seconds.
	//
	// example:
	//
	// 60
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// The transfer recipient, which can be either an agent ID or a skill group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent@ccc-test
	Transferee *string `json:"Transferee,omitempty" xml:"Transferee,omitempty"`
	// Destination type for the transfer. Valid values are AGENT, SKILL_GROUP, IVR, and EXTERNAL_NUMBER. If this parameter is not specified, the system determines the destination type based on the format of the target number. If the automatic detection is inaccurate, you must explicitly specify this parameter.
	//
	// example:
	//
	// SKILL_GROUP
	TransfereeType *string `json:"TransfereeType,omitempty" xml:"TransfereeType,omitempty"`
	// The transfer initiator. When the scenario involves directly transferring to an external number, the number specified by this parameter is used as the caller. This parameter is invalid when transferring to an internal agent or skill group; in such cases, the initiator is specified by the UserId parameter.
	//
	// example:
	//
	// 08314325****
	Transferor *string `json:"Transferor,omitempty" xml:"Transferor,omitempty"`
	// The agent ID that initiates a direct transfer.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BlindTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s BlindTransferRequest) GoString() string {
	return s.String()
}

func (s *BlindTransferRequest) GetCallPriority() *int32 {
	return s.CallPriority
}

func (s *BlindTransferRequest) GetContactFlowVariables() *string {
	return s.ContactFlowVariables
}

func (s *BlindTransferRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BlindTransferRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BlindTransferRequest) GetJobId() *string {
	return s.JobId
}

func (s *BlindTransferRequest) GetQueuingOverflowThreshold() *int64 {
	return s.QueuingOverflowThreshold
}

func (s *BlindTransferRequest) GetQueuingTimeoutSeconds() *int64 {
	return s.QueuingTimeoutSeconds
}

func (s *BlindTransferRequest) GetRoutingType() *string {
	return s.RoutingType
}

func (s *BlindTransferRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *BlindTransferRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *BlindTransferRequest) GetStrategyParams() *string {
	return s.StrategyParams
}

func (s *BlindTransferRequest) GetTags() *string {
	return s.Tags
}

func (s *BlindTransferRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *BlindTransferRequest) GetTransferee() *string {
	return s.Transferee
}

func (s *BlindTransferRequest) GetTransfereeType() *string {
	return s.TransfereeType
}

func (s *BlindTransferRequest) GetTransferor() *string {
	return s.Transferor
}

func (s *BlindTransferRequest) GetUserId() *string {
	return s.UserId
}

func (s *BlindTransferRequest) SetCallPriority(v int32) *BlindTransferRequest {
	s.CallPriority = &v
	return s
}

func (s *BlindTransferRequest) SetContactFlowVariables(v string) *BlindTransferRequest {
	s.ContactFlowVariables = &v
	return s
}

func (s *BlindTransferRequest) SetDeviceId(v string) *BlindTransferRequest {
	s.DeviceId = &v
	return s
}

func (s *BlindTransferRequest) SetInstanceId(v string) *BlindTransferRequest {
	s.InstanceId = &v
	return s
}

func (s *BlindTransferRequest) SetJobId(v string) *BlindTransferRequest {
	s.JobId = &v
	return s
}

func (s *BlindTransferRequest) SetQueuingOverflowThreshold(v int64) *BlindTransferRequest {
	s.QueuingOverflowThreshold = &v
	return s
}

func (s *BlindTransferRequest) SetQueuingTimeoutSeconds(v int64) *BlindTransferRequest {
	s.QueuingTimeoutSeconds = &v
	return s
}

func (s *BlindTransferRequest) SetRoutingType(v string) *BlindTransferRequest {
	s.RoutingType = &v
	return s
}

func (s *BlindTransferRequest) SetSkillGroupId(v string) *BlindTransferRequest {
	s.SkillGroupId = &v
	return s
}

func (s *BlindTransferRequest) SetStrategyName(v string) *BlindTransferRequest {
	s.StrategyName = &v
	return s
}

func (s *BlindTransferRequest) SetStrategyParams(v string) *BlindTransferRequest {
	s.StrategyParams = &v
	return s
}

func (s *BlindTransferRequest) SetTags(v string) *BlindTransferRequest {
	s.Tags = &v
	return s
}

func (s *BlindTransferRequest) SetTimeoutSeconds(v int32) *BlindTransferRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *BlindTransferRequest) SetTransferee(v string) *BlindTransferRequest {
	s.Transferee = &v
	return s
}

func (s *BlindTransferRequest) SetTransfereeType(v string) *BlindTransferRequest {
	s.TransfereeType = &v
	return s
}

func (s *BlindTransferRequest) SetTransferor(v string) *BlindTransferRequest {
	s.Transferor = &v
	return s
}

func (s *BlindTransferRequest) SetUserId(v string) *BlindTransferRequest {
	s.UserId = &v
	return s
}

func (s *BlindTransferRequest) Validate() error {
	return dara.Validate(s)
}
