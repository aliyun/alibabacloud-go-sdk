// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateMaliciousFileWhitelistConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigList(v []*BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) *BatchUpdateMaliciousFileWhitelistConfigRequest
	GetConfigList() []*BatchUpdateMaliciousFileWhitelistConfigRequestConfigList
}

type BatchUpdateMaliciousFileWhitelistConfigRequest struct {
	// The list of whitelist rules for sensitive file alerts in agentless detection.
	ConfigList []*BatchUpdateMaliciousFileWhitelistConfigRequestConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
}

func (s BatchUpdateMaliciousFileWhitelistConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateMaliciousFileWhitelistConfigRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequest) GetConfigList() []*BatchUpdateMaliciousFileWhitelistConfigRequestConfigList {
	return s.ConfigList
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequest) SetConfigList(v []*BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) *BatchUpdateMaliciousFileWhitelistConfigRequest {
	s.ConfigList = v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequest) Validate() error {
	if s.ConfigList != nil {
		for _, item := range s.ConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchUpdateMaliciousFileWhitelistConfigRequestConfigList struct {
	// The ID of the whitelist rule for sensitive file alerts in agentless detection. This parameter is optional. If you do not specify this parameter, a new rule is created. You can call the [ListMaliciousFileWhitelistConfigs](~~ListMaliciousFileWhitelistConfigs~~) operation to obtain the ID of the whitelist rule for sensitive files.
	//
	// example:
	//
	// 1
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The alerting name. Valid values:
	//
	// - **ALL**: all Alarm Metrics.
	//
	// example:
	//
	// ALL
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The field used to determine the whitelist for sensitive file alerts.
	//
	// example:
	//
	// fileMd5
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The expected field value to be whitelisted.
	//
	// example:
	//
	// b2cf9747ee49d8d9b105cf16e078cc16
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The operator used to evaluate the rule. Valid values:
	//
	// - strEqual: equal to.
	//
	// example:
	//
	// strEqual
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The business source. Valid values:
	//
	// - agentless: agentless detection.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the target to which the rule applies. Valid values:
	//
	// - ALL: all assets.
	//
	// - SELECTION_KEY: assets selected by using the asset selection component.
	//
	// example:
	//
	// ALL
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The scope of the target to which the rule applies. Valid values:
	//
	// - ALL: all assets.
	//
	// - Other values: the key of the asset scope selected by using the asset selection component.
	//
	// example:
	//
	// ALL
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) GoString() string {
	return s.String()
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) GetEventName() *string {
	return s.EventName
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) GetField() *string {
	return s.Field
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) GetFieldValue() *string {
	return s.FieldValue
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) GetOperator() *string {
	return s.Operator
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) GetSource() *string {
	return s.Source
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) GetTargetType() *string {
	return s.TargetType
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) GetTargetValue() *string {
	return s.TargetValue
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) SetConfigId(v int64) *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList {
	s.ConfigId = &v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) SetEventName(v string) *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList {
	s.EventName = &v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) SetField(v string) *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList {
	s.Field = &v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) SetFieldValue(v string) *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList {
	s.FieldValue = &v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) SetOperator(v string) *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList {
	s.Operator = &v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) SetSource(v string) *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList {
	s.Source = &v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) SetTargetType(v string) *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList {
	s.TargetType = &v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) SetTargetValue(v string) *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList {
	s.TargetValue = &v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigRequestConfigList) Validate() error {
	return dara.Validate(s)
}
