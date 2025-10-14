// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEventRuleTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactParameters(v []*PutEventRuleTargetsRequestContactParameters) *PutEventRuleTargetsRequest
	GetContactParameters() []*PutEventRuleTargetsRequestContactParameters
	SetFcParameters(v []*PutEventRuleTargetsRequestFcParameters) *PutEventRuleTargetsRequest
	GetFcParameters() []*PutEventRuleTargetsRequestFcParameters
	SetMnsParameters(v []*PutEventRuleTargetsRequestMnsParameters) *PutEventRuleTargetsRequest
	GetMnsParameters() []*PutEventRuleTargetsRequestMnsParameters
	SetOpenApiParameters(v []*PutEventRuleTargetsRequestOpenApiParameters) *PutEventRuleTargetsRequest
	GetOpenApiParameters() []*PutEventRuleTargetsRequestOpenApiParameters
	SetRegionId(v string) *PutEventRuleTargetsRequest
	GetRegionId() *string
	SetRuleName(v string) *PutEventRuleTargetsRequest
	GetRuleName() *string
	SetSlsParameters(v []*PutEventRuleTargetsRequestSlsParameters) *PutEventRuleTargetsRequest
	GetSlsParameters() []*PutEventRuleTargetsRequestSlsParameters
	SetWebhookParameters(v []*PutEventRuleTargetsRequestWebhookParameters) *PutEventRuleTargetsRequest
	GetWebhookParameters() []*PutEventRuleTargetsRequestWebhookParameters
}

type PutEventRuleTargetsRequest struct {
	// The information about the alert contact groups that receive alert notifications.
	ContactParameters []*PutEventRuleTargetsRequestContactParameters `json:"ContactParameters,omitempty" xml:"ContactParameters,omitempty" type:"Repeated"`
	// The information about the recipients in Function Compute.
	FcParameters []*PutEventRuleTargetsRequestFcParameters `json:"FcParameters,omitempty" xml:"FcParameters,omitempty" type:"Repeated"`
	// The notifications of Simple Message Queue (formerly MNS) (SMQ).
	MnsParameters []*PutEventRuleTargetsRequestMnsParameters `json:"MnsParameters,omitempty" xml:"MnsParameters,omitempty" type:"Repeated"`
	// The parameters of API callback notification.
	OpenApiParameters []*PutEventRuleTargetsRequestOpenApiParameters `json:"OpenApiParameters,omitempty" xml:"OpenApiParameters,omitempty" type:"Repeated"`
	RegionId          *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// testEventRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The information about the recipients in Simple Log Service.
	SlsParameters []*PutEventRuleTargetsRequestSlsParameters `json:"SlsParameters,omitempty" xml:"SlsParameters,omitempty" type:"Repeated"`
	// The information about the callback URLs that are used to receive alert notifications.
	WebhookParameters []*PutEventRuleTargetsRequestWebhookParameters `json:"WebhookParameters,omitempty" xml:"WebhookParameters,omitempty" type:"Repeated"`
}

func (s PutEventRuleTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsRequest) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsRequest) GetContactParameters() []*PutEventRuleTargetsRequestContactParameters {
	return s.ContactParameters
}

func (s *PutEventRuleTargetsRequest) GetFcParameters() []*PutEventRuleTargetsRequestFcParameters {
	return s.FcParameters
}

func (s *PutEventRuleTargetsRequest) GetMnsParameters() []*PutEventRuleTargetsRequestMnsParameters {
	return s.MnsParameters
}

func (s *PutEventRuleTargetsRequest) GetOpenApiParameters() []*PutEventRuleTargetsRequestOpenApiParameters {
	return s.OpenApiParameters
}

func (s *PutEventRuleTargetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PutEventRuleTargetsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *PutEventRuleTargetsRequest) GetSlsParameters() []*PutEventRuleTargetsRequestSlsParameters {
	return s.SlsParameters
}

func (s *PutEventRuleTargetsRequest) GetWebhookParameters() []*PutEventRuleTargetsRequestWebhookParameters {
	return s.WebhookParameters
}

func (s *PutEventRuleTargetsRequest) SetContactParameters(v []*PutEventRuleTargetsRequestContactParameters) *PutEventRuleTargetsRequest {
	s.ContactParameters = v
	return s
}

