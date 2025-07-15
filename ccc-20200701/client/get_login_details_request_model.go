// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatDeviceId(v string) *GetLoginDetailsRequest
	GetChatDeviceId() *string
	SetInstanceId(v string) *GetLoginDetailsRequest
	GetInstanceId() *string
	SetUserId(v string) *GetLoginDetailsRequest
	GetUserId() *string
}

type GetLoginDetailsRequest struct {
	ChatDeviceId *string `json:"ChatDeviceId,omitempty" xml:"ChatDeviceId,omitempty"`
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

func (s GetLoginDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLoginDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetLoginDetailsRequest) GetChatDeviceId() *string {
	return s.ChatDeviceId
}

func (s *GetLoginDetailsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLoginDetailsRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetLoginDetailsRequest) SetChatDeviceId(v string) *GetLoginDetailsRequest {
	s.ChatDeviceId = &v
	return s
}

func (s *GetLoginDetailsRequest) SetInstanceId(v string) *GetLoginDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLoginDetailsRequest) SetUserId(v string) *GetLoginDetailsRequest {
	s.UserId = &v
	return s
}

func (s *GetLoginDetailsRequest) Validate() error {
	return dara.Validate(s)
}
