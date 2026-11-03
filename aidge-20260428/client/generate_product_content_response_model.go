// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateProductContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateProductContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateProductContentResponse
	GetStatusCode() *int32
	SetBody(v *GenerateProductContentResponseBody) *GenerateProductContentResponse
	GetBody() *GenerateProductContentResponseBody
}

type GenerateProductContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateProductContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateProductContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateProductContentResponse) GoString() string {
	return s.String()
}

func (s *GenerateProductContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateProductContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateProductContentResponse) GetBody() *GenerateProductContentResponseBody {
	return s.Body
}

func (s *GenerateProductContentResponse) SetHeaders(v map[string]*string) *GenerateProductContentResponse {
	s.Headers = v
	return s
}

func (s *GenerateProductContentResponse) SetStatusCode(v int32) *GenerateProductContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateProductContentResponse) SetBody(v *GenerateProductContentResponseBody) *GenerateProductContentResponse {
	s.Body = v
	return s
}

func (s *GenerateProductContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
