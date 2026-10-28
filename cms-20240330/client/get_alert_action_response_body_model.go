// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertActions(v *GetAlertActionResponseBodyAlertActions) *GetAlertActionResponseBody
	GetAlertActions() *GetAlertActionResponseBodyAlertActions
	SetRequestId(v string) *GetAlertActionResponseBody
	GetRequestId() *string
}

type GetAlertActionResponseBody struct {
	// The list of alert action integration configurations.
	AlertActions *GetAlertActionResponseBodyAlertActions `json:"alertActions,omitempty" xml:"alertActions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAlertActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlertActionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlertActionResponseBody) GetAlertActions() *GetAlertActionResponseBodyAlertActions {
	return s.AlertActions
}

func (s *GetAlertActionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlertActionResponseBody) SetAlertActions(v *GetAlertActionResponseBodyAlertActions) *GetAlertActionResponseBody {
	s.AlertActions = v
	return s
}

func (s *GetAlertActionResponseBody) SetRequestId(v string) *GetAlertActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlertActionResponseBody) Validate() error {
	if s.AlertActions != nil {
		if err := s.AlertActions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlertActionResponseBodyAlertActions struct {
	// The unique identifier of the alert action integration.
	//
	// example:
	//
	// testId
	AlertActionId *string `json:"alertActionId,omitempty" xml:"alertActionId,omitempty"`
	// The name of the alert action integration.
	//
	// example:
	//
	// testName
	AlertActionName *string `json:"alertActionName,omitempty" xml:"alertActionName,omitempty"`
	// The event bus parameters.
	EbParam *GetAlertActionResponseBodyAlertActionsEbParam `json:"ebParam,omitempty" xml:"ebParam,omitempty" type:"Struct"`
	// The elastic scaling parameters.
	EssParam *GetAlertActionResponseBodyAlertActionsEssParam `json:"essParam,omitempty" xml:"essParam,omitempty" type:"Struct"`
	// The Function Compute 3.0 parameters.
	Fc3Param *GetAlertActionResponseBodyAlertActionsFc3Param `json:"fc3Param,omitempty" xml:"fc3Param,omitempty" type:"Struct"`
	// The Function Compute parameters.
	FcParam *GetAlertActionResponseBodyAlertActionsFcParam `json:"fcParam,omitempty" xml:"fcParam,omitempty" type:"Struct"`
	// The Simple Message Queue (formerly MNS) parameters.
	MnsParam *GetAlertActionResponseBodyAlertActionsMnsParam `json:"mnsParam,omitempty" xml:"mnsParam,omitempty" type:"Struct"`
	// The PagerDuty parameters.
	PagerDutyParam *GetAlertActionResponseBodyAlertActionsPagerDutyParam `json:"pagerDutyParam,omitempty" xml:"pagerDutyParam,omitempty" type:"Struct"`
	// The Simple Log Service parameters.
	SlsParam *GetAlertActionResponseBodyAlertActionsSlsParam `json:"slsParam,omitempty" xml:"slsParam,omitempty" type:"Struct"`
	// The type of the alert action integration.
	//
	// example:
	//
	// FC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The webhook parameters.
	WebhookParam *GetAlertActionResponseBodyAlertActionsWebhookParam `json:"webhookParam,omitempty" xml:"webhookParam,omitempty" type:"Struct"`
}

func (s GetAlertActionResponseBodyAlertActions) String() string {
	return dara.Prettify(s)
}

func (s GetAlertActionResponseBodyAlertActions) GoString() string {
	return s.String()
}

func (s *GetAlertActionResponseBodyAlertActions) GetAlertActionId() *string {
	return s.AlertActionId
}

func (s *GetAlertActionResponseBodyAlertActions) GetAlertActionName() *string {
	return s.AlertActionName
}

func (s *GetAlertActionResponseBodyAlertActions) GetEbParam() *GetAlertActionResponseBodyAlertActionsEbParam {
	return s.EbParam
}

func (s *GetAlertActionResponseBodyAlertActions) GetEssParam() *GetAlertActionResponseBodyAlertActionsEssParam {
	return s.EssParam
}

func (s *GetAlertActionResponseBodyAlertActions) GetFc3Param() *GetAlertActionResponseBodyAlertActionsFc3Param {
	return s.Fc3Param
}

func (s *GetAlertActionResponseBodyAlertActions) GetFcParam() *GetAlertActionResponseBodyAlertActionsFcParam {
	return s.FcParam
}

func (s *GetAlertActionResponseBodyAlertActions) GetMnsParam() *GetAlertActionResponseBodyAlertActionsMnsParam {
	return s.MnsParam
}

func (s *GetAlertActionResponseBodyAlertActions) GetPagerDutyParam() *GetAlertActionResponseBodyAlertActionsPagerDutyParam {
	return s.PagerDutyParam
}

func (s *GetAlertActionResponseBodyAlertActions) GetSlsParam() *GetAlertActionResponseBodyAlertActionsSlsParam {
	return s.SlsParam
}

func (s *GetAlertActionResponseBodyAlertActions) GetType() *string {
	return s.Type
}

func (s *GetAlertActionResponseBodyAlertActions) GetWebhookParam() *GetAlertActionResponseBodyAlertActionsWebhookParam {
	return s.WebhookParam
}

func (s *GetAlertActionResponseBodyAlertActions) SetAlertActionId(v string) *GetAlertActionResponseBodyAlertActions {
	s.AlertActionId = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActions) SetAlertActionName(v string) *GetAlertActionResponseBodyAlertActions {
	s.AlertActionName = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActions) SetEbParam(v *GetAlertActionResponseBodyAlertActionsEbParam) *GetAlertActionResponseBodyAlertActions {
	s.EbParam = v
	return s
}

func (s *GetAlertActionResponseBodyAlertActions) SetEssParam(v *GetAlertActionResponseBodyAlertActionsEssParam) *GetAlertActionResponseBodyAlertActions {
	s.EssParam = v
	return s
}

func (s *GetAlertActionResponseBodyAlertActions) SetFc3Param(v *GetAlertActionResponseBodyAlertActionsFc3Param) *GetAlertActionResponseBodyAlertActions {
	s.Fc3Param = v
	return s
}

func (s *GetAlertActionResponseBodyAlertActions) SetFcParam(v *GetAlertActionResponseBodyAlertActionsFcParam) *GetAlertActionResponseBodyAlertActions {
	s.FcParam = v
	return s
}

func (s *GetAlertActionResponseBodyAlertActions) SetMnsParam(v *GetAlertActionResponseBodyAlertActionsMnsParam) *GetAlertActionResponseBodyAlertActions {
	s.MnsParam = v
	return s
}

func (s *GetAlertActionResponseBodyAlertActions) SetPagerDutyParam(v *GetAlertActionResponseBodyAlertActionsPagerDutyParam) *GetAlertActionResponseBodyAlertActions {
	s.PagerDutyParam = v
	return s
}

func (s *GetAlertActionResponseBodyAlertActions) SetSlsParam(v *GetAlertActionResponseBodyAlertActionsSlsParam) *GetAlertActionResponseBodyAlertActions {
	s.SlsParam = v
	return s
}

func (s *GetAlertActionResponseBodyAlertActions) SetType(v string) *GetAlertActionResponseBodyAlertActions {
	s.Type = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActions) SetWebhookParam(v *GetAlertActionResponseBodyAlertActionsWebhookParam) *GetAlertActionResponseBodyAlertActions {
	s.WebhookParam = v
	return s
}

func (s *GetAlertActionResponseBodyAlertActions) Validate() error {
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

type GetAlertActionResponseBodyAlertActionsEbParam struct {
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

func (s GetAlertActionResponseBodyAlertActionsEbParam) String() string {
	return dara.Prettify(s)
}

func (s GetAlertActionResponseBodyAlertActionsEbParam) GoString() string {
	return s.String()
}

func (s *GetAlertActionResponseBodyAlertActionsEbParam) GetEbSource() *string {
	return s.EbSource
}

func (s *GetAlertActionResponseBodyAlertActionsEbParam) GetEventBusName() *string {
	return s.EventBusName
}

func (s *GetAlertActionResponseBodyAlertActionsEbParam) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAlertActionResponseBodyAlertActionsEbParam) GetSubject() *string {
	return s.Subject
}

func (s *GetAlertActionResponseBodyAlertActionsEbParam) SetEbSource(v string) *GetAlertActionResponseBodyAlertActionsEbParam {
	s.EbSource = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsEbParam) SetEventBusName(v string) *GetAlertActionResponseBodyAlertActionsEbParam {
	s.EventBusName = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsEbParam) SetRegionId(v string) *GetAlertActionResponseBodyAlertActionsEbParam {
	s.RegionId = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsEbParam) SetSubject(v string) *GetAlertActionResponseBodyAlertActionsEbParam {
	s.Subject = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsEbParam) Validate() error {
	return dara.Validate(s)
}

type GetAlertActionResponseBodyAlertActionsEssParam struct {
	// The ID of the elastic scaling group.
	//
	// example:
	//
	// test
	EssGroupId *string `json:"essGroupId,omitempty" xml:"essGroupId,omitempty"`
	// The ID of the elastic scaling rule.
	//
	// example:
	//
	// test
	EssRuleId *string `json:"essRuleId,omitempty" xml:"essRuleId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetAlertActionResponseBodyAlertActionsEssParam) String() string {
	return dara.Prettify(s)
}

