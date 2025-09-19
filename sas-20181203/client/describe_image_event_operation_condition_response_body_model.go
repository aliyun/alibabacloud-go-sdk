// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageEventOperationConditionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeImageEventOperationConditionResponseBody
	GetCode() *string
	SetData(v *DescribeImageEventOperationConditionResponseBodyData) *DescribeImageEventOperationConditionResponseBody
	GetData() *DescribeImageEventOperationConditionResponseBodyData
	SetMessage(v string) *DescribeImageEventOperationConditionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeImageEventOperationConditionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeImageEventOperationConditionResponseBody
	GetSuccess() *bool
}

type DescribeImageEventOperationConditionResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeImageEventOperationConditionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ADE57832-9666-511C-9A80-B87DE2E8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeImageEventOperationConditionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageEventOperationConditionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageEventOperationConditionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeImageEventOperationConditionResponseBody) GetData() *DescribeImageEventOperationConditionResponseBodyData {
	return s.Data
}

func (s *DescribeImageEventOperationConditionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeImageEventOperationConditionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageEventOperationConditionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeImageEventOperationConditionResponseBody) SetCode(v string) *DescribeImageEventOperationConditionResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageEventOperationConditionResponseBody) SetData(v *DescribeImageEventOperationConditionResponseBodyData) *DescribeImageEventOperationConditionResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageEventOperationConditionResponseBody) SetMessage(v string) *DescribeImageEventOperationConditionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeImageEventOperationConditionResponseBody) SetRequestId(v string) *DescribeImageEventOperationConditionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageEventOperationConditionResponseBody) SetSuccess(v bool) *DescribeImageEventOperationConditionResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeImageEventOperationConditionResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageEventOperationConditionResponseBodyData struct {
	// The alert type.
	//
	// 	- Only **sensitiveFile*	- may be returned.
	//
	// example:
	//
	// sensitiveFile
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The operations.
	Operations []*DescribeImageEventOperationConditionResponseBodyDataOperations `json:"Operations,omitempty" xml:"Operations,omitempty" type:"Repeated"`
	// The application scopes of the rules.
	Scenarios []*string `json:"Scenarios,omitempty" xml:"Scenarios,omitempty" type:"Repeated"`
}

func (s DescribeImageEventOperationConditionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageEventOperationConditionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageEventOperationConditionResponseBodyData) GetEventType() *string {
	return s.EventType
}

func (s *DescribeImageEventOperationConditionResponseBodyData) GetOperations() []*DescribeImageEventOperationConditionResponseBodyDataOperations {
	return s.Operations
}

func (s *DescribeImageEventOperationConditionResponseBodyData) GetScenarios() []*string {
	return s.Scenarios
}

func (s *DescribeImageEventOperationConditionResponseBodyData) SetEventType(v string) *DescribeImageEventOperationConditionResponseBodyData {
	s.EventType = &v
	return s
}

func (s *DescribeImageEventOperationConditionResponseBodyData) SetOperations(v []*DescribeImageEventOperationConditionResponseBodyDataOperations) *DescribeImageEventOperationConditionResponseBodyData {
	s.Operations = v
	return s
}

func (s *DescribeImageEventOperationConditionResponseBodyData) SetScenarios(v []*string) *DescribeImageEventOperationConditionResponseBodyData {
	s.Scenarios = v
	return s
}

func (s *DescribeImageEventOperationConditionResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeImageEventOperationConditionResponseBodyDataOperations struct {
	// The rule conditions.
	Conditions []*DescribeImageEventOperationConditionResponseBodyDataOperationsConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The operation code.
	//
	// 	- Only **whitelist*	- may be returned, which indicates that the alert event is added to the whitelist.
	//
	// example:
	//
	// whitelist
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The name of the operation.
	//
	// example:
	//
	// whitelist
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
}

func (s DescribeImageEventOperationConditionResponseBodyDataOperations) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageEventOperationConditionResponseBodyDataOperations) GoString() string {
	return s.String()
}

func (s *DescribeImageEventOperationConditionResponseBodyDataOperations) GetConditions() []*DescribeImageEventOperationConditionResponseBodyDataOperationsConditions {
	return s.Conditions
}

func (s *DescribeImageEventOperationConditionResponseBodyDataOperations) GetOperationCode() *string {
	return s.OperationCode
}

func (s *DescribeImageEventOperationConditionResponseBodyDataOperations) GetOperationName() *string {
	return s.OperationName
}

func (s *DescribeImageEventOperationConditionResponseBodyDataOperations) SetConditions(v []*DescribeImageEventOperationConditionResponseBodyDataOperationsConditions) *DescribeImageEventOperationConditionResponseBodyDataOperations {
	s.Conditions = v
	return s
}

func (s *DescribeImageEventOperationConditionResponseBodyDataOperations) SetOperationCode(v string) *DescribeImageEventOperationConditionResponseBodyDataOperations {
	s.OperationCode = &v
	return s
}

func (s *DescribeImageEventOperationConditionResponseBodyDataOperations) SetOperationName(v string) *DescribeImageEventOperationConditionResponseBodyDataOperations {
	s.OperationName = &v
	return s
}

func (s *DescribeImageEventOperationConditionResponseBodyDataOperations) Validate() error {
	return dara.Validate(s)
}

type DescribeImageEventOperationConditionResponseBodyDataOperationsConditions struct {
	// The keyword of the condition. Valid values:
	//
	// 	- **MD5**
	//
	// 	- **PATH**
	//
	// example:
	//
	// MD5
	ConditionKey *string `json:"ConditionKey,omitempty" xml:"ConditionKey,omitempty"`
	// The name of the condition.
	//
	// example:
	//
	// MD5
	ConditionName *string `json:"ConditionName,omitempty" xml:"ConditionName,omitempty"`
	// The matching types.
	SupportedMisType []*string `json:"SupportedMisType,omitempty" xml:"SupportedMisType,omitempty" type:"Repeated"`
}

func (s DescribeImageEventOperationConditionResponseBodyDataOperationsConditions) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageEventOperationConditionResponseBodyDataOperationsConditions) GoString() string {
	return s.String()
}

func (s *DescribeImageEventOperationConditionResponseBodyDataOperationsConditions) GetConditionKey() *string {
	return s.ConditionKey
}

func (s *DescribeImageEventOperationConditionResponseBodyDataOperationsConditions) GetConditionName() *string {
	return s.ConditionName
}

func (s *DescribeImageEventOperationConditionResponseBodyDataOperationsConditions) GetSupportedMisType() []*string {
	return s.SupportedMisType
}

func (s *DescribeImageEventOperationConditionResponseBodyDataOperationsConditions) SetConditionKey(v string) *DescribeImageEventOperationConditionResponseBodyDataOperationsConditions {
	s.ConditionKey = &v
	return s
}

func (s *DescribeImageEventOperationConditionResponseBodyDataOperationsConditions) SetConditionName(v string) *DescribeImageEventOperationConditionResponseBodyDataOperationsConditions {
	s.ConditionName = &v
	return s
}

func (s *DescribeImageEventOperationConditionResponseBodyDataOperationsConditions) SetSupportedMisType(v []*string) *DescribeImageEventOperationConditionResponseBodyDataOperationsConditions {
	s.SupportedMisType = v
	return s
}

func (s *DescribeImageEventOperationConditionResponseBodyDataOperationsConditions) Validate() error {
	return dara.Validate(s)
}
