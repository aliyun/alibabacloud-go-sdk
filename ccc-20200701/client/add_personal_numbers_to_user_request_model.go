// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPersonalNumbersToUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddPersonalNumbersToUserRequest
	GetInstanceId() *string
	SetNumberList(v string) *AddPersonalNumbersToUserRequest
	GetNumberList() *string
	SetUserId(v string) *AddPersonalNumbersToUserRequest
	GetUserId() *string
}

type AddPersonalNumbersToUserRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// List of numbers to be added. The numbers must be attachable to the agent corresponding to UserId. If you are unsure which numbers the agent can attach, refer to the ListPersonalNumbersOfUser API.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["0101234****","0105678****"]
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	// Agent ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddPersonalNumbersToUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPersonalNumbersToUserRequest) GoString() string {
	return s.String()
}

func (s *AddPersonalNumbersToUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddPersonalNumbersToUserRequest) GetNumberList() *string {
	return s.NumberList
}

func (s *AddPersonalNumbersToUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddPersonalNumbersToUserRequest) SetInstanceId(v string) *AddPersonalNumbersToUserRequest {
	s.InstanceId = &v
	return s
}

func (s *AddPersonalNumbersToUserRequest) SetNumberList(v string) *AddPersonalNumbersToUserRequest {
	s.NumberList = &v
	return s
}

func (s *AddPersonalNumbersToUserRequest) SetUserId(v string) *AddPersonalNumbersToUserRequest {
	s.UserId = &v
	return s
}

func (s *AddPersonalNumbersToUserRequest) Validate() error {
	return dara.Validate(s)
}
