// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertActionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertActionId(v string) *CreateAlertActionRequest
	GetAlertActionId() *string
	SetAlertActionName(v string) *CreateAlertActionRequest
	GetAlertActionName() *string
	SetEbParam(v *CreateAlertActionRequestEbParam) *CreateAlertActionRequest
	GetEbParam() *CreateAlertActionRequestEbParam
	SetEssParam(v *CreateAlertActionRequestEssParam) *CreateAlertActionRequest
	GetEssParam() *CreateAlertActionRequestEssParam
	SetFc3Param(v *CreateAlertActionRequestFc3Param) *CreateAlertActionRequest
	GetFc3Param() *CreateAlertActionRequestFc3Param
	SetFcParam(v *CreateAlertActionRequestFcParam) *CreateAlertActionRequest
	GetFcParam() *CreateAlertActionRequestFcParam
	SetMnsParam(v *CreateAlertActionRequestMnsParam) *CreateAlertActionRequest
	GetMnsParam() *CreateAlertActionRequestMnsParam
	SetPagerDutyParam(v *CreateAlertActionRequestPagerDutyParam) *CreateAlertActionRequest
	GetPagerDutyParam() *CreateAlertActionRequestPagerDutyParam
	SetSlsParam(v *CreateAlertActionRequestSlsParam) *CreateAlertActionRequest
	GetSlsParam() *CreateAlertActionRequestSlsParam
	SetType(v string) *CreateAlertActionRequest
	GetType() *string
	SetWebhookParam(v *CreateAlertActionRequestWebhookParam) *CreateAlertActionRequest
	GetWebhookParam() *CreateAlertActionRequestWebhookParam
}

type CreateAlertActionRequest struct {
	// The unique identifier of the alert action integration.
	//
	// example:
	//
	// testId
	AlertActionId *string `json:"alertActionId,omitempty" xml:"alertActionId,omitempty"`
	// The name of the alert action integration.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	AlertActionName *string `json:"alertActionName,omitempty" xml:"alertActionName,omitempty"`
	// The event bus parameters.
	EbParam *CreateAlertActionRequestEbParam `json:"ebParam,omitempty" xml:"ebParam,omitempty" type:"Struct"`
	// The elastic scaling parameters.
	EssParam *CreateAlertActionRequestEssParam `json:"essParam,omitempty" xml:"essParam,omitempty" type:"Struct"`
	// The Function Compute 3.0 parameters.
	Fc3Param *CreateAlertActionRequestFc3Param `json:"fc3Param,omitempty" xml:"fc3Param,omitempty" type:"Struct"`
	// The Function Compute parameters.
	FcParam *CreateAlertActionRequestFcParam `json:"fcParam,omitempty" xml:"fcParam,omitempty" type:"Struct"`
	// The Simple Message Queue (formerly MNS) (MSMQ) parameters.
	MnsParam *CreateAlertActionRequestMnsParam `json:"mnsParam,omitempty" xml:"mnsParam,omitempty" type:"Struct"`
	// The PagerDuty parameters.
	PagerDutyParam *CreateAlertActionRequestPagerDutyParam `json:"pagerDutyParam,omitempty" xml:"pagerDutyParam,omitempty" type:"Struct"`
	// The Simple Log Service parameters.
	SlsParam *CreateAlertActionRequestSlsParam `json:"slsParam,omitempty" xml:"slsParam,omitempty" type:"Struct"`
	// The type of the alert action integration.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The webhook parameters.
	WebhookParam *CreateAlertActionRequestWebhookParam `json:"webhookParam,omitempty" xml:"webhookParam,omitempty" type:"Struct"`
}

func (s CreateAlertActionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertActionRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertActionRequest) GetAlertActionId() *string {
	return s.AlertActionId
}

func (s *CreateAlertActionRequest) GetAlertActionName() *string {
	return s.AlertActionName
}

func (s *CreateAlertActionRequest) GetEbParam() *CreateAlertActionRequestEbParam {
	return s.EbParam
}

func (s *CreateAlertActionRequest) GetEssParam() *CreateAlertActionRequestEssParam {
	return s.EssParam
}

func (s *CreateAlertActionRequest) GetFc3Param() *CreateAlertActionRequestFc3Param {
	return s.Fc3Param
}

func (s *CreateAlertActionRequest) GetFcParam() *CreateAlertActionRequestFcParam {
	return s.FcParam
}

func (s *CreateAlertActionRequest) GetMnsParam() *CreateAlertActionRequestMnsParam {
	return s.MnsParam
}

func (s *CreateAlertActionRequest) GetPagerDutyParam() *CreateAlertActionRequestPagerDutyParam {
	return s.PagerDutyParam
}

