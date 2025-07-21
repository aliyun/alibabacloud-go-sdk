// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDataKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateDataKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateDataKeyResponse
	GetStatusCode() *int32
	SetBody(v *GenerateDataKeyResponseBody) *GenerateDataKeyResponse
	GetBody() *GenerateDataKeyResponseBody
}

type GenerateDataKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateDataKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateDataKeyResponse) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateDataKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateDataKeyResponse) GetBody() *GenerateDataKeyResponseBody {
	return s.Body
}

func (s *GenerateDataKeyResponse) SetHeaders(v map[string]*string) *GenerateDataKeyResponse {
	s.Headers = v
	return s
}

func (s *GenerateDataKeyResponse) SetStatusCode(v int32) *GenerateDataKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDataKeyResponse) SetBody(v *GenerateDataKeyResponseBody) *GenerateDataKeyResponse {
	s.Body = v
	return s
}

func (s *GenerateDataKeyResponse) Validate() error {
	return dara.Validate(s)
}
