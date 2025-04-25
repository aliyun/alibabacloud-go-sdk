// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type EventMatchRule struct {
	// example:
	//
	// true
	MatchState *bool   `json:"MatchState,omitempty" xml:"MatchState,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Prefix     *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Suffix     *string `json:"Suffix,omitempty" xml:"Suffix,omitempty"`
}

func (s EventMatchRule) String() string {
	return tea.Prettify(s)
}

func (s EventMatchRule) GoString() string {
	return s.String()
}

func (s *EventMatchRule) SetMatchState(v bool) *EventMatchRule {
	s.MatchState = &v
	return s
}

func (s *EventMatchRule) SetName(v string) *EventMatchRule {
	s.Name = &v
	return s
}

func (s *EventMatchRule) SetPrefix(v string) *EventMatchRule {
	s.Prefix = &v
	return s
}

func (s *EventMatchRule) SetSuffix(v string) *EventMatchRule {
	s.Suffix = &v
	return s
}

type AuthorizeEndpointAclRequest struct {
	// The ACL policy. Valid values:
	//
	// 	- **allow**: indicates that this operation is included in the Cidr whitelist. (Only the allow is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// allow
	AclStrategy *string `json:"AclStrategy,omitempty" xml:"AclStrategy,omitempty"`
	// The list of CIDR block.
	//
	// This parameter is required.
	CidrList []*string `json:"CidrList,omitempty" xml:"CidrList,omitempty" type:"Repeated"`
	// The type of the endpoint. Valid values:
	//
	// 	- **public**: indicates public endpoint. (Only the public endpoint is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s AuthorizeEndpointAclRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeEndpointAclRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeEndpointAclRequest) SetAclStrategy(v string) *AuthorizeEndpointAclRequest {
	s.AclStrategy = &v
	return s
}

func (s *AuthorizeEndpointAclRequest) SetCidrList(v []*string) *AuthorizeEndpointAclRequest {
	s.CidrList = v
	return s
}

func (s *AuthorizeEndpointAclRequest) SetEndpointType(v string) *AuthorizeEndpointAclRequest {
	s.EndpointType = &v
	return s
}

type AuthorizeEndpointAclShrinkRequest struct {
	// The ACL policy. Valid values:
	//
	// 	- **allow**: indicates that this operation is included in the Cidr whitelist. (Only the allow is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// allow
	AclStrategy *string `json:"AclStrategy,omitempty" xml:"AclStrategy,omitempty"`
	// The list of CIDR block.
	//
	// This parameter is required.
	CidrListShrink *string `json:"CidrList,omitempty" xml:"CidrList,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **public**: indicates public endpoint. (Only the public endpoint is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s AuthorizeEndpointAclShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeEndpointAclShrinkRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeEndpointAclShrinkRequest) SetAclStrategy(v string) *AuthorizeEndpointAclShrinkRequest {
	s.AclStrategy = &v
	return s
}

func (s *AuthorizeEndpointAclShrinkRequest) SetCidrListShrink(v string) *AuthorizeEndpointAclShrinkRequest {
	s.CidrListShrink = &v
	return s
}

func (s *AuthorizeEndpointAclShrinkRequest) SetEndpointType(v string) *AuthorizeEndpointAclShrinkRequest {
	s.EndpointType = &v
	return s
}

type AuthorizeEndpointAclResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthorizeEndpointAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeEndpointAclResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeEndpointAclResponseBody) SetCode(v int64) *AuthorizeEndpointAclResponseBody {
	s.Code = &v
	return s
}

func (s *AuthorizeEndpointAclResponseBody) SetMessage(v string) *AuthorizeEndpointAclResponseBody {
	s.Message = &v
	return s
}

func (s *AuthorizeEndpointAclResponseBody) SetRequestId(v string) *AuthorizeEndpointAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeEndpointAclResponseBody) SetStatus(v string) *AuthorizeEndpointAclResponseBody {
	s.Status = &v
	return s
}

func (s *AuthorizeEndpointAclResponseBody) SetSuccess(v bool) *AuthorizeEndpointAclResponseBody {
	s.Success = &v
	return s
}

type AuthorizeEndpointAclResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeEndpointAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeEndpointAclResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeEndpointAclResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeEndpointAclResponse) SetHeaders(v map[string]*string) *AuthorizeEndpointAclResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeEndpointAclResponse) SetStatusCode(v int32) *AuthorizeEndpointAclResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeEndpointAclResponse) SetBody(v *AuthorizeEndpointAclResponseBody) *AuthorizeEndpointAclResponse {
	s.Body = v
	return s
}

type CreateEventRuleRequest struct {
	// This parameter is required.
	Endpoints []*CreateEventRuleRequestEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// This parameter is required.
	EventTypes []*string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	MatchRules [][]*EventMatchRule `json:"MatchRules,omitempty" xml:"MatchRules,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// oss
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rule-xsXDW
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateEventRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateEventRuleRequest) SetEndpoints(v []*CreateEventRuleRequestEndpoints) *CreateEventRuleRequest {
	s.Endpoints = v
	return s
}

func (s *CreateEventRuleRequest) SetEventTypes(v []*string) *CreateEventRuleRequest {
	s.EventTypes = v
	return s
}

func (s *CreateEventRuleRequest) SetMatchRules(v [][]*EventMatchRule) *CreateEventRuleRequest {
	s.MatchRules = v
	return s
}

func (s *CreateEventRuleRequest) SetProductName(v string) *CreateEventRuleRequest {
	s.ProductName = &v
	return s
}

func (s *CreateEventRuleRequest) SetRuleName(v string) *CreateEventRuleRequest {
	s.RuleName = &v
	return s
}

type CreateEventRuleRequestEndpoints struct {
	// This parameter is required.
	//
	// example:
	//
	// http
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-xxx-queue
	EndpointValue *string `json:"EndpointValue,omitempty" xml:"EndpointValue,omitempty"`
}

func (s CreateEventRuleRequestEndpoints) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRuleRequestEndpoints) GoString() string {
	return s.String()
}

func (s *CreateEventRuleRequestEndpoints) SetEndpointType(v string) *CreateEventRuleRequestEndpoints {
	s.EndpointType = &v
	return s
}

func (s *CreateEventRuleRequestEndpoints) SetEndpointValue(v string) *CreateEventRuleRequestEndpoints {
	s.EndpointValue = &v
	return s
}

type CreateEventRuleShrinkRequest struct {
	// This parameter is required.
	EndpointsShrink *string `json:"Endpoints,omitempty" xml:"Endpoints,omitempty"`
	// This parameter is required.
	EventTypesShrink *string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty"`
	// This parameter is required.
	MatchRulesShrink *string `json:"MatchRules,omitempty" xml:"MatchRules,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rule-xsXDW
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateEventRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEventRuleShrinkRequest) SetEndpointsShrink(v string) *CreateEventRuleShrinkRequest {
	s.EndpointsShrink = &v
	return s
}

func (s *CreateEventRuleShrinkRequest) SetEventTypesShrink(v string) *CreateEventRuleShrinkRequest {
	s.EventTypesShrink = &v
	return s
}

func (s *CreateEventRuleShrinkRequest) SetMatchRulesShrink(v string) *CreateEventRuleShrinkRequest {
	s.MatchRulesShrink = &v
	return s
}

func (s *CreateEventRuleShrinkRequest) SetProductName(v string) *CreateEventRuleShrinkRequest {
	s.ProductName = &v
	return s
}

func (s *CreateEventRuleShrinkRequest) SetRuleName(v string) *CreateEventRuleShrinkRequest {
	s.RuleName = &v
	return s
}

type CreateEventRuleResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// rule-xsXDW
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEventRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventRuleResponseBody) SetCode(v int64) *CreateEventRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEventRuleResponseBody) SetData(v string) *CreateEventRuleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateEventRuleResponseBody) SetMessage(v string) *CreateEventRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEventRuleResponseBody) SetRequestId(v string) *CreateEventRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventRuleResponseBody) SetStatus(v string) *CreateEventRuleResponseBody {
	s.Status = &v
	return s
}

func (s *CreateEventRuleResponseBody) SetSuccess(v bool) *CreateEventRuleResponseBody {
	s.Success = &v
	return s
}

type CreateEventRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateEventRuleResponse) SetHeaders(v map[string]*string) *CreateEventRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateEventRuleResponse) SetStatusCode(v int32) *CreateEventRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventRuleResponse) SetBody(v *CreateEventRuleResponseBody) *CreateEventRuleResponse {
	s.Body = v
	return s
}

type CreateQueueRequest struct {
	// The period after which all messages sent to the queue are consumed. Valid values: 0 to 604800. Unit: seconds. Default value: 0
	//
	// example:
	//
	// 0
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The dead-letter queue policy.
	DlqPolicy *CreateQueueRequestDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	// Specifies whether to enable the log management feature. Valid values:
	//
	// 	- true: enabled.
	//
	// 	- false: disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// The maximum length of the message that is sent to the queue. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 65536
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is consumed. Valid values: 60 to 604800. Unit: seconds. Default value: 345600.
	//
	// example:
	//
	// 345600
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Valid values: 0 to 30. Unit: seconds. Default value: 0
	//
	// example:
	//
	// 0
	PollingWaitSeconds *int64 `json:"PollingWaitSeconds,omitempty" xml:"PollingWaitSeconds,omitempty"`
	// The name of the queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The tags.
	Tag []*CreateQueueRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 60
	VisibilityTimeout *int64 `json:"VisibilityTimeout,omitempty" xml:"VisibilityTimeout,omitempty"`
}

func (s CreateQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateQueueRequest) SetDelaySeconds(v int64) *CreateQueueRequest {
	s.DelaySeconds = &v
	return s
}

func (s *CreateQueueRequest) SetDlqPolicy(v *CreateQueueRequestDlqPolicy) *CreateQueueRequest {
	s.DlqPolicy = v
	return s
}

func (s *CreateQueueRequest) SetEnableLogging(v bool) *CreateQueueRequest {
	s.EnableLogging = &v
	return s
}

func (s *CreateQueueRequest) SetMaximumMessageSize(v int64) *CreateQueueRequest {
	s.MaximumMessageSize = &v
	return s
}

func (s *CreateQueueRequest) SetMessageRetentionPeriod(v int64) *CreateQueueRequest {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *CreateQueueRequest) SetPollingWaitSeconds(v int64) *CreateQueueRequest {
	s.PollingWaitSeconds = &v
	return s
}

func (s *CreateQueueRequest) SetQueueName(v string) *CreateQueueRequest {
	s.QueueName = &v
	return s
}

func (s *CreateQueueRequest) SetTag(v []*CreateQueueRequestTag) *CreateQueueRequest {
	s.Tag = v
	return s
}

func (s *CreateQueueRequest) SetVisibilityTimeout(v int64) *CreateQueueRequest {
	s.VisibilityTimeout = &v
	return s
}

type CreateQueueRequestDlqPolicy struct {
	// The queue to which dead-letter messages are delivered.
	//
	// example:
	//
	// deadLetterQueue
	DeadLetterTargetQueue *string `json:"DeadLetterTargetQueue,omitempty" xml:"DeadLetterTargetQueue,omitempty"`
	// Specifies whether to enable the dead-letter message delivery.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 3
	MaxReceiveCount *int32 `json:"MaxReceiveCount,omitempty" xml:"MaxReceiveCount,omitempty"`
}

func (s CreateQueueRequestDlqPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueRequestDlqPolicy) GoString() string {
	return s.String()
}

func (s *CreateQueueRequestDlqPolicy) SetDeadLetterTargetQueue(v string) *CreateQueueRequestDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *CreateQueueRequestDlqPolicy) SetEnabled(v bool) *CreateQueueRequestDlqPolicy {
	s.Enabled = &v
	return s
}

func (s *CreateQueueRequestDlqPolicy) SetMaxReceiveCount(v int32) *CreateQueueRequestDlqPolicy {
	s.MaxReceiveCount = &v
	return s
}

type CreateQueueRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateQueueRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueRequestTag) GoString() string {
	return s.String()
}

func (s *CreateQueueRequestTag) SetKey(v string) *CreateQueueRequestTag {
	s.Key = &v
	return s
}

func (s *CreateQueueRequestTag) SetValue(v string) *CreateQueueRequestTag {
	s.Value = &v
	return s
}

type CreateQueueShrinkRequest struct {
	// The period after which all messages sent to the queue are consumed. Valid values: 0 to 604800. Unit: seconds. Default value: 0
	//
	// example:
	//
	// 0
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The dead-letter queue policy.
	DlqPolicyShrink *string `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty"`
	// Specifies whether to enable the log management feature. Valid values:
	//
	// 	- true: enabled.
	//
	// 	- false: disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// The maximum length of the message that is sent to the queue. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 65536
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is consumed. Valid values: 60 to 604800. Unit: seconds. Default value: 345600.
	//
	// example:
	//
	// 345600
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Valid values: 0 to 30. Unit: seconds. Default value: 0
	//
	// example:
	//
	// 0
	PollingWaitSeconds *int64 `json:"PollingWaitSeconds,omitempty" xml:"PollingWaitSeconds,omitempty"`
	// The name of the queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The tags.
	Tag []*CreateQueueShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 60
	VisibilityTimeout *int64 `json:"VisibilityTimeout,omitempty" xml:"VisibilityTimeout,omitempty"`
}

