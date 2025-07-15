// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateExportWordTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateExportWordTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateExportWordTaskResponse
	GetStatusCode() *int32
	SetBody(v *GenerateExportWordTaskResponseBody) *GenerateExportWordTaskResponse
	GetBody() *GenerateExportWordTaskResponseBody
}

type GenerateExportWordTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateExportWordTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateExportWordTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateExportWordTaskResponse) GoString() string {
	return s.String()
}

func (s *GenerateExportWordTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateExportWordTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateExportWordTaskResponse) GetBody() *GenerateExportWordTaskResponseBody {
	return s.Body
}

func (s *GenerateExportWordTaskResponse) SetHeaders(v map[string]*string) *GenerateExportWordTaskResponse {
	s.Headers = v
	return s
}

func (s *GenerateExportWordTaskResponse) SetStatusCode(v int32) *GenerateExportWordTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateExportWordTaskResponse) SetBody(v *GenerateExportWordTaskResponseBody) *GenerateExportWordTaskResponse {
	s.Body = v
	return s
}

func (s *GenerateExportWordTaskResponse) Validate() error {
	return dara.Validate(s)
}
