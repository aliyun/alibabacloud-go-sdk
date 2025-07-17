// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *ListIntegrationResponseBodyPageInfo) *ListIntegrationResponseBody
	GetPageInfo() *ListIntegrationResponseBodyPageInfo
	SetRequestId(v string) *ListIntegrationResponseBody
	GetRequestId() *string
}

type ListIntegrationResponseBody struct {
	// The pagination information.
	PageInfo *ListIntegrationResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 34ED024E-9E31-434A-9E4E-D9D15C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIntegrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntegrationResponseBody) GetPageInfo() *ListIntegrationResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListIntegrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntegrationResponseBody) SetPageInfo(v *ListIntegrationResponseBodyPageInfo) *ListIntegrationResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListIntegrationResponseBody) SetRequestId(v string) *ListIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntegrationResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationResponseBodyPageInfo struct {
	// The information about each alert integration.
	Integrations []*ListIntegrationResponseBodyPageInfoIntegrations `json:"Integrations,omitempty" xml:"Integrations,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of alert integrations returned per page.
	//
	// example:
	//
	// 10
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of alert integrations.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListIntegrationResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListIntegrationResponseBodyPageInfo) GetIntegrations() []*ListIntegrationResponseBodyPageInfoIntegrations {
	return s.Integrations
}

func (s *ListIntegrationResponseBodyPageInfo) GetPage() *int64 {
	return s.Page
}

func (s *ListIntegrationResponseBodyPageInfo) GetSize() *int64 {
	return s.Size
}

func (s *ListIntegrationResponseBodyPageInfo) GetTotal() *int64 {
	return s.Total
}

func (s *ListIntegrationResponseBodyPageInfo) SetIntegrations(v []*ListIntegrationResponseBodyPageInfoIntegrations) *ListIntegrationResponseBodyPageInfo {
	s.Integrations = v
	return s
}

func (s *ListIntegrationResponseBodyPageInfo) SetPage(v int64) *ListIntegrationResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfo) SetSize(v int64) *ListIntegrationResponseBodyPageInfo {
	s.Size = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfo) SetTotal(v int64) *ListIntegrationResponseBodyPageInfo {
	s.Total = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationResponseBodyPageInfoIntegrations struct {
	// The endpoint of the alert integration.
	//
	// example:
	//
	// https://alerts.aliyuncs.com/api/v1/integrations/custom/ymQBN******
	ApiEndpoint *string `json:"ApiEndpoint,omitempty" xml:"ApiEndpoint,omitempty"`
	// The time when the alert integration was created.
	//
	// example:
	//
	// 2022-06-18
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information about the alert events.
	IntegrationDetail *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail `json:"IntegrationDetail,omitempty" xml:"IntegrationDetail,omitempty" type:"Struct"`
	// The ID of the alert integration.
	//
	// example:
	//
	// 1234
	IntegrationId *int64 `json:"IntegrationId,omitempty" xml:"IntegrationId,omitempty"`
	// The name of the alert integration.
	//
	// example:
	//
	// CloudMonitor integration
	IntegrationName *string `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	// The type of the alert integration. Valid values:
	//
	// 	- CLOUD_MONITOR: CloudMonitor
	//
	// 	- LOG_SERVICE: Log Service
	//
	// example:
	//
	// CLOUD_MONITOR
	IntegrationProductType *string `json:"IntegrationProductType,omitempty" xml:"IntegrationProductType,omitempty"`
	// The activity of the alert integration
	//
	// example:
	//
	// ready
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// The authentication token of the alert integration.
	//
	// example:
	//
	// ymQBN******
	ShortToken *string `json:"ShortToken,omitempty" xml:"ShortToken,omitempty"`
	// Indicates whether the alert integration was enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	State *bool `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListIntegrationResponseBodyPageInfoIntegrations) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationResponseBodyPageInfoIntegrations) GoString() string {
	return s.String()
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) GetApiEndpoint() *string {
	return s.ApiEndpoint
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) GetIntegrationDetail() *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail {
	return s.IntegrationDetail
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) GetIntegrationId() *int64 {
	return s.IntegrationId
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) GetIntegrationName() *string {
	return s.IntegrationName
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) GetIntegrationProductType() *string {
	return s.IntegrationProductType
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) GetLiveness() *string {
	return s.Liveness
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) GetShortToken() *string {
	return s.ShortToken
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) GetState() *bool {
	return s.State
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) SetApiEndpoint(v string) *ListIntegrationResponseBodyPageInfoIntegrations {
	s.ApiEndpoint = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) SetCreateTime(v string) *ListIntegrationResponseBodyPageInfoIntegrations {
	s.CreateTime = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) SetIntegrationDetail(v *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) *ListIntegrationResponseBodyPageInfoIntegrations {
	s.IntegrationDetail = v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) SetIntegrationId(v int64) *ListIntegrationResponseBodyPageInfoIntegrations {
	s.IntegrationId = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) SetIntegrationName(v string) *ListIntegrationResponseBodyPageInfoIntegrations {
	s.IntegrationName = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) SetIntegrationProductType(v string) *ListIntegrationResponseBodyPageInfoIntegrations {
	s.IntegrationProductType = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) SetLiveness(v string) *ListIntegrationResponseBodyPageInfoIntegrations {
	s.Liveness = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) SetShortToken(v string) *ListIntegrationResponseBodyPageInfoIntegrations {
	s.ShortToken = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) SetState(v bool) *ListIntegrationResponseBodyPageInfoIntegrations {
	s.State = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrations) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail struct {
	// Indicates whether alert events are automatically cleared. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	AutoRecover *bool `json:"AutoRecover,omitempty" xml:"AutoRecover,omitempty"`
	// The description of the alert integration.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The fields whose values are deduplicated.
	//
	// example:
	//
	// LABEL.alertname::LABEL.severity
	DuplicateKey *string `json:"DuplicateKey,omitempty" xml:"DuplicateKey,omitempty"`
	// The extended mapped fields of the alert source.
	ExtendedFieldRedefineRules []map[string]interface{} `json:"ExtendedFieldRedefineRules,omitempty" xml:"ExtendedFieldRedefineRules,omitempty" type:"Repeated"`
	// The predefined mapped fields of the alert source.
	FieldRedefineRules []map[string]interface{} `json:"FieldRedefineRules,omitempty" xml:"FieldRedefineRules,omitempty" type:"Repeated"`
	// The field for clearing alert events. The system queries alert events based on the field of alert clearing events and clears the alert events.
	//
	// > Only the Log Service alert integration supports the parameter.
	//
	// example:
	//
	// $.status
	InitiativeRecoverField *string `json:"InitiativeRecoverField,omitempty" xml:"InitiativeRecoverField,omitempty"`
	// The value of the field for clearing alert events. The system queries alert events based on the field of alert clearing events and clears the alert events.
	//
	// > Only the Log Service alert integration supports the parameter.
	//
	// example:
	//
	// ok
	InitiativeRecoverValue *string `json:"InitiativeRecoverValue,omitempty" xml:"InitiativeRecoverValue,omitempty"`
	// The time when alert events are automatically cleared. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 300
	RecoverTime *int64 `json:"RecoverTime,omitempty" xml:"RecoverTime,omitempty"`
	// The total number of alert events and the number of abnormal alert events in the last hour.
	Stat []*int64 `json:"Stat,omitempty" xml:"Stat,omitempty" type:"Repeated"`
}

