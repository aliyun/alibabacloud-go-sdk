// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDefaultServiceTestConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateDefaultServiceTestConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateDefaultServiceTestConfigResponse
	GetStatusCode() *int32
	SetBody(v *GenerateDefaultServiceTestConfigResponseBody) *GenerateDefaultServiceTestConfigResponse
	GetBody() *GenerateDefaultServiceTestConfigResponseBody
}

type GenerateDefaultServiceTestConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDefaultServiceTestConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateDefaultServiceTestConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateDefaultServiceTestConfigResponse) GoString() string {
	return s.String()
}

func (s *GenerateDefaultServiceTestConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateDefaultServiceTestConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateDefaultServiceTestConfigResponse) GetBody() *GenerateDefaultServiceTestConfigResponseBody {
	return s.Body
}

func (s *GenerateDefaultServiceTestConfigResponse) SetHeaders(v map[string]*string) *GenerateDefaultServiceTestConfigResponse {
	s.Headers = v
	return s
}

func (s *GenerateDefaultServiceTestConfigResponse) SetStatusCode(v int32) *GenerateDefaultServiceTestConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDefaultServiceTestConfigResponse) SetBody(v *GenerateDefaultServiceTestConfigResponseBody) *GenerateDefaultServiceTestConfigResponse {
	s.Body = v
	return s
}

func (s *GenerateDefaultServiceTestConfigResponse) Validate() error {
	return dara.Validate(s)
}
