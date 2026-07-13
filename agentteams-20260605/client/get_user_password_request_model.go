// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetUserPasswordRequest
	GetInstanceId() *string
	SetName(v string) *GetUserPasswordRequest
	GetName() *string
}

type GetUserPasswordRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetUserPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *GetUserPasswordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserPasswordRequest) GetName() *string {
	return s.Name
}

func (s *GetUserPasswordRequest) SetInstanceId(v string) *GetUserPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserPasswordRequest) SetName(v string) *GetUserPasswordRequest {
	s.Name = &v
	return s
}

func (s *GetUserPasswordRequest) Validate() error {
	return dara.Validate(s)
}
