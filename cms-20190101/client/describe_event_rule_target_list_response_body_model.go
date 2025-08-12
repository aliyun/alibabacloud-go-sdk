// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventRuleTargetListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeEventRuleTargetListResponseBody
	GetCode() *string
	SetContactParameters(v *DescribeEventRuleTargetListResponseBodyContactParameters) *DescribeEventRuleTargetListResponseBody
	GetContactParameters() *DescribeEventRuleTargetListResponseBodyContactParameters
	SetFcParameters(v *DescribeEventRuleTargetListResponseBodyFcParameters) *DescribeEventRuleTargetListResponseBody
	GetFcParameters() *DescribeEventRuleTargetListResponseBodyFcParameters
	SetMessage(v string) *DescribeEventRuleTargetListResponseBody
	GetMessage() *string
	SetMnsParameters(v *DescribeEventRuleTargetListResponseBodyMnsParameters) *DescribeEventRuleTargetListResponseBody
	GetMnsParameters() *DescribeEventRuleTargetListResponseBodyMnsParameters
	SetOpenApiParameters(v *DescribeEventRuleTargetListResponseBodyOpenApiParameters) *DescribeEventRuleTargetListResponseBody
	GetOpenApiParameters() *DescribeEventRuleTargetListResponseBodyOpenApiParameters
	SetRequestId(v string) *DescribeEventRuleTargetListResponseBody
	GetRequestId() *string
	SetSlsParameters(v *DescribeEventRuleTargetListResponseBodySlsParameters) *DescribeEventRuleTargetListResponseBody
	GetSlsParameters() *DescribeEventRuleTargetListResponseBodySlsParameters
	SetWebhookParameters(v *DescribeEventRuleTargetListResponseBodyWebhookParameters) *DescribeEventRuleTargetListResponseBody
	GetWebhookParameters() *DescribeEventRuleTargetListResponseBodyWebhookParameters
}

type DescribeEventRuleTargetListResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the recipients if alert notifications are sent to the alert contacts of an alert contact group.
	ContactParameters *DescribeEventRuleTargetListResponseBodyContactParameters `json:"ContactParameters,omitempty" xml:"ContactParameters,omitempty" type:"Struct"`
	// The information about the recipients in Function Compute.
	FcParameters *DescribeEventRuleTargetListResponseBodyFcParameters `json:"FcParameters,omitempty" xml:"FcParameters,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The notifications of Simple Message Queue (formerly MNS) (SMQ).
	MnsParameters *DescribeEventRuleTargetListResponseBodyMnsParameters `json:"MnsParameters,omitempty" xml:"MnsParameters,omitempty" type:"Struct"`
	// The information about the recipients in OpenAPI Explorer.
	OpenApiParameters *DescribeEventRuleTargetListResponseBodyOpenApiParameters `json:"OpenApiParameters,omitempty" xml:"OpenApiParameters,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 87170bc7-e28a-4c93-b9bf-90a1dbe84736
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the recipients in Log Service.
	SlsParameters *DescribeEventRuleTargetListResponseBodySlsParameters `json:"SlsParameters,omitempty" xml:"SlsParameters,omitempty" type:"Struct"`
	// The information about the recipients if alert notifications are sent by sending a request to a callback URL.
	WebhookParameters *DescribeEventRuleTargetListResponseBodyWebhookParameters `json:"WebhookParameters,omitempty" xml:"WebhookParameters,omitempty" type:"Struct"`
}

func (s DescribeEventRuleTargetListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeEventRuleTargetListResponseBody) GetContactParameters() *DescribeEventRuleTargetListResponseBodyContactParameters {
	return s.ContactParameters
}

func (s *DescribeEventRuleTargetListResponseBody) GetFcParameters() *DescribeEventRuleTargetListResponseBodyFcParameters {
	return s.FcParameters
}

func (s *DescribeEventRuleTargetListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEventRuleTargetListResponseBody) GetMnsParameters() *DescribeEventRuleTargetListResponseBodyMnsParameters {
	return s.MnsParameters
}

