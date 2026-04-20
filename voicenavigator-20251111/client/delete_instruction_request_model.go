// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstructionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteInstructionRequest
	GetCode() *string
	SetInstanceId(v string) *DeleteInstructionRequest
	GetInstanceId() *string
}

type DeleteInstructionRequest struct {
	// example:
	//
	// Transfer00
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteInstructionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstructionRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstructionRequest) GetCode() *string {
	return s.Code
}

func (s *DeleteInstructionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstructionRequest) SetCode(v string) *DeleteInstructionRequest {
	s.Code = &v
	return s
}

func (s *DeleteInstructionRequest) SetInstanceId(v string) *DeleteInstructionRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstructionRequest) Validate() error {
	return dara.Validate(s)
}
