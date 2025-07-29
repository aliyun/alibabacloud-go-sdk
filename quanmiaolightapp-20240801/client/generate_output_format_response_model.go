// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOutputFormatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateOutputFormatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateOutputFormatResponse
	GetStatusCode() *int32
	SetBody(v *GenerateOutputFormatResponseBody) *GenerateOutputFormatResponse
	GetBody() *GenerateOutputFormatResponseBody
}

type GenerateOutputFormatResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateOutputFormatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateOutputFormatResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateOutputFormatResponse) GoString() string {
	return s.String()
}

func (s *GenerateOutputFormatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateOutputFormatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateOutputFormatResponse) GetBody() *GenerateOutputFormatResponseBody {
	return s.Body
}

func (s *GenerateOutputFormatResponse) SetHeaders(v map[string]*string) *GenerateOutputFormatResponse {
	s.Headers = v
	return s
}

func (s *GenerateOutputFormatResponse) SetStatusCode(v int32) *GenerateOutputFormatResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateOutputFormatResponse) SetBody(v *GenerateOutputFormatResponseBody) *GenerateOutputFormatResponse {
	s.Body = v
	return s
}

func (s *GenerateOutputFormatResponse) Validate() error {
	return dara.Validate(s)
}
