// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageEventOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v string) *AddImageEventOperationRequest
	GetConditions() *string
	SetEventKey(v string) *AddImageEventOperationRequest
	GetEventKey() *string
	SetEventName(v string) *AddImageEventOperationRequest
	GetEventName() *string
	SetEventType(v string) *AddImageEventOperationRequest
	GetEventType() *string
	SetNote(v string) *AddImageEventOperationRequest
	GetNote() *string
	SetOperationCode(v string) *AddImageEventOperationRequest
	GetOperationCode() *string
	SetScenarios(v string) *AddImageEventOperationRequest
	GetScenarios() *string
	SetSource(v string) *AddImageEventOperationRequest
	GetSource() *string
}

type AddImageEventOperationRequest struct {
	// The rule conditions. The value is in the JSON format. Valid values of keys:
	//
	// 	- **condition**: the matching condition.
	//
	// 	- **type**: the matching type.
	//
	// 	- **value**: the matching value.
	//
	// example:
	//
	// [{\\"condition\\": \\"MD5\\", \\"type\\": \\"equals\\", \\"value\\": \\"0083a31cc0083a31ccf7c10367a6e783e\\"}]
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// The keyword of the alert item.
	//
	// example:
	//
	// PEM
	EventKey *string `json:"EventKey,omitempty" xml:"EventKey,omitempty"`
	// The name of the alert item.
	//
	// example:
	//
	// PEM
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The alert type.
	//
	// 	- Set the value to **sensitiveFile**.
	//
	// example:
	//
	// sensitiveFile
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The remarks that you want to add.
	//
	// example:
	//
	// test
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The operation code.
	//
	// 	- Set the value to **whitelist*	- to add the alert item to the whitelist.
	//
	// example:
	//
	// whitelist
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The application scope of the rule. The value is in the JSON format. Valid values of keys:
	//
	// 	- **type**
	//
	// 	- **value**
	//
	// example:
	//
	// {\\"type\\": \\"repo\\", \\"value\\": \\"test-aaa/shenzhen-repo-01\\"}
	Scenarios *string `json:"Scenarios,omitempty" xml:"Scenarios,omitempty"`
	// The source of the whitelist. Valid values:
	//
	// 	- **image**: image.
	//
	// 	- **agentless**: agentless detection.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s AddImageEventOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s AddImageEventOperationRequest) GoString() string {
	return s.String()
}

func (s *AddImageEventOperationRequest) GetConditions() *string {
	return s.Conditions
}

func (s *AddImageEventOperationRequest) GetEventKey() *string {
	return s.EventKey
}

func (s *AddImageEventOperationRequest) GetEventName() *string {
	return s.EventName
}

func (s *AddImageEventOperationRequest) GetEventType() *string {
	return s.EventType
}

func (s *AddImageEventOperationRequest) GetNote() *string {
	return s.Note
}

func (s *AddImageEventOperationRequest) GetOperationCode() *string {
	return s.OperationCode
}

func (s *AddImageEventOperationRequest) GetScenarios() *string {
	return s.Scenarios
}

func (s *AddImageEventOperationRequest) GetSource() *string {
	return s.Source
}

func (s *AddImageEventOperationRequest) SetConditions(v string) *AddImageEventOperationRequest {
	s.Conditions = &v
	return s
}

func (s *AddImageEventOperationRequest) SetEventKey(v string) *AddImageEventOperationRequest {
	s.EventKey = &v
	return s
}

func (s *AddImageEventOperationRequest) SetEventName(v string) *AddImageEventOperationRequest {
	s.EventName = &v
	return s
}

func (s *AddImageEventOperationRequest) SetEventType(v string) *AddImageEventOperationRequest {
	s.EventType = &v
	return s
}

func (s *AddImageEventOperationRequest) SetNote(v string) *AddImageEventOperationRequest {
	s.Note = &v
	return s
}

func (s *AddImageEventOperationRequest) SetOperationCode(v string) *AddImageEventOperationRequest {
	s.OperationCode = &v
	return s
}

func (s *AddImageEventOperationRequest) SetScenarios(v string) *AddImageEventOperationRequest {
	s.Scenarios = &v
	return s
}

func (s *AddImageEventOperationRequest) SetSource(v string) *AddImageEventOperationRequest {
	s.Source = &v
	return s
}

func (s *AddImageEventOperationRequest) Validate() error {
	return dara.Validate(s)
}
