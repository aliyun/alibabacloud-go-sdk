// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertActionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertActionName(v string) *UpdateAlertActionRequest
	GetAlertActionName() *string
	SetEbParam(v *UpdateAlertActionRequestEbParam) *UpdateAlertActionRequest
	GetEbParam() *UpdateAlertActionRequestEbParam
	SetEssParam(v *UpdateAlertActionRequestEssParam) *UpdateAlertActionRequest
	GetEssParam() *UpdateAlertActionRequestEssParam
	SetFc3Param(v *UpdateAlertActionRequestFc3Param) *UpdateAlertActionRequest
	GetFc3Param() *UpdateAlertActionRequestFc3Param
	SetFcParam(v *UpdateAlertActionRequestFcParam) *UpdateAlertActionRequest
	GetFcParam() *UpdateAlertActionRequestFcParam
	SetMnsParam(v *UpdateAlertActionRequestMnsParam) *UpdateAlertActionRequest
	GetMnsParam() *UpdateAlertActionRequestMnsParam
	SetPagerDutyParam(v *UpdateAlertActionRequestPagerDutyParam) *UpdateAlertActionRequest
	GetPagerDutyParam() *UpdateAlertActionRequestPagerDutyParam
	SetSlsParam(v *UpdateAlertActionRequestSlsParam) *UpdateAlertActionRequest
	GetSlsParam() *UpdateAlertActionRequestSlsParam
	SetType(v string) *UpdateAlertActionRequest
	GetType() *string
	SetWebhookParam(v *UpdateAlertActionRequestWebhookParam) *UpdateAlertActionRequest
	GetWebhookParam() *UpdateAlertActionRequestWebhookParam
}

type UpdateAlertActionRequest struct {
	// example:
	//
	// testName
	AlertActionName *string                                 `json:"alertActionName,omitempty" xml:"alertActionName,omitempty"`
	EbParam         *UpdateAlertActionRequestEbParam        `json:"ebParam,omitempty" xml:"ebParam,omitempty" type:"Struct"`
	EssParam        *UpdateAlertActionRequestEssParam       `json:"essParam,omitempty" xml:"essParam,omitempty" type:"Struct"`
	Fc3Param        *UpdateAlertActionRequestFc3Param       `json:"fc3Param,omitempty" xml:"fc3Param,omitempty" type:"Struct"`
	FcParam         *UpdateAlertActionRequestFcParam        `json:"fcParam,omitempty" xml:"fcParam,omitempty" type:"Struct"`
	MnsParam        *UpdateAlertActionRequestMnsParam       `json:"mnsParam,omitempty" xml:"mnsParam,omitempty" type:"Struct"`
	PagerDutyParam  *UpdateAlertActionRequestPagerDutyParam `json:"pagerDutyParam,omitempty" xml:"pagerDutyParam,omitempty" type:"Struct"`
	SlsParam        *UpdateAlertActionRequestSlsParam       `json:"slsParam,omitempty" xml:"slsParam,omitempty" type:"Struct"`
	// example:
	//
	// FC
	Type         *string                               `json:"type,omitempty" xml:"type,omitempty"`
	WebhookParam *UpdateAlertActionRequestWebhookParam `json:"webhookParam,omitempty" xml:"webhookParam,omitempty" type:"Struct"`
}

func (s UpdateAlertActionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertActionRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertActionRequest) GetAlertActionName() *string {
	return s.AlertActionName
}

func (s *UpdateAlertActionRequest) GetEbParam() *UpdateAlertActionRequestEbParam {
	return s.EbParam
}

func (s *UpdateAlertActionRequest) GetEssParam() *UpdateAlertActionRequestEssParam {
	return s.EssParam
}

func (s *UpdateAlertActionRequest) GetFc3Param() *UpdateAlertActionRequestFc3Param {
	return s.Fc3Param
}

func (s *UpdateAlertActionRequest) GetFcParam() *UpdateAlertActionRequestFcParam {
	return s.FcParam
}