func (s CreateQueueShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateQueueShrinkRequest) SetDelaySeconds(v int64) *CreateQueueShrinkRequest {
	s.DelaySeconds = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetDlqPolicyShrink(v string) *CreateQueueShrinkRequest {
	s.DlqPolicyShrink = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetEnableLogging(v bool) *CreateQueueShrinkRequest {
	s.EnableLogging = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetMaximumMessageSize(v int64) *CreateQueueShrinkRequest {
	s.MaximumMessageSize = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetMessageRetentionPeriod(v int64) *CreateQueueShrinkRequest {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetPollingWaitSeconds(v int64) *CreateQueueShrinkRequest {
	s.PollingWaitSeconds = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetQueueName(v string) *CreateQueueShrinkRequest {
	s.QueueName = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetTag(v []*CreateQueueShrinkRequestTag) *CreateQueueShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateQueueShrinkRequest) SetVisibilityTimeout(v int64) *CreateQueueShrinkRequest {
	s.VisibilityTimeout = &v
	return s
}

type CreateQueueShrinkRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateQueueShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateQueueShrinkRequestTag) SetKey(v string) *CreateQueueShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateQueueShrinkRequestTag) SetValue(v string) *CreateQueueShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateQueueResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123E62C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQueueResponseBody) SetCode(v int64) *CreateQueueResponseBody {
	s.Code = &v
	return s
}

func (s *CreateQueueResponseBody) SetData(v *CreateQueueResponseBodyData) *CreateQueueResponseBody {
	s.Data = v
	return s
}

func (s *CreateQueueResponseBody) SetMessage(v string) *CreateQueueResponseBody {
	s.Message = &v
	return s
}

func (s *CreateQueueResponseBody) SetRequestId(v string) *CreateQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQueueResponseBody) SetStatus(v string) *CreateQueueResponseBody {
	s.Status = &v
	return s
}

func (s *CreateQueueResponseBody) SetSuccess(v bool) *CreateQueueResponseBody {
	s.Success = &v
	return s
}

type CreateQueueResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateQueueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateQueueResponseBodyData) SetCode(v int64) *CreateQueueResponseBodyData {
	s.Code = &v
	return s
}

func (s *CreateQueueResponseBodyData) SetMessage(v string) *CreateQueueResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateQueueResponseBodyData) SetSuccess(v bool) *CreateQueueResponseBodyData {
	s.Success = &v
	return s
}

type CreateQueueResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueResponse) GoString() string {
	return s.String()
}

func (s *CreateQueueResponse) SetHeaders(v map[string]*string) *CreateQueueResponse {
	s.Headers = v
	return s
}

func (s *CreateQueueResponse) SetStatusCode(v int32) *CreateQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQueueResponse) SetBody(v *CreateQueueResponseBody) *CreateQueueResponse {
	s.Body = v
	return s
}

type CreateTopicRequest struct {
	// Specifies whether to enable the log management feature. Valid values:
	//
	// 	- true: enabled.
	//
	// 	- false: disabled. Default value: false.
	//
	// example:
	//
	// true
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// The maximum length of the message that is sent to the topic. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 10240
	MaxMessageSize *int64 `json:"MaxMessageSize,omitempty" xml:"MaxMessageSize,omitempty"`
	// The tags.
	Tag []*CreateTopicRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the topic that you want to create.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s CreateTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicRequest) GoString() string {
	return s.String()
}

func (s *CreateTopicRequest) SetEnableLogging(v bool) *CreateTopicRequest {
	s.EnableLogging = &v
	return s
}

func (s *CreateTopicRequest) SetMaxMessageSize(v int64) *CreateTopicRequest {
	s.MaxMessageSize = &v
	return s
}

func (s *CreateTopicRequest) SetTag(v []*CreateTopicRequestTag) *CreateTopicRequest {
	s.Tag = v
	return s
}

func (s *CreateTopicRequest) SetTopicName(v string) *CreateTopicRequest {
	s.TopicName = &v
	return s
}

type CreateTopicRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// joyce.wang
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTopicRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTopicRequestTag) SetKey(v string) *CreateTopicRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTopicRequestTag) SetValue(v string) *CreateTopicRequestTag {
	s.Value = &v
	return s
}

type CreateTopicResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateTopicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123E62C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTopicResponseBody) SetCode(v int64) *CreateTopicResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTopicResponseBody) SetData(v *CreateTopicResponseBodyData) *CreateTopicResponseBody {
	s.Data = v
	return s
}

func (s *CreateTopicResponseBody) SetMessage(v string) *CreateTopicResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTopicResponseBody) SetRequestId(v string) *CreateTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTopicResponseBody) SetStatus(v string) *CreateTopicResponseBody {
	s.Status = &v
	return s
}

func (s *CreateTopicResponseBody) SetSuccess(v bool) *CreateTopicResponseBody {
	s.Success = &v
	return s
}

type CreateTopicResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTopicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTopicResponseBodyData) SetCode(v int64) *CreateTopicResponseBodyData {
	s.Code = &v
	return s
}

func (s *CreateTopicResponseBodyData) SetMessage(v string) *CreateTopicResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateTopicResponseBodyData) SetSuccess(v bool) *CreateTopicResponseBodyData {
	s.Success = &v
	return s
}

type CreateTopicResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicResponse) GoString() string {
	return s.String()
}

func (s *CreateTopicResponse) SetHeaders(v map[string]*string) *CreateTopicResponse {
	s.Headers = v
	return s
}

func (s *CreateTopicResponse) SetStatusCode(v int32) *CreateTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTopicResponse) SetBody(v *CreateTopicResponseBody) *CreateTopicResponse {
	s.Body = v
	return s
}

type DeleteEventRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// oss
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rule-xsXDW
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DeleteEventRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventRuleRequest) SetProductName(v string) *DeleteEventRuleRequest {
	s.ProductName = &v
	return s
}

func (s *DeleteEventRuleRequest) SetRuleName(v string) *DeleteEventRuleRequest {
	s.RuleName = &v
	return s
}

type DeleteEventRuleResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventRuleResponseBody) SetCode(v int64) *DeleteEventRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventRuleResponseBody) SetMessage(v string) *DeleteEventRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventRuleResponseBody) SetRequestId(v string) *DeleteEventRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventRuleResponseBody) SetStatus(v string) *DeleteEventRuleResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteEventRuleResponseBody) SetSuccess(v bool) *DeleteEventRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteEventRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventRuleResponse) SetHeaders(v map[string]*string) *DeleteEventRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventRuleResponse) SetStatusCode(v int32) *DeleteEventRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventRuleResponse) SetBody(v *DeleteEventRuleResponseBody) *DeleteEventRuleResponse {
	s.Body = v
	return s
}

type DeleteQueueRequest struct {
	// The name of the queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// tf-testAccMNSQueue-525478433321945943
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s DeleteQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueueRequest) SetQueueName(v string) *DeleteQueueRequest {
	s.QueueName = &v
	return s
}

type DeleteQueueResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DeleteQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponseBody) SetCode(v int64) *DeleteQueueResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQueueResponseBody) SetData(v *DeleteQueueResponseBodyData) *DeleteQueueResponseBody {
	s.Data = v
	return s
}

func (s *DeleteQueueResponseBody) SetMessage(v string) *DeleteQueueResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQueueResponseBody) SetRequestId(v string) *DeleteQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQueueResponseBody) SetStatus(v string) *DeleteQueueResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteQueueResponseBody) SetSuccess(v bool) *DeleteQueueResponseBody {
	s.Success = &v
	return s
}

type DeleteQueueResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQueueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponseBodyData) SetCode(v int64) *DeleteQueueResponseBodyData {
	s.Code = &v
	return s
}

func (s *DeleteQueueResponseBodyData) SetMessage(v string) *DeleteQueueResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteQueueResponseBodyData) SetSuccess(v bool) *DeleteQueueResponseBodyData {
	s.Success = &v
	return s
}

type DeleteQueueResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponse) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponse) SetHeaders(v map[string]*string) *DeleteQueueResponse {
	s.Headers = v
	return s
}

func (s *DeleteQueueResponse) SetStatusCode(v int32) *DeleteQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQueueResponse) SetBody(v *DeleteQueueResponseBody) *DeleteQueueResponse {
	s.Body = v
	return s
}

type DeleteTopicRequest struct {
	// The name of the topic that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// tf-testAccMNSTopic-112965059402264645
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s DeleteTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTopicRequest) GoString() string {
	return s.String()
}

func (s *DeleteTopicRequest) SetTopicName(v string) *DeleteTopicRequest {
	s.TopicName = &v
	return s
}

type DeleteTopicResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponseBody) SetCode(v int64) *DeleteTopicResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTopicResponseBody) SetData(v map[string]interface{}) *DeleteTopicResponseBody {
	s.Data = v
	return s
}

func (s *DeleteTopicResponseBody) SetMessage(v string) *DeleteTopicResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTopicResponseBody) SetRequestId(v string) *DeleteTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTopicResponseBody) SetStatus(v string) *DeleteTopicResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteTopicResponseBody) SetSuccess(v bool) *DeleteTopicResponseBody {
	s.Success = &v
	return s
}

type DeleteTopicResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTopicResponse) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponse) SetHeaders(v map[string]*string) *DeleteTopicResponse {
	s.Headers = v
	return s
}

func (s *DeleteTopicResponse) SetStatusCode(v int32) *DeleteTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTopicResponse) SetBody(v *DeleteTopicResponseBody) *DeleteTopicResponse {
	s.Body = v
	return s
}

type DisableEndpointRequest struct {
	// The type of the endpoint. Value:
	//
	// 	- **public**: indicates an public endpoint. (Only the public endpoint is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s DisableEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableEndpointRequest) GoString() string {
	return s.String()
}

func (s *DisableEndpointRequest) SetEndpointType(v string) *DisableEndpointRequest {
	s.EndpointType = &v
	return s
}

type DisableEndpointResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DisableEndpointResponseBody) SetCode(v int64) *DisableEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *DisableEndpointResponseBody) SetMessage(v string) *DisableEndpointResponseBody {
	s.Message = &v
	return s
}

func (s *DisableEndpointResponseBody) SetRequestId(v string) *DisableEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableEndpointResponseBody) SetStatus(v string) *DisableEndpointResponseBody {
	s.Status = &v
	return s
}

func (s *DisableEndpointResponseBody) SetSuccess(v bool) *DisableEndpointResponseBody {
	s.Success = &v
	return s
}

type DisableEndpointResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableEndpointResponse) GoString() string {
	return s.String()
}

func (s *DisableEndpointResponse) SetHeaders(v map[string]*string) *DisableEndpointResponse {
	s.Headers = v
	return s
}

func (s *DisableEndpointResponse) SetStatusCode(v int32) *DisableEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableEndpointResponse) SetBody(v *DisableEndpointResponseBody) *DisableEndpointResponse {
	s.Body = v
	return s
}

type EnableEndpointRequest struct {
	// The type of the endpoint. Valid value:
	//
	// 	- **public**: indicates public endpoint. (Only the public is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s EnableEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableEndpointRequest) GoString() string {
	return s.String()
}

func (s *EnableEndpointRequest) SetEndpointType(v string) *EnableEndpointRequest {
	s.EndpointType = &v
	return s
}

type EnableEndpointResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *EnableEndpointResponseBody) SetCode(v int64) *EnableEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *EnableEndpointResponseBody) SetMessage(v string) *EnableEndpointResponseBody {
	s.Message = &v
	return s
}

func (s *EnableEndpointResponseBody) SetRequestId(v string) *EnableEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableEndpointResponseBody) SetStatus(v string) *EnableEndpointResponseBody {
	s.Status = &v
	return s
}

func (s *EnableEndpointResponseBody) SetSuccess(v bool) *EnableEndpointResponseBody {
	s.Success = &v
	return s
}

type EnableEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableEndpointResponse) GoString() string {
	return s.String()
}

func (s *EnableEndpointResponse) SetHeaders(v map[string]*string) *EnableEndpointResponse {
	s.Headers = v
	return s
}

func (s *EnableEndpointResponse) SetStatusCode(v int32) *EnableEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableEndpointResponse) SetBody(v *EnableEndpointResponseBody) *EnableEndpointResponse {
	s.Body = v
	return s
}

type GetEndpointAttributeRequest struct {
	// The type of the endpoint. Value:
	//
	// 	- **public**: indicates public endpoint. (Only the public is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s GetEndpointAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEndpointAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetEndpointAttributeRequest) SetEndpointType(v string) *GetEndpointAttributeRequest {
	s.EndpointType = &v
	return s
}

type GetEndpointAttributeResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *GetEndpointAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEndpointAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEndpointAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetEndpointAttributeResponseBody) SetCode(v int64) *GetEndpointAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *GetEndpointAttributeResponseBody) SetData(v *GetEndpointAttributeResponseBodyData) *GetEndpointAttributeResponseBody {
	s.Data = v
	return s
}

func (s *GetEndpointAttributeResponseBody) SetMessage(v string) *GetEndpointAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *GetEndpointAttributeResponseBody) SetRequestId(v string) *GetEndpointAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEndpointAttributeResponseBody) SetStatus(v string) *GetEndpointAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *GetEndpointAttributeResponseBody) SetSuccess(v bool) *GetEndpointAttributeResponseBody {
	s.Success = &v
	return s
}

type GetEndpointAttributeResponseBodyData struct {
	// The list of CIDR block.
	CidrList []*GetEndpointAttributeResponseBodyDataCidrList `json:"CidrList,omitempty" xml:"CidrList,omitempty" type:"Repeated"`
	// Specifies whether the endpoint is enabled.
	//
	// example:
	//
	// true
	EndpointEnabled *bool `json:"EndpointEnabled,omitempty" xml:"EndpointEnabled,omitempty"`
}

func (s GetEndpointAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEndpointAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEndpointAttributeResponseBodyData) SetCidrList(v []*GetEndpointAttributeResponseBodyDataCidrList) *GetEndpointAttributeResponseBodyData {
	s.CidrList = v
	return s
}

func (s *GetEndpointAttributeResponseBodyData) SetEndpointEnabled(v bool) *GetEndpointAttributeResponseBodyData {
	s.EndpointEnabled = &v
	return s
}

type GetEndpointAttributeResponseBodyDataCidrList struct {
	// The ACL policy. Valid values:
	//
	// 	- **allow**: indicates that the current endpoint allows access from the corresponding CIDR block. (Only allow is supported.)
	//
	// example:
	//
	// allow
	AclStrategy *string `json:"AclStrategy,omitempty" xml:"AclStrategy,omitempty"`
	// The CIDR block.
	//
	// example:
	//
	// 172.18.0.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1701951224000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s GetEndpointAttributeResponseBodyDataCidrList) String() string {
	return tea.Prettify(s)
}

