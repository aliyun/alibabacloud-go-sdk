// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPbcInstructionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPbcInstructionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPbcInstructionResponse
	GetStatusCode() *int32
	SetBody(v *PbcInstruction) *GetPbcInstructionResponse
	GetBody() *PbcInstruction
}

type GetPbcInstructionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcInstruction    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPbcInstructionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPbcInstructionResponse) GoString() string {
	return s.String()
}

func (s *GetPbcInstructionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPbcInstructionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPbcInstructionResponse) GetBody() *PbcInstruction {
	return s.Body
}

func (s *GetPbcInstructionResponse) SetHeaders(v map[string]*string) *GetPbcInstructionResponse {
	s.Headers = v
	return s
}

func (s *GetPbcInstructionResponse) SetStatusCode(v int32) *GetPbcInstructionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPbcInstructionResponse) SetBody(v *PbcInstruction) *GetPbcInstructionResponse {
	s.Body = v
	return s
}

func (s *GetPbcInstructionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