func (s GetAlertActionResponseBodyAlertActionsEssParam) GoString() string {
	return s.String()
}

func (s *GetAlertActionResponseBodyAlertActionsEssParam) GetEssGroupId() *string {
	return s.EssGroupId
}

func (s *GetAlertActionResponseBodyAlertActionsEssParam) GetEssRuleId() *string {
	return s.EssRuleId
}

func (s *GetAlertActionResponseBodyAlertActionsEssParam) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAlertActionResponseBodyAlertActionsEssParam) SetEssGroupId(v string) *GetAlertActionResponseBodyAlertActionsEssParam {
	s.EssGroupId = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsEssParam) SetEssRuleId(v string) *GetAlertActionResponseBodyAlertActionsEssParam {
	s.EssRuleId = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsEssParam) SetRegionId(v string) *GetAlertActionResponseBodyAlertActionsEssParam {
	s.RegionId = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsEssParam) Validate() error {
	return dara.Validate(s)
}

type GetAlertActionResponseBodyAlertActionsFc3Param struct {
	// The name of the function.
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
	// cn-qingdao
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetAlertActionResponseBodyAlertActionsFc3Param) String() string {
	return dara.Prettify(s)
}

func (s GetAlertActionResponseBodyAlertActionsFc3Param) GoString() string {
	return s.String()
}