func (s GetEndpointAttributeResponseBodyDataCidrList) GoString() string {
	return s.String()
}

func (s *GetEndpointAttributeResponseBodyDataCidrList) SetAclStrategy(v string) *GetEndpointAttributeResponseBodyDataCidrList {
	s.AclStrategy = &v
	return s
}

func (s *GetEndpointAttributeResponseBodyDataCidrList) SetCidr(v string) *GetEndpointAttributeResponseBodyDataCidrList {
	s.Cidr = &v
	return s
}

func (s *GetEndpointAttributeResponseBodyDataCidrList) SetCreateTime(v int64) *GetEndpointAttributeResponseBodyDataCidrList {
	s.CreateTime = &v
	return s
}

type GetEndpointAttributeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEndpointAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEndpointAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEndpointAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetEndpointAttributeResponse) SetHeaders(v map[string]*string) *GetEndpointAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetEndpointAttributeResponse) SetStatusCode(v int32) *GetEndpointAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEndpointAttributeResponse) SetBody(v *GetEndpointAttributeResponseBody) *GetEndpointAttributeResponse {
	s.Body = v
	return s
}

type GetQueueAttributesRequest struct {
	// The name of the queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo-queue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The tags.
	Tag []*GetQueueAttributesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetQueueAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQueueAttributesRequest) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesRequest) SetQueueName(v string) *GetQueueAttributesRequest {
	s.QueueName = &v
	return s
}

func (s *GetQueueAttributesRequest) SetTag(v []*GetQueueAttributesRequestTag) *GetQueueAttributesRequest {
	s.Tag = v
	return s
}

type GetQueueAttributesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetQueueAttributesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s GetQueueAttributesRequestTag) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesRequestTag) SetKey(v string) *GetQueueAttributesRequestTag {
	s.Key = &v
	return s
}

func (s *GetQueueAttributesRequestTag) SetValue(v string) *GetQueueAttributesRequestTag {
	s.Value = &v
	return s
}

type GetQueueAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetQueueAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueueAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQueueAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponseBody) SetCode(v int64) *GetQueueAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueueAttributesResponseBody) SetData(v *GetQueueAttributesResponseBodyData) *GetQueueAttributesResponseBody {
	s.Data = v
	return s
}

func (s *GetQueueAttributesResponseBody) SetMessage(v string) *GetQueueAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueueAttributesResponseBody) SetRequestId(v string) *GetQueueAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueueAttributesResponseBody) SetStatus(v string) *GetQueueAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *GetQueueAttributesResponseBody) SetSuccess(v bool) *GetQueueAttributesResponseBody {
	s.Success = &v
	return s
}

type GetQueueAttributesResponseBodyData struct {
	// The total number of messages that are in the Active state in the queue. The value is an approximate value. Default value: 0. We recommend that you do not use the return value and that you call CloudMonitor API operations to query the metric value.
	//
	// example:
	//
	// 20
	ActiveMessages *int64 `json:"ActiveMessages,omitempty" xml:"ActiveMessages,omitempty"`
	// The time when the queue was created.
	//
	// example:
	//
	// 1250700999
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The total number of messages that are in the Delayed state in the queue. The value is an approximate value. Default value: 0. We recommend that you do not use the return value and that you call CloudMonitor API operations to query the metric value.
	//
	// example:
	//
	// 0
	DelayMessages *int64 `json:"DelayMessages,omitempty" xml:"DelayMessages,omitempty"`
	// The period after which all messages sent to the queue are consumed. Unit: seconds.
	//
	// example:
	//
	// 30
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The dead-letter queue policy.
	DlqPolicy *GetQueueAttributesResponseBodyDataDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	// The total number of messages that are in the Inactive state in the queue. The value is an approximate value. Default value: 0. We recommend that you do not use the return value and that you call CloudMonitor API operations to query the metric value.
	//
	// example:
	//
	// 0
	InactiveMessages *int64 `json:"InactiveMessages,omitempty" xml:"InactiveMessages,omitempty"`
	// The time when the queue was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1250700999
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// Indicates whether the logging feature is enabled. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	LoggingEnabled *bool `json:"LoggingEnabled,omitempty" xml:"LoggingEnabled,omitempty"`
	// The maximum length of the message that is sent to the queue. Unit: bytes.
	//
	// example:
	//
	// 65536
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is received. Unit: seconds.
	//
	// example:
	//
	// 65536
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Unit: seconds.
	//
	// example:
	//
	// 0
	PollingWaitSeconds *int64 `json:"PollingWaitSeconds,omitempty" xml:"PollingWaitSeconds,omitempty"`
	// The name of the queue.
	//
	// example:
	//
	// demo-queue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The tag.
	Tags []*GetQueueAttributesResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 60
	VisibilityTimeout *int64 `json:"VisibilityTimeout,omitempty" xml:"VisibilityTimeout,omitempty"`
}

func (s GetQueueAttributesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQueueAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponseBodyData) SetActiveMessages(v int64) *GetQueueAttributesResponseBodyData {
	s.ActiveMessages = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetCreateTime(v int64) *GetQueueAttributesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetDelayMessages(v int64) *GetQueueAttributesResponseBodyData {
	s.DelayMessages = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetDelaySeconds(v int64) *GetQueueAttributesResponseBodyData {
	s.DelaySeconds = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetDlqPolicy(v *GetQueueAttributesResponseBodyDataDlqPolicy) *GetQueueAttributesResponseBodyData {
	s.DlqPolicy = v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetInactiveMessages(v int64) *GetQueueAttributesResponseBodyData {
	s.InactiveMessages = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetLastModifyTime(v int64) *GetQueueAttributesResponseBodyData {
	s.LastModifyTime = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetLoggingEnabled(v bool) *GetQueueAttributesResponseBodyData {
	s.LoggingEnabled = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetMaximumMessageSize(v int64) *GetQueueAttributesResponseBodyData {
	s.MaximumMessageSize = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetMessageRetentionPeriod(v int64) *GetQueueAttributesResponseBodyData {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetPollingWaitSeconds(v int64) *GetQueueAttributesResponseBodyData {
	s.PollingWaitSeconds = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetQueueName(v string) *GetQueueAttributesResponseBodyData {
	s.QueueName = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetTags(v []*GetQueueAttributesResponseBodyDataTags) *GetQueueAttributesResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetVisibilityTimeout(v int64) *GetQueueAttributesResponseBodyData {
	s.VisibilityTimeout = &v
	return s
}

type GetQueueAttributesResponseBodyDataDlqPolicy struct {
	// The queue to which dead-letter messages are delivered.
	//
	// example:
	//
	// deadLetterTargetQueue
	DeadLetterTargetQueue *string `json:"DeadLetterTargetQueue,omitempty" xml:"DeadLetterTargetQueue,omitempty"`
	// Specifies whether to enable the dead-letter message delivery.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 3
	MaxReceiveCount *string `json:"MaxReceiveCount,omitempty" xml:"MaxReceiveCount,omitempty"`
}

func (s GetQueueAttributesResponseBodyDataDlqPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetQueueAttributesResponseBodyDataDlqPolicy) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponseBodyDataDlqPolicy) SetDeadLetterTargetQueue(v string) *GetQueueAttributesResponseBodyDataDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *GetQueueAttributesResponseBodyDataDlqPolicy) SetEnabled(v bool) *GetQueueAttributesResponseBodyDataDlqPolicy {
	s.Enabled = &v
	return s
}

func (s *GetQueueAttributesResponseBodyDataDlqPolicy) SetMaxReceiveCount(v string) *GetQueueAttributesResponseBodyDataDlqPolicy {
	s.MaxReceiveCount = &v
	return s
}

type GetQueueAttributesResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetQueueAttributesResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s GetQueueAttributesResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponseBodyDataTags) SetTagKey(v string) *GetQueueAttributesResponseBodyDataTags {
	s.TagKey = &v
	return s
}

func (s *GetQueueAttributesResponseBodyDataTags) SetTagValue(v string) *GetQueueAttributesResponseBodyDataTags {
	s.TagValue = &v
	return s
}

type GetQueueAttributesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueueAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueueAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQueueAttributesResponse) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponse) SetHeaders(v map[string]*string) *GetQueueAttributesResponse {
	s.Headers = v
	return s
}

func (s *GetQueueAttributesResponse) SetStatusCode(v int32) *GetQueueAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueueAttributesResponse) SetBody(v *GetQueueAttributesResponseBody) *GetQueueAttributesResponse {
	s.Body = v
	return s
}

type GetSubscriptionAttributesRequest struct {
	// The name of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyTopic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetSubscriptionAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionAttributesRequest) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesRequest) SetSubscriptionName(v string) *GetSubscriptionAttributesRequest {
	s.SubscriptionName = &v
	return s
}

func (s *GetSubscriptionAttributesRequest) SetTopicName(v string) *GetSubscriptionAttributesRequest {
	s.TopicName = &v
	return s
}

type GetSubscriptionAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetSubscriptionAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSubscriptionAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesResponseBody) SetCode(v int64) *GetSubscriptionAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetData(v *GetSubscriptionAttributesResponseBodyData) *GetSubscriptionAttributesResponseBody {
	s.Data = v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetMessage(v string) *GetSubscriptionAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetRequestId(v string) *GetSubscriptionAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetStatus(v string) *GetSubscriptionAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetSuccess(v bool) *GetSubscriptionAttributesResponseBody {
	s.Success = &v
	return s
}

type GetSubscriptionAttributesResponseBodyData struct {
	// The time when the subscription was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554806
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The dead-letter queue policy.
	DlqPolicy *GetSubscriptionAttributesResponseBodyDataDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	// The endpoint to which the messages are pushed.
	//
	// example:
	//
	// http://example.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The tag that is used to filter messages. Only the messages that are attached with the specified tag can be pushed.
	//
	// example:
	//
	// important
	FilterTag *string `json:"FilterTag,omitempty" xml:"FilterTag,omitempty"`
	// The time when the subscription was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554962
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The content format of the messages that are pushed to the endpoint. Valid values:
	//
	// 	- XML
	//
	// 	- JSON
	//
	// 	- SIMPLIFIED
	//
	// example:
	//
	// XML
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" xml:"NotifyContentFormat,omitempty"`
	// The retry policy that is applied if an error occurs when Message Service (MNS) pushes messages to the endpoint. Valid values:
	//
	// 	- BACKOFF_RETRY
	//
	// 	- EXPONENTIAL_DECAY_RETRY
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The name of the subscription.
	//
	// example:
	//
	// MySubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// MyTopic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The Alibaba Cloud account ID of the topic owner.
	//
	// example:
	//
	// 123456789098****
	TopicOwner *string `json:"TopicOwner,omitempty" xml:"TopicOwner,omitempty"`
}

func (s GetSubscriptionAttributesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesResponseBodyData) SetCreateTime(v int64) *GetSubscriptionAttributesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetDlqPolicy(v *GetSubscriptionAttributesResponseBodyDataDlqPolicy) *GetSubscriptionAttributesResponseBodyData {
	s.DlqPolicy = v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetEndpoint(v string) *GetSubscriptionAttributesResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetFilterTag(v string) *GetSubscriptionAttributesResponseBodyData {
	s.FilterTag = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetLastModifyTime(v int64) *GetSubscriptionAttributesResponseBodyData {
	s.LastModifyTime = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetNotifyContentFormat(v string) *GetSubscriptionAttributesResponseBodyData {
	s.NotifyContentFormat = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetNotifyStrategy(v string) *GetSubscriptionAttributesResponseBodyData {
	s.NotifyStrategy = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetSubscriptionName(v string) *GetSubscriptionAttributesResponseBodyData {
	s.SubscriptionName = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetTopicName(v string) *GetSubscriptionAttributesResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetTopicOwner(v string) *GetSubscriptionAttributesResponseBodyData {
	s.TopicOwner = &v
	return s
}

type GetSubscriptionAttributesResponseBodyDataDlqPolicy struct {
	// The queue to which dead-letter messages are delivered.
	//
	// example:
	//
	// deadLetterTargetQueue
	DeadLetterTargetQueue *string `json:"DeadLetterTargetQueue,omitempty" xml:"DeadLetterTargetQueue,omitempty"`
	// Specifies whether to enable the dead-letter message delivery.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s GetSubscriptionAttributesResponseBodyDataDlqPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionAttributesResponseBodyDataDlqPolicy) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesResponseBodyDataDlqPolicy) SetDeadLetterTargetQueue(v string) *GetSubscriptionAttributesResponseBodyDataDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyDataDlqPolicy) SetEnabled(v bool) *GetSubscriptionAttributesResponseBodyDataDlqPolicy {
	s.Enabled = &v
	return s
}

type GetSubscriptionAttributesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubscriptionAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubscriptionAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionAttributesResponse) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesResponse) SetHeaders(v map[string]*string) *GetSubscriptionAttributesResponse {
	s.Headers = v
	return s
}

func (s *GetSubscriptionAttributesResponse) SetStatusCode(v int32) *GetSubscriptionAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubscriptionAttributesResponse) SetBody(v *GetSubscriptionAttributesResponseBody) *GetSubscriptionAttributesResponse {
	s.Body = v
	return s
}

type GetTopicAttributesRequest struct {
	// The tag.
	Tag []*GetTopicAttributesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo-topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetTopicAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTopicAttributesRequest) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesRequest) SetTag(v []*GetTopicAttributesRequestTag) *GetTopicAttributesRequest {
	s.Tag = v
	return s
}

func (s *GetTopicAttributesRequest) SetTopicName(v string) *GetTopicAttributesRequest {
	s.TopicName = &v
	return s
}

type GetTopicAttributesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTopicAttributesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s GetTopicAttributesRequestTag) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesRequestTag) SetKey(v string) *GetTopicAttributesRequestTag {
	s.Key = &v
	return s
}

func (s *GetTopicAttributesRequestTag) SetValue(v string) *GetTopicAttributesRequestTag {
	s.Value = &v
	return s
}

type GetTopicAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetTopicAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTopicAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTopicAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesResponseBody) SetCode(v int64) *GetTopicAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicAttributesResponseBody) SetData(v *GetTopicAttributesResponseBodyData) *GetTopicAttributesResponseBody {
	s.Data = v
	return s
}

func (s *GetTopicAttributesResponseBody) SetMessage(v string) *GetTopicAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicAttributesResponseBody) SetRequestId(v string) *GetTopicAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicAttributesResponseBody) SetStatus(v string) *GetTopicAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *GetTopicAttributesResponseBody) SetSuccess(v bool) *GetTopicAttributesResponseBody {
	s.Success = &v
	return s
}

