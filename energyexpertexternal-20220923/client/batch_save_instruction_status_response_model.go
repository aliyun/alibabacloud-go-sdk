// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSaveInstructionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSaveInstructionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSaveInstructionStatusResponse
	GetStatusCode() *int32
	SetBody(v *BatchSaveInstructionStatusResponseBody) *BatchSaveInstructionStatusResponse
	GetBody() *BatchSaveInstructionStatusResponseBody
}

type BatchSaveInstructionStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSaveInstructionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSaveInstructionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSaveInstructionStatusResponse) GoString() string {
	return s.String()
}

func (s *BatchSaveInstructionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSaveInstructionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSaveInstructionStatusResponse) GetBody() *BatchSaveInstructionStatusResponseBody {
	return s.Body
}

func (s *BatchSaveInstructionStatusResponse) SetHeaders(v map[string]*string) *BatchSaveInstructionStatusResponse {
	s.Headers = v
	return s
}

func (s *BatchSaveInstructionStatusResponse) SetStatusCode(v int32) *BatchSaveInstructionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSaveInstructionStatusResponse) SetBody(v *BatchSaveInstructionStatusResponseBody) *BatchSaveInstructionStatusResponse {
	s.Body = v
	return s
}

func (s *BatchSaveInstructionStatusResponse) Validate() error {
	return dara.Validate(s)
}