func (s *CreateAlertActionRequest) GetSlsParam() *CreateAlertActionRequestSlsParam {
	return s.SlsParam
}

func (s *CreateAlertActionRequest) GetType() *string {
	return s.Type
}

func (s *CreateAlertActionRequest) GetWebhookParam() *CreateAlertActionRequestWebhookParam {
	return s.WebhookParam
}

func (s *CreateAlertActionRequest) SetAlertActionId(v string) *CreateAlertActionRequest {
	s.AlertActionId = &v
	return s
}

func (s *CreateAlertActionRequest) SetAlertActionName(v string) *CreateAlertActionRequest {
	s.AlertActionName = &v
	return s
}

func (s *CreateAlertActionRequest) SetEbParam(v *CreateAlertActionRequestEbParam) *CreateAlertActionRequest {
	s.EbParam = v
	return s
}

func (s *CreateAlertActionRequest) SetEssParam(v *CreateAlertActionRequestEssParam) *CreateAlertActionRequest {
	s.EssParam = v
	return s
}

func (s *CreateAlertActionRequest) SetFc3Param(v *CreateAlertActionRequestFc3Param) *CreateAlertActionRequest {
	s.Fc3Param = v
	return s
}

func (s *CreateAlertActionRequest) SetFcParam(v *CreateAlertActionRequestFcParam) *CreateAlertActionRequest {
	s.FcParam = v
	return s
}

func (s *CreateAlertActionRequest) SetMnsParam(v *CreateAlertActionRequestMnsParam) *CreateAlertActionRequest {
	s.MnsParam = v
	return s
}

func (s *CreateAlertActionRequest) SetPagerDutyParam(v *CreateAlertActionRequestPagerDutyParam) *CreateAlertActionRequest {
	s.PagerDutyParam = v
	return s
}

func (s *CreateAlertActionRequest) SetSlsParam(v *CreateAlertActionRequestSlsParam) *CreateAlertActionRequest {
	s.SlsParam = v
	return s
}

func (s *CreateAlertActionRequest) SetType(v string) *CreateAlertActionRequest {
	s.Type = &v
	return s
}

func (s *CreateAlertActionRequest) SetWebhookParam(v *CreateAlertActionRequestWebhookParam) *CreateAlertActionRequest {
	s.WebhookParam = v
	return s
}

func (s *CreateAlertActionRequest) Validate() error {
	if s.EbParam != nil {
		if err := s.EbParam.Validate(); err != nil {
			return err
		}
	}
	if s.EssParam != nil {
		if err := s.EssParam.Validate(); err != nil {
			return err
		}
	}
	if s.Fc3Param != nil {
		if err := s.Fc3Param.Validate(); err != nil {
			return err
		}
	}
	if s.FcParam != nil {
		if err := s.FcParam.Validate(); err != nil {
			return err
		}
	}
	if s.MnsParam != nil {
		if err := s.MnsParam.Validate(); err != nil {
			return err
		}
	}
	if s.PagerDutyParam != nil {
		if err := s.PagerDutyParam.Validate(); err != nil {
			return err
		}
	}
	if s.SlsParam != nil {
		if err := s.SlsParam.Validate(); err != nil {
			return err
		}
	}
	if s.WebhookParam != nil {
		if err := s.WebhookParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAlertActionRequestEbParam struct {
	// The event source.
	//
	// example:
	//
	// test
	EbSource *string `json:"ebSource,omitempty" xml:"ebSource,omitempty"`
	// The name of the event bus.
	//
	// example:
	//
	// test
	EventBusName *string `json:"eventBusName,omitempty" xml:"eventBusName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The subject.
	//
	// example:
	//
	// test
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s CreateAlertActionRequestEbParam) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertActionRequestEbParam) GoString() string {
	return s.String()
}

func (s *CreateAlertActionRequestEbParam) GetEbSource() *string {
	return s.EbSource
}

func (s *CreateAlertActionRequestEbParam) GetEventBusName() *string {
	return s.EventBusName
}

func (s *CreateAlertActionRequestEbParam) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAlertActionRequestEbParam) GetSubject() *string {
	return s.Subject
}

func (s *CreateAlertActionRequestEbParam) SetEbSource(v string) *CreateAlertActionRequestEbParam {
	s.EbSource = &v
	return s
}

func (s *CreateAlertActionRequestEbParam) SetEventBusName(v string) *CreateAlertActionRequestEbParam {
	s.EventBusName = &v
	return s
}

func (s *CreateAlertActionRequestEbParam) SetRegionId(v string) *CreateAlertActionRequestEbParam {
	s.RegionId = &v
	return s
}

func (s *CreateAlertActionRequestEbParam) SetSubject(v string) *CreateAlertActionRequestEbParam {
	s.Subject = &v
	return s
}

func (s *CreateAlertActionRequestEbParam) Validate() error {
	return dara.Validate(s)
}

type CreateAlertActionRequestEssParam struct {
	// The ID of the elastic scaling group.
	//
	// example:
	//
	// testId
	EssGroupId *string `json:"essGroupId,omitempty" xml:"essGroupId,omitempty"`
	// The ID of the elastic scaling rule.
	//
	// example:
	//
	// testId
	EssRuleId *string `json:"essRuleId,omitempty" xml:"essRuleId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateAlertActionRequestEssParam) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertActionRequestEssParam) GoString() string {
	return s.String()
}

func (s *CreateAlertActionRequestEssParam) GetEssGroupId() *string {
	return s.EssGroupId
}

func (s *CreateAlertActionRequestEssParam) GetEssRuleId() *string {
	return s.EssRuleId
}

func (s *CreateAlertActionRequestEssParam) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAlertActionRequestEssParam) SetEssGroupId(v string) *CreateAlertActionRequestEssParam {
	s.EssGroupId = &v
	return s
}

func (s *CreateAlertActionRequestEssParam) SetEssRuleId(v string) *CreateAlertActionRequestEssParam {
	s.EssRuleId = &v
	return s
}

func (s *CreateAlertActionRequestEssParam) SetRegionId(v string) *CreateAlertActionRequestEssParam {
	s.RegionId = &v
	return s
}

func (s *CreateAlertActionRequestEssParam) Validate() error {
	return dara.Validate(s)
}

type CreateAlertActionRequestFc3Param struct {
	// The name of the function in Function Compute.
	//
	// example:
	//
	// test
	Function *string `json:"function,omitempty" xml:"function,omitempty"`
	// The version or alias of the function.
	//
	// example:
	//
	// test
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-nanjing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateAlertActionRequestFc3Param) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertActionRequestFc3Param) GoString() string {
	return s.String()
}