func (s ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) GoString() string {
	return s.String()
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) GetAutoRecover() *bool {
	return s.AutoRecover
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) GetDescription() *string {
	return s.Description
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) GetDuplicateKey() *string {
	return s.DuplicateKey
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) GetExtendedFieldRedefineRules() []map[string]interface{} {
	return s.ExtendedFieldRedefineRules
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) GetFieldRedefineRules() []map[string]interface{} {
	return s.FieldRedefineRules
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) GetInitiativeRecoverField() *string {
	return s.InitiativeRecoverField
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) GetInitiativeRecoverValue() *string {
	return s.InitiativeRecoverValue
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) GetRecoverTime() *int64 {
	return s.RecoverTime
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) GetStat() []*int64 {
	return s.Stat
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) SetAutoRecover(v bool) *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail {
	s.AutoRecover = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) SetDescription(v string) *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail {
	s.Description = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) SetDuplicateKey(v string) *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail {
	s.DuplicateKey = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) SetExtendedFieldRedefineRules(v []map[string]interface{}) *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail {
	s.ExtendedFieldRedefineRules = v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) SetFieldRedefineRules(v []map[string]interface{}) *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail {
	s.FieldRedefineRules = v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) SetInitiativeRecoverField(v string) *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail {
	s.InitiativeRecoverField = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) SetInitiativeRecoverValue(v string) *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail {
	s.InitiativeRecoverValue = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) SetRecoverTime(v int64) *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail {
	s.RecoverTime = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) SetStat(v []*int64) *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail {
	s.Stat = v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIntegrationsIntegrationDetail) Validate() error {
	return dara.Validate(s)
}
