// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageEventOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v string) *UpdateImageEventOperationRequest
	GetConditions() *string
	SetId(v int64) *UpdateImageEventOperationRequest
	GetId() *int64
	SetNote(v string) *UpdateImageEventOperationRequest
	GetNote() *string
	SetScenarios(v string) *UpdateImageEventOperationRequest
	GetScenarios() *string
	SetSource(v string) *UpdateImageEventOperationRequest
	GetSource() *string
}

type UpdateImageEventOperationRequest struct {
	// The rule conditions. Specify a value in the JSON format. You can specify the following keys:
	//
	// 	- **condition**: the matching condition.
	//
	// 	- **type**: the matching type.
	//
	// 	- **value**: the matching value.
	//
	// example:
	//
	// [{\\"condition\\": \\"MD5\\", \\"type\\": \\"equals\\", \\"value\\": \\"0083a31cc0083a31ccf7c10367a6e****\\"}]
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// The ID of the alert handling rule.
	//
	// > You can call the [DescribeImageEventOperationPage](~~DescribeImageEventOperationPage~~) operation to query the ID.
	//
	// example:
	//
	// 814163
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The remarks that you want to add.
	//
	// example:
	//
	// test
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The application scope of the rule. The value is in the JSON format. Valid values:
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

func (s UpdateImageEventOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageEventOperationRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageEventOperationRequest) GetConditions() *string {
	return s.Conditions
}

func (s *UpdateImageEventOperationRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateImageEventOperationRequest) GetNote() *string {
	return s.Note
}

func (s *UpdateImageEventOperationRequest) GetScenarios() *string {
	return s.Scenarios
}

func (s *UpdateImageEventOperationRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateImageEventOperationRequest) SetConditions(v string) *UpdateImageEventOperationRequest {
	s.Conditions = &v
	return s
}

func (s *UpdateImageEventOperationRequest) SetId(v int64) *UpdateImageEventOperationRequest {
	s.Id = &v
	return s
}

func (s *UpdateImageEventOperationRequest) SetNote(v string) *UpdateImageEventOperationRequest {
	s.Note = &v
	return s
}

func (s *UpdateImageEventOperationRequest) SetScenarios(v string) *UpdateImageEventOperationRequest {
	s.Scenarios = &v
	return s
}

func (s *UpdateImageEventOperationRequest) SetSource(v string) *UpdateImageEventOperationRequest {
	s.Source = &v
	return s
}

func (s *UpdateImageEventOperationRequest) Validate() error {
	return dara.Validate(s)
}