func (s *CreateAlertActionRequestFc3Param) GetFunction() *string {
	return s.Function
}

func (s *CreateAlertActionRequestFc3Param) GetQualifier() *string {
	return s.Qualifier
}

func (s *CreateAlertActionRequestFc3Param) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAlertActionRequestFc3Param) SetFunction(v string) *CreateAlertActionRequestFc3Param {
	s.Function = &v
	return s
}

func (s *CreateAlertActionRequestFc3Param) SetQualifier(v string) *CreateAlertActionRequestFc3Param {
	s.Qualifier = &v
	return s
}

func (s *CreateAlertActionRequestFc3Param) SetRegionId(v string) *CreateAlertActionRequestFc3Param {
	s.RegionId = &v
	return s
}

func (s *CreateAlertActionRequestFc3Param) Validate() error {
	return dara.Validate(s)
}

type CreateAlertActionRequestFcParam struct {
	// The name of the function in Function Compute.
	//
	// example:
	//
	// testFunction
	Function *string `json:"function,omitempty" xml:"function,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The service name of Function Compute.
	//
	// example:
	//
	// testService
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
}

func (s CreateAlertActionRequestFcParam) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertActionRequestFcParam) GoString() string {
	return s.String()
}

func (s *CreateAlertActionRequestFcParam) GetFunction() *string {
	return s.Function
}

func (s *CreateAlertActionRequestFcParam) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAlertActionRequestFcParam) GetService() *string {
	return s.Service
}

func (s *CreateAlertActionRequestFcParam) SetFunction(v string) *CreateAlertActionRequestFcParam {
	s.Function = &v
	return s
}

func (s *CreateAlertActionRequestFcParam) SetRegionId(v string) *CreateAlertActionRequestFcParam {
	s.RegionId = &v
	return s
}

func (s *CreateAlertActionRequestFcParam) SetService(v string) *CreateAlertActionRequestFcParam {
	s.Service = &v
	return s
}

func (s *CreateAlertActionRequestFcParam) Validate() error {
	return dara.Validate(s)
}

type CreateAlertActionRequestMnsParam struct {
	// The resource type of Simple Message Queue (formerly MNS) (MSMQ).
	//
	// example:
	//
	// queue
	MnsType *string `json:"mnsType,omitempty" xml:"mnsType,omitempty"`
	// The name of the resource.
	//
	// - If the resource type is queue, the resource name is the queue name.
	//
	// - If the resource type is topic, the resource name is the topic name.
	//
	// example:
	//
	// testName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateAlertActionRequestMnsParam) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertActionRequestMnsParam) GoString() string {
	return s.String()
}

func (s *CreateAlertActionRequestMnsParam) GetMnsType() *string {
	return s.MnsType
}

func (s *CreateAlertActionRequestMnsParam) GetName() *string {
	return s.Name
}

func (s *CreateAlertActionRequestMnsParam) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAlertActionRequestMnsParam) SetMnsType(v string) *CreateAlertActionRequestMnsParam {
	s.MnsType = &v
	return s
}

func (s *CreateAlertActionRequestMnsParam) SetName(v string) *CreateAlertActionRequestMnsParam {
	s.Name = &v
	return s
}

func (s *CreateAlertActionRequestMnsParam) SetRegionId(v string) *CreateAlertActionRequestMnsParam {
	s.RegionId = &v
	return s
}

func (s *CreateAlertActionRequestMnsParam) Validate() error {
	return dara.Validate(s)
}

type CreateAlertActionRequestPagerDutyParam struct {
	// The integration key of PagerDuty.
	//
	// example:
	//
	// fds3wr54rvcxfr4543t6435
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The integration webhook URL of PagerDuty. Both V1 and V2 are supported.
	//
	// example:
	//
	// https://events.pagerduty.com/v2/enqueue
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateAlertActionRequestPagerDutyParam) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertActionRequestPagerDutyParam) GoString() string {
	return s.String()
}

func (s *CreateAlertActionRequestPagerDutyParam) GetKey() *string {
	return s.Key
}

func (s *CreateAlertActionRequestPagerDutyParam) GetUrl() *string {
	return s.Url
}

func (s *CreateAlertActionRequestPagerDutyParam) SetKey(v string) *CreateAlertActionRequestPagerDutyParam {
	s.Key = &v
	return s
}

func (s *CreateAlertActionRequestPagerDutyParam) SetUrl(v string) *CreateAlertActionRequestPagerDutyParam {
	s.Url = &v
	return s
}

func (s *CreateAlertActionRequestPagerDutyParam) Validate() error {
	return dara.Validate(s)
}

type CreateAlertActionRequestSlsParam struct {
	// The name of the Logstore in Simple Log Service.
	//
	// example:
	//
	// testLogstore。
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the project in Simple Log Service.
	//
	// example:
	//
	// testProject
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateAlertActionRequestSlsParam) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertActionRequestSlsParam) GoString() string {
	return s.String()
}

func (s *CreateAlertActionRequestSlsParam) GetLogstore() *string {
	return s.Logstore
}

func (s *CreateAlertActionRequestSlsParam) GetProject() *string {
	return s.Project
}

func (s *CreateAlertActionRequestSlsParam) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAlertActionRequestSlsParam) SetLogstore(v string) *CreateAlertActionRequestSlsParam {
	s.Logstore = &v
	return s
}

func (s *CreateAlertActionRequestSlsParam) SetProject(v string) *CreateAlertActionRequestSlsParam {
	s.Project = &v
	return s
}

func (s *CreateAlertActionRequestSlsParam) SetRegionId(v string) *CreateAlertActionRequestSlsParam {
	s.RegionId = &v
	return s
}

func (s *CreateAlertActionRequestSlsParam) Validate() error {
	return dara.Validate(s)
}

type CreateAlertActionRequestWebhookParam struct {
	// The data format. This parameter takes effect only when the request method is POST.
	//
	// example:
	//
	// JSON
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The request header information.
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The webhook request method.
	//
	// example:
	//
	// POST
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// The alert callback URL.
	//
	// example:
	//
	// http://www.test.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateAlertActionRequestWebhookParam) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertActionRequestWebhookParam) GoString() string {
	return s.String()
}

func (s *CreateAlertActionRequestWebhookParam) GetContentType() *string {
	return s.ContentType
}

func (s *CreateAlertActionRequestWebhookParam) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlertActionRequestWebhookParam) GetMethod() *string {
	return s.Method
}

func (s *CreateAlertActionRequestWebhookParam) GetUrl() *string {
	return s.Url
}

func (s *CreateAlertActionRequestWebhookParam) SetContentType(v string) *CreateAlertActionRequestWebhookParam {
	s.ContentType = &v
	return s
}

func (s *CreateAlertActionRequestWebhookParam) SetHeaders(v map[string]*string) *CreateAlertActionRequestWebhookParam {
	s.Headers = v
	return s
}

func (s *CreateAlertActionRequestWebhookParam) SetMethod(v string) *CreateAlertActionRequestWebhookParam {
	s.Method = &v
	return s
}

func (s *CreateAlertActionRequestWebhookParam) SetUrl(v string) *CreateAlertActionRequestWebhookParam {
	s.Url = &v
	return s
}

func (s *CreateAlertActionRequestWebhookParam) Validate() error {
	return dara.Validate(s)
}
