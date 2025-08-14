// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProgramResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProgram(v *ChannelAssemblyProgram) *UpdateProgramResponseBody
	GetProgram() *ChannelAssemblyProgram
	SetRequestId(v string) *UpdateProgramResponseBody
	GetRequestId() *string
}

type UpdateProgramResponseBody struct {
	// The information about the program.
	Program *ChannelAssemblyProgram `json:"Program,omitempty" xml:"Program,omitempty"`
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProgramResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProgramResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProgramResponseBody) GetProgram() *ChannelAssemblyProgram {
	return s.Program
}

func (s *UpdateProgramResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProgramResponseBody) SetProgram(v *ChannelAssemblyProgram) *UpdateProgramResponseBody {
	s.Program = v
	return s
}

func (s *UpdateProgramResponseBody) SetRequestId(v string) *UpdateProgramResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProgramResponseBody) Validate() error {
	return dara.Validate(s)
}
