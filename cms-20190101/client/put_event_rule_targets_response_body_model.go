// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEventRuleTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutEventRuleTargetsResponseBody
	GetCode() *string
	SetFailedContactParameters(v *PutEventRuleTargetsResponseBodyFailedContactParameters) *PutEventRuleTargetsResponseBody
	GetFailedContactParameters() *PutEventRuleTargetsResponseBodyFailedContactParameters
	SetFailedFcParameters(v *PutEventRuleTargetsResponseBodyFailedFcParameters) *PutEventRuleTargetsResponseBody
	GetFailedFcParameters() *PutEventRuleTargetsResponseBodyFailedFcParameters
	SetFailedMnsParameters(v *PutEventRuleTargetsResponseBodyFailedMnsParameters) *PutEventRuleTargetsResponseBody
	GetFailedMnsParameters() *PutEventRuleTargetsResponseBodyFailedMnsParameters
	SetFailedParameterCount(v string) *PutEventRuleTargetsResponseBody
	GetFailedParameterCount() *string
	SetMessage(v string) *PutEventRuleTargetsResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutEventRuleTargetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutEventRuleTargetsResponseBody
	GetSuccess() *bool
}

type PutEventRuleTargetsResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is returned if the specified alert contact groups in the request failed to be created or modified.
	FailedContactParameters *PutEventRuleTargetsResponseBodyFailedContactParameters `json:"FailedContactParameters,omitempty" xml:"FailedContactParameters,omitempty" type:"Struct"`
	// This parameter is returned if the specified functions in the request failed to be created or modified in Function Compute.
	FailedFcParameters *PutEventRuleTargetsResponseBodyFailedFcParameters `json:"FailedFcParameters,omitempty" xml:"FailedFcParameters,omitempty" type:"Struct"`
	// This parameter is returned if the specified queues in the request failed to be created or modified in SMQ.
	FailedMnsParameters *PutEventRuleTargetsResponseBodyFailedMnsParameters `json:"FailedMnsParameters,omitempty" xml:"FailedMnsParameters,omitempty" type:"Struct"`
	// The number of resources that failed to be created or modified.
	//
	// example:
	//
	// 2
	FailedParameterCount *string `json:"FailedParameterCount,omitempty" xml:"FailedParameterCount,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 409C64DA-CF14-45DF-B463-471C790DD15A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutEventRuleTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutEventRuleTargetsResponseBody) GetFailedContactParameters() *PutEventRuleTargetsResponseBodyFailedContactParameters {
	return s.FailedContactParameters
}

func (s *PutEventRuleTargetsResponseBody) GetFailedFcParameters() *PutEventRuleTargetsResponseBodyFailedFcParameters {
	return s.FailedFcParameters
}

func (s *PutEventRuleTargetsResponseBody) GetFailedMnsParameters() *PutEventRuleTargetsResponseBodyFailedMnsParameters {
	return s.FailedMnsParameters
}

func (s *PutEventRuleTargetsResponseBody) GetFailedParameterCount() *string {
	return s.FailedParameterCount
}

func (s *PutEventRuleTargetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutEventRuleTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutEventRuleTargetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutEventRuleTargetsResponseBody) SetCode(v string) *PutEventRuleTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *PutEventRuleTargetsResponseBody) SetFailedContactParameters(v *PutEventRuleTargetsResponseBodyFailedContactParameters) *PutEventRuleTargetsResponseBody {
	s.FailedContactParameters = v
	return s
}

func (s *PutEventRuleTargetsResponseBody) SetFailedFcParameters(v *PutEventRuleTargetsResponseBodyFailedFcParameters) *PutEventRuleTargetsResponseBody {
	s.FailedFcParameters = v
	return s
}

func (s *PutEventRuleTargetsResponseBody) SetFailedMnsParameters(v *PutEventRuleTargetsResponseBodyFailedMnsParameters) *PutEventRuleTargetsResponseBody {
	s.FailedMnsParameters = v
	return s
}

func (s *PutEventRuleTargetsResponseBody) SetFailedParameterCount(v string) *PutEventRuleTargetsResponseBody {
	s.FailedParameterCount = &v
	return s
}

func (s *PutEventRuleTargetsResponseBody) SetMessage(v string) *PutEventRuleTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *PutEventRuleTargetsResponseBody) SetRequestId(v string) *PutEventRuleTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutEventRuleTargetsResponseBody) SetSuccess(v bool) *PutEventRuleTargetsResponseBody {
	s.Success = &v
	return s
}

func (s *PutEventRuleTargetsResponseBody) Validate() error {
	if s.FailedContactParameters != nil {
		if err := s.FailedContactParameters.Validate(); err != nil {
			return err
		}
	}
	if s.FailedFcParameters != nil {
		if err := s.FailedFcParameters.Validate(); err != nil {
			return err
		}
	}
	if s.FailedMnsParameters != nil {
		if err := s.FailedMnsParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutEventRuleTargetsResponseBodyFailedContactParameters struct {
	ContactParameter []*PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter `json:"ContactParameter,omitempty" xml:"ContactParameter,omitempty" type:"Repeated"`
}

func (s PutEventRuleTargetsResponseBodyFailedContactParameters) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsResponseBodyFailedContactParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponseBodyFailedContactParameters) GetContactParameter() []*PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter {
	return s.ContactParameter
}

func (s *PutEventRuleTargetsResponseBodyFailedContactParameters) SetContactParameter(v []*PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) *PutEventRuleTargetsResponseBodyFailedContactParameters {
	s.ContactParameter = v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedContactParameters) Validate() error {
	if s.ContactParameter != nil {
		for _, item := range s.ContactParameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter struct {
	// The name of the alert contact group.
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	// The ID of the recipient.
	//
	// example:
	//
	// 2
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The alert notification methods. Valid values:
	//
	// 4: Alert notifications are sent by using DingTalk and emails.
	//
	// example:
	//
	// 3
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) GetId() *int32 {
	return s.Id
}

func (s *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) GetLevel() *string {
	return s.Level
}

func (s *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) SetContactGroupName(v string) *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter {
	s.ContactGroupName = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) SetId(v int32) *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter {
	s.Id = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) SetLevel(v string) *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter {
	s.Level = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) Validate() error {
	return dara.Validate(s)
}

type PutEventRuleTargetsResponseBodyFailedFcParameters struct {
	FcParameter []*PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter `json:"FcParameter,omitempty" xml:"FcParameter,omitempty" type:"Repeated"`
}

func (s PutEventRuleTargetsResponseBodyFailedFcParameters) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsResponseBodyFailedFcParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParameters) GetFcParameter() []*PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter {
	return s.FcParameter
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParameters) SetFcParameter(v []*PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) *PutEventRuleTargetsResponseBodyFailedFcParameters {
	s.FcParameter = v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParameters) Validate() error {
	if s.FcParameter != nil {
		for _, item := range s.FcParameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter struct {
	// The name of the function.
	//
	// example:
	//
	// functionTest1
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The ID of the recipient.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the Function Compute service.
	//
	// example:
	//
	// serviceTest1
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) GetFunctionName() *string {
	return s.FunctionName
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) GetId() *int32 {
	return s.Id
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) GetRegion() *string {
	return s.Region
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) GetServiceName() *string {
	return s.ServiceName
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) SetFunctionName(v string) *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter {
	s.FunctionName = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) SetId(v int32) *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter {
	s.Id = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) SetRegion(v string) *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter {
	s.Region = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) SetServiceName(v string) *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter {
	s.ServiceName = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) Validate() error {
	return dara.Validate(s)
}

type PutEventRuleTargetsResponseBodyFailedMnsParameters struct {
	MnsParameter []*PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter `json:"MnsParameter,omitempty" xml:"MnsParameter,omitempty" type:"Repeated"`
}

func (s PutEventRuleTargetsResponseBodyFailedMnsParameters) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsResponseBodyFailedMnsParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponseBodyFailedMnsParameters) GetMnsParameter() []*PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter {
	return s.MnsParameter
}

func (s *PutEventRuleTargetsResponseBodyFailedMnsParameters) SetMnsParameter(v []*PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) *PutEventRuleTargetsResponseBodyFailedMnsParameters {
	s.MnsParameter = v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedMnsParameters) Validate() error {
	if s.MnsParameter != nil {
		for _, item := range s.MnsParameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter struct {
	// The ID of the recipient.
	//
	// example:
	//
	// 2
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the MNS queue.
	//
	// example:
	//
	// testQueue
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) GetId() *int32 {
	return s.Id
}

func (s *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) GetQueue() *string {
	return s.Queue
}

func (s *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) GetRegion() *string {
	return s.Region
}

func (s *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) SetId(v int32) *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter {
	s.Id = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) SetQueue(v string) *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter {
	s.Queue = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) SetRegion(v string) *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter {
	s.Region = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) Validate() error {
	return dara.Validate(s)
}