func (s *UpdateAlertActionRequest) GetMnsParam() *UpdateAlertActionRequestMnsParam {
	return s.MnsParam
}

func (s *UpdateAlertActionRequest) GetPagerDutyParam() *UpdateAlertActionRequestPagerDutyParam {
	return s.PagerDutyParam
}

func (s *UpdateAlertActionRequest) GetSlsParam() *UpdateAlertActionRequestSlsParam {
	return s.SlsParam
}

func (s *UpdateAlertActionRequest) GetType() *string {
	return s.Type
}

func (s *UpdateAlertActionRequest) GetWebhookParam() *UpdateAlertActionRequestWebhookParam {
	return s.WebhookParam
}

func (s *UpdateAlertActionRequest) SetAlertActionName(v string) *UpdateAlertActionRequest {
	s.AlertActionName = &v
	return s
}

func (s *UpdateAlertActionRequest) SetEbParam(v *UpdateAlertActionRequestEbParam) *UpdateAlertActionRequest {
	s.EbParam = v
	return s
}

func (s *UpdateAlertActionRequest) SetEssParam(v *UpdateAlertActionRequestEssParam) *UpdateAlertActionRequest {
	s.EssParam = v
	return s
}

func (s *UpdateAlertActionRequest) SetFc3Param(v *UpdateAlertActionRequestFc3Param) *UpdateAlertActionRequest {
	s.Fc3Param = v
	return s
}

func (s *UpdateAlertActionRequest) SetFcParam(v *UpdateAlertActionRequestFcParam) *UpdateAlertActionRequest {
	s.FcParam = v
	return s
}

func (s *UpdateAlertActionRequest) SetMnsParam(v *UpdateAlertActionRequestMnsParam) *UpdateAlertActionRequest {
	s.MnsParam = v
	return s
}

func (s *UpdateAlertActionRequest) SetPagerDutyParam(v *UpdateAlertActionRequestPagerDutyParam) *UpdateAlertActionRequest {
	s.PagerDutyParam = v
	return s
}

func (s *UpdateAlertActionRequest) SetSlsParam(v *UpdateAlertActionRequestSlsParam) *UpdateAlertActionRequest {
	s.SlsParam = v
	return s
}

func (s *UpdateAlertActionRequest) SetType(v string) *UpdateAlertActionRequest {
	s.Type = &v
	return s
}

func (s *UpdateAlertActionRequest) SetWebhookParam(v *UpdateAlertActionRequestWebhookParam) *UpdateAlertActionRequest {
	s.WebhookParam = v
	return s
}

func (s *UpdateAlertActionRequest) Validate() error {
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

type UpdateAlertActionRequestEbParam struct {
	// example:
	//
	// test
	EbSource *string `json:"ebSource,omitempty" xml:"ebSource,omitempty"`
	// example:
	//
	// test
	EventBusName *string `json:"eventBusName,omitempty" xml:"eventBusName,omitempty"`
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// test
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s UpdateAlertActionRequestEbParam) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertActionRequestEbParam) GoString() string {
	return s.String()
}

func (s *UpdateAlertActionRequestEbParam) GetEbSource() *string {
	return s.EbSource
}

func (s *UpdateAlertActionRequestEbParam) GetEventBusName() *string {
	return s.EventBusName
}

func (s *UpdateAlertActionRequestEbParam) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAlertActionRequestEbParam) GetSubject() *string {
	return s.Subject
}

func (s *UpdateAlertActionRequestEbParam) SetEbSource(v string) *UpdateAlertActionRequestEbParam {
	s.EbSource = &v
	return s
}

func (s *UpdateAlertActionRequestEbParam) SetEventBusName(v string) *UpdateAlertActionRequestEbParam {
	s.EventBusName = &v
	return s
}

func (s *UpdateAlertActionRequestEbParam) SetRegionId(v string) *UpdateAlertActionRequestEbParam {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertActionRequestEbParam) SetSubject(v string) *UpdateAlertActionRequestEbParam {
	s.Subject = &v
	return s
}

