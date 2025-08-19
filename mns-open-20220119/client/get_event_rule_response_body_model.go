// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetEventRuleResponseBody
	GetCode() *int64
	SetData(v *GetEventRuleResponseBodyData) *GetEventRuleResponseBody
	GetData() *GetEventRuleResponseBodyData
	SetMessage(v string) *GetEventRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEventRuleResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetEventRuleResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetEventRuleResponseBody
	GetSuccess() *bool
}

type GetEventRuleResponseBody struct {
	// example:
	//
	// 200
	Code *int64                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetEventRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEventRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventRuleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetEventRuleResponseBody) GetData() *GetEventRuleResponseBodyData {
	return s.Data
}

func (s *GetEventRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEventRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEventRuleResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetEventRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEventRuleResponseBody) SetCode(v int64) *GetEventRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetEventRuleResponseBody) SetData(v *GetEventRuleResponseBodyData) *GetEventRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetEventRuleResponseBody) SetMessage(v string) *GetEventRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetEventRuleResponseBody) SetRequestId(v string) *GetEventRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEventRuleResponseBody) SetStatus(v string) *GetEventRuleResponseBody {
	s.Status = &v
	return s
}

func (s *GetEventRuleResponseBody) SetSuccess(v bool) *GetEventRuleResponseBody {
	s.Success = &v
	return s
}

func (s *GetEventRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetEventRuleResponseBodyData struct {
	// example:
	//
	// --
	DeliveryMode *string                               `json:"DeliveryMode,omitempty" xml:"DeliveryMode,omitempty"`
	Endpoint     *GetEventRuleResponseBodyDataEndpoint `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Struct"`
	EventTypes   []*string                             `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	MatchRules   [][]*EventMatchRule                   `json:"MatchRules,omitempty" xml:"MatchRules,omitempty" type:"Repeated"`
	// example:
	//
	// event-recorder
	RuleName      *string                                      `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Subscriptions []*GetEventRuleResponseBodyDataSubscriptions `json:"Subscriptions,omitempty" xml:"Subscriptions,omitempty" type:"Repeated"`
	// example:
	//
	// demo-topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetEventRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEventRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEventRuleResponseBodyData) GetDeliveryMode() *string {
	return s.DeliveryMode
}

func (s *GetEventRuleResponseBodyData) GetEndpoint() *GetEventRuleResponseBodyDataEndpoint {
	return s.Endpoint
}

func (s *GetEventRuleResponseBodyData) GetEventTypes() []*string {
	return s.EventTypes
}

func (s *GetEventRuleResponseBodyData) GetMatchRules() [][]*EventMatchRule {
	return s.MatchRules
}

func (s *GetEventRuleResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *GetEventRuleResponseBodyData) GetSubscriptions() []*GetEventRuleResponseBodyDataSubscriptions {
	return s.Subscriptions
}

func (s *GetEventRuleResponseBodyData) GetTopicName() *string {
	return s.TopicName
}

func (s *GetEventRuleResponseBodyData) SetDeliveryMode(v string) *GetEventRuleResponseBodyData {
	s.DeliveryMode = &v
	return s
}

func (s *GetEventRuleResponseBodyData) SetEndpoint(v *GetEventRuleResponseBodyDataEndpoint) *GetEventRuleResponseBodyData {
	s.Endpoint = v
	return s
}

func (s *GetEventRuleResponseBodyData) SetEventTypes(v []*string) *GetEventRuleResponseBodyData {
	s.EventTypes = v
	return s
}

func (s *GetEventRuleResponseBodyData) SetMatchRules(v [][]*EventMatchRule) *GetEventRuleResponseBodyData {
	s.MatchRules = v
	return s
}

func (s *GetEventRuleResponseBodyData) SetRuleName(v string) *GetEventRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetEventRuleResponseBodyData) SetSubscriptions(v []*GetEventRuleResponseBodyDataSubscriptions) *GetEventRuleResponseBodyData {
	s.Subscriptions = v
	return s
}

func (s *GetEventRuleResponseBodyData) SetTopicName(v string) *GetEventRuleResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *GetEventRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetEventRuleResponseBodyDataEndpoint struct {
	// example:
	//
	// queue
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// example:
	//
	// retry-queue
	EndpointValue *string `json:"EndpointValue,omitempty" xml:"EndpointValue,omitempty"`
}

func (s GetEventRuleResponseBodyDataEndpoint) String() string {
	return dara.Prettify(s)
}

func (s GetEventRuleResponseBodyDataEndpoint) GoString() string {
	return s.String()
}

func (s *GetEventRuleResponseBodyDataEndpoint) GetEndpointType() *string {
	return s.EndpointType
}

func (s *GetEventRuleResponseBodyDataEndpoint) GetEndpointValue() *string {
	return s.EndpointValue
}

func (s *GetEventRuleResponseBodyDataEndpoint) SetEndpointType(v string) *GetEventRuleResponseBodyDataEndpoint {
	s.EndpointType = &v
	return s
}

func (s *GetEventRuleResponseBodyDataEndpoint) SetEndpointValue(v string) *GetEventRuleResponseBodyDataEndpoint {
	s.EndpointValue = &v
	return s
}

func (s *GetEventRuleResponseBodyDataEndpoint) Validate() error {
	return dara.Validate(s)
}

type GetEventRuleResponseBodyDataSubscriptions struct {
	// example:
	//
	// queue
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// example:
	//
	// retry-queue
	EndpointValue *string `json:"EndpointValue,omitempty" xml:"EndpointValue,omitempty"`
}

func (s GetEventRuleResponseBodyDataSubscriptions) String() string {
	return dara.Prettify(s)
}

func (s GetEventRuleResponseBodyDataSubscriptions) GoString() string {
	return s.String()
}

func (s *GetEventRuleResponseBodyDataSubscriptions) GetEndpointType() *string {
	return s.EndpointType
}

func (s *GetEventRuleResponseBodyDataSubscriptions) GetEndpointValue() *string {
	return s.EndpointValue
}

func (s *GetEventRuleResponseBodyDataSubscriptions) SetEndpointType(v string) *GetEventRuleResponseBodyDataSubscriptions {
	s.EndpointType = &v
	return s
}

func (s *GetEventRuleResponseBodyDataSubscriptions) SetEndpointValue(v string) *GetEventRuleResponseBodyDataSubscriptions {
	s.EndpointValue = &v
	return s
}

func (s *GetEventRuleResponseBodyDataSubscriptions) Validate() error {
	return dara.Validate(s)
}
