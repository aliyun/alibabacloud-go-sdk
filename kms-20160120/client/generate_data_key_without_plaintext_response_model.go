// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDataKeyWithoutPlaintextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateDataKeyWithoutPlaintextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateDataKeyWithoutPlaintextResponse
	GetStatusCode() *int32
	SetBody(v *GenerateDataKeyWithoutPlaintextResponseBody) *GenerateDataKeyWithoutPlaintextResponse
	GetBody() *GenerateDataKeyWithoutPlaintextResponseBody
}

type GenerateDataKeyWithoutPlaintextResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDataKeyWithoutPlaintextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateDataKeyWithoutPlaintextResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateDataKeyWithoutPlaintextResponse) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyWithoutPlaintextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateDataKeyWithoutPlaintextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateDataKeyWithoutPlaintextResponse) GetBody() *GenerateDataKeyWithoutPlaintextResponseBody {
	return s.Body
}

func (s *GenerateDataKeyWithoutPlaintextResponse) SetHeaders(v map[string]*string) *GenerateDataKeyWithoutPlaintextResponse {
	s.Headers = v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponse) SetStatusCode(v int32) *GenerateDataKeyWithoutPlaintextResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponse) SetBody(v *GenerateDataKeyWithoutPlaintextResponseBody) *GenerateDataKeyWithoutPlaintextResponse {
	s.Body = v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponse) Validate() error {
	return dara.Validate(s)
}