func (s *PutEventRuleTargetsRequest) SetFcParameters(v []*PutEventRuleTargetsRequestFcParameters) *PutEventRuleTargetsRequest {
	s.FcParameters = v
	return s
}

func (s *PutEventRuleTargetsRequest) SetMnsParameters(v []*PutEventRuleTargetsRequestMnsParameters) *PutEventRuleTargetsRequest {
	s.MnsParameters = v
	return s
}

func (s *PutEventRuleTargetsRequest) SetOpenApiParameters(v []*PutEventRuleTargetsRequestOpenApiParameters) *PutEventRuleTargetsRequest {
	s.OpenApiParameters = v
	return s
}

func (s *PutEventRuleTargetsRequest) SetRegionId(v string) *PutEventRuleTargetsRequest {
	s.RegionId = &v
	return s
}

func (s *PutEventRuleTargetsRequest) SetRuleName(v string) *PutEventRuleTargetsRequest {
	s.RuleName = &v
	return s
}

func (s *PutEventRuleTargetsRequest) SetSlsParameters(v []*PutEventRuleTargetsRequestSlsParameters) *PutEventRuleTargetsRequest {
	s.SlsParameters = v
	return s
}

func (s *PutEventRuleTargetsRequest) SetWebhookParameters(v []*PutEventRuleTargetsRequestWebhookParameters) *PutEventRuleTargetsRequest {
	s.WebhookParameters = v
	return s
}

