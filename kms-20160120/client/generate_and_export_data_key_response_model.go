// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAndExportDataKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateAndExportDataKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateAndExportDataKeyResponse
	GetStatusCode() *int32
	SetBody(v *GenerateAndExportDataKeyResponseBody) *GenerateAndExportDataKeyResponse
	GetBody() *GenerateAndExportDataKeyResponseBody
}

type GenerateAndExportDataKeyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateAndExportDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateAndExportDataKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateAndExportDataKeyResponse) GoString() string {
	return s.String()
}

func (s *GenerateAndExportDataKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateAndExportDataKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateAndExportDataKeyResponse) GetBody() *GenerateAndExportDataKeyResponseBody {
	return s.Body
}

func (s *GenerateAndExportDataKeyResponse) SetHeaders(v map[string]*string) *GenerateAndExportDataKeyResponse {
	s.Headers = v
	return s
}

func (s *GenerateAndExportDataKeyResponse) SetStatusCode(v int32) *GenerateAndExportDataKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAndExportDataKeyResponse) SetBody(v *GenerateAndExportDataKeyResponseBody) *GenerateAndExportDataKeyResponse {
	s.Body = v
	return s
}

func (s *GenerateAndExportDataKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
