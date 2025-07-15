// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateViewPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateViewPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateViewPointResponse
	GetStatusCode() *int32
	SetBody(v *GenerateViewPointResponseBody) *GenerateViewPointResponse
	GetBody() *GenerateViewPointResponseBody
}

type GenerateViewPointResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateViewPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateViewPointResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateViewPointResponse) GoString() string {
	return s.String()
}

func (s *GenerateViewPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateViewPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateViewPointResponse) GetBody() *GenerateViewPointResponseBody {
	return s.Body
}

func (s *GenerateViewPointResponse) SetHeaders(v map[string]*string) *GenerateViewPointResponse {
	s.Headers = v
	return s
}

func (s *GenerateViewPointResponse) SetStatusCode(v int32) *GenerateViewPointResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateViewPointResponse) SetBody(v *GenerateViewPointResponseBody) *GenerateViewPointResponse {
	s.Body = v
	return s
}

func (s *GenerateViewPointResponse) Validate() error {
	return dara.Validate(s)
}
