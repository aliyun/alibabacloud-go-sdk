// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetUserRequest
	GetInstanceId() *string
	SetName(v string) *GetUserRequest
	GetName() *string
}

type GetUserRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserRequest) GetName() *string {
	return s.Name
}

func (s *GetUserRequest) SetInstanceId(v string) *GetUserRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserRequest) SetName(v string) *GetUserRequest {
	s.Name = &v
	return s
}

func (s *GetUserRequest) Validate() error {
	return dara.Validate(s)
}