type GetTopicAttributesResponseBodyData struct {
	// The time when the topic was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554277
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the topic was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554460
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// Indicates whether the logging feature is enabled. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	LoggingEnabled *bool `json:"LoggingEnabled,omitempty" xml:"LoggingEnabled,omitempty"`
	// The maximum length of the message that is sent to the topic. Unit: bytes.
	//
	// example:
	//
	// 65536
	MaxMessageSize *int64 `json:"MaxMessageSize,omitempty" xml:"MaxMessageSize,omitempty"`
	// The number of messages in the topic.
	//
	// example:
	//
	// 0
	MessageCount *int64 `json:"MessageCount,omitempty" xml:"MessageCount,omitempty"`
	// The maximum duration for which a message is retained in the topic. After the specified retention period ends, the message is deleted regardless of whether the message is received. Unit: seconds.
	//
	// example:
	//
	// 86400
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The tags added to the resources.
	Tags []*GetTopicAttributesResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The name of the topic.
	//
	// example:
	//
	// demo-topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetTopicAttributesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTopicAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesResponseBodyData) SetCreateTime(v int64) *GetTopicAttributesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetLastModifyTime(v int64) *GetTopicAttributesResponseBodyData {
	s.LastModifyTime = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetLoggingEnabled(v bool) *GetTopicAttributesResponseBodyData {
	s.LoggingEnabled = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetMaxMessageSize(v int64) *GetTopicAttributesResponseBodyData {
	s.MaxMessageSize = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetMessageCount(v int64) *GetTopicAttributesResponseBodyData {
	s.MessageCount = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetMessageRetentionPeriod(v int64) *GetTopicAttributesResponseBodyData {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetTags(v []*GetTopicAttributesResponseBodyDataTags) *GetTopicAttributesResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetTopicName(v string) *GetTopicAttributesResponseBodyData {
	s.TopicName = &v
	return s
}

type GetTopicAttributesResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetTopicAttributesResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s GetTopicAttributesResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesResponseBodyDataTags) SetTagKey(v string) *GetTopicAttributesResponseBodyDataTags {
	s.TagKey = &v
	return s
}

func (s *GetTopicAttributesResponseBodyDataTags) SetTagValue(v string) *GetTopicAttributesResponseBodyDataTags {
	s.TagValue = &v
	return s
}

type GetTopicAttributesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTopicAttributesResponse) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesResponse) SetHeaders(v map[string]*string) *GetTopicAttributesResponse {
	s.Headers = v
	return s
}

func (s *GetTopicAttributesResponse) SetStatusCode(v int32) *GetTopicAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicAttributesResponse) SetBody(v *GetTopicAttributesResponseBody) *GetTopicAttributesResponse {
	s.Body = v
	return s
}

type ListQueueRequest struct {
	// The page number. Valid values: 1 to 100000000. If you set this parameter to a value smaller than 1, the value of this parameter is 1 by default. If you set this parameter to a value greater than 100000000, the value of this parameter is 100000000 by default.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Value values: 10 to 50. If you set this parameter to a value smaller than 10, the value of this parameter is 10 by default. If you set this parameter to a value greater than 50, the value of this parameter is 50 by default.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the queue.
	//
	// example:
	//
	// demo-queue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The tags.
	Tag []*ListQueueRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueueRequest) GoString() string {
	return s.String()
}

func (s *ListQueueRequest) SetPageNum(v int64) *ListQueueRequest {
	s.PageNum = &v
	return s
}

func (s *ListQueueRequest) SetPageSize(v int64) *ListQueueRequest {
	s.PageSize = &v
	return s
}

func (s *ListQueueRequest) SetQueueName(v string) *ListQueueRequest {
	s.QueueName = &v
	return s
}

func (s *ListQueueRequest) SetTag(v []*ListQueueRequestTag) *ListQueueRequest {
	s.Tag = v
	return s
}

type ListQueueRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQueueRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListQueueRequestTag) GoString() string {
	return s.String()
}

func (s *ListQueueRequestTag) SetKey(v string) *ListQueueRequestTag {
	s.Key = &v
	return s
}

func (s *ListQueueRequestTag) SetValue(v string) *ListQueueRequestTag {
	s.Value = &v
	return s
}

type ListQueueResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ListQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueueResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBody) SetCode(v int64) *ListQueueResponseBody {
	s.Code = &v
	return s
}

func (s *ListQueueResponseBody) SetData(v *ListQueueResponseBodyData) *ListQueueResponseBody {
	s.Data = v
	return s
}

func (s *ListQueueResponseBody) SetMessage(v string) *ListQueueResponseBody {
	s.Message = &v
	return s
}

func (s *ListQueueResponseBody) SetRequestId(v string) *ListQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueueResponseBody) SetStatus(v string) *ListQueueResponseBody {
	s.Status = &v
	return s
}

func (s *ListQueueResponseBody) SetSuccess(v bool) *ListQueueResponseBody {
	s.Success = &v
	return s
}

type ListQueueResponseBodyData struct {
	// The data returned on the current page.
	PageData []*ListQueueResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 3
	Pages *int64 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 130
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListQueueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBodyData) SetPageData(v []*ListQueueResponseBodyDataPageData) *ListQueueResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListQueueResponseBodyData) SetPageNum(v int64) *ListQueueResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListQueueResponseBodyData) SetPageSize(v int64) *ListQueueResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListQueueResponseBodyData) SetPages(v int64) *ListQueueResponseBodyData {
	s.Pages = &v
	return s
}

func (s *ListQueueResponseBodyData) SetSize(v int64) *ListQueueResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListQueueResponseBodyData) SetTotal(v int64) *ListQueueResponseBodyData {
	s.Total = &v
	return s
}

type ListQueueResponseBodyDataPageData struct {
	// The total number of messages that are in the Active state in the queue. The value is an approximate number. Default value: 0. We recommend that you do not use the return value and that you call CloudMonitor API operations to query the metric value.
	//
	// example:
	//
	// 20
	ActiveMessages *int64 `json:"ActiveMessages,omitempty" xml:"ActiveMessages,omitempty"`
	// The time when the queue was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1250700999
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The total number of the messages that are in the Delayed state in the queue. The value is an approximate number. Default value: 0. We recommend that you do not use the return value and that you call CloudMonitor API operations to query the metric value.
	//
	// example:
	//
	// 0
	DelayMessages *int64 `json:"DelayMessages,omitempty" xml:"DelayMessages,omitempty"`
	// The period after which all messages sent to the queue are consumed. Unit: seconds.
	//
	// example:
	//
	// 30
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The dead-letter queue policy.
	DlqPolicy *ListQueueResponseBodyDataPageDataDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	// The total number of the messages that are in the Inactive state in the queue. The value is an approximate number. Default value: 0. We recommend that you do not use the return value and that you call CloudMonitor API operations to query the metric value.
	//
	// example:
	//
	// 0
	InactiveMessages *int64 `json:"InactiveMessages,omitempty" xml:"InactiveMessages,omitempty"`
	// The time when the queue was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1250700999
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// Indicates whether the logging feature is enabled. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	LoggingEnabled *bool `json:"LoggingEnabled,omitempty" xml:"LoggingEnabled,omitempty"`
	// The maximum length of the message that is sent to the queue. Unit: bytes.
	//
	// example:
	//
	// 65536
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is received. Unit: seconds.
	//
	// example:
	//
	// 65536
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Unit: seconds.
	//
	// example:
	//
	// 0
	PollingWaitSeconds *int64 `json:"PollingWaitSeconds,omitempty" xml:"PollingWaitSeconds,omitempty"`
	// The name of the queue.
	//
	// example:
	//
	// demo-queue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The tags added to the resources.
	Tags []*ListQueueResponseBodyDataPageDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 60
	VisibilityTimeout *int64 `json:"VisibilityTimeout,omitempty" xml:"VisibilityTimeout,omitempty"`
}

