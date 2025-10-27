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
	// The whitelist rules.
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
	// The ID of the whitelist rule. If you do not specify this parameter, a whitelist rule is created.
	//
	// example:
	//
	// 1
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The name of the alert.
	//
	// 	- Set the value to **ALL**, which indicates all alert types.
	//
	// example:
	//
	// ALL
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The field that you want to use in the whitelist rule.
	//
	// example:
	//
	// fileMd5
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The value of the field that you want to use in the whitelist rule.
	//
	// example:
	//
	// b2cf9747ee49d8d9b105cf16e078cc16
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The logical operator that you want to use in the whitelist rule.
	//
	// 	- Set the value to strEqual, which indicates the equality operator (=).
	//
	// example:
	//
	// strEqual
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The feature to which this operation belongs.
	//
	// 	- Set the value to agentless, which indicates the agentless detection feature.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the assets on which you want the whitelist rule to take effect. Valid values:
	//
	// 	- ALL: all assets
	//
	// 	- SELECTION_KEY: selected assets
	//
	// example:
	//
	// ALL
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The assets on which you want the whitelist rule to take effect. Valid values:
	//
	// 	- ALL: all assets
	//
	// 	- Others: selected assets
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
