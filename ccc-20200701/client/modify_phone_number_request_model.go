// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPhoneNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactFlowId(v string) *ModifyPhoneNumberRequest
	GetContactFlowId() *string
	SetInstanceId(v string) *ModifyPhoneNumberRequest
	GetInstanceId() *string
	SetNumber(v string) *ModifyPhoneNumberRequest
	GetNumber() *string
	SetUsage(v string) *ModifyPhoneNumberRequest
	GetUsage() *string
}

type ModifyPhoneNumberRequest struct {
	// example:
	//
	// 78128960-bb00-4ddc-8a82-923a8c5bd22d
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0102134****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bidirection
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ModifyPhoneNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *ModifyPhoneNumberRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *ModifyPhoneNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPhoneNumberRequest) GetNumber() *string {
	return s.Number
}

func (s *ModifyPhoneNumberRequest) GetUsage() *string {
	return s.Usage
}

func (s *ModifyPhoneNumberRequest) SetContactFlowId(v string) *ModifyPhoneNumberRequest {
	s.ContactFlowId = &v
	return s
}

func (s *ModifyPhoneNumberRequest) SetInstanceId(v string) *ModifyPhoneNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPhoneNumberRequest) SetNumber(v string) *ModifyPhoneNumberRequest {
	s.Number = &v
	return s
}

func (s *ModifyPhoneNumberRequest) SetUsage(v string) *ModifyPhoneNumberRequest {
	s.Usage = &v
	return s
}

func (s *ModifyPhoneNumberRequest) Validate() error {
	return dara.Validate(s)
}
