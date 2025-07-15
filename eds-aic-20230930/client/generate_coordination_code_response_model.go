// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCoordinationCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateCoordinationCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateCoordinationCodeResponse
	GetStatusCode() *int32
	SetBody(v *GenerateCoordinationCodeResponseBody) *GenerateCoordinationCodeResponse
	GetBody() *GenerateCoordinationCodeResponseBody
}

type GenerateCoordinationCodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCoordinationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCoordinationCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateCoordinationCodeResponse) GoString() string {
	return s.String()
}

func (s *GenerateCoordinationCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateCoordinationCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateCoordinationCodeResponse) GetBody() *GenerateCoordinationCodeResponseBody {
	return s.Body
}

func (s *GenerateCoordinationCodeResponse) SetHeaders(v map[string]*string) *GenerateCoordinationCodeResponse {
	s.Headers = v
	return s
}

func (s *GenerateCoordinationCodeResponse) SetStatusCode(v int32) *GenerateCoordinationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCoordinationCodeResponse) SetBody(v *GenerateCoordinationCodeResponseBody) *GenerateCoordinationCodeResponse {
	s.Body = v
	return s
}

func (s *GenerateCoordinationCodeResponse) Validate() error {
	return dara.Validate(s)
}
