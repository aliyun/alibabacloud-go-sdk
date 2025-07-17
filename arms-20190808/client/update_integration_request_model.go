// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntegrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRecover(v bool) *UpdateIntegrationRequest
	GetAutoRecover() *bool
	SetDescription(v string) *UpdateIntegrationRequest
	GetDescription() *string
	SetDuplicateKey(v string) *UpdateIntegrationRequest
	GetDuplicateKey() *string
	SetExtendedFieldRedefineRules(v string) *UpdateIntegrationRequest
	GetExtendedFieldRedefineRules() *string
	SetFieldRedefineRules(v string) *UpdateIntegrationRequest
	GetFieldRedefineRules() *string
	SetInitiativeRecoverField(v string) *UpdateIntegrationRequest
	GetInitiativeRecoverField() *string
	SetInitiativeRecoverValue(v string) *UpdateIntegrationRequest
	GetInitiativeRecoverValue() *string
	SetIntegrationId(v int64) *UpdateIntegrationRequest
	GetIntegrationId() *int64
	SetIntegrationName(v string) *UpdateIntegrationRequest
	GetIntegrationName() *string
	SetIntegrationProductType(v string) *UpdateIntegrationRequest
	GetIntegrationProductType() *string
	SetLiveness(v string) *UpdateIntegrationRequest
	GetLiveness() *string
	SetRecoverTime(v int64) *UpdateIntegrationRequest
	GetRecoverTime() *int64
	SetStat(v string) *UpdateIntegrationRequest
	GetStat() *string
	SetState(v bool) *UpdateIntegrationRequest
	GetState() *bool
}