func (s ListQueueResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListQueueResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBodyDataPageData) SetActiveMessages(v int64) *ListQueueResponseBodyDataPageData {
	s.ActiveMessages = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetCreateTime(v int64) *ListQueueResponseBodyDataPageData {
	s.CreateTime = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetDelayMessages(v int64) *ListQueueResponseBodyDataPageData {
	s.DelayMessages = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetDelaySeconds(v int64) *ListQueueResponseBodyDataPageData {
	s.DelaySeconds = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetDlqPolicy(v *ListQueueResponseBodyDataPageDataDlqPolicy) *ListQueueResponseBodyDataPageData {
	s.DlqPolicy = v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetInactiveMessages(v int64) *ListQueueResponseBodyDataPageData {
	s.InactiveMessages = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetLastModifyTime(v int64) *ListQueueResponseBodyDataPageData {
	s.LastModifyTime = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetLoggingEnabled(v bool) *ListQueueResponseBodyDataPageData {
	s.LoggingEnabled = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetMaximumMessageSize(v int64) *ListQueueResponseBodyDataPageData {
	s.MaximumMessageSize = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetMessageRetentionPeriod(v int64) *ListQueueResponseBodyDataPageData {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetPollingWaitSeconds(v int64) *ListQueueResponseBodyDataPageData {
	s.PollingWaitSeconds = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetQueueName(v string) *ListQueueResponseBodyDataPageData {
	s.QueueName = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetTags(v []*ListQueueResponseBodyDataPageDataTags) *ListQueueResponseBodyDataPageData {
	s.Tags = v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetVisibilityTimeout(v int64) *ListQueueResponseBodyDataPageData {
	s.VisibilityTimeout = &v
	return s
}

type ListQueueResponseBodyDataPageDataDlqPolicy struct {
	// The queue to which dead-letter messages are delivered.
	//
	// example:
	//
	// deadLetterTargetQueue
	DeadLetterTargetQueue *string `json:"DeadLetterTargetQueue,omitempty" xml:"DeadLetterTargetQueue,omitempty"`
	// Specifies whether to enable the dead-letter message delivery.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 3
	MaxReceiveCount *string `json:"MaxReceiveCount,omitempty" xml:"MaxReceiveCount,omitempty"`
}

func (s ListQueueResponseBodyDataPageDataDlqPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListQueueResponseBodyDataPageDataDlqPolicy) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBodyDataPageDataDlqPolicy) SetDeadLetterTargetQueue(v string) *ListQueueResponseBodyDataPageDataDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *ListQueueResponseBodyDataPageDataDlqPolicy) SetEnabled(v bool) *ListQueueResponseBodyDataPageDataDlqPolicy {
	s.Enabled = &v
	return s
}

func (s *ListQueueResponseBodyDataPageDataDlqPolicy) SetMaxReceiveCount(v string) *ListQueueResponseBodyDataPageDataDlqPolicy {
	s.MaxReceiveCount = &v
	return s
}

type ListQueueResponseBodyDataPageDataTags struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListQueueResponseBodyDataPageDataTags) String() string {
	return tea.Prettify(s)
}

func (s ListQueueResponseBodyDataPageDataTags) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBodyDataPageDataTags) SetTagKey(v string) *ListQueueResponseBodyDataPageDataTags {
	s.TagKey = &v
	return s
}

func (s *ListQueueResponseBodyDataPageDataTags) SetTagValue(v string) *ListQueueResponseBodyDataPageDataTags {
	s.TagValue = &v
	return s
}

type ListQueueResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueueResponse) GoString() string {
	return s.String()
}

func (s *ListQueueResponse) SetHeaders(v map[string]*string) *ListQueueResponse {
	s.Headers = v
	return s
}

func (s *ListQueueResponse) SetStatusCode(v int32) *ListQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueueResponse) SetBody(v *ListQueueResponseBody) *ListQueueResponse {
	s.Body = v
	return s
}

type ListSubscriptionByTopicRequest struct {
	EndpointType  *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	EndpointValue *string `json:"EndpointValue,omitempty" xml:"EndpointValue,omitempty"`
	// The page number. Valid values: 1 to 100000000. If you set this parameter to a value smaller than 1, the value of this parameter is 1 by default. If you set this parameter to a value greater than 100000000, the value of this parameter is 100000000 by default.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Value values: 10 to 50. If you set this parameter to a value smaller than 10, the value of this parameter is 10 by default. If you set this parameter to a value greater than 50, the value of this parameter is 50 by default.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the subscription.
	//
	// example:
	//
	// demo-subscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The topic name.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s ListSubscriptionByTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionByTopicRequest) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicRequest) SetEndpointType(v string) *ListSubscriptionByTopicRequest {
	s.EndpointType = &v
	return s
}

func (s *ListSubscriptionByTopicRequest) SetEndpointValue(v string) *ListSubscriptionByTopicRequest {
	s.EndpointValue = &v
	return s
}

func (s *ListSubscriptionByTopicRequest) SetPageNum(v int64) *ListSubscriptionByTopicRequest {
	s.PageNum = &v
	return s
}

func (s *ListSubscriptionByTopicRequest) SetPageSize(v int64) *ListSubscriptionByTopicRequest {
	s.PageSize = &v
	return s
}

func (s *ListSubscriptionByTopicRequest) SetSubscriptionName(v string) *ListSubscriptionByTopicRequest {
	s.SubscriptionName = &v
	return s
}

func (s *ListSubscriptionByTopicRequest) SetTopicName(v string) *ListSubscriptionByTopicRequest {
	s.TopicName = &v
	return s
}

type ListSubscriptionByTopicResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ListSubscriptionByTopicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSubscriptionByTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionByTopicResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicResponseBody) SetCode(v int64) *ListSubscriptionByTopicResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetData(v *ListSubscriptionByTopicResponseBodyData) *ListSubscriptionByTopicResponseBody {
	s.Data = v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetMessage(v string) *ListSubscriptionByTopicResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetRequestId(v string) *ListSubscriptionByTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetStatus(v string) *ListSubscriptionByTopicResponseBody {
	s.Status = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetSuccess(v bool) *ListSubscriptionByTopicResponseBody {
	s.Success = &v
	return s
}

type ListSubscriptionByTopicResponseBodyData struct {
	// The data returned on the current page.
	PageData []*ListSubscriptionByTopicResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 3
	Pages *int64 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 130
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSubscriptionByTopicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionByTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicResponseBodyData) SetPageData(v []*ListSubscriptionByTopicResponseBodyDataPageData) *ListSubscriptionByTopicResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetPageNum(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetPageSize(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetPages(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.Pages = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetSize(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetTotal(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.Total = &v
	return s
}

type ListSubscriptionByTopicResponseBodyDataPageData struct {
	// The time when the subscription was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554806
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The dead-letter queue policy.
	DlqPolicy *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	// The endpoint to which the messages are pushed.
	//
	// example:
	//
	// http://example.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The tag that is used to filter messages. Only the messages that are attached with the specified tag can be pushed.
	//
	// example:
	//
	// important
	FilterTag *string `json:"FilterTag,omitempty" xml:"FilterTag,omitempty"`
	// The time when the subscription was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554806
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The content format of the messages that are pushed to the endpoint. Valid values:
	//
	// 	- XML
	//
	// 	- JSON
	//
	// 	- SIMPLIFIED
	//
	// example:
	//
	// XML
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" xml:"NotifyContentFormat,omitempty"`
	// The retry policy that is applied if an error occurs when Message Service (MNS) pushes messages to the endpoint. Valid values:
	//
	// 	- BACKOFF_RETRY
	//
	// 	- EXPONENTIAL_DECAY_RETRY
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The name of the subscription.
	//
	// example:
	//
	// MySubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// MyTopic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The Alibaba Cloud account ID of the topic owner.
	//
	// example:
	//
	// 123456789098****
	TopicOwner *string `json:"TopicOwner,omitempty" xml:"TopicOwner,omitempty"`
}

func (s ListSubscriptionByTopicResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionByTopicResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetCreateTime(v int64) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.CreateTime = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetDlqPolicy(v *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.DlqPolicy = v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetEndpoint(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.Endpoint = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetFilterTag(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.FilterTag = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetLastModifyTime(v int64) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.LastModifyTime = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetNotifyContentFormat(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.NotifyContentFormat = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetNotifyStrategy(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.NotifyStrategy = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetSubscriptionName(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.SubscriptionName = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetTopicName(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.TopicName = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetTopicOwner(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.TopicOwner = &v
	return s
}

type ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy struct {
	// The queue to which dead-letter messages are delivered.
	//
	// example:
	//
	// dead-letter-queue
	DeadLetterTargetQueue *string `json:"DeadLetterTargetQueue,omitempty" xml:"DeadLetterTargetQueue,omitempty"`
	// Specifies whether to enable the dead-letter message delivery.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy) SetDeadLetterTargetQueue(v string) *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy) SetEnabled(v bool) *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy {
	s.Enabled = &v
	return s
}

type ListSubscriptionByTopicResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubscriptionByTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubscriptionByTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionByTopicResponse) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicResponse) SetHeaders(v map[string]*string) *ListSubscriptionByTopicResponse {
	s.Headers = v
	return s
}

func (s *ListSubscriptionByTopicResponse) SetStatusCode(v int32) *ListSubscriptionByTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubscriptionByTopicResponse) SetBody(v *ListSubscriptionByTopicResponseBody) *ListSubscriptionByTopicResponse {
	s.Body = v
	return s
}

type ListTopicRequest struct {
	// The page number. Valid values: 1 to 100000000. If you set this parameter to a value smaller than 1, the value of this parameter is 1 by default. If you set this parameter to a value greater than 100000000, the value of this parameter is 100000000 by default.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Value values: 10 to 50. If you set this parameter to a value smaller than 10, the value of this parameter is 10 by default. If you set this parameter to a value greater than 50, the value of this parameter is 50 by default.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The tags.
	Tag []*ListTopicRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the topic.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s ListTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTopicRequest) GoString() string {
	return s.String()
}

func (s *ListTopicRequest) SetPageNum(v int64) *ListTopicRequest {
	s.PageNum = &v
	return s
}

func (s *ListTopicRequest) SetPageSize(v int64) *ListTopicRequest {
	s.PageSize = &v
	return s
}

func (s *ListTopicRequest) SetTag(v []*ListTopicRequestTag) *ListTopicRequest {
	s.Tag = v
	return s
}

func (s *ListTopicRequest) SetTopicName(v string) *ListTopicRequest {
	s.TopicName = &v
	return s
}

type ListTopicRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTopicRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTopicRequestTag) GoString() string {
	return s.String()
}

func (s *ListTopicRequestTag) SetKey(v string) *ListTopicRequestTag {
	s.Key = &v
	return s
}

func (s *ListTopicRequestTag) SetValue(v string) *ListTopicRequestTag {
	s.Value = &v
	return s
}

type ListTopicResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ListTopicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTopicResponseBody) GoString() string {
	return s.String()
}

func (s *ListTopicResponseBody) SetCode(v int64) *ListTopicResponseBody {
	s.Code = &v
	return s
}

func (s *ListTopicResponseBody) SetData(v *ListTopicResponseBodyData) *ListTopicResponseBody {
	s.Data = v
	return s
}

func (s *ListTopicResponseBody) SetMessage(v string) *ListTopicResponseBody {
	s.Message = &v
	return s
}

func (s *ListTopicResponseBody) SetRequestId(v string) *ListTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTopicResponseBody) SetStatus(v string) *ListTopicResponseBody {
	s.Status = &v
	return s
}

func (s *ListTopicResponseBody) SetSuccess(v bool) *ListTopicResponseBody {
	s.Success = &v
	return s
}

type ListTopicResponseBodyData struct {
	// The data returned on the current page.
	PageData []*ListTopicResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 130
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTopicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTopicResponseBodyData) SetPageData(v []*ListTopicResponseBodyDataPageData) *ListTopicResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListTopicResponseBodyData) SetPageNum(v int64) *ListTopicResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListTopicResponseBodyData) SetPageSize(v int64) *ListTopicResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTopicResponseBodyData) SetTotal(v int64) *ListTopicResponseBodyData {
	s.Total = &v
	return s
}

type ListTopicResponseBodyDataPageData struct {
	// The time when the subscription was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554962
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the subscription was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554962
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// Indicates whether the logging feature is enabled.
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	LoggingEnabled *bool `json:"LoggingEnabled,omitempty" xml:"LoggingEnabled,omitempty"`
	// The maximum length of the message that is sent to the topic. Unit: bytes.
	//
	// example:
	//
	// 65536
	MaxMessageSize *int64 `json:"MaxMessageSize,omitempty" xml:"MaxMessageSize,omitempty"`
	// The number of messages in the topic.
	//
	// example:
	//
	// 0
	MessageCount *int64 `json:"MessageCount,omitempty" xml:"MessageCount,omitempty"`
	// The maximum duration for which a message is retained in the topic. After the specified retention period ends, the message is deleted regardless of whether the message is received. Unit: seconds.
	//
	// example:
	//
	// 86400
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The tags added to the resources.
	Tags []*ListTopicResponseBodyDataPageDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The internal URL of the message topic. The internal URL can be accessed over an internal network.
	//
	// example:
	//
	// http:// 111111111****.mns.us-west-1-internal.aliyuncs.com/topics/testTopic
	TopicInnerUrl *string `json:"TopicInnerUrl,omitempty" xml:"TopicInnerUrl,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// demo-topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The URL of the message topic.
	//
	// example:
	//
	// http:// 111111111****.mns.us-west-1.aliyuncs.com/topics/testTopic
	TopicUrl *string `json:"TopicUrl,omitempty" xml:"TopicUrl,omitempty"`
}

func (s ListTopicResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListTopicResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListTopicResponseBodyDataPageData) SetCreateTime(v int64) *ListTopicResponseBodyDataPageData {
	s.CreateTime = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetLastModifyTime(v int64) *ListTopicResponseBodyDataPageData {
	s.LastModifyTime = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetLoggingEnabled(v bool) *ListTopicResponseBodyDataPageData {
	s.LoggingEnabled = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetMaxMessageSize(v int64) *ListTopicResponseBodyDataPageData {
	s.MaxMessageSize = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetMessageCount(v int64) *ListTopicResponseBodyDataPageData {
	s.MessageCount = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetMessageRetentionPeriod(v int64) *ListTopicResponseBodyDataPageData {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetTags(v []*ListTopicResponseBodyDataPageDataTags) *ListTopicResponseBodyDataPageData {
	s.Tags = v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetTopicInnerUrl(v string) *ListTopicResponseBodyDataPageData {
	s.TopicInnerUrl = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetTopicName(v string) *ListTopicResponseBodyDataPageData {
	s.TopicName = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetTopicUrl(v string) *ListTopicResponseBodyDataPageData {
	s.TopicUrl = &v
	return s
}

type ListTopicResponseBodyDataPageDataTags struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTopicResponseBodyDataPageDataTags) String() string {
	return tea.Prettify(s)
}

func (s ListTopicResponseBodyDataPageDataTags) GoString() string {
	return s.String()
}

func (s *ListTopicResponseBodyDataPageDataTags) SetTagKey(v string) *ListTopicResponseBodyDataPageDataTags {
	s.TagKey = &v
	return s
}

func (s *ListTopicResponseBodyDataPageDataTags) SetTagValue(v string) *ListTopicResponseBodyDataPageDataTags {
	s.TagValue = &v
	return s
}

type ListTopicResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTopicResponse) GoString() string {
	return s.String()
}

func (s *ListTopicResponse) SetHeaders(v map[string]*string) *ListTopicResponse {
	s.Headers = v
	return s
}

func (s *ListTopicResponse) SetStatusCode(v int32) *ListTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTopicResponse) SetBody(v *ListTopicResponseBody) *ListTopicResponse {
	s.Body = v
	return s
}

type RevokeEndpointAclRequest struct {
	// The ACL policy. Value:
	//
	// 	- **allow**: indicates that this operation is included in the Cidr whitelist. (Only the allow is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// allow
	AclStrategy *string `json:"AclStrategy,omitempty" xml:"AclStrategy,omitempty"`
	// The list of CIDR block.
	//
	// This parameter is required.
	CidrList []*string `json:"CidrList,omitempty" xml:"CidrList,omitempty" type:"Repeated"`
	// The type of the endpoint. Valid values:
	//
	// 	- **public**: indicates public endpoint. (Only the public is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s RevokeEndpointAclRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeEndpointAclRequest) GoString() string {
	return s.String()
}

func (s *RevokeEndpointAclRequest) SetAclStrategy(v string) *RevokeEndpointAclRequest {
	s.AclStrategy = &v
	return s
}

func (s *RevokeEndpointAclRequest) SetCidrList(v []*string) *RevokeEndpointAclRequest {
	s.CidrList = v
	return s
}

func (s *RevokeEndpointAclRequest) SetEndpointType(v string) *RevokeEndpointAclRequest {
	s.EndpointType = &v
	return s
}

type RevokeEndpointAclShrinkRequest struct {
	// The ACL policy. Value:
	//
	// 	- **allow**: indicates that this operation is included in the Cidr whitelist. (Only the allow is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// allow
	AclStrategy *string `json:"AclStrategy,omitempty" xml:"AclStrategy,omitempty"`
	// The list of CIDR block.
	//
	// This parameter is required.
	CidrListShrink *string `json:"CidrList,omitempty" xml:"CidrList,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **public**: indicates public endpoint. (Only the public is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s RevokeEndpointAclShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeEndpointAclShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevokeEndpointAclShrinkRequest) SetAclStrategy(v string) *RevokeEndpointAclShrinkRequest {
	s.AclStrategy = &v
	return s
}

func (s *RevokeEndpointAclShrinkRequest) SetCidrListShrink(v string) *RevokeEndpointAclShrinkRequest {
	s.CidrListShrink = &v
	return s
}

func (s *RevokeEndpointAclShrinkRequest) SetEndpointType(v string) *RevokeEndpointAclShrinkRequest {
	s.EndpointType = &v
	return s
}

type RevokeEndpointAclResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokeEndpointAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeEndpointAclResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeEndpointAclResponseBody) SetCode(v int64) *RevokeEndpointAclResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeEndpointAclResponseBody) SetMessage(v string) *RevokeEndpointAclResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeEndpointAclResponseBody) SetRequestId(v string) *RevokeEndpointAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeEndpointAclResponseBody) SetStatus(v string) *RevokeEndpointAclResponseBody {
	s.Status = &v
	return s
}

func (s *RevokeEndpointAclResponseBody) SetSuccess(v bool) *RevokeEndpointAclResponseBody {
	s.Success = &v
	return s
}

type RevokeEndpointAclResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeEndpointAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeEndpointAclResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeEndpointAclResponse) GoString() string {
	return s.String()
}

func (s *RevokeEndpointAclResponse) SetHeaders(v map[string]*string) *RevokeEndpointAclResponse {
	s.Headers = v
	return s
}

func (s *RevokeEndpointAclResponse) SetStatusCode(v int32) *RevokeEndpointAclResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeEndpointAclResponse) SetBody(v *RevokeEndpointAclResponseBody) *RevokeEndpointAclResponse {
	s.Body = v
	return s
}

type SetQueueAttributesRequest struct {
	// The period after which all messages sent to the queue are consumed. Valid values: 0 to 604800. Unit: seconds. Default value: 0
	//
	// example:
	//
	// 0
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The dead-letter queue policy.
	DlqPolicy *SetQueueAttributesRequestDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	// Specifies whether to enable the log management feature. Valid values:
	//
	// 	- true: enabled.
	//
	// 	- false: disabled. Default value: false.
	//
	// example:
	//
	// True
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// The maximum length of the message that is sent to the queue. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 1024
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is received. Valid values: 60 to 604800. Unit: seconds. Default value: 345600.
	//
	// example:
	//
	// 120
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Valid values: 0 to 30. Unit: seconds. Default value: 0
	//
	// example:
	//
	// 0
	PollingWaitSeconds *int64 `json:"PollingWaitSeconds,omitempty" xml:"PollingWaitSeconds,omitempty"`
	// The name of the queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// testqueue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 60
	VisibilityTimeout *int64 `json:"VisibilityTimeout,omitempty" xml:"VisibilityTimeout,omitempty"`
}

func (s SetQueueAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s SetQueueAttributesRequest) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesRequest) SetDelaySeconds(v int64) *SetQueueAttributesRequest {
	s.DelaySeconds = &v
	return s
}

func (s *SetQueueAttributesRequest) SetDlqPolicy(v *SetQueueAttributesRequestDlqPolicy) *SetQueueAttributesRequest {
	s.DlqPolicy = v
	return s
}

func (s *SetQueueAttributesRequest) SetEnableLogging(v bool) *SetQueueAttributesRequest {
	s.EnableLogging = &v
	return s
}

func (s *SetQueueAttributesRequest) SetMaximumMessageSize(v int64) *SetQueueAttributesRequest {
	s.MaximumMessageSize = &v
	return s
}

func (s *SetQueueAttributesRequest) SetMessageRetentionPeriod(v int64) *SetQueueAttributesRequest {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *SetQueueAttributesRequest) SetPollingWaitSeconds(v int64) *SetQueueAttributesRequest {
	s.PollingWaitSeconds = &v
	return s
}

func (s *SetQueueAttributesRequest) SetQueueName(v string) *SetQueueAttributesRequest {
	s.QueueName = &v
	return s
}

func (s *SetQueueAttributesRequest) SetVisibilityTimeout(v int64) *SetQueueAttributesRequest {
	s.VisibilityTimeout = &v
	return s
}

type SetQueueAttributesRequestDlqPolicy struct {
	// The queue to which dead-letter messages are delivered.
	//
	// example:
	//
	// deadLetterTargetQueue
	DeadLetterTargetQueue *string `json:"DeadLetterTargetQueue,omitempty" xml:"DeadLetterTargetQueue,omitempty"`
	// Specifies whether to enable the dead-letter message delivery.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 3
	MaxReceiveCount *int32 `json:"MaxReceiveCount,omitempty" xml:"MaxReceiveCount,omitempty"`
}

func (s SetQueueAttributesRequestDlqPolicy) String() string {
	return tea.Prettify(s)
}

func (s SetQueueAttributesRequestDlqPolicy) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesRequestDlqPolicy) SetDeadLetterTargetQueue(v string) *SetQueueAttributesRequestDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *SetQueueAttributesRequestDlqPolicy) SetEnabled(v bool) *SetQueueAttributesRequestDlqPolicy {
	s.Enabled = &v
	return s
}

func (s *SetQueueAttributesRequestDlqPolicy) SetMaxReceiveCount(v int32) *SetQueueAttributesRequestDlqPolicy {
	s.MaxReceiveCount = &v
	return s
}

type SetQueueAttributesShrinkRequest struct {
	// The period after which all messages sent to the queue are consumed. Valid values: 0 to 604800. Unit: seconds. Default value: 0
	//
	// example:
	//
	// 0
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The dead-letter queue policy.
	DlqPolicyShrink *string `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty"`
	// Specifies whether to enable the log management feature. Valid values:
	//
	// 	- true: enabled.
	//
	// 	- false: disabled. Default value: false.
	//
	// example:
	//
	// True
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// The maximum length of the message that is sent to the queue. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 1024
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is received. Valid values: 60 to 604800. Unit: seconds. Default value: 345600.
	//
	// example:
	//
	// 120
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Valid values: 0 to 30. Unit: seconds. Default value: 0
	//
	// example:
	//
	// 0
	PollingWaitSeconds *int64 `json:"PollingWaitSeconds,omitempty" xml:"PollingWaitSeconds,omitempty"`
	// The name of the queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// testqueue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 60
	VisibilityTimeout *int64 `json:"VisibilityTimeout,omitempty" xml:"VisibilityTimeout,omitempty"`
}

func (s SetQueueAttributesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetQueueAttributesShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesShrinkRequest) SetDelaySeconds(v int64) *SetQueueAttributesShrinkRequest {
	s.DelaySeconds = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetDlqPolicyShrink(v string) *SetQueueAttributesShrinkRequest {
	s.DlqPolicyShrink = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetEnableLogging(v bool) *SetQueueAttributesShrinkRequest {
	s.EnableLogging = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetMaximumMessageSize(v int64) *SetQueueAttributesShrinkRequest {
	s.MaximumMessageSize = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetMessageRetentionPeriod(v int64) *SetQueueAttributesShrinkRequest {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetPollingWaitSeconds(v int64) *SetQueueAttributesShrinkRequest {
	s.PollingWaitSeconds = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetQueueName(v string) *SetQueueAttributesShrinkRequest {
	s.QueueName = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetVisibilityTimeout(v int64) *SetQueueAttributesShrinkRequest {
	s.VisibilityTimeout = &v
	return s
}

type SetQueueAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *SetQueueAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetQueueAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetQueueAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesResponseBody) SetCode(v int64) *SetQueueAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *SetQueueAttributesResponseBody) SetData(v *SetQueueAttributesResponseBodyData) *SetQueueAttributesResponseBody {
	s.Data = v
	return s
}

func (s *SetQueueAttributesResponseBody) SetMessage(v string) *SetQueueAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *SetQueueAttributesResponseBody) SetRequestId(v string) *SetQueueAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetQueueAttributesResponseBody) SetStatus(v string) *SetQueueAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *SetQueueAttributesResponseBody) SetSuccess(v bool) *SetQueueAttributesResponseBody {
	s.Success = &v
	return s
}

type SetQueueAttributesResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetQueueAttributesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SetQueueAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesResponseBodyData) SetCode(v int64) *SetQueueAttributesResponseBodyData {
	s.Code = &v
	return s
}

func (s *SetQueueAttributesResponseBodyData) SetMessage(v string) *SetQueueAttributesResponseBodyData {
	s.Message = &v
	return s
}

func (s *SetQueueAttributesResponseBodyData) SetSuccess(v bool) *SetQueueAttributesResponseBodyData {
	s.Success = &v
	return s
}

type SetQueueAttributesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetQueueAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetQueueAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s SetQueueAttributesResponse) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesResponse) SetHeaders(v map[string]*string) *SetQueueAttributesResponse {
	s.Headers = v
	return s
}

func (s *SetQueueAttributesResponse) SetStatusCode(v int32) *SetQueueAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetQueueAttributesResponse) SetBody(v *SetQueueAttributesResponseBody) *SetQueueAttributesResponse {
	s.Body = v
	return s
}

