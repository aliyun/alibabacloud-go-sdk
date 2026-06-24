// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSourceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteOpenSourceAccountRequest
	GetInstanceId() *string
	SetUserName(v string) *DeleteOpenSourceAccountRequest
	GetUserName() *string
}

type DeleteOpenSourceAccountRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9n***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The username.
	//
	// example:
	//
	// KcUt5e5TbhkwaLyLl
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteOpenSourceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSourceAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteOpenSourceAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteOpenSourceAccountRequest) GetUserName() *string {
	return s.UserName
}

func (s *DeleteOpenSourceAccountRequest) SetInstanceId(v string) *DeleteOpenSourceAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteOpenSourceAccountRequest) SetUserName(v string) *DeleteOpenSourceAccountRequest {
	s.UserName = &v
	return s
}

func (s *DeleteOpenSourceAccountRequest) Validate() error {
	return dara.Validate(s)
}