func (s *GetAlertActionResponseBodyAlertActionsFc3Param) GetFunction() *string {
	return s.Function
}

func (s *GetAlertActionResponseBodyAlertActionsFc3Param) GetQualifier() *string {
	return s.Qualifier
}

func (s *GetAlertActionResponseBodyAlertActionsFc3Param) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAlertActionResponseBodyAlertActionsFc3Param) SetFunction(v string) *GetAlertActionResponseBodyAlertActionsFc3Param {
	s.Function = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsFc3Param) SetQualifier(v string) *GetAlertActionResponseBodyAlertActionsFc3Param {
	s.Qualifier = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsFc3Param) SetRegionId(v string) *GetAlertActionResponseBodyAlertActionsFc3Param {
	s.RegionId = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsFc3Param) Validate() error {
	return dara.Validate(s)
}

type GetAlertActionResponseBodyAlertActionsFcParam struct {
	// The name of the function in Function Compute.
	//
	// example:
	//
	// test
	Function *string `json:"function,omitempty" xml:"function,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The service name in Function Compute.
	//
	// example:
	//
	// test
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
}

func (s GetAlertActionResponseBodyAlertActionsFcParam) String() string {
	return dara.Prettify(s)
}

func (s GetAlertActionResponseBodyAlertActionsFcParam) GoString() string {
	return s.String()
}

func (s *GetAlertActionResponseBodyAlertActionsFcParam) GetFunction() *string {
	return s.Function
}

func (s *GetAlertActionResponseBodyAlertActionsFcParam) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAlertActionResponseBodyAlertActionsFcParam) GetService() *string {
	return s.Service
}

func (s *GetAlertActionResponseBodyAlertActionsFcParam) SetFunction(v string) *GetAlertActionResponseBodyAlertActionsFcParam {
	s.Function = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsFcParam) SetRegionId(v string) *GetAlertActionResponseBodyAlertActionsFcParam {
	s.RegionId = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsFcParam) SetService(v string) *GetAlertActionResponseBodyAlertActionsFcParam {
	s.Service = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsFcParam) Validate() error {
	return dara.Validate(s)
}

type GetAlertActionResponseBodyAlertActionsMnsParam struct {
	// The resource type of Simple Message Queue (formerly MNS).
	//
	// example:
	//
	// queue
	MnsType *string `json:"mnsType,omitempty" xml:"mnsType,omitempty"`
	// The name of the resource.
	//
	// - If the resource type is queue, the resource name is the name of the queue.
	//
	// - If the resource type is topic, the resource name is the name of the topic.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetAlertActionResponseBodyAlertActionsMnsParam) String() string {
	return dara.Prettify(s)
}

func (s GetAlertActionResponseBodyAlertActionsMnsParam) GoString() string {
	return s.String()
}

func (s *GetAlertActionResponseBodyAlertActionsMnsParam) GetMnsType() *string {
	return s.MnsType
}

func (s *GetAlertActionResponseBodyAlertActionsMnsParam) GetName() *string {
	return s.Name
}

func (s *GetAlertActionResponseBodyAlertActionsMnsParam) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAlertActionResponseBodyAlertActionsMnsParam) SetMnsType(v string) *GetAlertActionResponseBodyAlertActionsMnsParam {
	s.MnsType = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsMnsParam) SetName(v string) *GetAlertActionResponseBodyAlertActionsMnsParam {
	s.Name = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsMnsParam) SetRegionId(v string) *GetAlertActionResponseBodyAlertActionsMnsParam {
	s.RegionId = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsMnsParam) Validate() error {
	return dara.Validate(s)
}

type GetAlertActionResponseBodyAlertActionsPagerDutyParam struct {
	// The integration key of PagerDuty.
	//
	// example:
	//
	// EWR4WR4T4GREG4ET453T5Y
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The integration webhook URL for PagerDuty. V1 and V2 are supported.
	//
	// example:
	//
	// https://events.pagerduty.com/v2/enqueue
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAlertActionResponseBodyAlertActionsPagerDutyParam) String() string {
	return dara.Prettify(s)
}

func (s GetAlertActionResponseBodyAlertActionsPagerDutyParam) GoString() string {
	return s.String()
}

func (s *GetAlertActionResponseBodyAlertActionsPagerDutyParam) GetKey() *string {
	return s.Key
}

func (s *GetAlertActionResponseBodyAlertActionsPagerDutyParam) GetUrl() *string {
	return s.Url
}

func (s *GetAlertActionResponseBodyAlertActionsPagerDutyParam) SetKey(v string) *GetAlertActionResponseBodyAlertActionsPagerDutyParam {
	s.Key = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsPagerDutyParam) SetUrl(v string) *GetAlertActionResponseBodyAlertActionsPagerDutyParam {
	s.Url = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsPagerDutyParam) Validate() error {
	return dara.Validate(s)
}

type GetAlertActionResponseBodyAlertActionsSlsParam struct {
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// test
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// test
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetAlertActionResponseBodyAlertActionsSlsParam) String() string {
	return dara.Prettify(s)
}

func (s GetAlertActionResponseBodyAlertActionsSlsParam) GoString() string {
	return s.String()
}

func (s *GetAlertActionResponseBodyAlertActionsSlsParam) GetLogstore() *string {
	return s.Logstore
}

func (s *GetAlertActionResponseBodyAlertActionsSlsParam) GetProject() *string {
	return s.Project
}

func (s *GetAlertActionResponseBodyAlertActionsSlsParam) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAlertActionResponseBodyAlertActionsSlsParam) SetLogstore(v string) *GetAlertActionResponseBodyAlertActionsSlsParam {
	s.Logstore = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsSlsParam) SetProject(v string) *GetAlertActionResponseBodyAlertActionsSlsParam {
	s.Project = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsSlsParam) SetRegionId(v string) *GetAlertActionResponseBodyAlertActionsSlsParam {
	s.RegionId = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsSlsParam) Validate() error {
	return dara.Validate(s)
}

type GetAlertActionResponseBodyAlertActionsWebhookParam struct {
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
	// The URL for the alert callback.
	//
	// example:
	//
	// http://www.test.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAlertActionResponseBodyAlertActionsWebhookParam) String() string {
	return dara.Prettify(s)
}

func (s GetAlertActionResponseBodyAlertActionsWebhookParam) GoString() string {
	return s.String()
}

func (s *GetAlertActionResponseBodyAlertActionsWebhookParam) GetContentType() *string {
	return s.ContentType
}

func (s *GetAlertActionResponseBodyAlertActionsWebhookParam) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlertActionResponseBodyAlertActionsWebhookParam) GetMethod() *string {
	return s.Method
}

func (s *GetAlertActionResponseBodyAlertActionsWebhookParam) GetUrl() *string {
	return s.Url
}

func (s *GetAlertActionResponseBodyAlertActionsWebhookParam) SetContentType(v string) *GetAlertActionResponseBodyAlertActionsWebhookParam {
	s.ContentType = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsWebhookParam) SetHeaders(v map[string]*string) *GetAlertActionResponseBodyAlertActionsWebhookParam {
	s.Headers = v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsWebhookParam) SetMethod(v string) *GetAlertActionResponseBodyAlertActionsWebhookParam {
	s.Method = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsWebhookParam) SetUrl(v string) *GetAlertActionResponseBodyAlertActionsWebhookParam {
	s.Url = &v
	return s
}

func (s *GetAlertActionResponseBodyAlertActionsWebhookParam) Validate() error {
	return dara.Validate(s)
}
