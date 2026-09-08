// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPhoneNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactFlowId(v string) *AddPhoneNumbersRequest
	GetContactFlowId() *string
	SetInstanceId(v string) *AddPhoneNumbersRequest
	GetInstanceId() *string
	SetNumberGroupId(v string) *AddPhoneNumbersRequest
	GetNumberGroupId() *string
	SetNumberList(v string) *AddPhoneNumbersRequest
	GetNumberList() *string
	SetUsage(v string) *AddPhoneNumbersRequest
	GetUsage() *string
}

type AddPhoneNumbersRequest struct {
	// ID of the IVR contact flow to attach. This parameter is valid only when the number usage includes inbound calls. It is optional and defaults to empty.
	//
	// example:
	//
	// dDMD_0mif4hv
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Number group ID. You can view number grouping information in the Cloud Contact Center console. This parameter is optional and defaults to empty.
	//
	// example:
	//
	// 2cb77c29-5f60-4b90-b21e-9d2ba9833f14
	NumberGroupId *string `json:"NumberGroupId,omitempty" xml:"NumberGroupId,omitempty"`
	// List of phone numbers to add.
	//
	// example:
	//
	// ["0101234****", "0105678****"]
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	// Usage of the phone number. Note: If the provided number is a 400 number, the usage must be set to Inbound.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bidirection
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s AddPhoneNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPhoneNumbersRequest) GoString() string {
	return s.String()
}

func (s *AddPhoneNumbersRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *AddPhoneNumbersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddPhoneNumbersRequest) GetNumberGroupId() *string {
	return s.NumberGroupId
}

func (s *AddPhoneNumbersRequest) GetNumberList() *string {
	return s.NumberList
}

func (s *AddPhoneNumbersRequest) GetUsage() *string {
	return s.Usage
}

func (s *AddPhoneNumbersRequest) SetContactFlowId(v string) *AddPhoneNumbersRequest {
	s.ContactFlowId = &v
	return s
}

func (s *AddPhoneNumbersRequest) SetInstanceId(v string) *AddPhoneNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *AddPhoneNumbersRequest) SetNumberGroupId(v string) *AddPhoneNumbersRequest {
	s.NumberGroupId = &v
	return s
}

func (s *AddPhoneNumbersRequest) SetNumberList(v string) *AddPhoneNumbersRequest {
	s.NumberList = &v
	return s
}

func (s *AddPhoneNumbersRequest) SetUsage(v string) *AddPhoneNumbersRequest {
	s.Usage = &v
	return s
}

func (s *AddPhoneNumbersRequest) Validate() error {
	return dara.Validate(s)
}
