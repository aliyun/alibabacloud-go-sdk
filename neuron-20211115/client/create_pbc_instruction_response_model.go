// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcInstructionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePbcInstructionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePbcInstructionResponse
	GetStatusCode() *int32
	SetBody(v *PbcInstruction) *CreatePbcInstructionResponse
	GetBody() *PbcInstruction
}

type CreatePbcInstructionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcInstruction    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcInstructionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcInstructionResponse) GoString() string {
	return s.String()
}

func (s *CreatePbcInstructionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePbcInstructionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePbcInstructionResponse) GetBody() *PbcInstruction {
	return s.Body
}

func (s *CreatePbcInstructionResponse) SetHeaders(v map[string]*string) *CreatePbcInstructionResponse {
	s.Headers = v
	return s
}

func (s *CreatePbcInstructionResponse) SetStatusCode(v int32) *CreatePbcInstructionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePbcInstructionResponse) SetBody(v *PbcInstruction) *CreatePbcInstructionResponse {
	s.Body = v
	return s
}

func (s *CreatePbcInstructionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
