// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTrafficControlTaskCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *GenerateTrafficControlTaskCodeRequest
	GetEnvironment() *string
	SetInstanceId(v string) *GenerateTrafficControlTaskCodeRequest
	GetInstanceId() *string
}

type GenerateTrafficControlTaskCodeRequest struct {
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GenerateTrafficControlTaskCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateTrafficControlTaskCodeRequest) GoString() string {
	return s.String()
}

func (s *GenerateTrafficControlTaskCodeRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *GenerateTrafficControlTaskCodeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateTrafficControlTaskCodeRequest) SetEnvironment(v string) *GenerateTrafficControlTaskCodeRequest {
	s.Environment = &v
	return s
}

func (s *GenerateTrafficControlTaskCodeRequest) SetInstanceId(v string) *GenerateTrafficControlTaskCodeRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateTrafficControlTaskCodeRequest) Validate() error {
	return dara.Validate(s)
}
