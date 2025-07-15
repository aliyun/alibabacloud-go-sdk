// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePhoneNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *RemovePhoneNumbersRequest
	GetForce() *bool
	SetInstanceId(v string) *RemovePhoneNumbersRequest
	GetInstanceId() *string
	SetNumberList(v string) *RemovePhoneNumbersRequest
	GetNumberList() *string
}

type RemovePhoneNumbersRequest struct {
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
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
	// ["0101234****","0105678****"]
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
}

func (s RemovePhoneNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s RemovePhoneNumbersRequest) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumbersRequest) GetForce() *bool {
	return s.Force
}

func (s *RemovePhoneNumbersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemovePhoneNumbersRequest) GetNumberList() *string {
	return s.NumberList
}

func (s *RemovePhoneNumbersRequest) SetForce(v bool) *RemovePhoneNumbersRequest {
	s.Force = &v
	return s
}

func (s *RemovePhoneNumbersRequest) SetInstanceId(v string) *RemovePhoneNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *RemovePhoneNumbersRequest) SetNumberList(v string) *RemovePhoneNumbersRequest {
	s.NumberList = &v
	return s
}

func (s *RemovePhoneNumbersRequest) Validate() error {
	return dara.Validate(s)
}
