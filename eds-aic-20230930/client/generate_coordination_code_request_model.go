// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCoordinationCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GenerateCoordinationCodeRequest
	GetInstanceId() *string
	SetOwnerUserId(v string) *GenerateCoordinationCodeRequest
	GetOwnerUserId() *string
}

type GenerateCoordinationCodeRequest struct {
	// The ID of the instance.
	//
	// example:
	//
	// acp-2zecay9ponatdc4m****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the user to whom the current instance is assigned.
	//
	// example:
	//
	// xiaoming
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
}

func (s GenerateCoordinationCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCoordinationCodeRequest) GoString() string {
	return s.String()
}

func (s *GenerateCoordinationCodeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateCoordinationCodeRequest) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *GenerateCoordinationCodeRequest) SetInstanceId(v string) *GenerateCoordinationCodeRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) SetOwnerUserId(v string) *GenerateCoordinationCodeRequest {
	s.OwnerUserId = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) Validate() error {
	return dara.Validate(s)
}
