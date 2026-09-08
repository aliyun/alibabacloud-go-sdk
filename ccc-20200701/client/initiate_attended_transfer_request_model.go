// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateAttendedTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallPriority(v int32) *InitiateAttendedTransferRequest
	GetCallPriority() *int32
	SetDeviceId(v string) *InitiateAttendedTransferRequest
	GetDeviceId() *string
	SetInstanceId(v string) *InitiateAttendedTransferRequest
	GetInstanceId() *string
	SetJobId(v string) *InitiateAttendedTransferRequest
	GetJobId() *string
	SetQueuingOverflowThreshold(v int64) *InitiateAttendedTransferRequest
	GetQueuingOverflowThreshold() *int64
	SetQueuingTimeoutSeconds(v int64) *InitiateAttendedTransferRequest
	GetQueuingTimeoutSeconds() *int64
	SetRoutingType(v string) *InitiateAttendedTransferRequest
	GetRoutingType() *string
	SetStrategyName(v string) *InitiateAttendedTransferRequest
	GetStrategyName() *string
	SetStrategyParams(v string) *InitiateAttendedTransferRequest
	GetStrategyParams() *string
	SetTags(v string) *InitiateAttendedTransferRequest
	GetTags() *string
	SetTimeoutSeconds(v int32) *InitiateAttendedTransferRequest
	GetTimeoutSeconds() *int32
	SetTransferee(v string) *InitiateAttendedTransferRequest
	GetTransferee() *string
	SetTransfereeType(v string) *InitiateAttendedTransferRequest
	GetTransfereeType() *string
	SetTransferor(v string) *InitiateAttendedTransferRequest
	GetTransferor() *string
	SetUserId(v string) *InitiateAttendedTransferRequest
	GetUserId() *string
}

type InitiateAttendedTransferRequest struct {
	// The queuing priority when transferring to a skill group queue. Valid values range from 0 to 9, where 0 is the highest priority and 9 is the lowest.
	//
	// example:
	//
	// 5
	CallPriority *int32 `json:"CallPriority,omitempty" xml:"CallPriority,omitempty"`
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
	// The queuing timeout period in seconds when the transfer target is a skill group queue.
	//
	// example:
	//
	// 10
	QueuingTimeoutSeconds *int64 `json:"QueuingTimeoutSeconds,omitempty" xml:"QueuingTimeoutSeconds,omitempty"`
	// The call assignment type. Valid values are Automatic or Manual. If this parameter is empty, the default value is Automatic, which is also the current system\\"s default behavior. When Manual is selected, you must invoke APIs such as ClaimCall to assign the call to a specific agent.
	//
	// example:
	//
	// Automatic
	RoutingType *string `json:"RoutingType,omitempty" xml:"RoutingType,omitempty"`
	// The policy name for agent assignment when transferring to a skill group queue.
	//
	// example:
	//
	// MOST_IDLE，MOST_SKILLED，MOST_ACQUAINTED，CUSTOMIZED等
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The policy parameters for agent assignment when transferring to a skill group queue.
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
	// Ingest endpoint data, primarily used for extension requirements. Regular users do not need to concern themselves with this.
	//
	// example:
	//
	// a=b
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Timeout duration for the consultation transfer, in seconds. If the transferee does not answer within the specified time, the call is disconnected. This field is optional. Default value is 30 seconds.
	//
	// example:
	//
	// 60
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// The transferee, which can be an agent ID or a skill group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent2@ccc-test
	Transferee *string `json:"Transferee,omitempty" xml:"Transferee,omitempty"`
	// The destination type for the transfer. Valid values are AGENT, SKILL_GROUP, and EXTERNAL. If this parameter is not provided, the system determines the destination type based on the format of the target number. If inaccurate detection occurs, explicitly specify this parameter.
	//
	// example:
	//
	// SKILL_GROUP
	TransfereeType *string `json:"TransfereeType,omitempty" xml:"TransfereeType,omitempty"`
	// The party initiating the transfer. When transferring to an external number, this parameter specifies the caller number. This parameter is invalid when transferring to an internal agent or skill group; in such cases, the initiator is determined by the UserId parameter.
	//
	// example:
	//
	// 无
	Transferor *string `json:"Transferor,omitempty" xml:"Transferor,omitempty"`
	// The agent ID initiating the consultation transfer.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s InitiateAttendedTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s InitiateAttendedTransferRequest) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferRequest) GetCallPriority() *int32 {
	return s.CallPriority
}

func (s *InitiateAttendedTransferRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *InitiateAttendedTransferRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InitiateAttendedTransferRequest) GetJobId() *string {
	return s.JobId
}

func (s *InitiateAttendedTransferRequest) GetQueuingOverflowThreshold() *int64 {
	return s.QueuingOverflowThreshold
}

func (s *InitiateAttendedTransferRequest) GetQueuingTimeoutSeconds() *int64 {
	return s.QueuingTimeoutSeconds
}

func (s *InitiateAttendedTransferRequest) GetRoutingType() *string {
	return s.RoutingType
}

func (s *InitiateAttendedTransferRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *InitiateAttendedTransferRequest) GetStrategyParams() *string {
	return s.StrategyParams
}

func (s *InitiateAttendedTransferRequest) GetTags() *string {
	return s.Tags
}

func (s *InitiateAttendedTransferRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *InitiateAttendedTransferRequest) GetTransferee() *string {
	return s.Transferee
}

func (s *InitiateAttendedTransferRequest) GetTransfereeType() *string {
	return s.TransfereeType
}

func (s *InitiateAttendedTransferRequest) GetTransferor() *string {
	return s.Transferor
}

func (s *InitiateAttendedTransferRequest) GetUserId() *string {
	return s.UserId
}

func (s *InitiateAttendedTransferRequest) SetCallPriority(v int32) *InitiateAttendedTransferRequest {
	s.CallPriority = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetDeviceId(v string) *InitiateAttendedTransferRequest {
	s.DeviceId = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetInstanceId(v string) *InitiateAttendedTransferRequest {
	s.InstanceId = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetJobId(v string) *InitiateAttendedTransferRequest {
	s.JobId = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetQueuingOverflowThreshold(v int64) *InitiateAttendedTransferRequest {
	s.QueuingOverflowThreshold = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetQueuingTimeoutSeconds(v int64) *InitiateAttendedTransferRequest {
	s.QueuingTimeoutSeconds = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetRoutingType(v string) *InitiateAttendedTransferRequest {
	s.RoutingType = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetStrategyName(v string) *InitiateAttendedTransferRequest {
	s.StrategyName = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetStrategyParams(v string) *InitiateAttendedTransferRequest {
	s.StrategyParams = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetTags(v string) *InitiateAttendedTransferRequest {
	s.Tags = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetTimeoutSeconds(v int32) *InitiateAttendedTransferRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetTransferee(v string) *InitiateAttendedTransferRequest {
	s.Transferee = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetTransfereeType(v string) *InitiateAttendedTransferRequest {
	s.TransfereeType = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetTransferor(v string) *InitiateAttendedTransferRequest {
	s.Transferor = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetUserId(v string) *InitiateAttendedTransferRequest {
	s.UserId = &v
	return s
}

func (s *InitiateAttendedTransferRequest) Validate() error {
	return dara.Validate(s)
}
