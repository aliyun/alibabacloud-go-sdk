// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteUserRequest
	GetInstanceId() *string
	SetName(v string) *DeleteUserRequest
	GetName() *string
}

type DeleteUserRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteUserRequest) GetName() *string {
	return s.Name
}

func (s *DeleteUserRequest) SetInstanceId(v string) *DeleteUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUserRequest) SetName(v string) *DeleteUserRequest {
	s.Name = &v
	return s
}

func (s *DeleteUserRequest) Validate() error {
	return dara.Validate(s)
}