func (s *UpdateAlertActionRequestEbParam) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertActionRequestEssParam struct {
	// example:
	//
	// testId
	EssGroupId *string `json:"essGroupId,omitempty" xml:"essGroupId,omitempty"`
	// example:
	//
	// testId
	EssRuleId *string `json:"essRuleId,omitempty" xml:"essRuleId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s UpdateAlertActionRequestEssParam) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertActionRequestEssParam) GoString() string {
	return s.String()
}

func (s *UpdateAlertActionRequestEssParam) GetEssGroupId() *string {
	return s.EssGroupId
}

func (s *UpdateAlertActionRequestEssParam) GetEssRuleId() *string {
	return s.EssRuleId
}

func (s *UpdateAlertActionRequestEssParam) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAlertActionRequestEssParam) SetEssGroupId(v string) *UpdateAlertActionRequestEssParam {
	s.EssGroupId = &v
	return s
}

func (s *UpdateAlertActionRequestEssParam) SetEssRuleId(v string) *UpdateAlertActionRequestEssParam {
	s.EssRuleId = &v
	return s
}

func (s *UpdateAlertActionRequestEssParam) SetRegionId(v string) *UpdateAlertActionRequestEssParam {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertActionRequestEssParam) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertActionRequestFc3Param struct {
	// example:
	//
	// test
	Function *string `json:"function,omitempty" xml:"function,omitempty"`
	// example:
	//
	// test
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s UpdateAlertActionRequestFc3Param) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertActionRequestFc3Param) GoString() string {
	return s.String()
}

func (s *UpdateAlertActionRequestFc3Param) GetFunction() *string {
	return s.Function
}

func (s *UpdateAlertActionRequestFc3Param) GetQualifier() *string {
	return s.Qualifier
}

func (s *UpdateAlertActionRequestFc3Param) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAlertActionRequestFc3Param) SetFunction(v string) *UpdateAlertActionRequestFc3Param {
	s.Function = &v
	return s
}

func (s *UpdateAlertActionRequestFc3Param) SetQualifier(v string) *UpdateAlertActionRequestFc3Param {
	s.Qualifier = &v
	return s
}

func (s *UpdateAlertActionRequestFc3Param) SetRegionId(v string) *UpdateAlertActionRequestFc3Param {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertActionRequestFc3Param) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertActionRequestFcParam struct {
	// example:
	//
	// testName
	Function *string `json:"function,omitempty" xml:"function,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// testService
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
}

func (s UpdateAlertActionRequestFcParam) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertActionRequestFcParam) GoString() string {
	return s.String()
}

func (s *UpdateAlertActionRequestFcParam) GetFunction() *string {
	return s.Function
}

func (s *UpdateAlertActionRequestFcParam) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAlertActionRequestFcParam) GetService() *string {
	return s.Service
}

func (s *UpdateAlertActionRequestFcParam) SetFunction(v string) *UpdateAlertActionRequestFcParam {
	s.Function = &v
	return s
}

func (s *UpdateAlertActionRequestFcParam) SetRegionId(v string) *UpdateAlertActionRequestFcParam {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertActionRequestFcParam) SetService(v string) *UpdateAlertActionRequestFcParam {
	s.Service = &v
	return s
}

func (s *UpdateAlertActionRequestFcParam) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertActionRequestMnsParam struct {
	// example:
	//
	// queue
	MnsType *string `json:"mnsType,omitempty" xml:"mnsType,omitempty"`
	// example:
	//
	// testName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s UpdateAlertActionRequestMnsParam) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertActionRequestMnsParam) GoString() string {
	return s.String()
}

func (s *UpdateAlertActionRequestMnsParam) GetMnsType() *string {
	return s.MnsType
}

func (s *UpdateAlertActionRequestMnsParam) GetName() *string {
	return s.Name
}

