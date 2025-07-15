// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDoNotCallNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RemoveDoNotCallNumbersRequest
	GetInstanceId() *string
	SetNumberList(v string) *RemoveDoNotCallNumbersRequest
	GetNumberList() *string
}

type RemoveDoNotCallNumbersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ["1900000****","1312211****"]
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
}

func (s RemoveDoNotCallNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveDoNotCallNumbersRequest) GoString() string {
	return s.String()
}

func (s *RemoveDoNotCallNumbersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveDoNotCallNumbersRequest) GetNumberList() *string {
	return s.NumberList
}

func (s *RemoveDoNotCallNumbersRequest) SetInstanceId(v string) *RemoveDoNotCallNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveDoNotCallNumbersRequest) SetNumberList(v string) *RemoveDoNotCallNumbersRequest {
	s.NumberList = &v
	return s
}

func (s *RemoveDoNotCallNumbersRequest) Validate() error {
	return dara.Validate(s)
}