func (s *PutEventRuleTargetsRequest) Validate() error {
	if s.ContactParameters != nil {
		for _, item := range s.ContactParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FcParameters != nil {
		for _, item := range s.FcParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MnsParameters != nil {
		for _, item := range s.MnsParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OpenApiParameters != nil {
		for _, item := range s.OpenApiParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SlsParameters != nil {
		for _, item := range s.SlsParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WebhookParameters != nil {
		for _, item := range s.WebhookParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutEventRuleTargetsRequestContactParameters struct {
	// The name of the alert contact group. Valid values of N: 1 to 5.
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	// The ID of the recipient that receives alert notifications. Valid values of N: 1 to 5.
	//
	// example:
	//
	// 2
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The alert notification methods. Valid values of N: 1 to 5. Valid values:
	//
	// 4: Alert notifications are sent by using DingTalk and emails.
	//
	// example:
	//
	// 3
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s PutEventRuleTargetsRequestContactParameters) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsRequestContactParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsRequestContactParameters) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *PutEventRuleTargetsRequestContactParameters) GetId() *string {
	return s.Id
}

func (s *PutEventRuleTargetsRequestContactParameters) GetLevel() *string {
	return s.Level
}

func (s *PutEventRuleTargetsRequestContactParameters) SetContactGroupName(v string) *PutEventRuleTargetsRequestContactParameters {
	s.ContactGroupName = &v
	return s
}

func (s *PutEventRuleTargetsRequestContactParameters) SetId(v string) *PutEventRuleTargetsRequestContactParameters {
	s.Id = &v
	return s
}

func (s *PutEventRuleTargetsRequestContactParameters) SetLevel(v string) *PutEventRuleTargetsRequestContactParameters {
	s.Level = &v
	return s
}

func (s *PutEventRuleTargetsRequestContactParameters) Validate() error {
	return dara.Validate(s)
}

type PutEventRuleTargetsRequestFcParameters struct {
	// The name of the function. Valid values of N: 1 to 5.
	//
	// example:
	//
	// fc-test
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The ID of the recipient that receives alert notifications. Valid values of N: 1 to 5.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region where Function Compute is deployed. Valid values of N: 1 to 5.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the Function Compute service. Valid values of N: 1 to 5.
	//
	// example:
	//
	// fc-test
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s PutEventRuleTargetsRequestFcParameters) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsRequestFcParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsRequestFcParameters) GetFunctionName() *string {
	return s.FunctionName
}

func (s *PutEventRuleTargetsRequestFcParameters) GetId() *string {
	return s.Id
}

func (s *PutEventRuleTargetsRequestFcParameters) GetRegion() *string {
	return s.Region
}

func (s *PutEventRuleTargetsRequestFcParameters) GetServiceName() *string {
	return s.ServiceName
}

func (s *PutEventRuleTargetsRequestFcParameters) SetFunctionName(v string) *PutEventRuleTargetsRequestFcParameters {
	s.FunctionName = &v
	return s
}

func (s *PutEventRuleTargetsRequestFcParameters) SetId(v string) *PutEventRuleTargetsRequestFcParameters {
	s.Id = &v
	return s
}

func (s *PutEventRuleTargetsRequestFcParameters) SetRegion(v string) *PutEventRuleTargetsRequestFcParameters {
	s.Region = &v
	return s
}

func (s *PutEventRuleTargetsRequestFcParameters) SetServiceName(v string) *PutEventRuleTargetsRequestFcParameters {
	s.ServiceName = &v
	return s
}

func (s *PutEventRuleTargetsRequestFcParameters) Validate() error {
	return dara.Validate(s)
}

type PutEventRuleTargetsRequestMnsParameters struct {
	// The ID of the recipient that receives alert notifications. Valid values of N: 1 to 5.
	//
	// example:
	//
	// 3
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the SMQ queue. Valid values of N: 1 to 5.
	//
	// example:
	//
	// queue1
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// The region for SMQ. Valid values of N: 1 to 5.
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

func (s PutEventRuleTargetsRequestMnsParameters) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsRequestMnsParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsRequestMnsParameters) GetId() *string {
	return s.Id
}

func (s *PutEventRuleTargetsRequestMnsParameters) GetQueue() *string {
	return s.Queue
}

func (s *PutEventRuleTargetsRequestMnsParameters) GetRegion() *string {
	return s.Region
}

func (s *PutEventRuleTargetsRequestMnsParameters) GetTopic() *string {
	return s.Topic
}

func (s *PutEventRuleTargetsRequestMnsParameters) SetId(v string) *PutEventRuleTargetsRequestMnsParameters {
	s.Id = &v
	return s
}

func (s *PutEventRuleTargetsRequestMnsParameters) SetQueue(v string) *PutEventRuleTargetsRequestMnsParameters {
	s.Queue = &v
	return s
}

func (s *PutEventRuleTargetsRequestMnsParameters) SetRegion(v string) *PutEventRuleTargetsRequestMnsParameters {
	s.Region = &v
	return s
}

func (s *PutEventRuleTargetsRequestMnsParameters) SetTopic(v string) *PutEventRuleTargetsRequestMnsParameters {
	s.Topic = &v
	return s
}

func (s *PutEventRuleTargetsRequestMnsParameters) Validate() error {
	return dara.Validate(s)
}

type PutEventRuleTargetsRequestOpenApiParameters struct {
	// The API name.
	//
	// example:
	//
	// PutLogs
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the resource. Valid values of N: 1 to 5. Format: `arn:acs:${Service}:${Region}:${Account}:${ResourceType}/${ResourceId}`. Fields:
	//
	// 	- Service: the code of a cloud service
	//
	// 	- Region: the region ID
	//
	// 	- Account: the ID of an Alibaba Cloud account
	//
	// 	- ResourceType: the resource type
	//
	// 	- ResourceId: the resource ID
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the recipient that receives alert notifications sent by an API callback.
	//
	// example:
	//
	// 3
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters of the alert callback. Specify the parameters in the JSON format.
	//
	// example:
	//
	// {"customField1":"value1","customField2":"$.name"}
	JsonParams *string `json:"JsonParams,omitempty" xml:"JsonParams,omitempty"`
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

func (s PutEventRuleTargetsRequestOpenApiParameters) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsRequestOpenApiParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) GetAction() *string {
	return s.Action
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) GetArn() *string {
	return s.Arn
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) GetId() *string {
	return s.Id
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) GetJsonParams() *string {
	return s.JsonParams
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) GetProduct() *string {
	return s.Product
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) GetRegion() *string {
	return s.Region
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) GetRole() *string {
	return s.Role
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) GetVersion() *string {
	return s.Version
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) SetAction(v string) *PutEventRuleTargetsRequestOpenApiParameters {
	s.Action = &v
	return s
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) SetArn(v string) *PutEventRuleTargetsRequestOpenApiParameters {
	s.Arn = &v
	return s
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) SetId(v string) *PutEventRuleTargetsRequestOpenApiParameters {
	s.Id = &v
	return s
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) SetJsonParams(v string) *PutEventRuleTargetsRequestOpenApiParameters {
	s.JsonParams = &v
	return s
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) SetProduct(v string) *PutEventRuleTargetsRequestOpenApiParameters {
	s.Product = &v
	return s
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) SetRegion(v string) *PutEventRuleTargetsRequestOpenApiParameters {
	s.Region = &v
	return s
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) SetRole(v string) *PutEventRuleTargetsRequestOpenApiParameters {
	s.Role = &v
	return s
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) SetVersion(v string) *PutEventRuleTargetsRequestOpenApiParameters {
	s.Version = &v
	return s
}

func (s *PutEventRuleTargetsRequestOpenApiParameters) Validate() error {
	return dara.Validate(s)
}

type PutEventRuleTargetsRequestSlsParameters struct {
	// The ID of the recipient that receives alert notifications. Valid values of N: 1 to 5.
	//
	// example:
	//
	// 5
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the Simple Log Service Logstore. Valid values of N: 1 to 5.
	//
	// example:
	//
	// testlogstore
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The name of the Simple Log Service project. Valid values of N: 1 to 5.
	//
	// example:
	//
	// testproject
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The region where Simple Log Service is deployed. Valid values of N: 1 to 5.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s PutEventRuleTargetsRequestSlsParameters) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsRequestSlsParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsRequestSlsParameters) GetId() *string {
	return s.Id
}

func (s *PutEventRuleTargetsRequestSlsParameters) GetLogStore() *string {
	return s.LogStore
}

func (s *PutEventRuleTargetsRequestSlsParameters) GetProject() *string {
	return s.Project
}

func (s *PutEventRuleTargetsRequestSlsParameters) GetRegion() *string {
	return s.Region
}

func (s *PutEventRuleTargetsRequestSlsParameters) SetId(v string) *PutEventRuleTargetsRequestSlsParameters {
	s.Id = &v
	return s
}

func (s *PutEventRuleTargetsRequestSlsParameters) SetLogStore(v string) *PutEventRuleTargetsRequestSlsParameters {
	s.LogStore = &v
	return s
}

func (s *PutEventRuleTargetsRequestSlsParameters) SetProject(v string) *PutEventRuleTargetsRequestSlsParameters {
	s.Project = &v
	return s
}

func (s *PutEventRuleTargetsRequestSlsParameters) SetRegion(v string) *PutEventRuleTargetsRequestSlsParameters {
	s.Region = &v
	return s
}

func (s *PutEventRuleTargetsRequestSlsParameters) Validate() error {
	return dara.Validate(s)
}

type PutEventRuleTargetsRequestWebhookParameters struct {
	// The ID of the recipient that receives alert notifications. Valid values of N: 1 to 5.
	//
	// example:
	//
	// 4
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The HTTP request method. Valid values of N: 1 to 5.
	//
	// Valid values: GET and POST.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The name of the protocol. Valid values of N: 1 to 5. Valid values:
	//
	// 	- http
	//
	// 	- telnet
	//
	// 	- ping
	//
	// example:
	//
	// http
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The callback URL. Valid values of N: 1 to 5.
	//
	// example:
	//
	// http://www.aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PutEventRuleTargetsRequestWebhookParameters) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsRequestWebhookParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsRequestWebhookParameters) GetId() *string {
	return s.Id
}

func (s *PutEventRuleTargetsRequestWebhookParameters) GetMethod() *string {
	return s.Method
}

func (s *PutEventRuleTargetsRequestWebhookParameters) GetProtocol() *string {
	return s.Protocol
}

func (s *PutEventRuleTargetsRequestWebhookParameters) GetUrl() *string {
	return s.Url
}

func (s *PutEventRuleTargetsRequestWebhookParameters) SetId(v string) *PutEventRuleTargetsRequestWebhookParameters {
	s.Id = &v
	return s
}

func (s *PutEventRuleTargetsRequestWebhookParameters) SetMethod(v string) *PutEventRuleTargetsRequestWebhookParameters {
	s.Method = &v
	return s
}

func (s *PutEventRuleTargetsRequestWebhookParameters) SetProtocol(v string) *PutEventRuleTargetsRequestWebhookParameters {
	s.Protocol = &v
	return s
}

func (s *PutEventRuleTargetsRequestWebhookParameters) SetUrl(v string) *PutEventRuleTargetsRequestWebhookParameters {
	s.Url = &v
	return s
}

func (s *PutEventRuleTargetsRequestWebhookParameters) Validate() error {
	return dara.Validate(s)
}
