// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTemplateByScratchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateTemplateByScratchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateTemplateByScratchResponse
	GetStatusCode() *int32
	SetBody(v *GenerateTemplateByScratchResponseBody) *GenerateTemplateByScratchResponse
	GetBody() *GenerateTemplateByScratchResponseBody
}

type GenerateTemplateByScratchResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateTemplateByScratchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateTemplateByScratchResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateTemplateByScratchResponse) GoString() string {
	return s.String()
}

func (s *GenerateTemplateByScratchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateTemplateByScratchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateTemplateByScratchResponse) GetBody() *GenerateTemplateByScratchResponseBody {
	return s.Body
}

func (s *GenerateTemplateByScratchResponse) SetHeaders(v map[string]*string) *GenerateTemplateByScratchResponse {
	s.Headers = v
	return s
}

func (s *GenerateTemplateByScratchResponse) SetStatusCode(v int32) *GenerateTemplateByScratchResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTemplateByScratchResponse) SetBody(v *GenerateTemplateByScratchResponseBody) *GenerateTemplateByScratchResponse {
	s.Body = v
	return s
}

func (s *GenerateTemplateByScratchResponse) Validate() error {
	return dara.Validate(s)
}
