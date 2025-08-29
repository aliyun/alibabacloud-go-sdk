// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTrafficControlTaskConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GenerateTrafficControlTaskConfigRequest
	GetInstanceId() *string
}

type GenerateTrafficControlTaskConfigRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GenerateTrafficControlTaskConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateTrafficControlTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *GenerateTrafficControlTaskConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateTrafficControlTaskConfigRequest) SetInstanceId(v string) *GenerateTrafficControlTaskConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateTrafficControlTaskConfigRequest) Validate() error {
	return dara.Validate(s)
}
