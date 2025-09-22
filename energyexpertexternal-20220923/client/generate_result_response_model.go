// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateResultResponse
	GetStatusCode() *int32
	SetBody(v *GenerateResultResponseBody) *GenerateResultResponse
	GetBody() *GenerateResultResponseBody
}

type GenerateResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateResultResponse) GoString() string {
	return s.String()
}

func (s *GenerateResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateResultResponse) GetBody() *GenerateResultResponseBody {
	return s.Body
}

func (s *GenerateResultResponse) SetHeaders(v map[string]*string) *GenerateResultResponse {
	s.Headers = v
	return s
}

func (s *GenerateResultResponse) SetStatusCode(v int32) *GenerateResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateResultResponse) SetBody(v *GenerateResultResponseBody) *GenerateResultResponse {
	s.Body = v
	return s
}

func (s *GenerateResultResponse) Validate() error {
	return dara.Validate(s)
}