type SetSubscriptionAttributesRequest struct {
	// The dead-letter queue policy.
	DlqPolicy *SetSubscriptionAttributesRequestDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	// The retry policy that is applied if an error occurs when Message Service (MNS) pushes messages to the endpoint. Valid values:
	//
	// 	- BACKOFF_RETRY
	//
	// 	- EXPONENTIAL_DECAY_RETRY
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The name of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s SetSubscriptionAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSubscriptionAttributesRequest) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesRequest) SetDlqPolicy(v *SetSubscriptionAttributesRequestDlqPolicy) *SetSubscriptionAttributesRequest {
	s.DlqPolicy = v
	return s
}

func (s *SetSubscriptionAttributesRequest) SetNotifyStrategy(v string) *SetSubscriptionAttributesRequest {
	s.NotifyStrategy = &v
	return s
}

func (s *SetSubscriptionAttributesRequest) SetSubscriptionName(v string) *SetSubscriptionAttributesRequest {
	s.SubscriptionName = &v
	return s
}

func (s *SetSubscriptionAttributesRequest) SetTopicName(v string) *SetSubscriptionAttributesRequest {
	s.TopicName = &v
	return s
}

type SetSubscriptionAttributesRequestDlqPolicy struct {
	// The queue to which dead-letter messages are delivered.
	//
	// example:
	//
	// deadLetterTargetQueue
	DeadLetterTargetQueue *string `json:"DeadLetterTargetQueue,omitempty" xml:"DeadLetterTargetQueue,omitempty"`
	// Specifies whether to enable the dead-letter message delivery.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s SetSubscriptionAttributesRequestDlqPolicy) String() string {
	return tea.Prettify(s)
}

func (s SetSubscriptionAttributesRequestDlqPolicy) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesRequestDlqPolicy) SetDeadLetterTargetQueue(v string) *SetSubscriptionAttributesRequestDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *SetSubscriptionAttributesRequestDlqPolicy) SetEnabled(v bool) *SetSubscriptionAttributesRequestDlqPolicy {
	s.Enabled = &v
	return s
}

type SetSubscriptionAttributesShrinkRequest struct {
	// The dead-letter queue policy.
	DlqPolicyShrink *string `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty"`
	// The retry policy that is applied if an error occurs when Message Service (MNS) pushes messages to the endpoint. Valid values:
	//
	// 	- BACKOFF_RETRY
	//
	// 	- EXPONENTIAL_DECAY_RETRY
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The name of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s SetSubscriptionAttributesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSubscriptionAttributesShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesShrinkRequest) SetDlqPolicyShrink(v string) *SetSubscriptionAttributesShrinkRequest {
	s.DlqPolicyShrink = &v
	return s
}

func (s *SetSubscriptionAttributesShrinkRequest) SetNotifyStrategy(v string) *SetSubscriptionAttributesShrinkRequest {
	s.NotifyStrategy = &v
	return s
}

func (s *SetSubscriptionAttributesShrinkRequest) SetSubscriptionName(v string) *SetSubscriptionAttributesShrinkRequest {
	s.SubscriptionName = &v
	return s
}

func (s *SetSubscriptionAttributesShrinkRequest) SetTopicName(v string) *SetSubscriptionAttributesShrinkRequest {
	s.TopicName = &v
	return s
}

type SetSubscriptionAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *SetSubscriptionAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetSubscriptionAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSubscriptionAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesResponseBody) SetCode(v int64) *SetSubscriptionAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetData(v *SetSubscriptionAttributesResponseBodyData) *SetSubscriptionAttributesResponseBody {
	s.Data = v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetMessage(v string) *SetSubscriptionAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetRequestId(v string) *SetSubscriptionAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetStatus(v string) *SetSubscriptionAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetSuccess(v bool) *SetSubscriptionAttributesResponseBody {
	s.Success = &v
	return s
}

type SetSubscriptionAttributesResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetSubscriptionAttributesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SetSubscriptionAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesResponseBodyData) SetCode(v int64) *SetSubscriptionAttributesResponseBodyData {
	s.Code = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBodyData) SetMessage(v string) *SetSubscriptionAttributesResponseBodyData {
	s.Message = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBodyData) SetSuccess(v bool) *SetSubscriptionAttributesResponseBodyData {
	s.Success = &v
	return s
}

type SetSubscriptionAttributesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSubscriptionAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSubscriptionAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSubscriptionAttributesResponse) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesResponse) SetHeaders(v map[string]*string) *SetSubscriptionAttributesResponse {
	s.Headers = v
	return s
}

func (s *SetSubscriptionAttributesResponse) SetStatusCode(v int32) *SetSubscriptionAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSubscriptionAttributesResponse) SetBody(v *SetSubscriptionAttributesResponseBody) *SetSubscriptionAttributesResponse {
	s.Body = v
	return s
}

