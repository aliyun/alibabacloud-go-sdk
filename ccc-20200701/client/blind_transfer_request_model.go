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
	CallPriority         *int32  `json:"CallPriority,omitempty" xml:"CallPriority,omitempty"`
	ContactFlowVariables *string `json:"ContactFlowVariables,omitempty" xml:"ContactFlowVariables,omitempty"`
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// job-6538214103685****
	JobId                    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	QueuingOverflowThreshold *int64  `json:"QueuingOverflowThreshold,omitempty" xml:"QueuingOverflowThreshold,omitempty"`
	QueuingTimeoutSeconds    *int64  `json:"QueuingTimeoutSeconds,omitempty" xml:"QueuingTimeoutSeconds,omitempty"`
	RoutingType              *string `json:"RoutingType,omitempty" xml:"RoutingType,omitempty"`
	SkillGroupId             *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	StrategyName             *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	StrategyParams           *string `json:"StrategyParams,omitempty" xml:"StrategyParams,omitempty"`
	Tags                     *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// 60
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agent@ccc-test
	Transferee     *string `json:"Transferee,omitempty" xml:"Transferee,omitempty"`
	TransfereeType *string `json:"TransfereeType,omitempty" xml:"TransfereeType,omitempty"`
	// example:
	//
	// 08314325****
	Transferor *string `json:"Transferor,omitempty" xml:"Transferor,omitempty"`
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
