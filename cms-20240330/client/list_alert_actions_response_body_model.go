// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertActionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertActions(v []*ListAlertActionsResponseBodyAlertActions) *ListAlertActionsResponseBody
	GetAlertActions() []*ListAlertActionsResponseBodyAlertActions
	SetPageNumber(v int64) *ListAlertActionsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAlertActionsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListAlertActionsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListAlertActionsResponseBody
	GetTotal() *int32
}

type ListAlertActionsResponseBody struct {
	AlertActions []*ListAlertActionsResponseBodyAlertActions `json:"alertActions,omitempty" xml:"alertActions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 9
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAlertActionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertActionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBody) GetAlertActions() []*ListAlertActionsResponseBodyAlertActions {
	return s.AlertActions
}

func (s *ListAlertActionsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAlertActionsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlertActionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertActionsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListAlertActionsResponseBody) SetAlertActions(v []*ListAlertActionsResponseBodyAlertActions) *ListAlertActionsResponseBody {
	s.AlertActions = v
	return s
}

func (s *ListAlertActionsResponseBody) SetPageNumber(v int64) *ListAlertActionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAlertActionsResponseBody) SetPageSize(v int64) *ListAlertActionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAlertActionsResponseBody) SetRequestId(v string) *ListAlertActionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertActionsResponseBody) SetTotal(v int32) *ListAlertActionsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAlertActionsResponseBody) Validate() error {
	if s.AlertActions != nil {
		for _, item := range s.AlertActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAlertActionsResponseBodyAlertActions struct {
	// example:
	//
	// test
	AlertActionId *string `json:"alertActionId,omitempty" xml:"alertActionId,omitempty"`
	// example:
	//
	// testName
	AlertActionName *string                                                 `json:"alertActionName,omitempty" xml:"alertActionName,omitempty"`
	EbParam         *ListAlertActionsResponseBodyAlertActionsEbParam        `json:"ebParam,omitempty" xml:"ebParam,omitempty" type:"Struct"`
	EssParam        *ListAlertActionsResponseBodyAlertActionsEssParam       `json:"essParam,omitempty" xml:"essParam,omitempty" type:"Struct"`
	Fc3Param        *ListAlertActionsResponseBodyAlertActionsFc3Param       `json:"fc3Param,omitempty" xml:"fc3Param,omitempty" type:"Struct"`
	FcParam         *ListAlertActionsResponseBodyAlertActionsFcParam        `json:"fcParam,omitempty" xml:"fcParam,omitempty" type:"Struct"`
	MnsParam        *ListAlertActionsResponseBodyAlertActionsMnsParam       `json:"mnsParam,omitempty" xml:"mnsParam,omitempty" type:"Struct"`
	PagerDutyParam  *ListAlertActionsResponseBodyAlertActionsPagerDutyParam `json:"pagerDutyParam,omitempty" xml:"pagerDutyParam,omitempty" type:"Struct"`
	SlsParam        *ListAlertActionsResponseBodyAlertActionsSlsParam       `json:"slsParam,omitempty" xml:"slsParam,omitempty" type:"Struct"`
	// example:
	//
	// FC
	Type         *string                                               `json:"type,omitempty" xml:"type,omitempty"`
	WebhookParam *ListAlertActionsResponseBodyAlertActionsWebhookParam `json:"webhookParam,omitempty" xml:"webhookParam,omitempty" type:"Struct"`
}

func (s ListAlertActionsResponseBodyAlertActions) String() string {
	return dara.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActions) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActions) GetAlertActionId() *string {
	return s.AlertActionId
}

func (s *ListAlertActionsResponseBodyAlertActions) GetAlertActionName() *string {
	return s.AlertActionName
}

func (s *ListAlertActionsResponseBodyAlertActions) GetEbParam() *ListAlertActionsResponseBodyAlertActionsEbParam {
	return s.EbParam
}

func (s *ListAlertActionsResponseBodyAlertActions) GetEssParam() *ListAlertActionsResponseBodyAlertActionsEssParam {
	return s.EssParam
}

func (s *ListAlertActionsResponseBodyAlertActions) GetFc3Param() *ListAlertActionsResponseBodyAlertActionsFc3Param {
	return s.Fc3Param
}

func (s *ListAlertActionsResponseBodyAlertActions) GetFcParam() *ListAlertActionsResponseBodyAlertActionsFcParam {
	return s.FcParam
}

func (s *ListAlertActionsResponseBodyAlertActions) GetMnsParam() *ListAlertActionsResponseBodyAlertActionsMnsParam {
	return s.MnsParam
}

func (s *ListAlertActionsResponseBodyAlertActions) GetPagerDutyParam() *ListAlertActionsResponseBodyAlertActionsPagerDutyParam {
	return s.PagerDutyParam
}

func (s *ListAlertActionsResponseBodyAlertActions) GetSlsParam() *ListAlertActionsResponseBodyAlertActionsSlsParam {
	return s.SlsParam
}

func (s *ListAlertActionsResponseBodyAlertActions) GetType() *string {
	return s.Type
}

func (s *ListAlertActionsResponseBodyAlertActions) GetWebhookParam() *ListAlertActionsResponseBodyAlertActionsWebhookParam {
	return s.WebhookParam
}

func (s *ListAlertActionsResponseBodyAlertActions) SetAlertActionId(v string) *ListAlertActionsResponseBodyAlertActions {
	s.AlertActionId = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetAlertActionName(v string) *ListAlertActionsResponseBodyAlertActions {
	s.AlertActionName = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetEbParam(v *ListAlertActionsResponseBodyAlertActionsEbParam) *ListAlertActionsResponseBodyAlertActions {
	s.EbParam = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetEssParam(v *ListAlertActionsResponseBodyAlertActionsEssParam) *ListAlertActionsResponseBodyAlertActions {
	s.EssParam = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetFc3Param(v *ListAlertActionsResponseBodyAlertActionsFc3Param) *ListAlertActionsResponseBodyAlertActions {
	s.Fc3Param = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetFcParam(v *ListAlertActionsResponseBodyAlertActionsFcParam) *ListAlertActionsResponseBodyAlertActions {
	s.FcParam = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetMnsParam(v *ListAlertActionsResponseBodyAlertActionsMnsParam) *ListAlertActionsResponseBodyAlertActions {
	s.MnsParam = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetPagerDutyParam(v *ListAlertActionsResponseBodyAlertActionsPagerDutyParam) *ListAlertActionsResponseBodyAlertActions {
	s.PagerDutyParam = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetSlsParam(v *ListAlertActionsResponseBodyAlertActionsSlsParam) *ListAlertActionsResponseBodyAlertActions {
	s.SlsParam = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetType(v string) *ListAlertActionsResponseBodyAlertActions {
	s.Type = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetWebhookParam(v *ListAlertActionsResponseBodyAlertActionsWebhookParam) *ListAlertActionsResponseBodyAlertActions {
	s.WebhookParam = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) Validate() error {
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

type ListAlertActionsResponseBodyAlertActionsEbParam struct {
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

func (s ListAlertActionsResponseBodyAlertActionsEbParam) String() string {
	return dara.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActionsEbParam) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActionsEbParam) GetEbSource() *string {
	return s.EbSource
}

func (s *ListAlertActionsResponseBodyAlertActionsEbParam) GetEventBusName() *string {
	return s.EventBusName
}

func (s *ListAlertActionsResponseBodyAlertActionsEbParam) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAlertActionsResponseBodyAlertActionsEbParam) GetSubject() *string {
	return s.Subject
}

func (s *ListAlertActionsResponseBodyAlertActionsEbParam) SetEbSource(v string) *ListAlertActionsResponseBodyAlertActionsEbParam {
	s.EbSource = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsEbParam) SetEventBusName(v string) *ListAlertActionsResponseBodyAlertActionsEbParam {
	s.EventBusName = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsEbParam) SetRegionId(v string) *ListAlertActionsResponseBodyAlertActionsEbParam {
	s.RegionId = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsEbParam) SetSubject(v string) *ListAlertActionsResponseBodyAlertActionsEbParam {
	s.Subject = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsEbParam) Validate() error {
	return dara.Validate(s)
}

type ListAlertActionsResponseBodyAlertActionsEssParam struct {
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

func (s ListAlertActionsResponseBodyAlertActionsEssParam) String() string {
	return dara.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActionsEssParam) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActionsEssParam) GetEssGroupId() *string {
	return s.EssGroupId
}

func (s *ListAlertActionsResponseBodyAlertActionsEssParam) GetEssRuleId() *string {
	return s.EssRuleId
}

func (s *ListAlertActionsResponseBodyAlertActionsEssParam) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAlertActionsResponseBodyAlertActionsEssParam) SetEssGroupId(v string) *ListAlertActionsResponseBodyAlertActionsEssParam {
	s.EssGroupId = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsEssParam) SetEssRuleId(v string) *ListAlertActionsResponseBodyAlertActionsEssParam {
	s.EssRuleId = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsEssParam) SetRegionId(v string) *ListAlertActionsResponseBodyAlertActionsEssParam {
	s.RegionId = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsEssParam) Validate() error {
	return dara.Validate(s)
}

type ListAlertActionsResponseBodyAlertActionsFc3Param struct {
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
	// cn-qingdao
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListAlertActionsResponseBodyAlertActionsFc3Param) String() string {
	return dara.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActionsFc3Param) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActionsFc3Param) GetFunction() *string {
	return s.Function
}

func (s *ListAlertActionsResponseBodyAlertActionsFc3Param) GetQualifier() *string {
	return s.Qualifier
}

func (s *ListAlertActionsResponseBodyAlertActionsFc3Param) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAlertActionsResponseBodyAlertActionsFc3Param) SetFunction(v string) *ListAlertActionsResponseBodyAlertActionsFc3Param {
	s.Function = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsFc3Param) SetQualifier(v string) *ListAlertActionsResponseBodyAlertActionsFc3Param {
	s.Qualifier = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsFc3Param) SetRegionId(v string) *ListAlertActionsResponseBodyAlertActionsFc3Param {
	s.RegionId = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsFc3Param) Validate() error {
	return dara.Validate(s)
}

type ListAlertActionsResponseBodyAlertActionsFcParam struct {
	// example:
	//
	// test
	Function *string `json:"function,omitempty" xml:"function,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// test
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
}

