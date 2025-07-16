// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFCTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionARN(v string) *UpdateFCTriggerRequest
	GetFunctionARN() *string
	SetNotes(v string) *UpdateFCTriggerRequest
	GetNotes() *string
	SetRoleARN(v string) *UpdateFCTriggerRequest
	GetRoleARN() *string
	SetSourceARN(v string) *UpdateFCTriggerRequest
	GetSourceARN() *string
	SetTriggerARN(v string) *UpdateFCTriggerRequest
	GetTriggerARN() *string
}

type UpdateFCTriggerRequest struct {
	// The feature trigger.
	//
	// example:
	//
	// acs:fc:1223455566666:123:services/myservice/functions/myfunction
	FunctionARN *string `json:"FunctionARN,omitempty" xml:"FunctionARN,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test
	Notes *string `json:"Notes,omitempty" xml:"Notes,omitempty"`
	// The assigned RAM role.
	//
	// example:
	//
	// acs:ram:: 1234567890:role/aliyuncdneventnotificationrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// The resources and filters for event listening.
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

func (s UpdateFCTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFCTriggerRequest) GoString() string {
	return s.String()
}

func (s *UpdateFCTriggerRequest) GetFunctionARN() *string {
	return s.FunctionARN
}

func (s *UpdateFCTriggerRequest) GetNotes() *string {
	return s.Notes
}

func (s *UpdateFCTriggerRequest) GetRoleARN() *string {
	return s.RoleARN
}

func (s *UpdateFCTriggerRequest) GetSourceARN() *string {
	return s.SourceARN
}

func (s *UpdateFCTriggerRequest) GetTriggerARN() *string {
	return s.TriggerARN
}

func (s *UpdateFCTriggerRequest) SetFunctionARN(v string) *UpdateFCTriggerRequest {
	s.FunctionARN = &v
	return s
}

func (s *UpdateFCTriggerRequest) SetNotes(v string) *UpdateFCTriggerRequest {
	s.Notes = &v
	return s
}

func (s *UpdateFCTriggerRequest) SetRoleARN(v string) *UpdateFCTriggerRequest {
	s.RoleARN = &v
	return s
}

func (s *UpdateFCTriggerRequest) SetSourceARN(v string) *UpdateFCTriggerRequest {
	s.SourceARN = &v
	return s
}

func (s *UpdateFCTriggerRequest) SetTriggerARN(v string) *UpdateFCTriggerRequest {
	s.TriggerARN = &v
	return s
}

func (s *UpdateFCTriggerRequest) Validate() error {
	return dara.Validate(s)
}
