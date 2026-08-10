// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateYikeLoginTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateYikeLoginTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateYikeLoginTokenResponse
	GetStatusCode() *int32
	SetBody(v *GenerateYikeLoginTokenResponseBody) *GenerateYikeLoginTokenResponse
	GetBody() *GenerateYikeLoginTokenResponseBody
}

type GenerateYikeLoginTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateYikeLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateYikeLoginTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateYikeLoginTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateYikeLoginTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateYikeLoginTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateYikeLoginTokenResponse) GetBody() *GenerateYikeLoginTokenResponseBody {
	return s.Body
}

func (s *GenerateYikeLoginTokenResponse) SetHeaders(v map[string]*string) *GenerateYikeLoginTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateYikeLoginTokenResponse) SetStatusCode(v int32) *GenerateYikeLoginTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateYikeLoginTokenResponse) SetBody(v *GenerateYikeLoginTokenResponseBody) *GenerateYikeLoginTokenResponse {
	s.Body = v
	return s
}

func (s *GenerateYikeLoginTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
