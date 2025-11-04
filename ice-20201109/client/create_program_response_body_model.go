// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProgramResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProgram(v *ChannelAssemblyProgram) *CreateProgramResponseBody
	GetProgram() *ChannelAssemblyProgram
	SetRequestId(v string) *CreateProgramResponseBody
	GetRequestId() *string
}

type CreateProgramResponseBody struct {
	// The information about the program.
	Program *ChannelAssemblyProgram `json:"Program,omitempty" xml:"Program,omitempty"`
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProgramResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProgramResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProgramResponseBody) GetProgram() *ChannelAssemblyProgram {
	return s.Program
}

func (s *CreateProgramResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProgramResponseBody) SetProgram(v *ChannelAssemblyProgram) *CreateProgramResponseBody {
	s.Program = v
	return s
}

func (s *CreateProgramResponseBody) SetRequestId(v string) *CreateProgramResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProgramResponseBody) Validate() error {
	if s.Program != nil {
		if err := s.Program.Validate(); err != nil {
			return err
		}
	}
	return nil
}
