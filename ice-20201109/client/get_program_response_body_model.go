// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProgramResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProgram(v *ChannelAssemblyProgram) *GetProgramResponseBody
	GetProgram() *ChannelAssemblyProgram
	SetRequestId(v string) *GetProgramResponseBody
	GetRequestId() *string
}

type GetProgramResponseBody struct {
	// The information about the program.
	Program *ChannelAssemblyProgram `json:"Program,omitempty" xml:"Program,omitempty"`
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProgramResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProgramResponseBody) GoString() string {
	return s.String()
}

func (s *GetProgramResponseBody) GetProgram() *ChannelAssemblyProgram {
	return s.Program
}

func (s *GetProgramResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProgramResponseBody) SetProgram(v *ChannelAssemblyProgram) *GetProgramResponseBody {
	s.Program = v
	return s
}

func (s *GetProgramResponseBody) SetRequestId(v string) *GetProgramResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProgramResponseBody) Validate() error {
	if s.Program != nil {
		if err := s.Program.Validate(); err != nil {
			return err
		}
	}
	return nil
}
