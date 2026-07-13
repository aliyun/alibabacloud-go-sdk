// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ResetUserPasswordRequest
	GetInstanceId() *string
	SetName(v string) *ResetUserPasswordRequest
	GetName() *string
	SetPassword(v string) *ResetUserPasswordRequest
	GetPassword() *string
}

type ResetUserPasswordRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ResetUserPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetUserPasswordRequest) GetName() *string {
	return s.Name
}

func (s *ResetUserPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *ResetUserPasswordRequest) SetInstanceId(v string) *ResetUserPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetUserPasswordRequest) SetName(v string) *ResetUserPasswordRequest {
	s.Name = &v
	return s
}

func (s *ResetUserPasswordRequest) SetPassword(v string) *ResetUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *ResetUserPasswordRequest) Validate() error {
	return dara.Validate(s)
}
