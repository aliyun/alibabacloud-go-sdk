// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivilegesOfUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListPrivilegesOfUserRequest
	GetInstanceId() *string
}

type ListPrivilegesOfUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListPrivilegesOfUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivilegesOfUserRequest) GoString() string {
	return s.String()
}

func (s *ListPrivilegesOfUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPrivilegesOfUserRequest) SetInstanceId(v string) *ListPrivilegesOfUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListPrivilegesOfUserRequest) Validate() error {
	return dara.Validate(s)
}