type SetTopicAttributesRequest struct {
	// Specifies whether to enable the log management feature. Valid values:
	//
	// 	- true: enabled.
	//
	// 	- false: disabled. Default value: false.
	//
	// example:
	//
	// True
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// The maximum length of the message that is sent to the topic. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 65536
	MaxMessageSize *int64 `json:"MaxMessageSize,omitempty" xml:"MaxMessageSize,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s SetTopicAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s SetTopicAttributesRequest) GoString() string {
	return s.String()
}

func (s *SetTopicAttributesRequest) SetEnableLogging(v bool) *SetTopicAttributesRequest {
	s.EnableLogging = &v
	return s
}

func (s *SetTopicAttributesRequest) SetMaxMessageSize(v int64) *SetTopicAttributesRequest {
	s.MaxMessageSize = &v
	return s
}

func (s *SetTopicAttributesRequest) SetTopicName(v string) *SetTopicAttributesRequest {
	s.TopicName = &v
	return s
}

type SetTopicAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *SetTopicAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetTopicAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetTopicAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *SetTopicAttributesResponseBody) SetCode(v int64) *SetTopicAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *SetTopicAttributesResponseBody) SetData(v *SetTopicAttributesResponseBodyData) *SetTopicAttributesResponseBody {
	s.Data = v
	return s
}

func (s *SetTopicAttributesResponseBody) SetMessage(v string) *SetTopicAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *SetTopicAttributesResponseBody) SetRequestId(v string) *SetTopicAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetTopicAttributesResponseBody) SetStatus(v string) *SetTopicAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *SetTopicAttributesResponseBody) SetSuccess(v bool) *SetTopicAttributesResponseBody {
	s.Success = &v
	return s
}

type SetTopicAttributesResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetTopicAttributesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SetTopicAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetTopicAttributesResponseBodyData) SetCode(v int64) *SetTopicAttributesResponseBodyData {
	s.Code = &v
	return s
}

func (s *SetTopicAttributesResponseBodyData) SetMessage(v string) *SetTopicAttributesResponseBodyData {
	s.Message = &v
	return s
}

func (s *SetTopicAttributesResponseBodyData) SetSuccess(v bool) *SetTopicAttributesResponseBodyData {
	s.Success = &v
	return s
}

type SetTopicAttributesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetTopicAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetTopicAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s SetTopicAttributesResponse) GoString() string {
	return s.String()
}

func (s *SetTopicAttributesResponse) SetHeaders(v map[string]*string) *SetTopicAttributesResponse {
	s.Headers = v
	return s
}

func (s *SetTopicAttributesResponse) SetStatusCode(v int32) *SetTopicAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetTopicAttributesResponse) SetBody(v *SetTopicAttributesResponseBody) *SetTopicAttributesResponse {
	s.Body = v
	return s
}

type SubscribeRequest struct {
	// The dead-letter queue policy.
	DlqPolicy *SubscribeRequestDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	// The receiver endpoint. The format of the endpoint varies based on the terminal type.
	//
	// 	- If you set PushType to http, set Endpoint to an `HTTP URL that starts with http:// or https://`.
	//
	// 	- If you set PushType to queue, set Endpoint to a `queue name`.
	//
	// 	- If you set PushType to mpush, set Endpoint to an `AppKey`.
	//
	// 	- If you set PushType to alisms, set Endpoint to a `mobile number`.
	//
	// 	- If you set PushType to email, set Endpoint to an `email address`.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The tag that is used to filter messages. Only messages that have the same tag can be pushed. Set the value to a string of no more than 16 characters.
	//
	// By default, no tag is specified to filter messages.
	//
	// example:
	//
	// important
	MessageTag *string `json:"MessageTag,omitempty" xml:"MessageTag,omitempty"`
	// The content format of the messages that are pushed to the endpoint. Valid values:
	//
	// 	- XML
	//
	// 	- JSON
	//
	// 	- SIMPLIFIED
	//
	// example:
	//
	// XML
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" xml:"NotifyContentFormat,omitempty"`
	// The retry policy that is applied if an error occurs when Message Service (MNS) pushes messages to the endpoint. Valid values:
	//
	// 	- BACKOFF_RETRY
	//
	// 	- EXPONENTIAL_DECAY_RETRY
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The terminal type. Valid values:
	//
	// 	- http: HTTP services
	//
	// 	- queue: queues
	//
	// 	- mpush: mobile devices
	//
	// 	- alisms: Alibaba Cloud Short Message Service (SMS)
	//
	// 	- email: emails
	//
	// This parameter is required.
	//
	// example:
	//
	// queue
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// The name of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s SubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscribeRequest) GoString() string {
	return s.String()
}

func (s *SubscribeRequest) SetDlqPolicy(v *SubscribeRequestDlqPolicy) *SubscribeRequest {
	s.DlqPolicy = v
	return s
}

func (s *SubscribeRequest) SetEndpoint(v string) *SubscribeRequest {
	s.Endpoint = &v
	return s
}

func (s *SubscribeRequest) SetMessageTag(v string) *SubscribeRequest {
	s.MessageTag = &v
	return s
}

func (s *SubscribeRequest) SetNotifyContentFormat(v string) *SubscribeRequest {
	s.NotifyContentFormat = &v
	return s
}

func (s *SubscribeRequest) SetNotifyStrategy(v string) *SubscribeRequest {
	s.NotifyStrategy = &v
	return s
}

func (s *SubscribeRequest) SetPushType(v string) *SubscribeRequest {
	s.PushType = &v
	return s
}

func (s *SubscribeRequest) SetSubscriptionName(v string) *SubscribeRequest {
	s.SubscriptionName = &v
	return s
}

func (s *SubscribeRequest) SetTopicName(v string) *SubscribeRequest {
	s.TopicName = &v
	return s
}

type SubscribeRequestDlqPolicy struct {
	// The queue to which dead-letter messages are delivered.
	//
	// example:
	//
	// deadLetterTargetQueue
	DeadLetterTargetQueue *string `json:"DeadLetterTargetQueue,omitempty" xml:"DeadLetterTargetQueue,omitempty"`
	// Specifies whether to enable the dead-letter message delivery.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s SubscribeRequestDlqPolicy) String() string {
	return tea.Prettify(s)
}

func (s SubscribeRequestDlqPolicy) GoString() string {
	return s.String()
}

func (s *SubscribeRequestDlqPolicy) SetDeadLetterTargetQueue(v string) *SubscribeRequestDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *SubscribeRequestDlqPolicy) SetEnabled(v bool) *SubscribeRequestDlqPolicy {
	s.Enabled = &v
	return s
}

type SubscribeShrinkRequest struct {
	// The dead-letter queue policy.
	DlqPolicyShrink *string `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty"`
	// The receiver endpoint. The format of the endpoint varies based on the terminal type.
	//
	// 	- If you set PushType to http, set Endpoint to an `HTTP URL that starts with http:// or https://`.
	//
	// 	- If you set PushType to queue, set Endpoint to a `queue name`.
	//
	// 	- If you set PushType to mpush, set Endpoint to an `AppKey`.
	//
	// 	- If you set PushType to alisms, set Endpoint to a `mobile number`.
	//
	// 	- If you set PushType to email, set Endpoint to an `email address`.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The tag that is used to filter messages. Only messages that have the same tag can be pushed. Set the value to a string of no more than 16 characters.
	//
	// By default, no tag is specified to filter messages.
	//
	// example:
	//
	// important
	MessageTag *string `json:"MessageTag,omitempty" xml:"MessageTag,omitempty"`
	// The content format of the messages that are pushed to the endpoint. Valid values:
	//
	// 	- XML
	//
	// 	- JSON
	//
	// 	- SIMPLIFIED
	//
	// example:
	//
	// XML
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" xml:"NotifyContentFormat,omitempty"`
	// The retry policy that is applied if an error occurs when Message Service (MNS) pushes messages to the endpoint. Valid values:
	//
	// 	- BACKOFF_RETRY
	//
	// 	- EXPONENTIAL_DECAY_RETRY
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The terminal type. Valid values:
	//
	// 	- http: HTTP services
	//
	// 	- queue: queues
	//
	// 	- mpush: mobile devices
	//
	// 	- alisms: Alibaba Cloud Short Message Service (SMS)
	//
	// 	- email: emails
	//
	// This parameter is required.
	//
	// example:
	//
	// queue
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// The name of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s SubscribeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscribeShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubscribeShrinkRequest) SetDlqPolicyShrink(v string) *SubscribeShrinkRequest {
	s.DlqPolicyShrink = &v
	return s
}

func (s *SubscribeShrinkRequest) SetEndpoint(v string) *SubscribeShrinkRequest {
	s.Endpoint = &v
	return s
}

func (s *SubscribeShrinkRequest) SetMessageTag(v string) *SubscribeShrinkRequest {
	s.MessageTag = &v
	return s
}

func (s *SubscribeShrinkRequest) SetNotifyContentFormat(v string) *SubscribeShrinkRequest {
	s.NotifyContentFormat = &v
	return s
}

func (s *SubscribeShrinkRequest) SetNotifyStrategy(v string) *SubscribeShrinkRequest {
	s.NotifyStrategy = &v
	return s
}

func (s *SubscribeShrinkRequest) SetPushType(v string) *SubscribeShrinkRequest {
	s.PushType = &v
	return s
}

func (s *SubscribeShrinkRequest) SetSubscriptionName(v string) *SubscribeShrinkRequest {
	s.SubscriptionName = &v
	return s
}

func (s *SubscribeShrinkRequest) SetTopicName(v string) *SubscribeShrinkRequest {
	s.TopicName = &v
	return s
}

type SubscribeResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// {\\"Code\\": 200, \\"Success\\": True}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubscribeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeResponseBody) SetCode(v int64) *SubscribeResponseBody {
	s.Code = &v
	return s
}

func (s *SubscribeResponseBody) SetData(v string) *SubscribeResponseBody {
	s.Data = &v
	return s
}

func (s *SubscribeResponseBody) SetMessage(v string) *SubscribeResponseBody {
	s.Message = &v
	return s
}

func (s *SubscribeResponseBody) SetRequestId(v string) *SubscribeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubscribeResponseBody) SetStatus(v string) *SubscribeResponseBody {
	s.Status = &v
	return s
}

func (s *SubscribeResponseBody) SetSuccess(v bool) *SubscribeResponseBody {
	s.Success = &v
	return s
}

type SubscribeResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubscribeResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscribeResponse) GoString() string {
	return s.String()
}

func (s *SubscribeResponse) SetHeaders(v map[string]*string) *SubscribeResponse {
	s.Headers = v
	return s
}

func (s *SubscribeResponse) SetStatusCode(v int32) *SubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscribeResponse) SetBody(v *SubscribeResponseBody) *SubscribeResponse {
	s.Body = v
	return s
}

type UnsubscribeRequest struct {
	// The name of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s UnsubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeRequest) GoString() string {
	return s.String()
}

func (s *UnsubscribeRequest) SetSubscriptionName(v string) *UnsubscribeRequest {
	s.SubscriptionName = &v
	return s
}

func (s *UnsubscribeRequest) SetTopicName(v string) *UnsubscribeRequest {
	s.TopicName = &v
	return s
}

type UnsubscribeResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *UnsubscribeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnsubscribeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *UnsubscribeResponseBody) SetCode(v int64) *UnsubscribeResponseBody {
	s.Code = &v
	return s
}

func (s *UnsubscribeResponseBody) SetData(v *UnsubscribeResponseBodyData) *UnsubscribeResponseBody {
	s.Data = v
	return s
}

func (s *UnsubscribeResponseBody) SetMessage(v string) *UnsubscribeResponseBody {
	s.Message = &v
	return s
}

func (s *UnsubscribeResponseBody) SetRequestId(v string) *UnsubscribeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnsubscribeResponseBody) SetStatus(v string) *UnsubscribeResponseBody {
	s.Status = &v
	return s
}

func (s *UnsubscribeResponseBody) SetSuccess(v bool) *UnsubscribeResponseBody {
	s.Success = &v
	return s
}

type UnsubscribeResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnsubscribeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnsubscribeResponseBodyData) SetCode(v int64) *UnsubscribeResponseBodyData {
	s.Code = &v
	return s
}

func (s *UnsubscribeResponseBodyData) SetMessage(v string) *UnsubscribeResponseBodyData {
	s.Message = &v
	return s
}

func (s *UnsubscribeResponseBodyData) SetSuccess(v bool) *UnsubscribeResponseBodyData {
	s.Success = &v
	return s
}

type UnsubscribeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnsubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnsubscribeResponse) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeResponse) GoString() string {
	return s.String()
}

func (s *UnsubscribeResponse) SetHeaders(v map[string]*string) *UnsubscribeResponse {
	s.Headers = v
	return s
}

