// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTrafficControlTaskCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateTrafficControlTaskCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateTrafficControlTaskCodeResponse
	GetStatusCode() *int32
	SetBody(v *GenerateTrafficControlTaskCodeResponseBody) *GenerateTrafficControlTaskCodeResponse
	GetBody() *GenerateTrafficControlTaskCodeResponseBody
}

type GenerateTrafficControlTaskCodeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateTrafficControlTaskCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateTrafficControlTaskCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateTrafficControlTaskCodeResponse) GoString() string {
	return s.String()
}

func (s *GenerateTrafficControlTaskCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateTrafficControlTaskCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateTrafficControlTaskCodeResponse) GetBody() *GenerateTrafficControlTaskCodeResponseBody {
	return s.Body
}

func (s *GenerateTrafficControlTaskCodeResponse) SetHeaders(v map[string]*string) *GenerateTrafficControlTaskCodeResponse {
	s.Headers = v
	return s
}

func (s *GenerateTrafficControlTaskCodeResponse) SetStatusCode(v int32) *GenerateTrafficControlTaskCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTrafficControlTaskCodeResponse) SetBody(v *GenerateTrafficControlTaskCodeResponseBody) *GenerateTrafficControlTaskCodeResponse {
	s.Body = v
	return s
}

func (s *GenerateTrafficControlTaskCodeResponse) Validate() error {
	return dara.Validate(s)
}
