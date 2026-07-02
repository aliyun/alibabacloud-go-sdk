// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateMacResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateMacResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateMacResponse
	GetStatusCode() *int32
	SetBody(v *GenerateMacResponseBody) *GenerateMacResponse
	GetBody() *GenerateMacResponseBody
}

type GenerateMacResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateMacResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateMacResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateMacResponse) GoString() string {
	return s.String()
}

func (s *GenerateMacResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateMacResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateMacResponse) GetBody() *GenerateMacResponseBody {
	return s.Body
}

func (s *GenerateMacResponse) SetHeaders(v map[string]*string) *GenerateMacResponse {
	s.Headers = v
	return s
}

func (s *GenerateMacResponse) SetStatusCode(v int32) *GenerateMacResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateMacResponse) SetBody(v *GenerateMacResponseBody) *GenerateMacResponse {
	s.Body = v
	return s
}

func (s *GenerateMacResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