func (s ListAlertActionsResponseBodyAlertActionsFcParam) String() string {
	return dara.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActionsFcParam) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActionsFcParam) GetFunction() *string {
	return s.Function
}

func (s *ListAlertActionsResponseBodyAlertActionsFcParam) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAlertActionsResponseBodyAlertActionsFcParam) GetService() *string {
	return s.Service
}

func (s *ListAlertActionsResponseBodyAlertActionsFcParam) SetFunction(v string) *ListAlertActionsResponseBodyAlertActionsFcParam {
	s.Function = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsFcParam) SetRegionId(v string) *ListAlertActionsResponseBodyAlertActionsFcParam {
	s.RegionId = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsFcParam) SetService(v string) *ListAlertActionsResponseBodyAlertActionsFcParam {
	s.Service = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsFcParam) Validate() error {
	return dara.Validate(s)
}

type ListAlertActionsResponseBodyAlertActionsMnsParam struct {
	// example:
	//
	// queue
	MnsType *string `json:"mnsType,omitempty" xml:"mnsType,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListAlertActionsResponseBodyAlertActionsMnsParam) String() string {
	return dara.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActionsMnsParam) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActionsMnsParam) GetMnsType() *string {
	return s.MnsType
}

func (s *ListAlertActionsResponseBodyAlertActionsMnsParam) GetName() *string {
	return s.Name
}

func (s *ListAlertActionsResponseBodyAlertActionsMnsParam) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAlertActionsResponseBodyAlertActionsMnsParam) SetMnsType(v string) *ListAlertActionsResponseBodyAlertActionsMnsParam {
	s.MnsType = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsMnsParam) SetName(v string) *ListAlertActionsResponseBodyAlertActionsMnsParam {
	s.Name = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsMnsParam) SetRegionId(v string) *ListAlertActionsResponseBodyAlertActionsMnsParam {
	s.RegionId = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsMnsParam) Validate() error {
	return dara.Validate(s)
}

type ListAlertActionsResponseBodyAlertActionsPagerDutyParam struct {
	// example:
	//
	// fsfer4543t5t65g4t4
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// https://events.pagerduty.com/v2/enqueue
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListAlertActionsResponseBodyAlertActionsPagerDutyParam) String() string {
	return dara.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActionsPagerDutyParam) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActionsPagerDutyParam) GetKey() *string {
	return s.Key
}

func (s *ListAlertActionsResponseBodyAlertActionsPagerDutyParam) GetUrl() *string {
	return s.Url
}

func (s *ListAlertActionsResponseBodyAlertActionsPagerDutyParam) SetKey(v string) *ListAlertActionsResponseBodyAlertActionsPagerDutyParam {
	s.Key = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsPagerDutyParam) SetUrl(v string) *ListAlertActionsResponseBodyAlertActionsPagerDutyParam {
	s.Url = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsPagerDutyParam) Validate() error {
	return dara.Validate(s)
}

type ListAlertActionsResponseBodyAlertActionsSlsParam struct {
	// example:
	//
	// test
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// test
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListAlertActionsResponseBodyAlertActionsSlsParam) String() string {
	return dara.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActionsSlsParam) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActionsSlsParam) GetLogstore() *string {
	return s.Logstore
}

func (s *ListAlertActionsResponseBodyAlertActionsSlsParam) GetProject() *string {
	return s.Project
}

func (s *ListAlertActionsResponseBodyAlertActionsSlsParam) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAlertActionsResponseBodyAlertActionsSlsParam) SetLogstore(v string) *ListAlertActionsResponseBodyAlertActionsSlsParam {
	s.Logstore = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsSlsParam) SetProject(v string) *ListAlertActionsResponseBodyAlertActionsSlsParam {
	s.Project = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsSlsParam) SetRegionId(v string) *ListAlertActionsResponseBodyAlertActionsSlsParam {
	s.RegionId = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsSlsParam) Validate() error {
	return dara.Validate(s)
}

type ListAlertActionsResponseBodyAlertActionsWebhookParam struct {
	// example:
	//
	// JSON
	ContentType *string            `json:"contentType,omitempty" xml:"contentType,omitempty"`
	Headers     map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// example:
	//
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// example:
	//
	// http://www.test.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListAlertActionsResponseBodyAlertActionsWebhookParam) String() string {
	return dara.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActionsWebhookParam) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActionsWebhookParam) GetContentType() *string {
	return s.ContentType
}

func (s *ListAlertActionsResponseBodyAlertActionsWebhookParam) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertActionsResponseBodyAlertActionsWebhookParam) GetMethod() *string {
	return s.Method
}

func (s *ListAlertActionsResponseBodyAlertActionsWebhookParam) GetUrl() *string {
	return s.Url
}

func (s *ListAlertActionsResponseBodyAlertActionsWebhookParam) SetContentType(v string) *ListAlertActionsResponseBodyAlertActionsWebhookParam {
	s.ContentType = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsWebhookParam) SetHeaders(v map[string]*string) *ListAlertActionsResponseBodyAlertActionsWebhookParam {
	s.Headers = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsWebhookParam) SetMethod(v string) *ListAlertActionsResponseBodyAlertActionsWebhookParam {
	s.Method = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsWebhookParam) SetUrl(v string) *ListAlertActionsResponseBodyAlertActionsWebhookParam {
	s.Url = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsWebhookParam) Validate() error {
	return dara.Validate(s)
}
