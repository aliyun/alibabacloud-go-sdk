// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatRoutingProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetChatRoutingProfileRequest
	GetInstanceId() *string
}

type GetChatRoutingProfileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetChatRoutingProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatRoutingProfileRequest) GoString() string {
	return s.String()
}

func (s *GetChatRoutingProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetChatRoutingProfileRequest) SetInstanceId(v string) *GetChatRoutingProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *GetChatRoutingProfileRequest) Validate() error {
	return dara.Validate(s)
}
