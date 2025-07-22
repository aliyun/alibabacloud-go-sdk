// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetUserCommandRequest
	GetInstanceId() *string
	SetToken(v string) *GetUserCommandRequest
	GetToken() *string
}

type GetUserCommandRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetUserCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserCommandRequest) GoString() string {
	return s.String()
}

func (s *GetUserCommandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserCommandRequest) GetToken() *string {
	return s.Token
}

func (s *GetUserCommandRequest) SetInstanceId(v string) *GetUserCommandRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserCommandRequest) SetToken(v string) *GetUserCommandRequest {
	s.Token = &v
	return s
}

func (s *GetUserCommandRequest) Validate() error {
	return dara.Validate(s)
}
