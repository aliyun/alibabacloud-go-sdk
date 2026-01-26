// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntegrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIntegration(v *UpdateIntegrationResponseBodyIntegration) *UpdateIntegrationResponseBody
	GetIntegration() *UpdateIntegrationResponseBodyIntegration
	SetRequestId(v string) *UpdateIntegrationResponseBody
	GetRequestId() *string
}

type UpdateIntegrationResponseBody struct {
	// The Information about the alert integration.
	Integration *UpdateIntegrationResponseBodyIntegration `json:"Integration,omitempty" xml:"Integration,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 34ED024E-9E31-434A-9E4E-D9D15C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIntegrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationResponseBody) GetIntegration() *UpdateIntegrationResponseBodyIntegration {
	return s.Integration
}

func (s *UpdateIntegrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIntegrationResponseBody) SetIntegration(v *UpdateIntegrationResponseBodyIntegration) *UpdateIntegrationResponseBody {
	s.Integration = v
	return s
}

func (s *UpdateIntegrationResponseBody) SetRequestId(v string) *UpdateIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIntegrationResponseBody) Validate() error {
	if s.Integration != nil {
		if err := s.Integration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIntegrationResponseBodyIntegration struct {
	// The endpoint of the alert integration.
	//
	// example:
	//
	// https://alerts.aliyuncs.com/api/v1/integrations/custom/ymQBN******
	ApiEndpoint *string `json:"ApiEndpoint,omitempty" xml:"ApiEndpoint,omitempty"`
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
	// LABEL.dimensions::LABEL.ruleId
	DuplicateKey *string `json:"DuplicateKey,omitempty" xml:"DuplicateKey,omitempty"`
	// The extended mapped fields of the alert source.
	ExtendedFieldRedefineRules []map[string]interface{} `json:"ExtendedFieldRedefineRules,omitempty" xml:"ExtendedFieldRedefineRules,omitempty" type:"Repeated"`
	// The predefined mapped fields of the alert source.
	FieldRedefineRules []map[string]interface{} `json:"FieldRedefineRules,omitempty" xml:"FieldRedefineRules,omitempty" type:"Repeated"`
	// The field for clearing alert events. The system queries alert events based on the field of alert clearing events and clears the alert events.
	//
	// > Only Log Service supports this parameter.
	//
	// example:
	//
	// $.status
	InitiativeRecoverField *string `json:"InitiativeRecoverField,omitempty" xml:"InitiativeRecoverField,omitempty"`
	// The value of the field for clearing alert events. The system queries alert events based on the field of alert clearing events and clears the alert events.
	//
	// > Only Log Service supports this parameter.
	//
	// example:
	//
	// ok
	InitiativeRecoverValue *string `json:"InitiativeRecoverValue,omitempty" xml:"InitiativeRecoverValue,omitempty"`
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
	// The service of the alert integration. Valid values:
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
	// The time when alert events are automatically cleared. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 300
	RecoverTime *int64 `json:"RecoverTime,omitempty" xml:"RecoverTime,omitempty"`
	// The authentication token of the alert integration.
	//
	// example:
	//
	// ymQBN******
	ShortToken *string `json:"ShortToken,omitempty" xml:"ShortToken,omitempty"`
	// The total number of alert events and the number of abnormal alert events in the last hour.
	Stat []*int64 `json:"Stat,omitempty" xml:"Stat,omitempty" type:"Repeated"`
	// Indicates whether the alert integration is enabled. Valid values:
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

func (s UpdateIntegrationResponseBodyIntegration) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntegrationResponseBodyIntegration) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationResponseBodyIntegration) GetApiEndpoint() *string {
	return s.ApiEndpoint
}

func (s *UpdateIntegrationResponseBodyIntegration) GetAutoRecover() *bool {
	return s.AutoRecover
}

func (s *UpdateIntegrationResponseBodyIntegration) GetDescription() *string {
	return s.Description
}

func (s *UpdateIntegrationResponseBodyIntegration) GetDuplicateKey() *string {
	return s.DuplicateKey
}

func (s *UpdateIntegrationResponseBodyIntegration) GetExtendedFieldRedefineRules() []map[string]interface{} {
	return s.ExtendedFieldRedefineRules
}

func (s *UpdateIntegrationResponseBodyIntegration) GetFieldRedefineRules() []map[string]interface{} {
	return s.FieldRedefineRules
}

func (s *UpdateIntegrationResponseBodyIntegration) GetInitiativeRecoverField() *string {
	return s.InitiativeRecoverField
}

func (s *UpdateIntegrationResponseBodyIntegration) GetInitiativeRecoverValue() *string {
	return s.InitiativeRecoverValue
}

func (s *UpdateIntegrationResponseBodyIntegration) GetIntegrationId() *int64 {
	return s.IntegrationId
}

func (s *UpdateIntegrationResponseBodyIntegration) GetIntegrationName() *string {
	return s.IntegrationName
}

func (s *UpdateIntegrationResponseBodyIntegration) GetIntegrationProductType() *string {
	return s.IntegrationProductType
}

func (s *UpdateIntegrationResponseBodyIntegration) GetLiveness() *string {
	return s.Liveness
}

func (s *UpdateIntegrationResponseBodyIntegration) GetRecoverTime() *int64 {
	return s.RecoverTime
}

func (s *UpdateIntegrationResponseBodyIntegration) GetShortToken() *string {
	return s.ShortToken
}

func (s *UpdateIntegrationResponseBodyIntegration) GetStat() []*int64 {
	return s.Stat
}

func (s *UpdateIntegrationResponseBodyIntegration) GetState() *bool {
	return s.State
}

func (s *UpdateIntegrationResponseBodyIntegration) SetApiEndpoint(v string) *UpdateIntegrationResponseBodyIntegration {
	s.ApiEndpoint = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetAutoRecover(v bool) *UpdateIntegrationResponseBodyIntegration {
	s.AutoRecover = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetDescription(v string) *UpdateIntegrationResponseBodyIntegration {
	s.Description = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetDuplicateKey(v string) *UpdateIntegrationResponseBodyIntegration {
	s.DuplicateKey = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetExtendedFieldRedefineRules(v []map[string]interface{}) *UpdateIntegrationResponseBodyIntegration {
	s.ExtendedFieldRedefineRules = v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetFieldRedefineRules(v []map[string]interface{}) *UpdateIntegrationResponseBodyIntegration {
	s.FieldRedefineRules = v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetInitiativeRecoverField(v string) *UpdateIntegrationResponseBodyIntegration {
	s.InitiativeRecoverField = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetInitiativeRecoverValue(v string) *UpdateIntegrationResponseBodyIntegration {
	s.InitiativeRecoverValue = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetIntegrationId(v int64) *UpdateIntegrationResponseBodyIntegration {
	s.IntegrationId = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetIntegrationName(v string) *UpdateIntegrationResponseBodyIntegration {
	s.IntegrationName = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetIntegrationProductType(v string) *UpdateIntegrationResponseBodyIntegration {
	s.IntegrationProductType = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetLiveness(v string) *UpdateIntegrationResponseBodyIntegration {
	s.Liveness = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetRecoverTime(v int64) *UpdateIntegrationResponseBodyIntegration {
	s.RecoverTime = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetShortToken(v string) *UpdateIntegrationResponseBodyIntegration {
	s.ShortToken = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetStat(v []*int64) *UpdateIntegrationResponseBodyIntegration {
	s.Stat = v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetState(v bool) *UpdateIntegrationResponseBodyIntegration {
	s.State = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) Validate() error {
	return dara.Validate(s)
}