type UpdateIntegrationRequest struct {
	// Specifies whether to automatically clear alert events. Valid values:
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
	// LABEL.dimensions::LABEL.ruleId
	DuplicateKey *string `json:"DuplicateKey,omitempty" xml:"DuplicateKey,omitempty"`
	// The extended mapped fields are mapped to the fields of ARMS alerts. For more information, see the description of the ExtendedFieldRedefineRules parameter.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "redefineType":"EXTRACT",
	//
	//         "matchExpression":null,
	//
	//         "fieldName":"dimensions",
	//
	//         "expression":null,
	//
	//         "mappingRuleList":[
	//
	//         ],
	//
	//         "name":"dimensions",
	//
	//         "integrationId":1234,
	//
	//         "jsonPath":"$.dimensions",
	//
	//         "id":10013,
	//
	//         "fieldType":"LABEL"
	//
	//     },
	//
	//     {
	//
	//         "redefineType":"EXTRACT",
	//
	//         "matchExpression":null,
	//
	//         "fieldName":"expression",
	//
	//         "expression":null,
	//
	//         "mappingRuleList":[
	//
	//         ],
	//
	//         "name":"expression",
	//
	//         "integrationId":1234,
	//
	//         "jsonPath":"$.expression",
	//
	//         "id":10014,
	//
	//         "fieldType":"LABEL"
	//
	//     }
	//
	// ]
	ExtendedFieldRedefineRules *string `json:"ExtendedFieldRedefineRules,omitempty" xml:"ExtendedFieldRedefineRules,omitempty"`
	// The predefined mapped fields are mapped to the fields of ARMS alerts. The predefined mapped fields were generated when the alert integration was created. For more information, see the description of the FieldRedefineRules parameter.
	//
	// example:
	//
	// [ { "redefineType":"EXTRACT", "matchExpression":null, "fieldName":"alertname", "expression":null, "mappingRuleList":[ ], "name":"Alert name", "integrationId":1234, "jsonPath":"$.alertName", "id":10001, "fieldType":"LABEL" }, { "redefineType":"MAP", "matchExpression":null, "fieldName":"severity", "expression":null, "mappingRuleList":[ { "mappingValue":"critical", "mappingName":"P1", "mappingType":"MAP", "originValue":"CRITICAL" }, { "mappingValue":"error", "mappingName":"P2", "mappingType":"MAP", "originValue":"WARN" }, { "mappingValue":"warning", "mappingName":"P3", "mappingType":"MAP", "originValue":"INFO" } ], "name":"Alert level", "integrationId":1234, "jsonPath":"$.triggerLevel", "id":10002, "fieldType":"LABEL" }, { "redefineType":"EXTRACT", "matchExpression":null, "fieldName":"message", "expression":"{{$labels.namespace}} / {{$labels.dimensions}} Alert content {{ $labels.alertname }}, Current value {{$value}}.", "mappingRuleList":[ ], "name":"Alert description", "integrationId":1234, "jsonPath":null, "id":10003, "fieldType":"ANNOTATION" }, { "redefineType":"EXTRACT", "matchExpression":null, "fieldName":"value", "expression":null, "mappingRuleList":[ ], "name":"Alert sample value", "integrationId":1234, "jsonPath":"$.curValue", "id":10004, "fieldType":"ANNOTATION" }, { "redefineType":"EXTRACT", "matchExpression":null, "fieldName":"source", "expression":null, "mappingRuleList":[ ], "name":"Source", "integrationId":1234, "jsonPath":null, "id":10007, "fieldType":"LABEL" }, { "redefineType":"ADD", "matchExpression":null, "fieldName":"generatorUrl", "expression":"https://cloudmonitor.console.aliyun.com/index.htm#/alarmInfo/name={{$labels.ruleId}}\\&searchValue=\\&searchType=name\\&searchProduct=/history/all/searchKey:{{$labels.ruleId}},startTime:{{sub $startsAt 300000}},endTime:{{$endsAt}}", "mappingRuleList":[ ], "name":"Event URL", "integrationId":1234, "jsonPath":"https://cloudmonitor.console.aliyun.com/index.htm#/alarmInfo/name={{$labels.ruleId}}\\&searchValue=\\&searchType=name\\&searchProduct=/history/all/searchKey:{{$labels.ruleId}},startTime:{{sub $startsAt 300000}},endTime:{{$endsAt}}", "id":10012, "fieldType":"GENERATE_URL" } ]
	FieldRedefineRules *string `json:"FieldRedefineRules,omitempty" xml:"FieldRedefineRules,omitempty"`
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
	// The ID of the alert integration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	IntegrationId *int64 `json:"IntegrationId,omitempty" xml:"IntegrationId,omitempty"`
	// The name of the alert integration.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudMonitor integration
	IntegrationName *string `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	// The service of the alert integration. Valid values:
	//
	// 	- CLOUD_MONITOR: CloudMonitor
	//
	// 	- LOG_SERVICE: Log Service
	//
	// This parameter is required.
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
	// The period of time within which alert events are automatically cleared. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 300
	RecoverTime *int64 `json:"RecoverTime,omitempty" xml:"RecoverTime,omitempty"`
	// The total number of alert events and the number of abnormal alert events in the last hour.
	//
	// example:
	//
	// [0,0]
	Stat *string `json:"Stat,omitempty" xml:"Stat,omitempty"`
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

func (s UpdateIntegrationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntegrationRequest) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationRequest) GetAutoRecover() *bool {
	return s.AutoRecover
}

func (s *UpdateIntegrationRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateIntegrationRequest) GetDuplicateKey() *string {
	return s.DuplicateKey
}

func (s *UpdateIntegrationRequest) GetExtendedFieldRedefineRules() *string {
	return s.ExtendedFieldRedefineRules
}

func (s *UpdateIntegrationRequest) GetFieldRedefineRules() *string {
	return s.FieldRedefineRules
}

func (s *UpdateIntegrationRequest) GetInitiativeRecoverField() *string {
	return s.InitiativeRecoverField
}

func (s *UpdateIntegrationRequest) GetInitiativeRecoverValue() *string {
	return s.InitiativeRecoverValue
}

func (s *UpdateIntegrationRequest) GetIntegrationId() *int64 {
	return s.IntegrationId
}

func (s *UpdateIntegrationRequest) GetIntegrationName() *string {
	return s.IntegrationName
}

func (s *UpdateIntegrationRequest) GetIntegrationProductType() *string {
	return s.IntegrationProductType
}

func (s *UpdateIntegrationRequest) GetLiveness() *string {
	return s.Liveness
}

func (s *UpdateIntegrationRequest) GetRecoverTime() *int64 {
	return s.RecoverTime
}

func (s *UpdateIntegrationRequest) GetStat() *string {
	return s.Stat
}

func (s *UpdateIntegrationRequest) GetState() *bool {
	return s.State
}

func (s *UpdateIntegrationRequest) SetAutoRecover(v bool) *UpdateIntegrationRequest {
	s.AutoRecover = &v
	return s
}

func (s *UpdateIntegrationRequest) SetDescription(v string) *UpdateIntegrationRequest {
	s.Description = &v
	return s
}

func (s *UpdateIntegrationRequest) SetDuplicateKey(v string) *UpdateIntegrationRequest {
	s.DuplicateKey = &v
	return s
}

func (s *UpdateIntegrationRequest) SetExtendedFieldRedefineRules(v string) *UpdateIntegrationRequest {
	s.ExtendedFieldRedefineRules = &v
	return s
}

func (s *UpdateIntegrationRequest) SetFieldRedefineRules(v string) *UpdateIntegrationRequest {
	s.FieldRedefineRules = &v
	return s
}

func (s *UpdateIntegrationRequest) SetInitiativeRecoverField(v string) *UpdateIntegrationRequest {
	s.InitiativeRecoverField = &v
	return s
}

func (s *UpdateIntegrationRequest) SetInitiativeRecoverValue(v string) *UpdateIntegrationRequest {
	s.InitiativeRecoverValue = &v
	return s
}

func (s *UpdateIntegrationRequest) SetIntegrationId(v int64) *UpdateIntegrationRequest {
	s.IntegrationId = &v
	return s
}

func (s *UpdateIntegrationRequest) SetIntegrationName(v string) *UpdateIntegrationRequest {
	s.IntegrationName = &v
	return s
}

func (s *UpdateIntegrationRequest) SetIntegrationProductType(v string) *UpdateIntegrationRequest {
	s.IntegrationProductType = &v
	return s
}

func (s *UpdateIntegrationRequest) SetLiveness(v string) *UpdateIntegrationRequest {
	s.Liveness = &v
	return s
}

func (s *UpdateIntegrationRequest) SetRecoverTime(v int64) *UpdateIntegrationRequest {
	s.RecoverTime = &v
	return s
}

func (s *UpdateIntegrationRequest) SetStat(v string) *UpdateIntegrationRequest {
	s.Stat = &v
	return s
}

func (s *UpdateIntegrationRequest) SetState(v bool) *UpdateIntegrationRequest {
	s.State = &v
	return s
}

func (s *UpdateIntegrationRequest) Validate() error {
	return dara.Validate(s)
}
