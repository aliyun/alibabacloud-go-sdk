// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtension(v string) *GetUserRequest
	GetExtension() *string
	SetInstanceId(v string) *GetUserRequest
	GetInstanceId() *string
	SetUserId(v string) *GetUserRequest
	GetUserId() *string
}

type GetUserRequest struct {
	// example:
	//
	// 8003****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) GetExtension() *string {
	return s.Extension
}

func (s *GetUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetUserRequest) SetExtension(v string) *GetUserRequest {
	s.Extension = &v
	return s
}

func (s *GetUserRequest) SetInstanceId(v string) *GetUserRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

func (s *GetUserRequest) Validate() error {
	return dara.Validate(s)
}
