// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFCTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventMetaName(v string) *AddFCTriggerRequest
	GetEventMetaName() *string
	SetEventMetaVersion(v string) *AddFCTriggerRequest
	GetEventMetaVersion() *string
	SetFunctionARN(v string) *AddFCTriggerRequest
	GetFunctionARN() *string
	SetNotes(v string) *AddFCTriggerRequest
	GetNotes() *string
	SetRoleARN(v string) *AddFCTriggerRequest
	GetRoleARN() *string
	SetSourceARN(v string) *AddFCTriggerRequest
	GetSourceARN() *string
	SetTriggerARN(v string) *AddFCTriggerRequest
	GetTriggerARN() *string
}

type AddFCTriggerRequest struct {
	// The name of the event.
	//
	// This parameter is required.
	//
	// example:
	//
	// LogFileCreated
	EventMetaName *string `json:"EventMetaName,omitempty" xml:"EventMetaName,omitempty"`
	// The version of the event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	EventMetaVersion *string `json:"EventMetaVersion,omitempty" xml:"EventMetaVersion,omitempty"`
	// The feature trigger.
	//
	// example:
	//
	// acs:fc:1223455566666:123:services/myservice/functions/myfunction
	FunctionARN *string `json:"FunctionARN,omitempty" xml:"FunctionARN,omitempty"`
	// The remarks.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Notes *string `json:"Notes,omitempty" xml:"Notes,omitempty"`
	// The assigned Resource Access Management (RAM) role.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ram:: 1234567890:role/aliyuncdneventnotificationrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// The resources and filters for event listening.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:cdn:*:1234567890:domain/example.com
	SourceARN *string `json:"SourceARN,omitempty" xml:"SourceARN,omitempty"`
	// The trigger that corresponds to the Function Compute service.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:fc:cn-beijing: 1234567890:services/FCTestService/functions/printEvent/triggers/testtrigger
	TriggerARN *string `json:"TriggerARN,omitempty" xml:"TriggerARN,omitempty"`
}

func (s AddFCTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFCTriggerRequest) GoString() string {
	return s.String()
}

func (s *AddFCTriggerRequest) GetEventMetaName() *string {
	return s.EventMetaName
}

func (s *AddFCTriggerRequest) GetEventMetaVersion() *string {
	return s.EventMetaVersion
}

func (s *AddFCTriggerRequest) GetFunctionARN() *string {
	return s.FunctionARN
}

func (s *AddFCTriggerRequest) GetNotes() *string {
	return s.Notes
}

func (s *AddFCTriggerRequest) GetRoleARN() *string {
	return s.RoleARN
}

func (s *AddFCTriggerRequest) GetSourceARN() *string {
	return s.SourceARN
}

func (s *AddFCTriggerRequest) GetTriggerARN() *string {
	return s.TriggerARN
}

func (s *AddFCTriggerRequest) SetEventMetaName(v string) *AddFCTriggerRequest {
	s.EventMetaName = &v
	return s
}

func (s *AddFCTriggerRequest) SetEventMetaVersion(v string) *AddFCTriggerRequest {
	s.EventMetaVersion = &v
	return s
}

func (s *AddFCTriggerRequest) SetFunctionARN(v string) *AddFCTriggerRequest {
	s.FunctionARN = &v
	return s
}

func (s *AddFCTriggerRequest) SetNotes(v string) *AddFCTriggerRequest {
	s.Notes = &v
	return s
}

func (s *AddFCTriggerRequest) SetRoleARN(v string) *AddFCTriggerRequest {
	s.RoleARN = &v
	return s
}

func (s *AddFCTriggerRequest) SetSourceARN(v string) *AddFCTriggerRequest {
	s.SourceARN = &v
	return s
}

func (s *AddFCTriggerRequest) SetTriggerARN(v string) *AddFCTriggerRequest {
	s.TriggerARN = &v
	return s
}

func (s *AddFCTriggerRequest) Validate() error {
	return dara.Validate(s)
}
