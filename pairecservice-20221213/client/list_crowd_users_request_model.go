// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrowdUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListCrowdUsersRequest
	GetInstanceId() *string
}

type ListCrowdUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListCrowdUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCrowdUsersRequest) GoString() string {
	return s.String()
}

func (s *ListCrowdUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCrowdUsersRequest) SetInstanceId(v string) *ListCrowdUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCrowdUsersRequest) Validate() error {
	return dara.Validate(s)
}
