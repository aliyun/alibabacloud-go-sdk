// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPredictiveCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallee(v string) *StartPredictiveCallRequest
	GetCallee() *string
	SetCaller(v string) *StartPredictiveCallRequest
	GetCaller() *string
	SetContactFlowId(v string) *StartPredictiveCallRequest
	GetContactFlowId() *string
	SetContactFlowVariables(v string) *StartPredictiveCallRequest
	GetContactFlowVariables() *string
	SetInstanceId(v string) *StartPredictiveCallRequest
	GetInstanceId() *string
	SetMaskedCallee(v string) *StartPredictiveCallRequest
	GetMaskedCallee() *string
	SetSkillGroupId(v string) *StartPredictiveCallRequest
	GetSkillGroupId() *string
	SetTags(v string) *StartPredictiveCallRequest
	GetTags() *string
	SetTimeoutSeconds(v int32) *StartPredictiveCallRequest
	GetTimeoutSeconds() *int32
}

type StartPredictiveCallRequest struct {
	// The callee number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1312353****
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// The caller number, which must be an active outbound number under the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0109810****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// The IVR contact flow ID. After the callee answers, the call is automatically transferred into this IVR flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9774c36c-12fe-4e37-adce-89bc77ce****
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// The contact flow variables passed in as a JSON-formatted string of an array. Each array element is a key-value pair, where the key is the variable name and the value is the variable value. To use these variables in the IVR flow, create a Custom Parameter with the same name in the start node of the IVR associated with the specified contact flow ID.
	//
	// example:
	//
	// {"name":"王先生","time":"19点20分","address":"某某中心"}
	ContactFlowVariables *string `json:"ContactFlowVariables,omitempty" xml:"ContactFlowVariables,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The desensitized callee number. If this field is not empty, it indicates that the callee number must be desensitized. The Desensitization Rule is defined by the Customer. Simply enter the desensitized callee number here. Using a desensitized callee number means that in certain scenarios, only the desensitized number is visible, and the real callee number cannot be viewed.
	//
	// example:
	//
	// 1312353****
	MaskedCallee *string `json:"MaskedCallee,omitempty" xml:"MaskedCallee,omitempty"`
	// The skill group ID. This parameter is optional. If specified, the outbound number is selected only from the numbers associated with the specified skill group.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// The ingest endpoint data, primarily used for extension purposes. Regular users do not need to concern themselves with this.
	//
	// example:
	//
	// 无
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The timeout period, in seconds. If the call is not answered within the specified time, it is automatically disconnected.
	//
	// example:
	//
	// 10
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s StartPredictiveCallRequest) String() string {
	return dara.Prettify(s)
}

func (s StartPredictiveCallRequest) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallRequest) GetCallee() *string {
	return s.Callee
}

func (s *StartPredictiveCallRequest) GetCaller() *string {
	return s.Caller
}

func (s *StartPredictiveCallRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *StartPredictiveCallRequest) GetContactFlowVariables() *string {
	return s.ContactFlowVariables
}

func (s *StartPredictiveCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartPredictiveCallRequest) GetMaskedCallee() *string {
	return s.MaskedCallee
}

func (s *StartPredictiveCallRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *StartPredictiveCallRequest) GetTags() *string {
	return s.Tags
}

func (s *StartPredictiveCallRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *StartPredictiveCallRequest) SetCallee(v string) *StartPredictiveCallRequest {
	s.Callee = &v
	return s
}

func (s *StartPredictiveCallRequest) SetCaller(v string) *StartPredictiveCallRequest {
	s.Caller = &v
	return s
}

func (s *StartPredictiveCallRequest) SetContactFlowId(v string) *StartPredictiveCallRequest {
	s.ContactFlowId = &v
	return s
}

func (s *StartPredictiveCallRequest) SetContactFlowVariables(v string) *StartPredictiveCallRequest {
	s.ContactFlowVariables = &v
	return s
}

func (s *StartPredictiveCallRequest) SetInstanceId(v string) *StartPredictiveCallRequest {
	s.InstanceId = &v
	return s
}

func (s *StartPredictiveCallRequest) SetMaskedCallee(v string) *StartPredictiveCallRequest {
	s.MaskedCallee = &v
	return s
}

func (s *StartPredictiveCallRequest) SetSkillGroupId(v string) *StartPredictiveCallRequest {
	s.SkillGroupId = &v
	return s
}

func (s *StartPredictiveCallRequest) SetTags(v string) *StartPredictiveCallRequest {
	s.Tags = &v
	return s
}

func (s *StartPredictiveCallRequest) SetTimeoutSeconds(v int32) *StartPredictiveCallRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *StartPredictiveCallRequest) Validate() error {
	return dara.Validate(s)
}
