// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePbcInstructionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePbcInstructionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePbcInstructionResponse
	GetStatusCode() *int32
	SetBody(v *PbcInstruction) *UpdatePbcInstructionResponse
	GetBody() *PbcInstruction
}

type UpdatePbcInstructionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcInstruction    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePbcInstructionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePbcInstructionResponse) GoString() string {
	return s.String()
}

func (s *UpdatePbcInstructionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePbcInstructionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePbcInstructionResponse) GetBody() *PbcInstruction {
	return s.Body
}

func (s *UpdatePbcInstructionResponse) SetHeaders(v map[string]*string) *UpdatePbcInstructionResponse {
	s.Headers = v
	return s
}

func (s *UpdatePbcInstructionResponse) SetStatusCode(v int32) *UpdatePbcInstructionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePbcInstructionResponse) SetBody(v *PbcInstruction) *UpdatePbcInstructionResponse {
	s.Body = v
	return s
}

func (s *UpdatePbcInstructionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
