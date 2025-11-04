// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateKMSDataKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateKMSDataKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateKMSDataKeyResponse
	GetStatusCode() *int32
	SetBody(v *GenerateKMSDataKeyResponseBody) *GenerateKMSDataKeyResponse
	GetBody() *GenerateKMSDataKeyResponseBody
}

type GenerateKMSDataKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateKMSDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateKMSDataKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateKMSDataKeyResponse) GoString() string {
	return s.String()
}

func (s *GenerateKMSDataKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateKMSDataKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateKMSDataKeyResponse) GetBody() *GenerateKMSDataKeyResponseBody {
	return s.Body
}

func (s *GenerateKMSDataKeyResponse) SetHeaders(v map[string]*string) *GenerateKMSDataKeyResponse {
	s.Headers = v
	return s
}

func (s *GenerateKMSDataKeyResponse) SetStatusCode(v int32) *GenerateKMSDataKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateKMSDataKeyResponse) SetBody(v *GenerateKMSDataKeyResponseBody) *GenerateKMSDataKeyResponse {
	s.Body = v
	return s
}

func (s *GenerateKMSDataKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
