// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEditContactFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactFlowId(v string) *StartEditContactFlowRequest
	GetContactFlowId() *string
	SetInstanceId(v string) *StartEditContactFlowRequest
	GetInstanceId() *string
}

type StartEditContactFlowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 798e83a9-5140-4039-afa1-761ca4cca2df
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartEditContactFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s StartEditContactFlowRequest) GoString() string {
	return s.String()
}

func (s *StartEditContactFlowRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *StartEditContactFlowRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartEditContactFlowRequest) SetContactFlowId(v string) *StartEditContactFlowRequest {
	s.ContactFlowId = &v
	return s
}

func (s *StartEditContactFlowRequest) SetInstanceId(v string) *StartEditContactFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *StartEditContactFlowRequest) Validate() error {
	return dara.Validate(s)
}
