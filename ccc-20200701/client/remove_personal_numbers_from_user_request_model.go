// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePersonalNumbersFromUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RemovePersonalNumbersFromUserRequest
	GetInstanceId() *string
	SetNumberList(v string) *RemovePersonalNumbersFromUserRequest
	GetNumberList() *string
	SetUserId(v string) *RemovePersonalNumbersFromUserRequest
	GetUserId() *string
}

type RemovePersonalNumbersFromUserRequest struct {
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
	// ["0101234****", "0105678****"]
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemovePersonalNumbersFromUserRequest) String() string {
	return dara.Prettify(s)
}

func (s RemovePersonalNumbersFromUserRequest) GoString() string {
	return s.String()
}

func (s *RemovePersonalNumbersFromUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemovePersonalNumbersFromUserRequest) GetNumberList() *string {
	return s.NumberList
}

func (s *RemovePersonalNumbersFromUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *RemovePersonalNumbersFromUserRequest) SetInstanceId(v string) *RemovePersonalNumbersFromUserRequest {
	s.InstanceId = &v
	return s
}

func (s *RemovePersonalNumbersFromUserRequest) SetNumberList(v string) *RemovePersonalNumbersFromUserRequest {
	s.NumberList = &v
	return s
}

func (s *RemovePersonalNumbersFromUserRequest) SetUserId(v string) *RemovePersonalNumbersFromUserRequest {
	s.UserId = &v
	return s
}

func (s *RemovePersonalNumbersFromUserRequest) Validate() error {
	return dara.Validate(s)
}