func (s *UnsubscribeResponse) SetStatusCode(v int32) *UnsubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *UnsubscribeResponse) SetBody(v *UnsubscribeResponseBody) *UnsubscribeResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("mns-open"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to add one or more rules of access control lists (ACLs) for the endpoint of a type.
//
// @param tmpReq - AuthorizeEndpointAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeEndpointAclResponse
func (client *Client) AuthorizeEndpointAclWithOptions(tmpReq *AuthorizeEndpointAclRequest, runtime *util.RuntimeOptions) (_result *AuthorizeEndpointAclResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AuthorizeEndpointAclShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CidrList)) {
		request.CidrListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CidrList, tea.String("CidrList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclStrategy)) {
		query["AclStrategy"] = request.AclStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.CidrListShrink)) {
		query["CidrList"] = request.CidrListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthorizeEndpointAcl"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthorizeEndpointAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to add one or more rules of access control lists (ACLs) for the endpoint of a type.
//
// @param request - AuthorizeEndpointAclRequest
//
// @return AuthorizeEndpointAclResponse
func (client *Client) AuthorizeEndpointAcl(request *AuthorizeEndpointAclRequest) (_result *AuthorizeEndpointAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeEndpointAclResponse{}
	_body, _err := client.AuthorizeEndpointAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建事件规则
//
// @param tmpReq - CreateEventRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEventRuleResponse
func (client *Client) CreateEventRuleWithOptions(tmpReq *CreateEventRuleRequest, runtime *util.RuntimeOptions) (_result *CreateEventRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateEventRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Endpoints)) {
		request.EndpointsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Endpoints, tea.String("Endpoints"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.EventTypes)) {
		request.EventTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventTypes, tea.String("EventTypes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.MatchRules)) {
		request.MatchRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MatchRules, tea.String("MatchRules"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointsShrink)) {
		query["Endpoints"] = request.EndpointsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EventTypesShrink)) {
		query["EventTypes"] = request.EventTypesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MatchRulesShrink)) {
		query["MatchRules"] = request.MatchRulesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		query["ProductName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEventRule"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEventRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建事件规则
//
// @param request - CreateEventRuleRequest
//
// @return CreateEventRuleResponse
func (client *Client) CreateEventRule(request *CreateEventRuleRequest) (_result *CreateEventRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEventRuleResponse{}
	_body, _err := client.CreateEventRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a queue.
//
// @param tmpReq - CreateQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQueueResponse
func (client *Client) CreateQueueWithOptions(tmpReq *CreateQueueRequest, runtime *util.RuntimeOptions) (_result *CreateQueueResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateQueueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DlqPolicy)) {
		request.DlqPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DlqPolicy, tea.String("DlqPolicy"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DelaySeconds)) {
		query["DelaySeconds"] = request.DelaySeconds
	}

	if !tea.BoolValue(util.IsUnset(request.DlqPolicyShrink)) {
		query["DlqPolicy"] = request.DlqPolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EnableLogging)) {
		query["EnableLogging"] = request.EnableLogging
	}

	if !tea.BoolValue(util.IsUnset(request.MaximumMessageSize)) {
		query["MaximumMessageSize"] = request.MaximumMessageSize
	}

	if !tea.BoolValue(util.IsUnset(request.MessageRetentionPeriod)) {
		query["MessageRetentionPeriod"] = request.MessageRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PollingWaitSeconds)) {
		query["PollingWaitSeconds"] = request.PollingWaitSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VisibilityTimeout)) {
		query["VisibilityTimeout"] = request.VisibilityTimeout
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQueue"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a queue.
//
// @param request - CreateQueueRequest
//
// @return CreateQueueResponse
func (client *Client) CreateQueue(request *CreateQueueRequest) (_result *CreateQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateQueueResponse{}
	_body, _err := client.CreateQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a topic.
//
// @param request - CreateTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTopicResponse
func (client *Client) CreateTopicWithOptions(request *CreateTopicRequest, runtime *util.RuntimeOptions) (_result *CreateTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableLogging)) {
		body["EnableLogging"] = request.EnableLogging
	}

	if !tea.BoolValue(util.IsUnset(request.MaxMessageSize)) {
		body["MaxMessageSize"] = request.MaxMessageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		body["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTopic"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a topic.
//
// @param request - CreateTopicRequest
//
// @return CreateTopicResponse
func (client *Client) CreateTopic(request *CreateTopicRequest) (_result *CreateTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTopicResponse{}
	_body, _err := client.CreateTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除事件规则
//
// @param request - DeleteEventRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventRuleResponse
func (client *Client) DeleteEventRuleWithOptions(request *DeleteEventRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteEventRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		query["ProductName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEventRule"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEventRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除事件规则
//
// @param request - DeleteEventRuleRequest
//
// @return DeleteEventRuleResponse
func (client *Client) DeleteEventRule(request *DeleteEventRuleRequest) (_result *DeleteEventRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEventRuleResponse{}
	_body, _err := client.DeleteEventRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a queue.
//
// @param request - DeleteQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQueueResponse
func (client *Client) DeleteQueueWithOptions(request *DeleteQueueRequest, runtime *util.RuntimeOptions) (_result *DeleteQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteQueue"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a queue.
//
// @param request - DeleteQueueRequest
//
// @return DeleteQueueResponse
func (client *Client) DeleteQueue(request *DeleteQueueRequest) (_result *DeleteQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteQueueResponse{}
	_body, _err := client.DeleteQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a topic.
//
// @param request - DeleteTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTopicResponse
func (client *Client) DeleteTopicWithOptions(request *DeleteTopicRequest, runtime *util.RuntimeOptions) (_result *DeleteTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTopic"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a topic.
//
// @param request - DeleteTopicRequest
//
// @return DeleteTopicResponse
func (client *Client) DeleteTopic(request *DeleteTopicRequest) (_result *DeleteTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTopicResponse{}
	_body, _err := client.DeleteTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to disenable the endpoint of a type. After the endpoint is disabled, all requests from the endpoint are blocked and an error is returned.
//
// @param request - DisableEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableEndpointResponse
func (client *Client) DisableEndpointWithOptions(request *DisableEndpointRequest, runtime *util.RuntimeOptions) (_result *DisableEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableEndpoint"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to disenable the endpoint of a type. After the endpoint is disabled, all requests from the endpoint are blocked and an error is returned.
//
// @param request - DisableEndpointRequest
//
// @return DisableEndpointResponse
func (client *Client) DisableEndpoint(request *DisableEndpointRequest) (_result *DisableEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableEndpointResponse{}
	_body, _err := client.DisableEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to enable the endpoint of a type. If the endpoint is enabled, requests from the endpoint that are included in the access control lists (ACLs) are not blocked.
//
// @param request - EnableEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableEndpointResponse
func (client *Client) EnableEndpointWithOptions(request *EnableEndpointRequest, runtime *util.RuntimeOptions) (_result *EnableEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableEndpoint"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to enable the endpoint of a type. If the endpoint is enabled, requests from the endpoint that are included in the access control lists (ACLs) are not blocked.
//
// @param request - EnableEndpointRequest
//
// @return EnableEndpointResponse
func (client *Client) EnableEndpoint(request *EnableEndpointRequest) (_result *EnableEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableEndpointResponse{}
	_body, _err := client.EnableEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetEndpointAttribute
//
// @param request - GetEndpointAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEndpointAttributeResponse
func (client *Client) GetEndpointAttributeWithOptions(request *GetEndpointAttributeRequest, runtime *util.RuntimeOptions) (_result *GetEndpointAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEndpointAttribute"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEndpointAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetEndpointAttribute
//
// @param request - GetEndpointAttributeRequest
//
// @return GetEndpointAttributeResponse
func (client *Client) GetEndpointAttribute(request *GetEndpointAttributeRequest) (_result *GetEndpointAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEndpointAttributeResponse{}
	_body, _err := client.GetEndpointAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attributes of an existing queue.
//
// @param request - GetQueueAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueAttributesResponse
func (client *Client) GetQueueAttributesWithOptions(request *GetQueueAttributesRequest, runtime *util.RuntimeOptions) (_result *GetQueueAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQueueAttributes"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQueueAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of an existing queue.
//
// @param request - GetQueueAttributesRequest
//
// @return GetQueueAttributesResponse
func (client *Client) GetQueueAttributes(request *GetQueueAttributesRequest) (_result *GetQueueAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQueueAttributesResponse{}
	_body, _err := client.GetQueueAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attributes of a subscription.
//
// @param request - GetSubscriptionAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubscriptionAttributesResponse
func (client *Client) GetSubscriptionAttributesWithOptions(request *GetSubscriptionAttributesRequest, runtime *util.RuntimeOptions) (_result *GetSubscriptionAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubscriptionName)) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSubscriptionAttributes"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSubscriptionAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of a subscription.
//
// @param request - GetSubscriptionAttributesRequest
//
// @return GetSubscriptionAttributesResponse
func (client *Client) GetSubscriptionAttributes(request *GetSubscriptionAttributesRequest) (_result *GetSubscriptionAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSubscriptionAttributesResponse{}
	_body, _err := client.GetSubscriptionAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attributes of a topic.
//
// @param request - GetTopicAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicAttributesResponse
func (client *Client) GetTopicAttributesWithOptions(request *GetTopicAttributesRequest, runtime *util.RuntimeOptions) (_result *GetTopicAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTopicAttributes"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTopicAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of a topic.
//
// @param request - GetTopicAttributesRequest
//
// @return GetTopicAttributesResponse
func (client *Client) GetTopicAttributes(request *GetTopicAttributesRequest) (_result *GetTopicAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTopicAttributesResponse{}
	_body, _err := client.GetTopicAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all queues that belong to an Alibaba Cloud account. The queues are displayed by page.
//
// @param request - ListQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueueResponse
func (client *Client) ListQueueWithOptions(request *ListQueueRequest, runtime *util.RuntimeOptions) (_result *ListQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueue"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all queues that belong to an Alibaba Cloud account. The queues are displayed by page.
//
// @param request - ListQueueRequest
//
// @return ListQueueResponse
func (client *Client) ListQueue(request *ListQueueRequest) (_result *ListQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQueueResponse{}
	_body, _err := client.ListQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all subscriptions to a topic. The subscriptions are displayed by page.
//
// @param request - ListSubscriptionByTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubscriptionByTopicResponse
func (client *Client) ListSubscriptionByTopicWithOptions(request *ListSubscriptionByTopicRequest, runtime *util.RuntimeOptions) (_result *ListSubscriptionByTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointValue)) {
		query["EndpointValue"] = request.EndpointValue
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionName)) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSubscriptionByTopic"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubscriptionByTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all subscriptions to a topic. The subscriptions are displayed by page.
//
// @param request - ListSubscriptionByTopicRequest
//
// @return ListSubscriptionByTopicResponse
func (client *Client) ListSubscriptionByTopic(request *ListSubscriptionByTopicRequest) (_result *ListSubscriptionByTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSubscriptionByTopicResponse{}
	_body, _err := client.ListSubscriptionByTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the topics that belong to an Alibaba Cloud account. The topics are displayed by page.
//
// @param request - ListTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTopicResponse
func (client *Client) ListTopicWithOptions(request *ListTopicRequest, runtime *util.RuntimeOptions) (_result *ListTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTopic"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the topics that belong to an Alibaba Cloud account. The topics are displayed by page.
//
// @param request - ListTopicRequest
//
// @return ListTopicResponse
func (client *Client) ListTopic(request *ListTopicRequest) (_result *ListTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTopicResponse{}
	_body, _err := client.ListTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to delete one or more rules of access control lists (ACLs) for the endpoint of a type.
//
// @param tmpReq - RevokeEndpointAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeEndpointAclResponse
func (client *Client) RevokeEndpointAclWithOptions(tmpReq *RevokeEndpointAclRequest, runtime *util.RuntimeOptions) (_result *RevokeEndpointAclResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RevokeEndpointAclShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CidrList)) {
		request.CidrListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CidrList, tea.String("CidrList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclStrategy)) {
		query["AclStrategy"] = request.AclStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.CidrListShrink)) {
		query["CidrList"] = request.CidrListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeEndpointAcl"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeEndpointAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to delete one or more rules of access control lists (ACLs) for the endpoint of a type.
//
// @param request - RevokeEndpointAclRequest
//
// @return RevokeEndpointAclResponse
func (client *Client) RevokeEndpointAcl(request *RevokeEndpointAclRequest) (_result *RevokeEndpointAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeEndpointAclResponse{}
	_body, _err := client.RevokeEndpointAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a queue.
//
// @param tmpReq - SetQueueAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetQueueAttributesResponse
func (client *Client) SetQueueAttributesWithOptions(tmpReq *SetQueueAttributesRequest, runtime *util.RuntimeOptions) (_result *SetQueueAttributesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetQueueAttributesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DlqPolicy)) {
		request.DlqPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DlqPolicy, tea.String("DlqPolicy"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DelaySeconds)) {
		query["DelaySeconds"] = request.DelaySeconds
	}

	if !tea.BoolValue(util.IsUnset(request.DlqPolicyShrink)) {
		query["DlqPolicy"] = request.DlqPolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EnableLogging)) {
		query["EnableLogging"] = request.EnableLogging
	}

	if !tea.BoolValue(util.IsUnset(request.MaximumMessageSize)) {
		query["MaximumMessageSize"] = request.MaximumMessageSize
	}

	if !tea.BoolValue(util.IsUnset(request.MessageRetentionPeriod)) {
		query["MessageRetentionPeriod"] = request.MessageRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PollingWaitSeconds)) {
		query["PollingWaitSeconds"] = request.PollingWaitSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.VisibilityTimeout)) {
		query["VisibilityTimeout"] = request.VisibilityTimeout
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetQueueAttributes"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetQueueAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a queue.
//
// @param request - SetQueueAttributesRequest
//
// @return SetQueueAttributesResponse
func (client *Client) SetQueueAttributes(request *SetQueueAttributesRequest) (_result *SetQueueAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetQueueAttributesResponse{}
	_body, _err := client.SetQueueAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a subscription.
//
// @param tmpReq - SetSubscriptionAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSubscriptionAttributesResponse
func (client *Client) SetSubscriptionAttributesWithOptions(tmpReq *SetSubscriptionAttributesRequest, runtime *util.RuntimeOptions) (_result *SetSubscriptionAttributesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetSubscriptionAttributesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DlqPolicy)) {
		request.DlqPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DlqPolicy, tea.String("DlqPolicy"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DlqPolicyShrink)) {
		query["DlqPolicy"] = request.DlqPolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyStrategy)) {
		query["NotifyStrategy"] = request.NotifyStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionName)) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetSubscriptionAttributes"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetSubscriptionAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a subscription.
//
// @param request - SetSubscriptionAttributesRequest
//
// @return SetSubscriptionAttributesResponse
func (client *Client) SetSubscriptionAttributes(request *SetSubscriptionAttributesRequest) (_result *SetSubscriptionAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSubscriptionAttributesResponse{}
	_body, _err := client.SetSubscriptionAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a topic.
//
// @param request - SetTopicAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetTopicAttributesResponse
func (client *Client) SetTopicAttributesWithOptions(request *SetTopicAttributesRequest, runtime *util.RuntimeOptions) (_result *SetTopicAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableLogging)) {
		query["EnableLogging"] = request.EnableLogging
	}

	if !tea.BoolValue(util.IsUnset(request.MaxMessageSize)) {
		query["MaxMessageSize"] = request.MaxMessageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetTopicAttributes"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetTopicAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a topic.
//
// @param request - SetTopicAttributesRequest
//
// @return SetTopicAttributesResponse
func (client *Client) SetTopicAttributes(request *SetTopicAttributesRequest) (_result *SetTopicAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetTopicAttributesResponse{}
	_body, _err := client.SetTopicAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a subscription to a topic.
//
// @param tmpReq - SubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubscribeResponse
func (client *Client) SubscribeWithOptions(tmpReq *SubscribeRequest, runtime *util.RuntimeOptions) (_result *SubscribeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubscribeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DlqPolicy)) {
		request.DlqPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DlqPolicy, tea.String("DlqPolicy"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DlqPolicyShrink)) {
		query["DlqPolicy"] = request.DlqPolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
		query["Endpoint"] = request.Endpoint
	}

	if !tea.BoolValue(util.IsUnset(request.MessageTag)) {
		query["MessageTag"] = request.MessageTag
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyContentFormat)) {
		query["NotifyContentFormat"] = request.NotifyContentFormat
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyStrategy)) {
		query["NotifyStrategy"] = request.NotifyStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.PushType)) {
		query["PushType"] = request.PushType
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionName)) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Subscribe"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubscribeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a subscription to a topic.
//
// @param request - SubscribeRequest
//
// @return SubscribeResponse
func (client *Client) Subscribe(request *SubscribeRequest) (_result *SubscribeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubscribeResponse{}
	_body, _err := client.SubscribeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a subscription.
//
// @param request - UnsubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnsubscribeResponse
func (client *Client) UnsubscribeWithOptions(request *UnsubscribeRequest, runtime *util.RuntimeOptions) (_result *UnsubscribeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubscriptionName)) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Unsubscribe"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnsubscribeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a subscription.
//
// @param request - UnsubscribeRequest
//
// @return UnsubscribeResponse
func (client *Client) Unsubscribe(request *UnsubscribeRequest) (_result *UnsubscribeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnsubscribeResponse{}
	_body, _err := client.UnsubscribeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
