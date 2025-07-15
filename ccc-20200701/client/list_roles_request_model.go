// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRolesRequest
	GetInstanceId() *string
}

type ListRolesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRolesRequest) SetInstanceId(v string) *ListRolesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRolesRequest) Validate() error {
	return dara.Validate(s)
}