func (s *DescribeEventRuleTargetListResponseBody) GetOpenApiParameters() *DescribeEventRuleTargetListResponseBodyOpenApiParameters {
	return s.OpenApiParameters
}

func (s *DescribeEventRuleTargetListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventRuleTargetListResponseBody) GetSlsParameters() *DescribeEventRuleTargetListResponseBodySlsParameters {
	return s.SlsParameters
}

func (s *DescribeEventRuleTargetListResponseBody) GetWebhookParameters() *DescribeEventRuleTargetListResponseBodyWebhookParameters {
	return s.WebhookParameters
}

func (s *DescribeEventRuleTargetListResponseBody) SetCode(v string) *DescribeEventRuleTargetListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetContactParameters(v *DescribeEventRuleTargetListResponseBodyContactParameters) *DescribeEventRuleTargetListResponseBody {
	s.ContactParameters = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetFcParameters(v *DescribeEventRuleTargetListResponseBodyFcParameters) *DescribeEventRuleTargetListResponseBody {
	s.FcParameters = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetMessage(v string) *DescribeEventRuleTargetListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetMnsParameters(v *DescribeEventRuleTargetListResponseBodyMnsParameters) *DescribeEventRuleTargetListResponseBody {
	s.MnsParameters = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetOpenApiParameters(v *DescribeEventRuleTargetListResponseBodyOpenApiParameters) *DescribeEventRuleTargetListResponseBody {
	s.OpenApiParameters = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetRequestId(v string) *DescribeEventRuleTargetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetSlsParameters(v *DescribeEventRuleTargetListResponseBodySlsParameters) *DescribeEventRuleTargetListResponseBody {
	s.SlsParameters = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetWebhookParameters(v *DescribeEventRuleTargetListResponseBodyWebhookParameters) *DescribeEventRuleTargetListResponseBody {
	s.WebhookParameters = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleTargetListResponseBodyContactParameters struct {
	ContactParameter []*DescribeEventRuleTargetListResponseBodyContactParametersContactParameter `json:"ContactParameter,omitempty" xml:"ContactParameter,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleTargetListResponseBodyContactParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyContactParameters) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyContactParameters) GetContactParameter() []*DescribeEventRuleTargetListResponseBodyContactParametersContactParameter {
	return s.ContactParameter
}

func (s *DescribeEventRuleTargetListResponseBodyContactParameters) SetContactParameter(v []*DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) *DescribeEventRuleTargetListResponseBodyContactParameters {
	s.ContactParameter = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyContactParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleTargetListResponseBodyContactParametersContactParameter struct {
	// The name of the alert group.
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	// The ID of the recipient.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The alert notification methods. Valid values:
	//
	// 4: Alert notifications are sent by using DingTalk chatbots and emails.
	//
	// example:
	//
	// 3
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) GetId() *string {
	return s.Id
}

func (s *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) GetLevel() *string {
	return s.Level
}

func (s *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) SetContactGroupName(v string) *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter {
	s.ContactGroupName = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) SetId(v string) *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter {
	s.Id = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) SetLevel(v string) *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter {
	s.Level = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleTargetListResponseBodyFcParameters struct {
	FCParameter []*DescribeEventRuleTargetListResponseBodyFcParametersFCParameter `json:"FCParameter,omitempty" xml:"FCParameter,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleTargetListResponseBodyFcParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyFcParameters) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyFcParameters) GetFCParameter() []*DescribeEventRuleTargetListResponseBodyFcParametersFCParameter {
	return s.FCParameter
}

func (s *DescribeEventRuleTargetListResponseBodyFcParameters) SetFCParameter(v []*DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) *DescribeEventRuleTargetListResponseBodyFcParameters {
	s.FCParameter = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyFcParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleTargetListResponseBodyFcParametersFCParameter struct {
	// The Alibaba Cloud Resource Name (ARN) of the function.
	//
	// Format: `arn:acs:${Service}:${Region}:${Account}:${ResourceType}/${ResourceId}`. Fields:
	//
	// - Service: the code of an Alibaba Cloud service
	//
	// - Region: the region ID
	//
	// - Account: the ID of an Alibaba Cloud account
	//
	// - ResourceType: the resource type
	//
	// - ResourceId: the resource ID
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The name of the function.
	//
	// example:
	//
	// fcTest1
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The ID of the recipient.
	//
	// example:
	//
	// 3
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region where Function Compute is deployed.
	//
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the Function Compute service.
	//
	// example:
	//
	// service1
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) GetArn() *string {
	return s.Arn
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) GetId() *string {
	return s.Id
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) GetRegion() *string {
	return s.Region
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) SetArn(v string) *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter {
	s.Arn = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) SetFunctionName(v string) *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter {
	s.FunctionName = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) SetId(v string) *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter {
	s.Id = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) SetRegion(v string) *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter {
	s.Region = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) SetServiceName(v string) *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter {
	s.ServiceName = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleTargetListResponseBodyMnsParameters struct {
	MnsParameter []*DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter `json:"MnsParameter,omitempty" xml:"MnsParameter,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleTargetListResponseBodyMnsParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyMnsParameters) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParameters) GetMnsParameter() []*DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter {
	return s.MnsParameter
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParameters) SetMnsParameter(v []*DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) *DescribeEventRuleTargetListResponseBodyMnsParameters {
	s.MnsParameter = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter struct {
	// The ARN of the MNS queue.
	//
	// Format: `arn:acs:${Service}:${Region}:${Account}:${ResourceType}/${ResourceId}`. Fields:
	//
	// - Service: the code of an Alibaba Cloud service
	//
	// - Region: the region ID
	//
	// - Account: the ID of an Alibaba Cloud account
	//
	// - ResourceType: the resource type
	//
	// - ResourceId: the resource ID
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the recipient.
	//
	// example:
	//
	// 2
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the SMQ queue.
	//
	// example:
	//
	// testQueue
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// The region for SMQ.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The SMQ topic.
	//
	// example:
	//
	// topic_sample
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) GetArn() *string {
	return s.Arn
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) GetId() *string {
	return s.Id
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) GetQueue() *string {
	return s.Queue
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) GetRegion() *string {
	return s.Region
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) GetTopic() *string {
	return s.Topic
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) SetArn(v string) *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter {
	s.Arn = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) SetId(v string) *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter {
	s.Id = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) SetQueue(v string) *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter {
	s.Queue = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) SetRegion(v string) *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter {
	s.Region = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) SetTopic(v string) *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter {
	s.Topic = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleTargetListResponseBodyOpenApiParameters struct {
	OpenApiParameters []*DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters `json:"OpenApiParameters,omitempty" xml:"OpenApiParameters,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleTargetListResponseBodyOpenApiParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyOpenApiParameters) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParameters) GetOpenApiParameters() []*DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters {
	return s.OpenApiParameters
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParameters) SetOpenApiParameters(v []*DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) *DescribeEventRuleTargetListResponseBodyOpenApiParameters {
	s.OpenApiParameters = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters struct {
	// The name of the API operation.
	//
	// example:
	//
	// PutLogs
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The ARN of the API operation.
	//
	// Format: `arn:acs:${Service}:${Region}:${Account}:${ResourceType}/${ResourceId}`. Fields:
	//
	// - Service: the code of an Alibaba Cloud service
	//
	// - Region: the region ID
	//
	// - Account: the ID of an Alibaba Cloud account
	//
	// - ResourceType: the resource type
	//
	// - ResourceId: the resource ID The ARN of the Log Service Logstore.
	//
	// Format: `arn:acs:${Service}:${Region}:${Account}:${ResourceType}/${ResourceId}`. Fields:
	//
	// - Service: the code of an Alibaba Cloud service
	//
	// - Region: the region ID
	//
	// - Account: the ID of an Alibaba Cloud account
	//
	// - ResourceType: the resource type
	//
	// - ResourceId: the resource ID
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the recipient.
	//
	// example:
	//
	// 3
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the cloud service to which the API operation belongs.
	//
	// example:
	//
	// log
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region where the resource resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the role.
	//
	// example:
	//
	// MyRole
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The version of the API.
	//
	// example:
	//
	// 2019-01-01
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) GetAction() *string {
	return s.Action
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) GetArn() *string {
	return s.Arn
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) GetId() *string {
	return s.Id
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) GetProduct() *string {
	return s.Product
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) GetRegion() *string {
	return s.Region
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) GetRole() *string {
	return s.Role
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) GetVersion() *string {
	return s.Version
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) SetAction(v string) *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters {
	s.Action = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) SetArn(v string) *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters {
	s.Arn = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) SetId(v string) *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters {
	s.Id = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) SetProduct(v string) *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters {
	s.Product = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) SetRegion(v string) *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters {
	s.Region = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) SetRole(v string) *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters {
	s.Role = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) SetVersion(v string) *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters {
	s.Version = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyOpenApiParametersOpenApiParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleTargetListResponseBodySlsParameters struct {
	SlsParameter []*DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter `json:"SlsParameter,omitempty" xml:"SlsParameter,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleTargetListResponseBodySlsParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodySlsParameters) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodySlsParameters) GetSlsParameter() []*DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter {
	return s.SlsParameter
}

func (s *DescribeEventRuleTargetListResponseBodySlsParameters) SetSlsParameter(v []*DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) *DescribeEventRuleTargetListResponseBodySlsParameters {
	s.SlsParameter = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodySlsParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter struct {
	// The ARN of the Log Service Logstore.
	//
	// Format: `arn:acs:${Service}:${Region}:${Account}:${ResourceType}/${ResourceId}`. Fields:
	//
	// - Service: the code of an Alibaba Cloud service
	//
	// - Region: the region ID
	//
	// - Account: the ID of an Alibaba Cloud account
	//
	// - ResourceType: the resource type
	//
	// - ResourceId: the resource ID
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the recipient.
	//
	// example:
	//
	// 4
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the Logstore.
	//
	// example:
	//
	// logstore_test
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// project_test
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region where the Log Service project resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) GetArn() *string {
	return s.Arn
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) GetId() *string {
	return s.Id
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) GetLogStore() *string {
	return s.LogStore
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) GetProject() *string {
	return s.Project
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) GetRegion() *string {
	return s.Region
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) SetArn(v string) *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter {
	s.Arn = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) SetId(v string) *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter {
	s.Id = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) SetLogStore(v string) *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter {
	s.LogStore = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) SetProject(v string) *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter {
	s.Project = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) SetRegion(v string) *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter {
	s.Region = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleTargetListResponseBodyWebhookParameters struct {
	WebhookParameter []*DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter `json:"WebhookParameter,omitempty" xml:"WebhookParameter,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleTargetListResponseBodyWebhookParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyWebhookParameters) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParameters) GetWebhookParameter() []*DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter {
	return s.WebhookParameter
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParameters) SetWebhookParameter(v []*DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) *DescribeEventRuleTargetListResponseBodyWebhookParameters {
	s.WebhookParameter = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter struct {
	// The ID of the recipient.
	//
	// example:
	//
	// 5
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The HTTP request method. Valid values: GET and POST.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The callback URL.
	//
	// example:
	//
	// http://www.aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) GetId() *string {
	return s.Id
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) GetMethod() *string {
	return s.Method
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) GetUrl() *string {
	return s.Url
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) SetId(v string) *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter {
	s.Id = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) SetMethod(v string) *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter {
	s.Method = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) SetProtocol(v string) *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter {
	s.Protocol = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) SetUrl(v string) *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter {
	s.Url = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) Validate() error {
	return dara.Validate(s)
}