func (s *UpdateAlertActionRequestMnsParam) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAlertActionRequestMnsParam) SetMnsType(v string) *UpdateAlertActionRequestMnsParam {
	s.MnsType = &v
	return s
}

func (s *UpdateAlertActionRequestMnsParam) SetName(v string) *UpdateAlertActionRequestMnsParam {
	s.Name = &v
	return s
}

func (s *UpdateAlertActionRequestMnsParam) SetRegionId(v string) *UpdateAlertActionRequestMnsParam {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertActionRequestMnsParam) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertActionRequestPagerDutyParam struct {
	// example:
	//
	// fwe4rwr543t543t43t43t43
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// https://events.pagerduty.com/v2/enqueue
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UpdateAlertActionRequestPagerDutyParam) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertActionRequestPagerDutyParam) GoString() string {
	return s.String()
}

func (s *UpdateAlertActionRequestPagerDutyParam) GetKey() *string {
	return s.Key
}

func (s *UpdateAlertActionRequestPagerDutyParam) GetUrl() *string {
	return s.Url
}

func (s *UpdateAlertActionRequestPagerDutyParam) SetKey(v string) *UpdateAlertActionRequestPagerDutyParam {
	s.Key = &v
	return s
}

func (s *UpdateAlertActionRequestPagerDutyParam) SetUrl(v string) *UpdateAlertActionRequestPagerDutyParam {
	s.Url = &v
	return s
}

func (s *UpdateAlertActionRequestPagerDutyParam) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertActionRequestSlsParam struct {
	// example:
	//
	// testLogstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// testProject
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s UpdateAlertActionRequestSlsParam) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertActionRequestSlsParam) GoString() string {
	return s.String()
}

func (s *UpdateAlertActionRequestSlsParam) GetLogstore() *string {
	return s.Logstore
}

func (s *UpdateAlertActionRequestSlsParam) GetProject() *string {
	return s.Project
}

func (s *UpdateAlertActionRequestSlsParam) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAlertActionRequestSlsParam) SetLogstore(v string) *UpdateAlertActionRequestSlsParam {
	s.Logstore = &v
	return s
}

func (s *UpdateAlertActionRequestSlsParam) SetProject(v string) *UpdateAlertActionRequestSlsParam {
	s.Project = &v
	return s
}

func (s *UpdateAlertActionRequestSlsParam) SetRegionId(v string) *UpdateAlertActionRequestSlsParam {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertActionRequestSlsParam) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertActionRequestWebhookParam struct {
	// example:
	//
	// JSON
	ContentType *string            `json:"contentType,omitempty" xml:"contentType,omitempty"`
	Headers     map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// example:
	//
	// POST
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// example:
	//
	// http://www.test.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UpdateAlertActionRequestWebhookParam) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertActionRequestWebhookParam) GoString() string {
	return s.String()
}

func (s *UpdateAlertActionRequestWebhookParam) GetContentType() *string {
	return s.ContentType
}

func (s *UpdateAlertActionRequestWebhookParam) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlertActionRequestWebhookParam) GetMethod() *string {
	return s.Method
}

func (s *UpdateAlertActionRequestWebhookParam) GetUrl() *string {
	return s.Url
}

func (s *UpdateAlertActionRequestWebhookParam) SetContentType(v string) *UpdateAlertActionRequestWebhookParam {
	s.ContentType = &v
	return s
}

func (s *UpdateAlertActionRequestWebhookParam) SetHeaders(v map[string]*string) *UpdateAlertActionRequestWebhookParam {
	s.Headers = v
	return s
}

func (s *UpdateAlertActionRequestWebhookParam) SetMethod(v string) *UpdateAlertActionRequestWebhookParam {
	s.Method = &v
	return s
}

func (s *UpdateAlertActionRequestWebhookParam) SetUrl(v string) *UpdateAlertActionRequestWebhookParam {
	s.Url = &v
	return s
}

func (s *UpdateAlertActionRequestWebhookParam) Validate() error {
	return dara.Validate(s)
}
